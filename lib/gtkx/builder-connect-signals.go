// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtkx

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"github.com/pekim/gobbi/lib/gtk"
	"reflect"
	"sync"
	"unsafe"
)

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>

void GtkBuilderConnectSignal(GObject *object, gchar *class_name, gchar *signal_name, gchar *handler_name, gpointer user_data);

static void Gobbi_gtk_builder_connect(
	GtkBuilder *builder,
	GObject *object,
	const gchar *signal_name,
	const gchar *handler_name,
	GObject *connect_object,
	GConnectFlags flags,
	gpointer user_data
) {
	const gchar *class_name = G_OBJECT_CLASS_NAME(G_OBJECT_GET_CLASS(object));
	GtkBuilderConnectSignal(object, (gchar *)(class_name), (gchar *)(signal_name), (gchar *)(handler_name), user_data);
}

static void Gobbi_gtk_builder_connect_signals(GtkBuilder *builder) {
	gtk_builder_connect_signals_full(
		builder,
		Gobbi_gtk_builder_connect,
		builder
	);
}
*/
import "C"

var builderConnectSignalsHandlers = make(map[unsafe.Pointer]map[string]interface{})
var builderConnectSignalsHandlersLock sync.Mutex

func BuilderConnectSignals(builder *gtk.Builder, handlers map[string]interface{}) {
	cBuilder := builder.ToC()

	builderConnectSignalsHandlersLock.Lock()
	builderConnectSignalsHandlers[cBuilder] = handlers
	builderConnectSignalsHandlersLock.Unlock()

	C.Gobbi_gtk_builder_connect_signals(
		(*C.GtkBuilder)(cBuilder),
	)

	builderConnectSignalsHandlersLock.Lock()
	delete(builderConnectSignalsHandlers, cBuilder)
	builderConnectSignalsHandlersLock.Unlock()
}

//export GtkBuilderConnectSignal
func GtkBuilderConnectSignal(cObject *C.GObject, cClassName *C.gchar, cSignalName *C.gchar, cHandlerName *C.gchar, cBuilder C.gpointer) {
	className := C.GoString(cClassName)
	signalName := C.GoString(cSignalName)
	handlerName := C.GoString(cHandlerName)

	handlers, found := builderConnectSignalsHandlers[unsafe.Pointer(cBuilder)]
	if !found {
		panic("builder not found in handlers map")
	}

	handler, found := handlers[handlerName]
	if !found {
		fmt.Println("TODO: Not found handler", handlerName)
		return
	}

	ctorValue, found := gtk.GobjectClassGoTypeMap[className]
	if !found {
		fmt.Println("TODO: Not found class", className)
		return
	}

	ctorArgs := []reflect.Value{reflect.ValueOf(unsafe.Pointer(cObject))}
	ctorReturnValues := ctorValue.Call(ctorArgs)
	gtkInstance := ctorReturnValues[0]

	connectMethodName := "Connect" + generate.MakeExportedGoName(signalName)
	fmt.Println(connectMethodName)
	connectMethod := gtkInstance.MethodByName(connectMethodName)
	if !connectMethod.IsValid() {
		fmt.Println("TODO: Not found connect method", className, connectMethodName)
		return
	}

	connectArgs := []reflect.Value{reflect.ValueOf(handler)}
	connectMethod.Call(connectArgs)
}

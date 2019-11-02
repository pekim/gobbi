// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtkx

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"github.com/pekim/gobbi/lib/gobject"
	"github.com/pekim/gobbi/lib/gtk"
	"reflect"
	"strings"
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

type builderConnectSignalsHandlersWrapper struct {
	handlers map[string]interface{}
	err      error
}

var builderConnectSignalsHandlers = make(map[unsafe.Pointer]builderConnectSignalsHandlersWrapper)
var builderConnectSignalsHandlersLock sync.Mutex

func BuilderConnectSignals(builder *gtk.Builder, handlers map[string]interface{}) error {
	cBuilder := builder.ToC()
	handlersWrapper := builderConnectSignalsHandlersWrapper{
		handlers: handlers,
		err:      nil,
	}

	// add to map
	builderConnectSignalsHandlersLock.Lock()
	builderConnectSignalsHandlers[cBuilder] = handlersWrapper
	builderConnectSignalsHandlersLock.Unlock()

	// connect signals
	C.Gobbi_gtk_builder_connect_signals(
		(*C.GtkBuilder)(cBuilder),
	)

	// remove from map
	builderConnectSignalsHandlersLock.Lock()
	delete(builderConnectSignalsHandlers, cBuilder)
	builderConnectSignalsHandlersLock.Unlock()

	return handlersWrapper.err
}

//export GtkBuilderConnectSignal
func GtkBuilderConnectSignal(
	cObject *C.GObject,
	cClassName *C.gchar,
	cSignalName *C.gchar,
	cHandlerName *C.gchar,
	cBuilder C.gpointer, //user data
) {
	className := C.GoString(cClassName)
	signalName := C.GoString(cSignalName)
	handlerName := C.GoString(cHandlerName)

	handlersWrapper, found := builderConnectSignalsHandlers[unsafe.Pointer(cBuilder)]
	if !found {
		panic("builder not found in handlers map")
	}

	handler, found := handlersWrapper.handlers[handlerName]
	if !found {
		handlersWrapper.err = fmt.Errorf("Handler named '%s' not found", handlerName)
		return
	}

	goType, found := gtk.GobjectClassToGoTypeMetaMap[className]
	if !found {
		handlersWrapper.err = fmt.Errorf("Class with name '%s' not supported for Builder signal connecting", className)
		return
	}

	ctorArgs := []reflect.Value{reflect.ValueOf(unsafe.Pointer(cObject))}
	ctorReturnValues := goType.Ctor.Call(ctorArgs)
	gtkInstance := ctorReturnValues[0]

	// Connect notify signal.
	if strings.HasPrefix(signalName, "notify::") {
		propertyName := strings.TrimPrefix(signalName, "notify::")
		handlersWrapper.err = GtkBuilderConnectNotifySignalSignal(gtkInstance, handler, handlerName, propertyName, className)
		return
	}

	connectMethod, ok := gtkBuilderGetMethod(signalName, gtkInstance, className)
	if !ok {
		for _, ancestor := range goType.Ancestors {
			ancestorMethod := gtkInstance.MethodByName(ancestor.MethodName)
			if !ancestorMethod.IsValid() {
				continue
			}

			gtkInstance = ancestorMethod.Call(nil)[0]

			connectMethod, ok = gtkBuilderGetMethod(signalName, gtkInstance, className)
			if ok {
				break
			}
		}

		if !ok {
			handlersWrapper.err = fmt.Errorf("Signal '%s' not supported for class '%s' or any of its ancestor classes",
				signalName, className)
			return
		}
	}

	param0Type := connectMethod.Type().In(0)
	handlerType := reflect.TypeOf(handler)
	if !handlerType.AssignableTo(param0Type) {
		handlersWrapper.err = fmt.Errorf("Signature mistmatch for handler '%s' for signal '%s' of class '%s'",
			handlerName, signalName, className)
		return
	}

	connectArgs := []reflect.Value{reflect.ValueOf(handler)}
	connectMethod.Call(connectArgs)
}

func gtkBuilderGetMethod(signalName string, gtkInstance reflect.Value, className string) (reflect.Value, bool) {
	connectMethodName := "Connect" + generate.MakeExportedGoName(signalName)
	connectMethod := gtkInstance.MethodByName(connectMethodName)
	if !connectMethod.IsValid() {
		return reflect.Value{}, false
	}

	return connectMethod, true
}

func GtkBuilderConnectNotifySignalSignal(
	gtkInstance reflect.Value,
	handler interface{},
	handlerName string,
	propertyName string,
	className string,
) error {
	expectedHandlerType := reflect.TypeOf(func(*gobject.ParamSpec) {})
	handlerType := reflect.TypeOf(handler)
	if !handlerType.AssignableTo(expectedHandlerType) {
		fmt.Println("TODO: property signal handler of wrong type", propertyName)
		return fmt.Errorf("Signature mistmatch for handler '%s' for property '%s' nofifier  of class '%s'",
			handlerName, propertyName, className)
	}

	objectMethod := gtkInstance.MethodByName("Object")
	objectValue := objectMethod.Call(nil)[0]
	object := objectValue.Interface().(*gobject.Object)
	object.ConnectNotifyProperty(propertyName, handler.(func(*gobject.ParamSpec)))

	return nil
}

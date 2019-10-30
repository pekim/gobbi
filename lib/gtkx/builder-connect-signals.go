// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtkx

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gtk"
)

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>

static void Test_c(GObject *object) {
	const gchar *class_name = G_OBJECT_CLASS_NAME(G_OBJECT_GET_CLASS(object));
	printf("%s\n", class_name);
}

void GtkBuilderConnectSignal(GObject *object, gchar *class_name, gchar *signal_name, gchar *handler_name);

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
	GtkBuilderConnectSignal(object, (gchar *)(class_name), (gchar *)(signal_name), (gchar *)(handler_name));
}

static void Gobbi_gtk_builder_connect_signals(GtkBuilder *builder) {
	gtk_builder_connect_signals_full(
		builder,
		Gobbi_gtk_builder_connect,
		NULL
	);
}
*/
import "C"

func BuilderConnectSignals(builder *gtk.Builder) {
	C.Gobbi_gtk_builder_connect_signals(
		(*C.GtkBuilder)(builder.ToC()),
	)
}

//export GtkBuilderConnectSignal
func GtkBuilderConnectSignal(cObject *C.GObject, cClassName *C.gchar, cSignalName *C.gchar, cHandlerName *C.gchar) {
	className := C.GoString(cClassName)
	signalName := C.GoString(cSignalName)
	handlerName := C.GoString(cHandlerName)

	fmt.Println(cObject, className, signalName, handlerName)
}

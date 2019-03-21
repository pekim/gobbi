// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void callback_builderconnectfuncHandler(GObject *, GtkBuilder*, GObject*, gchar*, gchar*, GObject*, GConnectFlags, gpointer);

*/
import "C"

var callbackBuilderconnectfuncId int
var callbackBuilderconnectfuncMap = make(map[int]BuilderconnectfuncCallback)
var callbackBuilderconnectfuncLock sync.RWMutex

// BuilderconnectfuncCallback is a callback function for a 'BuilderConnectFunc' callback.
type BuilderconnectfuncCallback func(builder *Builder, object *gobject.Object, signalName string, handlerName string, connectObject *gobject.Object, flags gobject.ConnectFlags)

//export callback_builderconnectfuncHandler
func callback_builderconnectfuncHandler(_ *C.GObject, c_builder *C.GtkBuilder, c_object *C.GObject, c_signal_name *C.gchar, c_handler_name *C.gchar, c_connect_object *C.GObject, c_flags C.GConnectFlags, c_user_data C.gpointer) {
	callbackBuilderconnectfuncLock.RLock()
	defer callbackBuilderconnectfuncLock.RUnlock()

	builder := BuilderNewFromC(unsafe.Pointer(c_builder))

	object := gobject.ObjectNewFromC(unsafe.Pointer(c_object))

	signalName := C.GoString(c_signal_name)

	handlerName := C.GoString(c_handler_name)

	connectObject := gobject.ObjectNewFromC(unsafe.Pointer(c_connect_object))

	flags := gobject.ConnectFlags(c_flags)

	index := int(uintptr(c_user_data))
	callback := callbackBuilderconnectfuncMap[index]
	callback(builder, object, signalName, handlerName, connectObject, flags)
}

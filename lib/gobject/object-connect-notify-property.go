package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
/*

	void object_notifyHandler(GObject *, GParamSpec *, gpointer);

	static gulong Object_signal_connect_notify_property(gpointer instance, char *signal, gpointer data) {
		return g_signal_connect(instance, signal, G_CALLBACK(object_notifyHandler), data);
	}

*/
import "C"

import "unsafe"

/*
ConnectNotifyProperty connects the callback to the 'notify' signal
for a specific property of the Object.

For example, to connect to the notify signal for the "title" property
of a window.

	window.Object().ConnectNotifyProperty("title", func(pspec *gobject.ParamSpec) {
		fmt.Println("title property changed")
	})

In that example it is the "notify::title" signal that is being connected to.

The returned value represents the connection, and may be passed to DisconnectNotify to remove it.
*/
func (recv *Object) ConnectNotifyProperty(property string, callback ObjectSignalNotifyCallback) int {
	signalObjectNotifyLock.Lock()
	defer signalObjectNotifyLock.Unlock()

	signalObjectNotifyId++
	instance := C.gpointer(recv.native)

	signal := C.CString("notify::" + property)
	defer C.free(unsafe.Pointer(signal))

	handlerID := C.Object_signal_connect_notify_property(instance, signal, C.gpointer(uintptr(signalObjectNotifyId)))

	detail := signalObjectNotifyDetail{callback, handlerID}
	signalObjectNotifyMap[signalObjectNotifyId] = detail

	return signalObjectNotifyId
}

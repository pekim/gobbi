package main

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
#include <stdlib.h>

extern void destroyHandler();

static void signal_connect_destroy(gpointer instance,
                       const gchar *detailed_signal,
                       gpointer data) {
	g_signal_connect_data(instance, detailed_signal, destroyHandler, data, (GClosureNotify)NULL, 0);
}

*/
import "C"

import (
	"github.com/pekim/gobbi/lib/gtk"
	"sync"
	"unsafe"
)

var signalDestroyId int
var signalDestroyMap = make(map[int]DestroyCallback)
var signalDestroyLock sync.Mutex

type DestroyCallback func()

func connectDestroy(target *gtk.Widget, callback DestroyCallback) {
	signalDestroyLock.Lock()
	defer signalDestroyLock.Unlock()

	signalDestroyId++
	signalDestroyMap[signalDestroyId] = callback

	instance := C.gpointer(target.Object().ToC())

	detail := C.CString("destroy")
	defer C.free(unsafe.Pointer(detail))

	C.signal_connect_destroy(instance, detail,
		C.gpointer(uintptr(signalDestroyId)))
}

//export destroyHandler
func destroyHandler(target *C.GObject, data uintptr) {
	cb := signalDestroyMap[int(data)]
	cb()
}

//
//type KeyPressEventCallback func(evnt *gdk.EventKey)
//
//func connectKeyPressEvent(object *gobject.Object, callback KeyPressEventCallback) {
//	marshallCallback := func(event unsafe.Pointer) {
//		callback(gdk.EventKeyNewFromC(event))
//	}
//
//	connect(object, "key-press-event", marshallCallback)
//}

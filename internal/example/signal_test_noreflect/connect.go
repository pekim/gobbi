package main

/*
#cgo pkg-config: gobject-2.0
#cgo pkg-config: gdk-3.0
#include <glib-object.h>
#include <gdk/gdk.h>
#include <stdlib.h>

extern void destroyHandler();

static void signal_connect_destroy(gpointer instance, gpointer data) {
	g_signal_connect_data(instance, "destroy", destroyHandler, data, (GClosureNotify)NULL, 0);
}

extern void keyPressEventHandler();

static void signal_connect_key_press_event(gpointer instance, gpointer data) {
	g_signal_connect_data(instance, "key-press-event", keyPressEventHandler, data, (GClosureNotify)NULL, 0);
}

*/
import "C"

import (
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"sync"
	"unsafe"
)

// --------------------

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

	C.signal_connect_destroy(instance, C.gpointer(uintptr(signalDestroyId)))
}

//export destroyHandler
func destroyHandler(target *C.GObject, data uintptr) {
	cb := signalDestroyMap[int(data)]
	cb()
}

// --------------------

var signalKeyPressEventId int
var signalKeyPressEventMap = make(map[int]KeyPressEventCallback)
var signalKeyPressEventLock sync.Mutex

type KeyPressEventCallback func(event *gdk.EventKey)

func connectKeyPressEvent(target *gtk.Widget, callback KeyPressEventCallback) {
	signalKeyPressEventLock.Lock()
	defer signalKeyPressEventLock.Unlock()

	signalKeyPressEventId++
	signalKeyPressEventMap[signalKeyPressEventId] = callback

	instance := C.gpointer(target.Object().ToC())

	C.signal_connect_key_press_event(instance, C.gpointer(uintptr(signalKeyPressEventId)))
}

//export keyPressEventHandler
func keyPressEventHandler(target *C.GObject, event *C.GdkEventKey, data uintptr) {
	goEvent := gdk.EventKeyNewFromC(unsafe.Pointer(event))

	cb := signalKeyPressEventMap[int(data)]
	cb(goEvent)
}

package main

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
#include <stdio.h>

extern void goMarshal(GClosure *closure, GValue *retValue ,
	guint nParams, GValue *params,
	gpointer invocationHint, GValue *marshalData);

static inline GClosure* connect(gpointer instance,
						   const gchar *detailed_signal) {
	GClosure *closure;

	closure = g_closure_new_simple(sizeof(GClosure), NULL);
	g_closure_set_marshal(closure, (GClosureMarshal)(goMarshal));
	g_signal_connect_closure(instance, detailed_signal, closure, FALSE);

	return closure;
}

*/
import "C"
import (
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

var closures = make(map[*C.GClosure]interface{})
var closureLock sync.Mutex

func connect(object *gobject.Object, detail string, callback interface{}) {
	closureLock.Lock()
	defer closureLock.Unlock()

	closure := C.connect(C.gpointer(object.ToC()), C.CString(detail))
	closures[closure] = callback
}

// -----------------------

type DestroyCallback func()

func connectDestroy(object *gobject.Object, callback DestroyCallback) {
	marshallCallback := func() {
		callback()
	}

	connect(object, "destroy", marshallCallback)
}

type KeyPressEventCallback func(evnt *gdk.EventKey)

func connectKeyPressEvent(object *gobject.Object, callback KeyPressEventCallback) {
	marshallCallback := func(event unsafe.Pointer) {
		callback(gdk.EventKeyNewFromC(event))
	}

	connect(object, "key-press-event", marshallCallback)
}

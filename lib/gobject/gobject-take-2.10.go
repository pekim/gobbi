// +build gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"
import (
	"runtime"
	"unsafe"
)

func TakeRef(goObj interface{}, cPointer unsafe.Pointer) {
	gpointer := C.gpointer(cPointer)

	if C.g_object_is_floating(gpointer) == C.TRUE {
		C.g_object_ref_sink(gpointer)
	} else {
		C.g_object_ref(gpointer)
	}

	runtime.SetFinalizer(goObj, func(o interface{}) {
		C.g_object_unref(gpointer)
	})
}

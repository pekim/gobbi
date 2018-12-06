// +build !gobject_2.10,!gobject_2.12,!gobject_2.18,!gobject_2.24,!gobject_2.26,!gobject_2.28,!gobject_2.30,!gobject_2.32,!gobject_2.34,!gobject_2.36,!gobject_2.38,!gobject_2.42,!gobject_2.44,!gobject_2.46,!gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

func (recv *Object) Take() {
	// Do nothing as g_object_is_floating and g_object_ref_sink are not supported until gobject 2.10.
}

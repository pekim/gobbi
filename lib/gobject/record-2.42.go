// This is a generated file - DO NOT EDIT
// +build gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// InitFromInstance is a wrapper around the C function g_value_init_from_instance.
func (recv *Value) InitFromInstance(instance uintptr) {
	c_instance := (C.gpointer)(instance)

	C.g_value_init_from_instance((*C.GValue)(recv.native), c_instance)

	return
}

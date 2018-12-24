// This is a generated file - DO NOT EDIT
// +build gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypeGetInstanceCount is a wrapper around the C function g_type_get_instance_count.
func TypeGetInstanceCount(type_ Type) int32 {
	c_type := (C.GType)(type_)

	retC := C.g_type_get_instance_count(c_type)
	retGo := (int32)(retC)

	return retGo
}

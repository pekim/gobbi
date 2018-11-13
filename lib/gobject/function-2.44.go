// This is a generated file - DO NOT EDIT
// +build gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Returns the number of instances allocated of the particular type;
// this is only available if GLib is built with debugging support and
// the instance_count debug flag is set (by setting the GOBJECT_DEBUG
// variable to include instance-count).
/*

C function

g_type_get_instance_count
*/
func TypeGetInstanceCount(type_ Type) int32 {
	c_type := (C.GType)(type_)

	retC := C.g_type_get_instance_count(c_type)
	retGo := (int32)(retC)

	return retGo
}

// This is a generated file - DO NOT EDIT
// +build gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Gets the offset of the private data for instances of @g_class.
//
// This is how many bytes you should add to the instance pointer of a
// class in order to get the private data for the type represented by
// @g_class.
//
// You can only call this function after you have registered a private
// data area for @g_class using g_type_class_add_private().
/*

C function : g_type_class_get_instance_private_offset
*/
func (recv *TypeClass) GetInstancePrivateOffset() int32 {
	retC := C.g_type_class_get_instance_private_offset((C.gpointer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

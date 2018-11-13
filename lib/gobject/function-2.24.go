// This is a generated file - DO NOT EDIT
// +build gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Registers a private class structure for a classed type;
// when the class is allocated, the private structures for
// the class and all of its parent types are allocated
// sequentially in the same memory block as the public
// structures, and are zero-filled.
//
// This function should be called in the
// type's get_type() function after the type is registered.
// The private structure can be retrieved using the
// G_TYPE_CLASS_GET_PRIVATE() macro.
/*

C function

g_type_add_class_private
*/
func TypeAddClassPrivate(classType Type, privateSize uint64) {
	c_class_type := (C.GType)(classType)

	c_private_size := (C.gsize)(privateSize)

	C.g_type_add_class_private(c_class_type, c_private_size)

	return
}

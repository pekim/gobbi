// This is a generated file - DO NOT EDIT
// +build gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Blacklisted : g_value_get_schar

// Blacklisted : g_value_set_schar

// Blacklisted : g_weak_ref_clear

// Blacklisted : g_weak_ref_get

// Init is a wrapper around the C function g_weak_ref_init.
func (recv *WeakRef) Init(object *Object) {
	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	C.g_weak_ref_init((*C.GWeakRef)(recv.native), c_object)

	return
}

// Blacklisted : g_weak_ref_set

// This is a generated file - DO NOT EDIT
// +build gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unbind is a wrapper around the C function g_binding_unbind.
func (recv *Binding) Unbind() {
	C.g_binding_unbind((*C.GBinding)(recv.native))

	return
}

// Unsupported signal : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : g_object_new : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_valist : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_with_properties : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_newv : unsupported parameter object_type : no type generator for GType, GType

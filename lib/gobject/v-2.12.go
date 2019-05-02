// Code generated - DO NOT EDIT.
// +build gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_type_register_static_simple : unsupported parameter class_init : no type generator for ClassInitFunc (GClassInitFunc) for param class_init

// GetGtype is a wrapper around the C function g_value_get_gtype.
func (recv *Value) GetGtype() Type {
	retC := C.g_value_get_gtype((*C.GValue)(recv.native))
	retGo := (Type)(retC)

	return retGo
}

// SetGtype is a wrapper around the C function g_value_set_gtype.
func (recv *Value) SetGtype(vGtype Type) {
	c_v_gtype := (C.GType)(vGtype)

	C.g_value_set_gtype((*C.GValue)(recv.native), c_v_gtype)

	return
}

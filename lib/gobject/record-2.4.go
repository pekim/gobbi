// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// OverrideProperty is a wrapper around the C function g_object_class_override_property.
func (recv *ObjectClass) OverrideProperty(propertyId uint32, name string) {
	c_property_id := (C.guint)(propertyId)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_object_class_override_property((*C.GObjectClass)(recv.native), c_property_id, c_name)

	return
}

// TypeClassPeekStatic is a wrapper around the C function g_type_class_peek_static.
func TypeClassPeekStatic(type_ Type) TypeClass {
	c_type := (C.GType)(type_)

	retC := C.g_type_class_peek_static(c_type)
	retGo := *TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddPrivate is a wrapper around the C function g_type_class_add_private.
func (recv *TypeClass) AddPrivate(privateSize uint64) {
	c_private_size := (C.gsize)(privateSize)

	C.g_type_class_add_private((C.gpointer)(recv.native), c_private_size)

	return
}

// Unsupported : g_value_take_boxed : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Unsupported : g_value_take_object : unsupported parameter v_object : no type generator for gpointer (gpointer) for param v_object

// TakeParam is a wrapper around the C function g_value_take_param.
func (recv *Value) TakeParam(param *ParamSpec) {
	c_param := (*C.GParamSpec)(C.NULL)
	if param != nil {
		c_param = (*C.GParamSpec)(param.ToC())
	}

	C.g_value_take_param((*C.GValue)(recv.native), c_param)

	return
}

// TakeString is a wrapper around the C function g_value_take_string.
func (recv *Value) TakeString(vString string) {
	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	C.g_value_take_string((*C.GValue)(recv.native), c_v_string)

	return
}

// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GCClosure

// Blacklisted : GClosure

// Blacklisted : GClosureNotifyData

// Blacklisted : GEnumClass

// Blacklisted : GEnumValue

// Blacklisted : GFlagsClass

// Blacklisted : GFlagsValue

// Blacklisted : GInitiallyUnownedClass

// Blacklisted : GInterfaceInfo

// Blacklisted : GObjectClass

// Blacklisted : GObjectConstructParam

// Blacklisted : GParamSpecClass

// Blacklisted : GParamSpecPool

// Blacklisted : GParamSpecTypeInfo

// Blacklisted : GParameter

// Blacklisted : GSignalInvocationHint

// Blacklisted : GSignalQuery

// Blacklisted : GTypeClass

// Blacklisted : GTypeFundamentalInfo

// Blacklisted : GTypeInfo

// Blacklisted : GTypeInstance

// Blacklisted : GTypeInterface

// Blacklisted : GTypeModuleClass

// Blacklisted : GTypePluginClass

// Blacklisted : GTypeQuery

// Blacklisted : GTypeValueTable

// Value is a wrapper around the C record GValue.
type Value struct {
	native *C.GValue
	// Private : g_type
	// no type for data
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.GValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Value with another Value, and returns true if they represent the same GObject.
func (recv *Value) Equals(other *Value) bool {
	return other.ToC() == recv.ToC()
}

// g_value_register_transform_func : unsupported parameter transform_func : no type generator for ValueTransform (GValueTransform) for param transform_func
// Blacklisted : g_value_type_compatible

// Blacklisted : g_value_type_transformable

// Blacklisted : g_value_copy

// Unsupported : g_value_dup_boxed : no return generator

// Blacklisted : g_value_dup_object

// Blacklisted : g_value_dup_param

// Blacklisted : g_value_dup_string

// Blacklisted : g_value_fits_pointer

// Blacklisted : g_value_get_boolean

// Unsupported : g_value_get_boxed : no return generator

// Blacklisted : g_value_get_char

// Blacklisted : g_value_get_double

// Blacklisted : g_value_get_enum

// Blacklisted : g_value_get_flags

// Blacklisted : g_value_get_float

// Blacklisted : g_value_get_int

// Blacklisted : g_value_get_int64

// Blacklisted : g_value_get_long

// Blacklisted : g_value_get_object

// Blacklisted : g_value_get_param

// Unsupported : g_value_get_pointer : no return generator

// Blacklisted : g_value_get_string

// Blacklisted : g_value_get_uchar

// Blacklisted : g_value_get_uint

// Blacklisted : g_value_get_uint64

// Blacklisted : g_value_get_ulong

// Init is a wrapper around the C function g_value_init.
func (recv *Value) Init(gType Type) *Value {
	c_g_type := (C.GType)(gType)

	retC := C.g_value_init((*C.GValue)(recv.native), c_g_type)
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_value_peek_pointer : no return generator

// Blacklisted : g_value_reset

// Blacklisted : g_value_set_boolean

// Unsupported : g_value_set_boxed : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Unsupported : g_value_set_boxed_take_ownership : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Blacklisted : g_value_set_char

// Blacklisted : g_value_set_double

// Blacklisted : g_value_set_enum

// Blacklisted : g_value_set_flags

// Blacklisted : g_value_set_float

// Unsupported : g_value_set_instance : unsupported parameter instance : no type generator for gpointer (gpointer) for param instance

// Blacklisted : g_value_set_int

// Blacklisted : g_value_set_int64

// Blacklisted : g_value_set_long

// Blacklisted : g_value_set_object

// Unsupported : g_value_set_object_take_ownership : unsupported parameter v_object : no type generator for gpointer (gpointer) for param v_object

// Blacklisted : g_value_set_param

// Blacklisted : g_value_set_param_take_ownership

// Unsupported : g_value_set_pointer : unsupported parameter v_pointer : no type generator for gpointer (gpointer) for param v_pointer

// Unsupported : g_value_set_static_boxed : unsupported parameter v_boxed : no type generator for gpointer (gconstpointer) for param v_boxed

// Blacklisted : g_value_set_static_string

// Blacklisted : g_value_set_string

// Blacklisted : g_value_set_string_take_ownership

// Blacklisted : g_value_set_uchar

// Blacklisted : g_value_set_uint

// Blacklisted : g_value_set_uint64

// Blacklisted : g_value_set_ulong

// Blacklisted : g_value_transform

// Blacklisted : g_value_unset

// Blacklisted : GValueArray

// Blacklisted : GWeakRef

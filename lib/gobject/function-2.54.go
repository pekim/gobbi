// This is a generated file - DO NOT EDIT
// +build gobject_2.54

package gobject

import (
	"C"
	"unsafe"
)

// EnumToString is a wrapper around the C function g_enum_to_string.
func EnumToString(gEnumType Type, value int32) string {
	c_g_enum_type := (C.GType)(gEnumType)

	c_value := (C.gint)(value)

	retC := C.g_enum_to_string(c_g_enum_type, c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FlagsToString is a wrapper around the C function g_flags_to_string.
func FlagsToString(flagsType Type, value uint32) string {
	c_flags_type := (C.GType)(flagsType)

	c_value := (C.guint)(value)

	retC := C.g_flags_to_string(c_flags_type, c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

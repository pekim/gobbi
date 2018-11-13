// This is a generated file - DO NOT EDIT
// +build gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Pretty-prints @value in the form of the enum’s name.
//
// This is intended to be used for debugging purposes. The format of the output
// may change in the future.
/*

C function : g_enum_to_string
*/
func EnumToString(gEnumType Type, value int32) string {
	c_g_enum_type := (C.GType)(gEnumType)

	c_value := (C.gint)(value)

	retC := C.g_enum_to_string(c_g_enum_type, c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Pretty-prints @value in the form of the flag names separated by ` | ` and
// sorted. Any extra bits will be shown at the end as a hexadecimal number.
//
// This is intended to be used for debugging purposes. The format of the output
// may change in the future.
/*

C function : g_flags_to_string
*/
func FlagsToString(flagsType Type, value uint32) string {
	c_flags_type := (C.GType)(flagsType)

	c_value := (C.guint)(value)

	retC := C.g_flags_to_string(c_flags_type, c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

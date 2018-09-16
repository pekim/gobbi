// This is a generated file - DO NOT EDIT
// +build glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Utf8MakeValid is a wrapper around the C function g_utf8_make_valid.
func Utf8MakeValid(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_make_valid(c_str, c_len)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_uuid_string_is_valid : no return generator

// UuidStringRandom is a wrapper around the C function g_uuid_string_random.
func UuidStringRandom() string {
	retC := C.g_uuid_string_random()
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

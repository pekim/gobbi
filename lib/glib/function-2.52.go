// +build glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Utf8MakeValid is a wrapper around the C function g_utf8_make_valid.
func Utf8MakeValid(str string, len int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	C.g_utf8_make_valid()
}

// UuidStringIsValid is a wrapper around the C function g_uuid_string_is_valid.
func UuidStringIsValid(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.g_uuid_string_is_valid()
}

// UuidStringRandom is a wrapper around the C function g_uuid_string_random.
func UuidStringRandom() {
	C.g_uuid_string_random()
}

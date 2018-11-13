// This is a generated file - DO NOT EDIT
// +build glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// If the provided string is valid UTF-8, return a copy of it. If not,
// return a copy in which bytes that could not be interpreted as valid Unicode
// are replaced with the Unicode replacement character (U+FFFD).
//
// For example, this is an appropriate function to use if you have received
// a string that was incorrectly declared to be UTF-8, and you need a valid
// UTF-8 version of it that can be logged or displayed to the user, with the
// assumption that it is close enough to ASCII or UTF-8 to be mostly
// readable as-is.
/*

C function

g_utf8_make_valid
*/
func Utf8MakeValid(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_make_valid(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Parses the string @str and verify if it is a UUID.
//
// The function accepts the following syntax:
//
// - simple forms (e.g. `f81d4fae-7dec-11d0-a765-00a0c91e6bf6`)
//
// Note that hyphens are required within the UUID string itself,
// as per the aforementioned RFC.
/*

C function

g_uuid_string_is_valid
*/
func UuidStringIsValid(str string) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.g_uuid_string_is_valid(c_str)
	retGo := retC == C.TRUE

	return retGo
}

// Generates a random UUID (RFC 4122 version 4) as a string.
/*

C function

g_uuid_string_random
*/
func UuidStringRandom() string {
	retC := C.g_uuid_string_random()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

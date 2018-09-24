// This is a generated file - DO NOT EDIT
// +build glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// StrvContains is a wrapper around the C function g_strv_contains.
func StrvContains(strv string, str string) bool {
	c_strv := C.CString(strv)
	defer C.free(unsafe.Pointer(c_strv))

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.g_strv_contains(c_strv, c_str)
	retGo := retC == C.TRUE

	return retGo
}

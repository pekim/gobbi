// +build glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// StrvContains is a wrapper around the C function g_strv_contains.
func StrvContains(strv string, str string) {
	c_strv := C.CString(strv)
	defer C.free(unsafe.Pointer(c_strv))

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.g_strv_contains()
}

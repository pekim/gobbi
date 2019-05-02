// This is a generated file - DO NOT EDIT
// +build glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_build_filename_valist : unsupported parameter args : no type generator for va_list (va_list*) for param args

// Unsupported : g_clear_handle_id : unsupported parameter clear_func : no type generator for ClearHandleFunc (GClearHandleFunc) for param clear_func

// DateTimeNewFromIso8601 is a wrapper around the C function g_date_time_new_from_iso8601.
func DateTimeNewFromIso8601(text string, defaultTz *TimeZone) *DateTime {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_default_tz := (*C.GTimeZone)(C.NULL)
	if defaultTz != nil {
		c_default_tz = (*C.GTimeZone)(defaultTz.ToC())
	}

	retC := C.g_date_time_new_from_iso8601(c_text, c_default_tz)
	var retGo (*DateTime)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DateTimeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

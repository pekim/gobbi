// This is a generated file - DO NOT EDIT
// +build glib_2.54 glib_2.56

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

// AsciiStringToSigned is a wrapper around the C function g_ascii_string_to_signed.
func AsciiStringToSigned(str string, base uint32, min int64, max int64) (bool, int64, error) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_base := (C.guint)(base)

	c_min := (C.gint64)(min)

	c_max := (C.gint64)(max)

	var c_out_num C.gint64

	var cThrowableError *C.GError

	retC := C.g_ascii_string_to_signed(c_str, c_base, c_min, c_max, &c_out_num, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outNum := (int64)(c_out_num)

	return retGo, outNum, goError
}

// AsciiStringToUnsigned is a wrapper around the C function g_ascii_string_to_unsigned.
func AsciiStringToUnsigned(str string, base uint32, min uint64, max uint64) (bool, uint64, error) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_base := (C.guint)(base)

	c_min := (C.guint64)(min)

	c_max := (C.guint64)(max)

	var c_out_num C.guint64

	var cThrowableError *C.GError

	retC := C.g_ascii_string_to_unsigned(c_str, c_base, c_min, c_max, &c_out_num, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outNum := (uint64)(c_out_num)

	return retGo, outNum, goError
}

// Unsupported : g_ptr_array_find : unsupported parameter haystack : no type generator for gpointer (gpointer) for array param haystack

// Unsupported : g_ptr_array_find_with_equal_func : unsupported parameter haystack : no type generator for gpointer (gpointer) for array param haystack

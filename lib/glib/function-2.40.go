// +build glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// StrIsAscii is a wrapper around the C function g_str_is_ascii.
func StrIsAscii(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.g_str_is_ascii(c_str)
}

// Unsupported : g_str_match_string : unsupported parameter accept_alternates : type gboolean, gboolean

// StrToAscii is a wrapper around the C function g_str_to_ascii.
func StrToAscii(str string, fromLocale string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_from_locale := C.CString(fromLocale)
	defer C.free(unsafe.Pointer(c_from_locale))

	C.g_str_to_ascii(c_str, c_from_locale)
}

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : no type

// Unsupported : g_variant_parse_error_print_context : unsupported parameter error : type Error, GError*

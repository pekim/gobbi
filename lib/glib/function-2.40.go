// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// StrIsAscii is a wrapper around the C function g_str_is_ascii.
func StrIsAscii(str string) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.g_str_is_ascii(c_str)
	retGo := retC == C.TRUE

	return retGo
}

// StrMatchString is a wrapper around the C function g_str_match_string.
func StrMatchString(searchTerm string, potentialHit string, acceptAlternates bool) bool {
	c_search_term := C.CString(searchTerm)
	defer C.free(unsafe.Pointer(c_search_term))

	c_potential_hit := C.CString(potentialHit)
	defer C.free(unsafe.Pointer(c_potential_hit))

	c_accept_alternates :=
		boolToGboolean(acceptAlternates)

	retC := C.g_str_match_string(c_search_term, c_potential_hit, c_accept_alternates)
	retGo := retC == C.TRUE

	return retGo
}

// StrToAscii is a wrapper around the C function g_str_to_ascii.
func StrToAscii(str string, fromLocale string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_from_locale := C.CString(fromLocale)
	defer C.free(unsafe.Pointer(c_from_locale))

	retC := C.g_str_to_ascii(c_str, c_from_locale)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : output array param ascii_alternates

// VariantParseErrorPrintContext is a wrapper around the C function g_variant_parse_error_print_context.
func VariantParseErrorPrintContext(error *Error, sourceStr string) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	c_source_str := C.CString(sourceStr)
	defer C.free(unsafe.Pointer(c_source_str))

	retC := C.g_variant_parse_error_print_context(c_error, c_source_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

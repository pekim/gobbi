// This is a generated file - DO NOT EDIT
// +build glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// A convenience function for converting a string to a signed number.
//
// This function assumes that @str contains only a number of the given
// @base that is within inclusive bounds limited by @min and @max. If
// this is true, then the converted number is stored in @out_num. An
// empty string is not a valid input. A string with leading or
// trailing whitespace is also an invalid input.
//
// @base can be between 2 and 36 inclusive. Hexadecimal numbers must
// not be prefixed with "0x" or "0X". Such a problem does not exist
// for octal numbers, since they were usually prefixed with a zero
// which does not change the value of the parsed number.
//
// Parsing failures result in an error with the %G_NUMBER_PARSER_ERROR
// domain. If the input is invalid, the error code will be
// %G_NUMBER_PARSER_ERROR_INVALID. If the parsed number is out of
// bounds - %G_NUMBER_PARSER_ERROR_OUT_OF_BOUNDS.
//
// See g_ascii_strtoll() if you have more complex needs such as
// parsing a string which starts with a number, but then has other
// characters.
/*

C function : g_ascii_string_to_signed
*/
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

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outNum := (int64)(c_out_num)

	return retGo, outNum, goThrowableError
}

// A convenience function for converting a string to an unsigned number.
//
// This function assumes that @str contains only a number of the given
// @base that is within inclusive bounds limited by @min and @max. If
// this is true, then the converted number is stored in @out_num. An
// empty string is not a valid input. A string with leading or
// trailing whitespace is also an invalid input.
//
// @base can be between 2 and 36 inclusive. Hexadecimal numbers must
// not be prefixed with "0x" or "0X". Such a problem does not exist
// for octal numbers, since they were usually prefixed with a zero
// which does not change the value of the parsed number.
//
// Parsing failures result in an error with the %G_NUMBER_PARSER_ERROR
// domain. If the input is invalid, the error code will be
// %G_NUMBER_PARSER_ERROR_INVALID. If the parsed number is out of
// bounds - %G_NUMBER_PARSER_ERROR_OUT_OF_BOUNDS.
//
// See g_ascii_strtoull() if you have more complex needs such as
// parsing a string which starts with a number, but then has other
// characters.
/*

C function : g_ascii_string_to_unsigned
*/
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

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outNum := (uint64)(c_out_num)

	return retGo, outNum, goThrowableError
}

// Checks whether @needle exists in @haystack. If the element is found, %TRUE is
// returned and the element’s index is returned in @index_ (if non-%NULL).
// Otherwise, %FALSE is returned and @index_ is undefined. If @needle exists
// multiple times in @haystack, the index of the first instance is returned.
//
// This does pointer comparisons only. If you want to use more complex equality
// checks, such as string comparisons, use g_ptr_array_find_with_equal_func().
/*

C function : g_ptr_array_find
*/
func PtrArrayFind(haystack []uintptr, needle uintptr) (bool, uint32) {
	c_haystack := &haystack[0]

	c_needle := (C.gconstpointer)(needle)

	var c_index_ C.guint

	retC := C.g_ptr_array_find((*C.GPtrArray)(unsafe.Pointer(c_haystack)), c_needle, &c_index_)
	retGo := retC == C.TRUE

	index := (uint32)(c_index_)

	return retGo, index
}

// Unsupported : g_ptr_array_find_with_equal_func : unsupported parameter equal_func : no type generator for EqualFunc (GEqualFunc) for param equal_func

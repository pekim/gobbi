// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_checksum_type_get_length : unsupported parameter checksum_type : no param type for ChecksumType, GChecksumType

// Unsupported : g_compute_checksum_for_data : unsupported parameter checksum_type : no param type for ChecksumType, GChecksumType

// Unsupported : g_compute_checksum_for_string : unsupported parameter checksum_type : no param type for ChecksumType, GChecksumType

// Dpgettext is a wrapper around the C function g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgctxtid := C.CString(msgctxtid)
	defer C.free(unsafe.Pointer(c_msgctxtid))

	c_msgidoffset := (C.gsize)(msgidoffset)

	retC := C.g_dpgettext(c_domain, c_msgctxtid, c_msgidoffset)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FormatSizeForDisplay is a wrapper around the C function g_format_size_for_display.
func FormatSizeForDisplay(size uint64) string {
	c_size := (C.goffset)(size)

	retC := C.g_format_size_for_display(c_size)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in for string with indirection level of 2

// Unsupported : g_prefix_error : unsupported parameter err : no param type for Error, GError**

// Unsupported : g_propagate_prefixed_error : unsupported parameter dest : no param type for Error, GError**

// Strcmp0 is a wrapper around the C function g_strcmp0.
func Strcmp0(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_strcmp0(c_str1, c_str2)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_test_add_data_func : unsupported parameter test_func : no param type for TestDataFunc, GTestDataFunc

// Unsupported : g_test_add_func : unsupported parameter test_func : no param type for TestFunc, GTestFunc

// Unsupported : g_test_bug : no return type

// Unsupported : g_test_bug_base : no return type

// Unsupported : g_test_create_case : unsupported parameter data_setup : no param type for TestFixtureFunc, GTestFixtureFunc

// Unsupported : g_test_create_suite : no return type

// Unsupported : g_test_get_root : no return type

// Unsupported : g_test_init : unsupported parameter argc : no param type for gint, int*

// Unsupported : g_test_maximized_result : unsupported parameter ... : varargs

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no param type for DestroyNotify, GDestroyNotify

// Unsupported : g_test_queue_free : no return type

// TestRandDouble is a wrapper around the C function g_test_rand_double.
func TestRandDouble() float64 {
	retC := C.g_test_rand_double()
	retGo :=
		(float64)(retC)

	return retGo
}

// TestRandDoubleRange is a wrapper around the C function g_test_rand_double_range.
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) float64 {
	c_range_start := (C.double)(rangeStart)

	c_range_end := (C.double)(rangeEnd)

	retC := C.g_test_rand_double_range(c_range_start, c_range_end)
	retGo :=
		(float64)(retC)

	return retGo
}

// TestRandInt is a wrapper around the C function g_test_rand_int.
func TestRandInt() int32 {
	retC := C.g_test_rand_int()
	retGo :=
		(int32)(retC)

	return retGo
}

// TestRandIntRange is a wrapper around the C function g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_test_rand_int_range(c_begin, c_end)
	retGo :=
		(int32)(retC)

	return retGo
}

// TestRun is a wrapper around the C function g_test_run.
func TestRun() int32 {
	retC := C.g_test_run()
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_test_run_suite : unsupported parameter suite : no param type for TestSuite, GTestSuite*

// TestTimerElapsed is a wrapper around the C function g_test_timer_elapsed.
func TestTimerElapsed() float64 {
	retC := C.g_test_timer_elapsed()
	retGo :=
		(float64)(retC)

	return retGo
}

// TestTimerLast is a wrapper around the C function g_test_timer_last.
func TestTimerLast() float64 {
	retC := C.g_test_timer_last()
	retGo :=
		(float64)(retC)

	return retGo
}

// Unsupported : g_test_timer_start : no return type

// Unsupported : g_test_trap_fork : unsupported parameter test_trap_flags : no param type for TestTrapFlags, GTestTrapFlags

// Unsupported : g_test_trap_has_passed : no return type

// Unsupported : g_test_trap_reached_timeout : no return type

// Unsupported : g_uri_escape_string : unsupported parameter allow_utf8 : no param type for gboolean, gboolean

// UriParseScheme is a wrapper around the C function g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_uri_parse_scheme(c_uri)
	retGo :=
		C.GoString(retC)

	return retGo
}

// UriUnescapeSegment is a wrapper around the C function g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_escaped_string_end := C.CString(escapedStringEnd)
	defer C.free(unsafe.Pointer(c_escaped_string_end))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_segment(c_escaped_string, c_escaped_string_end, c_illegal_characters)
	retGo :=
		C.GoString(retC)

	return retGo
}

// UriUnescapeString is a wrapper around the C function g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_string(c_escaped_string, c_illegal_characters)
	retGo :=
		C.GoString(retC)

	return retGo
}

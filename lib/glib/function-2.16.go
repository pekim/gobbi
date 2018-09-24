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

// ChecksumTypeGetLength is a wrapper around the C function g_checksum_type_get_length.
func ChecksumTypeGetLength(checksumType ChecksumType) int64 {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_type_get_length(c_checksum_type)
	retGo := (int64)(retC)

	return retGo
}

// Unsupported : g_compute_checksum_for_data : unsupported parameter data : no param type

// ComputeChecksumForString is a wrapper around the C function g_compute_checksum_for_string.
func ComputeChecksumForString(checksumType ChecksumType, str string, length int64) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(length)

	retC := C.g_compute_checksum_for_string(c_checksum_type, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Dpgettext is a wrapper around the C function g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgctxtid := C.CString(msgctxtid)
	defer C.free(unsafe.Pointer(c_msgctxtid))

	c_msgidoffset := (C.gsize)(msgidoffset)

	retC := C.g_dpgettext(c_domain, c_msgctxtid, c_msgidoffset)
	retGo := C.GoString(retC)

	return retGo
}

// FormatSizeForDisplay is a wrapper around the C function g_format_size_for_display.
func FormatSizeForDisplay(size uint64) string {
	c_size := (C.goffset)(size)

	retC := C.g_format_size_for_display(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// Unsupported : g_prefix_error : unsupported parameter err : in string with indirection level of 2

// Unsupported : g_propagate_prefixed_error : unsupported parameter dest : in string with indirection level of 2

// Strcmp0 is a wrapper around the C function g_strcmp0.
func Strcmp0(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_strcmp0(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_test_add_data_func : unsupported parameter test_func : no type generator for TestDataFunc, GTestDataFunc

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc, GTestFunc

// Unsupported : g_test_bug : no return generator

// Unsupported : g_test_bug_base : no return generator

// Unsupported : g_test_create_case : unsupported parameter data_setup : no type generator for TestFixtureFunc, GTestFixtureFunc

// TestCreateSuite is a wrapper around the C function g_test_create_suite.
func TestCreateSuite(suiteName string) *TestSuite {
	c_suite_name := C.CString(suiteName)
	defer C.free(unsafe.Pointer(c_suite_name))

	retC := C.g_test_create_suite(c_suite_name)
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TestGetRoot is a wrapper around the C function g_test_get_root.
func TestGetRoot() *TestSuite {
	retC := C.g_test_get_root()
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_test_init : unsupported parameter argc : no type generator for gint, int*

// Unsupported : g_test_maximized_result : unsupported parameter ... : varargs

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_test_queue_free : no return generator

// TestRandDouble is a wrapper around the C function g_test_rand_double.
func TestRandDouble() float64 {
	retC := C.g_test_rand_double()
	retGo := (float64)(retC)

	return retGo
}

// TestRandDoubleRange is a wrapper around the C function g_test_rand_double_range.
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) float64 {
	c_range_start := (C.double)(rangeStart)

	c_range_end := (C.double)(rangeEnd)

	retC := C.g_test_rand_double_range(c_range_start, c_range_end)
	retGo := (float64)(retC)

	return retGo
}

// TestRandInt is a wrapper around the C function g_test_rand_int.
func TestRandInt() int32 {
	retC := C.g_test_rand_int()
	retGo := (int32)(retC)

	return retGo
}

// TestRandIntRange is a wrapper around the C function g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_test_rand_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// TestRun is a wrapper around the C function g_test_run.
func TestRun() int32 {
	retC := C.g_test_run()
	retGo := (int32)(retC)

	return retGo
}

// TestRunSuite is a wrapper around the C function g_test_run_suite.
func TestRunSuite(suite *TestSuite) int32 {
	c_suite := suite.toC()

	retC := C.g_test_run_suite(c_suite)
	retGo := (int32)(retC)

	return retGo
}

// TestTimerElapsed is a wrapper around the C function g_test_timer_elapsed.
func TestTimerElapsed() float64 {
	retC := C.g_test_timer_elapsed()
	retGo := (float64)(retC)

	return retGo
}

// TestTimerLast is a wrapper around the C function g_test_timer_last.
func TestTimerLast() float64 {
	retC := C.g_test_timer_last()
	retGo := (float64)(retC)

	return retGo
}

// Unsupported : g_test_timer_start : no return generator

// Unsupported : g_test_trap_fork : no return generator

// Unsupported : g_test_trap_has_passed : no return generator

// Unsupported : g_test_trap_reached_timeout : no return generator

// Unsupported : g_uri_escape_string : unsupported parameter allow_utf8 : no type generator for gboolean, gboolean

// UriParseScheme is a wrapper around the C function g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_uri_parse_scheme(c_uri)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

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
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriUnescapeString is a wrapper around the C function g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_string(c_escaped_string, c_illegal_characters)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

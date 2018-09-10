// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_checksum_type_get_length : unsupported parameter checksum_type : type ChecksumType, GChecksumType

// Unsupported : g_compute_checksum_for_data : unsupported parameter checksum_type : type ChecksumType, GChecksumType

// Unsupported : g_compute_checksum_for_string : unsupported parameter checksum_type : type ChecksumType, GChecksumType

// Dpgettext is a wrapper around the C function g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgctxtid := C.CString(msgctxtid)
	defer C.free(unsafe.Pointer(c_msgctxtid))

	c_msgidoffset := (C.gsize)(msgidoffset)

	C.g_dpgettext(c_domain, c_msgctxtid, c_msgidoffset)
}

// Unsupported : g_format_size_for_display : unsupported parameter size : type gint64, goffset

// Unsupported : g_markup_collect_attributes : unsupported parameter error : type Error, GError**

// Unsupported : g_prefix_error : unsupported parameter err : type Error, GError**

// Unsupported : g_propagate_prefixed_error : unsupported parameter dest : type Error, GError**

// Strcmp0 is a wrapper around the C function g_strcmp0.
func Strcmp0(str1 string, str2 string) {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	C.g_strcmp0(c_str1, c_str2)
}

// Unsupported : g_test_add_data_func : unsupported parameter test_data : type gpointer, gconstpointer

// Unsupported : g_test_add_func : unsupported parameter test_func : type TestFunc, GTestFunc

// TestBug is a wrapper around the C function g_test_bug.
func TestBug(bugUriSnippet string) {
	c_bug_uri_snippet := C.CString(bugUriSnippet)
	defer C.free(unsafe.Pointer(c_bug_uri_snippet))

	C.g_test_bug(c_bug_uri_snippet)
}

// TestBugBase is a wrapper around the C function g_test_bug_base.
func TestBugBase(uriPattern string) {
	c_uri_pattern := C.CString(uriPattern)
	defer C.free(unsafe.Pointer(c_uri_pattern))

	C.g_test_bug_base(c_uri_pattern)
}

// Unsupported : g_test_create_case : unsupported parameter test_data : type gpointer, gconstpointer

// TestCreateSuite is a wrapper around the C function g_test_create_suite.
func TestCreateSuite(suiteName string) {
	c_suite_name := C.CString(suiteName)
	defer C.free(unsafe.Pointer(c_suite_name))

	C.g_test_create_suite(c_suite_name)
}

// TestGetRoot is a wrapper around the C function g_test_get_root.
func TestGetRoot() {
	C.g_test_get_root()
}

// Unsupported : g_test_init : unsupported parameter argc : type gint, int*

// Unsupported : g_test_maximized_result : unsupported parameter maximized_quantity : type gdouble, double

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter minimized_quantity : type gdouble, double

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : type DestroyNotify, GDestroyNotify

// TestQueueFree is a wrapper around the C function g_test_queue_free.
func TestQueueFree(gfreePointer uintptr) {
	c_gfree_pointer := (C.gpointer)(gfreePointer)

	C.g_test_queue_free(c_gfree_pointer)
}

// TestRandDouble is a wrapper around the C function g_test_rand_double.
func TestRandDouble() {
	C.g_test_rand_double()
}

// Unsupported : g_test_rand_double_range : unsupported parameter range_start : type gdouble, double

// TestRandInt is a wrapper around the C function g_test_rand_int.
func TestRandInt() {
	C.g_test_rand_int()
}

// TestRandIntRange is a wrapper around the C function g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	C.g_test_rand_int_range(c_begin, c_end)
}

// TestRun is a wrapper around the C function g_test_run.
func TestRun() {
	C.g_test_run()
}

// Unsupported : g_test_run_suite : unsupported parameter suite : type TestSuite, GTestSuite*

// TestTimerElapsed is a wrapper around the C function g_test_timer_elapsed.
func TestTimerElapsed() {
	C.g_test_timer_elapsed()
}

// TestTimerLast is a wrapper around the C function g_test_timer_last.
func TestTimerLast() {
	C.g_test_timer_last()
}

// TestTimerStart is a wrapper around the C function g_test_timer_start.
func TestTimerStart() {
	C.g_test_timer_start()
}

// Unsupported : g_test_trap_fork : unsupported parameter test_trap_flags : type TestTrapFlags, GTestTrapFlags

// TestTrapHasPassed is a wrapper around the C function g_test_trap_has_passed.
func TestTrapHasPassed() {
	C.g_test_trap_has_passed()
}

// TestTrapReachedTimeout is a wrapper around the C function g_test_trap_reached_timeout.
func TestTrapReachedTimeout() {
	C.g_test_trap_reached_timeout()
}

// Unsupported : g_uri_escape_string : unsupported parameter allow_utf8 : type gboolean, gboolean

// UriParseScheme is a wrapper around the C function g_uri_parse_scheme.
func UriParseScheme(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.g_uri_parse_scheme(c_uri)
}

// UriUnescapeSegment is a wrapper around the C function g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_escaped_string_end := C.CString(escapedStringEnd)
	defer C.free(unsafe.Pointer(c_escaped_string_end))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	C.g_uri_unescape_segment(c_escaped_string, c_escaped_string_end, c_illegal_characters)
}

// UriUnescapeString is a wrapper around the C function g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	C.g_uri_unescape_string(c_escaped_string, c_illegal_characters)
}

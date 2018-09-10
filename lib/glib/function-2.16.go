// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_checksum_type_get_length : unsupported parameter checksum_type : type ChecksumType, GChecksumType

// Unsupported function: g_compute_checksum_for_data : unsupported parameter checksum_type : type ChecksumType, GChecksumType

// Unsupported function: g_compute_checksum_for_string : unsupported parameter checksum_type : type ChecksumType, GChecksumType

// Dpgettext is a wrapper around the C function g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) {}

// Unsupported function: g_format_size_for_display : unsupported parameter size : type gint64, goffset

// Unsupported function: g_markup_collect_attributes : unsupported parameter error : type Error, GError**

// Unsupported function: g_prefix_error : unsupported parameter err : type Error, GError**

// Unsupported function: g_propagate_prefixed_error : unsupported parameter dest : type Error, GError**

// Strcmp0 is a wrapper around the C function g_strcmp0.
func Strcmp0(str1 string, str2 string) {}

// Unsupported function: g_test_add_data_func : unsupported parameter test_data : type gpointer, gconstpointer

// Unsupported function: g_test_add_func : unsupported parameter test_func : type TestFunc, GTestFunc

// TestBug is a wrapper around the C function g_test_bug.
func TestBug(bugUriSnippet string) {}

// TestBugBase is a wrapper around the C function g_test_bug_base.
func TestBugBase(uriPattern string) {}

// Unsupported function: g_test_create_case : unsupported parameter test_data : type gpointer, gconstpointer

// TestCreateSuite is a wrapper around the C function g_test_create_suite.
func TestCreateSuite(suiteName string) {}

// TestGetRoot is a wrapper around the C function g_test_get_root.
func TestGetRoot() {}

// Unsupported function: g_test_init : unsupported parameter argc : type gint, int*

// Unsupported function: g_test_maximized_result : unsupported parameter maximized_quantity : type gdouble, double

// Unsupported function: g_test_message : unsupported parameter ... : varargs

// Unsupported function: g_test_minimized_result : unsupported parameter minimized_quantity : type gdouble, double

// Unsupported function: g_test_queue_destroy : unsupported parameter destroy_func : type DestroyNotify, GDestroyNotify

// TestQueueFree is a wrapper around the C function g_test_queue_free.
func TestQueueFree(gfreePointer uintptr) {}

// TestRandDouble is a wrapper around the C function g_test_rand_double.
func TestRandDouble() {}

// Unsupported function: g_test_rand_double_range : unsupported parameter range_start : type gdouble, double

// TestRandInt is a wrapper around the C function g_test_rand_int.
func TestRandInt() {}

// TestRandIntRange is a wrapper around the C function g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) {}

// TestRun is a wrapper around the C function g_test_run.
func TestRun() {}

// Unsupported function: g_test_run_suite : unsupported parameter suite : type TestSuite, GTestSuite*

// TestTimerElapsed is a wrapper around the C function g_test_timer_elapsed.
func TestTimerElapsed() {}

// TestTimerLast is a wrapper around the C function g_test_timer_last.
func TestTimerLast() {}

// TestTimerStart is a wrapper around the C function g_test_timer_start.
func TestTimerStart() {}

// Unsupported function: g_test_trap_fork : unsupported parameter test_trap_flags : type TestTrapFlags, GTestTrapFlags

// TestTrapHasPassed is a wrapper around the C function g_test_trap_has_passed.
func TestTrapHasPassed() {}

// TestTrapReachedTimeout is a wrapper around the C function g_test_trap_reached_timeout.
func TestTrapReachedTimeout() {}

// Unsupported function: g_uri_escape_string : unsupported parameter allow_utf8 : type gboolean, gboolean

// UriParseScheme is a wrapper around the C function g_uri_parse_scheme.
func UriParseScheme(uri string) {}

// UriUnescapeSegment is a wrapper around the C function g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) {}

// UriUnescapeString is a wrapper around the C function g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) {}

// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_checksum_type_get_length : unsupported parameter checksum_type : no param type for ChecksumType, GChecksumType

// Unsupported : g_compute_checksum_for_data : unsupported parameter checksum_type : no param type for ChecksumType, GChecksumType

// Unsupported : g_compute_checksum_for_string : unsupported parameter checksum_type : no param type for ChecksumType, GChecksumType

// Unsupported : g_dpgettext : no return type

// Unsupported : g_format_size_for_display : unsupported parameter size : no param type for gint64, goffset

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in for string with indirection level of 2

// Unsupported : g_prefix_error : unsupported parameter err : no param type for Error, GError**

// Unsupported : g_propagate_prefixed_error : unsupported parameter dest : no param type for Error, GError**

// Unsupported : g_strcmp0 : no return type

// Unsupported : g_test_add_data_func : unsupported parameter test_data : no param type for gpointer, gconstpointer

// Unsupported : g_test_add_func : unsupported parameter test_func : no param type for TestFunc, GTestFunc

// Unsupported : g_test_bug : no return type

// Unsupported : g_test_bug_base : no return type

// Unsupported : g_test_create_case : unsupported parameter test_data : no param type for gpointer, gconstpointer

// Unsupported : g_test_create_suite : no return type

// Unsupported : g_test_get_root : no return type

// Unsupported : g_test_init : unsupported parameter argc : no param type for gint, int*

// Unsupported : g_test_maximized_result : unsupported parameter maximized_quantity : no param type for gdouble, double

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter minimized_quantity : no param type for gdouble, double

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no param type for DestroyNotify, GDestroyNotify

// Unsupported : g_test_queue_free : no return type

// Unsupported : g_test_rand_double : no return type

// Unsupported : g_test_rand_double_range : unsupported parameter range_start : no param type for gdouble, double

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

// Unsupported : g_test_run : no return type

// Unsupported : g_test_run_suite : unsupported parameter suite : no param type for TestSuite, GTestSuite*

// Unsupported : g_test_timer_elapsed : no return type

// Unsupported : g_test_timer_last : no return type

// Unsupported : g_test_timer_start : no return type

// Unsupported : g_test_trap_fork : unsupported parameter test_trap_flags : no param type for TestTrapFlags, GTestTrapFlags

// Unsupported : g_test_trap_has_passed : no return type

// Unsupported : g_test_trap_reached_timeout : no return type

// Unsupported : g_uri_escape_string : unsupported parameter allow_utf8 : no param type for gboolean, gboolean

// Unsupported : g_uri_parse_scheme : no return type

// Unsupported : g_uri_unescape_segment : no return type

// Unsupported : g_uri_unescape_string : no return type

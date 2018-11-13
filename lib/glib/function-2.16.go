// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Gets the length in bytes of digests of type @checksum_type
/*

C function

g_checksum_type_get_length
*/
func ChecksumTypeGetLength(checksumType ChecksumType) int64 {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_type_get_length(c_checksum_type)
	retGo := (int64)(retC)

	return retGo
}

// Computes the checksum for a binary @data of @length. This is a
// convenience wrapper for g_checksum_new(), g_checksum_get_string()
// and g_checksum_free().
//
// The hexadecimal string returned will be in lower case.
/*

C function

g_compute_checksum_for_data
*/
func ComputeChecksumForData(checksumType ChecksumType, data []uint8) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_data := &data[0]

	c_length := (C.gsize)(len(data))

	retC := C.g_compute_checksum_for_data(c_checksum_type, (*C.guchar)(unsafe.Pointer(c_data)), c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Computes the checksum of a string.
//
// The hexadecimal string returned will be in lower case.
/*

C function

g_compute_checksum_for_string
*/
func ComputeChecksumForString(checksumType ChecksumType, str string) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(len(str))

	retC := C.g_compute_checksum_for_string(c_checksum_type, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// This function is a variant of g_dgettext() which supports
// a disambiguating message context. GNU gettext uses the
// '\004' character to separate the message context and
// message id in @msgctxtid.
// If 0 is passed as @msgidoffset, this function will fall back to
// trying to use the deprecated convention of using "|" as a separation
// character.
//
// This uses g_dgettext() internally. See that functions for differences
// with dgettext() proper.
//
// Applications should normally not use this function directly,
// but use the C_() macro for translations with context.
/*

C function

g_dpgettext
*/
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

// Formats a size (for example the size of a file) into a human
// readable string. Sizes are rounded to the nearest size prefix
// (KB, MB, GB) and are displayed rounded to the nearest tenth.
// E.g. the file size 3292528 bytes will be converted into the
// string "3.1 MB".
//
// The prefix units base is 1024 (i.e. 1 KB is 1024 bytes).
//
// This string should be freed with g_free() when not needed any longer.
/*

C function

g_format_size_for_display
*/
func FormatSizeForDisplay(size uint64) string {
	c_size := (C.goffset)(size)

	retC := C.g_format_size_for_display(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// Unsupported : g_prefix_error : unsupported parameter ... : varargs

// Unsupported : g_propagate_prefixed_error : unsupported parameter ... : varargs

// Compares @str1 and @str2 like strcmp(). Handles %NULL
// gracefully by sorting it before non-%NULL strings.
// Comparing two %NULL pointers returns 0.
/*

C function

g_strcmp0
*/
func Strcmp0(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_strcmp0(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_test_add_data_func : unsupported parameter test_func : no type generator for TestDataFunc (GTestDataFunc) for param test_func

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc (GTestFunc) for param test_func

// This function adds a message to test reports that
// associates a bug URI with a test case.
// Bug URIs are constructed from a base URI set with g_test_bug_base()
// and @bug_uri_snippet.
/*

C function

g_test_bug
*/
func TestBug(bugUriSnippet string) {
	c_bug_uri_snippet := C.CString(bugUriSnippet)
	defer C.free(unsafe.Pointer(c_bug_uri_snippet))

	C.g_test_bug(c_bug_uri_snippet)

	return
}

// Specify the base URI for bug reports.
//
// The base URI is used to construct bug report messages for
// g_test_message() when g_test_bug() is called.
// Calling this function outside of a test case sets the
// default base URI for all test cases. Calling it from within
// a test case changes the base URI for the scope of the test
// case only.
// Bug URIs are constructed by appending a bug specific URI
// portion to @uri_pattern, or by replacing the special string
// '\%s' within @uri_pattern if that is present.
/*

C function

g_test_bug_base
*/
func TestBugBase(uriPattern string) {
	c_uri_pattern := C.CString(uriPattern)
	defer C.free(unsafe.Pointer(c_uri_pattern))

	C.g_test_bug_base(c_uri_pattern)

	return
}

// Unsupported : g_test_create_case : unsupported parameter data_setup : no type generator for TestFixtureFunc (GTestFixtureFunc) for param data_setup

// Create a new test suite with the name @suite_name.
/*

C function

g_test_create_suite
*/
func TestCreateSuite(suiteName string) *TestSuite {
	c_suite_name := C.CString(suiteName)
	defer C.free(unsafe.Pointer(c_suite_name))

	retC := C.g_test_create_suite(c_suite_name)
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get the toplevel test suite for the test path API.
/*

C function

g_test_get_root
*/
func TestGetRoot() *TestSuite {
	retC := C.g_test_get_root()
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_test_init : unsupported parameter ... : varargs

// Unsupported : g_test_maximized_result : unsupported parameter ... : varargs

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// Enqueue a pointer to be released with g_free() during the next
// teardown phase. This is equivalent to calling g_test_queue_destroy()
// with a destroy callback of g_free().
/*

C function

g_test_queue_free
*/
func TestQueueFree(gfreePointer uintptr) {
	c_gfree_pointer := (C.gpointer)(gfreePointer)

	C.g_test_queue_free(c_gfree_pointer)

	return
}

// Get a reproducible random floating point number,
// see g_test_rand_int() for details on test case random numbers.
/*

C function

g_test_rand_double
*/
func TestRandDouble() float64 {
	retC := C.g_test_rand_double()
	retGo := (float64)(retC)

	return retGo
}

// Get a reproducible random floating pointer number out of a specified range,
// see g_test_rand_int() for details on test case random numbers.
/*

C function

g_test_rand_double_range
*/
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) float64 {
	c_range_start := (C.double)(rangeStart)

	c_range_end := (C.double)(rangeEnd)

	retC := C.g_test_rand_double_range(c_range_start, c_range_end)
	retGo := (float64)(retC)

	return retGo
}

// Get a reproducible random integer number.
//
// The random numbers generated by the g_test_rand_*() family of functions
// change with every new test program start, unless the --seed option is
// given when starting test programs.
//
// For individual test cases however, the random number generator is
// reseeded, to avoid dependencies between tests and to make --seed
// effective for all test cases.
/*

C function

g_test_rand_int
*/
func TestRandInt() int32 {
	retC := C.g_test_rand_int()
	retGo := (int32)(retC)

	return retGo
}

// Get a reproducible random integer number out of a specified range,
// see g_test_rand_int() for details on test case random numbers.
/*

C function

g_test_rand_int_range
*/
func TestRandIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_test_rand_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// Runs all tests under the toplevel suite which can be retrieved
// with g_test_get_root(). Similar to g_test_run_suite(), the test
// cases to be run are filtered according to test path arguments
// (`-p testpath` and `-s testpath`) as parsed by g_test_init().
// g_test_run_suite() or g_test_run() may only be called once in a
// program.
//
// In general, the tests and sub-suites within each suite are run in
// the order in which they are defined. However, note that prior to
// GLib 2.36, there was a bug in the `g_test_add_*`
// functions which caused them to create multiple suites with the same
// name, meaning that if you created tests "/foo/simple",
// "/bar/simple", and "/foo/using-bar" in that order, they would get
// run in that order (since g_test_run() would run the first "/foo"
// suite, then the "/bar" suite, then the second "/foo" suite). As of
// 2.36, this bug is fixed, and adding the tests in that order would
// result in a running order of "/foo/simple", "/foo/using-bar",
// "/bar/simple". If this new ordering is sub-optimal (because it puts
// more-complicated tests before simpler ones, making it harder to
// figure out exactly what has failed), you can fix it by changing the
// test paths to group tests by suite in a way that will result in the
// desired running order. Eg, "/simple/foo", "/simple/bar",
// "/complex/foo-using-bar".
//
// However, you should never make the actual result of a test depend
// on the order that tests are run in. If you need to ensure that some
// particular code runs before or after a given test case, use
// g_test_add(), which lets you specify setup and teardown functions.
//
// If all tests are skipped, this function will return 0 if
// producing TAP output, or 77 (treated as "skip test" by Automake) otherwise.
/*

C function

g_test_run
*/
func TestRun() int32 {
	retC := C.g_test_run()
	retGo := (int32)(retC)

	return retGo
}

// Execute the tests within @suite and all nested #GTestSuites.
// The test suites to be executed are filtered according to
// test path arguments (`-p testpath` and `-s testpath`) as parsed by
// g_test_init(). See the g_test_run() documentation for more
// information on the order that tests are run in.
//
// g_test_run_suite() or g_test_run() may only be called once
// in a program.
/*

C function

g_test_run_suite
*/
func TestRunSuite(suite *TestSuite) int32 {
	c_suite := (*C.GTestSuite)(C.NULL)
	if suite != nil {
		c_suite = (*C.GTestSuite)(suite.ToC())
	}

	retC := C.g_test_run_suite(c_suite)
	retGo := (int32)(retC)

	return retGo
}

// Get the time since the last start of the timer with g_test_timer_start().
/*

C function

g_test_timer_elapsed
*/
func TestTimerElapsed() float64 {
	retC := C.g_test_timer_elapsed()
	retGo := (float64)(retC)

	return retGo
}

// Report the last result of g_test_timer_elapsed().
/*

C function

g_test_timer_last
*/
func TestTimerLast() float64 {
	retC := C.g_test_timer_last()
	retGo := (float64)(retC)

	return retGo
}

// Start a timing test. Call g_test_timer_elapsed() when the task is supposed
// to be done. Call this function again to restart the timer.
/*

C function

g_test_timer_start
*/
func TestTimerStart() {
	C.g_test_timer_start()

	return
}

// Fork the current test program to execute a test case that might
// not return or that might abort.
//
// If @usec_timeout is non-0, the forked test case is aborted and
// considered failing if its run time exceeds it.
//
// The forking behavior can be configured with the #GTestTrapFlags flags.
//
// In the following example, the test code forks, the forked child
// process produces some sample output and exits successfully.
// The forking parent process then asserts successful child program
// termination and validates child program outputs.
//
// |[<!-- language="C" -->
// static void
// test_fork_patterns (void)
// {
// if (g_test_trap_fork (0, G_TEST_TRAP_SILENCE_STDOUT | G_TEST_TRAP_SILENCE_STDERR))
// {
// g_print ("some stdout text: somagic17\n");
// g_printerr ("some stderr text: semagic43\n");
// exit (0); // successful test run
// }
// g_test_trap_assert_passed ();
// g_test_trap_assert_stdout ("*somagic17*");
// g_test_trap_assert_stderr ("*semagic43*");
// }
// ]|
/*

C function

g_test_trap_fork
*/
func TestTrapFork(usecTimeout uint64, testTrapFlags TestTrapFlags) bool {
	c_usec_timeout := (C.guint64)(usecTimeout)

	c_test_trap_flags := (C.GTestTrapFlags)(testTrapFlags)

	retC := C.g_test_trap_fork(c_usec_timeout, c_test_trap_flags)
	retGo := retC == C.TRUE

	return retGo
}

// Check the result of the last g_test_trap_subprocess() call.
/*

C function

g_test_trap_has_passed
*/
func TestTrapHasPassed() bool {
	retC := C.g_test_trap_has_passed()
	retGo := retC == C.TRUE

	return retGo
}

// Check the result of the last g_test_trap_subprocess() call.
/*

C function

g_test_trap_reached_timeout
*/
func TestTrapReachedTimeout() bool {
	retC := C.g_test_trap_reached_timeout()
	retGo := retC == C.TRUE

	return retGo
}

// Escapes a string for use in a URI.
//
// Normally all characters that are not "unreserved" (i.e. ASCII alphanumerical
// characters plus dash, dot, underscore and tilde) are escaped.
// But if you specify characters in @reserved_chars_allowed they are not
// escaped. This is useful for the "reserved" characters in the URI
// specification, since those are allowed unescaped in some portions of
// a URI.
/*

C function

g_uri_escape_string
*/
func UriEscapeString(unescaped string, reservedCharsAllowed string, allowUtf8 bool) string {
	c_unescaped := C.CString(unescaped)
	defer C.free(unsafe.Pointer(c_unescaped))

	c_reserved_chars_allowed := C.CString(reservedCharsAllowed)
	defer C.free(unsafe.Pointer(c_reserved_chars_allowed))

	c_allow_utf8 :=
		boolToGboolean(allowUtf8)

	retC := C.g_uri_escape_string(c_unescaped, c_reserved_chars_allowed, c_allow_utf8)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the scheme portion of a URI string. RFC 3986 decodes the scheme as:
// |[
// URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ]
// ]|
// Common schemes include "file", "http", "svn+ssh", etc.
/*

C function

g_uri_parse_scheme
*/
func UriParseScheme(uri string) string {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_uri_parse_scheme(c_uri)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unescapes a segment of an escaped string.
//
// If any of the characters in @illegal_characters or the character zero appears
// as an escaped character in @escaped_string then that is an error and %NULL
// will be returned. This is useful it you want to avoid for instance having a
// slash being expanded in an escaped path element, which might confuse pathname
// handling.
/*

C function

g_uri_unescape_segment
*/
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

// Unescapes a whole escaped string.
//
// If any of the characters in @illegal_characters or the character zero appears
// as an escaped character in @escaped_string then that is an error and %NULL
// will be returned. This is useful it you want to avoid for instance having a
// slash being expanded in an escaped path element, which might confuse pathname
// handling.
/*

C function

g_uri_unescape_string
*/
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

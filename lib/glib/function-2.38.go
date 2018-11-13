// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Returns whether a test has already failed. This will
// be the case when g_test_fail(), g_test_incomplete()
// or g_test_skip() have been called, but also if an
// assertion has failed.
//
// This can be useful to return early from a test if
// continuing after a failed assertion might be harmful.
//
// The return value of this function is only meaningful
// if it is called from inside a test function.
/*

C function : g_test_failed
*/
func TestFailed() bool {
	retC := C.g_test_failed()
	retGo := retC == C.TRUE

	return retGo
}

// Gets the pathname of the directory containing test files of the type
// specified by @file_type.
//
// This is approximately the same as calling g_test_build_filename("."),
// but you don't need to free the return value.
/*

C function : g_test_get_dir
*/
func TestGetDir(fileType TestFileType) string {
	c_file_type := (C.GTestFileType)(fileType)

	retC := C.g_test_get_dir(c_file_type)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// Indicates that a test failed because of some incomplete
// functionality. This function can be called multiple times
// from the same test.
//
// Calling this function will not stop the test from running, you
// need to return from the test function yourself. So you can
// produce additional diagnostic messages or even continue running
// the test.
//
// If not called from inside a test, this function does nothing.
/*

C function : g_test_incomplete
*/
func TestIncomplete(msg string) {
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))

	C.g_test_incomplete(c_msg)

	return
}

// Changes the behaviour of g_assert_cmpstr(), g_assert_cmpint(),
// g_assert_cmpuint(), g_assert_cmphex(), g_assert_cmpfloat(),
// g_assert_true(), g_assert_false(), g_assert_null(), g_assert_no_error(),
// g_assert_error(), g_test_assert_expected_messages() and the various
// g_test_trap_assert_*() macros to not abort to program, but instead
// call g_test_fail() and continue. (This also changes the behavior of
// g_test_fail() so that it will not cause the test program to abort
// after completing the failed test.)
//
// Note that the g_assert_not_reached() and g_assert() are not
// affected by this.
//
// This function can only be called after g_test_init().
/*

C function : g_test_set_nonfatal_assertions
*/
func TestSetNonfatalAssertions() {
	C.g_test_set_nonfatal_assertions()

	return
}

// Indicates that a test was skipped.
//
// Calling this function will not stop the test from running, you
// need to return from the test function yourself. So you can
// produce additional diagnostic messages or even continue running
// the test.
//
// If not called from inside a test, this function does nothing.
/*

C function : g_test_skip
*/
func TestSkip(msg string) {
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))

	C.g_test_skip(c_msg)

	return
}

// Returns %TRUE (after g_test_init() has been called) if the test
// program is running under g_test_trap_subprocess().
/*

C function : g_test_subprocess
*/
func TestSubprocess() bool {
	retC := C.g_test_subprocess()
	retGo := retC == C.TRUE

	return retGo
}

// Respawns the test program to run only @test_path in a subprocess.
// This can be used for a test case that might not return, or that
// might abort.
//
// If @test_path is %NULL then the same test is re-run in a subprocess.
// You can use g_test_subprocess() to determine whether the test is in
// a subprocess or not.
//
// @test_path can also be the name of the parent test, followed by
// "`/subprocess/`" and then a name for the specific subtest (or just
// ending with "`/subprocess`" if the test only has one child test);
// tests with names of this form will automatically be skipped in the
// parent process.
//
// If @usec_timeout is non-0, the test subprocess is aborted and
// considered failing if its run time exceeds it.
//
// The subprocess behavior can be configured with the
// #GTestSubprocessFlags flags.
//
// You can use methods such as g_test_trap_assert_passed(),
// g_test_trap_assert_failed(), and g_test_trap_assert_stderr() to
// check the results of the subprocess. (But note that
// g_test_trap_assert_stdout() and g_test_trap_assert_stderr()
// cannot be used if @test_flags specifies that the child should
// inherit the parent stdout/stderr.)
//
// If your `main ()` needs to behave differently in
// the subprocess, you can call g_test_subprocess() (after calling
// g_test_init()) to see whether you are in a subprocess.
//
// The following example tests that calling
// `my_object_new(1000000)` will abort with an error
// message.
//
// |[<!-- language="C" -->
// static void
// test_create_large_object (void)
// {
// if (g_test_subprocess ())
// {
// my_object_new (1000000);
// return;
// }
//
// Reruns this same test in a subprocess
// g_test_trap_subprocess (NULL, 0, 0);
// g_test_trap_assert_failed ();
// g_test_trap_assert_stderr ("*ERROR*too large*");
// }
//
// int
// main (int argc, char **argv)
// {
// g_test_init (&argc, &argv, NULL);
//
// g_test_add_func ("/myobject/create_large_object",
// test_create_large_object);
// return g_test_run ();
// }
// ]|
/*

C function : g_test_trap_subprocess
*/
func TestTrapSubprocess(testPath string, usecTimeout uint64, testFlags TestSubprocessFlags) {
	c_test_path := C.CString(testPath)
	defer C.free(unsafe.Pointer(c_test_path))

	c_usec_timeout := (C.guint64)(usecTimeout)

	c_test_flags := (C.GTestSubprocessFlags)(testFlags)

	C.g_test_trap_subprocess(c_test_path, c_usec_timeout, c_test_flags)

	return
}

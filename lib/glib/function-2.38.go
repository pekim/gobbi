// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Unsupported : g_test_failed : no return generator

// TestGetDir is a wrapper around the C function g_test_get_dir.
func TestGetDir(fileType TestFileType) string {
	c_file_type := (C.GTestFileType)(fileType)

	retC := C.g_test_get_dir(c_file_type)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// Unsupported : g_test_incomplete : no return generator

// Unsupported : g_test_set_nonfatal_assertions : no return generator

// Unsupported : g_test_skip : no return generator

// Unsupported : g_test_subprocess : no return generator

// Unsupported : g_test_trap_subprocess : unsupported parameter test_flags : no type generator for TestSubprocessFlags, GTestSubprocessFlags

// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_test_build_filename : unsupported parameter file_type : type TestFileType, GTestFileType

// TestFailed is a wrapper around the C function g_test_failed.
func TestFailed() {}

// Unsupported function: g_test_get_dir : unsupported parameter file_type : type TestFileType, GTestFileType

// Unsupported function: g_test_get_filename : unsupported parameter file_type : type TestFileType, GTestFileType

// TestIncomplete is a wrapper around the C function g_test_incomplete.
func TestIncomplete(msg string) {}

// TestSetNonfatalAssertions is a wrapper around the C function g_test_set_nonfatal_assertions.
func TestSetNonfatalAssertions() {}

// TestSkip is a wrapper around the C function g_test_skip.
func TestSkip(msg string) {}

// TestSubprocess is a wrapper around the C function g_test_subprocess.
func TestSubprocess() {}

// Unsupported function: g_test_trap_subprocess : unsupported parameter test_flags : type TestSubprocessFlags, GTestSubprocessFlags

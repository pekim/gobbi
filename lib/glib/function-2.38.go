// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"C"
	"unsafe"
)

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// TestFailed is a wrapper around the C function g_test_failed.
func TestFailed() bool {
	retC := C.g_test_failed()
	retGo := retC == C.TRUE

	return retGo
}

// TestGetDir is a wrapper around the C function g_test_get_dir.
func TestGetDir(fileType TestFileType) string {
	c_file_type := (C.GTestFileType)(fileType)

	retC := C.g_test_get_dir(c_file_type)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// TestIncomplete is a wrapper around the C function g_test_incomplete.
func TestIncomplete(msg string) {
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))

	C.g_test_incomplete(c_msg)

	return
}

// TestSetNonfatalAssertions is a wrapper around the C function g_test_set_nonfatal_assertions.
func TestSetNonfatalAssertions() {
	C.g_test_set_nonfatal_assertions()

	return
}

// TestSkip is a wrapper around the C function g_test_skip.
func TestSkip(msg string) {
	c_msg := C.CString(msg)
	defer C.free(unsafe.Pointer(c_msg))

	C.g_test_skip(c_msg)

	return
}

// TestSubprocess is a wrapper around the C function g_test_subprocess.
func TestSubprocess() bool {
	retC := C.g_test_subprocess()
	retGo := retC == C.TRUE

	return retGo
}

// TestTrapSubprocess is a wrapper around the C function g_test_trap_subprocess.
func TestTrapSubprocess(testPath string, usecTimeout uint64, testFlags TestSubprocessFlags) {
	c_test_path := C.CString(testPath)
	defer C.free(unsafe.Pointer(c_test_path))

	c_usec_timeout := (C.guint64)(usecTimeout)

	c_test_flags := (C.GTestSubprocessFlags)(testFlags)

	C.g_test_trap_subprocess(c_test_path, c_usec_timeout, c_test_flags)

	return
}

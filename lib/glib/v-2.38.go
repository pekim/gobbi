// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	static GVariant* _g_variant_new_printf(const gchar* format_string) {
		return g_variant_new_printf(format_string);
    }
*/
import "C"

const KEY_FILE_DESKTOP_KEY_ACTIONS string = C.G_KEY_FILE_DESKTOP_KEY_ACTIONS
const KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE string = C.G_KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE

type TestFileType C.GTestFileType

const (
	TEST_DIST  TestFileType = 0
	TEST_BUILT TestFileType = 1
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

// VariantNewPrintf is a wrapper around the C function g_variant_new_printf.
func VariantNewPrintf(formatString string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_new_printf(c_format_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewTakeString is a wrapper around the C function g_variant_new_take_string.
func VariantNewTakeString(string_ string) *Variant {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_new_take_string(c_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

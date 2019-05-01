// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import call "github.com/pekim/gobbi/lib/internal/call"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

const KEY_FILE_DESKTOP_KEY_ACTIONS string = "Actions"
const KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE string = "DBusActivatable"

type TestFileType int

const (
	TEST_DIST  TestFileType = 0
	TEST_BUILT TestFileType = 1
)

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Unsupported : g_test_failed : return type :

// Unsupported : g_test_get_dir : return type :

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// TestSetNonfatalAssertions is a wrapper around the C function g_test_set_nonfatal_assertions.
func TestSetNonfatalAssertions() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(3336, &data)
	return
}

// Unsupported : g_test_subprocess : return type :

// Unsupported : g_variant_new_printf : return type :

// Unsupported : g_variant_new_take_string : return type :

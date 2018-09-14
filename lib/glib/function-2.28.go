// This is a generated file - DO NOT EDIT
// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_get_environ : no return type

// Unsupported : g_get_locale_variants : no return type

// GetMonotonicTime is a wrapper around the C function g_get_monotonic_time.
func GetMonotonicTime() int64 {
	retC := C.g_get_monotonic_time()
	retGo :=
		(int64)(retC)

	return retGo
}

// GetRealTime is a wrapper around the C function g_get_real_time.
func GetRealTime() int64 {
	retC := C.g_get_real_time()
	retGo :=
		(int64)(retC)

	return retGo
}

// Unsupported : g_get_user_runtime_dir : no return type

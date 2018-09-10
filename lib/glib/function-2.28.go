// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// GetEnviron is a wrapper around the C function g_get_environ.
func GetEnviron() {
	C.g_get_environ()
}

// GetLocaleVariants is a wrapper around the C function g_get_locale_variants.
func GetLocaleVariants(locale string) {
	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	C.g_get_locale_variants()
}

// GetMonotonicTime is a wrapper around the C function g_get_monotonic_time.
func GetMonotonicTime() {
	C.g_get_monotonic_time()
}

// GetRealTime is a wrapper around the C function g_get_real_time.
func GetRealTime() {
	C.g_get_real_time()
}

// GetUserRuntimeDir is a wrapper around the C function g_get_user_runtime_dir.
func GetUserRuntimeDir() {
	C.g_get_user_runtime_dir()
}

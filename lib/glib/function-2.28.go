// This is a generated file - DO NOT EDIT
// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// GetEnviron is a wrapper around the C function g_get_environ.
func GetEnviron() []string {
	retC := C.g_get_environ()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetLocaleVariants is a wrapper around the C function g_get_locale_variants.
func GetLocaleVariants(locale string) []string {
	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	retC := C.g_get_locale_variants(c_locale)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetMonotonicTime is a wrapper around the C function g_get_monotonic_time.
func GetMonotonicTime() int64 {
	retC := C.g_get_monotonic_time()
	retGo := (int64)(retC)

	return retGo
}

// GetRealTime is a wrapper around the C function g_get_real_time.
func GetRealTime() int64 {
	retC := C.g_get_real_time()
	retGo := (int64)(retC)

	return retGo
}

// GetUserRuntimeDir is a wrapper around the C function g_get_user_runtime_dir.
func GetUserRuntimeDir() string {
	retC := C.g_get_user_runtime_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_test_add_data_func : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

// Unsupported : g_test_add_vtable : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

// Unsupported : g_test_create_case : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

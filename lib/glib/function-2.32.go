// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_byte_array_new_take : array return type :

// EnvironGetenv is a wrapper around the C function g_environ_getenv.
func EnvironGetenv(envp []string, variable string) string {
	c_envp_array := make([]*C.gchar, len(envp), len(envp))
	for i, item := range envp {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_envp_array[i] = c
	}
	c_envp_arrayPtr := &c_envp_array[0]
	c_envp := (**C.gchar)(unsafe.Pointer(c_envp_arrayPtr))

	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_environ_getenv(c_envp, c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// EnvironSetenv is a wrapper around the C function g_environ_setenv.
func EnvironSetenv(envp []string, variable string, value string, overwrite bool) []string {
	c_envp_array := make([]*C.gchar, len(envp), len(envp))
	for i, item := range envp {
		c := C.CString(item)
		c_envp_array[i] = c
	}
	c_envp_arrayPtr := &c_envp_array[0]
	c_envp := (**C.gchar)(unsafe.Pointer(c_envp_arrayPtr))

	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	retC := C.g_environ_setenv(c_envp, c_variable, c_value, c_overwrite)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// EnvironUnsetenv is a wrapper around the C function g_environ_unsetenv.
func EnvironUnsetenv(envp []string, variable string) []string {
	c_envp_array := make([]*C.gchar, len(envp), len(envp))
	for i, item := range envp {
		c := C.CString(item)
		c_envp_array[i] = c
	}
	c_envp_arrayPtr := &c_envp_array[0]
	c_envp := (**C.gchar)(unsafe.Pointer(c_envp_arrayPtr))

	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_environ_unsetenv(c_envp, c_variable)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// Unsupported : g_hash_table_add : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_contains : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

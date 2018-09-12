// +build glib_2.2 glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// AsciiStrtoull is a wrapper around the C function g_ascii_strtoull.
func AsciiStrtoull(nptr string, base uint32) uint64 {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	c_base := (C.guint)(base)

	retC := C.g_ascii_strtoull(c_nptr, &c_endptr, c_base)
	retGo :=
		(uint64)(retC)

	return retGo
}

// Unsupported : g_fprintf : unsupported parameter file : no param type for gpointer, FILE*

// GetApplicationName is a wrapper around the C function g_get_application_name.
func GetApplicationName() string {
	retC := C.g_get_application_name()
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_printf : unsupported parameter ... : varargs

// Unsupported : g_set_application_name : no return type

// Unsupported : g_sprintf : unsupported parameter ... : varargs

// Unsupported : g_str_has_prefix : no return type

// Unsupported : g_str_has_suffix : no return type

// Utf8Strreverse is a wrapper around the C function g_utf8_strreverse.
func Utf8Strreverse(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strreverse(c_str, c_len)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_vfprintf : unsupported parameter file : no param type for gpointer, FILE*

// Unsupported : g_vprintf : unsupported parameter args : no param type for va_list, va_list

// Unsupported : g_vsprintf : unsupported parameter args : no param type for va_list, va_list

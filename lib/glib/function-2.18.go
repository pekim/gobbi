// This is a generated file - DO NOT EDIT
// +build glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Dgettext is a wrapper around the C function g_dgettext.
func Dgettext(domain string, msgid string) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	retC := C.g_dgettext(c_domain, c_msgid)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Dngettext is a wrapper around the C function g_dngettext.
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgid_plural := C.CString(msgidPlural)
	defer C.free(unsafe.Pointer(c_msgid_plural))

	c_n := (C.gulong)(n)

	retC := C.g_dngettext(c_domain, c_msgid, c_msgid_plural, c_n)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Dpgettext2 is a wrapper around the C function g_dpgettext2.
func Dpgettext2(domain string, context string, msgid string) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_context := C.CString(context)
	defer C.free(unsafe.Pointer(c_context))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	retC := C.g_dpgettext2(c_domain, c_context, c_msgid)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_set_error_literal : unsupported parameter err : no type generator for Error, GError**

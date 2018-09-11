// +build glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Dgettext is a wrapper around the C function g_dgettext.
func Dgettext(domain string, msgid string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	C.g_dgettext(c_domain, c_msgid)
}

// Dngettext is a wrapper around the C function g_dngettext.
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgid_plural := C.CString(msgidPlural)
	defer C.free(unsafe.Pointer(c_msgid_plural))

	c_n := (C.gulong)(n)

	C.g_dngettext(c_domain, c_msgid, c_msgid_plural, c_n)
}

// Dpgettext2 is a wrapper around the C function g_dpgettext2.
func Dpgettext2(domain string, context string, msgid string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_context := C.CString(context)
	defer C.free(unsafe.Pointer(c_context))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	C.g_dpgettext2(c_domain, c_context, c_msgid)
}

// Unsupported : g_set_error_literal : unsupported parameter err : no param type
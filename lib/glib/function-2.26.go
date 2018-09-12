// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_date_time_compare : unsupported parameter dt1 : no param type for gpointer, gconstpointer

// Unsupported : g_date_time_equal : unsupported parameter dt1 : no param type for gpointer, gconstpointer

// Unsupported : g_date_time_hash : unsupported parameter datetime : no param type for gpointer, gconstpointer

// Dcgettext is a wrapper around the C function g_dcgettext.
func Dcgettext(domain string, msgid string, category int32) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_category := (C.gint)(category)

	retC := C.g_dcgettext(c_domain, c_msgid, c_category)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_source_set_name_by_id : no return type

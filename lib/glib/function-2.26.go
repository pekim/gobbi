// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// DateTimeCompare is a wrapper around the C function g_date_time_compare.
func DateTimeCompare(dt1 uintptr, dt2 uintptr) int32 {
	c_dt1 := (C.gconstpointer)(dt1)

	c_dt2 := (C.gconstpointer)(dt2)

	retC := C.g_date_time_compare(c_dt1, c_dt2)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_date_time_equal : no return type

// DateTimeHash is a wrapper around the C function g_date_time_hash.
func DateTimeHash(datetime uintptr) uint32 {
	c_datetime := (C.gconstpointer)(datetime)

	retC := C.g_date_time_hash(c_datetime)
	retGo :=
		(uint32)(retC)

	return retGo
}

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
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_source_set_name_by_id : no return type

// This is a generated file - DO NOT EDIT
// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// A comparison function for #GDateTimes that is suitable
// as a #GCompareFunc. Both #GDateTimes must be non-%NULL.
/*

C function : g_date_time_compare
*/
func DateTimeCompare(dt1 uintptr, dt2 uintptr) int32 {
	c_dt1 := (C.gconstpointer)(dt1)

	c_dt2 := (C.gconstpointer)(dt2)

	retC := C.g_date_time_compare(c_dt1, c_dt2)
	retGo := (int32)(retC)

	return retGo
}

// Checks to see if @dt1 and @dt2 are equal.
//
// Equal here means that they represent the same moment after converting
// them to the same time zone.
/*

C function : g_date_time_equal
*/
func DateTimeEqual(dt1 uintptr, dt2 uintptr) bool {
	c_dt1 := (C.gconstpointer)(dt1)

	c_dt2 := (C.gconstpointer)(dt2)

	retC := C.g_date_time_equal(c_dt1, c_dt2)
	retGo := retC == C.TRUE

	return retGo
}

// Hashes @datetime into a #guint, suitable for use within #GHashTable.
/*

C function : g_date_time_hash
*/
func DateTimeHash(datetime uintptr) uint32 {
	c_datetime := (C.gconstpointer)(datetime)

	retC := C.g_date_time_hash(c_datetime)
	retGo := (uint32)(retC)

	return retGo
}

// This is a variant of g_dgettext() that allows specifying a locale
// category instead of always using `LC_MESSAGES`. See g_dgettext() for
// more information about how this functions differs from calling
// dcgettext() directly.
/*

C function : g_dcgettext
*/
func Dcgettext(domain string, msgid string, category int32) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_category := (C.gint)(category)

	retC := C.g_dcgettext(c_domain, c_msgid, c_category)
	retGo := C.GoString(retC)

	return retGo
}

// Sets the name of a source using its ID.
//
// This is a convenience utility to set source names from the return
// value of g_idle_add(), g_timeout_add(), etc.
//
// It is a programmer error to attempt to set the name of a non-existent
// source.
//
// More specifically: source IDs can be reissued after a source has been
// destroyed and therefore it is never valid to use this function with a
// source ID which may have already been removed.  An example is when
// scheduling an idle to run in another thread with g_idle_add(): the
// idle may already have run and been removed by the time this function
// is called on its (now invalid) source ID.  This source ID may have
// been reissued, leading to the operation being performed against the
// wrong source.
/*

C function : g_source_set_name_by_id
*/
func SourceSetNameById(tag uint32, name string) {
	c_tag := (C.guint)(tag)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_source_set_name_by_id(c_tag, c_name)

	return
}

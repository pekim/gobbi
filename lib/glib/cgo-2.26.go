// This is a generated file - DO NOT EDIT
// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_date_time_compare : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1

// Unsupported : g_date_time_equal : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1

// Unsupported : g_date_time_hash : unsupported parameter datetime : no type generator for gpointer (gconstpointer) for param datetime

// DateTime is a wrapper around the C record GDateTime.
type DateTime struct {
	native *C.GDateTime
}

func DateTimeNewFromC(u unsafe.Pointer) *DateTime {
	c := (*C.GDateTime)(u)
	if c == nil {
		return nil
	}

	g := &DateTime{native: c}

	return g
}

func (recv *DateTime) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TimeZone is a wrapper around the C record GTimeZone.
type TimeZone struct {
	native *C.GTimeZone
}

func TimeZoneNewFromC(u unsafe.Pointer) *TimeZone {
	c := (*C.GTimeZone)(u)
	if c == nil {
		return nil
	}

	g := &TimeZone{native: c}

	return g
}

func (recv *TimeZone) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

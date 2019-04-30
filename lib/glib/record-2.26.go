// This is a generated file - DO NOT EDIT
// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// DateTime is a wrapper around the C record GDateTime.
type DateTime struct {
	native unsafe.Pointer
}

func DateTimeNewFromC(u unsafe.Pointer) *DateTime {
	if u == nil {
		return nil
	}

	g := &DateTime{native: u}

	return g
}

func (recv *DateTime) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_date_time_new : return type :

// Unsupported : g_date_time_new_from_timeval_local : return type :

// Unsupported : g_date_time_new_from_timeval_utc : return type :

// Unsupported : g_date_time_new_from_unix_local : return type :

// Unsupported : g_date_time_new_from_unix_utc : return type :

// Unsupported : g_date_time_new_local : return type :

// Unsupported : g_date_time_new_now : return type :

// Unsupported : g_date_time_new_now_local : return type :

// Unsupported : g_date_time_new_now_utc : return type :

// Unsupported : g_date_time_new_utc : return type :

// TimeZone is a wrapper around the C record GTimeZone.
type TimeZone struct {
	native unsafe.Pointer
}

func TimeZoneNewFromC(u unsafe.Pointer) *TimeZone {
	if u == nil {
		return nil
	}

	g := &TimeZone{native: u}

	return g
}

func (recv *TimeZone) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_time_zone_new : return type :

// Unsupported : g_time_zone_new_local : return type :

// Unsupported : g_time_zone_new_utc : return type :

// Unsupported : g_variant_new_bytestring : return type :

// Unsupported : g_variant_new_bytestring_array : return type :

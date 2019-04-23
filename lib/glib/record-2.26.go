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

// Equals compares this DateTime with another DateTime, and returns true if they represent the same GObject.
func (recv *DateTime) Equals(other *DateTime) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_date_time_new

// Blacklisted : g_date_time_new_from_timeval_local

// Blacklisted : g_date_time_new_from_timeval_utc

// Blacklisted : g_date_time_new_from_unix_local

// Blacklisted : g_date_time_new_from_unix_utc

// Blacklisted : g_date_time_new_local

// Blacklisted : g_date_time_new_now

// Blacklisted : g_date_time_new_now_local

// Blacklisted : g_date_time_new_now_utc

// Blacklisted : g_date_time_new_utc

// g_date_time_compare : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1
// g_date_time_equal : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1
// g_date_time_hash : unsupported parameter datetime : no type generator for gpointer (gconstpointer) for param datetime
// Blacklisted : g_date_time_add

// Blacklisted : g_date_time_add_days

// Blacklisted : g_date_time_add_full

// Blacklisted : g_date_time_add_hours

// Blacklisted : g_date_time_add_minutes

// Blacklisted : g_date_time_add_months

// Blacklisted : g_date_time_add_seconds

// Blacklisted : g_date_time_add_weeks

// Blacklisted : g_date_time_add_years

// Blacklisted : g_date_time_difference

// Blacklisted : g_date_time_format

// Blacklisted : g_date_time_get_day_of_month

// Blacklisted : g_date_time_get_day_of_week

// GetDayOfYear is a wrapper around the C function g_date_time_get_day_of_year.
func (recv *DateTime) GetDayOfYear() int32 {
	retC := C.g_date_time_get_day_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_date_time_get_hour

// Blacklisted : g_date_time_get_microsecond

// Blacklisted : g_date_time_get_minute

// GetMonth is a wrapper around the C function g_date_time_get_month.
func (recv *DateTime) GetMonth() int32 {
	retC := C.g_date_time_get_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_date_time_get_second

// Blacklisted : g_date_time_get_seconds

// Blacklisted : g_date_time_get_timezone_abbreviation

// Blacklisted : g_date_time_get_utc_offset

// Blacklisted : g_date_time_get_week_numbering_year

// Blacklisted : g_date_time_get_week_of_year

// GetYear is a wrapper around the C function g_date_time_get_year.
func (recv *DateTime) GetYear() int32 {
	retC := C.g_date_time_get_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_date_time_get_ymd

// Blacklisted : g_date_time_is_daylight_savings

// Blacklisted : g_date_time_ref

// Blacklisted : g_date_time_to_local

// Blacklisted : g_date_time_to_timeval

// Blacklisted : g_date_time_to_timezone

// Blacklisted : g_date_time_to_unix

// Blacklisted : g_date_time_to_utc

// Blacklisted : g_date_time_unref

// Blacklisted : g_key_file_get_int64

// Blacklisted : g_key_file_get_uint64

// Blacklisted : g_key_file_set_int64

// Blacklisted : g_key_file_set_uint64

// Blacklisted : g_source_set_name_by_id

// Blacklisted : g_source_get_name

// Blacklisted : g_source_set_name

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

// Equals compares this TimeZone with another TimeZone, and returns true if they represent the same GObject.
func (recv *TimeZone) Equals(other *TimeZone) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_time_zone_new

// Blacklisted : g_time_zone_new_local

// Blacklisted : g_time_zone_new_utc

// Blacklisted : g_time_zone_adjust_time

// Blacklisted : g_time_zone_find_interval

// Blacklisted : g_time_zone_get_abbreviation

// Blacklisted : g_time_zone_get_offset

// Blacklisted : g_time_zone_is_dst

// Blacklisted : g_time_zone_ref

// Blacklisted : g_time_zone_unref

// Blacklisted : g_variant_new_bytestring

// Blacklisted : g_variant_new_bytestring_array

// Blacklisted : g_variant_compare

// Unsupported : g_variant_dup_bytestring : array return type :

// Blacklisted : g_variant_dup_bytestring_array

// Unsupported : g_variant_get_bytestring : array return type :

// Blacklisted : g_variant_get_bytestring_array

// Blacklisted : g_variant_is_floating

// Blacklisted : g_variant_builder_add_parsed

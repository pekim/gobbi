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

// DateTimeNew is a wrapper around the C function g_date_time_new.
func DateTimeNew(tz *TimeZone, year int32, month int32, day int32, hour int32, minute int32, seconds float64) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	c_year := (C.gint)(year)

	c_month := (C.gint)(month)

	c_day := (C.gint)(day)

	c_hour := (C.gint)(hour)

	c_minute := (C.gint)(minute)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_new(c_tz, c_year, c_month, c_day, c_hour, c_minute, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromTimevalLocal is a wrapper around the C function g_date_time_new_from_timeval_local.
func DateTimeNewFromTimevalLocal(tv *TimeVal) *DateTime {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_new_from_timeval_local(c_tv)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromTimevalUtc is a wrapper around the C function g_date_time_new_from_timeval_utc.
func DateTimeNewFromTimevalUtc(tv *TimeVal) *DateTime {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_new_from_timeval_utc(c_tv)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromUnixLocal is a wrapper around the C function g_date_time_new_from_unix_local.
func DateTimeNewFromUnixLocal(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_local(c_t)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromUnixUtc is a wrapper around the C function g_date_time_new_from_unix_utc.
func DateTimeNewFromUnixUtc(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_utc(c_t)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewLocal is a wrapper around the C function g_date_time_new_local.
func DateTimeNewLocal(year int32, month int32, day int32, hour int32, minute int32, seconds float64) *DateTime {
	c_year := (C.gint)(year)

	c_month := (C.gint)(month)

	c_day := (C.gint)(day)

	c_hour := (C.gint)(hour)

	c_minute := (C.gint)(minute)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_new_local(c_year, c_month, c_day, c_hour, c_minute, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewNow is a wrapper around the C function g_date_time_new_now.
func DateTimeNewNow(tz *TimeZone) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	retC := C.g_date_time_new_now(c_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewNowLocal is a wrapper around the C function g_date_time_new_now_local.
func DateTimeNewNowLocal() *DateTime {
	retC := C.g_date_time_new_now_local()
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewNowUtc is a wrapper around the C function g_date_time_new_now_utc.
func DateTimeNewNowUtc() *DateTime {
	retC := C.g_date_time_new_now_utc()
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewUtc is a wrapper around the C function g_date_time_new_utc.
func DateTimeNewUtc(year int32, month int32, day int32, hour int32, minute int32, seconds float64) *DateTime {
	c_year := (C.gint)(year)

	c_month := (C.gint)(month)

	c_day := (C.gint)(day)

	c_hour := (C.gint)(hour)

	c_minute := (C.gint)(minute)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_new_utc(c_year, c_month, c_day, c_hour, c_minute, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
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

// TimeZoneNew is a wrapper around the C function g_time_zone_new.
func TimeZoneNew(identifier string) *TimeZone {
	c_identifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(c_identifier))

	retC := C.g_time_zone_new(c_identifier)
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TimeZoneNewLocal is a wrapper around the C function g_time_zone_new_local.
func TimeZoneNewLocal() *TimeZone {
	retC := C.g_time_zone_new_local()
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TimeZoneNewUtc is a wrapper around the C function g_time_zone_new_utc.
func TimeZoneNewUtc() *TimeZone {
	retC := C.g_time_zone_new_utc()
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewBytestring is a wrapper around the C function g_variant_new_bytestring.
func VariantNewBytestring(string_ []uint8) *Variant {
	c_string_array := make([]C.guint8, len(string_)+1, len(string_)+1)
	for i, item := range string_ {
		c := (C.guint8)(item)
		c_string_array[i] = c
	}
	c_string_array[len(string_)] = 0
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	retC := C.g_variant_new_bytestring(c_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewBytestringArray is a wrapper around the C function g_variant_new_bytestring_array.
func VariantNewBytestringArray(strv []string) *Variant {
	c_strv_array := make([]*C.gchar, len(strv)+1, len(strv)+1)
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_array[len(strv)] = nil
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.gchar)(unsafe.Pointer(c_strv_arrayPtr))

	c_length := (C.gssize)(len(strv))

	retC := C.g_variant_new_bytestring_array(c_strv, c_length)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

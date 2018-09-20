// This is a generated file - DO NOT EDIT
// +build glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// DateTime is a wrapper around the C record GDateTime.
type DateTime struct {
	native *C.GDateTime
}

func dateTimeNewFromC(c *C.GDateTime) *DateTime {
	if c == nil {
		return nil
	}

	g := &DateTime{native: c}

	return g
}

// Unsupported : g_date_time_new : unsupported parameter tz : record param - coming soon

// Unsupported : g_date_time_new_from_iso8601 : unsupported parameter default_tz : record param - coming soon

// Unsupported : g_date_time_new_from_timeval_local : unsupported parameter tv : record param - coming soon

// Unsupported : g_date_time_new_from_timeval_utc : unsupported parameter tv : record param - coming soon

// DateTimeNewFromUnixLocal is a wrapper around the C function g_date_time_new_from_unix_local.
func DateTimeNewFromUnixLocal(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_local(c_t)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// DateTimeNewFromUnixUtc is a wrapper around the C function g_date_time_new_from_unix_utc.
func DateTimeNewFromUnixUtc(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_utc(c_t)
	retGo := dateTimeNewFromC(retC)

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
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// Unsupported : g_date_time_new_now : unsupported parameter tz : record param - coming soon

// DateTimeNewNowLocal is a wrapper around the C function g_date_time_new_now_local.
func DateTimeNewNowLocal() *DateTime {
	retC := C.g_date_time_new_now_local()
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// DateTimeNewNowUtc is a wrapper around the C function g_date_time_new_now_utc.
func DateTimeNewNowUtc() *DateTime {
	retC := C.g_date_time_new_now_utc()
	retGo := dateTimeNewFromC(retC)

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
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// Add is a wrapper around the C function g_date_time_add.
func Add(timespan TimeSpan) *DateTime {
	c_timespan := (C.GTimeSpan)(timespan)

	retC := C.g_date_time_add(c_timespan)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddDays is a wrapper around the C function g_date_time_add_days.
func AddDays(days int32) *DateTime {
	c_days := (C.gint)(days)

	retC := C.g_date_time_add_days(c_days)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddFull is a wrapper around the C function g_date_time_add_full.
func AddFull(years int32, months int32, days int32, hours int32, minutes int32, seconds float64) *DateTime {
	c_years := (C.gint)(years)

	c_months := (C.gint)(months)

	c_days := (C.gint)(days)

	c_hours := (C.gint)(hours)

	c_minutes := (C.gint)(minutes)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_full(c_years, c_months, c_days, c_hours, c_minutes, c_seconds)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddHours is a wrapper around the C function g_date_time_add_hours.
func AddHours(hours int32) *DateTime {
	c_hours := (C.gint)(hours)

	retC := C.g_date_time_add_hours(c_hours)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddMinutes is a wrapper around the C function g_date_time_add_minutes.
func AddMinutes(minutes int32) *DateTime {
	c_minutes := (C.gint)(minutes)

	retC := C.g_date_time_add_minutes(c_minutes)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddMonths is a wrapper around the C function g_date_time_add_months.
func AddMonths(months int32) *DateTime {
	c_months := (C.gint)(months)

	retC := C.g_date_time_add_months(c_months)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddSeconds is a wrapper around the C function g_date_time_add_seconds.
func AddSeconds(seconds float64) *DateTime {
	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_seconds(c_seconds)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddWeeks is a wrapper around the C function g_date_time_add_weeks.
func AddWeeks(weeks int32) *DateTime {
	c_weeks := (C.gint)(weeks)

	retC := C.g_date_time_add_weeks(c_weeks)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// AddYears is a wrapper around the C function g_date_time_add_years.
func AddYears(years int32) *DateTime {
	c_years := (C.gint)(years)

	retC := C.g_date_time_add_years(c_years)
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// Unsupported : g_date_time_difference : unsupported parameter begin : record param - coming soon

// Format is a wrapper around the C function g_date_time_format.
func Format(format string) string {
	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	retC := C.g_date_time_format(c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDayOfMonth is a wrapper around the C function g_date_time_get_day_of_month.
func GetDayOfMonth() int32 {
	retC := C.g_date_time_get_day_of_month()
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfWeek is a wrapper around the C function g_date_time_get_day_of_week.
func GetDayOfWeek() int32 {
	retC := C.g_date_time_get_day_of_week()
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfYear is a wrapper around the C function g_date_time_get_day_of_year.
func GetDayOfYear() int32 {
	retC := C.g_date_time_get_day_of_year()
	retGo := (int32)(retC)

	return retGo
}

// GetHour is a wrapper around the C function g_date_time_get_hour.
func GetHour() int32 {
	retC := C.g_date_time_get_hour()
	retGo := (int32)(retC)

	return retGo
}

// GetMicrosecond is a wrapper around the C function g_date_time_get_microsecond.
func GetMicrosecond() int32 {
	retC := C.g_date_time_get_microsecond()
	retGo := (int32)(retC)

	return retGo
}

// GetMinute is a wrapper around the C function g_date_time_get_minute.
func GetMinute() int32 {
	retC := C.g_date_time_get_minute()
	retGo := (int32)(retC)

	return retGo
}

// GetMonth is a wrapper around the C function g_date_time_get_month.
func GetMonth() int32 {
	retC := C.g_date_time_get_month()
	retGo := (int32)(retC)

	return retGo
}

// GetSecond is a wrapper around the C function g_date_time_get_second.
func GetSecond() int32 {
	retC := C.g_date_time_get_second()
	retGo := (int32)(retC)

	return retGo
}

// GetSeconds is a wrapper around the C function g_date_time_get_seconds.
func GetSeconds() float64 {
	retC := C.g_date_time_get_seconds()
	retGo := (float64)(retC)

	return retGo
}

// GetTimezoneAbbreviation is a wrapper around the C function g_date_time_get_timezone_abbreviation.
func GetTimezoneAbbreviation() string {
	retC := C.g_date_time_get_timezone_abbreviation()
	retGo := C.GoString(retC)

	return retGo
}

// GetUtcOffset is a wrapper around the C function g_date_time_get_utc_offset.
func GetUtcOffset() TimeSpan {
	retC := C.g_date_time_get_utc_offset()
	retGo := (TimeSpan)(retC)

	return retGo
}

// GetWeekNumberingYear is a wrapper around the C function g_date_time_get_week_numbering_year.
func GetWeekNumberingYear() int32 {
	retC := C.g_date_time_get_week_numbering_year()
	retGo := (int32)(retC)

	return retGo
}

// GetWeekOfYear is a wrapper around the C function g_date_time_get_week_of_year.
func GetWeekOfYear() int32 {
	retC := C.g_date_time_get_week_of_year()
	retGo := (int32)(retC)

	return retGo
}

// GetYear is a wrapper around the C function g_date_time_get_year.
func GetYear() int32 {
	retC := C.g_date_time_get_year()
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_date_time_get_ymd : unsupported parameter year : no type generator for gint, gint*

// Unsupported : g_date_time_is_daylight_savings : no return generator

// Ref is a wrapper around the C function g_date_time_ref.
func Ref() *DateTime {
	retC := C.g_date_time_ref()
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// ToLocal is a wrapper around the C function g_date_time_to_local.
func ToLocal() *DateTime {
	retC := C.g_date_time_to_local()
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// Unsupported : g_date_time_to_timeval : unsupported parameter tv : record param - coming soon

// Unsupported : g_date_time_to_timezone : unsupported parameter tz : record param - coming soon

// ToUnix is a wrapper around the C function g_date_time_to_unix.
func ToUnix() int64 {
	retC := C.g_date_time_to_unix()
	retGo := (int64)(retC)

	return retGo
}

// ToUtc is a wrapper around the C function g_date_time_to_utc.
func ToUtc() *DateTime {
	retC := C.g_date_time_to_utc()
	retGo := dateTimeNewFromC(retC)

	return retGo
}

// Unsupported : g_date_time_unref : no return generator

// TimeZone is a wrapper around the C record GTimeZone.
type TimeZone struct {
	native *C.GTimeZone
}

func timeZoneNewFromC(c *C.GTimeZone) *TimeZone {
	if c == nil {
		return nil
	}

	g := &TimeZone{native: c}

	return g
}

// TimeZoneNew is a wrapper around the C function g_time_zone_new.
func TimeZoneNew(identifier string) *TimeZone {
	c_identifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(c_identifier))

	retC := C.g_time_zone_new(c_identifier)
	retGo := timeZoneNewFromC(retC)

	return retGo
}

// TimeZoneNewLocal is a wrapper around the C function g_time_zone_new_local.
func TimeZoneNewLocal() *TimeZone {
	retC := C.g_time_zone_new_local()
	retGo := timeZoneNewFromC(retC)

	return retGo
}

// TimeZoneNewUtc is a wrapper around the C function g_time_zone_new_utc.
func TimeZoneNewUtc() *TimeZone {
	retC := C.g_time_zone_new_utc()
	retGo := timeZoneNewFromC(retC)

	return retGo
}

// Unsupported : g_time_zone_adjust_time : unsupported parameter time_ : no type generator for gint64, gint64*

// FindInterval is a wrapper around the C function g_time_zone_find_interval.
func FindInterval(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_find_interval(c_type, c_time_)
	retGo := (int32)(retC)

	return retGo
}

// GetAbbreviation is a wrapper around the C function g_time_zone_get_abbreviation.
func GetAbbreviation(interval int32) string {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_abbreviation(c_interval)
	retGo := C.GoString(retC)

	return retGo
}

// GetOffset is a wrapper around the C function g_time_zone_get_offset.
func GetOffset(interval int32) int32 {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_offset(c_interval)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_time_zone_is_dst : no return generator

// Ref is a wrapper around the C function g_time_zone_ref.
func Ref() *TimeZone {
	retC := C.g_time_zone_ref()
	retGo := timeZoneNewFromC(retC)

	return retGo
}

// Unsupported : g_time_zone_unref : no return generator

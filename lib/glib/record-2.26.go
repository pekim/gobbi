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

// Creates a new #GDateTime corresponding to the given date and time in
// the time zone @tz.
//
// The @year must be between 1 and 9999, @month between 1 and 12 and @day
// between 1 and 28, 29, 30 or 31 depending on the month and the year.
//
// @hour must be between 0 and 23 and @minute must be between 0 and 59.
//
// @seconds must be at least 0.0 and must be strictly less than 60.0.
// It will be rounded down to the nearest microsecond.
//
// If the given time is not representable in the given time zone (for
// example, 02:30 on March 14th 2010 in Toronto, due to daylight savings
// time) then the time will be rounded up to the nearest existing time
// (in this case, 03:00).  If this matters to you then you should verify
// the return value for containing the same as the numbers you gave.
//
// In the case that the given time is ambiguous in the given time zone
// (for example, 01:30 on November 7th 2010 in Toronto, due to daylight
// savings time) then the time falling within standard (ie:
// non-daylight) time is taken.
//
// It not considered a programmer error for the values to this function
// to be out of range, but in the case that they are, the function will
// return %NULL.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_new
*/
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

// Creates a #GDateTime corresponding to the given #GTimeVal @tv in the
// local time zone.
//
// The time contained in a #GTimeVal is always stored in the form of
// seconds elapsed since 1970-01-01 00:00:00 UTC, regardless of the
// local time offset.
//
// This call can fail (returning %NULL) if @tv represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_new_from_timeval_local
*/
func DateTimeNewFromTimevalLocal(tv *TimeVal) *DateTime {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_new_from_timeval_local(c_tv)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GDateTime corresponding to the given #GTimeVal @tv in UTC.
//
// The time contained in a #GTimeVal is always stored in the form of
// seconds elapsed since 1970-01-01 00:00:00 UTC.
//
// This call can fail (returning %NULL) if @tv represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_new_from_timeval_utc
*/
func DateTimeNewFromTimevalUtc(tv *TimeVal) *DateTime {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_new_from_timeval_utc(c_tv)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GDateTime corresponding to the given Unix time @t in the
// local time zone.
//
// Unix time is the number of seconds that have elapsed since 1970-01-01
// 00:00:00 UTC, regardless of the local time offset.
//
// This call can fail (returning %NULL) if @t represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_new_from_unix_local
*/
func DateTimeNewFromUnixLocal(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_local(c_t)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GDateTime corresponding to the given Unix time @t in UTC.
//
// Unix time is the number of seconds that have elapsed since 1970-01-01
// 00:00:00 UTC.
//
// This call can fail (returning %NULL) if @t represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_new_from_unix_utc
*/
func DateTimeNewFromUnixUtc(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_utc(c_t)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GDateTime corresponding to the given date and time in
// the local time zone.
//
// This call is equivalent to calling g_date_time_new() with the time
// zone returned by g_time_zone_new_local().
/*

C function : g_date_time_new_local
*/
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

// Creates a #GDateTime corresponding to this exact instant in the given
// time zone @tz.  The time is as accurate as the system allows, to a
// maximum accuracy of 1 microsecond.
//
// This function will always succeed unless the system clock is set to
// truly insane values (or unless GLib is still being used after the
// year 9999).
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_new_now
*/
func DateTimeNewNow(tz *TimeZone) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	retC := C.g_date_time_new_now(c_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GDateTime corresponding to this exact instant in the local
// time zone.
//
// This is equivalent to calling g_date_time_new_now() with the time
// zone returned by g_time_zone_new_local().
/*

C function : g_date_time_new_now_local
*/
func DateTimeNewNowLocal() *DateTime {
	retC := C.g_date_time_new_now_local()
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GDateTime corresponding to this exact instant in UTC.
//
// This is equivalent to calling g_date_time_new_now() with the time
// zone returned by g_time_zone_new_utc().
/*

C function : g_date_time_new_now_utc
*/
func DateTimeNewNowUtc() *DateTime {
	retC := C.g_date_time_new_now_utc()
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GDateTime corresponding to the given date and time in
// UTC.
//
// This call is equivalent to calling g_date_time_new() with the time
// zone returned by g_time_zone_new_utc().
/*

C function : g_date_time_new_utc
*/
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

// Creates a copy of @datetime and adds the specified timespan to the copy.
/*

C function : g_date_time_add
*/
func (recv *DateTime) Add(timespan TimeSpan) *DateTime {
	c_timespan := (C.GTimeSpan)(timespan)

	retC := C.g_date_time_add((*C.GDateTime)(recv.native), c_timespan)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @datetime and adds the specified number of days to the
// copy. Add negative values to subtract days.
/*

C function : g_date_time_add_days
*/
func (recv *DateTime) AddDays(days int32) *DateTime {
	c_days := (C.gint)(days)

	retC := C.g_date_time_add_days((*C.GDateTime)(recv.native), c_days)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GDateTime adding the specified values to the current date and
// time in @datetime. Add negative values to subtract.
/*

C function : g_date_time_add_full
*/
func (recv *DateTime) AddFull(years int32, months int32, days int32, hours int32, minutes int32, seconds float64) *DateTime {
	c_years := (C.gint)(years)

	c_months := (C.gint)(months)

	c_days := (C.gint)(days)

	c_hours := (C.gint)(hours)

	c_minutes := (C.gint)(minutes)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_full((*C.GDateTime)(recv.native), c_years, c_months, c_days, c_hours, c_minutes, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @datetime and adds the specified number of hours.
// Add negative values to subtract hours.
/*

C function : g_date_time_add_hours
*/
func (recv *DateTime) AddHours(hours int32) *DateTime {
	c_hours := (C.gint)(hours)

	retC := C.g_date_time_add_hours((*C.GDateTime)(recv.native), c_hours)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @datetime adding the specified number of minutes.
// Add negative values to subtract minutes.
/*

C function : g_date_time_add_minutes
*/
func (recv *DateTime) AddMinutes(minutes int32) *DateTime {
	c_minutes := (C.gint)(minutes)

	retC := C.g_date_time_add_minutes((*C.GDateTime)(recv.native), c_minutes)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @datetime and adds the specified number of months to the
// copy. Add negative values to subtract months.
//
// The day of the month of the resulting #GDateTime is clamped to the number
// of days in the updated calendar month. For example, if adding 1 month to
// 31st January 2018, the result would be 28th February 2018. In 2020 (a leap
// year), the result would be 29th February.
/*

C function : g_date_time_add_months
*/
func (recv *DateTime) AddMonths(months int32) *DateTime {
	c_months := (C.gint)(months)

	retC := C.g_date_time_add_months((*C.GDateTime)(recv.native), c_months)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @datetime and adds the specified number of seconds.
// Add negative values to subtract seconds.
/*

C function : g_date_time_add_seconds
*/
func (recv *DateTime) AddSeconds(seconds float64) *DateTime {
	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_seconds((*C.GDateTime)(recv.native), c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @datetime and adds the specified number of weeks to the
// copy. Add negative values to subtract weeks.
/*

C function : g_date_time_add_weeks
*/
func (recv *DateTime) AddWeeks(weeks int32) *DateTime {
	c_weeks := (C.gint)(weeks)

	retC := C.g_date_time_add_weeks((*C.GDateTime)(recv.native), c_weeks)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a copy of @datetime and adds the specified number of years to the
// copy. Add negative values to subtract years.
//
// As with g_date_time_add_months(), if the resulting date would be 29th
// February on a non-leap year, the day will be clamped to 28th February.
/*

C function : g_date_time_add_years
*/
func (recv *DateTime) AddYears(years int32) *DateTime {
	c_years := (C.gint)(years)

	retC := C.g_date_time_add_years((*C.GDateTime)(recv.native), c_years)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Calculates the difference in time between @end and @begin.  The
// #GTimeSpan that is returned is effectively @end - @begin (ie:
// positive if the first parameter is larger).
/*

C function : g_date_time_difference
*/
func (recv *DateTime) Difference(begin *DateTime) TimeSpan {
	c_begin := (*C.GDateTime)(C.NULL)
	if begin != nil {
		c_begin = (*C.GDateTime)(begin.ToC())
	}

	retC := C.g_date_time_difference((*C.GDateTime)(recv.native), c_begin)
	retGo := (TimeSpan)(retC)

	return retGo
}

// Creates a newly allocated string representing the requested @format.
//
// The format strings understood by this function are a subset of the
// strftime() format language as specified by C99.  The \%D, \%U and \%W
// conversions are not supported, nor is the 'E' modifier.  The GNU
// extensions \%k, \%l, \%s and \%P are supported, however, as are the
// '0', '_' and '-' modifiers.
//
// In contrast to strftime(), this function always produces a UTF-8
// string, regardless of the current locale.  Note that the rendering of
// many formats is locale-dependent and may not match the strftime()
// output exactly.
//
// The following format specifiers are supported:
//
// - \%a: the abbreviated weekday name according to the current locale
// - \%A: the full weekday name according to the current locale
// - \%b: the abbreviated month name according to the current locale
// - \%B: the full month name according to the current locale
// - \%c: the preferred date and time representation for the current locale
// - \%C: the century number (year/100) as a 2-digit integer (00-99)
// - \%d: the day of the month as a decimal number (range 01 to 31)
// - \%e: the day of the month as a decimal number (range  1 to 31)
// - \%F: equivalent to `%Y-%m-%d` (the ISO 8601 date format)
// - \%g: the last two digits of the ISO 8601 week-based year as a
// decimal number (00-99). This works well with \%V and \%u.
// - \%G: the ISO 8601 week-based year as a decimal number. This works
// well with \%V and \%u.
// - \%h: equivalent to \%b
// - \%H: the hour as a decimal number using a 24-hour clock (range 00 to 23)
// - \%I: the hour as a decimal number using a 12-hour clock (range 01 to 12)
// - \%j: the day of the year as a decimal number (range 001 to 366)
// - \%k: the hour (24-hour clock) as a decimal number (range 0 to 23);
// single digits are preceded by a blank
// - \%l: the hour (12-hour clock) as a decimal number (range 1 to 12);
// single digits are preceded by a blank
// - \%m: the month as a decimal number (range 01 to 12)
// - \%M: the minute as a decimal number (range 00 to 59)
// - \%p: either "AM" or "PM" according to the given time value, or the
// corresponding  strings for the current locale.  Noon is treated as
// "PM" and midnight as "AM".
// - \%P: like \%p but lowercase: "am" or "pm" or a corresponding string for
// the current locale
// - \%r: the time in a.m. or p.m. notation
// - \%R: the time in 24-hour notation (\%H:\%M)
// - \%s: the number of seconds since the Epoch, that is, since 1970-01-01
// 00:00:00 UTC
// - \%S: the second as a decimal number (range 00 to 60)
// - \%t: a tab character
// - \%T: the time in 24-hour notation with seconds (\%H:\%M:\%S)
// - \%u: the ISO 8601 standard day of the week as a decimal, range 1 to 7,
// Monday being 1. This works well with \%G and \%V.
// - \%V: the ISO 8601 standard week number of the current year as a decimal
// number, range 01 to 53, where week 1 is the first week that has at
// least 4 days in the new year. See g_date_time_get_week_of_year().
// This works well with \%G and \%u.
// - \%w: the day of the week as a decimal, range 0 to 6, Sunday being 0.
// This is not the ISO 8601 standard format -- use \%u instead.
// - \%x: the preferred date representation for the current locale without
// the time
// - \%X: the preferred time representation for the current locale without
// the date
// - \%y: the year as a decimal number without the century
// - \%Y: the year as a decimal number including the century
// - \%z: the time zone as an offset from UTC (+hhmm)
// - \%:z: the time zone as an offset from UTC (+hh:mm).
// This is a gnulib strftime() extension. Since: 2.38
// - \%::z: the time zone as an offset from UTC (+hh:mm:ss). This is a
// gnulib strftime() extension. Since: 2.38
// - \%:::z: the time zone as an offset from UTC, with : to necessary
// precision (e.g., -04, +05:30). This is a gnulib strftime() extension. Since: 2.38
// - \%Z: the time zone or name or abbreviation
// - \%\%: a literal \% character
//
// Some conversion specifications can be modified by preceding the
// conversion specifier by one or more modifier characters. The
// following modifiers are supported for many of the numeric
// conversions:
//
// - O: Use alternative numeric symbols, if the current locale supports those.
// - _: Pad a numeric result with spaces. This overrides the default padding
// for the specifier.
// - -: Do not pad a numeric result. This overrides the default padding
// for the specifier.
// - 0: Pad a numeric result with zeros. This overrides the default padding
// for the specifier.
//
// Additionally, when O is used with B, b, or h, it produces the alternative
// form of a month name. The alternative form should be used when the month
// name is used without a day number (e.g., standalone). It is required in
// some languages (Baltic, Slavic, Greek, and more) due to their grammatical
// rules. For other languages there is no difference. \%OB is a GNU and BSD
// strftime() extension expected to be added to the future POSIX specification,
// \%Ob and \%Oh are GNU strftime() extensions. Since: 2.56
/*

C function : g_date_time_format
*/
func (recv *DateTime) Format(format string) string {
	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	retC := C.g_date_time_format((*C.GDateTime)(recv.native), c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the day of the month represented by @datetime in the gregorian
// calendar.
/*

C function : g_date_time_get_day_of_month
*/
func (recv *DateTime) GetDayOfMonth() int32 {
	retC := C.g_date_time_get_day_of_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the ISO 8601 day of the week on which @datetime falls (1 is
// Monday, 2 is Tuesday... 7 is Sunday).
/*

C function : g_date_time_get_day_of_week
*/
func (recv *DateTime) GetDayOfWeek() int32 {
	retC := C.g_date_time_get_day_of_week((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the day of the year represented by @datetime in the Gregorian
// calendar.
/*

C function : g_date_time_get_day_of_year
*/
func (recv *DateTime) GetDayOfYear() int32 {
	retC := C.g_date_time_get_day_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the hour of the day represented by @datetime
/*

C function : g_date_time_get_hour
*/
func (recv *DateTime) GetHour() int32 {
	retC := C.g_date_time_get_hour((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the microsecond of the date represented by @datetime
/*

C function : g_date_time_get_microsecond
*/
func (recv *DateTime) GetMicrosecond() int32 {
	retC := C.g_date_time_get_microsecond((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the minute of the hour represented by @datetime
/*

C function : g_date_time_get_minute
*/
func (recv *DateTime) GetMinute() int32 {
	retC := C.g_date_time_get_minute((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the month of the year represented by @datetime in the Gregorian
// calendar.
/*

C function : g_date_time_get_month
*/
func (recv *DateTime) GetMonth() int32 {
	retC := C.g_date_time_get_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the second of the minute represented by @datetime
/*

C function : g_date_time_get_second
*/
func (recv *DateTime) GetSecond() int32 {
	retC := C.g_date_time_get_second((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the number of seconds since the start of the last minute,
// including the fractional part.
/*

C function : g_date_time_get_seconds
*/
func (recv *DateTime) GetSeconds() float64 {
	retC := C.g_date_time_get_seconds((*C.GDateTime)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Determines the time zone abbreviation to be used at the time and in
// the time zone of @datetime.
//
// For example, in Toronto this is currently "EST" during the winter
// months and "EDT" during the summer months when daylight savings
// time is in effect.
/*

C function : g_date_time_get_timezone_abbreviation
*/
func (recv *DateTime) GetTimezoneAbbreviation() string {
	retC := C.g_date_time_get_timezone_abbreviation((*C.GDateTime)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Determines the offset to UTC in effect at the time and in the time
// zone of @datetime.
//
// The offset is the number of microseconds that you add to UTC time to
// arrive at local time for the time zone (ie: negative numbers for time
// zones west of GMT, positive numbers for east).
//
// If @datetime represents UTC time, then the offset is always zero.
/*

C function : g_date_time_get_utc_offset
*/
func (recv *DateTime) GetUtcOffset() TimeSpan {
	retC := C.g_date_time_get_utc_offset((*C.GDateTime)(recv.native))
	retGo := (TimeSpan)(retC)

	return retGo
}

// Returns the ISO 8601 week-numbering year in which the week containing
// @datetime falls.
//
// This function, taken together with g_date_time_get_week_of_year() and
// g_date_time_get_day_of_week() can be used to determine the full ISO
// week date on which @datetime falls.
//
// This is usually equal to the normal Gregorian year (as returned by
// g_date_time_get_year()), except as detailed below:
//
// For Thursday, the week-numbering year is always equal to the usual
// calendar year.  For other days, the number is such that every day
// within a complete week (Monday to Sunday) is contained within the
// same week-numbering year.
//
// For Monday, Tuesday and Wednesday occurring near the end of the year,
// this may mean that the week-numbering year is one greater than the
// calendar year (so that these days have the same week-numbering year
// as the Thursday occurring early in the next year).
//
// For Friday, Saturday and Sunday occurring near the start of the year,
// this may mean that the week-numbering year is one less than the
// calendar year (so that these days have the same week-numbering year
// as the Thursday occurring late in the previous year).
//
// An equivalent description is that the week-numbering year is equal to
// the calendar year containing the majority of the days in the current
// week (Monday to Sunday).
//
// Note that January 1 0001 in the proleptic Gregorian calendar is a
// Monday, so this function never returns 0.
/*

C function : g_date_time_get_week_numbering_year
*/
func (recv *DateTime) GetWeekNumberingYear() int32 {
	retC := C.g_date_time_get_week_numbering_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the ISO 8601 week number for the week containing @datetime.
// The ISO 8601 week number is the same for every day of the week (from
// Moday through Sunday).  That can produce some unusual results
// (described below).
//
// The first week of the year is week 1.  This is the week that contains
// the first Thursday of the year.  Equivalently, this is the first week
// that has more than 4 of its days falling within the calendar year.
//
// The value 0 is never returned by this function.  Days contained
// within a year but occurring before the first ISO 8601 week of that
// year are considered as being contained in the last week of the
// previous year.  Similarly, the final days of a calendar year may be
// considered as being part of the first ISO 8601 week of the next year
// if 4 or more days of that week are contained within the new year.
/*

C function : g_date_time_get_week_of_year
*/
func (recv *DateTime) GetWeekOfYear() int32 {
	retC := C.g_date_time_get_week_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the year represented by @datetime in the Gregorian calendar.
/*

C function : g_date_time_get_year
*/
func (recv *DateTime) GetYear() int32 {
	retC := C.g_date_time_get_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the Gregorian day, month, and year of a given #GDateTime.
/*

C function : g_date_time_get_ymd
*/
func (recv *DateTime) GetYmd() (int32, int32, int32) {
	var c_year C.gint

	var c_month C.gint

	var c_day C.gint

	C.g_date_time_get_ymd((*C.GDateTime)(recv.native), &c_year, &c_month, &c_day)

	year := (int32)(c_year)

	month := (int32)(c_month)

	day := (int32)(c_day)

	return year, month, day
}

// Determines if daylight savings time is in effect at the time and in
// the time zone of @datetime.
/*

C function : g_date_time_is_daylight_savings
*/
func (recv *DateTime) IsDaylightSavings() bool {
	retC := C.g_date_time_is_daylight_savings((*C.GDateTime)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Atomically increments the reference count of @datetime by one.
/*

C function : g_date_time_ref
*/
func (recv *DateTime) Ref() *DateTime {
	retC := C.g_date_time_ref((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GDateTime corresponding to the same instant in time as
// @datetime, but in the local time zone.
//
// This call is equivalent to calling g_date_time_to_timezone() with the
// time zone returned by g_time_zone_new_local().
/*

C function : g_date_time_to_local
*/
func (recv *DateTime) ToLocal() *DateTime {
	retC := C.g_date_time_to_local((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Stores the instant in time that @datetime represents into @tv.
//
// The time contained in a #GTimeVal is always stored in the form of
// seconds elapsed since 1970-01-01 00:00:00 UTC, regardless of the time
// zone associated with @datetime.
//
// On systems where 'long' is 32bit (ie: all 32bit systems and all
// Windows systems), a #GTimeVal is incapable of storing the entire
// range of values that #GDateTime is capable of expressing.  On those
// systems, this function returns %FALSE to indicate that the time is
// out of range.
//
// On systems where 'long' is 64bit, this function never fails.
/*

C function : g_date_time_to_timeval
*/
func (recv *DateTime) ToTimeval(tv *TimeVal) bool {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_to_timeval((*C.GDateTime)(recv.native), c_tv)
	retGo := retC == C.TRUE

	return retGo
}

// Create a new #GDateTime corresponding to the same instant in time as
// @datetime, but in the time zone @tz.
//
// This call can fail in the case that the time goes out of bounds.  For
// example, converting 0001-01-01 00:00:00 UTC to a time zone west of
// Greenwich will fail (due to the year 0 being out of range).
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_to_timezone
*/
func (recv *DateTime) ToTimezone(tz *TimeZone) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	retC := C.g_date_time_to_timezone((*C.GDateTime)(recv.native), c_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gives the Unix time corresponding to @datetime, rounding down to the
// nearest second.
//
// Unix time is the number of seconds that have elapsed since 1970-01-01
// 00:00:00 UTC, regardless of the time zone associated with @datetime.
/*

C function : g_date_time_to_unix
*/
func (recv *DateTime) ToUnix() int64 {
	retC := C.g_date_time_to_unix((*C.GDateTime)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Creates a new #GDateTime corresponding to the same instant in time as
// @datetime, but in UTC.
//
// This call is equivalent to calling g_date_time_to_timezone() with the
// time zone returned by g_time_zone_new_utc().
/*

C function : g_date_time_to_utc
*/
func (recv *DateTime) ToUtc() *DateTime {
	retC := C.g_date_time_to_utc((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Atomically decrements the reference count of @datetime by one.
//
// When the reference count reaches zero, the resources allocated by
// @datetime are freed
/*

C function : g_date_time_unref
*/
func (recv *DateTime) Unref() {
	C.g_date_time_unref((*C.GDateTime)(recv.native))

	return
}

// Returns the value associated with @key under @group_name as a signed
// 64-bit integer. This is similar to g_key_file_get_integer() but can return
// 64-bit results without truncation.
/*

C function : g_key_file_get_int64
*/
func (recv *KeyFile) GetInt64(groupName string, key string) (int64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_int64((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Returns the value associated with @key under @group_name as an unsigned
// 64-bit integer. This is similar to g_key_file_get_integer() but can return
// large positive results without truncation.
/*

C function : g_key_file_get_uint64
*/
func (recv *KeyFile) GetUint64(groupName string, key string) (uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_uint64((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (uint64)(retC)

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Associates a new integer value with @key under @group_name.
// If @key cannot be found then it is created.
/*

C function : g_key_file_set_int64
*/
func (recv *KeyFile) SetInt64(groupName string, key string, value int64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint64)(value)

	C.g_key_file_set_int64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Associates a new integer value with @key under @group_name.
// If @key cannot be found then it is created.
/*

C function : g_key_file_set_uint64
*/
func (recv *KeyFile) SetUint64(groupName string, key string, value uint64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint64)(value)

	C.g_key_file_set_uint64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Returns the compile options that @regex was created with.
//
// Depending on the version of PCRE that is used, this may or may not
// include flags set by option expressions such as `(?i)` found at the
// top-level within the compiled pattern.
/*

C function : g_regex_get_compile_flags
*/
func (recv *Regex) GetCompileFlags() RegexCompileFlags {
	retC := C.g_regex_get_compile_flags((*C.GRegex)(recv.native))
	retGo := (RegexCompileFlags)(retC)

	return retGo
}

// Returns the match options that @regex was created with.
/*

C function : g_regex_get_match_flags
*/
func (recv *Regex) GetMatchFlags() RegexMatchFlags {
	retC := C.g_regex_get_match_flags((*C.GRegex)(recv.native))
	retGo := (RegexMatchFlags)(retC)

	return retGo
}

// Gets a name for the source, used in debugging and profiling.  The
// name may be #NULL if it has never been set with g_source_set_name().
/*

C function : g_source_get_name
*/
func (recv *Source) GetName() string {
	retC := C.g_source_get_name((*C.GSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets a name for the source, used in debugging and profiling.
// The name defaults to #NULL.
//
// The source name should describe in a human-readable way
// what the source does. For example, "X11 event queue"
// or "GTK+ repaint idle handler" or whatever it is.
//
// It is permitted to call this function multiple times, but is not
// recommended due to the potential performance impact.  For example,
// one could change the name in the "check" function of a #GSourceFuncs
// to include details like the event type in the source name.
//
// Use caution if changing the name while another thread may be
// accessing it with g_source_get_name(); that function does not copy
// the value, and changing the value will free it while the other thread
// may be attempting to use it.
/*

C function : g_source_set_name
*/
func (recv *Source) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_source_set_name((*C.GSource)(recv.native), c_name)

	return
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

// Creates a #GTimeZone corresponding to @identifier.
//
// @identifier can either be an RFC3339/ISO 8601 time offset or
// something that would pass as a valid value for the `TZ` environment
// variable (including %NULL).
//
// In Windows, @identifier can also be the unlocalized name of a time
// zone for standard time, for example "Pacific Standard Time".
//
// Valid RFC3339 time offsets are `"Z"` (for UTC) or
// `"±hh:mm"`.  ISO 8601 additionally specifies
// `"±hhmm"` and `"±hh"`.  Offsets are
// time values to be added to Coordinated Universal Time (UTC) to get
// the local time.
//
// In UNIX, the `TZ` environment variable typically corresponds
// to the name of a file in the zoneinfo database, or string in
// "std offset [dst [offset],start[/time],end[/time]]" (POSIX) format.
// There  are  no spaces in the specification. The name of standard
// and daylight savings time zone must be three or more alphabetic
// characters. Offsets are time values to be added to local time to
// get Coordinated Universal Time (UTC) and should be
// `"[±]hh[[:]mm[:ss]]"`.  Dates are either
// `"Jn"` (Julian day with n between 1 and 365, leap
// years not counted), `"n"` (zero-based Julian day
// with n between 0 and 365) or `"Mm.w.d"` (day d
// (0 <= d <= 6) of week w (1 <= w <= 5) of month m (1 <= m <= 12), day
// 0 is a Sunday).  Times are in local wall clock time, the default is
// 02:00:00.
//
// In Windows, the "tzn[+|–]hh[:mm[:ss]][dzn]" format is used, but also
// accepts POSIX format.  The Windows format uses US rules for all time
// zones; daylight savings time is 60 minutes behind the standard time
// with date and time of change taken from Pacific Standard Time.
// Offsets are time values to be added to the local time to get
// Coordinated Universal Time (UTC).
//
// g_time_zone_new_local() calls this function with the value of the
// `TZ` environment variable. This function itself is independent of
// the value of `TZ`, but if @identifier is %NULL then `/etc/localtime`
// will be consulted to discover the correct time zone on UNIX and the
// registry will be consulted or GetTimeZoneInformation() will be used
// to get the local time zone on Windows.
//
// If intervals are not available, only time zone rules from `TZ`
// environment variable or other means, then they will be computed
// from year 1900 to 2037.  If the maximum year for the rules is
// available and it is greater than 2037, then it will followed
// instead.
//
// See
// [RFC3339 §5.6](http://tools.ietf.org/html/rfc3339#section-5.6)
// for a precise definition of valid RFC3339 time offsets
// (the `time-offset` expansion) and ISO 8601 for the
// full list of valid time offsets.  See
// [The GNU C Library manual](http://www.gnu.org/s/libc/manual/html_node/TZ-Variable.html)
// for an explanation of the possible
// values of the `TZ` environment variable. See
// [Microsoft Time Zone Index Values](http://msdn.microsoft.com/en-us/library/ms912391%28v=winembedded.11%29.aspx)
// for the list of time zones on Windows.
//
// You should release the return value by calling g_time_zone_unref()
// when you are done with it.
/*

C function : g_time_zone_new
*/
func TimeZoneNew(identifier string) *TimeZone {
	c_identifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(c_identifier))

	retC := C.g_time_zone_new(c_identifier)
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GTimeZone corresponding to local time.  The local time
// zone may change between invocations to this function; for example,
// if the system administrator changes it.
//
// This is equivalent to calling g_time_zone_new() with the value of
// the `TZ` environment variable (including the possibility of %NULL).
//
// You should release the return value by calling g_time_zone_unref()
// when you are done with it.
/*

C function : g_time_zone_new_local
*/
func TimeZoneNewLocal() *TimeZone {
	retC := C.g_time_zone_new_local()
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GTimeZone corresponding to UTC.
//
// This is equivalent to calling g_time_zone_new() with a value like
// "Z", "UTC", "+00", etc.
//
// You should release the return value by calling g_time_zone_unref()
// when you are done with it.
/*

C function : g_time_zone_new_utc
*/
func TimeZoneNewUtc() *TimeZone {
	retC := C.g_time_zone_new_utc()
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finds an interval within @tz that corresponds to the given @time_,
// possibly adjusting @time_ if required to fit into an interval.
// The meaning of @time_ depends on @type.
//
// This function is similar to g_time_zone_find_interval(), with the
// difference that it always succeeds (by making the adjustments
// described below).
//
// In any of the cases where g_time_zone_find_interval() succeeds then
// this function returns the same value, without modifying @time_.
//
// This function may, however, modify @time_ in order to deal with
// non-existent times.  If the non-existent local @time_ of 02:30 were
// requested on March 14th 2010 in Toronto then this function would
// adjust @time_ to be 03:00 and return the interval containing the
// adjusted time.
/*

C function : g_time_zone_adjust_time
*/
func (recv *TimeZone) AdjustTime(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_adjust_time((*C.GTimeZone)(recv.native), c_type, &c_time_)
	retGo := (int32)(retC)

	return retGo
}

// Finds an the interval within @tz that corresponds to the given @time_.
// The meaning of @time_ depends on @type.
//
// If @type is %G_TIME_TYPE_UNIVERSAL then this function will always
// succeed (since universal time is monotonic and continuous).
//
// Otherwise @time_ is treated as local time.  The distinction between
// %G_TIME_TYPE_STANDARD and %G_TIME_TYPE_DAYLIGHT is ignored except in
// the case that the given @time_ is ambiguous.  In Toronto, for example,
// 01:30 on November 7th 2010 occurred twice (once inside of daylight
// savings time and the next, an hour later, outside of daylight savings
// time).  In this case, the different value of @type would result in a
// different interval being returned.
//
// It is still possible for this function to fail.  In Toronto, for
// example, 02:00 on March 14th 2010 does not exist (due to the leap
// forward to begin daylight savings time).  -1 is returned in that
// case.
/*

C function : g_time_zone_find_interval
*/
func (recv *TimeZone) FindInterval(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_find_interval((*C.GTimeZone)(recv.native), c_type, c_time_)
	retGo := (int32)(retC)

	return retGo
}

// Determines the time zone abbreviation to be used during a particular
// @interval of time in the time zone @tz.
//
// For example, in Toronto this is currently "EST" during the winter
// months and "EDT" during the summer months when daylight savings time
// is in effect.
/*

C function : g_time_zone_get_abbreviation
*/
func (recv *TimeZone) GetAbbreviation(interval int32) string {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_abbreviation((*C.GTimeZone)(recv.native), c_interval)
	retGo := C.GoString(retC)

	return retGo
}

// Determines the offset to UTC in effect during a particular @interval
// of time in the time zone @tz.
//
// The offset is the number of seconds that you add to UTC time to
// arrive at local time for @tz (ie: negative numbers for time zones
// west of GMT, positive numbers for east).
/*

C function : g_time_zone_get_offset
*/
func (recv *TimeZone) GetOffset(interval int32) int32 {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_offset((*C.GTimeZone)(recv.native), c_interval)
	retGo := (int32)(retC)

	return retGo
}

// Determines if daylight savings time is in effect during a particular
// @interval of time in the time zone @tz.
/*

C function : g_time_zone_is_dst
*/
func (recv *TimeZone) IsDst(interval int32) bool {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_is_dst((*C.GTimeZone)(recv.native), c_interval)
	retGo := retC == C.TRUE

	return retGo
}

// Increases the reference count on @tz.
/*

C function : g_time_zone_ref
*/
func (recv *TimeZone) Ref() *TimeZone {
	retC := C.g_time_zone_ref((*C.GTimeZone)(recv.native))
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decreases the reference count on @tz.
/*

C function : g_time_zone_unref
*/
func (recv *TimeZone) Unref() {
	C.g_time_zone_unref((*C.GTimeZone)(recv.native))

	return
}

// Unsupported : g_variant_builder_add_parsed : unsupported parameter ... : varargs

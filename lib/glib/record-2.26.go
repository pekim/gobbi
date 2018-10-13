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

// Unsupported : g_bytes_new : unsupported parameter data : no param type

// Unsupported : g_bytes_new_static : unsupported parameter data : no param type

// Unsupported : g_bytes_new_take : unsupported parameter data : no param type

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no param type

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
	c_tz := (*C.GTimeZone)(tz.ToC())

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
	c_tv := (*C.GTimeVal)(tv.ToC())

	retC := C.g_date_time_new_from_timeval_local(c_tv)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromTimevalUtc is a wrapper around the C function g_date_time_new_from_timeval_utc.
func DateTimeNewFromTimevalUtc(tv *TimeVal) *DateTime {
	c_tv := (*C.GTimeVal)(tv.ToC())

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
	c_tz := (*C.GTimeZone)(tz.ToC())

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

// Add is a wrapper around the C function g_date_time_add.
func (recv *DateTime) Add(timespan TimeSpan) *DateTime {
	c_timespan := (C.GTimeSpan)(timespan)

	retC := C.g_date_time_add((*C.GDateTime)(recv.native), c_timespan)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddDays is a wrapper around the C function g_date_time_add_days.
func (recv *DateTime) AddDays(days int32) *DateTime {
	c_days := (C.gint)(days)

	retC := C.g_date_time_add_days((*C.GDateTime)(recv.native), c_days)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFull is a wrapper around the C function g_date_time_add_full.
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

// AddHours is a wrapper around the C function g_date_time_add_hours.
func (recv *DateTime) AddHours(hours int32) *DateTime {
	c_hours := (C.gint)(hours)

	retC := C.g_date_time_add_hours((*C.GDateTime)(recv.native), c_hours)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddMinutes is a wrapper around the C function g_date_time_add_minutes.
func (recv *DateTime) AddMinutes(minutes int32) *DateTime {
	c_minutes := (C.gint)(minutes)

	retC := C.g_date_time_add_minutes((*C.GDateTime)(recv.native), c_minutes)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddMonths is a wrapper around the C function g_date_time_add_months.
func (recv *DateTime) AddMonths(months int32) *DateTime {
	c_months := (C.gint)(months)

	retC := C.g_date_time_add_months((*C.GDateTime)(recv.native), c_months)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddSeconds is a wrapper around the C function g_date_time_add_seconds.
func (recv *DateTime) AddSeconds(seconds float64) *DateTime {
	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_seconds((*C.GDateTime)(recv.native), c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddWeeks is a wrapper around the C function g_date_time_add_weeks.
func (recv *DateTime) AddWeeks(weeks int32) *DateTime {
	c_weeks := (C.gint)(weeks)

	retC := C.g_date_time_add_weeks((*C.GDateTime)(recv.native), c_weeks)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddYears is a wrapper around the C function g_date_time_add_years.
func (recv *DateTime) AddYears(years int32) *DateTime {
	c_years := (C.gint)(years)

	retC := C.g_date_time_add_years((*C.GDateTime)(recv.native), c_years)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Difference is a wrapper around the C function g_date_time_difference.
func (recv *DateTime) Difference(begin *DateTime) TimeSpan {
	c_begin := (*C.GDateTime)(begin.ToC())

	retC := C.g_date_time_difference((*C.GDateTime)(recv.native), c_begin)
	retGo := (TimeSpan)(retC)

	return retGo
}

// Format is a wrapper around the C function g_date_time_format.
func (recv *DateTime) Format(format string) string {
	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	retC := C.g_date_time_format((*C.GDateTime)(recv.native), c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDayOfMonth is a wrapper around the C function g_date_time_get_day_of_month.
func (recv *DateTime) GetDayOfMonth() int32 {
	retC := C.g_date_time_get_day_of_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfWeek is a wrapper around the C function g_date_time_get_day_of_week.
func (recv *DateTime) GetDayOfWeek() int32 {
	retC := C.g_date_time_get_day_of_week((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfYear is a wrapper around the C function g_date_time_get_day_of_year.
func (recv *DateTime) GetDayOfYear() int32 {
	retC := C.g_date_time_get_day_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHour is a wrapper around the C function g_date_time_get_hour.
func (recv *DateTime) GetHour() int32 {
	retC := C.g_date_time_get_hour((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMicrosecond is a wrapper around the C function g_date_time_get_microsecond.
func (recv *DateTime) GetMicrosecond() int32 {
	retC := C.g_date_time_get_microsecond((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinute is a wrapper around the C function g_date_time_get_minute.
func (recv *DateTime) GetMinute() int32 {
	retC := C.g_date_time_get_minute((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMonth is a wrapper around the C function g_date_time_get_month.
func (recv *DateTime) GetMonth() int32 {
	retC := C.g_date_time_get_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSecond is a wrapper around the C function g_date_time_get_second.
func (recv *DateTime) GetSecond() int32 {
	retC := C.g_date_time_get_second((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSeconds is a wrapper around the C function g_date_time_get_seconds.
func (recv *DateTime) GetSeconds() float64 {
	retC := C.g_date_time_get_seconds((*C.GDateTime)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetTimezoneAbbreviation is a wrapper around the C function g_date_time_get_timezone_abbreviation.
func (recv *DateTime) GetTimezoneAbbreviation() string {
	retC := C.g_date_time_get_timezone_abbreviation((*C.GDateTime)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUtcOffset is a wrapper around the C function g_date_time_get_utc_offset.
func (recv *DateTime) GetUtcOffset() TimeSpan {
	retC := C.g_date_time_get_utc_offset((*C.GDateTime)(recv.native))
	retGo := (TimeSpan)(retC)

	return retGo
}

// GetWeekNumberingYear is a wrapper around the C function g_date_time_get_week_numbering_year.
func (recv *DateTime) GetWeekNumberingYear() int32 {
	retC := C.g_date_time_get_week_numbering_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWeekOfYear is a wrapper around the C function g_date_time_get_week_of_year.
func (recv *DateTime) GetWeekOfYear() int32 {
	retC := C.g_date_time_get_week_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetYear is a wrapper around the C function g_date_time_get_year.
func (recv *DateTime) GetYear() int32 {
	retC := C.g_date_time_get_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetYmd is a wrapper around the C function g_date_time_get_ymd.
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

// IsDaylightSavings is a wrapper around the C function g_date_time_is_daylight_savings.
func (recv *DateTime) IsDaylightSavings() bool {
	retC := C.g_date_time_is_daylight_savings((*C.GDateTime)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_date_time_ref.
func (recv *DateTime) Ref() *DateTime {
	retC := C.g_date_time_ref((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToLocal is a wrapper around the C function g_date_time_to_local.
func (recv *DateTime) ToLocal() *DateTime {
	retC := C.g_date_time_to_local((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToTimeval is a wrapper around the C function g_date_time_to_timeval.
func (recv *DateTime) ToTimeval(tv *TimeVal) bool {
	c_tv := (*C.GTimeVal)(tv.ToC())

	retC := C.g_date_time_to_timeval((*C.GDateTime)(recv.native), c_tv)
	retGo := retC == C.TRUE

	return retGo
}

// ToTimezone is a wrapper around the C function g_date_time_to_timezone.
func (recv *DateTime) ToTimezone(tz *TimeZone) *DateTime {
	c_tz := (*C.GTimeZone)(tz.ToC())

	retC := C.g_date_time_to_timezone((*C.GDateTime)(recv.native), c_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToUnix is a wrapper around the C function g_date_time_to_unix.
func (recv *DateTime) ToUnix() int64 {
	retC := C.g_date_time_to_unix((*C.GDateTime)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// ToUtc is a wrapper around the C function g_date_time_to_utc.
func (recv *DateTime) ToUtc() *DateTime {
	retC := C.g_date_time_to_utc((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_date_time_unref.
func (recv *DateTime) Unref() {
	C.g_date_time_unref((*C.GDateTime)(recv.native))

	return
}

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// GetInt64 is a wrapper around the C function g_key_file_get_int64.
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

// GetUint64 is a wrapper around the C function g_key_file_get_uint64.
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

// SetInt64 is a wrapper around the C function g_key_file_set_int64.
func (recv *KeyFile) SetInt64(groupName string, key string, value int64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint64)(value)

	C.g_key_file_set_int64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetUint64 is a wrapper around the C function g_key_file_set_uint64.
func (recv *KeyFile) SetUint64(groupName string, key string, value uint64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint64)(value)

	C.g_key_file_set_uint64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// GetCompileFlags is a wrapper around the C function g_regex_get_compile_flags.
func (recv *Regex) GetCompileFlags() RegexCompileFlags {
	retC := C.g_regex_get_compile_flags((*C.GRegex)(recv.native))
	retGo := (RegexCompileFlags)(retC)

	return retGo
}

// GetMatchFlags is a wrapper around the C function g_regex_get_match_flags.
func (recv *Regex) GetMatchFlags() RegexMatchFlags {
	retC := C.g_regex_get_match_flags((*C.GRegex)(recv.native))
	retGo := (RegexMatchFlags)(retC)

	return retGo
}

// GetName is a wrapper around the C function g_source_get_name.
func (recv *Source) GetName() string {
	retC := C.g_source_get_name((*C.GSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetName is a wrapper around the C function g_source_set_name.
func (recv *Source) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_source_set_name((*C.GSource)(recv.native), c_name)

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

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

// AdjustTime is a wrapper around the C function g_time_zone_adjust_time.
func (recv *TimeZone) AdjustTime(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_adjust_time((*C.GTimeZone)(recv.native), c_type, &c_time_)
	retGo := (int32)(retC)

	return retGo
}

// FindInterval is a wrapper around the C function g_time_zone_find_interval.
func (recv *TimeZone) FindInterval(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_find_interval((*C.GTimeZone)(recv.native), c_type, c_time_)
	retGo := (int32)(retC)

	return retGo
}

// GetAbbreviation is a wrapper around the C function g_time_zone_get_abbreviation.
func (recv *TimeZone) GetAbbreviation(interval int32) string {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_abbreviation((*C.GTimeZone)(recv.native), c_interval)
	retGo := C.GoString(retC)

	return retGo
}

// GetOffset is a wrapper around the C function g_time_zone_get_offset.
func (recv *TimeZone) GetOffset(interval int32) int32 {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_offset((*C.GTimeZone)(recv.native), c_interval)
	retGo := (int32)(retC)

	return retGo
}

// IsDst is a wrapper around the C function g_time_zone_is_dst.
func (recv *TimeZone) IsDst(interval int32) bool {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_is_dst((*C.GTimeZone)(recv.native), c_interval)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_time_zone_ref.
func (recv *TimeZone) Ref() *TimeZone {
	retC := C.g_time_zone_ref((*C.GTimeZone)(recv.native))
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_time_zone_unref.
func (recv *TimeZone) Unref() {
	C.g_time_zone_unref((*C.GTimeZone)(recv.native))

	return
}

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_add_parsed : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant

// This is a generated file - DO NOT EDIT
// +build glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Copies a GDate to a newly-allocated GDate. If the input was invalid
// (as determined by g_date_valid()), the invalid state will be copied
// as is into the new object.
/*

C function : g_date_copy
*/
func (recv *Date) Copy() *Date {
	retC := C.g_date_copy((*C.GDate)(recv.native))
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GDateTime corresponding to the given
// [ISO 8601 formatted string](https://en.wikipedia.org/wiki/ISO_8601)
// @text. ISO 8601 strings of the form <date><sep><time><tz> are supported.
//
// <sep> is the separator and can be either 'T', 't' or ' '.
//
// <date> is in the form:
//
// - `YYYY-MM-DD` - Year/month/day, e.g. 2016-08-24.
// - `YYYYMMDD` - Same as above without dividers.
// - `YYYY-DDD` - Ordinal day where DDD is from 001 to 366, e.g. 2016-237.
// - `YYYYDDD` - Same as above without dividers.
// - `YYYY-Www-D` - Week day where ww is from 01 to 52 and D from 1-7,
// e.g. 2016-W34-3.
// - `YYYYWwwD` - Same as above without dividers.
//
// <time> is in the form:
//
// - `hh:mm:ss(.sss)` - Hours, minutes, seconds (subseconds), e.g. 22:10:42.123.
// - `hhmmss(.sss)` - Same as above without dividers.
//
// <tz> is an optional timezone suffix of the form:
//
// - `Z` - UTC.
// - `+hh:mm` or `-hh:mm` - Offset from UTC in hours and minutes, e.g. +12:00.
// - `+hh` or `-hh` - Offset from UTC in hours, e.g. +12.
//
// If the timezone is not provided in @text it must be provided in @default_tz
// (this field is otherwise ignored).
//
// This call can fail (returning %NULL) if @text is not a valid ISO 8601
// formatted string.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
/*

C function : g_date_time_new_from_iso8601
*/
func DateTimeNewFromIso8601(text string, defaultTz *TimeZone) *DateTime {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_default_tz := (*C.GTimeZone)(C.NULL)
	if defaultTz != nil {
		c_default_tz = (*C.GTimeZone)(defaultTz.ToC())
	}

	retC := C.g_date_time_new_from_iso8601(c_text, c_default_tz)
	var retGo (*DateTime)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DateTimeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Returns the actual locale which the result of
// g_key_file_get_locale_string() or g_key_file_get_locale_string_list()
// came from.
//
// If calling g_key_file_get_locale_string() or
// g_key_file_get_locale_string_list() with exactly the same @key_file,
// @group_name, @key and @locale, the result of those functions will
// have originally been tagged with the locale that is the result of
// this function.
/*

C function : g_key_file_get_locale_for_key
*/
func (recv *KeyFile) GetLocaleForKey(groupName string, key string, locale string) string {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	retC := C.g_key_file_get_locale_for_key((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

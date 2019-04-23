// This is a generated file - DO NOT EDIT

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

// Blacklisted : GArray

// Blacklisted : GAsyncQueue

// Blacklisted : GBookmarkFile

// Blacklisted : GByteArray

// Blacklisted : GCond

// Blacklisted : GData

// Date is a wrapper around the C record GDate.
type Date struct {
	native *C.GDate
	// Bitfield not supported : 32 julian_days
	// Bitfield not supported :  1 julian
	// Bitfield not supported :  1 dmy
	// Bitfield not supported :  6 day
	// Bitfield not supported :  4 month
	// Bitfield not supported : 16 year
}

func DateNewFromC(u unsafe.Pointer) *Date {
	c := (*C.GDate)(u)
	if c == nil {
		return nil
	}

	g := &Date{native: c}

	return g
}

func (recv *Date) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Date with another Date, and returns true if they represent the same GObject.
func (recv *Date) Equals(other *Date) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_date_new

// DateNewDmy is a wrapper around the C function g_date_new_dmy.
func DateNewDmy(day DateDay, month DateMonth, year DateYear) *Date {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_new_dmy(c_day, c_month, c_year)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_date_new_julian

// Blacklisted : g_date_get_days_in_month

// Blacklisted : g_date_get_monday_weeks_in_year

// Blacklisted : g_date_get_sunday_weeks_in_year

// Blacklisted : g_date_is_leap_year

// Blacklisted : g_date_strftime

// Blacklisted : g_date_valid_day

// Blacklisted : g_date_valid_dmy

// Blacklisted : g_date_valid_julian

// Blacklisted : g_date_valid_month

// Blacklisted : g_date_valid_weekday

// Blacklisted : g_date_valid_year

// Blacklisted : g_date_add_days

// Blacklisted : g_date_add_months

// Blacklisted : g_date_add_years

// Blacklisted : g_date_clamp

// Blacklisted : g_date_clear

// Blacklisted : g_date_compare

// Blacklisted : g_date_days_between

// Blacklisted : g_date_free

// GetDay is a wrapper around the C function g_date_get_day.
func (recv *Date) GetDay() DateDay {
	retC := C.g_date_get_day((*C.GDate)(recv.native))
	retGo := (DateDay)(retC)

	return retGo
}

// GetDayOfYear is a wrapper around the C function g_date_get_day_of_year.
func (recv *Date) GetDayOfYear() uint32 {
	retC := C.g_date_get_day_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Blacklisted : g_date_get_julian

// Blacklisted : g_date_get_monday_week_of_year

// GetMonth is a wrapper around the C function g_date_get_month.
func (recv *Date) GetMonth() DateMonth {
	retC := C.g_date_get_month((*C.GDate)(recv.native))
	retGo := (DateMonth)(retC)

	return retGo
}

// Blacklisted : g_date_get_sunday_week_of_year

// Blacklisted : g_date_get_weekday

// GetYear is a wrapper around the C function g_date_get_year.
func (recv *Date) GetYear() DateYear {
	retC := C.g_date_get_year((*C.GDate)(recv.native))
	retGo := (DateYear)(retC)

	return retGo
}

// Blacklisted : g_date_is_first_of_month

// Blacklisted : g_date_is_last_of_month

// Blacklisted : g_date_order

// Blacklisted : g_date_set_day

// Blacklisted : g_date_set_dmy

// Blacklisted : g_date_set_julian

// Blacklisted : g_date_set_month

// Blacklisted : g_date_set_parse

// Blacklisted : g_date_set_time

// Blacklisted : g_date_set_year

// Blacklisted : g_date_subtract_days

// Blacklisted : g_date_subtract_months

// Blacklisted : g_date_subtract_years

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer (tm*) for param tm

// Blacklisted : g_date_valid

// Blacklisted : GDebugKey

// Blacklisted : GDir

// Error is a wrapper around the C record GError.
type Error struct {
	native  *C.GError
	Domain  Quark
	Code    int32
	Message string
}

func ErrorNewFromC(u unsafe.Pointer) *Error {
	c := (*C.GError)(u)
	if c == nil {
		return nil
	}

	g := &Error{
		Code:    (int32)(c.code),
		Domain:  (Quark)(c.domain),
		Message: C.GoString(c.message),
		native:  c,
	}

	return g
}

func (recv *Error) ToC() unsafe.Pointer {
	recv.native.domain =
		(C.GQuark)(recv.Domain)
	recv.native.code =
		(C.gint)(recv.Code)
	recv.native.message =
		C.CString(recv.Message)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Error with another Error, and returns true if they represent the same GObject.
func (recv *Error) Equals(other *Error) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_error_new

// Blacklisted : g_error_new_literal

// Blacklisted : g_error_copy

// Blacklisted : g_error_free

// Blacklisted : g_error_matches

// Blacklisted : GHashTable

// Blacklisted : GHashTableIter

// Blacklisted : GHook

// Blacklisted : GHookList

// Blacklisted : GIConv

// Blacklisted : GIOChannel

// Blacklisted : GIOFuncs

// Blacklisted : GKeyFile

// Blacklisted : GList

// Blacklisted : GMainContext

// Blacklisted : GMainLoop

// Blacklisted : GMappedFile

// Blacklisted : GMarkupParseContext

// Blacklisted : GMarkupParser

// Blacklisted : GMatchInfo

// Blacklisted : GMemVTable

// Blacklisted : GNode

// Blacklisted : GOptionContext

// Blacklisted : GOptionEntry

// Blacklisted : GOptionGroup

// Blacklisted : GPatternSpec

// Blacklisted : GPollFD

// Blacklisted : GPrivate

// Blacklisted : GPtrArray

// Blacklisted : GQueue

// Blacklisted : GRand

// SList is a wrapper around the C record GSList.
type SList struct {
	native *C.GSList
	// data : no type generator for gpointer, gpointer
	// next : record
}

func SListNewFromC(u unsafe.Pointer) *SList {
	c := (*C.GSList)(u)
	if c == nil {
		return nil
	}

	g := &SList{native: c}

	return g
}

func (recv *SList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SList with another SList, and returns true if they represent the same GObject.
func (recv *SList) Equals(other *SList) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_slist_alloc

// g_slist_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Blacklisted : g_slist_concat

// Blacklisted : g_slist_copy

// Blacklisted : g_slist_delete_link

// g_slist_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func
// Blacklisted : g_slist_free

// Blacklisted : g_slist_free_1

// g_slist_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_insert : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Blacklisted : g_slist_last

// Blacklisted : g_slist_length

// Blacklisted : g_slist_nth

// g_slist_nth_data : no return generator
// Blacklisted : g_slist_position

// g_slist_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// Blacklisted : g_slist_remove_link

// Blacklisted : g_slist_reverse

// g_slist_sort : unsupported parameter compare_func : no type generator for CompareFunc (GCompareFunc) for param compare_func
// g_slist_sort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func
// Blacklisted : GScanner

// Blacklisted : GScannerConfig

// Blacklisted : GSequence

// Blacklisted : GSequenceIter

// Blacklisted : GSource

// Blacklisted : GSourceCallbackFuncs

// Blacklisted : GSourceFuncs

// Blacklisted : GSourcePrivate

// Blacklisted : GStatBuf

// Blacklisted : GString

// Blacklisted : GStringChunk

// Blacklisted : GTestCase

// Blacklisted : GTestConfig

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Blacklisted : GTestSuite

// Blacklisted : GThread

// Blacklisted : GThreadPool

// Blacklisted : GTimeVal

// Blacklisted : GTimer

// Blacklisted : GTrashStack

// Blacklisted : GTree

// Blacklisted : GVariantBuilder

// Blacklisted : GVariantIter

// Blacklisted : GVariantType

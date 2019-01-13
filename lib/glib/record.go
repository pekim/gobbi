// This is a generated file - DO NOT EDIT

package glib

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	static GError* _g_error_new(GQuark domain, gint code, const gchar* format) {
		return g_error_new(domain, code, format);
    }
*/
/*

	static void _g_scanner_error(GScanner* scanner, const gchar* format) {
		return g_scanner_error(scanner, format);
    }
*/
/*

	static void _g_scanner_warn(GScanner* scanner, const gchar* format) {
		return g_scanner_warn(scanner, format);
    }
*/
/*

	static void _g_string_append_printf(GString* string, const gchar* format) {
		return g_string_append_printf(string, format);
    }
*/
/*

	static void _g_string_printf(GString* string, const gchar* format) {
		return g_string_printf(string, format);
    }
*/
import "C"

// Array is a wrapper around the C record GArray.
type Array struct {
	native *C.GArray
	Data   string
	Len    uint32
}

func ArrayNewFromC(u unsafe.Pointer) *Array {
	c := (*C.GArray)(u)
	if c == nil {
		return nil
	}

	g := &Array{
		Data:   C.GoString(c.data),
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
}

func (recv *Array) ToC() unsafe.Pointer {
	recv.native.data =
		C.CString(recv.Data)
	recv.native.len =
		(C.guint)(recv.Len)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Array with another Array, and returns true if they represent the same GObject.
func (recv *Array) Equals(other *Array) bool {
	return other.ToC() == recv.ToC()
}

// g_array_append_vals : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_free : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_insert_vals : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_new : no type generator for gpointer (gpointer) for array return
// g_array_prepend_vals : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_remove_index : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_remove_index_fast : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_set_size : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_sized_new : no type generator for gpointer (gpointer) for array return
// g_array_sort : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_sort_with_data : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// AsyncQueue is a wrapper around the C record GAsyncQueue.
type AsyncQueue struct {
	native *C.GAsyncQueue
}

func AsyncQueueNewFromC(u unsafe.Pointer) *AsyncQueue {
	c := (*C.GAsyncQueue)(u)
	if c == nil {
		return nil
	}

	g := &AsyncQueue{native: c}

	return g
}

func (recv *AsyncQueue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncQueue with another AsyncQueue, and returns true if they represent the same GObject.
func (recv *AsyncQueue) Equals(other *AsyncQueue) bool {
	return other.ToC() == recv.ToC()
}

// AsyncQueueNew is a wrapper around the C function g_async_queue_new.
func AsyncQueueNew() *AsyncQueue {
	retC := C.g_async_queue_new()
	retGo := AsyncQueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Length is a wrapper around the C function g_async_queue_length.
func (recv *AsyncQueue) Length() int32 {
	retC := C.g_async_queue_length((*C.GAsyncQueue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// LengthUnlocked is a wrapper around the C function g_async_queue_length_unlocked.
func (recv *AsyncQueue) LengthUnlocked() int32 {
	retC := C.g_async_queue_length_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Lock is a wrapper around the C function g_async_queue_lock.
func (recv *AsyncQueue) Lock() {
	C.g_async_queue_lock((*C.GAsyncQueue)(recv.native))

	return
}

// Unsupported : g_async_queue_pop : no return generator

// Unsupported : g_async_queue_pop_unlocked : no return generator

// Unsupported : g_async_queue_push : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_async_queue_push_unlocked : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Ref is a wrapper around the C function g_async_queue_ref.
func (recv *AsyncQueue) Ref() *AsyncQueue {
	retC := C.g_async_queue_ref((*C.GAsyncQueue)(recv.native))
	retGo := AsyncQueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefUnlocked is a wrapper around the C function g_async_queue_ref_unlocked.
func (recv *AsyncQueue) RefUnlocked() {
	C.g_async_queue_ref_unlocked((*C.GAsyncQueue)(recv.native))

	return
}

// Unsupported : g_async_queue_timed_pop : no return generator

// Unsupported : g_async_queue_timed_pop_unlocked : no return generator

// Unsupported : g_async_queue_timeout_pop : no return generator

// Unsupported : g_async_queue_timeout_pop_unlocked : no return generator

// Unsupported : g_async_queue_try_pop : no return generator

// Unsupported : g_async_queue_try_pop_unlocked : no return generator

// Unlock is a wrapper around the C function g_async_queue_unlock.
func (recv *AsyncQueue) Unlock() {
	C.g_async_queue_unlock((*C.GAsyncQueue)(recv.native))

	return
}

// Unref is a wrapper around the C function g_async_queue_unref.
func (recv *AsyncQueue) Unref() {
	C.g_async_queue_unref((*C.GAsyncQueue)(recv.native))

	return
}

// UnrefAndUnlock is a wrapper around the C function g_async_queue_unref_and_unlock.
func (recv *AsyncQueue) UnrefAndUnlock() {
	C.g_async_queue_unref_and_unlock((*C.GAsyncQueue)(recv.native))

	return
}

// BookmarkFile is a wrapper around the C record GBookmarkFile.
type BookmarkFile struct {
	native *C.GBookmarkFile
}

func BookmarkFileNewFromC(u unsafe.Pointer) *BookmarkFile {
	c := (*C.GBookmarkFile)(u)
	if c == nil {
		return nil
	}

	g := &BookmarkFile{native: c}

	return g
}

func (recv *BookmarkFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BookmarkFile with another BookmarkFile, and returns true if they represent the same GObject.
func (recv *BookmarkFile) Equals(other *BookmarkFile) bool {
	return other.ToC() == recv.ToC()
}

// BookmarkFileErrorQuark is a wrapper around the C function g_bookmark_file_error_quark.
func BookmarkFileErrorQuark() Quark {
	retC := C.g_bookmark_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Blacklisted : GByteArray

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native *C.GCond
	// Private : p
	// Private : i
}

func CondNewFromC(u unsafe.Pointer) *Cond {
	c := (*C.GCond)(u)
	if c == nil {
		return nil
	}

	g := &Cond{native: c}

	return g
}

func (recv *Cond) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Cond with another Cond, and returns true if they represent the same GObject.
func (recv *Cond) Equals(other *Cond) bool {
	return other.ToC() == recv.ToC()
}

// Broadcast is a wrapper around the C function g_cond_broadcast.
func (recv *Cond) Broadcast() {
	C.g_cond_broadcast((*C.GCond)(recv.native))

	return
}

// Signal is a wrapper around the C function g_cond_signal.
func (recv *Cond) Signal() {
	C.g_cond_signal((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// Data is a wrapper around the C record GData.
type Data struct {
	native *C.GData
}

func DataNewFromC(u unsafe.Pointer) *Data {
	c := (*C.GData)(u)
	if c == nil {
		return nil
	}

	g := &Data{native: c}

	return g
}

func (recv *Data) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Data with another Data, and returns true if they represent the same GObject.
func (recv *Data) Equals(other *Data) bool {
	return other.ToC() == recv.ToC()
}

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

// DateNew is a wrapper around the C function g_date_new.
func DateNew() *Date {
	retC := C.g_date_new()
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateNewDmy is a wrapper around the C function g_date_new_dmy.
func DateNewDmy(day DateDay, month DateMonth, year DateYear) *Date {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_new_dmy(c_day, c_month, c_year)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateNewJulian is a wrapper around the C function g_date_new_julian.
func DateNewJulian(julianDay uint32) *Date {
	c_julian_day := (C.guint32)(julianDay)

	retC := C.g_date_new_julian(c_julian_day)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateGetDaysInMonth is a wrapper around the C function g_date_get_days_in_month.
func DateGetDaysInMonth(month DateMonth, year DateYear) uint8 {
	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_days_in_month(c_month, c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateGetMondayWeeksInYear is a wrapper around the C function g_date_get_monday_weeks_in_year.
func DateGetMondayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_monday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateGetSundayWeeksInYear is a wrapper around the C function g_date_get_sunday_weeks_in_year.
func DateGetSundayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_sunday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateIsLeapYear is a wrapper around the C function g_date_is_leap_year.
func DateIsLeapYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_is_leap_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// DateStrftime is a wrapper around the C function g_date_strftime.
func DateStrftime(s string, slen uint64, format string, date *Date) uint64 {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	c_slen := (C.gsize)(slen)

	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	c_date := (*C.GDate)(C.NULL)
	if date != nil {
		c_date = (*C.GDate)(date.ToC())
	}

	retC := C.g_date_strftime(c_s, c_slen, c_format, c_date)
	retGo := (uint64)(retC)

	return retGo
}

// DateValidDay is a wrapper around the C function g_date_valid_day.
func DateValidDay(day DateDay) bool {
	c_day := (C.GDateDay)(day)

	retC := C.g_date_valid_day(c_day)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidDmy is a wrapper around the C function g_date_valid_dmy.
func DateValidDmy(day DateDay, month DateMonth, year DateYear) bool {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_dmy(c_day, c_month, c_year)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidJulian is a wrapper around the C function g_date_valid_julian.
func DateValidJulian(julianDate uint32) bool {
	c_julian_date := (C.guint32)(julianDate)

	retC := C.g_date_valid_julian(c_julian_date)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidMonth is a wrapper around the C function g_date_valid_month.
func DateValidMonth(month DateMonth) bool {
	c_month := (C.GDateMonth)(month)

	retC := C.g_date_valid_month(c_month)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidWeekday is a wrapper around the C function g_date_valid_weekday.
func DateValidWeekday(weekday DateWeekday) bool {
	c_weekday := (C.GDateWeekday)(weekday)

	retC := C.g_date_valid_weekday(c_weekday)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidYear is a wrapper around the C function g_date_valid_year.
func DateValidYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// AddDays is a wrapper around the C function g_date_add_days.
func (recv *Date) AddDays(nDays uint32) {
	c_n_days := (C.guint)(nDays)

	C.g_date_add_days((*C.GDate)(recv.native), c_n_days)

	return
}

// AddMonths is a wrapper around the C function g_date_add_months.
func (recv *Date) AddMonths(nMonths uint32) {
	c_n_months := (C.guint)(nMonths)

	C.g_date_add_months((*C.GDate)(recv.native), c_n_months)

	return
}

// AddYears is a wrapper around the C function g_date_add_years.
func (recv *Date) AddYears(nYears uint32) {
	c_n_years := (C.guint)(nYears)

	C.g_date_add_years((*C.GDate)(recv.native), c_n_years)

	return
}

// Clamp is a wrapper around the C function g_date_clamp.
func (recv *Date) Clamp(minDate *Date, maxDate *Date) {
	c_min_date := (*C.GDate)(C.NULL)
	if minDate != nil {
		c_min_date = (*C.GDate)(minDate.ToC())
	}

	c_max_date := (*C.GDate)(C.NULL)
	if maxDate != nil {
		c_max_date = (*C.GDate)(maxDate.ToC())
	}

	C.g_date_clamp((*C.GDate)(recv.native), c_min_date, c_max_date)

	return
}

// Clear is a wrapper around the C function g_date_clear.
func (recv *Date) Clear(nDates uint32) {
	c_n_dates := (C.guint)(nDates)

	C.g_date_clear((*C.GDate)(recv.native), c_n_dates)

	return
}

// Compare is a wrapper around the C function g_date_compare.
func (recv *Date) Compare(rhs *Date) int32 {
	c_rhs := (*C.GDate)(C.NULL)
	if rhs != nil {
		c_rhs = (*C.GDate)(rhs.ToC())
	}

	retC := C.g_date_compare((*C.GDate)(recv.native), c_rhs)
	retGo := (int32)(retC)

	return retGo
}

// DaysBetween is a wrapper around the C function g_date_days_between.
func (recv *Date) DaysBetween(date2 *Date) int32 {
	c_date2 := (*C.GDate)(C.NULL)
	if date2 != nil {
		c_date2 = (*C.GDate)(date2.ToC())
	}

	retC := C.g_date_days_between((*C.GDate)(recv.native), c_date2)
	retGo := (int32)(retC)

	return retGo
}

// Free is a wrapper around the C function g_date_free.
func (recv *Date) Free() {
	C.g_date_free((*C.GDate)(recv.native))

	return
}

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

// GetJulian is a wrapper around the C function g_date_get_julian.
func (recv *Date) GetJulian() uint32 {
	retC := C.g_date_get_julian((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMondayWeekOfYear is a wrapper around the C function g_date_get_monday_week_of_year.
func (recv *Date) GetMondayWeekOfYear() uint32 {
	retC := C.g_date_get_monday_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMonth is a wrapper around the C function g_date_get_month.
func (recv *Date) GetMonth() DateMonth {
	retC := C.g_date_get_month((*C.GDate)(recv.native))
	retGo := (DateMonth)(retC)

	return retGo
}

// GetSundayWeekOfYear is a wrapper around the C function g_date_get_sunday_week_of_year.
func (recv *Date) GetSundayWeekOfYear() uint32 {
	retC := C.g_date_get_sunday_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetWeekday is a wrapper around the C function g_date_get_weekday.
func (recv *Date) GetWeekday() DateWeekday {
	retC := C.g_date_get_weekday((*C.GDate)(recv.native))
	retGo := (DateWeekday)(retC)

	return retGo
}

// GetYear is a wrapper around the C function g_date_get_year.
func (recv *Date) GetYear() DateYear {
	retC := C.g_date_get_year((*C.GDate)(recv.native))
	retGo := (DateYear)(retC)

	return retGo
}

// IsFirstOfMonth is a wrapper around the C function g_date_is_first_of_month.
func (recv *Date) IsFirstOfMonth() bool {
	retC := C.g_date_is_first_of_month((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsLastOfMonth is a wrapper around the C function g_date_is_last_of_month.
func (recv *Date) IsLastOfMonth() bool {
	retC := C.g_date_is_last_of_month((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Order is a wrapper around the C function g_date_order.
func (recv *Date) Order(date2 *Date) {
	c_date2 := (*C.GDate)(C.NULL)
	if date2 != nil {
		c_date2 = (*C.GDate)(date2.ToC())
	}

	C.g_date_order((*C.GDate)(recv.native), c_date2)

	return
}

// SetDay is a wrapper around the C function g_date_set_day.
func (recv *Date) SetDay(day DateDay) {
	c_day := (C.GDateDay)(day)

	C.g_date_set_day((*C.GDate)(recv.native), c_day)

	return
}

// SetDmy is a wrapper around the C function g_date_set_dmy.
func (recv *Date) SetDmy(day DateDay, month DateMonth, y DateYear) {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_y := (C.GDateYear)(y)

	C.g_date_set_dmy((*C.GDate)(recv.native), c_day, c_month, c_y)

	return
}

// SetJulian is a wrapper around the C function g_date_set_julian.
func (recv *Date) SetJulian(julianDate uint32) {
	c_julian_date := (C.guint32)(julianDate)

	C.g_date_set_julian((*C.GDate)(recv.native), c_julian_date)

	return
}

// SetMonth is a wrapper around the C function g_date_set_month.
func (recv *Date) SetMonth(month DateMonth) {
	c_month := (C.GDateMonth)(month)

	C.g_date_set_month((*C.GDate)(recv.native), c_month)

	return
}

// SetParse is a wrapper around the C function g_date_set_parse.
func (recv *Date) SetParse(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.g_date_set_parse((*C.GDate)(recv.native), c_str)

	return
}

// SetTime is a wrapper around the C function g_date_set_time.
func (recv *Date) SetTime(time Time) {
	c_time_ := (C.GTime)(time)

	C.g_date_set_time((*C.GDate)(recv.native), c_time_)

	return
}

// SetYear is a wrapper around the C function g_date_set_year.
func (recv *Date) SetYear(year DateYear) {
	c_year := (C.GDateYear)(year)

	C.g_date_set_year((*C.GDate)(recv.native), c_year)

	return
}

// SubtractDays is a wrapper around the C function g_date_subtract_days.
func (recv *Date) SubtractDays(nDays uint32) {
	c_n_days := (C.guint)(nDays)

	C.g_date_subtract_days((*C.GDate)(recv.native), c_n_days)

	return
}

// SubtractMonths is a wrapper around the C function g_date_subtract_months.
func (recv *Date) SubtractMonths(nMonths uint32) {
	c_n_months := (C.guint)(nMonths)

	C.g_date_subtract_months((*C.GDate)(recv.native), c_n_months)

	return
}

// SubtractYears is a wrapper around the C function g_date_subtract_years.
func (recv *Date) SubtractYears(nYears uint32) {
	c_n_years := (C.guint)(nYears)

	C.g_date_subtract_years((*C.GDate)(recv.native), c_n_years)

	return
}

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer (tm*) for param tm

// Valid is a wrapper around the C function g_date_valid.
func (recv *Date) Valid() bool {
	retC := C.g_date_valid((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// DebugKey is a wrapper around the C record GDebugKey.
type DebugKey struct {
	native *C.GDebugKey
	Key    string
	Value  uint32
}

func DebugKeyNewFromC(u unsafe.Pointer) *DebugKey {
	c := (*C.GDebugKey)(u)
	if c == nil {
		return nil
	}

	g := &DebugKey{
		Key:    C.GoString(c.key),
		Value:  (uint32)(c.value),
		native: c,
	}

	return g
}

func (recv *DebugKey) ToC() unsafe.Pointer {
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		(C.guint)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DebugKey with another DebugKey, and returns true if they represent the same GObject.
func (recv *DebugKey) Equals(other *DebugKey) bool {
	return other.ToC() == recv.ToC()
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native *C.GDir
}

func DirNewFromC(u unsafe.Pointer) *Dir {
	c := (*C.GDir)(u)
	if c == nil {
		return nil
	}

	g := &Dir{native: c}

	return g
}

func (recv *Dir) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Dir with another Dir, and returns true if they represent the same GObject.
func (recv *Dir) Equals(other *Dir) bool {
	return other.ToC() == recv.ToC()
}

// DirOpen is a wrapper around the C function g_dir_open.
func DirOpen(path string, flags uint32) (*Dir, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_flags := (C.guint)(flags)

	var cThrowableError *C.GError

	retC := C.g_dir_open(c_path, c_flags, &cThrowableError)
	retGo := DirNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Close is a wrapper around the C function g_dir_close.
func (recv *Dir) Close() {
	C.g_dir_close((*C.GDir)(recv.native))

	return
}

// ReadName is a wrapper around the C function g_dir_read_name.
func (recv *Dir) ReadName() string {
	retC := C.g_dir_read_name((*C.GDir)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Rewind is a wrapper around the C function g_dir_rewind.
func (recv *Dir) Rewind() {
	C.g_dir_rewind((*C.GDir)(recv.native))

	return
}

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

// ErrorNew is a wrapper around the C function g_error_new.
func ErrorNew(domain Quark, code int32, format string, args ...interface{}) *Error {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_error_new(c_domain, c_code, c_format)
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ErrorNewLiteral is a wrapper around the C function g_error_new_literal.
func ErrorNewLiteral(domain Quark, code int32, message string) *Error {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	retC := C.g_error_new_literal(c_domain, c_code, c_message)
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_error_copy.
func (recv *Error) Copy() *Error {
	retC := C.g_error_copy((*C.GError)(recv.native))
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_error_free.
func (recv *Error) Free() {
	C.g_error_free((*C.GError)(recv.native))

	return
}

// Matches is a wrapper around the C function g_error_matches.
func (recv *Error) Matches(domain Quark, code int32) bool {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	retC := C.g_error_matches((*C.GError)(recv.native), c_domain, c_code)
	retGo := retC == C.TRUE

	return retGo
}

// HashTable is a wrapper around the C record GHashTable.
type HashTable struct {
	native *C.GHashTable
}

func HashTableNewFromC(u unsafe.Pointer) *HashTable {
	c := (*C.GHashTable)(u)
	if c == nil {
		return nil
	}

	g := &HashTable{native: c}

	return g
}

func (recv *HashTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HashTable with another HashTable, and returns true if they represent the same GObject.
func (recv *HashTable) Equals(other *HashTable) bool {
	return other.ToC() == recv.ToC()
}

// HashTableDestroy is a wrapper around the C function g_hash_table_destroy.
func HashTableDestroy(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_destroy(c_hash_table)

	return
}

// g_hash_table_foreach : unsupported parameter func : no type generator for HFunc (GHFunc) for param func
// g_hash_table_foreach_remove : unsupported parameter func : no type generator for HRFunc (GHRFunc) for param func
// g_hash_table_foreach_steal : unsupported parameter func : no type generator for HRFunc (GHRFunc) for param func
// g_hash_table_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key
// g_hash_table_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// g_hash_table_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key
// g_hash_table_new : unsupported parameter hash_func : no type generator for HashFunc (GHashFunc) for param hash_func
// g_hash_table_new_full : unsupported parameter hash_func : no type generator for HashFunc (GHashFunc) for param hash_func
// g_hash_table_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// g_hash_table_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key
// HashTableSize is a wrapper around the C function g_hash_table_size.
func HashTableSize(hashTable *HashTable) uint32 {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_size(c_hash_table)
	retGo := (uint32)(retC)

	return retGo
}

// g_hash_table_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// HashTableIter is a wrapper around the C record GHashTableIter.
type HashTableIter struct {
	native *C.GHashTableIter
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
}

func HashTableIterNewFromC(u unsafe.Pointer) *HashTableIter {
	c := (*C.GHashTableIter)(u)
	if c == nil {
		return nil
	}

	g := &HashTableIter{native: c}

	return g
}

func (recv *HashTableIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HashTableIter with another HashTableIter, and returns true if they represent the same GObject.
func (recv *HashTableIter) Equals(other *HashTableIter) bool {
	return other.ToC() == recv.ToC()
}

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native *C.GHook
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
	RefCount uint32
	HookId   uint64
	Flags    uint32
	// _func : no type generator for gpointer, gpointer
	// destroy : no type generator for DestroyNotify, GDestroyNotify
}

func HookNewFromC(u unsafe.Pointer) *Hook {
	c := (*C.GHook)(u)
	if c == nil {
		return nil
	}

	g := &Hook{
		Flags:    (uint32)(c.flags),
		HookId:   (uint64)(c.hook_id),
		RefCount: (uint32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *Hook) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.guint)(recv.RefCount)
	recv.native.hook_id =
		(C.gulong)(recv.HookId)
	recv.native.flags =
		(C.guint)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Hook with another Hook, and returns true if they represent the same GObject.
func (recv *Hook) Equals(other *Hook) bool {
	return other.ToC() == recv.ToC()
}

// HookAlloc is a wrapper around the C function g_hook_alloc.
func HookAlloc(hookList *HookList) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	retC := C.g_hook_alloc(c_hook_list)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookDestroy is a wrapper around the C function g_hook_destroy.
func HookDestroy(hookList *HookList, hookId uint64) bool {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook_id := (C.gulong)(hookId)

	retC := C.g_hook_destroy(c_hook_list, c_hook_id)
	retGo := retC == C.TRUE

	return retGo
}

// HookDestroyLink is a wrapper around the C function g_hook_destroy_link.
func HookDestroyLink(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_destroy_link(c_hook_list, c_hook)

	return
}

// g_hook_find : unsupported parameter func : no type generator for HookFindFunc (GHookFindFunc) for param func
// g_hook_find_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_hook_find_func : unsupported parameter func : no type generator for gpointer (gpointer) for param func
// g_hook_find_func_data : unsupported parameter func : no type generator for gpointer (gpointer) for param func
// HookFirstValid is a wrapper around the C function g_hook_first_valid.
func HookFirstValid(hookList *HookList, mayBeInCall bool) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_may_be_in_call :=
		boolToGboolean(mayBeInCall)

	retC := C.g_hook_first_valid(c_hook_list, c_may_be_in_call)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookFree is a wrapper around the C function g_hook_free.
func HookFree(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_free(c_hook_list, c_hook)

	return
}

// HookGet is a wrapper around the C function g_hook_get.
func HookGet(hookList *HookList, hookId uint64) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook_id := (C.gulong)(hookId)

	retC := C.g_hook_get(c_hook_list, c_hook_id)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookInsertBefore is a wrapper around the C function g_hook_insert_before.
func HookInsertBefore(hookList *HookList, sibling *Hook, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_sibling := (*C.GHook)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GHook)(sibling.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_insert_before(c_hook_list, c_sibling, c_hook)

	return
}

// g_hook_insert_sorted : unsupported parameter func : no type generator for HookCompareFunc (GHookCompareFunc) for param func
// HookNextValid is a wrapper around the C function g_hook_next_valid.
func HookNextValid(hookList *HookList, hook *Hook, mayBeInCall bool) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	c_may_be_in_call :=
		boolToGboolean(mayBeInCall)

	retC := C.g_hook_next_valid(c_hook_list, c_hook, c_may_be_in_call)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookPrepend is a wrapper around the C function g_hook_prepend.
func HookPrepend(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_prepend(c_hook_list, c_hook)

	return
}

// HookRef is a wrapper around the C function g_hook_ref.
func HookRef(hookList *HookList, hook *Hook) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	retC := C.g_hook_ref(c_hook_list, c_hook)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookUnref is a wrapper around the C function g_hook_unref.
func HookUnref(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_unref(c_hook_list, c_hook)

	return
}

// CompareIds is a wrapper around the C function g_hook_compare_ids.
func (recv *Hook) CompareIds(sibling *Hook) int32 {
	c_sibling := (*C.GHook)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GHook)(sibling.ToC())
	}

	retC := C.g_hook_compare_ids((*C.GHook)(recv.native), c_sibling)
	retGo := (int32)(retC)

	return retGo
}

// HookList is a wrapper around the C record GHookList.
type HookList struct {
	native *C.GHookList
	SeqId  uint64
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	// hooks : record
	// dummy3 : no type generator for gpointer, gpointer
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func HookListNewFromC(u unsafe.Pointer) *HookList {
	c := (*C.GHookList)(u)
	if c == nil {
		return nil
	}

	g := &HookList{
		SeqId:  (uint64)(c.seq_id),
		native: c,
	}

	return g
}

func (recv *HookList) ToC() unsafe.Pointer {
	recv.native.seq_id =
		(C.gulong)(recv.SeqId)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HookList with another HookList, and returns true if they represent the same GObject.
func (recv *HookList) Equals(other *HookList) bool {
	return other.ToC() == recv.ToC()
}

// Clear is a wrapper around the C function g_hook_list_clear.
func (recv *HookList) Clear() {
	C.g_hook_list_clear((*C.GHookList)(recv.native))

	return
}

// Init is a wrapper around the C function g_hook_list_init.
func (recv *HookList) Init(hookSize uint32) {
	c_hook_size := (C.guint)(hookSize)

	C.g_hook_list_init((*C.GHookList)(recv.native), c_hook_size)

	return
}

// Invoke is a wrapper around the C function g_hook_list_invoke.
func (recv *HookList) Invoke(mayRecurse bool) {
	c_may_recurse :=
		boolToGboolean(mayRecurse)

	C.g_hook_list_invoke((*C.GHookList)(recv.native), c_may_recurse)

	return
}

// InvokeCheck is a wrapper around the C function g_hook_list_invoke_check.
func (recv *HookList) InvokeCheck(mayRecurse bool) {
	c_may_recurse :=
		boolToGboolean(mayRecurse)

	C.g_hook_list_invoke_check((*C.GHookList)(recv.native), c_may_recurse)

	return
}

// Unsupported : g_hook_list_marshal : unsupported parameter marshaller : no type generator for HookMarshaller (GHookMarshaller) for param marshaller

// Unsupported : g_hook_list_marshal_check : unsupported parameter marshaller : no type generator for HookCheckMarshaller (GHookCheckMarshaller) for param marshaller

// Blacklisted : GIConv

// IOChannel is a wrapper around the C record GIOChannel.
type IOChannel struct {
	native *C.GIOChannel
	// Private : ref_count
	// Private : funcs
	// Private : encoding
	// Private : read_cd
	// Private : write_cd
	// Private : line_term
	// Private : line_term_len
	// Private : buf_size
	// Private : read_buf
	// Private : encoded_read_buf
	// Private : write_buf
	// Private : partial_write_buf
	// Private : use_buffer
	// Private : do_encode
	// Private : close_on_unref
	// Private : is_readable
	// Private : is_writeable
	// Private : is_seekable
	// Private : reserved1
	// Private : reserved2
}

func IOChannelNewFromC(u unsafe.Pointer) *IOChannel {
	c := (*C.GIOChannel)(u)
	if c == nil {
		return nil
	}

	g := &IOChannel{native: c}

	return g
}

func (recv *IOChannel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOChannel with another IOChannel, and returns true if they represent the same GObject.
func (recv *IOChannel) Equals(other *IOChannel) bool {
	return other.ToC() == recv.ToC()
}

// IOChannelNewFile is a wrapper around the C function g_io_channel_new_file.
func IOChannelNewFile(filename string, mode string) (*IOChannel, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_mode := C.CString(mode)
	defer C.free(unsafe.Pointer(c_mode))

	var cThrowableError *C.GError

	retC := C.g_io_channel_new_file(c_filename, c_mode, &cThrowableError)
	retGo := IOChannelNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IOChannelUnixNew is a wrapper around the C function g_io_channel_unix_new.
func IOChannelUnixNew(fd int32) *IOChannel {
	c_fd := (C.int)(fd)

	retC := C.g_io_channel_unix_new(c_fd)
	retGo := IOChannelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IOChannelErrorFromErrno is a wrapper around the C function g_io_channel_error_from_errno.
func IOChannelErrorFromErrno(en int32) IOChannelError {
	c_en := (C.gint)(en)

	retC := C.g_io_channel_error_from_errno(c_en)
	retGo := (IOChannelError)(retC)

	return retGo
}

// IOChannelErrorQuark is a wrapper around the C function g_io_channel_error_quark.
func IOChannelErrorQuark() Quark {
	retC := C.g_io_channel_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Close is a wrapper around the C function g_io_channel_close.
func (recv *IOChannel) Close() {
	C.g_io_channel_close((*C.GIOChannel)(recv.native))

	return
}

// Flush is a wrapper around the C function g_io_channel_flush.
func (recv *IOChannel) Flush() (IOStatus, error) {
	var cThrowableError *C.GError

	retC := C.g_io_channel_flush((*C.GIOChannel)(recv.native), &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetBufferCondition is a wrapper around the C function g_io_channel_get_buffer_condition.
func (recv *IOChannel) GetBufferCondition() IOCondition {
	retC := C.g_io_channel_get_buffer_condition((*C.GIOChannel)(recv.native))
	retGo := (IOCondition)(retC)

	return retGo
}

// GetBufferSize is a wrapper around the C function g_io_channel_get_buffer_size.
func (recv *IOChannel) GetBufferSize() uint64 {
	retC := C.g_io_channel_get_buffer_size((*C.GIOChannel)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetBuffered is a wrapper around the C function g_io_channel_get_buffered.
func (recv *IOChannel) GetBuffered() bool {
	retC := C.g_io_channel_get_buffered((*C.GIOChannel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCloseOnUnref is a wrapper around the C function g_io_channel_get_close_on_unref.
func (recv *IOChannel) GetCloseOnUnref() bool {
	retC := C.g_io_channel_get_close_on_unref((*C.GIOChannel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEncoding is a wrapper around the C function g_io_channel_get_encoding.
func (recv *IOChannel) GetEncoding() string {
	retC := C.g_io_channel_get_encoding((*C.GIOChannel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_io_channel_get_flags.
func (recv *IOChannel) GetFlags() IOFlags {
	retC := C.g_io_channel_get_flags((*C.GIOChannel)(recv.native))
	retGo := (IOFlags)(retC)

	return retGo
}

// GetLineTerm is a wrapper around the C function g_io_channel_get_line_term.
func (recv *IOChannel) GetLineTerm(length int32) string {
	c_length := (C.gint)(length)

	retC := C.g_io_channel_get_line_term((*C.GIOChannel)(recv.native), &c_length)
	retGo := C.GoString(retC)

	return retGo
}

// Init is a wrapper around the C function g_io_channel_init.
func (recv *IOChannel) Init() {
	C.g_io_channel_init((*C.GIOChannel)(recv.native))

	return
}

// Read is a wrapper around the C function g_io_channel_read.
func (recv *IOChannel) Read(buf string, count uint64, bytesRead uint64) IOError {
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))

	c_count := (C.gsize)(count)

	c_bytes_read := (C.gsize)(bytesRead)

	retC := C.g_io_channel_read((*C.GIOChannel)(recv.native), c_buf, c_count, &c_bytes_read)
	retGo := (IOError)(retC)

	return retGo
}

// Unsupported : g_io_channel_read_chars : unsupported parameter buf : output array param buf

// ReadLine is a wrapper around the C function g_io_channel_read_line.
func (recv *IOChannel) ReadLine() (IOStatus, string, uint64, uint64, error) {
	var c_str_return *C.gchar

	var c_length C.gsize

	var c_terminator_pos C.gsize

	var cThrowableError *C.GError

	retC := C.g_io_channel_read_line((*C.GIOChannel)(recv.native), &c_str_return, &c_length, &c_terminator_pos, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	strReturn := C.GoString(c_str_return)
	defer C.free(unsafe.Pointer(c_str_return))

	length := (uint64)(c_length)

	terminatorPos := (uint64)(c_terminator_pos)

	return retGo, strReturn, length, terminatorPos, goError
}

// ReadLineString is a wrapper around the C function g_io_channel_read_line_string.
func (recv *IOChannel) ReadLineString(buffer *String, terminatorPos uint64) (IOStatus, error) {
	c_buffer := (*C.GString)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GString)(buffer.ToC())
	}

	c_terminator_pos := (C.gsize)(terminatorPos)

	var cThrowableError *C.GError

	retC := C.g_io_channel_read_line_string((*C.GIOChannel)(recv.native), c_buffer, &c_terminator_pos, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_io_channel_read_to_end : unsupported parameter str_return : output array param str_return

// ReadUnichar is a wrapper around the C function g_io_channel_read_unichar.
func (recv *IOChannel) ReadUnichar() (IOStatus, rune, error) {
	var c_thechar C.gunichar

	var cThrowableError *C.GError

	retC := C.g_io_channel_read_unichar((*C.GIOChannel)(recv.native), &c_thechar, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	thechar := (rune)(c_thechar)

	return retGo, thechar, goError
}

// Ref is a wrapper around the C function g_io_channel_ref.
func (recv *IOChannel) Ref() *IOChannel {
	retC := C.g_io_channel_ref((*C.GIOChannel)(recv.native))
	retGo := IOChannelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Seek is a wrapper around the C function g_io_channel_seek.
func (recv *IOChannel) Seek(offset int64, type_ SeekType) IOError {
	c_offset := (C.gint64)(offset)

	c_type := (C.GSeekType)(type_)

	retC := C.g_io_channel_seek((*C.GIOChannel)(recv.native), c_offset, c_type)
	retGo := (IOError)(retC)

	return retGo
}

// SeekPosition is a wrapper around the C function g_io_channel_seek_position.
func (recv *IOChannel) SeekPosition(offset int64, type_ SeekType) (IOStatus, error) {
	c_offset := (C.gint64)(offset)

	c_type := (C.GSeekType)(type_)

	var cThrowableError *C.GError

	retC := C.g_io_channel_seek_position((*C.GIOChannel)(recv.native), c_offset, c_type, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetBufferSize is a wrapper around the C function g_io_channel_set_buffer_size.
func (recv *IOChannel) SetBufferSize(size uint64) {
	c_size := (C.gsize)(size)

	C.g_io_channel_set_buffer_size((*C.GIOChannel)(recv.native), c_size)

	return
}

// SetBuffered is a wrapper around the C function g_io_channel_set_buffered.
func (recv *IOChannel) SetBuffered(buffered bool) {
	c_buffered :=
		boolToGboolean(buffered)

	C.g_io_channel_set_buffered((*C.GIOChannel)(recv.native), c_buffered)

	return
}

// SetCloseOnUnref is a wrapper around the C function g_io_channel_set_close_on_unref.
func (recv *IOChannel) SetCloseOnUnref(doClose bool) {
	c_do_close :=
		boolToGboolean(doClose)

	C.g_io_channel_set_close_on_unref((*C.GIOChannel)(recv.native), c_do_close)

	return
}

// SetEncoding is a wrapper around the C function g_io_channel_set_encoding.
func (recv *IOChannel) SetEncoding(encoding string) (IOStatus, error) {
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	var cThrowableError *C.GError

	retC := C.g_io_channel_set_encoding((*C.GIOChannel)(recv.native), c_encoding, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetFlags is a wrapper around the C function g_io_channel_set_flags.
func (recv *IOChannel) SetFlags(flags IOFlags) (IOStatus, error) {
	c_flags := (C.GIOFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_io_channel_set_flags((*C.GIOChannel)(recv.native), c_flags, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetLineTerm is a wrapper around the C function g_io_channel_set_line_term.
func (recv *IOChannel) SetLineTerm(lineTerm string, length int32) {
	c_line_term := C.CString(lineTerm)
	defer C.free(unsafe.Pointer(c_line_term))

	c_length := (C.gint)(length)

	C.g_io_channel_set_line_term((*C.GIOChannel)(recv.native), c_line_term, c_length)

	return
}

// Shutdown is a wrapper around the C function g_io_channel_shutdown.
func (recv *IOChannel) Shutdown(flush bool) (IOStatus, error) {
	c_flush :=
		boolToGboolean(flush)

	var cThrowableError *C.GError

	retC := C.g_io_channel_shutdown((*C.GIOChannel)(recv.native), c_flush, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnixGetFd is a wrapper around the C function g_io_channel_unix_get_fd.
func (recv *IOChannel) UnixGetFd() int32 {
	retC := C.g_io_channel_unix_get_fd((*C.GIOChannel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unref is a wrapper around the C function g_io_channel_unref.
func (recv *IOChannel) Unref() {
	C.g_io_channel_unref((*C.GIOChannel)(recv.native))

	return
}

// Write is a wrapper around the C function g_io_channel_write.
func (recv *IOChannel) Write(buf string, count uint64, bytesWritten uint64) IOError {
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))

	c_count := (C.gsize)(count)

	c_bytes_written := (C.gsize)(bytesWritten)

	retC := C.g_io_channel_write((*C.GIOChannel)(recv.native), c_buf, c_count, &c_bytes_written)
	retGo := (IOError)(retC)

	return retGo
}

// WriteChars is a wrapper around the C function g_io_channel_write_chars.
func (recv *IOChannel) WriteChars(buf []uint8, count int64) (IOStatus, uint64, error) {
	c_buf_array := make([]C.guint8, len(buf), len(buf))
	c_buf := &c_buf_array[0]

	c_count := (C.gssize)(count)

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_io_channel_write_chars((*C.GIOChannel)(recv.native), c_buf, c_count, &c_bytes_written, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goError
}

// WriteUnichar is a wrapper around the C function g_io_channel_write_unichar.
func (recv *IOChannel) WriteUnichar(thechar rune) (IOStatus, error) {
	c_thechar := (C.gunichar)(thechar)

	var cThrowableError *C.GError

	retC := C.g_io_channel_write_unichar((*C.GIOChannel)(recv.native), c_thechar, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IOFuncs is a wrapper around the C record GIOFuncs.
type IOFuncs struct {
	native *C.GIOFuncs
	// no type for io_read
	// no type for io_write
	// no type for io_seek
	// no type for io_close
	// no type for io_create_watch
	// no type for io_free
	// no type for io_set_flags
	// no type for io_get_flags
}

func IOFuncsNewFromC(u unsafe.Pointer) *IOFuncs {
	c := (*C.GIOFuncs)(u)
	if c == nil {
		return nil
	}

	g := &IOFuncs{native: c}

	return g
}

func (recv *IOFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOFuncs with another IOFuncs, and returns true if they represent the same GObject.
func (recv *IOFuncs) Equals(other *IOFuncs) bool {
	return other.ToC() == recv.ToC()
}

// KeyFile is a wrapper around the C record GKeyFile.
type KeyFile struct {
	native *C.GKeyFile
}

func KeyFileNewFromC(u unsafe.Pointer) *KeyFile {
	c := (*C.GKeyFile)(u)
	if c == nil {
		return nil
	}

	g := &KeyFile{native: c}

	return g
}

func (recv *KeyFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this KeyFile with another KeyFile, and returns true if they represent the same GObject.
func (recv *KeyFile) Equals(other *KeyFile) bool {
	return other.ToC() == recv.ToC()
}

// KeyFileErrorQuark is a wrapper around the C function g_key_file_error_quark.
func KeyFileErrorQuark() Quark {
	retC := C.g_key_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
}

func ListNewFromC(u unsafe.Pointer) *List {
	c := (*C.GList)(u)
	if c == nil {
		return nil
	}

	g := &List{native: c}

	return g
}

func (recv *List) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this List with another List, and returns true if they represent the same GObject.
func (recv *List) Equals(other *List) bool {
	return other.ToC() == recv.ToC()
}

// ListAlloc is a wrapper around the C function g_list_alloc.
func ListAlloc() *List {
	retC := C.g_list_alloc()
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// ListConcat is a wrapper around the C function g_list_concat.
func ListConcat(list1 *List, list2 *List) *List {
	c_list1 := (*C.GList)(C.NULL)
	if list1 != nil {
		c_list1 = (*C.GList)(list1.ToC())
	}

	c_list2 := (*C.GList)(C.NULL)
	if list2 != nil {
		c_list2 = (*C.GList)(list2.ToC())
	}

	retC := C.g_list_concat(c_list1, c_list2)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListCopy is a wrapper around the C function g_list_copy.
func ListCopy(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_copy(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListDeleteLink is a wrapper around the C function g_list_delete_link.
func ListDeleteLink(list *List, link *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	retC := C.g_list_delete_link(c_list, c_link_)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// ListFirst is a wrapper around the C function g_list_first.
func ListFirst(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_first(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func
// ListFree is a wrapper around the C function g_list_free.
func ListFree(list *List) {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	C.g_list_free(c_list)

	return
}

// ListFree1 is a wrapper around the C function g_list_free_1.
func ListFree1(list *List) {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	C.g_list_free_1(c_list)

	return
}

// g_list_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_insert : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// ListLast is a wrapper around the C function g_list_last.
func ListLast(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_last(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListLength is a wrapper around the C function g_list_length.
func ListLength(list *List) uint32 {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_length(c_list)
	retGo := (uint32)(retC)

	return retGo
}

// ListNth is a wrapper around the C function g_list_nth.
func ListNth(list *List, n uint32) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_list_nth(c_list, c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_nth_data : no return generator
// ListNthPrev is a wrapper around the C function g_list_nth_prev.
func ListNthPrev(list *List, n uint32) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_list_nth_prev(c_list, c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListPosition is a wrapper around the C function g_list_position.
func ListPosition(list *List, llink *List) int32 {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_llink := (*C.GList)(C.NULL)
	if llink != nil {
		c_llink = (*C.GList)(llink.ToC())
	}

	retC := C.g_list_position(c_list, c_llink)
	retGo := (int32)(retC)

	return retGo
}

// g_list_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// ListRemoveLink is a wrapper around the C function g_list_remove_link.
func ListRemoveLink(list *List, llink *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_llink := (*C.GList)(C.NULL)
	if llink != nil {
		c_llink = (*C.GList)(llink.ToC())
	}

	retC := C.g_list_remove_link(c_list, c_llink)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListReverse is a wrapper around the C function g_list_reverse.
func ListReverse(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_reverse(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_sort : unsupported parameter compare_func : no type generator for CompareFunc (GCompareFunc) for param compare_func
// g_list_sort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func
// MainContext is a wrapper around the C record GMainContext.
type MainContext struct {
	native *C.GMainContext
}

func MainContextNewFromC(u unsafe.Pointer) *MainContext {
	c := (*C.GMainContext)(u)
	if c == nil {
		return nil
	}

	g := &MainContext{native: c}

	return g
}

func (recv *MainContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MainContext with another MainContext, and returns true if they represent the same GObject.
func (recv *MainContext) Equals(other *MainContext) bool {
	return other.ToC() == recv.ToC()
}

// MainContextNew is a wrapper around the C function g_main_context_new.
func MainContextNew() *MainContext {
	retC := C.g_main_context_new()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MainContextDefault is a wrapper around the C function g_main_context_default.
func MainContextDefault() *MainContext {
	retC := C.g_main_context_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Acquire is a wrapper around the C function g_main_context_acquire.
func (recv *MainContext) Acquire() bool {
	retC := C.g_main_context_acquire((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AddPoll is a wrapper around the C function g_main_context_add_poll.
func (recv *MainContext) AddPoll(fd *PollFD, priority int32) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	c_priority := (C.gint)(priority)

	C.g_main_context_add_poll((*C.GMainContext)(recv.native), c_fd, c_priority)

	return
}

// Unsupported : g_main_context_check : unsupported parameter fds :

// Dispatch is a wrapper around the C function g_main_context_dispatch.
func (recv *MainContext) Dispatch() {
	C.g_main_context_dispatch((*C.GMainContext)(recv.native))

	return
}

// Unsupported : g_main_context_find_source_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// FindSourceById is a wrapper around the C function g_main_context_find_source_by_id.
func (recv *MainContext) FindSourceById(sourceId uint32) *Source {
	c_source_id := (C.guint)(sourceId)

	retC := C.g_main_context_find_source_by_id((*C.GMainContext)(recv.native), c_source_id)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_context_find_source_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_main_context_get_poll_func : no return generator

// Iteration is a wrapper around the C function g_main_context_iteration.
func (recv *MainContext) Iteration(mayBlock bool) bool {
	c_may_block :=
		boolToGboolean(mayBlock)

	retC := C.g_main_context_iteration((*C.GMainContext)(recv.native), c_may_block)
	retGo := retC == C.TRUE

	return retGo
}

// Pending is a wrapper around the C function g_main_context_pending.
func (recv *MainContext) Pending() bool {
	retC := C.g_main_context_pending((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Prepare is a wrapper around the C function g_main_context_prepare.
func (recv *MainContext) Prepare(priority int32) bool {
	c_priority := (C.gint)(priority)

	retC := C.g_main_context_prepare((*C.GMainContext)(recv.native), &c_priority)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_main_context_query : unsupported parameter fds : output array param fds

// Ref is a wrapper around the C function g_main_context_ref.
func (recv *MainContext) Ref() *MainContext {
	retC := C.g_main_context_ref((*C.GMainContext)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Release is a wrapper around the C function g_main_context_release.
func (recv *MainContext) Release() {
	C.g_main_context_release((*C.GMainContext)(recv.native))

	return
}

// RemovePoll is a wrapper around the C function g_main_context_remove_poll.
func (recv *MainContext) RemovePoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_main_context_remove_poll((*C.GMainContext)(recv.native), c_fd)

	return
}

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc (GPollFunc) for param func

// Unref is a wrapper around the C function g_main_context_unref.
func (recv *MainContext) Unref() {
	C.g_main_context_unref((*C.GMainContext)(recv.native))

	return
}

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// Wakeup is a wrapper around the C function g_main_context_wakeup.
func (recv *MainContext) Wakeup() {
	C.g_main_context_wakeup((*C.GMainContext)(recv.native))

	return
}

// MainLoop is a wrapper around the C record GMainLoop.
type MainLoop struct {
	native *C.GMainLoop
}

func MainLoopNewFromC(u unsafe.Pointer) *MainLoop {
	c := (*C.GMainLoop)(u)
	if c == nil {
		return nil
	}

	g := &MainLoop{native: c}

	return g
}

func (recv *MainLoop) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MainLoop with another MainLoop, and returns true if they represent the same GObject.
func (recv *MainLoop) Equals(other *MainLoop) bool {
	return other.ToC() == recv.ToC()
}

// MainLoopNew is a wrapper around the C function g_main_loop_new.
func MainLoopNew(context *MainContext, isRunning bool) *MainLoop {
	c_context := (*C.GMainContext)(C.NULL)
	if context != nil {
		c_context = (*C.GMainContext)(context.ToC())
	}

	c_is_running :=
		boolToGboolean(isRunning)

	retC := C.g_main_loop_new(c_context, c_is_running)
	retGo := MainLoopNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContext is a wrapper around the C function g_main_loop_get_context.
func (recv *MainLoop) GetContext() *MainContext {
	retC := C.g_main_loop_get_context((*C.GMainLoop)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsRunning is a wrapper around the C function g_main_loop_is_running.
func (recv *MainLoop) IsRunning() bool {
	retC := C.g_main_loop_is_running((*C.GMainLoop)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Quit is a wrapper around the C function g_main_loop_quit.
func (recv *MainLoop) Quit() {
	C.g_main_loop_quit((*C.GMainLoop)(recv.native))

	return
}

// Ref is a wrapper around the C function g_main_loop_ref.
func (recv *MainLoop) Ref() *MainLoop {
	retC := C.g_main_loop_ref((*C.GMainLoop)(recv.native))
	retGo := MainLoopNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Run is a wrapper around the C function g_main_loop_run.
func (recv *MainLoop) Run() {
	C.g_main_loop_run((*C.GMainLoop)(recv.native))

	return
}

// Unref is a wrapper around the C function g_main_loop_unref.
func (recv *MainLoop) Unref() {
	C.g_main_loop_unref((*C.GMainLoop)(recv.native))

	return
}

// MappedFile is a wrapper around the C record GMappedFile.
type MappedFile struct {
	native *C.GMappedFile
}

func MappedFileNewFromC(u unsafe.Pointer) *MappedFile {
	c := (*C.GMappedFile)(u)
	if c == nil {
		return nil
	}

	g := &MappedFile{native: c}

	return g
}

func (recv *MappedFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MappedFile with another MappedFile, and returns true if they represent the same GObject.
func (recv *MappedFile) Equals(other *MappedFile) bool {
	return other.ToC() == recv.ToC()
}

// Unref is a wrapper around the C function g_mapped_file_unref.
func (recv *MappedFile) Unref() {
	C.g_mapped_file_unref((*C.GMappedFile)(recv.native))

	return
}

// MarkupParseContext is a wrapper around the C record GMarkupParseContext.
type MarkupParseContext struct {
	native *C.GMarkupParseContext
}

func MarkupParseContextNewFromC(u unsafe.Pointer) *MarkupParseContext {
	c := (*C.GMarkupParseContext)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParseContext{native: c}

	return g
}

func (recv *MarkupParseContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkupParseContext with another MarkupParseContext, and returns true if they represent the same GObject.
func (recv *MarkupParseContext) Equals(other *MarkupParseContext) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// EndParse is a wrapper around the C function g_markup_parse_context_end_parse.
func (recv *MarkupParseContext) EndParse() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_markup_parse_context_end_parse((*C.GMarkupParseContext)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Free is a wrapper around the C function g_markup_parse_context_free.
func (recv *MarkupParseContext) Free() {
	C.g_markup_parse_context_free((*C.GMarkupParseContext)(recv.native))

	return
}

// GetPosition is a wrapper around the C function g_markup_parse_context_get_position.
func (recv *MarkupParseContext) GetPosition(lineNumber int32, charNumber int32) {
	c_line_number := (C.gint)(lineNumber)

	c_char_number := (C.gint)(charNumber)

	C.g_markup_parse_context_get_position((*C.GMarkupParseContext)(recv.native), &c_line_number, &c_char_number)

	return
}

// Parse is a wrapper around the C function g_markup_parse_context_parse.
func (recv *MarkupParseContext) Parse(text string, textLen int64) (bool, error) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_text_len := (C.gssize)(textLen)

	var cThrowableError *C.GError

	retC := C.g_markup_parse_context_parse((*C.GMarkupParseContext)(recv.native), c_text, c_text_len, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MarkupParser is a wrapper around the C record GMarkupParser.
type MarkupParser struct {
	native *C.GMarkupParser
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

func MarkupParserNewFromC(u unsafe.Pointer) *MarkupParser {
	c := (*C.GMarkupParser)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParser{native: c}

	return g
}

func (recv *MarkupParser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkupParser with another MarkupParser, and returns true if they represent the same GObject.
func (recv *MarkupParser) Equals(other *MarkupParser) bool {
	return other.ToC() == recv.ToC()
}

// MatchInfo is a wrapper around the C record GMatchInfo.
type MatchInfo struct {
	native *C.GMatchInfo
}

func MatchInfoNewFromC(u unsafe.Pointer) *MatchInfo {
	c := (*C.GMatchInfo)(u)
	if c == nil {
		return nil
	}

	g := &MatchInfo{native: c}

	return g
}

func (recv *MatchInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MatchInfo with another MatchInfo, and returns true if they represent the same GObject.
func (recv *MatchInfo) Equals(other *MatchInfo) bool {
	return other.ToC() == recv.ToC()
}

// MemVTable is a wrapper around the C record GMemVTable.
type MemVTable struct {
	native *C.GMemVTable
	// no type for malloc
	// no type for realloc
	// no type for free
	// no type for calloc
	// no type for try_malloc
	// no type for try_realloc
}

func MemVTableNewFromC(u unsafe.Pointer) *MemVTable {
	c := (*C.GMemVTable)(u)
	if c == nil {
		return nil
	}

	g := &MemVTable{native: c}

	return g
}

func (recv *MemVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemVTable with another MemVTable, and returns true if they represent the same GObject.
func (recv *MemVTable) Equals(other *MemVTable) bool {
	return other.ToC() == recv.ToC()
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native *C.GNode
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
	// parent : record
	// children : record
}

func NodeNewFromC(u unsafe.Pointer) *Node {
	c := (*C.GNode)(u)
	if c == nil {
		return nil
	}

	g := &Node{native: c}

	return g
}

func (recv *Node) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Node with another Node, and returns true if they represent the same GObject.
func (recv *Node) Equals(other *Node) bool {
	return other.ToC() == recv.ToC()
}

// g_node_new : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Unsupported : g_node_child_index : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// ChildPosition is a wrapper around the C function g_node_child_position.
func (recv *Node) ChildPosition(child *Node) int32 {
	c_child := (*C.GNode)(C.NULL)
	if child != nil {
		c_child = (*C.GNode)(child.ToC())
	}

	retC := C.g_node_child_position((*C.GNode)(recv.native), c_child)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_node_children_foreach : unsupported parameter func : no type generator for NodeForeachFunc (GNodeForeachFunc) for param func

// Copy is a wrapper around the C function g_node_copy.
func (recv *Node) Copy() *Node {
	retC := C.g_node_copy((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Depth is a wrapper around the C function g_node_depth.
func (recv *Node) Depth() uint32 {
	retC := C.g_node_depth((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Destroy is a wrapper around the C function g_node_destroy.
func (recv *Node) Destroy() {
	C.g_node_destroy((*C.GNode)(recv.native))

	return
}

// Unsupported : g_node_find : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_node_find_child : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// FirstSibling is a wrapper around the C function g_node_first_sibling.
func (recv *Node) FirstSibling() *Node {
	retC := C.g_node_first_sibling((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRoot is a wrapper around the C function g_node_get_root.
func (recv *Node) GetRoot() *Node {
	retC := C.g_node_get_root((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function g_node_insert.
func (recv *Node) Insert(position int32, node *Node) *Node {
	c_position := (C.gint)(position)

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert((*C.GNode)(recv.native), c_position, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertAfter is a wrapper around the C function g_node_insert_after.
func (recv *Node) InsertAfter(sibling *Node, node *Node) *Node {
	c_sibling := (*C.GNode)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GNode)(sibling.ToC())
	}

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert_after((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertBefore is a wrapper around the C function g_node_insert_before.
func (recv *Node) InsertBefore(sibling *Node, node *Node) *Node {
	c_sibling := (*C.GNode)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GNode)(sibling.ToC())
	}

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert_before((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsAncestor is a wrapper around the C function g_node_is_ancestor.
func (recv *Node) IsAncestor(descendant *Node) bool {
	c_descendant := (*C.GNode)(C.NULL)
	if descendant != nil {
		c_descendant = (*C.GNode)(descendant.ToC())
	}

	retC := C.g_node_is_ancestor((*C.GNode)(recv.native), c_descendant)
	retGo := retC == C.TRUE

	return retGo
}

// LastChild is a wrapper around the C function g_node_last_child.
func (recv *Node) LastChild() *Node {
	retC := C.g_node_last_child((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LastSibling is a wrapper around the C function g_node_last_sibling.
func (recv *Node) LastSibling() *Node {
	retC := C.g_node_last_sibling((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MaxHeight is a wrapper around the C function g_node_max_height.
func (recv *Node) MaxHeight() uint32 {
	retC := C.g_node_max_height((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NChildren is a wrapper around the C function g_node_n_children.
func (recv *Node) NChildren() uint32 {
	retC := C.g_node_n_children((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NNodes is a wrapper around the C function g_node_n_nodes.
func (recv *Node) NNodes(flags TraverseFlags) uint32 {
	c_flags := (C.GTraverseFlags)(flags)

	retC := C.g_node_n_nodes((*C.GNode)(recv.native), c_flags)
	retGo := (uint32)(retC)

	return retGo
}

// NthChild is a wrapper around the C function g_node_nth_child.
func (recv *Node) NthChild(n uint32) *Node {
	c_n := (C.guint)(n)

	retC := C.g_node_nth_child((*C.GNode)(recv.native), c_n)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prepend is a wrapper around the C function g_node_prepend.
func (recv *Node) Prepend(node *Node) *Node {
	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_prepend((*C.GNode)(recv.native), c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ReverseChildren is a wrapper around the C function g_node_reverse_children.
func (recv *Node) ReverseChildren() {
	C.g_node_reverse_children((*C.GNode)(recv.native))

	return
}

// Unsupported : g_node_traverse : unsupported parameter func : no type generator for NodeTraverseFunc (GNodeTraverseFunc) for param func

// Unlink is a wrapper around the C function g_node_unlink.
func (recv *Node) Unlink() {
	C.g_node_unlink((*C.GNode)(recv.native))

	return
}

// OptionContext is a wrapper around the C record GOptionContext.
type OptionContext struct {
	native *C.GOptionContext
}

func OptionContextNewFromC(u unsafe.Pointer) *OptionContext {
	c := (*C.GOptionContext)(u)
	if c == nil {
		return nil
	}

	g := &OptionContext{native: c}

	return g
}

func (recv *OptionContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionContext with another OptionContext, and returns true if they represent the same GObject.
func (recv *OptionContext) Equals(other *OptionContext) bool {
	return other.ToC() == recv.ToC()
}

// OptionEntry is a wrapper around the C record GOptionEntry.
type OptionEntry struct {
	native    *C.GOptionEntry
	LongName  string
	ShortName rune
	Flags     int32
	Arg       OptionArg
	// arg_data : no type generator for gpointer, gpointer
	Description    string
	ArgDescription string
}

func OptionEntryNewFromC(u unsafe.Pointer) *OptionEntry {
	c := (*C.GOptionEntry)(u)
	if c == nil {
		return nil
	}

	g := &OptionEntry{
		Arg:            (OptionArg)(c.arg),
		ArgDescription: C.GoString(c.arg_description),
		Description:    C.GoString(c.description),
		Flags:          (int32)(c.flags),
		LongName:       C.GoString(c.long_name),
		ShortName:      (rune)(c.short_name),
		native:         c,
	}

	return g
}

func (recv *OptionEntry) ToC() unsafe.Pointer {
	recv.native.long_name =
		C.CString(recv.LongName)
	recv.native.short_name =
		(C.gchar)(recv.ShortName)
	recv.native.flags =
		(C.gint)(recv.Flags)
	recv.native.arg =
		(C.GOptionArg)(recv.Arg)
	recv.native.description =
		C.CString(recv.Description)
	recv.native.arg_description =
		C.CString(recv.ArgDescription)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionEntry with another OptionEntry, and returns true if they represent the same GObject.
func (recv *OptionEntry) Equals(other *OptionEntry) bool {
	return other.ToC() == recv.ToC()
}

// OptionGroup is a wrapper around the C record GOptionGroup.
type OptionGroup struct {
	native *C.GOptionGroup
}

func OptionGroupNewFromC(u unsafe.Pointer) *OptionGroup {
	c := (*C.GOptionGroup)(u)
	if c == nil {
		return nil
	}

	g := &OptionGroup{native: c}

	return g
}

func (recv *OptionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionGroup with another OptionGroup, and returns true if they represent the same GObject.
func (recv *OptionGroup) Equals(other *OptionGroup) bool {
	return other.ToC() == recv.ToC()
}

// PatternSpec is a wrapper around the C record GPatternSpec.
type PatternSpec struct {
	native *C.GPatternSpec
}

func PatternSpecNewFromC(u unsafe.Pointer) *PatternSpec {
	c := (*C.GPatternSpec)(u)
	if c == nil {
		return nil
	}

	g := &PatternSpec{native: c}

	return g
}

func (recv *PatternSpec) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PatternSpec with another PatternSpec, and returns true if they represent the same GObject.
func (recv *PatternSpec) Equals(other *PatternSpec) bool {
	return other.ToC() == recv.ToC()
}

// PatternSpecNew is a wrapper around the C function g_pattern_spec_new.
func PatternSpecNew(pattern string) *PatternSpec {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	retC := C.g_pattern_spec_new(c_pattern)
	retGo := PatternSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_pattern_spec_equal.
func (recv *PatternSpec) Equal(pspec2 *PatternSpec) bool {
	c_pspec2 := (*C.GPatternSpec)(C.NULL)
	if pspec2 != nil {
		c_pspec2 = (*C.GPatternSpec)(pspec2.ToC())
	}

	retC := C.g_pattern_spec_equal((*C.GPatternSpec)(recv.native), c_pspec2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function g_pattern_spec_free.
func (recv *PatternSpec) Free() {
	C.g_pattern_spec_free((*C.GPatternSpec)(recv.native))

	return
}

// PollFD is a wrapper around the C record GPollFD.
type PollFD struct {
	native  *C.GPollFD
	Fd      int32
	Events  uint32
	Revents uint32
}

func PollFDNewFromC(u unsafe.Pointer) *PollFD {
	c := (*C.GPollFD)(u)
	if c == nil {
		return nil
	}

	g := &PollFD{
		Events:  (uint32)(c.events),
		Fd:      (int32)(c.fd),
		Revents: (uint32)(c.revents),
		native:  c,
	}

	return g
}

func (recv *PollFD) ToC() unsafe.Pointer {
	recv.native.fd =
		(C.gint)(recv.Fd)
	recv.native.events =
		(C.gushort)(recv.Events)
	recv.native.revents =
		(C.gushort)(recv.Revents)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollFD with another PollFD, and returns true if they represent the same GObject.
func (recv *PollFD) Equals(other *PollFD) bool {
	return other.ToC() == recv.ToC()
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native *C.GPrivate
	// Private : p
	// Private : notify
	// Private : future
}

func PrivateNewFromC(u unsafe.Pointer) *Private {
	c := (*C.GPrivate)(u)
	if c == nil {
		return nil
	}

	g := &Private{native: c}

	return g
}

func (recv *Private) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Private with another Private, and returns true if they represent the same GObject.
func (recv *Private) Equals(other *Private) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_private_get : no return generator

// Unsupported : g_private_set : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Blacklisted : GPtrArray

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native *C.GQueue
	// head : record
	// tail : record
	Length uint32
}

func QueueNewFromC(u unsafe.Pointer) *Queue {
	c := (*C.GQueue)(u)
	if c == nil {
		return nil
	}

	g := &Queue{
		Length: (uint32)(c.length),
		native: c,
	}

	return g
}

func (recv *Queue) ToC() unsafe.Pointer {
	recv.native.length =
		(C.guint)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Queue with another Queue, and returns true if they represent the same GObject.
func (recv *Queue) Equals(other *Queue) bool {
	return other.ToC() == recv.ToC()
}

// QueueNew is a wrapper around the C function g_queue_new.
func QueueNew() *Queue {
	retC := C.g_queue_new()
	retGo := QueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_queue_free.
func (recv *Queue) Free() {
	C.g_queue_free((*C.GQueue)(recv.native))

	return
}

// IsEmpty is a wrapper around the C function g_queue_is_empty.
func (recv *Queue) IsEmpty() bool {
	retC := C.g_queue_is_empty((*C.GQueue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_queue_peek_head : no return generator

// Unsupported : g_queue_peek_tail : no return generator

// Unsupported : g_queue_pop_head : no return generator

// PopHeadLink is a wrapper around the C function g_queue_pop_head_link.
func (recv *Queue) PopHeadLink() *List {
	retC := C.g_queue_pop_head_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_pop_tail : no return generator

// PopTailLink is a wrapper around the C function g_queue_pop_tail_link.
func (recv *Queue) PopTailLink() *List {
	retC := C.g_queue_pop_tail_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_push_head : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PushHeadLink is a wrapper around the C function g_queue_push_head_link.
func (recv *Queue) PushHeadLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_head_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Unsupported : g_queue_push_tail : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PushTailLink is a wrapper around the C function g_queue_push_tail_link.
func (recv *Queue) PushTailLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_tail_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native *C.GRand
}

func RandNewFromC(u unsafe.Pointer) *Rand {
	c := (*C.GRand)(u)
	if c == nil {
		return nil
	}

	g := &Rand{native: c}

	return g
}

func (recv *Rand) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Rand with another Rand, and returns true if they represent the same GObject.
func (recv *Rand) Equals(other *Rand) bool {
	return other.ToC() == recv.ToC()
}

// RandNew is a wrapper around the C function g_rand_new.
func RandNew() *Rand {
	retC := C.g_rand_new()
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RandNewWithSeed is a wrapper around the C function g_rand_new_with_seed.
func RandNewWithSeed(seed uint32) *Rand {
	c_seed := (C.guint32)(seed)

	retC := C.g_rand_new_with_seed(c_seed)
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Double is a wrapper around the C function g_rand_double.
func (recv *Rand) Double() float64 {
	retC := C.g_rand_double((*C.GRand)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// DoubleRange is a wrapper around the C function g_rand_double_range.
func (recv *Rand) DoubleRange(begin float64, end float64) float64 {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	retC := C.g_rand_double_range((*C.GRand)(recv.native), c_begin, c_end)
	retGo := (float64)(retC)

	return retGo
}

// Free is a wrapper around the C function g_rand_free.
func (recv *Rand) Free() {
	C.g_rand_free((*C.GRand)(recv.native))

	return
}

// Int is a wrapper around the C function g_rand_int.
func (recv *Rand) Int() uint32 {
	retC := C.g_rand_int((*C.GRand)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IntRange is a wrapper around the C function g_rand_int_range.
func (recv *Rand) IntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_rand_int_range((*C.GRand)(recv.native), c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// SetSeed is a wrapper around the C function g_rand_set_seed.
func (recv *Rand) SetSeed(seed uint32) {
	c_seed := (C.guint32)(seed)

	C.g_rand_set_seed((*C.GRand)(recv.native), c_seed)

	return
}

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

// SListAlloc is a wrapper around the C function g_slist_alloc.
func SListAlloc() *SList {
	retC := C.g_slist_alloc()
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// SListConcat is a wrapper around the C function g_slist_concat.
func SListConcat(list1 *SList, list2 *SList) *SList {
	c_list1 := (*C.GSList)(C.NULL)
	if list1 != nil {
		c_list1 = (*C.GSList)(list1.ToC())
	}

	c_list2 := (*C.GSList)(C.NULL)
	if list2 != nil {
		c_list2 = (*C.GSList)(list2.ToC())
	}

	retC := C.g_slist_concat(c_list1, c_list2)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListCopy is a wrapper around the C function g_slist_copy.
func SListCopy(list *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_copy(c_list)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListDeleteLink is a wrapper around the C function g_slist_delete_link.
func SListDeleteLink(list *SList, link *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_link_ := (*C.GSList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GSList)(link.ToC())
	}

	retC := C.g_slist_delete_link(c_list, c_link_)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func
// SListFree is a wrapper around the C function g_slist_free.
func SListFree(list *SList) {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	C.g_slist_free(c_list)

	return
}

// SListFree1 is a wrapper around the C function g_slist_free_1.
func SListFree1(list *SList) {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	C.g_slist_free_1(c_list)

	return
}

// g_slist_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_insert : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// SListLast is a wrapper around the C function g_slist_last.
func SListLast(list *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_last(c_list)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListLength is a wrapper around the C function g_slist_length.
func SListLength(list *SList) uint32 {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_length(c_list)
	retGo := (uint32)(retC)

	return retGo
}

// SListNth is a wrapper around the C function g_slist_nth.
func SListNth(list *SList, n uint32) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_slist_nth(c_list, c_n)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_nth_data : no return generator
// SListPosition is a wrapper around the C function g_slist_position.
func SListPosition(list *SList, llink *SList) int32 {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_llink := (*C.GSList)(C.NULL)
	if llink != nil {
		c_llink = (*C.GSList)(llink.ToC())
	}

	retC := C.g_slist_position(c_list, c_llink)
	retGo := (int32)(retC)

	return retGo
}

// g_slist_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// SListRemoveLink is a wrapper around the C function g_slist_remove_link.
func SListRemoveLink(list *SList, link *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_link_ := (*C.GSList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GSList)(link.ToC())
	}

	retC := C.g_slist_remove_link(c_list, c_link_)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListReverse is a wrapper around the C function g_slist_reverse.
func SListReverse(list *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_reverse(c_list)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_sort : unsupported parameter compare_func : no type generator for CompareFunc (GCompareFunc) for param compare_func
// g_slist_sort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func
// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native *C.GScanner
	// user_data : no type generator for gpointer, gpointer
	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	// qdata : record
	// config : record
	Token TokenType
	// value : no type generator for TokenValue, GTokenValue
	Line      uint32
	Position  uint32
	NextToken TokenType
	// next_value : no type generator for TokenValue, GTokenValue
	NextLine     uint32
	NextPosition uint32
	// Private : symbol_table
	// Private : input_fd
	// Private : text
	// Private : text_end
	// Private : buffer
	// Private : scope_id
	// msg_handler : no type generator for ScannerMsgFunc, GScannerMsgFunc
}

func ScannerNewFromC(u unsafe.Pointer) *Scanner {
	c := (*C.GScanner)(u)
	if c == nil {
		return nil
	}

	g := &Scanner{
		InputName:      C.GoString(c.input_name),
		Line:           (uint32)(c.line),
		MaxParseErrors: (uint32)(c.max_parse_errors),
		NextLine:       (uint32)(c.next_line),
		NextPosition:   (uint32)(c.next_position),
		NextToken:      (TokenType)(c.next_token),
		ParseErrors:    (uint32)(c.parse_errors),
		Position:       (uint32)(c.position),
		Token:          (TokenType)(c.token),
		native:         c,
	}

	return g
}

func (recv *Scanner) ToC() unsafe.Pointer {
	recv.native.max_parse_errors =
		(C.guint)(recv.MaxParseErrors)
	recv.native.parse_errors =
		(C.guint)(recv.ParseErrors)
	recv.native.input_name =
		C.CString(recv.InputName)
	recv.native.token =
		(C.GTokenType)(recv.Token)
	recv.native.line =
		(C.guint)(recv.Line)
	recv.native.position =
		(C.guint)(recv.Position)
	recv.native.next_token =
		(C.GTokenType)(recv.NextToken)
	recv.native.next_line =
		(C.guint)(recv.NextLine)
	recv.native.next_position =
		(C.guint)(recv.NextPosition)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Scanner with another Scanner, and returns true if they represent the same GObject.
func (recv *Scanner) Equals(other *Scanner) bool {
	return other.ToC() == recv.ToC()
}

// ScannerNew is a wrapper around the C function g_scanner_new.
func ScannerNew(configTempl *ScannerConfig) *Scanner {
	c_config_templ := (*C.GScannerConfig)(C.NULL)
	if configTempl != nil {
		c_config_templ = (*C.GScannerConfig)(configTempl.ToC())
	}

	retC := C.g_scanner_new(c_config_templ)
	retGo := ScannerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CurLine is a wrapper around the C function g_scanner_cur_line.
func (recv *Scanner) CurLine() uint32 {
	retC := C.g_scanner_cur_line((*C.GScanner)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// CurPosition is a wrapper around the C function g_scanner_cur_position.
func (recv *Scanner) CurPosition() uint32 {
	retC := C.g_scanner_cur_position((*C.GScanner)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// CurToken is a wrapper around the C function g_scanner_cur_token.
func (recv *Scanner) CurToken() TokenType {
	retC := C.g_scanner_cur_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Unsupported : g_scanner_cur_value : no return generator

// Destroy is a wrapper around the C function g_scanner_destroy.
func (recv *Scanner) Destroy() {
	C.g_scanner_destroy((*C.GScanner)(recv.native))

	return
}

// Eof is a wrapper around the C function g_scanner_eof.
func (recv *Scanner) Eof() bool {
	retC := C.g_scanner_eof((*C.GScanner)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Error is a wrapper around the C function g_scanner_error.
func (recv *Scanner) Error(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_scanner_error((*C.GScanner)(recv.native), c_format)

	return
}

// GetNextToken is a wrapper around the C function g_scanner_get_next_token.
func (recv *Scanner) GetNextToken() TokenType {
	retC := C.g_scanner_get_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// InputFile is a wrapper around the C function g_scanner_input_file.
func (recv *Scanner) InputFile(inputFd int32) {
	c_input_fd := (C.gint)(inputFd)

	C.g_scanner_input_file((*C.GScanner)(recv.native), c_input_fd)

	return
}

// InputText is a wrapper around the C function g_scanner_input_text.
func (recv *Scanner) InputText(text string, textLen uint32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_text_len := (C.guint)(textLen)

	C.g_scanner_input_text((*C.GScanner)(recv.native), c_text, c_text_len)

	return
}

// Unsupported : g_scanner_lookup_symbol : no return generator

// PeekNextToken is a wrapper around the C function g_scanner_peek_next_token.
func (recv *Scanner) PeekNextToken() TokenType {
	retC := C.g_scanner_peek_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Unsupported : g_scanner_scope_add_symbol : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc (GHFunc) for param func

// Unsupported : g_scanner_scope_lookup_symbol : no return generator

// ScopeRemoveSymbol is a wrapper around the C function g_scanner_scope_remove_symbol.
func (recv *Scanner) ScopeRemoveSymbol(scopeId uint32, symbol string) {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	C.g_scanner_scope_remove_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol)

	return
}

// SetScope is a wrapper around the C function g_scanner_set_scope.
func (recv *Scanner) SetScope(scopeId uint32) uint32 {
	c_scope_id := (C.guint)(scopeId)

	retC := C.g_scanner_set_scope((*C.GScanner)(recv.native), c_scope_id)
	retGo := (uint32)(retC)

	return retGo
}

// SyncFileOffset is a wrapper around the C function g_scanner_sync_file_offset.
func (recv *Scanner) SyncFileOffset() {
	C.g_scanner_sync_file_offset((*C.GScanner)(recv.native))

	return
}

// UnexpToken is a wrapper around the C function g_scanner_unexp_token.
func (recv *Scanner) UnexpToken(expectedToken TokenType, identifierSpec string, symbolSpec string, symbolName string, message string, isError int32) {
	c_expected_token := (C.GTokenType)(expectedToken)

	c_identifier_spec := C.CString(identifierSpec)
	defer C.free(unsafe.Pointer(c_identifier_spec))

	c_symbol_spec := C.CString(symbolSpec)
	defer C.free(unsafe.Pointer(c_symbol_spec))

	c_symbol_name := C.CString(symbolName)
	defer C.free(unsafe.Pointer(c_symbol_name))

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	c_is_error := (C.gint)(isError)

	C.g_scanner_unexp_token((*C.GScanner)(recv.native), c_expected_token, c_identifier_spec, c_symbol_spec, c_symbol_name, c_message, c_is_error)

	return
}

// Warn is a wrapper around the C function g_scanner_warn.
func (recv *Scanner) Warn(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_scanner_warn((*C.GScanner)(recv.native), c_format)

	return
}

// ScannerConfig is a wrapper around the C record GScannerConfig.
type ScannerConfig struct {
	native              *C.GScannerConfig
	CsetSkipCharacters  string
	CsetIdentifierFirst string
	CsetIdentifierNth   string
	CpairCommentSingle  string
	// Bitfield not supported :  1 case_sensitive
	// Bitfield not supported :  1 skip_comment_multi
	// Bitfield not supported :  1 skip_comment_single
	// Bitfield not supported :  1 scan_comment_multi
	// Bitfield not supported :  1 scan_identifier
	// Bitfield not supported :  1 scan_identifier_1char
	// Bitfield not supported :  1 scan_identifier_NULL
	// Bitfield not supported :  1 scan_symbols
	// Bitfield not supported :  1 scan_binary
	// Bitfield not supported :  1 scan_octal
	// Bitfield not supported :  1 scan_float
	// Bitfield not supported :  1 scan_hex
	// Bitfield not supported :  1 scan_hex_dollar
	// Bitfield not supported :  1 scan_string_sq
	// Bitfield not supported :  1 scan_string_dq
	// Bitfield not supported :  1 numbers_2_int
	// Bitfield not supported :  1 int_2_float
	// Bitfield not supported :  1 identifier_2_string
	// Bitfield not supported :  1 char_2_token
	// Bitfield not supported :  1 symbol_2_token
	// Bitfield not supported :  1 scope_0_fallback
	// Bitfield not supported :  1 store_int64
	// Private : padding_dummy
}

func ScannerConfigNewFromC(u unsafe.Pointer) *ScannerConfig {
	c := (*C.GScannerConfig)(u)
	if c == nil {
		return nil
	}

	g := &ScannerConfig{
		CpairCommentSingle:  C.GoString(c.cpair_comment_single),
		CsetIdentifierFirst: C.GoString(c.cset_identifier_first),
		CsetIdentifierNth:   C.GoString(c.cset_identifier_nth),
		CsetSkipCharacters:  C.GoString(c.cset_skip_characters),
		native:              c,
	}

	return g
}

func (recv *ScannerConfig) ToC() unsafe.Pointer {
	recv.native.cset_skip_characters =
		C.CString(recv.CsetSkipCharacters)
	recv.native.cset_identifier_first =
		C.CString(recv.CsetIdentifierFirst)
	recv.native.cset_identifier_nth =
		C.CString(recv.CsetIdentifierNth)
	recv.native.cpair_comment_single =
		C.CString(recv.CpairCommentSingle)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ScannerConfig with another ScannerConfig, and returns true if they represent the same GObject.
func (recv *ScannerConfig) Equals(other *ScannerConfig) bool {
	return other.ToC() == recv.ToC()
}

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native *C.GSequence
}

func SequenceNewFromC(u unsafe.Pointer) *Sequence {
	c := (*C.GSequence)(u)
	if c == nil {
		return nil
	}

	g := &Sequence{native: c}

	return g
}

func (recv *Sequence) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Sequence with another Sequence, and returns true if they represent the same GObject.
func (recv *Sequence) Equals(other *Sequence) bool {
	return other.ToC() == recv.ToC()
}

// SequenceIter is a wrapper around the C record GSequenceIter.
type SequenceIter struct {
	native *C.GSequenceIter
}

func SequenceIterNewFromC(u unsafe.Pointer) *SequenceIter {
	c := (*C.GSequenceIter)(u)
	if c == nil {
		return nil
	}

	g := &SequenceIter{native: c}

	return g
}

func (recv *SequenceIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SequenceIter with another SequenceIter, and returns true if they represent the same GObject.
func (recv *SequenceIter) Equals(other *SequenceIter) bool {
	return other.ToC() == recv.ToC()
}

// Source is a wrapper around the C record GSource.
type Source struct {
	native *C.GSource
	// Private : callback_data
	// Private : callback_funcs
	// Private : source_funcs
	// Private : ref_count
	// Private : context
	// Private : priority
	// Private : flags
	// Private : source_id
	// Private : poll_fds
	// Private : prev
	// Private : next
	// Private : name
	// Private : priv
}

func SourceNewFromC(u unsafe.Pointer) *Source {
	c := (*C.GSource)(u)
	if c == nil {
		return nil
	}

	g := &Source{native: c}

	return g
}

func (recv *Source) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Source with another Source, and returns true if they represent the same GObject.
func (recv *Source) Equals(other *Source) bool {
	return other.ToC() == recv.ToC()
}

// SourceNew is a wrapper around the C function g_source_new.
func SourceNew(sourceFuncs *SourceFuncs, structSize uint32) *Source {
	c_source_funcs := (*C.GSourceFuncs)(C.NULL)
	if sourceFuncs != nil {
		c_source_funcs = (*C.GSourceFuncs)(sourceFuncs.ToC())
	}

	c_struct_size := (C.guint)(structSize)

	retC := C.g_source_new(c_source_funcs, c_struct_size)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SourceRemove is a wrapper around the C function g_source_remove.
func SourceRemove(tag uint32) bool {
	c_tag := (C.guint)(tag)

	retC := C.g_source_remove(c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// g_source_remove_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data
// g_source_remove_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data
// AddPoll is a wrapper around the C function g_source_add_poll.
func (recv *Source) AddPoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_source_add_poll((*C.GSource)(recv.native), c_fd)

	return
}

// Attach is a wrapper around the C function g_source_attach.
func (recv *Source) Attach(context *MainContext) uint32 {
	c_context := (*C.GMainContext)(C.NULL)
	if context != nil {
		c_context = (*C.GMainContext)(context.ToC())
	}

	retC := C.g_source_attach((*C.GSource)(recv.native), c_context)
	retGo := (uint32)(retC)

	return retGo
}

// Destroy is a wrapper around the C function g_source_destroy.
func (recv *Source) Destroy() {
	C.g_source_destroy((*C.GSource)(recv.native))

	return
}

// GetCanRecurse is a wrapper around the C function g_source_get_can_recurse.
func (recv *Source) GetCanRecurse() bool {
	retC := C.g_source_get_can_recurse((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetContext is a wrapper around the C function g_source_get_context.
func (recv *Source) GetContext() *MainContext {
	retC := C.g_source_get_context((*C.GSource)(recv.native))
	var retGo (*MainContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MainContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetCurrentTime is a wrapper around the C function g_source_get_current_time.
func (recv *Source) GetCurrentTime(timeval *TimeVal) {
	c_timeval := (*C.GTimeVal)(C.NULL)
	if timeval != nil {
		c_timeval = (*C.GTimeVal)(timeval.ToC())
	}

	C.g_source_get_current_time((*C.GSource)(recv.native), c_timeval)

	return
}

// GetId is a wrapper around the C function g_source_get_id.
func (recv *Source) GetId() uint32 {
	retC := C.g_source_get_id((*C.GSource)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetPriority is a wrapper around the C function g_source_get_priority.
func (recv *Source) GetPriority() int32 {
	retC := C.g_source_get_priority((*C.GSource)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetReadyTime is a wrapper around the C function g_source_get_ready_time.
func (recv *Source) GetReadyTime() int64 {
	retC := C.g_source_get_ready_time((*C.GSource)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Ref is a wrapper around the C function g_source_ref.
func (recv *Source) Ref() *Source {
	retC := C.g_source_ref((*C.GSource)(recv.native))
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemovePoll is a wrapper around the C function g_source_remove_poll.
func (recv *Source) RemovePoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_source_remove_poll((*C.GSource)(recv.native), c_fd)

	return
}

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc (GSourceFunc) for param func

// Unsupported : g_source_set_callback_indirect : unsupported parameter callback_data : no type generator for gpointer (gpointer) for param callback_data

// SetCanRecurse is a wrapper around the C function g_source_set_can_recurse.
func (recv *Source) SetCanRecurse(canRecurse bool) {
	c_can_recurse :=
		boolToGboolean(canRecurse)

	C.g_source_set_can_recurse((*C.GSource)(recv.native), c_can_recurse)

	return
}

// SetPriority is a wrapper around the C function g_source_set_priority.
func (recv *Source) SetPriority(priority int32) {
	c_priority := (C.gint)(priority)

	C.g_source_set_priority((*C.GSource)(recv.native), c_priority)

	return
}

// Unref is a wrapper around the C function g_source_unref.
func (recv *Source) Unref() {
	C.g_source_unref((*C.GSource)(recv.native))

	return
}

// SourceCallbackFuncs is a wrapper around the C record GSourceCallbackFuncs.
type SourceCallbackFuncs struct {
	native *C.GSourceCallbackFuncs
	// no type for ref
	// no type for unref
	// no type for get
}

func SourceCallbackFuncsNewFromC(u unsafe.Pointer) *SourceCallbackFuncs {
	c := (*C.GSourceCallbackFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceCallbackFuncs{native: c}

	return g
}

func (recv *SourceCallbackFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SourceCallbackFuncs with another SourceCallbackFuncs, and returns true if they represent the same GObject.
func (recv *SourceCallbackFuncs) Equals(other *SourceCallbackFuncs) bool {
	return other.ToC() == recv.ToC()
}

// SourceFuncs is a wrapper around the C record GSourceFuncs.
type SourceFuncs struct {
	native *C.GSourceFuncs
	// no type for prepare
	// no type for check
	// no type for dispatch
	// no type for finalize
	// Private : closure_callback
	// Private : closure_marshal
}

func SourceFuncsNewFromC(u unsafe.Pointer) *SourceFuncs {
	c := (*C.GSourceFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceFuncs{native: c}

	return g
}

func (recv *SourceFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SourceFuncs with another SourceFuncs, and returns true if they represent the same GObject.
func (recv *SourceFuncs) Equals(other *SourceFuncs) bool {
	return other.ToC() == recv.ToC()
}

// SourcePrivate is a wrapper around the C record GSourcePrivate.
type SourcePrivate struct {
	native *C.GSourcePrivate
}

func SourcePrivateNewFromC(u unsafe.Pointer) *SourcePrivate {
	c := (*C.GSourcePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SourcePrivate{native: c}

	return g
}

func (recv *SourcePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SourcePrivate with another SourcePrivate, and returns true if they represent the same GObject.
func (recv *SourcePrivate) Equals(other *SourcePrivate) bool {
	return other.ToC() == recv.ToC()
}

// StatBuf is a wrapper around the C record GStatBuf.
type StatBuf struct {
	native *C.GStatBuf
}

func StatBufNewFromC(u unsafe.Pointer) *StatBuf {
	c := (*C.GStatBuf)(u)
	if c == nil {
		return nil
	}

	g := &StatBuf{native: c}

	return g
}

func (recv *StatBuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StatBuf with another StatBuf, and returns true if they represent the same GObject.
func (recv *StatBuf) Equals(other *StatBuf) bool {
	return other.ToC() == recv.ToC()
}

// String is a wrapper around the C record GString.
type String struct {
	native       *C.GString
	Str          string
	Len          uint64
	AllocatedLen uint64
}

func StringNewFromC(u unsafe.Pointer) *String {
	c := (*C.GString)(u)
	if c == nil {
		return nil
	}

	g := &String{
		AllocatedLen: (uint64)(c.allocated_len),
		Len:          (uint64)(c.len),
		Str:          C.GoString(c.str),
		native:       c,
	}

	return g
}

func (recv *String) ToC() unsafe.Pointer {
	recv.native.str =
		C.CString(recv.Str)
	recv.native.len =
		(C.gsize)(recv.Len)
	recv.native.allocated_len =
		(C.gsize)(recv.AllocatedLen)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this String with another String, and returns true if they represent the same GObject.
func (recv *String) Equals(other *String) bool {
	return other.ToC() == recv.ToC()
}

// Append is a wrapper around the C function g_string_append.
func (recv *String) Append(val string) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_append((*C.GString)(recv.native), c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendC is a wrapper around the C function g_string_append_c.
func (recv *String) AppendC(c rune) *String {
	c_c := (C.gchar)(c)

	retC := C.g_string_append_c((*C.GString)(recv.native), c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendLen is a wrapper around the C function g_string_append_len.
func (recv *String) AppendLen(val string, len int64) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_append_len((*C.GString)(recv.native), c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendPrintf is a wrapper around the C function g_string_append_printf.
func (recv *String) AppendPrintf(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_string_append_printf((*C.GString)(recv.native), c_format)

	return
}

// AppendUnichar is a wrapper around the C function g_string_append_unichar.
func (recv *String) AppendUnichar(wc rune) *String {
	c_wc := (C.gunichar)(wc)

	retC := C.g_string_append_unichar((*C.GString)(recv.native), c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AsciiDown is a wrapper around the C function g_string_ascii_down.
func (recv *String) AsciiDown() *String {
	retC := C.g_string_ascii_down((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AsciiUp is a wrapper around the C function g_string_ascii_up.
func (recv *String) AsciiUp() *String {
	retC := C.g_string_ascii_up((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Assign is a wrapper around the C function g_string_assign.
func (recv *String) Assign(rval string) *String {
	c_rval := C.CString(rval)
	defer C.free(unsafe.Pointer(c_rval))

	retC := C.g_string_assign((*C.GString)(recv.native), c_rval)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Down is a wrapper around the C function g_string_down.
func (recv *String) Down() *String {
	retC := C.g_string_down((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_string_equal.
func (recv *String) Equal(v2 *String) bool {
	c_v2 := (*C.GString)(C.NULL)
	if v2 != nil {
		c_v2 = (*C.GString)(v2.ToC())
	}

	retC := C.g_string_equal((*C.GString)(recv.native), c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Erase is a wrapper around the C function g_string_erase.
func (recv *String) Erase(pos int64, len int64) *String {
	c_pos := (C.gssize)(pos)

	c_len := (C.gssize)(len)

	retC := C.g_string_erase((*C.GString)(recv.native), c_pos, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_string_free.
func (recv *String) Free(freeSegment bool) string {
	c_free_segment :=
		boolToGboolean(freeSegment)

	retC := C.g_string_free((*C.GString)(recv.native), c_free_segment)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Hash is a wrapper around the C function g_string_hash.
func (recv *String) Hash() uint32 {
	retC := C.g_string_hash((*C.GString)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Insert is a wrapper around the C function g_string_insert.
func (recv *String) Insert(pos int64, val string) *String {
	c_pos := (C.gssize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_insert((*C.GString)(recv.native), c_pos, c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertC is a wrapper around the C function g_string_insert_c.
func (recv *String) InsertC(pos int64, c rune) *String {
	c_pos := (C.gssize)(pos)

	c_c := (C.gchar)(c)

	retC := C.g_string_insert_c((*C.GString)(recv.native), c_pos, c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertLen is a wrapper around the C function g_string_insert_len.
func (recv *String) InsertLen(pos int64, val string, len int64) *String {
	c_pos := (C.gssize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_insert_len((*C.GString)(recv.native), c_pos, c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertUnichar is a wrapper around the C function g_string_insert_unichar.
func (recv *String) InsertUnichar(pos int64, wc rune) *String {
	c_pos := (C.gssize)(pos)

	c_wc := (C.gunichar)(wc)

	retC := C.g_string_insert_unichar((*C.GString)(recv.native), c_pos, c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prepend is a wrapper around the C function g_string_prepend.
func (recv *String) Prepend(val string) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_prepend((*C.GString)(recv.native), c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrependC is a wrapper around the C function g_string_prepend_c.
func (recv *String) PrependC(c rune) *String {
	c_c := (C.gchar)(c)

	retC := C.g_string_prepend_c((*C.GString)(recv.native), c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrependLen is a wrapper around the C function g_string_prepend_len.
func (recv *String) PrependLen(val string, len int64) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_prepend_len((*C.GString)(recv.native), c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrependUnichar is a wrapper around the C function g_string_prepend_unichar.
func (recv *String) PrependUnichar(wc rune) *String {
	c_wc := (C.gunichar)(wc)

	retC := C.g_string_prepend_unichar((*C.GString)(recv.native), c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Printf is a wrapper around the C function g_string_printf.
func (recv *String) Printf(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_string_printf((*C.GString)(recv.native), c_format)

	return
}

// SetSize is a wrapper around the C function g_string_set_size.
func (recv *String) SetSize(len uint64) *String {
	c_len := (C.gsize)(len)

	retC := C.g_string_set_size((*C.GString)(recv.native), c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Truncate is a wrapper around the C function g_string_truncate.
func (recv *String) Truncate(len uint64) *String {
	c_len := (C.gsize)(len)

	retC := C.g_string_truncate((*C.GString)(recv.native), c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Up is a wrapper around the C function g_string_up.
func (recv *String) Up() *String {
	retC := C.g_string_up((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StringChunk is a wrapper around the C record GStringChunk.
type StringChunk struct {
	native *C.GStringChunk
}

func StringChunkNewFromC(u unsafe.Pointer) *StringChunk {
	c := (*C.GStringChunk)(u)
	if c == nil {
		return nil
	}

	g := &StringChunk{native: c}

	return g
}

func (recv *StringChunk) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StringChunk with another StringChunk, and returns true if they represent the same GObject.
func (recv *StringChunk) Equals(other *StringChunk) bool {
	return other.ToC() == recv.ToC()
}

// StringChunkNew is a wrapper around the C function g_string_chunk_new.
func StringChunkNew(size uint64) *StringChunk {
	c_size := (C.gsize)(size)

	retC := C.g_string_chunk_new(c_size)
	retGo := StringChunkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_string_chunk_free.
func (recv *StringChunk) Free() {
	C.g_string_chunk_free((*C.GStringChunk)(recv.native))

	return
}

// Insert is a wrapper around the C function g_string_chunk_insert.
func (recv *StringChunk) Insert(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_string_chunk_insert((*C.GStringChunk)(recv.native), c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// InsertConst is a wrapper around the C function g_string_chunk_insert_const.
func (recv *StringChunk) InsertConst(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_string_chunk_insert_const((*C.GStringChunk)(recv.native), c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// TestCase is a wrapper around the C record GTestCase.
type TestCase struct {
	native *C.GTestCase
}

func TestCaseNewFromC(u unsafe.Pointer) *TestCase {
	c := (*C.GTestCase)(u)
	if c == nil {
		return nil
	}

	g := &TestCase{native: c}

	return g
}

func (recv *TestCase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestCase with another TestCase, and returns true if they represent the same GObject.
func (recv *TestCase) Equals(other *TestCase) bool {
	return other.ToC() == recv.ToC()
}

// TestConfig is a wrapper around the C record GTestConfig.
type TestConfig struct {
	native          *C.GTestConfig
	TestInitialized bool
	TestQuick       bool
	TestPerf        bool
	TestVerbose     bool
	TestQuiet       bool
	TestUndefined   bool
}

func TestConfigNewFromC(u unsafe.Pointer) *TestConfig {
	c := (*C.GTestConfig)(u)
	if c == nil {
		return nil
	}

	g := &TestConfig{
		TestInitialized: c.test_initialized == C.TRUE,
		TestPerf:        c.test_perf == C.TRUE,
		TestQuick:       c.test_quick == C.TRUE,
		TestQuiet:       c.test_quiet == C.TRUE,
		TestUndefined:   c.test_undefined == C.TRUE,
		TestVerbose:     c.test_verbose == C.TRUE,
		native:          c,
	}

	return g
}

func (recv *TestConfig) ToC() unsafe.Pointer {
	recv.native.test_initialized =
		boolToGboolean(recv.TestInitialized)
	recv.native.test_quick =
		boolToGboolean(recv.TestQuick)
	recv.native.test_perf =
		boolToGboolean(recv.TestPerf)
	recv.native.test_verbose =
		boolToGboolean(recv.TestVerbose)
	recv.native.test_quiet =
		boolToGboolean(recv.TestQuiet)
	recv.native.test_undefined =
		boolToGboolean(recv.TestUndefined)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestConfig with another TestConfig, and returns true if they represent the same GObject.
func (recv *TestConfig) Equals(other *TestConfig) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// TestSuite is a wrapper around the C record GTestSuite.
type TestSuite struct {
	native *C.GTestSuite
}

func TestSuiteNewFromC(u unsafe.Pointer) *TestSuite {
	c := (*C.GTestSuite)(u)
	if c == nil {
		return nil
	}

	g := &TestSuite{native: c}

	return g
}

func (recv *TestSuite) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestSuite with another TestSuite, and returns true if they represent the same GObject.
func (recv *TestSuite) Equals(other *TestSuite) bool {
	return other.ToC() == recv.ToC()
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native *C.GThread
}

func ThreadNewFromC(u unsafe.Pointer) *Thread {
	c := (*C.GThread)(u)
	if c == nil {
		return nil
	}

	g := &Thread{native: c}

	return g
}

func (recv *Thread) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Thread with another Thread, and returns true if they represent the same GObject.
func (recv *Thread) Equals(other *Thread) bool {
	return other.ToC() == recv.ToC()
}

// ThreadErrorQuark is a wrapper around the C function g_thread_error_quark.
func ThreadErrorQuark() Quark {
	retC := C.g_thread_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// g_thread_exit : unsupported parameter retval : no type generator for gpointer (gpointer) for param retval
// ThreadSelf is a wrapper around the C function g_thread_self.
func ThreadSelf() *Thread {
	retC := C.g_thread_self()
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ThreadYield is a wrapper around the C function g_thread_yield.
func ThreadYield() {
	C.g_thread_yield()

	return
}

// Unsupported : g_thread_join : no return generator

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	// user_data : no type generator for gpointer, gpointer
	Exclusive bool
}

func ThreadPoolNewFromC(u unsafe.Pointer) *ThreadPool {
	c := (*C.GThreadPool)(u)
	if c == nil {
		return nil
	}

	g := &ThreadPool{
		Exclusive: c.exclusive == C.TRUE,
		native:    c,
	}

	return g
}

func (recv *ThreadPool) ToC() unsafe.Pointer {
	recv.native.exclusive =
		boolToGboolean(recv.Exclusive)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThreadPool with another ThreadPool, and returns true if they represent the same GObject.
func (recv *ThreadPool) Equals(other *ThreadPool) bool {
	return other.ToC() == recv.ToC()
}

// ThreadPoolGetMaxUnusedThreads is a wrapper around the C function g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() int32 {
	retC := C.g_thread_pool_get_max_unused_threads()
	retGo := (int32)(retC)

	return retGo
}

// ThreadPoolGetNumUnusedThreads is a wrapper around the C function g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() uint32 {
	retC := C.g_thread_pool_get_num_unused_threads()
	retGo := (uint32)(retC)

	return retGo
}

// g_thread_pool_new : unsupported parameter func : no type generator for Func (GFunc) for param func
// ThreadPoolSetMaxUnusedThreads is a wrapper around the C function g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	c_max_threads := (C.gint)(maxThreads)

	C.g_thread_pool_set_max_unused_threads(c_max_threads)

	return
}

// ThreadPoolStopUnusedThreads is a wrapper around the C function g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {
	C.g_thread_pool_stop_unused_threads()

	return
}

// Free is a wrapper around the C function g_thread_pool_free.
func (recv *ThreadPool) Free(immediate bool, wait bool) {
	c_immediate :=
		boolToGboolean(immediate)

	c_wait_ :=
		boolToGboolean(wait)

	C.g_thread_pool_free((*C.GThreadPool)(recv.native), c_immediate, c_wait_)

	return
}

// GetMaxThreads is a wrapper around the C function g_thread_pool_get_max_threads.
func (recv *ThreadPool) GetMaxThreads() int32 {
	retC := C.g_thread_pool_get_max_threads((*C.GThreadPool)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNumThreads is a wrapper around the C function g_thread_pool_get_num_threads.
func (recv *ThreadPool) GetNumThreads() uint32 {
	retC := C.g_thread_pool_get_num_threads((*C.GThreadPool)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_thread_pool_push : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// SetMaxThreads is a wrapper around the C function g_thread_pool_set_max_threads.
func (recv *ThreadPool) SetMaxThreads(maxThreads int32) (bool, error) {
	c_max_threads := (C.gint)(maxThreads)

	var cThrowableError *C.GError

	retC := C.g_thread_pool_set_max_threads((*C.GThreadPool)(recv.native), c_max_threads, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unprocessed is a wrapper around the C function g_thread_pool_unprocessed.
func (recv *ThreadPool) Unprocessed() uint32 {
	retC := C.g_thread_pool_unprocessed((*C.GThreadPool)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// TimeVal is a wrapper around the C record GTimeVal.
type TimeVal struct {
	native *C.GTimeVal
	TvSec  int64
	TvUsec int64
}

func TimeValNewFromC(u unsafe.Pointer) *TimeVal {
	c := (*C.GTimeVal)(u)
	if c == nil {
		return nil
	}

	g := &TimeVal{
		TvSec:  (int64)(c.tv_sec),
		TvUsec: (int64)(c.tv_usec),
		native: c,
	}

	return g
}

func (recv *TimeVal) ToC() unsafe.Pointer {
	recv.native.tv_sec =
		(C.glong)(recv.TvSec)
	recv.native.tv_usec =
		(C.glong)(recv.TvUsec)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TimeVal with another TimeVal, and returns true if they represent the same GObject.
func (recv *TimeVal) Equals(other *TimeVal) bool {
	return other.ToC() == recv.ToC()
}

// Add is a wrapper around the C function g_time_val_add.
func (recv *TimeVal) Add(microseconds int64) {
	c_microseconds := (C.glong)(microseconds)

	C.g_time_val_add((*C.GTimeVal)(recv.native), c_microseconds)

	return
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native *C.GTimer
}

func TimerNewFromC(u unsafe.Pointer) *Timer {
	c := (*C.GTimer)(u)
	if c == nil {
		return nil
	}

	g := &Timer{native: c}

	return g
}

func (recv *Timer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Timer with another Timer, and returns true if they represent the same GObject.
func (recv *Timer) Equals(other *Timer) bool {
	return other.ToC() == recv.ToC()
}

// TimerNew is a wrapper around the C function g_timer_new.
func TimerNew() *Timer {
	retC := C.g_timer_new()
	retGo := TimerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy is a wrapper around the C function g_timer_destroy.
func (recv *Timer) Destroy() {
	C.g_timer_destroy((*C.GTimer)(recv.native))

	return
}

// Elapsed is a wrapper around the C function g_timer_elapsed.
func (recv *Timer) Elapsed(microseconds uint64) float64 {
	c_microseconds := (C.gulong)(microseconds)

	retC := C.g_timer_elapsed((*C.GTimer)(recv.native), &c_microseconds)
	retGo := (float64)(retC)

	return retGo
}

// Reset is a wrapper around the C function g_timer_reset.
func (recv *Timer) Reset() {
	C.g_timer_reset((*C.GTimer)(recv.native))

	return
}

// Start is a wrapper around the C function g_timer_start.
func (recv *Timer) Start() {
	C.g_timer_start((*C.GTimer)(recv.native))

	return
}

// Stop is a wrapper around the C function g_timer_stop.
func (recv *Timer) Stop() {
	C.g_timer_stop((*C.GTimer)(recv.native))

	return
}

// TrashStack is a wrapper around the C record GTrashStack.
type TrashStack struct {
	native *C.GTrashStack
	// next : record
}

func TrashStackNewFromC(u unsafe.Pointer) *TrashStack {
	c := (*C.GTrashStack)(u)
	if c == nil {
		return nil
	}

	g := &TrashStack{native: c}

	return g
}

func (recv *TrashStack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TrashStack with another TrashStack, and returns true if they represent the same GObject.
func (recv *TrashStack) Equals(other *TrashStack) bool {
	return other.ToC() == recv.ToC()
}

// TrashStackHeight is a wrapper around the C function g_trash_stack_height.
func TrashStackHeight(stackP *TrashStack) uint32 {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	retC := C.g_trash_stack_height(c_stack_p)
	retGo := (uint32)(retC)

	return retGo
}

// g_trash_stack_peek : no return generator
// g_trash_stack_pop : no return generator
// g_trash_stack_push : unsupported parameter data_p : no type generator for gpointer (gpointer) for param data_p
// Tree is a wrapper around the C record GTree.
type Tree struct {
	native *C.GTree
}

func TreeNewFromC(u unsafe.Pointer) *Tree {
	c := (*C.GTree)(u)
	if c == nil {
		return nil
	}

	g := &Tree{native: c}

	return g
}

func (recv *Tree) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Tree with another Tree, and returns true if they represent the same GObject.
func (recv *Tree) Equals(other *Tree) bool {
	return other.ToC() == recv.ToC()
}

// g_tree_new : unsupported parameter key_compare_func : no type generator for CompareFunc (GCompareFunc) for param key_compare_func
// g_tree_new_full : unsupported parameter key_compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param key_compare_func
// g_tree_new_with_data : unsupported parameter key_compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param key_compare_func
// Destroy is a wrapper around the C function g_tree_destroy.
func (recv *Tree) Destroy() {
	C.g_tree_destroy((*C.GTree)(recv.native))

	return
}

// Unsupported : g_tree_foreach : unsupported parameter func : no type generator for TraverseFunc (GTraverseFunc) for param func

// Height is a wrapper around the C function g_tree_height.
func (recv *Tree) Height() int32 {
	retC := C.g_tree_height((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_tree_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_tree_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_tree_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key

// Nnodes is a wrapper around the C function g_tree_nnodes.
func (recv *Tree) Nnodes() int32 {
	retC := C.g_tree_nnodes((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_tree_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_tree_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_tree_search : unsupported parameter search_func : no type generator for CompareFunc (GCompareFunc) for param search_func

// Unsupported : g_tree_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_tree_traverse : unsupported parameter traverse_func : no type generator for TraverseFunc (GTraverseFunc) for param traverse_func

// VariantBuilder is a wrapper around the C record GVariantBuilder.
type VariantBuilder struct {
	native *C.GVariantBuilder
}

func VariantBuilderNewFromC(u unsafe.Pointer) *VariantBuilder {
	c := (*C.GVariantBuilder)(u)
	if c == nil {
		return nil
	}

	g := &VariantBuilder{native: c}

	return g
}

func (recv *VariantBuilder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VariantBuilder with another VariantBuilder, and returns true if they represent the same GObject.
func (recv *VariantBuilder) Equals(other *VariantBuilder) bool {
	return other.ToC() == recv.ToC()
}

// VariantIter is a wrapper around the C record GVariantIter.
type VariantIter struct {
	native *C.GVariantIter
	// Private : x
}

func VariantIterNewFromC(u unsafe.Pointer) *VariantIter {
	c := (*C.GVariantIter)(u)
	if c == nil {
		return nil
	}

	g := &VariantIter{native: c}

	return g
}

func (recv *VariantIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VariantIter with another VariantIter, and returns true if they represent the same GObject.
func (recv *VariantIter) Equals(other *VariantIter) bool {
	return other.ToC() == recv.ToC()
}

// VariantType is a wrapper around the C record GVariantType.
type VariantType struct {
	native *C.GVariantType
}

func VariantTypeNewFromC(u unsafe.Pointer) *VariantType {
	c := (*C.GVariantType)(u)
	if c == nil {
		return nil
	}

	g := &VariantType{native: c}

	return g
}

func (recv *VariantType) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VariantType with another VariantType, and returns true if they represent the same GObject.
func (recv *VariantType) Equals(other *VariantType) bool {
	return other.ToC() == recv.ToC()
}

// VariantTypeNewArray is a wrapper around the C function g_variant_type_new_array.
func VariantTypeNewArray(element *VariantType) *VariantType {
	c_element := (*C.GVariantType)(C.NULL)
	if element != nil {
		c_element = (*C.GVariantType)(element.ToC())
	}

	retC := C.g_variant_type_new_array(c_element)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeNewDictEntry is a wrapper around the C function g_variant_type_new_dict_entry.
func VariantTypeNewDictEntry(key *VariantType, value *VariantType) *VariantType {
	c_key := (*C.GVariantType)(C.NULL)
	if key != nil {
		c_key = (*C.GVariantType)(key.ToC())
	}

	c_value := (*C.GVariantType)(C.NULL)
	if value != nil {
		c_value = (*C.GVariantType)(value.ToC())
	}

	retC := C.g_variant_type_new_dict_entry(c_key, c_value)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeNewMaybe is a wrapper around the C function g_variant_type_new_maybe.
func VariantTypeNewMaybe(element *VariantType) *VariantType {
	c_element := (*C.GVariantType)(C.NULL)
	if element != nil {
		c_element = (*C.GVariantType)(element.ToC())
	}

	retC := C.g_variant_type_new_maybe(c_element)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_type_new_tuple : unsupported parameter items :

// VariantTypeChecked is a wrapper around the C function g_variant_type_checked_.
func VariantTypeChecked(arg0 string) *VariantType {
	c_arg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(c_arg0))

	retC := C.g_variant_type_checked_(c_arg0)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeStringIsValid is a wrapper around the C function g_variant_type_string_is_valid.
func VariantTypeStringIsValid(typeString string) bool {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	retC := C.g_variant_type_string_is_valid(c_type_string)
	retGo := retC == C.TRUE

	return retGo
}

// Copy is a wrapper around the C function g_variant_type_copy.
func (recv *VariantType) Copy() *VariantType {
	retC := C.g_variant_type_copy((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DupString is a wrapper around the C function g_variant_type_dup_string.
func (recv *VariantType) DupString() string {
	retC := C.g_variant_type_dup_string((*C.GVariantType)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Element is a wrapper around the C function g_variant_type_element.
func (recv *VariantType) Element() *VariantType {
	retC := C.g_variant_type_element((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_variant_type_equal.
func (recv *VariantType) Equal(type2 *VariantType) bool {
	c_type2 := (C.gconstpointer)(C.NULL)
	if type2 != nil {
		c_type2 = (C.gconstpointer)(type2.ToC())
	}

	retC := C.g_variant_type_equal((C.gconstpointer)(recv.native), c_type2)
	retGo := retC == C.TRUE

	return retGo
}

// First is a wrapper around the C function g_variant_type_first.
func (recv *VariantType) First() *VariantType {
	retC := C.g_variant_type_first((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_variant_type_free.
func (recv *VariantType) Free() {
	C.g_variant_type_free((*C.GVariantType)(recv.native))

	return
}

// GetStringLength is a wrapper around the C function g_variant_type_get_string_length.
func (recv *VariantType) GetStringLength() uint64 {
	retC := C.g_variant_type_get_string_length((*C.GVariantType)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Hash is a wrapper around the C function g_variant_type_hash.
func (recv *VariantType) Hash() uint32 {
	retC := C.g_variant_type_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsArray is a wrapper around the C function g_variant_type_is_array.
func (recv *VariantType) IsArray() bool {
	retC := C.g_variant_type_is_array((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsBasic is a wrapper around the C function g_variant_type_is_basic.
func (recv *VariantType) IsBasic() bool {
	retC := C.g_variant_type_is_basic((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsContainer is a wrapper around the C function g_variant_type_is_container.
func (recv *VariantType) IsContainer() bool {
	retC := C.g_variant_type_is_container((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsDefinite is a wrapper around the C function g_variant_type_is_definite.
func (recv *VariantType) IsDefinite() bool {
	retC := C.g_variant_type_is_definite((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsDictEntry is a wrapper around the C function g_variant_type_is_dict_entry.
func (recv *VariantType) IsDictEntry() bool {
	retC := C.g_variant_type_is_dict_entry((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsMaybe is a wrapper around the C function g_variant_type_is_maybe.
func (recv *VariantType) IsMaybe() bool {
	retC := C.g_variant_type_is_maybe((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSubtypeOf is a wrapper around the C function g_variant_type_is_subtype_of.
func (recv *VariantType) IsSubtypeOf(supertype *VariantType) bool {
	c_supertype := (*C.GVariantType)(C.NULL)
	if supertype != nil {
		c_supertype = (*C.GVariantType)(supertype.ToC())
	}

	retC := C.g_variant_type_is_subtype_of((*C.GVariantType)(recv.native), c_supertype)
	retGo := retC == C.TRUE

	return retGo
}

// IsTuple is a wrapper around the C function g_variant_type_is_tuple.
func (recv *VariantType) IsTuple() bool {
	retC := C.g_variant_type_is_tuple((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsVariant is a wrapper around the C function g_variant_type_is_variant.
func (recv *VariantType) IsVariant() bool {
	retC := C.g_variant_type_is_variant((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Key is a wrapper around the C function g_variant_type_key.
func (recv *VariantType) Key() *VariantType {
	retC := C.g_variant_type_key((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NItems is a wrapper around the C function g_variant_type_n_items.
func (recv *VariantType) NItems() uint64 {
	retC := C.g_variant_type_n_items((*C.GVariantType)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Next is a wrapper around the C function g_variant_type_next.
func (recv *VariantType) Next() *VariantType {
	retC := C.g_variant_type_next((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PeekString is a wrapper around the C function g_variant_type_peek_string.
func (recv *VariantType) PeekString() string {
	retC := C.g_variant_type_peek_string((*C.GVariantType)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Value is a wrapper around the C function g_variant_type_value.
func (recv *VariantType) Value() *VariantType {
	retC := C.g_variant_type_value((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

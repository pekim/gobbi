// This is a generated file - DO NOT EDIT

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
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

func (recv *Array) toC() *C.GArray {
	recv.native.data =
		C.CString(recv.Data)
	recv.native.len =
		(C.guint)(recv.Len)

	return recv.native
}

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

func (recv *AsyncQueue) toC() *C.GAsyncQueue {

	return recv.native
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

// Unsupported : g_async_queue_lock : no return generator

// Pop is a wrapper around the C function g_async_queue_pop.
func (recv *AsyncQueue) Pop() uintptr {
	retC := C.g_async_queue_pop((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// PopUnlocked is a wrapper around the C function g_async_queue_pop_unlocked.
func (recv *AsyncQueue) PopUnlocked() uintptr {
	retC := C.g_async_queue_pop_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_async_queue_push : no return generator

// Unsupported : g_async_queue_push_front : no return generator

// Unsupported : g_async_queue_push_front_unlocked : no return generator

// Unsupported : g_async_queue_push_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_push_sorted_unlocked : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_push_unlocked : no return generator

// Ref is a wrapper around the C function g_async_queue_ref.
func (recv *AsyncQueue) Ref() *AsyncQueue {
	retC := C.g_async_queue_ref((*C.GAsyncQueue)(recv.native))
	retGo := AsyncQueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_async_queue_ref_unlocked : no return generator

// Unsupported : g_async_queue_remove : no return generator

// Unsupported : g_async_queue_remove_unlocked : no return generator

// Unsupported : g_async_queue_sort : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_sort_unlocked : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// TimedPop is a wrapper around the C function g_async_queue_timed_pop.
func (recv *AsyncQueue) TimedPop(endTime *TimeVal) uintptr {
	c_end_time := endTime.toC()

	retC := C.g_async_queue_timed_pop((*C.GAsyncQueue)(recv.native), c_end_time)
	retGo := (uintptr)(retC)

	return retGo
}

// TimedPopUnlocked is a wrapper around the C function g_async_queue_timed_pop_unlocked.
func (recv *AsyncQueue) TimedPopUnlocked(endTime *TimeVal) uintptr {
	c_end_time := endTime.toC()

	retC := C.g_async_queue_timed_pop_unlocked((*C.GAsyncQueue)(recv.native), c_end_time)
	retGo := (uintptr)(retC)

	return retGo
}

// TimeoutPop is a wrapper around the C function g_async_queue_timeout_pop.
func (recv *AsyncQueue) TimeoutPop(timeout uint64) uintptr {
	c_timeout := (C.guint64)(timeout)

	retC := C.g_async_queue_timeout_pop((*C.GAsyncQueue)(recv.native), c_timeout)
	retGo := (uintptr)(retC)

	return retGo
}

// TimeoutPopUnlocked is a wrapper around the C function g_async_queue_timeout_pop_unlocked.
func (recv *AsyncQueue) TimeoutPopUnlocked(timeout uint64) uintptr {
	c_timeout := (C.guint64)(timeout)

	retC := C.g_async_queue_timeout_pop_unlocked((*C.GAsyncQueue)(recv.native), c_timeout)
	retGo := (uintptr)(retC)

	return retGo
}

// TryPop is a wrapper around the C function g_async_queue_try_pop.
func (recv *AsyncQueue) TryPop() uintptr {
	retC := C.g_async_queue_try_pop((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// TryPopUnlocked is a wrapper around the C function g_async_queue_try_pop_unlocked.
func (recv *AsyncQueue) TryPopUnlocked() uintptr {
	retC := C.g_async_queue_try_pop_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_async_queue_unlock : no return generator

// Unsupported : g_async_queue_unref : no return generator

// Unsupported : g_async_queue_unref_and_unlock : no return generator

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

func (recv *BookmarkFile) toC() *C.GBookmarkFile {

	return recv.native
}

// Unsupported : g_bookmark_file_add_application : no return generator

// Unsupported : g_bookmark_file_add_group : no return generator

// Unsupported : g_bookmark_file_free : no return generator

// Unsupported : g_bookmark_file_get_added : no return generator

// Unsupported : g_bookmark_file_get_app_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : g_bookmark_file_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_icon : no return generator

// Unsupported : g_bookmark_file_get_is_private : no return generator

// Unsupported : g_bookmark_file_get_modified : no return generator

// Unsupported : g_bookmark_file_get_uris : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_visited : no return generator

// Unsupported : g_bookmark_file_has_application : no return generator

// Unsupported : g_bookmark_file_has_group : no return generator

// Unsupported : g_bookmark_file_has_item : no return generator

// Unsupported : g_bookmark_file_load_from_data : unsupported parameter data : no param type

// Unsupported : g_bookmark_file_load_from_data_dirs : no return generator

// Unsupported : g_bookmark_file_load_from_file : no return generator

// Unsupported : g_bookmark_file_move_item : no return generator

// Unsupported : g_bookmark_file_remove_application : no return generator

// Unsupported : g_bookmark_file_remove_group : no return generator

// Unsupported : g_bookmark_file_remove_item : no return generator

// Unsupported : g_bookmark_file_set_added : unsupported parameter added : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_app_info : unsupported parameter stamp : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_description : no return generator

// Unsupported : g_bookmark_file_set_groups : unsupported parameter groups : no param type

// Unsupported : g_bookmark_file_set_icon : no return generator

// Unsupported : g_bookmark_file_set_is_private : unsupported parameter is_private : no type generator for gboolean, gboolean

// Unsupported : g_bookmark_file_set_mime_type : no return generator

// Unsupported : g_bookmark_file_set_modified : unsupported parameter modified : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_title : no return generator

// Unsupported : g_bookmark_file_set_visited : unsupported parameter visited : no type generator for glong, time_t

// Unsupported : g_bookmark_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_to_file : no return generator

// ByteArray is a wrapper around the C record GByteArray.
type ByteArray struct {
	native *C.GByteArray
	// data : no type generator for guint8, guint8*
	Len uint32
}

func ByteArrayNewFromC(u unsafe.Pointer) *ByteArray {
	c := (*C.GByteArray)(u)
	if c == nil {
		return nil
	}

	g := &ByteArray{
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
}

func (recv *ByteArray) toC() *C.GByteArray {
	recv.native.len =
		(C.guint)(recv.Len)

	return recv.native
}

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

func (recv *Cond) toC() *C.GCond {

	return recv.native
}

// Unsupported : g_cond_broadcast : no return generator

// Unsupported : g_cond_clear : no return generator

// Unsupported : g_cond_init : no return generator

// Unsupported : g_cond_signal : no return generator

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex, GMutex*

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

func (recv *Data) toC() *C.GData {

	return recv.native
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

func (recv *Date) toC() *C.GDate {

	return recv.native
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

// Unsupported : g_date_add_days : no return generator

// Unsupported : g_date_add_months : no return generator

// Unsupported : g_date_add_years : no return generator

// Unsupported : g_date_clamp : no return generator

// Unsupported : g_date_clear : no return generator

// Compare is a wrapper around the C function g_date_compare.
func (recv *Date) Compare(rhs *Date) int32 {
	c_rhs := rhs.toC()

	retC := C.g_date_compare((*C.GDate)(recv.native), c_rhs)
	retGo := (int32)(retC)

	return retGo
}

// DaysBetween is a wrapper around the C function g_date_days_between.
func (recv *Date) DaysBetween(date2 *Date) int32 {
	c_date2 := date2.toC()

	retC := C.g_date_days_between((*C.GDate)(recv.native), c_date2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_date_free : no return generator

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

// Unsupported : g_date_is_first_of_month : no return generator

// Unsupported : g_date_is_last_of_month : no return generator

// Unsupported : g_date_order : no return generator

// Unsupported : g_date_set_day : no return generator

// Unsupported : g_date_set_dmy : no return generator

// Unsupported : g_date_set_julian : no return generator

// Unsupported : g_date_set_month : no return generator

// Unsupported : g_date_set_parse : no return generator

// Unsupported : g_date_set_time : no return generator

// Unsupported : g_date_set_time_t : unsupported parameter timet : no type generator for glong, time_t

// Unsupported : g_date_set_time_val : no return generator

// Unsupported : g_date_set_year : no return generator

// Unsupported : g_date_subtract_days : no return generator

// Unsupported : g_date_subtract_months : no return generator

// Unsupported : g_date_subtract_years : no return generator

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer, tm*

// Unsupported : g_date_valid : no return generator

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

func (recv *DebugKey) toC() *C.GDebugKey {
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		(C.guint)(recv.Value)

	return recv.native
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

func (recv *Dir) toC() *C.GDir {

	return recv.native
}

// Unsupported : g_dir_close : no return generator

// ReadName is a wrapper around the C function g_dir_read_name.
func (recv *Dir) ReadName() string {
	retC := C.g_dir_read_name((*C.GDir)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dir_rewind : no return generator

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

func (recv *Error) toC() *C.GError {
	recv.native.domain =
		(C.GQuark)(recv.Domain)
	recv.native.code =
		(C.gint)(recv.Code)
	recv.native.message =
		C.CString(recv.Message)

	return recv.native
}

// Unsupported : g_error_new : unsupported parameter ... : varargs

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

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// Copy is a wrapper around the C function g_error_copy.
func (recv *Error) Copy() *Error {
	retC := C.g_error_copy((*C.GError)(recv.native))
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_error_free : no return generator

// Unsupported : g_error_matches : no return generator

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

func (recv *HashTable) toC() *C.GHashTable {

	return recv.native
}

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

func (recv *HashTableIter) toC() *C.GHashTableIter {

	return recv.native
}

// Unsupported : g_hash_table_iter_init : no return generator

// Unsupported : g_hash_table_iter_next : unsupported parameter key : no type generator for gpointer, gpointer*

// Unsupported : g_hash_table_iter_remove : no return generator

// Unsupported : g_hash_table_iter_replace : no return generator

// Unsupported : g_hash_table_iter_steal : no return generator

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native *C.GHook
	Data   uintptr
	// next : record
	// prev : record
	RefCount uint32
	HookId   uint64
	Flags    uint32
	Func     uintptr
	// destroy : no type generator for DestroyNotify, GDestroyNotify
}

func HookNewFromC(u unsafe.Pointer) *Hook {
	c := (*C.GHook)(u)
	if c == nil {
		return nil
	}

	g := &Hook{
		Data:     (uintptr)(c.data),
		Flags:    (uint32)(c.flags),
		Func:     (uintptr)(c._func),
		HookId:   (uint64)(c.hook_id),
		RefCount: (uint32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *Hook) toC() *C.GHook {
	recv.native.data =
		(C.gpointer)(recv.Data)
	recv.native.ref_count =
		(C.guint)(recv.RefCount)
	recv.native.hook_id =
		(C.gulong)(recv.HookId)
	recv.native.flags =
		(C.guint)(recv.Flags)
	recv.native._func =
		(C.gpointer)(recv.Func)

	return recv.native
}

// CompareIds is a wrapper around the C function g_hook_compare_ids.
func (recv *Hook) CompareIds(sibling *Hook) int32 {
	c_sibling := sibling.toC()

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
	Dummy3 uintptr
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func HookListNewFromC(u unsafe.Pointer) *HookList {
	c := (*C.GHookList)(u)
	if c == nil {
		return nil
	}

	g := &HookList{
		Dummy3: (uintptr)(c.dummy3),
		SeqId:  (uint64)(c.seq_id),
		native: c,
	}

	return g
}

func (recv *HookList) toC() *C.GHookList {
	recv.native.seq_id =
		(C.gulong)(recv.SeqId)
	recv.native.dummy3 =
		(C.gpointer)(recv.Dummy3)

	return recv.native
}

// Unsupported : g_hook_list_clear : no return generator

// Unsupported : g_hook_list_init : no return generator

// Unsupported : g_hook_list_invoke : unsupported parameter may_recurse : no type generator for gboolean, gboolean

// Unsupported : g_hook_list_invoke_check : unsupported parameter may_recurse : no type generator for gboolean, gboolean

// Unsupported : g_hook_list_marshal : unsupported parameter may_recurse : no type generator for gboolean, gboolean

// Unsupported : g_hook_list_marshal_check : unsupported parameter may_recurse : no type generator for gboolean, gboolean

// Blacklisted : GIConv

// Blacklisted : GIOChannel

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

func (recv *IOFuncs) toC() *C.GIOFuncs {

	return recv.native
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

func (recv *KeyFile) toC() *C.GKeyFile {

	return recv.native
}

// KeyFileNew is a wrapper around the C function g_key_file_new.
func KeyFileNew() *KeyFile {
	retC := C.g_key_file_new()
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_key_file_free : no return generator

// Unsupported : g_key_file_get_boolean : no return generator

// Unsupported : g_key_file_get_boolean_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_double_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_integer_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_keys : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_locale_string_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_string_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_has_group : no return generator

// Unsupported : g_key_file_has_key : no return generator

// Unsupported : g_key_file_load_from_bytes : no return generator

// Unsupported : g_key_file_load_from_data : no return generator

// Unsupported : g_key_file_load_from_data_dirs : no return generator

// Unsupported : g_key_file_load_from_dirs : unsupported parameter search_dirs : no param type

// Unsupported : g_key_file_load_from_file : no return generator

// Unsupported : g_key_file_remove_comment : no return generator

// Unsupported : g_key_file_remove_group : no return generator

// Unsupported : g_key_file_remove_key : no return generator

// Unsupported : g_key_file_save_to_file : no return generator

// Unsupported : g_key_file_set_boolean : unsupported parameter value : no type generator for gboolean, gboolean

// Unsupported : g_key_file_set_boolean_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_comment : no return generator

// Unsupported : g_key_file_set_double : no return generator

// Unsupported : g_key_file_set_double_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_int64 : no return generator

// Unsupported : g_key_file_set_integer : no return generator

// Unsupported : g_key_file_set_integer_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_list_separator : no return generator

// Unsupported : g_key_file_set_locale_string : no return generator

// Unsupported : g_key_file_set_locale_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_string : no return generator

// Unsupported : g_key_file_set_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_uint64 : no return generator

// Unsupported : g_key_file_set_value : no return generator

// Unsupported : g_key_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_unref : no return generator

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	Data   uintptr
	// next : record
	// prev : record
}

func ListNewFromC(u unsafe.Pointer) *List {
	c := (*C.GList)(u)
	if c == nil {
		return nil
	}

	g := &List{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *List) toC() *C.GList {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return recv.native
}

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

func (recv *MainContext) toC() *C.GMainContext {

	return recv.native
}

// MainContextNew is a wrapper around the C function g_main_context_new.
func MainContextNew() *MainContext {
	retC := C.g_main_context_new()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_context_acquire : no return generator

// Unsupported : g_main_context_add_poll : no return generator

// Unsupported : g_main_context_check : unsupported parameter fds : no param type

// Unsupported : g_main_context_dispatch : no return generator

// FindSourceByFuncsUserData is a wrapper around the C function g_main_context_find_source_by_funcs_user_data.
func (recv *MainContext) FindSourceByFuncsUserData(funcs *SourceFuncs, userData uintptr) *Source {
	c_funcs := funcs.toC()

	c_user_data := (C.gpointer)(userData)

	retC := C.g_main_context_find_source_by_funcs_user_data((*C.GMainContext)(recv.native), c_funcs, c_user_data)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindSourceById is a wrapper around the C function g_main_context_find_source_by_id.
func (recv *MainContext) FindSourceById(sourceId uint32) *Source {
	c_source_id := (C.guint)(sourceId)

	retC := C.g_main_context_find_source_by_id((*C.GMainContext)(recv.native), c_source_id)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindSourceByUserData is a wrapper around the C function g_main_context_find_source_by_user_data.
func (recv *MainContext) FindSourceByUserData(userData uintptr) *Source {
	c_user_data := (C.gpointer)(userData)

	retC := C.g_main_context_find_source_by_user_data((*C.GMainContext)(recv.native), c_user_data)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_context_get_poll_func : no return generator

// Unsupported : g_main_context_invoke : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_invoke_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_is_owner : no return generator

// Unsupported : g_main_context_iteration : unsupported parameter may_block : no type generator for gboolean, gboolean

// Unsupported : g_main_context_pending : no return generator

// Unsupported : g_main_context_pop_thread_default : no return generator

// Unsupported : g_main_context_prepare : unsupported parameter priority : no type generator for gint, gint*

// Unsupported : g_main_context_push_thread_default : no return generator

// Unsupported : g_main_context_query : unsupported parameter timeout_ : no type generator for gint, gint*

// Ref is a wrapper around the C function g_main_context_ref.
func (recv *MainContext) Ref() *MainContext {
	retC := C.g_main_context_ref((*C.GMainContext)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_context_release : no return generator

// Unsupported : g_main_context_remove_poll : no return generator

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc, GPollFunc

// Unsupported : g_main_context_unref : no return generator

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_main_context_wakeup : no return generator

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

func (recv *MainLoop) toC() *C.GMainLoop {

	return recv.native
}

// Unsupported : g_main_loop_new : unsupported parameter is_running : no type generator for gboolean, gboolean

// GetContext is a wrapper around the C function g_main_loop_get_context.
func (recv *MainLoop) GetContext() *MainContext {
	retC := C.g_main_loop_get_context((*C.GMainLoop)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_loop_is_running : no return generator

// Unsupported : g_main_loop_quit : no return generator

// Ref is a wrapper around the C function g_main_loop_ref.
func (recv *MainLoop) Ref() *MainLoop {
	retC := C.g_main_loop_ref((*C.GMainLoop)(recv.native))
	retGo := MainLoopNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_loop_run : no return generator

// Unsupported : g_main_loop_unref : no return generator

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

func (recv *MappedFile) toC() *C.GMappedFile {

	return recv.native
}

// Unsupported : g_mapped_file_new : unsupported parameter writable : no type generator for gboolean, gboolean

// Unsupported : g_mapped_file_new_from_fd : unsupported parameter writable : no type generator for gboolean, gboolean

// Unsupported : g_mapped_file_free : no return generator

// Unsupported : g_mapped_file_unref : no return generator

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

func (recv *MarkupParseContext) toC() *C.GMarkupParseContext {

	return recv.native
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_markup_parse_context_end_parse : no return generator

// Unsupported : g_markup_parse_context_free : no return generator

// Unsupported : g_markup_parse_context_get_position : unsupported parameter line_number : no type generator for gint, gint*

// Unsupported : g_markup_parse_context_parse : no return generator

// Unsupported : g_markup_parse_context_push : no return generator

// Unsupported : g_markup_parse_context_unref : no return generator

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

func (recv *MarkupParser) toC() *C.GMarkupParser {

	return recv.native
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

func (recv *MatchInfo) toC() *C.GMatchInfo {

	return recv.native
}

// Unsupported : g_match_info_fetch_all : no return type

// Unsupported : g_match_info_fetch_named_pos : unsupported parameter start_pos : no type generator for gint, gint*

// Unsupported : g_match_info_fetch_pos : unsupported parameter start_pos : no type generator for gint, gint*

// Unsupported : g_match_info_free : no return generator

// Unsupported : g_match_info_is_partial_match : no return generator

// Unsupported : g_match_info_matches : no return generator

// Unsupported : g_match_info_next : no return generator

// Unsupported : g_match_info_unref : no return generator

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

func (recv *MemVTable) toC() *C.GMemVTable {

	return recv.native
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native *C.GNode
	Data   uintptr
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

	g := &Node{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *Node) toC() *C.GNode {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return recv.native
}

// ChildIndex is a wrapper around the C function g_node_child_index.
func (recv *Node) ChildIndex(data uintptr) int32 {
	c_data := (C.gpointer)(data)

	retC := C.g_node_child_index((*C.GNode)(recv.native), c_data)
	retGo := (int32)(retC)

	return retGo
}

// ChildPosition is a wrapper around the C function g_node_child_position.
func (recv *Node) ChildPosition(child *Node) int32 {
	c_child := child.toC()

	retC := C.g_node_child_position((*C.GNode)(recv.native), c_child)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_node_children_foreach : unsupported parameter func : no type generator for NodeForeachFunc, GNodeForeachFunc

// Copy is a wrapper around the C function g_node_copy.
func (recv *Node) Copy() *Node {
	retC := C.g_node_copy((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc, GCopyFunc

// Depth is a wrapper around the C function g_node_depth.
func (recv *Node) Depth() uint32 {
	retC := C.g_node_depth((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_node_destroy : no return generator

// Find is a wrapper around the C function g_node_find.
func (recv *Node) Find(order TraverseType, flags TraverseFlags, data uintptr) *Node {
	c_order := (C.GTraverseType)(order)

	c_flags := (C.GTraverseFlags)(flags)

	c_data := (C.gpointer)(data)

	retC := C.g_node_find((*C.GNode)(recv.native), c_order, c_flags, c_data)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindChild is a wrapper around the C function g_node_find_child.
func (recv *Node) FindChild(flags TraverseFlags, data uintptr) *Node {
	c_flags := (C.GTraverseFlags)(flags)

	c_data := (C.gpointer)(data)

	retC := C.g_node_find_child((*C.GNode)(recv.native), c_flags, c_data)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

	c_node := node.toC()

	retC := C.g_node_insert((*C.GNode)(recv.native), c_position, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertAfter is a wrapper around the C function g_node_insert_after.
func (recv *Node) InsertAfter(sibling *Node, node *Node) *Node {
	c_sibling := sibling.toC()

	c_node := node.toC()

	retC := C.g_node_insert_after((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertBefore is a wrapper around the C function g_node_insert_before.
func (recv *Node) InsertBefore(sibling *Node, node *Node) *Node {
	c_sibling := sibling.toC()

	c_node := node.toC()

	retC := C.g_node_insert_before((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_node_is_ancestor : no return generator

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
	c_node := node.toC()

	retC := C.g_node_prepend((*C.GNode)(recv.native), c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_node_reverse_children : no return generator

// Unsupported : g_node_traverse : unsupported parameter func : no type generator for NodeTraverseFunc, GNodeTraverseFunc

// Unsupported : g_node_unlink : no return generator

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

func (recv *OptionContext) toC() *C.GOptionContext {

	return recv.native
}

// Unsupported : g_option_context_add_group : no return generator

// Unsupported : g_option_context_add_main_entries : no return generator

// Unsupported : g_option_context_free : no return generator

// Unsupported : g_option_context_get_help : unsupported parameter main_help : no type generator for gboolean, gboolean

// Unsupported : g_option_context_get_help_enabled : no return generator

// Unsupported : g_option_context_get_ignore_unknown_options : no return generator

// Unsupported : g_option_context_get_strict_posix : no return generator

// Unsupported : g_option_context_parse : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : g_option_context_parse_strv : unsupported parameter arguments : no param type

// Unsupported : g_option_context_set_description : no return generator

// Unsupported : g_option_context_set_help_enabled : unsupported parameter help_enabled : no type generator for gboolean, gboolean

// Unsupported : g_option_context_set_ignore_unknown_options : unsupported parameter ignore_unknown : no type generator for gboolean, gboolean

// Unsupported : g_option_context_set_main_group : no return generator

// Unsupported : g_option_context_set_strict_posix : unsupported parameter strict_posix : no type generator for gboolean, gboolean

// Unsupported : g_option_context_set_summary : no return generator

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// Unsupported : g_option_context_set_translation_domain : no return generator

// OptionEntry is a wrapper around the C record GOptionEntry.
type OptionEntry struct {
	native         *C.GOptionEntry
	LongName       string
	ShortName      rune
	Flags          int32
	Arg            OptionArg
	ArgData        uintptr
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
		ArgData:        (uintptr)(c.arg_data),
		ArgDescription: C.GoString(c.arg_description),
		Description:    C.GoString(c.description),
		Flags:          (int32)(c.flags),
		LongName:       C.GoString(c.long_name),
		ShortName:      (rune)(c.short_name),
		native:         c,
	}

	return g
}

func (recv *OptionEntry) toC() *C.GOptionEntry {
	recv.native.long_name =
		C.CString(recv.LongName)
	recv.native.short_name =
		(C.gchar)(recv.ShortName)
	recv.native.flags =
		(C.gint)(recv.Flags)
	recv.native.arg =
		(C.GOptionArg)(recv.Arg)
	recv.native.arg_data =
		(C.gpointer)(recv.ArgData)
	recv.native.description =
		C.CString(recv.Description)
	recv.native.arg_description =
		C.CString(recv.ArgDescription)

	return recv.native
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

func (recv *OptionGroup) toC() *C.GOptionGroup {

	return recv.native
}

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_option_group_add_entries : no return generator

// Unsupported : g_option_group_free : no return generator

// Unsupported : g_option_group_set_error_hook : unsupported parameter error_func : no type generator for OptionErrorFunc, GOptionErrorFunc

// Unsupported : g_option_group_set_parse_hooks : unsupported parameter pre_parse_func : no type generator for OptionParseFunc, GOptionParseFunc

// Unsupported : g_option_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// Unsupported : g_option_group_set_translation_domain : no return generator

// Unsupported : g_option_group_unref : no return generator

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

func (recv *PatternSpec) toC() *C.GPatternSpec {

	return recv.native
}

// Unsupported : g_pattern_spec_equal : no return generator

// Unsupported : g_pattern_spec_free : no return generator

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

func (recv *PollFD) toC() *C.GPollFD {
	recv.native.fd =
		(C.gint)(recv.Fd)
	recv.native.events =
		(C.gushort)(recv.Events)
	recv.native.revents =
		(C.gushort)(recv.Revents)

	return recv.native
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

func (recv *Private) toC() *C.GPrivate {

	return recv.native
}

// Get is a wrapper around the C function g_private_get.
func (recv *Private) Get() uintptr {
	retC := C.g_private_get((*C.GPrivate)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_private_replace : no return generator

// Unsupported : g_private_set : no return generator

// PtrArray is a wrapper around the C record GPtrArray.
type PtrArray struct {
	native *C.GPtrArray
	// pdata : no type generator for gpointer, gpointer*
	Len uint32
}

func PtrArrayNewFromC(u unsafe.Pointer) *PtrArray {
	c := (*C.GPtrArray)(u)
	if c == nil {
		return nil
	}

	g := &PtrArray{
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
}

func (recv *PtrArray) toC() *C.GPtrArray {
	recv.native.len =
		(C.guint)(recv.Len)

	return recv.native
}

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

func (recv *Queue) toC() *C.GQueue {
	recv.native.length =
		(C.guint)(recv.Length)

	return recv.native
}

// Unsupported : g_queue_clear : no return generator

// Unsupported : g_queue_delete_link : no return generator

// Unsupported : g_queue_find_custom : unsupported parameter func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func, GFunc

// Unsupported : g_queue_free : no return generator

// Unsupported : g_queue_free_full : unsupported parameter free_func : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_queue_init : no return generator

// Unsupported : g_queue_insert_after : no return generator

// Unsupported : g_queue_insert_before : no return generator

// Unsupported : g_queue_insert_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_queue_is_empty : no return generator

// PeekHead is a wrapper around the C function g_queue_peek_head.
func (recv *Queue) PeekHead() uintptr {
	retC := C.g_queue_peek_head((*C.GQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// PeekTail is a wrapper around the C function g_queue_peek_tail.
func (recv *Queue) PeekTail() uintptr {
	retC := C.g_queue_peek_tail((*C.GQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// PopHead is a wrapper around the C function g_queue_pop_head.
func (recv *Queue) PopHead() uintptr {
	retC := C.g_queue_pop_head((*C.GQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// PopHeadLink is a wrapper around the C function g_queue_pop_head_link.
func (recv *Queue) PopHeadLink() *List {
	retC := C.g_queue_pop_head_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopTail is a wrapper around the C function g_queue_pop_tail.
func (recv *Queue) PopTail() uintptr {
	retC := C.g_queue_pop_tail((*C.GQueue)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// PopTailLink is a wrapper around the C function g_queue_pop_tail_link.
func (recv *Queue) PopTailLink() *List {
	retC := C.g_queue_pop_tail_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_push_head : no return generator

// Unsupported : g_queue_push_head_link : no return generator

// Unsupported : g_queue_push_nth : no return generator

// Unsupported : g_queue_push_nth_link : no return generator

// Unsupported : g_queue_push_tail : no return generator

// Unsupported : g_queue_push_tail_link : no return generator

// Unsupported : g_queue_remove : no return generator

// Unsupported : g_queue_reverse : no return generator

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_queue_unlink : no return generator

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

func (recv *Rand) toC() *C.GRand {

	return recv.native
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

// Unsupported : g_rand_free : no return generator

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

// Unsupported : g_rand_set_seed : no return generator

// Unsupported : g_rand_set_seed_array : unsupported parameter seed : no type generator for guint32, const guint32*

// SList is a wrapper around the C record GSList.
type SList struct {
	native *C.GSList
	Data   uintptr
	// next : record
}

func SListNewFromC(u unsafe.Pointer) *SList {
	c := (*C.GSList)(u)
	if c == nil {
		return nil
	}

	g := &SList{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *SList) toC() *C.GSList {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return recv.native
}

// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native         *C.GScanner
	UserData       uintptr
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
		UserData:       (uintptr)(c.user_data),
		native:         c,
	}

	return g
}

func (recv *Scanner) toC() *C.GScanner {
	recv.native.user_data =
		(C.gpointer)(recv.UserData)
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

	return recv.native
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

// Unsupported : g_scanner_destroy : no return generator

// Unsupported : g_scanner_eof : no return generator

// Unsupported : g_scanner_error : unsupported parameter ... : varargs

// GetNextToken is a wrapper around the C function g_scanner_get_next_token.
func (recv *Scanner) GetNextToken() TokenType {
	retC := C.g_scanner_get_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Unsupported : g_scanner_input_file : no return generator

// Unsupported : g_scanner_input_text : no return generator

// LookupSymbol is a wrapper around the C function g_scanner_lookup_symbol.
func (recv *Scanner) LookupSymbol(symbol string) uintptr {
	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	retC := C.g_scanner_lookup_symbol((*C.GScanner)(recv.native), c_symbol)
	retGo := (uintptr)(retC)

	return retGo
}

// PeekNextToken is a wrapper around the C function g_scanner_peek_next_token.
func (recv *Scanner) PeekNextToken() TokenType {
	retC := C.g_scanner_peek_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Unsupported : g_scanner_scope_add_symbol : no return generator

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc, GHFunc

// ScopeLookupSymbol is a wrapper around the C function g_scanner_scope_lookup_symbol.
func (recv *Scanner) ScopeLookupSymbol(scopeId uint32, symbol string) uintptr {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	retC := C.g_scanner_scope_lookup_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_scanner_scope_remove_symbol : no return generator

// SetScope is a wrapper around the C function g_scanner_set_scope.
func (recv *Scanner) SetScope(scopeId uint32) uint32 {
	c_scope_id := (C.guint)(scopeId)

	retC := C.g_scanner_set_scope((*C.GScanner)(recv.native), c_scope_id)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_scanner_sync_file_offset : no return generator

// Unsupported : g_scanner_unexp_token : no return generator

// Unsupported : g_scanner_warn : unsupported parameter ... : varargs

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

func (recv *ScannerConfig) toC() *C.GScannerConfig {
	recv.native.cset_skip_characters =
		C.CString(recv.CsetSkipCharacters)
	recv.native.cset_identifier_first =
		C.CString(recv.CsetIdentifierFirst)
	recv.native.cset_identifier_nth =
		C.CString(recv.CsetIdentifierNth)
	recv.native.cpair_comment_single =
		C.CString(recv.CpairCommentSingle)

	return recv.native
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

func (recv *Sequence) toC() *C.GSequence {

	return recv.native
}

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func, GFunc

// Unsupported : g_sequence_free : no return generator

// Unsupported : g_sequence_insert_sorted : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_is_empty : no return generator

// Unsupported : g_sequence_lookup : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_lookup_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_search : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_search_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

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

func (recv *SequenceIter) toC() *C.GSequenceIter {

	return recv.native
}

// Unsupported : g_sequence_iter_is_begin : no return generator

// Unsupported : g_sequence_iter_is_end : no return generator

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

func (recv *Source) toC() *C.GSource {

	return recv.native
}

// SourceNew is a wrapper around the C function g_source_new.
func SourceNew(sourceFuncs *SourceFuncs, structSize uint32) *Source {
	c_source_funcs := sourceFuncs.toC()

	c_struct_size := (C.guint)(structSize)

	retC := C.g_source_new(c_source_funcs, c_struct_size)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_source_add_child_source : no return generator

// Unsupported : g_source_add_poll : no return generator

// Attach is a wrapper around the C function g_source_attach.
func (recv *Source) Attach(context *MainContext) uint32 {
	c_context := context.toC()

	retC := C.g_source_attach((*C.GSource)(recv.native), c_context)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_source_destroy : no return generator

// Unsupported : g_source_get_can_recurse : no return generator

// GetContext is a wrapper around the C function g_source_get_context.
func (recv *Source) GetContext() *MainContext {
	retC := C.g_source_get_context((*C.GSource)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_source_get_current_time : no return generator

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

// Unsupported : g_source_is_destroyed : no return generator

// Unsupported : g_source_modify_unix_fd : no return generator

// Ref is a wrapper around the C function g_source_ref.
func (recv *Source) Ref() *Source {
	retC := C.g_source_ref((*C.GSource)(recv.native))
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_source_remove_child_source : no return generator

// Unsupported : g_source_remove_poll : no return generator

// Unsupported : g_source_remove_unix_fd : no return generator

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_source_set_callback_indirect : no return generator

// Unsupported : g_source_set_can_recurse : unsupported parameter can_recurse : no type generator for gboolean, gboolean

// Unsupported : g_source_set_funcs : no return generator

// Unsupported : g_source_set_name : no return generator

// Unsupported : g_source_set_priority : no return generator

// Unsupported : g_source_set_ready_time : no return generator

// Unsupported : g_source_unref : no return generator

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

func (recv *SourceCallbackFuncs) toC() *C.GSourceCallbackFuncs {

	return recv.native
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

func (recv *SourceFuncs) toC() *C.GSourceFuncs {

	return recv.native
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

func (recv *SourcePrivate) toC() *C.GSourcePrivate {

	return recv.native
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

func (recv *StatBuf) toC() *C.GStatBuf {

	return recv.native
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

func (recv *String) toC() *C.GString {
	recv.native.str =
		C.CString(recv.Str)
	recv.native.len =
		(C.gsize)(recv.Len)
	recv.native.allocated_len =
		(C.gsize)(recv.AllocatedLen)

	return recv.native
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

// Unsupported : g_string_append_printf : unsupported parameter ... : varargs

// AppendUnichar is a wrapper around the C function g_string_append_unichar.
func (recv *String) AppendUnichar(wc rune) *String {
	c_wc := (C.gunichar)(wc)

	retC := C.g_string_append_unichar((*C.GString)(recv.native), c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_append_uri_escaped : unsupported parameter allow_utf8 : no type generator for gboolean, gboolean

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list, va_list

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

// Unsupported : g_string_equal : no return generator

// Erase is a wrapper around the C function g_string_erase.
func (recv *String) Erase(pos int64, len int64) *String {
	c_pos := (C.gssize)(pos)

	c_len := (C.gssize)(len)

	retC := C.g_string_erase((*C.GString)(recv.native), c_pos, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_free : unsupported parameter free_segment : no type generator for gboolean, gboolean

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

// Unsupported : g_string_printf : unsupported parameter ... : varargs

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

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list, va_list

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

func (recv *StringChunk) toC() *C.GStringChunk {

	return recv.native
}

// Unsupported : g_string_chunk_clear : no return generator

// Unsupported : g_string_chunk_free : no return generator

// Insert is a wrapper around the C function g_string_chunk_insert.
func (recv *StringChunk) Insert(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_string_chunk_insert((*C.GStringChunk)(recv.native), c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// InsertConst is a wrapper around the C function g_string_chunk_insert_const.
func (recv *StringChunk) InsertConst(string string) string {
	c_string := C.CString(string)
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

func (recv *TestCase) toC() *C.GTestCase {

	return recv.native
}

// TestConfig is a wrapper around the C record GTestConfig.
type TestConfig struct {
	native *C.GTestConfig
	// test_initialized : no type generator for gboolean, gboolean
	// test_quick : no type generator for gboolean, gboolean
	// test_perf : no type generator for gboolean, gboolean
	// test_verbose : no type generator for gboolean, gboolean
	// test_quiet : no type generator for gboolean, gboolean
	// test_undefined : no type generator for gboolean, gboolean
}

func TestConfigNewFromC(u unsafe.Pointer) *TestConfig {
	c := (*C.GTestConfig)(u)
	if c == nil {
		return nil
	}

	g := &TestConfig{native: c}

	return g
}

func (recv *TestConfig) toC() *C.GTestConfig {

	return recv.native
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

func (recv *TestSuite) toC() *C.GTestSuite {

	return recv.native
}

// Unsupported : g_test_suite_add : no return generator

// Unsupported : g_test_suite_add_suite : no return generator

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

func (recv *Thread) toC() *C.GThread {

	return recv.native
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Join is a wrapper around the C function g_thread_join.
func (recv *Thread) Join() uintptr {
	retC := C.g_thread_join((*C.GThread)(recv.native))
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_thread_unref : no return generator

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	UserData uintptr
	// exclusive : no type generator for gboolean, gboolean
}

func ThreadPoolNewFromC(u unsafe.Pointer) *ThreadPool {
	c := (*C.GThreadPool)(u)
	if c == nil {
		return nil
	}

	g := &ThreadPool{
		UserData: (uintptr)(c.user_data),
		native:   c,
	}

	return g
}

func (recv *ThreadPool) toC() *C.GThreadPool {
	recv.native.user_data =
		(C.gpointer)(recv.UserData)

	return recv.native
}

// Unsupported : g_thread_pool_free : unsupported parameter immediate : no type generator for gboolean, gboolean

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

// Unsupported : g_thread_pool_move_to_front : no return generator

// Unsupported : g_thread_pool_push : no return generator

// Unsupported : g_thread_pool_set_max_threads : no return generator

// Unsupported : g_thread_pool_set_sort_function : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

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

func (recv *TimeVal) toC() *C.GTimeVal {
	recv.native.tv_sec =
		(C.glong)(recv.TvSec)
	recv.native.tv_usec =
		(C.glong)(recv.TvUsec)

	return recv.native
}

// Unsupported : g_time_val_add : no return generator

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

func (recv *Timer) toC() *C.GTimer {

	return recv.native
}

// Unsupported : g_timer_continue : no return generator

// Unsupported : g_timer_destroy : no return generator

// Unsupported : g_timer_elapsed : unsupported parameter microseconds : no type generator for gulong, gulong*

// Unsupported : g_timer_reset : no return generator

// Unsupported : g_timer_start : no return generator

// Unsupported : g_timer_stop : no return generator

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

func (recv *TrashStack) toC() *C.GTrashStack {

	return recv.native
}

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

func (recv *Tree) toC() *C.GTree {

	return recv.native
}

// Unsupported : g_tree_destroy : no return generator

// Unsupported : g_tree_foreach : unsupported parameter func : no type generator for TraverseFunc, GTraverseFunc

// Height is a wrapper around the C function g_tree_height.
func (recv *Tree) Height() int32 {
	retC := C.g_tree_height((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_tree_insert : no return generator

// Lookup is a wrapper around the C function g_tree_lookup.
func (recv *Tree) Lookup(key uintptr) uintptr {
	c_key := (C.gconstpointer)(key)

	retC := C.g_tree_lookup((*C.GTree)(recv.native), c_key)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_tree_lookup_extended : unsupported parameter orig_key : no type generator for gpointer, gpointer*

// Nnodes is a wrapper around the C function g_tree_nnodes.
func (recv *Tree) Nnodes() int32 {
	retC := C.g_tree_nnodes((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_tree_remove : no return generator

// Unsupported : g_tree_replace : no return generator

// Unsupported : g_tree_search : unsupported parameter search_func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_tree_steal : no return generator

// Unsupported : g_tree_traverse : unsupported parameter traverse_func : no type generator for TraverseFunc, GTraverseFunc

// Unsupported : g_tree_unref : no return generator

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

func (recv *VariantBuilder) toC() *C.GVariantBuilder {

	return recv.native
}

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_add : unsupported parameter ... : varargs

// Unsupported : g_variant_builder_add_parsed : unsupported parameter ... : varargs

// Unsupported : g_variant_builder_add_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_builder_clear : no return generator

// Unsupported : g_variant_builder_close : no return generator

// Unsupported : g_variant_builder_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_builder_init : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_open : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_unref : no return generator

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

func (recv *VariantIter) toC() *C.GVariantIter {

	return recv.native
}

// Unsupported : g_variant_iter_free : no return generator

// Unsupported : g_variant_iter_init : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_iter_loop : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next_value : return type : Blacklisted record : GVariant

// Blacklisted : GVariantType

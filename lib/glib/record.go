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

// Blacklisted : g_async_queue_new

// Blacklisted : g_async_queue_length

// Blacklisted : g_async_queue_length_unlocked

// Blacklisted : g_async_queue_lock

// Unsupported : g_async_queue_pop : no return generator

// Unsupported : g_async_queue_pop_unlocked : no return generator

// Unsupported : g_async_queue_push : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_async_queue_push_unlocked : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_async_queue_ref

// Blacklisted : g_async_queue_ref_unlocked

// Unsupported : g_async_queue_timed_pop : no return generator

// Unsupported : g_async_queue_timed_pop_unlocked : no return generator

// Unsupported : g_async_queue_timeout_pop : no return generator

// Unsupported : g_async_queue_timeout_pop_unlocked : no return generator

// Unsupported : g_async_queue_try_pop : no return generator

// Unsupported : g_async_queue_try_pop_unlocked : no return generator

// Blacklisted : g_async_queue_unlock

// Blacklisted : g_async_queue_unref

// Blacklisted : g_async_queue_unref_and_unlock

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

// Blacklisted : g_bookmark_file_error_quark

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

// Blacklisted : g_cond_broadcast

// Blacklisted : g_cond_signal

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

// Blacklisted : g_dir_open

// Blacklisted : g_dir_close

// Blacklisted : g_dir_read_name

// Blacklisted : g_dir_rewind

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

// Blacklisted : g_hash_table_destroy

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
// Blacklisted : g_hash_table_size

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

// Blacklisted : g_hook_alloc

// Blacklisted : g_hook_destroy

// Blacklisted : g_hook_destroy_link

// g_hook_find : unsupported parameter func : no type generator for HookFindFunc (GHookFindFunc) for param func
// g_hook_find_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_hook_find_func : unsupported parameter func : no type generator for gpointer (gpointer) for param func
// g_hook_find_func_data : unsupported parameter func : no type generator for gpointer (gpointer) for param func
// Blacklisted : g_hook_first_valid

// Blacklisted : g_hook_free

// Blacklisted : g_hook_get

// Blacklisted : g_hook_insert_before

// g_hook_insert_sorted : unsupported parameter func : no type generator for HookCompareFunc (GHookCompareFunc) for param func
// Blacklisted : g_hook_next_valid

// Blacklisted : g_hook_prepend

// Blacklisted : g_hook_ref

// Blacklisted : g_hook_unref

// Blacklisted : g_hook_compare_ids

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

// Blacklisted : g_hook_list_clear

// Init is a wrapper around the C function g_hook_list_init.
func (recv *HookList) Init(hookSize uint32) {
	c_hook_size := (C.guint)(hookSize)

	C.g_hook_list_init((*C.GHookList)(recv.native), c_hook_size)

	return
}

// Blacklisted : g_hook_list_invoke

// Blacklisted : g_hook_list_invoke_check

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

// Blacklisted : g_io_channel_new_file

// Blacklisted : g_io_channel_unix_new

// Blacklisted : g_io_channel_error_from_errno

// Blacklisted : g_io_channel_error_quark

// Blacklisted : g_io_channel_close

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

// Blacklisted : g_io_channel_get_buffer_condition

// Blacklisted : g_io_channel_get_buffer_size

// Blacklisted : g_io_channel_get_buffered

// Blacklisted : g_io_channel_get_close_on_unref

// Blacklisted : g_io_channel_get_encoding

// Blacklisted : g_io_channel_get_flags

// Blacklisted : g_io_channel_get_line_term

// Init is a wrapper around the C function g_io_channel_init.
func (recv *IOChannel) Init() {
	C.g_io_channel_init((*C.GIOChannel)(recv.native))

	return
}

// Blacklisted : g_io_channel_read

// Unsupported : g_io_channel_read_chars : unsupported parameter buf : output array param buf

// Blacklisted : g_io_channel_read_line

// Blacklisted : g_io_channel_read_line_string

// Unsupported : g_io_channel_read_to_end : unsupported parameter str_return : output array param str_return

// Blacklisted : g_io_channel_read_unichar

// Blacklisted : g_io_channel_ref

// Blacklisted : g_io_channel_seek

// Blacklisted : g_io_channel_seek_position

// Blacklisted : g_io_channel_set_buffer_size

// Blacklisted : g_io_channel_set_buffered

// Blacklisted : g_io_channel_set_close_on_unref

// Blacklisted : g_io_channel_set_encoding

// Blacklisted : g_io_channel_set_flags

// Blacklisted : g_io_channel_set_line_term

// Blacklisted : g_io_channel_shutdown

// Blacklisted : g_io_channel_unix_get_fd

// Blacklisted : g_io_channel_unref

// Blacklisted : g_io_channel_write

// Blacklisted : g_io_channel_write_chars

// Blacklisted : g_io_channel_write_unichar

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

// Blacklisted : g_key_file_error_quark

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

// Blacklisted : g_list_alloc

// g_list_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Blacklisted : g_list_concat

// Blacklisted : g_list_copy

// Blacklisted : g_list_delete_link

// g_list_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// Blacklisted : g_list_first

// g_list_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func
// Blacklisted : g_list_free

// Blacklisted : g_list_free_1

// g_list_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_insert : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Blacklisted : g_list_last

// Blacklisted : g_list_length

// Blacklisted : g_list_nth

// g_list_nth_data : no return generator
// Blacklisted : g_list_nth_prev

// Blacklisted : g_list_position

// g_list_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// Blacklisted : g_list_remove_link

// Blacklisted : g_list_reverse

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

// Blacklisted : g_main_context_new

// Blacklisted : g_main_context_default

// Blacklisted : g_main_context_acquire

// Blacklisted : g_main_context_add_poll

// Unsupported : g_main_context_check : unsupported parameter fds :

// Blacklisted : g_main_context_dispatch

// Unsupported : g_main_context_find_source_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Blacklisted : g_main_context_find_source_by_id

// Unsupported : g_main_context_find_source_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_main_context_get_poll_func : no return generator

// Blacklisted : g_main_context_iteration

// Blacklisted : g_main_context_pending

// Blacklisted : g_main_context_prepare

// Unsupported : g_main_context_query : unsupported parameter fds : output array param fds

// Blacklisted : g_main_context_ref

// Blacklisted : g_main_context_release

// Blacklisted : g_main_context_remove_poll

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc (GPollFunc) for param func

// Blacklisted : g_main_context_unref

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// Blacklisted : g_main_context_wakeup

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

// Blacklisted : g_main_loop_new

// Blacklisted : g_main_loop_get_context

// Blacklisted : g_main_loop_is_running

// Blacklisted : g_main_loop_quit

// Blacklisted : g_main_loop_ref

// Blacklisted : g_main_loop_run

// Blacklisted : g_main_loop_unref

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

// Blacklisted : g_mapped_file_unref

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

// Blacklisted : g_markup_parse_context_end_parse

// Blacklisted : g_markup_parse_context_free

// Blacklisted : g_markup_parse_context_get_position

// Blacklisted : g_markup_parse_context_parse

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

// Blacklisted : g_node_child_position

// Unsupported : g_node_children_foreach : unsupported parameter func : no type generator for NodeForeachFunc (GNodeForeachFunc) for param func

// Blacklisted : g_node_copy

// Blacklisted : g_node_depth

// Destroy is a wrapper around the C function g_node_destroy.
func (recv *Node) Destroy() {
	C.g_node_destroy((*C.GNode)(recv.native))

	return
}

// Unsupported : g_node_find : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_node_find_child : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_node_first_sibling

// Blacklisted : g_node_get_root

// Blacklisted : g_node_insert

// Blacklisted : g_node_insert_after

// Blacklisted : g_node_insert_before

// Blacklisted : g_node_is_ancestor

// Blacklisted : g_node_last_child

// Blacklisted : g_node_last_sibling

// Blacklisted : g_node_max_height

// Blacklisted : g_node_n_children

// Blacklisted : g_node_n_nodes

// Blacklisted : g_node_nth_child

// Blacklisted : g_node_prepend

// Blacklisted : g_node_reverse_children

// Unsupported : g_node_traverse : unsupported parameter func : no type generator for NodeTraverseFunc (GNodeTraverseFunc) for param func

// Blacklisted : g_node_unlink

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

// Blacklisted : g_pattern_spec_new

// Blacklisted : g_pattern_spec_equal

// Blacklisted : g_pattern_spec_free

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

// Blacklisted : g_queue_new

// Blacklisted : g_queue_free

// Blacklisted : g_queue_is_empty

// Unsupported : g_queue_peek_head : no return generator

// Unsupported : g_queue_peek_tail : no return generator

// Unsupported : g_queue_pop_head : no return generator

// Blacklisted : g_queue_pop_head_link

// Unsupported : g_queue_pop_tail : no return generator

// Blacklisted : g_queue_pop_tail_link

// Unsupported : g_queue_push_head : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_queue_push_head_link

// Unsupported : g_queue_push_tail : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_queue_push_tail_link

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

// Blacklisted : g_rand_new

// Blacklisted : g_rand_new_with_seed

// Blacklisted : g_rand_double

// Blacklisted : g_rand_double_range

// Blacklisted : g_rand_free

// Blacklisted : g_rand_int

// Blacklisted : g_rand_int_range

// Blacklisted : g_rand_set_seed

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

// Blacklisted : g_scanner_new

// Blacklisted : g_scanner_cur_line

// Blacklisted : g_scanner_cur_position

// Blacklisted : g_scanner_cur_token

// Unsupported : g_scanner_cur_value : no return generator

// Destroy is a wrapper around the C function g_scanner_destroy.
func (recv *Scanner) Destroy() {
	C.g_scanner_destroy((*C.GScanner)(recv.native))

	return
}

// Blacklisted : g_scanner_eof

// Blacklisted : g_scanner_error

// Blacklisted : g_scanner_get_next_token

// Blacklisted : g_scanner_input_file

// Blacklisted : g_scanner_input_text

// Unsupported : g_scanner_lookup_symbol : no return generator

// Blacklisted : g_scanner_peek_next_token

// Unsupported : g_scanner_scope_add_symbol : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc (GHFunc) for param func

// Unsupported : g_scanner_scope_lookup_symbol : no return generator

// Blacklisted : g_scanner_scope_remove_symbol

// Blacklisted : g_scanner_set_scope

// Blacklisted : g_scanner_sync_file_offset

// Blacklisted : g_scanner_unexp_token

// Blacklisted : g_scanner_warn

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

// Blacklisted : g_source_new

// Blacklisted : g_source_remove

// g_source_remove_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data
// g_source_remove_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data
// Blacklisted : g_source_add_poll

// Blacklisted : g_source_attach

// Destroy is a wrapper around the C function g_source_destroy.
func (recv *Source) Destroy() {
	C.g_source_destroy((*C.GSource)(recv.native))

	return
}

// Blacklisted : g_source_get_can_recurse

// Blacklisted : g_source_get_context

// Blacklisted : g_source_get_current_time

// Blacklisted : g_source_get_id

// Blacklisted : g_source_get_priority

// Blacklisted : g_source_get_ready_time

// Blacklisted : g_source_ref

// Blacklisted : g_source_remove_poll

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc (GSourceFunc) for param func

// Unsupported : g_source_set_callback_indirect : unsupported parameter callback_data : no type generator for gpointer (gpointer) for param callback_data

// Blacklisted : g_source_set_can_recurse

// Blacklisted : g_source_set_priority

// Blacklisted : g_source_unref

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

// Blacklisted : g_string_append

// Blacklisted : g_string_append_c

// Blacklisted : g_string_append_len

// Blacklisted : g_string_append_printf

// Blacklisted : g_string_append_unichar

// Blacklisted : g_string_ascii_down

// Blacklisted : g_string_ascii_up

// Blacklisted : g_string_assign

// Blacklisted : g_string_down

// Blacklisted : g_string_equal

// Blacklisted : g_string_erase

// Blacklisted : g_string_free

// Blacklisted : g_string_hash

// Blacklisted : g_string_insert

// Blacklisted : g_string_insert_c

// Blacklisted : g_string_insert_len

// Blacklisted : g_string_insert_unichar

// Blacklisted : g_string_prepend

// Blacklisted : g_string_prepend_c

// Blacklisted : g_string_prepend_len

// Blacklisted : g_string_prepend_unichar

// Printf is a wrapper around the C function g_string_printf.
func (recv *String) Printf(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_string_printf((*C.GString)(recv.native), c_format)

	return
}

// Blacklisted : g_string_set_size

// Blacklisted : g_string_truncate

// Blacklisted : g_string_up

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

// Blacklisted : g_string_chunk_new

// Blacklisted : g_string_chunk_free

// Blacklisted : g_string_chunk_insert

// Blacklisted : g_string_chunk_insert_const

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

// Blacklisted : g_thread_error_quark

// g_thread_exit : unsupported parameter retval : no type generator for gpointer (gpointer) for param retval
// Blacklisted : g_thread_self

// Blacklisted : g_thread_yield

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

// Blacklisted : g_thread_pool_get_max_unused_threads

// Blacklisted : g_thread_pool_get_num_unused_threads

// g_thread_pool_new : unsupported parameter func : no type generator for Func (GFunc) for param func
// Blacklisted : g_thread_pool_set_max_unused_threads

// Blacklisted : g_thread_pool_stop_unused_threads

// Blacklisted : g_thread_pool_free

// Blacklisted : g_thread_pool_get_max_threads

// Blacklisted : g_thread_pool_get_num_threads

// Unsupported : g_thread_pool_push : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_thread_pool_set_max_threads

// Blacklisted : g_thread_pool_unprocessed

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

// Blacklisted : g_time_val_add

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

// Blacklisted : g_timer_new

// Destroy is a wrapper around the C function g_timer_destroy.
func (recv *Timer) Destroy() {
	C.g_timer_destroy((*C.GTimer)(recv.native))

	return
}

// Blacklisted : g_timer_elapsed

// Blacklisted : g_timer_reset

// Blacklisted : g_timer_start

// Blacklisted : g_timer_stop

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

// Blacklisted : g_trash_stack_height

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

// Blacklisted : g_tree_height

// Unsupported : g_tree_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_tree_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_tree_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key

// Blacklisted : g_tree_nnodes

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

// Blacklisted : g_variant_type_new_array

// Blacklisted : g_variant_type_new_dict_entry

// Blacklisted : g_variant_type_new_maybe

// Unsupported : g_variant_type_new_tuple : unsupported parameter items :

// Blacklisted : g_variant_type_checked_

// Blacklisted : g_variant_type_string_is_valid

// Blacklisted : g_variant_type_copy

// Blacklisted : g_variant_type_dup_string

// Blacklisted : g_variant_type_element

// Blacklisted : g_variant_type_equal

// Blacklisted : g_variant_type_first

// Blacklisted : g_variant_type_free

// Blacklisted : g_variant_type_get_string_length

// Blacklisted : g_variant_type_hash

// Blacklisted : g_variant_type_is_array

// Blacklisted : g_variant_type_is_basic

// Blacklisted : g_variant_type_is_container

// Blacklisted : g_variant_type_is_definite

// Blacklisted : g_variant_type_is_dict_entry

// Blacklisted : g_variant_type_is_maybe

// Blacklisted : g_variant_type_is_subtype_of

// Blacklisted : g_variant_type_is_tuple

// Blacklisted : g_variant_type_is_variant

// Blacklisted : g_variant_type_key

// Blacklisted : g_variant_type_n_items

// Blacklisted : g_variant_type_next

// Blacklisted : g_variant_type_peek_string

// Blacklisted : g_variant_type_value

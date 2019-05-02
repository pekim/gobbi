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
import "C"

// Unsupported : alias has no type generator for none, void

// Blacklisted : KEY_FILE_DESKTOP_ACTION_GROUP_PREFIX

// Blacklisted : LOG_DOMAIN

// Blacklisted : WIN32_MSG_HANDLE

// Blacklisted : GTestResult

// Unsupported : g_assertion_message_cmpnum : unsupported parameter arg1 : no type generator for long double (long double) for param arg1

// Unsupported : g_atexit : unsupported parameter func : no type generator for VoidFunc (GVoidFunc) for param func

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Blacklisted : g_byte_array_free

// Unsupported : g_byte_array_new : array return type :

// Unsupported : g_convert : array return type :

// Unsupported : g_convert_with_fallback : array return type :

// Unsupported : g_convert_with_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_datalist_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// Unsupported : g_datalist_get_data : no return generator

// Unsupported : g_datalist_id_get_data : no return generator

// Unsupported : g_datalist_id_remove_no_notify : no return generator

// Unsupported : g_datalist_id_set_data_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_dataset_destroy : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_foreach : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_get_data : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_remove_no_notify : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_set_data_full : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_direct_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_direct_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_file_get_contents : unsupported parameter contents : output array param contents

// Unsupported : g_free : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_hash_table_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key

// Unsupported : g_hash_table_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_remove_by_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_int_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_int_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_io_add_watch : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// Unsupported : g_io_add_watch_full : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// Unsupported : g_locale_from_utf8 : array return type :

// Unsupported : g_log_default_handler : unsupported parameter unused_data : no type generator for gpointer (gpointer) for param unused_data

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Blacklisted : g_log_structured_standard

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_malloc : no return generator

// Unsupported : g_malloc0 : no return generator

// Unsupported : g_memdup : unsupported parameter mem : no type generator for gpointer (gconstpointer) for param mem

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : no type generator for gpointer (gpointer*) for param nullify_location

// Blacklisted : g_number_parser_error_quark

// Unsupported : g_parse_debug_string : unsupported parameter keys :

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_qsort_with_data : unsupported parameter pbase : no type generator for gpointer (gconstpointer) for param pbase

// Unsupported : g_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_shell_parse_argv : unsupported parameter argcp : array length param argcp is pointer (gint*)

// Blacklisted : g_slice_get_config_state

// Unsupported : g_source_remove_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_source_remove_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_spawn_async : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_spawn_async_with_pipes : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : output array param standard_output

// Unsupported : g_spawn_sync : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_str_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_str_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Unsupported : g_strdup_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_strdupv : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strfreev : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Blacklisted : g_strv_get_type

// Unsupported : g_thread_exit : unsupported parameter retval : no type generator for gpointer (gpointer) for param retval

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_trash_stack_peek : no return generator

// Unsupported : g_trash_stack_pop : no return generator

// Unsupported : g_trash_stack_push : unsupported parameter data_p : no type generator for gpointer (gpointer) for param data_p

// Unsupported : g_try_malloc : no return generator

// Unsupported : g_try_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Blacklisted : g_ucs4_to_utf16

// Blacklisted : g_unichar_to_utf8

// Blacklisted : g_unicode_canonical_decomposition

// Blacklisted : g_unix_error_quark

// Blacklisted : g_utf16_to_ucs4

// Blacklisted : g_utf8_to_ucs4

// Blacklisted : g_utf8_to_ucs4_fast

// Blacklisted : g_utf8_to_utf16

// Blacklisted : g_variant_get_gtype

// Unsupported : g_variant_parse : unsupported parameter endptr : in string with indirection level of 2

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

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

// MainContextNew is a wrapper around the C function g_main_context_new.
func MainContextNew() *MainContext {
	retC := C.g_main_context_new()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
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

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

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

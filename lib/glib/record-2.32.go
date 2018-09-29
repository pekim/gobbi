// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_async_queue_push_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_push_sorted_unlocked : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_sort : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_sort_unlocked : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_bookmark_file_get_added : no return generator

// Unsupported : g_bookmark_file_get_app_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : g_bookmark_file_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_modified : no return generator

// Unsupported : g_bookmark_file_get_uris : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_visited : no return generator

// Unsupported : g_bookmark_file_load_from_data : unsupported parameter data : no param type

// Unsupported : g_bookmark_file_set_added : unsupported parameter added : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_app_info : unsupported parameter stamp : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_groups : unsupported parameter groups : no param type

// Unsupported : g_bookmark_file_set_modified : unsupported parameter modified : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_visited : unsupported parameter visited : no type generator for glong, time_t

// Unsupported : g_bookmark_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

// Bytes is a wrapper around the C record GBytes.
type Bytes struct {
	native *C.GBytes
}

func BytesNewFromC(u unsafe.Pointer) *Bytes {
	c := (*C.GBytes)(u)
	if c == nil {
		return nil
	}

	g := &Bytes{native: c}

	return g
}

func (recv *Bytes) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_bytes_new : unsupported parameter data : no param type

// Unsupported : g_bytes_new_static : unsupported parameter data : no param type

// Unsupported : g_bytes_new_take : unsupported parameter data : no param type

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no param type

// Compare is a wrapper around the C function g_bytes_compare.
func (recv *Bytes) Compare(bytes2 uintptr) int32 {
	c_bytes2 := (C.gconstpointer)(bytes2)

	retC := C.g_bytes_compare((C.gconstpointer)(recv.native), c_bytes2)
	retGo := (int32)(retC)

	return retGo
}

// Equal is a wrapper around the C function g_bytes_equal.
func (recv *Bytes) Equal(bytes2 uintptr) bool {
	c_bytes2 := (C.gconstpointer)(bytes2)

	retC := C.g_bytes_equal((C.gconstpointer)(recv.native), c_bytes2)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_bytes_get_data : unsupported parameter size : no type generator for gsize, gsize*

// GetSize is a wrapper around the C function g_bytes_get_size.
func (recv *Bytes) GetSize() uint64 {
	retC := C.g_bytes_get_size((*C.GBytes)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Hash is a wrapper around the C function g_bytes_hash.
func (recv *Bytes) Hash() uint32 {
	retC := C.g_bytes_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NewFromBytes is a wrapper around the C function g_bytes_new_from_bytes.
func (recv *Bytes) NewFromBytes(offset uint64, length uint64) *Bytes {
	c_offset := (C.gsize)(offset)

	c_length := (C.gsize)(length)

	retC := C.g_bytes_new_from_bytes((*C.GBytes)(recv.native), c_offset, c_length)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_bytes_ref.
func (recv *Bytes) Ref() *Bytes {
	retC := C.g_bytes_ref((*C.GBytes)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_bytes_unref.
func (recv *Bytes) Unref() {
	C.g_bytes_unref((*C.GBytes)(recv.native))

	return
}

// Unsupported : g_bytes_unref_to_array : no return type

// Unsupported : g_bytes_unref_to_data : unsupported parameter size : no type generator for gsize, gsize*

// Unsupported : g_checksum_get_digest : unsupported parameter buffer : no type generator for guint8, guint8*

// Unsupported : g_checksum_update : unsupported parameter data : no param type

// Clear is a wrapper around the C function g_cond_clear.
func (recv *Cond) Clear() {
	C.g_cond_clear((*C.GCond)(recv.native))

	return
}

// Init is a wrapper around the C function g_cond_init.
func (recv *Cond) Init() {
	C.g_cond_init((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_date_set_time_t : unsupported parameter timet : no type generator for glong, time_t

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer, tm*

// Unsupported : g_date_time_get_ymd : unsupported parameter year : no type generator for gint, gint*

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_hash_table_iter_next : unsupported parameter key : no type generator for gpointer, gpointer*

// Unsupported : g_hmac_get_digest : unsupported parameter buffer : no type generator for guint8, guint8*

// Unsupported : g_hmac_update : unsupported parameter data : no param type

// Unsupported : g_hook_list_marshal : unsupported parameter marshaller : no type generator for HookMarshaller, GHookMarshaller

// Unsupported : g_hook_list_marshal_check : unsupported parameter marshaller : no type generator for HookCheckMarshaller, GHookCheckMarshaller

// Blacklisted : GIConv

// Blacklisted : GIOChannel

// Unsupported : g_key_file_get_boolean_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_double_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_integer_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_keys : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_locale_string_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_string_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_load_from_dirs : unsupported parameter search_dirs : no param type

// Ref is a wrapper around the C function g_key_file_ref.
func (recv *KeyFile) Ref() *KeyFile {
	retC := C.g_key_file_ref((*C.GKeyFile)(recv.native))
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_key_file_set_boolean_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_double_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_integer_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_locale_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

// Unref is a wrapper around the C function g_key_file_unref.
func (recv *KeyFile) Unref() {
	C.g_key_file_unref((*C.GKeyFile)(recv.native))

	return
}

// Unsupported : g_main_context_check : unsupported parameter fds : no param type

// Unsupported : g_main_context_get_poll_func : no return generator

// Unsupported : g_main_context_invoke : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_invoke_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_prepare : unsupported parameter priority : no type generator for gint, gint*

// Unsupported : g_main_context_query : unsupported parameter timeout_ : no type generator for gint, gint*

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc, GPollFunc

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// MappedFileNewFromFd is a wrapper around the C function g_mapped_file_new_from_fd.
func MappedFileNewFromFd(fd int32, writable bool) (*MappedFile, error) {
	c_fd := (C.gint)(fd)

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new_from_fd(c_fd, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_markup_parse_context_get_position : unsupported parameter line_number : no type generator for gint, gint*

// Unsupported : g_match_info_fetch_all : no return type

// Unsupported : g_match_info_fetch_named_pos : unsupported parameter start_pos : no type generator for gint, gint*

// Unsupported : g_match_info_fetch_pos : unsupported parameter start_pos : no type generator for gint, gint*

// Unsupported : g_node_children_foreach : unsupported parameter func : no type generator for NodeForeachFunc, GNodeForeachFunc

// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc, GCopyFunc

// Unsupported : g_node_traverse : unsupported parameter func : no type generator for NodeTraverseFunc, GNodeTraverseFunc

// Unsupported : g_once_impl : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_option_context_parse : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : g_option_context_parse_strv : unsupported parameter arguments : no param type

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_option_group_set_error_hook : unsupported parameter error_func : no type generator for OptionErrorFunc, GOptionErrorFunc

// Unsupported : g_option_group_set_parse_hooks : unsupported parameter pre_parse_func : no type generator for OptionParseFunc, GOptionParseFunc

// Unsupported : g_option_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// Replace is a wrapper around the C function g_private_replace.
func (recv *Private) Replace(value uintptr) {
	c_value := (C.gpointer)(value)

	C.g_private_replace((*C.GPrivate)(recv.native), c_value)

	return
}

// Unsupported : g_queue_find_custom : unsupported parameter func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func, GFunc

// Unsupported : g_queue_free_full : unsupported parameter free_func : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_queue_insert_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc, GCompareDataFunc

// RWLock is a wrapper around the C record GRWLock.
type RWLock struct {
	native *C.GRWLock
	// Private : p
	// Private : i
}

func RWLockNewFromC(u unsafe.Pointer) *RWLock {
	c := (*C.GRWLock)(u)
	if c == nil {
		return nil
	}

	g := &RWLock{native: c}

	return g
}

func (recv *RWLock) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Clear is a wrapper around the C function g_rw_lock_clear.
func (recv *RWLock) Clear() {
	C.g_rw_lock_clear((*C.GRWLock)(recv.native))

	return
}

// Init is a wrapper around the C function g_rw_lock_init.
func (recv *RWLock) Init() {
	C.g_rw_lock_init((*C.GRWLock)(recv.native))

	return
}

// ReaderLock is a wrapper around the C function g_rw_lock_reader_lock.
func (recv *RWLock) ReaderLock() {
	C.g_rw_lock_reader_lock((*C.GRWLock)(recv.native))

	return
}

// ReaderTrylock is a wrapper around the C function g_rw_lock_reader_trylock.
func (recv *RWLock) ReaderTrylock() bool {
	retC := C.g_rw_lock_reader_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ReaderUnlock is a wrapper around the C function g_rw_lock_reader_unlock.
func (recv *RWLock) ReaderUnlock() {
	C.g_rw_lock_reader_unlock((*C.GRWLock)(recv.native))

	return
}

// WriterLock is a wrapper around the C function g_rw_lock_writer_lock.
func (recv *RWLock) WriterLock() {
	C.g_rw_lock_writer_lock((*C.GRWLock)(recv.native))

	return
}

// WriterTrylock is a wrapper around the C function g_rw_lock_writer_trylock.
func (recv *RWLock) WriterTrylock() bool {
	retC := C.g_rw_lock_writer_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// WriterUnlock is a wrapper around the C function g_rw_lock_writer_unlock.
func (recv *RWLock) WriterUnlock() {
	C.g_rw_lock_writer_unlock((*C.GRWLock)(recv.native))

	return
}

// Unsupported : g_rand_set_seed_array : unsupported parameter seed : no type generator for guint32, const guint32*

// RecMutex is a wrapper around the C record GRecMutex.
type RecMutex struct {
	native *C.GRecMutex
	// Private : p
	// Private : i
}

func RecMutexNewFromC(u unsafe.Pointer) *RecMutex {
	c := (*C.GRecMutex)(u)
	if c == nil {
		return nil
	}

	g := &RecMutex{native: c}

	return g
}

func (recv *RecMutex) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Clear is a wrapper around the C function g_rec_mutex_clear.
func (recv *RecMutex) Clear() {
	C.g_rec_mutex_clear((*C.GRecMutex)(recv.native))

	return
}

// Init is a wrapper around the C function g_rec_mutex_init.
func (recv *RecMutex) Init() {
	C.g_rec_mutex_init((*C.GRecMutex)(recv.native))

	return
}

// Lock is a wrapper around the C function g_rec_mutex_lock.
func (recv *RecMutex) Lock() {
	C.g_rec_mutex_lock((*C.GRecMutex)(recv.native))

	return
}

// Trylock is a wrapper around the C function g_rec_mutex_trylock.
func (recv *RecMutex) Trylock() bool {
	retC := C.g_rec_mutex_trylock((*C.GRecMutex)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unlock is a wrapper around the C function g_rec_mutex_unlock.
func (recv *RecMutex) Unlock() {
	C.g_rec_mutex_unlock((*C.GRecMutex)(recv.native))

	return
}

// Unsupported : g_regex_match : unsupported parameter match_info : record with indirection level of 2

// Unsupported : g_regex_match_all : unsupported parameter match_info : record with indirection level of 2

// Unsupported : g_regex_match_all_full : unsupported parameter string : no param type

// Unsupported : g_regex_match_full : unsupported parameter string : no param type

// Unsupported : g_regex_replace : unsupported parameter string : no param type

// Unsupported : g_regex_replace_eval : unsupported parameter string : no param type

// Unsupported : g_regex_replace_literal : unsupported parameter string : no param type

// Unsupported : g_regex_split : no return type

// Unsupported : g_regex_split_full : unsupported parameter string : no param type

// Unsupported : g_scanner_cur_value : no return generator

// Unsupported : g_scanner_error : unsupported parameter ... : varargs

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc, GHFunc

// Unsupported : g_scanner_warn : unsupported parameter ... : varargs

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func, GFunc

// Unsupported : g_sequence_insert_sorted : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_lookup : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_lookup_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_search : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_search_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_string_append_printf : unsupported parameter ... : varargs

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_string_printf : unsupported parameter ... : varargs

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Ref is a wrapper around the C function g_thread_ref.
func (recv *Thread) Ref() *Thread {
	retC := C.g_thread_ref((*C.GThread)(recv.native))
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_thread_unref.
func (recv *Thread) Unref() {
	C.g_thread_unref((*C.GThread)(recv.native))

	return
}

// Unsupported : g_thread_pool_set_sort_function : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_time_zone_adjust_time : unsupported parameter time_ : no type generator for gint64, gint64*

// Unsupported : g_timer_elapsed : unsupported parameter microseconds : no type generator for gulong, gulong*

// Unsupported : g_tree_foreach : unsupported parameter func : no type generator for TraverseFunc, GTraverseFunc

// Unsupported : g_tree_lookup_extended : unsupported parameter orig_key : no type generator for gpointer, gpointer*

// Unsupported : g_tree_search : unsupported parameter search_func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_tree_traverse : unsupported parameter traverse_func : no type generator for TraverseFunc, GTraverseFunc

// Blacklisted : GVariant

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_add : unsupported parameter ... : varargs

// Unsupported : g_variant_builder_add_parsed : unsupported parameter ... : varargs

// Unsupported : g_variant_builder_add_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_builder_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_builder_init : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_open : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant

// Unsupported : g_variant_dict_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_dict_init : unsupported parameter from_asv : Blacklisted record : GVariant

// Unsupported : g_variant_dict_insert : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_insert_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_dict_lookup : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_lookup_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Unsupported : g_variant_iter_init : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_iter_loop : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next_value : return type : Blacklisted record : GVariant

// Blacklisted : GVariantType

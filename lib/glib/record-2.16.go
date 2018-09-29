// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_bytes_new : unsupported parameter data : no param type

// Unsupported : g_bytes_new_static : unsupported parameter data : no param type

// Unsupported : g_bytes_new_take : unsupported parameter data : no param type

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no param type

// Unsupported : g_bytes_get_data : unsupported parameter size : no type generator for gsize, gsize*

// Unsupported : g_bytes_unref_to_array : no return type

// Unsupported : g_bytes_unref_to_data : unsupported parameter size : no type generator for gsize, gsize*

// Checksum is a wrapper around the C record GChecksum.
type Checksum struct {
	native *C.GChecksum
}

func ChecksumNewFromC(u unsafe.Pointer) *Checksum {
	c := (*C.GChecksum)(u)
	if c == nil {
		return nil
	}

	g := &Checksum{native: c}

	return g
}

func (recv *Checksum) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ChecksumNew is a wrapper around the C function g_checksum_new.
func ChecksumNew(checksumType ChecksumType) *Checksum {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_new(c_checksum_type)
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_checksum_copy.
func (recv *Checksum) Copy() *Checksum {
	retC := C.g_checksum_copy((*C.GChecksum)(recv.native))
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_checksum_free.
func (recv *Checksum) Free() {
	C.g_checksum_free((*C.GChecksum)(recv.native))

	return
}

// Unsupported : g_checksum_get_digest : unsupported parameter buffer : no type generator for guint8, guint8*

// GetString is a wrapper around the C function g_checksum_get_string.
func (recv *Checksum) GetString() string {
	retC := C.g_checksum_get_string((*C.GChecksum)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_checksum_update : unsupported parameter data : no param type

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_date_set_time_t : unsupported parameter timet : no type generator for glong, time_t

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer, tm*

// Unsupported : g_date_time_get_ymd : unsupported parameter year : no type generator for gint, gint*

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// GetHashTable is a wrapper around the C function g_hash_table_iter_get_hash_table.
func (recv *HashTableIter) GetHashTable() *HashTable {
	retC := C.g_hash_table_iter_get_hash_table((*C.GHashTableIter)(recv.native))
	retGo := HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_hash_table_iter_init.
func (recv *HashTableIter) Init(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	C.g_hash_table_iter_init((*C.GHashTableIter)(recv.native), c_hash_table)

	return
}

// Unsupported : g_hash_table_iter_next : unsupported parameter key : no type generator for gpointer, gpointer*

// Remove is a wrapper around the C function g_hash_table_iter_remove.
func (recv *HashTableIter) Remove() {
	C.g_hash_table_iter_remove((*C.GHashTableIter)(recv.native))

	return
}

// Steal is a wrapper around the C function g_hash_table_iter_steal.
func (recv *HashTableIter) Steal() {
	C.g_hash_table_iter_steal((*C.GHashTableIter)(recv.native))

	return
}

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

// Unsupported : g_key_file_set_boolean_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_double_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_integer_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_locale_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_main_context_check : unsupported parameter fds : no param type

// Unsupported : g_main_context_get_poll_func : no return generator

// Unsupported : g_main_context_invoke : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_invoke_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_prepare : unsupported parameter priority : no type generator for gint, gint*

// Unsupported : g_main_context_query : unsupported parameter timeout_ : no type generator for gint, gint*

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc, GPollFunc

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// GetElementStack is a wrapper around the C function g_markup_parse_context_get_element_stack.
func (recv *MarkupParseContext) GetElementStack() *SList {
	retC := C.g_markup_parse_context_get_element_stack((*C.GMarkupParseContext)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported : g_queue_find_custom : unsupported parameter func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func, GFunc

// Unsupported : g_queue_free_full : unsupported parameter free_func : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_queue_insert_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_rand_set_seed_array : unsupported parameter seed : no type generator for guint32, const guint32*

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

// AppendUriEscaped is a wrapper around the C function g_string_append_uri_escaped.
func (recv *String) AppendUriEscaped(unescaped string, reservedCharsAllowed string, allowUtf8 bool) *String {
	c_unescaped := C.CString(unescaped)
	defer C.free(unsafe.Pointer(c_unescaped))

	c_reserved_chars_allowed := C.CString(reservedCharsAllowed)
	defer C.free(unsafe.Pointer(c_reserved_chars_allowed))

	c_allow_utf8 :=
		boolToGboolean(allowUtf8)

	retC := C.g_string_append_uri_escaped((*C.GString)(recv.native), c_unescaped, c_reserved_chars_allowed, c_allow_utf8)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_string_printf : unsupported parameter ... : varargs

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Add is a wrapper around the C function g_test_suite_add.
func (recv *TestSuite) Add(testCase *TestCase) {
	c_test_case := (*C.GTestCase)(testCase.ToC())

	C.g_test_suite_add((*C.GTestSuite)(recv.native), c_test_case)

	return
}

// AddSuite is a wrapper around the C function g_test_suite_add_suite.
func (recv *TestSuite) AddSuite(nestedsuite *TestSuite) {
	c_nestedsuite := (*C.GTestSuite)(nestedsuite.ToC())

	C.g_test_suite_add_suite((*C.GTestSuite)(recv.native), c_nestedsuite)

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

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

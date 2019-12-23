// Code generated - DO NOT EDIT.
// +build glib_2.24

package glib

import "unsafe"

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-2.0/glib-object.h>
import "C"

type Array C.GArray
type AsyncQueue C.GAsyncQueue
type BookmarkFile C.GBookmarkFile
type ByteArray C.GByteArray
type Checksum C.GChecksum
type Cond C.GCond
type Data C.GData
type Date C.GDate
type DebugKey C.GDebugKey
type Dir C.GDir
type Error C.GError
type HashTable C.GHashTable
type HashTableIter C.GHashTableIter
type Hook C.GHook
type HookList C.GHookList
type IConv C.GIConv
type IOChannel C.GIOChannel
type IOFuncs C.GIOFuncs
type KeyFile C.GKeyFile
type List C.GList
type MainContext C.GMainContext
type MainLoop C.GMainLoop
type MappedFile C.GMappedFile
type MarkupParseContext C.GMarkupParseContext
type MarkupParser C.GMarkupParser
type MatchInfo C.GMatchInfo
type MemVTable C.GMemVTable
type Node C.GNode
type Once C.GOnce
type OptionContext C.GOptionContext
type OptionEntry C.GOptionEntry
type OptionGroup C.GOptionGroup
type PatternSpec C.GPatternSpec
type PollFD C.GPollFD
type Private C.GPrivate
type PtrArray C.GPtrArray
type Queue C.GQueue
type Rand C.GRand
type Regex C.GRegex
type SList C.GSList
type Scanner C.GScanner
type ScannerConfig C.GScannerConfig
type Sequence C.GSequence
type SequenceIter C.GSequenceIter
type Source C.GSource
type SourceCallbackFuncs C.GSourceCallbackFuncs
type SourceFuncs C.GSourceFuncs
type SourcePrivate C.GSourcePrivate
type StatBuf C.GStatBuf
type String C.GString
type StringChunk C.GStringChunk
type TestCase C.GTestCase
type TestConfig C.GTestConfig
type TestLogBuffer C.GTestLogBuffer

// UNSUPPORTED : TestLogMsg : blacklisted
type TestSuite C.GTestSuite
type Thread C.GThread
type ThreadPool C.GThreadPool
type TimeVal C.GTimeVal
type Timer C.GTimer
type TrashStack C.GTrashStack
type Tree C.GTree
type Variant C.GVariant
type VariantBuilder C.GVariantBuilder
type VariantIter C.GVariantIter
type VariantType C.GVariantType

func Fn_g_access(param0 string, param1 int) {}

func Fn_g_ascii_digit_value(param0 int8) {}

func Fn_g_ascii_dtostr(param0 string, param1 int, param2 float64) {}

func Fn_g_ascii_formatd(param0 string, param1 int, param2 string, param3 float64) {}

func Fn_g_ascii_strcasecmp(param0 string, param1 string) {}

func Fn_g_ascii_strdown(param0 string, param1 uint64) {}

func Fn_g_ascii_strncasecmp(param0 string, param1 string, param2 uint64) {}

func Fn_g_ascii_strtod(param0 string, param1 string) {}

func Fn_g_ascii_strtoll(param0 string, param1 string, param2 uint) {}

func Fn_g_ascii_strtoull(param0 string, param1 string, param2 uint) {}

func Fn_g_ascii_strup(param0 string, param1 uint64) {}

func Fn_g_ascii_tolower(param0 int8) {}

func Fn_g_ascii_toupper(param0 int8) {}

func Fn_g_ascii_xdigit_value(param0 int8) {}

func Fn_g_assert_warning(param0 string, param1 string, param2 int, param3 string, param4 string) {}

func Fn_g_assertion_message(param0 string, param1 string, param2 int, param3 string, param4 string) {}

func Fn_g_assertion_message_cmpnum(param0 string, param1 string, param2 int, param3 string, param4 string, param5 float64, param6 string, param7 float64, param8 int8) {
}

func Fn_g_assertion_message_cmpstr(param0 string, param1 string, param2 int, param3 string, param4 string, param5 string, param6 string, param7 string) {
}

func Fn_g_assertion_message_error(param0 string, param1 string, param2 int, param3 string, param4 string, param5 unsafe.Pointer, param6 uint32, param7 int) {
}

func Fn_g_assertion_message_expr(param0 string, param1 string, param2 int, param3 string, param4 string) {
}

// UNSUPPORTED : atexit : has callback

func Fn_g_atomic_int_add(param0 *int, param1 int) {}

func Fn_g_atomic_int_compare_and_exchange(param0 *int, param1 int, param2 int) {}

func Fn_g_atomic_int_dec_and_test(param0 *int) {}

func Fn_g_atomic_int_exchange_and_add(param0 *int, param1 int) {}

func Fn_g_atomic_int_get(param0 *int) {}

func Fn_g_atomic_int_inc(param0 *int) {}

func Fn_g_atomic_int_set(param0 *int, param1 int) {}

func Fn_g_atomic_pointer_compare_and_exchange(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {
}

func Fn_g_atomic_pointer_get(param0 unsafe.Pointer) {}

func Fn_g_atomic_pointer_set(param0 unsafe.Pointer, param1 *unsafe.Pointer) {}

// UNSUPPORTED : atomic_rc_box_release_full : has callback

func Fn_g_base64_decode(param0 string, param1 *uint64) {}

func Fn_g_base64_decode_inplace(param0 []uint8, param1 *uint64) {}

func Fn_g_base64_decode_step(param0 []uint8, param1 uint64, param2 []uint8, param3 *int, param4 *uint) {
}

func Fn_g_base64_encode(param0 []uint8, param1 uint64) {}

func Fn_g_base64_encode_close(param0 bool, param1 []uint8, param2 *int, param3 *int) {}

func Fn_g_base64_encode_step(param0 []uint8, param1 uint64, param2 bool, param3 []uint8, param4 *int, param5 *int) {
}

func Fn_g_basename(param0 string) {}

func Fn_g_bit_lock(param0 *int, param1 int) {}

func Fn_g_bit_nth_lsf(param0 uint64, param1 int) {}

func Fn_g_bit_nth_msf(param0 uint64, param1 int) {}

func Fn_g_bit_storage(param0 uint64) {}

func Fn_g_bit_trylock(param0 *int, param1 int) {}

func Fn_g_bit_unlock(param0 *int, param1 int) {}

func Fn_g_bookmark_file_error_quark() {
	C.g_bookmark_file_error_quark()
}

// UNSUPPORTED : build_filename : has varargs

// UNSUPPORTED : build_filename_valist : has va_list

func Fn_g_build_filenamev(param0 []string) {}

// UNSUPPORTED : build_path : has varargs

func Fn_g_build_pathv(param0 string, param1 []string) {}

func Fn_g_byte_array_free(param0 []uint8, param1 bool) {}

func Fn_g_byte_array_new() {
	C.g_byte_array_new()
}

func Fn_g_byte_array_unref(param0 []uint8) {}

func Fn_g_chdir(param0 string) {}

func Fn_glib_check_version(param0 uint, param1 uint, param2 uint) {}

func Fn_g_checksum_type_get_length(param0 int) {}

// UNSUPPORTED : child_watch_add : has callback

// UNSUPPORTED : child_watch_add_full : has callback

func Fn_g_child_watch_source_new(param0 int) {}

func Fn_g_clear_error(param0 *unsafe.Pointer) {}

// UNSUPPORTED : clear_handle_id : has callback

// UNSUPPORTED : clear_pointer : has callback

func Fn_g_compute_checksum_for_data(param0 int, param1 []uint8, param2 uint64) {}

func Fn_g_compute_checksum_for_string(param0 int, param1 string, param2 uint64) {}

func Fn_g_convert(param0 []uint8, param1 uint64, param2 string, param3 string, param4 *uint64, param5 *uint64) {
}

func Fn_g_convert_error_quark() {
	C.g_convert_error_quark()
}

func Fn_g_convert_with_fallback(param0 []uint8, param1 uint64, param2 string, param3 string, param4 string, param5 *uint64, param6 *uint64) {
}

func Fn_g_convert_with_iconv(param0 []uint8, param1 uint64, param2 IConv, param3 *uint64, param4 *uint64) {
}

func Fn_g_datalist_clear(param0 *unsafe.Pointer) {}

// UNSUPPORTED : datalist_foreach : has callback

func Fn_g_datalist_get_data(param0 *unsafe.Pointer, param1 string) {}

func Fn_g_datalist_get_flags(param0 *unsafe.Pointer) {}

// UNSUPPORTED : datalist_id_dup_data : has callback

func Fn_g_datalist_id_get_data(param0 *unsafe.Pointer, param1 uint32) {}

func Fn_g_datalist_id_remove_no_notify(param0 *unsafe.Pointer, param1 uint32) {}

// UNSUPPORTED : datalist_id_replace_data : has callback

// UNSUPPORTED : datalist_id_set_data_full : has callback

func Fn_g_datalist_init(param0 *unsafe.Pointer) {}

func Fn_g_datalist_set_flags(param0 *unsafe.Pointer, param1 uint) {}

func Fn_g_datalist_unset_flags(param0 *unsafe.Pointer, param1 uint) {}

func Fn_g_dataset_destroy(param0 unsafe.Pointer) {}

// UNSUPPORTED : dataset_foreach : has callback

func Fn_g_dataset_id_get_data(param0 unsafe.Pointer, param1 uint32) {}

func Fn_g_dataset_id_remove_no_notify(param0 unsafe.Pointer, param1 uint32) {}

// UNSUPPORTED : dataset_id_set_data_full : has callback

func Fn_g_date_get_days_in_month(param0 int, param1 uint16) {}

func Fn_g_date_get_monday_weeks_in_year(param0 uint16) {}

func Fn_g_date_get_sunday_weeks_in_year(param0 uint16) {}

func Fn_g_date_is_leap_year(param0 uint16) {}

func Fn_g_date_strftime(param0 string, param1 uint64, param2 string, param3 unsafe.Pointer) {}

func Fn_g_date_valid_day(param0 uint8) {}

func Fn_g_date_valid_dmy(param0 uint8, param1 int, param2 uint16) {}

func Fn_g_date_valid_julian(param0 uint32) {}

func Fn_g_date_valid_month(param0 int) {}

func Fn_g_date_valid_weekday(param0 int) {}

func Fn_g_date_valid_year(param0 uint16) {}

func Fn_g_dgettext(param0 string, param1 string) {}

func Fn_g_direct_equal(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_direct_hash(param0 unsafe.Pointer) {}

func Fn_g_dngettext(param0 string, param1 string, param2 string, param3 uint64) {}

func Fn_g_double_equal(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_double_hash(param0 unsafe.Pointer) {}

func Fn_g_dpgettext(param0 string, param1 string, param2 uint64) {}

func Fn_g_dpgettext2(param0 string, param1 string, param2 string) {}

func Fn_g_file_error_from_errno(param0 int) {}

func Fn_g_file_error_quark() {
	C.g_file_error_quark()
}

func Fn_g_file_get_contents(param0 string, param1 []uint8, param2 *uint64) {}

func Fn_g_file_open_tmp(param0 string, param1 string) {}

func Fn_g_file_read_link(param0 string) {}

func Fn_g_file_set_contents(param0 string, param1 []uint8, param2 uint64) {}

func Fn_g_file_test(param0 string, param1 int) {}

func Fn_g_filename_display_basename(param0 string) {}

func Fn_g_filename_display_name(param0 string) {}

func Fn_g_filename_from_uri(param0 string, param1 string) {}

func Fn_g_filename_from_utf8(param0 string, param1 uint64, param2 *uint64, param3 *uint64) {}

func Fn_g_filename_to_uri(param0 string, param1 string) {}

func Fn_g_filename_to_utf8(param0 string, param1 uint64, param2 *uint64, param3 *uint64) {}

func Fn_g_find_program_in_path(param0 string) {}

func Fn_g_format_size_for_display(param0 int64) {}

// UNSUPPORTED : fprintf : has varargs

func Fn_g_free(param0 *unsafe.Pointer) {}

func Fn_g_get_application_name() {
	C.g_get_application_name()
}

func Fn_g_get_charset(param0 string) {}

func Fn_g_get_codeset() {
	C.g_get_codeset()
}

func Fn_g_get_current_dir() {
	C.g_get_current_dir()
}

func Fn_g_get_current_time(param0 unsafe.Pointer) {}

func Fn_g_get_filename_charsets(param0 *[]string) {}

func Fn_g_get_home_dir() {
	C.g_get_home_dir()
}

func Fn_g_get_host_name() {
	C.g_get_host_name()
}

func Fn_g_get_language_names() {
	C.g_get_language_names()
}

func Fn_g_get_prgname() {
	C.g_get_prgname()
}

func Fn_g_get_real_name() {
	C.g_get_real_name()
}

func Fn_g_get_system_config_dirs() {
	C.g_get_system_config_dirs()
}

func Fn_g_get_system_data_dirs() {
	C.g_get_system_data_dirs()
}

func Fn_g_get_tmp_dir() {
	C.g_get_tmp_dir()
}

func Fn_g_get_user_cache_dir() {
	C.g_get_user_cache_dir()
}

func Fn_g_get_user_config_dir() {
	C.g_get_user_config_dir()
}

func Fn_g_get_user_data_dir() {
	C.g_get_user_data_dir()
}

func Fn_g_get_user_name() {
	C.g_get_user_name()
}

func Fn_g_get_user_special_dir(param0 int) {}

func Fn_g_getenv(param0 string) {}

func Fn_g_hash_table_destroy(param0 unsafe.Pointer) {}

func Fn_g_hash_table_insert(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {}

func Fn_g_hash_table_lookup(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_hash_table_lookup_extended(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 *unsafe.Pointer) {
}

func Fn_g_hash_table_remove(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_hash_table_remove_all(param0 unsafe.Pointer) {}

func Fn_g_hash_table_replace(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {
}

func Fn_g_hash_table_size(param0 unsafe.Pointer) {}

func Fn_g_hash_table_steal(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_hash_table_steal_all(param0 unsafe.Pointer) {}

func Fn_g_hash_table_unref(param0 unsafe.Pointer) {}

func Fn_g_hook_destroy(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_hook_destroy_link(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_hook_free(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_hook_insert_before(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_g_hook_prepend(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_hook_unref(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_hostname_is_ascii_encoded(param0 string) {}

func Fn_g_hostname_is_ip_address(param0 string) {}

func Fn_g_hostname_is_non_ascii(param0 string) {}

func Fn_g_hostname_to_ascii(param0 string) {}

func Fn_g_hostname_to_unicode(param0 string) {}

func Fn_g_iconv(param0 IConv, param1 string, param2 *uint64, param3 string, param4 *uint64) {}

func Fn_g_iconv_open(param0 string, param1 string) {}

// UNSUPPORTED : idle_add : has callback

// UNSUPPORTED : idle_add_full : has callback

func Fn_g_idle_remove_by_data(param0 *unsafe.Pointer) {}

func Fn_g_idle_source_new() {
	C.g_idle_source_new()
}

func Fn_g_int64_equal(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_int64_hash(param0 unsafe.Pointer) {}

func Fn_g_int_equal(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_int_hash(param0 unsafe.Pointer) {}

func Fn_g_intern_static_string(param0 string) {}

func Fn_g_intern_string(param0 string) {}

// UNSUPPORTED : io_add_watch : has callback

// UNSUPPORTED : io_add_watch_full : has callback

func Fn_g_io_channel_error_from_errno(param0 int) {}

func Fn_g_io_channel_error_quark() {
	C.g_io_channel_error_quark()
}

func Fn_g_io_create_watch(param0 unsafe.Pointer, param1 int) {}

func Fn_g_key_file_error_quark() {
	C.g_key_file_error_quark()
}

func Fn_g_listenv() {
	C.g_listenv()
}

func Fn_g_locale_from_utf8(param0 string, param1 uint64, param2 *uint64, param3 *uint64) {}

func Fn_g_locale_to_utf8(param0 []uint8, param1 uint64, param2 *uint64, param3 *uint64) {}

// UNSUPPORTED : log : has varargs

func Fn_g_log_default_handler(param0 string, param1 int, param2 string, param3 *unsafe.Pointer) {}

func Fn_g_log_remove_handler(param0 string, param1 uint) {}

func Fn_g_log_set_always_fatal(param0 int) {}

// UNSUPPORTED : log_set_default_handler : has callback

func Fn_g_log_set_fatal_mask(param0 string, param1 int) {}

// UNSUPPORTED : log_set_handler : has callback

// UNSUPPORTED : log_set_handler_full : has callback

// UNSUPPORTED : log_set_writer_func : has callback

// UNSUPPORTED : log_structured : has varargs

// UNSUPPORTED : log_structured_standard : has varargs

// UNSUPPORTED : logv : has va_list

func Fn_g_main_context_default() {
	C.g_main_context_default()
}

func Fn_g_main_context_get_thread_default() {
	C.g_main_context_get_thread_default()
}

func Fn_g_main_current_source() {
	C.g_main_current_source()
}

func Fn_g_main_depth() {
	C.g_main_depth()
}

func Fn_g_malloc(param0 uint64) {}

func Fn_g_malloc0(param0 uint64) {}

func Fn_g_malloc0_n(param0 uint64, param1 uint64) {}

func Fn_g_malloc_n(param0 uint64, param1 uint64) {}

// UNSUPPORTED : markup_collect_attributes : has varargs

func Fn_g_markup_error_quark() {
	C.g_markup_error_quark()
}

func Fn_g_markup_escape_text(param0 string, param1 uint64) {}

// UNSUPPORTED : markup_printf_escaped : has varargs

// UNSUPPORTED : markup_vprintf_escaped : has va_list

func Fn_g_mem_is_system_malloc() {
	C.g_mem_is_system_malloc()
}

func Fn_g_mem_profile() {
	C.g_mem_profile()
}

func Fn_g_mem_set_vtable(param0 unsafe.Pointer) {}

func Fn_g_memdup(param0 unsafe.Pointer, param1 uint) {}

func Fn_g_mkdir_with_parents(param0 string, param1 int) {}

func Fn_g_mkstemp(param0 string) {}

func Fn_g_mkstemp_full(param0 string, param1 int, param2 int) {}

func Fn_g_nullify_pointer(param0 *unsafe.Pointer) {}

func Fn_g_number_parser_error_quark() {
	C.g_number_parser_error_quark()
}

func Fn_g_on_error_query(param0 string) {}

func Fn_g_on_error_stack_trace(param0 string) {}

func Fn_g_once_init_enter(param0 unsafe.Pointer) {}

func Fn_g_once_init_leave(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_option_error_quark() {
	C.g_option_error_quark()
}

func Fn_g_parse_debug_string(param0 string, param1 []DebugKey, param2 uint) {}

func Fn_g_path_get_basename(param0 string) {}

func Fn_g_path_get_dirname(param0 string) {}

func Fn_g_path_is_absolute(param0 string) {}

func Fn_g_path_skip_root(param0 string) {}

func Fn_g_pattern_match(param0 unsafe.Pointer, param1 uint, param2 string, param3 string) {}

func Fn_g_pattern_match_simple(param0 string, param1 string) {}

func Fn_g_pattern_match_string(param0 unsafe.Pointer, param1 string) {}

func Fn_g_poll(param0 unsafe.Pointer, param1 uint, param2 int) {}

// UNSUPPORTED : prefix_error : has varargs

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

// UNSUPPORTED : printf : has varargs

// UNSUPPORTED : printf_string_upper_bound : has va_list

func Fn_g_propagate_error(param0 *unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : propagate_prefixed_error : has varargs

// UNSUPPORTED : ptr_array_find_with_equal_func : has callback

// UNSUPPORTED : qsort_with_data : has callback

func Fn_g_quark_from_static_string(param0 string) {}

func Fn_g_quark_from_string(param0 string) {}

func Fn_g_quark_to_string(param0 uint32) {}

func Fn_g_quark_try_string(param0 string) {}

func Fn_g_random_double() {
	C.g_random_double()
}

func Fn_g_random_double_range(param0 float64, param1 float64) {}

func Fn_g_random_int() {
	C.g_random_int()
}

func Fn_g_random_int_range(param0 int32, param1 int32) {}

func Fn_g_random_set_seed(param0 uint32) {}

// UNSUPPORTED : rc_box_release_full : has callback

func Fn_g_realloc(param0 *unsafe.Pointer, param1 uint64) {}

func Fn_g_realloc_n(param0 *unsafe.Pointer, param1 uint64, param2 uint64) {}

func Fn_g_regex_check_replacement(param0 string, param1 *bool) {}

func Fn_g_regex_error_quark() {
	C.g_regex_error_quark()
}

func Fn_g_regex_escape_string(param0 []string, param1 int) {}

func Fn_g_regex_match_simple(param0 string, param1 string, param2 int, param3 int) {}

func Fn_g_regex_split_simple(param0 string, param1 string, param2 int, param3 int) {}

func Fn_g_reload_user_special_dirs_cache() {
	C.g_reload_user_special_dirs_cache()
}

func Fn_g_return_if_fail_warning(param0 string, param1 string, param2 string) {}

func Fn_g_rmdir(param0 string) {}

func Fn_g_sequence_get(param0 unsafe.Pointer) {}

func Fn_g_sequence_insert_before(param0 unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_g_sequence_move(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_sequence_move_range(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_g_sequence_range_get_midpoint(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_sequence_remove(param0 unsafe.Pointer) {}

func Fn_g_sequence_remove_range(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_sequence_set(param0 unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_g_sequence_swap(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_set_application_name(param0 string) {}

// UNSUPPORTED : set_error : has varargs

func Fn_g_set_error_literal(param0 *unsafe.Pointer, param1 uint32, param2 int, param3 string) {}

func Fn_g_set_prgname(param0 string) {}

// UNSUPPORTED : set_print_handler : has callback

// UNSUPPORTED : set_printerr_handler : has callback

func Fn_g_setenv(param0 string, param1 string, param2 bool) {}

func Fn_g_shell_error_quark() {
	C.g_shell_error_quark()
}

func Fn_g_shell_parse_argv(param0 string, param1 *int, param2 *[]string) {}

func Fn_g_shell_quote(param0 string) {}

func Fn_g_shell_unquote(param0 string) {}

func Fn_g_slice_alloc(param0 uint64) {}

func Fn_g_slice_alloc0(param0 uint64) {}

func Fn_g_slice_copy(param0 uint64, param1 unsafe.Pointer) {}

func Fn_g_slice_free1(param0 uint64, param1 *unsafe.Pointer) {}

func Fn_g_slice_free_chain_with_offset(param0 uint64, param1 *unsafe.Pointer, param2 uint64) {}

func Fn_g_slice_get_config(param0 int) {}

func Fn_g_slice_get_config_state(param0 int, param1 int64, param2 *uint) {}

func Fn_g_slice_set_config(param0 int, param1 int64) {}

// UNSUPPORTED : snprintf : has varargs

func Fn_g_source_remove(param0 uint) {}

func Fn_g_source_remove_by_funcs_user_data(param0 unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_g_source_remove_by_user_data(param0 *unsafe.Pointer) {}

func Fn_g_spaced_primes_closest(param0 uint) {}

// UNSUPPORTED : spawn_async : has callback

// UNSUPPORTED : spawn_async_with_fds : has callback

// UNSUPPORTED : spawn_async_with_pipes : has callback

func Fn_g_spawn_close_pid(param0 int) {}

func Fn_g_spawn_command_line_async(param0 string) {}

func Fn_g_spawn_command_line_sync(param0 string, param1 []uint8, param2 []uint8, param3 *int) {}

func Fn_g_spawn_error_quark() {
	C.g_spawn_error_quark()
}

func Fn_g_spawn_exit_error_quark() {
	C.g_spawn_exit_error_quark()
}

// UNSUPPORTED : spawn_sync : has callback

// UNSUPPORTED : sprintf : has varargs

func Fn_g_stpcpy(param0 string, param1 string) {}

func Fn_g_str_equal(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_str_has_prefix(param0 string, param1 string) {}

func Fn_g_str_has_suffix(param0 string, param1 string) {}

func Fn_g_str_hash(param0 unsafe.Pointer) {}

func Fn_g_strcanon(param0 string, param1 string, param2 int8) {}

func Fn_g_strcasecmp(param0 string, param1 string) {}

func Fn_g_strchomp(param0 string) {}

func Fn_g_strchug(param0 string) {}

func Fn_g_strcmp0(param0 string, param1 string) {}

func Fn_g_strcompress(param0 string) {}

// UNSUPPORTED : strconcat : has varargs

func Fn_g_strdelimit(param0 string, param1 string, param2 int8) {}

func Fn_g_strdown(param0 string) {}

func Fn_g_strdup(param0 string) {}

// UNSUPPORTED : strdup_printf : has varargs

// UNSUPPORTED : strdup_vprintf : has va_list

func Fn_g_strdupv(param0 string) {}

func Fn_g_strerror(param0 int) {}

func Fn_g_strescape(param0 string, param1 string) {}

func Fn_g_strfreev(param0 string) {}

func Fn_g_string_new(param0 string) {}

func Fn_g_string_new_len(param0 string, param1 uint64) {}

func Fn_g_string_sized_new(param0 uint64) {}

func Fn_g_strip_context(param0 string, param1 string) {}

// UNSUPPORTED : strjoin : has varargs

func Fn_g_strjoinv(param0 string, param1 string) {}

func Fn_g_strlcat(param0 string, param1 string, param2 uint64) {}

func Fn_g_strlcpy(param0 string, param1 string, param2 uint64) {}

func Fn_g_strncasecmp(param0 string, param1 string, param2 uint) {}

func Fn_g_strndup(param0 string, param1 uint64) {}

func Fn_g_strnfill(param0 uint64, param1 int8) {}

func Fn_g_strreverse(param0 string) {}

func Fn_g_strrstr(param0 string, param1 string) {}

func Fn_g_strrstr_len(param0 string, param1 uint64, param2 string) {}

func Fn_g_strsignal(param0 int) {}

func Fn_g_strsplit(param0 string, param1 string, param2 int) {}

func Fn_g_strsplit_set(param0 string, param1 string, param2 int) {}

func Fn_g_strstr_len(param0 string, param1 uint64, param2 string) {}

func Fn_g_strtod(param0 string, param1 string) {}

func Fn_g_strup(param0 string) {}

func Fn_g_strv_get_type() {
	C.g_strv_get_type()
}

func Fn_g_strv_length(param0 string) {}

// UNSUPPORTED : test_add_data_func : has callback

// UNSUPPORTED : test_add_data_func_full : has callback

// UNSUPPORTED : test_add_func : has callback

// UNSUPPORTED : test_add_vtable : has callback

func Fn_g_test_assert_expected_messages_internal(param0 string, param1 string, param2 int, param3 string) {
}

func Fn_g_test_bug(param0 string) {}

func Fn_g_test_bug_base(param0 string) {}

// UNSUPPORTED : test_build_filename : has varargs

// UNSUPPORTED : test_create_case : has callback

func Fn_g_test_create_suite(param0 string) {}

// UNSUPPORTED : test_get_filename : has varargs

func Fn_g_test_get_root() {
	C.g_test_get_root()
}

// UNSUPPORTED : test_init : has varargs

// UNSUPPORTED : test_log_set_fatal_handler : has callback

func Fn_g_test_log_type_name(param0 int) {}

// UNSUPPORTED : test_maximized_result : has varargs

// UNSUPPORTED : test_message : has varargs

// UNSUPPORTED : test_minimized_result : has varargs

// UNSUPPORTED : test_queue_destroy : has callback

func Fn_g_test_queue_free(param0 *unsafe.Pointer) {}

func Fn_g_test_rand_double() {
	C.g_test_rand_double()
}

func Fn_g_test_rand_double_range(param0 float64, param1 float64) {}

func Fn_g_test_rand_int() {
	C.g_test_rand_int()
}

func Fn_g_test_rand_int_range(param0 int32, param1 int32) {}

func Fn_g_test_run() {
	C.g_test_run()
}

func Fn_g_test_run_suite(param0 unsafe.Pointer) {}

func Fn_g_test_timer_elapsed() {
	C.g_test_timer_elapsed()
}

func Fn_g_test_timer_last() {
	C.g_test_timer_last()
}

func Fn_g_test_timer_start() {
	C.g_test_timer_start()
}

func Fn_g_test_trap_assertions(param0 string, param1 string, param2 int, param3 string, param4 uint64, param5 string) {
}

func Fn_g_test_trap_fork(param0 uint64, param1 int) {}

func Fn_g_test_trap_has_passed() {
	C.g_test_trap_has_passed()
}

func Fn_g_test_trap_reached_timeout() {
	C.g_test_trap_reached_timeout()
}

func Fn_g_thread_error_quark() {
	C.g_thread_error_quark()
}

func Fn_g_thread_exit(param0 *unsafe.Pointer) {}

func Fn_g_thread_pool_get_max_idle_time() {
	C.g_thread_pool_get_max_idle_time()
}

func Fn_g_thread_pool_get_max_unused_threads() {
	C.g_thread_pool_get_max_unused_threads()
}

func Fn_g_thread_pool_get_num_unused_threads() {
	C.g_thread_pool_get_num_unused_threads()
}

func Fn_g_thread_pool_set_max_idle_time(param0 uint) {}

func Fn_g_thread_pool_set_max_unused_threads(param0 int) {}

func Fn_g_thread_pool_stop_unused_threads() {
	C.g_thread_pool_stop_unused_threads()
}

func Fn_g_thread_self() {
	C.g_thread_self()
}

func Fn_g_thread_yield() {
	C.g_thread_yield()
}

func Fn_g_time_val_from_iso8601(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : timeout_add : has callback

// UNSUPPORTED : timeout_add_full : has callback

// UNSUPPORTED : timeout_add_seconds : has callback

// UNSUPPORTED : timeout_add_seconds_full : has callback

func Fn_g_timeout_source_new(param0 uint) {}

func Fn_g_timeout_source_new_seconds(param0 uint) {}

func Fn_g_trash_stack_height(param0 *unsafe.Pointer) {}

func Fn_g_trash_stack_peek(param0 *unsafe.Pointer) {}

func Fn_g_trash_stack_pop(param0 *unsafe.Pointer) {}

func Fn_g_trash_stack_push(param0 *unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_g_try_malloc(param0 uint64) {}

func Fn_g_try_malloc0(param0 uint64) {}

func Fn_g_try_malloc0_n(param0 uint64, param1 uint64) {}

func Fn_g_try_malloc_n(param0 uint64, param1 uint64) {}

func Fn_g_try_realloc(param0 *unsafe.Pointer, param1 uint64) {}

func Fn_g_try_realloc_n(param0 *unsafe.Pointer, param1 uint64, param2 uint64) {}

func Fn_g_ucs4_to_utf16(param0 *rune, param1 int64, param2 *int64, param3 *int64) {}

func Fn_g_ucs4_to_utf8(param0 *rune, param1 int64, param2 *int64, param3 *int64) {}

func Fn_g_unichar_break_type(param0 rune) {}

func Fn_g_unichar_combining_class(param0 rune) {}

func Fn_g_unichar_digit_value(param0 rune) {}

func Fn_g_unichar_get_mirror_char(param0 rune, param1 *rune) {}

func Fn_g_unichar_get_script(param0 rune) {}

func Fn_g_unichar_isalnum(param0 rune) {}

func Fn_g_unichar_isalpha(param0 rune) {}

func Fn_g_unichar_iscntrl(param0 rune) {}

func Fn_g_unichar_isdefined(param0 rune) {}

func Fn_g_unichar_isdigit(param0 rune) {}

func Fn_g_unichar_isgraph(param0 rune) {}

func Fn_g_unichar_islower(param0 rune) {}

func Fn_g_unichar_ismark(param0 rune) {}

func Fn_g_unichar_isprint(param0 rune) {}

func Fn_g_unichar_ispunct(param0 rune) {}

func Fn_g_unichar_isspace(param0 rune) {}

func Fn_g_unichar_istitle(param0 rune) {}

func Fn_g_unichar_isupper(param0 rune) {}

func Fn_g_unichar_iswide(param0 rune) {}

func Fn_g_unichar_iswide_cjk(param0 rune) {}

func Fn_g_unichar_isxdigit(param0 rune) {}

func Fn_g_unichar_iszerowidth(param0 rune) {}

func Fn_g_unichar_to_utf8(param0 rune, param1 string) {}

func Fn_g_unichar_tolower(param0 rune) {}

func Fn_g_unichar_totitle(param0 rune) {}

func Fn_g_unichar_toupper(param0 rune) {}

func Fn_g_unichar_type(param0 rune) {}

func Fn_g_unichar_validate(param0 rune) {}

func Fn_g_unichar_xdigit_value(param0 rune) {}

func Fn_g_unicode_canonical_decomposition(param0 rune, param1 *uint64) {}

func Fn_g_unicode_canonical_ordering(param0 *rune, param1 uint64) {}

// UNSUPPORTED : unix_error_quark : blacklisted
// UNSUPPORTED : unix_fd_add : has callback

// UNSUPPORTED : unix_fd_add_full : has callback

// UNSUPPORTED : unix_signal_add : has callback

// UNSUPPORTED : unix_signal_add_full : has callback

func Fn_g_unlink(param0 string) {}

func Fn_g_unsetenv(param0 string) {}

func Fn_g_uri_escape_string(param0 string, param1 string, param2 bool) {}

func Fn_g_uri_list_extract_uris(param0 string) {}

func Fn_g_uri_parse_scheme(param0 string) {}

func Fn_g_uri_unescape_segment(param0 string, param1 string, param2 string) {}

func Fn_g_uri_unescape_string(param0 string, param1 string) {}

func Fn_g_usleep(param0 uint64) {}

func Fn_g_utf16_to_ucs4(param0 *uint16, param1 int64, param2 *int64, param3 *int64) {}

func Fn_g_utf16_to_utf8(param0 *uint16, param1 int64, param2 *int64, param3 *int64) {}

func Fn_g_utf8_casefold(param0 string, param1 uint64) {}

func Fn_g_utf8_collate(param0 string, param1 string) {}

func Fn_g_utf8_collate_key(param0 string, param1 uint64) {}

func Fn_g_utf8_collate_key_for_filename(param0 string, param1 uint64) {}

func Fn_g_utf8_find_next_char(param0 string, param1 string) {}

func Fn_g_utf8_find_prev_char(param0 string, param1 string) {}

func Fn_g_utf8_get_char(param0 string) {}

func Fn_g_utf8_get_char_validated(param0 string, param1 uint64) {}

func Fn_g_utf8_normalize(param0 string, param1 uint64, param2 int) {}

func Fn_g_utf8_offset_to_pointer(param0 string, param1 int64) {}

func Fn_g_utf8_pointer_to_offset(param0 string, param1 string) {}

func Fn_g_utf8_prev_char(param0 string) {}

func Fn_g_utf8_strchr(param0 string, param1 uint64, param2 rune) {}

func Fn_g_utf8_strdown(param0 string, param1 uint64) {}

func Fn_g_utf8_strlen(param0 string, param1 uint64) {}

func Fn_g_utf8_strncpy(param0 string, param1 string, param2 uint64) {}

func Fn_g_utf8_strrchr(param0 string, param1 uint64, param2 rune) {}

func Fn_g_utf8_strreverse(param0 string, param1 uint64) {}

func Fn_g_utf8_strup(param0 string, param1 uint64) {}

func Fn_g_utf8_to_ucs4(param0 string, param1 int64, param2 *int64, param3 *int64) {}

func Fn_g_utf8_to_ucs4_fast(param0 string, param1 int64, param2 *int64) {}

func Fn_g_utf8_to_utf16(param0 string, param1 int64, param2 *int64, param3 *int64) {}

func Fn_g_utf8_validate(param0 []uint8, param1 uint64, param2 string) {}

func Fn_g_variant_get_gtype() {
	C.g_variant_get_gtype()
}

func Fn_g_variant_is_object_path(param0 string) {}

func Fn_g_variant_is_signature(param0 string) {}

func Fn_g_variant_parse(param0 unsafe.Pointer, param1 string, param2 string, param3 string) {}

func Fn_g_variant_parse_error_quark() {
	C.g_variant_parse_error_quark()
}

func Fn_g_variant_parser_get_error_quark() {
	C.g_variant_parser_get_error_quark()
}

func Fn_g_variant_type_checked_(param0 string) {}

func Fn_g_variant_type_string_get_depth_(param0 string) {}

func Fn_g_variant_type_string_is_valid(param0 string) {}

func Fn_g_variant_type_string_scan(param0 string, param1 string, param2 string) {}

// UNSUPPORTED : vasprintf : has va_list

// UNSUPPORTED : vfprintf : has va_list

// UNSUPPORTED : vprintf : has va_list

// UNSUPPORTED : vsnprintf : has va_list

// UNSUPPORTED : vsprintf : has va_list

func Fn_g_warn_message(param0 string, param1 string, param2 int, param3 string, param4 string) {}

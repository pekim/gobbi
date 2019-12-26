// Code generated - DO NOT EDIT.
// +build !glib_2.4,!glib_2.6,!glib_2.14,!glib_2.16,!glib_2.22,!glib_2.24,!glib_2.26,!glib_2.30,!glib_2.32,!glib_2.38,!glib_2.40,!glib_2.50,!glib_2.54,!glib_2.60

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-2.0/glib-object.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type Array C.GArray
type AsyncQueue C.GAsyncQueue
type BookmarkFile C.GBookmarkFile
type ByteArray C.GByteArray
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
type OptionContext C.GOptionContext
type OptionEntry C.GOptionEntry
type OptionGroup C.GOptionGroup
type PatternSpec C.GPatternSpec
type PollFD C.GPollFD
type Private C.GPrivate
type PtrArray C.GPtrArray
type Queue C.GQueue
type Rand C.GRand
type SList C.GSList
type Scanner C.GScanner
type ScannerConfig C.GScannerConfig
type Sequence C.GSequence
type SequenceIter C.GSequenceIter
type Source C.GSource
type SourceCallbackFuncs C.GSourceCallbackFuncs
type SourceFuncs C.GSourceFuncs
type SourcePrivate C.GSourcePrivate
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
type VariantBuilder C.GVariantBuilder
type VariantIter C.GVariantIter
type VariantType C.GVariantType

// UNSUPPORTED : access : blacklisted
// UNSUPPORTED : ascii_digit_value : blacklisted
// UNSUPPORTED : ascii_dtostr : blacklisted
// UNSUPPORTED : ascii_formatd : blacklisted
// UNSUPPORTED : ascii_strcasecmp : blacklisted
// UNSUPPORTED : ascii_strdown : blacklisted
// UNSUPPORTED : ascii_string_to_signed : blacklisted
// UNSUPPORTED : ascii_string_to_unsigned : blacklisted
// UNSUPPORTED : ascii_strncasecmp : blacklisted
// UNSUPPORTED : ascii_strtod : blacklisted
// UNSUPPORTED : ascii_strtoll : blacklisted
// UNSUPPORTED : ascii_strtoull : blacklisted
// UNSUPPORTED : ascii_strup : blacklisted
// UNSUPPORTED : ascii_tolower : blacklisted
// UNSUPPORTED : ascii_toupper : blacklisted
// UNSUPPORTED : ascii_xdigit_value : blacklisted
// UNSUPPORTED : assert_warning : blacklisted
// UNSUPPORTED : assertion_message : blacklisted
// UNSUPPORTED : assertion_message_cmpnum : blacklisted
// UNSUPPORTED : assertion_message_cmpstr : blacklisted
// UNSUPPORTED : assertion_message_error : blacklisted
// UNSUPPORTED : assertion_message_expr : blacklisted
// UNSUPPORTED : atexit : blacklisted
// UNSUPPORTED : atomic_int_add : blacklisted
// UNSUPPORTED : atomic_int_and : blacklisted
// UNSUPPORTED : atomic_int_compare_and_exchange : blacklisted
// UNSUPPORTED : atomic_int_dec_and_test : blacklisted
// UNSUPPORTED : atomic_int_exchange_and_add : blacklisted
// UNSUPPORTED : atomic_int_get : blacklisted
// UNSUPPORTED : atomic_int_inc : blacklisted
// UNSUPPORTED : atomic_int_or : blacklisted
// UNSUPPORTED : atomic_int_set : blacklisted
// UNSUPPORTED : atomic_int_xor : blacklisted
// UNSUPPORTED : atomic_pointer_add : blacklisted
// UNSUPPORTED : atomic_pointer_and : blacklisted
// UNSUPPORTED : atomic_pointer_compare_and_exchange : blacklisted
// UNSUPPORTED : atomic_pointer_get : blacklisted
// UNSUPPORTED : atomic_pointer_or : blacklisted
// UNSUPPORTED : atomic_pointer_set : blacklisted
// UNSUPPORTED : atomic_pointer_xor : blacklisted
// UNSUPPORTED : atomic_rc_box_acquire : blacklisted
// UNSUPPORTED : atomic_rc_box_alloc : blacklisted
// UNSUPPORTED : atomic_rc_box_alloc0 : blacklisted
// UNSUPPORTED : atomic_rc_box_dup : blacklisted
// UNSUPPORTED : atomic_rc_box_get_size : blacklisted
// UNSUPPORTED : atomic_rc_box_release : blacklisted
// UNSUPPORTED : atomic_rc_box_release_full : blacklisted
// UNSUPPORTED : atomic_ref_count_compare : blacklisted
// UNSUPPORTED : atomic_ref_count_dec : blacklisted
// UNSUPPORTED : atomic_ref_count_inc : blacklisted
// UNSUPPORTED : atomic_ref_count_init : blacklisted
// UNSUPPORTED : base64_decode : blacklisted
// UNSUPPORTED : base64_decode_inplace : blacklisted
// UNSUPPORTED : base64_decode_step : blacklisted
// UNSUPPORTED : base64_encode : blacklisted
// UNSUPPORTED : base64_encode_close : blacklisted
// UNSUPPORTED : base64_encode_step : blacklisted
// UNSUPPORTED : basename : blacklisted
// UNSUPPORTED : bit_lock : blacklisted
// UNSUPPORTED : bit_nth_lsf : blacklisted
// UNSUPPORTED : bit_nth_msf : blacklisted
// UNSUPPORTED : bit_storage : blacklisted
// UNSUPPORTED : bit_trylock : blacklisted
// UNSUPPORTED : bit_unlock : blacklisted
// UNSUPPORTED : bookmark_file_error_quark : blacklisted
// UNSUPPORTED : build_filename : blacklisted
// UNSUPPORTED : build_filename_valist : blacklisted
// UNSUPPORTED : build_filenamev : blacklisted
// UNSUPPORTED : build_path : blacklisted
// UNSUPPORTED : build_pathv : blacklisted
// UNSUPPORTED : byte_array_free : blacklisted
// UNSUPPORTED : byte_array_free_to_bytes : blacklisted
// UNSUPPORTED : byte_array_new : blacklisted
// UNSUPPORTED : byte_array_new_take : blacklisted
// UNSUPPORTED : byte_array_unref : blacklisted
// UNSUPPORTED : canonicalize_filename : blacklisted
// UNSUPPORTED : chdir : blacklisted
// UNSUPPORTED : check_version : blacklisted
// UNSUPPORTED : checksum_type_get_length : blacklisted
// UNSUPPORTED : child_watch_add : blacklisted
// UNSUPPORTED : child_watch_add_full : blacklisted
// UNSUPPORTED : child_watch_source_new : blacklisted
// UNSUPPORTED : clear_error : blacklisted
// UNSUPPORTED : clear_handle_id : blacklisted
// UNSUPPORTED : clear_pointer : blacklisted
// UNSUPPORTED : close : blacklisted
// UNSUPPORTED : compute_checksum_for_bytes : blacklisted
// UNSUPPORTED : compute_checksum_for_data : blacklisted
// UNSUPPORTED : compute_checksum_for_string : blacklisted
// UNSUPPORTED : compute_hmac_for_bytes : blacklisted
// UNSUPPORTED : compute_hmac_for_data : blacklisted
// UNSUPPORTED : compute_hmac_for_string : blacklisted
// UNSUPPORTED : convert : blacklisted
// UNSUPPORTED : convert_error_quark : blacklisted
// UNSUPPORTED : convert_with_fallback : blacklisted
// UNSUPPORTED : convert_with_iconv : blacklisted
// UNSUPPORTED : datalist_clear : blacklisted
// UNSUPPORTED : datalist_foreach : blacklisted
// UNSUPPORTED : datalist_get_data : blacklisted
// UNSUPPORTED : datalist_get_flags : blacklisted
// UNSUPPORTED : datalist_id_dup_data : blacklisted
// UNSUPPORTED : datalist_id_get_data : blacklisted
// UNSUPPORTED : datalist_id_remove_no_notify : blacklisted
// UNSUPPORTED : datalist_id_replace_data : blacklisted
// UNSUPPORTED : datalist_id_set_data_full : blacklisted
// UNSUPPORTED : datalist_init : blacklisted
// UNSUPPORTED : datalist_set_flags : blacklisted
// UNSUPPORTED : datalist_unset_flags : blacklisted
// UNSUPPORTED : dataset_destroy : blacklisted
// UNSUPPORTED : dataset_foreach : blacklisted
// UNSUPPORTED : dataset_id_get_data : blacklisted
// UNSUPPORTED : dataset_id_remove_no_notify : blacklisted
// UNSUPPORTED : dataset_id_set_data_full : blacklisted
// UNSUPPORTED : date_get_days_in_month : blacklisted
// UNSUPPORTED : date_get_monday_weeks_in_year : blacklisted
// UNSUPPORTED : date_get_sunday_weeks_in_year : blacklisted
// UNSUPPORTED : date_is_leap_year : blacklisted
// UNSUPPORTED : date_strftime : blacklisted
// UNSUPPORTED : date_time_compare : blacklisted
// UNSUPPORTED : date_time_equal : blacklisted
// UNSUPPORTED : date_time_hash : blacklisted
// UNSUPPORTED : date_valid_day : blacklisted
// UNSUPPORTED : date_valid_dmy : blacklisted
// UNSUPPORTED : date_valid_julian : blacklisted
// UNSUPPORTED : date_valid_month : blacklisted
// UNSUPPORTED : date_valid_weekday : blacklisted
// UNSUPPORTED : date_valid_year : blacklisted
// UNSUPPORTED : dcgettext : blacklisted
// UNSUPPORTED : dgettext : blacklisted
// UNSUPPORTED : dir_make_tmp : blacklisted
// UNSUPPORTED : direct_equal : blacklisted
// UNSUPPORTED : direct_hash : blacklisted
// UNSUPPORTED : dngettext : blacklisted
// UNSUPPORTED : double_equal : blacklisted
// UNSUPPORTED : double_hash : blacklisted
// UNSUPPORTED : dpgettext : blacklisted
// UNSUPPORTED : dpgettext2 : blacklisted
// UNSUPPORTED : environ_getenv : blacklisted
// UNSUPPORTED : environ_setenv : blacklisted
// UNSUPPORTED : environ_unsetenv : blacklisted
// UNSUPPORTED : file_error_from_errno : blacklisted
// UNSUPPORTED : file_error_quark : blacklisted
// UNSUPPORTED : file_get_contents : blacklisted
// UNSUPPORTED : file_open_tmp : blacklisted
// UNSUPPORTED : file_read_link : blacklisted
// UNSUPPORTED : file_set_contents : blacklisted
// UNSUPPORTED : file_test : blacklisted
// UNSUPPORTED : filename_display_basename : blacklisted
// UNSUPPORTED : filename_display_name : blacklisted
// UNSUPPORTED : filename_from_uri : blacklisted
// UNSUPPORTED : filename_from_utf8 : blacklisted
// UNSUPPORTED : filename_to_uri : blacklisted
// UNSUPPORTED : filename_to_utf8 : blacklisted
// UNSUPPORTED : find_program_in_path : blacklisted
// UNSUPPORTED : format_size : blacklisted
// UNSUPPORTED : format_size_for_display : blacklisted
// UNSUPPORTED : format_size_full : blacklisted
// UNSUPPORTED : fprintf : blacklisted
// UNSUPPORTED : free : blacklisted
// UNSUPPORTED : get_application_name : blacklisted
// UNSUPPORTED : get_charset : blacklisted
// UNSUPPORTED : get_codeset : blacklisted
// UNSUPPORTED : get_console_charset : blacklisted
// UNSUPPORTED : get_current_dir : blacklisted
// UNSUPPORTED : get_current_time : blacklisted
// UNSUPPORTED : get_environ : blacklisted
// UNSUPPORTED : get_filename_charsets : blacklisted
// UNSUPPORTED : get_home_dir : blacklisted
// UNSUPPORTED : get_host_name : blacklisted
// UNSUPPORTED : get_language_names : blacklisted
// UNSUPPORTED : get_language_names_with_category : blacklisted
// UNSUPPORTED : get_locale_variants : blacklisted
// UNSUPPORTED : get_monotonic_time : blacklisted
// UNSUPPORTED : get_num_processors : blacklisted
// UNSUPPORTED : get_prgname : blacklisted
// UNSUPPORTED : get_real_name : blacklisted
// UNSUPPORTED : get_real_time : blacklisted
// UNSUPPORTED : get_system_config_dirs : blacklisted
// UNSUPPORTED : get_system_data_dirs : blacklisted
// UNSUPPORTED : get_tmp_dir : blacklisted
// UNSUPPORTED : get_user_cache_dir : blacklisted
// UNSUPPORTED : get_user_config_dir : blacklisted
// UNSUPPORTED : get_user_data_dir : blacklisted
// UNSUPPORTED : get_user_name : blacklisted
// UNSUPPORTED : get_user_runtime_dir : blacklisted
// UNSUPPORTED : get_user_special_dir : blacklisted
// UNSUPPORTED : getenv : blacklisted
// UNSUPPORTED : hash_table_add : blacklisted
// UNSUPPORTED : hash_table_contains : blacklisted
// UNSUPPORTED : hash_table_destroy : blacklisted
// UNSUPPORTED : hash_table_insert : blacklisted
// UNSUPPORTED : hash_table_lookup : blacklisted
// UNSUPPORTED : hash_table_lookup_extended : blacklisted
// UNSUPPORTED : hash_table_remove : blacklisted
// UNSUPPORTED : hash_table_remove_all : blacklisted
// UNSUPPORTED : hash_table_replace : blacklisted
// UNSUPPORTED : hash_table_size : blacklisted
// UNSUPPORTED : hash_table_steal : blacklisted
// UNSUPPORTED : hash_table_steal_all : blacklisted
// UNSUPPORTED : hash_table_steal_extended : blacklisted
// UNSUPPORTED : hash_table_unref : blacklisted
// UNSUPPORTED : hook_destroy : blacklisted
// UNSUPPORTED : hook_destroy_link : blacklisted
// UNSUPPORTED : hook_free : blacklisted
// UNSUPPORTED : hook_insert_before : blacklisted
// UNSUPPORTED : hook_prepend : blacklisted
// UNSUPPORTED : hook_unref : blacklisted
// UNSUPPORTED : hostname_is_ascii_encoded : blacklisted
// UNSUPPORTED : hostname_is_ip_address : blacklisted
// UNSUPPORTED : hostname_is_non_ascii : blacklisted
// UNSUPPORTED : hostname_to_ascii : blacklisted
// UNSUPPORTED : hostname_to_unicode : blacklisted
// UNSUPPORTED : iconv : blacklisted
// UNSUPPORTED : iconv_open : blacklisted
// UNSUPPORTED : idle_add : blacklisted
// UNSUPPORTED : idle_add_full : blacklisted
// UNSUPPORTED : idle_remove_by_data : blacklisted
// UNSUPPORTED : idle_source_new : blacklisted
// UNSUPPORTED : int64_equal : blacklisted
// UNSUPPORTED : int64_hash : blacklisted
// UNSUPPORTED : int_equal : blacklisted
// UNSUPPORTED : int_hash : blacklisted
// UNSUPPORTED : intern_static_string : blacklisted
// UNSUPPORTED : intern_string : blacklisted
// UNSUPPORTED : io_add_watch : blacklisted
// UNSUPPORTED : io_add_watch_full : blacklisted
// UNSUPPORTED : io_channel_error_from_errno : blacklisted
// UNSUPPORTED : io_channel_error_quark : blacklisted
// UNSUPPORTED : io_create_watch : blacklisted
// UNSUPPORTED : key_file_error_quark : blacklisted
// UNSUPPORTED : listenv : blacklisted
// UNSUPPORTED : locale_from_utf8 : blacklisted
// UNSUPPORTED : locale_to_utf8 : blacklisted
// UNSUPPORTED : log : blacklisted
// UNSUPPORTED : log_default_handler : blacklisted
// UNSUPPORTED : log_remove_handler : blacklisted
// UNSUPPORTED : log_set_always_fatal : blacklisted
// UNSUPPORTED : log_set_default_handler : blacklisted
// UNSUPPORTED : log_set_fatal_mask : blacklisted
// UNSUPPORTED : log_set_handler : blacklisted
// UNSUPPORTED : log_set_handler_full : blacklisted
// UNSUPPORTED : log_set_writer_func : blacklisted
// UNSUPPORTED : log_structured : blacklisted
// UNSUPPORTED : log_structured_array : blacklisted
// UNSUPPORTED : log_structured_standard : blacklisted
// UNSUPPORTED : log_variant : blacklisted
// UNSUPPORTED : log_writer_default : blacklisted
// UNSUPPORTED : log_writer_format_fields : blacklisted
// UNSUPPORTED : log_writer_is_journald : blacklisted
// UNSUPPORTED : log_writer_journald : blacklisted
// UNSUPPORTED : log_writer_standard_streams : blacklisted
// UNSUPPORTED : log_writer_supports_color : blacklisted
// UNSUPPORTED : logv : blacklisted
// UNSUPPORTED : main_context_default : blacklisted
// UNSUPPORTED : main_context_get_thread_default : blacklisted
// UNSUPPORTED : main_context_ref_thread_default : blacklisted
// UNSUPPORTED : main_current_source : blacklisted
// UNSUPPORTED : main_depth : blacklisted
// UNSUPPORTED : malloc : blacklisted
// UNSUPPORTED : malloc0 : blacklisted
// UNSUPPORTED : malloc0_n : blacklisted
// UNSUPPORTED : malloc_n : blacklisted
// UNSUPPORTED : markup_collect_attributes : blacklisted
// UNSUPPORTED : markup_error_quark : blacklisted
// UNSUPPORTED : markup_escape_text : blacklisted
// UNSUPPORTED : markup_printf_escaped : blacklisted
// UNSUPPORTED : markup_vprintf_escaped : blacklisted
// UNSUPPORTED : mem_is_system_malloc : blacklisted
// UNSUPPORTED : mem_profile : blacklisted
// UNSUPPORTED : mem_set_vtable : blacklisted
// UNSUPPORTED : memdup : blacklisted
// UNSUPPORTED : mkdir_with_parents : blacklisted
// UNSUPPORTED : mkdtemp : blacklisted
// UNSUPPORTED : mkdtemp_full : blacklisted
// UNSUPPORTED : mkstemp : blacklisted
// UNSUPPORTED : mkstemp_full : blacklisted
// UNSUPPORTED : nullify_pointer : blacklisted
// UNSUPPORTED : number_parser_error_quark : blacklisted
// UNSUPPORTED : on_error_query : blacklisted
// UNSUPPORTED : on_error_stack_trace : blacklisted
// UNSUPPORTED : once_init_enter : blacklisted
// UNSUPPORTED : once_init_leave : blacklisted
// UNSUPPORTED : option_error_quark : blacklisted
// UNSUPPORTED : parse_debug_string : blacklisted
// UNSUPPORTED : path_get_basename : blacklisted
// UNSUPPORTED : path_get_dirname : blacklisted
// UNSUPPORTED : path_is_absolute : blacklisted
// UNSUPPORTED : path_skip_root : blacklisted
// UNSUPPORTED : pattern_match : blacklisted
// UNSUPPORTED : pattern_match_simple : blacklisted
// UNSUPPORTED : pattern_match_string : blacklisted
// UNSUPPORTED : pointer_bit_lock : blacklisted
// UNSUPPORTED : pointer_bit_trylock : blacklisted
// UNSUPPORTED : pointer_bit_unlock : blacklisted
// UNSUPPORTED : poll : blacklisted
// UNSUPPORTED : prefix_error : blacklisted
// UNSUPPORTED : print : blacklisted
// UNSUPPORTED : printerr : blacklisted
// UNSUPPORTED : printf : blacklisted
// UNSUPPORTED : printf_string_upper_bound : blacklisted
// UNSUPPORTED : propagate_error : blacklisted
// UNSUPPORTED : propagate_prefixed_error : blacklisted
// UNSUPPORTED : ptr_array_find : blacklisted
// UNSUPPORTED : ptr_array_find_with_equal_func : blacklisted
// UNSUPPORTED : qsort_with_data : blacklisted
// UNSUPPORTED : quark_from_static_string : blacklisted
// UNSUPPORTED : quark_from_string : blacklisted
// UNSUPPORTED : quark_to_string : blacklisted
// UNSUPPORTED : quark_try_string : blacklisted
// UNSUPPORTED : random_double : blacklisted
// UNSUPPORTED : random_double_range : blacklisted
// UNSUPPORTED : random_int : blacklisted
// UNSUPPORTED : random_int_range : blacklisted
// UNSUPPORTED : random_set_seed : blacklisted
// UNSUPPORTED : rc_box_acquire : blacklisted
// UNSUPPORTED : rc_box_alloc : blacklisted
// UNSUPPORTED : rc_box_alloc0 : blacklisted
// UNSUPPORTED : rc_box_dup : blacklisted
// UNSUPPORTED : rc_box_get_size : blacklisted
// UNSUPPORTED : rc_box_release : blacklisted
// UNSUPPORTED : rc_box_release_full : blacklisted
// UNSUPPORTED : realloc : blacklisted
// UNSUPPORTED : realloc_n : blacklisted
// UNSUPPORTED : ref_count_compare : blacklisted
// UNSUPPORTED : ref_count_dec : blacklisted
// UNSUPPORTED : ref_count_inc : blacklisted
// UNSUPPORTED : ref_count_init : blacklisted
// UNSUPPORTED : ref_string_acquire : blacklisted
// UNSUPPORTED : ref_string_length : blacklisted
// UNSUPPORTED : ref_string_new : blacklisted
// UNSUPPORTED : ref_string_new_intern : blacklisted
// UNSUPPORTED : ref_string_new_len : blacklisted
// UNSUPPORTED : ref_string_release : blacklisted
// UNSUPPORTED : regex_check_replacement : blacklisted
// UNSUPPORTED : regex_error_quark : blacklisted
// UNSUPPORTED : regex_escape_nul : blacklisted
// UNSUPPORTED : regex_escape_string : blacklisted
// UNSUPPORTED : regex_match_simple : blacklisted
// UNSUPPORTED : regex_split_simple : blacklisted
// UNSUPPORTED : reload_user_special_dirs_cache : blacklisted
// UNSUPPORTED : return_if_fail_warning : blacklisted
// UNSUPPORTED : rmdir : blacklisted
// UNSUPPORTED : sequence_get : blacklisted
// UNSUPPORTED : sequence_insert_before : blacklisted
// UNSUPPORTED : sequence_move : blacklisted
// UNSUPPORTED : sequence_move_range : blacklisted
// UNSUPPORTED : sequence_range_get_midpoint : blacklisted
// UNSUPPORTED : sequence_remove : blacklisted
// UNSUPPORTED : sequence_remove_range : blacklisted
// UNSUPPORTED : sequence_set : blacklisted
// UNSUPPORTED : sequence_swap : blacklisted
// UNSUPPORTED : set_application_name : blacklisted
// UNSUPPORTED : set_error : blacklisted
// UNSUPPORTED : set_error_literal : blacklisted
// UNSUPPORTED : set_prgname : blacklisted
// UNSUPPORTED : set_print_handler : blacklisted
// UNSUPPORTED : set_printerr_handler : blacklisted
// UNSUPPORTED : setenv : blacklisted
// UNSUPPORTED : shell_error_quark : blacklisted
// UNSUPPORTED : shell_parse_argv : blacklisted
// UNSUPPORTED : shell_quote : blacklisted
// UNSUPPORTED : shell_unquote : blacklisted
// UNSUPPORTED : slice_alloc : blacklisted
// UNSUPPORTED : slice_alloc0 : blacklisted
// UNSUPPORTED : slice_copy : blacklisted
// UNSUPPORTED : slice_free1 : blacklisted
// UNSUPPORTED : slice_free_chain_with_offset : blacklisted
// UNSUPPORTED : slice_get_config : blacklisted
// UNSUPPORTED : slice_get_config_state : blacklisted
// UNSUPPORTED : slice_set_config : blacklisted
// UNSUPPORTED : snprintf : blacklisted
// UNSUPPORTED : source_remove : blacklisted
// UNSUPPORTED : source_remove_by_funcs_user_data : blacklisted
// UNSUPPORTED : source_remove_by_user_data : blacklisted
// UNSUPPORTED : source_set_name_by_id : blacklisted
// UNSUPPORTED : spaced_primes_closest : blacklisted
// UNSUPPORTED : spawn_async : blacklisted
// UNSUPPORTED : spawn_async_with_fds : blacklisted
// UNSUPPORTED : spawn_async_with_pipes : blacklisted
// UNSUPPORTED : spawn_check_exit_status : blacklisted
// UNSUPPORTED : spawn_close_pid : blacklisted
// UNSUPPORTED : spawn_command_line_async : blacklisted
// UNSUPPORTED : spawn_command_line_sync : blacklisted
// UNSUPPORTED : spawn_error_quark : blacklisted
// UNSUPPORTED : spawn_exit_error_quark : blacklisted
// UNSUPPORTED : spawn_sync : blacklisted
// UNSUPPORTED : sprintf : blacklisted
// UNSUPPORTED : stpcpy : blacklisted
// UNSUPPORTED : str_equal : blacklisted
// UNSUPPORTED : str_has_prefix : blacklisted
// UNSUPPORTED : str_has_suffix : blacklisted
// UNSUPPORTED : str_hash : blacklisted
// UNSUPPORTED : str_is_ascii : blacklisted
// UNSUPPORTED : str_match_string : blacklisted
// UNSUPPORTED : str_to_ascii : blacklisted
// UNSUPPORTED : str_tokenize_and_fold : blacklisted
// UNSUPPORTED : strcanon : blacklisted
// UNSUPPORTED : strcasecmp : blacklisted
// UNSUPPORTED : strchomp : blacklisted
// UNSUPPORTED : strchug : blacklisted
// UNSUPPORTED : strcmp0 : blacklisted
// UNSUPPORTED : strcompress : blacklisted
// UNSUPPORTED : strconcat : blacklisted
// UNSUPPORTED : strdelimit : blacklisted
// UNSUPPORTED : strdown : blacklisted
// UNSUPPORTED : strdup : blacklisted
// UNSUPPORTED : strdup_printf : blacklisted
// UNSUPPORTED : strdup_vprintf : blacklisted
// UNSUPPORTED : strdupv : blacklisted
// UNSUPPORTED : strerror : blacklisted
// UNSUPPORTED : strescape : blacklisted
// UNSUPPORTED : strfreev : blacklisted
// UNSUPPORTED : string_new : blacklisted
// UNSUPPORTED : string_new_len : blacklisted
// UNSUPPORTED : string_sized_new : blacklisted
// UNSUPPORTED : strip_context : blacklisted
// UNSUPPORTED : strjoin : blacklisted
// UNSUPPORTED : strjoinv : blacklisted
// UNSUPPORTED : strlcat : blacklisted
// UNSUPPORTED : strlcpy : blacklisted
// UNSUPPORTED : strncasecmp : blacklisted
// UNSUPPORTED : strndup : blacklisted
// UNSUPPORTED : strnfill : blacklisted
// UNSUPPORTED : strreverse : blacklisted
// UNSUPPORTED : strrstr : blacklisted
// UNSUPPORTED : strrstr_len : blacklisted
// UNSUPPORTED : strsignal : blacklisted
// UNSUPPORTED : strsplit : blacklisted
// UNSUPPORTED : strsplit_set : blacklisted
// UNSUPPORTED : strstr_len : blacklisted
// UNSUPPORTED : strtod : blacklisted
// UNSUPPORTED : strup : blacklisted
// UNSUPPORTED : strv_contains : blacklisted
// UNSUPPORTED : strv_equal : blacklisted
// UNSUPPORTED : strv_get_type : blacklisted
// UNSUPPORTED : strv_length : blacklisted
// UNSUPPORTED : test_add_data_func : blacklisted
// UNSUPPORTED : test_add_data_func_full : blacklisted
// UNSUPPORTED : test_add_func : blacklisted
// UNSUPPORTED : test_add_vtable : blacklisted
// UNSUPPORTED : test_assert_expected_messages_internal : blacklisted
// UNSUPPORTED : test_bug : blacklisted
// UNSUPPORTED : test_bug_base : blacklisted
// UNSUPPORTED : test_build_filename : blacklisted
// UNSUPPORTED : test_create_case : blacklisted
// UNSUPPORTED : test_create_suite : blacklisted
// UNSUPPORTED : test_expect_message : blacklisted
// UNSUPPORTED : test_fail : blacklisted
// UNSUPPORTED : test_failed : blacklisted
// UNSUPPORTED : test_get_dir : blacklisted
// UNSUPPORTED : test_get_filename : blacklisted
// UNSUPPORTED : test_get_root : blacklisted
// UNSUPPORTED : test_incomplete : blacklisted
// UNSUPPORTED : test_init : blacklisted
// UNSUPPORTED : test_log_set_fatal_handler : blacklisted
// UNSUPPORTED : test_log_type_name : blacklisted
// UNSUPPORTED : test_maximized_result : blacklisted
// UNSUPPORTED : test_message : blacklisted
// UNSUPPORTED : test_minimized_result : blacklisted
// UNSUPPORTED : test_queue_destroy : blacklisted
// UNSUPPORTED : test_queue_free : blacklisted
// UNSUPPORTED : test_rand_double : blacklisted
// UNSUPPORTED : test_rand_double_range : blacklisted
// UNSUPPORTED : test_rand_int : blacklisted
// UNSUPPORTED : test_rand_int_range : blacklisted
// UNSUPPORTED : test_run : blacklisted
// UNSUPPORTED : test_run_suite : blacklisted
// UNSUPPORTED : test_set_nonfatal_assertions : blacklisted
// UNSUPPORTED : test_skip : blacklisted
// UNSUPPORTED : test_subprocess : blacklisted
// UNSUPPORTED : test_summary : blacklisted
// UNSUPPORTED : test_timer_elapsed : blacklisted
// UNSUPPORTED : test_timer_last : blacklisted
// UNSUPPORTED : test_timer_start : blacklisted
// UNSUPPORTED : test_trap_assertions : blacklisted
// UNSUPPORTED : test_trap_fork : blacklisted
// UNSUPPORTED : test_trap_has_passed : blacklisted
// UNSUPPORTED : test_trap_reached_timeout : blacklisted
// UNSUPPORTED : test_trap_subprocess : blacklisted
// UNSUPPORTED : thread_error_quark : blacklisted
// UNSUPPORTED : thread_exit : blacklisted
// UNSUPPORTED : thread_pool_get_max_idle_time : blacklisted
// UNSUPPORTED : thread_pool_get_max_unused_threads : blacklisted
// UNSUPPORTED : thread_pool_get_num_unused_threads : blacklisted
// UNSUPPORTED : thread_pool_set_max_idle_time : blacklisted
// UNSUPPORTED : thread_pool_set_max_unused_threads : blacklisted
// UNSUPPORTED : thread_pool_stop_unused_threads : blacklisted
// UNSUPPORTED : thread_self : blacklisted
// UNSUPPORTED : thread_yield : blacklisted
// UNSUPPORTED : time_val_from_iso8601 : blacklisted
// UNSUPPORTED : timeout_add : blacklisted
// UNSUPPORTED : timeout_add_full : blacklisted
// UNSUPPORTED : timeout_add_seconds : blacklisted
// UNSUPPORTED : timeout_add_seconds_full : blacklisted
// UNSUPPORTED : timeout_source_new : blacklisted
// UNSUPPORTED : timeout_source_new_seconds : blacklisted
// UNSUPPORTED : trash_stack_height : blacklisted
// UNSUPPORTED : trash_stack_peek : blacklisted
// UNSUPPORTED : trash_stack_pop : blacklisted
// UNSUPPORTED : trash_stack_push : blacklisted
// UNSUPPORTED : try_malloc : blacklisted
// UNSUPPORTED : try_malloc0 : blacklisted
// UNSUPPORTED : try_malloc0_n : blacklisted
// UNSUPPORTED : try_malloc_n : blacklisted
// UNSUPPORTED : try_realloc : blacklisted
// UNSUPPORTED : try_realloc_n : blacklisted
// UNSUPPORTED : ucs4_to_utf16 : blacklisted
// UNSUPPORTED : ucs4_to_utf8 : blacklisted
// UNSUPPORTED : unichar_break_type : blacklisted
// UNSUPPORTED : unichar_combining_class : blacklisted
// UNSUPPORTED : unichar_compose : blacklisted
// UNSUPPORTED : unichar_decompose : blacklisted
// UNSUPPORTED : unichar_digit_value : blacklisted
// UNSUPPORTED : unichar_fully_decompose : blacklisted
// UNSUPPORTED : unichar_get_mirror_char : blacklisted
// UNSUPPORTED : unichar_get_script : blacklisted
// UNSUPPORTED : unichar_isalnum : blacklisted
// UNSUPPORTED : unichar_isalpha : blacklisted
// UNSUPPORTED : unichar_iscntrl : blacklisted
// UNSUPPORTED : unichar_isdefined : blacklisted
// UNSUPPORTED : unichar_isdigit : blacklisted
// UNSUPPORTED : unichar_isgraph : blacklisted
// UNSUPPORTED : unichar_islower : blacklisted
// UNSUPPORTED : unichar_ismark : blacklisted
// UNSUPPORTED : unichar_isprint : blacklisted
// UNSUPPORTED : unichar_ispunct : blacklisted
// UNSUPPORTED : unichar_isspace : blacklisted
// UNSUPPORTED : unichar_istitle : blacklisted
// UNSUPPORTED : unichar_isupper : blacklisted
// UNSUPPORTED : unichar_iswide : blacklisted
// UNSUPPORTED : unichar_iswide_cjk : blacklisted
// UNSUPPORTED : unichar_isxdigit : blacklisted
// UNSUPPORTED : unichar_iszerowidth : blacklisted
// UNSUPPORTED : unichar_to_utf8 : blacklisted
// UNSUPPORTED : unichar_tolower : blacklisted
// UNSUPPORTED : unichar_totitle : blacklisted
// UNSUPPORTED : unichar_toupper : blacklisted
// UNSUPPORTED : unichar_type : blacklisted
// UNSUPPORTED : unichar_validate : blacklisted
// UNSUPPORTED : unichar_xdigit_value : blacklisted
// UNSUPPORTED : unicode_canonical_decomposition : blacklisted
// UNSUPPORTED : unicode_canonical_ordering : blacklisted
// UNSUPPORTED : unicode_script_from_iso15924 : blacklisted
// UNSUPPORTED : unicode_script_to_iso15924 : blacklisted
// UNSUPPORTED : unix_error_quark : blacklisted
// UNSUPPORTED : unix_fd_add : blacklisted
// UNSUPPORTED : unix_fd_add_full : blacklisted
// UNSUPPORTED : unix_fd_source_new : blacklisted
// UNSUPPORTED : unix_open_pipe : blacklisted
// UNSUPPORTED : unix_set_fd_nonblocking : blacklisted
// UNSUPPORTED : unix_signal_add : blacklisted
// UNSUPPORTED : unix_signal_add_full : blacklisted
// UNSUPPORTED : unix_signal_source_new : blacklisted
// UNSUPPORTED : unlink : blacklisted
// UNSUPPORTED : unsetenv : blacklisted
// UNSUPPORTED : uri_escape_string : blacklisted
// UNSUPPORTED : uri_list_extract_uris : blacklisted
// UNSUPPORTED : uri_parse_scheme : blacklisted
// UNSUPPORTED : uri_unescape_segment : blacklisted
// UNSUPPORTED : uri_unescape_string : blacklisted
// UNSUPPORTED : usleep : blacklisted
// UNSUPPORTED : utf16_to_ucs4 : blacklisted
// UNSUPPORTED : utf16_to_utf8 : blacklisted
// UNSUPPORTED : utf8_casefold : blacklisted
// UNSUPPORTED : utf8_collate : blacklisted
// UNSUPPORTED : utf8_collate_key : blacklisted
// UNSUPPORTED : utf8_collate_key_for_filename : blacklisted
// UNSUPPORTED : utf8_find_next_char : blacklisted
// UNSUPPORTED : utf8_find_prev_char : blacklisted
// UNSUPPORTED : utf8_get_char : blacklisted
// UNSUPPORTED : utf8_get_char_validated : blacklisted
// UNSUPPORTED : utf8_make_valid : blacklisted
// UNSUPPORTED : utf8_normalize : blacklisted
// UNSUPPORTED : utf8_offset_to_pointer : blacklisted
// UNSUPPORTED : utf8_pointer_to_offset : blacklisted
// UNSUPPORTED : utf8_prev_char : blacklisted
// UNSUPPORTED : utf8_strchr : blacklisted
// UNSUPPORTED : utf8_strdown : blacklisted
// UNSUPPORTED : utf8_strlen : blacklisted
// UNSUPPORTED : utf8_strncpy : blacklisted
// UNSUPPORTED : utf8_strrchr : blacklisted
// UNSUPPORTED : utf8_strreverse : blacklisted
// UNSUPPORTED : utf8_strup : blacklisted
// UNSUPPORTED : utf8_substring : blacklisted
// UNSUPPORTED : utf8_to_ucs4 : blacklisted
// UNSUPPORTED : utf8_to_ucs4_fast : blacklisted
// UNSUPPORTED : utf8_to_utf16 : blacklisted
// UNSUPPORTED : utf8_validate : blacklisted
// UNSUPPORTED : utf8_validate_len : blacklisted
// UNSUPPORTED : uuid_string_is_valid : blacklisted
// UNSUPPORTED : uuid_string_random : blacklisted
// UNSUPPORTED : variant_get_gtype : blacklisted
// UNSUPPORTED : variant_is_object_path : blacklisted
// UNSUPPORTED : variant_is_signature : blacklisted
// UNSUPPORTED : variant_parse : blacklisted
// UNSUPPORTED : variant_parse_error_print_context : blacklisted
// UNSUPPORTED : variant_parse_error_quark : blacklisted
// UNSUPPORTED : variant_parser_get_error_quark : blacklisted
// UNSUPPORTED : variant_type_checked_ : blacklisted
// UNSUPPORTED : variant_type_string_get_depth_ : blacklisted
// UNSUPPORTED : variant_type_string_is_valid : blacklisted
// UNSUPPORTED : variant_type_string_scan : blacklisted
// UNSUPPORTED : vasprintf : blacklisted
// UNSUPPORTED : vfprintf : blacklisted
// UNSUPPORTED : vprintf : blacklisted
// UNSUPPORTED : vsnprintf : blacklisted
// UNSUPPORTED : vsprintf : blacklisted
// UNSUPPORTED : warn_message : blacklisted

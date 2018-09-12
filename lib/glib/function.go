package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// AsciiDigitValue is a wrapper around the C function g_ascii_digit_value.
func AsciiDigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_digit_value(c_c)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_ascii_dtostr : no return type

// Unsupported : g_ascii_formatd : no return type

// AsciiStrcasecmp is a wrapper around the C function g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_ascii_strcasecmp(c_s1, c_s2)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_ascii_strdown : no return type

// AsciiStrncasecmp is a wrapper around the C function g_ascii_strncasecmp.
func AsciiStrncasecmp(s1 string, s2 string, n uint64) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	c_n := (C.gsize)(n)

	retC := C.g_ascii_strncasecmp(c_s1, c_s2, c_n)
	retGo :=
		(int32)(retC)

	return retGo
}

// AsciiStrtod is a wrapper around the C function g_ascii_strtod.
func AsciiStrtod(nptr string) float64 {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_ascii_strtod(c_nptr, &c_endptr)
	retGo :=
		(float64)(retC)

	return retGo
}

// Unsupported : g_ascii_strup : no return type

// AsciiTolower is a wrapper around the C function g_ascii_tolower.
func AsciiTolower(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_tolower(c_c)
	retGo :=
		(rune)(retC)

	return retGo
}

// AsciiToupper is a wrapper around the C function g_ascii_toupper.
func AsciiToupper(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_toupper(c_c)
	retGo :=
		(rune)(retC)

	return retGo
}

// AsciiXdigitValue is a wrapper around the C function g_ascii_xdigit_value.
func AsciiXdigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_xdigit_value(c_c)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_assert_warning : unsupported parameter line : no param type

// Unsupported : g_assertion_message : unsupported parameter line : no param type

// Unsupported : g_assertion_message_cmpnum : unsupported parameter line : no param type

// Unsupported : g_assertion_message_cmpstr : unsupported parameter line : no param type

// Unsupported : g_assertion_message_error : unsupported parameter line : no param type

// Unsupported : g_assertion_message_expr : unsupported parameter line : no param type

// Unsupported : g_atexit : unsupported parameter func : no param type

// Unsupported : g_basename : unsupported parameter file_name : no param type

// BitNthLsf is a wrapper around the C function g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_lsf(c_mask, c_nth_bit)
	retGo :=
		(int32)(retC)

	return retGo
}

// BitNthMsf is a wrapper around the C function g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_msf(c_mask, c_nth_bit)
	retGo :=
		(int32)(retC)

	return retGo
}

// BitStorage is a wrapper around the C function g_bit_storage.
func BitStorage(number uint64) uint32 {
	c_number := (C.gulong)(number)

	retC := C.g_bit_storage(c_number)
	retGo :=
		(uint32)(retC)

	return retGo
}

// Unsupported : g_bookmark_file_error_quark : no return type

// Unsupported : g_build_filename : unsupported parameter first_element : no param type

// Unsupported : g_build_path : unsupported parameter separator : no param type

// Unsupported : g_byte_array_free : unsupported parameter array : no param type

// Unsupported : g_byte_array_new : no return type

// Unsupported : g_clear_error : no return type

// Unsupported : g_convert : unsupported parameter str : no param type

// Unsupported : g_convert_error_quark : no return type

// Unsupported : g_convert_with_fallback : unsupported parameter str : no param type

// Unsupported : g_convert_with_iconv : unsupported parameter str : no param type

// Unsupported : g_datalist_clear : unsupported parameter datalist : no param type

// Unsupported : g_datalist_foreach : unsupported parameter datalist : no param type

// Unsupported : g_datalist_get_data : unsupported parameter datalist : no param type

// Unsupported : g_datalist_id_get_data : unsupported parameter datalist : no param type

// Unsupported : g_datalist_id_remove_no_notify : unsupported parameter datalist : no param type

// Unsupported : g_datalist_id_set_data_full : unsupported parameter datalist : no param type

// Unsupported : g_datalist_init : unsupported parameter datalist : no param type

// Unsupported : g_dataset_destroy : unsupported parameter dataset_location : no param type

// Unsupported : g_dataset_foreach : unsupported parameter dataset_location : no param type

// Unsupported : g_dataset_id_get_data : unsupported parameter dataset_location : no param type

// Unsupported : g_dataset_id_remove_no_notify : unsupported parameter dataset_location : no param type

// Unsupported : g_dataset_id_set_data_full : unsupported parameter dataset_location : no param type

// Unsupported : g_date_get_days_in_month : unsupported parameter month : no param type

// Unsupported : g_date_get_monday_weeks_in_year : unsupported parameter year : no param type

// Unsupported : g_date_get_sunday_weeks_in_year : unsupported parameter year : no param type

// Unsupported : g_date_is_leap_year : unsupported parameter year : no param type

// Unsupported : g_date_strftime : unsupported parameter date : no param type

// Unsupported : g_date_valid_day : unsupported parameter day : no param type

// Unsupported : g_date_valid_dmy : unsupported parameter day : no param type

// Unsupported : g_date_valid_julian : no return type

// Unsupported : g_date_valid_month : unsupported parameter month : no param type

// Unsupported : g_date_valid_weekday : unsupported parameter weekday : no param type

// Unsupported : g_date_valid_year : unsupported parameter year : no param type

// Unsupported : g_direct_equal : unsupported parameter v1 : no param type

// Unsupported : g_direct_hash : unsupported parameter v : no param type

// Unsupported : g_file_error_from_errno : no return type

// Unsupported : g_file_error_quark : no return type

// Unsupported : g_file_get_contents : unsupported parameter filename : no param type

// Unsupported : g_file_open_tmp : unsupported parameter tmpl : no param type

// Unsupported : g_file_test : unsupported parameter filename : no param type

// Unsupported : g_filename_from_uri : no return type

// Unsupported : g_filename_from_utf8 : unsupported parameter bytes_read : no param type

// Unsupported : g_filename_to_uri : unsupported parameter filename : no param type

// Unsupported : g_filename_to_utf8 : unsupported parameter opsysstring : no param type

// Unsupported : g_find_program_in_path : unsupported parameter program : no param type

// Unsupported : g_free : no return type

// Unsupported : g_get_charset : no return type

// Unsupported : g_get_codeset : no return type

// Unsupported : g_get_current_dir : no return type

// Unsupported : g_get_current_time : unsupported parameter result : no param type

// Unsupported : g_get_home_dir : no return type

// Unsupported : g_get_prgname : no return type

// Unsupported : g_get_real_name : no return type

// Unsupported : g_get_tmp_dir : no return type

// Unsupported : g_get_user_name : no return type

// Unsupported : g_getenv : unsupported parameter variable : no param type

// Unsupported : g_hash_table_destroy : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_insert : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_lookup : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_lookup_extended : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_remove : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_replace : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_size : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_steal : unsupported parameter hash_table : no param type

// Unsupported : g_hook_destroy : unsupported parameter hook_list : no param type

// Unsupported : g_hook_destroy_link : unsupported parameter hook_list : no param type

// Unsupported : g_hook_free : unsupported parameter hook_list : no param type

// Unsupported : g_hook_insert_before : unsupported parameter hook_list : no param type

// Unsupported : g_hook_prepend : unsupported parameter hook_list : no param type

// Unsupported : g_hook_unref : unsupported parameter hook_list : no param type

// Unsupported : g_iconv : unsupported parameter converter : no param type

// Unsupported : g_iconv_open : no return type

// Unsupported : g_idle_add : unsupported parameter function : no param type

// Unsupported : g_idle_add_full : unsupported parameter function : no param type

// Unsupported : g_idle_remove_by_data : no return type

// Unsupported : g_idle_source_new : no return type

// Unsupported : g_int_equal : unsupported parameter v1 : no param type

// Unsupported : g_int_hash : unsupported parameter v : no param type

// Unsupported : g_io_add_watch : unsupported parameter channel : no param type

// Unsupported : g_io_add_watch_full : unsupported parameter channel : no param type

// Unsupported : g_io_channel_error_from_errno : no return type

// Unsupported : g_io_channel_error_quark : no return type

// Unsupported : g_io_create_watch : unsupported parameter channel : no param type

// Unsupported : g_key_file_error_quark : no return type

// Unsupported : g_locale_from_utf8 : unsupported parameter bytes_read : no param type

// Unsupported : g_locale_to_utf8 : unsupported parameter opsysstring : no param type

// Unsupported : g_log : unsupported parameter log_level : no param type

// Unsupported : g_log_default_handler : unsupported parameter log_level : no param type

// Unsupported : g_log_remove_handler : no return type

// Unsupported : g_log_set_always_fatal : unsupported parameter fatal_mask : no param type

// Unsupported : g_log_set_fatal_mask : unsupported parameter fatal_mask : no param type

// Unsupported : g_log_set_handler : unsupported parameter log_levels : no param type

// Unsupported : g_log_structured_standard : unsupported parameter log_level : no param type

// Unsupported : g_logv : unsupported parameter log_level : no param type

// Unsupported : g_main_context_default : no return type

// MainDepth is a wrapper around the C function g_main_depth.
func MainDepth() int32 {
	retC := C.g_main_depth()
	retGo :=
		(int32)(retC)

	return retGo
}

// Malloc is a wrapper around the C function g_malloc.
func Malloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc(c_n_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// Malloc0 is a wrapper around the C function g_malloc0.
func Malloc0(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc0(c_n_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// Unsupported : g_markup_error_quark : no return type

// Unsupported : g_markup_escape_text : no return type

// Unsupported : g_mem_is_system_malloc : no return type

// Unsupported : g_mem_profile : no return type

// Unsupported : g_mem_set_vtable : unsupported parameter vtable : no param type

// Unsupported : g_memdup : unsupported parameter mem : no param type

// Unsupported : g_mkstemp : unsupported parameter tmpl : no param type

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : no param type

// Unsupported : g_number_parser_error_quark : no return type

// Unsupported : g_on_error_query : no return type

// Unsupported : g_on_error_stack_trace : no return type

// Unsupported : g_option_error_quark : no return type

// Unsupported : g_parse_debug_string : unsupported parameter keys : no param type

// Unsupported : g_path_get_basename : unsupported parameter file_name : no param type

// Unsupported : g_path_get_dirname : unsupported parameter file_name : no param type

// Unsupported : g_path_is_absolute : unsupported parameter file_name : no param type

// Unsupported : g_path_skip_root : unsupported parameter file_name : no param type

// Unsupported : g_pattern_match : unsupported parameter pspec : no param type

// Unsupported : g_pattern_match_simple : no return type

// Unsupported : g_pattern_match_string : unsupported parameter pspec : no param type

// Unsupported : g_print : unsupported parameter ... : varargs

// Unsupported : g_printerr : unsupported parameter ... : varargs

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no param type

// Unsupported : g_propagate_error : unsupported parameter dest : no param type

// Unsupported : g_qsort_with_data : unsupported parameter pbase : no param type

// Unsupported : g_quark_from_static_string : no return type

// Unsupported : g_quark_from_string : no return type

// Unsupported : g_quark_to_string : unsupported parameter quark : no param type

// Unsupported : g_quark_try_string : no return type

// RandomDouble is a wrapper around the C function g_random_double.
func RandomDouble() float64 {
	retC := C.g_random_double()
	retGo :=
		(float64)(retC)

	return retGo
}

// RandomDoubleRange is a wrapper around the C function g_random_double_range.
func RandomDoubleRange(begin float64, end float64) float64 {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	retC := C.g_random_double_range(c_begin, c_end)
	retGo :=
		(float64)(retC)

	return retGo
}

// RandomInt is a wrapper around the C function g_random_int.
func RandomInt() uint32 {
	retC := C.g_random_int()
	retGo :=
		(uint32)(retC)

	return retGo
}

// RandomIntRange is a wrapper around the C function g_random_int_range.
func RandomIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_random_int_range(c_begin, c_end)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_random_set_seed : no return type

// Realloc is a wrapper around the C function g_realloc.
func Realloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_realloc(c_mem, c_n_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// Unsupported : g_regex_error_quark : no return type

// Unsupported : g_return_if_fail_warning : no return type

// Unsupported : g_set_error : unsupported parameter err : no param type

// Unsupported : g_set_prgname : no return type

// Unsupported : g_set_print_handler : unsupported parameter func : no param type

// Unsupported : g_set_printerr_handler : unsupported parameter func : no param type

// Unsupported : g_shell_error_quark : no return type

// Unsupported : g_shell_parse_argv : unsupported parameter command_line : no param type

// Unsupported : g_shell_quote : unsupported parameter unquoted_string : no param type

// Unsupported : g_shell_unquote : unsupported parameter quoted_string : no param type

// Unsupported : g_slice_get_config : unsupported parameter ckey : no param type

// Unsupported : g_slice_get_config_state : unsupported parameter ckey : no param type

// Unsupported : g_slice_set_config : unsupported parameter ckey : no param type

// Unsupported : g_snprintf : unsupported parameter ... : varargs

// Unsupported : g_source_remove : no return type

// Unsupported : g_source_remove_by_funcs_user_data : unsupported parameter funcs : no param type

// Unsupported : g_source_remove_by_user_data : no return type

// SpacedPrimesClosest is a wrapper around the C function g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	c_num := (C.guint)(num)

	retC := C.g_spaced_primes_closest(c_num)
	retGo :=
		(uint32)(retC)

	return retGo
}

// Unsupported : g_spawn_async : unsupported parameter working_directory : no param type

// Unsupported : g_spawn_async_with_pipes : unsupported parameter working_directory : no param type

// Unsupported : g_spawn_close_pid : unsupported parameter pid : no param type

// Unsupported : g_spawn_command_line_async : unsupported parameter command_line : no param type

// Unsupported : g_spawn_command_line_sync : unsupported parameter command_line : no param type

// Unsupported : g_spawn_error_quark : no return type

// Unsupported : g_spawn_exit_error_quark : no return type

// Unsupported : g_spawn_sync : unsupported parameter working_directory : no param type

// Unsupported : g_stpcpy : no return type

// Unsupported : g_str_equal : unsupported parameter v1 : no param type

// Unsupported : g_str_hash : unsupported parameter v : no param type

// Unsupported : g_strcanon : no return type

// Strcasecmp is a wrapper around the C function g_strcasecmp.
func Strcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_strcasecmp(c_s1, c_s2)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_strchomp : no return type

// Unsupported : g_strchug : no return type

// Unsupported : g_strcompress : no return type

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Unsupported : g_strdelimit : no return type

// Unsupported : g_strdown : no return type

// Unsupported : g_strdup : no return type

// Unsupported : g_strdup_printf : unsupported parameter ... : varargs

// Unsupported : g_strdup_vprintf : unsupported parameter args : no param type

// Unsupported : g_strdupv : unsupported parameter str_array : in for string with indirection level of 2

// Unsupported : g_strerror : no return type

// Unsupported : g_strescape : no return type

// Unsupported : g_strfreev : unsupported parameter str_array : in for string with indirection level of 2

// Unsupported : g_string_new : no return type

// Unsupported : g_string_new_len : no return type

// Unsupported : g_string_sized_new : no return type

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in for string with indirection level of 2

// Strlcat is a wrapper around the C function g_strlcat.
func Strlcat(dest string, src string, destSize uint64) uint64 {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_dest_size := (C.gsize)(destSize)

	retC := C.g_strlcat(c_dest, c_src, c_dest_size)
	retGo :=
		(uint64)(retC)

	return retGo
}

// Strlcpy is a wrapper around the C function g_strlcpy.
func Strlcpy(dest string, src string, destSize uint64) uint64 {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_dest_size := (C.gsize)(destSize)

	retC := C.g_strlcpy(c_dest, c_src, c_dest_size)
	retGo :=
		(uint64)(retC)

	return retGo
}

// Strncasecmp is a wrapper around the C function g_strncasecmp.
func Strncasecmp(s1 string, s2 string, n uint32) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	c_n := (C.guint)(n)

	retC := C.g_strncasecmp(c_s1, c_s2, c_n)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_strndup : no return type

// Unsupported : g_strnfill : no return type

// Unsupported : g_strreverse : no return type

// Unsupported : g_strrstr : no return type

// Unsupported : g_strrstr_len : no return type

// Unsupported : g_strsignal : no return type

// Unsupported : g_strsplit : no return type

// Unsupported : g_strstr_len : no return type

// Strtod is a wrapper around the C function g_strtod.
func Strtod(nptr string) float64 {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_strtod(c_nptr, &c_endptr)
	retGo :=
		(float64)(retC)

	return retGo
}

// Unsupported : g_strup : no return type

// Unsupported : g_strv_get_type : no return type

// Unsupported : g_test_add_vtable : unsupported parameter test_data : no param type

// Unsupported : g_test_assert_expected_messages_internal : unsupported parameter line : no param type

// Unsupported : g_test_log_type_name : unsupported parameter log_type : no param type

// Unsupported : g_test_trap_assertions : unsupported parameter line : no param type

// Unsupported : g_thread_error_quark : no return type

// Unsupported : g_thread_exit : no return type

// ThreadPoolGetMaxUnusedThreads is a wrapper around the C function g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() int32 {
	retC := C.g_thread_pool_get_max_unused_threads()
	retGo :=
		(int32)(retC)

	return retGo
}

// ThreadPoolGetNumUnusedThreads is a wrapper around the C function g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() uint32 {
	retC := C.g_thread_pool_get_num_unused_threads()
	retGo :=
		(uint32)(retC)

	return retGo
}

// Unsupported : g_thread_pool_set_max_unused_threads : no return type

// Unsupported : g_thread_pool_stop_unused_threads : no return type

// Unsupported : g_thread_self : no return type

// Unsupported : g_thread_yield : no return type

// Unsupported : g_timeout_add : unsupported parameter function : no param type

// Unsupported : g_timeout_add_full : unsupported parameter function : no param type

// Unsupported : g_timeout_source_new : no return type

// Unsupported : g_trash_stack_height : unsupported parameter stack_p : no param type

// Unsupported : g_trash_stack_peek : unsupported parameter stack_p : no param type

// Unsupported : g_trash_stack_pop : unsupported parameter stack_p : no param type

// Unsupported : g_trash_stack_push : unsupported parameter stack_p : no param type

// TryMalloc is a wrapper around the C function g_try_malloc.
func TryMalloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_malloc(c_n_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// TryRealloc is a wrapper around the C function g_try_realloc.
func TryRealloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_realloc(c_mem, c_n_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// Unsupported : g_ucs4_to_utf16 : unsupported parameter str : no param type

// Unsupported : g_ucs4_to_utf8 : unsupported parameter str : no param type

// Unsupported : g_unichar_break_type : no return type

// UnicharDigitValue is a wrapper around the C function g_unichar_digit_value.
func UnicharDigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_digit_value(c_c)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_unichar_isalnum : no return type

// Unsupported : g_unichar_isalpha : no return type

// Unsupported : g_unichar_iscntrl : no return type

// Unsupported : g_unichar_isdefined : no return type

// Unsupported : g_unichar_isdigit : no return type

// Unsupported : g_unichar_isgraph : no return type

// Unsupported : g_unichar_islower : no return type

// Unsupported : g_unichar_isprint : no return type

// Unsupported : g_unichar_ispunct : no return type

// Unsupported : g_unichar_isspace : no return type

// Unsupported : g_unichar_istitle : no return type

// Unsupported : g_unichar_isupper : no return type

// Unsupported : g_unichar_iswide : no return type

// Unsupported : g_unichar_isxdigit : no return type

// UnicharToUtf8 is a wrapper around the C function g_unichar_to_utf8.
func UnicharToUtf8(c rune) int32 {
	c_c := (C.gunichar)(c)

	var c_outbuf C.gchar

	retC := C.g_unichar_to_utf8(c_c, &c_outbuf)
	retGo :=
		(int32)(retC)

	return retGo
}

// UnicharTolower is a wrapper around the C function g_unichar_tolower.
func UnicharTolower(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_tolower(c_c)
	retGo :=
		(rune)(retC)

	return retGo
}

// UnicharTotitle is a wrapper around the C function g_unichar_totitle.
func UnicharTotitle(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_totitle(c_c)
	retGo :=
		(rune)(retC)

	return retGo
}

// UnicharToupper is a wrapper around the C function g_unichar_toupper.
func UnicharToupper(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_toupper(c_c)
	retGo :=
		(rune)(retC)

	return retGo
}

// Unsupported : g_unichar_type : no return type

// Unsupported : g_unichar_validate : no return type

// UnicharXdigitValue is a wrapper around the C function g_unichar_xdigit_value.
func UnicharXdigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_xdigit_value(c_c)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_unicode_canonical_decomposition : unsupported parameter result_len : no param type

// Unsupported : g_unicode_canonical_ordering : unsupported parameter string : no param type

// Unsupported : g_unix_error_quark : no return type

// Unsupported : g_usleep : no return type

// Unsupported : g_utf16_to_ucs4 : unsupported parameter str : no param type

// Unsupported : g_utf16_to_utf8 : unsupported parameter str : no param type

// Unsupported : g_utf8_casefold : no return type

// Utf8Collate is a wrapper around the C function g_utf8_collate.
func Utf8Collate(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_utf8_collate(c_str1, c_str2)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_utf8_collate_key : no return type

// Unsupported : g_utf8_find_next_char : no return type

// Unsupported : g_utf8_find_prev_char : no return type

// Utf8GetChar is a wrapper around the C function g_utf8_get_char.
func Utf8GetChar(p string) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_get_char(c_p)
	retGo :=
		(rune)(retC)

	return retGo
}

// Utf8GetCharValidated is a wrapper around the C function g_utf8_get_char_validated.
func Utf8GetCharValidated(p string, maxLen int64) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max_len := (C.gssize)(maxLen)

	retC := C.g_utf8_get_char_validated(c_p, c_max_len)
	retGo :=
		(rune)(retC)

	return retGo
}

// Unsupported : g_utf8_normalize : unsupported parameter mode : no param type

// Unsupported : g_utf8_offset_to_pointer : no return type

// Utf8PointerToOffset is a wrapper around the C function g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) int64 {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_pos := C.CString(pos)
	defer C.free(unsafe.Pointer(c_pos))

	retC := C.g_utf8_pointer_to_offset(c_str, c_pos)
	retGo :=
		(int64)(retC)

	return retGo
}

// Unsupported : g_utf8_prev_char : no return type

// Unsupported : g_utf8_strchr : no return type

// Unsupported : g_utf8_strdown : no return type

// Utf8Strlen is a wrapper around the C function g_utf8_strlen.
func Utf8Strlen(p string, max int64) int64 {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max := (C.gssize)(max)

	retC := C.g_utf8_strlen(c_p, c_max)
	retGo :=
		(int64)(retC)

	return retGo
}

// Unsupported : g_utf8_strncpy : no return type

// Unsupported : g_utf8_strrchr : no return type

// Unsupported : g_utf8_strup : no return type

// Unsupported : g_utf8_to_ucs4 : unsupported parameter items_read : no param type

// Unsupported : g_utf8_to_ucs4_fast : unsupported parameter items_written : no param type

// Unsupported : g_utf8_to_utf16 : unsupported parameter items_read : no param type

// Unsupported : g_utf8_validate : unsupported parameter str : no param type

// Unsupported : g_variant_get_gtype : no return type

// Unsupported : g_variant_parse : unsupported parameter type : no param type

// Unsupported : g_variant_parse_error_quark : no return type

// Unsupported : g_variant_parser_get_error_quark : no return type

// Unsupported : g_variant_type_checked_ : no return type

// Unsupported : g_variant_type_string_is_valid : no return type

// Unsupported : g_vsnprintf : unsupported parameter args : no param type

// Unsupported : g_warn_message : unsupported parameter line : no param type

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

// AsciiDigitValue is a wrapper around the C function g_ascii_digit_value.
func AsciiDigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_digit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_ascii_dtostr

// Blacklisted : g_ascii_formatd

// Blacklisted : g_ascii_strcasecmp

// Blacklisted : g_ascii_strdown

// Blacklisted : g_ascii_strncasecmp

// Blacklisted : g_ascii_strtod

// AsciiStrup is a wrapper around the C function g_ascii_strup.
func AsciiStrup(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strup(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_ascii_tolower

// Blacklisted : g_ascii_toupper

// Blacklisted : g_ascii_xdigit_value

// Blacklisted : g_assert_warning

// Blacklisted : g_assertion_message

// Unsupported : g_assertion_message_cmpnum : unsupported parameter arg1 : no type generator for long double (long double) for param arg1

// Blacklisted : g_assertion_message_cmpstr

// Blacklisted : g_assertion_message_error

// Blacklisted : g_assertion_message_expr

// Unsupported : g_atexit : unsupported parameter func : no type generator for VoidFunc (GVoidFunc) for param func

// Blacklisted : g_basename

// Blacklisted : g_bit_nth_lsf

// Blacklisted : g_bit_nth_msf

// Blacklisted : g_bit_storage

// Blacklisted : g_bookmark_file_error_quark

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Blacklisted : g_byte_array_free

// Unsupported : g_byte_array_new : array return type :

// Blacklisted : g_clear_error

// Unsupported : g_convert : array return type :

// Blacklisted : g_convert_error_quark

// Unsupported : g_convert_with_fallback : array return type :

// Unsupported : g_convert_with_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Blacklisted : g_datalist_clear

// Unsupported : g_datalist_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// Unsupported : g_datalist_get_data : no return generator

// Unsupported : g_datalist_id_get_data : no return generator

// Unsupported : g_datalist_id_remove_no_notify : no return generator

// Unsupported : g_datalist_id_set_data_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_datalist_init

// Unsupported : g_dataset_destroy : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_foreach : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_get_data : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_remove_no_notify : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_set_data_full : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

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

// Unsupported : g_direct_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_direct_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Blacklisted : g_file_error_from_errno

// Blacklisted : g_file_error_quark

// Unsupported : g_file_get_contents : unsupported parameter contents : output array param contents

// FileOpenTmp is a wrapper around the C function g_file_open_tmp.
func FileOpenTmp(tmpl string) (int32, string, error) {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var c_name_used *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_open_tmp(c_tmpl, &c_name_used, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	nameUsed := C.GoString(c_name_used)
	defer C.free(unsafe.Pointer(c_name_used))

	return retGo, nameUsed, goError
}

// Blacklisted : g_file_test

// FilenameFromUri is a wrapper around the C function g_filename_from_uri.
func FilenameFromUri(uri string) (string, string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_hostname *C.gchar

	var cThrowableError *C.GError

	retC := C.g_filename_from_uri(c_uri, &c_hostname, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	hostname := C.GoString(c_hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	return retGo, hostname, goError
}

// Blacklisted : g_filename_from_utf8

// Blacklisted : g_filename_to_uri

// Blacklisted : g_filename_to_utf8

// Blacklisted : g_find_program_in_path

// Unsupported : g_free : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Blacklisted : g_get_charset

// Blacklisted : g_get_codeset

// Blacklisted : g_get_current_dir

// Blacklisted : g_get_current_time

// Blacklisted : g_get_home_dir

// Blacklisted : g_get_prgname

// Blacklisted : g_get_real_name

// Blacklisted : g_get_tmp_dir

// Blacklisted : g_get_user_name

// Blacklisted : g_getenv

// Blacklisted : g_hash_table_destroy

// Unsupported : g_hash_table_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key

// Unsupported : g_hash_table_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Blacklisted : g_hash_table_size

// Unsupported : g_hash_table_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Blacklisted : g_hook_destroy

// Blacklisted : g_hook_destroy_link

// Blacklisted : g_hook_free

// Blacklisted : g_hook_insert_before

// Blacklisted : g_hook_prepend

// Blacklisted : g_hook_unref

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_remove_by_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_idle_source_new

// Unsupported : g_int_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_int_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_io_add_watch : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// Unsupported : g_io_add_watch_full : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// Blacklisted : g_io_channel_error_from_errno

// Blacklisted : g_io_channel_error_quark

// Blacklisted : g_io_create_watch

// Blacklisted : g_key_file_error_quark

// Unsupported : g_locale_from_utf8 : array return type :

// Blacklisted : g_locale_to_utf8

// Blacklisted : g_log

// Unsupported : g_log_default_handler : unsupported parameter unused_data : no type generator for gpointer (gpointer) for param unused_data

// Blacklisted : g_log_remove_handler

// Blacklisted : g_log_set_always_fatal

// Blacklisted : g_log_set_fatal_mask

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Blacklisted : g_log_structured_standard

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : g_main_context_default

// Blacklisted : g_main_depth

// Unsupported : g_malloc : no return generator

// Unsupported : g_malloc0 : no return generator

// Blacklisted : g_markup_error_quark

// Blacklisted : g_markup_escape_text

// Blacklisted : g_mem_is_system_malloc

// Blacklisted : g_mem_profile

// Blacklisted : g_mem_set_vtable

// Unsupported : g_memdup : unsupported parameter mem : no type generator for gpointer (gconstpointer) for param mem

// Blacklisted : g_mkstemp

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : no type generator for gpointer (gpointer*) for param nullify_location

// Blacklisted : g_number_parser_error_quark

// Blacklisted : g_on_error_query

// Blacklisted : g_on_error_stack_trace

// Blacklisted : g_option_error_quark

// Unsupported : g_parse_debug_string : unsupported parameter keys :

// Blacklisted : g_path_get_basename

// Blacklisted : g_path_get_dirname

// Blacklisted : g_path_is_absolute

// Blacklisted : g_path_skip_root

// Blacklisted : g_pattern_match

// Blacklisted : g_pattern_match_simple

// Blacklisted : g_pattern_match_string

// Blacklisted : g_print

// Blacklisted : g_printerr

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : g_propagate_error

// Unsupported : g_qsort_with_data : unsupported parameter pbase : no type generator for gpointer (gconstpointer) for param pbase

// Blacklisted : g_quark_from_static_string

// QuarkFromString is a wrapper around the C function g_quark_from_string.
func QuarkFromString(string_ string) Quark {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_from_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// QuarkToString is a wrapper around the C function g_quark_to_string.
func QuarkToString(quark Quark) string {
	c_quark := (C.GQuark)(quark)

	retC := C.g_quark_to_string(c_quark)
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : g_quark_try_string

// Blacklisted : g_random_double

// Blacklisted : g_random_double_range

// Blacklisted : g_random_int

// Blacklisted : g_random_int_range

// Blacklisted : g_random_set_seed

// Unsupported : g_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Blacklisted : g_regex_error_quark

// Blacklisted : g_return_if_fail_warning

// Blacklisted : g_set_error

// Blacklisted : g_set_prgname

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Blacklisted : g_shell_error_quark

// Unsupported : g_shell_parse_argv : unsupported parameter argcp : array length param argcp is pointer (gint*)

// Blacklisted : g_shell_quote

// Blacklisted : g_shell_unquote

// Blacklisted : g_slice_get_config

// Blacklisted : g_slice_get_config_state

// Blacklisted : g_slice_set_config

// Blacklisted : g_snprintf

// Blacklisted : g_source_remove

// Unsupported : g_source_remove_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_source_remove_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Blacklisted : g_spaced_primes_closest

// Unsupported : g_spawn_async : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_spawn_async_with_pipes : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Blacklisted : g_spawn_close_pid

// Blacklisted : g_spawn_command_line_async

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : output array param standard_output

// Blacklisted : g_spawn_error_quark

// Blacklisted : g_spawn_exit_error_quark

// Unsupported : g_spawn_sync : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Blacklisted : g_stpcpy

// Unsupported : g_str_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_str_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Blacklisted : g_strcanon

// Blacklisted : g_strcasecmp

// Blacklisted : g_strchomp

// Blacklisted : g_strchug

// Blacklisted : g_strcompress

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Blacklisted : g_strdelimit

// Blacklisted : g_strdown

// Blacklisted : g_strdup

// Blacklisted : g_strdup_printf

// Unsupported : g_strdup_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_strdupv : unsupported parameter str_array : in string with indirection level of 2

// Blacklisted : g_strerror

// Blacklisted : g_strescape

// Unsupported : g_strfreev : unsupported parameter str_array : in string with indirection level of 2

// Blacklisted : g_string_new

// Blacklisted : g_string_new_len

// Blacklisted : g_string_sized_new

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Blacklisted : g_strlcat

// Blacklisted : g_strlcpy

// Blacklisted : g_strncasecmp

// Blacklisted : g_strndup

// Blacklisted : g_strnfill

// Strreverse is a wrapper around the C function g_strreverse.
func Strreverse(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strreverse(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : g_strrstr

// Blacklisted : g_strrstr_len

// Blacklisted : g_strsignal

// Blacklisted : g_strsplit

// Blacklisted : g_strstr_len

// Blacklisted : g_strtod

// Blacklisted : g_strup

// Blacklisted : g_strv_get_type

// Blacklisted : g_test_assert_expected_messages_internal

// Blacklisted : g_test_log_type_name

// Blacklisted : g_test_trap_assertions

// Blacklisted : g_thread_error_quark

// Unsupported : g_thread_exit : unsupported parameter retval : no type generator for gpointer (gpointer) for param retval

// Blacklisted : g_thread_pool_get_max_unused_threads

// Blacklisted : g_thread_pool_get_num_unused_threads

// Blacklisted : g_thread_pool_set_max_unused_threads

// Blacklisted : g_thread_pool_stop_unused_threads

// Blacklisted : g_thread_self

// Blacklisted : g_thread_yield

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Blacklisted : g_timeout_source_new

// Blacklisted : g_trash_stack_height

// Unsupported : g_trash_stack_peek : no return generator

// Unsupported : g_trash_stack_pop : no return generator

// Unsupported : g_trash_stack_push : unsupported parameter data_p : no type generator for gpointer (gpointer) for param data_p

// Unsupported : g_try_malloc : no return generator

// Unsupported : g_try_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Blacklisted : g_ucs4_to_utf16

// Blacklisted : g_ucs4_to_utf8

// Blacklisted : g_unichar_break_type

// Blacklisted : g_unichar_digit_value

// Blacklisted : g_unichar_isalnum

// Blacklisted : g_unichar_isalpha

// Blacklisted : g_unichar_iscntrl

// Blacklisted : g_unichar_isdefined

// Blacklisted : g_unichar_isdigit

// Blacklisted : g_unichar_isgraph

// Blacklisted : g_unichar_islower

// Blacklisted : g_unichar_isprint

// Blacklisted : g_unichar_ispunct

// Blacklisted : g_unichar_isspace

// Blacklisted : g_unichar_istitle

// Blacklisted : g_unichar_isupper

// Blacklisted : g_unichar_iswide

// Blacklisted : g_unichar_isxdigit

// Blacklisted : g_unichar_to_utf8

// Blacklisted : g_unichar_tolower

// Blacklisted : g_unichar_totitle

// Blacklisted : g_unichar_toupper

// Blacklisted : g_unichar_type

// Blacklisted : g_unichar_validate

// Blacklisted : g_unichar_xdigit_value

// Blacklisted : g_unicode_canonical_decomposition

// Blacklisted : g_unicode_canonical_ordering

// Blacklisted : g_unix_error_quark

// Blacklisted : g_usleep

// Blacklisted : g_utf16_to_ucs4

// Blacklisted : g_utf16_to_utf8

// Blacklisted : g_utf8_casefold

// Blacklisted : g_utf8_collate

// Blacklisted : g_utf8_collate_key

// Blacklisted : g_utf8_find_next_char

// Blacklisted : g_utf8_find_prev_char

// Blacklisted : g_utf8_get_char

// Blacklisted : g_utf8_get_char_validated

// Blacklisted : g_utf8_normalize

// Blacklisted : g_utf8_offset_to_pointer

// Blacklisted : g_utf8_pointer_to_offset

// Blacklisted : g_utf8_prev_char

// Blacklisted : g_utf8_strchr

// Blacklisted : g_utf8_strdown

// Blacklisted : g_utf8_strlen

// Blacklisted : g_utf8_strncpy

// Blacklisted : g_utf8_strrchr

// Blacklisted : g_utf8_strup

// Blacklisted : g_utf8_to_ucs4

// Blacklisted : g_utf8_to_ucs4_fast

// Blacklisted : g_utf8_to_utf16

// Blacklisted : g_utf8_validate

// Blacklisted : g_variant_get_gtype

// Unsupported : g_variant_parse : unsupported parameter endptr : in string with indirection level of 2

// Blacklisted : g_variant_parse_error_quark

// Blacklisted : g_variant_parser_get_error_quark

// Blacklisted : g_variant_type_checked_

// Blacklisted : g_variant_type_string_is_valid

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : g_warn_message

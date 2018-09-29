// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_ascii_string_to_signed : unsupported parameter out_num : no type generator for gint64, gint64*

// Unsupported : g_ascii_string_to_unsigned : unsupported parameter out_num : no type generator for guint64, guint64*

// Unsupported : g_assert_warning : unsupported parameter line : no type generator for gint, const int

// Unsupported : g_assertion_message_cmpnum : unsupported parameter arg1 : no type generator for long double, long double

// Unsupported : g_atexit : unsupported parameter func : no type generator for VoidFunc, GVoidFunc

// Unsupported : g_atomic_int_add : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_and : unsupported parameter atomic : no type generator for guint, volatile guint*

// Unsupported : g_atomic_int_compare_and_exchange : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_dec_and_test : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_exchange_and_add : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_get : unsupported parameter atomic : no type generator for gint, volatile const gint*

// Unsupported : g_atomic_int_inc : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_or : unsupported parameter atomic : no type generator for guint, volatile guint*

// Unsupported : g_atomic_int_set : unsupported parameter atomic : no type generator for gint, volatile gint*

// Unsupported : g_atomic_int_xor : unsupported parameter atomic : no type generator for guint, volatile guint*

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_base64_decode : unsupported parameter out_len : no type generator for gsize, gsize*

// Unsupported : g_base64_decode_inplace : unsupported parameter text : no param type

// Unsupported : g_base64_decode_step : unsupported parameter in : no param type

// Unsupported : g_base64_encode : unsupported parameter data : no param type

// Unsupported : g_base64_encode_close : unsupported parameter out : no param type

// Unsupported : g_base64_encode_step : unsupported parameter in : no param type

// Unsupported : g_bit_lock : unsupported parameter address : no type generator for gint, volatile gint*

// Unsupported : g_bit_trylock : unsupported parameter address : no type generator for gint, volatile gint*

// Unsupported : g_bit_unlock : unsupported parameter address : no type generator for gint, volatile gint*

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_filename_valist : unsupported parameter args : no type generator for va_list, va_list*

// Unsupported : g_build_filenamev : unsupported parameter args : no param type

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Unsupported : g_build_pathv : unsupported parameter args : no param type

// Unsupported : g_byte_array_free : unsupported parameter array : no param type

// Unsupported : g_byte_array_free_to_bytes : unsupported parameter array : no param type

// Unsupported : g_byte_array_new : no return type

// Unsupported : g_byte_array_new_take : unsupported parameter data : no param type

// Unsupported : g_byte_array_unref : unsupported parameter array : no param type

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc, GChildWatchFunc

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc, GChildWatchFunc

// Unsupported : g_clear_handle_id : unsupported parameter tag_ptr : no type generator for guint, guint*

// Unsupported : g_clear_pointer : unsupported parameter pp : no type generator for gpointer, gpointer*

// Unsupported : g_compute_checksum_for_data : unsupported parameter data : no param type

// ComputeHmacForBytes is a wrapper around the C function g_compute_hmac_for_bytes.
func ComputeHmacForBytes(digestType ChecksumType, key *Bytes, data *Bytes) string {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key := (*C.GBytes)(key.ToC())

	c_data := (*C.GBytes)(data.ToC())

	retC := C.g_compute_hmac_for_bytes(c_digest_type, c_key, c_data)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_compute_hmac_for_data : unsupported parameter key : no param type

// Unsupported : g_compute_hmac_for_string : unsupported parameter key : no param type

// Unsupported : g_convert : unsupported parameter str : no param type

// Unsupported : g_convert_with_fallback : unsupported parameter str : no param type

// Unsupported : g_convert_with_iconv : unsupported parameter str : no param type

// Unsupported : g_datalist_clear : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_foreach : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_get_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_get_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_dup_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_get_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_remove_no_notify : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_replace_data : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_id_set_data_full : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_init : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_set_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_datalist_unset_flags : unsupported parameter datalist : record with indirection level of 2

// Unsupported : g_dataset_foreach : unsupported parameter func : no type generator for DataForeachFunc, GDataForeachFunc

// Unsupported : g_dataset_id_set_data_full : unsupported parameter destroy_func : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_environ_getenv : unsupported parameter envp : no param type

// Unsupported : g_environ_setenv : unsupported parameter envp : no param type

// Unsupported : g_environ_unsetenv : unsupported parameter envp : no param type

// Unsupported : g_file_get_contents : unsupported parameter contents : no param type

// Unsupported : g_file_set_contents : unsupported parameter contents : no param type

// Unsupported : g_filename_from_utf8 : unsupported parameter bytes_read : no type generator for gsize, gsize*

// Unsupported : g_filename_to_utf8 : unsupported parameter bytes_read : no type generator for gsize, gsize*

// Unsupported : g_fprintf : unsupported parameter file : no type generator for gpointer, FILE*

// Unsupported : g_get_environ : no return type

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// Unsupported : g_get_language_names : no return type

// Unsupported : g_get_locale_variants : no return type

// Unsupported : g_get_system_config_dirs : no return type

// Unsupported : g_get_system_data_dirs : no return type

// Unsupported : g_hash_table_lookup_extended : unsupported parameter orig_key : no type generator for gpointer, gpointer*

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_io_add_watch : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_io_add_watch_full : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_io_create_watch : unsupported parameter channel : Blacklisted record : GIOChannel

// Unsupported : g_listenv : no return type

// Unsupported : g_locale_from_utf8 : unsupported parameter bytes_read : no type generator for gsize, gsize*

// Unsupported : g_locale_to_utf8 : unsupported parameter opsysstring : no param type

// Unsupported : g_log : unsupported parameter ... : varargs

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no type generator for LogFunc, GLogFunc

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc, GLogFunc

// Unsupported : g_log_set_handler_full : unsupported parameter log_func : no type generator for LogFunc, GLogFunc

// Unsupported : g_log_set_writer_func : unsupported parameter func : no type generator for LogWriterFunc, GLogWriterFunc

// Unsupported : g_log_structured : unsupported parameter ... : varargs

// Unsupported : g_log_structured_array : unsupported parameter fields : no param type

// Unsupported : g_log_structured_standard : unsupported parameter ... : varargs

// Unsupported : g_log_variant : unsupported parameter fields : Blacklisted record : GVariant

// Unsupported : g_log_writer_default : unsupported parameter fields : no param type

// Unsupported : g_log_writer_format_fields : unsupported parameter fields : no param type

// LogWriterIsJournald is a wrapper around the C function g_log_writer_is_journald.
func LogWriterIsJournald(outputFd int32) bool {
	c_output_fd := (C.gint)(outputFd)

	retC := C.g_log_writer_is_journald(c_output_fd)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_log_writer_journald : unsupported parameter fields : no param type

// Unsupported : g_log_writer_standard_streams : unsupported parameter fields : no param type

// LogWriterSupportsColor is a wrapper around the C function g_log_writer_supports_color.
func LogWriterSupportsColor(outputFd int32) bool {
	c_output_fd := (C.gint)(outputFd)

	retC := C.g_log_writer_supports_color(c_output_fd)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : no type generator for gpointer, gpointer*

// Blacklisted : g_number_parser_error_quark

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer, void*

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer, void*

// Unsupported : g_parse_debug_string : unsupported parameter keys : no param type

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer, void*

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer, void*

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer, void*

// Unsupported : g_prefix_error : unsupported parameter err : record with indirection level of 2

// Unsupported : g_print : unsupported parameter ... : varargs

// Unsupported : g_printerr : unsupported parameter ... : varargs

// Unsupported : g_printf : unsupported parameter ... : varargs

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_propagate_error : unsupported parameter dest : record with indirection level of 2

// Unsupported : g_propagate_prefixed_error : unsupported parameter dest : record with indirection level of 2

// Unsupported : g_ptr_array_find : unsupported parameter haystack : no param type

// Unsupported : g_ptr_array_find_with_equal_func : unsupported parameter haystack : no param type

// Unsupported : g_qsort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_regex_escape_string : unsupported parameter string : no param type

// Unsupported : g_regex_split_simple : no return type

// Unsupported : g_set_error : unsupported parameter err : record with indirection level of 2

// Unsupported : g_set_error_literal : unsupported parameter err : record with indirection level of 2

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc, GPrintFunc

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc, GPrintFunc

// Unsupported : g_shell_parse_argv : unsupported parameter argcp : no type generator for gint, gint*

// Unsupported : g_slice_get_config_state : unsupported parameter n_values : no type generator for guint, guint*

// Unsupported : g_snprintf : unsupported parameter ... : varargs

// Unsupported : g_spawn_async : unsupported parameter argv : no param type

// Unsupported : g_spawn_async_with_pipes : unsupported parameter argv : no param type

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : no param type

// Unsupported : g_spawn_sync : unsupported parameter argv : no param type

// Unsupported : g_sprintf : unsupported parameter ... : varargs

// Unsupported : g_str_tokenize_and_fold : unsupported parameter ascii_alternates : no param type

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Unsupported : g_strdup_printf : unsupported parameter ... : varargs

// Unsupported : g_strdup_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_strdupv : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strfreev : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strsplit : no return type

// Unsupported : g_strsplit_set : no return type

// Unsupported : g_strv_get_type : no return generator

// Unsupported : g_strv_length : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_test_add_data_func : unsupported parameter test_func : no type generator for TestDataFunc, GTestDataFunc

// Unsupported : g_test_add_data_func_full : unsupported parameter test_func : no type generator for TestDataFunc, GTestDataFunc

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc, GTestFunc

// Unsupported : g_test_add_vtable : unsupported parameter data_setup : no type generator for TestFixtureFunc, GTestFixtureFunc

// Unsupported : g_test_build_filename : unsupported parameter ... : varargs

// Unsupported : g_test_create_case : unsupported parameter data_setup : no type generator for TestFixtureFunc, GTestFixtureFunc

// Unsupported : g_test_get_filename : unsupported parameter ... : varargs

// Unsupported : g_test_init : unsupported parameter argc : no type generator for gint, int*

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc, GTestLogFatalFunc

// Unsupported : g_test_maximized_result : unsupported parameter ... : varargs

// Unsupported : g_test_message : unsupported parameter ... : varargs

// Unsupported : g_test_minimized_result : unsupported parameter ... : varargs

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_trash_stack_height : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_peek : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_pop : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_trash_stack_push : unsupported parameter stack_p : record with indirection level of 2

// Unsupported : g_ucs4_to_utf16 : unsupported parameter str : no type generator for gunichar, const gunichar*

// Unsupported : g_ucs4_to_utf8 : unsupported parameter str : no type generator for gunichar, const gunichar*

// Unsupported : g_unichar_compose : unsupported parameter ch : no type generator for gunichar, gunichar*

// Unsupported : g_unichar_decompose : unsupported parameter a : no type generator for gunichar, gunichar*

// Unsupported : g_unichar_fully_decompose : unsupported parameter result : no type generator for gunichar, gunichar*

// Unsupported : g_unichar_get_mirror_char : unsupported parameter mirrored_ch : no type generator for gunichar, gunichar*

// Blacklisted : g_unichar_to_utf8

// Unsupported : g_unicode_canonical_decomposition : unsupported parameter result_len : no type generator for gsize, gsize*

// Unsupported : g_unicode_canonical_ordering : unsupported parameter string : no type generator for gunichar, gunichar*

// Blacklisted : g_unix_error_quark

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc, GUnixFDSourceFunc

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc, GUnixFDSourceFunc

// Unsupported : g_unix_open_pipe : unsupported parameter fds : no type generator for gint, gint*

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_uri_list_extract_uris : no return type

// Unsupported : g_utf16_to_ucs4 : unsupported parameter str : no type generator for guint16, const gunichar2*

// Unsupported : g_utf16_to_utf8 : unsupported parameter str : no type generator for guint16, const gunichar2*

// Unsupported : g_utf8_to_ucs4 : unsupported parameter items_read : no type generator for glong, glong*

// Unsupported : g_utf8_to_ucs4_fast : unsupported parameter items_written : no type generator for glong, glong*

// Unsupported : g_utf8_to_utf16 : unsupported parameter items_read : no type generator for glong, glong*

// Unsupported : g_utf8_validate : unsupported parameter str : no param type

// Unsupported : g_variant_get_gtype : no return generator

// Unsupported : g_variant_parse : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_type_checked_ : return type : Blacklisted record : GVariantType

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2

// Unsupported : g_vfprintf : unsupported parameter file : no type generator for gpointer, FILE*

// Unsupported : g_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_vsprintf : unsupported parameter args : no type generator for va_list, va_list

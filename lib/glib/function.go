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

// AsciiDtostr is a wrapper around the C function g_ascii_dtostr.
func AsciiDtostr(buffer string, bufLen int32, d float64) string {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_buf_len := (C.gint)(bufLen)

	c_d := (C.gdouble)(d)

	retC := C.g_ascii_dtostr(c_buffer, c_buf_len, c_d)
	retGo :=
		C.GoString(retC)

	return retGo
}

// AsciiFormatd is a wrapper around the C function g_ascii_formatd.
func AsciiFormatd(buffer string, bufLen int32, format string, d float64) string {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_buf_len := (C.gint)(bufLen)

	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	c_d := (C.gdouble)(d)

	retC := C.g_ascii_formatd(c_buffer, c_buf_len, c_format, c_d)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// AsciiStrdown is a wrapper around the C function g_ascii_strdown.
func AsciiStrdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strdown(c_str, c_len)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// AsciiStrup is a wrapper around the C function g_ascii_strup.
func AsciiStrup(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strup(c_str, c_len)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Unsupported : g_assert_warning : unsupported parameter line : no param type for gint, const int

// Unsupported : g_assertion_message : unsupported parameter line : no param type for gint, int

// Unsupported : g_assertion_message_cmpnum : unsupported parameter line : no param type for gint, int

// Unsupported : g_assertion_message_cmpstr : unsupported parameter line : no param type for gint, int

// Unsupported : g_assertion_message_error : unsupported parameter line : no param type for gint, int

// Unsupported : g_assertion_message_expr : unsupported parameter line : no param type for gint, int

// Unsupported : g_atexit : unsupported parameter func : no param type for VoidFunc, GVoidFunc

// Unsupported : g_basename : no return type

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

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Unsupported : g_byte_array_free : unsupported parameter array : no param type

// Unsupported : g_byte_array_new : no return type

// Unsupported : g_clear_error : no return type

// Unsupported : g_convert : unsupported parameter str : no param type

// Unsupported : g_convert_error_quark : no return type

// Unsupported : g_convert_with_fallback : unsupported parameter str : no param type

// Unsupported : g_convert_with_iconv : unsupported parameter str : no param type

// Unsupported : g_datalist_clear : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_foreach : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_get_data : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_id_get_data : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_id_remove_no_notify : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_id_set_data_full : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_datalist_init : unsupported parameter datalist : no param type for Data, GData**

// Unsupported : g_dataset_destroy : unsupported parameter dataset_location : no param type for gpointer, gconstpointer

// Unsupported : g_dataset_foreach : unsupported parameter dataset_location : no param type for gpointer, gconstpointer

// Unsupported : g_dataset_id_get_data : unsupported parameter dataset_location : no param type for gpointer, gconstpointer

// Unsupported : g_dataset_id_remove_no_notify : unsupported parameter dataset_location : no param type for gpointer, gconstpointer

// Unsupported : g_dataset_id_set_data_full : unsupported parameter dataset_location : no param type for gpointer, gconstpointer

// Unsupported : g_date_get_days_in_month : unsupported parameter month : no param type for DateMonth, GDateMonth

// Unsupported : g_date_get_monday_weeks_in_year : unsupported parameter year : no param type for DateYear, GDateYear

// Unsupported : g_date_get_sunday_weeks_in_year : unsupported parameter year : no param type for DateYear, GDateYear

// Unsupported : g_date_is_leap_year : unsupported parameter year : no param type for DateYear, GDateYear

// Unsupported : g_date_strftime : unsupported parameter date : no param type for Date, const GDate*

// Unsupported : g_date_valid_day : unsupported parameter day : no param type for DateDay, GDateDay

// Unsupported : g_date_valid_dmy : unsupported parameter day : no param type for DateDay, GDateDay

// Unsupported : g_date_valid_julian : no return type

// Unsupported : g_date_valid_month : unsupported parameter month : no param type for DateMonth, GDateMonth

// Unsupported : g_date_valid_weekday : unsupported parameter weekday : no param type for DateWeekday, GDateWeekday

// Unsupported : g_date_valid_year : unsupported parameter year : no param type for DateYear, GDateYear

// Unsupported : g_direct_equal : unsupported parameter v1 : no param type for gpointer, gconstpointer

// Unsupported : g_direct_hash : unsupported parameter v : no param type for gpointer, gconstpointer

// Unsupported : g_file_error_from_errno : no return type

// Unsupported : g_file_error_quark : no return type

// Unsupported : g_file_get_contents : unsupported parameter contents : no param type

// Unsupported : throws

// Unsupported : g_file_test : unsupported parameter test : no param type for FileTest, GFileTest

// Unsupported : g_filename_from_uri : no return type

// Unsupported : g_filename_from_utf8 : unsupported parameter bytes_read : no param type for gsize, gsize*

// Unsupported : throws

// Unsupported : g_filename_to_utf8 : unsupported parameter bytes_read : no param type for gsize, gsize*

// Unsupported : g_find_program_in_path : no return type

// Unsupported : g_free : no return type

// Unsupported : g_get_charset : no return type

// GetCodeset is a wrapper around the C function g_get_codeset.
func GetCodeset() string {
	retC := C.g_get_codeset()
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_get_current_dir : no return type

// Unsupported : g_get_current_time : unsupported parameter result : no param type for TimeVal, GTimeVal*

// Unsupported : g_get_home_dir : no return type

// GetPrgname is a wrapper around the C function g_get_prgname.
func GetPrgname() string {
	retC := C.g_get_prgname()
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_get_real_name : no return type

// Unsupported : g_get_tmp_dir : no return type

// Unsupported : g_get_user_name : no return type

// Unsupported : g_getenv : no return type

// Unsupported : g_hash_table_destroy : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_insert : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_lookup : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_lookup_extended : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_remove : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_replace : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_size : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_steal : unsupported parameter hash_table : no param type for GLib.HashTable, GHashTable*

// Unsupported : g_hook_destroy : unsupported parameter hook_list : no param type for HookList, GHookList*

// Unsupported : g_hook_destroy_link : unsupported parameter hook_list : no param type for HookList, GHookList*

// Unsupported : g_hook_free : unsupported parameter hook_list : no param type for HookList, GHookList*

// Unsupported : g_hook_insert_before : unsupported parameter hook_list : no param type for HookList, GHookList*

// Unsupported : g_hook_prepend : unsupported parameter hook_list : no param type for HookList, GHookList*

// Unsupported : g_hook_unref : unsupported parameter hook_list : no param type for HookList, GHookList*

// Unsupported : g_iconv : unsupported parameter converter : no param type for IConv, GIConv

// Unsupported : g_iconv_open : no return type

// Unsupported : g_idle_add : unsupported parameter function : no param type for SourceFunc, GSourceFunc

// Unsupported : g_idle_add_full : unsupported parameter function : no param type for SourceFunc, GSourceFunc

// Unsupported : g_idle_remove_by_data : no return type

// Unsupported : g_idle_source_new : no return type

// Unsupported : g_int_equal : unsupported parameter v1 : no param type for gpointer, gconstpointer

// Unsupported : g_int_hash : unsupported parameter v : no param type for gpointer, gconstpointer

// Unsupported : g_io_add_watch : unsupported parameter channel : no param type for IOChannel, GIOChannel*

// Unsupported : g_io_add_watch_full : unsupported parameter channel : no param type for IOChannel, GIOChannel*

// Unsupported : g_io_channel_error_from_errno : no return type

// Unsupported : g_io_channel_error_quark : no return type

// Unsupported : g_io_create_watch : unsupported parameter channel : no param type for IOChannel, GIOChannel*

// Unsupported : g_key_file_error_quark : no return type

// Unsupported : g_locale_from_utf8 : unsupported parameter bytes_read : no param type for gsize, gsize*

// Unsupported : g_locale_to_utf8 : unsupported parameter opsysstring : no param type

// Unsupported : g_log : unsupported parameter log_level : no param type for LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_default_handler : unsupported parameter log_level : no param type for LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_remove_handler : no return type

// Unsupported : g_log_set_always_fatal : unsupported parameter fatal_mask : no param type for LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_set_fatal_mask : unsupported parameter fatal_mask : no param type for LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_set_handler : unsupported parameter log_levels : no param type for LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_structured_standard : unsupported parameter log_level : no param type for LogLevelFlags, GLogLevelFlags

// Unsupported : g_logv : unsupported parameter log_level : no param type for LogLevelFlags, GLogLevelFlags

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

// MarkupEscapeText is a wrapper around the C function g_markup_escape_text.
func MarkupEscapeText(text string, length int64) string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gssize)(length)

	retC := C.g_markup_escape_text(c_text, c_length)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_mem_is_system_malloc : no return type

// Unsupported : g_mem_profile : no return type

// Unsupported : g_mem_set_vtable : unsupported parameter vtable : no param type for MemVTable, GMemVTable*

// Unsupported : g_memdup : unsupported parameter mem : no param type for gpointer, gconstpointer

// Mkstemp is a wrapper around the C function g_mkstemp.
func Mkstemp(tmpl string) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkstemp(c_tmpl)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : no param type for gpointer, gpointer*

// Unsupported : g_number_parser_error_quark : no return type

// Unsupported : g_on_error_query : no return type

// Unsupported : g_on_error_stack_trace : no return type

// Unsupported : g_option_error_quark : no return type

// Unsupported : g_parse_debug_string : unsupported parameter keys : no param type

// Unsupported : g_path_get_basename : no return type

// Unsupported : g_path_get_dirname : no return type

// Unsupported : g_path_is_absolute : no return type

// Unsupported : g_path_skip_root : no return type

// Unsupported : g_pattern_match : unsupported parameter pspec : no param type for PatternSpec, GPatternSpec*

// Unsupported : g_pattern_match_simple : no return type

// Unsupported : g_pattern_match_string : unsupported parameter pspec : no param type for PatternSpec, GPatternSpec*

// Unsupported : g_print : unsupported parameter ... : varargs

// Unsupported : g_printerr : unsupported parameter ... : varargs

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no param type for va_list, va_list

// Unsupported : g_propagate_error : unsupported parameter dest : no param type for Error, GError**

// Unsupported : g_qsort_with_data : unsupported parameter pbase : no param type for gpointer, gconstpointer

// Unsupported : g_quark_from_static_string : no return type

// Unsupported : g_quark_from_string : no return type

// Unsupported : g_quark_to_string : unsupported parameter quark : no param type for Quark, GQuark

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

// Unsupported : g_set_error : unsupported parameter err : no param type for Error, GError**

// Unsupported : g_set_prgname : no return type

// Unsupported : g_set_print_handler : unsupported parameter func : no param type for PrintFunc, GPrintFunc

// Unsupported : g_set_printerr_handler : unsupported parameter func : no param type for PrintFunc, GPrintFunc

// Unsupported : g_shell_error_quark : no return type

// Unsupported : g_shell_parse_argv : unsupported parameter argcp : no param type for gint, gint*

// Unsupported : g_shell_quote : no return type

// Unsupported : g_shell_unquote : no return type

// Unsupported : g_slice_get_config : unsupported parameter ckey : no param type for SliceConfig, GSliceConfig

// Unsupported : g_slice_get_config_state : unsupported parameter ckey : no param type for SliceConfig, GSliceConfig

// Unsupported : g_slice_set_config : unsupported parameter ckey : no param type for SliceConfig, GSliceConfig

// Unsupported : g_snprintf : unsupported parameter ... : varargs

// Unsupported : g_source_remove : no return type

// Unsupported : g_source_remove_by_funcs_user_data : unsupported parameter funcs : no param type for SourceFuncs, GSourceFuncs*

// Unsupported : g_source_remove_by_user_data : no return type

// SpacedPrimesClosest is a wrapper around the C function g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	c_num := (C.guint)(num)

	retC := C.g_spaced_primes_closest(c_num)
	retGo :=
		(uint32)(retC)

	return retGo
}

// Unsupported : g_spawn_async : unsupported parameter argv : no param type

// Unsupported : g_spawn_async_with_pipes : unsupported parameter argv : no param type

// Unsupported : g_spawn_close_pid : unsupported parameter pid : no param type for Pid, GPid

// Unsupported : g_spawn_command_line_async : no return type

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : no param type

// Unsupported : g_spawn_error_quark : no return type

// Unsupported : g_spawn_exit_error_quark : no return type

// Unsupported : g_spawn_sync : unsupported parameter argv : no param type

// Stpcpy is a wrapper around the C function g_stpcpy.
func Stpcpy(dest string, src string) string {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	retC := C.g_stpcpy(c_dest, c_src)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_str_equal : unsupported parameter v1 : no param type for gpointer, gconstpointer

// Unsupported : g_str_hash : unsupported parameter v : no param type for gpointer, gconstpointer

// Strcanon is a wrapper around the C function g_strcanon.
func Strcanon(string string, validChars string, substitutor rune) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_valid_chars := C.CString(validChars)
	defer C.free(unsafe.Pointer(c_valid_chars))

	c_substitutor := (C.gchar)(substitutor)

	retC := C.g_strcanon(c_string, c_valid_chars, c_substitutor)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Strchomp is a wrapper around the C function g_strchomp.
func Strchomp(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchomp(c_string)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strchug is a wrapper around the C function g_strchug.
func Strchug(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchug(c_string)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strcompress is a wrapper around the C function g_strcompress.
func Strcompress(source string) string {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	retC := C.g_strcompress(c_source)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Strdelimit is a wrapper around the C function g_strdelimit.
func Strdelimit(string string, delimiters string, newDelimiter rune) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiters := C.CString(delimiters)
	defer C.free(unsafe.Pointer(c_delimiters))

	c_new_delimiter := (C.gchar)(newDelimiter)

	retC := C.g_strdelimit(c_string, c_delimiters, c_new_delimiter)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strdown is a wrapper around the C function g_strdown.
func Strdown(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strdown(c_string)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strdup is a wrapper around the C function g_strdup.
func Strdup(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.g_strdup(c_str)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_strdup_printf : unsupported parameter ... : varargs

// Unsupported : g_strdup_vprintf : unsupported parameter args : no param type for va_list, va_list

// Unsupported : g_strdupv : unsupported parameter str_array : in for string with indirection level of 2

// Strerror is a wrapper around the C function g_strerror.
func Strerror(errnum int32) string {
	c_errnum := (C.gint)(errnum)

	retC := C.g_strerror(c_errnum)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strescape is a wrapper around the C function g_strescape.
func Strescape(source string, exceptions string) string {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	c_exceptions := C.CString(exceptions)
	defer C.free(unsafe.Pointer(c_exceptions))

	retC := C.g_strescape(c_source, c_exceptions)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Strndup is a wrapper around the C function g_strndup.
func Strndup(str string, n uint64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_n := (C.gsize)(n)

	retC := C.g_strndup(c_str, c_n)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strnfill is a wrapper around the C function g_strnfill.
func Strnfill(length uint64, fillChar rune) string {
	c_length := (C.gsize)(length)

	c_fill_char := (C.gchar)(fillChar)

	retC := C.g_strnfill(c_length, c_fill_char)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strreverse is a wrapper around the C function g_strreverse.
func Strreverse(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strreverse(c_string)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strrstr is a wrapper around the C function g_strrstr.
func Strrstr(haystack string, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strrstr(c_haystack, c_needle)
	retGo :=
		C.GoString(retC)

	return retGo
}

// StrrstrLen is a wrapper around the C function g_strrstr_len.
func StrrstrLen(haystack string, haystackLen int64, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_haystack_len := (C.gssize)(haystackLen)

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strrstr_len(c_haystack, c_haystack_len, c_needle)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Strsignal is a wrapper around the C function g_strsignal.
func Strsignal(signum int32) string {
	c_signum := (C.gint)(signum)

	retC := C.g_strsignal(c_signum)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strsplit : no return type

// StrstrLen is a wrapper around the C function g_strstr_len.
func StrstrLen(haystack string, haystackLen int64, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_haystack_len := (C.gssize)(haystackLen)

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strstr_len(c_haystack, c_haystack_len, c_needle)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Strup is a wrapper around the C function g_strup.
func Strup(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strup(c_string)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_strv_get_type : no return type

// Unsupported : g_test_add_vtable : unsupported parameter test_data : no param type for gpointer, gconstpointer

// Unsupported : g_test_assert_expected_messages_internal : unsupported parameter line : no param type for gint, int

// Unsupported : g_test_log_type_name : unsupported parameter log_type : no param type for TestLogType, GTestLogType

// Unsupported : g_test_trap_assertions : unsupported parameter line : no param type for gint, int

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

// Unsupported : g_timeout_add : unsupported parameter function : no param type for SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_full : unsupported parameter function : no param type for SourceFunc, GSourceFunc

// Unsupported : g_timeout_source_new : no return type

// Unsupported : g_trash_stack_height : unsupported parameter stack_p : no param type for TrashStack, GTrashStack**

// Unsupported : g_trash_stack_peek : unsupported parameter stack_p : no param type for TrashStack, GTrashStack**

// Unsupported : g_trash_stack_pop : unsupported parameter stack_p : no param type for TrashStack, GTrashStack**

// Unsupported : g_trash_stack_push : unsupported parameter stack_p : no param type for TrashStack, GTrashStack**

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

// Unsupported : g_ucs4_to_utf16 : unsupported parameter str : no param type for gunichar, const gunichar*

// Unsupported : g_ucs4_to_utf8 : unsupported parameter str : no param type for gunichar, const gunichar*

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

// Unsupported : g_unicode_canonical_decomposition : unsupported parameter result_len : no param type for gsize, gsize*

// Unsupported : g_unicode_canonical_ordering : unsupported parameter string : no param type for gunichar, gunichar*

// Unsupported : g_unix_error_quark : no return type

// Unsupported : g_usleep : no return type

// Unsupported : g_utf16_to_ucs4 : unsupported parameter str : no param type for guint16, const gunichar2*

// Unsupported : g_utf16_to_utf8 : unsupported parameter str : no param type for guint16, const gunichar2*

// Utf8Casefold is a wrapper around the C function g_utf8_casefold.
func Utf8Casefold(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_casefold(c_str, c_len)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Utf8CollateKey is a wrapper around the C function g_utf8_collate_key.
func Utf8CollateKey(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_collate_key(c_str, c_len)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Utf8FindNextChar is a wrapper around the C function g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_end := C.CString(end)
	defer C.free(unsafe.Pointer(c_end))

	retC := C.g_utf8_find_next_char(c_p, c_end)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Utf8FindPrevChar is a wrapper around the C function g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_find_prev_char(c_str, c_p)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Unsupported : g_utf8_normalize : unsupported parameter mode : no param type for NormalizeMode, GNormalizeMode

// Utf8OffsetToPointer is a wrapper around the C function g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_offset := (C.glong)(offset)

	retC := C.g_utf8_offset_to_pointer(c_str, c_offset)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Utf8PrevChar is a wrapper around the C function g_utf8_prev_char.
func Utf8PrevChar(p string) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_prev_char(c_p)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Utf8Strchr is a wrapper around the C function g_utf8_strchr.
func Utf8Strchr(p string, len int64, c rune) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	retC := C.g_utf8_strchr(c_p, c_len, c_c)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Utf8Strdown is a wrapper around the C function g_utf8_strdown.
func Utf8Strdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strdown(c_str, c_len)
	retGo :=
		C.GoString(retC)

	return retGo
}

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

// Utf8Strncpy is a wrapper around the C function g_utf8_strncpy.
func Utf8Strncpy(dest string, src string, n uint64) string {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_n := (C.gsize)(n)

	retC := C.g_utf8_strncpy(c_dest, c_src, c_n)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Utf8Strrchr is a wrapper around the C function g_utf8_strrchr.
func Utf8Strrchr(p string, len int64, c rune) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	retC := C.g_utf8_strrchr(c_p, c_len, c_c)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Utf8Strup is a wrapper around the C function g_utf8_strup.
func Utf8Strup(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strup(c_str, c_len)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_utf8_to_ucs4 : unsupported parameter items_read : no param type for glong, glong*

// Unsupported : g_utf8_to_ucs4_fast : unsupported parameter items_written : no param type for glong, glong*

// Unsupported : g_utf8_to_utf16 : unsupported parameter items_read : no param type for glong, glong*

// Unsupported : g_utf8_validate : unsupported parameter str : no param type

// Unsupported : g_variant_get_gtype : no return type

// Unsupported : g_variant_parse : unsupported parameter type : no param type for VariantType, const GVariantType*

// Unsupported : g_variant_parse_error_quark : no return type

// Unsupported : g_variant_parser_get_error_quark : no return type

// Unsupported : g_variant_type_checked_ : no return type

// Unsupported : g_variant_type_string_is_valid : no return type

// Unsupported : g_vsnprintf : unsupported parameter args : no param type for va_list, va_list

// Unsupported : g_warn_message : unsupported parameter line : no param type for gint, int

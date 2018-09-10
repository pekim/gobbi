package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// AsciiDigitValue is a wrapper around the C function g_ascii_digit_value.
func AsciiDigitValue(c rune) {
	c_c := (C.gchar)(c)

	C.g_ascii_digit_value(c_c)
}

// AsciiDtostr is a wrapper around the C function g_ascii_dtostr.
func AsciiDtostr(buffer string, bufLen int32, d float64) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_buf_len := (C.gint)(bufLen)

	c_d := (C.gdouble)(d)

	C.g_ascii_dtostr(c_buffer, c_buf_len, c_d)
}

// AsciiFormatd is a wrapper around the C function g_ascii_formatd.
func AsciiFormatd(buffer string, bufLen int32, format string, d float64) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_buf_len := (C.gint)(bufLen)

	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	c_d := (C.gdouble)(d)

	C.g_ascii_formatd(c_buffer, c_buf_len, c_format, c_d)
}

// AsciiStrcasecmp is a wrapper around the C function g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	C.g_ascii_strcasecmp(c_s1, c_s2)
}

// AsciiStrdown is a wrapper around the C function g_ascii_strdown.
func AsciiStrdown(str string, len int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	C.g_ascii_strdown(c_str, c_len)
}

// AsciiStrncasecmp is a wrapper around the C function g_ascii_strncasecmp.
func AsciiStrncasecmp(s1 string, s2 string, n uint64) {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	c_n := (C.gsize)(n)

	C.g_ascii_strncasecmp(c_s1, c_s2, c_n)
}

// AsciiStrtod is a wrapper around the C function g_ascii_strtod.
func AsciiStrtod(nptr string, endptr string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	C.g_ascii_strtod(c_nptr, c_endptr)
}

// AsciiStrup is a wrapper around the C function g_ascii_strup.
func AsciiStrup(str string, len int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	C.g_ascii_strup(c_str, c_len)
}

// AsciiTolower is a wrapper around the C function g_ascii_tolower.
func AsciiTolower(c rune) {
	c_c := (C.gchar)(c)

	C.g_ascii_tolower(c_c)
}

// AsciiToupper is a wrapper around the C function g_ascii_toupper.
func AsciiToupper(c rune) {
	c_c := (C.gchar)(c)

	C.g_ascii_toupper(c_c)
}

// AsciiXdigitValue is a wrapper around the C function g_ascii_xdigit_value.
func AsciiXdigitValue(c rune) {
	c_c := (C.gchar)(c)

	C.g_ascii_xdigit_value(c_c)
}

// Unsupported : g_assert_warning : unsupported parameter line : type gint, const int

// Unsupported : g_assertion_message : unsupported parameter line : type gint, int

// Unsupported : g_assertion_message_cmpnum : unsupported parameter line : type gint, int

// Unsupported : g_assertion_message_cmpstr : unsupported parameter line : type gint, int

// Unsupported : g_assertion_message_error : unsupported parameter line : type gint, int

// Unsupported : g_assertion_message_expr : unsupported parameter line : type gint, int

// Unsupported : g_atexit : unsupported parameter func : type VoidFunc, GVoidFunc

// Unsupported : g_basename : unsupported parameter file_name : type filename, const gchar*

// BitNthLsf is a wrapper around the C function g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	C.g_bit_nth_lsf(c_mask, c_nth_bit)
}

// BitNthMsf is a wrapper around the C function g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	C.g_bit_nth_msf(c_mask, c_nth_bit)
}

// BitStorage is a wrapper around the C function g_bit_storage.
func BitStorage(number uint64) {
	c_number := (C.gulong)(number)

	C.g_bit_storage(c_number)
}

// BookmarkFileErrorQuark is a wrapper around the C function g_bookmark_file_error_quark.
func BookmarkFileErrorQuark() {
	C.g_bookmark_file_error_quark()
}

// Unsupported : g_build_filename : unsupported parameter first_element : type filename, const gchar*

// Unsupported : g_build_path : unsupported parameter separator : type filename, const gchar*

// Unsupported : g_byte_array_free : unsupported parameter array : no type

// ByteArrayNew is a wrapper around the C function g_byte_array_new.
func ByteArrayNew() {
	C.g_byte_array_new()
}

// ClearError is a wrapper around the C function g_clear_error.
func ClearError() {
	C.g_clear_error()
}

// Unsupported : g_convert : unsupported parameter str : no type

// ConvertErrorQuark is a wrapper around the C function g_convert_error_quark.
func ConvertErrorQuark() {
	C.g_convert_error_quark()
}

// Unsupported : g_convert_with_fallback : unsupported parameter str : no type

// Unsupported : g_convert_with_iconv : unsupported parameter str : no type

// Unsupported : g_datalist_clear : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_foreach : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_get_data : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_id_get_data : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_id_remove_no_notify : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_id_set_data_full : unsupported parameter datalist : type Data, GData**

// Unsupported : g_datalist_init : unsupported parameter datalist : type Data, GData**

// Unsupported : g_dataset_destroy : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported : g_dataset_foreach : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported : g_dataset_id_get_data : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported : g_dataset_id_remove_no_notify : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported : g_dataset_id_set_data_full : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported : g_date_get_days_in_month : unsupported parameter month : type DateMonth, GDateMonth

// Unsupported : g_date_get_monday_weeks_in_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported : g_date_get_sunday_weeks_in_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported : g_date_is_leap_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported : g_date_strftime : unsupported parameter date : type Date, const GDate*

// Unsupported : g_date_valid_day : unsupported parameter day : type DateDay, GDateDay

// Unsupported : g_date_valid_dmy : unsupported parameter day : type DateDay, GDateDay

// DateValidJulian is a wrapper around the C function g_date_valid_julian.
func DateValidJulian(julianDate uint32) {
	c_julian_date := (C.guint32)(julianDate)

	C.g_date_valid_julian(c_julian_date)
}

// Unsupported : g_date_valid_month : unsupported parameter month : type DateMonth, GDateMonth

// Unsupported : g_date_valid_weekday : unsupported parameter weekday : type DateWeekday, GDateWeekday

// Unsupported : g_date_valid_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported : g_direct_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported : g_direct_hash : unsupported parameter v : type gpointer, gconstpointer

// FileErrorFromErrno is a wrapper around the C function g_file_error_from_errno.
func FileErrorFromErrno(errNo int32) {
	c_err_no := (C.gint)(errNo)

	C.g_file_error_from_errno(c_err_no)
}

// FileErrorQuark is a wrapper around the C function g_file_error_quark.
func FileErrorQuark() {
	C.g_file_error_quark()
}

// Unsupported : g_file_get_contents : unsupported parameter filename : type filename, const gchar*

// Unsupported : g_file_open_tmp : unsupported parameter tmpl : type filename, const gchar*

// Unsupported : g_file_test : unsupported parameter filename : type filename, const gchar*

// FilenameFromUri is a wrapper around the C function g_filename_from_uri.
func FilenameFromUri(uri string, hostname string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.g_filename_from_uri(c_uri, c_hostname)
}

// Unsupported : g_filename_from_utf8 : unsupported parameter bytes_read : type gsize, gsize*

// Unsupported : g_filename_to_uri : unsupported parameter filename : type filename, const gchar*

// Unsupported : g_filename_to_utf8 : unsupported parameter opsysstring : type filename, const gchar*

// Unsupported : g_find_program_in_path : unsupported parameter program : type filename, const gchar*

// Free is a wrapper around the C function g_free.
func Free(mem uintptr) {
	c_mem := (C.gpointer)(mem)

	C.g_free(c_mem)
}

// GetCharset is a wrapper around the C function g_get_charset.
func GetCharset(charset string) {
	C.g_get_charset(c_charset)
}

// GetCodeset is a wrapper around the C function g_get_codeset.
func GetCodeset() {
	C.g_get_codeset()
}

// GetCurrentDir is a wrapper around the C function g_get_current_dir.
func GetCurrentDir() {
	C.g_get_current_dir()
}

// Unsupported : g_get_current_time : unsupported parameter result : type TimeVal, GTimeVal*

// GetHomeDir is a wrapper around the C function g_get_home_dir.
func GetHomeDir() {
	C.g_get_home_dir()
}

// GetPrgname is a wrapper around the C function g_get_prgname.
func GetPrgname() {
	C.g_get_prgname()
}

// GetRealName is a wrapper around the C function g_get_real_name.
func GetRealName() {
	C.g_get_real_name()
}

// GetTmpDir is a wrapper around the C function g_get_tmp_dir.
func GetTmpDir() {
	C.g_get_tmp_dir()
}

// GetUserName is a wrapper around the C function g_get_user_name.
func GetUserName() {
	C.g_get_user_name()
}

// Unsupported : g_getenv : unsupported parameter variable : type filename, const gchar*

// Unsupported : g_hash_table_destroy : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_insert : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_lookup : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_lookup_extended : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_remove : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_replace : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_size : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_steal : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported : g_hook_destroy : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported : g_hook_destroy_link : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported : g_hook_free : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported : g_hook_insert_before : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported : g_hook_prepend : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported : g_hook_unref : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported : g_iconv : unsupported parameter converter : type IConv, GIConv

// IconvOpen is a wrapper around the C function g_iconv_open.
func IconvOpen(toCodeset string, fromCodeset string) {
	c_to_codeset := C.CString(toCodeset)
	defer C.free(unsafe.Pointer(c_to_codeset))

	c_from_codeset := C.CString(fromCodeset)
	defer C.free(unsafe.Pointer(c_from_codeset))

	C.g_iconv_open(c_to_codeset, c_from_codeset)
}

// Unsupported : g_idle_add : unsupported parameter function : type SourceFunc, GSourceFunc

// Unsupported : g_idle_add_full : unsupported parameter function : type SourceFunc, GSourceFunc

// IdleRemoveByData is a wrapper around the C function g_idle_remove_by_data.
func IdleRemoveByData(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_idle_remove_by_data(c_data)
}

// IdleSourceNew is a wrapper around the C function g_idle_source_new.
func IdleSourceNew() {
	C.g_idle_source_new()
}

// Unsupported : g_int_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported : g_int_hash : unsupported parameter v : type gpointer, gconstpointer

// Unsupported : g_io_add_watch : unsupported parameter channel : type IOChannel, GIOChannel*

// Unsupported : g_io_add_watch_full : unsupported parameter channel : type IOChannel, GIOChannel*

// IoChannelErrorFromErrno is a wrapper around the C function g_io_channel_error_from_errno.
func IoChannelErrorFromErrno(en int32) {
	c_en := (C.gint)(en)

	C.g_io_channel_error_from_errno(c_en)
}

// IoChannelErrorQuark is a wrapper around the C function g_io_channel_error_quark.
func IoChannelErrorQuark() {
	C.g_io_channel_error_quark()
}

// Unsupported : g_io_create_watch : unsupported parameter channel : type IOChannel, GIOChannel*

// KeyFileErrorQuark is a wrapper around the C function g_key_file_error_quark.
func KeyFileErrorQuark() {
	C.g_key_file_error_quark()
}

// Unsupported : g_locale_from_utf8 : unsupported parameter bytes_read : type gsize, gsize*

// Unsupported : g_locale_to_utf8 : unsupported parameter opsysstring : no type

// Unsupported : g_log : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_default_handler : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// LogRemoveHandler is a wrapper around the C function g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_handler_id := (C.guint)(handlerId)

	C.g_log_remove_handler(c_log_domain, c_handler_id)
}

// Unsupported : g_log_set_always_fatal : unsupported parameter fatal_mask : type LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_set_fatal_mask : unsupported parameter fatal_mask : type LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_set_handler : unsupported parameter log_levels : type LogLevelFlags, GLogLevelFlags

// Unsupported : g_log_structured_standard : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported : g_logv : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// MainContextDefault is a wrapper around the C function g_main_context_default.
func MainContextDefault() {
	C.g_main_context_default()
}

// MainDepth is a wrapper around the C function g_main_depth.
func MainDepth() {
	C.g_main_depth()
}

// Malloc is a wrapper around the C function g_malloc.
func Malloc(nBytes uint64) {
	c_n_bytes := (C.gsize)(nBytes)

	C.g_malloc(c_n_bytes)
}

// Malloc0 is a wrapper around the C function g_malloc0.
func Malloc0(nBytes uint64) {
	c_n_bytes := (C.gsize)(nBytes)

	C.g_malloc0(c_n_bytes)
}

// MarkupErrorQuark is a wrapper around the C function g_markup_error_quark.
func MarkupErrorQuark() {
	C.g_markup_error_quark()
}

// MarkupEscapeText is a wrapper around the C function g_markup_escape_text.
func MarkupEscapeText(text string, length int64) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gssize)(length)

	C.g_markup_escape_text(c_text, c_length)
}

// MemIsSystemMalloc is a wrapper around the C function g_mem_is_system_malloc.
func MemIsSystemMalloc() {
	C.g_mem_is_system_malloc()
}

// MemProfile is a wrapper around the C function g_mem_profile.
func MemProfile() {
	C.g_mem_profile()
}

// Unsupported : g_mem_set_vtable : unsupported parameter vtable : type MemVTable, GMemVTable*

// Unsupported : g_memdup : unsupported parameter mem : type gpointer, gconstpointer

// Unsupported : g_mkstemp : unsupported parameter tmpl : type filename, gchar*

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : type gpointer, gpointer*

// Blacklisted : g_number_parser_error_quark

// OnErrorQuery is a wrapper around the C function g_on_error_query.
func OnErrorQuery(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_query(c_prg_name)
}

// OnErrorStackTrace is a wrapper around the C function g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_stack_trace(c_prg_name)
}

// OptionErrorQuark is a wrapper around the C function g_option_error_quark.
func OptionErrorQuark() {
	C.g_option_error_quark()
}

// Unsupported : g_parse_debug_string : unsupported parameter keys : no type

// Unsupported : g_path_get_basename : unsupported parameter file_name : type filename, const gchar*

// Unsupported : g_path_get_dirname : unsupported parameter file_name : type filename, const gchar*

// Unsupported : g_path_is_absolute : unsupported parameter file_name : type filename, const gchar*

// Unsupported : g_path_skip_root : unsupported parameter file_name : type filename, const gchar*

// Unsupported : g_pattern_match : unsupported parameter pspec : type PatternSpec, GPatternSpec*

// PatternMatchSimple is a wrapper around the C function g_pattern_match_simple.
func PatternMatchSimple(pattern string, string string) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_pattern_match_simple(c_pattern, c_string)
}

// Unsupported : g_pattern_match_string : unsupported parameter pspec : type PatternSpec, GPatternSpec*

// Unsupported : g_print : unsupported parameter ... : varargs

// Unsupported : g_printerr : unsupported parameter ... : varargs

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : type va_list, va_list

// Unsupported : g_propagate_error : unsupported parameter dest : type Error, GError**

// Unsupported : g_qsort_with_data : unsupported parameter pbase : type gpointer, gconstpointer

// QuarkFromStaticString is a wrapper around the C function g_quark_from_static_string.
func QuarkFromStaticString(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_quark_from_static_string(c_string)
}

// QuarkFromString is a wrapper around the C function g_quark_from_string.
func QuarkFromString(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_quark_from_string(c_string)
}

// Unsupported : g_quark_to_string : unsupported parameter quark : type Quark, GQuark

// QuarkTryString is a wrapper around the C function g_quark_try_string.
func QuarkTryString(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_quark_try_string(c_string)
}

// RandomDouble is a wrapper around the C function g_random_double.
func RandomDouble() {
	C.g_random_double()
}

// RandomDoubleRange is a wrapper around the C function g_random_double_range.
func RandomDoubleRange(begin float64, end float64) {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	C.g_random_double_range(c_begin, c_end)
}

// RandomInt is a wrapper around the C function g_random_int.
func RandomInt() {
	C.g_random_int()
}

// RandomIntRange is a wrapper around the C function g_random_int_range.
func RandomIntRange(begin int32, end int32) {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	C.g_random_int_range(c_begin, c_end)
}

// RandomSetSeed is a wrapper around the C function g_random_set_seed.
func RandomSetSeed(seed uint32) {
	c_seed := (C.guint32)(seed)

	C.g_random_set_seed(c_seed)
}

// Realloc is a wrapper around the C function g_realloc.
func Realloc(mem uintptr, nBytes uint64) {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	C.g_realloc(c_mem, c_n_bytes)
}

// RegexErrorQuark is a wrapper around the C function g_regex_error_quark.
func RegexErrorQuark() {
	C.g_regex_error_quark()
}

// ReturnIfFailWarning is a wrapper around the C function g_return_if_fail_warning.
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_pretty_function := C.CString(prettyFunction)
	defer C.free(unsafe.Pointer(c_pretty_function))

	c_expression := C.CString(expression)
	defer C.free(unsafe.Pointer(c_expression))

	C.g_return_if_fail_warning(c_log_domain, c_pretty_function, c_expression)
}

// Unsupported : g_set_error : unsupported parameter err : type Error, GError**

// SetPrgname is a wrapper around the C function g_set_prgname.
func SetPrgname(prgname string) {
	c_prgname := C.CString(prgname)
	defer C.free(unsafe.Pointer(c_prgname))

	C.g_set_prgname(c_prgname)
}

// Unsupported : g_set_print_handler : unsupported parameter func : type PrintFunc, GPrintFunc

// Unsupported : g_set_printerr_handler : unsupported parameter func : type PrintFunc, GPrintFunc

// ShellErrorQuark is a wrapper around the C function g_shell_error_quark.
func ShellErrorQuark() {
	C.g_shell_error_quark()
}

// Unsupported : g_shell_parse_argv : unsupported parameter command_line : type filename, const gchar*

// Unsupported : g_shell_quote : unsupported parameter unquoted_string : type filename, const gchar*

// Unsupported : g_shell_unquote : unsupported parameter quoted_string : type filename, const gchar*

// Unsupported : g_slice_get_config : unsupported parameter ckey : type SliceConfig, GSliceConfig

// Unsupported : g_slice_get_config_state : unsupported parameter ckey : type SliceConfig, GSliceConfig

// Unsupported : g_slice_set_config : unsupported parameter ckey : type SliceConfig, GSliceConfig

// Unsupported : g_snprintf : unsupported parameter ... : varargs

// SourceRemove is a wrapper around the C function g_source_remove.
func SourceRemove(tag uint32) {
	c_tag := (C.guint)(tag)

	C.g_source_remove(c_tag)
}

// Unsupported : g_source_remove_by_funcs_user_data : unsupported parameter funcs : type SourceFuncs, GSourceFuncs*

// SourceRemoveByUserData is a wrapper around the C function g_source_remove_by_user_data.
func SourceRemoveByUserData(userData uintptr) {
	c_user_data := (C.gpointer)(userData)

	C.g_source_remove_by_user_data(c_user_data)
}

// SpacedPrimesClosest is a wrapper around the C function g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) {
	c_num := (C.guint)(num)

	C.g_spaced_primes_closest(c_num)
}

// Unsupported : g_spawn_async : unsupported parameter working_directory : type filename, const gchar*

// Unsupported : g_spawn_async_with_pipes : unsupported parameter working_directory : type filename, const gchar*

// Unsupported : g_spawn_close_pid : unsupported parameter pid : type Pid, GPid

// Unsupported : g_spawn_command_line_async : unsupported parameter command_line : type filename, const gchar*

// Unsupported : g_spawn_command_line_sync : unsupported parameter command_line : type filename, const gchar*

// SpawnErrorQuark is a wrapper around the C function g_spawn_error_quark.
func SpawnErrorQuark() {
	C.g_spawn_error_quark()
}

// SpawnExitErrorQuark is a wrapper around the C function g_spawn_exit_error_quark.
func SpawnExitErrorQuark() {
	C.g_spawn_exit_error_quark()
}

// Unsupported : g_spawn_sync : unsupported parameter working_directory : type filename, const gchar*

// Stpcpy is a wrapper around the C function g_stpcpy.
func Stpcpy(dest string, src string) {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	C.g_stpcpy(c_dest, c_src)
}

// Unsupported : g_str_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported : g_str_hash : unsupported parameter v : type gpointer, gconstpointer

// Strcanon is a wrapper around the C function g_strcanon.
func Strcanon(string string, validChars string, substitutor rune) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_valid_chars := C.CString(validChars)
	defer C.free(unsafe.Pointer(c_valid_chars))

	c_substitutor := (C.gchar)(substitutor)

	C.g_strcanon(c_string, c_valid_chars, c_substitutor)
}

// Strcasecmp is a wrapper around the C function g_strcasecmp.
func Strcasecmp(s1 string, s2 string) {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	C.g_strcasecmp(c_s1, c_s2)
}

// Strchomp is a wrapper around the C function g_strchomp.
func Strchomp(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_strchomp(c_string)
}

// Strchug is a wrapper around the C function g_strchug.
func Strchug(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_strchug(c_string)
}

// Strcompress is a wrapper around the C function g_strcompress.
func Strcompress(source string) {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	C.g_strcompress(c_source)
}

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Strdelimit is a wrapper around the C function g_strdelimit.
func Strdelimit(string string, delimiters string, newDelimiter rune) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiters := C.CString(delimiters)
	defer C.free(unsafe.Pointer(c_delimiters))

	c_new_delimiter := (C.gchar)(newDelimiter)

	C.g_strdelimit(c_string, c_delimiters, c_new_delimiter)
}

// Strdown is a wrapper around the C function g_strdown.
func Strdown(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_strdown(c_string)
}

// Strdup is a wrapper around the C function g_strdup.
func Strdup(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.g_strdup(c_str)
}

// Unsupported : g_strdup_printf : unsupported parameter ... : varargs

// Unsupported : g_strdup_vprintf : unsupported parameter args : type va_list, va_list

// Strdupv is a wrapper around the C function g_strdupv.
func Strdupv(strArray string) {
	c_str_array := C.CString(strArray)
	defer C.free(unsafe.Pointer(c_str_array))

	C.g_strdupv(c_str_array)
}

// Strerror is a wrapper around the C function g_strerror.
func Strerror(errnum int32) {
	c_errnum := (C.gint)(errnum)

	C.g_strerror(c_errnum)
}

// Strescape is a wrapper around the C function g_strescape.
func Strescape(source string, exceptions string) {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	c_exceptions := C.CString(exceptions)
	defer C.free(unsafe.Pointer(c_exceptions))

	C.g_strescape(c_source, c_exceptions)
}

// Strfreev is a wrapper around the C function g_strfreev.
func Strfreev(strArray string) {
	c_str_array := C.CString(strArray)
	defer C.free(unsafe.Pointer(c_str_array))

	C.g_strfreev(c_str_array)
}

// StringNew is a wrapper around the C function g_string_new.
func StringNew(init string) {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	C.g_string_new(c_init)
}

// StringNewLen is a wrapper around the C function g_string_new_len.
func StringNewLen(init string, len int64) {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	c_len := (C.gssize)(len)

	C.g_string_new_len(c_init, c_len)
}

// StringSizedNew is a wrapper around the C function g_string_sized_new.
func StringSizedNew(dflSize uint64) {
	c_dfl_size := (C.gsize)(dflSize)

	C.g_string_sized_new(c_dfl_size)
}

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Strjoinv is a wrapper around the C function g_strjoinv.
func Strjoinv(separator string, strArray string) {
	c_separator := C.CString(separator)
	defer C.free(unsafe.Pointer(c_separator))

	c_str_array := C.CString(strArray)
	defer C.free(unsafe.Pointer(c_str_array))

	C.g_strjoinv(c_separator, c_str_array)
}

// Strlcat is a wrapper around the C function g_strlcat.
func Strlcat(dest string, src string, destSize uint64) {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_dest_size := (C.gsize)(destSize)

	C.g_strlcat(c_dest, c_src, c_dest_size)
}

// Strlcpy is a wrapper around the C function g_strlcpy.
func Strlcpy(dest string, src string, destSize uint64) {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_dest_size := (C.gsize)(destSize)

	C.g_strlcpy(c_dest, c_src, c_dest_size)
}

// Strncasecmp is a wrapper around the C function g_strncasecmp.
func Strncasecmp(s1 string, s2 string, n uint32) {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	c_n := (C.guint)(n)

	C.g_strncasecmp(c_s1, c_s2, c_n)
}

// Strndup is a wrapper around the C function g_strndup.
func Strndup(str string, n uint64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_n := (C.gsize)(n)

	C.g_strndup(c_str, c_n)
}

// Strnfill is a wrapper around the C function g_strnfill.
func Strnfill(length uint64, fillChar rune) {
	c_length := (C.gsize)(length)

	c_fill_char := (C.gchar)(fillChar)

	C.g_strnfill(c_length, c_fill_char)
}

// Strreverse is a wrapper around the C function g_strreverse.
func Strreverse(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_strreverse(c_string)
}

// Strrstr is a wrapper around the C function g_strrstr.
func Strrstr(haystack string, needle string) {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	C.g_strrstr(c_haystack, c_needle)
}

// StrrstrLen is a wrapper around the C function g_strrstr_len.
func StrrstrLen(haystack string, haystackLen int64, needle string) {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_haystack_len := (C.gssize)(haystackLen)

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	C.g_strrstr_len(c_haystack, c_haystack_len, c_needle)
}

// Strsignal is a wrapper around the C function g_strsignal.
func Strsignal(signum int32) {
	c_signum := (C.gint)(signum)

	C.g_strsignal(c_signum)
}

// Strsplit is a wrapper around the C function g_strsplit.
func Strsplit(string string, delimiter string, maxTokens int32) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiter := C.CString(delimiter)
	defer C.free(unsafe.Pointer(c_delimiter))

	c_max_tokens := (C.gint)(maxTokens)

	C.g_strsplit(c_string, c_delimiter, c_max_tokens)
}

// StrstrLen is a wrapper around the C function g_strstr_len.
func StrstrLen(haystack string, haystackLen int64, needle string) {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_haystack_len := (C.gssize)(haystackLen)

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	C.g_strstr_len(c_haystack, c_haystack_len, c_needle)
}

// Strtod is a wrapper around the C function g_strtod.
func Strtod(nptr string, endptr string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	C.g_strtod(c_nptr, c_endptr)
}

// Strup is a wrapper around the C function g_strup.
func Strup(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_strup(c_string)
}

// Blacklisted : g_strv_get_type

// Unsupported : g_test_add_vtable : unsupported parameter test_data : type gpointer, gconstpointer

// Unsupported : g_test_assert_expected_messages_internal : unsupported parameter line : type gint, int

// Unsupported : g_test_log_type_name : unsupported parameter log_type : type TestLogType, GTestLogType

// Unsupported : g_test_trap_assertions : unsupported parameter line : type gint, int

// ThreadErrorQuark is a wrapper around the C function g_thread_error_quark.
func ThreadErrorQuark() {
	C.g_thread_error_quark()
}

// ThreadExit is a wrapper around the C function g_thread_exit.
func ThreadExit(retval uintptr) {
	c_retval := (C.gpointer)(retval)

	C.g_thread_exit(c_retval)
}

// ThreadPoolGetMaxUnusedThreads is a wrapper around the C function g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() {
	C.g_thread_pool_get_max_unused_threads()
}

// ThreadPoolGetNumUnusedThreads is a wrapper around the C function g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() {
	C.g_thread_pool_get_num_unused_threads()
}

// ThreadPoolSetMaxUnusedThreads is a wrapper around the C function g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	c_max_threads := (C.gint)(maxThreads)

	C.g_thread_pool_set_max_unused_threads(c_max_threads)
}

// ThreadPoolStopUnusedThreads is a wrapper around the C function g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {
	C.g_thread_pool_stop_unused_threads()
}

// ThreadSelf is a wrapper around the C function g_thread_self.
func ThreadSelf() {
	C.g_thread_self()
}

// ThreadYield is a wrapper around the C function g_thread_yield.
func ThreadYield() {
	C.g_thread_yield()
}

// Unsupported : g_timeout_add : unsupported parameter function : type SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_full : unsupported parameter function : type SourceFunc, GSourceFunc

// TimeoutSourceNew is a wrapper around the C function g_timeout_source_new.
func TimeoutSourceNew(interval uint32) {
	c_interval := (C.guint)(interval)

	C.g_timeout_source_new(c_interval)
}

// Unsupported : g_trash_stack_height : unsupported parameter stack_p : type TrashStack, GTrashStack**

// Unsupported : g_trash_stack_peek : unsupported parameter stack_p : type TrashStack, GTrashStack**

// Unsupported : g_trash_stack_pop : unsupported parameter stack_p : type TrashStack, GTrashStack**

// Unsupported : g_trash_stack_push : unsupported parameter stack_p : type TrashStack, GTrashStack**

// TryMalloc is a wrapper around the C function g_try_malloc.
func TryMalloc(nBytes uint64) {
	c_n_bytes := (C.gsize)(nBytes)

	C.g_try_malloc(c_n_bytes)
}

// TryRealloc is a wrapper around the C function g_try_realloc.
func TryRealloc(mem uintptr, nBytes uint64) {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	C.g_try_realloc(c_mem, c_n_bytes)
}

// Unsupported : g_ucs4_to_utf16 : unsupported parameter str : type gunichar, const gunichar*

// Unsupported : g_ucs4_to_utf8 : unsupported parameter str : type gunichar, const gunichar*

// UnicharBreakType is a wrapper around the C function g_unichar_break_type.
func UnicharBreakType(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_break_type(c_c)
}

// UnicharDigitValue is a wrapper around the C function g_unichar_digit_value.
func UnicharDigitValue(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_digit_value(c_c)
}

// UnicharIsalnum is a wrapper around the C function g_unichar_isalnum.
func UnicharIsalnum(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isalnum(c_c)
}

// UnicharIsalpha is a wrapper around the C function g_unichar_isalpha.
func UnicharIsalpha(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isalpha(c_c)
}

// UnicharIscntrl is a wrapper around the C function g_unichar_iscntrl.
func UnicharIscntrl(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_iscntrl(c_c)
}

// UnicharIsdefined is a wrapper around the C function g_unichar_isdefined.
func UnicharIsdefined(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isdefined(c_c)
}

// UnicharIsdigit is a wrapper around the C function g_unichar_isdigit.
func UnicharIsdigit(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isdigit(c_c)
}

// UnicharIsgraph is a wrapper around the C function g_unichar_isgraph.
func UnicharIsgraph(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isgraph(c_c)
}

// UnicharIslower is a wrapper around the C function g_unichar_islower.
func UnicharIslower(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_islower(c_c)
}

// UnicharIsprint is a wrapper around the C function g_unichar_isprint.
func UnicharIsprint(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isprint(c_c)
}

// UnicharIspunct is a wrapper around the C function g_unichar_ispunct.
func UnicharIspunct(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_ispunct(c_c)
}

// UnicharIsspace is a wrapper around the C function g_unichar_isspace.
func UnicharIsspace(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isspace(c_c)
}

// UnicharIstitle is a wrapper around the C function g_unichar_istitle.
func UnicharIstitle(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_istitle(c_c)
}

// UnicharIsupper is a wrapper around the C function g_unichar_isupper.
func UnicharIsupper(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isupper(c_c)
}

// UnicharIswide is a wrapper around the C function g_unichar_iswide.
func UnicharIswide(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_iswide(c_c)
}

// UnicharIsxdigit is a wrapper around the C function g_unichar_isxdigit.
func UnicharIsxdigit(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_isxdigit(c_c)
}

// UnicharToUtf8 is a wrapper around the C function g_unichar_to_utf8.
func UnicharToUtf8(c rune, outbuf string) {
	c_c := (C.gunichar)(c)

	C.g_unichar_to_utf8(c_c, c_outbuf)
}

// UnicharTolower is a wrapper around the C function g_unichar_tolower.
func UnicharTolower(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_tolower(c_c)
}

// UnicharTotitle is a wrapper around the C function g_unichar_totitle.
func UnicharTotitle(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_totitle(c_c)
}

// UnicharToupper is a wrapper around the C function g_unichar_toupper.
func UnicharToupper(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_toupper(c_c)
}

// UnicharType is a wrapper around the C function g_unichar_type.
func UnicharType(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_type(c_c)
}

// UnicharValidate is a wrapper around the C function g_unichar_validate.
func UnicharValidate(ch rune) {
	c_ch := (C.gunichar)(ch)

	C.g_unichar_validate(c_ch)
}

// UnicharXdigitValue is a wrapper around the C function g_unichar_xdigit_value.
func UnicharXdigitValue(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_xdigit_value(c_c)
}

// Unsupported : g_unicode_canonical_decomposition : unsupported parameter result_len : type gsize, gsize*

// Unsupported : g_unicode_canonical_ordering : unsupported parameter string : type gunichar, gunichar*

// Blacklisted : g_unix_error_quark

// Usleep is a wrapper around the C function g_usleep.
func Usleep(microseconds uint64) {
	c_microseconds := (C.gulong)(microseconds)

	C.g_usleep(c_microseconds)
}

// Unsupported : g_utf16_to_ucs4 : unsupported parameter str : type guint16, const gunichar2*

// Unsupported : g_utf16_to_utf8 : unsupported parameter str : type guint16, const gunichar2*

// Utf8Casefold is a wrapper around the C function g_utf8_casefold.
func Utf8Casefold(str string, len int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	C.g_utf8_casefold(c_str, c_len)
}

// Utf8Collate is a wrapper around the C function g_utf8_collate.
func Utf8Collate(str1 string, str2 string) {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	C.g_utf8_collate(c_str1, c_str2)
}

// Utf8CollateKey is a wrapper around the C function g_utf8_collate_key.
func Utf8CollateKey(str string, len int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	C.g_utf8_collate_key(c_str, c_len)
}

// Utf8FindNextChar is a wrapper around the C function g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_end := C.CString(end)
	defer C.free(unsafe.Pointer(c_end))

	C.g_utf8_find_next_char(c_p, c_end)
}

// Utf8FindPrevChar is a wrapper around the C function g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	C.g_utf8_find_prev_char(c_str, c_p)
}

// Utf8GetChar is a wrapper around the C function g_utf8_get_char.
func Utf8GetChar(p string) {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	C.g_utf8_get_char(c_p)
}

// Utf8GetCharValidated is a wrapper around the C function g_utf8_get_char_validated.
func Utf8GetCharValidated(p string, maxLen int64) {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max_len := (C.gssize)(maxLen)

	C.g_utf8_get_char_validated(c_p, c_max_len)
}

// Unsupported : g_utf8_normalize : unsupported parameter mode : type NormalizeMode, GNormalizeMode

// Utf8OffsetToPointer is a wrapper around the C function g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_offset := (C.glong)(offset)

	C.g_utf8_offset_to_pointer(c_str, c_offset)
}

// Utf8PointerToOffset is a wrapper around the C function g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_pos := C.CString(pos)
	defer C.free(unsafe.Pointer(c_pos))

	C.g_utf8_pointer_to_offset(c_str, c_pos)
}

// Utf8PrevChar is a wrapper around the C function g_utf8_prev_char.
func Utf8PrevChar(p string) {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	C.g_utf8_prev_char(c_p)
}

// Utf8Strchr is a wrapper around the C function g_utf8_strchr.
func Utf8Strchr(p string, len int64, c rune) {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	C.g_utf8_strchr(c_p, c_len, c_c)
}

// Utf8Strdown is a wrapper around the C function g_utf8_strdown.
func Utf8Strdown(str string, len int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	C.g_utf8_strdown(c_str, c_len)
}

// Utf8Strlen is a wrapper around the C function g_utf8_strlen.
func Utf8Strlen(p string, max int64) {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max := (C.gssize)(max)

	C.g_utf8_strlen(c_p, c_max)
}

// Utf8Strncpy is a wrapper around the C function g_utf8_strncpy.
func Utf8Strncpy(dest string, src string, n uint64) {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_n := (C.gsize)(n)

	C.g_utf8_strncpy(c_dest, c_src, c_n)
}

// Utf8Strrchr is a wrapper around the C function g_utf8_strrchr.
func Utf8Strrchr(p string, len int64, c rune) {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	C.g_utf8_strrchr(c_p, c_len, c_c)
}

// Utf8Strup is a wrapper around the C function g_utf8_strup.
func Utf8Strup(str string, len int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	C.g_utf8_strup(c_str, c_len)
}

// Unsupported : g_utf8_to_ucs4 : unsupported parameter items_read : type glong, glong*

// Unsupported : g_utf8_to_ucs4_fast : unsupported parameter items_written : type glong, glong*

// Unsupported : g_utf8_to_utf16 : unsupported parameter items_read : type glong, glong*

// Unsupported : g_utf8_validate : unsupported parameter str : no type

// Blacklisted : g_variant_get_gtype

// Unsupported : g_variant_parse : unsupported parameter type : type VariantType, const GVariantType*

// VariantParseErrorQuark is a wrapper around the C function g_variant_parse_error_quark.
func VariantParseErrorQuark() {
	C.g_variant_parse_error_quark()
}

// VariantParserGetErrorQuark is a wrapper around the C function g_variant_parser_get_error_quark.
func VariantParserGetErrorQuark() {
	C.g_variant_parser_get_error_quark()
}

// Blacklisted : g_variant_type_checked_

// VariantTypeStringIsValid is a wrapper around the C function g_variant_type_string_is_valid.
func VariantTypeStringIsValid(typeString string) {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	C.g_variant_type_string_is_valid(c_type_string)
}

// Unsupported : g_vsnprintf : unsupported parameter args : type va_list, va_list

// Unsupported : g_warn_message : unsupported parameter line : type gint, int

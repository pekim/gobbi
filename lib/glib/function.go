package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// AsciiDigitValue is a wrapper around the C function g_ascii_digit_value.
func AsciiDigitValue(c rune) {}

// AsciiDtostr is a wrapper around the C function g_ascii_dtostr.
func AsciiDtostr(buffer string, bufLen int32, d float64) {}

// AsciiFormatd is a wrapper around the C function g_ascii_formatd.
func AsciiFormatd(buffer string, bufLen int32, format string, d float64) {}

// AsciiStrcasecmp is a wrapper around the C function g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) {}

// AsciiStrdown is a wrapper around the C function g_ascii_strdown.
func AsciiStrdown(str string, len int64) {}

// AsciiStrncasecmp is a wrapper around the C function g_ascii_strncasecmp.
func AsciiStrncasecmp(s1 string, s2 string, n uint64) {}

// AsciiStrtod is a wrapper around the C function g_ascii_strtod.
func AsciiStrtod(nptr string, endptr string) {}

// AsciiStrup is a wrapper around the C function g_ascii_strup.
func AsciiStrup(str string, len int64) {}

// AsciiTolower is a wrapper around the C function g_ascii_tolower.
func AsciiTolower(c rune) {}

// AsciiToupper is a wrapper around the C function g_ascii_toupper.
func AsciiToupper(c rune) {}

// AsciiXdigitValue is a wrapper around the C function g_ascii_xdigit_value.
func AsciiXdigitValue(c rune) {}

// Unsupported function: g_assert_warning : unsupported parameter line : type gint, const int

// Unsupported function: g_assertion_message : unsupported parameter line : type gint, int

// Blacklisted function: g_assertion_message_cmpnum

// Unsupported function: g_assertion_message_cmpstr : unsupported parameter line : type gint, int

// Blacklisted function: g_assertion_message_error

// Unsupported function: g_assertion_message_expr : unsupported parameter line : type gint, int

// Unsupported function: g_atexit : unsupported parameter func : type VoidFunc, GVoidFunc

// Unsupported function: g_basename : unsupported parameter file_name : type filename, const gchar*

// BitNthLsf is a wrapper around the C function g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) {}

// BitNthMsf is a wrapper around the C function g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) {}

// BitStorage is a wrapper around the C function g_bit_storage.
func BitStorage(number uint64) {}

// BookmarkFileErrorQuark is a wrapper around the C function g_bookmark_file_error_quark.
func BookmarkFileErrorQuark() {}

// Unsupported function: g_build_filename : unsupported parameter first_element : type filename, const gchar*

// Unsupported function: g_build_path : unsupported parameter separator : type filename, const gchar*

// Unsupported function: g_byte_array_free : unsupported parameter array : no type

// ByteArrayNew is a wrapper around the C function g_byte_array_new.
func ByteArrayNew() {}

// ClearError is a wrapper around the C function g_clear_error.
func ClearError() {}

// Unsupported function: g_convert : unsupported parameter str : no type

// ConvertErrorQuark is a wrapper around the C function g_convert_error_quark.
func ConvertErrorQuark() {}

// Unsupported function: g_convert_with_fallback : unsupported parameter str : no type

// Unsupported function: g_convert_with_iconv : unsupported parameter str : no type

// Unsupported function: g_datalist_clear : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_datalist_foreach : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_datalist_get_data : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_datalist_id_get_data : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_datalist_id_remove_no_notify : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_datalist_id_set_data_full : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_datalist_init : unsupported parameter datalist : type Data, GData**

// Unsupported function: g_dataset_destroy : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported function: g_dataset_foreach : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported function: g_dataset_id_get_data : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported function: g_dataset_id_remove_no_notify : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported function: g_dataset_id_set_data_full : unsupported parameter dataset_location : type gpointer, gconstpointer

// Unsupported function: g_date_get_days_in_month : unsupported parameter month : type DateMonth, GDateMonth

// Unsupported function: g_date_get_monday_weeks_in_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported function: g_date_get_sunday_weeks_in_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported function: g_date_is_leap_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported function: g_date_strftime : unsupported parameter date : type Date, const GDate*

// Unsupported function: g_date_valid_day : unsupported parameter day : type DateDay, GDateDay

// Unsupported function: g_date_valid_dmy : unsupported parameter day : type DateDay, GDateDay

// DateValidJulian is a wrapper around the C function g_date_valid_julian.
func DateValidJulian(julianDate uint32) {}

// Unsupported function: g_date_valid_month : unsupported parameter month : type DateMonth, GDateMonth

// Unsupported function: g_date_valid_weekday : unsupported parameter weekday : type DateWeekday, GDateWeekday

// Unsupported function: g_date_valid_year : unsupported parameter year : type DateYear, GDateYear

// Unsupported function: g_direct_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported function: g_direct_hash : unsupported parameter v : type gpointer, gconstpointer

// FileErrorFromErrno is a wrapper around the C function g_file_error_from_errno.
func FileErrorFromErrno(errNo int32) {}

// FileErrorQuark is a wrapper around the C function g_file_error_quark.
func FileErrorQuark() {}

// Unsupported function: g_file_get_contents : unsupported parameter filename : type filename, const gchar*

// Unsupported function: g_file_open_tmp : unsupported parameter tmpl : type filename, const gchar*

// Unsupported function: g_file_test : unsupported parameter filename : type filename, const gchar*

// FilenameFromUri is a wrapper around the C function g_filename_from_uri.
func FilenameFromUri(uri string, hostname string) {}

// Unsupported function: g_filename_from_utf8 : unsupported parameter bytes_read : type gsize, gsize*

// Unsupported function: g_filename_to_uri : unsupported parameter filename : type filename, const gchar*

// Unsupported function: g_filename_to_utf8 : unsupported parameter opsysstring : type filename, const gchar*

// Unsupported function: g_find_program_in_path : unsupported parameter program : type filename, const gchar*

// Free is a wrapper around the C function g_free.
func Free(mem uintptr) {}

// GetCharset is a wrapper around the C function g_get_charset.
func GetCharset(charset string) {}

// GetCodeset is a wrapper around the C function g_get_codeset.
func GetCodeset() {}

// GetCurrentDir is a wrapper around the C function g_get_current_dir.
func GetCurrentDir() {}

// Unsupported function: g_get_current_time : unsupported parameter result : type TimeVal, GTimeVal*

// GetHomeDir is a wrapper around the C function g_get_home_dir.
func GetHomeDir() {}

// GetPrgname is a wrapper around the C function g_get_prgname.
func GetPrgname() {}

// GetRealName is a wrapper around the C function g_get_real_name.
func GetRealName() {}

// GetTmpDir is a wrapper around the C function g_get_tmp_dir.
func GetTmpDir() {}

// GetUserName is a wrapper around the C function g_get_user_name.
func GetUserName() {}

// Unsupported function: g_getenv : unsupported parameter variable : type filename, const gchar*

// Unsupported function: g_hash_table_destroy : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hash_table_insert : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hash_table_lookup : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hash_table_lookup_extended : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hash_table_remove : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hash_table_replace : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hash_table_size : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hash_table_steal : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// Unsupported function: g_hook_destroy : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported function: g_hook_destroy_link : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported function: g_hook_free : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported function: g_hook_insert_before : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported function: g_hook_prepend : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported function: g_hook_unref : unsupported parameter hook_list : type HookList, GHookList*

// Unsupported function: g_iconv : unsupported parameter converter : type IConv, GIConv

// IconvOpen is a wrapper around the C function g_iconv_open.
func IconvOpen(toCodeset string, fromCodeset string) {}

// Unsupported function: g_idle_add : unsupported parameter function : type SourceFunc, GSourceFunc

// Unsupported function: g_idle_add_full : unsupported parameter function : type SourceFunc, GSourceFunc

// IdleRemoveByData is a wrapper around the C function g_idle_remove_by_data.
func IdleRemoveByData(data uintptr) {}

// IdleSourceNew is a wrapper around the C function g_idle_source_new.
func IdleSourceNew() {}

// Unsupported function: g_int_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported function: g_int_hash : unsupported parameter v : type gpointer, gconstpointer

// Unsupported function: g_io_add_watch : unsupported parameter channel : type IOChannel, GIOChannel*

// Unsupported function: g_io_add_watch_full : unsupported parameter channel : type IOChannel, GIOChannel*

// IoChannelErrorFromErrno is a wrapper around the C function g_io_channel_error_from_errno.
func IoChannelErrorFromErrno(en int32) {}

// IoChannelErrorQuark is a wrapper around the C function g_io_channel_error_quark.
func IoChannelErrorQuark() {}

// Unsupported function: g_io_create_watch : unsupported parameter channel : type IOChannel, GIOChannel*

// KeyFileErrorQuark is a wrapper around the C function g_key_file_error_quark.
func KeyFileErrorQuark() {}

// Unsupported function: g_locale_from_utf8 : unsupported parameter bytes_read : type gsize, gsize*

// Unsupported function: g_locale_to_utf8 : unsupported parameter opsysstring : no type

// Unsupported function: g_log : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_default_handler : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// LogRemoveHandler is a wrapper around the C function g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) {}

// Unsupported function: g_log_set_always_fatal : unsupported parameter fatal_mask : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_set_fatal_mask : unsupported parameter fatal_mask : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_set_handler : unsupported parameter log_levels : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_structured_standard : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_logv : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// MainContextDefault is a wrapper around the C function g_main_context_default.
func MainContextDefault() {}

// MainDepth is a wrapper around the C function g_main_depth.
func MainDepth() {}

// Malloc is a wrapper around the C function g_malloc.
func Malloc(nBytes uint64) {}

// Malloc0 is a wrapper around the C function g_malloc0.
func Malloc0(nBytes uint64) {}

// MarkupErrorQuark is a wrapper around the C function g_markup_error_quark.
func MarkupErrorQuark() {}

// MarkupEscapeText is a wrapper around the C function g_markup_escape_text.
func MarkupEscapeText(text string, length int64) {}

// MemIsSystemMalloc is a wrapper around the C function g_mem_is_system_malloc.
func MemIsSystemMalloc() {}

// MemProfile is a wrapper around the C function g_mem_profile.
func MemProfile() {}

// Unsupported function: g_mem_set_vtable : unsupported parameter vtable : type MemVTable, GMemVTable*

// Unsupported function: g_memdup : unsupported parameter mem : type gpointer, gconstpointer

// Unsupported function: g_mkstemp : unsupported parameter tmpl : type filename, gchar*

// Blacklisted function: g_nullify_pointer

// NumberParserErrorQuark is a wrapper around the C function g_number_parser_error_quark.
func NumberParserErrorQuark() {}

// OnErrorQuery is a wrapper around the C function g_on_error_query.
func OnErrorQuery(prgName string) {}

// OnErrorStackTrace is a wrapper around the C function g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) {}

// OptionErrorQuark is a wrapper around the C function g_option_error_quark.
func OptionErrorQuark() {}

// Unsupported function: g_parse_debug_string : unsupported parameter keys : no type

// Unsupported function: g_path_get_basename : unsupported parameter file_name : type filename, const gchar*

// Unsupported function: g_path_get_dirname : unsupported parameter file_name : type filename, const gchar*

// Unsupported function: g_path_is_absolute : unsupported parameter file_name : type filename, const gchar*

// Unsupported function: g_path_skip_root : unsupported parameter file_name : type filename, const gchar*

// Unsupported function: g_pattern_match : unsupported parameter pspec : type PatternSpec, GPatternSpec*

// PatternMatchSimple is a wrapper around the C function g_pattern_match_simple.
func PatternMatchSimple(pattern string, string string) {}

// Unsupported function: g_pattern_match_string : unsupported parameter pspec : type PatternSpec, GPatternSpec*

// Unsupported function: g_print : unsupported parameter ... : varargs

// Unsupported function: g_printerr : unsupported parameter ... : varargs

// Unsupported function: g_printf_string_upper_bound : unsupported parameter args : type va_list, va_list

// Unsupported function: g_propagate_error : unsupported parameter dest : type Error, GError**

// Unsupported function: g_qsort_with_data : unsupported parameter pbase : type gpointer, gconstpointer

// QuarkFromStaticString is a wrapper around the C function g_quark_from_static_string.
func QuarkFromStaticString(string string) {}

// QuarkFromString is a wrapper around the C function g_quark_from_string.
func QuarkFromString(string string) {}

// Unsupported function: g_quark_to_string : unsupported parameter quark : type Quark, GQuark

// QuarkTryString is a wrapper around the C function g_quark_try_string.
func QuarkTryString(string string) {}

// RandomDouble is a wrapper around the C function g_random_double.
func RandomDouble() {}

// RandomDoubleRange is a wrapper around the C function g_random_double_range.
func RandomDoubleRange(begin float64, end float64) {}

// RandomInt is a wrapper around the C function g_random_int.
func RandomInt() {}

// RandomIntRange is a wrapper around the C function g_random_int_range.
func RandomIntRange(begin int32, end int32) {}

// RandomSetSeed is a wrapper around the C function g_random_set_seed.
func RandomSetSeed(seed uint32) {}

// Realloc is a wrapper around the C function g_realloc.
func Realloc(mem uintptr, nBytes uint64) {}

// RegexErrorQuark is a wrapper around the C function g_regex_error_quark.
func RegexErrorQuark() {}

// ReturnIfFailWarning is a wrapper around the C function g_return_if_fail_warning.
func ReturnIfFailWarning(logDomain string, prettyFunction string, expression string) {}

// Unsupported function: g_set_error : unsupported parameter err : type Error, GError**

// SetPrgname is a wrapper around the C function g_set_prgname.
func SetPrgname(prgname string) {}

// Unsupported function: g_set_print_handler : unsupported parameter func : type PrintFunc, GPrintFunc

// Unsupported function: g_set_printerr_handler : unsupported parameter func : type PrintFunc, GPrintFunc

// ShellErrorQuark is a wrapper around the C function g_shell_error_quark.
func ShellErrorQuark() {}

// Unsupported function: g_shell_parse_argv : unsupported parameter command_line : type filename, const gchar*

// Unsupported function: g_shell_quote : unsupported parameter unquoted_string : type filename, const gchar*

// Unsupported function: g_shell_unquote : unsupported parameter quoted_string : type filename, const gchar*

// Unsupported function: g_slice_get_config : unsupported parameter ckey : type SliceConfig, GSliceConfig

// Blacklisted function: g_slice_get_config_state

// Unsupported function: g_slice_set_config : unsupported parameter ckey : type SliceConfig, GSliceConfig

// Unsupported function: g_snprintf : unsupported parameter ... : varargs

// SourceRemove is a wrapper around the C function g_source_remove.
func SourceRemove(tag uint32) {}

// Unsupported function: g_source_remove_by_funcs_user_data : unsupported parameter funcs : type SourceFuncs, GSourceFuncs*

// SourceRemoveByUserData is a wrapper around the C function g_source_remove_by_user_data.
func SourceRemoveByUserData(userData uintptr) {}

// SpacedPrimesClosest is a wrapper around the C function g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) {}

// Unsupported function: g_spawn_async : unsupported parameter working_directory : type filename, const gchar*

// Unsupported function: g_spawn_async_with_pipes : unsupported parameter working_directory : type filename, const gchar*

// Unsupported function: g_spawn_close_pid : unsupported parameter pid : type Pid, GPid

// Unsupported function: g_spawn_command_line_async : unsupported parameter command_line : type filename, const gchar*

// Unsupported function: g_spawn_command_line_sync : unsupported parameter command_line : type filename, const gchar*

// SpawnErrorQuark is a wrapper around the C function g_spawn_error_quark.
func SpawnErrorQuark() {}

// SpawnExitErrorQuark is a wrapper around the C function g_spawn_exit_error_quark.
func SpawnExitErrorQuark() {}

// Unsupported function: g_spawn_sync : unsupported parameter working_directory : type filename, const gchar*

// Stpcpy is a wrapper around the C function g_stpcpy.
func Stpcpy(dest string, src string) {}

// Unsupported function: g_str_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported function: g_str_hash : unsupported parameter v : type gpointer, gconstpointer

// Strcanon is a wrapper around the C function g_strcanon.
func Strcanon(string string, validChars string, substitutor rune) {}

// Strcasecmp is a wrapper around the C function g_strcasecmp.
func Strcasecmp(s1 string, s2 string) {}

// Strchomp is a wrapper around the C function g_strchomp.
func Strchomp(string string) {}

// Strchug is a wrapper around the C function g_strchug.
func Strchug(string string) {}

// Strcompress is a wrapper around the C function g_strcompress.
func Strcompress(source string) {}

// Unsupported function: g_strconcat : unsupported parameter ... : varargs

// Strdelimit is a wrapper around the C function g_strdelimit.
func Strdelimit(string string, delimiters string, newDelimiter rune) {}

// Strdown is a wrapper around the C function g_strdown.
func Strdown(string string) {}

// Strdup is a wrapper around the C function g_strdup.
func Strdup(str string) {}

// Unsupported function: g_strdup_printf : unsupported parameter ... : varargs

// Unsupported function: g_strdup_vprintf : unsupported parameter args : type va_list, va_list

// Strdupv is a wrapper around the C function g_strdupv.
func Strdupv(strArray string) {}

// Strerror is a wrapper around the C function g_strerror.
func Strerror(errnum int32) {}

// Strescape is a wrapper around the C function g_strescape.
func Strescape(source string, exceptions string) {}

// Strfreev is a wrapper around the C function g_strfreev.
func Strfreev(strArray string) {}

// StringNew is a wrapper around the C function g_string_new.
func StringNew(init string) {}

// StringNewLen is a wrapper around the C function g_string_new_len.
func StringNewLen(init string, len int64) {}

// StringSizedNew is a wrapper around the C function g_string_sized_new.
func StringSizedNew(dflSize uint64) {}

// Unsupported function: g_strjoin : unsupported parameter ... : varargs

// Strjoinv is a wrapper around the C function g_strjoinv.
func Strjoinv(separator string, strArray string) {}

// Strlcat is a wrapper around the C function g_strlcat.
func Strlcat(dest string, src string, destSize uint64) {}

// Strlcpy is a wrapper around the C function g_strlcpy.
func Strlcpy(dest string, src string, destSize uint64) {}

// Strncasecmp is a wrapper around the C function g_strncasecmp.
func Strncasecmp(s1 string, s2 string, n uint32) {}

// Strndup is a wrapper around the C function g_strndup.
func Strndup(str string, n uint64) {}

// Strnfill is a wrapper around the C function g_strnfill.
func Strnfill(length uint64, fillChar rune) {}

// Strreverse is a wrapper around the C function g_strreverse.
func Strreverse(string string) {}

// Strrstr is a wrapper around the C function g_strrstr.
func Strrstr(haystack string, needle string) {}

// StrrstrLen is a wrapper around the C function g_strrstr_len.
func StrrstrLen(haystack string, haystackLen int64, needle string) {}

// Strsignal is a wrapper around the C function g_strsignal.
func Strsignal(signum int32) {}

// Strsplit is a wrapper around the C function g_strsplit.
func Strsplit(string string, delimiter string, maxTokens int32) {}

// StrstrLen is a wrapper around the C function g_strstr_len.
func StrstrLen(haystack string, haystackLen int64, needle string) {}

// Strtod is a wrapper around the C function g_strtod.
func Strtod(nptr string, endptr string) {}

// Strup is a wrapper around the C function g_strup.
func Strup(string string) {}

// StrvGetType is a wrapper around the C function g_strv_get_type.
func StrvGetType() {}

// Unsupported function: g_test_add_vtable : unsupported parameter test_data : type gpointer, gconstpointer

// Unsupported function: g_test_assert_expected_messages_internal : unsupported parameter line : type gint, int

// Unsupported function: g_test_log_type_name : unsupported parameter log_type : type TestLogType, GTestLogType

// Unsupported function: g_test_trap_assertions : unsupported parameter line : type gint, int

// ThreadErrorQuark is a wrapper around the C function g_thread_error_quark.
func ThreadErrorQuark() {}

// ThreadExit is a wrapper around the C function g_thread_exit.
func ThreadExit(retval uintptr) {}

// ThreadPoolGetMaxUnusedThreads is a wrapper around the C function g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() {}

// ThreadPoolGetNumUnusedThreads is a wrapper around the C function g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() {}

// ThreadPoolSetMaxUnusedThreads is a wrapper around the C function g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {}

// ThreadPoolStopUnusedThreads is a wrapper around the C function g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {}

// ThreadSelf is a wrapper around the C function g_thread_self.
func ThreadSelf() {}

// ThreadYield is a wrapper around the C function g_thread_yield.
func ThreadYield() {}

// Unsupported function: g_timeout_add : unsupported parameter function : type SourceFunc, GSourceFunc

// Unsupported function: g_timeout_add_full : unsupported parameter function : type SourceFunc, GSourceFunc

// TimeoutSourceNew is a wrapper around the C function g_timeout_source_new.
func TimeoutSourceNew(interval uint32) {}

// Unsupported function: g_trash_stack_height : unsupported parameter stack_p : type TrashStack, GTrashStack**

// Unsupported function: g_trash_stack_peek : unsupported parameter stack_p : type TrashStack, GTrashStack**

// Unsupported function: g_trash_stack_pop : unsupported parameter stack_p : type TrashStack, GTrashStack**

// Unsupported function: g_trash_stack_push : unsupported parameter stack_p : type TrashStack, GTrashStack**

// TryMalloc is a wrapper around the C function g_try_malloc.
func TryMalloc(nBytes uint64) {}

// TryRealloc is a wrapper around the C function g_try_realloc.
func TryRealloc(mem uintptr, nBytes uint64) {}

// Unsupported function: g_ucs4_to_utf16 : unsupported parameter str : type gunichar, const gunichar*

// Unsupported function: g_ucs4_to_utf8 : unsupported parameter str : type gunichar, const gunichar*

// UnicharBreakType is a wrapper around the C function g_unichar_break_type.
func UnicharBreakType(c rune) {}

// UnicharDigitValue is a wrapper around the C function g_unichar_digit_value.
func UnicharDigitValue(c rune) {}

// UnicharIsalnum is a wrapper around the C function g_unichar_isalnum.
func UnicharIsalnum(c rune) {}

// UnicharIsalpha is a wrapper around the C function g_unichar_isalpha.
func UnicharIsalpha(c rune) {}

// UnicharIscntrl is a wrapper around the C function g_unichar_iscntrl.
func UnicharIscntrl(c rune) {}

// UnicharIsdefined is a wrapper around the C function g_unichar_isdefined.
func UnicharIsdefined(c rune) {}

// UnicharIsdigit is a wrapper around the C function g_unichar_isdigit.
func UnicharIsdigit(c rune) {}

// UnicharIsgraph is a wrapper around the C function g_unichar_isgraph.
func UnicharIsgraph(c rune) {}

// UnicharIslower is a wrapper around the C function g_unichar_islower.
func UnicharIslower(c rune) {}

// UnicharIsprint is a wrapper around the C function g_unichar_isprint.
func UnicharIsprint(c rune) {}

// UnicharIspunct is a wrapper around the C function g_unichar_ispunct.
func UnicharIspunct(c rune) {}

// UnicharIsspace is a wrapper around the C function g_unichar_isspace.
func UnicharIsspace(c rune) {}

// UnicharIstitle is a wrapper around the C function g_unichar_istitle.
func UnicharIstitle(c rune) {}

// UnicharIsupper is a wrapper around the C function g_unichar_isupper.
func UnicharIsupper(c rune) {}

// UnicharIswide is a wrapper around the C function g_unichar_iswide.
func UnicharIswide(c rune) {}

// UnicharIsxdigit is a wrapper around the C function g_unichar_isxdigit.
func UnicharIsxdigit(c rune) {}

// UnicharToUtf8 is a wrapper around the C function g_unichar_to_utf8.
func UnicharToUtf8(c rune, outbuf string) {}

// UnicharTolower is a wrapper around the C function g_unichar_tolower.
func UnicharTolower(c rune) {}

// UnicharTotitle is a wrapper around the C function g_unichar_totitle.
func UnicharTotitle(c rune) {}

// UnicharToupper is a wrapper around the C function g_unichar_toupper.
func UnicharToupper(c rune) {}

// UnicharType is a wrapper around the C function g_unichar_type.
func UnicharType(c rune) {}

// UnicharValidate is a wrapper around the C function g_unichar_validate.
func UnicharValidate(ch rune) {}

// UnicharXdigitValue is a wrapper around the C function g_unichar_xdigit_value.
func UnicharXdigitValue(c rune) {}

// Blacklisted function: g_unicode_canonical_decomposition

// Blacklisted function: g_unicode_canonical_ordering

// UnixErrorQuark is a wrapper around the C function g_unix_error_quark.
func UnixErrorQuark() {}

// Usleep is a wrapper around the C function g_usleep.
func Usleep(microseconds uint64) {}

// Unsupported function: g_utf16_to_ucs4 : unsupported parameter str : type guint16, const gunichar2*

// Unsupported function: g_utf16_to_utf8 : unsupported parameter str : type guint16, const gunichar2*

// Utf8Casefold is a wrapper around the C function g_utf8_casefold.
func Utf8Casefold(str string, len int64) {}

// Utf8Collate is a wrapper around the C function g_utf8_collate.
func Utf8Collate(str1 string, str2 string) {}

// Utf8CollateKey is a wrapper around the C function g_utf8_collate_key.
func Utf8CollateKey(str string, len int64) {}

// Utf8FindNextChar is a wrapper around the C function g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) {}

// Utf8FindPrevChar is a wrapper around the C function g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) {}

// Utf8GetChar is a wrapper around the C function g_utf8_get_char.
func Utf8GetChar(p string) {}

// Utf8GetCharValidated is a wrapper around the C function g_utf8_get_char_validated.
func Utf8GetCharValidated(p string, maxLen int64) {}

// Unsupported function: g_utf8_normalize : unsupported parameter mode : type NormalizeMode, GNormalizeMode

// Utf8OffsetToPointer is a wrapper around the C function g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) {}

// Utf8PointerToOffset is a wrapper around the C function g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) {}

// Utf8PrevChar is a wrapper around the C function g_utf8_prev_char.
func Utf8PrevChar(p string) {}

// Utf8Strchr is a wrapper around the C function g_utf8_strchr.
func Utf8Strchr(p string, len int64, c rune) {}

// Utf8Strdown is a wrapper around the C function g_utf8_strdown.
func Utf8Strdown(str string, len int64) {}

// Utf8Strlen is a wrapper around the C function g_utf8_strlen.
func Utf8Strlen(p string, max int64) {}

// Utf8Strncpy is a wrapper around the C function g_utf8_strncpy.
func Utf8Strncpy(dest string, src string, n uint64) {}

// Utf8Strrchr is a wrapper around the C function g_utf8_strrchr.
func Utf8Strrchr(p string, len int64, c rune) {}

// Utf8Strup is a wrapper around the C function g_utf8_strup.
func Utf8Strup(str string, len int64) {}

// Unsupported function: g_utf8_to_ucs4 : unsupported parameter items_read : type glong, glong*

// Blacklisted function: g_utf8_to_ucs4_fast

// Unsupported function: g_utf8_to_utf16 : unsupported parameter items_read : type glong, glong*

// Unsupported function: g_utf8_validate : unsupported parameter str : no type

// VariantGetGtype is a wrapper around the C function g_variant_get_gtype.
func VariantGetGtype() {}

// Unsupported function: g_variant_parse : unsupported parameter type : type VariantType, const GVariantType*

// VariantParseErrorQuark is a wrapper around the C function g_variant_parse_error_quark.
func VariantParseErrorQuark() {}

// VariantParserGetErrorQuark is a wrapper around the C function g_variant_parser_get_error_quark.
func VariantParserGetErrorQuark() {}

// Blacklisted function: g_variant_type_checked_

// VariantTypeStringIsValid is a wrapper around the C function g_variant_type_string_is_valid.
func VariantTypeStringIsValid(typeString string) {}

// Unsupported function: g_vsnprintf : unsupported parameter args : type va_list, va_list

// Unsupported function: g_warn_message : unsupported parameter line : type gint, int

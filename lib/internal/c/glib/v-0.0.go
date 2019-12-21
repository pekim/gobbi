// Code generated - DO NOT EDIT.
// +build !glib_2.4,!glib_2.6,!glib_2.14,!glib_2.16,!glib_2.22,!glib_2.24,!glib_2.26,!glib_2.30,!glib_2.32,!glib_2.38,!glib_2.40,!glib_2.50,!glib_2.54,!glib_2.60

package glib

import (
	c "github.com/pekim/gobbi/lib/internal/c"
	"unsafe"
)

// #include <glib.h>
import "C"

// bitfields
type AsciiType C.GAsciiType
type FileTest C.GFileTest
type FormatSizeFlags C.GFormatSizeFlags
type HookFlagMask C.GHookFlagMask
type IOCondition C.GIOCondition
type IOFlags C.GIOFlags
type KeyFileFlags C.GKeyFileFlags
type LogLevelFlags C.GLogLevelFlags
type MarkupCollectType C.GMarkupCollectType
type MarkupParseFlags C.GMarkupParseFlags
type OptionFlags C.GOptionFlags
type SpawnFlags C.GSpawnFlags
type TestSubprocessFlags C.GTestSubprocessFlags
type TestTrapFlags C.GTestTrapFlags
type TraverseFlags C.GTraverseFlags

// enumerations
type BookmarkFileError C.GBookmarkFileError
type ConvertError C.GConvertError
type DateDMY C.GDateDMY
type DateMonth C.GDateMonth
type DateWeekday C.GDateWeekday
type ErrorType C.GErrorType
type FileError C.GFileError
type IOChannelError C.GIOChannelError
type IOError C.GIOError
type IOStatus C.GIOStatus
type KeyFileError C.GKeyFileError
type MarkupError C.GMarkupError
type NormalizeMode C.GNormalizeMode
type OptionArg C.GOptionArg
type OptionError C.GOptionError
type SeekType C.GSeekType
type ShellError C.GShellError
type SliceConfig C.GSliceConfig
type SpawnError C.GSpawnError
type TestLogType C.GTestLogType

// UNSUPPORTED : TestResult : blacklisted
type ThreadError C.GThreadError
type TimeType C.GTimeType
type TokenType C.GTokenType
type TraverseType C.GTraverseType
type UnicodeBreakType C.GUnicodeBreakType
type UnicodeScript C.GUnicodeScript
type UnicodeType C.GUnicodeType
type VariantParseError C.GVariantParseError

// records
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

// classes

// interfaces

func Fn_ascii_digit_value(c c.UndefinedParamType) {}

func Fn_ascii_dtostr(buffer c.UndefinedParamType, bufLen int, d float64) {}

func Fn_ascii_formatd(buffer c.UndefinedParamType, bufLen int, format c.UndefinedParamType, d float64) {
}

func Fn_ascii_strcasecmp(s1 c.UndefinedParamType, s2 c.UndefinedParamType) {}

func Fn_ascii_strdown(str c.UndefinedParamType, len uint64) {}

func Fn_ascii_strncasecmp(s1 c.UndefinedParamType, s2 c.UndefinedParamType, n uint64) {}

func Fn_ascii_strtod(nptr c.UndefinedParamType) {}

func Fn_ascii_strup(str c.UndefinedParamType, len uint64) {}

func Fn_ascii_tolower(c c.UndefinedParamType) {}

func Fn_ascii_toupper(c c.UndefinedParamType) {}

func Fn_ascii_xdigit_value(c c.UndefinedParamType) {}

func Fn_assert_warning(logDomain c.UndefinedParamType, file c.UndefinedParamType, line int, prettyFunction c.UndefinedParamType, expression c.UndefinedParamType) {
}

func Fn_assertion_message(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType, message c.UndefinedParamType) {
}

func Fn_assertion_message_cmpnum(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType, expr c.UndefinedParamType, arg1 float64, cmp c.UndefinedParamType, arg2 float64, numtype c.UndefinedParamType) {
}

func Fn_assertion_message_cmpstr(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType, expr c.UndefinedParamType, arg1 c.UndefinedParamType, cmp c.UndefinedParamType, arg2 c.UndefinedParamType) {
}

func Fn_assertion_message_error(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType, expr c.UndefinedParamType, error c.UndefinedParamType, errorDomain c.UndefinedParamType, errorCode int) {
}

func Fn_assertion_message_expr(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType, expr c.UndefinedParamType) {
}

func Fn_atexit(func_ c.UndefinedParamType) {}

func Fn_basename(fileName c.UndefinedParamType) {}

func Fn_bit_nth_lsf(mask uint64, nthBit int) {}

func Fn_bit_nth_msf(mask uint64, nthBit int) {}

func Fn_bit_storage(number uint64) {}

func Fn_bookmark_file_error_quark() {}

// UNSUPPORTED : build_filename : has varargs

// UNSUPPORTED : build_path : has varargs

func Fn_byte_array_free(array c.UndefinedParamType, freeSegment bool) {}

func Fn_byte_array_new() {}

func Fn_clear_error() {}

func Fn_convert(str c.UndefinedParamType, len uint64, toCodeset c.UndefinedParamType, fromCodeset c.UndefinedParamType) {
}

func Fn_convert_error_quark() {}

func Fn_convert_with_fallback(str c.UndefinedParamType, len uint64, toCodeset c.UndefinedParamType, fromCodeset c.UndefinedParamType, fallback c.UndefinedParamType) {
}

func Fn_convert_with_iconv(str c.UndefinedParamType, len uint64, converter c.UndefinedParamType) {}

func Fn_datalist_clear(datalist c.UndefinedParamType) {}

func Fn_datalist_foreach(datalist c.UndefinedParamType, func_ c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_datalist_get_data(datalist c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_datalist_id_get_data(datalist c.UndefinedParamType, keyId c.UndefinedParamType) {}

func Fn_datalist_id_remove_no_notify(datalist c.UndefinedParamType, keyId c.UndefinedParamType) {}

func Fn_datalist_id_set_data_full(datalist c.UndefinedParamType, keyId c.UndefinedParamType, data unsafe.Pointer, destroyFunc c.UndefinedParamType) {
}

func Fn_datalist_init(datalist c.UndefinedParamType) {}

func Fn_dataset_destroy(datasetLocation c.UndefinedParamType) {}

func Fn_dataset_foreach(datasetLocation c.UndefinedParamType, func_ c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_dataset_id_get_data(datasetLocation c.UndefinedParamType, keyId c.UndefinedParamType) {}

func Fn_dataset_id_remove_no_notify(datasetLocation c.UndefinedParamType, keyId c.UndefinedParamType) {
}

func Fn_dataset_id_set_data_full(datasetLocation c.UndefinedParamType, keyId c.UndefinedParamType, data unsafe.Pointer, destroyFunc c.UndefinedParamType) {
}

func Fn_date_get_days_in_month(month c.UndefinedParamType, year c.UndefinedParamType) {}

func Fn_date_get_monday_weeks_in_year(year c.UndefinedParamType) {}

func Fn_date_get_sunday_weeks_in_year(year c.UndefinedParamType) {}

func Fn_date_is_leap_year(year c.UndefinedParamType) {}

func Fn_date_strftime(s c.UndefinedParamType, slen uint64, format c.UndefinedParamType, date c.UndefinedParamType) {
}

func Fn_date_valid_day(day c.UndefinedParamType) {}

func Fn_date_valid_dmy(day c.UndefinedParamType, month c.UndefinedParamType, year c.UndefinedParamType) {
}

func Fn_date_valid_julian(julianDate uint32) {}

func Fn_date_valid_month(month c.UndefinedParamType) {}

func Fn_date_valid_weekday(weekday c.UndefinedParamType) {}

func Fn_date_valid_year(year c.UndefinedParamType) {}

func Fn_direct_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_direct_hash(v c.UndefinedParamType) {}

func Fn_file_error_from_errno(errNo int) {}

func Fn_file_error_quark() {}

func Fn_file_get_contents(filename c.UndefinedParamType) {}

func Fn_file_open_tmp(tmpl c.UndefinedParamType) {}

func Fn_file_test(filename c.UndefinedParamType, test c.UndefinedParamType) {}

func Fn_filename_from_uri(uri c.UndefinedParamType) {}

func Fn_filename_from_utf8(utf8string c.UndefinedParamType, len uint64) {}

func Fn_filename_to_uri(filename c.UndefinedParamType, hostname c.UndefinedParamType) {}

func Fn_filename_to_utf8(opsysstring c.UndefinedParamType, len uint64) {}

func Fn_find_program_in_path(program c.UndefinedParamType) {}

// UNSUPPORTED : fprintf : has varargs

func Fn_free(mem unsafe.Pointer) {}

func Fn_get_charset() {}

func Fn_get_codeset() {}

func Fn_get_current_dir() {}

func Fn_get_current_time(result c.UndefinedParamType) {}

func Fn_get_home_dir() {}

func Fn_get_prgname() {}

func Fn_get_real_name() {}

func Fn_get_tmp_dir() {}

func Fn_get_user_name() {}

func Fn_getenv(variable c.UndefinedParamType) {}

func Fn_hash_table_destroy(hashTable c.UndefinedParamType) {}

func Fn_hash_table_insert(hashTable c.UndefinedParamType, key unsafe.Pointer, value unsafe.Pointer) {}

func Fn_hash_table_lookup(hashTable c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_hash_table_lookup_extended(hashTable c.UndefinedParamType, lookupKey c.UndefinedParamType) {}

func Fn_hash_table_remove(hashTable c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_hash_table_replace(hashTable c.UndefinedParamType, key unsafe.Pointer, value unsafe.Pointer) {
}

func Fn_hash_table_size(hashTable c.UndefinedParamType) {}

func Fn_hash_table_steal(hashTable c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_hook_destroy(hookList c.UndefinedParamType, hookId uint64) {}

func Fn_hook_destroy_link(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_hook_free(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_hook_insert_before(hookList c.UndefinedParamType, sibling c.UndefinedParamType, hook c.UndefinedParamType) {
}

func Fn_hook_prepend(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_hook_unref(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_iconv(converter c.UndefinedParamType, inbuf c.UndefinedParamType, inbytesLeft *uint64, outbuf c.UndefinedParamType, outbytesLeft *uint64) {
}

func Fn_iconv_open(toCodeset c.UndefinedParamType, fromCodeset c.UndefinedParamType) {}

func Fn_idle_add(function c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_idle_add_full(priority int, function c.UndefinedParamType, data unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_idle_remove_by_data(data unsafe.Pointer) {}

func Fn_idle_source_new() {}

func Fn_int_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_int_hash(v c.UndefinedParamType) {}

func Fn_io_add_watch(channel c.UndefinedParamType, condition c.UndefinedParamType, func_ c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_io_add_watch_full(channel c.UndefinedParamType, priority int, condition c.UndefinedParamType, func_ c.UndefinedParamType, userData unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_io_channel_error_from_errno(en int) {}

func Fn_io_channel_error_quark() {}

func Fn_io_create_watch(channel c.UndefinedParamType, condition c.UndefinedParamType) {}

func Fn_key_file_error_quark() {}

func Fn_locale_from_utf8(utf8string c.UndefinedParamType, len uint64) {}

func Fn_locale_to_utf8(opsysstring c.UndefinedParamType, len uint64) {}

// UNSUPPORTED : log : has varargs

func Fn_log_default_handler(logDomain c.UndefinedParamType, logLevel c.UndefinedParamType, message c.UndefinedParamType, unusedData unsafe.Pointer) {
}

func Fn_log_remove_handler(logDomain c.UndefinedParamType, handlerId uint) {}

func Fn_log_set_always_fatal(fatalMask c.UndefinedParamType) {}

func Fn_log_set_fatal_mask(logDomain c.UndefinedParamType, fatalMask c.UndefinedParamType) {}

func Fn_log_set_handler(logDomain c.UndefinedParamType, logLevels c.UndefinedParamType, logFunc c.UndefinedParamType, userData unsafe.Pointer) {
}

// UNSUPPORTED : log_structured : has varargs

// UNSUPPORTED : log_structured_standard : has varargs

func Fn_logv(logDomain c.UndefinedParamType, logLevel c.UndefinedParamType, format c.UndefinedParamType, args c.UndefinedParamType) {
}

func Fn_main_context_default() {}

func Fn_main_depth() {}

func Fn_malloc(nBytes uint64) {}

func Fn_malloc0(nBytes uint64) {}

// UNSUPPORTED : markup_collect_attributes : has varargs

func Fn_markup_error_quark() {}

func Fn_markup_escape_text(text c.UndefinedParamType, length uint64) {}

// UNSUPPORTED : markup_printf_escaped : has varargs

func Fn_mem_is_system_malloc() {}

func Fn_mem_profile() {}

func Fn_mem_set_vtable(vtable c.UndefinedParamType) {}

func Fn_memdup(mem c.UndefinedParamType, byteSize uint) {}

func Fn_mkstemp(tmpl c.UndefinedParamType) {}

func Fn_nullify_pointer(nullifyLocation *unsafe.Pointer) {}

func Fn_number_parser_error_quark() {}

func Fn_on_error_query(prgName c.UndefinedParamType) {}

func Fn_on_error_stack_trace(prgName c.UndefinedParamType) {}

func Fn_option_error_quark() {}

func Fn_parse_debug_string(string_ c.UndefinedParamType, keys c.UndefinedParamType, nkeys uint) {}

func Fn_path_get_basename(fileName c.UndefinedParamType) {}

func Fn_path_get_dirname(fileName c.UndefinedParamType) {}

func Fn_path_is_absolute(fileName c.UndefinedParamType) {}

func Fn_path_skip_root(fileName c.UndefinedParamType) {}

func Fn_pattern_match(pspec c.UndefinedParamType, stringLength uint, string_ c.UndefinedParamType, stringReversed c.UndefinedParamType) {
}

func Fn_pattern_match_simple(pattern c.UndefinedParamType, string_ c.UndefinedParamType) {}

func Fn_pattern_match_string(pspec c.UndefinedParamType, string_ c.UndefinedParamType) {}

// UNSUPPORTED : prefix_error : has varargs

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

// UNSUPPORTED : printf : has varargs

func Fn_printf_string_upper_bound(format c.UndefinedParamType, args c.UndefinedParamType) {}

func Fn_propagate_error(src c.UndefinedParamType) {}

// UNSUPPORTED : propagate_prefixed_error : has varargs

func Fn_qsort_with_data(pbase c.UndefinedParamType, totalElems int, size uint64, compareFunc c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_quark_from_static_string(string_ c.UndefinedParamType) {}

func Fn_quark_from_string(string_ c.UndefinedParamType) {}

func Fn_quark_to_string(quark c.UndefinedParamType) {}

func Fn_quark_try_string(string_ c.UndefinedParamType) {}

func Fn_random_double() {}

func Fn_random_double_range(begin float64, end float64) {}

func Fn_random_int() {}

func Fn_random_int_range(begin int32, end int32) {}

func Fn_random_set_seed(seed uint32) {}

func Fn_realloc(mem unsafe.Pointer, nBytes uint64) {}

func Fn_regex_error_quark() {}

func Fn_return_if_fail_warning(logDomain c.UndefinedParamType, prettyFunction c.UndefinedParamType, expression c.UndefinedParamType) {
}

// UNSUPPORTED : set_error : has varargs

func Fn_set_prgname(prgname c.UndefinedParamType) {}

func Fn_set_print_handler(func_ c.UndefinedParamType) {}

func Fn_set_printerr_handler(func_ c.UndefinedParamType) {}

func Fn_shell_error_quark() {}

func Fn_shell_parse_argv(commandLine c.UndefinedParamType) {}

func Fn_shell_quote(unquotedString c.UndefinedParamType) {}

func Fn_shell_unquote(quotedString c.UndefinedParamType) {}

func Fn_slice_get_config(ckey c.UndefinedParamType) {}

func Fn_slice_get_config_state(ckey c.UndefinedParamType, address int64, nValues *uint) {}

func Fn_slice_set_config(ckey c.UndefinedParamType, value int64) {}

// UNSUPPORTED : snprintf : has varargs

func Fn_source_remove(tag uint) {}

func Fn_source_remove_by_funcs_user_data(funcs c.UndefinedParamType, userData unsafe.Pointer) {}

func Fn_source_remove_by_user_data(userData unsafe.Pointer) {}

func Fn_spaced_primes_closest(num uint) {}

func Fn_spawn_async(workingDirectory c.UndefinedParamType, argv c.UndefinedParamType, envp c.UndefinedParamType, flags c.UndefinedParamType, childSetup c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_spawn_async_with_pipes(workingDirectory c.UndefinedParamType, argv c.UndefinedParamType, envp c.UndefinedParamType, flags c.UndefinedParamType, childSetup c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_spawn_close_pid(pid c.UndefinedParamType) {}

func Fn_spawn_command_line_async(commandLine c.UndefinedParamType) {}

func Fn_spawn_command_line_sync(commandLine c.UndefinedParamType) {}

func Fn_spawn_error_quark() {}

func Fn_spawn_exit_error_quark() {}

func Fn_spawn_sync(workingDirectory c.UndefinedParamType, argv c.UndefinedParamType, envp c.UndefinedParamType, flags c.UndefinedParamType, childSetup c.UndefinedParamType, userData unsafe.Pointer) {
}

// UNSUPPORTED : sprintf : has varargs

func Fn_stpcpy(dest c.UndefinedParamType, src c.UndefinedParamType) {}

func Fn_str_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_str_hash(v c.UndefinedParamType) {}

func Fn_strcanon(string_ c.UndefinedParamType, validChars c.UndefinedParamType, substitutor c.UndefinedParamType) {
}

func Fn_strcasecmp(s1 c.UndefinedParamType, s2 c.UndefinedParamType) {}

func Fn_strchomp(string_ c.UndefinedParamType) {}

func Fn_strchug(string_ c.UndefinedParamType) {}

func Fn_strcompress(source c.UndefinedParamType) {}

// UNSUPPORTED : strconcat : has varargs

func Fn_strdelimit(string_ c.UndefinedParamType, delimiters c.UndefinedParamType, newDelimiter c.UndefinedParamType) {
}

func Fn_strdown(string_ c.UndefinedParamType) {}

func Fn_strdup(str c.UndefinedParamType) {}

// UNSUPPORTED : strdup_printf : has varargs

func Fn_strdup_vprintf(format c.UndefinedParamType, args c.UndefinedParamType) {}

func Fn_strdupv(strArray c.UndefinedParamType) {}

func Fn_strerror(errnum int) {}

func Fn_strescape(source c.UndefinedParamType, exceptions c.UndefinedParamType) {}

func Fn_strfreev(strArray c.UndefinedParamType) {}

func Fn_string_new(init c.UndefinedParamType) {}

func Fn_string_new_len(init c.UndefinedParamType, len uint64) {}

func Fn_string_sized_new(dflSize uint64) {}

// UNSUPPORTED : strjoin : has varargs

func Fn_strjoinv(separator c.UndefinedParamType, strArray c.UndefinedParamType) {}

func Fn_strlcat(dest c.UndefinedParamType, src c.UndefinedParamType, destSize uint64) {}

func Fn_strlcpy(dest c.UndefinedParamType, src c.UndefinedParamType, destSize uint64) {}

func Fn_strncasecmp(s1 c.UndefinedParamType, s2 c.UndefinedParamType, n uint) {}

func Fn_strndup(str c.UndefinedParamType, n uint64) {}

func Fn_strnfill(length uint64, fillChar c.UndefinedParamType) {}

func Fn_strreverse(string_ c.UndefinedParamType) {}

func Fn_strrstr(haystack c.UndefinedParamType, needle c.UndefinedParamType) {}

func Fn_strrstr_len(haystack c.UndefinedParamType, haystackLen uint64, needle c.UndefinedParamType) {}

func Fn_strsignal(signum int) {}

func Fn_strsplit(string_ c.UndefinedParamType, delimiter c.UndefinedParamType, maxTokens int) {}

func Fn_strstr_len(haystack c.UndefinedParamType, haystackLen uint64, needle c.UndefinedParamType) {}

func Fn_strtod(nptr c.UndefinedParamType) {}

func Fn_strup(string_ c.UndefinedParamType) {}

func Fn_strv_get_type() {}

func Fn_test_add_vtable(testpath c.UndefinedParamType, dataSize uint64, testData c.UndefinedParamType, dataSetup c.UndefinedParamType, dataTest c.UndefinedParamType, dataTeardown c.UndefinedParamType) {
}

func Fn_test_assert_expected_messages_internal(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType) {
}

// UNSUPPORTED : test_build_filename : has varargs

// UNSUPPORTED : test_get_filename : has varargs

// UNSUPPORTED : test_init : has varargs

func Fn_test_log_type_name(logType c.UndefinedParamType) {}

// UNSUPPORTED : test_maximized_result : has varargs

// UNSUPPORTED : test_message : has varargs

// UNSUPPORTED : test_minimized_result : has varargs

func Fn_test_trap_assertions(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType, assertionFlags uint64, pattern c.UndefinedParamType) {
}

func Fn_thread_error_quark() {}

func Fn_thread_exit(retval unsafe.Pointer) {}

func Fn_thread_pool_get_max_unused_threads() {}

func Fn_thread_pool_get_num_unused_threads() {}

func Fn_thread_pool_set_max_unused_threads(maxThreads int) {}

func Fn_thread_pool_stop_unused_threads() {}

func Fn_thread_self() {}

func Fn_thread_yield() {}

func Fn_timeout_add(interval uint, function c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_timeout_add_full(priority int, interval uint, function c.UndefinedParamType, data unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_timeout_source_new(interval uint) {}

func Fn_trash_stack_height(stackP c.UndefinedParamType) {}

func Fn_trash_stack_peek(stackP c.UndefinedParamType) {}

func Fn_trash_stack_pop(stackP c.UndefinedParamType) {}

func Fn_trash_stack_push(stackP c.UndefinedParamType, dataP unsafe.Pointer) {}

func Fn_try_malloc(nBytes uint64) {}

func Fn_try_realloc(mem unsafe.Pointer, nBytes uint64) {}

func Fn_ucs4_to_utf16(str *rune, len int64) {}

func Fn_ucs4_to_utf8(str *rune, len int64) {}

func Fn_unichar_break_type(c rune) {}

func Fn_unichar_digit_value(c rune) {}

func Fn_unichar_isalnum(c rune) {}

func Fn_unichar_isalpha(c rune) {}

func Fn_unichar_iscntrl(c rune) {}

func Fn_unichar_isdefined(c rune) {}

func Fn_unichar_isdigit(c rune) {}

func Fn_unichar_isgraph(c rune) {}

func Fn_unichar_islower(c rune) {}

func Fn_unichar_isprint(c rune) {}

func Fn_unichar_ispunct(c rune) {}

func Fn_unichar_isspace(c rune) {}

func Fn_unichar_istitle(c rune) {}

func Fn_unichar_isupper(c rune) {}

func Fn_unichar_iswide(c rune) {}

func Fn_unichar_isxdigit(c rune) {}

func Fn_unichar_to_utf8(c rune) {}

func Fn_unichar_tolower(c rune) {}

func Fn_unichar_totitle(c rune) {}

func Fn_unichar_toupper(c rune) {}

func Fn_unichar_type(c rune) {}

func Fn_unichar_validate(ch rune) {}

func Fn_unichar_xdigit_value(c rune) {}

func Fn_unicode_canonical_decomposition(ch rune, resultLen *uint64) {}

func Fn_unicode_canonical_ordering(string_ *rune, len uint64) {}

func Fn_unix_error_quark() {}

func Fn_usleep(microseconds uint64) {}

func Fn_utf16_to_ucs4(str c.UndefinedParamType, len int64) {}

func Fn_utf16_to_utf8(str c.UndefinedParamType, len int64) {}

func Fn_utf8_casefold(str c.UndefinedParamType, len uint64) {}

func Fn_utf8_collate(str1 c.UndefinedParamType, str2 c.UndefinedParamType) {}

func Fn_utf8_collate_key(str c.UndefinedParamType, len uint64) {}

func Fn_utf8_find_next_char(p c.UndefinedParamType, end c.UndefinedParamType) {}

func Fn_utf8_find_prev_char(str c.UndefinedParamType, p c.UndefinedParamType) {}

func Fn_utf8_get_char(p c.UndefinedParamType) {}

func Fn_utf8_get_char_validated(p c.UndefinedParamType, maxLen uint64) {}

func Fn_utf8_normalize(str c.UndefinedParamType, len uint64, mode c.UndefinedParamType) {}

func Fn_utf8_offset_to_pointer(str c.UndefinedParamType, offset int64) {}

func Fn_utf8_pointer_to_offset(str c.UndefinedParamType, pos c.UndefinedParamType) {}

func Fn_utf8_prev_char(p c.UndefinedParamType) {}

func Fn_utf8_strchr(p c.UndefinedParamType, len uint64, c rune) {}

func Fn_utf8_strdown(str c.UndefinedParamType, len uint64) {}

func Fn_utf8_strlen(p c.UndefinedParamType, max uint64) {}

func Fn_utf8_strncpy(dest c.UndefinedParamType, src c.UndefinedParamType, n uint64) {}

func Fn_utf8_strrchr(p c.UndefinedParamType, len uint64, c rune) {}

func Fn_utf8_strup(str c.UndefinedParamType, len uint64) {}

func Fn_utf8_to_ucs4(str c.UndefinedParamType, len int64) {}

func Fn_utf8_to_ucs4_fast(str c.UndefinedParamType, len int64) {}

func Fn_utf8_to_utf16(str c.UndefinedParamType, len int64) {}

func Fn_utf8_validate(str c.UndefinedParamType, maxLen uint64) {}

func Fn_variant_get_gtype() {}

func Fn_variant_parse(type_ c.UndefinedParamType, text c.UndefinedParamType, limit c.UndefinedParamType, endptr c.UndefinedParamType) {
}

func Fn_variant_parse_error_quark() {}

func Fn_variant_parser_get_error_quark() {}

func Fn_variant_type_checked_(arg0 c.UndefinedParamType) {}

func Fn_variant_type_string_get_depth_(typeString c.UndefinedParamType) {}

func Fn_variant_type_string_is_valid(typeString c.UndefinedParamType) {}

func Fn_vsnprintf(string_ c.UndefinedParamType, n uint64, format c.UndefinedParamType, args c.UndefinedParamType) {
}

func Fn_warn_message(domain c.UndefinedParamType, file c.UndefinedParamType, line int, func_ c.UndefinedParamType, warnexpr c.UndefinedParamType) {
}

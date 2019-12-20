// Code generated - DO NOT EDIT.
// +build glib_2.4

package glib

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
type OnceStatus C.GOnceStatus
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

func Fn_ascii_digit_value(c string) {}

func Fn_ascii_dtostr(buffer string, bufLen string, d string) {}

func Fn_ascii_formatd(buffer string, bufLen string, format string, d string) {}

func Fn_ascii_strcasecmp(s1 string, s2 string) {}

func Fn_ascii_strdown(str string, len string) {}

func Fn_ascii_strncasecmp(s1 string, s2 string, n string) {}

func Fn_ascii_strtod(nptr string, endptr string) {}

func Fn_ascii_strtoull(nptr string, endptr string, base string) {}

func Fn_ascii_strup(str string, len string) {}

func Fn_ascii_tolower(c string) {}

func Fn_ascii_toupper(c string) {}

func Fn_ascii_xdigit_value(c string) {}

func Fn_assert_warning(logDomain string, file string, line string, prettyFunction string, expression string) {
}

func Fn_assertion_message(domain string, file string, line string, func_ string, message string) {}

func Fn_assertion_message_cmpnum(domain string, file string, line string, func_ string, expr string, arg1 string, cmp string, arg2 string, numtype string) {
}

func Fn_assertion_message_cmpstr(domain string, file string, line string, func_ string, expr string, arg1 string, cmp string, arg2 string) {
}

func Fn_assertion_message_error(domain string, file string, line string, func_ string, expr string, error string, errorDomain string, errorCode string) {
}

func Fn_assertion_message_expr(domain string, file string, line string, func_ string, expr string) {}

func Fn_atexit(func_ string) {}

func Fn_atomic_int_add(atomic string, val string) {}

func Fn_atomic_int_compare_and_exchange(atomic string, oldval string, newval string) {}

func Fn_atomic_int_dec_and_test(atomic string) {}

func Fn_atomic_int_exchange_and_add(atomic string, val string) {}

func Fn_atomic_int_get(atomic string) {}

func Fn_atomic_int_inc(atomic string) {}

func Fn_atomic_int_set(atomic string, newval string) {}

func Fn_atomic_pointer_compare_and_exchange(atomic string, oldval string, newval string) {}

func Fn_atomic_pointer_get(atomic string) {}

func Fn_atomic_pointer_set(atomic string, newval string) {}

func Fn_basename(fileName string) {}

func Fn_bit_nth_lsf(mask string, nthBit string) {}

func Fn_bit_nth_msf(mask string, nthBit string) {}

func Fn_bit_storage(number string) {}

func Fn_bookmark_file_error_quark() {}

// UNSUPPORTED : build_filename : has varargs

// UNSUPPORTED : build_path : has varargs

func Fn_byte_array_free(array string, freeSegment string) {}

func Fn_byte_array_new() {}

func Fn_child_watch_add(pid string, function string, data string) {}

func Fn_child_watch_add_full(priority string, pid string, function string, data string, notify string) {}

func Fn_child_watch_source_new(pid string) {}

func Fn_clear_error() {}

func Fn_convert(str string, len string, toCodeset string, fromCodeset string, bytesRead string, bytesWritten string) {
}

func Fn_convert_error_quark() {}

func Fn_convert_with_fallback(str string, len string, toCodeset string, fromCodeset string, fallback string, bytesRead string, bytesWritten string) {
}

func Fn_convert_with_iconv(str string, len string, converter string, bytesRead string, bytesWritten string) {
}

func Fn_datalist_clear(datalist string) {}

func Fn_datalist_foreach(datalist string, func_ string, userData string) {}

func Fn_datalist_get_data(datalist string, key string) {}

func Fn_datalist_id_get_data(datalist string, keyId string) {}

func Fn_datalist_id_remove_no_notify(datalist string, keyId string) {}

func Fn_datalist_id_set_data_full(datalist string, keyId string, data string, destroyFunc string) {}

func Fn_datalist_init(datalist string) {}

func Fn_dataset_destroy(datasetLocation string) {}

func Fn_dataset_foreach(datasetLocation string, func_ string, userData string) {}

func Fn_dataset_id_get_data(datasetLocation string, keyId string) {}

func Fn_dataset_id_remove_no_notify(datasetLocation string, keyId string) {}

func Fn_dataset_id_set_data_full(datasetLocation string, keyId string, data string, destroyFunc string) {
}

func Fn_date_get_days_in_month(month string, year string) {}

func Fn_date_get_monday_weeks_in_year(year string) {}

func Fn_date_get_sunday_weeks_in_year(year string) {}

func Fn_date_is_leap_year(year string) {}

func Fn_date_strftime(s string, slen string, format string, date string) {}

func Fn_date_valid_day(day string) {}

func Fn_date_valid_dmy(day string, month string, year string) {}

func Fn_date_valid_julian(julianDate string) {}

func Fn_date_valid_month(month string) {}

func Fn_date_valid_weekday(weekday string) {}

func Fn_date_valid_year(year string) {}

func Fn_direct_equal(v1 string, v2 string) {}

func Fn_direct_hash(v string) {}

func Fn_file_error_from_errno(errNo string) {}

func Fn_file_error_quark() {}

func Fn_file_get_contents(filename string, contents string, length string) {}

func Fn_file_open_tmp(tmpl string, nameUsed string) {}

func Fn_file_read_link(filename string) {}

func Fn_file_test(filename string, test string) {}

func Fn_filename_from_uri(uri string, hostname string) {}

func Fn_filename_from_utf8(utf8string string, len string, bytesRead string, bytesWritten string) {}

func Fn_filename_to_uri(filename string, hostname string) {}

func Fn_filename_to_utf8(opsysstring string, len string, bytesRead string, bytesWritten string) {}

func Fn_find_program_in_path(program string) {}

// UNSUPPORTED : fprintf : has varargs

func Fn_free(mem string) {}

func Fn_get_application_name() {}

func Fn_get_charset(charset string) {}

func Fn_get_codeset() {}

func Fn_get_current_dir() {}

func Fn_get_current_time(result string) {}

func Fn_get_home_dir() {}

func Fn_get_prgname() {}

func Fn_get_real_name() {}

func Fn_get_tmp_dir() {}

func Fn_get_user_name() {}

func Fn_getenv(variable string) {}

func Fn_hash_table_destroy(hashTable string) {}

func Fn_hash_table_insert(hashTable string, key string, value string) {}

func Fn_hash_table_lookup(hashTable string, key string) {}

func Fn_hash_table_lookup_extended(hashTable string, lookupKey string, origKey string, value string) {}

func Fn_hash_table_remove(hashTable string, key string) {}

func Fn_hash_table_replace(hashTable string, key string, value string) {}

func Fn_hash_table_size(hashTable string) {}

func Fn_hash_table_steal(hashTable string, key string) {}

func Fn_hook_destroy(hookList string, hookId string) {}

func Fn_hook_destroy_link(hookList string, hook string) {}

func Fn_hook_free(hookList string, hook string) {}

func Fn_hook_insert_before(hookList string, sibling string, hook string) {}

func Fn_hook_prepend(hookList string, hook string) {}

func Fn_hook_unref(hookList string, hook string) {}

func Fn_iconv(converter string, inbuf string, inbytesLeft string, outbuf string, outbytesLeft string) {}

func Fn_iconv_open(toCodeset string, fromCodeset string) {}

func Fn_idle_add(function string, data string) {}

func Fn_idle_add_full(priority string, function string, data string, notify string) {}

func Fn_idle_remove_by_data(data string) {}

func Fn_idle_source_new() {}

func Fn_int_equal(v1 string, v2 string) {}

func Fn_int_hash(v string) {}

func Fn_io_add_watch(channel string, condition string, func_ string, userData string) {}

func Fn_io_add_watch_full(channel string, priority string, condition string, func_ string, userData string, notify string) {
}

func Fn_io_channel_error_from_errno(en string) {}

func Fn_io_channel_error_quark() {}

func Fn_io_create_watch(channel string, condition string) {}

func Fn_key_file_error_quark() {}

func Fn_locale_from_utf8(utf8string string, len string, bytesRead string, bytesWritten string) {}

func Fn_locale_to_utf8(opsysstring string, len string, bytesRead string, bytesWritten string) {}

// UNSUPPORTED : log : has varargs

func Fn_log_default_handler(logDomain string, logLevel string, message string, unusedData string) {}

func Fn_log_remove_handler(logDomain string, handlerId string) {}

func Fn_log_set_always_fatal(fatalMask string) {}

func Fn_log_set_fatal_mask(logDomain string, fatalMask string) {}

func Fn_log_set_handler(logDomain string, logLevels string, logFunc string, userData string) {}

// UNSUPPORTED : log_structured : has varargs

// UNSUPPORTED : log_structured_standard : has varargs

func Fn_logv(logDomain string, logLevel string, format string, args string) {}

func Fn_main_context_default() {}

func Fn_main_depth() {}

func Fn_malloc(nBytes string) {}

func Fn_malloc0(nBytes string) {}

// UNSUPPORTED : markup_collect_attributes : has varargs

func Fn_markup_error_quark() {}

func Fn_markup_escape_text(text string, length string) {}

// UNSUPPORTED : markup_printf_escaped : has varargs

func Fn_markup_vprintf_escaped(format string, args string) {}

func Fn_mem_is_system_malloc() {}

func Fn_mem_profile() {}

func Fn_mem_set_vtable(vtable string) {}

func Fn_memdup(mem string, byteSize string) {}

func Fn_mkstemp(tmpl string) {}

func Fn_nullify_pointer(nullifyLocation string) {}

func Fn_number_parser_error_quark() {}

func Fn_on_error_query(prgName string) {}

func Fn_on_error_stack_trace(prgName string) {}

func Fn_option_error_quark() {}

func Fn_parse_debug_string(string_ string, keys string, nkeys string) {}

func Fn_path_get_basename(fileName string) {}

func Fn_path_get_dirname(fileName string) {}

func Fn_path_is_absolute(fileName string) {}

func Fn_path_skip_root(fileName string) {}

func Fn_pattern_match(pspec string, stringLength string, string_ string, stringReversed string) {}

func Fn_pattern_match_simple(pattern string, string_ string) {}

func Fn_pattern_match_string(pspec string, string_ string) {}

// UNSUPPORTED : prefix_error : has varargs

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

// UNSUPPORTED : printf : has varargs

func Fn_printf_string_upper_bound(format string, args string) {}

func Fn_propagate_error(dest string, src string) {}

// UNSUPPORTED : propagate_prefixed_error : has varargs

func Fn_qsort_with_data(pbase string, totalElems string, size string, compareFunc string, userData string) {
}

func Fn_quark_from_static_string(string_ string) {}

func Fn_quark_from_string(string_ string) {}

func Fn_quark_to_string(quark string) {}

func Fn_quark_try_string(string_ string) {}

func Fn_random_double() {}

func Fn_random_double_range(begin string, end string) {}

func Fn_random_int() {}

func Fn_random_int_range(begin string, end string) {}

func Fn_random_set_seed(seed string) {}

func Fn_realloc(mem string, nBytes string) {}

func Fn_regex_error_quark() {}

func Fn_return_if_fail_warning(logDomain string, prettyFunction string, expression string) {}

func Fn_set_application_name(applicationName string) {}

// UNSUPPORTED : set_error : has varargs

func Fn_set_prgname(prgname string) {}

func Fn_set_print_handler(func_ string) {}

func Fn_set_printerr_handler(func_ string) {}

func Fn_setenv(variable string, value string, overwrite string) {}

func Fn_shell_error_quark() {}

func Fn_shell_parse_argv(commandLine string, argcp string, argvp string) {}

func Fn_shell_quote(unquotedString string) {}

func Fn_shell_unquote(quotedString string) {}

func Fn_slice_get_config(ckey string) {}

func Fn_slice_get_config_state(ckey string, address string, nValues string) {}

func Fn_slice_set_config(ckey string, value string) {}

// UNSUPPORTED : snprintf : has varargs

func Fn_source_remove(tag string) {}

func Fn_source_remove_by_funcs_user_data(funcs string, userData string) {}

func Fn_source_remove_by_user_data(userData string) {}

func Fn_spaced_primes_closest(num string) {}

func Fn_spawn_async(workingDirectory string, argv string, envp string, flags string, childSetup string, userData string, childPid string) {
}

func Fn_spawn_async_with_pipes(workingDirectory string, argv string, envp string, flags string, childSetup string, userData string, childPid string, standardInput string, standardOutput string, standardError string) {
}

func Fn_spawn_close_pid(pid string) {}

func Fn_spawn_command_line_async(commandLine string) {}

func Fn_spawn_command_line_sync(commandLine string, standardOutput string, standardError string, exitStatus string) {
}

func Fn_spawn_error_quark() {}

func Fn_spawn_exit_error_quark() {}

func Fn_spawn_sync(workingDirectory string, argv string, envp string, flags string, childSetup string, userData string, standardOutput string, standardError string, exitStatus string) {
}

// UNSUPPORTED : sprintf : has varargs

func Fn_stpcpy(dest string, src string) {}

func Fn_str_equal(v1 string, v2 string) {}

func Fn_str_has_prefix(str string, prefix string) {}

func Fn_str_has_suffix(str string, suffix string) {}

func Fn_str_hash(v string) {}

func Fn_strcanon(string_ string, validChars string, substitutor string) {}

func Fn_strcasecmp(s1 string, s2 string) {}

func Fn_strchomp(string_ string) {}

func Fn_strchug(string_ string) {}

func Fn_strcompress(source string) {}

// UNSUPPORTED : strconcat : has varargs

func Fn_strdelimit(string_ string, delimiters string, newDelimiter string) {}

func Fn_strdown(string_ string) {}

func Fn_strdup(str string) {}

// UNSUPPORTED : strdup_printf : has varargs

func Fn_strdup_vprintf(format string, args string) {}

func Fn_strdupv(strArray string) {}

func Fn_strerror(errnum string) {}

func Fn_strescape(source string, exceptions string) {}

func Fn_strfreev(strArray string) {}

func Fn_string_new(init string) {}

func Fn_string_new_len(init string, len string) {}

func Fn_string_sized_new(dflSize string) {}

func Fn_strip_context(msgid string, msgval string) {}

// UNSUPPORTED : strjoin : has varargs

func Fn_strjoinv(separator string, strArray string) {}

func Fn_strlcat(dest string, src string, destSize string) {}

func Fn_strlcpy(dest string, src string, destSize string) {}

func Fn_strncasecmp(s1 string, s2 string, n string) {}

func Fn_strndup(str string, n string) {}

func Fn_strnfill(length string, fillChar string) {}

func Fn_strreverse(string_ string) {}

func Fn_strrstr(haystack string, needle string) {}

func Fn_strrstr_len(haystack string, haystackLen string, needle string) {}

func Fn_strsignal(signum string) {}

func Fn_strsplit(string_ string, delimiter string, maxTokens string) {}

func Fn_strsplit_set(string_ string, delimiters string, maxTokens string) {}

func Fn_strstr_len(haystack string, haystackLen string, needle string) {}

func Fn_strtod(nptr string, endptr string) {}

func Fn_strup(string_ string) {}

func Fn_strv_get_type() {}

func Fn_test_add_vtable(testpath string, dataSize string, testData string, dataSetup string, dataTest string, dataTeardown string) {
}

func Fn_test_assert_expected_messages_internal(domain string, file string, line string, func_ string) {}

// UNSUPPORTED : test_build_filename : has varargs

// UNSUPPORTED : test_get_filename : has varargs

// UNSUPPORTED : test_init : has varargs

func Fn_test_log_type_name(logType string) {}

// UNSUPPORTED : test_maximized_result : has varargs

// UNSUPPORTED : test_message : has varargs

// UNSUPPORTED : test_minimized_result : has varargs

func Fn_test_trap_assertions(domain string, file string, line string, func_ string, assertionFlags string, pattern string) {
}

func Fn_thread_error_quark() {}

func Fn_thread_exit(retval string) {}

func Fn_thread_pool_get_max_unused_threads() {}

func Fn_thread_pool_get_num_unused_threads() {}

func Fn_thread_pool_set_max_unused_threads(maxThreads string) {}

func Fn_thread_pool_stop_unused_threads() {}

func Fn_thread_self() {}

func Fn_thread_yield() {}

func Fn_timeout_add(interval string, function string, data string) {}

func Fn_timeout_add_full(priority string, interval string, function string, data string, notify string) {}

func Fn_timeout_source_new(interval string) {}

func Fn_trash_stack_height(stackP string) {}

func Fn_trash_stack_peek(stackP string) {}

func Fn_trash_stack_pop(stackP string) {}

func Fn_trash_stack_push(stackP string, dataP string) {}

func Fn_try_malloc(nBytes string) {}

func Fn_try_realloc(mem string, nBytes string) {}

func Fn_ucs4_to_utf16(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_ucs4_to_utf8(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_unichar_break_type(c string) {}

func Fn_unichar_digit_value(c string) {}

func Fn_unichar_get_mirror_char(ch string, mirroredCh string) {}

func Fn_unichar_isalnum(c string) {}

func Fn_unichar_isalpha(c string) {}

func Fn_unichar_iscntrl(c string) {}

func Fn_unichar_isdefined(c string) {}

func Fn_unichar_isdigit(c string) {}

func Fn_unichar_isgraph(c string) {}

func Fn_unichar_islower(c string) {}

func Fn_unichar_isprint(c string) {}

func Fn_unichar_ispunct(c string) {}

func Fn_unichar_isspace(c string) {}

func Fn_unichar_istitle(c string) {}

func Fn_unichar_isupper(c string) {}

func Fn_unichar_iswide(c string) {}

func Fn_unichar_isxdigit(c string) {}

func Fn_unichar_to_utf8(c string, outbuf string) {}

func Fn_unichar_tolower(c string) {}

func Fn_unichar_totitle(c string) {}

func Fn_unichar_toupper(c string) {}

func Fn_unichar_type(c string) {}

func Fn_unichar_validate(ch string) {}

func Fn_unichar_xdigit_value(c string) {}

func Fn_unicode_canonical_decomposition(ch string, resultLen string) {}

func Fn_unicode_canonical_ordering(string_ string, len string) {}

func Fn_unix_error_quark() {}

func Fn_unsetenv(variable string) {}

func Fn_usleep(microseconds string) {}

func Fn_utf16_to_ucs4(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf16_to_utf8(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf8_casefold(str string, len string) {}

func Fn_utf8_collate(str1 string, str2 string) {}

func Fn_utf8_collate_key(str string, len string) {}

func Fn_utf8_find_next_char(p string, end string) {}

func Fn_utf8_find_prev_char(str string, p string) {}

func Fn_utf8_get_char(p string) {}

func Fn_utf8_get_char_validated(p string, maxLen string) {}

func Fn_utf8_normalize(str string, len string, mode string) {}

func Fn_utf8_offset_to_pointer(str string, offset string) {}

func Fn_utf8_pointer_to_offset(str string, pos string) {}

func Fn_utf8_prev_char(p string) {}

func Fn_utf8_strchr(p string, len string, c string) {}

func Fn_utf8_strdown(str string, len string) {}

func Fn_utf8_strlen(p string, max string) {}

func Fn_utf8_strncpy(dest string, src string, n string) {}

func Fn_utf8_strrchr(p string, len string, c string) {}

func Fn_utf8_strreverse(str string, len string) {}

func Fn_utf8_strup(str string, len string) {}

func Fn_utf8_to_ucs4(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf8_to_ucs4_fast(str string, len string, itemsWritten string) {}

func Fn_utf8_to_utf16(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf8_validate(str string, maxLen string, end string) {}

func Fn_variant_get_gtype() {}

func Fn_variant_parse(type_ string, text string, limit string, endptr string) {}

func Fn_variant_parse_error_quark() {}

func Fn_variant_parser_get_error_quark() {}

func Fn_variant_type_checked_(arg0 string) {}

func Fn_variant_type_string_get_depth_(typeString string) {}

func Fn_variant_type_string_is_valid(typeString string) {}

func Fn_vasprintf(string_ string, format string, args string) {}

func Fn_vfprintf(file string, format string, args string) {}

func Fn_vprintf(format string, args string) {}

func Fn_vsnprintf(string_ string, n string, format string, args string) {}

func Fn_vsprintf(string_ string, format string, args string) {}

func Fn_warn_message(domain string, file string, line string, func_ string, warnexpr string) {}

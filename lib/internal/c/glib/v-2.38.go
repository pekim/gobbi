// Code generated - DO NOT EDIT.
// +build glib_2.38

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
type RegexCompileFlags C.GRegexCompileFlags
type RegexMatchFlags C.GRegexMatchFlags
type SpawnFlags C.GSpawnFlags
type TestSubprocessFlags C.GTestSubprocessFlags
type TestTrapFlags C.GTestTrapFlags
type TraverseFlags C.GTraverseFlags

// enumerations
type BookmarkFileError C.GBookmarkFileError
type ChecksumType C.GChecksumType
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
type RegexError C.GRegexError
type SeekType C.GSeekType
type ShellError C.GShellError
type SliceConfig C.GSliceConfig
type SpawnError C.GSpawnError
type TestFileType C.GTestFileType
type TestLogType C.GTestLogType
type ThreadError C.GThreadError
type TimeType C.GTimeType
type TokenType C.GTokenType
type TraverseType C.GTraverseType
type UnicodeBreakType C.GUnicodeBreakType
type UnicodeScript C.GUnicodeScript
type UnicodeType C.GUnicodeType
type UserDirectory C.GUserDirectory
type VariantClass C.GVariantClass
type VariantParseError C.GVariantParseError

// records
type Array C.GArray
type AsyncQueue C.GAsyncQueue
type BookmarkFile C.GBookmarkFile
type ByteArray C.GByteArray
type Bytes C.GBytes
type Checksum C.GChecksum
type Cond C.GCond
type Data C.GData
type Date C.GDate
type DateTime C.GDateTime
type DebugKey C.GDebugKey
type Dir C.GDir
type Error C.GError
type HashTable C.GHashTable
type HashTableIter C.GHashTableIter
type Hmac C.GHmac
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
type RWLock C.GRWLock
type Rand C.GRand
type RecMutex C.GRecMutex
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
type TestSuite C.GTestSuite
type Thread C.GThread
type ThreadPool C.GThreadPool
type TimeVal C.GTimeVal
type TimeZone C.GTimeZone
type Timer C.GTimer
type TrashStack C.GTrashStack
type Tree C.GTree
type Variant C.GVariant
type VariantBuilder C.GVariantBuilder
type VariantIter C.GVariantIter
type VariantType C.GVariantType

// classes

// interfaces

func Fn_access(filename string, mode string) {}

func Fn_ascii_digit_value(c string) {}

func Fn_ascii_dtostr(buffer string, bufLen string, d string) {}

func Fn_ascii_formatd(buffer string, bufLen string, format string, d string) {}

func Fn_ascii_strcasecmp(s1 string, s2 string) {}

func Fn_ascii_strdown(str string, len string) {}

func Fn_ascii_strncasecmp(s1 string, s2 string, n string) {}

func Fn_ascii_strtod(nptr string, endptr string) {}

func Fn_ascii_strtoll(nptr string, endptr string, base string) {}

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

func Fn_atomic_int_and(atomic string, val string) {}

func Fn_atomic_int_compare_and_exchange(atomic string, oldval string, newval string) {}

func Fn_atomic_int_dec_and_test(atomic string) {}

func Fn_atomic_int_exchange_and_add(atomic string, val string) {}

func Fn_atomic_int_get(atomic string) {}

func Fn_atomic_int_inc(atomic string) {}

func Fn_atomic_int_or(atomic string, val string) {}

func Fn_atomic_int_set(atomic string, newval string) {}

func Fn_atomic_int_xor(atomic string, val string) {}

func Fn_atomic_pointer_add(atomic string, val string) {}

func Fn_atomic_pointer_and(atomic string, val string) {}

func Fn_atomic_pointer_compare_and_exchange(atomic string, oldval string, newval string) {}

func Fn_atomic_pointer_get(atomic string) {}

func Fn_atomic_pointer_or(atomic string, val string) {}

func Fn_atomic_pointer_set(atomic string, newval string) {}

func Fn_atomic_pointer_xor(atomic string, val string) {}

func Fn_base64_decode(text string, outLen string) {}

func Fn_base64_decode_inplace(text string, outLen string) {}

func Fn_base64_decode_step(in string, len string, out string, state string, save string) {}

func Fn_base64_encode(data string, len string) {}

func Fn_base64_encode_close(breakLines string, out string, state string, save string) {}

func Fn_base64_encode_step(in string, len string, breakLines string, out string, state string, save string) {
}

func Fn_basename(fileName string) {}

func Fn_bit_lock(address string, lockBit string) {}

func Fn_bit_nth_lsf(mask string, nthBit string) {}

func Fn_bit_nth_msf(mask string, nthBit string) {}

func Fn_bit_storage(number string) {}

func Fn_bit_trylock(address string, lockBit string) {}

func Fn_bit_unlock(address string, lockBit string) {}

func Fn_bookmark_file_error_quark() {}

// UNSUPPORTED : build_filename : has varargs

func Fn_build_filenamev(args string) {}

// UNSUPPORTED : build_path : has varargs

func Fn_build_pathv(separator string, args string) {}

func Fn_byte_array_free(array string, freeSegment string) {}

func Fn_byte_array_free_to_bytes(array string) {}

func Fn_byte_array_new() {}

func Fn_byte_array_new_take(data string, len string) {}

func Fn_byte_array_unref(array string) {}

func Fn_chdir(path string) {}

func Fn_check_version(requiredMajor string, requiredMinor string, requiredMicro string) {}

func Fn_checksum_type_get_length(checksumType string) {}

func Fn_child_watch_add(pid string, function string, data string) {}

func Fn_child_watch_add_full(priority string, pid string, function string, data string, notify string) {}

func Fn_child_watch_source_new(pid string) {}

func Fn_clear_error() {}

func Fn_clear_pointer(pp string, destroy string) {}

func Fn_close(fd string) {}

func Fn_compute_checksum_for_bytes(checksumType string, data string) {}

func Fn_compute_checksum_for_data(checksumType string, data string, length string) {}

func Fn_compute_checksum_for_string(checksumType string, str string, length string) {}

func Fn_compute_hmac_for_data(digestType string, key string, keyLen string, data string, length string) {}

func Fn_compute_hmac_for_string(digestType string, key string, keyLen string, str string, length string) {
}

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

func Fn_datalist_get_flags(datalist string) {}

func Fn_datalist_id_dup_data(datalist string, keyId string, dupFunc string, userData string) {}

func Fn_datalist_id_get_data(datalist string, keyId string) {}

func Fn_datalist_id_remove_no_notify(datalist string, keyId string) {}

func Fn_datalist_id_replace_data(datalist string, keyId string, oldval string, newval string, destroy string, oldDestroy string) {
}

func Fn_datalist_id_set_data_full(datalist string, keyId string, data string, destroyFunc string) {}

func Fn_datalist_init(datalist string) {}

func Fn_datalist_set_flags(datalist string, flags string) {}

func Fn_datalist_unset_flags(datalist string, flags string) {}

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

func Fn_date_time_compare(dt1 string, dt2 string) {}

func Fn_date_time_equal(dt1 string, dt2 string) {}

func Fn_date_time_hash(datetime string) {}

func Fn_date_valid_day(day string) {}

func Fn_date_valid_dmy(day string, month string, year string) {}

func Fn_date_valid_julian(julianDate string) {}

func Fn_date_valid_month(month string) {}

func Fn_date_valid_weekday(weekday string) {}

func Fn_date_valid_year(year string) {}

func Fn_dcgettext(domain string, msgid string, category string) {}

func Fn_dgettext(domain string, msgid string) {}

func Fn_dir_make_tmp(tmpl string) {}

func Fn_direct_equal(v1 string, v2 string) {}

func Fn_direct_hash(v string) {}

func Fn_dngettext(domain string, msgid string, msgidPlural string, n string) {}

func Fn_double_equal(v1 string, v2 string) {}

func Fn_double_hash(v string) {}

func Fn_dpgettext(domain string, msgctxtid string, msgidoffset string) {}

func Fn_dpgettext2(domain string, context string, msgid string) {}

func Fn_environ_getenv(envp string, variable string) {}

func Fn_environ_setenv(envp string, variable string, value string, overwrite string) {}

func Fn_environ_unsetenv(envp string, variable string) {}

func Fn_file_error_from_errno(errNo string) {}

func Fn_file_error_quark() {}

func Fn_file_get_contents(filename string, contents string, length string) {}

func Fn_file_open_tmp(tmpl string, nameUsed string) {}

func Fn_file_read_link(filename string) {}

func Fn_file_set_contents(filename string, contents string, length string) {}

func Fn_file_test(filename string, test string) {}

func Fn_filename_display_basename(filename string) {}

func Fn_filename_display_name(filename string) {}

func Fn_filename_from_uri(uri string, hostname string) {}

func Fn_filename_from_utf8(utf8string string, len string, bytesRead string, bytesWritten string) {}

func Fn_filename_to_uri(filename string, hostname string) {}

func Fn_filename_to_utf8(opsysstring string, len string, bytesRead string, bytesWritten string) {}

func Fn_find_program_in_path(program string) {}

func Fn_format_size(size string) {}

func Fn_format_size_for_display(size string) {}

func Fn_format_size_full(size string, flags string) {}

// UNSUPPORTED : fprintf : has varargs

func Fn_free(mem string) {}

func Fn_get_application_name() {}

func Fn_get_charset(charset string) {}

func Fn_get_codeset() {}

func Fn_get_current_dir() {}

func Fn_get_current_time(result string) {}

func Fn_get_environ() {}

func Fn_get_filename_charsets(filenameCharsets string) {}

func Fn_get_home_dir() {}

func Fn_get_host_name() {}

func Fn_get_language_names() {}

func Fn_get_locale_variants(locale string) {}

func Fn_get_monotonic_time() {}

func Fn_get_num_processors() {}

func Fn_get_prgname() {}

func Fn_get_real_name() {}

func Fn_get_real_time() {}

func Fn_get_system_config_dirs() {}

func Fn_get_system_data_dirs() {}

func Fn_get_tmp_dir() {}

func Fn_get_user_cache_dir() {}

func Fn_get_user_config_dir() {}

func Fn_get_user_data_dir() {}

func Fn_get_user_name() {}

func Fn_get_user_runtime_dir() {}

func Fn_get_user_special_dir(directory string) {}

func Fn_getenv(variable string) {}

func Fn_hash_table_add(hashTable string, key string) {}

func Fn_hash_table_contains(hashTable string, key string) {}

func Fn_hash_table_destroy(hashTable string) {}

func Fn_hash_table_insert(hashTable string, key string, value string) {}

func Fn_hash_table_lookup(hashTable string, key string) {}

func Fn_hash_table_lookup_extended(hashTable string, lookupKey string, origKey string, value string) {}

func Fn_hash_table_remove(hashTable string, key string) {}

func Fn_hash_table_remove_all(hashTable string) {}

func Fn_hash_table_replace(hashTable string, key string, value string) {}

func Fn_hash_table_size(hashTable string) {}

func Fn_hash_table_steal(hashTable string, key string) {}

func Fn_hash_table_steal_all(hashTable string) {}

func Fn_hash_table_unref(hashTable string) {}

func Fn_hook_destroy(hookList string, hookId string) {}

func Fn_hook_destroy_link(hookList string, hook string) {}

func Fn_hook_free(hookList string, hook string) {}

func Fn_hook_insert_before(hookList string, sibling string, hook string) {}

func Fn_hook_prepend(hookList string, hook string) {}

func Fn_hook_unref(hookList string, hook string) {}

func Fn_hostname_is_ascii_encoded(hostname string) {}

func Fn_hostname_is_ip_address(hostname string) {}

func Fn_hostname_is_non_ascii(hostname string) {}

func Fn_hostname_to_ascii(hostname string) {}

func Fn_hostname_to_unicode(hostname string) {}

func Fn_iconv(converter string, inbuf string, inbytesLeft string, outbuf string, outbytesLeft string) {}

func Fn_iconv_open(toCodeset string, fromCodeset string) {}

func Fn_idle_add(function string, data string) {}

func Fn_idle_add_full(priority string, function string, data string, notify string) {}

func Fn_idle_remove_by_data(data string) {}

func Fn_idle_source_new() {}

func Fn_int64_equal(v1 string, v2 string) {}

func Fn_int64_hash(v string) {}

func Fn_int_equal(v1 string, v2 string) {}

func Fn_int_hash(v string) {}

func Fn_intern_static_string(string_ string) {}

func Fn_intern_string(string_ string) {}

func Fn_io_add_watch(channel string, condition string, func_ string, userData string) {}

func Fn_io_add_watch_full(channel string, priority string, condition string, func_ string, userData string, notify string) {
}

func Fn_io_channel_error_from_errno(en string) {}

func Fn_io_channel_error_quark() {}

func Fn_io_create_watch(channel string, condition string) {}

func Fn_key_file_error_quark() {}

func Fn_listenv() {}

func Fn_locale_from_utf8(utf8string string, len string, bytesRead string, bytesWritten string) {}

func Fn_locale_to_utf8(opsysstring string, len string, bytesRead string, bytesWritten string) {}

// UNSUPPORTED : log : has varargs

func Fn_log_default_handler(logDomain string, logLevel string, message string, unusedData string) {}

func Fn_log_remove_handler(logDomain string, handlerId string) {}

func Fn_log_set_always_fatal(fatalMask string) {}

func Fn_log_set_default_handler(logFunc string, userData string) {}

func Fn_log_set_fatal_mask(logDomain string, fatalMask string) {}

func Fn_log_set_handler(logDomain string, logLevels string, logFunc string, userData string) {}

// UNSUPPORTED : log_structured : has varargs

// UNSUPPORTED : log_structured_standard : has varargs

func Fn_logv(logDomain string, logLevel string, format string, args string) {}

func Fn_main_context_default() {}

func Fn_main_context_get_thread_default() {}

func Fn_main_context_ref_thread_default() {}

func Fn_main_current_source() {}

func Fn_main_depth() {}

func Fn_malloc(nBytes string) {}

func Fn_malloc0(nBytes string) {}

func Fn_malloc0_n(nBlocks string, nBlockBytes string) {}

func Fn_malloc_n(nBlocks string, nBlockBytes string) {}

// UNSUPPORTED : markup_collect_attributes : has varargs

func Fn_markup_error_quark() {}

func Fn_markup_escape_text(text string, length string) {}

// UNSUPPORTED : markup_printf_escaped : has varargs

func Fn_markup_vprintf_escaped(format string, args string) {}

func Fn_mem_is_system_malloc() {}

func Fn_mem_profile() {}

func Fn_mem_set_vtable(vtable string) {}

func Fn_memdup(mem string, byteSize string) {}

func Fn_mkdir_with_parents(pathname string, mode string) {}

func Fn_mkdtemp(tmpl string) {}

func Fn_mkdtemp_full(tmpl string, mode string) {}

func Fn_mkstemp(tmpl string) {}

func Fn_mkstemp_full(tmpl string, flags string, mode string) {}

func Fn_nullify_pointer(nullifyLocation string) {}

func Fn_number_parser_error_quark() {}

func Fn_on_error_query(prgName string) {}

func Fn_on_error_stack_trace(prgName string) {}

func Fn_once_init_enter(location string) {}

func Fn_once_init_leave(location string, result string) {}

func Fn_option_error_quark() {}

func Fn_parse_debug_string(string_ string, keys string, nkeys string) {}

func Fn_path_get_basename(fileName string) {}

func Fn_path_get_dirname(fileName string) {}

func Fn_path_is_absolute(fileName string) {}

func Fn_path_skip_root(fileName string) {}

func Fn_pattern_match(pspec string, stringLength string, string_ string, stringReversed string) {}

func Fn_pattern_match_simple(pattern string, string_ string) {}

func Fn_pattern_match_string(pspec string, string_ string) {}

func Fn_pointer_bit_lock(address string, lockBit string) {}

func Fn_pointer_bit_trylock(address string, lockBit string) {}

func Fn_pointer_bit_unlock(address string, lockBit string) {}

func Fn_poll(fds string, nfds string, timeout string) {}

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

func Fn_realloc_n(mem string, nBlocks string, nBlockBytes string) {}

func Fn_regex_check_replacement(replacement string, hasReferences string) {}

func Fn_regex_error_quark() {}

func Fn_regex_escape_nul(string_ string, length string) {}

func Fn_regex_escape_string(string_ string, length string) {}

func Fn_regex_match_simple(pattern string, string_ string, compileOptions string, matchOptions string) {}

func Fn_regex_split_simple(pattern string, string_ string, compileOptions string, matchOptions string) {}

func Fn_reload_user_special_dirs_cache() {}

func Fn_return_if_fail_warning(logDomain string, prettyFunction string, expression string) {}

func Fn_rmdir(filename string) {}

func Fn_sequence_get(iter string) {}

func Fn_sequence_insert_before(iter string, data string) {}

func Fn_sequence_move(src string, dest string) {}

func Fn_sequence_move_range(dest string, begin string, end string) {}

func Fn_sequence_range_get_midpoint(begin string, end string) {}

func Fn_sequence_remove(iter string) {}

func Fn_sequence_remove_range(begin string, end string) {}

func Fn_sequence_set(iter string, data string) {}

func Fn_sequence_swap(a string, b string) {}

func Fn_set_application_name(applicationName string) {}

// UNSUPPORTED : set_error : has varargs

func Fn_set_error_literal(err string, domain string, code string, message string) {}

func Fn_set_prgname(prgname string) {}

func Fn_set_print_handler(func_ string) {}

func Fn_set_printerr_handler(func_ string) {}

func Fn_setenv(variable string, value string, overwrite string) {}

func Fn_shell_error_quark() {}

func Fn_shell_parse_argv(commandLine string, argcp string, argvp string) {}

func Fn_shell_quote(unquotedString string) {}

func Fn_shell_unquote(quotedString string) {}

func Fn_slice_alloc(blockSize string) {}

func Fn_slice_alloc0(blockSize string) {}

func Fn_slice_copy(blockSize string, memBlock string) {}

func Fn_slice_free1(blockSize string, memBlock string) {}

func Fn_slice_free_chain_with_offset(blockSize string, memChain string, nextOffset string) {}

func Fn_slice_get_config(ckey string) {}

func Fn_slice_get_config_state(ckey string, address string, nValues string) {}

func Fn_slice_set_config(ckey string, value string) {}

// UNSUPPORTED : snprintf : has varargs

func Fn_source_remove(tag string) {}

func Fn_source_remove_by_funcs_user_data(funcs string, userData string) {}

func Fn_source_remove_by_user_data(userData string) {}

func Fn_source_set_name_by_id(tag string, name string) {}

func Fn_spaced_primes_closest(num string) {}

func Fn_spawn_async(workingDirectory string, argv string, envp string, flags string, childSetup string, userData string, childPid string) {
}

func Fn_spawn_async_with_pipes(workingDirectory string, argv string, envp string, flags string, childSetup string, userData string, childPid string, standardInput string, standardOutput string, standardError string) {
}

func Fn_spawn_check_exit_status(exitStatus string) {}

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

func Fn_strcmp0(str1 string, str2 string) {}

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

func Fn_strv_length(strArray string) {}

func Fn_test_add_data_func(testpath string, testData string, testFunc string) {}

func Fn_test_add_data_func_full(testpath string, testData string, testFunc string, dataFreeFunc string) {
}

func Fn_test_add_func(testpath string, testFunc string) {}

func Fn_test_add_vtable(testpath string, dataSize string, testData string, dataSetup string, dataTest string, dataTeardown string) {
}

func Fn_test_assert_expected_messages_internal(domain string, file string, line string, func_ string) {}

func Fn_test_bug(bugUriSnippet string) {}

func Fn_test_bug_base(uriPattern string) {}

// UNSUPPORTED : test_build_filename : has varargs

func Fn_test_create_case(testName string, dataSize string, testData string, dataSetup string, dataTest string, dataTeardown string) {
}

func Fn_test_create_suite(suiteName string) {}

func Fn_test_expect_message(logDomain string, logLevel string, pattern string) {}

func Fn_test_fail() {}

func Fn_test_failed() {}

func Fn_test_get_dir(fileType string) {}

// UNSUPPORTED : test_get_filename : has varargs

func Fn_test_get_root() {}

func Fn_test_incomplete(msg string) {}

// UNSUPPORTED : test_init : has varargs

func Fn_test_log_set_fatal_handler(logFunc string, userData string) {}

func Fn_test_log_type_name(logType string) {}

// UNSUPPORTED : test_maximized_result : has varargs

// UNSUPPORTED : test_message : has varargs

// UNSUPPORTED : test_minimized_result : has varargs

func Fn_test_queue_destroy(destroyFunc string, destroyData string) {}

func Fn_test_queue_free(gfreePointer string) {}

func Fn_test_rand_double() {}

func Fn_test_rand_double_range(rangeStart string, rangeEnd string) {}

func Fn_test_rand_int() {}

func Fn_test_rand_int_range(begin string, end string) {}

func Fn_test_run() {}

func Fn_test_run_suite(suite string) {}

func Fn_test_set_nonfatal_assertions() {}

func Fn_test_skip(msg string) {}

func Fn_test_subprocess() {}

func Fn_test_timer_elapsed() {}

func Fn_test_timer_last() {}

func Fn_test_timer_start() {}

func Fn_test_trap_assertions(domain string, file string, line string, func_ string, assertionFlags string, pattern string) {
}

func Fn_test_trap_fork(usecTimeout string, testTrapFlags string) {}

func Fn_test_trap_has_passed() {}

func Fn_test_trap_reached_timeout() {}

func Fn_test_trap_subprocess(testPath string, usecTimeout string, testFlags string) {}

func Fn_thread_error_quark() {}

func Fn_thread_exit(retval string) {}

func Fn_thread_pool_get_max_idle_time() {}

func Fn_thread_pool_get_max_unused_threads() {}

func Fn_thread_pool_get_num_unused_threads() {}

func Fn_thread_pool_set_max_idle_time(interval string) {}

func Fn_thread_pool_set_max_unused_threads(maxThreads string) {}

func Fn_thread_pool_stop_unused_threads() {}

func Fn_thread_self() {}

func Fn_thread_yield() {}

func Fn_time_val_from_iso8601(isoDate string, time string) {}

func Fn_timeout_add(interval string, function string, data string) {}

func Fn_timeout_add_full(priority string, interval string, function string, data string, notify string) {}

func Fn_timeout_add_seconds(interval string, function string, data string) {}

func Fn_timeout_add_seconds_full(priority string, interval string, function string, data string, notify string) {
}

func Fn_timeout_source_new(interval string) {}

func Fn_timeout_source_new_seconds(interval string) {}

func Fn_trash_stack_height(stackP string) {}

func Fn_trash_stack_peek(stackP string) {}

func Fn_trash_stack_pop(stackP string) {}

func Fn_trash_stack_push(stackP string, dataP string) {}

func Fn_try_malloc(nBytes string) {}

func Fn_try_malloc0(nBytes string) {}

func Fn_try_malloc0_n(nBlocks string, nBlockBytes string) {}

func Fn_try_malloc_n(nBlocks string, nBlockBytes string) {}

func Fn_try_realloc(mem string, nBytes string) {}

func Fn_try_realloc_n(mem string, nBlocks string, nBlockBytes string) {}

func Fn_ucs4_to_utf16(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_ucs4_to_utf8(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_unichar_break_type(c string) {}

func Fn_unichar_combining_class(uc string) {}

func Fn_unichar_compose(a string, b string, ch string) {}

func Fn_unichar_decompose(ch string, a string, b string) {}

func Fn_unichar_digit_value(c string) {}

func Fn_unichar_fully_decompose(ch string, compat string, result string, resultLen string) {}

func Fn_unichar_get_mirror_char(ch string, mirroredCh string) {}

func Fn_unichar_get_script(ch string) {}

func Fn_unichar_isalnum(c string) {}

func Fn_unichar_isalpha(c string) {}

func Fn_unichar_iscntrl(c string) {}

func Fn_unichar_isdefined(c string) {}

func Fn_unichar_isdigit(c string) {}

func Fn_unichar_isgraph(c string) {}

func Fn_unichar_islower(c string) {}

func Fn_unichar_ismark(c string) {}

func Fn_unichar_isprint(c string) {}

func Fn_unichar_ispunct(c string) {}

func Fn_unichar_isspace(c string) {}

func Fn_unichar_istitle(c string) {}

func Fn_unichar_isupper(c string) {}

func Fn_unichar_iswide(c string) {}

func Fn_unichar_iswide_cjk(c string) {}

func Fn_unichar_isxdigit(c string) {}

func Fn_unichar_iszerowidth(c string) {}

func Fn_unichar_to_utf8(c string, outbuf string) {}

func Fn_unichar_tolower(c string) {}

func Fn_unichar_totitle(c string) {}

func Fn_unichar_toupper(c string) {}

func Fn_unichar_type(c string) {}

func Fn_unichar_validate(ch string) {}

func Fn_unichar_xdigit_value(c string) {}

func Fn_unicode_canonical_decomposition(ch string, resultLen string) {}

func Fn_unicode_canonical_ordering(string_ string, len string) {}

func Fn_unicode_script_from_iso15924(iso15924 string) {}

func Fn_unicode_script_to_iso15924(script string) {}

func Fn_unix_error_quark() {}

func Fn_unix_fd_add(fd string, condition string, function string, userData string) {}

func Fn_unix_fd_add_full(priority string, fd string, condition string, function string, userData string, notify string) {
}

func Fn_unix_fd_source_new(fd string, condition string) {}

func Fn_unix_open_pipe(fds string, flags string) {}

func Fn_unix_set_fd_nonblocking(fd string, nonblock string) {}

func Fn_unix_signal_add(signum string, handler string, userData string) {}

func Fn_unix_signal_add_full(priority string, signum string, handler string, userData string, notify string) {
}

func Fn_unix_signal_source_new(signum string) {}

func Fn_unlink(filename string) {}

func Fn_unsetenv(variable string) {}

func Fn_uri_escape_string(unescaped string, reservedCharsAllowed string, allowUtf8 string) {}

func Fn_uri_list_extract_uris(uriList string) {}

func Fn_uri_parse_scheme(uri string) {}

func Fn_uri_unescape_segment(escapedString string, escapedStringEnd string, illegalCharacters string) {}

func Fn_uri_unescape_string(escapedString string, illegalCharacters string) {}

func Fn_usleep(microseconds string) {}

func Fn_utf16_to_ucs4(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf16_to_utf8(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf8_casefold(str string, len string) {}

func Fn_utf8_collate(str1 string, str2 string) {}

func Fn_utf8_collate_key(str string, len string) {}

func Fn_utf8_collate_key_for_filename(str string, len string) {}

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

func Fn_utf8_substring(str string, startPos string, endPos string) {}

func Fn_utf8_to_ucs4(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf8_to_ucs4_fast(str string, len string, itemsWritten string) {}

func Fn_utf8_to_utf16(str string, len string, itemsRead string, itemsWritten string) {}

func Fn_utf8_validate(str string, maxLen string, end string) {}

func Fn_variant_get_gtype() {}

func Fn_variant_is_object_path(string_ string) {}

func Fn_variant_is_signature(string_ string) {}

func Fn_variant_parse(type_ string, text string, limit string, endptr string) {}

func Fn_variant_parse_error_quark() {}

func Fn_variant_parser_get_error_quark() {}

func Fn_variant_type_checked_(arg0 string) {}

func Fn_variant_type_string_get_depth_(typeString string) {}

func Fn_variant_type_string_is_valid(typeString string) {}

func Fn_variant_type_string_scan(string_ string, limit string, endptr string) {}

func Fn_vasprintf(string_ string, format string, args string) {}

func Fn_vfprintf(file string, format string, args string) {}

func Fn_vprintf(format string, args string) {}

func Fn_vsnprintf(string_ string, n string, format string, args string) {}

func Fn_vsprintf(string_ string, format string, args string) {}

func Fn_warn_message(domain string, file string, line string, func_ string, warnexpr string) {}

// Code generated - DO NOT EDIT.
// +build glib_2.14

package glib

import "unsafe"

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-2.0/glib-object.h>
import "C"

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

func Fn_access(filename string, mode int) {}

func Fn_ascii_digit_value(c int8) {}

func Fn_ascii_dtostr(buffer string, bufLen int, d float64) {}

func Fn_ascii_formatd(buffer string, bufLen int, format string, d float64) {}

func Fn_ascii_strcasecmp(s1 string, s2 string) {}

func Fn_ascii_strdown(str string, len uint64) {}

func Fn_ascii_strncasecmp(s1 string, s2 string, n uint64) {}

func Fn_ascii_strtod(nptr string, endptr string) {}

func Fn_ascii_strtoll(nptr string, endptr string, base uint) {}

func Fn_ascii_strtoull(nptr string, endptr string, base uint) {}

func Fn_ascii_strup(str string, len uint64) {}

func Fn_ascii_tolower(c int8) {}

func Fn_ascii_toupper(c int8) {}

func Fn_ascii_xdigit_value(c int8) {}

func Fn_assert_warning(logDomain string, file string, line int, prettyFunction string, expression string) {
}

func Fn_assertion_message(domain string, file string, line int, func_ string, message string) {}

func Fn_assertion_message_cmpnum(domain string, file string, line int, func_ string, expr string, arg1 float64, cmp string, arg2 float64, numtype int8) {
}

func Fn_assertion_message_cmpstr(domain string, file string, line int, func_ string, expr string, arg1 string, cmp string, arg2 string) {
}

func Fn_assertion_message_error(domain string, file string, line int, func_ string, expr string, error unsafe.Pointer, errorDomain uint32, errorCode int) {
}

func Fn_assertion_message_expr(domain string, file string, line int, func_ string, expr string) {}

// UNSUPPORTED : atexit : has callback

func Fn_atomic_int_add(atomic *int, val int) {}

func Fn_atomic_int_compare_and_exchange(atomic *int, oldval int, newval int) {}

func Fn_atomic_int_dec_and_test(atomic *int) {}

func Fn_atomic_int_exchange_and_add(atomic *int, val int) {}

func Fn_atomic_int_get(atomic *int) {}

func Fn_atomic_int_inc(atomic *int) {}

func Fn_atomic_int_set(atomic *int, newval int) {}

func Fn_atomic_pointer_compare_and_exchange(atomic unsafe.Pointer, oldval unsafe.Pointer, newval unsafe.Pointer) {
}

func Fn_atomic_pointer_get(atomic unsafe.Pointer) {}

func Fn_atomic_pointer_set(atomic unsafe.Pointer, newval unsafe.Pointer) {}

// UNSUPPORTED : atomic_rc_box_release_full : has callback

func Fn_base64_decode(text string, outLen *uint64) {}

func Fn_base64_decode_step(in *uint8, len uint64, out *uint8, state *int, save *uint) {}

func Fn_base64_encode(data *uint8, len uint64) {}

func Fn_base64_encode_close(breakLines bool, out *uint8, state *int, save *int) {}

func Fn_base64_encode_step(in *uint8, len uint64, breakLines bool, out *uint8, state *int, save *int) {
}

func Fn_basename(fileName string) {}

func Fn_bit_nth_lsf(mask uint64, nthBit int) {}

func Fn_bit_nth_msf(mask uint64, nthBit int) {}

func Fn_bit_storage(number uint64) {}

func Fn_bookmark_file_error_quark() {
	C.g_bookmark_file_error_quark()
}

// UNSUPPORTED : build_filename : has varargs

// UNSUPPORTED : build_filename_valist : has va_list

func Fn_build_filenamev(args *string) {}

// UNSUPPORTED : build_path : has varargs

func Fn_build_pathv(separator string, args *string) {}

func Fn_byte_array_free(array *uint8, freeSegment bool) {}

func Fn_byte_array_new() {
	C.g_byte_array_new()
}

func Fn_chdir(path string) {}

func Fn_check_version(requiredMajor uint, requiredMinor uint, requiredMicro uint) {}

// UNSUPPORTED : child_watch_add : has callback

// UNSUPPORTED : child_watch_add_full : has callback

func Fn_child_watch_source_new(pid int) {}

func Fn_clear_error(err *unsafe.Pointer) {}

// UNSUPPORTED : clear_handle_id : has callback

// UNSUPPORTED : clear_pointer : has callback

func Fn_convert(str *uint8, len uint64, toCodeset string, fromCodeset string, bytesRead *uint64, bytesWritten *uint64) {
}

func Fn_convert_error_quark() {
	C.g_convert_error_quark()
}

func Fn_convert_with_fallback(str *uint8, len uint64, toCodeset string, fromCodeset string, fallback string, bytesRead *uint64, bytesWritten *uint64) {
}

func Fn_convert_with_iconv(str *uint8, len uint64, converter IConv, bytesRead *uint64, bytesWritten *uint64) {
}

func Fn_datalist_clear(datalist *unsafe.Pointer) {}

// UNSUPPORTED : datalist_foreach : has callback

func Fn_datalist_get_data(datalist *unsafe.Pointer, key string) {}

func Fn_datalist_get_flags(datalist *unsafe.Pointer) {}

// UNSUPPORTED : datalist_id_dup_data : has callback

func Fn_datalist_id_get_data(datalist *unsafe.Pointer, keyId uint32) {}

func Fn_datalist_id_remove_no_notify(datalist *unsafe.Pointer, keyId uint32) {}

// UNSUPPORTED : datalist_id_replace_data : has callback

// UNSUPPORTED : datalist_id_set_data_full : has callback

func Fn_datalist_init(datalist *unsafe.Pointer) {}

func Fn_datalist_set_flags(datalist *unsafe.Pointer, flags uint) {}

func Fn_datalist_unset_flags(datalist *unsafe.Pointer, flags uint) {}

func Fn_dataset_destroy(datasetLocation unsafe.Pointer) {}

// UNSUPPORTED : dataset_foreach : has callback

func Fn_dataset_id_get_data(datasetLocation unsafe.Pointer, keyId uint32) {}

func Fn_dataset_id_remove_no_notify(datasetLocation unsafe.Pointer, keyId uint32) {}

// UNSUPPORTED : dataset_id_set_data_full : has callback

func Fn_date_get_days_in_month(month int, year uint16) {}

func Fn_date_get_monday_weeks_in_year(year uint16) {}

func Fn_date_get_sunday_weeks_in_year(year uint16) {}

func Fn_date_is_leap_year(year uint16) {}

func Fn_date_strftime(s string, slen uint64, format string, date unsafe.Pointer) {}

func Fn_date_valid_day(day uint8) {}

func Fn_date_valid_dmy(day uint8, month int, year uint16) {}

func Fn_date_valid_julian(julianDate uint32) {}

func Fn_date_valid_month(month int) {}

func Fn_date_valid_weekday(weekday int) {}

func Fn_date_valid_year(year uint16) {}

func Fn_direct_equal(v1 unsafe.Pointer, v2 unsafe.Pointer) {}

func Fn_direct_hash(v unsafe.Pointer) {}

func Fn_file_error_from_errno(errNo int) {}

func Fn_file_error_quark() {
	C.g_file_error_quark()
}

func Fn_file_get_contents(filename string, contents *uint8, length *uint64) {}

func Fn_file_open_tmp(tmpl string, nameUsed string) {}

func Fn_file_read_link(filename string) {}

func Fn_file_set_contents(filename string, contents *uint8, length uint64) {}

func Fn_file_test(filename string, test int) {}

func Fn_filename_display_basename(filename string) {}

func Fn_filename_display_name(filename string) {}

func Fn_filename_from_uri(uri string, hostname string) {}

func Fn_filename_from_utf8(utf8string string, len uint64, bytesRead *uint64, bytesWritten *uint64) {}

func Fn_filename_to_uri(filename string, hostname string) {}

func Fn_filename_to_utf8(opsysstring string, len uint64, bytesRead *uint64, bytesWritten *uint64) {}

func Fn_find_program_in_path(program string) {}

// UNSUPPORTED : fprintf : has varargs

func Fn_free(mem unsafe.Pointer) {}

func Fn_get_application_name() {
	C.g_get_application_name()
}

func Fn_get_charset(charset string) {}

func Fn_get_codeset() {
	C.g_get_codeset()
}

func Fn_get_current_dir() {
	C.g_get_current_dir()
}

func Fn_get_current_time(result unsafe.Pointer) {}

func Fn_get_filename_charsets(filenameCharsets *string) {}

func Fn_get_home_dir() {
	C.g_get_home_dir()
}

func Fn_get_host_name() {
	C.g_get_host_name()
}

func Fn_get_language_names() {
	C.g_get_language_names()
}

func Fn_get_prgname() {
	C.g_get_prgname()
}

func Fn_get_real_name() {
	C.g_get_real_name()
}

func Fn_get_system_config_dirs() {
	C.g_get_system_config_dirs()
}

func Fn_get_system_data_dirs() {
	C.g_get_system_data_dirs()
}

func Fn_get_tmp_dir() {
	C.g_get_tmp_dir()
}

func Fn_get_user_cache_dir() {
	C.g_get_user_cache_dir()
}

func Fn_get_user_config_dir() {
	C.g_get_user_config_dir()
}

func Fn_get_user_data_dir() {
	C.g_get_user_data_dir()
}

func Fn_get_user_name() {
	C.g_get_user_name()
}

func Fn_get_user_special_dir(directory int) {}

func Fn_getenv(variable string) {}

func Fn_hash_table_destroy(hashTable unsafe.Pointer) {}

func Fn_hash_table_insert(hashTable unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {}

func Fn_hash_table_lookup(hashTable unsafe.Pointer, key unsafe.Pointer) {}

func Fn_hash_table_lookup_extended(hashTable unsafe.Pointer, lookupKey unsafe.Pointer, origKey *unsafe.Pointer, value *unsafe.Pointer) {
}

func Fn_hash_table_remove(hashTable unsafe.Pointer, key unsafe.Pointer) {}

func Fn_hash_table_remove_all(hashTable unsafe.Pointer) {}

func Fn_hash_table_replace(hashTable unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {}

func Fn_hash_table_size(hashTable unsafe.Pointer) {}

func Fn_hash_table_steal(hashTable unsafe.Pointer, key unsafe.Pointer) {}

func Fn_hash_table_steal_all(hashTable unsafe.Pointer) {}

func Fn_hash_table_unref(hashTable unsafe.Pointer) {}

func Fn_hook_destroy(hookList unsafe.Pointer, hookId uint64) {}

func Fn_hook_destroy_link(hookList unsafe.Pointer, hook unsafe.Pointer) {}

func Fn_hook_free(hookList unsafe.Pointer, hook unsafe.Pointer) {}

func Fn_hook_insert_before(hookList unsafe.Pointer, sibling unsafe.Pointer, hook unsafe.Pointer) {}

func Fn_hook_prepend(hookList unsafe.Pointer, hook unsafe.Pointer) {}

func Fn_hook_unref(hookList unsafe.Pointer, hook unsafe.Pointer) {}

func Fn_iconv(converter IConv, inbuf string, inbytesLeft *uint64, outbuf string, outbytesLeft *uint64) {
}

func Fn_iconv_open(toCodeset string, fromCodeset string) {}

// UNSUPPORTED : idle_add : has callback

// UNSUPPORTED : idle_add_full : has callback

func Fn_idle_remove_by_data(data unsafe.Pointer) {}

func Fn_idle_source_new() {
	C.g_idle_source_new()
}

func Fn_int_equal(v1 unsafe.Pointer, v2 unsafe.Pointer) {}

func Fn_int_hash(v unsafe.Pointer) {}

func Fn_intern_static_string(string_ string) {}

func Fn_intern_string(string_ string) {}

// UNSUPPORTED : io_add_watch : has callback

// UNSUPPORTED : io_add_watch_full : has callback

func Fn_io_channel_error_from_errno(en int) {}

func Fn_io_channel_error_quark() {
	C.g_io_channel_error_quark()
}

func Fn_io_create_watch(channel unsafe.Pointer, condition int) {}

func Fn_key_file_error_quark() {
	C.g_key_file_error_quark()
}

func Fn_listenv() {
	C.g_listenv()
}

func Fn_locale_from_utf8(utf8string string, len uint64, bytesRead *uint64, bytesWritten *uint64) {}

func Fn_locale_to_utf8(opsysstring *uint8, len uint64, bytesRead *uint64, bytesWritten *uint64) {}

// UNSUPPORTED : log : has varargs

func Fn_log_default_handler(logDomain string, logLevel int, message string, unusedData unsafe.Pointer) {
}

func Fn_log_remove_handler(logDomain string, handlerId uint) {}

func Fn_log_set_always_fatal(fatalMask int) {}

// UNSUPPORTED : log_set_default_handler : has callback

func Fn_log_set_fatal_mask(logDomain string, fatalMask int) {}

// UNSUPPORTED : log_set_handler : has callback

// UNSUPPORTED : log_set_handler_full : has callback

// UNSUPPORTED : log_set_writer_func : has callback

// UNSUPPORTED : log_structured : has varargs

// UNSUPPORTED : log_structured_standard : has varargs

// UNSUPPORTED : logv : has va_list

func Fn_main_context_default() {
	C.g_main_context_default()
}

func Fn_main_current_source() {
	C.g_main_current_source()
}

func Fn_main_depth() {
	C.g_main_depth()
}

func Fn_malloc(nBytes uint64) {}

func Fn_malloc0(nBytes uint64) {}

// UNSUPPORTED : markup_collect_attributes : has varargs

func Fn_markup_error_quark() {
	C.g_markup_error_quark()
}

func Fn_markup_escape_text(text string, length uint64) {}

// UNSUPPORTED : markup_printf_escaped : has varargs

// UNSUPPORTED : markup_vprintf_escaped : has va_list

func Fn_mem_is_system_malloc() {
	C.g_mem_is_system_malloc()
}

func Fn_mem_profile() {
	C.g_mem_profile()
}

func Fn_mem_set_vtable(vtable unsafe.Pointer) {}

func Fn_memdup(mem unsafe.Pointer, byteSize uint) {}

func Fn_mkdir_with_parents(pathname string, mode int) {}

func Fn_mkstemp(tmpl string) {}

func Fn_nullify_pointer(nullifyLocation *unsafe.Pointer) {}

func Fn_number_parser_error_quark() {
	C.g_number_parser_error_quark()
}

func Fn_on_error_query(prgName string) {}

func Fn_on_error_stack_trace(prgName string) {}

func Fn_once_init_enter(location unsafe.Pointer) {}

func Fn_once_init_leave(location unsafe.Pointer, result uint64) {}

func Fn_option_error_quark() {
	C.g_option_error_quark()
}

func Fn_parse_debug_string(string_ string, keys *DebugKey, nkeys uint) {}

func Fn_path_get_basename(fileName string) {}

func Fn_path_get_dirname(fileName string) {}

func Fn_path_is_absolute(fileName string) {}

func Fn_path_skip_root(fileName string) {}

func Fn_pattern_match(pspec unsafe.Pointer, stringLength uint, string_ string, stringReversed string) {
}

func Fn_pattern_match_simple(pattern string, string_ string) {}

func Fn_pattern_match_string(pspec unsafe.Pointer, string_ string) {}

// UNSUPPORTED : prefix_error : has varargs

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

// UNSUPPORTED : printf : has varargs

// UNSUPPORTED : printf_string_upper_bound : has va_list

func Fn_propagate_error(dest *unsafe.Pointer, src unsafe.Pointer) {}

// UNSUPPORTED : propagate_prefixed_error : has varargs

// UNSUPPORTED : ptr_array_find_with_equal_func : has callback

// UNSUPPORTED : qsort_with_data : has callback

func Fn_quark_from_static_string(string_ string) {}

func Fn_quark_from_string(string_ string) {}

func Fn_quark_to_string(quark uint32) {}

func Fn_quark_try_string(string_ string) {}

func Fn_random_double() {
	C.g_random_double()
}

func Fn_random_double_range(begin float64, end float64) {}

func Fn_random_int() {
	C.g_random_int()
}

func Fn_random_int_range(begin int32, end int32) {}

func Fn_random_set_seed(seed uint32) {}

// UNSUPPORTED : rc_box_release_full : has callback

func Fn_realloc(mem unsafe.Pointer, nBytes uint64) {}

func Fn_regex_check_replacement(replacement string, hasReferences *bool) {}

func Fn_regex_error_quark() {
	C.g_regex_error_quark()
}

func Fn_regex_escape_string(string_ *string, length int) {}

func Fn_regex_match_simple(pattern string, string_ string, compileOptions int, matchOptions int) {}

func Fn_regex_split_simple(pattern string, string_ string, compileOptions int, matchOptions int) {}

func Fn_return_if_fail_warning(logDomain string, prettyFunction string, expression string) {}

func Fn_rmdir(filename string) {}

func Fn_sequence_get(iter unsafe.Pointer) {}

func Fn_sequence_insert_before(iter unsafe.Pointer, data unsafe.Pointer) {}

func Fn_sequence_move(src unsafe.Pointer, dest unsafe.Pointer) {}

func Fn_sequence_move_range(dest unsafe.Pointer, begin unsafe.Pointer, end unsafe.Pointer) {}

func Fn_sequence_range_get_midpoint(begin unsafe.Pointer, end unsafe.Pointer) {}

func Fn_sequence_remove(iter unsafe.Pointer) {}

func Fn_sequence_remove_range(begin unsafe.Pointer, end unsafe.Pointer) {}

func Fn_sequence_set(iter unsafe.Pointer, data unsafe.Pointer) {}

func Fn_sequence_swap(a unsafe.Pointer, b unsafe.Pointer) {}

func Fn_set_application_name(applicationName string) {}

// UNSUPPORTED : set_error : has varargs

func Fn_set_prgname(prgname string) {}

// UNSUPPORTED : set_print_handler : has callback

// UNSUPPORTED : set_printerr_handler : has callback

func Fn_setenv(variable string, value string, overwrite bool) {}

func Fn_shell_error_quark() {
	C.g_shell_error_quark()
}

func Fn_shell_parse_argv(commandLine string, argcp *int, argvp *string) {}

func Fn_shell_quote(unquotedString string) {}

func Fn_shell_unquote(quotedString string) {}

func Fn_slice_alloc(blockSize uint64) {}

func Fn_slice_alloc0(blockSize uint64) {}

func Fn_slice_copy(blockSize uint64, memBlock unsafe.Pointer) {}

func Fn_slice_free1(blockSize uint64, memBlock unsafe.Pointer) {}

func Fn_slice_free_chain_with_offset(blockSize uint64, memChain unsafe.Pointer, nextOffset uint64) {}

func Fn_slice_get_config(ckey int) {}

func Fn_slice_get_config_state(ckey int, address int64, nValues *uint) {}

func Fn_slice_set_config(ckey int, value int64) {}

// UNSUPPORTED : snprintf : has varargs

func Fn_source_remove(tag uint) {}

func Fn_source_remove_by_funcs_user_data(funcs unsafe.Pointer, userData unsafe.Pointer) {}

func Fn_source_remove_by_user_data(userData unsafe.Pointer) {}

func Fn_spaced_primes_closest(num uint) {}

// UNSUPPORTED : spawn_async : has callback

// UNSUPPORTED : spawn_async_with_fds : has callback

// UNSUPPORTED : spawn_async_with_pipes : has callback

func Fn_spawn_close_pid(pid int) {}

func Fn_spawn_command_line_async(commandLine string) {}

func Fn_spawn_command_line_sync(commandLine string, standardOutput *uint8, standardError *uint8, exitStatus *int) {
}

func Fn_spawn_error_quark() {
	C.g_spawn_error_quark()
}

func Fn_spawn_exit_error_quark() {
	C.g_spawn_exit_error_quark()
}

// UNSUPPORTED : spawn_sync : has callback

// UNSUPPORTED : sprintf : has varargs

func Fn_stpcpy(dest string, src string) {}

func Fn_str_equal(v1 unsafe.Pointer, v2 unsafe.Pointer) {}

func Fn_str_has_prefix(str string, prefix string) {}

func Fn_str_has_suffix(str string, suffix string) {}

func Fn_str_hash(v unsafe.Pointer) {}

func Fn_strcanon(string_ string, validChars string, substitutor int8) {}

func Fn_strcasecmp(s1 string, s2 string) {}

func Fn_strchomp(string_ string) {}

func Fn_strchug(string_ string) {}

func Fn_strcompress(source string) {}

// UNSUPPORTED : strconcat : has varargs

func Fn_strdelimit(string_ string, delimiters string, newDelimiter int8) {}

func Fn_strdown(string_ string) {}

func Fn_strdup(str string) {}

// UNSUPPORTED : strdup_printf : has varargs

// UNSUPPORTED : strdup_vprintf : has va_list

func Fn_strdupv(strArray string) {}

func Fn_strerror(errnum int) {}

func Fn_strescape(source string, exceptions string) {}

func Fn_strfreev(strArray string) {}

func Fn_string_new(init string) {}

func Fn_string_new_len(init string, len uint64) {}

func Fn_string_sized_new(dflSize uint64) {}

func Fn_strip_context(msgid string, msgval string) {}

// UNSUPPORTED : strjoin : has varargs

func Fn_strjoinv(separator string, strArray string) {}

func Fn_strlcat(dest string, src string, destSize uint64) {}

func Fn_strlcpy(dest string, src string, destSize uint64) {}

func Fn_strncasecmp(s1 string, s2 string, n uint) {}

func Fn_strndup(str string, n uint64) {}

func Fn_strnfill(length uint64, fillChar int8) {}

func Fn_strreverse(string_ string) {}

func Fn_strrstr(haystack string, needle string) {}

func Fn_strrstr_len(haystack string, haystackLen uint64, needle string) {}

func Fn_strsignal(signum int) {}

func Fn_strsplit(string_ string, delimiter string, maxTokens int) {}

func Fn_strsplit_set(string_ string, delimiters string, maxTokens int) {}

func Fn_strstr_len(haystack string, haystackLen uint64, needle string) {}

func Fn_strtod(nptr string, endptr string) {}

func Fn_strup(string_ string) {}

func Fn_strv_get_type() {
	C.g_strv_get_type()
}

func Fn_strv_length(strArray string) {}

// UNSUPPORTED : test_add_data_func : has callback

// UNSUPPORTED : test_add_data_func_full : has callback

// UNSUPPORTED : test_add_func : has callback

// UNSUPPORTED : test_add_vtable : has callback

func Fn_test_assert_expected_messages_internal(domain string, file string, line int, func_ string) {}

// UNSUPPORTED : test_build_filename : has varargs

// UNSUPPORTED : test_create_case : has callback

// UNSUPPORTED : test_get_filename : has varargs

// UNSUPPORTED : test_init : has varargs

// UNSUPPORTED : test_log_set_fatal_handler : has callback

func Fn_test_log_type_name(logType int) {}

// UNSUPPORTED : test_maximized_result : has varargs

// UNSUPPORTED : test_message : has varargs

// UNSUPPORTED : test_minimized_result : has varargs

// UNSUPPORTED : test_queue_destroy : has callback

func Fn_test_trap_assertions(domain string, file string, line int, func_ string, assertionFlags uint64, pattern string) {
}

func Fn_thread_error_quark() {
	C.g_thread_error_quark()
}

func Fn_thread_exit(retval unsafe.Pointer) {}

func Fn_thread_pool_get_max_idle_time() {
	C.g_thread_pool_get_max_idle_time()
}

func Fn_thread_pool_get_max_unused_threads() {
	C.g_thread_pool_get_max_unused_threads()
}

func Fn_thread_pool_get_num_unused_threads() {
	C.g_thread_pool_get_num_unused_threads()
}

func Fn_thread_pool_set_max_idle_time(interval uint) {}

func Fn_thread_pool_set_max_unused_threads(maxThreads int) {}

func Fn_thread_pool_stop_unused_threads() {
	C.g_thread_pool_stop_unused_threads()
}

func Fn_thread_self() {
	C.g_thread_self()
}

func Fn_thread_yield() {
	C.g_thread_yield()
}

func Fn_time_val_from_iso8601(isoDate string, time unsafe.Pointer) {}

// UNSUPPORTED : timeout_add : has callback

// UNSUPPORTED : timeout_add_full : has callback

// UNSUPPORTED : timeout_add_seconds : has callback

// UNSUPPORTED : timeout_add_seconds_full : has callback

func Fn_timeout_source_new(interval uint) {}

func Fn_timeout_source_new_seconds(interval uint) {}

func Fn_trash_stack_height(stackP *unsafe.Pointer) {}

func Fn_trash_stack_peek(stackP *unsafe.Pointer) {}

func Fn_trash_stack_pop(stackP *unsafe.Pointer) {}

func Fn_trash_stack_push(stackP *unsafe.Pointer, dataP unsafe.Pointer) {}

func Fn_try_malloc(nBytes uint64) {}

func Fn_try_malloc0(nBytes uint64) {}

func Fn_try_realloc(mem unsafe.Pointer, nBytes uint64) {}

func Fn_ucs4_to_utf16(str *rune, len int64, itemsRead *int64, itemsWritten *int64) {}

func Fn_ucs4_to_utf8(str *rune, len int64, itemsRead *int64, itemsWritten *int64) {}

func Fn_unichar_break_type(c rune) {}

func Fn_unichar_combining_class(uc rune) {}

func Fn_unichar_digit_value(c rune) {}

func Fn_unichar_get_mirror_char(ch rune, mirroredCh *rune) {}

func Fn_unichar_get_script(ch rune) {}

func Fn_unichar_isalnum(c rune) {}

func Fn_unichar_isalpha(c rune) {}

func Fn_unichar_iscntrl(c rune) {}

func Fn_unichar_isdefined(c rune) {}

func Fn_unichar_isdigit(c rune) {}

func Fn_unichar_isgraph(c rune) {}

func Fn_unichar_islower(c rune) {}

func Fn_unichar_ismark(c rune) {}

func Fn_unichar_isprint(c rune) {}

func Fn_unichar_ispunct(c rune) {}

func Fn_unichar_isspace(c rune) {}

func Fn_unichar_istitle(c rune) {}

func Fn_unichar_isupper(c rune) {}

func Fn_unichar_iswide(c rune) {}

func Fn_unichar_iswide_cjk(c rune) {}

func Fn_unichar_isxdigit(c rune) {}

func Fn_unichar_iszerowidth(c rune) {}

func Fn_unichar_to_utf8(c rune, outbuf string) {}

func Fn_unichar_tolower(c rune) {}

func Fn_unichar_totitle(c rune) {}

func Fn_unichar_toupper(c rune) {}

func Fn_unichar_type(c rune) {}

func Fn_unichar_validate(ch rune) {}

func Fn_unichar_xdigit_value(c rune) {}

func Fn_unicode_canonical_decomposition(ch rune, resultLen *uint64) {}

func Fn_unicode_canonical_ordering(string_ *rune, len uint64) {}

// UNSUPPORTED : unix_error_quark : blacklisted
// UNSUPPORTED : unix_fd_add : has callback

// UNSUPPORTED : unix_fd_add_full : has callback

// UNSUPPORTED : unix_signal_add : has callback

// UNSUPPORTED : unix_signal_add_full : has callback

func Fn_unlink(filename string) {}

func Fn_unsetenv(variable string) {}

func Fn_uri_list_extract_uris(uriList string) {}

func Fn_usleep(microseconds uint64) {}

func Fn_utf16_to_ucs4(str *uint16, len int64, itemsRead *int64, itemsWritten *int64) {}

func Fn_utf16_to_utf8(str *uint16, len int64, itemsRead *int64, itemsWritten *int64) {}

func Fn_utf8_casefold(str string, len uint64) {}

func Fn_utf8_collate(str1 string, str2 string) {}

func Fn_utf8_collate_key(str string, len uint64) {}

func Fn_utf8_collate_key_for_filename(str string, len uint64) {}

func Fn_utf8_find_next_char(p string, end string) {}

func Fn_utf8_find_prev_char(str string, p string) {}

func Fn_utf8_get_char(p string) {}

func Fn_utf8_get_char_validated(p string, maxLen uint64) {}

func Fn_utf8_normalize(str string, len uint64, mode int) {}

func Fn_utf8_offset_to_pointer(str string, offset int64) {}

func Fn_utf8_pointer_to_offset(str string, pos string) {}

func Fn_utf8_prev_char(p string) {}

func Fn_utf8_strchr(p string, len uint64, c rune) {}

func Fn_utf8_strdown(str string, len uint64) {}

func Fn_utf8_strlen(p string, max uint64) {}

func Fn_utf8_strncpy(dest string, src string, n uint64) {}

func Fn_utf8_strrchr(p string, len uint64, c rune) {}

func Fn_utf8_strreverse(str string, len uint64) {}

func Fn_utf8_strup(str string, len uint64) {}

func Fn_utf8_to_ucs4(str string, len int64, itemsRead *int64, itemsWritten *int64) {}

func Fn_utf8_to_ucs4_fast(str string, len int64, itemsWritten *int64) {}

func Fn_utf8_to_utf16(str string, len int64, itemsRead *int64, itemsWritten *int64) {}

func Fn_utf8_validate(str *uint8, maxLen uint64, end string) {}

func Fn_variant_get_gtype() {
	C.g_variant_get_gtype()
}

func Fn_variant_parse(type_ unsafe.Pointer, text string, limit string, endptr string) {}

func Fn_variant_parse_error_quark() {
	C.g_variant_parse_error_quark()
}

func Fn_variant_parser_get_error_quark() {
	C.g_variant_parser_get_error_quark()
}

func Fn_variant_type_checked_(arg0 string) {}

func Fn_variant_type_string_get_depth_(typeString string) {}

func Fn_variant_type_string_is_valid(typeString string) {}

// UNSUPPORTED : vasprintf : has va_list

// UNSUPPORTED : vfprintf : has va_list

// UNSUPPORTED : vprintf : has va_list

// UNSUPPORTED : vsnprintf : has va_list

// UNSUPPORTED : vsprintf : has va_list

func Fn_warn_message(domain string, file string, line int, func_ string, warnexpr string) {}

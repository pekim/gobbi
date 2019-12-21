// Code generated - DO NOT EDIT.
// +build glib_2.60

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
type LogWriterOutput C.GLogWriterOutput
type MarkupError C.GMarkupError
type NormalizeMode C.GNormalizeMode
type NumberParserError C.GNumberParserError
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

// UNSUPPORTED : TestResult : blacklisted
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
type LogField C.GLogField
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

// UNSUPPORTED : TestLogMsg : blacklisted
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
type VariantDict C.GVariantDict
type VariantIter C.GVariantIter
type VariantType C.GVariantType

// classes

// interfaces

func Fn_access(filename c.UndefinedParamType, mode int) {}

func Fn_ascii_digit_value(c c.UndefinedParamType) {}

func Fn_ascii_dtostr(buffer string, bufLen int, d float64) {}

func Fn_ascii_formatd(buffer string, bufLen int, format string, d float64) {}

func Fn_ascii_strcasecmp(s1 string, s2 string) {}

func Fn_ascii_strdown(str string, len uint64) {}

func Fn_ascii_string_to_signed(str string, base uint, min int64, max int64) {}

func Fn_ascii_string_to_unsigned(str string, base uint, min uint64, max uint64) {}

func Fn_ascii_strncasecmp(s1 string, s2 string, n uint64) {}

func Fn_ascii_strtod(nptr string) {}

func Fn_ascii_strtoll(nptr string, base uint) {}

func Fn_ascii_strtoull(nptr string, base uint) {}

func Fn_ascii_strup(str string, len uint64) {}

func Fn_ascii_tolower(c c.UndefinedParamType) {}

func Fn_ascii_toupper(c c.UndefinedParamType) {}

func Fn_ascii_xdigit_value(c c.UndefinedParamType) {}

func Fn_assert_warning(logDomain string, file string, line int, prettyFunction string, expression string) {
}

func Fn_assertion_message(domain string, file string, line int, func_ string, message string) {}

func Fn_assertion_message_cmpnum(domain string, file string, line int, func_ string, expr string, arg1 float64, cmp string, arg2 float64, numtype c.UndefinedParamType) {
}

func Fn_assertion_message_cmpstr(domain string, file string, line int, func_ string, expr string, arg1 string, cmp string, arg2 string) {
}

func Fn_assertion_message_error(domain string, file string, line int, func_ string, expr string, error c.UndefinedParamType, errorDomain c.UndefinedParamType, errorCode int) {
}

func Fn_assertion_message_expr(domain string, file string, line int, func_ string, expr string) {}

func Fn_atexit(func_ c.UndefinedParamType) {}

func Fn_atomic_int_add(atomic c.UndefinedParamType, val int) {}

func Fn_atomic_int_and(atomic c.UndefinedParamType, val uint) {}

func Fn_atomic_int_compare_and_exchange(atomic c.UndefinedParamType, oldval int, newval int) {}

func Fn_atomic_int_dec_and_test(atomic c.UndefinedParamType) {}

func Fn_atomic_int_exchange_and_add(atomic c.UndefinedParamType, val int) {}

func Fn_atomic_int_get(atomic c.UndefinedParamType) {}

func Fn_atomic_int_inc(atomic c.UndefinedParamType) {}

func Fn_atomic_int_or(atomic c.UndefinedParamType, val uint) {}

func Fn_atomic_int_set(atomic c.UndefinedParamType, newval int) {}

func Fn_atomic_int_xor(atomic c.UndefinedParamType, val uint) {}

func Fn_atomic_pointer_add(atomic c.UndefinedParamType, val uint64) {}

func Fn_atomic_pointer_and(atomic c.UndefinedParamType, val uint64) {}

func Fn_atomic_pointer_compare_and_exchange(atomic c.UndefinedParamType, oldval unsafe.Pointer, newval unsafe.Pointer) {
}

func Fn_atomic_pointer_get(atomic c.UndefinedParamType) {}

func Fn_atomic_pointer_or(atomic c.UndefinedParamType, val uint64) {}

func Fn_atomic_pointer_set(atomic c.UndefinedParamType, newval unsafe.Pointer) {}

func Fn_atomic_pointer_xor(atomic c.UndefinedParamType, val uint64) {}

func Fn_atomic_rc_box_acquire(memBlock unsafe.Pointer) {}

func Fn_atomic_rc_box_alloc(blockSize uint64) {}

func Fn_atomic_rc_box_alloc0(blockSize uint64) {}

func Fn_atomic_rc_box_dup(blockSize uint64, memBlock c.UndefinedParamType) {}

func Fn_atomic_rc_box_get_size(memBlock unsafe.Pointer) {}

func Fn_atomic_rc_box_release(memBlock unsafe.Pointer) {}

func Fn_atomic_rc_box_release_full(memBlock unsafe.Pointer, clearFunc c.UndefinedParamType) {}

func Fn_atomic_ref_count_compare(arc c.UndefinedParamType, val int) {}

func Fn_atomic_ref_count_dec(arc c.UndefinedParamType) {}

func Fn_atomic_ref_count_inc(arc c.UndefinedParamType) {}

func Fn_atomic_ref_count_init(arc c.UndefinedParamType) {}

func Fn_base64_decode(text string) {}

func Fn_base64_decode_inplace(text c.UndefinedParamType, outLen *uint64) {}

func Fn_base64_decode_step(in c.UndefinedParamType, len uint64, state *int, save *uint) {}

func Fn_base64_encode(data c.UndefinedParamType, len uint64) {}

func Fn_base64_encode_close(breakLines bool, state *int, save *int) {}

func Fn_base64_encode_step(in c.UndefinedParamType, len uint64, breakLines bool, state *int, save *int) {
}

func Fn_basename(fileName c.UndefinedParamType) {}

func Fn_bit_lock(address c.UndefinedParamType, lockBit int) {}

func Fn_bit_nth_lsf(mask uint64, nthBit int) {}

func Fn_bit_nth_msf(mask uint64, nthBit int) {}

func Fn_bit_storage(number uint64) {}

func Fn_bit_trylock(address c.UndefinedParamType, lockBit int) {}

func Fn_bit_unlock(address c.UndefinedParamType, lockBit int) {}

func Fn_bookmark_file_error_quark() {}

// UNSUPPORTED : build_filename : has varargs

func Fn_build_filename_valist(firstElement c.UndefinedParamType, args c.UndefinedParamType) {}

func Fn_build_filenamev(args c.UndefinedParamType) {}

// UNSUPPORTED : build_path : has varargs

func Fn_build_pathv(separator string, args c.UndefinedParamType) {}

func Fn_byte_array_free(array c.UndefinedParamType, freeSegment bool) {}

func Fn_byte_array_free_to_bytes(array c.UndefinedParamType) {}

func Fn_byte_array_new() {}

func Fn_byte_array_new_take(data c.UndefinedParamType, len uint64) {}

func Fn_byte_array_unref(array c.UndefinedParamType) {}

func Fn_canonicalize_filename(filename c.UndefinedParamType, relativeTo c.UndefinedParamType) {}

func Fn_chdir(path c.UndefinedParamType) {}

func Fn_check_version(requiredMajor uint, requiredMinor uint, requiredMicro uint) {}

func Fn_checksum_type_get_length(checksumType c.UndefinedParamType) {}

func Fn_child_watch_add(pid c.UndefinedParamType, function c.UndefinedParamType, data unsafe.Pointer) {
}

func Fn_child_watch_add_full(priority int, pid c.UndefinedParamType, function c.UndefinedParamType, data unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_child_watch_source_new(pid c.UndefinedParamType) {}

func Fn_clear_error() {}

func Fn_clear_handle_id(tagPtr *uint, clearFunc c.UndefinedParamType) {}

func Fn_clear_pointer(pp *unsafe.Pointer, destroy c.UndefinedParamType) {}

func Fn_close(fd int) {}

func Fn_compute_checksum_for_bytes(checksumType c.UndefinedParamType, data c.UndefinedParamType) {}

func Fn_compute_checksum_for_data(checksumType c.UndefinedParamType, data c.UndefinedParamType, length uint64) {
}

func Fn_compute_checksum_for_string(checksumType c.UndefinedParamType, str string, length uint64) {}

func Fn_compute_hmac_for_bytes(digestType c.UndefinedParamType, key c.UndefinedParamType, data c.UndefinedParamType) {
}

func Fn_compute_hmac_for_data(digestType c.UndefinedParamType, key c.UndefinedParamType, keyLen uint64, data c.UndefinedParamType, length uint64) {
}

func Fn_compute_hmac_for_string(digestType c.UndefinedParamType, key c.UndefinedParamType, keyLen uint64, str string, length uint64) {
}

func Fn_convert(str c.UndefinedParamType, len uint64, toCodeset string, fromCodeset string) {}

func Fn_convert_error_quark() {}

func Fn_convert_with_fallback(str c.UndefinedParamType, len uint64, toCodeset string, fromCodeset string, fallback string) {
}

func Fn_convert_with_iconv(str c.UndefinedParamType, len uint64, converter c.UndefinedParamType) {}

func Fn_datalist_clear(datalist c.UndefinedParamType) {}

func Fn_datalist_foreach(datalist c.UndefinedParamType, func_ c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_datalist_get_data(datalist c.UndefinedParamType, key string) {}

func Fn_datalist_get_flags(datalist c.UndefinedParamType) {}

func Fn_datalist_id_dup_data(datalist c.UndefinedParamType, keyId c.UndefinedParamType, dupFunc c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_datalist_id_get_data(datalist c.UndefinedParamType, keyId c.UndefinedParamType) {}

func Fn_datalist_id_remove_no_notify(datalist c.UndefinedParamType, keyId c.UndefinedParamType) {}

func Fn_datalist_id_replace_data(datalist c.UndefinedParamType, keyId c.UndefinedParamType, oldval unsafe.Pointer, newval unsafe.Pointer, destroy c.UndefinedParamType) {
}

func Fn_datalist_id_set_data_full(datalist c.UndefinedParamType, keyId c.UndefinedParamType, data unsafe.Pointer, destroyFunc c.UndefinedParamType) {
}

func Fn_datalist_init(datalist c.UndefinedParamType) {}

func Fn_datalist_set_flags(datalist c.UndefinedParamType, flags uint) {}

func Fn_datalist_unset_flags(datalist c.UndefinedParamType, flags uint) {}

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

func Fn_date_strftime(s string, slen uint64, format string, date c.UndefinedParamType) {}

func Fn_date_time_compare(dt1 c.UndefinedParamType, dt2 c.UndefinedParamType) {}

func Fn_date_time_equal(dt1 c.UndefinedParamType, dt2 c.UndefinedParamType) {}

func Fn_date_time_hash(datetime c.UndefinedParamType) {}

func Fn_date_valid_day(day c.UndefinedParamType) {}

func Fn_date_valid_dmy(day c.UndefinedParamType, month c.UndefinedParamType, year c.UndefinedParamType) {
}

func Fn_date_valid_julian(julianDate uint32) {}

func Fn_date_valid_month(month c.UndefinedParamType) {}

func Fn_date_valid_weekday(weekday c.UndefinedParamType) {}

func Fn_date_valid_year(year c.UndefinedParamType) {}

func Fn_dcgettext(domain string, msgid string, category int) {}

func Fn_dgettext(domain string, msgid string) {}

func Fn_dir_make_tmp(tmpl c.UndefinedParamType) {}

func Fn_direct_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_direct_hash(v c.UndefinedParamType) {}

func Fn_dngettext(domain string, msgid string, msgidPlural string, n uint64) {}

func Fn_double_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_double_hash(v c.UndefinedParamType) {}

func Fn_dpgettext(domain string, msgctxtid string, msgidoffset uint64) {}

func Fn_dpgettext2(domain string, context string, msgid string) {}

func Fn_environ_getenv(envp c.UndefinedParamType, variable c.UndefinedParamType) {}

func Fn_environ_setenv(envp c.UndefinedParamType, variable c.UndefinedParamType, value c.UndefinedParamType, overwrite bool) {
}

func Fn_environ_unsetenv(envp c.UndefinedParamType, variable c.UndefinedParamType) {}

func Fn_file_error_from_errno(errNo int) {}

func Fn_file_error_quark() {}

func Fn_file_get_contents(filename c.UndefinedParamType) {}

func Fn_file_open_tmp(tmpl c.UndefinedParamType) {}

func Fn_file_read_link(filename c.UndefinedParamType) {}

func Fn_file_set_contents(filename c.UndefinedParamType, contents c.UndefinedParamType, length uint64) {
}

func Fn_file_test(filename c.UndefinedParamType, test c.UndefinedParamType) {}

func Fn_filename_display_basename(filename c.UndefinedParamType) {}

func Fn_filename_display_name(filename c.UndefinedParamType) {}

func Fn_filename_from_uri(uri string) {}

func Fn_filename_from_utf8(utf8string string, len uint64) {}

func Fn_filename_to_uri(filename c.UndefinedParamType, hostname string) {}

func Fn_filename_to_utf8(opsysstring c.UndefinedParamType, len uint64) {}

func Fn_find_program_in_path(program c.UndefinedParamType) {}

func Fn_format_size(size uint64) {}

func Fn_format_size_for_display(size c.UndefinedParamType) {}

func Fn_format_size_full(size uint64, flags c.UndefinedParamType) {}

// UNSUPPORTED : fprintf : has varargs

func Fn_free(mem unsafe.Pointer) {}

func Fn_get_application_name() {}

func Fn_get_charset() {}

func Fn_get_codeset() {}

func Fn_get_current_dir() {}

func Fn_get_current_time(result c.UndefinedParamType) {}

func Fn_get_environ() {}

func Fn_get_filename_charsets() {}

func Fn_get_home_dir() {}

func Fn_get_host_name() {}

func Fn_get_language_names() {}

func Fn_get_language_names_with_category(categoryName string) {}

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

func Fn_get_user_special_dir(directory c.UndefinedParamType) {}

func Fn_getenv(variable c.UndefinedParamType) {}

func Fn_hash_table_add(hashTable c.UndefinedParamType, key unsafe.Pointer) {}

func Fn_hash_table_contains(hashTable c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_hash_table_destroy(hashTable c.UndefinedParamType) {}

func Fn_hash_table_insert(hashTable c.UndefinedParamType, key unsafe.Pointer, value unsafe.Pointer) {}

func Fn_hash_table_lookup(hashTable c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_hash_table_lookup_extended(hashTable c.UndefinedParamType, lookupKey c.UndefinedParamType) {}

func Fn_hash_table_remove(hashTable c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_hash_table_remove_all(hashTable c.UndefinedParamType) {}

func Fn_hash_table_replace(hashTable c.UndefinedParamType, key unsafe.Pointer, value unsafe.Pointer) {
}

func Fn_hash_table_size(hashTable c.UndefinedParamType) {}

func Fn_hash_table_steal(hashTable c.UndefinedParamType, key c.UndefinedParamType) {}

func Fn_hash_table_steal_all(hashTable c.UndefinedParamType) {}

func Fn_hash_table_steal_extended(hashTable c.UndefinedParamType, lookupKey c.UndefinedParamType) {}

func Fn_hash_table_unref(hashTable c.UndefinedParamType) {}

func Fn_hook_destroy(hookList c.UndefinedParamType, hookId uint64) {}

func Fn_hook_destroy_link(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_hook_free(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_hook_insert_before(hookList c.UndefinedParamType, sibling c.UndefinedParamType, hook c.UndefinedParamType) {
}

func Fn_hook_prepend(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_hook_unref(hookList c.UndefinedParamType, hook c.UndefinedParamType) {}

func Fn_hostname_is_ascii_encoded(hostname string) {}

func Fn_hostname_is_ip_address(hostname string) {}

func Fn_hostname_is_non_ascii(hostname string) {}

func Fn_hostname_to_ascii(hostname string) {}

func Fn_hostname_to_unicode(hostname string) {}

func Fn_iconv(converter c.UndefinedParamType, inbuf string, inbytesLeft *uint64, outbuf string, outbytesLeft *uint64) {
}

func Fn_iconv_open(toCodeset string, fromCodeset string) {}

func Fn_idle_add(function c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_idle_add_full(priority int, function c.UndefinedParamType, data unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_idle_remove_by_data(data unsafe.Pointer) {}

func Fn_idle_source_new() {}

func Fn_int64_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_int64_hash(v c.UndefinedParamType) {}

func Fn_int_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_int_hash(v c.UndefinedParamType) {}

func Fn_intern_static_string(string_ string) {}

func Fn_intern_string(string_ string) {}

func Fn_io_add_watch(channel c.UndefinedParamType, condition c.UndefinedParamType, func_ c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_io_add_watch_full(channel c.UndefinedParamType, priority int, condition c.UndefinedParamType, func_ c.UndefinedParamType, userData unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_io_channel_error_from_errno(en int) {}

func Fn_io_channel_error_quark() {}

func Fn_io_create_watch(channel c.UndefinedParamType, condition c.UndefinedParamType) {}

func Fn_key_file_error_quark() {}

func Fn_listenv() {}

func Fn_locale_from_utf8(utf8string string, len uint64) {}

func Fn_locale_to_utf8(opsysstring c.UndefinedParamType, len uint64) {}

// UNSUPPORTED : log : has varargs

func Fn_log_default_handler(logDomain string, logLevel c.UndefinedParamType, message string, unusedData unsafe.Pointer) {
}

func Fn_log_remove_handler(logDomain string, handlerId uint) {}

func Fn_log_set_always_fatal(fatalMask c.UndefinedParamType) {}

func Fn_log_set_default_handler(logFunc c.UndefinedParamType, userData unsafe.Pointer) {}

func Fn_log_set_fatal_mask(logDomain string, fatalMask c.UndefinedParamType) {}

func Fn_log_set_handler(logDomain string, logLevels c.UndefinedParamType, logFunc c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_log_set_handler_full(logDomain string, logLevels c.UndefinedParamType, logFunc c.UndefinedParamType, userData unsafe.Pointer, destroy c.UndefinedParamType) {
}

func Fn_log_set_writer_func(func_ c.UndefinedParamType, userData unsafe.Pointer, userDataFree c.UndefinedParamType) {
}

// UNSUPPORTED : log_structured : has varargs

func Fn_log_structured_array(logLevel c.UndefinedParamType, fields c.UndefinedParamType, nFields uint64) {
}

// UNSUPPORTED : log_structured_standard : has varargs

func Fn_log_variant(logDomain string, logLevel c.UndefinedParamType, fields c.UndefinedParamType) {}

func Fn_log_writer_default(logLevel c.UndefinedParamType, fields c.UndefinedParamType, nFields uint64, userData unsafe.Pointer) {
}

func Fn_log_writer_format_fields(logLevel c.UndefinedParamType, fields c.UndefinedParamType, nFields uint64, useColor bool) {
}

func Fn_log_writer_is_journald(outputFd int) {}

func Fn_log_writer_journald(logLevel c.UndefinedParamType, fields c.UndefinedParamType, nFields uint64, userData unsafe.Pointer) {
}

func Fn_log_writer_standard_streams(logLevel c.UndefinedParamType, fields c.UndefinedParamType, nFields uint64, userData unsafe.Pointer) {
}

func Fn_log_writer_supports_color(outputFd int) {}

func Fn_logv(logDomain string, logLevel c.UndefinedParamType, format string, args c.UndefinedParamType) {
}

func Fn_main_context_default() {}

func Fn_main_context_get_thread_default() {}

func Fn_main_context_ref_thread_default() {}

func Fn_main_current_source() {}

func Fn_main_depth() {}

func Fn_malloc(nBytes uint64) {}

func Fn_malloc0(nBytes uint64) {}

func Fn_malloc0_n(nBlocks uint64, nBlockBytes uint64) {}

func Fn_malloc_n(nBlocks uint64, nBlockBytes uint64) {}

// UNSUPPORTED : markup_collect_attributes : has varargs

func Fn_markup_error_quark() {}

func Fn_markup_escape_text(text string, length uint64) {}

// UNSUPPORTED : markup_printf_escaped : has varargs

func Fn_markup_vprintf_escaped(format string, args c.UndefinedParamType) {}

func Fn_mem_is_system_malloc() {}

func Fn_mem_profile() {}

func Fn_mem_set_vtable(vtable c.UndefinedParamType) {}

func Fn_memdup(mem c.UndefinedParamType, byteSize uint) {}

func Fn_mkdir_with_parents(pathname c.UndefinedParamType, mode int) {}

func Fn_mkdtemp(tmpl c.UndefinedParamType) {}

func Fn_mkdtemp_full(tmpl c.UndefinedParamType, mode int) {}

func Fn_mkstemp(tmpl c.UndefinedParamType) {}

func Fn_mkstemp_full(tmpl c.UndefinedParamType, flags int, mode int) {}

func Fn_nullify_pointer(nullifyLocation *unsafe.Pointer) {}

func Fn_number_parser_error_quark() {}

func Fn_on_error_query(prgName string) {}

func Fn_on_error_stack_trace(prgName string) {}

func Fn_once_init_enter(location c.UndefinedParamType) {}

func Fn_once_init_leave(location c.UndefinedParamType, result uint64) {}

func Fn_option_error_quark() {}

func Fn_parse_debug_string(string_ string, keys c.UndefinedParamType, nkeys uint) {}

func Fn_path_get_basename(fileName c.UndefinedParamType) {}

func Fn_path_get_dirname(fileName c.UndefinedParamType) {}

func Fn_path_is_absolute(fileName c.UndefinedParamType) {}

func Fn_path_skip_root(fileName c.UndefinedParamType) {}

func Fn_pattern_match(pspec c.UndefinedParamType, stringLength uint, string_ string, stringReversed string) {
}

func Fn_pattern_match_simple(pattern string, string_ string) {}

func Fn_pattern_match_string(pspec c.UndefinedParamType, string_ string) {}

func Fn_pointer_bit_lock(address c.UndefinedParamType, lockBit int) {}

func Fn_pointer_bit_trylock(address c.UndefinedParamType, lockBit int) {}

func Fn_pointer_bit_unlock(address c.UndefinedParamType, lockBit int) {}

func Fn_poll(fds c.UndefinedParamType, nfds uint, timeout int) {}

// UNSUPPORTED : prefix_error : has varargs

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

// UNSUPPORTED : printf : has varargs

func Fn_printf_string_upper_bound(format string, args c.UndefinedParamType) {}

func Fn_propagate_error(src c.UndefinedParamType) {}

// UNSUPPORTED : propagate_prefixed_error : has varargs

func Fn_ptr_array_find(haystack c.UndefinedParamType, needle c.UndefinedParamType) {}

func Fn_ptr_array_find_with_equal_func(haystack c.UndefinedParamType, needle c.UndefinedParamType, equalFunc c.UndefinedParamType) {
}

func Fn_qsort_with_data(pbase c.UndefinedParamType, totalElems int, size uint64, compareFunc c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_quark_from_static_string(string_ string) {}

func Fn_quark_from_string(string_ string) {}

func Fn_quark_to_string(quark c.UndefinedParamType) {}

func Fn_quark_try_string(string_ string) {}

func Fn_random_double() {}

func Fn_random_double_range(begin float64, end float64) {}

func Fn_random_int() {}

func Fn_random_int_range(begin int32, end int32) {}

func Fn_random_set_seed(seed uint32) {}

func Fn_rc_box_acquire(memBlock unsafe.Pointer) {}

func Fn_rc_box_alloc(blockSize uint64) {}

func Fn_rc_box_alloc0(blockSize uint64) {}

func Fn_rc_box_dup(blockSize uint64, memBlock c.UndefinedParamType) {}

func Fn_rc_box_get_size(memBlock unsafe.Pointer) {}

func Fn_rc_box_release(memBlock unsafe.Pointer) {}

func Fn_rc_box_release_full(memBlock unsafe.Pointer, clearFunc c.UndefinedParamType) {}

func Fn_realloc(mem unsafe.Pointer, nBytes uint64) {}

func Fn_realloc_n(mem unsafe.Pointer, nBlocks uint64, nBlockBytes uint64) {}

func Fn_ref_count_compare(rc c.UndefinedParamType, val int) {}

func Fn_ref_count_dec(rc c.UndefinedParamType) {}

func Fn_ref_count_inc(rc c.UndefinedParamType) {}

func Fn_ref_count_init(rc c.UndefinedParamType) {}

func Fn_ref_string_acquire(str string) {}

func Fn_ref_string_length(str string) {}

func Fn_ref_string_new(str string) {}

func Fn_ref_string_new_intern(str string) {}

func Fn_ref_string_new_len(str string, len uint64) {}

func Fn_ref_string_release(str string) {}

func Fn_regex_check_replacement(replacement string) {}

func Fn_regex_error_quark() {}

func Fn_regex_escape_nul(string_ string, length int) {}

func Fn_regex_escape_string(string_ c.UndefinedParamType, length int) {}

func Fn_regex_match_simple(pattern string, string_ string, compileOptions c.UndefinedParamType, matchOptions c.UndefinedParamType) {
}

func Fn_regex_split_simple(pattern string, string_ string, compileOptions c.UndefinedParamType, matchOptions c.UndefinedParamType) {
}

func Fn_reload_user_special_dirs_cache() {}

func Fn_return_if_fail_warning(logDomain string, prettyFunction string, expression string) {}

func Fn_rmdir(filename c.UndefinedParamType) {}

func Fn_sequence_get(iter c.UndefinedParamType) {}

func Fn_sequence_insert_before(iter c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_sequence_move(src c.UndefinedParamType, dest c.UndefinedParamType) {}

func Fn_sequence_move_range(dest c.UndefinedParamType, begin c.UndefinedParamType, end c.UndefinedParamType) {
}

func Fn_sequence_range_get_midpoint(begin c.UndefinedParamType, end c.UndefinedParamType) {}

func Fn_sequence_remove(iter c.UndefinedParamType) {}

func Fn_sequence_remove_range(begin c.UndefinedParamType, end c.UndefinedParamType) {}

func Fn_sequence_set(iter c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_sequence_swap(a c.UndefinedParamType, b c.UndefinedParamType) {}

func Fn_set_application_name(applicationName string) {}

// UNSUPPORTED : set_error : has varargs

func Fn_set_error_literal(domain c.UndefinedParamType, code int, message string) {}

func Fn_set_prgname(prgname string) {}

func Fn_set_print_handler(func_ c.UndefinedParamType) {}

func Fn_set_printerr_handler(func_ c.UndefinedParamType) {}

func Fn_setenv(variable c.UndefinedParamType, value c.UndefinedParamType, overwrite bool) {}

func Fn_shell_error_quark() {}

func Fn_shell_parse_argv(commandLine c.UndefinedParamType) {}

func Fn_shell_quote(unquotedString c.UndefinedParamType) {}

func Fn_shell_unquote(quotedString c.UndefinedParamType) {}

func Fn_slice_alloc(blockSize uint64) {}

func Fn_slice_alloc0(blockSize uint64) {}

func Fn_slice_copy(blockSize uint64, memBlock c.UndefinedParamType) {}

func Fn_slice_free1(blockSize uint64, memBlock unsafe.Pointer) {}

func Fn_slice_free_chain_with_offset(blockSize uint64, memChain unsafe.Pointer, nextOffset uint64) {}

func Fn_slice_get_config(ckey c.UndefinedParamType) {}

func Fn_slice_get_config_state(ckey c.UndefinedParamType, address int64, nValues *uint) {}

func Fn_slice_set_config(ckey c.UndefinedParamType, value int64) {}

// UNSUPPORTED : snprintf : has varargs

func Fn_source_remove(tag uint) {}

func Fn_source_remove_by_funcs_user_data(funcs c.UndefinedParamType, userData unsafe.Pointer) {}

func Fn_source_remove_by_user_data(userData unsafe.Pointer) {}

func Fn_source_set_name_by_id(tag uint, name string) {}

func Fn_spaced_primes_closest(num uint) {}

func Fn_spawn_async(workingDirectory c.UndefinedParamType, argv c.UndefinedParamType, envp c.UndefinedParamType, flags c.UndefinedParamType, childSetup c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_spawn_async_with_fds(workingDirectory c.UndefinedParamType, argv c.UndefinedParamType, envp c.UndefinedParamType, flags c.UndefinedParamType, childSetup c.UndefinedParamType, userData unsafe.Pointer, stdinFd int, stdoutFd int, stderrFd int) {
}

func Fn_spawn_async_with_pipes(workingDirectory c.UndefinedParamType, argv c.UndefinedParamType, envp c.UndefinedParamType, flags c.UndefinedParamType, childSetup c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_spawn_check_exit_status(exitStatus int) {}

func Fn_spawn_close_pid(pid c.UndefinedParamType) {}

func Fn_spawn_command_line_async(commandLine c.UndefinedParamType) {}

func Fn_spawn_command_line_sync(commandLine c.UndefinedParamType) {}

func Fn_spawn_error_quark() {}

func Fn_spawn_exit_error_quark() {}

func Fn_spawn_sync(workingDirectory c.UndefinedParamType, argv c.UndefinedParamType, envp c.UndefinedParamType, flags c.UndefinedParamType, childSetup c.UndefinedParamType, userData unsafe.Pointer) {
}

// UNSUPPORTED : sprintf : has varargs

func Fn_stpcpy(dest string, src string) {}

func Fn_str_equal(v1 c.UndefinedParamType, v2 c.UndefinedParamType) {}

func Fn_str_has_prefix(str string, prefix string) {}

func Fn_str_has_suffix(str string, suffix string) {}

func Fn_str_hash(v c.UndefinedParamType) {}

func Fn_str_is_ascii(str string) {}

func Fn_str_match_string(searchTerm string, potentialHit string, acceptAlternates bool) {}

func Fn_str_to_ascii(str string, fromLocale string) {}

func Fn_str_tokenize_and_fold(string_ string, translitLocale string) {}

func Fn_strcanon(string_ string, validChars string, substitutor c.UndefinedParamType) {}

func Fn_strcasecmp(s1 string, s2 string) {}

func Fn_strchomp(string_ string) {}

func Fn_strchug(string_ string) {}

func Fn_strcmp0(str1 string, str2 string) {}

func Fn_strcompress(source string) {}

// UNSUPPORTED : strconcat : has varargs

func Fn_strdelimit(string_ string, delimiters string, newDelimiter c.UndefinedParamType) {}

func Fn_strdown(string_ string) {}

func Fn_strdup(str string) {}

// UNSUPPORTED : strdup_printf : has varargs

func Fn_strdup_vprintf(format string, args c.UndefinedParamType) {}

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

func Fn_strnfill(length uint64, fillChar c.UndefinedParamType) {}

func Fn_strreverse(string_ string) {}

func Fn_strrstr(haystack string, needle string) {}

func Fn_strrstr_len(haystack string, haystackLen uint64, needle string) {}

func Fn_strsignal(signum int) {}

func Fn_strsplit(string_ string, delimiter string, maxTokens int) {}

func Fn_strsplit_set(string_ string, delimiters string, maxTokens int) {}

func Fn_strstr_len(haystack string, haystackLen uint64, needle string) {}

func Fn_strtod(nptr string) {}

func Fn_strup(string_ string) {}

func Fn_strv_contains(strv string, str string) {}

func Fn_strv_equal(strv1 string, strv2 string) {}

func Fn_strv_get_type() {}

func Fn_strv_length(strArray string) {}

func Fn_test_add_data_func(testpath string, testData c.UndefinedParamType, testFunc c.UndefinedParamType) {
}

func Fn_test_add_data_func_full(testpath string, testData unsafe.Pointer, testFunc c.UndefinedParamType, dataFreeFunc c.UndefinedParamType) {
}

func Fn_test_add_func(testpath string, testFunc c.UndefinedParamType) {}

func Fn_test_add_vtable(testpath string, dataSize uint64, testData c.UndefinedParamType, dataSetup c.UndefinedParamType, dataTest c.UndefinedParamType, dataTeardown c.UndefinedParamType) {
}

func Fn_test_assert_expected_messages_internal(domain string, file string, line int, func_ string) {}

func Fn_test_bug(bugUriSnippet string) {}

func Fn_test_bug_base(uriPattern string) {}

// UNSUPPORTED : test_build_filename : has varargs

func Fn_test_create_case(testName string, dataSize uint64, testData c.UndefinedParamType, dataSetup c.UndefinedParamType, dataTest c.UndefinedParamType, dataTeardown c.UndefinedParamType) {
}

func Fn_test_create_suite(suiteName string) {}

func Fn_test_expect_message(logDomain string, logLevel c.UndefinedParamType, pattern string) {}

func Fn_test_fail() {}

func Fn_test_failed() {}

func Fn_test_get_dir(fileType c.UndefinedParamType) {}

// UNSUPPORTED : test_get_filename : has varargs

func Fn_test_get_root() {}

func Fn_test_incomplete(msg string) {}

// UNSUPPORTED : test_init : has varargs

func Fn_test_log_set_fatal_handler(logFunc c.UndefinedParamType, userData unsafe.Pointer) {}

func Fn_test_log_type_name(logType c.UndefinedParamType) {}

// UNSUPPORTED : test_maximized_result : has varargs

// UNSUPPORTED : test_message : has varargs

// UNSUPPORTED : test_minimized_result : has varargs

func Fn_test_queue_destroy(destroyFunc c.UndefinedParamType, destroyData unsafe.Pointer) {}

func Fn_test_queue_free(gfreePointer unsafe.Pointer) {}

func Fn_test_rand_double() {}

func Fn_test_rand_double_range(rangeStart float64, rangeEnd float64) {}

func Fn_test_rand_int() {}

func Fn_test_rand_int_range(begin int32, end int32) {}

func Fn_test_run() {}

func Fn_test_run_suite(suite c.UndefinedParamType) {}

func Fn_test_set_nonfatal_assertions() {}

func Fn_test_skip(msg string) {}

func Fn_test_subprocess() {}

func Fn_test_timer_elapsed() {}

func Fn_test_timer_last() {}

func Fn_test_timer_start() {}

func Fn_test_trap_assertions(domain string, file string, line int, func_ string, assertionFlags uint64, pattern string) {
}

func Fn_test_trap_fork(usecTimeout uint64, testTrapFlags c.UndefinedParamType) {}

func Fn_test_trap_has_passed() {}

func Fn_test_trap_reached_timeout() {}

func Fn_test_trap_subprocess(testPath string, usecTimeout uint64, testFlags c.UndefinedParamType) {}

func Fn_thread_error_quark() {}

func Fn_thread_exit(retval unsafe.Pointer) {}

func Fn_thread_pool_get_max_idle_time() {}

func Fn_thread_pool_get_max_unused_threads() {}

func Fn_thread_pool_get_num_unused_threads() {}

func Fn_thread_pool_set_max_idle_time(interval uint) {}

func Fn_thread_pool_set_max_unused_threads(maxThreads int) {}

func Fn_thread_pool_stop_unused_threads() {}

func Fn_thread_self() {}

func Fn_thread_yield() {}

func Fn_time_val_from_iso8601(isoDate string) {}

func Fn_timeout_add(interval uint, function c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_timeout_add_full(priority int, interval uint, function c.UndefinedParamType, data unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_timeout_add_seconds(interval uint, function c.UndefinedParamType, data unsafe.Pointer) {}

func Fn_timeout_add_seconds_full(priority int, interval uint, function c.UndefinedParamType, data unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_timeout_source_new(interval uint) {}

func Fn_timeout_source_new_seconds(interval uint) {}

func Fn_trash_stack_height(stackP c.UndefinedParamType) {}

func Fn_trash_stack_peek(stackP c.UndefinedParamType) {}

func Fn_trash_stack_pop(stackP c.UndefinedParamType) {}

func Fn_trash_stack_push(stackP c.UndefinedParamType, dataP unsafe.Pointer) {}

func Fn_try_malloc(nBytes uint64) {}

func Fn_try_malloc0(nBytes uint64) {}

func Fn_try_malloc0_n(nBlocks uint64, nBlockBytes uint64) {}

func Fn_try_malloc_n(nBlocks uint64, nBlockBytes uint64) {}

func Fn_try_realloc(mem unsafe.Pointer, nBytes uint64) {}

func Fn_try_realloc_n(mem unsafe.Pointer, nBlocks uint64, nBlockBytes uint64) {}

func Fn_ucs4_to_utf16(str *rune, len int64) {}

func Fn_ucs4_to_utf8(str *rune, len int64) {}

func Fn_unichar_break_type(c rune) {}

func Fn_unichar_combining_class(uc rune) {}

func Fn_unichar_compose(a rune, b rune) {}

func Fn_unichar_decompose(ch rune) {}

func Fn_unichar_digit_value(c rune) {}

func Fn_unichar_fully_decompose(ch rune, compat bool, resultLen uint64) {}

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

func Fn_unichar_to_utf8(c rune) {}

func Fn_unichar_tolower(c rune) {}

func Fn_unichar_totitle(c rune) {}

func Fn_unichar_toupper(c rune) {}

func Fn_unichar_type(c rune) {}

func Fn_unichar_validate(ch rune) {}

func Fn_unichar_xdigit_value(c rune) {}

func Fn_unicode_canonical_decomposition(ch rune, resultLen *uint64) {}

func Fn_unicode_canonical_ordering(string_ *rune, len uint64) {}

func Fn_unicode_script_from_iso15924(iso15924 uint32) {}

func Fn_unicode_script_to_iso15924(script c.UndefinedParamType) {}

func Fn_unix_error_quark() {}

func Fn_unix_fd_add(fd int, condition c.UndefinedParamType, function c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_unix_fd_add_full(priority int, fd int, condition c.UndefinedParamType, function c.UndefinedParamType, userData unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_unix_fd_source_new(fd int, condition c.UndefinedParamType) {}

func Fn_unix_open_pipe(fds *int, flags int) {}

func Fn_unix_set_fd_nonblocking(fd int, nonblock bool) {}

func Fn_unix_signal_add(signum int, handler c.UndefinedParamType, userData unsafe.Pointer) {}

func Fn_unix_signal_add_full(priority int, signum int, handler c.UndefinedParamType, userData unsafe.Pointer, notify c.UndefinedParamType) {
}

func Fn_unix_signal_source_new(signum int) {}

func Fn_unlink(filename c.UndefinedParamType) {}

func Fn_unsetenv(variable c.UndefinedParamType) {}

func Fn_uri_escape_string(unescaped string, reservedCharsAllowed string, allowUtf8 bool) {}

func Fn_uri_list_extract_uris(uriList string) {}

func Fn_uri_parse_scheme(uri string) {}

func Fn_uri_unescape_segment(escapedString string, escapedStringEnd string, illegalCharacters string) {}

func Fn_uri_unescape_string(escapedString string, illegalCharacters string) {}

func Fn_usleep(microseconds uint64) {}

func Fn_utf16_to_ucs4(str c.UndefinedParamType, len int64) {}

func Fn_utf16_to_utf8(str c.UndefinedParamType, len int64) {}

func Fn_utf8_casefold(str string, len uint64) {}

func Fn_utf8_collate(str1 string, str2 string) {}

func Fn_utf8_collate_key(str string, len uint64) {}

func Fn_utf8_collate_key_for_filename(str string, len uint64) {}

func Fn_utf8_find_next_char(p string, end string) {}

func Fn_utf8_find_prev_char(str string, p string) {}

func Fn_utf8_get_char(p string) {}

func Fn_utf8_get_char_validated(p string, maxLen uint64) {}

func Fn_utf8_make_valid(str string, len uint64) {}

func Fn_utf8_normalize(str string, len uint64, mode c.UndefinedParamType) {}

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

func Fn_utf8_substring(str string, startPos int64, endPos int64) {}

func Fn_utf8_to_ucs4(str string, len int64) {}

func Fn_utf8_to_ucs4_fast(str string, len int64) {}

func Fn_utf8_to_utf16(str string, len int64) {}

func Fn_utf8_validate(str c.UndefinedParamType, maxLen uint64) {}

func Fn_utf8_validate_len(str c.UndefinedParamType, maxLen uint64) {}

func Fn_uuid_string_is_valid(str string) {}

func Fn_uuid_string_random() {}

func Fn_variant_get_gtype() {}

func Fn_variant_is_object_path(string_ string) {}

func Fn_variant_is_signature(string_ string) {}

func Fn_variant_parse(type_ c.UndefinedParamType, text string, limit string, endptr string) {}

func Fn_variant_parse_error_print_context(error c.UndefinedParamType, sourceStr string) {}

func Fn_variant_parse_error_quark() {}

func Fn_variant_parser_get_error_quark() {}

func Fn_variant_type_checked_(arg0 string) {}

func Fn_variant_type_string_get_depth_(typeString string) {}

func Fn_variant_type_string_is_valid(typeString string) {}

func Fn_variant_type_string_scan(string_ string, limit string) {}

func Fn_vasprintf(string_ string, format string, args c.UndefinedParamType) {}

func Fn_vfprintf(file c.UndefinedParamType, format string, args c.UndefinedParamType) {}

func Fn_vprintf(format string, args c.UndefinedParamType) {}

func Fn_vsnprintf(string_ string, n uint64, format string, args c.UndefinedParamType) {}

func Fn_vsprintf(string_ string, format string, args c.UndefinedParamType) {}

func Fn_warn_message(domain string, file string, line int, func_ string, warnexpr string) {}

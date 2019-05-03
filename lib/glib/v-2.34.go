// Code generated - DO NOT EDIT.
// +build glib_2.34

package glib

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	static void _g_log(const gchar* log_domain, GLogLevelFlags log_level, const gchar* format) {
		return g_log(log_domain, log_level, format);
    }
*/
/*

	static gchar* _g_markup_printf_escaped(const char* format) {
		return g_markup_printf_escaped(format);
    }
*/
/*

	static void _g_prefix_error(GError** err, const gchar* format) {
		return g_prefix_error(err, format);
    }
*/
/*

	static void _g_print(const gchar* format) {
		return g_print(format);
    }
*/
/*

	static void _g_printerr(const gchar* format) {
		return g_printerr(format);
    }
*/
/*

	static gint _g_printf(const gchar* format) {
		return g_printf(format);
    }
*/
/*

	static void _g_propagate_prefixed_error(GError** dest, GError* src, const gchar* format) {
		return g_propagate_prefixed_error(dest, src, format);
    }
*/
/*

	static void _g_set_error(GError** err, GQuark domain, gint code, const gchar* format) {
		return g_set_error(err, domain, code, format);
    }
*/
/*

	static gint _g_snprintf(gchar* string, gulong n, const gchar* format) {
		return g_snprintf(string, n, format);
    }
*/
/*

	static gint _g_sprintf(gchar* string, const gchar* format) {
		return g_sprintf(string, format);
    }
*/
/*

	static gchar* _g_strdup_printf(const gchar* format) {
		return g_strdup_printf(format);
    }
*/
/*

	static void _g_test_maximized_result(double maximized_quantity, const char* format) {
		return g_test_maximized_result(maximized_quantity, format);
    }
*/
/*

	static void _g_test_message(const char* format) {
		return g_test_message(format);
    }
*/
/*

	static void _g_test_minimized_result(double minimized_quantity, const char* format) {
		return g_test_minimized_result(minimized_quantity, format);
    }
*/
/*

	static GError* _g_error_new(GQuark domain, gint code, const gchar* format) {
		return g_error_new(domain, code, format);
    }
*/
/*

	static void _g_scanner_error(GScanner* scanner, const gchar* format) {
		return g_scanner_error(scanner, format);
    }
*/
/*

	static void _g_scanner_warn(GScanner* scanner, const gchar* format) {
		return g_scanner_warn(scanner, format);
    }
*/
/*

	static void _g_string_append_printf(GString* string, const gchar* format) {
		return g_string_append_printf(string, format);
    }
*/
/*

	static void _g_string_printf(GString* string, const gchar* format) {
		return g_string_printf(string, format);
    }
*/
/*

	static GVariant* _g_variant_new(const gchar* format_string) {
		return g_variant_new(format_string);
    }
*/
/*

	static GVariant* _g_variant_new_parsed(const gchar* format) {
		return g_variant_new_parsed(format);
    }
*/
/*

	static void _g_variant_get(GVariant* value, const gchar* format_string) {
		return g_variant_get(value, format_string);
    }
*/
/*

	static void _g_variant_get_child(GVariant* value, gsize index_, const gchar* format_string) {
		return g_variant_get_child(value, index_, format_string);
    }
*/
/*

	static gboolean _g_variant_lookup(GVariant* dictionary, const gchar* key, const gchar* format_string) {
		return g_variant_lookup(dictionary, key, format_string);
    }
*/
/*

	static void _g_variant_builder_add(GVariantBuilder* builder, const gchar* format_string) {
		return g_variant_builder_add(builder, format_string);
    }
*/
/*

	static void _g_variant_builder_add_parsed(GVariantBuilder* builder, const gchar* format) {
		return g_variant_builder_add_parsed(builder, format);
    }
*/
/*

	static gboolean _g_variant_iter_loop(GVariantIter* iter, const gchar* format_string) {
		return g_variant_iter_loop(iter, format_string);
    }
*/
/*

	static gboolean _g_variant_iter_next(GVariantIter* iter, const gchar* format_string) {
		return g_variant_iter_next(iter, format_string);
    }
*/
import "C"

// DateDay is a representation of the C alias GDateDay.
type DateDay uint8

// DateYear is a representation of the C alias GDateYear.
type DateYear uint16

// Unsupported : alias has no type generator for none, void

// Pid is a representation of the C alias GPid.
type Pid int32

// Quark is a representation of the C alias GQuark.
type Quark uint32

// Strv is a representation of the C alias GStrv.
type Strv string

// Time is a representation of the C alias GTime.
type Time int32

// TimeSpan is a representation of the C alias GTimeSpan.
type TimeSpan int64

// Type is a representation of the C alias GType.
type Type uint64

type AsciiType C.GAsciiType

const (
	ASCII_ALNUM  AsciiType = 1
	ASCII_ALPHA  AsciiType = 2
	ASCII_CNTRL  AsciiType = 4
	ASCII_DIGIT  AsciiType = 8
	ASCII_GRAPH  AsciiType = 16
	ASCII_LOWER  AsciiType = 32
	ASCII_PRINT  AsciiType = 64
	ASCII_PUNCT  AsciiType = 128
	ASCII_SPACE  AsciiType = 256
	ASCII_UPPER  AsciiType = 512
	ASCII_XDIGIT AsciiType = 1024
)

type GFileTest C.GFileTest

const (
	FILE_TEST_IS_REGULAR    GFileTest = 1
	FILE_TEST_IS_SYMLINK    GFileTest = 2
	FILE_TEST_IS_DIR        GFileTest = 4
	FILE_TEST_IS_EXECUTABLE GFileTest = 8
	FILE_TEST_EXISTS        GFileTest = 16
)

type FormatSizeFlags C.GFormatSizeFlags

const (
	FORMAT_SIZE_DEFAULT     FormatSizeFlags = 0
	FORMAT_SIZE_LONG_FORMAT FormatSizeFlags = 1
	FORMAT_SIZE_IEC_UNITS   FormatSizeFlags = 2
	FORMAT_SIZE_BITS        FormatSizeFlags = 4
)

type HookFlagMask C.GHookFlagMask

const (
	HOOK_FLAG_ACTIVE  HookFlagMask = 1
	HOOK_FLAG_IN_CALL HookFlagMask = 2
	HOOK_FLAG_MASK    HookFlagMask = 15
)

type IOCondition C.GIOCondition

const (
	IO_IN   IOCondition = 1
	IO_OUT  IOCondition = 4
	IO_PRI  IOCondition = 2
	IO_ERR  IOCondition = 8
	IO_HUP  IOCondition = 16
	IO_NVAL IOCondition = 32
)

type IOFlags C.GIOFlags

const (
	IO_FLAG_APPEND       IOFlags = 1
	IO_FLAG_NONBLOCK     IOFlags = 2
	IO_FLAG_IS_READABLE  IOFlags = 4
	IO_FLAG_IS_WRITABLE  IOFlags = 8
	IO_FLAG_IS_WRITEABLE IOFlags = 8
	IO_FLAG_IS_SEEKABLE  IOFlags = 16
	IO_FLAG_MASK         IOFlags = 31
	IO_FLAG_GET_MASK     IOFlags = 31
	IO_FLAG_SET_MASK     IOFlags = 3
)

type KeyFileFlags C.GKeyFileFlags

const (
	KEY_FILE_NONE              KeyFileFlags = 0
	KEY_FILE_KEEP_COMMENTS     KeyFileFlags = 1
	KEY_FILE_KEEP_TRANSLATIONS KeyFileFlags = 2
)

type LogLevelFlags C.GLogLevelFlags

const (
	LOG_FLAG_RECURSION LogLevelFlags = 1
	LOG_FLAG_FATAL     LogLevelFlags = 2
	LOG_LEVEL_ERROR    LogLevelFlags = 4
	LOG_LEVEL_CRITICAL LogLevelFlags = 8
	LOG_LEVEL_WARNING  LogLevelFlags = 16
	LOG_LEVEL_MESSAGE  LogLevelFlags = 32
	LOG_LEVEL_INFO     LogLevelFlags = 64
	LOG_LEVEL_DEBUG    LogLevelFlags = 128
	LOG_LEVEL_MASK     LogLevelFlags = -4
)

type MarkupCollectType C.GMarkupCollectType

const (
	MARKUP_COLLECT_INVALID  MarkupCollectType = 0
	MARKUP_COLLECT_STRING   MarkupCollectType = 1
	MARKUP_COLLECT_STRDUP   MarkupCollectType = 2
	MARKUP_COLLECT_BOOLEAN  MarkupCollectType = 3
	MARKUP_COLLECT_TRISTATE MarkupCollectType = 4
	MARKUP_COLLECT_OPTIONAL MarkupCollectType = 65536
)

type MarkupParseFlags C.GMarkupParseFlags

const (
	MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG MarkupParseFlags = 1
	MARKUP_TREAT_CDATA_AS_TEXT              MarkupParseFlags = 2
	MARKUP_PREFIX_ERROR_POSITION            MarkupParseFlags = 4
	MARKUP_IGNORE_QUALIFIED                 MarkupParseFlags = 8
)

type OptionFlags C.GOptionFlags

const (
	OPTION_FLAG_NONE         OptionFlags = 0
	OPTION_FLAG_HIDDEN       OptionFlags = 1
	OPTION_FLAG_IN_MAIN      OptionFlags = 2
	OPTION_FLAG_REVERSE      OptionFlags = 4
	OPTION_FLAG_NO_ARG       OptionFlags = 8
	OPTION_FLAG_FILENAME     OptionFlags = 16
	OPTION_FLAG_OPTIONAL_ARG OptionFlags = 32
	OPTION_FLAG_NOALIAS      OptionFlags = 64
)

type RegexCompileFlags C.GRegexCompileFlags

const (
	REGEX_CASELESS          RegexCompileFlags = 1
	REGEX_MULTILINE         RegexCompileFlags = 2
	REGEX_DOTALL            RegexCompileFlags = 4
	REGEX_EXTENDED          RegexCompileFlags = 8
	REGEX_ANCHORED          RegexCompileFlags = 16
	REGEX_DOLLAR_ENDONLY    RegexCompileFlags = 32
	REGEX_UNGREEDY          RegexCompileFlags = 512
	REGEX_RAW               RegexCompileFlags = 2048
	REGEX_NO_AUTO_CAPTURE   RegexCompileFlags = 4096
	REGEX_OPTIMIZE          RegexCompileFlags = 8192
	REGEX_FIRSTLINE         RegexCompileFlags = 262144
	REGEX_DUPNAMES          RegexCompileFlags = 524288
	REGEX_NEWLINE_CR        RegexCompileFlags = 1048576
	REGEX_NEWLINE_LF        RegexCompileFlags = 2097152
	REGEX_NEWLINE_CRLF      RegexCompileFlags = 3145728
	REGEX_NEWLINE_ANYCRLF   RegexCompileFlags = 5242880
	REGEX_BSR_ANYCRLF       RegexCompileFlags = 8388608
	REGEX_JAVASCRIPT_COMPAT RegexCompileFlags = 33554432
)

type RegexMatchFlags C.GRegexMatchFlags

const (
	REGEX_MATCH_ANCHORED         RegexMatchFlags = 16
	REGEX_MATCH_NOTBOL           RegexMatchFlags = 128
	REGEX_MATCH_NOTEOL           RegexMatchFlags = 256
	REGEX_MATCH_NOTEMPTY         RegexMatchFlags = 1024
	REGEX_MATCH_PARTIAL          RegexMatchFlags = 32768
	REGEX_MATCH_NEWLINE_CR       RegexMatchFlags = 1048576
	REGEX_MATCH_NEWLINE_LF       RegexMatchFlags = 2097152
	REGEX_MATCH_NEWLINE_CRLF     RegexMatchFlags = 3145728
	REGEX_MATCH_NEWLINE_ANY      RegexMatchFlags = 4194304
	REGEX_MATCH_NEWLINE_ANYCRLF  RegexMatchFlags = 5242880
	REGEX_MATCH_BSR_ANYCRLF      RegexMatchFlags = 8388608
	REGEX_MATCH_BSR_ANY          RegexMatchFlags = 16777216
	REGEX_MATCH_PARTIAL_SOFT     RegexMatchFlags = 32768
	REGEX_MATCH_PARTIAL_HARD     RegexMatchFlags = 134217728
	REGEX_MATCH_NOTEMPTY_ATSTART RegexMatchFlags = 268435456
)

type SpawnFlags C.GSpawnFlags

const (
	SPAWN_DEFAULT                SpawnFlags = 0
	SPAWN_LEAVE_DESCRIPTORS_OPEN SpawnFlags = 1
	SPAWN_DO_NOT_REAP_CHILD      SpawnFlags = 2
	SPAWN_SEARCH_PATH            SpawnFlags = 4
	SPAWN_STDOUT_TO_DEV_NULL     SpawnFlags = 8
	SPAWN_STDERR_TO_DEV_NULL     SpawnFlags = 16
	SPAWN_CHILD_INHERITS_STDIN   SpawnFlags = 32
	SPAWN_FILE_AND_ARGV_ZERO     SpawnFlags = 64
	SPAWN_SEARCH_PATH_FROM_ENVP  SpawnFlags = 128
	SPAWN_CLOEXEC_PIPES          SpawnFlags = 256
)

type TestSubprocessFlags C.GTestSubprocessFlags

const (
	TEST_SUBPROCESS_INHERIT_STDIN  TestSubprocessFlags = 1
	TEST_SUBPROCESS_INHERIT_STDOUT TestSubprocessFlags = 2
	TEST_SUBPROCESS_INHERIT_STDERR TestSubprocessFlags = 4
)

type TestTrapFlags C.GTestTrapFlags

const (
	TEST_TRAP_SILENCE_STDOUT TestTrapFlags = 128
	TEST_TRAP_SILENCE_STDERR TestTrapFlags = 256
	TEST_TRAP_INHERIT_STDIN  TestTrapFlags = 512
)

type TraverseFlags C.GTraverseFlags

const (
	TRAVERSE_LEAVES     TraverseFlags = 1
	TRAVERSE_NON_LEAVES TraverseFlags = 2
	TRAVERSE_ALL        TraverseFlags = 3
	TRAVERSE_MASK       TraverseFlags = 3
	TRAVERSE_LEAFS      TraverseFlags = 1
	TRAVERSE_NON_LEAFS  TraverseFlags = 2
)
const ANALYZER_ANALYZING int32 = C.G_ANALYZER_ANALYZING
const ASCII_DTOSTR_BUF_SIZE int32 = C.G_ASCII_DTOSTR_BUF_SIZE
const BIG_ENDIAN int32 = C.G_BIG_ENDIAN
const CSET_A_2_Z string = C.G_CSET_A_2_Z
const CSET_DIGITS string = C.G_CSET_DIGITS
const CSET_a_2_z string = C.G_CSET_a_2_z
const DATALIST_FLAGS_MASK int32 = C.G_DATALIST_FLAGS_MASK
const DATE_BAD_DAY int32 = C.G_DATE_BAD_DAY
const DATE_BAD_JULIAN int32 = C.G_DATE_BAD_JULIAN
const DATE_BAD_YEAR int32 = C.G_DATE_BAD_YEAR
const DIR_SEPARATOR int32 = C.G_DIR_SEPARATOR
const DIR_SEPARATOR_S string = C.G_DIR_SEPARATOR_S
const E float64 = C.G_E
const GINT16_FORMAT string = C.G_GINT16_FORMAT
const GINT16_MODIFIER string = C.G_GINT16_MODIFIER
const GINT32_FORMAT string = C.G_GINT32_FORMAT
const GINT32_MODIFIER string = C.G_GINT32_MODIFIER
const GINT64_FORMAT string = C.G_GINT64_FORMAT
const GINT64_MODIFIER string = C.G_GINT64_MODIFIER
const GINTPTR_FORMAT string = C.G_GINTPTR_FORMAT
const GINTPTR_MODIFIER string = C.G_GINTPTR_MODIFIER
const GNUC_FUNCTION string = C.G_GNUC_FUNCTION
const GNUC_PRETTY_FUNCTION string = C.G_GNUC_PRETTY_FUNCTION
const GSIZE_FORMAT string = C.G_GSIZE_FORMAT
const GSIZE_MODIFIER string = C.G_GSIZE_MODIFIER
const GSSIZE_FORMAT string = C.G_GSSIZE_FORMAT
const GSSIZE_MODIFIER string = C.G_GSSIZE_MODIFIER
const GUINT16_FORMAT string = C.G_GUINT16_FORMAT
const GUINT32_FORMAT string = C.G_GUINT32_FORMAT
const GUINT64_FORMAT string = C.G_GUINT64_FORMAT
const GUINTPTR_FORMAT string = C.G_GUINTPTR_FORMAT
const HAVE_GINT64 int32 = C.G_HAVE_GINT64
const HAVE_GNUC_VARARGS int32 = C.G_HAVE_GNUC_VARARGS
const HAVE_GNUC_VISIBILITY int32 = C.G_HAVE_GNUC_VISIBILITY
const HAVE_GROWING_STACK int32 = C.G_HAVE_GROWING_STACK
const HAVE_ISO_VARARGS int32 = C.G_HAVE_ISO_VARARGS
const HOOK_FLAG_USER_SHIFT int32 = C.G_HOOK_FLAG_USER_SHIFT
const IEEE754_DOUBLE_BIAS int32 = C.G_IEEE754_DOUBLE_BIAS
const IEEE754_FLOAT_BIAS int32 = C.G_IEEE754_FLOAT_BIAS

// Blacklisted : KEY_FILE_DESKTOP_ACTION_GROUP_PREFIX

const KEY_FILE_DESKTOP_GROUP string = C.G_KEY_FILE_DESKTOP_GROUP
const KEY_FILE_DESKTOP_KEY_CATEGORIES string = C.G_KEY_FILE_DESKTOP_KEY_CATEGORIES
const KEY_FILE_DESKTOP_KEY_COMMENT string = C.G_KEY_FILE_DESKTOP_KEY_COMMENT
const KEY_FILE_DESKTOP_KEY_EXEC string = C.G_KEY_FILE_DESKTOP_KEY_EXEC
const KEY_FILE_DESKTOP_KEY_FULLNAME string = C.G_KEY_FILE_DESKTOP_KEY_FULLNAME
const KEY_FILE_DESKTOP_KEY_GENERIC_NAME string = C.G_KEY_FILE_DESKTOP_KEY_GENERIC_NAME
const KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN string = C.G_KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN
const KEY_FILE_DESKTOP_KEY_HIDDEN string = C.G_KEY_FILE_DESKTOP_KEY_HIDDEN
const KEY_FILE_DESKTOP_KEY_ICON string = C.G_KEY_FILE_DESKTOP_KEY_ICON
const KEY_FILE_DESKTOP_KEY_KEYWORDS string = C.G_KEY_FILE_DESKTOP_KEY_KEYWORDS
const KEY_FILE_DESKTOP_KEY_MIME_TYPE string = C.G_KEY_FILE_DESKTOP_KEY_MIME_TYPE
const KEY_FILE_DESKTOP_KEY_NAME string = C.G_KEY_FILE_DESKTOP_KEY_NAME
const KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN string = C.G_KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN
const KEY_FILE_DESKTOP_KEY_NO_DISPLAY string = C.G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY
const KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN string = C.G_KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN
const KEY_FILE_DESKTOP_KEY_PATH string = C.G_KEY_FILE_DESKTOP_KEY_PATH
const KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY string = C.G_KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY
const KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS string = C.G_KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS
const KEY_FILE_DESKTOP_KEY_TERMINAL string = C.G_KEY_FILE_DESKTOP_KEY_TERMINAL
const KEY_FILE_DESKTOP_KEY_TRY_EXEC string = C.G_KEY_FILE_DESKTOP_KEY_TRY_EXEC
const KEY_FILE_DESKTOP_KEY_TYPE string = C.G_KEY_FILE_DESKTOP_KEY_TYPE
const KEY_FILE_DESKTOP_KEY_URL string = C.G_KEY_FILE_DESKTOP_KEY_URL
const KEY_FILE_DESKTOP_KEY_VERSION string = C.G_KEY_FILE_DESKTOP_KEY_VERSION
const KEY_FILE_DESKTOP_TYPE_APPLICATION string = C.G_KEY_FILE_DESKTOP_TYPE_APPLICATION
const KEY_FILE_DESKTOP_TYPE_DIRECTORY string = C.G_KEY_FILE_DESKTOP_TYPE_DIRECTORY
const KEY_FILE_DESKTOP_TYPE_LINK string = C.G_KEY_FILE_DESKTOP_TYPE_LINK
const LITTLE_ENDIAN int32 = C.G_LITTLE_ENDIAN
const LN10 float64 = C.G_LN10
const LN2 float64 = C.G_LN2
const LOG_2_BASE_10 float64 = C.G_LOG_2_BASE_10

// Blacklisted : LOG_DOMAIN

const LOG_FATAL_MASK int32 = C.G_LOG_FATAL_MASK
const LOG_LEVEL_USER_SHIFT int32 = C.G_LOG_LEVEL_USER_SHIFT
const MAJOR_VERSION int32 = C.GLIB_MAJOR_VERSION
const MAXINT16 int16 = C.G_MAXINT16
const MAXINT32 int32 = C.G_MAXINT32
const MAXINT64 int64 = C.G_MAXINT64
const MAXINT8 int8 = C.G_MAXINT8
const MAXUINT16 uint16 = C.G_MAXUINT16
const MAXUINT32 uint32 = C.G_MAXUINT32
const MAXUINT64 uint64 = C.G_MAXUINT64
const MAXUINT8 uint8 = C.G_MAXUINT8
const MICRO_VERSION int32 = C.GLIB_MICRO_VERSION
const MININT16 int16 = C.G_MININT16
const MININT32 int32 = C.G_MININT32
const MININT64 int64 = C.G_MININT64
const MININT8 int8 = C.G_MININT8
const MINOR_VERSION int32 = C.GLIB_MINOR_VERSION
const MODULE_SUFFIX string = C.G_MODULE_SUFFIX
const OPTION_REMAINING string = C.G_OPTION_REMAINING
const PDP_ENDIAN int32 = C.G_PDP_ENDIAN
const PI float64 = C.G_PI
const PI_2 float64 = C.G_PI_2
const PI_4 float64 = C.G_PI_4
const POLLFD_FORMAT string = C.G_POLLFD_FORMAT
const PRIORITY_DEFAULT int32 = C.G_PRIORITY_DEFAULT
const PRIORITY_DEFAULT_IDLE int32 = C.G_PRIORITY_DEFAULT_IDLE
const PRIORITY_HIGH int32 = C.G_PRIORITY_HIGH
const PRIORITY_HIGH_IDLE int32 = C.G_PRIORITY_HIGH_IDLE
const PRIORITY_LOW int32 = C.G_PRIORITY_LOW
const SEARCHPATH_SEPARATOR int32 = C.G_SEARCHPATH_SEPARATOR
const SEARCHPATH_SEPARATOR_S string = C.G_SEARCHPATH_SEPARATOR_S
const SIZEOF_LONG int32 = C.GLIB_SIZEOF_LONG
const SIZEOF_SIZE_T int32 = C.GLIB_SIZEOF_SIZE_T
const SIZEOF_SSIZE_T int32 = C.GLIB_SIZEOF_SSIZE_T
const SIZEOF_VOID_P int32 = C.GLIB_SIZEOF_VOID_P
const SOURCE_CONTINUE bool = true // C.G_SOURCE_CONTINUE
const SOURCE_REMOVE bool = false  // C.G_SOURCE_REMOVE
const SQRT2 float64 = C.G_SQRT2
const STR_DELIMITERS string = C.G_STR_DELIMITERS
const SYSDEF_AF_INET int32 = C.GLIB_SYSDEF_AF_INET
const SYSDEF_AF_INET6 int32 = C.GLIB_SYSDEF_AF_INET6
const SYSDEF_AF_UNIX int32 = C.GLIB_SYSDEF_AF_UNIX
const SYSDEF_MSG_DONTROUTE int32 = C.GLIB_SYSDEF_MSG_DONTROUTE
const SYSDEF_MSG_OOB int32 = C.GLIB_SYSDEF_MSG_OOB
const SYSDEF_MSG_PEEK int32 = C.GLIB_SYSDEF_MSG_PEEK
const TIME_SPAN_DAY int64 = C.G_TIME_SPAN_DAY
const TIME_SPAN_HOUR int64 = C.G_TIME_SPAN_HOUR
const TIME_SPAN_MILLISECOND int64 = C.G_TIME_SPAN_MILLISECOND
const TIME_SPAN_MINUTE int64 = C.G_TIME_SPAN_MINUTE
const TIME_SPAN_SECOND int64 = C.G_TIME_SPAN_SECOND
const UNICHAR_MAX_DECOMPOSITION_LENGTH int32 = C.G_UNICHAR_MAX_DECOMPOSITION_LENGTH
const URI_RESERVED_CHARS_GENERIC_DELIMITERS string = C.G_URI_RESERVED_CHARS_GENERIC_DELIMITERS
const URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS string = C.G_URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS
const USEC_PER_SEC int32 = C.G_USEC_PER_SEC
const VA_COPY_AS_ARRAY int32 = C.G_VA_COPY_AS_ARRAY
const VERSION_MIN_REQUIRED int32 = C.GLIB_VERSION_MIN_REQUIRED

// Blacklisted : WIN32_MSG_HANDLE

type BookmarkFileError C.GBookmarkFileError

const (
	BOOKMARK_FILE_ERROR_INVALID_URI        BookmarkFileError = 0
	BOOKMARK_FILE_ERROR_INVALID_VALUE      BookmarkFileError = 1
	BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED BookmarkFileError = 2
	BOOKMARK_FILE_ERROR_URI_NOT_FOUND      BookmarkFileError = 3
	BOOKMARK_FILE_ERROR_READ               BookmarkFileError = 4
	BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING   BookmarkFileError = 5
	BOOKMARK_FILE_ERROR_WRITE              BookmarkFileError = 6
	BOOKMARK_FILE_ERROR_FILE_NOT_FOUND     BookmarkFileError = 7
)

type ChecksumType C.GChecksumType

const (
	CHECKSUM_MD5    ChecksumType = 0
	CHECKSUM_SHA1   ChecksumType = 1
	CHECKSUM_SHA256 ChecksumType = 2
	CHECKSUM_SHA512 ChecksumType = 3
	CHECKSUM_SHA384 ChecksumType = 4
)

type ConvertError C.GConvertError

const (
	CONVERT_ERROR_NO_CONVERSION     ConvertError = 0
	CONVERT_ERROR_ILLEGAL_SEQUENCE  ConvertError = 1
	CONVERT_ERROR_FAILED            ConvertError = 2
	CONVERT_ERROR_PARTIAL_INPUT     ConvertError = 3
	CONVERT_ERROR_BAD_URI           ConvertError = 4
	CONVERT_ERROR_NOT_ABSOLUTE_PATH ConvertError = 5
	CONVERT_ERROR_NO_MEMORY         ConvertError = 6
	CONVERT_ERROR_EMBEDDED_NUL      ConvertError = 7
)

type DateDMY C.GDateDMY

const (
	DATE_DAY   DateDMY = 0
	DATE_MONTH DateDMY = 1
	DATE_YEAR  DateDMY = 2
)

type DateMonth C.GDateMonth

const (
	DATE_BAD_MONTH DateMonth = 0
	DATE_JANUARY   DateMonth = 1
	DATE_FEBRUARY  DateMonth = 2
	DATE_MARCH     DateMonth = 3
	DATE_APRIL     DateMonth = 4
	DATE_MAY       DateMonth = 5
	DATE_JUNE      DateMonth = 6
	DATE_JULY      DateMonth = 7
	DATE_AUGUST    DateMonth = 8
	DATE_SEPTEMBER DateMonth = 9
	DATE_OCTOBER   DateMonth = 10
	DATE_NOVEMBER  DateMonth = 11
	DATE_DECEMBER  DateMonth = 12
)

type DateWeekday C.GDateWeekday

const (
	DATE_BAD_WEEKDAY DateWeekday = 0
	DATE_MONDAY      DateWeekday = 1
	DATE_TUESDAY     DateWeekday = 2
	DATE_WEDNESDAY   DateWeekday = 3
	DATE_THURSDAY    DateWeekday = 4
	DATE_FRIDAY      DateWeekday = 5
	DATE_SATURDAY    DateWeekday = 6
	DATE_SUNDAY      DateWeekday = 7
)

type ErrorType C.GErrorType

const (
	ERR_UNKNOWN              ErrorType = 0
	ERR_UNEXP_EOF            ErrorType = 1
	ERR_UNEXP_EOF_IN_STRING  ErrorType = 2
	ERR_UNEXP_EOF_IN_COMMENT ErrorType = 3
	ERR_NON_DIGIT_IN_CONST   ErrorType = 4
	ERR_DIGIT_RADIX          ErrorType = 5
	ERR_FLOAT_RADIX          ErrorType = 6
	ERR_FLOAT_MALFORMED      ErrorType = 7
)

type FileError C.GFileError

const (
	FILE_ERROR_EXIST       FileError = 0
	FILE_ERROR_ISDIR       FileError = 1
	FILE_ERROR_ACCES       FileError = 2
	FILE_ERROR_NAMETOOLONG FileError = 3
	FILE_ERROR_NOENT       FileError = 4
	FILE_ERROR_NOTDIR      FileError = 5
	FILE_ERROR_NXIO        FileError = 6
	FILE_ERROR_NODEV       FileError = 7
	FILE_ERROR_ROFS        FileError = 8
	FILE_ERROR_TXTBSY      FileError = 9
	FILE_ERROR_FAULT       FileError = 10
	FILE_ERROR_LOOP        FileError = 11
	FILE_ERROR_NOSPC       FileError = 12
	FILE_ERROR_NOMEM       FileError = 13
	FILE_ERROR_MFILE       FileError = 14
	FILE_ERROR_NFILE       FileError = 15
	FILE_ERROR_BADF        FileError = 16
	FILE_ERROR_INVAL       FileError = 17
	FILE_ERROR_PIPE        FileError = 18
	FILE_ERROR_AGAIN       FileError = 19
	FILE_ERROR_INTR        FileError = 20
	FILE_ERROR_IO          FileError = 21
	FILE_ERROR_PERM        FileError = 22
	FILE_ERROR_NOSYS       FileError = 23
	FILE_ERROR_FAILED      FileError = 24
)

type IOChannelError C.GIOChannelError

const (
	IO_CHANNEL_ERROR_FBIG     IOChannelError = 0
	IO_CHANNEL_ERROR_INVAL    IOChannelError = 1
	IO_CHANNEL_ERROR_IO       IOChannelError = 2
	IO_CHANNEL_ERROR_ISDIR    IOChannelError = 3
	IO_CHANNEL_ERROR_NOSPC    IOChannelError = 4
	IO_CHANNEL_ERROR_NXIO     IOChannelError = 5
	IO_CHANNEL_ERROR_OVERFLOW IOChannelError = 6
	IO_CHANNEL_ERROR_PIPE     IOChannelError = 7
	IO_CHANNEL_ERROR_FAILED   IOChannelError = 8
)

type IOError C.GIOError

const (
	IO_ERROR_NONE    IOError = 0
	IO_ERROR_AGAIN   IOError = 1
	IO_ERROR_INVAL   IOError = 2
	IO_ERROR_UNKNOWN IOError = 3
)

type IOStatus C.GIOStatus

const (
	IO_STATUS_ERROR  IOStatus = 0
	IO_STATUS_NORMAL IOStatus = 1
	IO_STATUS_EOF    IOStatus = 2
	IO_STATUS_AGAIN  IOStatus = 3
)

type KeyFileError C.GKeyFileError

const (
	KEY_FILE_ERROR_UNKNOWN_ENCODING KeyFileError = 0
	KEY_FILE_ERROR_PARSE            KeyFileError = 1
	KEY_FILE_ERROR_NOT_FOUND        KeyFileError = 2
	KEY_FILE_ERROR_KEY_NOT_FOUND    KeyFileError = 3
	KEY_FILE_ERROR_GROUP_NOT_FOUND  KeyFileError = 4
	KEY_FILE_ERROR_INVALID_VALUE    KeyFileError = 5
)

type MarkupError C.GMarkupError

const (
	MARKUP_ERROR_BAD_UTF8          MarkupError = 0
	MARKUP_ERROR_EMPTY             MarkupError = 1
	MARKUP_ERROR_PARSE             MarkupError = 2
	MARKUP_ERROR_UNKNOWN_ELEMENT   MarkupError = 3
	MARKUP_ERROR_UNKNOWN_ATTRIBUTE MarkupError = 4
	MARKUP_ERROR_INVALID_CONTENT   MarkupError = 5
	MARKUP_ERROR_MISSING_ATTRIBUTE MarkupError = 6
)

type NormalizeMode C.GNormalizeMode

const (
	NORMALIZE_DEFAULT         NormalizeMode = 0
	NORMALIZE_NFD             NormalizeMode = 0
	NORMALIZE_DEFAULT_COMPOSE NormalizeMode = 1
	NORMALIZE_NFC             NormalizeMode = 1
	NORMALIZE_ALL             NormalizeMode = 2
	NORMALIZE_NFKD            NormalizeMode = 2
	NORMALIZE_ALL_COMPOSE     NormalizeMode = 3
	NORMALIZE_NFKC            NormalizeMode = 3
)

type OnceStatus C.GOnceStatus

const (
	ONCE_STATUS_NOTCALLED OnceStatus = 0
	ONCE_STATUS_PROGRESS  OnceStatus = 1
	ONCE_STATUS_READY     OnceStatus = 2
)

type OptionArg C.GOptionArg

const (
	OPTION_ARG_NONE           OptionArg = 0
	OPTION_ARG_STRING         OptionArg = 1
	OPTION_ARG_INT            OptionArg = 2
	OPTION_ARG_CALLBACK       OptionArg = 3
	OPTION_ARG_FILENAME       OptionArg = 4
	OPTION_ARG_STRING_ARRAY   OptionArg = 5
	OPTION_ARG_FILENAME_ARRAY OptionArg = 6
	OPTION_ARG_DOUBLE         OptionArg = 7
	OPTION_ARG_INT64          OptionArg = 8
)

type OptionError C.GOptionError

const (
	OPTION_ERROR_UNKNOWN_OPTION OptionError = 0
	OPTION_ERROR_BAD_VALUE      OptionError = 1
	OPTION_ERROR_FAILED         OptionError = 2
)

type RegexError C.GRegexError

const (
	REGEX_ERROR_COMPILE                                      RegexError = 0
	REGEX_ERROR_OPTIMIZE                                     RegexError = 1
	REGEX_ERROR_REPLACE                                      RegexError = 2
	REGEX_ERROR_MATCH                                        RegexError = 3
	REGEX_ERROR_INTERNAL                                     RegexError = 4
	REGEX_ERROR_STRAY_BACKSLASH                              RegexError = 101
	REGEX_ERROR_MISSING_CONTROL_CHAR                         RegexError = 102
	REGEX_ERROR_UNRECOGNIZED_ESCAPE                          RegexError = 103
	REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER                     RegexError = 104
	REGEX_ERROR_QUANTIFIER_TOO_BIG                           RegexError = 105
	REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS                 RegexError = 106
	REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS            RegexError = 107
	REGEX_ERROR_RANGE_OUT_OF_ORDER                           RegexError = 108
	REGEX_ERROR_NOTHING_TO_REPEAT                            RegexError = 109
	REGEX_ERROR_UNRECOGNIZED_CHARACTER                       RegexError = 112
	REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS              RegexError = 113
	REGEX_ERROR_UNMATCHED_PARENTHESIS                        RegexError = 114
	REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE              RegexError = 115
	REGEX_ERROR_UNTERMINATED_COMMENT                         RegexError = 118
	REGEX_ERROR_EXPRESSION_TOO_LARGE                         RegexError = 120
	REGEX_ERROR_MEMORY_ERROR                                 RegexError = 121
	REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND                   RegexError = 125
	REGEX_ERROR_MALFORMED_CONDITION                          RegexError = 126
	REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES                RegexError = 127
	REGEX_ERROR_ASSERTION_EXPECTED                           RegexError = 128
	REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME                     RegexError = 130
	REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED       RegexError = 131
	REGEX_ERROR_HEX_CODE_TOO_LARGE                           RegexError = 134
	REGEX_ERROR_INVALID_CONDITION                            RegexError = 135
	REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND              RegexError = 136
	REGEX_ERROR_INFINITE_LOOP                                RegexError = 140
	REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR           RegexError = 142
	REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME                    RegexError = 143
	REGEX_ERROR_MALFORMED_PROPERTY                           RegexError = 146
	REGEX_ERROR_UNKNOWN_PROPERTY                             RegexError = 147
	REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG                     RegexError = 148
	REGEX_ERROR_TOO_MANY_SUBPATTERNS                         RegexError = 149
	REGEX_ERROR_INVALID_OCTAL_VALUE                          RegexError = 151
	REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE                  RegexError = 154
	REGEX_ERROR_DEFINE_REPETION                              RegexError = 155
	REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS                 RegexError = 156
	REGEX_ERROR_MISSING_BACK_REFERENCE                       RegexError = 157
	REGEX_ERROR_INVALID_RELATIVE_REFERENCE                   RegexError = 158
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN RegexError = 159
	REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB            RegexError = 160
	REGEX_ERROR_NUMBER_TOO_BIG                               RegexError = 161
	REGEX_ERROR_MISSING_SUBPATTERN_NAME                      RegexError = 162
	REGEX_ERROR_MISSING_DIGIT                                RegexError = 163
	REGEX_ERROR_INVALID_DATA_CHARACTER                       RegexError = 164
	REGEX_ERROR_EXTRA_SUBPATTERN_NAME                        RegexError = 165
	REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED  RegexError = 166
	REGEX_ERROR_INVALID_CONTROL_CHAR                         RegexError = 168
	REGEX_ERROR_MISSING_NAME                                 RegexError = 169
	REGEX_ERROR_NOT_SUPPORTED_IN_CLASS                       RegexError = 171
	REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES                  RegexError = 172
	REGEX_ERROR_NAME_TOO_LONG                                RegexError = 175
	REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE                    RegexError = 176
)

type SeekType C.GSeekType

const (
	SEEK_CUR SeekType = 0
	SEEK_SET SeekType = 1
	SEEK_END SeekType = 2
)

type ShellError C.GShellError

const (
	SHELL_ERROR_BAD_QUOTING  ShellError = 0
	SHELL_ERROR_EMPTY_STRING ShellError = 1
	SHELL_ERROR_FAILED       ShellError = 2
)

type SliceConfig C.GSliceConfig

const (
	SLICE_CONFIG_ALWAYS_MALLOC      SliceConfig = 1
	SLICE_CONFIG_BYPASS_MAGAZINES   SliceConfig = 2
	SLICE_CONFIG_WORKING_SET_MSECS  SliceConfig = 3
	SLICE_CONFIG_COLOR_INCREMENT    SliceConfig = 4
	SLICE_CONFIG_CHUNK_SIZES        SliceConfig = 5
	SLICE_CONFIG_CONTENTION_COUNTER SliceConfig = 6
)

type SpawnError C.GSpawnError

const (
	SPAWN_ERROR_FORK        SpawnError = 0
	SPAWN_ERROR_READ        SpawnError = 1
	SPAWN_ERROR_CHDIR       SpawnError = 2
	SPAWN_ERROR_ACCES       SpawnError = 3
	SPAWN_ERROR_PERM        SpawnError = 4
	SPAWN_ERROR_TOO_BIG     SpawnError = 5
	SPAWN_ERROR_2BIG        SpawnError = 5
	SPAWN_ERROR_NOEXEC      SpawnError = 6
	SPAWN_ERROR_NAMETOOLONG SpawnError = 7
	SPAWN_ERROR_NOENT       SpawnError = 8
	SPAWN_ERROR_NOMEM       SpawnError = 9
	SPAWN_ERROR_NOTDIR      SpawnError = 10
	SPAWN_ERROR_LOOP        SpawnError = 11
	SPAWN_ERROR_TXTBUSY     SpawnError = 12
	SPAWN_ERROR_IO          SpawnError = 13
	SPAWN_ERROR_NFILE       SpawnError = 14
	SPAWN_ERROR_MFILE       SpawnError = 15
	SPAWN_ERROR_INVAL       SpawnError = 16
	SPAWN_ERROR_ISDIR       SpawnError = 17
	SPAWN_ERROR_LIBBAD      SpawnError = 18
	SPAWN_ERROR_FAILED      SpawnError = 19
)

type TestLogType C.GTestLogType

const (
	TEST_LOG_NONE         TestLogType = 0
	TEST_LOG_ERROR        TestLogType = 1
	TEST_LOG_START_BINARY TestLogType = 2
	TEST_LOG_LIST_CASE    TestLogType = 3
	TEST_LOG_SKIP_CASE    TestLogType = 4
	TEST_LOG_START_CASE   TestLogType = 5
	TEST_LOG_STOP_CASE    TestLogType = 6
	TEST_LOG_MIN_RESULT   TestLogType = 7
	TEST_LOG_MAX_RESULT   TestLogType = 8
	TEST_LOG_MESSAGE      TestLogType = 9
	TEST_LOG_START_SUITE  TestLogType = 10
	TEST_LOG_STOP_SUITE   TestLogType = 11
)

// Blacklisted : GTestResult

type ThreadError C.GThreadError

const (
	THREAD_ERROR_AGAIN ThreadError = 0
)

type TimeType C.GTimeType

const (
	TIME_TYPE_STANDARD  TimeType = 0
	TIME_TYPE_DAYLIGHT  TimeType = 1
	TIME_TYPE_UNIVERSAL TimeType = 2
)

type TokenType C.GTokenType

const (
	TOKEN_EOF             TokenType = 0
	TOKEN_LEFT_PAREN      TokenType = 40
	TOKEN_RIGHT_PAREN     TokenType = 41
	TOKEN_LEFT_CURLY      TokenType = 123
	TOKEN_RIGHT_CURLY     TokenType = 125
	TOKEN_LEFT_BRACE      TokenType = 91
	TOKEN_RIGHT_BRACE     TokenType = 93
	TOKEN_EQUAL_SIGN      TokenType = 61
	TOKEN_COMMA           TokenType = 44
	TOKEN_NONE            TokenType = 256
	TOKEN_ERROR           TokenType = 257
	TOKEN_CHAR            TokenType = 258
	TOKEN_BINARY          TokenType = 259
	TOKEN_OCTAL           TokenType = 260
	TOKEN_INT             TokenType = 261
	TOKEN_HEX             TokenType = 262
	TOKEN_FLOAT           TokenType = 263
	TOKEN_STRING          TokenType = 264
	TOKEN_SYMBOL          TokenType = 265
	TOKEN_IDENTIFIER      TokenType = 266
	TOKEN_IDENTIFIER_NULL TokenType = 267
	TOKEN_COMMENT_SINGLE  TokenType = 268
	TOKEN_COMMENT_MULTI   TokenType = 269
)

type TraverseType C.GTraverseType

const (
	IN_ORDER    TraverseType = 0
	PRE_ORDER   TraverseType = 1
	POST_ORDER  TraverseType = 2
	LEVEL_ORDER TraverseType = 3
)

type UnicodeBreakType C.GUnicodeBreakType

const (
	UNICODE_BREAK_MANDATORY                    UnicodeBreakType = 0
	UNICODE_BREAK_CARRIAGE_RETURN              UnicodeBreakType = 1
	UNICODE_BREAK_LINE_FEED                    UnicodeBreakType = 2
	UNICODE_BREAK_COMBINING_MARK               UnicodeBreakType = 3
	UNICODE_BREAK_SURROGATE                    UnicodeBreakType = 4
	UNICODE_BREAK_ZERO_WIDTH_SPACE             UnicodeBreakType = 5
	UNICODE_BREAK_INSEPARABLE                  UnicodeBreakType = 6
	UNICODE_BREAK_NON_BREAKING_GLUE            UnicodeBreakType = 7
	UNICODE_BREAK_CONTINGENT                   UnicodeBreakType = 8
	UNICODE_BREAK_SPACE                        UnicodeBreakType = 9
	UNICODE_BREAK_AFTER                        UnicodeBreakType = 10
	UNICODE_BREAK_BEFORE                       UnicodeBreakType = 11
	UNICODE_BREAK_BEFORE_AND_AFTER             UnicodeBreakType = 12
	UNICODE_BREAK_HYPHEN                       UnicodeBreakType = 13
	UNICODE_BREAK_NON_STARTER                  UnicodeBreakType = 14
	UNICODE_BREAK_OPEN_PUNCTUATION             UnicodeBreakType = 15
	UNICODE_BREAK_CLOSE_PUNCTUATION            UnicodeBreakType = 16
	UNICODE_BREAK_QUOTATION                    UnicodeBreakType = 17
	UNICODE_BREAK_EXCLAMATION                  UnicodeBreakType = 18
	UNICODE_BREAK_IDEOGRAPHIC                  UnicodeBreakType = 19
	UNICODE_BREAK_NUMERIC                      UnicodeBreakType = 20
	UNICODE_BREAK_INFIX_SEPARATOR              UnicodeBreakType = 21
	UNICODE_BREAK_SYMBOL                       UnicodeBreakType = 22
	UNICODE_BREAK_ALPHABETIC                   UnicodeBreakType = 23
	UNICODE_BREAK_PREFIX                       UnicodeBreakType = 24
	UNICODE_BREAK_POSTFIX                      UnicodeBreakType = 25
	UNICODE_BREAK_COMPLEX_CONTEXT              UnicodeBreakType = 26
	UNICODE_BREAK_AMBIGUOUS                    UnicodeBreakType = 27
	UNICODE_BREAK_UNKNOWN                      UnicodeBreakType = 28
	UNICODE_BREAK_NEXT_LINE                    UnicodeBreakType = 29
	UNICODE_BREAK_WORD_JOINER                  UnicodeBreakType = 30
	UNICODE_BREAK_HANGUL_L_JAMO                UnicodeBreakType = 31
	UNICODE_BREAK_HANGUL_V_JAMO                UnicodeBreakType = 32
	UNICODE_BREAK_HANGUL_T_JAMO                UnicodeBreakType = 33
	UNICODE_BREAK_HANGUL_LV_SYLLABLE           UnicodeBreakType = 34
	UNICODE_BREAK_HANGUL_LVT_SYLLABLE          UnicodeBreakType = 35
	UNICODE_BREAK_CLOSE_PARANTHESIS            UnicodeBreakType = 36
	UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER UnicodeBreakType = 37
	UNICODE_BREAK_HEBREW_LETTER                UnicodeBreakType = 38
	UNICODE_BREAK_REGIONAL_INDICATOR           UnicodeBreakType = 39
	UNICODE_BREAK_EMOJI_BASE                   UnicodeBreakType = 40
	UNICODE_BREAK_EMOJI_MODIFIER               UnicodeBreakType = 41
	UNICODE_BREAK_ZERO_WIDTH_JOINER            UnicodeBreakType = 42
)

type UnicodeScript C.GUnicodeScript

const (
	UNICODE_SCRIPT_INVALID_CODE           UnicodeScript = -1
	UNICODE_SCRIPT_COMMON                 UnicodeScript = 0
	UNICODE_SCRIPT_INHERITED              UnicodeScript = 1
	UNICODE_SCRIPT_ARABIC                 UnicodeScript = 2
	UNICODE_SCRIPT_ARMENIAN               UnicodeScript = 3
	UNICODE_SCRIPT_BENGALI                UnicodeScript = 4
	UNICODE_SCRIPT_BOPOMOFO               UnicodeScript = 5
	UNICODE_SCRIPT_CHEROKEE               UnicodeScript = 6
	UNICODE_SCRIPT_COPTIC                 UnicodeScript = 7
	UNICODE_SCRIPT_CYRILLIC               UnicodeScript = 8
	UNICODE_SCRIPT_DESERET                UnicodeScript = 9
	UNICODE_SCRIPT_DEVANAGARI             UnicodeScript = 10
	UNICODE_SCRIPT_ETHIOPIC               UnicodeScript = 11
	UNICODE_SCRIPT_GEORGIAN               UnicodeScript = 12
	UNICODE_SCRIPT_GOTHIC                 UnicodeScript = 13
	UNICODE_SCRIPT_GREEK                  UnicodeScript = 14
	UNICODE_SCRIPT_GUJARATI               UnicodeScript = 15
	UNICODE_SCRIPT_GURMUKHI               UnicodeScript = 16
	UNICODE_SCRIPT_HAN                    UnicodeScript = 17
	UNICODE_SCRIPT_HANGUL                 UnicodeScript = 18
	UNICODE_SCRIPT_HEBREW                 UnicodeScript = 19
	UNICODE_SCRIPT_HIRAGANA               UnicodeScript = 20
	UNICODE_SCRIPT_KANNADA                UnicodeScript = 21
	UNICODE_SCRIPT_KATAKANA               UnicodeScript = 22
	UNICODE_SCRIPT_KHMER                  UnicodeScript = 23
	UNICODE_SCRIPT_LAO                    UnicodeScript = 24
	UNICODE_SCRIPT_LATIN                  UnicodeScript = 25
	UNICODE_SCRIPT_MALAYALAM              UnicodeScript = 26
	UNICODE_SCRIPT_MONGOLIAN              UnicodeScript = 27
	UNICODE_SCRIPT_MYANMAR                UnicodeScript = 28
	UNICODE_SCRIPT_OGHAM                  UnicodeScript = 29
	UNICODE_SCRIPT_OLD_ITALIC             UnicodeScript = 30
	UNICODE_SCRIPT_ORIYA                  UnicodeScript = 31
	UNICODE_SCRIPT_RUNIC                  UnicodeScript = 32
	UNICODE_SCRIPT_SINHALA                UnicodeScript = 33
	UNICODE_SCRIPT_SYRIAC                 UnicodeScript = 34
	UNICODE_SCRIPT_TAMIL                  UnicodeScript = 35
	UNICODE_SCRIPT_TELUGU                 UnicodeScript = 36
	UNICODE_SCRIPT_THAANA                 UnicodeScript = 37
	UNICODE_SCRIPT_THAI                   UnicodeScript = 38
	UNICODE_SCRIPT_TIBETAN                UnicodeScript = 39
	UNICODE_SCRIPT_CANADIAN_ABORIGINAL    UnicodeScript = 40
	UNICODE_SCRIPT_YI                     UnicodeScript = 41
	UNICODE_SCRIPT_TAGALOG                UnicodeScript = 42
	UNICODE_SCRIPT_HANUNOO                UnicodeScript = 43
	UNICODE_SCRIPT_BUHID                  UnicodeScript = 44
	UNICODE_SCRIPT_TAGBANWA               UnicodeScript = 45
	UNICODE_SCRIPT_BRAILLE                UnicodeScript = 46
	UNICODE_SCRIPT_CYPRIOT                UnicodeScript = 47
	UNICODE_SCRIPT_LIMBU                  UnicodeScript = 48
	UNICODE_SCRIPT_OSMANYA                UnicodeScript = 49
	UNICODE_SCRIPT_SHAVIAN                UnicodeScript = 50
	UNICODE_SCRIPT_LINEAR_B               UnicodeScript = 51
	UNICODE_SCRIPT_TAI_LE                 UnicodeScript = 52
	UNICODE_SCRIPT_UGARITIC               UnicodeScript = 53
	UNICODE_SCRIPT_NEW_TAI_LUE            UnicodeScript = 54
	UNICODE_SCRIPT_BUGINESE               UnicodeScript = 55
	UNICODE_SCRIPT_GLAGOLITIC             UnicodeScript = 56
	UNICODE_SCRIPT_TIFINAGH               UnicodeScript = 57
	UNICODE_SCRIPT_SYLOTI_NAGRI           UnicodeScript = 58
	UNICODE_SCRIPT_OLD_PERSIAN            UnicodeScript = 59
	UNICODE_SCRIPT_KHAROSHTHI             UnicodeScript = 60
	UNICODE_SCRIPT_UNKNOWN                UnicodeScript = 61
	UNICODE_SCRIPT_BALINESE               UnicodeScript = 62
	UNICODE_SCRIPT_CUNEIFORM              UnicodeScript = 63
	UNICODE_SCRIPT_PHOENICIAN             UnicodeScript = 64
	UNICODE_SCRIPT_PHAGS_PA               UnicodeScript = 65
	UNICODE_SCRIPT_NKO                    UnicodeScript = 66
	UNICODE_SCRIPT_KAYAH_LI               UnicodeScript = 67
	UNICODE_SCRIPT_LEPCHA                 UnicodeScript = 68
	UNICODE_SCRIPT_REJANG                 UnicodeScript = 69
	UNICODE_SCRIPT_SUNDANESE              UnicodeScript = 70
	UNICODE_SCRIPT_SAURASHTRA             UnicodeScript = 71
	UNICODE_SCRIPT_CHAM                   UnicodeScript = 72
	UNICODE_SCRIPT_OL_CHIKI               UnicodeScript = 73
	UNICODE_SCRIPT_VAI                    UnicodeScript = 74
	UNICODE_SCRIPT_CARIAN                 UnicodeScript = 75
	UNICODE_SCRIPT_LYCIAN                 UnicodeScript = 76
	UNICODE_SCRIPT_LYDIAN                 UnicodeScript = 77
	UNICODE_SCRIPT_AVESTAN                UnicodeScript = 78
	UNICODE_SCRIPT_BAMUM                  UnicodeScript = 79
	UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS   UnicodeScript = 80
	UNICODE_SCRIPT_IMPERIAL_ARAMAIC       UnicodeScript = 81
	UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI  UnicodeScript = 82
	UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN UnicodeScript = 83
	UNICODE_SCRIPT_JAVANESE               UnicodeScript = 84
	UNICODE_SCRIPT_KAITHI                 UnicodeScript = 85
	UNICODE_SCRIPT_LISU                   UnicodeScript = 86
	UNICODE_SCRIPT_MEETEI_MAYEK           UnicodeScript = 87
	UNICODE_SCRIPT_OLD_SOUTH_ARABIAN      UnicodeScript = 88
	UNICODE_SCRIPT_OLD_TURKIC             UnicodeScript = 89
	UNICODE_SCRIPT_SAMARITAN              UnicodeScript = 90
	UNICODE_SCRIPT_TAI_THAM               UnicodeScript = 91
	UNICODE_SCRIPT_TAI_VIET               UnicodeScript = 92
	UNICODE_SCRIPT_BATAK                  UnicodeScript = 93
	UNICODE_SCRIPT_BRAHMI                 UnicodeScript = 94
	UNICODE_SCRIPT_MANDAIC                UnicodeScript = 95
	UNICODE_SCRIPT_CHAKMA                 UnicodeScript = 96
	UNICODE_SCRIPT_MEROITIC_CURSIVE       UnicodeScript = 97
	UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS   UnicodeScript = 98
	UNICODE_SCRIPT_MIAO                   UnicodeScript = 99
	UNICODE_SCRIPT_SHARADA                UnicodeScript = 100
	UNICODE_SCRIPT_SORA_SOMPENG           UnicodeScript = 101
	UNICODE_SCRIPT_TAKRI                  UnicodeScript = 102
	UNICODE_SCRIPT_BASSA_VAH              UnicodeScript = 103
	UNICODE_SCRIPT_CAUCASIAN_ALBANIAN     UnicodeScript = 104
	UNICODE_SCRIPT_DUPLOYAN               UnicodeScript = 105
	UNICODE_SCRIPT_ELBASAN                UnicodeScript = 106
	UNICODE_SCRIPT_GRANTHA                UnicodeScript = 107
	UNICODE_SCRIPT_KHOJKI                 UnicodeScript = 108
	UNICODE_SCRIPT_KHUDAWADI              UnicodeScript = 109
	UNICODE_SCRIPT_LINEAR_A               UnicodeScript = 110
	UNICODE_SCRIPT_MAHAJANI               UnicodeScript = 111
	UNICODE_SCRIPT_MANICHAEAN             UnicodeScript = 112
	UNICODE_SCRIPT_MENDE_KIKAKUI          UnicodeScript = 113
	UNICODE_SCRIPT_MODI                   UnicodeScript = 114
	UNICODE_SCRIPT_MRO                    UnicodeScript = 115
	UNICODE_SCRIPT_NABATAEAN              UnicodeScript = 116
	UNICODE_SCRIPT_OLD_NORTH_ARABIAN      UnicodeScript = 117
	UNICODE_SCRIPT_OLD_PERMIC             UnicodeScript = 118
	UNICODE_SCRIPT_PAHAWH_HMONG           UnicodeScript = 119
	UNICODE_SCRIPT_PALMYRENE              UnicodeScript = 120
	UNICODE_SCRIPT_PAU_CIN_HAU            UnicodeScript = 121
	UNICODE_SCRIPT_PSALTER_PAHLAVI        UnicodeScript = 122
	UNICODE_SCRIPT_SIDDHAM                UnicodeScript = 123
	UNICODE_SCRIPT_TIRHUTA                UnicodeScript = 124
	UNICODE_SCRIPT_WARANG_CITI            UnicodeScript = 125
	UNICODE_SCRIPT_AHOM                   UnicodeScript = 126
	UNICODE_SCRIPT_ANATOLIAN_HIEROGLYPHS  UnicodeScript = 127
	UNICODE_SCRIPT_HATRAN                 UnicodeScript = 128
	UNICODE_SCRIPT_MULTANI                UnicodeScript = 129
	UNICODE_SCRIPT_OLD_HUNGARIAN          UnicodeScript = 130
	UNICODE_SCRIPT_SIGNWRITING            UnicodeScript = 131
	UNICODE_SCRIPT_ADLAM                  UnicodeScript = 132
	UNICODE_SCRIPT_BHAIKSUKI              UnicodeScript = 133
	UNICODE_SCRIPT_MARCHEN                UnicodeScript = 134
	UNICODE_SCRIPT_NEWA                   UnicodeScript = 135
	UNICODE_SCRIPT_OSAGE                  UnicodeScript = 136
	UNICODE_SCRIPT_TANGUT                 UnicodeScript = 137
	UNICODE_SCRIPT_MASARAM_GONDI          UnicodeScript = 138
	UNICODE_SCRIPT_NUSHU                  UnicodeScript = 139
	UNICODE_SCRIPT_SOYOMBO                UnicodeScript = 140
	UNICODE_SCRIPT_ZANABAZAR_SQUARE       UnicodeScript = 141
)

type UnicodeType C.GUnicodeType

const (
	UNICODE_CONTROL             UnicodeType = 0
	UNICODE_FORMAT              UnicodeType = 1
	UNICODE_UNASSIGNED          UnicodeType = 2
	UNICODE_PRIVATE_USE         UnicodeType = 3
	UNICODE_SURROGATE           UnicodeType = 4
	UNICODE_LOWERCASE_LETTER    UnicodeType = 5
	UNICODE_MODIFIER_LETTER     UnicodeType = 6
	UNICODE_OTHER_LETTER        UnicodeType = 7
	UNICODE_TITLECASE_LETTER    UnicodeType = 8
	UNICODE_UPPERCASE_LETTER    UnicodeType = 9
	UNICODE_SPACING_MARK        UnicodeType = 10
	UNICODE_ENCLOSING_MARK      UnicodeType = 11
	UNICODE_NON_SPACING_MARK    UnicodeType = 12
	UNICODE_DECIMAL_NUMBER      UnicodeType = 13
	UNICODE_LETTER_NUMBER       UnicodeType = 14
	UNICODE_OTHER_NUMBER        UnicodeType = 15
	UNICODE_CONNECT_PUNCTUATION UnicodeType = 16
	UNICODE_DASH_PUNCTUATION    UnicodeType = 17
	UNICODE_CLOSE_PUNCTUATION   UnicodeType = 18
	UNICODE_FINAL_PUNCTUATION   UnicodeType = 19
	UNICODE_INITIAL_PUNCTUATION UnicodeType = 20
	UNICODE_OTHER_PUNCTUATION   UnicodeType = 21
	UNICODE_OPEN_PUNCTUATION    UnicodeType = 22
	UNICODE_CURRENCY_SYMBOL     UnicodeType = 23
	UNICODE_MODIFIER_SYMBOL     UnicodeType = 24
	UNICODE_MATH_SYMBOL         UnicodeType = 25
	UNICODE_OTHER_SYMBOL        UnicodeType = 26
	UNICODE_LINE_SEPARATOR      UnicodeType = 27
	UNICODE_PARAGRAPH_SEPARATOR UnicodeType = 28
	UNICODE_SPACE_SEPARATOR     UnicodeType = 29
)

type UserDirectory C.GUserDirectory

const (
	USER_DIRECTORY_DESKTOP      UserDirectory = 0
	USER_DIRECTORY_DOCUMENTS    UserDirectory = 1
	USER_DIRECTORY_DOWNLOAD     UserDirectory = 2
	USER_DIRECTORY_MUSIC        UserDirectory = 3
	USER_DIRECTORY_PICTURES     UserDirectory = 4
	USER_DIRECTORY_PUBLIC_SHARE UserDirectory = 5
	USER_DIRECTORY_TEMPLATES    UserDirectory = 6
	USER_DIRECTORY_VIDEOS       UserDirectory = 7
	USER_N_DIRECTORIES          UserDirectory = 8
)

type VariantClass C.GVariantClass

const (
	VARIANT_CLASS_BOOLEAN     VariantClass = 98
	VARIANT_CLASS_BYTE        VariantClass = 121
	VARIANT_CLASS_INT16       VariantClass = 110
	VARIANT_CLASS_UINT16      VariantClass = 113
	VARIANT_CLASS_INT32       VariantClass = 105
	VARIANT_CLASS_UINT32      VariantClass = 117
	VARIANT_CLASS_INT64       VariantClass = 120
	VARIANT_CLASS_UINT64      VariantClass = 116
	VARIANT_CLASS_HANDLE      VariantClass = 104
	VARIANT_CLASS_DOUBLE      VariantClass = 100
	VARIANT_CLASS_STRING      VariantClass = 115
	VARIANT_CLASS_OBJECT_PATH VariantClass = 111
	VARIANT_CLASS_SIGNATURE   VariantClass = 103
	VARIANT_CLASS_VARIANT     VariantClass = 118
	VARIANT_CLASS_MAYBE       VariantClass = 109
	VARIANT_CLASS_ARRAY       VariantClass = 97
	VARIANT_CLASS_TUPLE       VariantClass = 40
	VARIANT_CLASS_DICT_ENTRY  VariantClass = 123
)

type VariantParseError C.GVariantParseError

const (
	VARIANT_PARSE_ERROR_FAILED                       VariantParseError = 0
	VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED          VariantParseError = 1
	VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE            VariantParseError = 2
	VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED       VariantParseError = 3
	VARIANT_PARSE_ERROR_INPUT_NOT_AT_END             VariantParseError = 4
	VARIANT_PARSE_ERROR_INVALID_CHARACTER            VariantParseError = 5
	VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING        VariantParseError = 6
	VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH          VariantParseError = 7
	VARIANT_PARSE_ERROR_INVALID_SIGNATURE            VariantParseError = 8
	VARIANT_PARSE_ERROR_INVALID_TYPE_STRING          VariantParseError = 9
	VARIANT_PARSE_ERROR_NO_COMMON_TYPE               VariantParseError = 10
	VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE          VariantParseError = 11
	VARIANT_PARSE_ERROR_NUMBER_TOO_BIG               VariantParseError = 12
	VARIANT_PARSE_ERROR_TYPE_ERROR                   VariantParseError = 13
	VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN             VariantParseError = 14
	VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD              VariantParseError = 15
	VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT VariantParseError = 16
	VARIANT_PARSE_ERROR_VALUE_EXPECTED               VariantParseError = 17
)

// Access is a wrapper around the C function g_access.
func Access(filename string, mode int32) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_mode := (C.int)(mode)

	retC := C.g_access(c_filename, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// AsciiDigitValue is a wrapper around the C function g_ascii_digit_value.
func AsciiDigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_digit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// AsciiDtostr is a wrapper around the C function g_ascii_dtostr.
func AsciiDtostr(buffer string, bufLen int32, d float64) string {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_buf_len := (C.gint)(bufLen)

	c_d := (C.gdouble)(d)

	retC := C.g_ascii_dtostr(c_buffer, c_buf_len, c_d)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

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
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AsciiStrcasecmp is a wrapper around the C function g_ascii_strcasecmp.
func AsciiStrcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_ascii_strcasecmp(c_s1, c_s2)
	retGo := (int32)(retC)

	return retGo
}

// AsciiStrdown is a wrapper around the C function g_ascii_strdown.
func AsciiStrdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_ascii_strdown(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

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
	retGo := (int32)(retC)

	return retGo
}

// AsciiStrtod is a wrapper around the C function g_ascii_strtod.
func AsciiStrtod(nptr string) (float64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_ascii_strtod(c_nptr, &c_endptr)
	retGo := (float64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// AsciiStrtoll is a wrapper around the C function g_ascii_strtoll.
func AsciiStrtoll(nptr string, base uint32) (int64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	c_base := (C.guint)(base)

	retC := C.g_ascii_strtoll(c_nptr, &c_endptr, c_base)
	retGo := (int64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// AsciiStrtoull is a wrapper around the C function g_ascii_strtoull.
func AsciiStrtoull(nptr string, base uint32) (uint64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	c_base := (C.guint)(base)

	retC := C.g_ascii_strtoull(c_nptr, &c_endptr, c_base)
	retGo := (uint64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

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

// AsciiTolower is a wrapper around the C function g_ascii_tolower.
func AsciiTolower(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_tolower(c_c)
	retGo := (rune)(retC)

	return retGo
}

// AsciiToupper is a wrapper around the C function g_ascii_toupper.
func AsciiToupper(c rune) rune {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_toupper(c_c)
	retGo := (rune)(retC)

	return retGo
}

// AsciiXdigitValue is a wrapper around the C function g_ascii_xdigit_value.
func AsciiXdigitValue(c rune) int32 {
	c_c := (C.gchar)(c)

	retC := C.g_ascii_xdigit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// AssertWarning is a wrapper around the C function g_assert_warning.
func AssertWarning(logDomain string, file string, line int32, prettyFunction string, expression string) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_pretty_function := C.CString(prettyFunction)
	defer C.free(unsafe.Pointer(c_pretty_function))

	c_expression := C.CString(expression)
	defer C.free(unsafe.Pointer(c_expression))

	C.g_assert_warning(c_log_domain, c_file, c_line, c_pretty_function, c_expression)

	return
}

// AssertionMessage is a wrapper around the C function g_assertion_message.
func AssertionMessage(domain string, file string, line int32, func_ string, message string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	C.g_assertion_message(c_domain, c_file, c_line, c_func, c_message)

	return
}

// Unsupported : g_assertion_message_cmpnum : unsupported parameter arg1 : no type generator for long double (long double) for param arg1

// AssertionMessageCmpstr is a wrapper around the C function g_assertion_message_cmpstr.
func AssertionMessageCmpstr(domain string, file string, line int32, func_ string, expr string, arg1 string, cmp string, arg2 string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_expr := C.CString(expr)
	defer C.free(unsafe.Pointer(c_expr))

	c_arg1 := C.CString(arg1)
	defer C.free(unsafe.Pointer(c_arg1))

	c_cmp := C.CString(cmp)
	defer C.free(unsafe.Pointer(c_cmp))

	c_arg2 := C.CString(arg2)
	defer C.free(unsafe.Pointer(c_arg2))

	C.g_assertion_message_cmpstr(c_domain, c_file, c_line, c_func, c_expr, c_arg1, c_cmp, c_arg2)

	return
}

// AssertionMessageError is a wrapper around the C function g_assertion_message_error.
func AssertionMessageError(domain string, file string, line int32, func_ string, expr string, error *Error, errorDomain Quark, errorCode int32) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_expr := C.CString(expr)
	defer C.free(unsafe.Pointer(c_expr))

	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.int)(errorCode)

	C.g_assertion_message_error(c_domain, c_file, c_line, c_func, c_expr, c_error, c_error_domain, c_error_code)

	return
}

// AssertionMessageExpr is a wrapper around the C function g_assertion_message_expr.
func AssertionMessageExpr(domain string, file string, line int32, func_ string, expr string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_expr := C.CString(expr)
	defer C.free(unsafe.Pointer(c_expr))

	C.g_assertion_message_expr(c_domain, c_file, c_line, c_func, c_expr)

	return
}

// Unsupported : g_atexit : unsupported parameter func : no type generator for VoidFunc (GVoidFunc) for param func

// Blacklisted : g_atomic_int_add

// AtomicIntAnd is a wrapper around the C function g_atomic_int_and.
func AtomicIntAnd(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_and(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// Blacklisted : g_atomic_int_compare_and_exchange

// Blacklisted : g_atomic_int_dec_and_test

// Blacklisted : g_atomic_int_exchange_and_add

// Blacklisted : g_atomic_int_get

// Blacklisted : g_atomic_int_inc

// AtomicIntOr is a wrapper around the C function g_atomic_int_or.
func AtomicIntOr(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_or(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// Blacklisted : g_atomic_int_set

// AtomicIntXor is a wrapper around the C function g_atomic_int_xor.
func AtomicIntXor(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_xor(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_base64_decode : array return type :

// Unsupported : g_base64_decode_inplace : unsupported parameter out_len : array length param out_len is pointer (gsize*)

// Unsupported : g_base64_decode_step : unsupported parameter out : output array param out

// Base64Encode is a wrapper around the C function g_base64_encode.
func Base64Encode(data []uint8) string {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_len := (C.gsize)(len(data))

	retC := C.g_base64_encode(c_data, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_base64_encode_close : unsupported parameter out : output array param out

// Unsupported : g_base64_encode_step : unsupported parameter out : output array param out

// Basename is a wrapper around the C function g_basename.
func Basename(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_basename(c_file_name)
	retGo := C.GoString(retC)

	return retGo
}

// BitLock is a wrapper around the C function g_bit_lock.
func BitLock(address int32, lockBit int32) {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	C.g_bit_lock(&c_address, c_lock_bit)

	return
}

// BitNthLsf is a wrapper around the C function g_bit_nth_lsf.
func BitNthLsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_lsf(c_mask, c_nth_bit)
	retGo := (int32)(retC)

	return retGo
}

// BitNthMsf is a wrapper around the C function g_bit_nth_msf.
func BitNthMsf(mask uint64, nthBit int32) int32 {
	c_mask := (C.gulong)(mask)

	c_nth_bit := (C.gint)(nthBit)

	retC := C.g_bit_nth_msf(c_mask, c_nth_bit)
	retGo := (int32)(retC)

	return retGo
}

// BitStorage is a wrapper around the C function g_bit_storage.
func BitStorage(number uint64) uint32 {
	c_number := (C.gulong)(number)

	retC := C.g_bit_storage(c_number)
	retGo := (uint32)(retC)

	return retGo
}

// BitTrylock is a wrapper around the C function g_bit_trylock.
func BitTrylock(address int32, lockBit int32) bool {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	retC := C.g_bit_trylock(&c_address, c_lock_bit)
	retGo := retC == C.TRUE

	return retGo
}

// BitUnlock is a wrapper around the C function g_bit_unlock.
func BitUnlock(address int32, lockBit int32) {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	C.g_bit_unlock(&c_address, c_lock_bit)

	return
}

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// BuildFilenamev is a wrapper around the C function g_build_filenamev.
func BuildFilenamev(args []string) string {
	c_args_array := make([]*C.gchar, len(args)+1, len(args)+1)
	for i, item := range args {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_args_array[i] = c
	}
	c_args_array[len(args)] = nil
	c_args_arrayPtr := &c_args_array[0]
	c_args := (**C.gchar)(unsafe.Pointer(c_args_arrayPtr))

	retC := C.g_build_filenamev(c_args)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_build_path : unsupported parameter ... : varargs

// BuildPathv is a wrapper around the C function g_build_pathv.
func BuildPathv(separator string, args []string) string {
	c_separator := C.CString(separator)
	defer C.free(unsafe.Pointer(c_separator))

	c_args_array := make([]*C.gchar, len(args)+1, len(args)+1)
	for i, item := range args {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_args_array[i] = c
	}
	c_args_array[len(args)] = nil
	c_args_arrayPtr := &c_args_array[0]
	c_args := (**C.gchar)(unsafe.Pointer(c_args_arrayPtr))

	retC := C.g_build_pathv(c_separator, c_args)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_byte_array_free

// Unsupported : g_byte_array_new : array return type :

// Unsupported : g_byte_array_new_take : array return type :

// Chdir is a wrapper around the C function g_chdir.
func Chdir(path string) int32 {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_chdir(c_path)
	retGo := (int32)(retC)

	return retGo
}

// CheckVersion is a wrapper around the C function glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	c_required_major := (C.guint)(requiredMajor)

	c_required_minor := (C.guint)(requiredMinor)

	c_required_micro := (C.guint)(requiredMicro)

	retC := C.glib_check_version(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// ChildWatchSourceNew is a wrapper around the C function g_child_watch_source_new.
func ChildWatchSourceNew(pid Pid) *Source {
	c_pid := (C.GPid)(pid)

	retC := C.g_child_watch_source_new(c_pid)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ClearError is a wrapper around the C function g_clear_error.
func ClearError() error {
	var cThrowableError *C.GError

	C.g_clear_error(&cThrowableError)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return goError
}

// Unsupported : g_clear_pointer : unsupported parameter pp : no type generator for gpointer (gpointer*) for param pp

// ComputeChecksumForBytes is a wrapper around the C function g_compute_checksum_for_bytes.
func ComputeChecksumForBytes(checksumType ChecksumType, data *Bytes) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_data := (*C.GBytes)(C.NULL)
	if data != nil {
		c_data = (*C.GBytes)(data.ToC())
	}

	retC := C.g_compute_checksum_for_bytes(c_checksum_type, c_data)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ComputeChecksumForData is a wrapper around the C function g_compute_checksum_for_data.
func ComputeChecksumForData(checksumType ChecksumType, data []uint8) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	retC := C.g_compute_checksum_for_data(c_checksum_type, c_data, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ComputeChecksumForString is a wrapper around the C function g_compute_checksum_for_string.
func ComputeChecksumForString(checksumType ChecksumType, str string) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(len(str))

	retC := C.g_compute_checksum_for_string(c_checksum_type, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ComputeHmacForData is a wrapper around the C function g_compute_hmac_for_data.
func ComputeHmacForData(digestType ChecksumType, key []uint8, data []uint8) string {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key_array := make([]C.guchar, len(key)+1, len(key)+1)
	for i, item := range key {
		c := (C.guchar)(item)
		c_key_array[i] = c
	}
	c_key_array[len(key)] = 0
	c_key_arrayPtr := &c_key_array[0]
	c_key := (*C.guchar)(unsafe.Pointer(c_key_arrayPtr))

	c_key_len := (C.gsize)(len(key))

	c_data_array := make([]C.guchar, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guchar)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	retC := C.g_compute_hmac_for_data(c_digest_type, c_key, c_key_len, c_data, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ComputeHmacForString is a wrapper around the C function g_compute_hmac_for_string.
func ComputeHmacForString(digestType ChecksumType, key []uint8, str string) string {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key_array := make([]C.guchar, len(key)+1, len(key)+1)
	for i, item := range key {
		c := (C.guchar)(item)
		c_key_array[i] = c
	}
	c_key_array[len(key)] = 0
	c_key_arrayPtr := &c_key_array[0]
	c_key := (*C.guchar)(unsafe.Pointer(c_key_arrayPtr))

	c_key_len := (C.gsize)(len(key))

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(len(str))

	retC := C.g_compute_hmac_for_string(c_digest_type, c_key, c_key_len, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_convert : array return type :

// ConvertErrorQuark is a wrapper around the C function g_convert_error_quark.
func ConvertErrorQuark() Quark {
	retC := C.g_convert_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_convert_with_fallback : array return type :

// Unsupported : g_convert_with_iconv : unsupported parameter converter : Blacklisted record : GIConv

// DatalistClear is a wrapper around the C function g_datalist_clear.
func DatalistClear(datalist *Data) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	C.g_datalist_clear(c_datalist)

	return
}

// Unsupported : g_datalist_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// Unsupported : g_datalist_get_data : no return generator

// DatalistGetFlags is a wrapper around the C function g_datalist_get_flags.
func DatalistGetFlags(datalist *Data) uint32 {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	retC := C.g_datalist_get_flags(c_datalist)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_datalist_id_dup_data : unsupported parameter dup_func : no type generator for DuplicateFunc (GDuplicateFunc) for param dup_func

// Unsupported : g_datalist_id_get_data : no return generator

// Unsupported : g_datalist_id_remove_no_notify : no return generator

// Unsupported : g_datalist_id_replace_data : unsupported parameter oldval : no type generator for gpointer (gpointer) for param oldval

// Unsupported : g_datalist_id_set_data_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// DatalistInit is a wrapper around the C function g_datalist_init.
func DatalistInit(datalist *Data) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	C.g_datalist_init(c_datalist)

	return
}

// DatalistSetFlags is a wrapper around the C function g_datalist_set_flags.
func DatalistSetFlags(datalist *Data, flags uint32) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_flags := (C.guint)(flags)

	C.g_datalist_set_flags(c_datalist, c_flags)

	return
}

// DatalistUnsetFlags is a wrapper around the C function g_datalist_unset_flags.
func DatalistUnsetFlags(datalist *Data, flags uint32) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_flags := (C.guint)(flags)

	C.g_datalist_unset_flags(c_datalist, c_flags)

	return
}

// Unsupported : g_dataset_destroy : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_foreach : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_get_data : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_remove_no_notify : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_set_data_full : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_date_time_compare : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1

// Unsupported : g_date_time_equal : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1

// Unsupported : g_date_time_hash : unsupported parameter datetime : no type generator for gpointer (gconstpointer) for param datetime

// Dcgettext is a wrapper around the C function g_dcgettext.
func Dcgettext(domain string, msgid string, category int32) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_category := (C.gint)(category)

	retC := C.g_dcgettext(c_domain, c_msgid, c_category)
	retGo := C.GoString(retC)

	return retGo
}

// Dgettext is a wrapper around the C function g_dgettext.
func Dgettext(domain string, msgid string) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	retC := C.g_dgettext(c_domain, c_msgid)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_direct_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_direct_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Dngettext is a wrapper around the C function g_dngettext.
func Dngettext(domain string, msgid string, msgidPlural string, n uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgid_plural := C.CString(msgidPlural)
	defer C.free(unsafe.Pointer(c_msgid_plural))

	c_n := (C.gulong)(n)

	retC := C.g_dngettext(c_domain, c_msgid, c_msgid_plural, c_n)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_double_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_double_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Dpgettext is a wrapper around the C function g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgctxtid := C.CString(msgctxtid)
	defer C.free(unsafe.Pointer(c_msgctxtid))

	c_msgidoffset := (C.gsize)(msgidoffset)

	retC := C.g_dpgettext(c_domain, c_msgctxtid, c_msgidoffset)
	retGo := C.GoString(retC)

	return retGo
}

// Dpgettext2 is a wrapper around the C function g_dpgettext2.
func Dpgettext2(domain string, context string, msgid string) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_context := C.CString(context)
	defer C.free(unsafe.Pointer(c_context))

	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	retC := C.g_dpgettext2(c_domain, c_context, c_msgid)
	retGo := C.GoString(retC)

	return retGo
}

// EnvironGetenv is a wrapper around the C function g_environ_getenv.
func EnvironGetenv(envp []string, variable string) string {
	c_envp_array := make([]*C.gchar, len(envp)+1, len(envp)+1)
	for i, item := range envp {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_envp_array[i] = c
	}
	c_envp_array[len(envp)] = nil
	c_envp_arrayPtr := &c_envp_array[0]
	c_envp := (**C.gchar)(unsafe.Pointer(c_envp_arrayPtr))

	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_environ_getenv(c_envp, c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// EnvironSetenv is a wrapper around the C function g_environ_setenv.
func EnvironSetenv(envp []string, variable string, value string, overwrite bool) []string {
	c_envp_array := make([]*C.gchar, len(envp)+1, len(envp)+1)
	for i, item := range envp {
		c := C.CString(item)
		c_envp_array[i] = c
	}
	c_envp_array[len(envp)] = nil
	c_envp_arrayPtr := &c_envp_array[0]
	c_envp := (**C.gchar)(unsafe.Pointer(c_envp_arrayPtr))

	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	retC := C.g_environ_setenv(c_envp, c_variable, c_value, c_overwrite)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// EnvironUnsetenv is a wrapper around the C function g_environ_unsetenv.
func EnvironUnsetenv(envp []string, variable string) []string {
	c_envp_array := make([]*C.gchar, len(envp)+1, len(envp)+1)
	for i, item := range envp {
		c := C.CString(item)
		c_envp_array[i] = c
	}
	c_envp_array[len(envp)] = nil
	c_envp_arrayPtr := &c_envp_array[0]
	c_envp := (**C.gchar)(unsafe.Pointer(c_envp_arrayPtr))

	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_environ_unsetenv(c_envp, c_variable)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// FileErrorFromErrno is a wrapper around the C function g_file_error_from_errno.
func FileErrorFromErrno(errNo int32) FileError {
	c_err_no := (C.gint)(errNo)

	retC := C.g_file_error_from_errno(c_err_no)
	retGo := (FileError)(retC)

	return retGo
}

// FileErrorQuark is a wrapper around the C function g_file_error_quark.
func FileErrorQuark() Quark {
	retC := C.g_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

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

// FileReadLink is a wrapper around the C function g_file_read_link.
func FileReadLink(filename string) (string, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_file_read_link(c_filename, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// FileSetContents is a wrapper around the C function g_file_set_contents.
func FileSetContents(filename string, contents []uint8) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_contents_array := make([]C.guint8, len(contents)+1, len(contents)+1)
	for i, item := range contents {
		c := (C.guint8)(item)
		c_contents_array[i] = c
	}
	c_contents_array[len(contents)] = 0
	c_contents_arrayPtr := &c_contents_array[0]
	c_contents := (*C.gchar)(unsafe.Pointer(c_contents_arrayPtr))

	c_length := (C.gssize)(len(contents))

	var cThrowableError *C.GError

	retC := C.g_file_set_contents(c_filename, c_contents, c_length, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// FileTest is a wrapper around the C function g_file_test.
func FileTest(filename string, test GFileTest) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_test := (C.GFileTest)(test)

	retC := C.g_file_test(c_filename, c_test)
	retGo := retC == C.TRUE

	return retGo
}

// FilenameDisplayBasename is a wrapper around the C function g_filename_display_basename.
func FilenameDisplayBasename(filename string) string {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_filename_display_basename(c_filename)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FilenameDisplayName is a wrapper around the C function g_filename_display_name.
func FilenameDisplayName(filename string) string {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_filename_display_name(c_filename)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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

// FilenameFromUtf8 is a wrapper around the C function g_filename_from_utf8.
func FilenameFromUtf8(utf8string string, len int64) (string, uint64, uint64, error) {
	c_utf8string := C.CString(utf8string)
	defer C.free(unsafe.Pointer(c_utf8string))

	c_len := (C.gssize)(len)

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_filename_from_utf8(c_utf8string, c_len, &c_bytes_read, &c_bytes_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesRead, bytesWritten, goError
}

// FilenameToUri is a wrapper around the C function g_filename_to_uri.
func FilenameToUri(filename string, hostname string) (string, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	var cThrowableError *C.GError

	retC := C.g_filename_to_uri(c_filename, c_hostname, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// FilenameToUtf8 is a wrapper around the C function g_filename_to_utf8.
func FilenameToUtf8(opsysstring string, len int64) (string, uint64, uint64, error) {
	c_opsysstring := C.CString(opsysstring)
	defer C.free(unsafe.Pointer(c_opsysstring))

	c_len := (C.gssize)(len)

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_filename_to_utf8(c_opsysstring, c_len, &c_bytes_read, &c_bytes_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesRead, bytesWritten, goError
}

// FindProgramInPath is a wrapper around the C function g_find_program_in_path.
func FindProgramInPath(program string) string {
	c_program := C.CString(program)
	defer C.free(unsafe.Pointer(c_program))

	retC := C.g_find_program_in_path(c_program)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FormatSize is a wrapper around the C function g_format_size.
func FormatSize(size uint64) string {
	c_size := (C.guint64)(size)

	retC := C.g_format_size(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FormatSizeForDisplay is a wrapper around the C function g_format_size_for_display.
func FormatSizeForDisplay(size int64) string {
	c_size := (C.goffset)(size)

	retC := C.g_format_size_for_display(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FormatSizeFull is a wrapper around the C function g_format_size_full.
func FormatSizeFull(size uint64, flags FormatSizeFlags) string {
	c_size := (C.guint64)(size)

	c_flags := (C.GFormatSizeFlags)(flags)

	retC := C.g_format_size_full(c_size, c_flags)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_fprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Unsupported : g_free : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// GetApplicationName is a wrapper around the C function g_get_application_name.
func GetApplicationName() string {
	retC := C.g_get_application_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetCharset is a wrapper around the C function g_get_charset.
func GetCharset() (bool, string) {
	var c_charset *C.char

	retC := C.g_get_charset(&c_charset)
	retGo := retC == C.TRUE

	charset := C.GoString(c_charset)

	return retGo, charset
}

// GetCodeset is a wrapper around the C function g_get_codeset.
func GetCodeset() string {
	retC := C.g_get_codeset()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentDir is a wrapper around the C function g_get_current_dir.
func GetCurrentDir() string {
	retC := C.g_get_current_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentTime is a wrapper around the C function g_get_current_time.
func GetCurrentTime(result *TimeVal) {
	c_result := (*C.GTimeVal)(C.NULL)
	if result != nil {
		c_result = (*C.GTimeVal)(result.ToC())
	}

	C.g_get_current_time(c_result)

	return
}

// GetEnviron is a wrapper around the C function g_get_environ.
func GetEnviron() []string {
	retC := C.g_get_environ()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// GetHomeDir is a wrapper around the C function g_get_home_dir.
func GetHomeDir() string {
	retC := C.g_get_home_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetHostName is a wrapper around the C function g_get_host_name.
func GetHostName() string {
	retC := C.g_get_host_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetLanguageNames is a wrapper around the C function g_get_language_names.
func GetLanguageNames() []string {
	retC := C.g_get_language_names()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetLocaleVariants is a wrapper around the C function g_get_locale_variants.
func GetLocaleVariants(locale string) []string {
	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	retC := C.g_get_locale_variants(c_locale)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetMonotonicTime is a wrapper around the C function g_get_monotonic_time.
func GetMonotonicTime() int64 {
	retC := C.g_get_monotonic_time()
	retGo := (int64)(retC)

	return retGo
}

// GetPrgname is a wrapper around the C function g_get_prgname.
func GetPrgname() string {
	retC := C.g_get_prgname()
	retGo := C.GoString(retC)

	return retGo
}

// GetRealName is a wrapper around the C function g_get_real_name.
func GetRealName() string {
	retC := C.g_get_real_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetRealTime is a wrapper around the C function g_get_real_time.
func GetRealTime() int64 {
	retC := C.g_get_real_time()
	retGo := (int64)(retC)

	return retGo
}

// GetSystemConfigDirs is a wrapper around the C function g_get_system_config_dirs.
func GetSystemConfigDirs() []string {
	retC := C.g_get_system_config_dirs()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetSystemDataDirs is a wrapper around the C function g_get_system_data_dirs.
func GetSystemDataDirs() []string {
	retC := C.g_get_system_data_dirs()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetTmpDir is a wrapper around the C function g_get_tmp_dir.
func GetTmpDir() string {
	retC := C.g_get_tmp_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserCacheDir is a wrapper around the C function g_get_user_cache_dir.
func GetUserCacheDir() string {
	retC := C.g_get_user_cache_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserConfigDir is a wrapper around the C function g_get_user_config_dir.
func GetUserConfigDir() string {
	retC := C.g_get_user_config_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserDataDir is a wrapper around the C function g_get_user_data_dir.
func GetUserDataDir() string {
	retC := C.g_get_user_data_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserName is a wrapper around the C function g_get_user_name.
func GetUserName() string {
	retC := C.g_get_user_name()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserRuntimeDir is a wrapper around the C function g_get_user_runtime_dir.
func GetUserRuntimeDir() string {
	retC := C.g_get_user_runtime_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserSpecialDir is a wrapper around the C function g_get_user_special_dir.
func GetUserSpecialDir(directory UserDirectory) string {
	c_directory := (C.GUserDirectory)(directory)

	retC := C.g_get_user_special_dir(c_directory)
	retGo := C.GoString(retC)

	return retGo
}

// Getenv is a wrapper around the C function g_getenv.
func Getenv(variable string) string {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_getenv(c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_hash_table_add : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_contains : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key

// Unsupported : g_hash_table_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// HostnameIsAsciiEncoded is a wrapper around the C function g_hostname_is_ascii_encoded.
func HostnameIsAsciiEncoded(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ascii_encoded(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameIsIpAddress is a wrapper around the C function g_hostname_is_ip_address.
func HostnameIsIpAddress(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ip_address(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameIsNonAscii is a wrapper around the C function g_hostname_is_non_ascii.
func HostnameIsNonAscii(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_non_ascii(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameToAscii is a wrapper around the C function g_hostname_to_ascii.
func HostnameToAscii(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_ascii(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// HostnameToUnicode is a wrapper around the C function g_hostname_to_unicode.
func HostnameToUnicode(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_unicode(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_remove_by_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// IdleSourceNew is a wrapper around the C function g_idle_source_new.
func IdleSourceNew() *Source {
	retC := C.g_idle_source_new()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_int64_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_int64_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_int_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_int_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// InternStaticString is a wrapper around the C function g_intern_static_string.
func InternStaticString(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_static_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// InternString is a wrapper around the C function g_intern_string.
func InternString(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_io_add_watch : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// Unsupported : g_io_add_watch_full : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// IoCreateWatch is a wrapper around the C function g_io_create_watch.
func IoCreateWatch(channel *IOChannel, condition IOCondition) *Source {
	c_channel := (*C.GIOChannel)(C.NULL)
	if channel != nil {
		c_channel = (*C.GIOChannel)(channel.ToC())
	}

	c_condition := (C.GIOCondition)(condition)

	retC := C.g_io_create_watch(c_channel, c_condition)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Listenv is a wrapper around the C function g_listenv.
func Listenv() []string {
	retC := C.g_listenv()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// Unsupported : g_locale_from_utf8 : array return type :

// LocaleToUtf8 is a wrapper around the C function g_locale_to_utf8.
func LocaleToUtf8(opsysstring []uint8) (string, uint64, uint64, error) {
	c_opsysstring_array := make([]C.guint8, len(opsysstring)+1, len(opsysstring)+1)
	for i, item := range opsysstring {
		c := (C.guint8)(item)
		c_opsysstring_array[i] = c
	}
	c_opsysstring_array[len(opsysstring)] = 0
	c_opsysstring_arrayPtr := &c_opsysstring_array[0]
	c_opsysstring := (*C.gchar)(unsafe.Pointer(c_opsysstring_arrayPtr))

	c_len := (C.gssize)(len(opsysstring))

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_locale_to_utf8(c_opsysstring, c_len, &c_bytes_read, &c_bytes_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesRead, bytesWritten, goError
}

// Log is a wrapper around the C function g_log.
func Log(logDomain string, logLevel LogLevelFlags, format string, args ...interface{}) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_log_level := (C.GLogLevelFlags)(logLevel)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_log(c_log_domain, c_log_level, c_format)

	return
}

// Unsupported : g_log_default_handler : unsupported parameter unused_data : no type generator for gpointer (gpointer) for param unused_data

// LogRemoveHandler is a wrapper around the C function g_log_remove_handler.
func LogRemoveHandler(logDomain string, handlerId uint32) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_handler_id := (C.guint)(handlerId)

	C.g_log_remove_handler(c_log_domain, c_handler_id)

	return
}

// LogSetAlwaysFatal is a wrapper around the C function g_log_set_always_fatal.
func LogSetAlwaysFatal(fatalMask LogLevelFlags) LogLevelFlags {
	c_fatal_mask := (C.GLogLevelFlags)(fatalMask)

	retC := C.g_log_set_always_fatal(c_fatal_mask)
	retGo := (LogLevelFlags)(retC)

	return retGo
}

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// LogSetFatalMask is a wrapper around the C function g_log_set_fatal_mask.
func LogSetFatalMask(logDomain string, fatalMask LogLevelFlags) LogLevelFlags {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_fatal_mask := (C.GLogLevelFlags)(fatalMask)

	retC := C.g_log_set_fatal_mask(c_log_domain, c_fatal_mask)
	retGo := (LogLevelFlags)(retC)

	return retGo
}

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Blacklisted : g_log_structured_standard

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list (va_list) for param args

// MainCurrentSource is a wrapper around the C function g_main_current_source.
func MainCurrentSource() *Source {
	retC := C.g_main_current_source()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MainDepth is a wrapper around the C function g_main_depth.
func MainDepth() int32 {
	retC := C.g_main_depth()
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_malloc : no return generator

// Unsupported : g_malloc0 : no return generator

// Unsupported : g_malloc0_n : no return generator

// Unsupported : g_malloc_n : no return generator

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// MarkupErrorQuark is a wrapper around the C function g_markup_error_quark.
func MarkupErrorQuark() Quark {
	retC := C.g_markup_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// MarkupEscapeText is a wrapper around the C function g_markup_escape_text.
func MarkupEscapeText(text string) string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gssize)(len(text))

	retC := C.g_markup_escape_text(c_text, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// MarkupPrintfEscaped is a wrapper around the C function g_markup_printf_escaped.
func MarkupPrintfEscaped(format string, args ...interface{}) string {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_markup_printf_escaped(c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list (va_list) for param args

// MemIsSystemMalloc is a wrapper around the C function g_mem_is_system_malloc.
func MemIsSystemMalloc() bool {
	retC := C.g_mem_is_system_malloc()
	retGo := retC == C.TRUE

	return retGo
}

// MemProfile is a wrapper around the C function g_mem_profile.
func MemProfile() {
	C.g_mem_profile()

	return
}

// MemSetVtable is a wrapper around the C function g_mem_set_vtable.
func MemSetVtable(vtable *MemVTable) {
	c_vtable := (*C.GMemVTable)(C.NULL)
	if vtable != nil {
		c_vtable = (*C.GMemVTable)(vtable.ToC())
	}

	C.g_mem_set_vtable(c_vtable)

	return
}

// Unsupported : g_memdup : unsupported parameter mem : no type generator for gpointer (gconstpointer) for param mem

// MkdirWithParents is a wrapper around the C function g_mkdir_with_parents.
func MkdirWithParents(pathname string, mode int32) int32 {
	c_pathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(c_pathname))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdir_with_parents(c_pathname, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// Mkdtemp is a wrapper around the C function g_mkdtemp.
func Mkdtemp(tmpl string) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkdtemp(c_tmpl)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// MkdtempFull is a wrapper around the C function g_mkdtemp_full.
func MkdtempFull(tmpl string, mode int32) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdtemp_full(c_tmpl, c_mode)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Mkstemp is a wrapper around the C function g_mkstemp.
func Mkstemp(tmpl string) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkstemp(c_tmpl)
	retGo := (int32)(retC)

	return retGo
}

// MkstempFull is a wrapper around the C function g_mkstemp_full.
func MkstempFull(tmpl string, flags int32, mode int32) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_flags := (C.gint)(flags)

	c_mode := (C.gint)(mode)

	retC := C.g_mkstemp_full(c_tmpl, c_flags, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : no type generator for gpointer (gpointer*) for param nullify_location

// Blacklisted : g_number_parser_error_quark

// OnErrorQuery is a wrapper around the C function g_on_error_query.
func OnErrorQuery(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_query(c_prg_name)

	return
}

// OnErrorStackTrace is a wrapper around the C function g_on_error_stack_trace.
func OnErrorStackTrace(prgName string) {
	c_prg_name := C.CString(prgName)
	defer C.free(unsafe.Pointer(c_prg_name))

	C.g_on_error_stack_trace(c_prg_name)

	return
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location

// OptionErrorQuark is a wrapper around the C function g_option_error_quark.
func OptionErrorQuark() Quark {
	retC := C.g_option_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_parse_debug_string : unsupported parameter keys :

// PathGetBasename is a wrapper around the C function g_path_get_basename.
func PathGetBasename(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_get_basename(c_file_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// PathGetDirname is a wrapper around the C function g_path_get_dirname.
func PathGetDirname(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_get_dirname(c_file_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// PathIsAbsolute is a wrapper around the C function g_path_is_absolute.
func PathIsAbsolute(fileName string) bool {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_is_absolute(c_file_name)
	retGo := retC == C.TRUE

	return retGo
}

// PathSkipRoot is a wrapper around the C function g_path_skip_root.
func PathSkipRoot(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_path_skip_root(c_file_name)
	retGo := C.GoString(retC)

	return retGo
}

// PatternMatch is a wrapper around the C function g_pattern_match.
func PatternMatch(pspec *PatternSpec, stringLength uint32, string_ string, stringReversed string) bool {
	c_pspec := (*C.GPatternSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GPatternSpec)(pspec.ToC())
	}

	c_string_length := (C.guint)(stringLength)

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_string_reversed := C.CString(stringReversed)
	defer C.free(unsafe.Pointer(c_string_reversed))

	retC := C.g_pattern_match(c_pspec, c_string_length, c_string, c_string_reversed)
	retGo := retC == C.TRUE

	return retGo
}

// PatternMatchSimple is a wrapper around the C function g_pattern_match_simple.
func PatternMatchSimple(pattern string, string_ string) bool {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_pattern_match_simple(c_pattern, c_string)
	retGo := retC == C.TRUE

	return retGo
}

// PatternMatchString is a wrapper around the C function g_pattern_match_string.
func PatternMatchString(pspec *PatternSpec, string_ string) bool {
	c_pspec := (*C.GPatternSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GPatternSpec)(pspec.ToC())
	}

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_pattern_match_string(c_pspec, c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Poll is a wrapper around the C function g_poll.
func Poll(fds *PollFD, nfds uint32, timeout int32) int32 {
	c_fds := (*C.GPollFD)(C.NULL)
	if fds != nil {
		c_fds = (*C.GPollFD)(fds.ToC())
	}

	c_nfds := (C.guint)(nfds)

	c_timeout := (C.gint)(timeout)

	retC := C.g_poll(c_fds, c_nfds, c_timeout)
	retGo := (int32)(retC)

	return retGo
}

// PrefixError is a wrapper around the C function g_prefix_error.
func PrefixError(err *Error, format string, args ...interface{}) {
	c_err := (**C.GError)(C.NULL)
	if err != nil {
		c_err = (**C.GError)(err.ToC())
	}

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_prefix_error(c_err, c_format)

	return
}

// Print is a wrapper around the C function g_print.
func Print(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_print(c_format)

	return
}

// Printerr is a wrapper around the C function g_printerr.
func Printerr(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_printerr(c_format)

	return
}

// Printf is a wrapper around the C function g_printf.
func Printf(format string, args ...interface{}) int32 {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_printf(c_format)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list (va_list) for param args

// PropagateError is a wrapper around the C function g_propagate_error.
func PropagateError(src *Error) *Error {
	var c_dest *C.GError

	c_src := (*C.GError)(C.NULL)
	if src != nil {
		c_src = (*C.GError)(src.ToC())
	}

	C.g_propagate_error(&c_dest, c_src)

	dest := ErrorNewFromC(unsafe.Pointer(c_dest))

	return dest
}

// PropagatePrefixedError is a wrapper around the C function g_propagate_prefixed_error.
func PropagatePrefixedError(dest *Error, src *Error, format string, args ...interface{}) {
	c_dest := (**C.GError)(C.NULL)
	if dest != nil {
		c_dest = (**C.GError)(dest.ToC())
	}

	c_src := (*C.GError)(C.NULL)
	if src != nil {
		c_src = (*C.GError)(src.ToC())
	}

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_propagate_prefixed_error(c_dest, c_src, c_format)

	return
}

// Unsupported : g_qsort_with_data : unsupported parameter pbase : no type generator for gpointer (gconstpointer) for param pbase

// QuarkFromStaticString is a wrapper around the C function g_quark_from_static_string.
func QuarkFromStaticString(string_ string) Quark {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_from_static_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

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

// QuarkTryString is a wrapper around the C function g_quark_try_string.
func QuarkTryString(string_ string) Quark {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_quark_try_string(c_string)
	retGo := (Quark)(retC)

	return retGo
}

// RandomDouble is a wrapper around the C function g_random_double.
func RandomDouble() float64 {
	retC := C.g_random_double()
	retGo := (float64)(retC)

	return retGo
}

// RandomDoubleRange is a wrapper around the C function g_random_double_range.
func RandomDoubleRange(begin float64, end float64) float64 {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	retC := C.g_random_double_range(c_begin, c_end)
	retGo := (float64)(retC)

	return retGo
}

// RandomInt is a wrapper around the C function g_random_int.
func RandomInt() uint32 {
	retC := C.g_random_int()
	retGo := (uint32)(retC)

	return retGo
}

// RandomIntRange is a wrapper around the C function g_random_int_range.
func RandomIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_random_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// RandomSetSeed is a wrapper around the C function g_random_set_seed.
func RandomSetSeed(seed uint32) {
	c_seed := (C.guint32)(seed)

	C.g_random_set_seed(c_seed)

	return
}

// Unsupported : g_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// ReloadUserSpecialDirsCache is a wrapper around the C function g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()

	return
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

	return
}

// Rmdir is a wrapper around the C function g_rmdir.
func Rmdir(filename string) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_rmdir(c_filename)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_sequence_get : no return generator

// Unsupported : g_sequence_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_set : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// SetApplicationName is a wrapper around the C function g_set_application_name.
func SetApplicationName(applicationName string) {
	c_application_name := C.CString(applicationName)
	defer C.free(unsafe.Pointer(c_application_name))

	C.g_set_application_name(c_application_name)

	return
}

// SetError is a wrapper around the C function g_set_error.
func SetError(domain Quark, code int32, format string, args ...interface{}) *Error {
	var c_err *C.GError

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_set_error(&c_err, c_domain, c_code, c_format)

	err := ErrorNewFromC(unsafe.Pointer(c_err))

	return err
}

// SetErrorLiteral is a wrapper around the C function g_set_error_literal.
func SetErrorLiteral(domain Quark, code int32, message string) *Error {
	var c_err *C.GError

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	C.g_set_error_literal(&c_err, c_domain, c_code, c_message)

	err := ErrorNewFromC(unsafe.Pointer(c_err))

	return err
}

// SetPrgname is a wrapper around the C function g_set_prgname.
func SetPrgname(prgname string) {
	c_prgname := C.CString(prgname)
	defer C.free(unsafe.Pointer(c_prgname))

	C.g_set_prgname(c_prgname)

	return
}

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Setenv is a wrapper around the C function g_setenv.
func Setenv(variable string, value string, overwrite bool) bool {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	retC := C.g_setenv(c_variable, c_value, c_overwrite)
	retGo := retC == C.TRUE

	return retGo
}

// ShellErrorQuark is a wrapper around the C function g_shell_error_quark.
func ShellErrorQuark() Quark {
	retC := C.g_shell_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_shell_parse_argv : unsupported parameter argcp : array length param argcp is pointer (gint*)

// ShellQuote is a wrapper around the C function g_shell_quote.
func ShellQuote(unquotedString string) string {
	c_unquoted_string := C.CString(unquotedString)
	defer C.free(unsafe.Pointer(c_unquoted_string))

	retC := C.g_shell_quote(c_unquoted_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ShellUnquote is a wrapper around the C function g_shell_unquote.
func ShellUnquote(quotedString string) (string, error) {
	c_quoted_string := C.CString(quotedString)
	defer C.free(unsafe.Pointer(c_quoted_string))

	var cThrowableError *C.GError

	retC := C.g_shell_unquote(c_quoted_string, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_slice_alloc : no return generator

// Unsupported : g_slice_alloc0 : no return generator

// Unsupported : g_slice_copy : unsupported parameter mem_block : no type generator for gpointer (gconstpointer) for param mem_block

// Unsupported : g_slice_free1 : unsupported parameter mem_block : no type generator for gpointer (gpointer) for param mem_block

// Unsupported : g_slice_free_chain_with_offset : unsupported parameter mem_chain : no type generator for gpointer (gpointer) for param mem_chain

// SliceGetConfig is a wrapper around the C function g_slice_get_config.
func SliceGetConfig(ckey SliceConfig) int64 {
	c_ckey := (C.GSliceConfig)(ckey)

	retC := C.g_slice_get_config(c_ckey)
	retGo := (int64)(retC)

	return retGo
}

// Blacklisted : g_slice_get_config_state

// SliceSetConfig is a wrapper around the C function g_slice_set_config.
func SliceSetConfig(ckey SliceConfig, value int64) {
	c_ckey := (C.GSliceConfig)(ckey)

	c_value := (C.gint64)(value)

	C.g_slice_set_config(c_ckey, c_value)

	return
}

// Snprintf is a wrapper around the C function g_snprintf.
func Snprintf(string_ string, n uint64, format string, args ...interface{}) int32 {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_n := (C.gulong)(n)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_snprintf(c_string, c_n, c_format)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_source_remove_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_source_remove_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// SpacedPrimesClosest is a wrapper around the C function g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	c_num := (C.guint)(num)

	retC := C.g_spaced_primes_closest(c_num)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_spawn_async : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_spawn_async_with_pipes : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// SpawnCheckExitStatus is a wrapper around the C function g_spawn_check_exit_status.
func SpawnCheckExitStatus(exitStatus int32) (bool, error) {
	c_exit_status := (C.gint)(exitStatus)

	var cThrowableError *C.GError

	retC := C.g_spawn_check_exit_status(c_exit_status, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SpawnClosePid is a wrapper around the C function g_spawn_close_pid.
func SpawnClosePid(pid Pid) {
	c_pid := (C.GPid)(pid)

	C.g_spawn_close_pid(c_pid)

	return
}

// SpawnCommandLineAsync is a wrapper around the C function g_spawn_command_line_async.
func SpawnCommandLineAsync(commandLine string) (bool, error) {
	c_command_line := C.CString(commandLine)
	defer C.free(unsafe.Pointer(c_command_line))

	var cThrowableError *C.GError

	retC := C.g_spawn_command_line_async(c_command_line, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : output array param standard_output

// SpawnErrorQuark is a wrapper around the C function g_spawn_error_quark.
func SpawnErrorQuark() Quark {
	retC := C.g_spawn_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// SpawnExitErrorQuark is a wrapper around the C function g_spawn_exit_error_quark.
func SpawnExitErrorQuark() Quark {
	retC := C.g_spawn_exit_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Unsupported : g_spawn_sync : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Sprintf is a wrapper around the C function g_sprintf.
func Sprintf(string_ string, format string, args ...interface{}) int32 {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_sprintf(c_string, c_format)
	retGo := (int32)(retC)

	return retGo
}

// Stpcpy is a wrapper around the C function g_stpcpy.
func Stpcpy(dest string, src string) string {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	retC := C.g_stpcpy(c_dest, c_src)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_str_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// StrHasPrefix is a wrapper around the C function g_str_has_prefix.
func StrHasPrefix(str string, prefix string) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_prefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(c_prefix))

	retC := C.g_str_has_prefix(c_str, c_prefix)
	retGo := retC == C.TRUE

	return retGo
}

// StrHasSuffix is a wrapper around the C function g_str_has_suffix.
func StrHasSuffix(str string, suffix string) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_suffix := C.CString(suffix)
	defer C.free(unsafe.Pointer(c_suffix))

	retC := C.g_str_has_suffix(c_str, c_suffix)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_str_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Strcanon is a wrapper around the C function g_strcanon.
func Strcanon(string_ string, validChars string, substitutor rune) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_valid_chars := C.CString(validChars)
	defer C.free(unsafe.Pointer(c_valid_chars))

	c_substitutor := (C.gchar)(substitutor)

	retC := C.g_strcanon(c_string, c_valid_chars, c_substitutor)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strcasecmp is a wrapper around the C function g_strcasecmp.
func Strcasecmp(s1 string, s2 string) int32 {
	c_s1 := C.CString(s1)
	defer C.free(unsafe.Pointer(c_s1))

	c_s2 := C.CString(s2)
	defer C.free(unsafe.Pointer(c_s2))

	retC := C.g_strcasecmp(c_s1, c_s2)
	retGo := (int32)(retC)

	return retGo
}

// Strchomp is a wrapper around the C function g_strchomp.
func Strchomp(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchomp(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strchug is a wrapper around the C function g_strchug.
func Strchug(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strchug(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strcmp0 is a wrapper around the C function g_strcmp0.
func Strcmp0(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_strcmp0(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

// Strcompress is a wrapper around the C function g_strcompress.
func Strcompress(source string) string {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	retC := C.g_strcompress(c_source)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Strdelimit is a wrapper around the C function g_strdelimit.
func Strdelimit(string_ string, delimiters string, newDelimiter rune) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiters := C.CString(delimiters)
	defer C.free(unsafe.Pointer(c_delimiters))

	c_new_delimiter := (C.gchar)(newDelimiter)

	retC := C.g_strdelimit(c_string, c_delimiters, c_new_delimiter)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strdown is a wrapper around the C function g_strdown.
func Strdown(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strdown(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strdup is a wrapper around the C function g_strdup.
func Strdup(str string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.g_strdup(c_str)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// StrdupPrintf is a wrapper around the C function g_strdup_printf.
func StrdupPrintf(format string, args ...interface{}) string {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_strdup_printf(c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strdup_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_strdupv : unsupported parameter str_array : in string with indirection level of 2

// Strerror is a wrapper around the C function g_strerror.
func Strerror(errnum int32) string {
	c_errnum := (C.gint)(errnum)

	retC := C.g_strerror(c_errnum)
	retGo := C.GoString(retC)

	return retGo
}

// Strescape is a wrapper around the C function g_strescape.
func Strescape(source string, exceptions string) string {
	c_source := C.CString(source)
	defer C.free(unsafe.Pointer(c_source))

	c_exceptions := C.CString(exceptions)
	defer C.free(unsafe.Pointer(c_exceptions))

	retC := C.g_strescape(c_source, c_exceptions)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strfreev : unsupported parameter str_array : in string with indirection level of 2

// StringNew is a wrapper around the C function g_string_new.
func StringNew(init string) *String {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	retC := C.g_string_new(c_init)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StringNewLen is a wrapper around the C function g_string_new_len.
func StringNewLen(init string, len int64) *String {
	c_init := C.CString(init)
	defer C.free(unsafe.Pointer(c_init))

	c_len := (C.gssize)(len)

	retC := C.g_string_new_len(c_init, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StringSizedNew is a wrapper around the C function g_string_sized_new.
func StringSizedNew(dflSize uint64) *String {
	c_dfl_size := (C.gsize)(dflSize)

	retC := C.g_string_sized_new(c_dfl_size)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StripContext is a wrapper around the C function g_strip_context.
func StripContext(msgid string, msgval string) string {
	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgval := C.CString(msgval)
	defer C.free(unsafe.Pointer(c_msgval))

	retC := C.g_strip_context(c_msgid, c_msgval)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Strlcat is a wrapper around the C function g_strlcat.
func Strlcat(dest string, src string, destSize uint64) uint64 {
	c_dest := C.CString(dest)
	defer C.free(unsafe.Pointer(c_dest))

	c_src := C.CString(src)
	defer C.free(unsafe.Pointer(c_src))

	c_dest_size := (C.gsize)(destSize)

	retC := C.g_strlcat(c_dest, c_src, c_dest_size)
	retGo := (uint64)(retC)

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
	retGo := (uint64)(retC)

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
	retGo := (int32)(retC)

	return retGo
}

// Strndup is a wrapper around the C function g_strndup.
func Strndup(str string, n uint64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_n := (C.gsize)(n)

	retC := C.g_strndup(c_str, c_n)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strnfill is a wrapper around the C function g_strnfill.
func Strnfill(length uint64, fillChar rune) string {
	c_length := (C.gsize)(length)

	c_fill_char := (C.gchar)(fillChar)

	retC := C.g_strnfill(c_length, c_fill_char)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strreverse is a wrapper around the C function g_strreverse.
func Strreverse(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strreverse(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Strrstr is a wrapper around the C function g_strrstr.
func Strrstr(haystack string, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strrstr(c_haystack, c_needle)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

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
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strsignal is a wrapper around the C function g_strsignal.
func Strsignal(signum int32) string {
	c_signum := (C.gint)(signum)

	retC := C.g_strsignal(c_signum)
	retGo := C.GoString(retC)

	return retGo
}

// Strsplit is a wrapper around the C function g_strsplit.
func Strsplit(string_ string, delimiter string, maxTokens int32) []string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiter := C.CString(delimiter)
	defer C.free(unsafe.Pointer(c_delimiter))

	c_max_tokens := (C.gint)(maxTokens)

	retC := C.g_strsplit(c_string, c_delimiter, c_max_tokens)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// StrsplitSet is a wrapper around the C function g_strsplit_set.
func StrsplitSet(string_ string, delimiters string, maxTokens int32) []string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiters := C.CString(delimiters)
	defer C.free(unsafe.Pointer(c_delimiters))

	c_max_tokens := (C.gint)(maxTokens)

	retC := C.g_strsplit_set(c_string, c_delimiters, c_max_tokens)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// StrstrLen is a wrapper around the C function g_strstr_len.
func StrstrLen(haystack string, haystackLen int64, needle string) string {
	c_haystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(c_haystack))

	c_haystack_len := (C.gssize)(haystackLen)

	c_needle := C.CString(needle)
	defer C.free(unsafe.Pointer(c_needle))

	retC := C.g_strstr_len(c_haystack, c_haystack_len, c_needle)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Strtod is a wrapper around the C function g_strtod.
func Strtod(nptr string) (float64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	retC := C.g_strtod(c_nptr, &c_endptr)
	retGo := (float64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// Strup is a wrapper around the C function g_strup.
func Strup(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_strup(c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_strv_get_type

// Unsupported : g_strv_length : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_test_add_data_func : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

// Unsupported : g_test_add_data_func_full : unsupported parameter test_data : no type generator for gpointer (gpointer) for param test_data

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc (GTestFunc) for param test_func

// Unsupported : g_test_add_vtable : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

// TestAssertExpectedMessagesInternal is a wrapper around the C function g_test_assert_expected_messages_internal.
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func_ string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	C.g_test_assert_expected_messages_internal(c_domain, c_file, c_line, c_func)

	return
}

// TestBug is a wrapper around the C function g_test_bug.
func TestBug(bugUriSnippet string) {
	c_bug_uri_snippet := C.CString(bugUriSnippet)
	defer C.free(unsafe.Pointer(c_bug_uri_snippet))

	C.g_test_bug(c_bug_uri_snippet)

	return
}

// TestBugBase is a wrapper around the C function g_test_bug_base.
func TestBugBase(uriPattern string) {
	c_uri_pattern := C.CString(uriPattern)
	defer C.free(unsafe.Pointer(c_uri_pattern))

	C.g_test_bug_base(c_uri_pattern)

	return
}

// Unsupported : g_test_create_case : unsupported parameter test_data : no type generator for gpointer (gconstpointer) for param test_data

// TestCreateSuite is a wrapper around the C function g_test_create_suite.
func TestCreateSuite(suiteName string) *TestSuite {
	c_suite_name := C.CString(suiteName)
	defer C.free(unsafe.Pointer(c_suite_name))

	retC := C.g_test_create_suite(c_suite_name)
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TestExpectMessage is a wrapper around the C function g_test_expect_message.
func TestExpectMessage(logDomain string, logLevel LogLevelFlags, pattern string) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_log_level := (C.GLogLevelFlags)(logLevel)

	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.g_test_expect_message(c_log_domain, c_log_level, c_pattern)

	return
}

// TestFail is a wrapper around the C function g_test_fail.
func TestFail() {
	C.g_test_fail()

	return
}

// TestGetRoot is a wrapper around the C function g_test_get_root.
func TestGetRoot() *TestSuite {
	retC := C.g_test_get_root()
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_test_init : unsupported parameter ... : varargs

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc (GTestLogFatalFunc) for param log_func

// TestLogTypeName is a wrapper around the C function g_test_log_type_name.
func TestLogTypeName(logType TestLogType) string {
	c_log_type := (C.GTestLogType)(logType)

	retC := C.g_test_log_type_name(c_log_type)
	retGo := C.GoString(retC)

	return retGo
}

// TestMaximizedResult is a wrapper around the C function g_test_maximized_result.
func TestMaximizedResult(maximizedQuantity float64, format string, args ...interface{}) {
	c_maximized_quantity := (C.double)(maximizedQuantity)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_test_maximized_result(c_maximized_quantity, c_format)

	return
}

// TestMessage is a wrapper around the C function g_test_message.
func TestMessage(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_test_message(c_format)

	return
}

// TestMinimizedResult is a wrapper around the C function g_test_minimized_result.
func TestMinimizedResult(minimizedQuantity float64, format string, args ...interface{}) {
	c_minimized_quantity := (C.double)(minimizedQuantity)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_test_minimized_result(c_minimized_quantity, c_format)

	return
}

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// Unsupported : g_test_queue_free : unsupported parameter gfree_pointer : no type generator for gpointer (gpointer) for param gfree_pointer

// TestRandDouble is a wrapper around the C function g_test_rand_double.
func TestRandDouble() float64 {
	retC := C.g_test_rand_double()
	retGo := (float64)(retC)

	return retGo
}

// TestRandDoubleRange is a wrapper around the C function g_test_rand_double_range.
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) float64 {
	c_range_start := (C.double)(rangeStart)

	c_range_end := (C.double)(rangeEnd)

	retC := C.g_test_rand_double_range(c_range_start, c_range_end)
	retGo := (float64)(retC)

	return retGo
}

// TestRandInt is a wrapper around the C function g_test_rand_int.
func TestRandInt() int32 {
	retC := C.g_test_rand_int()
	retGo := (int32)(retC)

	return retGo
}

// TestRandIntRange is a wrapper around the C function g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_test_rand_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// TestRun is a wrapper around the C function g_test_run.
func TestRun() int32 {
	retC := C.g_test_run()
	retGo := (int32)(retC)

	return retGo
}

// TestRunSuite is a wrapper around the C function g_test_run_suite.
func TestRunSuite(suite *TestSuite) int32 {
	c_suite := (*C.GTestSuite)(C.NULL)
	if suite != nil {
		c_suite = (*C.GTestSuite)(suite.ToC())
	}

	retC := C.g_test_run_suite(c_suite)
	retGo := (int32)(retC)

	return retGo
}

// TestTimerElapsed is a wrapper around the C function g_test_timer_elapsed.
func TestTimerElapsed() float64 {
	retC := C.g_test_timer_elapsed()
	retGo := (float64)(retC)

	return retGo
}

// TestTimerLast is a wrapper around the C function g_test_timer_last.
func TestTimerLast() float64 {
	retC := C.g_test_timer_last()
	retGo := (float64)(retC)

	return retGo
}

// TestTimerStart is a wrapper around the C function g_test_timer_start.
func TestTimerStart() {
	C.g_test_timer_start()

	return
}

// TestTrapAssertions is a wrapper around the C function g_test_trap_assertions.
func TestTrapAssertions(domain string, file string, line int32, func_ string, assertionFlags uint64, pattern string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_assertion_flags := (C.guint64)(assertionFlags)

	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.g_test_trap_assertions(c_domain, c_file, c_line, c_func, c_assertion_flags, c_pattern)

	return
}

// TestTrapFork is a wrapper around the C function g_test_trap_fork.
func TestTrapFork(usecTimeout uint64, testTrapFlags TestTrapFlags) bool {
	c_usec_timeout := (C.guint64)(usecTimeout)

	c_test_trap_flags := (C.GTestTrapFlags)(testTrapFlags)

	retC := C.g_test_trap_fork(c_usec_timeout, c_test_trap_flags)
	retGo := retC == C.TRUE

	return retGo
}

// TestTrapHasPassed is a wrapper around the C function g_test_trap_has_passed.
func TestTrapHasPassed() bool {
	retC := C.g_test_trap_has_passed()
	retGo := retC == C.TRUE

	return retGo
}

// TestTrapReachedTimeout is a wrapper around the C function g_test_trap_reached_timeout.
func TestTrapReachedTimeout() bool {
	retC := C.g_test_trap_reached_timeout()
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_thread_exit : unsupported parameter retval : no type generator for gpointer (gpointer) for param retval

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// TimeoutSourceNew is a wrapper around the C function g_timeout_source_new.
func TimeoutSourceNew(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TimeoutSourceNewSeconds is a wrapper around the C function g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new_seconds(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_trash_stack_peek : no return generator

// Unsupported : g_trash_stack_pop : no return generator

// Unsupported : g_trash_stack_push : unsupported parameter data_p : no type generator for gpointer (gpointer) for param data_p

// Unsupported : g_try_malloc : no return generator

// Unsupported : g_try_malloc0 : no return generator

// Unsupported : g_try_malloc0_n : no return generator

// Unsupported : g_try_malloc_n : no return generator

// Unsupported : g_try_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_try_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Blacklisted : g_ucs4_to_utf16

// Ucs4ToUtf8 is a wrapper around the C function g_ucs4_to_utf8.
func Ucs4ToUtf8(str rune, len int64) (string, int64, int64, error) {
	c_str := (C.gunichar)(str)

	c_len := (C.glong)(len)

	var c_items_read C.glong

	var c_items_written C.glong

	var cThrowableError *C.GError

	retC := C.g_ucs4_to_utf8(&c_str, c_len, &c_items_read, &c_items_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	itemsRead := (int64)(c_items_read)

	itemsWritten := (int64)(c_items_written)

	return retGo, itemsRead, itemsWritten, goError
}

// UnicharBreakType is a wrapper around the C function g_unichar_break_type.
func UnicharBreakType(c rune) UnicodeBreakType {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_break_type(c_c)
	retGo := (UnicodeBreakType)(retC)

	return retGo
}

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) int32 {
	c_uc := (C.gunichar)(uc)

	retC := C.g_unichar_combining_class(c_uc)
	retGo := (int32)(retC)

	return retGo
}

// UnicharCompose is a wrapper around the C function g_unichar_compose.
func UnicharCompose(a rune, b rune, ch rune) bool {
	c_a := (C.gunichar)(a)

	c_b := (C.gunichar)(b)

	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_compose(c_a, c_b, &c_ch)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharDecompose is a wrapper around the C function g_unichar_decompose.
func UnicharDecompose(ch rune, a rune, b rune) bool {
	c_ch := (C.gunichar)(ch)

	c_a := (C.gunichar)(a)

	c_b := (C.gunichar)(b)

	retC := C.g_unichar_decompose(c_ch, &c_a, &c_b)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharDigitValue is a wrapper around the C function g_unichar_digit_value.
func UnicharDigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_digit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// UnicharFullyDecompose is a wrapper around the C function g_unichar_fully_decompose.
func UnicharFullyDecompose(ch rune, compat bool, result rune, resultLen uint64) uint64 {
	c_ch := (C.gunichar)(ch)

	c_compat :=
		boolToGboolean(compat)

	c_result := (C.gunichar)(result)

	c_result_len := (C.gsize)(resultLen)

	retC := C.g_unichar_fully_decompose(c_ch, c_compat, &c_result, c_result_len)
	retGo := (uint64)(retC)

	return retGo
}

// UnicharGetMirrorChar is a wrapper around the C function g_unichar_get_mirror_char.
func UnicharGetMirrorChar(ch rune, mirroredCh rune) bool {
	c_ch := (C.gunichar)(ch)

	c_mirrored_ch := (C.gunichar)(mirroredCh)

	retC := C.g_unichar_get_mirror_char(c_ch, &c_mirrored_ch)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharGetScript is a wrapper around the C function g_unichar_get_script.
func UnicharGetScript(ch rune) UnicodeScript {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_get_script(c_ch)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// UnicharIsalnum is a wrapper around the C function g_unichar_isalnum.
func UnicharIsalnum(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isalnum(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsalpha is a wrapper around the C function g_unichar_isalpha.
func UnicharIsalpha(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isalpha(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIscntrl is a wrapper around the C function g_unichar_iscntrl.
func UnicharIscntrl(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iscntrl(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsdefined is a wrapper around the C function g_unichar_isdefined.
func UnicharIsdefined(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isdefined(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsdigit is a wrapper around the C function g_unichar_isdigit.
func UnicharIsdigit(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isdigit(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsgraph is a wrapper around the C function g_unichar_isgraph.
func UnicharIsgraph(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isgraph(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIslower is a wrapper around the C function g_unichar_islower.
func UnicharIslower(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_islower(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsmark is a wrapper around the C function g_unichar_ismark.
func UnicharIsmark(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ismark(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsprint is a wrapper around the C function g_unichar_isprint.
func UnicharIsprint(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isprint(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIspunct is a wrapper around the C function g_unichar_ispunct.
func UnicharIspunct(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ispunct(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsspace is a wrapper around the C function g_unichar_isspace.
func UnicharIsspace(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isspace(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIstitle is a wrapper around the C function g_unichar_istitle.
func UnicharIstitle(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_istitle(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsupper is a wrapper around the C function g_unichar_isupper.
func UnicharIsupper(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isupper(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIswide is a wrapper around the C function g_unichar_iswide.
func UnicharIswide(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iswide(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIswideCjk is a wrapper around the C function g_unichar_iswide_cjk.
func UnicharIswideCjk(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iswide_cjk(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIsxdigit is a wrapper around the C function g_unichar_isxdigit.
func UnicharIsxdigit(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isxdigit(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharIszerowidth is a wrapper around the C function g_unichar_iszerowidth.
func UnicharIszerowidth(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iszerowidth(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : g_unichar_to_utf8

// UnicharTolower is a wrapper around the C function g_unichar_tolower.
func UnicharTolower(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_tolower(c_c)
	retGo := (rune)(retC)

	return retGo
}

// UnicharTotitle is a wrapper around the C function g_unichar_totitle.
func UnicharTotitle(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_totitle(c_c)
	retGo := (rune)(retC)

	return retGo
}

// UnicharToupper is a wrapper around the C function g_unichar_toupper.
func UnicharToupper(c rune) rune {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_toupper(c_c)
	retGo := (rune)(retC)

	return retGo
}

// UnicharType is a wrapper around the C function g_unichar_type.
func UnicharType(c rune) UnicodeType {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_type(c_c)
	retGo := (UnicodeType)(retC)

	return retGo
}

// UnicharValidate is a wrapper around the C function g_unichar_validate.
func UnicharValidate(ch rune) bool {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_validate(c_ch)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharXdigitValue is a wrapper around the C function g_unichar_xdigit_value.
func UnicharXdigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_xdigit_value(c_c)
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_unicode_canonical_decomposition

// UnicodeCanonicalOrdering is a wrapper around the C function g_unicode_canonical_ordering.
func UnicodeCanonicalOrdering(string_ rune, len uint64) {
	c_string := (C.gunichar)(string_)

	c_len := (C.gsize)(len)

	C.g_unicode_canonical_ordering(&c_string, c_len)

	return
}

// UnicodeScriptFromIso15924 is a wrapper around the C function g_unicode_script_from_iso15924.
func UnicodeScriptFromIso15924(iso15924 uint32) UnicodeScript {
	c_iso15924 := (C.guint32)(iso15924)

	retC := C.g_unicode_script_from_iso15924(c_iso15924)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// UnicodeScriptToIso15924 is a wrapper around the C function g_unicode_script_to_iso15924.
func UnicodeScriptToIso15924(script UnicodeScript) uint32 {
	c_script := (C.GUnicodeScript)(script)

	retC := C.g_unicode_script_to_iso15924(c_script)
	retGo := (uint32)(retC)

	return retGo
}

// Blacklisted : g_unix_error_quark

// UnixOpenPipe is a wrapper around the C function g_unix_open_pipe.
func UnixOpenPipe(fds int32, flags int32) (bool, error) {
	c_fds := (C.gint)(fds)

	c_flags := (C.gint)(flags)

	var cThrowableError *C.GError

	retC := C.g_unix_open_pipe(&c_fds, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnixSetFdNonblocking is a wrapper around the C function g_unix_set_fd_nonblocking.
func UnixSetFdNonblocking(fd int32, nonblock bool) (bool, error) {
	c_fd := (C.gint)(fd)

	c_nonblock :=
		boolToGboolean(nonblock)

	var cThrowableError *C.GError

	retC := C.g_unix_set_fd_nonblocking(c_fd, c_nonblock, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// UnixSignalSourceNew is a wrapper around the C function g_unix_signal_source_new.
func UnixSignalSourceNew(signum int32) *Source {
	c_signum := (C.gint)(signum)

	retC := C.g_unix_signal_source_new(c_signum)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unlink is a wrapper around the C function g_unlink.
func Unlink(filename string) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_unlink(c_filename)
	retGo := (int32)(retC)

	return retGo
}

// Unsetenv is a wrapper around the C function g_unsetenv.
func Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_unsetenv(c_variable)

	return
}

// UriEscapeString is a wrapper around the C function g_uri_escape_string.
func UriEscapeString(unescaped string, reservedCharsAllowed string, allowUtf8 bool) string {
	c_unescaped := C.CString(unescaped)
	defer C.free(unsafe.Pointer(c_unescaped))

	c_reserved_chars_allowed := C.CString(reservedCharsAllowed)
	defer C.free(unsafe.Pointer(c_reserved_chars_allowed))

	c_allow_utf8 :=
		boolToGboolean(allowUtf8)

	retC := C.g_uri_escape_string(c_unescaped, c_reserved_chars_allowed, c_allow_utf8)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriListExtractUris is a wrapper around the C function g_uri_list_extract_uris.
func UriListExtractUris(uriList string) []string {
	c_uri_list := C.CString(uriList)
	defer C.free(unsafe.Pointer(c_uri_list))

	retC := C.g_uri_list_extract_uris(c_uri_list)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// UriParseScheme is a wrapper around the C function g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_uri_parse_scheme(c_uri)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriUnescapeSegment is a wrapper around the C function g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_escaped_string_end := C.CString(escapedStringEnd)
	defer C.free(unsafe.Pointer(c_escaped_string_end))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_segment(c_escaped_string, c_escaped_string_end, c_illegal_characters)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriUnescapeString is a wrapper around the C function g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_string(c_escaped_string, c_illegal_characters)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Usleep is a wrapper around the C function g_usleep.
func Usleep(microseconds uint64) {
	c_microseconds := (C.gulong)(microseconds)

	C.g_usleep(c_microseconds)

	return
}

// Blacklisted : g_utf16_to_ucs4

// Utf16ToUtf8 is a wrapper around the C function g_utf16_to_utf8.
func Utf16ToUtf8(str uint16, len int64) (string, int64, int64, error) {
	c_str := (C.gunichar2)(str)

	c_len := (C.glong)(len)

	var c_items_read C.glong

	var c_items_written C.glong

	var cThrowableError *C.GError

	retC := C.g_utf16_to_utf8(&c_str, c_len, &c_items_read, &c_items_written, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	itemsRead := (int64)(c_items_read)

	itemsWritten := (int64)(c_items_written)

	return retGo, itemsRead, itemsWritten, goError
}

// Utf8Casefold is a wrapper around the C function g_utf8_casefold.
func Utf8Casefold(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_casefold(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Collate is a wrapper around the C function g_utf8_collate.
func Utf8Collate(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_utf8_collate(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

// Utf8CollateKey is a wrapper around the C function g_utf8_collate_key.
func Utf8CollateKey(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_collate_key(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8CollateKeyForFilename is a wrapper around the C function g_utf8_collate_key_for_filename.
func Utf8CollateKeyForFilename(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_collate_key_for_filename(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8FindNextChar is a wrapper around the C function g_utf8_find_next_char.
func Utf8FindNextChar(p string, end string) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_end := C.CString(end)
	defer C.free(unsafe.Pointer(c_end))

	retC := C.g_utf8_find_next_char(c_p, c_end)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8FindPrevChar is a wrapper around the C function g_utf8_find_prev_char.
func Utf8FindPrevChar(str string, p string) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_find_prev_char(c_str, c_p)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8GetChar is a wrapper around the C function g_utf8_get_char.
func Utf8GetChar(p string) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_get_char(c_p)
	retGo := (rune)(retC)

	return retGo
}

// Utf8GetCharValidated is a wrapper around the C function g_utf8_get_char_validated.
func Utf8GetCharValidated(p string, maxLen int64) rune {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max_len := (C.gssize)(maxLen)

	retC := C.g_utf8_get_char_validated(c_p, c_max_len)
	retGo := (rune)(retC)

	return retGo
}

// Utf8Normalize is a wrapper around the C function g_utf8_normalize.
func Utf8Normalize(str string, len int64, mode NormalizeMode) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	c_mode := (C.GNormalizeMode)(mode)

	retC := C.g_utf8_normalize(c_str, c_len, c_mode)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8OffsetToPointer is a wrapper around the C function g_utf8_offset_to_pointer.
func Utf8OffsetToPointer(str string, offset int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_offset := (C.glong)(offset)

	retC := C.g_utf8_offset_to_pointer(c_str, c_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8PointerToOffset is a wrapper around the C function g_utf8_pointer_to_offset.
func Utf8PointerToOffset(str string, pos string) int64 {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_pos := C.CString(pos)
	defer C.free(unsafe.Pointer(c_pos))

	retC := C.g_utf8_pointer_to_offset(c_str, c_pos)
	retGo := (int64)(retC)

	return retGo
}

// Utf8PrevChar is a wrapper around the C function g_utf8_prev_char.
func Utf8PrevChar(p string) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	retC := C.g_utf8_prev_char(c_p)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strchr is a wrapper around the C function g_utf8_strchr.
func Utf8Strchr(p string, len int64, c rune) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	retC := C.g_utf8_strchr(c_p, c_len, c_c)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strdown is a wrapper around the C function g_utf8_strdown.
func Utf8Strdown(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strdown(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strlen is a wrapper around the C function g_utf8_strlen.
func Utf8Strlen(p string, max int64) int64 {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_max := (C.gssize)(max)

	retC := C.g_utf8_strlen(c_p, c_max)
	retGo := (int64)(retC)

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
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strrchr is a wrapper around the C function g_utf8_strrchr.
func Utf8Strrchr(p string, len int64, c rune) string {
	c_p := C.CString(p)
	defer C.free(unsafe.Pointer(c_p))

	c_len := (C.gssize)(len)

	c_c := (C.gunichar)(c)

	retC := C.g_utf8_strrchr(c_p, c_len, c_c)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strreverse is a wrapper around the C function g_utf8_strreverse.
func Utf8Strreverse(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strreverse(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Strup is a wrapper around the C function g_utf8_strup.
func Utf8Strup(str string, len int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.gssize)(len)

	retC := C.g_utf8_strup(c_str, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Utf8Substring is a wrapper around the C function g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_start_pos := (C.glong)(startPos)

	c_end_pos := (C.glong)(endPos)

	retC := C.g_utf8_substring(c_str, c_start_pos, c_end_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_utf8_to_ucs4

// Blacklisted : g_utf8_to_ucs4_fast

// Blacklisted : g_utf8_to_utf16

// Utf8Validate is a wrapper around the C function g_utf8_validate.
func Utf8Validate(str []uint8) (bool, string) {
	c_str_array := make([]C.guint8, len(str)+1, len(str)+1)
	for i, item := range str {
		c := (C.guint8)(item)
		c_str_array[i] = c
	}
	c_str_array[len(str)] = 0
	c_str_arrayPtr := &c_str_array[0]
	c_str := (*C.gchar)(unsafe.Pointer(c_str_arrayPtr))

	c_max_len := (C.gssize)(len(str))

	var c_end *C.gchar

	retC := C.g_utf8_validate(c_str, c_max_len, &c_end)
	retGo := retC == C.TRUE

	end := C.GoString(c_end)

	return retGo, end
}

// Blacklisted : g_variant_get_gtype

// Unsupported : g_variant_parse : unsupported parameter endptr : in string with indirection level of 2

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2

// Unsupported : g_vfprintf : unsupported parameter file : no type generator for gpointer (FILE*) for param file

// Unsupported : g_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_vsprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// WarnMessage is a wrapper around the C function g_warn_message.
func WarnMessage(domain string, file string, line int32, func_ string, warnexpr string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_line := (C.int)(line)

	c_func := C.CString(func_)
	defer C.free(unsafe.Pointer(c_func))

	c_warnexpr := C.CString(warnexpr)
	defer C.free(unsafe.Pointer(c_warnexpr))

	C.g_warn_message(c_domain, c_file, c_line, c_func, c_warnexpr)

	return
}

// Array is a wrapper around the C record GArray.
type Array struct {
	native *C.GArray
	Data   string
	Len    uint32
}

func ArrayNewFromC(u unsafe.Pointer) *Array {
	c := (*C.GArray)(u)
	if c == nil {
		return nil
	}

	g := &Array{
		Data:   C.GoString(c.data),
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
}

func (recv *Array) ToC() unsafe.Pointer {
	recv.native.data =
		C.CString(recv.Data)
	recv.native.len =
		(C.guint)(recv.Len)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Array with another Array, and returns true if they represent the same GObject.
func (recv *Array) Equals(other *Array) bool {
	return other.ToC() == recv.ToC()
}

// g_array_append_vals : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_free : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_get_element_size : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_insert_vals : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_new : no type generator for gpointer (gpointer) for array return
// g_array_prepend_vals : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_ref : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_remove_index : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_remove_index_fast : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_remove_range : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_set_clear_func : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_set_size : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_sized_new : no type generator for gpointer (gpointer) for array return
// g_array_sort : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_sort_with_data : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_unref : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// AsyncQueue is a wrapper around the C record GAsyncQueue.
type AsyncQueue struct {
	native *C.GAsyncQueue
}

func AsyncQueueNewFromC(u unsafe.Pointer) *AsyncQueue {
	c := (*C.GAsyncQueue)(u)
	if c == nil {
		return nil
	}

	g := &AsyncQueue{native: c}

	return g
}

func (recv *AsyncQueue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncQueue with another AsyncQueue, and returns true if they represent the same GObject.
func (recv *AsyncQueue) Equals(other *AsyncQueue) bool {
	return other.ToC() == recv.ToC()
}

// AsyncQueueNew is a wrapper around the C function g_async_queue_new.
func AsyncQueueNew() *AsyncQueue {
	retC := C.g_async_queue_new()
	retGo := AsyncQueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_async_queue_new_full : unsupported parameter item_free_func : no type generator for DestroyNotify (GDestroyNotify) for param item_free_func
// Length is a wrapper around the C function g_async_queue_length.
func (recv *AsyncQueue) Length() int32 {
	retC := C.g_async_queue_length((*C.GAsyncQueue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// LengthUnlocked is a wrapper around the C function g_async_queue_length_unlocked.
func (recv *AsyncQueue) LengthUnlocked() int32 {
	retC := C.g_async_queue_length_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Lock is a wrapper around the C function g_async_queue_lock.
func (recv *AsyncQueue) Lock() {
	C.g_async_queue_lock((*C.GAsyncQueue)(recv.native))

	return
}

// Unsupported : g_async_queue_pop : no return generator

// Unsupported : g_async_queue_pop_unlocked : no return generator

// Unsupported : g_async_queue_push : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_async_queue_push_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_async_queue_push_sorted_unlocked : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_async_queue_push_unlocked : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Ref is a wrapper around the C function g_async_queue_ref.
func (recv *AsyncQueue) Ref() *AsyncQueue {
	retC := C.g_async_queue_ref((*C.GAsyncQueue)(recv.native))
	retGo := AsyncQueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefUnlocked is a wrapper around the C function g_async_queue_ref_unlocked.
func (recv *AsyncQueue) RefUnlocked() {
	C.g_async_queue_ref_unlocked((*C.GAsyncQueue)(recv.native))

	return
}

// Unsupported : g_async_queue_sort : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_async_queue_sort_unlocked : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_async_queue_timed_pop : no return generator

// Unsupported : g_async_queue_timed_pop_unlocked : no return generator

// Unsupported : g_async_queue_timeout_pop : no return generator

// Unsupported : g_async_queue_timeout_pop_unlocked : no return generator

// Unsupported : g_async_queue_try_pop : no return generator

// Unsupported : g_async_queue_try_pop_unlocked : no return generator

// Unlock is a wrapper around the C function g_async_queue_unlock.
func (recv *AsyncQueue) Unlock() {
	C.g_async_queue_unlock((*C.GAsyncQueue)(recv.native))

	return
}

// Unref is a wrapper around the C function g_async_queue_unref.
func (recv *AsyncQueue) Unref() {
	C.g_async_queue_unref((*C.GAsyncQueue)(recv.native))

	return
}

// UnrefAndUnlock is a wrapper around the C function g_async_queue_unref_and_unlock.
func (recv *AsyncQueue) UnrefAndUnlock() {
	C.g_async_queue_unref_and_unlock((*C.GAsyncQueue)(recv.native))

	return
}

// BookmarkFile is a wrapper around the C record GBookmarkFile.
type BookmarkFile struct {
	native *C.GBookmarkFile
}

func BookmarkFileNewFromC(u unsafe.Pointer) *BookmarkFile {
	c := (*C.GBookmarkFile)(u)
	if c == nil {
		return nil
	}

	g := &BookmarkFile{native: c}

	return g
}

func (recv *BookmarkFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BookmarkFile with another BookmarkFile, and returns true if they represent the same GObject.
func (recv *BookmarkFile) Equals(other *BookmarkFile) bool {
	return other.ToC() == recv.ToC()
}

// BookmarkFileErrorQuark is a wrapper around the C function g_bookmark_file_error_quark.
func BookmarkFileErrorQuark() Quark {
	retC := C.g_bookmark_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// BookmarkFileNew is a wrapper around the C function g_bookmark_file_new.
func BookmarkFileNew() *BookmarkFile {
	retC := C.g_bookmark_file_new()
	retGo := BookmarkFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddApplication is a wrapper around the C function g_bookmark_file_add_application.
func (recv *BookmarkFile) AddApplication(uri string, name string, exec string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_exec := C.CString(exec)
	defer C.free(unsafe.Pointer(c_exec))

	C.g_bookmark_file_add_application((*C.GBookmarkFile)(recv.native), c_uri, c_name, c_exec)

	return
}

// AddGroup is a wrapper around the C function g_bookmark_file_add_group.
func (recv *BookmarkFile) AddGroup(uri string, group string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	C.g_bookmark_file_add_group((*C.GBookmarkFile)(recv.native), c_uri, c_group)

	return
}

// Free is a wrapper around the C function g_bookmark_file_free.
func (recv *BookmarkFile) Free() {
	C.g_bookmark_file_free((*C.GBookmarkFile)(recv.native))

	return
}

// GetAdded is a wrapper around the C function g_bookmark_file_get_added.
func (recv *BookmarkFile) GetAdded(uri string) (int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_added((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAppInfo is a wrapper around the C function g_bookmark_file_get_app_info.
func (recv *BookmarkFile) GetAppInfo(uri string, name string) (bool, string, uint32, int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_exec *C.gchar

	var c_count C.guint

	var c_stamp C.time_t

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_app_info((*C.GBookmarkFile)(recv.native), c_uri, c_name, &c_exec, &c_count, &c_stamp, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	exec := C.GoString(c_exec)
	defer C.free(unsafe.Pointer(c_exec))

	count := (uint32)(c_count)

	stamp := (int64)(c_stamp)

	return retGo, exec, count, stamp, goError
}

// GetApplications is a wrapper around the C function g_bookmark_file_get_applications.
func (recv *BookmarkFile) GetApplications(uri string) ([]string, uint64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_applications((*C.GBookmarkFile)(recv.native), c_uri, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetDescription is a wrapper around the C function g_bookmark_file_get_description.
func (recv *BookmarkFile) GetDescription(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_description((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetGroups is a wrapper around the C function g_bookmark_file_get_groups.
func (recv *BookmarkFile) GetGroups(uri string) ([]string, uint64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_groups((*C.GBookmarkFile)(recv.native), c_uri, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetIcon is a wrapper around the C function g_bookmark_file_get_icon.
func (recv *BookmarkFile) GetIcon(uri string) (bool, string, string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_href *C.gchar

	var c_mime_type *C.gchar

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_icon((*C.GBookmarkFile)(recv.native), c_uri, &c_href, &c_mime_type, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	href := C.GoString(c_href)
	defer C.free(unsafe.Pointer(c_href))

	mimeType := C.GoString(c_mime_type)
	defer C.free(unsafe.Pointer(c_mime_type))

	return retGo, href, mimeType, goError
}

// GetIsPrivate is a wrapper around the C function g_bookmark_file_get_is_private.
func (recv *BookmarkFile) GetIsPrivate(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_is_private((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetMimeType is a wrapper around the C function g_bookmark_file_get_mime_type.
func (recv *BookmarkFile) GetMimeType(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_mime_type((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetModified is a wrapper around the C function g_bookmark_file_get_modified.
func (recv *BookmarkFile) GetModified(uri string) (int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_modified((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetSize is a wrapper around the C function g_bookmark_file_get_size.
func (recv *BookmarkFile) GetSize() int32 {
	retC := C.g_bookmark_file_get_size((*C.GBookmarkFile)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTitle is a wrapper around the C function g_bookmark_file_get_title.
func (recv *BookmarkFile) GetTitle(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_title((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetUris is a wrapper around the C function g_bookmark_file_get_uris.
func (recv *BookmarkFile) GetUris() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_bookmark_file_get_uris((*C.GBookmarkFile)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetVisited is a wrapper around the C function g_bookmark_file_get_visited.
func (recv *BookmarkFile) GetVisited(uri string) (int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_visited((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasApplication is a wrapper around the C function g_bookmark_file_has_application.
func (recv *BookmarkFile) HasApplication(uri string, name string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_has_application((*C.GBookmarkFile)(recv.native), c_uri, c_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasGroup is a wrapper around the C function g_bookmark_file_has_group.
func (recv *BookmarkFile) HasGroup(uri string, group string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_has_group((*C.GBookmarkFile)(recv.native), c_uri, c_group, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasItem is a wrapper around the C function g_bookmark_file_has_item.
func (recv *BookmarkFile) HasItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_bookmark_file_has_item((*C.GBookmarkFile)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// LoadFromData is a wrapper around the C function g_bookmark_file_load_from_data.
func (recv *BookmarkFile) LoadFromData(data []uint8) (bool, error) {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.gchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_data((*C.GBookmarkFile)(recv.native), c_data, c_length, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadFromDataDirs is a wrapper around the C function g_bookmark_file_load_from_data_dirs.
func (recv *BookmarkFile) LoadFromDataDirs(file string) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var c_full_path *C.gchar

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_data_dirs((*C.GBookmarkFile)(recv.native), c_file, &c_full_path, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goError
}

// LoadFromFile is a wrapper around the C function g_bookmark_file_load_from_file.
func (recv *BookmarkFile) LoadFromFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_file((*C.GBookmarkFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MoveItem is a wrapper around the C function g_bookmark_file_move_item.
func (recv *BookmarkFile) MoveItem(oldUri string, newUri string) (bool, error) {
	c_old_uri := C.CString(oldUri)
	defer C.free(unsafe.Pointer(c_old_uri))

	c_new_uri := C.CString(newUri)
	defer C.free(unsafe.Pointer(c_new_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_move_item((*C.GBookmarkFile)(recv.native), c_old_uri, c_new_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveApplication is a wrapper around the C function g_bookmark_file_remove_application.
func (recv *BookmarkFile) RemoveApplication(uri string, name string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_remove_application((*C.GBookmarkFile)(recv.native), c_uri, c_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveGroup is a wrapper around the C function g_bookmark_file_remove_group.
func (recv *BookmarkFile) RemoveGroup(uri string, group string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_remove_group((*C.GBookmarkFile)(recv.native), c_uri, c_group, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveItem is a wrapper around the C function g_bookmark_file_remove_item.
func (recv *BookmarkFile) RemoveItem(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_remove_item((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAdded is a wrapper around the C function g_bookmark_file_set_added.
func (recv *BookmarkFile) SetAdded(uri string, added int64) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_added := (C.time_t)(added)

	C.g_bookmark_file_set_added((*C.GBookmarkFile)(recv.native), c_uri, c_added)

	return
}

// SetAppInfo is a wrapper around the C function g_bookmark_file_set_app_info.
func (recv *BookmarkFile) SetAppInfo(uri string, name string, exec string, count int32, stamp int64) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_exec := C.CString(exec)
	defer C.free(unsafe.Pointer(c_exec))

	c_count := (C.gint)(count)

	c_stamp := (C.time_t)(stamp)

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_set_app_info((*C.GBookmarkFile)(recv.native), c_uri, c_name, c_exec, c_count, c_stamp, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetDescription is a wrapper around the C function g_bookmark_file_set_description.
func (recv *BookmarkFile) SetDescription(uri string, description string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_bookmark_file_set_description((*C.GBookmarkFile)(recv.native), c_uri, c_description)

	return
}

// SetGroups is a wrapper around the C function g_bookmark_file_set_groups.
func (recv *BookmarkFile) SetGroups(uri string, groups []string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_groups_array := make([]*C.gchar, len(groups)+1, len(groups)+1)
	for i, item := range groups {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_groups_array[i] = c
	}
	c_groups_array[len(groups)] = nil
	c_groups_arrayPtr := &c_groups_array[0]
	c_groups := (**C.gchar)(unsafe.Pointer(c_groups_arrayPtr))

	c_length := (C.gsize)(len(groups))

	C.g_bookmark_file_set_groups((*C.GBookmarkFile)(recv.native), c_uri, c_groups, c_length)

	return
}

// SetIcon is a wrapper around the C function g_bookmark_file_set_icon.
func (recv *BookmarkFile) SetIcon(uri string, href string, mimeType string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_href := C.CString(href)
	defer C.free(unsafe.Pointer(c_href))

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.g_bookmark_file_set_icon((*C.GBookmarkFile)(recv.native), c_uri, c_href, c_mime_type)

	return
}

// SetIsPrivate is a wrapper around the C function g_bookmark_file_set_is_private.
func (recv *BookmarkFile) SetIsPrivate(uri string, isPrivate bool) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_is_private :=
		boolToGboolean(isPrivate)

	C.g_bookmark_file_set_is_private((*C.GBookmarkFile)(recv.native), c_uri, c_is_private)

	return
}

// SetMimeType is a wrapper around the C function g_bookmark_file_set_mime_type.
func (recv *BookmarkFile) SetMimeType(uri string, mimeType string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.g_bookmark_file_set_mime_type((*C.GBookmarkFile)(recv.native), c_uri, c_mime_type)

	return
}

// SetModified is a wrapper around the C function g_bookmark_file_set_modified.
func (recv *BookmarkFile) SetModified(uri string, modified int64) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_modified := (C.time_t)(modified)

	C.g_bookmark_file_set_modified((*C.GBookmarkFile)(recv.native), c_uri, c_modified)

	return
}

// SetTitle is a wrapper around the C function g_bookmark_file_set_title.
func (recv *BookmarkFile) SetTitle(uri string, title string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_bookmark_file_set_title((*C.GBookmarkFile)(recv.native), c_uri, c_title)

	return
}

// SetVisited is a wrapper around the C function g_bookmark_file_set_visited.
func (recv *BookmarkFile) SetVisited(uri string, visited int64) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_visited := (C.time_t)(visited)

	C.g_bookmark_file_set_visited((*C.GBookmarkFile)(recv.native), c_uri, c_visited)

	return
}

// Unsupported : g_bookmark_file_to_data : array return type :

// ToFile is a wrapper around the C function g_bookmark_file_to_file.
func (recv *BookmarkFile) ToFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_to_file((*C.GBookmarkFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Blacklisted : GByteArray

// Bytes is a wrapper around the C record GBytes.
type Bytes struct {
	native *C.GBytes
}

func BytesNewFromC(u unsafe.Pointer) *Bytes {
	c := (*C.GBytes)(u)
	if c == nil {
		return nil
	}

	g := &Bytes{native: c}

	return g
}

func (recv *Bytes) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Bytes with another Bytes, and returns true if they represent the same GObject.
func (recv *Bytes) Equals(other *Bytes) bool {
	return other.ToC() == recv.ToC()
}

// BytesNew is a wrapper around the C function g_bytes_new.
func BytesNew(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BytesNewStatic is a wrapper around the C function g_bytes_new_static.
func BytesNewStatic(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_static(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BytesNewTake is a wrapper around the C function g_bytes_new_take.
func BytesNewTake(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_take(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_bytes_new_with_free_func : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// Compare is a wrapper around the C function g_bytes_compare.
func (recv *Bytes) Compare(bytes2 *Bytes) int32 {
	c_bytes2 := (C.gconstpointer)(C.NULL)
	if bytes2 != nil {
		c_bytes2 = (C.gconstpointer)(bytes2.ToC())
	}

	retC := C.g_bytes_compare((C.gconstpointer)(recv.native), c_bytes2)
	retGo := (int32)(retC)

	return retGo
}

// Equal is a wrapper around the C function g_bytes_equal.
func (recv *Bytes) Equal(bytes2 *Bytes) bool {
	c_bytes2 := (C.gconstpointer)(C.NULL)
	if bytes2 != nil {
		c_bytes2 = (C.gconstpointer)(bytes2.ToC())
	}

	retC := C.g_bytes_equal((C.gconstpointer)(recv.native), c_bytes2)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_bytes_get_data : array return type :

// GetSize is a wrapper around the C function g_bytes_get_size.
func (recv *Bytes) GetSize() uint64 {
	retC := C.g_bytes_get_size((*C.GBytes)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Hash is a wrapper around the C function g_bytes_hash.
func (recv *Bytes) Hash() uint32 {
	retC := C.g_bytes_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NewFromBytes is a wrapper around the C function g_bytes_new_from_bytes.
func (recv *Bytes) NewFromBytes(offset uint64, length uint64) *Bytes {
	c_offset := (C.gsize)(offset)

	c_length := (C.gsize)(length)

	retC := C.g_bytes_new_from_bytes((*C.GBytes)(recv.native), c_offset, c_length)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_bytes_ref.
func (recv *Bytes) Ref() *Bytes {
	retC := C.g_bytes_ref((*C.GBytes)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_bytes_unref.
func (recv *Bytes) Unref() {
	C.g_bytes_unref((*C.GBytes)(recv.native))

	return
}

// Unsupported : g_bytes_unref_to_array : array return type :

// Unsupported : g_bytes_unref_to_data : array return type :

// Checksum is a wrapper around the C record GChecksum.
type Checksum struct {
	native *C.GChecksum
}

func ChecksumNewFromC(u unsafe.Pointer) *Checksum {
	c := (*C.GChecksum)(u)
	if c == nil {
		return nil
	}

	g := &Checksum{native: c}

	return g
}

func (recv *Checksum) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Checksum with another Checksum, and returns true if they represent the same GObject.
func (recv *Checksum) Equals(other *Checksum) bool {
	return other.ToC() == recv.ToC()
}

// ChecksumNew is a wrapper around the C function g_checksum_new.
func ChecksumNew(checksumType ChecksumType) *Checksum {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_new(c_checksum_type)
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ChecksumTypeGetLength is a wrapper around the C function g_checksum_type_get_length.
func ChecksumTypeGetLength(checksumType ChecksumType) int64 {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_type_get_length(c_checksum_type)
	retGo := (int64)(retC)

	return retGo
}

// Copy is a wrapper around the C function g_checksum_copy.
func (recv *Checksum) Copy() *Checksum {
	retC := C.g_checksum_copy((*C.GChecksum)(recv.native))
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_checksum_free.
func (recv *Checksum) Free() {
	C.g_checksum_free((*C.GChecksum)(recv.native))

	return
}

// GetDigest is a wrapper around the C function g_checksum_get_digest.
func (recv *Checksum) GetDigest(buffer uint8, digestLen uint64) {
	c_buffer := (C.guint8)(buffer)

	c_digest_len := (C.gsize)(digestLen)

	C.g_checksum_get_digest((*C.GChecksum)(recv.native), &c_buffer, &c_digest_len)

	return
}

// GetString is a wrapper around the C function g_checksum_get_string.
func (recv *Checksum) GetString() string {
	retC := C.g_checksum_get_string((*C.GChecksum)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Reset is a wrapper around the C function g_checksum_reset.
func (recv *Checksum) Reset() {
	C.g_checksum_reset((*C.GChecksum)(recv.native))

	return
}

// Update is a wrapper around the C function g_checksum_update.
func (recv *Checksum) Update(data []uint8) {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gssize)(len(data))

	C.g_checksum_update((*C.GChecksum)(recv.native), c_data, c_length)

	return
}

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native *C.GCond
	// Private : p
	// Private : i
}

func CondNewFromC(u unsafe.Pointer) *Cond {
	c := (*C.GCond)(u)
	if c == nil {
		return nil
	}

	g := &Cond{native: c}

	return g
}

func (recv *Cond) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Cond with another Cond, and returns true if they represent the same GObject.
func (recv *Cond) Equals(other *Cond) bool {
	return other.ToC() == recv.ToC()
}

// Broadcast is a wrapper around the C function g_cond_broadcast.
func (recv *Cond) Broadcast() {
	C.g_cond_broadcast((*C.GCond)(recv.native))

	return
}

// Clear is a wrapper around the C function g_cond_clear.
func (recv *Cond) Clear() {
	C.g_cond_clear((*C.GCond)(recv.native))

	return
}

// Init is a wrapper around the C function g_cond_init.
func (recv *Cond) Init() {
	C.g_cond_init((*C.GCond)(recv.native))

	return
}

// Signal is a wrapper around the C function g_cond_signal.
func (recv *Cond) Signal() {
	C.g_cond_signal((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// Data is a wrapper around the C record GData.
type Data struct {
	native *C.GData
}

func DataNewFromC(u unsafe.Pointer) *Data {
	c := (*C.GData)(u)
	if c == nil {
		return nil
	}

	g := &Data{native: c}

	return g
}

func (recv *Data) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Data with another Data, and returns true if they represent the same GObject.
func (recv *Data) Equals(other *Data) bool {
	return other.ToC() == recv.ToC()
}

// Date is a wrapper around the C record GDate.
type Date struct {
	native *C.GDate
	// Bitfield not supported : 32 julian_days
	// Bitfield not supported :  1 julian
	// Bitfield not supported :  1 dmy
	// Bitfield not supported :  6 day
	// Bitfield not supported :  4 month
	// Bitfield not supported : 16 year
}

func DateNewFromC(u unsafe.Pointer) *Date {
	c := (*C.GDate)(u)
	if c == nil {
		return nil
	}

	g := &Date{native: c}

	return g
}

func (recv *Date) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Date with another Date, and returns true if they represent the same GObject.
func (recv *Date) Equals(other *Date) bool {
	return other.ToC() == recv.ToC()
}

// DateNew is a wrapper around the C function g_date_new.
func DateNew() *Date {
	retC := C.g_date_new()
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateNewDmy is a wrapper around the C function g_date_new_dmy.
func DateNewDmy(day DateDay, month DateMonth, year DateYear) *Date {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_new_dmy(c_day, c_month, c_year)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateNewJulian is a wrapper around the C function g_date_new_julian.
func DateNewJulian(julianDay uint32) *Date {
	c_julian_day := (C.guint32)(julianDay)

	retC := C.g_date_new_julian(c_julian_day)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateGetDaysInMonth is a wrapper around the C function g_date_get_days_in_month.
func DateGetDaysInMonth(month DateMonth, year DateYear) uint8 {
	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_days_in_month(c_month, c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateGetMondayWeeksInYear is a wrapper around the C function g_date_get_monday_weeks_in_year.
func DateGetMondayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_monday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateGetSundayWeeksInYear is a wrapper around the C function g_date_get_sunday_weeks_in_year.
func DateGetSundayWeeksInYear(year DateYear) uint8 {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_get_sunday_weeks_in_year(c_year)
	retGo := (uint8)(retC)

	return retGo
}

// DateIsLeapYear is a wrapper around the C function g_date_is_leap_year.
func DateIsLeapYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_is_leap_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// DateStrftime is a wrapper around the C function g_date_strftime.
func DateStrftime(s string, slen uint64, format string, date *Date) uint64 {
	c_s := C.CString(s)
	defer C.free(unsafe.Pointer(c_s))

	c_slen := (C.gsize)(slen)

	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	c_date := (*C.GDate)(C.NULL)
	if date != nil {
		c_date = (*C.GDate)(date.ToC())
	}

	retC := C.g_date_strftime(c_s, c_slen, c_format, c_date)
	retGo := (uint64)(retC)

	return retGo
}

// DateValidDay is a wrapper around the C function g_date_valid_day.
func DateValidDay(day DateDay) bool {
	c_day := (C.GDateDay)(day)

	retC := C.g_date_valid_day(c_day)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidDmy is a wrapper around the C function g_date_valid_dmy.
func DateValidDmy(day DateDay, month DateMonth, year DateYear) bool {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_dmy(c_day, c_month, c_year)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidJulian is a wrapper around the C function g_date_valid_julian.
func DateValidJulian(julianDate uint32) bool {
	c_julian_date := (C.guint32)(julianDate)

	retC := C.g_date_valid_julian(c_julian_date)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidMonth is a wrapper around the C function g_date_valid_month.
func DateValidMonth(month DateMonth) bool {
	c_month := (C.GDateMonth)(month)

	retC := C.g_date_valid_month(c_month)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidWeekday is a wrapper around the C function g_date_valid_weekday.
func DateValidWeekday(weekday DateWeekday) bool {
	c_weekday := (C.GDateWeekday)(weekday)

	retC := C.g_date_valid_weekday(c_weekday)
	retGo := retC == C.TRUE

	return retGo
}

// DateValidYear is a wrapper around the C function g_date_valid_year.
func DateValidYear(year DateYear) bool {
	c_year := (C.GDateYear)(year)

	retC := C.g_date_valid_year(c_year)
	retGo := retC == C.TRUE

	return retGo
}

// AddDays is a wrapper around the C function g_date_add_days.
func (recv *Date) AddDays(nDays uint32) {
	c_n_days := (C.guint)(nDays)

	C.g_date_add_days((*C.GDate)(recv.native), c_n_days)

	return
}

// AddMonths is a wrapper around the C function g_date_add_months.
func (recv *Date) AddMonths(nMonths uint32) {
	c_n_months := (C.guint)(nMonths)

	C.g_date_add_months((*C.GDate)(recv.native), c_n_months)

	return
}

// AddYears is a wrapper around the C function g_date_add_years.
func (recv *Date) AddYears(nYears uint32) {
	c_n_years := (C.guint)(nYears)

	C.g_date_add_years((*C.GDate)(recv.native), c_n_years)

	return
}

// Clamp is a wrapper around the C function g_date_clamp.
func (recv *Date) Clamp(minDate *Date, maxDate *Date) {
	c_min_date := (*C.GDate)(C.NULL)
	if minDate != nil {
		c_min_date = (*C.GDate)(minDate.ToC())
	}

	c_max_date := (*C.GDate)(C.NULL)
	if maxDate != nil {
		c_max_date = (*C.GDate)(maxDate.ToC())
	}

	C.g_date_clamp((*C.GDate)(recv.native), c_min_date, c_max_date)

	return
}

// Clear is a wrapper around the C function g_date_clear.
func (recv *Date) Clear(nDates uint32) {
	c_n_dates := (C.guint)(nDates)

	C.g_date_clear((*C.GDate)(recv.native), c_n_dates)

	return
}

// Compare is a wrapper around the C function g_date_compare.
func (recv *Date) Compare(rhs *Date) int32 {
	c_rhs := (*C.GDate)(C.NULL)
	if rhs != nil {
		c_rhs = (*C.GDate)(rhs.ToC())
	}

	retC := C.g_date_compare((*C.GDate)(recv.native), c_rhs)
	retGo := (int32)(retC)

	return retGo
}

// DaysBetween is a wrapper around the C function g_date_days_between.
func (recv *Date) DaysBetween(date2 *Date) int32 {
	c_date2 := (*C.GDate)(C.NULL)
	if date2 != nil {
		c_date2 = (*C.GDate)(date2.ToC())
	}

	retC := C.g_date_days_between((*C.GDate)(recv.native), c_date2)
	retGo := (int32)(retC)

	return retGo
}

// Free is a wrapper around the C function g_date_free.
func (recv *Date) Free() {
	C.g_date_free((*C.GDate)(recv.native))

	return
}

// GetDay is a wrapper around the C function g_date_get_day.
func (recv *Date) GetDay() DateDay {
	retC := C.g_date_get_day((*C.GDate)(recv.native))
	retGo := (DateDay)(retC)

	return retGo
}

// GetDayOfYear is a wrapper around the C function g_date_get_day_of_year.
func (recv *Date) GetDayOfYear() uint32 {
	retC := C.g_date_get_day_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetIso8601WeekOfYear is a wrapper around the C function g_date_get_iso8601_week_of_year.
func (recv *Date) GetIso8601WeekOfYear() uint32 {
	retC := C.g_date_get_iso8601_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetJulian is a wrapper around the C function g_date_get_julian.
func (recv *Date) GetJulian() uint32 {
	retC := C.g_date_get_julian((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMondayWeekOfYear is a wrapper around the C function g_date_get_monday_week_of_year.
func (recv *Date) GetMondayWeekOfYear() uint32 {
	retC := C.g_date_get_monday_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMonth is a wrapper around the C function g_date_get_month.
func (recv *Date) GetMonth() DateMonth {
	retC := C.g_date_get_month((*C.GDate)(recv.native))
	retGo := (DateMonth)(retC)

	return retGo
}

// GetSundayWeekOfYear is a wrapper around the C function g_date_get_sunday_week_of_year.
func (recv *Date) GetSundayWeekOfYear() uint32 {
	retC := C.g_date_get_sunday_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetWeekday is a wrapper around the C function g_date_get_weekday.
func (recv *Date) GetWeekday() DateWeekday {
	retC := C.g_date_get_weekday((*C.GDate)(recv.native))
	retGo := (DateWeekday)(retC)

	return retGo
}

// GetYear is a wrapper around the C function g_date_get_year.
func (recv *Date) GetYear() DateYear {
	retC := C.g_date_get_year((*C.GDate)(recv.native))
	retGo := (DateYear)(retC)

	return retGo
}

// IsFirstOfMonth is a wrapper around the C function g_date_is_first_of_month.
func (recv *Date) IsFirstOfMonth() bool {
	retC := C.g_date_is_first_of_month((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsLastOfMonth is a wrapper around the C function g_date_is_last_of_month.
func (recv *Date) IsLastOfMonth() bool {
	retC := C.g_date_is_last_of_month((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Order is a wrapper around the C function g_date_order.
func (recv *Date) Order(date2 *Date) {
	c_date2 := (*C.GDate)(C.NULL)
	if date2 != nil {
		c_date2 = (*C.GDate)(date2.ToC())
	}

	C.g_date_order((*C.GDate)(recv.native), c_date2)

	return
}

// SetDay is a wrapper around the C function g_date_set_day.
func (recv *Date) SetDay(day DateDay) {
	c_day := (C.GDateDay)(day)

	C.g_date_set_day((*C.GDate)(recv.native), c_day)

	return
}

// SetDmy is a wrapper around the C function g_date_set_dmy.
func (recv *Date) SetDmy(day DateDay, month DateMonth, y DateYear) {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_y := (C.GDateYear)(y)

	C.g_date_set_dmy((*C.GDate)(recv.native), c_day, c_month, c_y)

	return
}

// SetJulian is a wrapper around the C function g_date_set_julian.
func (recv *Date) SetJulian(julianDate uint32) {
	c_julian_date := (C.guint32)(julianDate)

	C.g_date_set_julian((*C.GDate)(recv.native), c_julian_date)

	return
}

// SetMonth is a wrapper around the C function g_date_set_month.
func (recv *Date) SetMonth(month DateMonth) {
	c_month := (C.GDateMonth)(month)

	C.g_date_set_month((*C.GDate)(recv.native), c_month)

	return
}

// SetParse is a wrapper around the C function g_date_set_parse.
func (recv *Date) SetParse(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.g_date_set_parse((*C.GDate)(recv.native), c_str)

	return
}

// SetTime is a wrapper around the C function g_date_set_time.
func (recv *Date) SetTime(time Time) {
	c_time_ := (C.GTime)(time)

	C.g_date_set_time((*C.GDate)(recv.native), c_time_)

	return
}

// SetTimeT is a wrapper around the C function g_date_set_time_t.
func (recv *Date) SetTimeT(timet int64) {
	c_timet := (C.time_t)(timet)

	C.g_date_set_time_t((*C.GDate)(recv.native), c_timet)

	return
}

// SetTimeVal is a wrapper around the C function g_date_set_time_val.
func (recv *Date) SetTimeVal(timeval *TimeVal) {
	c_timeval := (*C.GTimeVal)(C.NULL)
	if timeval != nil {
		c_timeval = (*C.GTimeVal)(timeval.ToC())
	}

	C.g_date_set_time_val((*C.GDate)(recv.native), c_timeval)

	return
}

// SetYear is a wrapper around the C function g_date_set_year.
func (recv *Date) SetYear(year DateYear) {
	c_year := (C.GDateYear)(year)

	C.g_date_set_year((*C.GDate)(recv.native), c_year)

	return
}

// SubtractDays is a wrapper around the C function g_date_subtract_days.
func (recv *Date) SubtractDays(nDays uint32) {
	c_n_days := (C.guint)(nDays)

	C.g_date_subtract_days((*C.GDate)(recv.native), c_n_days)

	return
}

// SubtractMonths is a wrapper around the C function g_date_subtract_months.
func (recv *Date) SubtractMonths(nMonths uint32) {
	c_n_months := (C.guint)(nMonths)

	C.g_date_subtract_months((*C.GDate)(recv.native), c_n_months)

	return
}

// SubtractYears is a wrapper around the C function g_date_subtract_years.
func (recv *Date) SubtractYears(nYears uint32) {
	c_n_years := (C.guint)(nYears)

	C.g_date_subtract_years((*C.GDate)(recv.native), c_n_years)

	return
}

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer (tm*) for param tm

// Valid is a wrapper around the C function g_date_valid.
func (recv *Date) Valid() bool {
	retC := C.g_date_valid((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// DateTime is a wrapper around the C record GDateTime.
type DateTime struct {
	native *C.GDateTime
}

func DateTimeNewFromC(u unsafe.Pointer) *DateTime {
	c := (*C.GDateTime)(u)
	if c == nil {
		return nil
	}

	g := &DateTime{native: c}

	return g
}

func (recv *DateTime) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DateTime with another DateTime, and returns true if they represent the same GObject.
func (recv *DateTime) Equals(other *DateTime) bool {
	return other.ToC() == recv.ToC()
}

// DateTimeNew is a wrapper around the C function g_date_time_new.
func DateTimeNew(tz *TimeZone, year int32, month int32, day int32, hour int32, minute int32, seconds float64) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	c_year := (C.gint)(year)

	c_month := (C.gint)(month)

	c_day := (C.gint)(day)

	c_hour := (C.gint)(hour)

	c_minute := (C.gint)(minute)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_new(c_tz, c_year, c_month, c_day, c_hour, c_minute, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromTimevalLocal is a wrapper around the C function g_date_time_new_from_timeval_local.
func DateTimeNewFromTimevalLocal(tv *TimeVal) *DateTime {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_new_from_timeval_local(c_tv)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromTimevalUtc is a wrapper around the C function g_date_time_new_from_timeval_utc.
func DateTimeNewFromTimevalUtc(tv *TimeVal) *DateTime {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_new_from_timeval_utc(c_tv)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromUnixLocal is a wrapper around the C function g_date_time_new_from_unix_local.
func DateTimeNewFromUnixLocal(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_local(c_t)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromUnixUtc is a wrapper around the C function g_date_time_new_from_unix_utc.
func DateTimeNewFromUnixUtc(t int64) *DateTime {
	c_t := (C.gint64)(t)

	retC := C.g_date_time_new_from_unix_utc(c_t)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewLocal is a wrapper around the C function g_date_time_new_local.
func DateTimeNewLocal(year int32, month int32, day int32, hour int32, minute int32, seconds float64) *DateTime {
	c_year := (C.gint)(year)

	c_month := (C.gint)(month)

	c_day := (C.gint)(day)

	c_hour := (C.gint)(hour)

	c_minute := (C.gint)(minute)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_new_local(c_year, c_month, c_day, c_hour, c_minute, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewNow is a wrapper around the C function g_date_time_new_now.
func DateTimeNewNow(tz *TimeZone) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	retC := C.g_date_time_new_now(c_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewNowLocal is a wrapper around the C function g_date_time_new_now_local.
func DateTimeNewNowLocal() *DateTime {
	retC := C.g_date_time_new_now_local()
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewNowUtc is a wrapper around the C function g_date_time_new_now_utc.
func DateTimeNewNowUtc() *DateTime {
	retC := C.g_date_time_new_now_utc()
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewUtc is a wrapper around the C function g_date_time_new_utc.
func DateTimeNewUtc(year int32, month int32, day int32, hour int32, minute int32, seconds float64) *DateTime {
	c_year := (C.gint)(year)

	c_month := (C.gint)(month)

	c_day := (C.gint)(day)

	c_hour := (C.gint)(hour)

	c_minute := (C.gint)(minute)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_new_utc(c_year, c_month, c_day, c_hour, c_minute, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_date_time_compare : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1
// g_date_time_equal : unsupported parameter dt1 : no type generator for gpointer (gconstpointer) for param dt1
// g_date_time_hash : unsupported parameter datetime : no type generator for gpointer (gconstpointer) for param datetime
// Add is a wrapper around the C function g_date_time_add.
func (recv *DateTime) Add(timespan TimeSpan) *DateTime {
	c_timespan := (C.GTimeSpan)(timespan)

	retC := C.g_date_time_add((*C.GDateTime)(recv.native), c_timespan)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddDays is a wrapper around the C function g_date_time_add_days.
func (recv *DateTime) AddDays(days int32) *DateTime {
	c_days := (C.gint)(days)

	retC := C.g_date_time_add_days((*C.GDateTime)(recv.native), c_days)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFull is a wrapper around the C function g_date_time_add_full.
func (recv *DateTime) AddFull(years int32, months int32, days int32, hours int32, minutes int32, seconds float64) *DateTime {
	c_years := (C.gint)(years)

	c_months := (C.gint)(months)

	c_days := (C.gint)(days)

	c_hours := (C.gint)(hours)

	c_minutes := (C.gint)(minutes)

	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_full((*C.GDateTime)(recv.native), c_years, c_months, c_days, c_hours, c_minutes, c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddHours is a wrapper around the C function g_date_time_add_hours.
func (recv *DateTime) AddHours(hours int32) *DateTime {
	c_hours := (C.gint)(hours)

	retC := C.g_date_time_add_hours((*C.GDateTime)(recv.native), c_hours)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddMinutes is a wrapper around the C function g_date_time_add_minutes.
func (recv *DateTime) AddMinutes(minutes int32) *DateTime {
	c_minutes := (C.gint)(minutes)

	retC := C.g_date_time_add_minutes((*C.GDateTime)(recv.native), c_minutes)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddMonths is a wrapper around the C function g_date_time_add_months.
func (recv *DateTime) AddMonths(months int32) *DateTime {
	c_months := (C.gint)(months)

	retC := C.g_date_time_add_months((*C.GDateTime)(recv.native), c_months)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddSeconds is a wrapper around the C function g_date_time_add_seconds.
func (recv *DateTime) AddSeconds(seconds float64) *DateTime {
	c_seconds := (C.gdouble)(seconds)

	retC := C.g_date_time_add_seconds((*C.GDateTime)(recv.native), c_seconds)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddWeeks is a wrapper around the C function g_date_time_add_weeks.
func (recv *DateTime) AddWeeks(weeks int32) *DateTime {
	c_weeks := (C.gint)(weeks)

	retC := C.g_date_time_add_weeks((*C.GDateTime)(recv.native), c_weeks)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddYears is a wrapper around the C function g_date_time_add_years.
func (recv *DateTime) AddYears(years int32) *DateTime {
	c_years := (C.gint)(years)

	retC := C.g_date_time_add_years((*C.GDateTime)(recv.native), c_years)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Difference is a wrapper around the C function g_date_time_difference.
func (recv *DateTime) Difference(begin *DateTime) TimeSpan {
	c_begin := (*C.GDateTime)(C.NULL)
	if begin != nil {
		c_begin = (*C.GDateTime)(begin.ToC())
	}

	retC := C.g_date_time_difference((*C.GDateTime)(recv.native), c_begin)
	retGo := (TimeSpan)(retC)

	return retGo
}

// Format is a wrapper around the C function g_date_time_format.
func (recv *DateTime) Format(format string) string {
	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	retC := C.g_date_time_format((*C.GDateTime)(recv.native), c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDayOfMonth is a wrapper around the C function g_date_time_get_day_of_month.
func (recv *DateTime) GetDayOfMonth() int32 {
	retC := C.g_date_time_get_day_of_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfWeek is a wrapper around the C function g_date_time_get_day_of_week.
func (recv *DateTime) GetDayOfWeek() int32 {
	retC := C.g_date_time_get_day_of_week((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDayOfYear is a wrapper around the C function g_date_time_get_day_of_year.
func (recv *DateTime) GetDayOfYear() int32 {
	retC := C.g_date_time_get_day_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHour is a wrapper around the C function g_date_time_get_hour.
func (recv *DateTime) GetHour() int32 {
	retC := C.g_date_time_get_hour((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMicrosecond is a wrapper around the C function g_date_time_get_microsecond.
func (recv *DateTime) GetMicrosecond() int32 {
	retC := C.g_date_time_get_microsecond((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinute is a wrapper around the C function g_date_time_get_minute.
func (recv *DateTime) GetMinute() int32 {
	retC := C.g_date_time_get_minute((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMonth is a wrapper around the C function g_date_time_get_month.
func (recv *DateTime) GetMonth() int32 {
	retC := C.g_date_time_get_month((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSecond is a wrapper around the C function g_date_time_get_second.
func (recv *DateTime) GetSecond() int32 {
	retC := C.g_date_time_get_second((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSeconds is a wrapper around the C function g_date_time_get_seconds.
func (recv *DateTime) GetSeconds() float64 {
	retC := C.g_date_time_get_seconds((*C.GDateTime)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetTimezoneAbbreviation is a wrapper around the C function g_date_time_get_timezone_abbreviation.
func (recv *DateTime) GetTimezoneAbbreviation() string {
	retC := C.g_date_time_get_timezone_abbreviation((*C.GDateTime)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUtcOffset is a wrapper around the C function g_date_time_get_utc_offset.
func (recv *DateTime) GetUtcOffset() TimeSpan {
	retC := C.g_date_time_get_utc_offset((*C.GDateTime)(recv.native))
	retGo := (TimeSpan)(retC)

	return retGo
}

// GetWeekNumberingYear is a wrapper around the C function g_date_time_get_week_numbering_year.
func (recv *DateTime) GetWeekNumberingYear() int32 {
	retC := C.g_date_time_get_week_numbering_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWeekOfYear is a wrapper around the C function g_date_time_get_week_of_year.
func (recv *DateTime) GetWeekOfYear() int32 {
	retC := C.g_date_time_get_week_of_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetYear is a wrapper around the C function g_date_time_get_year.
func (recv *DateTime) GetYear() int32 {
	retC := C.g_date_time_get_year((*C.GDateTime)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetYmd is a wrapper around the C function g_date_time_get_ymd.
func (recv *DateTime) GetYmd() (int32, int32, int32) {
	var c_year C.gint

	var c_month C.gint

	var c_day C.gint

	C.g_date_time_get_ymd((*C.GDateTime)(recv.native), &c_year, &c_month, &c_day)

	year := (int32)(c_year)

	month := (int32)(c_month)

	day := (int32)(c_day)

	return year, month, day
}

// IsDaylightSavings is a wrapper around the C function g_date_time_is_daylight_savings.
func (recv *DateTime) IsDaylightSavings() bool {
	retC := C.g_date_time_is_daylight_savings((*C.GDateTime)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_date_time_ref.
func (recv *DateTime) Ref() *DateTime {
	retC := C.g_date_time_ref((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToLocal is a wrapper around the C function g_date_time_to_local.
func (recv *DateTime) ToLocal() *DateTime {
	retC := C.g_date_time_to_local((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToTimeval is a wrapper around the C function g_date_time_to_timeval.
func (recv *DateTime) ToTimeval(tv *TimeVal) bool {
	c_tv := (*C.GTimeVal)(C.NULL)
	if tv != nil {
		c_tv = (*C.GTimeVal)(tv.ToC())
	}

	retC := C.g_date_time_to_timeval((*C.GDateTime)(recv.native), c_tv)
	retGo := retC == C.TRUE

	return retGo
}

// ToTimezone is a wrapper around the C function g_date_time_to_timezone.
func (recv *DateTime) ToTimezone(tz *TimeZone) *DateTime {
	c_tz := (*C.GTimeZone)(C.NULL)
	if tz != nil {
		c_tz = (*C.GTimeZone)(tz.ToC())
	}

	retC := C.g_date_time_to_timezone((*C.GDateTime)(recv.native), c_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToUnix is a wrapper around the C function g_date_time_to_unix.
func (recv *DateTime) ToUnix() int64 {
	retC := C.g_date_time_to_unix((*C.GDateTime)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// ToUtc is a wrapper around the C function g_date_time_to_utc.
func (recv *DateTime) ToUtc() *DateTime {
	retC := C.g_date_time_to_utc((*C.GDateTime)(recv.native))
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_date_time_unref.
func (recv *DateTime) Unref() {
	C.g_date_time_unref((*C.GDateTime)(recv.native))

	return
}

// DebugKey is a wrapper around the C record GDebugKey.
type DebugKey struct {
	native *C.GDebugKey
	Key    string
	Value  uint32
}

func DebugKeyNewFromC(u unsafe.Pointer) *DebugKey {
	c := (*C.GDebugKey)(u)
	if c == nil {
		return nil
	}

	g := &DebugKey{
		Key:    C.GoString(c.key),
		Value:  (uint32)(c.value),
		native: c,
	}

	return g
}

func (recv *DebugKey) ToC() unsafe.Pointer {
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		(C.guint)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DebugKey with another DebugKey, and returns true if they represent the same GObject.
func (recv *DebugKey) Equals(other *DebugKey) bool {
	return other.ToC() == recv.ToC()
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native *C.GDir
}

func DirNewFromC(u unsafe.Pointer) *Dir {
	c := (*C.GDir)(u)
	if c == nil {
		return nil
	}

	g := &Dir{native: c}

	return g
}

func (recv *Dir) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Dir with another Dir, and returns true if they represent the same GObject.
func (recv *Dir) Equals(other *Dir) bool {
	return other.ToC() == recv.ToC()
}

// DirMakeTmp is a wrapper around the C function g_dir_make_tmp.
func DirMakeTmp(tmpl string) (string, error) {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var cThrowableError *C.GError

	retC := C.g_dir_make_tmp(c_tmpl, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DirOpen is a wrapper around the C function g_dir_open.
func DirOpen(path string, flags uint32) (*Dir, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_flags := (C.guint)(flags)

	var cThrowableError *C.GError

	retC := C.g_dir_open(c_path, c_flags, &cThrowableError)
	retGo := DirNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Close is a wrapper around the C function g_dir_close.
func (recv *Dir) Close() {
	C.g_dir_close((*C.GDir)(recv.native))

	return
}

// ReadName is a wrapper around the C function g_dir_read_name.
func (recv *Dir) ReadName() string {
	retC := C.g_dir_read_name((*C.GDir)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Rewind is a wrapper around the C function g_dir_rewind.
func (recv *Dir) Rewind() {
	C.g_dir_rewind((*C.GDir)(recv.native))

	return
}

// Error is a wrapper around the C record GError.
type Error struct {
	native  *C.GError
	Domain  Quark
	Code    int32
	Message string
}

func ErrorNewFromC(u unsafe.Pointer) *Error {
	c := (*C.GError)(u)
	if c == nil {
		return nil
	}

	g := &Error{
		Code:    (int32)(c.code),
		Domain:  (Quark)(c.domain),
		Message: C.GoString(c.message),
		native:  c,
	}

	return g
}

func (recv *Error) ToC() unsafe.Pointer {
	recv.native.domain =
		(C.GQuark)(recv.Domain)
	recv.native.code =
		(C.gint)(recv.Code)
	recv.native.message =
		C.CString(recv.Message)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Error with another Error, and returns true if they represent the same GObject.
func (recv *Error) Equals(other *Error) bool {
	return other.ToC() == recv.ToC()
}

// ErrorNew is a wrapper around the C function g_error_new.
func ErrorNew(domain Quark, code int32, format string, args ...interface{}) *Error {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_error_new(c_domain, c_code, c_format)
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ErrorNewLiteral is a wrapper around the C function g_error_new_literal.
func ErrorNewLiteral(domain Quark, code int32, message string) *Error {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	retC := C.g_error_new_literal(c_domain, c_code, c_message)
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Copy is a wrapper around the C function g_error_copy.
func (recv *Error) Copy() *Error {
	retC := C.g_error_copy((*C.GError)(recv.native))
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_error_free.
func (recv *Error) Free() {
	C.g_error_free((*C.GError)(recv.native))

	return
}

// Matches is a wrapper around the C function g_error_matches.
func (recv *Error) Matches(domain Quark, code int32) bool {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	retC := C.g_error_matches((*C.GError)(recv.native), c_domain, c_code)
	retGo := retC == C.TRUE

	return retGo
}

// HashTable is a wrapper around the C record GHashTable.
type HashTable struct {
	native *C.GHashTable
}

func HashTableNewFromC(u unsafe.Pointer) *HashTable {
	c := (*C.GHashTable)(u)
	if c == nil {
		return nil
	}

	g := &HashTable{native: c}

	return g
}

func (recv *HashTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HashTable with another HashTable, and returns true if they represent the same GObject.
func (recv *HashTable) Equals(other *HashTable) bool {
	return other.ToC() == recv.ToC()
}

// g_hash_table_add : unsupported parameter key : no type generator for gpointer (gpointer) for param key
// g_hash_table_contains : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// HashTableDestroy is a wrapper around the C function g_hash_table_destroy.
func HashTableDestroy(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_destroy(c_hash_table)

	return
}

// g_hash_table_find : unsupported parameter predicate : no type generator for HRFunc (GHRFunc) for param predicate
// g_hash_table_foreach : unsupported parameter func : no type generator for HFunc (GHFunc) for param func
// g_hash_table_foreach_remove : unsupported parameter func : no type generator for HRFunc (GHRFunc) for param func
// g_hash_table_foreach_steal : unsupported parameter func : no type generator for HRFunc (GHRFunc) for param func
// HashTableGetKeys is a wrapper around the C function g_hash_table_get_keys.
func HashTableGetKeys(hashTable *HashTable) *List {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_get_keys(c_hash_table)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HashTableGetValues is a wrapper around the C function g_hash_table_get_values.
func HashTableGetValues(hashTable *HashTable) *List {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_get_values(c_hash_table)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_hash_table_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key
// g_hash_table_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// g_hash_table_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key
// g_hash_table_new : unsupported parameter hash_func : no type generator for HashFunc (GHashFunc) for param hash_func
// g_hash_table_new_full : unsupported parameter hash_func : no type generator for HashFunc (GHashFunc) for param hash_func
// HashTableRef is a wrapper around the C function g_hash_table_ref.
func HashTableRef(hashTable *HashTable) *HashTable {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_ref(c_hash_table)
	retGo := HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_hash_table_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// HashTableRemoveAll is a wrapper around the C function g_hash_table_remove_all.
func HashTableRemoveAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_remove_all(c_hash_table)

	return
}

// g_hash_table_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key
// HashTableSize is a wrapper around the C function g_hash_table_size.
func HashTableSize(hashTable *HashTable) uint32 {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_size(c_hash_table)
	retGo := (uint32)(retC)

	return retGo
}

// g_hash_table_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// HashTableStealAll is a wrapper around the C function g_hash_table_steal_all.
func HashTableStealAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_steal_all(c_hash_table)

	return
}

// HashTableUnref is a wrapper around the C function g_hash_table_unref.
func HashTableUnref(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_unref(c_hash_table)

	return
}

// HashTableIter is a wrapper around the C record GHashTableIter.
type HashTableIter struct {
	native *C.GHashTableIter
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
}

func HashTableIterNewFromC(u unsafe.Pointer) *HashTableIter {
	c := (*C.GHashTableIter)(u)
	if c == nil {
		return nil
	}

	g := &HashTableIter{native: c}

	return g
}

func (recv *HashTableIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HashTableIter with another HashTableIter, and returns true if they represent the same GObject.
func (recv *HashTableIter) Equals(other *HashTableIter) bool {
	return other.ToC() == recv.ToC()
}

// GetHashTable is a wrapper around the C function g_hash_table_iter_get_hash_table.
func (recv *HashTableIter) GetHashTable() *HashTable {
	retC := C.g_hash_table_iter_get_hash_table((*C.GHashTableIter)(recv.native))
	retGo := HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_hash_table_iter_init.
func (recv *HashTableIter) Init(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_iter_init((*C.GHashTableIter)(recv.native), c_hash_table)

	return
}

// Unsupported : g_hash_table_iter_next : unsupported parameter key : no type generator for gpointer (gpointer*) for param key

// Remove is a wrapper around the C function g_hash_table_iter_remove.
func (recv *HashTableIter) Remove() {
	C.g_hash_table_iter_remove((*C.GHashTableIter)(recv.native))

	return
}

// Unsupported : g_hash_table_iter_replace : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Steal is a wrapper around the C function g_hash_table_iter_steal.
func (recv *HashTableIter) Steal() {
	C.g_hash_table_iter_steal((*C.GHashTableIter)(recv.native))

	return
}

// Hmac is a wrapper around the C record GHmac.
type Hmac struct {
	native *C.GHmac
}

func HmacNewFromC(u unsafe.Pointer) *Hmac {
	c := (*C.GHmac)(u)
	if c == nil {
		return nil
	}

	g := &Hmac{native: c}

	return g
}

func (recv *Hmac) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Hmac with another Hmac, and returns true if they represent the same GObject.
func (recv *Hmac) Equals(other *Hmac) bool {
	return other.ToC() == recv.ToC()
}

// HmacNew is a wrapper around the C function g_hmac_new.
func HmacNew(digestType ChecksumType, key []uint8) *Hmac {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key_array := make([]C.guchar, len(key)+1, len(key)+1)
	for i, item := range key {
		c := (C.guchar)(item)
		c_key_array[i] = c
	}
	c_key_array[len(key)] = 0
	c_key_arrayPtr := &c_key_array[0]
	c_key := (*C.guchar)(unsafe.Pointer(c_key_arrayPtr))

	c_key_len := (C.gsize)(len(key))

	retC := C.g_hmac_new(c_digest_type, c_key, c_key_len)
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_hmac_copy.
func (recv *Hmac) Copy() *Hmac {
	retC := C.g_hmac_copy((*C.GHmac)(recv.native))
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDigest is a wrapper around the C function g_hmac_get_digest.
func (recv *Hmac) GetDigest(buffer uint8, digestLen uint64) {
	c_buffer := (C.guint8)(buffer)

	c_digest_len := (C.gsize)(digestLen)

	C.g_hmac_get_digest((*C.GHmac)(recv.native), &c_buffer, &c_digest_len)

	return
}

// GetString is a wrapper around the C function g_hmac_get_string.
func (recv *Hmac) GetString() string {
	retC := C.g_hmac_get_string((*C.GHmac)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Ref is a wrapper around the C function g_hmac_ref.
func (recv *Hmac) Ref() *Hmac {
	retC := C.g_hmac_ref((*C.GHmac)(recv.native))
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_hmac_unref.
func (recv *Hmac) Unref() {
	C.g_hmac_unref((*C.GHmac)(recv.native))

	return
}

// Update is a wrapper around the C function g_hmac_update.
func (recv *Hmac) Update(data []uint8) {
	c_data_array := make([]C.guchar, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guchar)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gssize)(len(data))

	C.g_hmac_update((*C.GHmac)(recv.native), c_data, c_length)

	return
}

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native *C.GHook
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
	RefCount uint32
	HookId   uint64
	Flags    uint32
	// _func : no type generator for gpointer, gpointer
	// destroy : no type generator for DestroyNotify, GDestroyNotify
}

func HookNewFromC(u unsafe.Pointer) *Hook {
	c := (*C.GHook)(u)
	if c == nil {
		return nil
	}

	g := &Hook{
		Flags:    (uint32)(c.flags),
		HookId:   (uint64)(c.hook_id),
		RefCount: (uint32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *Hook) ToC() unsafe.Pointer {
	recv.native.ref_count =
		(C.guint)(recv.RefCount)
	recv.native.hook_id =
		(C.gulong)(recv.HookId)
	recv.native.flags =
		(C.guint)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Hook with another Hook, and returns true if they represent the same GObject.
func (recv *Hook) Equals(other *Hook) bool {
	return other.ToC() == recv.ToC()
}

// HookAlloc is a wrapper around the C function g_hook_alloc.
func HookAlloc(hookList *HookList) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	retC := C.g_hook_alloc(c_hook_list)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookDestroy is a wrapper around the C function g_hook_destroy.
func HookDestroy(hookList *HookList, hookId uint64) bool {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook_id := (C.gulong)(hookId)

	retC := C.g_hook_destroy(c_hook_list, c_hook_id)
	retGo := retC == C.TRUE

	return retGo
}

// HookDestroyLink is a wrapper around the C function g_hook_destroy_link.
func HookDestroyLink(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_destroy_link(c_hook_list, c_hook)

	return
}

// g_hook_find : unsupported parameter func : no type generator for HookFindFunc (GHookFindFunc) for param func
// g_hook_find_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_hook_find_func : unsupported parameter func : no type generator for gpointer (gpointer) for param func
// g_hook_find_func_data : unsupported parameter func : no type generator for gpointer (gpointer) for param func
// HookFirstValid is a wrapper around the C function g_hook_first_valid.
func HookFirstValid(hookList *HookList, mayBeInCall bool) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_may_be_in_call :=
		boolToGboolean(mayBeInCall)

	retC := C.g_hook_first_valid(c_hook_list, c_may_be_in_call)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookFree is a wrapper around the C function g_hook_free.
func HookFree(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_free(c_hook_list, c_hook)

	return
}

// HookGet is a wrapper around the C function g_hook_get.
func HookGet(hookList *HookList, hookId uint64) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook_id := (C.gulong)(hookId)

	retC := C.g_hook_get(c_hook_list, c_hook_id)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookInsertBefore is a wrapper around the C function g_hook_insert_before.
func HookInsertBefore(hookList *HookList, sibling *Hook, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_sibling := (*C.GHook)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GHook)(sibling.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_insert_before(c_hook_list, c_sibling, c_hook)

	return
}

// g_hook_insert_sorted : unsupported parameter func : no type generator for HookCompareFunc (GHookCompareFunc) for param func
// HookNextValid is a wrapper around the C function g_hook_next_valid.
func HookNextValid(hookList *HookList, hook *Hook, mayBeInCall bool) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	c_may_be_in_call :=
		boolToGboolean(mayBeInCall)

	retC := C.g_hook_next_valid(c_hook_list, c_hook, c_may_be_in_call)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookPrepend is a wrapper around the C function g_hook_prepend.
func HookPrepend(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_prepend(c_hook_list, c_hook)

	return
}

// HookRef is a wrapper around the C function g_hook_ref.
func HookRef(hookList *HookList, hook *Hook) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	retC := C.g_hook_ref(c_hook_list, c_hook)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookUnref is a wrapper around the C function g_hook_unref.
func HookUnref(hookList *HookList, hook *Hook) {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_hook := (*C.GHook)(C.NULL)
	if hook != nil {
		c_hook = (*C.GHook)(hook.ToC())
	}

	C.g_hook_unref(c_hook_list, c_hook)

	return
}

// CompareIds is a wrapper around the C function g_hook_compare_ids.
func (recv *Hook) CompareIds(sibling *Hook) int32 {
	c_sibling := (*C.GHook)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GHook)(sibling.ToC())
	}

	retC := C.g_hook_compare_ids((*C.GHook)(recv.native), c_sibling)
	retGo := (int32)(retC)

	return retGo
}

// HookList is a wrapper around the C record GHookList.
type HookList struct {
	native *C.GHookList
	SeqId  uint64
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	// hooks : record
	// dummy3 : no type generator for gpointer, gpointer
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func HookListNewFromC(u unsafe.Pointer) *HookList {
	c := (*C.GHookList)(u)
	if c == nil {
		return nil
	}

	g := &HookList{
		SeqId:  (uint64)(c.seq_id),
		native: c,
	}

	return g
}

func (recv *HookList) ToC() unsafe.Pointer {
	recv.native.seq_id =
		(C.gulong)(recv.SeqId)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HookList with another HookList, and returns true if they represent the same GObject.
func (recv *HookList) Equals(other *HookList) bool {
	return other.ToC() == recv.ToC()
}

// Clear is a wrapper around the C function g_hook_list_clear.
func (recv *HookList) Clear() {
	C.g_hook_list_clear((*C.GHookList)(recv.native))

	return
}

// Init is a wrapper around the C function g_hook_list_init.
func (recv *HookList) Init(hookSize uint32) {
	c_hook_size := (C.guint)(hookSize)

	C.g_hook_list_init((*C.GHookList)(recv.native), c_hook_size)

	return
}

// Invoke is a wrapper around the C function g_hook_list_invoke.
func (recv *HookList) Invoke(mayRecurse bool) {
	c_may_recurse :=
		boolToGboolean(mayRecurse)

	C.g_hook_list_invoke((*C.GHookList)(recv.native), c_may_recurse)

	return
}

// InvokeCheck is a wrapper around the C function g_hook_list_invoke_check.
func (recv *HookList) InvokeCheck(mayRecurse bool) {
	c_may_recurse :=
		boolToGboolean(mayRecurse)

	C.g_hook_list_invoke_check((*C.GHookList)(recv.native), c_may_recurse)

	return
}

// Unsupported : g_hook_list_marshal : unsupported parameter marshaller : no type generator for HookMarshaller (GHookMarshaller) for param marshaller

// Unsupported : g_hook_list_marshal_check : unsupported parameter marshaller : no type generator for HookCheckMarshaller (GHookCheckMarshaller) for param marshaller

// Blacklisted : GIConv

// IOChannel is a wrapper around the C record GIOChannel.
type IOChannel struct {
	native *C.GIOChannel
	// Private : ref_count
	// Private : funcs
	// Private : encoding
	// Private : read_cd
	// Private : write_cd
	// Private : line_term
	// Private : line_term_len
	// Private : buf_size
	// Private : read_buf
	// Private : encoded_read_buf
	// Private : write_buf
	// Private : partial_write_buf
	// Private : use_buffer
	// Private : do_encode
	// Private : close_on_unref
	// Private : is_readable
	// Private : is_writeable
	// Private : is_seekable
	// Private : reserved1
	// Private : reserved2
}

func IOChannelNewFromC(u unsafe.Pointer) *IOChannel {
	c := (*C.GIOChannel)(u)
	if c == nil {
		return nil
	}

	g := &IOChannel{native: c}

	return g
}

func (recv *IOChannel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOChannel with another IOChannel, and returns true if they represent the same GObject.
func (recv *IOChannel) Equals(other *IOChannel) bool {
	return other.ToC() == recv.ToC()
}

// IOChannelNewFile is a wrapper around the C function g_io_channel_new_file.
func IOChannelNewFile(filename string, mode string) (*IOChannel, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_mode := C.CString(mode)
	defer C.free(unsafe.Pointer(c_mode))

	var cThrowableError *C.GError

	retC := C.g_io_channel_new_file(c_filename, c_mode, &cThrowableError)
	retGo := IOChannelNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IOChannelUnixNew is a wrapper around the C function g_io_channel_unix_new.
func IOChannelUnixNew(fd int32) *IOChannel {
	c_fd := (C.int)(fd)

	retC := C.g_io_channel_unix_new(c_fd)
	retGo := IOChannelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IOChannelErrorFromErrno is a wrapper around the C function g_io_channel_error_from_errno.
func IOChannelErrorFromErrno(en int32) IOChannelError {
	c_en := (C.gint)(en)

	retC := C.g_io_channel_error_from_errno(c_en)
	retGo := (IOChannelError)(retC)

	return retGo
}

// IOChannelErrorQuark is a wrapper around the C function g_io_channel_error_quark.
func IOChannelErrorQuark() Quark {
	retC := C.g_io_channel_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Close is a wrapper around the C function g_io_channel_close.
func (recv *IOChannel) Close() {
	C.g_io_channel_close((*C.GIOChannel)(recv.native))

	return
}

// Flush is a wrapper around the C function g_io_channel_flush.
func (recv *IOChannel) Flush() (IOStatus, error) {
	var cThrowableError *C.GError

	retC := C.g_io_channel_flush((*C.GIOChannel)(recv.native), &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetBufferCondition is a wrapper around the C function g_io_channel_get_buffer_condition.
func (recv *IOChannel) GetBufferCondition() IOCondition {
	retC := C.g_io_channel_get_buffer_condition((*C.GIOChannel)(recv.native))
	retGo := (IOCondition)(retC)

	return retGo
}

// GetBufferSize is a wrapper around the C function g_io_channel_get_buffer_size.
func (recv *IOChannel) GetBufferSize() uint64 {
	retC := C.g_io_channel_get_buffer_size((*C.GIOChannel)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetBuffered is a wrapper around the C function g_io_channel_get_buffered.
func (recv *IOChannel) GetBuffered() bool {
	retC := C.g_io_channel_get_buffered((*C.GIOChannel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCloseOnUnref is a wrapper around the C function g_io_channel_get_close_on_unref.
func (recv *IOChannel) GetCloseOnUnref() bool {
	retC := C.g_io_channel_get_close_on_unref((*C.GIOChannel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEncoding is a wrapper around the C function g_io_channel_get_encoding.
func (recv *IOChannel) GetEncoding() string {
	retC := C.g_io_channel_get_encoding((*C.GIOChannel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_io_channel_get_flags.
func (recv *IOChannel) GetFlags() IOFlags {
	retC := C.g_io_channel_get_flags((*C.GIOChannel)(recv.native))
	retGo := (IOFlags)(retC)

	return retGo
}

// GetLineTerm is a wrapper around the C function g_io_channel_get_line_term.
func (recv *IOChannel) GetLineTerm(length int32) string {
	c_length := (C.gint)(length)

	retC := C.g_io_channel_get_line_term((*C.GIOChannel)(recv.native), &c_length)
	retGo := C.GoString(retC)

	return retGo
}

// Init is a wrapper around the C function g_io_channel_init.
func (recv *IOChannel) Init() {
	C.g_io_channel_init((*C.GIOChannel)(recv.native))

	return
}

// Read is a wrapper around the C function g_io_channel_read.
func (recv *IOChannel) Read(buf string, count uint64, bytesRead uint64) IOError {
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))

	c_count := (C.gsize)(count)

	c_bytes_read := (C.gsize)(bytesRead)

	retC := C.g_io_channel_read((*C.GIOChannel)(recv.native), c_buf, c_count, &c_bytes_read)
	retGo := (IOError)(retC)

	return retGo
}

// Unsupported : g_io_channel_read_chars : unsupported parameter buf : output array param buf

// ReadLine is a wrapper around the C function g_io_channel_read_line.
func (recv *IOChannel) ReadLine() (IOStatus, string, uint64, uint64, error) {
	var c_str_return *C.gchar

	var c_length C.gsize

	var c_terminator_pos C.gsize

	var cThrowableError *C.GError

	retC := C.g_io_channel_read_line((*C.GIOChannel)(recv.native), &c_str_return, &c_length, &c_terminator_pos, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	strReturn := C.GoString(c_str_return)
	defer C.free(unsafe.Pointer(c_str_return))

	length := (uint64)(c_length)

	terminatorPos := (uint64)(c_terminator_pos)

	return retGo, strReturn, length, terminatorPos, goError
}

// ReadLineString is a wrapper around the C function g_io_channel_read_line_string.
func (recv *IOChannel) ReadLineString(buffer *String, terminatorPos uint64) (IOStatus, error) {
	c_buffer := (*C.GString)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GString)(buffer.ToC())
	}

	c_terminator_pos := (C.gsize)(terminatorPos)

	var cThrowableError *C.GError

	retC := C.g_io_channel_read_line_string((*C.GIOChannel)(recv.native), c_buffer, &c_terminator_pos, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_io_channel_read_to_end : unsupported parameter str_return : output array param str_return

// ReadUnichar is a wrapper around the C function g_io_channel_read_unichar.
func (recv *IOChannel) ReadUnichar() (IOStatus, rune, error) {
	var c_thechar C.gunichar

	var cThrowableError *C.GError

	retC := C.g_io_channel_read_unichar((*C.GIOChannel)(recv.native), &c_thechar, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	thechar := (rune)(c_thechar)

	return retGo, thechar, goError
}

// Ref is a wrapper around the C function g_io_channel_ref.
func (recv *IOChannel) Ref() *IOChannel {
	retC := C.g_io_channel_ref((*C.GIOChannel)(recv.native))
	retGo := IOChannelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Seek is a wrapper around the C function g_io_channel_seek.
func (recv *IOChannel) Seek(offset int64, type_ SeekType) IOError {
	c_offset := (C.gint64)(offset)

	c_type := (C.GSeekType)(type_)

	retC := C.g_io_channel_seek((*C.GIOChannel)(recv.native), c_offset, c_type)
	retGo := (IOError)(retC)

	return retGo
}

// SeekPosition is a wrapper around the C function g_io_channel_seek_position.
func (recv *IOChannel) SeekPosition(offset int64, type_ SeekType) (IOStatus, error) {
	c_offset := (C.gint64)(offset)

	c_type := (C.GSeekType)(type_)

	var cThrowableError *C.GError

	retC := C.g_io_channel_seek_position((*C.GIOChannel)(recv.native), c_offset, c_type, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetBufferSize is a wrapper around the C function g_io_channel_set_buffer_size.
func (recv *IOChannel) SetBufferSize(size uint64) {
	c_size := (C.gsize)(size)

	C.g_io_channel_set_buffer_size((*C.GIOChannel)(recv.native), c_size)

	return
}

// SetBuffered is a wrapper around the C function g_io_channel_set_buffered.
func (recv *IOChannel) SetBuffered(buffered bool) {
	c_buffered :=
		boolToGboolean(buffered)

	C.g_io_channel_set_buffered((*C.GIOChannel)(recv.native), c_buffered)

	return
}

// SetCloseOnUnref is a wrapper around the C function g_io_channel_set_close_on_unref.
func (recv *IOChannel) SetCloseOnUnref(doClose bool) {
	c_do_close :=
		boolToGboolean(doClose)

	C.g_io_channel_set_close_on_unref((*C.GIOChannel)(recv.native), c_do_close)

	return
}

// SetEncoding is a wrapper around the C function g_io_channel_set_encoding.
func (recv *IOChannel) SetEncoding(encoding string) (IOStatus, error) {
	c_encoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(c_encoding))

	var cThrowableError *C.GError

	retC := C.g_io_channel_set_encoding((*C.GIOChannel)(recv.native), c_encoding, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetFlags is a wrapper around the C function g_io_channel_set_flags.
func (recv *IOChannel) SetFlags(flags IOFlags) (IOStatus, error) {
	c_flags := (C.GIOFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_io_channel_set_flags((*C.GIOChannel)(recv.native), c_flags, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetLineTerm is a wrapper around the C function g_io_channel_set_line_term.
func (recv *IOChannel) SetLineTerm(lineTerm string, length int32) {
	c_line_term := C.CString(lineTerm)
	defer C.free(unsafe.Pointer(c_line_term))

	c_length := (C.gint)(length)

	C.g_io_channel_set_line_term((*C.GIOChannel)(recv.native), c_line_term, c_length)

	return
}

// Shutdown is a wrapper around the C function g_io_channel_shutdown.
func (recv *IOChannel) Shutdown(flush bool) (IOStatus, error) {
	c_flush :=
		boolToGboolean(flush)

	var cThrowableError *C.GError

	retC := C.g_io_channel_shutdown((*C.GIOChannel)(recv.native), c_flush, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnixGetFd is a wrapper around the C function g_io_channel_unix_get_fd.
func (recv *IOChannel) UnixGetFd() int32 {
	retC := C.g_io_channel_unix_get_fd((*C.GIOChannel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unref is a wrapper around the C function g_io_channel_unref.
func (recv *IOChannel) Unref() {
	C.g_io_channel_unref((*C.GIOChannel)(recv.native))

	return
}

// Write is a wrapper around the C function g_io_channel_write.
func (recv *IOChannel) Write(buf string, count uint64, bytesWritten uint64) IOError {
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))

	c_count := (C.gsize)(count)

	c_bytes_written := (C.gsize)(bytesWritten)

	retC := C.g_io_channel_write((*C.GIOChannel)(recv.native), c_buf, c_count, &c_bytes_written)
	retGo := (IOError)(retC)

	return retGo
}

// WriteChars is a wrapper around the C function g_io_channel_write_chars.
func (recv *IOChannel) WriteChars(buf []uint8, count int64) (IOStatus, uint64, error) {
	c_buf_array := make([]C.guint8, len(buf)+1, len(buf)+1)
	for i, item := range buf {
		c := (C.guint8)(item)
		c_buf_array[i] = c
	}
	c_buf_array[len(buf)] = 0
	c_buf_arrayPtr := &c_buf_array[0]
	c_buf := (*C.gchar)(unsafe.Pointer(c_buf_arrayPtr))

	c_count := (C.gssize)(count)

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_io_channel_write_chars((*C.GIOChannel)(recv.native), c_buf, c_count, &c_bytes_written, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goError
}

// WriteUnichar is a wrapper around the C function g_io_channel_write_unichar.
func (recv *IOChannel) WriteUnichar(thechar rune) (IOStatus, error) {
	c_thechar := (C.gunichar)(thechar)

	var cThrowableError *C.GError

	retC := C.g_io_channel_write_unichar((*C.GIOChannel)(recv.native), c_thechar, &cThrowableError)
	retGo := (IOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IOFuncs is a wrapper around the C record GIOFuncs.
type IOFuncs struct {
	native *C.GIOFuncs
	// no type for io_read
	// no type for io_write
	// no type for io_seek
	// no type for io_close
	// no type for io_create_watch
	// no type for io_free
	// no type for io_set_flags
	// no type for io_get_flags
}

func IOFuncsNewFromC(u unsafe.Pointer) *IOFuncs {
	c := (*C.GIOFuncs)(u)
	if c == nil {
		return nil
	}

	g := &IOFuncs{native: c}

	return g
}

func (recv *IOFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOFuncs with another IOFuncs, and returns true if they represent the same GObject.
func (recv *IOFuncs) Equals(other *IOFuncs) bool {
	return other.ToC() == recv.ToC()
}

// KeyFile is a wrapper around the C record GKeyFile.
type KeyFile struct {
	native *C.GKeyFile
}

func KeyFileNewFromC(u unsafe.Pointer) *KeyFile {
	c := (*C.GKeyFile)(u)
	if c == nil {
		return nil
	}

	g := &KeyFile{native: c}

	return g
}

func (recv *KeyFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this KeyFile with another KeyFile, and returns true if they represent the same GObject.
func (recv *KeyFile) Equals(other *KeyFile) bool {
	return other.ToC() == recv.ToC()
}

// KeyFileNew is a wrapper around the C function g_key_file_new.
func KeyFileNew() *KeyFile {
	retC := C.g_key_file_new()
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// KeyFileErrorQuark is a wrapper around the C function g_key_file_error_quark.
func KeyFileErrorQuark() Quark {
	retC := C.g_key_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Free is a wrapper around the C function g_key_file_free.
func (recv *KeyFile) Free() {
	C.g_key_file_free((*C.GKeyFile)(recv.native))

	return
}

// GetBoolean is a wrapper around the C function g_key_file_get_boolean.
func (recv *KeyFile) GetBoolean(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_boolean((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_key_file_get_boolean_list : array return type :

// GetComment is a wrapper around the C function g_key_file_get_comment.
func (recv *KeyFile) GetComment(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetDouble is a wrapper around the C function g_key_file_get_double.
func (recv *KeyFile) GetDouble(groupName string, key string) (float64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_double((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (float64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_key_file_get_double_list : array return type :

// GetGroups is a wrapper around the C function g_key_file_get_groups.
func (recv *KeyFile) GetGroups() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_key_file_get_groups((*C.GKeyFile)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetInt64 is a wrapper around the C function g_key_file_get_int64.
func (recv *KeyFile) GetInt64(groupName string, key string) (int64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_int64((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetInteger is a wrapper around the C function g_key_file_get_integer.
func (recv *KeyFile) GetInteger(groupName string, key string) (int32, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_integer((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_key_file_get_integer_list : array return type :

// GetKeys is a wrapper around the C function g_key_file_get_keys.
func (recv *KeyFile) GetKeys(groupName string) ([]string, uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_get_keys((*C.GKeyFile)(recv.native), c_group_name, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetLocaleString is a wrapper around the C function g_key_file_get_locale_string.
func (recv *KeyFile) GetLocaleString(groupName string, key string, locale string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_locale_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetLocaleStringList is a wrapper around the C function g_key_file_get_locale_string_list.
func (recv *KeyFile) GetLocaleStringList(groupName string, key string, locale string) ([]string, uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_get_locale_string_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetStartGroup is a wrapper around the C function g_key_file_get_start_group.
func (recv *KeyFile) GetStartGroup() string {
	retC := C.g_key_file_get_start_group((*C.GKeyFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetString is a wrapper around the C function g_key_file_get_string.
func (recv *KeyFile) GetString(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_string((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetStringList is a wrapper around the C function g_key_file_get_string_list.
func (recv *KeyFile) GetStringList(groupName string, key string) ([]string, uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_get_string_list((*C.GKeyFile)(recv.native), c_group_name, c_key, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetUint64 is a wrapper around the C function g_key_file_get_uint64.
func (recv *KeyFile) GetUint64(groupName string, key string) (uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_uint64((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (uint64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetValue is a wrapper around the C function g_key_file_get_value.
func (recv *KeyFile) GetValue(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_value((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasGroup is a wrapper around the C function g_key_file_has_group.
func (recv *KeyFile) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.g_key_file_has_group((*C.GKeyFile)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// HasKey is a wrapper around the C function g_key_file_has_key.
func (recv *KeyFile) HasKey(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_has_key((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadFromData is a wrapper around the C function g_key_file_load_from_data.
func (recv *KeyFile) LoadFromData(data string, length uint64, flags KeyFileFlags) (bool, error) {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	c_length := (C.gsize)(length)

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_data((*C.GKeyFile)(recv.native), c_data, c_length, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadFromDataDirs is a wrapper around the C function g_key_file_load_from_data_dirs.
func (recv *KeyFile) LoadFromDataDirs(file string, flags KeyFileFlags) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var c_full_path *C.gchar

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_data_dirs((*C.GKeyFile)(recv.native), c_file, &c_full_path, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goError
}

// LoadFromDirs is a wrapper around the C function g_key_file_load_from_dirs.
func (recv *KeyFile) LoadFromDirs(file string, searchDirs []string, flags KeyFileFlags) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_search_dirs_array := make([]*C.gchar, len(searchDirs)+1, len(searchDirs)+1)
	for i, item := range searchDirs {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_search_dirs_array[i] = c
	}
	c_search_dirs_array[len(searchDirs)] = nil
	c_search_dirs_arrayPtr := &c_search_dirs_array[0]
	c_search_dirs := (**C.gchar)(unsafe.Pointer(c_search_dirs_arrayPtr))

	var c_full_path *C.gchar

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_dirs((*C.GKeyFile)(recv.native), c_file, c_search_dirs, &c_full_path, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goError
}

// LoadFromFile is a wrapper around the C function g_key_file_load_from_file.
func (recv *KeyFile) LoadFromFile(file string, flags KeyFileFlags) (bool, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_file((*C.GKeyFile)(recv.native), c_file, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Ref is a wrapper around the C function g_key_file_ref.
func (recv *KeyFile) Ref() *KeyFile {
	retC := C.g_key_file_ref((*C.GKeyFile)(recv.native))
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveComment is a wrapper around the C function g_key_file_remove_comment.
func (recv *KeyFile) RemoveComment(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveGroup is a wrapper around the C function g_key_file_remove_group.
func (recv *KeyFile) RemoveGroup(groupName string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_group((*C.GKeyFile)(recv.native), c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveKey is a wrapper around the C function g_key_file_remove_key.
func (recv *KeyFile) RemoveKey(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_key((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetBoolean is a wrapper around the C function g_key_file_set_boolean.
func (recv *KeyFile) SetBoolean(groupName string, key string, value bool) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value :=
		boolToGboolean(value)

	C.g_key_file_set_boolean((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetBooleanList is a wrapper around the C function g_key_file_set_boolean_list.
func (recv *KeyFile) SetBooleanList(groupName string, key string, list []bool) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list_array := make([]C.gboolean, len(list)+1, len(list)+1)
	for i, item := range list {
		c :=
			boolToGboolean(item)
		c_list_array[i] = c
	}
	c_list_array[len(list)] = 0
	c_list_arrayPtr := &c_list_array[0]
	c_list := (*C.gboolean)(unsafe.Pointer(c_list_arrayPtr))

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_boolean_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_list, c_length)

	return
}

// SetComment is a wrapper around the C function g_key_file_set_comment.
func (recv *KeyFile) SetComment(groupName string, key string, comment string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_comment := C.CString(comment)
	defer C.free(unsafe.Pointer(c_comment))

	var cThrowableError *C.GError

	retC := C.g_key_file_set_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, c_comment, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetDouble is a wrapper around the C function g_key_file_set_double.
func (recv *KeyFile) SetDouble(groupName string, key string, value float64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	C.g_key_file_set_double((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetDoubleList is a wrapper around the C function g_key_file_set_double_list.
func (recv *KeyFile) SetDoubleList(groupName string, key string, list []float64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list_array := make([]C.gdouble, len(list)+1, len(list)+1)
	for i, item := range list {
		c := (C.gdouble)(item)
		c_list_array[i] = c
	}
	c_list_array[len(list)] = 0
	c_list_arrayPtr := &c_list_array[0]
	c_list := (*C.gdouble)(unsafe.Pointer(c_list_arrayPtr))

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_double_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_list, c_length)

	return
}

// SetInt64 is a wrapper around the C function g_key_file_set_int64.
func (recv *KeyFile) SetInt64(groupName string, key string, value int64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint64)(value)

	C.g_key_file_set_int64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetInteger is a wrapper around the C function g_key_file_set_integer.
func (recv *KeyFile) SetInteger(groupName string, key string, value int32) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	C.g_key_file_set_integer((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetIntegerList is a wrapper around the C function g_key_file_set_integer_list.
func (recv *KeyFile) SetIntegerList(groupName string, key string, list []int32) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list_array := make([]C.gint, len(list)+1, len(list)+1)
	for i, item := range list {
		c := (C.gint)(item)
		c_list_array[i] = c
	}
	c_list_array[len(list)] = 0
	c_list_arrayPtr := &c_list_array[0]
	c_list := (*C.gint)(unsafe.Pointer(c_list_arrayPtr))

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_integer_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_list, c_length)

	return
}

// SetListSeparator is a wrapper around the C function g_key_file_set_list_separator.
func (recv *KeyFile) SetListSeparator(separator rune) {
	c_separator := (C.gchar)(separator)

	C.g_key_file_set_list_separator((*C.GKeyFile)(recv.native), c_separator)

	return
}

// SetLocaleString is a wrapper around the C function g_key_file_set_locale_string.
func (recv *KeyFile) SetLocaleString(groupName string, key string, locale string, string_ string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	C.g_key_file_set_locale_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, c_string)

	return
}

// Blacklisted : g_key_file_set_locale_string_list

// SetString is a wrapper around the C function g_key_file_set_string.
func (recv *KeyFile) SetString(groupName string, key string, string_ string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	C.g_key_file_set_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_string)

	return
}

// Blacklisted : g_key_file_set_string_list

// SetUint64 is a wrapper around the C function g_key_file_set_uint64.
func (recv *KeyFile) SetUint64(groupName string, key string, value uint64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint64)(value)

	C.g_key_file_set_uint64((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetValue is a wrapper around the C function g_key_file_set_value.
func (recv *KeyFile) SetValue(groupName string, key string, value string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_key_file_set_value((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// ToData is a wrapper around the C function g_key_file_to_data.
func (recv *KeyFile) ToData() (string, uint64, error) {
	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_to_data((*C.GKeyFile)(recv.native), &c_length, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// Unref is a wrapper around the C function g_key_file_unref.
func (recv *KeyFile) Unref() {
	C.g_key_file_unref((*C.GKeyFile)(recv.native))

	return
}

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
}

func ListNewFromC(u unsafe.Pointer) *List {
	c := (*C.GList)(u)
	if c == nil {
		return nil
	}

	g := &List{native: c}

	return g
}

func (recv *List) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this List with another List, and returns true if they represent the same GObject.
func (recv *List) Equals(other *List) bool {
	return other.ToC() == recv.ToC()
}

// ListAlloc is a wrapper around the C function g_list_alloc.
func ListAlloc() *List {
	retC := C.g_list_alloc()
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// ListConcat is a wrapper around the C function g_list_concat.
func ListConcat(list1 *List, list2 *List) *List {
	c_list1 := (*C.GList)(C.NULL)
	if list1 != nil {
		c_list1 = (*C.GList)(list1.ToC())
	}

	c_list2 := (*C.GList)(C.NULL)
	if list2 != nil {
		c_list2 = (*C.GList)(list2.ToC())
	}

	retC := C.g_list_concat(c_list1, c_list2)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListCopy is a wrapper around the C function g_list_copy.
func ListCopy(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_copy(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_copy_deep : unsupported parameter func : no type generator for CopyFunc (GCopyFunc) for param func
// ListDeleteLink is a wrapper around the C function g_list_delete_link.
func ListDeleteLink(list *List, link *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	retC := C.g_list_delete_link(c_list, c_link_)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// ListFirst is a wrapper around the C function g_list_first.
func ListFirst(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_first(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func
// ListFree is a wrapper around the C function g_list_free.
func ListFree(list *List) {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	C.g_list_free(c_list)

	return
}

// ListFree1 is a wrapper around the C function g_list_free_1.
func ListFree1(list *List) {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	C.g_list_free_1(c_list)

	return
}

// g_list_free_full : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func
// g_list_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_insert : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_insert_sorted_with_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// ListLast is a wrapper around the C function g_list_last.
func ListLast(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_last(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListLength is a wrapper around the C function g_list_length.
func ListLength(list *List) uint32 {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_length(c_list)
	retGo := (uint32)(retC)

	return retGo
}

// ListNth is a wrapper around the C function g_list_nth.
func ListNth(list *List, n uint32) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_list_nth(c_list, c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_nth_data : no return generator
// ListNthPrev is a wrapper around the C function g_list_nth_prev.
func ListNthPrev(list *List, n uint32) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_list_nth_prev(c_list, c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListPosition is a wrapper around the C function g_list_position.
func ListPosition(list *List, llink *List) int32 {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_llink := (*C.GList)(C.NULL)
	if llink != nil {
		c_llink = (*C.GList)(llink.ToC())
	}

	retC := C.g_list_position(c_list, c_llink)
	retGo := (int32)(retC)

	return retGo
}

// g_list_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_list_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_list_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// ListRemoveLink is a wrapper around the C function g_list_remove_link.
func ListRemoveLink(list *List, llink *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_llink := (*C.GList)(C.NULL)
	if llink != nil {
		c_llink = (*C.GList)(llink.ToC())
	}

	retC := C.g_list_remove_link(c_list, c_llink)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListReverse is a wrapper around the C function g_list_reverse.
func ListReverse(list *List) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	retC := C.g_list_reverse(c_list)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_sort : unsupported parameter compare_func : no type generator for CompareFunc (GCompareFunc) for param compare_func
// g_list_sort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func
// MainContext is a wrapper around the C record GMainContext.
type MainContext struct {
	native *C.GMainContext
}

func MainContextNewFromC(u unsafe.Pointer) *MainContext {
	c := (*C.GMainContext)(u)
	if c == nil {
		return nil
	}

	g := &MainContext{native: c}

	return g
}

func (recv *MainContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MainContext with another MainContext, and returns true if they represent the same GObject.
func (recv *MainContext) Equals(other *MainContext) bool {
	return other.ToC() == recv.ToC()
}

// MainContextNew is a wrapper around the C function g_main_context_new.
func MainContextNew() *MainContext {
	retC := C.g_main_context_new()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MainContextDefault is a wrapper around the C function g_main_context_default.
func MainContextDefault() *MainContext {
	retC := C.g_main_context_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MainContextGetThreadDefault is a wrapper around the C function g_main_context_get_thread_default.
func MainContextGetThreadDefault() *MainContext {
	retC := C.g_main_context_get_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MainContextRefThreadDefault is a wrapper around the C function g_main_context_ref_thread_default.
func MainContextRefThreadDefault() *MainContext {
	retC := C.g_main_context_ref_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Acquire is a wrapper around the C function g_main_context_acquire.
func (recv *MainContext) Acquire() bool {
	retC := C.g_main_context_acquire((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AddPoll is a wrapper around the C function g_main_context_add_poll.
func (recv *MainContext) AddPoll(fd *PollFD, priority int32) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	c_priority := (C.gint)(priority)

	C.g_main_context_add_poll((*C.GMainContext)(recv.native), c_fd, c_priority)

	return
}

// Unsupported : g_main_context_check : unsupported parameter fds :

// Dispatch is a wrapper around the C function g_main_context_dispatch.
func (recv *MainContext) Dispatch() {
	C.g_main_context_dispatch((*C.GMainContext)(recv.native))

	return
}

// Unsupported : g_main_context_find_source_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// FindSourceById is a wrapper around the C function g_main_context_find_source_by_id.
func (recv *MainContext) FindSourceById(sourceId uint32) *Source {
	c_source_id := (C.guint)(sourceId)

	retC := C.g_main_context_find_source_by_id((*C.GMainContext)(recv.native), c_source_id)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_context_find_source_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_main_context_get_poll_func : no return generator

// Unsupported : g_main_context_invoke : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_main_context_invoke_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// IsOwner is a wrapper around the C function g_main_context_is_owner.
func (recv *MainContext) IsOwner() bool {
	retC := C.g_main_context_is_owner((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Iteration is a wrapper around the C function g_main_context_iteration.
func (recv *MainContext) Iteration(mayBlock bool) bool {
	c_may_block :=
		boolToGboolean(mayBlock)

	retC := C.g_main_context_iteration((*C.GMainContext)(recv.native), c_may_block)
	retGo := retC == C.TRUE

	return retGo
}

// Pending is a wrapper around the C function g_main_context_pending.
func (recv *MainContext) Pending() bool {
	retC := C.g_main_context_pending((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PopThreadDefault is a wrapper around the C function g_main_context_pop_thread_default.
func (recv *MainContext) PopThreadDefault() {
	C.g_main_context_pop_thread_default((*C.GMainContext)(recv.native))

	return
}

// Prepare is a wrapper around the C function g_main_context_prepare.
func (recv *MainContext) Prepare(priority int32) bool {
	c_priority := (C.gint)(priority)

	retC := C.g_main_context_prepare((*C.GMainContext)(recv.native), &c_priority)
	retGo := retC == C.TRUE

	return retGo
}

// PushThreadDefault is a wrapper around the C function g_main_context_push_thread_default.
func (recv *MainContext) PushThreadDefault() {
	C.g_main_context_push_thread_default((*C.GMainContext)(recv.native))

	return
}

// Unsupported : g_main_context_query : unsupported parameter fds : output array param fds

// Ref is a wrapper around the C function g_main_context_ref.
func (recv *MainContext) Ref() *MainContext {
	retC := C.g_main_context_ref((*C.GMainContext)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Release is a wrapper around the C function g_main_context_release.
func (recv *MainContext) Release() {
	C.g_main_context_release((*C.GMainContext)(recv.native))

	return
}

// RemovePoll is a wrapper around the C function g_main_context_remove_poll.
func (recv *MainContext) RemovePoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_main_context_remove_poll((*C.GMainContext)(recv.native), c_fd)

	return
}

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc (GPollFunc) for param func

// Unref is a wrapper around the C function g_main_context_unref.
func (recv *MainContext) Unref() {
	C.g_main_context_unref((*C.GMainContext)(recv.native))

	return
}

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// Wakeup is a wrapper around the C function g_main_context_wakeup.
func (recv *MainContext) Wakeup() {
	C.g_main_context_wakeup((*C.GMainContext)(recv.native))

	return
}

// MainLoop is a wrapper around the C record GMainLoop.
type MainLoop struct {
	native *C.GMainLoop
}

func MainLoopNewFromC(u unsafe.Pointer) *MainLoop {
	c := (*C.GMainLoop)(u)
	if c == nil {
		return nil
	}

	g := &MainLoop{native: c}

	return g
}

func (recv *MainLoop) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MainLoop with another MainLoop, and returns true if they represent the same GObject.
func (recv *MainLoop) Equals(other *MainLoop) bool {
	return other.ToC() == recv.ToC()
}

// MainLoopNew is a wrapper around the C function g_main_loop_new.
func MainLoopNew(context *MainContext, isRunning bool) *MainLoop {
	c_context := (*C.GMainContext)(C.NULL)
	if context != nil {
		c_context = (*C.GMainContext)(context.ToC())
	}

	c_is_running :=
		boolToGboolean(isRunning)

	retC := C.g_main_loop_new(c_context, c_is_running)
	retGo := MainLoopNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContext is a wrapper around the C function g_main_loop_get_context.
func (recv *MainLoop) GetContext() *MainContext {
	retC := C.g_main_loop_get_context((*C.GMainLoop)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsRunning is a wrapper around the C function g_main_loop_is_running.
func (recv *MainLoop) IsRunning() bool {
	retC := C.g_main_loop_is_running((*C.GMainLoop)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Quit is a wrapper around the C function g_main_loop_quit.
func (recv *MainLoop) Quit() {
	C.g_main_loop_quit((*C.GMainLoop)(recv.native))

	return
}

// Ref is a wrapper around the C function g_main_loop_ref.
func (recv *MainLoop) Ref() *MainLoop {
	retC := C.g_main_loop_ref((*C.GMainLoop)(recv.native))
	retGo := MainLoopNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Run is a wrapper around the C function g_main_loop_run.
func (recv *MainLoop) Run() {
	C.g_main_loop_run((*C.GMainLoop)(recv.native))

	return
}

// Unref is a wrapper around the C function g_main_loop_unref.
func (recv *MainLoop) Unref() {
	C.g_main_loop_unref((*C.GMainLoop)(recv.native))

	return
}

// MappedFile is a wrapper around the C record GMappedFile.
type MappedFile struct {
	native *C.GMappedFile
}

func MappedFileNewFromC(u unsafe.Pointer) *MappedFile {
	c := (*C.GMappedFile)(u)
	if c == nil {
		return nil
	}

	g := &MappedFile{native: c}

	return g
}

func (recv *MappedFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MappedFile with another MappedFile, and returns true if they represent the same GObject.
func (recv *MappedFile) Equals(other *MappedFile) bool {
	return other.ToC() == recv.ToC()
}

// MappedFileNew is a wrapper around the C function g_mapped_file_new.
func MappedFileNew(filename string, writable bool) (*MappedFile, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new(c_filename, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MappedFileNewFromFd is a wrapper around the C function g_mapped_file_new_from_fd.
func MappedFileNewFromFd(fd int32, writable bool) (*MappedFile, error) {
	c_fd := (C.gint)(fd)

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new_from_fd(c_fd, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Free is a wrapper around the C function g_mapped_file_free.
func (recv *MappedFile) Free() {
	C.g_mapped_file_free((*C.GMappedFile)(recv.native))

	return
}

// GetBytes is a wrapper around the C function g_mapped_file_get_bytes.
func (recv *MappedFile) GetBytes() *Bytes {
	retC := C.g_mapped_file_get_bytes((*C.GMappedFile)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContents is a wrapper around the C function g_mapped_file_get_contents.
func (recv *MappedFile) GetContents() string {
	retC := C.g_mapped_file_get_contents((*C.GMappedFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetLength is a wrapper around the C function g_mapped_file_get_length.
func (recv *MappedFile) GetLength() uint64 {
	retC := C.g_mapped_file_get_length((*C.GMappedFile)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Ref is a wrapper around the C function g_mapped_file_ref.
func (recv *MappedFile) Ref() *MappedFile {
	retC := C.g_mapped_file_ref((*C.GMappedFile)(recv.native))
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_mapped_file_unref.
func (recv *MappedFile) Unref() {
	C.g_mapped_file_unref((*C.GMappedFile)(recv.native))

	return
}

// MarkupParseContext is a wrapper around the C record GMarkupParseContext.
type MarkupParseContext struct {
	native *C.GMarkupParseContext
}

func MarkupParseContextNewFromC(u unsafe.Pointer) *MarkupParseContext {
	c := (*C.GMarkupParseContext)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParseContext{native: c}

	return g
}

func (recv *MarkupParseContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkupParseContext with another MarkupParseContext, and returns true if they represent the same GObject.
func (recv *MarkupParseContext) Equals(other *MarkupParseContext) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// EndParse is a wrapper around the C function g_markup_parse_context_end_parse.
func (recv *MarkupParseContext) EndParse() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_markup_parse_context_end_parse((*C.GMarkupParseContext)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Free is a wrapper around the C function g_markup_parse_context_free.
func (recv *MarkupParseContext) Free() {
	C.g_markup_parse_context_free((*C.GMarkupParseContext)(recv.native))

	return
}

// GetElement is a wrapper around the C function g_markup_parse_context_get_element.
func (recv *MarkupParseContext) GetElement() string {
	retC := C.g_markup_parse_context_get_element((*C.GMarkupParseContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetElementStack is a wrapper around the C function g_markup_parse_context_get_element_stack.
func (recv *MarkupParseContext) GetElementStack() *SList {
	retC := C.g_markup_parse_context_get_element_stack((*C.GMarkupParseContext)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPosition is a wrapper around the C function g_markup_parse_context_get_position.
func (recv *MarkupParseContext) GetPosition(lineNumber int32, charNumber int32) {
	c_line_number := (C.gint)(lineNumber)

	c_char_number := (C.gint)(charNumber)

	C.g_markup_parse_context_get_position((*C.GMarkupParseContext)(recv.native), &c_line_number, &c_char_number)

	return
}

// Unsupported : g_markup_parse_context_get_user_data : no return generator

// Parse is a wrapper around the C function g_markup_parse_context_parse.
func (recv *MarkupParseContext) Parse(text string, textLen int64) (bool, error) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_text_len := (C.gssize)(textLen)

	var cThrowableError *C.GError

	retC := C.g_markup_parse_context_parse((*C.GMarkupParseContext)(recv.native), c_text, c_text_len, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_markup_parse_context_pop : no return generator

// Unsupported : g_markup_parse_context_push : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// MarkupParser is a wrapper around the C record GMarkupParser.
type MarkupParser struct {
	native *C.GMarkupParser
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

func MarkupParserNewFromC(u unsafe.Pointer) *MarkupParser {
	c := (*C.GMarkupParser)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParser{native: c}

	return g
}

func (recv *MarkupParser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkupParser with another MarkupParser, and returns true if they represent the same GObject.
func (recv *MarkupParser) Equals(other *MarkupParser) bool {
	return other.ToC() == recv.ToC()
}

// MatchInfo is a wrapper around the C record GMatchInfo.
type MatchInfo struct {
	native *C.GMatchInfo
}

func MatchInfoNewFromC(u unsafe.Pointer) *MatchInfo {
	c := (*C.GMatchInfo)(u)
	if c == nil {
		return nil
	}

	g := &MatchInfo{native: c}

	return g
}

func (recv *MatchInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MatchInfo with another MatchInfo, and returns true if they represent the same GObject.
func (recv *MatchInfo) Equals(other *MatchInfo) bool {
	return other.ToC() == recv.ToC()
}

// ExpandReferences is a wrapper around the C function g_match_info_expand_references.
func (recv *MatchInfo) ExpandReferences(stringToExpand string) (string, error) {
	c_string_to_expand := C.CString(stringToExpand)
	defer C.free(unsafe.Pointer(c_string_to_expand))

	var cThrowableError *C.GError

	retC := C.g_match_info_expand_references((*C.GMatchInfo)(recv.native), c_string_to_expand, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Fetch is a wrapper around the C function g_match_info_fetch.
func (recv *MatchInfo) Fetch(matchNum int32) string {
	c_match_num := (C.gint)(matchNum)

	retC := C.g_match_info_fetch((*C.GMatchInfo)(recv.native), c_match_num)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FetchAll is a wrapper around the C function g_match_info_fetch_all.
func (recv *MatchInfo) FetchAll() []string {
	retC := C.g_match_info_fetch_all((*C.GMatchInfo)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// FetchNamed is a wrapper around the C function g_match_info_fetch_named.
func (recv *MatchInfo) FetchNamed(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_match_info_fetch_named((*C.GMatchInfo)(recv.native), c_name)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FetchNamedPos is a wrapper around the C function g_match_info_fetch_named_pos.
func (recv *MatchInfo) FetchNamedPos(name string) (bool, int32, int32) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_start_pos C.gint

	var c_end_pos C.gint

	retC := C.g_match_info_fetch_named_pos((*C.GMatchInfo)(recv.native), c_name, &c_start_pos, &c_end_pos)
	retGo := retC == C.TRUE

	startPos := (int32)(c_start_pos)

	endPos := (int32)(c_end_pos)

	return retGo, startPos, endPos
}

// FetchPos is a wrapper around the C function g_match_info_fetch_pos.
func (recv *MatchInfo) FetchPos(matchNum int32) (bool, int32, int32) {
	c_match_num := (C.gint)(matchNum)

	var c_start_pos C.gint

	var c_end_pos C.gint

	retC := C.g_match_info_fetch_pos((*C.GMatchInfo)(recv.native), c_match_num, &c_start_pos, &c_end_pos)
	retGo := retC == C.TRUE

	startPos := (int32)(c_start_pos)

	endPos := (int32)(c_end_pos)

	return retGo, startPos, endPos
}

// Free is a wrapper around the C function g_match_info_free.
func (recv *MatchInfo) Free() {
	C.g_match_info_free((*C.GMatchInfo)(recv.native))

	return
}

// GetMatchCount is a wrapper around the C function g_match_info_get_match_count.
func (recv *MatchInfo) GetMatchCount() int32 {
	retC := C.g_match_info_get_match_count((*C.GMatchInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_match_info_get_regex : return type : Blacklisted record : GRegex

// GetString is a wrapper around the C function g_match_info_get_string.
func (recv *MatchInfo) GetString() string {
	retC := C.g_match_info_get_string((*C.GMatchInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IsPartialMatch is a wrapper around the C function g_match_info_is_partial_match.
func (recv *MatchInfo) IsPartialMatch() bool {
	retC := C.g_match_info_is_partial_match((*C.GMatchInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Matches is a wrapper around the C function g_match_info_matches.
func (recv *MatchInfo) Matches() bool {
	retC := C.g_match_info_matches((*C.GMatchInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Next is a wrapper around the C function g_match_info_next.
func (recv *MatchInfo) Next() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_match_info_next((*C.GMatchInfo)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Ref is a wrapper around the C function g_match_info_ref.
func (recv *MatchInfo) Ref() *MatchInfo {
	retC := C.g_match_info_ref((*C.GMatchInfo)(recv.native))
	retGo := MatchInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_match_info_unref.
func (recv *MatchInfo) Unref() {
	C.g_match_info_unref((*C.GMatchInfo)(recv.native))

	return
}

// MemVTable is a wrapper around the C record GMemVTable.
type MemVTable struct {
	native *C.GMemVTable
	// no type for malloc
	// no type for realloc
	// no type for free
	// no type for calloc
	// no type for try_malloc
	// no type for try_realloc
}

func MemVTableNewFromC(u unsafe.Pointer) *MemVTable {
	c := (*C.GMemVTable)(u)
	if c == nil {
		return nil
	}

	g := &MemVTable{native: c}

	return g
}

func (recv *MemVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemVTable with another MemVTable, and returns true if they represent the same GObject.
func (recv *MemVTable) Equals(other *MemVTable) bool {
	return other.ToC() == recv.ToC()
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native *C.GNode
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
	// parent : record
	// children : record
}

func NodeNewFromC(u unsafe.Pointer) *Node {
	c := (*C.GNode)(u)
	if c == nil {
		return nil
	}

	g := &Node{native: c}

	return g
}

func (recv *Node) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Node with another Node, and returns true if they represent the same GObject.
func (recv *Node) Equals(other *Node) bool {
	return other.ToC() == recv.ToC()
}

// g_node_new : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Unsupported : g_node_child_index : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// ChildPosition is a wrapper around the C function g_node_child_position.
func (recv *Node) ChildPosition(child *Node) int32 {
	c_child := (*C.GNode)(C.NULL)
	if child != nil {
		c_child = (*C.GNode)(child.ToC())
	}

	retC := C.g_node_child_position((*C.GNode)(recv.native), c_child)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_node_children_foreach : unsupported parameter func : no type generator for NodeForeachFunc (GNodeForeachFunc) for param func

// Copy is a wrapper around the C function g_node_copy.
func (recv *Node) Copy() *Node {
	retC := C.g_node_copy((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc (GCopyFunc) for param copy_func

// Depth is a wrapper around the C function g_node_depth.
func (recv *Node) Depth() uint32 {
	retC := C.g_node_depth((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Destroy is a wrapper around the C function g_node_destroy.
func (recv *Node) Destroy() {
	C.g_node_destroy((*C.GNode)(recv.native))

	return
}

// Unsupported : g_node_find : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_node_find_child : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// FirstSibling is a wrapper around the C function g_node_first_sibling.
func (recv *Node) FirstSibling() *Node {
	retC := C.g_node_first_sibling((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRoot is a wrapper around the C function g_node_get_root.
func (recv *Node) GetRoot() *Node {
	retC := C.g_node_get_root((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function g_node_insert.
func (recv *Node) Insert(position int32, node *Node) *Node {
	c_position := (C.gint)(position)

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert((*C.GNode)(recv.native), c_position, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertAfter is a wrapper around the C function g_node_insert_after.
func (recv *Node) InsertAfter(sibling *Node, node *Node) *Node {
	c_sibling := (*C.GNode)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GNode)(sibling.ToC())
	}

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert_after((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertBefore is a wrapper around the C function g_node_insert_before.
func (recv *Node) InsertBefore(sibling *Node, node *Node) *Node {
	c_sibling := (*C.GNode)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GNode)(sibling.ToC())
	}

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert_before((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsAncestor is a wrapper around the C function g_node_is_ancestor.
func (recv *Node) IsAncestor(descendant *Node) bool {
	c_descendant := (*C.GNode)(C.NULL)
	if descendant != nil {
		c_descendant = (*C.GNode)(descendant.ToC())
	}

	retC := C.g_node_is_ancestor((*C.GNode)(recv.native), c_descendant)
	retGo := retC == C.TRUE

	return retGo
}

// LastChild is a wrapper around the C function g_node_last_child.
func (recv *Node) LastChild() *Node {
	retC := C.g_node_last_child((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LastSibling is a wrapper around the C function g_node_last_sibling.
func (recv *Node) LastSibling() *Node {
	retC := C.g_node_last_sibling((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MaxHeight is a wrapper around the C function g_node_max_height.
func (recv *Node) MaxHeight() uint32 {
	retC := C.g_node_max_height((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NChildren is a wrapper around the C function g_node_n_children.
func (recv *Node) NChildren() uint32 {
	retC := C.g_node_n_children((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NNodes is a wrapper around the C function g_node_n_nodes.
func (recv *Node) NNodes(flags TraverseFlags) uint32 {
	c_flags := (C.GTraverseFlags)(flags)

	retC := C.g_node_n_nodes((*C.GNode)(recv.native), c_flags)
	retGo := (uint32)(retC)

	return retGo
}

// NthChild is a wrapper around the C function g_node_nth_child.
func (recv *Node) NthChild(n uint32) *Node {
	c_n := (C.guint)(n)

	retC := C.g_node_nth_child((*C.GNode)(recv.native), c_n)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prepend is a wrapper around the C function g_node_prepend.
func (recv *Node) Prepend(node *Node) *Node {
	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_prepend((*C.GNode)(recv.native), c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ReverseChildren is a wrapper around the C function g_node_reverse_children.
func (recv *Node) ReverseChildren() {
	C.g_node_reverse_children((*C.GNode)(recv.native))

	return
}

// Unsupported : g_node_traverse : unsupported parameter func : no type generator for NodeTraverseFunc (GNodeTraverseFunc) for param func

// Unlink is a wrapper around the C function g_node_unlink.
func (recv *Node) Unlink() {
	C.g_node_unlink((*C.GNode)(recv.native))

	return
}

// Once is a wrapper around the C record GOnce.
type Once struct {
	native *C.GOnce
	Status OnceStatus
	// retval : no type generator for gpointer, volatile gpointer
}

func OnceNewFromC(u unsafe.Pointer) *Once {
	c := (*C.GOnce)(u)
	if c == nil {
		return nil
	}

	g := &Once{
		Status: (OnceStatus)(c.status),
		native: c,
	}

	return g
}

func (recv *Once) ToC() unsafe.Pointer {
	recv.native.status =
		(C.GOnceStatus)(recv.Status)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Once with another Once, and returns true if they represent the same GObject.
func (recv *Once) Equals(other *Once) bool {
	return other.ToC() == recv.ToC()
}

// g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location
// g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location
// Unsupported : g_once_impl : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// OptionContext is a wrapper around the C record GOptionContext.
type OptionContext struct {
	native *C.GOptionContext
}

func OptionContextNewFromC(u unsafe.Pointer) *OptionContext {
	c := (*C.GOptionContext)(u)
	if c == nil {
		return nil
	}

	g := &OptionContext{native: c}

	return g
}

func (recv *OptionContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionContext with another OptionContext, and returns true if they represent the same GObject.
func (recv *OptionContext) Equals(other *OptionContext) bool {
	return other.ToC() == recv.ToC()
}

// OptionContextNew is a wrapper around the C function g_option_context_new.
func OptionContextNew(parameterString string) *OptionContext {
	c_parameter_string := C.CString(parameterString)
	defer C.free(unsafe.Pointer(c_parameter_string))

	retC := C.g_option_context_new(c_parameter_string)
	retGo := OptionContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddGroup is a wrapper around the C function g_option_context_add_group.
func (recv *OptionContext) AddGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_option_context_add_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// AddMainEntries is a wrapper around the C function g_option_context_add_main_entries.
func (recv *OptionContext) AddMainEntries(entries *OptionEntry, translationDomain string) {
	c_entries := (*C.GOptionEntry)(C.NULL)
	if entries != nil {
		c_entries = (*C.GOptionEntry)(entries.ToC())
	}

	c_translation_domain := C.CString(translationDomain)
	defer C.free(unsafe.Pointer(c_translation_domain))

	C.g_option_context_add_main_entries((*C.GOptionContext)(recv.native), c_entries, c_translation_domain)

	return
}

// Free is a wrapper around the C function g_option_context_free.
func (recv *OptionContext) Free() {
	C.g_option_context_free((*C.GOptionContext)(recv.native))

	return
}

// GetDescription is a wrapper around the C function g_option_context_get_description.
func (recv *OptionContext) GetDescription() string {
	retC := C.g_option_context_get_description((*C.GOptionContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetHelp is a wrapper around the C function g_option_context_get_help.
func (recv *OptionContext) GetHelp(mainHelp bool, group *OptionGroup) string {
	c_main_help :=
		boolToGboolean(mainHelp)

	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	retC := C.g_option_context_get_help((*C.GOptionContext)(recv.native), c_main_help, c_group)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetHelpEnabled is a wrapper around the C function g_option_context_get_help_enabled.
func (recv *OptionContext) GetHelpEnabled() bool {
	retC := C.g_option_context_get_help_enabled((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIgnoreUnknownOptions is a wrapper around the C function g_option_context_get_ignore_unknown_options.
func (recv *OptionContext) GetIgnoreUnknownOptions() bool {
	retC := C.g_option_context_get_ignore_unknown_options((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMainGroup is a wrapper around the C function g_option_context_get_main_group.
func (recv *OptionContext) GetMainGroup() *OptionGroup {
	retC := C.g_option_context_get_main_group((*C.GOptionContext)(recv.native))
	retGo := OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSummary is a wrapper around the C function g_option_context_get_summary.
func (recv *OptionContext) GetSummary() string {
	retC := C.g_option_context_get_summary((*C.GOptionContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Parse is a wrapper around the C function g_option_context_parse.
func (recv *OptionContext) Parse(args []string) (bool, []string, error) {
	cArgc, cArgv := argsIn(args)

	var cThrowableError *C.GError

	retC := C.g_option_context_parse((*C.GOptionContext)(recv.native), &cArgc, &cArgv, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	args = argsOut(cArgc, cArgv)

	return retGo, args, goError
}

// SetDescription is a wrapper around the C function g_option_context_set_description.
func (recv *OptionContext) SetDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_option_context_set_description((*C.GOptionContext)(recv.native), c_description)

	return
}

// SetHelpEnabled is a wrapper around the C function g_option_context_set_help_enabled.
func (recv *OptionContext) SetHelpEnabled(helpEnabled bool) {
	c_help_enabled :=
		boolToGboolean(helpEnabled)

	C.g_option_context_set_help_enabled((*C.GOptionContext)(recv.native), c_help_enabled)

	return
}

// SetIgnoreUnknownOptions is a wrapper around the C function g_option_context_set_ignore_unknown_options.
func (recv *OptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) {
	c_ignore_unknown :=
		boolToGboolean(ignoreUnknown)

	C.g_option_context_set_ignore_unknown_options((*C.GOptionContext)(recv.native), c_ignore_unknown)

	return
}

// SetMainGroup is a wrapper around the C function g_option_context_set_main_group.
func (recv *OptionContext) SetMainGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_option_context_set_main_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// SetSummary is a wrapper around the C function g_option_context_set_summary.
func (recv *OptionContext) SetSummary(summary string) {
	c_summary := C.CString(summary)
	defer C.free(unsafe.Pointer(c_summary))

	C.g_option_context_set_summary((*C.GOptionContext)(recv.native), c_summary)

	return
}

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GTranslateFunc) for param func

// SetTranslationDomain is a wrapper around the C function g_option_context_set_translation_domain.
func (recv *OptionContext) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_context_set_translation_domain((*C.GOptionContext)(recv.native), c_domain)

	return
}

// OptionEntry is a wrapper around the C record GOptionEntry.
type OptionEntry struct {
	native    *C.GOptionEntry
	LongName  string
	ShortName rune
	Flags     int32
	Arg       OptionArg
	// arg_data : no type generator for gpointer, gpointer
	Description    string
	ArgDescription string
}

func OptionEntryNewFromC(u unsafe.Pointer) *OptionEntry {
	c := (*C.GOptionEntry)(u)
	if c == nil {
		return nil
	}

	g := &OptionEntry{
		Arg:            (OptionArg)(c.arg),
		ArgDescription: C.GoString(c.arg_description),
		Description:    C.GoString(c.description),
		Flags:          (int32)(c.flags),
		LongName:       C.GoString(c.long_name),
		ShortName:      (rune)(c.short_name),
		native:         c,
	}

	return g
}

func (recv *OptionEntry) ToC() unsafe.Pointer {
	recv.native.long_name =
		C.CString(recv.LongName)
	recv.native.short_name =
		(C.gchar)(recv.ShortName)
	recv.native.flags =
		(C.gint)(recv.Flags)
	recv.native.arg =
		(C.GOptionArg)(recv.Arg)
	recv.native.description =
		C.CString(recv.Description)
	recv.native.arg_description =
		C.CString(recv.ArgDescription)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionEntry with another OptionEntry, and returns true if they represent the same GObject.
func (recv *OptionEntry) Equals(other *OptionEntry) bool {
	return other.ToC() == recv.ToC()
}

// OptionGroup is a wrapper around the C record GOptionGroup.
type OptionGroup struct {
	native *C.GOptionGroup
}

func OptionGroupNewFromC(u unsafe.Pointer) *OptionGroup {
	c := (*C.GOptionGroup)(u)
	if c == nil {
		return nil
	}

	g := &OptionGroup{native: c}

	return g
}

func (recv *OptionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionGroup with another OptionGroup, and returns true if they represent the same GObject.
func (recv *OptionGroup) Equals(other *OptionGroup) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_option_group_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// AddEntries is a wrapper around the C function g_option_group_add_entries.
func (recv *OptionGroup) AddEntries(entries *OptionEntry) {
	c_entries := (*C.GOptionEntry)(C.NULL)
	if entries != nil {
		c_entries = (*C.GOptionEntry)(entries.ToC())
	}

	C.g_option_group_add_entries((*C.GOptionGroup)(recv.native), c_entries)

	return
}

// Free is a wrapper around the C function g_option_group_free.
func (recv *OptionGroup) Free() {
	C.g_option_group_free((*C.GOptionGroup)(recv.native))

	return
}

// Unsupported : g_option_group_set_error_hook : unsupported parameter error_func : no type generator for OptionErrorFunc (GOptionErrorFunc) for param error_func

// Unsupported : g_option_group_set_parse_hooks : unsupported parameter pre_parse_func : no type generator for OptionParseFunc (GOptionParseFunc) for param pre_parse_func

// Unsupported : g_option_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GTranslateFunc) for param func

// SetTranslationDomain is a wrapper around the C function g_option_group_set_translation_domain.
func (recv *OptionGroup) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_group_set_translation_domain((*C.GOptionGroup)(recv.native), c_domain)

	return
}

// PatternSpec is a wrapper around the C record GPatternSpec.
type PatternSpec struct {
	native *C.GPatternSpec
}

func PatternSpecNewFromC(u unsafe.Pointer) *PatternSpec {
	c := (*C.GPatternSpec)(u)
	if c == nil {
		return nil
	}

	g := &PatternSpec{native: c}

	return g
}

func (recv *PatternSpec) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PatternSpec with another PatternSpec, and returns true if they represent the same GObject.
func (recv *PatternSpec) Equals(other *PatternSpec) bool {
	return other.ToC() == recv.ToC()
}

// PatternSpecNew is a wrapper around the C function g_pattern_spec_new.
func PatternSpecNew(pattern string) *PatternSpec {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	retC := C.g_pattern_spec_new(c_pattern)
	retGo := PatternSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_pattern_spec_equal.
func (recv *PatternSpec) Equal(pspec2 *PatternSpec) bool {
	c_pspec2 := (*C.GPatternSpec)(C.NULL)
	if pspec2 != nil {
		c_pspec2 = (*C.GPatternSpec)(pspec2.ToC())
	}

	retC := C.g_pattern_spec_equal((*C.GPatternSpec)(recv.native), c_pspec2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function g_pattern_spec_free.
func (recv *PatternSpec) Free() {
	C.g_pattern_spec_free((*C.GPatternSpec)(recv.native))

	return
}

// PollFD is a wrapper around the C record GPollFD.
type PollFD struct {
	native  *C.GPollFD
	Fd      int32
	Events  uint32
	Revents uint32
}

func PollFDNewFromC(u unsafe.Pointer) *PollFD {
	c := (*C.GPollFD)(u)
	if c == nil {
		return nil
	}

	g := &PollFD{
		Events:  (uint32)(c.events),
		Fd:      (int32)(c.fd),
		Revents: (uint32)(c.revents),
		native:  c,
	}

	return g
}

func (recv *PollFD) ToC() unsafe.Pointer {
	recv.native.fd =
		(C.gint)(recv.Fd)
	recv.native.events =
		(C.gushort)(recv.Events)
	recv.native.revents =
		(C.gushort)(recv.Revents)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollFD with another PollFD, and returns true if they represent the same GObject.
func (recv *PollFD) Equals(other *PollFD) bool {
	return other.ToC() == recv.ToC()
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native *C.GPrivate
	// Private : p
	// Private : notify
	// Private : future
}

func PrivateNewFromC(u unsafe.Pointer) *Private {
	c := (*C.GPrivate)(u)
	if c == nil {
		return nil
	}

	g := &Private{native: c}

	return g
}

func (recv *Private) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Private with another Private, and returns true if they represent the same GObject.
func (recv *Private) Equals(other *Private) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_private_get : no return generator

// Unsupported : g_private_replace : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Unsupported : g_private_set : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Blacklisted : GPtrArray

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native *C.GQueue
	// head : record
	// tail : record
	Length uint32
}

func QueueNewFromC(u unsafe.Pointer) *Queue {
	c := (*C.GQueue)(u)
	if c == nil {
		return nil
	}

	g := &Queue{
		Length: (uint32)(c.length),
		native: c,
	}

	return g
}

func (recv *Queue) ToC() unsafe.Pointer {
	recv.native.length =
		(C.guint)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Queue with another Queue, and returns true if they represent the same GObject.
func (recv *Queue) Equals(other *Queue) bool {
	return other.ToC() == recv.ToC()
}

// QueueNew is a wrapper around the C function g_queue_new.
func QueueNew() *Queue {
	retC := C.g_queue_new()
	retGo := QueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Clear is a wrapper around the C function g_queue_clear.
func (recv *Queue) Clear() {
	C.g_queue_clear((*C.GQueue)(recv.native))

	return
}

// Copy is a wrapper around the C function g_queue_copy.
func (recv *Queue) Copy() *Queue {
	retC := C.g_queue_copy((*C.GQueue)(recv.native))
	retGo := QueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DeleteLink is a wrapper around the C function g_queue_delete_link.
func (recv *Queue) DeleteLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_delete_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Unsupported : g_queue_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func

// Free is a wrapper around the C function g_queue_free.
func (recv *Queue) Free() {
	C.g_queue_free((*C.GQueue)(recv.native))

	return
}

// Unsupported : g_queue_free_full : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// GetLength is a wrapper around the C function g_queue_get_length.
func (recv *Queue) GetLength() uint32 {
	retC := C.g_queue_get_length((*C.GQueue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_queue_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Init is a wrapper around the C function g_queue_init.
func (recv *Queue) Init() {
	C.g_queue_init((*C.GQueue)(recv.native))

	return
}

// Unsupported : g_queue_insert_after : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_queue_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_queue_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// IsEmpty is a wrapper around the C function g_queue_is_empty.
func (recv *Queue) IsEmpty() bool {
	retC := C.g_queue_is_empty((*C.GQueue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// LinkIndex is a wrapper around the C function g_queue_link_index.
func (recv *Queue) LinkIndex(link *List) int32 {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	retC := C.g_queue_link_index((*C.GQueue)(recv.native), c_link_)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_queue_peek_head : no return generator

// PeekHeadLink is a wrapper around the C function g_queue_peek_head_link.
func (recv *Queue) PeekHeadLink() *List {
	retC := C.g_queue_peek_head_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_peek_nth : no return generator

// PeekNthLink is a wrapper around the C function g_queue_peek_nth_link.
func (recv *Queue) PeekNthLink(n uint32) *List {
	c_n := (C.guint)(n)

	retC := C.g_queue_peek_nth_link((*C.GQueue)(recv.native), c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_peek_tail : no return generator

// PeekTailLink is a wrapper around the C function g_queue_peek_tail_link.
func (recv *Queue) PeekTailLink() *List {
	retC := C.g_queue_peek_tail_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_pop_head : no return generator

// PopHeadLink is a wrapper around the C function g_queue_pop_head_link.
func (recv *Queue) PopHeadLink() *List {
	retC := C.g_queue_pop_head_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_pop_nth : no return generator

// PopNthLink is a wrapper around the C function g_queue_pop_nth_link.
func (recv *Queue) PopNthLink(n uint32) *List {
	c_n := (C.guint)(n)

	retC := C.g_queue_pop_nth_link((*C.GQueue)(recv.native), c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_pop_tail : no return generator

// PopTailLink is a wrapper around the C function g_queue_pop_tail_link.
func (recv *Queue) PopTailLink() *List {
	retC := C.g_queue_pop_tail_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_push_head : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PushHeadLink is a wrapper around the C function g_queue_push_head_link.
func (recv *Queue) PushHeadLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_head_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Unsupported : g_queue_push_nth : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PushNthLink is a wrapper around the C function g_queue_push_nth_link.
func (recv *Queue) PushNthLink(n int32, link *List) {
	c_n := (C.gint)(n)

	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_nth_link((*C.GQueue)(recv.native), c_n, c_link_)

	return
}

// Unsupported : g_queue_push_tail : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PushTailLink is a wrapper around the C function g_queue_push_tail_link.
func (recv *Queue) PushTailLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_tail_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Unsupported : g_queue_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Reverse is a wrapper around the C function g_queue_reverse.
func (recv *Queue) Reverse() {
	C.g_queue_reverse((*C.GQueue)(recv.native))

	return
}

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func

// Unlink is a wrapper around the C function g_queue_unlink.
func (recv *Queue) Unlink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_unlink((*C.GQueue)(recv.native), c_link_)

	return
}

// RWLock is a wrapper around the C record GRWLock.
type RWLock struct {
	native *C.GRWLock
	// Private : p
	// Private : i
}

func RWLockNewFromC(u unsafe.Pointer) *RWLock {
	c := (*C.GRWLock)(u)
	if c == nil {
		return nil
	}

	g := &RWLock{native: c}

	return g
}

func (recv *RWLock) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RWLock with another RWLock, and returns true if they represent the same GObject.
func (recv *RWLock) Equals(other *RWLock) bool {
	return other.ToC() == recv.ToC()
}

// Clear is a wrapper around the C function g_rw_lock_clear.
func (recv *RWLock) Clear() {
	C.g_rw_lock_clear((*C.GRWLock)(recv.native))

	return
}

// Init is a wrapper around the C function g_rw_lock_init.
func (recv *RWLock) Init() {
	C.g_rw_lock_init((*C.GRWLock)(recv.native))

	return
}

// ReaderLock is a wrapper around the C function g_rw_lock_reader_lock.
func (recv *RWLock) ReaderLock() {
	C.g_rw_lock_reader_lock((*C.GRWLock)(recv.native))

	return
}

// ReaderTrylock is a wrapper around the C function g_rw_lock_reader_trylock.
func (recv *RWLock) ReaderTrylock() bool {
	retC := C.g_rw_lock_reader_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ReaderUnlock is a wrapper around the C function g_rw_lock_reader_unlock.
func (recv *RWLock) ReaderUnlock() {
	C.g_rw_lock_reader_unlock((*C.GRWLock)(recv.native))

	return
}

// WriterLock is a wrapper around the C function g_rw_lock_writer_lock.
func (recv *RWLock) WriterLock() {
	C.g_rw_lock_writer_lock((*C.GRWLock)(recv.native))

	return
}

// WriterTrylock is a wrapper around the C function g_rw_lock_writer_trylock.
func (recv *RWLock) WriterTrylock() bool {
	retC := C.g_rw_lock_writer_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// WriterUnlock is a wrapper around the C function g_rw_lock_writer_unlock.
func (recv *RWLock) WriterUnlock() {
	C.g_rw_lock_writer_unlock((*C.GRWLock)(recv.native))

	return
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native *C.GRand
}

func RandNewFromC(u unsafe.Pointer) *Rand {
	c := (*C.GRand)(u)
	if c == nil {
		return nil
	}

	g := &Rand{native: c}

	return g
}

func (recv *Rand) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Rand with another Rand, and returns true if they represent the same GObject.
func (recv *Rand) Equals(other *Rand) bool {
	return other.ToC() == recv.ToC()
}

// RandNew is a wrapper around the C function g_rand_new.
func RandNew() *Rand {
	retC := C.g_rand_new()
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RandNewWithSeed is a wrapper around the C function g_rand_new_with_seed.
func RandNewWithSeed(seed uint32) *Rand {
	c_seed := (C.guint32)(seed)

	retC := C.g_rand_new_with_seed(c_seed)
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RandNewWithSeedArray is a wrapper around the C function g_rand_new_with_seed_array.
func RandNewWithSeedArray(seed uint32, seedLength uint32) *Rand {
	c_seed := (C.guint32)(seed)

	c_seed_length := (C.guint)(seedLength)

	retC := C.g_rand_new_with_seed_array(&c_seed, c_seed_length)
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_rand_copy.
func (recv *Rand) Copy() *Rand {
	retC := C.g_rand_copy((*C.GRand)(recv.native))
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Double is a wrapper around the C function g_rand_double.
func (recv *Rand) Double() float64 {
	retC := C.g_rand_double((*C.GRand)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// DoubleRange is a wrapper around the C function g_rand_double_range.
func (recv *Rand) DoubleRange(begin float64, end float64) float64 {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	retC := C.g_rand_double_range((*C.GRand)(recv.native), c_begin, c_end)
	retGo := (float64)(retC)

	return retGo
}

// Free is a wrapper around the C function g_rand_free.
func (recv *Rand) Free() {
	C.g_rand_free((*C.GRand)(recv.native))

	return
}

// Int is a wrapper around the C function g_rand_int.
func (recv *Rand) Int() uint32 {
	retC := C.g_rand_int((*C.GRand)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IntRange is a wrapper around the C function g_rand_int_range.
func (recv *Rand) IntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_rand_int_range((*C.GRand)(recv.native), c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// SetSeed is a wrapper around the C function g_rand_set_seed.
func (recv *Rand) SetSeed(seed uint32) {
	c_seed := (C.guint32)(seed)

	C.g_rand_set_seed((*C.GRand)(recv.native), c_seed)

	return
}

// SetSeedArray is a wrapper around the C function g_rand_set_seed_array.
func (recv *Rand) SetSeedArray(seed uint32, seedLength uint32) {
	c_seed := (C.guint32)(seed)

	c_seed_length := (C.guint)(seedLength)

	C.g_rand_set_seed_array((*C.GRand)(recv.native), &c_seed, c_seed_length)

	return
}

// RecMutex is a wrapper around the C record GRecMutex.
type RecMutex struct {
	native *C.GRecMutex
	// Private : p
	// Private : i
}

func RecMutexNewFromC(u unsafe.Pointer) *RecMutex {
	c := (*C.GRecMutex)(u)
	if c == nil {
		return nil
	}

	g := &RecMutex{native: c}

	return g
}

func (recv *RecMutex) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RecMutex with another RecMutex, and returns true if they represent the same GObject.
func (recv *RecMutex) Equals(other *RecMutex) bool {
	return other.ToC() == recv.ToC()
}

// Clear is a wrapper around the C function g_rec_mutex_clear.
func (recv *RecMutex) Clear() {
	C.g_rec_mutex_clear((*C.GRecMutex)(recv.native))

	return
}

// Init is a wrapper around the C function g_rec_mutex_init.
func (recv *RecMutex) Init() {
	C.g_rec_mutex_init((*C.GRecMutex)(recv.native))

	return
}

// Lock is a wrapper around the C function g_rec_mutex_lock.
func (recv *RecMutex) Lock() {
	C.g_rec_mutex_lock((*C.GRecMutex)(recv.native))

	return
}

// Trylock is a wrapper around the C function g_rec_mutex_trylock.
func (recv *RecMutex) Trylock() bool {
	retC := C.g_rec_mutex_trylock((*C.GRecMutex)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unlock is a wrapper around the C function g_rec_mutex_unlock.
func (recv *RecMutex) Unlock() {
	C.g_rec_mutex_unlock((*C.GRecMutex)(recv.native))

	return
}

// Blacklisted : GRegex

// SList is a wrapper around the C record GSList.
type SList struct {
	native *C.GSList
	// data : no type generator for gpointer, gpointer
	// next : record
}

func SListNewFromC(u unsafe.Pointer) *SList {
	c := (*C.GSList)(u)
	if c == nil {
		return nil
	}

	g := &SList{native: c}

	return g
}

func (recv *SList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SList with another SList, and returns true if they represent the same GObject.
func (recv *SList) Equals(other *SList) bool {
	return other.ToC() == recv.ToC()
}

// SListAlloc is a wrapper around the C function g_slist_alloc.
func SListAlloc() *SList {
	retC := C.g_slist_alloc()
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// SListConcat is a wrapper around the C function g_slist_concat.
func SListConcat(list1 *SList, list2 *SList) *SList {
	c_list1 := (*C.GSList)(C.NULL)
	if list1 != nil {
		c_list1 = (*C.GSList)(list1.ToC())
	}

	c_list2 := (*C.GSList)(C.NULL)
	if list2 != nil {
		c_list2 = (*C.GSList)(list2.ToC())
	}

	retC := C.g_slist_concat(c_list1, c_list2)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListCopy is a wrapper around the C function g_slist_copy.
func SListCopy(list *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_copy(c_list)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_copy_deep : unsupported parameter func : no type generator for CopyFunc (GCopyFunc) for param func
// SListDeleteLink is a wrapper around the C function g_slist_delete_link.
func SListDeleteLink(list *SList, link *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_link_ := (*C.GSList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GSList)(link.ToC())
	}

	retC := C.g_slist_delete_link(c_list, c_link_)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func
// SListFree is a wrapper around the C function g_slist_free.
func SListFree(list *SList) {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	C.g_slist_free(c_list)

	return
}

// SListFree1 is a wrapper around the C function g_slist_free_1.
func SListFree1(list *SList) {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	C.g_slist_free_1(c_list)

	return
}

// g_slist_free_full : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func
// g_slist_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_insert : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_insert_sorted_with_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// SListLast is a wrapper around the C function g_slist_last.
func SListLast(list *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_last(c_list)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListLength is a wrapper around the C function g_slist_length.
func SListLength(list *SList) uint32 {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_length(c_list)
	retGo := (uint32)(retC)

	return retGo
}

// SListNth is a wrapper around the C function g_slist_nth.
func SListNth(list *SList, n uint32) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_slist_nth(c_list, c_n)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_nth_data : no return generator
// SListPosition is a wrapper around the C function g_slist_position.
func SListPosition(list *SList, llink *SList) int32 {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_llink := (*C.GSList)(C.NULL)
	if llink != nil {
		c_llink = (*C.GSList)(llink.ToC())
	}

	retC := C.g_slist_position(c_list, c_llink)
	retGo := (int32)(retC)

	return retGo
}

// g_slist_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_slist_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// g_slist_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data
// SListRemoveLink is a wrapper around the C function g_slist_remove_link.
func SListRemoveLink(list *SList, link *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_link_ := (*C.GSList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GSList)(link.ToC())
	}

	retC := C.g_slist_remove_link(c_list, c_link_)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListReverse is a wrapper around the C function g_slist_reverse.
func SListReverse(list *SList) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	retC := C.g_slist_reverse(c_list)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_sort : unsupported parameter compare_func : no type generator for CompareFunc (GCompareFunc) for param compare_func
// g_slist_sort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func
// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native *C.GScanner
	// user_data : no type generator for gpointer, gpointer
	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	// qdata : record
	// config : record
	Token TokenType
	// value : no type generator for TokenValue, GTokenValue
	Line      uint32
	Position  uint32
	NextToken TokenType
	// next_value : no type generator for TokenValue, GTokenValue
	NextLine     uint32
	NextPosition uint32
	// Private : symbol_table
	// Private : input_fd
	// Private : text
	// Private : text_end
	// Private : buffer
	// Private : scope_id
	// msg_handler : no type generator for ScannerMsgFunc, GScannerMsgFunc
}

func ScannerNewFromC(u unsafe.Pointer) *Scanner {
	c := (*C.GScanner)(u)
	if c == nil {
		return nil
	}

	g := &Scanner{
		InputName:      C.GoString(c.input_name),
		Line:           (uint32)(c.line),
		MaxParseErrors: (uint32)(c.max_parse_errors),
		NextLine:       (uint32)(c.next_line),
		NextPosition:   (uint32)(c.next_position),
		NextToken:      (TokenType)(c.next_token),
		ParseErrors:    (uint32)(c.parse_errors),
		Position:       (uint32)(c.position),
		Token:          (TokenType)(c.token),
		native:         c,
	}

	return g
}

func (recv *Scanner) ToC() unsafe.Pointer {
	recv.native.max_parse_errors =
		(C.guint)(recv.MaxParseErrors)
	recv.native.parse_errors =
		(C.guint)(recv.ParseErrors)
	recv.native.input_name =
		C.CString(recv.InputName)
	recv.native.token =
		(C.GTokenType)(recv.Token)
	recv.native.line =
		(C.guint)(recv.Line)
	recv.native.position =
		(C.guint)(recv.Position)
	recv.native.next_token =
		(C.GTokenType)(recv.NextToken)
	recv.native.next_line =
		(C.guint)(recv.NextLine)
	recv.native.next_position =
		(C.guint)(recv.NextPosition)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Scanner with another Scanner, and returns true if they represent the same GObject.
func (recv *Scanner) Equals(other *Scanner) bool {
	return other.ToC() == recv.ToC()
}

// ScannerNew is a wrapper around the C function g_scanner_new.
func ScannerNew(configTempl *ScannerConfig) *Scanner {
	c_config_templ := (*C.GScannerConfig)(C.NULL)
	if configTempl != nil {
		c_config_templ = (*C.GScannerConfig)(configTempl.ToC())
	}

	retC := C.g_scanner_new(c_config_templ)
	retGo := ScannerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CurLine is a wrapper around the C function g_scanner_cur_line.
func (recv *Scanner) CurLine() uint32 {
	retC := C.g_scanner_cur_line((*C.GScanner)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// CurPosition is a wrapper around the C function g_scanner_cur_position.
func (recv *Scanner) CurPosition() uint32 {
	retC := C.g_scanner_cur_position((*C.GScanner)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// CurToken is a wrapper around the C function g_scanner_cur_token.
func (recv *Scanner) CurToken() TokenType {
	retC := C.g_scanner_cur_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Unsupported : g_scanner_cur_value : no return generator

// Destroy is a wrapper around the C function g_scanner_destroy.
func (recv *Scanner) Destroy() {
	C.g_scanner_destroy((*C.GScanner)(recv.native))

	return
}

// Eof is a wrapper around the C function g_scanner_eof.
func (recv *Scanner) Eof() bool {
	retC := C.g_scanner_eof((*C.GScanner)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Error is a wrapper around the C function g_scanner_error.
func (recv *Scanner) Error(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_scanner_error((*C.GScanner)(recv.native), c_format)

	return
}

// GetNextToken is a wrapper around the C function g_scanner_get_next_token.
func (recv *Scanner) GetNextToken() TokenType {
	retC := C.g_scanner_get_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// InputFile is a wrapper around the C function g_scanner_input_file.
func (recv *Scanner) InputFile(inputFd int32) {
	c_input_fd := (C.gint)(inputFd)

	C.g_scanner_input_file((*C.GScanner)(recv.native), c_input_fd)

	return
}

// InputText is a wrapper around the C function g_scanner_input_text.
func (recv *Scanner) InputText(text string, textLen uint32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_text_len := (C.guint)(textLen)

	C.g_scanner_input_text((*C.GScanner)(recv.native), c_text, c_text_len)

	return
}

// Unsupported : g_scanner_lookup_symbol : no return generator

// PeekNextToken is a wrapper around the C function g_scanner_peek_next_token.
func (recv *Scanner) PeekNextToken() TokenType {
	retC := C.g_scanner_peek_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Unsupported : g_scanner_scope_add_symbol : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc (GHFunc) for param func

// Unsupported : g_scanner_scope_lookup_symbol : no return generator

// ScopeRemoveSymbol is a wrapper around the C function g_scanner_scope_remove_symbol.
func (recv *Scanner) ScopeRemoveSymbol(scopeId uint32, symbol string) {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	C.g_scanner_scope_remove_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol)

	return
}

// SetScope is a wrapper around the C function g_scanner_set_scope.
func (recv *Scanner) SetScope(scopeId uint32) uint32 {
	c_scope_id := (C.guint)(scopeId)

	retC := C.g_scanner_set_scope((*C.GScanner)(recv.native), c_scope_id)
	retGo := (uint32)(retC)

	return retGo
}

// SyncFileOffset is a wrapper around the C function g_scanner_sync_file_offset.
func (recv *Scanner) SyncFileOffset() {
	C.g_scanner_sync_file_offset((*C.GScanner)(recv.native))

	return
}

// UnexpToken is a wrapper around the C function g_scanner_unexp_token.
func (recv *Scanner) UnexpToken(expectedToken TokenType, identifierSpec string, symbolSpec string, symbolName string, message string, isError int32) {
	c_expected_token := (C.GTokenType)(expectedToken)

	c_identifier_spec := C.CString(identifierSpec)
	defer C.free(unsafe.Pointer(c_identifier_spec))

	c_symbol_spec := C.CString(symbolSpec)
	defer C.free(unsafe.Pointer(c_symbol_spec))

	c_symbol_name := C.CString(symbolName)
	defer C.free(unsafe.Pointer(c_symbol_name))

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	c_is_error := (C.gint)(isError)

	C.g_scanner_unexp_token((*C.GScanner)(recv.native), c_expected_token, c_identifier_spec, c_symbol_spec, c_symbol_name, c_message, c_is_error)

	return
}

// Warn is a wrapper around the C function g_scanner_warn.
func (recv *Scanner) Warn(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_scanner_warn((*C.GScanner)(recv.native), c_format)

	return
}

// ScannerConfig is a wrapper around the C record GScannerConfig.
type ScannerConfig struct {
	native              *C.GScannerConfig
	CsetSkipCharacters  string
	CsetIdentifierFirst string
	CsetIdentifierNth   string
	CpairCommentSingle  string
	// Bitfield not supported :  1 case_sensitive
	// Bitfield not supported :  1 skip_comment_multi
	// Bitfield not supported :  1 skip_comment_single
	// Bitfield not supported :  1 scan_comment_multi
	// Bitfield not supported :  1 scan_identifier
	// Bitfield not supported :  1 scan_identifier_1char
	// Bitfield not supported :  1 scan_identifier_NULL
	// Bitfield not supported :  1 scan_symbols
	// Bitfield not supported :  1 scan_binary
	// Bitfield not supported :  1 scan_octal
	// Bitfield not supported :  1 scan_float
	// Bitfield not supported :  1 scan_hex
	// Bitfield not supported :  1 scan_hex_dollar
	// Bitfield not supported :  1 scan_string_sq
	// Bitfield not supported :  1 scan_string_dq
	// Bitfield not supported :  1 numbers_2_int
	// Bitfield not supported :  1 int_2_float
	// Bitfield not supported :  1 identifier_2_string
	// Bitfield not supported :  1 char_2_token
	// Bitfield not supported :  1 symbol_2_token
	// Bitfield not supported :  1 scope_0_fallback
	// Bitfield not supported :  1 store_int64
	// Private : padding_dummy
}

func ScannerConfigNewFromC(u unsafe.Pointer) *ScannerConfig {
	c := (*C.GScannerConfig)(u)
	if c == nil {
		return nil
	}

	g := &ScannerConfig{
		CpairCommentSingle:  C.GoString(c.cpair_comment_single),
		CsetIdentifierFirst: C.GoString(c.cset_identifier_first),
		CsetIdentifierNth:   C.GoString(c.cset_identifier_nth),
		CsetSkipCharacters:  C.GoString(c.cset_skip_characters),
		native:              c,
	}

	return g
}

func (recv *ScannerConfig) ToC() unsafe.Pointer {
	recv.native.cset_skip_characters =
		C.CString(recv.CsetSkipCharacters)
	recv.native.cset_identifier_first =
		C.CString(recv.CsetIdentifierFirst)
	recv.native.cset_identifier_nth =
		C.CString(recv.CsetIdentifierNth)
	recv.native.cpair_comment_single =
		C.CString(recv.CpairCommentSingle)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ScannerConfig with another ScannerConfig, and returns true if they represent the same GObject.
func (recv *ScannerConfig) Equals(other *ScannerConfig) bool {
	return other.ToC() == recv.ToC()
}

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native *C.GSequence
}

func SequenceNewFromC(u unsafe.Pointer) *Sequence {
	c := (*C.GSequence)(u)
	if c == nil {
		return nil
	}

	g := &Sequence{native: c}

	return g
}

func (recv *Sequence) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Sequence with another Sequence, and returns true if they represent the same GObject.
func (recv *Sequence) Equals(other *Sequence) bool {
	return other.ToC() == recv.ToC()
}

// g_sequence_foreach_range : unsupported parameter func : no type generator for Func (GFunc) for param func
// g_sequence_get : no return generator
// g_sequence_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// SequenceMove is a wrapper around the C function g_sequence_move.
func SequenceMove(src *SequenceIter, dest *SequenceIter) {
	c_src := (*C.GSequenceIter)(C.NULL)
	if src != nil {
		c_src = (*C.GSequenceIter)(src.ToC())
	}

	c_dest := (*C.GSequenceIter)(C.NULL)
	if dest != nil {
		c_dest = (*C.GSequenceIter)(dest.ToC())
	}

	C.g_sequence_move(c_src, c_dest)

	return
}

// SequenceMoveRange is a wrapper around the C function g_sequence_move_range.
func SequenceMoveRange(dest *SequenceIter, begin *SequenceIter, end *SequenceIter) {
	c_dest := (*C.GSequenceIter)(C.NULL)
	if dest != nil {
		c_dest = (*C.GSequenceIter)(dest.ToC())
	}

	c_begin := (*C.GSequenceIter)(C.NULL)
	if begin != nil {
		c_begin = (*C.GSequenceIter)(begin.ToC())
	}

	c_end := (*C.GSequenceIter)(C.NULL)
	if end != nil {
		c_end = (*C.GSequenceIter)(end.ToC())
	}

	C.g_sequence_move_range(c_dest, c_begin, c_end)

	return
}

// g_sequence_new : unsupported parameter data_destroy : no type generator for DestroyNotify (GDestroyNotify) for param data_destroy
// SequenceRangeGetMidpoint is a wrapper around the C function g_sequence_range_get_midpoint.
func SequenceRangeGetMidpoint(begin *SequenceIter, end *SequenceIter) *SequenceIter {
	c_begin := (*C.GSequenceIter)(C.NULL)
	if begin != nil {
		c_begin = (*C.GSequenceIter)(begin.ToC())
	}

	c_end := (*C.GSequenceIter)(C.NULL)
	if end != nil {
		c_end = (*C.GSequenceIter)(end.ToC())
	}

	retC := C.g_sequence_range_get_midpoint(c_begin, c_end)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SequenceRemove is a wrapper around the C function g_sequence_remove.
func SequenceRemove(iter *SequenceIter) {
	c_iter := (*C.GSequenceIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GSequenceIter)(iter.ToC())
	}

	C.g_sequence_remove(c_iter)

	return
}

// SequenceRemoveRange is a wrapper around the C function g_sequence_remove_range.
func SequenceRemoveRange(begin *SequenceIter, end *SequenceIter) {
	c_begin := (*C.GSequenceIter)(C.NULL)
	if begin != nil {
		c_begin = (*C.GSequenceIter)(begin.ToC())
	}

	c_end := (*C.GSequenceIter)(C.NULL)
	if end != nil {
		c_end = (*C.GSequenceIter)(end.ToC())
	}

	C.g_sequence_remove_range(c_begin, c_end)

	return
}

// g_sequence_set : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_sequence_sort_changed : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func
// g_sequence_sort_changed_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param iter_cmp
// SequenceSwap is a wrapper around the C function g_sequence_swap.
func SequenceSwap(a *SequenceIter, b *SequenceIter) {
	c_a := (*C.GSequenceIter)(C.NULL)
	if a != nil {
		c_a = (*C.GSequenceIter)(a.ToC())
	}

	c_b := (*C.GSequenceIter)(C.NULL)
	if b != nil {
		c_b = (*C.GSequenceIter)(b.ToC())
	}

	C.g_sequence_swap(c_a, c_b)

	return
}

// Unsupported : g_sequence_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func

// Free is a wrapper around the C function g_sequence_free.
func (recv *Sequence) Free() {
	C.g_sequence_free((*C.GSequence)(recv.native))

	return
}

// GetBeginIter is a wrapper around the C function g_sequence_get_begin_iter.
func (recv *Sequence) GetBeginIter() *SequenceIter {
	retC := C.g_sequence_get_begin_iter((*C.GSequence)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEndIter is a wrapper around the C function g_sequence_get_end_iter.
func (recv *Sequence) GetEndIter() *SequenceIter {
	retC := C.g_sequence_get_end_iter((*C.GSequence)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIterAtPos is a wrapper around the C function g_sequence_get_iter_at_pos.
func (recv *Sequence) GetIterAtPos(pos int32) *SequenceIter {
	c_pos := (C.gint)(pos)

	retC := C.g_sequence_get_iter_at_pos((*C.GSequence)(recv.native), c_pos)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLength is a wrapper around the C function g_sequence_get_length.
func (recv *Sequence) GetLength() int32 {
	retC := C.g_sequence_get_length((*C.GSequence)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_sequence_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_lookup : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_lookup_iter : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_search : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_search_iter : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param cmp_func

// SequenceIter is a wrapper around the C record GSequenceIter.
type SequenceIter struct {
	native *C.GSequenceIter
}

func SequenceIterNewFromC(u unsafe.Pointer) *SequenceIter {
	c := (*C.GSequenceIter)(u)
	if c == nil {
		return nil
	}

	g := &SequenceIter{native: c}

	return g
}

func (recv *SequenceIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SequenceIter with another SequenceIter, and returns true if they represent the same GObject.
func (recv *SequenceIter) Equals(other *SequenceIter) bool {
	return other.ToC() == recv.ToC()
}

// Compare is a wrapper around the C function g_sequence_iter_compare.
func (recv *SequenceIter) Compare(b *SequenceIter) int32 {
	c_b := (*C.GSequenceIter)(C.NULL)
	if b != nil {
		c_b = (*C.GSequenceIter)(b.ToC())
	}

	retC := C.g_sequence_iter_compare((*C.GSequenceIter)(recv.native), c_b)
	retGo := (int32)(retC)

	return retGo
}

// GetPosition is a wrapper around the C function g_sequence_iter_get_position.
func (recv *SequenceIter) GetPosition() int32 {
	retC := C.g_sequence_iter_get_position((*C.GSequenceIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSequence is a wrapper around the C function g_sequence_iter_get_sequence.
func (recv *SequenceIter) GetSequence() *Sequence {
	retC := C.g_sequence_iter_get_sequence((*C.GSequenceIter)(recv.native))
	retGo := SequenceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsBegin is a wrapper around the C function g_sequence_iter_is_begin.
func (recv *SequenceIter) IsBegin() bool {
	retC := C.g_sequence_iter_is_begin((*C.GSequenceIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsEnd is a wrapper around the C function g_sequence_iter_is_end.
func (recv *SequenceIter) IsEnd() bool {
	retC := C.g_sequence_iter_is_end((*C.GSequenceIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Move is a wrapper around the C function g_sequence_iter_move.
func (recv *SequenceIter) Move(delta int32) *SequenceIter {
	c_delta := (C.gint)(delta)

	retC := C.g_sequence_iter_move((*C.GSequenceIter)(recv.native), c_delta)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Next is a wrapper around the C function g_sequence_iter_next.
func (recv *SequenceIter) Next() *SequenceIter {
	retC := C.g_sequence_iter_next((*C.GSequenceIter)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prev is a wrapper around the C function g_sequence_iter_prev.
func (recv *SequenceIter) Prev() *SequenceIter {
	retC := C.g_sequence_iter_prev((*C.GSequenceIter)(recv.native))
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Source is a wrapper around the C record GSource.
type Source struct {
	native *C.GSource
	// Private : callback_data
	// Private : callback_funcs
	// Private : source_funcs
	// Private : ref_count
	// Private : context
	// Private : priority
	// Private : flags
	// Private : source_id
	// Private : poll_fds
	// Private : prev
	// Private : next
	// Private : name
	// Private : priv
}

func SourceNewFromC(u unsafe.Pointer) *Source {
	c := (*C.GSource)(u)
	if c == nil {
		return nil
	}

	g := &Source{native: c}

	return g
}

func (recv *Source) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Source with another Source, and returns true if they represent the same GObject.
func (recv *Source) Equals(other *Source) bool {
	return other.ToC() == recv.ToC()
}

// SourceNew is a wrapper around the C function g_source_new.
func SourceNew(sourceFuncs *SourceFuncs, structSize uint32) *Source {
	c_source_funcs := (*C.GSourceFuncs)(C.NULL)
	if sourceFuncs != nil {
		c_source_funcs = (*C.GSourceFuncs)(sourceFuncs.ToC())
	}

	c_struct_size := (C.guint)(structSize)

	retC := C.g_source_new(c_source_funcs, c_struct_size)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SourceRemove is a wrapper around the C function g_source_remove.
func SourceRemove(tag uint32) bool {
	c_tag := (C.guint)(tag)

	retC := C.g_source_remove(c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// g_source_remove_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data
// g_source_remove_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data
// SourceSetNameById is a wrapper around the C function g_source_set_name_by_id.
func SourceSetNameById(tag uint32, name string) {
	c_tag := (C.guint)(tag)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_source_set_name_by_id(c_tag, c_name)

	return
}

// AddChildSource is a wrapper around the C function g_source_add_child_source.
func (recv *Source) AddChildSource(childSource *Source) {
	c_child_source := (*C.GSource)(C.NULL)
	if childSource != nil {
		c_child_source = (*C.GSource)(childSource.ToC())
	}

	C.g_source_add_child_source((*C.GSource)(recv.native), c_child_source)

	return
}

// AddPoll is a wrapper around the C function g_source_add_poll.
func (recv *Source) AddPoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_source_add_poll((*C.GSource)(recv.native), c_fd)

	return
}

// Attach is a wrapper around the C function g_source_attach.
func (recv *Source) Attach(context *MainContext) uint32 {
	c_context := (*C.GMainContext)(C.NULL)
	if context != nil {
		c_context = (*C.GMainContext)(context.ToC())
	}

	retC := C.g_source_attach((*C.GSource)(recv.native), c_context)
	retGo := (uint32)(retC)

	return retGo
}

// Destroy is a wrapper around the C function g_source_destroy.
func (recv *Source) Destroy() {
	C.g_source_destroy((*C.GSource)(recv.native))

	return
}

// GetCanRecurse is a wrapper around the C function g_source_get_can_recurse.
func (recv *Source) GetCanRecurse() bool {
	retC := C.g_source_get_can_recurse((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetContext is a wrapper around the C function g_source_get_context.
func (recv *Source) GetContext() *MainContext {
	retC := C.g_source_get_context((*C.GSource)(recv.native))
	var retGo (*MainContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MainContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetCurrentTime is a wrapper around the C function g_source_get_current_time.
func (recv *Source) GetCurrentTime(timeval *TimeVal) {
	c_timeval := (*C.GTimeVal)(C.NULL)
	if timeval != nil {
		c_timeval = (*C.GTimeVal)(timeval.ToC())
	}

	C.g_source_get_current_time((*C.GSource)(recv.native), c_timeval)

	return
}

// GetId is a wrapper around the C function g_source_get_id.
func (recv *Source) GetId() uint32 {
	retC := C.g_source_get_id((*C.GSource)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetName is a wrapper around the C function g_source_get_name.
func (recv *Source) GetName() string {
	retC := C.g_source_get_name((*C.GSource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPriority is a wrapper around the C function g_source_get_priority.
func (recv *Source) GetPriority() int32 {
	retC := C.g_source_get_priority((*C.GSource)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetReadyTime is a wrapper around the C function g_source_get_ready_time.
func (recv *Source) GetReadyTime() int64 {
	retC := C.g_source_get_ready_time((*C.GSource)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetTime is a wrapper around the C function g_source_get_time.
func (recv *Source) GetTime() int64 {
	retC := C.g_source_get_time((*C.GSource)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// IsDestroyed is a wrapper around the C function g_source_is_destroyed.
func (recv *Source) IsDestroyed() bool {
	retC := C.g_source_is_destroyed((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_source_ref.
func (recv *Source) Ref() *Source {
	retC := C.g_source_ref((*C.GSource)(recv.native))
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveChildSource is a wrapper around the C function g_source_remove_child_source.
func (recv *Source) RemoveChildSource(childSource *Source) {
	c_child_source := (*C.GSource)(C.NULL)
	if childSource != nil {
		c_child_source = (*C.GSource)(childSource.ToC())
	}

	C.g_source_remove_child_source((*C.GSource)(recv.native), c_child_source)

	return
}

// RemovePoll is a wrapper around the C function g_source_remove_poll.
func (recv *Source) RemovePoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_source_remove_poll((*C.GSource)(recv.native), c_fd)

	return
}

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc (GSourceFunc) for param func

// Unsupported : g_source_set_callback_indirect : unsupported parameter callback_data : no type generator for gpointer (gpointer) for param callback_data

// SetCanRecurse is a wrapper around the C function g_source_set_can_recurse.
func (recv *Source) SetCanRecurse(canRecurse bool) {
	c_can_recurse :=
		boolToGboolean(canRecurse)

	C.g_source_set_can_recurse((*C.GSource)(recv.native), c_can_recurse)

	return
}

// SetFuncs is a wrapper around the C function g_source_set_funcs.
func (recv *Source) SetFuncs(funcs *SourceFuncs) {
	c_funcs := (*C.GSourceFuncs)(C.NULL)
	if funcs != nil {
		c_funcs = (*C.GSourceFuncs)(funcs.ToC())
	}

	C.g_source_set_funcs((*C.GSource)(recv.native), c_funcs)

	return
}

// SetName is a wrapper around the C function g_source_set_name.
func (recv *Source) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_source_set_name((*C.GSource)(recv.native), c_name)

	return
}

// SetPriority is a wrapper around the C function g_source_set_priority.
func (recv *Source) SetPriority(priority int32) {
	c_priority := (C.gint)(priority)

	C.g_source_set_priority((*C.GSource)(recv.native), c_priority)

	return
}

// Unref is a wrapper around the C function g_source_unref.
func (recv *Source) Unref() {
	C.g_source_unref((*C.GSource)(recv.native))

	return
}

// SourceCallbackFuncs is a wrapper around the C record GSourceCallbackFuncs.
type SourceCallbackFuncs struct {
	native *C.GSourceCallbackFuncs
	// no type for ref
	// no type for unref
	// no type for get
}

func SourceCallbackFuncsNewFromC(u unsafe.Pointer) *SourceCallbackFuncs {
	c := (*C.GSourceCallbackFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceCallbackFuncs{native: c}

	return g
}

func (recv *SourceCallbackFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SourceCallbackFuncs with another SourceCallbackFuncs, and returns true if they represent the same GObject.
func (recv *SourceCallbackFuncs) Equals(other *SourceCallbackFuncs) bool {
	return other.ToC() == recv.ToC()
}

// SourceFuncs is a wrapper around the C record GSourceFuncs.
type SourceFuncs struct {
	native *C.GSourceFuncs
	// no type for prepare
	// no type for check
	// no type for dispatch
	// no type for finalize
	// Private : closure_callback
	// Private : closure_marshal
}

func SourceFuncsNewFromC(u unsafe.Pointer) *SourceFuncs {
	c := (*C.GSourceFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceFuncs{native: c}

	return g
}

func (recv *SourceFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SourceFuncs with another SourceFuncs, and returns true if they represent the same GObject.
func (recv *SourceFuncs) Equals(other *SourceFuncs) bool {
	return other.ToC() == recv.ToC()
}

// SourcePrivate is a wrapper around the C record GSourcePrivate.
type SourcePrivate struct {
	native *C.GSourcePrivate
}

func SourcePrivateNewFromC(u unsafe.Pointer) *SourcePrivate {
	c := (*C.GSourcePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SourcePrivate{native: c}

	return g
}

func (recv *SourcePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SourcePrivate with another SourcePrivate, and returns true if they represent the same GObject.
func (recv *SourcePrivate) Equals(other *SourcePrivate) bool {
	return other.ToC() == recv.ToC()
}

// StatBuf is a wrapper around the C record GStatBuf.
type StatBuf struct {
	native *C.GStatBuf
}

func StatBufNewFromC(u unsafe.Pointer) *StatBuf {
	c := (*C.GStatBuf)(u)
	if c == nil {
		return nil
	}

	g := &StatBuf{native: c}

	return g
}

func (recv *StatBuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StatBuf with another StatBuf, and returns true if they represent the same GObject.
func (recv *StatBuf) Equals(other *StatBuf) bool {
	return other.ToC() == recv.ToC()
}

// String is a wrapper around the C record GString.
type String struct {
	native       *C.GString
	Str          string
	Len          uint64
	AllocatedLen uint64
}

func StringNewFromC(u unsafe.Pointer) *String {
	c := (*C.GString)(u)
	if c == nil {
		return nil
	}

	g := &String{
		AllocatedLen: (uint64)(c.allocated_len),
		Len:          (uint64)(c.len),
		Str:          C.GoString(c.str),
		native:       c,
	}

	return g
}

func (recv *String) ToC() unsafe.Pointer {
	recv.native.str =
		C.CString(recv.Str)
	recv.native.len =
		(C.gsize)(recv.Len)
	recv.native.allocated_len =
		(C.gsize)(recv.AllocatedLen)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this String with another String, and returns true if they represent the same GObject.
func (recv *String) Equals(other *String) bool {
	return other.ToC() == recv.ToC()
}

// Append is a wrapper around the C function g_string_append.
func (recv *String) Append(val string) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_append((*C.GString)(recv.native), c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendC is a wrapper around the C function g_string_append_c.
func (recv *String) AppendC(c rune) *String {
	c_c := (C.gchar)(c)

	retC := C.g_string_append_c((*C.GString)(recv.native), c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendLen is a wrapper around the C function g_string_append_len.
func (recv *String) AppendLen(val string, len int64) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_append_len((*C.GString)(recv.native), c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendPrintf is a wrapper around the C function g_string_append_printf.
func (recv *String) AppendPrintf(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_string_append_printf((*C.GString)(recv.native), c_format)

	return
}

// AppendUnichar is a wrapper around the C function g_string_append_unichar.
func (recv *String) AppendUnichar(wc rune) *String {
	c_wc := (C.gunichar)(wc)

	retC := C.g_string_append_unichar((*C.GString)(recv.native), c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendUriEscaped is a wrapper around the C function g_string_append_uri_escaped.
func (recv *String) AppendUriEscaped(unescaped string, reservedCharsAllowed string, allowUtf8 bool) *String {
	c_unescaped := C.CString(unescaped)
	defer C.free(unsafe.Pointer(c_unescaped))

	c_reserved_chars_allowed := C.CString(reservedCharsAllowed)
	defer C.free(unsafe.Pointer(c_reserved_chars_allowed))

	c_allow_utf8 :=
		boolToGboolean(allowUtf8)

	retC := C.g_string_append_uri_escaped((*C.GString)(recv.native), c_unescaped, c_reserved_chars_allowed, c_allow_utf8)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// AsciiDown is a wrapper around the C function g_string_ascii_down.
func (recv *String) AsciiDown() *String {
	retC := C.g_string_ascii_down((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AsciiUp is a wrapper around the C function g_string_ascii_up.
func (recv *String) AsciiUp() *String {
	retC := C.g_string_ascii_up((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Assign is a wrapper around the C function g_string_assign.
func (recv *String) Assign(rval string) *String {
	c_rval := C.CString(rval)
	defer C.free(unsafe.Pointer(c_rval))

	retC := C.g_string_assign((*C.GString)(recv.native), c_rval)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Down is a wrapper around the C function g_string_down.
func (recv *String) Down() *String {
	retC := C.g_string_down((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_string_equal.
func (recv *String) Equal(v2 *String) bool {
	c_v2 := (*C.GString)(C.NULL)
	if v2 != nil {
		c_v2 = (*C.GString)(v2.ToC())
	}

	retC := C.g_string_equal((*C.GString)(recv.native), c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Erase is a wrapper around the C function g_string_erase.
func (recv *String) Erase(pos int64, len int64) *String {
	c_pos := (C.gssize)(pos)

	c_len := (C.gssize)(len)

	retC := C.g_string_erase((*C.GString)(recv.native), c_pos, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_string_free.
func (recv *String) Free(freeSegment bool) string {
	c_free_segment :=
		boolToGboolean(freeSegment)

	retC := C.g_string_free((*C.GString)(recv.native), c_free_segment)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FreeToBytes is a wrapper around the C function g_string_free_to_bytes.
func (recv *String) FreeToBytes() *Bytes {
	retC := C.g_string_free_to_bytes((*C.GString)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Hash is a wrapper around the C function g_string_hash.
func (recv *String) Hash() uint32 {
	retC := C.g_string_hash((*C.GString)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Insert is a wrapper around the C function g_string_insert.
func (recv *String) Insert(pos int64, val string) *String {
	c_pos := (C.gssize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_insert((*C.GString)(recv.native), c_pos, c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertC is a wrapper around the C function g_string_insert_c.
func (recv *String) InsertC(pos int64, c rune) *String {
	c_pos := (C.gssize)(pos)

	c_c := (C.gchar)(c)

	retC := C.g_string_insert_c((*C.GString)(recv.native), c_pos, c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertLen is a wrapper around the C function g_string_insert_len.
func (recv *String) InsertLen(pos int64, val string, len int64) *String {
	c_pos := (C.gssize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_insert_len((*C.GString)(recv.native), c_pos, c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertUnichar is a wrapper around the C function g_string_insert_unichar.
func (recv *String) InsertUnichar(pos int64, wc rune) *String {
	c_pos := (C.gssize)(pos)

	c_wc := (C.gunichar)(wc)

	retC := C.g_string_insert_unichar((*C.GString)(recv.native), c_pos, c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Overwrite is a wrapper around the C function g_string_overwrite.
func (recv *String) Overwrite(pos uint64, val string) *String {
	c_pos := (C.gsize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_overwrite((*C.GString)(recv.native), c_pos, c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OverwriteLen is a wrapper around the C function g_string_overwrite_len.
func (recv *String) OverwriteLen(pos uint64, val string, len int64) *String {
	c_pos := (C.gsize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_overwrite_len((*C.GString)(recv.native), c_pos, c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prepend is a wrapper around the C function g_string_prepend.
func (recv *String) Prepend(val string) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_prepend((*C.GString)(recv.native), c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrependC is a wrapper around the C function g_string_prepend_c.
func (recv *String) PrependC(c rune) *String {
	c_c := (C.gchar)(c)

	retC := C.g_string_prepend_c((*C.GString)(recv.native), c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrependLen is a wrapper around the C function g_string_prepend_len.
func (recv *String) PrependLen(val string, len int64) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_prepend_len((*C.GString)(recv.native), c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PrependUnichar is a wrapper around the C function g_string_prepend_unichar.
func (recv *String) PrependUnichar(wc rune) *String {
	c_wc := (C.gunichar)(wc)

	retC := C.g_string_prepend_unichar((*C.GString)(recv.native), c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Printf is a wrapper around the C function g_string_printf.
func (recv *String) Printf(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_string_printf((*C.GString)(recv.native), c_format)

	return
}

// SetSize is a wrapper around the C function g_string_set_size.
func (recv *String) SetSize(len uint64) *String {
	c_len := (C.gsize)(len)

	retC := C.g_string_set_size((*C.GString)(recv.native), c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Truncate is a wrapper around the C function g_string_truncate.
func (recv *String) Truncate(len uint64) *String {
	c_len := (C.gsize)(len)

	retC := C.g_string_truncate((*C.GString)(recv.native), c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Up is a wrapper around the C function g_string_up.
func (recv *String) Up() *String {
	retC := C.g_string_up((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// StringChunk is a wrapper around the C record GStringChunk.
type StringChunk struct {
	native *C.GStringChunk
}

func StringChunkNewFromC(u unsafe.Pointer) *StringChunk {
	c := (*C.GStringChunk)(u)
	if c == nil {
		return nil
	}

	g := &StringChunk{native: c}

	return g
}

func (recv *StringChunk) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StringChunk with another StringChunk, and returns true if they represent the same GObject.
func (recv *StringChunk) Equals(other *StringChunk) bool {
	return other.ToC() == recv.ToC()
}

// StringChunkNew is a wrapper around the C function g_string_chunk_new.
func StringChunkNew(size uint64) *StringChunk {
	c_size := (C.gsize)(size)

	retC := C.g_string_chunk_new(c_size)
	retGo := StringChunkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Clear is a wrapper around the C function g_string_chunk_clear.
func (recv *StringChunk) Clear() {
	C.g_string_chunk_clear((*C.GStringChunk)(recv.native))

	return
}

// Free is a wrapper around the C function g_string_chunk_free.
func (recv *StringChunk) Free() {
	C.g_string_chunk_free((*C.GStringChunk)(recv.native))

	return
}

// Insert is a wrapper around the C function g_string_chunk_insert.
func (recv *StringChunk) Insert(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_string_chunk_insert((*C.GStringChunk)(recv.native), c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// InsertConst is a wrapper around the C function g_string_chunk_insert_const.
func (recv *StringChunk) InsertConst(string_ string) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_string_chunk_insert_const((*C.GStringChunk)(recv.native), c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// InsertLen is a wrapper around the C function g_string_chunk_insert_len.
func (recv *StringChunk) InsertLen(string_ string, len int64) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_len := (C.gssize)(len)

	retC := C.g_string_chunk_insert_len((*C.GStringChunk)(recv.native), c_string, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// TestCase is a wrapper around the C record GTestCase.
type TestCase struct {
	native *C.GTestCase
}

func TestCaseNewFromC(u unsafe.Pointer) *TestCase {
	c := (*C.GTestCase)(u)
	if c == nil {
		return nil
	}

	g := &TestCase{native: c}

	return g
}

func (recv *TestCase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestCase with another TestCase, and returns true if they represent the same GObject.
func (recv *TestCase) Equals(other *TestCase) bool {
	return other.ToC() == recv.ToC()
}

// TestConfig is a wrapper around the C record GTestConfig.
type TestConfig struct {
	native          *C.GTestConfig
	TestInitialized bool
	TestQuick       bool
	TestPerf        bool
	TestVerbose     bool
	TestQuiet       bool
	TestUndefined   bool
}

func TestConfigNewFromC(u unsafe.Pointer) *TestConfig {
	c := (*C.GTestConfig)(u)
	if c == nil {
		return nil
	}

	g := &TestConfig{
		TestInitialized: c.test_initialized == C.TRUE,
		TestPerf:        c.test_perf == C.TRUE,
		TestQuick:       c.test_quick == C.TRUE,
		TestQuiet:       c.test_quiet == C.TRUE,
		TestUndefined:   c.test_undefined == C.TRUE,
		TestVerbose:     c.test_verbose == C.TRUE,
		native:          c,
	}

	return g
}

func (recv *TestConfig) ToC() unsafe.Pointer {
	recv.native.test_initialized =
		boolToGboolean(recv.TestInitialized)
	recv.native.test_quick =
		boolToGboolean(recv.TestQuick)
	recv.native.test_perf =
		boolToGboolean(recv.TestPerf)
	recv.native.test_verbose =
		boolToGboolean(recv.TestVerbose)
	recv.native.test_quiet =
		boolToGboolean(recv.TestQuiet)
	recv.native.test_undefined =
		boolToGboolean(recv.TestUndefined)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestConfig with another TestConfig, and returns true if they represent the same GObject.
func (recv *TestConfig) Equals(other *TestConfig) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// TestSuite is a wrapper around the C record GTestSuite.
type TestSuite struct {
	native *C.GTestSuite
}

func TestSuiteNewFromC(u unsafe.Pointer) *TestSuite {
	c := (*C.GTestSuite)(u)
	if c == nil {
		return nil
	}

	g := &TestSuite{native: c}

	return g
}

func (recv *TestSuite) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TestSuite with another TestSuite, and returns true if they represent the same GObject.
func (recv *TestSuite) Equals(other *TestSuite) bool {
	return other.ToC() == recv.ToC()
}

// Add is a wrapper around the C function g_test_suite_add.
func (recv *TestSuite) Add(testCase *TestCase) {
	c_test_case := (*C.GTestCase)(C.NULL)
	if testCase != nil {
		c_test_case = (*C.GTestCase)(testCase.ToC())
	}

	C.g_test_suite_add((*C.GTestSuite)(recv.native), c_test_case)

	return
}

// AddSuite is a wrapper around the C function g_test_suite_add_suite.
func (recv *TestSuite) AddSuite(nestedsuite *TestSuite) {
	c_nestedsuite := (*C.GTestSuite)(C.NULL)
	if nestedsuite != nil {
		c_nestedsuite = (*C.GTestSuite)(nestedsuite.ToC())
	}

	C.g_test_suite_add_suite((*C.GTestSuite)(recv.native), c_nestedsuite)

	return
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native *C.GThread
}

func ThreadNewFromC(u unsafe.Pointer) *Thread {
	c := (*C.GThread)(u)
	if c == nil {
		return nil
	}

	g := &Thread{native: c}

	return g
}

func (recv *Thread) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Thread with another Thread, and returns true if they represent the same GObject.
func (recv *Thread) Equals(other *Thread) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// ThreadErrorQuark is a wrapper around the C function g_thread_error_quark.
func ThreadErrorQuark() Quark {
	retC := C.g_thread_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// g_thread_exit : unsupported parameter retval : no type generator for gpointer (gpointer) for param retval
// ThreadSelf is a wrapper around the C function g_thread_self.
func ThreadSelf() *Thread {
	retC := C.g_thread_self()
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ThreadYield is a wrapper around the C function g_thread_yield.
func ThreadYield() {
	C.g_thread_yield()

	return
}

// Unsupported : g_thread_join : no return generator

// Ref is a wrapper around the C function g_thread_ref.
func (recv *Thread) Ref() *Thread {
	retC := C.g_thread_ref((*C.GThread)(recv.native))
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_thread_unref.
func (recv *Thread) Unref() {
	C.g_thread_unref((*C.GThread)(recv.native))

	return
}

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	// user_data : no type generator for gpointer, gpointer
	Exclusive bool
}

func ThreadPoolNewFromC(u unsafe.Pointer) *ThreadPool {
	c := (*C.GThreadPool)(u)
	if c == nil {
		return nil
	}

	g := &ThreadPool{
		Exclusive: c.exclusive == C.TRUE,
		native:    c,
	}

	return g
}

func (recv *ThreadPool) ToC() unsafe.Pointer {
	recv.native.exclusive =
		boolToGboolean(recv.Exclusive)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThreadPool with another ThreadPool, and returns true if they represent the same GObject.
func (recv *ThreadPool) Equals(other *ThreadPool) bool {
	return other.ToC() == recv.ToC()
}

// ThreadPoolGetMaxIdleTime is a wrapper around the C function g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() uint32 {
	retC := C.g_thread_pool_get_max_idle_time()
	retGo := (uint32)(retC)

	return retGo
}

// ThreadPoolGetMaxUnusedThreads is a wrapper around the C function g_thread_pool_get_max_unused_threads.
func ThreadPoolGetMaxUnusedThreads() int32 {
	retC := C.g_thread_pool_get_max_unused_threads()
	retGo := (int32)(retC)

	return retGo
}

// ThreadPoolGetNumUnusedThreads is a wrapper around the C function g_thread_pool_get_num_unused_threads.
func ThreadPoolGetNumUnusedThreads() uint32 {
	retC := C.g_thread_pool_get_num_unused_threads()
	retGo := (uint32)(retC)

	return retGo
}

// g_thread_pool_new : unsupported parameter func : no type generator for Func (GFunc) for param func
// ThreadPoolSetMaxIdleTime is a wrapper around the C function g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {
	c_interval := (C.guint)(interval)

	C.g_thread_pool_set_max_idle_time(c_interval)

	return
}

// ThreadPoolSetMaxUnusedThreads is a wrapper around the C function g_thread_pool_set_max_unused_threads.
func ThreadPoolSetMaxUnusedThreads(maxThreads int32) {
	c_max_threads := (C.gint)(maxThreads)

	C.g_thread_pool_set_max_unused_threads(c_max_threads)

	return
}

// ThreadPoolStopUnusedThreads is a wrapper around the C function g_thread_pool_stop_unused_threads.
func ThreadPoolStopUnusedThreads() {
	C.g_thread_pool_stop_unused_threads()

	return
}

// Free is a wrapper around the C function g_thread_pool_free.
func (recv *ThreadPool) Free(immediate bool, wait bool) {
	c_immediate :=
		boolToGboolean(immediate)

	c_wait_ :=
		boolToGboolean(wait)

	C.g_thread_pool_free((*C.GThreadPool)(recv.native), c_immediate, c_wait_)

	return
}

// GetMaxThreads is a wrapper around the C function g_thread_pool_get_max_threads.
func (recv *ThreadPool) GetMaxThreads() int32 {
	retC := C.g_thread_pool_get_max_threads((*C.GThreadPool)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNumThreads is a wrapper around the C function g_thread_pool_get_num_threads.
func (recv *ThreadPool) GetNumThreads() uint32 {
	retC := C.g_thread_pool_get_num_threads((*C.GThreadPool)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_thread_pool_push : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// SetMaxThreads is a wrapper around the C function g_thread_pool_set_max_threads.
func (recv *ThreadPool) SetMaxThreads(maxThreads int32) (bool, error) {
	c_max_threads := (C.gint)(maxThreads)

	var cThrowableError *C.GError

	retC := C.g_thread_pool_set_max_threads((*C.GThreadPool)(recv.native), c_max_threads, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_thread_pool_set_sort_function : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unprocessed is a wrapper around the C function g_thread_pool_unprocessed.
func (recv *ThreadPool) Unprocessed() uint32 {
	retC := C.g_thread_pool_unprocessed((*C.GThreadPool)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// TimeVal is a wrapper around the C record GTimeVal.
type TimeVal struct {
	native *C.GTimeVal
	TvSec  int64
	TvUsec int64
}

func TimeValNewFromC(u unsafe.Pointer) *TimeVal {
	c := (*C.GTimeVal)(u)
	if c == nil {
		return nil
	}

	g := &TimeVal{
		TvSec:  (int64)(c.tv_sec),
		TvUsec: (int64)(c.tv_usec),
		native: c,
	}

	return g
}

func (recv *TimeVal) ToC() unsafe.Pointer {
	recv.native.tv_sec =
		(C.glong)(recv.TvSec)
	recv.native.tv_usec =
		(C.glong)(recv.TvUsec)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TimeVal with another TimeVal, and returns true if they represent the same GObject.
func (recv *TimeVal) Equals(other *TimeVal) bool {
	return other.ToC() == recv.ToC()
}

// TimeValFromIso8601 is a wrapper around the C function g_time_val_from_iso8601.
func TimeValFromIso8601(isoDate string) (bool, *TimeVal) {
	c_iso_date := C.CString(isoDate)
	defer C.free(unsafe.Pointer(c_iso_date))

	var c_time_ C.GTimeVal

	retC := C.g_time_val_from_iso8601(c_iso_date, &c_time_)
	retGo := retC == C.TRUE

	time := TimeValNewFromC(unsafe.Pointer(&c_time_))

	return retGo, time
}

// Add is a wrapper around the C function g_time_val_add.
func (recv *TimeVal) Add(microseconds int64) {
	c_microseconds := (C.glong)(microseconds)

	C.g_time_val_add((*C.GTimeVal)(recv.native), c_microseconds)

	return
}

// ToIso8601 is a wrapper around the C function g_time_val_to_iso8601.
func (recv *TimeVal) ToIso8601() string {
	retC := C.g_time_val_to_iso8601((*C.GTimeVal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// TimeZone is a wrapper around the C record GTimeZone.
type TimeZone struct {
	native *C.GTimeZone
}

func TimeZoneNewFromC(u unsafe.Pointer) *TimeZone {
	c := (*C.GTimeZone)(u)
	if c == nil {
		return nil
	}

	g := &TimeZone{native: c}

	return g
}

func (recv *TimeZone) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TimeZone with another TimeZone, and returns true if they represent the same GObject.
func (recv *TimeZone) Equals(other *TimeZone) bool {
	return other.ToC() == recv.ToC()
}

// TimeZoneNew is a wrapper around the C function g_time_zone_new.
func TimeZoneNew(identifier string) *TimeZone {
	c_identifier := C.CString(identifier)
	defer C.free(unsafe.Pointer(c_identifier))

	retC := C.g_time_zone_new(c_identifier)
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TimeZoneNewLocal is a wrapper around the C function g_time_zone_new_local.
func TimeZoneNewLocal() *TimeZone {
	retC := C.g_time_zone_new_local()
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TimeZoneNewUtc is a wrapper around the C function g_time_zone_new_utc.
func TimeZoneNewUtc() *TimeZone {
	retC := C.g_time_zone_new_utc()
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AdjustTime is a wrapper around the C function g_time_zone_adjust_time.
func (recv *TimeZone) AdjustTime(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_adjust_time((*C.GTimeZone)(recv.native), c_type, &c_time_)
	retGo := (int32)(retC)

	return retGo
}

// FindInterval is a wrapper around the C function g_time_zone_find_interval.
func (recv *TimeZone) FindInterval(type_ TimeType, time int64) int32 {
	c_type := (C.GTimeType)(type_)

	c_time_ := (C.gint64)(time)

	retC := C.g_time_zone_find_interval((*C.GTimeZone)(recv.native), c_type, c_time_)
	retGo := (int32)(retC)

	return retGo
}

// GetAbbreviation is a wrapper around the C function g_time_zone_get_abbreviation.
func (recv *TimeZone) GetAbbreviation(interval int32) string {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_abbreviation((*C.GTimeZone)(recv.native), c_interval)
	retGo := C.GoString(retC)

	return retGo
}

// GetOffset is a wrapper around the C function g_time_zone_get_offset.
func (recv *TimeZone) GetOffset(interval int32) int32 {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_get_offset((*C.GTimeZone)(recv.native), c_interval)
	retGo := (int32)(retC)

	return retGo
}

// IsDst is a wrapper around the C function g_time_zone_is_dst.
func (recv *TimeZone) IsDst(interval int32) bool {
	c_interval := (C.gint)(interval)

	retC := C.g_time_zone_is_dst((*C.GTimeZone)(recv.native), c_interval)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_time_zone_ref.
func (recv *TimeZone) Ref() *TimeZone {
	retC := C.g_time_zone_ref((*C.GTimeZone)(recv.native))
	retGo := TimeZoneNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_time_zone_unref.
func (recv *TimeZone) Unref() {
	C.g_time_zone_unref((*C.GTimeZone)(recv.native))

	return
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native *C.GTimer
}

func TimerNewFromC(u unsafe.Pointer) *Timer {
	c := (*C.GTimer)(u)
	if c == nil {
		return nil
	}

	g := &Timer{native: c}

	return g
}

func (recv *Timer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Timer with another Timer, and returns true if they represent the same GObject.
func (recv *Timer) Equals(other *Timer) bool {
	return other.ToC() == recv.ToC()
}

// TimerNew is a wrapper around the C function g_timer_new.
func TimerNew() *Timer {
	retC := C.g_timer_new()
	retGo := TimerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Continue is a wrapper around the C function g_timer_continue.
func (recv *Timer) Continue() {
	C.g_timer_continue((*C.GTimer)(recv.native))

	return
}

// Destroy is a wrapper around the C function g_timer_destroy.
func (recv *Timer) Destroy() {
	C.g_timer_destroy((*C.GTimer)(recv.native))

	return
}

// Elapsed is a wrapper around the C function g_timer_elapsed.
func (recv *Timer) Elapsed(microseconds uint64) float64 {
	c_microseconds := (C.gulong)(microseconds)

	retC := C.g_timer_elapsed((*C.GTimer)(recv.native), &c_microseconds)
	retGo := (float64)(retC)

	return retGo
}

// Reset is a wrapper around the C function g_timer_reset.
func (recv *Timer) Reset() {
	C.g_timer_reset((*C.GTimer)(recv.native))

	return
}

// Start is a wrapper around the C function g_timer_start.
func (recv *Timer) Start() {
	C.g_timer_start((*C.GTimer)(recv.native))

	return
}

// Stop is a wrapper around the C function g_timer_stop.
func (recv *Timer) Stop() {
	C.g_timer_stop((*C.GTimer)(recv.native))

	return
}

// TrashStack is a wrapper around the C record GTrashStack.
type TrashStack struct {
	native *C.GTrashStack
	// next : record
}

func TrashStackNewFromC(u unsafe.Pointer) *TrashStack {
	c := (*C.GTrashStack)(u)
	if c == nil {
		return nil
	}

	g := &TrashStack{native: c}

	return g
}

func (recv *TrashStack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TrashStack with another TrashStack, and returns true if they represent the same GObject.
func (recv *TrashStack) Equals(other *TrashStack) bool {
	return other.ToC() == recv.ToC()
}

// TrashStackHeight is a wrapper around the C function g_trash_stack_height.
func TrashStackHeight(stackP *TrashStack) uint32 {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	retC := C.g_trash_stack_height(c_stack_p)
	retGo := (uint32)(retC)

	return retGo
}

// g_trash_stack_peek : no return generator
// g_trash_stack_pop : no return generator
// g_trash_stack_push : unsupported parameter data_p : no type generator for gpointer (gpointer) for param data_p
// Tree is a wrapper around the C record GTree.
type Tree struct {
	native *C.GTree
}

func TreeNewFromC(u unsafe.Pointer) *Tree {
	c := (*C.GTree)(u)
	if c == nil {
		return nil
	}

	g := &Tree{native: c}

	return g
}

func (recv *Tree) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Tree with another Tree, and returns true if they represent the same GObject.
func (recv *Tree) Equals(other *Tree) bool {
	return other.ToC() == recv.ToC()
}

// g_tree_new : unsupported parameter key_compare_func : no type generator for CompareFunc (GCompareFunc) for param key_compare_func
// g_tree_new_full : unsupported parameter key_compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param key_compare_func
// g_tree_new_with_data : unsupported parameter key_compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param key_compare_func
// Destroy is a wrapper around the C function g_tree_destroy.
func (recv *Tree) Destroy() {
	C.g_tree_destroy((*C.GTree)(recv.native))

	return
}

// Unsupported : g_tree_foreach : unsupported parameter func : no type generator for TraverseFunc (GTraverseFunc) for param func

// Height is a wrapper around the C function g_tree_height.
func (recv *Tree) Height() int32 {
	retC := C.g_tree_height((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_tree_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_tree_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_tree_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key

// Nnodes is a wrapper around the C function g_tree_nnodes.
func (recv *Tree) Nnodes() int32 {
	retC := C.g_tree_nnodes((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Ref is a wrapper around the C function g_tree_ref.
func (recv *Tree) Ref() *Tree {
	retC := C.g_tree_ref((*C.GTree)(recv.native))
	retGo := TreeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_tree_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_tree_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_tree_search : unsupported parameter search_func : no type generator for CompareFunc (GCompareFunc) for param search_func

// Unsupported : g_tree_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_tree_traverse : unsupported parameter traverse_func : no type generator for TraverseFunc (GTraverseFunc) for param traverse_func

// Unref is a wrapper around the C function g_tree_unref.
func (recv *Tree) Unref() {
	C.g_tree_unref((*C.GTree)(recv.native))

	return
}

// Variant is a wrapper around the C record GVariant.
type Variant struct {
	native *C.GVariant
}

func VariantNewFromC(u unsafe.Pointer) *Variant {
	c := (*C.GVariant)(u)
	if c == nil {
		return nil
	}

	g := &Variant{native: c}

	return g
}

func (recv *Variant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Variant with another Variant, and returns true if they represent the same GObject.
func (recv *Variant) Equals(other *Variant) bool {
	return other.ToC() == recv.ToC()
}

// VariantNew is a wrapper around the C function g_variant_new.
func VariantNew(formatString string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_new(c_format_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_array : unsupported parameter children :

// VariantNewBoolean is a wrapper around the C function g_variant_new_boolean.
func VariantNewBoolean(value bool) *Variant {
	c_value :=
		boolToGboolean(value)

	retC := C.g_variant_new_boolean(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewByte is a wrapper around the C function g_variant_new_byte.
func VariantNewByte(value uint8) *Variant {
	c_value := (C.guchar)(value)

	retC := C.g_variant_new_byte(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewBytestring is a wrapper around the C function g_variant_new_bytestring.
func VariantNewBytestring(string_ []uint8) *Variant {
	c_string_array := make([]C.guint8, len(string_)+1, len(string_)+1)
	for i, item := range string_ {
		c := (C.guint8)(item)
		c_string_array[i] = c
	}
	c_string_array[len(string_)] = 0
	c_string_arrayPtr := &c_string_array[0]
	c_string := (*C.gchar)(unsafe.Pointer(c_string_arrayPtr))

	retC := C.g_variant_new_bytestring(c_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewBytestringArray is a wrapper around the C function g_variant_new_bytestring_array.
func VariantNewBytestringArray(strv []string) *Variant {
	c_strv_array := make([]*C.gchar, len(strv)+1, len(strv)+1)
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_array[len(strv)] = nil
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.gchar)(unsafe.Pointer(c_strv_arrayPtr))

	c_length := (C.gssize)(len(strv))

	retC := C.g_variant_new_bytestring_array(c_strv, c_length)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewDictEntry is a wrapper around the C function g_variant_new_dict_entry.
func VariantNewDictEntry(key *Variant, value *Variant) *Variant {
	c_key := (*C.GVariant)(C.NULL)
	if key != nil {
		c_key = (*C.GVariant)(key.ToC())
	}

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_new_dict_entry(c_key, c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewDouble is a wrapper around the C function g_variant_new_double.
func VariantNewDouble(value float64) *Variant {
	c_value := (C.gdouble)(value)

	retC := C.g_variant_new_double(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_fixed_array : unsupported parameter elements : no type generator for gpointer (gconstpointer) for param elements

// Unsupported : g_variant_new_from_data : unsupported parameter notify : no type generator for DestroyNotify (GDestroyNotify) for param notify

// VariantNewHandle is a wrapper around the C function g_variant_new_handle.
func VariantNewHandle(value int32) *Variant {
	c_value := (C.gint32)(value)

	retC := C.g_variant_new_handle(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt16 is a wrapper around the C function g_variant_new_int16.
func VariantNewInt16(value int16) *Variant {
	c_value := (C.gint16)(value)

	retC := C.g_variant_new_int16(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt32 is a wrapper around the C function g_variant_new_int32.
func VariantNewInt32(value int32) *Variant {
	c_value := (C.gint32)(value)

	retC := C.g_variant_new_int32(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewInt64 is a wrapper around the C function g_variant_new_int64.
func VariantNewInt64(value int64) *Variant {
	c_value := (C.gint64)(value)

	retC := C.g_variant_new_int64(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewMaybe is a wrapper around the C function g_variant_new_maybe.
func VariantNewMaybe(childType *VariantType, child *Variant) *Variant {
	c_child_type := (*C.GVariantType)(C.NULL)
	if childType != nil {
		c_child_type = (*C.GVariantType)(childType.ToC())
	}

	c_child := (*C.GVariant)(C.NULL)
	if child != nil {
		c_child = (*C.GVariant)(child.ToC())
	}

	retC := C.g_variant_new_maybe(c_child_type, c_child)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewObjectPath is a wrapper around the C function g_variant_new_object_path.
func VariantNewObjectPath(objectPath string) *Variant {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_variant_new_object_path(c_object_path)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewObjv is a wrapper around the C function g_variant_new_objv.
func VariantNewObjv(strv []string) *Variant {
	c_strv_array := make([]*C.gchar, len(strv)+1, len(strv)+1)
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_array[len(strv)] = nil
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.gchar)(unsafe.Pointer(c_strv_arrayPtr))

	c_length := (C.gssize)(len(strv))

	retC := C.g_variant_new_objv(c_strv, c_length)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewParsed is a wrapper around the C function g_variant_new_parsed.
func VariantNewParsed(format string, args ...interface{}) *Variant {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_variant_new_parsed(c_format)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_parsed_va : unsupported parameter app : no type generator for va_list (va_list*) for param app

// VariantNewSignature is a wrapper around the C function g_variant_new_signature.
func VariantNewSignature(signature string) *Variant {
	c_signature := C.CString(signature)
	defer C.free(unsafe.Pointer(c_signature))

	retC := C.g_variant_new_signature(c_signature)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewString is a wrapper around the C function g_variant_new_string.
func VariantNewString(string_ string) *Variant {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_new_string(c_string)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewStrv is a wrapper around the C function g_variant_new_strv.
func VariantNewStrv(strv []string) *Variant {
	c_strv_array := make([]*C.gchar, len(strv)+1, len(strv)+1)
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_array[len(strv)] = nil
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.gchar)(unsafe.Pointer(c_strv_arrayPtr))

	c_length := (C.gssize)(len(strv))

	retC := C.g_variant_new_strv(c_strv, c_length)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_tuple : unsupported parameter children :

// VariantNewUint16 is a wrapper around the C function g_variant_new_uint16.
func VariantNewUint16(value uint16) *Variant {
	c_value := (C.guint16)(value)

	retC := C.g_variant_new_uint16(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewUint32 is a wrapper around the C function g_variant_new_uint32.
func VariantNewUint32(value uint32) *Variant {
	c_value := (C.guint32)(value)

	retC := C.g_variant_new_uint32(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantNewUint64 is a wrapper around the C function g_variant_new_uint64.
func VariantNewUint64(value uint64) *Variant {
	c_value := (C.guint64)(value)

	retC := C.g_variant_new_uint64(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_new_va : unsupported parameter endptr : in string with indirection level of 2

// VariantNewVariant is a wrapper around the C function g_variant_new_variant.
func VariantNewVariant(value *Variant) *Variant {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_new_variant(c_value)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantIsObjectPath is a wrapper around the C function g_variant_is_object_path.
func VariantIsObjectPath(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_object_path(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// VariantIsSignature is a wrapper around the C function g_variant_is_signature.
func VariantIsSignature(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_signature(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// g_variant_parse : unsupported parameter endptr : in string with indirection level of 2
// VariantParseErrorQuark is a wrapper around the C function g_variant_parse_error_quark.
func VariantParseErrorQuark() Quark {
	retC := C.g_variant_parse_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// VariantParserGetErrorQuark is a wrapper around the C function g_variant_parser_get_error_quark.
func VariantParserGetErrorQuark() Quark {
	retC := C.g_variant_parser_get_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// Byteswap is a wrapper around the C function g_variant_byteswap.
func (recv *Variant) Byteswap() *Variant {
	retC := C.g_variant_byteswap((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckFormatString is a wrapper around the C function g_variant_check_format_string.
func (recv *Variant) CheckFormatString(formatString string, copyOnly bool) bool {
	c_format_string := C.CString(formatString)
	defer C.free(unsafe.Pointer(c_format_string))

	c_copy_only :=
		boolToGboolean(copyOnly)

	retC := C.g_variant_check_format_string((*C.GVariant)(recv.native), c_format_string, c_copy_only)
	retGo := retC == C.TRUE

	return retGo
}

// Classify is a wrapper around the C function g_variant_classify.
func (recv *Variant) Classify() VariantClass {
	retC := C.g_variant_classify((*C.GVariant)(recv.native))
	retGo := (VariantClass)(retC)

	return retGo
}

// Compare is a wrapper around the C function g_variant_compare.
func (recv *Variant) Compare(two *Variant) int32 {
	c_two := (C.gconstpointer)(C.NULL)
	if two != nil {
		c_two = (C.gconstpointer)(two.ToC())
	}

	retC := C.g_variant_compare((C.gconstpointer)(recv.native), c_two)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_variant_dup_bytestring : array return type :

// DupBytestringArray is a wrapper around the C function g_variant_dup_bytestring_array.
func (recv *Variant) DupBytestringArray() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_bytestring_array((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// DupObjv is a wrapper around the C function g_variant_dup_objv.
func (recv *Variant) DupObjv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_objv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// DupString is a wrapper around the C function g_variant_dup_string.
func (recv *Variant) DupString() (string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_string((*C.GVariant)(recv.native), &c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	length := (uint64)(c_length)

	return retGo, length
}

// DupStrv is a wrapper around the C function g_variant_dup_strv.
func (recv *Variant) DupStrv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_strv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// Equal is a wrapper around the C function g_variant_equal.
func (recv *Variant) Equal(two *Variant) bool {
	c_two := (C.gconstpointer)(C.NULL)
	if two != nil {
		c_two = (C.gconstpointer)(two.ToC())
	}

	retC := C.g_variant_equal((C.gconstpointer)(recv.native), c_two)
	retGo := retC == C.TRUE

	return retGo
}

// Get is a wrapper around the C function g_variant_get.
func (recv *Variant) Get(formatString string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_variant_get((*C.GVariant)(recv.native), c_format_string)

	return
}

// GetBoolean is a wrapper around the C function g_variant_get_boolean.
func (recv *Variant) GetBoolean() bool {
	retC := C.g_variant_get_boolean((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetByte is a wrapper around the C function g_variant_get_byte.
func (recv *Variant) GetByte() uint8 {
	retC := C.g_variant_get_byte((*C.GVariant)(recv.native))
	retGo := (uint8)(retC)

	return retGo
}

// Unsupported : g_variant_get_bytestring : array return type :

// GetBytestringArray is a wrapper around the C function g_variant_get_bytestring_array.
func (recv *Variant) GetBytestringArray() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_bytestring_array((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	length := (uint64)(c_length)

	return retGo, length
}

// GetChild is a wrapper around the C function g_variant_get_child.
func (recv *Variant) GetChild(index uint64, formatString string, args ...interface{}) {
	c_index_ := (C.gsize)(index)

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_variant_get_child((*C.GVariant)(recv.native), c_index_, c_format_string)

	return
}

// GetChildValue is a wrapper around the C function g_variant_get_child_value.
func (recv *Variant) GetChildValue(index uint64) *Variant {
	c_index_ := (C.gsize)(index)

	retC := C.g_variant_get_child_value((*C.GVariant)(recv.native), c_index_)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_get_data : no return generator

// GetDouble is a wrapper around the C function g_variant_get_double.
func (recv *Variant) GetDouble() float64 {
	retC := C.g_variant_get_double((*C.GVariant)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Unsupported : g_variant_get_fixed_array : no type generator for gpointer (gconstpointer) for array return

// GetHandle is a wrapper around the C function g_variant_get_handle.
func (recv *Variant) GetHandle() int32 {
	retC := C.g_variant_get_handle((*C.GVariant)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInt16 is a wrapper around the C function g_variant_get_int16.
func (recv *Variant) GetInt16() int16 {
	retC := C.g_variant_get_int16((*C.GVariant)(recv.native))
	retGo := (int16)(retC)

	return retGo
}

// GetInt32 is a wrapper around the C function g_variant_get_int32.
func (recv *Variant) GetInt32() int32 {
	retC := C.g_variant_get_int32((*C.GVariant)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInt64 is a wrapper around the C function g_variant_get_int64.
func (recv *Variant) GetInt64() int64 {
	retC := C.g_variant_get_int64((*C.GVariant)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetMaybe is a wrapper around the C function g_variant_get_maybe.
func (recv *Variant) GetMaybe() *Variant {
	retC := C.g_variant_get_maybe((*C.GVariant)(recv.native))
	var retGo (*Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetNormalForm is a wrapper around the C function g_variant_get_normal_form.
func (recv *Variant) GetNormalForm() *Variant {
	retC := C.g_variant_get_normal_form((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetObjv is a wrapper around the C function g_variant_get_objv.
func (recv *Variant) GetObjv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_objv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	length := (uint64)(c_length)

	return retGo, length
}

// GetSize is a wrapper around the C function g_variant_get_size.
func (recv *Variant) GetSize() uint64 {
	retC := C.g_variant_get_size((*C.GVariant)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetString is a wrapper around the C function g_variant_get_string.
func (recv *Variant) GetString() (string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_string((*C.GVariant)(recv.native), &c_length)
	retGo := C.GoString(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetStrv is a wrapper around the C function g_variant_get_strv.
func (recv *Variant) GetStrv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_strv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	length := (uint64)(c_length)

	return retGo, length
}

// GetType is a wrapper around the C function g_variant_get_type.
func (recv *Variant) GetType() *VariantType {
	retC := C.g_variant_get_type((*C.GVariant)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTypeString is a wrapper around the C function g_variant_get_type_string.
func (recv *Variant) GetTypeString() string {
	retC := C.g_variant_get_type_string((*C.GVariant)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUint16 is a wrapper around the C function g_variant_get_uint16.
func (recv *Variant) GetUint16() uint16 {
	retC := C.g_variant_get_uint16((*C.GVariant)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetUint32 is a wrapper around the C function g_variant_get_uint32.
func (recv *Variant) GetUint32() uint32 {
	retC := C.g_variant_get_uint32((*C.GVariant)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetUint64 is a wrapper around the C function g_variant_get_uint64.
func (recv *Variant) GetUint64() uint64 {
	retC := C.g_variant_get_uint64((*C.GVariant)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_variant_get_va : unsupported parameter endptr : in string with indirection level of 2

// GetVariant is a wrapper around the C function g_variant_get_variant.
func (recv *Variant) GetVariant() *Variant {
	retC := C.g_variant_get_variant((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Hash is a wrapper around the C function g_variant_hash.
func (recv *Variant) Hash() uint32 {
	retC := C.g_variant_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsContainer is a wrapper around the C function g_variant_is_container.
func (recv *Variant) IsContainer() bool {
	retC := C.g_variant_is_container((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsFloating is a wrapper around the C function g_variant_is_floating.
func (recv *Variant) IsFloating() bool {
	retC := C.g_variant_is_floating((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsNormalForm is a wrapper around the C function g_variant_is_normal_form.
func (recv *Variant) IsNormalForm() bool {
	retC := C.g_variant_is_normal_form((*C.GVariant)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsOfType is a wrapper around the C function g_variant_is_of_type.
func (recv *Variant) IsOfType(type_ *VariantType) bool {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	retC := C.g_variant_is_of_type((*C.GVariant)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// IterNew is a wrapper around the C function g_variant_iter_new.
func (recv *Variant) IterNew() *VariantIter {
	retC := C.g_variant_iter_new((*C.GVariant)(recv.native))
	retGo := VariantIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lookup is a wrapper around the C function g_variant_lookup.
func (recv *Variant) Lookup(key string, formatString string, args ...interface{}) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_lookup((*C.GVariant)(recv.native), c_key, c_format_string)
	retGo := retC == C.TRUE

	return retGo
}

// LookupValue is a wrapper around the C function g_variant_lookup_value.
func (recv *Variant) LookupValue(key string, expectedType *VariantType) *Variant {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_expected_type := (*C.GVariantType)(C.NULL)
	if expectedType != nil {
		c_expected_type = (*C.GVariantType)(expectedType.ToC())
	}

	retC := C.g_variant_lookup_value((*C.GVariant)(recv.native), c_key, c_expected_type)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NChildren is a wrapper around the C function g_variant_n_children.
func (recv *Variant) NChildren() uint64 {
	retC := C.g_variant_n_children((*C.GVariant)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Print is a wrapper around the C function g_variant_print.
func (recv *Variant) Print(typeAnnotate bool) string {
	c_type_annotate :=
		boolToGboolean(typeAnnotate)

	retC := C.g_variant_print((*C.GVariant)(recv.native), c_type_annotate)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// PrintString is a wrapper around the C function g_variant_print_string.
func (recv *Variant) PrintString(string_ *String, typeAnnotate bool) *String {
	c_string := (*C.GString)(C.NULL)
	if string_ != nil {
		c_string = (*C.GString)(string_.ToC())
	}

	c_type_annotate :=
		boolToGboolean(typeAnnotate)

	retC := C.g_variant_print_string((*C.GVariant)(recv.native), c_string, c_type_annotate)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_variant_ref.
func (recv *Variant) Ref() *Variant {
	retC := C.g_variant_ref((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RefSink is a wrapper around the C function g_variant_ref_sink.
func (recv *Variant) RefSink() *Variant {
	retC := C.g_variant_ref_sink((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_store : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// TakeRef is a wrapper around the C function g_variant_take_ref.
func (recv *Variant) TakeRef() *Variant {
	retC := C.g_variant_take_ref((*C.GVariant)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_variant_unref.
func (recv *Variant) Unref() {
	C.g_variant_unref((*C.GVariant)(recv.native))

	return
}

// VariantBuilder is a wrapper around the C record GVariantBuilder.
type VariantBuilder struct {
	native *C.GVariantBuilder
}

func VariantBuilderNewFromC(u unsafe.Pointer) *VariantBuilder {
	c := (*C.GVariantBuilder)(u)
	if c == nil {
		return nil
	}

	g := &VariantBuilder{native: c}

	return g
}

func (recv *VariantBuilder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VariantBuilder with another VariantBuilder, and returns true if they represent the same GObject.
func (recv *VariantBuilder) Equals(other *VariantBuilder) bool {
	return other.ToC() == recv.ToC()
}

// VariantBuilderNew is a wrapper around the C function g_variant_builder_new.
func VariantBuilderNew(type_ *VariantType) *VariantBuilder {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	retC := C.g_variant_builder_new(c_type)
	retGo := VariantBuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function g_variant_builder_add.
func (recv *VariantBuilder) Add(formatString string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._g_variant_builder_add((*C.GVariantBuilder)(recv.native), c_format_string)

	return
}

// AddParsed is a wrapper around the C function g_variant_builder_add_parsed.
func (recv *VariantBuilder) AddParsed(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_variant_builder_add_parsed((*C.GVariantBuilder)(recv.native), c_format)

	return
}

// AddValue is a wrapper around the C function g_variant_builder_add_value.
func (recv *VariantBuilder) AddValue(value *Variant) {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_variant_builder_add_value((*C.GVariantBuilder)(recv.native), c_value)

	return
}

// Clear is a wrapper around the C function g_variant_builder_clear.
func (recv *VariantBuilder) Clear() {
	C.g_variant_builder_clear((*C.GVariantBuilder)(recv.native))

	return
}

// Close is a wrapper around the C function g_variant_builder_close.
func (recv *VariantBuilder) Close() {
	C.g_variant_builder_close((*C.GVariantBuilder)(recv.native))

	return
}

// End is a wrapper around the C function g_variant_builder_end.
func (recv *VariantBuilder) End() *Variant {
	retC := C.g_variant_builder_end((*C.GVariantBuilder)(recv.native))
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_variant_builder_init.
func (recv *VariantBuilder) Init(type_ *VariantType) {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	C.g_variant_builder_init((*C.GVariantBuilder)(recv.native), c_type)

	return
}

// Open is a wrapper around the C function g_variant_builder_open.
func (recv *VariantBuilder) Open(type_ *VariantType) {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	C.g_variant_builder_open((*C.GVariantBuilder)(recv.native), c_type)

	return
}

// Ref is a wrapper around the C function g_variant_builder_ref.
func (recv *VariantBuilder) Ref() *VariantBuilder {
	retC := C.g_variant_builder_ref((*C.GVariantBuilder)(recv.native))
	retGo := VariantBuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_variant_builder_unref.
func (recv *VariantBuilder) Unref() {
	C.g_variant_builder_unref((*C.GVariantBuilder)(recv.native))

	return
}

// VariantIter is a wrapper around the C record GVariantIter.
type VariantIter struct {
	native *C.GVariantIter
	// Private : x
}

func VariantIterNewFromC(u unsafe.Pointer) *VariantIter {
	c := (*C.GVariantIter)(u)
	if c == nil {
		return nil
	}

	g := &VariantIter{native: c}

	return g
}

func (recv *VariantIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VariantIter with another VariantIter, and returns true if they represent the same GObject.
func (recv *VariantIter) Equals(other *VariantIter) bool {
	return other.ToC() == recv.ToC()
}

// Copy is a wrapper around the C function g_variant_iter_copy.
func (recv *VariantIter) Copy() *VariantIter {
	retC := C.g_variant_iter_copy((*C.GVariantIter)(recv.native))
	retGo := VariantIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_variant_iter_free.
func (recv *VariantIter) Free() {
	C.g_variant_iter_free((*C.GVariantIter)(recv.native))

	return
}

// Init is a wrapper around the C function g_variant_iter_init.
func (recv *VariantIter) Init(value *Variant) uint64 {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_variant_iter_init((*C.GVariantIter)(recv.native), c_value)
	retGo := (uint64)(retC)

	return retGo
}

// Loop is a wrapper around the C function g_variant_iter_loop.
func (recv *VariantIter) Loop(formatString string, args ...interface{}) bool {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_iter_loop((*C.GVariantIter)(recv.native), c_format_string)
	retGo := retC == C.TRUE

	return retGo
}

// NChildren is a wrapper around the C function g_variant_iter_n_children.
func (recv *VariantIter) NChildren() uint64 {
	retC := C.g_variant_iter_n_children((*C.GVariantIter)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Next is a wrapper around the C function g_variant_iter_next.
func (recv *VariantIter) Next(formatString string, args ...interface{}) bool {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	retC := C._g_variant_iter_next((*C.GVariantIter)(recv.native), c_format_string)
	retGo := retC == C.TRUE

	return retGo
}

// NextValue is a wrapper around the C function g_variant_iter_next_value.
func (recv *VariantIter) NextValue() *Variant {
	retC := C.g_variant_iter_next_value((*C.GVariantIter)(recv.native))
	var retGo (*Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// VariantType is a wrapper around the C record GVariantType.
type VariantType struct {
	native *C.GVariantType
}

func VariantTypeNewFromC(u unsafe.Pointer) *VariantType {
	c := (*C.GVariantType)(u)
	if c == nil {
		return nil
	}

	g := &VariantType{native: c}

	return g
}

func (recv *VariantType) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VariantType with another VariantType, and returns true if they represent the same GObject.
func (recv *VariantType) Equals(other *VariantType) bool {
	return other.ToC() == recv.ToC()
}

// VariantTypeNew is a wrapper around the C function g_variant_type_new.
func VariantTypeNew(typeString string) *VariantType {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	retC := C.g_variant_type_new(c_type_string)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeNewArray is a wrapper around the C function g_variant_type_new_array.
func VariantTypeNewArray(element *VariantType) *VariantType {
	c_element := (*C.GVariantType)(C.NULL)
	if element != nil {
		c_element = (*C.GVariantType)(element.ToC())
	}

	retC := C.g_variant_type_new_array(c_element)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeNewDictEntry is a wrapper around the C function g_variant_type_new_dict_entry.
func VariantTypeNewDictEntry(key *VariantType, value *VariantType) *VariantType {
	c_key := (*C.GVariantType)(C.NULL)
	if key != nil {
		c_key = (*C.GVariantType)(key.ToC())
	}

	c_value := (*C.GVariantType)(C.NULL)
	if value != nil {
		c_value = (*C.GVariantType)(value.ToC())
	}

	retC := C.g_variant_type_new_dict_entry(c_key, c_value)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeNewMaybe is a wrapper around the C function g_variant_type_new_maybe.
func VariantTypeNewMaybe(element *VariantType) *VariantType {
	c_element := (*C.GVariantType)(C.NULL)
	if element != nil {
		c_element = (*C.GVariantType)(element.ToC())
	}

	retC := C.g_variant_type_new_maybe(c_element)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_type_new_tuple : unsupported parameter items :

// VariantTypeChecked is a wrapper around the C function g_variant_type_checked_.
func VariantTypeChecked(arg0 string) *VariantType {
	c_arg0 := C.CString(arg0)
	defer C.free(unsafe.Pointer(c_arg0))

	retC := C.g_variant_type_checked_(c_arg0)
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VariantTypeStringIsValid is a wrapper around the C function g_variant_type_string_is_valid.
func VariantTypeStringIsValid(typeString string) bool {
	c_type_string := C.CString(typeString)
	defer C.free(unsafe.Pointer(c_type_string))

	retC := C.g_variant_type_string_is_valid(c_type_string)
	retGo := retC == C.TRUE

	return retGo
}

// VariantTypeStringScan is a wrapper around the C function g_variant_type_string_scan.
func VariantTypeStringScan(string_ string, limit string) (bool, string) {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_limit := C.CString(limit)
	defer C.free(unsafe.Pointer(c_limit))

	var c_endptr *C.gchar

	retC := C.g_variant_type_string_scan(c_string, c_limit, &c_endptr)
	retGo := retC == C.TRUE

	endptr := C.GoString(c_endptr)
	defer C.free(unsafe.Pointer(c_endptr))

	return retGo, endptr
}

// Copy is a wrapper around the C function g_variant_type_copy.
func (recv *VariantType) Copy() *VariantType {
	retC := C.g_variant_type_copy((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DupString is a wrapper around the C function g_variant_type_dup_string.
func (recv *VariantType) DupString() string {
	retC := C.g_variant_type_dup_string((*C.GVariantType)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Element is a wrapper around the C function g_variant_type_element.
func (recv *VariantType) Element() *VariantType {
	retC := C.g_variant_type_element((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_variant_type_equal.
func (recv *VariantType) Equal(type2 *VariantType) bool {
	c_type2 := (C.gconstpointer)(C.NULL)
	if type2 != nil {
		c_type2 = (C.gconstpointer)(type2.ToC())
	}

	retC := C.g_variant_type_equal((C.gconstpointer)(recv.native), c_type2)
	retGo := retC == C.TRUE

	return retGo
}

// First is a wrapper around the C function g_variant_type_first.
func (recv *VariantType) First() *VariantType {
	retC := C.g_variant_type_first((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_variant_type_free.
func (recv *VariantType) Free() {
	C.g_variant_type_free((*C.GVariantType)(recv.native))

	return
}

// GetStringLength is a wrapper around the C function g_variant_type_get_string_length.
func (recv *VariantType) GetStringLength() uint64 {
	retC := C.g_variant_type_get_string_length((*C.GVariantType)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Hash is a wrapper around the C function g_variant_type_hash.
func (recv *VariantType) Hash() uint32 {
	retC := C.g_variant_type_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsArray is a wrapper around the C function g_variant_type_is_array.
func (recv *VariantType) IsArray() bool {
	retC := C.g_variant_type_is_array((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsBasic is a wrapper around the C function g_variant_type_is_basic.
func (recv *VariantType) IsBasic() bool {
	retC := C.g_variant_type_is_basic((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsContainer is a wrapper around the C function g_variant_type_is_container.
func (recv *VariantType) IsContainer() bool {
	retC := C.g_variant_type_is_container((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsDefinite is a wrapper around the C function g_variant_type_is_definite.
func (recv *VariantType) IsDefinite() bool {
	retC := C.g_variant_type_is_definite((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsDictEntry is a wrapper around the C function g_variant_type_is_dict_entry.
func (recv *VariantType) IsDictEntry() bool {
	retC := C.g_variant_type_is_dict_entry((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsMaybe is a wrapper around the C function g_variant_type_is_maybe.
func (recv *VariantType) IsMaybe() bool {
	retC := C.g_variant_type_is_maybe((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSubtypeOf is a wrapper around the C function g_variant_type_is_subtype_of.
func (recv *VariantType) IsSubtypeOf(supertype *VariantType) bool {
	c_supertype := (*C.GVariantType)(C.NULL)
	if supertype != nil {
		c_supertype = (*C.GVariantType)(supertype.ToC())
	}

	retC := C.g_variant_type_is_subtype_of((*C.GVariantType)(recv.native), c_supertype)
	retGo := retC == C.TRUE

	return retGo
}

// IsTuple is a wrapper around the C function g_variant_type_is_tuple.
func (recv *VariantType) IsTuple() bool {
	retC := C.g_variant_type_is_tuple((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsVariant is a wrapper around the C function g_variant_type_is_variant.
func (recv *VariantType) IsVariant() bool {
	retC := C.g_variant_type_is_variant((*C.GVariantType)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Key is a wrapper around the C function g_variant_type_key.
func (recv *VariantType) Key() *VariantType {
	retC := C.g_variant_type_key((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NItems is a wrapper around the C function g_variant_type_n_items.
func (recv *VariantType) NItems() uint64 {
	retC := C.g_variant_type_n_items((*C.GVariantType)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Next is a wrapper around the C function g_variant_type_next.
func (recv *VariantType) Next() *VariantType {
	retC := C.g_variant_type_next((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PeekString is a wrapper around the C function g_variant_type_peek_string.
func (recv *VariantType) PeekString() string {
	retC := C.g_variant_type_peek_string((*C.GVariantType)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Value is a wrapper around the C function g_variant_type_value.
func (recv *VariantType) Value() *VariantType {
	retC := C.g_variant_type_value((*C.GVariantType)(recv.native))
	retGo := VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

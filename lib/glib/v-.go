// Code generated - DO NOT EDIT.
// +build !glib_2.2,!glib_2.4,!glib_2.6,!glib_2.8,!glib_2.10,!glib_2.12,!glib_2.14,!glib_2.16,!glib_2.18,!glib_2.20,!glib_2.22,!glib_2.24,!glib_2.26,!glib_2.28,!glib_2.30,!glib_2.32,!glib_2.34,!glib_2.36,!glib_2.38,!glib_2.40,!glib_2.44,!glib_2.46,!glib_2.48,!glib_2.50,!glib_2.52,!glib_2.54,!glib_2.56

package glib

import (
	"fmt"
	"reflect"
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

	static gchar* _g_strdup_printf(const gchar* format) {
		return g_strdup_printf(format);
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
import "C"

var gobjectClassGoTypeMap = make(map[string]reflect.Type)

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
const GINT32_FORMAT string = C.G_GINT32_FORMAT
const GINT64_FORMAT string = C.G_GINT64_FORMAT
const GNUC_FUNCTION string = C.G_GNUC_FUNCTION
const GNUC_PRETTY_FUNCTION string = C.G_GNUC_PRETTY_FUNCTION
const GUINT16_FORMAT string = C.G_GUINT16_FORMAT
const GUINT32_FORMAT string = C.G_GUINT32_FORMAT
const GUINT64_FORMAT string = C.G_GUINT64_FORMAT
const HAVE_GINT64 int32 = C.G_HAVE_GINT64
const HAVE_GNUC_VARARGS int32 = C.G_HAVE_GNUC_VARARGS
const HAVE_GNUC_VISIBILITY int32 = C.G_HAVE_GNUC_VISIBILITY
const HAVE_GROWING_STACK int32 = C.G_HAVE_GROWING_STACK
const HAVE_ISO_VARARGS int32 = C.G_HAVE_ISO_VARARGS
const HOOK_FLAG_USER_SHIFT int32 = C.G_HOOK_FLAG_USER_SHIFT
const IEEE754_DOUBLE_BIAS int32 = C.G_IEEE754_DOUBLE_BIAS
const IEEE754_FLOAT_BIAS int32 = C.G_IEEE754_FLOAT_BIAS

// Blacklisted : KEY_FILE_DESKTOP_ACTION_GROUP_PREFIX

const KEY_FILE_DESKTOP_KEY_FULLNAME string = C.G_KEY_FILE_DESKTOP_KEY_FULLNAME
const KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN string = C.G_KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN
const KEY_FILE_DESKTOP_KEY_KEYWORDS string = C.G_KEY_FILE_DESKTOP_KEY_KEYWORDS
const LITTLE_ENDIAN int32 = C.G_LITTLE_ENDIAN
const LN10 float64 = C.G_LN10
const LN2 float64 = C.G_LN2
const LOG_2_BASE_10 float64 = C.G_LOG_2_BASE_10

// Blacklisted : LOG_DOMAIN

const LOG_FATAL_MASK int32 = C.G_LOG_FATAL_MASK
const LOG_LEVEL_USER_SHIFT int32 = C.G_LOG_LEVEL_USER_SHIFT
const MAJOR_VERSION int32 = C.GLIB_MAJOR_VERSION
const MAXINT64 int64 = C.G_MAXINT64
const MAXUINT64 uint64 = C.G_MAXUINT64
const MICRO_VERSION int32 = C.GLIB_MICRO_VERSION
const MININT64 int64 = C.G_MININT64
const MINOR_VERSION int32 = C.GLIB_MINOR_VERSION
const MODULE_SUFFIX string = C.G_MODULE_SUFFIX
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
const SQRT2 float64 = C.G_SQRT2
const STR_DELIMITERS string = C.G_STR_DELIMITERS
const SYSDEF_AF_INET int32 = C.GLIB_SYSDEF_AF_INET
const SYSDEF_AF_INET6 int32 = C.GLIB_SYSDEF_AF_INET6
const SYSDEF_AF_UNIX int32 = C.GLIB_SYSDEF_AF_UNIX
const SYSDEF_MSG_DONTROUTE int32 = C.GLIB_SYSDEF_MSG_DONTROUTE
const SYSDEF_MSG_OOB int32 = C.GLIB_SYSDEF_MSG_OOB
const SYSDEF_MSG_PEEK int32 = C.GLIB_SYSDEF_MSG_PEEK
const URI_RESERVED_CHARS_GENERIC_DELIMITERS string = C.G_URI_RESERVED_CHARS_GENERIC_DELIMITERS
const URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS string = C.G_URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS
const USEC_PER_SEC int32 = C.G_USEC_PER_SEC
const VA_COPY_AS_ARRAY int32 = C.G_VA_COPY_AS_ARRAY

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

// Basename is a wrapper around the C function g_basename.
func Basename(fileName string) string {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	retC := C.g_basename(c_file_name)
	retGo := C.GoString(retC)

	return retGo
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

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Blacklisted : g_byte_array_free

// Unsupported : g_byte_array_new : array return type :

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

// DatalistGetData is a wrapper around the C function g_datalist_get_data.
func DatalistGetData(datalist *Data, key string) uintptr {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_datalist_get_data(c_datalist, c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// DatalistIdGetData is a wrapper around the C function g_datalist_id_get_data.
func DatalistIdGetData(datalist *Data, keyId Quark) uintptr {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_datalist_id_get_data(c_datalist, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// DatalistIdRemoveNoNotify is a wrapper around the C function g_datalist_id_remove_no_notify.
func DatalistIdRemoveNoNotify(datalist *Data, keyId Quark) uintptr {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_datalist_id_remove_no_notify(c_datalist, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_datalist_id_set_data_full : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// DatalistInit is a wrapper around the C function g_datalist_init.
func DatalistInit(datalist *Data) {
	c_datalist := (**C.GData)(C.NULL)
	if datalist != nil {
		c_datalist = (**C.GData)(datalist.ToC())
	}

	C.g_datalist_init(c_datalist)

	return
}

// DatasetDestroy is a wrapper around the C function g_dataset_destroy.
func DatasetDestroy(datasetLocation uintptr) {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	C.g_dataset_destroy(c_dataset_location)

	return
}

// Unsupported : g_dataset_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// DatasetIdGetData is a wrapper around the C function g_dataset_id_get_data.
func DatasetIdGetData(datasetLocation uintptr, keyId Quark) uintptr {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_dataset_id_get_data(c_dataset_location, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// DatasetIdRemoveNoNotify is a wrapper around the C function g_dataset_id_remove_no_notify.
func DatasetIdRemoveNoNotify(datasetLocation uintptr, keyId Quark) uintptr {
	c_dataset_location := (C.gconstpointer)(datasetLocation)

	c_key_id := (C.GQuark)(keyId)

	retC := C.g_dataset_id_remove_no_notify(c_dataset_location, c_key_id)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dataset_id_set_data_full : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// DirectEqual is a wrapper around the C function g_direct_equal.
func DirectEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_direct_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// DirectHash is a wrapper around the C function g_direct_hash.
func DirectHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_direct_hash(c_v)
	retGo := (uint32)(retC)

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

// FileTest is a wrapper around the C function g_file_test.
func FileTest(filename string, test GFileTest) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_test := (C.GFileTest)(test)

	retC := C.g_file_test(c_filename, c_test)
	retGo := retC == C.TRUE

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

// Free is a wrapper around the C function g_free.
func Free(mem uintptr) {
	c_mem := (C.gpointer)(mem)

	C.g_free(c_mem)

	return
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

// GetHomeDir is a wrapper around the C function g_get_home_dir.
func GetHomeDir() string {
	retC := C.g_get_home_dir()
	retGo := C.GoString(retC)

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

// GetTmpDir is a wrapper around the C function g_get_tmp_dir.
func GetTmpDir() string {
	retC := C.g_get_tmp_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserName is a wrapper around the C function g_get_user_name.
func GetUserName() string {
	retC := C.g_get_user_name()
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

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type : Blacklisted record : GIConv

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// IdleRemoveByData is a wrapper around the C function g_idle_remove_by_data.
func IdleRemoveByData(data uintptr) bool {
	c_data := (C.gpointer)(data)

	retC := C.g_idle_remove_by_data(c_data)
	retGo := retC == C.TRUE

	return retGo
}

// IdleSourceNew is a wrapper around the C function g_idle_source_new.
func IdleSourceNew() *Source {
	retC := C.g_idle_source_new()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IntEqual is a wrapper around the C function g_int_equal.
func IntEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_int_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// IntHash is a wrapper around the C function g_int_hash.
func IntHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_int_hash(c_v)
	retGo := (uint32)(retC)

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

// LogDefaultHandler is a wrapper around the C function g_log_default_handler.
func LogDefaultHandler(logDomain string, logLevel LogLevelFlags, message string, unusedData uintptr) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_log_level := (C.GLogLevelFlags)(logLevel)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	c_unused_data := (C.gpointer)(unusedData)

	C.g_log_default_handler(c_log_domain, c_log_level, c_message, c_unused_data)

	return
}

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

// MainDepth is a wrapper around the C function g_main_depth.
func MainDepth() int32 {
	retC := C.g_main_depth()
	retGo := (int32)(retC)

	return retGo
}

// Malloc is a wrapper around the C function g_malloc.
func Malloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Malloc0 is a wrapper around the C function g_malloc0.
func Malloc0(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_malloc0(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// Memdup is a wrapper around the C function g_memdup.
func Memdup(mem uintptr, byteSize uint32) uintptr {
	c_mem := (C.gconstpointer)(mem)

	c_byte_size := (C.guint)(byteSize)

	retC := C.g_memdup(c_mem, c_byte_size)
	retGo := (uintptr)(unsafe.Pointer(retC))

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

// NullifyPointer is a wrapper around the C function g_nullify_pointer.
func NullifyPointer(nullifyLocation uintptr) {
	c_nullify_location := (C.gpointer)(nullifyLocation)

	C.g_nullify_pointer(&c_nullify_location)

	return
}

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

// Unsupported : g_qsort_with_data : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func

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

// Realloc is a wrapper around the C function g_realloc.
func Realloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_realloc(c_mem, c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
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

// SetPrgname is a wrapper around the C function g_set_prgname.
func SetPrgname(prgname string) {
	c_prgname := C.CString(prgname)
	defer C.free(unsafe.Pointer(c_prgname))

	C.g_set_prgname(c_prgname)

	return
}

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

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

// SpacedPrimesClosest is a wrapper around the C function g_spaced_primes_closest.
func SpacedPrimesClosest(num uint32) uint32 {
	c_num := (C.guint)(num)

	retC := C.g_spaced_primes_closest(c_num)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_spawn_async : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_spawn_async_with_pipes : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

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

// StrEqual is a wrapper around the C function g_str_equal.
func StrEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_str_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// StrHash is a wrapper around the C function g_str_hash.
func StrHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_str_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

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

// TestLogTypeName is a wrapper around the C function g_test_log_type_name.
func TestLogTypeName(logType TestLogType) string {
	c_log_type := (C.GTestLogType)(logType)

	retC := C.g_test_log_type_name(c_log_type)
	retGo := C.GoString(retC)

	return retGo
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

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// TimeoutSourceNew is a wrapper around the C function g_timeout_source_new.
func TimeoutSourceNew(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TryMalloc is a wrapper around the C function g_try_malloc.
func TryMalloc(nBytes uint64) uintptr {
	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_malloc(c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TryRealloc is a wrapper around the C function g_try_realloc.
func TryRealloc(mem uintptr, nBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_bytes := (C.gsize)(nBytes)

	retC := C.g_try_realloc(c_mem, c_n_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// UnicharDigitValue is a wrapper around the C function g_unichar_digit_value.
func UnicharDigitValue(c rune) int32 {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_digit_value(c_c)
	retGo := (int32)(retC)

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

// UnicharIsxdigit is a wrapper around the C function g_unichar_isxdigit.
func UnicharIsxdigit(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_isxdigit(c_c)
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

// Blacklisted : g_unix_error_quark

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

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

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

// Blacklisted : GArray

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

// Pop is a wrapper around the C function g_async_queue_pop.
func (recv *AsyncQueue) Pop() uintptr {
	retC := C.g_async_queue_pop((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PopUnlocked is a wrapper around the C function g_async_queue_pop_unlocked.
func (recv *AsyncQueue) PopUnlocked() uintptr {
	retC := C.g_async_queue_pop_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Push is a wrapper around the C function g_async_queue_push.
func (recv *AsyncQueue) Push(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_async_queue_push((*C.GAsyncQueue)(recv.native), c_data)

	return
}

// PushUnlocked is a wrapper around the C function g_async_queue_push_unlocked.
func (recv *AsyncQueue) PushUnlocked(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_async_queue_push_unlocked((*C.GAsyncQueue)(recv.native), c_data)

	return
}

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

// TimedPop is a wrapper around the C function g_async_queue_timed_pop.
func (recv *AsyncQueue) TimedPop(endTime *TimeVal) uintptr {
	c_end_time := (*C.GTimeVal)(C.NULL)
	if endTime != nil {
		c_end_time = (*C.GTimeVal)(endTime.ToC())
	}

	retC := C.g_async_queue_timed_pop((*C.GAsyncQueue)(recv.native), c_end_time)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TimedPopUnlocked is a wrapper around the C function g_async_queue_timed_pop_unlocked.
func (recv *AsyncQueue) TimedPopUnlocked(endTime *TimeVal) uintptr {
	c_end_time := (*C.GTimeVal)(C.NULL)
	if endTime != nil {
		c_end_time = (*C.GTimeVal)(endTime.ToC())
	}

	retC := C.g_async_queue_timed_pop_unlocked((*C.GAsyncQueue)(recv.native), c_end_time)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TimeoutPop is a wrapper around the C function g_async_queue_timeout_pop.
func (recv *AsyncQueue) TimeoutPop(timeout uint64) uintptr {
	c_timeout := (C.guint64)(timeout)

	retC := C.g_async_queue_timeout_pop((*C.GAsyncQueue)(recv.native), c_timeout)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TimeoutPopUnlocked is a wrapper around the C function g_async_queue_timeout_pop_unlocked.
func (recv *AsyncQueue) TimeoutPopUnlocked(timeout uint64) uintptr {
	c_timeout := (C.guint64)(timeout)

	retC := C.g_async_queue_timeout_pop_unlocked((*C.GAsyncQueue)(recv.native), c_timeout)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TryPop is a wrapper around the C function g_async_queue_try_pop.
func (recv *AsyncQueue) TryPop() uintptr {
	retC := C.g_async_queue_try_pop((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TryPopUnlocked is a wrapper around the C function g_async_queue_try_pop_unlocked.
func (recv *AsyncQueue) TryPopUnlocked() uintptr {
	retC := C.g_async_queue_try_pop_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// Blacklisted : GByteArray

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

// Signal is a wrapper around the C function g_cond_signal.
func (recv *Cond) Signal() {
	C.g_cond_signal((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

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

// Blacklisted : g_date_to_struct_tm

// Valid is a wrapper around the C function g_date_valid.
func (recv *Date) Valid() bool {
	retC := C.g_date_valid((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
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

// HashTableDestroy is a wrapper around the C function g_hash_table_destroy.
func HashTableDestroy(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_destroy(c_hash_table)

	return
}

// g_hash_table_foreach : unsupported parameter func : no type generator for HFunc (GHFunc) for param func
// g_hash_table_foreach_remove : unsupported parameter func : no type generator for HRFunc (GHRFunc) for param func
// g_hash_table_foreach_steal : unsupported parameter func : no type generator for HRFunc (GHRFunc) for param func
// HashTableInsert is a wrapper around the C function g_hash_table_insert.
func HashTableInsert(hashTable *HashTable, key uintptr, value uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	retC := C.g_hash_table_insert(c_hash_table, c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// HashTableLookup is a wrapper around the C function g_hash_table_lookup.
func HashTableLookup(hashTable *HashTable, key uintptr) uintptr {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_lookup(c_hash_table, c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// HashTableLookupExtended is a wrapper around the C function g_hash_table_lookup_extended.
func HashTableLookupExtended(hashTable *HashTable, lookupKey uintptr) (bool, uintptr, uintptr) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_lookup_key := (C.gconstpointer)(lookupKey)

	var c_orig_key C.gpointer

	var c_value C.gpointer

	retC := C.g_hash_table_lookup_extended(c_hash_table, c_lookup_key, &c_orig_key, &c_value)
	retGo := retC == C.TRUE

	origKey := (uintptr)(unsafe.Pointer(&c_orig_key))

	value := (uintptr)(unsafe.Pointer(&c_value))

	return retGo, origKey, value
}

// g_hash_table_new : unsupported parameter hash_func : no type generator for HashFunc (GHashFunc) for param hash_func
// g_hash_table_new_full : unsupported parameter hash_func : no type generator for HashFunc (GHashFunc) for param hash_func
// HashTableRemove is a wrapper around the C function g_hash_table_remove.
func HashTableRemove(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_remove(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
}

// HashTableReplace is a wrapper around the C function g_hash_table_replace.
func HashTableReplace(hashTable *HashTable, key uintptr, value uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	retC := C.g_hash_table_replace(c_hash_table, c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

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

// HashTableSteal is a wrapper around the C function g_hash_table_steal.
func HashTableSteal(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_steal(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
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

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native *C.GHook
	Data   uintptr
	// next : record
	// prev : record
	RefCount uint32
	HookId   uint64
	Flags    uint32
	Func     uintptr
	// destroy : no type generator for DestroyNotify, GDestroyNotify
}

func HookNewFromC(u unsafe.Pointer) *Hook {
	c := (*C.GHook)(u)
	if c == nil {
		return nil
	}

	g := &Hook{
		Data:     (uintptr)(c.data),
		Flags:    (uint32)(c.flags),
		Func:     (uintptr)(c._func),
		HookId:   (uint64)(c.hook_id),
		RefCount: (uint32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *Hook) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)
	recv.native.ref_count =
		(C.guint)(recv.RefCount)
	recv.native.hook_id =
		(C.gulong)(recv.HookId)
	recv.native.flags =
		(C.guint)(recv.Flags)
	recv.native._func =
		(C.gpointer)(recv.Func)

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
// HookFindData is a wrapper around the C function g_hook_find_data.
func HookFindData(hookList *HookList, needValids bool, data uintptr) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_need_valids :=
		boolToGboolean(needValids)

	c_data := (C.gpointer)(data)

	retC := C.g_hook_find_data(c_hook_list, c_need_valids, c_data)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookFindFunc is a wrapper around the C function g_hook_find_func.
func HookFindFunc(hookList *HookList, needValids bool, func_ uintptr) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_need_valids :=
		boolToGboolean(needValids)

	c_func := (C.gpointer)(func_)

	retC := C.g_hook_find_func(c_hook_list, c_need_valids, c_func)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HookFindFuncData is a wrapper around the C function g_hook_find_func_data.
func HookFindFuncData(hookList *HookList, needValids bool, func_ uintptr, data uintptr) *Hook {
	c_hook_list := (*C.GHookList)(C.NULL)
	if hookList != nil {
		c_hook_list = (*C.GHookList)(hookList.ToC())
	}

	c_need_valids :=
		boolToGboolean(needValids)

	c_func := (C.gpointer)(func_)

	c_data := (C.gpointer)(data)

	retC := C.g_hook_find_func_data(c_hook_list, c_need_valids, c_func, c_data)
	retGo := HookNewFromC(unsafe.Pointer(retC))

	return retGo
}

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
	Dummy3 uintptr
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func HookListNewFromC(u unsafe.Pointer) *HookList {
	c := (*C.GHookList)(u)
	if c == nil {
		return nil
	}

	g := &HookList{
		Dummy3: (uintptr)(c.dummy3),
		SeqId:  (uint64)(c.seq_id),
		native: c,
	}

	return g
}

func (recv *HookList) ToC() unsafe.Pointer {
	recv.native.seq_id =
		(C.gulong)(recv.SeqId)
	recv.native.dummy3 =
		(C.gpointer)(recv.Dummy3)

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

// KeyFileErrorQuark is a wrapper around the C function g_key_file_error_quark.
func KeyFileErrorQuark() Quark {
	retC := C.g_key_file_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	Data   uintptr
	// next : record
	// prev : record
}

func ListNewFromC(u unsafe.Pointer) *List {
	c := (*C.GList)(u)
	if c == nil {
		return nil
	}

	g := &List{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *List) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

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

// ListAppend is a wrapper around the C function g_list_append.
func ListAppend(list *List, data uintptr) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_data := (C.gpointer)(data)

	retC := C.g_list_append(c_list, c_data)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// ListFind is a wrapper around the C function g_list_find.
func ListFind(list *List, data uintptr) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_list_find(c_list, c_data)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_find_custom : unsupported parameter func : no type generator for CompareFunc (GCompareFunc) for param func
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

// ListIndex is a wrapper around the C function g_list_index.
func ListIndex(list *List, data uintptr) int32 {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_list_index(c_list, c_data)
	retGo := (int32)(retC)

	return retGo
}

// ListInsert is a wrapper around the C function g_list_insert.
func ListInsert(list *List, data uintptr, position int32) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_data := (C.gpointer)(data)

	c_position := (C.gint)(position)

	retC := C.g_list_insert(c_list, c_data, c_position)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListInsertBefore is a wrapper around the C function g_list_insert_before.
func ListInsertBefore(list *List, sibling *List, data uintptr) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_sibling := (*C.GList)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GList)(sibling.ToC())
	}

	c_data := (C.gpointer)(data)

	retC := C.g_list_insert_before(c_list, c_sibling, c_data)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_list_insert_sorted : unsupported parameter func : no type generator for CompareFunc (GCompareFunc) for param func
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

// ListNthData is a wrapper around the C function g_list_nth_data.
func ListNthData(list *List, n uint32) uintptr {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_list_nth_data(c_list, c_n)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// ListPrepend is a wrapper around the C function g_list_prepend.
func ListPrepend(list *List, data uintptr) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_data := (C.gpointer)(data)

	retC := C.g_list_prepend(c_list, c_data)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListRemove is a wrapper around the C function g_list_remove.
func ListRemove(list *List, data uintptr) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_list_remove(c_list, c_data)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListRemoveAll is a wrapper around the C function g_list_remove_all.
func ListRemoveAll(list *List, data uintptr) *List {
	c_list := (*C.GList)(C.NULL)
	if list != nil {
		c_list = (*C.GList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_list_remove_all(c_list, c_data)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// FindSourceByFuncsUserData is a wrapper around the C function g_main_context_find_source_by_funcs_user_data.
func (recv *MainContext) FindSourceByFuncsUserData(funcs *SourceFuncs, userData uintptr) *Source {
	c_funcs := (*C.GSourceFuncs)(C.NULL)
	if funcs != nil {
		c_funcs = (*C.GSourceFuncs)(funcs.ToC())
	}

	c_user_data := (C.gpointer)(userData)

	retC := C.g_main_context_find_source_by_funcs_user_data((*C.GMainContext)(recv.native), c_funcs, c_user_data)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindSourceById is a wrapper around the C function g_main_context_find_source_by_id.
func (recv *MainContext) FindSourceById(sourceId uint32) *Source {
	c_source_id := (C.guint)(sourceId)

	retC := C.g_main_context_find_source_by_id((*C.GMainContext)(recv.native), c_source_id)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindSourceByUserData is a wrapper around the C function g_main_context_find_source_by_user_data.
func (recv *MainContext) FindSourceByUserData(userData uintptr) *Source {
	c_user_data := (C.gpointer)(userData)

	retC := C.g_main_context_find_source_by_user_data((*C.GMainContext)(recv.native), c_user_data)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_context_get_poll_func : no return generator

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

// Prepare is a wrapper around the C function g_main_context_prepare.
func (recv *MainContext) Prepare(priority int32) bool {
	c_priority := (C.gint)(priority)

	retC := C.g_main_context_prepare((*C.GMainContext)(recv.native), &c_priority)
	retGo := retC == C.TRUE

	return retGo
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

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

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

// GetPosition is a wrapper around the C function g_markup_parse_context_get_position.
func (recv *MarkupParseContext) GetPosition(lineNumber int32, charNumber int32) {
	c_line_number := (C.gint)(lineNumber)

	c_char_number := (C.gint)(charNumber)

	C.g_markup_parse_context_get_position((*C.GMarkupParseContext)(recv.native), &c_line_number, &c_char_number)

	return
}

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
	Data   uintptr
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

	g := &Node{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *Node) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Node with another Node, and returns true if they represent the same GObject.
func (recv *Node) Equals(other *Node) bool {
	return other.ToC() == recv.ToC()
}

// NodeNew is a wrapper around the C function g_node_new.
func NodeNew(data uintptr) *Node {
	c_data := (C.gpointer)(data)

	retC := C.g_node_new(c_data)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ChildIndex is a wrapper around the C function g_node_child_index.
func (recv *Node) ChildIndex(data uintptr) int32 {
	c_data := (C.gpointer)(data)

	retC := C.g_node_child_index((*C.GNode)(recv.native), c_data)
	retGo := (int32)(retC)

	return retGo
}

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

// Find is a wrapper around the C function g_node_find.
func (recv *Node) Find(order TraverseType, flags TraverseFlags, data uintptr) *Node {
	c_order := (C.GTraverseType)(order)

	c_flags := (C.GTraverseFlags)(flags)

	c_data := (C.gpointer)(data)

	retC := C.g_node_find((*C.GNode)(recv.native), c_order, c_flags, c_data)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FindChild is a wrapper around the C function g_node_find_child.
func (recv *Node) FindChild(flags TraverseFlags, data uintptr) *Node {
	c_flags := (C.GTraverseFlags)(flags)

	c_data := (C.gpointer)(data)

	retC := C.g_node_find_child((*C.GNode)(recv.native), c_flags, c_data)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// OptionEntry is a wrapper around the C record GOptionEntry.
type OptionEntry struct {
	native         *C.GOptionEntry
	LongName       string
	ShortName      rune
	Flags          int32
	Arg            OptionArg
	ArgData        uintptr
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
		ArgData:        (uintptr)(c.arg_data),
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
	recv.native.arg_data =
		(C.gpointer)(recv.ArgData)
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

// Get is a wrapper around the C function g_private_get.
func (recv *Private) Get() uintptr {
	retC := C.g_private_get((*C.GPrivate)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Set is a wrapper around the C function g_private_set.
func (recv *Private) Set(value uintptr) {
	c_value := (C.gpointer)(value)

	C.g_private_set((*C.GPrivate)(recv.native), c_value)

	return
}

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

// Free is a wrapper around the C function g_queue_free.
func (recv *Queue) Free() {
	C.g_queue_free((*C.GQueue)(recv.native))

	return
}

// IsEmpty is a wrapper around the C function g_queue_is_empty.
func (recv *Queue) IsEmpty() bool {
	retC := C.g_queue_is_empty((*C.GQueue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PeekHead is a wrapper around the C function g_queue_peek_head.
func (recv *Queue) PeekHead() uintptr {
	retC := C.g_queue_peek_head((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PeekTail is a wrapper around the C function g_queue_peek_tail.
func (recv *Queue) PeekTail() uintptr {
	retC := C.g_queue_peek_tail((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PopHead is a wrapper around the C function g_queue_pop_head.
func (recv *Queue) PopHead() uintptr {
	retC := C.g_queue_pop_head((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PopHeadLink is a wrapper around the C function g_queue_pop_head_link.
func (recv *Queue) PopHeadLink() *List {
	retC := C.g_queue_pop_head_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopTail is a wrapper around the C function g_queue_pop_tail.
func (recv *Queue) PopTail() uintptr {
	retC := C.g_queue_pop_tail((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PopTailLink is a wrapper around the C function g_queue_pop_tail_link.
func (recv *Queue) PopTailLink() *List {
	retC := C.g_queue_pop_tail_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PushHead is a wrapper around the C function g_queue_push_head.
func (recv *Queue) PushHead(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_queue_push_head((*C.GQueue)(recv.native), c_data)

	return
}

// PushHeadLink is a wrapper around the C function g_queue_push_head_link.
func (recv *Queue) PushHeadLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_head_link((*C.GQueue)(recv.native), c_link_)

	return
}

// PushTail is a wrapper around the C function g_queue_push_tail.
func (recv *Queue) PushTail(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_queue_push_tail((*C.GQueue)(recv.native), c_data)

	return
}

// PushTailLink is a wrapper around the C function g_queue_push_tail_link.
func (recv *Queue) PushTailLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_tail_link((*C.GQueue)(recv.native), c_link_)

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

// SList is a wrapper around the C record GSList.
type SList struct {
	native *C.GSList
	Data   uintptr
	// next : record
}

func SListNewFromC(u unsafe.Pointer) *SList {
	c := (*C.GSList)(u)
	if c == nil {
		return nil
	}

	g := &SList{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *SList) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

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

// SListAppend is a wrapper around the C function g_slist_append.
func SListAppend(list *SList, data uintptr) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_data := (C.gpointer)(data)

	retC := C.g_slist_append(c_list, c_data)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// SListFind is a wrapper around the C function g_slist_find.
func SListFind(list *SList, data uintptr) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_slist_find(c_list, c_data)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_find_custom : unsupported parameter func : no type generator for CompareFunc (GCompareFunc) for param func
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

// SListIndex is a wrapper around the C function g_slist_index.
func SListIndex(list *SList, data uintptr) int32 {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_slist_index(c_list, c_data)
	retGo := (int32)(retC)

	return retGo
}

// SListInsert is a wrapper around the C function g_slist_insert.
func SListInsert(list *SList, data uintptr, position int32) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_data := (C.gpointer)(data)

	c_position := (C.gint)(position)

	retC := C.g_slist_insert(c_list, c_data, c_position)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListInsertBefore is a wrapper around the C function g_slist_insert_before.
func SListInsertBefore(slist *SList, sibling *SList, data uintptr) *SList {
	c_slist := (*C.GSList)(C.NULL)
	if slist != nil {
		c_slist = (*C.GSList)(slist.ToC())
	}

	c_sibling := (*C.GSList)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GSList)(sibling.ToC())
	}

	c_data := (C.gpointer)(data)

	retC := C.g_slist_insert_before(c_slist, c_sibling, c_data)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_insert_sorted : unsupported parameter func : no type generator for CompareFunc (GCompareFunc) for param func
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

// SListNthData is a wrapper around the C function g_slist_nth_data.
func SListNthData(list *SList, n uint32) uintptr {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_n := (C.guint)(n)

	retC := C.g_slist_nth_data(c_list, c_n)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// SListPrepend is a wrapper around the C function g_slist_prepend.
func SListPrepend(list *SList, data uintptr) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_data := (C.gpointer)(data)

	retC := C.g_slist_prepend(c_list, c_data)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListRemove is a wrapper around the C function g_slist_remove.
func SListRemove(list *SList, data uintptr) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_slist_remove(c_list, c_data)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SListRemoveAll is a wrapper around the C function g_slist_remove_all.
func SListRemoveAll(list *SList, data uintptr) *SList {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	c_data := (C.gconstpointer)(data)

	retC := C.g_slist_remove_all(c_list, c_data)
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

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
	native         *C.GScanner
	UserData       uintptr
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
		UserData:       (uintptr)(c.user_data),
		native:         c,
	}

	return g
}

func (recv *Scanner) ToC() unsafe.Pointer {
	recv.native.user_data =
		(C.gpointer)(recv.UserData)
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

// LookupSymbol is a wrapper around the C function g_scanner_lookup_symbol.
func (recv *Scanner) LookupSymbol(symbol string) uintptr {
	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	retC := C.g_scanner_lookup_symbol((*C.GScanner)(recv.native), c_symbol)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PeekNextToken is a wrapper around the C function g_scanner_peek_next_token.
func (recv *Scanner) PeekNextToken() TokenType {
	retC := C.g_scanner_peek_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// ScopeAddSymbol is a wrapper around the C function g_scanner_scope_add_symbol.
func (recv *Scanner) ScopeAddSymbol(scopeId uint32, symbol string, value uintptr) {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	c_value := (C.gpointer)(value)

	C.g_scanner_scope_add_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol, c_value)

	return
}

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc (GHFunc) for param func

// ScopeLookupSymbol is a wrapper around the C function g_scanner_scope_lookup_symbol.
func (recv *Scanner) ScopeLookupSymbol(scopeId uint32, symbol string) uintptr {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	retC := C.g_scanner_scope_lookup_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// SourceRemoveByFuncsUserData is a wrapper around the C function g_source_remove_by_funcs_user_data.
func SourceRemoveByFuncsUserData(funcs *SourceFuncs, userData uintptr) bool {
	c_funcs := (*C.GSourceFuncs)(C.NULL)
	if funcs != nil {
		c_funcs = (*C.GSourceFuncs)(funcs.ToC())
	}

	c_user_data := (C.gpointer)(userData)

	retC := C.g_source_remove_by_funcs_user_data(c_funcs, c_user_data)
	retGo := retC == C.TRUE

	return retGo
}

// SourceRemoveByUserData is a wrapper around the C function g_source_remove_by_user_data.
func SourceRemoveByUserData(userData uintptr) bool {
	c_user_data := (C.gpointer)(userData)

	retC := C.g_source_remove_by_user_data(c_user_data)
	retGo := retC == C.TRUE

	return retGo
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

// Ref is a wrapper around the C function g_source_ref.
func (recv *Source) Ref() *Source {
	retC := C.g_source_ref((*C.GSource)(recv.native))
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
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

// SetCallbackIndirect is a wrapper around the C function g_source_set_callback_indirect.
func (recv *Source) SetCallbackIndirect(callbackData uintptr, callbackFuncs *SourceCallbackFuncs) {
	c_callback_data := (C.gpointer)(callbackData)

	c_callback_funcs := (*C.GSourceCallbackFuncs)(C.NULL)
	if callbackFuncs != nil {
		c_callback_funcs = (*C.GSourceCallbackFuncs)(callbackFuncs.ToC())
	}

	C.g_source_set_callback_indirect((*C.GSource)(recv.native), c_callback_data, c_callback_funcs)

	return
}

// SetCanRecurse is a wrapper around the C function g_source_set_can_recurse.
func (recv *Source) SetCanRecurse(canRecurse bool) {
	c_can_recurse :=
		boolToGboolean(canRecurse)

	C.g_source_set_can_recurse((*C.GSource)(recv.native), c_can_recurse)

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

// ThreadErrorQuark is a wrapper around the C function g_thread_error_quark.
func ThreadErrorQuark() Quark {
	retC := C.g_thread_error_quark()
	retGo := (Quark)(retC)

	return retGo
}

// ThreadExit is a wrapper around the C function g_thread_exit.
func ThreadExit(retval uintptr) {
	c_retval := (C.gpointer)(retval)

	C.g_thread_exit(c_retval)

	return
}

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

// Join is a wrapper around the C function g_thread_join.
func (recv *Thread) Join() uintptr {
	retC := C.g_thread_join((*C.GThread)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	UserData  uintptr
	Exclusive bool
}

func ThreadPoolNewFromC(u unsafe.Pointer) *ThreadPool {
	c := (*C.GThreadPool)(u)
	if c == nil {
		return nil
	}

	g := &ThreadPool{
		Exclusive: c.exclusive == C.TRUE,
		UserData:  (uintptr)(c.user_data),
		native:    c,
	}

	return g
}

func (recv *ThreadPool) ToC() unsafe.Pointer {
	recv.native.user_data =
		(C.gpointer)(recv.UserData)
	recv.native.exclusive =
		boolToGboolean(recv.Exclusive)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThreadPool with another ThreadPool, and returns true if they represent the same GObject.
func (recv *ThreadPool) Equals(other *ThreadPool) bool {
	return other.ToC() == recv.ToC()
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

// Push is a wrapper around the C function g_thread_pool_push.
func (recv *ThreadPool) Push(data uintptr) (bool, error) {
	c_data := (C.gpointer)(data)

	var cThrowableError *C.GError

	retC := C.g_thread_pool_push((*C.GThreadPool)(recv.native), c_data, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

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

// Add is a wrapper around the C function g_time_val_add.
func (recv *TimeVal) Add(microseconds int64) {
	c_microseconds := (C.glong)(microseconds)

	C.g_time_val_add((*C.GTimeVal)(recv.native), c_microseconds)

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

// TrashStackPeek is a wrapper around the C function g_trash_stack_peek.
func TrashStackPeek(stackP *TrashStack) uintptr {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	retC := C.g_trash_stack_peek(c_stack_p)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TrashStackPop is a wrapper around the C function g_trash_stack_pop.
func TrashStackPop(stackP *TrashStack) uintptr {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	retC := C.g_trash_stack_pop(c_stack_p)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// TrashStackPush is a wrapper around the C function g_trash_stack_push.
func TrashStackPush(stackP *TrashStack, dataP uintptr) {
	c_stack_p := (**C.GTrashStack)(C.NULL)
	if stackP != nil {
		c_stack_p = (**C.GTrashStack)(stackP.ToC())
	}

	c_data_p := (C.gpointer)(dataP)

	C.g_trash_stack_push(c_stack_p, c_data_p)

	return
}

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

// Insert is a wrapper around the C function g_tree_insert.
func (recv *Tree) Insert(key uintptr, value uintptr) {
	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	C.g_tree_insert((*C.GTree)(recv.native), c_key, c_value)

	return
}

// Lookup is a wrapper around the C function g_tree_lookup.
func (recv *Tree) Lookup(key uintptr) uintptr {
	c_key := (C.gconstpointer)(key)

	retC := C.g_tree_lookup((*C.GTree)(recv.native), c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// LookupExtended is a wrapper around the C function g_tree_lookup_extended.
func (recv *Tree) LookupExtended(lookupKey uintptr, origKey uintptr, value uintptr) bool {
	c_lookup_key := (C.gconstpointer)(lookupKey)

	c_orig_key := (C.gpointer)(origKey)

	c_value := (C.gpointer)(value)

	retC := C.g_tree_lookup_extended((*C.GTree)(recv.native), c_lookup_key, &c_orig_key, &c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Nnodes is a wrapper around the C function g_tree_nnodes.
func (recv *Tree) Nnodes() int32 {
	retC := C.g_tree_nnodes((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Remove is a wrapper around the C function g_tree_remove.
func (recv *Tree) Remove(key uintptr) bool {
	c_key := (C.gconstpointer)(key)

	retC := C.g_tree_remove((*C.GTree)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Replace is a wrapper around the C function g_tree_replace.
func (recv *Tree) Replace(key uintptr, value uintptr) {
	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	C.g_tree_replace((*C.GTree)(recv.native), c_key, c_value)

	return
}

// Unsupported : g_tree_search : unsupported parameter search_func : no type generator for CompareFunc (GCompareFunc) for param search_func

// Steal is a wrapper around the C function g_tree_steal.
func (recv *Tree) Steal(key uintptr) bool {
	c_key := (C.gconstpointer)(key)

	retC := C.g_tree_steal((*C.GTree)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_tree_traverse : unsupported parameter traverse_func : no type generator for TraverseFunc (GTraverseFunc) for param traverse_func

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

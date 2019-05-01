// This is a generated file - DO NOT EDIT

package glib

import call "github.com/pekim/gobbi/lib/internal/call"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
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

type AsciiType int

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

type GFileTest int

const (
	FILE_TEST_IS_REGULAR    GFileTest = 1
	FILE_TEST_IS_SYMLINK    GFileTest = 2
	FILE_TEST_IS_DIR        GFileTest = 4
	FILE_TEST_IS_EXECUTABLE GFileTest = 8
	FILE_TEST_EXISTS        GFileTest = 16
)

type FormatSizeFlags int

const (
	FORMAT_SIZE_DEFAULT     FormatSizeFlags = 0
	FORMAT_SIZE_LONG_FORMAT FormatSizeFlags = 1
	FORMAT_SIZE_IEC_UNITS   FormatSizeFlags = 2
	FORMAT_SIZE_BITS        FormatSizeFlags = 4
)

type HookFlagMask int

const (
	HOOK_FLAG_ACTIVE  HookFlagMask = 1
	HOOK_FLAG_IN_CALL HookFlagMask = 2
	HOOK_FLAG_MASK    HookFlagMask = 15
)

type IOCondition int

const (
	IO_IN   IOCondition = 1
	IO_OUT  IOCondition = 4
	IO_PRI  IOCondition = 2
	IO_ERR  IOCondition = 8
	IO_HUP  IOCondition = 16
	IO_NVAL IOCondition = 32
)

type IOFlags int

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

type KeyFileFlags int

const (
	KEY_FILE_NONE              KeyFileFlags = 0
	KEY_FILE_KEEP_COMMENTS     KeyFileFlags = 1
	KEY_FILE_KEEP_TRANSLATIONS KeyFileFlags = 2
)

type LogLevelFlags int

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

type MarkupCollectType int

const (
	MARKUP_COLLECT_INVALID  MarkupCollectType = 0
	MARKUP_COLLECT_STRING   MarkupCollectType = 1
	MARKUP_COLLECT_STRDUP   MarkupCollectType = 2
	MARKUP_COLLECT_BOOLEAN  MarkupCollectType = 3
	MARKUP_COLLECT_TRISTATE MarkupCollectType = 4
	MARKUP_COLLECT_OPTIONAL MarkupCollectType = 65536
)

type MarkupParseFlags int

const (
	MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG MarkupParseFlags = 1
	MARKUP_TREAT_CDATA_AS_TEXT              MarkupParseFlags = 2
	MARKUP_PREFIX_ERROR_POSITION            MarkupParseFlags = 4
	MARKUP_IGNORE_QUALIFIED                 MarkupParseFlags = 8
)

type OptionFlags int

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

type SpawnFlags int

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

type TestSubprocessFlags int

const (
	TEST_SUBPROCESS_INHERIT_STDIN  TestSubprocessFlags = 1
	TEST_SUBPROCESS_INHERIT_STDOUT TestSubprocessFlags = 2
	TEST_SUBPROCESS_INHERIT_STDERR TestSubprocessFlags = 4
)

type TestTrapFlags int

const (
	TEST_TRAP_SILENCE_STDOUT TestTrapFlags = 128
	TEST_TRAP_SILENCE_STDERR TestTrapFlags = 256
	TEST_TRAP_INHERIT_STDIN  TestTrapFlags = 512
)

type TraverseFlags int

const (
	TRAVERSE_LEAVES     TraverseFlags = 1
	TRAVERSE_NON_LEAVES TraverseFlags = 2
	TRAVERSE_ALL        TraverseFlags = 3
	TRAVERSE_MASK       TraverseFlags = 3
	TRAVERSE_LEAFS      TraverseFlags = 1
	TRAVERSE_NON_LEAFS  TraverseFlags = 2
)
const ANALYZER_ANALYZING int32 = 1
const ASCII_DTOSTR_BUF_SIZE int32 = 39
const BIG_ENDIAN int32 = 4321
const CSET_A_2_Z string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const CSET_DIGITS string = "0123456789"
const CSET_a_2_z string = "abcdefghijklmnopqrstuvwxyz"
const DATALIST_FLAGS_MASK int32 = 3
const DATE_BAD_DAY int32 = 0
const DATE_BAD_JULIAN int32 = 0
const DATE_BAD_YEAR int32 = 0
const DIR_SEPARATOR int32 = 47
const DIR_SEPARATOR_S string = "/"
const E float64 = 0
const GINT16_FORMAT string = "hi"
const GINT32_FORMAT string = "i"
const GINT64_FORMAT string = "li"
const GNUC_FUNCTION string = ""
const GNUC_PRETTY_FUNCTION string = ""
const GUINT16_FORMAT string = "hu"
const GUINT32_FORMAT string = "u"
const GUINT64_FORMAT string = "lu"
const HAVE_GINT64 int32 = 1
const HAVE_GNUC_VARARGS int32 = 1
const HAVE_GNUC_VISIBILITY int32 = 1
const HAVE_GROWING_STACK int32 = 0
const HAVE_ISO_VARARGS int32 = 1
const HOOK_FLAG_USER_SHIFT int32 = 4
const IEEE754_DOUBLE_BIAS int32 = 1023
const IEEE754_FLOAT_BIAS int32 = 127

// Blacklisted : KEY_FILE_DESKTOP_ACTION_GROUP_PREFIX

const KEY_FILE_DESKTOP_KEY_FULLNAME string = "X-GNOME-FullName"
const KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN string = "X-GNOME-Gettext-Domain"
const KEY_FILE_DESKTOP_KEY_KEYWORDS string = "Keywords"
const LITTLE_ENDIAN int32 = 1234
const LN10 float64 = 0
const LN2 float64 = 0
const LOG_2_BASE_10 float64 = 0

// Blacklisted : LOG_DOMAIN

const LOG_FATAL_MASK int32 = 5
const LOG_LEVEL_USER_SHIFT int32 = 8
const MAJOR_VERSION int32 = 2
const MAXINT64 int64 = 9223372036854775807
const MAXUINT64 uint64 = 9223372036854775807
const MICRO_VERSION int32 = 1
const MININT64 int64 = -9223372036854775808
const MINOR_VERSION int32 = 56
const MODULE_SUFFIX string = "so"
const PDP_ENDIAN int32 = 3412
const PI float64 = 0
const PI_2 float64 = 0
const PI_4 float64 = 0
const POLLFD_FORMAT string = "%d"
const PRIORITY_DEFAULT int32 = 0
const PRIORITY_DEFAULT_IDLE int32 = 200
const PRIORITY_HIGH int32 = -100
const PRIORITY_HIGH_IDLE int32 = 100
const PRIORITY_LOW int32 = 300
const SEARCHPATH_SEPARATOR int32 = 58
const SEARCHPATH_SEPARATOR_S string = ":"
const SIZEOF_LONG int32 = 8
const SIZEOF_SIZE_T int32 = 8
const SIZEOF_SSIZE_T int32 = 8
const SIZEOF_VOID_P int32 = 8
const SQRT2 float64 = 0
const STR_DELIMITERS string = "_-|> <."
const SYSDEF_AF_INET int32 = 2
const SYSDEF_AF_INET6 int32 = 10
const SYSDEF_AF_UNIX int32 = 1
const SYSDEF_MSG_DONTROUTE int32 = 4
const SYSDEF_MSG_OOB int32 = 1
const SYSDEF_MSG_PEEK int32 = 2
const URI_RESERVED_CHARS_GENERIC_DELIMITERS string = ":/?#[]@"
const URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS string = "!$&'()*+,;="
const USEC_PER_SEC int32 = 1000000
const VA_COPY_AS_ARRAY int32 = 1

// Blacklisted : WIN32_MSG_HANDLE

type BookmarkFileError int

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

type ConvertError int

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

type DateDMY int

const (
	DATE_DAY   DateDMY = 0
	DATE_MONTH DateDMY = 1
	DATE_YEAR  DateDMY = 2
)

type DateMonth int

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

type DateWeekday int

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

type ErrorType int

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

type FileError int

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

type IOChannelError int

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

type IOError int

const (
	IO_ERROR_NONE    IOError = 0
	IO_ERROR_AGAIN   IOError = 1
	IO_ERROR_INVAL   IOError = 2
	IO_ERROR_UNKNOWN IOError = 3
)

type IOStatus int

const (
	IO_STATUS_ERROR  IOStatus = 0
	IO_STATUS_NORMAL IOStatus = 1
	IO_STATUS_EOF    IOStatus = 2
	IO_STATUS_AGAIN  IOStatus = 3
)

type KeyFileError int

const (
	KEY_FILE_ERROR_UNKNOWN_ENCODING KeyFileError = 0
	KEY_FILE_ERROR_PARSE            KeyFileError = 1
	KEY_FILE_ERROR_NOT_FOUND        KeyFileError = 2
	KEY_FILE_ERROR_KEY_NOT_FOUND    KeyFileError = 3
	KEY_FILE_ERROR_GROUP_NOT_FOUND  KeyFileError = 4
	KEY_FILE_ERROR_INVALID_VALUE    KeyFileError = 5
)

type MarkupError int

const (
	MARKUP_ERROR_BAD_UTF8          MarkupError = 0
	MARKUP_ERROR_EMPTY             MarkupError = 1
	MARKUP_ERROR_PARSE             MarkupError = 2
	MARKUP_ERROR_UNKNOWN_ELEMENT   MarkupError = 3
	MARKUP_ERROR_UNKNOWN_ATTRIBUTE MarkupError = 4
	MARKUP_ERROR_INVALID_CONTENT   MarkupError = 5
	MARKUP_ERROR_MISSING_ATTRIBUTE MarkupError = 6
)

type NormalizeMode int

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

type OptionArg int

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

type OptionError int

const (
	OPTION_ERROR_UNKNOWN_OPTION OptionError = 0
	OPTION_ERROR_BAD_VALUE      OptionError = 1
	OPTION_ERROR_FAILED         OptionError = 2
)

type SeekType int

const (
	SEEK_CUR SeekType = 0
	SEEK_SET SeekType = 1
	SEEK_END SeekType = 2
)

type ShellError int

const (
	SHELL_ERROR_BAD_QUOTING  ShellError = 0
	SHELL_ERROR_EMPTY_STRING ShellError = 1
	SHELL_ERROR_FAILED       ShellError = 2
)

type SliceConfig int

const (
	SLICE_CONFIG_ALWAYS_MALLOC      SliceConfig = 1
	SLICE_CONFIG_BYPASS_MAGAZINES   SliceConfig = 2
	SLICE_CONFIG_WORKING_SET_MSECS  SliceConfig = 3
	SLICE_CONFIG_COLOR_INCREMENT    SliceConfig = 4
	SLICE_CONFIG_CHUNK_SIZES        SliceConfig = 5
	SLICE_CONFIG_CONTENTION_COUNTER SliceConfig = 6
)

type SpawnError int

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

type TestLogType int

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

type ThreadError int

const (
	THREAD_ERROR_AGAIN ThreadError = 0
)

type TimeType int

const (
	TIME_TYPE_STANDARD  TimeType = 0
	TIME_TYPE_DAYLIGHT  TimeType = 1
	TIME_TYPE_UNIVERSAL TimeType = 2
)

type TokenType int

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

type TraverseType int

const (
	IN_ORDER    TraverseType = 0
	PRE_ORDER   TraverseType = 1
	POST_ORDER  TraverseType = 2
	LEVEL_ORDER TraverseType = 3
)

type UnicodeBreakType int

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

type UnicodeScript int

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

type UnicodeType int

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

type VariantParseError int

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

// Unsupported : g_ascii_dtostr : return type :

// Unsupported : g_ascii_formatd : return type :

// Unsupported : g_ascii_strdown : return type :

// Unsupported : g_ascii_strup : return type :

// Unsupported : g_assertion_message_cmpnum : unsupported parameter arg1 : no type generator for long double (long double) for param arg1

// Unsupported : g_atexit : unsupported parameter func : no type generator for VoidFunc (GVoidFunc) for param func

// Unsupported : g_basename : return type :

// Unsupported : g_bookmark_file_error_quark : return type :

// Unsupported : g_build_filename : unsupported parameter ... : varargs

// Unsupported : g_build_path : unsupported parameter ... : varargs

// Blacklisted : g_byte_array_free

// Unsupported : g_byte_array_new : array return type :

// Unsupported : throws

// Unsupported : g_convert : array return type :

// Unsupported : g_convert_error_quark : return type :

// Unsupported : g_convert_with_fallback : array return type :

// Unsupported : g_convert_with_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_datalist_foreach : unsupported parameter func : no type generator for DataForeachFunc (GDataForeachFunc) for param func

// Unsupported : g_datalist_get_data : no return generator

// Unsupported : g_datalist_id_get_data : no return generator

// Unsupported : g_datalist_id_remove_no_notify : no return generator

// Unsupported : g_datalist_id_set_data_full : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_dataset_destroy : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_foreach : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_get_data : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_remove_no_notify : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_dataset_id_set_data_full : unsupported parameter dataset_location : no type generator for gpointer (gconstpointer) for param dataset_location

// Unsupported : g_date_is_leap_year : return type :

// Unsupported : g_date_valid_day : return type :

// Unsupported : g_date_valid_dmy : return type :

// Unsupported : g_date_valid_julian : return type :

// Unsupported : g_date_valid_month : return type :

// Unsupported : g_date_valid_weekday : return type :

// Unsupported : g_date_valid_year : return type :

// Unsupported : g_direct_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_direct_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_file_error_from_errno : return type :

// Unsupported : g_file_error_quark : return type :

// Unsupported : g_file_get_contents : unsupported parameter contents : output array param contents

// Unsupported : throws

// Unsupported : g_file_test : return type :

// Unsupported : g_filename_from_uri : return type :

// Unsupported : g_filename_from_utf8 : return type :

// Unsupported : g_filename_to_uri : return type :

// Unsupported : g_filename_to_utf8 : return type :

// Unsupported : g_find_program_in_path : return type :

// Unsupported : g_free : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_get_charset : return type :

// Unsupported : g_get_codeset : return type :

// Unsupported : g_get_current_dir : return type :

// Unsupported : g_get_home_dir : return type :

// Unsupported : g_get_prgname : return type :

// Unsupported : g_get_real_name : return type :

// Unsupported : g_get_tmp_dir : return type :

// Unsupported : g_get_user_name : return type :

// Unsupported : g_getenv : return type :

// Unsupported : g_hash_table_insert : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_lookup : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_lookup_extended : unsupported parameter lookup_key : no type generator for gpointer (gconstpointer) for param lookup_key

// Unsupported : g_hash_table_remove : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hash_table_replace : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_steal : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

// Unsupported : g_hook_destroy : return type :

// Unsupported : g_iconv : unsupported parameter converter : Blacklisted record : GIConv

// Unsupported : g_iconv_open : return type :

// Unsupported : g_idle_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_idle_remove_by_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_idle_source_new : return type :

// Unsupported : g_int_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_int_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_io_add_watch : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// Unsupported : g_io_add_watch_full : unsupported parameter func : no type generator for IOFunc (GIOFunc) for param func

// Unsupported : g_io_channel_error_from_errno : return type :

// Unsupported : g_io_channel_error_quark : return type :

// Unsupported : g_io_create_watch : return type :

// Unsupported : g_key_file_error_quark : return type :

// Unsupported : g_locale_from_utf8 : array return type :

// Unsupported : g_locale_to_utf8 : return type :

// Unsupported : g_log_default_handler : unsupported parameter unused_data : no type generator for gpointer (gpointer) for param unused_data

// Unsupported : g_log_set_always_fatal : return type :

// Unsupported : g_log_set_fatal_mask : return type :

// Unsupported : g_log_set_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Blacklisted : g_log_structured_standard

// Unsupported : g_logv : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_main_context_default : return type :

// MainDepth is a wrapper around the C function g_main_depth.
func MainDepth() int32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_INT},
	}
	call.Function(2018, &data)
	ret := data.Return.Int32()

	return ret
}

// Unsupported : g_malloc : no return generator

// Unsupported : g_malloc0 : no return generator

// Unsupported : g_markup_error_quark : return type :

// Unsupported : g_markup_escape_text : return type :

// Unsupported : g_mem_is_system_malloc : return type :

// MemProfile is a wrapper around the C function g_mem_profile.
func MemProfile() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(2071, &data)
	return
}

// Unsupported : g_memdup : unsupported parameter mem : no type generator for gpointer (gconstpointer) for param mem

// Unsupported : g_nullify_pointer : unsupported parameter nullify_location : no type generator for gpointer (gpointer*) for param nullify_location

// Unsupported : g_number_parser_error_quark : return type :

// Unsupported : g_option_error_quark : return type :

// Unsupported : g_parse_debug_string : unsupported parameter keys :

// Unsupported : g_path_get_basename : return type :

// Unsupported : g_path_get_dirname : return type :

// Unsupported : g_path_is_absolute : return type :

// Unsupported : g_path_skip_root : return type :

// Unsupported : g_pattern_match : return type :

// Unsupported : g_pattern_match_simple : return type :

// Unsupported : g_pattern_match_string : return type :

// Unsupported : g_printf_string_upper_bound : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_qsort_with_data : unsupported parameter pbase : no type generator for gpointer (gconstpointer) for param pbase

// Unsupported : g_quark_from_static_string : return type :

// Unsupported : g_quark_from_string : return type :

// Unsupported : g_quark_to_string : return type :

// Unsupported : g_quark_try_string : return type :

// RandomDouble is a wrapper around the C function g_random_double.
func RandomDouble() float64 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_DOUBLE},
	}
	call.Function(2567, &data)
	ret := data.Return.Float64()

	return ret
}

// RandomInt is a wrapper around the C function g_random_int.
func RandomInt() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(2569, &data)
	ret := data.Return.Uint32()

	return ret
}

// Unsupported : g_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_regex_error_quark : return type :

// Unsupported : g_set_print_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_set_printerr_handler : unsupported parameter func : no type generator for PrintFunc (GPrintFunc) for param func

// Unsupported : g_shell_error_quark : return type :

// Unsupported : g_shell_parse_argv : unsupported parameter argcp : array length param argcp is pointer (gint*)

// Unsupported : g_shell_quote : return type :

// Unsupported : g_shell_unquote : return type :

// Blacklisted : g_slice_get_config_state

// Unsupported : g_source_remove : return type :

// Unsupported : g_source_remove_by_funcs_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_source_remove_by_user_data : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_spawn_async : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_spawn_async_with_pipes : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_spawn_command_line_async : return type :

// Unsupported : g_spawn_command_line_sync : unsupported parameter standard_output : output array param standard_output

// Unsupported : g_spawn_error_quark : return type :

// Unsupported : g_spawn_exit_error_quark : return type :

// Unsupported : g_spawn_sync : unsupported parameter child_setup : no type generator for SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Unsupported : g_stpcpy : return type :

// Unsupported : g_str_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_str_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// Unsupported : g_strcanon : return type :

// Unsupported : g_strchomp : return type :

// Unsupported : g_strchug : return type :

// Unsupported : g_strcompress : return type :

// Unsupported : g_strconcat : unsupported parameter ... : varargs

// Unsupported : g_strdelimit : return type :

// Unsupported : g_strdown : return type :

// Unsupported : g_strdup : return type :

// Unsupported : g_strdup_printf : return type :

// Unsupported : g_strdup_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_strdupv : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strerror : return type :

// Unsupported : g_strescape : return type :

// Unsupported : g_strfreev : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_string_new : return type :

// Unsupported : g_string_new_len : return type :

// Unsupported : g_string_sized_new : return type :

// Unsupported : g_strjoin : unsupported parameter ... : varargs

// Unsupported : g_strjoinv : unsupported parameter str_array : in string with indirection level of 2

// Unsupported : g_strndup : return type :

// Unsupported : g_strnfill : return type :

// Unsupported : g_strreverse : return type :

// Unsupported : g_strrstr : return type :

// Unsupported : g_strrstr_len : return type :

// Unsupported : g_strsignal : return type :

// Unsupported : g_strsplit : array return type :

// Unsupported : g_strstr_len : return type :

// Unsupported : g_strup : return type :

// Unsupported : g_strv_get_type : return type :

// Unsupported : g_test_log_type_name : return type :

// Unsupported : g_thread_error_quark : return type :

// Unsupported : g_thread_exit : unsupported parameter retval : no type generator for gpointer (gpointer) for param retval

// Unsupported : g_thread_self : return type :

// Unsupported : g_timeout_add : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_add_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_timeout_source_new : return type :

// Unsupported : g_trash_stack_peek : no return generator

// Unsupported : g_trash_stack_pop : no return generator

// Unsupported : g_trash_stack_push : unsupported parameter data_p : no type generator for gpointer (gpointer) for param data_p

// Unsupported : g_try_malloc : no return generator

// Unsupported : g_try_realloc : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : throws

// Unsupported : g_ucs4_to_utf8 : return type :

// Unsupported : g_unichar_break_type : return type :

// Unsupported : g_unichar_isalnum : return type :

// Unsupported : g_unichar_isalpha : return type :

// Unsupported : g_unichar_iscntrl : return type :

// Unsupported : g_unichar_isdefined : return type :

// Unsupported : g_unichar_isdigit : return type :

// Unsupported : g_unichar_isgraph : return type :

// Unsupported : g_unichar_islower : return type :

// Unsupported : g_unichar_isprint : return type :

// Unsupported : g_unichar_ispunct : return type :

// Unsupported : g_unichar_isspace : return type :

// Unsupported : g_unichar_istitle : return type :

// Unsupported : g_unichar_isupper : return type :

// Unsupported : g_unichar_iswide : return type :

// Unsupported : g_unichar_isxdigit : return type :

// Blacklisted : g_unichar_to_utf8

// Unsupported : g_unichar_type : return type :

// Unsupported : g_unichar_validate : return type :

// Blacklisted : g_unicode_canonical_decomposition

// Unsupported : g_unix_error_quark : return type :

// Unsupported : throws

// Unsupported : g_utf16_to_utf8 : return type :

// Unsupported : g_utf8_casefold : return type :

// Unsupported : g_utf8_collate_key : return type :

// Unsupported : g_utf8_find_next_char : return type :

// Unsupported : g_utf8_find_prev_char : return type :

// Unsupported : g_utf8_normalize : return type :

// Unsupported : g_utf8_offset_to_pointer : return type :

// Unsupported : g_utf8_prev_char : return type :

// Unsupported : g_utf8_strchr : return type :

// Unsupported : g_utf8_strdown : return type :

// Unsupported : g_utf8_strncpy : return type :

// Unsupported : g_utf8_strrchr : return type :

// Unsupported : g_utf8_strup : return type :

// Unsupported : throws

// Blacklisted : g_utf8_to_ucs4_fast

// Unsupported : throws

// Unsupported : g_utf8_validate : return type :

// Unsupported : g_variant_get_gtype : return type :

// Unsupported : g_variant_parse : unsupported parameter endptr : in string with indirection level of 2

// Unsupported : g_variant_parse_error_quark : return type :

// Unsupported : g_variant_parser_get_error_quark : return type :

// Unsupported : g_variant_type_checked_ : return type :

// Unsupported : g_variant_type_string_is_valid : return type :

// Unsupported : g_vsnprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : GByteArray

// Unsupported : g_date_new : return type :

// Unsupported : g_date_new_dmy : return type :

// Unsupported : g_date_new_julian : return type :

// Unsupported : g_error_new : return type :

// Unsupported : g_error_new_literal : return type :

// Blacklisted : GIConv

// Unsupported : g_io_channel_new_file : return type :

// Unsupported : g_io_channel_unix_new : return type :

// Unsupported : g_main_context_new : return type :

// Unsupported : g_main_loop_new : return type :

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Blacklisted : GPtrArray

// Unsupported : g_source_new : return type :

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Unsupported : g_variant_type_new_array : return type :

// Unsupported : g_variant_type_new_dict_entry : return type :

// Unsupported : g_variant_type_new_maybe : return type :

// Unsupported : g_variant_type_new_tuple : unsupported parameter items :

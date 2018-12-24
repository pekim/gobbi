// This is a generated file - DO NOT EDIT

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

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

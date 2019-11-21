// Code generated - DO NOT EDIT.

package glib

// BookmarkFileError is a representation of the C type GBookmarkFileError.
//
type BookmarkFileError int

const (
	// BookmarkFileError_InvalidUri is a representation of the C type G_BOOKMARK_FILE_ERROR_INVALID_URI.
	BookmarkFileError_InvalidUri BookmarkFileError = 0
	// BookmarkFileError_InvalidValue is a representation of the C type G_BOOKMARK_FILE_ERROR_INVALID_VALUE.
	BookmarkFileError_InvalidValue BookmarkFileError = 1
	// BookmarkFileError_AppNotRegistered is a representation of the C type G_BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED.
	BookmarkFileError_AppNotRegistered BookmarkFileError = 2
	// BookmarkFileError_UriNotFound is a representation of the C type G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
	BookmarkFileError_UriNotFound BookmarkFileError = 3
	// BookmarkFileError_Read is a representation of the C type G_BOOKMARK_FILE_ERROR_READ.
	BookmarkFileError_Read BookmarkFileError = 4
	// BookmarkFileError_UnknownEncoding is a representation of the C type G_BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING.
	BookmarkFileError_UnknownEncoding BookmarkFileError = 5
	// BookmarkFileError_Write is a representation of the C type G_BOOKMARK_FILE_ERROR_WRITE.
	BookmarkFileError_Write BookmarkFileError = 6
	// BookmarkFileError_FileNotFound is a representation of the C type G_BOOKMARK_FILE_ERROR_FILE_NOT_FOUND.
	BookmarkFileError_FileNotFound BookmarkFileError = 7
)

// ChecksumType is a representation of the C type GChecksumType.
//
// since 2.16
type ChecksumType int

const (
	// ChecksumType_Md5 is a representation of the C type G_CHECKSUM_MD5.
	ChecksumType_Md5 ChecksumType = 0
	// ChecksumType_Sha1 is a representation of the C type G_CHECKSUM_SHA1.
	ChecksumType_Sha1 ChecksumType = 1
	// ChecksumType_Sha256 is a representation of the C type G_CHECKSUM_SHA256.
	ChecksumType_Sha256 ChecksumType = 2
	// ChecksumType_Sha512 is a representation of the C type G_CHECKSUM_SHA512.
	ChecksumType_Sha512 ChecksumType = 3
	// ChecksumType_Sha384 is a representation of the C type G_CHECKSUM_SHA384.
	ChecksumType_Sha384 ChecksumType = 4
)

// ConvertError is a representation of the C type GConvertError.
//
type ConvertError int

const (
	// ConvertError_NoConversion is a representation of the C type G_CONVERT_ERROR_NO_CONVERSION.
	ConvertError_NoConversion ConvertError = 0
	// ConvertError_IllegalSequence is a representation of the C type G_CONVERT_ERROR_ILLEGAL_SEQUENCE.
	ConvertError_IllegalSequence ConvertError = 1
	// ConvertError_Failed is a representation of the C type G_CONVERT_ERROR_FAILED.
	ConvertError_Failed ConvertError = 2
	// ConvertError_PartialInput is a representation of the C type G_CONVERT_ERROR_PARTIAL_INPUT.
	ConvertError_PartialInput ConvertError = 3
	// ConvertError_BadUri is a representation of the C type G_CONVERT_ERROR_BAD_URI.
	ConvertError_BadUri ConvertError = 4
	// ConvertError_NotAbsolutePath is a representation of the C type G_CONVERT_ERROR_NOT_ABSOLUTE_PATH.
	ConvertError_NotAbsolutePath ConvertError = 5
	// ConvertError_NoMemory is a representation of the C type G_CONVERT_ERROR_NO_MEMORY.
	ConvertError_NoMemory ConvertError = 6
	// ConvertError_EmbeddedNul is a representation of the C type G_CONVERT_ERROR_EMBEDDED_NUL.
	ConvertError_EmbeddedNul ConvertError = 7
)

// DateDMY is a representation of the C type GDateDMY.
//
type DateDMY int

const (
	// DateDMY_Day is a representation of the C type G_DATE_DAY.
	DateDMY_Day DateDMY = 0
	// DateDMY_Month is a representation of the C type G_DATE_MONTH.
	DateDMY_Month DateDMY = 1
	// DateDMY_Year is a representation of the C type G_DATE_YEAR.
	DateDMY_Year DateDMY = 2
)

// DateMonth is a representation of the C type GDateMonth.
//
type DateMonth int

const (
	// DateMonth_BadMonth is a representation of the C type G_DATE_BAD_MONTH.
	DateMonth_BadMonth DateMonth = 0
	// DateMonth_January is a representation of the C type G_DATE_JANUARY.
	DateMonth_January DateMonth = 1
	// DateMonth_February is a representation of the C type G_DATE_FEBRUARY.
	DateMonth_February DateMonth = 2
	// DateMonth_March is a representation of the C type G_DATE_MARCH.
	DateMonth_March DateMonth = 3
	// DateMonth_April is a representation of the C type G_DATE_APRIL.
	DateMonth_April DateMonth = 4
	// DateMonth_May is a representation of the C type G_DATE_MAY.
	DateMonth_May DateMonth = 5
	// DateMonth_June is a representation of the C type G_DATE_JUNE.
	DateMonth_June DateMonth = 6
	// DateMonth_July is a representation of the C type G_DATE_JULY.
	DateMonth_July DateMonth = 7
	// DateMonth_August is a representation of the C type G_DATE_AUGUST.
	DateMonth_August DateMonth = 8
	// DateMonth_September is a representation of the C type G_DATE_SEPTEMBER.
	DateMonth_September DateMonth = 9
	// DateMonth_October is a representation of the C type G_DATE_OCTOBER.
	DateMonth_October DateMonth = 10
	// DateMonth_November is a representation of the C type G_DATE_NOVEMBER.
	DateMonth_November DateMonth = 11
	// DateMonth_December is a representation of the C type G_DATE_DECEMBER.
	DateMonth_December DateMonth = 12
)

// DateWeekday is a representation of the C type GDateWeekday.
//
type DateWeekday int

const (
	// DateWeekday_BadWeekday is a representation of the C type G_DATE_BAD_WEEKDAY.
	DateWeekday_BadWeekday DateWeekday = 0
	// DateWeekday_Monday is a representation of the C type G_DATE_MONDAY.
	DateWeekday_Monday DateWeekday = 1
	// DateWeekday_Tuesday is a representation of the C type G_DATE_TUESDAY.
	DateWeekday_Tuesday DateWeekday = 2
	// DateWeekday_Wednesday is a representation of the C type G_DATE_WEDNESDAY.
	DateWeekday_Wednesday DateWeekday = 3
	// DateWeekday_Thursday is a representation of the C type G_DATE_THURSDAY.
	DateWeekday_Thursday DateWeekday = 4
	// DateWeekday_Friday is a representation of the C type G_DATE_FRIDAY.
	DateWeekday_Friday DateWeekday = 5
	// DateWeekday_Saturday is a representation of the C type G_DATE_SATURDAY.
	DateWeekday_Saturday DateWeekday = 6
	// DateWeekday_Sunday is a representation of the C type G_DATE_SUNDAY.
	DateWeekday_Sunday DateWeekday = 7
)

// ErrorType is a representation of the C type GErrorType.
//
type ErrorType int

const (
	// ErrorType_Unknown is a representation of the C type G_ERR_UNKNOWN.
	ErrorType_Unknown ErrorType = 0
	// ErrorType_UnexpEof is a representation of the C type G_ERR_UNEXP_EOF.
	ErrorType_UnexpEof ErrorType = 1
	// ErrorType_UnexpEofInString is a representation of the C type G_ERR_UNEXP_EOF_IN_STRING.
	ErrorType_UnexpEofInString ErrorType = 2
	// ErrorType_UnexpEofInComment is a representation of the C type G_ERR_UNEXP_EOF_IN_COMMENT.
	ErrorType_UnexpEofInComment ErrorType = 3
	// ErrorType_NonDigitInConst is a representation of the C type G_ERR_NON_DIGIT_IN_CONST.
	ErrorType_NonDigitInConst ErrorType = 4
	// ErrorType_DigitRadix is a representation of the C type G_ERR_DIGIT_RADIX.
	ErrorType_DigitRadix ErrorType = 5
	// ErrorType_FloatRadix is a representation of the C type G_ERR_FLOAT_RADIX.
	ErrorType_FloatRadix ErrorType = 6
	// ErrorType_FloatMalformed is a representation of the C type G_ERR_FLOAT_MALFORMED.
	ErrorType_FloatMalformed ErrorType = 7
)

// FileError is a representation of the C type GFileError.
//
type FileError int

const (
	// FileError_Exist is a representation of the C type G_FILE_ERROR_EXIST.
	FileError_Exist FileError = 0
	// FileError_Isdir is a representation of the C type G_FILE_ERROR_ISDIR.
	FileError_Isdir FileError = 1
	// FileError_Acces is a representation of the C type G_FILE_ERROR_ACCES.
	FileError_Acces FileError = 2
	// FileError_Nametoolong is a representation of the C type G_FILE_ERROR_NAMETOOLONG.
	FileError_Nametoolong FileError = 3
	// FileError_Noent is a representation of the C type G_FILE_ERROR_NOENT.
	FileError_Noent FileError = 4
	// FileError_Notdir is a representation of the C type G_FILE_ERROR_NOTDIR.
	FileError_Notdir FileError = 5
	// FileError_Nxio is a representation of the C type G_FILE_ERROR_NXIO.
	FileError_Nxio FileError = 6
	// FileError_Nodev is a representation of the C type G_FILE_ERROR_NODEV.
	FileError_Nodev FileError = 7
	// FileError_Rofs is a representation of the C type G_FILE_ERROR_ROFS.
	FileError_Rofs FileError = 8
	// FileError_Txtbsy is a representation of the C type G_FILE_ERROR_TXTBSY.
	FileError_Txtbsy FileError = 9
	// FileError_Fault is a representation of the C type G_FILE_ERROR_FAULT.
	FileError_Fault FileError = 10
	// FileError_Loop is a representation of the C type G_FILE_ERROR_LOOP.
	FileError_Loop FileError = 11
	// FileError_Nospc is a representation of the C type G_FILE_ERROR_NOSPC.
	FileError_Nospc FileError = 12
	// FileError_Nomem is a representation of the C type G_FILE_ERROR_NOMEM.
	FileError_Nomem FileError = 13
	// FileError_Mfile is a representation of the C type G_FILE_ERROR_MFILE.
	FileError_Mfile FileError = 14
	// FileError_Nfile is a representation of the C type G_FILE_ERROR_NFILE.
	FileError_Nfile FileError = 15
	// FileError_Badf is a representation of the C type G_FILE_ERROR_BADF.
	FileError_Badf FileError = 16
	// FileError_Inval is a representation of the C type G_FILE_ERROR_INVAL.
	FileError_Inval FileError = 17
	// FileError_Pipe is a representation of the C type G_FILE_ERROR_PIPE.
	FileError_Pipe FileError = 18
	// FileError_Again is a representation of the C type G_FILE_ERROR_AGAIN.
	FileError_Again FileError = 19
	// FileError_Intr is a representation of the C type G_FILE_ERROR_INTR.
	FileError_Intr FileError = 20
	// FileError_Io is a representation of the C type G_FILE_ERROR_IO.
	FileError_Io FileError = 21
	// FileError_Perm is a representation of the C type G_FILE_ERROR_PERM.
	FileError_Perm FileError = 22
	// FileError_Nosys is a representation of the C type G_FILE_ERROR_NOSYS.
	FileError_Nosys FileError = 23
	// FileError_Failed is a representation of the C type G_FILE_ERROR_FAILED.
	FileError_Failed FileError = 24
)

// IOChannelError is a representation of the C type GIOChannelError.
//
type IOChannelError int

const (
	// IOChannelError_Fbig is a representation of the C type G_IO_CHANNEL_ERROR_FBIG.
	IOChannelError_Fbig IOChannelError = 0
	// IOChannelError_Inval is a representation of the C type G_IO_CHANNEL_ERROR_INVAL.
	IOChannelError_Inval IOChannelError = 1
	// IOChannelError_Io is a representation of the C type G_IO_CHANNEL_ERROR_IO.
	IOChannelError_Io IOChannelError = 2
	// IOChannelError_Isdir is a representation of the C type G_IO_CHANNEL_ERROR_ISDIR.
	IOChannelError_Isdir IOChannelError = 3
	// IOChannelError_Nospc is a representation of the C type G_IO_CHANNEL_ERROR_NOSPC.
	IOChannelError_Nospc IOChannelError = 4
	// IOChannelError_Nxio is a representation of the C type G_IO_CHANNEL_ERROR_NXIO.
	IOChannelError_Nxio IOChannelError = 5
	// IOChannelError_Overflow is a representation of the C type G_IO_CHANNEL_ERROR_OVERFLOW.
	IOChannelError_Overflow IOChannelError = 6
	// IOChannelError_Pipe is a representation of the C type G_IO_CHANNEL_ERROR_PIPE.
	IOChannelError_Pipe IOChannelError = 7
	// IOChannelError_Failed is a representation of the C type G_IO_CHANNEL_ERROR_FAILED.
	IOChannelError_Failed IOChannelError = 8
)

// IOError is a representation of the C type GIOError.
//
type IOError int

const (
	// IOError_None is a representation of the C type G_IO_ERROR_NONE.
	IOError_None IOError = 0
	// IOError_Again is a representation of the C type G_IO_ERROR_AGAIN.
	IOError_Again IOError = 1
	// IOError_Inval is a representation of the C type G_IO_ERROR_INVAL.
	IOError_Inval IOError = 2
	// IOError_Unknown is a representation of the C type G_IO_ERROR_UNKNOWN.
	IOError_Unknown IOError = 3
)

// IOStatus is a representation of the C type GIOStatus.
//
type IOStatus int

const (
	// IOStatus_Error is a representation of the C type G_IO_STATUS_ERROR.
	IOStatus_Error IOStatus = 0
	// IOStatus_Normal is a representation of the C type G_IO_STATUS_NORMAL.
	IOStatus_Normal IOStatus = 1
	// IOStatus_Eof is a representation of the C type G_IO_STATUS_EOF.
	IOStatus_Eof IOStatus = 2
	// IOStatus_Again is a representation of the C type G_IO_STATUS_AGAIN.
	IOStatus_Again IOStatus = 3
)

// KeyFileError is a representation of the C type GKeyFileError.
//
type KeyFileError int

const (
	// KeyFileError_UnknownEncoding is a representation of the C type G_KEY_FILE_ERROR_UNKNOWN_ENCODING.
	KeyFileError_UnknownEncoding KeyFileError = 0
	// KeyFileError_Parse is a representation of the C type G_KEY_FILE_ERROR_PARSE.
	KeyFileError_Parse KeyFileError = 1
	// KeyFileError_NotFound is a representation of the C type G_KEY_FILE_ERROR_NOT_FOUND.
	KeyFileError_NotFound KeyFileError = 2
	// KeyFileError_KeyNotFound is a representation of the C type G_KEY_FILE_ERROR_KEY_NOT_FOUND.
	KeyFileError_KeyNotFound KeyFileError = 3
	// KeyFileError_GroupNotFound is a representation of the C type G_KEY_FILE_ERROR_GROUP_NOT_FOUND.
	KeyFileError_GroupNotFound KeyFileError = 4
	// KeyFileError_InvalidValue is a representation of the C type G_KEY_FILE_ERROR_INVALID_VALUE.
	KeyFileError_InvalidValue KeyFileError = 5
)

// LogWriterOutput is a representation of the C type GLogWriterOutput.
//
// since 2.50
type LogWriterOutput int

const (
	// LogWriterOutput_Handled is a representation of the C type G_LOG_WRITER_HANDLED.
	LogWriterOutput_Handled LogWriterOutput = 1
	// LogWriterOutput_Unhandled is a representation of the C type G_LOG_WRITER_UNHANDLED.
	LogWriterOutput_Unhandled LogWriterOutput = 0
)

// MarkupError is a representation of the C type GMarkupError.
//
type MarkupError int

const (
	// MarkupError_BadUtf8 is a representation of the C type G_MARKUP_ERROR_BAD_UTF8.
	MarkupError_BadUtf8 MarkupError = 0
	// MarkupError_Empty is a representation of the C type G_MARKUP_ERROR_EMPTY.
	MarkupError_Empty MarkupError = 1
	// MarkupError_Parse is a representation of the C type G_MARKUP_ERROR_PARSE.
	MarkupError_Parse MarkupError = 2
	// MarkupError_UnknownElement is a representation of the C type G_MARKUP_ERROR_UNKNOWN_ELEMENT.
	MarkupError_UnknownElement MarkupError = 3
	// MarkupError_UnknownAttribute is a representation of the C type G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE.
	MarkupError_UnknownAttribute MarkupError = 4
	// MarkupError_InvalidContent is a representation of the C type G_MARKUP_ERROR_INVALID_CONTENT.
	MarkupError_InvalidContent MarkupError = 5
	// MarkupError_MissingAttribute is a representation of the C type G_MARKUP_ERROR_MISSING_ATTRIBUTE.
	MarkupError_MissingAttribute MarkupError = 6
)

// NormalizeMode is a representation of the C type GNormalizeMode.
//
type NormalizeMode int

const (
	// NormalizeMode_Default is a representation of the C type G_NORMALIZE_DEFAULT.
	NormalizeMode_Default NormalizeMode = 0
	// NormalizeMode_Nfd is a representation of the C type G_NORMALIZE_NFD.
	NormalizeMode_Nfd NormalizeMode = 0
	// NormalizeMode_DefaultCompose is a representation of the C type G_NORMALIZE_DEFAULT_COMPOSE.
	NormalizeMode_DefaultCompose NormalizeMode = 1
	// NormalizeMode_Nfc is a representation of the C type G_NORMALIZE_NFC.
	NormalizeMode_Nfc NormalizeMode = 1
	// NormalizeMode_All is a representation of the C type G_NORMALIZE_ALL.
	NormalizeMode_All NormalizeMode = 2
	// NormalizeMode_Nfkd is a representation of the C type G_NORMALIZE_NFKD.
	NormalizeMode_Nfkd NormalizeMode = 2
	// NormalizeMode_AllCompose is a representation of the C type G_NORMALIZE_ALL_COMPOSE.
	NormalizeMode_AllCompose NormalizeMode = 3
	// NormalizeMode_Nfkc is a representation of the C type G_NORMALIZE_NFKC.
	NormalizeMode_Nfkc NormalizeMode = 3
)

// NumberParserError is a representation of the C type GNumberParserError.
//
// since 2.54
type NumberParserError int

const (
	// NumberParserError_Invalid is a representation of the C type G_NUMBER_PARSER_ERROR_INVALID.
	NumberParserError_Invalid NumberParserError = 0
	// NumberParserError_OutOfBounds is a representation of the C type G_NUMBER_PARSER_ERROR_OUT_OF_BOUNDS.
	NumberParserError_OutOfBounds NumberParserError = 1
)

// OnceStatus is a representation of the C type GOnceStatus.
//
// since 2.4
type OnceStatus int

const (
	// OnceStatus_Notcalled is a representation of the C type G_ONCE_STATUS_NOTCALLED.
	OnceStatus_Notcalled OnceStatus = 0
	// OnceStatus_Progress is a representation of the C type G_ONCE_STATUS_PROGRESS.
	OnceStatus_Progress OnceStatus = 1
	// OnceStatus_Ready is a representation of the C type G_ONCE_STATUS_READY.
	OnceStatus_Ready OnceStatus = 2
)

// OptionArg is a representation of the C type GOptionArg.
//
type OptionArg int

const (
	// OptionArg_None is a representation of the C type G_OPTION_ARG_NONE.
	OptionArg_None OptionArg = 0
	// OptionArg_String is a representation of the C type G_OPTION_ARG_STRING.
	OptionArg_String OptionArg = 1
	// OptionArg_Int is a representation of the C type G_OPTION_ARG_INT.
	OptionArg_Int OptionArg = 2
	// OptionArg_Callback is a representation of the C type G_OPTION_ARG_CALLBACK.
	OptionArg_Callback OptionArg = 3
	// OptionArg_Filename is a representation of the C type G_OPTION_ARG_FILENAME.
	OptionArg_Filename OptionArg = 4
	// OptionArg_StringArray is a representation of the C type G_OPTION_ARG_STRING_ARRAY.
	OptionArg_StringArray OptionArg = 5
	// OptionArg_FilenameArray is a representation of the C type G_OPTION_ARG_FILENAME_ARRAY.
	OptionArg_FilenameArray OptionArg = 6
	// OptionArg_Double is a representation of the C type G_OPTION_ARG_DOUBLE.
	OptionArg_Double OptionArg = 7
	// OptionArg_Int64 is a representation of the C type G_OPTION_ARG_INT64.
	OptionArg_Int64 OptionArg = 8
)

// OptionError is a representation of the C type GOptionError.
//
type OptionError int

const (
	// OptionError_UnknownOption is a representation of the C type G_OPTION_ERROR_UNKNOWN_OPTION.
	OptionError_UnknownOption OptionError = 0
	// OptionError_BadValue is a representation of the C type G_OPTION_ERROR_BAD_VALUE.
	OptionError_BadValue OptionError = 1
	// OptionError_Failed is a representation of the C type G_OPTION_ERROR_FAILED.
	OptionError_Failed OptionError = 2
)

// RegexError is a representation of the C type GRegexError.
//
// since 2.14
type RegexError int

const (
	// RegexError_Compile is a representation of the C type G_REGEX_ERROR_COMPILE.
	RegexError_Compile RegexError = 0
	// RegexError_Optimize is a representation of the C type G_REGEX_ERROR_OPTIMIZE.
	RegexError_Optimize RegexError = 1
	// RegexError_Replace is a representation of the C type G_REGEX_ERROR_REPLACE.
	RegexError_Replace RegexError = 2
	// RegexError_Match is a representation of the C type G_REGEX_ERROR_MATCH.
	RegexError_Match RegexError = 3
	// RegexError_Internal is a representation of the C type G_REGEX_ERROR_INTERNAL.
	RegexError_Internal RegexError = 4
	// RegexError_StrayBackslash is a representation of the C type G_REGEX_ERROR_STRAY_BACKSLASH.
	RegexError_StrayBackslash RegexError = 101
	// RegexError_MissingControlChar is a representation of the C type G_REGEX_ERROR_MISSING_CONTROL_CHAR.
	RegexError_MissingControlChar RegexError = 102
	// RegexError_UnrecognizedEscape is a representation of the C type G_REGEX_ERROR_UNRECOGNIZED_ESCAPE.
	RegexError_UnrecognizedEscape RegexError = 103
	// RegexError_QuantifiersOutOfOrder is a representation of the C type G_REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER.
	RegexError_QuantifiersOutOfOrder RegexError = 104
	// RegexError_QuantifierTooBig is a representation of the C type G_REGEX_ERROR_QUANTIFIER_TOO_BIG.
	RegexError_QuantifierTooBig RegexError = 105
	// RegexError_UnterminatedCharacterClass is a representation of the C type G_REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS.
	RegexError_UnterminatedCharacterClass RegexError = 106
	// RegexError_InvalidEscapeInCharacterClass is a representation of the C type G_REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS.
	RegexError_InvalidEscapeInCharacterClass RegexError = 107
	// RegexError_RangeOutOfOrder is a representation of the C type G_REGEX_ERROR_RANGE_OUT_OF_ORDER.
	RegexError_RangeOutOfOrder RegexError = 108
	// RegexError_NothingToRepeat is a representation of the C type G_REGEX_ERROR_NOTHING_TO_REPEAT.
	RegexError_NothingToRepeat RegexError = 109
	// RegexError_UnrecognizedCharacter is a representation of the C type G_REGEX_ERROR_UNRECOGNIZED_CHARACTER.
	RegexError_UnrecognizedCharacter RegexError = 112
	// RegexError_PosixNamedClassOutsideClass is a representation of the C type G_REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS.
	RegexError_PosixNamedClassOutsideClass RegexError = 113
	// RegexError_UnmatchedParenthesis is a representation of the C type G_REGEX_ERROR_UNMATCHED_PARENTHESIS.
	RegexError_UnmatchedParenthesis RegexError = 114
	// RegexError_InexistentSubpatternReference is a representation of the C type G_REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE.
	RegexError_InexistentSubpatternReference RegexError = 115
	// RegexError_UnterminatedComment is a representation of the C type G_REGEX_ERROR_UNTERMINATED_COMMENT.
	RegexError_UnterminatedComment RegexError = 118
	// RegexError_ExpressionTooLarge is a representation of the C type G_REGEX_ERROR_EXPRESSION_TOO_LARGE.
	RegexError_ExpressionTooLarge RegexError = 120
	// RegexError_MemoryError is a representation of the C type G_REGEX_ERROR_MEMORY_ERROR.
	RegexError_MemoryError RegexError = 121
	// RegexError_VariableLengthLookbehind is a representation of the C type G_REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND.
	RegexError_VariableLengthLookbehind RegexError = 125
	// RegexError_MalformedCondition is a representation of the C type G_REGEX_ERROR_MALFORMED_CONDITION.
	RegexError_MalformedCondition RegexError = 126
	// RegexError_TooManyConditionalBranches is a representation of the C type G_REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES.
	RegexError_TooManyConditionalBranches RegexError = 127
	// RegexError_AssertionExpected is a representation of the C type G_REGEX_ERROR_ASSERTION_EXPECTED.
	RegexError_AssertionExpected RegexError = 128
	// RegexError_UnknownPosixClassName is a representation of the C type G_REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME.
	RegexError_UnknownPosixClassName RegexError = 130
	// RegexError_PosixCollatingElementsNotSupported is a representation of the C type G_REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED.
	RegexError_PosixCollatingElementsNotSupported RegexError = 131
	// RegexError_HexCodeTooLarge is a representation of the C type G_REGEX_ERROR_HEX_CODE_TOO_LARGE.
	RegexError_HexCodeTooLarge RegexError = 134
	// RegexError_InvalidCondition is a representation of the C type G_REGEX_ERROR_INVALID_CONDITION.
	RegexError_InvalidCondition RegexError = 135
	// RegexError_SingleByteMatchInLookbehind is a representation of the C type G_REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND.
	RegexError_SingleByteMatchInLookbehind RegexError = 136
	// RegexError_InfiniteLoop is a representation of the C type G_REGEX_ERROR_INFINITE_LOOP.
	RegexError_InfiniteLoop RegexError = 140
	// RegexError_MissingSubpatternNameTerminator is a representation of the C type G_REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR.
	RegexError_MissingSubpatternNameTerminator RegexError = 142
	// RegexError_DuplicateSubpatternName is a representation of the C type G_REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME.
	RegexError_DuplicateSubpatternName RegexError = 143
	// RegexError_MalformedProperty is a representation of the C type G_REGEX_ERROR_MALFORMED_PROPERTY.
	RegexError_MalformedProperty RegexError = 146
	// RegexError_UnknownProperty is a representation of the C type G_REGEX_ERROR_UNKNOWN_PROPERTY.
	RegexError_UnknownProperty RegexError = 147
	// RegexError_SubpatternNameTooLong is a representation of the C type G_REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG.
	RegexError_SubpatternNameTooLong RegexError = 148
	// RegexError_TooManySubpatterns is a representation of the C type G_REGEX_ERROR_TOO_MANY_SUBPATTERNS.
	RegexError_TooManySubpatterns RegexError = 149
	// RegexError_InvalidOctalValue is a representation of the C type G_REGEX_ERROR_INVALID_OCTAL_VALUE.
	RegexError_InvalidOctalValue RegexError = 151
	// RegexError_TooManyBranchesInDefine is a representation of the C type G_REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE.
	RegexError_TooManyBranchesInDefine RegexError = 154
	// RegexError_DefineRepetion is a representation of the C type G_REGEX_ERROR_DEFINE_REPETION.
	RegexError_DefineRepetion RegexError = 155
	// RegexError_InconsistentNewlineOptions is a representation of the C type G_REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS.
	RegexError_InconsistentNewlineOptions RegexError = 156
	// RegexError_MissingBackReference is a representation of the C type G_REGEX_ERROR_MISSING_BACK_REFERENCE.
	RegexError_MissingBackReference RegexError = 157
	// RegexError_InvalidRelativeReference is a representation of the C type G_REGEX_ERROR_INVALID_RELATIVE_REFERENCE.
	RegexError_InvalidRelativeReference RegexError = 158
	// RegexError_BacktrackingControlVerbArgumentForbidden is a representation of the C type G_REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN.
	RegexError_BacktrackingControlVerbArgumentForbidden RegexError = 159
	// RegexError_UnknownBacktrackingControlVerb is a representation of the C type G_REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB.
	RegexError_UnknownBacktrackingControlVerb RegexError = 160
	// RegexError_NumberTooBig is a representation of the C type G_REGEX_ERROR_NUMBER_TOO_BIG.
	RegexError_NumberTooBig RegexError = 161
	// RegexError_MissingSubpatternName is a representation of the C type G_REGEX_ERROR_MISSING_SUBPATTERN_NAME.
	RegexError_MissingSubpatternName RegexError = 162
	// RegexError_MissingDigit is a representation of the C type G_REGEX_ERROR_MISSING_DIGIT.
	RegexError_MissingDigit RegexError = 163
	// RegexError_InvalidDataCharacter is a representation of the C type G_REGEX_ERROR_INVALID_DATA_CHARACTER.
	RegexError_InvalidDataCharacter RegexError = 164
	// RegexError_ExtraSubpatternName is a representation of the C type G_REGEX_ERROR_EXTRA_SUBPATTERN_NAME.
	RegexError_ExtraSubpatternName RegexError = 165
	// RegexError_BacktrackingControlVerbArgumentRequired is a representation of the C type G_REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED.
	RegexError_BacktrackingControlVerbArgumentRequired RegexError = 166
	// RegexError_InvalidControlChar is a representation of the C type G_REGEX_ERROR_INVALID_CONTROL_CHAR.
	RegexError_InvalidControlChar RegexError = 168
	// RegexError_MissingName is a representation of the C type G_REGEX_ERROR_MISSING_NAME.
	RegexError_MissingName RegexError = 169
	// RegexError_NotSupportedInClass is a representation of the C type G_REGEX_ERROR_NOT_SUPPORTED_IN_CLASS.
	RegexError_NotSupportedInClass RegexError = 171
	// RegexError_TooManyForwardReferences is a representation of the C type G_REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES.
	RegexError_TooManyForwardReferences RegexError = 172
	// RegexError_NameTooLong is a representation of the C type G_REGEX_ERROR_NAME_TOO_LONG.
	RegexError_NameTooLong RegexError = 175
	// RegexError_CharacterValueTooLarge is a representation of the C type G_REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE.
	RegexError_CharacterValueTooLarge RegexError = 176
)

// SeekType is a representation of the C type GSeekType.
//
type SeekType int

const (
	// SeekType_Cur is a representation of the C type G_SEEK_CUR.
	SeekType_Cur SeekType = 0
	// SeekType_Set is a representation of the C type G_SEEK_SET.
	SeekType_Set SeekType = 1
	// SeekType_End is a representation of the C type G_SEEK_END.
	SeekType_End SeekType = 2
)

// ShellError is a representation of the C type GShellError.
//
type ShellError int

const (
	// ShellError_BadQuoting is a representation of the C type G_SHELL_ERROR_BAD_QUOTING.
	ShellError_BadQuoting ShellError = 0
	// ShellError_EmptyString is a representation of the C type G_SHELL_ERROR_EMPTY_STRING.
	ShellError_EmptyString ShellError = 1
	// ShellError_Failed is a representation of the C type G_SHELL_ERROR_FAILED.
	ShellError_Failed ShellError = 2
)

// SliceConfig is a representation of the C type GSliceConfig.
//
type SliceConfig int

const (
	// SliceConfig_AlwaysMalloc is a representation of the C type G_SLICE_CONFIG_ALWAYS_MALLOC.
	SliceConfig_AlwaysMalloc SliceConfig = 1
	// SliceConfig_BypassMagazines is a representation of the C type G_SLICE_CONFIG_BYPASS_MAGAZINES.
	SliceConfig_BypassMagazines SliceConfig = 2
	// SliceConfig_WorkingSetMsecs is a representation of the C type G_SLICE_CONFIG_WORKING_SET_MSECS.
	SliceConfig_WorkingSetMsecs SliceConfig = 3
	// SliceConfig_ColorIncrement is a representation of the C type G_SLICE_CONFIG_COLOR_INCREMENT.
	SliceConfig_ColorIncrement SliceConfig = 4
	// SliceConfig_ChunkSizes is a representation of the C type G_SLICE_CONFIG_CHUNK_SIZES.
	SliceConfig_ChunkSizes SliceConfig = 5
	// SliceConfig_ContentionCounter is a representation of the C type G_SLICE_CONFIG_CONTENTION_COUNTER.
	SliceConfig_ContentionCounter SliceConfig = 6
)

// SpawnError is a representation of the C type GSpawnError.
//
type SpawnError int

const (
	// SpawnError_Fork is a representation of the C type G_SPAWN_ERROR_FORK.
	SpawnError_Fork SpawnError = 0
	// SpawnError_Read is a representation of the C type G_SPAWN_ERROR_READ.
	SpawnError_Read SpawnError = 1
	// SpawnError_Chdir is a representation of the C type G_SPAWN_ERROR_CHDIR.
	SpawnError_Chdir SpawnError = 2
	// SpawnError_Acces is a representation of the C type G_SPAWN_ERROR_ACCES.
	SpawnError_Acces SpawnError = 3
	// SpawnError_Perm is a representation of the C type G_SPAWN_ERROR_PERM.
	SpawnError_Perm SpawnError = 4
	// SpawnError_TooBig is a representation of the C type G_SPAWN_ERROR_TOO_BIG.
	SpawnError_TooBig SpawnError = 5
	// SpawnError_2big is a representation of the C type G_SPAWN_ERROR_2BIG.
	SpawnError_2big SpawnError = 5
	// SpawnError_Noexec is a representation of the C type G_SPAWN_ERROR_NOEXEC.
	SpawnError_Noexec SpawnError = 6
	// SpawnError_Nametoolong is a representation of the C type G_SPAWN_ERROR_NAMETOOLONG.
	SpawnError_Nametoolong SpawnError = 7
	// SpawnError_Noent is a representation of the C type G_SPAWN_ERROR_NOENT.
	SpawnError_Noent SpawnError = 8
	// SpawnError_Nomem is a representation of the C type G_SPAWN_ERROR_NOMEM.
	SpawnError_Nomem SpawnError = 9
	// SpawnError_Notdir is a representation of the C type G_SPAWN_ERROR_NOTDIR.
	SpawnError_Notdir SpawnError = 10
	// SpawnError_Loop is a representation of the C type G_SPAWN_ERROR_LOOP.
	SpawnError_Loop SpawnError = 11
	// SpawnError_Txtbusy is a representation of the C type G_SPAWN_ERROR_TXTBUSY.
	SpawnError_Txtbusy SpawnError = 12
	// SpawnError_Io is a representation of the C type G_SPAWN_ERROR_IO.
	SpawnError_Io SpawnError = 13
	// SpawnError_Nfile is a representation of the C type G_SPAWN_ERROR_NFILE.
	SpawnError_Nfile SpawnError = 14
	// SpawnError_Mfile is a representation of the C type G_SPAWN_ERROR_MFILE.
	SpawnError_Mfile SpawnError = 15
	// SpawnError_Inval is a representation of the C type G_SPAWN_ERROR_INVAL.
	SpawnError_Inval SpawnError = 16
	// SpawnError_Isdir is a representation of the C type G_SPAWN_ERROR_ISDIR.
	SpawnError_Isdir SpawnError = 17
	// SpawnError_Libbad is a representation of the C type G_SPAWN_ERROR_LIBBAD.
	SpawnError_Libbad SpawnError = 18
	// SpawnError_Failed is a representation of the C type G_SPAWN_ERROR_FAILED.
	SpawnError_Failed SpawnError = 19
)

// TestFileType is a representation of the C type GTestFileType.
//
// since 2.38
type TestFileType int

const (
	// TestFileType_Dist is a representation of the C type G_TEST_DIST.
	TestFileType_Dist TestFileType = 0
	// TestFileType_Built is a representation of the C type G_TEST_BUILT.
	TestFileType_Built TestFileType = 1
)

// TestLogType is a representation of the C type GTestLogType.
//
type TestLogType int

const (
	// TestLogType_None is a representation of the C type G_TEST_LOG_NONE.
	TestLogType_None TestLogType = 0
	// TestLogType_Error is a representation of the C type G_TEST_LOG_ERROR.
	TestLogType_Error TestLogType = 1
	// TestLogType_StartBinary is a representation of the C type G_TEST_LOG_START_BINARY.
	TestLogType_StartBinary TestLogType = 2
	// TestLogType_ListCase is a representation of the C type G_TEST_LOG_LIST_CASE.
	TestLogType_ListCase TestLogType = 3
	// TestLogType_SkipCase is a representation of the C type G_TEST_LOG_SKIP_CASE.
	TestLogType_SkipCase TestLogType = 4
	// TestLogType_StartCase is a representation of the C type G_TEST_LOG_START_CASE.
	TestLogType_StartCase TestLogType = 5
	// TestLogType_StopCase is a representation of the C type G_TEST_LOG_STOP_CASE.
	TestLogType_StopCase TestLogType = 6
	// TestLogType_MinResult is a representation of the C type G_TEST_LOG_MIN_RESULT.
	TestLogType_MinResult TestLogType = 7
	// TestLogType_MaxResult is a representation of the C type G_TEST_LOG_MAX_RESULT.
	TestLogType_MaxResult TestLogType = 8
	// TestLogType_Message is a representation of the C type G_TEST_LOG_MESSAGE.
	TestLogType_Message TestLogType = 9
	// TestLogType_StartSuite is a representation of the C type G_TEST_LOG_START_SUITE.
	TestLogType_StartSuite TestLogType = 10
	// TestLogType_StopSuite is a representation of the C type G_TEST_LOG_STOP_SUITE.
	TestLogType_StopSuite TestLogType = 11
)

// TestResult is a representation of the C type GTestResult.
//
type TestResult int

const (
	// TestResult_Success is a representation of the C type G_TEST_RUN_SUCCESS.
	TestResult_Success TestResult = 0
	// TestResult_Skipped is a representation of the C type G_TEST_RUN_SKIPPED.
	TestResult_Skipped TestResult = 1
	// TestResult_Failure is a representation of the C type G_TEST_RUN_FAILURE.
	TestResult_Failure TestResult = 2
	// TestResult_Incomplete is a representation of the C type G_TEST_RUN_INCOMPLETE.
	TestResult_Incomplete TestResult = 3
)

// ThreadError is a representation of the C type GThreadError.
//
type ThreadError int

const (
	// ThreadError_ThreadErrorAgain is a representation of the C type G_THREAD_ERROR_AGAIN.
	ThreadError_ThreadErrorAgain ThreadError = 0
)

// TimeType is a representation of the C type GTimeType.
//
type TimeType int

const (
	// TimeType_Standard is a representation of the C type G_TIME_TYPE_STANDARD.
	TimeType_Standard TimeType = 0
	// TimeType_Daylight is a representation of the C type G_TIME_TYPE_DAYLIGHT.
	TimeType_Daylight TimeType = 1
	// TimeType_Universal is a representation of the C type G_TIME_TYPE_UNIVERSAL.
	TimeType_Universal TimeType = 2
)

// TokenType is a representation of the C type GTokenType.
//
type TokenType int

const (
	// TokenType_Eof is a representation of the C type G_TOKEN_EOF.
	TokenType_Eof TokenType = 0
	// TokenType_LeftParen is a representation of the C type G_TOKEN_LEFT_PAREN.
	TokenType_LeftParen TokenType = 40
	// TokenType_RightParen is a representation of the C type G_TOKEN_RIGHT_PAREN.
	TokenType_RightParen TokenType = 41
	// TokenType_LeftCurly is a representation of the C type G_TOKEN_LEFT_CURLY.
	TokenType_LeftCurly TokenType = 123
	// TokenType_RightCurly is a representation of the C type G_TOKEN_RIGHT_CURLY.
	TokenType_RightCurly TokenType = 125
	// TokenType_LeftBrace is a representation of the C type G_TOKEN_LEFT_BRACE.
	TokenType_LeftBrace TokenType = 91
	// TokenType_RightBrace is a representation of the C type G_TOKEN_RIGHT_BRACE.
	TokenType_RightBrace TokenType = 93
	// TokenType_EqualSign is a representation of the C type G_TOKEN_EQUAL_SIGN.
	TokenType_EqualSign TokenType = 61
	// TokenType_Comma is a representation of the C type G_TOKEN_COMMA.
	TokenType_Comma TokenType = 44
	// TokenType_None is a representation of the C type G_TOKEN_NONE.
	TokenType_None TokenType = 256
	// TokenType_Error is a representation of the C type G_TOKEN_ERROR.
	TokenType_Error TokenType = 257
	// TokenType_Char is a representation of the C type G_TOKEN_CHAR.
	TokenType_Char TokenType = 258
	// TokenType_Binary is a representation of the C type G_TOKEN_BINARY.
	TokenType_Binary TokenType = 259
	// TokenType_Octal is a representation of the C type G_TOKEN_OCTAL.
	TokenType_Octal TokenType = 260
	// TokenType_Int is a representation of the C type G_TOKEN_INT.
	TokenType_Int TokenType = 261
	// TokenType_Hex is a representation of the C type G_TOKEN_HEX.
	TokenType_Hex TokenType = 262
	// TokenType_Float is a representation of the C type G_TOKEN_FLOAT.
	TokenType_Float TokenType = 263
	// TokenType_String is a representation of the C type G_TOKEN_STRING.
	TokenType_String TokenType = 264
	// TokenType_Symbol is a representation of the C type G_TOKEN_SYMBOL.
	TokenType_Symbol TokenType = 265
	// TokenType_Identifier is a representation of the C type G_TOKEN_IDENTIFIER.
	TokenType_Identifier TokenType = 266
	// TokenType_IdentifierNull is a representation of the C type G_TOKEN_IDENTIFIER_NULL.
	TokenType_IdentifierNull TokenType = 267
	// TokenType_CommentSingle is a representation of the C type G_TOKEN_COMMENT_SINGLE.
	TokenType_CommentSingle TokenType = 268
	// TokenType_CommentMulti is a representation of the C type G_TOKEN_COMMENT_MULTI.
	TokenType_CommentMulti TokenType = 269
)

// TraverseType is a representation of the C type GTraverseType.
//
type TraverseType int

const (
	// TraverseType_InOrder is a representation of the C type G_IN_ORDER.
	TraverseType_InOrder TraverseType = 0
	// TraverseType_PreOrder is a representation of the C type G_PRE_ORDER.
	TraverseType_PreOrder TraverseType = 1
	// TraverseType_PostOrder is a representation of the C type G_POST_ORDER.
	TraverseType_PostOrder TraverseType = 2
	// TraverseType_LevelOrder is a representation of the C type G_LEVEL_ORDER.
	TraverseType_LevelOrder TraverseType = 3
)

// UnicodeBreakType is a representation of the C type GUnicodeBreakType.
//
type UnicodeBreakType int

const (
	// UnicodeBreakType_Mandatory is a representation of the C type G_UNICODE_BREAK_MANDATORY.
	UnicodeBreakType_Mandatory UnicodeBreakType = 0
	// UnicodeBreakType_CarriageReturn is a representation of the C type G_UNICODE_BREAK_CARRIAGE_RETURN.
	UnicodeBreakType_CarriageReturn UnicodeBreakType = 1
	// UnicodeBreakType_LineFeed is a representation of the C type G_UNICODE_BREAK_LINE_FEED.
	UnicodeBreakType_LineFeed UnicodeBreakType = 2
	// UnicodeBreakType_CombiningMark is a representation of the C type G_UNICODE_BREAK_COMBINING_MARK.
	UnicodeBreakType_CombiningMark UnicodeBreakType = 3
	// UnicodeBreakType_Surrogate is a representation of the C type G_UNICODE_BREAK_SURROGATE.
	UnicodeBreakType_Surrogate UnicodeBreakType = 4
	// UnicodeBreakType_ZeroWidthSpace is a representation of the C type G_UNICODE_BREAK_ZERO_WIDTH_SPACE.
	UnicodeBreakType_ZeroWidthSpace UnicodeBreakType = 5
	// UnicodeBreakType_Inseparable is a representation of the C type G_UNICODE_BREAK_INSEPARABLE.
	UnicodeBreakType_Inseparable UnicodeBreakType = 6
	// UnicodeBreakType_NonBreakingGlue is a representation of the C type G_UNICODE_BREAK_NON_BREAKING_GLUE.
	UnicodeBreakType_NonBreakingGlue UnicodeBreakType = 7
	// UnicodeBreakType_Contingent is a representation of the C type G_UNICODE_BREAK_CONTINGENT.
	UnicodeBreakType_Contingent UnicodeBreakType = 8
	// UnicodeBreakType_Space is a representation of the C type G_UNICODE_BREAK_SPACE.
	UnicodeBreakType_Space UnicodeBreakType = 9
	// UnicodeBreakType_After is a representation of the C type G_UNICODE_BREAK_AFTER.
	UnicodeBreakType_After UnicodeBreakType = 10
	// UnicodeBreakType_Before is a representation of the C type G_UNICODE_BREAK_BEFORE.
	UnicodeBreakType_Before UnicodeBreakType = 11
	// UnicodeBreakType_BeforeAndAfter is a representation of the C type G_UNICODE_BREAK_BEFORE_AND_AFTER.
	UnicodeBreakType_BeforeAndAfter UnicodeBreakType = 12
	// UnicodeBreakType_Hyphen is a representation of the C type G_UNICODE_BREAK_HYPHEN.
	UnicodeBreakType_Hyphen UnicodeBreakType = 13
	// UnicodeBreakType_NonStarter is a representation of the C type G_UNICODE_BREAK_NON_STARTER.
	UnicodeBreakType_NonStarter UnicodeBreakType = 14
	// UnicodeBreakType_OpenPunctuation is a representation of the C type G_UNICODE_BREAK_OPEN_PUNCTUATION.
	UnicodeBreakType_OpenPunctuation UnicodeBreakType = 15
	// UnicodeBreakType_ClosePunctuation is a representation of the C type G_UNICODE_BREAK_CLOSE_PUNCTUATION.
	UnicodeBreakType_ClosePunctuation UnicodeBreakType = 16
	// UnicodeBreakType_Quotation is a representation of the C type G_UNICODE_BREAK_QUOTATION.
	UnicodeBreakType_Quotation UnicodeBreakType = 17
	// UnicodeBreakType_Exclamation is a representation of the C type G_UNICODE_BREAK_EXCLAMATION.
	UnicodeBreakType_Exclamation UnicodeBreakType = 18
	// UnicodeBreakType_Ideographic is a representation of the C type G_UNICODE_BREAK_IDEOGRAPHIC.
	UnicodeBreakType_Ideographic UnicodeBreakType = 19
	// UnicodeBreakType_Numeric is a representation of the C type G_UNICODE_BREAK_NUMERIC.
	UnicodeBreakType_Numeric UnicodeBreakType = 20
	// UnicodeBreakType_InfixSeparator is a representation of the C type G_UNICODE_BREAK_INFIX_SEPARATOR.
	UnicodeBreakType_InfixSeparator UnicodeBreakType = 21
	// UnicodeBreakType_Symbol is a representation of the C type G_UNICODE_BREAK_SYMBOL.
	UnicodeBreakType_Symbol UnicodeBreakType = 22
	// UnicodeBreakType_Alphabetic is a representation of the C type G_UNICODE_BREAK_ALPHABETIC.
	UnicodeBreakType_Alphabetic UnicodeBreakType = 23
	// UnicodeBreakType_Prefix is a representation of the C type G_UNICODE_BREAK_PREFIX.
	UnicodeBreakType_Prefix UnicodeBreakType = 24
	// UnicodeBreakType_Postfix is a representation of the C type G_UNICODE_BREAK_POSTFIX.
	UnicodeBreakType_Postfix UnicodeBreakType = 25
	// UnicodeBreakType_ComplexContext is a representation of the C type G_UNICODE_BREAK_COMPLEX_CONTEXT.
	UnicodeBreakType_ComplexContext UnicodeBreakType = 26
	// UnicodeBreakType_Ambiguous is a representation of the C type G_UNICODE_BREAK_AMBIGUOUS.
	UnicodeBreakType_Ambiguous UnicodeBreakType = 27
	// UnicodeBreakType_Unknown is a representation of the C type G_UNICODE_BREAK_UNKNOWN.
	UnicodeBreakType_Unknown UnicodeBreakType = 28
	// UnicodeBreakType_NextLine is a representation of the C type G_UNICODE_BREAK_NEXT_LINE.
	UnicodeBreakType_NextLine UnicodeBreakType = 29
	// UnicodeBreakType_WordJoiner is a representation of the C type G_UNICODE_BREAK_WORD_JOINER.
	UnicodeBreakType_WordJoiner UnicodeBreakType = 30
	// UnicodeBreakType_HangulLJamo is a representation of the C type G_UNICODE_BREAK_HANGUL_L_JAMO.
	UnicodeBreakType_HangulLJamo UnicodeBreakType = 31
	// UnicodeBreakType_HangulVJamo is a representation of the C type G_UNICODE_BREAK_HANGUL_V_JAMO.
	UnicodeBreakType_HangulVJamo UnicodeBreakType = 32
	// UnicodeBreakType_HangulTJamo is a representation of the C type G_UNICODE_BREAK_HANGUL_T_JAMO.
	UnicodeBreakType_HangulTJamo UnicodeBreakType = 33
	// UnicodeBreakType_HangulLvSyllable is a representation of the C type G_UNICODE_BREAK_HANGUL_LV_SYLLABLE.
	UnicodeBreakType_HangulLvSyllable UnicodeBreakType = 34
	// UnicodeBreakType_HangulLvtSyllable is a representation of the C type G_UNICODE_BREAK_HANGUL_LVT_SYLLABLE.
	UnicodeBreakType_HangulLvtSyllable UnicodeBreakType = 35
	// UnicodeBreakType_CloseParanthesis is a representation of the C type G_UNICODE_BREAK_CLOSE_PARANTHESIS.
	UnicodeBreakType_CloseParanthesis UnicodeBreakType = 36
	// UnicodeBreakType_ConditionalJapaneseStarter is a representation of the C type G_UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER.
	UnicodeBreakType_ConditionalJapaneseStarter UnicodeBreakType = 37
	// UnicodeBreakType_HebrewLetter is a representation of the C type G_UNICODE_BREAK_HEBREW_LETTER.
	UnicodeBreakType_HebrewLetter UnicodeBreakType = 38
	// UnicodeBreakType_RegionalIndicator is a representation of the C type G_UNICODE_BREAK_REGIONAL_INDICATOR.
	UnicodeBreakType_RegionalIndicator UnicodeBreakType = 39
	// UnicodeBreakType_EmojiBase is a representation of the C type G_UNICODE_BREAK_EMOJI_BASE.
	UnicodeBreakType_EmojiBase UnicodeBreakType = 40
	// UnicodeBreakType_EmojiModifier is a representation of the C type G_UNICODE_BREAK_EMOJI_MODIFIER.
	UnicodeBreakType_EmojiModifier UnicodeBreakType = 41
	// UnicodeBreakType_ZeroWidthJoiner is a representation of the C type G_UNICODE_BREAK_ZERO_WIDTH_JOINER.
	UnicodeBreakType_ZeroWidthJoiner UnicodeBreakType = 42
)

// UnicodeScript is a representation of the C type GUnicodeScript.
//
type UnicodeScript int

const (
	// UnicodeScript_InvalidCode is a representation of the C type G_UNICODE_SCRIPT_INVALID_CODE.
	UnicodeScript_InvalidCode UnicodeScript = -1
	// UnicodeScript_Common is a representation of the C type G_UNICODE_SCRIPT_COMMON.
	UnicodeScript_Common UnicodeScript = 0
	// UnicodeScript_Inherited is a representation of the C type G_UNICODE_SCRIPT_INHERITED.
	UnicodeScript_Inherited UnicodeScript = 1
	// UnicodeScript_Arabic is a representation of the C type G_UNICODE_SCRIPT_ARABIC.
	UnicodeScript_Arabic UnicodeScript = 2
	// UnicodeScript_Armenian is a representation of the C type G_UNICODE_SCRIPT_ARMENIAN.
	UnicodeScript_Armenian UnicodeScript = 3
	// UnicodeScript_Bengali is a representation of the C type G_UNICODE_SCRIPT_BENGALI.
	UnicodeScript_Bengali UnicodeScript = 4
	// UnicodeScript_Bopomofo is a representation of the C type G_UNICODE_SCRIPT_BOPOMOFO.
	UnicodeScript_Bopomofo UnicodeScript = 5
	// UnicodeScript_Cherokee is a representation of the C type G_UNICODE_SCRIPT_CHEROKEE.
	UnicodeScript_Cherokee UnicodeScript = 6
	// UnicodeScript_Coptic is a representation of the C type G_UNICODE_SCRIPT_COPTIC.
	UnicodeScript_Coptic UnicodeScript = 7
	// UnicodeScript_Cyrillic is a representation of the C type G_UNICODE_SCRIPT_CYRILLIC.
	UnicodeScript_Cyrillic UnicodeScript = 8
	// UnicodeScript_Deseret is a representation of the C type G_UNICODE_SCRIPT_DESERET.
	UnicodeScript_Deseret UnicodeScript = 9
	// UnicodeScript_Devanagari is a representation of the C type G_UNICODE_SCRIPT_DEVANAGARI.
	UnicodeScript_Devanagari UnicodeScript = 10
	// UnicodeScript_Ethiopic is a representation of the C type G_UNICODE_SCRIPT_ETHIOPIC.
	UnicodeScript_Ethiopic UnicodeScript = 11
	// UnicodeScript_Georgian is a representation of the C type G_UNICODE_SCRIPT_GEORGIAN.
	UnicodeScript_Georgian UnicodeScript = 12
	// UnicodeScript_Gothic is a representation of the C type G_UNICODE_SCRIPT_GOTHIC.
	UnicodeScript_Gothic UnicodeScript = 13
	// UnicodeScript_Greek is a representation of the C type G_UNICODE_SCRIPT_GREEK.
	UnicodeScript_Greek UnicodeScript = 14
	// UnicodeScript_Gujarati is a representation of the C type G_UNICODE_SCRIPT_GUJARATI.
	UnicodeScript_Gujarati UnicodeScript = 15
	// UnicodeScript_Gurmukhi is a representation of the C type G_UNICODE_SCRIPT_GURMUKHI.
	UnicodeScript_Gurmukhi UnicodeScript = 16
	// UnicodeScript_Han is a representation of the C type G_UNICODE_SCRIPT_HAN.
	UnicodeScript_Han UnicodeScript = 17
	// UnicodeScript_Hangul is a representation of the C type G_UNICODE_SCRIPT_HANGUL.
	UnicodeScript_Hangul UnicodeScript = 18
	// UnicodeScript_Hebrew is a representation of the C type G_UNICODE_SCRIPT_HEBREW.
	UnicodeScript_Hebrew UnicodeScript = 19
	// UnicodeScript_Hiragana is a representation of the C type G_UNICODE_SCRIPT_HIRAGANA.
	UnicodeScript_Hiragana UnicodeScript = 20
	// UnicodeScript_Kannada is a representation of the C type G_UNICODE_SCRIPT_KANNADA.
	UnicodeScript_Kannada UnicodeScript = 21
	// UnicodeScript_Katakana is a representation of the C type G_UNICODE_SCRIPT_KATAKANA.
	UnicodeScript_Katakana UnicodeScript = 22
	// UnicodeScript_Khmer is a representation of the C type G_UNICODE_SCRIPT_KHMER.
	UnicodeScript_Khmer UnicodeScript = 23
	// UnicodeScript_Lao is a representation of the C type G_UNICODE_SCRIPT_LAO.
	UnicodeScript_Lao UnicodeScript = 24
	// UnicodeScript_Latin is a representation of the C type G_UNICODE_SCRIPT_LATIN.
	UnicodeScript_Latin UnicodeScript = 25
	// UnicodeScript_Malayalam is a representation of the C type G_UNICODE_SCRIPT_MALAYALAM.
	UnicodeScript_Malayalam UnicodeScript = 26
	// UnicodeScript_Mongolian is a representation of the C type G_UNICODE_SCRIPT_MONGOLIAN.
	UnicodeScript_Mongolian UnicodeScript = 27
	// UnicodeScript_Myanmar is a representation of the C type G_UNICODE_SCRIPT_MYANMAR.
	UnicodeScript_Myanmar UnicodeScript = 28
	// UnicodeScript_Ogham is a representation of the C type G_UNICODE_SCRIPT_OGHAM.
	UnicodeScript_Ogham UnicodeScript = 29
	// UnicodeScript_OldItalic is a representation of the C type G_UNICODE_SCRIPT_OLD_ITALIC.
	UnicodeScript_OldItalic UnicodeScript = 30
	// UnicodeScript_Oriya is a representation of the C type G_UNICODE_SCRIPT_ORIYA.
	UnicodeScript_Oriya UnicodeScript = 31
	// UnicodeScript_Runic is a representation of the C type G_UNICODE_SCRIPT_RUNIC.
	UnicodeScript_Runic UnicodeScript = 32
	// UnicodeScript_Sinhala is a representation of the C type G_UNICODE_SCRIPT_SINHALA.
	UnicodeScript_Sinhala UnicodeScript = 33
	// UnicodeScript_Syriac is a representation of the C type G_UNICODE_SCRIPT_SYRIAC.
	UnicodeScript_Syriac UnicodeScript = 34
	// UnicodeScript_Tamil is a representation of the C type G_UNICODE_SCRIPT_TAMIL.
	UnicodeScript_Tamil UnicodeScript = 35
	// UnicodeScript_Telugu is a representation of the C type G_UNICODE_SCRIPT_TELUGU.
	UnicodeScript_Telugu UnicodeScript = 36
	// UnicodeScript_Thaana is a representation of the C type G_UNICODE_SCRIPT_THAANA.
	UnicodeScript_Thaana UnicodeScript = 37
	// UnicodeScript_Thai is a representation of the C type G_UNICODE_SCRIPT_THAI.
	UnicodeScript_Thai UnicodeScript = 38
	// UnicodeScript_Tibetan is a representation of the C type G_UNICODE_SCRIPT_TIBETAN.
	UnicodeScript_Tibetan UnicodeScript = 39
	// UnicodeScript_CanadianAboriginal is a representation of the C type G_UNICODE_SCRIPT_CANADIAN_ABORIGINAL.
	UnicodeScript_CanadianAboriginal UnicodeScript = 40
	// UnicodeScript_Yi is a representation of the C type G_UNICODE_SCRIPT_YI.
	UnicodeScript_Yi UnicodeScript = 41
	// UnicodeScript_Tagalog is a representation of the C type G_UNICODE_SCRIPT_TAGALOG.
	UnicodeScript_Tagalog UnicodeScript = 42
	// UnicodeScript_Hanunoo is a representation of the C type G_UNICODE_SCRIPT_HANUNOO.
	UnicodeScript_Hanunoo UnicodeScript = 43
	// UnicodeScript_Buhid is a representation of the C type G_UNICODE_SCRIPT_BUHID.
	UnicodeScript_Buhid UnicodeScript = 44
	// UnicodeScript_Tagbanwa is a representation of the C type G_UNICODE_SCRIPT_TAGBANWA.
	UnicodeScript_Tagbanwa UnicodeScript = 45
	// UnicodeScript_Braille is a representation of the C type G_UNICODE_SCRIPT_BRAILLE.
	UnicodeScript_Braille UnicodeScript = 46
	// UnicodeScript_Cypriot is a representation of the C type G_UNICODE_SCRIPT_CYPRIOT.
	UnicodeScript_Cypriot UnicodeScript = 47
	// UnicodeScript_Limbu is a representation of the C type G_UNICODE_SCRIPT_LIMBU.
	UnicodeScript_Limbu UnicodeScript = 48
	// UnicodeScript_Osmanya is a representation of the C type G_UNICODE_SCRIPT_OSMANYA.
	UnicodeScript_Osmanya UnicodeScript = 49
	// UnicodeScript_Shavian is a representation of the C type G_UNICODE_SCRIPT_SHAVIAN.
	UnicodeScript_Shavian UnicodeScript = 50
	// UnicodeScript_LinearB is a representation of the C type G_UNICODE_SCRIPT_LINEAR_B.
	UnicodeScript_LinearB UnicodeScript = 51
	// UnicodeScript_TaiLe is a representation of the C type G_UNICODE_SCRIPT_TAI_LE.
	UnicodeScript_TaiLe UnicodeScript = 52
	// UnicodeScript_Ugaritic is a representation of the C type G_UNICODE_SCRIPT_UGARITIC.
	UnicodeScript_Ugaritic UnicodeScript = 53
	// UnicodeScript_NewTaiLue is a representation of the C type G_UNICODE_SCRIPT_NEW_TAI_LUE.
	UnicodeScript_NewTaiLue UnicodeScript = 54
	// UnicodeScript_Buginese is a representation of the C type G_UNICODE_SCRIPT_BUGINESE.
	UnicodeScript_Buginese UnicodeScript = 55
	// UnicodeScript_Glagolitic is a representation of the C type G_UNICODE_SCRIPT_GLAGOLITIC.
	UnicodeScript_Glagolitic UnicodeScript = 56
	// UnicodeScript_Tifinagh is a representation of the C type G_UNICODE_SCRIPT_TIFINAGH.
	UnicodeScript_Tifinagh UnicodeScript = 57
	// UnicodeScript_SylotiNagri is a representation of the C type G_UNICODE_SCRIPT_SYLOTI_NAGRI.
	UnicodeScript_SylotiNagri UnicodeScript = 58
	// UnicodeScript_OldPersian is a representation of the C type G_UNICODE_SCRIPT_OLD_PERSIAN.
	UnicodeScript_OldPersian UnicodeScript = 59
	// UnicodeScript_Kharoshthi is a representation of the C type G_UNICODE_SCRIPT_KHAROSHTHI.
	UnicodeScript_Kharoshthi UnicodeScript = 60
	// UnicodeScript_Unknown is a representation of the C type G_UNICODE_SCRIPT_UNKNOWN.
	UnicodeScript_Unknown UnicodeScript = 61
	// UnicodeScript_Balinese is a representation of the C type G_UNICODE_SCRIPT_BALINESE.
	UnicodeScript_Balinese UnicodeScript = 62
	// UnicodeScript_Cuneiform is a representation of the C type G_UNICODE_SCRIPT_CUNEIFORM.
	UnicodeScript_Cuneiform UnicodeScript = 63
	// UnicodeScript_Phoenician is a representation of the C type G_UNICODE_SCRIPT_PHOENICIAN.
	UnicodeScript_Phoenician UnicodeScript = 64
	// UnicodeScript_PhagsPa is a representation of the C type G_UNICODE_SCRIPT_PHAGS_PA.
	UnicodeScript_PhagsPa UnicodeScript = 65
	// UnicodeScript_Nko is a representation of the C type G_UNICODE_SCRIPT_NKO.
	UnicodeScript_Nko UnicodeScript = 66
	// UnicodeScript_KayahLi is a representation of the C type G_UNICODE_SCRIPT_KAYAH_LI.
	UnicodeScript_KayahLi UnicodeScript = 67
	// UnicodeScript_Lepcha is a representation of the C type G_UNICODE_SCRIPT_LEPCHA.
	UnicodeScript_Lepcha UnicodeScript = 68
	// UnicodeScript_Rejang is a representation of the C type G_UNICODE_SCRIPT_REJANG.
	UnicodeScript_Rejang UnicodeScript = 69
	// UnicodeScript_Sundanese is a representation of the C type G_UNICODE_SCRIPT_SUNDANESE.
	UnicodeScript_Sundanese UnicodeScript = 70
	// UnicodeScript_Saurashtra is a representation of the C type G_UNICODE_SCRIPT_SAURASHTRA.
	UnicodeScript_Saurashtra UnicodeScript = 71
	// UnicodeScript_Cham is a representation of the C type G_UNICODE_SCRIPT_CHAM.
	UnicodeScript_Cham UnicodeScript = 72
	// UnicodeScript_OlChiki is a representation of the C type G_UNICODE_SCRIPT_OL_CHIKI.
	UnicodeScript_OlChiki UnicodeScript = 73
	// UnicodeScript_Vai is a representation of the C type G_UNICODE_SCRIPT_VAI.
	UnicodeScript_Vai UnicodeScript = 74
	// UnicodeScript_Carian is a representation of the C type G_UNICODE_SCRIPT_CARIAN.
	UnicodeScript_Carian UnicodeScript = 75
	// UnicodeScript_Lycian is a representation of the C type G_UNICODE_SCRIPT_LYCIAN.
	UnicodeScript_Lycian UnicodeScript = 76
	// UnicodeScript_Lydian is a representation of the C type G_UNICODE_SCRIPT_LYDIAN.
	UnicodeScript_Lydian UnicodeScript = 77
	// UnicodeScript_Avestan is a representation of the C type G_UNICODE_SCRIPT_AVESTAN.
	UnicodeScript_Avestan UnicodeScript = 78
	// UnicodeScript_Bamum is a representation of the C type G_UNICODE_SCRIPT_BAMUM.
	UnicodeScript_Bamum UnicodeScript = 79
	// UnicodeScript_EgyptianHieroglyphs is a representation of the C type G_UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS.
	UnicodeScript_EgyptianHieroglyphs UnicodeScript = 80
	// UnicodeScript_ImperialAramaic is a representation of the C type G_UNICODE_SCRIPT_IMPERIAL_ARAMAIC.
	UnicodeScript_ImperialAramaic UnicodeScript = 81
	// UnicodeScript_InscriptionalPahlavi is a representation of the C type G_UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI.
	UnicodeScript_InscriptionalPahlavi UnicodeScript = 82
	// UnicodeScript_InscriptionalParthian is a representation of the C type G_UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN.
	UnicodeScript_InscriptionalParthian UnicodeScript = 83
	// UnicodeScript_Javanese is a representation of the C type G_UNICODE_SCRIPT_JAVANESE.
	UnicodeScript_Javanese UnicodeScript = 84
	// UnicodeScript_Kaithi is a representation of the C type G_UNICODE_SCRIPT_KAITHI.
	UnicodeScript_Kaithi UnicodeScript = 85
	// UnicodeScript_Lisu is a representation of the C type G_UNICODE_SCRIPT_LISU.
	UnicodeScript_Lisu UnicodeScript = 86
	// UnicodeScript_MeeteiMayek is a representation of the C type G_UNICODE_SCRIPT_MEETEI_MAYEK.
	UnicodeScript_MeeteiMayek UnicodeScript = 87
	// UnicodeScript_OldSouthArabian is a representation of the C type G_UNICODE_SCRIPT_OLD_SOUTH_ARABIAN.
	UnicodeScript_OldSouthArabian UnicodeScript = 88
	// UnicodeScript_OldTurkic is a representation of the C type G_UNICODE_SCRIPT_OLD_TURKIC.
	UnicodeScript_OldTurkic UnicodeScript = 89
	// UnicodeScript_Samaritan is a representation of the C type G_UNICODE_SCRIPT_SAMARITAN.
	UnicodeScript_Samaritan UnicodeScript = 90
	// UnicodeScript_TaiTham is a representation of the C type G_UNICODE_SCRIPT_TAI_THAM.
	UnicodeScript_TaiTham UnicodeScript = 91
	// UnicodeScript_TaiViet is a representation of the C type G_UNICODE_SCRIPT_TAI_VIET.
	UnicodeScript_TaiViet UnicodeScript = 92
	// UnicodeScript_Batak is a representation of the C type G_UNICODE_SCRIPT_BATAK.
	UnicodeScript_Batak UnicodeScript = 93
	// UnicodeScript_Brahmi is a representation of the C type G_UNICODE_SCRIPT_BRAHMI.
	UnicodeScript_Brahmi UnicodeScript = 94
	// UnicodeScript_Mandaic is a representation of the C type G_UNICODE_SCRIPT_MANDAIC.
	UnicodeScript_Mandaic UnicodeScript = 95
	// UnicodeScript_Chakma is a representation of the C type G_UNICODE_SCRIPT_CHAKMA.
	UnicodeScript_Chakma UnicodeScript = 96
	// UnicodeScript_MeroiticCursive is a representation of the C type G_UNICODE_SCRIPT_MEROITIC_CURSIVE.
	UnicodeScript_MeroiticCursive UnicodeScript = 97
	// UnicodeScript_MeroiticHieroglyphs is a representation of the C type G_UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS.
	UnicodeScript_MeroiticHieroglyphs UnicodeScript = 98
	// UnicodeScript_Miao is a representation of the C type G_UNICODE_SCRIPT_MIAO.
	UnicodeScript_Miao UnicodeScript = 99
	// UnicodeScript_Sharada is a representation of the C type G_UNICODE_SCRIPT_SHARADA.
	UnicodeScript_Sharada UnicodeScript = 100
	// UnicodeScript_SoraSompeng is a representation of the C type G_UNICODE_SCRIPT_SORA_SOMPENG.
	UnicodeScript_SoraSompeng UnicodeScript = 101
	// UnicodeScript_Takri is a representation of the C type G_UNICODE_SCRIPT_TAKRI.
	UnicodeScript_Takri UnicodeScript = 102
	// UnicodeScript_BassaVah is a representation of the C type G_UNICODE_SCRIPT_BASSA_VAH.
	UnicodeScript_BassaVah UnicodeScript = 103
	// UnicodeScript_CaucasianAlbanian is a representation of the C type G_UNICODE_SCRIPT_CAUCASIAN_ALBANIAN.
	UnicodeScript_CaucasianAlbanian UnicodeScript = 104
	// UnicodeScript_Duployan is a representation of the C type G_UNICODE_SCRIPT_DUPLOYAN.
	UnicodeScript_Duployan UnicodeScript = 105
	// UnicodeScript_Elbasan is a representation of the C type G_UNICODE_SCRIPT_ELBASAN.
	UnicodeScript_Elbasan UnicodeScript = 106
	// UnicodeScript_Grantha is a representation of the C type G_UNICODE_SCRIPT_GRANTHA.
	UnicodeScript_Grantha UnicodeScript = 107
	// UnicodeScript_Khojki is a representation of the C type G_UNICODE_SCRIPT_KHOJKI.
	UnicodeScript_Khojki UnicodeScript = 108
	// UnicodeScript_Khudawadi is a representation of the C type G_UNICODE_SCRIPT_KHUDAWADI.
	UnicodeScript_Khudawadi UnicodeScript = 109
	// UnicodeScript_LinearA is a representation of the C type G_UNICODE_SCRIPT_LINEAR_A.
	UnicodeScript_LinearA UnicodeScript = 110
	// UnicodeScript_Mahajani is a representation of the C type G_UNICODE_SCRIPT_MAHAJANI.
	UnicodeScript_Mahajani UnicodeScript = 111
	// UnicodeScript_Manichaean is a representation of the C type G_UNICODE_SCRIPT_MANICHAEAN.
	UnicodeScript_Manichaean UnicodeScript = 112
	// UnicodeScript_MendeKikakui is a representation of the C type G_UNICODE_SCRIPT_MENDE_KIKAKUI.
	UnicodeScript_MendeKikakui UnicodeScript = 113
	// UnicodeScript_Modi is a representation of the C type G_UNICODE_SCRIPT_MODI.
	UnicodeScript_Modi UnicodeScript = 114
	// UnicodeScript_Mro is a representation of the C type G_UNICODE_SCRIPT_MRO.
	UnicodeScript_Mro UnicodeScript = 115
	// UnicodeScript_Nabataean is a representation of the C type G_UNICODE_SCRIPT_NABATAEAN.
	UnicodeScript_Nabataean UnicodeScript = 116
	// UnicodeScript_OldNorthArabian is a representation of the C type G_UNICODE_SCRIPT_OLD_NORTH_ARABIAN.
	UnicodeScript_OldNorthArabian UnicodeScript = 117
	// UnicodeScript_OldPermic is a representation of the C type G_UNICODE_SCRIPT_OLD_PERMIC.
	UnicodeScript_OldPermic UnicodeScript = 118
	// UnicodeScript_PahawhHmong is a representation of the C type G_UNICODE_SCRIPT_PAHAWH_HMONG.
	UnicodeScript_PahawhHmong UnicodeScript = 119
	// UnicodeScript_Palmyrene is a representation of the C type G_UNICODE_SCRIPT_PALMYRENE.
	UnicodeScript_Palmyrene UnicodeScript = 120
	// UnicodeScript_PauCinHau is a representation of the C type G_UNICODE_SCRIPT_PAU_CIN_HAU.
	UnicodeScript_PauCinHau UnicodeScript = 121
	// UnicodeScript_PsalterPahlavi is a representation of the C type G_UNICODE_SCRIPT_PSALTER_PAHLAVI.
	UnicodeScript_PsalterPahlavi UnicodeScript = 122
	// UnicodeScript_Siddham is a representation of the C type G_UNICODE_SCRIPT_SIDDHAM.
	UnicodeScript_Siddham UnicodeScript = 123
	// UnicodeScript_Tirhuta is a representation of the C type G_UNICODE_SCRIPT_TIRHUTA.
	UnicodeScript_Tirhuta UnicodeScript = 124
	// UnicodeScript_WarangCiti is a representation of the C type G_UNICODE_SCRIPT_WARANG_CITI.
	UnicodeScript_WarangCiti UnicodeScript = 125
	// UnicodeScript_Ahom is a representation of the C type G_UNICODE_SCRIPT_AHOM.
	UnicodeScript_Ahom UnicodeScript = 126
	// UnicodeScript_AnatolianHieroglyphs is a representation of the C type G_UNICODE_SCRIPT_ANATOLIAN_HIEROGLYPHS.
	UnicodeScript_AnatolianHieroglyphs UnicodeScript = 127
	// UnicodeScript_Hatran is a representation of the C type G_UNICODE_SCRIPT_HATRAN.
	UnicodeScript_Hatran UnicodeScript = 128
	// UnicodeScript_Multani is a representation of the C type G_UNICODE_SCRIPT_MULTANI.
	UnicodeScript_Multani UnicodeScript = 129
	// UnicodeScript_OldHungarian is a representation of the C type G_UNICODE_SCRIPT_OLD_HUNGARIAN.
	UnicodeScript_OldHungarian UnicodeScript = 130
	// UnicodeScript_Signwriting is a representation of the C type G_UNICODE_SCRIPT_SIGNWRITING.
	UnicodeScript_Signwriting UnicodeScript = 131
	// UnicodeScript_Adlam is a representation of the C type G_UNICODE_SCRIPT_ADLAM.
	UnicodeScript_Adlam UnicodeScript = 132
	// UnicodeScript_Bhaiksuki is a representation of the C type G_UNICODE_SCRIPT_BHAIKSUKI.
	UnicodeScript_Bhaiksuki UnicodeScript = 133
	// UnicodeScript_Marchen is a representation of the C type G_UNICODE_SCRIPT_MARCHEN.
	UnicodeScript_Marchen UnicodeScript = 134
	// UnicodeScript_Newa is a representation of the C type G_UNICODE_SCRIPT_NEWA.
	UnicodeScript_Newa UnicodeScript = 135
	// UnicodeScript_Osage is a representation of the C type G_UNICODE_SCRIPT_OSAGE.
	UnicodeScript_Osage UnicodeScript = 136
	// UnicodeScript_Tangut is a representation of the C type G_UNICODE_SCRIPT_TANGUT.
	UnicodeScript_Tangut UnicodeScript = 137
	// UnicodeScript_MasaramGondi is a representation of the C type G_UNICODE_SCRIPT_MASARAM_GONDI.
	UnicodeScript_MasaramGondi UnicodeScript = 138
	// UnicodeScript_Nushu is a representation of the C type G_UNICODE_SCRIPT_NUSHU.
	UnicodeScript_Nushu UnicodeScript = 139
	// UnicodeScript_Soyombo is a representation of the C type G_UNICODE_SCRIPT_SOYOMBO.
	UnicodeScript_Soyombo UnicodeScript = 140
	// UnicodeScript_ZanabazarSquare is a representation of the C type G_UNICODE_SCRIPT_ZANABAZAR_SQUARE.
	UnicodeScript_ZanabazarSquare UnicodeScript = 141
	// UnicodeScript_Dogra is a representation of the C type G_UNICODE_SCRIPT_DOGRA.
	UnicodeScript_Dogra UnicodeScript = 142
	// UnicodeScript_GunjalaGondi is a representation of the C type G_UNICODE_SCRIPT_GUNJALA_GONDI.
	UnicodeScript_GunjalaGondi UnicodeScript = 143
	// UnicodeScript_HanifiRohingya is a representation of the C type G_UNICODE_SCRIPT_HANIFI_ROHINGYA.
	UnicodeScript_HanifiRohingya UnicodeScript = 144
	// UnicodeScript_Makasar is a representation of the C type G_UNICODE_SCRIPT_MAKASAR.
	UnicodeScript_Makasar UnicodeScript = 145
	// UnicodeScript_Medefaidrin is a representation of the C type G_UNICODE_SCRIPT_MEDEFAIDRIN.
	UnicodeScript_Medefaidrin UnicodeScript = 146
	// UnicodeScript_OldSogdian is a representation of the C type G_UNICODE_SCRIPT_OLD_SOGDIAN.
	UnicodeScript_OldSogdian UnicodeScript = 147
	// UnicodeScript_Sogdian is a representation of the C type G_UNICODE_SCRIPT_SOGDIAN.
	UnicodeScript_Sogdian UnicodeScript = 148
	// UnicodeScript_Elymaic is a representation of the C type G_UNICODE_SCRIPT_ELYMAIC.
	UnicodeScript_Elymaic UnicodeScript = 149
	// UnicodeScript_Nandinagari is a representation of the C type G_UNICODE_SCRIPT_NANDINAGARI.
	UnicodeScript_Nandinagari UnicodeScript = 150
	// UnicodeScript_NyiakengPuachueHmong is a representation of the C type G_UNICODE_SCRIPT_NYIAKENG_PUACHUE_HMONG.
	UnicodeScript_NyiakengPuachueHmong UnicodeScript = 151
	// UnicodeScript_Wancho is a representation of the C type G_UNICODE_SCRIPT_WANCHO.
	UnicodeScript_Wancho UnicodeScript = 152
)

// UnicodeType is a representation of the C type GUnicodeType.
//
type UnicodeType int

const (
	// UnicodeType_Control is a representation of the C type G_UNICODE_CONTROL.
	UnicodeType_Control UnicodeType = 0
	// UnicodeType_Format is a representation of the C type G_UNICODE_FORMAT.
	UnicodeType_Format UnicodeType = 1
	// UnicodeType_Unassigned is a representation of the C type G_UNICODE_UNASSIGNED.
	UnicodeType_Unassigned UnicodeType = 2
	// UnicodeType_PrivateUse is a representation of the C type G_UNICODE_PRIVATE_USE.
	UnicodeType_PrivateUse UnicodeType = 3
	// UnicodeType_Surrogate is a representation of the C type G_UNICODE_SURROGATE.
	UnicodeType_Surrogate UnicodeType = 4
	// UnicodeType_LowercaseLetter is a representation of the C type G_UNICODE_LOWERCASE_LETTER.
	UnicodeType_LowercaseLetter UnicodeType = 5
	// UnicodeType_ModifierLetter is a representation of the C type G_UNICODE_MODIFIER_LETTER.
	UnicodeType_ModifierLetter UnicodeType = 6
	// UnicodeType_OtherLetter is a representation of the C type G_UNICODE_OTHER_LETTER.
	UnicodeType_OtherLetter UnicodeType = 7
	// UnicodeType_TitlecaseLetter is a representation of the C type G_UNICODE_TITLECASE_LETTER.
	UnicodeType_TitlecaseLetter UnicodeType = 8
	// UnicodeType_UppercaseLetter is a representation of the C type G_UNICODE_UPPERCASE_LETTER.
	UnicodeType_UppercaseLetter UnicodeType = 9
	// UnicodeType_SpacingMark is a representation of the C type G_UNICODE_SPACING_MARK.
	UnicodeType_SpacingMark UnicodeType = 10
	// UnicodeType_EnclosingMark is a representation of the C type G_UNICODE_ENCLOSING_MARK.
	UnicodeType_EnclosingMark UnicodeType = 11
	// UnicodeType_NonSpacingMark is a representation of the C type G_UNICODE_NON_SPACING_MARK.
	UnicodeType_NonSpacingMark UnicodeType = 12
	// UnicodeType_DecimalNumber is a representation of the C type G_UNICODE_DECIMAL_NUMBER.
	UnicodeType_DecimalNumber UnicodeType = 13
	// UnicodeType_LetterNumber is a representation of the C type G_UNICODE_LETTER_NUMBER.
	UnicodeType_LetterNumber UnicodeType = 14
	// UnicodeType_OtherNumber is a representation of the C type G_UNICODE_OTHER_NUMBER.
	UnicodeType_OtherNumber UnicodeType = 15
	// UnicodeType_ConnectPunctuation is a representation of the C type G_UNICODE_CONNECT_PUNCTUATION.
	UnicodeType_ConnectPunctuation UnicodeType = 16
	// UnicodeType_DashPunctuation is a representation of the C type G_UNICODE_DASH_PUNCTUATION.
	UnicodeType_DashPunctuation UnicodeType = 17
	// UnicodeType_ClosePunctuation is a representation of the C type G_UNICODE_CLOSE_PUNCTUATION.
	UnicodeType_ClosePunctuation UnicodeType = 18
	// UnicodeType_FinalPunctuation is a representation of the C type G_UNICODE_FINAL_PUNCTUATION.
	UnicodeType_FinalPunctuation UnicodeType = 19
	// UnicodeType_InitialPunctuation is a representation of the C type G_UNICODE_INITIAL_PUNCTUATION.
	UnicodeType_InitialPunctuation UnicodeType = 20
	// UnicodeType_OtherPunctuation is a representation of the C type G_UNICODE_OTHER_PUNCTUATION.
	UnicodeType_OtherPunctuation UnicodeType = 21
	// UnicodeType_OpenPunctuation is a representation of the C type G_UNICODE_OPEN_PUNCTUATION.
	UnicodeType_OpenPunctuation UnicodeType = 22
	// UnicodeType_CurrencySymbol is a representation of the C type G_UNICODE_CURRENCY_SYMBOL.
	UnicodeType_CurrencySymbol UnicodeType = 23
	// UnicodeType_ModifierSymbol is a representation of the C type G_UNICODE_MODIFIER_SYMBOL.
	UnicodeType_ModifierSymbol UnicodeType = 24
	// UnicodeType_MathSymbol is a representation of the C type G_UNICODE_MATH_SYMBOL.
	UnicodeType_MathSymbol UnicodeType = 25
	// UnicodeType_OtherSymbol is a representation of the C type G_UNICODE_OTHER_SYMBOL.
	UnicodeType_OtherSymbol UnicodeType = 26
	// UnicodeType_LineSeparator is a representation of the C type G_UNICODE_LINE_SEPARATOR.
	UnicodeType_LineSeparator UnicodeType = 27
	// UnicodeType_ParagraphSeparator is a representation of the C type G_UNICODE_PARAGRAPH_SEPARATOR.
	UnicodeType_ParagraphSeparator UnicodeType = 28
	// UnicodeType_SpaceSeparator is a representation of the C type G_UNICODE_SPACE_SEPARATOR.
	UnicodeType_SpaceSeparator UnicodeType = 29
)

// UserDirectory is a representation of the C type GUserDirectory.
//
// since 2.14
type UserDirectory int

const (
	// UserDirectory_DirectoryDesktop is a representation of the C type G_USER_DIRECTORY_DESKTOP.
	UserDirectory_DirectoryDesktop UserDirectory = 0
	// UserDirectory_DirectoryDocuments is a representation of the C type G_USER_DIRECTORY_DOCUMENTS.
	UserDirectory_DirectoryDocuments UserDirectory = 1
	// UserDirectory_DirectoryDownload is a representation of the C type G_USER_DIRECTORY_DOWNLOAD.
	UserDirectory_DirectoryDownload UserDirectory = 2
	// UserDirectory_DirectoryMusic is a representation of the C type G_USER_DIRECTORY_MUSIC.
	UserDirectory_DirectoryMusic UserDirectory = 3
	// UserDirectory_DirectoryPictures is a representation of the C type G_USER_DIRECTORY_PICTURES.
	UserDirectory_DirectoryPictures UserDirectory = 4
	// UserDirectory_DirectoryPublicShare is a representation of the C type G_USER_DIRECTORY_PUBLIC_SHARE.
	UserDirectory_DirectoryPublicShare UserDirectory = 5
	// UserDirectory_DirectoryTemplates is a representation of the C type G_USER_DIRECTORY_TEMPLATES.
	UserDirectory_DirectoryTemplates UserDirectory = 6
	// UserDirectory_DirectoryVideos is a representation of the C type G_USER_DIRECTORY_VIDEOS.
	UserDirectory_DirectoryVideos UserDirectory = 7
	// UserDirectory_NDirectories is a representation of the C type G_USER_N_DIRECTORIES.
	UserDirectory_NDirectories UserDirectory = 8
)

// VariantClass is a representation of the C type GVariantClass.
//
// since 2.24
type VariantClass int

const (
	// VariantClass_Boolean is a representation of the C type G_VARIANT_CLASS_BOOLEAN.
	VariantClass_Boolean VariantClass = 98
	// VariantClass_Byte is a representation of the C type G_VARIANT_CLASS_BYTE.
	VariantClass_Byte VariantClass = 121
	// VariantClass_Int16 is a representation of the C type G_VARIANT_CLASS_INT16.
	VariantClass_Int16 VariantClass = 110
	// VariantClass_Uint16 is a representation of the C type G_VARIANT_CLASS_UINT16.
	VariantClass_Uint16 VariantClass = 113
	// VariantClass_Int32 is a representation of the C type G_VARIANT_CLASS_INT32.
	VariantClass_Int32 VariantClass = 105
	// VariantClass_Uint32 is a representation of the C type G_VARIANT_CLASS_UINT32.
	VariantClass_Uint32 VariantClass = 117
	// VariantClass_Int64 is a representation of the C type G_VARIANT_CLASS_INT64.
	VariantClass_Int64 VariantClass = 120
	// VariantClass_Uint64 is a representation of the C type G_VARIANT_CLASS_UINT64.
	VariantClass_Uint64 VariantClass = 116
	// VariantClass_Handle is a representation of the C type G_VARIANT_CLASS_HANDLE.
	VariantClass_Handle VariantClass = 104
	// VariantClass_Double is a representation of the C type G_VARIANT_CLASS_DOUBLE.
	VariantClass_Double VariantClass = 100
	// VariantClass_String is a representation of the C type G_VARIANT_CLASS_STRING.
	VariantClass_String VariantClass = 115
	// VariantClass_ObjectPath is a representation of the C type G_VARIANT_CLASS_OBJECT_PATH.
	VariantClass_ObjectPath VariantClass = 111
	// VariantClass_Signature is a representation of the C type G_VARIANT_CLASS_SIGNATURE.
	VariantClass_Signature VariantClass = 103
	// VariantClass_Variant is a representation of the C type G_VARIANT_CLASS_VARIANT.
	VariantClass_Variant VariantClass = 118
	// VariantClass_Maybe is a representation of the C type G_VARIANT_CLASS_MAYBE.
	VariantClass_Maybe VariantClass = 109
	// VariantClass_Array is a representation of the C type G_VARIANT_CLASS_ARRAY.
	VariantClass_Array VariantClass = 97
	// VariantClass_Tuple is a representation of the C type G_VARIANT_CLASS_TUPLE.
	VariantClass_Tuple VariantClass = 40
	// VariantClass_DictEntry is a representation of the C type G_VARIANT_CLASS_DICT_ENTRY.
	VariantClass_DictEntry VariantClass = 123
)

// VariantParseError is a representation of the C type GVariantParseError.
//
type VariantParseError int

const (
	// VariantParseError_Failed is a representation of the C type G_VARIANT_PARSE_ERROR_FAILED.
	VariantParseError_Failed VariantParseError = 0
	// VariantParseError_BasicTypeExpected is a representation of the C type G_VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED.
	VariantParseError_BasicTypeExpected VariantParseError = 1
	// VariantParseError_CannotInferType is a representation of the C type G_VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE.
	VariantParseError_CannotInferType VariantParseError = 2
	// VariantParseError_DefiniteTypeExpected is a representation of the C type G_VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED.
	VariantParseError_DefiniteTypeExpected VariantParseError = 3
	// VariantParseError_InputNotAtEnd is a representation of the C type G_VARIANT_PARSE_ERROR_INPUT_NOT_AT_END.
	VariantParseError_InputNotAtEnd VariantParseError = 4
	// VariantParseError_InvalidCharacter is a representation of the C type G_VARIANT_PARSE_ERROR_INVALID_CHARACTER.
	VariantParseError_InvalidCharacter VariantParseError = 5
	// VariantParseError_InvalidFormatString is a representation of the C type G_VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING.
	VariantParseError_InvalidFormatString VariantParseError = 6
	// VariantParseError_InvalidObjectPath is a representation of the C type G_VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH.
	VariantParseError_InvalidObjectPath VariantParseError = 7
	// VariantParseError_InvalidSignature is a representation of the C type G_VARIANT_PARSE_ERROR_INVALID_SIGNATURE.
	VariantParseError_InvalidSignature VariantParseError = 8
	// VariantParseError_InvalidTypeString is a representation of the C type G_VARIANT_PARSE_ERROR_INVALID_TYPE_STRING.
	VariantParseError_InvalidTypeString VariantParseError = 9
	// VariantParseError_NoCommonType is a representation of the C type G_VARIANT_PARSE_ERROR_NO_COMMON_TYPE.
	VariantParseError_NoCommonType VariantParseError = 10
	// VariantParseError_NumberOutOfRange is a representation of the C type G_VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE.
	VariantParseError_NumberOutOfRange VariantParseError = 11
	// VariantParseError_NumberTooBig is a representation of the C type G_VARIANT_PARSE_ERROR_NUMBER_TOO_BIG.
	VariantParseError_NumberTooBig VariantParseError = 12
	// VariantParseError_TypeError is a representation of the C type G_VARIANT_PARSE_ERROR_TYPE_ERROR.
	VariantParseError_TypeError VariantParseError = 13
	// VariantParseError_UnexpectedToken is a representation of the C type G_VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN.
	VariantParseError_UnexpectedToken VariantParseError = 14
	// VariantParseError_UnknownKeyword is a representation of the C type G_VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD.
	VariantParseError_UnknownKeyword VariantParseError = 15
	// VariantParseError_UnterminatedStringConstant is a representation of the C type G_VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT.
	VariantParseError_UnterminatedStringConstant VariantParseError = 16
	// VariantParseError_ValueExpected is a representation of the C type G_VARIANT_PARSE_ERROR_VALUE_EXPECTED.
	VariantParseError_ValueExpected VariantParseError = 17
)

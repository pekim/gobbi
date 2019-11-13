// Code generated - DO NOT EDIT.

package glib

// Bookmarkfileerror is a representation of the C type BookmarkFileError.
type Bookmarkfileerror int

const (
	// invalid_uri
	GBookmarkFileErrorInvalidUri Bookmarkfileerror = 0
	// invalid_value
	GBookmarkFileErrorInvalidValue Bookmarkfileerror = 1
	// app_not_registered
	GBookmarkFileErrorAppNotRegistered Bookmarkfileerror = 2
	// uri_not_found
	GBookmarkFileErrorUriNotFound Bookmarkfileerror = 3
	// read
	GBookmarkFileErrorRead Bookmarkfileerror = 4
	// unknown_encoding
	GBookmarkFileErrorUnknownEncoding Bookmarkfileerror = 5
	// write
	GBookmarkFileErrorWrite Bookmarkfileerror = 6
	// file_not_found
	GBookmarkFileErrorFileNotFound Bookmarkfileerror = 7
)

// Checksumtype is a representation of the C type ChecksumType.
//
// since 2.16
type Checksumtype int

const (
	// md5
	GChecksumMd5 Checksumtype = 0
	// sha1
	GChecksumSha1 Checksumtype = 1
	// sha256
	GChecksumSha256 Checksumtype = 2
	// sha512
	GChecksumSha512 Checksumtype = 3
	// sha384
	GChecksumSha384 Checksumtype = 4
)

// Converterror is a representation of the C type ConvertError.
type Converterror int

const (
	// no_conversion
	GConvertErrorNoConversion Converterror = 0
	// illegal_sequence
	GConvertErrorIllegalSequence Converterror = 1
	// failed
	GConvertErrorFailed Converterror = 2
	// partial_input
	GConvertErrorPartialInput Converterror = 3
	// bad_uri
	GConvertErrorBadUri Converterror = 4
	// not_absolute_path
	GConvertErrorNotAbsolutePath Converterror = 5
	// no_memory
	GConvertErrorNoMemory Converterror = 6
	// embedded_nul
	GConvertErrorEmbeddedNul Converterror = 7
)

// Datedmy is a representation of the C type DateDMY.
type Datedmy int

const (
	// day
	GDateDay Datedmy = 0
	// month
	GDateMonth Datedmy = 1
	// year
	GDateYear Datedmy = 2
)

// Datemonth is a representation of the C type DateMonth.
type Datemonth int

const (
	// bad_month
	GDateBadMonth Datemonth = 0
	// january
	GDateJanuary Datemonth = 1
	// february
	GDateFebruary Datemonth = 2
	// march
	GDateMarch Datemonth = 3
	// april
	GDateApril Datemonth = 4
	// may
	GDateMay Datemonth = 5
	// june
	GDateJune Datemonth = 6
	// july
	GDateJuly Datemonth = 7
	// august
	GDateAugust Datemonth = 8
	// september
	GDateSeptember Datemonth = 9
	// october
	GDateOctober Datemonth = 10
	// november
	GDateNovember Datemonth = 11
	// december
	GDateDecember Datemonth = 12
)

// Dateweekday is a representation of the C type DateWeekday.
type Dateweekday int

const (
	// bad_weekday
	GDateBadWeekday Dateweekday = 0
	// monday
	GDateMonday Dateweekday = 1
	// tuesday
	GDateTuesday Dateweekday = 2
	// wednesday
	GDateWednesday Dateweekday = 3
	// thursday
	GDateThursday Dateweekday = 4
	// friday
	GDateFriday Dateweekday = 5
	// saturday
	GDateSaturday Dateweekday = 6
	// sunday
	GDateSunday Dateweekday = 7
)

// Errortype is a representation of the C type ErrorType.
type Errortype int

const (
	// unknown
	GErrUnknown Errortype = 0
	// unexp_eof
	GErrUnexpEof Errortype = 1
	// unexp_eof_in_string
	GErrUnexpEofInString Errortype = 2
	// unexp_eof_in_comment
	GErrUnexpEofInComment Errortype = 3
	// non_digit_in_const
	GErrNonDigitInConst Errortype = 4
	// digit_radix
	GErrDigitRadix Errortype = 5
	// float_radix
	GErrFloatRadix Errortype = 6
	// float_malformed
	GErrFloatMalformed Errortype = 7
)

// Fileerror is a representation of the C type FileError.
type Fileerror int

const (
	// exist
	GFileErrorExist Fileerror = 0
	// isdir
	GFileErrorIsdir Fileerror = 1
	// acces
	GFileErrorAcces Fileerror = 2
	// nametoolong
	GFileErrorNametoolong Fileerror = 3
	// noent
	GFileErrorNoent Fileerror = 4
	// notdir
	GFileErrorNotdir Fileerror = 5
	// nxio
	GFileErrorNxio Fileerror = 6
	// nodev
	GFileErrorNodev Fileerror = 7
	// rofs
	GFileErrorRofs Fileerror = 8
	// txtbsy
	GFileErrorTxtbsy Fileerror = 9
	// fault
	GFileErrorFault Fileerror = 10
	// loop
	GFileErrorLoop Fileerror = 11
	// nospc
	GFileErrorNospc Fileerror = 12
	// nomem
	GFileErrorNomem Fileerror = 13
	// mfile
	GFileErrorMfile Fileerror = 14
	// nfile
	GFileErrorNfile Fileerror = 15
	// badf
	GFileErrorBadf Fileerror = 16
	// inval
	GFileErrorInval Fileerror = 17
	// pipe
	GFileErrorPipe Fileerror = 18
	// again
	GFileErrorAgain Fileerror = 19
	// intr
	GFileErrorIntr Fileerror = 20
	// io
	GFileErrorIo Fileerror = 21
	// perm
	GFileErrorPerm Fileerror = 22
	// nosys
	GFileErrorNosys Fileerror = 23
	// failed
	GFileErrorFailed Fileerror = 24
)

// Iochannelerror is a representation of the C type IOChannelError.
type Iochannelerror int

const (
	// fbig
	GIoChannelErrorFbig Iochannelerror = 0
	// inval
	GIoChannelErrorInval Iochannelerror = 1
	// io
	GIoChannelErrorIo Iochannelerror = 2
	// isdir
	GIoChannelErrorIsdir Iochannelerror = 3
	// nospc
	GIoChannelErrorNospc Iochannelerror = 4
	// nxio
	GIoChannelErrorNxio Iochannelerror = 5
	// overflow
	GIoChannelErrorOverflow Iochannelerror = 6
	// pipe
	GIoChannelErrorPipe Iochannelerror = 7
	// failed
	GIoChannelErrorFailed Iochannelerror = 8
)

// Ioerror is a representation of the C type IOError.
type Ioerror int

const (
	// none
	GIoErrorNone Ioerror = 0
	// again
	GIoErrorAgain Ioerror = 1
	// inval
	GIoErrorInval Ioerror = 2
	// unknown
	GIoErrorUnknown Ioerror = 3
)

// Iostatus is a representation of the C type IOStatus.
type Iostatus int

const (
	// error
	GIoStatusError Iostatus = 0
	// normal
	GIoStatusNormal Iostatus = 1
	// eof
	GIoStatusEof Iostatus = 2
	// again
	GIoStatusAgain Iostatus = 3
)

// Keyfileerror is a representation of the C type KeyFileError.
type Keyfileerror int

const (
	// unknown_encoding
	GKeyFileErrorUnknownEncoding Keyfileerror = 0
	// parse
	GKeyFileErrorParse Keyfileerror = 1
	// not_found
	GKeyFileErrorNotFound Keyfileerror = 2
	// key_not_found
	GKeyFileErrorKeyNotFound Keyfileerror = 3
	// group_not_found
	GKeyFileErrorGroupNotFound Keyfileerror = 4
	// invalid_value
	GKeyFileErrorInvalidValue Keyfileerror = 5
)

// Logwriteroutput is a representation of the C type LogWriterOutput.
//
// since 2.50
type Logwriteroutput int

const (
	// handled
	GLogWriterHandled Logwriteroutput = 1
	// unhandled
	GLogWriterUnhandled Logwriteroutput = 0
)

// Markuperror is a representation of the C type MarkupError.
type Markuperror int

const (
	// bad_utf8
	GMarkupErrorBadUtf8 Markuperror = 0
	// empty
	GMarkupErrorEmpty Markuperror = 1
	// parse
	GMarkupErrorParse Markuperror = 2
	// unknown_element
	GMarkupErrorUnknownElement Markuperror = 3
	// unknown_attribute
	GMarkupErrorUnknownAttribute Markuperror = 4
	// invalid_content
	GMarkupErrorInvalidContent Markuperror = 5
	// missing_attribute
	GMarkupErrorMissingAttribute Markuperror = 6
)

// Normalizemode is a representation of the C type NormalizeMode.
type Normalizemode int

const (
	// default
	GNormalizeDefault Normalizemode = 0
	// nfd
	GNormalizeNfd Normalizemode = 0
	// default_compose
	GNormalizeDefaultCompose Normalizemode = 1
	// nfc
	GNormalizeNfc Normalizemode = 1
	// all
	GNormalizeAll Normalizemode = 2
	// nfkd
	GNormalizeNfkd Normalizemode = 2
	// all_compose
	GNormalizeAllCompose Normalizemode = 3
	// nfkc
	GNormalizeNfkc Normalizemode = 3
)

// Numberparsererror is a representation of the C type NumberParserError.
//
// since 2.54
type Numberparsererror int

const (
	// invalid
	GNumberParserErrorInvalid Numberparsererror = 0
	// out_of_bounds
	GNumberParserErrorOutOfBounds Numberparsererror = 1
)

// Oncestatus is a representation of the C type OnceStatus.
//
// since 2.4
type Oncestatus int

const (
	// notcalled
	GOnceStatusNotcalled Oncestatus = 0
	// progress
	GOnceStatusProgress Oncestatus = 1
	// ready
	GOnceStatusReady Oncestatus = 2
)

// Optionarg is a representation of the C type OptionArg.
type Optionarg int

const (
	// none
	GOptionArgNone Optionarg = 0
	// string
	GOptionArgString Optionarg = 1
	// int
	GOptionArgInt Optionarg = 2
	// callback
	GOptionArgCallback Optionarg = 3
	// filename
	GOptionArgFilename Optionarg = 4
	// string_array
	GOptionArgStringArray Optionarg = 5
	// filename_array
	GOptionArgFilenameArray Optionarg = 6
	// double
	GOptionArgDouble Optionarg = 7
	// int64
	GOptionArgInt64 Optionarg = 8
)

// Optionerror is a representation of the C type OptionError.
type Optionerror int

const (
	// unknown_option
	GOptionErrorUnknownOption Optionerror = 0
	// bad_value
	GOptionErrorBadValue Optionerror = 1
	// failed
	GOptionErrorFailed Optionerror = 2
)

// Regexerror is a representation of the C type RegexError.
//
// since 2.14
type Regexerror int

const (
	// compile
	GRegexErrorCompile Regexerror = 0
	// optimize
	GRegexErrorOptimize Regexerror = 1
	// replace
	GRegexErrorReplace Regexerror = 2
	// match
	GRegexErrorMatch Regexerror = 3
	// internal
	GRegexErrorInternal Regexerror = 4
	// stray_backslash
	GRegexErrorStrayBackslash Regexerror = 101
	// missing_control_char
	GRegexErrorMissingControlChar Regexerror = 102
	// unrecognized_escape
	GRegexErrorUnrecognizedEscape Regexerror = 103
	// quantifiers_out_of_order
	GRegexErrorQuantifiersOutOfOrder Regexerror = 104
	// quantifier_too_big
	GRegexErrorQuantifierTooBig Regexerror = 105
	// unterminated_character_class
	GRegexErrorUnterminatedCharacterClass Regexerror = 106
	// invalid_escape_in_character_class
	GRegexErrorInvalidEscapeInCharacterClass Regexerror = 107
	// range_out_of_order
	GRegexErrorRangeOutOfOrder Regexerror = 108
	// nothing_to_repeat
	GRegexErrorNothingToRepeat Regexerror = 109
	// unrecognized_character
	GRegexErrorUnrecognizedCharacter Regexerror = 112
	// posix_named_class_outside_class
	GRegexErrorPosixNamedClassOutsideClass Regexerror = 113
	// unmatched_parenthesis
	GRegexErrorUnmatchedParenthesis Regexerror = 114
	// inexistent_subpattern_reference
	GRegexErrorInexistentSubpatternReference Regexerror = 115
	// unterminated_comment
	GRegexErrorUnterminatedComment Regexerror = 118
	// expression_too_large
	GRegexErrorExpressionTooLarge Regexerror = 120
	// memory_error
	GRegexErrorMemoryError Regexerror = 121
	// variable_length_lookbehind
	GRegexErrorVariableLengthLookbehind Regexerror = 125
	// malformed_condition
	GRegexErrorMalformedCondition Regexerror = 126
	// too_many_conditional_branches
	GRegexErrorTooManyConditionalBranches Regexerror = 127
	// assertion_expected
	GRegexErrorAssertionExpected Regexerror = 128
	// unknown_posix_class_name
	GRegexErrorUnknownPosixClassName Regexerror = 130
	// posix_collating_elements_not_supported
	GRegexErrorPosixCollatingElementsNotSupported Regexerror = 131
	// hex_code_too_large
	GRegexErrorHexCodeTooLarge Regexerror = 134
	// invalid_condition
	GRegexErrorInvalidCondition Regexerror = 135
	// single_byte_match_in_lookbehind
	GRegexErrorSingleByteMatchInLookbehind Regexerror = 136
	// infinite_loop
	GRegexErrorInfiniteLoop Regexerror = 140
	// missing_subpattern_name_terminator
	GRegexErrorMissingSubpatternNameTerminator Regexerror = 142
	// duplicate_subpattern_name
	GRegexErrorDuplicateSubpatternName Regexerror = 143
	// malformed_property
	GRegexErrorMalformedProperty Regexerror = 146
	// unknown_property
	GRegexErrorUnknownProperty Regexerror = 147
	// subpattern_name_too_long
	GRegexErrorSubpatternNameTooLong Regexerror = 148
	// too_many_subpatterns
	GRegexErrorTooManySubpatterns Regexerror = 149
	// invalid_octal_value
	GRegexErrorInvalidOctalValue Regexerror = 151
	// too_many_branches_in_define
	GRegexErrorTooManyBranchesInDefine Regexerror = 154
	// define_repetion
	GRegexErrorDefineRepetion Regexerror = 155
	// inconsistent_newline_options
	GRegexErrorInconsistentNewlineOptions Regexerror = 156
	// missing_back_reference
	GRegexErrorMissingBackReference Regexerror = 157
	// invalid_relative_reference
	GRegexErrorInvalidRelativeReference Regexerror = 158
	// backtracking_control_verb_argument_forbidden
	GRegexErrorBacktrackingControlVerbArgumentForbidden Regexerror = 159
	// unknown_backtracking_control_verb
	GRegexErrorUnknownBacktrackingControlVerb Regexerror = 160
	// number_too_big
	GRegexErrorNumberTooBig Regexerror = 161
	// missing_subpattern_name
	GRegexErrorMissingSubpatternName Regexerror = 162
	// missing_digit
	GRegexErrorMissingDigit Regexerror = 163
	// invalid_data_character
	GRegexErrorInvalidDataCharacter Regexerror = 164
	// extra_subpattern_name
	GRegexErrorExtraSubpatternName Regexerror = 165
	// backtracking_control_verb_argument_required
	GRegexErrorBacktrackingControlVerbArgumentRequired Regexerror = 166
	// invalid_control_char
	GRegexErrorInvalidControlChar Regexerror = 168
	// missing_name
	GRegexErrorMissingName Regexerror = 169
	// not_supported_in_class
	GRegexErrorNotSupportedInClass Regexerror = 171
	// too_many_forward_references
	GRegexErrorTooManyForwardReferences Regexerror = 172
	// name_too_long
	GRegexErrorNameTooLong Regexerror = 175
	// character_value_too_large
	GRegexErrorCharacterValueTooLarge Regexerror = 176
)

// Seektype is a representation of the C type SeekType.
type Seektype int

const (
	// cur
	GSeekCur Seektype = 0
	// set
	GSeekSet Seektype = 1
	// end
	GSeekEnd Seektype = 2
)

// Shellerror is a representation of the C type ShellError.
type Shellerror int

const (
	// bad_quoting
	GShellErrorBadQuoting Shellerror = 0
	// empty_string
	GShellErrorEmptyString Shellerror = 1
	// failed
	GShellErrorFailed Shellerror = 2
)

// Sliceconfig is a representation of the C type SliceConfig.
type Sliceconfig int

const (
	// always_malloc
	GSliceConfigAlwaysMalloc Sliceconfig = 1
	// bypass_magazines
	GSliceConfigBypassMagazines Sliceconfig = 2
	// working_set_msecs
	GSliceConfigWorkingSetMsecs Sliceconfig = 3
	// color_increment
	GSliceConfigColorIncrement Sliceconfig = 4
	// chunk_sizes
	GSliceConfigChunkSizes Sliceconfig = 5
	// contention_counter
	GSliceConfigContentionCounter Sliceconfig = 6
)

// Spawnerror is a representation of the C type SpawnError.
type Spawnerror int

const (
	// fork
	GSpawnErrorFork Spawnerror = 0
	// read
	GSpawnErrorRead Spawnerror = 1
	// chdir
	GSpawnErrorChdir Spawnerror = 2
	// acces
	GSpawnErrorAcces Spawnerror = 3
	// perm
	GSpawnErrorPerm Spawnerror = 4
	// too_big
	GSpawnErrorTooBig Spawnerror = 5
	// 2big
	GSpawnError2big Spawnerror = 5
	// noexec
	GSpawnErrorNoexec Spawnerror = 6
	// nametoolong
	GSpawnErrorNametoolong Spawnerror = 7
	// noent
	GSpawnErrorNoent Spawnerror = 8
	// nomem
	GSpawnErrorNomem Spawnerror = 9
	// notdir
	GSpawnErrorNotdir Spawnerror = 10
	// loop
	GSpawnErrorLoop Spawnerror = 11
	// txtbusy
	GSpawnErrorTxtbusy Spawnerror = 12
	// io
	GSpawnErrorIo Spawnerror = 13
	// nfile
	GSpawnErrorNfile Spawnerror = 14
	// mfile
	GSpawnErrorMfile Spawnerror = 15
	// inval
	GSpawnErrorInval Spawnerror = 16
	// isdir
	GSpawnErrorIsdir Spawnerror = 17
	// libbad
	GSpawnErrorLibbad Spawnerror = 18
	// failed
	GSpawnErrorFailed Spawnerror = 19
)

// Testfiletype is a representation of the C type TestFileType.
//
// since 2.38
type Testfiletype int

const (
	// dist
	GTestDist Testfiletype = 0
	// built
	GTestBuilt Testfiletype = 1
)

// Testlogtype is a representation of the C type TestLogType.
type Testlogtype int

const (
	// none
	GTestLogNone Testlogtype = 0
	// error
	GTestLogError Testlogtype = 1
	// start_binary
	GTestLogStartBinary Testlogtype = 2
	// list_case
	GTestLogListCase Testlogtype = 3
	// skip_case
	GTestLogSkipCase Testlogtype = 4
	// start_case
	GTestLogStartCase Testlogtype = 5
	// stop_case
	GTestLogStopCase Testlogtype = 6
	// min_result
	GTestLogMinResult Testlogtype = 7
	// max_result
	GTestLogMaxResult Testlogtype = 8
	// message
	GTestLogMessage Testlogtype = 9
	// start_suite
	GTestLogStartSuite Testlogtype = 10
	// stop_suite
	GTestLogStopSuite Testlogtype = 11
)

// Testresult is a representation of the C type TestResult.
type Testresult int

const (
	// success
	GTestRunSuccess Testresult = 0
	// skipped
	GTestRunSkipped Testresult = 1
	// failure
	GTestRunFailure Testresult = 2
	// incomplete
	GTestRunIncomplete Testresult = 3
)

// Threaderror is a representation of the C type ThreadError.
type Threaderror int

const (
	// thread_error_again
	GThreadErrorAgain Threaderror = 0
)

// Timetype is a representation of the C type TimeType.
type Timetype int

const (
	// standard
	GTimeTypeStandard Timetype = 0
	// daylight
	GTimeTypeDaylight Timetype = 1
	// universal
	GTimeTypeUniversal Timetype = 2
)

// Tokentype is a representation of the C type TokenType.
type Tokentype int

const (
	// eof
	GTokenEof Tokentype = 0
	// left_paren
	GTokenLeftParen Tokentype = 40
	// right_paren
	GTokenRightParen Tokentype = 41
	// left_curly
	GTokenLeftCurly Tokentype = 123
	// right_curly
	GTokenRightCurly Tokentype = 125
	// left_brace
	GTokenLeftBrace Tokentype = 91
	// right_brace
	GTokenRightBrace Tokentype = 93
	// equal_sign
	GTokenEqualSign Tokentype = 61
	// comma
	GTokenComma Tokentype = 44
	// none
	GTokenNone Tokentype = 256
	// error
	GTokenError Tokentype = 257
	// char
	GTokenChar Tokentype = 258
	// binary
	GTokenBinary Tokentype = 259
	// octal
	GTokenOctal Tokentype = 260
	// int
	GTokenInt Tokentype = 261
	// hex
	GTokenHex Tokentype = 262
	// float
	GTokenFloat Tokentype = 263
	// string
	GTokenString Tokentype = 264
	// symbol
	GTokenSymbol Tokentype = 265
	// identifier
	GTokenIdentifier Tokentype = 266
	// identifier_null
	GTokenIdentifierNull Tokentype = 267
	// comment_single
	GTokenCommentSingle Tokentype = 268
	// comment_multi
	GTokenCommentMulti Tokentype = 269
)

// Traversetype is a representation of the C type TraverseType.
type Traversetype int

const (
	// in_order
	GInOrder Traversetype = 0
	// pre_order
	GPreOrder Traversetype = 1
	// post_order
	GPostOrder Traversetype = 2
	// level_order
	GLevelOrder Traversetype = 3
)

// Unicodebreaktype is a representation of the C type UnicodeBreakType.
type Unicodebreaktype int

const (
	// mandatory
	GUnicodeBreakMandatory Unicodebreaktype = 0
	// carriage_return
	GUnicodeBreakCarriageReturn Unicodebreaktype = 1
	// line_feed
	GUnicodeBreakLineFeed Unicodebreaktype = 2
	// combining_mark
	GUnicodeBreakCombiningMark Unicodebreaktype = 3
	// surrogate
	GUnicodeBreakSurrogate Unicodebreaktype = 4
	// zero_width_space
	GUnicodeBreakZeroWidthSpace Unicodebreaktype = 5
	// inseparable
	GUnicodeBreakInseparable Unicodebreaktype = 6
	// non_breaking_glue
	GUnicodeBreakNonBreakingGlue Unicodebreaktype = 7
	// contingent
	GUnicodeBreakContingent Unicodebreaktype = 8
	// space
	GUnicodeBreakSpace Unicodebreaktype = 9
	// after
	GUnicodeBreakAfter Unicodebreaktype = 10
	// before
	GUnicodeBreakBefore Unicodebreaktype = 11
	// before_and_after
	GUnicodeBreakBeforeAndAfter Unicodebreaktype = 12
	// hyphen
	GUnicodeBreakHyphen Unicodebreaktype = 13
	// non_starter
	GUnicodeBreakNonStarter Unicodebreaktype = 14
	// open_punctuation
	GUnicodeBreakOpenPunctuation Unicodebreaktype = 15
	// close_punctuation
	GUnicodeBreakClosePunctuation Unicodebreaktype = 16
	// quotation
	GUnicodeBreakQuotation Unicodebreaktype = 17
	// exclamation
	GUnicodeBreakExclamation Unicodebreaktype = 18
	// ideographic
	GUnicodeBreakIdeographic Unicodebreaktype = 19
	// numeric
	GUnicodeBreakNumeric Unicodebreaktype = 20
	// infix_separator
	GUnicodeBreakInfixSeparator Unicodebreaktype = 21
	// symbol
	GUnicodeBreakSymbol Unicodebreaktype = 22
	// alphabetic
	GUnicodeBreakAlphabetic Unicodebreaktype = 23
	// prefix
	GUnicodeBreakPrefix Unicodebreaktype = 24
	// postfix
	GUnicodeBreakPostfix Unicodebreaktype = 25
	// complex_context
	GUnicodeBreakComplexContext Unicodebreaktype = 26
	// ambiguous
	GUnicodeBreakAmbiguous Unicodebreaktype = 27
	// unknown
	GUnicodeBreakUnknown Unicodebreaktype = 28
	// next_line
	GUnicodeBreakNextLine Unicodebreaktype = 29
	// word_joiner
	GUnicodeBreakWordJoiner Unicodebreaktype = 30
	// hangul_l_jamo
	GUnicodeBreakHangulLJamo Unicodebreaktype = 31
	// hangul_v_jamo
	GUnicodeBreakHangulVJamo Unicodebreaktype = 32
	// hangul_t_jamo
	GUnicodeBreakHangulTJamo Unicodebreaktype = 33
	// hangul_lv_syllable
	GUnicodeBreakHangulLvSyllable Unicodebreaktype = 34
	// hangul_lvt_syllable
	GUnicodeBreakHangulLvtSyllable Unicodebreaktype = 35
	// close_paranthesis
	GUnicodeBreakCloseParanthesis Unicodebreaktype = 36
	// conditional_japanese_starter
	GUnicodeBreakConditionalJapaneseStarter Unicodebreaktype = 37
	// hebrew_letter
	GUnicodeBreakHebrewLetter Unicodebreaktype = 38
	// regional_indicator
	GUnicodeBreakRegionalIndicator Unicodebreaktype = 39
	// emoji_base
	GUnicodeBreakEmojiBase Unicodebreaktype = 40
	// emoji_modifier
	GUnicodeBreakEmojiModifier Unicodebreaktype = 41
	// zero_width_joiner
	GUnicodeBreakZeroWidthJoiner Unicodebreaktype = 42
)

// Unicodescript is a representation of the C type UnicodeScript.
type Unicodescript int

const (
	// invalid_code
	GUnicodeScriptInvalidCode Unicodescript = -1
	// common
	GUnicodeScriptCommon Unicodescript = 0
	// inherited
	GUnicodeScriptInherited Unicodescript = 1
	// arabic
	GUnicodeScriptArabic Unicodescript = 2
	// armenian
	GUnicodeScriptArmenian Unicodescript = 3
	// bengali
	GUnicodeScriptBengali Unicodescript = 4
	// bopomofo
	GUnicodeScriptBopomofo Unicodescript = 5
	// cherokee
	GUnicodeScriptCherokee Unicodescript = 6
	// coptic
	GUnicodeScriptCoptic Unicodescript = 7
	// cyrillic
	GUnicodeScriptCyrillic Unicodescript = 8
	// deseret
	GUnicodeScriptDeseret Unicodescript = 9
	// devanagari
	GUnicodeScriptDevanagari Unicodescript = 10
	// ethiopic
	GUnicodeScriptEthiopic Unicodescript = 11
	// georgian
	GUnicodeScriptGeorgian Unicodescript = 12
	// gothic
	GUnicodeScriptGothic Unicodescript = 13
	// greek
	GUnicodeScriptGreek Unicodescript = 14
	// gujarati
	GUnicodeScriptGujarati Unicodescript = 15
	// gurmukhi
	GUnicodeScriptGurmukhi Unicodescript = 16
	// han
	GUnicodeScriptHan Unicodescript = 17
	// hangul
	GUnicodeScriptHangul Unicodescript = 18
	// hebrew
	GUnicodeScriptHebrew Unicodescript = 19
	// hiragana
	GUnicodeScriptHiragana Unicodescript = 20
	// kannada
	GUnicodeScriptKannada Unicodescript = 21
	// katakana
	GUnicodeScriptKatakana Unicodescript = 22
	// khmer
	GUnicodeScriptKhmer Unicodescript = 23
	// lao
	GUnicodeScriptLao Unicodescript = 24
	// latin
	GUnicodeScriptLatin Unicodescript = 25
	// malayalam
	GUnicodeScriptMalayalam Unicodescript = 26
	// mongolian
	GUnicodeScriptMongolian Unicodescript = 27
	// myanmar
	GUnicodeScriptMyanmar Unicodescript = 28
	// ogham
	GUnicodeScriptOgham Unicodescript = 29
	// old_italic
	GUnicodeScriptOldItalic Unicodescript = 30
	// oriya
	GUnicodeScriptOriya Unicodescript = 31
	// runic
	GUnicodeScriptRunic Unicodescript = 32
	// sinhala
	GUnicodeScriptSinhala Unicodescript = 33
	// syriac
	GUnicodeScriptSyriac Unicodescript = 34
	// tamil
	GUnicodeScriptTamil Unicodescript = 35
	// telugu
	GUnicodeScriptTelugu Unicodescript = 36
	// thaana
	GUnicodeScriptThaana Unicodescript = 37
	// thai
	GUnicodeScriptThai Unicodescript = 38
	// tibetan
	GUnicodeScriptTibetan Unicodescript = 39
	// canadian_aboriginal
	GUnicodeScriptCanadianAboriginal Unicodescript = 40
	// yi
	GUnicodeScriptYi Unicodescript = 41
	// tagalog
	GUnicodeScriptTagalog Unicodescript = 42
	// hanunoo
	GUnicodeScriptHanunoo Unicodescript = 43
	// buhid
	GUnicodeScriptBuhid Unicodescript = 44
	// tagbanwa
	GUnicodeScriptTagbanwa Unicodescript = 45
	// braille
	GUnicodeScriptBraille Unicodescript = 46
	// cypriot
	GUnicodeScriptCypriot Unicodescript = 47
	// limbu
	GUnicodeScriptLimbu Unicodescript = 48
	// osmanya
	GUnicodeScriptOsmanya Unicodescript = 49
	// shavian
	GUnicodeScriptShavian Unicodescript = 50
	// linear_b
	GUnicodeScriptLinearB Unicodescript = 51
	// tai_le
	GUnicodeScriptTaiLe Unicodescript = 52
	// ugaritic
	GUnicodeScriptUgaritic Unicodescript = 53
	// new_tai_lue
	GUnicodeScriptNewTaiLue Unicodescript = 54
	// buginese
	GUnicodeScriptBuginese Unicodescript = 55
	// glagolitic
	GUnicodeScriptGlagolitic Unicodescript = 56
	// tifinagh
	GUnicodeScriptTifinagh Unicodescript = 57
	// syloti_nagri
	GUnicodeScriptSylotiNagri Unicodescript = 58
	// old_persian
	GUnicodeScriptOldPersian Unicodescript = 59
	// kharoshthi
	GUnicodeScriptKharoshthi Unicodescript = 60
	// unknown
	GUnicodeScriptUnknown Unicodescript = 61
	// balinese
	GUnicodeScriptBalinese Unicodescript = 62
	// cuneiform
	GUnicodeScriptCuneiform Unicodescript = 63
	// phoenician
	GUnicodeScriptPhoenician Unicodescript = 64
	// phags_pa
	GUnicodeScriptPhagsPa Unicodescript = 65
	// nko
	GUnicodeScriptNko Unicodescript = 66
	// kayah_li
	GUnicodeScriptKayahLi Unicodescript = 67
	// lepcha
	GUnicodeScriptLepcha Unicodescript = 68
	// rejang
	GUnicodeScriptRejang Unicodescript = 69
	// sundanese
	GUnicodeScriptSundanese Unicodescript = 70
	// saurashtra
	GUnicodeScriptSaurashtra Unicodescript = 71
	// cham
	GUnicodeScriptCham Unicodescript = 72
	// ol_chiki
	GUnicodeScriptOlChiki Unicodescript = 73
	// vai
	GUnicodeScriptVai Unicodescript = 74
	// carian
	GUnicodeScriptCarian Unicodescript = 75
	// lycian
	GUnicodeScriptLycian Unicodescript = 76
	// lydian
	GUnicodeScriptLydian Unicodescript = 77
	// avestan
	GUnicodeScriptAvestan Unicodescript = 78
	// bamum
	GUnicodeScriptBamum Unicodescript = 79
	// egyptian_hieroglyphs
	GUnicodeScriptEgyptianHieroglyphs Unicodescript = 80
	// imperial_aramaic
	GUnicodeScriptImperialAramaic Unicodescript = 81
	// inscriptional_pahlavi
	GUnicodeScriptInscriptionalPahlavi Unicodescript = 82
	// inscriptional_parthian
	GUnicodeScriptInscriptionalParthian Unicodescript = 83
	// javanese
	GUnicodeScriptJavanese Unicodescript = 84
	// kaithi
	GUnicodeScriptKaithi Unicodescript = 85
	// lisu
	GUnicodeScriptLisu Unicodescript = 86
	// meetei_mayek
	GUnicodeScriptMeeteiMayek Unicodescript = 87
	// old_south_arabian
	GUnicodeScriptOldSouthArabian Unicodescript = 88
	// old_turkic
	GUnicodeScriptOldTurkic Unicodescript = 89
	// samaritan
	GUnicodeScriptSamaritan Unicodescript = 90
	// tai_tham
	GUnicodeScriptTaiTham Unicodescript = 91
	// tai_viet
	GUnicodeScriptTaiViet Unicodescript = 92
	// batak
	GUnicodeScriptBatak Unicodescript = 93
	// brahmi
	GUnicodeScriptBrahmi Unicodescript = 94
	// mandaic
	GUnicodeScriptMandaic Unicodescript = 95
	// chakma
	GUnicodeScriptChakma Unicodescript = 96
	// meroitic_cursive
	GUnicodeScriptMeroiticCursive Unicodescript = 97
	// meroitic_hieroglyphs
	GUnicodeScriptMeroiticHieroglyphs Unicodescript = 98
	// miao
	GUnicodeScriptMiao Unicodescript = 99
	// sharada
	GUnicodeScriptSharada Unicodescript = 100
	// sora_sompeng
	GUnicodeScriptSoraSompeng Unicodescript = 101
	// takri
	GUnicodeScriptTakri Unicodescript = 102
	// bassa_vah
	GUnicodeScriptBassaVah Unicodescript = 103
	// caucasian_albanian
	GUnicodeScriptCaucasianAlbanian Unicodescript = 104
	// duployan
	GUnicodeScriptDuployan Unicodescript = 105
	// elbasan
	GUnicodeScriptElbasan Unicodescript = 106
	// grantha
	GUnicodeScriptGrantha Unicodescript = 107
	// khojki
	GUnicodeScriptKhojki Unicodescript = 108
	// khudawadi
	GUnicodeScriptKhudawadi Unicodescript = 109
	// linear_a
	GUnicodeScriptLinearA Unicodescript = 110
	// mahajani
	GUnicodeScriptMahajani Unicodescript = 111
	// manichaean
	GUnicodeScriptManichaean Unicodescript = 112
	// mende_kikakui
	GUnicodeScriptMendeKikakui Unicodescript = 113
	// modi
	GUnicodeScriptModi Unicodescript = 114
	// mro
	GUnicodeScriptMro Unicodescript = 115
	// nabataean
	GUnicodeScriptNabataean Unicodescript = 116
	// old_north_arabian
	GUnicodeScriptOldNorthArabian Unicodescript = 117
	// old_permic
	GUnicodeScriptOldPermic Unicodescript = 118
	// pahawh_hmong
	GUnicodeScriptPahawhHmong Unicodescript = 119
	// palmyrene
	GUnicodeScriptPalmyrene Unicodescript = 120
	// pau_cin_hau
	GUnicodeScriptPauCinHau Unicodescript = 121
	// psalter_pahlavi
	GUnicodeScriptPsalterPahlavi Unicodescript = 122
	// siddham
	GUnicodeScriptSiddham Unicodescript = 123
	// tirhuta
	GUnicodeScriptTirhuta Unicodescript = 124
	// warang_citi
	GUnicodeScriptWarangCiti Unicodescript = 125
	// ahom
	GUnicodeScriptAhom Unicodescript = 126
	// anatolian_hieroglyphs
	GUnicodeScriptAnatolianHieroglyphs Unicodescript = 127
	// hatran
	GUnicodeScriptHatran Unicodescript = 128
	// multani
	GUnicodeScriptMultani Unicodescript = 129
	// old_hungarian
	GUnicodeScriptOldHungarian Unicodescript = 130
	// signwriting
	GUnicodeScriptSignwriting Unicodescript = 131
	// adlam
	GUnicodeScriptAdlam Unicodescript = 132
	// bhaiksuki
	GUnicodeScriptBhaiksuki Unicodescript = 133
	// marchen
	GUnicodeScriptMarchen Unicodescript = 134
	// newa
	GUnicodeScriptNewa Unicodescript = 135
	// osage
	GUnicodeScriptOsage Unicodescript = 136
	// tangut
	GUnicodeScriptTangut Unicodescript = 137
	// masaram_gondi
	GUnicodeScriptMasaramGondi Unicodescript = 138
	// nushu
	GUnicodeScriptNushu Unicodescript = 139
	// soyombo
	GUnicodeScriptSoyombo Unicodescript = 140
	// zanabazar_square
	GUnicodeScriptZanabazarSquare Unicodescript = 141
)

// Unicodetype is a representation of the C type UnicodeType.
type Unicodetype int

const (
	// control
	GUnicodeControl Unicodetype = 0
	// format
	GUnicodeFormat Unicodetype = 1
	// unassigned
	GUnicodeUnassigned Unicodetype = 2
	// private_use
	GUnicodePrivateUse Unicodetype = 3
	// surrogate
	GUnicodeSurrogate Unicodetype = 4
	// lowercase_letter
	GUnicodeLowercaseLetter Unicodetype = 5
	// modifier_letter
	GUnicodeModifierLetter Unicodetype = 6
	// other_letter
	GUnicodeOtherLetter Unicodetype = 7
	// titlecase_letter
	GUnicodeTitlecaseLetter Unicodetype = 8
	// uppercase_letter
	GUnicodeUppercaseLetter Unicodetype = 9
	// spacing_mark
	GUnicodeSpacingMark Unicodetype = 10
	// enclosing_mark
	GUnicodeEnclosingMark Unicodetype = 11
	// non_spacing_mark
	GUnicodeNonSpacingMark Unicodetype = 12
	// decimal_number
	GUnicodeDecimalNumber Unicodetype = 13
	// letter_number
	GUnicodeLetterNumber Unicodetype = 14
	// other_number
	GUnicodeOtherNumber Unicodetype = 15
	// connect_punctuation
	GUnicodeConnectPunctuation Unicodetype = 16
	// dash_punctuation
	GUnicodeDashPunctuation Unicodetype = 17
	// close_punctuation
	GUnicodeClosePunctuation Unicodetype = 18
	// final_punctuation
	GUnicodeFinalPunctuation Unicodetype = 19
	// initial_punctuation
	GUnicodeInitialPunctuation Unicodetype = 20
	// other_punctuation
	GUnicodeOtherPunctuation Unicodetype = 21
	// open_punctuation
	GUnicodeOpenPunctuation Unicodetype = 22
	// currency_symbol
	GUnicodeCurrencySymbol Unicodetype = 23
	// modifier_symbol
	GUnicodeModifierSymbol Unicodetype = 24
	// math_symbol
	GUnicodeMathSymbol Unicodetype = 25
	// other_symbol
	GUnicodeOtherSymbol Unicodetype = 26
	// line_separator
	GUnicodeLineSeparator Unicodetype = 27
	// paragraph_separator
	GUnicodeParagraphSeparator Unicodetype = 28
	// space_separator
	GUnicodeSpaceSeparator Unicodetype = 29
)

// Userdirectory is a representation of the C type UserDirectory.
//
// since 2.14
type Userdirectory int

const (
	// directory_desktop
	GUserDirectoryDesktop Userdirectory = 0
	// directory_documents
	GUserDirectoryDocuments Userdirectory = 1
	// directory_download
	GUserDirectoryDownload Userdirectory = 2
	// directory_music
	GUserDirectoryMusic Userdirectory = 3
	// directory_pictures
	GUserDirectoryPictures Userdirectory = 4
	// directory_public_share
	GUserDirectoryPublicShare Userdirectory = 5
	// directory_templates
	GUserDirectoryTemplates Userdirectory = 6
	// directory_videos
	GUserDirectoryVideos Userdirectory = 7
	// n_directories
	GUserNDirectories Userdirectory = 8
)

// Variantclass is a representation of the C type VariantClass.
//
// since 2.24
type Variantclass int

const (
	// boolean
	GVariantClassBoolean Variantclass = 98
	// byte
	GVariantClassByte Variantclass = 121
	// int16
	GVariantClassInt16 Variantclass = 110
	// uint16
	GVariantClassUint16 Variantclass = 113
	// int32
	GVariantClassInt32 Variantclass = 105
	// uint32
	GVariantClassUint32 Variantclass = 117
	// int64
	GVariantClassInt64 Variantclass = 120
	// uint64
	GVariantClassUint64 Variantclass = 116
	// handle
	GVariantClassHandle Variantclass = 104
	// double
	GVariantClassDouble Variantclass = 100
	// string
	GVariantClassString Variantclass = 115
	// object_path
	GVariantClassObjectPath Variantclass = 111
	// signature
	GVariantClassSignature Variantclass = 103
	// variant
	GVariantClassVariant Variantclass = 118
	// maybe
	GVariantClassMaybe Variantclass = 109
	// array
	GVariantClassArray Variantclass = 97
	// tuple
	GVariantClassTuple Variantclass = 40
	// dict_entry
	GVariantClassDictEntry Variantclass = 123
)

// Variantparseerror is a representation of the C type VariantParseError.
type Variantparseerror int

const (
	// failed
	GVariantParseErrorFailed Variantparseerror = 0
	// basic_type_expected
	GVariantParseErrorBasicTypeExpected Variantparseerror = 1
	// cannot_infer_type
	GVariantParseErrorCannotInferType Variantparseerror = 2
	// definite_type_expected
	GVariantParseErrorDefiniteTypeExpected Variantparseerror = 3
	// input_not_at_end
	GVariantParseErrorInputNotAtEnd Variantparseerror = 4
	// invalid_character
	GVariantParseErrorInvalidCharacter Variantparseerror = 5
	// invalid_format_string
	GVariantParseErrorInvalidFormatString Variantparseerror = 6
	// invalid_object_path
	GVariantParseErrorInvalidObjectPath Variantparseerror = 7
	// invalid_signature
	GVariantParseErrorInvalidSignature Variantparseerror = 8
	// invalid_type_string
	GVariantParseErrorInvalidTypeString Variantparseerror = 9
	// no_common_type
	GVariantParseErrorNoCommonType Variantparseerror = 10
	// number_out_of_range
	GVariantParseErrorNumberOutOfRange Variantparseerror = 11
	// number_too_big
	GVariantParseErrorNumberTooBig Variantparseerror = 12
	// type_error
	GVariantParseErrorTypeError Variantparseerror = 13
	// unexpected_token
	GVariantParseErrorUnexpectedToken Variantparseerror = 14
	// unknown_keyword
	GVariantParseErrorUnknownKeyword Variantparseerror = 15
	// unterminated_string_constant
	GVariantParseErrorUnterminatedStringConstant Variantparseerror = 16
	// value_expected
	GVariantParseErrorValueExpected Variantparseerror = 17
)

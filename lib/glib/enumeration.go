// Code generated - DO NOT EDIT.

package glib

// Bookmarkfileerror is a representation of the C type BookmarkFileError.
type Bookmarkfileerror int

const (
	Bookmarkfileerror_InvalidUri       Bookmarkfileerror = 0
	Bookmarkfileerror_InvalidValue     Bookmarkfileerror = 1
	Bookmarkfileerror_AppNotRegistered Bookmarkfileerror = 2
	Bookmarkfileerror_UriNotFound      Bookmarkfileerror = 3
	Bookmarkfileerror_Read             Bookmarkfileerror = 4
	Bookmarkfileerror_UnknownEncoding  Bookmarkfileerror = 5
	Bookmarkfileerror_Write            Bookmarkfileerror = 6
	Bookmarkfileerror_FileNotFound     Bookmarkfileerror = 7
)

// Checksumtype is a representation of the C type ChecksumType.
//
// since 2.16
type Checksumtype int

const (
	Checksumtype_Md5    Checksumtype = 0
	Checksumtype_Sha1   Checksumtype = 1
	Checksumtype_Sha256 Checksumtype = 2
	Checksumtype_Sha512 Checksumtype = 3
	Checksumtype_Sha384 Checksumtype = 4
)

// Converterror is a representation of the C type ConvertError.
type Converterror int

const (
	Converterror_NoConversion    Converterror = 0
	Converterror_IllegalSequence Converterror = 1
	Converterror_Failed          Converterror = 2
	Converterror_PartialInput    Converterror = 3
	Converterror_BadUri          Converterror = 4
	Converterror_NotAbsolutePath Converterror = 5
	Converterror_NoMemory        Converterror = 6
	Converterror_EmbeddedNul     Converterror = 7
)

// Datedmy is a representation of the C type DateDMY.
type Datedmy int

const (
	Datedmy_Day   Datedmy = 0
	Datedmy_Month Datedmy = 1
	Datedmy_Year  Datedmy = 2
)

// Datemonth is a representation of the C type DateMonth.
type Datemonth int

const (
	Datemonth_BadMonth  Datemonth = 0
	Datemonth_January   Datemonth = 1
	Datemonth_February  Datemonth = 2
	Datemonth_March     Datemonth = 3
	Datemonth_April     Datemonth = 4
	Datemonth_May       Datemonth = 5
	Datemonth_June      Datemonth = 6
	Datemonth_July      Datemonth = 7
	Datemonth_August    Datemonth = 8
	Datemonth_September Datemonth = 9
	Datemonth_October   Datemonth = 10
	Datemonth_November  Datemonth = 11
	Datemonth_December  Datemonth = 12
)

// Dateweekday is a representation of the C type DateWeekday.
type Dateweekday int

const (
	Dateweekday_BadWeekday Dateweekday = 0
	Dateweekday_Monday     Dateweekday = 1
	Dateweekday_Tuesday    Dateweekday = 2
	Dateweekday_Wednesday  Dateweekday = 3
	Dateweekday_Thursday   Dateweekday = 4
	Dateweekday_Friday     Dateweekday = 5
	Dateweekday_Saturday   Dateweekday = 6
	Dateweekday_Sunday     Dateweekday = 7
)

// Errortype is a representation of the C type ErrorType.
type Errortype int

const (
	Errortype_Unknown           Errortype = 0
	Errortype_UnexpEof          Errortype = 1
	Errortype_UnexpEofInString  Errortype = 2
	Errortype_UnexpEofInComment Errortype = 3
	Errortype_NonDigitInConst   Errortype = 4
	Errortype_DigitRadix        Errortype = 5
	Errortype_FloatRadix        Errortype = 6
	Errortype_FloatMalformed    Errortype = 7
)

// Fileerror is a representation of the C type FileError.
type Fileerror int

const (
	Fileerror_Exist       Fileerror = 0
	Fileerror_Isdir       Fileerror = 1
	Fileerror_Acces       Fileerror = 2
	Fileerror_Nametoolong Fileerror = 3
	Fileerror_Noent       Fileerror = 4
	Fileerror_Notdir      Fileerror = 5
	Fileerror_Nxio        Fileerror = 6
	Fileerror_Nodev       Fileerror = 7
	Fileerror_Rofs        Fileerror = 8
	Fileerror_Txtbsy      Fileerror = 9
	Fileerror_Fault       Fileerror = 10
	Fileerror_Loop        Fileerror = 11
	Fileerror_Nospc       Fileerror = 12
	Fileerror_Nomem       Fileerror = 13
	Fileerror_Mfile       Fileerror = 14
	Fileerror_Nfile       Fileerror = 15
	Fileerror_Badf        Fileerror = 16
	Fileerror_Inval       Fileerror = 17
	Fileerror_Pipe        Fileerror = 18
	Fileerror_Again       Fileerror = 19
	Fileerror_Intr        Fileerror = 20
	Fileerror_Io          Fileerror = 21
	Fileerror_Perm        Fileerror = 22
	Fileerror_Nosys       Fileerror = 23
	Fileerror_Failed      Fileerror = 24
)

// Iochannelerror is a representation of the C type IOChannelError.
type Iochannelerror int

const (
	Iochannelerror_Fbig     Iochannelerror = 0
	Iochannelerror_Inval    Iochannelerror = 1
	Iochannelerror_Io       Iochannelerror = 2
	Iochannelerror_Isdir    Iochannelerror = 3
	Iochannelerror_Nospc    Iochannelerror = 4
	Iochannelerror_Nxio     Iochannelerror = 5
	Iochannelerror_Overflow Iochannelerror = 6
	Iochannelerror_Pipe     Iochannelerror = 7
	Iochannelerror_Failed   Iochannelerror = 8
)

// Ioerror is a representation of the C type IOError.
type Ioerror int

const (
	Ioerror_None    Ioerror = 0
	Ioerror_Again   Ioerror = 1
	Ioerror_Inval   Ioerror = 2
	Ioerror_Unknown Ioerror = 3
)

// Iostatus is a representation of the C type IOStatus.
type Iostatus int

const (
	Iostatus_Error  Iostatus = 0
	Iostatus_Normal Iostatus = 1
	Iostatus_Eof    Iostatus = 2
	Iostatus_Again  Iostatus = 3
)

// Keyfileerror is a representation of the C type KeyFileError.
type Keyfileerror int

const (
	Keyfileerror_UnknownEncoding Keyfileerror = 0
	Keyfileerror_Parse           Keyfileerror = 1
	Keyfileerror_NotFound        Keyfileerror = 2
	Keyfileerror_KeyNotFound     Keyfileerror = 3
	Keyfileerror_GroupNotFound   Keyfileerror = 4
	Keyfileerror_InvalidValue    Keyfileerror = 5
)

// Logwriteroutput is a representation of the C type LogWriterOutput.
//
// since 2.50
type Logwriteroutput int

const (
	Logwriteroutput_Handled   Logwriteroutput = 1
	Logwriteroutput_Unhandled Logwriteroutput = 0
)

// Markuperror is a representation of the C type MarkupError.
type Markuperror int

const (
	Markuperror_BadUtf8          Markuperror = 0
	Markuperror_Empty            Markuperror = 1
	Markuperror_Parse            Markuperror = 2
	Markuperror_UnknownElement   Markuperror = 3
	Markuperror_UnknownAttribute Markuperror = 4
	Markuperror_InvalidContent   Markuperror = 5
	Markuperror_MissingAttribute Markuperror = 6
)

// Normalizemode is a representation of the C type NormalizeMode.
type Normalizemode int

const (
	Normalizemode_Default        Normalizemode = 0
	Normalizemode_Nfd            Normalizemode = 0
	Normalizemode_DefaultCompose Normalizemode = 1
	Normalizemode_Nfc            Normalizemode = 1
	Normalizemode_All            Normalizemode = 2
	Normalizemode_Nfkd           Normalizemode = 2
	Normalizemode_AllCompose     Normalizemode = 3
	Normalizemode_Nfkc           Normalizemode = 3
)

// Numberparsererror is a representation of the C type NumberParserError.
//
// since 2.54
type Numberparsererror int

const (
	Numberparsererror_Invalid     Numberparsererror = 0
	Numberparsererror_OutOfBounds Numberparsererror = 1
)

// Oncestatus is a representation of the C type OnceStatus.
//
// since 2.4
type Oncestatus int

const (
	Oncestatus_Notcalled Oncestatus = 0
	Oncestatus_Progress  Oncestatus = 1
	Oncestatus_Ready     Oncestatus = 2
)

// Optionarg is a representation of the C type OptionArg.
type Optionarg int

const (
	Optionarg_None          Optionarg = 0
	Optionarg_String        Optionarg = 1
	Optionarg_Int           Optionarg = 2
	Optionarg_Callback      Optionarg = 3
	Optionarg_Filename      Optionarg = 4
	Optionarg_StringArray   Optionarg = 5
	Optionarg_FilenameArray Optionarg = 6
	Optionarg_Double        Optionarg = 7
	Optionarg_Int64         Optionarg = 8
)

// Optionerror is a representation of the C type OptionError.
type Optionerror int

const (
	Optionerror_UnknownOption Optionerror = 0
	Optionerror_BadValue      Optionerror = 1
	Optionerror_Failed        Optionerror = 2
)

// Regexerror is a representation of the C type RegexError.
//
// since 2.14
type Regexerror int

const (
	Regexerror_Compile                                  Regexerror = 0
	Regexerror_Optimize                                 Regexerror = 1
	Regexerror_Replace                                  Regexerror = 2
	Regexerror_Match                                    Regexerror = 3
	Regexerror_Internal                                 Regexerror = 4
	Regexerror_StrayBackslash                           Regexerror = 101
	Regexerror_MissingControlChar                       Regexerror = 102
	Regexerror_UnrecognizedEscape                       Regexerror = 103
	Regexerror_QuantifiersOutOfOrder                    Regexerror = 104
	Regexerror_QuantifierTooBig                         Regexerror = 105
	Regexerror_UnterminatedCharacterClass               Regexerror = 106
	Regexerror_InvalidEscapeInCharacterClass            Regexerror = 107
	Regexerror_RangeOutOfOrder                          Regexerror = 108
	Regexerror_NothingToRepeat                          Regexerror = 109
	Regexerror_UnrecognizedCharacter                    Regexerror = 112
	Regexerror_PosixNamedClassOutsideClass              Regexerror = 113
	Regexerror_UnmatchedParenthesis                     Regexerror = 114
	Regexerror_InexistentSubpatternReference            Regexerror = 115
	Regexerror_UnterminatedComment                      Regexerror = 118
	Regexerror_ExpressionTooLarge                       Regexerror = 120
	Regexerror_MemoryError                              Regexerror = 121
	Regexerror_VariableLengthLookbehind                 Regexerror = 125
	Regexerror_MalformedCondition                       Regexerror = 126
	Regexerror_TooManyConditionalBranches               Regexerror = 127
	Regexerror_AssertionExpected                        Regexerror = 128
	Regexerror_UnknownPosixClassName                    Regexerror = 130
	Regexerror_PosixCollatingElementsNotSupported       Regexerror = 131
	Regexerror_HexCodeTooLarge                          Regexerror = 134
	Regexerror_InvalidCondition                         Regexerror = 135
	Regexerror_SingleByteMatchInLookbehind              Regexerror = 136
	Regexerror_InfiniteLoop                             Regexerror = 140
	Regexerror_MissingSubpatternNameTerminator          Regexerror = 142
	Regexerror_DuplicateSubpatternName                  Regexerror = 143
	Regexerror_MalformedProperty                        Regexerror = 146
	Regexerror_UnknownProperty                          Regexerror = 147
	Regexerror_SubpatternNameTooLong                    Regexerror = 148
	Regexerror_TooManySubpatterns                       Regexerror = 149
	Regexerror_InvalidOctalValue                        Regexerror = 151
	Regexerror_TooManyBranchesInDefine                  Regexerror = 154
	Regexerror_DefineRepetion                           Regexerror = 155
	Regexerror_InconsistentNewlineOptions               Regexerror = 156
	Regexerror_MissingBackReference                     Regexerror = 157
	Regexerror_InvalidRelativeReference                 Regexerror = 158
	Regexerror_BacktrackingControlVerbArgumentForbidden Regexerror = 159
	Regexerror_UnknownBacktrackingControlVerb           Regexerror = 160
	Regexerror_NumberTooBig                             Regexerror = 161
	Regexerror_MissingSubpatternName                    Regexerror = 162
	Regexerror_MissingDigit                             Regexerror = 163
	Regexerror_InvalidDataCharacter                     Regexerror = 164
	Regexerror_ExtraSubpatternName                      Regexerror = 165
	Regexerror_BacktrackingControlVerbArgumentRequired  Regexerror = 166
	Regexerror_InvalidControlChar                       Regexerror = 168
	Regexerror_MissingName                              Regexerror = 169
	Regexerror_NotSupportedInClass                      Regexerror = 171
	Regexerror_TooManyForwardReferences                 Regexerror = 172
	Regexerror_NameTooLong                              Regexerror = 175
	Regexerror_CharacterValueTooLarge                   Regexerror = 176
)

// Seektype is a representation of the C type SeekType.
type Seektype int

const (
	Seektype_Cur Seektype = 0
	Seektype_Set Seektype = 1
	Seektype_End Seektype = 2
)

// Shellerror is a representation of the C type ShellError.
type Shellerror int

const (
	Shellerror_BadQuoting  Shellerror = 0
	Shellerror_EmptyString Shellerror = 1
	Shellerror_Failed      Shellerror = 2
)

// Sliceconfig is a representation of the C type SliceConfig.
type Sliceconfig int

const (
	Sliceconfig_AlwaysMalloc      Sliceconfig = 1
	Sliceconfig_BypassMagazines   Sliceconfig = 2
	Sliceconfig_WorkingSetMsecs   Sliceconfig = 3
	Sliceconfig_ColorIncrement    Sliceconfig = 4
	Sliceconfig_ChunkSizes        Sliceconfig = 5
	Sliceconfig_ContentionCounter Sliceconfig = 6
)

// Spawnerror is a representation of the C type SpawnError.
type Spawnerror int

const (
	Spawnerror_Fork        Spawnerror = 0
	Spawnerror_Read        Spawnerror = 1
	Spawnerror_Chdir       Spawnerror = 2
	Spawnerror_Acces       Spawnerror = 3
	Spawnerror_Perm        Spawnerror = 4
	Spawnerror_TooBig      Spawnerror = 5
	Spawnerror_2big        Spawnerror = 5
	Spawnerror_Noexec      Spawnerror = 6
	Spawnerror_Nametoolong Spawnerror = 7
	Spawnerror_Noent       Spawnerror = 8
	Spawnerror_Nomem       Spawnerror = 9
	Spawnerror_Notdir      Spawnerror = 10
	Spawnerror_Loop        Spawnerror = 11
	Spawnerror_Txtbusy     Spawnerror = 12
	Spawnerror_Io          Spawnerror = 13
	Spawnerror_Nfile       Spawnerror = 14
	Spawnerror_Mfile       Spawnerror = 15
	Spawnerror_Inval       Spawnerror = 16
	Spawnerror_Isdir       Spawnerror = 17
	Spawnerror_Libbad      Spawnerror = 18
	Spawnerror_Failed      Spawnerror = 19
)

// Testfiletype is a representation of the C type TestFileType.
//
// since 2.38
type Testfiletype int

const (
	Testfiletype_Dist  Testfiletype = 0
	Testfiletype_Built Testfiletype = 1
)

// Testlogtype is a representation of the C type TestLogType.
type Testlogtype int

const (
	Testlogtype_None        Testlogtype = 0
	Testlogtype_Error       Testlogtype = 1
	Testlogtype_StartBinary Testlogtype = 2
	Testlogtype_ListCase    Testlogtype = 3
	Testlogtype_SkipCase    Testlogtype = 4
	Testlogtype_StartCase   Testlogtype = 5
	Testlogtype_StopCase    Testlogtype = 6
	Testlogtype_MinResult   Testlogtype = 7
	Testlogtype_MaxResult   Testlogtype = 8
	Testlogtype_Message     Testlogtype = 9
	Testlogtype_StartSuite  Testlogtype = 10
	Testlogtype_StopSuite   Testlogtype = 11
)

// Testresult is a representation of the C type TestResult.
type Testresult int

const (
	Testresult_Success    Testresult = 0
	Testresult_Skipped    Testresult = 1
	Testresult_Failure    Testresult = 2
	Testresult_Incomplete Testresult = 3
)

// Threaderror is a representation of the C type ThreadError.
type Threaderror int

const (
	Threaderror_ThreadErrorAgain Threaderror = 0
)

// Timetype is a representation of the C type TimeType.
type Timetype int

const (
	Timetype_Standard  Timetype = 0
	Timetype_Daylight  Timetype = 1
	Timetype_Universal Timetype = 2
)

// Tokentype is a representation of the C type TokenType.
type Tokentype int

const (
	Tokentype_Eof            Tokentype = 0
	Tokentype_LeftParen      Tokentype = 40
	Tokentype_RightParen     Tokentype = 41
	Tokentype_LeftCurly      Tokentype = 123
	Tokentype_RightCurly     Tokentype = 125
	Tokentype_LeftBrace      Tokentype = 91
	Tokentype_RightBrace     Tokentype = 93
	Tokentype_EqualSign      Tokentype = 61
	Tokentype_Comma          Tokentype = 44
	Tokentype_None           Tokentype = 256
	Tokentype_Error          Tokentype = 257
	Tokentype_Char           Tokentype = 258
	Tokentype_Binary         Tokentype = 259
	Tokentype_Octal          Tokentype = 260
	Tokentype_Int            Tokentype = 261
	Tokentype_Hex            Tokentype = 262
	Tokentype_Float          Tokentype = 263
	Tokentype_String         Tokentype = 264
	Tokentype_Symbol         Tokentype = 265
	Tokentype_Identifier     Tokentype = 266
	Tokentype_IdentifierNull Tokentype = 267
	Tokentype_CommentSingle  Tokentype = 268
	Tokentype_CommentMulti   Tokentype = 269
)

// Traversetype is a representation of the C type TraverseType.
type Traversetype int

const (
	Traversetype_InOrder    Traversetype = 0
	Traversetype_PreOrder   Traversetype = 1
	Traversetype_PostOrder  Traversetype = 2
	Traversetype_LevelOrder Traversetype = 3
)

// Unicodebreaktype is a representation of the C type UnicodeBreakType.
type Unicodebreaktype int

const (
	Unicodebreaktype_Mandatory                  Unicodebreaktype = 0
	Unicodebreaktype_CarriageReturn             Unicodebreaktype = 1
	Unicodebreaktype_LineFeed                   Unicodebreaktype = 2
	Unicodebreaktype_CombiningMark              Unicodebreaktype = 3
	Unicodebreaktype_Surrogate                  Unicodebreaktype = 4
	Unicodebreaktype_ZeroWidthSpace             Unicodebreaktype = 5
	Unicodebreaktype_Inseparable                Unicodebreaktype = 6
	Unicodebreaktype_NonBreakingGlue            Unicodebreaktype = 7
	Unicodebreaktype_Contingent                 Unicodebreaktype = 8
	Unicodebreaktype_Space                      Unicodebreaktype = 9
	Unicodebreaktype_After                      Unicodebreaktype = 10
	Unicodebreaktype_Before                     Unicodebreaktype = 11
	Unicodebreaktype_BeforeAndAfter             Unicodebreaktype = 12
	Unicodebreaktype_Hyphen                     Unicodebreaktype = 13
	Unicodebreaktype_NonStarter                 Unicodebreaktype = 14
	Unicodebreaktype_OpenPunctuation            Unicodebreaktype = 15
	Unicodebreaktype_ClosePunctuation           Unicodebreaktype = 16
	Unicodebreaktype_Quotation                  Unicodebreaktype = 17
	Unicodebreaktype_Exclamation                Unicodebreaktype = 18
	Unicodebreaktype_Ideographic                Unicodebreaktype = 19
	Unicodebreaktype_Numeric                    Unicodebreaktype = 20
	Unicodebreaktype_InfixSeparator             Unicodebreaktype = 21
	Unicodebreaktype_Symbol                     Unicodebreaktype = 22
	Unicodebreaktype_Alphabetic                 Unicodebreaktype = 23
	Unicodebreaktype_Prefix                     Unicodebreaktype = 24
	Unicodebreaktype_Postfix                    Unicodebreaktype = 25
	Unicodebreaktype_ComplexContext             Unicodebreaktype = 26
	Unicodebreaktype_Ambiguous                  Unicodebreaktype = 27
	Unicodebreaktype_Unknown                    Unicodebreaktype = 28
	Unicodebreaktype_NextLine                   Unicodebreaktype = 29
	Unicodebreaktype_WordJoiner                 Unicodebreaktype = 30
	Unicodebreaktype_HangulLJamo                Unicodebreaktype = 31
	Unicodebreaktype_HangulVJamo                Unicodebreaktype = 32
	Unicodebreaktype_HangulTJamo                Unicodebreaktype = 33
	Unicodebreaktype_HangulLvSyllable           Unicodebreaktype = 34
	Unicodebreaktype_HangulLvtSyllable          Unicodebreaktype = 35
	Unicodebreaktype_CloseParanthesis           Unicodebreaktype = 36
	Unicodebreaktype_ConditionalJapaneseStarter Unicodebreaktype = 37
	Unicodebreaktype_HebrewLetter               Unicodebreaktype = 38
	Unicodebreaktype_RegionalIndicator          Unicodebreaktype = 39
	Unicodebreaktype_EmojiBase                  Unicodebreaktype = 40
	Unicodebreaktype_EmojiModifier              Unicodebreaktype = 41
	Unicodebreaktype_ZeroWidthJoiner            Unicodebreaktype = 42
)

// Unicodescript is a representation of the C type UnicodeScript.
type Unicodescript int

const (
	Unicodescript_InvalidCode           Unicodescript = -1
	Unicodescript_Common                Unicodescript = 0
	Unicodescript_Inherited             Unicodescript = 1
	Unicodescript_Arabic                Unicodescript = 2
	Unicodescript_Armenian              Unicodescript = 3
	Unicodescript_Bengali               Unicodescript = 4
	Unicodescript_Bopomofo              Unicodescript = 5
	Unicodescript_Cherokee              Unicodescript = 6
	Unicodescript_Coptic                Unicodescript = 7
	Unicodescript_Cyrillic              Unicodescript = 8
	Unicodescript_Deseret               Unicodescript = 9
	Unicodescript_Devanagari            Unicodescript = 10
	Unicodescript_Ethiopic              Unicodescript = 11
	Unicodescript_Georgian              Unicodescript = 12
	Unicodescript_Gothic                Unicodescript = 13
	Unicodescript_Greek                 Unicodescript = 14
	Unicodescript_Gujarati              Unicodescript = 15
	Unicodescript_Gurmukhi              Unicodescript = 16
	Unicodescript_Han                   Unicodescript = 17
	Unicodescript_Hangul                Unicodescript = 18
	Unicodescript_Hebrew                Unicodescript = 19
	Unicodescript_Hiragana              Unicodescript = 20
	Unicodescript_Kannada               Unicodescript = 21
	Unicodescript_Katakana              Unicodescript = 22
	Unicodescript_Khmer                 Unicodescript = 23
	Unicodescript_Lao                   Unicodescript = 24
	Unicodescript_Latin                 Unicodescript = 25
	Unicodescript_Malayalam             Unicodescript = 26
	Unicodescript_Mongolian             Unicodescript = 27
	Unicodescript_Myanmar               Unicodescript = 28
	Unicodescript_Ogham                 Unicodescript = 29
	Unicodescript_OldItalic             Unicodescript = 30
	Unicodescript_Oriya                 Unicodescript = 31
	Unicodescript_Runic                 Unicodescript = 32
	Unicodescript_Sinhala               Unicodescript = 33
	Unicodescript_Syriac                Unicodescript = 34
	Unicodescript_Tamil                 Unicodescript = 35
	Unicodescript_Telugu                Unicodescript = 36
	Unicodescript_Thaana                Unicodescript = 37
	Unicodescript_Thai                  Unicodescript = 38
	Unicodescript_Tibetan               Unicodescript = 39
	Unicodescript_CanadianAboriginal    Unicodescript = 40
	Unicodescript_Yi                    Unicodescript = 41
	Unicodescript_Tagalog               Unicodescript = 42
	Unicodescript_Hanunoo               Unicodescript = 43
	Unicodescript_Buhid                 Unicodescript = 44
	Unicodescript_Tagbanwa              Unicodescript = 45
	Unicodescript_Braille               Unicodescript = 46
	Unicodescript_Cypriot               Unicodescript = 47
	Unicodescript_Limbu                 Unicodescript = 48
	Unicodescript_Osmanya               Unicodescript = 49
	Unicodescript_Shavian               Unicodescript = 50
	Unicodescript_LinearB               Unicodescript = 51
	Unicodescript_TaiLe                 Unicodescript = 52
	Unicodescript_Ugaritic              Unicodescript = 53
	Unicodescript_NewTaiLue             Unicodescript = 54
	Unicodescript_Buginese              Unicodescript = 55
	Unicodescript_Glagolitic            Unicodescript = 56
	Unicodescript_Tifinagh              Unicodescript = 57
	Unicodescript_SylotiNagri           Unicodescript = 58
	Unicodescript_OldPersian            Unicodescript = 59
	Unicodescript_Kharoshthi            Unicodescript = 60
	Unicodescript_Unknown               Unicodescript = 61
	Unicodescript_Balinese              Unicodescript = 62
	Unicodescript_Cuneiform             Unicodescript = 63
	Unicodescript_Phoenician            Unicodescript = 64
	Unicodescript_PhagsPa               Unicodescript = 65
	Unicodescript_Nko                   Unicodescript = 66
	Unicodescript_KayahLi               Unicodescript = 67
	Unicodescript_Lepcha                Unicodescript = 68
	Unicodescript_Rejang                Unicodescript = 69
	Unicodescript_Sundanese             Unicodescript = 70
	Unicodescript_Saurashtra            Unicodescript = 71
	Unicodescript_Cham                  Unicodescript = 72
	Unicodescript_OlChiki               Unicodescript = 73
	Unicodescript_Vai                   Unicodescript = 74
	Unicodescript_Carian                Unicodescript = 75
	Unicodescript_Lycian                Unicodescript = 76
	Unicodescript_Lydian                Unicodescript = 77
	Unicodescript_Avestan               Unicodescript = 78
	Unicodescript_Bamum                 Unicodescript = 79
	Unicodescript_EgyptianHieroglyphs   Unicodescript = 80
	Unicodescript_ImperialAramaic       Unicodescript = 81
	Unicodescript_InscriptionalPahlavi  Unicodescript = 82
	Unicodescript_InscriptionalParthian Unicodescript = 83
	Unicodescript_Javanese              Unicodescript = 84
	Unicodescript_Kaithi                Unicodescript = 85
	Unicodescript_Lisu                  Unicodescript = 86
	Unicodescript_MeeteiMayek           Unicodescript = 87
	Unicodescript_OldSouthArabian       Unicodescript = 88
	Unicodescript_OldTurkic             Unicodescript = 89
	Unicodescript_Samaritan             Unicodescript = 90
	Unicodescript_TaiTham               Unicodescript = 91
	Unicodescript_TaiViet               Unicodescript = 92
	Unicodescript_Batak                 Unicodescript = 93
	Unicodescript_Brahmi                Unicodescript = 94
	Unicodescript_Mandaic               Unicodescript = 95
	Unicodescript_Chakma                Unicodescript = 96
	Unicodescript_MeroiticCursive       Unicodescript = 97
	Unicodescript_MeroiticHieroglyphs   Unicodescript = 98
	Unicodescript_Miao                  Unicodescript = 99
	Unicodescript_Sharada               Unicodescript = 100
	Unicodescript_SoraSompeng           Unicodescript = 101
	Unicodescript_Takri                 Unicodescript = 102
	Unicodescript_BassaVah              Unicodescript = 103
	Unicodescript_CaucasianAlbanian     Unicodescript = 104
	Unicodescript_Duployan              Unicodescript = 105
	Unicodescript_Elbasan               Unicodescript = 106
	Unicodescript_Grantha               Unicodescript = 107
	Unicodescript_Khojki                Unicodescript = 108
	Unicodescript_Khudawadi             Unicodescript = 109
	Unicodescript_LinearA               Unicodescript = 110
	Unicodescript_Mahajani              Unicodescript = 111
	Unicodescript_Manichaean            Unicodescript = 112
	Unicodescript_MendeKikakui          Unicodescript = 113
	Unicodescript_Modi                  Unicodescript = 114
	Unicodescript_Mro                   Unicodescript = 115
	Unicodescript_Nabataean             Unicodescript = 116
	Unicodescript_OldNorthArabian       Unicodescript = 117
	Unicodescript_OldPermic             Unicodescript = 118
	Unicodescript_PahawhHmong           Unicodescript = 119
	Unicodescript_Palmyrene             Unicodescript = 120
	Unicodescript_PauCinHau             Unicodescript = 121
	Unicodescript_PsalterPahlavi        Unicodescript = 122
	Unicodescript_Siddham               Unicodescript = 123
	Unicodescript_Tirhuta               Unicodescript = 124
	Unicodescript_WarangCiti            Unicodescript = 125
	Unicodescript_Ahom                  Unicodescript = 126
	Unicodescript_AnatolianHieroglyphs  Unicodescript = 127
	Unicodescript_Hatran                Unicodescript = 128
	Unicodescript_Multani               Unicodescript = 129
	Unicodescript_OldHungarian          Unicodescript = 130
	Unicodescript_Signwriting           Unicodescript = 131
	Unicodescript_Adlam                 Unicodescript = 132
	Unicodescript_Bhaiksuki             Unicodescript = 133
	Unicodescript_Marchen               Unicodescript = 134
	Unicodescript_Newa                  Unicodescript = 135
	Unicodescript_Osage                 Unicodescript = 136
	Unicodescript_Tangut                Unicodescript = 137
	Unicodescript_MasaramGondi          Unicodescript = 138
	Unicodescript_Nushu                 Unicodescript = 139
	Unicodescript_Soyombo               Unicodescript = 140
	Unicodescript_ZanabazarSquare       Unicodescript = 141
)

// Unicodetype is a representation of the C type UnicodeType.
type Unicodetype int

const (
	Unicodetype_Control            Unicodetype = 0
	Unicodetype_Format             Unicodetype = 1
	Unicodetype_Unassigned         Unicodetype = 2
	Unicodetype_PrivateUse         Unicodetype = 3
	Unicodetype_Surrogate          Unicodetype = 4
	Unicodetype_LowercaseLetter    Unicodetype = 5
	Unicodetype_ModifierLetter     Unicodetype = 6
	Unicodetype_OtherLetter        Unicodetype = 7
	Unicodetype_TitlecaseLetter    Unicodetype = 8
	Unicodetype_UppercaseLetter    Unicodetype = 9
	Unicodetype_SpacingMark        Unicodetype = 10
	Unicodetype_EnclosingMark      Unicodetype = 11
	Unicodetype_NonSpacingMark     Unicodetype = 12
	Unicodetype_DecimalNumber      Unicodetype = 13
	Unicodetype_LetterNumber       Unicodetype = 14
	Unicodetype_OtherNumber        Unicodetype = 15
	Unicodetype_ConnectPunctuation Unicodetype = 16
	Unicodetype_DashPunctuation    Unicodetype = 17
	Unicodetype_ClosePunctuation   Unicodetype = 18
	Unicodetype_FinalPunctuation   Unicodetype = 19
	Unicodetype_InitialPunctuation Unicodetype = 20
	Unicodetype_OtherPunctuation   Unicodetype = 21
	Unicodetype_OpenPunctuation    Unicodetype = 22
	Unicodetype_CurrencySymbol     Unicodetype = 23
	Unicodetype_ModifierSymbol     Unicodetype = 24
	Unicodetype_MathSymbol         Unicodetype = 25
	Unicodetype_OtherSymbol        Unicodetype = 26
	Unicodetype_LineSeparator      Unicodetype = 27
	Unicodetype_ParagraphSeparator Unicodetype = 28
	Unicodetype_SpaceSeparator     Unicodetype = 29
)

// Userdirectory is a representation of the C type UserDirectory.
//
// since 2.14
type Userdirectory int

const (
	Userdirectory_DirectoryDesktop     Userdirectory = 0
	Userdirectory_DirectoryDocuments   Userdirectory = 1
	Userdirectory_DirectoryDownload    Userdirectory = 2
	Userdirectory_DirectoryMusic       Userdirectory = 3
	Userdirectory_DirectoryPictures    Userdirectory = 4
	Userdirectory_DirectoryPublicShare Userdirectory = 5
	Userdirectory_DirectoryTemplates   Userdirectory = 6
	Userdirectory_DirectoryVideos      Userdirectory = 7
	Userdirectory_NDirectories         Userdirectory = 8
)

// Variantclass is a representation of the C type VariantClass.
//
// since 2.24
type Variantclass int

const (
	Variantclass_Boolean    Variantclass = 98
	Variantclass_Byte       Variantclass = 121
	Variantclass_Int16      Variantclass = 110
	Variantclass_Uint16     Variantclass = 113
	Variantclass_Int32      Variantclass = 105
	Variantclass_Uint32     Variantclass = 117
	Variantclass_Int64      Variantclass = 120
	Variantclass_Uint64     Variantclass = 116
	Variantclass_Handle     Variantclass = 104
	Variantclass_Double     Variantclass = 100
	Variantclass_String     Variantclass = 115
	Variantclass_ObjectPath Variantclass = 111
	Variantclass_Signature  Variantclass = 103
	Variantclass_Variant    Variantclass = 118
	Variantclass_Maybe      Variantclass = 109
	Variantclass_Array      Variantclass = 97
	Variantclass_Tuple      Variantclass = 40
	Variantclass_DictEntry  Variantclass = 123
)

// Variantparseerror is a representation of the C type VariantParseError.
type Variantparseerror int

const (
	Variantparseerror_Failed                     Variantparseerror = 0
	Variantparseerror_BasicTypeExpected          Variantparseerror = 1
	Variantparseerror_CannotInferType            Variantparseerror = 2
	Variantparseerror_DefiniteTypeExpected       Variantparseerror = 3
	Variantparseerror_InputNotAtEnd              Variantparseerror = 4
	Variantparseerror_InvalidCharacter           Variantparseerror = 5
	Variantparseerror_InvalidFormatString        Variantparseerror = 6
	Variantparseerror_InvalidObjectPath          Variantparseerror = 7
	Variantparseerror_InvalidSignature           Variantparseerror = 8
	Variantparseerror_InvalidTypeString          Variantparseerror = 9
	Variantparseerror_NoCommonType               Variantparseerror = 10
	Variantparseerror_NumberOutOfRange           Variantparseerror = 11
	Variantparseerror_NumberTooBig               Variantparseerror = 12
	Variantparseerror_TypeError                  Variantparseerror = 13
	Variantparseerror_UnexpectedToken            Variantparseerror = 14
	Variantparseerror_UnknownKeyword             Variantparseerror = 15
	Variantparseerror_UnterminatedStringConstant Variantparseerror = 16
	Variantparseerror_ValueExpected              Variantparseerror = 17
)

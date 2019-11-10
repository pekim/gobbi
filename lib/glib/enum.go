// Code generated - DO NOT EDIT.

package glib

type BookmarkFileError uint32

const (
	BookmarkFileError_InvalidUri       BookmarkFileError = int64(0)
	BookmarkFileError_InvalidValue     BookmarkFileError = int64(1)
	BookmarkFileError_AppNotRegistered BookmarkFileError = int64(2)
	BookmarkFileError_UriNotFound      BookmarkFileError = int64(3)
	BookmarkFileError_Read             BookmarkFileError = int64(4)
	BookmarkFileError_UnknownEncoding  BookmarkFileError = int64(5)
	BookmarkFileError_Write            BookmarkFileError = int64(6)
	BookmarkFileError_FileNotFound     BookmarkFileError = int64(7)
)

type ChecksumType uint32

const (
	ChecksumType_Md5    ChecksumType = int64(0)
	ChecksumType_Sha1   ChecksumType = int64(1)
	ChecksumType_Sha256 ChecksumType = int64(2)
	ChecksumType_Sha512 ChecksumType = int64(3)
	ChecksumType_Sha384 ChecksumType = int64(4)
)

type ConvertError uint32

const (
	ConvertError_NoConversion    ConvertError = int64(0)
	ConvertError_IllegalSequence ConvertError = int64(1)
	ConvertError_Failed          ConvertError = int64(2)
	ConvertError_PartialInput    ConvertError = int64(3)
	ConvertError_BadUri          ConvertError = int64(4)
	ConvertError_NotAbsolutePath ConvertError = int64(5)
	ConvertError_NoMemory        ConvertError = int64(6)
	ConvertError_EmbeddedNul     ConvertError = int64(7)
)

type DateDMY uint32

const (
	DateDMY_Day   DateDMY = int64(0)
	DateDMY_Month DateDMY = int64(1)
	DateDMY_Year  DateDMY = int64(2)
)

type DateMonth uint32

const (
	DateMonth_BadMonth  DateMonth = int64(0)
	DateMonth_January   DateMonth = int64(1)
	DateMonth_February  DateMonth = int64(2)
	DateMonth_March     DateMonth = int64(3)
	DateMonth_April     DateMonth = int64(4)
	DateMonth_May       DateMonth = int64(5)
	DateMonth_June      DateMonth = int64(6)
	DateMonth_July      DateMonth = int64(7)
	DateMonth_August    DateMonth = int64(8)
	DateMonth_September DateMonth = int64(9)
	DateMonth_October   DateMonth = int64(10)
	DateMonth_November  DateMonth = int64(11)
	DateMonth_December  DateMonth = int64(12)
)

type DateWeekday uint32

const (
	DateWeekday_BadWeekday DateWeekday = int64(0)
	DateWeekday_Monday     DateWeekday = int64(1)
	DateWeekday_Tuesday    DateWeekday = int64(2)
	DateWeekday_Wednesday  DateWeekday = int64(3)
	DateWeekday_Thursday   DateWeekday = int64(4)
	DateWeekday_Friday     DateWeekday = int64(5)
	DateWeekday_Saturday   DateWeekday = int64(6)
	DateWeekday_Sunday     DateWeekday = int64(7)
)

type ErrorType uint32

const (
	ErrorType_Unknown           ErrorType = int64(0)
	ErrorType_UnexpEof          ErrorType = int64(1)
	ErrorType_UnexpEofInString  ErrorType = int64(2)
	ErrorType_UnexpEofInComment ErrorType = int64(3)
	ErrorType_NonDigitInConst   ErrorType = int64(4)
	ErrorType_DigitRadix        ErrorType = int64(5)
	ErrorType_FloatRadix        ErrorType = int64(6)
	ErrorType_FloatMalformed    ErrorType = int64(7)
)

type FileError uint32

const (
	FileError_Exist       FileError = int64(0)
	FileError_Isdir       FileError = int64(1)
	FileError_Acces       FileError = int64(2)
	FileError_Nametoolong FileError = int64(3)
	FileError_Noent       FileError = int64(4)
	FileError_Notdir      FileError = int64(5)
	FileError_Nxio        FileError = int64(6)
	FileError_Nodev       FileError = int64(7)
	FileError_Rofs        FileError = int64(8)
	FileError_Txtbsy      FileError = int64(9)
	FileError_Fault       FileError = int64(10)
	FileError_Loop        FileError = int64(11)
	FileError_Nospc       FileError = int64(12)
	FileError_Nomem       FileError = int64(13)
	FileError_Mfile       FileError = int64(14)
	FileError_Nfile       FileError = int64(15)
	FileError_Badf        FileError = int64(16)
	FileError_Inval       FileError = int64(17)
	FileError_Pipe        FileError = int64(18)
	FileError_Again       FileError = int64(19)
	FileError_Intr        FileError = int64(20)
	FileError_Io          FileError = int64(21)
	FileError_Perm        FileError = int64(22)
	FileError_Nosys       FileError = int64(23)
	FileError_Failed      FileError = int64(24)
)

type IOChannelError uint32

const (
	IOChannelError_Fbig     IOChannelError = int64(0)
	IOChannelError_Inval    IOChannelError = int64(1)
	IOChannelError_Io       IOChannelError = int64(2)
	IOChannelError_Isdir    IOChannelError = int64(3)
	IOChannelError_Nospc    IOChannelError = int64(4)
	IOChannelError_Nxio     IOChannelError = int64(5)
	IOChannelError_Overflow IOChannelError = int64(6)
	IOChannelError_Pipe     IOChannelError = int64(7)
	IOChannelError_Failed   IOChannelError = int64(8)
)

type IOError uint32

const (
	IOError_None    IOError = int64(0)
	IOError_Again   IOError = int64(1)
	IOError_Inval   IOError = int64(2)
	IOError_Unknown IOError = int64(3)
)

type IOStatus uint32

const (
	IOStatus_Error  IOStatus = int64(0)
	IOStatus_Normal IOStatus = int64(1)
	IOStatus_Eof    IOStatus = int64(2)
	IOStatus_Again  IOStatus = int64(3)
)

type KeyFileError uint32

const (
	KeyFileError_UnknownEncoding KeyFileError = int64(0)
	KeyFileError_Parse           KeyFileError = int64(1)
	KeyFileError_NotFound        KeyFileError = int64(2)
	KeyFileError_KeyNotFound     KeyFileError = int64(3)
	KeyFileError_GroupNotFound   KeyFileError = int64(4)
	KeyFileError_InvalidValue    KeyFileError = int64(5)
)

type LogWriterOutput uint32

const (
	LogWriterOutput_Handled   LogWriterOutput = int64(1)
	LogWriterOutput_Unhandled LogWriterOutput = int64(0)
)

type MarkupError uint32

const (
	MarkupError_BadUtf8          MarkupError = int64(0)
	MarkupError_Empty            MarkupError = int64(1)
	MarkupError_Parse            MarkupError = int64(2)
	MarkupError_UnknownElement   MarkupError = int64(3)
	MarkupError_UnknownAttribute MarkupError = int64(4)
	MarkupError_InvalidContent   MarkupError = int64(5)
	MarkupError_MissingAttribute MarkupError = int64(6)
)

type NormalizeMode uint32

const (
	NormalizeMode_Default        NormalizeMode = int64(0)
	NormalizeMode_Nfd            NormalizeMode = int64(0)
	NormalizeMode_DefaultCompose NormalizeMode = int64(1)
	NormalizeMode_Nfc            NormalizeMode = int64(1)
	NormalizeMode_All            NormalizeMode = int64(2)
	NormalizeMode_Nfkd           NormalizeMode = int64(2)
	NormalizeMode_AllCompose     NormalizeMode = int64(3)
	NormalizeMode_Nfkc           NormalizeMode = int64(3)
)

type NumberParserError uint32

const (
	NumberParserError_Invalid     NumberParserError = int64(0)
	NumberParserError_OutOfBounds NumberParserError = int64(1)
)

type OnceStatus uint32

const (
	OnceStatus_Notcalled OnceStatus = int64(0)
	OnceStatus_Progress  OnceStatus = int64(1)
	OnceStatus_Ready     OnceStatus = int64(2)
)

type OptionArg uint32

const (
	OptionArg_None          OptionArg = int64(0)
	OptionArg_String        OptionArg = int64(1)
	OptionArg_Int           OptionArg = int64(2)
	OptionArg_Callback      OptionArg = int64(3)
	OptionArg_Filename      OptionArg = int64(4)
	OptionArg_StringArray   OptionArg = int64(5)
	OptionArg_FilenameArray OptionArg = int64(6)
	OptionArg_Double        OptionArg = int64(7)
	OptionArg_Int64         OptionArg = int64(8)
)

type OptionError uint32

const (
	OptionError_UnknownOption OptionError = int64(0)
	OptionError_BadValue      OptionError = int64(1)
	OptionError_Failed        OptionError = int64(2)
)

type RegexError uint32

const (
	RegexError_Compile                                  RegexError = int64(0)
	RegexError_Optimize                                 RegexError = int64(1)
	RegexError_Replace                                  RegexError = int64(2)
	RegexError_Match                                    RegexError = int64(3)
	RegexError_Internal                                 RegexError = int64(4)
	RegexError_StrayBackslash                           RegexError = int64(101)
	RegexError_MissingControlChar                       RegexError = int64(102)
	RegexError_UnrecognizedEscape                       RegexError = int64(103)
	RegexError_QuantifiersOutOfOrder                    RegexError = int64(104)
	RegexError_QuantifierTooBig                         RegexError = int64(105)
	RegexError_UnterminatedCharacterClass               RegexError = int64(106)
	RegexError_InvalidEscapeInCharacterClass            RegexError = int64(107)
	RegexError_RangeOutOfOrder                          RegexError = int64(108)
	RegexError_NothingToRepeat                          RegexError = int64(109)
	RegexError_UnrecognizedCharacter                    RegexError = int64(112)
	RegexError_PosixNamedClassOutsideClass              RegexError = int64(113)
	RegexError_UnmatchedParenthesis                     RegexError = int64(114)
	RegexError_InexistentSubpatternReference            RegexError = int64(115)
	RegexError_UnterminatedComment                      RegexError = int64(118)
	RegexError_ExpressionTooLarge                       RegexError = int64(120)
	RegexError_MemoryError                              RegexError = int64(121)
	RegexError_VariableLengthLookbehind                 RegexError = int64(125)
	RegexError_MalformedCondition                       RegexError = int64(126)
	RegexError_TooManyConditionalBranches               RegexError = int64(127)
	RegexError_AssertionExpected                        RegexError = int64(128)
	RegexError_UnknownPosixClassName                    RegexError = int64(130)
	RegexError_PosixCollatingElementsNotSupported       RegexError = int64(131)
	RegexError_HexCodeTooLarge                          RegexError = int64(134)
	RegexError_InvalidCondition                         RegexError = int64(135)
	RegexError_SingleByteMatchInLookbehind              RegexError = int64(136)
	RegexError_InfiniteLoop                             RegexError = int64(140)
	RegexError_MissingSubpatternNameTerminator          RegexError = int64(142)
	RegexError_DuplicateSubpatternName                  RegexError = int64(143)
	RegexError_MalformedProperty                        RegexError = int64(146)
	RegexError_UnknownProperty                          RegexError = int64(147)
	RegexError_SubpatternNameTooLong                    RegexError = int64(148)
	RegexError_TooManySubpatterns                       RegexError = int64(149)
	RegexError_InvalidOctalValue                        RegexError = int64(151)
	RegexError_TooManyBranchesInDefine                  RegexError = int64(154)
	RegexError_DefineRepetion                           RegexError = int64(155)
	RegexError_InconsistentNewlineOptions               RegexError = int64(156)
	RegexError_MissingBackReference                     RegexError = int64(157)
	RegexError_InvalidRelativeReference                 RegexError = int64(158)
	RegexError_BacktrackingControlVerbArgumentForbidden RegexError = int64(159)
	RegexError_UnknownBacktrackingControlVerb           RegexError = int64(160)
	RegexError_NumberTooBig                             RegexError = int64(161)
	RegexError_MissingSubpatternName                    RegexError = int64(162)
	RegexError_MissingDigit                             RegexError = int64(163)
	RegexError_InvalidDataCharacter                     RegexError = int64(164)
	RegexError_ExtraSubpatternName                      RegexError = int64(165)
	RegexError_BacktrackingControlVerbArgumentRequired  RegexError = int64(166)
	RegexError_InvalidControlChar                       RegexError = int64(168)
	RegexError_MissingName                              RegexError = int64(169)
	RegexError_NotSupportedInClass                      RegexError = int64(171)
	RegexError_TooManyForwardReferences                 RegexError = int64(172)
	RegexError_NameTooLong                              RegexError = int64(175)
	RegexError_CharacterValueTooLarge                   RegexError = int64(176)
)

type SeekType uint32

const (
	SeekType_Cur SeekType = int64(0)
	SeekType_Set SeekType = int64(1)
	SeekType_End SeekType = int64(2)
)

type ShellError uint32

const (
	ShellError_BadQuoting  ShellError = int64(0)
	ShellError_EmptyString ShellError = int64(1)
	ShellError_Failed      ShellError = int64(2)
)

type SliceConfig uint32

const (
	SliceConfig_AlwaysMalloc      SliceConfig = int64(1)
	SliceConfig_BypassMagazines   SliceConfig = int64(2)
	SliceConfig_WorkingSetMsecs   SliceConfig = int64(3)
	SliceConfig_ColorIncrement    SliceConfig = int64(4)
	SliceConfig_ChunkSizes        SliceConfig = int64(5)
	SliceConfig_ContentionCounter SliceConfig = int64(6)
)

type SpawnError uint32

const (
	SpawnError_Fork        SpawnError = int64(0)
	SpawnError_Read        SpawnError = int64(1)
	SpawnError_Chdir       SpawnError = int64(2)
	SpawnError_Acces       SpawnError = int64(3)
	SpawnError_Perm        SpawnError = int64(4)
	SpawnError_TooBig      SpawnError = int64(5)
	SpawnError_2big        SpawnError = int64(5)
	SpawnError_Noexec      SpawnError = int64(6)
	SpawnError_Nametoolong SpawnError = int64(7)
	SpawnError_Noent       SpawnError = int64(8)
	SpawnError_Nomem       SpawnError = int64(9)
	SpawnError_Notdir      SpawnError = int64(10)
	SpawnError_Loop        SpawnError = int64(11)
	SpawnError_Txtbusy     SpawnError = int64(12)
	SpawnError_Io          SpawnError = int64(13)
	SpawnError_Nfile       SpawnError = int64(14)
	SpawnError_Mfile       SpawnError = int64(15)
	SpawnError_Inval       SpawnError = int64(16)
	SpawnError_Isdir       SpawnError = int64(17)
	SpawnError_Libbad      SpawnError = int64(18)
	SpawnError_Failed      SpawnError = int64(19)
)

type TestFileType uint32

const (
	TestFileType_Dist  TestFileType = int64(0)
	TestFileType_Built TestFileType = int64(1)
)

type TestLogType uint32

const (
	TestLogType_None        TestLogType = int64(0)
	TestLogType_Error       TestLogType = int64(1)
	TestLogType_StartBinary TestLogType = int64(2)
	TestLogType_ListCase    TestLogType = int64(3)
	TestLogType_SkipCase    TestLogType = int64(4)
	TestLogType_StartCase   TestLogType = int64(5)
	TestLogType_StopCase    TestLogType = int64(6)
	TestLogType_MinResult   TestLogType = int64(7)
	TestLogType_MaxResult   TestLogType = int64(8)
	TestLogType_Message     TestLogType = int64(9)
	TestLogType_StartSuite  TestLogType = int64(10)
	TestLogType_StopSuite   TestLogType = int64(11)
)

type TestResult uint32

const (
	TestResult_Success    TestResult = int64(0)
	TestResult_Skipped    TestResult = int64(1)
	TestResult_Failure    TestResult = int64(2)
	TestResult_Incomplete TestResult = int64(3)
)

type ThreadError uint32

const (
	ThreadError_ThreadErrorAgain ThreadError = int64(0)
)

type TimeType uint32

const (
	TimeType_Standard  TimeType = int64(0)
	TimeType_Daylight  TimeType = int64(1)
	TimeType_Universal TimeType = int64(2)
)

type TokenType uint32

const (
	TokenType_Eof            TokenType = int64(0)
	TokenType_LeftParen      TokenType = int64(40)
	TokenType_RightParen     TokenType = int64(41)
	TokenType_LeftCurly      TokenType = int64(123)
	TokenType_RightCurly     TokenType = int64(125)
	TokenType_LeftBrace      TokenType = int64(91)
	TokenType_RightBrace     TokenType = int64(93)
	TokenType_EqualSign      TokenType = int64(61)
	TokenType_Comma          TokenType = int64(44)
	TokenType_None           TokenType = int64(256)
	TokenType_Error          TokenType = int64(257)
	TokenType_Char           TokenType = int64(258)
	TokenType_Binary         TokenType = int64(259)
	TokenType_Octal          TokenType = int64(260)
	TokenType_Int            TokenType = int64(261)
	TokenType_Hex            TokenType = int64(262)
	TokenType_Float          TokenType = int64(263)
	TokenType_String         TokenType = int64(264)
	TokenType_Symbol         TokenType = int64(265)
	TokenType_Identifier     TokenType = int64(266)
	TokenType_IdentifierNull TokenType = int64(267)
	TokenType_CommentSingle  TokenType = int64(268)
	TokenType_CommentMulti   TokenType = int64(269)
)

type TraverseType uint32

const (
	TraverseType_InOrder    TraverseType = int64(0)
	TraverseType_PreOrder   TraverseType = int64(1)
	TraverseType_PostOrder  TraverseType = int64(2)
	TraverseType_LevelOrder TraverseType = int64(3)
)

type UnicodeBreakType uint32

const (
	UnicodeBreakType_Mandatory                  UnicodeBreakType = int64(0)
	UnicodeBreakType_CarriageReturn             UnicodeBreakType = int64(1)
	UnicodeBreakType_LineFeed                   UnicodeBreakType = int64(2)
	UnicodeBreakType_CombiningMark              UnicodeBreakType = int64(3)
	UnicodeBreakType_Surrogate                  UnicodeBreakType = int64(4)
	UnicodeBreakType_ZeroWidthSpace             UnicodeBreakType = int64(5)
	UnicodeBreakType_Inseparable                UnicodeBreakType = int64(6)
	UnicodeBreakType_NonBreakingGlue            UnicodeBreakType = int64(7)
	UnicodeBreakType_Contingent                 UnicodeBreakType = int64(8)
	UnicodeBreakType_Space                      UnicodeBreakType = int64(9)
	UnicodeBreakType_After                      UnicodeBreakType = int64(10)
	UnicodeBreakType_Before                     UnicodeBreakType = int64(11)
	UnicodeBreakType_BeforeAndAfter             UnicodeBreakType = int64(12)
	UnicodeBreakType_Hyphen                     UnicodeBreakType = int64(13)
	UnicodeBreakType_NonStarter                 UnicodeBreakType = int64(14)
	UnicodeBreakType_OpenPunctuation            UnicodeBreakType = int64(15)
	UnicodeBreakType_ClosePunctuation           UnicodeBreakType = int64(16)
	UnicodeBreakType_Quotation                  UnicodeBreakType = int64(17)
	UnicodeBreakType_Exclamation                UnicodeBreakType = int64(18)
	UnicodeBreakType_Ideographic                UnicodeBreakType = int64(19)
	UnicodeBreakType_Numeric                    UnicodeBreakType = int64(20)
	UnicodeBreakType_InfixSeparator             UnicodeBreakType = int64(21)
	UnicodeBreakType_Symbol                     UnicodeBreakType = int64(22)
	UnicodeBreakType_Alphabetic                 UnicodeBreakType = int64(23)
	UnicodeBreakType_Prefix                     UnicodeBreakType = int64(24)
	UnicodeBreakType_Postfix                    UnicodeBreakType = int64(25)
	UnicodeBreakType_ComplexContext             UnicodeBreakType = int64(26)
	UnicodeBreakType_Ambiguous                  UnicodeBreakType = int64(27)
	UnicodeBreakType_Unknown                    UnicodeBreakType = int64(28)
	UnicodeBreakType_NextLine                   UnicodeBreakType = int64(29)
	UnicodeBreakType_WordJoiner                 UnicodeBreakType = int64(30)
	UnicodeBreakType_HangulLJamo                UnicodeBreakType = int64(31)
	UnicodeBreakType_HangulVJamo                UnicodeBreakType = int64(32)
	UnicodeBreakType_HangulTJamo                UnicodeBreakType = int64(33)
	UnicodeBreakType_HangulLvSyllable           UnicodeBreakType = int64(34)
	UnicodeBreakType_HangulLvtSyllable          UnicodeBreakType = int64(35)
	UnicodeBreakType_CloseParanthesis           UnicodeBreakType = int64(36)
	UnicodeBreakType_ConditionalJapaneseStarter UnicodeBreakType = int64(37)
	UnicodeBreakType_HebrewLetter               UnicodeBreakType = int64(38)
	UnicodeBreakType_RegionalIndicator          UnicodeBreakType = int64(39)
	UnicodeBreakType_EmojiBase                  UnicodeBreakType = int64(40)
	UnicodeBreakType_EmojiModifier              UnicodeBreakType = int64(41)
	UnicodeBreakType_ZeroWidthJoiner            UnicodeBreakType = int64(42)
)

type UnicodeScript int32

const (
	UnicodeScript_InvalidCode           UnicodeScript = int64(-1)
	UnicodeScript_Common                UnicodeScript = int64(0)
	UnicodeScript_Inherited             UnicodeScript = int64(1)
	UnicodeScript_Arabic                UnicodeScript = int64(2)
	UnicodeScript_Armenian              UnicodeScript = int64(3)
	UnicodeScript_Bengali               UnicodeScript = int64(4)
	UnicodeScript_Bopomofo              UnicodeScript = int64(5)
	UnicodeScript_Cherokee              UnicodeScript = int64(6)
	UnicodeScript_Coptic                UnicodeScript = int64(7)
	UnicodeScript_Cyrillic              UnicodeScript = int64(8)
	UnicodeScript_Deseret               UnicodeScript = int64(9)
	UnicodeScript_Devanagari            UnicodeScript = int64(10)
	UnicodeScript_Ethiopic              UnicodeScript = int64(11)
	UnicodeScript_Georgian              UnicodeScript = int64(12)
	UnicodeScript_Gothic                UnicodeScript = int64(13)
	UnicodeScript_Greek                 UnicodeScript = int64(14)
	UnicodeScript_Gujarati              UnicodeScript = int64(15)
	UnicodeScript_Gurmukhi              UnicodeScript = int64(16)
	UnicodeScript_Han                   UnicodeScript = int64(17)
	UnicodeScript_Hangul                UnicodeScript = int64(18)
	UnicodeScript_Hebrew                UnicodeScript = int64(19)
	UnicodeScript_Hiragana              UnicodeScript = int64(20)
	UnicodeScript_Kannada               UnicodeScript = int64(21)
	UnicodeScript_Katakana              UnicodeScript = int64(22)
	UnicodeScript_Khmer                 UnicodeScript = int64(23)
	UnicodeScript_Lao                   UnicodeScript = int64(24)
	UnicodeScript_Latin                 UnicodeScript = int64(25)
	UnicodeScript_Malayalam             UnicodeScript = int64(26)
	UnicodeScript_Mongolian             UnicodeScript = int64(27)
	UnicodeScript_Myanmar               UnicodeScript = int64(28)
	UnicodeScript_Ogham                 UnicodeScript = int64(29)
	UnicodeScript_OldItalic             UnicodeScript = int64(30)
	UnicodeScript_Oriya                 UnicodeScript = int64(31)
	UnicodeScript_Runic                 UnicodeScript = int64(32)
	UnicodeScript_Sinhala               UnicodeScript = int64(33)
	UnicodeScript_Syriac                UnicodeScript = int64(34)
	UnicodeScript_Tamil                 UnicodeScript = int64(35)
	UnicodeScript_Telugu                UnicodeScript = int64(36)
	UnicodeScript_Thaana                UnicodeScript = int64(37)
	UnicodeScript_Thai                  UnicodeScript = int64(38)
	UnicodeScript_Tibetan               UnicodeScript = int64(39)
	UnicodeScript_CanadianAboriginal    UnicodeScript = int64(40)
	UnicodeScript_Yi                    UnicodeScript = int64(41)
	UnicodeScript_Tagalog               UnicodeScript = int64(42)
	UnicodeScript_Hanunoo               UnicodeScript = int64(43)
	UnicodeScript_Buhid                 UnicodeScript = int64(44)
	UnicodeScript_Tagbanwa              UnicodeScript = int64(45)
	UnicodeScript_Braille               UnicodeScript = int64(46)
	UnicodeScript_Cypriot               UnicodeScript = int64(47)
	UnicodeScript_Limbu                 UnicodeScript = int64(48)
	UnicodeScript_Osmanya               UnicodeScript = int64(49)
	UnicodeScript_Shavian               UnicodeScript = int64(50)
	UnicodeScript_LinearB               UnicodeScript = int64(51)
	UnicodeScript_TaiLe                 UnicodeScript = int64(52)
	UnicodeScript_Ugaritic              UnicodeScript = int64(53)
	UnicodeScript_NewTaiLue             UnicodeScript = int64(54)
	UnicodeScript_Buginese              UnicodeScript = int64(55)
	UnicodeScript_Glagolitic            UnicodeScript = int64(56)
	UnicodeScript_Tifinagh              UnicodeScript = int64(57)
	UnicodeScript_SylotiNagri           UnicodeScript = int64(58)
	UnicodeScript_OldPersian            UnicodeScript = int64(59)
	UnicodeScript_Kharoshthi            UnicodeScript = int64(60)
	UnicodeScript_Unknown               UnicodeScript = int64(61)
	UnicodeScript_Balinese              UnicodeScript = int64(62)
	UnicodeScript_Cuneiform             UnicodeScript = int64(63)
	UnicodeScript_Phoenician            UnicodeScript = int64(64)
	UnicodeScript_PhagsPa               UnicodeScript = int64(65)
	UnicodeScript_Nko                   UnicodeScript = int64(66)
	UnicodeScript_KayahLi               UnicodeScript = int64(67)
	UnicodeScript_Lepcha                UnicodeScript = int64(68)
	UnicodeScript_Rejang                UnicodeScript = int64(69)
	UnicodeScript_Sundanese             UnicodeScript = int64(70)
	UnicodeScript_Saurashtra            UnicodeScript = int64(71)
	UnicodeScript_Cham                  UnicodeScript = int64(72)
	UnicodeScript_OlChiki               UnicodeScript = int64(73)
	UnicodeScript_Vai                   UnicodeScript = int64(74)
	UnicodeScript_Carian                UnicodeScript = int64(75)
	UnicodeScript_Lycian                UnicodeScript = int64(76)
	UnicodeScript_Lydian                UnicodeScript = int64(77)
	UnicodeScript_Avestan               UnicodeScript = int64(78)
	UnicodeScript_Bamum                 UnicodeScript = int64(79)
	UnicodeScript_EgyptianHieroglyphs   UnicodeScript = int64(80)
	UnicodeScript_ImperialAramaic       UnicodeScript = int64(81)
	UnicodeScript_InscriptionalPahlavi  UnicodeScript = int64(82)
	UnicodeScript_InscriptionalParthian UnicodeScript = int64(83)
	UnicodeScript_Javanese              UnicodeScript = int64(84)
	UnicodeScript_Kaithi                UnicodeScript = int64(85)
	UnicodeScript_Lisu                  UnicodeScript = int64(86)
	UnicodeScript_MeeteiMayek           UnicodeScript = int64(87)
	UnicodeScript_OldSouthArabian       UnicodeScript = int64(88)
	UnicodeScript_OldTurkic             UnicodeScript = int64(89)
	UnicodeScript_Samaritan             UnicodeScript = int64(90)
	UnicodeScript_TaiTham               UnicodeScript = int64(91)
	UnicodeScript_TaiViet               UnicodeScript = int64(92)
	UnicodeScript_Batak                 UnicodeScript = int64(93)
	UnicodeScript_Brahmi                UnicodeScript = int64(94)
	UnicodeScript_Mandaic               UnicodeScript = int64(95)
	UnicodeScript_Chakma                UnicodeScript = int64(96)
	UnicodeScript_MeroiticCursive       UnicodeScript = int64(97)
	UnicodeScript_MeroiticHieroglyphs   UnicodeScript = int64(98)
	UnicodeScript_Miao                  UnicodeScript = int64(99)
	UnicodeScript_Sharada               UnicodeScript = int64(100)
	UnicodeScript_SoraSompeng           UnicodeScript = int64(101)
	UnicodeScript_Takri                 UnicodeScript = int64(102)
	UnicodeScript_BassaVah              UnicodeScript = int64(103)
	UnicodeScript_CaucasianAlbanian     UnicodeScript = int64(104)
	UnicodeScript_Duployan              UnicodeScript = int64(105)
	UnicodeScript_Elbasan               UnicodeScript = int64(106)
	UnicodeScript_Grantha               UnicodeScript = int64(107)
	UnicodeScript_Khojki                UnicodeScript = int64(108)
	UnicodeScript_Khudawadi             UnicodeScript = int64(109)
	UnicodeScript_LinearA               UnicodeScript = int64(110)
	UnicodeScript_Mahajani              UnicodeScript = int64(111)
	UnicodeScript_Manichaean            UnicodeScript = int64(112)
	UnicodeScript_MendeKikakui          UnicodeScript = int64(113)
	UnicodeScript_Modi                  UnicodeScript = int64(114)
	UnicodeScript_Mro                   UnicodeScript = int64(115)
	UnicodeScript_Nabataean             UnicodeScript = int64(116)
	UnicodeScript_OldNorthArabian       UnicodeScript = int64(117)
	UnicodeScript_OldPermic             UnicodeScript = int64(118)
	UnicodeScript_PahawhHmong           UnicodeScript = int64(119)
	UnicodeScript_Palmyrene             UnicodeScript = int64(120)
	UnicodeScript_PauCinHau             UnicodeScript = int64(121)
	UnicodeScript_PsalterPahlavi        UnicodeScript = int64(122)
	UnicodeScript_Siddham               UnicodeScript = int64(123)
	UnicodeScript_Tirhuta               UnicodeScript = int64(124)
	UnicodeScript_WarangCiti            UnicodeScript = int64(125)
	UnicodeScript_Ahom                  UnicodeScript = int64(126)
	UnicodeScript_AnatolianHieroglyphs  UnicodeScript = int64(127)
	UnicodeScript_Hatran                UnicodeScript = int64(128)
	UnicodeScript_Multani               UnicodeScript = int64(129)
	UnicodeScript_OldHungarian          UnicodeScript = int64(130)
	UnicodeScript_Signwriting           UnicodeScript = int64(131)
	UnicodeScript_Adlam                 UnicodeScript = int64(132)
	UnicodeScript_Bhaiksuki             UnicodeScript = int64(133)
	UnicodeScript_Marchen               UnicodeScript = int64(134)
	UnicodeScript_Newa                  UnicodeScript = int64(135)
	UnicodeScript_Osage                 UnicodeScript = int64(136)
	UnicodeScript_Tangut                UnicodeScript = int64(137)
	UnicodeScript_MasaramGondi          UnicodeScript = int64(138)
	UnicodeScript_Nushu                 UnicodeScript = int64(139)
	UnicodeScript_Soyombo               UnicodeScript = int64(140)
	UnicodeScript_ZanabazarSquare       UnicodeScript = int64(141)
)

type UnicodeType uint32

const (
	UnicodeType_Control            UnicodeType = int64(0)
	UnicodeType_Format             UnicodeType = int64(1)
	UnicodeType_Unassigned         UnicodeType = int64(2)
	UnicodeType_PrivateUse         UnicodeType = int64(3)
	UnicodeType_Surrogate          UnicodeType = int64(4)
	UnicodeType_LowercaseLetter    UnicodeType = int64(5)
	UnicodeType_ModifierLetter     UnicodeType = int64(6)
	UnicodeType_OtherLetter        UnicodeType = int64(7)
	UnicodeType_TitlecaseLetter    UnicodeType = int64(8)
	UnicodeType_UppercaseLetter    UnicodeType = int64(9)
	UnicodeType_SpacingMark        UnicodeType = int64(10)
	UnicodeType_EnclosingMark      UnicodeType = int64(11)
	UnicodeType_NonSpacingMark     UnicodeType = int64(12)
	UnicodeType_DecimalNumber      UnicodeType = int64(13)
	UnicodeType_LetterNumber       UnicodeType = int64(14)
	UnicodeType_OtherNumber        UnicodeType = int64(15)
	UnicodeType_ConnectPunctuation UnicodeType = int64(16)
	UnicodeType_DashPunctuation    UnicodeType = int64(17)
	UnicodeType_ClosePunctuation   UnicodeType = int64(18)
	UnicodeType_FinalPunctuation   UnicodeType = int64(19)
	UnicodeType_InitialPunctuation UnicodeType = int64(20)
	UnicodeType_OtherPunctuation   UnicodeType = int64(21)
	UnicodeType_OpenPunctuation    UnicodeType = int64(22)
	UnicodeType_CurrencySymbol     UnicodeType = int64(23)
	UnicodeType_ModifierSymbol     UnicodeType = int64(24)
	UnicodeType_MathSymbol         UnicodeType = int64(25)
	UnicodeType_OtherSymbol        UnicodeType = int64(26)
	UnicodeType_LineSeparator      UnicodeType = int64(27)
	UnicodeType_ParagraphSeparator UnicodeType = int64(28)
	UnicodeType_SpaceSeparator     UnicodeType = int64(29)
)

type UserDirectory uint32

const (
	UserDirectory_DirectoryDesktop     UserDirectory = int64(0)
	UserDirectory_DirectoryDocuments   UserDirectory = int64(1)
	UserDirectory_DirectoryDownload    UserDirectory = int64(2)
	UserDirectory_DirectoryMusic       UserDirectory = int64(3)
	UserDirectory_DirectoryPictures    UserDirectory = int64(4)
	UserDirectory_DirectoryPublicShare UserDirectory = int64(5)
	UserDirectory_DirectoryTemplates   UserDirectory = int64(6)
	UserDirectory_DirectoryVideos      UserDirectory = int64(7)
	UserDirectory_NDirectories         UserDirectory = int64(8)
)

type VariantClass uint32

const (
	VariantClass_Boolean    VariantClass = int64(98)
	VariantClass_Byte       VariantClass = int64(121)
	VariantClass_Int16      VariantClass = int64(110)
	VariantClass_Uint16     VariantClass = int64(113)
	VariantClass_Int32      VariantClass = int64(105)
	VariantClass_Uint32     VariantClass = int64(117)
	VariantClass_Int64      VariantClass = int64(120)
	VariantClass_Uint64     VariantClass = int64(116)
	VariantClass_Handle     VariantClass = int64(104)
	VariantClass_Double     VariantClass = int64(100)
	VariantClass_String     VariantClass = int64(115)
	VariantClass_ObjectPath VariantClass = int64(111)
	VariantClass_Signature  VariantClass = int64(103)
	VariantClass_Variant    VariantClass = int64(118)
	VariantClass_Maybe      VariantClass = int64(109)
	VariantClass_Array      VariantClass = int64(97)
	VariantClass_Tuple      VariantClass = int64(40)
	VariantClass_DictEntry  VariantClass = int64(123)
)

type VariantParseError uint32

const (
	VariantParseError_Failed                     VariantParseError = int64(0)
	VariantParseError_BasicTypeExpected          VariantParseError = int64(1)
	VariantParseError_CannotInferType            VariantParseError = int64(2)
	VariantParseError_DefiniteTypeExpected       VariantParseError = int64(3)
	VariantParseError_InputNotAtEnd              VariantParseError = int64(4)
	VariantParseError_InvalidCharacter           VariantParseError = int64(5)
	VariantParseError_InvalidFormatString        VariantParseError = int64(6)
	VariantParseError_InvalidObjectPath          VariantParseError = int64(7)
	VariantParseError_InvalidSignature           VariantParseError = int64(8)
	VariantParseError_InvalidTypeString          VariantParseError = int64(9)
	VariantParseError_NoCommonType               VariantParseError = int64(10)
	VariantParseError_NumberOutOfRange           VariantParseError = int64(11)
	VariantParseError_NumberTooBig               VariantParseError = int64(12)
	VariantParseError_TypeError                  VariantParseError = int64(13)
	VariantParseError_UnexpectedToken            VariantParseError = int64(14)
	VariantParseError_UnknownKeyword             VariantParseError = int64(15)
	VariantParseError_UnterminatedStringConstant VariantParseError = int64(16)
	VariantParseError_ValueExpected              VariantParseError = int64(17)
)

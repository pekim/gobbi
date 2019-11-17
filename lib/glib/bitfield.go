// Code generated - DO NOT EDIT.

package glib

// AsciiType is a representation of the C type GAsciiType.
type AsciiType int

const (
	// AsciiType_Alnum is a representation of the C type G_ASCII_ALNUM.
	AsciiType_Alnum AsciiType = 1
	// AsciiType_Alpha is a representation of the C type G_ASCII_ALPHA.
	AsciiType_Alpha AsciiType = 2
	// AsciiType_Cntrl is a representation of the C type G_ASCII_CNTRL.
	AsciiType_Cntrl AsciiType = 4
	// AsciiType_Digit is a representation of the C type G_ASCII_DIGIT.
	AsciiType_Digit AsciiType = 8
	// AsciiType_Graph is a representation of the C type G_ASCII_GRAPH.
	AsciiType_Graph AsciiType = 16
	// AsciiType_Lower is a representation of the C type G_ASCII_LOWER.
	AsciiType_Lower AsciiType = 32
	// AsciiType_Print is a representation of the C type G_ASCII_PRINT.
	AsciiType_Print AsciiType = 64
	// AsciiType_Punct is a representation of the C type G_ASCII_PUNCT.
	AsciiType_Punct AsciiType = 128
	// AsciiType_Space is a representation of the C type G_ASCII_SPACE.
	AsciiType_Space AsciiType = 256
	// AsciiType_Upper is a representation of the C type G_ASCII_UPPER.
	AsciiType_Upper AsciiType = 512
	// AsciiType_Xdigit is a representation of the C type G_ASCII_XDIGIT.
	AsciiType_Xdigit AsciiType = 1024
)

// FileTest is a representation of the C type GFileTest.
type FileTest int

const (
	// FileTest_IsRegular is a representation of the C type G_FILE_TEST_IS_REGULAR.
	FileTest_IsRegular FileTest = 1
	// FileTest_IsSymlink is a representation of the C type G_FILE_TEST_IS_SYMLINK.
	FileTest_IsSymlink FileTest = 2
	// FileTest_IsDir is a representation of the C type G_FILE_TEST_IS_DIR.
	FileTest_IsDir FileTest = 4
	// FileTest_IsExecutable is a representation of the C type G_FILE_TEST_IS_EXECUTABLE.
	FileTest_IsExecutable FileTest = 8
	// FileTest_Exists is a representation of the C type G_FILE_TEST_EXISTS.
	FileTest_Exists FileTest = 16
)

// FormatSizeFlags is a representation of the C type GFormatSizeFlags.
type FormatSizeFlags int

const (
	// FormatSizeFlags_Default is a representation of the C type G_FORMAT_SIZE_DEFAULT.
	FormatSizeFlags_Default FormatSizeFlags = 0
	// FormatSizeFlags_LongFormat is a representation of the C type G_FORMAT_SIZE_LONG_FORMAT.
	FormatSizeFlags_LongFormat FormatSizeFlags = 1
	// FormatSizeFlags_IecUnits is a representation of the C type G_FORMAT_SIZE_IEC_UNITS.
	FormatSizeFlags_IecUnits FormatSizeFlags = 2
	// FormatSizeFlags_Bits is a representation of the C type G_FORMAT_SIZE_BITS.
	FormatSizeFlags_Bits FormatSizeFlags = 4
)

// HookFlagMask is a representation of the C type GHookFlagMask.
type HookFlagMask int

const (
	// HookFlagMask_Active is a representation of the C type G_HOOK_FLAG_ACTIVE.
	HookFlagMask_Active HookFlagMask = 1
	// HookFlagMask_InCall is a representation of the C type G_HOOK_FLAG_IN_CALL.
	HookFlagMask_InCall HookFlagMask = 2
	// HookFlagMask_Mask is a representation of the C type G_HOOK_FLAG_MASK.
	HookFlagMask_Mask HookFlagMask = 15
)

// IOCondition is a representation of the C type GIOCondition.
type IOCondition int

const (
	// IOCondition_In is a representation of the C type G_IO_IN.
	IOCondition_In IOCondition = 1
	// IOCondition_Out is a representation of the C type G_IO_OUT.
	IOCondition_Out IOCondition = 4
	// IOCondition_Pri is a representation of the C type G_IO_PRI.
	IOCondition_Pri IOCondition = 2
	// IOCondition_Err is a representation of the C type G_IO_ERR.
	IOCondition_Err IOCondition = 8
	// IOCondition_Hup is a representation of the C type G_IO_HUP.
	IOCondition_Hup IOCondition = 16
	// IOCondition_Nval is a representation of the C type G_IO_NVAL.
	IOCondition_Nval IOCondition = 32
)

// IOFlags is a representation of the C type GIOFlags.
type IOFlags int

const (
	// IOFlags_Append is a representation of the C type G_IO_FLAG_APPEND.
	IOFlags_Append IOFlags = 1
	// IOFlags_Nonblock is a representation of the C type G_IO_FLAG_NONBLOCK.
	IOFlags_Nonblock IOFlags = 2
	// IOFlags_IsReadable is a representation of the C type G_IO_FLAG_IS_READABLE.
	IOFlags_IsReadable IOFlags = 4
	// IOFlags_IsWritable is a representation of the C type G_IO_FLAG_IS_WRITABLE.
	IOFlags_IsWritable IOFlags = 8
	// IOFlags_IsWriteable is a representation of the C type G_IO_FLAG_IS_WRITEABLE.
	IOFlags_IsWriteable IOFlags = 8
	// IOFlags_IsSeekable is a representation of the C type G_IO_FLAG_IS_SEEKABLE.
	IOFlags_IsSeekable IOFlags = 16
	// IOFlags_Mask is a representation of the C type G_IO_FLAG_MASK.
	IOFlags_Mask IOFlags = 31
	// IOFlags_GetMask is a representation of the C type G_IO_FLAG_GET_MASK.
	IOFlags_GetMask IOFlags = 31
	// IOFlags_SetMask is a representation of the C type G_IO_FLAG_SET_MASK.
	IOFlags_SetMask IOFlags = 3
)

// KeyFileFlags is a representation of the C type GKeyFileFlags.
type KeyFileFlags int

const (
	// KeyFileFlags_None is a representation of the C type G_KEY_FILE_NONE.
	KeyFileFlags_None KeyFileFlags = 0
	// KeyFileFlags_KeepComments is a representation of the C type G_KEY_FILE_KEEP_COMMENTS.
	KeyFileFlags_KeepComments KeyFileFlags = 1
	// KeyFileFlags_KeepTranslations is a representation of the C type G_KEY_FILE_KEEP_TRANSLATIONS.
	KeyFileFlags_KeepTranslations KeyFileFlags = 2
)

// LogLevelFlags is a representation of the C type GLogLevelFlags.
type LogLevelFlags int

const (
	// LogLevelFlags_FlagRecursion is a representation of the C type G_LOG_FLAG_RECURSION.
	LogLevelFlags_FlagRecursion LogLevelFlags = 1
	// LogLevelFlags_FlagFatal is a representation of the C type G_LOG_FLAG_FATAL.
	LogLevelFlags_FlagFatal LogLevelFlags = 2
	// LogLevelFlags_LevelError is a representation of the C type G_LOG_LEVEL_ERROR.
	LogLevelFlags_LevelError LogLevelFlags = 4
	// LogLevelFlags_LevelCritical is a representation of the C type G_LOG_LEVEL_CRITICAL.
	LogLevelFlags_LevelCritical LogLevelFlags = 8
	// LogLevelFlags_LevelWarning is a representation of the C type G_LOG_LEVEL_WARNING.
	LogLevelFlags_LevelWarning LogLevelFlags = 16
	// LogLevelFlags_LevelMessage is a representation of the C type G_LOG_LEVEL_MESSAGE.
	LogLevelFlags_LevelMessage LogLevelFlags = 32
	// LogLevelFlags_LevelInfo is a representation of the C type G_LOG_LEVEL_INFO.
	LogLevelFlags_LevelInfo LogLevelFlags = 64
	// LogLevelFlags_LevelDebug is a representation of the C type G_LOG_LEVEL_DEBUG.
	LogLevelFlags_LevelDebug LogLevelFlags = 128
	// LogLevelFlags_LevelMask is a representation of the C type G_LOG_LEVEL_MASK.
	LogLevelFlags_LevelMask LogLevelFlags = -4
)

// MarkupCollectType is a representation of the C type GMarkupCollectType.
type MarkupCollectType int

const (
	// MarkupCollectType_Invalid is a representation of the C type G_MARKUP_COLLECT_INVALID.
	MarkupCollectType_Invalid MarkupCollectType = 0
	// MarkupCollectType_String is a representation of the C type G_MARKUP_COLLECT_STRING.
	MarkupCollectType_String MarkupCollectType = 1
	// MarkupCollectType_Strdup is a representation of the C type G_MARKUP_COLLECT_STRDUP.
	MarkupCollectType_Strdup MarkupCollectType = 2
	// MarkupCollectType_Boolean is a representation of the C type G_MARKUP_COLLECT_BOOLEAN.
	MarkupCollectType_Boolean MarkupCollectType = 3
	// MarkupCollectType_Tristate is a representation of the C type G_MARKUP_COLLECT_TRISTATE.
	MarkupCollectType_Tristate MarkupCollectType = 4
	// MarkupCollectType_Optional is a representation of the C type G_MARKUP_COLLECT_OPTIONAL.
	MarkupCollectType_Optional MarkupCollectType = 65536
)

// MarkupParseFlags is a representation of the C type GMarkupParseFlags.
type MarkupParseFlags int

const (
	// MarkupParseFlags_DoNotUseThisUnsupportedFlag is a representation of the C type G_MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG.
	MarkupParseFlags_DoNotUseThisUnsupportedFlag MarkupParseFlags = 1
	// MarkupParseFlags_TreatCdataAsText is a representation of the C type G_MARKUP_TREAT_CDATA_AS_TEXT.
	MarkupParseFlags_TreatCdataAsText MarkupParseFlags = 2
	// MarkupParseFlags_PrefixErrorPosition is a representation of the C type G_MARKUP_PREFIX_ERROR_POSITION.
	MarkupParseFlags_PrefixErrorPosition MarkupParseFlags = 4
	// MarkupParseFlags_IgnoreQualified is a representation of the C type G_MARKUP_IGNORE_QUALIFIED.
	MarkupParseFlags_IgnoreQualified MarkupParseFlags = 8
)

// OptionFlags is a representation of the C type GOptionFlags.
type OptionFlags int

const (
	// OptionFlags_None is a representation of the C type G_OPTION_FLAG_NONE.
	OptionFlags_None OptionFlags = 0
	// OptionFlags_Hidden is a representation of the C type G_OPTION_FLAG_HIDDEN.
	OptionFlags_Hidden OptionFlags = 1
	// OptionFlags_InMain is a representation of the C type G_OPTION_FLAG_IN_MAIN.
	OptionFlags_InMain OptionFlags = 2
	// OptionFlags_Reverse is a representation of the C type G_OPTION_FLAG_REVERSE.
	OptionFlags_Reverse OptionFlags = 4
	// OptionFlags_NoArg is a representation of the C type G_OPTION_FLAG_NO_ARG.
	OptionFlags_NoArg OptionFlags = 8
	// OptionFlags_Filename is a representation of the C type G_OPTION_FLAG_FILENAME.
	OptionFlags_Filename OptionFlags = 16
	// OptionFlags_OptionalArg is a representation of the C type G_OPTION_FLAG_OPTIONAL_ARG.
	OptionFlags_OptionalArg OptionFlags = 32
	// OptionFlags_Noalias is a representation of the C type G_OPTION_FLAG_NOALIAS.
	OptionFlags_Noalias OptionFlags = 64
)

// RegexCompileFlags is a representation of the C type GRegexCompileFlags.
//
// since 2.14
type RegexCompileFlags int

const (
	// RegexCompileFlags_Caseless is a representation of the C type G_REGEX_CASELESS.
	RegexCompileFlags_Caseless RegexCompileFlags = 1
	// RegexCompileFlags_Multiline is a representation of the C type G_REGEX_MULTILINE.
	RegexCompileFlags_Multiline RegexCompileFlags = 2
	// RegexCompileFlags_Dotall is a representation of the C type G_REGEX_DOTALL.
	RegexCompileFlags_Dotall RegexCompileFlags = 4
	// RegexCompileFlags_Extended is a representation of the C type G_REGEX_EXTENDED.
	RegexCompileFlags_Extended RegexCompileFlags = 8
	// RegexCompileFlags_Anchored is a representation of the C type G_REGEX_ANCHORED.
	RegexCompileFlags_Anchored RegexCompileFlags = 16
	// RegexCompileFlags_DollarEndonly is a representation of the C type G_REGEX_DOLLAR_ENDONLY.
	RegexCompileFlags_DollarEndonly RegexCompileFlags = 32
	// RegexCompileFlags_Ungreedy is a representation of the C type G_REGEX_UNGREEDY.
	RegexCompileFlags_Ungreedy RegexCompileFlags = 512
	// RegexCompileFlags_Raw is a representation of the C type G_REGEX_RAW.
	RegexCompileFlags_Raw RegexCompileFlags = 2048
	// RegexCompileFlags_NoAutoCapture is a representation of the C type G_REGEX_NO_AUTO_CAPTURE.
	RegexCompileFlags_NoAutoCapture RegexCompileFlags = 4096
	// RegexCompileFlags_Optimize is a representation of the C type G_REGEX_OPTIMIZE.
	RegexCompileFlags_Optimize RegexCompileFlags = 8192
	// RegexCompileFlags_Firstline is a representation of the C type G_REGEX_FIRSTLINE.
	RegexCompileFlags_Firstline RegexCompileFlags = 262144
	// RegexCompileFlags_Dupnames is a representation of the C type G_REGEX_DUPNAMES.
	RegexCompileFlags_Dupnames RegexCompileFlags = 524288
	// RegexCompileFlags_NewlineCr is a representation of the C type G_REGEX_NEWLINE_CR.
	RegexCompileFlags_NewlineCr RegexCompileFlags = 1048576
	// RegexCompileFlags_NewlineLf is a representation of the C type G_REGEX_NEWLINE_LF.
	RegexCompileFlags_NewlineLf RegexCompileFlags = 2097152
	// RegexCompileFlags_NewlineCrlf is a representation of the C type G_REGEX_NEWLINE_CRLF.
	RegexCompileFlags_NewlineCrlf RegexCompileFlags = 3145728
	// RegexCompileFlags_NewlineAnycrlf is a representation of the C type G_REGEX_NEWLINE_ANYCRLF.
	RegexCompileFlags_NewlineAnycrlf RegexCompileFlags = 5242880
	// RegexCompileFlags_BsrAnycrlf is a representation of the C type G_REGEX_BSR_ANYCRLF.
	RegexCompileFlags_BsrAnycrlf RegexCompileFlags = 8388608
	// RegexCompileFlags_JavascriptCompat is a representation of the C type G_REGEX_JAVASCRIPT_COMPAT.
	RegexCompileFlags_JavascriptCompat RegexCompileFlags = 33554432
)

// RegexMatchFlags is a representation of the C type GRegexMatchFlags.
//
// since 2.14
type RegexMatchFlags int

const (
	// RegexMatchFlags_Anchored is a representation of the C type G_REGEX_MATCH_ANCHORED.
	RegexMatchFlags_Anchored RegexMatchFlags = 16
	// RegexMatchFlags_Notbol is a representation of the C type G_REGEX_MATCH_NOTBOL.
	RegexMatchFlags_Notbol RegexMatchFlags = 128
	// RegexMatchFlags_Noteol is a representation of the C type G_REGEX_MATCH_NOTEOL.
	RegexMatchFlags_Noteol RegexMatchFlags = 256
	// RegexMatchFlags_Notempty is a representation of the C type G_REGEX_MATCH_NOTEMPTY.
	RegexMatchFlags_Notempty RegexMatchFlags = 1024
	// RegexMatchFlags_Partial is a representation of the C type G_REGEX_MATCH_PARTIAL.
	RegexMatchFlags_Partial RegexMatchFlags = 32768
	// RegexMatchFlags_NewlineCr is a representation of the C type G_REGEX_MATCH_NEWLINE_CR.
	RegexMatchFlags_NewlineCr RegexMatchFlags = 1048576
	// RegexMatchFlags_NewlineLf is a representation of the C type G_REGEX_MATCH_NEWLINE_LF.
	RegexMatchFlags_NewlineLf RegexMatchFlags = 2097152
	// RegexMatchFlags_NewlineCrlf is a representation of the C type G_REGEX_MATCH_NEWLINE_CRLF.
	RegexMatchFlags_NewlineCrlf RegexMatchFlags = 3145728
	// RegexMatchFlags_NewlineAny is a representation of the C type G_REGEX_MATCH_NEWLINE_ANY.
	RegexMatchFlags_NewlineAny RegexMatchFlags = 4194304
	// RegexMatchFlags_NewlineAnycrlf is a representation of the C type G_REGEX_MATCH_NEWLINE_ANYCRLF.
	RegexMatchFlags_NewlineAnycrlf RegexMatchFlags = 5242880
	// RegexMatchFlags_BsrAnycrlf is a representation of the C type G_REGEX_MATCH_BSR_ANYCRLF.
	RegexMatchFlags_BsrAnycrlf RegexMatchFlags = 8388608
	// RegexMatchFlags_BsrAny is a representation of the C type G_REGEX_MATCH_BSR_ANY.
	RegexMatchFlags_BsrAny RegexMatchFlags = 16777216
	// RegexMatchFlags_PartialSoft is a representation of the C type G_REGEX_MATCH_PARTIAL_SOFT.
	RegexMatchFlags_PartialSoft RegexMatchFlags = 32768
	// RegexMatchFlags_PartialHard is a representation of the C type G_REGEX_MATCH_PARTIAL_HARD.
	RegexMatchFlags_PartialHard RegexMatchFlags = 134217728
	// RegexMatchFlags_NotemptyAtstart is a representation of the C type G_REGEX_MATCH_NOTEMPTY_ATSTART.
	RegexMatchFlags_NotemptyAtstart RegexMatchFlags = 268435456
)

// SpawnFlags is a representation of the C type GSpawnFlags.
type SpawnFlags int

const (
	// SpawnFlags_Default is a representation of the C type G_SPAWN_DEFAULT.
	SpawnFlags_Default SpawnFlags = 0
	// SpawnFlags_LeaveDescriptorsOpen is a representation of the C type G_SPAWN_LEAVE_DESCRIPTORS_OPEN.
	SpawnFlags_LeaveDescriptorsOpen SpawnFlags = 1
	// SpawnFlags_DoNotReapChild is a representation of the C type G_SPAWN_DO_NOT_REAP_CHILD.
	SpawnFlags_DoNotReapChild SpawnFlags = 2
	// SpawnFlags_SearchPath is a representation of the C type G_SPAWN_SEARCH_PATH.
	SpawnFlags_SearchPath SpawnFlags = 4
	// SpawnFlags_StdoutToDevNull is a representation of the C type G_SPAWN_STDOUT_TO_DEV_NULL.
	SpawnFlags_StdoutToDevNull SpawnFlags = 8
	// SpawnFlags_StderrToDevNull is a representation of the C type G_SPAWN_STDERR_TO_DEV_NULL.
	SpawnFlags_StderrToDevNull SpawnFlags = 16
	// SpawnFlags_ChildInheritsStdin is a representation of the C type G_SPAWN_CHILD_INHERITS_STDIN.
	SpawnFlags_ChildInheritsStdin SpawnFlags = 32
	// SpawnFlags_FileAndArgvZero is a representation of the C type G_SPAWN_FILE_AND_ARGV_ZERO.
	SpawnFlags_FileAndArgvZero SpawnFlags = 64
	// SpawnFlags_SearchPathFromEnvp is a representation of the C type G_SPAWN_SEARCH_PATH_FROM_ENVP.
	SpawnFlags_SearchPathFromEnvp SpawnFlags = 128
	// SpawnFlags_CloexecPipes is a representation of the C type G_SPAWN_CLOEXEC_PIPES.
	SpawnFlags_CloexecPipes SpawnFlags = 256
)

// TestSubprocessFlags is a representation of the C type GTestSubprocessFlags.
type TestSubprocessFlags int

const (
	// TestSubprocessFlags_Stdin is a representation of the C type G_TEST_SUBPROCESS_INHERIT_STDIN.
	TestSubprocessFlags_Stdin TestSubprocessFlags = 1
	// TestSubprocessFlags_Stdout is a representation of the C type G_TEST_SUBPROCESS_INHERIT_STDOUT.
	TestSubprocessFlags_Stdout TestSubprocessFlags = 2
	// TestSubprocessFlags_Stderr is a representation of the C type G_TEST_SUBPROCESS_INHERIT_STDERR.
	TestSubprocessFlags_Stderr TestSubprocessFlags = 4
)

// TestTrapFlags is a representation of the C type GTestTrapFlags.
type TestTrapFlags int

const (
	// TestTrapFlags_SilenceStdout is a representation of the C type G_TEST_TRAP_SILENCE_STDOUT.
	TestTrapFlags_SilenceStdout TestTrapFlags = 128
	// TestTrapFlags_SilenceStderr is a representation of the C type G_TEST_TRAP_SILENCE_STDERR.
	TestTrapFlags_SilenceStderr TestTrapFlags = 256
	// TestTrapFlags_InheritStdin is a representation of the C type G_TEST_TRAP_INHERIT_STDIN.
	TestTrapFlags_InheritStdin TestTrapFlags = 512
)

// TraverseFlags is a representation of the C type GTraverseFlags.
type TraverseFlags int

const (
	// TraverseFlags_Leaves is a representation of the C type G_TRAVERSE_LEAVES.
	TraverseFlags_Leaves TraverseFlags = 1
	// TraverseFlags_NonLeaves is a representation of the C type G_TRAVERSE_NON_LEAVES.
	TraverseFlags_NonLeaves TraverseFlags = 2
	// TraverseFlags_All is a representation of the C type G_TRAVERSE_ALL.
	TraverseFlags_All TraverseFlags = 3
	// TraverseFlags_Mask is a representation of the C type G_TRAVERSE_MASK.
	TraverseFlags_Mask TraverseFlags = 3
	// TraverseFlags_Leafs is a representation of the C type G_TRAVERSE_LEAFS.
	TraverseFlags_Leafs TraverseFlags = 1
	// TraverseFlags_NonLeafs is a representation of the C type G_TRAVERSE_NON_LEAFS.
	TraverseFlags_NonLeafs TraverseFlags = 2
)

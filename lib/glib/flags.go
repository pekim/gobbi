// Code generated - DO NOT EDIT.

package glib

type AsciiType uint32

const (
	AsciiType_Alnum  AsciiType = int64(1)
	AsciiType_Alpha  AsciiType = int64(2)
	AsciiType_Cntrl  AsciiType = int64(4)
	AsciiType_Digit  AsciiType = int64(8)
	AsciiType_Graph  AsciiType = int64(16)
	AsciiType_Lower  AsciiType = int64(32)
	AsciiType_Print  AsciiType = int64(64)
	AsciiType_Punct  AsciiType = int64(128)
	AsciiType_Space  AsciiType = int64(256)
	AsciiType_Upper  AsciiType = int64(512)
	AsciiType_Xdigit AsciiType = int64(1024)
)

type FileTest uint32

const (
	FileTest_IsRegular    FileTest = int64(1)
	FileTest_IsSymlink    FileTest = int64(2)
	FileTest_IsDir        FileTest = int64(4)
	FileTest_IsExecutable FileTest = int64(8)
	FileTest_Exists       FileTest = int64(16)
)

type FormatSizeFlags uint32

const (
	FormatSizeFlags_Default    FormatSizeFlags = int64(0)
	FormatSizeFlags_LongFormat FormatSizeFlags = int64(1)
	FormatSizeFlags_IecUnits   FormatSizeFlags = int64(2)
	FormatSizeFlags_Bits       FormatSizeFlags = int64(4)
)

type HookFlagMask uint32

const (
	HookFlagMask_Active HookFlagMask = int64(1)
	HookFlagMask_InCall HookFlagMask = int64(2)
	HookFlagMask_Mask   HookFlagMask = int64(15)
)

type IOCondition uint32

const (
	IOCondition_In   IOCondition = int64(1)
	IOCondition_Out  IOCondition = int64(4)
	IOCondition_Pri  IOCondition = int64(2)
	IOCondition_Err  IOCondition = int64(8)
	IOCondition_Hup  IOCondition = int64(16)
	IOCondition_Nval IOCondition = int64(32)
)

type IOFlags uint32

const (
	IOFlags_Append      IOFlags = int64(1)
	IOFlags_Nonblock    IOFlags = int64(2)
	IOFlags_IsReadable  IOFlags = int64(4)
	IOFlags_IsWritable  IOFlags = int64(8)
	IOFlags_IsWriteable IOFlags = int64(8)
	IOFlags_IsSeekable  IOFlags = int64(16)
	IOFlags_Mask        IOFlags = int64(31)
	IOFlags_GetMask     IOFlags = int64(31)
	IOFlags_SetMask     IOFlags = int64(3)
)

type KeyFileFlags uint32

const (
	KeyFileFlags_None             KeyFileFlags = int64(0)
	KeyFileFlags_KeepComments     KeyFileFlags = int64(1)
	KeyFileFlags_KeepTranslations KeyFileFlags = int64(2)
)

type LogLevelFlags int32

const (
	LogLevelFlags_FlagRecursion LogLevelFlags = int64(1)
	LogLevelFlags_FlagFatal     LogLevelFlags = int64(2)
	LogLevelFlags_LevelError    LogLevelFlags = int64(4)
	LogLevelFlags_LevelCritical LogLevelFlags = int64(8)
	LogLevelFlags_LevelWarning  LogLevelFlags = int64(16)
	LogLevelFlags_LevelMessage  LogLevelFlags = int64(32)
	LogLevelFlags_LevelInfo     LogLevelFlags = int64(64)
	LogLevelFlags_LevelDebug    LogLevelFlags = int64(128)
	LogLevelFlags_LevelMask     LogLevelFlags = int64(-4)
)

type MarkupCollectType uint32

const (
	MarkupCollectType_Invalid  MarkupCollectType = int64(0)
	MarkupCollectType_String   MarkupCollectType = int64(1)
	MarkupCollectType_Strdup   MarkupCollectType = int64(2)
	MarkupCollectType_Boolean  MarkupCollectType = int64(3)
	MarkupCollectType_Tristate MarkupCollectType = int64(4)
	MarkupCollectType_Optional MarkupCollectType = int64(65536)
)

type MarkupParseFlags uint32

const (
	MarkupParseFlags_DoNotUseThisUnsupportedFlag MarkupParseFlags = int64(1)
	MarkupParseFlags_TreatCdataAsText            MarkupParseFlags = int64(2)
	MarkupParseFlags_PrefixErrorPosition         MarkupParseFlags = int64(4)
	MarkupParseFlags_IgnoreQualified             MarkupParseFlags = int64(8)
)

type OptionFlags uint32

const (
	OptionFlags_None        OptionFlags = int64(0)
	OptionFlags_Hidden      OptionFlags = int64(1)
	OptionFlags_InMain      OptionFlags = int64(2)
	OptionFlags_Reverse     OptionFlags = int64(4)
	OptionFlags_NoArg       OptionFlags = int64(8)
	OptionFlags_Filename    OptionFlags = int64(16)
	OptionFlags_OptionalArg OptionFlags = int64(32)
	OptionFlags_Noalias     OptionFlags = int64(64)
)

type RegexCompileFlags uint32

const (
	RegexCompileFlags_Caseless         RegexCompileFlags = int64(1)
	RegexCompileFlags_Multiline        RegexCompileFlags = int64(2)
	RegexCompileFlags_Dotall           RegexCompileFlags = int64(4)
	RegexCompileFlags_Extended         RegexCompileFlags = int64(8)
	RegexCompileFlags_Anchored         RegexCompileFlags = int64(16)
	RegexCompileFlags_DollarEndonly    RegexCompileFlags = int64(32)
	RegexCompileFlags_Ungreedy         RegexCompileFlags = int64(512)
	RegexCompileFlags_Raw              RegexCompileFlags = int64(2048)
	RegexCompileFlags_NoAutoCapture    RegexCompileFlags = int64(4096)
	RegexCompileFlags_Optimize         RegexCompileFlags = int64(8192)
	RegexCompileFlags_Firstline        RegexCompileFlags = int64(262144)
	RegexCompileFlags_Dupnames         RegexCompileFlags = int64(524288)
	RegexCompileFlags_NewlineCr        RegexCompileFlags = int64(1048576)
	RegexCompileFlags_NewlineLf        RegexCompileFlags = int64(2097152)
	RegexCompileFlags_NewlineCrlf      RegexCompileFlags = int64(3145728)
	RegexCompileFlags_NewlineAnycrlf   RegexCompileFlags = int64(5242880)
	RegexCompileFlags_BsrAnycrlf       RegexCompileFlags = int64(8388608)
	RegexCompileFlags_JavascriptCompat RegexCompileFlags = int64(33554432)
)

type RegexMatchFlags uint32

const (
	RegexMatchFlags_Anchored        RegexMatchFlags = int64(16)
	RegexMatchFlags_Notbol          RegexMatchFlags = int64(128)
	RegexMatchFlags_Noteol          RegexMatchFlags = int64(256)
	RegexMatchFlags_Notempty        RegexMatchFlags = int64(1024)
	RegexMatchFlags_Partial         RegexMatchFlags = int64(32768)
	RegexMatchFlags_NewlineCr       RegexMatchFlags = int64(1048576)
	RegexMatchFlags_NewlineLf       RegexMatchFlags = int64(2097152)
	RegexMatchFlags_NewlineCrlf     RegexMatchFlags = int64(3145728)
	RegexMatchFlags_NewlineAny      RegexMatchFlags = int64(4194304)
	RegexMatchFlags_NewlineAnycrlf  RegexMatchFlags = int64(5242880)
	RegexMatchFlags_BsrAnycrlf      RegexMatchFlags = int64(8388608)
	RegexMatchFlags_BsrAny          RegexMatchFlags = int64(16777216)
	RegexMatchFlags_PartialSoft     RegexMatchFlags = int64(32768)
	RegexMatchFlags_PartialHard     RegexMatchFlags = int64(134217728)
	RegexMatchFlags_NotemptyAtstart RegexMatchFlags = int64(268435456)
)

type SpawnFlags uint32

const (
	SpawnFlags_Default              SpawnFlags = int64(0)
	SpawnFlags_LeaveDescriptorsOpen SpawnFlags = int64(1)
	SpawnFlags_DoNotReapChild       SpawnFlags = int64(2)
	SpawnFlags_SearchPath           SpawnFlags = int64(4)
	SpawnFlags_StdoutToDevNull      SpawnFlags = int64(8)
	SpawnFlags_StderrToDevNull      SpawnFlags = int64(16)
	SpawnFlags_ChildInheritsStdin   SpawnFlags = int64(32)
	SpawnFlags_FileAndArgvZero      SpawnFlags = int64(64)
	SpawnFlags_SearchPathFromEnvp   SpawnFlags = int64(128)
	SpawnFlags_CloexecPipes         SpawnFlags = int64(256)
)

type TestSubprocessFlags uint32

const (
	TestSubprocessFlags_Stdin  TestSubprocessFlags = int64(1)
	TestSubprocessFlags_Stdout TestSubprocessFlags = int64(2)
	TestSubprocessFlags_Stderr TestSubprocessFlags = int64(4)
)

type TestTrapFlags uint32

const (
	TestTrapFlags_SilenceStdout TestTrapFlags = int64(128)
	TestTrapFlags_SilenceStderr TestTrapFlags = int64(256)
	TestTrapFlags_InheritStdin  TestTrapFlags = int64(512)
)

type TraverseFlags uint32

const (
	TraverseFlags_Leaves    TraverseFlags = int64(1)
	TraverseFlags_NonLeaves TraverseFlags = int64(2)
	TraverseFlags_All       TraverseFlags = int64(3)
	TraverseFlags_Mask      TraverseFlags = int64(3)
	TraverseFlags_Leafs     TraverseFlags = int64(1)
	TraverseFlags_NonLeafs  TraverseFlags = int64(2)
)

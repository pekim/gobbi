// Code generated - DO NOT EDIT.

package glib

// Asciitype is a representation of the C type GAsciiType.
type Asciitype int

const (
	// Asciitype_Alnum is a representation of the C type G_ASCII_ALNUM.
	Asciitype_Alnum Asciitype = 1
	// Asciitype_Alpha is a representation of the C type G_ASCII_ALPHA.
	Asciitype_Alpha Asciitype = 2
	// Asciitype_Cntrl is a representation of the C type G_ASCII_CNTRL.
	Asciitype_Cntrl Asciitype = 4
	// Asciitype_Digit is a representation of the C type G_ASCII_DIGIT.
	Asciitype_Digit Asciitype = 8
	// Asciitype_Graph is a representation of the C type G_ASCII_GRAPH.
	Asciitype_Graph Asciitype = 16
	// Asciitype_Lower is a representation of the C type G_ASCII_LOWER.
	Asciitype_Lower Asciitype = 32
	// Asciitype_Print is a representation of the C type G_ASCII_PRINT.
	Asciitype_Print Asciitype = 64
	// Asciitype_Punct is a representation of the C type G_ASCII_PUNCT.
	Asciitype_Punct Asciitype = 128
	// Asciitype_Space is a representation of the C type G_ASCII_SPACE.
	Asciitype_Space Asciitype = 256
	// Asciitype_Upper is a representation of the C type G_ASCII_UPPER.
	Asciitype_Upper Asciitype = 512
	// Asciitype_Xdigit is a representation of the C type G_ASCII_XDIGIT.
	Asciitype_Xdigit Asciitype = 1024
)

// Filetest is a representation of the C type GFileTest.
type Filetest int

const (
	// Filetest_IsRegular is a representation of the C type G_FILE_TEST_IS_REGULAR.
	Filetest_IsRegular Filetest = 1
	// Filetest_IsSymlink is a representation of the C type G_FILE_TEST_IS_SYMLINK.
	Filetest_IsSymlink Filetest = 2
	// Filetest_IsDir is a representation of the C type G_FILE_TEST_IS_DIR.
	Filetest_IsDir Filetest = 4
	// Filetest_IsExecutable is a representation of the C type G_FILE_TEST_IS_EXECUTABLE.
	Filetest_IsExecutable Filetest = 8
	// Filetest_Exists is a representation of the C type G_FILE_TEST_EXISTS.
	Filetest_Exists Filetest = 16
)

// Formatsizeflags is a representation of the C type GFormatSizeFlags.
type Formatsizeflags int

const (
	// Formatsizeflags_Default is a representation of the C type G_FORMAT_SIZE_DEFAULT.
	Formatsizeflags_Default Formatsizeflags = 0
	// Formatsizeflags_LongFormat is a representation of the C type G_FORMAT_SIZE_LONG_FORMAT.
	Formatsizeflags_LongFormat Formatsizeflags = 1
	// Formatsizeflags_IecUnits is a representation of the C type G_FORMAT_SIZE_IEC_UNITS.
	Formatsizeflags_IecUnits Formatsizeflags = 2
	// Formatsizeflags_Bits is a representation of the C type G_FORMAT_SIZE_BITS.
	Formatsizeflags_Bits Formatsizeflags = 4
)

// Hookflagmask is a representation of the C type GHookFlagMask.
type Hookflagmask int

const (
	// Hookflagmask_Active is a representation of the C type G_HOOK_FLAG_ACTIVE.
	Hookflagmask_Active Hookflagmask = 1
	// Hookflagmask_InCall is a representation of the C type G_HOOK_FLAG_IN_CALL.
	Hookflagmask_InCall Hookflagmask = 2
	// Hookflagmask_Mask is a representation of the C type G_HOOK_FLAG_MASK.
	Hookflagmask_Mask Hookflagmask = 15
)

// Iocondition is a representation of the C type GIOCondition.
type Iocondition int

const (
	// Iocondition_In is a representation of the C type G_IO_IN.
	Iocondition_In Iocondition = 1
	// Iocondition_Out is a representation of the C type G_IO_OUT.
	Iocondition_Out Iocondition = 4
	// Iocondition_Pri is a representation of the C type G_IO_PRI.
	Iocondition_Pri Iocondition = 2
	// Iocondition_Err is a representation of the C type G_IO_ERR.
	Iocondition_Err Iocondition = 8
	// Iocondition_Hup is a representation of the C type G_IO_HUP.
	Iocondition_Hup Iocondition = 16
	// Iocondition_Nval is a representation of the C type G_IO_NVAL.
	Iocondition_Nval Iocondition = 32
)

// Ioflags is a representation of the C type GIOFlags.
type Ioflags int

const (
	// Ioflags_Append is a representation of the C type G_IO_FLAG_APPEND.
	Ioflags_Append Ioflags = 1
	// Ioflags_Nonblock is a representation of the C type G_IO_FLAG_NONBLOCK.
	Ioflags_Nonblock Ioflags = 2
	// Ioflags_IsReadable is a representation of the C type G_IO_FLAG_IS_READABLE.
	Ioflags_IsReadable Ioflags = 4
	// Ioflags_IsWritable is a representation of the C type G_IO_FLAG_IS_WRITABLE.
	Ioflags_IsWritable Ioflags = 8
	// Ioflags_IsWriteable is a representation of the C type G_IO_FLAG_IS_WRITEABLE.
	Ioflags_IsWriteable Ioflags = 8
	// Ioflags_IsSeekable is a representation of the C type G_IO_FLAG_IS_SEEKABLE.
	Ioflags_IsSeekable Ioflags = 16
	// Ioflags_Mask is a representation of the C type G_IO_FLAG_MASK.
	Ioflags_Mask Ioflags = 31
	// Ioflags_GetMask is a representation of the C type G_IO_FLAG_GET_MASK.
	Ioflags_GetMask Ioflags = 31
	// Ioflags_SetMask is a representation of the C type G_IO_FLAG_SET_MASK.
	Ioflags_SetMask Ioflags = 3
)

// Keyfileflags is a representation of the C type GKeyFileFlags.
type Keyfileflags int

const (
	// Keyfileflags_None is a representation of the C type G_KEY_FILE_NONE.
	Keyfileflags_None Keyfileflags = 0
	// Keyfileflags_KeepComments is a representation of the C type G_KEY_FILE_KEEP_COMMENTS.
	Keyfileflags_KeepComments Keyfileflags = 1
	// Keyfileflags_KeepTranslations is a representation of the C type G_KEY_FILE_KEEP_TRANSLATIONS.
	Keyfileflags_KeepTranslations Keyfileflags = 2
)

// Loglevelflags is a representation of the C type GLogLevelFlags.
type Loglevelflags int

const (
	// Loglevelflags_FlagRecursion is a representation of the C type G_LOG_FLAG_RECURSION.
	Loglevelflags_FlagRecursion Loglevelflags = 1
	// Loglevelflags_FlagFatal is a representation of the C type G_LOG_FLAG_FATAL.
	Loglevelflags_FlagFatal Loglevelflags = 2
	// Loglevelflags_LevelError is a representation of the C type G_LOG_LEVEL_ERROR.
	Loglevelflags_LevelError Loglevelflags = 4
	// Loglevelflags_LevelCritical is a representation of the C type G_LOG_LEVEL_CRITICAL.
	Loglevelflags_LevelCritical Loglevelflags = 8
	// Loglevelflags_LevelWarning is a representation of the C type G_LOG_LEVEL_WARNING.
	Loglevelflags_LevelWarning Loglevelflags = 16
	// Loglevelflags_LevelMessage is a representation of the C type G_LOG_LEVEL_MESSAGE.
	Loglevelflags_LevelMessage Loglevelflags = 32
	// Loglevelflags_LevelInfo is a representation of the C type G_LOG_LEVEL_INFO.
	Loglevelflags_LevelInfo Loglevelflags = 64
	// Loglevelflags_LevelDebug is a representation of the C type G_LOG_LEVEL_DEBUG.
	Loglevelflags_LevelDebug Loglevelflags = 128
	// Loglevelflags_LevelMask is a representation of the C type G_LOG_LEVEL_MASK.
	Loglevelflags_LevelMask Loglevelflags = -4
)

// Markupcollecttype is a representation of the C type GMarkupCollectType.
type Markupcollecttype int

const (
	// Markupcollecttype_Invalid is a representation of the C type G_MARKUP_COLLECT_INVALID.
	Markupcollecttype_Invalid Markupcollecttype = 0
	// Markupcollecttype_String is a representation of the C type G_MARKUP_COLLECT_STRING.
	Markupcollecttype_String Markupcollecttype = 1
	// Markupcollecttype_Strdup is a representation of the C type G_MARKUP_COLLECT_STRDUP.
	Markupcollecttype_Strdup Markupcollecttype = 2
	// Markupcollecttype_Boolean is a representation of the C type G_MARKUP_COLLECT_BOOLEAN.
	Markupcollecttype_Boolean Markupcollecttype = 3
	// Markupcollecttype_Tristate is a representation of the C type G_MARKUP_COLLECT_TRISTATE.
	Markupcollecttype_Tristate Markupcollecttype = 4
	// Markupcollecttype_Optional is a representation of the C type G_MARKUP_COLLECT_OPTIONAL.
	Markupcollecttype_Optional Markupcollecttype = 65536
)

// Markupparseflags is a representation of the C type GMarkupParseFlags.
type Markupparseflags int

const (
	// Markupparseflags_DoNotUseThisUnsupportedFlag is a representation of the C type G_MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG.
	Markupparseflags_DoNotUseThisUnsupportedFlag Markupparseflags = 1
	// Markupparseflags_TreatCdataAsText is a representation of the C type G_MARKUP_TREAT_CDATA_AS_TEXT.
	Markupparseflags_TreatCdataAsText Markupparseflags = 2
	// Markupparseflags_PrefixErrorPosition is a representation of the C type G_MARKUP_PREFIX_ERROR_POSITION.
	Markupparseflags_PrefixErrorPosition Markupparseflags = 4
	// Markupparseflags_IgnoreQualified is a representation of the C type G_MARKUP_IGNORE_QUALIFIED.
	Markupparseflags_IgnoreQualified Markupparseflags = 8
)

// Optionflags is a representation of the C type GOptionFlags.
type Optionflags int

const (
	// Optionflags_None is a representation of the C type G_OPTION_FLAG_NONE.
	Optionflags_None Optionflags = 0
	// Optionflags_Hidden is a representation of the C type G_OPTION_FLAG_HIDDEN.
	Optionflags_Hidden Optionflags = 1
	// Optionflags_InMain is a representation of the C type G_OPTION_FLAG_IN_MAIN.
	Optionflags_InMain Optionflags = 2
	// Optionflags_Reverse is a representation of the C type G_OPTION_FLAG_REVERSE.
	Optionflags_Reverse Optionflags = 4
	// Optionflags_NoArg is a representation of the C type G_OPTION_FLAG_NO_ARG.
	Optionflags_NoArg Optionflags = 8
	// Optionflags_Filename is a representation of the C type G_OPTION_FLAG_FILENAME.
	Optionflags_Filename Optionflags = 16
	// Optionflags_OptionalArg is a representation of the C type G_OPTION_FLAG_OPTIONAL_ARG.
	Optionflags_OptionalArg Optionflags = 32
	// Optionflags_Noalias is a representation of the C type G_OPTION_FLAG_NOALIAS.
	Optionflags_Noalias Optionflags = 64
)

// Regexcompileflags is a representation of the C type GRegexCompileFlags.
//
// since 2.14
type Regexcompileflags int

const (
	// Regexcompileflags_Caseless is a representation of the C type G_REGEX_CASELESS.
	Regexcompileflags_Caseless Regexcompileflags = 1
	// Regexcompileflags_Multiline is a representation of the C type G_REGEX_MULTILINE.
	Regexcompileflags_Multiline Regexcompileflags = 2
	// Regexcompileflags_Dotall is a representation of the C type G_REGEX_DOTALL.
	Regexcompileflags_Dotall Regexcompileflags = 4
	// Regexcompileflags_Extended is a representation of the C type G_REGEX_EXTENDED.
	Regexcompileflags_Extended Regexcompileflags = 8
	// Regexcompileflags_Anchored is a representation of the C type G_REGEX_ANCHORED.
	Regexcompileflags_Anchored Regexcompileflags = 16
	// Regexcompileflags_DollarEndonly is a representation of the C type G_REGEX_DOLLAR_ENDONLY.
	Regexcompileflags_DollarEndonly Regexcompileflags = 32
	// Regexcompileflags_Ungreedy is a representation of the C type G_REGEX_UNGREEDY.
	Regexcompileflags_Ungreedy Regexcompileflags = 512
	// Regexcompileflags_Raw is a representation of the C type G_REGEX_RAW.
	Regexcompileflags_Raw Regexcompileflags = 2048
	// Regexcompileflags_NoAutoCapture is a representation of the C type G_REGEX_NO_AUTO_CAPTURE.
	Regexcompileflags_NoAutoCapture Regexcompileflags = 4096
	// Regexcompileflags_Optimize is a representation of the C type G_REGEX_OPTIMIZE.
	Regexcompileflags_Optimize Regexcompileflags = 8192
	// Regexcompileflags_Firstline is a representation of the C type G_REGEX_FIRSTLINE.
	Regexcompileflags_Firstline Regexcompileflags = 262144
	// Regexcompileflags_Dupnames is a representation of the C type G_REGEX_DUPNAMES.
	Regexcompileflags_Dupnames Regexcompileflags = 524288
	// Regexcompileflags_NewlineCr is a representation of the C type G_REGEX_NEWLINE_CR.
	Regexcompileflags_NewlineCr Regexcompileflags = 1048576
	// Regexcompileflags_NewlineLf is a representation of the C type G_REGEX_NEWLINE_LF.
	Regexcompileflags_NewlineLf Regexcompileflags = 2097152
	// Regexcompileflags_NewlineCrlf is a representation of the C type G_REGEX_NEWLINE_CRLF.
	Regexcompileflags_NewlineCrlf Regexcompileflags = 3145728
	// Regexcompileflags_NewlineAnycrlf is a representation of the C type G_REGEX_NEWLINE_ANYCRLF.
	Regexcompileflags_NewlineAnycrlf Regexcompileflags = 5242880
	// Regexcompileflags_BsrAnycrlf is a representation of the C type G_REGEX_BSR_ANYCRLF.
	Regexcompileflags_BsrAnycrlf Regexcompileflags = 8388608
	// Regexcompileflags_JavascriptCompat is a representation of the C type G_REGEX_JAVASCRIPT_COMPAT.
	Regexcompileflags_JavascriptCompat Regexcompileflags = 33554432
)

// Regexmatchflags is a representation of the C type GRegexMatchFlags.
//
// since 2.14
type Regexmatchflags int

const (
	// Regexmatchflags_Anchored is a representation of the C type G_REGEX_MATCH_ANCHORED.
	Regexmatchflags_Anchored Regexmatchflags = 16
	// Regexmatchflags_Notbol is a representation of the C type G_REGEX_MATCH_NOTBOL.
	Regexmatchflags_Notbol Regexmatchflags = 128
	// Regexmatchflags_Noteol is a representation of the C type G_REGEX_MATCH_NOTEOL.
	Regexmatchflags_Noteol Regexmatchflags = 256
	// Regexmatchflags_Notempty is a representation of the C type G_REGEX_MATCH_NOTEMPTY.
	Regexmatchflags_Notempty Regexmatchflags = 1024
	// Regexmatchflags_Partial is a representation of the C type G_REGEX_MATCH_PARTIAL.
	Regexmatchflags_Partial Regexmatchflags = 32768
	// Regexmatchflags_NewlineCr is a representation of the C type G_REGEX_MATCH_NEWLINE_CR.
	Regexmatchflags_NewlineCr Regexmatchflags = 1048576
	// Regexmatchflags_NewlineLf is a representation of the C type G_REGEX_MATCH_NEWLINE_LF.
	Regexmatchflags_NewlineLf Regexmatchflags = 2097152
	// Regexmatchflags_NewlineCrlf is a representation of the C type G_REGEX_MATCH_NEWLINE_CRLF.
	Regexmatchflags_NewlineCrlf Regexmatchflags = 3145728
	// Regexmatchflags_NewlineAny is a representation of the C type G_REGEX_MATCH_NEWLINE_ANY.
	Regexmatchflags_NewlineAny Regexmatchflags = 4194304
	// Regexmatchflags_NewlineAnycrlf is a representation of the C type G_REGEX_MATCH_NEWLINE_ANYCRLF.
	Regexmatchflags_NewlineAnycrlf Regexmatchflags = 5242880
	// Regexmatchflags_BsrAnycrlf is a representation of the C type G_REGEX_MATCH_BSR_ANYCRLF.
	Regexmatchflags_BsrAnycrlf Regexmatchflags = 8388608
	// Regexmatchflags_BsrAny is a representation of the C type G_REGEX_MATCH_BSR_ANY.
	Regexmatchflags_BsrAny Regexmatchflags = 16777216
	// Regexmatchflags_PartialSoft is a representation of the C type G_REGEX_MATCH_PARTIAL_SOFT.
	Regexmatchflags_PartialSoft Regexmatchflags = 32768
	// Regexmatchflags_PartialHard is a representation of the C type G_REGEX_MATCH_PARTIAL_HARD.
	Regexmatchflags_PartialHard Regexmatchflags = 134217728
	// Regexmatchflags_NotemptyAtstart is a representation of the C type G_REGEX_MATCH_NOTEMPTY_ATSTART.
	Regexmatchflags_NotemptyAtstart Regexmatchflags = 268435456
)

// Spawnflags is a representation of the C type GSpawnFlags.
type Spawnflags int

const (
	// Spawnflags_Default is a representation of the C type G_SPAWN_DEFAULT.
	Spawnflags_Default Spawnflags = 0
	// Spawnflags_LeaveDescriptorsOpen is a representation of the C type G_SPAWN_LEAVE_DESCRIPTORS_OPEN.
	Spawnflags_LeaveDescriptorsOpen Spawnflags = 1
	// Spawnflags_DoNotReapChild is a representation of the C type G_SPAWN_DO_NOT_REAP_CHILD.
	Spawnflags_DoNotReapChild Spawnflags = 2
	// Spawnflags_SearchPath is a representation of the C type G_SPAWN_SEARCH_PATH.
	Spawnflags_SearchPath Spawnflags = 4
	// Spawnflags_StdoutToDevNull is a representation of the C type G_SPAWN_STDOUT_TO_DEV_NULL.
	Spawnflags_StdoutToDevNull Spawnflags = 8
	// Spawnflags_StderrToDevNull is a representation of the C type G_SPAWN_STDERR_TO_DEV_NULL.
	Spawnflags_StderrToDevNull Spawnflags = 16
	// Spawnflags_ChildInheritsStdin is a representation of the C type G_SPAWN_CHILD_INHERITS_STDIN.
	Spawnflags_ChildInheritsStdin Spawnflags = 32
	// Spawnflags_FileAndArgvZero is a representation of the C type G_SPAWN_FILE_AND_ARGV_ZERO.
	Spawnflags_FileAndArgvZero Spawnflags = 64
	// Spawnflags_SearchPathFromEnvp is a representation of the C type G_SPAWN_SEARCH_PATH_FROM_ENVP.
	Spawnflags_SearchPathFromEnvp Spawnflags = 128
	// Spawnflags_CloexecPipes is a representation of the C type G_SPAWN_CLOEXEC_PIPES.
	Spawnflags_CloexecPipes Spawnflags = 256
)

// Testsubprocessflags is a representation of the C type GTestSubprocessFlags.
type Testsubprocessflags int

const (
	// Testsubprocessflags_Stdin is a representation of the C type G_TEST_SUBPROCESS_INHERIT_STDIN.
	Testsubprocessflags_Stdin Testsubprocessflags = 1
	// Testsubprocessflags_Stdout is a representation of the C type G_TEST_SUBPROCESS_INHERIT_STDOUT.
	Testsubprocessflags_Stdout Testsubprocessflags = 2
	// Testsubprocessflags_Stderr is a representation of the C type G_TEST_SUBPROCESS_INHERIT_STDERR.
	Testsubprocessflags_Stderr Testsubprocessflags = 4
)

// Testtrapflags is a representation of the C type GTestTrapFlags.
type Testtrapflags int

const (
	// Testtrapflags_SilenceStdout is a representation of the C type G_TEST_TRAP_SILENCE_STDOUT.
	Testtrapflags_SilenceStdout Testtrapflags = 128
	// Testtrapflags_SilenceStderr is a representation of the C type G_TEST_TRAP_SILENCE_STDERR.
	Testtrapflags_SilenceStderr Testtrapflags = 256
	// Testtrapflags_InheritStdin is a representation of the C type G_TEST_TRAP_INHERIT_STDIN.
	Testtrapflags_InheritStdin Testtrapflags = 512
)

// Traverseflags is a representation of the C type GTraverseFlags.
type Traverseflags int

const (
	// Traverseflags_Leaves is a representation of the C type G_TRAVERSE_LEAVES.
	Traverseflags_Leaves Traverseflags = 1
	// Traverseflags_NonLeaves is a representation of the C type G_TRAVERSE_NON_LEAVES.
	Traverseflags_NonLeaves Traverseflags = 2
	// Traverseflags_All is a representation of the C type G_TRAVERSE_ALL.
	Traverseflags_All Traverseflags = 3
	// Traverseflags_Mask is a representation of the C type G_TRAVERSE_MASK.
	Traverseflags_Mask Traverseflags = 3
	// Traverseflags_Leafs is a representation of the C type G_TRAVERSE_LEAFS.
	Traverseflags_Leafs Traverseflags = 1
	// Traverseflags_NonLeafs is a representation of the C type G_TRAVERSE_NON_LEAFS.
	Traverseflags_NonLeafs Traverseflags = 2
)

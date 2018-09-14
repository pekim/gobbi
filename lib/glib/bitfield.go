// This is a generated file - DO NOT EDIT

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

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

type FileTest C.GFileTest

const (
	FILE_TEST_IS_REGULAR    FileTest = 1
	FILE_TEST_IS_SYMLINK    FileTest = 2
	FILE_TEST_IS_DIR        FileTest = 4
	FILE_TEST_IS_EXECUTABLE FileTest = 8
	FILE_TEST_EXISTS        FileTest = 16
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

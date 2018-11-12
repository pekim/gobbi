# `glib` bitfields

## `AsciiType`



- ASCII_ALNUM = 1
- ASCII_ALPHA = 2
- ASCII_CNTRL = 4
- ASCII_DIGIT = 8
- ASCII_GRAPH = 16
- ASCII_LOWER = 32
- ASCII_PRINT = 64
- ASCII_PUNCT = 128
- ASCII_SPACE = 256
- ASCII_UPPER = 512
- ASCII_XDIGIT = 1024

C - `GAsciiType`

## `FileTest`

A test to perform on a file using g_file_test().

- FILE_TEST_IS_REGULAR = 1
- FILE_TEST_IS_SYMLINK = 2
- FILE_TEST_IS_DIR = 4
- FILE_TEST_IS_EXECUTABLE = 8
- FILE_TEST_EXISTS = 16

C - `GFileTest`

## `FormatSizeFlags`

Flags to modify the format of the string returned by g_format_size_full().

- FORMAT_SIZE_DEFAULT = 0
- FORMAT_SIZE_LONG_FORMAT = 1
- FORMAT_SIZE_IEC_UNITS = 2
- FORMAT_SIZE_BITS = 4

C - `GFormatSizeFlags`

## `HookFlagMask`

Flags used internally in the #GHook implementation.

- HOOK_FLAG_ACTIVE = 1
- HOOK_FLAG_IN_CALL = 2
- HOOK_FLAG_MASK = 15

C - `GHookFlagMask`

## `IOCondition`

A bitwise combination representing a condition to watch for on an
event source.

- IO_IN = 1
- IO_OUT = 4
- IO_PRI = 2
- IO_ERR = 8
- IO_HUP = 16
- IO_NVAL = 32

C - `GIOCondition`

## `IOFlags`

Specifies properties of a #GIOChannel. Some of the flags can only be
read with g_io_channel_get_flags(), but not changed with
g_io_channel_set_flags().

- IO_FLAG_APPEND = 1
- IO_FLAG_NONBLOCK = 2
- IO_FLAG_IS_READABLE = 4
- IO_FLAG_IS_WRITABLE = 8
- IO_FLAG_IS_WRITEABLE = 8
- IO_FLAG_IS_SEEKABLE = 16
- IO_FLAG_MASK = 31
- IO_FLAG_GET_MASK = 31
- IO_FLAG_SET_MASK = 3

C - `GIOFlags`

## `KeyFileFlags`

Flags which influence the parsing.

- KEY_FILE_NONE = 0
- KEY_FILE_KEEP_COMMENTS = 1
- KEY_FILE_KEEP_TRANSLATIONS = 2

C - `GKeyFileFlags`

## `LogLevelFlags`

Flags specifying the level of log messages.

It is possible to change how GLib treats messages of the various
levels using g_log_set_handler() and g_log_set_fatal_mask().

- LOG_FLAG_RECURSION = 1
- LOG_FLAG_FATAL = 2
- LOG_LEVEL_ERROR = 4
- LOG_LEVEL_CRITICAL = 8
- LOG_LEVEL_WARNING = 16
- LOG_LEVEL_MESSAGE = 32
- LOG_LEVEL_INFO = 64
- LOG_LEVEL_DEBUG = 128
- LOG_LEVEL_MASK = -4

C - `GLogLevelFlags`

## `MarkupCollectType`

A mixed enumerated type and flags field. You must specify one type
(string, strdup, boolean, tristate).  Additionally, you may  optionally
bitwise OR the type with the flag %G_MARKUP_COLLECT_OPTIONAL.

It is likely that this enum will be extended in the future to
support other types.

- MARKUP_COLLECT_INVALID = 0
- MARKUP_COLLECT_STRING = 1
- MARKUP_COLLECT_STRDUP = 2
- MARKUP_COLLECT_BOOLEAN = 3
- MARKUP_COLLECT_TRISTATE = 4
- MARKUP_COLLECT_OPTIONAL = 65536

C - `GMarkupCollectType`

## `MarkupParseFlags`

Flags that affect the behaviour of the parser.

- MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG = 1
- MARKUP_TREAT_CDATA_AS_TEXT = 2
- MARKUP_PREFIX_ERROR_POSITION = 4
- MARKUP_IGNORE_QUALIFIED = 8

C - `GMarkupParseFlags`

## `OptionFlags`

Flags which modify individual options.

- OPTION_FLAG_NONE = 0
- OPTION_FLAG_HIDDEN = 1
- OPTION_FLAG_IN_MAIN = 2
- OPTION_FLAG_REVERSE = 4
- OPTION_FLAG_NO_ARG = 8
- OPTION_FLAG_FILENAME = 16
- OPTION_FLAG_OPTIONAL_ARG = 32
- OPTION_FLAG_NOALIAS = 64

C - `GOptionFlags`

## `RegexCompileFlags`

Flags specifying compile-time options.

- REGEX_CASELESS = 1
- REGEX_MULTILINE = 2
- REGEX_DOTALL = 4
- REGEX_EXTENDED = 8
- REGEX_ANCHORED = 16
- REGEX_DOLLAR_ENDONLY = 32
- REGEX_UNGREEDY = 512
- REGEX_RAW = 2048
- REGEX_NO_AUTO_CAPTURE = 4096
- REGEX_OPTIMIZE = 8192
- REGEX_FIRSTLINE = 262144
- REGEX_DUPNAMES = 524288
- REGEX_NEWLINE_CR = 1048576
- REGEX_NEWLINE_LF = 2097152
- REGEX_NEWLINE_CRLF = 3145728
- REGEX_NEWLINE_ANYCRLF = 5242880
- REGEX_BSR_ANYCRLF = 8388608
- REGEX_JAVASCRIPT_COMPAT = 33554432

C - `GRegexCompileFlags`

## `RegexMatchFlags`

Flags specifying match-time options.

- REGEX_MATCH_ANCHORED = 16
- REGEX_MATCH_NOTBOL = 128
- REGEX_MATCH_NOTEOL = 256
- REGEX_MATCH_NOTEMPTY = 1024
- REGEX_MATCH_PARTIAL = 32768
- REGEX_MATCH_NEWLINE_CR = 1048576
- REGEX_MATCH_NEWLINE_LF = 2097152
- REGEX_MATCH_NEWLINE_CRLF = 3145728
- REGEX_MATCH_NEWLINE_ANY = 4194304
- REGEX_MATCH_NEWLINE_ANYCRLF = 5242880
- REGEX_MATCH_BSR_ANYCRLF = 8388608
- REGEX_MATCH_BSR_ANY = 16777216
- REGEX_MATCH_PARTIAL_SOFT = 32768
- REGEX_MATCH_PARTIAL_HARD = 134217728
- REGEX_MATCH_NOTEMPTY_ATSTART = 268435456

C - `GRegexMatchFlags`

## `SpawnFlags`

Flags passed to g_spawn_sync(), g_spawn_async() and g_spawn_async_with_pipes().

- SPAWN_DEFAULT = 0
- SPAWN_LEAVE_DESCRIPTORS_OPEN = 1
- SPAWN_DO_NOT_REAP_CHILD = 2
- SPAWN_SEARCH_PATH = 4
- SPAWN_STDOUT_TO_DEV_NULL = 8
- SPAWN_STDERR_TO_DEV_NULL = 16
- SPAWN_CHILD_INHERITS_STDIN = 32
- SPAWN_FILE_AND_ARGV_ZERO = 64
- SPAWN_SEARCH_PATH_FROM_ENVP = 128
- SPAWN_CLOEXEC_PIPES = 256

C - `GSpawnFlags`

## `TestSubprocessFlags`

Flags to pass to g_test_trap_subprocess() to control input and output.

Note that in contrast with g_test_trap_fork(), the default is to
not show stdout and stderr.

- TEST_SUBPROCESS_INHERIT_STDIN = 1
- TEST_SUBPROCESS_INHERIT_STDOUT = 2
- TEST_SUBPROCESS_INHERIT_STDERR = 4

C - `GTestSubprocessFlags`

## `TestTrapFlags`

Test traps are guards around forked tests.
These flags determine what traps to set.

- TEST_TRAP_SILENCE_STDOUT = 128
- TEST_TRAP_SILENCE_STDERR = 256
- TEST_TRAP_INHERIT_STDIN = 512

C - `GTestTrapFlags`

## `TraverseFlags`

Specifies which nodes are visited during several of the tree
functions, including g_node_traverse() and g_node_find().

- TRAVERSE_LEAVES = 1
- TRAVERSE_NON_LEAVES = 2
- TRAVERSE_ALL = 3
- TRAVERSE_MASK = 3
- TRAVERSE_LEAFS = 1
- TRAVERSE_NON_LEAFS = 2

C - `GTraverseFlags`

(int=32)
- SPAWN_FILE_AND_ARGV_ZERO = %!s(int=64)
- SPAWN_SEARCH_PATH_FROM_ENVP = %!s(int=128)
- SPAWN_CLOEXEC_PIPES = %!s(int=256)
## `TestSubprocessFlags`

Flags to pass to g_test_trap_subprocess() to control input and output.

Note that in contrast with g_test_trap_fork(), the default is to
not show stdout and stderr.

C - `GTestSubprocessFlags`

- TEST_SUBPROCESS_INHERIT_STDIN = %!s(int=1)
- TEST_SUBPROCESS_INHERIT_STDOUT = %!s(int=2)
- TEST_SUBPROCESS_INHERIT_STDERR = %!s(int=4)
## `TestTrapFlags`

Test traps are guards around forked tests.
These flags determine what traps to set.

C - `GTestTrapFlags`

- TEST_TRAP_SILENCE_STDOUT = %!s(int=128)
- TEST_TRAP_SILENCE_STDERR = %!s(int=256)
- TEST_TRAP_INHERIT_STDIN = %!s(int=512)
## `TraverseFlags`

Specifies which nodes are visited during several of the tree
functions, including g_node_traverse() and g_node_find().

C - `GTraverseFlags`

- TRAVERSE_LEAVES = %!s(int=1)
- TRAVERSE_NON_LEAVES = %!s(int=2)
- TRAVERSE_ALL = %!s(int=3)
- TRAVERSE_MASK = %!s(int=3)
- TRAVERSE_LEAFS = %!s(int=1)
- TRAVERSE_NON_LEAFS = %!s(int=2)

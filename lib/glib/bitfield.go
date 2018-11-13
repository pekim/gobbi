// This is a generated file - DO NOT EDIT

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// A test to perform on a file using g_file_test().
type GFileTest C.GFileTest

const (
	/*
	   %TRUE if the file is a regular file
	       (not a directory). Note that this test will also return %TRUE
	       if the tested file is a symlink to a regular file.
	*/
	FILE_TEST_IS_REGULAR GFileTest = 1
	// %TRUE if the file is a symlink.
	FILE_TEST_IS_SYMLINK GFileTest = 2
	// %TRUE if the file is a directory.
	FILE_TEST_IS_DIR GFileTest = 4
	// %TRUE if the file is executable.
	FILE_TEST_IS_EXECUTABLE GFileTest = 8
	/*
	   %TRUE if the file exists. It may or may not
	       be a regular file.
	*/
	FILE_TEST_EXISTS GFileTest = 16
)

// Flags to modify the format of the string returned by g_format_size_full().
type FormatSizeFlags C.GFormatSizeFlags

const (
	// behave the same as g_format_size()
	FORMAT_SIZE_DEFAULT FormatSizeFlags = 0
	/*
	   include the exact number of bytes as part
	       of the returned string.  For example, "45.6 kB (45,612 bytes)".
	*/
	FORMAT_SIZE_LONG_FORMAT FormatSizeFlags = 1
	/*
	   use IEC (base 1024) units with "KiB"-style
	       suffixes. IEC units should only be used for reporting things with
	       a strong "power of 2" basis, like RAM sizes or RAID stripe sizes.
	       Network and storage sizes should be reported in the normal SI units.
	*/
	FORMAT_SIZE_IEC_UNITS FormatSizeFlags = 2
	/*
	   set the size as a quantity in bits, rather than
	       bytes, and return units in bits. For example, ‘Mb’ rather than ‘MB’.
	*/
	FORMAT_SIZE_BITS FormatSizeFlags = 4
)

// Flags used internally in the #GHook implementation.
type HookFlagMask C.GHookFlagMask

const (
	// set if the hook has not been destroyed
	HOOK_FLAG_ACTIVE HookFlagMask = 1
	// set if the hook is currently being run
	HOOK_FLAG_IN_CALL HookFlagMask = 2
	/*
	   A mask covering all bits reserved for
	     hook flags; see %G_HOOK_FLAG_USER_SHIFT
	*/
	HOOK_FLAG_MASK HookFlagMask = 15
)

/*
A bitwise combination representing a condition to watch for on an
event source.
*/
type IOCondition C.GIOCondition

const (
	// There is data to read.
	IO_IN IOCondition = 1
	// Data can be written (without blocking).
	IO_OUT IOCondition = 4
	// There is urgent data to read.
	IO_PRI IOCondition = 2
	// Error condition.
	IO_ERR IOCondition = 8
	/*
	   Hung up (the connection has been broken, usually for
	              pipes and sockets).
	*/
	IO_HUP IOCondition = 16
	// Invalid request. The file descriptor is not open.
	IO_NVAL IOCondition = 32
)

/*
Specifies properties of a #GIOChannel. Some of the flags can only be
read with g_io_channel_get_flags(), but not changed with
g_io_channel_set_flags().
*/
type IOFlags C.GIOFlags

const (
	/*
	   turns on append mode, corresponds to %O_APPEND
	       (see the documentation of the UNIX open() syscall)
	*/
	IO_FLAG_APPEND IOFlags = 1
	/*
	   turns on nonblocking mode, corresponds to
	       %O_NONBLOCK/%O_NDELAY (see the documentation of the UNIX open()
	       syscall)
	*/
	IO_FLAG_NONBLOCK IOFlags = 2
	/*
	   indicates that the io channel is readable.
	       This flag cannot be changed.
	*/
	IO_FLAG_IS_READABLE IOFlags = 4
	/*
	   indicates that the io channel is writable.
	       This flag cannot be changed.
	*/
	IO_FLAG_IS_WRITABLE IOFlags = 8
	/*
	   a misspelled version of @G_IO_FLAG_IS_WRITABLE
	       that existed before the spelling was fixed in GLib 2.30. It is kept
	       here for compatibility reasons. Deprecated since 2.30
	*/
	IO_FLAG_IS_WRITEABLE IOFlags = 8
	/*
	   indicates that the io channel is seekable,
	       i.e. that g_io_channel_seek_position() can be used on it.
	       This flag cannot be changed.
	*/
	IO_FLAG_IS_SEEKABLE IOFlags = 16
	// the mask that specifies all the valid flags.
	IO_FLAG_MASK IOFlags = 31
	/*
	   the mask of the flags that are returned from
	       g_io_channel_get_flags()
	*/
	IO_FLAG_GET_MASK IOFlags = 31
	/*
	   the mask of the flags that the user can modify
	       with g_io_channel_set_flags()
	*/
	IO_FLAG_SET_MASK IOFlags = 3
)

// Flags which influence the parsing.
type KeyFileFlags C.GKeyFileFlags

const (
	// No flags, default behaviour
	KEY_FILE_NONE KeyFileFlags = 0
	/*
	   Use this flag if you plan to write the
	       (possibly modified) contents of the key file back to a file;
	       otherwise all comments will be lost when the key file is
	       written back.
	*/
	KEY_FILE_KEEP_COMMENTS KeyFileFlags = 1
	/*
	   Use this flag if you plan to write the
	       (possibly modified) contents of the key file back to a file;
	       otherwise only the translations for the current language will be
	       written back.
	*/
	KEY_FILE_KEEP_TRANSLATIONS KeyFileFlags = 2
)

/*
Flags specifying the level of log messages.

It is possible to change how GLib treats messages of the various
levels using g_log_set_handler() and g_log_set_fatal_mask().
*/
type LogLevelFlags C.GLogLevelFlags

const (
	// internal flag
	LOG_FLAG_RECURSION LogLevelFlags = 1
	// internal flag
	LOG_FLAG_FATAL LogLevelFlags = 2
	/*
	   log level for errors, see g_error().
	       This level is also used for messages produced by g_assert().
	*/
	LOG_LEVEL_ERROR LogLevelFlags = 4
	/*
	   log level for critical warning messages, see
	       g_critical().
	       This level is also used for messages produced by g_return_if_fail()
	       and g_return_val_if_fail().
	*/
	LOG_LEVEL_CRITICAL LogLevelFlags = 8
	// log level for warnings, see g_warning()
	LOG_LEVEL_WARNING LogLevelFlags = 16
	// log level for messages, see g_message()
	LOG_LEVEL_MESSAGE LogLevelFlags = 32
	// log level for informational messages, see g_info()
	LOG_LEVEL_INFO LogLevelFlags = 64
	// log level for debug messages, see g_debug()
	LOG_LEVEL_DEBUG LogLevelFlags = 128
	// a mask including all log levels
	LOG_LEVEL_MASK LogLevelFlags = -4
)

/*
A mixed enumerated type and flags field. You must specify one type
(string, strdup, boolean, tristate).  Additionally, you may  optionally
bitwise OR the type with the flag %G_MARKUP_COLLECT_OPTIONAL.

It is likely that this enum will be extended in the future to
support other types.
*/
type MarkupCollectType C.GMarkupCollectType

const (
	/*
	   used to terminate the list of attributes
	       to collect
	*/
	MARKUP_COLLECT_INVALID MarkupCollectType = 0
	/*
	   collect the string pointer directly from
	       the attribute_values[] array. Expects a parameter of type (const
	       char **). If %G_MARKUP_COLLECT_OPTIONAL is specified and the
	       attribute isn't present then the pointer will be set to %NULL
	*/
	MARKUP_COLLECT_STRING MarkupCollectType = 1
	/*
	   as with %G_MARKUP_COLLECT_STRING, but
	       expects a parameter of type (char **) and g_strdup()s the
	       returned pointer. The pointer must be freed with g_free()
	*/
	MARKUP_COLLECT_STRDUP MarkupCollectType = 2
	/*
	   expects a parameter of type (gboolean *)
	       and parses the attribute value as a boolean. Sets %FALSE if the
	       attribute isn't present. Valid boolean values consist of
	       (case-insensitive) "false", "f", "no", "n", "0" and "true", "t",
	       "yes", "y", "1"
	*/
	MARKUP_COLLECT_BOOLEAN MarkupCollectType = 3
	/*
	   as with %G_MARKUP_COLLECT_BOOLEAN, but
	       in the case of a missing attribute a value is set that compares
	       equal to neither %FALSE nor %TRUE G_MARKUP_COLLECT_OPTIONAL is
	       implied
	*/
	MARKUP_COLLECT_TRISTATE MarkupCollectType = 4
	/*
	   can be bitwise ORed with the other fields.
	       If present, allows the attribute not to appear. A default value
	       is set depending on what value type is used
	*/
	MARKUP_COLLECT_OPTIONAL MarkupCollectType = 65536
)

// Flags that affect the behaviour of the parser.
type MarkupParseFlags C.GMarkupParseFlags

const (
	// flag you should not use
	MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG MarkupParseFlags = 1
	/*
	   When this flag is set, CDATA marked
	       sections are not passed literally to the @passthrough function of
	       the parser. Instead, the content of the section (without the
	       `<![CDATA[` and `]]>`) is
	       passed to the @text function. This flag was added in GLib 2.12
	*/
	MARKUP_TREAT_CDATA_AS_TEXT MarkupParseFlags = 2
	/*
	   Normally errors caught by GMarkup
	       itself have line/column information prefixed to them to let the
	       caller know the location of the error. When this flag is set the
	       location information is also prefixed to errors generated by the
	       #GMarkupParser implementation functions
	*/
	MARKUP_PREFIX_ERROR_POSITION MarkupParseFlags = 4
	/*
	   Ignore (don't report) qualified
	       attributes and tags, along with their contents.  A qualified
	       attribute or tag is one that contains ':' in its name (ie: is in
	       another namespace).  Since: 2.40.
	*/
	MARKUP_IGNORE_QUALIFIED MarkupParseFlags = 8
)

// Flags which modify individual options.
type OptionFlags C.GOptionFlags

const (
	// No flags. Since: 2.42.
	OPTION_FLAG_NONE OptionFlags = 0
	// The option doesn't appear in `--help` output.
	OPTION_FLAG_HIDDEN OptionFlags = 1
	/*
	   The option appears in the main section of the
	       `--help` output, even if it is defined in a group.
	*/
	OPTION_FLAG_IN_MAIN OptionFlags = 2
	/*
	   For options of the %G_OPTION_ARG_NONE kind, this
	       flag indicates that the sense of the option is reversed.
	*/
	OPTION_FLAG_REVERSE OptionFlags = 4
	/*
	   For options of the %G_OPTION_ARG_CALLBACK kind,
	       this flag indicates that the callback does not take any argument
	       (like a %G_OPTION_ARG_NONE option). Since 2.8
	*/
	OPTION_FLAG_NO_ARG OptionFlags = 8
	/*
	   For options of the %G_OPTION_ARG_CALLBACK
	       kind, this flag indicates that the argument should be passed to the
	       callback in the GLib filename encoding rather than UTF-8. Since 2.8
	*/
	OPTION_FLAG_FILENAME OptionFlags = 16
	/*
	   For options of the %G_OPTION_ARG_CALLBACK
	       kind, this flag indicates that the argument supply is optional.
	       If no argument is given then data of %GOptionParseFunc will be
	       set to NULL. Since 2.8
	*/
	OPTION_FLAG_OPTIONAL_ARG OptionFlags = 32
	/*
	   This flag turns off the automatic conflict
	       resolution which prefixes long option names with `groupname-` if
	       there is a conflict. This option should only be used in situations
	       where aliasing is necessary to model some legacy commandline interface.
	       It is not safe to use this option, unless all option groups are under
	       your direct control. Since 2.8.
	*/
	OPTION_FLAG_NOALIAS OptionFlags = 64
)

// Flags passed to g_spawn_sync(), g_spawn_async() and g_spawn_async_with_pipes().
type SpawnFlags C.GSpawnFlags

const (
	// no flags, default behaviour
	SPAWN_DEFAULT SpawnFlags = 0
	/*
	   the parent's open file descriptors will
	       be inherited by the child; otherwise all descriptors except stdin,
	       stdout and stderr will be closed before calling exec() in the child.
	*/
	SPAWN_LEAVE_DESCRIPTORS_OPEN SpawnFlags = 1
	/*
	   the child will not be automatically reaped;
	       you must use g_child_watch_add() yourself (or call waitpid() or handle
	       `SIGCHLD` yourself), or the child will become a zombie.
	*/
	SPAWN_DO_NOT_REAP_CHILD SpawnFlags = 2
	/*
	   `argv[0]` need not be an absolute path, it will be
	       looked for in the user's `PATH`.
	*/
	SPAWN_SEARCH_PATH SpawnFlags = 4
	/*
	   the child's standard output will be discarded,
	       instead of going to the same location as the parent's standard output.
	*/
	SPAWN_STDOUT_TO_DEV_NULL SpawnFlags = 8
	// the child's standard error will be discarded.
	SPAWN_STDERR_TO_DEV_NULL SpawnFlags = 16
	/*
	   the child will inherit the parent's standard
	       input (by default, the child's standard input is attached to `/dev/null`).
	*/
	SPAWN_CHILD_INHERITS_STDIN SpawnFlags = 32
	/*
	   the first element of `argv` is the file to
	       execute, while the remaining elements are the actual argument vector
	       to pass to the file. Normally g_spawn_async_with_pipes() uses `argv[0]`
	       as the file to execute, and passes all of `argv` to the child.
	*/
	SPAWN_FILE_AND_ARGV_ZERO SpawnFlags = 64
	/*
	   if `argv[0]` is not an abolute path,
	       it will be looked for in the `PATH` from the passed child environment.
	       Since: 2.34
	*/
	SPAWN_SEARCH_PATH_FROM_ENVP SpawnFlags = 128
	/*
	   create all pipes with the `O_CLOEXEC` flag set.
	       Since: 2.40
	*/
	SPAWN_CLOEXEC_PIPES SpawnFlags = 256
)

/*
Flags to pass to g_test_trap_subprocess() to control input and output.

Note that in contrast with g_test_trap_fork(), the default is to
not show stdout and stderr.
*/
type TestSubprocessFlags C.GTestSubprocessFlags

const (
	/*
	   If this flag is given, the child
	       process will inherit the parent's stdin. Otherwise, the child's
	       stdin is redirected to `/dev/null`.
	*/
	TEST_SUBPROCESS_INHERIT_STDIN TestSubprocessFlags = 1
	/*
	   If this flag is given, the child
	       process will inherit the parent's stdout. Otherwise, the child's
	       stdout will not be visible, but it will be captured to allow
	       later tests with g_test_trap_assert_stdout().
	*/
	TEST_SUBPROCESS_INHERIT_STDOUT TestSubprocessFlags = 2
	/*
	   If this flag is given, the child
	       process will inherit the parent's stderr. Otherwise, the child's
	       stderr will not be visible, but it will be captured to allow
	       later tests with g_test_trap_assert_stderr().
	*/
	TEST_SUBPROCESS_INHERIT_STDERR TestSubprocessFlags = 4
)

/*
Test traps are guards around forked tests.
These flags determine what traps to set.
*/
type TestTrapFlags C.GTestTrapFlags

const (
	/*
	   Redirect stdout of the test child to
	       `/dev/null` so it cannot be observed on the console during test
	       runs. The actual output is still captured though to allow later
	       tests with g_test_trap_assert_stdout().
	*/
	TEST_TRAP_SILENCE_STDOUT TestTrapFlags = 128
	/*
	   Redirect stderr of the test child to
	       `/dev/null` so it cannot be observed on the console during test
	       runs. The actual output is still captured though to allow later
	       tests with g_test_trap_assert_stderr().
	*/
	TEST_TRAP_SILENCE_STDERR TestTrapFlags = 256
	/*
	   If this flag is given, stdin of the
	       child process is shared with stdin of its parent process.
	       It is redirected to `/dev/null` otherwise.
	*/
	TEST_TRAP_INHERIT_STDIN TestTrapFlags = 512
)

/*
Specifies which nodes are visited during several of the tree
functions, including g_node_traverse() and g_node_find().
*/
type TraverseFlags C.GTraverseFlags

const (
	/*
	   only leaf nodes should be visited. This name has
	                       been introduced in 2.6, for older version use
	                       %G_TRAVERSE_LEAFS.
	*/
	TRAVERSE_LEAVES TraverseFlags = 1
	/*
	   only non-leaf nodes should be visited. This
	                           name has been introduced in 2.6, for older
	                           version use %G_TRAVERSE_NON_LEAFS.
	*/
	TRAVERSE_NON_LEAVES TraverseFlags = 2
	// all nodes should be visited.
	TRAVERSE_ALL TraverseFlags = 3
	// a mask of all traverse flags.
	TRAVERSE_MASK TraverseFlags = 3
	// identical to %G_TRAVERSE_LEAVES.
	TRAVERSE_LEAFS TraverseFlags = 1
	// identical to %G_TRAVERSE_NON_LEAVES.
	TRAVERSE_NON_LEAFS TraverseFlags = 2
)

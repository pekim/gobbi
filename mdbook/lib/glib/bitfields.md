# `glib` bitfields

## `AsciiType`



C - `GAsciiType`

## `FileTest`

A test to perform on a file using g_file_test().

C - `GFileTest`

## `FormatSizeFlags`

Flags to modify the format of the string returned by g_format_size_full().

C - `GFormatSizeFlags`

## `HookFlagMask`

Flags used internally in the #GHook implementation.

C - `GHookFlagMask`

## `IOCondition`

A bitwise combination representing a condition to watch for on an
event source.

C - `GIOCondition`

## `IOFlags`

Specifies properties of a #GIOChannel. Some of the flags can only be
read with g_io_channel_get_flags(), but not changed with
g_io_channel_set_flags().

C - `GIOFlags`

## `KeyFileFlags`

Flags which influence the parsing.

C - `GKeyFileFlags`

## `LogLevelFlags`

Flags specifying the level of log messages.

It is possible to change how GLib treats messages of the various
levels using g_log_set_handler() and g_log_set_fatal_mask().

C - `GLogLevelFlags`

## `MarkupCollectType`

A mixed enumerated type and flags field. You must specify one type
(string, strdup, boolean, tristate).  Additionally, you may  optionally
bitwise OR the type with the flag %G_MARKUP_COLLECT_OPTIONAL.

It is likely that this enum will be extended in the future to
support other types.

C - `GMarkupCollectType`

## `MarkupParseFlags`

Flags that affect the behaviour of the parser.

C - `GMarkupParseFlags`

## `OptionFlags`

Flags which modify individual options.

C - `GOptionFlags`

## `RegexCompileFlags`

Flags specifying compile-time options.

C - `GRegexCompileFlags`

## `RegexMatchFlags`

Flags specifying match-time options.

C - `GRegexMatchFlags`

## `SpawnFlags`

Flags passed to g_spawn_sync(), g_spawn_async() and g_spawn_async_with_pipes().

C - `GSpawnFlags`

## `TestSubprocessFlags`

Flags to pass to g_test_trap_subprocess() to control input and output.

Note that in contrast with g_test_trap_fork(), the default is to
not show stdout and stderr.

C - `GTestSubprocessFlags`

## `TestTrapFlags`

Test traps are guards around forked tests.
These flags determine what traps to set.

C - `GTestTrapFlags`

## `TraverseFlags`

Specifies which nodes are visited during several of the tree
functions, including g_node_traverse() and g_node_find().

C - `GTraverseFlags`


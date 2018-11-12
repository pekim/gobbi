# `glib` enums

## `BookmarkFileError`

Error codes returned by bookmark file parsing.

C - `GBookmarkFileError`

## `ChecksumType`

The hashing algorithm to be used by #GChecksum when performing the
digest of some data.

Note that the #GChecksumType enumeration may be extended at a later
date to include new hashing algorithm types.

C - `GChecksumType`

## `ConvertError`

Error codes returned by character set conversion routines.

C - `GConvertError`

## `DateDMY`

This enumeration isn't used in the API, but may be useful if you need
to mark a number as a day, month, or year.

C - `GDateDMY`

## `DateMonth`

Enumeration representing a month; values are #G_DATE_JANUARY,
&num;G_DATE_FEBRUARY, etc. #G_DATE_BAD_MONTH is the invalid value.

C - `GDateMonth`

## `DateWeekday`

Enumeration representing a day of the week; #G_DATE_MONDAY,
&num;G_DATE_TUESDAY, etc. #G_DATE_BAD_WEEKDAY is an invalid weekday.

C - `GDateWeekday`

## `ErrorType`

The possible errors, used in the @v_error field
of #GTokenValue, when the token is a %G_TOKEN_ERROR.

C - `GErrorType`

## `FileError`

Values corresponding to @errno codes returned from file operations
on UNIX. Unlike @errno codes, GFileError values are available on
all systems, even Windows. The exact meaning of each code depends
on what sort of file operation you were performing; the UNIX
documentation gives more details. The following error code descriptions
come from the GNU C Library manual, and are under the copyright
of that manual.

It's not very portable to make detailed assumptions about exactly
which errors will be returned from a given operation. Some errors
don't occur on some systems, etc., sometimes there are subtle
differences in when a system will report a given error, etc.

C - `GFileError`

## `IOChannelError`

Error codes returned by #GIOChannel operations.

C - `GIOChannelError`

## `IOError`

#GIOError is only used by the deprecated functions
g_io_channel_read(), g_io_channel_write(), and g_io_channel_seek().

C - `GIOError`

## `IOStatus`

Stati returned by most of the #GIOFuncs functions.

C - `GIOStatus`

## `KeyFileError`

Error codes returned by key file parsing.

C - `GKeyFileError`

## `LogWriterOutput`

Return values from #GLogWriterFuncs to indicate whether the given log entry
was successfully handled by the writer, or whether there was an error in
handling it (and hence a fallback writer should be used).

If a #GLogWriterFunc ignores a log entry, it should return
%G_LOG_WRITER_HANDLED.

C - `GLogWriterOutput`

## `MarkupError`

Error codes returned by markup parsing.

C - `GMarkupError`

## `NormalizeMode`

Defines how a Unicode string is transformed in a canonical
form, standardizing such issues as whether a character with
an accent is represented as a base character and combining
accent or as a single precomposed character. Unicode strings
should generally be normalized before comparing them.

C - `GNormalizeMode`

## `NumberParserError`

Error codes returned by functions converting a string to a number.

C - `GNumberParserError`

## `OnceStatus`

The possible statuses of a one-time initialization function
controlled by a #GOnce struct.

C - `GOnceStatus`

## `OptionArg`

The #GOptionArg enum values determine which type of extra argument the
options expect to find. If an option expects an extra argument, it can
be specified in several ways; with a short option: `-x arg`, with a long
option: `--name arg` or combined in a single argument: `--name=arg`.

C - `GOptionArg`

## `OptionError`

Error codes returned by option parsing.

C - `GOptionError`

## `RegexError`

Error codes returned by regular expressions functions.

C - `GRegexError`

## `SeekType`

An enumeration specifying the base position for a
g_io_channel_seek_position() operation.

C - `GSeekType`

## `ShellError`

Error codes returned by shell functions.

C - `GShellError`

## `SliceConfig`



C - `GSliceConfig`

## `SpawnError`

Error codes returned by spawning processes.

C - `GSpawnError`

## `TestFileType`

The type of file to return the filename for, when used with
g_test_build_filename().

These two options correspond rather directly to the 'dist' and
'built' terminology that automake uses and are explicitly used to
distinguish between the 'srcdir' and 'builddir' being separate.  All
files in your project should either be dist (in the
`EXTRA_DIST` or `dist_schema_DATA`
sense, in which case they will always be in the srcdir) or built (in
the `BUILT_SOURCES` sense, in which case they will
always be in the builddir).

Note: as a general rule of automake, files that are generated only as
part of the build-from-git process (but then are distributed with the
tarball) always go in srcdir (even if doing a srcdir != builddir
build from git) and are considered as distributed files.

C - `GTestFileType`

## `TestLogType`



C - `GTestLogType`

## `ThreadError`

Possible errors of thread related functions.

C - `GThreadError`

## `TimeType`

Disambiguates a given time in two ways.

First, specifies if the given time is in universal or local time.

Second, if the time is in local time, specifies if it is local
standard time or local daylight time.  This is important for the case
where the same local time occurs twice (during daylight savings time
transitions, for example).

C - `GTimeType`

## `TokenType`

The possible types of token returned from each
g_scanner_get_next_token() call.

C - `GTokenType`

## `TraverseType`

Specifies the type of traveral performed by g_tree_traverse(),
g_node_traverse() and g_node_find(). The different orders are
illustrated here:
- In order: A, B, C, D, E, F, G, H, I
  ![](Sorted_binary_tree_inorder.svg)
- Pre order: F, B, A, D, C, E, G, I, H
  ![](Sorted_binary_tree_preorder.svg)
- Post order: A, C, E, D, B, H, I, G, F
  ![](Sorted_binary_tree_postorder.svg)
- Level order: F, B, G, A, D, I, C, E, H
  ![](Sorted_binary_tree_breadth-first_traversal.svg)

C - `GTraverseType`

## `UnicodeBreakType`

These are the possible line break classifications.

Since new unicode versions may add new types here, applications should be ready
to handle unknown values. They may be regarded as %G_UNICODE_BREAK_UNKNOWN.

See [Unicode Line Breaking Algorithm](http://www.unicode.org/unicode/reports/tr14/).

C - `GUnicodeBreakType`

## `UnicodeScript`

The #GUnicodeScript enumeration identifies different writing
systems. The values correspond to the names as defined in the
Unicode standard. The enumeration has been added in GLib 2.14,
and is interchangeable with #PangoScript.

Note that new types may be added in the future. Applications
should be ready to handle unknown values.
See [Unicode Standard Annex #24: Script names](http://www.unicode.org/reports/tr24/).

C - `GUnicodeScript`

## `UnicodeType`

These are the possible character classifications from the
Unicode specification.
See [Unicode Character Database](http://www.unicode.org/reports/tr44/#General_Category_Values).

C - `GUnicodeType`

## `UserDirectory`

These are logical ids for special directories which are defined
depending on the platform used. You should use g_get_user_special_dir()
to retrieve the full path associated to the logical id.

The #GUserDirectory enumeration can be extended at later date. Not
every platform has a directory for every logical id in this
enumeration.

C - `GUserDirectory`

## `VariantClass`

The range of possible top-level types of #GVariant instances.

C - `GVariantClass`

## `VariantParseError`

Error codes returned by parsing text-format GVariants.

C - `GVariantParseError`


# `glib` enums

## `BookmarkFileError`

Error codes returned by bookmark file parsing.

C - `GBookmarkFileError`

### BOOKMARK_FILE_ERROR_INVALID_URI = 0
URI was ill-formed

### BOOKMARK_FILE_ERROR_INVALID_VALUE = 1
a requested field was not found

### BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED = 2
a requested application did
    not register a bookmark

### BOOKMARK_FILE_ERROR_URI_NOT_FOUND = 3
a requested URI was not found

### BOOKMARK_FILE_ERROR_READ = 4
document was ill formed

### BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING = 5
the text being parsed was
    in an unknown encoding

### BOOKMARK_FILE_ERROR_WRITE = 6
an error occurred while writing

### BOOKMARK_FILE_ERROR_FILE_NOT_FOUND = 7
requested file was not found


## `ChecksumType`

The hashing algorithm to be used by #GChecksum when performing the
digest of some data.

Note that the #GChecksumType enumeration may be extended at a later
date to include new hashing algorithm types.

C - `GChecksumType`

### CHECKSUM_MD5 = 0
Use the MD5 hashing algorithm

### CHECKSUM_SHA1 = 1
Use the SHA-1 hashing algorithm

### CHECKSUM_SHA256 = 2
Use the SHA-256 hashing algorithm

### CHECKSUM_SHA512 = 3
Use the SHA-512 hashing algorithm (Since: 2.36)

### CHECKSUM_SHA384 = 4
Use the SHA-384 hashing algorithm (Since: 2.51)


## `ConvertError`

Error codes returned by character set conversion routines.

C - `GConvertError`

### CONVERT_ERROR_NO_CONVERSION = 0
Conversion between the requested character
    sets is not supported.

### CONVERT_ERROR_ILLEGAL_SEQUENCE = 1
Invalid byte sequence in conversion input;
   or the character sequence could not be represented in the target
   character set.

### CONVERT_ERROR_FAILED = 2
Conversion failed for some reason.

### CONVERT_ERROR_PARTIAL_INPUT = 3
Partial character sequence at end of input.

### CONVERT_ERROR_BAD_URI = 4
URI is invalid.

### CONVERT_ERROR_NOT_ABSOLUTE_PATH = 5
Pathname is not an absolute path.

### CONVERT_ERROR_NO_MEMORY = 6
No memory available. Since: 2.40

### CONVERT_ERROR_EMBEDDED_NUL = 7
An embedded NUL character is present in
    conversion output where a NUL-terminated string is expected.
    Since: 2.56


## `DateDMY`

This enumeration isn't used in the API, but may be useful if you need
to mark a number as a day, month, or year.

C - `GDateDMY`

### DATE_DAY = 0
a day

### DATE_MONTH = 1
a month

### DATE_YEAR = 2
a year


## `DateMonth`

Enumeration representing a month; values are #G_DATE_JANUARY,
&num;G_DATE_FEBRUARY, etc. #G_DATE_BAD_MONTH is the invalid value.

C - `GDateMonth`

### DATE_BAD_MONTH = 0
invalid value

### DATE_JANUARY = 1
January

### DATE_FEBRUARY = 2
February

### DATE_MARCH = 3
March

### DATE_APRIL = 4
April

### DATE_MAY = 5
May

### DATE_JUNE = 6
June

### DATE_JULY = 7
July

### DATE_AUGUST = 8
August

### DATE_SEPTEMBER = 9
September

### DATE_OCTOBER = 10
October

### DATE_NOVEMBER = 11
November

### DATE_DECEMBER = 12
December


## `DateWeekday`

Enumeration representing a day of the week; #G_DATE_MONDAY,
&num;G_DATE_TUESDAY, etc. #G_DATE_BAD_WEEKDAY is an invalid weekday.

C - `GDateWeekday`

### DATE_BAD_WEEKDAY = 0
invalid value

### DATE_MONDAY = 1
Monday

### DATE_TUESDAY = 2
Tuesday

### DATE_WEDNESDAY = 3
Wednesday

### DATE_THURSDAY = 4
Thursday

### DATE_FRIDAY = 5
Friday

### DATE_SATURDAY = 6
Saturday

### DATE_SUNDAY = 7
Sunday


## `ErrorType`

The possible errors, used in the @v_error field
of #GTokenValue, when the token is a %G_TOKEN_ERROR.

C - `GErrorType`

### ERR_UNKNOWN = 0
unknown error

### ERR_UNEXP_EOF = 1
unexpected end of file

### ERR_UNEXP_EOF_IN_STRING = 2
unterminated string constant

### ERR_UNEXP_EOF_IN_COMMENT = 3
unterminated comment

### ERR_NON_DIGIT_IN_CONST = 4
non-digit character in a number

### ERR_DIGIT_RADIX = 5
digit beyond radix in a number

### ERR_FLOAT_RADIX = 6
non-decimal floating point number

### ERR_FLOAT_MALFORMED = 7
malformed floating point number


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

### FILE_ERROR_EXIST = 0
Operation not permitted; only the owner of
    the file (or other resource) or processes with special privileges
    can perform the operation.

### FILE_ERROR_ISDIR = 1
File is a directory; you cannot open a directory
    for writing, or create or remove hard links to it.

### FILE_ERROR_ACCES = 2
Permission denied; the file permissions do not
    allow the attempted operation.

### FILE_ERROR_NAMETOOLONG = 3
Filename too long.

### FILE_ERROR_NOENT = 4
No such file or directory. This is a "file
    doesn't exist" error for ordinary files that are referenced in
    contexts where they are expected to already exist.

### FILE_ERROR_NOTDIR = 5
A file that isn't a directory was specified when
    a directory is required.

### FILE_ERROR_NXIO = 6
No such device or address. The system tried to
    use the device represented by a file you specified, and it
    couldn't find the device. This can mean that the device file was
    installed incorrectly, or that the physical device is missing or
    not correctly attached to the computer.

### FILE_ERROR_NODEV = 7
The underlying file system of the specified file
    does not support memory mapping.

### FILE_ERROR_ROFS = 8
The directory containing the new link can't be
    modified because it's on a read-only file system.

### FILE_ERROR_TXTBSY = 9
Text file busy.

### FILE_ERROR_FAULT = 10
You passed in a pointer to bad memory.
    (GLib won't reliably return this, don't pass in pointers to bad
    memory.)

### FILE_ERROR_LOOP = 11
Too many levels of symbolic links were encountered
    in looking up a file name. This often indicates a cycle of symbolic
    links.

### FILE_ERROR_NOSPC = 12
No space left on device; write operation on a
    file failed because the disk is full.

### FILE_ERROR_NOMEM = 13
No memory available. The system cannot allocate
    more virtual memory because its capacity is full.

### FILE_ERROR_MFILE = 14
The current process has too many files open and
    can't open any more. Duplicate descriptors do count toward this
    limit.

### FILE_ERROR_NFILE = 15
There are too many distinct file openings in the
    entire system.

### FILE_ERROR_BADF = 16
Bad file descriptor; for example, I/O on a
    descriptor that has been closed or reading from a descriptor open
    only for writing (or vice versa).

### FILE_ERROR_INVAL = 17
Invalid argument. This is used to indicate
    various kinds of problems with passing the wrong argument to a
    library function.

### FILE_ERROR_PIPE = 18
Broken pipe; there is no process reading from the
    other end of a pipe. Every library function that returns this
    error code also generates a 'SIGPIPE' signal; this signal
    terminates the program if not handled or blocked. Thus, your
    program will never actually see this code unless it has handled
    or blocked 'SIGPIPE'.

### FILE_ERROR_AGAIN = 19
Resource temporarily unavailable; the call might
    work if you try again later.

### FILE_ERROR_INTR = 20
Interrupted function call; an asynchronous signal
    occurred and prevented completion of the call. When this
    happens, you should try the call again.

### FILE_ERROR_IO = 21
Input/output error; usually used for physical read
   or write errors. i.e. the disk or other physical device hardware
   is returning errors.

### FILE_ERROR_PERM = 22
Operation not permitted; only the owner of the
   file (or other resource) or processes with special privileges can
   perform the operation.

### FILE_ERROR_NOSYS = 23
Function not implemented; this indicates that
   the system is missing some functionality.

### FILE_ERROR_FAILED = 24
Does not correspond to a UNIX error code; this
   is the standard "failed for unspecified reason" error code present
   in all #GError error code enumerations. Returned if no specific
   code applies.


## `IOChannelError`

Error codes returned by #GIOChannel operations.

C - `GIOChannelError`

### IO_CHANNEL_ERROR_FBIG = 0
File too large.

### IO_CHANNEL_ERROR_INVAL = 1
Invalid argument.

### IO_CHANNEL_ERROR_IO = 2
IO error.

### IO_CHANNEL_ERROR_ISDIR = 3
File is a directory.

### IO_CHANNEL_ERROR_NOSPC = 4
No space left on device.

### IO_CHANNEL_ERROR_NXIO = 5
No such device or address.

### IO_CHANNEL_ERROR_OVERFLOW = 6
Value too large for defined datatype.

### IO_CHANNEL_ERROR_PIPE = 7
Broken pipe.

### IO_CHANNEL_ERROR_FAILED = 8
Some other error.


## `IOError`

#GIOError is only used by the deprecated functions
g_io_channel_read(), g_io_channel_write(), and g_io_channel_seek().

C - `GIOError`

### IO_ERROR_NONE = 0
no error

### IO_ERROR_AGAIN = 1
an EAGAIN error occurred

### IO_ERROR_INVAL = 2
an EINVAL error occurred

### IO_ERROR_UNKNOWN = 3
another error occurred


## `IOStatus`

Stati returned by most of the #GIOFuncs functions.

C - `GIOStatus`

### IO_STATUS_ERROR = 0
An error occurred.

### IO_STATUS_NORMAL = 1
Success.

### IO_STATUS_EOF = 2
End of file.

### IO_STATUS_AGAIN = 3
Resource temporarily unavailable.


## `KeyFileError`

Error codes returned by key file parsing.

C - `GKeyFileError`

### KEY_FILE_ERROR_UNKNOWN_ENCODING = 0
the text being parsed was in
    an unknown encoding

### KEY_FILE_ERROR_PARSE = 1
document was ill-formed

### KEY_FILE_ERROR_NOT_FOUND = 2
the file was not found

### KEY_FILE_ERROR_KEY_NOT_FOUND = 3
a requested key was not found

### KEY_FILE_ERROR_GROUP_NOT_FOUND = 4
a requested group was not found

### KEY_FILE_ERROR_INVALID_VALUE = 5
a value could not be parsed


## `LogWriterOutput`

Return values from #GLogWriterFuncs to indicate whether the given log entry
was successfully handled by the writer, or whether there was an error in
handling it (and hence a fallback writer should be used).

If a #GLogWriterFunc ignores a log entry, it should return
%G_LOG_WRITER_HANDLED.

C - `GLogWriterOutput`

### LOG_WRITER_HANDLED = 1
Log writer has handled the log entry.

### LOG_WRITER_UNHANDLED = 0
Log writer could not handle the log entry.


## `MarkupError`

Error codes returned by markup parsing.

C - `GMarkupError`

### MARKUP_ERROR_BAD_UTF8 = 0
text being parsed was not valid UTF-8

### MARKUP_ERROR_EMPTY = 1
document contained nothing, or only whitespace

### MARKUP_ERROR_PARSE = 2
document was ill-formed

### MARKUP_ERROR_UNKNOWN_ELEMENT = 3
error should be set by #GMarkupParser
    functions; element wasn't known

### MARKUP_ERROR_UNKNOWN_ATTRIBUTE = 4
error should be set by #GMarkupParser
    functions; attribute wasn't known

### MARKUP_ERROR_INVALID_CONTENT = 5
error should be set by #GMarkupParser
    functions; content was invalid

### MARKUP_ERROR_MISSING_ATTRIBUTE = 6
error should be set by #GMarkupParser
    functions; a required attribute was missing


## `NormalizeMode`

Defines how a Unicode string is transformed in a canonical
form, standardizing such issues as whether a character with
an accent is represented as a base character and combining
accent or as a single precomposed character. Unicode strings
should generally be normalized before comparing them.

C - `GNormalizeMode`

### NORMALIZE_DEFAULT = 0
standardize differences that do not affect the
    text content, such as the above-mentioned accent representation

### NORMALIZE_NFD = 0
another name for %G_NORMALIZE_DEFAULT

### NORMALIZE_DEFAULT_COMPOSE = 1
like %G_NORMALIZE_DEFAULT, but with
    composed forms rather than a maximally decomposed form

### NORMALIZE_NFC = 1
another name for %G_NORMALIZE_DEFAULT_COMPOSE

### NORMALIZE_ALL = 2
beyond %G_NORMALIZE_DEFAULT also standardize the
    "compatibility" characters in Unicode, such as SUPERSCRIPT THREE
    to the standard forms (in this case DIGIT THREE). Formatting
    information may be lost but for most text operations such
    characters should be considered the same

### NORMALIZE_NFKD = 2
another name for %G_NORMALIZE_ALL

### NORMALIZE_ALL_COMPOSE = 3
like %G_NORMALIZE_ALL, but with composed
    forms rather than a maximally decomposed form

### NORMALIZE_NFKC = 3
another name for %G_NORMALIZE_ALL_COMPOSE


## `NumberParserError`

Error codes returned by functions converting a string to a number.

C - `GNumberParserError`

### NUMBER_PARSER_ERROR_INVALID = 0
String was not a valid number.

### NUMBER_PARSER_ERROR_OUT_OF_BOUNDS = 1
String was a number, but out of bounds.


## `OnceStatus`

The possible statuses of a one-time initialization function
controlled by a #GOnce struct.

C - `GOnceStatus`

### ONCE_STATUS_NOTCALLED = 0
the function has not been called yet.

### ONCE_STATUS_PROGRESS = 1
the function call is currently in progress.

### ONCE_STATUS_READY = 2
the function has been called.


## `OptionArg`

The #GOptionArg enum values determine which type of extra argument the
options expect to find. If an option expects an extra argument, it can
be specified in several ways; with a short option: `-x arg`, with a long
option: `--name arg` or combined in a single argument: `--name=arg`.

C - `GOptionArg`

### OPTION_ARG_NONE = 0
No extra argument. This is useful for simple flags.

### OPTION_ARG_STRING = 1
The option takes a string argument.

### OPTION_ARG_INT = 2
The option takes an integer argument.

### OPTION_ARG_CALLBACK = 3
The option provides a callback (of type
    #GOptionArgFunc) to parse the extra argument.

### OPTION_ARG_FILENAME = 4
The option takes a filename as argument.

### OPTION_ARG_STRING_ARRAY = 5
The option takes a string argument, multiple
    uses of the option are collected into an array of strings.

### OPTION_ARG_FILENAME_ARRAY = 6
The option takes a filename as argument,
    multiple uses of the option are collected into an array of strings.

### OPTION_ARG_DOUBLE = 7
The option takes a double argument. The argument
    can be formatted either for the user's locale or for the "C" locale.
    Since 2.12

### OPTION_ARG_INT64 = 8
The option takes a 64-bit integer. Like
    %G_OPTION_ARG_INT but for larger numbers. The number can be in
    decimal base, or in hexadecimal (when prefixed with `0x`, for
    example, `0xffffffff`). Since 2.12


## `OptionError`

Error codes returned by option parsing.

C - `GOptionError`

### OPTION_ERROR_UNKNOWN_OPTION = 0
An option was not known to the parser.
 This error will only be reported, if the parser hasn't been instructed
 to ignore unknown options, see g_option_context_set_ignore_unknown_options().

### OPTION_ERROR_BAD_VALUE = 1
A value couldn't be parsed.

### OPTION_ERROR_FAILED = 2
A #GOptionArgFunc callback failed.


## `RegexError`

Error codes returned by regular expressions functions.

C - `GRegexError`

### REGEX_ERROR_COMPILE = 0
Compilation of the regular expression failed.

### REGEX_ERROR_OPTIMIZE = 1
Optimization of the regular expression failed.

### REGEX_ERROR_REPLACE = 2
Replacement failed due to an ill-formed replacement
    string.

### REGEX_ERROR_MATCH = 3
The match process failed.

### REGEX_ERROR_INTERNAL = 4
Internal error of the regular expression engine.
    Since 2.16

### REGEX_ERROR_STRAY_BACKSLASH = 101
"\\" at end of pattern. Since 2.16

### REGEX_ERROR_MISSING_CONTROL_CHAR = 102
"\\c" at end of pattern. Since 2.16

### REGEX_ERROR_UNRECOGNIZED_ESCAPE = 103
Unrecognized character follows "\\".
    Since 2.16

### REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER = 104
Numbers out of order in "{}"
    quantifier. Since 2.16

### REGEX_ERROR_QUANTIFIER_TOO_BIG = 105
Number too big in "{}" quantifier.
    Since 2.16

### REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS = 106
Missing terminating "]" for
    character class. Since 2.16

### REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS = 107
Invalid escape sequence
    in character class. Since 2.16

### REGEX_ERROR_RANGE_OUT_OF_ORDER = 108
Range out of order in character class.
    Since 2.16

### REGEX_ERROR_NOTHING_TO_REPEAT = 109
Nothing to repeat. Since 2.16

### REGEX_ERROR_UNRECOGNIZED_CHARACTER = 112
Unrecognized character after "(?",
    "(?<" or "(?P". Since 2.16

### REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS = 113
POSIX named classes are
    supported only within a class. Since 2.16

### REGEX_ERROR_UNMATCHED_PARENTHESIS = 114
Missing terminating ")" or ")"
    without opening "(". Since 2.16

### REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE = 115
Reference to non-existent
    subpattern. Since 2.16

### REGEX_ERROR_UNTERMINATED_COMMENT = 118
Missing terminating ")" after comment.
    Since 2.16

### REGEX_ERROR_EXPRESSION_TOO_LARGE = 120
Regular expression too large.
    Since 2.16

### REGEX_ERROR_MEMORY_ERROR = 121
Failed to get memory. Since 2.16

### REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND = 125
Lookbehind assertion is not
    fixed length. Since 2.16

### REGEX_ERROR_MALFORMED_CONDITION = 126
Malformed number or name after "(?(".
    Since 2.16

### REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES = 127
Conditional group contains
    more than two branches. Since 2.16

### REGEX_ERROR_ASSERTION_EXPECTED = 128
Assertion expected after "(?(".
    Since 2.16

### REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME = 130
Unknown POSIX class name.
    Since 2.16

### REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED = 131
POSIX collating
    elements are not supported. Since 2.16

### REGEX_ERROR_HEX_CODE_TOO_LARGE = 134
Character value in "\\x{...}" sequence
    is too large. Since 2.16

### REGEX_ERROR_INVALID_CONDITION = 135
Invalid condition "(?(0)". Since 2.16

### REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND = 136
\\C not allowed in
    lookbehind assertion. Since 2.16

### REGEX_ERROR_INFINITE_LOOP = 140
Recursive call could loop indefinitely.
    Since 2.16

### REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR = 142
Missing terminator
    in subpattern name. Since 2.16

### REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME = 143
Two named subpatterns have
    the same name. Since 2.16

### REGEX_ERROR_MALFORMED_PROPERTY = 146
Malformed "\\P" or "\\p" sequence.
    Since 2.16

### REGEX_ERROR_UNKNOWN_PROPERTY = 147
Unknown property name after "\\P" or
    "\\p". Since 2.16

### REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG = 148
Subpattern name is too long
    (maximum 32 characters). Since 2.16

### REGEX_ERROR_TOO_MANY_SUBPATTERNS = 149
Too many named subpatterns (maximum
    10,000). Since 2.16

### REGEX_ERROR_INVALID_OCTAL_VALUE = 151
Octal value is greater than "\\377".
    Since 2.16

### REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE = 154
"DEFINE" group contains more
    than one branch. Since 2.16

### REGEX_ERROR_DEFINE_REPETION = 155
Repeating a "DEFINE" group is not allowed.
    This error is never raised. Since: 2.16 Deprecated: 2.34

### REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS = 156
Inconsistent newline options.
    Since 2.16

### REGEX_ERROR_MISSING_BACK_REFERENCE = 157
"\\g" is not followed by a braced,
     angle-bracketed, or quoted name or number, or by a plain number. Since: 2.16

### REGEX_ERROR_INVALID_RELATIVE_REFERENCE = 158
relative reference must not be zero. Since: 2.34

### REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_FORBIDDEN = 159
the backtracing
    control verb used does not allow an argument. Since: 2.34

### REGEX_ERROR_UNKNOWN_BACKTRACKING_CONTROL_VERB = 160
unknown backtracing
    control verb. Since: 2.34

### REGEX_ERROR_NUMBER_TOO_BIG = 161
number is too big in escape sequence. Since: 2.34

### REGEX_ERROR_MISSING_SUBPATTERN_NAME = 162
Missing subpattern name. Since: 2.34

### REGEX_ERROR_MISSING_DIGIT = 163
Missing digit. Since 2.34

### REGEX_ERROR_INVALID_DATA_CHARACTER = 164
In JavaScript compatibility mode,
    "[" is an invalid data character. Since: 2.34

### REGEX_ERROR_EXTRA_SUBPATTERN_NAME = 165
different names for subpatterns of the
    same number are not allowed. Since: 2.34

### REGEX_ERROR_BACKTRACKING_CONTROL_VERB_ARGUMENT_REQUIRED = 166
the backtracing control
    verb requires an argument. Since: 2.34

### REGEX_ERROR_INVALID_CONTROL_CHAR = 168
"\\c" must be followed by an ASCII
    character. Since: 2.34

### REGEX_ERROR_MISSING_NAME = 169
"\\k" is not followed by a braced, angle-bracketed, or
    quoted name. Since: 2.34

### REGEX_ERROR_NOT_SUPPORTED_IN_CLASS = 171
"\\N" is not supported in a class. Since: 2.34

### REGEX_ERROR_TOO_MANY_FORWARD_REFERENCES = 172
too many forward references. Since: 2.34

### REGEX_ERROR_NAME_TOO_LONG = 175
the name is too long in "(*MARK)", "(*PRUNE)",
    "(*SKIP)", or "(*THEN)". Since: 2.34

### REGEX_ERROR_CHARACTER_VALUE_TOO_LARGE = 176
the character value in the \\u sequence is
    too large. Since: 2.34


## `SeekType`

An enumeration specifying the base position for a
g_io_channel_seek_position() operation.

C - `GSeekType`

### SEEK_CUR = 0
the current position in the file.

### SEEK_SET = 1
the start of the file.

### SEEK_END = 2
the end of the file.


## `ShellError`

Error codes returned by shell functions.

C - `GShellError`

### SHELL_ERROR_BAD_QUOTING = 0
Mismatched or otherwise mangled quoting.

### SHELL_ERROR_EMPTY_STRING = 1
String to be parsed was empty.

### SHELL_ERROR_FAILED = 2
Some other error.


## `SliceConfig`



C - `GSliceConfig`

### SLICE_CONFIG_ALWAYS_MALLOC = 1


### SLICE_CONFIG_BYPASS_MAGAZINES = 2


### SLICE_CONFIG_WORKING_SET_MSECS = 3


### SLICE_CONFIG_COLOR_INCREMENT = 4


### SLICE_CONFIG_CHUNK_SIZES = 5


### SLICE_CONFIG_CONTENTION_COUNTER = 6



## `SpawnError`

Error codes returned by spawning processes.

C - `GSpawnError`

### SPAWN_ERROR_FORK = 0
Fork failed due to lack of memory.

### SPAWN_ERROR_READ = 1
Read or select on pipes failed.

### SPAWN_ERROR_CHDIR = 2
Changing to working directory failed.

### SPAWN_ERROR_ACCES = 3
execv() returned `EACCES`

### SPAWN_ERROR_PERM = 4
execv() returned `EPERM`

### SPAWN_ERROR_TOO_BIG = 5
execv() returned `E2BIG`

### SPAWN_ERROR_2BIG = 5
deprecated alias for %G_SPAWN_ERROR_TOO_BIG

### SPAWN_ERROR_NOEXEC = 6
execv() returned `ENOEXEC`

### SPAWN_ERROR_NAMETOOLONG = 7
execv() returned `ENAMETOOLONG`

### SPAWN_ERROR_NOENT = 8
execv() returned `ENOENT`

### SPAWN_ERROR_NOMEM = 9
execv() returned `ENOMEM`

### SPAWN_ERROR_NOTDIR = 10
execv() returned `ENOTDIR`

### SPAWN_ERROR_LOOP = 11
execv() returned `ELOOP`

### SPAWN_ERROR_TXTBUSY = 12
execv() returned `ETXTBUSY`

### SPAWN_ERROR_IO = 13
execv() returned `EIO`

### SPAWN_ERROR_NFILE = 14
execv() returned `ENFILE`

### SPAWN_ERROR_MFILE = 15
execv() returned `EMFILE`

### SPAWN_ERROR_INVAL = 16
execv() returned `EINVAL`

### SPAWN_ERROR_ISDIR = 17
execv() returned `EISDIR`

### SPAWN_ERROR_LIBBAD = 18
execv() returned `ELIBBAD`

### SPAWN_ERROR_FAILED = 19
Some other fatal failure,
  `error->message` should explain.


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

### TEST_DIST = 0
a file that was included in the distribution tarball

### TEST_BUILT = 1
a file that was built on the compiling machine


## `TestLogType`



C - `GTestLogType`

### TEST_LOG_NONE = 0


### TEST_LOG_ERROR = 1


### TEST_LOG_START_BINARY = 2


### TEST_LOG_LIST_CASE = 3


### TEST_LOG_SKIP_CASE = 4


### TEST_LOG_START_CASE = 5


### TEST_LOG_STOP_CASE = 6


### TEST_LOG_MIN_RESULT = 7


### TEST_LOG_MAX_RESULT = 8


### TEST_LOG_MESSAGE = 9


### TEST_LOG_START_SUITE = 10


### TEST_LOG_STOP_SUITE = 11



## `ThreadError`

Possible errors of thread related functions.

C - `GThreadError`

### THREAD_ERROR_AGAIN = 0
a thread couldn't be created due to resource
                       shortage. Try again later.


## `TimeType`

Disambiguates a given time in two ways.

First, specifies if the given time is in universal or local time.

Second, if the time is in local time, specifies if it is local
standard time or local daylight time.  This is important for the case
where the same local time occurs twice (during daylight savings time
transitions, for example).

C - `GTimeType`

### TIME_TYPE_STANDARD = 0
the time is in local standard time

### TIME_TYPE_DAYLIGHT = 1
the time is in local daylight time

### TIME_TYPE_UNIVERSAL = 2
the time is in UTC


## `TokenType`

The possible types of token returned from each
g_scanner_get_next_token() call.

C - `GTokenType`

### TOKEN_EOF = 0
the end of the file

### TOKEN_LEFT_PAREN = 40
a '(' character

### TOKEN_RIGHT_PAREN = 41
a ')' character

### TOKEN_LEFT_CURLY = 123
a '{' character

### TOKEN_RIGHT_CURLY = 125
a '}' character

### TOKEN_LEFT_BRACE = 91
a '[' character

### TOKEN_RIGHT_BRACE = 93
a ']' character

### TOKEN_EQUAL_SIGN = 61
a '=' character

### TOKEN_COMMA = 44
a ',' character

### TOKEN_NONE = 256
not a token

### TOKEN_ERROR = 257
an error occurred

### TOKEN_CHAR = 258
a character

### TOKEN_BINARY = 259
a binary integer

### TOKEN_OCTAL = 260
an octal integer

### TOKEN_INT = 261
an integer

### TOKEN_HEX = 262
a hex integer

### TOKEN_FLOAT = 263
a floating point number

### TOKEN_STRING = 264
a string

### TOKEN_SYMBOL = 265
a symbol

### TOKEN_IDENTIFIER = 266
an identifier

### TOKEN_IDENTIFIER_NULL = 267
a null identifier

### TOKEN_COMMENT_SINGLE = 268
one line comment

### TOKEN_COMMENT_MULTI = 269
multi line comment


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

### IN_ORDER = 0
vists a node's left child first, then the node itself,
             then its right child. This is the one to use if you
             want the output sorted according to the compare
             function.

### PRE_ORDER = 1
visits a node, then its children.

### POST_ORDER = 2
visits the node's children, then the node itself.

### LEVEL_ORDER = 3
is not implemented for
             [balanced binary trees][glib-Balanced-Binary-Trees].
             For [n-ary trees][glib-N-ary-Trees], it
             vists the root node first, then its children, then
             its grandchildren, and so on. Note that this is less
             efficient than the other orders.


## `UnicodeBreakType`

These are the possible line break classifications.

Since new unicode versions may add new types here, applications should be ready
to handle unknown values. They may be regarded as %G_UNICODE_BREAK_UNKNOWN.

See [Unicode Line Breaking Algorithm](http://www.unicode.org/unicode/reports/tr14/).

C - `GUnicodeBreakType`

### UNICODE_BREAK_MANDATORY = 0
Mandatory Break (BK)

### UNICODE_BREAK_CARRIAGE_RETURN = 1
Carriage Return (CR)

### UNICODE_BREAK_LINE_FEED = 2
Line Feed (LF)

### UNICODE_BREAK_COMBINING_MARK = 3
Attached Characters and Combining Marks (CM)

### UNICODE_BREAK_SURROGATE = 4
Surrogates (SG)

### UNICODE_BREAK_ZERO_WIDTH_SPACE = 5
Zero Width Space (ZW)

### UNICODE_BREAK_INSEPARABLE = 6
Inseparable (IN)

### UNICODE_BREAK_NON_BREAKING_GLUE = 7
Non-breaking ("Glue") (GL)

### UNICODE_BREAK_CONTINGENT = 8
Contingent Break Opportunity (CB)

### UNICODE_BREAK_SPACE = 9
Space (SP)

### UNICODE_BREAK_AFTER = 10
Break Opportunity After (BA)

### UNICODE_BREAK_BEFORE = 11
Break Opportunity Before (BB)

### UNICODE_BREAK_BEFORE_AND_AFTER = 12
Break Opportunity Before and After (B2)

### UNICODE_BREAK_HYPHEN = 13
Hyphen (HY)

### UNICODE_BREAK_NON_STARTER = 14
Nonstarter (NS)

### UNICODE_BREAK_OPEN_PUNCTUATION = 15
Opening Punctuation (OP)

### UNICODE_BREAK_CLOSE_PUNCTUATION = 16
Closing Punctuation (CL)

### UNICODE_BREAK_QUOTATION = 17
Ambiguous Quotation (QU)

### UNICODE_BREAK_EXCLAMATION = 18
Exclamation/Interrogation (EX)

### UNICODE_BREAK_IDEOGRAPHIC = 19
Ideographic (ID)

### UNICODE_BREAK_NUMERIC = 20
Numeric (NU)

### UNICODE_BREAK_INFIX_SEPARATOR = 21
Infix Separator (Numeric) (IS)

### UNICODE_BREAK_SYMBOL = 22
Symbols Allowing Break After (SY)

### UNICODE_BREAK_ALPHABETIC = 23
Ordinary Alphabetic and Symbol Characters (AL)

### UNICODE_BREAK_PREFIX = 24
Prefix (Numeric) (PR)

### UNICODE_BREAK_POSTFIX = 25
Postfix (Numeric) (PO)

### UNICODE_BREAK_COMPLEX_CONTEXT = 26
Complex Content Dependent (South East Asian) (SA)

### UNICODE_BREAK_AMBIGUOUS = 27
Ambiguous (Alphabetic or Ideographic) (AI)

### UNICODE_BREAK_UNKNOWN = 28
Unknown (XX)

### UNICODE_BREAK_NEXT_LINE = 29
Next Line (NL)

### UNICODE_BREAK_WORD_JOINER = 30
Word Joiner (WJ)

### UNICODE_BREAK_HANGUL_L_JAMO = 31
Hangul L Jamo (JL)

### UNICODE_BREAK_HANGUL_V_JAMO = 32
Hangul V Jamo (JV)

### UNICODE_BREAK_HANGUL_T_JAMO = 33
Hangul T Jamo (JT)

### UNICODE_BREAK_HANGUL_LV_SYLLABLE = 34
Hangul LV Syllable (H2)

### UNICODE_BREAK_HANGUL_LVT_SYLLABLE = 35
Hangul LVT Syllable (H3)

### UNICODE_BREAK_CLOSE_PARANTHESIS = 36
Closing Parenthesis (CP). Since 2.28

### UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER = 37
Conditional Japanese Starter (CJ). Since: 2.32

### UNICODE_BREAK_HEBREW_LETTER = 38
Hebrew Letter (HL). Since: 2.32

### UNICODE_BREAK_REGIONAL_INDICATOR = 39
Regional Indicator (RI). Since: 2.36

### UNICODE_BREAK_EMOJI_BASE = 40
Emoji Base (EB). Since: 2.50

### UNICODE_BREAK_EMOJI_MODIFIER = 41
Emoji Modifier (EM). Since: 2.50

### UNICODE_BREAK_ZERO_WIDTH_JOINER = 42
Zero Width Joiner (ZWJ). Since: 2.50


## `UnicodeScript`

The #GUnicodeScript enumeration identifies different writing
systems. The values correspond to the names as defined in the
Unicode standard. The enumeration has been added in GLib 2.14,
and is interchangeable with #PangoScript.

Note that new types may be added in the future. Applications
should be ready to handle unknown values.
See [Unicode Standard Annex #24: Script names](http://www.unicode.org/reports/tr24/).

C - `GUnicodeScript`

### UNICODE_SCRIPT_INVALID_CODE = -1
a value never returned from g_unichar_get_script()

### UNICODE_SCRIPT_COMMON = 0
a character used by multiple different scripts

### UNICODE_SCRIPT_INHERITED = 1
a mark glyph that takes its script from the
                              base glyph to which it is attached

### UNICODE_SCRIPT_ARABIC = 2
Arabic

### UNICODE_SCRIPT_ARMENIAN = 3
Armenian

### UNICODE_SCRIPT_BENGALI = 4
Bengali

### UNICODE_SCRIPT_BOPOMOFO = 5
Bopomofo

### UNICODE_SCRIPT_CHEROKEE = 6
Cherokee

### UNICODE_SCRIPT_COPTIC = 7
Coptic

### UNICODE_SCRIPT_CYRILLIC = 8
Cyrillic

### UNICODE_SCRIPT_DESERET = 9
Deseret

### UNICODE_SCRIPT_DEVANAGARI = 10
Devanagari

### UNICODE_SCRIPT_ETHIOPIC = 11
Ethiopic

### UNICODE_SCRIPT_GEORGIAN = 12
Georgian

### UNICODE_SCRIPT_GOTHIC = 13
Gothic

### UNICODE_SCRIPT_GREEK = 14
Greek

### UNICODE_SCRIPT_GUJARATI = 15
Gujarati

### UNICODE_SCRIPT_GURMUKHI = 16
Gurmukhi

### UNICODE_SCRIPT_HAN = 17
Han

### UNICODE_SCRIPT_HANGUL = 18
Hangul

### UNICODE_SCRIPT_HEBREW = 19
Hebrew

### UNICODE_SCRIPT_HIRAGANA = 20
Hiragana

### UNICODE_SCRIPT_KANNADA = 21
Kannada

### UNICODE_SCRIPT_KATAKANA = 22
Katakana

### UNICODE_SCRIPT_KHMER = 23
Khmer

### UNICODE_SCRIPT_LAO = 24
Lao

### UNICODE_SCRIPT_LATIN = 25
Latin

### UNICODE_SCRIPT_MALAYALAM = 26
Malayalam

### UNICODE_SCRIPT_MONGOLIAN = 27
Mongolian

### UNICODE_SCRIPT_MYANMAR = 28
Myanmar

### UNICODE_SCRIPT_OGHAM = 29
Ogham

### UNICODE_SCRIPT_OLD_ITALIC = 30
Old Italic

### UNICODE_SCRIPT_ORIYA = 31
Oriya

### UNICODE_SCRIPT_RUNIC = 32
Runic

### UNICODE_SCRIPT_SINHALA = 33
Sinhala

### UNICODE_SCRIPT_SYRIAC = 34
Syriac

### UNICODE_SCRIPT_TAMIL = 35
Tamil

### UNICODE_SCRIPT_TELUGU = 36
Telugu

### UNICODE_SCRIPT_THAANA = 37
Thaana

### UNICODE_SCRIPT_THAI = 38
Thai

### UNICODE_SCRIPT_TIBETAN = 39
Tibetan

### UNICODE_SCRIPT_CANADIAN_ABORIGINAL = 40
Canadian Aboriginal

### UNICODE_SCRIPT_YI = 41
Yi

### UNICODE_SCRIPT_TAGALOG = 42
Tagalog

### UNICODE_SCRIPT_HANUNOO = 43
Hanunoo

### UNICODE_SCRIPT_BUHID = 44
Buhid

### UNICODE_SCRIPT_TAGBANWA = 45
Tagbanwa

### UNICODE_SCRIPT_BRAILLE = 46
Braille

### UNICODE_SCRIPT_CYPRIOT = 47
Cypriot

### UNICODE_SCRIPT_LIMBU = 48
Limbu

### UNICODE_SCRIPT_OSMANYA = 49
Osmanya

### UNICODE_SCRIPT_SHAVIAN = 50
Shavian

### UNICODE_SCRIPT_LINEAR_B = 51
Linear B

### UNICODE_SCRIPT_TAI_LE = 52
Tai Le

### UNICODE_SCRIPT_UGARITIC = 53
Ugaritic

### UNICODE_SCRIPT_NEW_TAI_LUE = 54
New Tai Lue

### UNICODE_SCRIPT_BUGINESE = 55
Buginese

### UNICODE_SCRIPT_GLAGOLITIC = 56
Glagolitic

### UNICODE_SCRIPT_TIFINAGH = 57
Tifinagh

### UNICODE_SCRIPT_SYLOTI_NAGRI = 58
Syloti Nagri

### UNICODE_SCRIPT_OLD_PERSIAN = 59
Old Persian

### UNICODE_SCRIPT_KHAROSHTHI = 60
Kharoshthi

### UNICODE_SCRIPT_UNKNOWN = 61
an unassigned code point

### UNICODE_SCRIPT_BALINESE = 62
Balinese

### UNICODE_SCRIPT_CUNEIFORM = 63
Cuneiform

### UNICODE_SCRIPT_PHOENICIAN = 64
Phoenician

### UNICODE_SCRIPT_PHAGS_PA = 65
Phags-pa

### UNICODE_SCRIPT_NKO = 66
N'Ko

### UNICODE_SCRIPT_KAYAH_LI = 67
Kayah Li. Since 2.16.3

### UNICODE_SCRIPT_LEPCHA = 68
Lepcha. Since 2.16.3

### UNICODE_SCRIPT_REJANG = 69
Rejang. Since 2.16.3

### UNICODE_SCRIPT_SUNDANESE = 70
Sundanese. Since 2.16.3

### UNICODE_SCRIPT_SAURASHTRA = 71
Saurashtra. Since 2.16.3

### UNICODE_SCRIPT_CHAM = 72
Cham. Since 2.16.3

### UNICODE_SCRIPT_OL_CHIKI = 73
Ol Chiki. Since 2.16.3

### UNICODE_SCRIPT_VAI = 74
Vai. Since 2.16.3

### UNICODE_SCRIPT_CARIAN = 75
Carian. Since 2.16.3

### UNICODE_SCRIPT_LYCIAN = 76
Lycian. Since 2.16.3

### UNICODE_SCRIPT_LYDIAN = 77
Lydian. Since 2.16.3

### UNICODE_SCRIPT_AVESTAN = 78
Avestan. Since 2.26

### UNICODE_SCRIPT_BAMUM = 79
Bamum. Since 2.26

### UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS = 80
Egyptian Hieroglpyhs. Since 2.26

### UNICODE_SCRIPT_IMPERIAL_ARAMAIC = 81
Imperial Aramaic. Since 2.26

### UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI = 82
Inscriptional Pahlavi. Since 2.26

### UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN = 83
Inscriptional Parthian. Since 2.26

### UNICODE_SCRIPT_JAVANESE = 84
Javanese. Since 2.26

### UNICODE_SCRIPT_KAITHI = 85
Kaithi. Since 2.26

### UNICODE_SCRIPT_LISU = 86
Lisu. Since 2.26

### UNICODE_SCRIPT_MEETEI_MAYEK = 87
Meetei Mayek. Since 2.26

### UNICODE_SCRIPT_OLD_SOUTH_ARABIAN = 88
Old South Arabian. Since 2.26

### UNICODE_SCRIPT_OLD_TURKIC = 89
Old Turkic. Since 2.28

### UNICODE_SCRIPT_SAMARITAN = 90
Samaritan. Since 2.26

### UNICODE_SCRIPT_TAI_THAM = 91
Tai Tham. Since 2.26

### UNICODE_SCRIPT_TAI_VIET = 92
Tai Viet. Since 2.26

### UNICODE_SCRIPT_BATAK = 93
Batak. Since 2.28

### UNICODE_SCRIPT_BRAHMI = 94
Brahmi. Since 2.28

### UNICODE_SCRIPT_MANDAIC = 95
Mandaic. Since 2.28

### UNICODE_SCRIPT_CHAKMA = 96
Chakma. Since: 2.32

### UNICODE_SCRIPT_MEROITIC_CURSIVE = 97
Meroitic Cursive. Since: 2.32

### UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS = 98
Meroitic Hieroglyphs. Since: 2.32

### UNICODE_SCRIPT_MIAO = 99
Miao. Since: 2.32

### UNICODE_SCRIPT_SHARADA = 100
Sharada. Since: 2.32

### UNICODE_SCRIPT_SORA_SOMPENG = 101
Sora Sompeng. Since: 2.32

### UNICODE_SCRIPT_TAKRI = 102
Takri. Since: 2.32

### UNICODE_SCRIPT_BASSA_VAH = 103
Bassa. Since: 2.42

### UNICODE_SCRIPT_CAUCASIAN_ALBANIAN = 104
Caucasian Albanian. Since: 2.42

### UNICODE_SCRIPT_DUPLOYAN = 105
Duployan. Since: 2.42

### UNICODE_SCRIPT_ELBASAN = 106
Elbasan. Since: 2.42

### UNICODE_SCRIPT_GRANTHA = 107
Grantha. Since: 2.42

### UNICODE_SCRIPT_KHOJKI = 108
Kjohki. Since: 2.42

### UNICODE_SCRIPT_KHUDAWADI = 109
Khudawadi, Sindhi. Since: 2.42

### UNICODE_SCRIPT_LINEAR_A = 110
Linear A. Since: 2.42

### UNICODE_SCRIPT_MAHAJANI = 111
Mahajani. Since: 2.42

### UNICODE_SCRIPT_MANICHAEAN = 112
Manichaean. Since: 2.42

### UNICODE_SCRIPT_MENDE_KIKAKUI = 113
Mende Kikakui. Since: 2.42

### UNICODE_SCRIPT_MODI = 114
Modi. Since: 2.42

### UNICODE_SCRIPT_MRO = 115
Mro. Since: 2.42

### UNICODE_SCRIPT_NABATAEAN = 116
Nabataean. Since: 2.42

### UNICODE_SCRIPT_OLD_NORTH_ARABIAN = 117
Old North Arabian. Since: 2.42

### UNICODE_SCRIPT_OLD_PERMIC = 118
Old Permic. Since: 2.42

### UNICODE_SCRIPT_PAHAWH_HMONG = 119
Pahawh Hmong. Since: 2.42

### UNICODE_SCRIPT_PALMYRENE = 120
Palmyrene. Since: 2.42

### UNICODE_SCRIPT_PAU_CIN_HAU = 121
Pau Cin Hau. Since: 2.42

### UNICODE_SCRIPT_PSALTER_PAHLAVI = 122
Psalter Pahlavi. Since: 2.42

### UNICODE_SCRIPT_SIDDHAM = 123
Siddham. Since: 2.42

### UNICODE_SCRIPT_TIRHUTA = 124
Tirhuta. Since: 2.42

### UNICODE_SCRIPT_WARANG_CITI = 125
Warang Citi. Since: 2.42

### UNICODE_SCRIPT_AHOM = 126
Ahom. Since: 2.48

### UNICODE_SCRIPT_ANATOLIAN_HIEROGLYPHS = 127
Anatolian Hieroglyphs. Since: 2.48

### UNICODE_SCRIPT_HATRAN = 128
Hatran. Since: 2.48

### UNICODE_SCRIPT_MULTANI = 129
Multani. Since: 2.48

### UNICODE_SCRIPT_OLD_HUNGARIAN = 130
Old Hungarian. Since: 2.48

### UNICODE_SCRIPT_SIGNWRITING = 131
Signwriting. Since: 2.48

### UNICODE_SCRIPT_ADLAM = 132
Adlam. Since: 2.50

### UNICODE_SCRIPT_BHAIKSUKI = 133
Bhaiksuki. Since: 2.50

### UNICODE_SCRIPT_MARCHEN = 134
Marchen. Since: 2.50

### UNICODE_SCRIPT_NEWA = 135
Newa. Since: 2.50

### UNICODE_SCRIPT_OSAGE = 136
Osage. Since: 2.50

### UNICODE_SCRIPT_TANGUT = 137
Tangut. Since: 2.50

### UNICODE_SCRIPT_MASARAM_GONDI = 138
Masaram Gondi. Since: 2.54

### UNICODE_SCRIPT_NUSHU = 139
Nushu. Since: 2.54

### UNICODE_SCRIPT_SOYOMBO = 140
Soyombo. Since: 2.54

### UNICODE_SCRIPT_ZANABAZAR_SQUARE = 141
Zanabazar Square. Since: 2.54


## `UnicodeType`

These are the possible character classifications from the
Unicode specification.
See [Unicode Character Database](http://www.unicode.org/reports/tr44/#General_Category_Values).

C - `GUnicodeType`

### UNICODE_CONTROL = 0
General category "Other, Control" (Cc)

### UNICODE_FORMAT = 1
General category "Other, Format" (Cf)

### UNICODE_UNASSIGNED = 2
General category "Other, Not Assigned" (Cn)

### UNICODE_PRIVATE_USE = 3
General category "Other, Private Use" (Co)

### UNICODE_SURROGATE = 4
General category "Other, Surrogate" (Cs)

### UNICODE_LOWERCASE_LETTER = 5
General category "Letter, Lowercase" (Ll)

### UNICODE_MODIFIER_LETTER = 6
General category "Letter, Modifier" (Lm)

### UNICODE_OTHER_LETTER = 7
General category "Letter, Other" (Lo)

### UNICODE_TITLECASE_LETTER = 8
General category "Letter, Titlecase" (Lt)

### UNICODE_UPPERCASE_LETTER = 9
General category "Letter, Uppercase" (Lu)

### UNICODE_SPACING_MARK = 10
General category "Mark, Spacing" (Mc)

### UNICODE_ENCLOSING_MARK = 11
General category "Mark, Enclosing" (Me)

### UNICODE_NON_SPACING_MARK = 12
General category "Mark, Nonspacing" (Mn)

### UNICODE_DECIMAL_NUMBER = 13
General category "Number, Decimal Digit" (Nd)

### UNICODE_LETTER_NUMBER = 14
General category "Number, Letter" (Nl)

### UNICODE_OTHER_NUMBER = 15
General category "Number, Other" (No)

### UNICODE_CONNECT_PUNCTUATION = 16
General category "Punctuation, Connector" (Pc)

### UNICODE_DASH_PUNCTUATION = 17
General category "Punctuation, Dash" (Pd)

### UNICODE_CLOSE_PUNCTUATION = 18
General category "Punctuation, Close" (Pe)

### UNICODE_FINAL_PUNCTUATION = 19
General category "Punctuation, Final quote" (Pf)

### UNICODE_INITIAL_PUNCTUATION = 20
General category "Punctuation, Initial quote" (Pi)

### UNICODE_OTHER_PUNCTUATION = 21
General category "Punctuation, Other" (Po)

### UNICODE_OPEN_PUNCTUATION = 22
General category "Punctuation, Open" (Ps)

### UNICODE_CURRENCY_SYMBOL = 23
General category "Symbol, Currency" (Sc)

### UNICODE_MODIFIER_SYMBOL = 24
General category "Symbol, Modifier" (Sk)

### UNICODE_MATH_SYMBOL = 25
General category "Symbol, Math" (Sm)

### UNICODE_OTHER_SYMBOL = 26
General category "Symbol, Other" (So)

### UNICODE_LINE_SEPARATOR = 27
General category "Separator, Line" (Zl)

### UNICODE_PARAGRAPH_SEPARATOR = 28
General category "Separator, Paragraph" (Zp)

### UNICODE_SPACE_SEPARATOR = 29
General category "Separator, Space" (Zs)


## `UserDirectory`

These are logical ids for special directories which are defined
depending on the platform used. You should use g_get_user_special_dir()
to retrieve the full path associated to the logical id.

The #GUserDirectory enumeration can be extended at later date. Not
every platform has a directory for every logical id in this
enumeration.

C - `GUserDirectory`

### USER_DIRECTORY_DESKTOP = 0
the user's Desktop directory

### USER_DIRECTORY_DOCUMENTS = 1
the user's Documents directory

### USER_DIRECTORY_DOWNLOAD = 2
the user's Downloads directory

### USER_DIRECTORY_MUSIC = 3
the user's Music directory

### USER_DIRECTORY_PICTURES = 4
the user's Pictures directory

### USER_DIRECTORY_PUBLIC_SHARE = 5
the user's shared directory

### USER_DIRECTORY_TEMPLATES = 6
the user's Templates directory

### USER_DIRECTORY_VIDEOS = 7
the user's Movies directory

### USER_N_DIRECTORIES = 8
the number of enum values


## `VariantClass`

The range of possible top-level types of #GVariant instances.

C - `GVariantClass`

### VARIANT_CLASS_BOOLEAN = 98
The #GVariant is a boolean.

### VARIANT_CLASS_BYTE = 121
The #GVariant is a byte.

### VARIANT_CLASS_INT16 = 110
The #GVariant is a signed 16 bit integer.

### VARIANT_CLASS_UINT16 = 113
The #GVariant is an unsigned 16 bit integer.

### VARIANT_CLASS_INT32 = 105
The #GVariant is a signed 32 bit integer.

### VARIANT_CLASS_UINT32 = 117
The #GVariant is an unsigned 32 bit integer.

### VARIANT_CLASS_INT64 = 120
The #GVariant is a signed 64 bit integer.

### VARIANT_CLASS_UINT64 = 116
The #GVariant is an unsigned 64 bit integer.

### VARIANT_CLASS_HANDLE = 104
The #GVariant is a file handle index.

### VARIANT_CLASS_DOUBLE = 100
The #GVariant is a double precision floating
                         point value.

### VARIANT_CLASS_STRING = 115
The #GVariant is a normal string.

### VARIANT_CLASS_OBJECT_PATH = 111
The #GVariant is a D-Bus object path
                              string.

### VARIANT_CLASS_SIGNATURE = 103
The #GVariant is a D-Bus signature string.

### VARIANT_CLASS_VARIANT = 118
The #GVariant is a variant.

### VARIANT_CLASS_MAYBE = 109
The #GVariant is a maybe-typed value.

### VARIANT_CLASS_ARRAY = 97
The #GVariant is an array.

### VARIANT_CLASS_TUPLE = 40
The #GVariant is a tuple.

### VARIANT_CLASS_DICT_ENTRY = 123
The #GVariant is a dictionary entry.


## `VariantParseError`

Error codes returned by parsing text-format GVariants.

C - `GVariantParseError`

### VARIANT_PARSE_ERROR_FAILED = 0
generic error (unused)

### VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED = 1
a non-basic #GVariantType was given where a basic type was expected

### VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE = 2
cannot infer the #GVariantType

### VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED = 3
an indefinite #GVariantType was given where a definite type was expected

### VARIANT_PARSE_ERROR_INPUT_NOT_AT_END = 4
extra data after parsing finished

### VARIANT_PARSE_ERROR_INVALID_CHARACTER = 5
invalid character in number or unicode escape

### VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING = 6
not a valid #GVariant format string

### VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH = 7
not a valid object path

### VARIANT_PARSE_ERROR_INVALID_SIGNATURE = 8
not a valid type signature

### VARIANT_PARSE_ERROR_INVALID_TYPE_STRING = 9
not a valid #GVariant type string

### VARIANT_PARSE_ERROR_NO_COMMON_TYPE = 10
could not find a common type for array entries

### VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE = 11
the numerical value is out of range of the given type

### VARIANT_PARSE_ERROR_NUMBER_TOO_BIG = 12
the numerical value is out of range for any type

### VARIANT_PARSE_ERROR_TYPE_ERROR = 13
cannot parse as variant of the specified type

### VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN = 14
an unexpected token was encountered

### VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD = 15
an unknown keyword was encountered

### VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT = 16
unterminated string constant

### VARIANT_PARSE_ERROR_VALUE_EXPECTED = 17
no value given



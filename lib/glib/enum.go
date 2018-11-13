// This is a generated file - DO NOT EDIT

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Error codes returned by bookmark file parsing.
type BookmarkFileError C.GBookmarkFileError

const (
	// URI was ill-formed
	BOOKMARK_FILE_ERROR_INVALID_URI BookmarkFileError = 0

	// a requested field was not found
	BOOKMARK_FILE_ERROR_INVALID_VALUE BookmarkFileError = 1

	// a requested application did
	// not register a bookmark
	BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED BookmarkFileError = 2

	// a requested URI was not found
	BOOKMARK_FILE_ERROR_URI_NOT_FOUND BookmarkFileError = 3

	// document was ill formed
	BOOKMARK_FILE_ERROR_READ BookmarkFileError = 4

	// the text being parsed was
	// in an unknown encoding
	BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING BookmarkFileError = 5

	// an error occurred while writing
	BOOKMARK_FILE_ERROR_WRITE BookmarkFileError = 6

	// requested file was not found
	BOOKMARK_FILE_ERROR_FILE_NOT_FOUND BookmarkFileError = 7
)

// Error codes returned by character set conversion routines.
type ConvertError C.GConvertError

const (
	// Conversion between the requested character
	// sets is not supported.
	CONVERT_ERROR_NO_CONVERSION ConvertError = 0

	// Invalid byte sequence in conversion input;
	// or the character sequence could not be represented in the target
	// character set.
	CONVERT_ERROR_ILLEGAL_SEQUENCE ConvertError = 1

	// Conversion failed for some reason.
	CONVERT_ERROR_FAILED ConvertError = 2

	// Partial character sequence at end of input.
	CONVERT_ERROR_PARTIAL_INPUT ConvertError = 3

	// URI is invalid.
	CONVERT_ERROR_BAD_URI ConvertError = 4

	// Pathname is not an absolute path.
	CONVERT_ERROR_NOT_ABSOLUTE_PATH ConvertError = 5

	// No memory available. Since: 2.40
	CONVERT_ERROR_NO_MEMORY ConvertError = 6

	// An embedded NUL character is present in
	// conversion output where a NUL-terminated string is expected.
	// Since: 2.56
	CONVERT_ERROR_EMBEDDED_NUL ConvertError = 7
)

// This enumeration isn't used in the API, but may be useful if you need
// to mark a number as a day, month, or year.
type DateDMY C.GDateDMY

const (
	// a day
	DATE_DAY DateDMY = 0

	// a month
	DATE_MONTH DateDMY = 1

	// a year
	DATE_YEAR DateDMY = 2
)

// Enumeration representing a month; values are #G_DATE_JANUARY,
// #G_DATE_FEBRUARY, etc. #G_DATE_BAD_MONTH is the invalid value.
type DateMonth C.GDateMonth

const (
	// invalid value
	DATE_BAD_MONTH DateMonth = 0

	// January
	DATE_JANUARY DateMonth = 1

	// February
	DATE_FEBRUARY DateMonth = 2

	// March
	DATE_MARCH DateMonth = 3

	// April
	DATE_APRIL DateMonth = 4

	// May
	DATE_MAY DateMonth = 5

	// June
	DATE_JUNE DateMonth = 6

	// July
	DATE_JULY DateMonth = 7

	// August
	DATE_AUGUST DateMonth = 8

	// September
	DATE_SEPTEMBER DateMonth = 9

	// October
	DATE_OCTOBER DateMonth = 10

	// November
	DATE_NOVEMBER DateMonth = 11

	// December
	DATE_DECEMBER DateMonth = 12
)

// Enumeration representing a day of the week; #G_DATE_MONDAY,
// #G_DATE_TUESDAY, etc. #G_DATE_BAD_WEEKDAY is an invalid weekday.
type DateWeekday C.GDateWeekday

const (
	// invalid value
	DATE_BAD_WEEKDAY DateWeekday = 0

	// Monday
	DATE_MONDAY DateWeekday = 1

	// Tuesday
	DATE_TUESDAY DateWeekday = 2

	// Wednesday
	DATE_WEDNESDAY DateWeekday = 3

	// Thursday
	DATE_THURSDAY DateWeekday = 4

	// Friday
	DATE_FRIDAY DateWeekday = 5

	// Saturday
	DATE_SATURDAY DateWeekday = 6

	// Sunday
	DATE_SUNDAY DateWeekday = 7
)

// The possible errors, used in the @v_error field
// of #GTokenValue, when the token is a %G_TOKEN_ERROR.
type ErrorType C.GErrorType

const (
	// unknown error
	ERR_UNKNOWN ErrorType = 0

	// unexpected end of file
	ERR_UNEXP_EOF ErrorType = 1

	// unterminated string constant
	ERR_UNEXP_EOF_IN_STRING ErrorType = 2

	// unterminated comment
	ERR_UNEXP_EOF_IN_COMMENT ErrorType = 3

	// non-digit character in a number
	ERR_NON_DIGIT_IN_CONST ErrorType = 4

	// digit beyond radix in a number
	ERR_DIGIT_RADIX ErrorType = 5

	// non-decimal floating point number
	ERR_FLOAT_RADIX ErrorType = 6

	// malformed floating point number
	ERR_FLOAT_MALFORMED ErrorType = 7
)

// Values corresponding to @errno codes returned from file operations
// on UNIX. Unlike @errno codes, GFileError values are available on
// all systems, even Windows. The exact meaning of each code depends
// on what sort of file operation you were performing; the UNIX
// documentation gives more details. The following error code descriptions
// come from the GNU C Library manual, and are under the copyright
// of that manual.
//
// It's not very portable to make detailed assumptions about exactly
// which errors will be returned from a given operation. Some errors
// don't occur on some systems, etc., sometimes there are subtle
// differences in when a system will report a given error, etc.
type FileError C.GFileError

const (
	// Operation not permitted; only the owner of
	// the file (or other resource) or processes with special privileges
	// can perform the operation.
	FILE_ERROR_EXIST FileError = 0

	// File is a directory; you cannot open a directory
	// for writing, or create or remove hard links to it.
	FILE_ERROR_ISDIR FileError = 1

	// Permission denied; the file permissions do not
	// allow the attempted operation.
	FILE_ERROR_ACCES FileError = 2

	// Filename too long.
	FILE_ERROR_NAMETOOLONG FileError = 3

	// No such file or directory. This is a "file
	// doesn't exist" error for ordinary files that are referenced in
	// contexts where they are expected to already exist.
	FILE_ERROR_NOENT FileError = 4

	// A file that isn't a directory was specified when
	// a directory is required.
	FILE_ERROR_NOTDIR FileError = 5

	// No such device or address. The system tried to
	// use the device represented by a file you specified, and it
	// couldn't find the device. This can mean that the device file was
	// installed incorrectly, or that the physical device is missing or
	// not correctly attached to the computer.
	FILE_ERROR_NXIO FileError = 6

	// The underlying file system of the specified file
	// does not support memory mapping.
	FILE_ERROR_NODEV FileError = 7

	// The directory containing the new link can't be
	// modified because it's on a read-only file system.
	FILE_ERROR_ROFS FileError = 8

	// Text file busy.
	FILE_ERROR_TXTBSY FileError = 9

	// You passed in a pointer to bad memory.
	// (GLib won't reliably return this, don't pass in pointers to bad
	// memory.)
	FILE_ERROR_FAULT FileError = 10

	// Too many levels of symbolic links were encountered
	// in looking up a file name. This often indicates a cycle of symbolic
	// links.
	FILE_ERROR_LOOP FileError = 11

	// No space left on device; write operation on a
	// file failed because the disk is full.
	FILE_ERROR_NOSPC FileError = 12

	// No memory available. The system cannot allocate
	// more virtual memory because its capacity is full.
	FILE_ERROR_NOMEM FileError = 13

	// The current process has too many files open and
	// can't open any more. Duplicate descriptors do count toward this
	// limit.
	FILE_ERROR_MFILE FileError = 14

	// There are too many distinct file openings in the
	// entire system.
	FILE_ERROR_NFILE FileError = 15

	// Bad file descriptor; for example, I/O on a
	// descriptor that has been closed or reading from a descriptor open
	// only for writing (or vice versa).
	FILE_ERROR_BADF FileError = 16

	// Invalid argument. This is used to indicate
	// various kinds of problems with passing the wrong argument to a
	// library function.
	FILE_ERROR_INVAL FileError = 17

	// Broken pipe; there is no process reading from the
	// other end of a pipe. Every library function that returns this
	// error code also generates a 'SIGPIPE' signal; this signal
	// terminates the program if not handled or blocked. Thus, your
	// program will never actually see this code unless it has handled
	// or blocked 'SIGPIPE'.
	FILE_ERROR_PIPE FileError = 18

	// Resource temporarily unavailable; the call might
	// work if you try again later.
	FILE_ERROR_AGAIN FileError = 19

	// Interrupted function call; an asynchronous signal
	// occurred and prevented completion of the call. When this
	// happens, you should try the call again.
	FILE_ERROR_INTR FileError = 20

	// Input/output error; usually used for physical read
	// or write errors. i.e. the disk or other physical device hardware
	// is returning errors.
	FILE_ERROR_IO FileError = 21

	// Operation not permitted; only the owner of the
	// file (or other resource) or processes with special privileges can
	// perform the operation.
	FILE_ERROR_PERM FileError = 22

	// Function not implemented; this indicates that
	// the system is missing some functionality.
	FILE_ERROR_NOSYS FileError = 23

	// Does not correspond to a UNIX error code; this
	// is the standard "failed for unspecified reason" error code present
	// in all #GError error code enumerations. Returned if no specific
	// code applies.
	FILE_ERROR_FAILED FileError = 24
)

// Error codes returned by #GIOChannel operations.
type IOChannelError C.GIOChannelError

const (
	// File too large.
	IO_CHANNEL_ERROR_FBIG IOChannelError = 0

	// Invalid argument.
	IO_CHANNEL_ERROR_INVAL IOChannelError = 1

	// IO error.
	IO_CHANNEL_ERROR_IO IOChannelError = 2

	// File is a directory.
	IO_CHANNEL_ERROR_ISDIR IOChannelError = 3

	// No space left on device.
	IO_CHANNEL_ERROR_NOSPC IOChannelError = 4

	// No such device or address.
	IO_CHANNEL_ERROR_NXIO IOChannelError = 5

	// Value too large for defined datatype.
	IO_CHANNEL_ERROR_OVERFLOW IOChannelError = 6

	// Broken pipe.
	IO_CHANNEL_ERROR_PIPE IOChannelError = 7

	// Some other error.
	IO_CHANNEL_ERROR_FAILED IOChannelError = 8
)

// #GIOError is only used by the deprecated functions
// g_io_channel_read(), g_io_channel_write(), and g_io_channel_seek().
type IOError C.GIOError

const (
	// no error
	IO_ERROR_NONE IOError = 0

	// an EAGAIN error occurred
	IO_ERROR_AGAIN IOError = 1

	// an EINVAL error occurred
	IO_ERROR_INVAL IOError = 2

	// another error occurred
	IO_ERROR_UNKNOWN IOError = 3
)

// Stati returned by most of the #GIOFuncs functions.
type IOStatus C.GIOStatus

const (
	// An error occurred.
	IO_STATUS_ERROR IOStatus = 0

	// Success.
	IO_STATUS_NORMAL IOStatus = 1

	// End of file.
	IO_STATUS_EOF IOStatus = 2

	// Resource temporarily unavailable.
	IO_STATUS_AGAIN IOStatus = 3
)

// Error codes returned by key file parsing.
type KeyFileError C.GKeyFileError

const (
	// the text being parsed was in
	// an unknown encoding
	KEY_FILE_ERROR_UNKNOWN_ENCODING KeyFileError = 0

	// document was ill-formed
	KEY_FILE_ERROR_PARSE KeyFileError = 1

	// the file was not found
	KEY_FILE_ERROR_NOT_FOUND KeyFileError = 2

	// a requested key was not found
	KEY_FILE_ERROR_KEY_NOT_FOUND KeyFileError = 3

	// a requested group was not found
	KEY_FILE_ERROR_GROUP_NOT_FOUND KeyFileError = 4

	// a value could not be parsed
	KEY_FILE_ERROR_INVALID_VALUE KeyFileError = 5
)

// Error codes returned by markup parsing.
type MarkupError C.GMarkupError

const (
	// text being parsed was not valid UTF-8
	MARKUP_ERROR_BAD_UTF8 MarkupError = 0

	// document contained nothing, or only whitespace
	MARKUP_ERROR_EMPTY MarkupError = 1

	// document was ill-formed
	MARKUP_ERROR_PARSE MarkupError = 2

	// error should be set by #GMarkupParser
	// functions; element wasn't known
	MARKUP_ERROR_UNKNOWN_ELEMENT MarkupError = 3

	// error should be set by #GMarkupParser
	// functions; attribute wasn't known
	MARKUP_ERROR_UNKNOWN_ATTRIBUTE MarkupError = 4

	// error should be set by #GMarkupParser
	// functions; content was invalid
	MARKUP_ERROR_INVALID_CONTENT MarkupError = 5

	// error should be set by #GMarkupParser
	// functions; a required attribute was missing
	MARKUP_ERROR_MISSING_ATTRIBUTE MarkupError = 6
)

// Defines how a Unicode string is transformed in a canonical
// form, standardizing such issues as whether a character with
// an accent is represented as a base character and combining
// accent or as a single precomposed character. Unicode strings
// should generally be normalized before comparing them.
type NormalizeMode C.GNormalizeMode

const (
	// standardize differences that do not affect the
	// text content, such as the above-mentioned accent representation
	NORMALIZE_DEFAULT NormalizeMode = 0

	// another name for %G_NORMALIZE_DEFAULT
	NORMALIZE_NFD NormalizeMode = 0

	// like %G_NORMALIZE_DEFAULT, but with
	// composed forms rather than a maximally decomposed form
	NORMALIZE_DEFAULT_COMPOSE NormalizeMode = 1

	// another name for %G_NORMALIZE_DEFAULT_COMPOSE
	NORMALIZE_NFC NormalizeMode = 1

	// beyond %G_NORMALIZE_DEFAULT also standardize the
	// "compatibility" characters in Unicode, such as SUPERSCRIPT THREE
	// to the standard forms (in this case DIGIT THREE). Formatting
	// information may be lost but for most text operations such
	// characters should be considered the same
	NORMALIZE_ALL NormalizeMode = 2

	// another name for %G_NORMALIZE_ALL
	NORMALIZE_NFKD NormalizeMode = 2

	// like %G_NORMALIZE_ALL, but with composed
	// forms rather than a maximally decomposed form
	NORMALIZE_ALL_COMPOSE NormalizeMode = 3

	// another name for %G_NORMALIZE_ALL_COMPOSE
	NORMALIZE_NFKC NormalizeMode = 3
)

// The #GOptionArg enum values determine which type of extra argument the
// options expect to find. If an option expects an extra argument, it can
// be specified in several ways; with a short option: `-x arg`, with a long
// option: `--name arg` or combined in a single argument: `--name=arg`.
type OptionArg C.GOptionArg

const (
	// No extra argument. This is useful for simple flags.
	OPTION_ARG_NONE OptionArg = 0

	// The option takes a string argument.
	OPTION_ARG_STRING OptionArg = 1

	// The option takes an integer argument.
	OPTION_ARG_INT OptionArg = 2

	// The option provides a callback (of type
	// #GOptionArgFunc) to parse the extra argument.
	OPTION_ARG_CALLBACK OptionArg = 3

	// The option takes a filename as argument.
	OPTION_ARG_FILENAME OptionArg = 4

	// The option takes a string argument, multiple
	// uses of the option are collected into an array of strings.
	OPTION_ARG_STRING_ARRAY OptionArg = 5

	// The option takes a filename as argument,
	// multiple uses of the option are collected into an array of strings.
	OPTION_ARG_FILENAME_ARRAY OptionArg = 6

	// The option takes a double argument. The argument
	// can be formatted either for the user's locale or for the "C" locale.
	// Since 2.12
	OPTION_ARG_DOUBLE OptionArg = 7

	// The option takes a 64-bit integer. Like
	// %G_OPTION_ARG_INT but for larger numbers. The number can be in
	// decimal base, or in hexadecimal (when prefixed with `0x`, for
	// example, `0xffffffff`). Since 2.12
	OPTION_ARG_INT64 OptionArg = 8
)

// Error codes returned by option parsing.
type OptionError C.GOptionError

const (
	// An option was not known to the parser.
	// This error will only be reported, if the parser hasn't been instructed
	// to ignore unknown options, see g_option_context_set_ignore_unknown_options().
	OPTION_ERROR_UNKNOWN_OPTION OptionError = 0

	// A value couldn't be parsed.
	OPTION_ERROR_BAD_VALUE OptionError = 1

	// A #GOptionArgFunc callback failed.
	OPTION_ERROR_FAILED OptionError = 2
)

// An enumeration specifying the base position for a
// g_io_channel_seek_position() operation.
type SeekType C.GSeekType

const (
	// the current position in the file.
	SEEK_CUR SeekType = 0

	// the start of the file.
	SEEK_SET SeekType = 1

	// the end of the file.
	SEEK_END SeekType = 2
)

// Error codes returned by shell functions.
type ShellError C.GShellError

const (
	// Mismatched or otherwise mangled quoting.
	SHELL_ERROR_BAD_QUOTING ShellError = 0

	// String to be parsed was empty.
	SHELL_ERROR_EMPTY_STRING ShellError = 1

	// Some other error.
	SHELL_ERROR_FAILED ShellError = 2
)

type SliceConfig C.GSliceConfig

const (
	SLICE_CONFIG_ALWAYS_MALLOC SliceConfig = 1

	SLICE_CONFIG_BYPASS_MAGAZINES SliceConfig = 2

	SLICE_CONFIG_WORKING_SET_MSECS SliceConfig = 3

	SLICE_CONFIG_COLOR_INCREMENT SliceConfig = 4

	SLICE_CONFIG_CHUNK_SIZES SliceConfig = 5

	SLICE_CONFIG_CONTENTION_COUNTER SliceConfig = 6
)

// Error codes returned by spawning processes.
type SpawnError C.GSpawnError

const (
	// Fork failed due to lack of memory.
	SPAWN_ERROR_FORK SpawnError = 0

	// Read or select on pipes failed.
	SPAWN_ERROR_READ SpawnError = 1

	// Changing to working directory failed.
	SPAWN_ERROR_CHDIR SpawnError = 2

	// execv() returned `EACCES`
	SPAWN_ERROR_ACCES SpawnError = 3

	// execv() returned `EPERM`
	SPAWN_ERROR_PERM SpawnError = 4

	// execv() returned `E2BIG`
	SPAWN_ERROR_TOO_BIG SpawnError = 5

	// deprecated alias for %G_SPAWN_ERROR_TOO_BIG
	SPAWN_ERROR_2BIG SpawnError = 5

	// execv() returned `ENOEXEC`
	SPAWN_ERROR_NOEXEC SpawnError = 6

	// execv() returned `ENAMETOOLONG`
	SPAWN_ERROR_NAMETOOLONG SpawnError = 7

	// execv() returned `ENOENT`
	SPAWN_ERROR_NOENT SpawnError = 8

	// execv() returned `ENOMEM`
	SPAWN_ERROR_NOMEM SpawnError = 9

	// execv() returned `ENOTDIR`
	SPAWN_ERROR_NOTDIR SpawnError = 10

	// execv() returned `ELOOP`
	SPAWN_ERROR_LOOP SpawnError = 11

	// execv() returned `ETXTBUSY`
	SPAWN_ERROR_TXTBUSY SpawnError = 12

	// execv() returned `EIO`
	SPAWN_ERROR_IO SpawnError = 13

	// execv() returned `ENFILE`
	SPAWN_ERROR_NFILE SpawnError = 14

	// execv() returned `EMFILE`
	SPAWN_ERROR_MFILE SpawnError = 15

	// execv() returned `EINVAL`
	SPAWN_ERROR_INVAL SpawnError = 16

	// execv() returned `EISDIR`
	SPAWN_ERROR_ISDIR SpawnError = 17

	// execv() returned `ELIBBAD`
	SPAWN_ERROR_LIBBAD SpawnError = 18

	// Some other fatal failure,
	// `error->message` should explain.
	SPAWN_ERROR_FAILED SpawnError = 19
)

type TestLogType C.GTestLogType

const (
	TEST_LOG_NONE TestLogType = 0

	TEST_LOG_ERROR TestLogType = 1

	TEST_LOG_START_BINARY TestLogType = 2

	TEST_LOG_LIST_CASE TestLogType = 3

	TEST_LOG_SKIP_CASE TestLogType = 4

	TEST_LOG_START_CASE TestLogType = 5

	TEST_LOG_STOP_CASE TestLogType = 6

	TEST_LOG_MIN_RESULT TestLogType = 7

	TEST_LOG_MAX_RESULT TestLogType = 8

	TEST_LOG_MESSAGE TestLogType = 9

	TEST_LOG_START_SUITE TestLogType = 10

	TEST_LOG_STOP_SUITE TestLogType = 11
)

// Blacklisted : GTestResult

// Possible errors of thread related functions.
type ThreadError C.GThreadError

const (
	// a thread couldn't be created due to resource
	// shortage. Try again later.
	THREAD_ERROR_AGAIN ThreadError = 0
)

// Disambiguates a given time in two ways.
//
// First, specifies if the given time is in universal or local time.
//
// Second, if the time is in local time, specifies if it is local
// standard time or local daylight time.  This is important for the case
// where the same local time occurs twice (during daylight savings time
// transitions, for example).
type TimeType C.GTimeType

const (
	// the time is in local standard time
	TIME_TYPE_STANDARD TimeType = 0

	// the time is in local daylight time
	TIME_TYPE_DAYLIGHT TimeType = 1

	// the time is in UTC
	TIME_TYPE_UNIVERSAL TimeType = 2
)

// The possible types of token returned from each
// g_scanner_get_next_token() call.
type TokenType C.GTokenType

const (
	// the end of the file
	TOKEN_EOF TokenType = 0

	// a '(' character
	TOKEN_LEFT_PAREN TokenType = 40

	// a ')' character
	TOKEN_RIGHT_PAREN TokenType = 41

	// a '{' character
	TOKEN_LEFT_CURLY TokenType = 123

	// a '}' character
	TOKEN_RIGHT_CURLY TokenType = 125

	// a '[' character
	TOKEN_LEFT_BRACE TokenType = 91

	// a ']' character
	TOKEN_RIGHT_BRACE TokenType = 93

	// a '=' character
	TOKEN_EQUAL_SIGN TokenType = 61

	// a ',' character
	TOKEN_COMMA TokenType = 44

	// not a token
	TOKEN_NONE TokenType = 256

	// an error occurred
	TOKEN_ERROR TokenType = 257

	// a character
	TOKEN_CHAR TokenType = 258

	// a binary integer
	TOKEN_BINARY TokenType = 259

	// an octal integer
	TOKEN_OCTAL TokenType = 260

	// an integer
	TOKEN_INT TokenType = 261

	// a hex integer
	TOKEN_HEX TokenType = 262

	// a floating point number
	TOKEN_FLOAT TokenType = 263

	// a string
	TOKEN_STRING TokenType = 264

	// a symbol
	TOKEN_SYMBOL TokenType = 265

	// an identifier
	TOKEN_IDENTIFIER TokenType = 266

	// a null identifier
	TOKEN_IDENTIFIER_NULL TokenType = 267

	// one line comment
	TOKEN_COMMENT_SINGLE TokenType = 268

	// multi line comment
	TOKEN_COMMENT_MULTI TokenType = 269
)

// Specifies the type of traveral performed by g_tree_traverse(),
// g_node_traverse() and g_node_find(). The different orders are
// illustrated here:
// - In order: A, B, C, D, E, F, G, H, I
// ![](Sorted_binary_tree_inorder.svg)
// - Pre order: F, B, A, D, C, E, G, I, H
// ![](Sorted_binary_tree_preorder.svg)
// - Post order: A, C, E, D, B, H, I, G, F
// ![](Sorted_binary_tree_postorder.svg)
// - Level order: F, B, G, A, D, I, C, E, H
// ![](Sorted_binary_tree_breadth-first_traversal.svg)
type TraverseType C.GTraverseType

const (
	// vists a node's left child first, then the node itself,
	// then its right child. This is the one to use if you
	// want the output sorted according to the compare
	// function.
	IN_ORDER TraverseType = 0

	// visits a node, then its children.
	PRE_ORDER TraverseType = 1

	// visits the node's children, then the node itself.
	POST_ORDER TraverseType = 2

	// is not implemented for
	// [balanced binary trees][glib-Balanced-Binary-Trees].
	// For [n-ary trees][glib-N-ary-Trees], it
	// vists the root node first, then its children, then
	// its grandchildren, and so on. Note that this is less
	// efficient than the other orders.
	LEVEL_ORDER TraverseType = 3
)

// These are the possible line break classifications.
//
// Since new unicode versions may add new types here, applications should be ready
// to handle unknown values. They may be regarded as %G_UNICODE_BREAK_UNKNOWN.
//
// See [Unicode Line Breaking Algorithm](http://www.unicode.org/unicode/reports/tr14/).
type UnicodeBreakType C.GUnicodeBreakType

const (
	// Mandatory Break (BK)
	UNICODE_BREAK_MANDATORY UnicodeBreakType = 0

	// Carriage Return (CR)
	UNICODE_BREAK_CARRIAGE_RETURN UnicodeBreakType = 1

	// Line Feed (LF)
	UNICODE_BREAK_LINE_FEED UnicodeBreakType = 2

	// Attached Characters and Combining Marks (CM)
	UNICODE_BREAK_COMBINING_MARK UnicodeBreakType = 3

	// Surrogates (SG)
	UNICODE_BREAK_SURROGATE UnicodeBreakType = 4

	// Zero Width Space (ZW)
	UNICODE_BREAK_ZERO_WIDTH_SPACE UnicodeBreakType = 5

	// Inseparable (IN)
	UNICODE_BREAK_INSEPARABLE UnicodeBreakType = 6

	// Non-breaking ("Glue") (GL)
	UNICODE_BREAK_NON_BREAKING_GLUE UnicodeBreakType = 7

	// Contingent Break Opportunity (CB)
	UNICODE_BREAK_CONTINGENT UnicodeBreakType = 8

	// Space (SP)
	UNICODE_BREAK_SPACE UnicodeBreakType = 9

	// Break Opportunity After (BA)
	UNICODE_BREAK_AFTER UnicodeBreakType = 10

	// Break Opportunity Before (BB)
	UNICODE_BREAK_BEFORE UnicodeBreakType = 11

	// Break Opportunity Before and After (B2)
	UNICODE_BREAK_BEFORE_AND_AFTER UnicodeBreakType = 12

	// Hyphen (HY)
	UNICODE_BREAK_HYPHEN UnicodeBreakType = 13

	// Nonstarter (NS)
	UNICODE_BREAK_NON_STARTER UnicodeBreakType = 14

	// Opening Punctuation (OP)
	UNICODE_BREAK_OPEN_PUNCTUATION UnicodeBreakType = 15

	// Closing Punctuation (CL)
	UNICODE_BREAK_CLOSE_PUNCTUATION UnicodeBreakType = 16

	// Ambiguous Quotation (QU)
	UNICODE_BREAK_QUOTATION UnicodeBreakType = 17

	// Exclamation/Interrogation (EX)
	UNICODE_BREAK_EXCLAMATION UnicodeBreakType = 18

	// Ideographic (ID)
	UNICODE_BREAK_IDEOGRAPHIC UnicodeBreakType = 19

	// Numeric (NU)
	UNICODE_BREAK_NUMERIC UnicodeBreakType = 20

	// Infix Separator (Numeric) (IS)
	UNICODE_BREAK_INFIX_SEPARATOR UnicodeBreakType = 21

	// Symbols Allowing Break After (SY)
	UNICODE_BREAK_SYMBOL UnicodeBreakType = 22

	// Ordinary Alphabetic and Symbol Characters (AL)
	UNICODE_BREAK_ALPHABETIC UnicodeBreakType = 23

	// Prefix (Numeric) (PR)
	UNICODE_BREAK_PREFIX UnicodeBreakType = 24

	// Postfix (Numeric) (PO)
	UNICODE_BREAK_POSTFIX UnicodeBreakType = 25

	// Complex Content Dependent (South East Asian) (SA)
	UNICODE_BREAK_COMPLEX_CONTEXT UnicodeBreakType = 26

	// Ambiguous (Alphabetic or Ideographic) (AI)
	UNICODE_BREAK_AMBIGUOUS UnicodeBreakType = 27

	// Unknown (XX)
	UNICODE_BREAK_UNKNOWN UnicodeBreakType = 28

	// Next Line (NL)
	UNICODE_BREAK_NEXT_LINE UnicodeBreakType = 29

	// Word Joiner (WJ)
	UNICODE_BREAK_WORD_JOINER UnicodeBreakType = 30

	// Hangul L Jamo (JL)
	UNICODE_BREAK_HANGUL_L_JAMO UnicodeBreakType = 31

	// Hangul V Jamo (JV)
	UNICODE_BREAK_HANGUL_V_JAMO UnicodeBreakType = 32

	// Hangul T Jamo (JT)
	UNICODE_BREAK_HANGUL_T_JAMO UnicodeBreakType = 33

	// Hangul LV Syllable (H2)
	UNICODE_BREAK_HANGUL_LV_SYLLABLE UnicodeBreakType = 34

	// Hangul LVT Syllable (H3)
	UNICODE_BREAK_HANGUL_LVT_SYLLABLE UnicodeBreakType = 35

	// Closing Parenthesis (CP). Since 2.28
	UNICODE_BREAK_CLOSE_PARANTHESIS UnicodeBreakType = 36

	// Conditional Japanese Starter (CJ). Since: 2.32
	UNICODE_BREAK_CONDITIONAL_JAPANESE_STARTER UnicodeBreakType = 37

	// Hebrew Letter (HL). Since: 2.32
	UNICODE_BREAK_HEBREW_LETTER UnicodeBreakType = 38

	// Regional Indicator (RI). Since: 2.36
	UNICODE_BREAK_REGIONAL_INDICATOR UnicodeBreakType = 39

	// Emoji Base (EB). Since: 2.50
	UNICODE_BREAK_EMOJI_BASE UnicodeBreakType = 40

	// Emoji Modifier (EM). Since: 2.50
	UNICODE_BREAK_EMOJI_MODIFIER UnicodeBreakType = 41

	// Zero Width Joiner (ZWJ). Since: 2.50
	UNICODE_BREAK_ZERO_WIDTH_JOINER UnicodeBreakType = 42
)

// The #GUnicodeScript enumeration identifies different writing
// systems. The values correspond to the names as defined in the
// Unicode standard. The enumeration has been added in GLib 2.14,
// and is interchangeable with #PangoScript.
//
// Note that new types may be added in the future. Applications
// should be ready to handle unknown values.
// See [Unicode Standard Annex #24: Script names](http://www.unicode.org/reports/tr24/).
type UnicodeScript C.GUnicodeScript

const (
	// a value never returned from g_unichar_get_script()
	UNICODE_SCRIPT_INVALID_CODE UnicodeScript = -1

	// a character used by multiple different scripts
	UNICODE_SCRIPT_COMMON UnicodeScript = 0

	// a mark glyph that takes its script from the
	// base glyph to which it is attached
	UNICODE_SCRIPT_INHERITED UnicodeScript = 1

	// Arabic
	UNICODE_SCRIPT_ARABIC UnicodeScript = 2

	// Armenian
	UNICODE_SCRIPT_ARMENIAN UnicodeScript = 3

	// Bengali
	UNICODE_SCRIPT_BENGALI UnicodeScript = 4

	// Bopomofo
	UNICODE_SCRIPT_BOPOMOFO UnicodeScript = 5

	// Cherokee
	UNICODE_SCRIPT_CHEROKEE UnicodeScript = 6

	// Coptic
	UNICODE_SCRIPT_COPTIC UnicodeScript = 7

	// Cyrillic
	UNICODE_SCRIPT_CYRILLIC UnicodeScript = 8

	// Deseret
	UNICODE_SCRIPT_DESERET UnicodeScript = 9

	// Devanagari
	UNICODE_SCRIPT_DEVANAGARI UnicodeScript = 10

	// Ethiopic
	UNICODE_SCRIPT_ETHIOPIC UnicodeScript = 11

	// Georgian
	UNICODE_SCRIPT_GEORGIAN UnicodeScript = 12

	// Gothic
	UNICODE_SCRIPT_GOTHIC UnicodeScript = 13

	// Greek
	UNICODE_SCRIPT_GREEK UnicodeScript = 14

	// Gujarati
	UNICODE_SCRIPT_GUJARATI UnicodeScript = 15

	// Gurmukhi
	UNICODE_SCRIPT_GURMUKHI UnicodeScript = 16

	// Han
	UNICODE_SCRIPT_HAN UnicodeScript = 17

	// Hangul
	UNICODE_SCRIPT_HANGUL UnicodeScript = 18

	// Hebrew
	UNICODE_SCRIPT_HEBREW UnicodeScript = 19

	// Hiragana
	UNICODE_SCRIPT_HIRAGANA UnicodeScript = 20

	// Kannada
	UNICODE_SCRIPT_KANNADA UnicodeScript = 21

	// Katakana
	UNICODE_SCRIPT_KATAKANA UnicodeScript = 22

	// Khmer
	UNICODE_SCRIPT_KHMER UnicodeScript = 23

	// Lao
	UNICODE_SCRIPT_LAO UnicodeScript = 24

	// Latin
	UNICODE_SCRIPT_LATIN UnicodeScript = 25

	// Malayalam
	UNICODE_SCRIPT_MALAYALAM UnicodeScript = 26

	// Mongolian
	UNICODE_SCRIPT_MONGOLIAN UnicodeScript = 27

	// Myanmar
	UNICODE_SCRIPT_MYANMAR UnicodeScript = 28

	// Ogham
	UNICODE_SCRIPT_OGHAM UnicodeScript = 29

	// Old Italic
	UNICODE_SCRIPT_OLD_ITALIC UnicodeScript = 30

	// Oriya
	UNICODE_SCRIPT_ORIYA UnicodeScript = 31

	// Runic
	UNICODE_SCRIPT_RUNIC UnicodeScript = 32

	// Sinhala
	UNICODE_SCRIPT_SINHALA UnicodeScript = 33

	// Syriac
	UNICODE_SCRIPT_SYRIAC UnicodeScript = 34

	// Tamil
	UNICODE_SCRIPT_TAMIL UnicodeScript = 35

	// Telugu
	UNICODE_SCRIPT_TELUGU UnicodeScript = 36

	// Thaana
	UNICODE_SCRIPT_THAANA UnicodeScript = 37

	// Thai
	UNICODE_SCRIPT_THAI UnicodeScript = 38

	// Tibetan
	UNICODE_SCRIPT_TIBETAN UnicodeScript = 39

	// Canadian Aboriginal
	UNICODE_SCRIPT_CANADIAN_ABORIGINAL UnicodeScript = 40

	// Yi
	UNICODE_SCRIPT_YI UnicodeScript = 41

	// Tagalog
	UNICODE_SCRIPT_TAGALOG UnicodeScript = 42

	// Hanunoo
	UNICODE_SCRIPT_HANUNOO UnicodeScript = 43

	// Buhid
	UNICODE_SCRIPT_BUHID UnicodeScript = 44

	// Tagbanwa
	UNICODE_SCRIPT_TAGBANWA UnicodeScript = 45

	// Braille
	UNICODE_SCRIPT_BRAILLE UnicodeScript = 46

	// Cypriot
	UNICODE_SCRIPT_CYPRIOT UnicodeScript = 47

	// Limbu
	UNICODE_SCRIPT_LIMBU UnicodeScript = 48

	// Osmanya
	UNICODE_SCRIPT_OSMANYA UnicodeScript = 49

	// Shavian
	UNICODE_SCRIPT_SHAVIAN UnicodeScript = 50

	// Linear B
	UNICODE_SCRIPT_LINEAR_B UnicodeScript = 51

	// Tai Le
	UNICODE_SCRIPT_TAI_LE UnicodeScript = 52

	// Ugaritic
	UNICODE_SCRIPT_UGARITIC UnicodeScript = 53

	// New Tai Lue
	UNICODE_SCRIPT_NEW_TAI_LUE UnicodeScript = 54

	// Buginese
	UNICODE_SCRIPT_BUGINESE UnicodeScript = 55

	// Glagolitic
	UNICODE_SCRIPT_GLAGOLITIC UnicodeScript = 56

	// Tifinagh
	UNICODE_SCRIPT_TIFINAGH UnicodeScript = 57

	// Syloti Nagri
	UNICODE_SCRIPT_SYLOTI_NAGRI UnicodeScript = 58

	// Old Persian
	UNICODE_SCRIPT_OLD_PERSIAN UnicodeScript = 59

	// Kharoshthi
	UNICODE_SCRIPT_KHAROSHTHI UnicodeScript = 60

	// an unassigned code point
	UNICODE_SCRIPT_UNKNOWN UnicodeScript = 61

	// Balinese
	UNICODE_SCRIPT_BALINESE UnicodeScript = 62

	// Cuneiform
	UNICODE_SCRIPT_CUNEIFORM UnicodeScript = 63

	// Phoenician
	UNICODE_SCRIPT_PHOENICIAN UnicodeScript = 64

	// Phags-pa
	UNICODE_SCRIPT_PHAGS_PA UnicodeScript = 65

	// N'Ko
	UNICODE_SCRIPT_NKO UnicodeScript = 66

	// Kayah Li. Since 2.16.3
	UNICODE_SCRIPT_KAYAH_LI UnicodeScript = 67

	// Lepcha. Since 2.16.3
	UNICODE_SCRIPT_LEPCHA UnicodeScript = 68

	// Rejang. Since 2.16.3
	UNICODE_SCRIPT_REJANG UnicodeScript = 69

	// Sundanese. Since 2.16.3
	UNICODE_SCRIPT_SUNDANESE UnicodeScript = 70

	// Saurashtra. Since 2.16.3
	UNICODE_SCRIPT_SAURASHTRA UnicodeScript = 71

	// Cham. Since 2.16.3
	UNICODE_SCRIPT_CHAM UnicodeScript = 72

	// Ol Chiki. Since 2.16.3
	UNICODE_SCRIPT_OL_CHIKI UnicodeScript = 73

	// Vai. Since 2.16.3
	UNICODE_SCRIPT_VAI UnicodeScript = 74

	// Carian. Since 2.16.3
	UNICODE_SCRIPT_CARIAN UnicodeScript = 75

	// Lycian. Since 2.16.3
	UNICODE_SCRIPT_LYCIAN UnicodeScript = 76

	// Lydian. Since 2.16.3
	UNICODE_SCRIPT_LYDIAN UnicodeScript = 77

	// Avestan. Since 2.26
	UNICODE_SCRIPT_AVESTAN UnicodeScript = 78

	// Bamum. Since 2.26
	UNICODE_SCRIPT_BAMUM UnicodeScript = 79

	// Egyptian Hieroglpyhs. Since 2.26
	UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS UnicodeScript = 80

	// Imperial Aramaic. Since 2.26
	UNICODE_SCRIPT_IMPERIAL_ARAMAIC UnicodeScript = 81

	// Inscriptional Pahlavi. Since 2.26
	UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI UnicodeScript = 82

	// Inscriptional Parthian. Since 2.26
	UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN UnicodeScript = 83

	// Javanese. Since 2.26
	UNICODE_SCRIPT_JAVANESE UnicodeScript = 84

	// Kaithi. Since 2.26
	UNICODE_SCRIPT_KAITHI UnicodeScript = 85

	// Lisu. Since 2.26
	UNICODE_SCRIPT_LISU UnicodeScript = 86

	// Meetei Mayek. Since 2.26
	UNICODE_SCRIPT_MEETEI_MAYEK UnicodeScript = 87

	// Old South Arabian. Since 2.26
	UNICODE_SCRIPT_OLD_SOUTH_ARABIAN UnicodeScript = 88

	// Old Turkic. Since 2.28
	UNICODE_SCRIPT_OLD_TURKIC UnicodeScript = 89

	// Samaritan. Since 2.26
	UNICODE_SCRIPT_SAMARITAN UnicodeScript = 90

	// Tai Tham. Since 2.26
	UNICODE_SCRIPT_TAI_THAM UnicodeScript = 91

	// Tai Viet. Since 2.26
	UNICODE_SCRIPT_TAI_VIET UnicodeScript = 92

	// Batak. Since 2.28
	UNICODE_SCRIPT_BATAK UnicodeScript = 93

	// Brahmi. Since 2.28
	UNICODE_SCRIPT_BRAHMI UnicodeScript = 94

	// Mandaic. Since 2.28
	UNICODE_SCRIPT_MANDAIC UnicodeScript = 95

	// Chakma. Since: 2.32
	UNICODE_SCRIPT_CHAKMA UnicodeScript = 96

	// Meroitic Cursive. Since: 2.32
	UNICODE_SCRIPT_MEROITIC_CURSIVE UnicodeScript = 97

	// Meroitic Hieroglyphs. Since: 2.32
	UNICODE_SCRIPT_MEROITIC_HIEROGLYPHS UnicodeScript = 98

	// Miao. Since: 2.32
	UNICODE_SCRIPT_MIAO UnicodeScript = 99

	// Sharada. Since: 2.32
	UNICODE_SCRIPT_SHARADA UnicodeScript = 100

	// Sora Sompeng. Since: 2.32
	UNICODE_SCRIPT_SORA_SOMPENG UnicodeScript = 101

	// Takri. Since: 2.32
	UNICODE_SCRIPT_TAKRI UnicodeScript = 102

	// Bassa. Since: 2.42
	UNICODE_SCRIPT_BASSA_VAH UnicodeScript = 103

	// Caucasian Albanian. Since: 2.42
	UNICODE_SCRIPT_CAUCASIAN_ALBANIAN UnicodeScript = 104

	// Duployan. Since: 2.42
	UNICODE_SCRIPT_DUPLOYAN UnicodeScript = 105

	// Elbasan. Since: 2.42
	UNICODE_SCRIPT_ELBASAN UnicodeScript = 106

	// Grantha. Since: 2.42
	UNICODE_SCRIPT_GRANTHA UnicodeScript = 107

	// Kjohki. Since: 2.42
	UNICODE_SCRIPT_KHOJKI UnicodeScript = 108

	// Khudawadi, Sindhi. Since: 2.42
	UNICODE_SCRIPT_KHUDAWADI UnicodeScript = 109

	// Linear A. Since: 2.42
	UNICODE_SCRIPT_LINEAR_A UnicodeScript = 110

	// Mahajani. Since: 2.42
	UNICODE_SCRIPT_MAHAJANI UnicodeScript = 111

	// Manichaean. Since: 2.42
	UNICODE_SCRIPT_MANICHAEAN UnicodeScript = 112

	// Mende Kikakui. Since: 2.42
	UNICODE_SCRIPT_MENDE_KIKAKUI UnicodeScript = 113

	// Modi. Since: 2.42
	UNICODE_SCRIPT_MODI UnicodeScript = 114

	// Mro. Since: 2.42
	UNICODE_SCRIPT_MRO UnicodeScript = 115

	// Nabataean. Since: 2.42
	UNICODE_SCRIPT_NABATAEAN UnicodeScript = 116

	// Old North Arabian. Since: 2.42
	UNICODE_SCRIPT_OLD_NORTH_ARABIAN UnicodeScript = 117

	// Old Permic. Since: 2.42
	UNICODE_SCRIPT_OLD_PERMIC UnicodeScript = 118

	// Pahawh Hmong. Since: 2.42
	UNICODE_SCRIPT_PAHAWH_HMONG UnicodeScript = 119

	// Palmyrene. Since: 2.42
	UNICODE_SCRIPT_PALMYRENE UnicodeScript = 120

	// Pau Cin Hau. Since: 2.42
	UNICODE_SCRIPT_PAU_CIN_HAU UnicodeScript = 121

	// Psalter Pahlavi. Since: 2.42
	UNICODE_SCRIPT_PSALTER_PAHLAVI UnicodeScript = 122

	// Siddham. Since: 2.42
	UNICODE_SCRIPT_SIDDHAM UnicodeScript = 123

	// Tirhuta. Since: 2.42
	UNICODE_SCRIPT_TIRHUTA UnicodeScript = 124

	// Warang Citi. Since: 2.42
	UNICODE_SCRIPT_WARANG_CITI UnicodeScript = 125

	// Ahom. Since: 2.48
	UNICODE_SCRIPT_AHOM UnicodeScript = 126

	// Anatolian Hieroglyphs. Since: 2.48
	UNICODE_SCRIPT_ANATOLIAN_HIEROGLYPHS UnicodeScript = 127

	// Hatran. Since: 2.48
	UNICODE_SCRIPT_HATRAN UnicodeScript = 128

	// Multani. Since: 2.48
	UNICODE_SCRIPT_MULTANI UnicodeScript = 129

	// Old Hungarian. Since: 2.48
	UNICODE_SCRIPT_OLD_HUNGARIAN UnicodeScript = 130

	// Signwriting. Since: 2.48
	UNICODE_SCRIPT_SIGNWRITING UnicodeScript = 131

	// Adlam. Since: 2.50
	UNICODE_SCRIPT_ADLAM UnicodeScript = 132

	// Bhaiksuki. Since: 2.50
	UNICODE_SCRIPT_BHAIKSUKI UnicodeScript = 133

	// Marchen. Since: 2.50
	UNICODE_SCRIPT_MARCHEN UnicodeScript = 134

	// Newa. Since: 2.50
	UNICODE_SCRIPT_NEWA UnicodeScript = 135

	// Osage. Since: 2.50
	UNICODE_SCRIPT_OSAGE UnicodeScript = 136

	// Tangut. Since: 2.50
	UNICODE_SCRIPT_TANGUT UnicodeScript = 137

	// Masaram Gondi. Since: 2.54
	UNICODE_SCRIPT_MASARAM_GONDI UnicodeScript = 138

	// Nushu. Since: 2.54
	UNICODE_SCRIPT_NUSHU UnicodeScript = 139

	// Soyombo. Since: 2.54
	UNICODE_SCRIPT_SOYOMBO UnicodeScript = 140

	// Zanabazar Square. Since: 2.54
	UNICODE_SCRIPT_ZANABAZAR_SQUARE UnicodeScript = 141
)

// These are the possible character classifications from the
// Unicode specification.
// See [Unicode Character Database](http://www.unicode.org/reports/tr44/#General_Category_Values).
type UnicodeType C.GUnicodeType

const (
	// General category "Other, Control" (Cc)
	UNICODE_CONTROL UnicodeType = 0

	// General category "Other, Format" (Cf)
	UNICODE_FORMAT UnicodeType = 1

	// General category "Other, Not Assigned" (Cn)
	UNICODE_UNASSIGNED UnicodeType = 2

	// General category "Other, Private Use" (Co)
	UNICODE_PRIVATE_USE UnicodeType = 3

	// General category "Other, Surrogate" (Cs)
	UNICODE_SURROGATE UnicodeType = 4

	// General category "Letter, Lowercase" (Ll)
	UNICODE_LOWERCASE_LETTER UnicodeType = 5

	// General category "Letter, Modifier" (Lm)
	UNICODE_MODIFIER_LETTER UnicodeType = 6

	// General category "Letter, Other" (Lo)
	UNICODE_OTHER_LETTER UnicodeType = 7

	// General category "Letter, Titlecase" (Lt)
	UNICODE_TITLECASE_LETTER UnicodeType = 8

	// General category "Letter, Uppercase" (Lu)
	UNICODE_UPPERCASE_LETTER UnicodeType = 9

	// General category "Mark, Spacing" (Mc)
	UNICODE_SPACING_MARK UnicodeType = 10

	// General category "Mark, Enclosing" (Me)
	UNICODE_ENCLOSING_MARK UnicodeType = 11

	// General category "Mark, Nonspacing" (Mn)
	UNICODE_NON_SPACING_MARK UnicodeType = 12

	// General category "Number, Decimal Digit" (Nd)
	UNICODE_DECIMAL_NUMBER UnicodeType = 13

	// General category "Number, Letter" (Nl)
	UNICODE_LETTER_NUMBER UnicodeType = 14

	// General category "Number, Other" (No)
	UNICODE_OTHER_NUMBER UnicodeType = 15

	// General category "Punctuation, Connector" (Pc)
	UNICODE_CONNECT_PUNCTUATION UnicodeType = 16

	// General category "Punctuation, Dash" (Pd)
	UNICODE_DASH_PUNCTUATION UnicodeType = 17

	// General category "Punctuation, Close" (Pe)
	UNICODE_CLOSE_PUNCTUATION UnicodeType = 18

	// General category "Punctuation, Final quote" (Pf)
	UNICODE_FINAL_PUNCTUATION UnicodeType = 19

	// General category "Punctuation, Initial quote" (Pi)
	UNICODE_INITIAL_PUNCTUATION UnicodeType = 20

	// General category "Punctuation, Other" (Po)
	UNICODE_OTHER_PUNCTUATION UnicodeType = 21

	// General category "Punctuation, Open" (Ps)
	UNICODE_OPEN_PUNCTUATION UnicodeType = 22

	// General category "Symbol, Currency" (Sc)
	UNICODE_CURRENCY_SYMBOL UnicodeType = 23

	// General category "Symbol, Modifier" (Sk)
	UNICODE_MODIFIER_SYMBOL UnicodeType = 24

	// General category "Symbol, Math" (Sm)
	UNICODE_MATH_SYMBOL UnicodeType = 25

	// General category "Symbol, Other" (So)
	UNICODE_OTHER_SYMBOL UnicodeType = 26

	// General category "Separator, Line" (Zl)
	UNICODE_LINE_SEPARATOR UnicodeType = 27

	// General category "Separator, Paragraph" (Zp)
	UNICODE_PARAGRAPH_SEPARATOR UnicodeType = 28

	// General category "Separator, Space" (Zs)
	UNICODE_SPACE_SEPARATOR UnicodeType = 29
)

// Error codes returned by parsing text-format GVariants.
type VariantParseError C.GVariantParseError

const (
	// generic error (unused)
	VARIANT_PARSE_ERROR_FAILED VariantParseError = 0

	// a non-basic #GVariantType was given where a basic type was expected
	VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED VariantParseError = 1

	// cannot infer the #GVariantType
	VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE VariantParseError = 2

	// an indefinite #GVariantType was given where a definite type was expected
	VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED VariantParseError = 3

	// extra data after parsing finished
	VARIANT_PARSE_ERROR_INPUT_NOT_AT_END VariantParseError = 4

	// invalid character in number or unicode escape
	VARIANT_PARSE_ERROR_INVALID_CHARACTER VariantParseError = 5

	// not a valid #GVariant format string
	VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING VariantParseError = 6

	// not a valid object path
	VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH VariantParseError = 7

	// not a valid type signature
	VARIANT_PARSE_ERROR_INVALID_SIGNATURE VariantParseError = 8

	// not a valid #GVariant type string
	VARIANT_PARSE_ERROR_INVALID_TYPE_STRING VariantParseError = 9

	// could not find a common type for array entries
	VARIANT_PARSE_ERROR_NO_COMMON_TYPE VariantParseError = 10

	// the numerical value is out of range of the given type
	VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE VariantParseError = 11

	// the numerical value is out of range for any type
	VARIANT_PARSE_ERROR_NUMBER_TOO_BIG VariantParseError = 12

	// cannot parse as variant of the specified type
	VARIANT_PARSE_ERROR_TYPE_ERROR VariantParseError = 13

	// an unexpected token was encountered
	VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN VariantParseError = 14

	// an unknown keyword was encountered
	VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD VariantParseError = 15

	// unterminated string constant
	VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT VariantParseError = 16

	// no value given
	VARIANT_PARSE_ERROR_VALUE_EXPECTED VariantParseError = 17
)

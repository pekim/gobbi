// This is a generated file - DO NOT EDIT

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

const ANALYZER_ANALYZING int = C.G_ANALYZER_ANALYZING

// A good size for a buffer to be passed into g_ascii_dtostr().
// It is guaranteed to be enough for all output of that function
// on systems with 64bit IEEE-compatible doubles.
//
// The typical usage would be something like:
// |[<!-- language="C" -->
// char buf[G_ASCII_DTOSTR_BUF_SIZE];
//
// fprintf (out, "value=%s\n", g_ascii_dtostr (buf, sizeof (buf), value));
// ]|
const ASCII_DTOSTR_BUF_SIZE int = C.G_ASCII_DTOSTR_BUF_SIZE

// Specifies one of the possible types of byte order.
// See #G_BYTE_ORDER.
const BIG_ENDIAN int = C.G_BIG_ENDIAN

// The set of uppercase ASCII alphabet characters.
// Used for specifying valid identifier characters
// in #GScannerConfig.
const CSET_A_2_Z string = C.G_CSET_A_2_Z

// The set of ASCII digits.
// Used for specifying valid identifier characters
// in #GScannerConfig.
const CSET_DIGITS string = C.G_CSET_DIGITS

// The set of lowercase ASCII alphabet characters.
// Used for specifying valid identifier characters
// in #GScannerConfig.
const CSET_a_2_z string = C.G_CSET_a_2_z

// A bitmask that restricts the possible flags passed to
// g_datalist_set_flags(). Passing a flags value where
// flags & ~G_DATALIST_FLAGS_MASK != 0 is an error.
const DATALIST_FLAGS_MASK int = C.G_DATALIST_FLAGS_MASK

// Represents an invalid #GDateDay.
const DATE_BAD_DAY int = C.G_DATE_BAD_DAY

// Represents an invalid Julian day number.
const DATE_BAD_JULIAN int = C.G_DATE_BAD_JULIAN

// Represents an invalid year.
const DATE_BAD_YEAR int = C.G_DATE_BAD_YEAR

// The directory separator character.
// This is '/' on UNIX machines and '\' under Windows.
const DIR_SEPARATOR int = C.G_DIR_SEPARATOR

// The directory separator as a string.
// This is "/" on UNIX machines and "\" under Windows.
const DIR_SEPARATOR_S string = C.G_DIR_SEPARATOR_S

// Unsupported : type gdouble for E

// This is the platform dependent conversion specifier for scanning and
// printing values of type #gint16. It is a string literal, but doesn't
// include the percent-sign, such that you can add precision and length
// modifiers between percent-sign and conversion specifier.
//
// |[<!-- language="C" -->
// gint16 in;
// gint32 out;
// sscanf ("42", "%" G_GINT16_FORMAT, &in)
// out = in * 1000;
// g_print ("%" G_GINT32_FORMAT, out);
// ]|
const GINT16_FORMAT string = C.G_GINT16_FORMAT

// This is the platform dependent conversion specifier for scanning
// and printing values of type #gint32. See also #G_GINT16_FORMAT.
const GINT32_FORMAT string = C.G_GINT32_FORMAT

// This is the platform dependent conversion specifier for scanning
// and printing values of type #gint64. See also #G_GINT16_FORMAT.
//
// Some platforms do not support scanning and printing 64-bit integers,
// even though the types are supported. On such platforms %G_GINT64_FORMAT
// is not defined. Note that scanf() may not support 64-bit integers, even
// if %G_GINT64_FORMAT is defined. Due to its weak error handling, scanf()
// is not recommended for parsing anyway; consider using g_ascii_strtoull()
// instead.
const GINT64_FORMAT string = C.G_GINT64_FORMAT

// Expands to "" on all modern compilers, and to  __FUNCTION__ on gcc
// version 2.x. Don't use it.
const GNUC_FUNCTION string = C.G_GNUC_FUNCTION

// Expands to "" on all modern compilers, and to __PRETTY_FUNCTION__
// on gcc version 2.x. Don't use it.
const GNUC_PRETTY_FUNCTION string = C.G_GNUC_PRETTY_FUNCTION

// This is the platform dependent conversion specifier for scanning
// and printing values of type #guint16. See also #G_GINT16_FORMAT
const GUINT16_FORMAT string = C.G_GUINT16_FORMAT

// This is the platform dependent conversion specifier for scanning
// and printing values of type #guint32. See also #G_GINT16_FORMAT.
const GUINT32_FORMAT string = C.G_GUINT32_FORMAT

// This is the platform dependent conversion specifier for scanning
// and printing values of type #guint64. See also #G_GINT16_FORMAT.
//
// Some platforms do not support scanning and printing 64-bit integers,
// even though the types are supported. On such platforms %G_GUINT64_FORMAT
// is not defined.  Note that scanf() may not support 64-bit integers, even
// if %G_GINT64_FORMAT is defined. Due to its weak error handling, scanf()
// is not recommended for parsing anyway; consider using g_ascii_strtoull()
// instead.
const GUINT64_FORMAT string = C.G_GUINT64_FORMAT
const HAVE_GINT64 int = C.G_HAVE_GINT64
const HAVE_GNUC_VARARGS int = C.G_HAVE_GNUC_VARARGS

// Defined to 1 if gcc-style visibility handling is supported.
const HAVE_GNUC_VISIBILITY int = C.G_HAVE_GNUC_VISIBILITY
const HAVE_GROWING_STACK int = C.G_HAVE_GROWING_STACK
const HAVE_ISO_VARARGS int = C.G_HAVE_ISO_VARARGS

// The position of the first bit which is not reserved for internal
// use be the #GHook implementation, i.e.
// `1 << G_HOOK_FLAG_USER_SHIFT` is the first
// bit which can be used for application-defined flags.
const HOOK_FLAG_USER_SHIFT int = C.G_HOOK_FLAG_USER_SHIFT

// The bias by which exponents in double-precision floats are offset.
const IEEE754_DOUBLE_BIAS int = C.G_IEEE754_DOUBLE_BIAS

// The bias by which exponents in single-precision floats are offset.
const IEEE754_FLOAT_BIAS int = C.G_IEEE754_FLOAT_BIAS

// Blacklisted : KEY_FILE_DESKTOP_ACTION_GROUP_PREFIX

const KEY_FILE_DESKTOP_KEY_FULLNAME string = C.G_KEY_FILE_DESKTOP_KEY_FULLNAME
const KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN string = C.G_KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN
const KEY_FILE_DESKTOP_KEY_KEYWORDS string = C.G_KEY_FILE_DESKTOP_KEY_KEYWORDS

// Specifies one of the possible types of byte order.
// See #G_BYTE_ORDER.
const LITTLE_ENDIAN int = C.G_LITTLE_ENDIAN

// Unsupported : type gdouble for LN10

// Unsupported : type gdouble for LN2

// Unsupported : type gdouble for LOG_2_BASE_10

// Unsupported : type gchar for LOG_DOMAIN

// GLib log levels that are considered fatal by default.
//
// This is not used if structured logging is enabled; see
// [Using Structured Logging][using-structured-logging].
const LOG_FATAL_MASK int = C.G_LOG_FATAL_MASK

// Log levels below 1<<G_LOG_LEVEL_USER_SHIFT are used by GLib.
// Higher bits can be used for user-defined log levels.
const LOG_LEVEL_USER_SHIFT int = C.G_LOG_LEVEL_USER_SHIFT

// The major version number of the GLib library.
//
// Like #glib_major_version, but from the headers used at
// application compile time, rather than from the library
// linked against at application run time.
const MAJOR_VERSION int = C.GLIB_MAJOR_VERSION

// Unsupported : type gint64 for MAXINT64

// Unsupported : type guint64 for MAXUINT64

// The micro version number of the GLib library.
//
// Like #gtk_micro_version, but from the headers used at
// application compile time, rather than from the library
// linked against at application run time.
const MICRO_VERSION int = C.GLIB_MICRO_VERSION

// Unsupported : type gint64 for MININT64

// The minor version number of the GLib library.
//
// Like #gtk_minor_version, but from the headers used at
// application compile time, rather than from the library
// linked against at application run time.
const MINOR_VERSION int = C.GLIB_MINOR_VERSION
const MODULE_SUFFIX string = C.G_MODULE_SUFFIX

// Specifies one of the possible types of byte order
// (currently unused). See #G_BYTE_ORDER.
const PDP_ENDIAN int = C.G_PDP_ENDIAN

// Unsupported : type gdouble for PI

// Unsupported : type gdouble for PI_2

// Unsupported : type gdouble for PI_4

// A format specifier that can be used in printf()-style format strings
// when printing the @fd member of a #GPollFD.
const POLLFD_FORMAT string = C.G_POLLFD_FORMAT

// Use this for default priority event sources.
//
// In GLib this priority is used when adding timeout functions
// with g_timeout_add(). In GDK this priority is used for events
// from the X server.
const PRIORITY_DEFAULT int = C.G_PRIORITY_DEFAULT

// Use this for default priority idle functions.
//
// In GLib this priority is used when adding idle functions with
// g_idle_add().
const PRIORITY_DEFAULT_IDLE int = C.G_PRIORITY_DEFAULT_IDLE

// Use this for high priority event sources.
//
// It is not used within GLib or GTK+.
const PRIORITY_HIGH int = C.G_PRIORITY_HIGH

// Use this for high priority idle functions.
//
// GTK+ uses #G_PRIORITY_HIGH_IDLE + 10 for resizing operations,
// and #G_PRIORITY_HIGH_IDLE + 20 for redrawing operations. (This is
// done to ensure that any pending resizes are processed before any
// pending redraws, so that widgets are not redrawn twice unnecessarily.)
const PRIORITY_HIGH_IDLE int = C.G_PRIORITY_HIGH_IDLE

// Use this for very low priority background tasks.
//
// It is not used within GLib or GTK+.
const PRIORITY_LOW int = C.G_PRIORITY_LOW

// The search path separator character.
// This is ':' on UNIX machines and ';' under Windows.
const SEARCHPATH_SEPARATOR int = C.G_SEARCHPATH_SEPARATOR

// The search path separator as a string.
// This is ":" on UNIX machines and ";" under Windows.
const SEARCHPATH_SEPARATOR_S string = C.G_SEARCHPATH_SEPARATOR_S
const SIZEOF_LONG int = C.GLIB_SIZEOF_LONG
const SIZEOF_SIZE_T int = C.GLIB_SIZEOF_SIZE_T
const SIZEOF_SSIZE_T int = C.GLIB_SIZEOF_SSIZE_T
const SIZEOF_VOID_P int = C.GLIB_SIZEOF_VOID_P

// Unsupported : type gdouble for SQRT2

// The standard delimiters, used in g_strdelimit().
const STR_DELIMITERS string = C.G_STR_DELIMITERS
const SYSDEF_AF_INET int = C.GLIB_SYSDEF_AF_INET
const SYSDEF_AF_INET6 int = C.GLIB_SYSDEF_AF_INET6
const SYSDEF_AF_UNIX int = C.GLIB_SYSDEF_AF_UNIX
const SYSDEF_MSG_DONTROUTE int = C.GLIB_SYSDEF_MSG_DONTROUTE
const SYSDEF_MSG_OOB int = C.GLIB_SYSDEF_MSG_OOB
const SYSDEF_MSG_PEEK int = C.GLIB_SYSDEF_MSG_PEEK

// Generic delimiters characters as defined in RFC 3986. Includes ":/?#[]@".
const URI_RESERVED_CHARS_GENERIC_DELIMITERS string = C.G_URI_RESERVED_CHARS_GENERIC_DELIMITERS

// Subcomponent delimiter characters as defined in RFC 3986. Includes "!$&'()*+,;=".
const URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS string = C.G_URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS

// Number of microseconds in one second (1 million).
// This macro is provided for code readability.
const USEC_PER_SEC int = C.G_USEC_PER_SEC
const VA_COPY_AS_ARRAY int = C.G_VA_COPY_AS_ARRAY

// Blacklisted : WIN32_MSG_HANDLE

# `glib` Constants

## `ANALYZER_ANALYZING`



C - `G_ANALYZER_ANALYZING`

## `ASCII_DTOSTR_BUF_SIZE`

A good size for a buffer to be passed into g_ascii_dtostr().
It is guaranteed to be enough for all output of that function
on systems with 64bit IEEE-compatible doubles.

The typical usage would be something like:
|[<!-- language="C" -->
  char buf[G_ASCII_DTOSTR_BUF_SIZE];

  fprintf (out, "value=%s\n", g_ascii_dtostr (buf, sizeof (buf), value));
]|

C - `G_ASCII_DTOSTR_BUF_SIZE`

## `BIG_ENDIAN`

Specifies one of the possible types of byte order.
See #G_BYTE_ORDER.

C - `G_BIG_ENDIAN`

## `CSET_A_2_Z`

The set of uppercase ASCII alphabet characters.
Used for specifying valid identifier characters
in #GScannerConfig.

C - `G_CSET_A_2_Z`

## `CSET_DIGITS`

The set of ASCII digits.
Used for specifying valid identifier characters
in #GScannerConfig.

C - `G_CSET_DIGITS`

## `CSET_a_2_z`

The set of lowercase ASCII alphabet characters.
Used for specifying valid identifier characters
in #GScannerConfig.

C - `G_CSET_a_2_z`

## `DATALIST_FLAGS_MASK`

A bitmask that restricts the possible flags passed to
g_datalist_set_flags(). Passing a flags value where
flags & ~G_DATALIST_FLAGS_MASK != 0 is an error.

C - `G_DATALIST_FLAGS_MASK`

## `DATE_BAD_DAY`

Represents an invalid #GDateDay.

C - `G_DATE_BAD_DAY`

## `DATE_BAD_JULIAN`

Represents an invalid Julian day number.

C - `G_DATE_BAD_JULIAN`

## `DATE_BAD_YEAR`

Represents an invalid year.

C - `G_DATE_BAD_YEAR`

## `DIR_SEPARATOR`

The directory separator character.
This is '/' on UNIX machines and '\' under Windows.

C - `G_DIR_SEPARATOR`

## `DIR_SEPARATOR_S`

The directory separator as a string.
This is "/" on UNIX machines and "\" under Windows.

C - `G_DIR_SEPARATOR_S`

## `GINT16_FORMAT`

This is the platform dependent conversion specifier for scanning and
printing values of type #gint16. It is a string literal, but doesn't
include the percent-sign, such that you can add precision and length
modifiers between percent-sign and conversion specifier.

|[<!-- language="C" -->
gint16 in;
gint32 out;
sscanf ("42", "%" G_GINT16_FORMAT, &in)
out = in * 1000;
g_print ("%" G_GINT32_FORMAT, out);
]|

C - `G_GINT16_FORMAT`

## `GINT16_MODIFIER`

The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint16 or #guint16. It
is a string literal, but doesn't include the percent-sign, such
that you can add precision and length modifiers between percent-sign
and conversion specifier and append a conversion specifier.

The following example prints "0x7b";
|[<!-- language="C" -->
gint16 value = 123;
g_print ("%#" G_GINT16_MODIFIER "x", value);
]|

C - `G_GINT16_MODIFIER`

## `GINT32_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #gint32. See also #G_GINT16_FORMAT.

C - `G_GINT32_FORMAT`

## `GINT32_MODIFIER`

The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint32 or #guint32. It
is a string literal. See also #G_GINT16_MODIFIER.

C - `G_GINT32_MODIFIER`

## `GINT64_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #gint64. See also #G_GINT16_FORMAT.

Some platforms do not support scanning and printing 64-bit integers,
even though the types are supported. On such platforms %G_GINT64_FORMAT
is not defined. Note that scanf() may not support 64-bit integers, even
if %G_GINT64_FORMAT is defined. Due to its weak error handling, scanf()
is not recommended for parsing anyway; consider using g_ascii_strtoull()
instead.

C - `G_GINT64_FORMAT`

## `GINT64_MODIFIER`

The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint64 or #guint64.
It is a string literal.

Some platforms do not support printing 64-bit integers, even
though the types are supported. On such platforms %G_GINT64_MODIFIER
is not defined.

C - `G_GINT64_MODIFIER`

## `GINTPTR_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #gintptr.

C - `G_GINTPTR_FORMAT`

## `GINTPTR_MODIFIER`

The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gintptr or #guintptr.
It is a string literal.

C - `G_GINTPTR_MODIFIER`

## `GNUC_FUNCTION`

Expands to "" on all modern compilers, and to  __FUNCTION__ on gcc
version 2.x. Don't use it.

C - `G_GNUC_FUNCTION`

## `GNUC_PRETTY_FUNCTION`

Expands to "" on all modern compilers, and to __PRETTY_FUNCTION__
on gcc version 2.x. Don't use it.

C - `G_GNUC_PRETTY_FUNCTION`

## `GSIZE_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #gsize. See also #G_GINT16_FORMAT.

C - `G_GSIZE_FORMAT`

## `GSIZE_MODIFIER`

The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gsize. It
is a string literal.

C - `G_GSIZE_MODIFIER`

## `GSSIZE_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #gssize. See also #G_GINT16_FORMAT.

C - `G_GSSIZE_FORMAT`

## `GSSIZE_MODIFIER`

The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gssize. It
is a string literal.

C - `G_GSSIZE_MODIFIER`

## `GUINT16_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #guint16. See also #G_GINT16_FORMAT

C - `G_GUINT16_FORMAT`

## `GUINT32_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #guint32. See also #G_GINT16_FORMAT.

C - `G_GUINT32_FORMAT`

## `GUINT64_FORMAT`

This is the platform dependent conversion specifier for scanning
and printing values of type #guint64. See also #G_GINT16_FORMAT.

Some platforms do not support scanning and printing 64-bit integers,
even though the types are supported. On such platforms %G_GUINT64_FORMAT
is not defined.  Note that scanf() may not support 64-bit integers, even
if %G_GINT64_FORMAT is defined. Due to its weak error handling, scanf()
is not recommended for parsing anyway; consider using g_ascii_strtoull()
instead.

C - `G_GUINT64_FORMAT`

## `GUINTPTR_FORMAT`

This is the platform dependent conversion specifier
for scanning and printing values of type #guintptr.

C - `G_GUINTPTR_FORMAT`

## `HAVE_GINT64`



C - `G_HAVE_GINT64`

## `HAVE_GNUC_VARARGS`



C - `G_HAVE_GNUC_VARARGS`

## `HAVE_GNUC_VISIBILITY`

Defined to 1 if gcc-style visibility handling is supported.

C - `G_HAVE_GNUC_VISIBILITY`

## `HAVE_GROWING_STACK`



C - `G_HAVE_GROWING_STACK`

## `HAVE_ISO_VARARGS`



C - `G_HAVE_ISO_VARARGS`

## `HOOK_FLAG_USER_SHIFT`

The position of the first bit which is not reserved for internal
use be the #GHook implementation, i.e.
`1 << G_HOOK_FLAG_USER_SHIFT` is the first
bit which can be used for application-defined flags.

C - `G_HOOK_FLAG_USER_SHIFT`

## `IEEE754_DOUBLE_BIAS`

The bias by which exponents in double-precision floats are offset.

C - `G_IEEE754_DOUBLE_BIAS`

## `IEEE754_FLOAT_BIAS`

The bias by which exponents in single-precision floats are offset.

C - `G_IEEE754_FLOAT_BIAS`

## `KEY_FILE_DESKTOP_GROUP`

The name of the main group of a desktop entry file, as defined in the
[Desktop Entry Specification](http://freedesktop.org/Standards/desktop-entry-spec).
Consult the specification for more
details about the meanings of the keys below.

C - `G_KEY_FILE_DESKTOP_GROUP`

## `KEY_FILE_DESKTOP_KEY_ACTIONS`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string list
giving the available application actions.

C - `G_KEY_FILE_DESKTOP_KEY_ACTIONS`

## `KEY_FILE_DESKTOP_KEY_CATEGORIES`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list
of strings giving the categories in which the desktop entry
should be shown in a menu.

C - `G_KEY_FILE_DESKTOP_KEY_CATEGORIES`

## `KEY_FILE_DESKTOP_KEY_COMMENT`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the tooltip for the desktop entry.

C - `G_KEY_FILE_DESKTOP_KEY_COMMENT`

## `KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean set to true
if the application is D-Bus activatable.

C - `G_KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE`

## `KEY_FILE_DESKTOP_KEY_EXEC`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the command line to execute. It is only valid for desktop
entries with the `Application` type.

C - `G_KEY_FILE_DESKTOP_KEY_EXEC`

## `KEY_FILE_DESKTOP_KEY_FULLNAME`



C - `G_KEY_FILE_DESKTOP_KEY_FULLNAME`

## `KEY_FILE_DESKTOP_KEY_GENERIC_NAME`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the generic name of the desktop entry.

C - `G_KEY_FILE_DESKTOP_KEY_GENERIC_NAME`

## `KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN`



C - `G_KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN`

## `KEY_FILE_DESKTOP_KEY_HIDDEN`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the desktop entry has been deleted by the user.

C - `G_KEY_FILE_DESKTOP_KEY_HIDDEN`

## `KEY_FILE_DESKTOP_KEY_ICON`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the name of the icon to be displayed for the desktop
entry.

C - `G_KEY_FILE_DESKTOP_KEY_ICON`

## `KEY_FILE_DESKTOP_KEY_KEYWORDS`



C - `G_KEY_FILE_DESKTOP_KEY_KEYWORDS`

## `KEY_FILE_DESKTOP_KEY_MIME_TYPE`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list
of strings giving the MIME types supported by this desktop entry.

C - `G_KEY_FILE_DESKTOP_KEY_MIME_TYPE`

## `KEY_FILE_DESKTOP_KEY_NAME`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the specific name of the desktop entry.

C - `G_KEY_FILE_DESKTOP_KEY_NAME`

## `KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list of
strings identifying the environments that should not display the
desktop entry.

C - `G_KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN`

## `KEY_FILE_DESKTOP_KEY_NO_DISPLAY`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the desktop entry should be shown in menus.

C - `G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY`

## `KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list of
strings identifying the environments that should display the
desktop entry.

C - `G_KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN`

## `KEY_FILE_DESKTOP_KEY_PATH`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
containing the working directory to run the program in. It is only
valid for desktop entries with the `Application` type.

C - `G_KEY_FILE_DESKTOP_KEY_PATH`

## `KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the application supports the
[Startup Notification Protocol Specification](http://www.freedesktop.org/Standards/startup-notification-spec).

C - `G_KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY`

## `KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is string
identifying the WM class or name hint of a window that the application
will create, which can be used to emulate Startup Notification with
older applications.

C - `G_KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS`

## `KEY_FILE_DESKTOP_KEY_TERMINAL`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the program should be run in a terminal window.
It is only valid for desktop entries with the
`Application` type.

C - `G_KEY_FILE_DESKTOP_KEY_TERMINAL`

## `KEY_FILE_DESKTOP_KEY_TRY_EXEC`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the file name of a binary on disk used to determine if the
program is actually installed. It is only valid for desktop entries
with the `Application` type.

C - `G_KEY_FILE_DESKTOP_KEY_TRY_EXEC`

## `KEY_FILE_DESKTOP_KEY_TYPE`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the type of the desktop entry. Usually
&num;G_KEY_FILE_DESKTOP_TYPE_APPLICATION,
&num;G_KEY_FILE_DESKTOP_TYPE_LINK, or
&num;G_KEY_FILE_DESKTOP_TYPE_DIRECTORY.

C - `G_KEY_FILE_DESKTOP_KEY_TYPE`

## `KEY_FILE_DESKTOP_KEY_URL`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the URL to access. It is only valid for desktop entries
with the `Link` type.

C - `G_KEY_FILE_DESKTOP_KEY_URL`

## `KEY_FILE_DESKTOP_KEY_VERSION`

A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the version of the Desktop Entry Specification used for
the desktop entry file.

C - `G_KEY_FILE_DESKTOP_KEY_VERSION`

## `KEY_FILE_DESKTOP_TYPE_APPLICATION`

The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
entries representing applications.

C - `G_KEY_FILE_DESKTOP_TYPE_APPLICATION`

## `KEY_FILE_DESKTOP_TYPE_DIRECTORY`

The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
entries representing directories.

C - `G_KEY_FILE_DESKTOP_TYPE_DIRECTORY`

## `KEY_FILE_DESKTOP_TYPE_LINK`

The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
entries representing links to documents.

C - `G_KEY_FILE_DESKTOP_TYPE_LINK`

## `LITTLE_ENDIAN`

Specifies one of the possible types of byte order.
See #G_BYTE_ORDER.

C - `G_LITTLE_ENDIAN`

## `LOG_FATAL_MASK`

GLib log levels that are considered fatal by default.

This is not used if structured logging is enabled; see
[Using Structured Logging][using-structured-logging].

C - `G_LOG_FATAL_MASK`

## `LOG_LEVEL_USER_SHIFT`

Log levels below 1<<G_LOG_LEVEL_USER_SHIFT are used by GLib.
Higher bits can be used for user-defined log levels.

C - `G_LOG_LEVEL_USER_SHIFT`

## `MAJOR_VERSION`

The major version number of the GLib library.

Like #glib_major_version, but from the headers used at
application compile time, rather than from the library
linked against at application run time.

C - `GLIB_MAJOR_VERSION`

## `MICRO_VERSION`

The micro version number of the GLib library.

Like #gtk_micro_version, but from the headers used at
application compile time, rather than from the library
linked against at application run time.

C - `GLIB_MICRO_VERSION`

## `MINOR_VERSION`

The minor version number of the GLib library.

Like #gtk_minor_version, but from the headers used at
application compile time, rather than from the library
linked against at application run time.

C - `GLIB_MINOR_VERSION`

## `MODULE_SUFFIX`



C - `G_MODULE_SUFFIX`

## `OPTION_REMAINING`

If a long option in the main group has this name, it is not treated as a
regular option. Instead it collects all non-option arguments which would
otherwise be left in `argv`. The option must be of type
%G_OPTION_ARG_CALLBACK, %G_OPTION_ARG_STRING_ARRAY
or %G_OPTION_ARG_FILENAME_ARRAY.


Using #G_OPTION_REMAINING instead of simply scanning `argv`
for leftover arguments has the advantage that GOption takes care of
necessary encoding conversions for strings or filenames.

C - `G_OPTION_REMAINING`

## `PDP_ENDIAN`

Specifies one of the possible types of byte order
(currently unused). See #G_BYTE_ORDER.

C - `G_PDP_ENDIAN`

## `PID_FORMAT`

A format specifier that can be used in printf()-style format strings
when printing a #GPid.

C - `G_PID_FORMAT`

## `POLLFD_FORMAT`

A format specifier that can be used in printf()-style format strings
when printing the @fd member of a #GPollFD.

C - `G_POLLFD_FORMAT`

## `PRIORITY_DEFAULT`

Use this for default priority event sources.

In GLib this priority is used when adding timeout functions
with g_timeout_add(). In GDK this priority is used for events
from the X server.

C - `G_PRIORITY_DEFAULT`

## `PRIORITY_DEFAULT_IDLE`

Use this for default priority idle functions.

In GLib this priority is used when adding idle functions with
g_idle_add().

C - `G_PRIORITY_DEFAULT_IDLE`

## `PRIORITY_HIGH`

Use this for high priority event sources.

It is not used within GLib or GTK+.

C - `G_PRIORITY_HIGH`

## `PRIORITY_HIGH_IDLE`

Use this for high priority idle functions.

GTK+ uses #G_PRIORITY_HIGH_IDLE + 10 for resizing operations,
and #G_PRIORITY_HIGH_IDLE + 20 for redrawing operations. (This is
done to ensure that any pending resizes are processed before any
pending redraws, so that widgets are not redrawn twice unnecessarily.)

C - `G_PRIORITY_HIGH_IDLE`

## `PRIORITY_LOW`

Use this for very low priority background tasks.

It is not used within GLib or GTK+.

C - `G_PRIORITY_LOW`

## `SEARCHPATH_SEPARATOR`

The search path separator character.
This is ':' on UNIX machines and ';' under Windows.

C - `G_SEARCHPATH_SEPARATOR`

## `SEARCHPATH_SEPARATOR_S`

The search path separator as a string.
This is ":" on UNIX machines and ";" under Windows.

C - `G_SEARCHPATH_SEPARATOR_S`

## `SIZEOF_LONG`



C - `GLIB_SIZEOF_LONG`

## `SIZEOF_SIZE_T`



C - `GLIB_SIZEOF_SIZE_T`

## `SIZEOF_SSIZE_T`



C - `GLIB_SIZEOF_SSIZE_T`

## `SIZEOF_VOID_P`



C - `GLIB_SIZEOF_VOID_P`

## `SOURCE_CONTINUE`

Use this macro as the return value of a #GSourceFunc to leave
the #GSource in the main loop.

C - `G_SOURCE_CONTINUE`

## `SOURCE_REMOVE`

Use this macro as the return value of a #GSourceFunc to remove
the #GSource from the main loop.

C - `G_SOURCE_REMOVE`

## `STR_DELIMITERS`

The standard delimiters, used in g_strdelimit().

C - `G_STR_DELIMITERS`

## `SYSDEF_AF_INET`



C - `GLIB_SYSDEF_AF_INET`

## `SYSDEF_AF_INET6`



C - `GLIB_SYSDEF_AF_INET6`

## `SYSDEF_AF_UNIX`



C - `GLIB_SYSDEF_AF_UNIX`

## `SYSDEF_MSG_DONTROUTE`



C - `GLIB_SYSDEF_MSG_DONTROUTE`

## `SYSDEF_MSG_OOB`



C - `GLIB_SYSDEF_MSG_OOB`

## `SYSDEF_MSG_PEEK`



C - `GLIB_SYSDEF_MSG_PEEK`

## `UNICHAR_MAX_DECOMPOSITION_LENGTH`

The maximum length (in codepoints) of a compatibility or canonical
decomposition of a single Unicode character.

This is as defined by Unicode 6.1.

C - `G_UNICHAR_MAX_DECOMPOSITION_LENGTH`

## `URI_RESERVED_CHARS_GENERIC_DELIMITERS`

Generic delimiters characters as defined in RFC 3986. Includes ":/?#[]@".

C - `G_URI_RESERVED_CHARS_GENERIC_DELIMITERS`

## `URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS`

Subcomponent delimiter characters as defined in RFC 3986. Includes "!$&'()*+,;=".

C - `G_URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS`

## `USEC_PER_SEC`

Number of microseconds in one second (1 million).
This macro is provided for code readability.

C - `G_USEC_PER_SEC`

## `VA_COPY_AS_ARRAY`



C - `G_VA_COPY_AS_ARRAY`

## `VERSION_MIN_REQUIRED`

A macro that should be defined by the user prior to including
the glib.h header.
The definition should be one of the predefined GLib version
macros: %GLIB_VERSION_2_26, %GLIB_VERSION_2_28,...

This macro defines the earliest version of GLib that the package is
required to be able to compile against.

If the compiler is configured to warn about the use of deprecated
functions, then using functions that were deprecated in version
%GLIB_VERSION_MIN_REQUIRED or earlier will cause warnings (but
using functions deprecated in later releases will not).

C - `GLIB_VERSION_MIN_REQUIRED`


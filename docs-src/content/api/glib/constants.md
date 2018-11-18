+++
title = "constants"
+++
<p class="api-heading">ANALYZER_ANALYZING</p>
<div class="api-notes">
  <p class="api-ctype">G_ANALYZER_ANALYZING</p>
</div>
<p class="api-heading">ASCII_DTOSTR_BUF_SIZE</p>
<p class="api-doc">A good size for a buffer to be passed into g_ascii_dtostr().
It is guaranteed to be enough for all output of that function
on systems with 64bit IEEE-compatible doubles.

The typical usage would be something like:
|[<!-- language="C" -->
  char buf[G_ASCII_DTOSTR_BUF_SIZE];

  fprintf (out, "value=%s\n", g_ascii_dtostr (buf, sizeof (buf), value));
]|</p>
<div class="api-notes">
  <p class="api-ctype">G_ASCII_DTOSTR_BUF_SIZE</p>
</div>
<p class="api-heading">BIG_ENDIAN</p>
<p class="api-doc">Specifies one of the possible types of byte order.
See #G_BYTE_ORDER.</p>
<div class="api-notes">
  <p class="api-ctype">G_BIG_ENDIAN</p>
</div>
<p class="api-heading">CSET_A_2_Z</p>
<p class="api-doc">The set of uppercase ASCII alphabet characters.
Used for specifying valid identifier characters
in #GScannerConfig.</p>
<div class="api-notes">
  <p class="api-ctype">G_CSET_A_2_Z</p>
</div>
<p class="api-heading">CSET_DIGITS</p>
<p class="api-doc">The set of ASCII digits.
Used for specifying valid identifier characters
in #GScannerConfig.</p>
<div class="api-notes">
  <p class="api-ctype">G_CSET_DIGITS</p>
</div>
<p class="api-heading">CSET_a_2_z</p>
<p class="api-doc">The set of lowercase ASCII alphabet characters.
Used for specifying valid identifier characters
in #GScannerConfig.</p>
<div class="api-notes">
  <p class="api-ctype">G_CSET_a_2_z</p>
</div>
<p class="api-heading">DATALIST_FLAGS_MASK</p>
<p class="api-doc">A bitmask that restricts the possible flags passed to
g_datalist_set_flags(). Passing a flags value where
flags & ~G_DATALIST_FLAGS_MASK != 0 is an error.</p>
<div class="api-notes">
  <p class="api-ctype">G_DATALIST_FLAGS_MASK</p>
</div>
<p class="api-heading">DATE_BAD_DAY</p>
<p class="api-doc">Represents an invalid #GDateDay.</p>
<div class="api-notes">
  <p class="api-ctype">G_DATE_BAD_DAY</p>
</div>
<p class="api-heading">DATE_BAD_JULIAN</p>
<p class="api-doc">Represents an invalid Julian day number.</p>
<div class="api-notes">
  <p class="api-ctype">G_DATE_BAD_JULIAN</p>
</div>
<p class="api-heading">DATE_BAD_YEAR</p>
<p class="api-doc">Represents an invalid year.</p>
<div class="api-notes">
  <p class="api-ctype">G_DATE_BAD_YEAR</p>
</div>
<p class="api-heading">DIR_SEPARATOR</p>
<p class="api-doc">The directory separator character.
This is '/' on UNIX machines and '\' under Windows.</p>
<div class="api-notes">
  <p class="api-ctype">G_DIR_SEPARATOR</p>
</div>
<p class="api-heading">DIR_SEPARATOR_S</p>
<p class="api-doc">The directory separator as a string.
This is "/" on UNIX machines and "\" under Windows.</p>
<div class="api-notes">
  <p class="api-ctype">G_DIR_SEPARATOR_S</p>
</div>
<p class="api-heading">E</p>
<p class="api-doc">The base of natural logarithms.</p>
<div class="api-notes">
  <p class="api-ctype">G_E</p>
</div>
<p class="api-heading">GINT16_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning and
printing values of type #gint16. It is a string literal, but doesn't
include the percent-sign, such that you can add precision and length
modifiers between percent-sign and conversion specifier.

|[<!-- language="C" -->
gint16 in;
gint32 out;
sscanf ("42", "%" G_GINT16_FORMAT, &in)
out = in * 1000;
g_print ("%" G_GINT32_FORMAT, out);
]|</p>
<div class="api-notes">
  <p class="api-ctype">G_GINT16_FORMAT</p>
</div>
<p class="api-heading">GINT16_MODIFIER</p>
<p class="api-doc">The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint16 or #guint16. It
is a string literal, but doesn't include the percent-sign, such
that you can add precision and length modifiers between percent-sign
and conversion specifier and append a conversion specifier.

The following example prints "0x7b";
|[<!-- language="C" -->
gint16 value = 123;
g_print ("%#" G_GINT16_MODIFIER "x", value);
]|</p>
<div class="api-notes">
  <p class="api-ctype">G_GINT16_MODIFIER</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">GINT32_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #gint32. See also #G_GINT16_FORMAT.</p>
<div class="api-notes">
  <p class="api-ctype">G_GINT32_FORMAT</p>
</div>
<p class="api-heading">GINT32_MODIFIER</p>
<p class="api-doc">The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint32 or #guint32. It
is a string literal. See also #G_GINT16_MODIFIER.</p>
<div class="api-notes">
  <p class="api-ctype">G_GINT32_MODIFIER</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">GINT64_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #gint64. See also #G_GINT16_FORMAT.

Some platforms do not support scanning and printing 64-bit integers,
even though the types are supported. On such platforms %G_GINT64_FORMAT
is not defined. Note that scanf() may not support 64-bit integers, even
if %G_GINT64_FORMAT is defined. Due to its weak error handling, scanf()
is not recommended for parsing anyway; consider using g_ascii_strtoull()
instead.</p>
<div class="api-notes">
  <p class="api-ctype">G_GINT64_FORMAT</p>
</div>
<p class="api-heading">GINT64_MODIFIER</p>
<p class="api-doc">The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint64 or #guint64.
It is a string literal.

Some platforms do not support printing 64-bit integers, even
though the types are supported. On such platforms %G_GINT64_MODIFIER
is not defined.</p>
<div class="api-notes">
  <p class="api-ctype">G_GINT64_MODIFIER</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">GINTPTR_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #gintptr.</p>
<div class="api-notes">
  <p class="api-ctype">G_GINTPTR_FORMAT</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">GINTPTR_MODIFIER</p>
<p class="api-doc">The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gintptr or #guintptr.
It is a string literal.</p>
<div class="api-notes">
  <p class="api-ctype">G_GINTPTR_MODIFIER</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">GNUC_FUNCTION</p>
<p class="api-doc">Expands to "" on all modern compilers, and to  __FUNCTION__ on gcc
version 2.x. Don't use it.</p>
<div class="api-notes">
  <p class="api-ctype">G_GNUC_FUNCTION</p>
</div>
<p class="api-heading">GNUC_PRETTY_FUNCTION</p>
<p class="api-doc">Expands to "" on all modern compilers, and to __PRETTY_FUNCTION__
on gcc version 2.x. Don't use it.</p>
<div class="api-notes">
  <p class="api-ctype">G_GNUC_PRETTY_FUNCTION</p>
</div>
<p class="api-heading">GSIZE_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #gsize. See also #G_GINT16_FORMAT.</p>
<div class="api-notes">
  <p class="api-ctype">G_GSIZE_FORMAT</p>
  <p class="api-since">since 2.6</p>
</div>
<p class="api-heading">GSIZE_MODIFIER</p>
<p class="api-doc">The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gsize. It
is a string literal.</p>
<div class="api-notes">
  <p class="api-ctype">G_GSIZE_MODIFIER</p>
  <p class="api-since">since 2.6</p>
</div>
<p class="api-heading">GSSIZE_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #gssize. See also #G_GINT16_FORMAT.</p>
<div class="api-notes">
  <p class="api-ctype">G_GSSIZE_FORMAT</p>
  <p class="api-since">since 2.6</p>
</div>
<p class="api-heading">GSSIZE_MODIFIER</p>
<p class="api-doc">The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gssize. It
is a string literal.</p>
<div class="api-notes">
  <p class="api-ctype">G_GSSIZE_MODIFIER</p>
  <p class="api-since">since 2.6</p>
</div>
<p class="api-heading">GUINT16_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #guint16. See also #G_GINT16_FORMAT</p>
<div class="api-notes">
  <p class="api-ctype">G_GUINT16_FORMAT</p>
</div>
<p class="api-heading">GUINT32_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #guint32. See also #G_GINT16_FORMAT.</p>
<div class="api-notes">
  <p class="api-ctype">G_GUINT32_FORMAT</p>
</div>
<p class="api-heading">GUINT64_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier for scanning
and printing values of type #guint64. See also #G_GINT16_FORMAT.

Some platforms do not support scanning and printing 64-bit integers,
even though the types are supported. On such platforms %G_GUINT64_FORMAT
is not defined.  Note that scanf() may not support 64-bit integers, even
if %G_GINT64_FORMAT is defined. Due to its weak error handling, scanf()
is not recommended for parsing anyway; consider using g_ascii_strtoull()
instead.</p>
<div class="api-notes">
  <p class="api-ctype">G_GUINT64_FORMAT</p>
</div>
<p class="api-heading">GUINTPTR_FORMAT</p>
<p class="api-doc">This is the platform dependent conversion specifier
for scanning and printing values of type #guintptr.</p>
<div class="api-notes">
  <p class="api-ctype">G_GUINTPTR_FORMAT</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">HAVE_GINT64</p>
<div class="api-notes">
  <p class="api-ctype">G_HAVE_GINT64</p>
</div>
<p class="api-heading">HAVE_GNUC_VARARGS</p>
<div class="api-notes">
  <p class="api-ctype">G_HAVE_GNUC_VARARGS</p>
</div>
<p class="api-heading">HAVE_GNUC_VISIBILITY</p>
<p class="api-doc">Defined to 1 if gcc-style visibility handling is supported.</p>
<div class="api-notes">
  <p class="api-ctype">G_HAVE_GNUC_VISIBILITY</p>
</div>
<p class="api-heading">HAVE_GROWING_STACK</p>
<div class="api-notes">
  <p class="api-ctype">G_HAVE_GROWING_STACK</p>
</div>
<p class="api-heading">HAVE_ISO_VARARGS</p>
<div class="api-notes">
  <p class="api-ctype">G_HAVE_ISO_VARARGS</p>
</div>
<p class="api-heading">HOOK_FLAG_USER_SHIFT</p>
<p class="api-doc">The position of the first bit which is not reserved for internal
use be the #GHook implementation, i.e.
`1 << G_HOOK_FLAG_USER_SHIFT` is the first
bit which can be used for application-defined flags.</p>
<div class="api-notes">
  <p class="api-ctype">G_HOOK_FLAG_USER_SHIFT</p>
</div>
<p class="api-heading">IEEE754_DOUBLE_BIAS</p>
<p class="api-doc">The bias by which exponents in double-precision floats are offset.</p>
<div class="api-notes">
  <p class="api-ctype">G_IEEE754_DOUBLE_BIAS</p>
</div>
<p class="api-heading">IEEE754_FLOAT_BIAS</p>
<p class="api-doc">The bias by which exponents in single-precision floats are offset.</p>
<div class="api-notes">
  <p class="api-ctype">G_IEEE754_FLOAT_BIAS</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_GROUP</p>
<p class="api-doc">The name of the main group of a desktop entry file, as defined in the
[Desktop Entry Specification](http://freedesktop.org/Standards/desktop-entry-spec).
Consult the specification for more
details about the meanings of the keys below.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_GROUP</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_ACTIONS</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string list
giving the available application actions.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_ACTIONS</p>
  <p class="api-since">since 2.38</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_CATEGORIES</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list
of strings giving the categories in which the desktop entry
should be shown in a menu.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_CATEGORIES</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_COMMENT</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the tooltip for the desktop entry.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_COMMENT</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean set to true
if the application is D-Bus activatable.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE</p>
  <p class="api-since">since 2.38</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_EXEC</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the command line to execute. It is only valid for desktop
entries with the `Application` type.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_EXEC</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_FULLNAME</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_FULLNAME</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_GENERIC_NAME</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the generic name of the desktop entry.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_GENERIC_NAME</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_HIDDEN</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the desktop entry has been deleted by the user.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_HIDDEN</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_ICON</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the name of the icon to be displayed for the desktop
entry.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_ICON</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_KEYWORDS</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_KEYWORDS</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_MIME_TYPE</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list
of strings giving the MIME types supported by this desktop entry.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_MIME_TYPE</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_NAME</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a localized
string giving the specific name of the desktop entry.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_NAME</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list of
strings identifying the environments that should not display the
desktop entry.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_NO_DISPLAY</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the desktop entry should be shown in menus.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_NO_DISPLAY</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a list of
strings identifying the environments that should display the
desktop entry.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_PATH</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
containing the working directory to run the program in. It is only
valid for desktop entries with the `Application` type.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_PATH</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the application supports the
[Startup Notification Protocol Specification](http://www.freedesktop.org/Standards/startup-notification-spec).</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is string
identifying the WM class or name hint of a window that the application
will create, which can be used to emulate Startup Notification with
older applications.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_TERMINAL</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a boolean
stating whether the program should be run in a terminal window.
It is only valid for desktop entries with the
`Application` type.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_TERMINAL</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_TRY_EXEC</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the file name of a binary on disk used to determine if the
program is actually installed. It is only valid for desktop entries
with the `Application` type.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_TRY_EXEC</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_TYPE</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the type of the desktop entry. Usually
#G_KEY_FILE_DESKTOP_TYPE_APPLICATION,
#G_KEY_FILE_DESKTOP_TYPE_LINK, or
#G_KEY_FILE_DESKTOP_TYPE_DIRECTORY.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_TYPE</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_URL</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the URL to access. It is only valid for desktop entries
with the `Link` type.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_URL</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_KEY_VERSION</p>
<p class="api-doc">A key under #G_KEY_FILE_DESKTOP_GROUP, whose value is a string
giving the version of the Desktop Entry Specification used for
the desktop entry file.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_KEY_VERSION</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_TYPE_APPLICATION</p>
<p class="api-doc">The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
entries representing applications.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_TYPE_APPLICATION</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_TYPE_DIRECTORY</p>
<p class="api-doc">The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
entries representing directories.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_TYPE_DIRECTORY</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">KEY_FILE_DESKTOP_TYPE_LINK</p>
<p class="api-doc">The value of the #G_KEY_FILE_DESKTOP_KEY_TYPE, key for desktop
entries representing links to documents.</p>
<div class="api-notes">
  <p class="api-ctype">G_KEY_FILE_DESKTOP_TYPE_LINK</p>
  <p class="api-since">since 2.14</p>
</div>
<p class="api-heading">LITTLE_ENDIAN</p>
<p class="api-doc">Specifies one of the possible types of byte order.
See #G_BYTE_ORDER.</p>
<div class="api-notes">
  <p class="api-ctype">G_LITTLE_ENDIAN</p>
</div>
<p class="api-heading">LN10</p>
<p class="api-doc">The natural logarithm of 10.</p>
<div class="api-notes">
  <p class="api-ctype">G_LN10</p>
</div>
<p class="api-heading">LN2</p>
<p class="api-doc">The natural logarithm of 2.</p>
<div class="api-notes">
  <p class="api-ctype">G_LN2</p>
</div>
<p class="api-heading">LOG_2_BASE_10</p>
<p class="api-doc">Multiplying the base 2 exponent by this number yields the base 10 exponent.</p>
<div class="api-notes">
  <p class="api-ctype">G_LOG_2_BASE_10</p>
</div>
<p class="api-heading">LOG_DOMAIN</p>
<p class="api-doc">Defines the log domain. See [Log Domains](#log-domains).

Libraries should define this so that any messages
which they log can be differentiated from messages from other
libraries and application code. But be careful not to define
it in any public header files.

Log domains must be unique, and it is recommended that they are the
application or library name, optionally followed by a hyphen and a sub-domain
name. For example, `bloatpad` or `bloatpad-io`.

If undefined, it defaults to the default %NULL (or `""`) log domain; this is
not advisable, as it cannot be filtered against using the `G_MESSAGES_DEBUG`
environment variable.

For example, GTK+ uses this in its `Makefile.am`:
|[
AM_CPPFLAGS = -DG_LOG_DOMAIN=\"Gtk\"
]|

Applications can choose to leave it as the default %NULL (or `""`)
domain. However, defining the domain offers the same advantages as
above.</p>
<div class="api-notes">
  <p class="api-ctype">G_LOG_DOMAIN</p>
</div>
<p class="api-heading">LOG_FATAL_MASK</p>
<p class="api-doc">GLib log levels that are considered fatal by default.

This is not used if structured logging is enabled; see
[Using Structured Logging][using-structured-logging].</p>
<div class="api-notes">
  <p class="api-ctype">G_LOG_FATAL_MASK</p>
</div>
<p class="api-heading">LOG_LEVEL_USER_SHIFT</p>
<p class="api-doc">Log levels below 1<<G_LOG_LEVEL_USER_SHIFT are used by GLib.
Higher bits can be used for user-defined log levels.</p>
<div class="api-notes">
  <p class="api-ctype">G_LOG_LEVEL_USER_SHIFT</p>
</div>
<p class="api-heading">MAJOR_VERSION</p>
<p class="api-doc">The major version number of the GLib library.

Like #glib_major_version, but from the headers used at
application compile time, rather than from the library
linked against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_MAJOR_VERSION</p>
</div>
<p class="api-heading">MAXINT16</p>
<p class="api-doc">The maximum value which can be held in a #gint16.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXINT16</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MAXINT32</p>
<p class="api-doc">The maximum value which can be held in a #gint32.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXINT32</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MAXINT64</p>
<p class="api-doc">The maximum value which can be held in a #gint64.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXINT64</p>
</div>
<p class="api-heading">MAXINT8</p>
<p class="api-doc">The maximum value which can be held in a #gint8.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXINT8</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MAXUINT16</p>
<p class="api-doc">The maximum value which can be held in a #guint16.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXUINT16</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MAXUINT32</p>
<p class="api-doc">The maximum value which can be held in a #guint32.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXUINT32</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MAXUINT64</p>
<p class="api-doc">The maximum value which can be held in a #guint64.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXUINT64</p>
</div>
<p class="api-heading">MAXUINT8</p>
<p class="api-doc">The maximum value which can be held in a #guint8.</p>
<div class="api-notes">
  <p class="api-ctype">G_MAXUINT8</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MICRO_VERSION</p>
<p class="api-doc">The micro version number of the GLib library.

Like #gtk_micro_version, but from the headers used at
application compile time, rather than from the library
linked against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_MICRO_VERSION</p>
</div>
<p class="api-heading">MININT16</p>
<p class="api-doc">The minimum value which can be held in a #gint16.</p>
<div class="api-notes">
  <p class="api-ctype">G_MININT16</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MININT32</p>
<p class="api-doc">The minimum value which can be held in a #gint32.</p>
<div class="api-notes">
  <p class="api-ctype">G_MININT32</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MININT64</p>
<p class="api-doc">The minimum value which can be held in a #gint64.</p>
<div class="api-notes">
  <p class="api-ctype">G_MININT64</p>
</div>
<p class="api-heading">MININT8</p>
<p class="api-doc">The minimum value which can be held in a #gint8.</p>
<div class="api-notes">
  <p class="api-ctype">G_MININT8</p>
  <p class="api-since">since 2.4</p>
</div>
<p class="api-heading">MINOR_VERSION</p>
<p class="api-doc">The minor version number of the GLib library.

Like #gtk_minor_version, but from the headers used at
application compile time, rather than from the library
linked against at application run time.</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_MINOR_VERSION</p>
</div>
<p class="api-heading">MODULE_SUFFIX</p>
<div class="api-notes">
  <p class="api-ctype">G_MODULE_SUFFIX</p>
</div>
<p class="api-heading">OPTION_REMAINING</p>
<p class="api-doc">If a long option in the main group has this name, it is not treated as a
regular option. Instead it collects all non-option arguments which would
otherwise be left in `argv`. The option must be of type
%G_OPTION_ARG_CALLBACK, %G_OPTION_ARG_STRING_ARRAY
or %G_OPTION_ARG_FILENAME_ARRAY.


Using #G_OPTION_REMAINING instead of simply scanning `argv`
for leftover arguments has the advantage that GOption takes care of
necessary encoding conversions for strings or filenames.</p>
<div class="api-notes">
  <p class="api-ctype">G_OPTION_REMAINING</p>
  <p class="api-since">since 2.6</p>
</div>
<p class="api-heading">PDP_ENDIAN</p>
<p class="api-doc">Specifies one of the possible types of byte order
(currently unused). See #G_BYTE_ORDER.</p>
<div class="api-notes">
  <p class="api-ctype">G_PDP_ENDIAN</p>
</div>
<p class="api-heading">PI</p>
<p class="api-doc">The value of pi (ratio of circle's circumference to its diameter).</p>
<div class="api-notes">
  <p class="api-ctype">G_PI</p>
</div>
<p class="api-heading">PID_FORMAT</p>
<p class="api-doc">A format specifier that can be used in printf()-style format strings
when printing a #GPid.</p>
<div class="api-notes">
  <p class="api-ctype">G_PID_FORMAT</p>
  <p class="api-since">since 2.50</p>
</div>
<p class="api-heading">PI_2</p>
<p class="api-doc">Pi divided by 2.</p>
<div class="api-notes">
  <p class="api-ctype">G_PI_2</p>
</div>
<p class="api-heading">PI_4</p>
<p class="api-doc">Pi divided by 4.</p>
<div class="api-notes">
  <p class="api-ctype">G_PI_4</p>
</div>
<p class="api-heading">POLLFD_FORMAT</p>
<p class="api-doc">A format specifier that can be used in printf()-style format strings
when printing the @fd member of a #GPollFD.</p>
<div class="api-notes">
  <p class="api-ctype">G_POLLFD_FORMAT</p>
</div>
<p class="api-heading">PRIORITY_DEFAULT</p>
<p class="api-doc">Use this for default priority event sources.

In GLib this priority is used when adding timeout functions
with g_timeout_add(). In GDK this priority is used for events
from the X server.</p>
<div class="api-notes">
  <p class="api-ctype">G_PRIORITY_DEFAULT</p>
</div>
<p class="api-heading">PRIORITY_DEFAULT_IDLE</p>
<p class="api-doc">Use this for default priority idle functions.

In GLib this priority is used when adding idle functions with
g_idle_add().</p>
<div class="api-notes">
  <p class="api-ctype">G_PRIORITY_DEFAULT_IDLE</p>
</div>
<p class="api-heading">PRIORITY_HIGH</p>
<p class="api-doc">Use this for high priority event sources.

It is not used within GLib or GTK+.</p>
<div class="api-notes">
  <p class="api-ctype">G_PRIORITY_HIGH</p>
</div>
<p class="api-heading">PRIORITY_HIGH_IDLE</p>
<p class="api-doc">Use this for high priority idle functions.

GTK+ uses #G_PRIORITY_HIGH_IDLE + 10 for resizing operations,
and #G_PRIORITY_HIGH_IDLE + 20 for redrawing operations. (This is
done to ensure that any pending resizes are processed before any
pending redraws, so that widgets are not redrawn twice unnecessarily.)</p>
<div class="api-notes">
  <p class="api-ctype">G_PRIORITY_HIGH_IDLE</p>
</div>
<p class="api-heading">PRIORITY_LOW</p>
<p class="api-doc">Use this for very low priority background tasks.

It is not used within GLib or GTK+.</p>
<div class="api-notes">
  <p class="api-ctype">G_PRIORITY_LOW</p>
</div>
<p class="api-heading">SEARCHPATH_SEPARATOR</p>
<p class="api-doc">The search path separator character.
This is ':' on UNIX machines and ';' under Windows.</p>
<div class="api-notes">
  <p class="api-ctype">G_SEARCHPATH_SEPARATOR</p>
</div>
<p class="api-heading">SEARCHPATH_SEPARATOR_S</p>
<p class="api-doc">The search path separator as a string.
This is ":" on UNIX machines and ";" under Windows.</p>
<div class="api-notes">
  <p class="api-ctype">G_SEARCHPATH_SEPARATOR_S</p>
</div>
<p class="api-heading">SIZEOF_LONG</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SIZEOF_LONG</p>
</div>
<p class="api-heading">SIZEOF_SIZE_T</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SIZEOF_SIZE_T</p>
</div>
<p class="api-heading">SIZEOF_SSIZE_T</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SIZEOF_SSIZE_T</p>
</div>
<p class="api-heading">SIZEOF_VOID_P</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SIZEOF_VOID_P</p>
</div>
<p class="api-heading">SOURCE_CONTINUE</p>
<p class="api-doc">Use this macro as the return value of a #GSourceFunc to leave
the #GSource in the main loop.</p>
<div class="api-notes">
  <p class="api-ctype">G_SOURCE_CONTINUE</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">SOURCE_REMOVE</p>
<p class="api-doc">Use this macro as the return value of a #GSourceFunc to remove
the #GSource from the main loop.</p>
<div class="api-notes">
  <p class="api-ctype">G_SOURCE_REMOVE</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">SQRT2</p>
<p class="api-doc">The square root of two.</p>
<div class="api-notes">
  <p class="api-ctype">G_SQRT2</p>
</div>
<p class="api-heading">STR_DELIMITERS</p>
<p class="api-doc">The standard delimiters, used in g_strdelimit().</p>
<div class="api-notes">
  <p class="api-ctype">G_STR_DELIMITERS</p>
</div>
<p class="api-heading">SYSDEF_AF_INET</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SYSDEF_AF_INET</p>
</div>
<p class="api-heading">SYSDEF_AF_INET6</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SYSDEF_AF_INET6</p>
</div>
<p class="api-heading">SYSDEF_AF_UNIX</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SYSDEF_AF_UNIX</p>
</div>
<p class="api-heading">SYSDEF_MSG_DONTROUTE</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SYSDEF_MSG_DONTROUTE</p>
</div>
<p class="api-heading">SYSDEF_MSG_OOB</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SYSDEF_MSG_OOB</p>
</div>
<p class="api-heading">SYSDEF_MSG_PEEK</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_SYSDEF_MSG_PEEK</p>
</div>
<p class="api-heading">TIME_SPAN_DAY</p>
<p class="api-doc">Evaluates to a time span of one day.</p>
<div class="api-notes">
  <p class="api-ctype">G_TIME_SPAN_DAY</p>
  <p class="api-since">since 2.26</p>
</div>
<p class="api-heading">TIME_SPAN_HOUR</p>
<p class="api-doc">Evaluates to a time span of one hour.</p>
<div class="api-notes">
  <p class="api-ctype">G_TIME_SPAN_HOUR</p>
  <p class="api-since">since 2.26</p>
</div>
<p class="api-heading">TIME_SPAN_MILLISECOND</p>
<p class="api-doc">Evaluates to a time span of one millisecond.</p>
<div class="api-notes">
  <p class="api-ctype">G_TIME_SPAN_MILLISECOND</p>
  <p class="api-since">since 2.26</p>
</div>
<p class="api-heading">TIME_SPAN_MINUTE</p>
<p class="api-doc">Evaluates to a time span of one minute.</p>
<div class="api-notes">
  <p class="api-ctype">G_TIME_SPAN_MINUTE</p>
  <p class="api-since">since 2.26</p>
</div>
<p class="api-heading">TIME_SPAN_SECOND</p>
<p class="api-doc">Evaluates to a time span of one second.</p>
<div class="api-notes">
  <p class="api-ctype">G_TIME_SPAN_SECOND</p>
  <p class="api-since">since 2.26</p>
</div>
<p class="api-heading">UNICHAR_MAX_DECOMPOSITION_LENGTH</p>
<p class="api-doc">The maximum length (in codepoints) of a compatibility or canonical
decomposition of a single Unicode character.

This is as defined by Unicode 6.1.</p>
<div class="api-notes">
  <p class="api-ctype">G_UNICHAR_MAX_DECOMPOSITION_LENGTH</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">URI_RESERVED_CHARS_GENERIC_DELIMITERS</p>
<p class="api-doc">Generic delimiters characters as defined in RFC 3986. Includes ":/?#[]@".</p>
<div class="api-notes">
  <p class="api-ctype">G_URI_RESERVED_CHARS_GENERIC_DELIMITERS</p>
</div>
<p class="api-heading">URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS</p>
<p class="api-doc">Subcomponent delimiter characters as defined in RFC 3986. Includes "!$&'()*+,;=".</p>
<div class="api-notes">
  <p class="api-ctype">G_URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS</p>
</div>
<p class="api-heading">USEC_PER_SEC</p>
<p class="api-doc">Number of microseconds in one second (1 million).
This macro is provided for code readability.</p>
<div class="api-notes">
  <p class="api-ctype">G_USEC_PER_SEC</p>
</div>
<p class="api-heading">VA_COPY_AS_ARRAY</p>
<div class="api-notes">
  <p class="api-ctype">G_VA_COPY_AS_ARRAY</p>
</div>
<p class="api-heading">VERSION_MIN_REQUIRED</p>
<p class="api-doc">A macro that should be defined by the user prior to including
the glib.h header.
The definition should be one of the predefined GLib version
macros: %GLIB_VERSION_2_26, %GLIB_VERSION_2_28,...

This macro defines the earliest version of GLib that the package is
required to be able to compile against.

If the compiler is configured to warn about the use of deprecated
functions, then using functions that were deprecated in version
%GLIB_VERSION_MIN_REQUIRED or earlier will cause warnings (but
using functions deprecated in later releases will not).</p>
<div class="api-notes">
  <p class="api-ctype">GLIB_VERSION_MIN_REQUIRED</p>
  <p class="api-since">since 2.32</p>
</div>
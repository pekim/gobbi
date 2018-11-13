// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

/*
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
*/
const GINT16_MODIFIER string = C.G_GINT16_MODIFIER

/*
The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint32 or #guint32. It
is a string literal. See also #G_GINT16_MODIFIER.
*/
const GINT32_MODIFIER string = C.G_GINT32_MODIFIER

/*
The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gint64 or #guint64.
It is a string literal.

Some platforms do not support printing 64-bit integers, even
though the types are supported. On such platforms %G_GINT64_MODIFIER
is not defined.
*/
const GINT64_MODIFIER string = C.G_GINT64_MODIFIER

// Unsupported : type gint16 for MAXINT16

// Unsupported : type gint32 for MAXINT32

// Unsupported : type gint8 for MAXINT8

// Unsupported : type guint16 for MAXUINT16

// Unsupported : type guint32 for MAXUINT32

// Unsupported : type guint8 for MAXUINT8

// Unsupported : type gint16 for MININT16

// Unsupported : type gint32 for MININT32

// Unsupported : type gint8 for MININT8

// This is a generated file - DO NOT EDIT
// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// This is the platform dependent conversion specifier for scanning
// and printing values of type #gsize. See also #G_GINT16_FORMAT.
const GSIZE_FORMAT string = C.G_GSIZE_FORMAT

// The platform dependent length modifier for conversion specifiers
// for scanning and printing values of type #gsize. It
// is a string literal.
const GSIZE_MODIFIER string = C.G_GSIZE_MODIFIER

// This is the platform dependent conversion specifier for scanning
// and printing values of type #gssize. See also #G_GINT16_FORMAT.
const GSSIZE_FORMAT string = C.G_GSSIZE_FORMAT

// The platform dependent length modifier for conversion specifiers
// for scanning and printing values of type #gssize. It
// is a string literal.
const GSSIZE_MODIFIER string = C.G_GSSIZE_MODIFIER

// If a long option in the main group has this name, it is not treated as a
// regular option. Instead it collects all non-option arguments which would
// otherwise be left in `argv`. The option must be of type
// %G_OPTION_ARG_CALLBACK, %G_OPTION_ARG_STRING_ARRAY
// or %G_OPTION_ARG_FILENAME_ARRAY.
//
//
// Using #G_OPTION_REMAINING instead of simply scanning `argv`
// for leftover arguments has the advantage that GOption takes care of
// necessary encoding conversions for strings or filenames.
const OPTION_REMAINING string = C.G_OPTION_REMAINING

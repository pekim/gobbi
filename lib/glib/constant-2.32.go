// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

/*
Use this macro as the return value of a #GSourceFunc to leave
the #GSource in the main loop.
*/
const SOURCE_CONTINUE bool = true // C.G_SOURCE_CONTINUE
/*
Use this macro as the return value of a #GSourceFunc to remove
the #GSource from the main loop.
*/
const SOURCE_REMOVE bool = false // C.G_SOURCE_REMOVE
/*
The maximum length (in codepoints) of a compatibility or canonical
decomposition of a single Unicode character.

This is as defined by Unicode 6.1.
*/
const UNICHAR_MAX_DECOMPOSITION_LENGTH int = C.G_UNICHAR_MAX_DECOMPOSITION_LENGTH

/*
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
*/
const VERSION_MIN_REQUIRED int = C.GLIB_VERSION_MIN_REQUIRED

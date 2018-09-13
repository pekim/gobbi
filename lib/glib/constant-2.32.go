// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : type gboolean for SOURCE_CONTINUE

// Unsupported : type gboolean for SOURCE_REMOVE

const UNICHAR_MAX_DECOMPOSITION_LENGTH int = C.G_UNICHAR_MAX_DECOMPOSITION_LENGTH
const VERSION_MIN_REQUIRED int = C.GLIB_VERSION_MIN_REQUIRED

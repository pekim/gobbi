// This is a generated file - DO NOT EDIT

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

func boolToGboolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

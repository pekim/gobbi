// This is a generated file - DO NOT EDIT

package gdkpixbuf

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

func boolToGboolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

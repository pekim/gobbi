// This is a generated file - DO NOT EDIT

package cairo

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib/gstdio.h>
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

func boolToGboolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

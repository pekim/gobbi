// This is a generated file - DO NOT EDIT

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

func boolToGboolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
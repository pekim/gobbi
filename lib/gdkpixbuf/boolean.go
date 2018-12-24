// This is a generated file - DO NOT EDIT

package gdkpixbuf

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

func boolToGboolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

// Code generated - DO NOT EDIT.

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <../gdk/gdk_event.h>
// #include <stdlib.h>
import "C"

func boolToGboolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

func boolToGboolean(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}

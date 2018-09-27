// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetLocaleDirection is a wrapper around the C function gtk_get_locale_direction.
func GetLocaleDirection() TextDirection {
	retC := C.gtk_get_locale_direction()
	retGo := (TextDirection)(retC)

	return retGo
}

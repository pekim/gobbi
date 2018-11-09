// This is a generated file - DO NOT EDIT
// +build gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetRevealed is a wrapper around the C function gtk_info_bar_get_revealed.
func (recv *InfoBar) GetRevealed() bool {
	retC := C.gtk_info_bar_get_revealed((*C.GtkInfoBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetRevealed is a wrapper around the C function gtk_info_bar_set_revealed.
func (recv *InfoBar) SetRevealed(revealed bool) {
	c_revealed :=
		boolToGboolean(revealed)

	C.gtk_info_bar_set_revealed((*C.GtkInfoBar)(recv.native), c_revealed)

	return
}

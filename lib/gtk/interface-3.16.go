// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Returns the size of a non-scrolling border around the
// outside of the scrollable. An example for this would
// be treeview headers. GTK+ can use this information to
// display overlayed graphics, like the overshoot indication,
// at the right position.
/*

C function : gtk_scrollable_get_border
*/
func (recv *Scrollable) GetBorder() (bool, *Border) {
	var c_border C.GtkBorder

	retC := C.gtk_scrollable_get_border((*C.GtkScrollable)(recv.native), &c_border)
	retGo := retC == C.TRUE

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return retGo, border
}

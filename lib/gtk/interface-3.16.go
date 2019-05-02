// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	"unsafe"
)

// GetBorder is a wrapper around the C function gtk_scrollable_get_border.
func (recv *Scrollable) GetBorder() (bool, *Border) {
	var c_border C.GtkBorder

	retC := C.gtk_scrollable_get_border((*C.GtkScrollable)(recv.native), &c_border)
	retGo := retC == C.TRUE

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return retGo, border
}

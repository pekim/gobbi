// This is a generated file - DO NOT EDIT
// +build gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetChildAtPos is a wrapper around the C function gtk_flow_box_get_child_at_pos.
func (recv *FlowBox) GetChildAtPos(x int32, y int32) *FlowBoxChild {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_flow_box_get_child_at_pos((*C.GtkFlowBox)(recv.native), c_x, c_y)
	var retGo (*FlowBoxChild)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FlowBoxChildNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

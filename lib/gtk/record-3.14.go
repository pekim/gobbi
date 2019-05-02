// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "C"

// IterGetState is a wrapper around the C function gtk_widget_path_iter_get_state.
func (recv *WidgetPath) IterGetState(pos int32) StateFlags {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_state((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := (StateFlags)(retC)

	return retGo
}

// IterSetState is a wrapper around the C function gtk_widget_path_iter_set_state.
func (recv *WidgetPath) IterSetState(pos int32, state StateFlags) {
	c_pos := (C.gint)(pos)

	c_state := (C.GtkStateFlags)(state)

	C.gtk_widget_path_iter_set_state((*C.GtkWidgetPath)(recv.native), c_pos, c_state)

	return
}

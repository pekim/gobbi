// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Returns the state flags corresponding to the widget found at
// the position @pos in the widget hierarchy defined by
// @path
/*

C function

gtk_widget_path_iter_get_state
*/
func (recv *WidgetPath) IterGetState(pos int32) StateFlags {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_state((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := (StateFlags)(retC)

	return retGo
}

// Sets the widget name for the widget found at position @pos
// in the widget hierarchy defined by @path.
//
// If you want to update just a single state flag, you need to do
// this manually, as this function updates all state flags.
//
// Setting a flag
//
// |[<!-- language="C" -->
// gtk_widget_path_iter_set_state (path, pos, gtk_widget_path_iter_get_state (path, pos) | flag);
// ]|
//
// Unsetting a flag
//
// |[<!-- language="C" -->
// gtk_widget_path_iter_set_state (path, pos, gtk_widget_path_iter_get_state (path, pos) & ~flag);
// ]|
/*

C function

gtk_widget_path_iter_set_state
*/
func (recv *WidgetPath) IterSetState(pos int32, state StateFlags) {
	c_pos := (C.gint)(pos)

	c_state := (C.GtkStateFlags)(state)

	C.gtk_widget_path_iter_set_state((*C.GtkWidgetPath)(recv.native), c_pos, c_state)

	return
}

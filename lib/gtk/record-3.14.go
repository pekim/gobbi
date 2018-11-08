// This is a generated file - DO NOT EDIT
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_target_list_new : unsupported parameter targets : no type generator for TargetEntry (GtkTargetEntry) for array param targets

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no type generator for gint (gint) for array param indices

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

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

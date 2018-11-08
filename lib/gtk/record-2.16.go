// This is a generated file - DO NOT EDIT
// +build gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// ActivatableIface is a wrapper around the C record GtkActivatableIface.
type ActivatableIface struct {
	native *C.GtkActivatableIface
	// Private : g_iface
	// no type for update
	// no type for sync_action_properties
}

func ActivatableIfaceNewFromC(u unsafe.Pointer) *ActivatableIface {
	c := (*C.GtkActivatableIface)(u)
	if c == nil {
		return nil
	}

	g := &ActivatableIface{native: c}

	return g
}

func (recv *ActivatableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_selection_data_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_new : unsupported parameter targets :

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

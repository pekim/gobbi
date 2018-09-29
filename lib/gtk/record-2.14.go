// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// BorderNew is a wrapper around the C function gtk_border_new.
func BorderNew() *Border {
	retC := C.gtk_border_new()
	retGo := BorderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_selection_data_get_data : no return type

// Unsupported : gtk_selection_data_get_data_type : return type : Blacklisted record : GdkAtom

// GetDisplay is a wrapper around the C function gtk_selection_data_get_display.
func (recv *SelectionData) GetDisplay() *gdk.Display {
	retC := C.gtk_selection_data_get_display((*C.GtkSelectionData)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFormat is a wrapper around the C function gtk_selection_data_get_format.
func (recv *SelectionData) GetFormat() int32 {
	retC := C.gtk_selection_data_get_format((*C.GtkSelectionData)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLength is a wrapper around the C function gtk_selection_data_get_length.
func (recv *SelectionData) GetLength() int32 {
	retC := C.gtk_selection_data_get_length((*C.GtkSelectionData)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_selection_data_get_target : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_target_list_new : unsupported parameter targets : no param type

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no param type

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

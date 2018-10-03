// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_paper_size_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// SymbolicColorNewWin32 is a wrapper around the C function gtk_symbolic_color_new_win32.
func SymbolicColorNewWin32(themeClass string, id int32) *SymbolicColor {
	c_theme_class := C.CString(themeClass)
	defer C.free(unsafe.Pointer(c_theme_class))

	c_id := (C.gint)(id)

	retC := C.gtk_symbolic_color_new_win32(c_theme_class, c_id)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_target_list_new : unsupported parameter targets : no param type

// Unsupported : gtk_tree_path_new_from_indices : unsupported parameter ... : varargs

// Unsupported : gtk_tree_path_new_from_indicesv : unsupported parameter indices : no param type

// Unsupported : gtk_tree_row_reference_new : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_row_reference_new_proxy : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*
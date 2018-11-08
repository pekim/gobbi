// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetCurrentName is a wrapper around the C function gtk_file_chooser_get_current_name.
func (recv *FileChooser) GetCurrentName() string {
	retC := C.gtk_file_chooser_get_current_name((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_model_rows_reordered_with_length : unsupported parameter new_order : no type generator for gint (gint) for array param new_order

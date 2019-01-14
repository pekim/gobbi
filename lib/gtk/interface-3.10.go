// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// RowsReorderedWithLength is a wrapper around the C function gtk_tree_model_rows_reordered_with_length.
func (recv *TreeModel) RowsReorderedWithLength(path *TreePath, iter *TreeIter, newOrder []int32) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_new_order_array := make([]C.gint, len(newOrder), len(newOrder))
	c_new_order_arrayPtr := &c_new_order_array[0]
	c_new_order := (*C.gint)(unsafe.Pointer(c_new_order_arrayPtr))

	c_length := (C.gint)(len(newOrder))

	C.gtk_tree_model_rows_reordered_with_length((*C.GtkTreeModel)(recv.native), c_path, c_iter, c_new_order, c_length)

	return
}

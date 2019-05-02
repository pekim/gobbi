// This is a generated file - DO NOT EDIT
// +build gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// ComboBoxNewWithEntry is a wrapper around the C function gtk_combo_box_new_with_entry.
func ComboBoxNewWithEntry() *ComboBox {
	retC := C.gtk_combo_box_new_with_entry()
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxNewWithModelAndEntry is a wrapper around the C function gtk_combo_box_new_with_model_and_entry.
func ComboBoxNewWithModelAndEntry(model *TreeModel) *ComboBox {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_combo_box_new_with_model_and_entry(c_model)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxTextNew is a wrapper around the C function gtk_combo_box_text_new.
func ComboBoxTextNew() *ComboBoxText {
	retC := C.gtk_combo_box_text_new()
	retGo := ComboBoxTextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxTextNewWithEntry is a wrapper around the C function gtk_combo_box_text_new_with_entry.
func ComboBoxTextNewWithEntry() *ComboBoxText {
	retC := C.gtk_combo_box_text_new_with_entry()
	retGo := ComboBoxTextNewFromC(unsafe.Pointer(retC))

	return retGo
}

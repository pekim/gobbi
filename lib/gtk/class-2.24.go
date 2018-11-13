// This is a generated file - DO NOT EDIT
// +build gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported signal 'activate-link' for AboutDialog : unsupported parameter uri : type utf8 :

// Creates a new empty #GtkComboBox with an entry.
/*

C function : gtk_combo_box_new_with_entry
*/
func ComboBoxNewWithEntry() *ComboBox {
	retC := C.gtk_combo_box_new_with_entry()
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new empty #GtkComboBox with an entry
// and with the model initialized to @model.
/*

C function : gtk_combo_box_new_with_model_and_entry
*/
func ComboBoxNewWithModelAndEntry(model *TreeModel) *ComboBox {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_combo_box_new_with_model_and_entry(c_model)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the column which @combo_box is using to get the strings
// from to display in the internal entry.
/*

C function : gtk_combo_box_get_entry_text_column
*/
func (recv *ComboBox) GetEntryTextColumn() int32 {
	retC := C.gtk_combo_box_get_entry_text_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns whether the combo box has an entry.
/*

C function : gtk_combo_box_get_has_entry
*/
func (recv *ComboBox) GetHasEntry() bool {
	retC := C.gtk_combo_box_get_has_entry((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the model column which @combo_box should use to get strings from
// to be @text_column. The column @text_column in the model of @combo_box
// must be of type %G_TYPE_STRING.
//
// This is only relevant if @combo_box has been created with
// #GtkComboBox:has-entry as %TRUE.
/*

C function : gtk_combo_box_set_entry_text_column
*/
func (recv *ComboBox) SetEntryTextColumn(textColumn int32) {
	c_text_column := (C.gint)(textColumn)

	C.gtk_combo_box_set_entry_text_column((*C.GtkComboBox)(recv.native), c_text_column)

	return
}

// Creates a new #GtkComboBoxText, which is a #GtkComboBox just displaying
// strings.
/*

C function : gtk_combo_box_text_new
*/
func ComboBoxTextNew() *ComboBoxText {
	retC := C.gtk_combo_box_text_new()
	retGo := ComboBoxTextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkComboBoxText, which is a #GtkComboBox just displaying
// strings. The combo box created by this function has an entry.
/*

C function : gtk_combo_box_text_new_with_entry
*/
func ComboBoxTextNewWithEntry() *ComboBoxText {
	retC := C.gtk_combo_box_text_new_with_entry()
	retGo := ComboBoxTextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends @text to the list of strings stored in @combo_box.
// If @id is non-%NULL then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a
// position of -1.
/*

C function : gtk_combo_box_text_append
*/
func (recv *ComboBoxText) Append(id string, text string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_append((*C.GtkComboBoxText)(recv.native), c_id, c_text)

	return
}

// Appends @text to the list of strings stored in @combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a
// position of -1.
/*

C function : gtk_combo_box_text_append_text
*/
func (recv *ComboBoxText) AppendText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_append_text((*C.GtkComboBoxText)(recv.native), c_text)

	return
}

// Returns the currently active string in @combo_box, or %NULL
// if none is selected. If @combo_box contains an entry, this
// function will return its contents (which will not necessarily
// be an item from the list).
/*

C function : gtk_combo_box_text_get_active_text
*/
func (recv *ComboBoxText) GetActiveText() string {
	retC := C.gtk_combo_box_text_get_active_text((*C.GtkComboBoxText)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Inserts @text at @position in the list of strings stored in @combo_box.
//
// If @position is negative then @text is appended.
//
// This is the same as calling gtk_combo_box_text_insert() with a %NULL
// ID string.
/*

C function : gtk_combo_box_text_insert_text
*/
func (recv *ComboBoxText) InsertText(position int32, text string) {
	c_position := (C.gint)(position)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_insert_text((*C.GtkComboBoxText)(recv.native), c_position, c_text)

	return
}

// Prepends @text to the list of strings stored in @combo_box.
// If @id is non-%NULL then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a
// position of 0.
/*

C function : gtk_combo_box_text_prepend
*/
func (recv *ComboBoxText) Prepend(id string, text string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_prepend((*C.GtkComboBoxText)(recv.native), c_id, c_text)

	return
}

// Prepends @text to the list of strings stored in @combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a
// position of 0.
/*

C function : gtk_combo_box_text_prepend_text
*/
func (recv *ComboBoxText) PrependText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_prepend_text((*C.GtkComboBoxText)(recv.native), c_text)

	return
}

// Removes the string at @position from @combo_box.
/*

C function : gtk_combo_box_text_remove
*/
func (recv *ComboBoxText) Remove(position int32) {
	c_position := (C.gint)(position)

	C.gtk_combo_box_text_remove((*C.GtkComboBoxText)(recv.native), c_position)

	return
}

// Gets the current group name for @notebook.
/*

C function : gtk_notebook_get_group_name
*/
func (recv *Notebook) GetGroupName() string {
	retC := C.gtk_notebook_get_group_name((*C.GtkNotebook)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets a group name for @notebook.
//
// Notebooks with the same name will be able to exchange tabs
// via drag and drop. A notebook with a %NULL group name will
// not be able to exchange tabs with any other notebook.
/*

C function : gtk_notebook_set_group_name
*/
func (recv *Notebook) SetGroupName(groupName string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_notebook_set_group_name((*C.GtkNotebook)(recv.native), c_group_name)

	return
}

// Gets the number of digits to round the value to when
// it changes. See #GtkRange::change-value.
/*

C function : gtk_range_get_round_digits
*/
func (recv *Range) GetRoundDigits() int32 {
	retC := C.gtk_range_get_round_digits((*C.GtkRange)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the number of digits to round the value to when
// it changes. See #GtkRange::change-value.
/*

C function : gtk_range_set_round_digits
*/
func (recv *Range) SetRoundDigits(roundDigits int32) {
	c_round_digits := (C.gint)(roundDigits)

	C.gtk_range_set_round_digits((*C.GtkRange)(recv.native), c_round_digits)

	return
}

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

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// ComboBoxNewWithEntry is a wrapper around the C function gtk_combo_box_new_with_entry.
func ComboBoxNewWithEntry() *ComboBox {
	retC := C.gtk_combo_box_new_with_entry()
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetEntryTextColumn is a wrapper around the C function gtk_combo_box_get_entry_text_column.
func (recv *ComboBox) GetEntryTextColumn() int32 {
	retC := C.gtk_combo_box_get_entry_text_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHasEntry is a wrapper around the C function gtk_combo_box_get_has_entry.
func (recv *ComboBox) GetHasEntry() bool {
	retC := C.gtk_combo_box_get_has_entry((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetEntryTextColumn is a wrapper around the C function gtk_combo_box_set_entry_text_column.
func (recv *ComboBox) SetEntryTextColumn(textColumn int32) {
	c_text_column := (C.gint)(textColumn)

	C.gtk_combo_box_set_entry_text_column((*C.GtkComboBox)(recv.native), c_text_column)

	return
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

// Append is a wrapper around the C function gtk_combo_box_text_append.
func (recv *ComboBoxText) Append(id string, text string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_append((*C.GtkComboBoxText)(recv.native), c_id, c_text)

	return
}

// AppendText is a wrapper around the C function gtk_combo_box_text_append_text.
func (recv *ComboBoxText) AppendText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_append_text((*C.GtkComboBoxText)(recv.native), c_text)

	return
}

// GetActiveText is a wrapper around the C function gtk_combo_box_text_get_active_text.
func (recv *ComboBoxText) GetActiveText() string {
	retC := C.gtk_combo_box_text_get_active_text((*C.GtkComboBoxText)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// InsertText is a wrapper around the C function gtk_combo_box_text_insert_text.
func (recv *ComboBoxText) InsertText(position int32, text string) {
	c_position := (C.gint)(position)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_insert_text((*C.GtkComboBoxText)(recv.native), c_position, c_text)

	return
}

// Prepend is a wrapper around the C function gtk_combo_box_text_prepend.
func (recv *ComboBoxText) Prepend(id string, text string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_prepend((*C.GtkComboBoxText)(recv.native), c_id, c_text)

	return
}

// PrependText is a wrapper around the C function gtk_combo_box_text_prepend_text.
func (recv *ComboBoxText) PrependText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_prepend_text((*C.GtkComboBoxText)(recv.native), c_text)

	return
}

// Remove is a wrapper around the C function gtk_combo_box_text_remove.
func (recv *ComboBoxText) Remove(position int32) {
	c_position := (C.gint)(position)

	C.gtk_combo_box_text_remove((*C.GtkComboBoxText)(recv.native), c_position)

	return
}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// GetGroupName is a wrapper around the C function gtk_notebook_get_group_name.
func (recv *Notebook) GetGroupName() string {
	retC := C.gtk_notebook_get_group_name((*C.GtkNotebook)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetGroupName is a wrapper around the C function gtk_notebook_set_group_name.
func (recv *Notebook) SetGroupName(groupName string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_notebook_set_group_name((*C.GtkNotebook)(recv.native), c_group_name)

	return
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// GetRoundDigits is a wrapper around the C function gtk_range_get_round_digits.
func (recv *Range) GetRoundDigits() int32 {
	retC := C.gtk_range_get_round_digits((*C.GtkRange)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetRoundDigits is a wrapper around the C function gtk_range_set_round_digits.
func (recv *Range) SetRoundDigits(roundDigits int32) {
	c_round_digits := (C.gint)(roundDigits)

	C.gtk_range_set_round_digits((*C.GtkRange)(recv.native), c_round_digits)

	return
}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

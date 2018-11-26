// This is a generated file - DO NOT EDIT
// +build gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	gboolean aboutdialog_activateLinkHandler(GObject *, gchar*, gpointer);

	static gulong AboutDialog_signal_connect_activate_link(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-link", G_CALLBACK(aboutdialog_activateLinkHandler), data);
	}

*/
import "C"

type signalAboutDialogActivateLinkDetail struct {
	callback  AboutDialogSignalActivateLinkCallback
	handlerID C.gulong
}

var signalAboutDialogActivateLinkId int
var signalAboutDialogActivateLinkMap = make(map[int]signalAboutDialogActivateLinkDetail)
var signalAboutDialogActivateLinkLock sync.Mutex

// AboutDialogSignalActivateLinkCallback is a callback function for a 'activate-link' signal emitted from a AboutDialog.
type AboutDialogSignalActivateLinkCallback func(uri string) bool

/*
ConnectActivateLink connects the callback to the 'activate-link' signal for the AboutDialog.

The returned value represents the connection, and may be passed to DisconnectActivateLink to remove it.
*/
func (recv *AboutDialog) ConnectActivateLink(callback AboutDialogSignalActivateLinkCallback) int {
	signalAboutDialogActivateLinkLock.Lock()
	defer signalAboutDialogActivateLinkLock.Unlock()

	signalAboutDialogActivateLinkId++
	instance := C.gpointer(recv.native)
	handlerID := C.AboutDialog_signal_connect_activate_link(instance, C.gpointer(uintptr(signalAboutDialogActivateLinkId)))

	detail := signalAboutDialogActivateLinkDetail{callback, handlerID}
	signalAboutDialogActivateLinkMap[signalAboutDialogActivateLinkId] = detail

	return signalAboutDialogActivateLinkId
}

/*
DisconnectActivateLink disconnects a callback from the 'activate-link' signal for the AboutDialog.

The connectionID should be a value returned from a call to ConnectActivateLink.
*/
func (recv *AboutDialog) DisconnectActivateLink(connectionID int) {
	signalAboutDialogActivateLinkLock.Lock()
	defer signalAboutDialogActivateLinkLock.Unlock()

	detail, exists := signalAboutDialogActivateLinkMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAboutDialogActivateLinkMap, connectionID)
}

//export aboutdialog_activateLinkHandler
func aboutdialog_activateLinkHandler(_ *C.GObject, c_uri C.gchar, data C.gpointer) C.gboolean {

	index := int(uintptr(data))
	callback := signalAboutDialogActivateLinkMap[index].callback
	retGo := callback(uri)
	retC :=
		boolToGboolean(retGo)
	return retC
}

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

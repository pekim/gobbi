// This is a generated file - DO NOT EDIT
// +build gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
var signalAboutDialogActivateLinkLock sync.RWMutex

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
func aboutdialog_activateLinkHandler(_ *C.GObject, c_uri *C.gchar, data C.gpointer) C.gboolean {
	signalAboutDialogActivateLinkLock.RLock()
	defer signalAboutDialogActivateLinkLock.RUnlock()

	uri := C.GoString(c_uri)

	index := int(uintptr(data))
	callback := signalAboutDialogActivateLinkMap[index].callback
	retGo := callback(uri)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_combo_box_new_with_entry

// Blacklisted : gtk_combo_box_new_with_model_and_entry

// Blacklisted : gtk_combo_box_get_entry_text_column

// Blacklisted : gtk_combo_box_get_has_entry

// Blacklisted : gtk_combo_box_set_entry_text_column

// Blacklisted : gtk_combo_box_text_new

// Blacklisted : gtk_combo_box_text_new_with_entry

// Blacklisted : gtk_combo_box_text_append

// Blacklisted : gtk_combo_box_text_append_text

// Blacklisted : gtk_combo_box_text_get_active_text

// Blacklisted : gtk_combo_box_text_insert_text

// Blacklisted : gtk_combo_box_text_prepend

// Blacklisted : gtk_combo_box_text_prepend_text

// Blacklisted : gtk_combo_box_text_remove

// Blacklisted : gtk_notebook_get_group_name

// Blacklisted : gtk_notebook_set_group_name

// Blacklisted : gtk_range_get_round_digits

// Blacklisted : gtk_range_set_round_digits

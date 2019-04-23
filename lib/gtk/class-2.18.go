// This is a generated file - DO NOT EDIT
// +build gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void entrybuffer_deletedTextHandler(GObject *, guint, guint, gpointer);

	static gulong EntryBuffer_signal_connect_deleted_text(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "deleted-text", G_CALLBACK(entrybuffer_deletedTextHandler), data);
	}

*/
/*

	void entrybuffer_insertedTextHandler(GObject *, guint, gchar*, guint, gpointer);

	static gulong EntryBuffer_signal_connect_inserted_text(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "inserted-text", G_CALLBACK(entrybuffer_insertedTextHandler), data);
	}

*/
/*

	void infobar_closeHandler(GObject *, gpointer);

	static gulong InfoBar_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", G_CALLBACK(infobar_closeHandler), data);
	}

*/
/*

	void infobar_responseHandler(GObject *, gint, gpointer);

	static gulong InfoBar_signal_connect_response(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "response", G_CALLBACK(infobar_responseHandler), data);
	}

*/
/*

	void label_activateCurrentLinkHandler(GObject *, gpointer);

	static gulong Label_signal_connect_activate_current_link(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-current-link", G_CALLBACK(label_activateCurrentLinkHandler), data);
	}

*/
/*

	gboolean label_activateLinkHandler(GObject *, gchar*, gpointer);

	static gulong Label_signal_connect_activate_link(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-link", G_CALLBACK(label_activateLinkHandler), data);
	}

*/
/*

	void printoperation_updateCustomWidgetHandler(GObject *, GtkWidget *, GtkPageSetup *, GtkPrintSettings *, gpointer);

	static gulong PrintOperation_signal_connect_update_custom_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "update-custom-widget", G_CALLBACK(printoperation_updateCustomWidgetHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_cell_renderer_get_alignment

// Blacklisted : gtk_cell_renderer_get_padding

// Blacklisted : gtk_cell_renderer_get_sensitive

// Blacklisted : gtk_cell_renderer_get_visible

// Blacklisted : gtk_cell_renderer_set_alignment

// Blacklisted : gtk_cell_renderer_set_padding

// Blacklisted : gtk_cell_renderer_set_sensitive

// Blacklisted : gtk_cell_renderer_set_visible

// Blacklisted : gtk_cell_renderer_toggle_get_activatable

// Blacklisted : gtk_cell_renderer_toggle_set_activatable

// Blacklisted : gtk_entry_new_with_buffer

// Blacklisted : gtk_entry_get_buffer

// Blacklisted : gtk_entry_set_buffer

type signalEntryBufferDeletedTextDetail struct {
	callback  EntryBufferSignalDeletedTextCallback
	handlerID C.gulong
}

var signalEntryBufferDeletedTextId int
var signalEntryBufferDeletedTextMap = make(map[int]signalEntryBufferDeletedTextDetail)
var signalEntryBufferDeletedTextLock sync.RWMutex

// EntryBufferSignalDeletedTextCallback is a callback function for a 'deleted-text' signal emitted from a EntryBuffer.
type EntryBufferSignalDeletedTextCallback func(position uint32, nChars uint32)

/*
ConnectDeletedText connects the callback to the 'deleted-text' signal for the EntryBuffer.

The returned value represents the connection, and may be passed to DisconnectDeletedText to remove it.
*/
func (recv *EntryBuffer) ConnectDeletedText(callback EntryBufferSignalDeletedTextCallback) int {
	signalEntryBufferDeletedTextLock.Lock()
	defer signalEntryBufferDeletedTextLock.Unlock()

	signalEntryBufferDeletedTextId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryBuffer_signal_connect_deleted_text(instance, C.gpointer(uintptr(signalEntryBufferDeletedTextId)))

	detail := signalEntryBufferDeletedTextDetail{callback, handlerID}
	signalEntryBufferDeletedTextMap[signalEntryBufferDeletedTextId] = detail

	return signalEntryBufferDeletedTextId
}

/*
DisconnectDeletedText disconnects a callback from the 'deleted-text' signal for the EntryBuffer.

The connectionID should be a value returned from a call to ConnectDeletedText.
*/
func (recv *EntryBuffer) DisconnectDeletedText(connectionID int) {
	signalEntryBufferDeletedTextLock.Lock()
	defer signalEntryBufferDeletedTextLock.Unlock()

	detail, exists := signalEntryBufferDeletedTextMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryBufferDeletedTextMap, connectionID)
}

//export entrybuffer_deletedTextHandler
func entrybuffer_deletedTextHandler(_ *C.GObject, c_position C.guint, c_n_chars C.guint, data C.gpointer) {
	signalEntryBufferDeletedTextLock.RLock()
	defer signalEntryBufferDeletedTextLock.RUnlock()

	position := uint32(c_position)

	nChars := uint32(c_n_chars)

	index := int(uintptr(data))
	callback := signalEntryBufferDeletedTextMap[index].callback
	callback(position, nChars)
}

type signalEntryBufferInsertedTextDetail struct {
	callback  EntryBufferSignalInsertedTextCallback
	handlerID C.gulong
}

var signalEntryBufferInsertedTextId int
var signalEntryBufferInsertedTextMap = make(map[int]signalEntryBufferInsertedTextDetail)
var signalEntryBufferInsertedTextLock sync.RWMutex

// EntryBufferSignalInsertedTextCallback is a callback function for a 'inserted-text' signal emitted from a EntryBuffer.
type EntryBufferSignalInsertedTextCallback func(position uint32, chars string, nChars uint32)

/*
ConnectInsertedText connects the callback to the 'inserted-text' signal for the EntryBuffer.

The returned value represents the connection, and may be passed to DisconnectInsertedText to remove it.
*/
func (recv *EntryBuffer) ConnectInsertedText(callback EntryBufferSignalInsertedTextCallback) int {
	signalEntryBufferInsertedTextLock.Lock()
	defer signalEntryBufferInsertedTextLock.Unlock()

	signalEntryBufferInsertedTextId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryBuffer_signal_connect_inserted_text(instance, C.gpointer(uintptr(signalEntryBufferInsertedTextId)))

	detail := signalEntryBufferInsertedTextDetail{callback, handlerID}
	signalEntryBufferInsertedTextMap[signalEntryBufferInsertedTextId] = detail

	return signalEntryBufferInsertedTextId
}

/*
DisconnectInsertedText disconnects a callback from the 'inserted-text' signal for the EntryBuffer.

The connectionID should be a value returned from a call to ConnectInsertedText.
*/
func (recv *EntryBuffer) DisconnectInsertedText(connectionID int) {
	signalEntryBufferInsertedTextLock.Lock()
	defer signalEntryBufferInsertedTextLock.Unlock()

	detail, exists := signalEntryBufferInsertedTextMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryBufferInsertedTextMap, connectionID)
}

//export entrybuffer_insertedTextHandler
func entrybuffer_insertedTextHandler(_ *C.GObject, c_position C.guint, c_chars *C.gchar, c_n_chars C.guint, data C.gpointer) {
	signalEntryBufferInsertedTextLock.RLock()
	defer signalEntryBufferInsertedTextLock.RUnlock()

	position := uint32(c_position)

	chars := C.GoString(c_chars)

	nChars := uint32(c_n_chars)

	index := int(uintptr(data))
	callback := signalEntryBufferInsertedTextMap[index].callback
	callback(position, chars, nChars)
}

// Blacklisted : gtk_entry_buffer_new

// Blacklisted : gtk_entry_buffer_delete_text

// Blacklisted : gtk_entry_buffer_emit_deleted_text

// Blacklisted : gtk_entry_buffer_emit_inserted_text

// Blacklisted : gtk_entry_buffer_get_bytes

// Blacklisted : gtk_entry_buffer_get_length

// Blacklisted : gtk_entry_buffer_get_max_length

// GetText is a wrapper around the C function gtk_entry_buffer_get_text.
func (recv *EntryBuffer) GetText() string {
	retC := C.gtk_entry_buffer_get_text((*C.GtkEntryBuffer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : gtk_entry_buffer_insert_text

// Blacklisted : gtk_entry_buffer_set_max_length

// SetText is a wrapper around the C function gtk_entry_buffer_set_text.
func (recv *EntryBuffer) SetText(chars string, nChars int32) {
	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_buffer_set_text((*C.GtkEntryBuffer)(recv.native), c_chars, c_n_chars)

	return
}

// Blacklisted : gtk_icon_view_get_item_padding

// Blacklisted : gtk_icon_view_set_item_padding

type signalInfoBarCloseDetail struct {
	callback  InfoBarSignalCloseCallback
	handlerID C.gulong
}

var signalInfoBarCloseId int
var signalInfoBarCloseMap = make(map[int]signalInfoBarCloseDetail)
var signalInfoBarCloseLock sync.RWMutex

// InfoBarSignalCloseCallback is a callback function for a 'close' signal emitted from a InfoBar.
type InfoBarSignalCloseCallback func()

/*
ConnectClose connects the callback to the 'close' signal for the InfoBar.

The returned value represents the connection, and may be passed to DisconnectClose to remove it.
*/
func (recv *InfoBar) ConnectClose(callback InfoBarSignalCloseCallback) int {
	signalInfoBarCloseLock.Lock()
	defer signalInfoBarCloseLock.Unlock()

	signalInfoBarCloseId++
	instance := C.gpointer(recv.native)
	handlerID := C.InfoBar_signal_connect_close(instance, C.gpointer(uintptr(signalInfoBarCloseId)))

	detail := signalInfoBarCloseDetail{callback, handlerID}
	signalInfoBarCloseMap[signalInfoBarCloseId] = detail

	return signalInfoBarCloseId
}

/*
DisconnectClose disconnects a callback from the 'close' signal for the InfoBar.

The connectionID should be a value returned from a call to ConnectClose.
*/
func (recv *InfoBar) DisconnectClose(connectionID int) {
	signalInfoBarCloseLock.Lock()
	defer signalInfoBarCloseLock.Unlock()

	detail, exists := signalInfoBarCloseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalInfoBarCloseMap, connectionID)
}

//export infobar_closeHandler
func infobar_closeHandler(_ *C.GObject, data C.gpointer) {
	signalInfoBarCloseLock.RLock()
	defer signalInfoBarCloseLock.RUnlock()

	index := int(uintptr(data))
	callback := signalInfoBarCloseMap[index].callback
	callback()
}

type signalInfoBarResponseDetail struct {
	callback  InfoBarSignalResponseCallback
	handlerID C.gulong
}

var signalInfoBarResponseId int
var signalInfoBarResponseMap = make(map[int]signalInfoBarResponseDetail)
var signalInfoBarResponseLock sync.RWMutex

// InfoBarSignalResponseCallback is a callback function for a 'response' signal emitted from a InfoBar.
type InfoBarSignalResponseCallback func(responseId int32)

/*
ConnectResponse connects the callback to the 'response' signal for the InfoBar.

The returned value represents the connection, and may be passed to DisconnectResponse to remove it.
*/
func (recv *InfoBar) ConnectResponse(callback InfoBarSignalResponseCallback) int {
	signalInfoBarResponseLock.Lock()
	defer signalInfoBarResponseLock.Unlock()

	signalInfoBarResponseId++
	instance := C.gpointer(recv.native)
	handlerID := C.InfoBar_signal_connect_response(instance, C.gpointer(uintptr(signalInfoBarResponseId)))

	detail := signalInfoBarResponseDetail{callback, handlerID}
	signalInfoBarResponseMap[signalInfoBarResponseId] = detail

	return signalInfoBarResponseId
}

/*
DisconnectResponse disconnects a callback from the 'response' signal for the InfoBar.

The connectionID should be a value returned from a call to ConnectResponse.
*/
func (recv *InfoBar) DisconnectResponse(connectionID int) {
	signalInfoBarResponseLock.Lock()
	defer signalInfoBarResponseLock.Unlock()

	detail, exists := signalInfoBarResponseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalInfoBarResponseMap, connectionID)
}

//export infobar_responseHandler
func infobar_responseHandler(_ *C.GObject, c_response_id C.gint, data C.gpointer) {
	signalInfoBarResponseLock.RLock()
	defer signalInfoBarResponseLock.RUnlock()

	responseId := int32(c_response_id)

	index := int(uintptr(data))
	callback := signalInfoBarResponseMap[index].callback
	callback(responseId)
}

// Blacklisted : gtk_info_bar_new

// Blacklisted : gtk_info_bar_add_action_widget

// Blacklisted : gtk_info_bar_add_button

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// Blacklisted : gtk_info_bar_get_action_area

// Blacklisted : gtk_info_bar_get_content_area

// Blacklisted : gtk_info_bar_get_message_type

// Blacklisted : gtk_info_bar_response

// Blacklisted : gtk_info_bar_set_default_response

// Blacklisted : gtk_info_bar_set_message_type

// Blacklisted : gtk_info_bar_set_response_sensitive

type signalLabelActivateCurrentLinkDetail struct {
	callback  LabelSignalActivateCurrentLinkCallback
	handlerID C.gulong
}

var signalLabelActivateCurrentLinkId int
var signalLabelActivateCurrentLinkMap = make(map[int]signalLabelActivateCurrentLinkDetail)
var signalLabelActivateCurrentLinkLock sync.RWMutex

// LabelSignalActivateCurrentLinkCallback is a callback function for a 'activate-current-link' signal emitted from a Label.
type LabelSignalActivateCurrentLinkCallback func()

/*
ConnectActivateCurrentLink connects the callback to the 'activate-current-link' signal for the Label.

The returned value represents the connection, and may be passed to DisconnectActivateCurrentLink to remove it.
*/
func (recv *Label) ConnectActivateCurrentLink(callback LabelSignalActivateCurrentLinkCallback) int {
	signalLabelActivateCurrentLinkLock.Lock()
	defer signalLabelActivateCurrentLinkLock.Unlock()

	signalLabelActivateCurrentLinkId++
	instance := C.gpointer(recv.native)
	handlerID := C.Label_signal_connect_activate_current_link(instance, C.gpointer(uintptr(signalLabelActivateCurrentLinkId)))

	detail := signalLabelActivateCurrentLinkDetail{callback, handlerID}
	signalLabelActivateCurrentLinkMap[signalLabelActivateCurrentLinkId] = detail

	return signalLabelActivateCurrentLinkId
}

/*
DisconnectActivateCurrentLink disconnects a callback from the 'activate-current-link' signal for the Label.

The connectionID should be a value returned from a call to ConnectActivateCurrentLink.
*/
func (recv *Label) DisconnectActivateCurrentLink(connectionID int) {
	signalLabelActivateCurrentLinkLock.Lock()
	defer signalLabelActivateCurrentLinkLock.Unlock()

	detail, exists := signalLabelActivateCurrentLinkMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalLabelActivateCurrentLinkMap, connectionID)
}

//export label_activateCurrentLinkHandler
func label_activateCurrentLinkHandler(_ *C.GObject, data C.gpointer) {
	signalLabelActivateCurrentLinkLock.RLock()
	defer signalLabelActivateCurrentLinkLock.RUnlock()

	index := int(uintptr(data))
	callback := signalLabelActivateCurrentLinkMap[index].callback
	callback()
}

type signalLabelActivateLinkDetail struct {
	callback  LabelSignalActivateLinkCallback
	handlerID C.gulong
}

var signalLabelActivateLinkId int
var signalLabelActivateLinkMap = make(map[int]signalLabelActivateLinkDetail)
var signalLabelActivateLinkLock sync.RWMutex

// LabelSignalActivateLinkCallback is a callback function for a 'activate-link' signal emitted from a Label.
type LabelSignalActivateLinkCallback func(uri string) bool

/*
ConnectActivateLink connects the callback to the 'activate-link' signal for the Label.

The returned value represents the connection, and may be passed to DisconnectActivateLink to remove it.
*/
func (recv *Label) ConnectActivateLink(callback LabelSignalActivateLinkCallback) int {
	signalLabelActivateLinkLock.Lock()
	defer signalLabelActivateLinkLock.Unlock()

	signalLabelActivateLinkId++
	instance := C.gpointer(recv.native)
	handlerID := C.Label_signal_connect_activate_link(instance, C.gpointer(uintptr(signalLabelActivateLinkId)))

	detail := signalLabelActivateLinkDetail{callback, handlerID}
	signalLabelActivateLinkMap[signalLabelActivateLinkId] = detail

	return signalLabelActivateLinkId
}

/*
DisconnectActivateLink disconnects a callback from the 'activate-link' signal for the Label.

The connectionID should be a value returned from a call to ConnectActivateLink.
*/
func (recv *Label) DisconnectActivateLink(connectionID int) {
	signalLabelActivateLinkLock.Lock()
	defer signalLabelActivateLinkLock.Unlock()

	detail, exists := signalLabelActivateLinkMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalLabelActivateLinkMap, connectionID)
}

//export label_activateLinkHandler
func label_activateLinkHandler(_ *C.GObject, c_uri *C.gchar, data C.gpointer) C.gboolean {
	signalLabelActivateLinkLock.RLock()
	defer signalLabelActivateLinkLock.RUnlock()

	uri := C.GoString(c_uri)

	index := int(uintptr(data))
	callback := signalLabelActivateLinkMap[index].callback
	retGo := callback(uri)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_label_get_current_uri

// Blacklisted : gtk_label_get_track_visited_links

// Blacklisted : gtk_label_set_track_visited_links

// Blacklisted : gtk_menu_get_reserve_toggle_size

// Blacklisted : gtk_menu_set_reserve_toggle_size

type signalPrintOperationUpdateCustomWidgetDetail struct {
	callback  PrintOperationSignalUpdateCustomWidgetCallback
	handlerID C.gulong
}

var signalPrintOperationUpdateCustomWidgetId int
var signalPrintOperationUpdateCustomWidgetMap = make(map[int]signalPrintOperationUpdateCustomWidgetDetail)
var signalPrintOperationUpdateCustomWidgetLock sync.RWMutex

// PrintOperationSignalUpdateCustomWidgetCallback is a callback function for a 'update-custom-widget' signal emitted from a PrintOperation.
type PrintOperationSignalUpdateCustomWidgetCallback func(widget *Widget, setup *PageSetup, settings *PrintSettings)

/*
ConnectUpdateCustomWidget connects the callback to the 'update-custom-widget' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectUpdateCustomWidget to remove it.
*/
func (recv *PrintOperation) ConnectUpdateCustomWidget(callback PrintOperationSignalUpdateCustomWidgetCallback) int {
	signalPrintOperationUpdateCustomWidgetLock.Lock()
	defer signalPrintOperationUpdateCustomWidgetLock.Unlock()

	signalPrintOperationUpdateCustomWidgetId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_update_custom_widget(instance, C.gpointer(uintptr(signalPrintOperationUpdateCustomWidgetId)))

	detail := signalPrintOperationUpdateCustomWidgetDetail{callback, handlerID}
	signalPrintOperationUpdateCustomWidgetMap[signalPrintOperationUpdateCustomWidgetId] = detail

	return signalPrintOperationUpdateCustomWidgetId
}

/*
DisconnectUpdateCustomWidget disconnects a callback from the 'update-custom-widget' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectUpdateCustomWidget.
*/
func (recv *PrintOperation) DisconnectUpdateCustomWidget(connectionID int) {
	signalPrintOperationUpdateCustomWidgetLock.Lock()
	defer signalPrintOperationUpdateCustomWidgetLock.Unlock()

	detail, exists := signalPrintOperationUpdateCustomWidgetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationUpdateCustomWidgetMap, connectionID)
}

//export printoperation_updateCustomWidgetHandler
func printoperation_updateCustomWidgetHandler(_ *C.GObject, c_widget *C.GtkWidget, c_setup *C.GtkPageSetup, c_settings *C.GtkPrintSettings, data C.gpointer) {
	signalPrintOperationUpdateCustomWidgetLock.RLock()
	defer signalPrintOperationUpdateCustomWidgetLock.RUnlock()

	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

	setup := PageSetupNewFromC(unsafe.Pointer(c_setup))

	settings := PrintSettingsNewFromC(unsafe.Pointer(c_settings))

	index := int(uintptr(data))
	callback := signalPrintOperationUpdateCustomWidgetMap[index].callback
	callback(widget, setup, settings)
}

// Blacklisted : gtk_print_operation_get_embed_page_setup

// Blacklisted : gtk_print_operation_get_has_selection

// Blacklisted : gtk_print_operation_get_n_pages_to_print

// Blacklisted : gtk_print_operation_get_support_selection

// Blacklisted : gtk_print_operation_set_embed_page_setup

// Blacklisted : gtk_print_operation_set_has_selection

// Blacklisted : gtk_print_operation_set_support_selection

// Blacklisted : gtk_range_get_flippable

// Blacklisted : gtk_range_set_flippable

// GetTitle is a wrapper around the C function gtk_status_icon_get_title.
func (recv *StatusIcon) GetTitle() string {
	retC := C.gtk_status_icon_get_title((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetTitle is a wrapper around the C function gtk_status_icon_set_title.
func (recv *StatusIcon) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_status_icon_set_title((*C.GtkStatusIcon)(recv.native), c_title)

	return
}

// Blacklisted : gtk_widget_get_allocation

// Blacklisted : gtk_widget_get_app_paintable

// Blacklisted : gtk_widget_get_can_default

// Blacklisted : gtk_widget_get_can_focus

// Blacklisted : gtk_widget_get_double_buffered

// Blacklisted : gtk_widget_get_has_window

// Blacklisted : gtk_widget_get_receives_default

// Blacklisted : gtk_widget_get_sensitive

// Blacklisted : gtk_widget_get_state

// Blacklisted : gtk_widget_get_visible

// Blacklisted : gtk_widget_has_default

// Blacklisted : gtk_widget_has_focus

// Blacklisted : gtk_widget_has_grab

// Blacklisted : gtk_widget_is_drawable

// Blacklisted : gtk_widget_is_sensitive

// Blacklisted : gtk_widget_is_toplevel

// Blacklisted : gtk_widget_set_allocation

// Blacklisted : gtk_widget_set_can_default

// Blacklisted : gtk_widget_set_can_focus

// Blacklisted : gtk_widget_set_has_window

// Blacklisted : gtk_widget_set_receives_default

// Blacklisted : gtk_widget_set_visible

// Blacklisted : gtk_widget_set_window

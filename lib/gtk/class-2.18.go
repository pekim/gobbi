// This is a generated file - DO NOT EDIT
// +build gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// GetAlignment is a wrapper around the C function gtk_cell_renderer_get_alignment.
func (recv *CellRenderer) GetAlignment() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_cell_renderer_get_alignment((*C.GtkCellRenderer)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// GetPadding is a wrapper around the C function gtk_cell_renderer_get_padding.
func (recv *CellRenderer) GetPadding() (int32, int32) {
	var c_xpad C.gint

	var c_ypad C.gint

	C.gtk_cell_renderer_get_padding((*C.GtkCellRenderer)(recv.native), &c_xpad, &c_ypad)

	xpad := (int32)(c_xpad)

	ypad := (int32)(c_ypad)

	return xpad, ypad
}

// GetSensitive is a wrapper around the C function gtk_cell_renderer_get_sensitive.
func (recv *CellRenderer) GetSensitive() bool {
	retC := C.gtk_cell_renderer_get_sensitive((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisible is a wrapper around the C function gtk_cell_renderer_get_visible.
func (recv *CellRenderer) GetVisible() bool {
	retC := C.gtk_cell_renderer_get_visible((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlignment is a wrapper around the C function gtk_cell_renderer_set_alignment.
func (recv *CellRenderer) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_cell_renderer_set_alignment((*C.GtkCellRenderer)(recv.native), c_xalign, c_yalign)

	return
}

// SetPadding is a wrapper around the C function gtk_cell_renderer_set_padding.
func (recv *CellRenderer) SetPadding(xpad int32, ypad int32) {
	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_cell_renderer_set_padding((*C.GtkCellRenderer)(recv.native), c_xpad, c_ypad)

	return
}

// SetSensitive is a wrapper around the C function gtk_cell_renderer_set_sensitive.
func (recv *CellRenderer) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_cell_renderer_set_sensitive((*C.GtkCellRenderer)(recv.native), c_sensitive)

	return
}

// SetVisible is a wrapper around the C function gtk_cell_renderer_set_visible.
func (recv *CellRenderer) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_cell_renderer_set_visible((*C.GtkCellRenderer)(recv.native), c_visible)

	return
}

// GetActivatable is a wrapper around the C function gtk_cell_renderer_toggle_get_activatable.
func (recv *CellRendererToggle) GetActivatable() bool {
	retC := C.gtk_cell_renderer_toggle_get_activatable((*C.GtkCellRendererToggle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActivatable is a wrapper around the C function gtk_cell_renderer_toggle_set_activatable.
func (recv *CellRendererToggle) SetActivatable(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_cell_renderer_toggle_set_activatable((*C.GtkCellRendererToggle)(recv.native), c_setting)

	return
}

// EntryNewWithBuffer is a wrapper around the C function gtk_entry_new_with_buffer.
func EntryNewWithBuffer(buffer *EntryBuffer) *Entry {
	c_buffer := (*C.GtkEntryBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkEntryBuffer)(buffer.ToC())
	}

	retC := C.gtk_entry_new_with_buffer(c_buffer)
	retGo := EntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBuffer is a wrapper around the C function gtk_entry_get_buffer.
func (recv *Entry) GetBuffer() *EntryBuffer {
	retC := C.gtk_entry_get_buffer((*C.GtkEntry)(recv.native))
	retGo := EntryBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetBuffer is a wrapper around the C function gtk_entry_set_buffer.
func (recv *Entry) SetBuffer(buffer *EntryBuffer) {
	c_buffer := (*C.GtkEntryBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkEntryBuffer)(buffer.ToC())
	}

	C.gtk_entry_set_buffer((*C.GtkEntry)(recv.native), c_buffer)

	return
}

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

// EntryBufferNew is a wrapper around the C function gtk_entry_buffer_new.
func EntryBufferNew(initialChars string, nInitialChars int32) *EntryBuffer {
	c_initial_chars := C.CString(initialChars)
	defer C.free(unsafe.Pointer(c_initial_chars))

	c_n_initial_chars := (C.gint)(nInitialChars)

	retC := C.gtk_entry_buffer_new(c_initial_chars, c_n_initial_chars)
	retGo := EntryBufferNewFromC(unsafe.Pointer(retC))

	C.g_object_unref((C.gpointer)(retC))

	return retGo
}

// DeleteText is a wrapper around the C function gtk_entry_buffer_delete_text.
func (recv *EntryBuffer) DeleteText(position uint32, nChars int32) uint32 {
	c_position := (C.guint)(position)

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_entry_buffer_delete_text((*C.GtkEntryBuffer)(recv.native), c_position, c_n_chars)
	retGo := (uint32)(retC)

	return retGo
}

// EmitDeletedText is a wrapper around the C function gtk_entry_buffer_emit_deleted_text.
func (recv *EntryBuffer) EmitDeletedText(position uint32, nChars uint32) {
	c_position := (C.guint)(position)

	c_n_chars := (C.guint)(nChars)

	C.gtk_entry_buffer_emit_deleted_text((*C.GtkEntryBuffer)(recv.native), c_position, c_n_chars)

	return
}

// EmitInsertedText is a wrapper around the C function gtk_entry_buffer_emit_inserted_text.
func (recv *EntryBuffer) EmitInsertedText(position uint32, chars string, nChars uint32) {
	c_position := (C.guint)(position)

	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.guint)(nChars)

	C.gtk_entry_buffer_emit_inserted_text((*C.GtkEntryBuffer)(recv.native), c_position, c_chars, c_n_chars)

	return
}

// GetBytes is a wrapper around the C function gtk_entry_buffer_get_bytes.
func (recv *EntryBuffer) GetBytes() uint64 {
	retC := C.gtk_entry_buffer_get_bytes((*C.GtkEntryBuffer)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetLength is a wrapper around the C function gtk_entry_buffer_get_length.
func (recv *EntryBuffer) GetLength() uint32 {
	retC := C.gtk_entry_buffer_get_length((*C.GtkEntryBuffer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMaxLength is a wrapper around the C function gtk_entry_buffer_get_max_length.
func (recv *EntryBuffer) GetMaxLength() int32 {
	retC := C.gtk_entry_buffer_get_max_length((*C.GtkEntryBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetText is a wrapper around the C function gtk_entry_buffer_get_text.
func (recv *EntryBuffer) GetText() string {
	retC := C.gtk_entry_buffer_get_text((*C.GtkEntryBuffer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// InsertText is a wrapper around the C function gtk_entry_buffer_insert_text.
func (recv *EntryBuffer) InsertText(position uint32, chars string, nChars int32) uint32 {
	c_position := (C.guint)(position)

	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_entry_buffer_insert_text((*C.GtkEntryBuffer)(recv.native), c_position, c_chars, c_n_chars)
	retGo := (uint32)(retC)

	return retGo
}

// SetMaxLength is a wrapper around the C function gtk_entry_buffer_set_max_length.
func (recv *EntryBuffer) SetMaxLength(maxLength int32) {
	c_max_length := (C.gint)(maxLength)

	C.gtk_entry_buffer_set_max_length((*C.GtkEntryBuffer)(recv.native), c_max_length)

	return
}

// SetText is a wrapper around the C function gtk_entry_buffer_set_text.
func (recv *EntryBuffer) SetText(chars string, nChars int32) {
	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_buffer_set_text((*C.GtkEntryBuffer)(recv.native), c_chars, c_n_chars)

	return
}

// GetItemPadding is a wrapper around the C function gtk_icon_view_get_item_padding.
func (recv *IconView) GetItemPadding() int32 {
	retC := C.gtk_icon_view_get_item_padding((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetItemPadding is a wrapper around the C function gtk_icon_view_set_item_padding.
func (recv *IconView) SetItemPadding(itemPadding int32) {
	c_item_padding := (C.gint)(itemPadding)

	C.gtk_icon_view_set_item_padding((*C.GtkIconView)(recv.native), c_item_padding)

	return
}

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

// InfoBarNew is a wrapper around the C function gtk_info_bar_new.
func InfoBarNew() *InfoBar {
	retC := C.gtk_info_bar_new()
	retGo := InfoBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddActionWidget is a wrapper around the C function gtk_info_bar_add_action_widget.
func (recv *InfoBar) AddActionWidget(child *Widget, responseId int32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_add_action_widget((*C.GtkInfoBar)(recv.native), c_child, c_response_id)

	return
}

// AddButton is a wrapper around the C function gtk_info_bar_add_button.
func (recv *InfoBar) AddButton(buttonText string, responseId int32) *Button {
	c_button_text := C.CString(buttonText)
	defer C.free(unsafe.Pointer(c_button_text))

	c_response_id := (C.gint)(responseId)

	retC := C.gtk_info_bar_add_button((*C.GtkInfoBar)(recv.native), c_button_text, c_response_id)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// GetActionArea is a wrapper around the C function gtk_info_bar_get_action_area.
func (recv *InfoBar) GetActionArea() *Widget {
	retC := C.gtk_info_bar_get_action_area((*C.GtkInfoBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContentArea is a wrapper around the C function gtk_info_bar_get_content_area.
func (recv *InfoBar) GetContentArea() *Widget {
	retC := C.gtk_info_bar_get_content_area((*C.GtkInfoBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMessageType is a wrapper around the C function gtk_info_bar_get_message_type.
func (recv *InfoBar) GetMessageType() MessageType {
	retC := C.gtk_info_bar_get_message_type((*C.GtkInfoBar)(recv.native))
	retGo := (MessageType)(retC)

	return retGo
}

// Response is a wrapper around the C function gtk_info_bar_response.
func (recv *InfoBar) Response(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_response((*C.GtkInfoBar)(recv.native), c_response_id)

	return
}

// SetDefaultResponse is a wrapper around the C function gtk_info_bar_set_default_response.
func (recv *InfoBar) SetDefaultResponse(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_set_default_response((*C.GtkInfoBar)(recv.native), c_response_id)

	return
}

// SetMessageType is a wrapper around the C function gtk_info_bar_set_message_type.
func (recv *InfoBar) SetMessageType(messageType MessageType) {
	c_message_type := (C.GtkMessageType)(messageType)

	C.gtk_info_bar_set_message_type((*C.GtkInfoBar)(recv.native), c_message_type)

	return
}

// SetResponseSensitive is a wrapper around the C function gtk_info_bar_set_response_sensitive.
func (recv *InfoBar) SetResponseSensitive(responseId int32, setting bool) {
	c_response_id := (C.gint)(responseId)

	c_setting :=
		boolToGboolean(setting)

	C.gtk_info_bar_set_response_sensitive((*C.GtkInfoBar)(recv.native), c_response_id, c_setting)

	return
}

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

// GetCurrentUri is a wrapper around the C function gtk_label_get_current_uri.
func (recv *Label) GetCurrentUri() string {
	retC := C.gtk_label_get_current_uri((*C.GtkLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTrackVisitedLinks is a wrapper around the C function gtk_label_get_track_visited_links.
func (recv *Label) GetTrackVisitedLinks() bool {
	retC := C.gtk_label_get_track_visited_links((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetTrackVisitedLinks is a wrapper around the C function gtk_label_set_track_visited_links.
func (recv *Label) SetTrackVisitedLinks(trackLinks bool) {
	c_track_links :=
		boolToGboolean(trackLinks)

	C.gtk_label_set_track_visited_links((*C.GtkLabel)(recv.native), c_track_links)

	return
}

// GetReserveToggleSize is a wrapper around the C function gtk_menu_get_reserve_toggle_size.
func (recv *Menu) GetReserveToggleSize() bool {
	retC := C.gtk_menu_get_reserve_toggle_size((*C.GtkMenu)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetReserveToggleSize is a wrapper around the C function gtk_menu_set_reserve_toggle_size.
func (recv *Menu) SetReserveToggleSize(reserveToggleSize bool) {
	c_reserve_toggle_size :=
		boolToGboolean(reserveToggleSize)

	C.gtk_menu_set_reserve_toggle_size((*C.GtkMenu)(recv.native), c_reserve_toggle_size)

	return
}

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

// GetEmbedPageSetup is a wrapper around the C function gtk_print_operation_get_embed_page_setup.
func (recv *PrintOperation) GetEmbedPageSetup() bool {
	retC := C.gtk_print_operation_get_embed_page_setup((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasSelection is a wrapper around the C function gtk_print_operation_get_has_selection.
func (recv *PrintOperation) GetHasSelection() bool {
	retC := C.gtk_print_operation_get_has_selection((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetNPagesToPrint is a wrapper around the C function gtk_print_operation_get_n_pages_to_print.
func (recv *PrintOperation) GetNPagesToPrint() int32 {
	retC := C.gtk_print_operation_get_n_pages_to_print((*C.GtkPrintOperation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSupportSelection is a wrapper around the C function gtk_print_operation_get_support_selection.
func (recv *PrintOperation) GetSupportSelection() bool {
	retC := C.gtk_print_operation_get_support_selection((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetEmbedPageSetup is a wrapper around the C function gtk_print_operation_set_embed_page_setup.
func (recv *PrintOperation) SetEmbedPageSetup(embed bool) {
	c_embed :=
		boolToGboolean(embed)

	C.gtk_print_operation_set_embed_page_setup((*C.GtkPrintOperation)(recv.native), c_embed)

	return
}

// SetHasSelection is a wrapper around the C function gtk_print_operation_set_has_selection.
func (recv *PrintOperation) SetHasSelection(hasSelection bool) {
	c_has_selection :=
		boolToGboolean(hasSelection)

	C.gtk_print_operation_set_has_selection((*C.GtkPrintOperation)(recv.native), c_has_selection)

	return
}

// SetSupportSelection is a wrapper around the C function gtk_print_operation_set_support_selection.
func (recv *PrintOperation) SetSupportSelection(supportSelection bool) {
	c_support_selection :=
		boolToGboolean(supportSelection)

	C.gtk_print_operation_set_support_selection((*C.GtkPrintOperation)(recv.native), c_support_selection)

	return
}

// GetFlippable is a wrapper around the C function gtk_range_get_flippable.
func (recv *Range) GetFlippable() bool {
	retC := C.gtk_range_get_flippable((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFlippable is a wrapper around the C function gtk_range_set_flippable.
func (recv *Range) SetFlippable(flippable bool) {
	c_flippable :=
		boolToGboolean(flippable)

	C.gtk_range_set_flippable((*C.GtkRange)(recv.native), c_flippable)

	return
}

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

// Unsupported : gtk_widget_get_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// GetAppPaintable is a wrapper around the C function gtk_widget_get_app_paintable.
func (recv *Widget) GetAppPaintable() bool {
	retC := C.gtk_widget_get_app_paintable((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanDefault is a wrapper around the C function gtk_widget_get_can_default.
func (recv *Widget) GetCanDefault() bool {
	retC := C.gtk_widget_get_can_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanFocus is a wrapper around the C function gtk_widget_get_can_focus.
func (recv *Widget) GetCanFocus() bool {
	retC := C.gtk_widget_get_can_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDoubleBuffered is a wrapper around the C function gtk_widget_get_double_buffered.
func (recv *Widget) GetDoubleBuffered() bool {
	retC := C.gtk_widget_get_double_buffered((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasWindow is a wrapper around the C function gtk_widget_get_has_window.
func (recv *Widget) GetHasWindow() bool {
	retC := C.gtk_widget_get_has_window((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetReceivesDefault is a wrapper around the C function gtk_widget_get_receives_default.
func (recv *Widget) GetReceivesDefault() bool {
	retC := C.gtk_widget_get_receives_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSensitive is a wrapper around the C function gtk_widget_get_sensitive.
func (recv *Widget) GetSensitive() bool {
	retC := C.gtk_widget_get_sensitive((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetState is a wrapper around the C function gtk_widget_get_state.
func (recv *Widget) GetState() StateType {
	retC := C.gtk_widget_get_state((*C.GtkWidget)(recv.native))
	retGo := (StateType)(retC)

	return retGo
}

// GetVisible is a wrapper around the C function gtk_widget_get_visible.
func (recv *Widget) GetVisible() bool {
	retC := C.gtk_widget_get_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasDefault is a wrapper around the C function gtk_widget_has_default.
func (recv *Widget) HasDefault() bool {
	retC := C.gtk_widget_has_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasFocus is a wrapper around the C function gtk_widget_has_focus.
func (recv *Widget) HasFocus() bool {
	retC := C.gtk_widget_has_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasGrab is a wrapper around the C function gtk_widget_has_grab.
func (recv *Widget) HasGrab() bool {
	retC := C.gtk_widget_has_grab((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsDrawable is a wrapper around the C function gtk_widget_is_drawable.
func (recv *Widget) IsDrawable() bool {
	retC := C.gtk_widget_is_drawable((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSensitive is a wrapper around the C function gtk_widget_is_sensitive.
func (recv *Widget) IsSensitive() bool {
	retC := C.gtk_widget_is_sensitive((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsToplevel is a wrapper around the C function gtk_widget_is_toplevel.
func (recv *Widget) IsToplevel() bool {
	retC := C.gtk_widget_is_toplevel((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_set_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// SetCanDefault is a wrapper around the C function gtk_widget_set_can_default.
func (recv *Widget) SetCanDefault(canDefault bool) {
	c_can_default :=
		boolToGboolean(canDefault)

	C.gtk_widget_set_can_default((*C.GtkWidget)(recv.native), c_can_default)

	return
}

// SetCanFocus is a wrapper around the C function gtk_widget_set_can_focus.
func (recv *Widget) SetCanFocus(canFocus bool) {
	c_can_focus :=
		boolToGboolean(canFocus)

	C.gtk_widget_set_can_focus((*C.GtkWidget)(recv.native), c_can_focus)

	return
}

// SetHasWindow is a wrapper around the C function gtk_widget_set_has_window.
func (recv *Widget) SetHasWindow(hasWindow bool) {
	c_has_window :=
		boolToGboolean(hasWindow)

	C.gtk_widget_set_has_window((*C.GtkWidget)(recv.native), c_has_window)

	return
}

// SetReceivesDefault is a wrapper around the C function gtk_widget_set_receives_default.
func (recv *Widget) SetReceivesDefault(receivesDefault bool) {
	c_receives_default :=
		boolToGboolean(receivesDefault)

	C.gtk_widget_set_receives_default((*C.GtkWidget)(recv.native), c_receives_default)

	return
}

// SetVisible is a wrapper around the C function gtk_widget_set_visible.
func (recv *Widget) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_widget_set_visible((*C.GtkWidget)(recv.native), c_visible)

	return
}

// SetWindow is a wrapper around the C function gtk_widget_set_window.
func (recv *Widget) SetWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_widget_set_window((*C.GtkWidget)(recv.native), c_window)

	return
}

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

	void infobar_closeHandler(GObject *, gpointer);

	static gulong InfoBar_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", G_CALLBACK(infobar_closeHandler), data);
	}

*/
/*

	void label_activateCurrentLinkHandler(GObject *, gpointer);

	static gulong Label_signal_connect_activate_current_link(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-current-link", G_CALLBACK(label_activateCurrentLinkHandler), data);
	}

*/
/*

	void printoperation_updateCustomWidgetHandler(GObject *, GtkWidget *, GtkPageSetup *, GtkPrintSettings *, gpointer);

	static gulong PrintOperation_signal_connect_update_custom_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "update-custom-widget", G_CALLBACK(printoperation_updateCustomWidgetHandler), data);
	}

*/
import "C"

// Fills in @xalign and @yalign with the appropriate values of @cell.
/*

C function

gtk_cell_renderer_get_alignment
*/
func (recv *CellRenderer) GetAlignment() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_cell_renderer_get_alignment((*C.GtkCellRenderer)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// Fills in @xpad and @ypad with the appropriate values of @cell.
/*

C function

gtk_cell_renderer_get_padding
*/
func (recv *CellRenderer) GetPadding() (int32, int32) {
	var c_xpad C.gint

	var c_ypad C.gint

	C.gtk_cell_renderer_get_padding((*C.GtkCellRenderer)(recv.native), &c_xpad, &c_ypad)

	xpad := (int32)(c_xpad)

	ypad := (int32)(c_ypad)

	return xpad, ypad
}

// Returns the cell renderer’s sensitivity.
/*

C function

gtk_cell_renderer_get_sensitive
*/
func (recv *CellRenderer) GetSensitive() bool {
	retC := C.gtk_cell_renderer_get_sensitive((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the cell renderer’s visibility.
/*

C function

gtk_cell_renderer_get_visible
*/
func (recv *CellRenderer) GetVisible() bool {
	retC := C.gtk_cell_renderer_get_visible((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the renderer’s alignment within its available space.
/*

C function

gtk_cell_renderer_set_alignment
*/
func (recv *CellRenderer) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_cell_renderer_set_alignment((*C.GtkCellRenderer)(recv.native), c_xalign, c_yalign)

	return
}

// Sets the renderer’s padding.
/*

C function

gtk_cell_renderer_set_padding
*/
func (recv *CellRenderer) SetPadding(xpad int32, ypad int32) {
	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_cell_renderer_set_padding((*C.GtkCellRenderer)(recv.native), c_xpad, c_ypad)

	return
}

// Sets the cell renderer’s sensitivity.
/*

C function

gtk_cell_renderer_set_sensitive
*/
func (recv *CellRenderer) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_cell_renderer_set_sensitive((*C.GtkCellRenderer)(recv.native), c_sensitive)

	return
}

// Sets the cell renderer’s visibility.
/*

C function

gtk_cell_renderer_set_visible
*/
func (recv *CellRenderer) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_cell_renderer_set_visible((*C.GtkCellRenderer)(recv.native), c_visible)

	return
}

// Returns whether the cell renderer is activatable. See
// gtk_cell_renderer_toggle_set_activatable().
/*

C function

gtk_cell_renderer_toggle_get_activatable
*/
func (recv *CellRendererToggle) GetActivatable() bool {
	retC := C.gtk_cell_renderer_toggle_get_activatable((*C.GtkCellRendererToggle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Makes the cell renderer activatable.
/*

C function

gtk_cell_renderer_toggle_set_activatable
*/
func (recv *CellRendererToggle) SetActivatable(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_cell_renderer_toggle_set_activatable((*C.GtkCellRendererToggle)(recv.native), c_setting)

	return
}

// Creates a new entry with the specified text buffer.
/*

C function

gtk_entry_new_with_buffer
*/
func EntryNewWithBuffer(buffer *EntryBuffer) *Entry {
	c_buffer := (*C.GtkEntryBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkEntryBuffer)(buffer.ToC())
	}

	retC := C.gtk_entry_new_with_buffer(c_buffer)
	retGo := EntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get the #GtkEntryBuffer object which holds the text for
// this widget.
/*

C function

gtk_entry_get_buffer
*/
func (recv *Entry) GetBuffer() *EntryBuffer {
	retC := C.gtk_entry_get_buffer((*C.GtkEntry)(recv.native))
	retGo := EntryBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set the #GtkEntryBuffer object which holds the text for
// this widget.
/*

C function

gtk_entry_set_buffer
*/
func (recv *Entry) SetBuffer(buffer *EntryBuffer) {
	c_buffer := (*C.GtkEntryBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkEntryBuffer)(buffer.ToC())
	}

	C.gtk_entry_set_buffer((*C.GtkEntry)(recv.native), c_buffer)

	return
}

// Unsupported signal 'deleted-text' for EntryBuffer : unsupported parameter position : type guint :

// Unsupported signal 'inserted-text' for EntryBuffer : unsupported parameter position : type guint :

// Create a new GtkEntryBuffer object.
//
// Optionally, specify initial text to set in the buffer.
/*

C function

gtk_entry_buffer_new
*/
func EntryBufferNew(initialChars string, nInitialChars int32) *EntryBuffer {
	c_initial_chars := C.CString(initialChars)
	defer C.free(unsafe.Pointer(c_initial_chars))

	c_n_initial_chars := (C.gint)(nInitialChars)

	retC := C.gtk_entry_buffer_new(c_initial_chars, c_n_initial_chars)
	retGo := EntryBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Deletes a sequence of characters from the buffer. @n_chars characters are
// deleted starting at @position. If @n_chars is negative, then all characters
// until the end of the text are deleted.
//
// If @position or @n_chars are out of bounds, then they are coerced to sane
// values.
//
// Note that the positions are specified in characters, not bytes.
/*

C function

gtk_entry_buffer_delete_text
*/
func (recv *EntryBuffer) DeleteText(position uint32, nChars int32) uint32 {
	c_position := (C.guint)(position)

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_entry_buffer_delete_text((*C.GtkEntryBuffer)(recv.native), c_position, c_n_chars)
	retGo := (uint32)(retC)

	return retGo
}

// Used when subclassing #GtkEntryBuffer
/*

C function

gtk_entry_buffer_emit_deleted_text
*/
func (recv *EntryBuffer) EmitDeletedText(position uint32, nChars uint32) {
	c_position := (C.guint)(position)

	c_n_chars := (C.guint)(nChars)

	C.gtk_entry_buffer_emit_deleted_text((*C.GtkEntryBuffer)(recv.native), c_position, c_n_chars)

	return
}

// Used when subclassing #GtkEntryBuffer
/*

C function

gtk_entry_buffer_emit_inserted_text
*/
func (recv *EntryBuffer) EmitInsertedText(position uint32, chars string, nChars uint32) {
	c_position := (C.guint)(position)

	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.guint)(nChars)

	C.gtk_entry_buffer_emit_inserted_text((*C.GtkEntryBuffer)(recv.native), c_position, c_chars, c_n_chars)

	return
}

// Retrieves the length in bytes of the buffer.
// See gtk_entry_buffer_get_length().
/*

C function

gtk_entry_buffer_get_bytes
*/
func (recv *EntryBuffer) GetBytes() uint64 {
	retC := C.gtk_entry_buffer_get_bytes((*C.GtkEntryBuffer)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Retrieves the length in characters of the buffer.
/*

C function

gtk_entry_buffer_get_length
*/
func (recv *EntryBuffer) GetLength() uint32 {
	retC := C.gtk_entry_buffer_get_length((*C.GtkEntryBuffer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Retrieves the maximum allowed length of the text in
// @buffer. See gtk_entry_buffer_set_max_length().
/*

C function

gtk_entry_buffer_get_max_length
*/
func (recv *EntryBuffer) GetMaxLength() int32 {
	retC := C.gtk_entry_buffer_get_max_length((*C.GtkEntryBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the contents of the buffer.
//
// The memory pointer returned by this call will not change
// unless this object emits a signal, or is finalized.
/*

C function

gtk_entry_buffer_get_text
*/
func (recv *EntryBuffer) GetText() string {
	retC := C.gtk_entry_buffer_get_text((*C.GtkEntryBuffer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Inserts @n_chars characters of @chars into the contents of the
// buffer, at position @position.
//
// If @n_chars is negative, then characters from chars will be inserted
// until a null-terminator is found. If @position or @n_chars are out of
// bounds, or the maximum buffer text length is exceeded, then they are
// coerced to sane values.
//
// Note that the position and length are in characters, not in bytes.
/*

C function

gtk_entry_buffer_insert_text
*/
func (recv *EntryBuffer) InsertText(position uint32, chars string, nChars int32) uint32 {
	c_position := (C.guint)(position)

	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_entry_buffer_insert_text((*C.GtkEntryBuffer)(recv.native), c_position, c_chars, c_n_chars)
	retGo := (uint32)(retC)

	return retGo
}

// Sets the maximum allowed length of the contents of the buffer. If
// the current contents are longer than the given length, then they
// will be truncated to fit.
/*

C function

gtk_entry_buffer_set_max_length
*/
func (recv *EntryBuffer) SetMaxLength(maxLength int32) {
	c_max_length := (C.gint)(maxLength)

	C.gtk_entry_buffer_set_max_length((*C.GtkEntryBuffer)(recv.native), c_max_length)

	return
}

// Sets the text in the buffer.
//
// This is roughly equivalent to calling gtk_entry_buffer_delete_text()
// and gtk_entry_buffer_insert_text().
//
// Note that @n_chars is in characters, not in bytes.
/*

C function

gtk_entry_buffer_set_text
*/
func (recv *EntryBuffer) SetText(chars string, nChars int32) {
	c_chars := C.CString(chars)
	defer C.free(unsafe.Pointer(c_chars))

	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_buffer_set_text((*C.GtkEntryBuffer)(recv.native), c_chars, c_n_chars)

	return
}

// Returns the value of the ::item-padding property.
/*

C function

gtk_icon_view_get_item_padding
*/
func (recv *IconView) GetItemPadding() int32 {
	retC := C.gtk_icon_view_get_item_padding((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the #GtkIconView:item-padding property which specifies the padding
// around each of the icon view’s items.
/*

C function

gtk_icon_view_set_item_padding
*/
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
var signalInfoBarCloseLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalInfoBarCloseMap[index].callback
	callback()
}

// Unsupported signal 'response' for InfoBar : unsupported parameter response_id : type gint :

// Creates a new #GtkInfoBar object.
/*

C function

gtk_info_bar_new
*/
func InfoBarNew() *InfoBar {
	retC := C.gtk_info_bar_new()
	retGo := InfoBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add an activatable widget to the action area of a #GtkInfoBar,
// connecting a signal handler that will emit the #GtkInfoBar::response
// signal on the message area when the widget is activated. The widget
// is appended to the end of the message areas action area.
/*

C function

gtk_info_bar_add_action_widget
*/
func (recv *InfoBar) AddActionWidget(child *Widget, responseId int32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_add_action_widget((*C.GtkInfoBar)(recv.native), c_child, c_response_id)

	return
}

// Adds a button with the given text and sets things up so that
// clicking the button will emit the “response” signal with the given
// response_id. The button is appended to the end of the info bars's
// action area. The button widget is returned, but usually you don't
// need it.
/*

C function

gtk_info_bar_add_button
*/
func (recv *InfoBar) AddButton(buttonText string, responseId int32) *Button {
	c_button_text := C.CString(buttonText)
	defer C.free(unsafe.Pointer(c_button_text))

	c_response_id := (C.gint)(responseId)

	retC := C.gtk_info_bar_add_button((*C.GtkInfoBar)(recv.native), c_button_text, c_response_id)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// Returns the action area of @info_bar.
/*

C function

gtk_info_bar_get_action_area
*/
func (recv *InfoBar) GetActionArea() *Widget {
	retC := C.gtk_info_bar_get_action_area((*C.GtkInfoBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the content area of @info_bar.
/*

C function

gtk_info_bar_get_content_area
*/
func (recv *InfoBar) GetContentArea() *Widget {
	retC := C.gtk_info_bar_get_content_area((*C.GtkInfoBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the message type of the message area.
/*

C function

gtk_info_bar_get_message_type
*/
func (recv *InfoBar) GetMessageType() MessageType {
	retC := C.gtk_info_bar_get_message_type((*C.GtkInfoBar)(recv.native))
	retGo := (MessageType)(retC)

	return retGo
}

// Emits the “response” signal with the given @response_id.
/*

C function

gtk_info_bar_response
*/
func (recv *InfoBar) Response(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_response((*C.GtkInfoBar)(recv.native), c_response_id)

	return
}

// Sets the last widget in the info bar’s action area with
// the given response_id as the default widget for the dialog.
// Pressing “Enter” normally activates the default widget.
//
// Note that this function currently requires @info_bar to
// be added to a widget hierarchy.
/*

C function

gtk_info_bar_set_default_response
*/
func (recv *InfoBar) SetDefaultResponse(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_info_bar_set_default_response((*C.GtkInfoBar)(recv.native), c_response_id)

	return
}

// Sets the message type of the message area.
//
// GTK+ uses this type to determine how the message is displayed.
/*

C function

gtk_info_bar_set_message_type
*/
func (recv *InfoBar) SetMessageType(messageType MessageType) {
	c_message_type := (C.GtkMessageType)(messageType)

	C.gtk_info_bar_set_message_type((*C.GtkInfoBar)(recv.native), c_message_type)

	return
}

// Calls gtk_widget_set_sensitive (widget, setting) for each
// widget in the info bars’s action area with the given response_id.
// A convenient way to sensitize/desensitize dialog buttons.
/*

C function

gtk_info_bar_set_response_sensitive
*/
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
var signalLabelActivateCurrentLinkLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalLabelActivateCurrentLinkMap[index].callback
	callback()
}

// Unsupported signal 'activate-link' for Label : unsupported parameter uri : type utf8 :

// Returns the URI for the currently active link in the label.
// The active link is the one under the mouse pointer or, in a
// selectable label, the link in which the text cursor is currently
// positioned.
//
// This function is intended for use in a #GtkLabel::activate-link handler
// or for use in a #GtkWidget::query-tooltip handler.
/*

C function

gtk_label_get_current_uri
*/
func (recv *Label) GetCurrentUri() string {
	retC := C.gtk_label_get_current_uri((*C.GtkLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns whether the label is currently keeping track
// of clicked links.
/*

C function

gtk_label_get_track_visited_links
*/
func (recv *Label) GetTrackVisitedLinks() bool {
	retC := C.gtk_label_get_track_visited_links((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the label should keep track of clicked
// links (and use a different color for them).
/*

C function

gtk_label_set_track_visited_links
*/
func (recv *Label) SetTrackVisitedLinks(trackLinks bool) {
	c_track_links :=
		boolToGboolean(trackLinks)

	C.gtk_label_set_track_visited_links((*C.GtkLabel)(recv.native), c_track_links)

	return
}

// Returns whether the menu reserves space for toggles and
// icons, regardless of their actual presence.
/*

C function

gtk_menu_get_reserve_toggle_size
*/
func (recv *Menu) GetReserveToggleSize() bool {
	retC := C.gtk_menu_get_reserve_toggle_size((*C.GtkMenu)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the menu should reserve space for drawing toggles
// or icons, regardless of their actual presence.
/*

C function

gtk_menu_set_reserve_toggle_size
*/
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
var signalPrintOperationUpdateCustomWidgetLock sync.Mutex

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
	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

	setup := PageSetupNewFromC(unsafe.Pointer(c_setup))

	settings := PrintSettingsNewFromC(unsafe.Pointer(c_settings))

	index := int(uintptr(data))
	callback := signalPrintOperationUpdateCustomWidgetMap[index].callback
	callback(widget, setup, settings)
}

// Gets the value of #GtkPrintOperation:embed-page-setup property.
/*

C function

gtk_print_operation_get_embed_page_setup
*/
func (recv *PrintOperation) GetEmbedPageSetup() bool {
	retC := C.gtk_print_operation_get_embed_page_setup((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value of #GtkPrintOperation:has-selection property.
/*

C function

gtk_print_operation_get_has_selection
*/
func (recv *PrintOperation) GetHasSelection() bool {
	retC := C.gtk_print_operation_get_has_selection((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the number of pages that will be printed.
//
// Note that this value is set during print preparation phase
// (%GTK_PRINT_STATUS_PREPARING), so this function should never be
// called before the data generation phase (%GTK_PRINT_STATUS_GENERATING_DATA).
// You can connect to the #GtkPrintOperation::status-changed signal
// and call gtk_print_operation_get_n_pages_to_print() when
// print status is %GTK_PRINT_STATUS_GENERATING_DATA.
// This is typically used to track the progress of print operation.
/*

C function

gtk_print_operation_get_n_pages_to_print
*/
func (recv *PrintOperation) GetNPagesToPrint() int32 {
	retC := C.gtk_print_operation_get_n_pages_to_print((*C.GtkPrintOperation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the value of #GtkPrintOperation:support-selection property.
/*

C function

gtk_print_operation_get_support_selection
*/
func (recv *PrintOperation) GetSupportSelection() bool {
	retC := C.gtk_print_operation_get_support_selection((*C.GtkPrintOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Embed page size combo box and orientation combo box into page setup page.
// Selected page setup is stored as default page setup in #GtkPrintOperation.
/*

C function

gtk_print_operation_set_embed_page_setup
*/
func (recv *PrintOperation) SetEmbedPageSetup(embed bool) {
	c_embed :=
		boolToGboolean(embed)

	C.gtk_print_operation_set_embed_page_setup((*C.GtkPrintOperation)(recv.native), c_embed)

	return
}

// Sets whether there is a selection to print.
//
// Application has to set number of pages to which the selection
// will draw by gtk_print_operation_set_n_pages() in a callback of
// #GtkPrintOperation::begin-print.
/*

C function

gtk_print_operation_set_has_selection
*/
func (recv *PrintOperation) SetHasSelection(hasSelection bool) {
	c_has_selection :=
		boolToGboolean(hasSelection)

	C.gtk_print_operation_set_has_selection((*C.GtkPrintOperation)(recv.native), c_has_selection)

	return
}

// Sets whether selection is supported by #GtkPrintOperation.
/*

C function

gtk_print_operation_set_support_selection
*/
func (recv *PrintOperation) SetSupportSelection(supportSelection bool) {
	c_support_selection :=
		boolToGboolean(supportSelection)

	C.gtk_print_operation_set_support_selection((*C.GtkPrintOperation)(recv.native), c_support_selection)

	return
}

// Gets the value set by gtk_range_set_flippable().
/*

C function

gtk_range_get_flippable
*/
func (recv *Range) GetFlippable() bool {
	retC := C.gtk_range_get_flippable((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// If a range is flippable, it will switch its direction if it is
// horizontal and its direction is %GTK_TEXT_DIR_RTL.
//
// See gtk_widget_get_direction().
/*

C function

gtk_range_set_flippable
*/
func (recv *Range) SetFlippable(flippable bool) {
	c_flippable :=
		boolToGboolean(flippable)

	C.gtk_range_set_flippable((*C.GtkRange)(recv.native), c_flippable)

	return
}

// Gets the title of this tray icon. See gtk_status_icon_set_title().
/*

C function

gtk_status_icon_get_title
*/
func (recv *StatusIcon) GetTitle() string {
	retC := C.gtk_status_icon_get_title((*C.GtkStatusIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the title of this tray icon.
// This should be a short, human-readable, localized string
// describing the tray icon. It may be used by tools like screen
// readers to render the tray icon.
/*

C function

gtk_status_icon_set_title
*/
func (recv *StatusIcon) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_status_icon_set_title((*C.GtkStatusIcon)(recv.native), c_title)

	return
}

// Unsupported : gtk_widget_get_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Determines whether the application intends to draw on the widget in
// an #GtkWidget::draw handler.
//
// See gtk_widget_set_app_paintable()
/*

C function

gtk_widget_get_app_paintable
*/
func (recv *Widget) GetAppPaintable() bool {
	retC := C.gtk_widget_get_app_paintable((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget can be a default widget. See
// gtk_widget_set_can_default().
/*

C function

gtk_widget_get_can_default
*/
func (recv *Widget) GetCanDefault() bool {
	retC := C.gtk_widget_get_can_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget can own the input focus. See
// gtk_widget_set_can_focus().
/*

C function

gtk_widget_get_can_focus
*/
func (recv *Widget) GetCanFocus() bool {
	retC := C.gtk_widget_get_can_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether the widget is double buffered.
//
// See gtk_widget_set_double_buffered()
/*

C function

gtk_widget_get_double_buffered
*/
func (recv *Widget) GetDoubleBuffered() bool {
	retC := C.gtk_widget_get_double_buffered((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget has a #GdkWindow of its own. See
// gtk_widget_set_has_window().
/*

C function

gtk_widget_get_has_window
*/
func (recv *Widget) GetHasWindow() bool {
	retC := C.gtk_widget_get_has_window((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget is always treated as the default widget
// within its toplevel when it has the focus, even if another widget
// is the default.
//
// See gtk_widget_set_receives_default().
/*

C function

gtk_widget_get_receives_default
*/
func (recv *Widget) GetReceivesDefault() bool {
	retC := C.gtk_widget_get_receives_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the widget’s sensitivity (in the sense of returning
// the value that has been set using gtk_widget_set_sensitive()).
//
// The effective sensitivity of a widget is however determined by both its
// own and its parent widget’s sensitivity. See gtk_widget_is_sensitive().
/*

C function

gtk_widget_get_sensitive
*/
func (recv *Widget) GetSensitive() bool {
	retC := C.gtk_widget_get_sensitive((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the widget’s state. See gtk_widget_set_state().
/*

C function

gtk_widget_get_state
*/
func (recv *Widget) GetState() StateType {
	retC := C.gtk_widget_get_state((*C.GtkWidget)(recv.native))
	retGo := (StateType)(retC)

	return retGo
}

// Determines whether the widget is visible. If you want to
// take into account whether the widget’s parent is also marked as
// visible, use gtk_widget_is_visible() instead.
//
// This function does not check if the widget is obscured in any way.
//
// See gtk_widget_set_visible().
/*

C function

gtk_widget_get_visible
*/
func (recv *Widget) GetVisible() bool {
	retC := C.gtk_widget_get_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget is the current default widget within its
// toplevel. See gtk_widget_set_can_default().
/*

C function

gtk_widget_has_default
*/
func (recv *Widget) HasDefault() bool {
	retC := C.gtk_widget_has_default((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines if the widget has the global input focus. See
// gtk_widget_is_focus() for the difference between having the global
// input focus, and only having the focus within a toplevel.
/*

C function

gtk_widget_has_focus
*/
func (recv *Widget) HasFocus() bool {
	retC := C.gtk_widget_has_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether the widget is currently grabbing events, so it
// is the only widget receiving input events (keyboard and mouse).
//
// See also gtk_grab_add().
/*

C function

gtk_widget_has_grab
*/
func (recv *Widget) HasGrab() bool {
	retC := C.gtk_widget_has_grab((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget can be drawn to. A widget can be drawn
// to if it is mapped and visible.
/*

C function

gtk_widget_is_drawable
*/
func (recv *Widget) IsDrawable() bool {
	retC := C.gtk_widget_is_drawable((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the widget’s effective sensitivity, which means
// it is sensitive itself and also its parent widget is sensitive
/*

C function

gtk_widget_is_sensitive
*/
func (recv *Widget) IsSensitive() bool {
	retC := C.gtk_widget_is_sensitive((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether @widget is a toplevel widget.
//
// Currently only #GtkWindow and #GtkInvisible (and out-of-process
// #GtkPlugs) are toplevel widgets. Toplevel widgets have no parent
// widget.
/*

C function

gtk_widget_is_toplevel
*/
func (recv *Widget) IsToplevel() bool {
	retC := C.gtk_widget_is_toplevel((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_set_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Specifies whether @widget can be a default widget. See
// gtk_widget_grab_default() for details about the meaning of
// “default”.
/*

C function

gtk_widget_set_can_default
*/
func (recv *Widget) SetCanDefault(canDefault bool) {
	c_can_default :=
		boolToGboolean(canDefault)

	C.gtk_widget_set_can_default((*C.GtkWidget)(recv.native), c_can_default)

	return
}

// Specifies whether @widget can own the input focus. See
// gtk_widget_grab_focus() for actually setting the input focus on a
// widget.
/*

C function

gtk_widget_set_can_focus
*/
func (recv *Widget) SetCanFocus(canFocus bool) {
	c_can_focus :=
		boolToGboolean(canFocus)

	C.gtk_widget_set_can_focus((*C.GtkWidget)(recv.native), c_can_focus)

	return
}

// Specifies whether @widget has a #GdkWindow of its own. Note that
// all realized widgets have a non-%NULL “window” pointer
// (gtk_widget_get_window() never returns a %NULL window when a widget
// is realized), but for many of them it’s actually the #GdkWindow of
// one of its parent widgets. Widgets that do not create a %window for
// themselves in #GtkWidget::realize must announce this by
// calling this function with @has_window = %FALSE.
//
// This function should only be called by widget implementations,
// and they should call it in their init() function.
/*

C function

gtk_widget_set_has_window
*/
func (recv *Widget) SetHasWindow(hasWindow bool) {
	c_has_window :=
		boolToGboolean(hasWindow)

	C.gtk_widget_set_has_window((*C.GtkWidget)(recv.native), c_has_window)

	return
}

// Specifies whether @widget will be treated as the default widget
// within its toplevel when it has the focus, even if another widget
// is the default.
//
// See gtk_widget_grab_default() for details about the meaning of
// “default”.
/*

C function

gtk_widget_set_receives_default
*/
func (recv *Widget) SetReceivesDefault(receivesDefault bool) {
	c_receives_default :=
		boolToGboolean(receivesDefault)

	C.gtk_widget_set_receives_default((*C.GtkWidget)(recv.native), c_receives_default)

	return
}

// Sets the visibility state of @widget. Note that setting this to
// %TRUE doesn’t mean the widget is actually viewable, see
// gtk_widget_get_visible().
//
// This function simply calls gtk_widget_show() or gtk_widget_hide()
// but is nicer to use when the visibility of the widget depends on
// some condition.
/*

C function

gtk_widget_set_visible
*/
func (recv *Widget) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_widget_set_visible((*C.GtkWidget)(recv.native), c_visible)

	return
}

// Sets a widget’s window. This function should only be used in a
// widget’s #GtkWidget::realize implementation. The %window passed is
// usually either new window created with gdk_window_new(), or the
// window of its parent widget as returned by
// gtk_widget_get_parent_window().
//
// Widgets must indicate whether they will create their own #GdkWindow
// by calling gtk_widget_set_has_window(). This is usually done in the
// widget’s init() function.
//
// Note that this function does not add any reference to @window.
/*

C function

gtk_widget_set_window
*/
func (recv *Widget) SetWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_widget_set_window((*C.GtkWidget)(recv.native), c_window)

	return
}

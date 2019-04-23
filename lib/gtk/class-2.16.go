// This is a generated file - DO NOT EDIT
// +build gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
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

	void entry_iconPressHandler(GObject *, GtkEntryIconPosition, GdkEventButton *, gpointer);

	static gulong Entry_signal_connect_icon_press(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "icon-press", G_CALLBACK(entry_iconPressHandler), data);
	}

*/
/*

	void entry_iconReleaseHandler(GObject *, GtkEntryIconPosition, GdkEventButton *, gpointer);

	static gulong Entry_signal_connect_icon_release(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "icon-release", G_CALLBACK(entry_iconReleaseHandler), data);
	}

*/
/*

	gboolean statusicon_queryTooltipHandler(GObject *, gint, gint, gboolean, GtkTooltip *, gpointer);

	static gulong StatusIcon_signal_connect_query_tooltip(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-tooltip", G_CALLBACK(statusicon_queryTooltipHandler), data);
	}

*/
/*

	gboolean statusicon_scrollEventHandler(GObject *, GdkEventScroll *, gpointer);

	static gulong StatusIcon_signal_connect_scroll_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "scroll-event", G_CALLBACK(statusicon_scrollEventHandler), data);
	}

*/
/*

	void textbuffer_pasteDoneHandler(GObject *, GtkClipboard *, gpointer);

	static gulong TextBuffer_signal_connect_paste_done(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "paste-done", G_CALLBACK(textbuffer_pasteDoneHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_action_block_activate

// Blacklisted : gtk_action_get_gicon

// Blacklisted : gtk_action_get_icon_name

// Blacklisted : gtk_action_get_is_important

// Blacklisted : gtk_action_get_label

// Blacklisted : gtk_action_get_short_label

// Blacklisted : gtk_action_get_stock_id

// Blacklisted : gtk_action_get_tooltip

// Blacklisted : gtk_action_get_visible_horizontal

// Blacklisted : gtk_action_get_visible_vertical

// Blacklisted : gtk_action_set_gicon

// Blacklisted : gtk_action_set_icon_name

// Blacklisted : gtk_action_set_is_important

// Blacklisted : gtk_action_set_label

// Blacklisted : gtk_action_set_short_label

// Blacklisted : gtk_action_set_stock_id

// Blacklisted : gtk_action_set_tooltip

// Blacklisted : gtk_action_set_visible_horizontal

// Blacklisted : gtk_action_set_visible_vertical

// Blacklisted : gtk_action_unblock_activate

// Blacklisted : gtk_cell_view_get_model

type signalEntryIconPressDetail struct {
	callback  EntrySignalIconPressCallback
	handlerID C.gulong
}

var signalEntryIconPressId int
var signalEntryIconPressMap = make(map[int]signalEntryIconPressDetail)
var signalEntryIconPressLock sync.RWMutex

// EntrySignalIconPressCallback is a callback function for a 'icon-press' signal emitted from a Entry.
type EntrySignalIconPressCallback func(iconPos EntryIconPosition, event *gdk.EventButton)

/*
ConnectIconPress connects the callback to the 'icon-press' signal for the Entry.

The returned value represents the connection, and may be passed to DisconnectIconPress to remove it.
*/
func (recv *Entry) ConnectIconPress(callback EntrySignalIconPressCallback) int {
	signalEntryIconPressLock.Lock()
	defer signalEntryIconPressLock.Unlock()

	signalEntryIconPressId++
	instance := C.gpointer(recv.native)
	handlerID := C.Entry_signal_connect_icon_press(instance, C.gpointer(uintptr(signalEntryIconPressId)))

	detail := signalEntryIconPressDetail{callback, handlerID}
	signalEntryIconPressMap[signalEntryIconPressId] = detail

	return signalEntryIconPressId
}

/*
DisconnectIconPress disconnects a callback from the 'icon-press' signal for the Entry.

The connectionID should be a value returned from a call to ConnectIconPress.
*/
func (recv *Entry) DisconnectIconPress(connectionID int) {
	signalEntryIconPressLock.Lock()
	defer signalEntryIconPressLock.Unlock()

	detail, exists := signalEntryIconPressMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryIconPressMap, connectionID)
}

//export entry_iconPressHandler
func entry_iconPressHandler(_ *C.GObject, c_icon_pos C.GtkEntryIconPosition, c_event *C.GdkEventButton, data C.gpointer) {
	signalEntryIconPressLock.RLock()
	defer signalEntryIconPressLock.RUnlock()

	iconPos := EntryIconPosition(c_icon_pos)

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalEntryIconPressMap[index].callback
	callback(iconPos, event)
}

type signalEntryIconReleaseDetail struct {
	callback  EntrySignalIconReleaseCallback
	handlerID C.gulong
}

var signalEntryIconReleaseId int
var signalEntryIconReleaseMap = make(map[int]signalEntryIconReleaseDetail)
var signalEntryIconReleaseLock sync.RWMutex

// EntrySignalIconReleaseCallback is a callback function for a 'icon-release' signal emitted from a Entry.
type EntrySignalIconReleaseCallback func(iconPos EntryIconPosition, event *gdk.EventButton)

/*
ConnectIconRelease connects the callback to the 'icon-release' signal for the Entry.

The returned value represents the connection, and may be passed to DisconnectIconRelease to remove it.
*/
func (recv *Entry) ConnectIconRelease(callback EntrySignalIconReleaseCallback) int {
	signalEntryIconReleaseLock.Lock()
	defer signalEntryIconReleaseLock.Unlock()

	signalEntryIconReleaseId++
	instance := C.gpointer(recv.native)
	handlerID := C.Entry_signal_connect_icon_release(instance, C.gpointer(uintptr(signalEntryIconReleaseId)))

	detail := signalEntryIconReleaseDetail{callback, handlerID}
	signalEntryIconReleaseMap[signalEntryIconReleaseId] = detail

	return signalEntryIconReleaseId
}

/*
DisconnectIconRelease disconnects a callback from the 'icon-release' signal for the Entry.

The connectionID should be a value returned from a call to ConnectIconRelease.
*/
func (recv *Entry) DisconnectIconRelease(connectionID int) {
	signalEntryIconReleaseLock.Lock()
	defer signalEntryIconReleaseLock.Unlock()

	detail, exists := signalEntryIconReleaseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryIconReleaseMap, connectionID)
}

//export entry_iconReleaseHandler
func entry_iconReleaseHandler(_ *C.GObject, c_icon_pos C.GtkEntryIconPosition, c_event *C.GdkEventButton, data C.gpointer) {
	signalEntryIconReleaseLock.RLock()
	defer signalEntryIconReleaseLock.RUnlock()

	iconPos := EntryIconPosition(c_icon_pos)

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalEntryIconReleaseMap[index].callback
	callback(iconPos, event)
}

// Blacklisted : gtk_entry_get_current_icon_drag_source

// Blacklisted : gtk_entry_get_icon_activatable

// Blacklisted : gtk_entry_get_icon_at_pos

// Blacklisted : gtk_entry_get_icon_gicon

// Blacklisted : gtk_entry_get_icon_name

// Blacklisted : gtk_entry_get_icon_pixbuf

// Blacklisted : gtk_entry_get_icon_sensitive

// Blacklisted : gtk_entry_get_icon_stock

// Blacklisted : gtk_entry_get_icon_storage_type

// Blacklisted : gtk_entry_get_icon_tooltip_markup

// Blacklisted : gtk_entry_get_icon_tooltip_text

// Blacklisted : gtk_entry_get_progress_fraction

// Blacklisted : gtk_entry_get_progress_pulse_step

// Blacklisted : gtk_entry_progress_pulse

// Blacklisted : gtk_entry_set_icon_activatable

// Blacklisted : gtk_entry_set_icon_drag_source

// Blacklisted : gtk_entry_set_icon_from_gicon

// Blacklisted : gtk_entry_set_icon_from_icon_name

// Blacklisted : gtk_entry_set_icon_from_pixbuf

// Blacklisted : gtk_entry_set_icon_from_stock

// Blacklisted : gtk_entry_set_icon_sensitive

// Blacklisted : gtk_entry_set_icon_tooltip_markup

// Blacklisted : gtk_entry_set_icon_tooltip_text

// Blacklisted : gtk_entry_set_progress_fraction

// Blacklisted : gtk_entry_set_progress_pulse_step

// Blacklisted : gtk_entry_unset_invisible_char

// Blacklisted : gtk_im_multicontext_get_context_id

// Blacklisted : gtk_im_multicontext_set_context_id

// Blacklisted : gtk_image_menu_item_get_always_show_image

// Blacklisted : gtk_image_menu_item_get_use_stock

// Blacklisted : gtk_image_menu_item_set_accel_group

// Blacklisted : gtk_image_menu_item_set_always_show_image

// Blacklisted : gtk_image_menu_item_set_use_stock

// Blacklisted : gtk_menu_item_get_label

// Blacklisted : gtk_menu_item_get_use_underline

// Blacklisted : gtk_menu_item_set_label

// Blacklisted : gtk_menu_item_set_use_underline

// Blacklisted : gtk_print_operation_draw_page_finish

// Blacklisted : gtk_print_operation_set_defer_drawing

// Blacklisted : gtk_print_settings_get_printer_lpi

// Blacklisted : gtk_print_settings_get_resolution_x

// Blacklisted : gtk_print_settings_get_resolution_y

// Blacklisted : gtk_print_settings_set_printer_lpi

// Blacklisted : gtk_print_settings_set_resolution_xy

// Blacklisted : gtk_scale_add_mark

// Blacklisted : gtk_scale_clear_marks

type signalStatusIconQueryTooltipDetail struct {
	callback  StatusIconSignalQueryTooltipCallback
	handlerID C.gulong
}

var signalStatusIconQueryTooltipId int
var signalStatusIconQueryTooltipMap = make(map[int]signalStatusIconQueryTooltipDetail)
var signalStatusIconQueryTooltipLock sync.RWMutex

// StatusIconSignalQueryTooltipCallback is a callback function for a 'query-tooltip' signal emitted from a StatusIcon.
type StatusIconSignalQueryTooltipCallback func(x int32, y int32, keyboardMode bool, tooltip *Tooltip) bool

/*
ConnectQueryTooltip connects the callback to the 'query-tooltip' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectQueryTooltip to remove it.
*/
func (recv *StatusIcon) ConnectQueryTooltip(callback StatusIconSignalQueryTooltipCallback) int {
	signalStatusIconQueryTooltipLock.Lock()
	defer signalStatusIconQueryTooltipLock.Unlock()

	signalStatusIconQueryTooltipId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_query_tooltip(instance, C.gpointer(uintptr(signalStatusIconQueryTooltipId)))

	detail := signalStatusIconQueryTooltipDetail{callback, handlerID}
	signalStatusIconQueryTooltipMap[signalStatusIconQueryTooltipId] = detail

	return signalStatusIconQueryTooltipId
}

/*
DisconnectQueryTooltip disconnects a callback from the 'query-tooltip' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectQueryTooltip.
*/
func (recv *StatusIcon) DisconnectQueryTooltip(connectionID int) {
	signalStatusIconQueryTooltipLock.Lock()
	defer signalStatusIconQueryTooltipLock.Unlock()

	detail, exists := signalStatusIconQueryTooltipMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconQueryTooltipMap, connectionID)
}

//export statusicon_queryTooltipHandler
func statusicon_queryTooltipHandler(_ *C.GObject, c_x C.gint, c_y C.gint, c_keyboard_mode C.gboolean, c_tooltip *C.GtkTooltip, data C.gpointer) C.gboolean {
	signalStatusIconQueryTooltipLock.RLock()
	defer signalStatusIconQueryTooltipLock.RUnlock()

	x := int32(c_x)

	y := int32(c_y)

	keyboardMode := c_keyboard_mode == C.TRUE

	tooltip := TooltipNewFromC(unsafe.Pointer(c_tooltip))

	index := int(uintptr(data))
	callback := signalStatusIconQueryTooltipMap[index].callback
	retGo := callback(x, y, keyboardMode, tooltip)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalStatusIconScrollEventDetail struct {
	callback  StatusIconSignalScrollEventCallback
	handlerID C.gulong
}

var signalStatusIconScrollEventId int
var signalStatusIconScrollEventMap = make(map[int]signalStatusIconScrollEventDetail)
var signalStatusIconScrollEventLock sync.RWMutex

// StatusIconSignalScrollEventCallback is a callback function for a 'scroll-event' signal emitted from a StatusIcon.
type StatusIconSignalScrollEventCallback func(event *gdk.EventScroll) bool

/*
ConnectScrollEvent connects the callback to the 'scroll-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectScrollEvent to remove it.
*/
func (recv *StatusIcon) ConnectScrollEvent(callback StatusIconSignalScrollEventCallback) int {
	signalStatusIconScrollEventLock.Lock()
	defer signalStatusIconScrollEventLock.Unlock()

	signalStatusIconScrollEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_scroll_event(instance, C.gpointer(uintptr(signalStatusIconScrollEventId)))

	detail := signalStatusIconScrollEventDetail{callback, handlerID}
	signalStatusIconScrollEventMap[signalStatusIconScrollEventId] = detail

	return signalStatusIconScrollEventId
}

/*
DisconnectScrollEvent disconnects a callback from the 'scroll-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectScrollEvent.
*/
func (recv *StatusIcon) DisconnectScrollEvent(connectionID int) {
	signalStatusIconScrollEventLock.Lock()
	defer signalStatusIconScrollEventLock.Unlock()

	detail, exists := signalStatusIconScrollEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconScrollEventMap, connectionID)
}

//export statusicon_scrollEventHandler
func statusicon_scrollEventHandler(_ *C.GObject, c_event *C.GdkEventScroll, data C.gpointer) C.gboolean {
	signalStatusIconScrollEventLock.RLock()
	defer signalStatusIconScrollEventLock.RUnlock()

	event := gdk.EventScrollNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconScrollEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_status_icon_get_has_tooltip

// Blacklisted : gtk_status_icon_get_tooltip_markup

// Blacklisted : gtk_status_icon_get_tooltip_text

// Blacklisted : gtk_status_icon_set_has_tooltip

// Blacklisted : gtk_status_icon_set_tooltip_markup

// Blacklisted : gtk_status_icon_set_tooltip_text

// Unsupported : gtk_style_get : unsupported parameter ... : varargs

// Blacklisted : gtk_style_get_style_property

// Unsupported : gtk_style_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

type signalTextBufferPasteDoneDetail struct {
	callback  TextBufferSignalPasteDoneCallback
	handlerID C.gulong
}

var signalTextBufferPasteDoneId int
var signalTextBufferPasteDoneMap = make(map[int]signalTextBufferPasteDoneDetail)
var signalTextBufferPasteDoneLock sync.RWMutex

// TextBufferSignalPasteDoneCallback is a callback function for a 'paste-done' signal emitted from a TextBuffer.
type TextBufferSignalPasteDoneCallback func(clipboard *Clipboard)

/*
ConnectPasteDone connects the callback to the 'paste-done' signal for the TextBuffer.

The returned value represents the connection, and may be passed to DisconnectPasteDone to remove it.
*/
func (recv *TextBuffer) ConnectPasteDone(callback TextBufferSignalPasteDoneCallback) int {
	signalTextBufferPasteDoneLock.Lock()
	defer signalTextBufferPasteDoneLock.Unlock()

	signalTextBufferPasteDoneId++
	instance := C.gpointer(recv.native)
	handlerID := C.TextBuffer_signal_connect_paste_done(instance, C.gpointer(uintptr(signalTextBufferPasteDoneId)))

	detail := signalTextBufferPasteDoneDetail{callback, handlerID}
	signalTextBufferPasteDoneMap[signalTextBufferPasteDoneId] = detail

	return signalTextBufferPasteDoneId
}

/*
DisconnectPasteDone disconnects a callback from the 'paste-done' signal for the TextBuffer.

The connectionID should be a value returned from a call to ConnectPasteDone.
*/
func (recv *TextBuffer) DisconnectPasteDone(connectionID int) {
	signalTextBufferPasteDoneLock.Lock()
	defer signalTextBufferPasteDoneLock.Unlock()

	detail, exists := signalTextBufferPasteDoneMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextBufferPasteDoneMap, connectionID)
}

//export textbuffer_pasteDoneHandler
func textbuffer_pasteDoneHandler(_ *C.GObject, c_clipboard *C.GtkClipboard, data C.gpointer) {
	signalTextBufferPasteDoneLock.RLock()
	defer signalTextBufferPasteDoneLock.RUnlock()

	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	index := int(uintptr(data))
	callback := signalTextBufferPasteDoneMap[index].callback
	callback(clipboard)
}

// Blacklisted : gtk_window_get_default_icon_name

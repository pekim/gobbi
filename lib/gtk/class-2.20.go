// This is a generated file - DO NOT EDIT
// +build gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	void entry_preeditChangedHandler(GObject *, gchar*, gpointer);

	static gulong Entry_signal_connect_preedit_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "preedit-changed", G_CALLBACK(entry_preeditChangedHandler), data);
	}

*/
/*

	void textview_preeditChangedHandler(GObject *, gchar*, gpointer);

	static gulong TextView_signal_connect_preedit_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "preedit-changed", G_CALLBACK(textview_preeditChangedHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_action_get_always_show_image

// Blacklisted : gtk_action_set_always_show_image

// Blacklisted : gtk_cell_renderer_spinner_new

// Blacklisted : gtk_dialog_get_widget_for_response

type signalEntryPreeditChangedDetail struct {
	callback  EntrySignalPreeditChangedCallback
	handlerID C.gulong
}

var signalEntryPreeditChangedId int
var signalEntryPreeditChangedMap = make(map[int]signalEntryPreeditChangedDetail)
var signalEntryPreeditChangedLock sync.RWMutex

// EntrySignalPreeditChangedCallback is a callback function for a 'preedit-changed' signal emitted from a Entry.
type EntrySignalPreeditChangedCallback func(preedit string)

/*
ConnectPreeditChanged connects the callback to the 'preedit-changed' signal for the Entry.

The returned value represents the connection, and may be passed to DisconnectPreeditChanged to remove it.
*/
func (recv *Entry) ConnectPreeditChanged(callback EntrySignalPreeditChangedCallback) int {
	signalEntryPreeditChangedLock.Lock()
	defer signalEntryPreeditChangedLock.Unlock()

	signalEntryPreeditChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Entry_signal_connect_preedit_changed(instance, C.gpointer(uintptr(signalEntryPreeditChangedId)))

	detail := signalEntryPreeditChangedDetail{callback, handlerID}
	signalEntryPreeditChangedMap[signalEntryPreeditChangedId] = detail

	return signalEntryPreeditChangedId
}

/*
DisconnectPreeditChanged disconnects a callback from the 'preedit-changed' signal for the Entry.

The connectionID should be a value returned from a call to ConnectPreeditChanged.
*/
func (recv *Entry) DisconnectPreeditChanged(connectionID int) {
	signalEntryPreeditChangedLock.Lock()
	defer signalEntryPreeditChangedLock.Unlock()

	detail, exists := signalEntryPreeditChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryPreeditChangedMap, connectionID)
}

//export entry_preeditChangedHandler
func entry_preeditChangedHandler(_ *C.GObject, c_preedit *C.gchar, data C.gpointer) {
	signalEntryPreeditChangedLock.RLock()
	defer signalEntryPreeditChangedLock.RUnlock()

	preedit := C.GoString(c_preedit)

	index := int(uintptr(data))
	callback := signalEntryPreeditChangedMap[index].callback
	callback(preedit)
}

// Blacklisted : gtk_notebook_get_action_widget

// Blacklisted : gtk_notebook_set_action_widget

// Blacklisted : gtk_offscreen_window_new

// Blacklisted : gtk_offscreen_window_get_pixbuf

// Blacklisted : gtk_offscreen_window_get_surface

// Blacklisted : gtk_paned_get_handle_window

// Blacklisted : gtk_print_context_get_hard_margins

// Blacklisted : gtk_range_get_min_slider_size

// Blacklisted : gtk_range_get_range_rect

// Blacklisted : gtk_range_get_slider_range

// Blacklisted : gtk_range_get_slider_size_fixed

// Blacklisted : gtk_range_set_min_slider_size

// Blacklisted : gtk_range_set_slider_size_fixed

// Blacklisted : gtk_spinner_new

// Blacklisted : gtk_spinner_start

// Blacklisted : gtk_spinner_stop

// Blacklisted : gtk_status_icon_set_name

// Blacklisted : gtk_statusbar_get_message_area

type signalTextViewPreeditChangedDetail struct {
	callback  TextViewSignalPreeditChangedCallback
	handlerID C.gulong
}

var signalTextViewPreeditChangedId int
var signalTextViewPreeditChangedMap = make(map[int]signalTextViewPreeditChangedDetail)
var signalTextViewPreeditChangedLock sync.RWMutex

// TextViewSignalPreeditChangedCallback is a callback function for a 'preedit-changed' signal emitted from a TextView.
type TextViewSignalPreeditChangedCallback func(preedit string)

/*
ConnectPreeditChanged connects the callback to the 'preedit-changed' signal for the TextView.

The returned value represents the connection, and may be passed to DisconnectPreeditChanged to remove it.
*/
func (recv *TextView) ConnectPreeditChanged(callback TextViewSignalPreeditChangedCallback) int {
	signalTextViewPreeditChangedLock.Lock()
	defer signalTextViewPreeditChangedLock.Unlock()

	signalTextViewPreeditChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.TextView_signal_connect_preedit_changed(instance, C.gpointer(uintptr(signalTextViewPreeditChangedId)))

	detail := signalTextViewPreeditChangedDetail{callback, handlerID}
	signalTextViewPreeditChangedMap[signalTextViewPreeditChangedId] = detail

	return signalTextViewPreeditChangedId
}

/*
DisconnectPreeditChanged disconnects a callback from the 'preedit-changed' signal for the TextView.

The connectionID should be a value returned from a call to ConnectPreeditChanged.
*/
func (recv *TextView) DisconnectPreeditChanged(connectionID int) {
	signalTextViewPreeditChangedLock.Lock()
	defer signalTextViewPreeditChangedLock.Unlock()

	detail, exists := signalTextViewPreeditChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTextViewPreeditChangedMap, connectionID)
}

//export textview_preeditChangedHandler
func textview_preeditChangedHandler(_ *C.GObject, c_preedit *C.gchar, data C.gpointer) {
	signalTextViewPreeditChangedLock.RLock()
	defer signalTextViewPreeditChangedLock.RUnlock()

	preedit := C.GoString(c_preedit)

	index := int(uintptr(data))
	callback := signalTextViewPreeditChangedMap[index].callback
	callback(preedit)
}

// Blacklisted : gtk_tool_item_get_ellipsize_mode

// Blacklisted : gtk_tool_item_get_text_alignment

// Blacklisted : gtk_tool_item_get_text_orientation

// Blacklisted : gtk_tool_item_get_text_size_group

// Blacklisted : gtk_tool_item_group_new

// Blacklisted : gtk_tool_item_group_get_collapsed

// Blacklisted : gtk_tool_item_group_get_drop_item

// Blacklisted : gtk_tool_item_group_get_ellipsize

// Blacklisted : gtk_tool_item_group_get_header_relief

// Blacklisted : gtk_tool_item_group_get_item_position

// Blacklisted : gtk_tool_item_group_get_label

// Blacklisted : gtk_tool_item_group_get_label_widget

// Blacklisted : gtk_tool_item_group_get_n_items

// Blacklisted : gtk_tool_item_group_get_nth_item

// Blacklisted : gtk_tool_item_group_insert

// Blacklisted : gtk_tool_item_group_set_collapsed

// Blacklisted : gtk_tool_item_group_set_ellipsize

// Blacklisted : gtk_tool_item_group_set_header_relief

// Blacklisted : gtk_tool_item_group_set_item_position

// Blacklisted : gtk_tool_item_group_set_label

// Blacklisted : gtk_tool_item_group_set_label_widget

// Blacklisted : gtk_tool_palette_new

// Blacklisted : gtk_tool_palette_get_drag_target_group

// Blacklisted : gtk_tool_palette_get_drag_target_item

// Blacklisted : gtk_tool_palette_add_drag_dest

// Blacklisted : gtk_tool_palette_get_drag_item

// Blacklisted : gtk_tool_palette_get_drop_group

// Blacklisted : gtk_tool_palette_get_drop_item

// Blacklisted : gtk_tool_palette_get_exclusive

// Blacklisted : gtk_tool_palette_get_expand

// Blacklisted : gtk_tool_palette_get_group_position

// Blacklisted : gtk_tool_palette_get_hadjustment

// Blacklisted : gtk_tool_palette_get_icon_size

// Blacklisted : gtk_tool_palette_get_style

// Blacklisted : gtk_tool_palette_get_vadjustment

// Blacklisted : gtk_tool_palette_set_drag_source

// Blacklisted : gtk_tool_palette_set_exclusive

// Blacklisted : gtk_tool_palette_set_expand

// Blacklisted : gtk_tool_palette_set_group_position

// Blacklisted : gtk_tool_palette_set_icon_size

// Blacklisted : gtk_tool_palette_set_style

// Blacklisted : gtk_tool_palette_unset_icon_size

// Blacklisted : gtk_tool_palette_unset_style

// Blacklisted : gtk_tooltip_set_icon_from_gicon

// Blacklisted : gtk_viewport_get_bin_window

// Blacklisted : gtk_widget_get_mapped

// Blacklisted : gtk_widget_get_realized

// Blacklisted : gtk_widget_get_requisition

// Blacklisted : gtk_widget_has_rc_style

// Unsupported : gtk_widget_send_focus_change : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_widget_set_mapped

// Blacklisted : gtk_widget_set_realized

// Blacklisted : gtk_widget_style_attach

// Blacklisted : gtk_window_get_mnemonics_visible

// Blacklisted : gtk_window_get_window_type

// Blacklisted : gtk_window_set_mnemonics_visible

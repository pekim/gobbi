// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	void cellrenderercombo_changedHandler(GObject *, gchar*, GtkTreeIter *, gpointer);

	static gulong CellRendererCombo_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(cellrenderercombo_changedHandler), data);
	}

*/
/*

	gboolean statusicon_buttonPressEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong StatusIcon_signal_connect_button_press_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-press-event", G_CALLBACK(statusicon_buttonPressEventHandler), data);
	}

*/
/*

	gboolean statusicon_buttonReleaseEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong StatusIcon_signal_connect_button_release_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-release-event", G_CALLBACK(statusicon_buttonReleaseEventHandler), data);
	}

*/
/*

	gboolean widget_damageEventHandler(GObject *, GdkEventExpose *, gpointer);

	static gulong Widget_signal_connect_damage_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "damage-event", G_CALLBACK(widget_damageEventHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_accel_group_get_is_locked

// Blacklisted : gtk_accel_group_get_modifier_mask

// Blacklisted : gtk_adjustment_configure

// Blacklisted : gtk_adjustment_get_lower

// Blacklisted : gtk_adjustment_get_page_increment

// Blacklisted : gtk_adjustment_get_page_size

// Blacklisted : gtk_adjustment_get_step_increment

// Blacklisted : gtk_adjustment_get_upper

// Blacklisted : gtk_adjustment_set_lower

// Blacklisted : gtk_adjustment_set_page_increment

// Blacklisted : gtk_adjustment_set_page_size

// Blacklisted : gtk_adjustment_set_step_increment

// Blacklisted : gtk_adjustment_set_upper

// Blacklisted : gtk_builder_add_objects_from_file

// Blacklisted : gtk_builder_add_objects_from_string

// Blacklisted : gtk_calendar_get_detail_height_rows

// Blacklisted : gtk_calendar_get_detail_width_chars

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc (GtkCalendarDetailFunc) for param func

// Blacklisted : gtk_calendar_set_detail_height_rows

// Blacklisted : gtk_calendar_set_detail_width_chars

type signalCellRendererComboChangedDetail struct {
	callback  CellRendererComboSignalChangedCallback
	handlerID C.gulong
}

var signalCellRendererComboChangedId int
var signalCellRendererComboChangedMap = make(map[int]signalCellRendererComboChangedDetail)
var signalCellRendererComboChangedLock sync.RWMutex

// CellRendererComboSignalChangedCallback is a callback function for a 'changed' signal emitted from a CellRendererCombo.
type CellRendererComboSignalChangedCallback func(pathString string, newIter *TreeIter)

/*
ConnectChanged connects the callback to the 'changed' signal for the CellRendererCombo.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *CellRendererCombo) ConnectChanged(callback CellRendererComboSignalChangedCallback) int {
	signalCellRendererComboChangedLock.Lock()
	defer signalCellRendererComboChangedLock.Unlock()

	signalCellRendererComboChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRendererCombo_signal_connect_changed(instance, C.gpointer(uintptr(signalCellRendererComboChangedId)))

	detail := signalCellRendererComboChangedDetail{callback, handlerID}
	signalCellRendererComboChangedMap[signalCellRendererComboChangedId] = detail

	return signalCellRendererComboChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the CellRendererCombo.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *CellRendererCombo) DisconnectChanged(connectionID int) {
	signalCellRendererComboChangedLock.Lock()
	defer signalCellRendererComboChangedLock.Unlock()

	detail, exists := signalCellRendererComboChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererComboChangedMap, connectionID)
}

//export cellrenderercombo_changedHandler
func cellrenderercombo_changedHandler(_ *C.GObject, c_path_string *C.gchar, c_new_iter *C.GtkTreeIter, data C.gpointer) {
	signalCellRendererComboChangedLock.RLock()
	defer signalCellRendererComboChangedLock.RUnlock()

	pathString := C.GoString(c_path_string)

	newIter := TreeIterNewFromC(unsafe.Pointer(c_new_iter))

	index := int(uintptr(data))
	callback := signalCellRendererComboChangedMap[index].callback
	callback(pathString, newIter)
}

// Unsupported : gtk_clipboard_request_uris : unsupported parameter callback : no type generator for ClipboardURIReceivedFunc (GtkClipboardURIReceivedFunc) for param callback

// Blacklisted : gtk_clipboard_wait_for_uris

// Blacklisted : gtk_clipboard_wait_is_uris_available

// Blacklisted : gtk_color_selection_dialog_get_color_selection

// Blacklisted : gtk_combo_box_get_button_sensitivity

// Blacklisted : gtk_combo_box_set_button_sensitivity

// Blacklisted : gtk_container_get_focus_child

// Blacklisted : gtk_dialog_get_action_area

// Blacklisted : gtk_dialog_get_content_area

// Blacklisted : gtk_entry_get_overwrite_mode

// Blacklisted : gtk_entry_get_text_length

// Blacklisted : gtk_entry_set_overwrite_mode

// Blacklisted : gtk_font_selection_get_face

// Blacklisted : gtk_font_selection_get_face_list

// Blacklisted : gtk_font_selection_get_family

// Blacklisted : gtk_font_selection_get_family_list

// Blacklisted : gtk_font_selection_get_preview_entry

// Blacklisted : gtk_font_selection_get_size

// Blacklisted : gtk_font_selection_get_size_entry

// Blacklisted : gtk_font_selection_get_size_list

// Blacklisted : gtk_font_selection_dialog_get_cancel_button

// Blacklisted : gtk_font_selection_dialog_get_ok_button

// Blacklisted : gtk_hsv_new

// Blacklisted : gtk_hsv_to_rgb

// Blacklisted : gtk_hsv_get_color

// Blacklisted : gtk_hsv_get_metrics

// Blacklisted : gtk_hsv_is_adjusting

// Blacklisted : gtk_hsv_set_color

// Blacklisted : gtk_hsv_set_metrics

// Blacklisted : gtk_handle_box_get_child_detached

// Blacklisted : gtk_icon_info_new_for_pixbuf

// Blacklisted : gtk_icon_theme_lookup_by_gicon

// Blacklisted : gtk_image_new_from_gicon

// Blacklisted : gtk_image_get_gicon

// Blacklisted : gtk_image_set_from_gicon

// Blacklisted : gtk_layout_get_bin_window

// Blacklisted : gtk_link_button_get_visited

// Blacklisted : gtk_link_button_set_visited

// Blacklisted : gtk_menu_get_accel_path

// Blacklisted : gtk_menu_get_monitor

// Blacklisted : gtk_menu_item_get_accel_path

// Blacklisted : gtk_message_dialog_get_image

// Blacklisted : gtk_mount_operation_new

// Blacklisted : gtk_mount_operation_get_parent

// Blacklisted : gtk_mount_operation_get_screen

// Blacklisted : gtk_mount_operation_is_showing

// Blacklisted : gtk_mount_operation_set_parent

// Blacklisted : gtk_mount_operation_set_screen

// Blacklisted : gtk_page_setup_load_file

// Blacklisted : gtk_page_setup_load_key_file

// Blacklisted : gtk_print_settings_get_number_up_layout

// Blacklisted : gtk_print_settings_load_file

// Blacklisted : gtk_print_settings_load_key_file

// Blacklisted : gtk_print_settings_set_number_up_layout

// Blacklisted : gtk_scale_button_get_minus_button

// Blacklisted : gtk_scale_button_get_plus_button

// Blacklisted : gtk_scale_button_get_popup

type signalStatusIconButtonPressEventDetail struct {
	callback  StatusIconSignalButtonPressEventCallback
	handlerID C.gulong
}

var signalStatusIconButtonPressEventId int
var signalStatusIconButtonPressEventMap = make(map[int]signalStatusIconButtonPressEventDetail)
var signalStatusIconButtonPressEventLock sync.RWMutex

// StatusIconSignalButtonPressEventCallback is a callback function for a 'button-press-event' signal emitted from a StatusIcon.
type StatusIconSignalButtonPressEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonPressEvent connects the callback to the 'button-press-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectButtonPressEvent to remove it.
*/
func (recv *StatusIcon) ConnectButtonPressEvent(callback StatusIconSignalButtonPressEventCallback) int {
	signalStatusIconButtonPressEventLock.Lock()
	defer signalStatusIconButtonPressEventLock.Unlock()

	signalStatusIconButtonPressEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_button_press_event(instance, C.gpointer(uintptr(signalStatusIconButtonPressEventId)))

	detail := signalStatusIconButtonPressEventDetail{callback, handlerID}
	signalStatusIconButtonPressEventMap[signalStatusIconButtonPressEventId] = detail

	return signalStatusIconButtonPressEventId
}

/*
DisconnectButtonPressEvent disconnects a callback from the 'button-press-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectButtonPressEvent.
*/
func (recv *StatusIcon) DisconnectButtonPressEvent(connectionID int) {
	signalStatusIconButtonPressEventLock.Lock()
	defer signalStatusIconButtonPressEventLock.Unlock()

	detail, exists := signalStatusIconButtonPressEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconButtonPressEventMap, connectionID)
}

//export statusicon_buttonPressEventHandler
func statusicon_buttonPressEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	signalStatusIconButtonPressEventLock.RLock()
	defer signalStatusIconButtonPressEventLock.RUnlock()

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconButtonPressEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalStatusIconButtonReleaseEventDetail struct {
	callback  StatusIconSignalButtonReleaseEventCallback
	handlerID C.gulong
}

var signalStatusIconButtonReleaseEventId int
var signalStatusIconButtonReleaseEventMap = make(map[int]signalStatusIconButtonReleaseEventDetail)
var signalStatusIconButtonReleaseEventLock sync.RWMutex

// StatusIconSignalButtonReleaseEventCallback is a callback function for a 'button-release-event' signal emitted from a StatusIcon.
type StatusIconSignalButtonReleaseEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonReleaseEvent connects the callback to the 'button-release-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectButtonReleaseEvent to remove it.
*/
func (recv *StatusIcon) ConnectButtonReleaseEvent(callback StatusIconSignalButtonReleaseEventCallback) int {
	signalStatusIconButtonReleaseEventLock.Lock()
	defer signalStatusIconButtonReleaseEventLock.Unlock()

	signalStatusIconButtonReleaseEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_button_release_event(instance, C.gpointer(uintptr(signalStatusIconButtonReleaseEventId)))

	detail := signalStatusIconButtonReleaseEventDetail{callback, handlerID}
	signalStatusIconButtonReleaseEventMap[signalStatusIconButtonReleaseEventId] = detail

	return signalStatusIconButtonReleaseEventId
}

/*
DisconnectButtonReleaseEvent disconnects a callback from the 'button-release-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectButtonReleaseEvent.
*/
func (recv *StatusIcon) DisconnectButtonReleaseEvent(connectionID int) {
	signalStatusIconButtonReleaseEventLock.Lock()
	defer signalStatusIconButtonReleaseEventLock.Unlock()

	detail, exists := signalStatusIconButtonReleaseEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconButtonReleaseEventMap, connectionID)
}

//export statusicon_buttonReleaseEventHandler
func statusicon_buttonReleaseEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	signalStatusIconButtonReleaseEventLock.RLock()
	defer signalStatusIconButtonReleaseEventLock.RUnlock()

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconButtonReleaseEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_status_icon_new_from_gicon

// Blacklisted : gtk_status_icon_get_gicon

// Blacklisted : gtk_status_icon_get_x11_window_id

// Blacklisted : gtk_status_icon_set_from_gicon

// Blacklisted : gtk_tool_item_toolbar_reconfigured

// Blacklisted : gtk_tooltip_set_icon_from_icon_name

// Unsupported : gtk_tree_selection_get_select_function : no return generator

type signalWidgetDamageEventDetail struct {
	callback  WidgetSignalDamageEventCallback
	handlerID C.gulong
}

var signalWidgetDamageEventId int
var signalWidgetDamageEventMap = make(map[int]signalWidgetDamageEventDetail)
var signalWidgetDamageEventLock sync.RWMutex

// WidgetSignalDamageEventCallback is a callback function for a 'damage-event' signal emitted from a Widget.
type WidgetSignalDamageEventCallback func(event *gdk.EventExpose) bool

/*
ConnectDamageEvent connects the callback to the 'damage-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDamageEvent to remove it.
*/
func (recv *Widget) ConnectDamageEvent(callback WidgetSignalDamageEventCallback) int {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	signalWidgetDamageEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_damage_event(instance, C.gpointer(uintptr(signalWidgetDamageEventId)))

	detail := signalWidgetDamageEventDetail{callback, handlerID}
	signalWidgetDamageEventMap[signalWidgetDamageEventId] = detail

	return signalWidgetDamageEventId
}

/*
DisconnectDamageEvent disconnects a callback from the 'damage-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDamageEvent.
*/
func (recv *Widget) DisconnectDamageEvent(connectionID int) {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	detail, exists := signalWidgetDamageEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDamageEventMap, connectionID)
}

//export widget_damageEventHandler
func widget_damageEventHandler(_ *C.GObject, c_event *C.GdkEventExpose, data C.gpointer) C.gboolean {
	signalWidgetDamageEventLock.RLock()
	defer signalWidgetDamageEventLock.RUnlock()

	event := gdk.EventExposeNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetDamageEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_widget_get_window

// Blacklisted : gtk_window_get_default_widget

// Blacklisted : gtk_window_group_list_windows

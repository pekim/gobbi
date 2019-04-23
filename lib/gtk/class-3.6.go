// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

	void levelbar_offsetChangedHandler(GObject *, gchar*, gpointer);

	static gulong LevelBar_signal_connect_offset_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "offset-changed", G_CALLBACK(levelbar_offsetChangedHandler), data);
	}

*/
import "C"

// Blacklisted : gtk_accel_label_set_accel

// Blacklisted : gtk_action_group_get_accel_group

// Blacklisted : gtk_action_group_set_accel_group

// Blacklisted : gtk_application_get_active_window

// Blacklisted : gtk_application_get_window_by_id

// Blacklisted : gtk_application_window_get_id

// Blacklisted : gtk_button_get_always_show_image

// Blacklisted : gtk_button_set_always_show_image

// Blacklisted : gtk_entry_get_attributes

// Blacklisted : gtk_entry_get_input_hints

// Blacklisted : gtk_entry_get_input_purpose

// Blacklisted : gtk_entry_set_attributes

// Blacklisted : gtk_entry_set_input_hints

// Blacklisted : gtk_entry_set_input_purpose

// Blacklisted : gtk_icon_view_get_cell_rect

type signalLevelBarOffsetChangedDetail struct {
	callback  LevelBarSignalOffsetChangedCallback
	handlerID C.gulong
}

var signalLevelBarOffsetChangedId int
var signalLevelBarOffsetChangedMap = make(map[int]signalLevelBarOffsetChangedDetail)
var signalLevelBarOffsetChangedLock sync.RWMutex

// LevelBarSignalOffsetChangedCallback is a callback function for a 'offset-changed' signal emitted from a LevelBar.
type LevelBarSignalOffsetChangedCallback func(name string)

/*
ConnectOffsetChanged connects the callback to the 'offset-changed' signal for the LevelBar.

The returned value represents the connection, and may be passed to DisconnectOffsetChanged to remove it.
*/
func (recv *LevelBar) ConnectOffsetChanged(callback LevelBarSignalOffsetChangedCallback) int {
	signalLevelBarOffsetChangedLock.Lock()
	defer signalLevelBarOffsetChangedLock.Unlock()

	signalLevelBarOffsetChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.LevelBar_signal_connect_offset_changed(instance, C.gpointer(uintptr(signalLevelBarOffsetChangedId)))

	detail := signalLevelBarOffsetChangedDetail{callback, handlerID}
	signalLevelBarOffsetChangedMap[signalLevelBarOffsetChangedId] = detail

	return signalLevelBarOffsetChangedId
}

/*
DisconnectOffsetChanged disconnects a callback from the 'offset-changed' signal for the LevelBar.

The connectionID should be a value returned from a call to ConnectOffsetChanged.
*/
func (recv *LevelBar) DisconnectOffsetChanged(connectionID int) {
	signalLevelBarOffsetChangedLock.Lock()
	defer signalLevelBarOffsetChangedLock.Unlock()

	detail, exists := signalLevelBarOffsetChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalLevelBarOffsetChangedMap, connectionID)
}

//export levelbar_offsetChangedHandler
func levelbar_offsetChangedHandler(_ *C.GObject, c_name *C.gchar, data C.gpointer) {
	signalLevelBarOffsetChangedLock.RLock()
	defer signalLevelBarOffsetChangedLock.RUnlock()

	name := C.GoString(c_name)

	index := int(uintptr(data))
	callback := signalLevelBarOffsetChangedMap[index].callback
	callback(name)
}

// Blacklisted : gtk_level_bar_new

// Blacklisted : gtk_level_bar_new_for_interval

// Blacklisted : gtk_level_bar_add_offset_value

// Blacklisted : gtk_level_bar_get_max_value

// Blacklisted : gtk_level_bar_get_min_value

// Blacklisted : gtk_level_bar_get_mode

// Blacklisted : gtk_level_bar_get_offset_value

// Blacklisted : gtk_level_bar_get_value

// Blacklisted : gtk_level_bar_remove_offset_value

// Blacklisted : gtk_level_bar_set_max_value

// Blacklisted : gtk_level_bar_set_min_value

// Blacklisted : gtk_level_bar_set_mode

// Blacklisted : gtk_level_bar_set_value

// Blacklisted : gtk_menu_button_new

// Blacklisted : gtk_menu_button_get_align_widget

// Blacklisted : gtk_menu_button_get_direction

// Blacklisted : gtk_menu_button_get_menu_model

// Blacklisted : gtk_menu_button_get_popup

// Blacklisted : gtk_menu_button_set_align_widget

// Blacklisted : gtk_menu_button_set_direction

// Blacklisted : gtk_menu_button_set_menu_model

// Blacklisted : gtk_menu_button_set_popup

// Blacklisted : gtk_menu_shell_bind_model

// Blacklisted : gtk_search_entry_new

// Blacklisted : gtk_text_view_get_input_hints

// Blacklisted : gtk_text_view_get_input_purpose

// Blacklisted : gtk_text_view_set_input_hints

// Blacklisted : gtk_text_view_set_input_purpose

// Blacklisted : gtk_widget_insert_action_group

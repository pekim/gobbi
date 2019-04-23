// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void screen_sizeChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_size_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "size-changed", G_CALLBACK(screen_sizeChangedHandler), data);
	}

*/
import "C"

type signalScreenSizeChangedDetail struct {
	callback  ScreenSignalSizeChangedCallback
	handlerID C.gulong
}

var signalScreenSizeChangedId int
var signalScreenSizeChangedMap = make(map[int]signalScreenSizeChangedDetail)
var signalScreenSizeChangedLock sync.RWMutex

// ScreenSignalSizeChangedCallback is a callback function for a 'size-changed' signal emitted from a Screen.
type ScreenSignalSizeChangedCallback func()

/*
ConnectSizeChanged connects the callback to the 'size-changed' signal for the Screen.

The returned value represents the connection, and may be passed to DisconnectSizeChanged to remove it.
*/
func (recv *Screen) ConnectSizeChanged(callback ScreenSignalSizeChangedCallback) int {
	signalScreenSizeChangedLock.Lock()
	defer signalScreenSizeChangedLock.Unlock()

	signalScreenSizeChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Screen_signal_connect_size_changed(instance, C.gpointer(uintptr(signalScreenSizeChangedId)))

	detail := signalScreenSizeChangedDetail{callback, handlerID}
	signalScreenSizeChangedMap[signalScreenSizeChangedId] = detail

	return signalScreenSizeChangedId
}

/*
DisconnectSizeChanged disconnects a callback from the 'size-changed' signal for the Screen.

The connectionID should be a value returned from a call to ConnectSizeChanged.
*/
func (recv *Screen) DisconnectSizeChanged(connectionID int) {
	signalScreenSizeChangedLock.Lock()
	defer signalScreenSizeChangedLock.Unlock()

	detail, exists := signalScreenSizeChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScreenSizeChangedMap, connectionID)
}

//export screen_sizeChangedHandler
func screen_sizeChangedHandler(_ *C.GObject, data C.gpointer) {
	signalScreenSizeChangedLock.RLock()
	defer signalScreenSizeChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScreenSizeChangedMap[index].callback
	callback()
}

// Blacklisted : gdk_screen_get_default

// Blacklisted : gdk_screen_get_display

// Blacklisted : gdk_screen_get_height

// Blacklisted : gdk_screen_get_height_mm

// Blacklisted : gdk_screen_get_monitor_at_point

// Blacklisted : gdk_screen_get_monitor_at_window

// Blacklisted : gdk_screen_get_monitor_geometry

// Blacklisted : gdk_screen_get_n_monitors

// Blacklisted : gdk_screen_get_number

// Blacklisted : gdk_screen_get_root_window

// Blacklisted : gdk_screen_get_setting

// Blacklisted : gdk_screen_get_system_visual

// Blacklisted : gdk_screen_get_toplevel_windows

// Blacklisted : gdk_screen_get_width

// Blacklisted : gdk_screen_get_width_mm

// Blacklisted : gdk_screen_list_visuals

// Blacklisted : gdk_screen_make_display_name

// Blacklisted : gdk_window_fullscreen

// Blacklisted : gdk_window_set_skip_pager_hint

// Blacklisted : gdk_window_set_skip_taskbar_hint

// Blacklisted : gdk_window_unfullscreen

// This is a generated file - DO NOT EDIT
// +build gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void screen_compositedChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_composited_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "composited-changed", G_CALLBACK(screen_compositedChangedHandler), data);
	}

*/
import "C"

type signalScreenCompositedChangedDetail struct {
	callback  ScreenSignalCompositedChangedCallback
	handlerID C.gulong
}

var signalScreenCompositedChangedId int
var signalScreenCompositedChangedMap = make(map[int]signalScreenCompositedChangedDetail)
var signalScreenCompositedChangedLock sync.RWMutex

// ScreenSignalCompositedChangedCallback is a callback function for a 'composited-changed' signal emitted from a Screen.
type ScreenSignalCompositedChangedCallback func()

/*
ConnectCompositedChanged connects the callback to the 'composited-changed' signal for the Screen.

The returned value represents the connection, and may be passed to DisconnectCompositedChanged to remove it.
*/
func (recv *Screen) ConnectCompositedChanged(callback ScreenSignalCompositedChangedCallback) int {
	signalScreenCompositedChangedLock.Lock()
	defer signalScreenCompositedChangedLock.Unlock()

	signalScreenCompositedChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Screen_signal_connect_composited_changed(instance, C.gpointer(uintptr(signalScreenCompositedChangedId)))

	detail := signalScreenCompositedChangedDetail{callback, handlerID}
	signalScreenCompositedChangedMap[signalScreenCompositedChangedId] = detail

	return signalScreenCompositedChangedId
}

/*
DisconnectCompositedChanged disconnects a callback from the 'composited-changed' signal for the Screen.

The connectionID should be a value returned from a call to ConnectCompositedChanged.
*/
func (recv *Screen) DisconnectCompositedChanged(connectionID int) {
	signalScreenCompositedChangedLock.Lock()
	defer signalScreenCompositedChangedLock.Unlock()

	detail, exists := signalScreenCompositedChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScreenCompositedChangedMap, connectionID)
}

//export screen_compositedChangedHandler
func screen_compositedChangedHandler(_ *C.GObject, data C.gpointer) {
	signalScreenCompositedChangedLock.RLock()
	defer signalScreenCompositedChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScreenCompositedChangedMap[index].callback
	callback()
}

// Blacklisted : gdk_screen_get_active_window

// Blacklisted : gdk_screen_get_font_options

// Blacklisted : gdk_screen_get_resolution

// Blacklisted : gdk_screen_get_window_stack

// Blacklisted : gdk_screen_is_composited

// Blacklisted : gdk_screen_set_font_options

// Blacklisted : gdk_screen_set_resolution

// Blacklisted : gdk_window_get_type_hint

// Blacklisted : gdk_window_input_shape_combine_region

// Blacklisted : gdk_window_merge_child_input_shapes

// Blacklisted : gdk_window_set_child_input_shapes

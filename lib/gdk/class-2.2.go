// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void display_closedHandler(GObject *, gboolean, gpointer);

	static gulong Display_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(display_closedHandler), data);
	}

*/
/*

	void displaymanager_displayOpenedHandler(GObject *, GdkDisplay *, gpointer);

	static gulong DisplayManager_signal_connect_display_opened(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "display-opened", G_CALLBACK(displaymanager_displayOpenedHandler), data);
	}

*/
/*

	void keymap_keysChangedHandler(GObject *, gpointer);

	static gulong Keymap_signal_connect_keys_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keys-changed", G_CALLBACK(keymap_keysChangedHandler), data);
	}

*/
/*

	void screen_sizeChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_size_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "size-changed", G_CALLBACK(screen_sizeChangedHandler), data);
	}

*/
import "C"

// Blacklisted : gdk_cursor_new_for_display

// Blacklisted : gdk_cursor_get_display

type signalDisplayClosedDetail struct {
	callback  DisplaySignalClosedCallback
	handlerID C.gulong
}

var signalDisplayClosedId int
var signalDisplayClosedMap = make(map[int]signalDisplayClosedDetail)
var signalDisplayClosedLock sync.RWMutex

// DisplaySignalClosedCallback is a callback function for a 'closed' signal emitted from a Display.
type DisplaySignalClosedCallback func(isError bool)

/*
ConnectClosed connects the callback to the 'closed' signal for the Display.

The returned value represents the connection, and may be passed to DisconnectClosed to remove it.
*/
func (recv *Display) ConnectClosed(callback DisplaySignalClosedCallback) int {
	signalDisplayClosedLock.Lock()
	defer signalDisplayClosedLock.Unlock()

	signalDisplayClosedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Display_signal_connect_closed(instance, C.gpointer(uintptr(signalDisplayClosedId)))

	detail := signalDisplayClosedDetail{callback, handlerID}
	signalDisplayClosedMap[signalDisplayClosedId] = detail

	return signalDisplayClosedId
}

/*
DisconnectClosed disconnects a callback from the 'closed' signal for the Display.

The connectionID should be a value returned from a call to ConnectClosed.
*/
func (recv *Display) DisconnectClosed(connectionID int) {
	signalDisplayClosedLock.Lock()
	defer signalDisplayClosedLock.Unlock()

	detail, exists := signalDisplayClosedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplayClosedMap, connectionID)
}

//export display_closedHandler
func display_closedHandler(_ *C.GObject, c_is_error C.gboolean, data C.gpointer) {
	signalDisplayClosedLock.RLock()
	defer signalDisplayClosedLock.RUnlock()

	isError := c_is_error == C.TRUE

	index := int(uintptr(data))
	callback := signalDisplayClosedMap[index].callback
	callback(isError)
}

// Blacklisted : gdk_display_get_default

// Blacklisted : gdk_display_open

// Blacklisted : gdk_display_beep

// Blacklisted : gdk_display_close

// Blacklisted : gdk_display_get_default_screen

// Unsupported : gdk_display_get_event : no return generator

// Blacklisted : gdk_display_get_n_screens

// Blacklisted : gdk_display_get_name

// Unsupported : gdk_display_get_pointer : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Blacklisted : gdk_display_get_screen

// Blacklisted : gdk_display_get_window_at_pointer

// Blacklisted : gdk_display_keyboard_ungrab

// Blacklisted : gdk_display_list_devices

// Unsupported : gdk_display_peek_event : no return generator

// Blacklisted : gdk_display_pointer_is_grabbed

// Blacklisted : gdk_display_pointer_ungrab

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event (const GdkEvent*) for param event

// Blacklisted : gdk_display_set_double_click_time

// Blacklisted : gdk_display_sync

type signalDisplayManagerDisplayOpenedDetail struct {
	callback  DisplayManagerSignalDisplayOpenedCallback
	handlerID C.gulong
}

var signalDisplayManagerDisplayOpenedId int
var signalDisplayManagerDisplayOpenedMap = make(map[int]signalDisplayManagerDisplayOpenedDetail)
var signalDisplayManagerDisplayOpenedLock sync.RWMutex

// DisplayManagerSignalDisplayOpenedCallback is a callback function for a 'display-opened' signal emitted from a DisplayManager.
type DisplayManagerSignalDisplayOpenedCallback func(display *Display)

/*
ConnectDisplayOpened connects the callback to the 'display-opened' signal for the DisplayManager.

The returned value represents the connection, and may be passed to DisconnectDisplayOpened to remove it.
*/
func (recv *DisplayManager) ConnectDisplayOpened(callback DisplayManagerSignalDisplayOpenedCallback) int {
	signalDisplayManagerDisplayOpenedLock.Lock()
	defer signalDisplayManagerDisplayOpenedLock.Unlock()

	signalDisplayManagerDisplayOpenedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DisplayManager_signal_connect_display_opened(instance, C.gpointer(uintptr(signalDisplayManagerDisplayOpenedId)))

	detail := signalDisplayManagerDisplayOpenedDetail{callback, handlerID}
	signalDisplayManagerDisplayOpenedMap[signalDisplayManagerDisplayOpenedId] = detail

	return signalDisplayManagerDisplayOpenedId
}

/*
DisconnectDisplayOpened disconnects a callback from the 'display-opened' signal for the DisplayManager.

The connectionID should be a value returned from a call to ConnectDisplayOpened.
*/
func (recv *DisplayManager) DisconnectDisplayOpened(connectionID int) {
	signalDisplayManagerDisplayOpenedLock.Lock()
	defer signalDisplayManagerDisplayOpenedLock.Unlock()

	detail, exists := signalDisplayManagerDisplayOpenedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplayManagerDisplayOpenedMap, connectionID)
}

//export displaymanager_displayOpenedHandler
func displaymanager_displayOpenedHandler(_ *C.GObject, c_display *C.GdkDisplay, data C.gpointer) {
	signalDisplayManagerDisplayOpenedLock.RLock()
	defer signalDisplayManagerDisplayOpenedLock.RUnlock()

	display := DisplayNewFromC(unsafe.Pointer(c_display))

	index := int(uintptr(data))
	callback := signalDisplayManagerDisplayOpenedMap[index].callback
	callback(display)
}

// Blacklisted : gdk_display_manager_get

// Blacklisted : gdk_display_manager_get_default_display

// Blacklisted : gdk_display_manager_list_displays

// Blacklisted : gdk_display_manager_set_default_display

type signalKeymapKeysChangedDetail struct {
	callback  KeymapSignalKeysChangedCallback
	handlerID C.gulong
}

var signalKeymapKeysChangedId int
var signalKeymapKeysChangedMap = make(map[int]signalKeymapKeysChangedDetail)
var signalKeymapKeysChangedLock sync.RWMutex

// KeymapSignalKeysChangedCallback is a callback function for a 'keys-changed' signal emitted from a Keymap.
type KeymapSignalKeysChangedCallback func()

/*
ConnectKeysChanged connects the callback to the 'keys-changed' signal for the Keymap.

The returned value represents the connection, and may be passed to DisconnectKeysChanged to remove it.
*/
func (recv *Keymap) ConnectKeysChanged(callback KeymapSignalKeysChangedCallback) int {
	signalKeymapKeysChangedLock.Lock()
	defer signalKeymapKeysChangedLock.Unlock()

	signalKeymapKeysChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Keymap_signal_connect_keys_changed(instance, C.gpointer(uintptr(signalKeymapKeysChangedId)))

	detail := signalKeymapKeysChangedDetail{callback, handlerID}
	signalKeymapKeysChangedMap[signalKeymapKeysChangedId] = detail

	return signalKeymapKeysChangedId
}

/*
DisconnectKeysChanged disconnects a callback from the 'keys-changed' signal for the Keymap.

The connectionID should be a value returned from a call to ConnectKeysChanged.
*/
func (recv *Keymap) DisconnectKeysChanged(connectionID int) {
	signalKeymapKeysChangedLock.Lock()
	defer signalKeymapKeysChangedLock.Unlock()

	detail, exists := signalKeymapKeysChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalKeymapKeysChangedMap, connectionID)
}

//export keymap_keysChangedHandler
func keymap_keysChangedHandler(_ *C.GObject, data C.gpointer) {
	signalKeymapKeysChangedLock.RLock()
	defer signalKeymapKeysChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalKeymapKeysChangedMap[index].callback
	callback()
}

// Blacklisted : gdk_keymap_get_for_display

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

// Blacklisted : gdk_visual_get_screen

// Blacklisted : gdk_window_fullscreen

// Blacklisted : gdk_window_set_skip_pager_hint

// Blacklisted : gdk_window_set_skip_taskbar_hint

// Blacklisted : gdk_window_unfullscreen

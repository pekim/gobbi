// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void screen_monitorsChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_monitors_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "monitors-changed", G_CALLBACK(screen_monitorsChangedHandler), data);
	}

*/
import "C"

type signalScreenMonitorsChangedDetail struct {
	callback  ScreenSignalMonitorsChangedCallback
	handlerID C.gulong
}

var signalScreenMonitorsChangedId int
var signalScreenMonitorsChangedMap = make(map[int]signalScreenMonitorsChangedDetail)
var signalScreenMonitorsChangedLock sync.RWMutex

// ScreenSignalMonitorsChangedCallback is a callback function for a 'monitors-changed' signal emitted from a Screen.
type ScreenSignalMonitorsChangedCallback func()

/*
ConnectMonitorsChanged connects the callback to the 'monitors-changed' signal for the Screen.

The returned value represents the connection, and may be passed to DisconnectMonitorsChanged to remove it.
*/
func (recv *Screen) ConnectMonitorsChanged(callback ScreenSignalMonitorsChangedCallback) int {
	signalScreenMonitorsChangedLock.Lock()
	defer signalScreenMonitorsChangedLock.Unlock()

	signalScreenMonitorsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Screen_signal_connect_monitors_changed(instance, C.gpointer(uintptr(signalScreenMonitorsChangedId)))

	detail := signalScreenMonitorsChangedDetail{callback, handlerID}
	signalScreenMonitorsChangedMap[signalScreenMonitorsChangedId] = detail

	return signalScreenMonitorsChangedId
}

/*
DisconnectMonitorsChanged disconnects a callback from the 'monitors-changed' signal for the Screen.

The connectionID should be a value returned from a call to ConnectMonitorsChanged.
*/
func (recv *Screen) DisconnectMonitorsChanged(connectionID int) {
	signalScreenMonitorsChangedLock.Lock()
	defer signalScreenMonitorsChangedLock.Unlock()

	detail, exists := signalScreenMonitorsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScreenMonitorsChangedMap, connectionID)
}

//export screen_monitorsChangedHandler
func screen_monitorsChangedHandler(_ *C.GObject, data C.gpointer) {
	signalScreenMonitorsChangedLock.RLock()
	defer signalScreenMonitorsChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScreenMonitorsChangedMap[index].callback
	callback()
}

// Blacklisted : gdk_screen_get_monitor_height_mm

// Blacklisted : gdk_screen_get_monitor_plug_name

// Blacklisted : gdk_screen_get_monitor_width_mm

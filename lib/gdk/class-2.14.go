// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"fmt"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void Screen_monitorsChangedHandler();

	static gulong Screen_signal_connect_monitors_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "monitors-changed", Screen_monitorsChangedHandler, data);
	}

*/
import "C"

// AppLaunchContextNew is a wrapper around the C function gdk_app_launch_context_new.
func AppLaunchContextNew() *AppLaunchContext {
	retC := C.gdk_app_launch_context_new()
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDesktop is a wrapper around the C function gdk_app_launch_context_set_desktop.
func (recv *AppLaunchContext) SetDesktop(desktop int32) {
	c_desktop := (C.gint)(desktop)

	C.gdk_app_launch_context_set_desktop((*C.GdkAppLaunchContext)(recv.native), c_desktop)

	return
}

// SetDisplay is a wrapper around the C function gdk_app_launch_context_set_display.
func (recv *AppLaunchContext) SetDisplay(display *Display) {
	c_display := (*C.GdkDisplay)(display.ToC())

	C.gdk_app_launch_context_set_display((*C.GdkAppLaunchContext)(recv.native), c_display)

	return
}

// Unsupported : gdk_app_launch_context_set_icon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SetIconName is a wrapper around the C function gdk_app_launch_context_set_icon_name.
func (recv *AppLaunchContext) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gdk_app_launch_context_set_icon_name((*C.GdkAppLaunchContext)(recv.native), c_icon_name)

	return
}

// SetScreen is a wrapper around the C function gdk_app_launch_context_set_screen.
func (recv *AppLaunchContext) SetScreen(screen *Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gdk_app_launch_context_set_screen((*C.GdkAppLaunchContext)(recv.native), c_screen)

	return
}

// SetTimestamp is a wrapper around the C function gdk_app_launch_context_set_timestamp.
func (recv *AppLaunchContext) SetTimestamp(timestamp uint32) {
	c_timestamp := (C.guint32)(timestamp)

	C.gdk_app_launch_context_set_timestamp((*C.GdkAppLaunchContext)(recv.native), c_timestamp)

	return
}

var signalScreenMonitorsChangedId int
var signalScreenMonitorsChangedMap = make(map[int]ScreenSignalMonitorsChangedCallback)
var signalScreenMonitorsChangedLock sync.Mutex

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
	signalScreenMonitorsChangedMap[signalScreenMonitorsChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Screen_signal_connect_monitors_changed(instance, C.gpointer(uintptr(signalScreenMonitorsChangedId)))
	return int(retC)
}

/*
DisconnectMonitorsChanged disconnects a callback from the 'monitors-changed' signal for the Screen.

The connectionID should be a value returned from a call to ConnectMonitorsChanged.
*/
func (recv *Screen) DisconnectMonitorsChanged(connectionID int) {
	signalScreenMonitorsChangedLock.Lock()
	defer signalScreenMonitorsChangedLock.Unlock()

	_, exists := signalScreenMonitorsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalScreenMonitorsChangedMap, connectionID)
}

//export Screen_monitorsChangedHandler
func Screen_monitorsChangedHandler() {
	fmt.Println("cb")
}

// GetMonitorHeightMm is a wrapper around the C function gdk_screen_get_monitor_height_mm.
func (recv *Screen) GetMonitorHeightMm(monitorNum int32) int32 {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_height_mm((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := (int32)(retC)

	return retGo
}

// GetMonitorPlugName is a wrapper around the C function gdk_screen_get_monitor_plug_name.
func (recv *Screen) GetMonitorPlugName(monitorNum int32) string {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_plug_name((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetMonitorWidthMm is a wrapper around the C function gdk_screen_get_monitor_width_mm.
func (recv *Screen) GetMonitorWidthMm(monitorNum int32) int32 {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_width_mm((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported signal 'create-surface' for Window : return value cairo.Surface :

// Unsupported signal 'pick-embedded-child' for Window : return value Window :

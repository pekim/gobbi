// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void screen_monitorsChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_monitors_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "monitors-changed", G_CALLBACK(screen_monitorsChangedHandler), data);
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
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	C.gdk_app_launch_context_set_display((*C.GdkAppLaunchContext)(recv.native), c_display)

	return
}

// SetIcon is a wrapper around the C function gdk_app_launch_context_set_icon.
func (recv *AppLaunchContext) SetIcon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gdk_app_launch_context_set_icon((*C.GdkAppLaunchContext)(recv.native), c_icon)

	return
}

// SetIconName is a wrapper around the C function gdk_app_launch_context_set_icon_name.
func (recv *AppLaunchContext) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gdk_app_launch_context_set_icon_name((*C.GdkAppLaunchContext)(recv.native), c_icon_name)

	return
}

// SetScreen is a wrapper around the C function gdk_app_launch_context_set_screen.
func (recv *AppLaunchContext) SetScreen(screen *Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gdk_app_launch_context_set_screen((*C.GdkAppLaunchContext)(recv.native), c_screen)

	return
}

// SetTimestamp is a wrapper around the C function gdk_app_launch_context_set_timestamp.
func (recv *AppLaunchContext) SetTimestamp(timestamp uint32) {
	c_timestamp := (C.guint32)(timestamp)

	C.gdk_app_launch_context_set_timestamp((*C.GdkAppLaunchContext)(recv.native), c_timestamp)

	return
}

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

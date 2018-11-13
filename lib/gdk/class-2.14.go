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

// Creates a new #GdkAppLaunchContext.
/*

C function : gdk_app_launch_context_new
*/
func AppLaunchContextNew() *AppLaunchContext {
	retC := C.gdk_app_launch_context_new()
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the workspace on which applications will be launched when
// using this context when running under a window manager that
// supports multiple workspaces, as described in the
// [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec).
//
// When the workspace is not specified or @desktop is set to -1,
// it is up to the window manager to pick one, typically it will
// be the current workspace.
/*

C function : gdk_app_launch_context_set_desktop
*/
func (recv *AppLaunchContext) SetDesktop(desktop int32) {
	c_desktop := (C.gint)(desktop)

	C.gdk_app_launch_context_set_desktop((*C.GdkAppLaunchContext)(recv.native), c_desktop)

	return
}

// Sets the display on which applications will be launched when
// using this context. See also gdk_app_launch_context_set_screen().
/*

C function : gdk_app_launch_context_set_display
*/
func (recv *AppLaunchContext) SetDisplay(display *Display) {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	C.gdk_app_launch_context_set_display((*C.GdkAppLaunchContext)(recv.native), c_display)

	return
}

// Sets the icon for applications that are launched with this
// context.
//
// Window Managers can use this information when displaying startup
// notification.
//
// See also gdk_app_launch_context_set_icon_name().
/*

C function : gdk_app_launch_context_set_icon
*/
func (recv *AppLaunchContext) SetIcon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gdk_app_launch_context_set_icon((*C.GdkAppLaunchContext)(recv.native), c_icon)

	return
}

// Sets the icon for applications that are launched with this context.
// The @icon_name will be interpreted in the same way as the Icon field
// in desktop files. See also gdk_app_launch_context_set_icon().
//
// If both @icon and @icon_name are set, the @icon_name takes priority.
// If neither @icon or @icon_name is set, the icon is taken from either
// the file that is passed to launched application or from the #GAppInfo
// for the launched application itself.
/*

C function : gdk_app_launch_context_set_icon_name
*/
func (recv *AppLaunchContext) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gdk_app_launch_context_set_icon_name((*C.GdkAppLaunchContext)(recv.native), c_icon_name)

	return
}

// Sets the screen on which applications will be launched when
// using this context. See also gdk_app_launch_context_set_display().
//
// If both @screen and @display are set, the @screen takes priority.
// If neither @screen or @display are set, the default screen and
// display are used.
/*

C function : gdk_app_launch_context_set_screen
*/
func (recv *AppLaunchContext) SetScreen(screen *Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gdk_app_launch_context_set_screen((*C.GdkAppLaunchContext)(recv.native), c_screen)

	return
}

// Sets the timestamp of @context. The timestamp should ideally
// be taken from the event that triggered the launch.
//
// Window managers can use this information to avoid moving the
// focus to the newly launched application when the user is busy
// typing in another window. This is also known as 'focus stealing
// prevention'.
/*

C function : gdk_app_launch_context_set_timestamp
*/
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
	index := int(uintptr(data))
	callback := signalScreenMonitorsChangedMap[index].callback
	callback()
}

// Gets the height in millimeters of the specified monitor.
/*

C function : gdk_screen_get_monitor_height_mm
*/
func (recv *Screen) GetMonitorHeightMm(monitorNum int32) int32 {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_height_mm((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := (int32)(retC)

	return retGo
}

// Returns the output name of the specified monitor.
// Usually something like VGA, DVI, or TV, not the actual
// product name of the display device.
/*

C function : gdk_screen_get_monitor_plug_name
*/
func (recv *Screen) GetMonitorPlugName(monitorNum int32) string {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_plug_name((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the width in millimeters of the specified monitor, if available.
/*

C function : gdk_screen_get_monitor_width_mm
*/
func (recv *Screen) GetMonitorWidthMm(monitorNum int32) int32 {
	c_monitor_num := (C.gint)(monitorNum)

	retC := C.gdk_screen_get_monitor_width_mm((*C.GdkScreen)(recv.native), c_monitor_num)
	retGo := (int32)(retC)

	return retGo
}

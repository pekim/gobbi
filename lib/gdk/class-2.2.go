// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// CursorNewForDisplay is a wrapper around the C function gdk_cursor_new_for_display.
func CursorNewForDisplay(display *Display, cursorType CursorType) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new_for_display(c_display, c_cursor_type)
	gobject.ObjectNewFromC(unsafe.Pointer(retC)).Take()
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_cursor_get_display.
func (recv *Cursor) GetDisplay() *Display {
	retC := C.gdk_cursor_get_display((*C.GdkCursor)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// DisplayGetDefault is a wrapper around the C function gdk_display_get_default.
func DisplayGetDefault() *Display {
	retC := C.gdk_display_get_default()
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// DisplayOpen is a wrapper around the C function gdk_display_open.
func DisplayOpen(displayName string) *Display {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	retC := C.gdk_display_open(c_display_name)
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Beep is a wrapper around the C function gdk_display_beep.
func (recv *Display) Beep() {
	C.gdk_display_beep((*C.GdkDisplay)(recv.native))

	return
}

// Close is a wrapper around the C function gdk_display_close.
func (recv *Display) Close() {
	C.gdk_display_close((*C.GdkDisplay)(recv.native))

	return
}

// GetDefaultScreen is a wrapper around the C function gdk_display_get_default_screen.
func (recv *Display) GetDefaultScreen() *Screen {
	retC := C.gdk_display_get_default_screen((*C.GdkDisplay)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_display_get_event : no return generator

// GetNScreens is a wrapper around the C function gdk_display_get_n_screens.
func (recv *Display) GetNScreens() int32 {
	retC := C.gdk_display_get_n_screens((*C.GdkDisplay)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetName is a wrapper around the C function gdk_display_get_name.
func (recv *Display) GetName() string {
	retC := C.gdk_display_get_name((*C.GdkDisplay)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gdk_display_get_pointer : unsupported parameter mask : GdkModifierType* with indirection level of 1

// GetScreen is a wrapper around the C function gdk_display_get_screen.
func (recv *Display) GetScreen(screenNum int32) *Screen {
	c_screen_num := (C.gint)(screenNum)

	retC := C.gdk_display_get_screen((*C.GdkDisplay)(recv.native), c_screen_num)
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindowAtPointer is a wrapper around the C function gdk_display_get_window_at_pointer.
func (recv *Display) GetWindowAtPointer() (*Window, int32, int32) {
	var c_win_x C.gint

	var c_win_y C.gint

	retC := C.gdk_display_get_window_at_pointer((*C.GdkDisplay)(recv.native), &c_win_x, &c_win_y)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	winX := (int32)(c_win_x)

	winY := (int32)(c_win_y)

	return retGo, winX, winY
}

// KeyboardUngrab is a wrapper around the C function gdk_display_keyboard_ungrab.
func (recv *Display) KeyboardUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_display_keyboard_ungrab((*C.GdkDisplay)(recv.native), c_time_)

	return
}

// ListDevices is a wrapper around the C function gdk_display_list_devices.
func (recv *Display) ListDevices() *glib.List {
	retC := C.gdk_display_list_devices((*C.GdkDisplay)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_display_peek_event : no return generator

// PointerIsGrabbed is a wrapper around the C function gdk_display_pointer_is_grabbed.
func (recv *Display) PointerIsGrabbed() bool {
	retC := C.gdk_display_pointer_is_grabbed((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PointerUngrab is a wrapper around the C function gdk_display_pointer_ungrab.
func (recv *Display) PointerUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_display_pointer_ungrab((*C.GdkDisplay)(recv.native), c_time_)

	return
}

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event (const GdkEvent*) for param event

// SetDoubleClickTime is a wrapper around the C function gdk_display_set_double_click_time.
func (recv *Display) SetDoubleClickTime(msec uint32) {
	c_msec := (C.guint)(msec)

	C.gdk_display_set_double_click_time((*C.GdkDisplay)(recv.native), c_msec)

	return
}

// Sync is a wrapper around the C function gdk_display_sync.
func (recv *Display) Sync() {
	C.gdk_display_sync((*C.GdkDisplay)(recv.native))

	return
}

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

// DisplayManagerGet is a wrapper around the C function gdk_display_manager_get.
func DisplayManagerGet() *DisplayManager {
	retC := C.gdk_display_manager_get()
	retGo := DisplayManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultDisplay is a wrapper around the C function gdk_display_manager_get_default_display.
func (recv *DisplayManager) GetDefaultDisplay() *Display {
	retC := C.gdk_display_manager_get_default_display((*C.GdkDisplayManager)(recv.native))
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ListDisplays is a wrapper around the C function gdk_display_manager_list_displays.
func (recv *DisplayManager) ListDisplays() *glib.SList {
	retC := C.gdk_display_manager_list_displays((*C.GdkDisplayManager)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDefaultDisplay is a wrapper around the C function gdk_display_manager_set_default_display.
func (recv *DisplayManager) SetDefaultDisplay(display *Display) {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	C.gdk_display_manager_set_default_display((*C.GdkDisplayManager)(recv.native), c_display)

	return
}

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

// KeymapGetForDisplay is a wrapper around the C function gdk_keymap_get_for_display.
func KeymapGetForDisplay(display *Display) *Keymap {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	retC := C.gdk_keymap_get_for_display(c_display)
	retGo := KeymapNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// ScreenGetDefault is a wrapper around the C function gdk_screen_get_default.
func ScreenGetDefault() *Screen {
	retC := C.gdk_screen_get_default()
	var retGo (*Screen)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ScreenNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_screen_get_display.
func (recv *Screen) GetDisplay() *Display {
	retC := C.gdk_screen_get_display((*C.GdkScreen)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHeight is a wrapper around the C function gdk_screen_get_height.
func (recv *Screen) GetHeight() int32 {
	retC := C.gdk_screen_get_height((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHeightMm is a wrapper around the C function gdk_screen_get_height_mm.
func (recv *Screen) GetHeightMm() int32 {
	retC := C.gdk_screen_get_height_mm((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMonitorAtPoint is a wrapper around the C function gdk_screen_get_monitor_at_point.
func (recv *Screen) GetMonitorAtPoint(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gdk_screen_get_monitor_at_point((*C.GdkScreen)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// GetMonitorAtWindow is a wrapper around the C function gdk_screen_get_monitor_at_window.
func (recv *Screen) GetMonitorAtWindow(window *Window) int32 {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_screen_get_monitor_at_window((*C.GdkScreen)(recv.native), c_window)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

// GetNMonitors is a wrapper around the C function gdk_screen_get_n_monitors.
func (recv *Screen) GetNMonitors() int32 {
	retC := C.gdk_screen_get_n_monitors((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNumber is a wrapper around the C function gdk_screen_get_number.
func (recv *Screen) GetNumber() int32 {
	retC := C.gdk_screen_get_number((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRootWindow is a wrapper around the C function gdk_screen_get_root_window.
func (recv *Screen) GetRootWindow() *Window {
	retC := C.gdk_screen_get_root_window((*C.GdkScreen)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSetting is a wrapper around the C function gdk_screen_get_setting.
func (recv *Screen) GetSetting(name string, value *gobject.Value) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.gdk_screen_get_setting((*C.GdkScreen)(recv.native), c_name, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// GetSystemVisual is a wrapper around the C function gdk_screen_get_system_visual.
func (recv *Screen) GetSystemVisual() *Visual {
	retC := C.gdk_screen_get_system_visual((*C.GdkScreen)(recv.native))
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetToplevelWindows is a wrapper around the C function gdk_screen_get_toplevel_windows.
func (recv *Screen) GetToplevelWindows() *glib.List {
	retC := C.gdk_screen_get_toplevel_windows((*C.GdkScreen)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gdk_screen_get_width.
func (recv *Screen) GetWidth() int32 {
	retC := C.gdk_screen_get_width((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWidthMm is a wrapper around the C function gdk_screen_get_width_mm.
func (recv *Screen) GetWidthMm() int32 {
	retC := C.gdk_screen_get_width_mm((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// ListVisuals is a wrapper around the C function gdk_screen_list_visuals.
func (recv *Screen) ListVisuals() *glib.List {
	retC := C.gdk_screen_list_visuals((*C.GdkScreen)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MakeDisplayName is a wrapper around the C function gdk_screen_make_display_name.
func (recv *Screen) MakeDisplayName() string {
	retC := C.gdk_screen_make_display_name((*C.GdkScreen)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gdk_visual_get_screen.
func (recv *Visual) GetScreen() *Screen {
	retC := C.gdk_visual_get_screen((*C.GdkVisual)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Fullscreen is a wrapper around the C function gdk_window_fullscreen.
func (recv *Window) Fullscreen() {
	C.gdk_window_fullscreen((*C.GdkWindow)(recv.native))

	return
}

// SetSkipPagerHint is a wrapper around the C function gdk_window_set_skip_pager_hint.
func (recv *Window) SetSkipPagerHint(skipsPager bool) {
	c_skips_pager :=
		boolToGboolean(skipsPager)

	C.gdk_window_set_skip_pager_hint((*C.GdkWindow)(recv.native), c_skips_pager)

	return
}

// SetSkipTaskbarHint is a wrapper around the C function gdk_window_set_skip_taskbar_hint.
func (recv *Window) SetSkipTaskbarHint(skipsTaskbar bool) {
	c_skips_taskbar :=
		boolToGboolean(skipsTaskbar)

	C.gdk_window_set_skip_taskbar_hint((*C.GdkWindow)(recv.native), c_skips_taskbar)

	return
}

// Unfullscreen is a wrapper around the C function gdk_window_unfullscreen.
func (recv *Window) Unfullscreen() {
	C.gdk_window_unfullscreen((*C.GdkWindow)(recv.native))

	return
}

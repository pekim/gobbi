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

// Creates a new cursor from the set of builtin cursors.
/*

C function : gdk_cursor_new_for_display
*/
func CursorNewForDisplay(display *Display, cursorType CursorType) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new_for_display(c_display, c_cursor_type)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the display on which the #GdkCursor is defined.
/*

C function : gdk_cursor_get_display
*/
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
var signalDisplayClosedLock sync.Mutex

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
	isError := c_is_error == C.TRUE

	index := int(uintptr(data))
	callback := signalDisplayClosedMap[index].callback
	callback(isError)
}

// Emits a short beep on @display
/*

C function : gdk_display_beep
*/
func (recv *Display) Beep() {
	C.gdk_display_beep((*C.GdkDisplay)(recv.native))

	return
}

// Closes the connection to the windowing system for the given display,
// and cleans up associated resources.
/*

C function : gdk_display_close
*/
func (recv *Display) Close() {
	C.gdk_display_close((*C.GdkDisplay)(recv.native))

	return
}

// Get the default #GdkScreen for @display.
/*

C function : gdk_display_get_default_screen
*/
func (recv *Display) GetDefaultScreen() *Screen {
	retC := C.gdk_display_get_default_screen((*C.GdkDisplay)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_display_get_event : no return generator

// Gets the number of screen managed by the @display.
/*

C function : gdk_display_get_n_screens
*/
func (recv *Display) GetNScreens() int32 {
	retC := C.gdk_display_get_n_screens((*C.GdkDisplay)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the name of the display.
/*

C function : gdk_display_get_name
*/
func (recv *Display) GetName() string {
	retC := C.gdk_display_get_name((*C.GdkDisplay)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gdk_display_get_pointer : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Returns a screen object for one of the screens of the display.
/*

C function : gdk_display_get_screen
*/
func (recv *Display) GetScreen(screenNum int32) *Screen {
	c_screen_num := (C.gint)(screenNum)

	retC := C.gdk_display_get_screen((*C.GdkDisplay)(recv.native), c_screen_num)
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains the window underneath the mouse pointer, returning the location
// of the pointer in that window in @win_x, @win_y for @screen. Returns %NULL
// if the window under the mouse pointer is not known to GDK (for example,
// belongs to another application).
/*

C function : gdk_display_get_window_at_pointer
*/
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

// Release any keyboard grab
/*

C function : gdk_display_keyboard_ungrab
*/
func (recv *Display) KeyboardUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_display_keyboard_ungrab((*C.GdkDisplay)(recv.native), c_time_)

	return
}

// Returns the list of available input devices attached to @display.
// The list is statically allocated and should not be freed.
/*

C function : gdk_display_list_devices
*/
func (recv *Display) ListDevices() *glib.List {
	retC := C.gdk_display_list_devices((*C.GdkDisplay)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_display_peek_event : no return generator

// Test if the pointer is grabbed.
/*

C function : gdk_display_pointer_is_grabbed
*/
func (recv *Display) PointerIsGrabbed() bool {
	retC := C.gdk_display_pointer_is_grabbed((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Release any pointer grab.
/*

C function : gdk_display_pointer_ungrab
*/
func (recv *Display) PointerUngrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_display_pointer_ungrab((*C.GdkDisplay)(recv.native), c_time_)

	return
}

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event (const GdkEvent*) for param event

// Sets the double click time (two clicks within this time interval
// count as a double click and result in a #GDK_2BUTTON_PRESS event).
// Applications should not set this, it is a global
// user-configured setting.
/*

C function : gdk_display_set_double_click_time
*/
func (recv *Display) SetDoubleClickTime(msec uint32) {
	c_msec := (C.guint)(msec)

	C.gdk_display_set_double_click_time((*C.GdkDisplay)(recv.native), c_msec)

	return
}

// Flushes any requests queued for the windowing system and waits until all
// requests have been handled. This is often used for making sure that the
// display is synchronized with the current state of the program. Calling
// gdk_display_sync() before gdk_error_trap_pop() makes sure that any errors
// generated from earlier requests are handled before the error trap is
// removed.
//
// This is most useful for X11. On windowing systems where requests are
// handled synchronously, this function will do nothing.
/*

C function : gdk_display_sync
*/
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
var signalDisplayManagerDisplayOpenedLock sync.Mutex

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
	display := DisplayNewFromC(unsafe.Pointer(c_display))

	index := int(uintptr(data))
	callback := signalDisplayManagerDisplayOpenedMap[index].callback
	callback(display)
}

// Gets the default #GdkDisplay.
/*

C function : gdk_display_manager_get_default_display
*/
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

// List all currently open displays.
/*

C function : gdk_display_manager_list_displays
*/
func (recv *DisplayManager) ListDisplays() *glib.SList {
	retC := C.gdk_display_manager_list_displays((*C.GdkDisplayManager)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets @display as the default display.
/*

C function : gdk_display_manager_set_default_display
*/
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
var signalKeymapKeysChangedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalKeymapKeysChangedMap[index].callback
	callback()
}

type signalScreenSizeChangedDetail struct {
	callback  ScreenSignalSizeChangedCallback
	handlerID C.gulong
}

var signalScreenSizeChangedId int
var signalScreenSizeChangedMap = make(map[int]signalScreenSizeChangedDetail)
var signalScreenSizeChangedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalScreenSizeChangedMap[index].callback
	callback()
}

// Gets the display to which the @screen belongs.
/*

C function : gdk_screen_get_display
*/
func (recv *Screen) GetDisplay() *Display {
	retC := C.gdk_screen_get_display((*C.GdkScreen)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the height of @screen in pixels. The returned size is in
// ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
/*

C function : gdk_screen_get_height
*/
func (recv *Screen) GetHeight() int32 {
	retC := C.gdk_screen_get_height((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the height of @screen in millimeters.
//
// Note that this value is somewhat ill-defined when the screen
// has multiple monitors of different resolution. It is recommended
// to use the monitor dimensions instead.
/*

C function : gdk_screen_get_height_mm
*/
func (recv *Screen) GetHeightMm() int32 {
	retC := C.gdk_screen_get_height_mm((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the monitor number in which the point (@x,@y) is located.
/*

C function : gdk_screen_get_monitor_at_point
*/
func (recv *Screen) GetMonitorAtPoint(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gdk_screen_get_monitor_at_point((*C.GdkScreen)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of the monitor in which the largest area of the
// bounding rectangle of @window resides.
/*

C function : gdk_screen_get_monitor_at_window
*/
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

// Returns the number of monitors which @screen consists of.
/*

C function : gdk_screen_get_n_monitors
*/
func (recv *Screen) GetNMonitors() int32 {
	retC := C.gdk_screen_get_n_monitors((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the index of @screen among the screens in the display
// to which it belongs. (See gdk_screen_get_display())
/*

C function : gdk_screen_get_number
*/
func (recv *Screen) GetNumber() int32 {
	retC := C.gdk_screen_get_number((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the root window of @screen.
/*

C function : gdk_screen_get_root_window
*/
func (recv *Screen) GetRootWindow() *Window {
	retC := C.gdk_screen_get_root_window((*C.GdkScreen)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Retrieves a desktop-wide setting such as double-click time
// for the #GdkScreen @screen.
//
// FIXME needs a list of valid settings here, or a link to
// more information.
/*

C function : gdk_screen_get_setting
*/
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

// Get the system’s default visual for @screen.
// This is the visual for the root window of the display.
// The return value should not be freed.
/*

C function : gdk_screen_get_system_visual
*/
func (recv *Screen) GetSystemVisual() *Visual {
	retC := C.gdk_screen_get_system_visual((*C.GdkScreen)(recv.native))
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains a list of all toplevel windows known to GDK on the screen @screen.
// A toplevel window is a child of the root window (see
// gdk_get_default_root_window()).
//
// The returned list should be freed with g_list_free(), but
// its elements need not be freed.
/*

C function : gdk_screen_get_toplevel_windows
*/
func (recv *Screen) GetToplevelWindows() *glib.List {
	retC := C.gdk_screen_get_toplevel_windows((*C.GdkScreen)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the width of @screen in pixels. The returned size is in
// ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
/*

C function : gdk_screen_get_width
*/
func (recv *Screen) GetWidth() int32 {
	retC := C.gdk_screen_get_width((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the width of @screen in millimeters.
//
// Note that this value is somewhat ill-defined when the screen
// has multiple monitors of different resolution. It is recommended
// to use the monitor dimensions instead.
/*

C function : gdk_screen_get_width_mm
*/
func (recv *Screen) GetWidthMm() int32 {
	retC := C.gdk_screen_get_width_mm((*C.GdkScreen)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Lists the available visuals for the specified @screen.
// A visual describes a hardware image data format.
// For example, a visual might support 24-bit color, or 8-bit color,
// and might expect pixels to be in a certain format.
//
// Call g_list_free() on the return value when you’re finished with it.
/*

C function : gdk_screen_list_visuals
*/
func (recv *Screen) ListVisuals() *glib.List {
	retC := C.gdk_screen_list_visuals((*C.GdkScreen)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Determines the name to pass to gdk_display_open() to get
// a #GdkDisplay with this screen as the default screen.
/*

C function : gdk_screen_make_display_name
*/
func (recv *Screen) MakeDisplayName() string {
	retC := C.gdk_screen_make_display_name((*C.GdkScreen)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the screen to which this visual belongs
/*

C function : gdk_visual_get_screen
*/
func (recv *Visual) GetScreen() *Screen {
	retC := C.gdk_visual_get_screen((*C.GdkVisual)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Moves the window into fullscreen mode. This means the
// window covers the entire screen and is above any panels
// or task bars.
//
// If the window was already fullscreen, then this function does nothing.
//
// On X11, asks the window manager to put @window in a fullscreen
// state, if the window manager supports this operation. Not all
// window managers support this, and some deliberately ignore it or
// don’t have a concept of “fullscreen”; so you can’t rely on the
// fullscreenification actually happening. But it will happen with
// most standard window managers, and GDK makes a best effort to get
// it to happen.
/*

C function : gdk_window_fullscreen
*/
func (recv *Window) Fullscreen() {
	C.gdk_window_fullscreen((*C.GdkWindow)(recv.native))

	return
}

// Toggles whether a window should appear in a pager (workspace
// switcher, or other desktop utility program that displays a small
// thumbnail representation of the windows on the desktop). If a
// window’s semantic type as specified with gdk_window_set_type_hint()
// already fully describes the window, this function should
// not be called in addition, instead you should
// allow the window to be treated according to standard policy for
// its semantic type.
/*

C function : gdk_window_set_skip_pager_hint
*/
func (recv *Window) SetSkipPagerHint(skipsPager bool) {
	c_skips_pager :=
		boolToGboolean(skipsPager)

	C.gdk_window_set_skip_pager_hint((*C.GdkWindow)(recv.native), c_skips_pager)

	return
}

// Toggles whether a window should appear in a task list or window
// list. If a window’s semantic type as specified with
// gdk_window_set_type_hint() already fully describes the window, this
// function should not be called in addition,
// instead you should allow the window to be treated according to
// standard policy for its semantic type.
/*

C function : gdk_window_set_skip_taskbar_hint
*/
func (recv *Window) SetSkipTaskbarHint(skipsTaskbar bool) {
	c_skips_taskbar :=
		boolToGboolean(skipsTaskbar)

	C.gdk_window_set_skip_taskbar_hint((*C.GdkWindow)(recv.native), c_skips_taskbar)

	return
}

// Moves the window out of fullscreen mode. If the window was not
// fullscreen, does nothing.
//
// On X11, asks the window manager to move @window out of the fullscreen
// state, if the window manager supports this operation. Not all
// window managers support this, and some deliberately ignore it or
// don’t have a concept of “fullscreen”; so you can’t rely on the
// unfullscreenification actually happening. But it will happen with
// most standard window managers, and GDK makes a best effort to get
// it to happen.
/*

C function : gdk_window_unfullscreen
*/
func (recv *Window) Unfullscreen() {
	C.gdk_window_unfullscreen((*C.GdkWindow)(recv.native))

	return
}

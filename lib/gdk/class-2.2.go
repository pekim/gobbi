// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_app_launch_context_set_icon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// CursorNewForDisplay is a wrapper around the C function gdk_cursor_new_for_display.
func CursorNewForDisplay(display *Display, cursorType CursorType) *Cursor {
	c_display := (*C.GdkDisplay)(display.ToC())

	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new_for_display(c_display, c_cursor_type)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_cursor_get_display.
func (recv *Cursor) GetDisplay() *Display {
	retC := C.gdk_cursor_get_display((*C.GdkCursor)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_cursor_get_surface : unsupported parameter x_hot : no type generator for gdouble, gdouble*

// Unsupported : gdk_device_get_axis : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_axis_value : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_history : unsupported parameter events : no param type

// Unsupported : gdk_device_get_key : unsupported parameter keyval : no type generator for guint, guint*

// Unsupported : gdk_device_get_position : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_device_get_position_double : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_device_get_state : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_window_at_position : unsupported parameter win_x : no type generator for gint, gint*

// Unsupported : gdk_device_get_window_at_position_double : unsupported parameter win_x : no type generator for gdouble, gdouble*

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

// Unsupported : gdk_display_get_maximal_cursor_size : unsupported parameter width : no type generator for guint, guint*

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

// Unsupported : gdk_display_get_pointer : unsupported parameter screen : record with indirection level of 2

// GetScreen is a wrapper around the C function gdk_display_get_screen.
func (recv *Display) GetScreen(screenNum int32) *Screen {
	c_screen_num := (C.gint)(screenNum)

	retC := C.gdk_display_get_screen((*C.GdkDisplay)(recv.native), c_screen_num)
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_display_get_window_at_pointer : unsupported parameter win_x : no type generator for gint, gint*

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

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// SetDoubleClickTime is a wrapper around the C function gdk_display_set_double_click_time.
func (recv *Display) SetDoubleClickTime(msec uint32) {
	c_msec := (C.guint)(msec)

	C.gdk_display_set_double_click_time((*C.GdkDisplay)(recv.native), c_msec)

	return
}

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

// Sync is a wrapper around the C function gdk_display_sync.
func (recv *Display) Sync() {
	C.gdk_display_sync((*C.GdkDisplay)(recv.native))

	return
}

// GetDefaultDisplay is a wrapper around the C function gdk_display_manager_get_default_display.
func (recv *DisplayManager) GetDefaultDisplay() *Display {
	retC := C.gdk_display_manager_get_default_display((*C.GdkDisplayManager)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

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
	c_display := (*C.GdkDisplay)(display.ToC())

	C.gdk_display_manager_set_default_display((*C.GdkDisplayManager)(recv.native), c_display)

	return
}

// Unsupported : gdk_frame_clock_get_refresh_info : unsupported parameter refresh_interval_return : no type generator for gint64, gint64*

// Unsupported : gdk_gl_context_get_required_version : unsupported parameter major : no type generator for gint, int*

// Unsupported : gdk_gl_context_get_version : unsupported parameter major : no type generator for gint, int*

// Unsupported : gdk_keymap_add_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_map_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_translate_keyboard_state : unsupported parameter keyval : no type generator for guint, guint*

// Unsupported : gdk_monitor_get_geometry : unsupported parameter geometry : Blacklisted record : GdkRectangle

// Unsupported : gdk_monitor_get_workarea : unsupported parameter workarea : Blacklisted record : GdkRectangle

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
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gdk_screen_get_monitor_at_window((*C.GdkScreen)(recv.native), c_window)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

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

	c_value := (*C.GValue)(value.ToC())

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

// Unsupported : gdk_seat_grab : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_visual_get_blue_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_green_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_red_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// GetScreen is a wrapper around the C function gdk_visual_get_screen.
func (recv *Visual) GetScreen() *Screen {
	retC := C.gdk_visual_get_screen((*C.GdkVisual)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

// Fullscreen is a wrapper around the C function gdk_window_fullscreen.
func (recv *Window) Fullscreen() {
	C.gdk_window_fullscreen((*C.GdkWindow)(recv.native))

	return
}

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// Unsupported : gdk_window_get_device_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_device_position_double : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_get_drag_protocol : unsupported parameter target : record with indirection level of 2

// Unsupported : gdk_window_get_frame_extents : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_get_geometry : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_origin : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_coords : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_origin : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_user_data : unsupported parameter data : no type generator for gpointer, gpointer*

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc, GdkWindowChildFunc

// Unsupported : gdk_window_invalidate_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

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

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*

// Unfullscreen is a wrapper around the C function gdk_window_unfullscreen.
func (recv *Window) Unfullscreen() {
	C.gdk_window_unfullscreen((*C.GdkWindow)(recv.native))

	return
}

// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
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

// Unsupported : gdk_display_get_event : no return generator

// Unsupported : gdk_display_get_maximal_cursor_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gdk_display_get_pointer : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_display_get_window_at_pointer : unsupported parameter win_x : no type generator for gint, gint*

// Unsupported : gdk_display_peek_event : no return generator

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

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

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

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

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_seat_grab : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_visual_get_blue_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_green_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_red_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

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

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*

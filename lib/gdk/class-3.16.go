// This is a generated file - DO NOT EDIT
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_app_launch_context_set_icon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gdk_cursor_get_surface : unsupported parameter x_hot : no type generator for gdouble, gdouble*

// Unsupported : gdk_device_get_axis : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_axis_value : unsupported parameter axes : no param type

// Unsupported : gdk_device_get_history : unsupported parameter events : no param type

// Unsupported : gdk_device_get_key : unsupported parameter keyval : no type generator for guint, guint*

// Unsupported : gdk_device_get_position : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_device_get_position_double : unsupported parameter screen : record with indirection level of 2

// GetProductId is a wrapper around the C function gdk_device_get_product_id.
func (recv *Device) GetProductId() string {
	retC := C.gdk_device_get_product_id((*C.GdkDevice)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gdk_device_get_state : unsupported parameter axes : no param type

// GetVendorId is a wrapper around the C function gdk_device_get_vendor_id.
func (recv *Device) GetVendorId() string {
	retC := C.gdk_device_get_vendor_id((*C.GdkDevice)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

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

// GetDebugEnabled is a wrapper around the C function gdk_gl_context_get_debug_enabled.
func (recv *GLContext) GetDebugEnabled() bool {
	retC := C.gdk_gl_context_get_debug_enabled((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_gl_context_get_display.
func (recv *GLContext) GetDisplay() *Display {
	retC := C.gdk_gl_context_get_display((*C.GdkGLContext)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetForwardCompatible is a wrapper around the C function gdk_gl_context_get_forward_compatible.
func (recv *GLContext) GetForwardCompatible() bool {
	retC := C.gdk_gl_context_get_forward_compatible((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_gl_context_get_required_version : unsupported parameter major : no type generator for gint, int*

// GetSharedContext is a wrapper around the C function gdk_gl_context_get_shared_context.
func (recv *GLContext) GetSharedContext() *GLContext {
	retC := C.gdk_gl_context_get_shared_context((*C.GdkGLContext)(recv.native))
	retGo := GLContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_gl_context_get_version : unsupported parameter major : no type generator for gint, int*

// GetWindow is a wrapper around the C function gdk_gl_context_get_window.
func (recv *GLContext) GetWindow() *Window {
	retC := C.gdk_gl_context_get_window((*C.GdkGLContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MakeCurrent is a wrapper around the C function gdk_gl_context_make_current.
func (recv *GLContext) MakeCurrent() {
	C.gdk_gl_context_make_current((*C.GdkGLContext)(recv.native))

	return
}

// Realize is a wrapper around the C function gdk_gl_context_realize.
func (recv *GLContext) Realize() (bool, error) {
	var cThrowableError *C.GError

	retC := C.gdk_gl_context_realize((*C.GdkGLContext)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetDebugEnabled is a wrapper around the C function gdk_gl_context_set_debug_enabled.
func (recv *GLContext) SetDebugEnabled(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.gdk_gl_context_set_debug_enabled((*C.GdkGLContext)(recv.native), c_enabled)

	return
}

// SetForwardCompatible is a wrapper around the C function gdk_gl_context_set_forward_compatible.
func (recv *GLContext) SetForwardCompatible(compatible bool) {
	c_compatible :=
		boolToGboolean(compatible)

	C.gdk_gl_context_set_forward_compatible((*C.GdkGLContext)(recv.native), c_compatible)

	return
}

// SetRequiredVersion is a wrapper around the C function gdk_gl_context_set_required_version.
func (recv *GLContext) SetRequiredVersion(major int32, minor int32) {
	c_major := (C.int)(major)

	c_minor := (C.int)(minor)

	C.gdk_gl_context_set_required_version((*C.GdkGLContext)(recv.native), c_major, c_minor)

	return
}

// Unsupported : gdk_keymap_add_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_get_entries_for_keycode : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_get_entries_for_keyval : unsupported parameter keys : no param type

// Unsupported : gdk_keymap_map_virtual_modifiers : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gdk_keymap_translate_keyboard_state : unsupported parameter keyval : no type generator for guint, guint*

// Unsupported : gdk_monitor_get_geometry : unsupported parameter geometry : Blacklisted record : GdkRectangle

// Unsupported : gdk_monitor_get_workarea : unsupported parameter workarea : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_seat_grab : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_visual_get_blue_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_green_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_red_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// CreateGlContext is a wrapper around the C function gdk_window_create_gl_context.
func (recv *Window) CreateGlContext() (*GLContext, error) {
	var cThrowableError *C.GError

	retC := C.gdk_window_create_gl_context((*C.GdkWindow)(recv.native), &cThrowableError)
	retGo := GLContextNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

// MarkPaintFromClip is a wrapper around the C function gdk_window_mark_paint_from_clip.
func (recv *Window) MarkPaintFromClip(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(cr.ToC())

	C.gdk_window_mark_paint_from_clip((*C.GdkWindow)(recv.native), c_cr)

	return
}

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*

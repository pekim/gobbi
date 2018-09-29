// This is a generated file - DO NOT EDIT
// +build gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

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

// GetCursorType is a wrapper around the C function gdk_cursor_get_cursor_type.
func (recv *Cursor) GetCursorType() CursorType {
	retC := C.gdk_cursor_get_cursor_type((*C.GdkCursor)(recv.native))
	retGo := (CursorType)(retC)

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

// Unsupported : gdk_display_get_event : no return generator

// Unsupported : gdk_display_get_maximal_cursor_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gdk_display_get_pointer : unsupported parameter screen : record with indirection level of 2

// Unsupported : gdk_display_get_window_at_pointer : unsupported parameter win_x : no type generator for gint, gint*

// IsClosed is a wrapper around the C function gdk_display_is_closed.
func (recv *Display) IsClosed() bool {
	retC := C.gdk_display_is_closed((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_display_peek_event : no return generator

// Unsupported : gdk_display_put_event : unsupported parameter event : no type generator for Event, const GdkEvent*

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

// GetActions is a wrapper around the C function gdk_drag_context_get_actions.
func (recv *DragContext) GetActions() DragAction {
	retC := C.gdk_drag_context_get_actions((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// GetSelectedAction is a wrapper around the C function gdk_drag_context_get_selected_action.
func (recv *DragContext) GetSelectedAction() DragAction {
	retC := C.gdk_drag_context_get_selected_action((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// GetSourceWindow is a wrapper around the C function gdk_drag_context_get_source_window.
func (recv *DragContext) GetSourceWindow() *Window {
	retC := C.gdk_drag_context_get_source_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSuggestedAction is a wrapper around the C function gdk_drag_context_get_suggested_action.
func (recv *DragContext) GetSuggestedAction() DragAction {
	retC := C.gdk_drag_context_get_suggested_action((*C.GdkDragContext)(recv.native))
	retGo := (DragAction)(retC)

	return retGo
}

// ListTargets is a wrapper around the C function gdk_drag_context_list_targets.
func (recv *DragContext) ListTargets() *glib.List {
	retC := C.gdk_drag_context_list_targets((*C.GdkDragContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
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

// Unsupported : gdk_screen_get_monitor_geometry : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_screen_get_monitor_workarea : unsupported parameter dest : Blacklisted record : GdkRectangle

// Unsupported : gdk_seat_grab : unsupported parameter event : no type generator for Event, const GdkEvent*

// GetBitsPerRgb is a wrapper around the C function gdk_visual_get_bits_per_rgb.
func (recv *Visual) GetBitsPerRgb() int32 {
	retC := C.gdk_visual_get_bits_per_rgb((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_visual_get_blue_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// GetByteOrder is a wrapper around the C function gdk_visual_get_byte_order.
func (recv *Visual) GetByteOrder() ByteOrder {
	retC := C.gdk_visual_get_byte_order((*C.GdkVisual)(recv.native))
	retGo := (ByteOrder)(retC)

	return retGo
}

// GetColormapSize is a wrapper around the C function gdk_visual_get_colormap_size.
func (recv *Visual) GetColormapSize() int32 {
	retC := C.gdk_visual_get_colormap_size((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDepth is a wrapper around the C function gdk_visual_get_depth.
func (recv *Visual) GetDepth() int32 {
	retC := C.gdk_visual_get_depth((*C.GdkVisual)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_visual_get_green_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// Unsupported : gdk_visual_get_red_pixel_details : unsupported parameter mask : no type generator for guint32, guint32*

// GetVisualType is a wrapper around the C function gdk_visual_get_visual_type.
func (recv *Visual) GetVisualType() VisualType {
	retC := C.gdk_visual_get_visual_type((*C.GdkVisual)(recv.native))
	retGo := (VisualType)(retC)

	return retGo
}

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_begin_paint_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_coords_from_parent : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_coords_to_parent : unsupported parameter parent_x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_create_similar_image_surface : unsupported parameter format : no type generator for gint, cairo_format_t

// CreateSimilarSurface is a wrapper around the C function gdk_window_create_similar_surface.
func (recv *Window) CreateSimilarSurface(content cairo.Content, width int32, height int32) *cairo.Surface {
	c_content := (C.cairo_content_t)(content)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_window_create_similar_surface((*C.GdkWindow)(recv.native), c_content, c_width, c_height)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAcceptFocus is a wrapper around the C function gdk_window_get_accept_focus.
func (recv *Window) GetAcceptFocus() bool {
	retC := C.gdk_window_get_accept_focus((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBackgroundPattern is a wrapper around the C function gdk_window_get_background_pattern.
func (recv *Window) GetBackgroundPattern() *cairo.Pattern {
	retC := C.gdk_window_get_background_pattern((*C.GdkWindow)(recv.native))
	retGo := cairo.PatternNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetComposited is a wrapper around the C function gdk_window_get_composited.
func (recv *Window) GetComposited() bool {
	retC := C.gdk_window_get_composited((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// Unsupported : gdk_window_get_device_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_device_position_double : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gdk_window_get_drag_protocol : unsupported parameter target : record with indirection level of 2

// GetEffectiveParent is a wrapper around the C function gdk_window_get_effective_parent.
func (recv *Window) GetEffectiveParent() *Window {
	retC := C.gdk_window_get_effective_parent((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEffectiveToplevel is a wrapper around the C function gdk_window_get_effective_toplevel.
func (recv *Window) GetEffectiveToplevel() *Window {
	retC := C.gdk_window_get_effective_toplevel((*C.GdkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusOnMap is a wrapper around the C function gdk_window_get_focus_on_map.
func (recv *Window) GetFocusOnMap() bool {
	retC := C.gdk_window_get_focus_on_map((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_get_frame_extents : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gdk_window_get_geometry : unsupported parameter x : no type generator for gint, gint*

// GetModalHint is a wrapper around the C function gdk_window_get_modal_hint.
func (recv *Window) GetModalHint() bool {
	retC := C.gdk_window_get_modal_hint((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_get_origin : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_position : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_coords : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gdk_window_get_root_origin : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gdk_window_get_user_data : unsupported parameter data : no type generator for gpointer, gpointer*

// HasNative is a wrapper around the C function gdk_window_has_native.
func (recv *Window) HasNative() bool {
	retC := C.gdk_window_has_native((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc, GdkWindowChildFunc

// Unsupported : gdk_window_invalidate_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// IsInputOnly is a wrapper around the C function gdk_window_is_input_only.
func (recv *Window) IsInputOnly() bool {
	retC := C.gdk_window_is_input_only((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsShaped is a wrapper around the C function gdk_window_is_shaped.
func (recv *Window) IsShaped() bool {
	retC := C.gdk_window_is_shaped((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc, GdkFilterFunc

// Unsupported : gdk_window_set_invalidate_handler : unsupported parameter handler : no type generator for WindowInvalidateHandlerFunc, GdkWindowInvalidateHandlerFunc

// Unsupported : gdk_window_show_window_menu : unsupported parameter event : no type generator for Event, GdkEvent*

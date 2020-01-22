// Code generated - DO NOT EDIT.
// +build gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

func Fn_gdk_cursor_get_cursor_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkCursor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_cursor_get_cursor_type(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_display_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_is_closed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_drag_context_get_actions(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_get_actions(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_drag_context_get_selected_action(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_get_selected_action(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_drag_context_get_source_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_get_source_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_context_get_suggested_action(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_get_suggested_action(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_drag_context_list_targets(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_list_targets(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

func Fn_gdk_visual_get_bits_per_rgb(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	ret := C.gdk_visual_get_bits_per_rgb(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_visual_get_blue_pixel_details(paramInstance unsafe.Pointer, param0 *uint32, param1 *int, param2 *int) {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint32)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gdk_visual_get_blue_pixel_details(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_visual_get_byte_order(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	ret := C.gdk_visual_get_byte_order(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_visual_get_colormap_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	ret := C.gdk_visual_get_colormap_size(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_visual_get_depth(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	ret := C.gdk_visual_get_depth(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_visual_get_green_pixel_details(paramInstance unsafe.Pointer, param0 *uint32, param1 *int, param2 *int) {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint32)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gdk_visual_get_green_pixel_details(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_visual_get_red_pixel_details(paramInstance unsafe.Pointer, param0 *uint32, param1 *int, param2 *int) {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint32)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gdk_visual_get_red_pixel_details(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_visual_get_visual_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	ret := C.gdk_visual_get_visual_type(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_coords_from_parent(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 *float64, param3 *float64) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	C.gdk_window_coords_from_parent(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_window_coords_to_parent(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 *float64, param3 *float64) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	C.gdk_window_coords_to_parent(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_window_create_similar_surface(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.cairo_content_t)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	ret := C.gdk_window_create_similar_surface(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_get_accept_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_accept_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_window_get_background_pattern(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_background_pattern(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_composited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_composited(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_window_get_effective_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_effective_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_effective_toplevel(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_effective_toplevel(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_focus_on_map(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_focus_on_map(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_window_get_modal_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_modal_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_window_has_native(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_has_native(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

func Fn_gdk_window_is_input_only(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_is_input_only(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_window_is_shaped(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_is_shaped(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

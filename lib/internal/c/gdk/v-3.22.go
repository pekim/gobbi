// Code generated - DO NOT EDIT.
// +build gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type EventPadAxis C.GdkEventPadAxis
type EventPadButton C.GdkEventPadButton
type EventPadGroupMode C.GdkEventPadGroupMode

func Fn_gdk_cairo_get_drawing_context(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	ret := C.gdk_cairo_get_drawing_context(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

func Fn_gdk_pango_context_get_for_display(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	ret := C.gdk_pango_context_get_for_display(cValue0)

	return unsafe.Pointer(ret)
}

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

func Fn_gdk_device_get_axes(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_axes(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_device_tool_get_hardware_id(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GdkDeviceTool)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_tool_get_hardware_id(cValueInstance)

	return (uint64)(ret)
}

func Fn_gdk_device_tool_get_serial(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GdkDeviceTool)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_tool_get_serial(cValueInstance)

	return (uint64)(ret)
}

func Fn_gdk_device_tool_get_tool_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDeviceTool)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_tool_get_tool_type(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_display_get_monitor(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	ret := C.gdk_display_get_monitor(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_monitor_at_point(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	ret := C.gdk_display_get_monitor_at_point(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_monitor_at_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	ret := C.gdk_display_get_monitor_at_window(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_n_monitors(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_n_monitors(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_display_get_primary_monitor(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_primary_monitor(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drawing_context_get_cairo_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDrawingContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drawing_context_get_cairo_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drawing_context_get_clip(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDrawingContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drawing_context_get_clip(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drawing_context_get_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDrawingContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drawing_context_get_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drawing_context_is_valid(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDrawingContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drawing_context_is_valid(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_gl_context_get_use_es(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_gl_context_get_use_es(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_gl_context_set_use_es(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	C.gdk_gl_context_set_use_es(cValueInstance, cValue0)
}

func Fn_gdk_monitor_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_monitor_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_monitor_get_geometry(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gdk_monitor_get_geometry(cValueInstance, cValue0)
}

func Fn_gdk_monitor_get_height_mm(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_monitor_get_height_mm(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_monitor_get_refresh_rate(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_monitor_get_refresh_rate(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_monitor_get_scale_factor(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_monitor_get_scale_factor(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_monitor_get_subpixel_layout(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_monitor_get_subpixel_layout(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_monitor_get_width_mm(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_monitor_get_width_mm(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_monitor_get_workarea(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gdk_monitor_get_workarea(cValueInstance, cValue0)
}

func Fn_gdk_monitor_is_primary(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_monitor_is_primary(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_begin_draw_frame(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	ret := C.gdk_window_begin_draw_frame(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_end_draw_frame(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDrawingContext)(unsafe.Pointer(param0))

	C.gdk_window_end_draw_frame(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_device_pad_get_feature_group(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.GdkDevicePad)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkDevicePadFeature)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gdk_device_pad_get_feature_group(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gdk_device_pad_get_group_n_modes(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GdkDevicePad)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gdk_device_pad_get_group_n_modes(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gdk_device_pad_get_n_features(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GdkDevicePad)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkDevicePadFeature)(param0)

	ret := C.gdk_device_pad_get_n_features(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gdk_device_pad_get_n_groups(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDevicePad)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_pad_get_n_groups(cValueInstance)

	return (int)(ret)
}

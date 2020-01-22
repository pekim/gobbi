// Code generated - DO NOT EDIT.
// +build gdk_2.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

func Fn_gdk_cairo_set_source_window(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	C.gdk_cairo_set_source_window(cValue0, cValue1, cValue2, cValue3)
}

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

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

func Fn_gdk_device_get_n_keys(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_n_keys(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_height(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_window_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_visual(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_visual(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_width(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

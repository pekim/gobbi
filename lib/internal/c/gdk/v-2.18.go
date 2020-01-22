// Code generated - DO NOT EDIT.
// +build gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

func Fn_gdk_offscreen_window_get_embedder(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	ret := C.gdk_offscreen_window_get_embedder(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_offscreen_window_set_embedder(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	C.gdk_offscreen_window_set_embedder(cValue0, cValue1)
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

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_ensure_native(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_ensure_native(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_window_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_flush(cValueInstance)
}

func Fn_gdk_window_geometry_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_geometry_changed(cValueInstance)
}

func Fn_gdk_window_get_cursor(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_cursor(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_root_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gdk_window_get_root_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

func Fn_gdk_window_is_destroyed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_is_destroyed(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_restack(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gdk_window_restack(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

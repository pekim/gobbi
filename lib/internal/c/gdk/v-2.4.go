// Code generated - DO NOT EDIT.
// +build gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

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

func Fn_gdk_cursor_new_from_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	ret := C.gdk_cursor_new_from_pixbuf(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_display_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_flush(cValueInstance)
}

func Fn_gdk_display_get_default_cursor_size(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_default_cursor_size(cValueInstance)

	return (uint)(ret)
}

func Fn_gdk_display_get_default_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_default_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_maximal_cursor_size(paramInstance unsafe.Pointer, param0 *uint, param1 *uint) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	C.gdk_display_get_maximal_cursor_size(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_display_set_double_click_distance(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gdk_display_set_double_click_distance(cValueInstance, cValue0)
}

func Fn_gdk_display_supports_cursor_alpha(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_supports_cursor_alpha(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_cursor_color(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_supports_cursor_color(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_set_accept_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_accept_focus(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_window_set_keep_above(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_keep_above(cValueInstance, cValue0)
}

func Fn_gdk_window_set_keep_below(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_keep_below(cValueInstance, cValue0)
}

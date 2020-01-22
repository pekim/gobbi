// Code generated - DO NOT EDIT.
// +build gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func Fn_gdk_cairo_surface_create_from_pixbuf(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	cValue1 := (C.int)(param1)

	cValue2 := (*C.GdkWindow)(unsafe.Pointer(param2))

	ret := C.gdk_cairo_surface_create_from_pixbuf(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

func Fn_gdk_set_allowed_backends(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gdk_set_allowed_backends(cValue0)
}

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

func Fn_gdk_cursor_new_from_surface(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_surface_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	ret := C.gdk_cursor_new_from_surface(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_get_surface(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) unsafe.Pointer {
	cValueInstance := (*C.GdkCursor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gdk_cursor_get_surface(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

func Fn_gdk_device_get_position_double(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *float64, param2 *float64) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	C.gdk_device_get_position_double(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_screen_get_monitor_scale_factor(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gdk_screen_get_monitor_scale_factor(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_create_similar_image_surface(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.cairo_format_t)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	ret := C.gdk_window_create_similar_image_surface(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_get_children_with_user_data(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	ret := C.gdk_window_get_children_with_user_data(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_device_position_double(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *float64, param2 *float64, param3 *int) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkModifierType)(unsafe.Pointer(param3))

	ret := C.gdk_window_get_device_position_double(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_scale_factor(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_scale_factor(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_window_set_opaque_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gdk_window_set_opaque_region(cValueInstance, cValue0)
}

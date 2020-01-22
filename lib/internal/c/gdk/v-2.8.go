// Code generated - DO NOT EDIT.
// +build gdk_2.8

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

type EventGrabBroken C.GdkEventGrabBroken

func Fn_gdk_cairo_create(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	ret := C.gdk_cairo_create(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cairo_rectangle(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gdk_cairo_rectangle(cValue0, cValue1)
}

func Fn_gdk_cairo_region(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_region_t)(unsafe.Pointer(param1))

	C.gdk_cairo_region(cValue0, cValue1)
}

func Fn_gdk_cairo_set_source_color(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gdk_cairo_set_source_color(cValue0, cValue1)
}

func Fn_gdk_cairo_set_source_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	C.gdk_cairo_set_source_pixbuf(cValue0, cValue1, cValue2, cValue3)
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

func Fn_gdk_cursor_new_from_name(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gdk_cursor_new_from_name(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_get_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkCursor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_cursor_get_image(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_display_warp_pointer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gdk_display_warp_pointer(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_screen_get_rgba_visual(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_rgba_visual(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

func Fn_gdk_window_move_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gdk_window_move_region(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_window_set_urgency_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_urgency_hint(cValueInstance, cValue0)
}

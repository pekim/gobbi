// Code generated - DO NOT EDIT.
// +build gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func Fn_gdk_rectangle_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GdkRectangle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gdk_rectangle_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gdk_drag_begin_from_point(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int) unsafe.Pointer {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))

	cValue2 := (*C.GList)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	ret := C.gdk_drag_begin_from_point(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_drop_done(param0 unsafe.Pointer, param1 bool) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gdk_drag_drop_done(cValue0, cValue1)
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

func Fn_gdk_device_get_seat(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_seat(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_display_get_default_seat(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_default_seat(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_list_seats(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_list_seats(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_context_get_drag_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_get_drag_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_context_manage_dnd(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) bool {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GdkDragAction)(param1)

	ret := C.gdk_drag_context_manage_dnd(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gdk_drag_context_set_hotspot(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gdk_drag_context_set_hotspot(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_gl_context_is_legacy(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_gl_context_is_legacy(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_seat_get_capabilities(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkSeat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_seat_get_capabilities(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_seat_get_keyboard(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkSeat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_seat_get_keyboard(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_seat_get_pointer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkSeat)(unsafe.Pointer(paramInstance))

	ret := C.gdk_seat_get_pointer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_seat_get_slaves(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GdkSeat)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkSeatCapabilities)(param0)

	ret := C.gdk_seat_get_slaves(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

func Fn_gdk_seat_ungrab(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkSeat)(unsafe.Pointer(paramInstance))

	C.gdk_seat_ungrab(cValueInstance)
}

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

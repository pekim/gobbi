// Code generated - DO NOT EDIT.
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func Fn_gdk_cairo_draw_from_gl(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 int, param6 int, param7 int, param8 int) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	cValue5 := (C.int)(param5)

	cValue6 := (C.int)(param6)

	cValue7 := (C.int)(param7)

	cValue8 := (C.int)(param8)

	C.gdk_cairo_draw_from_gl(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8)
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

func Fn_gdk_device_get_product_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_product_id(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_device_get_vendor_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_vendor_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gdk_gl_context_get_debug_enabled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_gl_context_get_debug_enabled(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_gl_context_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_gl_context_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_gl_context_get_forward_compatible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_gl_context_get_forward_compatible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_gl_context_get_required_version(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.int)(unsafe.Pointer(param0))

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.gdk_gl_context_get_required_version(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_gl_context_get_shared_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_gl_context_get_shared_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_gl_context_get_version(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.int)(unsafe.Pointer(param0))

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.gdk_gl_context_get_version(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_gl_context_get_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_gl_context_get_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_gl_context_make_current(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	C.gdk_gl_context_make_current(cValueInstance)
}

func Fn_gdk_gl_context_realize(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.gdk_gl_context_realize(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_gdk_gl_context_set_debug_enabled(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_gl_context_set_debug_enabled(cValueInstance, cValue0)
}

func Fn_gdk_gl_context_set_forward_compatible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_gl_context_set_forward_compatible(cValueInstance, cValue0)
}

func Fn_gdk_gl_context_set_required_version(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkGLContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	C.gdk_gl_context_set_required_version(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_gl_context_clear_current() {
	C.gdk_gl_context_clear_current()
}

func Fn_gdk_gl_context_get_current() unsafe.Pointer {
	ret := C.gdk_gl_context_get_current()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_create_gl_context(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.gdk_window_create_gl_context(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

func Fn_gdk_window_mark_paint_from_clip(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	C.gdk_window_mark_paint_from_clip(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

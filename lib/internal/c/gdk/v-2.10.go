// Code generated - DO NOT EDIT.
// +build gdk_2.10

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
func Fn_gdk_atom_intern_static_string(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gdk_atom_intern_static_string(cValue0)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_display_supports_input_shapes(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_supports_input_shapes(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_shapes(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_supports_shapes(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_screen_get_active_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_active_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_font_options(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_font_options(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_resolution(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_resolution(cValueInstance)

	return (float64)(ret)
}

func Fn_gdk_screen_get_window_stack(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_window_stack(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_is_composited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_is_composited(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_screen_set_font_options(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_font_options_t)(unsafe.Pointer(param0))

	C.gdk_screen_set_font_options(cValueInstance, cValue0)
}

func Fn_gdk_screen_set_resolution(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gdk_screen_set_resolution(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_get_type_hint(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_type_hint(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_window_input_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gdk_window_input_shape_combine_region(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

func Fn_gdk_window_merge_child_input_shapes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_merge_child_input_shapes(cValueInstance)
}

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_set_child_input_shapes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_set_child_input_shapes(cValueInstance)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

// Code generated - DO NOT EDIT.
// +build gdk_2.12

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
func Fn_gdk_color_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkColor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_color_to_string(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

func Fn_gdk_notify_startup_complete_with_id(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gdk_notify_startup_complete_with_id(cValue0)
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

func Fn_gdk_display_supports_composite(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_supports_composite(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_keymap_have_bidi_layouts(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))

	ret := C.gdk_keymap_have_bidi_layouts(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_beep(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_beep(cValueInstance)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_set_composited(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_composited(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_window_set_opacity(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gdk_window_set_opacity(cValueInstance, cValue0)
}

func Fn_gdk_window_set_startup_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gdk_window_set_startup_id(cValueInstance, cValue0)
}

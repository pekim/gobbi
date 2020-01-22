// Code generated - DO NOT EDIT.
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

func Fn_gdk_test_render_sync(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gdk_test_render_sync(cValue0)
}

func Fn_gdk_test_simulate_button(param0 unsafe.Pointer, param1 int, param2 int, param3 uint, param4 int, param5 int) bool {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.guint)(param3)

	cValue4 := (C.GdkModifierType)(param4)

	cValue5 := (C.GdkEventType)(param5)

	ret := C.gdk_test_simulate_button(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_gdk_test_simulate_key(param0 unsafe.Pointer, param1 int, param2 int, param3 uint, param4 int, param5 int) bool {
	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.guint)(param3)

	cValue4 := (C.GdkModifierType)(param4)

	cValue5 := (C.GdkEventType)(param5)

	ret := C.gdk_test_simulate_key(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

func Fn_gdk_app_launch_context_new() unsafe.Pointer {
	ret := C.gdk_app_launch_context_new()

	return unsafe.Pointer(ret)
}

func Fn_gdk_app_launch_context_set_desktop(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gdk_app_launch_context_set_desktop(cValueInstance, cValue0)
}

func Fn_gdk_app_launch_context_set_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	C.gdk_app_launch_context_set_display(cValueInstance, cValue0)
}

func Fn_gdk_app_launch_context_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gdk_app_launch_context_set_icon(cValueInstance, cValue0)
}

func Fn_gdk_app_launch_context_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gdk_app_launch_context_set_icon_name(cValueInstance, cValue0)
}

func Fn_gdk_app_launch_context_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gdk_app_launch_context_set_screen(cValueInstance, cValue0)
}

func Fn_gdk_app_launch_context_set_timestamp(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.gdk_app_launch_context_set_timestamp(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_screen_get_monitor_height_mm(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gdk_screen_get_monitor_height_mm(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_plug_name(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gdk_screen_get_monitor_plug_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gdk_screen_get_monitor_width_mm(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gdk_screen_get_monitor_width_mm(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

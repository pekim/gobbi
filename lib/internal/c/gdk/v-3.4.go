// Code generated - DO NOT EDIT.
// +build gdk_3.4

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

func Fn_gdk_keymap_get_modifier_mask(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkModifierIntent)(param0)

	ret := C.gdk_keymap_get_modifier_mask(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gdk_keymap_get_modifier_state(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))

	ret := C.gdk_keymap_get_modifier_state(cValueInstance)

	return (uint)(ret)
}

func Fn_gdk_screen_get_monitor_workarea(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gdk_screen_get_monitor_workarea(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_begin_move_drag_for_device(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 uint32) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.guint32)(param4)

	C.gdk_window_begin_move_drag_for_device(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gdk_window_begin_resize_drag_for_device(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 uint32) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkWindowEdge)(param0)

	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	cValue5 := (C.guint32)(param5)

	C.gdk_window_begin_resize_drag_for_device(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

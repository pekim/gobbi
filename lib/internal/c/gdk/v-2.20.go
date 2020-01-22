// Code generated - DO NOT EDIT.
// +build gdk_2.20

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

func Fn_gdk_device_get_axis_use(paramInstance unsafe.Pointer, param0 uint) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gdk_device_get_axis_use(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

func Fn_gdk_device_get_has_cursor(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_has_cursor(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_device_get_key(paramInstance unsafe.Pointer, param0 uint, param1 *uint, param2 *int) bool {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkModifierType)(unsafe.Pointer(param2))

	ret := C.gdk_device_get_key(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gdk_device_get_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_device_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gdk_device_get_source(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_source(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_keymap_add_virtual_modifiers(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkModifierType)(unsafe.Pointer(param0))

	C.gdk_keymap_add_virtual_modifiers(cValueInstance, cValue0)
}

func Fn_gdk_keymap_map_virtual_modifiers(paramInstance unsafe.Pointer, param0 *int) bool {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkModifierType)(unsafe.Pointer(param0))

	ret := C.gdk_keymap_map_virtual_modifiers(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gdk_screen_get_primary_monitor(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_primary_monitor(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

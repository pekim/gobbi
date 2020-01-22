// Code generated - DO NOT EDIT.
// +build gdk_2.6

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

type EventOwnerChange C.GdkEventOwnerChange

func Fn_gdk_drag_drop_succeeded(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	ret := C.gdk_drag_drop_succeeded(cValue0)

	return toGoBool(ret)
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

func Fn_gdk_display_request_selection_notification(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(param0)

	ret := C.gdk_display_request_selection_notification(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gdk_display_store_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint32, param2 []Atom, param3 int) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.guint32)(param1)

	cValue2 := (*C.GdkAtom)(unsafe.Pointer(&param2[0]))

	cValue3 := (C.gint)(param3)

	C.gdk_display_store_clipboard(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_display_supports_clipboard_persistence(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_supports_clipboard_persistence(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_display_supports_selection_notification(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_supports_selection_notification(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

func Fn_gdk_window_configure_finished(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_configure_finished(cValueInstance)
}

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_enable_synchronized_configure(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_enable_synchronized_configure(cValueInstance)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_set_focus_on_map(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_focus_on_map(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

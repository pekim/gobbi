// Code generated - DO NOT EDIT.
// +build gdk_3.0

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
func Fn_gdk_rgba_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkRGBA)(unsafe.Pointer(paramInstance))

	ret := C.gdk_rgba_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_rgba_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (C.gconstpointer)(paramInstance)

	cValue0 := (C.gconstpointer)(param0)

	ret := C.gdk_rgba_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gdk_rgba_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkRGBA)(unsafe.Pointer(paramInstance))

	C.gdk_rgba_free(cValueInstance)
}

func Fn_gdk_rgba_hash(paramInstance unsafe.Pointer) uint {
	cValueInstance := (C.gconstpointer)(paramInstance)

	ret := C.gdk_rgba_hash(cValueInstance)

	return (uint)(ret)
}

func Fn_gdk_rgba_parse(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GdkRGBA)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gdk_rgba_parse(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gdk_rgba_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkRGBA)(unsafe.Pointer(paramInstance))

	ret := C.gdk_rgba_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gdk_cairo_set_source_rgba(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gdk_cairo_set_source_rgba(cValue0, cValue1)
}

func Fn_gdk_disable_multidevice() {
	C.gdk_disable_multidevice()
}

func Fn_gdk_error_trap_pop_ignored() {
	C.gdk_error_trap_pop_ignored()
}

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

func Fn_gdk_events_get_angle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *float64) bool {
	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEvent)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	ret := C.gdk_events_get_angle(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gdk_events_get_center(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *float64, param3 *float64) bool {
	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEvent)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	ret := C.gdk_events_get_center(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gdk_events_get_distance(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *float64) bool {
	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEvent)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	ret := C.gdk_events_get_distance(cValue0, cValue1, cValue2)

	return toGoBool(ret)
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

func Fn_gdk_device_get_associated_device(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_associated_device(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

func Fn_gdk_device_get_device_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_device_type(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_device_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_device_get_n_axes(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_get_n_axes(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_device_get_position(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int, param2 *int) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gdk_device_get_position(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_device_get_window_at_position(paramInstance unsafe.Pointer, param0 *int, param1 *int) unsafe.Pointer {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gdk_device_get_window_at_position(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gdk_device_get_window_at_position_double(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) unsafe.Pointer {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gdk_device_get_window_at_position_double(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gdk_device_grab(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 bool, param3 int, param4 unsafe.Pointer, param5 uint32) int {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GdkGrabOwnership)(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.GdkEventMask)(param3)

	cValue4 := (*C.GdkCursor)(unsafe.Pointer(param4))

	cValue5 := (C.guint32)(param5)

	ret := C.gdk_device_grab(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return (int)(ret)
}

func Fn_gdk_device_list_axes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_list_axes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_device_ungrab(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.gdk_device_ungrab(cValueInstance, cValue0)
}

func Fn_gdk_device_warp(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GdkDevice)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gdk_device_warp(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_device_manager_get_client_pointer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDeviceManager)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_manager_get_client_pointer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_device_manager_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDeviceManager)(unsafe.Pointer(paramInstance))

	ret := C.gdk_device_manager_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_device_manager_list_devices(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GdkDeviceManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkDeviceType)(param0)

	ret := C.gdk_device_manager_list_devices(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_app_launch_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_app_launch_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_device_manager(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_device_manager(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_has_pending(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_has_pending(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_display_notify_startup_complete(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gdk_display_notify_startup_complete(cValueInstance, cValue0)
}

func Fn_gdk_display_manager_open_display(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplayManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gdk_display_manager_open_display(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_context_get_dest_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_get_dest_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_drag_context_get_protocol(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDragContext)(unsafe.Pointer(paramInstance))

	ret := C.gdk_drag_context_get_protocol(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_keymap_get_num_lock_state(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkKeymap)(unsafe.Pointer(paramInstance))

	ret := C.gdk_keymap_get_num_lock_state(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_get_device_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gdk_window_get_device_cursor(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gdk_window_get_device_events(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gdk_window_get_device_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int, param3 *int) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkModifierType)(unsafe.Pointer(param3))

	ret := C.gdk_window_get_device_position(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_drag_protocol(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) int {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GdkWindow)(unsafe.Pointer(param0))

	ret := C.gdk_window_get_drag_protocol(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gdk_window_get_support_multidevice(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_support_multidevice(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_set_device_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkCursor)(unsafe.Pointer(param1))

	C.gdk_window_set_device_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_set_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (C.GdkEventMask)(param1)

	C.gdk_window_set_device_events(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_window_set_source_events(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkInputSource)(param0)

	cValue1 := (C.GdkEventMask)(param1)

	C.gdk_window_set_source_events(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_window_set_support_multidevice(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_support_multidevice(cValueInstance, cValue0)
}

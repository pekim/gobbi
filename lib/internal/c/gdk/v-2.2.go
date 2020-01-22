// Code generated - DO NOT EDIT.
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

import "unsafe"

// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func Fn_gdk_drag_find_window_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 *unsafe.Pointer, param6 *int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkScreen)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	cValue5 := (**C.GdkWindow)(unsafe.Pointer(param5))

	cValue6 := (*C.GdkDragProtocol)(unsafe.Pointer(param6))

	C.gdk_drag_find_window_for_screen(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

func Fn_gdk_get_display_arg_name() string {
	ret := C.gdk_get_display_arg_name()

	return C.GoString(ret)
}

func Fn_gdk_notify_startup_complete() {
	C.gdk_notify_startup_complete()
}

func Fn_gdk_pango_context_get_for_screen(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gdk_pango_context_get_for_screen(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

func Fn_gdk_parse_args(param0 *int, param1 *[]string) {
	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	var cValue1ArrayPointer **C.gchar
	cValue1 := &cValue1ArrayPointer
	param1Indirected := *param1
	param1IndirectedLen := len(param1Indirected)
	cValue1Array := C.malloc((C.ulong)(param1IndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1IndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1IndirectedLen:param1IndirectedLen]
	for param1Indirectedi, param1IndirectedString := range param1Indirected {
		param1IndirectedSlice[param1Indirectedi] = (*C.gchar)(C.CString(param1IndirectedString))
		defer C.free(unsafe.Pointer(param1IndirectedSlice[param1Indirectedi]))
	}
	if len(param1IndirectedSlice) > 0 {
		cValue1ArrayPointer = &param1IndirectedSlice[0]
	}

	C.gdk_parse_args(cValue0, cValue1)

	param1OutLen := int(*cValue0)
	param1Out := make([]string, param1OutLen, param1OutLen)
	if param1OutLen > 0 {
		param1OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1ArrayPointer))[:param1OutLen:param1OutLen]
		for param1Outi, param1OutCString := range param1OutCSlice {
			param1Out[param1Outi] = C.GoString(param1OutCString)
		}
	}
	*param1 = param1Out
}

func Fn_gdk_selection_owner_get_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(param1)

	ret := C.gdk_selection_owner_get_for_display(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gdk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 uint32, param4 bool) bool {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(param2)

	cValue3 := (C.guint32)(param3)

	cValue4 := toCBool(param4)

	ret := C.gdk_selection_owner_set_for_display(cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gdk_selection_send_notify_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 uint32) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(param2)

	cValue3 := (C.GdkAtom)(param3)

	cValue4 := (C.GdkAtom)(param4)

	cValue5 := (C.guint32)(param5)

	C.gdk_selection_send_notify_for_display(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

func Fn_gdk_cursor_new_for_display(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.GdkCursorType)(param1)

	ret := C.gdk_cursor_new_for_display(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gdk_cursor_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkCursor)(unsafe.Pointer(paramInstance))

	ret := C.gdk_cursor_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_device_get_axis : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_axis_value : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gdk_device_get_state : parameter 'axes' is array parameter without length parameter

func Fn_gdk_display_beep(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_beep(cValueInstance)
}

func Fn_gdk_display_close(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_close(cValueInstance)
}

func Fn_gdk_display_get_default_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_default_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_event(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_event(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_n_screens(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_n_screens(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_display_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gdk_display_get_pointer(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int, param2 *int, param3 *int) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkModifierType)(unsafe.Pointer(param3))

	C.gdk_display_get_pointer(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gdk_display_get_screen(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gdk_display_get_screen(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_get_window_at_pointer(paramInstance unsafe.Pointer, param0 *int, param1 *int) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gdk_display_get_window_at_pointer(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_keyboard_ungrab(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.gdk_display_keyboard_ungrab(cValueInstance, cValue0)
}

func Fn_gdk_display_list_devices(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_list_devices(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_peek_event(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_peek_event(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_pointer_is_grabbed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_pointer_is_grabbed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_display_pointer_ungrab(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.gdk_display_pointer_ungrab(cValueInstance, cValue0)
}

func Fn_gdk_display_put_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gdk_display_put_event(cValueInstance, cValue0)
}

func Fn_gdk_display_set_double_click_time(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gdk_display_set_double_click_time(cValueInstance, cValue0)
}

func Fn_gdk_display_sync(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplay)(unsafe.Pointer(paramInstance))

	C.gdk_display_sync(cValueInstance)
}

func Fn_gdk_display_get_default() unsafe.Pointer {
	ret := C.gdk_display_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_open(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gdk_display_open(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_manager_get_default_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplayManager)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_manager_get_default_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_manager_list_displays(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkDisplayManager)(unsafe.Pointer(paramInstance))

	ret := C.gdk_display_manager_list_displays(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_display_manager_set_default_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GdkDisplayManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	C.gdk_display_manager_set_default_display(cValueInstance, cValue0)
}

func Fn_gdk_display_manager_get() unsafe.Pointer {
	ret := C.gdk_display_manager_get()

	return unsafe.Pointer(ret)
}

func Fn_gdk_keymap_get_for_display(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	ret := C.gdk_keymap_get_for_display(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_height(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_screen_get_height_mm(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_height_mm(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_at_point(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gdk_screen_get_monitor_at_point(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_at_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	ret := C.gdk_screen_get_monitor_at_window(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gdk_screen_get_monitor_geometry(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gdk_screen_get_monitor_geometry(cValueInstance, cValue0, cValue1)
}

func Fn_gdk_screen_get_n_monitors(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_n_monitors(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_screen_get_number(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_number(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_screen_get_root_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_root_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_setting(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	ret := C.gdk_screen_get_setting(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gdk_screen_get_system_visual(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_system_visual(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_toplevel_windows(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_toplevel_windows(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_get_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_width(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_screen_get_width_mm(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_get_width_mm(cValueInstance)

	return (int)(ret)
}

func Fn_gdk_screen_list_visuals(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_list_visuals(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_screen_make_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GdkScreen)(unsafe.Pointer(paramInstance))

	ret := C.gdk_screen_make_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gdk_screen_get_default() unsafe.Pointer {
	ret := C.gdk_screen_get_default()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

func Fn_gdk_visual_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkVisual)(unsafe.Pointer(paramInstance))

	ret := C.gdk_visual_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_fullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_fullscreen(cValueInstance)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

func Fn_gdk_window_set_skip_pager_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_skip_pager_hint(cValueInstance, cValue0)
}

func Fn_gdk_window_set_skip_taskbar_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gdk_window_set_skip_taskbar_hint(cValueInstance, cValue0)
}

func Fn_gdk_window_unfullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	C.gdk_window_unfullscreen(cValueInstance)
}

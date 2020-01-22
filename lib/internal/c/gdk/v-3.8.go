// Code generated - DO NOT EDIT.
// +build gdk_3.8

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
func Fn_gdk_frame_timings_get_complete(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GdkFrameTimings)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_timings_get_complete(cValueInstance)

	return toGoBool(ret)
}

func Fn_gdk_frame_timings_get_frame_counter(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GdkFrameTimings)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_timings_get_frame_counter(cValueInstance)

	return (int64)(ret)
}

func Fn_gdk_frame_timings_get_predicted_presentation_time(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GdkFrameTimings)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_timings_get_predicted_presentation_time(cValueInstance)

	return (int64)(ret)
}

func Fn_gdk_frame_timings_get_presentation_time(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GdkFrameTimings)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_timings_get_presentation_time(cValueInstance)

	return (int64)(ret)
}

func Fn_gdk_frame_timings_get_refresh_interval(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GdkFrameTimings)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_timings_get_refresh_interval(cValueInstance)

	return (int64)(ret)
}

func Fn_gdk_frame_timings_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkFrameTimings)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_timings_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_frame_timings_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkFrameTimings)(unsafe.Pointer(paramInstance))

	C.gdk_frame_timings_unref(cValueInstance)
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

func Fn_gdk_frame_clock_begin_updating(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	C.gdk_frame_clock_begin_updating(cValueInstance)
}

func Fn_gdk_frame_clock_end_updating(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	C.gdk_frame_clock_end_updating(cValueInstance)
}

func Fn_gdk_frame_clock_get_current_timings(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_clock_get_current_timings(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_frame_clock_get_frame_counter(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_clock_get_frame_counter(cValueInstance)

	return (int64)(ret)
}

func Fn_gdk_frame_clock_get_frame_time(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_clock_get_frame_time(cValueInstance)

	return (int64)(ret)
}

func Fn_gdk_frame_clock_get_history_start(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	ret := C.gdk_frame_clock_get_history_start(cValueInstance)

	return (int64)(ret)
}

func Fn_gdk_frame_clock_get_refresh_info(paramInstance unsafe.Pointer, param0 int64, param1 *int64, param2 *int64) {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint64)(param0)

	cValue1 := (*C.gint64)(unsafe.Pointer(param1))

	cValue2 := (*C.gint64)(unsafe.Pointer(param2))

	C.gdk_frame_clock_get_refresh_info(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gdk_frame_clock_get_timings(paramInstance unsafe.Pointer, param0 int64) unsafe.Pointer {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint64)(param0)

	ret := C.gdk_frame_clock_get_timings(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gdk_frame_clock_request_phase(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkFrameClock)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkFrameClockPhase)(param0)

	C.gdk_frame_clock_request_phase(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_seat_grab : parameter 'prepare_func' is callback

// UNSUPPORTED : gdk_window_add_filter : parameter 'function' is callback

// UNSUPPORTED : gdk_window_destroy_notify : blacklisted

func Fn_gdk_window_get_frame_clock(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_frame_clock(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gdk_window_get_fullscreen_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gdk_window_get_fullscreen_mode(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gdk_window_invalidate_maybe_recurse : parameter 'child_func' is callback

// UNSUPPORTED : gdk_window_remove_filter : parameter 'function' is callback

func Fn_gdk_window_set_fullscreen_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GdkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkFullscreenMode)(param0)

	C.gdk_window_set_fullscreen_mode(cValueInstance, cValue0)
}

// UNSUPPORTED : gdk_window_set_invalidate_handler : parameter 'handler' is callback

// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"C"
	"unsafe"
)

// BeginUpdating is a wrapper around the C function gdk_frame_clock_begin_updating.
func (recv *FrameClock) BeginUpdating() {
	C.gdk_frame_clock_begin_updating((*C.GdkFrameClock)(recv.native))

	return
}

// EndUpdating is a wrapper around the C function gdk_frame_clock_end_updating.
func (recv *FrameClock) EndUpdating() {
	C.gdk_frame_clock_end_updating((*C.GdkFrameClock)(recv.native))

	return
}

// GetCurrentTimings is a wrapper around the C function gdk_frame_clock_get_current_timings.
func (recv *FrameClock) GetCurrentTimings() *FrameTimings {
	retC := C.gdk_frame_clock_get_current_timings((*C.GdkFrameClock)(recv.native))
	var retGo (*FrameTimings)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FrameTimingsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFrameCounter is a wrapper around the C function gdk_frame_clock_get_frame_counter.
func (recv *FrameClock) GetFrameCounter() int64 {
	retC := C.gdk_frame_clock_get_frame_counter((*C.GdkFrameClock)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetFrameTime is a wrapper around the C function gdk_frame_clock_get_frame_time.
func (recv *FrameClock) GetFrameTime() int64 {
	retC := C.gdk_frame_clock_get_frame_time((*C.GdkFrameClock)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetHistoryStart is a wrapper around the C function gdk_frame_clock_get_history_start.
func (recv *FrameClock) GetHistoryStart() int64 {
	retC := C.gdk_frame_clock_get_history_start((*C.GdkFrameClock)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetRefreshInfo is a wrapper around the C function gdk_frame_clock_get_refresh_info.
func (recv *FrameClock) GetRefreshInfo(baseTime int64) (int64, int64) {
	c_base_time := (C.gint64)(baseTime)

	var c_refresh_interval_return C.gint64

	var c_presentation_time_return C.gint64

	C.gdk_frame_clock_get_refresh_info((*C.GdkFrameClock)(recv.native), c_base_time, &c_refresh_interval_return, &c_presentation_time_return)

	refreshIntervalReturn := (int64)(c_refresh_interval_return)

	presentationTimeReturn := (int64)(c_presentation_time_return)

	return refreshIntervalReturn, presentationTimeReturn
}

// GetTimings is a wrapper around the C function gdk_frame_clock_get_timings.
func (recv *FrameClock) GetTimings(frameCounter int64) *FrameTimings {
	c_frame_counter := (C.gint64)(frameCounter)

	retC := C.gdk_frame_clock_get_timings((*C.GdkFrameClock)(recv.native), c_frame_counter)
	var retGo (*FrameTimings)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FrameTimingsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// RequestPhase is a wrapper around the C function gdk_frame_clock_request_phase.
func (recv *FrameClock) RequestPhase(phase FrameClockPhase) {
	c_phase := (C.GdkFrameClockPhase)(phase)

	C.gdk_frame_clock_request_phase((*C.GdkFrameClock)(recv.native), c_phase)

	return
}

// GetFrameClock is a wrapper around the C function gdk_window_get_frame_clock.
func (recv *Window) GetFrameClock() *FrameClock {
	retC := C.gdk_window_get_frame_clock((*C.GdkWindow)(recv.native))
	retGo := FrameClockNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFullscreenMode is a wrapper around the C function gdk_window_get_fullscreen_mode.
func (recv *Window) GetFullscreenMode() FullscreenMode {
	retC := C.gdk_window_get_fullscreen_mode((*C.GdkWindow)(recv.native))
	retGo := (FullscreenMode)(retC)

	return retGo
}

// SetFullscreenMode is a wrapper around the C function gdk_window_set_fullscreen_mode.
func (recv *Window) SetFullscreenMode(mode FullscreenMode) {
	c_mode := (C.GdkFullscreenMode)(mode)

	C.gdk_window_set_fullscreen_mode((*C.GdkWindow)(recv.native), c_mode)

	return
}

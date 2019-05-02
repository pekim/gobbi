// Code generated - DO NOT EDIT.
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type FrameClockPhase C.GdkFrameClockPhase

const (
	GDK_FRAME_CLOCK_PHASE_NONE          FrameClockPhase = 0
	GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS  FrameClockPhase = 1
	GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT  FrameClockPhase = 2
	GDK_FRAME_CLOCK_PHASE_UPDATE        FrameClockPhase = 4
	GDK_FRAME_CLOCK_PHASE_LAYOUT        FrameClockPhase = 8
	GDK_FRAME_CLOCK_PHASE_PAINT         FrameClockPhase = 16
	GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS FrameClockPhase = 32
	GDK_FRAME_CLOCK_PHASE_AFTER_PAINT   FrameClockPhase = 64
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

type FullscreenMode C.GdkFullscreenMode

const (
	GDK_FULLSCREEN_ON_CURRENT_MONITOR FullscreenMode = 0
	GDK_FULLSCREEN_ON_ALL_MONITORS    FullscreenMode = 1
)

// GetComplete is a wrapper around the C function gdk_frame_timings_get_complete.
func (recv *FrameTimings) GetComplete() bool {
	retC := C.gdk_frame_timings_get_complete((*C.GdkFrameTimings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFrameCounter is a wrapper around the C function gdk_frame_timings_get_frame_counter.
func (recv *FrameTimings) GetFrameCounter() int64 {
	retC := C.gdk_frame_timings_get_frame_counter((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetPredictedPresentationTime is a wrapper around the C function gdk_frame_timings_get_predicted_presentation_time.
func (recv *FrameTimings) GetPredictedPresentationTime() int64 {
	retC := C.gdk_frame_timings_get_predicted_presentation_time((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetPresentationTime is a wrapper around the C function gdk_frame_timings_get_presentation_time.
func (recv *FrameTimings) GetPresentationTime() int64 {
	retC := C.gdk_frame_timings_get_presentation_time((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetRefreshInterval is a wrapper around the C function gdk_frame_timings_get_refresh_interval.
func (recv *FrameTimings) GetRefreshInterval() int64 {
	retC := C.gdk_frame_timings_get_refresh_interval((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Ref is a wrapper around the C function gdk_frame_timings_ref.
func (recv *FrameTimings) Ref() *FrameTimings {
	retC := C.gdk_frame_timings_ref((*C.GdkFrameTimings)(recv.native))
	retGo := FrameTimingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gdk_frame_timings_unref.
func (recv *FrameTimings) Unref() {
	C.gdk_frame_timings_unref((*C.GdkFrameTimings)(recv.native))

	return
}

// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Starts updates for an animation. Until a matching call to
// gdk_frame_clock_end_updating() is made, the frame clock will continually
// request a new frame with the %GDK_FRAME_CLOCK_PHASE_UPDATE phase.
// This function may be called multiple times and frames will be
// requested until gdk_frame_clock_end_updating() is called the same
// number of times.
/*

C function : gdk_frame_clock_begin_updating
*/
func (recv *FrameClock) BeginUpdating() {
	C.gdk_frame_clock_begin_updating((*C.GdkFrameClock)(recv.native))

	return
}

// Stops updates for an animation. See the documentation for
// gdk_frame_clock_begin_updating().
/*

C function : gdk_frame_clock_end_updating
*/
func (recv *FrameClock) EndUpdating() {
	C.gdk_frame_clock_end_updating((*C.GdkFrameClock)(recv.native))

	return
}

// Gets the frame timings for the current frame.
/*

C function : gdk_frame_clock_get_current_timings
*/
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

// A #GdkFrameClock maintains a 64-bit counter that increments for
// each frame drawn.
/*

C function : gdk_frame_clock_get_frame_counter
*/
func (recv *FrameClock) GetFrameCounter() int64 {
	retC := C.gdk_frame_clock_get_frame_counter((*C.GdkFrameClock)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Gets the time that should currently be used for animations.  Inside
// the processing of a frame, it’s the time used to compute the
// animation position of everything in a frame. Outside of a frame, it's
// the time of the conceptual “previous frame,” which may be either
// the actual previous frame time, or if that’s too old, an updated
// time.
/*

C function : gdk_frame_clock_get_frame_time
*/
func (recv *FrameClock) GetFrameTime() int64 {
	retC := C.gdk_frame_clock_get_frame_time((*C.GdkFrameClock)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// #GdkFrameClock internally keeps a history of #GdkFrameTimings
// objects for recent frames that can be retrieved with
// gdk_frame_clock_get_timings(). The set of stored frames
// is the set from the counter values given by
// gdk_frame_clock_get_history_start() and
// gdk_frame_clock_get_frame_counter(), inclusive.
/*

C function : gdk_frame_clock_get_history_start
*/
func (recv *FrameClock) GetHistoryStart() int64 {
	retC := C.gdk_frame_clock_get_history_start((*C.GdkFrameClock)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Using the frame history stored in the frame clock, finds the last
// known presentation time and refresh interval, and assuming that
// presentation times are separated by the refresh interval,
// predicts a presentation time that is a multiple of the refresh
// interval after the last presentation time, and later than @base_time.
/*

C function : gdk_frame_clock_get_refresh_info
*/
func (recv *FrameClock) GetRefreshInfo(baseTime int64) (int64, int64) {
	c_base_time := (C.gint64)(baseTime)

	var c_refresh_interval_return C.gint64

	var c_presentation_time_return C.gint64

	C.gdk_frame_clock_get_refresh_info((*C.GdkFrameClock)(recv.native), c_base_time, &c_refresh_interval_return, &c_presentation_time_return)

	refreshIntervalReturn := (int64)(c_refresh_interval_return)

	presentationTimeReturn := (int64)(c_presentation_time_return)

	return refreshIntervalReturn, presentationTimeReturn
}

// Retrieves a #GdkFrameTimings object holding timing information
// for the current frame or a recent frame. The #GdkFrameTimings
// object may not yet be complete: see gdk_frame_timings_get_complete().
/*

C function : gdk_frame_clock_get_timings
*/
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

// Asks the frame clock to run a particular phase. The signal
// corresponding the requested phase will be emitted the next
// time the frame clock processes. Multiple calls to
// gdk_frame_clock_request_phase() will be combined together
// and only one frame processed. If you are displaying animated
// content and want to continually request the
// %GDK_FRAME_CLOCK_PHASE_UPDATE phase for a period of time,
// you should use gdk_frame_clock_begin_updating() instead, since
// this allows GTK+ to adjust system parameters to get maximally
// smooth animations.
/*

C function : gdk_frame_clock_request_phase
*/
func (recv *FrameClock) RequestPhase(phase FrameClockPhase) {
	c_phase := (C.GdkFrameClockPhase)(phase)

	C.gdk_frame_clock_request_phase((*C.GdkFrameClock)(recv.native), c_phase)

	return
}

// Gets the frame clock for the window. The frame clock for a window
// never changes unless the window is reparented to a new toplevel
// window.
/*

C function : gdk_window_get_frame_clock
*/
func (recv *Window) GetFrameClock() *FrameClock {
	retC := C.gdk_window_get_frame_clock((*C.GdkWindow)(recv.native))
	retGo := FrameClockNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains the #GdkFullscreenMode of the @window.
/*

C function : gdk_window_get_fullscreen_mode
*/
func (recv *Window) GetFullscreenMode() FullscreenMode {
	retC := C.gdk_window_get_fullscreen_mode((*C.GdkWindow)(recv.native))
	retGo := (FullscreenMode)(retC)

	return retGo
}

// Specifies whether the @window should span over all monitors (in a multi-head
// setup) or only the current monitor when in fullscreen mode.
//
// The @mode argument is from the #GdkFullscreenMode enumeration.
// If #GDK_FULLSCREEN_ON_ALL_MONITORS is specified, the fullscreen @window will
// span over all monitors from the #GdkScreen.
//
// On X11, searches through the list of monitors from the #GdkScreen the ones
// which delimit the 4 edges of the entire #GdkScreen and will ask the window
// manager to span the @window over these monitors.
//
// If the XINERAMA extension is not available or not usable, this function
// has no effect.
//
// Not all window managers support this, so you can’t rely on the fullscreen
// window to span over the multiple monitors when #GDK_FULLSCREEN_ON_ALL_MONITORS
// is specified.
/*

C function : gdk_window_set_fullscreen_mode
*/
func (recv *Window) SetFullscreenMode(mode FullscreenMode) {
	c_mode := (C.GdkFullscreenMode)(mode)

	C.gdk_window_set_fullscreen_mode((*C.GdkWindow)(recv.native), c_mode)

	return
}

// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// The timing information in a #GdkFrameTimings is filled in
// incrementally as the frame as drawn and passed off to the
// window system for processing and display to the user. The
// accessor functions for #GdkFrameTimings can return 0 to
// indicate an unavailable value for two reasons: either because
// the information is not yet available, or because it isn't
// available at all. Once gdk_frame_timings_get_complete() returns
// %TRUE for a frame, you can be certain that no further values
// will become available and be stored in the #GdkFrameTimings.
/*

C function : gdk_frame_timings_get_complete
*/
func (recv *FrameTimings) GetComplete() bool {
	retC := C.gdk_frame_timings_get_complete((*C.GdkFrameTimings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the frame counter value of the #GdkFrameClock when this
// this frame was drawn.
/*

C function : gdk_frame_timings_get_frame_counter
*/
func (recv *FrameTimings) GetFrameCounter() int64 {
	retC := C.gdk_frame_timings_get_frame_counter((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Gets the predicted time at which this frame will be displayed. Although
// no predicted time may be available, if one is available, it will
// be available while the frame is being generated, in contrast to
// gdk_frame_timings_get_presentation_time(), which is only available
// after the frame has been presented. In general, if you are simply
// animating, you should use gdk_frame_clock_get_frame_time() rather
// than this function, but this function is useful for applications
// that want exact control over latency. For example, a movie player
// may want this information for Audio/Video synchronization.
/*

C function : gdk_frame_timings_get_predicted_presentation_time
*/
func (recv *FrameTimings) GetPredictedPresentationTime() int64 {
	retC := C.gdk_frame_timings_get_predicted_presentation_time((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Reurns the presentation time. This is the time at which the frame
// became visible to the user.
/*

C function : gdk_frame_timings_get_presentation_time
*/
func (recv *FrameTimings) GetPresentationTime() int64 {
	retC := C.gdk_frame_timings_get_presentation_time((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Gets the natural interval between presentation times for
// the display that this frame was displayed on. Frame presentation
// usually happens during the “vertical blanking interval”.
/*

C function : gdk_frame_timings_get_refresh_interval
*/
func (recv *FrameTimings) GetRefreshInterval() int64 {
	retC := C.gdk_frame_timings_get_refresh_interval((*C.GdkFrameTimings)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Increases the reference count of @timings.
/*

C function : gdk_frame_timings_ref
*/
func (recv *FrameTimings) Ref() *FrameTimings {
	retC := C.gdk_frame_timings_ref((*C.GdkFrameTimings)(recv.native))
	retGo := FrameTimingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decreases the reference count of @timings. If @timings
// is no longer referenced, it will be freed.
/*

C function : gdk_frame_timings_unref
*/
func (recv *FrameTimings) Unref() {
	C.gdk_frame_timings_unref((*C.GdkFrameTimings)(recv.native))

	return
}

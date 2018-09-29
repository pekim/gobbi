// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

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

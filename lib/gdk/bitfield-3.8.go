// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

/*
#GdkFrameClockPhase is used to represent the different paint clock
phases that can be requested. The elements of the enumeration
correspond to the signals of #GdkFrameClock.
*/
type FrameClockPhase C.GdkFrameClockPhase

const (
	// no phase
	GDK_FRAME_CLOCK_PHASE_NONE FrameClockPhase = 0
	// corresponds to GdkFrameClock::flush-events. Should not be handled by applications.
	GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS FrameClockPhase = 1
	// corresponds to GdkFrameClock::before-paint. Should not be handled by applications.
	GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT FrameClockPhase = 2
	// corresponds to GdkFrameClock::update.
	GDK_FRAME_CLOCK_PHASE_UPDATE FrameClockPhase = 4
	// corresponds to GdkFrameClock::layout.
	GDK_FRAME_CLOCK_PHASE_LAYOUT FrameClockPhase = 8
	// corresponds to GdkFrameClock::paint.
	GDK_FRAME_CLOCK_PHASE_PAINT FrameClockPhase = 16
	// corresponds to GdkFrameClock::resume-events. Should not be handled by applications.
	GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS FrameClockPhase = 32
	// corresponds to GdkFrameClock::after-paint. Should not be handled by applications.
	GDK_FRAME_CLOCK_PHASE_AFTER_PAINT FrameClockPhase = 64
)

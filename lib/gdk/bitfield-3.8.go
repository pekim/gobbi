// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

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

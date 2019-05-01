// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type FrameClockPhase int

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

type FullscreenMode int

const (
	GDK_FULLSCREEN_ON_CURRENT_MONITOR FullscreenMode = 0
	GDK_FULLSCREEN_ON_ALL_MONITORS    FullscreenMode = 1
)

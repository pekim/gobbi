// This is a generated file - DO NOT EDIT
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Indicates which monitor (in a multi-head setup) a window should span over
// when in fullscreen mode.
type FullscreenMode C.GdkFullscreenMode

const (
	// Fullscreen on current monitor only.
	GDK_FULLSCREEN_ON_CURRENT_MONITOR FullscreenMode = 0
	// Span across all monitors when fullscreen.
	GDK_FULLSCREEN_ON_ALL_MONITORS FullscreenMode = 1
)

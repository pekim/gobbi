// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

/*
Positioning hints for aligning a window relative to a rectangle.

These hints determine how the window should be positioned in the case that
the window would fall off-screen if placed in its ideal position.

For example, %GDK_ANCHOR_FLIP_X will replace %GDK_GRAVITY_NORTH_WEST with
%GDK_GRAVITY_NORTH_EAST and vice versa if the window extends beyond the left
or right edges of the monitor.

If %GDK_ANCHOR_SLIDE_X is set, the window can be shifted horizontally to fit
on-screen. If %GDK_ANCHOR_RESIZE_X is set, the window can be shrunken
horizontally to fit.

In general, when multiple flags are set, flipping should take precedence over
sliding, which should take precedence over resizing.
*/
type AnchorHints C.GdkAnchorHints

const (
	// allow flipping anchors horizontally
	GDK_ANCHOR_FLIP_X AnchorHints = 1
	// allow flipping anchors vertically
	GDK_ANCHOR_FLIP_Y AnchorHints = 2
	// allow sliding window horizontally
	GDK_ANCHOR_SLIDE_X AnchorHints = 4
	// allow sliding window vertically
	GDK_ANCHOR_SLIDE_Y AnchorHints = 8
	// allow resizing window horizontally
	GDK_ANCHOR_RESIZE_X AnchorHints = 16
	// allow resizing window vertically
	GDK_ANCHOR_RESIZE_Y AnchorHints = 32
	// allow flipping anchors on both axes
	GDK_ANCHOR_FLIP AnchorHints = 3
	// allow sliding window on both axes
	GDK_ANCHOR_SLIDE AnchorHints = 12
	// allow resizing window on both axes
	GDK_ANCHOR_RESIZE AnchorHints = 48
)

// Flags describing the current capabilities of a device/tool.
type AxisFlags C.GdkAxisFlags

const (
	// X axis is present
	GDK_AXIS_FLAG_X AxisFlags = 2
	// Y axis is present
	GDK_AXIS_FLAG_Y AxisFlags = 4
	// Pressure axis is present
	GDK_AXIS_FLAG_PRESSURE AxisFlags = 8
	// X tilt axis is present
	GDK_AXIS_FLAG_XTILT AxisFlags = 16
	// Y tilt axis is present
	GDK_AXIS_FLAG_YTILT AxisFlags = 32
	// Wheel axis is present
	GDK_AXIS_FLAG_WHEEL AxisFlags = 64
	// Distance axis is present
	GDK_AXIS_FLAG_DISTANCE AxisFlags = 128
	// Z-axis rotation is present
	GDK_AXIS_FLAG_ROTATION AxisFlags = 256
	// Slider axis is present
	GDK_AXIS_FLAG_SLIDER AxisFlags = 512
)

// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type AnchorHints C.GdkAnchorHints

const (
	GDK_ANCHOR_FLIP_X   AnchorHints = 1
	GDK_ANCHOR_FLIP_Y   AnchorHints = 2
	GDK_ANCHOR_SLIDE_X  AnchorHints = 4
	GDK_ANCHOR_SLIDE_Y  AnchorHints = 8
	GDK_ANCHOR_RESIZE_X AnchorHints = 16
	GDK_ANCHOR_RESIZE_Y AnchorHints = 32
	GDK_ANCHOR_FLIP     AnchorHints = 3
	GDK_ANCHOR_SLIDE    AnchorHints = 12
	GDK_ANCHOR_RESIZE   AnchorHints = 48
)

type AxisFlags C.GdkAxisFlags

const (
	GDK_AXIS_FLAG_X        AxisFlags = 2
	GDK_AXIS_FLAG_Y        AxisFlags = 4
	GDK_AXIS_FLAG_PRESSURE AxisFlags = 8
	GDK_AXIS_FLAG_XTILT    AxisFlags = 16
	GDK_AXIS_FLAG_YTILT    AxisFlags = 32
	GDK_AXIS_FLAG_WHEEL    AxisFlags = 64
	GDK_AXIS_FLAG_DISTANCE AxisFlags = 128
	GDK_AXIS_FLAG_ROTATION AxisFlags = 256
	GDK_AXIS_FLAG_SLIDER   AxisFlags = 512
)

// This is a generated file - DO NOT EDIT
// +build gdk_3.22
// +build gdk_3.22
// +build gdk_3.22
// +build gdk_3.22
// +build gdk_3.22
// +build gdk_3.22
// +build gdk_3.22
// +build gdk_3.22
// +build gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type AnchorHints int

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

type AxisFlags int

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

type DevicePadFeature int

const (
	GDK_DEVICE_PAD_FEATURE_BUTTON DevicePadFeature = 0
	GDK_DEVICE_PAD_FEATURE_RING   DevicePadFeature = 1
	GDK_DEVICE_PAD_FEATURE_STRIP  DevicePadFeature = 2
)

type DeviceToolType int

const (
	GDK_DEVICE_TOOL_TYPE_UNKNOWN  DeviceToolType = 0
	GDK_DEVICE_TOOL_TYPE_PEN      DeviceToolType = 1
	GDK_DEVICE_TOOL_TYPE_ERASER   DeviceToolType = 2
	GDK_DEVICE_TOOL_TYPE_BRUSH    DeviceToolType = 3
	GDK_DEVICE_TOOL_TYPE_PENCIL   DeviceToolType = 4
	GDK_DEVICE_TOOL_TYPE_AIRBRUSH DeviceToolType = 5
	GDK_DEVICE_TOOL_TYPE_MOUSE    DeviceToolType = 6
	GDK_DEVICE_TOOL_TYPE_LENS     DeviceToolType = 7
)

type SubpixelLayout int

const (
	GDK_SUBPIXEL_LAYOUT_UNKNOWN        SubpixelLayout = 0
	GDK_SUBPIXEL_LAYOUT_NONE           SubpixelLayout = 1
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB SubpixelLayout = 2
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR SubpixelLayout = 3
	GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB   SubpixelLayout = 4
	GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR   SubpixelLayout = 5
)

// Unsupported : gdk_cairo_get_drawing_context : return type :

// Unsupported : gdk_pango_context_get_for_display : return type :

// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// A pad feature.
type DevicePadFeature C.GdkDevicePadFeature

const (
	// a button
	GDK_DEVICE_PAD_FEATURE_BUTTON DevicePadFeature = 0
	// a ring-shaped interactive area
	GDK_DEVICE_PAD_FEATURE_RING DevicePadFeature = 1
	// a straight interactive area
	GDK_DEVICE_PAD_FEATURE_STRIP DevicePadFeature = 2
)

// Indicates the specific type of tool being used being a tablet. Such as an
// airbrush, pencil, etc.
type DeviceToolType C.GdkDeviceToolType

const (
	// Tool is of an unknown type.
	GDK_DEVICE_TOOL_TYPE_UNKNOWN DeviceToolType = 0
	// Tool is a standard tablet stylus.
	GDK_DEVICE_TOOL_TYPE_PEN DeviceToolType = 1
	// Tool is standard tablet eraser.
	GDK_DEVICE_TOOL_TYPE_ERASER DeviceToolType = 2
	// Tool is a brush stylus.
	GDK_DEVICE_TOOL_TYPE_BRUSH DeviceToolType = 3
	// Tool is a pencil stylus.
	GDK_DEVICE_TOOL_TYPE_PENCIL DeviceToolType = 4
	// Tool is an airbrush stylus.
	GDK_DEVICE_TOOL_TYPE_AIRBRUSH DeviceToolType = 5
	// Tool is a mouse.
	GDK_DEVICE_TOOL_TYPE_MOUSE DeviceToolType = 6
	// Tool is a lens cursor.
	GDK_DEVICE_TOOL_TYPE_LENS DeviceToolType = 7
)

// This enumeration describes how the red, green and blue components
// of physical pixels on an output device are laid out.
type SubpixelLayout C.GdkSubpixelLayout

const (
	// The layout is not known
	GDK_SUBPIXEL_LAYOUT_UNKNOWN SubpixelLayout = 0
	// Not organized in this way
	GDK_SUBPIXEL_LAYOUT_NONE SubpixelLayout = 1
	// The layout is horizontal, the order is RGB
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB SubpixelLayout = 2
	// The layout is horizontal, the order is BGR
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR SubpixelLayout = 3
	// The layout is vertical, the order is RGB
	GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB SubpixelLayout = 4
	// The layout is vertical, the order is BGR
	GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR SubpixelLayout = 5
)

// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type DevicePadFeature C.GdkDevicePadFeature

const (
	GDK_DEVICE_PAD_FEATURE_BUTTON DevicePadFeature = 0
	GDK_DEVICE_PAD_FEATURE_RING   DevicePadFeature = 1
	GDK_DEVICE_PAD_FEATURE_STRIP  DevicePadFeature = 2
)

type DeviceToolType C.GdkDeviceToolType

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

type SubpixelLayout C.GdkSubpixelLayout

const (
	GDK_SUBPIXEL_LAYOUT_UNKNOWN        SubpixelLayout = 0
	GDK_SUBPIXEL_LAYOUT_NONE           SubpixelLayout = 1
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB SubpixelLayout = 2
	GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR SubpixelLayout = 3
	GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB   SubpixelLayout = 4
	GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR   SubpixelLayout = 5
)

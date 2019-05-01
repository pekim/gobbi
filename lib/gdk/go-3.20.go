// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22
// +build gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type SeatCapabilities int

const (
	GDK_SEAT_CAPABILITY_NONE          SeatCapabilities = 0
	GDK_SEAT_CAPABILITY_POINTER       SeatCapabilities = 1
	GDK_SEAT_CAPABILITY_TOUCH         SeatCapabilities = 2
	GDK_SEAT_CAPABILITY_TABLET_STYLUS SeatCapabilities = 4
	GDK_SEAT_CAPABILITY_KEYBOARD      SeatCapabilities = 8
	GDK_SEAT_CAPABILITY_ALL_POINTING  SeatCapabilities = 7
	GDK_SEAT_CAPABILITY_ALL           SeatCapabilities = 15
)

type DragCancelReason int

const (
	GDK_DRAG_CANCEL_NO_TARGET      DragCancelReason = 0
	GDK_DRAG_CANCEL_USER_CANCELLED DragCancelReason = 1
	GDK_DRAG_CANCEL_ERROR          DragCancelReason = 2
)

// Unsupported : gdk_drag_begin_from_point : return type :

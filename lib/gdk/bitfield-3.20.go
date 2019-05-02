// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

import "C"

type SeatCapabilities C.GdkSeatCapabilities

const (
	GDK_SEAT_CAPABILITY_NONE          SeatCapabilities = 0
	GDK_SEAT_CAPABILITY_POINTER       SeatCapabilities = 1
	GDK_SEAT_CAPABILITY_TOUCH         SeatCapabilities = 2
	GDK_SEAT_CAPABILITY_TABLET_STYLUS SeatCapabilities = 4
	GDK_SEAT_CAPABILITY_KEYBOARD      SeatCapabilities = 8
	GDK_SEAT_CAPABILITY_ALL_POINTING  SeatCapabilities = 7
	GDK_SEAT_CAPABILITY_ALL           SeatCapabilities = 15
)

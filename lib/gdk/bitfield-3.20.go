// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Flags describing the seat capabilities.
type SeatCapabilities C.GdkSeatCapabilities

const (
	// No input capabilities
	GDK_SEAT_CAPABILITY_NONE SeatCapabilities = 0
	// The seat has a pointer (e.g. mouse)
	GDK_SEAT_CAPABILITY_POINTER SeatCapabilities = 1
	// The seat has touchscreen(s) attached
	GDK_SEAT_CAPABILITY_TOUCH SeatCapabilities = 2
	// The seat has drawing tablet(s) attached
	GDK_SEAT_CAPABILITY_TABLET_STYLUS SeatCapabilities = 4
	// The seat has keyboard(s) attached
	GDK_SEAT_CAPABILITY_KEYBOARD SeatCapabilities = 8
	// The union of all pointing capabilities
	GDK_SEAT_CAPABILITY_ALL_POINTING SeatCapabilities = 7
	// The union of all capabilities
	GDK_SEAT_CAPABILITY_ALL SeatCapabilities = 15
)

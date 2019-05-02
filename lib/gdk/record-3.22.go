// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Equals compares this DevicePadInterface with another DevicePadInterface, and returns true if they represent the same GObject.
func (recv *DevicePadInterface) Equals(other *DevicePadInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DrawingContextClass with another DrawingContextClass, and returns true if they represent the same GObject.
func (recv *DrawingContextClass) Equals(other *DrawingContextClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventPadAxis with another EventPadAxis, and returns true if they represent the same GObject.
func (recv *EventPadAxis) Equals(other *EventPadAxis) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventPadButton with another EventPadButton, and returns true if they represent the same GObject.
func (recv *EventPadButton) Equals(other *EventPadButton) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this EventPadGroupMode with another EventPadGroupMode, and returns true if they represent the same GObject.
func (recv *EventPadGroupMode) Equals(other *EventPadGroupMode) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this MonitorClass with another MonitorClass, and returns true if they represent the same GObject.
func (recv *MonitorClass) Equals(other *MonitorClass) bool {
	return other.ToC() == recv.ToC()
}

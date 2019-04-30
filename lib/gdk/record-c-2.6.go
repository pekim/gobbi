// This is a generated file - DO NOT EDIT
// +build gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// EventOwnerChange is a wrapper around the C record GdkEventOwnerChange.
type EventOwnerChange struct {
	native *C.GdkEventOwnerChange
	Type   EventType
	// window : record
	SendEvent int8
	// owner : record
	Reason OwnerChange
	// selection : record
	Time          uint32
	SelectionTime uint32
}

func EventOwnerChangeNewFromC(u unsafe.Pointer) *EventOwnerChange {
	c := (*C.GdkEventOwnerChange)(u)
	if c == nil {
		return nil
	}

	g := &EventOwnerChange{
		Reason:        (OwnerChange)(c.reason),
		SelectionTime: (uint32)(c.selection_time),
		SendEvent:     (int8)(c.send_event),
		Time:          (uint32)(c.time),
		Type:          (EventType)(c._type),
		native:        c,
	}

	return g
}

func (recv *EventOwnerChange) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

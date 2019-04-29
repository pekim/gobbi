// This is a generated file - DO NOT EDIT
// +build gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// EventOwnerChange is a wrapper around the C record GdkEventOwnerChange.
type EventOwnerChange struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventOwnerChange{native: u}

	return g
}

func (recv *EventOwnerChange) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

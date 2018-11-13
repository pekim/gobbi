// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Generated when a pointer or keyboard grab is broken. On X11, this happens
// when the grab window becomes unviewable (i.e. it or one of its ancestors
// is unmapped), or if the same application grabs the pointer or keyboard
// again. Note that implicit grabs (which are initiated by button presses)
// can also cause #GdkEventGrabBroken events.
/*

C record/class : GdkEventGrabBroken
*/
type EventGrabBroken struct {
	native *C.GdkEventGrabBroken
	Type   EventType
	// window : record
	SendEvent int8
	Keyboard  bool
	Implicit  bool
	// grab_window : record
}

func EventGrabBrokenNewFromC(u unsafe.Pointer) *EventGrabBroken {
	c := (*C.GdkEventGrabBroken)(u)
	if c == nil {
		return nil
	}

	g := &EventGrabBroken{
		Implicit:  c.implicit == C.TRUE,
		Keyboard:  c.keyboard == C.TRUE,
		SendEvent: (int8)(c.send_event),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventGrabBroken) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.keyboard =
		boolToGboolean(recv.Keyboard)
	recv.native.implicit =
		boolToGboolean(recv.Implicit)

	return (unsafe.Pointer)(recv.native)
}

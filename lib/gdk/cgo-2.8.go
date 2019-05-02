// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// CursorNewFromName is a wrapper around the C function gdk_cursor_new_from_name.
func CursorNewFromName(display *Display, name string) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gdk_cursor_new_from_name(c_display, c_name)
	var retGo (*Cursor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CursorNewFromC(unsafe.Pointer(retC))
	}

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// EventGrabBroken is a wrapper around the C record GdkEventGrabBroken.
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

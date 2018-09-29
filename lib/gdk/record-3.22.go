// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// DevicePadInterface is a wrapper around the C record GdkDevicePadInterface.
type DevicePadInterface struct {
	native *C.GdkDevicePadInterface
}

func DevicePadInterfaceNewFromC(u unsafe.Pointer) *DevicePadInterface {
	c := (*C.GdkDevicePadInterface)(u)
	if c == nil {
		return nil
	}

	g := &DevicePadInterface{native: c}

	return g
}

func (recv *DevicePadInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DrawingContextClass is a wrapper around the C record GdkDrawingContextClass.
type DrawingContextClass struct {
	native *C.GdkDrawingContextClass
}

func DrawingContextClassNewFromC(u unsafe.Pointer) *DrawingContextClass {
	c := (*C.GdkDrawingContextClass)(u)
	if c == nil {
		return nil
	}

	g := &DrawingContextClass{native: c}

	return g
}

func (recv *DrawingContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventPadAxis is a wrapper around the C record GdkEventPadAxis.
type EventPadAxis struct {
	native *C.GdkEventPadAxis
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Index     uint32
	Mode      uint32
	Value     float64
}

func EventPadAxisNewFromC(u unsafe.Pointer) *EventPadAxis {
	c := (*C.GdkEventPadAxis)(u)
	if c == nil {
		return nil
	}

	g := &EventPadAxis{
		Group:     (uint32)(c.group),
		Index:     (uint32)(c.index),
		Mode:      (uint32)(c.mode),
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		Value:     (float64)(c.value),
		native:    c,
	}

	return g
}

func (recv *EventPadAxis) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.group =
		(C.guint)(recv.Group)
	recv.native.index =
		(C.guint)(recv.Index)
	recv.native.mode =
		(C.guint)(recv.Mode)
	recv.native.value =
		(C.gdouble)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// EventPadButton is a wrapper around the C record GdkEventPadButton.
type EventPadButton struct {
	native *C.GdkEventPadButton
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Button    uint32
	Mode      uint32
}

func EventPadButtonNewFromC(u unsafe.Pointer) *EventPadButton {
	c := (*C.GdkEventPadButton)(u)
	if c == nil {
		return nil
	}

	g := &EventPadButton{
		Button:    (uint32)(c.button),
		Group:     (uint32)(c.group),
		Mode:      (uint32)(c.mode),
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventPadButton) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.group =
		(C.guint)(recv.Group)
	recv.native.button =
		(C.guint)(recv.Button)
	recv.native.mode =
		(C.guint)(recv.Mode)

	return (unsafe.Pointer)(recv.native)
}

// EventPadGroupMode is a wrapper around the C record GdkEventPadGroupMode.
type EventPadGroupMode struct {
	native *C.GdkEventPadGroupMode
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Mode      uint32
}

func EventPadGroupModeNewFromC(u unsafe.Pointer) *EventPadGroupMode {
	c := (*C.GdkEventPadGroupMode)(u)
	if c == nil {
		return nil
	}

	g := &EventPadGroupMode{
		Group:     (uint32)(c.group),
		Mode:      (uint32)(c.mode),
		SendEvent: (int8)(c.send_event),
		Time:      (uint32)(c.time),
		Type:      (EventType)(c._type),
		native:    c,
	}

	return g
}

func (recv *EventPadGroupMode) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.group =
		(C.guint)(recv.Group)
	recv.native.mode =
		(C.guint)(recv.Mode)

	return (unsafe.Pointer)(recv.native)
}

// MonitorClass is a wrapper around the C record GdkMonitorClass.
type MonitorClass struct {
	native *C.GdkMonitorClass
}

func MonitorClassNewFromC(u unsafe.Pointer) *MonitorClass {
	c := (*C.GdkMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &MonitorClass{native: c}

	return g
}

func (recv *MonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

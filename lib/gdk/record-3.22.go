// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import "unsafe"

// DevicePadInterface is a wrapper around the C record GdkDevicePadInterface.
type DevicePadInterface struct {
	native unsafe.Pointer
}

func DevicePadInterfaceNewFromC(u unsafe.Pointer) *DevicePadInterface {
	if u == nil {
		return nil
	}

	g := &DevicePadInterface{native: u}

	return g
}

func (recv *DevicePadInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DrawingContextClass is a wrapper around the C record GdkDrawingContextClass.
type DrawingContextClass struct {
	native unsafe.Pointer
}

func DrawingContextClassNewFromC(u unsafe.Pointer) *DrawingContextClass {
	if u == nil {
		return nil
	}

	g := &DrawingContextClass{native: u}

	return g
}

func (recv *DrawingContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventPadAxis is a wrapper around the C record GdkEventPadAxis.
type EventPadAxis struct {
	native unsafe.Pointer
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
	if u == nil {
		return nil
	}

	g := &EventPadAxis{native: u}

	return g
}

func (recv *EventPadAxis) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventPadButton is a wrapper around the C record GdkEventPadButton.
type EventPadButton struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Button    uint32
	Mode      uint32
}

func EventPadButtonNewFromC(u unsafe.Pointer) *EventPadButton {
	if u == nil {
		return nil
	}

	g := &EventPadButton{native: u}

	return g
}

func (recv *EventPadButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventPadGroupMode is a wrapper around the C record GdkEventPadGroupMode.
type EventPadGroupMode struct {
	native unsafe.Pointer
	Type   EventType
	// window : record
	SendEvent int8
	Time      uint32
	Group     uint32
	Mode      uint32
}

func EventPadGroupModeNewFromC(u unsafe.Pointer) *EventPadGroupMode {
	if u == nil {
		return nil
	}

	g := &EventPadGroupMode{native: u}

	return g
}

func (recv *EventPadGroupMode) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MonitorClass is a wrapper around the C record GdkMonitorClass.
type MonitorClass struct {
	native unsafe.Pointer
}

func MonitorClassNewFromC(u unsafe.Pointer) *MonitorClass {
	if u == nil {
		return nil
	}

	g := &MonitorClass{native: u}

	return g
}

func (recv *MonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

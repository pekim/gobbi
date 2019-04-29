// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import "unsafe"

// DevicePadInterface is a wrapper around the C record GdkDevicePadInterface.
type DevicePadInterface struct {
	native unsafe.Pointer
}

// DrawingContextClass is a wrapper around the C record GdkDrawingContextClass.
type DrawingContextClass struct {
	native unsafe.Pointer
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

// MonitorClass is a wrapper around the C record GdkMonitorClass.
type MonitorClass struct {
	native unsafe.Pointer
}

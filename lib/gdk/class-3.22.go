// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import "unsafe"

// DeviceTool is a wrapper around the C record GdkDeviceTool.
type DeviceTool struct {
	native unsafe.Pointer
}

func DeviceToolNewFromC(u unsafe.Pointer) *DeviceTool {
	if u == nil {
		return nil
	}

	g := &DeviceTool{native: u}

	return g
}

func (recv *DeviceTool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DrawingContext is a wrapper around the C record GdkDrawingContext.
type DrawingContext struct {
	native unsafe.Pointer
}

func DrawingContextNewFromC(u unsafe.Pointer) *DrawingContext {
	if u == nil {
		return nil
	}

	g := &DrawingContext{native: u}

	return g
}

func (recv *DrawingContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Monitor is a wrapper around the C record GdkMonitor.
type Monitor struct {
	native unsafe.Pointer
}

func MonitorNewFromC(u unsafe.Pointer) *Monitor {
	if u == nil {
		return nil
	}

	g := &Monitor{native: u}

	return g
}

func (recv *Monitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Seat is a wrapper around the C record GdkSeat.
type Seat struct {
	native unsafe.Pointer
	// parent_instance : record
}

func SeatNewFromC(u unsafe.Pointer) *Seat {
	if u == nil {
		return nil
	}

	g := &Seat{native: u}

	return g
}

func (recv *Seat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

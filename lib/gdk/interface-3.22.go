// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import "unsafe"

// DevicePad is a wrapper around the C record GdkDevicePad.
type DevicePad struct {
	native unsafe.Pointer
}

func DevicePadNewFromC(u unsafe.Pointer) *DevicePad {
	if u == nil {
		return nil
	}

	g := &DevicePad{native: u}

	return g
}

func (recv *DevicePad) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

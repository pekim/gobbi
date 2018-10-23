// This is a generated file - DO NOT EDIT

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// DevicePad is a wrapper around the C record GdkDevicePad.
type DevicePad struct {
	native *C.GdkDevicePad
}

func DevicePadNewFromC(u unsafe.Pointer) *DevicePad {
	c := (*C.GdkDevicePad)(u)
	if c == nil {
		return nil
	}

	g := &DevicePad{native: c}

	return g
}

func (recv *DevicePad) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

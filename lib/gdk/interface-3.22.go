// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// Equals compares this DevicePad with another DevicePad, and returns true if they represent the same GObject.
func (recv *DevicePad) Equals(other *DevicePad) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : gdk_device_pad_get_feature_group

// Blacklisted : gdk_device_pad_get_group_n_modes

// Blacklisted : gdk_device_pad_get_n_features

// Blacklisted : gdk_device_pad_get_n_groups

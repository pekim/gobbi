// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Variant is a wrapper around the C record GVariant.
type Variant struct {
	native *C.GVariant
}

func VariantNewFromC(u unsafe.Pointer) *Variant {
	c := (*C.GVariant)(u)
	if c == nil {
		return nil
	}

	g := &Variant{native: c}

	return g
}

func (recv *Variant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

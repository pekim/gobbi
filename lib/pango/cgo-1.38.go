// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_attr_background_alpha_new : return type :

// Unsupported : pango_attr_foreground_alpha_new : return type :

// AttrFontFeatures is a wrapper around the C record PangoAttrFontFeatures.
type AttrFontFeatures struct {
	native *C.PangoAttrFontFeatures
	// attr : record
	Features string
}

func AttrFontFeaturesNewFromC(u unsafe.Pointer) *AttrFontFeatures {
	c := (*C.PangoAttrFontFeatures)(u)
	if c == nil {
		return nil
	}

	g := &AttrFontFeatures{
		Features: C.GoString(c.features),
		native:   c,
	}

	return g
}

func (recv *AttrFontFeatures) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

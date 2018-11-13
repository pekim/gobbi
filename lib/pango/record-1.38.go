// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

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
	recv.native.features =
		C.CString(recv.Features)

	return (unsafe.Pointer)(recv.native)
}

// Calculates the scale factor of a matrix on the width and height of the font.
// That is, @xscale is the scale factor in the direction of the X coordinate,
// and @yscale is the scale factor in the direction perpendicular to the
// vector that the X coordinate is mapped to.
//
// Note that output numbers will always be non-negative.
/*

C function : pango_matrix_get_font_scale_factors
*/
func (recv *Matrix) GetFontScaleFactors() (float64, float64) {
	var c_xscale C.double

	var c_yscale C.double

	C.pango_matrix_get_font_scale_factors((*C.PangoMatrix)(recv.native), &c_xscale, &c_yscale)

	xscale := (float64)(c_xscale)

	yscale := (float64)(c_yscale)

	return xscale, yscale
}

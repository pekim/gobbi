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

// Equals compares this AttrFontFeatures with another AttrFontFeatures, and returns true if they represent the same GObject.
func (recv *AttrFontFeatures) Equals(other *AttrFontFeatures) bool {
	return other.ToC() == recv.ToC()
}

// AttrFontFeaturesNew is a wrapper around the C function pango_attr_font_features_new.
func AttrFontFeaturesNew(features string) *Attribute {
	c_features := C.CString(features)
	defer C.free(unsafe.Pointer(c_features))

	retC := C.pango_attr_font_features_new(c_features)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontScaleFactors is a wrapper around the C function pango_matrix_get_font_scale_factors.
func (recv *Matrix) GetFontScaleFactors() (float64, float64) {
	var c_xscale C.double

	var c_yscale C.double

	C.pango_matrix_get_font_scale_factors((*C.PangoMatrix)(recv.native), &c_xscale, &c_yscale)

	xscale := (float64)(c_xscale)

	yscale := (float64)(c_yscale)

	return xscale, yscale
}

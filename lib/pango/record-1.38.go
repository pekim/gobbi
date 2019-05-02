// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

import (
	"C"
	"unsafe"
)

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

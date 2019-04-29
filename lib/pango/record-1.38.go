// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

import "unsafe"

// AttrFontFeatures is a wrapper around the C record PangoAttrFontFeatures.
type AttrFontFeatures struct {
	native unsafe.Pointer
	// attr : record
	Features string
}

func AttrFontFeaturesNewFromC(u unsafe.Pointer) *AttrFontFeatures {
	if u == nil {
		return nil
	}

	g := &AttrFontFeatures{native: u}

	return g
}

func (recv *AttrFontFeatures) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// GlyphItemIter is a wrapper around the C record PangoGlyphItemIter.
type GlyphItemIter struct {
	native unsafe.Pointer
	// glyph_item : record
	Text       string
	StartGlyph int32
	StartIndex int32
	StartChar  int32
	EndGlyph   int32
	EndIndex   int32
	EndChar    int32
}

func GlyphItemIterNewFromC(u unsafe.Pointer) *GlyphItemIter {
	if u == nil {
		return nil
	}

	g := &GlyphItemIter{native: u}

	return g
}

func (recv *GlyphItemIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

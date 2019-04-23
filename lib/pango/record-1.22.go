// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GlyphItemIter is a wrapper around the C record PangoGlyphItemIter.
type GlyphItemIter struct {
	native *C.PangoGlyphItemIter
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
	c := (*C.PangoGlyphItemIter)(u)
	if c == nil {
		return nil
	}

	g := &GlyphItemIter{
		EndChar:    (int32)(c.end_char),
		EndGlyph:   (int32)(c.end_glyph),
		EndIndex:   (int32)(c.end_index),
		StartChar:  (int32)(c.start_char),
		StartGlyph: (int32)(c.start_glyph),
		StartIndex: (int32)(c.start_index),
		Text:       C.GoString(c.text),
		native:     c,
	}

	return g
}

func (recv *GlyphItemIter) ToC() unsafe.Pointer {
	recv.native.text =
		C.CString(recv.Text)
	recv.native.start_glyph =
		(C.int)(recv.StartGlyph)
	recv.native.start_index =
		(C.int)(recv.StartIndex)
	recv.native.start_char =
		(C.int)(recv.StartChar)
	recv.native.end_glyph =
		(C.int)(recv.EndGlyph)
	recv.native.end_index =
		(C.int)(recv.EndIndex)
	recv.native.end_char =
		(C.int)(recv.EndChar)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GlyphItemIter with another GlyphItemIter, and returns true if they represent the same GObject.
func (recv *GlyphItemIter) Equals(other *GlyphItemIter) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : pango_glyph_item_iter_copy

// Blacklisted : pango_glyph_item_iter_free

// Blacklisted : pango_glyph_item_iter_init_end

// Blacklisted : pango_glyph_item_iter_init_start

// Blacklisted : pango_glyph_item_iter_next_cluster

// Blacklisted : pango_glyph_item_iter_prev_cluster

// Unsupported : pango_language_get_scripts : array return type :

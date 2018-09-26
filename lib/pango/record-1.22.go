// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.31.0 pango_1.32 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
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

// Copy is a wrapper around the C function pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	retC := C.pango_glyph_item_iter_copy((*C.PangoGlyphItemIter)(recv.native))
	retGo := GlyphItemIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_glyph_item_iter_free : no return generator

// InitEnd is a wrapper around the C function pango_glyph_item_iter_init_end.
func (recv *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(glyphItem.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.pango_glyph_item_iter_init_end((*C.PangoGlyphItemIter)(recv.native), c_glyph_item, c_text)
	retGo := retC == C.TRUE

	return retGo
}

// InitStart is a wrapper around the C function pango_glyph_item_iter_init_start.
func (recv *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(glyphItem.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.pango_glyph_item_iter_init_start((*C.PangoGlyphItemIter)(recv.native), c_glyph_item, c_text)
	retGo := retC == C.TRUE

	return retGo
}

// NextCluster is a wrapper around the C function pango_glyph_item_iter_next_cluster.
func (recv *GlyphItemIter) NextCluster() bool {
	retC := C.pango_glyph_item_iter_next_cluster((*C.PangoGlyphItemIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PrevCluster is a wrapper around the C function pango_glyph_item_iter_prev_cluster.
func (recv *GlyphItemIter) PrevCluster() bool {
	retC := C.pango_glyph_item_iter_prev_cluster((*C.PangoGlyphItemIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

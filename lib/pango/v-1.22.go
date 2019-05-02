// Code generated - DO NOT EDIT.
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// CreateContext is a wrapper around the C function pango_font_map_create_context.
func (recv *FontMap) CreateContext() *Context {
	retC := C.pango_font_map_create_context((*C.PangoFontMap)(recv.native))
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBaseline is a wrapper around the C function pango_layout_get_baseline.
func (recv *Layout) GetBaseline() int32 {
	retC := C.pango_layout_get_baseline((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// DrawGlyphItem is a wrapper around the C function pango_renderer_draw_glyph_item.
func (recv *Renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x int32, y int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_glyph_item := (*C.PangoGlyphItem)(C.NULL)
	if glyphItem != nil {
		c_glyph_item = (*C.PangoGlyphItem)(glyphItem.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_glyph_item((*C.PangoRenderer)(recv.native), c_text, c_glyph_item, c_x, c_y)

	return
}

type BidiType C.PangoBidiType

const (
	PANGO_BIDI_TYPE_L   BidiType = 0
	PANGO_BIDI_TYPE_LRE BidiType = 1
	PANGO_BIDI_TYPE_LRO BidiType = 2
	PANGO_BIDI_TYPE_R   BidiType = 3
	PANGO_BIDI_TYPE_AL  BidiType = 4
	PANGO_BIDI_TYPE_RLE BidiType = 5
	PANGO_BIDI_TYPE_RLO BidiType = 6
	PANGO_BIDI_TYPE_PDF BidiType = 7
	PANGO_BIDI_TYPE_EN  BidiType = 8
	PANGO_BIDI_TYPE_ES  BidiType = 9
	PANGO_BIDI_TYPE_ET  BidiType = 10
	PANGO_BIDI_TYPE_AN  BidiType = 11
	PANGO_BIDI_TYPE_CS  BidiType = 12
	PANGO_BIDI_TYPE_NSM BidiType = 13
	PANGO_BIDI_TYPE_BN  BidiType = 14
	PANGO_BIDI_TYPE_B   BidiType = 15
	PANGO_BIDI_TYPE_S   BidiType = 16
	PANGO_BIDI_TYPE_WS  BidiType = 17
	PANGO_BIDI_TYPE_ON  BidiType = 18
)

// BidiTypeForUnichar is a wrapper around the C function pango_bidi_type_for_unichar.
func BidiTypeForUnichar(ch rune) BidiType {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_bidi_type_for_unichar(c_ch)
	retGo := (BidiType)(retC)

	return retGo
}

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

// Copy is a wrapper around the C function pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	retC := C.pango_glyph_item_iter_copy((*C.PangoGlyphItemIter)(recv.native))
	var retGo (*GlyphItemIter)
	if retC == nil {
		retGo = nil
	} else {
		retGo = GlyphItemIterNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Free is a wrapper around the C function pango_glyph_item_iter_free.
func (recv *GlyphItemIter) Free() {
	C.pango_glyph_item_iter_free((*C.PangoGlyphItemIter)(recv.native))

	return
}

// InitEnd is a wrapper around the C function pango_glyph_item_iter_init_end.
func (recv *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(C.NULL)
	if glyphItem != nil {
		c_glyph_item = (*C.PangoGlyphItem)(glyphItem.ToC())
	}

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.pango_glyph_item_iter_init_end((*C.PangoGlyphItemIter)(recv.native), c_glyph_item, c_text)
	retGo := retC == C.TRUE

	return retGo
}

// InitStart is a wrapper around the C function pango_glyph_item_iter_init_start.
func (recv *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	c_glyph_item := (*C.PangoGlyphItem)(C.NULL)
	if glyphItem != nil {
		c_glyph_item = (*C.PangoGlyphItem)(glyphItem.ToC())
	}

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

// Unsupported : pango_language_get_scripts : array return type :

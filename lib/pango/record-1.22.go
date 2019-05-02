// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	"C"
	"unsafe"
)

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

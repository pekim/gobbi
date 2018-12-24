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

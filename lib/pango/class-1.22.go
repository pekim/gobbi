// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Creates a #PangoContext connected to @fontmap.  This is equivalent
// to pango_context_new() followed by pango_context_set_font_map().
//
// If you are using Pango as part of a higher-level system,
// that system may have it's own way of create a #PangoContext.
// For instance, the GTK+ toolkit has, among others,
// gdk_pango_context_get_for_screen(), and
// gtk_widget_get_pango_context().  Use those instead.
/*

C function

pango_font_map_create_context
*/
func (recv *FontMap) CreateContext() *Context {
	retC := C.pango_font_map_create_context((*C.PangoFontMap)(recv.native))
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the Y position of baseline of the first line in @layout.
/*

C function

pango_layout_get_baseline
*/
func (recv *Layout) GetBaseline() int32 {
	retC := C.pango_layout_get_baseline((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Draws the glyphs in @glyph_item with the specified #PangoRenderer,
// embedding the text associated with the glyphs in the output if the
// output format supports it (PDF for example).
//
// Note that @text is the start of the text for layout, which is then
// indexed by <literal>@glyph_item->item->offset</literal>.
//
// If @text is %NULL, this simply calls pango_renderer_draw_glyphs().
//
// The default implementation of this method simply falls back to
// pango_renderer_draw_glyphs().
/*

C function

pango_renderer_draw_glyph_item
*/
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

// Code generated - DO NOT EDIT.
// +build pangocairo_1.22

package pangocairo

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
import "C"

// CreateContext is a wrapper around the C function pango_cairo_create_context.
func CreateContext(cr *cairo.Context) *pango.Context {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	retC := C.pango_cairo_create_context(c_cr)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ShowGlyphItem is a wrapper around the C function pango_cairo_show_glyph_item.
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_glyph_item := (*C.PangoGlyphItem)(C.NULL)
	if glyphItem != nil {
		c_glyph_item = (*C.PangoGlyphItem)(glyphItem.ToC())
	}

	C.pango_cairo_show_glyph_item(c_cr, c_text, c_glyph_item)

	return
}

// SetDefault is a wrapper around the C function pango_cairo_font_map_set_default.
func (recv *FontMap) SetDefault() {
	C.pango_cairo_font_map_set_default((*C.PangoCairoFontMap)(recv.native))

	return
}

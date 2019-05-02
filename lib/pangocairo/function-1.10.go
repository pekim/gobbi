// This is a generated file - DO NOT EDIT
// +build pangocairo_1.10 pangocairo_1.14 pangocairo_1.18 pangocairo_1.22

package pangocairo

import (
	"C"
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// ContextGetFontOptions is a wrapper around the C function pango_cairo_context_get_font_options.
func ContextGetFontOptions(context *pango.Context) *cairo.FontOptions {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	retC := C.pango_cairo_context_get_font_options(c_context)
	var retGo (*cairo.FontOptions)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.FontOptionsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ContextGetResolution is a wrapper around the C function pango_cairo_context_get_resolution.
func ContextGetResolution(context *pango.Context) float64 {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	retC := C.pango_cairo_context_get_resolution(c_context)
	retGo := (float64)(retC)

	return retGo
}

// ContextSetFontOptions is a wrapper around the C function pango_cairo_context_set_font_options.
func ContextSetFontOptions(context *pango.Context, options *cairo.FontOptions) {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_options := (*C.cairo_font_options_t)(C.NULL)
	if options != nil {
		c_options = (*C.cairo_font_options_t)(options.ToC())
	}

	C.pango_cairo_context_set_font_options(c_context, c_options)

	return
}

// ContextSetResolution is a wrapper around the C function pango_cairo_context_set_resolution.
func ContextSetResolution(context *pango.Context, dpi float64) {
	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	c_dpi := (C.double)(dpi)

	C.pango_cairo_context_set_resolution(c_context, c_dpi)

	return
}

// CreateLayout is a wrapper around the C function pango_cairo_create_layout.
func CreateLayout(cr *cairo.Context) *pango.Layout {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	retC := C.pango_cairo_create_layout(c_cr)
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_cairo_glyph_string_path : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// LayoutLinePath is a wrapper around the C function pango_cairo_layout_line_path.
func LayoutLinePath(cr *cairo.Context, line *pango.LayoutLine) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	C.pango_cairo_layout_line_path(c_cr, c_line)

	return
}

// LayoutPath is a wrapper around the C function pango_cairo_layout_path.
func LayoutPath(cr *cairo.Context, layout *pango.Layout) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	C.pango_cairo_layout_path(c_cr, c_layout)

	return
}

// Unsupported : pango_cairo_show_glyph_string : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// ShowLayout is a wrapper around the C function pango_cairo_show_layout.
func ShowLayout(cr *cairo.Context, layout *pango.Layout) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	C.pango_cairo_show_layout(c_cr, c_layout)

	return
}

// ShowLayoutLine is a wrapper around the C function pango_cairo_show_layout_line.
func ShowLayoutLine(cr *cairo.Context, line *pango.LayoutLine) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	C.pango_cairo_show_layout_line(c_cr, c_line)

	return
}

// UpdateContext is a wrapper around the C function pango_cairo_update_context.
func UpdateContext(cr *cairo.Context, context *pango.Context) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_context := (*C.PangoContext)(C.NULL)
	if context != nil {
		c_context = (*C.PangoContext)(context.ToC())
	}

	C.pango_cairo_update_context(c_cr, c_context)

	return
}

// UpdateLayout is a wrapper around the C function pango_cairo_update_layout.
func UpdateLayout(cr *cairo.Context, layout *pango.Layout) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	C.pango_cairo_update_layout(c_cr, c_layout)

	return
}

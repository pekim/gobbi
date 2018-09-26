// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.16 pango_1.22 pango_1.26 pango_1.31.0 pango_1.32 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Renderer is a wrapper around the C record PangoRenderer.
type Renderer struct {
	native *C.PangoRenderer
	// Private : parent_instance
	// Private : underline
	// Private : strikethrough
	// Private : active_count
	// matrix : record
	// Private : priv
}

func RendererNewFromC(u unsafe.Pointer) *Renderer {
	c := (*C.PangoRenderer)(u)
	if c == nil {
		return nil
	}

	g := &Renderer{native: c}

	return g
}

func (recv *Renderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_renderer_activate : no return generator

// Unsupported : pango_renderer_deactivate : no return generator

// Unsupported : pango_renderer_draw_error_underline : no return generator

// Unsupported : pango_renderer_draw_glyph : no return generator

// Unsupported : pango_renderer_draw_glyph_item : no return generator

// Unsupported : pango_renderer_draw_glyphs : no return generator

// Unsupported : pango_renderer_draw_layout : no return generator

// Unsupported : pango_renderer_draw_layout_line : no return generator

// Unsupported : pango_renderer_draw_rectangle : no return generator

// Unsupported : pango_renderer_draw_trapezoid : no return generator

// GetColor is a wrapper around the C function pango_renderer_get_color.
func (recv *Renderer) GetColor(part RenderPart) *Color {
	c_part := (C.PangoRenderPart)(part)

	retC := C.pango_renderer_get_color((*C.PangoRenderer)(recv.native), c_part)
	retGo := ColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMatrix is a wrapper around the C function pango_renderer_get_matrix.
func (recv *Renderer) GetMatrix() *Matrix {
	retC := C.pango_renderer_get_matrix((*C.PangoRenderer)(recv.native))
	retGo := MatrixNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_renderer_part_changed : no return generator

// Unsupported : pango_renderer_set_alpha : no return generator

// Unsupported : pango_renderer_set_color : no return generator

// Unsupported : pango_renderer_set_matrix : no return generator

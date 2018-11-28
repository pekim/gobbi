// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetFontDescription is a wrapper around the C function pango_layout_get_font_description.
func (recv *Layout) GetFontDescription() *FontDescription {
	retC := C.pango_layout_get_font_description((*C.PangoLayout)(recv.native))
	var retGo (*FontDescription)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontDescriptionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

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

// Equals compares this Renderer with another Renderer, and returns true if they represent the same GObject.
func (recv *Renderer) Equals(other *Renderer) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Renderer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Renderer.
// Exercise care, as this is a potentially dangerous function if the Object is not a Renderer.
func CastToRenderer(object *gobject.Object) *Renderer {
	return RendererNewFromC(object.ToC())
}

// Activate is a wrapper around the C function pango_renderer_activate.
func (recv *Renderer) Activate() {
	C.pango_renderer_activate((*C.PangoRenderer)(recv.native))

	return
}

// Deactivate is a wrapper around the C function pango_renderer_deactivate.
func (recv *Renderer) Deactivate() {
	C.pango_renderer_deactivate((*C.PangoRenderer)(recv.native))

	return
}

// DrawErrorUnderline is a wrapper around the C function pango_renderer_draw_error_underline.
func (recv *Renderer) DrawErrorUnderline(x int32, y int32, width int32, height int32) {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_error_underline((*C.PangoRenderer)(recv.native), c_x, c_y, c_width, c_height)

	return
}

// DrawGlyph is a wrapper around the C function pango_renderer_draw_glyph.
func (recv *Renderer) DrawGlyph(font *Font, glyph Glyph, x float64, y float64) {
	c_font := (*C.PangoFont)(C.NULL)
	if font != nil {
		c_font = (*C.PangoFont)(font.ToC())
	}

	c_glyph := (C.PangoGlyph)(glyph)

	c_x := (C.double)(x)

	c_y := (C.double)(y)

	C.pango_renderer_draw_glyph((*C.PangoRenderer)(recv.native), c_font, c_glyph, c_x, c_y)

	return
}

// Unsupported : pango_renderer_draw_glyphs : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// DrawLayout is a wrapper around the C function pango_renderer_draw_layout.
func (recv *Renderer) DrawLayout(layout *Layout, x int32, y int32) {
	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_layout((*C.PangoRenderer)(recv.native), c_layout, c_x, c_y)

	return
}

// DrawLayoutLine is a wrapper around the C function pango_renderer_draw_layout_line.
func (recv *Renderer) DrawLayoutLine(line *LayoutLine, x int32, y int32) {
	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_renderer_draw_layout_line((*C.PangoRenderer)(recv.native), c_line, c_x, c_y)

	return
}

// DrawRectangle is a wrapper around the C function pango_renderer_draw_rectangle.
func (recv *Renderer) DrawRectangle(part RenderPart, x int32, y int32, width int32, height int32) {
	c_part := (C.PangoRenderPart)(part)

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_rectangle((*C.PangoRenderer)(recv.native), c_part, c_x, c_y, c_width, c_height)

	return
}

// DrawTrapezoid is a wrapper around the C function pango_renderer_draw_trapezoid.
func (recv *Renderer) DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	c_part := (C.PangoRenderPart)(part)

	c_y1_ := (C.double)(y1)

	c_x11 := (C.double)(x11)

	c_x21 := (C.double)(x21)

	c_y2 := (C.double)(y2)

	c_x12 := (C.double)(x12)

	c_x22 := (C.double)(x22)

	C.pango_renderer_draw_trapezoid((*C.PangoRenderer)(recv.native), c_part, c_y1_, c_x11, c_x21, c_y2, c_x12, c_x22)

	return
}

// GetColor is a wrapper around the C function pango_renderer_get_color.
func (recv *Renderer) GetColor(part RenderPart) *Color {
	c_part := (C.PangoRenderPart)(part)

	retC := C.pango_renderer_get_color((*C.PangoRenderer)(recv.native), c_part)
	var retGo (*Color)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ColorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetMatrix is a wrapper around the C function pango_renderer_get_matrix.
func (recv *Renderer) GetMatrix() *Matrix {
	retC := C.pango_renderer_get_matrix((*C.PangoRenderer)(recv.native))
	var retGo (*Matrix)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MatrixNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PartChanged is a wrapper around the C function pango_renderer_part_changed.
func (recv *Renderer) PartChanged(part RenderPart) {
	c_part := (C.PangoRenderPart)(part)

	C.pango_renderer_part_changed((*C.PangoRenderer)(recv.native), c_part)

	return
}

// SetColor is a wrapper around the C function pango_renderer_set_color.
func (recv *Renderer) SetColor(part RenderPart, color *Color) {
	c_part := (C.PangoRenderPart)(part)

	c_color := (*C.PangoColor)(C.NULL)
	if color != nil {
		c_color = (*C.PangoColor)(color.ToC())
	}

	C.pango_renderer_set_color((*C.PangoRenderer)(recv.native), c_part, c_color)

	return
}

// SetMatrix is a wrapper around the C function pango_renderer_set_matrix.
func (recv *Renderer) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	C.pango_renderer_set_matrix((*C.PangoRenderer)(recv.native), c_matrix)

	return
}

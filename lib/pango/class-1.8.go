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

// Gets the font description for the layout, if any.
/*

C function : pango_layout_get_font_description
*/
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

// Object upcasts to *Object
func (recv *Renderer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Renderer.
// Exercise care, as this is a potentially dangerous function if the Object is not a Renderer.
func CastToRenderer(object *gobject.Object) *Renderer {
	return RendererNewFromC(object.ToC())
}

// Does initial setup before rendering operations on @renderer.
// pango_renderer_deactivate() should be called when done drawing.
// Calls such as pango_renderer_draw_layout() automatically
// activate the layout before drawing on it. Calls to
// pango_renderer_activate() and pango_renderer_deactivate() can
// be nested and the renderer will only be initialized and
// deinitialized once.
/*

C function : pango_renderer_activate
*/
func (recv *Renderer) Activate() {
	C.pango_renderer_activate((*C.PangoRenderer)(recv.native))

	return
}

// Cleans up after rendering operations on @renderer. See
// docs for pango_renderer_activate().
/*

C function : pango_renderer_deactivate
*/
func (recv *Renderer) Deactivate() {
	C.pango_renderer_deactivate((*C.PangoRenderer)(recv.native))

	return
}

// Draw a squiggly line that approximately covers the given rectangle
// in the style of an underline used to indicate a spelling error.
// (The width of the underline is rounded to an integer number
// of up/down segments and the resulting rectangle is centered
// in the original rectangle)
//
// This should be called while @renderer is already active.  Use
// pango_renderer_activate() to activate a renderer.
/*

C function : pango_renderer_draw_error_underline
*/
func (recv *Renderer) DrawErrorUnderline(x int32, y int32, width int32, height int32) {
	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_error_underline((*C.PangoRenderer)(recv.native), c_x, c_y, c_width, c_height)

	return
}

// Draws a single glyph with coordinates in device space.
/*

C function : pango_renderer_draw_glyph
*/
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

// Draws @layout with the specified #PangoRenderer.
/*

C function : pango_renderer_draw_layout
*/
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

// Draws @line with the specified #PangoRenderer.
/*

C function : pango_renderer_draw_layout_line
*/
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

// Draws an axis-aligned rectangle in user space coordinates with the
// specified #PangoRenderer.
//
// This should be called while @renderer is already active.  Use
// pango_renderer_activate() to activate a renderer.
/*

C function : pango_renderer_draw_rectangle
*/
func (recv *Renderer) DrawRectangle(part RenderPart, x int32, y int32, width int32, height int32) {
	c_part := (C.PangoRenderPart)(part)

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.pango_renderer_draw_rectangle((*C.PangoRenderer)(recv.native), c_part, c_x, c_y, c_width, c_height)

	return
}

// Draws a trapezoid with the parallel sides aligned with the X axis
// using the given #PangoRenderer; coordinates are in device space.
/*

C function : pango_renderer_draw_trapezoid
*/
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

// Gets the current rendering color for the specified part.
/*

C function : pango_renderer_get_color
*/
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

// Gets the transformation matrix that will be applied when
// rendering. See pango_renderer_set_matrix().
/*

C function : pango_renderer_get_matrix
*/
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

// Informs Pango that the way that the rendering is done
// for @part has changed in a way that would prevent multiple
// pieces being joined together into one drawing call. For
// instance, if a subclass of #PangoRenderer was to add a stipple
// option for drawing underlines, it needs to call
//
// <informalexample><programlisting>
// pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE);
// </programlisting></informalexample>
//
// When the stipple changes or underlines with different stipples
// might be joined together. Pango automatically calls this for
// changes to colors. (See pango_renderer_set_color())
/*

C function : pango_renderer_part_changed
*/
func (recv *Renderer) PartChanged(part RenderPart) {
	c_part := (C.PangoRenderPart)(part)

	C.pango_renderer_part_changed((*C.PangoRenderer)(recv.native), c_part)

	return
}

// Sets the color for part of the rendering.
// Also see pango_renderer_set_alpha().
/*

C function : pango_renderer_set_color
*/
func (recv *Renderer) SetColor(part RenderPart, color *Color) {
	c_part := (C.PangoRenderPart)(part)

	c_color := (*C.PangoColor)(C.NULL)
	if color != nil {
		c_color = (*C.PangoColor)(color.ToC())
	}

	C.pango_renderer_set_color((*C.PangoRenderer)(recv.native), c_part, c_color)

	return
}

// Sets the transformation matrix that will be applied when rendering.
/*

C function : pango_renderer_set_matrix
*/
func (recv *Renderer) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	C.pango_renderer_set_matrix((*C.PangoRenderer)(recv.native), c_matrix)

	return
}

// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-Paths.html.

func (ctx *Context) CopyPath() *Path {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_copy_path(c_ctx)
	retGo := PathNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Path) {
		o.Destroy()
	})

	return retGo
}

func (ctx *Context) CopyPathFlat() *Path {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_copy_path_flat(c_ctx)
	retGo := PathNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Path) {
		o.Destroy()
	})

	return retGo
}

func (path *Path) Destroy() {
	c_path := (*C.cairo_path_t)(path.ToC())
	C.cairo_path_destroy(c_path)
}

func (ctx *Context) AppendPath(path *Path) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_path := (*C.cairo_path_t)(path.ToC())
	C.cairo_append_path(c_ctx, c_path)
}

func (ctx *Context) GetCurrentPoint() (float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_x C.double
	var c_y C.double

	C.cairo_get_current_point(c_ctx, &c_x, &c_y)

	return float64(c_x), float64(c_y)
}

func (ctx *Context) NewPath() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_new_path(c_ctx)
}

func (ctx *Context) ClosePath() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_close_path(c_ctx)
}

func (ctx *Context) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_xc := (C.double)(xc)
	c_yc := (C.double)(yc)
	c_radius := (C.double)(radius)
	c_angle1 := (C.double)(angle1)
	c_angle2 := (C.double)(angle2)

	C.cairo_arc(c_ctx, c_xc, c_yc, c_radius, c_angle1, c_angle2)
}

func (ctx *Context) ArcNegative(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_xc := (C.double)(xc)
	c_yc := (C.double)(yc)
	c_radius := (C.double)(radius)
	c_angle1 := (C.double)(angle1)
	c_angle2 := (C.double)(angle2)

	C.cairo_arc_negative(c_ctx, c_xc, c_yc, c_radius, c_angle1, c_angle2)
}

func (ctx *Context) CurveTo(
	x1 float64, y1 float64,
	x2 float64, y2 float64,
	x3 float64, y3 float64,
) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x1 := (C.double)(x1)
	c_y1 := (C.double)(y1)
	c_x2 := (C.double)(x2)
	c_y2 := (C.double)(y2)
	c_x3 := (C.double)(x3)
	c_y3 := (C.double)(y3)

	C.cairo_curve_to(c_ctx, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3)
}

func (ctx *Context) LineTo(x float64, y float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	C.cairo_line_to(c_ctx, c_x, c_y)
}

func (ctx *Context) MoveTo(x float64, y float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	C.cairo_move_to(c_ctx, c_x, c_y)
}

func (ctx *Context) Rectangle(x float64, y float64, width float64, height float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)
	c_width := (C.double)(width)
	c_height := (C.double)(height)

	C.cairo_rectangle(c_ctx, c_x, c_y, c_width, c_height)
}

func (ctx *Context) GlyphPath(glyphs []Glyph) {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	numGlyphs := len(glyphs)
	c_numGlyphs := C.int(numGlyphs)

	c_glyphs := make([]C.cairo_glyph_t, numGlyphs, numGlyphs)
	for i, glyph := range glyphs {
		c_glyphs[i] = glyph.toC()
	}

	C.cairo_glyph_path(c_ctx, &c_glyphs[0], c_numGlyphs)
}

func (ctx *Context) TextPath(text string) {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.cairo_text_path(c_ctx, c_text)
}

func (ctx *Context) RelCurveTo(
	x1 float64, y1 float64,
	x2 float64, y2 float64,
	x3 float64, y3 float64,
) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x1 := (C.double)(x1)
	c_y1 := (C.double)(y1)
	c_x2 := (C.double)(x2)
	c_y2 := (C.double)(y2)
	c_x3 := (C.double)(x3)
	c_y3 := (C.double)(y3)

	C.cairo_rel_curve_to(c_ctx, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3)
}

func (ctx *Context) RelLineTo(x float64, y float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	C.cairo_rel_line_to(c_ctx, c_x, c_y)
}

func (ctx *Context) RelMoveTo(x float64, y float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	C.cairo_rel_move_to(c_ctx, c_x, c_y)
}

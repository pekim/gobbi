// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.10

package cairo

// #include <cairo/cairo.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-cairo-t.html.

func Create(surface *Surface) *Context {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_create(c_surface)
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Context) {
		o.Destroy()
	})

	return retGo
}

func (ctx *Context) Reference() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_reference(c_ctx)
}

func (ctx *Context) Destroy() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_destroy(c_ctx)
}

func (ctx *Context) Status() Status {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_status(c_ctx)
	return Status(retC)
}

func (ctx *Context) Save() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_save(c_ctx)
}

func (ctx *Context) Restore() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_restore(c_ctx)
}

func (ctx *Context) GetTarget() *Surface {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_target(c_ctx)
	retGo := SurfaceNewFromC(unsafe.Pointer(retC))

	retGo.Reference()
	runtime.SetFinalizer(retGo, func(o *Surface) {
		o.Destroy()
	})

	return retGo
}

func (ctx *Context) SetDash(dashes []float64, offset float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_dashes := (*C.double)(&dashes[0])
	c_num_dashes := C.int(len(dashes))
	c_offset := (C.double)(offset)
	C.cairo_set_dash(c_ctx, c_dashes, c_num_dashes, c_offset)
}

func (ctx *Context) SetSourceRGB(red float64, green float64, blue float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)

	C.cairo_set_source_rgb(c_ctx, c_red, c_green, c_blue)
}

func (ctx *Context) SetSourceRGBA(red float64, green float64, blue float64, alpha float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)
	c_alpha := (C.double)(alpha)

	C.cairo_set_source_rgba(c_ctx, c_red, c_green, c_blue, c_alpha)
}

func (ctx *Context) SetSource(source *Pattern) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_source := (*C.cairo_pattern_t)(source.ToC())
	C.cairo_set_source(c_ctx, c_source)
}

func (ctx *Context) SetSourceSurface(surface *Surface, x float64, y float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_surface := (*C.cairo_surface_t)(surface.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)
	C.cairo_set_source_surface(c_ctx, c_surface, c_x, c_y)
}

func (ctx *Context) GetSource() *Pattern {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_source(c_ctx)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	retGo.Reference()
	runtime.SetFinalizer(retGo, func(o *Pattern) {
		o.Destroy()
	})

	return retGo
}

func (ctx *Context) SetAntiAlias(antialias Antialias) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_antialias := (C.cairo_antialias_t)(antialias)
	C.cairo_set_antialias(c_ctx, c_antialias)
}

func (ctx *Context) GetAntiAlias() Antialias {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_antialias(c_ctx)
	return Antialias(retC)
}

func (ctx *Context) SetFillRule(fillRule FillRule) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_fillRule := (C.cairo_fill_rule_t)(fillRule)
	C.cairo_set_fill_rule(c_ctx, c_fillRule)
}

func (ctx *Context) GetFillRule() FillRule {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_fill_rule(c_ctx)
	return FillRule(retC)
}

func (ctx *Context) SetLineCap(lineCap LineCap) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_lineCap := (C.cairo_line_cap_t)(lineCap)
	C.cairo_set_line_cap(c_ctx, c_lineCap)
}

func (ctx *Context) GetLineCap() LineCap {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_line_cap(c_ctx)
	return LineCap(retC)
}

func (ctx *Context) SetLineJoin(lineJoin LineJoin) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_lineJoin := (C.cairo_line_join_t)(lineJoin)
	C.cairo_set_line_join(c_ctx, c_lineJoin)
}

func (ctx *Context) GetLineJoin() LineJoin {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_line_join(c_ctx)
	return LineJoin(retC)
}

func (ctx *Context) SetLineWidth(lineWidth float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_lineWidth := (C.double)(lineWidth)
	C.cairo_set_line_width(c_ctx, c_lineWidth)
}

func (ctx *Context) GetLineWidth() float64 {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_line_width(c_ctx)
	return float64(retC)
}

func (ctx *Context) SetMiterLimit(miterLimit float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_miterLimit := (C.double)(miterLimit)
	C.cairo_set_miter_limit(c_ctx, c_miterLimit)
}

func (ctx *Context) GetMiterLimit() float64 {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_miter_limit(c_ctx)
	return float64(retC)
}

func (ctx *Context) SetOperator(operator Operator) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_operator := (C.cairo_operator_t)(operator)
	C.cairo_set_operator(c_ctx, c_operator)
}

func (ctx *Context) GetOperator() Operator {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_operator(c_ctx)
	return Operator(retC)
}

func (ctx *Context) SetTolerance(tolerance float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_tolerance := (C.double)(tolerance)
	C.cairo_set_tolerance(c_ctx, c_tolerance)
}

func (ctx *Context) GetTolerance() float64 {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_tolerance(c_ctx)
	return float64(retC)
}

func (ctx *Context) Clip() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_clip(c_ctx)
}

func (ctx *Context) ClipPreserve() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_clip_preserve(c_ctx)
}

func (ctx *Context) ResetClip() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_reset_clip(c_ctx)
}

func (ctx *Context) Fill() {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	C.cairo_fill(c_ctx)
}

func (ctx *Context) FillPreserve() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_fill_preserve(c_ctx)
}

func (ctx *Context) FillExtents() (float64, float64, float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_x1 C.double
	var c_y1 C.double
	var c_x2 C.double
	var c_y2 C.double

	C.cairo_fill_extents(c_ctx, &c_x1, &c_y1, &c_x2, &c_y2)

	return float64(c_x1), float64(c_y1), float64(c_x2), float64(c_y2)
}

func (ctx *Context) InFill(x float64, y float64) bool {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	retC := C.cairo_in_fill(c_ctx, c_x, c_y)
	return retC == 1
}

func (ctx *Context) Mask(pattern *Pattern) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_pattern := (*C.cairo_pattern_t)(pattern.ToC())
	C.cairo_mask(c_ctx, c_pattern)
}

func (ctx *Context) SetMaskSurface(surface *Surface, surfaceX float64, surfaceY float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_surface := (*C.cairo_surface_t)(surface.ToC())
	c_surfaceX := (C.double)(surfaceX)
	c_surfaceY := (C.double)(surfaceY)
	C.cairo_set_source_surface(c_ctx, c_surface, c_surfaceX, c_surfaceY)
}

func (ctx *Context) Paint() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_paint(c_ctx)
}

func (ctx *Context) PaintWithAlpha(alpha float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_alpha := (C.double)(alpha)
	C.cairo_paint_with_alpha(c_ctx, c_alpha)
}

func (ctx *Context) Stroke() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_stroke(c_ctx)
}

func (ctx *Context) StrokePreserve() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_stroke_preserve(c_ctx)
}

func (ctx *Context) StrokeExtents() (float64, float64, float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_x1 C.double
	var c_y1 C.double
	var c_x2 C.double
	var c_y2 C.double

	C.cairo_stroke_extents(c_ctx, &c_x1, &c_y1, &c_x2, &c_y2)

	return float64(c_x1), float64(c_y1), float64(c_x2), float64(c_y2)
}

func (ctx *Context) InStroke(x float64, y float64) bool {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	retC := C.cairo_in_stroke(c_ctx, c_x, c_y)
	return retC == 1
}

func (ctx *Context) CopyPage() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_copy_page(c_ctx)
}

func (ctx *Context) ShowPage() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_show_page(c_ctx)
}

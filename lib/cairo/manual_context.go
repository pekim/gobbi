package cairo

// #include <cairo/cairo.h>
import "C"
import (
	"runtime"
	"unsafe"
)

func Create(surface *Surface) *Context {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.cairo_create(c_surface)
	retGo := ContextNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Context) {
		o.Destroy()
	})

	return retGo
}

func (ctx *Context) GetReferenceCount() int {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_reference_count(c_ctx)
	return int(retC)
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

func (ctx *Context) PushGroup() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_push_group(c_ctx)
}

func (ctx *Context) PushGroupWithContent(content Content) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_content := (C.cairo_content_t)(content)
	C.cairo_push_group_with_content(c_ctx, c_content)
}

func (ctx *Context) PopGroup() *Pattern {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_pop_group(c_ctx)
	retGo := PatternNewFromC(unsafe.Pointer(retC))

	runtime.SetFinalizer(retGo, func(o *Pattern) {
		o.Destroy()
	})

	return retGo
}

func (ctx *Context) PopGroupToSource() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_pop_group_to_source(c_ctx)
}

func (ctx *Context) GetGroupTarget() *Surface {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_group_target(c_ctx)
	retGo := SurfaceNewFromC(unsafe.Pointer(retC))

	retGo.Reference()
	runtime.SetFinalizer(retGo, func(o *Surface) {
		o.Destroy()
	})

	return retGo
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

func (ctx *Context) SetDash(dashes []float64, offset float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_dashes := (*C.double)(&dashes[0])
	c_num_dashes := C.int(len(dashes))
	c_offset := (C.double)(offset)
	C.cairo_set_dash(c_ctx, c_dashes, c_num_dashes, c_offset)
}

func (ctx *Context) GetDashCount() int {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	retC := C.cairo_get_dash_count(c_ctx)
	return int(retC)
}

func (ctx *Context) GetDash() ([]float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	dashCount := ctx.GetDashCount()
	c_dashes := make([]C.double, dashCount, dashCount)
	var c_offset C.double

	C.cairo_get_dash(c_ctx, &c_dashes[0], &c_offset)
	dashes := (*[1 << 30]float64)(unsafe.Pointer(&c_dashes[0]))[:dashCount:dashCount]

	return dashes, float64(c_offset)
}

//void 	cairo_set_fill_rule ()
//cairo_fill_rule_t 	cairo_get_fill_rule ()
//void 	cairo_set_line_cap ()
//cairo_line_cap_t 	cairo_get_line_cap ()
//void 	cairo_set_line_join ()
//cairo_line_join_t 	cairo_get_line_join ()
//void 	cairo_set_line_width ()
//double 	cairo_get_line_width ()
//void 	cairo_set_miter_limit ()
//double 	cairo_get_miter_limit ()
//void 	cairo_set_operator ()
//cairo_operator_t 	cairo_get_operator ()
//void 	cairo_set_tolerance ()
//double 	cairo_get_tolerance ()
//void 	cairo_clip ()
//void 	cairo_clip_preserve ()
//void 	cairo_clip_extents ()
//cairo_bool_t 	cairo_in_clip ()
//void 	cairo_reset_clip ()
//void 	cairo_rectangle_list_destroy ()
//cairo_rectangle_list_t * 	cairo_copy_clip_rectangle_list ()

func (ctx *Context) Fill() {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	C.cairo_fill(c_ctx)
}

//void 	cairo_fill_preserve ()
//void 	cairo_fill_extents ()
//cairo_bool_t 	cairo_in_fill ()
//void 	cairo_mask ()
//void 	cairo_mask_surface ()
//void 	cairo_paint ()
//void 	cairo_paint_with_alpha ()
//void 	cairo_stroke ()
//void 	cairo_stroke_preserve ()
//void 	cairo_stroke_extents ()
//cairo_bool_t 	cairo_in_stroke ()
//void 	cairo_copy_page ()
//void 	cairo_show_page ()
//unsigned int 	cairo_get_reference_count ()
//cairo_status_t 	cairo_set_user_data ()
//void * 	cairo_get_user_data ()

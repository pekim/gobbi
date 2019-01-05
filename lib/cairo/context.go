package cairo

// #include <cairo/cairo.h>
import "C"

func (ctx *Context) Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_xc := (C.double)(xc)
	c_yc := (C.double)(yc)
	c_radius := (C.double)(radius)
	c_angle1 := (C.double)(angle1)
	c_angle2 := (C.double)(angle2)

	C.cairo_arc(c_ctx, c_xc, c_yc, c_radius, c_angle1, c_angle2)
}

func (ctx *Context) Fill() {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	C.cairo_fill(c_ctx)
}

func (ctx *Context) SetSourceRGBA(red float64, green float64, blue float64, alpha float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_red := (C.double)(red)
	c_green := (C.double)(green)
	c_blue := (C.double)(blue)
	c_alpha := (C.double)(alpha)

	C.cairo_set_source_rgba(c_ctx, c_red, c_green, c_blue, c_alpha)
}

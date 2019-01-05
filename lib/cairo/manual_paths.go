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

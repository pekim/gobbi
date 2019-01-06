// +build cairo_1.10

package cairo

// #include <cairo/cairo.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-cairo-t.html.

func (ctx *Context) InClip(x float64, y float64) bool {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_x := (C.double)(x)
	c_y := (C.double)(y)

	retC := C.cairo_in_clip(c_ctx, c_x, c_y)
	return retC == 1
}

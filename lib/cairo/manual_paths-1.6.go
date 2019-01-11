// +build cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-Paths.html.

func (ctx *Context) HasCurrentPoint() bool {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	retC := C.cairo_has_current_point(c_ctx)

	return retC == 1
}

func (ctx *Context) PathExtents() (float64, float64, float64, float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_x1 C.double
	var c_y1 C.double
	var c_x2 C.double
	var c_y2 C.double

	C.cairo_path_extents(c_ctx, &c_x1, &c_y1, &c_x2, &c_y2)

	return float64(c_x1), float64(c_y1), float64(c_x2), float64(c_y2)
}

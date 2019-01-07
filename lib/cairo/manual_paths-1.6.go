// +build cairo_1.6 cairo_1.10

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

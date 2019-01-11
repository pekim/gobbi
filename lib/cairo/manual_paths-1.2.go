// +build cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-Paths.html.

func (ctx *Context) NewSubPath() {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	C.cairo_new_sub_path(c_ctx)
}

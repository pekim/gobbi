// +build cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-text.html.

func (ctx *Context) SetScaledFont(scaledFont *ScaledFont) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_scaledFont := (*C.cairo_scaled_font_t)(scaledFont.ToC())

	C.cairo_set_scaled_font(c_ctx, c_scaledFont)
}

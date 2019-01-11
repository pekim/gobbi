// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-cairo-scaled-font-t.html.

func (sc *ScaledFont) reference() {
	c_ff := (*C.cairo_scaled_font_t)(sc.ToC())
	C.cairo_scaled_font_reference(c_ff)
}

func (sc *ScaledFont) destroy() {
	c_ff := (*C.cairo_scaled_font_t)(sc.ToC())
	C.cairo_scaled_font_destroy(c_ff)
}

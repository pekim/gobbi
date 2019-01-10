// +build cairo_1.4 cairo_1.6 cairo cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

type FontExtents struct {
	Ascent      float64
	Descent     float64
	Height      float64
	MaxXAdvance float64
	MaxYAdvance float64
}

func fontExtentsFromC(fe *C.cairo_font_extents_t) FontExtents {
	return FontExtents{
		Ascent:      float64(fe.ascent),
		Descent:     float64(fe.descent),
		Height:      float64(fe.height),
		MaxXAdvance: float64(fe.max_x_advance),
		MaxYAdvance: float64(fe.max_y_advance),
	}
}

func (fe FontExtents) toC() *C.cairo_font_extents_t {
	return &C.cairo_font_extents_t{
		ascent:        C.double(fe.Ascent),
		descent:       C.double(fe.Descent),
		height:        C.double(fe.Height),
		max_x_advance: C.double(fe.MaxXAdvance),
		max_y_advance: C.double(fe.MaxYAdvance),
	}
}

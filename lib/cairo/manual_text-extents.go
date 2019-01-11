// +build cairo_1.4 cairo_1.6 cairo cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

type TextExtents struct {
	XBearing float64
	YBearing float64
	Height   float64
	Width    float64
	XAdvance float64
	YAdvance float64
}

func textExtentsFromC(te *C.cairo_text_extents_t) TextExtents {
	return TextExtents{
		XBearing: float64(te.x_bearing),
		YBearing: float64(te.y_bearing),
		Height:   float64(te.height),
		Width:    float64(te.width),
		XAdvance: float64(te.x_advance),
		YAdvance: float64(te.y_advance),
	}
}

func (te TextExtents) toC() *C.cairo_text_extents_t {
	return &C.cairo_text_extents_t{
		x_bearing: C.double(te.XBearing),
		y_bearing: C.double(te.YBearing),
		height:    C.double(te.Height),
		width:     C.double(te.Width),
		x_advance: C.double(te.XAdvance),
		y_advance: C.double(te.YAdvance),
	}
}

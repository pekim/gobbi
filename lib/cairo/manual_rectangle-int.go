// +build cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

type RectangleInt struct {
	X      int
	Y      int
	Width  int
	Height int
}

func rectangleIntFromC(r *C.cairo_rectangle_int_t) RectangleInt {
	return RectangleInt{
		X:      int(r.x),
		Y:      int(r.y),
		Width:  int(r.width),
		Height: int(r.height),
	}
}

func (r RectangleInt) toC() *C.cairo_rectangle_int_t {
	return &C.cairo_rectangle_int_t{
		x:      C.int(r.X),
		y:      C.int(r.X),
		width:  C.int(r.Width),
		height: C.int(r.Height),
	}
}

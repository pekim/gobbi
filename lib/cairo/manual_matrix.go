// +build cairo_1.4 cairo_1.6 cairo cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

type Matrix struct {
	XX float64
	YX float64
	XY float64
	YY float64
	X0 float64
	Y0 float64
}

func matrixFromC(m *C.cairo_matrix_t) Matrix {
	return Matrix{
		XX: float64(m.xx),
		YX: float64(m.yx),
		XY: float64(m.xy),
		YY: float64(m.yy),
		X0: float64(m.x0),
		Y0: float64(m.y0),
	}
}

func (m Matrix) toC() *C.cairo_matrix_t {
	return &C.cairo_matrix_t{
		xx: C.double(m.XX),
		yx: C.double(m.YX),
		xy: C.double(m.XY),
		yy: C.double(m.YY),
		x0: C.double(m.X0),
		y0: C.double(m.Y0),
	}
}

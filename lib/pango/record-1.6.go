// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.16 pango_1.22 pango_1.26 pango_1.31.0 pango_1.32 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Matrix is a wrapper around the C record PangoMatrix.
type Matrix struct {
	native *C.PangoMatrix
	Xx     float64
	Xy     float64
	Yx     float64
	Yy     float64
	X0     float64
	Y0     float64
}

func MatrixNewFromC(u unsafe.Pointer) *Matrix {
	c := (*C.PangoMatrix)(u)
	if c == nil {
		return nil
	}

	g := &Matrix{
		X0:     (float64)(c.x0),
		Xx:     (float64)(c.xx),
		Xy:     (float64)(c.xy),
		Y0:     (float64)(c.y0),
		Yx:     (float64)(c.yx),
		Yy:     (float64)(c.yy),
		native: c,
	}

	return g
}

func (recv *Matrix) ToC() unsafe.Pointer {
	recv.native.xx =
		(C.double)(recv.Xx)
	recv.native.xy =
		(C.double)(recv.Xy)
	recv.native.yx =
		(C.double)(recv.Yx)
	recv.native.yy =
		(C.double)(recv.Yy)
	recv.native.x0 =
		(C.double)(recv.X0)
	recv.native.y0 =
		(C.double)(recv.Y0)

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_matrix_concat : no return generator

// Copy is a wrapper around the C function pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	retC := C.pango_matrix_copy((*C.PangoMatrix)(recv.native))
	retGo := MatrixNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_matrix_free : no return generator

// Unsupported : pango_matrix_get_font_scale_factors : unsupported parameter xscale : no type generator for gdouble, double*

// Unsupported : pango_matrix_rotate : no return generator

// Unsupported : pango_matrix_scale : no return generator

// Unsupported : pango_matrix_transform_distance : unsupported parameter dx : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_pixel_rectangle : no return generator

// Unsupported : pango_matrix_transform_point : unsupported parameter x : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_rectangle : no return generator

// Unsupported : pango_matrix_translate : no return generator

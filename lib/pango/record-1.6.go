// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// Matrix is a wrapper around the C record PangoMatrix.
type Matrix struct {
	native unsafe.Pointer
	Xx     float64
	Xy     float64
	Yx     float64
	Yy     float64
	X0     float64
	Y0     float64
}

func MatrixNewFromC(u unsafe.Pointer) *Matrix {
	if u == nil {
		return nil
	}

	g := &Matrix{native: u}

	return g
}

func (recv *Matrix) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

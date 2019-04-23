// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Blacklisted : pango_font_metrics_get_strikethrough_position

// Blacklisted : pango_font_metrics_get_strikethrough_thickness

// Blacklisted : pango_font_metrics_get_underline_position

// Blacklisted : pango_font_metrics_get_underline_thickness

// Blacklisted : pango_glyph_item_free

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs :

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

// Equals compares this Matrix with another Matrix, and returns true if they represent the same GObject.
func (recv *Matrix) Equals(other *Matrix) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : pango_matrix_concat

// Blacklisted : pango_matrix_copy

// Blacklisted : pango_matrix_free

// Blacklisted : pango_matrix_rotate

// Blacklisted : pango_matrix_scale

// Blacklisted : pango_matrix_translate

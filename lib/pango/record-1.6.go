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

// GetStrikethroughPosition is a wrapper around the C function pango_font_metrics_get_strikethrough_position.
func (recv *FontMetrics) GetStrikethroughPosition() int32 {
	retC := C.pango_font_metrics_get_strikethrough_position((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStrikethroughThickness is a wrapper around the C function pango_font_metrics_get_strikethrough_thickness.
func (recv *FontMetrics) GetStrikethroughThickness() int32 {
	retC := C.pango_font_metrics_get_strikethrough_thickness((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUnderlinePosition is a wrapper around the C function pango_font_metrics_get_underline_position.
func (recv *FontMetrics) GetUnderlinePosition() int32 {
	retC := C.pango_font_metrics_get_underline_position((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUnderlineThickness is a wrapper around the C function pango_font_metrics_get_underline_thickness.
func (recv *FontMetrics) GetUnderlineThickness() int32 {
	retC := C.pango_font_metrics_get_underline_thickness((*C.PangoFontMetrics)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Free is a wrapper around the C function pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	C.pango_glyph_item_free((*C.PangoGlyphItem)(recv.native))

	return
}

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

// Concat is a wrapper around the C function pango_matrix_concat.
func (recv *Matrix) Concat(newMatrix *Matrix) {
	c_new_matrix := (*C.PangoMatrix)(C.NULL)
	if newMatrix != nil {
		c_new_matrix = (*C.PangoMatrix)(newMatrix.ToC())
	}

	C.pango_matrix_concat((*C.PangoMatrix)(recv.native), c_new_matrix)

	return
}

// Copy is a wrapper around the C function pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	retC := C.pango_matrix_copy((*C.PangoMatrix)(recv.native))
	var retGo (*Matrix)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MatrixNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Free is a wrapper around the C function pango_matrix_free.
func (recv *Matrix) Free() {
	C.pango_matrix_free((*C.PangoMatrix)(recv.native))

	return
}

// Rotate is a wrapper around the C function pango_matrix_rotate.
func (recv *Matrix) Rotate(degrees float64) {
	c_degrees := (C.double)(degrees)

	C.pango_matrix_rotate((*C.PangoMatrix)(recv.native), c_degrees)

	return
}

// Scale is a wrapper around the C function pango_matrix_scale.
func (recv *Matrix) Scale(scaleX float64, scaleY float64) {
	c_scale_x := (C.double)(scaleX)

	c_scale_y := (C.double)(scaleY)

	C.pango_matrix_scale((*C.PangoMatrix)(recv.native), c_scale_x, c_scale_y)

	return
}

// Translate is a wrapper around the C function pango_matrix_translate.
func (recv *Matrix) Translate(tx float64, ty float64) {
	c_tx := (C.double)(tx)

	c_ty := (C.double)(ty)

	C.pango_matrix_translate((*C.PangoMatrix)(recv.native), c_tx, c_ty)

	return
}

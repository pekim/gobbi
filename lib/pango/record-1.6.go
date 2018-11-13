// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Frees a #PangoGlyphItem and resources to which it points.
/*

C function : pango_glyph_item_free
*/
func (recv *GlyphItem) Free() {
	C.pango_glyph_item_free((*C.PangoGlyphItem)(recv.native))

	return
}

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs :

// A structure specifying a transformation between user-space
// coordinates and device coordinates. The transformation
// is given by
//
// <programlisting>
// x_device = x_user * matrix->xx + y_user * matrix->xy + matrix->x0;
// y_device = x_user * matrix->yx + y_user * matrix->yy + matrix->y0;
// </programlisting>
/*

C record/class : PangoMatrix
*/
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

// Changes the transformation represented by @matrix to be the
// transformation given by first applying transformation
// given by @new_matrix then applying the original transformation.
/*

C function : pango_matrix_concat
*/
func (recv *Matrix) Concat(newMatrix *Matrix) {
	c_new_matrix := (*C.PangoMatrix)(C.NULL)
	if newMatrix != nil {
		c_new_matrix = (*C.PangoMatrix)(newMatrix.ToC())
	}

	C.pango_matrix_concat((*C.PangoMatrix)(recv.native), c_new_matrix)

	return
}

// Copies a #PangoMatrix.
/*

C function : pango_matrix_copy
*/
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

// Free a #PangoMatrix created with pango_matrix_copy().
/*

C function : pango_matrix_free
*/
func (recv *Matrix) Free() {
	C.pango_matrix_free((*C.PangoMatrix)(recv.native))

	return
}

// Changes the transformation represented by @matrix to be the
// transformation given by first rotating by @degrees degrees
// counter-clockwise then applying the original transformation.
/*

C function : pango_matrix_rotate
*/
func (recv *Matrix) Rotate(degrees float64) {
	c_degrees := (C.double)(degrees)

	C.pango_matrix_rotate((*C.PangoMatrix)(recv.native), c_degrees)

	return
}

// Changes the transformation represented by @matrix to be the
// transformation given by first scaling by @sx in the X direction
// and @sy in the Y direction then applying the original
// transformation.
/*

C function : pango_matrix_scale
*/
func (recv *Matrix) Scale(scaleX float64, scaleY float64) {
	c_scale_x := (C.double)(scaleX)

	c_scale_y := (C.double)(scaleY)

	C.pango_matrix_scale((*C.PangoMatrix)(recv.native), c_scale_x, c_scale_y)

	return
}

// Changes the transformation represented by @matrix to be the
// transformation given by first translating by (@tx, @ty)
// then applying the original transformation.
/*

C function : pango_matrix_translate
*/
func (recv *Matrix) Translate(tx float64, ty float64) {
	c_tx := (C.double)(tx)

	c_ty := (C.double)(ty)

	C.pango_matrix_translate((*C.PangoMatrix)(recv.native), c_tx, c_ty)

	return
}

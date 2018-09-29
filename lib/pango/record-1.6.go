// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_attr_iterator_get_font : unsupported parameter language : record with indirection level of 2

// Unsupported : pango_attr_iterator_range : unsupported parameter start : no type generator for gint, gint*

// Unsupported : pango_attr_list_filter : unsupported parameter func : no type generator for AttrFilterFunc, PangoAttrFilterFunc

// Unsupported : pango_coverage_to_bytes : unsupported parameter bytes : no param type

// Free is a wrapper around the C function pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	C.pango_glyph_item_free((*C.PangoGlyphItem)(recv.native))

	return
}

// Unsupported : pango_glyph_item_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_item_letter_space : unsupported parameter log_attrs : no param type

// Unsupported : pango_glyph_string_get_logical_widths : unsupported parameter logical_widths : no param type

// Unsupported : pango_glyph_string_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_glyph_string_x_to_index : unsupported parameter index_ : no type generator for gint, int*

// Unsupported : pango_language_get_scripts : unsupported parameter num_scripts : no type generator for gint, int*

// Unsupported : pango_layout_line_get_x_ranges : unsupported parameter ranges : no param type

// Unsupported : pango_layout_line_index_to_x : unsupported parameter x_pos : no type generator for gint, int*

// Unsupported : pango_layout_line_x_to_index : unsupported parameter index_ : no type generator for gint, int*

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

// Concat is a wrapper around the C function pango_matrix_concat.
func (recv *Matrix) Concat(newMatrix *Matrix) {
	c_new_matrix := (*C.PangoMatrix)(newMatrix.ToC())

	C.pango_matrix_concat((*C.PangoMatrix)(recv.native), c_new_matrix)

	return
}

// Copy is a wrapper around the C function pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	retC := C.pango_matrix_copy((*C.PangoMatrix)(recv.native))
	retGo := MatrixNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function pango_matrix_free.
func (recv *Matrix) Free() {
	C.pango_matrix_free((*C.PangoMatrix)(recv.native))

	return
}

// Unsupported : pango_matrix_get_font_scale_factors : unsupported parameter xscale : no type generator for gdouble, double*

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

// Unsupported : pango_matrix_transform_distance : unsupported parameter dx : no type generator for gdouble, double*

// Unsupported : pango_matrix_transform_point : unsupported parameter x : no type generator for gdouble, double*

// Translate is a wrapper around the C function pango_matrix_translate.
func (recv *Matrix) Translate(tx float64, ty float64) {
	c_tx := (C.double)(tx)

	c_ty := (C.double)(ty)

	C.pango_matrix_translate((*C.PangoMatrix)(recv.native), c_tx, c_ty)

	return
}

// Unsupported : pango_script_iter_get_range : unsupported parameter script : PangoScript* with indirection level of 1

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

// Unsupported : pango_tab_array_get_tab : unsupported parameter alignment : PangoTabAlign* with indirection level of 1

// Unsupported : pango_tab_array_get_tabs : unsupported parameter alignments : PangoTabAlign** with indirection level of 2

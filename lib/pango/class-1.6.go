// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetFontMap is a wrapper around the C function pango_context_get_font_map.
func (recv *Context) GetFontMap() *FontMap {
	retC := C.pango_context_get_font_map((*C.PangoContext)(recv.native))
	retGo := FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMatrix is a wrapper around the C function pango_context_get_matrix.
func (recv *Context) GetMatrix() *Matrix {
	retC := C.pango_context_get_matrix((*C.PangoContext)(recv.native))
	retGo := MatrixNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_context_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_context_list_families : unsupported parameter families : no param type

// SetMatrix is a wrapper around the C function pango_context_set_matrix.
func (recv *Context) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(matrix.ToC())

	C.pango_context_set_matrix((*C.PangoContext)(recv.native), c_matrix)

	return
}

// Blacklisted : PangoEngine

// Unsupported : pango_font_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Unsupported : pango_font_face_list_sizes : unsupported parameter sizes : no param type

// Unsupported : pango_font_family_list_faces : unsupported parameter faces : no param type

// Unsupported : pango_font_map_list_families : unsupported parameter families : no param type

// Unsupported : pango_fontset_foreach : unsupported parameter func : no type generator for FontsetForeachFunc, PangoFontsetForeachFunc

// Unsupported : pango_fontset_get_metrics : return type : Blacklisted record : PangoFontMetrics

// Blacklisted : PangoFontsetSimple

// GetEllipsize is a wrapper around the C function pango_layout_get_ellipsize.
func (recv *Layout) GetEllipsize() EllipsizeMode {
	retC := C.pango_layout_get_ellipsize((*C.PangoLayout)(recv.native))
	retGo := (EllipsizeMode)(retC)

	return retGo
}

// Unsupported : pango_layout_get_iter : return type : Blacklisted record : PangoLayoutIter

// Unsupported : pango_layout_get_log_attrs : unsupported parameter attrs : no param type

// Unsupported : pango_layout_get_log_attrs_readonly : unsupported parameter n_attrs : no type generator for gint, gint*

// Unsupported : pango_layout_get_pixel_size : unsupported parameter width : no type generator for gint, int*

// Unsupported : pango_layout_get_size : unsupported parameter width : no type generator for gint, int*

// Unsupported : pango_layout_index_to_line_x : unsupported parameter line : no type generator for gint, int*

// Unsupported : pango_layout_move_cursor_visually : unsupported parameter new_index : no type generator for gint, int*

// SetEllipsize is a wrapper around the C function pango_layout_set_ellipsize.
func (recv *Layout) SetEllipsize(ellipsize EllipsizeMode) {
	c_ellipsize := (C.PangoEllipsizeMode)(ellipsize)

	C.pango_layout_set_ellipsize((*C.PangoLayout)(recv.native), c_ellipsize)

	return
}

// Unsupported : pango_layout_set_markup_with_accel : unsupported parameter accel_char : no type generator for gunichar, gunichar*

// Unsupported : pango_layout_xy_to_index : unsupported parameter index_ : no type generator for gint, int*

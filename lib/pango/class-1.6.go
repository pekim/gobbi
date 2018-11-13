// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Gets the #PangoFontMap used to look up fonts for this context.
/*

C function : pango_context_get_font_map
*/
func (recv *Context) GetFontMap() *FontMap {
	retC := C.pango_context_get_font_map((*C.PangoContext)(recv.native))
	retGo := FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the transformation matrix that will be applied when
// rendering with this context. See pango_context_set_matrix().
/*

C function : pango_context_get_matrix
*/
func (recv *Context) GetMatrix() *Matrix {
	retC := C.pango_context_get_matrix((*C.PangoContext)(recv.native))
	var retGo (*Matrix)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MatrixNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Sets the transformation matrix that will be applied when rendering
// with this context. Note that reported metrics are in the user space
// coordinates before the application of the matrix, not device-space
// coordinates after the application of the matrix. So, they don't scale
// with the matrix, though they may change slightly for different
// matrices, depending on how the text is fit to the pixel grid.
/*

C function : pango_context_set_matrix
*/
func (recv *Context) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	C.pango_context_set_matrix((*C.PangoContext)(recv.native), c_matrix)

	return
}

// Gets the type of ellipsization being performed for @layout.
// See pango_layout_set_ellipsize()
/*

C function : pango_layout_get_ellipsize
*/
func (recv *Layout) GetEllipsize() EllipsizeMode {
	retC := C.pango_layout_get_ellipsize((*C.PangoLayout)(recv.native))
	retGo := (EllipsizeMode)(retC)

	return retGo
}

// Sets the type of ellipsization being performed for @layout.
// Depending on the ellipsization mode @ellipsize text is
// removed from the start, middle, or end of text so they
// fit within the width and height of layout set with
// pango_layout_set_width() and pango_layout_set_height().
//
// If the layout contains characters such as newlines that
// force it to be layed out in multiple paragraphs, then whether
// each paragraph is ellipsized separately or the entire layout
// is ellipsized as a whole depends on the set height of the layout.
// See pango_layout_set_height() for details.
/*

C function : pango_layout_set_ellipsize
*/
func (recv *Layout) SetEllipsize(ellipsize EllipsizeMode) {
	c_ellipsize := (C.PangoEllipsizeMode)(ellipsize)

	C.pango_layout_set_ellipsize((*C.PangoLayout)(recv.native), c_ellipsize)

	return
}

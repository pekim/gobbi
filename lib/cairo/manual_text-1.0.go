// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-text.html.

func (ctx *Context) SelectFontFace(family string, slant FontSlant, weight FontWeight) {
	c_ctx := (*C.cairo_t)(ctx.ToC())

	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	c_slant := C.cairo_font_slant_t(slant)
	c_weight := C.cairo_font_weight_t(weight)

	C.cairo_select_font_face(c_ctx, c_family, c_slant, c_weight)
}

func (ctx *Context) SetFontSize(size float64) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_size := C.double(size)

	C.cairo_set_font_size(c_ctx, c_size)
}

func (ctx *Context) SetFontMatrix(matrix Matrix) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_matrix := matrix.toC()

	C.cairo_set_font_matrix(c_ctx, c_matrix)
}

func (ctx *Context) GetFontMatrix() Matrix {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	var c_matrix C.cairo_matrix_t

	C.cairo_get_font_matrix(c_ctx, &c_matrix)
	return matrixFromC(&c_matrix)
}

func (ctx *Context) SetFontOptions(options FontOptions) {
	c_ctx := (*C.cairo_t)(ctx.ToC())
	c_options := (*C.cairo_font_options_t)(options.ToC())

	C.cairo_set_font_options(c_ctx, c_options)
}

// void	cairo_select_font_face ()
// void	cairo_set_font_size ()
// void	cairo_set_font_matrix ()
// void	cairo_get_font_matrix ()
// void	cairo_set_font_options ()
// void	cairo_get_font_options ()
// void	cairo_set_font_face ()
// cairo_font_face_t *	cairo_get_font_face ()
// void	cairo_show_text ()
// void	cairo_show_glyphs ()
// void	cairo_font_extents ()
// void	cairo_text_extents ()
// void	cairo_glyph_extents ()

// 1.2
// cairo_scaled_font_t *	cairo_get_scaled_font ()

// 1.4
// void	cairo_set_scaled_font ()

// 1.8
// void	cairo_show_text_glyphs ()
// cairo_font_face_t *	cairo_toy_font_face_create ()
// const char *	cairo_toy_font_face_get_family ()
// cairo_font_slant_t	cairo_toy_font_face_get_slant ()
// cairo_font_weight_t	cairo_toy_font_face_get_weight ()
// cairo_glyph_t *	cairo_glyph_allocate ()
// void	cairo_glyph_free ()
// cairo_text_cluster_t *	cairo_text_cluster_allocate ()
// void	cairo_text_cluster_free ()

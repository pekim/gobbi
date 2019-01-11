// +build cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-text.html.

// void	cairo_show_text_glyphs ()

func ToyFontFaceCreate(family string, slant FontSlant, weight FontWeight) {
	c_family := C.CString(family)
	defer C.free(unsafe.Pointer(c_family))

	c_slant := C.cairo_font_slant_t(slant)
	c_weight := C.cairo_font_weight_t(weight)

	C.cairo_toy_font_face_create(c_family, c_slant, c_weight)
}

// cairo_font_face_t *	cairo_toy_font_face_create ()
// const char *	cairo_toy_font_face_get_family ()
// cairo_font_slant_t	cairo_toy_font_face_get_slant ()
// cairo_font_weight_t	cairo_toy_font_face_get_weight ()
// cairo_glyph_t *	cairo_glyph_allocate ()
// void	cairo_glyph_free ()
// cairo_text_cluster_t *	cairo_text_cluster_allocate ()
// void	cairo_text_cluster_free ()

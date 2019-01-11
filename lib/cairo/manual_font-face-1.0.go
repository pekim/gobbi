// +build cairo_1.0 cairo_1.2 cairo_1.4 cairo_1.6 cairo_1.8 cairo_1.10 cairo_1.12

package cairo

// #include <cairo/cairo.h>
import "C"

// The functions in this file are in the same order as they are
// documented at https://cairographics.org/manual/cairo-cairo-font-face-t.html.

func (ff *FontFace) reference() {
	c_ff := (*C.cairo_font_face_t)(ff.ToC())
	C.cairo_font_face_reference(c_ff)
}

func (ff *FontFace) destroy() {
	c_ff := (*C.cairo_font_face_t)(ff.ToC())
	C.cairo_font_face_destroy(c_ff)
}

func (ff *FontFace) Status() Status {
	c_ff := (*C.cairo_font_face_t)(ff.ToC())

	retC := C.cairo_font_face_status(c_ff)
	return Status(retC)
}

func (ff *FontFace) GetType() FontType {
	c_ff := (*C.cairo_font_face_t)(ff.ToC())

	retC := C.cairo_font_face_get_type(c_ff)
	return FontType(retC)
}

// Gobbi takes care of the referencing.
// So this function should not be needed by applications.
//
// unsigned int	cairo_font_face_get_reference_count ()

// cairo_status_t	cairo_font_face_set_user_data ()
// void *	cairo_font_face_get_user_data ()

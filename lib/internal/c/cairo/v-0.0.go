// Code generated - DO NOT EDIT.

package cairo

// #include <cairo-gobject.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
}

type Context C.cairo_t
type Device C.cairo_device_t
type Surface C.cairo_surface_t
type Matrix C.cairo_matrix_t
type Pattern C.cairo_pattern_t
type Region C.cairo_region_t
type FontOptions C.cairo_font_options_t
type FontFace C.cairo_font_face_t
type ScaledFont C.cairo_scaled_font_t
type Path C.cairo_path_t
type Rectangle C.cairo_rectangle_t
type RectangleInt C.cairo_rectangle_int_t

// UNSUPPORTED : image_surface_create : blacklisted

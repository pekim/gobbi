// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Returns a textual specification of @color in the hexadecimal form
// <literal>&num;rrrrggggbbbb</literal>, where <literal>r</literal>,
// <literal>g</literal> and <literal>b</literal> are hex digits representing
// the red, green, and blue components respectively.
/*

C function

pango_color_to_string
*/
func (recv *Color) ToString() string {
	retC := C.pango_color_to_string((*C.PangoColor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the gravity field of a font description. See
// pango_font_description_set_gravity().
/*

C function

pango_font_description_get_gravity
*/
func (recv *FontDescription) GetGravity() Gravity {
	retC := C.pango_font_description_get_gravity((*C.PangoFontDescription)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// Sets the gravity field of a font description. The gravity field
// specifies how the glyphs should be rotated.  If @gravity is
// %PANGO_GRAVITY_AUTO, this actually unsets the gravity mask on
// the font description.
//
// This function is seldom useful to the user.  Gravity should normally
// be set on a #PangoContext.
/*

C function

pango_font_description_set_gravity
*/
func (recv *FontDescription) SetGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_font_description_set_gravity((*C.PangoFontDescription)(recv.native), c_gravity)

	return
}

// Transforms the distance vector (@dx,@dy) by @matrix. This is
// similar to pango_matrix_transform_point() except that the translation
// components of the transformation are ignored. The calculation of
// the returned vector is as follows:
//
// <programlisting>
// dx2 = dx1 * xx + dy1 * xy;
// dy2 = dx1 * yx + dy1 * yy;
// </programlisting>
//
// Affine transformations are position invariant, so the same vector
// always transforms to the same vector. If (@x1,@y1) transforms
// to (@x2,@y2) then (@x1+@dx1,@y1+@dy1) will transform to
// (@x1+@dx2,@y1+@dy2) for all values of @x1 and @x2.
/*

C function

pango_matrix_transform_distance
*/
func (recv *Matrix) TransformDistance(dx float64, dy float64) {
	c_dx := (C.double)(dx)

	c_dy := (C.double)(dy)

	C.pango_matrix_transform_distance((*C.PangoMatrix)(recv.native), &c_dx, &c_dy)

	return
}

// First transforms the @rect using @matrix, then calculates the bounding box
// of the transformed rectangle.  The rectangle should be in device units
// (pixels).
//
// This function is useful for example when you want to draw a rotated
// @PangoLayout to an image buffer, and want to know how large the image
// should be and how much you should shift the layout when rendering.
//
// For better accuracy, you should use pango_matrix_transform_rectangle() on
// original rectangle in Pango units and convert to pixels afterward
// using pango_extents_to_pixels()'s first argument.
/*

C function

pango_matrix_transform_pixel_rectangle
*/
func (recv *Matrix) TransformPixelRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_pixel_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}

// Transforms the point (@x, @y) by @matrix.
/*

C function

pango_matrix_transform_point
*/
func (recv *Matrix) TransformPoint(x float64, y float64) {
	c_x := (C.double)(x)

	c_y := (C.double)(y)

	C.pango_matrix_transform_point((*C.PangoMatrix)(recv.native), &c_x, &c_y)

	return
}

// First transforms @rect using @matrix, then calculates the bounding box
// of the transformed rectangle.  The rectangle should be in Pango units.
//
// This function is useful for example when you want to draw a rotated
// @PangoLayout to an image buffer, and want to know how large the image
// should be and how much you should shift the layout when rendering.
//
// If you have a rectangle in device units (pixels), use
// pango_matrix_transform_pixel_rectangle().
//
// If you have the rectangle in Pango units and want to convert to
// transformed pixel bounding box, it is more accurate to transform it first
// (using this function) and pass the result to pango_extents_to_pixels(),
// first argument, for an inclusive rounded rectangle.
// However, there are valid reasons that you may want to convert
// to pixels first and then transform, for example when the transformed
// coordinates may overflow in Pango units (large matrix translation for
// example).
/*

C function

pango_matrix_transform_rectangle
*/
func (recv *Matrix) TransformRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}

// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// ToString is a wrapper around the C function pango_color_to_string.
func (recv *Color) ToString() string {
	retC := C.pango_color_to_string((*C.PangoColor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetGravity is a wrapper around the C function pango_font_description_get_gravity.
func (recv *FontDescription) GetGravity() Gravity {
	retC := C.pango_font_description_get_gravity((*C.PangoFontDescription)(recv.native))
	retGo := (Gravity)(retC)

	return retGo
}

// SetGravity is a wrapper around the C function pango_font_description_set_gravity.
func (recv *FontDescription) SetGravity(gravity Gravity) {
	c_gravity := (C.PangoGravity)(gravity)

	C.pango_font_description_set_gravity((*C.PangoFontDescription)(recv.native), c_gravity)

	return
}

// TransformDistance is a wrapper around the C function pango_matrix_transform_distance.
func (recv *Matrix) TransformDistance(dx float64, dy float64) {
	c_dx := (C.double)(dx)

	c_dy := (C.double)(dy)

	C.pango_matrix_transform_distance((*C.PangoMatrix)(recv.native), &c_dx, &c_dy)

	return
}

// TransformPixelRectangle is a wrapper around the C function pango_matrix_transform_pixel_rectangle.
func (recv *Matrix) TransformPixelRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_pixel_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}

// TransformPoint is a wrapper around the C function pango_matrix_transform_point.
func (recv *Matrix) TransformPoint(x float64, y float64) {
	c_x := (C.double)(x)

	c_y := (C.double)(y)

	C.pango_matrix_transform_point((*C.PangoMatrix)(recv.native), &c_x, &c_y)

	return
}

// TransformRectangle is a wrapper around the C function pango_matrix_transform_rectangle.
func (recv *Matrix) TransformRectangle(rect *Rectangle) {
	c_rect := (*C.PangoRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.PangoRectangle)(rect.ToC())
	}

	C.pango_matrix_transform_rectangle((*C.PangoMatrix)(recv.native), c_rect)

	return
}

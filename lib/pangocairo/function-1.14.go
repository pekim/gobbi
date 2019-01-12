// This is a generated file - DO NOT EDIT
// +build pangocairo_1.14 pangocairo_1.18 pangocairo_1.22

package pangocairo

import cairo "github.com/pekim/gobbi/lib/cairo"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
import "C"

// ErrorUnderlinePath is a wrapper around the C function pango_cairo_error_underline_path.
func ErrorUnderlinePath(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.double)(x)

	c_y := (C.double)(y)

	c_width := (C.double)(width)

	c_height := (C.double)(height)

	C.pango_cairo_error_underline_path(c_cr, c_x, c_y, c_width, c_height)

	return
}

// ShowErrorUnderline is a wrapper around the C function pango_cairo_show_error_underline.
func ShowErrorUnderline(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.double)(x)

	c_y := (C.double)(y)

	c_width := (C.double)(width)

	c_height := (C.double)(height)

	C.pango_cairo_show_error_underline(c_cr, c_x, c_y, c_width, c_height)

	return
}

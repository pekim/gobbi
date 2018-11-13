// This is a generated file - DO NOT EDIT
// +build pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Quantizes the thickness and position of a line, typically an
// underline or strikethrough, to whole device pixels, that is integer
// multiples of %PANGO_SCALE. The purpose of this function is to avoid
// such lines looking blurry.
//
// Care is taken to make sure @thickness is at least one pixel when this
// function returns, but returned @position may become zero as a result
// of rounding.
/*

C function : pango_quantize_line_geometry
*/
func QuantizeLineGeometry(thickness int32, position int32) {
	c_thickness := (C.int)(thickness)

	c_position := (C.int)(position)

	C.pango_quantize_line_geometry(&c_thickness, &c_position)

	return
}

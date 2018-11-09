// This is a generated file - DO NOT EDIT
// +build pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetFontScaleFactor is a wrapper around the C function pango_matrix_get_font_scale_factor.
func (recv *Matrix) GetFontScaleFactor() float64 {
	retC := C.pango_matrix_get_font_scale_factor((*C.PangoMatrix)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

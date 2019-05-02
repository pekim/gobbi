// This is a generated file - DO NOT EDIT
// +build pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "C"

// IsZeroWidth is a wrapper around the C function pango_is_zero_width.
func IsZeroWidth(ch rune) bool {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_is_zero_width(c_ch)
	retGo := retC == C.TRUE

	return retGo
}

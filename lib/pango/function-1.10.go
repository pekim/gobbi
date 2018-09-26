// This is a generated file - DO NOT EDIT
// +build pango_1.10 pango_1.12 pango_1.16 pango_1.22 pango_1.26 pango_1.31.0 pango_1.32 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// IsZeroWidth is a wrapper around the C function pango_is_zero_width.
func IsZeroWidth(ch rune) bool {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_is_zero_width(c_ch)
	retGo := retC == C.TRUE

	return retGo
}

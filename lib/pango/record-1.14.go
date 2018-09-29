// This is a generated file - DO NOT EDIT
// +build pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetWidth is a wrapper around the C function pango_glyph_string_get_width.
func (recv *GlyphString) GetWidth() int32 {
	retC := C.pango_glyph_string_get_width((*C.PangoGlyphString)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

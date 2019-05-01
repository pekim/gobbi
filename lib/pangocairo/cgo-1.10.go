// This is a generated file - DO NOT EDIT
// +build pangocairo_1.10 pangocairo_1.14 pangocairo_1.18 pangocairo_1.22

package pangocairo

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_cairo_glyph_string_path : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_cairo_show_glyph_string : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// FontMap is a wrapper around the C record PangoCairoFontMap.
type FontMap struct {
	native *C.PangoCairoFontMap
}

func FontMapNewFromC(u unsafe.Pointer) *FontMap {
	c := (*C.PangoCairoFontMap)(u)
	if c == nil {
		return nil
	}

	g := &FontMap{native: c}

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

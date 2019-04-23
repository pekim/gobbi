// This is a generated file - DO NOT EDIT
// +build pangocairo_1.18 pangocairo_1.22

package pangocairo

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
import "C"

// Font is a wrapper around the C record PangoCairoFont.
type Font struct {
	native *C.PangoCairoFont
}

func FontNewFromC(u unsafe.Pointer) *Font {
	c := (*C.PangoCairoFont)(u)
	if c == nil {
		return nil
	}

	g := &Font{native: c}

	return g
}

func (recv *Font) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Font with another Font, and returns true if they represent the same GObject.
func (recv *Font) Equals(other *Font) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : pango_cairo_font_get_scaled_font

// Blacklisted : pango_cairo_font_map_new_for_font_type

// Blacklisted : pango_cairo_font_map_get_font_type

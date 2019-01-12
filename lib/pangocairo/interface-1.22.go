// This is a generated file - DO NOT EDIT
// +build pangocairo_1.22

package pangocairo

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
import "C"

// SetDefault is a wrapper around the C function pango_cairo_font_map_set_default.
func (recv *FontMap) SetDefault() {
	C.pango_cairo_font_map_set_default((*C.PangoCairoFontMap)(recv.native))

	return
}

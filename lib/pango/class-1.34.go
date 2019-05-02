// This is a generated file - DO NOT EDIT
// +build pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Changed is a wrapper around the C function pango_font_map_changed.
func (recv *FontMap) Changed() {
	C.pango_font_map_changed((*C.PangoFontMap)(recv.native))

	return
}

// This is a generated file - DO NOT EDIT
// +build pango_1.34 pango_1.38

package pango

import "C"

// Changed is a wrapper around the C function pango_font_map_changed.
func (recv *FontMap) Changed() {
	C.pango_font_map_changed((*C.PangoFontMap)(recv.native))

	return
}

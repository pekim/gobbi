// This is a generated file - DO NOT EDIT
// +build pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Forces a change in the context, which will cause any #PangoContext
// using this fontmap to change.
//
// This function is only useful when implementing a new backend
// for Pango, something applications won't do. Backends should
// call this function if they have attached extra data to the context
// and such data is changed.
/*

C function

pango_font_map_changed
*/
func (recv *FontMap) Changed() {
	C.pango_font_map_changed((*C.PangoFontMap)(recv.native))

	return
}

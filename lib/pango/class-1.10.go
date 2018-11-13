// This is a generated file - DO NOT EDIT
// +build pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Gets the font map for which the font was created.
//
// Note that the font maintains a <firstterm>weak</firstterm> reference
// to the font map, so if all references to font map are dropped, the font
// map will be finalized even if there are fonts created with the font
// map that are still alive.  In that case this function will return %NULL.
// It is the responsibility of the user to ensure that the font map is kept
// alive.  In most uses this is not an issue as a #PangoContext holds
// a reference to the font map.
/*

C function

pango_font_get_font_map
*/
func (recv *Font) GetFontMap() *FontMap {
	retC := C.pango_font_get_font_map((*C.PangoFont)(recv.native))
	var retGo (*FontMap)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontMapNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

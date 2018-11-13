// This is a generated file - DO NOT EDIT
// +build pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Initializes @attr's klass to @klass,
// it's start_index to %PANGO_ATTR_INDEX_FROM_TEXT_BEGINNING
// and end_index to %PANGO_ATTR_INDEX_TO_TEXT_END
// such that the attribute applies
// to the entire text by default.
/*

C function : pango_attribute_init
*/
func (recv *Attribute) Init(klass *AttrClass) {
	c_klass := (*C.PangoAttrClass)(C.NULL)
	if klass != nil {
		c_klass = (*C.PangoAttrClass)(klass.ToC())
	}

	C.pango_attribute_init((*C.PangoAttribute)(recv.native), c_klass)

	return
}

// Make a deep copy of an existing #PangoGlyphItem structure.
/*

C function : pango_glyph_item_copy
*/
func (recv *GlyphItem) Copy() *GlyphItem {
	retC := C.pango_glyph_item_copy((*C.PangoGlyphItem)(recv.native))
	var retGo (*GlyphItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = GlyphItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

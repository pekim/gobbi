// This is a generated file - DO NOT EDIT
// +build pango_1.32 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// ShapeFull is a wrapper around the C function pango_shape_full.
func ShapeFull(itemText string, itemLength int32, paragraphText string, paragraphLength int32, analysis *Analysis, glyphs *GlyphString) {
	c_item_text := C.CString(itemText)
	defer C.free(unsafe.Pointer(c_item_text))

	c_item_length := (C.gint)(itemLength)

	c_paragraph_text := C.CString(paragraphText)
	defer C.free(unsafe.Pointer(c_paragraph_text))

	c_paragraph_length := (C.gint)(paragraphLength)

	c_analysis := (*C.PangoAnalysis)(analysis.ToC())

	c_glyphs := (*C.PangoGlyphString)(glyphs.ToC())

	C.pango_shape_full(c_item_text, c_item_length, c_paragraph_text, c_paragraph_length, c_analysis, c_glyphs)

	return
}

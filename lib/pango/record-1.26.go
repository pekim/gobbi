// This is a generated file - DO NOT EDIT
// +build pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetLogicalWidths is a wrapper around the C function pango_glyph_item_get_logical_widths.
func (recv *GlyphItem) GetLogicalWidths(text string, logicalWidths []int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_logical_widths := &logicalWidths[0]

	C.pango_glyph_item_get_logical_widths((*C.PangoGlyphItem)(recv.native), c_text, (*C.int)(unsafe.Pointer(c_logical_widths)))

	return
}

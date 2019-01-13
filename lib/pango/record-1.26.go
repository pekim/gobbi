// This is a generated file - DO NOT EDIT
// +build pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetLogicalWidths is a wrapper around the C function pango_glyph_item_get_logical_widths.
func (recv *GlyphItem) GetLogicalWidths(text string, logicalWidths []int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_logical_widths_array := make([]C.int, len(logicalWidths), len(logicalWidths))
	c_logical_widths := &c_logical_widths_array[0]

	C.pango_glyph_item_get_logical_widths((*C.PangoGlyphItem)(recv.native), c_text, c_logical_widths)

	return
}

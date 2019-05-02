// Code generated - DO NOT EDIT.
// +build atk_1.3 atk_1.4 atk_1.6 atk_1.9 atk_1.12 atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Unsupported : atk_text_free_ranges : unsupported parameter ranges :

// atk_text_free_ranges : unsupported parameter ranges :
// Unsupported : atk_text_get_bounded_ranges : array return type :

// GetRangeExtents is a wrapper around the C function atk_text_get_range_extents.
func (recv *Text) GetRangeExtents(startOffset int32, endOffset int32, coordType CoordType) *TextRectangle {
	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	c_coord_type := (C.AtkCoordType)(coordType)

	var c_rect C.AtkTextRectangle

	C.atk_text_get_range_extents((*C.AtkText)(recv.native), c_start_offset, c_end_offset, c_coord_type, &c_rect)

	rect := TextRectangleNewFromC(unsafe.Pointer(&c_rect))

	return rect
}

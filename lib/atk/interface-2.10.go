// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GetStringAtOffset is a wrapper around the C function atk_text_get_string_at_offset.
func (recv *Text) GetStringAtOffset(offset int32, granularity TextGranularity) (string, int32, int32) {
	c_offset := (C.gint)(offset)

	c_granularity := (C.AtkTextGranularity)(granularity)

	var c_start_offset C.gint

	var c_end_offset C.gint

	retC := C.atk_text_get_string_at_offset((*C.AtkText)(recv.native), c_offset, c_granularity, &c_start_offset, &c_end_offset)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	startOffset := (int32)(c_start_offset)

	endOffset := (int32)(c_end_offset)

	return retGo, startOffset, endOffset
}

// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Gets a portion of the text exposed through an #AtkText according to a given @offset
// and a specific @granularity, along with the start and end offsets defining the
// boundaries of such a portion of text.
//
// If @granularity is ATK_TEXT_GRANULARITY_CHAR the character at the
// offset is returned.
//
// If @granularity is ATK_TEXT_GRANULARITY_WORD the returned string
// is from the word start at or before the offset to the word start after
// the offset.
//
// The returned string will contain the word at the offset if the offset
// is inside a word and will contain the word before the offset if the
// offset is not inside a word.
//
// If @granularity is ATK_TEXT_GRANULARITY_SENTENCE the returned string
// is from the sentence start at or before the offset to the sentence
// start after the offset.
//
// The returned string will contain the sentence at the offset if the offset
// is inside a sentence and will contain the sentence before the offset
// if the offset is not inside a sentence.
//
// If @granularity is ATK_TEXT_GRANULARITY_LINE the returned string
// is from the line start at or before the offset to the line
// start after the offset.
//
// If @granularity is ATK_TEXT_GRANULARITY_PARAGRAPH the returned string
// is from the start of the paragraph at or before the offset to the start
// of the following paragraph after the offset.
/*

C function : atk_text_get_string_at_offset
*/
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

// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

type BidiType C.PangoBidiType

const (
	PANGO_BIDI_TYPE_L   BidiType = 0
	PANGO_BIDI_TYPE_LRE BidiType = 1
	PANGO_BIDI_TYPE_LRO BidiType = 2
	PANGO_BIDI_TYPE_R   BidiType = 3
	PANGO_BIDI_TYPE_AL  BidiType = 4
	PANGO_BIDI_TYPE_RLE BidiType = 5
	PANGO_BIDI_TYPE_RLO BidiType = 6
	PANGO_BIDI_TYPE_PDF BidiType = 7
	PANGO_BIDI_TYPE_EN  BidiType = 8
	PANGO_BIDI_TYPE_ES  BidiType = 9
	PANGO_BIDI_TYPE_ET  BidiType = 10
	PANGO_BIDI_TYPE_AN  BidiType = 11
	PANGO_BIDI_TYPE_CS  BidiType = 12
	PANGO_BIDI_TYPE_NSM BidiType = 13
	PANGO_BIDI_TYPE_BN  BidiType = 14
	PANGO_BIDI_TYPE_B   BidiType = 15
	PANGO_BIDI_TYPE_S   BidiType = 16
	PANGO_BIDI_TYPE_WS  BidiType = 17
	PANGO_BIDI_TYPE_ON  BidiType = 18
)

// BidiTypeForUnichar is a wrapper around the C function pango_bidi_type_for_unichar.
func BidiTypeForUnichar(ch rune) BidiType {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_bidi_type_for_unichar(c_ch)
	retGo := (BidiType)(retC)

	return retGo
}

// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// The #PangoBidiType type represents the bidirectional character
// type of a Unicode character as specified by the
// <ulink url="http://www.unicode.org/reports/tr9/">Unicode bidirectional algorithm</ulink>.
type BidiType C.PangoBidiType

const (
	// Left-to-Right
	PANGO_BIDI_TYPE_L BidiType = 0
	// Left-to-Right Embedding
	PANGO_BIDI_TYPE_LRE BidiType = 1
	// Left-to-Right Override
	PANGO_BIDI_TYPE_LRO BidiType = 2
	// Right-to-Left
	PANGO_BIDI_TYPE_R BidiType = 3
	// Right-to-Left Arabic
	PANGO_BIDI_TYPE_AL BidiType = 4
	// Right-to-Left Embedding
	PANGO_BIDI_TYPE_RLE BidiType = 5
	// Right-to-Left Override
	PANGO_BIDI_TYPE_RLO BidiType = 6
	// Pop Directional Format
	PANGO_BIDI_TYPE_PDF BidiType = 7
	// European Number
	PANGO_BIDI_TYPE_EN BidiType = 8
	// European Number Separator
	PANGO_BIDI_TYPE_ES BidiType = 9
	// European Number Terminator
	PANGO_BIDI_TYPE_ET BidiType = 10
	// Arabic Number
	PANGO_BIDI_TYPE_AN BidiType = 11
	// Common Number Separator
	PANGO_BIDI_TYPE_CS BidiType = 12
	// Nonspacing Mark
	PANGO_BIDI_TYPE_NSM BidiType = 13
	// Boundary Neutral
	PANGO_BIDI_TYPE_BN BidiType = 14
	// Paragraph Separator
	PANGO_BIDI_TYPE_B BidiType = 15
	// Segment Separator
	PANGO_BIDI_TYPE_S BidiType = 16
	// Whitespace
	PANGO_BIDI_TYPE_WS BidiType = 17
	// Other Neutrals
	PANGO_BIDI_TYPE_ON BidiType = 18
)

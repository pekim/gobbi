// This is a generated file - DO NOT EDIT

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// The bits in a #PangoFontMask correspond to fields in a
// #PangoFontDescription that have been set.
type FontMask C.PangoFontMask

const (
	// the font family is specified.
	PANGO_FONT_MASK_FAMILY FontMask = 1

	// the font style is specified.
	PANGO_FONT_MASK_STYLE FontMask = 2

	// the font variant is specified.
	PANGO_FONT_MASK_VARIANT FontMask = 4

	// the font weight is specified.
	PANGO_FONT_MASK_WEIGHT FontMask = 8

	// the font stretch is specified.
	PANGO_FONT_MASK_STRETCH FontMask = 16

	// the font size is specified.
	PANGO_FONT_MASK_SIZE FontMask = 32

	// the font gravity is specified (Since: 1.16.)
	PANGO_FONT_MASK_GRAVITY FontMask = 64
)

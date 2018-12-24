// This is a generated file - DO NOT EDIT

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

type FontMask C.PangoFontMask

const (
	PANGO_FONT_MASK_FAMILY  FontMask = 1
	PANGO_FONT_MASK_STYLE   FontMask = 2
	PANGO_FONT_MASK_VARIANT FontMask = 4
	PANGO_FONT_MASK_WEIGHT  FontMask = 8
	PANGO_FONT_MASK_STRETCH FontMask = 16
	PANGO_FONT_MASK_SIZE    FontMask = 32
	PANGO_FONT_MASK_GRAVITY FontMask = 64
)

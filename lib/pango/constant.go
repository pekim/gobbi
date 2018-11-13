// This is a generated file - DO NOT EDIT

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Blacklisted : ENGINE_TYPE_LANG

// Blacklisted : ENGINE_TYPE_SHAPE

// Unsupported : type Glyph for GLYPH_EMPTY

// Unsupported : type Glyph for GLYPH_UNKNOWN_FLAG

// Blacklisted : RENDER_TYPE_NONE

// The %PANGO_SCALE macro represents the scale between dimensions used
// for Pango distances and device units. (The definition of device
// units is dependent on the output device; it will typically be pixels
// for a screen, and points for a printer.) %PANGO_SCALE is currently
// 1024, but this may be changed in the future.
//
// When setting font sizes, device units are always considered to be
// points (as in "12 point font"), rather than pixels.
const SCALE int = C.PANGO_SCALE

// Blacklisted : UNKNOWN_GLYPH_HEIGHT

// Blacklisted : UNKNOWN_GLYPH_WIDTH

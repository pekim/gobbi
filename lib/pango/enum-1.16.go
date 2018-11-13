// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// The #PangoGravity type represents the orientation of glyphs in a segment
// of text.  This is useful when rendering vertical text layouts.  In
// those situations, the layout is rotated using a non-identity PangoMatrix,
// and then glyph orientation is controlled using #PangoGravity.
// Not every value in this enumeration makes sense for every usage of
// #PangoGravity; for example, %PANGO_GRAVITY_AUTO only can be passed to
// pango_context_set_base_gravity() and can only be returned by
// pango_context_get_base_gravity().
//
// See also: #PangoGravityHint
type Gravity C.PangoGravity

const (
	// Glyphs stand upright (default)
	PANGO_GRAVITY_SOUTH Gravity = 0
	// Glyphs are rotated 90 degrees clockwise
	PANGO_GRAVITY_EAST Gravity = 1
	// Glyphs are upside-down
	PANGO_GRAVITY_NORTH Gravity = 2
	// Glyphs are rotated 90 degrees counter-clockwise
	PANGO_GRAVITY_WEST Gravity = 3
	// Gravity is resolved from the context matrix
	PANGO_GRAVITY_AUTO Gravity = 4
)

// The #PangoGravityHint defines how horizontal scripts should behave in a
// vertical context.  That is, English excerpt in a vertical paragraph for
// example.
//
// See #PangoGravity.
type GravityHint C.PangoGravityHint

const (
	// scripts will take their natural gravity based
	// on the base gravity and the script.  This is the default.
	PANGO_GRAVITY_HINT_NATURAL GravityHint = 0
	// always use the base gravity set, regardless of
	// the script.
	PANGO_GRAVITY_HINT_STRONG GravityHint = 1
	// for scripts not in their natural direction (eg.
	// Latin in East gravity), choose per-script gravity such that every script
	// respects the line progression.  This means, Latin and Arabic will take
	// opposite gravities and both flow top-to-bottom for example.
	PANGO_GRAVITY_HINT_LINE GravityHint = 2
)

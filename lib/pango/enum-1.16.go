// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

type Gravity C.PangoGravity

const (
	PANGO_GRAVITY_SOUTH Gravity = 0
	PANGO_GRAVITY_EAST  Gravity = 1
	PANGO_GRAVITY_NORTH Gravity = 2
	PANGO_GRAVITY_WEST  Gravity = 3
	PANGO_GRAVITY_AUTO  Gravity = 4
)

type GravityHint C.PangoGravityHint

const (
	PANGO_GRAVITY_HINT_NATURAL GravityHint = 0
	PANGO_GRAVITY_HINT_STRONG  GravityHint = 1
	PANGO_GRAVITY_HINT_LINE    GravityHint = 2
)

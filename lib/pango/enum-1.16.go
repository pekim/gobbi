// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "C"

type Gravity C.PangoGravity

const (
	PANGO_GRAVITY_SOUTH Gravity = 0
	PANGO_GRAVITY_EAST  Gravity = 1
	PANGO_GRAVITY_NORTH Gravity = 2
	PANGO_GRAVITY_WEST  Gravity = 3
	PANGO_GRAVITY_AUTO  Gravity = 4
)

// GravityGetForMatrix is a wrapper around the C function pango_gravity_get_for_matrix.
func GravityGetForMatrix(matrix *Matrix) Gravity {
	c_matrix := (*C.PangoMatrix)(C.NULL)
	if matrix != nil {
		c_matrix = (*C.PangoMatrix)(matrix.ToC())
	}

	retC := C.pango_gravity_get_for_matrix(c_matrix)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityGetForScript is a wrapper around the C function pango_gravity_get_for_script.
func GravityGetForScript(script Script, baseGravity Gravity, hint GravityHint) Gravity {
	c_script := (C.PangoScript)(script)

	c_base_gravity := (C.PangoGravity)(baseGravity)

	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_gravity_get_for_script(c_script, c_base_gravity, c_hint)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityToRotation is a wrapper around the C function pango_gravity_to_rotation.
func GravityToRotation(gravity Gravity) float64 {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_gravity_to_rotation(c_gravity)
	retGo := (float64)(retC)

	return retGo
}

type GravityHint C.PangoGravityHint

const (
	PANGO_GRAVITY_HINT_NATURAL GravityHint = 0
	PANGO_GRAVITY_HINT_STRONG  GravityHint = 1
	PANGO_GRAVITY_HINT_LINE    GravityHint = 2
)

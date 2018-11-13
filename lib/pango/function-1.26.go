// This is a generated file - DO NOT EDIT
// +build pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Based on the script, East Asian width, base gravity, and hint,
// returns actual gravity to use in laying out a single character
// or #PangoItem.
//
// This function is similar to pango_gravity_get_for_script() except
// that this function makes a distinction between narrow/half-width and
// wide/full-width characters also.  Wide/full-width characters always
// stand <emphasis>upright</emphasis>, that is, they always take the base gravity,
// whereas narrow/full-width characters are always rotated in vertical
// context.
//
// If @base_gravity is %PANGO_GRAVITY_AUTO, it is first replaced with the
// preferred gravity of @script.
/*

C function

pango_gravity_get_for_script_and_width
*/
func GravityGetForScriptAndWidth(script Script, wide bool, baseGravity Gravity, hint GravityHint) Gravity {
	c_script := (C.PangoScript)(script)

	c_wide :=
		boolToGboolean(wide)

	c_base_gravity := (C.PangoGravity)(baseGravity)

	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_gravity_get_for_script_and_width(c_script, c_wide, c_base_gravity, c_hint)
	retGo := (Gravity)(retC)

	return retGo
}

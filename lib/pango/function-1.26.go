// This is a generated file - DO NOT EDIT
// +build pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GravityGetForScriptAndWidth is a wrapper around the C function pango_gravity_get_for_script_and_width.
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

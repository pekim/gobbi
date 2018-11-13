// This is a generated file - DO NOT EDIT
// +build pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Returns whether a #PangoFontFace is synthesized by the underlying
// font rendering engine from another face, perhaps by shearing, emboldening,
// or lightening it.
/*

C function

pango_font_face_is_synthesized
*/
func (recv *FontFace) IsSynthesized() bool {
	retC := C.pango_font_face_is_synthesized((*C.PangoFontFace)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

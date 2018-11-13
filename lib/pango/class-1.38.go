// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Gets the current alpha for the specified part.
/*

C function

pango_renderer_get_alpha
*/
func (recv *Renderer) GetAlpha(part RenderPart) uint16 {
	c_part := (C.PangoRenderPart)(part)

	retC := C.pango_renderer_get_alpha((*C.PangoRenderer)(recv.native), c_part)
	retGo := (uint16)(retC)

	return retGo
}

// Sets the alpha for part of the rendering.
// Note that the alpha may only be used if a color is
// specified for @part as well.
/*

C function

pango_renderer_set_alpha
*/
func (recv *Renderer) SetAlpha(part RenderPart, alpha uint16) {
	c_part := (C.PangoRenderPart)(part)

	c_alpha := (C.guint16)(alpha)

	C.pango_renderer_set_alpha((*C.PangoRenderer)(recv.native), c_part, c_alpha)

	return
}

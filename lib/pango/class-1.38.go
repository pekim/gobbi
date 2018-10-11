// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

import gobject "github.com/pekim/gobbi/lib/gobject"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

func (recv *Context) Object() *gobject.Object {}

func (recv *EngineLang) Engine() *Engine {}

func (recv *EngineShape) Engine() *Engine {}

func (recv *Font) Object() *gobject.Object {}

func (recv *FontFace) Object() *gobject.Object {}

func (recv *FontFamily) Object() *gobject.Object {}

func (recv *FontMap) Object() *gobject.Object {}

func (recv *Fontset) Object() *gobject.Object {}

func (recv *Layout) Object() *gobject.Object {}

// GetAlpha is a wrapper around the C function pango_renderer_get_alpha.
func (recv *Renderer) GetAlpha(part RenderPart) uint16 {
	c_part := (C.PangoRenderPart)(part)

	retC := C.pango_renderer_get_alpha((*C.PangoRenderer)(recv.native), c_part)
	retGo := (uint16)(retC)

	return retGo
}

// SetAlpha is a wrapper around the C function pango_renderer_set_alpha.
func (recv *Renderer) SetAlpha(part RenderPart, alpha uint16) {
	c_part := (C.PangoRenderPart)(part)

	c_alpha := (C.guint16)(alpha)

	C.pango_renderer_set_alpha((*C.PangoRenderer)(recv.native), c_part, c_alpha)

	return
}

func (recv *Renderer) Object() *gobject.Object {}

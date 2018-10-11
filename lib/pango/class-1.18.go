// This is a generated file - DO NOT EDIT
// +build pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

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

// IsSynthesized is a wrapper around the C function pango_font_face_is_synthesized.
func (recv *FontFace) IsSynthesized() bool {
	retC := C.pango_font_face_is_synthesized((*C.PangoFontFace)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

func (recv *FontFace) Object() *gobject.Object {}

func (recv *FontFamily) Object() *gobject.Object {}

func (recv *FontMap) Object() *gobject.Object {}

func (recv *Fontset) Object() *gobject.Object {}

func (recv *Layout) Object() *gobject.Object {}

func (recv *Renderer) Object() *gobject.Object {}

// This is a generated file - DO NOT EDIT
// +build pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

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

// GetCharacterCount is a wrapper around the C function pango_layout_get_character_count.
func (recv *Layout) GetCharacterCount() int32 {
	retC := C.pango_layout_get_character_count((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : pango_layout_get_log_attrs_readonly : no return type

func (recv *Layout) Object() *gobject.Object {}

func (recv *Renderer) Object() *gobject.Object {}

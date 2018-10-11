// This is a generated file - DO NOT EDIT
// +build pango_1.32.4 pango_1.34 pango_1.38

package pango

import gobject "github.com/pekim/gobbi/lib/gobject"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Changed is a wrapper around the C function pango_context_changed.
func (recv *Context) Changed() {
	C.pango_context_changed((*C.PangoContext)(recv.native))

	return
}

// GetSerial is a wrapper around the C function pango_context_get_serial.
func (recv *Context) GetSerial() uint32 {
	retC := C.pango_context_get_serial((*C.PangoContext)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

func (recv *Context) Object() *gobject.Object {}

func (recv *EngineLang) Engine() *Engine {}

func (recv *EngineShape) Engine() *Engine {}

func (recv *Font) Object() *gobject.Object {}

func (recv *FontFace) Object() *gobject.Object {}

func (recv *FontFamily) Object() *gobject.Object {}

// GetSerial is a wrapper around the C function pango_font_map_get_serial.
func (recv *FontMap) GetSerial() uint32 {
	retC := C.pango_font_map_get_serial((*C.PangoFontMap)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

func (recv *FontMap) Object() *gobject.Object {}

func (recv *Fontset) Object() *gobject.Object {}

// GetSerial is a wrapper around the C function pango_layout_get_serial.
func (recv *Layout) GetSerial() uint32 {
	retC := C.pango_layout_get_serial((*C.PangoLayout)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

func (recv *Layout) Object() *gobject.Object {}

func (recv *Renderer) Object() *gobject.Object {}

// This is a generated file - DO NOT EDIT
// +build pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

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

// GetHeight is a wrapper around the C function pango_layout_get_height.
func (recv *Layout) GetHeight() int32 {
	retC := C.pango_layout_get_height((*C.PangoLayout)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetHeight is a wrapper around the C function pango_layout_set_height.
func (recv *Layout) SetHeight(height int32) {
	c_height := (C.int)(height)

	C.pango_layout_set_height((*C.PangoLayout)(recv.native), c_height)

	return
}

func (recv *Layout) Object() *gobject.Object {}

// GetLayout is a wrapper around the C function pango_renderer_get_layout.
func (recv *Renderer) GetLayout() *Layout {
	retC := C.pango_renderer_get_layout((*C.PangoRenderer)(recv.native))
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLayoutLine is a wrapper around the C function pango_renderer_get_layout_line.
func (recv *Renderer) GetLayoutLine() *LayoutLine {
	retC := C.pango_renderer_get_layout_line((*C.PangoRenderer)(recv.native))
	retGo := LayoutLineNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *Renderer) Object() *gobject.Object {}

// This is a generated file - DO NOT EDIT
// +build pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetFontMap is a wrapper around the C function pango_context_get_font_map.
func (recv *Context) GetFontMap() *FontMap {
	retC := C.pango_context_get_font_map((*C.PangoContext)(recv.native))
	retGo := FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMatrix is a wrapper around the C function pango_context_get_matrix.
func (recv *Context) GetMatrix() *Matrix {
	retC := C.pango_context_get_matrix((*C.PangoContext)(recv.native))
	retGo := MatrixNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetMatrix is a wrapper around the C function pango_context_set_matrix.
func (recv *Context) SetMatrix(matrix *Matrix) {
	c_matrix := (*C.PangoMatrix)(matrix.ToC())

	C.pango_context_set_matrix((*C.PangoContext)(recv.native), c_matrix)

	return
}

func (recv *Context) Object() *gobject.Object {}

func (recv *EngineLang) Engine() *Engine {}

func (recv *EngineShape) Engine() *Engine {}

func (recv *Font) Object() *gobject.Object {}

func (recv *FontFace) Object() *gobject.Object {}

func (recv *FontFamily) Object() *gobject.Object {}

func (recv *FontMap) Object() *gobject.Object {}

func (recv *Fontset) Object() *gobject.Object {}

// GetEllipsize is a wrapper around the C function pango_layout_get_ellipsize.
func (recv *Layout) GetEllipsize() EllipsizeMode {
	retC := C.pango_layout_get_ellipsize((*C.PangoLayout)(recv.native))
	retGo := (EllipsizeMode)(retC)

	return retGo
}

// SetEllipsize is a wrapper around the C function pango_layout_set_ellipsize.
func (recv *Layout) SetEllipsize(ellipsize EllipsizeMode) {
	c_ellipsize := (C.PangoEllipsizeMode)(ellipsize)

	C.pango_layout_set_ellipsize((*C.PangoLayout)(recv.native), c_ellipsize)

	return
}

func (recv *Layout) Object() *gobject.Object {}

func (recv *Renderer) Object() *gobject.Object {}

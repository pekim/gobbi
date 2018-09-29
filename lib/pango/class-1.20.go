// This is a generated file - DO NOT EDIT
// +build pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

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

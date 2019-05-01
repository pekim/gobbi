// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Renderer is a wrapper around the C record PangoRenderer.
type Renderer struct {
	native *C.PangoRenderer
	// Private : parent_instance
	// Private : underline
	// Private : strikethrough
	// Private : active_count
	// matrix : record
	// Private : priv
}

func RendererNewFromC(u unsafe.Pointer) *Renderer {
	c := (*C.PangoRenderer)(u)
	if c == nil {
		return nil
	}

	g := &Renderer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Renderer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Renderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_attr_strikethrough_color_new : return type :

// Unsupported : pango_attr_underline_color_new : return type :

// RendererClass is a wrapper around the C record PangoRendererClass.
type RendererClass struct {
	native *C.PangoRendererClass
	// Private : parent_class
	// no type for draw_glyphs
	// no type for draw_rectangle
	// no type for draw_error_underline
	// no type for draw_shape
	// no type for draw_trapezoid
	// no type for draw_glyph
	// no type for part_changed
	// no type for begin
	// no type for end
	// no type for prepare_run
	// no type for draw_glyph_item
	// no type for _pango_reserved2
	// no type for _pango_reserved3
	// no type for _pango_reserved4
}

func RendererClassNewFromC(u unsafe.Pointer) *RendererClass {
	c := (*C.PangoRendererClass)(u)
	if c == nil {
		return nil
	}

	g := &RendererClass{native: c}

	return g
}

func (recv *RendererClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

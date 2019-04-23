// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Blacklisted : pango_layout_get_font_description

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

// Equals compares this Renderer with another Renderer, and returns true if they represent the same GObject.
func (recv *Renderer) Equals(other *Renderer) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Renderer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Renderer.
// Exercise care, as this is a potentially dangerous function if the Object is not a Renderer.
func CastToRenderer(object *gobject.Object) *Renderer {
	return RendererNewFromC(object.ToC())
}

// Blacklisted : pango_renderer_activate

// Blacklisted : pango_renderer_deactivate

// Blacklisted : pango_renderer_draw_error_underline

// Blacklisted : pango_renderer_draw_glyph

// Unsupported : pango_renderer_draw_glyphs : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Blacklisted : pango_renderer_draw_layout

// Blacklisted : pango_renderer_draw_layout_line

// Blacklisted : pango_renderer_draw_rectangle

// Blacklisted : pango_renderer_draw_trapezoid

// Blacklisted : pango_renderer_get_color

// Blacklisted : pango_renderer_get_matrix

// Blacklisted : pango_renderer_part_changed

// Blacklisted : pango_renderer_set_color

// Blacklisted : pango_renderer_set_matrix

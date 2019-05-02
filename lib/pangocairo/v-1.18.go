// Code generated - DO NOT EDIT.
// +build pangocairo_1.18 pangocairo_1.22

package pangocairo

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_cairo_context_get_shape_renderer : unsupported parameter data : no type generator for gpointer (gpointer*) for param data

// Unsupported : pango_cairo_context_set_shape_renderer : unsupported parameter func : no type generator for ShapeRendererFunc (PangoCairoShapeRendererFunc) for param func

// Font is a wrapper around the C record PangoCairoFont.
type Font struct {
	native *C.PangoCairoFont
}

func FontNewFromC(u unsafe.Pointer) *Font {
	c := (*C.PangoCairoFont)(u)
	if c == nil {
		return nil
	}

	g := &Font{native: c}

	return g
}

func (recv *Font) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Font with another Font, and returns true if they represent the same GObject.
func (recv *Font) Equals(other *Font) bool {
	return other.ToC() == recv.ToC()
}

// GetScaledFont is a wrapper around the C function pango_cairo_font_get_scaled_font.
func (recv *Font) GetScaledFont() *cairo.ScaledFont {
	retC := C.pango_cairo_font_get_scaled_font((*C.PangoCairoFont)(recv.native))
	var retGo (*cairo.ScaledFont)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.ScaledFontNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// FontMapNewForFontType is a wrapper around the C function pango_cairo_font_map_new_for_font_type.
func FontMapNewForFontType(fonttype cairo.FontType) *pango.FontMap {
	c_fonttype := (C.cairo_font_type_t)(fonttype)

	retC := C.pango_cairo_font_map_new_for_font_type(c_fonttype)
	var retGo (*pango.FontMap)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.FontMapNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFontType is a wrapper around the C function pango_cairo_font_map_get_font_type.
func (recv *FontMap) GetFontType() cairo.FontType {
	retC := C.pango_cairo_font_map_get_font_type((*C.PangoCairoFontMap)(recv.native))
	retGo := (cairo.FontType)(retC)

	return retGo
}

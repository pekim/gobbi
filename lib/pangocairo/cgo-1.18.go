// This is a generated file - DO NOT EDIT
// +build pangocairo_1.18 pangocairo_1.22

package pangocairo

import "unsafe"

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

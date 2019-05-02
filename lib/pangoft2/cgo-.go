// This is a generated file - DO NOT EDIT

package pangoft2

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangoft2.h>
// #include <stdlib.h>
import "C"

// FontMap is a wrapper around the C record PangoFT2FontMap.
type FontMap struct {
	native *C.PangoFT2FontMap
}

func FontMapNewFromC(u unsafe.Pointer) *FontMap {
	c := (*C.PangoFT2FontMap)(u)
	if c == nil {
		return nil
	}

	g := &FontMap{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontMap) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontMapNew is a wrapper around the C function pango_ft2_font_map_new.
func FontMapNew() *FontMap {
	retC := C.pango_ft2_font_map_new()
	retGo := FontMapNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : pango_ft2_render : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_ft2_render_transformed : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

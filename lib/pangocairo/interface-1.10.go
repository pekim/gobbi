// This is a generated file - DO NOT EDIT
// +build pangocairo_1.10 pangocairo_1.14 pangocairo_1.18 pangocairo_1.22

package pangocairo

import (
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangocairo.h>
// #include <stdlib.h>
import "C"

// FontMap is a wrapper around the C record PangoCairoFontMap.
type FontMap struct {
	native *C.PangoCairoFontMap
}

func FontMapNewFromC(u unsafe.Pointer) *FontMap {
	c := (*C.PangoCairoFontMap)(u)
	if c == nil {
		return nil
	}

	g := &FontMap{native: c}

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FontMap with another FontMap, and returns true if they represent the same GObject.
func (recv *FontMap) Equals(other *FontMap) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : pango_cairo_font_map_get_default

// FontMapNew is a wrapper around the C function pango_cairo_font_map_new.
func FontMapNew() *pango.FontMap {
	retC := C.pango_cairo_font_map_new()
	retGo := pango.FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : pango_cairo_font_map_create_context

// Blacklisted : pango_cairo_font_map_get_resolution

// Blacklisted : pango_cairo_font_map_set_resolution

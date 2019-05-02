// This is a generated file - DO NOT EDIT
// +build pangocairo_1.10 pangocairo_1.14 pangocairo_1.18 pangocairo_1.22

package pangocairo

import (
	"C"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// Equals compares this FontMap with another FontMap, and returns true if they represent the same GObject.
func (recv *FontMap) Equals(other *FontMap) bool {
	return other.ToC() == recv.ToC()
}

// FontMapGetDefault is a wrapper around the C function pango_cairo_font_map_get_default.
func FontMapGetDefault() *pango.FontMap {
	retC := C.pango_cairo_font_map_get_default()
	retGo := pango.FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontMapNew is a wrapper around the C function pango_cairo_font_map_new.
func FontMapNew() *pango.FontMap {
	retC := C.pango_cairo_font_map_new()
	retGo := pango.FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateContext is a wrapper around the C function pango_cairo_font_map_create_context.
func (recv *FontMap) CreateContext() *pango.Context {
	retC := C.pango_cairo_font_map_create_context((*C.PangoCairoFontMap)(recv.native))
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetResolution is a wrapper around the C function pango_cairo_font_map_get_resolution.
func (recv *FontMap) GetResolution() float64 {
	retC := C.pango_cairo_font_map_get_resolution((*C.PangoCairoFontMap)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetResolution is a wrapper around the C function pango_cairo_font_map_set_resolution.
func (recv *FontMap) SetResolution(dpi float64) {
	c_dpi := (C.double)(dpi)

	C.pango_cairo_font_map_set_resolution((*C.PangoCairoFontMap)(recv.native), c_dpi)

	return
}

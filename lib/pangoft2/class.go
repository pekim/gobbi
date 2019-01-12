// This is a generated file - DO NOT EDIT

package pangoft2

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
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

// Equals compares this FontMap with another FontMap, and returns true if they represent the same GObject.
func (recv *FontMap) Equals(other *FontMap) bool {
	return other.ToC() == recv.ToC()
}

// FontMap upcasts to *FontMap
func (recv *FontMap) FontMap() *pango.FontMap {
	return pango.FontMapNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FontMap) Object() *gobject.Object {
	return recv.FontMap().Object()
}

// CastToWidget down casts any arbitrary Object to FontMap.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontMap.
func CastToFontMap(object *gobject.Object) *FontMap {
	return FontMapNewFromC(object.ToC())
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

// FontMapForDisplay is a wrapper around the C function pango_ft2_font_map_for_display.
func FontMapForDisplay() *pango.FontMap {
	retC := C.pango_ft2_font_map_for_display()
	retGo := pango.FontMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateContext is a wrapper around the C function pango_ft2_font_map_create_context.
func (recv *FontMap) CreateContext() *pango.Context {
	retC := C.pango_ft2_font_map_create_context((*C.PangoFT2FontMap)(recv.native))
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_ft2_font_map_set_default_substitute : unsupported parameter func : no type generator for SubstituteFunc (PangoFT2SubstituteFunc) for param func

// SetResolution is a wrapper around the C function pango_ft2_font_map_set_resolution.
func (recv *FontMap) SetResolution(dpiX float64, dpiY float64) {
	c_dpi_x := (C.double)(dpiX)

	c_dpi_y := (C.double)(dpiY)

	C.pango_ft2_font_map_set_resolution((*C.PangoFT2FontMap)(recv.native), c_dpi_x, c_dpi_y)

	return
}

// SubstituteChanged is a wrapper around the C function pango_ft2_font_map_substitute_changed.
func (recv *FontMap) SubstituteChanged() {
	C.pango_ft2_font_map_substitute_changed((*C.PangoFT2FontMap)(recv.native))

	return
}

// Code generated - DO NOT EDIT.

package pangoft2

import (
	freetype2 "github.com/pekim/gobbi/lib/freetype2"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"reflect"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangoft2.h>
// #include <stdlib.h>
import "C"

var gobjectClassGoTypeMap = make(map[string]reflect.Type)

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

// FontGetCoverage is a wrapper around the C function pango_ft2_font_get_coverage.
func FontGetCoverage(font *pango.Font, language *pango.Language) *pango.Coverage {
	c_font := (*C.PangoFont)(C.NULL)
	if font != nil {
		c_font = (*C.PangoFont)(font.ToC())
	}

	c_language := (*C.PangoLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.PangoLanguage)(language.ToC())
	}

	retC := C.pango_ft2_font_get_coverage(c_font, c_language)
	retGo := pango.CoverageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontGetFace is a wrapper around the C function pango_ft2_font_get_face.
func FontGetFace(font *pango.Font) freetype2.Face {
	c_font := (*C.PangoFont)(C.NULL)
	if font != nil {
		c_font = (*C.PangoFont)(font.ToC())
	}

	retC := C.pango_ft2_font_get_face(c_font)
	retGo := *freetype2.FaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontGetKerning is a wrapper around the C function pango_ft2_font_get_kerning.
func FontGetKerning(font *pango.Font, left pango.Glyph, right pango.Glyph) int32 {
	c_font := (*C.PangoFont)(C.NULL)
	if font != nil {
		c_font = (*C.PangoFont)(font.ToC())
	}

	c_left := (C.PangoGlyph)(left)

	c_right := (C.PangoGlyph)(right)

	retC := C.pango_ft2_font_get_kerning(c_font, c_left, c_right)
	retGo := (int32)(retC)

	return retGo
}

// GetContext is a wrapper around the C function pango_ft2_get_context.
func GetContext(dpiX float64, dpiY float64) *pango.Context {
	c_dpi_x := (C.double)(dpiX)

	c_dpi_y := (C.double)(dpiY)

	retC := C.pango_ft2_get_context(c_dpi_x, c_dpi_y)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUnknownGlyph is a wrapper around the C function pango_ft2_get_unknown_glyph.
func GetUnknownGlyph(font *pango.Font) pango.Glyph {
	c_font := (*C.PangoFont)(C.NULL)
	if font != nil {
		c_font = (*C.PangoFont)(font.ToC())
	}

	retC := C.pango_ft2_get_unknown_glyph(c_font)
	retGo := (pango.Glyph)(retC)

	return retGo
}

// Unsupported : pango_ft2_render : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// RenderLayout is a wrapper around the C function pango_ft2_render_layout.
func RenderLayout(bitmap *freetype2.Bitmap, layout *pango.Layout, x int32, y int32) {
	c_bitmap := (*C.FT_Bitmap)(C.NULL)
	if bitmap != nil {
		c_bitmap = (*C.FT_Bitmap)(bitmap.ToC())
	}

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_ft2_render_layout(c_bitmap, c_layout, c_x, c_y)

	return
}

// RenderLayoutLine is a wrapper around the C function pango_ft2_render_layout_line.
func RenderLayoutLine(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int32, y int32) {
	c_bitmap := (*C.FT_Bitmap)(C.NULL)
	if bitmap != nil {
		c_bitmap = (*C.FT_Bitmap)(bitmap.ToC())
	}

	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_ft2_render_layout_line(c_bitmap, c_line, c_x, c_y)

	return
}

// RenderLayoutLineSubpixel is a wrapper around the C function pango_ft2_render_layout_line_subpixel.
func RenderLayoutLineSubpixel(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int32, y int32) {
	c_bitmap := (*C.FT_Bitmap)(C.NULL)
	if bitmap != nil {
		c_bitmap = (*C.FT_Bitmap)(bitmap.ToC())
	}

	c_line := (*C.PangoLayoutLine)(C.NULL)
	if line != nil {
		c_line = (*C.PangoLayoutLine)(line.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_ft2_render_layout_line_subpixel(c_bitmap, c_line, c_x, c_y)

	return
}

// RenderLayoutSubpixel is a wrapper around the C function pango_ft2_render_layout_subpixel.
func RenderLayoutSubpixel(bitmap *freetype2.Bitmap, layout *pango.Layout, x int32, y int32) {
	c_bitmap := (*C.FT_Bitmap)(C.NULL)
	if bitmap != nil {
		c_bitmap = (*C.FT_Bitmap)(bitmap.ToC())
	}

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	C.pango_ft2_render_layout_subpixel(c_bitmap, c_layout, c_x, c_y)

	return
}

// Unsupported : pango_ft2_render_transformed : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// ShutdownDisplay is a wrapper around the C function pango_ft2_shutdown_display.
func ShutdownDisplay() {
	C.pango_ft2_shutdown_display()

	return
}

// This is a generated file - DO NOT EDIT

package pangoft2

import (
	freetype2 "github.com/pekim/gobbi/lib/freetype2"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangoft2.h>
// #include <stdlib.h>
import "C"

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

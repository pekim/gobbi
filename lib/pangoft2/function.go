// Code generated - DO NOT EDIT.

package pangoft2

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	freetype2 "github.com/pekim/gobbi/lib/freetype2"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
)

var fontGetCoverageFunction *gi.Function
var fontGetCoverageFunction_Once sync.Once

func fontGetCoverageFunction_Set() error {
	var err error
	fontGetCoverageFunction_Once.Do(func() {
		fontGetCoverageFunction, err = gi.FunctionInvokerNew("PangoFT2", "font_get_coverage")
	})
	return err
}

// FontGetCoverage is a representation of the C type pango_ft2_font_get_coverage.
func FontGetCoverage(font *pango.Font, language *pango.Language) *pango.Coverage {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(font.Native())
	inArgs[1].SetPointer(language.Native())

	var ret gi.Argument

	err := fontGetCoverageFunction_Set()
	if err == nil {
		ret = fontGetCoverageFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.CoverageNewFromNative(ret.Pointer())

	return retGo
}

var fontGetFaceFunction *gi.Function
var fontGetFaceFunction_Once sync.Once

func fontGetFaceFunction_Set() error {
	var err error
	fontGetFaceFunction_Once.Do(func() {
		fontGetFaceFunction, err = gi.FunctionInvokerNew("PangoFT2", "font_get_face")
	})
	return err
}

// FontGetFace is a representation of the C type pango_ft2_font_get_face.
func FontGetFace(font *pango.Font) *freetype2.Face {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(font.Native())

	var ret gi.Argument

	err := fontGetFaceFunction_Set()
	if err == nil {
		ret = fontGetFaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := freetype2.FaceNewFromNative(ret.Pointer())

	return retGo
}

var fontGetKerningFunction *gi.Function
var fontGetKerningFunction_Once sync.Once

func fontGetKerningFunction_Set() error {
	var err error
	fontGetKerningFunction_Once.Do(func() {
		fontGetKerningFunction, err = gi.FunctionInvokerNew("PangoFT2", "font_get_kerning")
	})
	return err
}

// FontGetKerning is a representation of the C type pango_ft2_font_get_kerning.
func FontGetKerning(font *pango.Font, left pango.Glyph, right pango.Glyph) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(font.Native())
	inArgs[1].SetUint32(uint32(left))
	inArgs[2].SetUint32(uint32(right))

	var ret gi.Argument

	err := fontGetKerningFunction_Set()
	if err == nil {
		ret = fontGetKerningFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var getContextFunction *gi.Function
var getContextFunction_Once sync.Once

func getContextFunction_Set() error {
	var err error
	getContextFunction_Once.Do(func() {
		getContextFunction, err = gi.FunctionInvokerNew("PangoFT2", "get_context")
	})
	return err
}

// GetContext is a representation of the C type pango_ft2_get_context.
func GetContext(dpiX float64, dpiY float64) *pango.Context {
	var inArgs [2]gi.Argument
	inArgs[0].SetFloat64(dpiX)
	inArgs[1].SetFloat64(dpiY)

	var ret gi.Argument

	err := getContextFunction_Set()
	if err == nil {
		ret = getContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.ContextNewFromNative(ret.Pointer())

	return retGo
}

var getUnknownGlyphFunction *gi.Function
var getUnknownGlyphFunction_Once sync.Once

func getUnknownGlyphFunction_Set() error {
	var err error
	getUnknownGlyphFunction_Once.Do(func() {
		getUnknownGlyphFunction, err = gi.FunctionInvokerNew("PangoFT2", "get_unknown_glyph")
	})
	return err
}

// GetUnknownGlyph is a representation of the C type pango_ft2_get_unknown_glyph.
func GetUnknownGlyph(font *pango.Font) pango.Glyph {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(font.Native())

	var ret gi.Argument

	err := getUnknownGlyphFunction_Set()
	if err == nil {
		ret = getUnknownGlyphFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.Glyph(ret.Uint32())

	return retGo
}

var renderFunction *gi.Function
var renderFunction_Once sync.Once

func renderFunction_Set() error {
	var err error
	renderFunction_Once.Do(func() {
		renderFunction, err = gi.FunctionInvokerNew("PangoFT2", "render")
	})
	return err
}

// Render is a representation of the C type pango_ft2_render.
func Render(bitmap *freetype2.Bitmap, font *pango.Font, glyphs *pango.GlyphString, x int32, y int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(bitmap.Native())
	inArgs[1].SetPointer(font.Native())
	inArgs[2].SetPointer(glyphs.Native())
	inArgs[3].SetInt32(x)
	inArgs[4].SetInt32(y)

	err := renderFunction_Set()
	if err == nil {
		renderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var renderLayoutFunction *gi.Function
var renderLayoutFunction_Once sync.Once

func renderLayoutFunction_Set() error {
	var err error
	renderLayoutFunction_Once.Do(func() {
		renderLayoutFunction, err = gi.FunctionInvokerNew("PangoFT2", "render_layout")
	})
	return err
}

// RenderLayout is a representation of the C type pango_ft2_render_layout.
func RenderLayout(bitmap *freetype2.Bitmap, layout *pango.Layout, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(bitmap.Native())
	inArgs[1].SetPointer(layout.Native())
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := renderLayoutFunction_Set()
	if err == nil {
		renderLayoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var renderLayoutLineFunction *gi.Function
var renderLayoutLineFunction_Once sync.Once

func renderLayoutLineFunction_Set() error {
	var err error
	renderLayoutLineFunction_Once.Do(func() {
		renderLayoutLineFunction, err = gi.FunctionInvokerNew("PangoFT2", "render_layout_line")
	})
	return err
}

// RenderLayoutLine is a representation of the C type pango_ft2_render_layout_line.
func RenderLayoutLine(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(bitmap.Native())
	inArgs[1].SetPointer(line.Native())
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := renderLayoutLineFunction_Set()
	if err == nil {
		renderLayoutLineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var renderLayoutLineSubpixelFunction *gi.Function
var renderLayoutLineSubpixelFunction_Once sync.Once

func renderLayoutLineSubpixelFunction_Set() error {
	var err error
	renderLayoutLineSubpixelFunction_Once.Do(func() {
		renderLayoutLineSubpixelFunction, err = gi.FunctionInvokerNew("PangoFT2", "render_layout_line_subpixel")
	})
	return err
}

// RenderLayoutLineSubpixel is a representation of the C type pango_ft2_render_layout_line_subpixel.
func RenderLayoutLineSubpixel(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(bitmap.Native())
	inArgs[1].SetPointer(line.Native())
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := renderLayoutLineSubpixelFunction_Set()
	if err == nil {
		renderLayoutLineSubpixelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var renderLayoutSubpixelFunction *gi.Function
var renderLayoutSubpixelFunction_Once sync.Once

func renderLayoutSubpixelFunction_Set() error {
	var err error
	renderLayoutSubpixelFunction_Once.Do(func() {
		renderLayoutSubpixelFunction, err = gi.FunctionInvokerNew("PangoFT2", "render_layout_subpixel")
	})
	return err
}

// RenderLayoutSubpixel is a representation of the C type pango_ft2_render_layout_subpixel.
func RenderLayoutSubpixel(bitmap *freetype2.Bitmap, layout *pango.Layout, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(bitmap.Native())
	inArgs[1].SetPointer(layout.Native())
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := renderLayoutSubpixelFunction_Set()
	if err == nil {
		renderLayoutSubpixelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var renderTransformedFunction *gi.Function
var renderTransformedFunction_Once sync.Once

func renderTransformedFunction_Set() error {
	var err error
	renderTransformedFunction_Once.Do(func() {
		renderTransformedFunction, err = gi.FunctionInvokerNew("PangoFT2", "render_transformed")
	})
	return err
}

// RenderTransformed is a representation of the C type pango_ft2_render_transformed.
func RenderTransformed(bitmap *freetype2.Bitmap, matrix *pango.Matrix, font *pango.Font, glyphs *pango.GlyphString, x int32, y int32) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(bitmap.Native())
	inArgs[1].SetPointer(matrix.Native())
	inArgs[2].SetPointer(font.Native())
	inArgs[3].SetPointer(glyphs.Native())
	inArgs[4].SetInt32(x)
	inArgs[5].SetInt32(y)

	err := renderTransformedFunction_Set()
	if err == nil {
		renderTransformedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var shutdownDisplayFunction *gi.Function
var shutdownDisplayFunction_Once sync.Once

func shutdownDisplayFunction_Set() error {
	var err error
	shutdownDisplayFunction_Once.Do(func() {
		shutdownDisplayFunction, err = gi.FunctionInvokerNew("PangoFT2", "shutdown_display")
	})
	return err
}

// ShutdownDisplay is a representation of the C type pango_ft2_shutdown_display.
func ShutdownDisplay() {

	err := shutdownDisplayFunction_Set()
	if err == nil {
		shutdownDisplayFunction.Invoke(nil, nil)
	}

	return
}

// Code generated - DO NOT EDIT.

package pangoft2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
)

// UNSUPPORTED : C value 'pango_ft2_font_get_coverage' : parameter 'language' of type 'Pango.Language' not supported

// UNSUPPORTED : C value 'pango_ft2_font_get_face' : return type 'freetype2.Face' not supported

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

// UNSUPPORTED : C value 'pango_ft2_render' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout_line' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout_line_subpixel' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_layout_subpixel' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

// UNSUPPORTED : C value 'pango_ft2_render_transformed' : parameter 'bitmap' of type 'freetype2.Bitmap' not supported

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

// Code generated - DO NOT EDIT.

package pangocairo

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
)

// UNSUPPORTED : C value 'pango_cairo_context_get_font_options' : return type 'cairo.FontOptions' not supported

var contextGetResolutionFunction *gi.Function
var contextGetResolutionFunction_Once sync.Once

func contextGetResolutionFunction_Set() error {
	var err error
	contextGetResolutionFunction_Once.Do(func() {
		contextGetResolutionFunction, err = gi.FunctionInvokerNew("PangoCairo", "context_get_resolution")
	})
	return err
}

// ContextGetResolution is a representation of the C type pango_cairo_context_get_resolution.
func ContextGetResolution(context *pango.Context) float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := contextGetResolutionFunction_Set()
	if err == nil {
		ret = contextGetResolutionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

// UNSUPPORTED : C value 'pango_cairo_context_get_shape_renderer' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'pango_cairo_context_set_font_options' : parameter 'options' of type 'cairo.FontOptions' not supported

var contextSetResolutionFunction *gi.Function
var contextSetResolutionFunction_Once sync.Once

func contextSetResolutionFunction_Set() error {
	var err error
	contextSetResolutionFunction_Once.Do(func() {
		contextSetResolutionFunction, err = gi.FunctionInvokerNew("PangoCairo", "context_set_resolution")
	})
	return err
}

// ContextSetResolution is a representation of the C type pango_cairo_context_set_resolution.
func ContextSetResolution(context *pango.Context, dpi float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetFloat64(dpi)

	err := contextSetResolutionFunction_Set()
	if err == nil {
		contextSetResolutionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_cairo_context_set_shape_renderer' : parameter 'func' of type 'ShapeRendererFunc' not supported

// UNSUPPORTED : C value 'pango_cairo_create_context' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_create_layout' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_error_underline_path' : parameter 'cr' of type 'cairo.Context' not supported

var fontMapGetDefaultFunction *gi.Function
var fontMapGetDefaultFunction_Once sync.Once

func fontMapGetDefaultFunction_Set() error {
	var err error
	fontMapGetDefaultFunction_Once.Do(func() {
		fontMapGetDefaultFunction, err = gi.FunctionInvokerNew("PangoCairo", "font_map_get_default")
	})
	return err
}

// FontMapGetDefault is a representation of the C type pango_cairo_font_map_get_default.
func FontMapGetDefault() *pango.FontMap {

	var ret gi.Argument

	err := fontMapGetDefaultFunction_Set()
	if err == nil {
		ret = fontMapGetDefaultFunction.Invoke(nil, nil)
	}

	retGo := pango.FontMapNewFromNative(ret.Pointer())

	return retGo
}

var fontMapNewFunction *gi.Function
var fontMapNewFunction_Once sync.Once

func fontMapNewFunction_Set() error {
	var err error
	fontMapNewFunction_Once.Do(func() {
		fontMapNewFunction, err = gi.FunctionInvokerNew("PangoCairo", "font_map_new")
	})
	return err
}

// FontMapNew is a representation of the C type pango_cairo_font_map_new.
func FontMapNew() *pango.FontMap {

	var ret gi.Argument

	err := fontMapNewFunction_Set()
	if err == nil {
		ret = fontMapNewFunction.Invoke(nil, nil)
	}

	retGo := pango.FontMapNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_cairo_font_map_new_for_font_type' : parameter 'fonttype' of type 'cairo.FontType' not supported

// UNSUPPORTED : C value 'pango_cairo_glyph_string_path' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_layout_line_path' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_layout_path' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_show_error_underline' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_show_glyph_item' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_show_glyph_string' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_show_layout' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_show_layout_line' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_update_context' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'pango_cairo_update_layout' : parameter 'cr' of type 'cairo.Context' not supported

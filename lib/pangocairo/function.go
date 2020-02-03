// Code generated - DO NOT EDIT.

package pangocairo

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
)

var contextGetFontOptionsFunction *gi.Function
var contextGetFontOptionsFunction_Once sync.Once

func contextGetFontOptionsFunction_Set() error {
	var err error
	contextGetFontOptionsFunction_Once.Do(func() {
		contextGetFontOptionsFunction, err = gi.FunctionInvokerNew("PangoCairo", "context_get_font_options")
	})
	return err
}

// ContextGetFontOptions is a representation of the C type pango_cairo_context_get_font_options.
func ContextGetFontOptions(context *pango.Context) *cairo.FontOptions {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := contextGetFontOptionsFunction_Set()
	if err == nil {
		ret = contextGetFontOptionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := cairo.FontOptionsNewFromNative(ret.Pointer())

	return retGo
}

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

// UNSUPPORTED : C value 'pango_cairo_context_get_shape_renderer' : return type not supported

var contextSetFontOptionsFunction *gi.Function
var contextSetFontOptionsFunction_Once sync.Once

func contextSetFontOptionsFunction_Set() error {
	var err error
	contextSetFontOptionsFunction_Once.Do(func() {
		contextSetFontOptionsFunction, err = gi.FunctionInvokerNew("PangoCairo", "context_set_font_options")
	})
	return err
}

// ContextSetFontOptions is a representation of the C type pango_cairo_context_set_font_options.
func ContextSetFontOptions(context *pango.Context, options *cairo.FontOptions) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetPointer(options.Native())

	err := contextSetFontOptionsFunction_Set()
	if err == nil {
		contextSetFontOptionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

var createContextFunction *gi.Function
var createContextFunction_Once sync.Once

func createContextFunction_Set() error {
	var err error
	createContextFunction_Once.Do(func() {
		createContextFunction, err = gi.FunctionInvokerNew("PangoCairo", "create_context")
	})
	return err
}

// CreateContext is a representation of the C type pango_cairo_create_context.
func CreateContext(cr *cairo.Context) *pango.Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(cr.Native())

	var ret gi.Argument

	err := createContextFunction_Set()
	if err == nil {
		ret = createContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.ContextNewFromNative(ret.Pointer())

	return retGo
}

var createLayoutFunction *gi.Function
var createLayoutFunction_Once sync.Once

func createLayoutFunction_Set() error {
	var err error
	createLayoutFunction_Once.Do(func() {
		createLayoutFunction, err = gi.FunctionInvokerNew("PangoCairo", "create_layout")
	})
	return err
}

// CreateLayout is a representation of the C type pango_cairo_create_layout.
func CreateLayout(cr *cairo.Context) *pango.Layout {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(cr.Native())

	var ret gi.Argument

	err := createLayoutFunction_Set()
	if err == nil {
		ret = createLayoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.LayoutNewFromNative(ret.Pointer())

	return retGo
}

var errorUnderlinePathFunction *gi.Function
var errorUnderlinePathFunction_Once sync.Once

func errorUnderlinePathFunction_Set() error {
	var err error
	errorUnderlinePathFunction_Once.Do(func() {
		errorUnderlinePathFunction, err = gi.FunctionInvokerNew("PangoCairo", "error_underline_path")
	})
	return err
}

// ErrorUnderlinePath is a representation of the C type pango_cairo_error_underline_path.
func ErrorUnderlinePath(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetFloat64(x)
	inArgs[2].SetFloat64(y)
	inArgs[3].SetFloat64(width)
	inArgs[4].SetFloat64(height)

	err := errorUnderlinePathFunction_Set()
	if err == nil {
		errorUnderlinePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

var glyphStringPathFunction *gi.Function
var glyphStringPathFunction_Once sync.Once

func glyphStringPathFunction_Set() error {
	var err error
	glyphStringPathFunction_Once.Do(func() {
		glyphStringPathFunction, err = gi.FunctionInvokerNew("PangoCairo", "glyph_string_path")
	})
	return err
}

// GlyphStringPath is a representation of the C type pango_cairo_glyph_string_path.
func GlyphStringPath(cr *cairo.Context, font *pango.Font, glyphs *pango.GlyphString) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(font.Native())
	inArgs[2].SetPointer(glyphs.Native())

	err := glyphStringPathFunction_Set()
	if err == nil {
		glyphStringPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutLinePathFunction *gi.Function
var layoutLinePathFunction_Once sync.Once

func layoutLinePathFunction_Set() error {
	var err error
	layoutLinePathFunction_Once.Do(func() {
		layoutLinePathFunction, err = gi.FunctionInvokerNew("PangoCairo", "layout_line_path")
	})
	return err
}

// LayoutLinePath is a representation of the C type pango_cairo_layout_line_path.
func LayoutLinePath(cr *cairo.Context, line *pango.LayoutLine) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(line.Native())

	err := layoutLinePathFunction_Set()
	if err == nil {
		layoutLinePathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutPathFunction *gi.Function
var layoutPathFunction_Once sync.Once

func layoutPathFunction_Set() error {
	var err error
	layoutPathFunction_Once.Do(func() {
		layoutPathFunction, err = gi.FunctionInvokerNew("PangoCairo", "layout_path")
	})
	return err
}

// LayoutPath is a representation of the C type pango_cairo_layout_path.
func LayoutPath(cr *cairo.Context, layout *pango.Layout) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(layout.Native())

	err := layoutPathFunction_Set()
	if err == nil {
		layoutPathFunction.Invoke(inArgs[:], nil)
	}

	return
}

var showErrorUnderlineFunction *gi.Function
var showErrorUnderlineFunction_Once sync.Once

func showErrorUnderlineFunction_Set() error {
	var err error
	showErrorUnderlineFunction_Once.Do(func() {
		showErrorUnderlineFunction, err = gi.FunctionInvokerNew("PangoCairo", "show_error_underline")
	})
	return err
}

// ShowErrorUnderline is a representation of the C type pango_cairo_show_error_underline.
func ShowErrorUnderline(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetFloat64(x)
	inArgs[2].SetFloat64(y)
	inArgs[3].SetFloat64(width)
	inArgs[4].SetFloat64(height)

	err := showErrorUnderlineFunction_Set()
	if err == nil {
		showErrorUnderlineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var showGlyphItemFunction *gi.Function
var showGlyphItemFunction_Once sync.Once

func showGlyphItemFunction_Set() error {
	var err error
	showGlyphItemFunction_Once.Do(func() {
		showGlyphItemFunction, err = gi.FunctionInvokerNew("PangoCairo", "show_glyph_item")
	})
	return err
}

// ShowGlyphItem is a representation of the C type pango_cairo_show_glyph_item.
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetString(text)
	inArgs[2].SetPointer(glyphItem.Native())

	err := showGlyphItemFunction_Set()
	if err == nil {
		showGlyphItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

var showGlyphStringFunction *gi.Function
var showGlyphStringFunction_Once sync.Once

func showGlyphStringFunction_Set() error {
	var err error
	showGlyphStringFunction_Once.Do(func() {
		showGlyphStringFunction, err = gi.FunctionInvokerNew("PangoCairo", "show_glyph_string")
	})
	return err
}

// ShowGlyphString is a representation of the C type pango_cairo_show_glyph_string.
func ShowGlyphString(cr *cairo.Context, font *pango.Font, glyphs *pango.GlyphString) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(font.Native())
	inArgs[2].SetPointer(glyphs.Native())

	err := showGlyphStringFunction_Set()
	if err == nil {
		showGlyphStringFunction.Invoke(inArgs[:], nil)
	}

	return
}

var showLayoutFunction *gi.Function
var showLayoutFunction_Once sync.Once

func showLayoutFunction_Set() error {
	var err error
	showLayoutFunction_Once.Do(func() {
		showLayoutFunction, err = gi.FunctionInvokerNew("PangoCairo", "show_layout")
	})
	return err
}

// ShowLayout is a representation of the C type pango_cairo_show_layout.
func ShowLayout(cr *cairo.Context, layout *pango.Layout) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(layout.Native())

	err := showLayoutFunction_Set()
	if err == nil {
		showLayoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var showLayoutLineFunction *gi.Function
var showLayoutLineFunction_Once sync.Once

func showLayoutLineFunction_Set() error {
	var err error
	showLayoutLineFunction_Once.Do(func() {
		showLayoutLineFunction, err = gi.FunctionInvokerNew("PangoCairo", "show_layout_line")
	})
	return err
}

// ShowLayoutLine is a representation of the C type pango_cairo_show_layout_line.
func ShowLayoutLine(cr *cairo.Context, line *pango.LayoutLine) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(line.Native())

	err := showLayoutLineFunction_Set()
	if err == nil {
		showLayoutLineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var updateContextFunction *gi.Function
var updateContextFunction_Once sync.Once

func updateContextFunction_Set() error {
	var err error
	updateContextFunction_Once.Do(func() {
		updateContextFunction, err = gi.FunctionInvokerNew("PangoCairo", "update_context")
	})
	return err
}

// UpdateContext is a representation of the C type pango_cairo_update_context.
func UpdateContext(cr *cairo.Context, context *pango.Context) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(context.Native())

	err := updateContextFunction_Set()
	if err == nil {
		updateContextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var updateLayoutFunction *gi.Function
var updateLayoutFunction_Once sync.Once

func updateLayoutFunction_Set() error {
	var err error
	updateLayoutFunction_Once.Do(func() {
		updateLayoutFunction, err = gi.FunctionInvokerNew("PangoCairo", "update_layout")
	})
	return err
}

// UpdateLayout is a representation of the C type pango_cairo_update_layout.
func UpdateLayout(cr *cairo.Context, layout *pango.Layout) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(cr.Native())
	inArgs[1].SetPointer(layout.Native())

	err := updateLayoutFunction_Set()
	if err == nil {
		updateLayoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

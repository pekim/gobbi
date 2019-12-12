// Code generated - DO NOT EDIT.

package pangocairo

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

var fontInterface *gi.Interface
var fontInterface_Once sync.Once

func fontInterface_Set() error {
	var err error
	fontInterface_Once.Do(func() {
		fontInterface, err = gi.InterfaceNew("PangoCairo", "Font")
	})
	return err
}

type Font struct {
	native unsafe.Pointer
}

func FontNewFromNative(native unsafe.Pointer) *Font {
	err := fontInterface_Set()
	if err != nil {
		return nil
	}

	instance := &Font{native: native}

	return instance
}

/*
CastToFont down casts any arbitrary Object to Font.
Exercise care, as this is a potentially dangerous function
if the Object is not a Font.
*/
func CastToFont(object *gobject.Object) *Font {
	return FontNewFromNative(object.Native())
}

// Equals compares this Font with another Font, and returns true if they represent the same Object.
func (recv *Font) Equals(other *Font) bool {
	return other.Native() == recv.Native()
}

func (recv *Font) Native() unsafe.Pointer {
	return recv.native
}

var fontGetScaledFontFunction *gi.Function
var fontGetScaledFontFunction_Once sync.Once

func fontGetScaledFontFunction_Set() error {
	var err error
	fontGetScaledFontFunction_Once.Do(func() {
		err = fontInterface_Set()
		if err != nil {
			return
		}
		fontGetScaledFontFunction, err = fontInterface.InvokerNew("get_scaled_font")
	})
	return err
}

// GetScaledFont is a representation of the C type pango_cairo_font_get_scaled_font.
func (recv *Font) GetScaledFont() *cairo.ScaledFont {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontGetScaledFontFunction_Set()
	if err == nil {
		ret = fontGetScaledFontFunction.Invoke(inArgs[:], nil)
	}

	retGo := cairo.ScaledFontNewFromNative(ret.Pointer())

	return retGo
}

var fontMapInterface *gi.Interface
var fontMapInterface_Once sync.Once

func fontMapInterface_Set() error {
	var err error
	fontMapInterface_Once.Do(func() {
		fontMapInterface, err = gi.InterfaceNew("PangoCairo", "FontMap")
	})
	return err
}

type FontMap struct {
	native unsafe.Pointer
}

func FontMapNewFromNative(native unsafe.Pointer) *FontMap {
	err := fontMapInterface_Set()
	if err != nil {
		return nil
	}

	instance := &FontMap{native: native}

	return instance
}

/*
CastToFontMap down casts any arbitrary Object to FontMap.
Exercise care, as this is a potentially dangerous function
if the Object is not a FontMap.
*/
func CastToFontMap(object *gobject.Object) *FontMap {
	return FontMapNewFromNative(object.Native())
}

// Equals compares this FontMap with another FontMap, and returns true if they represent the same Object.
func (recv *FontMap) Equals(other *FontMap) bool {
	return other.Native() == recv.Native()
}

func (recv *FontMap) Native() unsafe.Pointer {
	return recv.native
}

var fontMapCreateContextFunction *gi.Function
var fontMapCreateContextFunction_Once sync.Once

func fontMapCreateContextFunction_Set() error {
	var err error
	fontMapCreateContextFunction_Once.Do(func() {
		err = fontMapInterface_Set()
		if err != nil {
			return
		}
		fontMapCreateContextFunction, err = fontMapInterface.InvokerNew("create_context")
	})
	return err
}

// CreateContext is a representation of the C type pango_cairo_font_map_create_context.
func (recv *FontMap) CreateContext() *pango.Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontMapCreateContextFunction_Set()
	if err == nil {
		ret = fontMapCreateContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.ContextNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_cairo_font_map_get_font_type' : return type 'cairo.FontType' not supported

var fontMapGetResolutionFunction *gi.Function
var fontMapGetResolutionFunction_Once sync.Once

func fontMapGetResolutionFunction_Set() error {
	var err error
	fontMapGetResolutionFunction_Once.Do(func() {
		err = fontMapInterface_Set()
		if err != nil {
			return
		}
		fontMapGetResolutionFunction, err = fontMapInterface.InvokerNew("get_resolution")
	})
	return err
}

// GetResolution is a representation of the C type pango_cairo_font_map_get_resolution.
func (recv *FontMap) GetResolution() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontMapGetResolutionFunction_Set()
	if err == nil {
		ret = fontMapGetResolutionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var fontMapSetDefaultFunction *gi.Function
var fontMapSetDefaultFunction_Once sync.Once

func fontMapSetDefaultFunction_Set() error {
	var err error
	fontMapSetDefaultFunction_Once.Do(func() {
		err = fontMapInterface_Set()
		if err != nil {
			return
		}
		fontMapSetDefaultFunction, err = fontMapInterface.InvokerNew("set_default")
	})
	return err
}

// SetDefault is a representation of the C type pango_cairo_font_map_set_default.
func (recv *FontMap) SetDefault() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := fontMapSetDefaultFunction_Set()
	if err == nil {
		fontMapSetDefaultFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontMapSetResolutionFunction *gi.Function
var fontMapSetResolutionFunction_Once sync.Once

func fontMapSetResolutionFunction_Set() error {
	var err error
	fontMapSetResolutionFunction_Once.Do(func() {
		err = fontMapInterface_Set()
		if err != nil {
			return
		}
		fontMapSetResolutionFunction, err = fontMapInterface.InvokerNew("set_resolution")
	})
	return err
}

// SetResolution is a representation of the C type pango_cairo_font_map_set_resolution.
func (recv *FontMap) SetResolution(dpi float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetFloat64(dpi)

	err := fontMapSetResolutionFunction_Set()
	if err == nil {
		fontMapSetResolutionFunction.Invoke(inArgs[:], nil)
	}

	return
}

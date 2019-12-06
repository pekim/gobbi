// Code generated - DO NOT EDIT.

package pangoft2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

var fontMapObject *gi.Object
var fontMapObject_Once sync.Once

func fontMapObject_Set() error {
	var err error
	fontMapObject_Once.Do(func() {
		fontMapObject, err = gi.ObjectNew("PangoFT2", "FontMap")
	})
	return err
}

type FontMap struct {
	native unsafe.Pointer
}

func FontMapNewFromNative(native unsafe.Pointer) *FontMap {
	return &FontMap{native: native}
}

// FontMap upcasts to *FontMap
func (recv *FontMap) FontMap() *pango.FontMap {
	return pango.FontMapNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *FontMap) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

var fontMapNewFunction *gi.Function
var fontMapNewFunction_Once sync.Once

func fontMapNewFunction_Set() error {
	var err error
	fontMapNewFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapNewFunction, err = fontMapObject.InvokerNew("new")
	})
	return err
}

// FontMapNew is a representation of the C type pango_ft2_font_map_new.
func FontMapNew() *FontMap {

	var ret gi.Argument

	err := fontMapNewFunction_Set()
	if err == nil {
		ret = fontMapNewFunction.Invoke(nil, nil)
	}

	retGo := FontMapNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_ft2_font_map_create_context' : return type 'Pango.Context' not supported

// UNSUPPORTED : C value 'pango_ft2_font_map_set_default_substitute' : parameter 'func' of type 'SubstituteFunc' not supported

var fontMapSetResolutionFunction *gi.Function
var fontMapSetResolutionFunction_Once sync.Once

func fontMapSetResolutionFunction_Set() error {
	var err error
	fontMapSetResolutionFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapSetResolutionFunction, err = fontMapObject.InvokerNew("set_resolution")
	})
	return err
}

// SetResolution is a representation of the C type pango_ft2_font_map_set_resolution.
func (recv *FontMap) SetResolution(dpiX float64, dpiY float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(dpiX)
	inArgs[2].SetFloat64(dpiY)

	err := fontMapSetResolutionFunction_Set()
	if err == nil {
		fontMapSetResolutionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontMapSubstituteChangedFunction *gi.Function
var fontMapSubstituteChangedFunction_Once sync.Once

func fontMapSubstituteChangedFunction_Set() error {
	var err error
	fontMapSubstituteChangedFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapSubstituteChangedFunction, err = fontMapObject.InvokerNew("substitute_changed")
	})
	return err
}

// SubstituteChanged is a representation of the C type pango_ft2_font_map_substitute_changed.
func (recv *FontMap) SubstituteChanged() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontMapSubstituteChangedFunction_Set()
	if err == nil {
		fontMapSubstituteChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

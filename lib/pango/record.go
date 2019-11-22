// Code generated - DO NOT EDIT.

package pango

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var analysisStruct *gi.Struct
var analysisStruct_Once sync.Once

func analysisStruct_Set() error {
	var err error
	analysisStruct_Once.Do(func() {
		analysisStruct, err = gi.StructNew("Pango", "Analysis")
	})
	return err
}

type Analysis struct {
	native uintptr
	// UNSUPPORTED : C value 'shape_engine' : no Go type for 'EngineShape'
	// UNSUPPORTED : C value 'lang_engine' : no Go type for 'EngineLang'
	// UNSUPPORTED : C value 'font' : no Go type for 'Font'
	Level    uint8
	Gravity  uint8
	Flags    uint8
	Script   uint8
	Language *Language
	// UNSUPPORTED : C value 'extra_attrs' : no Go type for 'GLib.SList'
}

var attrClassStruct *gi.Struct
var attrClassStruct_Once sync.Once

func attrClassStruct_Set() error {
	var err error
	attrClassStruct_Once.Do(func() {
		attrClassStruct, err = gi.StructNew("Pango", "AttrClass")
	})
	return err
}

type AttrClass struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'AttrType'
	// UNSUPPORTED : C value 'copy' : missing Type
	// UNSUPPORTED : C value 'destroy' : missing Type
	// UNSUPPORTED : C value 'equal' : missing Type
}

var attrColorStruct *gi.Struct
var attrColorStruct_Once sync.Once

func attrColorStruct_Set() error {
	var err error
	attrColorStruct_Once.Do(func() {
		attrColorStruct, err = gi.StructNew("Pango", "AttrColor")
	})
	return err
}

type AttrColor struct {
	native uintptr
	Attr   *Attribute
	Color  *Color
}

var attrFloatStruct *gi.Struct
var attrFloatStruct_Once sync.Once

func attrFloatStruct_Set() error {
	var err error
	attrFloatStruct_Once.Do(func() {
		attrFloatStruct, err = gi.StructNew("Pango", "AttrFloat")
	})
	return err
}

type AttrFloat struct {
	native uintptr
	Attr   *Attribute
	// UNSUPPORTED : C value 'value' : no Go type for 'gdouble'
}

var attrFontDescStruct *gi.Struct
var attrFontDescStruct_Once sync.Once

func attrFontDescStruct_Set() error {
	var err error
	attrFontDescStruct_Once.Do(func() {
		attrFontDescStruct, err = gi.StructNew("Pango", "AttrFontDesc")
	})
	return err
}

type AttrFontDesc struct {
	native uintptr
	Attr   *Attribute
	Desc   *FontDescription
}

var attrFontFeaturesStruct *gi.Struct
var attrFontFeaturesStruct_Once sync.Once

func attrFontFeaturesStruct_Set() error {
	var err error
	attrFontFeaturesStruct_Once.Do(func() {
		attrFontFeaturesStruct, err = gi.StructNew("Pango", "AttrFontFeatures")
	})
	return err
}

type AttrFontFeatures struct {
	native   uintptr
	Attr     *Attribute
	Features string
}

var attrIntStruct *gi.Struct
var attrIntStruct_Once sync.Once

func attrIntStruct_Set() error {
	var err error
	attrIntStruct_Once.Do(func() {
		attrIntStruct, err = gi.StructNew("Pango", "AttrInt")
	})
	return err
}

type AttrInt struct {
	native uintptr
	Attr   *Attribute
	Value  int32
}

var attrIteratorStruct *gi.Struct
var attrIteratorStruct_Once sync.Once

func attrIteratorStruct_Set() error {
	var err error
	attrIteratorStruct_Once.Do(func() {
		attrIteratorStruct, err = gi.StructNew("Pango", "AttrIterator")
	})
	return err
}

type AttrIterator struct {
	native uintptr
}

var attrIteratorCopyFunction *gi.Function
var attrIteratorCopyFunction_Once sync.Once

func attrIteratorCopyFunction_Set() error {
	var err error
	attrIteratorCopyFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorCopyFunction, err = attrIteratorStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_attr_iterator_copy.
func (recv *AttrIterator) Copy() *AttrIterator {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrIteratorCopyFunction_Set()
	if err == nil {
		ret = attrIteratorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

var attrIteratorDestroyFunction *gi.Function
var attrIteratorDestroyFunction_Once sync.Once

func attrIteratorDestroyFunction_Set() error {
	var err error
	attrIteratorDestroyFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorDestroyFunction, err = attrIteratorStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type pango_attr_iterator_destroy.
func (recv *AttrIterator) Destroy() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attrIteratorDestroyFunction_Set()
	if err == nil {
		attrIteratorDestroyFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_attr_iterator_get' : parameter 'type' of type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_attrs' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_font' : parameter 'desc' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_next' : return type 'gboolean' not supported

var attrIteratorRangeFunction *gi.Function
var attrIteratorRangeFunction_Once sync.Once

func attrIteratorRangeFunction_Set() error {
	var err error
	attrIteratorRangeFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorRangeFunction, err = attrIteratorStruct.InvokerNew("range")
	})
	return err
}

// Range is a representation of the C type pango_attr_iterator_range.
func (recv *AttrIterator) Range() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := attrIteratorRangeFunction_Set()
	if err == nil {
		attrIteratorRangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var attrLanguageStruct *gi.Struct
var attrLanguageStruct_Once sync.Once

func attrLanguageStruct_Set() error {
	var err error
	attrLanguageStruct_Once.Do(func() {
		attrLanguageStruct, err = gi.StructNew("Pango", "AttrLanguage")
	})
	return err
}

type AttrLanguage struct {
	native uintptr
	Attr   *Attribute
	Value  *Language
}

var attrListStruct *gi.Struct
var attrListStruct_Once sync.Once

func attrListStruct_Set() error {
	var err error
	attrListStruct_Once.Do(func() {
		attrListStruct, err = gi.StructNew("Pango", "AttrList")
	})
	return err
}

type AttrList struct {
	native uintptr
}

var attrListNewFunction *gi.Function
var attrListNewFunction_Once sync.Once

func attrListNewFunction_Set() error {
	var err error
	attrListNewFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListNewFunction, err = attrListStruct.InvokerNew("new")
	})
	return err
}

// AttrListNew is a representation of the C type pango_attr_list_new.
func AttrListNew() *AttrList {

	var ret gi.Argument

	err := attrListNewFunction_Set()
	if err == nil {
		ret = attrListNewFunction.Invoke(nil, nil)
	}

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_change' : parameter 'attr' of type 'Attribute' not supported

var attrListCopyFunction *gi.Function
var attrListCopyFunction_Once sync.Once

func attrListCopyFunction_Set() error {
	var err error
	attrListCopyFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListCopyFunction, err = attrListStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_attr_list_copy.
func (recv *AttrList) Copy() *AttrList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListCopyFunction_Set()
	if err == nil {
		ret = attrListCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_filter' : parameter 'func' of type 'AttrFilterFunc' not supported

var attrListGetIteratorFunction *gi.Function
var attrListGetIteratorFunction_Once sync.Once

func attrListGetIteratorFunction_Set() error {
	var err error
	attrListGetIteratorFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListGetIteratorFunction, err = attrListStruct.InvokerNew("get_iterator")
	})
	return err
}

// GetIterator is a representation of the C type pango_attr_list_get_iterator.
func (recv *AttrList) GetIterator() *AttrIterator {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListGetIteratorFunction_Set()
	if err == nil {
		ret = attrListGetIteratorFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_insert' : parameter 'attr' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_list_insert_before' : parameter 'attr' of type 'Attribute' not supported

var attrListRefFunction *gi.Function
var attrListRefFunction_Once sync.Once

func attrListRefFunction_Set() error {
	var err error
	attrListRefFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListRefFunction, err = attrListStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_attr_list_ref.
func (recv *AttrList) Ref() *AttrList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListRefFunction_Set()
	if err == nil {
		ret = attrListRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_splice' : parameter 'other' of type 'AttrList' not supported

var attrListUnrefFunction *gi.Function
var attrListUnrefFunction_Once sync.Once

func attrListUnrefFunction_Set() error {
	var err error
	attrListUnrefFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListUnrefFunction, err = attrListStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_attr_list_unref.
func (recv *AttrList) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attrListUnrefFunction_Set()
	if err == nil {
		attrListUnrefFunction.Invoke(inArgs[:], nil)
	}

}

var attrShapeStruct *gi.Struct
var attrShapeStruct_Once sync.Once

func attrShapeStruct_Set() error {
	var err error
	attrShapeStruct_Once.Do(func() {
		attrShapeStruct, err = gi.StructNew("Pango", "AttrShape")
	})
	return err
}

type AttrShape struct {
	native      uintptr
	Attr        *Attribute
	InkRect     *Rectangle
	LogicalRect *Rectangle
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'copy_func' : no Go type for 'AttrDataCopyFunc'
	// UNSUPPORTED : C value 'destroy_func' : no Go type for 'GLib.DestroyNotify'
}

var attrSizeStruct *gi.Struct
var attrSizeStruct_Once sync.Once

func attrSizeStruct_Set() error {
	var err error
	attrSizeStruct_Once.Do(func() {
		attrSizeStruct, err = gi.StructNew("Pango", "AttrSize")
	})
	return err
}

type AttrSize struct {
	native   uintptr
	Attr     *Attribute
	Size     int32
	Absolute uint32
}

var attrStringStruct *gi.Struct
var attrStringStruct_Once sync.Once

func attrStringStruct_Set() error {
	var err error
	attrStringStruct_Once.Do(func() {
		attrStringStruct, err = gi.StructNew("Pango", "AttrString")
	})
	return err
}

type AttrString struct {
	native uintptr
	Attr   *Attribute
	Value  string
}

var attributeStruct *gi.Struct
var attributeStruct_Once sync.Once

func attributeStruct_Set() error {
	var err error
	attributeStruct_Once.Do(func() {
		attributeStruct, err = gi.StructNew("Pango", "Attribute")
	})
	return err
}

type Attribute struct {
	native     uintptr
	Klass      *AttrClass
	StartIndex uint32
	EndIndex   uint32
}

var attributeCopyFunction *gi.Function
var attributeCopyFunction_Once sync.Once

func attributeCopyFunction_Set() error {
	var err error
	attributeCopyFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeCopyFunction, err = attributeStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_attribute_copy.
func (recv *Attribute) Copy() *Attribute {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attributeCopyFunction_Set()
	if err == nil {
		ret = attributeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var attributeDestroyFunction *gi.Function
var attributeDestroyFunction_Once sync.Once

func attributeDestroyFunction_Set() error {
	var err error
	attributeDestroyFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeDestroyFunction, err = attributeStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type pango_attribute_destroy.
func (recv *Attribute) Destroy() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attributeDestroyFunction_Set()
	if err == nil {
		attributeDestroyFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_attribute_equal' : parameter 'attr2' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attribute_init' : parameter 'klass' of type 'AttrClass' not supported

var colorStruct *gi.Struct
var colorStruct_Once sync.Once

func colorStruct_Set() error {
	var err error
	colorStruct_Once.Do(func() {
		colorStruct, err = gi.StructNew("Pango", "Color")
	})
	return err
}

type Color struct {
	native uintptr
	Red    uint16
	Green  uint16
	Blue   uint16
}

var colorCopyFunction *gi.Function
var colorCopyFunction_Once sync.Once

func colorCopyFunction_Set() error {
	var err error
	colorCopyFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorCopyFunction, err = colorStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_color_copy.
func (recv *Color) Copy() *Color {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorCopyFunction_Set()
	if err == nil {
		ret = colorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Color{native: ret.Pointer()}

	return retGo
}

var colorFreeFunction *gi.Function
var colorFreeFunction_Once sync.Once

func colorFreeFunction_Set() error {
	var err error
	colorFreeFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorFreeFunction, err = colorStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_color_free.
func (recv *Color) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := colorFreeFunction_Set()
	if err == nil {
		colorFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_color_parse' : return type 'gboolean' not supported

var colorToStringFunction *gi.Function
var colorToStringFunction_Once sync.Once

func colorToStringFunction_Set() error {
	var err error
	colorToStringFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorToStringFunction, err = colorStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type pango_color_to_string.
func (recv *Color) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorToStringFunction_Set()
	if err == nil {
		ret = colorToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var contextClassStruct *gi.Struct
var contextClassStruct_Once sync.Once

func contextClassStruct_Set() error {
	var err error
	contextClassStruct_Once.Do(func() {
		contextClassStruct, err = gi.StructNew("Pango", "ContextClass")
	})
	return err
}

type ContextClass struct {
	native uintptr
}

var coverageStruct *gi.Struct
var coverageStruct_Once sync.Once

func coverageStruct_Set() error {
	var err error
	coverageStruct_Once.Do(func() {
		coverageStruct, err = gi.StructNew("Pango", "Coverage")
	})
	return err
}

type Coverage struct {
	native uintptr
}

var coverageCopyFunction *gi.Function
var coverageCopyFunction_Once sync.Once

func coverageCopyFunction_Set() error {
	var err error
	coverageCopyFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageCopyFunction, err = coverageStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_coverage_copy.
func (recv *Coverage) Copy() *Coverage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := coverageCopyFunction_Set()
	if err == nil {
		ret = coverageCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_get' : return type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_max' : parameter 'other' of type 'Coverage' not supported

var coverageRefFunction *gi.Function
var coverageRefFunction_Once sync.Once

func coverageRefFunction_Set() error {
	var err error
	coverageRefFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageRefFunction, err = coverageStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_coverage_ref.
func (recv *Coverage) Ref() *Coverage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := coverageRefFunction_Set()
	if err == nil {
		ret = coverageRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_set' : parameter 'level' of type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_to_bytes' : parameter 'bytes' has no type

var coverageUnrefFunction *gi.Function
var coverageUnrefFunction_Once sync.Once

func coverageUnrefFunction_Set() error {
	var err error
	coverageUnrefFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageUnrefFunction, err = coverageStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_coverage_unref.
func (recv *Coverage) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := coverageUnrefFunction_Set()
	if err == nil {
		coverageUnrefFunction.Invoke(inArgs[:], nil)
	}

}

var engineClassStruct *gi.Struct
var engineClassStruct_Once sync.Once

func engineClassStruct_Set() error {
	var err error
	engineClassStruct_Once.Do(func() {
		engineClassStruct, err = gi.StructNew("Pango", "EngineClass")
	})
	return err
}

type EngineClass struct {
	native uintptr
}

var engineInfoStruct *gi.Struct
var engineInfoStruct_Once sync.Once

func engineInfoStruct_Set() error {
	var err error
	engineInfoStruct_Once.Do(func() {
		engineInfoStruct, err = gi.StructNew("Pango", "EngineInfo")
	})
	return err
}

type EngineInfo struct {
	native     uintptr
	Id         string
	EngineType string
	RenderType string
	Scripts    *EngineScriptInfo
	NScripts   int32
}

var engineLangClassStruct *gi.Struct
var engineLangClassStruct_Once sync.Once

func engineLangClassStruct_Set() error {
	var err error
	engineLangClassStruct_Once.Do(func() {
		engineLangClassStruct, err = gi.StructNew("Pango", "EngineLangClass")
	})
	return err
}

type EngineLangClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_break' : missing Type
}

var engineScriptInfoStruct *gi.Struct
var engineScriptInfoStruct_Once sync.Once

func engineScriptInfoStruct_Set() error {
	var err error
	engineScriptInfoStruct_Once.Do(func() {
		engineScriptInfoStruct, err = gi.StructNew("Pango", "EngineScriptInfo")
	})
	return err
}

type EngineScriptInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'script' : no Go type for 'Script'
	Langs string
}

var engineShapeClassStruct *gi.Struct
var engineShapeClassStruct_Once sync.Once

func engineShapeClassStruct_Set() error {
	var err error
	engineShapeClassStruct_Once.Do(func() {
		engineShapeClassStruct, err = gi.StructNew("Pango", "EngineShapeClass")
	})
	return err
}

type EngineShapeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_shape' : missing Type
	// UNSUPPORTED : C value 'covers' : missing Type
}

var fontClassStruct *gi.Struct
var fontClassStruct_Once sync.Once

func fontClassStruct_Set() error {
	var err error
	fontClassStruct_Once.Do(func() {
		fontClassStruct, err = gi.StructNew("Pango", "FontClass")
	})
	return err
}

type FontClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'describe' : missing Type
	// UNSUPPORTED : C value 'get_coverage' : missing Type
	// UNSUPPORTED : C value 'find_shaper' : missing Type
	// UNSUPPORTED : C value 'get_glyph_extents' : missing Type
	// UNSUPPORTED : C value 'get_metrics' : missing Type
	// UNSUPPORTED : C value 'get_font_map' : missing Type
	// UNSUPPORTED : C value 'describe_absolute' : missing Type
	// UNSUPPORTED : C value '_pango_reserved1' : missing Type
	// UNSUPPORTED : C value '_pango_reserved2' : missing Type
}

var fontDescriptionStruct *gi.Struct
var fontDescriptionStruct_Once sync.Once

func fontDescriptionStruct_Set() error {
	var err error
	fontDescriptionStruct_Once.Do(func() {
		fontDescriptionStruct, err = gi.StructNew("Pango", "FontDescription")
	})
	return err
}

type FontDescription struct {
	native uintptr
}

var fontDescriptionNewFunction *gi.Function
var fontDescriptionNewFunction_Once sync.Once

func fontDescriptionNewFunction_Set() error {
	var err error
	fontDescriptionNewFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionNewFunction, err = fontDescriptionStruct.InvokerNew("new")
	})
	return err
}

// FontDescriptionNew is a representation of the C type pango_font_description_new.
func FontDescriptionNew() *FontDescription {

	var ret gi.Argument

	err := fontDescriptionNewFunction_Set()
	if err == nil {
		ret = fontDescriptionNewFunction.Invoke(nil, nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_better_match' : parameter 'old_match' of type 'FontDescription' not supported

var fontDescriptionCopyFunction *gi.Function
var fontDescriptionCopyFunction_Once sync.Once

func fontDescriptionCopyFunction_Set() error {
	var err error
	fontDescriptionCopyFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionCopyFunction, err = fontDescriptionStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_font_description_copy.
func (recv *FontDescription) Copy() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionCopyFunction_Set()
	if err == nil {
		ret = fontDescriptionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var fontDescriptionCopyStaticFunction *gi.Function
var fontDescriptionCopyStaticFunction_Once sync.Once

func fontDescriptionCopyStaticFunction_Set() error {
	var err error
	fontDescriptionCopyStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionCopyStaticFunction, err = fontDescriptionStruct.InvokerNew("copy_static")
	})
	return err
}

// CopyStatic is a representation of the C type pango_font_description_copy_static.
func (recv *FontDescription) CopyStatic() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionCopyStaticFunction_Set()
	if err == nil {
		ret = fontDescriptionCopyStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_equal' : parameter 'desc2' of type 'FontDescription' not supported

var fontDescriptionFreeFunction *gi.Function
var fontDescriptionFreeFunction_Once sync.Once

func fontDescriptionFreeFunction_Set() error {
	var err error
	fontDescriptionFreeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionFreeFunction, err = fontDescriptionStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_font_description_free.
func (recv *FontDescription) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontDescriptionFreeFunction_Set()
	if err == nil {
		fontDescriptionFreeFunction.Invoke(inArgs[:], nil)
	}

}

var fontDescriptionGetFamilyFunction *gi.Function
var fontDescriptionGetFamilyFunction_Once sync.Once

func fontDescriptionGetFamilyFunction_Set() error {
	var err error
	fontDescriptionGetFamilyFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetFamilyFunction, err = fontDescriptionStruct.InvokerNew("get_family")
	})
	return err
}

// GetFamily is a representation of the C type pango_font_description_get_family.
func (recv *FontDescription) GetFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetFamilyFunction_Set()
	if err == nil {
		ret = fontDescriptionGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_gravity' : return type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_font_description_get_set_fields' : return type 'FontMask' not supported

var fontDescriptionGetSizeFunction *gi.Function
var fontDescriptionGetSizeFunction_Once sync.Once

func fontDescriptionGetSizeFunction_Set() error {
	var err error
	fontDescriptionGetSizeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetSizeFunction, err = fontDescriptionStruct.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type pango_font_description_get_size.
func (recv *FontDescription) GetSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetSizeFunction_Set()
	if err == nil {
		ret = fontDescriptionGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_size_is_absolute' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_font_description_get_stretch' : return type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_get_style' : return type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_get_variant' : return type 'Variant' not supported

var fontDescriptionGetVariationsFunction *gi.Function
var fontDescriptionGetVariationsFunction_Once sync.Once

func fontDescriptionGetVariationsFunction_Set() error {
	var err error
	fontDescriptionGetVariationsFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetVariationsFunction, err = fontDescriptionStruct.InvokerNew("get_variations")
	})
	return err
}

// GetVariations is a representation of the C type pango_font_description_get_variations.
func (recv *FontDescription) GetVariations() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetVariationsFunction_Set()
	if err == nil {
		ret = fontDescriptionGetVariationsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_weight' : return type 'Weight' not supported

var fontDescriptionHashFunction *gi.Function
var fontDescriptionHashFunction_Once sync.Once

func fontDescriptionHashFunction_Set() error {
	var err error
	fontDescriptionHashFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionHashFunction, err = fontDescriptionStruct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type pango_font_description_hash.
func (recv *FontDescription) Hash() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionHashFunction_Set()
	if err == nil {
		ret = fontDescriptionHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_merge' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_merge_static' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_set_absolute_size' : parameter 'size' of type 'gdouble' not supported

var fontDescriptionSetFamilyFunction *gi.Function
var fontDescriptionSetFamilyFunction_Once sync.Once

func fontDescriptionSetFamilyFunction_Set() error {
	var err error
	fontDescriptionSetFamilyFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetFamilyFunction, err = fontDescriptionStruct.InvokerNew("set_family")
	})
	return err
}

// SetFamily is a representation of the C type pango_font_description_set_family.
func (recv *FontDescription) SetFamily(family string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	err := fontDescriptionSetFamilyFunction_Set()
	if err == nil {
		fontDescriptionSetFamilyFunction.Invoke(inArgs[:], nil)
	}

}

var fontDescriptionSetFamilyStaticFunction *gi.Function
var fontDescriptionSetFamilyStaticFunction_Once sync.Once

func fontDescriptionSetFamilyStaticFunction_Set() error {
	var err error
	fontDescriptionSetFamilyStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetFamilyStaticFunction, err = fontDescriptionStruct.InvokerNew("set_family_static")
	})
	return err
}

// SetFamilyStatic is a representation of the C type pango_font_description_set_family_static.
func (recv *FontDescription) SetFamilyStatic(family string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	err := fontDescriptionSetFamilyStaticFunction_Set()
	if err == nil {
		fontDescriptionSetFamilyStaticFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_font_description_set_gravity' : parameter 'gravity' of type 'Gravity' not supported

var fontDescriptionSetSizeFunction *gi.Function
var fontDescriptionSetSizeFunction_Once sync.Once

func fontDescriptionSetSizeFunction_Set() error {
	var err error
	fontDescriptionSetSizeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetSizeFunction, err = fontDescriptionStruct.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type pango_font_description_set_size.
func (recv *FontDescription) SetSize(size int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(size)

	err := fontDescriptionSetSizeFunction_Set()
	if err == nil {
		fontDescriptionSetSizeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_font_description_set_stretch' : parameter 'stretch' of type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_set_style' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_set_variant' : parameter 'variant' of type 'Variant' not supported

var fontDescriptionSetVariationsFunction *gi.Function
var fontDescriptionSetVariationsFunction_Once sync.Once

func fontDescriptionSetVariationsFunction_Set() error {
	var err error
	fontDescriptionSetVariationsFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetVariationsFunction, err = fontDescriptionStruct.InvokerNew("set_variations")
	})
	return err
}

// SetVariations is a representation of the C type pango_font_description_set_variations.
func (recv *FontDescription) SetVariations(settings string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	err := fontDescriptionSetVariationsFunction_Set()
	if err == nil {
		fontDescriptionSetVariationsFunction.Invoke(inArgs[:], nil)
	}

}

var fontDescriptionSetVariationsStaticFunction *gi.Function
var fontDescriptionSetVariationsStaticFunction_Once sync.Once

func fontDescriptionSetVariationsStaticFunction_Set() error {
	var err error
	fontDescriptionSetVariationsStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetVariationsStaticFunction, err = fontDescriptionStruct.InvokerNew("set_variations_static")
	})
	return err
}

// SetVariationsStatic is a representation of the C type pango_font_description_set_variations_static.
func (recv *FontDescription) SetVariationsStatic(settings string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	err := fontDescriptionSetVariationsStaticFunction_Set()
	if err == nil {
		fontDescriptionSetVariationsStaticFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_font_description_set_weight' : parameter 'weight' of type 'Weight' not supported

var fontDescriptionToFilenameFunction *gi.Function
var fontDescriptionToFilenameFunction_Once sync.Once

func fontDescriptionToFilenameFunction_Set() error {
	var err error
	fontDescriptionToFilenameFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionToFilenameFunction, err = fontDescriptionStruct.InvokerNew("to_filename")
	})
	return err
}

// ToFilename is a representation of the C type pango_font_description_to_filename.
func (recv *FontDescription) ToFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionToFilenameFunction_Set()
	if err == nil {
		ret = fontDescriptionToFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fontDescriptionToStringFunction *gi.Function
var fontDescriptionToStringFunction_Once sync.Once

func fontDescriptionToStringFunction_Set() error {
	var err error
	fontDescriptionToStringFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionToStringFunction, err = fontDescriptionStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type pango_font_description_to_string.
func (recv *FontDescription) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionToStringFunction_Set()
	if err == nil {
		ret = fontDescriptionToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_unset_fields' : parameter 'to_unset' of type 'FontMask' not supported

var fontFaceClassStruct *gi.Struct
var fontFaceClassStruct_Once sync.Once

func fontFaceClassStruct_Set() error {
	var err error
	fontFaceClassStruct_Once.Do(func() {
		fontFaceClassStruct, err = gi.StructNew("Pango", "FontFaceClass")
	})
	return err
}

type FontFaceClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_face_name' : missing Type
	// UNSUPPORTED : C value 'describe' : missing Type
	// UNSUPPORTED : C value 'list_sizes' : missing Type
	// UNSUPPORTED : C value 'is_synthesized' : missing Type
	// UNSUPPORTED : C value '_pango_reserved3' : missing Type
	// UNSUPPORTED : C value '_pango_reserved4' : missing Type
}

var fontFamilyClassStruct *gi.Struct
var fontFamilyClassStruct_Once sync.Once

func fontFamilyClassStruct_Set() error {
	var err error
	fontFamilyClassStruct_Once.Do(func() {
		fontFamilyClassStruct, err = gi.StructNew("Pango", "FontFamilyClass")
	})
	return err
}

type FontFamilyClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'list_faces' : missing Type
	// UNSUPPORTED : C value 'get_name' : missing Type
	// UNSUPPORTED : C value 'is_monospace' : missing Type
	// UNSUPPORTED : C value '_pango_reserved2' : missing Type
	// UNSUPPORTED : C value '_pango_reserved3' : missing Type
	// UNSUPPORTED : C value '_pango_reserved4' : missing Type
}

var fontMapClassStruct *gi.Struct
var fontMapClassStruct_Once sync.Once

func fontMapClassStruct_Set() error {
	var err error
	fontMapClassStruct_Once.Do(func() {
		fontMapClassStruct, err = gi.StructNew("Pango", "FontMapClass")
	})
	return err
}

type FontMapClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'load_font' : missing Type
	// UNSUPPORTED : C value 'list_families' : missing Type
	// UNSUPPORTED : C value 'load_fontset' : missing Type
	ShapeEngineType string
	// UNSUPPORTED : C value 'get_serial' : missing Type
	// UNSUPPORTED : C value 'changed' : missing Type
	// UNSUPPORTED : C value '_pango_reserved1' : missing Type
	// UNSUPPORTED : C value '_pango_reserved2' : missing Type
}

var fontMetricsStruct *gi.Struct
var fontMetricsStruct_Once sync.Once

func fontMetricsStruct_Set() error {
	var err error
	fontMetricsStruct_Once.Do(func() {
		fontMetricsStruct, err = gi.StructNew("Pango", "FontMetrics")
	})
	return err
}

type FontMetrics struct {
	native uintptr
}

var fontMetricsNewFunction *gi.Function
var fontMetricsNewFunction_Once sync.Once

func fontMetricsNewFunction_Set() error {
	var err error
	fontMetricsNewFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsNewFunction, err = fontMetricsStruct.InvokerNew("new")
	})
	return err
}

// FontMetricsNew is a representation of the C type pango_font_metrics_new.
func FontMetricsNew() *FontMetrics {

	var ret gi.Argument

	err := fontMetricsNewFunction_Set()
	if err == nil {
		ret = fontMetricsNewFunction.Invoke(nil, nil)
	}

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var fontMetricsGetApproximateCharWidthFunction *gi.Function
var fontMetricsGetApproximateCharWidthFunction_Once sync.Once

func fontMetricsGetApproximateCharWidthFunction_Set() error {
	var err error
	fontMetricsGetApproximateCharWidthFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetApproximateCharWidthFunction, err = fontMetricsStruct.InvokerNew("get_approximate_char_width")
	})
	return err
}

// GetApproximateCharWidth is a representation of the C type pango_font_metrics_get_approximate_char_width.
func (recv *FontMetrics) GetApproximateCharWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetApproximateCharWidthFunction_Set()
	if err == nil {
		ret = fontMetricsGetApproximateCharWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetApproximateDigitWidthFunction *gi.Function
var fontMetricsGetApproximateDigitWidthFunction_Once sync.Once

func fontMetricsGetApproximateDigitWidthFunction_Set() error {
	var err error
	fontMetricsGetApproximateDigitWidthFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetApproximateDigitWidthFunction, err = fontMetricsStruct.InvokerNew("get_approximate_digit_width")
	})
	return err
}

// GetApproximateDigitWidth is a representation of the C type pango_font_metrics_get_approximate_digit_width.
func (recv *FontMetrics) GetApproximateDigitWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetApproximateDigitWidthFunction_Set()
	if err == nil {
		ret = fontMetricsGetApproximateDigitWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetAscentFunction *gi.Function
var fontMetricsGetAscentFunction_Once sync.Once

func fontMetricsGetAscentFunction_Set() error {
	var err error
	fontMetricsGetAscentFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetAscentFunction, err = fontMetricsStruct.InvokerNew("get_ascent")
	})
	return err
}

// GetAscent is a representation of the C type pango_font_metrics_get_ascent.
func (recv *FontMetrics) GetAscent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetAscentFunction_Set()
	if err == nil {
		ret = fontMetricsGetAscentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetDescentFunction *gi.Function
var fontMetricsGetDescentFunction_Once sync.Once

func fontMetricsGetDescentFunction_Set() error {
	var err error
	fontMetricsGetDescentFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetDescentFunction, err = fontMetricsStruct.InvokerNew("get_descent")
	})
	return err
}

// GetDescent is a representation of the C type pango_font_metrics_get_descent.
func (recv *FontMetrics) GetDescent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetDescentFunction_Set()
	if err == nil {
		ret = fontMetricsGetDescentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughPositionFunction *gi.Function
var fontMetricsGetStrikethroughPositionFunction_Once sync.Once

func fontMetricsGetStrikethroughPositionFunction_Set() error {
	var err error
	fontMetricsGetStrikethroughPositionFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetStrikethroughPositionFunction, err = fontMetricsStruct.InvokerNew("get_strikethrough_position")
	})
	return err
}

// GetStrikethroughPosition is a representation of the C type pango_font_metrics_get_strikethrough_position.
func (recv *FontMetrics) GetStrikethroughPosition() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetStrikethroughPositionFunction_Set()
	if err == nil {
		ret = fontMetricsGetStrikethroughPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughThicknessFunction *gi.Function
var fontMetricsGetStrikethroughThicknessFunction_Once sync.Once

func fontMetricsGetStrikethroughThicknessFunction_Set() error {
	var err error
	fontMetricsGetStrikethroughThicknessFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetStrikethroughThicknessFunction, err = fontMetricsStruct.InvokerNew("get_strikethrough_thickness")
	})
	return err
}

// GetStrikethroughThickness is a representation of the C type pango_font_metrics_get_strikethrough_thickness.
func (recv *FontMetrics) GetStrikethroughThickness() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetStrikethroughThicknessFunction_Set()
	if err == nil {
		ret = fontMetricsGetStrikethroughThicknessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlinePositionFunction *gi.Function
var fontMetricsGetUnderlinePositionFunction_Once sync.Once

func fontMetricsGetUnderlinePositionFunction_Set() error {
	var err error
	fontMetricsGetUnderlinePositionFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetUnderlinePositionFunction, err = fontMetricsStruct.InvokerNew("get_underline_position")
	})
	return err
}

// GetUnderlinePosition is a representation of the C type pango_font_metrics_get_underline_position.
func (recv *FontMetrics) GetUnderlinePosition() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetUnderlinePositionFunction_Set()
	if err == nil {
		ret = fontMetricsGetUnderlinePositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlineThicknessFunction *gi.Function
var fontMetricsGetUnderlineThicknessFunction_Once sync.Once

func fontMetricsGetUnderlineThicknessFunction_Set() error {
	var err error
	fontMetricsGetUnderlineThicknessFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetUnderlineThicknessFunction, err = fontMetricsStruct.InvokerNew("get_underline_thickness")
	})
	return err
}

// GetUnderlineThickness is a representation of the C type pango_font_metrics_get_underline_thickness.
func (recv *FontMetrics) GetUnderlineThickness() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetUnderlineThicknessFunction_Set()
	if err == nil {
		ret = fontMetricsGetUnderlineThicknessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsRefFunction *gi.Function
var fontMetricsRefFunction_Once sync.Once

func fontMetricsRefFunction_Set() error {
	var err error
	fontMetricsRefFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsRefFunction, err = fontMetricsStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_font_metrics_ref.
func (recv *FontMetrics) Ref() *FontMetrics {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsRefFunction_Set()
	if err == nil {
		ret = fontMetricsRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var fontMetricsUnrefFunction *gi.Function
var fontMetricsUnrefFunction_Once sync.Once

func fontMetricsUnrefFunction_Set() error {
	var err error
	fontMetricsUnrefFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsUnrefFunction, err = fontMetricsStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_font_metrics_unref.
func (recv *FontMetrics) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontMetricsUnrefFunction_Set()
	if err == nil {
		fontMetricsUnrefFunction.Invoke(inArgs[:], nil)
	}

}

var fontsetClassStruct *gi.Struct
var fontsetClassStruct_Once sync.Once

func fontsetClassStruct_Set() error {
	var err error
	fontsetClassStruct_Once.Do(func() {
		fontsetClassStruct, err = gi.StructNew("Pango", "FontsetClass")
	})
	return err
}

type FontsetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'get_font' : missing Type
	// UNSUPPORTED : C value 'get_metrics' : missing Type
	// UNSUPPORTED : C value 'get_language' : missing Type
	// UNSUPPORTED : C value 'foreach' : missing Type
	// UNSUPPORTED : C value '_pango_reserved1' : missing Type
	// UNSUPPORTED : C value '_pango_reserved2' : missing Type
	// UNSUPPORTED : C value '_pango_reserved3' : missing Type
	// UNSUPPORTED : C value '_pango_reserved4' : missing Type
}

var fontsetSimpleClassStruct *gi.Struct
var fontsetSimpleClassStruct_Once sync.Once

func fontsetSimpleClassStruct_Set() error {
	var err error
	fontsetSimpleClassStruct_Once.Do(func() {
		fontsetSimpleClassStruct, err = gi.StructNew("Pango", "FontsetSimpleClass")
	})
	return err
}

type FontsetSimpleClass struct {
	native uintptr
}

var glyphGeometryStruct *gi.Struct
var glyphGeometryStruct_Once sync.Once

func glyphGeometryStruct_Set() error {
	var err error
	glyphGeometryStruct_Once.Do(func() {
		glyphGeometryStruct, err = gi.StructNew("Pango", "GlyphGeometry")
	})
	return err
}

type GlyphGeometry struct {
	native  uintptr
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

var glyphInfoStruct *gi.Struct
var glyphInfoStruct_Once sync.Once

func glyphInfoStruct_Set() error {
	var err error
	glyphInfoStruct_Once.Do(func() {
		glyphInfoStruct, err = gi.StructNew("Pango", "GlyphInfo")
	})
	return err
}

type GlyphInfo struct {
	native   uintptr
	Glyph    Glyph
	Geometry *GlyphGeometry
	Attr     *GlyphVisAttr
}

var glyphItemStruct *gi.Struct
var glyphItemStruct_Once sync.Once

func glyphItemStruct_Set() error {
	var err error
	glyphItemStruct_Once.Do(func() {
		glyphItemStruct, err = gi.StructNew("Pango", "GlyphItem")
	})
	return err
}

type GlyphItem struct {
	native uintptr
	Item   *Item
	Glyphs *GlyphString
}

// UNSUPPORTED : C value 'pango_glyph_item_apply_attrs' : parameter 'list' of type 'AttrList' not supported

var glyphItemCopyFunction *gi.Function
var glyphItemCopyFunction_Once sync.Once

func glyphItemCopyFunction_Set() error {
	var err error
	glyphItemCopyFunction_Once.Do(func() {
		err = glyphItemStruct_Set()
		if err != nil {
			return
		}
		glyphItemCopyFunction, err = glyphItemStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_glyph_item_copy.
func (recv *GlyphItem) Copy() *GlyphItem {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemCopyFunction_Set()
	if err == nil {
		ret = glyphItemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

var glyphItemFreeFunction *gi.Function
var glyphItemFreeFunction_Once sync.Once

func glyphItemFreeFunction_Set() error {
	var err error
	glyphItemFreeFunction_Once.Do(func() {
		err = glyphItemStruct_Set()
		if err != nil {
			return
		}
		glyphItemFreeFunction, err = glyphItemStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphItemFreeFunction_Set()
	if err == nil {
		glyphItemFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_glyph_item_get_logical_widths' : parameter 'logical_widths' has no type

// UNSUPPORTED : C value 'pango_glyph_item_letter_space' : parameter 'log_attrs' has no type

var glyphItemSplitFunction *gi.Function
var glyphItemSplitFunction_Once sync.Once

func glyphItemSplitFunction_Set() error {
	var err error
	glyphItemSplitFunction_Once.Do(func() {
		err = glyphItemStruct_Set()
		if err != nil {
			return
		}
		glyphItemSplitFunction, err = glyphItemStruct.InvokerNew("split")
	})
	return err
}

// Split is a representation of the C type pango_glyph_item_split.
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(splitIndex)

	var ret gi.Argument

	err := glyphItemSplitFunction_Set()
	if err == nil {
		ret = glyphItemSplitFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

var glyphItemIterStruct *gi.Struct
var glyphItemIterStruct_Once sync.Once

func glyphItemIterStruct_Set() error {
	var err error
	glyphItemIterStruct_Once.Do(func() {
		glyphItemIterStruct, err = gi.StructNew("Pango", "GlyphItemIter")
	})
	return err
}

type GlyphItemIter struct {
	native     uintptr
	GlyphItem  *GlyphItem
	Text       string
	StartGlyph int32
	StartIndex int32
	StartChar  int32
	EndGlyph   int32
	EndIndex   int32
	EndChar    int32
}

var glyphItemIterCopyFunction *gi.Function
var glyphItemIterCopyFunction_Once sync.Once

func glyphItemIterCopyFunction_Set() error {
	var err error
	glyphItemIterCopyFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterCopyFunction, err = glyphItemIterStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemIterCopyFunction_Set()
	if err == nil {
		ret = glyphItemIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GlyphItemIter{native: ret.Pointer()}

	return retGo
}

var glyphItemIterFreeFunction *gi.Function
var glyphItemIterFreeFunction_Once sync.Once

func glyphItemIterFreeFunction_Set() error {
	var err error
	glyphItemIterFreeFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterFreeFunction, err = glyphItemIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_glyph_item_iter_free.
func (recv *GlyphItemIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphItemIterFreeFunction_Set()
	if err == nil {
		glyphItemIterFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_end' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_start' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_next_cluster' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_prev_cluster' : return type 'gboolean' not supported

var glyphStringStruct *gi.Struct
var glyphStringStruct_Once sync.Once

func glyphStringStruct_Set() error {
	var err error
	glyphStringStruct_Once.Do(func() {
		glyphStringStruct, err = gi.StructNew("Pango", "GlyphString")
	})
	return err
}

type GlyphString struct {
	native    uintptr
	NumGlyphs int32
	// UNSUPPORTED : C value 'glyphs' : missing Type
	LogClusters int32
}

var glyphStringNewFunction *gi.Function
var glyphStringNewFunction_Once sync.Once

func glyphStringNewFunction_Set() error {
	var err error
	glyphStringNewFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringNewFunction, err = glyphStringStruct.InvokerNew("new")
	})
	return err
}

// GlyphStringNew is a representation of the C type pango_glyph_string_new.
func GlyphStringNew() *GlyphString {

	var ret gi.Argument

	err := glyphStringNewFunction_Set()
	if err == nil {
		ret = glyphStringNewFunction.Invoke(nil, nil)
	}

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

var glyphStringCopyFunction *gi.Function
var glyphStringCopyFunction_Once sync.Once

func glyphStringCopyFunction_Set() error {
	var err error
	glyphStringCopyFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringCopyFunction, err = glyphStringStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_glyph_string_copy.
func (recv *GlyphString) Copy() *GlyphString {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphStringCopyFunction_Set()
	if err == nil {
		ret = glyphStringCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_extents' : parameter 'font' of type 'Font' not supported

// UNSUPPORTED : C value 'pango_glyph_string_extents_range' : parameter 'font' of type 'Font' not supported

var glyphStringFreeFunction *gi.Function
var glyphStringFreeFunction_Once sync.Once

func glyphStringFreeFunction_Set() error {
	var err error
	glyphStringFreeFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringFreeFunction, err = glyphStringStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_glyph_string_free.
func (recv *GlyphString) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphStringFreeFunction_Set()
	if err == nil {
		glyphStringFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_glyph_string_get_logical_widths' : parameter 'logical_widths' has no type

var glyphStringGetWidthFunction *gi.Function
var glyphStringGetWidthFunction_Once sync.Once

func glyphStringGetWidthFunction_Set() error {
	var err error
	glyphStringGetWidthFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringGetWidthFunction, err = glyphStringStruct.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type pango_glyph_string_get_width.
func (recv *GlyphString) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphStringGetWidthFunction_Set()
	if err == nil {
		ret = glyphStringGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_index_to_x' : parameter 'analysis' of type 'Analysis' not supported

var glyphStringSetSizeFunction *gi.Function
var glyphStringSetSizeFunction_Once sync.Once

func glyphStringSetSizeFunction_Set() error {
	var err error
	glyphStringSetSizeFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringSetSizeFunction, err = glyphStringStruct.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type pango_glyph_string_set_size.
func (recv *GlyphString) SetSize(newLen int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newLen)

	err := glyphStringSetSizeFunction_Set()
	if err == nil {
		glyphStringSetSizeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_glyph_string_x_to_index' : parameter 'analysis' of type 'Analysis' not supported

var glyphVisAttrStruct *gi.Struct
var glyphVisAttrStruct_Once sync.Once

func glyphVisAttrStruct_Set() error {
	var err error
	glyphVisAttrStruct_Once.Do(func() {
		glyphVisAttrStruct, err = gi.StructNew("Pango", "GlyphVisAttr")
	})
	return err
}

type GlyphVisAttr struct {
	native         uintptr
	IsClusterStart uint32
}

var includedModuleStruct *gi.Struct
var includedModuleStruct_Once sync.Once

func includedModuleStruct_Set() error {
	var err error
	includedModuleStruct_Once.Do(func() {
		includedModuleStruct, err = gi.StructNew("Pango", "IncludedModule")
	})
	return err
}

type IncludedModule struct {
	native uintptr
	// UNSUPPORTED : C value 'list' : missing Type
	// UNSUPPORTED : C value 'init' : missing Type
	// UNSUPPORTED : C value 'exit' : missing Type
	// UNSUPPORTED : C value 'create' : missing Type
}

var itemStruct *gi.Struct
var itemStruct_Once sync.Once

func itemStruct_Set() error {
	var err error
	itemStruct_Once.Do(func() {
		itemStruct, err = gi.StructNew("Pango", "Item")
	})
	return err
}

type Item struct {
	native   uintptr
	Offset   int32
	Length   int32
	NumChars int32
	Analysis *Analysis
}

var itemNewFunction *gi.Function
var itemNewFunction_Once sync.Once

func itemNewFunction_Set() error {
	var err error
	itemNewFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemNewFunction, err = itemStruct.InvokerNew("new")
	})
	return err
}

// ItemNew is a representation of the C type pango_item_new.
func ItemNew() *Item {

	var ret gi.Argument

	err := itemNewFunction_Set()
	if err == nil {
		ret = itemNewFunction.Invoke(nil, nil)
	}

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var itemCopyFunction *gi.Function
var itemCopyFunction_Once sync.Once

func itemCopyFunction_Set() error {
	var err error
	itemCopyFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemCopyFunction, err = itemStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_item_copy.
func (recv *Item) Copy() *Item {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := itemCopyFunction_Set()
	if err == nil {
		ret = itemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var itemFreeFunction *gi.Function
var itemFreeFunction_Once sync.Once

func itemFreeFunction_Set() error {
	var err error
	itemFreeFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemFreeFunction, err = itemStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_item_free.
func (recv *Item) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := itemFreeFunction_Set()
	if err == nil {
		itemFreeFunction.Invoke(inArgs[:], nil)
	}

}

var itemSplitFunction *gi.Function
var itemSplitFunction_Once sync.Once

func itemSplitFunction_Set() error {
	var err error
	itemSplitFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemSplitFunction, err = itemStruct.InvokerNew("split")
	})
	return err
}

// Split is a representation of the C type pango_item_split.
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(splitIndex)
	inArgs[2].SetInt32(splitOffset)

	var ret gi.Argument

	err := itemSplitFunction_Set()
	if err == nil {
		ret = itemSplitFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var languageStruct *gi.Struct
var languageStruct_Once sync.Once

func languageStruct_Set() error {
	var err error
	languageStruct_Once.Do(func() {
		languageStruct, err = gi.StructNew("Pango", "Language")
	})
	return err
}

type Language struct {
	native uintptr
}

var languageGetSampleStringFunction *gi.Function
var languageGetSampleStringFunction_Once sync.Once

func languageGetSampleStringFunction_Set() error {
	var err error
	languageGetSampleStringFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetSampleStringFunction, err = languageStruct.InvokerNew("get_sample_string")
	})
	return err
}

// GetSampleString is a representation of the C type pango_language_get_sample_string.
func (recv *Language) GetSampleString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := languageGetSampleStringFunction_Set()
	if err == nil {
		ret = languageGetSampleStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageGetScriptsFunction *gi.Function
var languageGetScriptsFunction_Once sync.Once

func languageGetScriptsFunction_Set() error {
	var err error
	languageGetScriptsFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetScriptsFunction, err = languageStruct.InvokerNew("get_scripts")
	})
	return err
}

// GetScripts is a representation of the C type pango_language_get_scripts.
func (recv *Language) GetScripts() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := languageGetScriptsFunction_Set()
	if err == nil {
		languageGetScriptsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'pango_language_includes_script' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_language_matches' : return type 'gboolean' not supported

var languageToStringFunction *gi.Function
var languageToStringFunction_Once sync.Once

func languageToStringFunction_Set() error {
	var err error
	languageToStringFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageToStringFunction, err = languageStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type pango_language_to_string.
func (recv *Language) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := languageToStringFunction_Set()
	if err == nil {
		ret = languageToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var layoutClassStruct *gi.Struct
var layoutClassStruct_Once sync.Once

func layoutClassStruct_Set() error {
	var err error
	layoutClassStruct_Once.Do(func() {
		layoutClassStruct, err = gi.StructNew("Pango", "LayoutClass")
	})
	return err
}

type LayoutClass struct {
	native uintptr
}

var layoutIterStruct *gi.Struct
var layoutIterStruct_Once sync.Once

func layoutIterStruct_Set() error {
	var err error
	layoutIterStruct_Once.Do(func() {
		layoutIterStruct, err = gi.StructNew("Pango", "LayoutIter")
	})
	return err
}

type LayoutIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_layout_iter_at_last_line' : return type 'gboolean' not supported

var layoutIterCopyFunction *gi.Function
var layoutIterCopyFunction_Once sync.Once

func layoutIterCopyFunction_Set() error {
	var err error
	layoutIterCopyFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterCopyFunction, err = layoutIterStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_layout_iter_copy.
func (recv *LayoutIter) Copy() *LayoutIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterCopyFunction_Set()
	if err == nil {
		ret = layoutIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutIter{native: ret.Pointer()}

	return retGo
}

var layoutIterFreeFunction *gi.Function
var layoutIterFreeFunction_Once sync.Once

func layoutIterFreeFunction_Set() error {
	var err error
	layoutIterFreeFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterFreeFunction, err = layoutIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_layout_iter_free.
func (recv *LayoutIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := layoutIterFreeFunction_Set()
	if err == nil {
		layoutIterFreeFunction.Invoke(inArgs[:], nil)
	}

}

var layoutIterGetBaselineFunction *gi.Function
var layoutIterGetBaselineFunction_Once sync.Once

func layoutIterGetBaselineFunction_Set() error {
	var err error
	layoutIterGetBaselineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetBaselineFunction, err = layoutIterStruct.InvokerNew("get_baseline")
	})
	return err
}

// GetBaseline is a representation of the C type pango_layout_iter_get_baseline.
func (recv *LayoutIter) GetBaseline() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetBaselineFunction_Set()
	if err == nil {
		ret = layoutIterGetBaselineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_char_extents' : parameter 'logical_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_cluster_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetIndexFunction *gi.Function
var layoutIterGetIndexFunction_Once sync.Once

func layoutIterGetIndexFunction_Set() error {
	var err error
	layoutIterGetIndexFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetIndexFunction, err = layoutIterStruct.InvokerNew("get_index")
	})
	return err
}

// GetIndex is a representation of the C type pango_layout_iter_get_index.
func (recv *LayoutIter) GetIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetIndexFunction_Set()
	if err == nil {
		ret = layoutIterGetIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_layout' : return type 'Layout' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_layout_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetLineFunction *gi.Function
var layoutIterGetLineFunction_Once sync.Once

func layoutIterGetLineFunction_Set() error {
	var err error
	layoutIterGetLineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineFunction, err = layoutIterStruct.InvokerNew("get_line")
	})
	return err
}

// GetLine is a representation of the C type pango_layout_iter_get_line.
func (recv *LayoutIter) GetLine() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetLineFunction_Set()
	if err == nil {
		ret = layoutIterGetLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_line_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetLineReadonlyFunction *gi.Function
var layoutIterGetLineReadonlyFunction_Once sync.Once

func layoutIterGetLineReadonlyFunction_Set() error {
	var err error
	layoutIterGetLineReadonlyFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineReadonlyFunction, err = layoutIterStruct.InvokerNew("get_line_readonly")
	})
	return err
}

// GetLineReadonly is a representation of the C type pango_layout_iter_get_line_readonly.
func (recv *LayoutIter) GetLineReadonly() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetLineReadonlyFunction_Set()
	if err == nil {
		ret = layoutIterGetLineReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var layoutIterGetLineYrangeFunction *gi.Function
var layoutIterGetLineYrangeFunction_Once sync.Once

func layoutIterGetLineYrangeFunction_Set() error {
	var err error
	layoutIterGetLineYrangeFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineYrangeFunction, err = layoutIterStruct.InvokerNew("get_line_yrange")
	})
	return err
}

// GetLineYrange is a representation of the C type pango_layout_iter_get_line_yrange.
func (recv *LayoutIter) GetLineYrange() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetLineYrangeFunction_Set()
	if err == nil {
		layoutIterGetLineYrangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

// UNSUPPORTED : C value 'pango_layout_iter_get_run' : return type 'LayoutRun' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_run_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_run_readonly' : return type 'LayoutRun' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_char' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_cluster' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_layout_iter_next_run' : return type 'gboolean' not supported

var layoutLineStruct *gi.Struct
var layoutLineStruct_Once sync.Once

func layoutLineStruct_Set() error {
	var err error
	layoutLineStruct_Once.Do(func() {
		layoutLineStruct, err = gi.StructNew("Pango", "LayoutLine")
	})
	return err
}

type LayoutLine struct {
	native uintptr
	// UNSUPPORTED : C value 'layout' : no Go type for 'Layout'
	StartIndex int32
	Length     int32
	// UNSUPPORTED : C value 'runs' : no Go type for 'GLib.SList'
	IsParagraphStart uint32
	ResolvedDir      uint32
}

// UNSUPPORTED : C value 'pango_layout_line_get_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_line_get_pixel_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_line_get_x_ranges' : parameter 'ranges' has no type

// UNSUPPORTED : C value 'pango_layout_line_index_to_x' : parameter 'trailing' of type 'gboolean' not supported

var layoutLineRefFunction *gi.Function
var layoutLineRefFunction_Once sync.Once

func layoutLineRefFunction_Set() error {
	var err error
	layoutLineRefFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineRefFunction, err = layoutLineStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_layout_line_ref.
func (recv *LayoutLine) Ref() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutLineRefFunction_Set()
	if err == nil {
		ret = layoutLineRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var layoutLineUnrefFunction *gi.Function
var layoutLineUnrefFunction_Once sync.Once

func layoutLineUnrefFunction_Set() error {
	var err error
	layoutLineUnrefFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineUnrefFunction, err = layoutLineStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_layout_line_unref.
func (recv *LayoutLine) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := layoutLineUnrefFunction_Set()
	if err == nil {
		layoutLineUnrefFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_layout_line_x_to_index' : return type 'gboolean' not supported

var logAttrStruct *gi.Struct
var logAttrStruct_Once sync.Once

func logAttrStruct_Set() error {
	var err error
	logAttrStruct_Once.Do(func() {
		logAttrStruct, err = gi.StructNew("Pango", "LogAttr")
	})
	return err
}

type LogAttr struct {
	native                    uintptr
	IsLineBreak               uint32
	IsMandatoryBreak          uint32
	IsCharBreak               uint32
	IsWhite                   uint32
	IsCursorPosition          uint32
	IsWordStart               uint32
	IsWordEnd                 uint32
	IsSentenceBoundary        uint32
	IsSentenceStart           uint32
	IsSentenceEnd             uint32
	BackspaceDeletesCharacter uint32
	IsExpandableSpace         uint32
	IsWordBoundary            uint32
}

var mapStruct *gi.Struct
var mapStruct_Once sync.Once

func mapStruct_Set() error {
	var err error
	mapStruct_Once.Do(func() {
		mapStruct, err = gi.StructNew("Pango", "Map")
	})
	return err
}

type Map struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_map_get_engine' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_map_get_engines' : parameter 'script' of type 'Script' not supported

var mapEntryStruct *gi.Struct
var mapEntryStruct_Once sync.Once

func mapEntryStruct_Set() error {
	var err error
	mapEntryStruct_Once.Do(func() {
		mapEntryStruct, err = gi.StructNew("Pango", "MapEntry")
	})
	return err
}

type MapEntry struct {
	native uintptr
}

var matrixStruct *gi.Struct
var matrixStruct_Once sync.Once

func matrixStruct_Set() error {
	var err error
	matrixStruct_Once.Do(func() {
		matrixStruct, err = gi.StructNew("Pango", "Matrix")
	})
	return err
}

type Matrix struct {
	native uintptr
	// UNSUPPORTED : C value 'xx' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'xy' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'yx' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'yy' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'x0' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y0' : no Go type for 'gdouble'
}

// UNSUPPORTED : C value 'pango_matrix_concat' : parameter 'new_matrix' of type 'Matrix' not supported

var matrixCopyFunction *gi.Function
var matrixCopyFunction_Once sync.Once

func matrixCopyFunction_Set() error {
	var err error
	matrixCopyFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixCopyFunction, err = matrixStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matrixCopyFunction_Set()
	if err == nil {
		ret = matrixCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Matrix{native: ret.Pointer()}

	return retGo
}

var matrixFreeFunction *gi.Function
var matrixFreeFunction_Once sync.Once

func matrixFreeFunction_Set() error {
	var err error
	matrixFreeFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixFreeFunction, err = matrixStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_matrix_free.
func (recv *Matrix) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := matrixFreeFunction_Set()
	if err == nil {
		matrixFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_matrix_get_font_scale_factor' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_get_font_scale_factors' : parameter 'xscale' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_rotate' : parameter 'degrees' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_scale' : parameter 'scale_x' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_distance' : parameter 'dx' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_pixel_rectangle' : parameter 'rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_point' : parameter 'x' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_matrix_transform_rectangle' : parameter 'rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_matrix_translate' : parameter 'tx' of type 'gdouble' not supported

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() error {
	var err error
	rectangleStruct_Once.Do(func() {
		rectangleStruct, err = gi.StructNew("Pango", "Rectangle")
	})
	return err
}

type Rectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

var rendererClassStruct *gi.Struct
var rendererClassStruct_Once sync.Once

func rendererClassStruct_Set() error {
	var err error
	rendererClassStruct_Once.Do(func() {
		rendererClassStruct, err = gi.StructNew("Pango", "RendererClass")
	})
	return err
}

type RendererClass struct {
	native uintptr
	// UNSUPPORTED : C value 'draw_glyphs' : missing Type
	// UNSUPPORTED : C value 'draw_rectangle' : missing Type
	// UNSUPPORTED : C value 'draw_error_underline' : missing Type
	// UNSUPPORTED : C value 'draw_shape' : missing Type
	// UNSUPPORTED : C value 'draw_trapezoid' : missing Type
	// UNSUPPORTED : C value 'draw_glyph' : missing Type
	// UNSUPPORTED : C value 'part_changed' : missing Type
	// UNSUPPORTED : C value 'begin' : missing Type
	// UNSUPPORTED : C value 'end' : missing Type
	// UNSUPPORTED : C value 'prepare_run' : missing Type
	// UNSUPPORTED : C value 'draw_glyph_item' : missing Type
	// UNSUPPORTED : C value '_pango_reserved2' : missing Type
	// UNSUPPORTED : C value '_pango_reserved3' : missing Type
	// UNSUPPORTED : C value '_pango_reserved4' : missing Type
}

var rendererPrivateStruct *gi.Struct
var rendererPrivateStruct_Once sync.Once

func rendererPrivateStruct_Set() error {
	var err error
	rendererPrivateStruct_Once.Do(func() {
		rendererPrivateStruct, err = gi.StructNew("Pango", "RendererPrivate")
	})
	return err
}

type RendererPrivate struct {
	native uintptr
}

var scriptIterStruct *gi.Struct
var scriptIterStruct_Once sync.Once

func scriptIterStruct_Set() error {
	var err error
	scriptIterStruct_Once.Do(func() {
		scriptIterStruct, err = gi.StructNew("Pango", "ScriptIter")
	})
	return err
}

type ScriptIter struct {
	native uintptr
}

var scriptIterFreeFunction *gi.Function
var scriptIterFreeFunction_Once sync.Once

func scriptIterFreeFunction_Set() error {
	var err error
	scriptIterFreeFunction_Once.Do(func() {
		err = scriptIterStruct_Set()
		if err != nil {
			return
		}
		scriptIterFreeFunction, err = scriptIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_script_iter_free.
func (recv *ScriptIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := scriptIterFreeFunction_Set()
	if err == nil {
		scriptIterFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_script_iter_get_range' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_script_iter_next' : return type 'gboolean' not supported

var tabArrayStruct *gi.Struct
var tabArrayStruct_Once sync.Once

func tabArrayStruct_Set() error {
	var err error
	tabArrayStruct_Once.Do(func() {
		tabArrayStruct, err = gi.StructNew("Pango", "TabArray")
	})
	return err
}

type TabArray struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_tab_array_new' : parameter 'positions_in_pixels' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_tab_array_new_with_positions' : parameter 'positions_in_pixels' of type 'gboolean' not supported

var tabArrayCopyFunction *gi.Function
var tabArrayCopyFunction_Once sync.Once

func tabArrayCopyFunction_Set() error {
	var err error
	tabArrayCopyFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayCopyFunction, err = tabArrayStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_tab_array_copy.
func (recv *TabArray) Copy() *TabArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayCopyFunction_Set()
	if err == nil {
		ret = tabArrayCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TabArray{native: ret.Pointer()}

	return retGo
}

var tabArrayFreeFunction *gi.Function
var tabArrayFreeFunction_Once sync.Once

func tabArrayFreeFunction_Set() error {
	var err error
	tabArrayFreeFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayFreeFunction, err = tabArrayStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_tab_array_free.
func (recv *TabArray) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := tabArrayFreeFunction_Set()
	if err == nil {
		tabArrayFreeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_tab_array_get_positions_in_pixels' : return type 'gboolean' not supported

var tabArrayGetSizeFunction *gi.Function
var tabArrayGetSizeFunction_Once sync.Once

func tabArrayGetSizeFunction_Set() error {
	var err error
	tabArrayGetSizeFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayGetSizeFunction, err = tabArrayStruct.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type pango_tab_array_get_size.
func (recv *TabArray) GetSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayGetSizeFunction_Set()
	if err == nil {
		ret = tabArrayGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_tab_array_get_tab' : parameter 'alignment' of type 'TabAlign' not supported

// UNSUPPORTED : C value 'pango_tab_array_get_tabs' : parameter 'alignments' of type 'TabAlign' not supported

var tabArrayResizeFunction *gi.Function
var tabArrayResizeFunction_Once sync.Once

func tabArrayResizeFunction_Set() error {
	var err error
	tabArrayResizeFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayResizeFunction, err = tabArrayStruct.InvokerNew("resize")
	})
	return err
}

// Resize is a representation of the C type pango_tab_array_resize.
func (recv *TabArray) Resize(newSize int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newSize)

	err := tabArrayResizeFunction_Set()
	if err == nil {
		tabArrayResizeFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'pango_tab_array_set_tab' : parameter 'alignment' of type 'TabAlign' not supported

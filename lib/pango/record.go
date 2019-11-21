// Code generated - DO NOT EDIT.

package pango

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var analysisStruct *gi.Struct
var analysisStructOnce sync.Once

func analysisStructSet() {
	analysisStructOnce.Do(func() {
		analysisStruct = gi.StructNew("Pango", "Analysis")
	})
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
var attrClassStructOnce sync.Once

func attrClassStructSet() {
	attrClassStructOnce.Do(func() {
		attrClassStruct = gi.StructNew("Pango", "AttrClass")
	})
}

type AttrClass struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'AttrType'
	// UNSUPPORTED : C value 'copy' : missing Type
	// UNSUPPORTED : C value 'destroy' : missing Type
	// UNSUPPORTED : C value 'equal' : missing Type
}

var attrColorStruct *gi.Struct
var attrColorStructOnce sync.Once

func attrColorStructSet() {
	attrColorStructOnce.Do(func() {
		attrColorStruct = gi.StructNew("Pango", "AttrColor")
	})
}

type AttrColor struct {
	native uintptr
	Attr   *Attribute
	Color  *Color
}

var attrFloatStruct *gi.Struct
var attrFloatStructOnce sync.Once

func attrFloatStructSet() {
	attrFloatStructOnce.Do(func() {
		attrFloatStruct = gi.StructNew("Pango", "AttrFloat")
	})
}

type AttrFloat struct {
	native uintptr
	Attr   *Attribute
	// UNSUPPORTED : C value 'value' : no Go type for 'gdouble'
}

var attrFontDescStruct *gi.Struct
var attrFontDescStructOnce sync.Once

func attrFontDescStructSet() {
	attrFontDescStructOnce.Do(func() {
		attrFontDescStruct = gi.StructNew("Pango", "AttrFontDesc")
	})
}

type AttrFontDesc struct {
	native uintptr
	Attr   *Attribute
	Desc   *FontDescription
}

var attrFontFeaturesStruct *gi.Struct
var attrFontFeaturesStructOnce sync.Once

func attrFontFeaturesStructSet() {
	attrFontFeaturesStructOnce.Do(func() {
		attrFontFeaturesStruct = gi.StructNew("Pango", "AttrFontFeatures")
	})
}

type AttrFontFeatures struct {
	native   uintptr
	Attr     *Attribute
	Features string
}

var attrIntStruct *gi.Struct
var attrIntStructOnce sync.Once

func attrIntStructSet() {
	attrIntStructOnce.Do(func() {
		attrIntStruct = gi.StructNew("Pango", "AttrInt")
	})
}

type AttrInt struct {
	native uintptr
	Attr   *Attribute
	Value  int32
}

var attrIteratorStruct *gi.Struct
var attrIteratorStructOnce sync.Once

func attrIteratorStructSet() {
	attrIteratorStructOnce.Do(func() {
		attrIteratorStruct = gi.StructNew("Pango", "AttrIterator")
	})
}

type AttrIterator struct {
	native uintptr
}

var attrIteratorCopyFunction *gi.Function
var attrIteratorCopyFunctionOnce sync.Once

func attrIteratorCopyFunctionSet() {
	attrIteratorCopyFunctionOnce.Do(func() {
		attrIteratorCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyAttrIteratorInvoker *gi.Function

// Copy is a representation of the C type pango_attr_iterator_copy.
func (recv *AttrIterator) Copy() *AttrIterator {
	if copyAttrIteratorInvoker == nil {
		copyAttrIteratorInvoker = gi.StructFunctionInvokerNew("Pango", "AttrIterator", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyAttrIteratorInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

var attrIteratorDestroyFunction *gi.Function
var attrIteratorDestroyFunctionOnce sync.Once

func attrIteratorDestroyFunctionSet() {
	attrIteratorDestroyFunctionOnce.Do(func() {
		attrIteratorDestroyFunction = gi.FunctionInvokerNew("Pango", "destroy")
	})
}

var destroyAttrIteratorInvoker *gi.Function

// Destroy is a representation of the C type pango_attr_iterator_destroy.
func (recv *AttrIterator) Destroy() {
	if destroyAttrIteratorInvoker == nil {
		destroyAttrIteratorInvoker = gi.StructFunctionInvokerNew("Pango", "AttrIterator", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyAttrIteratorInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_attr_iterator_get' : parameter 'type' of type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_attrs' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_font' : parameter 'desc' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_next' : return type 'gboolean' not supported

var attrIteratorRangeFunction *gi.Function
var attrIteratorRangeFunctionOnce sync.Once

func attrIteratorRangeFunctionSet() {
	attrIteratorRangeFunctionOnce.Do(func() {
		attrIteratorRangeFunction = gi.FunctionInvokerNew("Pango", "range")
	})
}

var rangeAttrIteratorInvoker *gi.Function

// Range is a representation of the C type pango_attr_iterator_range.
func (recv *AttrIterator) Range() (int32, int32) {
	if rangeAttrIteratorInvoker == nil {
		rangeAttrIteratorInvoker = gi.StructFunctionInvokerNew("Pango", "AttrIterator", "range")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	rangeAttrIteratorInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var attrLanguageStruct *gi.Struct
var attrLanguageStructOnce sync.Once

func attrLanguageStructSet() {
	attrLanguageStructOnce.Do(func() {
		attrLanguageStruct = gi.StructNew("Pango", "AttrLanguage")
	})
}

type AttrLanguage struct {
	native uintptr
	Attr   *Attribute
	Value  *Language
}

var attrListStruct *gi.Struct
var attrListStructOnce sync.Once

func attrListStructSet() {
	attrListStructOnce.Do(func() {
		attrListStruct = gi.StructNew("Pango", "AttrList")
	})
}

type AttrList struct {
	native uintptr
}

var attrListNewFunction *gi.Function
var attrListNewFunctionOnce sync.Once

func attrListNewFunctionSet() {
	attrListNewFunctionOnce.Do(func() {
		attrListNewFunction = gi.FunctionInvokerNew("Pango", "new")
	})
}

var newAttrListInvoker *gi.Function

// AttrListNew is a representation of the C type pango_attr_list_new.
func AttrListNew() *AttrList {
	if newAttrListInvoker == nil {
		newAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "new")
	}

	ret := newAttrListInvoker.Invoke(nil, nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_change' : parameter 'attr' of type 'Attribute' not supported

var attrListCopyFunction *gi.Function
var attrListCopyFunctionOnce sync.Once

func attrListCopyFunctionSet() {
	attrListCopyFunctionOnce.Do(func() {
		attrListCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyAttrListInvoker *gi.Function

// Copy is a representation of the C type pango_attr_list_copy.
func (recv *AttrList) Copy() *AttrList {
	if copyAttrListInvoker == nil {
		copyAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyAttrListInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_filter' : parameter 'func' of type 'AttrFilterFunc' not supported

var attrListGetIteratorFunction *gi.Function
var attrListGetIteratorFunctionOnce sync.Once

func attrListGetIteratorFunctionSet() {
	attrListGetIteratorFunctionOnce.Do(func() {
		attrListGetIteratorFunction = gi.FunctionInvokerNew("Pango", "get_iterator")
	})
}

var getIteratorAttrListInvoker *gi.Function

// GetIterator is a representation of the C type pango_attr_list_get_iterator.
func (recv *AttrList) GetIterator() *AttrIterator {
	if getIteratorAttrListInvoker == nil {
		getIteratorAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "get_iterator")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIteratorAttrListInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_insert' : parameter 'attr' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_list_insert_before' : parameter 'attr' of type 'Attribute' not supported

var attrListRefFunction *gi.Function
var attrListRefFunctionOnce sync.Once

func attrListRefFunctionSet() {
	attrListRefFunctionOnce.Do(func() {
		attrListRefFunction = gi.FunctionInvokerNew("Pango", "ref")
	})
}

var refAttrListInvoker *gi.Function

// Ref is a representation of the C type pango_attr_list_ref.
func (recv *AttrList) Ref() *AttrList {
	if refAttrListInvoker == nil {
		refAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refAttrListInvoker.Invoke(inArgs[:], nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_splice' : parameter 'other' of type 'AttrList' not supported

var attrListUnrefFunction *gi.Function
var attrListUnrefFunctionOnce sync.Once

func attrListUnrefFunctionSet() {
	attrListUnrefFunctionOnce.Do(func() {
		attrListUnrefFunction = gi.FunctionInvokerNew("Pango", "unref")
	})
}

var unrefAttrListInvoker *gi.Function

// Unref is a representation of the C type pango_attr_list_unref.
func (recv *AttrList) Unref() {
	if unrefAttrListInvoker == nil {
		unrefAttrListInvoker = gi.StructFunctionInvokerNew("Pango", "AttrList", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefAttrListInvoker.Invoke(inArgs[:], nil)

}

var attrShapeStruct *gi.Struct
var attrShapeStructOnce sync.Once

func attrShapeStructSet() {
	attrShapeStructOnce.Do(func() {
		attrShapeStruct = gi.StructNew("Pango", "AttrShape")
	})
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
var attrSizeStructOnce sync.Once

func attrSizeStructSet() {
	attrSizeStructOnce.Do(func() {
		attrSizeStruct = gi.StructNew("Pango", "AttrSize")
	})
}

type AttrSize struct {
	native   uintptr
	Attr     *Attribute
	Size     int32
	Absolute uint32
}

var attrStringStruct *gi.Struct
var attrStringStructOnce sync.Once

func attrStringStructSet() {
	attrStringStructOnce.Do(func() {
		attrStringStruct = gi.StructNew("Pango", "AttrString")
	})
}

type AttrString struct {
	native uintptr
	Attr   *Attribute
	Value  string
}

var attributeStruct *gi.Struct
var attributeStructOnce sync.Once

func attributeStructSet() {
	attributeStructOnce.Do(func() {
		attributeStruct = gi.StructNew("Pango", "Attribute")
	})
}

type Attribute struct {
	native     uintptr
	Klass      *AttrClass
	StartIndex uint32
	EndIndex   uint32
}

var attributeCopyFunction *gi.Function
var attributeCopyFunctionOnce sync.Once

func attributeCopyFunctionSet() {
	attributeCopyFunctionOnce.Do(func() {
		attributeCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyAttributeInvoker *gi.Function

// Copy is a representation of the C type pango_attribute_copy.
func (recv *Attribute) Copy() *Attribute {
	if copyAttributeInvoker == nil {
		copyAttributeInvoker = gi.StructFunctionInvokerNew("Pango", "Attribute", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyAttributeInvoker.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var attributeDestroyFunction *gi.Function
var attributeDestroyFunctionOnce sync.Once

func attributeDestroyFunctionSet() {
	attributeDestroyFunctionOnce.Do(func() {
		attributeDestroyFunction = gi.FunctionInvokerNew("Pango", "destroy")
	})
}

var destroyAttributeInvoker *gi.Function

// Destroy is a representation of the C type pango_attribute_destroy.
func (recv *Attribute) Destroy() {
	if destroyAttributeInvoker == nil {
		destroyAttributeInvoker = gi.StructFunctionInvokerNew("Pango", "Attribute", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyAttributeInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_attribute_equal' : parameter 'attr2' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attribute_init' : parameter 'klass' of type 'AttrClass' not supported

var colorStruct *gi.Struct
var colorStructOnce sync.Once

func colorStructSet() {
	colorStructOnce.Do(func() {
		colorStruct = gi.StructNew("Pango", "Color")
	})
}

type Color struct {
	native uintptr
	Red    uint16
	Green  uint16
	Blue   uint16
}

var colorCopyFunction *gi.Function
var colorCopyFunctionOnce sync.Once

func colorCopyFunctionSet() {
	colorCopyFunctionOnce.Do(func() {
		colorCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyColorInvoker *gi.Function

// Copy is a representation of the C type pango_color_copy.
func (recv *Color) Copy() *Color {
	if copyColorInvoker == nil {
		copyColorInvoker = gi.StructFunctionInvokerNew("Pango", "Color", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyColorInvoker.Invoke(inArgs[:], nil)

	retGo := &Color{native: ret.Pointer()}

	return retGo
}

var colorFreeFunction *gi.Function
var colorFreeFunctionOnce sync.Once

func colorFreeFunctionSet() {
	colorFreeFunctionOnce.Do(func() {
		colorFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeColorInvoker *gi.Function

// Free is a representation of the C type pango_color_free.
func (recv *Color) Free() {
	if freeColorInvoker == nil {
		freeColorInvoker = gi.StructFunctionInvokerNew("Pango", "Color", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeColorInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_color_parse' : return type 'gboolean' not supported

var colorToStringFunction *gi.Function
var colorToStringFunctionOnce sync.Once

func colorToStringFunctionSet() {
	colorToStringFunctionOnce.Do(func() {
		colorToStringFunction = gi.FunctionInvokerNew("Pango", "to_string")
	})
}

var toStringColorInvoker *gi.Function

// ToString is a representation of the C type pango_color_to_string.
func (recv *Color) ToString() string {
	if toStringColorInvoker == nil {
		toStringColorInvoker = gi.StructFunctionInvokerNew("Pango", "Color", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringColorInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var contextClassStruct *gi.Struct
var contextClassStructOnce sync.Once

func contextClassStructSet() {
	contextClassStructOnce.Do(func() {
		contextClassStruct = gi.StructNew("Pango", "ContextClass")
	})
}

type ContextClass struct {
	native uintptr
}

var coverageStruct *gi.Struct
var coverageStructOnce sync.Once

func coverageStructSet() {
	coverageStructOnce.Do(func() {
		coverageStruct = gi.StructNew("Pango", "Coverage")
	})
}

type Coverage struct {
	native uintptr
}

var coverageCopyFunction *gi.Function
var coverageCopyFunctionOnce sync.Once

func coverageCopyFunctionSet() {
	coverageCopyFunctionOnce.Do(func() {
		coverageCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyCoverageInvoker *gi.Function

// Copy is a representation of the C type pango_coverage_copy.
func (recv *Coverage) Copy() *Coverage {
	if copyCoverageInvoker == nil {
		copyCoverageInvoker = gi.StructFunctionInvokerNew("Pango", "Coverage", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyCoverageInvoker.Invoke(inArgs[:], nil)

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_get' : return type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_max' : parameter 'other' of type 'Coverage' not supported

var coverageRefFunction *gi.Function
var coverageRefFunctionOnce sync.Once

func coverageRefFunctionSet() {
	coverageRefFunctionOnce.Do(func() {
		coverageRefFunction = gi.FunctionInvokerNew("Pango", "ref")
	})
}

var refCoverageInvoker *gi.Function

// Ref is a representation of the C type pango_coverage_ref.
func (recv *Coverage) Ref() *Coverage {
	if refCoverageInvoker == nil {
		refCoverageInvoker = gi.StructFunctionInvokerNew("Pango", "Coverage", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refCoverageInvoker.Invoke(inArgs[:], nil)

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_set' : parameter 'level' of type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_to_bytes' : parameter 'bytes' has no type

var coverageUnrefFunction *gi.Function
var coverageUnrefFunctionOnce sync.Once

func coverageUnrefFunctionSet() {
	coverageUnrefFunctionOnce.Do(func() {
		coverageUnrefFunction = gi.FunctionInvokerNew("Pango", "unref")
	})
}

var unrefCoverageInvoker *gi.Function

// Unref is a representation of the C type pango_coverage_unref.
func (recv *Coverage) Unref() {
	if unrefCoverageInvoker == nil {
		unrefCoverageInvoker = gi.StructFunctionInvokerNew("Pango", "Coverage", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefCoverageInvoker.Invoke(inArgs[:], nil)

}

var engineClassStruct *gi.Struct
var engineClassStructOnce sync.Once

func engineClassStructSet() {
	engineClassStructOnce.Do(func() {
		engineClassStruct = gi.StructNew("Pango", "EngineClass")
	})
}

type EngineClass struct {
	native uintptr
}

var engineInfoStruct *gi.Struct
var engineInfoStructOnce sync.Once

func engineInfoStructSet() {
	engineInfoStructOnce.Do(func() {
		engineInfoStruct = gi.StructNew("Pango", "EngineInfo")
	})
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
var engineLangClassStructOnce sync.Once

func engineLangClassStructSet() {
	engineLangClassStructOnce.Do(func() {
		engineLangClassStruct = gi.StructNew("Pango", "EngineLangClass")
	})
}

type EngineLangClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_break' : missing Type
}

var engineScriptInfoStruct *gi.Struct
var engineScriptInfoStructOnce sync.Once

func engineScriptInfoStructSet() {
	engineScriptInfoStructOnce.Do(func() {
		engineScriptInfoStruct = gi.StructNew("Pango", "EngineScriptInfo")
	})
}

type EngineScriptInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'script' : no Go type for 'Script'
	Langs string
}

var engineShapeClassStruct *gi.Struct
var engineShapeClassStructOnce sync.Once

func engineShapeClassStructSet() {
	engineShapeClassStructOnce.Do(func() {
		engineShapeClassStruct = gi.StructNew("Pango", "EngineShapeClass")
	})
}

type EngineShapeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_shape' : missing Type
	// UNSUPPORTED : C value 'covers' : missing Type
}

var fontClassStruct *gi.Struct
var fontClassStructOnce sync.Once

func fontClassStructSet() {
	fontClassStructOnce.Do(func() {
		fontClassStruct = gi.StructNew("Pango", "FontClass")
	})
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
var fontDescriptionStructOnce sync.Once

func fontDescriptionStructSet() {
	fontDescriptionStructOnce.Do(func() {
		fontDescriptionStruct = gi.StructNew("Pango", "FontDescription")
	})
}

type FontDescription struct {
	native uintptr
}

var fontDescriptionNewFunction *gi.Function
var fontDescriptionNewFunctionOnce sync.Once

func fontDescriptionNewFunctionSet() {
	fontDescriptionNewFunctionOnce.Do(func() {
		fontDescriptionNewFunction = gi.FunctionInvokerNew("Pango", "new")
	})
}

var newFontDescriptionInvoker *gi.Function

// FontDescriptionNew is a representation of the C type pango_font_description_new.
func FontDescriptionNew() *FontDescription {
	if newFontDescriptionInvoker == nil {
		newFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "new")
	}

	ret := newFontDescriptionInvoker.Invoke(nil, nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_better_match' : parameter 'old_match' of type 'FontDescription' not supported

var fontDescriptionCopyFunction *gi.Function
var fontDescriptionCopyFunctionOnce sync.Once

func fontDescriptionCopyFunctionSet() {
	fontDescriptionCopyFunctionOnce.Do(func() {
		fontDescriptionCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyFontDescriptionInvoker *gi.Function

// Copy is a representation of the C type pango_font_description_copy.
func (recv *FontDescription) Copy() *FontDescription {
	if copyFontDescriptionInvoker == nil {
		copyFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var fontDescriptionCopyStaticFunction *gi.Function
var fontDescriptionCopyStaticFunctionOnce sync.Once

func fontDescriptionCopyStaticFunctionSet() {
	fontDescriptionCopyStaticFunctionOnce.Do(func() {
		fontDescriptionCopyStaticFunction = gi.FunctionInvokerNew("Pango", "copy_static")
	})
}

var copyStaticFontDescriptionInvoker *gi.Function

// CopyStatic is a representation of the C type pango_font_description_copy_static.
func (recv *FontDescription) CopyStatic() *FontDescription {
	if copyStaticFontDescriptionInvoker == nil {
		copyStaticFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "copy_static")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyStaticFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_equal' : parameter 'desc2' of type 'FontDescription' not supported

var fontDescriptionFreeFunction *gi.Function
var fontDescriptionFreeFunctionOnce sync.Once

func fontDescriptionFreeFunctionSet() {
	fontDescriptionFreeFunctionOnce.Do(func() {
		fontDescriptionFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeFontDescriptionInvoker *gi.Function

// Free is a representation of the C type pango_font_description_free.
func (recv *FontDescription) Free() {
	if freeFontDescriptionInvoker == nil {
		freeFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

var fontDescriptionGetFamilyFunction *gi.Function
var fontDescriptionGetFamilyFunctionOnce sync.Once

func fontDescriptionGetFamilyFunctionSet() {
	fontDescriptionGetFamilyFunctionOnce.Do(func() {
		fontDescriptionGetFamilyFunction = gi.FunctionInvokerNew("Pango", "get_family")
	})
}

var getFamilyFontDescriptionInvoker *gi.Function

// GetFamily is a representation of the C type pango_font_description_get_family.
func (recv *FontDescription) GetFamily() string {
	if getFamilyFontDescriptionInvoker == nil {
		getFamilyFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "get_family")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFamilyFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_gravity' : return type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_font_description_get_set_fields' : return type 'FontMask' not supported

var fontDescriptionGetSizeFunction *gi.Function
var fontDescriptionGetSizeFunctionOnce sync.Once

func fontDescriptionGetSizeFunctionSet() {
	fontDescriptionGetSizeFunctionOnce.Do(func() {
		fontDescriptionGetSizeFunction = gi.FunctionInvokerNew("Pango", "get_size")
	})
}

var getSizeFontDescriptionInvoker *gi.Function

// GetSize is a representation of the C type pango_font_description_get_size.
func (recv *FontDescription) GetSize() int32 {
	if getSizeFontDescriptionInvoker == nil {
		getSizeFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "get_size")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSizeFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_size_is_absolute' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_font_description_get_stretch' : return type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_get_style' : return type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_get_variant' : return type 'Variant' not supported

var fontDescriptionGetVariationsFunction *gi.Function
var fontDescriptionGetVariationsFunctionOnce sync.Once

func fontDescriptionGetVariationsFunctionSet() {
	fontDescriptionGetVariationsFunctionOnce.Do(func() {
		fontDescriptionGetVariationsFunction = gi.FunctionInvokerNew("Pango", "get_variations")
	})
}

var getVariationsFontDescriptionInvoker *gi.Function

// GetVariations is a representation of the C type pango_font_description_get_variations.
func (recv *FontDescription) GetVariations() string {
	if getVariationsFontDescriptionInvoker == nil {
		getVariationsFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "get_variations")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getVariationsFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_weight' : return type 'Weight' not supported

var fontDescriptionHashFunction *gi.Function
var fontDescriptionHashFunctionOnce sync.Once

func fontDescriptionHashFunctionSet() {
	fontDescriptionHashFunctionOnce.Do(func() {
		fontDescriptionHashFunction = gi.FunctionInvokerNew("Pango", "hash")
	})
}

var hashFontDescriptionInvoker *gi.Function

// Hash is a representation of the C type pango_font_description_hash.
func (recv *FontDescription) Hash() uint32 {
	if hashFontDescriptionInvoker == nil {
		hashFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_merge' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_merge_static' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_set_absolute_size' : parameter 'size' of type 'gdouble' not supported

var fontDescriptionSetFamilyFunction *gi.Function
var fontDescriptionSetFamilyFunctionOnce sync.Once

func fontDescriptionSetFamilyFunctionSet() {
	fontDescriptionSetFamilyFunctionOnce.Do(func() {
		fontDescriptionSetFamilyFunction = gi.FunctionInvokerNew("Pango", "set_family")
	})
}

var setFamilyFontDescriptionInvoker *gi.Function

// SetFamily is a representation of the C type pango_font_description_set_family.
func (recv *FontDescription) SetFamily(family string) {
	if setFamilyFontDescriptionInvoker == nil {
		setFamilyFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_family")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	setFamilyFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

var fontDescriptionSetFamilyStaticFunction *gi.Function
var fontDescriptionSetFamilyStaticFunctionOnce sync.Once

func fontDescriptionSetFamilyStaticFunctionSet() {
	fontDescriptionSetFamilyStaticFunctionOnce.Do(func() {
		fontDescriptionSetFamilyStaticFunction = gi.FunctionInvokerNew("Pango", "set_family_static")
	})
}

var setFamilyStaticFontDescriptionInvoker *gi.Function

// SetFamilyStatic is a representation of the C type pango_font_description_set_family_static.
func (recv *FontDescription) SetFamilyStatic(family string) {
	if setFamilyStaticFontDescriptionInvoker == nil {
		setFamilyStaticFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_family_static")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	setFamilyStaticFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_gravity' : parameter 'gravity' of type 'Gravity' not supported

var fontDescriptionSetSizeFunction *gi.Function
var fontDescriptionSetSizeFunctionOnce sync.Once

func fontDescriptionSetSizeFunctionSet() {
	fontDescriptionSetSizeFunctionOnce.Do(func() {
		fontDescriptionSetSizeFunction = gi.FunctionInvokerNew("Pango", "set_size")
	})
}

var setSizeFontDescriptionInvoker *gi.Function

// SetSize is a representation of the C type pango_font_description_set_size.
func (recv *FontDescription) SetSize(size int32) {
	if setSizeFontDescriptionInvoker == nil {
		setSizeFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_size")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(size)

	setSizeFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_stretch' : parameter 'stretch' of type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_set_style' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_set_variant' : parameter 'variant' of type 'Variant' not supported

var fontDescriptionSetVariationsFunction *gi.Function
var fontDescriptionSetVariationsFunctionOnce sync.Once

func fontDescriptionSetVariationsFunctionSet() {
	fontDescriptionSetVariationsFunctionOnce.Do(func() {
		fontDescriptionSetVariationsFunction = gi.FunctionInvokerNew("Pango", "set_variations")
	})
}

var setVariationsFontDescriptionInvoker *gi.Function

// SetVariations is a representation of the C type pango_font_description_set_variations.
func (recv *FontDescription) SetVariations(settings string) {
	if setVariationsFontDescriptionInvoker == nil {
		setVariationsFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_variations")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	setVariationsFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

var fontDescriptionSetVariationsStaticFunction *gi.Function
var fontDescriptionSetVariationsStaticFunctionOnce sync.Once

func fontDescriptionSetVariationsStaticFunctionSet() {
	fontDescriptionSetVariationsStaticFunctionOnce.Do(func() {
		fontDescriptionSetVariationsStaticFunction = gi.FunctionInvokerNew("Pango", "set_variations_static")
	})
}

var setVariationsStaticFontDescriptionInvoker *gi.Function

// SetVariationsStatic is a representation of the C type pango_font_description_set_variations_static.
func (recv *FontDescription) SetVariationsStatic(settings string) {
	if setVariationsStaticFontDescriptionInvoker == nil {
		setVariationsStaticFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "set_variations_static")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	setVariationsStaticFontDescriptionInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_weight' : parameter 'weight' of type 'Weight' not supported

var fontDescriptionToFilenameFunction *gi.Function
var fontDescriptionToFilenameFunctionOnce sync.Once

func fontDescriptionToFilenameFunctionSet() {
	fontDescriptionToFilenameFunctionOnce.Do(func() {
		fontDescriptionToFilenameFunction = gi.FunctionInvokerNew("Pango", "to_filename")
	})
}

var toFilenameFontDescriptionInvoker *gi.Function

// ToFilename is a representation of the C type pango_font_description_to_filename.
func (recv *FontDescription) ToFilename() string {
	if toFilenameFontDescriptionInvoker == nil {
		toFilenameFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "to_filename")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toFilenameFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var fontDescriptionToStringFunction *gi.Function
var fontDescriptionToStringFunctionOnce sync.Once

func fontDescriptionToStringFunctionSet() {
	fontDescriptionToStringFunctionOnce.Do(func() {
		fontDescriptionToStringFunction = gi.FunctionInvokerNew("Pango", "to_string")
	})
}

var toStringFontDescriptionInvoker *gi.Function

// ToString is a representation of the C type pango_font_description_to_string.
func (recv *FontDescription) ToString() string {
	if toStringFontDescriptionInvoker == nil {
		toStringFontDescriptionInvoker = gi.StructFunctionInvokerNew("Pango", "FontDescription", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringFontDescriptionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_unset_fields' : parameter 'to_unset' of type 'FontMask' not supported

var fontFaceClassStruct *gi.Struct
var fontFaceClassStructOnce sync.Once

func fontFaceClassStructSet() {
	fontFaceClassStructOnce.Do(func() {
		fontFaceClassStruct = gi.StructNew("Pango", "FontFaceClass")
	})
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
var fontFamilyClassStructOnce sync.Once

func fontFamilyClassStructSet() {
	fontFamilyClassStructOnce.Do(func() {
		fontFamilyClassStruct = gi.StructNew("Pango", "FontFamilyClass")
	})
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
var fontMapClassStructOnce sync.Once

func fontMapClassStructSet() {
	fontMapClassStructOnce.Do(func() {
		fontMapClassStruct = gi.StructNew("Pango", "FontMapClass")
	})
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
var fontMetricsStructOnce sync.Once

func fontMetricsStructSet() {
	fontMetricsStructOnce.Do(func() {
		fontMetricsStruct = gi.StructNew("Pango", "FontMetrics")
	})
}

type FontMetrics struct {
	native uintptr
}

var fontMetricsNewFunction *gi.Function
var fontMetricsNewFunctionOnce sync.Once

func fontMetricsNewFunctionSet() {
	fontMetricsNewFunctionOnce.Do(func() {
		fontMetricsNewFunction = gi.FunctionInvokerNew("Pango", "new")
	})
}

var newFontMetricsInvoker *gi.Function

// FontMetricsNew is a representation of the C type pango_font_metrics_new.
func FontMetricsNew() *FontMetrics {
	if newFontMetricsInvoker == nil {
		newFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "new")
	}

	ret := newFontMetricsInvoker.Invoke(nil, nil)

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var fontMetricsGetApproximateCharWidthFunction *gi.Function
var fontMetricsGetApproximateCharWidthFunctionOnce sync.Once

func fontMetricsGetApproximateCharWidthFunctionSet() {
	fontMetricsGetApproximateCharWidthFunctionOnce.Do(func() {
		fontMetricsGetApproximateCharWidthFunction = gi.FunctionInvokerNew("Pango", "get_approximate_char_width")
	})
}

var getApproximateCharWidthFontMetricsInvoker *gi.Function

// GetApproximateCharWidth is a representation of the C type pango_font_metrics_get_approximate_char_width.
func (recv *FontMetrics) GetApproximateCharWidth() int32 {
	if getApproximateCharWidthFontMetricsInvoker == nil {
		getApproximateCharWidthFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_approximate_char_width")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getApproximateCharWidthFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetApproximateDigitWidthFunction *gi.Function
var fontMetricsGetApproximateDigitWidthFunctionOnce sync.Once

func fontMetricsGetApproximateDigitWidthFunctionSet() {
	fontMetricsGetApproximateDigitWidthFunctionOnce.Do(func() {
		fontMetricsGetApproximateDigitWidthFunction = gi.FunctionInvokerNew("Pango", "get_approximate_digit_width")
	})
}

var getApproximateDigitWidthFontMetricsInvoker *gi.Function

// GetApproximateDigitWidth is a representation of the C type pango_font_metrics_get_approximate_digit_width.
func (recv *FontMetrics) GetApproximateDigitWidth() int32 {
	if getApproximateDigitWidthFontMetricsInvoker == nil {
		getApproximateDigitWidthFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_approximate_digit_width")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getApproximateDigitWidthFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetAscentFunction *gi.Function
var fontMetricsGetAscentFunctionOnce sync.Once

func fontMetricsGetAscentFunctionSet() {
	fontMetricsGetAscentFunctionOnce.Do(func() {
		fontMetricsGetAscentFunction = gi.FunctionInvokerNew("Pango", "get_ascent")
	})
}

var getAscentFontMetricsInvoker *gi.Function

// GetAscent is a representation of the C type pango_font_metrics_get_ascent.
func (recv *FontMetrics) GetAscent() int32 {
	if getAscentFontMetricsInvoker == nil {
		getAscentFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_ascent")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getAscentFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetDescentFunction *gi.Function
var fontMetricsGetDescentFunctionOnce sync.Once

func fontMetricsGetDescentFunctionSet() {
	fontMetricsGetDescentFunctionOnce.Do(func() {
		fontMetricsGetDescentFunction = gi.FunctionInvokerNew("Pango", "get_descent")
	})
}

var getDescentFontMetricsInvoker *gi.Function

// GetDescent is a representation of the C type pango_font_metrics_get_descent.
func (recv *FontMetrics) GetDescent() int32 {
	if getDescentFontMetricsInvoker == nil {
		getDescentFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_descent")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescentFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughPositionFunction *gi.Function
var fontMetricsGetStrikethroughPositionFunctionOnce sync.Once

func fontMetricsGetStrikethroughPositionFunctionSet() {
	fontMetricsGetStrikethroughPositionFunctionOnce.Do(func() {
		fontMetricsGetStrikethroughPositionFunction = gi.FunctionInvokerNew("Pango", "get_strikethrough_position")
	})
}

var getStrikethroughPositionFontMetricsInvoker *gi.Function

// GetStrikethroughPosition is a representation of the C type pango_font_metrics_get_strikethrough_position.
func (recv *FontMetrics) GetStrikethroughPosition() int32 {
	if getStrikethroughPositionFontMetricsInvoker == nil {
		getStrikethroughPositionFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_strikethrough_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStrikethroughPositionFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughThicknessFunction *gi.Function
var fontMetricsGetStrikethroughThicknessFunctionOnce sync.Once

func fontMetricsGetStrikethroughThicknessFunctionSet() {
	fontMetricsGetStrikethroughThicknessFunctionOnce.Do(func() {
		fontMetricsGetStrikethroughThicknessFunction = gi.FunctionInvokerNew("Pango", "get_strikethrough_thickness")
	})
}

var getStrikethroughThicknessFontMetricsInvoker *gi.Function

// GetStrikethroughThickness is a representation of the C type pango_font_metrics_get_strikethrough_thickness.
func (recv *FontMetrics) GetStrikethroughThickness() int32 {
	if getStrikethroughThicknessFontMetricsInvoker == nil {
		getStrikethroughThicknessFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_strikethrough_thickness")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStrikethroughThicknessFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlinePositionFunction *gi.Function
var fontMetricsGetUnderlinePositionFunctionOnce sync.Once

func fontMetricsGetUnderlinePositionFunctionSet() {
	fontMetricsGetUnderlinePositionFunctionOnce.Do(func() {
		fontMetricsGetUnderlinePositionFunction = gi.FunctionInvokerNew("Pango", "get_underline_position")
	})
}

var getUnderlinePositionFontMetricsInvoker *gi.Function

// GetUnderlinePosition is a representation of the C type pango_font_metrics_get_underline_position.
func (recv *FontMetrics) GetUnderlinePosition() int32 {
	if getUnderlinePositionFontMetricsInvoker == nil {
		getUnderlinePositionFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_underline_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUnderlinePositionFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlineThicknessFunction *gi.Function
var fontMetricsGetUnderlineThicknessFunctionOnce sync.Once

func fontMetricsGetUnderlineThicknessFunctionSet() {
	fontMetricsGetUnderlineThicknessFunctionOnce.Do(func() {
		fontMetricsGetUnderlineThicknessFunction = gi.FunctionInvokerNew("Pango", "get_underline_thickness")
	})
}

var getUnderlineThicknessFontMetricsInvoker *gi.Function

// GetUnderlineThickness is a representation of the C type pango_font_metrics_get_underline_thickness.
func (recv *FontMetrics) GetUnderlineThickness() int32 {
	if getUnderlineThicknessFontMetricsInvoker == nil {
		getUnderlineThicknessFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "get_underline_thickness")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUnderlineThicknessFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsRefFunction *gi.Function
var fontMetricsRefFunctionOnce sync.Once

func fontMetricsRefFunctionSet() {
	fontMetricsRefFunctionOnce.Do(func() {
		fontMetricsRefFunction = gi.FunctionInvokerNew("Pango", "ref")
	})
}

var refFontMetricsInvoker *gi.Function

// Ref is a representation of the C type pango_font_metrics_ref.
func (recv *FontMetrics) Ref() *FontMetrics {
	if refFontMetricsInvoker == nil {
		refFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refFontMetricsInvoker.Invoke(inArgs[:], nil)

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var fontMetricsUnrefFunction *gi.Function
var fontMetricsUnrefFunctionOnce sync.Once

func fontMetricsUnrefFunctionSet() {
	fontMetricsUnrefFunctionOnce.Do(func() {
		fontMetricsUnrefFunction = gi.FunctionInvokerNew("Pango", "unref")
	})
}

var unrefFontMetricsInvoker *gi.Function

// Unref is a representation of the C type pango_font_metrics_unref.
func (recv *FontMetrics) Unref() {
	if unrefFontMetricsInvoker == nil {
		unrefFontMetricsInvoker = gi.StructFunctionInvokerNew("Pango", "FontMetrics", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefFontMetricsInvoker.Invoke(inArgs[:], nil)

}

var fontsetClassStruct *gi.Struct
var fontsetClassStructOnce sync.Once

func fontsetClassStructSet() {
	fontsetClassStructOnce.Do(func() {
		fontsetClassStruct = gi.StructNew("Pango", "FontsetClass")
	})
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
var fontsetSimpleClassStructOnce sync.Once

func fontsetSimpleClassStructSet() {
	fontsetSimpleClassStructOnce.Do(func() {
		fontsetSimpleClassStruct = gi.StructNew("Pango", "FontsetSimpleClass")
	})
}

type FontsetSimpleClass struct {
	native uintptr
}

var glyphGeometryStruct *gi.Struct
var glyphGeometryStructOnce sync.Once

func glyphGeometryStructSet() {
	glyphGeometryStructOnce.Do(func() {
		glyphGeometryStruct = gi.StructNew("Pango", "GlyphGeometry")
	})
}

type GlyphGeometry struct {
	native  uintptr
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

var glyphInfoStruct *gi.Struct
var glyphInfoStructOnce sync.Once

func glyphInfoStructSet() {
	glyphInfoStructOnce.Do(func() {
		glyphInfoStruct = gi.StructNew("Pango", "GlyphInfo")
	})
}

type GlyphInfo struct {
	native   uintptr
	Glyph    Glyph
	Geometry *GlyphGeometry
	Attr     *GlyphVisAttr
}

var glyphItemStruct *gi.Struct
var glyphItemStructOnce sync.Once

func glyphItemStructSet() {
	glyphItemStructOnce.Do(func() {
		glyphItemStruct = gi.StructNew("Pango", "GlyphItem")
	})
}

type GlyphItem struct {
	native uintptr
	Item   *Item
	Glyphs *GlyphString
}

// UNSUPPORTED : C value 'pango_glyph_item_apply_attrs' : parameter 'list' of type 'AttrList' not supported

var glyphItemCopyFunction *gi.Function
var glyphItemCopyFunctionOnce sync.Once

func glyphItemCopyFunctionSet() {
	glyphItemCopyFunctionOnce.Do(func() {
		glyphItemCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyGlyphItemInvoker *gi.Function

// Copy is a representation of the C type pango_glyph_item_copy.
func (recv *GlyphItem) Copy() *GlyphItem {
	if copyGlyphItemInvoker == nil {
		copyGlyphItemInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItem", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyGlyphItemInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

var glyphItemFreeFunction *gi.Function
var glyphItemFreeFunctionOnce sync.Once

func glyphItemFreeFunctionSet() {
	glyphItemFreeFunctionOnce.Do(func() {
		glyphItemFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeGlyphItemInvoker *gi.Function

// Free is a representation of the C type pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	if freeGlyphItemInvoker == nil {
		freeGlyphItemInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItem", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeGlyphItemInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_item_get_logical_widths' : parameter 'logical_widths' has no type

// UNSUPPORTED : C value 'pango_glyph_item_letter_space' : parameter 'log_attrs' has no type

var glyphItemSplitFunction *gi.Function
var glyphItemSplitFunctionOnce sync.Once

func glyphItemSplitFunctionSet() {
	glyphItemSplitFunctionOnce.Do(func() {
		glyphItemSplitFunction = gi.FunctionInvokerNew("Pango", "split")
	})
}

var splitGlyphItemInvoker *gi.Function

// Split is a representation of the C type pango_glyph_item_split.
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	if splitGlyphItemInvoker == nil {
		splitGlyphItemInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItem", "split")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(splitIndex)

	ret := splitGlyphItemInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

var glyphItemIterStruct *gi.Struct
var glyphItemIterStructOnce sync.Once

func glyphItemIterStructSet() {
	glyphItemIterStructOnce.Do(func() {
		glyphItemIterStruct = gi.StructNew("Pango", "GlyphItemIter")
	})
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
var glyphItemIterCopyFunctionOnce sync.Once

func glyphItemIterCopyFunctionSet() {
	glyphItemIterCopyFunctionOnce.Do(func() {
		glyphItemIterCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyGlyphItemIterInvoker *gi.Function

// Copy is a representation of the C type pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	if copyGlyphItemIterInvoker == nil {
		copyGlyphItemIterInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItemIter", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyGlyphItemIterInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphItemIter{native: ret.Pointer()}

	return retGo
}

var glyphItemIterFreeFunction *gi.Function
var glyphItemIterFreeFunctionOnce sync.Once

func glyphItemIterFreeFunctionSet() {
	glyphItemIterFreeFunctionOnce.Do(func() {
		glyphItemIterFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeGlyphItemIterInvoker *gi.Function

// Free is a representation of the C type pango_glyph_item_iter_free.
func (recv *GlyphItemIter) Free() {
	if freeGlyphItemIterInvoker == nil {
		freeGlyphItemIterInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphItemIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeGlyphItemIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_end' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_start' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_next_cluster' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_prev_cluster' : return type 'gboolean' not supported

var glyphStringStruct *gi.Struct
var glyphStringStructOnce sync.Once

func glyphStringStructSet() {
	glyphStringStructOnce.Do(func() {
		glyphStringStruct = gi.StructNew("Pango", "GlyphString")
	})
}

type GlyphString struct {
	native    uintptr
	NumGlyphs int32
	// UNSUPPORTED : C value 'glyphs' : missing Type
	LogClusters int32
}

var glyphStringNewFunction *gi.Function
var glyphStringNewFunctionOnce sync.Once

func glyphStringNewFunctionSet() {
	glyphStringNewFunctionOnce.Do(func() {
		glyphStringNewFunction = gi.FunctionInvokerNew("Pango", "new")
	})
}

var newGlyphStringInvoker *gi.Function

// GlyphStringNew is a representation of the C type pango_glyph_string_new.
func GlyphStringNew() *GlyphString {
	if newGlyphStringInvoker == nil {
		newGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "new")
	}

	ret := newGlyphStringInvoker.Invoke(nil, nil)

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

var glyphStringCopyFunction *gi.Function
var glyphStringCopyFunctionOnce sync.Once

func glyphStringCopyFunctionSet() {
	glyphStringCopyFunctionOnce.Do(func() {
		glyphStringCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyGlyphStringInvoker *gi.Function

// Copy is a representation of the C type pango_glyph_string_copy.
func (recv *GlyphString) Copy() *GlyphString {
	if copyGlyphStringInvoker == nil {
		copyGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyGlyphStringInvoker.Invoke(inArgs[:], nil)

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_extents' : parameter 'font' of type 'Font' not supported

// UNSUPPORTED : C value 'pango_glyph_string_extents_range' : parameter 'font' of type 'Font' not supported

var glyphStringFreeFunction *gi.Function
var glyphStringFreeFunctionOnce sync.Once

func glyphStringFreeFunctionSet() {
	glyphStringFreeFunctionOnce.Do(func() {
		glyphStringFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeGlyphStringInvoker *gi.Function

// Free is a representation of the C type pango_glyph_string_free.
func (recv *GlyphString) Free() {
	if freeGlyphStringInvoker == nil {
		freeGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeGlyphStringInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_string_get_logical_widths' : parameter 'logical_widths' has no type

var glyphStringGetWidthFunction *gi.Function
var glyphStringGetWidthFunctionOnce sync.Once

func glyphStringGetWidthFunctionSet() {
	glyphStringGetWidthFunctionOnce.Do(func() {
		glyphStringGetWidthFunction = gi.FunctionInvokerNew("Pango", "get_width")
	})
}

var getWidthGlyphStringInvoker *gi.Function

// GetWidth is a representation of the C type pango_glyph_string_get_width.
func (recv *GlyphString) GetWidth() int32 {
	if getWidthGlyphStringInvoker == nil {
		getWidthGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "get_width")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getWidthGlyphStringInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_index_to_x' : parameter 'analysis' of type 'Analysis' not supported

var glyphStringSetSizeFunction *gi.Function
var glyphStringSetSizeFunctionOnce sync.Once

func glyphStringSetSizeFunctionSet() {
	glyphStringSetSizeFunctionOnce.Do(func() {
		glyphStringSetSizeFunction = gi.FunctionInvokerNew("Pango", "set_size")
	})
}

var setSizeGlyphStringInvoker *gi.Function

// SetSize is a representation of the C type pango_glyph_string_set_size.
func (recv *GlyphString) SetSize(newLen int32) {
	if setSizeGlyphStringInvoker == nil {
		setSizeGlyphStringInvoker = gi.StructFunctionInvokerNew("Pango", "GlyphString", "set_size")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newLen)

	setSizeGlyphStringInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_string_x_to_index' : parameter 'analysis' of type 'Analysis' not supported

var glyphVisAttrStruct *gi.Struct
var glyphVisAttrStructOnce sync.Once

func glyphVisAttrStructSet() {
	glyphVisAttrStructOnce.Do(func() {
		glyphVisAttrStruct = gi.StructNew("Pango", "GlyphVisAttr")
	})
}

type GlyphVisAttr struct {
	native         uintptr
	IsClusterStart uint32
}

var includedModuleStruct *gi.Struct
var includedModuleStructOnce sync.Once

func includedModuleStructSet() {
	includedModuleStructOnce.Do(func() {
		includedModuleStruct = gi.StructNew("Pango", "IncludedModule")
	})
}

type IncludedModule struct {
	native uintptr
	// UNSUPPORTED : C value 'list' : missing Type
	// UNSUPPORTED : C value 'init' : missing Type
	// UNSUPPORTED : C value 'exit' : missing Type
	// UNSUPPORTED : C value 'create' : missing Type
}

var itemStruct *gi.Struct
var itemStructOnce sync.Once

func itemStructSet() {
	itemStructOnce.Do(func() {
		itemStruct = gi.StructNew("Pango", "Item")
	})
}

type Item struct {
	native   uintptr
	Offset   int32
	Length   int32
	NumChars int32
	Analysis *Analysis
}

var itemNewFunction *gi.Function
var itemNewFunctionOnce sync.Once

func itemNewFunctionSet() {
	itemNewFunctionOnce.Do(func() {
		itemNewFunction = gi.FunctionInvokerNew("Pango", "new")
	})
}

var newItemInvoker *gi.Function

// ItemNew is a representation of the C type pango_item_new.
func ItemNew() *Item {
	if newItemInvoker == nil {
		newItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "new")
	}

	ret := newItemInvoker.Invoke(nil, nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var itemCopyFunction *gi.Function
var itemCopyFunctionOnce sync.Once

func itemCopyFunctionSet() {
	itemCopyFunctionOnce.Do(func() {
		itemCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyItemInvoker *gi.Function

// Copy is a representation of the C type pango_item_copy.
func (recv *Item) Copy() *Item {
	if copyItemInvoker == nil {
		copyItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyItemInvoker.Invoke(inArgs[:], nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var itemFreeFunction *gi.Function
var itemFreeFunctionOnce sync.Once

func itemFreeFunctionSet() {
	itemFreeFunctionOnce.Do(func() {
		itemFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeItemInvoker *gi.Function

// Free is a representation of the C type pango_item_free.
func (recv *Item) Free() {
	if freeItemInvoker == nil {
		freeItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeItemInvoker.Invoke(inArgs[:], nil)

}

var itemSplitFunction *gi.Function
var itemSplitFunctionOnce sync.Once

func itemSplitFunctionSet() {
	itemSplitFunctionOnce.Do(func() {
		itemSplitFunction = gi.FunctionInvokerNew("Pango", "split")
	})
}

var splitItemInvoker *gi.Function

// Split is a representation of the C type pango_item_split.
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	if splitItemInvoker == nil {
		splitItemInvoker = gi.StructFunctionInvokerNew("Pango", "Item", "split")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(splitIndex)
	inArgs[2].SetInt32(splitOffset)

	ret := splitItemInvoker.Invoke(inArgs[:], nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var languageStruct *gi.Struct
var languageStructOnce sync.Once

func languageStructSet() {
	languageStructOnce.Do(func() {
		languageStruct = gi.StructNew("Pango", "Language")
	})
}

type Language struct {
	native uintptr
}

var languageGetSampleStringFunction *gi.Function
var languageGetSampleStringFunctionOnce sync.Once

func languageGetSampleStringFunctionSet() {
	languageGetSampleStringFunctionOnce.Do(func() {
		languageGetSampleStringFunction = gi.FunctionInvokerNew("Pango", "get_sample_string")
	})
}

var getSampleStringLanguageInvoker *gi.Function

// GetSampleString is a representation of the C type pango_language_get_sample_string.
func (recv *Language) GetSampleString() string {
	if getSampleStringLanguageInvoker == nil {
		getSampleStringLanguageInvoker = gi.StructFunctionInvokerNew("Pango", "Language", "get_sample_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSampleStringLanguageInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var languageGetScriptsFunction *gi.Function
var languageGetScriptsFunctionOnce sync.Once

func languageGetScriptsFunctionSet() {
	languageGetScriptsFunctionOnce.Do(func() {
		languageGetScriptsFunction = gi.FunctionInvokerNew("Pango", "get_scripts")
	})
}

var getScriptsLanguageInvoker *gi.Function

// GetScripts is a representation of the C type pango_language_get_scripts.
func (recv *Language) GetScripts() int32 {
	if getScriptsLanguageInvoker == nil {
		getScriptsLanguageInvoker = gi.StructFunctionInvokerNew("Pango", "Language", "get_scripts")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	getScriptsLanguageInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'pango_language_includes_script' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_language_matches' : return type 'gboolean' not supported

var languageToStringFunction *gi.Function
var languageToStringFunctionOnce sync.Once

func languageToStringFunctionSet() {
	languageToStringFunctionOnce.Do(func() {
		languageToStringFunction = gi.FunctionInvokerNew("Pango", "to_string")
	})
}

var toStringLanguageInvoker *gi.Function

// ToString is a representation of the C type pango_language_to_string.
func (recv *Language) ToString() string {
	if toStringLanguageInvoker == nil {
		toStringLanguageInvoker = gi.StructFunctionInvokerNew("Pango", "Language", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringLanguageInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var layoutClassStruct *gi.Struct
var layoutClassStructOnce sync.Once

func layoutClassStructSet() {
	layoutClassStructOnce.Do(func() {
		layoutClassStruct = gi.StructNew("Pango", "LayoutClass")
	})
}

type LayoutClass struct {
	native uintptr
}

var layoutIterStruct *gi.Struct
var layoutIterStructOnce sync.Once

func layoutIterStructSet() {
	layoutIterStructOnce.Do(func() {
		layoutIterStruct = gi.StructNew("Pango", "LayoutIter")
	})
}

type LayoutIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_layout_iter_at_last_line' : return type 'gboolean' not supported

var layoutIterCopyFunction *gi.Function
var layoutIterCopyFunctionOnce sync.Once

func layoutIterCopyFunctionSet() {
	layoutIterCopyFunctionOnce.Do(func() {
		layoutIterCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyLayoutIterInvoker *gi.Function

// Copy is a representation of the C type pango_layout_iter_copy.
func (recv *LayoutIter) Copy() *LayoutIter {
	if copyLayoutIterInvoker == nil {
		copyLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutIter{native: ret.Pointer()}

	return retGo
}

var layoutIterFreeFunction *gi.Function
var layoutIterFreeFunctionOnce sync.Once

func layoutIterFreeFunctionSet() {
	layoutIterFreeFunctionOnce.Do(func() {
		layoutIterFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeLayoutIterInvoker *gi.Function

// Free is a representation of the C type pango_layout_iter_free.
func (recv *LayoutIter) Free() {
	if freeLayoutIterInvoker == nil {
		freeLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeLayoutIterInvoker.Invoke(inArgs[:], nil)

}

var layoutIterGetBaselineFunction *gi.Function
var layoutIterGetBaselineFunctionOnce sync.Once

func layoutIterGetBaselineFunctionSet() {
	layoutIterGetBaselineFunctionOnce.Do(func() {
		layoutIterGetBaselineFunction = gi.FunctionInvokerNew("Pango", "get_baseline")
	})
}

var getBaselineLayoutIterInvoker *gi.Function

// GetBaseline is a representation of the C type pango_layout_iter_get_baseline.
func (recv *LayoutIter) GetBaseline() int32 {
	if getBaselineLayoutIterInvoker == nil {
		getBaselineLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_baseline")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getBaselineLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_char_extents' : parameter 'logical_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_cluster_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetIndexFunction *gi.Function
var layoutIterGetIndexFunctionOnce sync.Once

func layoutIterGetIndexFunctionSet() {
	layoutIterGetIndexFunctionOnce.Do(func() {
		layoutIterGetIndexFunction = gi.FunctionInvokerNew("Pango", "get_index")
	})
}

var getIndexLayoutIterInvoker *gi.Function

// GetIndex is a representation of the C type pango_layout_iter_get_index.
func (recv *LayoutIter) GetIndex() int32 {
	if getIndexLayoutIterInvoker == nil {
		getIndexLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_index")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIndexLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_layout' : return type 'Layout' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_layout_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetLineFunction *gi.Function
var layoutIterGetLineFunctionOnce sync.Once

func layoutIterGetLineFunctionSet() {
	layoutIterGetLineFunctionOnce.Do(func() {
		layoutIterGetLineFunction = gi.FunctionInvokerNew("Pango", "get_line")
	})
}

var getLineLayoutIterInvoker *gi.Function

// GetLine is a representation of the C type pango_layout_iter_get_line.
func (recv *LayoutIter) GetLine() *LayoutLine {
	if getLineLayoutIterInvoker == nil {
		getLineLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLineLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_line_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetLineReadonlyFunction *gi.Function
var layoutIterGetLineReadonlyFunctionOnce sync.Once

func layoutIterGetLineReadonlyFunctionSet() {
	layoutIterGetLineReadonlyFunctionOnce.Do(func() {
		layoutIterGetLineReadonlyFunction = gi.FunctionInvokerNew("Pango", "get_line_readonly")
	})
}

var getLineReadonlyLayoutIterInvoker *gi.Function

// GetLineReadonly is a representation of the C type pango_layout_iter_get_line_readonly.
func (recv *LayoutIter) GetLineReadonly() *LayoutLine {
	if getLineReadonlyLayoutIterInvoker == nil {
		getLineReadonlyLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_line_readonly")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLineReadonlyLayoutIterInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var layoutIterGetLineYrangeFunction *gi.Function
var layoutIterGetLineYrangeFunctionOnce sync.Once

func layoutIterGetLineYrangeFunctionSet() {
	layoutIterGetLineYrangeFunctionOnce.Do(func() {
		layoutIterGetLineYrangeFunction = gi.FunctionInvokerNew("Pango", "get_line_yrange")
	})
}

var getLineYrangeLayoutIterInvoker *gi.Function

// GetLineYrange is a representation of the C type pango_layout_iter_get_line_yrange.
func (recv *LayoutIter) GetLineYrange() (int32, int32) {
	if getLineYrangeLayoutIterInvoker == nil {
		getLineYrangeLayoutIterInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutIter", "get_line_yrange")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	getLineYrangeLayoutIterInvoker.Invoke(inArgs[:], outArgs[:])

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
var layoutLineStructOnce sync.Once

func layoutLineStructSet() {
	layoutLineStructOnce.Do(func() {
		layoutLineStruct = gi.StructNew("Pango", "LayoutLine")
	})
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
var layoutLineRefFunctionOnce sync.Once

func layoutLineRefFunctionSet() {
	layoutLineRefFunctionOnce.Do(func() {
		layoutLineRefFunction = gi.FunctionInvokerNew("Pango", "ref")
	})
}

var refLayoutLineInvoker *gi.Function

// Ref is a representation of the C type pango_layout_line_ref.
func (recv *LayoutLine) Ref() *LayoutLine {
	if refLayoutLineInvoker == nil {
		refLayoutLineInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutLine", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refLayoutLineInvoker.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var layoutLineUnrefFunction *gi.Function
var layoutLineUnrefFunctionOnce sync.Once

func layoutLineUnrefFunctionSet() {
	layoutLineUnrefFunctionOnce.Do(func() {
		layoutLineUnrefFunction = gi.FunctionInvokerNew("Pango", "unref")
	})
}

var unrefLayoutLineInvoker *gi.Function

// Unref is a representation of the C type pango_layout_line_unref.
func (recv *LayoutLine) Unref() {
	if unrefLayoutLineInvoker == nil {
		unrefLayoutLineInvoker = gi.StructFunctionInvokerNew("Pango", "LayoutLine", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefLayoutLineInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_layout_line_x_to_index' : return type 'gboolean' not supported

var logAttrStruct *gi.Struct
var logAttrStructOnce sync.Once

func logAttrStructSet() {
	logAttrStructOnce.Do(func() {
		logAttrStruct = gi.StructNew("Pango", "LogAttr")
	})
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
var mapStructOnce sync.Once

func mapStructSet() {
	mapStructOnce.Do(func() {
		mapStruct = gi.StructNew("Pango", "Map")
	})
}

type Map struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_map_get_engine' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_map_get_engines' : parameter 'script' of type 'Script' not supported

var mapEntryStruct *gi.Struct
var mapEntryStructOnce sync.Once

func mapEntryStructSet() {
	mapEntryStructOnce.Do(func() {
		mapEntryStruct = gi.StructNew("Pango", "MapEntry")
	})
}

type MapEntry struct {
	native uintptr
}

var matrixStruct *gi.Struct
var matrixStructOnce sync.Once

func matrixStructSet() {
	matrixStructOnce.Do(func() {
		matrixStruct = gi.StructNew("Pango", "Matrix")
	})
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
var matrixCopyFunctionOnce sync.Once

func matrixCopyFunctionSet() {
	matrixCopyFunctionOnce.Do(func() {
		matrixCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyMatrixInvoker *gi.Function

// Copy is a representation of the C type pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	if copyMatrixInvoker == nil {
		copyMatrixInvoker = gi.StructFunctionInvokerNew("Pango", "Matrix", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyMatrixInvoker.Invoke(inArgs[:], nil)

	retGo := &Matrix{native: ret.Pointer()}

	return retGo
}

var matrixFreeFunction *gi.Function
var matrixFreeFunctionOnce sync.Once

func matrixFreeFunctionSet() {
	matrixFreeFunctionOnce.Do(func() {
		matrixFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeMatrixInvoker *gi.Function

// Free is a representation of the C type pango_matrix_free.
func (recv *Matrix) Free() {
	if freeMatrixInvoker == nil {
		freeMatrixInvoker = gi.StructFunctionInvokerNew("Pango", "Matrix", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMatrixInvoker.Invoke(inArgs[:], nil)

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
var rectangleStructOnce sync.Once

func rectangleStructSet() {
	rectangleStructOnce.Do(func() {
		rectangleStruct = gi.StructNew("Pango", "Rectangle")
	})
}

type Rectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

var rendererClassStruct *gi.Struct
var rendererClassStructOnce sync.Once

func rendererClassStructSet() {
	rendererClassStructOnce.Do(func() {
		rendererClassStruct = gi.StructNew("Pango", "RendererClass")
	})
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
var rendererPrivateStructOnce sync.Once

func rendererPrivateStructSet() {
	rendererPrivateStructOnce.Do(func() {
		rendererPrivateStruct = gi.StructNew("Pango", "RendererPrivate")
	})
}

type RendererPrivate struct {
	native uintptr
}

var scriptIterStruct *gi.Struct
var scriptIterStructOnce sync.Once

func scriptIterStructSet() {
	scriptIterStructOnce.Do(func() {
		scriptIterStruct = gi.StructNew("Pango", "ScriptIter")
	})
}

type ScriptIter struct {
	native uintptr
}

var scriptIterFreeFunction *gi.Function
var scriptIterFreeFunctionOnce sync.Once

func scriptIterFreeFunctionSet() {
	scriptIterFreeFunctionOnce.Do(func() {
		scriptIterFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeScriptIterInvoker *gi.Function

// Free is a representation of the C type pango_script_iter_free.
func (recv *ScriptIter) Free() {
	if freeScriptIterInvoker == nil {
		freeScriptIterInvoker = gi.StructFunctionInvokerNew("Pango", "ScriptIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeScriptIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_script_iter_get_range' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_script_iter_next' : return type 'gboolean' not supported

var tabArrayStruct *gi.Struct
var tabArrayStructOnce sync.Once

func tabArrayStructSet() {
	tabArrayStructOnce.Do(func() {
		tabArrayStruct = gi.StructNew("Pango", "TabArray")
	})
}

type TabArray struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_tab_array_new' : parameter 'positions_in_pixels' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_tab_array_new_with_positions' : parameter 'positions_in_pixels' of type 'gboolean' not supported

var tabArrayCopyFunction *gi.Function
var tabArrayCopyFunctionOnce sync.Once

func tabArrayCopyFunctionSet() {
	tabArrayCopyFunctionOnce.Do(func() {
		tabArrayCopyFunction = gi.FunctionInvokerNew("Pango", "copy")
	})
}

var copyTabArrayInvoker *gi.Function

// Copy is a representation of the C type pango_tab_array_copy.
func (recv *TabArray) Copy() *TabArray {
	if copyTabArrayInvoker == nil {
		copyTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTabArrayInvoker.Invoke(inArgs[:], nil)

	retGo := &TabArray{native: ret.Pointer()}

	return retGo
}

var tabArrayFreeFunction *gi.Function
var tabArrayFreeFunctionOnce sync.Once

func tabArrayFreeFunctionSet() {
	tabArrayFreeFunctionOnce.Do(func() {
		tabArrayFreeFunction = gi.FunctionInvokerNew("Pango", "free")
	})
}

var freeTabArrayInvoker *gi.Function

// Free is a representation of the C type pango_tab_array_free.
func (recv *TabArray) Free() {
	if freeTabArrayInvoker == nil {
		freeTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTabArrayInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_tab_array_get_positions_in_pixels' : return type 'gboolean' not supported

var tabArrayGetSizeFunction *gi.Function
var tabArrayGetSizeFunctionOnce sync.Once

func tabArrayGetSizeFunctionSet() {
	tabArrayGetSizeFunctionOnce.Do(func() {
		tabArrayGetSizeFunction = gi.FunctionInvokerNew("Pango", "get_size")
	})
}

var getSizeTabArrayInvoker *gi.Function

// GetSize is a representation of the C type pango_tab_array_get_size.
func (recv *TabArray) GetSize() int32 {
	if getSizeTabArrayInvoker == nil {
		getSizeTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "get_size")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSizeTabArrayInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_tab_array_get_tab' : parameter 'alignment' of type 'TabAlign' not supported

// UNSUPPORTED : C value 'pango_tab_array_get_tabs' : parameter 'alignments' of type 'TabAlign' not supported

var tabArrayResizeFunction *gi.Function
var tabArrayResizeFunctionOnce sync.Once

func tabArrayResizeFunctionSet() {
	tabArrayResizeFunctionOnce.Do(func() {
		tabArrayResizeFunction = gi.FunctionInvokerNew("Pango", "resize")
	})
}

var resizeTabArrayInvoker *gi.Function

// Resize is a representation of the C type pango_tab_array_resize.
func (recv *TabArray) Resize(newSize int32) {
	if resizeTabArrayInvoker == nil {
		resizeTabArrayInvoker = gi.StructFunctionInvokerNew("Pango", "TabArray", "resize")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newSize)

	resizeTabArrayInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_tab_array_set_tab' : parameter 'alignment' of type 'TabAlign' not supported

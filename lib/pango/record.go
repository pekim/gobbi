// Code generated - DO NOT EDIT.

package pango

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var analysisStruct *gi.Struct
var analysisStruct_Once sync.Once

func analysisStruct_Set() {
	analysisStruct_Once.Do(func() {
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
var attrClassStruct_Once sync.Once

func attrClassStruct_Set() {
	attrClassStruct_Once.Do(func() {
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
var attrColorStruct_Once sync.Once

func attrColorStruct_Set() {
	attrColorStruct_Once.Do(func() {
		attrColorStruct = gi.StructNew("Pango", "AttrColor")
	})
}

type AttrColor struct {
	native uintptr
	Attr   *Attribute
	Color  *Color
}

var attrFloatStruct *gi.Struct
var attrFloatStruct_Once sync.Once

func attrFloatStruct_Set() {
	attrFloatStruct_Once.Do(func() {
		attrFloatStruct = gi.StructNew("Pango", "AttrFloat")
	})
}

type AttrFloat struct {
	native uintptr
	Attr   *Attribute
	// UNSUPPORTED : C value 'value' : no Go type for 'gdouble'
}

var attrFontDescStruct *gi.Struct
var attrFontDescStruct_Once sync.Once

func attrFontDescStruct_Set() {
	attrFontDescStruct_Once.Do(func() {
		attrFontDescStruct = gi.StructNew("Pango", "AttrFontDesc")
	})
}

type AttrFontDesc struct {
	native uintptr
	Attr   *Attribute
	Desc   *FontDescription
}

var attrFontFeaturesStruct *gi.Struct
var attrFontFeaturesStruct_Once sync.Once

func attrFontFeaturesStruct_Set() {
	attrFontFeaturesStruct_Once.Do(func() {
		attrFontFeaturesStruct = gi.StructNew("Pango", "AttrFontFeatures")
	})
}

type AttrFontFeatures struct {
	native   uintptr
	Attr     *Attribute
	Features string
}

var attrIntStruct *gi.Struct
var attrIntStruct_Once sync.Once

func attrIntStruct_Set() {
	attrIntStruct_Once.Do(func() {
		attrIntStruct = gi.StructNew("Pango", "AttrInt")
	})
}

type AttrInt struct {
	native uintptr
	Attr   *Attribute
	Value  int32
}

var attrIteratorStruct *gi.Struct
var attrIteratorStruct_Once sync.Once

func attrIteratorStruct_Set() {
	attrIteratorStruct_Once.Do(func() {
		attrIteratorStruct = gi.StructNew("Pango", "AttrIterator")
	})
}

type AttrIterator struct {
	native uintptr
}

var attrIteratorCopyFunction *gi.Function
var attrIteratorCopyFunction_Once sync.Once

func attrIteratorCopyFunction_Set() {
	attrIteratorCopyFunction_Once.Do(func() {
		attrIteratorStruct_Set()
		attrIteratorCopyFunction = attrIteratorStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_attr_iterator_copy.
func (recv *AttrIterator) Copy() *AttrIterator {
	attrIteratorCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := attrIteratorCopyFunction.Invoke(inArgs[:], nil)

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

var attrIteratorDestroyFunction *gi.Function
var attrIteratorDestroyFunction_Once sync.Once

func attrIteratorDestroyFunction_Set() {
	attrIteratorDestroyFunction_Once.Do(func() {
		attrIteratorStruct_Set()
		attrIteratorDestroyFunction = attrIteratorStruct.InvokerNew("destroy")
	})
}

// Destroy is a representation of the C type pango_attr_iterator_destroy.
func (recv *AttrIterator) Destroy() {
	attrIteratorDestroyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	attrIteratorDestroyFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_attr_iterator_get' : parameter 'type' of type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_attrs' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_font' : parameter 'desc' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_next' : return type 'gboolean' not supported

var attrIteratorRangeFunction *gi.Function
var attrIteratorRangeFunction_Once sync.Once

func attrIteratorRangeFunction_Set() {
	attrIteratorRangeFunction_Once.Do(func() {
		attrIteratorStruct_Set()
		attrIteratorRangeFunction = attrIteratorStruct.InvokerNew("range")
	})
}

// Range is a representation of the C type pango_attr_iterator_range.
func (recv *AttrIterator) Range() (int32, int32) {
	attrIteratorRangeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	attrIteratorRangeFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var attrLanguageStruct *gi.Struct
var attrLanguageStruct_Once sync.Once

func attrLanguageStruct_Set() {
	attrLanguageStruct_Once.Do(func() {
		attrLanguageStruct = gi.StructNew("Pango", "AttrLanguage")
	})
}

type AttrLanguage struct {
	native uintptr
	Attr   *Attribute
	Value  *Language
}

var attrListStruct *gi.Struct
var attrListStruct_Once sync.Once

func attrListStruct_Set() {
	attrListStruct_Once.Do(func() {
		attrListStruct = gi.StructNew("Pango", "AttrList")
	})
}

type AttrList struct {
	native uintptr
}

var attrListNewFunction *gi.Function
var attrListNewFunction_Once sync.Once

func attrListNewFunction_Set() {
	attrListNewFunction_Once.Do(func() {
		attrListStruct_Set()
		attrListNewFunction = attrListStruct.InvokerNew("new")
	})
}

// AttrListNew is a representation of the C type pango_attr_list_new.
func AttrListNew() *AttrList {
	attrListNewFunction_Set()

	ret := attrListNewFunction.Invoke(nil, nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_change' : parameter 'attr' of type 'Attribute' not supported

var attrListCopyFunction *gi.Function
var attrListCopyFunction_Once sync.Once

func attrListCopyFunction_Set() {
	attrListCopyFunction_Once.Do(func() {
		attrListStruct_Set()
		attrListCopyFunction = attrListStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_attr_list_copy.
func (recv *AttrList) Copy() *AttrList {
	attrListCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := attrListCopyFunction.Invoke(inArgs[:], nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_filter' : parameter 'func' of type 'AttrFilterFunc' not supported

var attrListGetIteratorFunction *gi.Function
var attrListGetIteratorFunction_Once sync.Once

func attrListGetIteratorFunction_Set() {
	attrListGetIteratorFunction_Once.Do(func() {
		attrListStruct_Set()
		attrListGetIteratorFunction = attrListStruct.InvokerNew("get_iterator")
	})
}

// GetIterator is a representation of the C type pango_attr_list_get_iterator.
func (recv *AttrList) GetIterator() *AttrIterator {
	attrListGetIteratorFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := attrListGetIteratorFunction.Invoke(inArgs[:], nil)

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_insert' : parameter 'attr' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_list_insert_before' : parameter 'attr' of type 'Attribute' not supported

var attrListRefFunction *gi.Function
var attrListRefFunction_Once sync.Once

func attrListRefFunction_Set() {
	attrListRefFunction_Once.Do(func() {
		attrListStruct_Set()
		attrListRefFunction = attrListStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type pango_attr_list_ref.
func (recv *AttrList) Ref() *AttrList {
	attrListRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := attrListRefFunction.Invoke(inArgs[:], nil)

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_splice' : parameter 'other' of type 'AttrList' not supported

var attrListUnrefFunction *gi.Function
var attrListUnrefFunction_Once sync.Once

func attrListUnrefFunction_Set() {
	attrListUnrefFunction_Once.Do(func() {
		attrListStruct_Set()
		attrListUnrefFunction = attrListStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type pango_attr_list_unref.
func (recv *AttrList) Unref() {
	attrListUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	attrListUnrefFunction.Invoke(inArgs[:], nil)

}

var attrShapeStruct *gi.Struct
var attrShapeStruct_Once sync.Once

func attrShapeStruct_Set() {
	attrShapeStruct_Once.Do(func() {
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
var attrSizeStruct_Once sync.Once

func attrSizeStruct_Set() {
	attrSizeStruct_Once.Do(func() {
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
var attrStringStruct_Once sync.Once

func attrStringStruct_Set() {
	attrStringStruct_Once.Do(func() {
		attrStringStruct = gi.StructNew("Pango", "AttrString")
	})
}

type AttrString struct {
	native uintptr
	Attr   *Attribute
	Value  string
}

var attributeStruct *gi.Struct
var attributeStruct_Once sync.Once

func attributeStruct_Set() {
	attributeStruct_Once.Do(func() {
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
var attributeCopyFunction_Once sync.Once

func attributeCopyFunction_Set() {
	attributeCopyFunction_Once.Do(func() {
		attributeStruct_Set()
		attributeCopyFunction = attributeStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_attribute_copy.
func (recv *Attribute) Copy() *Attribute {
	attributeCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := attributeCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var attributeDestroyFunction *gi.Function
var attributeDestroyFunction_Once sync.Once

func attributeDestroyFunction_Set() {
	attributeDestroyFunction_Once.Do(func() {
		attributeStruct_Set()
		attributeDestroyFunction = attributeStruct.InvokerNew("destroy")
	})
}

// Destroy is a representation of the C type pango_attribute_destroy.
func (recv *Attribute) Destroy() {
	attributeDestroyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	attributeDestroyFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_attribute_equal' : parameter 'attr2' of type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attribute_init' : parameter 'klass' of type 'AttrClass' not supported

var colorStruct *gi.Struct
var colorStruct_Once sync.Once

func colorStruct_Set() {
	colorStruct_Once.Do(func() {
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
var colorCopyFunction_Once sync.Once

func colorCopyFunction_Set() {
	colorCopyFunction_Once.Do(func() {
		colorStruct_Set()
		colorCopyFunction = colorStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_color_copy.
func (recv *Color) Copy() *Color {
	colorCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := colorCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Color{native: ret.Pointer()}

	return retGo
}

var colorFreeFunction *gi.Function
var colorFreeFunction_Once sync.Once

func colorFreeFunction_Set() {
	colorFreeFunction_Once.Do(func() {
		colorStruct_Set()
		colorFreeFunction = colorStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_color_free.
func (recv *Color) Free() {
	colorFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	colorFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_color_parse' : return type 'gboolean' not supported

var colorToStringFunction *gi.Function
var colorToStringFunction_Once sync.Once

func colorToStringFunction_Set() {
	colorToStringFunction_Once.Do(func() {
		colorStruct_Set()
		colorToStringFunction = colorStruct.InvokerNew("to_string")
	})
}

// ToString is a representation of the C type pango_color_to_string.
func (recv *Color) ToString() string {
	colorToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := colorToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var contextClassStruct *gi.Struct
var contextClassStruct_Once sync.Once

func contextClassStruct_Set() {
	contextClassStruct_Once.Do(func() {
		contextClassStruct = gi.StructNew("Pango", "ContextClass")
	})
}

type ContextClass struct {
	native uintptr
}

var coverageStruct *gi.Struct
var coverageStruct_Once sync.Once

func coverageStruct_Set() {
	coverageStruct_Once.Do(func() {
		coverageStruct = gi.StructNew("Pango", "Coverage")
	})
}

type Coverage struct {
	native uintptr
}

var coverageCopyFunction *gi.Function
var coverageCopyFunction_Once sync.Once

func coverageCopyFunction_Set() {
	coverageCopyFunction_Once.Do(func() {
		coverageStruct_Set()
		coverageCopyFunction = coverageStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_coverage_copy.
func (recv *Coverage) Copy() *Coverage {
	coverageCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := coverageCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_get' : return type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_max' : parameter 'other' of type 'Coverage' not supported

var coverageRefFunction *gi.Function
var coverageRefFunction_Once sync.Once

func coverageRefFunction_Set() {
	coverageRefFunction_Once.Do(func() {
		coverageStruct_Set()
		coverageRefFunction = coverageStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type pango_coverage_ref.
func (recv *Coverage) Ref() *Coverage {
	coverageRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := coverageRefFunction.Invoke(inArgs[:], nil)

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_coverage_set' : parameter 'level' of type 'CoverageLevel' not supported

// UNSUPPORTED : C value 'pango_coverage_to_bytes' : parameter 'bytes' has no type

var coverageUnrefFunction *gi.Function
var coverageUnrefFunction_Once sync.Once

func coverageUnrefFunction_Set() {
	coverageUnrefFunction_Once.Do(func() {
		coverageStruct_Set()
		coverageUnrefFunction = coverageStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type pango_coverage_unref.
func (recv *Coverage) Unref() {
	coverageUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	coverageUnrefFunction.Invoke(inArgs[:], nil)

}

var engineClassStruct *gi.Struct
var engineClassStruct_Once sync.Once

func engineClassStruct_Set() {
	engineClassStruct_Once.Do(func() {
		engineClassStruct = gi.StructNew("Pango", "EngineClass")
	})
}

type EngineClass struct {
	native uintptr
}

var engineInfoStruct *gi.Struct
var engineInfoStruct_Once sync.Once

func engineInfoStruct_Set() {
	engineInfoStruct_Once.Do(func() {
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
var engineLangClassStruct_Once sync.Once

func engineLangClassStruct_Set() {
	engineLangClassStruct_Once.Do(func() {
		engineLangClassStruct = gi.StructNew("Pango", "EngineLangClass")
	})
}

type EngineLangClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_break' : missing Type
}

var engineScriptInfoStruct *gi.Struct
var engineScriptInfoStruct_Once sync.Once

func engineScriptInfoStruct_Set() {
	engineScriptInfoStruct_Once.Do(func() {
		engineScriptInfoStruct = gi.StructNew("Pango", "EngineScriptInfo")
	})
}

type EngineScriptInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'script' : no Go type for 'Script'
	Langs string
}

var engineShapeClassStruct *gi.Struct
var engineShapeClassStruct_Once sync.Once

func engineShapeClassStruct_Set() {
	engineShapeClassStruct_Once.Do(func() {
		engineShapeClassStruct = gi.StructNew("Pango", "EngineShapeClass")
	})
}

type EngineShapeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'script_shape' : missing Type
	// UNSUPPORTED : C value 'covers' : missing Type
}

var fontClassStruct *gi.Struct
var fontClassStruct_Once sync.Once

func fontClassStruct_Set() {
	fontClassStruct_Once.Do(func() {
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
var fontDescriptionStruct_Once sync.Once

func fontDescriptionStruct_Set() {
	fontDescriptionStruct_Once.Do(func() {
		fontDescriptionStruct = gi.StructNew("Pango", "FontDescription")
	})
}

type FontDescription struct {
	native uintptr
}

var fontDescriptionNewFunction *gi.Function
var fontDescriptionNewFunction_Once sync.Once

func fontDescriptionNewFunction_Set() {
	fontDescriptionNewFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionNewFunction = fontDescriptionStruct.InvokerNew("new")
	})
}

// FontDescriptionNew is a representation of the C type pango_font_description_new.
func FontDescriptionNew() *FontDescription {
	fontDescriptionNewFunction_Set()

	ret := fontDescriptionNewFunction.Invoke(nil, nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_better_match' : parameter 'old_match' of type 'FontDescription' not supported

var fontDescriptionCopyFunction *gi.Function
var fontDescriptionCopyFunction_Once sync.Once

func fontDescriptionCopyFunction_Set() {
	fontDescriptionCopyFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionCopyFunction = fontDescriptionStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_font_description_copy.
func (recv *FontDescription) Copy() *FontDescription {
	fontDescriptionCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionCopyFunction.Invoke(inArgs[:], nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var fontDescriptionCopyStaticFunction *gi.Function
var fontDescriptionCopyStaticFunction_Once sync.Once

func fontDescriptionCopyStaticFunction_Set() {
	fontDescriptionCopyStaticFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionCopyStaticFunction = fontDescriptionStruct.InvokerNew("copy_static")
	})
}

// CopyStatic is a representation of the C type pango_font_description_copy_static.
func (recv *FontDescription) CopyStatic() *FontDescription {
	fontDescriptionCopyStaticFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionCopyStaticFunction.Invoke(inArgs[:], nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_equal' : parameter 'desc2' of type 'FontDescription' not supported

var fontDescriptionFreeFunction *gi.Function
var fontDescriptionFreeFunction_Once sync.Once

func fontDescriptionFreeFunction_Set() {
	fontDescriptionFreeFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionFreeFunction = fontDescriptionStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_font_description_free.
func (recv *FontDescription) Free() {
	fontDescriptionFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	fontDescriptionFreeFunction.Invoke(inArgs[:], nil)

}

var fontDescriptionGetFamilyFunction *gi.Function
var fontDescriptionGetFamilyFunction_Once sync.Once

func fontDescriptionGetFamilyFunction_Set() {
	fontDescriptionGetFamilyFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionGetFamilyFunction = fontDescriptionStruct.InvokerNew("get_family")
	})
}

// GetFamily is a representation of the C type pango_font_description_get_family.
func (recv *FontDescription) GetFamily() string {
	fontDescriptionGetFamilyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionGetFamilyFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_gravity' : return type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_font_description_get_set_fields' : return type 'FontMask' not supported

var fontDescriptionGetSizeFunction *gi.Function
var fontDescriptionGetSizeFunction_Once sync.Once

func fontDescriptionGetSizeFunction_Set() {
	fontDescriptionGetSizeFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionGetSizeFunction = fontDescriptionStruct.InvokerNew("get_size")
	})
}

// GetSize is a representation of the C type pango_font_description_get_size.
func (recv *FontDescription) GetSize() int32 {
	fontDescriptionGetSizeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionGetSizeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_size_is_absolute' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_font_description_get_stretch' : return type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_get_style' : return type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_get_variant' : return type 'Variant' not supported

var fontDescriptionGetVariationsFunction *gi.Function
var fontDescriptionGetVariationsFunction_Once sync.Once

func fontDescriptionGetVariationsFunction_Set() {
	fontDescriptionGetVariationsFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionGetVariationsFunction = fontDescriptionStruct.InvokerNew("get_variations")
	})
}

// GetVariations is a representation of the C type pango_font_description_get_variations.
func (recv *FontDescription) GetVariations() string {
	fontDescriptionGetVariationsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionGetVariationsFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_weight' : return type 'Weight' not supported

var fontDescriptionHashFunction *gi.Function
var fontDescriptionHashFunction_Once sync.Once

func fontDescriptionHashFunction_Set() {
	fontDescriptionHashFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionHashFunction = fontDescriptionStruct.InvokerNew("hash")
	})
}

// Hash is a representation of the C type pango_font_description_hash.
func (recv *FontDescription) Hash() uint32 {
	fontDescriptionHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionHashFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_merge' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_merge_static' : parameter 'desc_to_merge' of type 'FontDescription' not supported

// UNSUPPORTED : C value 'pango_font_description_set_absolute_size' : parameter 'size' of type 'gdouble' not supported

var fontDescriptionSetFamilyFunction *gi.Function
var fontDescriptionSetFamilyFunction_Once sync.Once

func fontDescriptionSetFamilyFunction_Set() {
	fontDescriptionSetFamilyFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionSetFamilyFunction = fontDescriptionStruct.InvokerNew("set_family")
	})
}

// SetFamily is a representation of the C type pango_font_description_set_family.
func (recv *FontDescription) SetFamily(family string) {
	fontDescriptionSetFamilyFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	fontDescriptionSetFamilyFunction.Invoke(inArgs[:], nil)

}

var fontDescriptionSetFamilyStaticFunction *gi.Function
var fontDescriptionSetFamilyStaticFunction_Once sync.Once

func fontDescriptionSetFamilyStaticFunction_Set() {
	fontDescriptionSetFamilyStaticFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionSetFamilyStaticFunction = fontDescriptionStruct.InvokerNew("set_family_static")
	})
}

// SetFamilyStatic is a representation of the C type pango_font_description_set_family_static.
func (recv *FontDescription) SetFamilyStatic(family string) {
	fontDescriptionSetFamilyStaticFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	fontDescriptionSetFamilyStaticFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_gravity' : parameter 'gravity' of type 'Gravity' not supported

var fontDescriptionSetSizeFunction *gi.Function
var fontDescriptionSetSizeFunction_Once sync.Once

func fontDescriptionSetSizeFunction_Set() {
	fontDescriptionSetSizeFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionSetSizeFunction = fontDescriptionStruct.InvokerNew("set_size")
	})
}

// SetSize is a representation of the C type pango_font_description_set_size.
func (recv *FontDescription) SetSize(size int32) {
	fontDescriptionSetSizeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(size)

	fontDescriptionSetSizeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_stretch' : parameter 'stretch' of type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_font_description_set_style' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'pango_font_description_set_variant' : parameter 'variant' of type 'Variant' not supported

var fontDescriptionSetVariationsFunction *gi.Function
var fontDescriptionSetVariationsFunction_Once sync.Once

func fontDescriptionSetVariationsFunction_Set() {
	fontDescriptionSetVariationsFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionSetVariationsFunction = fontDescriptionStruct.InvokerNew("set_variations")
	})
}

// SetVariations is a representation of the C type pango_font_description_set_variations.
func (recv *FontDescription) SetVariations(settings string) {
	fontDescriptionSetVariationsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	fontDescriptionSetVariationsFunction.Invoke(inArgs[:], nil)

}

var fontDescriptionSetVariationsStaticFunction *gi.Function
var fontDescriptionSetVariationsStaticFunction_Once sync.Once

func fontDescriptionSetVariationsStaticFunction_Set() {
	fontDescriptionSetVariationsStaticFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionSetVariationsStaticFunction = fontDescriptionStruct.InvokerNew("set_variations_static")
	})
}

// SetVariationsStatic is a representation of the C type pango_font_description_set_variations_static.
func (recv *FontDescription) SetVariationsStatic(settings string) {
	fontDescriptionSetVariationsStaticFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	fontDescriptionSetVariationsStaticFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_font_description_set_weight' : parameter 'weight' of type 'Weight' not supported

var fontDescriptionToFilenameFunction *gi.Function
var fontDescriptionToFilenameFunction_Once sync.Once

func fontDescriptionToFilenameFunction_Set() {
	fontDescriptionToFilenameFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionToFilenameFunction = fontDescriptionStruct.InvokerNew("to_filename")
	})
}

// ToFilename is a representation of the C type pango_font_description_to_filename.
func (recv *FontDescription) ToFilename() string {
	fontDescriptionToFilenameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionToFilenameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var fontDescriptionToStringFunction *gi.Function
var fontDescriptionToStringFunction_Once sync.Once

func fontDescriptionToStringFunction_Set() {
	fontDescriptionToStringFunction_Once.Do(func() {
		fontDescriptionStruct_Set()
		fontDescriptionToStringFunction = fontDescriptionStruct.InvokerNew("to_string")
	})
}

// ToString is a representation of the C type pango_font_description_to_string.
func (recv *FontDescription) ToString() string {
	fontDescriptionToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontDescriptionToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_unset_fields' : parameter 'to_unset' of type 'FontMask' not supported

var fontFaceClassStruct *gi.Struct
var fontFaceClassStruct_Once sync.Once

func fontFaceClassStruct_Set() {
	fontFaceClassStruct_Once.Do(func() {
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
var fontFamilyClassStruct_Once sync.Once

func fontFamilyClassStruct_Set() {
	fontFamilyClassStruct_Once.Do(func() {
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
var fontMapClassStruct_Once sync.Once

func fontMapClassStruct_Set() {
	fontMapClassStruct_Once.Do(func() {
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
var fontMetricsStruct_Once sync.Once

func fontMetricsStruct_Set() {
	fontMetricsStruct_Once.Do(func() {
		fontMetricsStruct = gi.StructNew("Pango", "FontMetrics")
	})
}

type FontMetrics struct {
	native uintptr
}

var fontMetricsNewFunction *gi.Function
var fontMetricsNewFunction_Once sync.Once

func fontMetricsNewFunction_Set() {
	fontMetricsNewFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsNewFunction = fontMetricsStruct.InvokerNew("new")
	})
}

// FontMetricsNew is a representation of the C type pango_font_metrics_new.
func FontMetricsNew() *FontMetrics {
	fontMetricsNewFunction_Set()

	ret := fontMetricsNewFunction.Invoke(nil, nil)

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var fontMetricsGetApproximateCharWidthFunction *gi.Function
var fontMetricsGetApproximateCharWidthFunction_Once sync.Once

func fontMetricsGetApproximateCharWidthFunction_Set() {
	fontMetricsGetApproximateCharWidthFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetApproximateCharWidthFunction = fontMetricsStruct.InvokerNew("get_approximate_char_width")
	})
}

// GetApproximateCharWidth is a representation of the C type pango_font_metrics_get_approximate_char_width.
func (recv *FontMetrics) GetApproximateCharWidth() int32 {
	fontMetricsGetApproximateCharWidthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetApproximateCharWidthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetApproximateDigitWidthFunction *gi.Function
var fontMetricsGetApproximateDigitWidthFunction_Once sync.Once

func fontMetricsGetApproximateDigitWidthFunction_Set() {
	fontMetricsGetApproximateDigitWidthFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetApproximateDigitWidthFunction = fontMetricsStruct.InvokerNew("get_approximate_digit_width")
	})
}

// GetApproximateDigitWidth is a representation of the C type pango_font_metrics_get_approximate_digit_width.
func (recv *FontMetrics) GetApproximateDigitWidth() int32 {
	fontMetricsGetApproximateDigitWidthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetApproximateDigitWidthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetAscentFunction *gi.Function
var fontMetricsGetAscentFunction_Once sync.Once

func fontMetricsGetAscentFunction_Set() {
	fontMetricsGetAscentFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetAscentFunction = fontMetricsStruct.InvokerNew("get_ascent")
	})
}

// GetAscent is a representation of the C type pango_font_metrics_get_ascent.
func (recv *FontMetrics) GetAscent() int32 {
	fontMetricsGetAscentFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetAscentFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetDescentFunction *gi.Function
var fontMetricsGetDescentFunction_Once sync.Once

func fontMetricsGetDescentFunction_Set() {
	fontMetricsGetDescentFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetDescentFunction = fontMetricsStruct.InvokerNew("get_descent")
	})
}

// GetDescent is a representation of the C type pango_font_metrics_get_descent.
func (recv *FontMetrics) GetDescent() int32 {
	fontMetricsGetDescentFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetDescentFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughPositionFunction *gi.Function
var fontMetricsGetStrikethroughPositionFunction_Once sync.Once

func fontMetricsGetStrikethroughPositionFunction_Set() {
	fontMetricsGetStrikethroughPositionFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetStrikethroughPositionFunction = fontMetricsStruct.InvokerNew("get_strikethrough_position")
	})
}

// GetStrikethroughPosition is a representation of the C type pango_font_metrics_get_strikethrough_position.
func (recv *FontMetrics) GetStrikethroughPosition() int32 {
	fontMetricsGetStrikethroughPositionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetStrikethroughPositionFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughThicknessFunction *gi.Function
var fontMetricsGetStrikethroughThicknessFunction_Once sync.Once

func fontMetricsGetStrikethroughThicknessFunction_Set() {
	fontMetricsGetStrikethroughThicknessFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetStrikethroughThicknessFunction = fontMetricsStruct.InvokerNew("get_strikethrough_thickness")
	})
}

// GetStrikethroughThickness is a representation of the C type pango_font_metrics_get_strikethrough_thickness.
func (recv *FontMetrics) GetStrikethroughThickness() int32 {
	fontMetricsGetStrikethroughThicknessFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetStrikethroughThicknessFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlinePositionFunction *gi.Function
var fontMetricsGetUnderlinePositionFunction_Once sync.Once

func fontMetricsGetUnderlinePositionFunction_Set() {
	fontMetricsGetUnderlinePositionFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetUnderlinePositionFunction = fontMetricsStruct.InvokerNew("get_underline_position")
	})
}

// GetUnderlinePosition is a representation of the C type pango_font_metrics_get_underline_position.
func (recv *FontMetrics) GetUnderlinePosition() int32 {
	fontMetricsGetUnderlinePositionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetUnderlinePositionFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlineThicknessFunction *gi.Function
var fontMetricsGetUnderlineThicknessFunction_Once sync.Once

func fontMetricsGetUnderlineThicknessFunction_Set() {
	fontMetricsGetUnderlineThicknessFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsGetUnderlineThicknessFunction = fontMetricsStruct.InvokerNew("get_underline_thickness")
	})
}

// GetUnderlineThickness is a representation of the C type pango_font_metrics_get_underline_thickness.
func (recv *FontMetrics) GetUnderlineThickness() int32 {
	fontMetricsGetUnderlineThicknessFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsGetUnderlineThicknessFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var fontMetricsRefFunction *gi.Function
var fontMetricsRefFunction_Once sync.Once

func fontMetricsRefFunction_Set() {
	fontMetricsRefFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsRefFunction = fontMetricsStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type pango_font_metrics_ref.
func (recv *FontMetrics) Ref() *FontMetrics {
	fontMetricsRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := fontMetricsRefFunction.Invoke(inArgs[:], nil)

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var fontMetricsUnrefFunction *gi.Function
var fontMetricsUnrefFunction_Once sync.Once

func fontMetricsUnrefFunction_Set() {
	fontMetricsUnrefFunction_Once.Do(func() {
		fontMetricsStruct_Set()
		fontMetricsUnrefFunction = fontMetricsStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type pango_font_metrics_unref.
func (recv *FontMetrics) Unref() {
	fontMetricsUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	fontMetricsUnrefFunction.Invoke(inArgs[:], nil)

}

var fontsetClassStruct *gi.Struct
var fontsetClassStruct_Once sync.Once

func fontsetClassStruct_Set() {
	fontsetClassStruct_Once.Do(func() {
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
var fontsetSimpleClassStruct_Once sync.Once

func fontsetSimpleClassStruct_Set() {
	fontsetSimpleClassStruct_Once.Do(func() {
		fontsetSimpleClassStruct = gi.StructNew("Pango", "FontsetSimpleClass")
	})
}

type FontsetSimpleClass struct {
	native uintptr
}

var glyphGeometryStruct *gi.Struct
var glyphGeometryStruct_Once sync.Once

func glyphGeometryStruct_Set() {
	glyphGeometryStruct_Once.Do(func() {
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
var glyphInfoStruct_Once sync.Once

func glyphInfoStruct_Set() {
	glyphInfoStruct_Once.Do(func() {
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
var glyphItemStruct_Once sync.Once

func glyphItemStruct_Set() {
	glyphItemStruct_Once.Do(func() {
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
var glyphItemCopyFunction_Once sync.Once

func glyphItemCopyFunction_Set() {
	glyphItemCopyFunction_Once.Do(func() {
		glyphItemStruct_Set()
		glyphItemCopyFunction = glyphItemStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_glyph_item_copy.
func (recv *GlyphItem) Copy() *GlyphItem {
	glyphItemCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := glyphItemCopyFunction.Invoke(inArgs[:], nil)

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

var glyphItemFreeFunction *gi.Function
var glyphItemFreeFunction_Once sync.Once

func glyphItemFreeFunction_Set() {
	glyphItemFreeFunction_Once.Do(func() {
		glyphItemStruct_Set()
		glyphItemFreeFunction = glyphItemStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	glyphItemFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	glyphItemFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_item_get_logical_widths' : parameter 'logical_widths' has no type

// UNSUPPORTED : C value 'pango_glyph_item_letter_space' : parameter 'log_attrs' has no type

var glyphItemSplitFunction *gi.Function
var glyphItemSplitFunction_Once sync.Once

func glyphItemSplitFunction_Set() {
	glyphItemSplitFunction_Once.Do(func() {
		glyphItemStruct_Set()
		glyphItemSplitFunction = glyphItemStruct.InvokerNew("split")
	})
}

// Split is a representation of the C type pango_glyph_item_split.
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	glyphItemSplitFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(splitIndex)

	ret := glyphItemSplitFunction.Invoke(inArgs[:], nil)

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo
}

var glyphItemIterStruct *gi.Struct
var glyphItemIterStruct_Once sync.Once

func glyphItemIterStruct_Set() {
	glyphItemIterStruct_Once.Do(func() {
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
var glyphItemIterCopyFunction_Once sync.Once

func glyphItemIterCopyFunction_Set() {
	glyphItemIterCopyFunction_Once.Do(func() {
		glyphItemIterStruct_Set()
		glyphItemIterCopyFunction = glyphItemIterStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	glyphItemIterCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := glyphItemIterCopyFunction.Invoke(inArgs[:], nil)

	retGo := &GlyphItemIter{native: ret.Pointer()}

	return retGo
}

var glyphItemIterFreeFunction *gi.Function
var glyphItemIterFreeFunction_Once sync.Once

func glyphItemIterFreeFunction_Set() {
	glyphItemIterFreeFunction_Once.Do(func() {
		glyphItemIterStruct_Set()
		glyphItemIterFreeFunction = glyphItemIterStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_glyph_item_iter_free.
func (recv *GlyphItemIter) Free() {
	glyphItemIterFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	glyphItemIterFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_end' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_init_start' : parameter 'glyph_item' of type 'GlyphItem' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_next_cluster' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_glyph_item_iter_prev_cluster' : return type 'gboolean' not supported

var glyphStringStruct *gi.Struct
var glyphStringStruct_Once sync.Once

func glyphStringStruct_Set() {
	glyphStringStruct_Once.Do(func() {
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
var glyphStringNewFunction_Once sync.Once

func glyphStringNewFunction_Set() {
	glyphStringNewFunction_Once.Do(func() {
		glyphStringStruct_Set()
		glyphStringNewFunction = glyphStringStruct.InvokerNew("new")
	})
}

// GlyphStringNew is a representation of the C type pango_glyph_string_new.
func GlyphStringNew() *GlyphString {
	glyphStringNewFunction_Set()

	ret := glyphStringNewFunction.Invoke(nil, nil)

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

var glyphStringCopyFunction *gi.Function
var glyphStringCopyFunction_Once sync.Once

func glyphStringCopyFunction_Set() {
	glyphStringCopyFunction_Once.Do(func() {
		glyphStringStruct_Set()
		glyphStringCopyFunction = glyphStringStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_glyph_string_copy.
func (recv *GlyphString) Copy() *GlyphString {
	glyphStringCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := glyphStringCopyFunction.Invoke(inArgs[:], nil)

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_extents' : parameter 'font' of type 'Font' not supported

// UNSUPPORTED : C value 'pango_glyph_string_extents_range' : parameter 'font' of type 'Font' not supported

var glyphStringFreeFunction *gi.Function
var glyphStringFreeFunction_Once sync.Once

func glyphStringFreeFunction_Set() {
	glyphStringFreeFunction_Once.Do(func() {
		glyphStringStruct_Set()
		glyphStringFreeFunction = glyphStringStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_glyph_string_free.
func (recv *GlyphString) Free() {
	glyphStringFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	glyphStringFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_string_get_logical_widths' : parameter 'logical_widths' has no type

var glyphStringGetWidthFunction *gi.Function
var glyphStringGetWidthFunction_Once sync.Once

func glyphStringGetWidthFunction_Set() {
	glyphStringGetWidthFunction_Once.Do(func() {
		glyphStringStruct_Set()
		glyphStringGetWidthFunction = glyphStringStruct.InvokerNew("get_width")
	})
}

// GetWidth is a representation of the C type pango_glyph_string_get_width.
func (recv *GlyphString) GetWidth() int32 {
	glyphStringGetWidthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := glyphStringGetWidthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_glyph_string_index_to_x' : parameter 'analysis' of type 'Analysis' not supported

var glyphStringSetSizeFunction *gi.Function
var glyphStringSetSizeFunction_Once sync.Once

func glyphStringSetSizeFunction_Set() {
	glyphStringSetSizeFunction_Once.Do(func() {
		glyphStringStruct_Set()
		glyphStringSetSizeFunction = glyphStringStruct.InvokerNew("set_size")
	})
}

// SetSize is a representation of the C type pango_glyph_string_set_size.
func (recv *GlyphString) SetSize(newLen int32) {
	glyphStringSetSizeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newLen)

	glyphStringSetSizeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_glyph_string_x_to_index' : parameter 'analysis' of type 'Analysis' not supported

var glyphVisAttrStruct *gi.Struct
var glyphVisAttrStruct_Once sync.Once

func glyphVisAttrStruct_Set() {
	glyphVisAttrStruct_Once.Do(func() {
		glyphVisAttrStruct = gi.StructNew("Pango", "GlyphVisAttr")
	})
}

type GlyphVisAttr struct {
	native         uintptr
	IsClusterStart uint32
}

var includedModuleStruct *gi.Struct
var includedModuleStruct_Once sync.Once

func includedModuleStruct_Set() {
	includedModuleStruct_Once.Do(func() {
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
var itemStruct_Once sync.Once

func itemStruct_Set() {
	itemStruct_Once.Do(func() {
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
var itemNewFunction_Once sync.Once

func itemNewFunction_Set() {
	itemNewFunction_Once.Do(func() {
		itemStruct_Set()
		itemNewFunction = itemStruct.InvokerNew("new")
	})
}

// ItemNew is a representation of the C type pango_item_new.
func ItemNew() *Item {
	itemNewFunction_Set()

	ret := itemNewFunction.Invoke(nil, nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var itemCopyFunction *gi.Function
var itemCopyFunction_Once sync.Once

func itemCopyFunction_Set() {
	itemCopyFunction_Once.Do(func() {
		itemStruct_Set()
		itemCopyFunction = itemStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_item_copy.
func (recv *Item) Copy() *Item {
	itemCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := itemCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var itemFreeFunction *gi.Function
var itemFreeFunction_Once sync.Once

func itemFreeFunction_Set() {
	itemFreeFunction_Once.Do(func() {
		itemStruct_Set()
		itemFreeFunction = itemStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_item_free.
func (recv *Item) Free() {
	itemFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	itemFreeFunction.Invoke(inArgs[:], nil)

}

var itemSplitFunction *gi.Function
var itemSplitFunction_Once sync.Once

func itemSplitFunction_Set() {
	itemSplitFunction_Once.Do(func() {
		itemStruct_Set()
		itemSplitFunction = itemStruct.InvokerNew("split")
	})
}

// Split is a representation of the C type pango_item_split.
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	itemSplitFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(splitIndex)
	inArgs[2].SetInt32(splitOffset)

	ret := itemSplitFunction.Invoke(inArgs[:], nil)

	retGo := &Item{native: ret.Pointer()}

	return retGo
}

var languageStruct *gi.Struct
var languageStruct_Once sync.Once

func languageStruct_Set() {
	languageStruct_Once.Do(func() {
		languageStruct = gi.StructNew("Pango", "Language")
	})
}

type Language struct {
	native uintptr
}

var languageGetSampleStringFunction *gi.Function
var languageGetSampleStringFunction_Once sync.Once

func languageGetSampleStringFunction_Set() {
	languageGetSampleStringFunction_Once.Do(func() {
		languageStruct_Set()
		languageGetSampleStringFunction = languageStruct.InvokerNew("get_sample_string")
	})
}

// GetSampleString is a representation of the C type pango_language_get_sample_string.
func (recv *Language) GetSampleString() string {
	languageGetSampleStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := languageGetSampleStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var languageGetScriptsFunction *gi.Function
var languageGetScriptsFunction_Once sync.Once

func languageGetScriptsFunction_Set() {
	languageGetScriptsFunction_Once.Do(func() {
		languageStruct_Set()
		languageGetScriptsFunction = languageStruct.InvokerNew("get_scripts")
	})
}

// GetScripts is a representation of the C type pango_language_get_scripts.
func (recv *Language) GetScripts() int32 {
	languageGetScriptsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	languageGetScriptsFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'pango_language_includes_script' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_language_matches' : return type 'gboolean' not supported

var languageToStringFunction *gi.Function
var languageToStringFunction_Once sync.Once

func languageToStringFunction_Set() {
	languageToStringFunction_Once.Do(func() {
		languageStruct_Set()
		languageToStringFunction = languageStruct.InvokerNew("to_string")
	})
}

// ToString is a representation of the C type pango_language_to_string.
func (recv *Language) ToString() string {
	languageToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := languageToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var layoutClassStruct *gi.Struct
var layoutClassStruct_Once sync.Once

func layoutClassStruct_Set() {
	layoutClassStruct_Once.Do(func() {
		layoutClassStruct = gi.StructNew("Pango", "LayoutClass")
	})
}

type LayoutClass struct {
	native uintptr
}

var layoutIterStruct *gi.Struct
var layoutIterStruct_Once sync.Once

func layoutIterStruct_Set() {
	layoutIterStruct_Once.Do(func() {
		layoutIterStruct = gi.StructNew("Pango", "LayoutIter")
	})
}

type LayoutIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_layout_iter_at_last_line' : return type 'gboolean' not supported

var layoutIterCopyFunction *gi.Function
var layoutIterCopyFunction_Once sync.Once

func layoutIterCopyFunction_Set() {
	layoutIterCopyFunction_Once.Do(func() {
		layoutIterStruct_Set()
		layoutIterCopyFunction = layoutIterStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_layout_iter_copy.
func (recv *LayoutIter) Copy() *LayoutIter {
	layoutIterCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := layoutIterCopyFunction.Invoke(inArgs[:], nil)

	retGo := &LayoutIter{native: ret.Pointer()}

	return retGo
}

var layoutIterFreeFunction *gi.Function
var layoutIterFreeFunction_Once sync.Once

func layoutIterFreeFunction_Set() {
	layoutIterFreeFunction_Once.Do(func() {
		layoutIterStruct_Set()
		layoutIterFreeFunction = layoutIterStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_layout_iter_free.
func (recv *LayoutIter) Free() {
	layoutIterFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	layoutIterFreeFunction.Invoke(inArgs[:], nil)

}

var layoutIterGetBaselineFunction *gi.Function
var layoutIterGetBaselineFunction_Once sync.Once

func layoutIterGetBaselineFunction_Set() {
	layoutIterGetBaselineFunction_Once.Do(func() {
		layoutIterStruct_Set()
		layoutIterGetBaselineFunction = layoutIterStruct.InvokerNew("get_baseline")
	})
}

// GetBaseline is a representation of the C type pango_layout_iter_get_baseline.
func (recv *LayoutIter) GetBaseline() int32 {
	layoutIterGetBaselineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := layoutIterGetBaselineFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_char_extents' : parameter 'logical_rect' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_cluster_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetIndexFunction *gi.Function
var layoutIterGetIndexFunction_Once sync.Once

func layoutIterGetIndexFunction_Set() {
	layoutIterGetIndexFunction_Once.Do(func() {
		layoutIterStruct_Set()
		layoutIterGetIndexFunction = layoutIterStruct.InvokerNew("get_index")
	})
}

// GetIndex is a representation of the C type pango_layout_iter_get_index.
func (recv *LayoutIter) GetIndex() int32 {
	layoutIterGetIndexFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := layoutIterGetIndexFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_layout' : return type 'Layout' not supported

// UNSUPPORTED : C value 'pango_layout_iter_get_layout_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetLineFunction *gi.Function
var layoutIterGetLineFunction_Once sync.Once

func layoutIterGetLineFunction_Set() {
	layoutIterGetLineFunction_Once.Do(func() {
		layoutIterStruct_Set()
		layoutIterGetLineFunction = layoutIterStruct.InvokerNew("get_line")
	})
}

// GetLine is a representation of the C type pango_layout_iter_get_line.
func (recv *LayoutIter) GetLine() *LayoutLine {
	layoutIterGetLineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := layoutIterGetLineFunction.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_iter_get_line_extents' : parameter 'ink_rect' of type 'Rectangle' not supported

var layoutIterGetLineReadonlyFunction *gi.Function
var layoutIterGetLineReadonlyFunction_Once sync.Once

func layoutIterGetLineReadonlyFunction_Set() {
	layoutIterGetLineReadonlyFunction_Once.Do(func() {
		layoutIterStruct_Set()
		layoutIterGetLineReadonlyFunction = layoutIterStruct.InvokerNew("get_line_readonly")
	})
}

// GetLineReadonly is a representation of the C type pango_layout_iter_get_line_readonly.
func (recv *LayoutIter) GetLineReadonly() *LayoutLine {
	layoutIterGetLineReadonlyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := layoutIterGetLineReadonlyFunction.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var layoutIterGetLineYrangeFunction *gi.Function
var layoutIterGetLineYrangeFunction_Once sync.Once

func layoutIterGetLineYrangeFunction_Set() {
	layoutIterGetLineYrangeFunction_Once.Do(func() {
		layoutIterStruct_Set()
		layoutIterGetLineYrangeFunction = layoutIterStruct.InvokerNew("get_line_yrange")
	})
}

// GetLineYrange is a representation of the C type pango_layout_iter_get_line_yrange.
func (recv *LayoutIter) GetLineYrange() (int32, int32) {
	layoutIterGetLineYrangeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	layoutIterGetLineYrangeFunction.Invoke(inArgs[:], outArgs[:])

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

func layoutLineStruct_Set() {
	layoutLineStruct_Once.Do(func() {
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
var layoutLineRefFunction_Once sync.Once

func layoutLineRefFunction_Set() {
	layoutLineRefFunction_Once.Do(func() {
		layoutLineStruct_Set()
		layoutLineRefFunction = layoutLineStruct.InvokerNew("ref")
	})
}

// Ref is a representation of the C type pango_layout_line_ref.
func (recv *LayoutLine) Ref() *LayoutLine {
	layoutLineRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := layoutLineRefFunction.Invoke(inArgs[:], nil)

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var layoutLineUnrefFunction *gi.Function
var layoutLineUnrefFunction_Once sync.Once

func layoutLineUnrefFunction_Set() {
	layoutLineUnrefFunction_Once.Do(func() {
		layoutLineStruct_Set()
		layoutLineUnrefFunction = layoutLineStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type pango_layout_line_unref.
func (recv *LayoutLine) Unref() {
	layoutLineUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	layoutLineUnrefFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_layout_line_x_to_index' : return type 'gboolean' not supported

var logAttrStruct *gi.Struct
var logAttrStruct_Once sync.Once

func logAttrStruct_Set() {
	logAttrStruct_Once.Do(func() {
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
var mapStruct_Once sync.Once

func mapStruct_Set() {
	mapStruct_Once.Do(func() {
		mapStruct = gi.StructNew("Pango", "Map")
	})
}

type Map struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_map_get_engine' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_map_get_engines' : parameter 'script' of type 'Script' not supported

var mapEntryStruct *gi.Struct
var mapEntryStruct_Once sync.Once

func mapEntryStruct_Set() {
	mapEntryStruct_Once.Do(func() {
		mapEntryStruct = gi.StructNew("Pango", "MapEntry")
	})
}

type MapEntry struct {
	native uintptr
}

var matrixStruct *gi.Struct
var matrixStruct_Once sync.Once

func matrixStruct_Set() {
	matrixStruct_Once.Do(func() {
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
var matrixCopyFunction_Once sync.Once

func matrixCopyFunction_Set() {
	matrixCopyFunction_Once.Do(func() {
		matrixStruct_Set()
		matrixCopyFunction = matrixStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	matrixCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := matrixCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Matrix{native: ret.Pointer()}

	return retGo
}

var matrixFreeFunction *gi.Function
var matrixFreeFunction_Once sync.Once

func matrixFreeFunction_Set() {
	matrixFreeFunction_Once.Do(func() {
		matrixStruct_Set()
		matrixFreeFunction = matrixStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_matrix_free.
func (recv *Matrix) Free() {
	matrixFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	matrixFreeFunction.Invoke(inArgs[:], nil)

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

func rectangleStruct_Set() {
	rectangleStruct_Once.Do(func() {
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
var rendererClassStruct_Once sync.Once

func rendererClassStruct_Set() {
	rendererClassStruct_Once.Do(func() {
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
var rendererPrivateStruct_Once sync.Once

func rendererPrivateStruct_Set() {
	rendererPrivateStruct_Once.Do(func() {
		rendererPrivateStruct = gi.StructNew("Pango", "RendererPrivate")
	})
}

type RendererPrivate struct {
	native uintptr
}

var scriptIterStruct *gi.Struct
var scriptIterStruct_Once sync.Once

func scriptIterStruct_Set() {
	scriptIterStruct_Once.Do(func() {
		scriptIterStruct = gi.StructNew("Pango", "ScriptIter")
	})
}

type ScriptIter struct {
	native uintptr
}

var scriptIterFreeFunction *gi.Function
var scriptIterFreeFunction_Once sync.Once

func scriptIterFreeFunction_Set() {
	scriptIterFreeFunction_Once.Do(func() {
		scriptIterStruct_Set()
		scriptIterFreeFunction = scriptIterStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_script_iter_free.
func (recv *ScriptIter) Free() {
	scriptIterFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	scriptIterFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_script_iter_get_range' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_script_iter_next' : return type 'gboolean' not supported

var tabArrayStruct *gi.Struct
var tabArrayStruct_Once sync.Once

func tabArrayStruct_Set() {
	tabArrayStruct_Once.Do(func() {
		tabArrayStruct = gi.StructNew("Pango", "TabArray")
	})
}

type TabArray struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_tab_array_new' : parameter 'positions_in_pixels' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_tab_array_new_with_positions' : parameter 'positions_in_pixels' of type 'gboolean' not supported

var tabArrayCopyFunction *gi.Function
var tabArrayCopyFunction_Once sync.Once

func tabArrayCopyFunction_Set() {
	tabArrayCopyFunction_Once.Do(func() {
		tabArrayStruct_Set()
		tabArrayCopyFunction = tabArrayStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type pango_tab_array_copy.
func (recv *TabArray) Copy() *TabArray {
	tabArrayCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := tabArrayCopyFunction.Invoke(inArgs[:], nil)

	retGo := &TabArray{native: ret.Pointer()}

	return retGo
}

var tabArrayFreeFunction *gi.Function
var tabArrayFreeFunction_Once sync.Once

func tabArrayFreeFunction_Set() {
	tabArrayFreeFunction_Once.Do(func() {
		tabArrayStruct_Set()
		tabArrayFreeFunction = tabArrayStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type pango_tab_array_free.
func (recv *TabArray) Free() {
	tabArrayFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	tabArrayFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_tab_array_get_positions_in_pixels' : return type 'gboolean' not supported

var tabArrayGetSizeFunction *gi.Function
var tabArrayGetSizeFunction_Once sync.Once

func tabArrayGetSizeFunction_Set() {
	tabArrayGetSizeFunction_Once.Do(func() {
		tabArrayStruct_Set()
		tabArrayGetSizeFunction = tabArrayStruct.InvokerNew("get_size")
	})
}

// GetSize is a representation of the C type pango_tab_array_get_size.
func (recv *TabArray) GetSize() int32 {
	tabArrayGetSizeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := tabArrayGetSizeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_tab_array_get_tab' : parameter 'alignment' of type 'TabAlign' not supported

// UNSUPPORTED : C value 'pango_tab_array_get_tabs' : parameter 'alignments' of type 'TabAlign' not supported

var tabArrayResizeFunction *gi.Function
var tabArrayResizeFunction_Once sync.Once

func tabArrayResizeFunction_Set() {
	tabArrayResizeFunction_Once.Do(func() {
		tabArrayStruct_Set()
		tabArrayResizeFunction = tabArrayStruct.InvokerNew("resize")
	})
}

// Resize is a representation of the C type pango_tab_array_resize.
func (recv *TabArray) Resize(newSize int32) {
	tabArrayResizeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newSize)

	tabArrayResizeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'pango_tab_array_set_tab' : parameter 'alignment' of type 'TabAlign' not supported

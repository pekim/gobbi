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
	native uintptr
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
func (recv *AttrIterator) Copy() (*AttrIterator, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrIteratorCopyFunction_Set()
	if err == nil {
		ret = attrIteratorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo, err
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
func (recv *AttrIterator) Destroy() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attrIteratorDestroyFunction_Set()
	if err == nil {
		attrIteratorDestroyFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'pango_attr_iterator_get' : parameter 'type' of type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_attrs' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_font' : parameter 'extra_attrs' of type 'GLib.SList' not supported

var attrIteratorNextFunction *gi.Function
var attrIteratorNextFunction_Once sync.Once

func attrIteratorNextFunction_Set() error {
	var err error
	attrIteratorNextFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorNextFunction, err = attrIteratorStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type pango_attr_iterator_next.
func (recv *AttrIterator) Next() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrIteratorNextFunction_Set()
	if err == nil {
		ret = attrIteratorNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *AttrIterator) Range() (int32, int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := attrIteratorRangeFunction_Set()
	if err == nil {
		attrIteratorRangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1, err
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
func AttrListNew() (*AttrList, error) {

	var ret gi.Argument

	err := attrListNewFunction_Set()
	if err == nil {
		ret = attrListNewFunction.Invoke(nil, nil)
	}

	retGo := &AttrList{native: ret.Pointer()}

	return retGo, err
}

var attrListChangeFunction *gi.Function
var attrListChangeFunction_Once sync.Once

func attrListChangeFunction_Set() error {
	var err error
	attrListChangeFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListChangeFunction, err = attrListStruct.InvokerNew("change")
	})
	return err
}

// Change is a representation of the C type pango_attr_list_change.
func (recv *AttrList) Change(attr *Attribute) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr.native)

	err := attrListChangeFunction_Set()
	if err == nil {
		attrListChangeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *AttrList) Copy() (*AttrList, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListCopyFunction_Set()
	if err == nil {
		ret = attrListCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrList{native: ret.Pointer()}

	return retGo, err
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
func (recv *AttrList) GetIterator() (*AttrIterator, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListGetIteratorFunction_Set()
	if err == nil {
		ret = attrListGetIteratorFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrIterator{native: ret.Pointer()}

	return retGo, err
}

var attrListInsertFunction *gi.Function
var attrListInsertFunction_Once sync.Once

func attrListInsertFunction_Set() error {
	var err error
	attrListInsertFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListInsertFunction, err = attrListStruct.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type pango_attr_list_insert.
func (recv *AttrList) Insert(attr *Attribute) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr.native)

	err := attrListInsertFunction_Set()
	if err == nil {
		attrListInsertFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var attrListInsertBeforeFunction *gi.Function
var attrListInsertBeforeFunction_Once sync.Once

func attrListInsertBeforeFunction_Set() error {
	var err error
	attrListInsertBeforeFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListInsertBeforeFunction, err = attrListStruct.InvokerNew("insert_before")
	})
	return err
}

// InsertBefore is a representation of the C type pango_attr_list_insert_before.
func (recv *AttrList) InsertBefore(attr *Attribute) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr.native)

	err := attrListInsertBeforeFunction_Set()
	if err == nil {
		attrListInsertBeforeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *AttrList) Ref() (*AttrList, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListRefFunction_Set()
	if err == nil {
		ret = attrListRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrList{native: ret.Pointer()}

	return retGo, err
}

var attrListSpliceFunction *gi.Function
var attrListSpliceFunction_Once sync.Once

func attrListSpliceFunction_Set() error {
	var err error
	attrListSpliceFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListSpliceFunction, err = attrListStruct.InvokerNew("splice")
	})
	return err
}

// Splice is a representation of the C type pango_attr_list_splice.
func (recv *AttrList) Splice(other *AttrList, pos int32, len int32) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(other.native)
	inArgs[2].SetInt32(pos)
	inArgs[3].SetInt32(len)

	err := attrListSpliceFunction_Set()
	if err == nil {
		attrListSpliceFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *AttrList) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attrListUnrefFunction_Set()
	if err == nil {
		attrListUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
func (recv *Attribute) Copy() (*Attribute, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attributeCopyFunction_Set()
	if err == nil {
		ret = attributeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Attribute{native: ret.Pointer()}

	return retGo, err
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
func (recv *Attribute) Destroy() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attributeDestroyFunction_Set()
	if err == nil {
		attributeDestroyFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var attributeEqualFunction *gi.Function
var attributeEqualFunction_Once sync.Once

func attributeEqualFunction_Set() error {
	var err error
	attributeEqualFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeEqualFunction, err = attributeStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type pango_attribute_equal.
func (recv *Attribute) Equal(attr2 *Attribute) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr2.native)

	var ret gi.Argument

	err := attributeEqualFunction_Set()
	if err == nil {
		ret = attributeEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var attributeInitFunction *gi.Function
var attributeInitFunction_Once sync.Once

func attributeInitFunction_Set() error {
	var err error
	attributeInitFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeInitFunction, err = attributeStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type pango_attribute_init.
func (recv *Attribute) Init(klass *AttrClass) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(klass.native)

	err := attributeInitFunction_Set()
	if err == nil {
		attributeInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *Color) Copy() (*Color, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorCopyFunction_Set()
	if err == nil {
		ret = colorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Color{native: ret.Pointer()}

	return retGo, err
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
func (recv *Color) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := colorFreeFunction_Set()
	if err == nil {
		colorFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var colorParseFunction *gi.Function
var colorParseFunction_Once sync.Once

func colorParseFunction_Set() error {
	var err error
	colorParseFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorParseFunction, err = colorStruct.InvokerNew("parse")
	})
	return err
}

// Parse is a representation of the C type pango_color_parse.
func (recv *Color) Parse(spec string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(spec)

	var ret gi.Argument

	err := colorParseFunction_Set()
	if err == nil {
		ret = colorParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *Color) ToString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorToStringFunction_Set()
	if err == nil {
		ret = colorToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
func (recv *Coverage) Copy() (*Coverage, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := coverageCopyFunction_Set()
	if err == nil {
		ret = coverageCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Coverage{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'pango_coverage_get' : return type 'CoverageLevel' not supported

var coverageMaxFunction *gi.Function
var coverageMaxFunction_Once sync.Once

func coverageMaxFunction_Set() error {
	var err error
	coverageMaxFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageMaxFunction, err = coverageStruct.InvokerNew("max")
	})
	return err
}

// Max is a representation of the C type pango_coverage_max.
func (recv *Coverage) Max(other *Coverage) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(other.native)

	err := coverageMaxFunction_Set()
	if err == nil {
		coverageMaxFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *Coverage) Ref() (*Coverage, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := coverageRefFunction_Set()
	if err == nil {
		ret = coverageRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Coverage{native: ret.Pointer()}

	return retGo, err
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
func (recv *Coverage) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := coverageUnrefFunction_Set()
	if err == nil {
		coverageUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
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
	native uintptr
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
func FontDescriptionNew() (*FontDescription, error) {

	var ret gi.Argument

	err := fontDescriptionNewFunction_Set()
	if err == nil {
		ret = fontDescriptionNewFunction.Invoke(nil, nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo, err
}

var fontDescriptionBetterMatchFunction *gi.Function
var fontDescriptionBetterMatchFunction_Once sync.Once

func fontDescriptionBetterMatchFunction_Set() error {
	var err error
	fontDescriptionBetterMatchFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionBetterMatchFunction, err = fontDescriptionStruct.InvokerNew("better_match")
	})
	return err
}

// BetterMatch is a representation of the C type pango_font_description_better_match.
func (recv *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(oldMatch.native)
	inArgs[2].SetPointer(newMatch.native)

	var ret gi.Argument

	err := fontDescriptionBetterMatchFunction_Set()
	if err == nil {
		ret = fontDescriptionBetterMatchFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *FontDescription) Copy() (*FontDescription, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionCopyFunction_Set()
	if err == nil {
		ret = fontDescriptionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo, err
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
func (recv *FontDescription) CopyStatic() (*FontDescription, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionCopyStaticFunction_Set()
	if err == nil {
		ret = fontDescriptionCopyStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo, err
}

var fontDescriptionEqualFunction *gi.Function
var fontDescriptionEqualFunction_Once sync.Once

func fontDescriptionEqualFunction_Set() error {
	var err error
	fontDescriptionEqualFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionEqualFunction, err = fontDescriptionStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type pango_font_description_equal.
func (recv *FontDescription) Equal(desc2 *FontDescription) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(desc2.native)

	var ret gi.Argument

	err := fontDescriptionEqualFunction_Set()
	if err == nil {
		ret = fontDescriptionEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *FontDescription) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontDescriptionFreeFunction_Set()
	if err == nil {
		fontDescriptionFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *FontDescription) GetFamily() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetFamilyFunction_Set()
	if err == nil {
		ret = fontDescriptionGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
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
func (recv *FontDescription) GetSize() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetSizeFunction_Set()
	if err == nil {
		ret = fontDescriptionGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var fontDescriptionGetSizeIsAbsoluteFunction *gi.Function
var fontDescriptionGetSizeIsAbsoluteFunction_Once sync.Once

func fontDescriptionGetSizeIsAbsoluteFunction_Set() error {
	var err error
	fontDescriptionGetSizeIsAbsoluteFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetSizeIsAbsoluteFunction, err = fontDescriptionStruct.InvokerNew("get_size_is_absolute")
	})
	return err
}

// GetSizeIsAbsolute is a representation of the C type pango_font_description_get_size_is_absolute.
func (recv *FontDescription) GetSizeIsAbsolute() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetSizeIsAbsoluteFunction_Set()
	if err == nil {
		ret = fontDescriptionGetSizeIsAbsoluteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *FontDescription) GetVariations() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetVariationsFunction_Set()
	if err == nil {
		ret = fontDescriptionGetVariationsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
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
func (recv *FontDescription) Hash() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionHashFunction_Set()
	if err == nil {
		ret = fontDescriptionHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var fontDescriptionMergeFunction *gi.Function
var fontDescriptionMergeFunction_Once sync.Once

func fontDescriptionMergeFunction_Set() error {
	var err error
	fontDescriptionMergeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionMergeFunction, err = fontDescriptionStruct.InvokerNew("merge")
	})
	return err
}

// Merge is a representation of the C type pango_font_description_merge.
func (recv *FontDescription) Merge(descToMerge *FontDescription, replaceExisting bool) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(descToMerge.native)
	inArgs[2].SetBoolean(replaceExisting)

	err := fontDescriptionMergeFunction_Set()
	if err == nil {
		fontDescriptionMergeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var fontDescriptionMergeStaticFunction *gi.Function
var fontDescriptionMergeStaticFunction_Once sync.Once

func fontDescriptionMergeStaticFunction_Set() error {
	var err error
	fontDescriptionMergeStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionMergeStaticFunction, err = fontDescriptionStruct.InvokerNew("merge_static")
	})
	return err
}

// MergeStatic is a representation of the C type pango_font_description_merge_static.
func (recv *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting bool) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(descToMerge.native)
	inArgs[2].SetBoolean(replaceExisting)

	err := fontDescriptionMergeStaticFunction_Set()
	if err == nil {
		fontDescriptionMergeStaticFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var fontDescriptionSetAbsoluteSizeFunction *gi.Function
var fontDescriptionSetAbsoluteSizeFunction_Once sync.Once

func fontDescriptionSetAbsoluteSizeFunction_Set() error {
	var err error
	fontDescriptionSetAbsoluteSizeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetAbsoluteSizeFunction, err = fontDescriptionStruct.InvokerNew("set_absolute_size")
	})
	return err
}

// SetAbsoluteSize is a representation of the C type pango_font_description_set_absolute_size.
func (recv *FontDescription) SetAbsoluteSize(size float64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(size)

	err := fontDescriptionSetAbsoluteSizeFunction_Set()
	if err == nil {
		fontDescriptionSetAbsoluteSizeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *FontDescription) SetFamily(family string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	err := fontDescriptionSetFamilyFunction_Set()
	if err == nil {
		fontDescriptionSetFamilyFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *FontDescription) SetFamilyStatic(family string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	err := fontDescriptionSetFamilyStaticFunction_Set()
	if err == nil {
		fontDescriptionSetFamilyStaticFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *FontDescription) SetSize(size int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(size)

	err := fontDescriptionSetSizeFunction_Set()
	if err == nil {
		fontDescriptionSetSizeFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *FontDescription) SetVariations(settings string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	err := fontDescriptionSetVariationsFunction_Set()
	if err == nil {
		fontDescriptionSetVariationsFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *FontDescription) SetVariationsStatic(settings string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	err := fontDescriptionSetVariationsStaticFunction_Set()
	if err == nil {
		fontDescriptionSetVariationsStaticFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *FontDescription) ToFilename() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionToFilenameFunction_Set()
	if err == nil {
		ret = fontDescriptionToFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
func (recv *FontDescription) ToString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionToStringFunction_Set()
	if err == nil {
		ret = fontDescriptionToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
func FontMetricsNew() (*FontMetrics, error) {

	var ret gi.Argument

	err := fontMetricsNewFunction_Set()
	if err == nil {
		ret = fontMetricsNewFunction.Invoke(nil, nil)
	}

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo, err
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
func (recv *FontMetrics) GetApproximateCharWidth() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetApproximateCharWidthFunction_Set()
	if err == nil {
		ret = fontMetricsGetApproximateCharWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) GetApproximateDigitWidth() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetApproximateDigitWidthFunction_Set()
	if err == nil {
		ret = fontMetricsGetApproximateDigitWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) GetAscent() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetAscentFunction_Set()
	if err == nil {
		ret = fontMetricsGetAscentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) GetDescent() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetDescentFunction_Set()
	if err == nil {
		ret = fontMetricsGetDescentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) GetStrikethroughPosition() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetStrikethroughPositionFunction_Set()
	if err == nil {
		ret = fontMetricsGetStrikethroughPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) GetStrikethroughThickness() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetStrikethroughThicknessFunction_Set()
	if err == nil {
		ret = fontMetricsGetStrikethroughThicknessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) GetUnderlinePosition() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetUnderlinePositionFunction_Set()
	if err == nil {
		ret = fontMetricsGetUnderlinePositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) GetUnderlineThickness() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetUnderlineThicknessFunction_Set()
	if err == nil {
		ret = fontMetricsGetUnderlineThicknessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *FontMetrics) Ref() (*FontMetrics, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsRefFunction_Set()
	if err == nil {
		ret = fontMetricsRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo, err
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
func (recv *FontMetrics) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontMetricsUnrefFunction_Set()
	if err == nil {
		fontMetricsUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
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
	native uintptr
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
	native uintptr
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
}

// UNSUPPORTED : C value 'pango_glyph_item_apply_attrs' : return type 'GLib.SList' not supported

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
func (recv *GlyphItem) Copy() (*GlyphItem, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemCopyFunction_Set()
	if err == nil {
		ret = glyphItemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GlyphItem{native: ret.Pointer()}

	return retGo, err
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
func (recv *GlyphItem) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphItemFreeFunction_Set()
	if err == nil {
		glyphItemFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *GlyphItem) Split(text string, splitIndex int32) (*GlyphItem, error) {
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

	return retGo, err
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
	native uintptr
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
func (recv *GlyphItemIter) Copy() (*GlyphItemIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemIterCopyFunction_Set()
	if err == nil {
		ret = glyphItemIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GlyphItemIter{native: ret.Pointer()}

	return retGo, err
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
func (recv *GlyphItemIter) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphItemIterFreeFunction_Set()
	if err == nil {
		glyphItemIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var glyphItemIterInitEndFunction *gi.Function
var glyphItemIterInitEndFunction_Once sync.Once

func glyphItemIterInitEndFunction_Set() error {
	var err error
	glyphItemIterInitEndFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterInitEndFunction, err = glyphItemIterStruct.InvokerNew("init_end")
	})
	return err
}

// InitEnd is a representation of the C type pango_glyph_item_iter_init_end.
func (recv *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(glyphItem.native)
	inArgs[2].SetString(text)

	var ret gi.Argument

	err := glyphItemIterInitEndFunction_Set()
	if err == nil {
		ret = glyphItemIterInitEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var glyphItemIterInitStartFunction *gi.Function
var glyphItemIterInitStartFunction_Once sync.Once

func glyphItemIterInitStartFunction_Set() error {
	var err error
	glyphItemIterInitStartFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterInitStartFunction, err = glyphItemIterStruct.InvokerNew("init_start")
	})
	return err
}

// InitStart is a representation of the C type pango_glyph_item_iter_init_start.
func (recv *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(glyphItem.native)
	inArgs[2].SetString(text)

	var ret gi.Argument

	err := glyphItemIterInitStartFunction_Set()
	if err == nil {
		ret = glyphItemIterInitStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var glyphItemIterNextClusterFunction *gi.Function
var glyphItemIterNextClusterFunction_Once sync.Once

func glyphItemIterNextClusterFunction_Set() error {
	var err error
	glyphItemIterNextClusterFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterNextClusterFunction, err = glyphItemIterStruct.InvokerNew("next_cluster")
	})
	return err
}

// NextCluster is a representation of the C type pango_glyph_item_iter_next_cluster.
func (recv *GlyphItemIter) NextCluster() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemIterNextClusterFunction_Set()
	if err == nil {
		ret = glyphItemIterNextClusterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var glyphItemIterPrevClusterFunction *gi.Function
var glyphItemIterPrevClusterFunction_Once sync.Once

func glyphItemIterPrevClusterFunction_Set() error {
	var err error
	glyphItemIterPrevClusterFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterPrevClusterFunction, err = glyphItemIterStruct.InvokerNew("prev_cluster")
	})
	return err
}

// PrevCluster is a representation of the C type pango_glyph_item_iter_prev_cluster.
func (recv *GlyphItemIter) PrevCluster() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemIterPrevClusterFunction_Set()
	if err == nil {
		ret = glyphItemIterPrevClusterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
	native uintptr
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
func GlyphStringNew() (*GlyphString, error) {

	var ret gi.Argument

	err := glyphStringNewFunction_Set()
	if err == nil {
		ret = glyphStringNewFunction.Invoke(nil, nil)
	}

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo, err
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
func (recv *GlyphString) Copy() (*GlyphString, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphStringCopyFunction_Set()
	if err == nil {
		ret = glyphStringCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GlyphString{native: ret.Pointer()}

	return retGo, err
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
func (recv *GlyphString) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphStringFreeFunction_Set()
	if err == nil {
		glyphStringFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *GlyphString) GetWidth() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphStringGetWidthFunction_Set()
	if err == nil {
		ret = glyphStringGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var glyphStringIndexToXFunction *gi.Function
var glyphStringIndexToXFunction_Once sync.Once

func glyphStringIndexToXFunction_Set() error {
	var err error
	glyphStringIndexToXFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringIndexToXFunction, err = glyphStringStruct.InvokerNew("index_to_x")
	})
	return err
}

// IndexToX is a representation of the C type pango_glyph_string_index_to_x.
func (recv *GlyphString) IndexToX(text string, length int32, analysis *Analysis, index int32, trailing bool) (int32, error) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(length)
	inArgs[3].SetPointer(analysis.native)
	inArgs[4].SetInt32(index)
	inArgs[5].SetBoolean(trailing)

	var outArgs [1]gi.Argument

	err := glyphStringIndexToXFunction_Set()
	if err == nil {
		glyphStringIndexToXFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0, err
}

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
func (recv *GlyphString) SetSize(newLen int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newLen)

	err := glyphStringSetSizeFunction_Set()
	if err == nil {
		glyphStringSetSizeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var glyphStringXToIndexFunction *gi.Function
var glyphStringXToIndexFunction_Once sync.Once

func glyphStringXToIndexFunction_Set() error {
	var err error
	glyphStringXToIndexFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringXToIndexFunction, err = glyphStringStruct.InvokerNew("x_to_index")
	})
	return err
}

// XToIndex is a representation of the C type pango_glyph_string_x_to_index.
func (recv *GlyphString) XToIndex(text string, length int32, analysis *Analysis, xPos int32) (int32, int32, error) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(length)
	inArgs[3].SetPointer(analysis.native)
	inArgs[4].SetInt32(xPos)

	var outArgs [2]gi.Argument

	err := glyphStringXToIndexFunction_Set()
	if err == nil {
		glyphStringXToIndexFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1, err
}

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
	native uintptr
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
	native uintptr
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
func ItemNew() (*Item, error) {

	var ret gi.Argument

	err := itemNewFunction_Set()
	if err == nil {
		ret = itemNewFunction.Invoke(nil, nil)
	}

	retGo := &Item{native: ret.Pointer()}

	return retGo, err
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
func (recv *Item) Copy() (*Item, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := itemCopyFunction_Set()
	if err == nil {
		ret = itemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Item{native: ret.Pointer()}

	return retGo, err
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
func (recv *Item) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := itemFreeFunction_Set()
	if err == nil {
		itemFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *Item) Split(splitIndex int32, splitOffset int32) (*Item, error) {
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

	return retGo, err
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
func (recv *Language) GetSampleString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := languageGetSampleStringFunction_Set()
	if err == nil {
		ret = languageGetSampleStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
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
func (recv *Language) GetScripts() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := languageGetScriptsFunction_Set()
	if err == nil {
		languageGetScriptsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0, err
}

// UNSUPPORTED : C value 'pango_language_includes_script' : parameter 'script' of type 'Script' not supported

var languageMatchesFunction *gi.Function
var languageMatchesFunction_Once sync.Once

func languageMatchesFunction_Set() error {
	var err error
	languageMatchesFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageMatchesFunction, err = languageStruct.InvokerNew("matches")
	})
	return err
}

// Matches is a representation of the C type pango_language_matches.
func (recv *Language) Matches(rangeList string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(rangeList)

	var ret gi.Argument

	err := languageMatchesFunction_Set()
	if err == nil {
		ret = languageMatchesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *Language) ToString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := languageToStringFunction_Set()
	if err == nil {
		ret = languageToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
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

var layoutIterAtLastLineFunction *gi.Function
var layoutIterAtLastLineFunction_Once sync.Once

func layoutIterAtLastLineFunction_Set() error {
	var err error
	layoutIterAtLastLineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterAtLastLineFunction, err = layoutIterStruct.InvokerNew("at_last_line")
	})
	return err
}

// AtLastLine is a representation of the C type pango_layout_iter_at_last_line.
func (recv *LayoutIter) AtLastLine() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterAtLastLineFunction_Set()
	if err == nil {
		ret = layoutIterAtLastLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *LayoutIter) Copy() (*LayoutIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterCopyFunction_Set()
	if err == nil {
		ret = layoutIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutIter{native: ret.Pointer()}

	return retGo, err
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
func (recv *LayoutIter) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := layoutIterFreeFunction_Set()
	if err == nil {
		layoutIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
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
func (recv *LayoutIter) GetBaseline() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetBaselineFunction_Set()
	if err == nil {
		ret = layoutIterGetBaselineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var layoutIterGetCharExtentsFunction *gi.Function
var layoutIterGetCharExtentsFunction_Once sync.Once

func layoutIterGetCharExtentsFunction_Set() error {
	var err error
	layoutIterGetCharExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetCharExtentsFunction, err = layoutIterStruct.InvokerNew("get_char_extents")
	})
	return err
}

// GetCharExtents is a representation of the C type pango_layout_iter_get_char_extents.
func (recv *LayoutIter) GetCharExtents() (*Rectangle, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := layoutIterGetCharExtentsFunction_Set()
	if err == nil {
		layoutIterGetCharExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}

	return out0, err
}

var layoutIterGetClusterExtentsFunction *gi.Function
var layoutIterGetClusterExtentsFunction_Once sync.Once

func layoutIterGetClusterExtentsFunction_Set() error {
	var err error
	layoutIterGetClusterExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetClusterExtentsFunction, err = layoutIterStruct.InvokerNew("get_cluster_extents")
	})
	return err
}

// GetClusterExtents is a representation of the C type pango_layout_iter_get_cluster_extents.
func (recv *LayoutIter) GetClusterExtents() (*Rectangle, *Rectangle, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetClusterExtentsFunction_Set()
	if err == nil {
		layoutIterGetClusterExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1, err
}

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
func (recv *LayoutIter) GetIndex() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetIndexFunction_Set()
	if err == nil {
		ret = layoutIterGetIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'pango_layout_iter_get_layout' : return type 'Layout' not supported

var layoutIterGetLayoutExtentsFunction *gi.Function
var layoutIterGetLayoutExtentsFunction_Once sync.Once

func layoutIterGetLayoutExtentsFunction_Set() error {
	var err error
	layoutIterGetLayoutExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLayoutExtentsFunction, err = layoutIterStruct.InvokerNew("get_layout_extents")
	})
	return err
}

// GetLayoutExtents is a representation of the C type pango_layout_iter_get_layout_extents.
func (recv *LayoutIter) GetLayoutExtents() (*Rectangle, *Rectangle, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetLayoutExtentsFunction_Set()
	if err == nil {
		layoutIterGetLayoutExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1, err
}

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
func (recv *LayoutIter) GetLine() (*LayoutLine, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetLineFunction_Set()
	if err == nil {
		ret = layoutIterGetLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo, err
}

var layoutIterGetLineExtentsFunction *gi.Function
var layoutIterGetLineExtentsFunction_Once sync.Once

func layoutIterGetLineExtentsFunction_Set() error {
	var err error
	layoutIterGetLineExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineExtentsFunction, err = layoutIterStruct.InvokerNew("get_line_extents")
	})
	return err
}

// GetLineExtents is a representation of the C type pango_layout_iter_get_line_extents.
func (recv *LayoutIter) GetLineExtents() (*Rectangle, *Rectangle, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetLineExtentsFunction_Set()
	if err == nil {
		layoutIterGetLineExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1, err
}

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
func (recv *LayoutIter) GetLineReadonly() (*LayoutLine, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetLineReadonlyFunction_Set()
	if err == nil {
		ret = layoutIterGetLineReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo, err
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
func (recv *LayoutIter) GetLineYrange() (int32, int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetLineYrangeFunction_Set()
	if err == nil {
		layoutIterGetLineYrangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1, err
}

// UNSUPPORTED : C value 'pango_layout_iter_get_run' : return type 'LayoutRun' not supported

var layoutIterGetRunExtentsFunction *gi.Function
var layoutIterGetRunExtentsFunction_Once sync.Once

func layoutIterGetRunExtentsFunction_Set() error {
	var err error
	layoutIterGetRunExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetRunExtentsFunction, err = layoutIterStruct.InvokerNew("get_run_extents")
	})
	return err
}

// GetRunExtents is a representation of the C type pango_layout_iter_get_run_extents.
func (recv *LayoutIter) GetRunExtents() (*Rectangle, *Rectangle, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetRunExtentsFunction_Set()
	if err == nil {
		layoutIterGetRunExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1, err
}

// UNSUPPORTED : C value 'pango_layout_iter_get_run_readonly' : return type 'LayoutRun' not supported

var layoutIterNextCharFunction *gi.Function
var layoutIterNextCharFunction_Once sync.Once

func layoutIterNextCharFunction_Set() error {
	var err error
	layoutIterNextCharFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextCharFunction, err = layoutIterStruct.InvokerNew("next_char")
	})
	return err
}

// NextChar is a representation of the C type pango_layout_iter_next_char.
func (recv *LayoutIter) NextChar() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextCharFunction_Set()
	if err == nil {
		ret = layoutIterNextCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var layoutIterNextClusterFunction *gi.Function
var layoutIterNextClusterFunction_Once sync.Once

func layoutIterNextClusterFunction_Set() error {
	var err error
	layoutIterNextClusterFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextClusterFunction, err = layoutIterStruct.InvokerNew("next_cluster")
	})
	return err
}

// NextCluster is a representation of the C type pango_layout_iter_next_cluster.
func (recv *LayoutIter) NextCluster() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextClusterFunction_Set()
	if err == nil {
		ret = layoutIterNextClusterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var layoutIterNextLineFunction *gi.Function
var layoutIterNextLineFunction_Once sync.Once

func layoutIterNextLineFunction_Set() error {
	var err error
	layoutIterNextLineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextLineFunction, err = layoutIterStruct.InvokerNew("next_line")
	})
	return err
}

// NextLine is a representation of the C type pango_layout_iter_next_line.
func (recv *LayoutIter) NextLine() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextLineFunction_Set()
	if err == nil {
		ret = layoutIterNextLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var layoutIterNextRunFunction *gi.Function
var layoutIterNextRunFunction_Once sync.Once

func layoutIterNextRunFunction_Set() error {
	var err error
	layoutIterNextRunFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextRunFunction, err = layoutIterStruct.InvokerNew("next_run")
	})
	return err
}

// NextRun is a representation of the C type pango_layout_iter_next_run.
func (recv *LayoutIter) NextRun() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextRunFunction_Set()
	if err == nil {
		ret = layoutIterNextRunFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
}

var layoutLineGetExtentsFunction *gi.Function
var layoutLineGetExtentsFunction_Once sync.Once

func layoutLineGetExtentsFunction_Set() error {
	var err error
	layoutLineGetExtentsFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineGetExtentsFunction, err = layoutLineStruct.InvokerNew("get_extents")
	})
	return err
}

// GetExtents is a representation of the C type pango_layout_line_get_extents.
func (recv *LayoutLine) GetExtents() (*Rectangle, *Rectangle, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutLineGetExtentsFunction_Set()
	if err == nil {
		layoutLineGetExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1, err
}

var layoutLineGetPixelExtentsFunction *gi.Function
var layoutLineGetPixelExtentsFunction_Once sync.Once

func layoutLineGetPixelExtentsFunction_Set() error {
	var err error
	layoutLineGetPixelExtentsFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineGetPixelExtentsFunction, err = layoutLineStruct.InvokerNew("get_pixel_extents")
	})
	return err
}

// GetPixelExtents is a representation of the C type pango_layout_line_get_pixel_extents.
func (recv *LayoutLine) GetPixelExtents() (*Rectangle, *Rectangle, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutLineGetPixelExtentsFunction_Set()
	if err == nil {
		layoutLineGetPixelExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1, err
}

// UNSUPPORTED : C value 'pango_layout_line_get_x_ranges' : parameter 'ranges' has no type

var layoutLineIndexToXFunction *gi.Function
var layoutLineIndexToXFunction_Once sync.Once

func layoutLineIndexToXFunction_Set() error {
	var err error
	layoutLineIndexToXFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineIndexToXFunction, err = layoutLineStruct.InvokerNew("index_to_x")
	})
	return err
}

// IndexToX is a representation of the C type pango_layout_line_index_to_x.
func (recv *LayoutLine) IndexToX(index int32, trailing bool) (int32, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)
	inArgs[2].SetBoolean(trailing)

	var outArgs [1]gi.Argument

	err := layoutLineIndexToXFunction_Set()
	if err == nil {
		layoutLineIndexToXFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0, err
}

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
func (recv *LayoutLine) Ref() (*LayoutLine, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutLineRefFunction_Set()
	if err == nil {
		ret = layoutLineRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo, err
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
func (recv *LayoutLine) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := layoutLineUnrefFunction_Set()
	if err == nil {
		layoutLineUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var layoutLineXToIndexFunction *gi.Function
var layoutLineXToIndexFunction_Once sync.Once

func layoutLineXToIndexFunction_Set() error {
	var err error
	layoutLineXToIndexFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineXToIndexFunction, err = layoutLineStruct.InvokerNew("x_to_index")
	})
	return err
}

// XToIndex is a representation of the C type pango_layout_line_x_to_index.
func (recv *LayoutLine) XToIndex(xPos int32) (bool, int32, int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(xPos)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := layoutLineXToIndexFunction_Set()
	if err == nil {
		ret = layoutLineXToIndexFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1, err
}

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
	native uintptr
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
}

var matrixConcatFunction *gi.Function
var matrixConcatFunction_Once sync.Once

func matrixConcatFunction_Set() error {
	var err error
	matrixConcatFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixConcatFunction, err = matrixStruct.InvokerNew("concat")
	})
	return err
}

// Concat is a representation of the C type pango_matrix_concat.
func (recv *Matrix) Concat(newMatrix *Matrix) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(newMatrix.native)

	err := matrixConcatFunction_Set()
	if err == nil {
		matrixConcatFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *Matrix) Copy() (*Matrix, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matrixCopyFunction_Set()
	if err == nil {
		ret = matrixCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Matrix{native: ret.Pointer()}

	return retGo, err
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
func (recv *Matrix) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := matrixFreeFunction_Set()
	if err == nil {
		matrixFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var matrixGetFontScaleFactorFunction *gi.Function
var matrixGetFontScaleFactorFunction_Once sync.Once

func matrixGetFontScaleFactorFunction_Set() error {
	var err error
	matrixGetFontScaleFactorFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixGetFontScaleFactorFunction, err = matrixStruct.InvokerNew("get_font_scale_factor")
	})
	return err
}

// GetFontScaleFactor is a representation of the C type pango_matrix_get_font_scale_factor.
func (recv *Matrix) GetFontScaleFactor() (float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matrixGetFontScaleFactorFunction_Set()
	if err == nil {
		ret = matrixGetFontScaleFactorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var matrixGetFontScaleFactorsFunction *gi.Function
var matrixGetFontScaleFactorsFunction_Once sync.Once

func matrixGetFontScaleFactorsFunction_Set() error {
	var err error
	matrixGetFontScaleFactorsFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixGetFontScaleFactorsFunction, err = matrixStruct.InvokerNew("get_font_scale_factors")
	})
	return err
}

// GetFontScaleFactors is a representation of the C type pango_matrix_get_font_scale_factors.
func (recv *Matrix) GetFontScaleFactors() (float64, float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := matrixGetFontScaleFactorsFunction_Set()
	if err == nil {
		matrixGetFontScaleFactorsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1, err
}

var matrixRotateFunction *gi.Function
var matrixRotateFunction_Once sync.Once

func matrixRotateFunction_Set() error {
	var err error
	matrixRotateFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixRotateFunction, err = matrixStruct.InvokerNew("rotate")
	})
	return err
}

// Rotate is a representation of the C type pango_matrix_rotate.
func (recv *Matrix) Rotate(degrees float64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(degrees)

	err := matrixRotateFunction_Set()
	if err == nil {
		matrixRotateFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var matrixScaleFunction *gi.Function
var matrixScaleFunction_Once sync.Once

func matrixScaleFunction_Set() error {
	var err error
	matrixScaleFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixScaleFunction, err = matrixStruct.InvokerNew("scale")
	})
	return err
}

// Scale is a representation of the C type pango_matrix_scale.
func (recv *Matrix) Scale(scaleX float64, scaleY float64) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(scaleX)
	inArgs[2].SetFloat64(scaleY)

	err := matrixScaleFunction_Set()
	if err == nil {
		matrixScaleFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var matrixTransformDistanceFunction *gi.Function
var matrixTransformDistanceFunction_Once sync.Once

func matrixTransformDistanceFunction_Set() error {
	var err error
	matrixTransformDistanceFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformDistanceFunction, err = matrixStruct.InvokerNew("transform_distance")
	})
	return err
}

// TransformDistance is a representation of the C type pango_matrix_transform_distance.
func (recv *Matrix) TransformDistance(dx float64, dy float64) (float64, float64, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(dx)
	inArgs[2].SetFloat64(dy)

	var outArgs [2]gi.Argument

	err := matrixTransformDistanceFunction_Set()
	if err == nil {
		matrixTransformDistanceFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1, err
}

var matrixTransformPixelRectangleFunction *gi.Function
var matrixTransformPixelRectangleFunction_Once sync.Once

func matrixTransformPixelRectangleFunction_Set() error {
	var err error
	matrixTransformPixelRectangleFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformPixelRectangleFunction, err = matrixStruct.InvokerNew("transform_pixel_rectangle")
	})
	return err
}

// TransformPixelRectangle is a representation of the C type pango_matrix_transform_pixel_rectangle.
func (recv *Matrix) TransformPixelRectangle(rect *Rectangle) (*Rectangle, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(rect.native)

	var outArgs [1]gi.Argument

	err := matrixTransformPixelRectangleFunction_Set()
	if err == nil {
		matrixTransformPixelRectangleFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}

	return out0, err
}

var matrixTransformPointFunction *gi.Function
var matrixTransformPointFunction_Once sync.Once

func matrixTransformPointFunction_Set() error {
	var err error
	matrixTransformPointFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformPointFunction, err = matrixStruct.InvokerNew("transform_point")
	})
	return err
}

// TransformPoint is a representation of the C type pango_matrix_transform_point.
func (recv *Matrix) TransformPoint(x float64, y float64) (float64, float64, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(x)
	inArgs[2].SetFloat64(y)

	var outArgs [2]gi.Argument

	err := matrixTransformPointFunction_Set()
	if err == nil {
		matrixTransformPointFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1, err
}

var matrixTransformRectangleFunction *gi.Function
var matrixTransformRectangleFunction_Once sync.Once

func matrixTransformRectangleFunction_Set() error {
	var err error
	matrixTransformRectangleFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformRectangleFunction, err = matrixStruct.InvokerNew("transform_rectangle")
	})
	return err
}

// TransformRectangle is a representation of the C type pango_matrix_transform_rectangle.
func (recv *Matrix) TransformRectangle(rect *Rectangle) (*Rectangle, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(rect.native)

	var outArgs [1]gi.Argument

	err := matrixTransformRectangleFunction_Set()
	if err == nil {
		matrixTransformRectangleFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}

	return out0, err
}

var matrixTranslateFunction *gi.Function
var matrixTranslateFunction_Once sync.Once

func matrixTranslateFunction_Set() error {
	var err error
	matrixTranslateFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTranslateFunction, err = matrixStruct.InvokerNew("translate")
	})
	return err
}

// Translate is a representation of the C type pango_matrix_translate.
func (recv *Matrix) Translate(tx float64, ty float64) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(tx)
	inArgs[2].SetFloat64(ty)

	err := matrixTranslateFunction_Set()
	if err == nil {
		matrixTranslateFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
func (recv *ScriptIter) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := scriptIterFreeFunction_Set()
	if err == nil {
		scriptIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'pango_script_iter_get_range' : parameter 'script' of type 'Script' not supported

var scriptIterNextFunction *gi.Function
var scriptIterNextFunction_Once sync.Once

func scriptIterNextFunction_Set() error {
	var err error
	scriptIterNextFunction_Once.Do(func() {
		err = scriptIterStruct_Set()
		if err != nil {
			return
		}
		scriptIterNextFunction, err = scriptIterStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type pango_script_iter_next.
func (recv *ScriptIter) Next() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scriptIterNextFunction_Set()
	if err == nil {
		ret = scriptIterNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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

var tabArrayNewFunction *gi.Function
var tabArrayNewFunction_Once sync.Once

func tabArrayNewFunction_Set() error {
	var err error
	tabArrayNewFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayNewFunction, err = tabArrayStruct.InvokerNew("new")
	})
	return err
}

// TabArrayNew is a representation of the C type pango_tab_array_new.
func TabArrayNew(initialSize int32, positionsInPixels bool) (*TabArray, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(initialSize)
	inArgs[1].SetBoolean(positionsInPixels)

	var ret gi.Argument

	err := tabArrayNewFunction_Set()
	if err == nil {
		ret = tabArrayNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TabArray{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'pango_tab_array_new_with_positions' : parameter 'first_alignment' of type 'TabAlign' not supported

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
func (recv *TabArray) Copy() (*TabArray, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayCopyFunction_Set()
	if err == nil {
		ret = tabArrayCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TabArray{native: ret.Pointer()}

	return retGo, err
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
func (recv *TabArray) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := tabArrayFreeFunction_Set()
	if err == nil {
		tabArrayFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var tabArrayGetPositionsInPixelsFunction *gi.Function
var tabArrayGetPositionsInPixelsFunction_Once sync.Once

func tabArrayGetPositionsInPixelsFunction_Set() error {
	var err error
	tabArrayGetPositionsInPixelsFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayGetPositionsInPixelsFunction, err = tabArrayStruct.InvokerNew("get_positions_in_pixels")
	})
	return err
}

// GetPositionsInPixels is a representation of the C type pango_tab_array_get_positions_in_pixels.
func (recv *TabArray) GetPositionsInPixels() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayGetPositionsInPixelsFunction_Set()
	if err == nil {
		ret = tabArrayGetPositionsInPixelsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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
func (recv *TabArray) GetSize() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayGetSizeFunction_Set()
	if err == nil {
		ret = tabArrayGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
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
func (recv *TabArray) Resize(newSize int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newSize)

	err := tabArrayResizeFunction_Set()
	if err == nil {
		tabArrayResizeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'pango_tab_array_set_tab' : parameter 'alignment' of type 'TabAlign' not supported

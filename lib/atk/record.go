// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var actionIfaceStruct *gi.Struct
var actionIfaceStruct_Once sync.Once

func actionIfaceStruct_Set() error {
	var err error
	actionIfaceStruct_Once.Do(func() {
		actionIfaceStruct, err = gi.StructNew("Atk", "ActionIface")
	})
	return err
}

type ActionIface struct {
	native uintptr
}

var attributeStruct *gi.Struct
var attributeStruct_Once sync.Once

func attributeStruct_Set() error {
	var err error
	attributeStruct_Once.Do(func() {
		attributeStruct, err = gi.StructNew("Atk", "Attribute")
	})
	return err
}

type Attribute struct {
	native uintptr
}

var componentIfaceStruct *gi.Struct
var componentIfaceStruct_Once sync.Once

func componentIfaceStruct_Set() error {
	var err error
	componentIfaceStruct_Once.Do(func() {
		componentIfaceStruct, err = gi.StructNew("Atk", "ComponentIface")
	})
	return err
}

type ComponentIface struct {
	native uintptr
}

var documentIfaceStruct *gi.Struct
var documentIfaceStruct_Once sync.Once

func documentIfaceStruct_Set() error {
	var err error
	documentIfaceStruct_Once.Do(func() {
		documentIfaceStruct, err = gi.StructNew("Atk", "DocumentIface")
	})
	return err
}

type DocumentIface struct {
	native uintptr
}

var editableTextIfaceStruct *gi.Struct
var editableTextIfaceStruct_Once sync.Once

func editableTextIfaceStruct_Set() error {
	var err error
	editableTextIfaceStruct_Once.Do(func() {
		editableTextIfaceStruct, err = gi.StructNew("Atk", "EditableTextIface")
	})
	return err
}

type EditableTextIface struct {
	native uintptr
}

var gObjectAccessibleClassStruct *gi.Struct
var gObjectAccessibleClassStruct_Once sync.Once

func gObjectAccessibleClassStruct_Set() error {
	var err error
	gObjectAccessibleClassStruct_Once.Do(func() {
		gObjectAccessibleClassStruct, err = gi.StructNew("Atk", "GObjectAccessibleClass")
	})
	return err
}

type GObjectAccessibleClass struct {
	native uintptr
}

var hyperlinkClassStruct *gi.Struct
var hyperlinkClassStruct_Once sync.Once

func hyperlinkClassStruct_Set() error {
	var err error
	hyperlinkClassStruct_Once.Do(func() {
		hyperlinkClassStruct, err = gi.StructNew("Atk", "HyperlinkClass")
	})
	return err
}

type HyperlinkClass struct {
	native uintptr
}

var hyperlinkImplIfaceStruct *gi.Struct
var hyperlinkImplIfaceStruct_Once sync.Once

func hyperlinkImplIfaceStruct_Set() error {
	var err error
	hyperlinkImplIfaceStruct_Once.Do(func() {
		hyperlinkImplIfaceStruct, err = gi.StructNew("Atk", "HyperlinkImplIface")
	})
	return err
}

type HyperlinkImplIface struct {
	native uintptr
}

var hypertextIfaceStruct *gi.Struct
var hypertextIfaceStruct_Once sync.Once

func hypertextIfaceStruct_Set() error {
	var err error
	hypertextIfaceStruct_Once.Do(func() {
		hypertextIfaceStruct, err = gi.StructNew("Atk", "HypertextIface")
	})
	return err
}

type HypertextIface struct {
	native uintptr
}

var imageIfaceStruct *gi.Struct
var imageIfaceStruct_Once sync.Once

func imageIfaceStruct_Set() error {
	var err error
	imageIfaceStruct_Once.Do(func() {
		imageIfaceStruct, err = gi.StructNew("Atk", "ImageIface")
	})
	return err
}

type ImageIface struct {
	native uintptr
}

var implementorStruct *gi.Struct
var implementorStruct_Once sync.Once

func implementorStruct_Set() error {
	var err error
	implementorStruct_Once.Do(func() {
		implementorStruct, err = gi.StructNew("Atk", "Implementor")
	})
	return err
}

type Implementor struct {
	native uintptr
}

// UNSUPPORTED : C value 'atk_implementor_ref_accessible' : return type 'Object' not supported

var keyEventStructStruct *gi.Struct
var keyEventStructStruct_Once sync.Once

func keyEventStructStruct_Set() error {
	var err error
	keyEventStructStruct_Once.Do(func() {
		keyEventStructStruct, err = gi.StructNew("Atk", "KeyEventStruct")
	})
	return err
}

type KeyEventStruct struct {
	native uintptr
}

var miscClassStruct *gi.Struct
var miscClassStruct_Once sync.Once

func miscClassStruct_Set() error {
	var err error
	miscClassStruct_Once.Do(func() {
		miscClassStruct, err = gi.StructNew("Atk", "MiscClass")
	})
	return err
}

type MiscClass struct {
	native uintptr
}

var noOpObjectClassStruct *gi.Struct
var noOpObjectClassStruct_Once sync.Once

func noOpObjectClassStruct_Set() error {
	var err error
	noOpObjectClassStruct_Once.Do(func() {
		noOpObjectClassStruct, err = gi.StructNew("Atk", "NoOpObjectClass")
	})
	return err
}

type NoOpObjectClass struct {
	native uintptr
}

var noOpObjectFactoryClassStruct *gi.Struct
var noOpObjectFactoryClassStruct_Once sync.Once

func noOpObjectFactoryClassStruct_Set() error {
	var err error
	noOpObjectFactoryClassStruct_Once.Do(func() {
		noOpObjectFactoryClassStruct, err = gi.StructNew("Atk", "NoOpObjectFactoryClass")
	})
	return err
}

type NoOpObjectFactoryClass struct {
	native uintptr
}

var objectClassStruct *gi.Struct
var objectClassStruct_Once sync.Once

func objectClassStruct_Set() error {
	var err error
	objectClassStruct_Once.Do(func() {
		objectClassStruct, err = gi.StructNew("Atk", "ObjectClass")
	})
	return err
}

type ObjectClass struct {
	native uintptr
}

var objectFactoryClassStruct *gi.Struct
var objectFactoryClassStruct_Once sync.Once

func objectFactoryClassStruct_Set() error {
	var err error
	objectFactoryClassStruct_Once.Do(func() {
		objectFactoryClassStruct, err = gi.StructNew("Atk", "ObjectFactoryClass")
	})
	return err
}

type ObjectFactoryClass struct {
	native uintptr
}

var plugClassStruct *gi.Struct
var plugClassStruct_Once sync.Once

func plugClassStruct_Set() error {
	var err error
	plugClassStruct_Once.Do(func() {
		plugClassStruct, err = gi.StructNew("Atk", "PlugClass")
	})
	return err
}

type PlugClass struct {
	native uintptr
}

var propertyValuesStruct *gi.Struct
var propertyValuesStruct_Once sync.Once

func propertyValuesStruct_Set() error {
	var err error
	propertyValuesStruct_Once.Do(func() {
		propertyValuesStruct, err = gi.StructNew("Atk", "PropertyValues")
	})
	return err
}

type PropertyValues struct {
	native uintptr
}

var rangeStruct *gi.Struct
var rangeStruct_Once sync.Once

func rangeStruct_Set() error {
	var err error
	rangeStruct_Once.Do(func() {
		rangeStruct, err = gi.StructNew("Atk", "Range")
	})
	return err
}

type Range struct {
	native uintptr
}

var rangeNewFunction *gi.Function
var rangeNewFunction_Once sync.Once

func rangeNewFunction_Set() error {
	var err error
	rangeNewFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeNewFunction, err = rangeStruct.InvokerNew("new")
	})
	return err
}

// RangeNew is a representation of the C type atk_range_new.
func RangeNew(lowerLimit float64, upperLimit float64, description string) (*Range, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetFloat64(lowerLimit)
	inArgs[1].SetFloat64(upperLimit)
	inArgs[2].SetString(description)

	var ret gi.Argument

	err := rangeNewFunction_Set()
	if err == nil {
		ret = rangeNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Range{native: ret.Pointer()}

	return retGo, err
}

var rangeCopyFunction *gi.Function
var rangeCopyFunction_Once sync.Once

func rangeCopyFunction_Set() error {
	var err error
	rangeCopyFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeCopyFunction, err = rangeStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type atk_range_copy.
func (recv *Range) Copy() (*Range, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rangeCopyFunction_Set()
	if err == nil {
		ret = rangeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Range{native: ret.Pointer()}

	return retGo, err
}

var rangeFreeFunction *gi.Function
var rangeFreeFunction_Once sync.Once

func rangeFreeFunction_Set() error {
	var err error
	rangeFreeFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeFreeFunction, err = rangeStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type atk_range_free.
func (recv *Range) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rangeFreeFunction_Set()
	if err == nil {
		rangeFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var rangeGetDescriptionFunction *gi.Function
var rangeGetDescriptionFunction_Once sync.Once

func rangeGetDescriptionFunction_Set() error {
	var err error
	rangeGetDescriptionFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeGetDescriptionFunction, err = rangeStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type atk_range_get_description.
func (recv *Range) GetDescription() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rangeGetDescriptionFunction_Set()
	if err == nil {
		ret = rangeGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var rangeGetLowerLimitFunction *gi.Function
var rangeGetLowerLimitFunction_Once sync.Once

func rangeGetLowerLimitFunction_Set() error {
	var err error
	rangeGetLowerLimitFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeGetLowerLimitFunction, err = rangeStruct.InvokerNew("get_lower_limit")
	})
	return err
}

// GetLowerLimit is a representation of the C type atk_range_get_lower_limit.
func (recv *Range) GetLowerLimit() (float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rangeGetLowerLimitFunction_Set()
	if err == nil {
		ret = rangeGetLowerLimitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var rangeGetUpperLimitFunction *gi.Function
var rangeGetUpperLimitFunction_Once sync.Once

func rangeGetUpperLimitFunction_Set() error {
	var err error
	rangeGetUpperLimitFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeGetUpperLimitFunction, err = rangeStruct.InvokerNew("get_upper_limit")
	})
	return err
}

// GetUpperLimit is a representation of the C type atk_range_get_upper_limit.
func (recv *Range) GetUpperLimit() (float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rangeGetUpperLimitFunction_Set()
	if err == nil {
		ret = rangeGetUpperLimitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() error {
	var err error
	rectangleStruct_Once.Do(func() {
		rectangleStruct, err = gi.StructNew("Atk", "Rectangle")
	})
	return err
}

type Rectangle struct {
	native uintptr
}

var registryClassStruct *gi.Struct
var registryClassStruct_Once sync.Once

func registryClassStruct_Set() error {
	var err error
	registryClassStruct_Once.Do(func() {
		registryClassStruct, err = gi.StructNew("Atk", "RegistryClass")
	})
	return err
}

type RegistryClass struct {
	native uintptr
}

var relationClassStruct *gi.Struct
var relationClassStruct_Once sync.Once

func relationClassStruct_Set() error {
	var err error
	relationClassStruct_Once.Do(func() {
		relationClassStruct, err = gi.StructNew("Atk", "RelationClass")
	})
	return err
}

type RelationClass struct {
	native uintptr
}

var relationSetClassStruct *gi.Struct
var relationSetClassStruct_Once sync.Once

func relationSetClassStruct_Set() error {
	var err error
	relationSetClassStruct_Once.Do(func() {
		relationSetClassStruct, err = gi.StructNew("Atk", "RelationSetClass")
	})
	return err
}

type RelationSetClass struct {
	native uintptr
}

var selectionIfaceStruct *gi.Struct
var selectionIfaceStruct_Once sync.Once

func selectionIfaceStruct_Set() error {
	var err error
	selectionIfaceStruct_Once.Do(func() {
		selectionIfaceStruct, err = gi.StructNew("Atk", "SelectionIface")
	})
	return err
}

type SelectionIface struct {
	native uintptr
}

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() error {
	var err error
	socketClassStruct_Once.Do(func() {
		socketClassStruct, err = gi.StructNew("Atk", "SocketClass")
	})
	return err
}

type SocketClass struct {
	native uintptr
}

var stateSetClassStruct *gi.Struct
var stateSetClassStruct_Once sync.Once

func stateSetClassStruct_Set() error {
	var err error
	stateSetClassStruct_Once.Do(func() {
		stateSetClassStruct, err = gi.StructNew("Atk", "StateSetClass")
	})
	return err
}

type StateSetClass struct {
	native uintptr
}

var streamableContentIfaceStruct *gi.Struct
var streamableContentIfaceStruct_Once sync.Once

func streamableContentIfaceStruct_Set() error {
	var err error
	streamableContentIfaceStruct_Once.Do(func() {
		streamableContentIfaceStruct, err = gi.StructNew("Atk", "StreamableContentIface")
	})
	return err
}

type StreamableContentIface struct {
	native uintptr
}

var tableCellIfaceStruct *gi.Struct
var tableCellIfaceStruct_Once sync.Once

func tableCellIfaceStruct_Set() error {
	var err error
	tableCellIfaceStruct_Once.Do(func() {
		tableCellIfaceStruct, err = gi.StructNew("Atk", "TableCellIface")
	})
	return err
}

type TableCellIface struct {
	native uintptr
}

var tableIfaceStruct *gi.Struct
var tableIfaceStruct_Once sync.Once

func tableIfaceStruct_Set() error {
	var err error
	tableIfaceStruct_Once.Do(func() {
		tableIfaceStruct, err = gi.StructNew("Atk", "TableIface")
	})
	return err
}

type TableIface struct {
	native uintptr
}

var textIfaceStruct *gi.Struct
var textIfaceStruct_Once sync.Once

func textIfaceStruct_Set() error {
	var err error
	textIfaceStruct_Once.Do(func() {
		textIfaceStruct, err = gi.StructNew("Atk", "TextIface")
	})
	return err
}

type TextIface struct {
	native uintptr
}

var textRangeStruct *gi.Struct
var textRangeStruct_Once sync.Once

func textRangeStruct_Set() error {
	var err error
	textRangeStruct_Once.Do(func() {
		textRangeStruct, err = gi.StructNew("Atk", "TextRange")
	})
	return err
}

type TextRange struct {
	native uintptr
}

var textRectangleStruct *gi.Struct
var textRectangleStruct_Once sync.Once

func textRectangleStruct_Set() error {
	var err error
	textRectangleStruct_Once.Do(func() {
		textRectangleStruct, err = gi.StructNew("Atk", "TextRectangle")
	})
	return err
}

type TextRectangle struct {
	native uintptr
}

var utilClassStruct *gi.Struct
var utilClassStruct_Once sync.Once

func utilClassStruct_Set() error {
	var err error
	utilClassStruct_Once.Do(func() {
		utilClassStruct, err = gi.StructNew("Atk", "UtilClass")
	})
	return err
}

type UtilClass struct {
	native uintptr
}

var valueIfaceStruct *gi.Struct
var valueIfaceStruct_Once sync.Once

func valueIfaceStruct_Set() error {
	var err error
	valueIfaceStruct_Once.Do(func() {
		valueIfaceStruct, err = gi.StructNew("Atk", "ValueIface")
	})
	return err
}

type ValueIface struct {
	native uintptr
}

var windowIfaceStruct *gi.Struct
var windowIfaceStruct_Once sync.Once

func windowIfaceStruct_Set() error {
	var err error
	windowIfaceStruct_Once.Do(func() {
		windowIfaceStruct, err = gi.StructNew("Atk", "WindowIface")
	})
	return err
}

type WindowIface struct {
	native uintptr
}

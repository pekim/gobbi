// Code generated - DO NOT EDIT.

package pango

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var contextObject *gi.Object
var contextObject_Once sync.Once

func contextObject_Set() error {
	var err error
	contextObject_Once.Do(func() {
		contextObject, err = gi.ObjectNew("Pango", "Context")
	})
	return err
}

type Context struct {
	native unsafe.Pointer
}

func ContextNewFromNative(native unsafe.Pointer) *Context {
	instance := &Context{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Context) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToContext down casts any arbitrary Object to Context.
Exercise care, as this is a potentially dangerous function
if the Object is not a Context.
*/
func (recv *Context) CastToContext(object *gobject.Object) *Context {
	return ContextNewFromNative(object.Native())
}

func (recv *Context) Native() unsafe.Pointer {
	return recv.native
}

var contextNewFunction *gi.Function
var contextNewFunction_Once sync.Once

func contextNewFunction_Set() error {
	var err error
	contextNewFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextNewFunction, err = contextObject.InvokerNew("new")
	})
	return err
}

// ContextNew is a representation of the C type pango_context_new.
func ContextNew() *Context {

	var ret gi.Argument

	err := contextNewFunction_Set()
	if err == nil {
		ret = contextNewFunction.Invoke(nil, nil)
	}

	retGo := ContextNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var contextChangedFunction *gi.Function
var contextChangedFunction_Once sync.Once

func contextChangedFunction_Set() error {
	var err error
	contextChangedFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextChangedFunction, err = contextObject.InvokerNew("changed")
	})
	return err
}

// Changed is a representation of the C type pango_context_changed.
func (recv *Context) Changed() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := contextChangedFunction_Set()
	if err == nil {
		contextChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextGetBaseDirFunction *gi.Function
var contextGetBaseDirFunction_Once sync.Once

func contextGetBaseDirFunction_Set() error {
	var err error
	contextGetBaseDirFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetBaseDirFunction, err = contextObject.InvokerNew("get_base_dir")
	})
	return err
}

// GetBaseDir is a representation of the C type pango_context_get_base_dir.
func (recv *Context) GetBaseDir() Direction {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetBaseDirFunction_Set()
	if err == nil {
		ret = contextGetBaseDirFunction.Invoke(inArgs[:], nil)
	}

	retGo := Direction(ret.Int32())

	return retGo
}

var contextGetBaseGravityFunction *gi.Function
var contextGetBaseGravityFunction_Once sync.Once

func contextGetBaseGravityFunction_Set() error {
	var err error
	contextGetBaseGravityFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetBaseGravityFunction, err = contextObject.InvokerNew("get_base_gravity")
	})
	return err
}

// GetBaseGravity is a representation of the C type pango_context_get_base_gravity.
func (recv *Context) GetBaseGravity() Gravity {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetBaseGravityFunction_Set()
	if err == nil {
		ret = contextGetBaseGravityFunction.Invoke(inArgs[:], nil)
	}

	retGo := Gravity(ret.Int32())

	return retGo
}

var contextGetFontDescriptionFunction *gi.Function
var contextGetFontDescriptionFunction_Once sync.Once

func contextGetFontDescriptionFunction_Set() error {
	var err error
	contextGetFontDescriptionFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetFontDescriptionFunction, err = contextObject.InvokerNew("get_font_description")
	})
	return err
}

// GetFontDescription is a representation of the C type pango_context_get_font_description.
func (recv *Context) GetFontDescription() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetFontDescriptionFunction_Set()
	if err == nil {
		ret = contextGetFontDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var contextGetFontMapFunction *gi.Function
var contextGetFontMapFunction_Once sync.Once

func contextGetFontMapFunction_Set() error {
	var err error
	contextGetFontMapFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetFontMapFunction, err = contextObject.InvokerNew("get_font_map")
	})
	return err
}

// GetFontMap is a representation of the C type pango_context_get_font_map.
func (recv *Context) GetFontMap() *FontMap {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetFontMapFunction_Set()
	if err == nil {
		ret = contextGetFontMapFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontMapNewFromNative(ret.Pointer())

	return retGo
}

var contextGetGravityFunction *gi.Function
var contextGetGravityFunction_Once sync.Once

func contextGetGravityFunction_Set() error {
	var err error
	contextGetGravityFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetGravityFunction, err = contextObject.InvokerNew("get_gravity")
	})
	return err
}

// GetGravity is a representation of the C type pango_context_get_gravity.
func (recv *Context) GetGravity() Gravity {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetGravityFunction_Set()
	if err == nil {
		ret = contextGetGravityFunction.Invoke(inArgs[:], nil)
	}

	retGo := Gravity(ret.Int32())

	return retGo
}

var contextGetGravityHintFunction *gi.Function
var contextGetGravityHintFunction_Once sync.Once

func contextGetGravityHintFunction_Set() error {
	var err error
	contextGetGravityHintFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetGravityHintFunction, err = contextObject.InvokerNew("get_gravity_hint")
	})
	return err
}

// GetGravityHint is a representation of the C type pango_context_get_gravity_hint.
func (recv *Context) GetGravityHint() GravityHint {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetGravityHintFunction_Set()
	if err == nil {
		ret = contextGetGravityHintFunction.Invoke(inArgs[:], nil)
	}

	retGo := GravityHint(ret.Int32())

	return retGo
}

var contextGetLanguageFunction *gi.Function
var contextGetLanguageFunction_Once sync.Once

func contextGetLanguageFunction_Set() error {
	var err error
	contextGetLanguageFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetLanguageFunction, err = contextObject.InvokerNew("get_language")
	})
	return err
}

// GetLanguage is a representation of the C type pango_context_get_language.
func (recv *Context) GetLanguage() *Language {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetLanguageFunction_Set()
	if err == nil {
		ret = contextGetLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := LanguageNewFromNative(ret.Pointer())

	return retGo
}

var contextGetMatrixFunction *gi.Function
var contextGetMatrixFunction_Once sync.Once

func contextGetMatrixFunction_Set() error {
	var err error
	contextGetMatrixFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetMatrixFunction, err = contextObject.InvokerNew("get_matrix")
	})
	return err
}

// GetMatrix is a representation of the C type pango_context_get_matrix.
func (recv *Context) GetMatrix() *Matrix {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetMatrixFunction_Set()
	if err == nil {
		ret = contextGetMatrixFunction.Invoke(inArgs[:], nil)
	}

	retGo := MatrixNewFromNative(ret.Pointer())

	return retGo
}

var contextGetMetricsFunction *gi.Function
var contextGetMetricsFunction_Once sync.Once

func contextGetMetricsFunction_Set() error {
	var err error
	contextGetMetricsFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetMetricsFunction, err = contextObject.InvokerNew("get_metrics")
	})
	return err
}

// GetMetrics is a representation of the C type pango_context_get_metrics.
func (recv *Context) GetMetrics(desc *FontDescription, language *Language) *FontMetrics {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(desc.Native())
	inArgs[2].SetPointer(language.Native())

	var ret gi.Argument

	err := contextGetMetricsFunction_Set()
	if err == nil {
		ret = contextGetMetricsFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontMetricsNewFromNative(ret.Pointer())

	return retGo
}

var contextGetSerialFunction *gi.Function
var contextGetSerialFunction_Once sync.Once

func contextGetSerialFunction_Set() error {
	var err error
	contextGetSerialFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextGetSerialFunction, err = contextObject.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type pango_context_get_serial.
func (recv *Context) GetSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := contextGetSerialFunction_Set()
	if err == nil {
		ret = contextGetSerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'pango_context_list_families' : parameter 'families' of type 'nil' not supported

var contextLoadFontFunction *gi.Function
var contextLoadFontFunction_Once sync.Once

func contextLoadFontFunction_Set() error {
	var err error
	contextLoadFontFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextLoadFontFunction, err = contextObject.InvokerNew("load_font")
	})
	return err
}

// LoadFont is a representation of the C type pango_context_load_font.
func (recv *Context) LoadFont(desc *FontDescription) *Font {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(desc.Native())

	var ret gi.Argument

	err := contextLoadFontFunction_Set()
	if err == nil {
		ret = contextLoadFontFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontNewFromNative(ret.Pointer())

	return retGo
}

var contextLoadFontsetFunction *gi.Function
var contextLoadFontsetFunction_Once sync.Once

func contextLoadFontsetFunction_Set() error {
	var err error
	contextLoadFontsetFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextLoadFontsetFunction, err = contextObject.InvokerNew("load_fontset")
	})
	return err
}

// LoadFontset is a representation of the C type pango_context_load_fontset.
func (recv *Context) LoadFontset(desc *FontDescription, language *Language) *Fontset {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(desc.Native())
	inArgs[2].SetPointer(language.Native())

	var ret gi.Argument

	err := contextLoadFontsetFunction_Set()
	if err == nil {
		ret = contextLoadFontsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontsetNewFromNative(ret.Pointer())

	return retGo
}

var contextSetBaseDirFunction *gi.Function
var contextSetBaseDirFunction_Once sync.Once

func contextSetBaseDirFunction_Set() error {
	var err error
	contextSetBaseDirFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetBaseDirFunction, err = contextObject.InvokerNew("set_base_dir")
	})
	return err
}

// SetBaseDir is a representation of the C type pango_context_set_base_dir.
func (recv *Context) SetBaseDir(direction Direction) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(direction))

	err := contextSetBaseDirFunction_Set()
	if err == nil {
		contextSetBaseDirFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextSetBaseGravityFunction *gi.Function
var contextSetBaseGravityFunction_Once sync.Once

func contextSetBaseGravityFunction_Set() error {
	var err error
	contextSetBaseGravityFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetBaseGravityFunction, err = contextObject.InvokerNew("set_base_gravity")
	})
	return err
}

// SetBaseGravity is a representation of the C type pango_context_set_base_gravity.
func (recv *Context) SetBaseGravity(gravity Gravity) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(gravity))

	err := contextSetBaseGravityFunction_Set()
	if err == nil {
		contextSetBaseGravityFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextSetFontDescriptionFunction *gi.Function
var contextSetFontDescriptionFunction_Once sync.Once

func contextSetFontDescriptionFunction_Set() error {
	var err error
	contextSetFontDescriptionFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetFontDescriptionFunction, err = contextObject.InvokerNew("set_font_description")
	})
	return err
}

// SetFontDescription is a representation of the C type pango_context_set_font_description.
func (recv *Context) SetFontDescription(desc *FontDescription) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(desc.Native())

	err := contextSetFontDescriptionFunction_Set()
	if err == nil {
		contextSetFontDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextSetFontMapFunction *gi.Function
var contextSetFontMapFunction_Once sync.Once

func contextSetFontMapFunction_Set() error {
	var err error
	contextSetFontMapFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetFontMapFunction, err = contextObject.InvokerNew("set_font_map")
	})
	return err
}

// SetFontMap is a representation of the C type pango_context_set_font_map.
func (recv *Context) SetFontMap(fontMap *FontMap) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(fontMap.Native())

	err := contextSetFontMapFunction_Set()
	if err == nil {
		contextSetFontMapFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextSetGravityHintFunction *gi.Function
var contextSetGravityHintFunction_Once sync.Once

func contextSetGravityHintFunction_Set() error {
	var err error
	contextSetGravityHintFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetGravityHintFunction, err = contextObject.InvokerNew("set_gravity_hint")
	})
	return err
}

// SetGravityHint is a representation of the C type pango_context_set_gravity_hint.
func (recv *Context) SetGravityHint(hint GravityHint) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(hint))

	err := contextSetGravityHintFunction_Set()
	if err == nil {
		contextSetGravityHintFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextSetLanguageFunction *gi.Function
var contextSetLanguageFunction_Once sync.Once

func contextSetLanguageFunction_Set() error {
	var err error
	contextSetLanguageFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetLanguageFunction, err = contextObject.InvokerNew("set_language")
	})
	return err
}

// SetLanguage is a representation of the C type pango_context_set_language.
func (recv *Context) SetLanguage(language *Language) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(language.Native())

	err := contextSetLanguageFunction_Set()
	if err == nil {
		contextSetLanguageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var contextSetMatrixFunction *gi.Function
var contextSetMatrixFunction_Once sync.Once

func contextSetMatrixFunction_Set() error {
	var err error
	contextSetMatrixFunction_Once.Do(func() {
		err = contextObject_Set()
		if err != nil {
			return
		}
		contextSetMatrixFunction, err = contextObject.InvokerNew("set_matrix")
	})
	return err
}

// SetMatrix is a representation of the C type pango_context_set_matrix.
func (recv *Context) SetMatrix(matrix *Matrix) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(matrix.Native())

	err := contextSetMatrixFunction_Set()
	if err == nil {
		contextSetMatrixFunction.Invoke(inArgs[:], nil)
	}

	return
}

var engineObject *gi.Object
var engineObject_Once sync.Once

func engineObject_Set() error {
	var err error
	engineObject_Once.Do(func() {
		engineObject, err = gi.ObjectNew("Pango", "Engine")
	})
	return err
}

type Engine struct {
	native unsafe.Pointer
}

func EngineNewFromNative(native unsafe.Pointer) *Engine {
	instance := &Engine{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Engine) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToEngine down casts any arbitrary Object to Engine.
Exercise care, as this is a potentially dangerous function
if the Object is not a Engine.
*/
func (recv *Engine) CastToEngine(object *gobject.Object) *Engine {
	return EngineNewFromNative(object.Native())
}

func (recv *Engine) Native() unsafe.Pointer {
	return recv.native
}

var engineLangObject *gi.Object
var engineLangObject_Once sync.Once

func engineLangObject_Set() error {
	var err error
	engineLangObject_Once.Do(func() {
		engineLangObject, err = gi.ObjectNew("Pango", "EngineLang")
	})
	return err
}

type EngineLang struct {
	native unsafe.Pointer
}

func EngineLangNewFromNative(native unsafe.Pointer) *EngineLang {
	instance := &EngineLang{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Engine upcasts to *Engine
func (recv *EngineLang) Engine() *Engine {
	return EngineNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *EngineLang) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToEngineLang down casts any arbitrary Object to EngineLang.
Exercise care, as this is a potentially dangerous function
if the Object is not a EngineLang.
*/
func (recv *EngineLang) CastToEngineLang(object *gobject.Object) *EngineLang {
	return EngineLangNewFromNative(object.Native())
}

func (recv *EngineLang) Native() unsafe.Pointer {
	return recv.native
}

var engineShapeObject *gi.Object
var engineShapeObject_Once sync.Once

func engineShapeObject_Set() error {
	var err error
	engineShapeObject_Once.Do(func() {
		engineShapeObject, err = gi.ObjectNew("Pango", "EngineShape")
	})
	return err
}

type EngineShape struct {
	native unsafe.Pointer
}

func EngineShapeNewFromNative(native unsafe.Pointer) *EngineShape {
	instance := &EngineShape{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Engine upcasts to *Engine
func (recv *EngineShape) Engine() *Engine {
	return EngineNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *EngineShape) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToEngineShape down casts any arbitrary Object to EngineShape.
Exercise care, as this is a potentially dangerous function
if the Object is not a EngineShape.
*/
func (recv *EngineShape) CastToEngineShape(object *gobject.Object) *EngineShape {
	return EngineShapeNewFromNative(object.Native())
}

func (recv *EngineShape) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *EngineShape) FieldParentInstance() *Engine {
	argValue := gi.ObjectFieldGet(engineShapeObject, recv.Native(), "parent_instance")
	value := EngineNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *EngineShape) SetFieldParentInstance(value *Engine) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(engineShapeObject, recv.Native(), "parent_instance", argValue)
}

var fontObject *gi.Object
var fontObject_Once sync.Once

func fontObject_Set() error {
	var err error
	fontObject_Once.Do(func() {
		fontObject, err = gi.ObjectNew("Pango", "Font")
	})
	return err
}

type Font struct {
	native unsafe.Pointer
}

func FontNewFromNative(native unsafe.Pointer) *Font {
	instance := &Font{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Font) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFont down casts any arbitrary Object to Font.
Exercise care, as this is a potentially dangerous function
if the Object is not a Font.
*/
func (recv *Font) CastToFont(object *gobject.Object) *Font {
	return FontNewFromNative(object.Native())
}

func (recv *Font) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Font) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(fontObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Font) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fontObject, recv.Native(), "parent_instance", argValue)
}

var fontDescribeFunction *gi.Function
var fontDescribeFunction_Once sync.Once

func fontDescribeFunction_Set() error {
	var err error
	fontDescribeFunction_Once.Do(func() {
		err = fontObject_Set()
		if err != nil {
			return
		}
		fontDescribeFunction, err = fontObject.InvokerNew("describe")
	})
	return err
}

// Describe is a representation of the C type pango_font_describe.
func (recv *Font) Describe() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontDescribeFunction_Set()
	if err == nil {
		ret = fontDescribeFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var fontDescribeWithAbsoluteSizeFunction *gi.Function
var fontDescribeWithAbsoluteSizeFunction_Once sync.Once

func fontDescribeWithAbsoluteSizeFunction_Set() error {
	var err error
	fontDescribeWithAbsoluteSizeFunction_Once.Do(func() {
		err = fontObject_Set()
		if err != nil {
			return
		}
		fontDescribeWithAbsoluteSizeFunction, err = fontObject.InvokerNew("describe_with_absolute_size")
	})
	return err
}

// DescribeWithAbsoluteSize is a representation of the C type pango_font_describe_with_absolute_size.
func (recv *Font) DescribeWithAbsoluteSize() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontDescribeWithAbsoluteSizeFunction_Set()
	if err == nil {
		ret = fontDescribeWithAbsoluteSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var fontFindShaperFunction *gi.Function
var fontFindShaperFunction_Once sync.Once

func fontFindShaperFunction_Set() error {
	var err error
	fontFindShaperFunction_Once.Do(func() {
		err = fontObject_Set()
		if err != nil {
			return
		}
		fontFindShaperFunction, err = fontObject.InvokerNew("find_shaper")
	})
	return err
}

// FindShaper is a representation of the C type pango_font_find_shaper.
func (recv *Font) FindShaper(language *Language, ch uint32) *EngineShape {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(language.Native())
	inArgs[2].SetUint32(ch)

	var ret gi.Argument

	err := fontFindShaperFunction_Set()
	if err == nil {
		ret = fontFindShaperFunction.Invoke(inArgs[:], nil)
	}

	retGo := EngineShapeNewFromNative(ret.Pointer())

	return retGo
}

var fontGetCoverageFunction *gi.Function
var fontGetCoverageFunction_Once sync.Once

func fontGetCoverageFunction_Set() error {
	var err error
	fontGetCoverageFunction_Once.Do(func() {
		err = fontObject_Set()
		if err != nil {
			return
		}
		fontGetCoverageFunction, err = fontObject.InvokerNew("get_coverage")
	})
	return err
}

// GetCoverage is a representation of the C type pango_font_get_coverage.
func (recv *Font) GetCoverage(language *Language) *Coverage {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(language.Native())

	var ret gi.Argument

	err := fontGetCoverageFunction_Set()
	if err == nil {
		ret = fontGetCoverageFunction.Invoke(inArgs[:], nil)
	}

	retGo := CoverageNewFromNative(ret.Pointer())

	return retGo
}

var fontGetFontMapFunction *gi.Function
var fontGetFontMapFunction_Once sync.Once

func fontGetFontMapFunction_Set() error {
	var err error
	fontGetFontMapFunction_Once.Do(func() {
		err = fontObject_Set()
		if err != nil {
			return
		}
		fontGetFontMapFunction, err = fontObject.InvokerNew("get_font_map")
	})
	return err
}

// GetFontMap is a representation of the C type pango_font_get_font_map.
func (recv *Font) GetFontMap() *FontMap {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontGetFontMapFunction_Set()
	if err == nil {
		ret = fontGetFontMapFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontMapNewFromNative(ret.Pointer())

	return retGo
}

var fontGetGlyphExtentsFunction *gi.Function
var fontGetGlyphExtentsFunction_Once sync.Once

func fontGetGlyphExtentsFunction_Set() error {
	var err error
	fontGetGlyphExtentsFunction_Once.Do(func() {
		err = fontObject_Set()
		if err != nil {
			return
		}
		fontGetGlyphExtentsFunction, err = fontObject.InvokerNew("get_glyph_extents")
	})
	return err
}

// GetGlyphExtents is a representation of the C type pango_font_get_glyph_extents.
func (recv *Font) GetGlyphExtents(glyph Glyph) (*Rectangle, *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(uint32(glyph))

	var outArgs [2]gi.Argument

	err := fontGetGlyphExtentsFunction_Set()
	if err == nil {
		fontGetGlyphExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var fontGetMetricsFunction *gi.Function
var fontGetMetricsFunction_Once sync.Once

func fontGetMetricsFunction_Set() error {
	var err error
	fontGetMetricsFunction_Once.Do(func() {
		err = fontObject_Set()
		if err != nil {
			return
		}
		fontGetMetricsFunction, err = fontObject.InvokerNew("get_metrics")
	})
	return err
}

// GetMetrics is a representation of the C type pango_font_get_metrics.
func (recv *Font) GetMetrics(language *Language) *FontMetrics {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(language.Native())

	var ret gi.Argument

	err := fontGetMetricsFunction_Set()
	if err == nil {
		ret = fontGetMetricsFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontMetricsNewFromNative(ret.Pointer())

	return retGo
}

var fontFaceObject *gi.Object
var fontFaceObject_Once sync.Once

func fontFaceObject_Set() error {
	var err error
	fontFaceObject_Once.Do(func() {
		fontFaceObject, err = gi.ObjectNew("Pango", "FontFace")
	})
	return err
}

type FontFace struct {
	native unsafe.Pointer
}

func FontFaceNewFromNative(native unsafe.Pointer) *FontFace {
	instance := &FontFace{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FontFace) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFontFace down casts any arbitrary Object to FontFace.
Exercise care, as this is a potentially dangerous function
if the Object is not a FontFace.
*/
func (recv *FontFace) CastToFontFace(object *gobject.Object) *FontFace {
	return FontFaceNewFromNative(object.Native())
}

func (recv *FontFace) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FontFace) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(fontFaceObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FontFace) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fontFaceObject, recv.Native(), "parent_instance", argValue)
}

var fontFaceDescribeFunction *gi.Function
var fontFaceDescribeFunction_Once sync.Once

func fontFaceDescribeFunction_Set() error {
	var err error
	fontFaceDescribeFunction_Once.Do(func() {
		err = fontFaceObject_Set()
		if err != nil {
			return
		}
		fontFaceDescribeFunction, err = fontFaceObject.InvokerNew("describe")
	})
	return err
}

// Describe is a representation of the C type pango_font_face_describe.
func (recv *FontFace) Describe() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontFaceDescribeFunction_Set()
	if err == nil {
		ret = fontFaceDescribeFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var fontFaceGetFaceNameFunction *gi.Function
var fontFaceGetFaceNameFunction_Once sync.Once

func fontFaceGetFaceNameFunction_Set() error {
	var err error
	fontFaceGetFaceNameFunction_Once.Do(func() {
		err = fontFaceObject_Set()
		if err != nil {
			return
		}
		fontFaceGetFaceNameFunction, err = fontFaceObject.InvokerNew("get_face_name")
	})
	return err
}

// GetFaceName is a representation of the C type pango_font_face_get_face_name.
func (recv *FontFace) GetFaceName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontFaceGetFaceNameFunction_Set()
	if err == nil {
		ret = fontFaceGetFaceNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fontFaceIsSynthesizedFunction *gi.Function
var fontFaceIsSynthesizedFunction_Once sync.Once

func fontFaceIsSynthesizedFunction_Set() error {
	var err error
	fontFaceIsSynthesizedFunction_Once.Do(func() {
		err = fontFaceObject_Set()
		if err != nil {
			return
		}
		fontFaceIsSynthesizedFunction, err = fontFaceObject.InvokerNew("is_synthesized")
	})
	return err
}

// IsSynthesized is a representation of the C type pango_font_face_is_synthesized.
func (recv *FontFace) IsSynthesized() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontFaceIsSynthesizedFunction_Set()
	if err == nil {
		ret = fontFaceIsSynthesizedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_face_list_sizes' : parameter 'sizes' of type 'nil' not supported

var fontFamilyObject *gi.Object
var fontFamilyObject_Once sync.Once

func fontFamilyObject_Set() error {
	var err error
	fontFamilyObject_Once.Do(func() {
		fontFamilyObject, err = gi.ObjectNew("Pango", "FontFamily")
	})
	return err
}

type FontFamily struct {
	native unsafe.Pointer
}

func FontFamilyNewFromNative(native unsafe.Pointer) *FontFamily {
	instance := &FontFamily{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FontFamily) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFontFamily down casts any arbitrary Object to FontFamily.
Exercise care, as this is a potentially dangerous function
if the Object is not a FontFamily.
*/
func (recv *FontFamily) CastToFontFamily(object *gobject.Object) *FontFamily {
	return FontFamilyNewFromNative(object.Native())
}

func (recv *FontFamily) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FontFamily) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(fontFamilyObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FontFamily) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fontFamilyObject, recv.Native(), "parent_instance", argValue)
}

var fontFamilyGetNameFunction *gi.Function
var fontFamilyGetNameFunction_Once sync.Once

func fontFamilyGetNameFunction_Set() error {
	var err error
	fontFamilyGetNameFunction_Once.Do(func() {
		err = fontFamilyObject_Set()
		if err != nil {
			return
		}
		fontFamilyGetNameFunction, err = fontFamilyObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type pango_font_family_get_name.
func (recv *FontFamily) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontFamilyGetNameFunction_Set()
	if err == nil {
		ret = fontFamilyGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fontFamilyIsMonospaceFunction *gi.Function
var fontFamilyIsMonospaceFunction_Once sync.Once

func fontFamilyIsMonospaceFunction_Set() error {
	var err error
	fontFamilyIsMonospaceFunction_Once.Do(func() {
		err = fontFamilyObject_Set()
		if err != nil {
			return
		}
		fontFamilyIsMonospaceFunction, err = fontFamilyObject.InvokerNew("is_monospace")
	})
	return err
}

// IsMonospace is a representation of the C type pango_font_family_is_monospace.
func (recv *FontFamily) IsMonospace() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontFamilyIsMonospaceFunction_Set()
	if err == nil {
		ret = fontFamilyIsMonospaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_family_list_faces' : parameter 'faces' of type 'nil' not supported

var fontMapObject *gi.Object
var fontMapObject_Once sync.Once

func fontMapObject_Set() error {
	var err error
	fontMapObject_Once.Do(func() {
		fontMapObject, err = gi.ObjectNew("Pango", "FontMap")
	})
	return err
}

type FontMap struct {
	native unsafe.Pointer
}

func FontMapNewFromNative(native unsafe.Pointer) *FontMap {
	instance := &FontMap{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *FontMap) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFontMap down casts any arbitrary Object to FontMap.
Exercise care, as this is a potentially dangerous function
if the Object is not a FontMap.
*/
func (recv *FontMap) CastToFontMap(object *gobject.Object) *FontMap {
	return FontMapNewFromNative(object.Native())
}

func (recv *FontMap) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *FontMap) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(fontMapObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *FontMap) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fontMapObject, recv.Native(), "parent_instance", argValue)
}

var fontMapChangedFunction *gi.Function
var fontMapChangedFunction_Once sync.Once

func fontMapChangedFunction_Set() error {
	var err error
	fontMapChangedFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapChangedFunction, err = fontMapObject.InvokerNew("changed")
	})
	return err
}

// Changed is a representation of the C type pango_font_map_changed.
func (recv *FontMap) Changed() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := fontMapChangedFunction_Set()
	if err == nil {
		fontMapChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontMapCreateContextFunction *gi.Function
var fontMapCreateContextFunction_Once sync.Once

func fontMapCreateContextFunction_Set() error {
	var err error
	fontMapCreateContextFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapCreateContextFunction, err = fontMapObject.InvokerNew("create_context")
	})
	return err
}

// CreateContext is a representation of the C type pango_font_map_create_context.
func (recv *FontMap) CreateContext() *Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontMapCreateContextFunction_Set()
	if err == nil {
		ret = fontMapCreateContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ContextNewFromNative(ret.Pointer())

	return retGo
}

var fontMapGetSerialFunction *gi.Function
var fontMapGetSerialFunction_Once sync.Once

func fontMapGetSerialFunction_Set() error {
	var err error
	fontMapGetSerialFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapGetSerialFunction, err = fontMapObject.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type pango_font_map_get_serial.
func (recv *FontMap) GetSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontMapGetSerialFunction_Set()
	if err == nil {
		ret = fontMapGetSerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var fontMapGetShapeEngineTypeFunction *gi.Function
var fontMapGetShapeEngineTypeFunction_Once sync.Once

func fontMapGetShapeEngineTypeFunction_Set() error {
	var err error
	fontMapGetShapeEngineTypeFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapGetShapeEngineTypeFunction, err = fontMapObject.InvokerNew("get_shape_engine_type")
	})
	return err
}

// GetShapeEngineType is a representation of the C type pango_font_map_get_shape_engine_type.
func (recv *FontMap) GetShapeEngineType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontMapGetShapeEngineTypeFunction_Set()
	if err == nil {
		ret = fontMapGetShapeEngineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_map_list_families' : parameter 'families' of type 'nil' not supported

var fontMapLoadFontFunction *gi.Function
var fontMapLoadFontFunction_Once sync.Once

func fontMapLoadFontFunction_Set() error {
	var err error
	fontMapLoadFontFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapLoadFontFunction, err = fontMapObject.InvokerNew("load_font")
	})
	return err
}

// LoadFont is a representation of the C type pango_font_map_load_font.
func (recv *FontMap) LoadFont(context *Context, desc *FontDescription) *Font {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(context.Native())
	inArgs[2].SetPointer(desc.Native())

	var ret gi.Argument

	err := fontMapLoadFontFunction_Set()
	if err == nil {
		ret = fontMapLoadFontFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontNewFromNative(ret.Pointer())

	return retGo
}

var fontMapLoadFontsetFunction *gi.Function
var fontMapLoadFontsetFunction_Once sync.Once

func fontMapLoadFontsetFunction_Set() error {
	var err error
	fontMapLoadFontsetFunction_Once.Do(func() {
		err = fontMapObject_Set()
		if err != nil {
			return
		}
		fontMapLoadFontsetFunction, err = fontMapObject.InvokerNew("load_fontset")
	})
	return err
}

// LoadFontset is a representation of the C type pango_font_map_load_fontset.
func (recv *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) *Fontset {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(context.Native())
	inArgs[2].SetPointer(desc.Native())
	inArgs[3].SetPointer(language.Native())

	var ret gi.Argument

	err := fontMapLoadFontsetFunction_Set()
	if err == nil {
		ret = fontMapLoadFontsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontsetNewFromNative(ret.Pointer())

	return retGo
}

var fontsetObject *gi.Object
var fontsetObject_Once sync.Once

func fontsetObject_Set() error {
	var err error
	fontsetObject_Once.Do(func() {
		fontsetObject, err = gi.ObjectNew("Pango", "Fontset")
	})
	return err
}

type Fontset struct {
	native unsafe.Pointer
}

func FontsetNewFromNative(native unsafe.Pointer) *Fontset {
	instance := &Fontset{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Fontset) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFontset down casts any arbitrary Object to Fontset.
Exercise care, as this is a potentially dangerous function
if the Object is not a Fontset.
*/
func (recv *Fontset) CastToFontset(object *gobject.Object) *Fontset {
	return FontsetNewFromNative(object.Native())
}

func (recv *Fontset) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInstance returns the C field 'parent_instance'.
func (recv *Fontset) FieldParentInstance() *gobject.Object {
	argValue := gi.ObjectFieldGet(fontsetObject, recv.Native(), "parent_instance")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInstance sets the value of the C field 'parent_instance'.
func (recv *Fontset) SetFieldParentInstance(value *gobject.Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(fontsetObject, recv.Native(), "parent_instance", argValue)
}

// UNSUPPORTED : C value 'pango_fontset_foreach' : parameter 'func' of type 'FontsetForeachFunc' not supported

var fontsetGetFontFunction *gi.Function
var fontsetGetFontFunction_Once sync.Once

func fontsetGetFontFunction_Set() error {
	var err error
	fontsetGetFontFunction_Once.Do(func() {
		err = fontsetObject_Set()
		if err != nil {
			return
		}
		fontsetGetFontFunction, err = fontsetObject.InvokerNew("get_font")
	})
	return err
}

// GetFont is a representation of the C type pango_fontset_get_font.
func (recv *Fontset) GetFont(wc uint32) *Font {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(wc)

	var ret gi.Argument

	err := fontsetGetFontFunction_Set()
	if err == nil {
		ret = fontsetGetFontFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontNewFromNative(ret.Pointer())

	return retGo
}

var fontsetGetMetricsFunction *gi.Function
var fontsetGetMetricsFunction_Once sync.Once

func fontsetGetMetricsFunction_Set() error {
	var err error
	fontsetGetMetricsFunction_Once.Do(func() {
		err = fontsetObject_Set()
		if err != nil {
			return
		}
		fontsetGetMetricsFunction, err = fontsetObject.InvokerNew("get_metrics")
	})
	return err
}

// GetMetrics is a representation of the C type pango_fontset_get_metrics.
func (recv *Fontset) GetMetrics() *FontMetrics {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontsetGetMetricsFunction_Set()
	if err == nil {
		ret = fontsetGetMetricsFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontMetricsNewFromNative(ret.Pointer())

	return retGo
}

var fontsetSimpleObject *gi.Object
var fontsetSimpleObject_Once sync.Once

func fontsetSimpleObject_Set() error {
	var err error
	fontsetSimpleObject_Once.Do(func() {
		fontsetSimpleObject, err = gi.ObjectNew("Pango", "FontsetSimple")
	})
	return err
}

type FontsetSimple struct {
	native unsafe.Pointer
}

func FontsetSimpleNewFromNative(native unsafe.Pointer) *FontsetSimple {
	instance := &FontsetSimple{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Fontset upcasts to *Fontset
func (recv *FontsetSimple) Fontset() *Fontset {
	return FontsetNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *FontsetSimple) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToFontsetSimple down casts any arbitrary Object to FontsetSimple.
Exercise care, as this is a potentially dangerous function
if the Object is not a FontsetSimple.
*/
func (recv *FontsetSimple) CastToFontsetSimple(object *gobject.Object) *FontsetSimple {
	return FontsetSimpleNewFromNative(object.Native())
}

func (recv *FontsetSimple) Native() unsafe.Pointer {
	return recv.native
}

var fontsetSimpleNewFunction *gi.Function
var fontsetSimpleNewFunction_Once sync.Once

func fontsetSimpleNewFunction_Set() error {
	var err error
	fontsetSimpleNewFunction_Once.Do(func() {
		err = fontsetSimpleObject_Set()
		if err != nil {
			return
		}
		fontsetSimpleNewFunction, err = fontsetSimpleObject.InvokerNew("new")
	})
	return err
}

// FontsetSimpleNew is a representation of the C type pango_fontset_simple_new.
func FontsetSimpleNew(language *Language) *FontsetSimple {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(language.Native())

	var ret gi.Argument

	err := fontsetSimpleNewFunction_Set()
	if err == nil {
		ret = fontsetSimpleNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontsetSimpleNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var fontsetSimpleAppendFunction *gi.Function
var fontsetSimpleAppendFunction_Once sync.Once

func fontsetSimpleAppendFunction_Set() error {
	var err error
	fontsetSimpleAppendFunction_Once.Do(func() {
		err = fontsetSimpleObject_Set()
		if err != nil {
			return
		}
		fontsetSimpleAppendFunction, err = fontsetSimpleObject.InvokerNew("append")
	})
	return err
}

// Append is a representation of the C type pango_fontset_simple_append.
func (recv *FontsetSimple) Append(font *Font) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(font.Native())

	err := fontsetSimpleAppendFunction_Set()
	if err == nil {
		fontsetSimpleAppendFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontsetSimpleSizeFunction *gi.Function
var fontsetSimpleSizeFunction_Once sync.Once

func fontsetSimpleSizeFunction_Set() error {
	var err error
	fontsetSimpleSizeFunction_Once.Do(func() {
		err = fontsetSimpleObject_Set()
		if err != nil {
			return
		}
		fontsetSimpleSizeFunction, err = fontsetSimpleObject.InvokerNew("size")
	})
	return err
}

// Size is a representation of the C type pango_fontset_simple_size.
func (recv *FontsetSimple) Size() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontsetSimpleSizeFunction_Set()
	if err == nil {
		ret = fontsetSimpleSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutObject *gi.Object
var layoutObject_Once sync.Once

func layoutObject_Set() error {
	var err error
	layoutObject_Once.Do(func() {
		layoutObject, err = gi.ObjectNew("Pango", "Layout")
	})
	return err
}

type Layout struct {
	native unsafe.Pointer
}

func LayoutNewFromNative(native unsafe.Pointer) *Layout {
	instance := &Layout{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Layout) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToLayout down casts any arbitrary Object to Layout.
Exercise care, as this is a potentially dangerous function
if the Object is not a Layout.
*/
func (recv *Layout) CastToLayout(object *gobject.Object) *Layout {
	return LayoutNewFromNative(object.Native())
}

func (recv *Layout) Native() unsafe.Pointer {
	return recv.native
}

var layoutNewFunction *gi.Function
var layoutNewFunction_Once sync.Once

func layoutNewFunction_Set() error {
	var err error
	layoutNewFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutNewFunction, err = layoutObject.InvokerNew("new")
	})
	return err
}

// LayoutNew is a representation of the C type pango_layout_new.
func LayoutNew(context *Context) *Layout {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := layoutNewFunction_Set()
	if err == nil {
		ret = layoutNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var layoutContextChangedFunction *gi.Function
var layoutContextChangedFunction_Once sync.Once

func layoutContextChangedFunction_Set() error {
	var err error
	layoutContextChangedFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutContextChangedFunction, err = layoutObject.InvokerNew("context_changed")
	})
	return err
}

// ContextChanged is a representation of the C type pango_layout_context_changed.
func (recv *Layout) ContextChanged() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := layoutContextChangedFunction_Set()
	if err == nil {
		layoutContextChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutCopyFunction *gi.Function
var layoutCopyFunction_Once sync.Once

func layoutCopyFunction_Set() error {
	var err error
	layoutCopyFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutCopyFunction, err = layoutObject.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_layout_copy.
func (recv *Layout) Copy() *Layout {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutCopyFunction_Set()
	if err == nil {
		ret = layoutCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutNewFromNative(ret.Pointer())

	return retGo
}

var layoutGetAlignmentFunction *gi.Function
var layoutGetAlignmentFunction_Once sync.Once

func layoutGetAlignmentFunction_Set() error {
	var err error
	layoutGetAlignmentFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetAlignmentFunction, err = layoutObject.InvokerNew("get_alignment")
	})
	return err
}

// GetAlignment is a representation of the C type pango_layout_get_alignment.
func (recv *Layout) GetAlignment() Alignment {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetAlignmentFunction_Set()
	if err == nil {
		ret = layoutGetAlignmentFunction.Invoke(inArgs[:], nil)
	}

	retGo := Alignment(ret.Int32())

	return retGo
}

var layoutGetAttributesFunction *gi.Function
var layoutGetAttributesFunction_Once sync.Once

func layoutGetAttributesFunction_Set() error {
	var err error
	layoutGetAttributesFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetAttributesFunction, err = layoutObject.InvokerNew("get_attributes")
	})
	return err
}

// GetAttributes is a representation of the C type pango_layout_get_attributes.
func (recv *Layout) GetAttributes() *AttrList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetAttributesFunction_Set()
	if err == nil {
		ret = layoutGetAttributesFunction.Invoke(inArgs[:], nil)
	}

	retGo := AttrListNewFromNative(ret.Pointer())

	return retGo
}

var layoutGetAutoDirFunction *gi.Function
var layoutGetAutoDirFunction_Once sync.Once

func layoutGetAutoDirFunction_Set() error {
	var err error
	layoutGetAutoDirFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetAutoDirFunction, err = layoutObject.InvokerNew("get_auto_dir")
	})
	return err
}

// GetAutoDir is a representation of the C type pango_layout_get_auto_dir.
func (recv *Layout) GetAutoDir() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetAutoDirFunction_Set()
	if err == nil {
		ret = layoutGetAutoDirFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutGetBaselineFunction *gi.Function
var layoutGetBaselineFunction_Once sync.Once

func layoutGetBaselineFunction_Set() error {
	var err error
	layoutGetBaselineFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetBaselineFunction, err = layoutObject.InvokerNew("get_baseline")
	})
	return err
}

// GetBaseline is a representation of the C type pango_layout_get_baseline.
func (recv *Layout) GetBaseline() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetBaselineFunction_Set()
	if err == nil {
		ret = layoutGetBaselineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetCharacterCountFunction *gi.Function
var layoutGetCharacterCountFunction_Once sync.Once

func layoutGetCharacterCountFunction_Set() error {
	var err error
	layoutGetCharacterCountFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetCharacterCountFunction, err = layoutObject.InvokerNew("get_character_count")
	})
	return err
}

// GetCharacterCount is a representation of the C type pango_layout_get_character_count.
func (recv *Layout) GetCharacterCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetCharacterCountFunction_Set()
	if err == nil {
		ret = layoutGetCharacterCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetContextFunction *gi.Function
var layoutGetContextFunction_Once sync.Once

func layoutGetContextFunction_Set() error {
	var err error
	layoutGetContextFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetContextFunction, err = layoutObject.InvokerNew("get_context")
	})
	return err
}

// GetContext is a representation of the C type pango_layout_get_context.
func (recv *Layout) GetContext() *Context {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetContextFunction_Set()
	if err == nil {
		ret = layoutGetContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ContextNewFromNative(ret.Pointer())

	return retGo
}

var layoutGetCursorPosFunction *gi.Function
var layoutGetCursorPosFunction_Once sync.Once

func layoutGetCursorPosFunction_Set() error {
	var err error
	layoutGetCursorPosFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetCursorPosFunction, err = layoutObject.InvokerNew("get_cursor_pos")
	})
	return err
}

// GetCursorPos is a representation of the C type pango_layout_get_cursor_pos.
func (recv *Layout) GetCursorPos(index int32) (*Rectangle, *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(index)

	var outArgs [2]gi.Argument

	err := layoutGetCursorPosFunction_Set()
	if err == nil {
		layoutGetCursorPosFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var layoutGetEllipsizeFunction *gi.Function
var layoutGetEllipsizeFunction_Once sync.Once

func layoutGetEllipsizeFunction_Set() error {
	var err error
	layoutGetEllipsizeFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetEllipsizeFunction, err = layoutObject.InvokerNew("get_ellipsize")
	})
	return err
}

// GetEllipsize is a representation of the C type pango_layout_get_ellipsize.
func (recv *Layout) GetEllipsize() EllipsizeMode {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetEllipsizeFunction_Set()
	if err == nil {
		ret = layoutGetEllipsizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := EllipsizeMode(ret.Int32())

	return retGo
}

var layoutGetExtentsFunction *gi.Function
var layoutGetExtentsFunction_Once sync.Once

func layoutGetExtentsFunction_Set() error {
	var err error
	layoutGetExtentsFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetExtentsFunction, err = layoutObject.InvokerNew("get_extents")
	})
	return err
}

// GetExtents is a representation of the C type pango_layout_get_extents.
func (recv *Layout) GetExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument

	err := layoutGetExtentsFunction_Set()
	if err == nil {
		layoutGetExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var layoutGetFontDescriptionFunction *gi.Function
var layoutGetFontDescriptionFunction_Once sync.Once

func layoutGetFontDescriptionFunction_Set() error {
	var err error
	layoutGetFontDescriptionFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetFontDescriptionFunction, err = layoutObject.InvokerNew("get_font_description")
	})
	return err
}

// GetFontDescription is a representation of the C type pango_layout_get_font_description.
func (recv *Layout) GetFontDescription() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetFontDescriptionFunction_Set()
	if err == nil {
		ret = layoutGetFontDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var layoutGetHeightFunction *gi.Function
var layoutGetHeightFunction_Once sync.Once

func layoutGetHeightFunction_Set() error {
	var err error
	layoutGetHeightFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetHeightFunction, err = layoutObject.InvokerNew("get_height")
	})
	return err
}

// GetHeight is a representation of the C type pango_layout_get_height.
func (recv *Layout) GetHeight() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetHeightFunction_Set()
	if err == nil {
		ret = layoutGetHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetIndentFunction *gi.Function
var layoutGetIndentFunction_Once sync.Once

func layoutGetIndentFunction_Set() error {
	var err error
	layoutGetIndentFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetIndentFunction, err = layoutObject.InvokerNew("get_indent")
	})
	return err
}

// GetIndent is a representation of the C type pango_layout_get_indent.
func (recv *Layout) GetIndent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetIndentFunction_Set()
	if err == nil {
		ret = layoutGetIndentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetIterFunction *gi.Function
var layoutGetIterFunction_Once sync.Once

func layoutGetIterFunction_Set() error {
	var err error
	layoutGetIterFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetIterFunction, err = layoutObject.InvokerNew("get_iter")
	})
	return err
}

// GetIter is a representation of the C type pango_layout_get_iter.
func (recv *Layout) GetIter() *LayoutIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetIterFunction_Set()
	if err == nil {
		ret = layoutGetIterFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutIterNewFromNative(ret.Pointer())

	return retGo
}

var layoutGetJustifyFunction *gi.Function
var layoutGetJustifyFunction_Once sync.Once

func layoutGetJustifyFunction_Set() error {
	var err error
	layoutGetJustifyFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetJustifyFunction, err = layoutObject.InvokerNew("get_justify")
	})
	return err
}

// GetJustify is a representation of the C type pango_layout_get_justify.
func (recv *Layout) GetJustify() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetJustifyFunction_Set()
	if err == nil {
		ret = layoutGetJustifyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutGetLineFunction *gi.Function
var layoutGetLineFunction_Once sync.Once

func layoutGetLineFunction_Set() error {
	var err error
	layoutGetLineFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetLineFunction, err = layoutObject.InvokerNew("get_line")
	})
	return err
}

// GetLine is a representation of the C type pango_layout_get_line.
func (recv *Layout) GetLine(line int32) *LayoutLine {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(line)

	var ret gi.Argument

	err := layoutGetLineFunction_Set()
	if err == nil {
		ret = layoutGetLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutLineNewFromNative(ret.Pointer())

	return retGo
}

var layoutGetLineCountFunction *gi.Function
var layoutGetLineCountFunction_Once sync.Once

func layoutGetLineCountFunction_Set() error {
	var err error
	layoutGetLineCountFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetLineCountFunction, err = layoutObject.InvokerNew("get_line_count")
	})
	return err
}

// GetLineCount is a representation of the C type pango_layout_get_line_count.
func (recv *Layout) GetLineCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetLineCountFunction_Set()
	if err == nil {
		ret = layoutGetLineCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetLineReadonlyFunction *gi.Function
var layoutGetLineReadonlyFunction_Once sync.Once

func layoutGetLineReadonlyFunction_Set() error {
	var err error
	layoutGetLineReadonlyFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetLineReadonlyFunction, err = layoutObject.InvokerNew("get_line_readonly")
	})
	return err
}

// GetLineReadonly is a representation of the C type pango_layout_get_line_readonly.
func (recv *Layout) GetLineReadonly(line int32) *LayoutLine {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(line)

	var ret gi.Argument

	err := layoutGetLineReadonlyFunction_Set()
	if err == nil {
		ret = layoutGetLineReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutLineNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_get_lines' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_layout_get_lines_readonly' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_layout_get_log_attrs' : parameter 'attrs' of type 'nil' not supported

var layoutGetLogAttrsReadonlyFunction *gi.Function
var layoutGetLogAttrsReadonlyFunction_Once sync.Once

func layoutGetLogAttrsReadonlyFunction_Set() error {
	var err error
	layoutGetLogAttrsReadonlyFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetLogAttrsReadonlyFunction, err = layoutObject.InvokerNew("get_log_attrs_readonly")
	})
	return err
}

// GetLogAttrsReadonly is a representation of the C type pango_layout_get_log_attrs_readonly.
func (recv *Layout) GetLogAttrsReadonly() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := layoutGetLogAttrsReadonlyFunction_Set()
	if err == nil {
		layoutGetLogAttrsReadonlyFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var layoutGetPixelExtentsFunction *gi.Function
var layoutGetPixelExtentsFunction_Once sync.Once

func layoutGetPixelExtentsFunction_Set() error {
	var err error
	layoutGetPixelExtentsFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetPixelExtentsFunction, err = layoutObject.InvokerNew("get_pixel_extents")
	})
	return err
}

// GetPixelExtents is a representation of the C type pango_layout_get_pixel_extents.
func (recv *Layout) GetPixelExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument

	err := layoutGetPixelExtentsFunction_Set()
	if err == nil {
		layoutGetPixelExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var layoutGetPixelSizeFunction *gi.Function
var layoutGetPixelSizeFunction_Once sync.Once

func layoutGetPixelSizeFunction_Set() error {
	var err error
	layoutGetPixelSizeFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetPixelSizeFunction, err = layoutObject.InvokerNew("get_pixel_size")
	})
	return err
}

// GetPixelSize is a representation of the C type pango_layout_get_pixel_size.
func (recv *Layout) GetPixelSize() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument

	err := layoutGetPixelSizeFunction_Set()
	if err == nil {
		layoutGetPixelSizeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var layoutGetSerialFunction *gi.Function
var layoutGetSerialFunction_Once sync.Once

func layoutGetSerialFunction_Set() error {
	var err error
	layoutGetSerialFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetSerialFunction, err = layoutObject.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type pango_layout_get_serial.
func (recv *Layout) GetSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetSerialFunction_Set()
	if err == nil {
		ret = layoutGetSerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var layoutGetSingleParagraphModeFunction *gi.Function
var layoutGetSingleParagraphModeFunction_Once sync.Once

func layoutGetSingleParagraphModeFunction_Set() error {
	var err error
	layoutGetSingleParagraphModeFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetSingleParagraphModeFunction, err = layoutObject.InvokerNew("get_single_paragraph_mode")
	})
	return err
}

// GetSingleParagraphMode is a representation of the C type pango_layout_get_single_paragraph_mode.
func (recv *Layout) GetSingleParagraphMode() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetSingleParagraphModeFunction_Set()
	if err == nil {
		ret = layoutGetSingleParagraphModeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutGetSizeFunction *gi.Function
var layoutGetSizeFunction_Once sync.Once

func layoutGetSizeFunction_Set() error {
	var err error
	layoutGetSizeFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetSizeFunction, err = layoutObject.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type pango_layout_get_size.
func (recv *Layout) GetSize() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument

	err := layoutGetSizeFunction_Set()
	if err == nil {
		layoutGetSizeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var layoutGetSpacingFunction *gi.Function
var layoutGetSpacingFunction_Once sync.Once

func layoutGetSpacingFunction_Set() error {
	var err error
	layoutGetSpacingFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetSpacingFunction, err = layoutObject.InvokerNew("get_spacing")
	})
	return err
}

// GetSpacing is a representation of the C type pango_layout_get_spacing.
func (recv *Layout) GetSpacing() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetSpacingFunction_Set()
	if err == nil {
		ret = layoutGetSpacingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetTabsFunction *gi.Function
var layoutGetTabsFunction_Once sync.Once

func layoutGetTabsFunction_Set() error {
	var err error
	layoutGetTabsFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetTabsFunction, err = layoutObject.InvokerNew("get_tabs")
	})
	return err
}

// GetTabs is a representation of the C type pango_layout_get_tabs.
func (recv *Layout) GetTabs() *TabArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetTabsFunction_Set()
	if err == nil {
		ret = layoutGetTabsFunction.Invoke(inArgs[:], nil)
	}

	retGo := TabArrayNewFromNative(ret.Pointer())

	return retGo
}

var layoutGetTextFunction *gi.Function
var layoutGetTextFunction_Once sync.Once

func layoutGetTextFunction_Set() error {
	var err error
	layoutGetTextFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetTextFunction, err = layoutObject.InvokerNew("get_text")
	})
	return err
}

// GetText is a representation of the C type pango_layout_get_text.
func (recv *Layout) GetText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetTextFunction_Set()
	if err == nil {
		ret = layoutGetTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var layoutGetUnknownGlyphsCountFunction *gi.Function
var layoutGetUnknownGlyphsCountFunction_Once sync.Once

func layoutGetUnknownGlyphsCountFunction_Set() error {
	var err error
	layoutGetUnknownGlyphsCountFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetUnknownGlyphsCountFunction, err = layoutObject.InvokerNew("get_unknown_glyphs_count")
	})
	return err
}

// GetUnknownGlyphsCount is a representation of the C type pango_layout_get_unknown_glyphs_count.
func (recv *Layout) GetUnknownGlyphsCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetUnknownGlyphsCountFunction_Set()
	if err == nil {
		ret = layoutGetUnknownGlyphsCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetWidthFunction *gi.Function
var layoutGetWidthFunction_Once sync.Once

func layoutGetWidthFunction_Set() error {
	var err error
	layoutGetWidthFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetWidthFunction, err = layoutObject.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type pango_layout_get_width.
func (recv *Layout) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetWidthFunction_Set()
	if err == nil {
		ret = layoutGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutGetWrapFunction *gi.Function
var layoutGetWrapFunction_Once sync.Once

func layoutGetWrapFunction_Set() error {
	var err error
	layoutGetWrapFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutGetWrapFunction, err = layoutObject.InvokerNew("get_wrap")
	})
	return err
}

// GetWrap is a representation of the C type pango_layout_get_wrap.
func (recv *Layout) GetWrap() WrapMode {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutGetWrapFunction_Set()
	if err == nil {
		ret = layoutGetWrapFunction.Invoke(inArgs[:], nil)
	}

	retGo := WrapMode(ret.Int32())

	return retGo
}

var layoutIndexToLineXFunction *gi.Function
var layoutIndexToLineXFunction_Once sync.Once

func layoutIndexToLineXFunction_Set() error {
	var err error
	layoutIndexToLineXFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutIndexToLineXFunction, err = layoutObject.InvokerNew("index_to_line_x")
	})
	return err
}

// IndexToLineX is a representation of the C type pango_layout_index_to_line_x.
func (recv *Layout) IndexToLineX(index int32, trailing bool) (int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(index)
	inArgs[2].SetBoolean(trailing)

	var outArgs [2]gi.Argument

	err := layoutIndexToLineXFunction_Set()
	if err == nil {
		layoutIndexToLineXFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var layoutIndexToPosFunction *gi.Function
var layoutIndexToPosFunction_Once sync.Once

func layoutIndexToPosFunction_Set() error {
	var err error
	layoutIndexToPosFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutIndexToPosFunction, err = layoutObject.InvokerNew("index_to_pos")
	})
	return err
}

// IndexToPos is a representation of the C type pango_layout_index_to_pos.
func (recv *Layout) IndexToPos(index int32) *Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(index)

	var outArgs [1]gi.Argument

	err := layoutIndexToPosFunction_Set()
	if err == nil {
		layoutIndexToPosFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())

	return out0
}

var layoutIsEllipsizedFunction *gi.Function
var layoutIsEllipsizedFunction_Once sync.Once

func layoutIsEllipsizedFunction_Set() error {
	var err error
	layoutIsEllipsizedFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutIsEllipsizedFunction, err = layoutObject.InvokerNew("is_ellipsized")
	})
	return err
}

// IsEllipsized is a representation of the C type pango_layout_is_ellipsized.
func (recv *Layout) IsEllipsized() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutIsEllipsizedFunction_Set()
	if err == nil {
		ret = layoutIsEllipsizedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutIsWrappedFunction *gi.Function
var layoutIsWrappedFunction_Once sync.Once

func layoutIsWrappedFunction_Set() error {
	var err error
	layoutIsWrappedFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutIsWrappedFunction, err = layoutObject.InvokerNew("is_wrapped")
	})
	return err
}

// IsWrapped is a representation of the C type pango_layout_is_wrapped.
func (recv *Layout) IsWrapped() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := layoutIsWrappedFunction_Set()
	if err == nil {
		ret = layoutIsWrappedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutMoveCursorVisuallyFunction *gi.Function
var layoutMoveCursorVisuallyFunction_Once sync.Once

func layoutMoveCursorVisuallyFunction_Set() error {
	var err error
	layoutMoveCursorVisuallyFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutMoveCursorVisuallyFunction, err = layoutObject.InvokerNew("move_cursor_visually")
	})
	return err
}

// MoveCursorVisually is a representation of the C type pango_layout_move_cursor_visually.
func (recv *Layout) MoveCursorVisually(strong bool, oldIndex int32, oldTrailing int32, direction int32) (int32, int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(strong)
	inArgs[2].SetInt32(oldIndex)
	inArgs[3].SetInt32(oldTrailing)
	inArgs[4].SetInt32(direction)

	var outArgs [2]gi.Argument

	err := layoutMoveCursorVisuallyFunction_Set()
	if err == nil {
		layoutMoveCursorVisuallyFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var layoutSetAlignmentFunction *gi.Function
var layoutSetAlignmentFunction_Once sync.Once

func layoutSetAlignmentFunction_Set() error {
	var err error
	layoutSetAlignmentFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetAlignmentFunction, err = layoutObject.InvokerNew("set_alignment")
	})
	return err
}

// SetAlignment is a representation of the C type pango_layout_set_alignment.
func (recv *Layout) SetAlignment(alignment Alignment) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(alignment))

	err := layoutSetAlignmentFunction_Set()
	if err == nil {
		layoutSetAlignmentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetAttributesFunction *gi.Function
var layoutSetAttributesFunction_Once sync.Once

func layoutSetAttributesFunction_Set() error {
	var err error
	layoutSetAttributesFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetAttributesFunction, err = layoutObject.InvokerNew("set_attributes")
	})
	return err
}

// SetAttributes is a representation of the C type pango_layout_set_attributes.
func (recv *Layout) SetAttributes(attrs *AttrList) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(attrs.Native())

	err := layoutSetAttributesFunction_Set()
	if err == nil {
		layoutSetAttributesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetAutoDirFunction *gi.Function
var layoutSetAutoDirFunction_Once sync.Once

func layoutSetAutoDirFunction_Set() error {
	var err error
	layoutSetAutoDirFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetAutoDirFunction, err = layoutObject.InvokerNew("set_auto_dir")
	})
	return err
}

// SetAutoDir is a representation of the C type pango_layout_set_auto_dir.
func (recv *Layout) SetAutoDir(autoDir bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(autoDir)

	err := layoutSetAutoDirFunction_Set()
	if err == nil {
		layoutSetAutoDirFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetEllipsizeFunction *gi.Function
var layoutSetEllipsizeFunction_Once sync.Once

func layoutSetEllipsizeFunction_Set() error {
	var err error
	layoutSetEllipsizeFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetEllipsizeFunction, err = layoutObject.InvokerNew("set_ellipsize")
	})
	return err
}

// SetEllipsize is a representation of the C type pango_layout_set_ellipsize.
func (recv *Layout) SetEllipsize(ellipsize EllipsizeMode) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(ellipsize))

	err := layoutSetEllipsizeFunction_Set()
	if err == nil {
		layoutSetEllipsizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetFontDescriptionFunction *gi.Function
var layoutSetFontDescriptionFunction_Once sync.Once

func layoutSetFontDescriptionFunction_Set() error {
	var err error
	layoutSetFontDescriptionFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetFontDescriptionFunction, err = layoutObject.InvokerNew("set_font_description")
	})
	return err
}

// SetFontDescription is a representation of the C type pango_layout_set_font_description.
func (recv *Layout) SetFontDescription(desc *FontDescription) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(desc.Native())

	err := layoutSetFontDescriptionFunction_Set()
	if err == nil {
		layoutSetFontDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetHeightFunction *gi.Function
var layoutSetHeightFunction_Once sync.Once

func layoutSetHeightFunction_Set() error {
	var err error
	layoutSetHeightFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetHeightFunction, err = layoutObject.InvokerNew("set_height")
	})
	return err
}

// SetHeight is a representation of the C type pango_layout_set_height.
func (recv *Layout) SetHeight(height int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(height)

	err := layoutSetHeightFunction_Set()
	if err == nil {
		layoutSetHeightFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetIndentFunction *gi.Function
var layoutSetIndentFunction_Once sync.Once

func layoutSetIndentFunction_Set() error {
	var err error
	layoutSetIndentFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetIndentFunction, err = layoutObject.InvokerNew("set_indent")
	})
	return err
}

// SetIndent is a representation of the C type pango_layout_set_indent.
func (recv *Layout) SetIndent(indent int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(indent)

	err := layoutSetIndentFunction_Set()
	if err == nil {
		layoutSetIndentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetJustifyFunction *gi.Function
var layoutSetJustifyFunction_Once sync.Once

func layoutSetJustifyFunction_Set() error {
	var err error
	layoutSetJustifyFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetJustifyFunction, err = layoutObject.InvokerNew("set_justify")
	})
	return err
}

// SetJustify is a representation of the C type pango_layout_set_justify.
func (recv *Layout) SetJustify(justify bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(justify)

	err := layoutSetJustifyFunction_Set()
	if err == nil {
		layoutSetJustifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetMarkupFunction *gi.Function
var layoutSetMarkupFunction_Once sync.Once

func layoutSetMarkupFunction_Set() error {
	var err error
	layoutSetMarkupFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetMarkupFunction, err = layoutObject.InvokerNew("set_markup")
	})
	return err
}

// SetMarkup is a representation of the C type pango_layout_set_markup.
func (recv *Layout) SetMarkup(markup string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(markup)
	inArgs[2].SetInt32(length)

	err := layoutSetMarkupFunction_Set()
	if err == nil {
		layoutSetMarkupFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_layout_set_markup_with_accel' : parameter 'accel_marker' of type 'gunichar' not supported

var layoutSetSingleParagraphModeFunction *gi.Function
var layoutSetSingleParagraphModeFunction_Once sync.Once

func layoutSetSingleParagraphModeFunction_Set() error {
	var err error
	layoutSetSingleParagraphModeFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetSingleParagraphModeFunction, err = layoutObject.InvokerNew("set_single_paragraph_mode")
	})
	return err
}

// SetSingleParagraphMode is a representation of the C type pango_layout_set_single_paragraph_mode.
func (recv *Layout) SetSingleParagraphMode(setting bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(setting)

	err := layoutSetSingleParagraphModeFunction_Set()
	if err == nil {
		layoutSetSingleParagraphModeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetSpacingFunction *gi.Function
var layoutSetSpacingFunction_Once sync.Once

func layoutSetSpacingFunction_Set() error {
	var err error
	layoutSetSpacingFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetSpacingFunction, err = layoutObject.InvokerNew("set_spacing")
	})
	return err
}

// SetSpacing is a representation of the C type pango_layout_set_spacing.
func (recv *Layout) SetSpacing(spacing int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(spacing)

	err := layoutSetSpacingFunction_Set()
	if err == nil {
		layoutSetSpacingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetTabsFunction *gi.Function
var layoutSetTabsFunction_Once sync.Once

func layoutSetTabsFunction_Set() error {
	var err error
	layoutSetTabsFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetTabsFunction, err = layoutObject.InvokerNew("set_tabs")
	})
	return err
}

// SetTabs is a representation of the C type pango_layout_set_tabs.
func (recv *Layout) SetTabs(tabs *TabArray) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(tabs.Native())

	err := layoutSetTabsFunction_Set()
	if err == nil {
		layoutSetTabsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetTextFunction *gi.Function
var layoutSetTextFunction_Once sync.Once

func layoutSetTextFunction_Set() error {
	var err error
	layoutSetTextFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetTextFunction, err = layoutObject.InvokerNew("set_text")
	})
	return err
}

// SetText is a representation of the C type pango_layout_set_text.
func (recv *Layout) SetText(text string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(length)

	err := layoutSetTextFunction_Set()
	if err == nil {
		layoutSetTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetWidthFunction *gi.Function
var layoutSetWidthFunction_Once sync.Once

func layoutSetWidthFunction_Set() error {
	var err error
	layoutSetWidthFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetWidthFunction, err = layoutObject.InvokerNew("set_width")
	})
	return err
}

// SetWidth is a representation of the C type pango_layout_set_width.
func (recv *Layout) SetWidth(width int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(width)

	err := layoutSetWidthFunction_Set()
	if err == nil {
		layoutSetWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutSetWrapFunction *gi.Function
var layoutSetWrapFunction_Once sync.Once

func layoutSetWrapFunction_Set() error {
	var err error
	layoutSetWrapFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutSetWrapFunction, err = layoutObject.InvokerNew("set_wrap")
	})
	return err
}

// SetWrap is a representation of the C type pango_layout_set_wrap.
func (recv *Layout) SetWrap(wrap WrapMode) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(wrap))

	err := layoutSetWrapFunction_Set()
	if err == nil {
		layoutSetWrapFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutXyToIndexFunction *gi.Function
var layoutXyToIndexFunction_Once sync.Once

func layoutXyToIndexFunction_Set() error {
	var err error
	layoutXyToIndexFunction_Once.Do(func() {
		err = layoutObject_Set()
		if err != nil {
			return
		}
		layoutXyToIndexFunction, err = layoutObject.InvokerNew("xy_to_index")
	})
	return err
}

// XyToIndex is a representation of the C type pango_layout_xy_to_index.
func (recv *Layout) XyToIndex(x int32, y int32) (bool, int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := layoutXyToIndexFunction_Set()
	if err == nil {
		ret = layoutXyToIndexFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var rendererObject *gi.Object
var rendererObject_Once sync.Once

func rendererObject_Set() error {
	var err error
	rendererObject_Once.Do(func() {
		rendererObject, err = gi.ObjectNew("Pango", "Renderer")
	})
	return err
}

type Renderer struct {
	native unsafe.Pointer
}

func RendererNewFromNative(native unsafe.Pointer) *Renderer {
	instance := &Renderer{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Renderer) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRenderer down casts any arbitrary Object to Renderer.
Exercise care, as this is a potentially dangerous function
if the Object is not a Renderer.
*/
func (recv *Renderer) CastToRenderer(object *gobject.Object) *Renderer {
	return RendererNewFromNative(object.Native())
}

func (recv *Renderer) Native() unsafe.Pointer {
	return recv.native
}

// FieldMatrix returns the C field 'matrix'.
func (recv *Renderer) FieldMatrix() *Matrix {
	argValue := gi.ObjectFieldGet(rendererObject, recv.Native(), "matrix")
	value := MatrixNewFromNative(argValue.Pointer())
	return value
}

// SetFieldMatrix sets the value of the C field 'matrix'.
func (recv *Renderer) SetFieldMatrix(value *Matrix) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(rendererObject, recv.Native(), "matrix", argValue)
}

var rendererActivateFunction *gi.Function
var rendererActivateFunction_Once sync.Once

func rendererActivateFunction_Set() error {
	var err error
	rendererActivateFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererActivateFunction, err = rendererObject.InvokerNew("activate")
	})
	return err
}

// Activate is a representation of the C type pango_renderer_activate.
func (recv *Renderer) Activate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := rendererActivateFunction_Set()
	if err == nil {
		rendererActivateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDeactivateFunction *gi.Function
var rendererDeactivateFunction_Once sync.Once

func rendererDeactivateFunction_Set() error {
	var err error
	rendererDeactivateFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDeactivateFunction, err = rendererObject.InvokerNew("deactivate")
	})
	return err
}

// Deactivate is a representation of the C type pango_renderer_deactivate.
func (recv *Renderer) Deactivate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := rendererDeactivateFunction_Set()
	if err == nil {
		rendererDeactivateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawErrorUnderlineFunction *gi.Function
var rendererDrawErrorUnderlineFunction_Once sync.Once

func rendererDrawErrorUnderlineFunction_Set() error {
	var err error
	rendererDrawErrorUnderlineFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawErrorUnderlineFunction, err = rendererObject.InvokerNew("draw_error_underline")
	})
	return err
}

// DrawErrorUnderline is a representation of the C type pango_renderer_draw_error_underline.
func (recv *Renderer) DrawErrorUnderline(x int32, y int32, width int32, height int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)

	err := rendererDrawErrorUnderlineFunction_Set()
	if err == nil {
		rendererDrawErrorUnderlineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawGlyphFunction *gi.Function
var rendererDrawGlyphFunction_Once sync.Once

func rendererDrawGlyphFunction_Set() error {
	var err error
	rendererDrawGlyphFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawGlyphFunction, err = rendererObject.InvokerNew("draw_glyph")
	})
	return err
}

// DrawGlyph is a representation of the C type pango_renderer_draw_glyph.
func (recv *Renderer) DrawGlyph(font *Font, glyph Glyph, x float64, y float64) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(font.Native())
	inArgs[2].SetUint32(uint32(glyph))
	inArgs[3].SetFloat64(x)
	inArgs[4].SetFloat64(y)

	err := rendererDrawGlyphFunction_Set()
	if err == nil {
		rendererDrawGlyphFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawGlyphItemFunction *gi.Function
var rendererDrawGlyphItemFunction_Once sync.Once

func rendererDrawGlyphItemFunction_Set() error {
	var err error
	rendererDrawGlyphItemFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawGlyphItemFunction, err = rendererObject.InvokerNew("draw_glyph_item")
	})
	return err
}

// DrawGlyphItem is a representation of the C type pango_renderer_draw_glyph_item.
func (recv *Renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x int32, y int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(text)
	inArgs[2].SetPointer(glyphItem.Native())
	inArgs[3].SetInt32(x)
	inArgs[4].SetInt32(y)

	err := rendererDrawGlyphItemFunction_Set()
	if err == nil {
		rendererDrawGlyphItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawGlyphsFunction *gi.Function
var rendererDrawGlyphsFunction_Once sync.Once

func rendererDrawGlyphsFunction_Set() error {
	var err error
	rendererDrawGlyphsFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawGlyphsFunction, err = rendererObject.InvokerNew("draw_glyphs")
	})
	return err
}

// DrawGlyphs is a representation of the C type pango_renderer_draw_glyphs.
func (recv *Renderer) DrawGlyphs(font *Font, glyphs *GlyphString, x int32, y int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(font.Native())
	inArgs[2].SetPointer(glyphs.Native())
	inArgs[3].SetInt32(x)
	inArgs[4].SetInt32(y)

	err := rendererDrawGlyphsFunction_Set()
	if err == nil {
		rendererDrawGlyphsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawLayoutFunction *gi.Function
var rendererDrawLayoutFunction_Once sync.Once

func rendererDrawLayoutFunction_Set() error {
	var err error
	rendererDrawLayoutFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawLayoutFunction, err = rendererObject.InvokerNew("draw_layout")
	})
	return err
}

// DrawLayout is a representation of the C type pango_renderer_draw_layout.
func (recv *Renderer) DrawLayout(layout *Layout, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(layout.Native())
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := rendererDrawLayoutFunction_Set()
	if err == nil {
		rendererDrawLayoutFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawLayoutLineFunction *gi.Function
var rendererDrawLayoutLineFunction_Once sync.Once

func rendererDrawLayoutLineFunction_Set() error {
	var err error
	rendererDrawLayoutLineFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawLayoutLineFunction, err = rendererObject.InvokerNew("draw_layout_line")
	})
	return err
}

// DrawLayoutLine is a representation of the C type pango_renderer_draw_layout_line.
func (recv *Renderer) DrawLayoutLine(line *LayoutLine, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(line.Native())
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := rendererDrawLayoutLineFunction_Set()
	if err == nil {
		rendererDrawLayoutLineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawRectangleFunction *gi.Function
var rendererDrawRectangleFunction_Once sync.Once

func rendererDrawRectangleFunction_Set() error {
	var err error
	rendererDrawRectangleFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawRectangleFunction, err = rendererObject.InvokerNew("draw_rectangle")
	})
	return err
}

// DrawRectangle is a representation of the C type pango_renderer_draw_rectangle.
func (recv *Renderer) DrawRectangle(part RenderPart, x int32, y int32, width int32, height int32) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(part))
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)
	inArgs[4].SetInt32(width)
	inArgs[5].SetInt32(height)

	err := rendererDrawRectangleFunction_Set()
	if err == nil {
		rendererDrawRectangleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererDrawTrapezoidFunction *gi.Function
var rendererDrawTrapezoidFunction_Once sync.Once

func rendererDrawTrapezoidFunction_Set() error {
	var err error
	rendererDrawTrapezoidFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererDrawTrapezoidFunction, err = rendererObject.InvokerNew("draw_trapezoid")
	})
	return err
}

// DrawTrapezoid is a representation of the C type pango_renderer_draw_trapezoid.
func (recv *Renderer) DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	var inArgs [8]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(part))
	inArgs[2].SetFloat64(y1)
	inArgs[3].SetFloat64(x11)
	inArgs[4].SetFloat64(x21)
	inArgs[5].SetFloat64(y2)
	inArgs[6].SetFloat64(x12)
	inArgs[7].SetFloat64(x22)

	err := rendererDrawTrapezoidFunction_Set()
	if err == nil {
		rendererDrawTrapezoidFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererGetAlphaFunction *gi.Function
var rendererGetAlphaFunction_Once sync.Once

func rendererGetAlphaFunction_Set() error {
	var err error
	rendererGetAlphaFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererGetAlphaFunction, err = rendererObject.InvokerNew("get_alpha")
	})
	return err
}

// GetAlpha is a representation of the C type pango_renderer_get_alpha.
func (recv *Renderer) GetAlpha(part RenderPart) uint16 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(part))

	var ret gi.Argument

	err := rendererGetAlphaFunction_Set()
	if err == nil {
		ret = rendererGetAlphaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var rendererGetColorFunction *gi.Function
var rendererGetColorFunction_Once sync.Once

func rendererGetColorFunction_Set() error {
	var err error
	rendererGetColorFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererGetColorFunction, err = rendererObject.InvokerNew("get_color")
	})
	return err
}

// GetColor is a representation of the C type pango_renderer_get_color.
func (recv *Renderer) GetColor(part RenderPart) *Color {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(part))

	var ret gi.Argument

	err := rendererGetColorFunction_Set()
	if err == nil {
		ret = rendererGetColorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ColorNewFromNative(ret.Pointer())

	return retGo
}

var rendererGetLayoutFunction *gi.Function
var rendererGetLayoutFunction_Once sync.Once

func rendererGetLayoutFunction_Set() error {
	var err error
	rendererGetLayoutFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererGetLayoutFunction, err = rendererObject.InvokerNew("get_layout")
	})
	return err
}

// GetLayout is a representation of the C type pango_renderer_get_layout.
func (recv *Renderer) GetLayout() *Layout {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := rendererGetLayoutFunction_Set()
	if err == nil {
		ret = rendererGetLayoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutNewFromNative(ret.Pointer())

	return retGo
}

var rendererGetLayoutLineFunction *gi.Function
var rendererGetLayoutLineFunction_Once sync.Once

func rendererGetLayoutLineFunction_Set() error {
	var err error
	rendererGetLayoutLineFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererGetLayoutLineFunction, err = rendererObject.InvokerNew("get_layout_line")
	})
	return err
}

// GetLayoutLine is a representation of the C type pango_renderer_get_layout_line.
func (recv *Renderer) GetLayoutLine() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := rendererGetLayoutLineFunction_Set()
	if err == nil {
		ret = rendererGetLayoutLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutLineNewFromNative(ret.Pointer())

	return retGo
}

var rendererGetMatrixFunction *gi.Function
var rendererGetMatrixFunction_Once sync.Once

func rendererGetMatrixFunction_Set() error {
	var err error
	rendererGetMatrixFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererGetMatrixFunction, err = rendererObject.InvokerNew("get_matrix")
	})
	return err
}

// GetMatrix is a representation of the C type pango_renderer_get_matrix.
func (recv *Renderer) GetMatrix() *Matrix {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := rendererGetMatrixFunction_Set()
	if err == nil {
		ret = rendererGetMatrixFunction.Invoke(inArgs[:], nil)
	}

	retGo := MatrixNewFromNative(ret.Pointer())

	return retGo
}

var rendererPartChangedFunction *gi.Function
var rendererPartChangedFunction_Once sync.Once

func rendererPartChangedFunction_Set() error {
	var err error
	rendererPartChangedFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererPartChangedFunction, err = rendererObject.InvokerNew("part_changed")
	})
	return err
}

// PartChanged is a representation of the C type pango_renderer_part_changed.
func (recv *Renderer) PartChanged(part RenderPart) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(part))

	err := rendererPartChangedFunction_Set()
	if err == nil {
		rendererPartChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererSetAlphaFunction *gi.Function
var rendererSetAlphaFunction_Once sync.Once

func rendererSetAlphaFunction_Set() error {
	var err error
	rendererSetAlphaFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererSetAlphaFunction, err = rendererObject.InvokerNew("set_alpha")
	})
	return err
}

// SetAlpha is a representation of the C type pango_renderer_set_alpha.
func (recv *Renderer) SetAlpha(part RenderPart, alpha uint16) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(part))
	inArgs[2].SetUint16(alpha)

	err := rendererSetAlphaFunction_Set()
	if err == nil {
		rendererSetAlphaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererSetColorFunction *gi.Function
var rendererSetColorFunction_Once sync.Once

func rendererSetColorFunction_Set() error {
	var err error
	rendererSetColorFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererSetColorFunction, err = rendererObject.InvokerNew("set_color")
	})
	return err
}

// SetColor is a representation of the C type pango_renderer_set_color.
func (recv *Renderer) SetColor(part RenderPart, color *Color) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(part))
	inArgs[2].SetPointer(color.Native())

	err := rendererSetColorFunction_Set()
	if err == nil {
		rendererSetColorFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rendererSetMatrixFunction *gi.Function
var rendererSetMatrixFunction_Once sync.Once

func rendererSetMatrixFunction_Set() error {
	var err error
	rendererSetMatrixFunction_Once.Do(func() {
		err = rendererObject_Set()
		if err != nil {
			return
		}
		rendererSetMatrixFunction, err = rendererObject.InvokerNew("set_matrix")
	})
	return err
}

// SetMatrix is a representation of the C type pango_renderer_set_matrix.
func (recv *Renderer) SetMatrix(matrix *Matrix) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(matrix.Native())

	err := rendererSetMatrixFunction_Set()
	if err == nil {
		rendererSetMatrixFunction.Invoke(inArgs[:], nil)
	}

	return
}

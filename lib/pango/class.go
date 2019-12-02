// Code generated - DO NOT EDIT.

package pango

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
)

var contextStruct *gi.Struct
var contextStruct_Once sync.Once

func contextStruct_Set() error {
	var err error
	contextStruct_Once.Do(func() {
		contextStruct, err = gi.StructNew("Pango", "Context")
	})
	return err
}

type Context struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_context_new' : return type 'Context' not supported

var contextChangedFunction *gi.Function
var contextChangedFunction_Once sync.Once

func contextChangedFunction_Set() error {
	var err error
	contextChangedFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextChangedFunction, err = contextStruct.InvokerNew("changed")
	})
	return err
}

// Changed is a representation of the C type pango_context_changed.
func (recv *Context) Changed() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := contextChangedFunction_Set()
	if err == nil {
		contextChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_context_get_base_dir' : return type 'Direction' not supported

// UNSUPPORTED : C value 'pango_context_get_base_gravity' : return type 'Gravity' not supported

var contextGetFontDescriptionFunction *gi.Function
var contextGetFontDescriptionFunction_Once sync.Once

func contextGetFontDescriptionFunction_Set() error {
	var err error
	contextGetFontDescriptionFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextGetFontDescriptionFunction, err = contextStruct.InvokerNew("get_font_description")
	})
	return err
}

// GetFontDescription is a representation of the C type pango_context_get_font_description.
func (recv *Context) GetFontDescription() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := contextGetFontDescriptionFunction_Set()
	if err == nil {
		ret = contextGetFontDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_context_get_font_map' : return type 'FontMap' not supported

// UNSUPPORTED : C value 'pango_context_get_gravity' : return type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_context_get_gravity_hint' : return type 'GravityHint' not supported

var contextGetLanguageFunction *gi.Function
var contextGetLanguageFunction_Once sync.Once

func contextGetLanguageFunction_Set() error {
	var err error
	contextGetLanguageFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextGetLanguageFunction, err = contextStruct.InvokerNew("get_language")
	})
	return err
}

// GetLanguage is a representation of the C type pango_context_get_language.
func (recv *Context) GetLanguage() *Language {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := contextGetLanguageFunction_Set()
	if err == nil {
		ret = contextGetLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Language{native: ret.Pointer()}

	return retGo
}

var contextGetMatrixFunction *gi.Function
var contextGetMatrixFunction_Once sync.Once

func contextGetMatrixFunction_Set() error {
	var err error
	contextGetMatrixFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextGetMatrixFunction, err = contextStruct.InvokerNew("get_matrix")
	})
	return err
}

// GetMatrix is a representation of the C type pango_context_get_matrix.
func (recv *Context) GetMatrix() *Matrix {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := contextGetMatrixFunction_Set()
	if err == nil {
		ret = contextGetMatrixFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Matrix{native: ret.Pointer()}

	return retGo
}

var contextGetMetricsFunction *gi.Function
var contextGetMetricsFunction_Once sync.Once

func contextGetMetricsFunction_Set() error {
	var err error
	contextGetMetricsFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextGetMetricsFunction, err = contextStruct.InvokerNew("get_metrics")
	})
	return err
}

// GetMetrics is a representation of the C type pango_context_get_metrics.
func (recv *Context) GetMetrics(desc *FontDescription, language *Language) *FontMetrics {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(desc.native)
	inArgs[2].SetPointer(language.native)

	var ret gi.Argument

	err := contextGetMetricsFunction_Set()
	if err == nil {
		ret = contextGetMetricsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

var contextGetSerialFunction *gi.Function
var contextGetSerialFunction_Once sync.Once

func contextGetSerialFunction_Set() error {
	var err error
	contextGetSerialFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextGetSerialFunction, err = contextStruct.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type pango_context_get_serial.
func (recv *Context) GetSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := contextGetSerialFunction_Set()
	if err == nil {
		ret = contextGetSerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'pango_context_list_families' : parameter 'families' of type 'nil' not supported

// UNSUPPORTED : C value 'pango_context_load_font' : return type 'Font' not supported

// UNSUPPORTED : C value 'pango_context_load_fontset' : return type 'Fontset' not supported

// UNSUPPORTED : C value 'pango_context_set_base_dir' : parameter 'direction' of type 'Direction' not supported

// UNSUPPORTED : C value 'pango_context_set_base_gravity' : parameter 'gravity' of type 'Gravity' not supported

var contextSetFontDescriptionFunction *gi.Function
var contextSetFontDescriptionFunction_Once sync.Once

func contextSetFontDescriptionFunction_Set() error {
	var err error
	contextSetFontDescriptionFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextSetFontDescriptionFunction, err = contextStruct.InvokerNew("set_font_description")
	})
	return err
}

// SetFontDescription is a representation of the C type pango_context_set_font_description.
func (recv *Context) SetFontDescription(desc *FontDescription) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(desc.native)

	err := contextSetFontDescriptionFunction_Set()
	if err == nil {
		contextSetFontDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_context_set_font_map' : parameter 'font_map' of type 'FontMap' not supported

// UNSUPPORTED : C value 'pango_context_set_gravity_hint' : parameter 'hint' of type 'GravityHint' not supported

var contextSetLanguageFunction *gi.Function
var contextSetLanguageFunction_Once sync.Once

func contextSetLanguageFunction_Set() error {
	var err error
	contextSetLanguageFunction_Once.Do(func() {
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextSetLanguageFunction, err = contextStruct.InvokerNew("set_language")
	})
	return err
}

// SetLanguage is a representation of the C type pango_context_set_language.
func (recv *Context) SetLanguage(language *Language) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(language.native)

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
		err = contextStruct_Set()
		if err != nil {
			return
		}
		contextSetMatrixFunction, err = contextStruct.InvokerNew("set_matrix")
	})
	return err
}

// SetMatrix is a representation of the C type pango_context_set_matrix.
func (recv *Context) SetMatrix(matrix *Matrix) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(matrix.native)

	err := contextSetMatrixFunction_Set()
	if err == nil {
		contextSetMatrixFunction.Invoke(inArgs[:], nil)
	}

	return
}

var engineStruct *gi.Struct
var engineStruct_Once sync.Once

func engineStruct_Set() error {
	var err error
	engineStruct_Once.Do(func() {
		engineStruct, err = gi.StructNew("Pango", "Engine")
	})
	return err
}

type Engine struct {
	native uintptr
}

// EngineStruct creates an uninitialised Engine.
func EngineStruct() *Engine {
	err := engineStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Engine{native: engineStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEngine)
	return structGo
}
func finalizeEngine(obj *Engine) {
	engineStruct.Free(obj.native)
}

var engineLangStruct *gi.Struct
var engineLangStruct_Once sync.Once

func engineLangStruct_Set() error {
	var err error
	engineLangStruct_Once.Do(func() {
		engineLangStruct, err = gi.StructNew("Pango", "EngineLang")
	})
	return err
}

type EngineLang struct {
	native uintptr
}

// EngineLangStruct creates an uninitialised EngineLang.
func EngineLangStruct() *EngineLang {
	err := engineLangStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EngineLang{native: engineLangStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEngineLang)
	return structGo
}
func finalizeEngineLang(obj *EngineLang) {
	engineLangStruct.Free(obj.native)
}

var engineShapeStruct *gi.Struct
var engineShapeStruct_Once sync.Once

func engineShapeStruct_Set() error {
	var err error
	engineShapeStruct_Once.Do(func() {
		engineShapeStruct, err = gi.StructNew("Pango", "EngineShape")
	})
	return err
}

type EngineShape struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'Engine'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'Engine'

// EngineShapeStruct creates an uninitialised EngineShape.
func EngineShapeStruct() *EngineShape {
	err := engineShapeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EngineShape{native: engineShapeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEngineShape)
	return structGo
}
func finalizeEngineShape(obj *EngineShape) {
	engineShapeStruct.Free(obj.native)
}

var fontStruct *gi.Struct
var fontStruct_Once sync.Once

func fontStruct_Set() error {
	var err error
	fontStruct_Once.Do(func() {
		fontStruct, err = gi.StructNew("Pango", "Font")
	})
	return err
}

type Font struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

var fontDescribeFunction *gi.Function
var fontDescribeFunction_Once sync.Once

func fontDescribeFunction_Set() error {
	var err error
	fontDescribeFunction_Once.Do(func() {
		err = fontStruct_Set()
		if err != nil {
			return
		}
		fontDescribeFunction, err = fontStruct.InvokerNew("describe")
	})
	return err
}

// Describe is a representation of the C type pango_font_describe.
func (recv *Font) Describe() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescribeFunction_Set()
	if err == nil {
		ret = fontDescribeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var fontDescribeWithAbsoluteSizeFunction *gi.Function
var fontDescribeWithAbsoluteSizeFunction_Once sync.Once

func fontDescribeWithAbsoluteSizeFunction_Set() error {
	var err error
	fontDescribeWithAbsoluteSizeFunction_Once.Do(func() {
		err = fontStruct_Set()
		if err != nil {
			return
		}
		fontDescribeWithAbsoluteSizeFunction, err = fontStruct.InvokerNew("describe_with_absolute_size")
	})
	return err
}

// DescribeWithAbsoluteSize is a representation of the C type pango_font_describe_with_absolute_size.
func (recv *Font) DescribeWithAbsoluteSize() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescribeWithAbsoluteSizeFunction_Set()
	if err == nil {
		ret = fontDescribeWithAbsoluteSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_find_shaper' : return type 'EngineShape' not supported

var fontGetCoverageFunction *gi.Function
var fontGetCoverageFunction_Once sync.Once

func fontGetCoverageFunction_Set() error {
	var err error
	fontGetCoverageFunction_Once.Do(func() {
		err = fontStruct_Set()
		if err != nil {
			return
		}
		fontGetCoverageFunction, err = fontStruct.InvokerNew("get_coverage")
	})
	return err
}

// GetCoverage is a representation of the C type pango_font_get_coverage.
func (recv *Font) GetCoverage(language *Language) *Coverage {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(language.native)

	var ret gi.Argument

	err := fontGetCoverageFunction_Set()
	if err == nil {
		ret = fontGetCoverageFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Coverage{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_font_get_font_map' : return type 'FontMap' not supported

var fontGetGlyphExtentsFunction *gi.Function
var fontGetGlyphExtentsFunction_Once sync.Once

func fontGetGlyphExtentsFunction_Set() error {
	var err error
	fontGetGlyphExtentsFunction_Once.Do(func() {
		err = fontStruct_Set()
		if err != nil {
			return
		}
		fontGetGlyphExtentsFunction, err = fontStruct.InvokerNew("get_glyph_extents")
	})
	return err
}

// GetGlyphExtents is a representation of the C type pango_font_get_glyph_extents.
func (recv *Font) GetGlyphExtents(glyph Glyph) (*Rectangle, *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(uint32(glyph))

	var outArgs [2]gi.Argument

	err := fontGetGlyphExtentsFunction_Set()
	if err == nil {
		fontGetGlyphExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1
}

var fontGetMetricsFunction *gi.Function
var fontGetMetricsFunction_Once sync.Once

func fontGetMetricsFunction_Set() error {
	var err error
	fontGetMetricsFunction_Once.Do(func() {
		err = fontStruct_Set()
		if err != nil {
			return
		}
		fontGetMetricsFunction, err = fontStruct.InvokerNew("get_metrics")
	})
	return err
}

// GetMetrics is a representation of the C type pango_font_get_metrics.
func (recv *Font) GetMetrics(language *Language) *FontMetrics {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(language.native)

	var ret gi.Argument

	err := fontGetMetricsFunction_Set()
	if err == nil {
		ret = fontGetMetricsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

// FontStruct creates an uninitialised Font.
func FontStruct() *Font {
	err := fontStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Font{native: fontStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFont)
	return structGo
}
func finalizeFont(obj *Font) {
	fontStruct.Free(obj.native)
}

var fontFaceStruct *gi.Struct
var fontFaceStruct_Once sync.Once

func fontFaceStruct_Set() error {
	var err error
	fontFaceStruct_Once.Do(func() {
		fontFaceStruct, err = gi.StructNew("Pango", "FontFace")
	})
	return err
}

type FontFace struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

var fontFaceDescribeFunction *gi.Function
var fontFaceDescribeFunction_Once sync.Once

func fontFaceDescribeFunction_Set() error {
	var err error
	fontFaceDescribeFunction_Once.Do(func() {
		err = fontFaceStruct_Set()
		if err != nil {
			return
		}
		fontFaceDescribeFunction, err = fontFaceStruct.InvokerNew("describe")
	})
	return err
}

// Describe is a representation of the C type pango_font_face_describe.
func (recv *FontFace) Describe() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontFaceDescribeFunction_Set()
	if err == nil {
		ret = fontFaceDescribeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var fontFaceGetFaceNameFunction *gi.Function
var fontFaceGetFaceNameFunction_Once sync.Once

func fontFaceGetFaceNameFunction_Set() error {
	var err error
	fontFaceGetFaceNameFunction_Once.Do(func() {
		err = fontFaceStruct_Set()
		if err != nil {
			return
		}
		fontFaceGetFaceNameFunction, err = fontFaceStruct.InvokerNew("get_face_name")
	})
	return err
}

// GetFaceName is a representation of the C type pango_font_face_get_face_name.
func (recv *FontFace) GetFaceName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = fontFaceStruct_Set()
		if err != nil {
			return
		}
		fontFaceIsSynthesizedFunction, err = fontFaceStruct.InvokerNew("is_synthesized")
	})
	return err
}

// IsSynthesized is a representation of the C type pango_font_face_is_synthesized.
func (recv *FontFace) IsSynthesized() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontFaceIsSynthesizedFunction_Set()
	if err == nil {
		ret = fontFaceIsSynthesizedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_face_list_sizes' : parameter 'sizes' of type 'nil' not supported

// FontFaceStruct creates an uninitialised FontFace.
func FontFaceStruct() *FontFace {
	err := fontFaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontFace{native: fontFaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontFace)
	return structGo
}
func finalizeFontFace(obj *FontFace) {
	fontFaceStruct.Free(obj.native)
}

var fontFamilyStruct *gi.Struct
var fontFamilyStruct_Once sync.Once

func fontFamilyStruct_Set() error {
	var err error
	fontFamilyStruct_Once.Do(func() {
		fontFamilyStruct, err = gi.StructNew("Pango", "FontFamily")
	})
	return err
}

type FontFamily struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

var fontFamilyGetNameFunction *gi.Function
var fontFamilyGetNameFunction_Once sync.Once

func fontFamilyGetNameFunction_Set() error {
	var err error
	fontFamilyGetNameFunction_Once.Do(func() {
		err = fontFamilyStruct_Set()
		if err != nil {
			return
		}
		fontFamilyGetNameFunction, err = fontFamilyStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type pango_font_family_get_name.
func (recv *FontFamily) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = fontFamilyStruct_Set()
		if err != nil {
			return
		}
		fontFamilyIsMonospaceFunction, err = fontFamilyStruct.InvokerNew("is_monospace")
	})
	return err
}

// IsMonospace is a representation of the C type pango_font_family_is_monospace.
func (recv *FontFamily) IsMonospace() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontFamilyIsMonospaceFunction_Set()
	if err == nil {
		ret = fontFamilyIsMonospaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'pango_font_family_list_faces' : parameter 'faces' of type 'nil' not supported

// FontFamilyStruct creates an uninitialised FontFamily.
func FontFamilyStruct() *FontFamily {
	err := fontFamilyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontFamily{native: fontFamilyStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontFamily)
	return structGo
}
func finalizeFontFamily(obj *FontFamily) {
	fontFamilyStruct.Free(obj.native)
}

var fontMapStruct *gi.Struct
var fontMapStruct_Once sync.Once

func fontMapStruct_Set() error {
	var err error
	fontMapStruct_Once.Do(func() {
		fontMapStruct, err = gi.StructNew("Pango", "FontMap")
	})
	return err
}

type FontMap struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

var fontMapChangedFunction *gi.Function
var fontMapChangedFunction_Once sync.Once

func fontMapChangedFunction_Set() error {
	var err error
	fontMapChangedFunction_Once.Do(func() {
		err = fontMapStruct_Set()
		if err != nil {
			return
		}
		fontMapChangedFunction, err = fontMapStruct.InvokerNew("changed")
	})
	return err
}

// Changed is a representation of the C type pango_font_map_changed.
func (recv *FontMap) Changed() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontMapChangedFunction_Set()
	if err == nil {
		fontMapChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_font_map_create_context' : return type 'Context' not supported

var fontMapGetSerialFunction *gi.Function
var fontMapGetSerialFunction_Once sync.Once

func fontMapGetSerialFunction_Set() error {
	var err error
	fontMapGetSerialFunction_Once.Do(func() {
		err = fontMapStruct_Set()
		if err != nil {
			return
		}
		fontMapGetSerialFunction, err = fontMapStruct.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type pango_font_map_get_serial.
func (recv *FontMap) GetSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = fontMapStruct_Set()
		if err != nil {
			return
		}
		fontMapGetShapeEngineTypeFunction, err = fontMapStruct.InvokerNew("get_shape_engine_type")
	})
	return err
}

// GetShapeEngineType is a representation of the C type pango_font_map_get_shape_engine_type.
func (recv *FontMap) GetShapeEngineType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMapGetShapeEngineTypeFunction_Set()
	if err == nil {
		ret = fontMapGetShapeEngineTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_map_list_families' : parameter 'families' of type 'nil' not supported

// UNSUPPORTED : C value 'pango_font_map_load_font' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'pango_font_map_load_fontset' : parameter 'context' of type 'Context' not supported

// FontMapStruct creates an uninitialised FontMap.
func FontMapStruct() *FontMap {
	err := fontMapStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontMap{native: fontMapStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontMap)
	return structGo
}
func finalizeFontMap(obj *FontMap) {
	fontMapStruct.Free(obj.native)
}

var fontsetStruct *gi.Struct
var fontsetStruct_Once sync.Once

func fontsetStruct_Set() error {
	var err error
	fontsetStruct_Once.Do(func() {
		fontsetStruct, err = gi.StructNew("Pango", "Fontset")
	})
	return err
}

type Fontset struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'pango_fontset_foreach' : parameter 'func' of type 'FontsetForeachFunc' not supported

// UNSUPPORTED : C value 'pango_fontset_get_font' : return type 'Font' not supported

var fontsetGetMetricsFunction *gi.Function
var fontsetGetMetricsFunction_Once sync.Once

func fontsetGetMetricsFunction_Set() error {
	var err error
	fontsetGetMetricsFunction_Once.Do(func() {
		err = fontsetStruct_Set()
		if err != nil {
			return
		}
		fontsetGetMetricsFunction, err = fontsetStruct.InvokerNew("get_metrics")
	})
	return err
}

// GetMetrics is a representation of the C type pango_fontset_get_metrics.
func (recv *Fontset) GetMetrics() *FontMetrics {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontsetGetMetricsFunction_Set()
	if err == nil {
		ret = fontsetGetMetricsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontMetrics{native: ret.Pointer()}

	return retGo
}

// FontsetStruct creates an uninitialised Fontset.
func FontsetStruct() *Fontset {
	err := fontsetStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Fontset{native: fontsetStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontset)
	return structGo
}
func finalizeFontset(obj *Fontset) {
	fontsetStruct.Free(obj.native)
}

var fontsetSimpleStruct *gi.Struct
var fontsetSimpleStruct_Once sync.Once

func fontsetSimpleStruct_Set() error {
	var err error
	fontsetSimpleStruct_Once.Do(func() {
		fontsetSimpleStruct, err = gi.StructNew("Pango", "FontsetSimple")
	})
	return err
}

type FontsetSimple struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_fontset_simple_new' : return type 'FontsetSimple' not supported

// UNSUPPORTED : C value 'pango_fontset_simple_append' : parameter 'font' of type 'Font' not supported

var fontsetSimpleSizeFunction *gi.Function
var fontsetSimpleSizeFunction_Once sync.Once

func fontsetSimpleSizeFunction_Set() error {
	var err error
	fontsetSimpleSizeFunction_Once.Do(func() {
		err = fontsetSimpleStruct_Set()
		if err != nil {
			return
		}
		fontsetSimpleSizeFunction, err = fontsetSimpleStruct.InvokerNew("size")
	})
	return err
}

// Size is a representation of the C type pango_fontset_simple_size.
func (recv *FontsetSimple) Size() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontsetSimpleSizeFunction_Set()
	if err == nil {
		ret = fontsetSimpleSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutStruct *gi.Struct
var layoutStruct_Once sync.Once

func layoutStruct_Set() error {
	var err error
	layoutStruct_Once.Do(func() {
		layoutStruct, err = gi.StructNew("Pango", "Layout")
	})
	return err
}

type Layout struct {
	native uintptr
}

// UNSUPPORTED : C value 'pango_layout_new' : parameter 'context' of type 'Context' not supported

var layoutContextChangedFunction *gi.Function
var layoutContextChangedFunction_Once sync.Once

func layoutContextChangedFunction_Set() error {
	var err error
	layoutContextChangedFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutContextChangedFunction, err = layoutStruct.InvokerNew("context_changed")
	})
	return err
}

// ContextChanged is a representation of the C type pango_layout_context_changed.
func (recv *Layout) ContextChanged() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := layoutContextChangedFunction_Set()
	if err == nil {
		layoutContextChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_layout_copy' : return type 'Layout' not supported

// UNSUPPORTED : C value 'pango_layout_get_alignment' : return type 'Alignment' not supported

var layoutGetAttributesFunction *gi.Function
var layoutGetAttributesFunction_Once sync.Once

func layoutGetAttributesFunction_Set() error {
	var err error
	layoutGetAttributesFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetAttributesFunction, err = layoutStruct.InvokerNew("get_attributes")
	})
	return err
}

// GetAttributes is a representation of the C type pango_layout_get_attributes.
func (recv *Layout) GetAttributes() *AttrList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutGetAttributesFunction_Set()
	if err == nil {
		ret = layoutGetAttributesFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AttrList{native: ret.Pointer()}

	return retGo
}

var layoutGetAutoDirFunction *gi.Function
var layoutGetAutoDirFunction_Once sync.Once

func layoutGetAutoDirFunction_Set() error {
	var err error
	layoutGetAutoDirFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetAutoDirFunction, err = layoutStruct.InvokerNew("get_auto_dir")
	})
	return err
}

// GetAutoDir is a representation of the C type pango_layout_get_auto_dir.
func (recv *Layout) GetAutoDir() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetBaselineFunction, err = layoutStruct.InvokerNew("get_baseline")
	})
	return err
}

// GetBaseline is a representation of the C type pango_layout_get_baseline.
func (recv *Layout) GetBaseline() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetCharacterCountFunction, err = layoutStruct.InvokerNew("get_character_count")
	})
	return err
}

// GetCharacterCount is a representation of the C type pango_layout_get_character_count.
func (recv *Layout) GetCharacterCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutGetCharacterCountFunction_Set()
	if err == nil {
		ret = layoutGetCharacterCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_get_context' : return type 'Context' not supported

var layoutGetCursorPosFunction *gi.Function
var layoutGetCursorPosFunction_Once sync.Once

func layoutGetCursorPosFunction_Set() error {
	var err error
	layoutGetCursorPosFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetCursorPosFunction, err = layoutStruct.InvokerNew("get_cursor_pos")
	})
	return err
}

// GetCursorPos is a representation of the C type pango_layout_get_cursor_pos.
func (recv *Layout) GetCursorPos(index int32) (*Rectangle, *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	var outArgs [2]gi.Argument

	err := layoutGetCursorPosFunction_Set()
	if err == nil {
		layoutGetCursorPosFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1
}

// UNSUPPORTED : C value 'pango_layout_get_ellipsize' : return type 'EllipsizeMode' not supported

var layoutGetExtentsFunction *gi.Function
var layoutGetExtentsFunction_Once sync.Once

func layoutGetExtentsFunction_Set() error {
	var err error
	layoutGetExtentsFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetExtentsFunction, err = layoutStruct.InvokerNew("get_extents")
	})
	return err
}

// GetExtents is a representation of the C type pango_layout_get_extents.
func (recv *Layout) GetExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutGetExtentsFunction_Set()
	if err == nil {
		layoutGetExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1
}

var layoutGetFontDescriptionFunction *gi.Function
var layoutGetFontDescriptionFunction_Once sync.Once

func layoutGetFontDescriptionFunction_Set() error {
	var err error
	layoutGetFontDescriptionFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetFontDescriptionFunction, err = layoutStruct.InvokerNew("get_font_description")
	})
	return err
}

// GetFontDescription is a representation of the C type pango_layout_get_font_description.
func (recv *Layout) GetFontDescription() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutGetFontDescriptionFunction_Set()
	if err == nil {
		ret = layoutGetFontDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var layoutGetHeightFunction *gi.Function
var layoutGetHeightFunction_Once sync.Once

func layoutGetHeightFunction_Set() error {
	var err error
	layoutGetHeightFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetHeightFunction, err = layoutStruct.InvokerNew("get_height")
	})
	return err
}

// GetHeight is a representation of the C type pango_layout_get_height.
func (recv *Layout) GetHeight() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetIndentFunction, err = layoutStruct.InvokerNew("get_indent")
	})
	return err
}

// GetIndent is a representation of the C type pango_layout_get_indent.
func (recv *Layout) GetIndent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetIterFunction, err = layoutStruct.InvokerNew("get_iter")
	})
	return err
}

// GetIter is a representation of the C type pango_layout_get_iter.
func (recv *Layout) GetIter() *LayoutIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutGetIterFunction_Set()
	if err == nil {
		ret = layoutGetIterFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutIter{native: ret.Pointer()}

	return retGo
}

var layoutGetJustifyFunction *gi.Function
var layoutGetJustifyFunction_Once sync.Once

func layoutGetJustifyFunction_Set() error {
	var err error
	layoutGetJustifyFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetJustifyFunction, err = layoutStruct.InvokerNew("get_justify")
	})
	return err
}

// GetJustify is a representation of the C type pango_layout_get_justify.
func (recv *Layout) GetJustify() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetLineFunction, err = layoutStruct.InvokerNew("get_line")
	})
	return err
}

// GetLine is a representation of the C type pango_layout_get_line.
func (recv *Layout) GetLine(line int32) *LayoutLine {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(line)

	var ret gi.Argument

	err := layoutGetLineFunction_Set()
	if err == nil {
		ret = layoutGetLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var layoutGetLineCountFunction *gi.Function
var layoutGetLineCountFunction_Once sync.Once

func layoutGetLineCountFunction_Set() error {
	var err error
	layoutGetLineCountFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetLineCountFunction, err = layoutStruct.InvokerNew("get_line_count")
	})
	return err
}

// GetLineCount is a representation of the C type pango_layout_get_line_count.
func (recv *Layout) GetLineCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetLineReadonlyFunction, err = layoutStruct.InvokerNew("get_line_readonly")
	})
	return err
}

// GetLineReadonly is a representation of the C type pango_layout_get_line_readonly.
func (recv *Layout) GetLineReadonly(line int32) *LayoutLine {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(line)

	var ret gi.Argument

	err := layoutGetLineReadonlyFunction_Set()
	if err == nil {
		ret = layoutGetLineReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetLogAttrsReadonlyFunction, err = layoutStruct.InvokerNew("get_log_attrs_readonly")
	})
	return err
}

// GetLogAttrsReadonly is a representation of the C type pango_layout_get_log_attrs_readonly.
func (recv *Layout) GetLogAttrsReadonly() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetPixelExtentsFunction, err = layoutStruct.InvokerNew("get_pixel_extents")
	})
	return err
}

// GetPixelExtents is a representation of the C type pango_layout_get_pixel_extents.
func (recv *Layout) GetPixelExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutGetPixelExtentsFunction_Set()
	if err == nil {
		layoutGetPixelExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}
	out1 := &Rectangle{native: outArgs[1].Pointer()}

	return out0, out1
}

var layoutGetPixelSizeFunction *gi.Function
var layoutGetPixelSizeFunction_Once sync.Once

func layoutGetPixelSizeFunction_Set() error {
	var err error
	layoutGetPixelSizeFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetPixelSizeFunction, err = layoutStruct.InvokerNew("get_pixel_size")
	})
	return err
}

// GetPixelSize is a representation of the C type pango_layout_get_pixel_size.
func (recv *Layout) GetPixelSize() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetSerialFunction, err = layoutStruct.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type pango_layout_get_serial.
func (recv *Layout) GetSerial() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetSingleParagraphModeFunction, err = layoutStruct.InvokerNew("get_single_paragraph_mode")
	})
	return err
}

// GetSingleParagraphMode is a representation of the C type pango_layout_get_single_paragraph_mode.
func (recv *Layout) GetSingleParagraphMode() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetSizeFunction, err = layoutStruct.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type pango_layout_get_size.
func (recv *Layout) GetSize() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetSpacingFunction, err = layoutStruct.InvokerNew("get_spacing")
	})
	return err
}

// GetSpacing is a representation of the C type pango_layout_get_spacing.
func (recv *Layout) GetSpacing() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetTabsFunction, err = layoutStruct.InvokerNew("get_tabs")
	})
	return err
}

// GetTabs is a representation of the C type pango_layout_get_tabs.
func (recv *Layout) GetTabs() *TabArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutGetTabsFunction_Set()
	if err == nil {
		ret = layoutGetTabsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TabArray{native: ret.Pointer()}

	return retGo
}

var layoutGetTextFunction *gi.Function
var layoutGetTextFunction_Once sync.Once

func layoutGetTextFunction_Set() error {
	var err error
	layoutGetTextFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetTextFunction, err = layoutStruct.InvokerNew("get_text")
	})
	return err
}

// GetText is a representation of the C type pango_layout_get_text.
func (recv *Layout) GetText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetUnknownGlyphsCountFunction, err = layoutStruct.InvokerNew("get_unknown_glyphs_count")
	})
	return err
}

// GetUnknownGlyphsCount is a representation of the C type pango_layout_get_unknown_glyphs_count.
func (recv *Layout) GetUnknownGlyphsCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutGetWidthFunction, err = layoutStruct.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type pango_layout_get_width.
func (recv *Layout) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutGetWidthFunction_Set()
	if err == nil {
		ret = layoutGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'pango_layout_get_wrap' : return type 'WrapMode' not supported

var layoutIndexToLineXFunction *gi.Function
var layoutIndexToLineXFunction_Once sync.Once

func layoutIndexToLineXFunction_Set() error {
	var err error
	layoutIndexToLineXFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutIndexToLineXFunction, err = layoutStruct.InvokerNew("index_to_line_x")
	})
	return err
}

// IndexToLineX is a representation of the C type pango_layout_index_to_line_x.
func (recv *Layout) IndexToLineX(index int32, trailing bool) (int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutIndexToPosFunction, err = layoutStruct.InvokerNew("index_to_pos")
	})
	return err
}

// IndexToPos is a representation of the C type pango_layout_index_to_pos.
func (recv *Layout) IndexToPos(index int32) *Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	var outArgs [1]gi.Argument

	err := layoutIndexToPosFunction_Set()
	if err == nil {
		layoutIndexToPosFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{native: outArgs[0].Pointer()}

	return out0
}

var layoutIsEllipsizedFunction *gi.Function
var layoutIsEllipsizedFunction_Once sync.Once

func layoutIsEllipsizedFunction_Set() error {
	var err error
	layoutIsEllipsizedFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutIsEllipsizedFunction, err = layoutStruct.InvokerNew("is_ellipsized")
	})
	return err
}

// IsEllipsized is a representation of the C type pango_layout_is_ellipsized.
func (recv *Layout) IsEllipsized() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutIsWrappedFunction, err = layoutStruct.InvokerNew("is_wrapped")
	})
	return err
}

// IsWrapped is a representation of the C type pango_layout_is_wrapped.
func (recv *Layout) IsWrapped() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutMoveCursorVisuallyFunction, err = layoutStruct.InvokerNew("move_cursor_visually")
	})
	return err
}

// MoveCursorVisually is a representation of the C type pango_layout_move_cursor_visually.
func (recv *Layout) MoveCursorVisually(strong bool, oldIndex int32, oldTrailing int32, direction int32) (int32, int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
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

// UNSUPPORTED : C value 'pango_layout_set_alignment' : parameter 'alignment' of type 'Alignment' not supported

var layoutSetAttributesFunction *gi.Function
var layoutSetAttributesFunction_Once sync.Once

func layoutSetAttributesFunction_Set() error {
	var err error
	layoutSetAttributesFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetAttributesFunction, err = layoutStruct.InvokerNew("set_attributes")
	})
	return err
}

// SetAttributes is a representation of the C type pango_layout_set_attributes.
func (recv *Layout) SetAttributes(attrs *AttrList) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attrs.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetAutoDirFunction, err = layoutStruct.InvokerNew("set_auto_dir")
	})
	return err
}

// SetAutoDir is a representation of the C type pango_layout_set_auto_dir.
func (recv *Layout) SetAutoDir(autoDir bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(autoDir)

	err := layoutSetAutoDirFunction_Set()
	if err == nil {
		layoutSetAutoDirFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_layout_set_ellipsize' : parameter 'ellipsize' of type 'EllipsizeMode' not supported

var layoutSetFontDescriptionFunction *gi.Function
var layoutSetFontDescriptionFunction_Once sync.Once

func layoutSetFontDescriptionFunction_Set() error {
	var err error
	layoutSetFontDescriptionFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetFontDescriptionFunction, err = layoutStruct.InvokerNew("set_font_description")
	})
	return err
}

// SetFontDescription is a representation of the C type pango_layout_set_font_description.
func (recv *Layout) SetFontDescription(desc *FontDescription) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(desc.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetHeightFunction, err = layoutStruct.InvokerNew("set_height")
	})
	return err
}

// SetHeight is a representation of the C type pango_layout_set_height.
func (recv *Layout) SetHeight(height int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetIndentFunction, err = layoutStruct.InvokerNew("set_indent")
	})
	return err
}

// SetIndent is a representation of the C type pango_layout_set_indent.
func (recv *Layout) SetIndent(indent int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetJustifyFunction, err = layoutStruct.InvokerNew("set_justify")
	})
	return err
}

// SetJustify is a representation of the C type pango_layout_set_justify.
func (recv *Layout) SetJustify(justify bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetMarkupFunction, err = layoutStruct.InvokerNew("set_markup")
	})
	return err
}

// SetMarkup is a representation of the C type pango_layout_set_markup.
func (recv *Layout) SetMarkup(markup string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetSingleParagraphModeFunction, err = layoutStruct.InvokerNew("set_single_paragraph_mode")
	})
	return err
}

// SetSingleParagraphMode is a representation of the C type pango_layout_set_single_paragraph_mode.
func (recv *Layout) SetSingleParagraphMode(setting bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetSpacingFunction, err = layoutStruct.InvokerNew("set_spacing")
	})
	return err
}

// SetSpacing is a representation of the C type pango_layout_set_spacing.
func (recv *Layout) SetSpacing(spacing int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetTabsFunction, err = layoutStruct.InvokerNew("set_tabs")
	})
	return err
}

// SetTabs is a representation of the C type pango_layout_set_tabs.
func (recv *Layout) SetTabs(tabs *TabArray) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tabs.native)

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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetTextFunction, err = layoutStruct.InvokerNew("set_text")
	})
	return err
}

// SetText is a representation of the C type pango_layout_set_text.
func (recv *Layout) SetText(text string, length int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutSetWidthFunction, err = layoutStruct.InvokerNew("set_width")
	})
	return err
}

// SetWidth is a representation of the C type pango_layout_set_width.
func (recv *Layout) SetWidth(width int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(width)

	err := layoutSetWidthFunction_Set()
	if err == nil {
		layoutSetWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_layout_set_wrap' : parameter 'wrap' of type 'WrapMode' not supported

var layoutXyToIndexFunction *gi.Function
var layoutXyToIndexFunction_Once sync.Once

func layoutXyToIndexFunction_Set() error {
	var err error
	layoutXyToIndexFunction_Once.Do(func() {
		err = layoutStruct_Set()
		if err != nil {
			return
		}
		layoutXyToIndexFunction, err = layoutStruct.InvokerNew("xy_to_index")
	})
	return err
}

// XyToIndex is a representation of the C type pango_layout_xy_to_index.
func (recv *Layout) XyToIndex(x int32, y int32) (bool, int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
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

var rendererStruct *gi.Struct
var rendererStruct_Once sync.Once

func rendererStruct_Set() error {
	var err error
	rendererStruct_Once.Do(func() {
		rendererStruct, err = gi.StructNew("Pango", "Renderer")
	})
	return err
}

type Renderer struct {
	native uintptr
}

// FieldMatrix returns the C field 'matrix'.
func (recv *Renderer) FieldMatrix() *Matrix {
	argValue := gi.FieldGet(rendererStruct, recv.native, "matrix")
	value := &Matrix{native: argValue.Pointer()}
	return value
}

// SetFieldMatrix sets the value of the C field 'matrix'.
func (recv *Renderer) SetFieldMatrix(value *Matrix) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(rendererStruct, recv.native, "matrix", argValue)
}

var rendererActivateFunction *gi.Function
var rendererActivateFunction_Once sync.Once

func rendererActivateFunction_Set() error {
	var err error
	rendererActivateFunction_Once.Do(func() {
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererActivateFunction, err = rendererStruct.InvokerNew("activate")
	})
	return err
}

// Activate is a representation of the C type pango_renderer_activate.
func (recv *Renderer) Activate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererDeactivateFunction, err = rendererStruct.InvokerNew("deactivate")
	})
	return err
}

// Deactivate is a representation of the C type pango_renderer_deactivate.
func (recv *Renderer) Deactivate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

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
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererDrawErrorUnderlineFunction, err = rendererStruct.InvokerNew("draw_error_underline")
	})
	return err
}

// DrawErrorUnderline is a representation of the C type pango_renderer_draw_error_underline.
func (recv *Renderer) DrawErrorUnderline(x int32, y int32, width int32, height int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
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

// UNSUPPORTED : C value 'pango_renderer_draw_glyph' : parameter 'font' of type 'Font' not supported

var rendererDrawGlyphItemFunction *gi.Function
var rendererDrawGlyphItemFunction_Once sync.Once

func rendererDrawGlyphItemFunction_Set() error {
	var err error
	rendererDrawGlyphItemFunction_Once.Do(func() {
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererDrawGlyphItemFunction, err = rendererStruct.InvokerNew("draw_glyph_item")
	})
	return err
}

// DrawGlyphItem is a representation of the C type pango_renderer_draw_glyph_item.
func (recv *Renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x int32, y int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetPointer(glyphItem.native)
	inArgs[3].SetInt32(x)
	inArgs[4].SetInt32(y)

	err := rendererDrawGlyphItemFunction_Set()
	if err == nil {
		rendererDrawGlyphItemFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_renderer_draw_glyphs' : parameter 'font' of type 'Font' not supported

// UNSUPPORTED : C value 'pango_renderer_draw_layout' : parameter 'layout' of type 'Layout' not supported

var rendererDrawLayoutLineFunction *gi.Function
var rendererDrawLayoutLineFunction_Once sync.Once

func rendererDrawLayoutLineFunction_Set() error {
	var err error
	rendererDrawLayoutLineFunction_Once.Do(func() {
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererDrawLayoutLineFunction, err = rendererStruct.InvokerNew("draw_layout_line")
	})
	return err
}

// DrawLayoutLine is a representation of the C type pango_renderer_draw_layout_line.
func (recv *Renderer) DrawLayoutLine(line *LayoutLine, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(line.native)
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := rendererDrawLayoutLineFunction_Set()
	if err == nil {
		rendererDrawLayoutLineFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_renderer_draw_rectangle' : parameter 'part' of type 'RenderPart' not supported

// UNSUPPORTED : C value 'pango_renderer_draw_trapezoid' : parameter 'part' of type 'RenderPart' not supported

// UNSUPPORTED : C value 'pango_renderer_get_alpha' : parameter 'part' of type 'RenderPart' not supported

// UNSUPPORTED : C value 'pango_renderer_get_color' : parameter 'part' of type 'RenderPart' not supported

// UNSUPPORTED : C value 'pango_renderer_get_layout' : return type 'Layout' not supported

var rendererGetLayoutLineFunction *gi.Function
var rendererGetLayoutLineFunction_Once sync.Once

func rendererGetLayoutLineFunction_Set() error {
	var err error
	rendererGetLayoutLineFunction_Once.Do(func() {
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererGetLayoutLineFunction, err = rendererStruct.InvokerNew("get_layout_line")
	})
	return err
}

// GetLayoutLine is a representation of the C type pango_renderer_get_layout_line.
func (recv *Renderer) GetLayoutLine() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rendererGetLayoutLineFunction_Set()
	if err == nil {
		ret = rendererGetLayoutLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := &LayoutLine{native: ret.Pointer()}

	return retGo
}

var rendererGetMatrixFunction *gi.Function
var rendererGetMatrixFunction_Once sync.Once

func rendererGetMatrixFunction_Set() error {
	var err error
	rendererGetMatrixFunction_Once.Do(func() {
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererGetMatrixFunction, err = rendererStruct.InvokerNew("get_matrix")
	})
	return err
}

// GetMatrix is a representation of the C type pango_renderer_get_matrix.
func (recv *Renderer) GetMatrix() *Matrix {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rendererGetMatrixFunction_Set()
	if err == nil {
		ret = rendererGetMatrixFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Matrix{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_renderer_part_changed' : parameter 'part' of type 'RenderPart' not supported

// UNSUPPORTED : C value 'pango_renderer_set_alpha' : parameter 'part' of type 'RenderPart' not supported

// UNSUPPORTED : C value 'pango_renderer_set_color' : parameter 'part' of type 'RenderPart' not supported

var rendererSetMatrixFunction *gi.Function
var rendererSetMatrixFunction_Once sync.Once

func rendererSetMatrixFunction_Set() error {
	var err error
	rendererSetMatrixFunction_Once.Do(func() {
		err = rendererStruct_Set()
		if err != nil {
			return
		}
		rendererSetMatrixFunction, err = rendererStruct.InvokerNew("set_matrix")
	})
	return err
}

// SetMatrix is a representation of the C type pango_renderer_set_matrix.
func (recv *Renderer) SetMatrix(matrix *Matrix) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(matrix.native)

	err := rendererSetMatrixFunction_Set()
	if err == nil {
		rendererSetMatrixFunction.Invoke(inArgs[:], nil)
	}

	return
}

// RendererStruct creates an uninitialised Renderer.
func RendererStruct() *Renderer {
	err := rendererStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Renderer{native: rendererStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRenderer)
	return structGo
}
func finalizeRenderer(obj *Renderer) {
	rendererStruct.Free(obj.native)
}

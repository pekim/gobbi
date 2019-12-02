// Code generated - DO NOT EDIT.

package cairo

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
		contextStruct, err = gi.StructNew("cairo", "Context")
	})
	return err
}

type Context struct {
	Native uintptr
}

// ContextStruct creates an uninitialised Context.
func ContextStruct() *Context {
	err := contextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Context{}
	structGo.Native = contextStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeContext)
	return structGo
}
func finalizeContext(obj *Context) {
	contextStruct.Free(obj.Native)
}

var deviceStruct *gi.Struct
var deviceStruct_Once sync.Once

func deviceStruct_Set() error {
	var err error
	deviceStruct_Once.Do(func() {
		deviceStruct, err = gi.StructNew("cairo", "Device")
	})
	return err
}

type Device struct {
	Native uintptr
}

// DeviceStruct creates an uninitialised Device.
func DeviceStruct() *Device {
	err := deviceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Device{}
	structGo.Native = deviceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDevice)
	return structGo
}
func finalizeDevice(obj *Device) {
	deviceStruct.Free(obj.Native)
}

var surfaceStruct *gi.Struct
var surfaceStruct_Once sync.Once

func surfaceStruct_Set() error {
	var err error
	surfaceStruct_Once.Do(func() {
		surfaceStruct, err = gi.StructNew("cairo", "Surface")
	})
	return err
}

type Surface struct {
	Native uintptr
}

// SurfaceStruct creates an uninitialised Surface.
func SurfaceStruct() *Surface {
	err := surfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Surface{}
	structGo.Native = surfaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSurface)
	return structGo
}
func finalizeSurface(obj *Surface) {
	surfaceStruct.Free(obj.Native)
}

var matrixStruct *gi.Struct
var matrixStruct_Once sync.Once

func matrixStruct_Set() error {
	var err error
	matrixStruct_Once.Do(func() {
		matrixStruct, err = gi.StructNew("cairo", "Matrix")
	})
	return err
}

type Matrix struct {
	Native uintptr
}

// MatrixStruct creates an uninitialised Matrix.
func MatrixStruct() *Matrix {
	err := matrixStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Matrix{}
	structGo.Native = matrixStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMatrix)
	return structGo
}
func finalizeMatrix(obj *Matrix) {
	matrixStruct.Free(obj.Native)
}

var patternStruct *gi.Struct
var patternStruct_Once sync.Once

func patternStruct_Set() error {
	var err error
	patternStruct_Once.Do(func() {
		patternStruct, err = gi.StructNew("cairo", "Pattern")
	})
	return err
}

type Pattern struct {
	Native uintptr
}

// PatternStruct creates an uninitialised Pattern.
func PatternStruct() *Pattern {
	err := patternStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Pattern{}
	structGo.Native = patternStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizePattern)
	return structGo
}
func finalizePattern(obj *Pattern) {
	patternStruct.Free(obj.Native)
}

var regionStruct *gi.Struct
var regionStruct_Once sync.Once

func regionStruct_Set() error {
	var err error
	regionStruct_Once.Do(func() {
		regionStruct, err = gi.StructNew("cairo", "Region")
	})
	return err
}

type Region struct {
	Native uintptr
}

// RegionStruct creates an uninitialised Region.
func RegionStruct() *Region {
	err := regionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Region{}
	structGo.Native = regionStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeRegion)
	return structGo
}
func finalizeRegion(obj *Region) {
	regionStruct.Free(obj.Native)
}

var fontOptionsStruct *gi.Struct
var fontOptionsStruct_Once sync.Once

func fontOptionsStruct_Set() error {
	var err error
	fontOptionsStruct_Once.Do(func() {
		fontOptionsStruct, err = gi.StructNew("cairo", "FontOptions")
	})
	return err
}

type FontOptions struct {
	Native uintptr
}

// FontOptionsStruct creates an uninitialised FontOptions.
func FontOptionsStruct() *FontOptions {
	err := fontOptionsStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontOptions{}
	structGo.Native = fontOptionsStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFontOptions)
	return structGo
}
func finalizeFontOptions(obj *FontOptions) {
	fontOptionsStruct.Free(obj.Native)
}

var fontFaceStruct *gi.Struct
var fontFaceStruct_Once sync.Once

func fontFaceStruct_Set() error {
	var err error
	fontFaceStruct_Once.Do(func() {
		fontFaceStruct, err = gi.StructNew("cairo", "FontFace")
	})
	return err
}

type FontFace struct {
	Native uintptr
}

// FontFaceStruct creates an uninitialised FontFace.
func FontFaceStruct() *FontFace {
	err := fontFaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontFace{}
	structGo.Native = fontFaceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFontFace)
	return structGo
}
func finalizeFontFace(obj *FontFace) {
	fontFaceStruct.Free(obj.Native)
}

var scaledFontStruct *gi.Struct
var scaledFontStruct_Once sync.Once

func scaledFontStruct_Set() error {
	var err error
	scaledFontStruct_Once.Do(func() {
		scaledFontStruct, err = gi.StructNew("cairo", "ScaledFont")
	})
	return err
}

type ScaledFont struct {
	Native uintptr
}

// ScaledFontStruct creates an uninitialised ScaledFont.
func ScaledFontStruct() *ScaledFont {
	err := scaledFontStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaledFont{}
	structGo.Native = scaledFontStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeScaledFont)
	return structGo
}
func finalizeScaledFont(obj *ScaledFont) {
	scaledFontStruct.Free(obj.Native)
}

var pathStruct *gi.Struct
var pathStruct_Once sync.Once

func pathStruct_Set() error {
	var err error
	pathStruct_Once.Do(func() {
		pathStruct, err = gi.StructNew("cairo", "Path")
	})
	return err
}

type Path struct {
	Native uintptr
}

// PathStruct creates an uninitialised Path.
func PathStruct() *Path {
	err := pathStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Path{}
	structGo.Native = pathStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizePath)
	return structGo
}
func finalizePath(obj *Path) {
	pathStruct.Free(obj.Native)
}

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() error {
	var err error
	rectangleStruct_Once.Do(func() {
		rectangleStruct, err = gi.StructNew("cairo", "Rectangle")
	})
	return err
}

type Rectangle struct {
	Native uintptr
}

// FieldX returns the C field 'x'.
func (recv *Rectangle) FieldX() float64 {
	argValue := gi.FieldGet(rectangleStruct, recv.Native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Rectangle) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(rectangleStruct, recv.Native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Rectangle) FieldY() float64 {
	argValue := gi.FieldGet(rectangleStruct, recv.Native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Rectangle) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(rectangleStruct, recv.Native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *Rectangle) FieldWidth() float64 {
	argValue := gi.FieldGet(rectangleStruct, recv.Native, "width")
	value := argValue.Float64()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *Rectangle) SetFieldWidth(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(rectangleStruct, recv.Native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *Rectangle) FieldHeight() float64 {
	argValue := gi.FieldGet(rectangleStruct, recv.Native, "height")
	value := argValue.Float64()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *Rectangle) SetFieldHeight(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(rectangleStruct, recv.Native, "height", argValue)
}

// RectangleStruct creates an uninitialised Rectangle.
func RectangleStruct() *Rectangle {
	err := rectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Rectangle{}
	structGo.Native = rectangleStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeRectangle)
	return structGo
}
func finalizeRectangle(obj *Rectangle) {
	rectangleStruct.Free(obj.Native)
}

var rectangleIntStruct *gi.Struct
var rectangleIntStruct_Once sync.Once

func rectangleIntStruct_Set() error {
	var err error
	rectangleIntStruct_Once.Do(func() {
		rectangleIntStruct, err = gi.StructNew("cairo", "RectangleInt")
	})
	return err
}

type RectangleInt struct {
	Native uintptr
}

// FieldX returns the C field 'x'.
func (recv *RectangleInt) FieldX() int32 {
	argValue := gi.FieldGet(rectangleIntStruct, recv.Native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *RectangleInt) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(rectangleIntStruct, recv.Native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *RectangleInt) FieldY() int32 {
	argValue := gi.FieldGet(rectangleIntStruct, recv.Native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *RectangleInt) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(rectangleIntStruct, recv.Native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *RectangleInt) FieldWidth() int32 {
	argValue := gi.FieldGet(rectangleIntStruct, recv.Native, "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *RectangleInt) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(rectangleIntStruct, recv.Native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *RectangleInt) FieldHeight() int32 {
	argValue := gi.FieldGet(rectangleIntStruct, recv.Native, "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *RectangleInt) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(rectangleIntStruct, recv.Native, "height", argValue)
}

// RectangleIntStruct creates an uninitialised RectangleInt.
func RectangleIntStruct() *RectangleInt {
	err := rectangleIntStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RectangleInt{}
	structGo.Native = rectangleIntStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeRectangleInt)
	return structGo
}
func finalizeRectangleInt(obj *RectangleInt) {
	rectangleIntStruct.Free(obj.Native)
}

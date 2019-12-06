// Code generated - DO NOT EDIT.

package cairo

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
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
	native unsafe.Pointer
}

func ContextNewFromNative(native unsafe.Pointer) *Context {
	return &Context{native: native}
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

// ContextStruct creates an uninitialised Context.
func ContextStruct() *Context {
	err := contextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextNewFromNative(contextStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContext)
	return structGo
}
func finalizeContext(obj *Context) {
	contextStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func DeviceNewFromNative(native unsafe.Pointer) *Device {
	return &Device{native: native}
}

/*
CastToDevice down casts any arbitrary Object to Device.
Exercise care, as this is a potentially dangerous function
if the Object is not a Device.
*/
func (recv *Device) CastToDevice(object *gobject.Object) *Device {
	return DeviceNewFromNative(object.Native())
}

func (recv *Device) Native() unsafe.Pointer {
	return recv.native
}

// DeviceStruct creates an uninitialised Device.
func DeviceStruct() *Device {
	err := deviceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DeviceNewFromNative(deviceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDevice)
	return structGo
}
func finalizeDevice(obj *Device) {
	deviceStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func SurfaceNewFromNative(native unsafe.Pointer) *Surface {
	return &Surface{native: native}
}

/*
CastToSurface down casts any arbitrary Object to Surface.
Exercise care, as this is a potentially dangerous function
if the Object is not a Surface.
*/
func (recv *Surface) CastToSurface(object *gobject.Object) *Surface {
	return SurfaceNewFromNative(object.Native())
}

func (recv *Surface) Native() unsafe.Pointer {
	return recv.native
}

// SurfaceStruct creates an uninitialised Surface.
func SurfaceStruct() *Surface {
	err := surfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SurfaceNewFromNative(surfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSurface)
	return structGo
}
func finalizeSurface(obj *Surface) {
	surfaceStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func MatrixNewFromNative(native unsafe.Pointer) *Matrix {
	return &Matrix{native: native}
}

/*
CastToMatrix down casts any arbitrary Object to Matrix.
Exercise care, as this is a potentially dangerous function
if the Object is not a Matrix.
*/
func (recv *Matrix) CastToMatrix(object *gobject.Object) *Matrix {
	return MatrixNewFromNative(object.Native())
}

func (recv *Matrix) Native() unsafe.Pointer {
	return recv.native
}

// MatrixStruct creates an uninitialised Matrix.
func MatrixStruct() *Matrix {
	err := matrixStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MatrixNewFromNative(matrixStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMatrix)
	return structGo
}
func finalizeMatrix(obj *Matrix) {
	matrixStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func PatternNewFromNative(native unsafe.Pointer) *Pattern {
	return &Pattern{native: native}
}

/*
CastToPattern down casts any arbitrary Object to Pattern.
Exercise care, as this is a potentially dangerous function
if the Object is not a Pattern.
*/
func (recv *Pattern) CastToPattern(object *gobject.Object) *Pattern {
	return PatternNewFromNative(object.Native())
}

func (recv *Pattern) Native() unsafe.Pointer {
	return recv.native
}

// PatternStruct creates an uninitialised Pattern.
func PatternStruct() *Pattern {
	err := patternStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PatternNewFromNative(patternStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePattern)
	return structGo
}
func finalizePattern(obj *Pattern) {
	patternStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func RegionNewFromNative(native unsafe.Pointer) *Region {
	return &Region{native: native}
}

/*
CastToRegion down casts any arbitrary Object to Region.
Exercise care, as this is a potentially dangerous function
if the Object is not a Region.
*/
func (recv *Region) CastToRegion(object *gobject.Object) *Region {
	return RegionNewFromNative(object.Native())
}

func (recv *Region) Native() unsafe.Pointer {
	return recv.native
}

// RegionStruct creates an uninitialised Region.
func RegionStruct() *Region {
	err := regionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RegionNewFromNative(regionStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRegion)
	return structGo
}
func finalizeRegion(obj *Region) {
	regionStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func FontOptionsNewFromNative(native unsafe.Pointer) *FontOptions {
	return &FontOptions{native: native}
}

/*
CastToFontOptions down casts any arbitrary Object to FontOptions.
Exercise care, as this is a potentially dangerous function
if the Object is not a FontOptions.
*/
func (recv *FontOptions) CastToFontOptions(object *gobject.Object) *FontOptions {
	return FontOptionsNewFromNative(object.Native())
}

func (recv *FontOptions) Native() unsafe.Pointer {
	return recv.native
}

// FontOptionsStruct creates an uninitialised FontOptions.
func FontOptionsStruct() *FontOptions {
	err := fontOptionsStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontOptionsNewFromNative(fontOptionsStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontOptions)
	return structGo
}
func finalizeFontOptions(obj *FontOptions) {
	fontOptionsStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func FontFaceNewFromNative(native unsafe.Pointer) *FontFace {
	return &FontFace{native: native}
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

// FontFaceStruct creates an uninitialised FontFace.
func FontFaceStruct() *FontFace {
	err := fontFaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontFaceNewFromNative(fontFaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontFace)
	return structGo
}
func finalizeFontFace(obj *FontFace) {
	fontFaceStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func ScaledFontNewFromNative(native unsafe.Pointer) *ScaledFont {
	return &ScaledFont{native: native}
}

/*
CastToScaledFont down casts any arbitrary Object to ScaledFont.
Exercise care, as this is a potentially dangerous function
if the Object is not a ScaledFont.
*/
func (recv *ScaledFont) CastToScaledFont(object *gobject.Object) *ScaledFont {
	return ScaledFontNewFromNative(object.Native())
}

func (recv *ScaledFont) Native() unsafe.Pointer {
	return recv.native
}

// ScaledFontStruct creates an uninitialised ScaledFont.
func ScaledFontStruct() *ScaledFont {
	err := scaledFontStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ScaledFontNewFromNative(scaledFontStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeScaledFont)
	return structGo
}
func finalizeScaledFont(obj *ScaledFont) {
	scaledFontStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func PathNewFromNative(native unsafe.Pointer) *Path {
	return &Path{native: native}
}

/*
CastToPath down casts any arbitrary Object to Path.
Exercise care, as this is a potentially dangerous function
if the Object is not a Path.
*/
func (recv *Path) CastToPath(object *gobject.Object) *Path {
	return PathNewFromNative(object.Native())
}

func (recv *Path) Native() unsafe.Pointer {
	return recv.native
}

// PathStruct creates an uninitialised Path.
func PathStruct() *Path {
	err := pathStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PathNewFromNative(pathStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePath)
	return structGo
}
func finalizePath(obj *Path) {
	pathStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func RectangleNewFromNative(native unsafe.Pointer) *Rectangle {
	return &Rectangle{native: native}
}

/*
CastToRectangle down casts any arbitrary Object to Rectangle.
Exercise care, as this is a potentially dangerous function
if the Object is not a Rectangle.
*/
func (recv *Rectangle) CastToRectangle(object *gobject.Object) *Rectangle {
	return RectangleNewFromNative(object.Native())
}

func (recv *Rectangle) Native() unsafe.Pointer {
	return recv.native
}

// FieldX returns the C field 'x'.
func (recv *Rectangle) FieldX() float64 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "x")
	value := argValue.Float64()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Rectangle) SetFieldX(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Rectangle) FieldY() float64 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "y")
	value := argValue.Float64()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Rectangle) SetFieldY(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *Rectangle) FieldWidth() float64 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "width")
	value := argValue.Float64()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *Rectangle) SetFieldWidth(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *Rectangle) FieldHeight() float64 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "height")
	value := argValue.Float64()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *Rectangle) SetFieldHeight(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "height", argValue)
}

// RectangleStruct creates an uninitialised Rectangle.
func RectangleStruct() *Rectangle {
	err := rectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RectangleNewFromNative(rectangleStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRectangle)
	return structGo
}
func finalizeRectangle(obj *Rectangle) {
	rectangleStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func RectangleIntNewFromNative(native unsafe.Pointer) *RectangleInt {
	return &RectangleInt{native: native}
}

/*
CastToRectangleInt down casts any arbitrary Object to RectangleInt.
Exercise care, as this is a potentially dangerous function
if the Object is not a RectangleInt.
*/
func (recv *RectangleInt) CastToRectangleInt(object *gobject.Object) *RectangleInt {
	return RectangleIntNewFromNative(object.Native())
}

func (recv *RectangleInt) Native() unsafe.Pointer {
	return recv.native
}

// FieldX returns the C field 'x'.
func (recv *RectangleInt) FieldX() int32 {
	argValue := gi.StructFieldGet(rectangleIntStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *RectangleInt) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleIntStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *RectangleInt) FieldY() int32 {
	argValue := gi.StructFieldGet(rectangleIntStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *RectangleInt) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleIntStruct, recv.native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *RectangleInt) FieldWidth() int32 {
	argValue := gi.StructFieldGet(rectangleIntStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *RectangleInt) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleIntStruct, recv.native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *RectangleInt) FieldHeight() int32 {
	argValue := gi.StructFieldGet(rectangleIntStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *RectangleInt) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleIntStruct, recv.native, "height", argValue)
}

// RectangleIntStruct creates an uninitialised RectangleInt.
func RectangleIntStruct() *RectangleInt {
	err := rectangleIntStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RectangleIntNewFromNative(rectangleIntStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRectangleInt)
	return structGo
}
func finalizeRectangleInt(obj *RectangleInt) {
	rectangleIntStruct.Free(obj.native)
}

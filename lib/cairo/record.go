// Code generated - DO NOT EDIT.

package cairo

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var contextStruct *gi.Struct
var contextStructOnce sync.Once

func contextStructSet() {
	contextStructOnce.Do(func() {
		contextStruct = gi.StructNew("cairo", "Context")
	})
}

type Context struct {
	native uintptr
}

var deviceStruct *gi.Struct
var deviceStructOnce sync.Once

func deviceStructSet() {
	deviceStructOnce.Do(func() {
		deviceStruct = gi.StructNew("cairo", "Device")
	})
}

type Device struct {
	native uintptr
}

var surfaceStruct *gi.Struct
var surfaceStructOnce sync.Once

func surfaceStructSet() {
	surfaceStructOnce.Do(func() {
		surfaceStruct = gi.StructNew("cairo", "Surface")
	})
}

type Surface struct {
	native uintptr
}

var matrixStruct *gi.Struct
var matrixStructOnce sync.Once

func matrixStructSet() {
	matrixStructOnce.Do(func() {
		matrixStruct = gi.StructNew("cairo", "Matrix")
	})
}

type Matrix struct {
	native uintptr
}

var patternStruct *gi.Struct
var patternStructOnce sync.Once

func patternStructSet() {
	patternStructOnce.Do(func() {
		patternStruct = gi.StructNew("cairo", "Pattern")
	})
}

type Pattern struct {
	native uintptr
}

var regionStruct *gi.Struct
var regionStructOnce sync.Once

func regionStructSet() {
	regionStructOnce.Do(func() {
		regionStruct = gi.StructNew("cairo", "Region")
	})
}

type Region struct {
	native uintptr
}

var fontOptionsStruct *gi.Struct
var fontOptionsStructOnce sync.Once

func fontOptionsStructSet() {
	fontOptionsStructOnce.Do(func() {
		fontOptionsStruct = gi.StructNew("cairo", "FontOptions")
	})
}

type FontOptions struct {
	native uintptr
}

var fontFaceStruct *gi.Struct
var fontFaceStructOnce sync.Once

func fontFaceStructSet() {
	fontFaceStructOnce.Do(func() {
		fontFaceStruct = gi.StructNew("cairo", "FontFace")
	})
}

type FontFace struct {
	native uintptr
}

var scaledFontStruct *gi.Struct
var scaledFontStructOnce sync.Once

func scaledFontStructSet() {
	scaledFontStructOnce.Do(func() {
		scaledFontStruct = gi.StructNew("cairo", "ScaledFont")
	})
}

type ScaledFont struct {
	native uintptr
}

var pathStruct *gi.Struct
var pathStructOnce sync.Once

func pathStructSet() {
	pathStructOnce.Do(func() {
		pathStruct = gi.StructNew("cairo", "Path")
	})
}

type Path struct {
	native uintptr
}

var rectangleStruct *gi.Struct
var rectangleStructOnce sync.Once

func rectangleStructSet() {
	rectangleStructOnce.Do(func() {
		rectangleStruct = gi.StructNew("cairo", "Rectangle")
	})
}

type Rectangle struct {
	native uintptr
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'width' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'height' : no Go type for 'gdouble'
}

var rectangleIntStruct *gi.Struct
var rectangleIntStructOnce sync.Once

func rectangleIntStructSet() {
	rectangleIntStructOnce.Do(func() {
		rectangleIntStruct = gi.StructNew("cairo", "RectangleInt")
	})
}

type RectangleInt struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

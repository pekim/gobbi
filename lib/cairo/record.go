// Code generated - DO NOT EDIT.

package cairo

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var contextStruct *gi.Struct
var contextStruct_Once sync.Once

func contextStruct_Set() {
	contextStruct_Once.Do(func() {
		contextStruct = gi.StructNew("cairo", "Context")
	})
}

type Context struct {
	native uintptr
}

var deviceStruct *gi.Struct
var deviceStruct_Once sync.Once

func deviceStruct_Set() {
	deviceStruct_Once.Do(func() {
		deviceStruct = gi.StructNew("cairo", "Device")
	})
}

type Device struct {
	native uintptr
}

var surfaceStruct *gi.Struct
var surfaceStruct_Once sync.Once

func surfaceStruct_Set() {
	surfaceStruct_Once.Do(func() {
		surfaceStruct = gi.StructNew("cairo", "Surface")
	})
}

type Surface struct {
	native uintptr
}

var matrixStruct *gi.Struct
var matrixStruct_Once sync.Once

func matrixStruct_Set() {
	matrixStruct_Once.Do(func() {
		matrixStruct = gi.StructNew("cairo", "Matrix")
	})
}

type Matrix struct {
	native uintptr
}

var patternStruct *gi.Struct
var patternStruct_Once sync.Once

func patternStruct_Set() {
	patternStruct_Once.Do(func() {
		patternStruct = gi.StructNew("cairo", "Pattern")
	})
}

type Pattern struct {
	native uintptr
}

var regionStruct *gi.Struct
var regionStruct_Once sync.Once

func regionStruct_Set() {
	regionStruct_Once.Do(func() {
		regionStruct = gi.StructNew("cairo", "Region")
	})
}

type Region struct {
	native uintptr
}

var fontOptionsStruct *gi.Struct
var fontOptionsStruct_Once sync.Once

func fontOptionsStruct_Set() {
	fontOptionsStruct_Once.Do(func() {
		fontOptionsStruct = gi.StructNew("cairo", "FontOptions")
	})
}

type FontOptions struct {
	native uintptr
}

var fontFaceStruct *gi.Struct
var fontFaceStruct_Once sync.Once

func fontFaceStruct_Set() {
	fontFaceStruct_Once.Do(func() {
		fontFaceStruct = gi.StructNew("cairo", "FontFace")
	})
}

type FontFace struct {
	native uintptr
}

var scaledFontStruct *gi.Struct
var scaledFontStruct_Once sync.Once

func scaledFontStruct_Set() {
	scaledFontStruct_Once.Do(func() {
		scaledFontStruct = gi.StructNew("cairo", "ScaledFont")
	})
}

type ScaledFont struct {
	native uintptr
}

var pathStruct *gi.Struct
var pathStruct_Once sync.Once

func pathStruct_Set() {
	pathStruct_Once.Do(func() {
		pathStruct = gi.StructNew("cairo", "Path")
	})
}

type Path struct {
	native uintptr
}

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() {
	rectangleStruct_Once.Do(func() {
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
var rectangleIntStruct_Once sync.Once

func rectangleIntStruct_Set() {
	rectangleIntStruct_Once.Do(func() {
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

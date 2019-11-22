// Code generated - DO NOT EDIT.

package cairo

import (
	gi "github.com/pekim/gobbi/internal/gi"
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
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
	native uintptr
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'width' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'height' : no Go type for 'gdouble'
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
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

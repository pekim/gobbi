// Code generated - DO NOT EDIT.

package cairo

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var ContextStruct *gi.Struct
var ContextStructOnce sync.Once

func ContextStructSet() {
	ContextStructOnce.Do(func() {
		ContextStruct = gi.StructNew("cairo", "Context")
	})
}

type Context struct {
	native uintptr
}

var DeviceStruct *gi.Struct
var DeviceStructOnce sync.Once

func DeviceStructSet() {
	DeviceStructOnce.Do(func() {
		DeviceStruct = gi.StructNew("cairo", "Device")
	})
}

type Device struct {
	native uintptr
}

var SurfaceStruct *gi.Struct
var SurfaceStructOnce sync.Once

func SurfaceStructSet() {
	SurfaceStructOnce.Do(func() {
		SurfaceStruct = gi.StructNew("cairo", "Surface")
	})
}

type Surface struct {
	native uintptr
}

var MatrixStruct *gi.Struct
var MatrixStructOnce sync.Once

func MatrixStructSet() {
	MatrixStructOnce.Do(func() {
		MatrixStruct = gi.StructNew("cairo", "Matrix")
	})
}

type Matrix struct {
	native uintptr
}

var PatternStruct *gi.Struct
var PatternStructOnce sync.Once

func PatternStructSet() {
	PatternStructOnce.Do(func() {
		PatternStruct = gi.StructNew("cairo", "Pattern")
	})
}

type Pattern struct {
	native uintptr
}

var RegionStruct *gi.Struct
var RegionStructOnce sync.Once

func RegionStructSet() {
	RegionStructOnce.Do(func() {
		RegionStruct = gi.StructNew("cairo", "Region")
	})
}

type Region struct {
	native uintptr
}

var FontOptionsStruct *gi.Struct
var FontOptionsStructOnce sync.Once

func FontOptionsStructSet() {
	FontOptionsStructOnce.Do(func() {
		FontOptionsStruct = gi.StructNew("cairo", "FontOptions")
	})
}

type FontOptions struct {
	native uintptr
}

var FontFaceStruct *gi.Struct
var FontFaceStructOnce sync.Once

func FontFaceStructSet() {
	FontFaceStructOnce.Do(func() {
		FontFaceStruct = gi.StructNew("cairo", "FontFace")
	})
}

type FontFace struct {
	native uintptr
}

var ScaledFontStruct *gi.Struct
var ScaledFontStructOnce sync.Once

func ScaledFontStructSet() {
	ScaledFontStructOnce.Do(func() {
		ScaledFontStruct = gi.StructNew("cairo", "ScaledFont")
	})
}

type ScaledFont struct {
	native uintptr
}

var PathStruct *gi.Struct
var PathStructOnce sync.Once

func PathStructSet() {
	PathStructOnce.Do(func() {
		PathStruct = gi.StructNew("cairo", "Path")
	})
}

type Path struct {
	native uintptr
}

var RectangleStruct *gi.Struct
var RectangleStructOnce sync.Once

func RectangleStructSet() {
	RectangleStructOnce.Do(func() {
		RectangleStruct = gi.StructNew("cairo", "Rectangle")
	})
}

type Rectangle struct {
	native uintptr
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'width' : no Go type for 'gdouble'
	// UNSUPPORTED : C value 'height' : no Go type for 'gdouble'
}

var RectangleIntStruct *gi.Struct
var RectangleIntStructOnce sync.Once

func RectangleIntStructSet() {
	RectangleIntStructOnce.Do(func() {
		RectangleIntStruct = gi.StructNew("cairo", "RectangleInt")
	})
}

type RectangleInt struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

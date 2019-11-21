// Code generated - DO NOT EDIT.

package freetype2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var bitmapStruct *gi.Struct
var bitmapStructOnce sync.Once

func bitmapStructSet() {
	bitmapStructOnce.Do(func() {
		bitmapStruct = gi.StructNew("freetype2", "Bitmap")
	})
}

type Bitmap struct {
	native uintptr
}

var faceStruct *gi.Struct
var faceStructOnce sync.Once

func faceStructSet() {
	faceStructOnce.Do(func() {
		faceStruct = gi.StructNew("freetype2", "Face")
	})
}

type Face struct {
	native uintptr
}

var libraryStruct *gi.Struct
var libraryStructOnce sync.Once

func libraryStructSet() {
	libraryStructOnce.Do(func() {
		libraryStruct = gi.StructNew("freetype2", "Library")
	})
}

type Library struct {
	native uintptr
}

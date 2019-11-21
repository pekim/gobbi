// Code generated - DO NOT EDIT.

package freetype2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var bitmapStruct *gi.Struct
var bitmapStruct_Once sync.Once

func bitmapStruct_Set() {
	bitmapStruct_Once.Do(func() {
		bitmapStruct = gi.StructNew("freetype2", "Bitmap")
	})
}

type Bitmap struct {
	native uintptr
}

var faceStruct *gi.Struct
var faceStruct_Once sync.Once

func faceStruct_Set() {
	faceStruct_Once.Do(func() {
		faceStruct = gi.StructNew("freetype2", "Face")
	})
}

type Face struct {
	native uintptr
}

var libraryStruct *gi.Struct
var libraryStruct_Once sync.Once

func libraryStruct_Set() {
	libraryStruct_Once.Do(func() {
		libraryStruct = gi.StructNew("freetype2", "Library")
	})
}

type Library struct {
	native uintptr
}

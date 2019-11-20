// Code generated - DO NOT EDIT.

package freetype2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var BitmapStruct *gi.Struct
var BitmapStructOnce sync.Once

func BitmapStructSet() {
	BitmapStructOnce.Do(func() {
		BitmapStruct = gi.StructNew("freetype2", "Bitmap")
	})
}

type Bitmap struct {
	native uintptr
}

var FaceStruct *gi.Struct
var FaceStructOnce sync.Once

func FaceStructSet() {
	FaceStructOnce.Do(func() {
		FaceStruct = gi.StructNew("freetype2", "Face")
	})
}

type Face struct {
	native uintptr
}

var LibraryStruct *gi.Struct
var LibraryStructOnce sync.Once

func LibraryStructSet() {
	LibraryStructOnce.Do(func() {
		LibraryStruct = gi.StructNew("freetype2", "Library")
	})
}

type Library struct {
	native uintptr
}

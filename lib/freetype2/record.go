// Code generated - DO NOT EDIT.

package freetype2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var bitmapStruct *gi.Struct
var bitmapStruct_Once sync.Once

func bitmapStruct_Set() error {
	var err error
	bitmapStruct_Once.Do(func() {
		bitmapStruct, err = gi.StructNew("freetype2", "Bitmap")
	})
	return err
}

type Bitmap struct {
	native uintptr
}

var faceStruct *gi.Struct
var faceStruct_Once sync.Once

func faceStruct_Set() error {
	var err error
	faceStruct_Once.Do(func() {
		faceStruct, err = gi.StructNew("freetype2", "Face")
	})
	return err
}

type Face struct {
	native uintptr
}

var libraryStruct *gi.Struct
var libraryStruct_Once sync.Once

func libraryStruct_Set() error {
	var err error
	libraryStruct_Once.Do(func() {
		libraryStruct, err = gi.StructNew("freetype2", "Library")
	})
	return err
}

type Library struct {
	native uintptr
}

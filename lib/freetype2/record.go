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

// BitmapStruct creates an uninitialised Bitmap.
func BitmapStruct() *Bitmap {
	err := bitmapStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Bitmap{native: bitmapStruct.Alloc()}
	return structGo
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

// FaceStruct creates an uninitialised Face.
func FaceStruct() *Face {
	err := faceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Face{native: faceStruct.Alloc()}
	return structGo
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

// LibraryStruct creates an uninitialised Library.
func LibraryStruct() *Library {
	err := libraryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Library{native: libraryStruct.Alloc()}
	return structGo
}

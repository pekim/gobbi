// Code generated - DO NOT EDIT.

package freetype2

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
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
	native unsafe.Pointer
}

func BitmapNewFromNative(native unsafe.Pointer) *Bitmap {
	return &Bitmap{native: native}
}

/*
CastToBitmap down casts any arbitrary Object to Bitmap.
Exercise care, as this is a potentially dangerous function
if the Object is not a Bitmap.
*/
func (recv *Bitmap) CastToBitmap(object *gobject.Object) *Bitmap {
	return BitmapNewFromNative(object.Native())
}

func (recv *Bitmap) Native() unsafe.Pointer {
	return recv.native
}

// BitmapStruct creates an uninitialised Bitmap.
func BitmapStruct() *Bitmap {
	err := bitmapStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BitmapNewFromNative(bitmapStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBitmap)
	return structGo
}
func finalizeBitmap(obj *Bitmap) {
	bitmapStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func FaceNewFromNative(native unsafe.Pointer) *Face {
	return &Face{native: native}
}

/*
CastToFace down casts any arbitrary Object to Face.
Exercise care, as this is a potentially dangerous function
if the Object is not a Face.
*/
func (recv *Face) CastToFace(object *gobject.Object) *Face {
	return FaceNewFromNative(object.Native())
}

func (recv *Face) Native() unsafe.Pointer {
	return recv.native
}

// FaceStruct creates an uninitialised Face.
func FaceStruct() *Face {
	err := faceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FaceNewFromNative(faceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFace)
	return structGo
}
func finalizeFace(obj *Face) {
	faceStruct.Free(obj.native)
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
	native unsafe.Pointer
}

func LibraryNewFromNative(native unsafe.Pointer) *Library {
	return &Library{native: native}
}

/*
CastToLibrary down casts any arbitrary Object to Library.
Exercise care, as this is a potentially dangerous function
if the Object is not a Library.
*/
func (recv *Library) CastToLibrary(object *gobject.Object) *Library {
	return LibraryNewFromNative(object.Native())
}

func (recv *Library) Native() unsafe.Pointer {
	return recv.native
}

// LibraryStruct creates an uninitialised Library.
func LibraryStruct() *Library {
	err := libraryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LibraryNewFromNative(libraryStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLibrary)
	return structGo
}
func finalizeLibrary(obj *Library) {
	libraryStruct.Free(obj.native)
}

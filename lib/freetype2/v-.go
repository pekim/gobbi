// Code generated - DO NOT EDIT.

package freetype2

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <stdlib.h>
import "C"

// Unsupported : alias has no type generator for int32,

// LibraryVersion is a wrapper around the C function FT_Library_Version.
func LibraryVersion() {
	C.FT_Library_Version()

	return
}

// Bitmap is a wrapper around the C record FT_Bitmap.
type Bitmap struct {
	native *C.FT_Bitmap
}

func BitmapNewFromC(u unsafe.Pointer) *Bitmap {
	c := (*C.FT_Bitmap)(u)
	if c == nil {
		return nil
	}

	g := &Bitmap{native: c}

	return g
}

func (recv *Bitmap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Bitmap with another Bitmap, and returns true if they represent the same GObject.
func (recv *Bitmap) Equals(other *Bitmap) bool {
	return other.ToC() == recv.ToC()
}

// Face is a wrapper around the C record FT_Face.
type Face struct {
	native *C.FT_Face
}

func FaceNewFromC(u unsafe.Pointer) *Face {
	c := (*C.FT_Face)(u)
	if c == nil {
		return nil
	}

	g := &Face{native: c}

	return g
}

func (recv *Face) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Face with another Face, and returns true if they represent the same GObject.
func (recv *Face) Equals(other *Face) bool {
	return other.ToC() == recv.ToC()
}

// Library is a wrapper around the C record FT_Library.
type Library struct {
	native *C.FT_Library
}

func LibraryNewFromC(u unsafe.Pointer) *Library {
	c := (*C.FT_Library)(u)
	if c == nil {
		return nil
	}

	g := &Library{native: c}

	return g
}

func (recv *Library) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Library with another Library, and returns true if they represent the same GObject.
func (recv *Library) Equals(other *Library) bool {
	return other.ToC() == recv.ToC()
}

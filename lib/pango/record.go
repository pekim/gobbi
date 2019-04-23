// This is a generated file - DO NOT EDIT

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Blacklisted : PangoAnalysis

// Blacklisted : PangoAttrClass

// Blacklisted : PangoAttrColor

// Blacklisted : PangoAttrFloat

// Blacklisted : PangoAttrFontDesc

// Blacklisted : PangoAttrInt

// Blacklisted : PangoAttrIterator

// Blacklisted : PangoAttrLanguage

// Blacklisted : PangoAttrList

// Blacklisted : PangoAttrShape

// Blacklisted : PangoAttrSize

// Blacklisted : PangoAttrString

// Blacklisted : PangoAttribute

// Blacklisted : PangoColor

// Blacklisted : PangoContextClass

// Blacklisted : PangoCoverage

// Blacklisted : PangoEngineClass

// Blacklisted : PangoEngineInfo

// Blacklisted : PangoEngineLangClass

// Blacklisted : PangoEngineScriptInfo

// Blacklisted : PangoEngineShapeClass

// Blacklisted : PangoFontClass

// Blacklisted : PangoFontDescription

// Blacklisted : PangoFontFaceClass

// Blacklisted : PangoFontFamilyClass

// Blacklisted : PangoFontMapClass

// Blacklisted : PangoFontMetrics

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// Blacklisted : PangoGlyphGeometry

// Blacklisted : PangoGlyphInfo

// GlyphItem is a wrapper around the C record PangoGlyphItem.
type GlyphItem struct {
	native *C.PangoGlyphItem
	// item : record
	// glyphs : record
}

func GlyphItemNewFromC(u unsafe.Pointer) *GlyphItem {
	c := (*C.PangoGlyphItem)(u)
	if c == nil {
		return nil
	}

	g := &GlyphItem{native: c}

	return g
}

func (recv *GlyphItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GlyphItem with another GlyphItem, and returns true if they represent the same GObject.
func (recv *GlyphItem) Equals(other *GlyphItem) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : PangoGlyphString

// Blacklisted : PangoGlyphVisAttr

// Blacklisted : PangoIncludedModule

// Blacklisted : PangoItem

// Blacklisted : PangoLanguage

// Blacklisted : PangoLayoutClass

// Blacklisted : PangoLayoutIter

// Blacklisted : PangoLayoutLine

// Blacklisted : PangoLogAttr

// Blacklisted : PangoMap

// Blacklisted : PangoMapEntry

// Rectangle is a wrapper around the C record PangoRectangle.
type Rectangle struct {
	native *C.PangoRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	c := (*C.PangoRectangle)(u)
	if c == nil {
		return nil
	}

	g := &Rectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.int)(recv.X)
	recv.native.y =
		(C.int)(recv.Y)
	recv.native.width =
		(C.int)(recv.Width)
	recv.native.height =
		(C.int)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same GObject.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : PangoRendererPrivate

// Blacklisted : PangoScriptForLang

// Blacklisted : PangoScriptIter

// Blacklisted : PangoTabArray

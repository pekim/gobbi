// This is a generated file - DO NOT EDIT

package pango

import "unsafe"

// Context is a wrapper around the C record PangoContext.
type Context struct {
	native unsafe.Pointer
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	if u == nil {
		return nil
	}

	g := &Context{native: u}

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoEngine

// EngineLang is a wrapper around the C record PangoEngineLang.
type EngineLang struct {
	native unsafe.Pointer
	// Private : parent_instance
}

func EngineLangNewFromC(u unsafe.Pointer) *EngineLang {
	if u == nil {
		return nil
	}

	g := &EngineLang{native: u}

	return g
}

func (recv *EngineLang) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EngineShape is a wrapper around the C record PangoEngineShape.
type EngineShape struct {
	native unsafe.Pointer
	// parent_instance : record
}

func EngineShapeNewFromC(u unsafe.Pointer) *EngineShape {
	if u == nil {
		return nil
	}

	g := &EngineShape{native: u}

	return g
}

func (recv *EngineShape) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Font is a wrapper around the C record PangoFont.
type Font struct {
	native unsafe.Pointer
	// parent_instance : record
}

func FontNewFromC(u unsafe.Pointer) *Font {
	if u == nil {
		return nil
	}

	g := &Font{native: u}

	return g
}

func (recv *Font) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontFace is a wrapper around the C record PangoFontFace.
type FontFace struct {
	native unsafe.Pointer
	// parent_instance : record
}

func FontFaceNewFromC(u unsafe.Pointer) *FontFace {
	if u == nil {
		return nil
	}

	g := &FontFace{native: u}

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontFamily is a wrapper around the C record PangoFontFamily.
type FontFamily struct {
	native unsafe.Pointer
	// parent_instance : record
}

func FontFamilyNewFromC(u unsafe.Pointer) *FontFamily {
	if u == nil {
		return nil
	}

	g := &FontFamily{native: u}

	return g
}

func (recv *FontFamily) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontMap is a wrapper around the C record PangoFontMap.
type FontMap struct {
	native unsafe.Pointer
	// parent_instance : record
}

func FontMapNewFromC(u unsafe.Pointer) *FontMap {
	if u == nil {
		return nil
	}

	g := &FontMap{native: u}

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Fontset is a wrapper around the C record PangoFontset.
type Fontset struct {
	native unsafe.Pointer
	// parent_instance : record
}

func FontsetNewFromC(u unsafe.Pointer) *Fontset {
	if u == nil {
		return nil
	}

	g := &Fontset{native: u}

	return g
}

func (recv *Fontset) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoFontsetSimple

// Layout is a wrapper around the C record PangoLayout.
type Layout struct {
	native unsafe.Pointer
}

func LayoutNewFromC(u unsafe.Pointer) *Layout {
	if u == nil {
		return nil
	}

	g := &Layout{native: u}

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// This is a generated file - DO NOT EDIT

package pango

import "unsafe"

// Context is a wrapper around the C record PangoContext.
type Context struct {
	native unsafe.Pointer
}

// Blacklisted : PangoEngine

// EngineLang is a wrapper around the C record PangoEngineLang.
type EngineLang struct {
	native unsafe.Pointer
	// Private : parent_instance
}

// EngineShape is a wrapper around the C record PangoEngineShape.
type EngineShape struct {
	native unsafe.Pointer
	// parent_instance : record
}

// Font is a wrapper around the C record PangoFont.
type Font struct {
	native unsafe.Pointer
	// parent_instance : record
}

// FontFace is a wrapper around the C record PangoFontFace.
type FontFace struct {
	native unsafe.Pointer
	// parent_instance : record
}

// FontFamily is a wrapper around the C record PangoFontFamily.
type FontFamily struct {
	native unsafe.Pointer
	// parent_instance : record
}

// FontMap is a wrapper around the C record PangoFontMap.
type FontMap struct {
	native unsafe.Pointer
	// parent_instance : record
}

// Fontset is a wrapper around the C record PangoFontset.
type Fontset struct {
	native unsafe.Pointer
	// parent_instance : record
}

// Blacklisted : PangoFontsetSimple

// Layout is a wrapper around the C record PangoLayout.
type Layout struct {
	native unsafe.Pointer
}

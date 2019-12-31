// Code generated - DO NOT EDIT.
// +build pango_1.31

package pango

import "unsafe"

// Glyph is a representation of the C alias PangoGlyph.
type Glyph uint32

// GlyphUnit is a representation of the C alias PangoGlyphUnit.
type GlyphUnit int32

// LayoutRun is a representation of the C alias PangoLayoutRun.
type LayoutRun GlyphItem

// ANALYSIS_FLAG_CENTERED_BASELINE is a representation of the C constant PANGO_ANALYSIS_FLAG_CENTERED_BASELINE.
//
// since 1.16
const ANALYSIS_FLAG_CENTERED_BASELINE = 1

// ATTR_INDEX_FROM_TEXT_BEGINNING is a representation of the C constant PANGO_ATTR_INDEX_FROM_TEXT_BEGINNING.
//
// since 1.24
const ATTR_INDEX_FROM_TEXT_BEGINNING = 0

// ENGINE_TYPE_LANG is a representation of the C constant PANGO_ENGINE_TYPE_LANG.
const ENGINE_TYPE_LANG = "PangoEngineLang"

// ENGINE_TYPE_SHAPE is a representation of the C constant PANGO_ENGINE_TYPE_SHAPE.
const ENGINE_TYPE_SHAPE = "PangoEngineShape"

// GLYPH_EMPTY is a representation of the C constant PANGO_GLYPH_EMPTY.
const GLYPH_EMPTY = uint64(0xfffffff)

// GLYPH_INVALID_INPUT is a representation of the C constant PANGO_GLYPH_INVALID_INPUT.
//
// since 1.20
const GLYPH_INVALID_INPUT = uint64(0xffffffff)

// GLYPH_UNKNOWN_FLAG is a representation of the C constant PANGO_GLYPH_UNKNOWN_FLAG.
const GLYPH_UNKNOWN_FLAG = uint64(0x10000000)

// RENDER_TYPE_NONE is a representation of the C constant PANGO_RENDER_TYPE_NONE.
const RENDER_TYPE_NONE = "PangoRenderNone"

// SCALE is a representation of the C constant PANGO_SCALE.
const SCALE = 1024

// UNKNOWN_GLYPH_HEIGHT is a representation of the C constant PANGO_UNKNOWN_GLYPH_HEIGHT.
const UNKNOWN_GLYPH_HEIGHT = 14

// UNKNOWN_GLYPH_WIDTH is a representation of the C constant PANGO_UNKNOWN_GLYPH_WIDTH.
const UNKNOWN_GLYPH_WIDTH = 10

// Analysis is a representation of the C record PangoAnalysis.
type Analysis struct {
	native unsafe.Pointer
}

// AttrClass is a representation of the C record PangoAttrClass.
type AttrClass struct {
	native unsafe.Pointer
}

// AttrColor is a representation of the C record PangoAttrColor.
type AttrColor struct {
	native unsafe.Pointer
}

// AttrFloat is a representation of the C record PangoAttrFloat.
type AttrFloat struct {
	native unsafe.Pointer
}

// AttrFontDesc is a representation of the C record PangoAttrFontDesc.
type AttrFontDesc struct {
	native unsafe.Pointer
}

// AttrInt is a representation of the C record PangoAttrInt.
type AttrInt struct {
	native unsafe.Pointer
}

// AttrIterator is a representation of the C record PangoAttrIterator.
type AttrIterator struct {
	native unsafe.Pointer
}

// AttrLanguage is a representation of the C record PangoAttrLanguage.
type AttrLanguage struct {
	native unsafe.Pointer
}

// AttrList is a representation of the C record PangoAttrList.
type AttrList struct {
	native unsafe.Pointer
}

// AttrShape is a representation of the C record PangoAttrShape.
type AttrShape struct {
	native unsafe.Pointer
}

// AttrSize is a representation of the C record PangoAttrSize.
type AttrSize struct {
	native unsafe.Pointer
}

// AttrString is a representation of the C record PangoAttrString.
type AttrString struct {
	native unsafe.Pointer
}

// Attribute is a representation of the C record PangoAttribute.
type Attribute struct {
	native unsafe.Pointer
}

// Color is a representation of the C record PangoColor.
type Color struct {
	native unsafe.Pointer
}

// ContextClass is a representation of the C record PangoContextClass.
type ContextClass struct {
	native unsafe.Pointer
}

// Coverage is a representation of the C record PangoCoverage.
type Coverage struct {
	native unsafe.Pointer
}

// UNSUPPORTED : EngineClass : blacklisted

// UNSUPPORTED : EngineInfo : blacklisted

// UNSUPPORTED : EngineLangClass : blacklisted

// UNSUPPORTED : EngineScriptInfo : blacklisted

// UNSUPPORTED : EngineShapeClass : blacklisted

// UNSUPPORTED : FontClass : blacklisted

// FontDescription is a representation of the C record PangoFontDescription.
type FontDescription struct {
	native unsafe.Pointer
}

// UNSUPPORTED : FontFaceClass : blacklisted

// UNSUPPORTED : FontFamilyClass : blacklisted

// UNSUPPORTED : FontMapClass : blacklisted

// FontMetrics is a representation of the C record PangoFontMetrics.
type FontMetrics struct {
	native unsafe.Pointer
}

// UNSUPPORTED : FontsetClass : blacklisted

// UNSUPPORTED : FontsetSimpleClass : blacklisted

// GlyphGeometry is a representation of the C record PangoGlyphGeometry.
type GlyphGeometry struct {
	native unsafe.Pointer
}

// GlyphInfo is a representation of the C record PangoGlyphInfo.
type GlyphInfo struct {
	native unsafe.Pointer
}

// GlyphItem is a representation of the C record PangoGlyphItem.
type GlyphItem struct {
	native unsafe.Pointer
}

// GlyphItemIter is a representation of the C record PangoGlyphItemIter.
type GlyphItemIter struct {
	native unsafe.Pointer
}

// GlyphString is a representation of the C record PangoGlyphString.
type GlyphString struct {
	native unsafe.Pointer
}

// GlyphVisAttr is a representation of the C record PangoGlyphVisAttr.
type GlyphVisAttr struct {
	native unsafe.Pointer
}

// UNSUPPORTED : IncludedModule : blacklisted

// Item is a representation of the C record PangoItem.
type Item struct {
	native unsafe.Pointer
}

// Language is a representation of the C record PangoLanguage.
type Language struct {
	native unsafe.Pointer
}

// LayoutClass is a representation of the C record PangoLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
}

// LayoutIter is a representation of the C record PangoLayoutIter.
type LayoutIter struct {
	native unsafe.Pointer
}

// LayoutLine is a representation of the C record PangoLayoutLine.
type LayoutLine struct {
	native unsafe.Pointer
}

// LogAttr is a representation of the C record PangoLogAttr.
type LogAttr struct {
	native unsafe.Pointer
}

// UNSUPPORTED : Map : blacklisted

// UNSUPPORTED : MapEntry : blacklisted

// Matrix is a representation of the C record PangoMatrix.
type Matrix struct {
	native unsafe.Pointer
}

// Rectangle is a representation of the C record PangoRectangle.
type Rectangle struct {
	native unsafe.Pointer
}

// RendererClass is a representation of the C record PangoRendererClass.
type RendererClass struct {
	native unsafe.Pointer
}

// RendererPrivate is a representation of the C record PangoRendererPrivate.
type RendererPrivate struct {
	native unsafe.Pointer
}

// ScriptIter is a representation of the C record PangoScriptIter.
type ScriptIter struct {
	native unsafe.Pointer
}

// TabArray is a representation of the C record PangoTabArray.
type TabArray struct {
	native unsafe.Pointer
}

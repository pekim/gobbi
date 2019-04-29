// This is a generated file - DO NOT EDIT

package pango

import "unsafe"

// Analysis is a wrapper around the C record PangoAnalysis.
type Analysis struct {
	native unsafe.Pointer
	// shape_engine : record
	// lang_engine : record
	// font : record
	Level   uint8
	Gravity uint8
	Flags   uint8
	Script  uint8
	// language : record
	// extra_attrs : record
}

// AttrClass is a wrapper around the C record PangoAttrClass.
type AttrClass struct {
	native unsafe.Pointer
	Type   AttrType
	// no type for copy
	// no type for destroy
	// no type for equal
}

// AttrColor is a wrapper around the C record PangoAttrColor.
type AttrColor struct {
	native unsafe.Pointer
	// attr : record
	// color : record
}

// AttrFloat is a wrapper around the C record PangoAttrFloat.
type AttrFloat struct {
	native unsafe.Pointer
	// attr : record
	Value float64
}

// AttrFontDesc is a wrapper around the C record PangoAttrFontDesc.
type AttrFontDesc struct {
	native unsafe.Pointer
	// attr : record
	// desc : record
}

// AttrInt is a wrapper around the C record PangoAttrInt.
type AttrInt struct {
	native unsafe.Pointer
	// attr : record
	Value int32
}

// AttrIterator is a wrapper around the C record PangoAttrIterator.
type AttrIterator struct {
	native unsafe.Pointer
}

// AttrLanguage is a wrapper around the C record PangoAttrLanguage.
type AttrLanguage struct {
	native unsafe.Pointer
	// attr : record
	// value : record
}

// AttrList is a wrapper around the C record PangoAttrList.
type AttrList struct {
	native unsafe.Pointer
}

// AttrShape is a wrapper around the C record PangoAttrShape.
type AttrShape struct {
	native unsafe.Pointer
	// attr : record
	// ink_rect : record
	// logical_rect : record
	// data : no type generator for gpointer, gpointer
	// copy_func : no type generator for AttrDataCopyFunc, PangoAttrDataCopyFunc
	// destroy_func : no type generator for GLib.DestroyNotify, GDestroyNotify
}

// AttrSize is a wrapper around the C record PangoAttrSize.
type AttrSize struct {
	native unsafe.Pointer
	// attr : record
	Size int32
	// Bitfield not supported :  1 absolute
}

// AttrString is a wrapper around the C record PangoAttrString.
type AttrString struct {
	native unsafe.Pointer
	// attr : record
	Value string
}

// Attribute is a wrapper around the C record PangoAttribute.
type Attribute struct {
	native unsafe.Pointer
	// klass : record
	StartIndex uint32
	EndIndex   uint32
}

// Color is a wrapper around the C record PangoColor.
type Color struct {
	native unsafe.Pointer
	Red    uint16
	Green  uint16
	Blue   uint16
}

// ContextClass is a wrapper around the C record PangoContextClass.
type ContextClass struct {
	native unsafe.Pointer
}

// Coverage is a wrapper around the C record PangoCoverage.
type Coverage struct {
	native unsafe.Pointer
}

// Blacklisted : PangoEngineClass

// Blacklisted : PangoEngineInfo

// Blacklisted : PangoEngineLangClass

// Blacklisted : PangoEngineScriptInfo

// Blacklisted : PangoEngineShapeClass

// Blacklisted : PangoFontClass

// FontDescription is a wrapper around the C record PangoFontDescription.
type FontDescription struct {
	native unsafe.Pointer
}

// Blacklisted : PangoFontFaceClass

// Blacklisted : PangoFontFamilyClass

// Blacklisted : PangoFontMapClass

// FontMetrics is a wrapper around the C record PangoFontMetrics.
type FontMetrics struct {
	native unsafe.Pointer
	// Private : ref_count
	// Private : ascent
	// Private : descent
	// Private : approximate_char_width
	// Private : approximate_digit_width
	// Private : underline_position
	// Private : underline_thickness
	// Private : strikethrough_position
	// Private : strikethrough_thickness
}

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// GlyphGeometry is a wrapper around the C record PangoGlyphGeometry.
type GlyphGeometry struct {
	native  unsafe.Pointer
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

// GlyphInfo is a wrapper around the C record PangoGlyphInfo.
type GlyphInfo struct {
	native unsafe.Pointer
	Glyph  Glyph
	// geometry : record
	// attr : record
}

// GlyphItem is a wrapper around the C record PangoGlyphItem.
type GlyphItem struct {
	native unsafe.Pointer
	// item : record
	// glyphs : record
}

// Blacklisted : PangoGlyphString

// GlyphVisAttr is a wrapper around the C record PangoGlyphVisAttr.
type GlyphVisAttr struct {
	native unsafe.Pointer
	// Bitfield not supported :  1 is_cluster_start
}

// Blacklisted : PangoIncludedModule

// Item is a wrapper around the C record PangoItem.
type Item struct {
	native   unsafe.Pointer
	Offset   int32
	Length   int32
	NumChars int32
	// analysis : record
}

// Language is a wrapper around the C record PangoLanguage.
type Language struct {
	native unsafe.Pointer
}

// LayoutClass is a wrapper around the C record PangoLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
}

// Blacklisted : PangoLayoutIter

// LayoutLine is a wrapper around the C record PangoLayoutLine.
type LayoutLine struct {
	native unsafe.Pointer
	// layout : record
	StartIndex int32
	Length     int32
	// runs : record
	// Bitfield not supported :  1 is_paragraph_start
	// Bitfield not supported :  3 resolved_dir
}

// LogAttr is a wrapper around the C record PangoLogAttr.
type LogAttr struct {
	native unsafe.Pointer
	// Bitfield not supported :  1 is_line_break
	// Bitfield not supported :  1 is_mandatory_break
	// Bitfield not supported :  1 is_char_break
	// Bitfield not supported :  1 is_white
	// Bitfield not supported :  1 is_cursor_position
	// Bitfield not supported :  1 is_word_start
	// Bitfield not supported :  1 is_word_end
	// Bitfield not supported :  1 is_sentence_boundary
	// Bitfield not supported :  1 is_sentence_start
	// Bitfield not supported :  1 is_sentence_end
	// Bitfield not supported :  1 backspace_deletes_character
	// Bitfield not supported :  1 is_expandable_space
	// Bitfield not supported :  1 is_word_boundary
}

// Blacklisted : PangoMap

// Blacklisted : PangoMapEntry

// Rectangle is a wrapper around the C record PangoRectangle.
type Rectangle struct {
	native unsafe.Pointer
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// RendererPrivate is a wrapper around the C record PangoRendererPrivate.
type RendererPrivate struct {
	native unsafe.Pointer
}

// Blacklisted : PangoScriptForLang

// ScriptIter is a wrapper around the C record PangoScriptIter.
type ScriptIter struct {
	native unsafe.Pointer
}

// TabArray is a wrapper around the C record PangoTabArray.
type TabArray struct {
	native unsafe.Pointer
}

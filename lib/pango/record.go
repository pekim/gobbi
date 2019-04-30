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

func AnalysisNewFromC(u unsafe.Pointer) *Analysis {
	if u == nil {
		return nil
	}

	g := &Analysis{native: u}

	return g
}

func (recv *Analysis) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrClass is a wrapper around the C record PangoAttrClass.
type AttrClass struct {
	native unsafe.Pointer
	Type   AttrType
	// no type for copy
	// no type for destroy
	// no type for equal
}

func AttrClassNewFromC(u unsafe.Pointer) *AttrClass {
	if u == nil {
		return nil
	}

	g := &AttrClass{native: u}

	return g
}

func (recv *AttrClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrColor is a wrapper around the C record PangoAttrColor.
type AttrColor struct {
	native unsafe.Pointer
	// attr : record
	// color : record
}

func AttrColorNewFromC(u unsafe.Pointer) *AttrColor {
	if u == nil {
		return nil
	}

	g := &AttrColor{native: u}

	return g
}

func (recv *AttrColor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrFloat is a wrapper around the C record PangoAttrFloat.
type AttrFloat struct {
	native unsafe.Pointer
	// attr : record
	Value float64
}

func AttrFloatNewFromC(u unsafe.Pointer) *AttrFloat {
	if u == nil {
		return nil
	}

	g := &AttrFloat{native: u}

	return g
}

func (recv *AttrFloat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrFontDesc is a wrapper around the C record PangoAttrFontDesc.
type AttrFontDesc struct {
	native unsafe.Pointer
	// attr : record
	// desc : record
}

func AttrFontDescNewFromC(u unsafe.Pointer) *AttrFontDesc {
	if u == nil {
		return nil
	}

	g := &AttrFontDesc{native: u}

	return g
}

func (recv *AttrFontDesc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrInt is a wrapper around the C record PangoAttrInt.
type AttrInt struct {
	native unsafe.Pointer
	// attr : record
	Value int32
}

func AttrIntNewFromC(u unsafe.Pointer) *AttrInt {
	if u == nil {
		return nil
	}

	g := &AttrInt{native: u}

	return g
}

func (recv *AttrInt) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrIterator is a wrapper around the C record PangoAttrIterator.
type AttrIterator struct {
	native unsafe.Pointer
}

func AttrIteratorNewFromC(u unsafe.Pointer) *AttrIterator {
	if u == nil {
		return nil
	}

	g := &AttrIterator{native: u}

	return g
}

func (recv *AttrIterator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrLanguage is a wrapper around the C record PangoAttrLanguage.
type AttrLanguage struct {
	native unsafe.Pointer
	// attr : record
	// value : record
}

func AttrLanguageNewFromC(u unsafe.Pointer) *AttrLanguage {
	if u == nil {
		return nil
	}

	g := &AttrLanguage{native: u}

	return g
}

func (recv *AttrLanguage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrList is a wrapper around the C record PangoAttrList.
type AttrList struct {
	native unsafe.Pointer
}

func AttrListNewFromC(u unsafe.Pointer) *AttrList {
	if u == nil {
		return nil
	}

	g := &AttrList{native: u}

	return g
}

func (recv *AttrList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_attr_list_new : return type :

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

func AttrShapeNewFromC(u unsafe.Pointer) *AttrShape {
	if u == nil {
		return nil
	}

	g := &AttrShape{native: u}

	return g
}

func (recv *AttrShape) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrSize is a wrapper around the C record PangoAttrSize.
type AttrSize struct {
	native unsafe.Pointer
	// attr : record
	Size int32
	// Bitfield not supported :  1 absolute
}

func AttrSizeNewFromC(u unsafe.Pointer) *AttrSize {
	if u == nil {
		return nil
	}

	g := &AttrSize{native: u}

	return g
}

func (recv *AttrSize) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrString is a wrapper around the C record PangoAttrString.
type AttrString struct {
	native unsafe.Pointer
	// attr : record
	Value string
}

func AttrStringNewFromC(u unsafe.Pointer) *AttrString {
	if u == nil {
		return nil
	}

	g := &AttrString{native: u}

	return g
}

func (recv *AttrString) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Attribute is a wrapper around the C record PangoAttribute.
type Attribute struct {
	native unsafe.Pointer
	// klass : record
	StartIndex uint32
	EndIndex   uint32
}

func AttributeNewFromC(u unsafe.Pointer) *Attribute {
	if u == nil {
		return nil
	}

	g := &Attribute{native: u}

	return g
}

func (recv *Attribute) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Color is a wrapper around the C record PangoColor.
type Color struct {
	native unsafe.Pointer
	Red    uint16
	Green  uint16
	Blue   uint16
}

func ColorNewFromC(u unsafe.Pointer) *Color {
	if u == nil {
		return nil
	}

	g := &Color{native: u}

	return g
}

func (recv *Color) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContextClass is a wrapper around the C record PangoContextClass.
type ContextClass struct {
	native unsafe.Pointer
}

func ContextClassNewFromC(u unsafe.Pointer) *ContextClass {
	if u == nil {
		return nil
	}

	g := &ContextClass{native: u}

	return g
}

func (recv *ContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Coverage is a wrapper around the C record PangoCoverage.
type Coverage struct {
	native unsafe.Pointer
}

func CoverageNewFromC(u unsafe.Pointer) *Coverage {
	if u == nil {
		return nil
	}

	g := &Coverage{native: u}

	return g
}

func (recv *Coverage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FontDescriptionNewFromC(u unsafe.Pointer) *FontDescription {
	if u == nil {
		return nil
	}

	g := &FontDescription{native: u}

	return g
}

func (recv *FontDescription) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_font_description_new : return type :

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

func FontMetricsNewFromC(u unsafe.Pointer) *FontMetrics {
	if u == nil {
		return nil
	}

	g := &FontMetrics{native: u}

	return g
}

func (recv *FontMetrics) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_font_metrics_new : return type :

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// GlyphGeometry is a wrapper around the C record PangoGlyphGeometry.
type GlyphGeometry struct {
	native  unsafe.Pointer
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

func GlyphGeometryNewFromC(u unsafe.Pointer) *GlyphGeometry {
	if u == nil {
		return nil
	}

	g := &GlyphGeometry{native: u}

	return g
}

func (recv *GlyphGeometry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GlyphInfo is a wrapper around the C record PangoGlyphInfo.
type GlyphInfo struct {
	native unsafe.Pointer
	Glyph  Glyph
	// geometry : record
	// attr : record
}

func GlyphInfoNewFromC(u unsafe.Pointer) *GlyphInfo {
	if u == nil {
		return nil
	}

	g := &GlyphInfo{native: u}

	return g
}

func (recv *GlyphInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GlyphItem is a wrapper around the C record PangoGlyphItem.
type GlyphItem struct {
	native unsafe.Pointer
	// item : record
	// glyphs : record
}

func GlyphItemNewFromC(u unsafe.Pointer) *GlyphItem {
	if u == nil {
		return nil
	}

	g := &GlyphItem{native: u}

	return g
}

func (recv *GlyphItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoGlyphString

// GlyphVisAttr is a wrapper around the C record PangoGlyphVisAttr.
type GlyphVisAttr struct {
	native unsafe.Pointer
	// Bitfield not supported :  1 is_cluster_start
}

func GlyphVisAttrNewFromC(u unsafe.Pointer) *GlyphVisAttr {
	if u == nil {
		return nil
	}

	g := &GlyphVisAttr{native: u}

	return g
}

func (recv *GlyphVisAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ItemNewFromC(u unsafe.Pointer) *Item {
	if u == nil {
		return nil
	}

	g := &Item{native: u}

	return g
}

func (recv *Item) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_item_new : return type :

// Language is a wrapper around the C record PangoLanguage.
type Language struct {
	native unsafe.Pointer
}

func LanguageNewFromC(u unsafe.Pointer) *Language {
	if u == nil {
		return nil
	}

	g := &Language{native: u}

	return g
}

func (recv *Language) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LayoutClass is a wrapper around the C record PangoLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
}

func LayoutClassNewFromC(u unsafe.Pointer) *LayoutClass {
	if u == nil {
		return nil
	}

	g := &LayoutClass{native: u}

	return g
}

func (recv *LayoutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LayoutLineNewFromC(u unsafe.Pointer) *LayoutLine {
	if u == nil {
		return nil
	}

	g := &LayoutLine{native: u}

	return g
}

func (recv *LayoutLine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func LogAttrNewFromC(u unsafe.Pointer) *LogAttr {
	if u == nil {
		return nil
	}

	g := &LogAttr{native: u}

	return g
}

func (recv *LogAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	if u == nil {
		return nil
	}

	g := &Rectangle{native: u}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererPrivate is a wrapper around the C record PangoRendererPrivate.
type RendererPrivate struct {
	native unsafe.Pointer
}

func RendererPrivateNewFromC(u unsafe.Pointer) *RendererPrivate {
	if u == nil {
		return nil
	}

	g := &RendererPrivate{native: u}

	return g
}

func (recv *RendererPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoScriptForLang

// ScriptIter is a wrapper around the C record PangoScriptIter.
type ScriptIter struct {
	native unsafe.Pointer
}

func ScriptIterNewFromC(u unsafe.Pointer) *ScriptIter {
	if u == nil {
		return nil
	}

	g := &ScriptIter{native: u}

	return g
}

func (recv *ScriptIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TabArray is a wrapper around the C record PangoTabArray.
type TabArray struct {
	native unsafe.Pointer
}

func TabArrayNewFromC(u unsafe.Pointer) *TabArray {
	if u == nil {
		return nil
	}

	g := &TabArray{native: u}

	return g
}

func (recv *TabArray) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : pango_tab_array_new : return type :

// Unsupported : pango_tab_array_new_with_positions : unsupported parameter ... : varargs

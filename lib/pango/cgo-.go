// This is a generated file - DO NOT EDIT

package pango

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Context is a wrapper around the C record PangoContext.
type Context struct {
	native *C.PangoContext
}

func ContextNewFromC(u unsafe.Pointer) *Context {
	c := (*C.PangoContext)(u)
	if c == nil {
		return nil
	}

	g := &Context{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Context) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Context) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoEngine

// EngineLang is a wrapper around the C record PangoEngineLang.
type EngineLang struct {
	native *C.PangoEngineLang
	// Private : parent_instance
}

func EngineLangNewFromC(u unsafe.Pointer) *EngineLang {
	c := (*C.PangoEngineLang)(u)
	if c == nil {
		return nil
	}

	g := &EngineLang{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EngineLang) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EngineLang) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EngineShape is a wrapper around the C record PangoEngineShape.
type EngineShape struct {
	native *C.PangoEngineShape
	// parent_instance : record
}

func EngineShapeNewFromC(u unsafe.Pointer) *EngineShape {
	c := (*C.PangoEngineShape)(u)
	if c == nil {
		return nil
	}

	g := &EngineShape{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EngineShape) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EngineShape) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Font is a wrapper around the C record PangoFont.
type Font struct {
	native *C.PangoFont
	// parent_instance : record
}

func FontNewFromC(u unsafe.Pointer) *Font {
	c := (*C.PangoFont)(u)
	if c == nil {
		return nil
	}

	g := &Font{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Font) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Font) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontFace is a wrapper around the C record PangoFontFace.
type FontFace struct {
	native *C.PangoFontFace
	// parent_instance : record
}

func FontFaceNewFromC(u unsafe.Pointer) *FontFace {
	c := (*C.PangoFontFace)(u)
	if c == nil {
		return nil
	}

	g := &FontFace{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontFace) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontFace) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontFamily is a wrapper around the C record PangoFontFamily.
type FontFamily struct {
	native *C.PangoFontFamily
	// parent_instance : record
}

func FontFamilyNewFromC(u unsafe.Pointer) *FontFamily {
	c := (*C.PangoFontFamily)(u)
	if c == nil {
		return nil
	}

	g := &FontFamily{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontFamily) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontFamily) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FontMap is a wrapper around the C record PangoFontMap.
type FontMap struct {
	native *C.PangoFontMap
	// parent_instance : record
}

func FontMapNewFromC(u unsafe.Pointer) *FontMap {
	c := (*C.PangoFontMap)(u)
	if c == nil {
		return nil
	}

	g := &FontMap{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FontMap) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FontMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Fontset is a wrapper around the C record PangoFontset.
type Fontset struct {
	native *C.PangoFontset
	// parent_instance : record
}

func FontsetNewFromC(u unsafe.Pointer) *Fontset {
	c := (*C.PangoFontset)(u)
	if c == nil {
		return nil
	}

	g := &Fontset{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Fontset) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Fontset) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoFontsetSimple

// Layout is a wrapper around the C record PangoLayout.
type Layout struct {
	native *C.PangoLayout
}

func LayoutNewFromC(u unsafe.Pointer) *Layout {
	c := (*C.PangoLayout)(u)
	if c == nil {
		return nil
	}

	g := &Layout{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Layout) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : ENGINE_TYPE_LANG

// Blacklisted : ENGINE_TYPE_SHAPE

// Unsupported : type Glyph for GLYPH_EMPTY

// Unsupported : type Glyph for GLYPH_UNKNOWN_FLAG

// Blacklisted : RENDER_TYPE_NONE

// Blacklisted : UNKNOWN_GLYPH_HEIGHT

// Blacklisted : UNKNOWN_GLYPH_WIDTH

// Unsupported : pango_break : unsupported parameter attrs :

// Blacklisted : pango_config_key_get

// Blacklisted : pango_config_key_get_system

// Blacklisted : pango_default_break

// Unsupported : pango_find_map : return type : Blacklisted record : PangoMap

// Blacklisted : pango_get_lib_subdirectory

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs :

// Blacklisted : pango_get_sysconf_subdirectory

// Unsupported : pango_lookup_aliases : unsupported parameter families : output array param families

// Unsupported : pango_module_register : unsupported parameter module : Blacklisted record : PangoIncludedModule

// Unsupported : pango_parse_stretch : unsupported parameter stretch : PangoStretch* with indirection level of 1

// Unsupported : pango_parse_style : unsupported parameter style : PangoStyle* with indirection level of 1

// Unsupported : pango_parse_variant : unsupported parameter variant : PangoVariant* with indirection level of 1

// Unsupported : pango_parse_weight : unsupported parameter weight : PangoWeight* with indirection level of 1

// Unsupported : pango_read_line : unsupported parameter stream : no type generator for gpointer (FILE*) for param stream

// Unsupported : pango_scan_int : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_string : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_word : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_shape : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Analysis is a wrapper around the C record PangoAnalysis.
type Analysis struct {
	native *C.PangoAnalysis
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
	c := (*C.PangoAnalysis)(u)
	if c == nil {
		return nil
	}

	g := &Analysis{
		Flags:   (uint8)(c.flags),
		Gravity: (uint8)(c.gravity),
		Level:   (uint8)(c.level),
		Script:  (uint8)(c.script),
		native:  c,
	}

	return g
}

func (recv *Analysis) ToC() unsafe.Pointer {
	recv.native.level =
		(C.guint8)(recv.Level)
	recv.native.gravity =
		(C.guint8)(recv.Gravity)
	recv.native.flags =
		(C.guint8)(recv.Flags)
	recv.native.script =
		(C.guint8)(recv.Script)

	return (unsafe.Pointer)(recv.native)
}

// AttrClass is a wrapper around the C record PangoAttrClass.
type AttrClass struct {
	native *C.PangoAttrClass
	Type   AttrType
	// no type for copy
	// no type for destroy
	// no type for equal
}

func AttrClassNewFromC(u unsafe.Pointer) *AttrClass {
	c := (*C.PangoAttrClass)(u)
	if c == nil {
		return nil
	}

	g := &AttrClass{
		Type:   (AttrType)(c._type),
		native: c,
	}

	return g
}

func (recv *AttrClass) ToC() unsafe.Pointer {
	recv.native._type =
		(C.PangoAttrType)(recv.Type)

	return (unsafe.Pointer)(recv.native)
}

// AttrColor is a wrapper around the C record PangoAttrColor.
type AttrColor struct {
	native *C.PangoAttrColor
	// attr : record
	// color : record
}

func AttrColorNewFromC(u unsafe.Pointer) *AttrColor {
	c := (*C.PangoAttrColor)(u)
	if c == nil {
		return nil
	}

	g := &AttrColor{native: c}

	return g
}

func (recv *AttrColor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrFloat is a wrapper around the C record PangoAttrFloat.
type AttrFloat struct {
	native *C.PangoAttrFloat
	// attr : record
	Value float64
}

func AttrFloatNewFromC(u unsafe.Pointer) *AttrFloat {
	c := (*C.PangoAttrFloat)(u)
	if c == nil {
		return nil
	}

	g := &AttrFloat{
		Value:  (float64)(c.value),
		native: c,
	}

	return g
}

func (recv *AttrFloat) ToC() unsafe.Pointer {
	recv.native.value =
		(C.double)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// AttrFontDesc is a wrapper around the C record PangoAttrFontDesc.
type AttrFontDesc struct {
	native *C.PangoAttrFontDesc
	// attr : record
	// desc : record
}

func AttrFontDescNewFromC(u unsafe.Pointer) *AttrFontDesc {
	c := (*C.PangoAttrFontDesc)(u)
	if c == nil {
		return nil
	}

	g := &AttrFontDesc{native: c}

	return g
}

func (recv *AttrFontDesc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrInt is a wrapper around the C record PangoAttrInt.
type AttrInt struct {
	native *C.PangoAttrInt
	// attr : record
	Value int32
}

func AttrIntNewFromC(u unsafe.Pointer) *AttrInt {
	c := (*C.PangoAttrInt)(u)
	if c == nil {
		return nil
	}

	g := &AttrInt{
		Value:  (int32)(c.value),
		native: c,
	}

	return g
}

func (recv *AttrInt) ToC() unsafe.Pointer {
	recv.native.value =
		(C.int)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// AttrIterator is a wrapper around the C record PangoAttrIterator.
type AttrIterator struct {
	native *C.PangoAttrIterator
}

func AttrIteratorNewFromC(u unsafe.Pointer) *AttrIterator {
	c := (*C.PangoAttrIterator)(u)
	if c == nil {
		return nil
	}

	g := &AttrIterator{native: c}

	return g
}

func (recv *AttrIterator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrLanguage is a wrapper around the C record PangoAttrLanguage.
type AttrLanguage struct {
	native *C.PangoAttrLanguage
	// attr : record
	// value : record
}

func AttrLanguageNewFromC(u unsafe.Pointer) *AttrLanguage {
	c := (*C.PangoAttrLanguage)(u)
	if c == nil {
		return nil
	}

	g := &AttrLanguage{native: c}

	return g
}

func (recv *AttrLanguage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrList is a wrapper around the C record PangoAttrList.
type AttrList struct {
	native *C.PangoAttrList
}

func AttrListNewFromC(u unsafe.Pointer) *AttrList {
	c := (*C.PangoAttrList)(u)
	if c == nil {
		return nil
	}

	g := &AttrList{native: c}

	return g
}

func (recv *AttrList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrShape is a wrapper around the C record PangoAttrShape.
type AttrShape struct {
	native *C.PangoAttrShape
	// attr : record
	// ink_rect : record
	// logical_rect : record
	// data : no type generator for gpointer, gpointer
	// copy_func : no type generator for AttrDataCopyFunc, PangoAttrDataCopyFunc
	// destroy_func : no type generator for GLib.DestroyNotify, GDestroyNotify
}

func AttrShapeNewFromC(u unsafe.Pointer) *AttrShape {
	c := (*C.PangoAttrShape)(u)
	if c == nil {
		return nil
	}

	g := &AttrShape{native: c}

	return g
}

func (recv *AttrShape) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AttrSize is a wrapper around the C record PangoAttrSize.
type AttrSize struct {
	native *C.PangoAttrSize
	// attr : record
	Size int32
	// Bitfield not supported :  1 absolute
}

func AttrSizeNewFromC(u unsafe.Pointer) *AttrSize {
	c := (*C.PangoAttrSize)(u)
	if c == nil {
		return nil
	}

	g := &AttrSize{
		Size:   (int32)(c.size),
		native: c,
	}

	return g
}

func (recv *AttrSize) ToC() unsafe.Pointer {
	recv.native.size =
		(C.int)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// AttrString is a wrapper around the C record PangoAttrString.
type AttrString struct {
	native *C.PangoAttrString
	// attr : record
	Value string
}

func AttrStringNewFromC(u unsafe.Pointer) *AttrString {
	c := (*C.PangoAttrString)(u)
	if c == nil {
		return nil
	}

	g := &AttrString{
		Value:  C.GoString(c.value),
		native: c,
	}

	return g
}

func (recv *AttrString) ToC() unsafe.Pointer {
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Attribute is a wrapper around the C record PangoAttribute.
type Attribute struct {
	native *C.PangoAttribute
	// klass : record
	StartIndex uint32
	EndIndex   uint32
}

func AttributeNewFromC(u unsafe.Pointer) *Attribute {
	c := (*C.PangoAttribute)(u)
	if c == nil {
		return nil
	}

	g := &Attribute{
		EndIndex:   (uint32)(c.end_index),
		StartIndex: (uint32)(c.start_index),
		native:     c,
	}

	return g
}

func (recv *Attribute) ToC() unsafe.Pointer {
	recv.native.start_index =
		(C.guint)(recv.StartIndex)
	recv.native.end_index =
		(C.guint)(recv.EndIndex)

	return (unsafe.Pointer)(recv.native)
}

// Color is a wrapper around the C record PangoColor.
type Color struct {
	native *C.PangoColor
	Red    uint16
	Green  uint16
	Blue   uint16
}

func ColorNewFromC(u unsafe.Pointer) *Color {
	c := (*C.PangoColor)(u)
	if c == nil {
		return nil
	}

	g := &Color{
		Blue:   (uint16)(c.blue),
		Green:  (uint16)(c.green),
		Red:    (uint16)(c.red),
		native: c,
	}

	return g
}

func (recv *Color) ToC() unsafe.Pointer {
	recv.native.red =
		(C.guint16)(recv.Red)
	recv.native.green =
		(C.guint16)(recv.Green)
	recv.native.blue =
		(C.guint16)(recv.Blue)

	return (unsafe.Pointer)(recv.native)
}

// ContextClass is a wrapper around the C record PangoContextClass.
type ContextClass struct {
	native *C.PangoContextClass
}

func ContextClassNewFromC(u unsafe.Pointer) *ContextClass {
	c := (*C.PangoContextClass)(u)
	if c == nil {
		return nil
	}

	g := &ContextClass{native: c}

	return g
}

func (recv *ContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Coverage is a wrapper around the C record PangoCoverage.
type Coverage struct {
	native *C.PangoCoverage
}

func CoverageNewFromC(u unsafe.Pointer) *Coverage {
	c := (*C.PangoCoverage)(u)
	if c == nil {
		return nil
	}

	g := &Coverage{native: c}

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
	native *C.PangoFontDescription
}

func FontDescriptionNewFromC(u unsafe.Pointer) *FontDescription {
	c := (*C.PangoFontDescription)(u)
	if c == nil {
		return nil
	}

	g := &FontDescription{native: c}

	return g
}

func (recv *FontDescription) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoFontFaceClass

// Blacklisted : PangoFontFamilyClass

// Blacklisted : PangoFontMapClass

// FontMetrics is a wrapper around the C record PangoFontMetrics.
type FontMetrics struct {
	native *C.PangoFontMetrics
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
	c := (*C.PangoFontMetrics)(u)
	if c == nil {
		return nil
	}

	g := &FontMetrics{native: c}

	return g
}

func (recv *FontMetrics) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoFontsetClass

// Blacklisted : PangoFontsetSimpleClass

// GlyphGeometry is a wrapper around the C record PangoGlyphGeometry.
type GlyphGeometry struct {
	native  *C.PangoGlyphGeometry
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}

func GlyphGeometryNewFromC(u unsafe.Pointer) *GlyphGeometry {
	c := (*C.PangoGlyphGeometry)(u)
	if c == nil {
		return nil
	}

	g := &GlyphGeometry{
		Width:   (GlyphUnit)(c.width),
		XOffset: (GlyphUnit)(c.x_offset),
		YOffset: (GlyphUnit)(c.y_offset),
		native:  c,
	}

	return g
}

func (recv *GlyphGeometry) ToC() unsafe.Pointer {
	recv.native.width =
		(C.PangoGlyphUnit)(recv.Width)
	recv.native.x_offset =
		(C.PangoGlyphUnit)(recv.XOffset)
	recv.native.y_offset =
		(C.PangoGlyphUnit)(recv.YOffset)

	return (unsafe.Pointer)(recv.native)
}

// GlyphInfo is a wrapper around the C record PangoGlyphInfo.
type GlyphInfo struct {
	native *C.PangoGlyphInfo
	Glyph  Glyph
	// geometry : record
	// attr : record
}

func GlyphInfoNewFromC(u unsafe.Pointer) *GlyphInfo {
	c := (*C.PangoGlyphInfo)(u)
	if c == nil {
		return nil
	}

	g := &GlyphInfo{
		Glyph:  (Glyph)(c.glyph),
		native: c,
	}

	return g
}

func (recv *GlyphInfo) ToC() unsafe.Pointer {
	recv.native.glyph =
		(C.PangoGlyph)(recv.Glyph)

	return (unsafe.Pointer)(recv.native)
}

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

// Blacklisted : PangoGlyphString

// GlyphVisAttr is a wrapper around the C record PangoGlyphVisAttr.
type GlyphVisAttr struct {
	native *C.PangoGlyphVisAttr
	// Bitfield not supported :  1 is_cluster_start
}

func GlyphVisAttrNewFromC(u unsafe.Pointer) *GlyphVisAttr {
	c := (*C.PangoGlyphVisAttr)(u)
	if c == nil {
		return nil
	}

	g := &GlyphVisAttr{native: c}

	return g
}

func (recv *GlyphVisAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoIncludedModule

// Item is a wrapper around the C record PangoItem.
type Item struct {
	native   *C.PangoItem
	Offset   int32
	Length   int32
	NumChars int32
	// analysis : record
}

func ItemNewFromC(u unsafe.Pointer) *Item {
	c := (*C.PangoItem)(u)
	if c == nil {
		return nil
	}

	g := &Item{
		Length:   (int32)(c.length),
		NumChars: (int32)(c.num_chars),
		Offset:   (int32)(c.offset),
		native:   c,
	}

	return g
}

func (recv *Item) ToC() unsafe.Pointer {
	recv.native.offset =
		(C.gint)(recv.Offset)
	recv.native.length =
		(C.gint)(recv.Length)
	recv.native.num_chars =
		(C.gint)(recv.NumChars)

	return (unsafe.Pointer)(recv.native)
}

// Language is a wrapper around the C record PangoLanguage.
type Language struct {
	native *C.PangoLanguage
}

func LanguageNewFromC(u unsafe.Pointer) *Language {
	c := (*C.PangoLanguage)(u)
	if c == nil {
		return nil
	}

	g := &Language{native: c}

	return g
}

func (recv *Language) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LayoutClass is a wrapper around the C record PangoLayoutClass.
type LayoutClass struct {
	native *C.PangoLayoutClass
}

func LayoutClassNewFromC(u unsafe.Pointer) *LayoutClass {
	c := (*C.PangoLayoutClass)(u)
	if c == nil {
		return nil
	}

	g := &LayoutClass{native: c}

	return g
}

func (recv *LayoutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoLayoutIter

// LayoutLine is a wrapper around the C record PangoLayoutLine.
type LayoutLine struct {
	native *C.PangoLayoutLine
	// layout : record
	StartIndex int32
	Length     int32
	// runs : record
	// Bitfield not supported :  1 is_paragraph_start
	// Bitfield not supported :  3 resolved_dir
}

func LayoutLineNewFromC(u unsafe.Pointer) *LayoutLine {
	c := (*C.PangoLayoutLine)(u)
	if c == nil {
		return nil
	}

	g := &LayoutLine{
		Length:     (int32)(c.length),
		StartIndex: (int32)(c.start_index),
		native:     c,
	}

	return g
}

func (recv *LayoutLine) ToC() unsafe.Pointer {
	recv.native.start_index =
		(C.gint)(recv.StartIndex)
	recv.native.length =
		(C.gint)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// LogAttr is a wrapper around the C record PangoLogAttr.
type LogAttr struct {
	native *C.PangoLogAttr
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
	c := (*C.PangoLogAttr)(u)
	if c == nil {
		return nil
	}

	g := &LogAttr{native: c}

	return g
}

func (recv *LogAttr) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

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

// RendererPrivate is a wrapper around the C record PangoRendererPrivate.
type RendererPrivate struct {
	native *C.PangoRendererPrivate
}

func RendererPrivateNewFromC(u unsafe.Pointer) *RendererPrivate {
	c := (*C.PangoRendererPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RendererPrivate{native: c}

	return g
}

func (recv *RendererPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : PangoScriptForLang

// ScriptIter is a wrapper around the C record PangoScriptIter.
type ScriptIter struct {
	native *C.PangoScriptIter
}

func ScriptIterNewFromC(u unsafe.Pointer) *ScriptIter {
	c := (*C.PangoScriptIter)(u)
	if c == nil {
		return nil
	}

	g := &ScriptIter{native: c}

	return g
}

func (recv *ScriptIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TabArray is a wrapper around the C record PangoTabArray.
type TabArray struct {
	native *C.PangoTabArray
}

func TabArrayNewFromC(u unsafe.Pointer) *TabArray {
	c := (*C.PangoTabArray)(u)
	if c == nil {
		return nil
	}

	g := &TabArray{native: c}

	return g
}

func (recv *TabArray) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

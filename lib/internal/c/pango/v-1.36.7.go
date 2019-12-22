// Code generated - DO NOT EDIT.
// +build pango_1.36.7

package pango

import (
	c "github.com/pekim/gobbi/lib/internal/c"
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	"unsafe"
)

// #include <pango/pango.h>
import "C"

// aliases
type Glyph C.PangoGlyph
type GlyphUnit C.PangoGlyphUnit
type LayoutRun C.PangoLayoutRun

// bitfields
type FontMask C.PangoFontMask

// enumerations
type Alignment C.PangoAlignment
type AttrType C.PangoAttrType
type BidiType C.PangoBidiType
type CoverageLevel C.PangoCoverageLevel
type Direction C.PangoDirection
type EllipsizeMode C.PangoEllipsizeMode
type Gravity C.PangoGravity
type GravityHint C.PangoGravityHint
type RenderPart C.PangoRenderPart
type Script C.PangoScript
type Stretch C.PangoStretch
type Style C.PangoStyle
type TabAlign C.PangoTabAlign
type Underline C.PangoUnderline
type Variant C.PangoVariant
type Weight C.PangoWeight
type WrapMode C.PangoWrapMode

// unions

// records
type Analysis C.PangoAnalysis
type AttrClass C.PangoAttrClass
type AttrColor C.PangoAttrColor
type AttrFloat C.PangoAttrFloat
type AttrFontDesc C.PangoAttrFontDesc
type AttrInt C.PangoAttrInt
type AttrIterator C.PangoAttrIterator
type AttrLanguage C.PangoAttrLanguage
type AttrList C.PangoAttrList
type AttrShape C.PangoAttrShape
type AttrSize C.PangoAttrSize
type AttrString C.PangoAttrString
type Attribute C.PangoAttribute
type Color C.PangoColor
type ContextClass C.PangoContextClass
type Coverage C.PangoCoverage

// UNSUPPORTED : EngineClass : blacklisted
// UNSUPPORTED : EngineInfo : blacklisted
// UNSUPPORTED : EngineLangClass : blacklisted
// UNSUPPORTED : EngineScriptInfo : blacklisted
// UNSUPPORTED : EngineShapeClass : blacklisted
// UNSUPPORTED : FontClass : blacklisted
type FontDescription C.PangoFontDescription

// UNSUPPORTED : FontFaceClass : blacklisted
// UNSUPPORTED : FontFamilyClass : blacklisted
// UNSUPPORTED : FontMapClass : blacklisted
type FontMetrics C.PangoFontMetrics

// UNSUPPORTED : FontsetClass : blacklisted
// UNSUPPORTED : FontsetSimpleClass : blacklisted
type GlyphGeometry C.PangoGlyphGeometry
type GlyphInfo C.PangoGlyphInfo
type GlyphItem C.PangoGlyphItem
type GlyphItemIter C.PangoGlyphItemIter
type GlyphString C.PangoGlyphString
type GlyphVisAttr C.PangoGlyphVisAttr

// UNSUPPORTED : IncludedModule : blacklisted
type Item C.PangoItem
type Language C.PangoLanguage
type LayoutClass C.PangoLayoutClass
type LayoutIter C.PangoLayoutIter
type LayoutLine C.PangoLayoutLine
type LogAttr C.PangoLogAttr

// UNSUPPORTED : Map : blacklisted
// UNSUPPORTED : MapEntry : blacklisted
type Matrix C.PangoMatrix
type Rectangle C.PangoRectangle
type RendererClass C.PangoRendererClass
type RendererPrivate C.PangoRendererPrivate
type ScriptIter C.PangoScriptIter
type TabArray C.PangoTabArray

// classes
type Context C.PangoContext

// UNSUPPORTED : Engine : blacklisted
type EngineLang C.PangoEngineLang
type EngineShape C.PangoEngineShape
type Font C.PangoFont
type FontFace C.PangoFontFace
type FontFamily C.PangoFontFamily
type FontMap C.PangoFontMap
type Fontset C.PangoFontset

// UNSUPPORTED : FontsetSimple : blacklisted
type Layout C.PangoLayout
type Renderer C.PangoRenderer

// interfaces

func Fn_attr_background_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_fallback_new(enableFallback bool) {}

func Fn_attr_family_new(family string) {}

func Fn_attr_foreground_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_gravity_hint_new(hint GravityHint) {}

func Fn_attr_gravity_new(gravity Gravity) {}

func Fn_attr_letter_spacing_new(letterSpacing int) {}

func Fn_attr_rise_new(rise int) {}

func Fn_attr_scale_new(scaleFactor float64) {}

func Fn_attr_stretch_new(stretch Stretch) {}

func Fn_attr_strikethrough_color_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_strikethrough_new(strikethrough bool) {}

func Fn_attr_style_new(style Style) {}

func Fn_attr_type_get_name(type_ AttrType) {}

func Fn_attr_type_register(name string) {}

func Fn_attr_underline_color_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_underline_new(underline Underline) {}

func Fn_attr_variant_new(variant Variant) {}

func Fn_attr_weight_new(weight Weight) {}

func Fn_bidi_type_for_unichar(ch rune) {}

func Fn_break(text string, length int, analysis *Analysis, attrs c.UndefinedParamType, attrsLen int) {}

func Fn_config_key_get(key string) {}

func Fn_config_key_get_system(key string) {}

func Fn_default_break(text string, length int, analysis *Analysis, attrs *LogAttr, attrsLen int) {}

func Fn_extents_to_pixels(inclusive *Rectangle, nearest *Rectangle) {}

func Fn_find_base_dir(text string, length int) {}

func Fn_find_map(language *Language, engineTypeId uint, renderTypeId uint) {}

func Fn_find_paragraph_boundary(text string, length int) {}

func Fn_font_description_from_string(str string) {}

func Fn_get_lib_subdirectory() {}

func Fn_get_log_attrs(text string, length int, level int, language *Language, logAttrs c.UndefinedParamType, attrsLen int) {
}

func Fn_get_mirror_char(ch rune, mirroredCh *rune) {}

func Fn_get_sysconf_subdirectory() {}

func Fn_gravity_get_for_matrix(matrix *Matrix) {}

func Fn_gravity_get_for_script(script Script, baseGravity Gravity, hint GravityHint) {}

func Fn_gravity_get_for_script_and_width(script Script, wide bool, baseGravity Gravity, hint GravityHint) {
}

func Fn_gravity_to_rotation(gravity Gravity) {}

func Fn_is_zero_width(ch rune) {}

func Fn_itemize(context *Context, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) {
}

func Fn_itemize_with_base_dir(context *Context, baseDir Direction, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) {
}

func Fn_language_from_string(language string) {}

func Fn_language_get_default() {}

func Fn_log2vis_get_embedding_levels(text string, length int, pbaseDir *Direction) {}

func Fn_lookup_aliases(fontname string) {}

func Fn_markup_parser_finish(context *glib.MarkupParseContext) {}

func Fn_markup_parser_new(accelMarker rune) {}

// UNSUPPORTED : module_register : blacklisted
func Fn_parse_enum(type_ glib.Type, str string, warn bool) {}

func Fn_parse_markup(markupText string, length int, accelMarker rune) {}

func Fn_parse_stretch(str string, warn bool) {}

func Fn_parse_style(str string, warn bool) {}

func Fn_parse_variant(str string, warn bool) {}

func Fn_parse_weight(str string, warn bool) {}

func Fn_quantize_line_geometry(thickness *int, position *int) {}

func Fn_read_line(stream unsafe.Pointer) {}

func Fn_reorder_items(logicalItems *glib.List) {}

func Fn_scan_int(pos string) {}

func Fn_scan_string(pos string) {}

func Fn_scan_word(pos string) {}

func Fn_script_for_unichar(ch rune) {}

func Fn_script_get_sample_language(script Script) {}

func Fn_shape(text string, length int, analysis *Analysis, glyphs *GlyphString) {}

func Fn_shape_full(itemText string, itemLength int, paragraphText string, paragraphLength int, analysis *Analysis, glyphs *GlyphString) {
}

func Fn_skip_space(pos string) {}

func Fn_split_file_list(str string) {}

func Fn_trim_string(str string) {}

func Fn_unichar_direction(ch rune) {}

func Fn_units_from_double(d float64) {}

func Fn_units_to_double(i int) {}

func Fn_version() {}

func Fn_version_check(requiredMajor int, requiredMinor int, requiredMicro int) {}

func Fn_version_string() {}

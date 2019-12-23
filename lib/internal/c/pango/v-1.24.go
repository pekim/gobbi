// Code generated - DO NOT EDIT.
// +build pango_1.24

package pango

import "unsafe"

// #include <pango/pango.h>
import "C"

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

func Fn_attr_background_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_fallback_new(enableFallback bool) {}

func Fn_attr_family_new(family string) {}

func Fn_attr_foreground_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_gravity_hint_new(hint int) {}

func Fn_attr_gravity_new(gravity int) {}

func Fn_attr_letter_spacing_new(letterSpacing int) {}

func Fn_attr_rise_new(rise int) {}

func Fn_attr_scale_new(scaleFactor float64) {}

func Fn_attr_stretch_new(stretch int) {}

func Fn_attr_strikethrough_color_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_strikethrough_new(strikethrough bool) {}

func Fn_attr_style_new(style int) {}

func Fn_attr_type_get_name(type_ int) {}

func Fn_attr_type_register(name string) {}

func Fn_attr_underline_color_new(red uint16, green uint16, blue uint16) {}

func Fn_attr_underline_new(underline int) {}

func Fn_attr_variant_new(variant int) {}

func Fn_attr_weight_new(weight int) {}

func Fn_bidi_type_for_unichar(ch rune) {}

func Fn_break(text string, length int, analysis unsafe.Pointer, attrs *LogAttr, attrsLen int) {}

func Fn_config_key_get(key string) {}

func Fn_config_key_get_system(key string) {}

func Fn_default_break(text string, length int, analysis unsafe.Pointer, attrs unsafe.Pointer, attrsLen int) {
}

func Fn_extents_to_pixels(inclusive unsafe.Pointer, nearest unsafe.Pointer) {}

func Fn_find_base_dir(text string, length int) {}

func Fn_find_map(language unsafe.Pointer, engineTypeId uint, renderTypeId uint) {}

func Fn_find_paragraph_boundary(text string, length int, paragraphDelimiterIndex *int, nextParagraphStart *int) {
}

func Fn_font_description_from_string(str string) {}

// UNSUPPORTED : get_lib_subdirectory : blacklisted
func Fn_get_log_attrs(text string, length int, level int, language unsafe.Pointer, logAttrs *LogAttr, attrsLen int) {
}

func Fn_get_mirror_char(ch rune, mirroredCh *rune) {}

// UNSUPPORTED : get_sysconf_subdirectory : blacklisted
func Fn_gravity_get_for_matrix(matrix unsafe.Pointer) {}

func Fn_gravity_get_for_script(script int, baseGravity int, hint int) {}

func Fn_gravity_to_rotation(gravity int) {}

func Fn_is_zero_width(ch rune) {}

func Fn_itemize(context unsafe.Pointer, text string, startIndex int, length int, attrs unsafe.Pointer, cachedIter unsafe.Pointer) {
}

func Fn_itemize_with_base_dir(context unsafe.Pointer, baseDir int, text string, startIndex int, length int, attrs unsafe.Pointer, cachedIter unsafe.Pointer) {
}

func Fn_language_from_string(language string) {}

func Fn_language_get_default() {
	C.pango_language_get_default()
}

func Fn_log2vis_get_embedding_levels(text string, length int, pbaseDir int) {}

func Fn_lookup_aliases(fontname string, families *string, nFamilies *int) {}

// UNSUPPORTED : module_register : blacklisted
func Fn_parse_enum(type_ uint64, str string, value *int, warn bool, possibleValues string) {}

func Fn_parse_markup(markupText string, length int, accelMarker rune, attrList *unsafe.Pointer, text string, accelChar *rune) {
}

func Fn_parse_stretch(str string, stretch int, warn bool) {}

func Fn_parse_style(str string, style int, warn bool) {}

func Fn_parse_variant(str string, variant int, warn bool) {}

func Fn_parse_weight(str string, weight int, warn bool) {}

func Fn_quantize_line_geometry(thickness *int, position *int) {}

func Fn_read_line(stream unsafe.Pointer, str unsafe.Pointer) {}

func Fn_reorder_items(logicalItems unsafe.Pointer) {}

func Fn_scan_int(pos string, out *int) {}

func Fn_scan_string(pos string, out unsafe.Pointer) {}

func Fn_scan_word(pos string, out unsafe.Pointer) {}

func Fn_script_for_unichar(ch rune) {}

func Fn_script_get_sample_language(script int) {}

func Fn_shape(text string, length int, analysis unsafe.Pointer, glyphs unsafe.Pointer) {}

func Fn_skip_space(pos string) {}

func Fn_split_file_list(str string) {}

func Fn_trim_string(str string) {}

func Fn_unichar_direction(ch rune) {}

func Fn_units_from_double(d float64) {}

func Fn_units_to_double(i int) {}

func Fn_version() {
	C.pango_version()
}

func Fn_version_check(requiredMajor int, requiredMinor int, requiredMicro int) {}

func Fn_version_string() {
	C.pango_version_string()
}

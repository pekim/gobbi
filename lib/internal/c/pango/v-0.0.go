// Code generated - DO NOT EDIT.
// +build !pango_1.6,!pango_1.8,!pango_1.16,!pango_1.20,!pango_1.22,!pango_1.24,!pango_1.36.7,!pango_1.38,!pango_1.42

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
type Rectangle C.PangoRectangle
type RendererPrivate C.PangoRendererPrivate
type ScriptIter C.PangoScriptIter
type TabArray C.PangoTabArray

func Fn_attr_background_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_attr_family_new(param0 string) {}

func Fn_attr_foreground_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_attr_rise_new(param0 int) {}

func Fn_attr_scale_new(param0 float64) {}

func Fn_attr_stretch_new(param0 int) {}

func Fn_attr_strikethrough_new(param0 bool) {}

func Fn_attr_style_new(param0 int) {}

func Fn_attr_type_register(param0 string) {}

func Fn_attr_underline_new(param0 int) {}

func Fn_attr_variant_new(param0 int) {}

func Fn_attr_weight_new(param0 int) {}

func Fn_break(param0 string, param1 int, param2 unsafe.Pointer, param3 []LogAttr, param4 int) {}

func Fn_config_key_get(param0 string) {}

func Fn_config_key_get_system(param0 string) {}

func Fn_default_break(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_find_map(param0 unsafe.Pointer, param1 uint, param2 uint) {}

func Fn_find_paragraph_boundary(param0 string, param1 int, param2 *int, param3 *int) {}

func Fn_font_description_from_string(param0 string) {}

// UNSUPPORTED : get_lib_subdirectory : blacklisted
func Fn_get_log_attrs(param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 []LogAttr, param5 int) {
}

func Fn_get_mirror_char(param0 rune, param1 *rune) {}

// UNSUPPORTED : get_sysconf_subdirectory : blacklisted
func Fn_itemize(param0 unsafe.Pointer, param1 string, param2 int, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_language_from_string(param0 string) {}

func Fn_lookup_aliases(param0 string, param1 *[]string, param2 *int) {}

// UNSUPPORTED : module_register : blacklisted
func Fn_parse_markup(param0 string, param1 int, param2 rune, param3 *unsafe.Pointer, param4 string, param5 *rune) {
}

func Fn_parse_stretch(param0 string, param1 int, param2 bool) {}

func Fn_parse_style(param0 string, param1 int, param2 bool) {}

func Fn_parse_variant(param0 string, param1 int, param2 bool) {}

func Fn_parse_weight(param0 string, param1 int, param2 bool) {}

func Fn_read_line(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_reorder_items(param0 unsafe.Pointer) {}

func Fn_scan_int(param0 string, param1 *int) {}

func Fn_scan_string(param0 string, param1 unsafe.Pointer) {}

func Fn_scan_word(param0 string, param1 unsafe.Pointer) {}

func Fn_shape(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {}

func Fn_skip_space(param0 string) {}

func Fn_split_file_list(param0 string) {}

func Fn_trim_string(param0 string) {}

func Fn_unichar_direction(param0 rune) {}

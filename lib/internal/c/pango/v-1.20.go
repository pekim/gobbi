// Code generated - DO NOT EDIT.
// +build pango_1.20

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
type Matrix C.PangoMatrix
type Rectangle C.PangoRectangle
type RendererClass C.PangoRendererClass
type RendererPrivate C.PangoRendererPrivate
type ScriptIter C.PangoScriptIter
type TabArray C.PangoTabArray

func Fn_attr_background_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_attr_fallback_new(param0 bool) {}

func Fn_attr_family_new(param0 string) {}

func Fn_attr_foreground_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_attr_gravity_hint_new(param0 int) {}

func Fn_attr_gravity_new(param0 int) {}

func Fn_attr_letter_spacing_new(param0 int) {}

func Fn_attr_rise_new(param0 int) {}

func Fn_attr_scale_new(param0 float64) {}

func Fn_attr_stretch_new(param0 int) {}

func Fn_attr_strikethrough_color_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_attr_strikethrough_new(param0 bool) {}

func Fn_attr_style_new(param0 int) {}

func Fn_attr_type_register(param0 string) {}

func Fn_attr_underline_color_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_attr_underline_new(param0 int) {}

func Fn_attr_variant_new(param0 int) {}

func Fn_attr_weight_new(param0 int) {}

func Fn_break(param0 string, param1 int, param2 unsafe.Pointer, param3 []LogAttr, param4 int) {}

func Fn_config_key_get(param0 string) {}

func Fn_config_key_get_system(param0 string) {}

func Fn_default_break(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_extents_to_pixels(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_find_base_dir(param0 string, param1 int) {}

func Fn_find_map(param0 unsafe.Pointer, param1 uint, param2 uint) {}

func Fn_find_paragraph_boundary(param0 string, param1 int, param2 *int, param3 *int) {}

func Fn_font_description_from_string(param0 string) {}

// UNSUPPORTED : get_lib_subdirectory : blacklisted
func Fn_get_log_attrs(param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 []LogAttr, param5 int) {
}

func Fn_get_mirror_char(param0 rune, param1 *rune) {}

// UNSUPPORTED : get_sysconf_subdirectory : blacklisted
func Fn_gravity_get_for_matrix(param0 unsafe.Pointer) {}

func Fn_gravity_get_for_script(param0 int, param1 int, param2 int) {}

func Fn_gravity_to_rotation(param0 int) {}

func Fn_is_zero_width(param0 rune) {}

func Fn_itemize(param0 unsafe.Pointer, param1 string, param2 int, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_itemize_with_base_dir(param0 unsafe.Pointer, param1 int, param2 string, param3 int, param4 int, param5 unsafe.Pointer, param6 unsafe.Pointer) {
}

func Fn_language_from_string(param0 string) {}

func Fn_language_get_default() {
	C.pango_language_get_default()
}

func Fn_log2vis_get_embedding_levels(param0 string, param1 int, param2 int) {}

func Fn_lookup_aliases(param0 string, param1 *[]string, param2 *int) {}

// UNSUPPORTED : module_register : blacklisted
func Fn_parse_enum(param0 uint64, param1 string, param2 *int, param3 bool, param4 string) {}

func Fn_parse_markup(param0 string, param1 int, param2 rune, param3 *unsafe.Pointer, param4 string, param5 *rune) {
}

func Fn_parse_stretch(param0 string, param1 int, param2 bool) {}

func Fn_parse_style(param0 string, param1 int, param2 bool) {}

func Fn_parse_variant(param0 string, param1 int, param2 bool) {}

func Fn_parse_weight(param0 string, param1 int, param2 bool) {}

func Fn_quantize_line_geometry(param0 *int, param1 *int) {}

func Fn_read_line(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_reorder_items(param0 unsafe.Pointer) {}

func Fn_scan_int(param0 string, param1 *int) {}

func Fn_scan_string(param0 string, param1 unsafe.Pointer) {}

func Fn_scan_word(param0 string, param1 unsafe.Pointer) {}

func Fn_script_for_unichar(param0 rune) {}

func Fn_script_get_sample_language(param0 int) {}

func Fn_shape(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {}

func Fn_skip_space(param0 string) {}

func Fn_split_file_list(param0 string) {}

func Fn_trim_string(param0 string) {}

func Fn_unichar_direction(param0 rune) {}

func Fn_units_from_double(param0 float64) {}

func Fn_units_to_double(param0 int) {}

func Fn_version() {
	C.pango_version()
}

func Fn_version_check(param0 int, param1 int, param2 int) {}

func Fn_version_string() {
	C.pango_version_string()
}

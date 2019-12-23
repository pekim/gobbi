// Code generated - DO NOT EDIT.
// +build !pango_1.6,!pango_1.8,!pango_1.16,!pango_1.20,!pango_1.22,!pango_1.24,!pango_1.36.7,!pango_1.38,!pango_1.42

package pango

import "unsafe"

// #include <pango/pango.h>
import "C"

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

func Fn_pango_attr_background_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_pango_attr_family_new(param0 string) {}

func Fn_pango_attr_foreground_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_pango_attr_rise_new(param0 int) {}

func Fn_pango_attr_scale_new(param0 float64) {}

func Fn_pango_attr_stretch_new(param0 int) {}

func Fn_pango_attr_strikethrough_new(param0 bool) {}

func Fn_pango_attr_style_new(param0 int) {}

func Fn_pango_attr_type_register(param0 string) {}

func Fn_pango_attr_underline_new(param0 int) {}

func Fn_pango_attr_variant_new(param0 int) {}

func Fn_pango_attr_weight_new(param0 int) {}

func Fn_pango_break(param0 string, param1 int, param2 unsafe.Pointer, param3 []LogAttr, param4 int) {
}

func Fn_pango_config_key_get(param0 string) {}

func Fn_pango_config_key_get_system(param0 string) {}

func Fn_pango_default_break(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_pango_find_map(param0 unsafe.Pointer, param1 uint, param2 uint) {}

func Fn_pango_find_paragraph_boundary(param0 string, param1 int, param2 *int, param3 *int) {}

func Fn_pango_font_description_from_string(param0 string) {}

// UNSUPPORTED : get_lib_subdirectory : blacklisted
func Fn_pango_get_log_attrs(param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 []LogAttr, param5 int) {
}

func Fn_pango_get_mirror_char(param0 rune, param1 *rune) {}

// UNSUPPORTED : get_sysconf_subdirectory : blacklisted
func Fn_pango_itemize(param0 unsafe.Pointer, param1 string, param2 int, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_pango_language_from_string(param0 string) {}

func Fn_pango_lookup_aliases(param0 string, param1 *[]string, param2 *int) {}

// UNSUPPORTED : module_register : blacklisted
func Fn_pango_parse_markup(param0 string, param1 int, param2 rune, param3 *unsafe.Pointer, param4 string, param5 *rune) {
}

func Fn_pango_parse_stretch(param0 string, param1 int, param2 bool) {}

func Fn_pango_parse_style(param0 string, param1 int, param2 bool) {}

func Fn_pango_parse_variant(param0 string, param1 int, param2 bool) {}

func Fn_pango_parse_weight(param0 string, param1 int, param2 bool) {}

func Fn_pango_read_line(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_reorder_items(param0 unsafe.Pointer) {}

func Fn_pango_scan_int(param0 string, param1 *int) {}

func Fn_pango_scan_string(param0 string, param1 unsafe.Pointer) {}

func Fn_pango_scan_word(param0 string, param1 unsafe.Pointer) {}

func Fn_pango_shape(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {}

func Fn_pango_skip_space(param0 string) {}

func Fn_pango_split_file_list(param0 string) {}

func Fn_pango_trim_string(param0 string) {}

func Fn_pango_unichar_direction(param0 rune) {}

func Fn_pango_context_new() {
	C.pango_context_new()
}

func Fn_pango_context_get_base_dir(paramInstance unsafe.Pointer) {}

func Fn_pango_context_get_font_description(paramInstance unsafe.Pointer) {}

func Fn_pango_context_get_language(paramInstance unsafe.Pointer) {}

func Fn_pango_context_get_metrics(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_pango_context_list_families(paramInstance unsafe.Pointer, param0 []*unsafe.Pointer, param1 *int) {
}

func Fn_pango_context_load_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_context_load_fontset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_pango_context_set_base_dir(paramInstance unsafe.Pointer, param0 int) {}

func Fn_pango_context_set_font_description(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_context_set_font_map(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_context_set_language(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_font_describe(paramInstance unsafe.Pointer) {}

func Fn_pango_font_find_shaper(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint32) {}

func Fn_pango_font_get_coverage(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_font_get_glyph_extents(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_pango_font_get_metrics(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_font_descriptions_free(param0 []unsafe.Pointer, param1 int) {}

func Fn_pango_font_face_describe(paramInstance unsafe.Pointer) {}

func Fn_pango_font_face_get_face_name(paramInstance unsafe.Pointer) {}

func Fn_pango_font_family_get_name(paramInstance unsafe.Pointer) {}

func Fn_pango_font_family_list_faces(paramInstance unsafe.Pointer, param0 []*unsafe.Pointer, param1 *int) {
}

func Fn_pango_font_map_list_families(paramInstance unsafe.Pointer, param0 []*unsafe.Pointer, param1 *int) {
}

func Fn_pango_font_map_load_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_pango_font_map_load_fontset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

// UNSUPPORTED : foreach : has callback

func Fn_pango_fontset_get_font(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_pango_fontset_get_metrics(paramInstance unsafe.Pointer) {}

func Fn_pango_fontset_simple_new(param0 unsafe.Pointer) {}

func Fn_pango_fontset_simple_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_fontset_simple_size(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_new(param0 unsafe.Pointer) {}

func Fn_pango_layout_context_changed(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_copy(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_alignment(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_attributes(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_context(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_cursor_pos(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_pango_layout_get_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_pango_layout_get_indent(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_iter(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_justify(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_line(paramInstance unsafe.Pointer, param0 int) {}

func Fn_pango_layout_get_line_count(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_lines(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_log_attrs(paramInstance unsafe.Pointer, param0 []unsafe.Pointer, param1 *int) {
}

func Fn_pango_layout_get_pixel_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_pango_layout_get_pixel_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_pango_layout_get_single_paragraph_mode(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_pango_layout_get_spacing(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_tabs(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_text(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_width(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_get_wrap(paramInstance unsafe.Pointer) {}

func Fn_pango_layout_index_to_line_x(paramInstance unsafe.Pointer, param0 int, param1 bool, param2 *int, param3 *int) {
}

func Fn_pango_layout_index_to_pos(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_pango_layout_move_cursor_visually(paramInstance unsafe.Pointer, param0 bool, param1 int, param2 int, param3 int, param4 *int, param5 *int) {
}

func Fn_pango_layout_set_alignment(paramInstance unsafe.Pointer, param0 int) {}

func Fn_pango_layout_set_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_layout_set_font_description(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_layout_set_indent(paramInstance unsafe.Pointer, param0 int) {}

func Fn_pango_layout_set_justify(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_pango_layout_set_markup(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_pango_layout_set_markup_with_accel(paramInstance unsafe.Pointer, param0 string, param1 int, param2 rune, param3 *rune) {
}

func Fn_pango_layout_set_single_paragraph_mode(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_pango_layout_set_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_pango_layout_set_tabs(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_pango_layout_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_pango_layout_set_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_pango_layout_set_wrap(paramInstance unsafe.Pointer, param0 int) {}

func Fn_pango_layout_xy_to_index(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

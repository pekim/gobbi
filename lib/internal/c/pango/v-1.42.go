// Code generated - DO NOT EDIT.
// +build pango_1.42

package pango

import "unsafe"

// #include <pango/pango.h>
import "C"

type Analysis C.PangoAnalysis
type AttrClass C.PangoAttrClass
type AttrColor C.PangoAttrColor
type AttrFloat C.PangoAttrFloat
type AttrFontDesc C.PangoAttrFontDesc
type AttrFontFeatures C.PangoAttrFontFeatures
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

func Fn_pango_attr_background_alpha_new(param0 uint16) {}

func Fn_pango_attr_background_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_pango_attr_fallback_new(param0 bool) {}

func Fn_pango_attr_family_new(param0 string) {}

func Fn_pango_attr_foreground_alpha_new(param0 uint16) {}

func Fn_pango_attr_foreground_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_pango_attr_gravity_hint_new(param0 int) {}

func Fn_pango_attr_gravity_new(param0 int) {}

func Fn_pango_attr_letter_spacing_new(param0 int) {}

func Fn_pango_attr_rise_new(param0 int) {}

func Fn_pango_attr_scale_new(param0 float64) {}

func Fn_pango_attr_stretch_new(param0 int) {}

func Fn_pango_attr_strikethrough_color_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_pango_attr_strikethrough_new(param0 bool) {}

func Fn_pango_attr_style_new(param0 int) {}

func Fn_pango_attr_type_get_name(param0 int) {}

func Fn_pango_attr_type_register(param0 string) {}

func Fn_pango_attr_underline_color_new(param0 uint16, param1 uint16, param2 uint16) {}

func Fn_pango_attr_underline_new(param0 int) {}

func Fn_pango_attr_variant_new(param0 int) {}

func Fn_pango_attr_weight_new(param0 int) {}

func Fn_pango_bidi_type_for_unichar(param0 rune) {}

func Fn_pango_break(param0 string, param1 int, param2 unsafe.Pointer, param3 []LogAttr, param4 int) {
}

func Fn_pango_config_key_get(param0 string) {}

func Fn_pango_config_key_get_system(param0 string) {}

func Fn_pango_default_break(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_pango_extents_to_pixels(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_find_base_dir(param0 string, param1 int) {}

func Fn_pango_find_map(param0 unsafe.Pointer, param1 uint, param2 uint) {}

func Fn_pango_find_paragraph_boundary(param0 string, param1 int, param2 *int, param3 *int) {}

func Fn_pango_font_description_from_string(param0 string) {}

// UNSUPPORTED : get_lib_subdirectory : blacklisted
func Fn_pango_get_log_attrs(param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 []LogAttr, param5 int) {
}

func Fn_pango_get_mirror_char(param0 rune, param1 *rune) {}

// UNSUPPORTED : get_sysconf_subdirectory : blacklisted
func Fn_pango_gravity_get_for_matrix(param0 unsafe.Pointer) {}

func Fn_pango_gravity_get_for_script(param0 int, param1 int, param2 int) {}

func Fn_pango_gravity_get_for_script_and_width(param0 int, param1 bool, param2 int, param3 int) {}

func Fn_pango_gravity_to_rotation(param0 int) {}

func Fn_pango_is_zero_width(param0 rune) {}

func Fn_pango_itemize(param0 unsafe.Pointer, param1 string, param2 int, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_pango_itemize_with_base_dir(param0 unsafe.Pointer, param1 int, param2 string, param3 int, param4 int, param5 unsafe.Pointer, param6 unsafe.Pointer) {
}

func Fn_pango_language_from_string(param0 string) {}

func Fn_pango_language_get_default() {
	C.pango_language_get_default()
}

func Fn_pango_log2vis_get_embedding_levels(param0 string, param1 int, param2 int) {}

func Fn_pango_lookup_aliases(param0 string, param1 *[]string, param2 *int) {}

func Fn_pango_markup_parser_finish(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 string, param3 *rune) {
}

func Fn_pango_markup_parser_new(param0 rune) {}

// UNSUPPORTED : module_register : blacklisted
func Fn_pango_parse_enum(param0 uint64, param1 string, param2 *int, param3 bool, param4 string) {}

func Fn_pango_parse_markup(param0 string, param1 int, param2 rune, param3 *unsafe.Pointer, param4 string, param5 *rune) {
}

func Fn_pango_parse_stretch(param0 string, param1 int, param2 bool) {}

func Fn_pango_parse_style(param0 string, param1 int, param2 bool) {}

func Fn_pango_parse_variant(param0 string, param1 int, param2 bool) {}

func Fn_pango_parse_weight(param0 string, param1 int, param2 bool) {}

func Fn_pango_quantize_line_geometry(param0 *int, param1 *int) {}

func Fn_pango_read_line(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_reorder_items(param0 unsafe.Pointer) {}

func Fn_pango_scan_int(param0 string, param1 *int) {}

func Fn_pango_scan_string(param0 string, param1 unsafe.Pointer) {}

func Fn_pango_scan_word(param0 string, param1 unsafe.Pointer) {}

func Fn_pango_script_for_unichar(param0 rune) {}

func Fn_pango_script_get_sample_language(param0 int) {}

func Fn_pango_shape(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {}

func Fn_pango_shape_full(param0 string, param1 int, param2 string, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_pango_skip_space(param0 string) {}

func Fn_pango_split_file_list(param0 string) {}

func Fn_pango_trim_string(param0 string) {}

func Fn_pango_unichar_direction(param0 rune) {}

func Fn_pango_units_from_double(param0 float64) {}

func Fn_pango_units_to_double(param0 int) {}

func Fn_pango_version() {
	C.pango_version()
}

func Fn_pango_version_check(param0 int, param1 int, param2 int) {}

func Fn_pango_version_string() {
	C.pango_version_string()
}

func Fn_pango_context_new() {
	C.pango_context_new()
}

func Fn_pango_context_changed() {}

func Fn_pango_context_get_base_dir() {}

func Fn_pango_context_get_base_gravity() {}

func Fn_pango_context_get_font_description() {}

func Fn_pango_context_get_font_map() {}

func Fn_pango_context_get_gravity() {}

func Fn_pango_context_get_gravity_hint() {}

func Fn_pango_context_get_language() {}

func Fn_pango_context_get_matrix() {}

func Fn_pango_context_get_metrics(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_context_get_serial() {}

func Fn_pango_context_list_families(param0 []*unsafe.Pointer, param1 *int) {}

func Fn_pango_context_load_font(param0 unsafe.Pointer) {}

func Fn_pango_context_load_fontset(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_context_set_base_dir(param0 int) {}

func Fn_pango_context_set_base_gravity(param0 int) {}

func Fn_pango_context_set_font_description(param0 unsafe.Pointer) {}

func Fn_pango_context_set_font_map(param0 unsafe.Pointer) {}

func Fn_pango_context_set_gravity_hint(param0 int) {}

func Fn_pango_context_set_language(param0 unsafe.Pointer) {}

func Fn_pango_context_set_matrix(param0 unsafe.Pointer) {}

func Fn_pango_font_describe() {}

func Fn_pango_font_describe_with_absolute_size() {}

func Fn_pango_font_find_shaper(param0 unsafe.Pointer, param1 uint32) {}

func Fn_pango_font_get_coverage(param0 unsafe.Pointer) {}

func Fn_pango_font_get_font_map() {}

func Fn_pango_font_get_glyph_extents(param0 uint32, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_pango_font_get_metrics(param0 unsafe.Pointer) {}

func Fn_pango_font_descriptions_free(param0 []unsafe.Pointer, param1 int) {}

func Fn_pango_font_face_describe() {}

func Fn_pango_font_face_get_face_name() {}

func Fn_pango_font_face_is_synthesized() {}

func Fn_pango_font_face_list_sizes(param0 []*int, param1 *int) {}

func Fn_pango_font_family_get_name() {}

func Fn_pango_font_family_is_monospace() {}

func Fn_pango_font_family_list_faces(param0 []*unsafe.Pointer, param1 *int) {}

func Fn_pango_font_map_changed() {}

func Fn_pango_font_map_create_context() {}

func Fn_pango_font_map_get_serial() {}

func Fn_pango_font_map_get_shape_engine_type() {}

func Fn_pango_font_map_list_families(param0 []*unsafe.Pointer, param1 *int) {}

func Fn_pango_font_map_load_font(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_font_map_load_fontset(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

// UNSUPPORTED : foreach : has callback

func Fn_pango_fontset_get_font(param0 uint) {}

func Fn_pango_fontset_get_metrics() {}

func Fn_pango_fontset_simple_new(param0 unsafe.Pointer) {}

func Fn_pango_fontset_simple_append(param0 unsafe.Pointer) {}

func Fn_pango_fontset_simple_size() {}

func Fn_pango_layout_new(param0 unsafe.Pointer) {}

func Fn_pango_layout_context_changed() {}

func Fn_pango_layout_copy() {}

func Fn_pango_layout_get_alignment() {}

func Fn_pango_layout_get_attributes() {}

func Fn_pango_layout_get_auto_dir() {}

func Fn_pango_layout_get_baseline() {}

func Fn_pango_layout_get_character_count() {}

func Fn_pango_layout_get_context() {}

func Fn_pango_layout_get_cursor_pos(param0 int, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_pango_layout_get_ellipsize() {}

func Fn_pango_layout_get_extents(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_layout_get_font_description() {}

func Fn_pango_layout_get_height() {}

func Fn_pango_layout_get_indent() {}

func Fn_pango_layout_get_iter() {}

func Fn_pango_layout_get_justify() {}

func Fn_pango_layout_get_line(param0 int) {}

func Fn_pango_layout_get_line_count() {}

func Fn_pango_layout_get_line_readonly(param0 int) {}

func Fn_pango_layout_get_lines() {}

func Fn_pango_layout_get_lines_readonly() {}

func Fn_pango_layout_get_log_attrs(param0 []unsafe.Pointer, param1 *int) {}

func Fn_pango_layout_get_log_attrs_readonly(param0 *int) {}

func Fn_pango_layout_get_pixel_extents(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_pango_layout_get_pixel_size(param0 *int, param1 *int) {}

func Fn_pango_layout_get_serial() {}

func Fn_pango_layout_get_single_paragraph_mode() {}

func Fn_pango_layout_get_size(param0 *int, param1 *int) {}

func Fn_pango_layout_get_spacing() {}

func Fn_pango_layout_get_tabs() {}

func Fn_pango_layout_get_text() {}

func Fn_pango_layout_get_unknown_glyphs_count() {}

func Fn_pango_layout_get_width() {}

func Fn_pango_layout_get_wrap() {}

func Fn_pango_layout_index_to_line_x(param0 int, param1 bool, param2 *int, param3 *int) {}

func Fn_pango_layout_index_to_pos(param0 int, param1 unsafe.Pointer) {}

func Fn_pango_layout_is_ellipsized() {}

func Fn_pango_layout_is_wrapped() {}

func Fn_pango_layout_move_cursor_visually(param0 bool, param1 int, param2 int, param3 int, param4 *int, param5 *int) {
}

func Fn_pango_layout_set_alignment(param0 int) {}

func Fn_pango_layout_set_attributes(param0 unsafe.Pointer) {}

func Fn_pango_layout_set_auto_dir(param0 bool) {}

func Fn_pango_layout_set_ellipsize(param0 int) {}

func Fn_pango_layout_set_font_description(param0 unsafe.Pointer) {}

func Fn_pango_layout_set_height(param0 int) {}

func Fn_pango_layout_set_indent(param0 int) {}

func Fn_pango_layout_set_justify(param0 bool) {}

func Fn_pango_layout_set_markup(param0 string, param1 int) {}

func Fn_pango_layout_set_markup_with_accel(param0 string, param1 int, param2 rune, param3 *rune) {}

func Fn_pango_layout_set_single_paragraph_mode(param0 bool) {}

func Fn_pango_layout_set_spacing(param0 int) {}

func Fn_pango_layout_set_tabs(param0 unsafe.Pointer) {}

func Fn_pango_layout_set_text(param0 string, param1 int) {}

func Fn_pango_layout_set_width(param0 int) {}

func Fn_pango_layout_set_wrap(param0 int) {}

func Fn_pango_layout_xy_to_index(param0 int, param1 int, param2 *int, param3 *int) {}

func Fn_pango_renderer_activate() {}

func Fn_pango_renderer_deactivate() {}

func Fn_pango_renderer_draw_error_underline(param0 int, param1 int, param2 int, param3 int) {}

func Fn_pango_renderer_draw_glyph(param0 unsafe.Pointer, param1 uint32, param2 float64, param3 float64) {
}

func Fn_pango_renderer_draw_glyph_item(param0 string, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_pango_renderer_draw_glyphs(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_pango_renderer_draw_layout(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_pango_renderer_draw_layout_line(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_pango_renderer_draw_rectangle(param0 int, param1 int, param2 int, param3 int, param4 int) {}

func Fn_pango_renderer_draw_trapezoid(param0 int, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64, param6 float64) {
}

func Fn_pango_renderer_get_alpha(param0 int) {}

func Fn_pango_renderer_get_color(param0 int) {}

func Fn_pango_renderer_get_layout() {}

func Fn_pango_renderer_get_layout_line() {}

func Fn_pango_renderer_get_matrix() {}

func Fn_pango_renderer_part_changed(param0 int) {}

func Fn_pango_renderer_set_alpha(param0 int, param1 uint16) {}

func Fn_pango_renderer_set_color(param0 int, param1 unsafe.Pointer) {}

func Fn_pango_renderer_set_matrix(param0 unsafe.Pointer) {}

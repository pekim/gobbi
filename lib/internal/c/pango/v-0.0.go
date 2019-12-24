// Code generated - DO NOT EDIT.
// +build !pango_1.6,!pango_1.8,!pango_1.16,!pango_1.20,!pango_1.22,!pango_1.24,!pango_1.36.7,!pango_1.38,!pango_1.42

package pango

import "unsafe"

// #include <pango/pango.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
}

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

func Fn_pango_attr_background_new(param0 uint16, param1 uint16, param2 uint16) {
	cValue0 := (C.guint16)(param0)
	cValue1 := (C.guint16)(param1)
	cValue2 := (C.guint16)(param2)

	C.pango_attr_background_new(cValue0, cValue1, cValue2)
}

func Fn_pango_attr_family_new(param0 string) {
	// has string param
}

func Fn_pango_attr_foreground_new(param0 uint16, param1 uint16, param2 uint16) {
	cValue0 := (C.guint16)(param0)
	cValue1 := (C.guint16)(param1)
	cValue2 := (C.guint16)(param2)

	C.pango_attr_foreground_new(cValue0, cValue1, cValue2)
}

func Fn_pango_attr_rise_new(param0 int) {
	cValue0 := (C.int)(param0)

	C.pango_attr_rise_new(cValue0)
}

func Fn_pango_attr_scale_new(param0 float64) {
	cValue0 := (C.double)(param0)

	C.pango_attr_scale_new(cValue0)
}

func Fn_pango_attr_stretch_new(param0 int) {
	cValue0 := (C.PangoStretch)(param0)

	C.pango_attr_stretch_new(cValue0)
}

func Fn_pango_attr_strikethrough_new(param0 bool) {
	cValue0 := toCBool(param0)

	C.pango_attr_strikethrough_new(cValue0)
}

func Fn_pango_attr_style_new(param0 int) {
	cValue0 := (C.PangoStyle)(param0)

	C.pango_attr_style_new(cValue0)
}

func Fn_pango_attr_type_register(param0 string) {
	// has string param
}

func Fn_pango_attr_underline_new(param0 int) {
	cValue0 := (C.PangoUnderline)(param0)

	C.pango_attr_underline_new(cValue0)
}

func Fn_pango_attr_variant_new(param0 int) {
	cValue0 := (C.PangoVariant)(param0)

	C.pango_attr_variant_new(cValue0)
}

func Fn_pango_attr_weight_new(param0 int) {
	cValue0 := (C.PangoWeight)(param0)

	C.pango_attr_weight_new(cValue0)
}

func Fn_pango_break(param0 string, param1 int, param2 unsafe.Pointer, param3 []LogAttr, param4 int) {
	// has array param
}

// UNSUPPORTED : config_key_get : blacklisted
// UNSUPPORTED : config_key_get_system : blacklisted
// UNSUPPORTED : default_break : blacklisted
// UNSUPPORTED : find_map : blacklisted
func Fn_pango_find_paragraph_boundary(param0 string, param1 int, param2 *int, param3 *int) {
	// has string param
}

func Fn_pango_font_description_from_string(param0 string) {
	// has string param
}

// UNSUPPORTED : get_lib_subdirectory : blacklisted
func Fn_pango_get_log_attrs(param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 []LogAttr, param5 int) {
	// has array param
}

func Fn_pango_get_mirror_char(param0 rune, param1 *rune) {
	cValue0 := (C.gunichar)(param0)
	cValue1 := (*C.gunichar)(unsafe.Pointer(param1))

	C.pango_get_mirror_char(cValue0, cValue1)
}

// UNSUPPORTED : get_sysconf_subdirectory : blacklisted
func Fn_pango_itemize(param0 unsafe.Pointer, param1 string, param2 int, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	// has string param
}

func Fn_pango_language_from_string(param0 string) {
	// has string param
}

func Fn_pango_lookup_aliases(param0 string, param1 *[]string, param2 *int) {
	// has array param
}

// UNSUPPORTED : module_register : blacklisted
func Fn_pango_parse_markup(param0 string, param1 int, param2 rune, param3 *unsafe.Pointer, param4 string, param5 *rune, error unsafe.Pointer) {
	// has string param
}

func Fn_pango_parse_stretch(param0 string, param1 *int, param2 bool) {
	// has string param
}

func Fn_pango_parse_style(param0 string, param1 *int, param2 bool) {
	// has string param
}

func Fn_pango_parse_variant(param0 string, param1 *int, param2 bool) {
	// has string param
}

func Fn_pango_parse_weight(param0 string, param1 *int, param2 bool) {
	// has string param
}

func Fn_pango_read_line(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.FILE)(unsafe.Pointer(param0))
	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	C.pango_read_line(cValue0, cValue1)
}

func Fn_pango_reorder_items(param0 unsafe.Pointer) {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.pango_reorder_items(cValue0)
}

func Fn_pango_scan_int(param0 string, param1 *int) {
	// has string param
}

func Fn_pango_scan_string(param0 string, param1 unsafe.Pointer) {
	// has string param
}

func Fn_pango_scan_word(param0 string, param1 unsafe.Pointer) {
	// has string param
}

func Fn_pango_shape(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
	// has string param
}

func Fn_pango_skip_space(param0 string) {
	// has string param
}

func Fn_pango_split_file_list(param0 string) {
	// has string param
}

func Fn_pango_trim_string(param0 string) {
	// has string param
}

func Fn_pango_unichar_direction(param0 rune) {
	cValue0 := (C.gunichar)(param0)

	C.pango_unichar_direction(cValue0)
}

func Fn_pango_context_new() {

	C.pango_context_new()
}

func Fn_pango_context_get_base_dir(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	C.pango_context_get_base_dir(cValueInstance)
}

func Fn_pango_context_get_font_description(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	C.pango_context_get_font_description(cValueInstance)
}

func Fn_pango_context_get_language(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	C.pango_context_get_language(cValueInstance)
}

func Fn_pango_context_get_metrics(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))
	cValue1 := (*C.PangoLanguage)(unsafe.Pointer(param1))

	C.pango_context_get_metrics(cValueInstance, cValue0, cValue1)
}

func Fn_pango_context_list_families(paramInstance unsafe.Pointer, param0 []*unsafe.Pointer, param1 *int) {
	// has array param
}

func Fn_pango_context_load_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.pango_context_load_font(cValueInstance, cValue0)
}

func Fn_pango_context_load_fontset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))
	cValue1 := (*C.PangoLanguage)(unsafe.Pointer(param1))

	C.pango_context_load_fontset(cValueInstance, cValue0, cValue1)
}

func Fn_pango_context_set_base_dir(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))
	cValue0 := (C.PangoDirection)(param0)

	C.pango_context_set_base_dir(cValueInstance, cValue0)
}

func Fn_pango_context_set_font_description(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.pango_context_set_font_description(cValueInstance, cValue0)
}

func Fn_pango_context_set_font_map(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoFontMap)(unsafe.Pointer(param0))

	C.pango_context_set_font_map(cValueInstance, cValue0)
}

func Fn_pango_context_set_language(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	C.pango_context_set_language(cValueInstance, cValue0)
}

func Fn_pango_font_describe(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	C.pango_font_describe(cValueInstance)
}

func Fn_pango_font_find_shaper(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint32) {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))
	cValue1 := (C.guint32)(param1)

	C.pango_font_find_shaper(cValueInstance, cValue0, cValue1)
}

func Fn_pango_font_get_coverage(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	C.pango_font_get_coverage(cValueInstance, cValue0)
}

func Fn_pango_font_get_glyph_extents(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))
	cValue0 := (C.PangoGlyph)(param0)
	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))
	cValue2 := (*C.PangoRectangle)(unsafe.Pointer(param2))

	C.pango_font_get_glyph_extents(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_font_get_metrics(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	C.pango_font_get_metrics(cValueInstance, cValue0)
}

func Fn_pango_font_descriptions_free(param0 []unsafe.Pointer, param1 int) {
	// has array param
}

func Fn_pango_font_face_describe(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFontFace)(unsafe.Pointer(paramInstance))

	C.pango_font_face_describe(cValueInstance)
}

func Fn_pango_font_face_get_face_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFontFace)(unsafe.Pointer(paramInstance))

	C.pango_font_face_get_face_name(cValueInstance)
}

func Fn_pango_font_family_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFontFamily)(unsafe.Pointer(paramInstance))

	C.pango_font_family_get_name(cValueInstance)
}

func Fn_pango_font_family_list_faces(paramInstance unsafe.Pointer, param0 []*unsafe.Pointer, param1 *int) {
	// has array param
}

// UNSUPPORTED : get_shape_engine_type : blacklisted
func Fn_pango_font_map_list_families(paramInstance unsafe.Pointer, param0 []*unsafe.Pointer, param1 *int) {
	// has array param
}

func Fn_pango_font_map_load_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))
	cValue1 := (*C.PangoFontDescription)(unsafe.Pointer(param1))

	C.pango_font_map_load_font(cValueInstance, cValue0, cValue1)
}

func Fn_pango_font_map_load_fontset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))
	cValue1 := (*C.PangoFontDescription)(unsafe.Pointer(param1))
	cValue2 := (*C.PangoLanguage)(unsafe.Pointer(param2))

	C.pango_font_map_load_fontset(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : foreach : has callback

func Fn_pango_fontset_get_font(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.PangoFontset)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.pango_fontset_get_font(cValueInstance, cValue0)
}

func Fn_pango_fontset_get_metrics(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFontset)(unsafe.Pointer(paramInstance))

	C.pango_fontset_get_metrics(cValueInstance)
}

// UNSUPPORTED : new : blacklisted
// UNSUPPORTED : append : blacklisted
// UNSUPPORTED : size : blacklisted
func Fn_pango_layout_new(param0 unsafe.Pointer) {
	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))

	C.pango_layout_new(cValue0)
}

func Fn_pango_layout_context_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_context_changed(cValueInstance)
}

func Fn_pango_layout_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_copy(cValueInstance)
}

func Fn_pango_layout_get_alignment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_alignment(cValueInstance)
}

func Fn_pango_layout_get_attributes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_attributes(cValueInstance)
}

func Fn_pango_layout_get_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_context(cValueInstance)
}

func Fn_pango_layout_get_cursor_pos(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))
	cValue2 := (*C.PangoRectangle)(unsafe.Pointer(param2))

	C.pango_layout_get_cursor_pos(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_layout_get_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))
	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_get_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_get_indent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_indent(cValueInstance)
}

func Fn_pango_layout_get_iter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_iter(cValueInstance)
}

func Fn_pango_layout_get_justify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_justify(cValueInstance)
}

func Fn_pango_layout_get_line(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.pango_layout_get_line(cValueInstance, cValue0)
}

func Fn_pango_layout_get_line_count(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_line_count(cValueInstance)
}

func Fn_pango_layout_get_lines(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_lines(cValueInstance)
}

func Fn_pango_layout_get_log_attrs(paramInstance unsafe.Pointer, param0 []unsafe.Pointer, param1 *int) {
	// has array param
}

func Fn_pango_layout_get_pixel_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))
	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_get_pixel_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_get_pixel_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.int)(unsafe.Pointer(param0))
	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_layout_get_pixel_size(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_get_single_paragraph_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_single_paragraph_mode(cValueInstance)
}

func Fn_pango_layout_get_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.int)(unsafe.Pointer(param0))
	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_layout_get_size(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_get_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_spacing(cValueInstance)
}

func Fn_pango_layout_get_tabs(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_tabs(cValueInstance)
}

func Fn_pango_layout_get_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_text(cValueInstance)
}

func Fn_pango_layout_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_width(cValueInstance)
}

func Fn_pango_layout_get_wrap(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_get_wrap(cValueInstance)
}

func Fn_pango_layout_index_to_line_x(paramInstance unsafe.Pointer, param0 int, param1 bool, param2 *int, param3 *int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := toCBool(param1)
	cValue2 := (*C.int)(unsafe.Pointer(param2))
	cValue3 := (*C.int)(unsafe.Pointer(param3))

	C.pango_layout_index_to_line_x(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_layout_index_to_pos(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_index_to_pos(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_move_cursor_visually(paramInstance unsafe.Pointer, param0 bool, param1 int, param2 int, param3 int, param4 *int, param5 *int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)
	cValue1 := (C.int)(param1)
	cValue2 := (C.int)(param2)
	cValue3 := (C.int)(param3)
	cValue4 := (*C.int)(unsafe.Pointer(param4))
	cValue5 := (*C.int)(unsafe.Pointer(param5))

	C.pango_layout_move_cursor_visually(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_pango_layout_set_alignment(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.PangoAlignment)(param0)

	C.pango_layout_set_alignment(cValueInstance, cValue0)
}

func Fn_pango_layout_set_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoAttrList)(unsafe.Pointer(param0))

	C.pango_layout_set_attributes(cValueInstance, cValue0)
}

func Fn_pango_layout_set_font_description(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.pango_layout_set_font_description(cValueInstance, cValue0)
}

func Fn_pango_layout_set_indent(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.pango_layout_set_indent(cValueInstance, cValue0)
}

func Fn_pango_layout_set_justify(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.pango_layout_set_justify(cValueInstance, cValue0)
}

func Fn_pango_layout_set_markup(paramInstance unsafe.Pointer, param0 string, param1 int) {
	// has string param
}

func Fn_pango_layout_set_markup_with_accel(paramInstance unsafe.Pointer, param0 string, param1 int, param2 rune, param3 *rune) {
	// has string param
}

func Fn_pango_layout_set_single_paragraph_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.pango_layout_set_single_paragraph_mode(cValueInstance, cValue0)
}

func Fn_pango_layout_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.pango_layout_set_spacing(cValueInstance, cValue0)
}

func Fn_pango_layout_set_tabs(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.PangoTabArray)(unsafe.Pointer(param0))

	C.pango_layout_set_tabs(cValueInstance, cValue0)
}

func Fn_pango_layout_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {
	// has string param
}

func Fn_pango_layout_set_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.pango_layout_set_width(cValueInstance, cValue0)
}

func Fn_pango_layout_set_wrap(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.PangoWrapMode)(param0)

	C.pango_layout_set_wrap(cValueInstance, cValue0)
}

func Fn_pango_layout_xy_to_index(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)
	cValue1 := (C.int)(param1)
	cValue2 := (*C.int)(unsafe.Pointer(param2))
	cValue3 := (*C.int)(unsafe.Pointer(param3))

	C.pango_layout_xy_to_index(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// Code generated - DO NOT EDIT.
// +build pango_1.42

package pango

import "unsafe"

// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

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

func Fn_pango_attr_background_alpha_new(param0 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	ret := C.pango_attr_background_alpha_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_background_new(param0 uint16, param1 uint16, param2 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	ret := C.pango_attr_background_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_fallback_new(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.pango_attr_fallback_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_family_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_attr_family_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_foreground_alpha_new(param0 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	ret := C.pango_attr_foreground_alpha_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_foreground_new(param0 uint16, param1 uint16, param2 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	ret := C.pango_attr_foreground_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_gravity_hint_new(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoGravityHint)(param0)

	ret := C.pango_attr_gravity_hint_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_gravity_new(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoGravity)(param0)

	ret := C.pango_attr_gravity_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_letter_spacing_new(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.pango_attr_letter_spacing_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_rise_new(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.pango_attr_rise_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_scale_new(param0 float64) unsafe.Pointer {
	cValue0 := (C.double)(param0)

	ret := C.pango_attr_scale_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_stretch_new(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoStretch)(param0)

	ret := C.pango_attr_stretch_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_strikethrough_color_new(param0 uint16, param1 uint16, param2 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	ret := C.pango_attr_strikethrough_color_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_strikethrough_new(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.pango_attr_strikethrough_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_style_new(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoStyle)(param0)

	ret := C.pango_attr_style_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_type_get_name(param0 int) string {
	cValue0 := (C.PangoAttrType)(param0)

	ret := C.pango_attr_type_get_name(cValue0)

	return C.GoString(ret)
}

func Fn_pango_attr_type_register(param0 string) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_attr_type_register(cValue0)

	return (int)(ret)
}

func Fn_pango_attr_underline_color_new(param0 uint16, param1 uint16, param2 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	ret := C.pango_attr_underline_color_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_underline_new(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoUnderline)(param0)

	ret := C.pango_attr_underline_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_variant_new(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoVariant)(param0)

	ret := C.pango_attr_variant_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_weight_new(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoWeight)(param0)

	ret := C.pango_attr_weight_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_bidi_type_for_unichar(param0 rune) int {
	cValue0 := (C.gunichar)(param0)

	ret := C.pango_bidi_type_for_unichar(cValue0)

	return (int)(ret)
}

// UNSUPPORTED : pango_break : has non-string array param attrs

// UNSUPPORTED : pango_config_key_get : blacklisted
// UNSUPPORTED : pango_config_key_get_system : blacklisted
// UNSUPPORTED : pango_default_break : blacklisted
func Fn_pango_extents_to_pixels(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_extents_to_pixels(cValue0, cValue1)
}

func Fn_pango_find_base_dir(param0 string, param1 int) int {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.pango_find_base_dir(cValue0, cValue1)

	return (int)(ret)
}

// UNSUPPORTED : pango_find_map : blacklisted
func Fn_pango_find_paragraph_boundary(param0 string, param1 int, param2 *int, param3 *int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.pango_find_paragraph_boundary(cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_font_description_from_string(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_font_description_from_string(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted
// UNSUPPORTED : pango_get_log_attrs : has non-string array param log_attrs

func Fn_pango_get_mirror_char(param0 rune, param1 *rune) bool {
	cValue0 := (C.gunichar)(param0)

	cValue1 := (*C.gunichar)(unsafe.Pointer(param1))

	ret := C.pango_get_mirror_char(cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted
func Fn_pango_gravity_get_for_matrix(param0 unsafe.Pointer) int {
	cValue0 := (*C.PangoMatrix)(unsafe.Pointer(param0))

	ret := C.pango_gravity_get_for_matrix(cValue0)

	return (int)(ret)
}

func Fn_pango_gravity_get_for_script(param0 int, param1 int, param2 int) int {
	cValue0 := (C.PangoScript)(param0)

	cValue1 := (C.PangoGravity)(param1)

	cValue2 := (C.PangoGravityHint)(param2)

	ret := C.pango_gravity_get_for_script(cValue0, cValue1, cValue2)

	return (int)(ret)
}

func Fn_pango_gravity_get_for_script_and_width(param0 int, param1 bool, param2 int, param3 int) int {
	cValue0 := (C.PangoScript)(param0)

	cValue1 := toCBool(param1)

	cValue2 := (C.PangoGravity)(param2)

	cValue3 := (C.PangoGravityHint)(param3)

	ret := C.pango_gravity_get_for_script_and_width(cValue0, cValue1, cValue2, cValue3)

	return (int)(ret)
}

func Fn_pango_gravity_to_rotation(param0 int) float64 {
	cValue0 := (C.PangoGravity)(param0)

	ret := C.pango_gravity_to_rotation(cValue0)

	return (float64)(ret)
}

func Fn_pango_is_zero_width(param0 rune) bool {
	cValue0 := (C.gunichar)(param0)

	ret := C.pango_is_zero_width(cValue0)

	return toGoBool(ret)
}

func Fn_pango_itemize(param0 unsafe.Pointer, param1 string, param2 int, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (*C.PangoAttrList)(unsafe.Pointer(param4))

	cValue5 := (*C.PangoAttrIterator)(unsafe.Pointer(param5))

	ret := C.pango_itemize(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_pango_itemize_with_base_dir(param0 unsafe.Pointer, param1 int, param2 string, param3 int, param4 int, param5 unsafe.Pointer, param6 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))

	cValue1 := (C.PangoDirection)(param1)

	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	cValue5 := (*C.PangoAttrList)(unsafe.Pointer(param5))

	cValue6 := (*C.PangoAttrIterator)(unsafe.Pointer(param6))

	ret := C.pango_itemize_with_base_dir(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_pango_language_from_string(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_language_from_string(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_language_get_default() unsafe.Pointer {
	ret := C.pango_language_get_default()

	return unsafe.Pointer(ret)
}

func Fn_pango_log2vis_get_embedding_levels(param0 string, param1 int, param2 *int) *uint8 {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (*C.PangoDirection)(unsafe.Pointer(param2))

	ret := C.pango_log2vis_get_embedding_levels(cValue0, cValue1, cValue2)

	return (*uint8)(ret)
}

// UNSUPPORTED : pango_lookup_aliases : blacklisted
func Fn_pango_markup_parser_finish(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *string, param3 *rune, error unsafe.Pointer) bool {
	cValue0 := (*C.GMarkupParseContext)(unsafe.Pointer(param0))

	cValue1 := (**C.PangoAttrList)(unsafe.Pointer(param1))

	var cValue2String *C.gchar
	cValue2 := &cValue2String

	cValue3 := (*C.gunichar)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.pango_markup_parser_finish(cValue0, cValue1, cValue2, cValue3, cError)

	param2String := C.GoString(cValue2String)
	*param2 = param2String

	return toGoBool(ret)
}

func Fn_pango_markup_parser_new(param0 rune) unsafe.Pointer {
	cValue0 := (C.gunichar)(param0)

	ret := C.pango_markup_parser_new(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_module_register : blacklisted
func Fn_pango_parse_enum(param0 uint64, param1 string, param2 *int, param3 bool, param4 *string) bool {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.int)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	var cValue4String *C.gchar
	cValue4 := &cValue4String

	ret := C.pango_parse_enum(cValue0, cValue1, cValue2, cValue3, cValue4)

	param4String := C.GoString(cValue4String)
	*param4 = param4String

	return toGoBool(ret)
}

func Fn_pango_parse_markup(param0 string, param1 int, param2 rune, param3 *unsafe.Pointer, param4 *string, param5 *rune, error unsafe.Pointer) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.gunichar)(param2)

	cValue3 := (**C.PangoAttrList)(unsafe.Pointer(param3))

	var cValue4String *C.gchar
	cValue4 := &cValue4String

	cValue5 := (*C.gunichar)(unsafe.Pointer(param5))

	cError := (**C.GError)(error)

	ret := C.pango_parse_markup(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cError)

	param4String := C.GoString(cValue4String)
	*param4 = param4String

	return toGoBool(ret)
}

func Fn_pango_parse_stretch(param0 string, param1 *int, param2 bool) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoStretch)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.pango_parse_stretch(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_pango_parse_style(param0 string, param1 *int, param2 bool) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoStyle)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.pango_parse_style(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_pango_parse_variant(param0 string, param1 *int, param2 bool) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoVariant)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.pango_parse_variant(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_pango_parse_weight(param0 string, param1 *int, param2 bool) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoWeight)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.pango_parse_weight(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_pango_quantize_line_geometry(param0 *int, param1 *int) {
	cValue0 := (*C.int)(unsafe.Pointer(param0))

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_quantize_line_geometry(cValue0, cValue1)
}

func Fn_pango_read_line(param0 unsafe.Pointer, param1 unsafe.Pointer) int {
	cValue0 := (*C.FILE)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	ret := C.pango_read_line(cValue0, cValue1)

	return (int)(ret)
}

func Fn_pango_reorder_items(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	ret := C.pango_reorder_items(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_scan_int(param0 *string, param1 *int) bool {
	var cValue0String *C.gchar
	cValue0String = (*C.gchar)(C.CString(*param0))
	cValue0 := &cValue0String

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	ret := C.pango_scan_int(cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

func Fn_pango_scan_string(param0 *string, param1 unsafe.Pointer) bool {
	var cValue0String *C.gchar
	cValue0String = (*C.gchar)(C.CString(*param0))
	cValue0 := &cValue0String

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	ret := C.pango_scan_string(cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

func Fn_pango_scan_word(param0 *string, param1 unsafe.Pointer) bool {
	var cValue0String *C.gchar
	cValue0String = (*C.gchar)(C.CString(*param0))
	cValue0 := &cValue0String

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	ret := C.pango_scan_word(cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

func Fn_pango_script_for_unichar(param0 rune) int {
	cValue0 := (C.gunichar)(param0)

	ret := C.pango_script_for_unichar(cValue0)

	return (int)(ret)
}

func Fn_pango_script_get_sample_language(param0 int) unsafe.Pointer {
	cValue0 := (C.PangoScript)(param0)

	ret := C.pango_script_get_sample_language(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_shape(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.PangoAnalysis)(unsafe.Pointer(param2))

	cValue3 := (*C.PangoGlyphString)(unsafe.Pointer(param3))

	C.pango_shape(cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_shape_full(param0 string, param1 int, param2 string, param3 int, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gint)(param3)

	cValue4 := (*C.PangoAnalysis)(unsafe.Pointer(param4))

	cValue5 := (*C.PangoGlyphString)(unsafe.Pointer(param5))

	C.pango_shape_full(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_pango_skip_space(param0 *string) bool {
	var cValue0String *C.gchar
	cValue0String = (*C.gchar)(C.CString(*param0))
	cValue0 := &cValue0String

	ret := C.pango_skip_space(cValue0)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

// UNSUPPORTED : pango_split_file_list : has array return

func Fn_pango_trim_string(param0 string) string {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_trim_string(cValue0)

	return C.GoString(ret)
}

func Fn_pango_unichar_direction(param0 rune) int {
	cValue0 := (C.gunichar)(param0)

	ret := C.pango_unichar_direction(cValue0)

	return (int)(ret)
}

func Fn_pango_units_from_double(param0 float64) int {
	cValue0 := (C.double)(param0)

	ret := C.pango_units_from_double(cValue0)

	return (int)(ret)
}

func Fn_pango_units_to_double(param0 int) float64 {
	cValue0 := (C.int)(param0)

	ret := C.pango_units_to_double(cValue0)

	return (float64)(ret)
}

func Fn_pango_version() int {
	ret := C.pango_version()

	return (int)(ret)
}

func Fn_pango_version_check(param0 int, param1 int, param2 int) string {
	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	ret := C.pango_version_check(cValue0, cValue1, cValue2)

	return C.GoString(ret)
}

func Fn_pango_version_string() string {
	ret := C.pango_version_string()

	return C.GoString(ret)
}

func Fn_pango_context_new() unsafe.Pointer {
	ret := C.pango_context_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_context_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	C.pango_context_changed(cValueInstance)
}

func Fn_pango_context_get_base_dir(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_base_dir(cValueInstance)

	return (int)(ret)
}

func Fn_pango_context_get_base_gravity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_base_gravity(cValueInstance)

	return (int)(ret)
}

func Fn_pango_context_get_font_description(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_font_description(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_font_map(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_font_map(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_gravity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_gravity(cValueInstance)

	return (int)(ret)
}

func Fn_pango_context_get_gravity_hint(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_gravity_hint(cValueInstance)

	return (int)(ret)
}

func Fn_pango_context_get_language(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_language(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_matrix(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_matrix(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_metrics(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoLanguage)(unsafe.Pointer(param1))

	ret := C.pango_context_get_metrics(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_serial(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_serial(cValueInstance)

	return (uint)(ret)
}

// UNSUPPORTED : pango_context_list_families : has non-string array param families

func Fn_pango_context_load_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	ret := C.pango_context_load_font(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_load_fontset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoLanguage)(unsafe.Pointer(param1))

	ret := C.pango_context_load_fontset(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_set_base_dir(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoDirection)(param0)

	C.pango_context_set_base_dir(cValueInstance, cValue0)
}

func Fn_pango_context_set_base_gravity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoGravity)(param0)

	C.pango_context_set_base_gravity(cValueInstance, cValue0)
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

func Fn_pango_context_set_gravity_hint(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoGravityHint)(param0)

	C.pango_context_set_gravity_hint(cValueInstance, cValue0)
}

func Fn_pango_context_set_language(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	C.pango_context_set_language(cValueInstance, cValue0)
}

func Fn_pango_context_set_matrix(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoMatrix)(unsafe.Pointer(param0))

	C.pango_context_set_matrix(cValueInstance, cValue0)
}

func Fn_pango_font_describe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_describe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_describe_with_absolute_size(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_describe_with_absolute_size(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_find_shaper(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint32) unsafe.Pointer {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	cValue1 := (C.guint32)(param1)

	ret := C.pango_font_find_shaper(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_get_coverage(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	ret := C.pango_font_get_coverage(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_get_font_map(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_get_font_map(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_get_glyph_extents(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoGlyph)(param0)

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.PangoRectangle)(unsafe.Pointer(param2))

	C.pango_font_get_glyph_extents(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_font_get_metrics(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFont)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	ret := C.pango_font_get_metrics(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_font_descriptions_free : has non-string array param descs

func Fn_pango_font_face_describe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontFace)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_face_describe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_face_get_face_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoFontFace)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_face_get_face_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_font_face_is_synthesized(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoFontFace)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_face_is_synthesized(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_font_face_list_sizes : has non-string array param sizes

func Fn_pango_font_family_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoFontFamily)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_family_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_font_family_is_monospace(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoFontFamily)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_family_is_monospace(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_font_family_list_faces : has non-string array param faces

func Fn_pango_font_map_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	C.pango_font_map_changed(cValueInstance)
}

func Fn_pango_font_map_create_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_map_create_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_map_get_serial(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_map_get_serial(cValueInstance)

	return (uint)(ret)
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted
// UNSUPPORTED : pango_font_map_list_families : has non-string array param families

func Fn_pango_font_map_load_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoFontDescription)(unsafe.Pointer(param1))

	ret := C.pango_font_map_load_font(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_map_load_fontset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoFontDescription)(unsafe.Pointer(param1))

	cValue2 := (*C.PangoLanguage)(unsafe.Pointer(param2))

	ret := C.pango_font_map_load_fontset(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_fontset_foreach : has callback

func Fn_pango_fontset_get_font(paramInstance unsafe.Pointer, param0 uint) unsafe.Pointer {
	cValueInstance := (*C.PangoFontset)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.pango_fontset_get_font(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_fontset_get_metrics(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontset)(unsafe.Pointer(paramInstance))

	ret := C.pango_fontset_get_metrics(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_fontset_simple_new : blacklisted
// UNSUPPORTED : pango_fontset_simple_append : blacklisted
// UNSUPPORTED : pango_fontset_simple_size : blacklisted
func Fn_pango_layout_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.PangoContext)(unsafe.Pointer(param0))

	ret := C.pango_layout_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_context_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	C.pango_layout_context_changed(cValueInstance)
}

func Fn_pango_layout_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_alignment(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_alignment(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_attributes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_attributes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_auto_dir(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_auto_dir(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_get_baseline(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_baseline(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_character_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_character_count(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_cursor_pos(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.PangoRectangle)(unsafe.Pointer(param2))

	C.pango_layout_get_cursor_pos(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_layout_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_get_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_get_font_description(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_font_description(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_height(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_indent(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_indent(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_iter(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_iter(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_justify(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_justify(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_get_line(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	ret := C.pango_layout_get_line(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_line_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_line_count(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_line_readonly(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	ret := C.pango_layout_get_line_readonly(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_lines(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_lines(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_lines_readonly(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_lines_readonly(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_layout_get_log_attrs : has non-string array param attrs

// UNSUPPORTED : pango_layout_get_log_attrs_readonly : has array return

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

func Fn_pango_layout_get_serial(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_serial(cValueInstance)

	return (uint)(ret)
}

func Fn_pango_layout_get_single_paragraph_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_single_paragraph_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_get_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.int)(unsafe.Pointer(param0))

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_layout_get_size(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_tabs(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_tabs(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_layout_get_unknown_glyphs_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_unknown_glyphs_count(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_width(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_get_wrap(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_wrap(cValueInstance)

	return (int)(ret)
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

func Fn_pango_layout_is_ellipsized(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_is_ellipsized(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_is_wrapped(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_is_wrapped(cValueInstance)

	return toGoBool(ret)
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

func Fn_pango_layout_set_auto_dir(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.pango_layout_set_auto_dir(cValueInstance, cValue0)
}

func Fn_pango_layout_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.pango_layout_set_ellipsize(cValueInstance, cValue0)
}

func Fn_pango_layout_set_font_description(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.pango_layout_set_font_description(cValueInstance, cValue0)
}

func Fn_pango_layout_set_height(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	C.pango_layout_set_height(cValueInstance, cValue0)
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
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	C.pango_layout_set_markup(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_set_markup_with_accel(paramInstance unsafe.Pointer, param0 string, param1 int, param2 rune, param3 *rune) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.gunichar)(param2)

	cValue3 := (*C.gunichar)(unsafe.Pointer(param3))

	C.pango_layout_set_markup_with_accel(cValueInstance, cValue0, cValue1, cValue2, cValue3)
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
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	C.pango_layout_set_text(cValueInstance, cValue0, cValue1)
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

func Fn_pango_layout_xy_to_index(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) bool {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (*C.int)(unsafe.Pointer(param2))

	cValue3 := (*C.int)(unsafe.Pointer(param3))

	ret := C.pango_layout_xy_to_index(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_pango_renderer_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	C.pango_renderer_activate(cValueInstance)
}

func Fn_pango_renderer_deactivate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	C.pango_renderer_deactivate(cValueInstance)
}

func Fn_pango_renderer_draw_error_underline(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	C.pango_renderer_draw_error_underline(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_renderer_draw_glyph(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint32, param2 float64, param3 float64) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFont)(unsafe.Pointer(param0))

	cValue1 := (C.PangoGlyph)(param1)

	cValue2 := (C.double)(param2)

	cValue3 := (C.double)(param3)

	C.pango_renderer_draw_glyph(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_renderer_draw_glyph_item(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 int) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoGlyphItem)(unsafe.Pointer(param1))

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	C.pango_renderer_draw_glyph_item(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_renderer_draw_glyphs(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFont)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoGlyphString)(unsafe.Pointer(param1))

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	C.pango_renderer_draw_glyphs(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_renderer_draw_layout(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoLayout)(unsafe.Pointer(param0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	C.pango_renderer_draw_layout(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_renderer_draw_layout_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoLayoutLine)(unsafe.Pointer(param0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	C.pango_renderer_draw_layout_line(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_renderer_draw_rectangle(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 int) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (C.int)(param3)

	cValue4 := (C.int)(param4)

	C.pango_renderer_draw_rectangle(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_pango_renderer_draw_trapezoid(paramInstance unsafe.Pointer, param0 int, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64, param6 float64) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	cValue1 := (C.double)(param1)

	cValue2 := (C.double)(param2)

	cValue3 := (C.double)(param3)

	cValue4 := (C.double)(param4)

	cValue5 := (C.double)(param5)

	cValue6 := (C.double)(param6)

	C.pango_renderer_draw_trapezoid(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_pango_renderer_get_alpha(paramInstance unsafe.Pointer, param0 int) uint16 {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	ret := C.pango_renderer_get_alpha(cValueInstance, cValue0)

	return (uint16)(ret)
}

func Fn_pango_renderer_get_color(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	ret := C.pango_renderer_get_color(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_renderer_get_layout(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	ret := C.pango_renderer_get_layout(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_renderer_get_layout_line(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	ret := C.pango_renderer_get_layout_line(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_renderer_get_matrix(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	ret := C.pango_renderer_get_matrix(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_renderer_part_changed(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	C.pango_renderer_part_changed(cValueInstance, cValue0)
}

func Fn_pango_renderer_set_alpha(paramInstance unsafe.Pointer, param0 int, param1 uint16) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	cValue1 := (C.guint16)(param1)

	C.pango_renderer_set_alpha(cValueInstance, cValue0, cValue1)
}

func Fn_pango_renderer_set_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	cValue1 := (*C.PangoColor)(unsafe.Pointer(param1))

	C.pango_renderer_set_color(cValueInstance, cValue0, cValue1)
}

func Fn_pango_renderer_set_matrix(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoMatrix)(unsafe.Pointer(param0))

	C.pango_renderer_set_matrix(cValueInstance, cValue0)
}

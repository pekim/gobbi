// Code generated - DO NOT EDIT.
// +build pango_1.14

package pango

import "unsafe"

// #include <pango/pango.h>
// #include <pango/pango-font.h>
// #include <pango/pango-modules.h>
// #include <stdlib.h>
/*

static PangoTabArray* c_pango_tab_array_new_with_positions(gint size, gboolean positions_in_pixels, PangoTabAlign first_alignment, gint first_position) {
    return pango_tab_array_new_with_positions(size, positions_in_pixels, first_alignment, first_position, NULL);
}
*/
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

func Fn_pango_attr_font_desc_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	ret := C.pango_attr_font_desc_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_iterator_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoAttrIterator)(unsafe.Pointer(paramInstance))

	ret := C.pango_attr_iterator_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_iterator_destroy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoAttrIterator)(unsafe.Pointer(paramInstance))

	C.pango_attr_iterator_destroy(cValueInstance)
}

func Fn_pango_attr_iterator_get(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.PangoAttrIterator)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoAttrType)(param0)

	ret := C.pango_attr_iterator_get(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_iterator_get_attrs(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoAttrIterator)(unsafe.Pointer(paramInstance))

	ret := C.pango_attr_iterator_get_attrs(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_iterator_get_font : parameter 'language' is non array with indirect count > 1

func Fn_pango_attr_iterator_next(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoAttrIterator)(unsafe.Pointer(paramInstance))

	ret := C.pango_attr_iterator_next(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_attr_iterator_range(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.PangoAttrIterator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.pango_attr_iterator_range(cValueInstance, cValue0, cValue1)
}

func Fn_pango_attr_language_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.PangoLanguage)(unsafe.Pointer(param0))

	ret := C.pango_attr_language_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_list_new() unsafe.Pointer {
	ret := C.pango_attr_list_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_list_change(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoAttribute)(unsafe.Pointer(param0))

	C.pango_attr_list_change(cValueInstance, cValue0)
}

func Fn_pango_attr_list_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	ret := C.pango_attr_list_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

func Fn_pango_attr_list_get_iterator(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	ret := C.pango_attr_list_get_iterator(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_list_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoAttribute)(unsafe.Pointer(param0))

	C.pango_attr_list_insert(cValueInstance, cValue0)
}

func Fn_pango_attr_list_insert_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoAttribute)(unsafe.Pointer(param0))

	C.pango_attr_list_insert_before(cValueInstance, cValue0)
}

func Fn_pango_attr_list_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	ret := C.pango_attr_list_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_list_splice(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoAttrList)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.pango_attr_list_splice(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_attr_list_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoAttrList)(unsafe.Pointer(paramInstance))

	C.pango_attr_list_unref(cValueInstance)
}

func Fn_pango_attr_shape_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	ret := C.pango_attr_shape_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

func Fn_pango_attr_size_new(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.pango_attr_size_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_size_new_absolute(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.pango_attr_size_new_absolute(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attribute_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoAttribute)(unsafe.Pointer(paramInstance))

	ret := C.pango_attribute_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_attribute_destroy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoAttribute)(unsafe.Pointer(paramInstance))

	C.pango_attribute_destroy(cValueInstance)
}

func Fn_pango_attribute_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.PangoAttribute)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoAttribute)(unsafe.Pointer(param0))

	ret := C.pango_attribute_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_pango_color_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoColor)(unsafe.Pointer(paramInstance))

	ret := C.pango_color_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_color_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoColor)(unsafe.Pointer(paramInstance))

	C.pango_color_free(cValueInstance)
}

func Fn_pango_color_parse(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.PangoColor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_color_parse(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_pango_coverage_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoCoverage)(unsafe.Pointer(paramInstance))

	ret := C.pango_coverage_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_coverage_get(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.PangoCoverage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	ret := C.pango_coverage_get(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_pango_coverage_max(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoCoverage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoCoverage)(unsafe.Pointer(param0))

	C.pango_coverage_max(cValueInstance, cValue0)
}

func Fn_pango_coverage_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoCoverage)(unsafe.Pointer(paramInstance))

	ret := C.pango_coverage_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_coverage_set(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.PangoCoverage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.PangoCoverageLevel)(param1)

	C.pango_coverage_set(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

func Fn_pango_coverage_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoCoverage)(unsafe.Pointer(paramInstance))

	C.pango_coverage_unref(cValueInstance)
}

func Fn_pango_coverage_from_bytes(param0 []uint8, param1 int) unsafe.Pointer {
	cValue0 := (*C.guchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.int)(param1)

	ret := C.pango_coverage_from_bytes(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_coverage_new() unsafe.Pointer {
	ret := C.pango_coverage_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_font_description_new() unsafe.Pointer {
	ret := C.pango_font_description_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_font_description_better_match(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoFontDescription)(unsafe.Pointer(param1))

	ret := C.pango_font_description_better_match(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_pango_font_description_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_description_copy_static(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_copy_static(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_description_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	ret := C.pango_font_description_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_pango_font_description_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	C.pango_font_description_free(cValueInstance)
}

func Fn_pango_font_description_get_family(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_family(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_font_description_get_set_fields(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_set_fields(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_description_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_description_get_size_is_absolute(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_size_is_absolute(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_font_description_get_stretch(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_stretch(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_description_get_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_style(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_description_get_variant(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_variant(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_description_get_weight(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_weight(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_description_hash(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_hash(cValueInstance)

	return (uint)(ret)
}

func Fn_pango_font_description_merge(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.pango_font_description_merge(cValueInstance, cValue0, cValue1)
}

func Fn_pango_font_description_merge_static(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.pango_font_description_merge_static(cValueInstance, cValue0, cValue1)
}

func Fn_pango_font_description_set_absolute_size(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.double)(param0)

	C.pango_font_description_set_absolute_size(cValueInstance, cValue0)
}

func Fn_pango_font_description_set_family(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.pango_font_description_set_family(cValueInstance, cValue0)
}

func Fn_pango_font_description_set_family_static(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.pango_font_description_set_family_static(cValueInstance, cValue0)
}

func Fn_pango_font_description_set_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.pango_font_description_set_size(cValueInstance, cValue0)
}

func Fn_pango_font_description_set_stretch(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoStretch)(param0)

	C.pango_font_description_set_stretch(cValueInstance, cValue0)
}

func Fn_pango_font_description_set_style(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoStyle)(param0)

	C.pango_font_description_set_style(cValueInstance, cValue0)
}

func Fn_pango_font_description_set_variant(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoVariant)(param0)

	C.pango_font_description_set_variant(cValueInstance, cValue0)
}

func Fn_pango_font_description_set_weight(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoWeight)(param0)

	C.pango_font_description_set_weight(cValueInstance, cValue0)
}

func Fn_pango_font_description_to_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_to_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_font_description_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_font_description_unset_fields(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoFontMask)(param0)

	C.pango_font_description_unset_fields(cValueInstance, cValue0)
}

func Fn_pango_font_description_from_string(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_font_description_from_string(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_font_metrics_new : blacklisted

func Fn_pango_font_metrics_get_approximate_char_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_approximate_char_width(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_approximate_digit_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_approximate_digit_width(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_ascent(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_ascent(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_descent(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_descent(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_strikethrough_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_strikethrough_position(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_strikethrough_thickness(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_strikethrough_thickness(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_underline_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_underline_position(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_underline_thickness(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_get_underline_thickness(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_metrics_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_metrics_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_metrics_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoFontMetrics)(unsafe.Pointer(paramInstance))

	C.pango_font_metrics_unref(cValueInstance)
}

func Fn_pango_glyph_item_apply_attrs(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoGlyphItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoAttrList)(unsafe.Pointer(param1))

	ret := C.pango_glyph_item_apply_attrs(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_item_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoGlyphItem)(unsafe.Pointer(paramInstance))

	C.pango_glyph_item_free(cValueInstance)
}

func Fn_pango_glyph_item_letter_space(paramInstance unsafe.Pointer, param0 string, param1 []LogAttr, param2 int) {
	cValueInstance := (*C.PangoGlyphItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoLogAttr)(unsafe.Pointer(&param1[0]))

	cValue2 := (C.int)(param2)

	C.pango_glyph_item_letter_space(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_glyph_item_split(paramInstance unsafe.Pointer, param0 string, param1 int) unsafe.Pointer {
	cValueInstance := (*C.PangoGlyphItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	ret := C.pango_glyph_item_split(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_string_new() unsafe.Pointer {
	ret := C.pango_glyph_string_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_string_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	ret := C.pango_glyph_string_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_string_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFont)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.PangoRectangle)(unsafe.Pointer(param2))

	C.pango_glyph_string_extents(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_glyph_string_extents_range(paramInstance unsafe.Pointer, param0 int, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer) {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	cValue2 := (*C.PangoFont)(unsafe.Pointer(param2))

	cValue3 := (*C.PangoRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.PangoRectangle)(unsafe.Pointer(param4))

	C.pango_glyph_string_extents_range(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_pango_glyph_string_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	C.pango_glyph_string_free(cValueInstance)
}

func Fn_pango_glyph_string_get_logical_widths(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, param3 []int) {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (*C.int)(unsafe.Pointer(&param3[0]))

	C.pango_glyph_string_get_logical_widths(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_pango_glyph_string_get_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	ret := C.pango_glyph_string_get_width(cValueInstance)

	return (int)(ret)
}

func Fn_pango_glyph_string_index_to_x(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, param3 int, param4 bool, param5 *int) {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (*C.PangoAnalysis)(unsafe.Pointer(param2))

	cValue3 := (C.int)(param3)

	cValue4 := toCBool(param4)

	cValue5 := (*C.int)(unsafe.Pointer(param5))

	C.pango_glyph_string_index_to_x(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_pango_glyph_string_set_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.pango_glyph_string_set_size(cValueInstance, cValue0)
}

func Fn_pango_glyph_string_x_to_index(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, param3 int, param4 *int, param5 *int) {
	cValueInstance := (*C.PangoGlyphString)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (*C.PangoAnalysis)(unsafe.Pointer(param2))

	cValue3 := (C.int)(param3)

	cValue4 := (*C.int)(unsafe.Pointer(param4))

	cValue5 := (*C.int)(unsafe.Pointer(param5))

	C.pango_glyph_string_x_to_index(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_pango_item_new() unsafe.Pointer {
	ret := C.pango_item_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_item_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoItem)(unsafe.Pointer(paramInstance))

	ret := C.pango_item_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_item_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoItem)(unsafe.Pointer(paramInstance))

	C.pango_item_free(cValueInstance)
}

func Fn_pango_item_split(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.PangoItem)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	ret := C.pango_item_split(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_language_get_sample_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoLanguage)(unsafe.Pointer(paramInstance))

	ret := C.pango_language_get_sample_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_language_includes_script(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.PangoLanguage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoScript)(param0)

	ret := C.pango_language_includes_script(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_pango_language_matches(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.PangoLanguage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_language_matches(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_pango_language_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoLanguage)(unsafe.Pointer(paramInstance))

	ret := C.pango_language_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_pango_language_from_string(param0 *string) unsafe.Pointer {
	var cValue0Value (*C.char)
	if param0 != nil {
		cValue0Value = (*C.char)(C.CString(*param0))
		defer C.free(unsafe.Pointer(cValue0Value))
	}
	cValue0 := cValue0Value

	ret := C.pango_language_from_string(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_iter_at_last_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_at_last_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	C.pango_layout_iter_free(cValueInstance)
}

func Fn_pango_layout_iter_get_baseline(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_get_baseline(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_iter_get_char_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	C.pango_layout_iter_get_char_extents(cValueInstance, cValue0)
}

func Fn_pango_layout_iter_get_cluster_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_iter_get_cluster_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_iter_get_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_get_index(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_iter_get_layout_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_iter_get_layout_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_iter_get_line(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_get_line(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_iter_get_line_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_iter_get_line_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_iter_get_line_yrange(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.int)(unsafe.Pointer(param0))

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_layout_iter_get_line_yrange(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_iter_get_run(paramInstance unsafe.Pointer) *GlyphItem {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_get_run(cValueInstance)

	return (*GlyphItem)(ret)
}

func Fn_pango_layout_iter_get_run_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_iter_get_run_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_iter_next_char(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_next_char(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_next_cluster(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_next_cluster(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_next_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_next_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_next_run(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_next_run(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_line_get_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutLine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_line_get_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_line_get_pixel_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutLine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_layout_line_get_pixel_extents(cValueInstance, cValue0, cValue1)
}

func Fn_pango_layout_line_get_x_ranges(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *[]int, param3 *int) {
	cValueInstance := (*C.PangoLayoutLine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (C.int)(param1)

	var cValue2ArrayPointer (*C.int)
	cValue2 := &cValue2ArrayPointer

	cValue3 := (*C.int)(unsafe.Pointer(param3))

	C.pango_layout_line_get_x_ranges(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	param2OutLen := int(*cValue3)
	param2Out := make([]int, param2OutLen, param2OutLen)
	if param2OutLen > 0 {
		param2Out = (*[1 << 30](int))(unsafe.Pointer(cValue2ArrayPointer))[:param2OutLen:param2OutLen]
	}
	*param2 = param2Out
}

func Fn_pango_layout_line_index_to_x(paramInstance unsafe.Pointer, param0 int, param1 bool, param2 *int) {
	cValueInstance := (*C.PangoLayoutLine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := toCBool(param1)

	cValue2 := (*C.int)(unsafe.Pointer(param2))

	C.pango_layout_line_index_to_x(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_pango_layout_line_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayoutLine)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_line_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_line_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoLayoutLine)(unsafe.Pointer(paramInstance))

	C.pango_layout_line_unref(cValueInstance)
}

func Fn_pango_layout_line_x_to_index(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) bool {
	cValueInstance := (*C.PangoLayoutLine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	cValue2 := (*C.int)(unsafe.Pointer(param2))

	ret := C.pango_layout_line_x_to_index(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_map_get_engine : blacklisted

// UNSUPPORTED : pango_map_get_engines : blacklisted

func Fn_pango_matrix_concat(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoMatrix)(unsafe.Pointer(param0))

	C.pango_matrix_concat(cValueInstance, cValue0)
}

func Fn_pango_matrix_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	ret := C.pango_matrix_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_matrix_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	C.pango_matrix_free(cValueInstance)
}

func Fn_pango_matrix_get_font_scale_factor(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	ret := C.pango_matrix_get_font_scale_factor(cValueInstance)

	return (float64)(ret)
}

func Fn_pango_matrix_rotate(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (C.double)(param0)

	C.pango_matrix_rotate(cValueInstance, cValue0)
}

func Fn_pango_matrix_scale(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (C.double)(param0)

	cValue1 := (C.double)(param1)

	C.pango_matrix_scale(cValueInstance, cValue0, cValue1)
}

func Fn_pango_matrix_translate(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (C.double)(param0)

	cValue1 := (C.double)(param1)

	C.pango_matrix_translate(cValueInstance, cValue0, cValue1)
}

func Fn_pango_script_iter_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoScriptIter)(unsafe.Pointer(paramInstance))

	C.pango_script_iter_free(cValueInstance)
}

// UNSUPPORTED : pango_script_iter_get_range : parameter 'start' is non array with indirect count > 1

func Fn_pango_script_iter_next(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoScriptIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_script_iter_next(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_script_iter_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	ret := C.pango_script_iter_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_tab_array_new(param0 int, param1 bool) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	ret := C.pango_tab_array_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_pango_tab_array_new_with_positions(param0 int, param1 bool, param2 int, param3 int) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	cValue2 := (C.PangoTabAlign)(param2)

	cValue3 := (C.gint)(param3)

	ret := C.c_pango_tab_array_new_with_positions(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_pango_tab_array_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoTabArray)(unsafe.Pointer(paramInstance))

	ret := C.pango_tab_array_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_tab_array_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoTabArray)(unsafe.Pointer(paramInstance))

	C.pango_tab_array_free(cValueInstance)
}

func Fn_pango_tab_array_get_positions_in_pixels(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoTabArray)(unsafe.Pointer(paramInstance))

	ret := C.pango_tab_array_get_positions_in_pixels(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_tab_array_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoTabArray)(unsafe.Pointer(paramInstance))

	ret := C.pango_tab_array_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_pango_tab_array_get_tab(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.PangoTabArray)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.PangoTabAlign)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.pango_tab_array_get_tab(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'alignments' is non array with indirect count > 1

func Fn_pango_tab_array_resize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoTabArray)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.pango_tab_array_resize(cValueInstance, cValue0)
}

func Fn_pango_tab_array_set_tab(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) {
	cValueInstance := (*C.PangoTabArray)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.PangoTabAlign)(param1)

	cValue2 := (C.gint)(param2)

	C.pango_tab_array_set_tab(cValueInstance, cValue0, cValue1, cValue2)
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

func Fn_pango_attr_foreground_new(param0 uint16, param1 uint16, param2 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	ret := C.pango_attr_foreground_new(cValue0, cValue1, cValue2)

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

func Fn_pango_break(param0 string, param1 int, param2 unsafe.Pointer, param3 []LogAttr, param4 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (*C.PangoAnalysis)(unsafe.Pointer(param2))

	cValue3 := (*C.PangoLogAttr)(unsafe.Pointer(&param3[0]))

	cValue4 := (C.int)(param4)

	C.pango_break(cValue0, cValue1, cValue2, cValue3, cValue4)
}

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

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

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

func Fn_pango_get_log_attrs(param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 []LogAttr, param5 int) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (*C.PangoLanguage)(unsafe.Pointer(param3))

	cValue4 := (*C.PangoLogAttr)(unsafe.Pointer(&param4[0]))

	cValue5 := (C.int)(param5)

	C.pango_get_log_attrs(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_pango_get_mirror_char(param0 rune, param1 *rune) bool {
	cValue0 := (C.gunichar)(param0)

	cValue1 := (*C.gunichar)(unsafe.Pointer(param1))

	ret := C.pango_get_mirror_char(cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

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

func Fn_pango_log2vis_get_embedding_levels(param0 string, param1 int, param2 *int) *uint8 {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	cValue2 := (*C.PangoDirection)(unsafe.Pointer(param2))

	ret := C.pango_log2vis_get_embedding_levels(cValue0, cValue1, cValue2)

	return (*uint8)(ret)
}

// UNSUPPORTED : pango_lookup_aliases : blacklisted

// UNSUPPORTED : pango_markup_parser_finish : parameter 'attr_list' is non array with indirect count > 1

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_parse_enum : parameter 'possible_values' is non array with indirect count > 1

// UNSUPPORTED : pango_parse_markup : parameter 'attr_list' is non array with indirect count > 1

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

// UNSUPPORTED : pango_scan_int : parameter 'pos' is non array with indirect count > 1

// UNSUPPORTED : pango_scan_string : parameter 'pos' is non array with indirect count > 1

// UNSUPPORTED : pango_scan_word : parameter 'pos' is non array with indirect count > 1

func Fn_pango_shape(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.PangoAnalysis)(unsafe.Pointer(param2))

	cValue3 := (*C.PangoGlyphString)(unsafe.Pointer(param3))

	C.pango_shape(cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : pango_skip_space : parameter 'pos' is non array with indirect count > 1

// UNSUPPORTED : pango_split_file_list : no array length

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

func Fn_pango_context_new() unsafe.Pointer {
	ret := C.pango_context_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_base_dir(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_base_dir(cValueInstance)

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

func Fn_pango_context_list_families(paramInstance unsafe.Pointer, param0 *[]unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer (**C.PangoFontFamily)
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_context_list_families(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]unsafe.Pointer, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0Out = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
	}
	*param0 = param0Out
}

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

func Fn_pango_font_descriptions_free(param0 []unsafe.Pointer, param1 int) {
	cValue0 := (**C.PangoFontDescription)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.int)(param1)

	C.pango_font_descriptions_free(cValue0, cValue1)
}

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

func Fn_pango_font_face_list_sizes(paramInstance unsafe.Pointer, param0 *[]int, param1 *int) {
	cValueInstance := (*C.PangoFontFace)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer (*C.int)
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_font_face_list_sizes(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]int, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0Out = (*[1 << 30](int))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
	}
	*param0 = param0Out
}

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

func Fn_pango_font_family_list_faces(paramInstance unsafe.Pointer, param0 *[]unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.PangoFontFamily)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer (**C.PangoFontFace)
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_font_family_list_faces(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]unsafe.Pointer, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0Out = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
	}
	*param0 = param0Out
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

func Fn_pango_font_map_list_families(paramInstance unsafe.Pointer, param0 *[]unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer (**C.PangoFontFamily)
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.pango_font_map_list_families(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]unsafe.Pointer, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0Out = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
	}
	*param0 = param0Out
}

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

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

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

func Fn_pango_layout_get_lines(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_lines(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_log_attrs(paramInstance unsafe.Pointer, param0 *[]unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer (*C.PangoLogAttr)
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.pango_layout_get_log_attrs(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]unsafe.Pointer, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0Out = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
	}
	*param0 = param0Out
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

func Fn_pango_renderer_get_color(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	ret := C.pango_renderer_get_color(cValueInstance, cValue0)

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

// Code generated - DO NOT EDIT.
// +build pango_1.2

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
type Rectangle C.PangoRectangle
type RendererPrivate C.PangoRendererPrivate
type ScriptIter C.PangoScriptIter
type TabArray C.PangoTabArray

func Fn_pango_attr_font_desc_new(desc unsafe.Pointer) unsafe.Pointer {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_attr_font_desc_new(c_desc)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_iterator_copy(iterator unsafe.Pointer) unsafe.Pointer {
	c_iterator := (*C.PangoAttrIterator)(unsafe.Pointer(iterator))

	ret := C.pango_attr_iterator_copy(c_iterator)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_iterator_destroy(iterator unsafe.Pointer) {
	c_iterator := (*C.PangoAttrIterator)(unsafe.Pointer(iterator))

	C.pango_attr_iterator_destroy(c_iterator)
}

func Fn_pango_attr_iterator_get(iterator unsafe.Pointer, type_ int) unsafe.Pointer {
	c_iterator := (*C.PangoAttrIterator)(unsafe.Pointer(iterator))

	c_type_ := (C.PangoAttrType)(type_)

	ret := C.pango_attr_iterator_get(c_iterator, c_type_)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_iterator_get_attrs(iterator unsafe.Pointer) unsafe.Pointer {
	c_iterator := (*C.PangoAttrIterator)(unsafe.Pointer(iterator))

	ret := C.pango_attr_iterator_get_attrs(c_iterator)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_iterator_get_font : parameter 'language' is non array with indirect count > 1

func Fn_pango_attr_iterator_next(iterator unsafe.Pointer) bool {
	c_iterator := (*C.PangoAttrIterator)(unsafe.Pointer(iterator))

	ret := C.pango_attr_iterator_next(c_iterator)

	return toGoBool(ret)
}

func Fn_pango_attr_iterator_range(iterator unsafe.Pointer, start *int, end *int) {
	c_iterator := (*C.PangoAttrIterator)(unsafe.Pointer(iterator))

	c_start := (*C.gint)(unsafe.Pointer(start))

	c_end := (*C.gint)(unsafe.Pointer(end))

	C.pango_attr_iterator_range(c_iterator, c_start, c_end)
}

func Fn_pango_attr_language_new(language unsafe.Pointer) unsafe.Pointer {
	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_attr_language_new(c_language)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_list_new() unsafe.Pointer {
	ret := C.pango_attr_list_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_list_change(list unsafe.Pointer, attr unsafe.Pointer) {
	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	c_attr := (*C.PangoAttribute)(unsafe.Pointer(attr))

	C.pango_attr_list_change(c_list, c_attr)
}

func Fn_pango_attr_list_copy(list unsafe.Pointer) unsafe.Pointer {
	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	ret := C.pango_attr_list_copy(c_list)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

func Fn_pango_attr_list_get_iterator(list unsafe.Pointer) unsafe.Pointer {
	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	ret := C.pango_attr_list_get_iterator(c_list)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_list_insert(list unsafe.Pointer, attr unsafe.Pointer) {
	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	c_attr := (*C.PangoAttribute)(unsafe.Pointer(attr))

	C.pango_attr_list_insert(c_list, c_attr)
}

func Fn_pango_attr_list_insert_before(list unsafe.Pointer, attr unsafe.Pointer) {
	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	c_attr := (*C.PangoAttribute)(unsafe.Pointer(attr))

	C.pango_attr_list_insert_before(c_list, c_attr)
}

func Fn_pango_attr_list_splice(list unsafe.Pointer, other unsafe.Pointer, pos int, len_ int) {
	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	c_other := (*C.PangoAttrList)(unsafe.Pointer(other))

	c_pos := (C.gint)(pos)

	c_len_ := (C.gint)(len_)

	C.pango_attr_list_splice(c_list, c_other, c_pos, c_len_)
}

func Fn_pango_attr_list_unref(list unsafe.Pointer) {
	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	C.pango_attr_list_unref(c_list)
}

func Fn_pango_attr_shape_new(inkRect unsafe.Pointer, logicalRect unsafe.Pointer) unsafe.Pointer {
	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	ret := C.pango_attr_shape_new(c_inkRect, c_logicalRect)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

func Fn_pango_attr_size_new(size int) unsafe.Pointer {
	c_size := (C.int)(size)

	ret := C.pango_attr_size_new(c_size)

	return unsafe.Pointer(ret)
}

func Fn_pango_attribute_copy(attr unsafe.Pointer) unsafe.Pointer {
	c_attr := (*C.PangoAttribute)(unsafe.Pointer(attr))

	ret := C.pango_attribute_copy(c_attr)

	return unsafe.Pointer(ret)
}

func Fn_pango_attribute_destroy(attr unsafe.Pointer) {
	c_attr := (*C.PangoAttribute)(unsafe.Pointer(attr))

	C.pango_attribute_destroy(c_attr)
}

func Fn_pango_attribute_equal(attr1 unsafe.Pointer, attr2 unsafe.Pointer) bool {
	c_attr1 := (*C.PangoAttribute)(unsafe.Pointer(attr1))

	c_attr2 := (*C.PangoAttribute)(unsafe.Pointer(attr2))

	ret := C.pango_attribute_equal(c_attr1, c_attr2)

	return toGoBool(ret)
}

func Fn_pango_color_copy(src unsafe.Pointer) unsafe.Pointer {
	c_src := (*C.PangoColor)(unsafe.Pointer(src))

	ret := C.pango_color_copy(c_src)

	return unsafe.Pointer(ret)
}

func Fn_pango_color_free(color unsafe.Pointer) {
	c_color := (*C.PangoColor)(unsafe.Pointer(color))

	C.pango_color_free(c_color)
}

func Fn_pango_color_parse(color unsafe.Pointer, spec string) bool {
	c_color := (*C.PangoColor)(unsafe.Pointer(color))

	c_spec := (*C.char)(C.CString(spec))
	defer C.free(unsafe.Pointer(c_spec))

	ret := C.pango_color_parse(c_color, c_spec)

	return toGoBool(ret)
}

func Fn_pango_coverage_copy(coverage unsafe.Pointer) unsafe.Pointer {
	c_coverage := (*C.PangoCoverage)(unsafe.Pointer(coverage))

	ret := C.pango_coverage_copy(c_coverage)

	return unsafe.Pointer(ret)
}

func Fn_pango_coverage_get(coverage unsafe.Pointer, index int) int {
	c_coverage := (*C.PangoCoverage)(unsafe.Pointer(coverage))

	c_index := (C.int)(index)

	ret := C.pango_coverage_get(c_coverage, c_index)

	return (int)(ret)
}

func Fn_pango_coverage_max(coverage unsafe.Pointer, other unsafe.Pointer) {
	c_coverage := (*C.PangoCoverage)(unsafe.Pointer(coverage))

	c_other := (*C.PangoCoverage)(unsafe.Pointer(other))

	C.pango_coverage_max(c_coverage, c_other)
}

func Fn_pango_coverage_ref(coverage unsafe.Pointer) unsafe.Pointer {
	c_coverage := (*C.PangoCoverage)(unsafe.Pointer(coverage))

	ret := C.pango_coverage_ref(c_coverage)

	return unsafe.Pointer(ret)
}

func Fn_pango_coverage_set(coverage unsafe.Pointer, index int, level int) {
	c_coverage := (*C.PangoCoverage)(unsafe.Pointer(coverage))

	c_index := (C.int)(index)

	c_level := (C.PangoCoverageLevel)(level)

	C.pango_coverage_set(c_coverage, c_index, c_level)
}

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

func Fn_pango_coverage_unref(coverage unsafe.Pointer) {
	c_coverage := (*C.PangoCoverage)(unsafe.Pointer(coverage))

	C.pango_coverage_unref(c_coverage)
}

func Fn_pango_coverage_from_bytes(bytes []uint8, nBytes int) unsafe.Pointer {
	c_bytes := (*C.guchar)(unsafe.Pointer(&bytes[0]))

	c_nBytes := (C.int)(nBytes)

	ret := C.pango_coverage_from_bytes(c_bytes, c_nBytes)

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

func Fn_pango_font_description_better_match(desc unsafe.Pointer, oldMatch unsafe.Pointer, newMatch unsafe.Pointer) bool {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_oldMatch := (*C.PangoFontDescription)(unsafe.Pointer(oldMatch))

	c_newMatch := (*C.PangoFontDescription)(unsafe.Pointer(newMatch))

	ret := C.pango_font_description_better_match(c_desc, c_oldMatch, c_newMatch)

	return toGoBool(ret)
}

func Fn_pango_font_description_copy(desc unsafe.Pointer) unsafe.Pointer {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_copy(c_desc)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_description_copy_static(desc unsafe.Pointer) unsafe.Pointer {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_copy_static(c_desc)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_description_equal(desc1 unsafe.Pointer, desc2 unsafe.Pointer) bool {
	c_desc1 := (*C.PangoFontDescription)(unsafe.Pointer(desc1))

	c_desc2 := (*C.PangoFontDescription)(unsafe.Pointer(desc2))

	ret := C.pango_font_description_equal(c_desc1, c_desc2)

	return toGoBool(ret)
}

func Fn_pango_font_description_free(desc unsafe.Pointer) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	C.pango_font_description_free(c_desc)
}

func Fn_pango_font_description_get_family(desc unsafe.Pointer) string {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_get_family(c_desc)

	return C.GoString(ret)
}

func Fn_pango_font_description_get_set_fields(desc unsafe.Pointer) int {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_get_set_fields(c_desc)

	return (int)(ret)
}

func Fn_pango_font_description_get_size(desc unsafe.Pointer) int {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_get_size(c_desc)

	return (int)(ret)
}

func Fn_pango_font_description_get_stretch(desc unsafe.Pointer) int {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_get_stretch(c_desc)

	return (int)(ret)
}

func Fn_pango_font_description_get_style(desc unsafe.Pointer) int {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_get_style(c_desc)

	return (int)(ret)
}

func Fn_pango_font_description_get_variant(desc unsafe.Pointer) int {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_get_variant(c_desc)

	return (int)(ret)
}

func Fn_pango_font_description_get_weight(desc unsafe.Pointer) int {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_get_weight(c_desc)

	return (int)(ret)
}

func Fn_pango_font_description_hash(desc unsafe.Pointer) uint {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_hash(c_desc)

	return (uint)(ret)
}

func Fn_pango_font_description_merge(desc unsafe.Pointer, descToMerge unsafe.Pointer, replaceExisting bool) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_descToMerge := (*C.PangoFontDescription)(unsafe.Pointer(descToMerge))

	c_replaceExisting := toCBool(replaceExisting)

	C.pango_font_description_merge(c_desc, c_descToMerge, c_replaceExisting)
}

func Fn_pango_font_description_merge_static(desc unsafe.Pointer, descToMerge unsafe.Pointer, replaceExisting bool) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_descToMerge := (*C.PangoFontDescription)(unsafe.Pointer(descToMerge))

	c_replaceExisting := toCBool(replaceExisting)

	C.pango_font_description_merge_static(c_desc, c_descToMerge, c_replaceExisting)
}

func Fn_pango_font_description_set_family(desc unsafe.Pointer, family string) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_family := (*C.char)(C.CString(family))
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family(c_desc, c_family)
}

func Fn_pango_font_description_set_family_static(desc unsafe.Pointer, family string) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_family := (*C.char)(C.CString(family))
	defer C.free(unsafe.Pointer(c_family))

	C.pango_font_description_set_family_static(c_desc, c_family)
}

func Fn_pango_font_description_set_size(desc unsafe.Pointer, size int) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_size := (C.gint)(size)

	C.pango_font_description_set_size(c_desc, c_size)
}

func Fn_pango_font_description_set_stretch(desc unsafe.Pointer, stretch int) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_stretch := (C.PangoStretch)(stretch)

	C.pango_font_description_set_stretch(c_desc, c_stretch)
}

func Fn_pango_font_description_set_style(desc unsafe.Pointer, style int) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_style := (C.PangoStyle)(style)

	C.pango_font_description_set_style(c_desc, c_style)
}

func Fn_pango_font_description_set_variant(desc unsafe.Pointer, variant int) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_variant := (C.PangoVariant)(variant)

	C.pango_font_description_set_variant(c_desc, c_variant)
}

func Fn_pango_font_description_set_weight(desc unsafe.Pointer, weight int) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_weight := (C.PangoWeight)(weight)

	C.pango_font_description_set_weight(c_desc, c_weight)
}

func Fn_pango_font_description_to_filename(desc unsafe.Pointer) string {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_to_filename(c_desc)

	return C.GoString(ret)
}

func Fn_pango_font_description_to_string(desc unsafe.Pointer) string {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_description_to_string(c_desc)

	return C.GoString(ret)
}

func Fn_pango_font_description_unset_fields(desc unsafe.Pointer, toUnset int) {
	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_toUnset := (C.PangoFontMask)(toUnset)

	C.pango_font_description_unset_fields(c_desc, c_toUnset)
}

func Fn_pango_font_description_from_string(str string) unsafe.Pointer {
	c_str := (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	ret := C.pango_font_description_from_string(c_str)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_font_metrics_new : blacklisted

func Fn_pango_font_metrics_get_approximate_char_width(metrics unsafe.Pointer) int {
	c_metrics := (*C.PangoFontMetrics)(unsafe.Pointer(metrics))

	ret := C.pango_font_metrics_get_approximate_char_width(c_metrics)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_approximate_digit_width(metrics unsafe.Pointer) int {
	c_metrics := (*C.PangoFontMetrics)(unsafe.Pointer(metrics))

	ret := C.pango_font_metrics_get_approximate_digit_width(c_metrics)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_ascent(metrics unsafe.Pointer) int {
	c_metrics := (*C.PangoFontMetrics)(unsafe.Pointer(metrics))

	ret := C.pango_font_metrics_get_ascent(c_metrics)

	return (int)(ret)
}

func Fn_pango_font_metrics_get_descent(metrics unsafe.Pointer) int {
	c_metrics := (*C.PangoFontMetrics)(unsafe.Pointer(metrics))

	ret := C.pango_font_metrics_get_descent(c_metrics)

	return (int)(ret)
}

func Fn_pango_font_metrics_ref(metrics unsafe.Pointer) unsafe.Pointer {
	c_metrics := (*C.PangoFontMetrics)(unsafe.Pointer(metrics))

	ret := C.pango_font_metrics_ref(c_metrics)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_metrics_unref(metrics unsafe.Pointer) {
	c_metrics := (*C.PangoFontMetrics)(unsafe.Pointer(metrics))

	C.pango_font_metrics_unref(c_metrics)
}

func Fn_pango_glyph_item_apply_attrs(glyphItem unsafe.Pointer, text string, list unsafe.Pointer) unsafe.Pointer {
	c_glyphItem := (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem))

	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_list := (*C.PangoAttrList)(unsafe.Pointer(list))

	ret := C.pango_glyph_item_apply_attrs(c_glyphItem, c_text, c_list)

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_item_split(orig unsafe.Pointer, text string, splitIndex int) unsafe.Pointer {
	c_orig := (*C.PangoGlyphItem)(unsafe.Pointer(orig))

	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_splitIndex := (C.int)(splitIndex)

	ret := C.pango_glyph_item_split(c_orig, c_text, c_splitIndex)

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_string_new() unsafe.Pointer {
	ret := C.pango_glyph_string_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_string_copy(string_ unsafe.Pointer) unsafe.Pointer {
	c_string_ := (*C.PangoGlyphString)(unsafe.Pointer(string_))

	ret := C.pango_glyph_string_copy(c_string_)

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_string_extents(glyphs unsafe.Pointer, font unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_glyphs := (*C.PangoGlyphString)(unsafe.Pointer(glyphs))

	c_font := (*C.PangoFont)(unsafe.Pointer(font))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_glyph_string_extents(c_glyphs, c_font, c_inkRect, c_logicalRect)
}

func Fn_pango_glyph_string_extents_range(glyphs unsafe.Pointer, start int, end int, font unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_glyphs := (*C.PangoGlyphString)(unsafe.Pointer(glyphs))

	c_start := (C.int)(start)

	c_end := (C.int)(end)

	c_font := (*C.PangoFont)(unsafe.Pointer(font))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_glyph_string_extents_range(c_glyphs, c_start, c_end, c_font, c_inkRect, c_logicalRect)
}

func Fn_pango_glyph_string_free(string_ unsafe.Pointer) {
	c_string_ := (*C.PangoGlyphString)(unsafe.Pointer(string_))

	C.pango_glyph_string_free(c_string_)
}

func Fn_pango_glyph_string_get_logical_widths(glyphs unsafe.Pointer, text string, length int, embeddingLevel int, logicalWidths []int) {
	c_glyphs := (*C.PangoGlyphString)(unsafe.Pointer(glyphs))

	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	c_embeddingLevel := (C.int)(embeddingLevel)

	c_logicalWidths := (*C.int)(unsafe.Pointer(&logicalWidths[0]))

	C.pango_glyph_string_get_logical_widths(c_glyphs, c_text, c_length, c_embeddingLevel, c_logicalWidths)
}

func Fn_pango_glyph_string_index_to_x(glyphs unsafe.Pointer, text string, length int, analysis unsafe.Pointer, index int, trailing bool, xPos *int) {
	c_glyphs := (*C.PangoGlyphString)(unsafe.Pointer(glyphs))

	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	c_analysis := (*C.PangoAnalysis)(unsafe.Pointer(analysis))

	c_index := (C.int)(index)

	c_trailing := toCBool(trailing)

	c_xPos := (*C.int)(unsafe.Pointer(xPos))

	C.pango_glyph_string_index_to_x(c_glyphs, c_text, c_length, c_analysis, c_index, c_trailing, c_xPos)
}

func Fn_pango_glyph_string_set_size(string_ unsafe.Pointer, newLen int) {
	c_string_ := (*C.PangoGlyphString)(unsafe.Pointer(string_))

	c_newLen := (C.gint)(newLen)

	C.pango_glyph_string_set_size(c_string_, c_newLen)
}

func Fn_pango_glyph_string_x_to_index(glyphs unsafe.Pointer, text string, length int, analysis unsafe.Pointer, xPos int, index *int, trailing *int) {
	c_glyphs := (*C.PangoGlyphString)(unsafe.Pointer(glyphs))

	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	c_analysis := (*C.PangoAnalysis)(unsafe.Pointer(analysis))

	c_xPos := (C.int)(xPos)

	c_index := (*C.int)(unsafe.Pointer(index))

	c_trailing := (*C.int)(unsafe.Pointer(trailing))

	C.pango_glyph_string_x_to_index(c_glyphs, c_text, c_length, c_analysis, c_xPos, c_index, c_trailing)
}

func Fn_pango_item_new() unsafe.Pointer {
	ret := C.pango_item_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_item_copy(item unsafe.Pointer) unsafe.Pointer {
	c_item := (*C.PangoItem)(unsafe.Pointer(item))

	ret := C.pango_item_copy(c_item)

	return unsafe.Pointer(ret)
}

func Fn_pango_item_free(item unsafe.Pointer) {
	c_item := (*C.PangoItem)(unsafe.Pointer(item))

	C.pango_item_free(c_item)
}

func Fn_pango_item_split(orig unsafe.Pointer, splitIndex int, splitOffset int) unsafe.Pointer {
	c_orig := (*C.PangoItem)(unsafe.Pointer(orig))

	c_splitIndex := (C.int)(splitIndex)

	c_splitOffset := (C.int)(splitOffset)

	ret := C.pango_item_split(c_orig, c_splitIndex, c_splitOffset)

	return unsafe.Pointer(ret)
}

func Fn_pango_language_get_sample_string(language unsafe.Pointer) string {
	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_language_get_sample_string(c_language)

	return C.GoString(ret)
}

func Fn_pango_language_matches(language unsafe.Pointer, rangeList string) bool {
	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	c_rangeList := (*C.char)(C.CString(rangeList))
	defer C.free(unsafe.Pointer(c_rangeList))

	ret := C.pango_language_matches(c_language, c_rangeList)

	return toGoBool(ret)
}

func Fn_pango_language_to_string(language unsafe.Pointer) string {
	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_language_to_string(c_language)

	return C.GoString(ret)
}

func Fn_pango_language_from_string(language *string) unsafe.Pointer {
	var c_languageValue (*C.char)
	if language != nil {
		c_languageValue = (*C.char)(C.CString(*language))
		defer C.free(unsafe.Pointer(c_languageValue))
	}
	c_language := c_languageValue

	ret := C.pango_language_from_string(c_language)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_iter_at_last_line(iter unsafe.Pointer) bool {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_at_last_line(c_iter)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_free(iter unsafe.Pointer) {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	C.pango_layout_iter_free(c_iter)
}

func Fn_pango_layout_iter_get_baseline(iter unsafe.Pointer) int {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_get_baseline(c_iter)

	return (int)(ret)
}

func Fn_pango_layout_iter_get_char_extents(iter unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_iter_get_char_extents(c_iter, c_logicalRect)
}

func Fn_pango_layout_iter_get_cluster_extents(iter unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_iter_get_cluster_extents(c_iter, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_iter_get_index(iter unsafe.Pointer) int {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_get_index(c_iter)

	return (int)(ret)
}

func Fn_pango_layout_iter_get_layout_extents(iter unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_iter_get_layout_extents(c_iter, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_iter_get_line(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_get_line(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_iter_get_line_extents(iter unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_iter_get_line_extents(c_iter, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_iter_get_line_yrange(iter unsafe.Pointer, y0 *int, y1 *int) {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	c_y0 := (*C.int)(unsafe.Pointer(y0))

	c_y1 := (*C.int)(unsafe.Pointer(y1))

	C.pango_layout_iter_get_line_yrange(c_iter, c_y0, c_y1)
}

func Fn_pango_layout_iter_get_run(iter unsafe.Pointer) *GlyphItem {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_get_run(c_iter)

	return (*GlyphItem)(ret)
}

func Fn_pango_layout_iter_get_run_extents(iter unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_iter_get_run_extents(c_iter, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_iter_next_char(iter unsafe.Pointer) bool {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_next_char(c_iter)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_next_cluster(iter unsafe.Pointer) bool {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_next_cluster(c_iter)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_next_line(iter unsafe.Pointer) bool {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_next_line(c_iter)

	return toGoBool(ret)
}

func Fn_pango_layout_iter_next_run(iter unsafe.Pointer) bool {
	c_iter := (*C.PangoLayoutIter)(unsafe.Pointer(iter))

	ret := C.pango_layout_iter_next_run(c_iter)

	return toGoBool(ret)
}

func Fn_pango_layout_line_get_extents(line unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_line := (*C.PangoLayoutLine)(unsafe.Pointer(line))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_line_get_extents(c_line, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_line_get_pixel_extents(layoutLine unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_layoutLine := (*C.PangoLayoutLine)(unsafe.Pointer(layoutLine))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_line_get_pixel_extents(c_layoutLine, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_line_get_x_ranges(line unsafe.Pointer, startIndex int, endIndex int, ranges *[]int, nRanges *int) {
	c_line := (*C.PangoLayoutLine)(unsafe.Pointer(line))

	c_startIndex := (C.int)(startIndex)

	c_endIndex := (C.int)(endIndex)

	var c_rangesArrayPointer (*C.int)
	c_ranges := &c_rangesArrayPointer

	c_nRanges := (*C.int)(unsafe.Pointer(nRanges))

	C.pango_layout_line_get_x_ranges(c_line, c_startIndex, c_endIndex, c_ranges, c_nRanges)

	rangesOutLen := int(*c_nRanges)
	rangesOut := make([]int, rangesOutLen, rangesOutLen)
	if rangesOutLen > 0 {
		rangesOut = (*[1 << 30](int))(unsafe.Pointer(c_rangesArrayPointer))[:rangesOutLen:rangesOutLen]
	}
	*ranges = rangesOut
}

func Fn_pango_layout_line_index_to_x(line unsafe.Pointer, index int, trailing bool, xPos *int) {
	c_line := (*C.PangoLayoutLine)(unsafe.Pointer(line))

	c_index := (C.int)(index)

	c_trailing := toCBool(trailing)

	c_xPos := (*C.int)(unsafe.Pointer(xPos))

	C.pango_layout_line_index_to_x(c_line, c_index, c_trailing, c_xPos)
}

func Fn_pango_layout_line_unref(line unsafe.Pointer) {
	c_line := (*C.PangoLayoutLine)(unsafe.Pointer(line))

	C.pango_layout_line_unref(c_line)
}

func Fn_pango_layout_line_x_to_index(line unsafe.Pointer, xPos int, index *int, trailing *int) bool {
	c_line := (*C.PangoLayoutLine)(unsafe.Pointer(line))

	c_xPos := (C.int)(xPos)

	c_index := (*C.int)(unsafe.Pointer(index))

	c_trailing := (*C.int)(unsafe.Pointer(trailing))

	ret := C.pango_layout_line_x_to_index(c_line, c_xPos, c_index, c_trailing)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_map_get_engine : blacklisted

// UNSUPPORTED : pango_map_get_engines : blacklisted

// UNSUPPORTED : pango_script_iter_get_range : parameter 'start' is non array with indirect count > 1

func Fn_pango_tab_array_new(initialSize int, positionsInPixels bool) unsafe.Pointer {
	c_initialSize := (C.gint)(initialSize)

	c_positionsInPixels := toCBool(positionsInPixels)

	ret := C.pango_tab_array_new(c_initialSize, c_positionsInPixels)

	return unsafe.Pointer(ret)
}

func Fn_pango_tab_array_new_with_positions(size int, positionsInPixels bool, firstAlignment int, firstPosition int) unsafe.Pointer {
	c_size := (C.gint)(size)

	c_positionsInPixels := toCBool(positionsInPixels)

	c_firstAlignment := (C.PangoTabAlign)(firstAlignment)

	c_firstPosition := (C.gint)(firstPosition)

	ret := C.c_pango_tab_array_new_with_positions(c_size, c_positionsInPixels, c_firstAlignment, c_firstPosition)

	return unsafe.Pointer(ret)
}

func Fn_pango_tab_array_copy(src unsafe.Pointer) unsafe.Pointer {
	c_src := (*C.PangoTabArray)(unsafe.Pointer(src))

	ret := C.pango_tab_array_copy(c_src)

	return unsafe.Pointer(ret)
}

func Fn_pango_tab_array_free(tabArray unsafe.Pointer) {
	c_tabArray := (*C.PangoTabArray)(unsafe.Pointer(tabArray))

	C.pango_tab_array_free(c_tabArray)
}

func Fn_pango_tab_array_get_positions_in_pixels(tabArray unsafe.Pointer) bool {
	c_tabArray := (*C.PangoTabArray)(unsafe.Pointer(tabArray))

	ret := C.pango_tab_array_get_positions_in_pixels(c_tabArray)

	return toGoBool(ret)
}

func Fn_pango_tab_array_get_size(tabArray unsafe.Pointer) int {
	c_tabArray := (*C.PangoTabArray)(unsafe.Pointer(tabArray))

	ret := C.pango_tab_array_get_size(c_tabArray)

	return (int)(ret)
}

func Fn_pango_tab_array_get_tab(tabArray unsafe.Pointer, tabIndex int, alignment *int, location *int) {
	c_tabArray := (*C.PangoTabArray)(unsafe.Pointer(tabArray))

	c_tabIndex := (C.gint)(tabIndex)

	c_alignment := (*C.PangoTabAlign)(unsafe.Pointer(alignment))

	c_location := (*C.gint)(unsafe.Pointer(location))

	C.pango_tab_array_get_tab(c_tabArray, c_tabIndex, c_alignment, c_location)
}

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'alignments' is non array with indirect count > 1

func Fn_pango_tab_array_resize(tabArray unsafe.Pointer, newSize int) {
	c_tabArray := (*C.PangoTabArray)(unsafe.Pointer(tabArray))

	c_newSize := (C.gint)(newSize)

	C.pango_tab_array_resize(c_tabArray, c_newSize)
}

func Fn_pango_tab_array_set_tab(tabArray unsafe.Pointer, tabIndex int, alignment int, location int) {
	c_tabArray := (*C.PangoTabArray)(unsafe.Pointer(tabArray))

	c_tabIndex := (C.gint)(tabIndex)

	c_alignment := (C.PangoTabAlign)(alignment)

	c_location := (C.gint)(location)

	C.pango_tab_array_set_tab(c_tabArray, c_tabIndex, c_alignment, c_location)
}

func Fn_pango_attr_background_new(red uint16, green uint16, blue uint16) unsafe.Pointer {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	ret := C.pango_attr_background_new(c_red, c_green, c_blue)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_family_new(family string) unsafe.Pointer {
	c_family := (*C.char)(C.CString(family))
	defer C.free(unsafe.Pointer(c_family))

	ret := C.pango_attr_family_new(c_family)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_foreground_new(red uint16, green uint16, blue uint16) unsafe.Pointer {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	ret := C.pango_attr_foreground_new(c_red, c_green, c_blue)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_rise_new(rise int) unsafe.Pointer {
	c_rise := (C.int)(rise)

	ret := C.pango_attr_rise_new(c_rise)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_scale_new(scaleFactor float64) unsafe.Pointer {
	c_scaleFactor := (C.double)(scaleFactor)

	ret := C.pango_attr_scale_new(c_scaleFactor)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_stretch_new(stretch int) unsafe.Pointer {
	c_stretch := (C.PangoStretch)(stretch)

	ret := C.pango_attr_stretch_new(c_stretch)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_strikethrough_new(strikethrough bool) unsafe.Pointer {
	c_strikethrough := toCBool(strikethrough)

	ret := C.pango_attr_strikethrough_new(c_strikethrough)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_style_new(style int) unsafe.Pointer {
	c_style := (C.PangoStyle)(style)

	ret := C.pango_attr_style_new(c_style)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_underline_new(underline int) unsafe.Pointer {
	c_underline := (C.PangoUnderline)(underline)

	ret := C.pango_attr_underline_new(c_underline)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_variant_new(variant int) unsafe.Pointer {
	c_variant := (C.PangoVariant)(variant)

	ret := C.pango_attr_variant_new(c_variant)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_weight_new(weight int) unsafe.Pointer {
	c_weight := (C.PangoWeight)(weight)

	ret := C.pango_attr_weight_new(c_weight)

	return unsafe.Pointer(ret)
}

func Fn_pango_break(text string, length int, analysis unsafe.Pointer, attrs []LogAttr, attrsLen int) {
	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	c_analysis := (*C.PangoAnalysis)(unsafe.Pointer(analysis))

	c_attrs := (*C.PangoLogAttr)(unsafe.Pointer(&attrs[0]))

	c_attrsLen := (C.int)(attrsLen)

	C.pango_break(c_text, c_length, c_analysis, c_attrs, c_attrsLen)
}

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// UNSUPPORTED : pango_find_map : blacklisted

func Fn_pango_find_paragraph_boundary(text string, length int, paragraphDelimiterIndex *int, nextParagraphStart *int) {
	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gint)(length)

	c_paragraphDelimiterIndex := (*C.gint)(unsafe.Pointer(paragraphDelimiterIndex))

	c_nextParagraphStart := (*C.gint)(unsafe.Pointer(nextParagraphStart))

	C.pango_find_paragraph_boundary(c_text, c_length, c_paragraphDelimiterIndex, c_nextParagraphStart)
}

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

func Fn_pango_get_log_attrs(text string, length int, level int, language unsafe.Pointer, logAttrs []LogAttr, attrsLen int) {
	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	c_level := (C.int)(level)

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	c_logAttrs := (*C.PangoLogAttr)(unsafe.Pointer(&logAttrs[0]))

	c_attrsLen := (C.int)(attrsLen)

	C.pango_get_log_attrs(c_text, c_length, c_level, c_language, c_logAttrs, c_attrsLen)
}

func Fn_pango_get_mirror_char(ch rune, mirroredCh *rune) bool {
	c_ch := (C.gunichar)(ch)

	c_mirroredCh := (*C.gunichar)(unsafe.Pointer(mirroredCh))

	ret := C.pango_get_mirror_char(c_ch, c_mirroredCh)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

func Fn_pango_itemize(context unsafe.Pointer, text string, startIndex int, length int, attrs unsafe.Pointer, cachedIter unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_startIndex := (C.int)(startIndex)

	c_length := (C.int)(length)

	c_attrs := (*C.PangoAttrList)(unsafe.Pointer(attrs))

	c_cachedIter := (*C.PangoAttrIterator)(unsafe.Pointer(cachedIter))

	ret := C.pango_itemize(c_context, c_text, c_startIndex, c_length, c_attrs, c_cachedIter)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_lookup_aliases : blacklisted

// UNSUPPORTED : pango_markup_parser_finish : parameter 'attr_list' is non array with indirect count > 1

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_parse_enum : parameter 'possible_values' is non array with indirect count > 1

// UNSUPPORTED : pango_parse_markup : parameter 'attr_list' is non array with indirect count > 1

func Fn_pango_parse_stretch(str string, stretch *int, warn bool) bool {
	c_str := (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_stretch := (*C.PangoStretch)(unsafe.Pointer(stretch))

	c_warn := toCBool(warn)

	ret := C.pango_parse_stretch(c_str, c_stretch, c_warn)

	return toGoBool(ret)
}

func Fn_pango_parse_style(str string, style *int, warn bool) bool {
	c_str := (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_style := (*C.PangoStyle)(unsafe.Pointer(style))

	c_warn := toCBool(warn)

	ret := C.pango_parse_style(c_str, c_style, c_warn)

	return toGoBool(ret)
}

func Fn_pango_parse_variant(str string, variant *int, warn bool) bool {
	c_str := (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_variant := (*C.PangoVariant)(unsafe.Pointer(variant))

	c_warn := toCBool(warn)

	ret := C.pango_parse_variant(c_str, c_variant, c_warn)

	return toGoBool(ret)
}

func Fn_pango_parse_weight(str string, weight *int, warn bool) bool {
	c_str := (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_weight := (*C.PangoWeight)(unsafe.Pointer(weight))

	c_warn := toCBool(warn)

	ret := C.pango_parse_weight(c_str, c_weight, c_warn)

	return toGoBool(ret)
}

func Fn_pango_read_line(stream unsafe.Pointer, str unsafe.Pointer) int {
	c_stream := (*C.FILE)(unsafe.Pointer(stream))

	c_str := (*C.GString)(unsafe.Pointer(str))

	ret := C.pango_read_line(c_stream, c_str)

	return (int)(ret)
}

func Fn_pango_reorder_items(logicalItems unsafe.Pointer) unsafe.Pointer {
	c_logicalItems := (*C.GList)(unsafe.Pointer(logicalItems))

	ret := C.pango_reorder_items(c_logicalItems)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_scan_int : parameter 'pos' is non array with indirect count > 1

// UNSUPPORTED : pango_scan_string : parameter 'pos' is non array with indirect count > 1

// UNSUPPORTED : pango_scan_word : parameter 'pos' is non array with indirect count > 1

func Fn_pango_shape(text string, length int, analysis unsafe.Pointer, glyphs unsafe.Pointer) {
	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gint)(length)

	c_analysis := (*C.PangoAnalysis)(unsafe.Pointer(analysis))

	c_glyphs := (*C.PangoGlyphString)(unsafe.Pointer(glyphs))

	C.pango_shape(c_text, c_length, c_analysis, c_glyphs)
}

// UNSUPPORTED : pango_skip_space : parameter 'pos' is non array with indirect count > 1

// UNSUPPORTED : pango_split_file_list : no array length

func Fn_pango_trim_string(str string) string {
	c_str := (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	ret := C.pango_trim_string(c_str)

	return C.GoString(ret)
}

func Fn_pango_unichar_direction(ch rune) int {
	c_ch := (C.gunichar)(ch)

	ret := C.pango_unichar_direction(c_ch)

	return (int)(ret)
}

func Fn_pango_context_new() unsafe.Pointer {
	ret := C.pango_context_new()

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_base_dir(context unsafe.Pointer) int {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	ret := C.pango_context_get_base_dir(c_context)

	return (int)(ret)
}

func Fn_pango_context_get_font_description(context unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	ret := C.pango_context_get_font_description(c_context)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_language(context unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	ret := C.pango_context_get_language(c_context)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_metrics(context unsafe.Pointer, desc unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_context_get_metrics(c_context, c_desc, c_language)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_list_families(context unsafe.Pointer, families *[]unsafe.Pointer, nFamilies *int) {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	var c_familiesArrayPointer (**C.PangoFontFamily)
	c_families := &c_familiesArrayPointer

	c_nFamilies := (*C.int)(unsafe.Pointer(nFamilies))

	C.pango_context_list_families(c_context, c_families, c_nFamilies)

	familiesOutLen := int(*c_nFamilies)
	familiesOut := make([]unsafe.Pointer, familiesOutLen, familiesOutLen)
	if familiesOutLen > 0 {
		familiesOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_familiesArrayPointer))[:familiesOutLen:familiesOutLen]
	}
	*families = familiesOut
}

func Fn_pango_context_load_font(context unsafe.Pointer, desc unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_context_load_font(c_context, c_desc)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_load_fontset(context unsafe.Pointer, desc unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_context_load_fontset(c_context, c_desc, c_language)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_set_base_dir(context unsafe.Pointer, direction int) {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_direction := (C.PangoDirection)(direction)

	C.pango_context_set_base_dir(c_context, c_direction)
}

func Fn_pango_context_set_font_description(context unsafe.Pointer, desc unsafe.Pointer) {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	C.pango_context_set_font_description(c_context, c_desc)
}

func Fn_pango_context_set_font_map(context unsafe.Pointer, fontMap unsafe.Pointer) {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_fontMap := (*C.PangoFontMap)(unsafe.Pointer(fontMap))

	C.pango_context_set_font_map(c_context, c_fontMap)
}

func Fn_pango_context_set_language(context unsafe.Pointer, language unsafe.Pointer) {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	C.pango_context_set_language(c_context, c_language)
}

func Fn_pango_font_describe(font unsafe.Pointer) unsafe.Pointer {
	c_font := (*C.PangoFont)(unsafe.Pointer(font))

	ret := C.pango_font_describe(c_font)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_find_shaper(font unsafe.Pointer, language unsafe.Pointer, ch uint32) unsafe.Pointer {
	c_font := (*C.PangoFont)(unsafe.Pointer(font))

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	c_ch := (C.guint32)(ch)

	ret := C.pango_font_find_shaper(c_font, c_language, c_ch)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_get_coverage(font unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	c_font := (*C.PangoFont)(unsafe.Pointer(font))

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_font_get_coverage(c_font, c_language)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_get_glyph_extents(font unsafe.Pointer, glyph uint32, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_font := (*C.PangoFont)(unsafe.Pointer(font))

	c_glyph := (C.PangoGlyph)(glyph)

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_font_get_glyph_extents(c_font, c_glyph, c_inkRect, c_logicalRect)
}

func Fn_pango_font_get_metrics(font unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	c_font := (*C.PangoFont)(unsafe.Pointer(font))

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_font_get_metrics(c_font, c_language)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_descriptions_free(descs []unsafe.Pointer, nDescs int) {
	c_descs := (**C.PangoFontDescription)(unsafe.Pointer(&descs[0]))

	c_nDescs := (C.int)(nDescs)

	C.pango_font_descriptions_free(c_descs, c_nDescs)
}

func Fn_pango_font_face_describe(face unsafe.Pointer) unsafe.Pointer {
	c_face := (*C.PangoFontFace)(unsafe.Pointer(face))

	ret := C.pango_font_face_describe(c_face)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_face_get_face_name(face unsafe.Pointer) string {
	c_face := (*C.PangoFontFace)(unsafe.Pointer(face))

	ret := C.pango_font_face_get_face_name(c_face)

	return C.GoString(ret)
}

func Fn_pango_font_family_get_name(family unsafe.Pointer) string {
	c_family := (*C.PangoFontFamily)(unsafe.Pointer(family))

	ret := C.pango_font_family_get_name(c_family)

	return C.GoString(ret)
}

func Fn_pango_font_family_list_faces(family unsafe.Pointer, faces *[]unsafe.Pointer, nFaces *int) {
	c_family := (*C.PangoFontFamily)(unsafe.Pointer(family))

	var c_facesArrayPointer (**C.PangoFontFace)
	c_faces := &c_facesArrayPointer

	c_nFaces := (*C.int)(unsafe.Pointer(nFaces))

	C.pango_font_family_list_faces(c_family, c_faces, c_nFaces)

	facesOutLen := int(*c_nFaces)
	facesOut := make([]unsafe.Pointer, facesOutLen, facesOutLen)
	if facesOutLen > 0 {
		facesOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_facesArrayPointer))[:facesOutLen:facesOutLen]
	}
	*faces = facesOut
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

func Fn_pango_font_map_list_families(fontmap unsafe.Pointer, families *[]unsafe.Pointer, nFamilies *int) {
	c_fontmap := (*C.PangoFontMap)(unsafe.Pointer(fontmap))

	var c_familiesArrayPointer (**C.PangoFontFamily)
	c_families := &c_familiesArrayPointer

	c_nFamilies := (*C.int)(unsafe.Pointer(nFamilies))

	C.pango_font_map_list_families(c_fontmap, c_families, c_nFamilies)

	familiesOutLen := int(*c_nFamilies)
	familiesOut := make([]unsafe.Pointer, familiesOutLen, familiesOutLen)
	if familiesOutLen > 0 {
		familiesOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_familiesArrayPointer))[:familiesOutLen:familiesOutLen]
	}
	*families = familiesOut
}

func Fn_pango_font_map_load_font(fontmap unsafe.Pointer, context unsafe.Pointer, desc unsafe.Pointer) unsafe.Pointer {
	c_fontmap := (*C.PangoFontMap)(unsafe.Pointer(fontmap))

	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	ret := C.pango_font_map_load_font(c_fontmap, c_context, c_desc)

	return unsafe.Pointer(ret)
}

func Fn_pango_font_map_load_fontset(fontmap unsafe.Pointer, context unsafe.Pointer, desc unsafe.Pointer, language unsafe.Pointer) unsafe.Pointer {
	c_fontmap := (*C.PangoFontMap)(unsafe.Pointer(fontmap))

	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	c_language := (*C.PangoLanguage)(unsafe.Pointer(language))

	ret := C.pango_font_map_load_fontset(c_fontmap, c_context, c_desc, c_language)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

func Fn_pango_fontset_get_font(fontset unsafe.Pointer, wc uint) unsafe.Pointer {
	c_fontset := (*C.PangoFontset)(unsafe.Pointer(fontset))

	c_wc := (C.guint)(wc)

	ret := C.pango_fontset_get_font(c_fontset, c_wc)

	return unsafe.Pointer(ret)
}

func Fn_pango_fontset_get_metrics(fontset unsafe.Pointer) unsafe.Pointer {
	c_fontset := (*C.PangoFontset)(unsafe.Pointer(fontset))

	ret := C.pango_fontset_get_metrics(c_fontset)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

func Fn_pango_layout_new(context unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.PangoContext)(unsafe.Pointer(context))

	ret := C.pango_layout_new(c_context)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_context_changed(layout unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	C.pango_layout_context_changed(c_layout)
}

func Fn_pango_layout_copy(src unsafe.Pointer) unsafe.Pointer {
	c_src := (*C.PangoLayout)(unsafe.Pointer(src))

	ret := C.pango_layout_copy(c_src)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_alignment(layout unsafe.Pointer) int {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_alignment(c_layout)

	return (int)(ret)
}

func Fn_pango_layout_get_attributes(layout unsafe.Pointer) unsafe.Pointer {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_attributes(c_layout)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_context(layout unsafe.Pointer) unsafe.Pointer {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_context(c_layout)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_cursor_pos(layout unsafe.Pointer, index int, strongPos unsafe.Pointer, weakPos unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_index := (C.int)(index)

	c_strongPos := (*C.PangoRectangle)(unsafe.Pointer(strongPos))

	c_weakPos := (*C.PangoRectangle)(unsafe.Pointer(weakPos))

	C.pango_layout_get_cursor_pos(c_layout, c_index, c_strongPos, c_weakPos)
}

func Fn_pango_layout_get_extents(layout unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_get_extents(c_layout, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_get_indent(layout unsafe.Pointer) int {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_indent(c_layout)

	return (int)(ret)
}

func Fn_pango_layout_get_iter(layout unsafe.Pointer) unsafe.Pointer {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_iter(c_layout)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_justify(layout unsafe.Pointer) bool {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_justify(c_layout)

	return toGoBool(ret)
}

func Fn_pango_layout_get_line(layout unsafe.Pointer, line int) unsafe.Pointer {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_line := (C.int)(line)

	ret := C.pango_layout_get_line(c_layout, c_line)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_line_count(layout unsafe.Pointer) int {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_line_count(c_layout)

	return (int)(ret)
}

func Fn_pango_layout_get_lines(layout unsafe.Pointer) unsafe.Pointer {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_lines(c_layout)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_log_attrs(layout unsafe.Pointer, attrs *[]unsafe.Pointer, nAttrs *int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	var c_attrsArrayPointer (*C.PangoLogAttr)
	c_attrs := &c_attrsArrayPointer

	c_nAttrs := (*C.gint)(unsafe.Pointer(nAttrs))

	C.pango_layout_get_log_attrs(c_layout, c_attrs, c_nAttrs)

	attrsOutLen := int(*c_nAttrs)
	attrsOut := make([]unsafe.Pointer, attrsOutLen, attrsOutLen)
	if attrsOutLen > 0 {
		attrsOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_attrsArrayPointer))[:attrsOutLen:attrsOutLen]
	}
	*attrs = attrsOut
}

func Fn_pango_layout_get_pixel_extents(layout unsafe.Pointer, inkRect unsafe.Pointer, logicalRect unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_inkRect := (*C.PangoRectangle)(unsafe.Pointer(inkRect))

	c_logicalRect := (*C.PangoRectangle)(unsafe.Pointer(logicalRect))

	C.pango_layout_get_pixel_extents(c_layout, c_inkRect, c_logicalRect)
}

func Fn_pango_layout_get_pixel_size(layout unsafe.Pointer, width *int, height *int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_width := (*C.int)(unsafe.Pointer(width))

	c_height := (*C.int)(unsafe.Pointer(height))

	C.pango_layout_get_pixel_size(c_layout, c_width, c_height)
}

func Fn_pango_layout_get_single_paragraph_mode(layout unsafe.Pointer) bool {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_single_paragraph_mode(c_layout)

	return toGoBool(ret)
}

func Fn_pango_layout_get_size(layout unsafe.Pointer, width *int, height *int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_width := (*C.int)(unsafe.Pointer(width))

	c_height := (*C.int)(unsafe.Pointer(height))

	C.pango_layout_get_size(c_layout, c_width, c_height)
}

func Fn_pango_layout_get_spacing(layout unsafe.Pointer) int {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_spacing(c_layout)

	return (int)(ret)
}

func Fn_pango_layout_get_tabs(layout unsafe.Pointer) unsafe.Pointer {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_tabs(c_layout)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_text(layout unsafe.Pointer) string {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_text(c_layout)

	return C.GoString(ret)
}

func Fn_pango_layout_get_width(layout unsafe.Pointer) int {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_width(c_layout)

	return (int)(ret)
}

func Fn_pango_layout_get_wrap(layout unsafe.Pointer) int {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	ret := C.pango_layout_get_wrap(c_layout)

	return (int)(ret)
}

func Fn_pango_layout_index_to_line_x(layout unsafe.Pointer, index int, trailing bool, line *int, xPos *int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_index := (C.int)(index)

	c_trailing := toCBool(trailing)

	c_line := (*C.int)(unsafe.Pointer(line))

	c_xPos := (*C.int)(unsafe.Pointer(xPos))

	C.pango_layout_index_to_line_x(c_layout, c_index, c_trailing, c_line, c_xPos)
}

func Fn_pango_layout_index_to_pos(layout unsafe.Pointer, index int, pos unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_index := (C.int)(index)

	c_pos := (*C.PangoRectangle)(unsafe.Pointer(pos))

	C.pango_layout_index_to_pos(c_layout, c_index, c_pos)
}

func Fn_pango_layout_move_cursor_visually(layout unsafe.Pointer, strong bool, oldIndex int, oldTrailing int, direction int, newIndex *int, newTrailing *int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_strong := toCBool(strong)

	c_oldIndex := (C.int)(oldIndex)

	c_oldTrailing := (C.int)(oldTrailing)

	c_direction := (C.int)(direction)

	c_newIndex := (*C.int)(unsafe.Pointer(newIndex))

	c_newTrailing := (*C.int)(unsafe.Pointer(newTrailing))

	C.pango_layout_move_cursor_visually(c_layout, c_strong, c_oldIndex, c_oldTrailing, c_direction, c_newIndex, c_newTrailing)
}

func Fn_pango_layout_set_alignment(layout unsafe.Pointer, alignment int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_alignment := (C.PangoAlignment)(alignment)

	C.pango_layout_set_alignment(c_layout, c_alignment)
}

func Fn_pango_layout_set_attributes(layout unsafe.Pointer, attrs unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_attrs := (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.pango_layout_set_attributes(c_layout, c_attrs)
}

func Fn_pango_layout_set_font_description(layout unsafe.Pointer, desc unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_desc := (*C.PangoFontDescription)(unsafe.Pointer(desc))

	C.pango_layout_set_font_description(c_layout, c_desc)
}

func Fn_pango_layout_set_indent(layout unsafe.Pointer, indent int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_indent := (C.int)(indent)

	C.pango_layout_set_indent(c_layout, c_indent)
}

func Fn_pango_layout_set_justify(layout unsafe.Pointer, justify bool) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_justify := toCBool(justify)

	C.pango_layout_set_justify(c_layout, c_justify)
}

func Fn_pango_layout_set_markup(layout unsafe.Pointer, markup string, length int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_markup := (*C.char)(C.CString(markup))
	defer C.free(unsafe.Pointer(c_markup))

	c_length := (C.int)(length)

	C.pango_layout_set_markup(c_layout, c_markup, c_length)
}

func Fn_pango_layout_set_markup_with_accel(layout unsafe.Pointer, markup string, length int, accelMarker rune, accelChar *rune) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_markup := (*C.char)(C.CString(markup))
	defer C.free(unsafe.Pointer(c_markup))

	c_length := (C.int)(length)

	c_accelMarker := (C.gunichar)(accelMarker)

	c_accelChar := (*C.gunichar)(unsafe.Pointer(accelChar))

	C.pango_layout_set_markup_with_accel(c_layout, c_markup, c_length, c_accelMarker, c_accelChar)
}

func Fn_pango_layout_set_single_paragraph_mode(layout unsafe.Pointer, setting bool) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_setting := toCBool(setting)

	C.pango_layout_set_single_paragraph_mode(c_layout, c_setting)
}

func Fn_pango_layout_set_spacing(layout unsafe.Pointer, spacing int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_spacing := (C.int)(spacing)

	C.pango_layout_set_spacing(c_layout, c_spacing)
}

func Fn_pango_layout_set_tabs(layout unsafe.Pointer, tabs unsafe.Pointer) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_tabs := (*C.PangoTabArray)(unsafe.Pointer(tabs))

	C.pango_layout_set_tabs(c_layout, c_tabs)
}

func Fn_pango_layout_set_text(layout unsafe.Pointer, text string, length int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_text := (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.int)(length)

	C.pango_layout_set_text(c_layout, c_text, c_length)
}

func Fn_pango_layout_set_width(layout unsafe.Pointer, width int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_width := (C.int)(width)

	C.pango_layout_set_width(c_layout, c_width)
}

func Fn_pango_layout_set_wrap(layout unsafe.Pointer, wrap int) {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_wrap := (C.PangoWrapMode)(wrap)

	C.pango_layout_set_wrap(c_layout, c_wrap)
}

func Fn_pango_layout_xy_to_index(layout unsafe.Pointer, x int, y int, index *int, trailing *int) bool {
	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_index := (*C.int)(unsafe.Pointer(index))

	c_trailing := (*C.int)(unsafe.Pointer(trailing))

	ret := C.pango_layout_xy_to_index(c_layout, c_x, c_y, c_index, c_trailing)

	return toGoBool(ret)
}

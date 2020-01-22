// Code generated - DO NOT EDIT.
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.24 pango_1.26 pango_1.30 pango_1.31 pango_1.32 pango_1.32.4 pango_1.34 pango_1.36.7 pango_1.38 pango_1.42

package pango

import "unsafe"

// #include <pango/pango.h>
// #include <pango/pango-font.h>
// #include <pango/pango-modules.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : EngineClass : blacklisted
// UNSUPPORTED : EngineInfo : blacklisted
// UNSUPPORTED : EngineLangClass : blacklisted
// UNSUPPORTED : EngineScriptInfo : blacklisted
// UNSUPPORTED : EngineShapeClass : blacklisted
// UNSUPPORTED : FontClass : blacklisted
// UNSUPPORTED : FontFaceClass : blacklisted
// UNSUPPORTED : FontFamilyClass : blacklisted
// UNSUPPORTED : FontMapClass : blacklisted
// UNSUPPORTED : FontsetClass : blacklisted
// UNSUPPORTED : FontsetSimpleClass : blacklisted
// UNSUPPORTED : IncludedModule : blacklisted
// UNSUPPORTED : Map : blacklisted
// UNSUPPORTED : MapEntry : blacklisted
// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

func Fn_pango_color_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.PangoColor)(unsafe.Pointer(paramInstance))

	ret := C.pango_color_to_string(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

func Fn_pango_font_description_get_gravity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_gravity(cValueInstance)

	return (int)(ret)
}

func Fn_pango_font_description_set_gravity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoGravity)(param0)

	C.pango_font_description_set_gravity(cValueInstance, cValue0)
}

// UNSUPPORTED : pango_font_metrics_new : blacklisted

// UNSUPPORTED : pango_glyph_item_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_item_letter_space : parameter 'log_attrs' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_string_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

func Fn_pango_language_get_default() unsafe.Pointer {
	ret := C.pango_language_get_default()

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_iter_get_line_readonly(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_get_line_readonly(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_iter_get_run_readonly(paramInstance unsafe.Pointer) *GlyphItem {
	cValueInstance := (*C.PangoLayoutIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_iter_get_run_readonly(cValueInstance)

	return (*GlyphItem)(ret)
}

// UNSUPPORTED : pango_map_get_engine : blacklisted

// UNSUPPORTED : pango_map_get_engines : blacklisted

func Fn_pango_matrix_transform_distance(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.double)(unsafe.Pointer(param0))

	cValue1 := (*C.double)(unsafe.Pointer(param1))

	C.pango_matrix_transform_distance(cValueInstance, cValue0, cValue1)
}

func Fn_pango_matrix_transform_pixel_rectangle(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	C.pango_matrix_transform_pixel_rectangle(cValueInstance, cValue0)
}

func Fn_pango_matrix_transform_point(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.double)(unsafe.Pointer(param0))

	cValue1 := (*C.double)(unsafe.Pointer(param1))

	C.pango_matrix_transform_point(cValueInstance, cValue0, cValue1)
}

func Fn_pango_matrix_transform_rectangle(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	C.pango_matrix_transform_rectangle(cValueInstance, cValue0)
}

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'locations' is array parameter without length parameter

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

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

func Fn_pango_extents_to_pixels(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.PangoRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.PangoRectangle)(unsafe.Pointer(param1))

	C.pango_extents_to_pixels(cValue0, cValue1)
}

// UNSUPPORTED : pango_find_map : blacklisted

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

// UNSUPPORTED : pango_lookup_aliases : blacklisted

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

// UNSUPPORTED : pango_split_file_list : no array length

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

func Fn_pango_context_get_base_gravity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_base_gravity(cValueInstance)

	return (int)(ret)
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

func Fn_pango_context_set_base_gravity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoGravity)(param0)

	C.pango_context_set_base_gravity(cValueInstance, cValue0)
}

func Fn_pango_context_set_gravity_hint(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoGravityHint)(param0)

	C.pango_context_set_gravity_hint(cValueInstance, cValue0)
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

func Fn_pango_layout_get_line_readonly(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	ret := C.pango_layout_get_line_readonly(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_lines_readonly(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_lines_readonly(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_layout_get_unknown_glyphs_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_unknown_glyphs_count(cValueInstance)

	return (int)(ret)
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

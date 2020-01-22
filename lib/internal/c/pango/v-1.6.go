// Code generated - DO NOT EDIT.
// +build pango_1.6

package pango

import "unsafe"

// #include <pango/pango.h>
// #include <pango/pango-font.h>
// #include <pango/pango-modules.h>
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
type Matrix C.PangoMatrix

// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

// UNSUPPORTED : pango_font_metrics_new : blacklisted

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

func Fn_pango_glyph_item_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoGlyphItem)(unsafe.Pointer(paramInstance))

	C.pango_glyph_item_free(cValueInstance)
}

// UNSUPPORTED : pango_glyph_item_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_item_letter_space : parameter 'log_attrs' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_string_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

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

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'locations' is array parameter without length parameter

func Fn_pango_attr_letter_spacing_new(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.pango_attr_letter_spacing_new(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// UNSUPPORTED : pango_find_map : blacklisted

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

// UNSUPPORTED : pango_lookup_aliases : blacklisted

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_split_file_list : no array length

func Fn_pango_context_get_font_map(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_font_map(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_get_matrix(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_matrix(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_context_set_matrix(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoMatrix)(unsafe.Pointer(param0))

	C.pango_context_set_matrix(cValueInstance, cValue0)
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

func Fn_pango_layout_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_pango_layout_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.pango_layout_set_ellipsize(cValueInstance, cValue0)
}

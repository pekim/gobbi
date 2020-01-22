// Code generated - DO NOT EDIT.
// +build pango_1.8

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
type RendererClass C.PangoRendererClass

// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

func Fn_pango_attr_size_new_absolute(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.pango_attr_size_new_absolute(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

func Fn_pango_font_description_get_size_is_absolute(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_description_get_size_is_absolute(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_font_description_set_absolute_size(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.PangoFontDescription)(unsafe.Pointer(paramInstance))

	cValue0 := (C.double)(param0)

	C.pango_font_description_set_absolute_size(cValueInstance, cValue0)
}

// UNSUPPORTED : pango_font_metrics_new : blacklisted

// UNSUPPORTED : pango_glyph_item_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_item_letter_space : parameter 'log_attrs' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_string_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_map_get_engine : blacklisted

// UNSUPPORTED : pango_map_get_engines : blacklisted

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'locations' is array parameter without length parameter

func Fn_pango_attr_strikethrough_color_new(param0 uint16, param1 uint16, param2 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	ret := C.pango_attr_strikethrough_color_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_underline_color_new(param0 uint16, param1 uint16, param2 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	ret := C.pango_attr_underline_color_new(cValue0, cValue1, cValue2)

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

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

func Fn_pango_layout_get_font_description(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_font_description(cValueInstance)

	return unsafe.Pointer(ret)
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

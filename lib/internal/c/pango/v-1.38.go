// Code generated - DO NOT EDIT.
// +build pango_1.38 pango_1.42

package pango

import "unsafe"

// #include <pango/pango.h>
// #include <pango/pango-font.h>
// #include <pango/pango-modules.h>
// #include <stdlib.h>
import "C"

type AttrFontFeatures C.PangoAttrFontFeatures

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
func Fn_pango_attr_font_features_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.pango_attr_font_features_new(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

// UNSUPPORTED : pango_font_metrics_new : blacklisted

// UNSUPPORTED : pango_glyph_item_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_item_letter_space : parameter 'log_attrs' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_string_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_map_get_engine : blacklisted

// UNSUPPORTED : pango_map_get_engines : blacklisted

func Fn_pango_matrix_get_font_scale_factors(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
	cValueInstance := (*C.PangoMatrix)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.double)(unsafe.Pointer(param0))

	cValue1 := (*C.double)(unsafe.Pointer(param1))

	C.pango_matrix_get_font_scale_factors(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'locations' is array parameter without length parameter

func Fn_pango_attr_background_alpha_new(param0 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	ret := C.pango_attr_background_alpha_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_pango_attr_foreground_alpha_new(param0 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	ret := C.pango_attr_foreground_alpha_new(cValue0)

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

func Fn_pango_renderer_get_alpha(paramInstance unsafe.Pointer, param0 int) uint16 {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	ret := C.pango_renderer_get_alpha(cValueInstance, cValue0)

	return (uint16)(ret)
}

func Fn_pango_renderer_set_alpha(paramInstance unsafe.Pointer, param0 int, param1 uint16) {
	cValueInstance := (*C.PangoRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoRenderPart)(param0)

	cValue1 := (C.guint16)(param1)

	C.pango_renderer_set_alpha(cValueInstance, cValue0, cValue1)
}

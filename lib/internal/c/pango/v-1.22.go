// Code generated - DO NOT EDIT.
// +build pango_1.22

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
type GlyphItemIter C.PangoGlyphItemIter

// UNSUPPORTED : IncludedModule : blacklisted
// UNSUPPORTED : Map : blacklisted
// UNSUPPORTED : MapEntry : blacklisted
// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

// UNSUPPORTED : pango_font_metrics_new : blacklisted

// UNSUPPORTED : pango_glyph_item_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_item_letter_space : parameter 'log_attrs' is array parameter without length parameter

func Fn_pango_glyph_item_iter_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoGlyphItemIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_glyph_item_iter_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_pango_glyph_item_iter_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoGlyphItemIter)(unsafe.Pointer(paramInstance))

	C.pango_glyph_item_iter_free(cValueInstance)
}

func Fn_pango_glyph_item_iter_init_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) bool {
	cValueInstance := (*C.PangoGlyphItemIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoGlyphItem)(unsafe.Pointer(param0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.pango_glyph_item_iter_init_end(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_pango_glyph_item_iter_init_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) bool {
	cValueInstance := (*C.PangoGlyphItemIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoGlyphItem)(unsafe.Pointer(param0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.pango_glyph_item_iter_init_start(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_pango_glyph_item_iter_next_cluster(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoGlyphItemIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_glyph_item_iter_next_cluster(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_glyph_item_iter_prev_cluster(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoGlyphItemIter)(unsafe.Pointer(paramInstance))

	ret := C.pango_glyph_item_iter_prev_cluster(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_glyph_string_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

func Fn_pango_language_get_scripts(paramInstance unsafe.Pointer, param0 *int) []int {
	cValueInstance := (*C.PangoLanguage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.int)(unsafe.Pointer(param0))

	ret := C.pango_language_get_scripts(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]int, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](int))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : pango_map_get_engine : blacklisted

// UNSUPPORTED : pango_map_get_engines : blacklisted

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'locations' is array parameter without length parameter

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// UNSUPPORTED : pango_find_map : blacklisted

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

// UNSUPPORTED : pango_lookup_aliases : blacklisted

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_split_file_list : no array length

func Fn_pango_font_map_create_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_map_create_context(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

func Fn_pango_layout_get_baseline(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_baseline(cValueInstance)

	return (int)(ret)
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

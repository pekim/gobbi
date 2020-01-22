// Code generated - DO NOT EDIT.
// +build pango_1.2 pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.24 pango_1.26 pango_1.30 pango_1.31 pango_1.32 pango_1.32.4 pango_1.34 pango_1.36.7 pango_1.38 pango_1.42

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
func Fn_pango_attr_iterator_get_attrs(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoAttrIterator)(unsafe.Pointer(paramInstance))

	ret := C.pango_attr_iterator_get_attrs(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

// UNSUPPORTED : pango_font_metrics_new : blacklisted

func Fn_pango_glyph_item_apply_attrs(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.PangoGlyphItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.PangoAttrList)(unsafe.Pointer(param1))

	ret := C.pango_glyph_item_apply_attrs(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_glyph_item_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_item_letter_space : parameter 'log_attrs' is array parameter without length parameter

func Fn_pango_glyph_item_split(paramInstance unsafe.Pointer, param0 string, param1 int) unsafe.Pointer {
	cValueInstance := (*C.PangoGlyphItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.int)(param1)

	ret := C.pango_glyph_item_split(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : pango_glyph_string_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

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

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

// Code generated - DO NOT EDIT.
// +build pango_1.32.4

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
// UNSUPPORTED : pango_attr_list_filter : parameter 'func' is callback

// UNSUPPORTED : pango_attr_shape_new_with_data : parameter 'copy_func' is callback

// UNSUPPORTED : pango_coverage_to_bytes : blacklisted

// UNSUPPORTED : pango_font_metrics_new : blacklisted

// UNSUPPORTED : pango_glyph_item_get_logical_widths : parameter 'logical_widths' is array parameter without length parameter

// UNSUPPORTED : pango_glyph_item_letter_space : parameter 'log_attrs' is array parameter without length parameter

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

func Fn_pango_context_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	C.pango_context_changed(cValueInstance)
}

func Fn_pango_context_get_serial(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.PangoContext)(unsafe.Pointer(paramInstance))

	ret := C.pango_context_get_serial(cValueInstance)

	return (uint)(ret)
}

func Fn_pango_font_map_get_serial(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.PangoFontMap)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_map_get_serial(cValueInstance)

	return (uint)(ret)
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

func Fn_pango_layout_get_serial(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_serial(cValueInstance)

	return (uint)(ret)
}

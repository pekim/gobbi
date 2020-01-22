// Code generated - DO NOT EDIT.
// +build pango_1.4

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

func Fn_pango_language_includes_script(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.PangoLanguage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoScript)(param0)

	ret := C.pango_language_includes_script(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_map_get_engine : blacklisted

// UNSUPPORTED : pango_map_get_engines : blacklisted

func Fn_pango_script_iter_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.PangoScriptIter)(unsafe.Pointer(paramInstance))

	C.pango_script_iter_free(cValueInstance)
}

func Fn_pango_script_iter_get_range(paramInstance unsafe.Pointer, param0 *string, param1 *string, param2 *int) {
	cValueInstance := (*C.PangoScriptIter)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cValue2 := (*C.PangoScript)(unsafe.Pointer(param2))

	C.pango_script_iter_get_range(cValueInstance, cValue0, cValue1, cValue2)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	param1String := C.GoString(cValue1String)
	*param1 = param1String
}

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

// UNSUPPORTED : pango_tab_array_get_tabs : parameter 'locations' is array parameter without length parameter

func Fn_pango_attr_fallback_new(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.pango_attr_fallback_new(cValue0)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

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

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_split_file_list : no array length

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

func Fn_pango_font_family_is_monospace(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoFontFamily)(unsafe.Pointer(paramInstance))

	ret := C.pango_font_family_is_monospace(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : pango_font_map_get_shape_engine_type : blacklisted

// UNSUPPORTED : pango_fontset_foreach : parameter 'func' is callback

// UNSUPPORTED : pango_fontset_simple_new : blacklisted

// UNSUPPORTED : pango_fontset_simple_append : blacklisted

// UNSUPPORTED : pango_fontset_simple_size : blacklisted

func Fn_pango_layout_get_auto_dir(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	ret := C.pango_layout_get_auto_dir(cValueInstance)

	return toGoBool(ret)
}

func Fn_pango_layout_set_auto_dir(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.PangoLayout)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.pango_layout_set_auto_dir(cValueInstance, cValue0)
}

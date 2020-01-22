// Code generated - DO NOT EDIT.
// +build pango_1.4 pango_1.6 pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.24 pango_1.26 pango_1.30 pango_1.31 pango_1.32 pango_1.32.4 pango_1.34 pango_1.36.7 pango_1.38 pango_1.42

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/internal/c/pango"
)

// AttrFallbackNew wraps the C function pango_attr_fallback_new.
//
// since 1.4
func AttrFallbackNew(enableFallback bool) *Attribute {
	sys_enableFallback := enableFallback
	retSys := pango.Fn_pango_attr_fallback_new(sys_enableFallback)
	ret := AttributeNewFromC(retSys)

	return ret
}

// UNSUPPORTED : pango_break : has array param, attrs

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// FindBaseDir wraps the C function pango_find_base_dir.
//
// since 1.4
func FindBaseDir(text string, length int) int {
	sys_text := text
	sys_length := length
	retSys := pango.Fn_pango_find_base_dir(sys_text, sys_length)
	ret := retSys

	return ret
}

// UNSUPPORTED : pango_find_map : blacklisted

// UNSUPPORTED : pango_find_paragraph_boundary : has [in]out param, paragraph_delimiter_index

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_log_attrs : has array param, log_attrs

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

// ItemizeWithBaseDir wraps the C function pango_itemize_with_base_dir.
//
// since 1.4
func ItemizeWithBaseDir(context *Context, baseDir Direction, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	sys_context := context.ToC()
	sys_baseDir := (int)(baseDir)
	sys_text := text
	sys_startIndex := startIndex
	sys_length := length
	sys_attrs := attrs.ToC()
	sys_cachedIter := cachedIter.ToC()
	retSys := pango.Fn_pango_itemize_with_base_dir(sys_context, sys_baseDir, sys_text, sys_startIndex, sys_length, sys_attrs, sys_cachedIter)
	ret := glib.ListNewFromC(retSys)

	return ret
}

// Log2visGetEmbeddingLevels wraps the C function pango_log2vis_get_embedding_levels.
//
// since 1.4
func Log2visGetEmbeddingLevels(text string, length int, pbaseDir *Direction) *uint8 {
	sys_text := text
	sys_length := length
	sys_pbaseDir := (*int)(pbaseDir)
	retSys := pango.Fn_pango_log2vis_get_embedding_levels(sys_text, sys_length, sys_pbaseDir)
	ret := retSys

	return ret
}

// UNSUPPORTED : pango_lookup_aliases : blacklisted

// UNSUPPORTED : pango_markup_parser_finish : throws

// UNSUPPORTED : pango_module_register : blacklisted

// UNSUPPORTED : pango_parse_enum : has [in]out param, value

// UNSUPPORTED : pango_parse_markup : throws

// UNSUPPORTED : pango_parse_stretch : has [in]out param, stretch

// UNSUPPORTED : pango_parse_style : has [in]out param, style

// UNSUPPORTED : pango_parse_variant : has [in]out param, variant

// UNSUPPORTED : pango_parse_weight : has [in]out param, weight

// UNSUPPORTED : pango_read_line : has [in]out param, str

// UNSUPPORTED : pango_scan_int : has [in]out param, out

// UNSUPPORTED : pango_scan_string : has [in]out param, out

// UNSUPPORTED : pango_scan_word : has [in]out param, out

// UNSUPPORTED : pango_split_file_list : no array length

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

// UNSUPPORTED : Engine : blacklisted

// UNSUPPORTED : FontsetSimple : blacklisted

// Code generated - DO NOT EDIT.
// +build pango_1.38

package pango

import (
	pango "github.com/pekim/gobbi/lib/internal/c/pango"
	"unsafe"
)

// AttrBackgroundAlphaNew wraps the C function pango_attr_background_alpha_new.
//
// since 1.38
func AttrBackgroundAlphaNew(alpha uint16) *Attribute {
	sys_alpha := alpha
	retSys := pango.Fn_pango_attr_background_alpha_new(sys_alpha)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrForegroundAlphaNew wraps the C function pango_attr_foreground_alpha_new.
//
// since 1.38
func AttrForegroundAlphaNew(alpha uint16) *Attribute {
	sys_alpha := alpha
	retSys := pango.Fn_pango_attr_foreground_alpha_new(sys_alpha)
	ret := AttributeNewFromC(retSys)

	return ret
}

// UNSUPPORTED : pango_break : has array param, attrs

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// UNSUPPORTED : pango_find_map : blacklisted

// UNSUPPORTED : pango_find_paragraph_boundary : has [in]out param, paragraph_delimiter_index

// UNSUPPORTED : pango_get_lib_subdirectory : blacklisted

// UNSUPPORTED : pango_get_log_attrs : has array param, log_attrs

// UNSUPPORTED : pango_get_sysconf_subdirectory : blacklisted

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

// AttrFontFeatures is a representation of the C record PangoAttrFontFeatures.
//
// since 1.38
type AttrFontFeatures struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoAttrFontFeatures that represents the AttrFontFeatures.
func (recv *AttrFontFeatures) ToC() unsafe.Pointer {
	return recv.native
}

// AttrFontFeaturesNewFromC creates a new AttrFontFeatures from a pointer to the C PangoAttrFontFeatures that represents the AttrFontFeatures.
func AttrFontFeaturesNewFromC(native unsafe.Pointer) *AttrFontFeatures {
	return &AttrFontFeatures{native: native}
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

// UNSUPPORTED : Engine : blacklisted

// UNSUPPORTED : FontsetSimple : blacklisted

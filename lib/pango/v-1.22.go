// Code generated - DO NOT EDIT.
// +build pango_1.22 pango_1.24 pango_1.26 pango_1.30 pango_1.31 pango_1.32 pango_1.32.4 pango_1.34 pango_1.36.7 pango_1.38 pango_1.42

package pango

import "unsafe"

// BidiType is a representation of the C enumeration PangoBidiType.
type BidiType int

// BidiType_l is a representation of the C enumeration member PANGO_BIDI_TYPE_L.
const BidiType_l = BidiType(0)

// BidiType_lre is a representation of the C enumeration member PANGO_BIDI_TYPE_LRE.
const BidiType_lre = BidiType(1)

// BidiType_lro is a representation of the C enumeration member PANGO_BIDI_TYPE_LRO.
const BidiType_lro = BidiType(2)

// BidiType_r is a representation of the C enumeration member PANGO_BIDI_TYPE_R.
const BidiType_r = BidiType(3)

// BidiType_al is a representation of the C enumeration member PANGO_BIDI_TYPE_AL.
const BidiType_al = BidiType(4)

// BidiType_rle is a representation of the C enumeration member PANGO_BIDI_TYPE_RLE.
const BidiType_rle = BidiType(5)

// BidiType_rlo is a representation of the C enumeration member PANGO_BIDI_TYPE_RLO.
const BidiType_rlo = BidiType(6)

// BidiType_pdf is a representation of the C enumeration member PANGO_BIDI_TYPE_PDF.
const BidiType_pdf = BidiType(7)

// BidiType_en is a representation of the C enumeration member PANGO_BIDI_TYPE_EN.
const BidiType_en = BidiType(8)

// BidiType_es is a representation of the C enumeration member PANGO_BIDI_TYPE_ES.
const BidiType_es = BidiType(9)

// BidiType_et is a representation of the C enumeration member PANGO_BIDI_TYPE_ET.
const BidiType_et = BidiType(10)

// BidiType_an is a representation of the C enumeration member PANGO_BIDI_TYPE_AN.
const BidiType_an = BidiType(11)

// BidiType_cs is a representation of the C enumeration member PANGO_BIDI_TYPE_CS.
const BidiType_cs = BidiType(12)

// BidiType_nsm is a representation of the C enumeration member PANGO_BIDI_TYPE_NSM.
const BidiType_nsm = BidiType(13)

// BidiType_bn is a representation of the C enumeration member PANGO_BIDI_TYPE_BN.
const BidiType_bn = BidiType(14)

// BidiType_b is a representation of the C enumeration member PANGO_BIDI_TYPE_B.
const BidiType_b = BidiType(15)

// BidiType_s is a representation of the C enumeration member PANGO_BIDI_TYPE_S.
const BidiType_s = BidiType(16)

// BidiType_ws is a representation of the C enumeration member PANGO_BIDI_TYPE_WS.
const BidiType_ws = BidiType(17)

// BidiType_on is a representation of the C enumeration member PANGO_BIDI_TYPE_ON.
const BidiType_on = BidiType(18)

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

// GlyphItemIter is a representation of the C record PangoGlyphItemIter.
//
// since 1.22
type GlyphItemIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C PangoGlyphItemIter that represents the GlyphItemIter.
func (recv *GlyphItemIter) ToC() unsafe.Pointer {
	return recv.native
}

// GlyphItemIterNewFromC creates a new GlyphItemIter from a pointer to the C PangoGlyphItemIter that represents the GlyphItemIter.
func GlyphItemIterNewFromC(native unsafe.Pointer) *GlyphItemIter {
	return &GlyphItemIter{native: native}
}

// UNSUPPORTED : IncludedModule : blacklisted

// UNSUPPORTED : Map : blacklisted

// UNSUPPORTED : MapEntry : blacklisted

// UNSUPPORTED : Engine : blacklisted

// UNSUPPORTED : FontsetSimple : blacklisted

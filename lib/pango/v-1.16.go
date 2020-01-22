// Code generated - DO NOT EDIT.
// +build pango_1.16

package pango

import pango "github.com/pekim/gobbi/lib/internal/c/pango"

// ANALYSIS_FLAG_CENTERED_BASELINE is a representation of the C constant PANGO_ANALYSIS_FLAG_CENTERED_BASELINE.
//
// since 1.16
const ANALYSIS_FLAG_CENTERED_BASELINE = 1

// Gravity is a representation of the C enumeration PangoGravity.
type Gravity int

// Gravity_south is a representation of the C enumeration member PANGO_GRAVITY_SOUTH.
const Gravity_south = Gravity(0)

// Gravity_east is a representation of the C enumeration member PANGO_GRAVITY_EAST.
const Gravity_east = Gravity(1)

// Gravity_north is a representation of the C enumeration member PANGO_GRAVITY_NORTH.
const Gravity_north = Gravity(2)

// Gravity_west is a representation of the C enumeration member PANGO_GRAVITY_WEST.
const Gravity_west = Gravity(3)

// Gravity_auto is a representation of the C enumeration member PANGO_GRAVITY_AUTO.
const Gravity_auto = Gravity(4)

// GravityHint is a representation of the C enumeration PangoGravityHint.
type GravityHint int

// GravityHint_natural is a representation of the C enumeration member PANGO_GRAVITY_HINT_NATURAL.
const GravityHint_natural = GravityHint(0)

// GravityHint_strong is a representation of the C enumeration member PANGO_GRAVITY_HINT_STRONG.
const GravityHint_strong = GravityHint(1)

// GravityHint_line is a representation of the C enumeration member PANGO_GRAVITY_HINT_LINE.
const GravityHint_line = GravityHint(2)

// AttrGravityHintNew wraps the C function pango_attr_gravity_hint_new.
//
// since 1.16
func AttrGravityHintNew(hint GravityHint) *Attribute {
	sys_hint := (int)(hint)
	retSys := pango.Fn_pango_attr_gravity_hint_new(sys_hint)
	ret := AttributeNewFromC(retSys)

	return ret
}

// AttrGravityNew wraps the C function pango_attr_gravity_new.
//
// since 1.16
func AttrGravityNew(gravity Gravity) *Attribute {
	sys_gravity := (int)(gravity)
	retSys := pango.Fn_pango_attr_gravity_new(sys_gravity)
	ret := AttributeNewFromC(retSys)

	return ret
}

// UNSUPPORTED : pango_break : has array param, attrs

// UNSUPPORTED : pango_config_key_get : blacklisted

// UNSUPPORTED : pango_config_key_get_system : blacklisted

// UNSUPPORTED : pango_default_break : blacklisted

// ExtentsToPixels wraps the C function pango_extents_to_pixels.
//
// since 1.16
func ExtentsToPixels(inclusive *Rectangle, nearest *Rectangle) {
	sys_inclusive := inclusive.ToC()
	sys_nearest := nearest.ToC()
	pango.Fn_pango_extents_to_pixels(sys_inclusive, sys_nearest)
}

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

// UnitsFromDouble wraps the C function pango_units_from_double.
//
// since 1.16
func UnitsFromDouble(d float64) int {
	sys_d := d
	retSys := pango.Fn_pango_units_from_double(sys_d)
	ret := retSys

	return ret
}

// UnitsToDouble wraps the C function pango_units_to_double.
//
// since 1.16
func UnitsToDouble(i int) float64 {
	sys_i := i
	retSys := pango.Fn_pango_units_to_double(sys_i)
	ret := retSys

	return ret
}

// Version wraps the C function pango_version.
//
// since 1.16
func Version() int {
	retSys := pango.Fn_pango_version()
	ret := retSys

	return ret
}

// VersionCheck wraps the C function pango_version_check.
//
// since 1.16
func VersionCheck(requiredMajor int, requiredMinor int, requiredMicro int) string {
	sys_requiredMajor := requiredMajor
	sys_requiredMinor := requiredMinor
	sys_requiredMicro := requiredMicro
	retSys := pango.Fn_pango_version_check(sys_requiredMajor, sys_requiredMinor, sys_requiredMicro)
	ret := retSys

	return ret
}

// VersionString wraps the C function pango_version_string.
//
// since 1.16
func VersionString() string {
	retSys := pango.Fn_pango_version_string()
	ret := retSys

	return ret
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

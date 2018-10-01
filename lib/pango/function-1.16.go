// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// AttrGravityHintNew is a wrapper around the C function pango_attr_gravity_hint_new.
func AttrGravityHintNew(hint GravityHint) *Attribute {
	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_attr_gravity_hint_new(c_hint)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrGravityNew is a wrapper around the C function pango_attr_gravity_new.
func AttrGravityNew(gravity Gravity) *Attribute {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_attr_gravity_new(c_gravity)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_break : unsupported parameter attrs : no param type

// ExtentsToPixels is a wrapper around the C function pango_extents_to_pixels.
func ExtentsToPixels(inclusive *Rectangle, nearest *Rectangle) {
	c_inclusive := (*C.PangoRectangle)(inclusive.ToC())

	c_nearest := (*C.PangoRectangle)(nearest.ToC())

	C.pango_extents_to_pixels(c_inclusive, c_nearest)

	return
}

// Unsupported : pango_find_map : return type : Blacklisted record : PangoMap

// Unsupported : pango_get_log_attrs : unsupported parameter log_attrs : no param type

// GravityGetForMatrix is a wrapper around the C function pango_gravity_get_for_matrix.
func GravityGetForMatrix(matrix *Matrix) Gravity {
	c_matrix := (*C.PangoMatrix)(matrix.ToC())

	retC := C.pango_gravity_get_for_matrix(c_matrix)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityGetForScript is a wrapper around the C function pango_gravity_get_for_script.
func GravityGetForScript(script Script, baseGravity Gravity, hint GravityHint) Gravity {
	c_script := (C.PangoScript)(script)

	c_base_gravity := (C.PangoGravity)(baseGravity)

	c_hint := (C.PangoGravityHint)(hint)

	retC := C.pango_gravity_get_for_script(c_script, c_base_gravity, c_hint)
	retGo := (Gravity)(retC)

	return retGo
}

// GravityToRotation is a wrapper around the C function pango_gravity_to_rotation.
func GravityToRotation(gravity Gravity) float64 {
	c_gravity := (C.PangoGravity)(gravity)

	retC := C.pango_gravity_to_rotation(c_gravity)
	retGo := (float64)(retC)

	return retGo
}

// LanguageGetDefault is a wrapper around the C function pango_language_get_default.
func LanguageGetDefault() *Language {
	retC := C.pango_language_get_default()
	retGo := LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : pango_log2vis_get_embedding_levels : unsupported parameter pbase_dir : PangoDirection* with indirection level of 1

// Unsupported : pango_lookup_aliases : unsupported parameter families : no param type

// Unsupported : pango_markup_parser_finish : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_module_register : unsupported parameter module : Blacklisted record : PangoIncludedModule

// Unsupported : pango_parse_enum : unsupported parameter type : no type generator for GType, GType

// Unsupported : pango_parse_markup : unsupported parameter attr_list : record with indirection level of 2

// Unsupported : pango_parse_stretch : unsupported parameter stretch : PangoStretch* with indirection level of 1

// Unsupported : pango_parse_style : unsupported parameter style : PangoStyle* with indirection level of 1

// Unsupported : pango_parse_variant : unsupported parameter variant : PangoVariant* with indirection level of 1

// Unsupported : pango_parse_weight : unsupported parameter weight : PangoWeight* with indirection level of 1

// Unsupported : pango_read_line : unsupported parameter stream : no type generator for gpointer, FILE*

// Unsupported : pango_scan_int : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_string : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_scan_word : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_shape : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_shape_full : unsupported parameter glyphs : Blacklisted record : PangoGlyphString

// Unsupported : pango_skip_space : unsupported parameter pos : in string with indirection level of 2

// Unsupported : pango_split_file_list : no return type

// UnitsFromDouble is a wrapper around the C function pango_units_from_double.
func UnitsFromDouble(d float64) int32 {
	c_d := (C.double)(d)

	retC := C.pango_units_from_double(c_d)
	retGo := (int32)(retC)

	return retGo
}

// UnitsToDouble is a wrapper around the C function pango_units_to_double.
func UnitsToDouble(i int32) float64 {
	c_i := (C.int)(i)

	retC := C.pango_units_to_double(c_i)
	retGo := (float64)(retC)

	return retGo
}

// Version is a wrapper around the C function pango_version.
func Version() int32 {
	retC := C.pango_version()
	retGo := (int32)(retC)

	return retGo
}

// VersionCheck is a wrapper around the C function pango_version_check.
func VersionCheck(requiredMajor int32, requiredMinor int32, requiredMicro int32) string {
	c_required_major := (C.int)(requiredMajor)

	c_required_minor := (C.int)(requiredMinor)

	c_required_micro := (C.int)(requiredMicro)

	retC := C.pango_version_check(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// VersionString is a wrapper around the C function pango_version_string.
func VersionString() string {
	retC := C.pango_version_string()
	retGo := C.GoString(retC)

	return retGo
}

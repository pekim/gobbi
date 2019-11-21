// Code generated - DO NOT EDIT.

package pango

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var attrBackgroundAlphaNewFunction *gi.Function
var attrBackgroundAlphaNewFunction_Once sync.Once

func attrBackgroundAlphaNewFunction_Set() {
	attrBackgroundAlphaNewFunction_Once.Do(func() {
		attrBackgroundAlphaNewFunction = gi.FunctionInvokerNew("Pango", "attr_background_alpha_new")
	})
}

// AttrBackgroundAlphaNew is a representation of the C type pango_attr_background_alpha_new.
func AttrBackgroundAlphaNew(alpha uint16) *Attribute {
	attrBackgroundAlphaNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(alpha)

	ret := attrBackgroundAlphaNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var attrBackgroundNewFunction *gi.Function
var attrBackgroundNewFunction_Once sync.Once

func attrBackgroundNewFunction_Set() {
	attrBackgroundNewFunction_Once.Do(func() {
		attrBackgroundNewFunction = gi.FunctionInvokerNew("Pango", "attr_background_new")
	})
}

// AttrBackgroundNew is a representation of the C type pango_attr_background_new.
func AttrBackgroundNew(red uint16, green uint16, blue uint16) *Attribute {
	attrBackgroundNewFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetUint16(red)
	inArgs[1].SetUint16(green)
	inArgs[2].SetUint16(blue)

	ret := attrBackgroundNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_fallback_new' : parameter 'enable_fallback' of type 'gboolean' not supported

var attrFamilyNewFunction *gi.Function
var attrFamilyNewFunction_Once sync.Once

func attrFamilyNewFunction_Set() {
	attrFamilyNewFunction_Once.Do(func() {
		attrFamilyNewFunction = gi.FunctionInvokerNew("Pango", "attr_family_new")
	})
}

// AttrFamilyNew is a representation of the C type pango_attr_family_new.
func AttrFamilyNew(family string) *Attribute {
	attrFamilyNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(family)

	ret := attrFamilyNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var attrForegroundAlphaNewFunction *gi.Function
var attrForegroundAlphaNewFunction_Once sync.Once

func attrForegroundAlphaNewFunction_Set() {
	attrForegroundAlphaNewFunction_Once.Do(func() {
		attrForegroundAlphaNewFunction = gi.FunctionInvokerNew("Pango", "attr_foreground_alpha_new")
	})
}

// AttrForegroundAlphaNew is a representation of the C type pango_attr_foreground_alpha_new.
func AttrForegroundAlphaNew(alpha uint16) *Attribute {
	attrForegroundAlphaNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(alpha)

	ret := attrForegroundAlphaNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var attrForegroundNewFunction *gi.Function
var attrForegroundNewFunction_Once sync.Once

func attrForegroundNewFunction_Set() {
	attrForegroundNewFunction_Once.Do(func() {
		attrForegroundNewFunction = gi.FunctionInvokerNew("Pango", "attr_foreground_new")
	})
}

// AttrForegroundNew is a representation of the C type pango_attr_foreground_new.
func AttrForegroundNew(red uint16, green uint16, blue uint16) *Attribute {
	attrForegroundNewFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetUint16(red)
	inArgs[1].SetUint16(green)
	inArgs[2].SetUint16(blue)

	ret := attrForegroundNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_gravity_hint_new' : parameter 'hint' of type 'GravityHint' not supported

// UNSUPPORTED : C value 'pango_attr_gravity_new' : parameter 'gravity' of type 'Gravity' not supported

var attrLetterSpacingNewFunction *gi.Function
var attrLetterSpacingNewFunction_Once sync.Once

func attrLetterSpacingNewFunction_Set() {
	attrLetterSpacingNewFunction_Once.Do(func() {
		attrLetterSpacingNewFunction = gi.FunctionInvokerNew("Pango", "attr_letter_spacing_new")
	})
}

// AttrLetterSpacingNew is a representation of the C type pango_attr_letter_spacing_new.
func AttrLetterSpacingNew(letterSpacing int32) *Attribute {
	attrLetterSpacingNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(letterSpacing)

	ret := attrLetterSpacingNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

var attrRiseNewFunction *gi.Function
var attrRiseNewFunction_Once sync.Once

func attrRiseNewFunction_Set() {
	attrRiseNewFunction_Once.Do(func() {
		attrRiseNewFunction = gi.FunctionInvokerNew("Pango", "attr_rise_new")
	})
}

// AttrRiseNew is a representation of the C type pango_attr_rise_new.
func AttrRiseNew(rise int32) *Attribute {
	attrRiseNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(rise)

	ret := attrRiseNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_scale_new' : parameter 'scale_factor' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_attr_stretch_new' : parameter 'stretch' of type 'Stretch' not supported

var attrStrikethroughColorNewFunction *gi.Function
var attrStrikethroughColorNewFunction_Once sync.Once

func attrStrikethroughColorNewFunction_Set() {
	attrStrikethroughColorNewFunction_Once.Do(func() {
		attrStrikethroughColorNewFunction = gi.FunctionInvokerNew("Pango", "attr_strikethrough_color_new")
	})
}

// AttrStrikethroughColorNew is a representation of the C type pango_attr_strikethrough_color_new.
func AttrStrikethroughColorNew(red uint16, green uint16, blue uint16) *Attribute {
	attrStrikethroughColorNewFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetUint16(red)
	inArgs[1].SetUint16(green)
	inArgs[2].SetUint16(blue)

	ret := attrStrikethroughColorNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_strikethrough_new' : parameter 'strikethrough' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_attr_style_new' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'pango_attr_type_get_name' : parameter 'type' of type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_type_register' : return type 'AttrType' not supported

var attrUnderlineColorNewFunction *gi.Function
var attrUnderlineColorNewFunction_Once sync.Once

func attrUnderlineColorNewFunction_Set() {
	attrUnderlineColorNewFunction_Once.Do(func() {
		attrUnderlineColorNewFunction = gi.FunctionInvokerNew("Pango", "attr_underline_color_new")
	})
}

// AttrUnderlineColorNew is a representation of the C type pango_attr_underline_color_new.
func AttrUnderlineColorNew(red uint16, green uint16, blue uint16) *Attribute {
	attrUnderlineColorNewFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetUint16(red)
	inArgs[1].SetUint16(green)
	inArgs[2].SetUint16(blue)

	ret := attrUnderlineColorNewFunction.Invoke(inArgs[:], nil)

	retGo := &Attribute{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_underline_new' : parameter 'underline' of type 'Underline' not supported

// UNSUPPORTED : C value 'pango_attr_variant_new' : parameter 'variant' of type 'Variant' not supported

// UNSUPPORTED : C value 'pango_attr_weight_new' : parameter 'weight' of type 'Weight' not supported

// UNSUPPORTED : C value 'pango_bidi_type_for_unichar' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_break' : parameter 'analysis' of type 'Analysis' not supported

var configKeyGetFunction *gi.Function
var configKeyGetFunction_Once sync.Once

func configKeyGetFunction_Set() {
	configKeyGetFunction_Once.Do(func() {
		configKeyGetFunction = gi.FunctionInvokerNew("Pango", "config_key_get")
	})
}

// ConfigKeyGet is a representation of the C type pango_config_key_get.
func ConfigKeyGet(key string) string {
	configKeyGetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(key)

	ret := configKeyGetFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var configKeyGetSystemFunction *gi.Function
var configKeyGetSystemFunction_Once sync.Once

func configKeyGetSystemFunction_Set() {
	configKeyGetSystemFunction_Once.Do(func() {
		configKeyGetSystemFunction = gi.FunctionInvokerNew("Pango", "config_key_get_system")
	})
}

// ConfigKeyGetSystem is a representation of the C type pango_config_key_get_system.
func ConfigKeyGetSystem(key string) string {
	configKeyGetSystemFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(key)

	ret := configKeyGetSystemFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'pango_default_break' : parameter 'analysis' of type 'Analysis' not supported

// UNSUPPORTED : C value 'pango_extents_to_pixels' : parameter 'inclusive' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_find_base_dir' : return type 'Direction' not supported

// UNSUPPORTED : C value 'pango_find_map' : parameter 'language' of type 'Language' not supported

var findParagraphBoundaryFunction *gi.Function
var findParagraphBoundaryFunction_Once sync.Once

func findParagraphBoundaryFunction_Set() {
	findParagraphBoundaryFunction_Once.Do(func() {
		findParagraphBoundaryFunction = gi.FunctionInvokerNew("Pango", "find_paragraph_boundary")
	})
}

// FindParagraphBoundary is a representation of the C type pango_find_paragraph_boundary.
func FindParagraphBoundary(text string, length int32) (int32, int32) {
	findParagraphBoundaryFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(text)
	inArgs[1].SetInt32(length)

	var outArgs [2]gi.Argument

	findParagraphBoundaryFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var fontDescriptionFromStringFunction *gi.Function
var fontDescriptionFromStringFunction_Once sync.Once

func fontDescriptionFromStringFunction_Set() {
	fontDescriptionFromStringFunction_Once.Do(func() {
		fontDescriptionFromStringFunction = gi.FunctionInvokerNew("Pango", "font_description_from_string")
	})
}

// FontDescriptionFromString is a representation of the C type pango_font_description_from_string.
func FontDescriptionFromString(str string) *FontDescription {
	fontDescriptionFromStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := fontDescriptionFromStringFunction.Invoke(inArgs[:], nil)

	retGo := &FontDescription{native: ret.Pointer()}

	return retGo
}

var getLibSubdirectoryFunction *gi.Function
var getLibSubdirectoryFunction_Once sync.Once

func getLibSubdirectoryFunction_Set() {
	getLibSubdirectoryFunction_Once.Do(func() {
		getLibSubdirectoryFunction = gi.FunctionInvokerNew("Pango", "get_lib_subdirectory")
	})
}

// GetLibSubdirectory is a representation of the C type pango_get_lib_subdirectory.
func GetLibSubdirectory() string {
	getLibSubdirectoryFunction_Set()

	ret := getLibSubdirectoryFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_get_log_attrs' : parameter 'language' of type 'Language' not supported

// UNSUPPORTED : C value 'pango_get_mirror_char' : parameter 'ch' of type 'gunichar' not supported

var getSysconfSubdirectoryFunction *gi.Function
var getSysconfSubdirectoryFunction_Once sync.Once

func getSysconfSubdirectoryFunction_Set() {
	getSysconfSubdirectoryFunction_Once.Do(func() {
		getSysconfSubdirectoryFunction = gi.FunctionInvokerNew("Pango", "get_sysconf_subdirectory")
	})
}

// GetSysconfSubdirectory is a representation of the C type pango_get_sysconf_subdirectory.
func GetSysconfSubdirectory() string {
	getSysconfSubdirectoryFunction_Set()

	ret := getSysconfSubdirectoryFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'pango_gravity_get_for_matrix' : parameter 'matrix' of type 'Matrix' not supported

// UNSUPPORTED : C value 'pango_gravity_get_for_script' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_gravity_get_for_script_and_width' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_gravity_to_rotation' : parameter 'gravity' of type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_is_zero_width' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_itemize' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'pango_itemize_with_base_dir' : parameter 'context' of type 'Context' not supported

var languageFromStringFunction *gi.Function
var languageFromStringFunction_Once sync.Once

func languageFromStringFunction_Set() {
	languageFromStringFunction_Once.Do(func() {
		languageFromStringFunction = gi.FunctionInvokerNew("Pango", "language_from_string")
	})
}

// LanguageFromString is a representation of the C type pango_language_from_string.
func LanguageFromString(language string) *Language {
	languageFromStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(language)

	ret := languageFromStringFunction.Invoke(inArgs[:], nil)

	retGo := &Language{native: ret.Pointer()}

	return retGo
}

var languageGetDefaultFunction *gi.Function
var languageGetDefaultFunction_Once sync.Once

func languageGetDefaultFunction_Set() {
	languageGetDefaultFunction_Once.Do(func() {
		languageGetDefaultFunction = gi.FunctionInvokerNew("Pango", "language_get_default")
	})
}

// LanguageGetDefault is a representation of the C type pango_language_get_default.
func LanguageGetDefault() *Language {
	languageGetDefaultFunction_Set()

	ret := languageGetDefaultFunction.Invoke(nil, nil)

	retGo := &Language{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'pango_log2vis_get_embedding_levels' : parameter 'pbase_dir' of type 'Direction' not supported

// UNSUPPORTED : C value 'pango_lookup_aliases' : parameter 'families' has no type

// UNSUPPORTED : C value 'pango_markup_parser_finish' : parameter 'context' of type 'GLib.MarkupParseContext' not supported

// UNSUPPORTED : C value 'pango_markup_parser_new' : parameter 'accel_marker' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_module_register' : parameter 'module' of type 'IncludedModule' not supported

// UNSUPPORTED : C value 'pango_parse_enum' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'pango_parse_markup' : parameter 'accel_marker' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_parse_stretch' : parameter 'stretch' of type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_parse_style' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'pango_parse_variant' : parameter 'variant' of type 'Variant' not supported

// UNSUPPORTED : C value 'pango_parse_weight' : parameter 'weight' of type 'Weight' not supported

var quantizeLineGeometryFunction *gi.Function
var quantizeLineGeometryFunction_Once sync.Once

func quantizeLineGeometryFunction_Set() {
	quantizeLineGeometryFunction_Once.Do(func() {
		quantizeLineGeometryFunction = gi.FunctionInvokerNew("Pango", "quantize_line_geometry")
	})
}

// QuantizeLineGeometry is a representation of the C type pango_quantize_line_geometry.
func QuantizeLineGeometry(thickness int32, position int32) (int32, int32) {
	quantizeLineGeometryFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(thickness)
	inArgs[1].SetInt32(position)

	var outArgs [2]gi.Argument

	quantizeLineGeometryFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

// UNSUPPORTED : C value 'pango_read_line' : parameter 'stream' of type 'gpointer' not supported

// UNSUPPORTED : C value 'pango_reorder_items' : parameter 'logical_items' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'pango_scan_int' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_scan_string' : parameter 'out' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'pango_scan_word' : parameter 'out' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'pango_script_for_unichar' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_script_get_sample_language' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_shape' : parameter 'analysis' of type 'Analysis' not supported

// UNSUPPORTED : C value 'pango_shape_full' : parameter 'analysis' of type 'Analysis' not supported

// UNSUPPORTED : C value 'pango_skip_space' : return type 'gboolean' not supported

var splitFileListFunction *gi.Function
var splitFileListFunction_Once sync.Once

func splitFileListFunction_Set() {
	splitFileListFunction_Once.Do(func() {
		splitFileListFunction = gi.FunctionInvokerNew("Pango", "split_file_list")
	})
}

// SplitFileList is a representation of the C type pango_split_file_list.
func SplitFileList(str string) {
	splitFileListFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	splitFileListFunction.Invoke(inArgs[:], nil)

}

var trimStringFunction *gi.Function
var trimStringFunction_Once sync.Once

func trimStringFunction_Set() {
	trimStringFunction_Once.Do(func() {
		trimStringFunction = gi.FunctionInvokerNew("Pango", "trim_string")
	})
}

// TrimString is a representation of the C type pango_trim_string.
func TrimString(str string) string {
	trimStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := trimStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'pango_unichar_direction' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_units_from_double' : parameter 'd' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_units_to_double' : return type 'gdouble' not supported

var versionFunction *gi.Function
var versionFunction_Once sync.Once

func versionFunction_Set() {
	versionFunction_Once.Do(func() {
		versionFunction = gi.FunctionInvokerNew("Pango", "version")
	})
}

// Version is a representation of the C type pango_version.
func Version() int32 {
	versionFunction_Set()

	ret := versionFunction.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

var versionCheckFunction *gi.Function
var versionCheckFunction_Once sync.Once

func versionCheckFunction_Set() {
	versionCheckFunction_Once.Do(func() {
		versionCheckFunction = gi.FunctionInvokerNew("Pango", "version_check")
	})
}

// VersionCheck is a representation of the C type pango_version_check.
func VersionCheck(requiredMajor int32, requiredMinor int32, requiredMicro int32) string {
	versionCheckFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetInt32(requiredMajor)
	inArgs[1].SetInt32(requiredMinor)
	inArgs[2].SetInt32(requiredMicro)

	ret := versionCheckFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var versionStringFunction *gi.Function
var versionStringFunction_Once sync.Once

func versionStringFunction_Set() {
	versionStringFunction_Once.Do(func() {
		versionStringFunction = gi.FunctionInvokerNew("Pango", "version_string")
	})
}

// VersionString is a representation of the C type pango_version_string.
func VersionString() string {
	versionStringFunction_Set()

	ret := versionStringFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

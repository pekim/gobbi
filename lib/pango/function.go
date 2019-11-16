// Code generated - DO NOT EDIT.

package pango

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'pango_attr_background_alpha_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_background_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_fallback_new' : parameter 'enable_fallback' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_attr_family_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_foreground_alpha_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_foreground_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_gravity_hint_new' : parameter 'hint' of type 'GravityHint' not supported

// UNSUPPORTED : C value 'pango_attr_gravity_new' : parameter 'gravity' of type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_attr_letter_spacing_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_rise_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_scale_new' : parameter 'scale_factor' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_attr_stretch_new' : parameter 'stretch' of type 'Stretch' not supported

// UNSUPPORTED : C value 'pango_attr_strikethrough_color_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_strikethrough_new' : parameter 'strikethrough' of type 'gboolean' not supported

// UNSUPPORTED : C value 'pango_attr_style_new' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'pango_attr_type_get_name' : parameter 'type' of type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_type_register' : return type 'AttrType' not supported

// UNSUPPORTED : C value 'pango_attr_underline_color_new' : return type 'Attribute' not supported

// UNSUPPORTED : C value 'pango_attr_underline_new' : parameter 'underline' of type 'Underline' not supported

// UNSUPPORTED : C value 'pango_attr_variant_new' : parameter 'variant' of type 'Variant' not supported

// UNSUPPORTED : C value 'pango_attr_weight_new' : parameter 'weight' of type 'Weight' not supported

// UNSUPPORTED : C value 'pango_bidi_type_for_unichar' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_break' : parameter 'analysis' of type 'Analysis' not supported

var configKeyGetInvoker *gi.Function

// ConfigKeyGet is a representation of the C type pango_config_key_get.
func ConfigKeyGet(key string) string {
	if configKeyGetInvoker == nil {
		configKeyGetInvoker = gi.FunctionInvokerNew("Pango", "config_key_get")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(key)

	ret := configKeyGetInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

var configKeyGetSystemInvoker *gi.Function

// ConfigKeyGetSystem is a representation of the C type pango_config_key_get_system.
func ConfigKeyGetSystem(key string) string {
	if configKeyGetSystemInvoker == nil {
		configKeyGetSystemInvoker = gi.FunctionInvokerNew("Pango", "config_key_get_system")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(key)

	ret := configKeyGetSystemInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'pango_default_break' : parameter 'analysis' of type 'Analysis' not supported

// UNSUPPORTED : C value 'pango_extents_to_pixels' : parameter 'inclusive' of type 'Rectangle' not supported

// UNSUPPORTED : C value 'pango_find_base_dir' : return type 'Direction' not supported

// UNSUPPORTED : C value 'pango_find_map' : parameter 'language' of type 'Language' not supported

// UNSUPPORTED : C value 'pango_find_paragraph_boundary' : parameter 'paragraph_delimiter_index' with direction 'out' not supported

// UNSUPPORTED : C value 'pango_font_description_from_string' : return type 'FontDescription' not supported

var getLibSubdirectoryInvoker *gi.Function

// GetLibSubdirectory is a representation of the C type pango_get_lib_subdirectory.
func GetLibSubdirectory() string {
	if getLibSubdirectoryInvoker == nil {
		getLibSubdirectoryInvoker = gi.FunctionInvokerNew("Pango", "get_lib_subdirectory")
	}

	ret := getLibSubdirectoryInvoker.Invoke(nil)
	return ret.String(false)
}

// UNSUPPORTED : C value 'pango_get_log_attrs' : parameter 'language' of type 'Language' not supported

// UNSUPPORTED : C value 'pango_get_mirror_char' : parameter 'ch' of type 'gunichar' not supported

var getSysconfSubdirectoryInvoker *gi.Function

// GetSysconfSubdirectory is a representation of the C type pango_get_sysconf_subdirectory.
func GetSysconfSubdirectory() string {
	if getSysconfSubdirectoryInvoker == nil {
		getSysconfSubdirectoryInvoker = gi.FunctionInvokerNew("Pango", "get_sysconf_subdirectory")
	}

	ret := getSysconfSubdirectoryInvoker.Invoke(nil)
	return ret.String(false)
}

// UNSUPPORTED : C value 'pango_gravity_get_for_matrix' : parameter 'matrix' of type 'Matrix' not supported

// UNSUPPORTED : C value 'pango_gravity_get_for_script' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_gravity_get_for_script_and_width' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_gravity_to_rotation' : parameter 'gravity' of type 'Gravity' not supported

// UNSUPPORTED : C value 'pango_is_zero_width' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_itemize' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'pango_itemize_with_base_dir' : parameter 'context' of type 'Context' not supported

// UNSUPPORTED : C value 'pango_language_from_string' : return type 'Language' not supported

// UNSUPPORTED : C value 'pango_language_get_default' : return type 'Language' not supported

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

// UNSUPPORTED : C value 'pango_quantize_line_geometry' : parameter 'thickness' with direction 'inout' not supported

// UNSUPPORTED : C value 'pango_read_line' : parameter 'stream' of type 'gpointer' not supported

// UNSUPPORTED : C value 'pango_reorder_items' : parameter 'logical_items' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'pango_scan_int' : parameter 'pos' with direction 'inout' not supported

// UNSUPPORTED : C value 'pango_scan_string' : parameter 'pos' with direction 'inout' not supported

// UNSUPPORTED : C value 'pango_scan_word' : parameter 'pos' with direction 'inout' not supported

// UNSUPPORTED : C value 'pango_script_for_unichar' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_script_get_sample_language' : parameter 'script' of type 'Script' not supported

// UNSUPPORTED : C value 'pango_shape' : parameter 'analysis' of type 'Analysis' not supported

// UNSUPPORTED : C value 'pango_shape_full' : parameter 'analysis' of type 'Analysis' not supported

// UNSUPPORTED : C value 'pango_skip_space' : parameter 'pos' with direction 'inout' not supported

var splitFileListInvoker *gi.Function

// SplitFileList is a representation of the C type pango_split_file_list.
func SplitFileList(str string) {
	if splitFileListInvoker == nil {
		splitFileListInvoker = gi.FunctionInvokerNew("Pango", "split_file_list")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	splitFileListInvoker.Invoke(inArgs[:])
}

var trimStringInvoker *gi.Function

// TrimString is a representation of the C type pango_trim_string.
func TrimString(str string) string {
	if trimStringInvoker == nil {
		trimStringInvoker = gi.FunctionInvokerNew("Pango", "trim_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := trimStringInvoker.Invoke(inArgs[:])
	return ret.String(true)
}

// UNSUPPORTED : C value 'pango_unichar_direction' : parameter 'ch' of type 'gunichar' not supported

// UNSUPPORTED : C value 'pango_units_from_double' : parameter 'd' of type 'gdouble' not supported

// UNSUPPORTED : C value 'pango_units_to_double' : return type 'gdouble' not supported

var versionInvoker *gi.Function

// Version is a representation of the C type pango_version.
func Version() int32 {
	if versionInvoker == nil {
		versionInvoker = gi.FunctionInvokerNew("Pango", "version")
	}

	ret := versionInvoker.Invoke(nil)
	return ret.Int32()
}

var versionCheckInvoker *gi.Function

// VersionCheck is a representation of the C type pango_version_check.
func VersionCheck(requiredMajor int32, requiredMinor int32, requiredMicro int32) string {
	if versionCheckInvoker == nil {
		versionCheckInvoker = gi.FunctionInvokerNew("Pango", "version_check")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetInt32(requiredMajor)
	inArgs[1].SetInt32(requiredMinor)
	inArgs[2].SetInt32(requiredMicro)

	ret := versionCheckInvoker.Invoke(inArgs[:])
	return ret.String(false)
}

var versionStringInvoker *gi.Function

// VersionString is a representation of the C type pango_version_string.
func VersionString() string {
	if versionStringInvoker == nil {
		versionStringInvoker = gi.FunctionInvokerNew("Pango", "version_string")
	}

	ret := versionStringInvoker.Invoke(nil)
	return ret.String(false)
}

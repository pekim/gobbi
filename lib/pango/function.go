// Code generated - DO NOT EDIT.

package pango

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'pango_attr_background_alpha_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_background_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_fallback_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_family_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_foreground_alpha_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_foreground_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_gravity_hint_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_gravity_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_letter_spacing_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_rise_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_scale_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_stretch_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_strikethrough_color_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_strikethrough_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_style_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_type_get_name' : has parameters

// UNSUPPORTED : C value 'pango_attr_type_register' : has parameters

// UNSUPPORTED : C value 'pango_attr_underline_color_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_underline_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_variant_new' : has parameters

// UNSUPPORTED : C value 'pango_attr_weight_new' : has parameters

// UNSUPPORTED : C value 'pango_bidi_type_for_unichar' : has parameters

// UNSUPPORTED : C value 'pango_break' : has parameters

// UNSUPPORTED : C value 'pango_config_key_get' : has parameters

// UNSUPPORTED : C value 'pango_config_key_get_system' : has parameters

// UNSUPPORTED : C value 'pango_default_break' : has parameters

// UNSUPPORTED : C value 'pango_extents_to_pixels' : has parameters

// UNSUPPORTED : C value 'pango_find_base_dir' : has parameters

// UNSUPPORTED : C value 'pango_find_map' : has parameters

// UNSUPPORTED : C value 'pango_find_paragraph_boundary' : has parameters

// UNSUPPORTED : C value 'pango_font_description_from_string' : has parameters

// UNSUPPORTED : C value 'pango_get_lib_subdirectory' : return type not supported

// UNSUPPORTED : C value 'pango_get_log_attrs' : has parameters

// UNSUPPORTED : C value 'pango_get_mirror_char' : has parameters

// UNSUPPORTED : C value 'pango_get_sysconf_subdirectory' : return type not supported

// UNSUPPORTED : C value 'pango_gravity_get_for_matrix' : has parameters

// UNSUPPORTED : C value 'pango_gravity_get_for_script' : has parameters

// UNSUPPORTED : C value 'pango_gravity_get_for_script_and_width' : has parameters

// UNSUPPORTED : C value 'pango_gravity_to_rotation' : has parameters

// UNSUPPORTED : C value 'pango_is_zero_width' : has parameters

// UNSUPPORTED : C value 'pango_itemize' : has parameters

// UNSUPPORTED : C value 'pango_itemize_with_base_dir' : has parameters

// UNSUPPORTED : C value 'pango_language_from_string' : has parameters

// UNSUPPORTED : C value 'pango_language_get_default' : return type not supported

// UNSUPPORTED : C value 'pango_log2vis_get_embedding_levels' : has parameters

// UNSUPPORTED : C value 'pango_lookup_aliases' : has parameters

// UNSUPPORTED : C value 'pango_markup_parser_finish' : has parameters

// UNSUPPORTED : C value 'pango_markup_parser_new' : has parameters

// UNSUPPORTED : C value 'pango_module_register' : has parameters

// UNSUPPORTED : C value 'pango_parse_enum' : has parameters

// UNSUPPORTED : C value 'pango_parse_markup' : has parameters

// UNSUPPORTED : C value 'pango_parse_stretch' : has parameters

// UNSUPPORTED : C value 'pango_parse_style' : has parameters

// UNSUPPORTED : C value 'pango_parse_variant' : has parameters

// UNSUPPORTED : C value 'pango_parse_weight' : has parameters

// UNSUPPORTED : C value 'pango_quantize_line_geometry' : has parameters

// UNSUPPORTED : C value 'pango_read_line' : has parameters

// UNSUPPORTED : C value 'pango_reorder_items' : has parameters

// UNSUPPORTED : C value 'pango_scan_int' : has parameters

// UNSUPPORTED : C value 'pango_scan_string' : has parameters

// UNSUPPORTED : C value 'pango_scan_word' : has parameters

// UNSUPPORTED : C value 'pango_script_for_unichar' : has parameters

// UNSUPPORTED : C value 'pango_script_get_sample_language' : has parameters

// UNSUPPORTED : C value 'pango_shape' : has parameters

// UNSUPPORTED : C value 'pango_shape_full' : has parameters

// UNSUPPORTED : C value 'pango_skip_space' : has parameters

// UNSUPPORTED : C value 'pango_split_file_list' : has parameters

// UNSUPPORTED : C value 'pango_trim_string' : has parameters

// UNSUPPORTED : C value 'pango_unichar_direction' : has parameters

// UNSUPPORTED : C value 'pango_units_from_double' : has parameters

// UNSUPPORTED : C value 'pango_units_to_double' : has parameters

var versionInvoker *gi.Function

// Version is a representation of the C type pango_version.
func Version() int32 {
	if versionInvoker == nil {
		versionInvoker = gi.FunctionInvokerNew("Pango", "version")
	}

	versionInvoker.Invoke()
	return ret.Int32()
}

// UNSUPPORTED : C value 'pango_version_check' : has parameters

// UNSUPPORTED : C value 'pango_version_string' : return type not supported

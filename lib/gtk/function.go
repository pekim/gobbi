// Code generated - DO NOT EDIT.

package gtk

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'gtk_accel_groups_activate' : has parameters

// UNSUPPORTED : C value 'gtk_accel_groups_from_object' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_get_default_mod_mask' : return type not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_label' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_get_label_with_keycode' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_name' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_name_with_keycode' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_parse' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_parse_with_keycode' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_set_default_mod_mask' : has parameters

// UNSUPPORTED : C value 'gtk_accelerator_valid' : has parameters

// UNSUPPORTED : C value 'gtk_alternative_dialog_button_order' : has parameters

// UNSUPPORTED : C value 'gtk_binding_entry_add_signal_from_string' : has parameters

// UNSUPPORTED : C value 'gtk_binding_entry_add_signall' : has parameters

// UNSUPPORTED : C value 'gtk_binding_entry_remove' : has parameters

// UNSUPPORTED : C value 'gtk_binding_entry_skip' : has parameters

// UNSUPPORTED : C value 'gtk_binding_set_by_class' : has parameters

// UNSUPPORTED : C value 'gtk_binding_set_find' : has parameters

// UNSUPPORTED : C value 'gtk_binding_set_new' : has parameters

// UNSUPPORTED : C value 'gtk_bindings_activate' : has parameters

// UNSUPPORTED : C value 'gtk_bindings_activate_event' : has parameters

// UNSUPPORTED : C value 'gtk_builder_error_quark' : return type not supported

// UNSUPPORTED : C value 'gtk_cairo_should_draw_window' : has parameters

// UNSUPPORTED : C value 'gtk_cairo_transform_to_window' : has parameters

// UNSUPPORTED : C value 'gtk_check_version' : has parameters

// UNSUPPORTED : C value 'gtk_css_provider_error_quark' : return type not supported

// UNSUPPORTED : C value 'gtk_device_grab_add' : has parameters

// UNSUPPORTED : C value 'gtk_device_grab_remove' : has parameters

// UNSUPPORTED : C value 'gtk_disable_setlocale' : return type not supported

// UNSUPPORTED : C value 'gtk_distribute_natural_allocation' : has parameters

// UNSUPPORTED : C value 'gtk_drag_cancel' : has parameters

// UNSUPPORTED : C value 'gtk_drag_finish' : has parameters

// UNSUPPORTED : C value 'gtk_drag_get_source_widget' : has parameters

// UNSUPPORTED : C value 'gtk_drag_set_icon_default' : has parameters

// UNSUPPORTED : C value 'gtk_drag_set_icon_gicon' : has parameters

// UNSUPPORTED : C value 'gtk_drag_set_icon_name' : has parameters

// UNSUPPORTED : C value 'gtk_drag_set_icon_pixbuf' : has parameters

// UNSUPPORTED : C value 'gtk_drag_set_icon_stock' : has parameters

// UNSUPPORTED : C value 'gtk_drag_set_icon_surface' : has parameters

// UNSUPPORTED : C value 'gtk_drag_set_icon_widget' : has parameters

// UNSUPPORTED : C value 'gtk_draw_insertion_cursor' : has parameters

// UNSUPPORTED : C value 'gtk_events_pending' : return type not supported

// UNSUPPORTED : C value 'gtk_false' : return type not supported

// UNSUPPORTED : C value 'gtk_file_chooser_error_quark' : return type not supported

var getBinaryAgeInvoker *gi.Function

// GetBinaryAge is a representation of the C type gtk_get_binary_age.
func GetBinaryAge() uint32 {
	if getBinaryAgeInvoker == nil {
		getBinaryAgeInvoker = gi.FunctionInvokerNew("Gtk", "get_binary_age")
	}

	ret := getBinaryAgeInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'gtk_get_current_event' : return type not supported

// UNSUPPORTED : C value 'gtk_get_current_event_device' : return type not supported

// UNSUPPORTED : C value 'gtk_get_current_event_state' : has parameters

var getCurrentEventTimeInvoker *gi.Function

// GetCurrentEventTime is a representation of the C type gtk_get_current_event_time.
func GetCurrentEventTime() uint32 {
	if getCurrentEventTimeInvoker == nil {
		getCurrentEventTimeInvoker = gi.FunctionInvokerNew("Gtk", "get_current_event_time")
	}

	ret := getCurrentEventTimeInvoker.Invoke()
	return ret.Uint32()
}

var getDebugFlagsInvoker *gi.Function

// GetDebugFlags is a representation of the C type gtk_get_debug_flags.
func GetDebugFlags() uint32 {
	if getDebugFlagsInvoker == nil {
		getDebugFlagsInvoker = gi.FunctionInvokerNew("Gtk", "get_debug_flags")
	}

	ret := getDebugFlagsInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'gtk_get_default_language' : return type not supported

// UNSUPPORTED : C value 'gtk_get_event_widget' : has parameters

var getInterfaceAgeInvoker *gi.Function

// GetInterfaceAge is a representation of the C type gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	if getInterfaceAgeInvoker == nil {
		getInterfaceAgeInvoker = gi.FunctionInvokerNew("Gtk", "get_interface_age")
	}

	ret := getInterfaceAgeInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'gtk_get_locale_direction' : return type not supported

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type gtk_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Gtk", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type gtk_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Gtk", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	return ret.Uint32()
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type gtk_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Gtk", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'gtk_get_option_group' : has parameters

// UNSUPPORTED : C value 'gtk_grab_get_current' : return type not supported

// UNSUPPORTED : C value 'gtk_icon_size_from_name' : has parameters

// UNSUPPORTED : C value 'gtk_icon_size_get_name' : has parameters

// UNSUPPORTED : C value 'gtk_icon_size_lookup' : has parameters

// UNSUPPORTED : C value 'gtk_icon_size_lookup_for_settings' : has parameters

// UNSUPPORTED : C value 'gtk_icon_size_register' : has parameters

// UNSUPPORTED : C value 'gtk_icon_size_register_alias' : has parameters

// UNSUPPORTED : C value 'gtk_icon_theme_error_quark' : return type not supported

// UNSUPPORTED : C value 'gtk_init' : has parameters

// UNSUPPORTED : C value 'gtk_init_check' : has parameters

// UNSUPPORTED : C value 'gtk_init_with_args' : has parameters

// UNSUPPORTED : C value 'gtk_key_snooper_install' : has parameters

// UNSUPPORTED : C value 'gtk_key_snooper_remove' : has parameters

// UNSUPPORTED : C value 'gtk_main' : return type not supported

// UNSUPPORTED : C value 'gtk_main_do_event' : has parameters

// UNSUPPORTED : C value 'gtk_main_iteration' : return type not supported

// UNSUPPORTED : C value 'gtk_main_iteration_do' : has parameters

var mainLevelInvoker *gi.Function

// MainLevel is a representation of the C type gtk_main_level.
func MainLevel() uint32 {
	if mainLevelInvoker == nil {
		mainLevelInvoker = gi.FunctionInvokerNew("Gtk", "main_level")
	}

	ret := mainLevelInvoker.Invoke()
	return ret.Uint32()
}

// UNSUPPORTED : C value 'gtk_main_quit' : return type not supported

// UNSUPPORTED : C value 'gtk_paint_arrow' : has parameters

// UNSUPPORTED : C value 'gtk_paint_box' : has parameters

// UNSUPPORTED : C value 'gtk_paint_box_gap' : has parameters

// UNSUPPORTED : C value 'gtk_paint_check' : has parameters

// UNSUPPORTED : C value 'gtk_paint_diamond' : has parameters

// UNSUPPORTED : C value 'gtk_paint_expander' : has parameters

// UNSUPPORTED : C value 'gtk_paint_extension' : has parameters

// UNSUPPORTED : C value 'gtk_paint_flat_box' : has parameters

// UNSUPPORTED : C value 'gtk_paint_focus' : has parameters

// UNSUPPORTED : C value 'gtk_paint_handle' : has parameters

// UNSUPPORTED : C value 'gtk_paint_hline' : has parameters

// UNSUPPORTED : C value 'gtk_paint_layout' : has parameters

// UNSUPPORTED : C value 'gtk_paint_option' : has parameters

// UNSUPPORTED : C value 'gtk_paint_resize_grip' : has parameters

// UNSUPPORTED : C value 'gtk_paint_shadow' : has parameters

// UNSUPPORTED : C value 'gtk_paint_shadow_gap' : has parameters

// UNSUPPORTED : C value 'gtk_paint_slider' : has parameters

// UNSUPPORTED : C value 'gtk_paint_spinner' : has parameters

// UNSUPPORTED : C value 'gtk_paint_tab' : has parameters

// UNSUPPORTED : C value 'gtk_paint_vline' : has parameters

// UNSUPPORTED : C value 'gtk_paper_size_get_default' : return type not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_paper_sizes' : has parameters

// UNSUPPORTED : C value 'gtk_parse_args' : has parameters

// UNSUPPORTED : C value 'gtk_print_error_quark' : return type not supported

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog' : has parameters

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog_async' : has parameters

// UNSUPPORTED : C value 'gtk_propagate_event' : has parameters

// UNSUPPORTED : C value 'gtk_rc_add_default_file' : has parameters

// UNSUPPORTED : C value 'gtk_rc_find_module_in_path' : has parameters

// UNSUPPORTED : C value 'gtk_rc_find_pixmap_in_path' : has parameters

var rcGetDefaultFilesInvoker *gi.Function

// RcGetDefaultFiles is a representation of the C type gtk_rc_get_default_files.
func RcGetDefaultFiles() {
	if rcGetDefaultFilesInvoker == nil {
		rcGetDefaultFilesInvoker = gi.FunctionInvokerNew("Gtk", "rc_get_default_files")
	}

	rcGetDefaultFilesInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_rc_get_im_module_file' : return type not supported

// UNSUPPORTED : C value 'gtk_rc_get_im_module_path' : return type not supported

// UNSUPPORTED : C value 'gtk_rc_get_module_dir' : return type not supported

// UNSUPPORTED : C value 'gtk_rc_get_style' : has parameters

// UNSUPPORTED : C value 'gtk_rc_get_style_by_paths' : has parameters

// UNSUPPORTED : C value 'gtk_rc_get_theme_dir' : return type not supported

// UNSUPPORTED : C value 'gtk_rc_parse' : has parameters

// UNSUPPORTED : C value 'gtk_rc_parse_color' : has parameters

// UNSUPPORTED : C value 'gtk_rc_parse_color_full' : has parameters

// UNSUPPORTED : C value 'gtk_rc_parse_priority' : has parameters

// UNSUPPORTED : C value 'gtk_rc_parse_state' : has parameters

// UNSUPPORTED : C value 'gtk_rc_parse_string' : has parameters

// UNSUPPORTED : C value 'gtk_rc_property_parse_border' : has parameters

// UNSUPPORTED : C value 'gtk_rc_property_parse_color' : has parameters

// UNSUPPORTED : C value 'gtk_rc_property_parse_enum' : has parameters

// UNSUPPORTED : C value 'gtk_rc_property_parse_flags' : has parameters

// UNSUPPORTED : C value 'gtk_rc_property_parse_requisition' : has parameters

// UNSUPPORTED : C value 'gtk_rc_reparse_all' : return type not supported

// UNSUPPORTED : C value 'gtk_rc_reparse_all_for_settings' : has parameters

// UNSUPPORTED : C value 'gtk_rc_reset_styles' : has parameters

// UNSUPPORTED : C value 'gtk_rc_scanner_new' : return type not supported

// UNSUPPORTED : C value 'gtk_rc_set_default_files' : has parameters

// UNSUPPORTED : C value 'gtk_recent_chooser_error_quark' : return type not supported

// UNSUPPORTED : C value 'gtk_recent_manager_error_quark' : return type not supported

// UNSUPPORTED : C value 'gtk_render_activity' : has parameters

// UNSUPPORTED : C value 'gtk_render_arrow' : has parameters

// UNSUPPORTED : C value 'gtk_render_background' : has parameters

// UNSUPPORTED : C value 'gtk_render_background_get_clip' : has parameters

// UNSUPPORTED : C value 'gtk_render_check' : has parameters

// UNSUPPORTED : C value 'gtk_render_expander' : has parameters

// UNSUPPORTED : C value 'gtk_render_extension' : has parameters

// UNSUPPORTED : C value 'gtk_render_focus' : has parameters

// UNSUPPORTED : C value 'gtk_render_frame' : has parameters

// UNSUPPORTED : C value 'gtk_render_frame_gap' : has parameters

// UNSUPPORTED : C value 'gtk_render_handle' : has parameters

// UNSUPPORTED : C value 'gtk_render_icon' : has parameters

// UNSUPPORTED : C value 'gtk_render_icon_pixbuf' : has parameters

// UNSUPPORTED : C value 'gtk_render_icon_surface' : has parameters

// UNSUPPORTED : C value 'gtk_render_insertion_cursor' : has parameters

// UNSUPPORTED : C value 'gtk_render_layout' : has parameters

// UNSUPPORTED : C value 'gtk_render_line' : has parameters

// UNSUPPORTED : C value 'gtk_render_option' : has parameters

// UNSUPPORTED : C value 'gtk_render_slider' : has parameters

// UNSUPPORTED : C value 'gtk_rgb_to_hsv' : has parameters

// UNSUPPORTED : C value 'gtk_selection_add_target' : has parameters

// UNSUPPORTED : C value 'gtk_selection_add_targets' : has parameters

// UNSUPPORTED : C value 'gtk_selection_clear_targets' : has parameters

// UNSUPPORTED : C value 'gtk_selection_convert' : has parameters

// UNSUPPORTED : C value 'gtk_selection_owner_set' : has parameters

// UNSUPPORTED : C value 'gtk_selection_owner_set_for_display' : has parameters

// UNSUPPORTED : C value 'gtk_selection_remove_all' : has parameters

// UNSUPPORTED : C value 'gtk_set_debug_flags' : has parameters

// UNSUPPORTED : C value 'gtk_show_about_dialog' : has parameters

// UNSUPPORTED : C value 'gtk_show_uri' : has parameters

// UNSUPPORTED : C value 'gtk_show_uri_on_window' : has parameters

// UNSUPPORTED : C value 'gtk_stock_add' : has parameters

// UNSUPPORTED : C value 'gtk_stock_add_static' : has parameters

// UNSUPPORTED : C value 'gtk_stock_list_ids' : return type not supported

// UNSUPPORTED : C value 'gtk_stock_lookup' : has parameters

// UNSUPPORTED : C value 'gtk_stock_set_translate_func' : has parameters

// UNSUPPORTED : C value 'gtk_target_table_free' : has parameters

// UNSUPPORTED : C value 'gtk_target_table_new_from_list' : has parameters

// UNSUPPORTED : C value 'gtk_targets_include_image' : has parameters

// UNSUPPORTED : C value 'gtk_targets_include_rich_text' : has parameters

// UNSUPPORTED : C value 'gtk_targets_include_text' : has parameters

// UNSUPPORTED : C value 'gtk_targets_include_uri' : has parameters

// UNSUPPORTED : C value 'gtk_test_create_simple_window' : has parameters

// UNSUPPORTED : C value 'gtk_test_create_widget' : has parameters

// UNSUPPORTED : C value 'gtk_test_display_button_window' : has parameters

// UNSUPPORTED : C value 'gtk_test_find_label' : has parameters

// UNSUPPORTED : C value 'gtk_test_find_sibling' : has parameters

// UNSUPPORTED : C value 'gtk_test_find_widget' : has parameters

// UNSUPPORTED : C value 'gtk_test_init' : has parameters

// UNSUPPORTED : C value 'gtk_test_list_all_types' : has parameters

// UNSUPPORTED : C value 'gtk_test_register_all_types' : return type not supported

// UNSUPPORTED : C value 'gtk_test_slider_get_value' : has parameters

// UNSUPPORTED : C value 'gtk_test_slider_set_perc' : has parameters

// UNSUPPORTED : C value 'gtk_test_spin_button_click' : has parameters

// UNSUPPORTED : C value 'gtk_test_text_get' : has parameters

// UNSUPPORTED : C value 'gtk_test_text_set' : has parameters

// UNSUPPORTED : C value 'gtk_test_widget_click' : has parameters

// UNSUPPORTED : C value 'gtk_test_widget_send_key' : has parameters

// UNSUPPORTED : C value 'gtk_test_widget_wait_for_draw' : has parameters

// UNSUPPORTED : C value 'gtk_tree_get_row_drag_data' : has parameters

// UNSUPPORTED : C value 'gtk_tree_row_reference_deleted' : has parameters

// UNSUPPORTED : C value 'gtk_tree_row_reference_inserted' : has parameters

// UNSUPPORTED : C value 'gtk_tree_row_reference_reordered' : has parameters

// UNSUPPORTED : C value 'gtk_tree_set_row_drag_data' : has parameters

// UNSUPPORTED : C value 'gtk_true' : return type not supported

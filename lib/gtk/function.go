// Code generated - DO NOT EDIT.

package gtk

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'gtk_accel_groups_activate' : non trivial function

// UNSUPPORTED : C value 'gtk_accel_groups_from_object' : non trivial function

var acceleratorGetDefaultModMaskInvoker *gi.Function

// AcceleratorGetDefaultModMask is a representation of the C type gtk_accelerator_get_default_mod_mask.
func AcceleratorGetDefaultModMask() {
	if acceleratorGetDefaultModMaskInvoker == nil {
		acceleratorGetDefaultModMaskInvoker = gi.FunctionInvokerNew("Gtk", "accelerator_get_default_mod_mask")
	}

	acceleratorGetDefaultModMaskInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_accelerator_get_label' : non trivial function

// UNSUPPORTED : C value 'gtk_accelerator_get_label_with_keycode' : non trivial function

// UNSUPPORTED : C value 'gtk_accelerator_name' : non trivial function

// UNSUPPORTED : C value 'gtk_accelerator_name_with_keycode' : non trivial function

// UNSUPPORTED : C value 'gtk_accelerator_parse' : non trivial function

// UNSUPPORTED : C value 'gtk_accelerator_parse_with_keycode' : non trivial function

// UNSUPPORTED : C value 'gtk_accelerator_set_default_mod_mask' : non trivial function

// UNSUPPORTED : C value 'gtk_accelerator_valid' : non trivial function

// UNSUPPORTED : C value 'gtk_alternative_dialog_button_order' : non trivial function

// UNSUPPORTED : C value 'gtk_binding_entry_add_signal_from_string' : non trivial function

// UNSUPPORTED : C value 'gtk_binding_entry_add_signall' : non trivial function

// UNSUPPORTED : C value 'gtk_binding_entry_remove' : non trivial function

// UNSUPPORTED : C value 'gtk_binding_entry_skip' : non trivial function

// UNSUPPORTED : C value 'gtk_binding_set_by_class' : non trivial function

// UNSUPPORTED : C value 'gtk_binding_set_find' : non trivial function

// UNSUPPORTED : C value 'gtk_binding_set_new' : non trivial function

// UNSUPPORTED : C value 'gtk_bindings_activate' : non trivial function

// UNSUPPORTED : C value 'gtk_bindings_activate_event' : non trivial function

var builderErrorQuarkInvoker *gi.Function

// BuilderErrorQuark is a representation of the C type gtk_builder_error_quark.
func BuilderErrorQuark() {
	if builderErrorQuarkInvoker == nil {
		builderErrorQuarkInvoker = gi.FunctionInvokerNew("Gtk", "builder_error_quark")
	}

	builderErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_cairo_should_draw_window' : non trivial function

// UNSUPPORTED : C value 'gtk_cairo_transform_to_window' : non trivial function

// UNSUPPORTED : C value 'gtk_check_version' : non trivial function

var cssProviderErrorQuarkInvoker *gi.Function

// CssProviderErrorQuark is a representation of the C type gtk_css_provider_error_quark.
func CssProviderErrorQuark() {
	if cssProviderErrorQuarkInvoker == nil {
		cssProviderErrorQuarkInvoker = gi.FunctionInvokerNew("Gtk", "css_provider_error_quark")
	}

	cssProviderErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_device_grab_add' : non trivial function

// UNSUPPORTED : C value 'gtk_device_grab_remove' : non trivial function

var disableSetlocaleInvoker *gi.Function

// DisableSetlocale is a representation of the C type gtk_disable_setlocale.
func DisableSetlocale() {
	if disableSetlocaleInvoker == nil {
		disableSetlocaleInvoker = gi.FunctionInvokerNew("Gtk", "disable_setlocale")
	}

	disableSetlocaleInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_distribute_natural_allocation' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_cancel' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_finish' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_get_source_widget' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_set_icon_default' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_set_icon_gicon' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_set_icon_name' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_set_icon_pixbuf' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_set_icon_stock' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_set_icon_surface' : non trivial function

// UNSUPPORTED : C value 'gtk_drag_set_icon_widget' : non trivial function

// UNSUPPORTED : C value 'gtk_draw_insertion_cursor' : non trivial function

var eventsPendingInvoker *gi.Function

// EventsPending is a representation of the C type gtk_events_pending.
func EventsPending() {
	if eventsPendingInvoker == nil {
		eventsPendingInvoker = gi.FunctionInvokerNew("Gtk", "events_pending")
	}

	eventsPendingInvoker.Invoke()
}

var falseInvoker *gi.Function

// False is a representation of the C type gtk_false.
func False() {
	if falseInvoker == nil {
		falseInvoker = gi.FunctionInvokerNew("Gtk", "false")
	}

	falseInvoker.Invoke()
}

var fileChooserErrorQuarkInvoker *gi.Function

// FileChooserErrorQuark is a representation of the C type gtk_file_chooser_error_quark.
func FileChooserErrorQuark() {
	if fileChooserErrorQuarkInvoker == nil {
		fileChooserErrorQuarkInvoker = gi.FunctionInvokerNew("Gtk", "file_chooser_error_quark")
	}

	fileChooserErrorQuarkInvoker.Invoke()
}

var getBinaryAgeInvoker *gi.Function

// GetBinaryAge is a representation of the C type gtk_get_binary_age.
func GetBinaryAge() uint32 {
	if getBinaryAgeInvoker == nil {
		getBinaryAgeInvoker = gi.FunctionInvokerNew("Gtk", "get_binary_age")
	}

	ret := getBinaryAgeInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getCurrentEventInvoker *gi.Function

// GetCurrentEvent is a representation of the C type gtk_get_current_event.
func GetCurrentEvent() {
	if getCurrentEventInvoker == nil {
		getCurrentEventInvoker = gi.FunctionInvokerNew("Gtk", "get_current_event")
	}

	getCurrentEventInvoker.Invoke()
}

var getCurrentEventDeviceInvoker *gi.Function

// GetCurrentEventDevice is a representation of the C type gtk_get_current_event_device.
func GetCurrentEventDevice() {
	if getCurrentEventDeviceInvoker == nil {
		getCurrentEventDeviceInvoker = gi.FunctionInvokerNew("Gtk", "get_current_event_device")
	}

	getCurrentEventDeviceInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_get_current_event_state' : non trivial function

var getCurrentEventTimeInvoker *gi.Function

// GetCurrentEventTime is a representation of the C type gtk_get_current_event_time.
func GetCurrentEventTime() {
	if getCurrentEventTimeInvoker == nil {
		getCurrentEventTimeInvoker = gi.FunctionInvokerNew("Gtk", "get_current_event_time")
	}

	getCurrentEventTimeInvoker.Invoke()
}

var getDebugFlagsInvoker *gi.Function

// GetDebugFlags is a representation of the C type gtk_get_debug_flags.
func GetDebugFlags() uint32 {
	if getDebugFlagsInvoker == nil {
		getDebugFlagsInvoker = gi.FunctionInvokerNew("Gtk", "get_debug_flags")
	}

	ret := getDebugFlagsInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getDefaultLanguageInvoker *gi.Function

// GetDefaultLanguage is a representation of the C type gtk_get_default_language.
func GetDefaultLanguage() {
	if getDefaultLanguageInvoker == nil {
		getDefaultLanguageInvoker = gi.FunctionInvokerNew("Gtk", "get_default_language")
	}

	getDefaultLanguageInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_get_event_widget' : non trivial function

var getInterfaceAgeInvoker *gi.Function

// GetInterfaceAge is a representation of the C type gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	if getInterfaceAgeInvoker == nil {
		getInterfaceAgeInvoker = gi.FunctionInvokerNew("Gtk", "get_interface_age")
	}

	ret := getInterfaceAgeInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getLocaleDirectionInvoker *gi.Function

// GetLocaleDirection is a representation of the C type gtk_get_locale_direction.
func GetLocaleDirection() {
	if getLocaleDirectionInvoker == nil {
		getLocaleDirectionInvoker = gi.FunctionInvokerNew("Gtk", "get_locale_direction")
	}

	getLocaleDirectionInvoker.Invoke()
}

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type gtk_get_major_version.
func GetMajorVersion() uint32 {
	if getMajorVersionInvoker == nil {
		getMajorVersionInvoker = gi.FunctionInvokerNew("Gtk", "get_major_version")
	}

	ret := getMajorVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type gtk_get_micro_version.
func GetMicroVersion() uint32 {
	if getMicroVersionInvoker == nil {
		getMicroVersionInvoker = gi.FunctionInvokerNew("Gtk", "get_micro_version")
	}

	ret := getMicroVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type gtk_get_minor_version.
func GetMinorVersion() uint32 {
	if getMinorVersionInvoker == nil {
		getMinorVersionInvoker = gi.FunctionInvokerNew("Gtk", "get_minor_version")
	}

	ret := getMinorVersionInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

// UNSUPPORTED : C value 'gtk_get_option_group' : non trivial function

var grabGetCurrentInvoker *gi.Function

// GrabGetCurrent is a representation of the C type gtk_grab_get_current.
func GrabGetCurrent() {
	if grabGetCurrentInvoker == nil {
		grabGetCurrentInvoker = gi.FunctionInvokerNew("Gtk", "grab_get_current")
	}

	grabGetCurrentInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_icon_size_from_name' : non trivial function

// UNSUPPORTED : C value 'gtk_icon_size_get_name' : non trivial function

// UNSUPPORTED : C value 'gtk_icon_size_lookup' : non trivial function

// UNSUPPORTED : C value 'gtk_icon_size_lookup_for_settings' : non trivial function

// UNSUPPORTED : C value 'gtk_icon_size_register' : non trivial function

// UNSUPPORTED : C value 'gtk_icon_size_register_alias' : non trivial function

var iconThemeErrorQuarkInvoker *gi.Function

// IconThemeErrorQuark is a representation of the C type gtk_icon_theme_error_quark.
func IconThemeErrorQuark() {
	if iconThemeErrorQuarkInvoker == nil {
		iconThemeErrorQuarkInvoker = gi.FunctionInvokerNew("Gtk", "icon_theme_error_quark")
	}

	iconThemeErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_init' : non trivial function

// UNSUPPORTED : C value 'gtk_init_check' : non trivial function

// UNSUPPORTED : C value 'gtk_init_with_args' : non trivial function

// UNSUPPORTED : C value 'gtk_key_snooper_install' : non trivial function

// UNSUPPORTED : C value 'gtk_key_snooper_remove' : non trivial function

var mainInvoker *gi.Function

// Main is a representation of the C type gtk_main.
func Main() {
	if mainInvoker == nil {
		mainInvoker = gi.FunctionInvokerNew("Gtk", "main")
	}

	mainInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_main_do_event' : non trivial function

var mainIterationInvoker *gi.Function

// MainIteration is a representation of the C type gtk_main_iteration.
func MainIteration() {
	if mainIterationInvoker == nil {
		mainIterationInvoker = gi.FunctionInvokerNew("Gtk", "main_iteration")
	}

	mainIterationInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_main_iteration_do' : non trivial function

var mainLevelInvoker *gi.Function

// MainLevel is a representation of the C type gtk_main_level.
func MainLevel() uint32 {
	if mainLevelInvoker == nil {
		mainLevelInvoker = gi.FunctionInvokerNew("Gtk", "main_level")
	}

	ret := mainLevelInvoker.Invoke()
	retValue := ret.Uint32()
	return retValue
}

var mainQuitInvoker *gi.Function

// MainQuit is a representation of the C type gtk_main_quit.
func MainQuit() {
	if mainQuitInvoker == nil {
		mainQuitInvoker = gi.FunctionInvokerNew("Gtk", "main_quit")
	}

	mainQuitInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_paint_arrow' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_box' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_box_gap' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_check' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_diamond' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_expander' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_extension' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_flat_box' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_focus' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_handle' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_hline' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_layout' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_option' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_resize_grip' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_shadow' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_shadow_gap' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_slider' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_spinner' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_tab' : non trivial function

// UNSUPPORTED : C value 'gtk_paint_vline' : non trivial function

var paperSizeGetDefaultInvoker *gi.Function

// PaperSizeGetDefault is a representation of the C type gtk_paper_size_get_default.
func PaperSizeGetDefault() {
	if paperSizeGetDefaultInvoker == nil {
		paperSizeGetDefaultInvoker = gi.FunctionInvokerNew("Gtk", "paper_size_get_default")
	}

	paperSizeGetDefaultInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_paper_size_get_paper_sizes' : non trivial function

// UNSUPPORTED : C value 'gtk_parse_args' : non trivial function

var printErrorQuarkInvoker *gi.Function

// PrintErrorQuark is a representation of the C type gtk_print_error_quark.
func PrintErrorQuark() {
	if printErrorQuarkInvoker == nil {
		printErrorQuarkInvoker = gi.FunctionInvokerNew("Gtk", "print_error_quark")
	}

	printErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog' : non trivial function

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog_async' : non trivial function

// UNSUPPORTED : C value 'gtk_propagate_event' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_add_default_file' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_find_module_in_path' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_find_pixmap_in_path' : non trivial function

var rcGetDefaultFilesInvoker *gi.Function

// RcGetDefaultFiles is a representation of the C type gtk_rc_get_default_files.
func RcGetDefaultFiles() {
	if rcGetDefaultFilesInvoker == nil {
		rcGetDefaultFilesInvoker = gi.FunctionInvokerNew("Gtk", "rc_get_default_files")
	}

	rcGetDefaultFilesInvoker.Invoke()
}

var rcGetImModuleFileInvoker *gi.Function

// RcGetImModuleFile is a representation of the C type gtk_rc_get_im_module_file.
func RcGetImModuleFile() {
	if rcGetImModuleFileInvoker == nil {
		rcGetImModuleFileInvoker = gi.FunctionInvokerNew("Gtk", "rc_get_im_module_file")
	}

	rcGetImModuleFileInvoker.Invoke()
}

var rcGetImModulePathInvoker *gi.Function

// RcGetImModulePath is a representation of the C type gtk_rc_get_im_module_path.
func RcGetImModulePath() {
	if rcGetImModulePathInvoker == nil {
		rcGetImModulePathInvoker = gi.FunctionInvokerNew("Gtk", "rc_get_im_module_path")
	}

	rcGetImModulePathInvoker.Invoke()
}

var rcGetModuleDirInvoker *gi.Function

// RcGetModuleDir is a representation of the C type gtk_rc_get_module_dir.
func RcGetModuleDir() {
	if rcGetModuleDirInvoker == nil {
		rcGetModuleDirInvoker = gi.FunctionInvokerNew("Gtk", "rc_get_module_dir")
	}

	rcGetModuleDirInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_rc_get_style' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_get_style_by_paths' : non trivial function

var rcGetThemeDirInvoker *gi.Function

// RcGetThemeDir is a representation of the C type gtk_rc_get_theme_dir.
func RcGetThemeDir() {
	if rcGetThemeDirInvoker == nil {
		rcGetThemeDirInvoker = gi.FunctionInvokerNew("Gtk", "rc_get_theme_dir")
	}

	rcGetThemeDirInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_rc_parse' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_parse_color' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_parse_color_full' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_parse_priority' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_parse_state' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_parse_string' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_property_parse_border' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_property_parse_color' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_property_parse_enum' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_property_parse_flags' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_property_parse_requisition' : non trivial function

var rcReparseAllInvoker *gi.Function

// RcReparseAll is a representation of the C type gtk_rc_reparse_all.
func RcReparseAll() {
	if rcReparseAllInvoker == nil {
		rcReparseAllInvoker = gi.FunctionInvokerNew("Gtk", "rc_reparse_all")
	}

	rcReparseAllInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_rc_reparse_all_for_settings' : non trivial function

// UNSUPPORTED : C value 'gtk_rc_reset_styles' : non trivial function

var rcScannerNewInvoker *gi.Function

// RcScannerNew is a representation of the C type gtk_rc_scanner_new.
func RcScannerNew() {
	if rcScannerNewInvoker == nil {
		rcScannerNewInvoker = gi.FunctionInvokerNew("Gtk", "rc_scanner_new")
	}

	rcScannerNewInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_rc_set_default_files' : non trivial function

var recentChooserErrorQuarkInvoker *gi.Function

// RecentChooserErrorQuark is a representation of the C type gtk_recent_chooser_error_quark.
func RecentChooserErrorQuark() {
	if recentChooserErrorQuarkInvoker == nil {
		recentChooserErrorQuarkInvoker = gi.FunctionInvokerNew("Gtk", "recent_chooser_error_quark")
	}

	recentChooserErrorQuarkInvoker.Invoke()
}

var recentManagerErrorQuarkInvoker *gi.Function

// RecentManagerErrorQuark is a representation of the C type gtk_recent_manager_error_quark.
func RecentManagerErrorQuark() {
	if recentManagerErrorQuarkInvoker == nil {
		recentManagerErrorQuarkInvoker = gi.FunctionInvokerNew("Gtk", "recent_manager_error_quark")
	}

	recentManagerErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_render_activity' : non trivial function

// UNSUPPORTED : C value 'gtk_render_arrow' : non trivial function

// UNSUPPORTED : C value 'gtk_render_background' : non trivial function

// UNSUPPORTED : C value 'gtk_render_background_get_clip' : non trivial function

// UNSUPPORTED : C value 'gtk_render_check' : non trivial function

// UNSUPPORTED : C value 'gtk_render_expander' : non trivial function

// UNSUPPORTED : C value 'gtk_render_extension' : non trivial function

// UNSUPPORTED : C value 'gtk_render_focus' : non trivial function

// UNSUPPORTED : C value 'gtk_render_frame' : non trivial function

// UNSUPPORTED : C value 'gtk_render_frame_gap' : non trivial function

// UNSUPPORTED : C value 'gtk_render_handle' : non trivial function

// UNSUPPORTED : C value 'gtk_render_icon' : non trivial function

// UNSUPPORTED : C value 'gtk_render_icon_pixbuf' : non trivial function

// UNSUPPORTED : C value 'gtk_render_icon_surface' : non trivial function

// UNSUPPORTED : C value 'gtk_render_insertion_cursor' : non trivial function

// UNSUPPORTED : C value 'gtk_render_layout' : non trivial function

// UNSUPPORTED : C value 'gtk_render_line' : non trivial function

// UNSUPPORTED : C value 'gtk_render_option' : non trivial function

// UNSUPPORTED : C value 'gtk_render_slider' : non trivial function

// UNSUPPORTED : C value 'gtk_rgb_to_hsv' : non trivial function

// UNSUPPORTED : C value 'gtk_selection_add_target' : non trivial function

// UNSUPPORTED : C value 'gtk_selection_add_targets' : non trivial function

// UNSUPPORTED : C value 'gtk_selection_clear_targets' : non trivial function

// UNSUPPORTED : C value 'gtk_selection_convert' : non trivial function

// UNSUPPORTED : C value 'gtk_selection_owner_set' : non trivial function

// UNSUPPORTED : C value 'gtk_selection_owner_set_for_display' : non trivial function

// UNSUPPORTED : C value 'gtk_selection_remove_all' : non trivial function

// UNSUPPORTED : C value 'gtk_set_debug_flags' : non trivial function

// UNSUPPORTED : C value 'gtk_show_about_dialog' : non trivial function

// UNSUPPORTED : C value 'gtk_show_uri' : non trivial function

// UNSUPPORTED : C value 'gtk_show_uri_on_window' : non trivial function

// UNSUPPORTED : C value 'gtk_stock_add' : non trivial function

// UNSUPPORTED : C value 'gtk_stock_add_static' : non trivial function

var stockListIdsInvoker *gi.Function

// StockListIds is a representation of the C type gtk_stock_list_ids.
func StockListIds() {
	if stockListIdsInvoker == nil {
		stockListIdsInvoker = gi.FunctionInvokerNew("Gtk", "stock_list_ids")
	}

	stockListIdsInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_stock_lookup' : non trivial function

// UNSUPPORTED : C value 'gtk_stock_set_translate_func' : non trivial function

// UNSUPPORTED : C value 'gtk_target_table_free' : non trivial function

// UNSUPPORTED : C value 'gtk_target_table_new_from_list' : non trivial function

// UNSUPPORTED : C value 'gtk_targets_include_image' : non trivial function

// UNSUPPORTED : C value 'gtk_targets_include_rich_text' : non trivial function

// UNSUPPORTED : C value 'gtk_targets_include_text' : non trivial function

// UNSUPPORTED : C value 'gtk_targets_include_uri' : non trivial function

// UNSUPPORTED : C value 'gtk_test_create_simple_window' : non trivial function

// UNSUPPORTED : C value 'gtk_test_create_widget' : non trivial function

// UNSUPPORTED : C value 'gtk_test_display_button_window' : non trivial function

// UNSUPPORTED : C value 'gtk_test_find_label' : non trivial function

// UNSUPPORTED : C value 'gtk_test_find_sibling' : non trivial function

// UNSUPPORTED : C value 'gtk_test_find_widget' : non trivial function

// UNSUPPORTED : C value 'gtk_test_init' : non trivial function

// UNSUPPORTED : C value 'gtk_test_list_all_types' : non trivial function

var testRegisterAllTypesInvoker *gi.Function

// TestRegisterAllTypes is a representation of the C type gtk_test_register_all_types.
func TestRegisterAllTypes() {
	if testRegisterAllTypesInvoker == nil {
		testRegisterAllTypesInvoker = gi.FunctionInvokerNew("Gtk", "test_register_all_types")
	}

	testRegisterAllTypesInvoker.Invoke()
}

// UNSUPPORTED : C value 'gtk_test_slider_get_value' : non trivial function

// UNSUPPORTED : C value 'gtk_test_slider_set_perc' : non trivial function

// UNSUPPORTED : C value 'gtk_test_spin_button_click' : non trivial function

// UNSUPPORTED : C value 'gtk_test_text_get' : non trivial function

// UNSUPPORTED : C value 'gtk_test_text_set' : non trivial function

// UNSUPPORTED : C value 'gtk_test_widget_click' : non trivial function

// UNSUPPORTED : C value 'gtk_test_widget_send_key' : non trivial function

// UNSUPPORTED : C value 'gtk_test_widget_wait_for_draw' : non trivial function

// UNSUPPORTED : C value 'gtk_tree_get_row_drag_data' : non trivial function

// UNSUPPORTED : C value 'gtk_tree_row_reference_deleted' : non trivial function

// UNSUPPORTED : C value 'gtk_tree_row_reference_inserted' : non trivial function

// UNSUPPORTED : C value 'gtk_tree_row_reference_reordered' : non trivial function

// UNSUPPORTED : C value 'gtk_tree_set_row_drag_data' : non trivial function

var trueInvoker *gi.Function

// True is a representation of the C type gtk_true.
func True() {
	if trueInvoker == nil {
		trueInvoker = gi.FunctionInvokerNew("Gtk", "true")
	}

	trueInvoker.Invoke()
}

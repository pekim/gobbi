// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'gtk_accel_groups_activate' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_accel_groups_from_object' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_default_mod_mask' : return type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_label' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_label_with_keycode' : parameter 'display' of type 'Gdk.Display' not supported

// UNSUPPORTED : C value 'gtk_accelerator_name' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_name_with_keycode' : parameter 'display' of type 'Gdk.Display' not supported

// UNSUPPORTED : C value 'gtk_accelerator_parse' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_parse_with_keycode' : parameter 'accelerator_codes' has no type

// UNSUPPORTED : C value 'gtk_accelerator_set_default_mod_mask' : parameter 'default_mod_mask' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_valid' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_alternative_dialog_button_order' : parameter 'screen' of type 'Gdk.Screen' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_add_signal_from_string' : parameter 'binding_set' of type 'BindingSet' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_add_signall' : parameter 'binding_set' of type 'BindingSet' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_remove' : parameter 'binding_set' of type 'BindingSet' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_skip' : parameter 'binding_set' of type 'BindingSet' not supported

// UNSUPPORTED : C value 'gtk_binding_set_by_class' : parameter 'object_class' of type 'gpointer' not supported

var bindingSetFindFunction *gi.Function
var bindingSetFindFunction_Once sync.Once

func bindingSetFindFunction_Set() {
	bindingSetFindFunction_Once.Do(func() {
		bindingSetFindFunction = gi.FunctionInvokerNew("Gtk", "binding_set_find")
	})
}

var bindingSetFindInvoker *gi.Function

// BindingSetFind is a representation of the C type gtk_binding_set_find.
func BindingSetFind(setName string) *BindingSet {
	bindingSetFindFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(setName)

	ret := bindingSetFindFunction.Invoke(inArgs[:], nil)

	retGo := &BindingSet{native: ret.Pointer()}

	return retGo
}

var bindingSetNewFunction *gi.Function
var bindingSetNewFunction_Once sync.Once

func bindingSetNewFunction_Set() {
	bindingSetNewFunction_Once.Do(func() {
		bindingSetNewFunction = gi.FunctionInvokerNew("Gtk", "binding_set_new")
	})
}

var bindingSetNewInvoker *gi.Function

// BindingSetNew is a representation of the C type gtk_binding_set_new.
func BindingSetNew(setName string) *BindingSet {
	bindingSetNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(setName)

	ret := bindingSetNewFunction.Invoke(inArgs[:], nil)

	retGo := &BindingSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_bindings_activate' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_bindings_activate_event' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_builder_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_cairo_should_draw_window' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_cairo_transform_to_window' : parameter 'cr' of type 'cairo.Context' not supported

var checkVersionFunction *gi.Function
var checkVersionFunction_Once sync.Once

func checkVersionFunction_Set() {
	checkVersionFunction_Once.Do(func() {
		checkVersionFunction = gi.FunctionInvokerNew("Gtk", "check_version")
	})
}

var checkVersionInvoker *gi.Function

// CheckVersion is a representation of the C type gtk_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	checkVersionFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(requiredMajor)
	inArgs[1].SetUint32(requiredMinor)
	inArgs[2].SetUint32(requiredMicro)

	ret := checkVersionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_provider_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_device_grab_add' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_device_grab_remove' : parameter 'widget' of type 'Widget' not supported

var disableSetlocaleFunction *gi.Function
var disableSetlocaleFunction_Once sync.Once

func disableSetlocaleFunction_Set() {
	disableSetlocaleFunction_Once.Do(func() {
		disableSetlocaleFunction = gi.FunctionInvokerNew("Gtk", "disable_setlocale")
	})
}

var disableSetlocaleInvoker *gi.Function

// DisableSetlocale is a representation of the C type gtk_disable_setlocale.
func DisableSetlocale() {
	disableSetlocaleFunction_Set()

	disableSetlocaleFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gtk_distribute_natural_allocation' : parameter 'sizes' of type 'RequestedSize' not supported

// UNSUPPORTED : C value 'gtk_drag_cancel' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_finish' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_get_source_widget' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_default' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_gicon' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_name' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_pixbuf' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_stock' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_surface' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_widget' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_draw_insertion_cursor' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_events_pending' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_false' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_file_chooser_error_quark' : return type 'GLib.Quark' not supported

var getBinaryAgeFunction *gi.Function
var getBinaryAgeFunction_Once sync.Once

func getBinaryAgeFunction_Set() {
	getBinaryAgeFunction_Once.Do(func() {
		getBinaryAgeFunction = gi.FunctionInvokerNew("Gtk", "get_binary_age")
	})
}

var getBinaryAgeInvoker *gi.Function

// GetBinaryAge is a representation of the C type gtk_get_binary_age.
func GetBinaryAge() uint32 {
	getBinaryAgeFunction_Set()

	ret := getBinaryAgeFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_current_event' : return type 'Gdk.Event' not supported

// UNSUPPORTED : C value 'gtk_get_current_event_device' : return type 'Gdk.Device' not supported

// UNSUPPORTED : C value 'gtk_get_current_event_state' : parameter 'state' of type 'Gdk.ModifierType' not supported

var getCurrentEventTimeFunction *gi.Function
var getCurrentEventTimeFunction_Once sync.Once

func getCurrentEventTimeFunction_Set() {
	getCurrentEventTimeFunction_Once.Do(func() {
		getCurrentEventTimeFunction = gi.FunctionInvokerNew("Gtk", "get_current_event_time")
	})
}

var getCurrentEventTimeInvoker *gi.Function

// GetCurrentEventTime is a representation of the C type gtk_get_current_event_time.
func GetCurrentEventTime() uint32 {
	getCurrentEventTimeFunction_Set()

	ret := getCurrentEventTimeFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getDebugFlagsFunction *gi.Function
var getDebugFlagsFunction_Once sync.Once

func getDebugFlagsFunction_Set() {
	getDebugFlagsFunction_Once.Do(func() {
		getDebugFlagsFunction = gi.FunctionInvokerNew("Gtk", "get_debug_flags")
	})
}

var getDebugFlagsInvoker *gi.Function

// GetDebugFlags is a representation of the C type gtk_get_debug_flags.
func GetDebugFlags() uint32 {
	getDebugFlagsFunction_Set()

	ret := getDebugFlagsFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_default_language' : return type 'Pango.Language' not supported

// UNSUPPORTED : C value 'gtk_get_event_widget' : parameter 'event' of type 'Gdk.Event' not supported

var getInterfaceAgeFunction *gi.Function
var getInterfaceAgeFunction_Once sync.Once

func getInterfaceAgeFunction_Set() {
	getInterfaceAgeFunction_Once.Do(func() {
		getInterfaceAgeFunction = gi.FunctionInvokerNew("Gtk", "get_interface_age")
	})
}

var getInterfaceAgeInvoker *gi.Function

// GetInterfaceAge is a representation of the C type gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	getInterfaceAgeFunction_Set()

	ret := getInterfaceAgeFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_locale_direction' : return type 'TextDirection' not supported

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() {
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction = gi.FunctionInvokerNew("Gtk", "get_major_version")
	})
}

var getMajorVersionInvoker *gi.Function

// GetMajorVersion is a representation of the C type gtk_get_major_version.
func GetMajorVersion() uint32 {
	getMajorVersionFunction_Set()

	ret := getMajorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() {
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction = gi.FunctionInvokerNew("Gtk", "get_micro_version")
	})
}

var getMicroVersionInvoker *gi.Function

// GetMicroVersion is a representation of the C type gtk_get_micro_version.
func GetMicroVersion() uint32 {
	getMicroVersionFunction_Set()

	ret := getMicroVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() {
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction = gi.FunctionInvokerNew("Gtk", "get_minor_version")
	})
}

var getMinorVersionInvoker *gi.Function

// GetMinorVersion is a representation of the C type gtk_get_minor_version.
func GetMinorVersion() uint32 {
	getMinorVersionFunction_Set()

	ret := getMinorVersionFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_option_group' : parameter 'open_default_display' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_grab_get_current' : return type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_icon_size_from_name' : return type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_get_name' : parameter 'size' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_lookup' : parameter 'size' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_lookup_for_settings' : parameter 'settings' of type 'Settings' not supported

// UNSUPPORTED : C value 'gtk_icon_size_register' : return type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_register_alias' : parameter 'target' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_theme_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_init' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_init_check' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_init_with_args' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_key_snooper_install' : parameter 'snooper' of type 'KeySnoopFunc' not supported

var keySnooperRemoveFunction *gi.Function
var keySnooperRemoveFunction_Once sync.Once

func keySnooperRemoveFunction_Set() {
	keySnooperRemoveFunction_Once.Do(func() {
		keySnooperRemoveFunction = gi.FunctionInvokerNew("Gtk", "key_snooper_remove")
	})
}

var keySnooperRemoveInvoker *gi.Function

// KeySnooperRemove is a representation of the C type gtk_key_snooper_remove.
func KeySnooperRemove(snooperHandlerId uint32) {
	keySnooperRemoveFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(snooperHandlerId)

	keySnooperRemoveFunction.Invoke(inArgs[:], nil)

}

var mainFunction *gi.Function
var mainFunction_Once sync.Once

func mainFunction_Set() {
	mainFunction_Once.Do(func() {
		mainFunction = gi.FunctionInvokerNew("Gtk", "main")
	})
}

var mainInvoker *gi.Function

// Main is a representation of the C type gtk_main.
func Main() {
	mainFunction_Set()

	mainFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gtk_main_do_event' : parameter 'event' of type 'Gdk.Event' not supported

// UNSUPPORTED : C value 'gtk_main_iteration' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_main_iteration_do' : parameter 'blocking' of type 'gboolean' not supported

var mainLevelFunction *gi.Function
var mainLevelFunction_Once sync.Once

func mainLevelFunction_Set() {
	mainLevelFunction_Once.Do(func() {
		mainLevelFunction = gi.FunctionInvokerNew("Gtk", "main_level")
	})
}

var mainLevelInvoker *gi.Function

// MainLevel is a representation of the C type gtk_main_level.
func MainLevel() uint32 {
	mainLevelFunction_Set()

	ret := mainLevelFunction.Invoke(nil, nil)

	retGo := ret.Uint32()

	return retGo
}

var mainQuitFunction *gi.Function
var mainQuitFunction_Once sync.Once

func mainQuitFunction_Set() {
	mainQuitFunction_Once.Do(func() {
		mainQuitFunction = gi.FunctionInvokerNew("Gtk", "main_quit")
	})
}

var mainQuitInvoker *gi.Function

// MainQuit is a representation of the C type gtk_main_quit.
func MainQuit() {
	mainQuitFunction_Set()

	mainQuitFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gtk_paint_arrow' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_box' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_box_gap' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_check' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_diamond' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_expander' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_extension' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_flat_box' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_focus' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_handle' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_hline' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_layout' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_option' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_resize_grip' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_shadow' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_shadow_gap' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_slider' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_spinner' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_tab' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_vline' : parameter 'style' of type 'Style' not supported

var paperSizeGetDefaultFunction *gi.Function
var paperSizeGetDefaultFunction_Once sync.Once

func paperSizeGetDefaultFunction_Set() {
	paperSizeGetDefaultFunction_Once.Do(func() {
		paperSizeGetDefaultFunction = gi.FunctionInvokerNew("Gtk", "paper_size_get_default")
	})
}

var paperSizeGetDefaultInvoker *gi.Function

// PaperSizeGetDefault is a representation of the C type gtk_paper_size_get_default.
func PaperSizeGetDefault() string {
	paperSizeGetDefaultFunction_Set()

	ret := paperSizeGetDefaultFunction.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_get_paper_sizes' : parameter 'include_custom' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_parse_args' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_print_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog_async' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_propagate_event' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_rc_add_default_file' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'gtk_rc_find_module_in_path' : return type 'filename' not supported

// UNSUPPORTED : C value 'gtk_rc_find_pixmap_in_path' : parameter 'settings' of type 'Settings' not supported

var rcGetDefaultFilesFunction *gi.Function
var rcGetDefaultFilesFunction_Once sync.Once

func rcGetDefaultFilesFunction_Set() {
	rcGetDefaultFilesFunction_Once.Do(func() {
		rcGetDefaultFilesFunction = gi.FunctionInvokerNew("Gtk", "rc_get_default_files")
	})
}

var rcGetDefaultFilesInvoker *gi.Function

// RcGetDefaultFiles is a representation of the C type gtk_rc_get_default_files.
func RcGetDefaultFiles() {
	rcGetDefaultFilesFunction_Set()

	rcGetDefaultFilesFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gtk_rc_get_im_module_file' : return type 'filename' not supported

// UNSUPPORTED : C value 'gtk_rc_get_im_module_path' : return type 'filename' not supported

// UNSUPPORTED : C value 'gtk_rc_get_module_dir' : return type 'filename' not supported

// UNSUPPORTED : C value 'gtk_rc_get_style' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_rc_get_style_by_paths' : parameter 'settings' of type 'Settings' not supported

var rcGetThemeDirFunction *gi.Function
var rcGetThemeDirFunction_Once sync.Once

func rcGetThemeDirFunction_Set() {
	rcGetThemeDirFunction_Once.Do(func() {
		rcGetThemeDirFunction = gi.FunctionInvokerNew("Gtk", "rc_get_theme_dir")
	})
}

var rcGetThemeDirInvoker *gi.Function

// RcGetThemeDir is a representation of the C type gtk_rc_get_theme_dir.
func RcGetThemeDir() string {
	rcGetThemeDirFunction_Set()

	ret := rcGetThemeDirFunction.Invoke(nil, nil)

	retGo := ret.String(true)

	return retGo
}

var rcParseFunction *gi.Function
var rcParseFunction_Once sync.Once

func rcParseFunction_Set() {
	rcParseFunction_Once.Do(func() {
		rcParseFunction = gi.FunctionInvokerNew("Gtk", "rc_parse")
	})
}

var rcParseInvoker *gi.Function

// RcParse is a representation of the C type gtk_rc_parse.
func RcParse(filename string) {
	rcParseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	rcParseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_rc_parse_color' : parameter 'scanner' of type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_parse_color_full' : parameter 'scanner' of type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_parse_priority' : parameter 'scanner' of type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_parse_state' : parameter 'scanner' of type 'GLib.Scanner' not supported

var rcParseStringFunction *gi.Function
var rcParseStringFunction_Once sync.Once

func rcParseStringFunction_Set() {
	rcParseStringFunction_Once.Do(func() {
		rcParseStringFunction = gi.FunctionInvokerNew("Gtk", "rc_parse_string")
	})
}

var rcParseStringInvoker *gi.Function

// RcParseString is a representation of the C type gtk_rc_parse_string.
func RcParseString(rcString string) {
	rcParseStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(rcString)

	rcParseStringFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_rc_property_parse_border' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_color' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_enum' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_flags' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_requisition' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_reparse_all' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_rc_reparse_all_for_settings' : parameter 'settings' of type 'Settings' not supported

// UNSUPPORTED : C value 'gtk_rc_reset_styles' : parameter 'settings' of type 'Settings' not supported

// UNSUPPORTED : C value 'gtk_rc_scanner_new' : return type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_set_default_files' : parameter 'filenames' has no type

// UNSUPPORTED : C value 'gtk_recent_chooser_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_recent_manager_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_render_activity' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_arrow' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_background' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_background_get_clip' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_check' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_expander' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_extension' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_focus' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_frame' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_frame_gap' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_handle' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_icon' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_icon_pixbuf' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_icon_surface' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_insertion_cursor' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_layout' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_line' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_option' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_slider' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_rgb_to_hsv' : parameter 'r' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_selection_add_target' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_add_targets' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_clear_targets' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_convert' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_owner_set' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_owner_set_for_display' : parameter 'display' of type 'Gdk.Display' not supported

// UNSUPPORTED : C value 'gtk_selection_remove_all' : parameter 'widget' of type 'Widget' not supported

var setDebugFlagsFunction *gi.Function
var setDebugFlagsFunction_Once sync.Once

func setDebugFlagsFunction_Set() {
	setDebugFlagsFunction_Once.Do(func() {
		setDebugFlagsFunction = gi.FunctionInvokerNew("Gtk", "set_debug_flags")
	})
}

var setDebugFlagsInvoker *gi.Function

// SetDebugFlags is a representation of the C type gtk_set_debug_flags.
func SetDebugFlags(flags uint32) {
	setDebugFlagsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(flags)

	setDebugFlagsFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_show_about_dialog' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_show_uri' : parameter 'screen' of type 'Gdk.Screen' not supported

// UNSUPPORTED : C value 'gtk_show_uri_on_window' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_stock_add' : parameter 'items' has no type

// UNSUPPORTED : C value 'gtk_stock_add_static' : parameter 'items' has no type

// UNSUPPORTED : C value 'gtk_stock_list_ids' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_stock_lookup' : parameter 'item' of type 'StockItem' not supported

// UNSUPPORTED : C value 'gtk_stock_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

// UNSUPPORTED : C value 'gtk_target_table_free' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_target_table_new_from_list' : parameter 'list' of type 'TargetList' not supported

// UNSUPPORTED : C value 'gtk_targets_include_image' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_targets_include_rich_text' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_targets_include_text' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_targets_include_uri' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_test_create_simple_window' : return type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_create_widget' : parameter 'widget_type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_test_display_button_window' : parameter '...' has no type

// UNSUPPORTED : C value 'gtk_test_find_label' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_find_sibling' : parameter 'base_widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_find_widget' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_init' : parameter 'argvp' has no type

var testListAllTypesFunction *gi.Function
var testListAllTypesFunction_Once sync.Once

func testListAllTypesFunction_Set() {
	testListAllTypesFunction_Once.Do(func() {
		testListAllTypesFunction = gi.FunctionInvokerNew("Gtk", "test_list_all_types")
	})
}

var testListAllTypesInvoker *gi.Function

// TestListAllTypes is a representation of the C type gtk_test_list_all_types.
func TestListAllTypes() uint32 {
	testListAllTypesFunction_Set()

	var outArgs [1]gi.Argument

	testListAllTypesFunction.Invoke(nil, outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

var testRegisterAllTypesFunction *gi.Function
var testRegisterAllTypesFunction_Once sync.Once

func testRegisterAllTypesFunction_Set() {
	testRegisterAllTypesFunction_Once.Do(func() {
		testRegisterAllTypesFunction = gi.FunctionInvokerNew("Gtk", "test_register_all_types")
	})
}

var testRegisterAllTypesInvoker *gi.Function

// TestRegisterAllTypes is a representation of the C type gtk_test_register_all_types.
func TestRegisterAllTypes() {
	testRegisterAllTypesFunction_Set()

	testRegisterAllTypesFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gtk_test_slider_get_value' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_slider_set_perc' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_spin_button_click' : parameter 'spinner' of type 'SpinButton' not supported

// UNSUPPORTED : C value 'gtk_test_text_get' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_text_set' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_widget_click' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_widget_send_key' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_widget_wait_for_draw' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_tree_get_row_drag_data' : parameter 'selection_data' of type 'SelectionData' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_deleted' : parameter 'proxy' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_inserted' : parameter 'proxy' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_reordered' : parameter 'proxy' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_tree_set_row_drag_data' : parameter 'selection_data' of type 'SelectionData' not supported

// UNSUPPORTED : C value 'gtk_true' : return type 'gboolean' not supported

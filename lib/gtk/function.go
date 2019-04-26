// This is a generated file - DO NOT EDIT

package gtk

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : gtk_accel_groups_activate : return type :

// Unsupported : gtk_accel_groups_from_object : return type :

// Unsupported : gtk_accelerator_get_default_mod_mask : return type :

// Unsupported : gtk_accelerator_name : return type :

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Unsupported : gtk_accelerator_valid : return type :

// Unsupported : gtk_binding_set_by_class : unsupported parameter object_class : no type generator for gpointer (gpointer) for param object_class

// Unsupported : gtk_binding_set_find : return type :

// Unsupported : gtk_binding_set_new : return type :

// Unsupported : gtk_bindings_activate : return type :

// Unsupported : gtk_builder_error_quark : return type :

// Unsupported : gtk_check_version : return type :

// Unsupported : gtk_css_provider_error_quark : return type :

// DisableSetlocale is a wrapper around the C function gtk_disable_setlocale.
func DisableSetlocale() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(5471, &data)
	return
}

// Unsupported : gtk_drag_get_source_widget : return type :

// Unsupported : gtk_events_pending : return type :

// Unsupported : gtk_false : return type :

// Unsupported : gtk_get_current_event : no return generator

// Unsupported : gtk_get_current_event_device : return type :

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// GetCurrentEventTime is a wrapper around the C function gtk_get_current_event_time.
func GetCurrentEventTime() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5906, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetDebugFlags is a wrapper around the C function gtk_get_debug_flags.
func GetDebugFlags() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5907, &data)
	ret := data.Return.Uint32()

	return ret
}

// Unsupported : gtk_get_default_language : return type :

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_grab_get_current : return type :

// Unsupported : gtk_icon_size_from_name : return type :

// Unsupported : gtk_icon_size_get_name : return type :

// Unsupported : gtk_icon_size_lookup : return type :

// Unsupported : gtk_icon_size_register : return type :

// Unsupported : gtk_icon_theme_error_quark : return type :

// Unsupported : gtk_init_check : return type :

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// Main is a wrapper around the C function gtk_main.
func Main() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(6390, &data)
	return
}

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_main_iteration : return type :

// Unsupported : gtk_main_iteration_do : return type :

// MainLevel is a wrapper around the C function gtk_main_level.
func MainLevel() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(6394, &data)
	ret := data.Return.Uint32()

	return ret
}

// MainQuit is a wrapper around the C function gtk_main_quit.
func MainQuit() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(6395, &data)
	return
}

// Unsupported : gtk_parse_args : return type :

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_rc_find_module_in_path : return type :

// Unsupported : gtk_rc_find_pixmap_in_path : return type :

// Unsupported : gtk_rc_get_default_files : array return type :

// Unsupported : gtk_rc_get_im_module_file : return type :

// Unsupported : gtk_rc_get_im_module_path : return type :

// Unsupported : gtk_rc_get_module_dir : return type :

// Unsupported : gtk_rc_get_style : return type :

// Unsupported : gtk_rc_get_style_by_paths : return type :

// Unsupported : gtk_rc_get_theme_dir : return type :

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Unsupported : gtk_rc_property_parse_border : return type :

// Unsupported : gtk_rc_property_parse_color : return type :

// Unsupported : gtk_rc_property_parse_enum : return type :

// Unsupported : gtk_rc_property_parse_flags : return type :

// Unsupported : gtk_rc_property_parse_requisition : return type :

// Unsupported : gtk_rc_reparse_all : return type :

// Unsupported : gtk_rc_reparse_all_for_settings : return type :

// Unsupported : gtk_rc_scanner_new : return type :

// Unsupported : gtk_recent_chooser_error_quark : return type :

// Unsupported : gtk_recent_manager_error_quark : return type :

// Unsupported : gtk_selection_add_targets : unsupported parameter targets :

// Unsupported : gtk_selection_convert : return type :

// Unsupported : gtk_selection_owner_set : return type :

// Unsupported : gtk_stock_add : unsupported parameter items :

// Unsupported : gtk_stock_add_static : unsupported parameter items :

// Unsupported : gtk_stock_list_ids : return type :

// Unsupported : gtk_stock_lookup : return type :

// Unsupported : gtk_tree_get_row_drag_data : return type :

// Unsupported : gtk_tree_set_row_drag_data : return type :

// Unsupported : gtk_true : return type :

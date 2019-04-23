// This is a generated file - DO NOT EDIT

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Blacklisted : gtk_accel_groups_activate

// Blacklisted : gtk_accel_groups_from_object

// Blacklisted : gtk_accelerator_get_default_mod_mask

// Blacklisted : gtk_accelerator_name

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Blacklisted : gtk_accelerator_set_default_mod_mask

// Blacklisted : gtk_accelerator_valid

// Blacklisted : gtk_binding_entry_add_signall

// Blacklisted : gtk_binding_entry_remove

// Unsupported : gtk_binding_set_by_class : unsupported parameter object_class : no type generator for gpointer (gpointer) for param object_class

// Blacklisted : gtk_binding_set_find

// Blacklisted : gtk_binding_set_new

// Blacklisted : gtk_bindings_activate

// Blacklisted : gtk_builder_error_quark

// Blacklisted : gtk_check_version

// Blacklisted : gtk_css_provider_error_quark

// Blacklisted : gtk_disable_setlocale

// Blacklisted : gtk_distribute_natural_allocation

// Blacklisted : gtk_drag_finish

// Blacklisted : gtk_drag_get_source_widget

// Blacklisted : gtk_drag_set_icon_default

// Blacklisted : gtk_drag_set_icon_pixbuf

// Blacklisted : gtk_drag_set_icon_stock

// Blacklisted : gtk_drag_set_icon_surface

// Blacklisted : gtk_drag_set_icon_widget

// Blacklisted : gtk_events_pending

// Blacklisted : gtk_false

// Unsupported : gtk_get_current_event : no return generator

// Blacklisted : gtk_get_current_event_device

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// Blacklisted : gtk_get_current_event_time

// Blacklisted : gtk_get_debug_flags

// Blacklisted : gtk_get_default_language

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_grab_get_current

// Blacklisted : gtk_icon_size_from_name

// Blacklisted : gtk_icon_size_get_name

// Blacklisted : gtk_icon_size_lookup

// Blacklisted : gtk_icon_size_register

// Blacklisted : gtk_icon_size_register_alias

// Blacklisted : gtk_icon_theme_error_quark

// Init is a wrapper around the C function gtk_init.
func Init(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gtk_init(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// Blacklisted : gtk_init_check

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// Blacklisted : gtk_key_snooper_remove

// Blacklisted : gtk_main

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_main_iteration

// Blacklisted : gtk_main_iteration_do

// Blacklisted : gtk_main_level

// Blacklisted : gtk_main_quit

// Blacklisted : gtk_paint_arrow

// Blacklisted : gtk_paint_box

// Blacklisted : gtk_paint_box_gap

// Blacklisted : gtk_paint_check

// Blacklisted : gtk_paint_diamond

// Blacklisted : gtk_paint_expander

// Blacklisted : gtk_paint_extension

// Blacklisted : gtk_paint_flat_box

// Blacklisted : gtk_paint_focus

// Blacklisted : gtk_paint_handle

// Blacklisted : gtk_paint_hline

// Blacklisted : gtk_paint_layout

// Blacklisted : gtk_paint_option

// Blacklisted : gtk_paint_resize_grip

// Blacklisted : gtk_paint_shadow

// Blacklisted : gtk_paint_shadow_gap

// Blacklisted : gtk_paint_slider

// Blacklisted : gtk_paint_spinner

// Blacklisted : gtk_paint_tab

// Blacklisted : gtk_paint_vline

// Blacklisted : gtk_parse_args

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Blacklisted : gtk_rc_add_default_file

// Blacklisted : gtk_rc_find_module_in_path

// Blacklisted : gtk_rc_find_pixmap_in_path

// Blacklisted : gtk_rc_get_default_files

// Blacklisted : gtk_rc_get_im_module_file

// Blacklisted : gtk_rc_get_im_module_path

// Blacklisted : gtk_rc_get_module_dir

// Blacklisted : gtk_rc_get_style

// Blacklisted : gtk_rc_get_style_by_paths

// Blacklisted : gtk_rc_get_theme_dir

// Blacklisted : gtk_rc_parse

// Blacklisted : gtk_rc_parse_color

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Blacklisted : gtk_rc_parse_string

// Blacklisted : gtk_rc_property_parse_border

// Blacklisted : gtk_rc_property_parse_color

// Blacklisted : gtk_rc_property_parse_enum

// Blacklisted : gtk_rc_property_parse_flags

// Blacklisted : gtk_rc_property_parse_requisition

// Blacklisted : gtk_rc_reparse_all

// Blacklisted : gtk_rc_reparse_all_for_settings

// Blacklisted : gtk_rc_scanner_new

// Blacklisted : gtk_rc_set_default_files

// Blacklisted : gtk_recent_chooser_error_quark

// Blacklisted : gtk_recent_manager_error_quark

// Blacklisted : gtk_selection_add_target

// Unsupported : gtk_selection_add_targets : unsupported parameter targets :

// Blacklisted : gtk_selection_clear_targets

// Blacklisted : gtk_selection_convert

// Blacklisted : gtk_selection_owner_set

// Blacklisted : gtk_selection_remove_all

// Blacklisted : gtk_set_debug_flags

// Unsupported : gtk_stock_add : unsupported parameter items :

// Unsupported : gtk_stock_add_static : unsupported parameter items :

// Blacklisted : gtk_stock_list_ids

// Blacklisted : gtk_stock_lookup

// Blacklisted : gtk_tree_get_row_drag_data

// Blacklisted : gtk_tree_row_reference_deleted

// Blacklisted : gtk_tree_row_reference_inserted

// Blacklisted : gtk_tree_row_reference_reordered

// Blacklisted : gtk_tree_set_row_drag_data

// Blacklisted : gtk_true

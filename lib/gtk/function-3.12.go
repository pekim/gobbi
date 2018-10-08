// This is a generated file - DO NOT EDIT
// +build gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_codes : no param type

// Unsupported : gtk_drag_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_draw_insertion_cursor : unsupported parameter location : Blacklisted record : GdkRectangle

// Unsupported : gtk_get_current_event : no return generator

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetLocaleDirection is a wrapper around the C function gtk_get_locale_direction.
func GetLocaleDirection() TextDirection {
	retC := C.gtk_get_locale_direction()
	retGo := (TextDirection)(retC)

	return retGo
}

// Unsupported : gtk_icon_size_from_name : no return generator

// Unsupported : gtk_icon_size_get_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_lookup : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_lookup_for_settings : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_register : no return generator

// Unsupported : gtk_icon_size_register_alias : unsupported parameter target : no type generator for gint, GtkIconSize

// Unsupported : gtk_init_with_args : unsupported parameter entries : no param type

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc, GtkKeySnoopFunc

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc, GtkPageSetupDoneFunc

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_rc_get_default_files : no return type

// Unsupported : gtk_rc_get_style_by_paths : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Unsupported : gtk_rc_set_default_files : unsupported parameter filenames : no param type

// Unsupported : gtk_render_background_get_clip : unsupported parameter out_clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_selection_add_target : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_add_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_clear_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_show_about_dialog : unsupported parameter ... : varargs

// Unsupported : gtk_stock_add : unsupported parameter items : no param type

// Unsupported : gtk_stock_add_static : unsupported parameter items : no param type

// Unsupported : gtk_stock_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GtkTranslateFunc

// Unsupported : gtk_target_table_free : unsupported parameter targets : no param type

// Unsupported : gtk_target_table_new_from_list : no return type

// Unsupported : gtk_targets_include_image : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_text : unsupported parameter targets : no param type

// Unsupported : gtk_targets_include_uri : unsupported parameter targets : no param type

// Unsupported : gtk_test_create_widget : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// Unsupported : gtk_test_find_sibling : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_find_widget : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_test_init : unsupported parameter argvp : no param type

// Unsupported : gtk_test_list_all_types : no return type

// Unsupported : gtk_tree_get_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_row_reference_reordered : unsupported parameter new_order : no param type

// Unsupported : gtk_tree_set_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

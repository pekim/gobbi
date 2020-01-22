// Code generated - DO NOT EDIT.
// +build gtk_2.4

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
)

// STOCK_DIALOG_AUTHENTICATION is a representation of the C constant GTK_STOCK_DIALOG_AUTHENTICATION.
//
// since 2.4
const STOCK_DIALOG_AUTHENTICATION = "gtk-dialog-authentication"

// STOCK_HARDDISK is a representation of the C constant GTK_STOCK_HARDDISK.
//
// since 2.4
const STOCK_HARDDISK = "gtk-harddisk"

// STOCK_INDENT is a representation of the C constant GTK_STOCK_INDENT.
//
// since 2.4
const STOCK_INDENT = "gtk-indent"

// STOCK_NETWORK is a representation of the C constant GTK_STOCK_NETWORK.
//
// since 2.4
const STOCK_NETWORK = "gtk-network"

// STOCK_UNINDENT is a representation of the C constant GTK_STOCK_UNINDENT.
//
// since 2.4
const STOCK_UNINDENT = "gtk-unindent"

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// BindingsActivateEvent wraps the C function gtk_bindings_activate_event.
//
// since 2.4
func BindingsActivateEvent(object *gobject.Object, event *gdk.EventKey) bool {
	sys_object := object.ToC()
	sys_event := event.ToC()
	retSys := gtk.Fn_gtk_bindings_activate_event(sys_object, sys_event)
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_get_current_event_state : has [in]out param, state

// UNSUPPORTED : gtk_icon_size_lookup : has [in]out param, width

// UNSUPPORTED : gtk_icon_size_lookup_for_settings : has [in]out param, width

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_parse_args : has array param, argv

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_parse_color : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_color_full : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_state : has [in]out param, state

// RcResetStyles wraps the C function gtk_rc_reset_styles.
//
// since 2.4
func RcResetStyles(settings *Settings) {
	sys_settings := settings.ToC()
	gtk.Fn_gtk_rc_reset_styles(sys_settings)
}

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_render_background_get_clip : has [in]out param, out_clip

// UNSUPPORTED : gtk_rgb_to_hsv : has [in]out param, h

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

// UNSUPPORTED : gtk_show_uri : throws

// UNSUPPORTED : gtk_show_uri_on_window : throws

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

// UNSUPPORTED : gtk_stock_lookup : has [in]out param, item

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

// UNSUPPORTED : gtk_target_table_new_from_list : has [in]out param, n_targets

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

// UNSUPPORTED : gtk_test_init : has array param, argvp

// UNSUPPORTED : gtk_test_list_all_types : has [in]out param, n_types

// UNSUPPORTED : gtk_tree_get_row_drag_data : has [in]out param, tree_model

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// UNSUPPORTED : GestureStylusClass : blacklisted

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// UNSUPPORTED : StackAccessibleClass : blacklisted

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// UNSUPPORTED : StackAccessible : blacklisted

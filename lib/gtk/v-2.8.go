// Code generated - DO NOT EDIT.
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
)

// STOCK_FULLSCREEN is a representation of the C constant GTK_STOCK_FULLSCREEN.
//
// since 2.8
const STOCK_FULLSCREEN = "gtk-fullscreen"

// STOCK_INFO is a representation of the C constant GTK_STOCK_INFO.
//
// since 2.8
const STOCK_INFO = "gtk-info"

// STOCK_LEAVE_FULLSCREEN is a representation of the C constant GTK_STOCK_LEAVE_FULLSCREEN.
//
// since 2.8
const STOCK_LEAVE_FULLSCREEN = "gtk-leave-fullscreen"

// FileChooserConfirmation is a representation of the C enumeration GtkFileChooserConfirmation.
type FileChooserConfirmation int

// FileChooserConfirmation_confirm is a representation of the C enumeration member GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM.
const FileChooserConfirmation_confirm = FileChooserConfirmation(0)

// FileChooserConfirmation_accept_filename is a representation of the C enumeration member GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME.
const FileChooserConfirmation_accept_filename = FileChooserConfirmation(1)

// FileChooserConfirmation_select_again is a representation of the C enumeration member GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN.
const FileChooserConfirmation_select_again = FileChooserConfirmation(2)

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// DragSetIconName wraps the C function gtk_drag_set_icon_name.
//
// since 2.8
func DragSetIconName(context *gdk.DragContext, iconName string, hotX int, hotY int) {
	sys_context := context.ToC()
	sys_iconName := iconName
	sys_hotX := hotX
	sys_hotY := hotY
	gtk.Fn_gtk_drag_set_icon_name(sys_context, sys_iconName, sys_hotX, sys_hotY)
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

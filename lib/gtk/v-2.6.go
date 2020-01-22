// Code generated - DO NOT EDIT.
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
)

// STOCK_ABOUT is a representation of the C constant GTK_STOCK_ABOUT.
//
// since 2.6
const STOCK_ABOUT = "gtk-about"

// STOCK_CONNECT is a representation of the C constant GTK_STOCK_CONNECT.
//
// since 2.6
const STOCK_CONNECT = "gtk-connect"

// STOCK_DIRECTORY is a representation of the C constant GTK_STOCK_DIRECTORY.
//
// since 2.6
const STOCK_DIRECTORY = "gtk-directory"

// STOCK_DISCONNECT is a representation of the C constant GTK_STOCK_DISCONNECT.
//
// since 2.6
const STOCK_DISCONNECT = "gtk-disconnect"

// STOCK_EDIT is a representation of the C constant GTK_STOCK_EDIT.
//
// since 2.6
const STOCK_EDIT = "gtk-edit"

// STOCK_FILE is a representation of the C constant GTK_STOCK_FILE.
//
// since 2.6
const STOCK_FILE = "gtk-file"

// STOCK_MEDIA_FORWARD is a representation of the C constant GTK_STOCK_MEDIA_FORWARD.
//
// since 2.6
const STOCK_MEDIA_FORWARD = "gtk-media-forward"

// STOCK_MEDIA_NEXT is a representation of the C constant GTK_STOCK_MEDIA_NEXT.
//
// since 2.6
const STOCK_MEDIA_NEXT = "gtk-media-next"

// STOCK_MEDIA_PAUSE is a representation of the C constant GTK_STOCK_MEDIA_PAUSE.
//
// since 2.6
const STOCK_MEDIA_PAUSE = "gtk-media-pause"

// STOCK_MEDIA_PLAY is a representation of the C constant GTK_STOCK_MEDIA_PLAY.
//
// since 2.6
const STOCK_MEDIA_PLAY = "gtk-media-play"

// STOCK_MEDIA_PREVIOUS is a representation of the C constant GTK_STOCK_MEDIA_PREVIOUS.
//
// since 2.6
const STOCK_MEDIA_PREVIOUS = "gtk-media-previous"

// STOCK_MEDIA_RECORD is a representation of the C constant GTK_STOCK_MEDIA_RECORD.
//
// since 2.6
const STOCK_MEDIA_RECORD = "gtk-media-record"

// STOCK_MEDIA_REWIND is a representation of the C constant GTK_STOCK_MEDIA_REWIND.
//
// since 2.6
const STOCK_MEDIA_REWIND = "gtk-media-rewind"

// STOCK_MEDIA_STOP is a representation of the C constant GTK_STOCK_MEDIA_STOP.
//
// since 2.6
const STOCK_MEDIA_STOP = "gtk-media-stop"

// AcceleratorGetLabel wraps the C function gtk_accelerator_get_label.
//
// since 2.6
func AcceleratorGetLabel(acceleratorKey uint, acceleratorMods gdk.ModifierType) string {
	sys_acceleratorKey := acceleratorKey
	sys_acceleratorMods := (int)(acceleratorMods)
	retSys := gtk.Fn_gtk_accelerator_get_label(sys_acceleratorKey, sys_acceleratorMods)
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// AlternativeDialogButtonOrder wraps the C function gtk_alternative_dialog_button_order.
//
// since 2.6
func AlternativeDialogButtonOrder(screen *gdk.Screen) bool {
	sys_screen := screen.ToC()
	retSys := gtk.Fn_gtk_alternative_dialog_button_order(sys_screen)
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_get_current_event_state : has [in]out param, state

// GetOptionGroup wraps the C function gtk_get_option_group.
//
// since 2.6
func GetOptionGroup(openDefaultDisplay bool) *glib.OptionGroup {
	sys_openDefaultDisplay := openDefaultDisplay
	retSys := gtk.Fn_gtk_get_option_group(sys_openDefaultDisplay)
	ret := glib.OptionGroupNewFromC(retSys)

	return ret
}

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

// ShowAboutDialog wraps the C function gtk_show_about_dialog.
//
// since 2.6
func ShowAboutDialog(parent *Window, firstPropertyName string) {
	sys_parent := parent.ToC()
	sys_firstPropertyName := firstPropertyName
	gtk.Fn_gtk_show_about_dialog(sys_parent, sys_firstPropertyName)
}

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

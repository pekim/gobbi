// Code generated - DO NOT EDIT.
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// ApplicationInhibitFlags is a representation of the C bitfield GtkApplicationInhibitFlags.
type ApplicationInhibitFlags int

// ApplicationInhibitFlags_logout is a representation of the C bitfield member GTK_APPLICATION_INHIBIT_LOGOUT.
const ApplicationInhibitFlags_logout = ApplicationInhibitFlags(1)

// ApplicationInhibitFlags_switch is a representation of the C bitfield member GTK_APPLICATION_INHIBIT_SWITCH.
const ApplicationInhibitFlags_switch = ApplicationInhibitFlags(2)

// ApplicationInhibitFlags_suspend is a representation of the C bitfield member GTK_APPLICATION_INHIBIT_SUSPEND.
const ApplicationInhibitFlags_suspend = ApplicationInhibitFlags(4)

// ApplicationInhibitFlags_idle is a representation of the C bitfield member GTK_APPLICATION_INHIBIT_IDLE.
const ApplicationInhibitFlags_idle = ApplicationInhibitFlags(8)

// AcceleratorGetLabelWithKeycode wraps the C function gtk_accelerator_get_label_with_keycode.
//
// since 3.4
func AcceleratorGetLabelWithKeycode(display *gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	sys_display := display.ToC()
	sys_acceleratorKey := acceleratorKey
	sys_keycode := keycode
	sys_acceleratorMods := (int)(acceleratorMods)
	retSys := gtk.Fn_gtk_accelerator_get_label_with_keycode(sys_display, sys_acceleratorKey, sys_keycode, sys_acceleratorMods)
	ret := retSys

	return ret
}

// AcceleratorNameWithKeycode wraps the C function gtk_accelerator_name_with_keycode.
//
// since 3.4
func AcceleratorNameWithKeycode(display *gdk.Display, acceleratorKey uint, keycode uint, acceleratorMods gdk.ModifierType) string {
	sys_display := display.ToC()
	sys_acceleratorKey := acceleratorKey
	sys_keycode := keycode
	sys_acceleratorMods := (int)(acceleratorMods)
	retSys := gtk.Fn_gtk_accelerator_name_with_keycode(sys_display, sys_acceleratorKey, sys_keycode, sys_acceleratorMods)
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

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

// RenderInsertionCursor wraps the C function gtk_render_insertion_cursor.
//
// since 3.4
func RenderInsertionCursor(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout, index int, direction pango.Direction) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_layout := layout.ToC()
	sys_index := index
	sys_direction := (int)(direction)
	gtk.Fn_gtk_render_insertion_cursor(sys_context, sys_cr, sys_x, sys_y, sys_layout, sys_index, sys_direction)
}

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

// ColorChooserDialog is a representation of the C record GtkColorChooserDialog.
//
// since 3.4
type ColorChooserDialog struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooserDialog that represents the ColorChooserDialog.
func (recv *ColorChooserDialog) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserDialogNewFromC creates a new ColorChooserDialog from a pointer to the C GtkColorChooserDialog that represents the ColorChooserDialog.
func ColorChooserDialogNewFromC(native unsafe.Pointer) *ColorChooserDialog {
	return &ColorChooserDialog{native: native}
}

// ColorChooserWidget is a representation of the C record GtkColorChooserWidget.
//
// since 3.4
type ColorChooserWidget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooserWidget that represents the ColorChooserWidget.
func (recv *ColorChooserWidget) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserWidgetNewFromC creates a new ColorChooserWidget from a pointer to the C GtkColorChooserWidget that represents the ColorChooserWidget.
func ColorChooserWidgetNewFromC(native unsafe.Pointer) *ColorChooserWidget {
	return &ColorChooserWidget{native: native}
}

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// UNSUPPORTED : StackAccessible : blacklisted

// Actionable is a representation of the C interface GtkActionable.
//
// since 3.4
type Actionable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkActionable that represents the Actionable.
func (recv *Actionable) ToC() unsafe.Pointer {
	return recv.native
}

// ActionableNewFromC creates a new Actionable from a pointer to the C GtkActionable that represents the Actionable.
func ActionableNewFromC(native unsafe.Pointer) *Actionable {
	return &Actionable{native: native}
}

// ColorChooser is a representation of the C interface GtkColorChooser.
//
// since 3.4
type ColorChooser struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkColorChooser that represents the ColorChooser.
func (recv *ColorChooser) ToC() unsafe.Pointer {
	return recv.native
}

// ColorChooserNewFromC creates a new ColorChooser from a pointer to the C GtkColorChooser that represents the ColorChooser.
func ColorChooserNewFromC(native unsafe.Pointer) *ColorChooser {
	return &ColorChooser{native: native}
}

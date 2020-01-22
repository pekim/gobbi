// Code generated - DO NOT EDIT.
// +build gtk_3.16

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
	"unsafe"
)

// STYLE_CLASS_LABEL is a representation of the C constant GTK_STYLE_CLASS_LABEL.
//
// since 3.16
const STYLE_CLASS_LABEL = "label"

// STYLE_CLASS_MONOSPACE is a representation of the C constant GTK_STYLE_CLASS_MONOSPACE.
//
// since 3.16
const STYLE_CLASS_MONOSPACE = "monospace"

// STYLE_CLASS_PAPER is a representation of the C constant GTK_STYLE_CLASS_PAPER.
//
// since 3.16
const STYLE_CLASS_PAPER = "paper"

// STYLE_CLASS_STATUSBAR is a representation of the C constant GTK_STYLE_CLASS_STATUSBAR.
//
// since 3.16
const STYLE_CLASS_STATUSBAR = "statusbar"

// STYLE_CLASS_TOUCH_SELECTION is a representation of the C constant GTK_STYLE_CLASS_TOUCH_SELECTION.
//
// since 3.16
const STYLE_CLASS_TOUCH_SELECTION = "touch-selection"

// STYLE_CLASS_UNDERSHOOT is a representation of the C constant GTK_STYLE_CLASS_UNDERSHOOT.
//
// since 3.16
const STYLE_CLASS_UNDERSHOOT = "undershoot"

// STYLE_CLASS_WIDE is a representation of the C constant GTK_STYLE_CLASS_WIDE.
//
// since 3.16
const STYLE_CLASS_WIDE = "wide"

// TextExtendSelection is a representation of the C enumeration GtkTextExtendSelection.
type TextExtendSelection int

// TextExtendSelection_word is a representation of the C enumeration member GTK_TEXT_EXTEND_SELECTION_WORD.
const TextExtendSelection_word = TextExtendSelection(0)

// TextExtendSelection_line is a representation of the C enumeration member GTK_TEXT_EXTEND_SELECTION_LINE.
const TextExtendSelection_line = TextExtendSelection(1)

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// DragCancel wraps the C function gtk_drag_cancel.
//
// since 3.16
func DragCancel(context *gdk.DragContext) {
	sys_context := context.ToC()
	gtk.Fn_gtk_drag_cancel(sys_context)
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

// GLAreaClass is a representation of the C record GtkGLAreaClass.
//
// since 3.16
type GLAreaClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGLAreaClass that represents the GLAreaClass.
func (recv *GLAreaClass) ToC() unsafe.Pointer {
	return recv.native
}

// GLAreaClassNewFromC creates a new GLAreaClass from a pointer to the C GtkGLAreaClass that represents the GLAreaClass.
func GLAreaClassNewFromC(native unsafe.Pointer) *GLAreaClass {
	return &GLAreaClass{native: native}
}

// UNSUPPORTED : GestureStylusClass : blacklisted

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// UNSUPPORTED : StackAccessibleClass : blacklisted

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// GLArea is a representation of the C record GtkGLArea.
//
// since 3.16
type GLArea struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkGLArea that represents the GLArea.
func (recv *GLArea) ToC() unsafe.Pointer {
	return recv.native
}

// GLAreaNewFromC creates a new GLArea from a pointer to the C GtkGLArea that represents the GLArea.
func GLAreaNewFromC(native unsafe.Pointer) *GLArea {
	return &GLArea{native: native}
}

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// UNSUPPORTED : StackAccessible : blacklisted

// StackSidebar is a representation of the C record GtkStackSidebar.
//
// since 3.16
type StackSidebar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkStackSidebar that represents the StackSidebar.
func (recv *StackSidebar) ToC() unsafe.Pointer {
	return recv.native
}

// StackSidebarNewFromC creates a new StackSidebar from a pointer to the C GtkStackSidebar that represents the StackSidebar.
func StackSidebarNewFromC(native unsafe.Pointer) *StackSidebar {
	return &StackSidebar{native: native}
}

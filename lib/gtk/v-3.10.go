// Code generated - DO NOT EDIT.
// +build gtk_3.10

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
	"unsafe"
)

// BaselinePosition is a representation of the C enumeration GtkBaselinePosition.
type BaselinePosition int

// BaselinePosition_top is a representation of the C enumeration member GTK_BASELINE_POSITION_TOP.
const BaselinePosition_top = BaselinePosition(0)

// BaselinePosition_center is a representation of the C enumeration member GTK_BASELINE_POSITION_CENTER.
const BaselinePosition_center = BaselinePosition(1)

// BaselinePosition_bottom is a representation of the C enumeration member GTK_BASELINE_POSITION_BOTTOM.
const BaselinePosition_bottom = BaselinePosition(2)

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

// RenderIconSurface wraps the C function gtk_render_icon_surface.
//
// since 3.10
func RenderIconSurface(context *StyleContext, cr *cairo.Context, surface *cairo.Surface, x float64, y float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_surface := surface.ToC()
	sys_x := x
	sys_y := y
	gtk.Fn_gtk_render_icon_surface(sys_context, sys_cr, sys_surface, sys_x, sys_y)
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

// TestWidgetWaitForDraw wraps the C function gtk_test_widget_wait_for_draw.
//
// since 3.10
func TestWidgetWaitForDraw(widget *Widget) {
	sys_widget := widget.ToC()
	gtk.Fn_gtk_test_widget_wait_for_draw(sys_widget)
}

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

// SearchBar is a representation of the C record GtkSearchBar.
//
// since 3.10
type SearchBar struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GtkSearchBar that represents the SearchBar.
func (recv *SearchBar) ToC() unsafe.Pointer {
	return recv.native
}

// SearchBarNewFromC creates a new SearchBar from a pointer to the C GtkSearchBar that represents the SearchBar.
func SearchBarNewFromC(native unsafe.Pointer) *SearchBar {
	return &SearchBar{native: native}
}

// UNSUPPORTED : StackAccessible : blacklisted

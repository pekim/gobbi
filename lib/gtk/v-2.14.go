// Code generated - DO NOT EDIT.
// +build gtk_2.14

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
)

// STOCK_PAGE_SETUP is a representation of the C constant GTK_STOCK_PAGE_SETUP.
//
// since 2.14
const STOCK_PAGE_SETUP = "gtk-page-setup"

// STOCK_PRINT_ERROR is a representation of the C constant GTK_STOCK_PRINT_ERROR.
//
// since 2.14
const STOCK_PRINT_ERROR = "gtk-print-error"

// STOCK_PRINT_PAUSED is a representation of the C constant GTK_STOCK_PRINT_PAUSED.
//
// since 2.14
const STOCK_PRINT_PAUSED = "gtk-print-paused"

// STOCK_PRINT_REPORT is a representation of the C constant GTK_STOCK_PRINT_REPORT.
//
// since 2.14
const STOCK_PRINT_REPORT = "gtk-print-report"

// STOCK_PRINT_WARNING is a representation of the C constant GTK_STOCK_PRINT_WARNING.
//
// since 2.14
const STOCK_PRINT_WARNING = "gtk-print-warning"

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

// TestCreateSimpleWindow wraps the C function gtk_test_create_simple_window.
//
// since 2.14
func TestCreateSimpleWindow(windowTitle string, dialogText string) *Widget {
	sys_windowTitle := windowTitle
	sys_dialogText := dialogText
	retSys := gtk.Fn_gtk_test_create_simple_window(sys_windowTitle, sys_dialogText)
	ret := WidgetNewFromC(retSys)

	return ret
}

// TestCreateWidget wraps the C function gtk_test_create_widget.
//
// since 2.14
func TestCreateWidget(widgetType uint64, firstPropertyName string) *Widget {
	sys_widgetType := widgetType
	sys_firstPropertyName := firstPropertyName
	retSys := gtk.Fn_gtk_test_create_widget(sys_widgetType, sys_firstPropertyName)
	ret := WidgetNewFromC(retSys)

	return ret
}

// TestDisplayButtonWindow wraps the C function gtk_test_display_button_window.
//
// since 2.14
func TestDisplayButtonWindow(windowTitle string, dialogText string) *Widget {
	sys_windowTitle := windowTitle
	sys_dialogText := dialogText
	retSys := gtk.Fn_gtk_test_display_button_window(sys_windowTitle, sys_dialogText)
	ret := WidgetNewFromC(retSys)

	return ret
}

// TestFindLabel wraps the C function gtk_test_find_label.
//
// since 2.14
func TestFindLabel(widget *Widget, labelPattern string) *Widget {
	sys_widget := widget.ToC()
	sys_labelPattern := labelPattern
	retSys := gtk.Fn_gtk_test_find_label(sys_widget, sys_labelPattern)
	ret := WidgetNewFromC(retSys)

	return ret
}

// TestFindSibling wraps the C function gtk_test_find_sibling.
//
// since 2.14
func TestFindSibling(baseWidget *Widget, widgetType uint64) *Widget {
	sys_baseWidget := baseWidget.ToC()
	sys_widgetType := widgetType
	retSys := gtk.Fn_gtk_test_find_sibling(sys_baseWidget, sys_widgetType)
	ret := WidgetNewFromC(retSys)

	return ret
}

// TestFindWidget wraps the C function gtk_test_find_widget.
//
// since 2.14
func TestFindWidget(widget *Widget, labelPattern string, widgetType uint64) *Widget {
	sys_widget := widget.ToC()
	sys_labelPattern := labelPattern
	sys_widgetType := widgetType
	retSys := gtk.Fn_gtk_test_find_widget(sys_widget, sys_labelPattern, sys_widgetType)
	ret := WidgetNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gtk_test_init : has array param, argvp

// UNSUPPORTED : gtk_test_list_all_types : has [in]out param, n_types

// TestRegisterAllTypes wraps the C function gtk_test_register_all_types.
//
// since 2.14
func TestRegisterAllTypes() {
	gtk.Fn_gtk_test_register_all_types()
}

// TestSliderGetValue wraps the C function gtk_test_slider_get_value.
//
// since 2.14
func TestSliderGetValue(widget *Widget) float64 {
	sys_widget := widget.ToC()
	retSys := gtk.Fn_gtk_test_slider_get_value(sys_widget)
	ret := retSys

	return ret
}

// TestSliderSetPerc wraps the C function gtk_test_slider_set_perc.
//
// since 2.14
func TestSliderSetPerc(widget *Widget, percentage float64) {
	sys_widget := widget.ToC()
	sys_percentage := percentage
	gtk.Fn_gtk_test_slider_set_perc(sys_widget, sys_percentage)
}

// TestSpinButtonClick wraps the C function gtk_test_spin_button_click.
//
// since 2.14
func TestSpinButtonClick(spinner *SpinButton, button uint, upwards bool) bool {
	sys_spinner := spinner.ToC()
	sys_button := button
	sys_upwards := upwards
	retSys := gtk.Fn_gtk_test_spin_button_click(sys_spinner, sys_button, sys_upwards)
	ret := retSys

	return ret
}

// TestTextGet wraps the C function gtk_test_text_get.
//
// since 2.14
func TestTextGet(widget *Widget) string {
	sys_widget := widget.ToC()
	retSys := gtk.Fn_gtk_test_text_get(sys_widget)
	ret := retSys

	return ret
}

// TestTextSet wraps the C function gtk_test_text_set.
//
// since 2.14
func TestTextSet(widget *Widget, string_ string) {
	sys_widget := widget.ToC()
	sys_string_ := string_
	gtk.Fn_gtk_test_text_set(sys_widget, sys_string_)
}

// TestWidgetClick wraps the C function gtk_test_widget_click.
//
// since 2.14
func TestWidgetClick(widget *Widget, button uint, modifiers gdk.ModifierType) bool {
	sys_widget := widget.ToC()
	sys_button := button
	sys_modifiers := (int)(modifiers)
	retSys := gtk.Fn_gtk_test_widget_click(sys_widget, sys_button, sys_modifiers)
	ret := retSys

	return ret
}

// TestWidgetSendKey wraps the C function gtk_test_widget_send_key.
//
// since 2.14
func TestWidgetSendKey(widget *Widget, keyval uint, modifiers gdk.ModifierType) bool {
	sys_widget := widget.ToC()
	sys_keyval := keyval
	sys_modifiers := (int)(modifiers)
	retSys := gtk.Fn_gtk_test_widget_send_key(sys_widget, sys_keyval, sys_modifiers)
	ret := retSys

	return ret
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

// UNSUPPORTED : StackAccessible : blacklisted

// Code generated - DO NOT EDIT.
// +build gtk_3.10

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
func Fn_gtk_icon_set_render_icon_surface(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	cValue2 := (C.int)(param2)

	cValue3 := (*C.GdkWindow)(unsafe.Pointer(param3))

	ret := C.gtk_icon_set_render_icon_surface(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_selection_data_get_data : no array length

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

func Fn_gtk_widget_class_bind_template_child_full(paramInstance unsafe.Pointer, param0 string, param1 bool, param2 uint64) {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	cValue2 := (C.gssize)(param2)

	C.gtk_widget_class_bind_template_child_full(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_widget_class_set_template(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	C.gtk_widget_class_set_template(cValueInstance, cValue0)
}

func Fn_gtk_widget_class_set_template_from_resource(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_class_set_template_from_resource(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_render_icon_surface(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 float64, param4 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (*C.cairo_surface_t)(unsafe.Pointer(param2))

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	C.gtk_render_icon_surface(cValue0, cValue1, cValue2, cValue3, cValue4)
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

func Fn_gtk_test_widget_wait_for_draw(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_test_widget_wait_for_draw(cValue0)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_add_credit_section : parameter 'people' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_get_artists : no array length

// UNSUPPORTED : gtk_about_dialog_get_authors : no array length

// UNSUPPORTED : gtk_about_dialog_get_documenters : no array length

// UNSUPPORTED : gtk_about_dialog_set_artists : parameter 'artists' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_authors : parameter 'authors' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_documenters : parameter 'documenters' is array parameter without length parameter

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_application_get_accels_for_action : no array length

// UNSUPPORTED : gtk_application_get_actions_for_accel : no array length

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

// UNSUPPORTED : gtk_application_set_accels_for_action : parameter 'accels' is array parameter without length parameter

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

func Fn_gtk_box_get_baseline_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_box_get_baseline_position(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_box_set_baseline_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkBaselinePosition)(param0)

	C.gtk_box_set_baseline_position(cValueInstance, cValue0)
}

func Fn_gtk_builder_new_from_file(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_builder_new_from_file(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_builder_new_from_resource(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_builder_new_from_resource(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_builder_new_from_string(param0 string, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gssize)(param1)

	ret := C.gtk_builder_new_from_string(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

func Fn_gtk_builder_get_application(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	ret := C.gtk_builder_get_application(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_builder_set_application(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkApplication)(unsafe.Pointer(param0))

	C.gtk_builder_set_application(cValueInstance, cValue0)
}

func Fn_gtk_button_new_from_icon_name(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_button_new_from_icon_name(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

// UNSUPPORTED : gtk_cell_area_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_entry_get_tabs(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_tabs(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_set_tabs(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoTabArray)(unsafe.Pointer(param0))

	C.gtk_entry_set_tabs(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

func Fn_gtk_grid_get_baseline_row(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	ret := C.gtk_grid_get_baseline_row(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_grid_get_row_baseline_position(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_grid_get_row_baseline_position(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_grid_remove_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_grid_remove_column(cValueInstance, cValue0)
}

func Fn_gtk_grid_remove_row(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_grid_remove_row(cValueInstance, cValue0)
}

func Fn_gtk_grid_set_baseline_row(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_grid_set_baseline_row(cValueInstance, cValue0)
}

func Fn_gtk_grid_set_row_baseline_position(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GtkBaselinePosition)(param1)

	C.gtk_grid_set_row_baseline_position(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_header_bar_new() unsafe.Pointer {
	ret := C.gtk_header_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_header_bar_get_custom_title(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_header_bar_get_custom_title(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_header_bar_get_show_close_button(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_header_bar_get_show_close_button(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_header_bar_get_subtitle(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_header_bar_get_subtitle(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_header_bar_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_header_bar_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_header_bar_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_header_bar_pack_end(cValueInstance, cValue0)
}

func Fn_gtk_header_bar_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_header_bar_pack_start(cValueInstance, cValue0)
}

func Fn_gtk_header_bar_set_custom_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_header_bar_set_custom_title(cValueInstance, cValue0)
}

func Fn_gtk_header_bar_set_show_close_button(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_header_bar_set_show_close_button(cValueInstance, cValue0)
}

func Fn_gtk_header_bar_set_subtitle(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_header_bar_set_subtitle(cValueInstance, cValue0)
}

func Fn_gtk_header_bar_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_header_bar_set_title(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

func Fn_gtk_icon_info_get_base_scale(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_base_scale(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

func Fn_gtk_icon_info_load_surface(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_surface(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

func Fn_gtk_icon_theme_load_icon_for_scale(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, param3 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.GtkIconLookupFlags)(param3)

	cError := (**C.GError)(error)

	ret := C.gtk_icon_theme_load_icon_for_scale(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_load_surface(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.GdkWindow)(unsafe.Pointer(param3))

	cValue4 := (C.GtkIconLookupFlags)(param4)

	cError := (**C.GError)(error)

	ret := C.gtk_icon_theme_load_surface(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_lookup_by_gicon_for_scale(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.GtkIconLookupFlags)(param3)

	ret := C.gtk_icon_theme_lookup_by_gicon_for_scale(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_lookup_icon_for_scale(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, param3 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.GtkIconLookupFlags)(param3)

	ret := C.gtk_icon_theme_lookup_icon_for_scale(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_image_new_from_surface(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.cairo_surface_t)(unsafe.Pointer(param0))

	ret := C.gtk_image_new_from_surface(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_set_from_surface(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_surface_t)(unsafe.Pointer(param0))

	C.gtk_image_set_from_surface(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_get_show_close_button(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_info_bar_get_show_close_button(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_info_bar_set_show_close_button(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_info_bar_set_show_close_button(cValueInstance, cValue0)
}

func Fn_gtk_label_get_lines(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_lines(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_set_lines(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_label_set_lines(cValueInstance, cValue0)
}

func Fn_gtk_list_box_new() unsafe.Pointer {
	ret := C.gtk_list_box_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

func Fn_gtk_list_box_drag_highlight_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkListBoxRow)(unsafe.Pointer(param0))

	C.gtk_list_box_drag_highlight_row(cValueInstance, cValue0)
}

func Fn_gtk_list_box_drag_unhighlight_row(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	C.gtk_list_box_drag_unhighlight_row(cValueInstance)
}

func Fn_gtk_list_box_get_activate_on_single_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_get_activate_on_single_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_list_box_get_adjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_get_adjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_box_get_row_at_index(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_list_box_get_row_at_index(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_box_get_row_at_y(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_list_box_get_row_at_y(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_box_get_selected_row(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_get_selected_row(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_box_get_selection_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_get_selection_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_list_box_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_list_box_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_list_box_invalidate_filter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	C.gtk_list_box_invalidate_filter(cValueInstance)
}

func Fn_gtk_list_box_invalidate_headers(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	C.gtk_list_box_invalidate_headers(cValueInstance)
}

func Fn_gtk_list_box_invalidate_sort(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	C.gtk_list_box_invalidate_sort(cValueInstance)
}

func Fn_gtk_list_box_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_list_box_prepend(cValueInstance, cValue0)
}

func Fn_gtk_list_box_select_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkListBoxRow)(unsafe.Pointer(param0))

	C.gtk_list_box_select_row(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

func Fn_gtk_list_box_set_activate_on_single_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_list_box_set_activate_on_single_click(cValueInstance, cValue0)
}

func Fn_gtk_list_box_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_list_box_set_adjustment(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

func Fn_gtk_list_box_set_placeholder(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_list_box_set_placeholder(cValueInstance, cValue0)
}

func Fn_gtk_list_box_set_selection_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSelectionMode)(param0)

	C.gtk_list_box_set_selection_mode(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_list_box_row_new() unsafe.Pointer {
	ret := C.gtk_list_box_row_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_box_row_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	C.gtk_list_box_row_changed(cValueInstance)
}

func Fn_gtk_list_box_row_get_header(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_row_get_header(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_box_row_get_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_row_get_index(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_list_box_row_set_header(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_list_box_row_set_header(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_places_sidebar_new() unsafe.Pointer {
	ret := C.gtk_places_sidebar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_places_sidebar_add_shortcut(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.gtk_places_sidebar_add_shortcut(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_get_location(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_location(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_places_sidebar_get_nth_bookmark(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_places_sidebar_get_nth_bookmark(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_places_sidebar_get_open_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_open_flags(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_places_sidebar_get_show_desktop(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_show_desktop(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_places_sidebar_list_shortcuts(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_list_shortcuts(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_places_sidebar_remove_shortcut(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.gtk_places_sidebar_remove_shortcut(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_set_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.gtk_places_sidebar_set_location(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_set_open_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPlacesOpenFlags)(param0)

	C.gtk_places_sidebar_set_open_flags(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_set_show_connect_to_server(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_places_sidebar_set_show_connect_to_server(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_set_show_desktop(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_places_sidebar_set_show_desktop(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

func Fn_gtk_revealer_new() unsafe.Pointer {
	ret := C.gtk_revealer_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_revealer_get_child_revealed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRevealer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_revealer_get_child_revealed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_revealer_get_reveal_child(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRevealer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_revealer_get_reveal_child(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_revealer_get_transition_duration(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkRevealer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_revealer_get_transition_duration(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_revealer_get_transition_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRevealer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_revealer_get_transition_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_revealer_set_reveal_child(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRevealer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_revealer_set_reveal_child(cValueInstance, cValue0)
}

func Fn_gtk_revealer_set_transition_duration(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkRevealer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_revealer_set_transition_duration(cValueInstance, cValue0)
}

func Fn_gtk_revealer_set_transition_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRevealer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkRevealerTransitionType)(param0)

	C.gtk_revealer_set_transition_type(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_search_bar_new() unsafe.Pointer {
	ret := C.gtk_search_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_search_bar_connect_entry(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkSearchBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkEntry)(unsafe.Pointer(param0))

	C.gtk_search_bar_connect_entry(cValueInstance, cValue0)
}

func Fn_gtk_search_bar_get_search_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSearchBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_search_bar_get_search_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_search_bar_get_show_close_button(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSearchBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_search_bar_get_show_close_button(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_search_bar_handle_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSearchBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_search_bar_handle_event(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_search_bar_set_search_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSearchBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_search_bar_set_search_mode(cValueInstance, cValue0)
}

func Fn_gtk_search_bar_set_show_close_button(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSearchBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_search_bar_set_show_close_button(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_stack_new() unsafe.Pointer {
	ret := C.gtk_stack_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_stack_add_named(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_stack_add_named(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_stack_add_titled(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 string) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_stack_add_titled(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_stack_get_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_stack_get_transition_duration(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_transition_duration(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_stack_get_transition_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_transition_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_stack_get_visible_child(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_visible_child(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_stack_get_visible_child_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_visible_child_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_stack_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_stack_set_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_stack_set_transition_duration(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_stack_set_transition_duration(cValueInstance, cValue0)
}

func Fn_gtk_stack_set_transition_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStackTransitionType)(param0)

	C.gtk_stack_set_transition_type(cValueInstance, cValue0)
}

func Fn_gtk_stack_set_visible_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_stack_set_visible_child(cValueInstance, cValue0)
}

func Fn_gtk_stack_set_visible_child_full(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStackTransitionType)(param1)

	C.gtk_stack_set_visible_child_full(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_stack_set_visible_child_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_stack_set_visible_child_name(cValueInstance, cValue0)
}

func Fn_gtk_stack_switcher_new() unsafe.Pointer {
	ret := C.gtk_stack_switcher_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_stack_switcher_get_stack(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStackSwitcher)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_switcher_get_stack(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_stack_switcher_set_stack(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStackSwitcher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStack)(unsafe.Pointer(param0))

	C.gtk_stack_switcher_set_stack(cValueInstance, cValue0)
}

func Fn_gtk_style_context_get_scale(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_scale(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_style_context_set_scale(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_style_context_set_scale(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_drag_begin_with_coordinates(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer, param4 int, param5 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	cValue1 := (C.GdkDragAction)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.GdkEvent)(unsafe.Pointer(param3))

	cValue4 := (C.gint)(param4)

	cValue5 := (C.gint)(param5)

	ret := C.gtk_drag_begin_with_coordinates(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_allocated_baseline(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_allocated_baseline(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_preferred_height_and_baseline_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_widget_get_preferred_height_and_baseline_for_width(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_widget_get_scale_factor(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_scale_factor(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_valign_with_baseline(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_valign_with_baseline(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_init_template(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_init_template(cValueInstance)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_size_allocate_with_baseline(paramInstance unsafe.Pointer, param0 *gdk.Rectangle, param1 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_widget_size_allocate_with_baseline(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_close(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_close(cValueInstance)
}

func Fn_gtk_window_set_titlebar(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_window_set_titlebar(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_file_chooser_get_current_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_current_name(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

func Fn_gtk_tree_model_rows_reordered_with_length(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 []int, param3 int) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(&param2[0]))

	cValue3 := (C.gint)(param3)

	C.gtk_tree_model_rows_reordered_with_length(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

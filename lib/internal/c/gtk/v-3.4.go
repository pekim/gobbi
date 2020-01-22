// Code generated - DO NOT EDIT.
// +build gtk_3.4

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static void c_gtk_actionable_set_action_target(GtkActionable* actionable, const gchar* format_string) {
    return gtk_actionable_set_action_target(actionable, format_string, NULL);
}
*/
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
// UNSUPPORTED : gtk_selection_data_get_data : no array length

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

func Fn_gtk_symbolic_color_new_win32(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_symbolic_color_new_win32(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_accelerator_get_label_with_keycode(param0 unsafe.Pointer, param1 uint, param2 uint, param3 int) string {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GdkModifierType)(param3)

	ret := C.gtk_accelerator_get_label_with_keycode(cValue0, cValue1, cValue2, cValue3)

	return C.GoString(ret)
}

func Fn_gtk_accelerator_name_with_keycode(param0 unsafe.Pointer, param1 uint, param2 uint, param3 int) string {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GdkModifierType)(param3)

	ret := C.gtk_accelerator_name_with_keycode(cValue0, cValue1, cValue2, cValue3)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_render_insertion_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 unsafe.Pointer, param5 int, param6 int) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (*C.PangoLayout)(unsafe.Pointer(param4))

	cValue5 := (C.int)(param5)

	cValue6 := (C.PangoDirection)(param6)

	C.gtk_render_insertion_cursor(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

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

func Fn_gtk_application_add_accelerator(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.gtk_application_add_accelerator(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_application_get_accels_for_action : no array length

// UNSUPPORTED : gtk_application_get_actions_for_accel : no array length

func Fn_gtk_application_get_app_menu(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_get_app_menu(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_application_get_menubar(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_get_menubar(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_application_inhibit(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 string) uint {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkApplicationInhibitFlags)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.gtk_application_inhibit(cValueInstance, cValue0, cValue1, cValue2)

	return (uint)(ret)
}

func Fn_gtk_application_is_inhibited(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkApplicationInhibitFlags)(param0)

	ret := C.gtk_application_is_inhibited(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

func Fn_gtk_application_remove_accelerator(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.gtk_application_remove_accelerator(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_application_set_accels_for_action : parameter 'accels' is array parameter without length parameter

func Fn_gtk_application_set_app_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	C.gtk_application_set_app_menu(cValueInstance, cValue0)
}

func Fn_gtk_application_set_menubar(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	C.gtk_application_set_menubar(cValueInstance, cValue0)
}

func Fn_gtk_application_uninhibit(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_application_uninhibit(cValueInstance, cValue0)
}

func Fn_gtk_application_window_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkApplication)(unsafe.Pointer(param0))

	ret := C.gtk_application_window_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_application_window_get_show_menubar(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkApplicationWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_window_get_show_menubar(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_application_window_set_show_menubar(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkApplicationWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_application_window_set_show_menubar(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

func Fn_gtk_builder_add_from_resource(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_builder_add_from_resource(cValueInstance, cValue0, cError)

	return (uint)(ret)
}

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

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

func Fn_gtk_color_chooser_dialog_new(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	ret := C.gtk_color_chooser_dialog_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_chooser_widget_new() unsafe.Pointer {
	ret := C.gtk_color_chooser_widget_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_entry_completion_compute_prefix(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_entry_completion_compute_prefix(cValueInstance, cValue0)

	return C.GoString(ret)
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

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_image_new_from_resource(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_image_new_from_resource(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_menu_new_from_model(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	ret := C.gtk_menu_new_from_model(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_bar_new_from_model(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	ret := C.gtk_menu_bar_new_from_model(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

func Fn_gtk_scale_get_has_origin(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_get_has_origin(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_scale_set_has_origin(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_scale_set_has_origin(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scrolled_window_get_capture_button_press(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_capture_button_press(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_scrolled_window_get_kinetic_scrolling(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_kinetic_scrolling(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_scrolled_window_set_capture_button_press(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_scrolled_window_set_capture_button_press(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_kinetic_scrolling(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_scrolled_window_set_kinetic_scrolling(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_style_context_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	C.gtk_style_context_set_parent(cValueInstance, cValue0)
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

func Fn_gtk_tree_view_get_n_columns(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_n_columns(cValueInstance)

	return (uint)(ret)
}

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

func Fn_gtk_ui_manager_add_ui_from_resource(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_ui_manager_add_ui_from_resource(cValueInstance, cValue0, cError)

	return (uint)(ret)
}

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_get_modifier_mask(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkModifierIntent)(param0)

	ret := C.gtk_widget_get_modifier_mask(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_get_attached_to(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_attached_to(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_hide_titlebar_when_maximized(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_hide_titlebar_when_maximized(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_set_attached_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_window_set_attached_to(cValueInstance, cValue0)
}

func Fn_gtk_window_set_hide_titlebar_when_maximized(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_hide_titlebar_when_maximized(cValueInstance, cValue0)
}

func Fn_gtk_actionable_get_action_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkActionable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_actionable_get_action_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_actionable_get_action_target_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkActionable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_actionable_get_action_target_value(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_actionable_set_action_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkActionable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_actionable_set_action_name(cValueInstance, cValue0)
}

func Fn_gtk_actionable_set_action_target(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkActionable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.c_gtk_actionable_set_action_target(cValueInstance, cValue0)
}

func Fn_gtk_actionable_set_action_target_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.gtk_actionable_set_action_target_value(cValueInstance, cValue0)
}

func Fn_gtk_actionable_set_detailed_action_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkActionable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_actionable_set_detailed_action_name(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

func Fn_gtk_color_chooser_add_palette(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 []gdk.RGBA) {
	cValueInstance := (*C.GtkColorChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.GdkRGBA)(unsafe.Pointer(&param3[0]))

	C.gtk_color_chooser_add_palette(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_color_chooser_get_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_chooser_get_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_chooser_get_use_alpha(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkColorChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_chooser_get_use_alpha(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_color_chooser_set_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_chooser_set_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_chooser_set_use_alpha(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkColorChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_color_chooser_set_use_alpha(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

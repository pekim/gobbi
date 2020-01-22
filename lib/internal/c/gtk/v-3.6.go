// Code generated - DO NOT EDIT.
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
// UNSUPPORTED : gtk_selection_data_get_data : no array length

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

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

func Fn_gtk_accel_label_set_accel(paramInstance unsafe.Pointer, param0 uint, param1 int) {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	C.gtk_accel_label_set_accel(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

func Fn_gtk_action_group_get_accel_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_get_accel_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_group_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_action_group_set_accel_group(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_application_get_accels_for_action : no array length

// UNSUPPORTED : gtk_application_get_actions_for_accel : no array length

func Fn_gtk_application_get_active_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_get_active_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_application_get_window_by_id(paramInstance unsafe.Pointer, param0 uint) unsafe.Pointer {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_application_get_window_by_id(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

// UNSUPPORTED : gtk_application_set_accels_for_action : parameter 'accels' is array parameter without length parameter

func Fn_gtk_application_window_get_id(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkApplicationWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_window_get_id(cValueInstance)

	return (uint)(ret)
}

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_button_get_always_show_image(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_always_show_image(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_button_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_button_set_always_show_image(cValueInstance, cValue0)
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

func Fn_gtk_entry_get_attributes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_attributes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_input_hints(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_input_hints(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_get_input_purpose(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_input_purpose(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_set_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoAttrList)(unsafe.Pointer(param0))

	C.gtk_entry_set_attributes(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_input_hints(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkInputHints)(param0)

	C.gtk_entry_set_input_hints(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_input_purpose(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkInputPurpose)(param0)

	C.gtk_entry_set_input_purpose(cValueInstance, cValue0)
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

func Fn_gtk_icon_view_get_cell_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	ret := C.gtk_icon_view_get_cell_rect(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_level_bar_new() unsafe.Pointer {
	ret := C.gtk_level_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_level_bar_new_for_interval(param0 float64, param1 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	ret := C.gtk_level_bar_new_for_interval(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_level_bar_add_offset_value(paramInstance unsafe.Pointer, param0 string, param1 float64) {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	C.gtk_level_bar_add_offset_value(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_level_bar_get_max_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_level_bar_get_max_value(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_level_bar_get_min_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_level_bar_get_min_value(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_level_bar_get_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_level_bar_get_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_level_bar_get_offset_value(paramInstance unsafe.Pointer, param0 string, param1 *float64) bool {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_level_bar_get_offset_value(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_level_bar_get_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_level_bar_get_value(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_level_bar_remove_offset_value(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_level_bar_remove_offset_value(cValueInstance, cValue0)
}

func Fn_gtk_level_bar_set_max_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_level_bar_set_max_value(cValueInstance, cValue0)
}

func Fn_gtk_level_bar_set_min_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_level_bar_set_min_value(cValueInstance, cValue0)
}

func Fn_gtk_level_bar_set_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkLevelBarMode)(param0)

	C.gtk_level_bar_set_mode(cValueInstance, cValue0)
}

func Fn_gtk_level_bar_set_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkLevelBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_level_bar_set_value(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_button_new() unsafe.Pointer {
	ret := C.gtk_menu_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_button_get_align_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_button_get_align_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_button_get_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_button_get_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_menu_button_get_menu_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_button_get_menu_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_button_get_popup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_button_get_popup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_button_set_align_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_button_set_align_widget(cValueInstance, cValue0)
}

func Fn_gtk_menu_button_set_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkArrowType)(param0)

	C.gtk_menu_button_set_direction(cValueInstance, cValue0)
}

func Fn_gtk_menu_button_set_menu_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	C.gtk_menu_button_set_menu_model(cValueInstance, cValue0)
}

func Fn_gtk_menu_button_set_popup(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_button_set_popup(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_bind_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 bool) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := toCBool(param2)

	C.gtk_menu_shell_bind_model(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_search_entry_new() unsafe.Pointer {
	ret := C.gtk_search_entry_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_view_get_input_hints(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_input_hints(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_input_purpose(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_input_purpose(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_set_input_hints(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkInputHints)(param0)

	C.gtk_text_view_set_input_hints(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_input_purpose(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkInputPurpose)(param0)

	C.gtk_text_view_set_input_purpose(cValueInstance, cValue0)
}

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

func Fn_gtk_widget_insert_action_group(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GActionGroup)(unsafe.Pointer(param1))

	C.gtk_widget_insert_action_group(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

// Code generated - DO NOT EDIT.
// +build gtk_2.8

package gtk

import "unsafe"

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
// UNSUPPORTED : gtk_selection_data_get_data : no array length

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

func Fn_gtk_text_iter_backward_visible_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_visible_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_lines(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_visible_lines(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

func Fn_gtk_text_iter_forward_visible_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_visible_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_lines(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_visible_lines(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_row_reference_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeRowReference)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_row_reference_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func Fn_gtk_drag_set_icon_name(param0 unsafe.Pointer, param1 string, param2 int, param3 int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_drag_set_icon_name(cValue0, cValue1, cValue2, cValue3)
}

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

func Fn_gtk_about_dialog_get_wrap_license(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_wrap_license(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_about_dialog_set_artists : parameter 'artists' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_authors : parameter 'authors' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_documenters : parameter 'documenters' is array parameter without length parameter

func Fn_gtk_about_dialog_set_wrap_license(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_about_dialog_set_wrap_license(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

func Fn_gtk_action_get_accel_closure(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_accel_closure(cValueInstance)

	return unsafe.Pointer(ret)
}

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

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

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

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_dialog_get_response_for_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_dialog_get_response_for_widget(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_entry_completion_get_popup_set_width(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_popup_set_width(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_popup_single_match(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_popup_single_match(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

func Fn_gtk_entry_completion_set_popup_set_width(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_popup_set_width(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_popup_single_match(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_popup_single_match(cValueInstance, cValue0)
}

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

func Fn_gtk_icon_view_create_drag_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_create_drag_icon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_enable_model_drag_dest(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int, param2 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GdkDragAction)(param2)

	C.gtk_icon_view_enable_model_drag_dest(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_view_enable_model_drag_source(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkModifierType)(param0)

	cValue1 := (*C.GtkTargetEntry)(unsafe.Pointer(&param1[0]))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.GdkDragAction)(param3)

	C.gtk_icon_view_enable_model_drag_source(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_view_get_cursor(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkCellRenderer)(unsafe.Pointer(param1))

	ret := C.gtk_icon_view_get_cursor(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_dest_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *int) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkIconViewDropPosition)(unsafe.Pointer(param3))

	ret := C.gtk_icon_view_get_dest_item_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_drag_dest_item(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconViewDropPosition)(unsafe.Pointer(param1))

	C.gtk_icon_view_get_drag_dest_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_get_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (**C.GtkCellRenderer)(unsafe.Pointer(param3))

	ret := C.gtk_icon_view_get_item_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_reorderable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_reorderable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreePath)(unsafe.Pointer(param1))

	ret := C.gtk_icon_view_get_visible_range(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_scroll_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 float32, param3 float32) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := (C.gfloat)(param3)

	C.gtk_icon_view_scroll_to_path(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_icon_view_set_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_icon_view_set_cursor(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_view_set_drag_dest_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconViewDropPosition)(param1)

	C.gtk_icon_view_set_drag_dest_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_set_reorderable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_view_set_reorderable(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_unset_model_drag_dest(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_unset_model_drag_dest(cValueInstance)
}

func Fn_gtk_icon_view_unset_model_drag_source(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_unset_model_drag_source(cValueInstance)
}

func Fn_gtk_image_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	C.gtk_image_clear(cValueInstance)
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

func Fn_gtk_menu_bar_get_child_pack_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_bar_get_child_pack_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_menu_bar_get_pack_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_bar_get_pack_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_menu_bar_set_child_pack_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPackDirection)(param0)

	C.gtk_menu_bar_set_child_pack_direction(cValueInstance, cValue0)
}

func Fn_gtk_menu_bar_set_pack_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPackDirection)(param0)

	C.gtk_menu_bar_set_pack_direction(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_get_take_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_shell_get_take_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_shell_set_take_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_shell_set_take_focus(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scrolled_window_get_hscrollbar(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_hscrollbar(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_get_vscrollbar(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_vscrollbar(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_size_group_get_ignore_hidden(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_size_group_get_ignore_hidden(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_size_group_set_ignore_hidden(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_size_group_set_ignore_hidden(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

func Fn_gtk_tool_button_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_button_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_button_set_icon_name(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

func Fn_gtk_tree_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreePath)(unsafe.Pointer(param1))

	ret := C.gtk_tree_view_get_visible_range(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

func Fn_gtk_tree_view_column_queue_resize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_queue_resize(cValueInstance)
}

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_drag_source_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_drag_source_set_icon_name(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_get_urgency_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_urgency_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_present_with_time(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.gtk_window_present_with_time(cValueInstance, cValue0)
}

func Fn_gtk_window_set_urgency_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_urgency_hint(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_file_chooser_get_do_overwrite_confirmation(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_do_overwrite_confirmation(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_do_overwrite_confirmation(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_do_overwrite_confirmation(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

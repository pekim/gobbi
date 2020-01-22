// Code generated - DO NOT EDIT.
// +build gtk_2.2

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

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

func Fn_gtk_tree_row_reference_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeRowReference)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_row_reference_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

func Fn_gtk_widget_class_find_style_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_widget_class_find_style_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

func Fn_gtk_widget_class_list_style_properties(paramInstance unsafe.Pointer, param0 *uint) []unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	ret := C.gtk_widget_class_list_style_properties(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]unsafe.Pointer, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 uint32) bool {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(param2)

	cValue3 := (C.guint32)(param3)

	ret := C.gtk_selection_owner_set_for_display(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
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

func Fn_gtk_clipboard_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

func Fn_gtk_clipboard_get_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(param1)

	ret := C.gtk_clipboard_get_for_display(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

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

func Fn_gtk_invisible_new_for_screen(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_invisible_new_for_screen(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkInvisible)(unsafe.Pointer(paramInstance))

	ret := C.gtk_invisible_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkInvisible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_invisible_set_screen(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_list_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_list_store_iter_is_valid(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_list_store_move_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_move_after(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_list_store_move_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_move_before(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_list_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_swap(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_menu_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_select_first(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_shell_select_first(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_n_pages(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_n_pages(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_plug_new_for_display(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.Window)(param1)

	ret := C.gtk_plug_new_for_display(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_plug_construct_for_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64) {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.Window)(param1)

	C.gtk_plug_construct_for_display(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_settings_get_for_screen(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_settings_get_for_screen(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

func Fn_gtk_tree_model_sort_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_sort_iter_is_valid(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_selection_count_selected_rows(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_selection_count_selected_rows(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

func Fn_gtk_tree_selection_get_selected_rows(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_tree_selection_get_selected_rows(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

func Fn_gtk_tree_selection_unselect_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_selection_unselect_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_store_iter_is_valid(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_store_move_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_move_after(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_store_move_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_move_before(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_tree_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_swap(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_expand_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_view_expand_to_path(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

func Fn_gtk_tree_view_set_cursor_on_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkCellRenderer)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	C.gtk_tree_view_set_cursor_on_cell(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

func Fn_gtk_tree_view_column_focus_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_tree_view_column_focus_cell(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_get_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(param0)

	ret := C.gtk_widget_get_clipboard(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_root_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_root_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_has_screen(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_screen(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_fullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_fullscreen(cValueInstance)
}

func Fn_gtk_window_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_skip_pager_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_skip_pager_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_skip_taskbar_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_skip_taskbar_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_set_icon_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_window_set_icon_from_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_window_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_window_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_window_set_skip_pager_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_skip_pager_hint(cValueInstance, cValue0)
}

func Fn_gtk_window_set_skip_taskbar_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_skip_taskbar_hint(cValueInstance, cValue0)
}

func Fn_gtk_window_unfullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_unfullscreen(cValueInstance)
}

func Fn_gtk_window_set_auto_startup_notification(param0 bool) {
	cValue0 := toCBool(param0)

	C.gtk_window_set_auto_startup_notification(cValue0)
}

func Fn_gtk_window_set_default_icon_from_file(param0 string, error unsafe.Pointer) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_window_set_default_icon_from_file(cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

func Fn_gtk_tree_model_get_string_from_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_get_string_from_iter(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

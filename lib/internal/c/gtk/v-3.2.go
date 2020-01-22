// Code generated - DO NOT EDIT.
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type CssSection C.GtkCssSection

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
func Fn_gtk_css_section_get_end_line(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_get_end_line(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_css_section_get_end_position(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_get_end_position(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_css_section_get_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_get_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_css_section_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_get_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_css_section_get_section_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_get_section_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_css_section_get_start_line(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_get_start_line(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_css_section_get_start_position(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_get_start_position(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_css_section_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_section_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_css_section_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCssSection)(unsafe.Pointer(paramInstance))

	C.gtk_css_section_unref(cValueInstance)
}

// UNSUPPORTED : gtk_selection_data_get_data : no array length

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

func Fn_gtk_text_iter_assign(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_iter_assign(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

func Fn_gtk_widget_class_set_accessible_role(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.AtkRole)(param0)

	C.gtk_widget_class_set_accessible_role(cValueInstance, cValue0)
}

func Fn_gtk_widget_class_set_accessible_type(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	C.gtk_widget_class_set_accessible_type(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_widget_path_append_for_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_widget_path_append_for_widget(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_widget_path_append_with_siblings(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint) int {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	ret := C.gtk_widget_path_append_with_siblings(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_widget_path_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_widget_path_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	C.gtk_widget_path_unref(cValueInstance)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func Fn_gtk_drag_set_icon_gicon(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GIcon)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_drag_set_icon_gicon(cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_render_icon(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 float64, param4 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkPixbuf)(unsafe.Pointer(param2))

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	C.gtk_render_icon(cValue0, cValue1, cValue2, cValue3, cValue4)
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

func Fn_gtk_adjustment_get_minimum_increment(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_minimum_increment(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_app_chooser_button_get_show_default_item(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_button_get_show_default_item(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_button_set_show_default_item(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_button_set_show_default_item(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_application_get_accels_for_action : no array length

// UNSUPPORTED : gtk_application_get_actions_for_accel : no array length

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

// UNSUPPORTED : gtk_application_set_accels_for_action : parameter 'accels' is array parameter without length parameter

func Fn_gtk_assistant_remove_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_assistant_remove_page(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_button_box_get_child_non_homogeneous(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_button_box_get_child_non_homogeneous(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_button_box_set_child_non_homogeneous(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_button_box_set_child_non_homogeneous(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_container_child_notify(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_container_child_notify(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_css_provider_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkCssProvider)(unsafe.Pointer(paramInstance))

	ret := C.gtk_css_provider_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_placeholder_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_placeholder_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_entry_set_placeholder_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_entry_set_placeholder_text(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

func Fn_gtk_expander_get_resize_toplevel(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_resize_toplevel(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_expander_set_resize_toplevel(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_resize_toplevel(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_font_chooser_dialog_new(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	ret := C.gtk_font_chooser_dialog_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_chooser_widget_new() unsafe.Pointer {
	ret := C.gtk_font_chooser_widget_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

func Fn_gtk_grid_get_child_at(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_grid_get_child_at(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_grid_insert_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_grid_insert_column(cValueInstance, cValue0)
}

func Fn_gtk_grid_insert_next_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkPositionType)(param1)

	C.gtk_grid_insert_next_to(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_grid_insert_row(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_grid_insert_row(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_lock_button_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GPermission)(unsafe.Pointer(param0))

	ret := C.gtk_lock_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_lock_button_get_permission(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLockButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_lock_button_get_permission(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_lock_button_set_permission(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkLockButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GPermission)(unsafe.Pointer(param0))

	C.gtk_lock_button_set_permission(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_overlay_new() unsafe.Pointer {
	ret := C.gtk_overlay_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_overlay_add_overlay(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkOverlay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_overlay_add_overlay(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

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

func Fn_gtk_tree_view_column_get_x_offset(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_x_offset(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_drag_source_set_icon_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_drag_source_set_icon_gicon(cValueInstance, cValue0)
}

func Fn_gtk_widget_has_visible_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_visible_focus(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_get_focus_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_focus_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_set_focus_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_focus_visible(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_font_chooser_get_font(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_font(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_chooser_get_font_desc(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_font_desc(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_chooser_get_font_face(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_font_face(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_chooser_get_font_family(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_font_family(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_chooser_get_font_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_font_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_font_chooser_get_preview_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_preview_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_chooser_get_show_preview_entry(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_show_preview_entry(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

func Fn_gtk_font_chooser_set_font(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_chooser_set_font(cValueInstance, cValue0)
}

func Fn_gtk_font_chooser_set_font_desc(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.gtk_font_chooser_set_font_desc(cValueInstance, cValue0)
}

func Fn_gtk_font_chooser_set_preview_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_chooser_set_preview_text(cValueInstance, cValue0)
}

func Fn_gtk_font_chooser_set_show_preview_entry(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_chooser_set_show_preview_entry(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

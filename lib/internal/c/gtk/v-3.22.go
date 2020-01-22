// Code generated - DO NOT EDIT.
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

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
func Fn_gtk_paper_size_new_from_gvariant(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	ret := C.gtk_paper_size_new_from_gvariant(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_to_gvariant(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_to_gvariant(cValueInstance)

	return unsafe.Pointer(ret)
}

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

func Fn_gtk_show_uri_on_window(param0 unsafe.Pointer, param1 string, param2 uint32, error unsafe.Pointer) bool {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.guint32)(param2)

	cError := (**C.GError)(error)

	ret := C.gtk_show_uri_on_window(cValue0, cValue1, cValue2, cError)

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

func Fn_gtk_clipboard_get_selection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_get_selection(cValueInstance)

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

func Fn_gtk_file_filter_new_from_gvariant(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	ret := C.gtk_file_filter_new_from_gvariant(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

func Fn_gtk_file_filter_to_gvariant(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_filter_to_gvariant(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_gl_area_get_use_es(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gl_area_get_use_es(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gl_area_set_use_es(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_gl_area_set_use_es(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

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

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

func Fn_gtk_menu_place_on_monitor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkMonitor)(unsafe.Pointer(param0))

	C.gtk_menu_place_on_monitor(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

func Fn_gtk_menu_popup_at_pointer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gtk_menu_popup_at_pointer(cValueInstance, cValue0)
}

func Fn_gtk_menu_popup_at_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (C.GdkGravity)(param2)

	cValue3 := (C.GdkGravity)(param3)

	cValue4 := (*C.GdkEvent)(unsafe.Pointer(param4))

	C.gtk_menu_popup_at_rect(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_menu_popup_at_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkGravity)(param1)

	cValue2 := (C.GdkGravity)(param2)

	cValue3 := (*C.GdkEvent)(unsafe.Pointer(param3))

	C.gtk_menu_popup_at_widget(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_pad_controller_new(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.GActionGroup)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkDevice)(unsafe.Pointer(param2))

	ret := C.gtk_pad_controller_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_pad_controller_set_action(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 string, param4 string) {
	cValueInstance := (*C.GtkPadController)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPadActionType)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	C.gtk_pad_controller_set_action(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_pad_controller_set_action_entries(paramInstance unsafe.Pointer, param0 []PadActionEntry, param1 int) {
	cValueInstance := (*C.GtkPadController)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPadActionEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	C.gtk_pad_controller_set_action_entries(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_new_from_gvariant(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	ret := C.gtk_page_setup_new_from_gvariant(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_to_gvariant(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_page_setup_to_gvariant(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_popover_popdown(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	C.gtk_popover_popdown(cValueInstance)
}

func Fn_gtk_popover_popup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	C.gtk_popover_popup(cValueInstance)
}

func Fn_gtk_print_settings_new_from_gvariant(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	ret := C.gtk_print_settings_new_from_gvariant(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_print_settings_to_gvariant(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_to_gvariant(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scrolled_window_get_max_content_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_max_content_height(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_max_content_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_max_content_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_propagate_natural_height(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_propagate_natural_height(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_scrolled_window_get_propagate_natural_width(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_propagate_natural_width(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_scrolled_window_set_max_content_height(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_scrolled_window_set_max_content_height(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_max_content_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_scrolled_window_set_max_content_width(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_propagate_natural_height(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_scrolled_window_set_propagate_natural_height(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_propagate_natural_width(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_scrolled_window_set_propagate_natural_width(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_shortcut_label_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_shortcut_label_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_shortcut_label_get_accelerator(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkShortcutLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_shortcut_label_get_accelerator(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_shortcut_label_get_disabled_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkShortcutLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_shortcut_label_get_disabled_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_shortcut_label_set_accelerator(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkShortcutLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_shortcut_label_set_accelerator(cValueInstance, cValue0)
}

func Fn_gtk_shortcut_label_set_disabled_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkShortcutLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_shortcut_label_set_disabled_text(cValueInstance, cValue0)
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

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_file_chooser_get_choice(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_get_choice(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_remove_choice(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_remove_choice(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_choice(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_file_chooser_set_choice(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

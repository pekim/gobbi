// Code generated - DO NOT EDIT.
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
type GLAreaClass C.GtkGLAreaClass

// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
func Fn_gtk_paper_size_new_from_ipp(param0 string, param1 float64, param2 float64) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	ret := C.gtk_paper_size_new_from_ipp(cValue0, cValue1, cValue2)

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

func Fn_gtk_drag_cancel(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	C.gtk_drag_cancel(cValue0)
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

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

func Fn_gtk_clipboard_get_default(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	ret := C.gtk_clipboard_get_default(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_css_provider_load_from_resource(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkCssProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_css_provider_load_from_resource(cValueInstance, cValue0)
}

func Fn_gtk_entry_grab_focus_without_selecting(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_grab_focus_without_selecting(cValueInstance)
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

func Fn_gtk_gl_area_new() unsafe.Pointer {
	ret := C.gtk_gl_area_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_gl_area_attach_buffers(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	C.gtk_gl_area_attach_buffers(cValueInstance)
}

func Fn_gtk_gl_area_get_auto_render(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gl_area_get_auto_render(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gl_area_get_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gl_area_get_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gl_area_get_error(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gl_area_get_error(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gl_area_get_has_alpha(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gl_area_get_has_alpha(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gl_area_get_has_depth_buffer(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gl_area_get_has_depth_buffer(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gl_area_get_has_stencil_buffer(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gl_area_get_has_stencil_buffer(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gl_area_get_required_version(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_gl_area_get_required_version(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_gl_area_make_current(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	C.gtk_gl_area_make_current(cValueInstance)
}

func Fn_gtk_gl_area_queue_render(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	C.gtk_gl_area_queue_render(cValueInstance)
}

func Fn_gtk_gl_area_set_auto_render(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_gl_area_set_auto_render(cValueInstance, cValue0)
}

func Fn_gtk_gl_area_set_error(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.gtk_gl_area_set_error(cValueInstance, cValue0)
}

func Fn_gtk_gl_area_set_has_alpha(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_gl_area_set_has_alpha(cValueInstance, cValue0)
}

func Fn_gtk_gl_area_set_has_depth_buffer(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_gl_area_set_has_depth_buffer(cValueInstance, cValue0)
}

func Fn_gtk_gl_area_set_has_stencil_buffer(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_gl_area_set_has_stencil_buffer(cValueInstance, cValue0)
}

func Fn_gtk_gl_area_set_required_version(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkGLArea)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_gl_area_set_required_version(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_label_get_xalign(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_xalign(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_label_get_yalign(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_yalign(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_label_set_xalign(paramInstance unsafe.Pointer, param0 float32) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	C.gtk_label_set_xalign(cValueInstance, cValue0)
}

func Fn_gtk_label_set_yalign(paramInstance unsafe.Pointer, param0 float32) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	C.gtk_label_set_yalign(cValueInstance, cValue0)
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

func Fn_gtk_model_button_new() unsafe.Pointer {
	ret := C.gtk_model_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_detach_tab(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_detach_tab(cValueInstance, cValue0)
}

func Fn_gtk_paned_get_wide_handle(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_wide_handle(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_paned_set_wide_handle(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_paned_set_wide_handle(cValueInstance, cValue0)
}

func Fn_gtk_popover_get_transitions_enabled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	ret := C.gtk_popover_get_transitions_enabled(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_popover_set_transitions_enabled(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_popover_set_transitions_enabled(cValueInstance, cValue0)
}

func Fn_gtk_popover_menu_new() unsafe.Pointer {
	ret := C.gtk_popover_menu_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_popover_menu_open_submenu(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPopoverMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_popover_menu_open_submenu(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scrolled_window_get_overlay_scrolling(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_overlay_scrolling(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_scrolled_window_set_overlay_scrolling(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_scrolled_window_set_overlay_scrolling(cValueInstance, cValue0)
}

func Fn_gtk_search_entry_handle_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSearchEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_search_entry_handle_event(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_stack_get_hhomogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_hhomogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_stack_get_vhomogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_vhomogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_stack_set_hhomogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_stack_set_hhomogeneous(cValueInstance, cValue0)
}

func Fn_gtk_stack_set_vhomogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_stack_set_vhomogeneous(cValueInstance, cValue0)
}

func Fn_gtk_stack_sidebar_new() unsafe.Pointer {
	ret := C.gtk_stack_sidebar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_stack_sidebar_get_stack(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStackSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_sidebar_get_stack(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_stack_sidebar_set_stack(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStackSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStack)(unsafe.Pointer(param0))

	C.gtk_stack_sidebar_set_stack(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

func Fn_gtk_text_buffer_insert_markup(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_text_buffer_insert_markup(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_view_get_monospace(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_monospace(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_set_monospace(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_view_set_monospace(cValueInstance, cValue0)
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

func Fn_gtk_widget_get_action_group(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_widget_get_action_group(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_get_titlebar(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_titlebar(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_scrollable_get_border(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBorder)(unsafe.Pointer(param0))

	ret := C.gtk_scrollable_get_border(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

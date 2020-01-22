// Code generated - DO NOT EDIT.
// +build gtk_2.20

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

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

func Fn_gtk_action_get_always_show_image(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_always_show_image(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_always_show_image(cValueInstance, cValue0)
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

func Fn_gtk_cell_renderer_spinner_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_spinner_new()

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

func Fn_gtk_dialog_get_widget_for_response(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_dialog_get_widget_for_response(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_notebook_get_action_widget(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPackType)(param0)

	ret := C.gtk_notebook_get_action_widget(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_set_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkPackType)(param1)

	C.gtk_notebook_set_action_widget(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_offscreen_window_new() unsafe.Pointer {
	ret := C.gtk_offscreen_window_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_offscreen_window_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkOffscreenWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_offscreen_window_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_offscreen_window_get_surface(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkOffscreenWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_offscreen_window_get_surface(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paned_get_handle_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_handle_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_hard_margins(paramInstance unsafe.Pointer, param0 *float64, param1 *float64, param2 *float64, param3 *float64) bool {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	ret := C.gtk_print_context_get_hard_margins(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_range_get_min_slider_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_min_slider_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_range_get_range_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_range_get_range_rect(cValueInstance, cValue0)
}

func Fn_gtk_range_get_slider_range(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_range_get_slider_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_range_get_slider_size_fixed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_slider_size_fixed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_set_min_slider_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_range_set_min_slider_size(cValueInstance, cValue0)
}

func Fn_gtk_range_set_slider_size_fixed(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_slider_size_fixed(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_spinner_new() unsafe.Pointer {
	ret := C.gtk_spinner_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_spinner_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinner)(unsafe.Pointer(paramInstance))

	C.gtk_spinner_start(cValueInstance)
}

func Fn_gtk_spinner_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinner)(unsafe.Pointer(paramInstance))

	C.gtk_spinner_stop(cValueInstance)
}

func Fn_gtk_status_icon_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_name(cValueInstance, cValue0)
}

func Fn_gtk_statusbar_get_message_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_statusbar_get_message_area(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

func Fn_gtk_tool_item_get_ellipsize_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_ellipsize_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_text_alignment(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_text_alignment(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_tool_item_get_text_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_text_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_text_size_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_text_size_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_tool_item_group_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_get_collapsed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_collapsed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_group_get_drop_item(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_tool_item_group_get_drop_item(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_group_get_header_relief(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_header_relief(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_group_get_item_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	ret := C.gtk_tool_item_group_get_item_position(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tool_item_group_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_item_group_get_label_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_label_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_get_n_items(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_n_items(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_tool_item_group_get_nth_item(paramInstance unsafe.Pointer, param0 uint) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_tool_item_group_get_nth_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_tool_item_group_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_item_group_set_collapsed(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_group_set_collapsed(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.gtk_tool_item_group_set_ellipsize(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_header_relief(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkReliefStyle)(param0)

	C.gtk_tool_item_group_set_header_relief(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_item_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_tool_item_group_set_item_position(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_item_group_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_item_group_set_label(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tool_item_group_set_label_widget(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_new() unsafe.Pointer {
	ret := C.gtk_tool_palette_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_add_drag_dest(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkDestDefaults)(param1)

	cValue2 := (C.GtkToolPaletteDragTargets)(param2)

	cValue3 := (C.GdkDragAction)(param3)

	C.gtk_tool_palette_add_drag_dest(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tool_palette_get_drag_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkSelectionData)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_drag_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_drop_group(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_tool_palette_get_drop_group(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_drop_item(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_tool_palette_get_drop_item(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_exclusive(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_exclusive(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tool_palette_get_expand(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_expand(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tool_palette_get_group_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_group_position(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tool_palette_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_icon_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_icon_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_palette_get_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_palette_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_set_drag_source(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkToolPaletteDragTargets)(param0)

	C.gtk_tool_palette_set_drag_source(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_set_exclusive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_tool_palette_set_exclusive(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_palette_set_expand(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_tool_palette_set_expand(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_palette_set_group_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_tool_palette_set_group_position(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_palette_set_icon_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkIconSize)(param0)

	C.gtk_tool_palette_set_icon_size(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_set_style(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkToolbarStyle)(param0)

	C.gtk_tool_palette_set_style(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_unset_icon_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	C.gtk_tool_palette_unset_icon_size(cValueInstance)
}

func Fn_gtk_tool_palette_unset_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	C.gtk_tool_palette_unset_style(cValueInstance)
}

func Fn_gtk_tool_palette_get_drag_target_group() unsafe.Pointer {
	ret := C.gtk_tool_palette_get_drag_target_group()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_drag_target_item() unsafe.Pointer {
	ret := C.gtk_tool_palette_get_drag_target_item()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tooltip_set_icon_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_tooltip_set_icon_from_gicon(cValueInstance, cValue0, cValue1)
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

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

func Fn_gtk_viewport_get_bin_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	ret := C.gtk_viewport_get_bin_window(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_get_mapped(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_mapped(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_realized(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_realized(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_requisition(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRequisition)(unsafe.Pointer(param0))

	C.gtk_widget_get_requisition(cValueInstance, cValue0)
}

func Fn_gtk_widget_has_rc_style(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_rc_style(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_send_focus_change(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_widget_send_focus_change(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_set_mapped(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_mapped(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_realized(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_realized(cValueInstance, cValue0)
}

func Fn_gtk_widget_style_attach(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_style_attach(cValueInstance)
}

func Fn_gtk_window_get_mnemonics_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_mnemonics_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_window_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_window_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_window_set_mnemonics_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_mnemonics_visible(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_tool_shell_get_ellipsize_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_ellipsize_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_text_alignment(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_text_alignment(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_tool_shell_get_text_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_text_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_text_size_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_text_size_group(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

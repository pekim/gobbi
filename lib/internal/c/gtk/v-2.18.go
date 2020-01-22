// Code generated - DO NOT EDIT.
// +build gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

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

static void c_gtk_info_bar_add_buttons(GtkInfoBar* info_bar, const gchar* first_button_text) {
    return gtk_info_bar_add_buttons(info_bar, first_button_text, NULL);
}
*/
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

func Fn_gtk_cell_renderer_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gfloat)(unsafe.Pointer(param0))

	cValue1 := (*C.gfloat)(unsafe.Pointer(param1))

	C.gtk_cell_renderer_get_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_get_padding(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_renderer_get_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_get_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	C.gtk_cell_renderer_set_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_set_padding(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_renderer_set_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_set_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_toggle_get_activatable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_toggle_get_activatable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_toggle_set_activatable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_toggle_set_activatable(cValueInstance, cValue0)
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

func Fn_gtk_entry_new_with_buffer(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkEntryBuffer)(unsafe.Pointer(param0))

	ret := C.gtk_entry_new_with_buffer(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_buffer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_buffer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_set_buffer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkEntryBuffer)(unsafe.Pointer(param0))

	C.gtk_entry_set_buffer(cValueInstance, cValue0)
}

func Fn_gtk_entry_buffer_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_entry_buffer_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_buffer_delete_text(paramInstance unsafe.Pointer, param0 uint, param1 int) uint {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_entry_buffer_delete_text(cValueInstance, cValue0, cValue1)

	return (uint)(ret)
}

func Fn_gtk_entry_buffer_emit_deleted_text(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_entry_buffer_emit_deleted_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_buffer_emit_inserted_text(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 uint) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.guint)(param2)

	C.gtk_entry_buffer_emit_inserted_text(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_entry_buffer_get_bytes(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_bytes(cValueInstance)

	return (uint64)(ret)
}

func Fn_gtk_entry_buffer_get_length(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_length(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_entry_buffer_get_max_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_max_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_buffer_get_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_entry_buffer_insert_text(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 int) uint {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	ret := C.gtk_entry_buffer_insert_text(cValueInstance, cValue0, cValue1, cValue2)

	return (uint)(ret)
}

func Fn_gtk_entry_buffer_set_max_length(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_buffer_set_max_length(cValueInstance, cValue0)
}

func Fn_gtk_entry_buffer_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_entry_buffer_set_text(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_icon_view_get_item_padding(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_item_padding(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_icon_view_set_item_padding(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_item_padding(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_new() unsafe.Pointer {
	ret := C.gtk_info_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_info_bar_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_info_bar_add_action_widget(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_info_bar_add_button(paramInstance unsafe.Pointer, param0 string, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_info_bar_add_button(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_info_bar_add_buttons(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.c_gtk_info_bar_add_buttons(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_get_action_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_info_bar_get_action_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_info_bar_get_content_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_info_bar_get_content_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_info_bar_get_message_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_info_bar_get_message_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_info_bar_response(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_info_bar_response(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_set_default_response(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_info_bar_set_default_response(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_set_message_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkMessageType)(param0)

	C.gtk_info_bar_set_message_type(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_set_response_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	C.gtk_info_bar_set_response_sensitive(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_label_get_current_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_current_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_label_get_track_visited_links(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_track_visited_links(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_set_track_visited_links(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_track_visited_links(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

func Fn_gtk_menu_get_reserve_toggle_size(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_reserve_toggle_size(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_set_reserve_toggle_size(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_set_reserve_toggle_size(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_get_embed_page_setup(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_embed_page_setup(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_get_has_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_has_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_get_n_pages_to_print(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_n_pages_to_print(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_operation_get_support_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_support_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_set_embed_page_setup(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_embed_page_setup(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_has_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_has_selection(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_support_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_support_selection(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_range_get_flippable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_flippable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_set_flippable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_flippable(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_status_icon_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_title(cValueInstance, cValue0)
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

func Fn_gtk_widget_get_allocation(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_get_allocation(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_app_paintable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_app_paintable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_can_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_can_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_can_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_can_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_double_buffered(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_double_buffered(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_has_window(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_has_window(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_receives_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_receives_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_state(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_state(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_has_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_has_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_has_grab(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_grab(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_drawable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_drawable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_toplevel(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_toplevel(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_set_allocation(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_set_allocation(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_can_default(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_can_default(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_can_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_can_focus(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_has_window(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_has_window(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_receives_default(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_receives_default(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_widget_set_window(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_file_chooser_get_create_folders(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_create_folders(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_create_folders(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_create_folders(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

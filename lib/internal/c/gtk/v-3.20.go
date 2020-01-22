// Code generated - DO NOT EDIT.
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

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

func Fn_gtk_text_iter_starts_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_starts_tag(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

func Fn_gtk_widget_class_get_css_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_class_get_css_name(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_widget_class_set_css_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_class_set_css_name(cValueInstance, cValue0)
}

func Fn_gtk_widget_path_iter_get_object_name(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_get_object_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_widget_path_iter_set_object_name(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_set_object_name(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_render_background_get_clip(param0 unsafe.Pointer, param1 float64, param2 float64, param3 float64, param4 float64, param5 unsafe.Pointer) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (*C.GdkRectangle)(unsafe.Pointer(param5))

	C.gtk_render_background_get_clip(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
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

func Fn_gtk_application_window_get_help_overlay(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkApplicationWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_window_get_help_overlay(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_application_window_set_help_overlay(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplicationWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkShortcutsWindow)(unsafe.Pointer(param0))

	C.gtk_application_window_set_help_overlay(cValueInstance, cValue0)
}

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

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

func Fn_gtk_file_chooser_native_new(param0 string, param1 unsafe.Pointer, param2 int, param3 string, param4 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	cValue2 := (C.GtkFileChooserAction)(param2)

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	ret := C.gtk_file_chooser_native_new(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_native_get_accept_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooserNative)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_native_get_accept_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_native_get_cancel_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooserNative)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_native_get_cancel_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_native_set_accept_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooserNative)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_native_set_accept_label(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_native_set_cancel_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooserNative)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_native_set_cancel_label(cValueInstance, cValue0)
}

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

func Fn_gtk_native_dialog_destroy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	C.gtk_native_dialog_destroy(cValueInstance)
}

func Fn_gtk_native_dialog_get_modal(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_native_dialog_get_modal(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_native_dialog_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_native_dialog_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_native_dialog_get_transient_for(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_native_dialog_get_transient_for(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_native_dialog_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_native_dialog_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_native_dialog_hide(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	C.gtk_native_dialog_hide(cValueInstance)
}

func Fn_gtk_native_dialog_run(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_native_dialog_run(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_native_dialog_set_modal(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_native_dialog_set_modal(cValueInstance, cValue0)
}

func Fn_gtk_native_dialog_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_native_dialog_set_title(cValueInstance, cValue0)
}

func Fn_gtk_native_dialog_set_transient_for(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_native_dialog_set_transient_for(cValueInstance, cValue0)
}

func Fn_gtk_native_dialog_show(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNativeDialog)(unsafe.Pointer(paramInstance))

	C.gtk_native_dialog_show(cValueInstance)
}

func Fn_gtk_popover_get_constrain_to(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	ret := C.gtk_popover_get_constrain_to(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_popover_set_constrain_to(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPopoverConstraint)(param0)

	C.gtk_popover_set_constrain_to(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_settings_reset_property(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_settings_reset_property(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_style_context_to_string(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStyleContextPrintFlags)(param0)

	ret := C.gtk_style_context_to_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

func Fn_gtk_text_tag_changed(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextTag)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_tag_changed(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_view_reset_cursor_blink(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_reset_cursor_blink(cValueInstance)
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

func Fn_gtk_widget_get_allocated_size(paramInstance unsafe.Pointer, param0 *gdk.Rectangle, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	cValue1 := (*C.int)(unsafe.Pointer(param1))

	C.gtk_widget_get_allocated_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_focus_on_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_focus_on_click(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_queue_allocate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_queue_allocate(cValueInstance)
}

func Fn_gtk_widget_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_focus_on_click(cValueInstance, cValue0)
}

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

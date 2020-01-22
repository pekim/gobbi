// Code generated - DO NOT EDIT.
// +build gtk_3.12

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

func Fn_gtk_tree_path_new_from_indicesv(param0 []int, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.gint)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	ret := C.gtk_tree_path_new_from_indicesv(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func Fn_gtk_get_locale_direction() int {
	ret := C.gtk_get_locale_direction()

	return (int)(ret)
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

func Fn_gtk_accel_label_get_accel(paramInstance unsafe.Pointer, param0 *uint, param1 *int) {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkModifierType)(unsafe.Pointer(param1))

	C.gtk_accel_label_get_accel(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

func Fn_gtk_action_bar_new() unsafe.Pointer {
	ret := C.gtk_action_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_bar_get_center_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkActionBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_bar_get_center_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_bar_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_action_bar_pack_end(cValueInstance, cValue0)
}

func Fn_gtk_action_bar_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_action_bar_pack_start(cValueInstance, cValue0)
}

func Fn_gtk_action_bar_set_center_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_action_bar_set_center_widget(cValueInstance, cValue0)
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

func Fn_gtk_box_get_center_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_box_get_center_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_box_set_center_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_box_set_center_widget(cValueInstance, cValue0)
}

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

func Fn_gtk_dialog_get_header_bar(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_dialog_get_header_bar(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_max_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_max_width_chars(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_set_max_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_set_max_width_chars(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

func Fn_gtk_flow_box_new() unsafe.Pointer {
	ret := C.gtk_flow_box_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

func Fn_gtk_flow_box_get_activate_on_single_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_activate_on_single_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_flow_box_get_child_at_index(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_flow_box_get_child_at_index(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_flow_box_get_column_spacing(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_column_spacing(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_flow_box_get_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_flow_box_get_max_children_per_line(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_max_children_per_line(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_flow_box_get_min_children_per_line(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_min_children_per_line(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_flow_box_get_row_spacing(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_row_spacing(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_flow_box_get_selected_children(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_selected_children(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_flow_box_get_selection_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_get_selection_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_flow_box_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_flow_box_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_flow_box_invalidate_filter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	C.gtk_flow_box_invalidate_filter(cValueInstance)
}

func Fn_gtk_flow_box_invalidate_sort(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	C.gtk_flow_box_invalidate_sort(cValueInstance)
}

func Fn_gtk_flow_box_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	C.gtk_flow_box_select_all(cValueInstance)
}

func Fn_gtk_flow_box_select_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFlowBoxChild)(unsafe.Pointer(param0))

	C.gtk_flow_box_select_child(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

func Fn_gtk_flow_box_set_activate_on_single_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_flow_box_set_activate_on_single_click(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_set_column_spacing(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_flow_box_set_column_spacing(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

func Fn_gtk_flow_box_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_flow_box_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_flow_box_set_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_set_max_children_per_line(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_flow_box_set_max_children_per_line(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_set_min_children_per_line(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_flow_box_set_min_children_per_line(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_set_row_spacing(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_flow_box_set_row_spacing(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_set_selection_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSelectionMode)(param0)

	C.gtk_flow_box_set_selection_mode(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_flow_box_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_flow_box_set_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	C.gtk_flow_box_unselect_all(cValueInstance)
}

func Fn_gtk_flow_box_unselect_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFlowBoxChild)(unsafe.Pointer(param0))

	C.gtk_flow_box_unselect_child(cValueInstance, cValue0)
}

func Fn_gtk_flow_box_child_new() unsafe.Pointer {
	ret := C.gtk_flow_box_child_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_flow_box_child_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFlowBoxChild)(unsafe.Pointer(paramInstance))

	C.gtk_flow_box_child_changed(cValueInstance)
}

func Fn_gtk_flow_box_child_get_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFlowBoxChild)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_child_get_index(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_flow_box_child_is_selected(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFlowBoxChild)(unsafe.Pointer(paramInstance))

	ret := C.gtk_flow_box_child_is_selected(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

func Fn_gtk_header_bar_get_decoration_layout(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_header_bar_get_decoration_layout(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_header_bar_get_has_subtitle(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_header_bar_get_has_subtitle(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_header_bar_set_decoration_layout(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_header_bar_set_decoration_layout(cValueInstance, cValue0)
}

func Fn_gtk_header_bar_set_has_subtitle(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkHeaderBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_header_bar_set_has_subtitle(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

func Fn_gtk_icon_info_is_symbolic(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_is_symbolic(cValueInstance)

	return toGoBool(ret)
}

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

func Fn_gtk_menu_button_get_popover(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_button_get_popover(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_button_get_use_popover(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_button_get_use_popover(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_button_set_popover(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_button_set_popover(cValueInstance, cValue0)
}

func Fn_gtk_menu_button_set_use_popover(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_button_set_use_popover(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_get_local_only(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_local_only(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_places_sidebar_set_local_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_places_sidebar_set_local_only(cValueInstance, cValue0)
}

func Fn_gtk_popover_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_popover_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_popover_new_from_model(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	ret := C.gtk_popover_new_from_model(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_popover_bind_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_popover_bind_model(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_popover_get_modal(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	ret := C.gtk_popover_get_modal(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_popover_get_relative_to(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	ret := C.gtk_popover_get_relative_to(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_popover_set_modal(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_popover_set_modal(cValueInstance, cValue0)
}

func Fn_gtk_popover_set_pointing_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_popover_set_pointing_to(cValueInstance, cValue0)
}

func Fn_gtk_popover_set_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPositionType)(param0)

	C.gtk_popover_set_position(cValueInstance, cValue0)
}

func Fn_gtk_popover_set_relative_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_popover_set_relative_to(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_stack_get_child_by_name(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_stack_get_child_by_name(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_stack_get_transition_running(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_transition_running(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_widget_get_margin_end(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_end(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_start(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_start(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_set_margin_end(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_end(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_start(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_start(cValueInstance, cValue0)
}

func Fn_gtk_window_is_maximized(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_is_maximized(cValueInstance)

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

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

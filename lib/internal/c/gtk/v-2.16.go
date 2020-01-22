// Code generated - DO NOT EDIT.
// +build gtk_2.16

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static void c_gtk_style_get(GtkStyle* style, GType widget_type, const gchar* first_property_name) {
    return gtk_style_get(style, widget_type, first_property_name, NULL);
}
*/
/*

static void c_gtk_style_get_valist(GtkStyle* style, GType widget_type, const gchar* first_property_name) {
    return gtk_style_get_valist(style, widget_type, first_property_name, NULL);
}
*/
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

type ActivatableIface C.GtkActivatableIface

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
// UNSUPPORTED : gtk_selection_data_get_data : no array length

func Fn_gtk_selection_data_get_selection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_selection(cValueInstance)

	return unsafe.Pointer(ret)
}

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

func Fn_gtk_action_block_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_block_activate(cValueInstance)
}

func Fn_gtk_action_get_gicon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_gicon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_is_important(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_is_important(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_short_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_short_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_stock_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_stock_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_tooltip(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_tooltip(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_visible_horizontal(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_visible_horizontal(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_visible_vertical(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_visible_vertical(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_set_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_action_set_gicon(cValueInstance, cValue0)
}

func Fn_gtk_action_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_action_set_is_important(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_is_important(cValueInstance, cValue0)
}

func Fn_gtk_action_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_label(cValueInstance, cValue0)
}

func Fn_gtk_action_set_short_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_short_label(cValueInstance, cValue0)
}

func Fn_gtk_action_set_stock_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_stock_id(cValueInstance, cValue0)
}

func Fn_gtk_action_set_tooltip(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_tooltip(cValueInstance, cValue0)
}

func Fn_gtk_action_set_visible_horizontal(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_visible_horizontal(cValueInstance, cValue0)
}

func Fn_gtk_action_set_visible_vertical(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_visible_vertical(cValueInstance, cValue0)
}

func Fn_gtk_action_unblock_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_unblock_activate(cValueInstance)
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

func Fn_gtk_cell_view_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_model(cValueInstance)

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

func Fn_gtk_entry_get_current_icon_drag_source(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_current_icon_drag_source(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_get_icon_activatable(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_activatable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_icon_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_entry_get_icon_at_pos(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_entry_get_icon_gicon(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_gicon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_icon_name(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_icon_pixbuf(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_pixbuf(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_icon_sensitive(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_sensitive(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_icon_stock(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_stock(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_icon_storage_type(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_storage_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_entry_get_icon_tooltip_markup(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_tooltip_markup(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_icon_tooltip_text(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_tooltip_text(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_progress_fraction(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_progress_fraction(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_entry_get_progress_pulse_step(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_progress_pulse_step(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_entry_progress_pulse(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_progress_pulse(cValueInstance)
}

func Fn_gtk_entry_set_icon_activatable(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := toCBool(param1)

	C.gtk_entry_set_icon_activatable(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_drag_source(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, param2 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GtkTargetList)(unsafe.Pointer(param1))

	cValue2 := (C.GdkDragAction)(param2)

	C.gtk_entry_set_icon_drag_source(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_entry_set_icon_from_gicon(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GIcon)(unsafe.Pointer(param1))

	C.gtk_entry_set_icon_from_gicon(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_from_icon_name(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_from_icon_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_from_pixbuf(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_entry_set_icon_from_pixbuf(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_from_stock(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_from_stock(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := toCBool(param1)

	C.gtk_entry_set_icon_sensitive(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_tooltip_markup(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_tooltip_markup(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_tooltip_text(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_tooltip_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_progress_fraction(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_entry_set_progress_fraction(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_progress_pulse_step(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_entry_set_progress_pulse_step(cValueInstance, cValue0)
}

func Fn_gtk_entry_unset_invisible_char(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_unset_invisible_char(cValueInstance)
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

func Fn_gtk_im_multicontext_get_context_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_im_multicontext_get_context_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_im_multicontext_set_context_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_im_multicontext_set_context_id(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_image_menu_item_get_always_show_image(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_menu_item_get_always_show_image(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_image_menu_item_get_use_stock(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_menu_item_get_use_stock(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_image_menu_item_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_image_menu_item_set_accel_group(cValueInstance, cValue0)
}

func Fn_gtk_image_menu_item_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_image_menu_item_set_always_show_image(cValueInstance, cValue0)
}

func Fn_gtk_image_menu_item_set_use_stock(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_image_menu_item_set_use_stock(cValueInstance, cValue0)
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

func Fn_gtk_menu_item_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_menu_item_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_use_underline(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_item_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_item_set_label(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_item_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_draw_page_finish(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_draw_page_finish(cValueInstance)
}

func Fn_gtk_print_operation_set_defer_drawing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_set_defer_drawing(cValueInstance)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_print_settings_get_printer_lpi(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_printer_lpi(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_resolution_x(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_resolution_x(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_resolution_y(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_resolution_y(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_set_printer_lpi(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_print_settings_set_printer_lpi(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_resolution_xy(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_print_settings_set_resolution_xy(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

func Fn_gtk_scale_add_mark(paramInstance unsafe.Pointer, param0 float64, param1 int, param2 string) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkPositionType)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_scale_add_mark(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_scale_clear_marks(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	C.gtk_scale_clear_marks(cValueInstance)
}

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_status_icon_get_has_tooltip(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_has_tooltip(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_status_icon_get_tooltip_markup(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_tooltip_markup(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_tooltip_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_tooltip_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_set_has_tooltip(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_status_icon_set_has_tooltip(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_style_get(paramInstance unsafe.Pointer, param0 uint64, param1 string) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_style_get(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_get_style_property(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_style_get_style_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_style_get_valist(paramInstance unsafe.Pointer, param0 uint64, param1 string) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_style_get_valist(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_window_get_default_icon_name() string {
	ret := C.gtk_window_get_default_icon_name()

	return C.GoString(ret)
}

func Fn_gtk_activatable_do_set_related_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_activatable_do_set_related_action(cValueInstance, cValue0)
}

func Fn_gtk_activatable_get_related_action(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_activatable_get_related_action(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_activatable_get_use_action_appearance(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_activatable_get_use_action_appearance(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_activatable_set_related_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_activatable_set_related_action(cValueInstance, cValue0)
}

func Fn_gtk_activatable_set_use_action_appearance(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_activatable_set_use_action_appearance(cValueInstance, cValue0)
}

func Fn_gtk_activatable_sync_action_properties(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_activatable_sync_action_properties(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

func Fn_gtk_orientable_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkOrientable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_orientable_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_orientable_set_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkOrientable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	C.gtk_orientable_set_orientation(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

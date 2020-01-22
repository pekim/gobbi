// Code generated - DO NOT EDIT.
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static GtkWidget* c_gtk_test_create_widget(GType widget_type, const gchar* first_property_name) {
    return gtk_test_create_widget(widget_type, first_property_name, NULL);
}
*/
/*

static GtkWidget* c_gtk_test_display_button_window(const gchar* window_title, const gchar* dialog_text) {
    return gtk_test_display_button_window(window_title, dialog_text, NULL);
}
*/
/*

static void c_gtk_test_init(int* argcp, char*** argvp) {
    return gtk_test_init(argcp, argvp, NULL);
}
*/
import "C"

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
func Fn_gtk_border_new() unsafe.Pointer {
	ret := C.gtk_border_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_selection_data_get_data : no array length

func Fn_gtk_selection_data_get_data_type(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_data_type(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_selection_data_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_selection_data_get_format(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_format(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_selection_data_get_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_selection_data_get_target(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_target(cValueInstance)

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

func Fn_gtk_rgb_to_hsv(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	cValue4 := (*C.gdouble)(unsafe.Pointer(param4))

	cValue5 := (*C.gdouble)(unsafe.Pointer(param5))

	C.gtk_rgb_to_hsv(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_show_uri(param0 unsafe.Pointer, param1 string, param2 uint32, error unsafe.Pointer) bool {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.guint32)(param2)

	cError := (**C.GError)(error)

	ret := C.gtk_show_uri(cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

func Fn_gtk_test_create_simple_window(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_test_create_simple_window(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_create_widget(param0 uint64, param1 string) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.c_gtk_test_create_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_display_button_window(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.c_gtk_test_display_button_window(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_find_label(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_test_find_label(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_find_sibling(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.gtk_test_find_sibling(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_find_widget(param0 unsafe.Pointer, param1 string, param2 uint64) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GType)(param2)

	ret := C.gtk_test_find_widget(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_init(param0 *int, param1 *[]string) {
	cValue0 := (*C.int)(unsafe.Pointer(param0))

	var cValue1ArrayPointer **C.gchar
	cValue1 := &cValue1ArrayPointer
	param1Indirected := *param1
	param1IndirectedLen := len(param1Indirected)
	cValue1Array := C.malloc((C.ulong)(param1IndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1IndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1IndirectedLen:param1IndirectedLen]
	for param1Indirectedi, param1IndirectedString := range param1Indirected {
		param1IndirectedSlice[param1Indirectedi] = (*C.gchar)(C.CString(param1IndirectedString))
		defer C.free(unsafe.Pointer(param1IndirectedSlice[param1Indirectedi]))
	}
	if len(param1IndirectedSlice) > 0 {
		cValue1ArrayPointer = &param1IndirectedSlice[0]
	}

	C.c_gtk_test_init(cValue0, cValue1)

	param1OutLen := int(*cValue0)
	param1Out := make([]string, param1OutLen, param1OutLen)
	if param1OutLen > 0 {
		param1OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1ArrayPointer))[:param1OutLen:param1OutLen]
		for param1Outi, param1OutCString := range param1OutCSlice {
			param1Out[param1Outi] = C.GoString(param1OutCString)
		}
	}
	*param1 = param1Out
}

func Fn_gtk_test_list_all_types(param0 *uint) []uint64 {
	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	ret := C.gtk_test_list_all_types(cValue0)

	retLen := int(*cValue0)
	retGo := make([]uint64, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint64))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gtk_test_register_all_types() {
	C.gtk_test_register_all_types()
}

func Fn_gtk_test_slider_get_value(param0 unsafe.Pointer) float64 {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_test_slider_get_value(cValue0)

	return (float64)(ret)
}

func Fn_gtk_test_slider_set_perc(param0 unsafe.Pointer, param1 float64) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.double)(param1)

	C.gtk_test_slider_set_perc(cValue0, cValue1)
}

func Fn_gtk_test_spin_button_click(param0 unsafe.Pointer, param1 uint, param2 bool) bool {
	cValue0 := (*C.GtkSpinButton)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := toCBool(param2)

	ret := C.gtk_test_spin_button_click(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_test_text_get(param0 unsafe.Pointer) string {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_test_text_get(cValue0)

	return C.GoString(ret)
}

func Fn_gtk_test_text_set(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_test_text_set(cValue0, cValue1)
}

func Fn_gtk_test_widget_click(param0 unsafe.Pointer, param1 uint, param2 int) bool {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	ret := C.gtk_test_widget_click(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_test_widget_send_key(param0 unsafe.Pointer, param1 uint, param2 int) bool {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	ret := C.gtk_test_widget_send_key(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_add_credit_section : parameter 'people' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_get_artists : no array length

// UNSUPPORTED : gtk_about_dialog_get_authors : no array length

// UNSUPPORTED : gtk_about_dialog_get_documenters : no array length

// UNSUPPORTED : gtk_about_dialog_set_artists : parameter 'artists' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_authors : parameter 'authors' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_documenters : parameter 'documenters' is array parameter without length parameter

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

func Fn_gtk_accel_group_get_is_locked(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accel_group_get_is_locked(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_accel_group_get_modifier_mask(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accel_group_get_modifier_mask(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

func Fn_gtk_adjustment_configure(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_adjustment_configure(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_adjustment_get_lower(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_lower(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_page_increment(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_page_increment(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_page_size(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_page_size(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_step_increment(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_step_increment(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_upper(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_upper(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_set_lower(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_lower(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_page_increment(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_page_increment(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_page_size(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_page_size(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_step_increment(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_step_increment(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_upper(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_upper(cValueInstance, cValue0)
}

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

func Fn_gtk_calendar_get_detail_height_rows(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_calendar_get_detail_height_rows(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_calendar_get_detail_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_calendar_get_detail_width_chars(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

func Fn_gtk_calendar_set_detail_height_rows(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_calendar_set_detail_height_rows(cValueInstance, cValue0)
}

func Fn_gtk_calendar_set_detail_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_calendar_set_detail_width_chars(cValueInstance, cValue0)
}

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

func Fn_gtk_clipboard_wait_is_uris_available(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_is_uris_available(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

func Fn_gtk_color_selection_dialog_get_color_selection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkColorSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_selection_dialog_get_color_selection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_get_button_sensitivity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_button_sensitivity(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

func Fn_gtk_combo_box_set_button_sensitivity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSensitivityType)(param0)

	C.gtk_combo_box_set_button_sensitivity(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_container_get_focus_child(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_get_focus_child(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_get_action_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_dialog_get_action_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_get_content_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_dialog_get_content_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_overwrite_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_overwrite_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_text_length(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_text_length(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_entry_set_overwrite_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_set_overwrite_mode(cValueInstance, cValue0)
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

func Fn_gtk_font_selection_get_face(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_face(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_face_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_face_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_family(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_family(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_family_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_family_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_preview_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_preview_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_font_selection_get_size_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_size_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_size_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_size_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_dialog_get_cancel_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_dialog_get_cancel_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_dialog_get_ok_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_dialog_get_ok_button(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

func Fn_gtk_hsv_new() unsafe.Pointer {
	ret := C.gtk_hsv_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_hsv_get_color(paramInstance unsafe.Pointer, param0 *float64, param1 *float64, param2 *float64) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	C.gtk_hsv_get_color(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_hsv_get_metrics(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_hsv_get_metrics(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_hsv_is_adjusting(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	ret := C.gtk_hsv_is_adjusting(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_hsv_set_color(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 float64) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (C.double)(param0)

	cValue1 := (C.double)(param1)

	cValue2 := (C.double)(param2)

	C.gtk_hsv_set_color(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_hsv_set_metrics(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_hsv_set_metrics(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_hsv_to_rgb(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	cValue4 := (*C.gdouble)(unsafe.Pointer(param4))

	cValue5 := (*C.gdouble)(unsafe.Pointer(param5))

	C.gtk_hsv_to_rgb(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_handle_box_get_child_detached(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_handle_box_get_child_detached(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

func Fn_gtk_icon_info_new_for_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkIconTheme)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	ret := C.gtk_icon_info_new_for_pixbuf(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

func Fn_gtk_icon_theme_lookup_by_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	ret := C.gtk_icon_theme_lookup_by_gicon(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_image_new_from_gicon(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_image_new_from_gicon(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_get_gicon(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconSize)(unsafe.Pointer(param1))

	C.gtk_image_get_gicon(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_set_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_set_from_gicon(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_layout_get_bin_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_layout_get_bin_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_link_button_get_visited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_link_button_get_visited(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_link_button_set_visited(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_link_button_set_visited(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

func Fn_gtk_menu_get_accel_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_accel_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_menu_get_monitor(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_monitor(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_item_get_accel_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_accel_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_message_dialog_get_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_message_dialog_get_image(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_mount_operation_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	ret := C.gtk_mount_operation_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_mount_operation_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_mount_operation_get_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_mount_operation_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_mount_operation_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_mount_operation_is_showing(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_mount_operation_is_showing(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_mount_operation_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_mount_operation_set_parent(cValueInstance, cValue0)
}

func Fn_gtk_mount_operation_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_mount_operation_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_load_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_load_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_page_setup_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_load_key_file(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_gtk_plug_get_embedded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	ret := C.gtk_plug_get_embedded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_plug_get_socket_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	ret := C.gtk_plug_get_socket_window(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_print_settings_get_number_up_layout(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_number_up_layout(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_load_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_load_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_load_key_file(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_set_number_up_layout(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkNumberUpLayout)(param0)

	C.gtk_print_settings_set_number_up_layout(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scale_button_get_minus_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_minus_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_plus_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_plus_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_popup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_popup(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_socket_get_plug_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSocket)(unsafe.Pointer(paramInstance))

	ret := C.gtk_socket_get_plug_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_gicon(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	ret := C.gtk_status_icon_new_from_gicon(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_gicon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_gicon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_x11_window_id(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_x11_window_id(cValueInstance)

	return (uint32)(ret)
}

func Fn_gtk_status_icon_set_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_status_icon_set_from_gicon(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

func Fn_gtk_tool_item_toolbar_reconfigured(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_toolbar_reconfigured(cValueInstance)
}

func Fn_gtk_tooltip_set_icon_from_icon_name(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_tooltip_set_icon_from_icon_name(cValueInstance, cValue0, cValue1)
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

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_get_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_window(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_get_default_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_default_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_group_list_windows(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_group_list_windows(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_file_chooser_get_current_folder_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_current_folder_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_files(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_files(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_preview_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_select_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_select_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_current_folder_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_set_current_folder_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_set_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_unselect_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.gtk_file_chooser_unselect_file(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_tool_shell_get_icon_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_icon_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_relief_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_relief_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_rebuild_menu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	C.gtk_tool_shell_rebuild_menu(cValueInstance)
}

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

// Code generated - DO NOT EDIT.
// +build gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

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

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_widget_path_iter_get_state(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_get_state(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_widget_path_iter_set_state(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GtkStateFlags)(param1)

	C.gtk_widget_path_iter_set_state(cValueInstance, cValue0, cValue1)
}

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

func Fn_gtk_application_get_menu_by_id(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_application_get_menu_by_id(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

func Fn_gtk_application_prefers_app_menu(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_prefers_app_menu(cValueInstance)

	return toGoBool(ret)
}

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

func Fn_gtk_cell_area_attribute_get_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) int {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_cell_area_attribute_get_column(cValueInstance, cValue0, cValue1)

	return (int)(ret)
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

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

func Fn_gtk_event_controller_get_propagation_phase(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEventController)(unsafe.Pointer(paramInstance))

	ret := C.gtk_event_controller_get_propagation_phase(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_event_controller_get_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEventController)(unsafe.Pointer(paramInstance))

	ret := C.gtk_event_controller_get_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_event_controller_handle_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEventController)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_event_controller_handle_event(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_event_controller_reset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEventController)(unsafe.Pointer(paramInstance))

	C.gtk_event_controller_reset(cValueInstance)
}

func Fn_gtk_event_controller_set_propagation_phase(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEventController)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPropagationPhase)(param0)

	C.gtk_event_controller_set_propagation_phase(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_gesture_get_bounding_box(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_get_bounding_box(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_gesture_get_bounding_box_center(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_gesture_get_bounding_box_center(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_gesture_get_device(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_get_device(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_get_last_updated_sequence(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_get_last_updated_sequence(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_get_point(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *float64, param2 *float64) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventSequence)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	ret := C.gtk_gesture_get_point(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_gesture_get_sequence_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventSequence)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_get_sequence_state(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_gesture_get_sequences(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_get_sequences(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_get_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_get_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkGesture)(unsafe.Pointer(param0))

	C.gtk_gesture_group(cValueInstance, cValue0)
}

func Fn_gtk_gesture_handles_sequence(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventSequence)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_handles_sequence(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_gesture_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gesture_is_grouped_with(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkGesture)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_is_grouped_with(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_gesture_is_recognized(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_is_recognized(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gesture_set_sequence_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventSequence)(unsafe.Pointer(param0))

	cValue1 := (C.GtkEventSequenceState)(param1)

	ret := C.gtk_gesture_set_sequence_state(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_gesture_set_state(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEventSequenceState)(param0)

	ret := C.gtk_gesture_set_state(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_gesture_set_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_gesture_set_window(cValueInstance, cValue0)
}

func Fn_gtk_gesture_ungroup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	C.gtk_gesture_ungroup(cValueInstance)
}

func Fn_gtk_gesture_drag_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_drag_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_drag_get_offset(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) bool {
	cValueInstance := (*C.GtkGestureDrag)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_gesture_drag_get_offset(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_gesture_drag_get_start_point(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) bool {
	cValueInstance := (*C.GtkGestureDrag)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_gesture_drag_get_start_point(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_gesture_long_press_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_long_press_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_multi_press_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_multi_press_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_multi_press_get_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGestureMultiPress)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_multi_press_get_area(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_gesture_multi_press_set_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkGestureMultiPress)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_gesture_multi_press_set_area(cValueInstance, cValue0)
}

func Fn_gtk_gesture_pan_new(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkOrientation)(param1)

	ret := C.gtk_gesture_pan_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_pan_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkGesturePan)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_pan_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_gesture_pan_set_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkGesturePan)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	C.gtk_gesture_pan_set_orientation(cValueInstance, cValue0)
}

func Fn_gtk_gesture_rotate_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_rotate_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_rotate_get_angle_delta(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkGestureRotate)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_rotate_get_angle_delta(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_gesture_single_get_button(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_single_get_button(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_gesture_single_get_current_button(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_single_get_current_button(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_gesture_single_get_current_sequence(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_single_get_current_sequence(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_single_get_exclusive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_single_get_exclusive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gesture_single_get_touch_only(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_single_get_touch_only(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_gesture_single_set_button(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_gesture_single_set_button(cValueInstance, cValue0)
}

func Fn_gtk_gesture_single_set_exclusive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_gesture_single_set_exclusive(cValueInstance, cValue0)
}

func Fn_gtk_gesture_single_set_touch_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGestureSingle)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_gesture_single_set_touch_only(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

func Fn_gtk_gesture_swipe_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_swipe_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_swipe_get_velocity(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) bool {
	cValueInstance := (*C.GtkGestureSwipe)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_gesture_swipe_get_velocity(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_gesture_zoom_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_zoom_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gesture_zoom_get_scale_delta(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkGestureZoom)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gesture_zoom_get_scale_delta(cValueInstance)

	return (float64)(ret)
}

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

func Fn_gtk_icon_theme_add_resource_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_add_resource_path(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

func Fn_gtk_list_box_get_selected_rows(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_get_selected_rows(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_box_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	C.gtk_list_box_select_all(cValueInstance)
}

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_list_box_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	C.gtk_list_box_unselect_all(cValueInstance)
}

func Fn_gtk_list_box_unselect_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkListBoxRow)(unsafe.Pointer(param0))

	C.gtk_list_box_unselect_row(cValueInstance, cValue0)
}

func Fn_gtk_list_box_row_get_activatable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_row_get_activatable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_list_box_row_get_selectable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_row_get_selectable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_list_box_row_is_selected(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_list_box_row_is_selected(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_list_box_row_set_activatable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_list_box_row_set_activatable(cValueInstance, cValue0)
}

func Fn_gtk_list_box_row_set_selectable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkListBoxRow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_list_box_row_set_selectable(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_places_sidebar_get_show_enter_location(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_show_enter_location(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_places_sidebar_set_show_enter_location(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_places_sidebar_set_show_enter_location(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

func Fn_gtk_switch_get_state(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSwitch)(unsafe.Pointer(paramInstance))

	ret := C.gtk_switch_get_state(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_switch_set_state(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSwitch)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_switch_set_state(cValueInstance, cValue0)
}

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

func Fn_gtk_widget_get_clip(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_get_clip(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_set_clip(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_set_clip(cValueInstance, cValue0)
}

func Fn_gtk_window_set_interactive_debugging(param0 bool) {
	cValue0 := toCBool(param0)

	C.gtk_window_set_interactive_debugging(cValue0)
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

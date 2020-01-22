// Code generated - DO NOT EDIT.
// +build gtk_2.12

package gtk

import (
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

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
func Fn_gtk_binding_entry_skip(param0 unsafe.Pointer, param1 uint, param2 int) {
	cValue0 := (*C.GtkBindingSet)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_binding_entry_skip(cValue0, cValue1, cValue2)
}

func Fn_gtk_paper_size_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_paper_size_new_from_key_file(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_paper_size_to_key_file(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_paper_size_get_paper_sizes(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.gtk_paper_size_get_paper_sizes(cValue0)

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

func Fn_gtk_rc_parse_color_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) uint {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRcStyle)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkColor)(unsafe.Pointer(param2))

	ret := C.gtk_rc_parse_color_full(cValue0, cValue1, cValue2)

	return (uint)(ret)
}

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_add_credit_section : parameter 'people' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_get_artists : no array length

// UNSUPPORTED : gtk_about_dialog_get_authors : no array length

// UNSUPPORTED : gtk_about_dialog_get_documenters : no array length

func Fn_gtk_about_dialog_get_program_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_program_name(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_about_dialog_set_artists : parameter 'artists' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_authors : parameter 'authors' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_documenters : parameter 'documenters' is array parameter without length parameter

func Fn_gtk_about_dialog_set_program_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_program_name(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

func Fn_gtk_action_create_menu(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_create_menu(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_builder_new() unsafe.Pointer {
	ret := C.gtk_builder_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

func Fn_gtk_builder_add_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_builder_add_from_file(cValueInstance, cValue0, cError)

	return (uint)(ret)
}

func Fn_gtk_builder_add_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gsize)(param1)

	cError := (**C.GError)(error)

	ret := C.gtk_builder_add_from_string(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

func Fn_gtk_builder_connect_signals(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.gtk_builder_connect_signals(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

func Fn_gtk_builder_get_object(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_builder_get_object(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_builder_get_objects(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	ret := C.gtk_builder_get_objects(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_builder_get_translation_domain(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	ret := C.gtk_builder_get_translation_domain(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_builder_get_type_from_name(paramInstance unsafe.Pointer, param0 string) uint64 {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_builder_get_type_from_name(cValueInstance, cValue0)

	return (uint64)(ret)
}

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_builder_set_translation_domain(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_builder_set_translation_domain(cValueInstance, cValue0)
}

func Fn_gtk_builder_value_from_string(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.gtk_builder_value_from_string(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_gtk_builder_value_from_string_type(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.gtk_builder_value_from_string_type(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_entry_get_cursor_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_cursor_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_set_cursor_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_entry_set_cursor_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_get_completion_prefix(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_completion_prefix(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_entry_completion_get_inline_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_inline_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_set_inline_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_inline_selection(cValueInstance, cValue0)
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

func Fn_gtk_icon_theme_list_contexts(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_theme_list_contexts(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_convert_widget_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_icon_view_convert_widget_to_bin_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_view_get_tooltip_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_tooltip_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := (**C.GtkTreeModel)(unsafe.Pointer(param3))

	cValue4 := (**C.GtkTreePath)(unsafe.Pointer(param4))

	cValue5 := (*C.GtkTreeIter)(unsafe.Pointer(param5))

	ret := C.gtk_icon_view_get_tooltip_context(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_icon_view_set_tooltip_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkCellRenderer)(unsafe.Pointer(param2))

	C.gtk_icon_view_set_tooltip_cell(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_view_set_tooltip_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_tooltip_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_tooltip_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_icon_view_set_tooltip_item(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_list_store_set_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(&param1[0]))

	cValue2 := (*C.GValue)(unsafe.Pointer(&param2[0]))

	cValue3 := (C.gint)(param3)

	C.gtk_list_store_set_valuesv(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_tool_button_set_arrow_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_menu_tool_button_set_arrow_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_tool_button_set_arrow_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_new_from_key_file(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_to_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_to_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_page_setup_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_page_setup_to_key_file(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_settings_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_new_from_key_file(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_print_settings_to_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_to_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_print_settings_to_key_file(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_range_get_fill_level(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_fill_level(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_range_get_restrict_to_fill_level(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_restrict_to_fill_level(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_get_show_fill_level(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_show_fill_level(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_set_fill_level(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_range_set_fill_level(cValueInstance, cValue0)
}

func Fn_gtk_range_set_restrict_to_fill_level(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_restrict_to_fill_level(cValueInstance, cValue0)
}

func Fn_gtk_range_set_show_fill_level(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_show_fill_level(cValueInstance, cValue0)
}

func Fn_gtk_recent_action_new(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.gtk_recent_action_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_action_new_for_manager(param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.GtkRecentManager)(unsafe.Pointer(param4))

	ret := C.gtk_recent_action_new_for_manager(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_action_get_show_numbers(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_action_get_show_numbers(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_action_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_action_set_show_numbers(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scale_button_get_adjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_adjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_value(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_scale_button_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scale_button_set_adjustment(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scale_button_set_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_scale_button_set_value(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_status_icon_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_status_icon_set_screen(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

func Fn_gtk_text_buffer_add_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_add_mark(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

func Fn_gtk_text_mark_new(param0 string, param1 bool) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	ret := C.gtk_text_mark_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

func Fn_gtk_tool_item_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_item_set_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_item_set_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_custom(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tooltip_set_custom(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_tooltip_set_icon(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_icon_from_stock(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_tooltip_set_icon_from_stock(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tooltip_set_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tooltip_set_markup(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tooltip_set_text(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_tip_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_tooltip_set_tip_area(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_trigger_tooltip_query(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	C.gtk_tooltip_trigger_tooltip_query(cValue0)
}

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_tree_store_set_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(&param1[0]))

	cValue2 := (*C.GValue)(unsafe.Pointer(&param2[0]))

	cValue3 := (C.gint)(param3)

	C.gtk_tree_store_set_valuesv(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_bin_window_to_tree_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_bin_window_to_tree_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_bin_window_to_widget_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_bin_window_to_widget_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_tree_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_tree_to_bin_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_tree_to_widget_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_tree_to_widget_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_widget_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_widget_to_bin_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_widget_to_tree_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_widget_to_tree_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_get_level_indentation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_level_indentation(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

func Fn_gtk_tree_view_get_show_expanders(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_show_expanders(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_tooltip_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_tooltip_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := (**C.GtkTreeModel)(unsafe.Pointer(param3))

	cValue4 := (**C.GtkTreePath)(unsafe.Pointer(param4))

	cValue5 := (*C.GtkTreeIter)(unsafe.Pointer(param5))

	ret := C.gtk_tree_view_get_tooltip_context(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

func Fn_gtk_tree_view_is_rubber_banding_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_is_rubber_banding_active(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_level_indentation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_set_level_indentation(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_show_expanders(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_show_expanders(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_tooltip_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkCellRenderer)(unsafe.Pointer(param3))

	C.gtk_tree_view_set_tooltip_cell(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_set_tooltip_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_set_tooltip_column(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_tooltip_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_view_set_tooltip_row(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_column_get_tree_view(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_tree_view(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

func Fn_gtk_volume_button_new() unsafe.Pointer {
	ret := C.gtk_volume_button_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_error_bell(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_error_bell(cValueInstance)
}

func Fn_gtk_widget_get_has_tooltip(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_has_tooltip(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_tooltip_markup(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_tooltip_markup(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_tooltip_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_tooltip_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_tooltip_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_tooltip_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_keynav_failed(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkDirectionType)(param0)

	ret := C.gtk_widget_keynav_failed(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_modify_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_widget_modify_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_has_tooltip(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_has_tooltip(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_set_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_set_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_tooltip_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_widget_set_tooltip_window(cValueInstance, cValue0)
}

func Fn_gtk_widget_trigger_tooltip_query(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_trigger_tooltip_query(cValueInstance)
}

func Fn_gtk_window_get_opacity(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_opacity(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_window_set_opacity(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_window_set_opacity(cValueInstance, cValue0)
}

func Fn_gtk_window_set_startup_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_startup_id(cValueInstance, cValue0)
}

func Fn_gtk_buildable_add_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_buildable_add_child(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_buildable_construct_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_buildable_construct_child(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_buildable_custom_finished(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gpointer)(param3)

	C.gtk_buildable_custom_finished(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_buildable_custom_tag_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 *unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gpointer)(unsafe.Pointer(param3))

	C.gtk_buildable_custom_tag_end(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_buildable_custom_tag_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GMarkupParser)(unsafe.Pointer(param3))

	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))

	ret := C.gtk_buildable_custom_tag_start(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_buildable_get_internal_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_buildable_get_internal_child(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_buildable_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_buildable_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_buildable_parser_finished(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	C.gtk_buildable_parser_finished(cValueInstance, cValue0)
}

func Fn_gtk_buildable_set_buildable_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_buildable_set_buildable_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_buildable_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_buildable_set_name(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

func Fn_gtk_cell_layout_get_cells(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_layout_get_cells(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

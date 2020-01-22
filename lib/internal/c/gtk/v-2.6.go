// Code generated - DO NOT EDIT.
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import (
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static void c_gtk_show_about_dialog(GtkWindow* parent, const gchar* first_property_name) {
    return gtk_show_about_dialog(parent, first_property_name, NULL);
}
*/
/*

static void c_gtk_dialog_set_alternative_button_order(GtkDialog* dialog, gint first_response_id) {
    return gtk_dialog_set_alternative_button_order(dialog, first_response_id, NULL);
}
*/
/*

static void c_gtk_list_store_insert_with_values(GtkListStore* list_store, GtkTreeIter* iter, gint position) {
    return gtk_list_store_insert_with_values(list_store, iter, position, NULL);
}
*/
/*

static void c_gtk_message_dialog_format_secondary_markup(GtkMessageDialog* message_dialog, const gchar* message_format) {
    return gtk_message_dialog_format_secondary_markup(message_dialog, message_format, NULL);
}
*/
/*

static void c_gtk_message_dialog_format_secondary_text(GtkMessageDialog* message_dialog, const gchar* message_format) {
    return gtk_message_dialog_format_secondary_text(message_dialog, message_format, NULL);
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

func Fn_gtk_selection_data_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

func Fn_gtk_selection_data_set_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_selection_data_set_pixbuf(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

func Fn_gtk_selection_data_targets_include_image(paramInstance unsafe.Pointer, param0 bool) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gtk_selection_data_targets_include_image(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_target_list_add_image_targets(paramInstance unsafe.Pointer, param0 uint, param1 bool) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := toCBool(param1)

	C.gtk_target_list_add_image_targets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_target_list_add_text_targets(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_target_list_add_text_targets(cValueInstance, cValue0)
}

func Fn_gtk_target_list_add_uri_targets(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_target_list_add_uri_targets(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_accelerator_get_label(param0 uint, param1 int) string {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	ret := C.gtk_accelerator_get_label(cValue0, cValue1)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func Fn_gtk_alternative_dialog_button_order(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_alternative_dialog_button_order(cValue0)

	return toGoBool(ret)
}

func Fn_gtk_get_option_group(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.gtk_get_option_group(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_show_about_dialog(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_show_about_dialog(cValue0, cValue1)
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_about_dialog_new() unsafe.Pointer {
	ret := C.gtk_about_dialog_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_about_dialog_add_credit_section : parameter 'people' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_get_artists : no array length

// UNSUPPORTED : gtk_about_dialog_get_authors : no array length

func Fn_gtk_about_dialog_get_comments(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_comments(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_copyright(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_copyright(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_about_dialog_get_documenters : no array length

func Fn_gtk_about_dialog_get_license(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_license(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_logo(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_logo(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_about_dialog_get_logo_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_logo_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_translator_credits(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_translator_credits(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_version(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_version(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_website(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_website(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_website_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_website_label(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_about_dialog_set_artists : parameter 'artists' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_authors : parameter 'authors' is array parameter without length parameter

func Fn_gtk_about_dialog_set_comments(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_comments(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_copyright(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_copyright(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_about_dialog_set_documenters : parameter 'documenters' is array parameter without length parameter

func Fn_gtk_about_dialog_set_license(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_license(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_logo(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_about_dialog_set_logo(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_logo_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_logo_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_translator_credits(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_translator_credits(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_version(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_version(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_website(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_website(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_website_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_website_label(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

func Fn_gtk_action_get_accel_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_accel_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_action_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_visible(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

func Fn_gtk_action_group_translate_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_action_group_translate_string(cValueInstance, cValue0)

	return C.GoString(ret)
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

func Fn_gtk_button_get_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_image(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_button_set_image(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

// UNSUPPORTED : gtk_cell_area_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : parameter 'callback' is callback

func Fn_gtk_cell_renderer_stop_editing(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_stop_editing(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_combo_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_combo_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_progress_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_progress_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new() unsafe.Pointer {
	ret := C.gtk_cell_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_context(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellAreaContext)(unsafe.Pointer(param1))

	ret := C.gtk_cell_view_new_with_context(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_markup(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_cell_view_new_with_markup(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_pixbuf(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_cell_view_new_with_pixbuf(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_text(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_cell_view_new_with_text(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_get_displayed_row(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_displayed_row(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_get_size_of_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRequisition)(unsafe.Pointer(param1))

	ret := C.gtk_cell_view_get_size_of_row(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_cell_view_set_background_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_background_color(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_displayed_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_displayed_row(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_model(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

func Fn_gtk_clipboard_set_can_store(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	C.gtk_clipboard_set_can_store(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_clipboard_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_clipboard_set_image(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

func Fn_gtk_clipboard_store(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_store(cValueInstance)
}

func Fn_gtk_clipboard_wait_for_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_for_image(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

func Fn_gtk_clipboard_wait_is_image_available(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_is_image_available(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_clipboard_wait_is_target_available(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(param0)

	ret := C.gtk_clipboard_wait_is_target_available(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

func Fn_gtk_combo_box_get_column_span_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_column_span_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_focus_on_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_focus_on_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_combo_box_get_popup_accessible(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_popup_accessible(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

func Fn_gtk_combo_box_get_row_span_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_row_span_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_wrap_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_wrap_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_set_add_tearoffs(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_combo_box_set_add_tearoffs(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_combo_box_set_focus_on_click(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_dialog_set_alternative_button_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.c_gtk_dialog_set_alternative_button_order(cValueInstance, cValue0)
}

func Fn_gtk_dialog_set_alternative_button_order_from_array(paramInstance unsafe.Pointer, param0 int, param1 []int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(&param1[0]))

	C.gtk_dialog_set_alternative_button_order_from_array(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_completion_get_inline_completion(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_inline_completion(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_popup_completion(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_popup_completion(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_text_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_text_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_completion_insert_prefix(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_insert_prefix(cValueInstance)
}

func Fn_gtk_entry_completion_set_inline_completion(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_inline_completion(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

func Fn_gtk_entry_completion_set_popup_completion(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_popup_completion(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

func Fn_gtk_file_chooser_button_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkFileChooserAction)(param1)

	ret := C.gtk_file_chooser_button_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_button_new_with_dialog(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_file_chooser_button_new_with_dialog(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_button_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_button_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_button_get_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_button_get_width_chars(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_file_chooser_button_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_button_set_title(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_button_set_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_file_chooser_button_set_width_chars(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

func Fn_gtk_file_filter_add_pixbuf_formats(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	C.gtk_file_filter_add_pixbuf_formats(cValueInstance)
}

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

func Fn_gtk_icon_view_new() unsafe.Pointer {
	ret := C.gtk_icon_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_new_with_model(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_new_with_model(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_get_column_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_column_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_columns(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_columns(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_item_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_item_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_item_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_item_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_margin(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_margin(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_markup_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_markup_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_get_path_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_icon_view_get_path_at_pos(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_get_pixbuf_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_pixbuf_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_row_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_row_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_selected_items(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_selected_items(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_get_selection_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_selection_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_text_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_text_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_item_activated(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_item_activated(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_path_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_path_is_selected(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_select_all(cValueInstance)
}

func Fn_gtk_icon_view_select_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_select_path(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_icon_view_set_column_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_column_spacing(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_columns(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_columns(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_item_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	C.gtk_icon_view_set_item_orientation(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_item_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_item_width(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_margin(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_markup_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_markup_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_icon_view_set_model(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_pixbuf_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_pixbuf_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_row_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_row_spacing(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_selection_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSelectionMode)(param0)

	C.gtk_icon_view_set_selection_mode(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_text_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_text_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_unselect_all(cValueInstance)
}

func Fn_gtk_icon_view_unselect_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_unselect_path(cValueInstance, cValue0)
}

func Fn_gtk_image_new_from_icon_name(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_image_new_from_icon_name(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_get_icon_name(paramInstance unsafe.Pointer, param0 *string, param1 *int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (*C.GtkIconSize)(unsafe.Pointer(param1))

	C.gtk_image_get_icon_name(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String
}

func Fn_gtk_image_get_pixel_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_get_pixel_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_image_set_from_icon_name(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_set_from_icon_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_set_pixel_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_image_set_pixel_size(cValueInstance, cValue0)
}

func Fn_gtk_label_get_angle(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_angle(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_label_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_get_max_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_max_width_chars(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_get_single_line_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_single_line_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_get_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_width_chars(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_set_angle(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_label_set_angle(cValueInstance, cValue0)
}

func Fn_gtk_label_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.gtk_label_set_ellipsize(cValueInstance, cValue0)
}

func Fn_gtk_label_set_max_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_label_set_max_width_chars(cValueInstance, cValue0)
}

func Fn_gtk_label_set_single_line_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_single_line_mode(cValueInstance, cValue0)
}

func Fn_gtk_label_set_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_label_set_width_chars(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_list_store_insert_with_values(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.c_gtk_list_store_insert_with_values(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_list_store_insert_with_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 []int, param3 []gobject.Value, param4 int) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(&param2[0]))

	cValue3 := (*C.GValue)(unsafe.Pointer(&param3[0]))

	cValue4 := (C.gint)(param4)

	C.gtk_list_store_insert_with_valuesv(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_get_for_attach_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_menu_get_for_attach_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_tool_button_new(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_menu_tool_button_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_tool_button_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_menu_tool_button_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_tool_button_get_menu(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_tool_button_get_menu(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_tool_button_set_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_tool_button_set_menu(cValueInstance, cValue0)
}

func Fn_gtk_message_dialog_format_secondary_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.c_gtk_message_dialog_format_secondary_markup(cValueInstance, cValue0)
}

func Fn_gtk_message_dialog_format_secondary_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.c_gtk_message_dialog_format_secondary_text(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_progress_bar_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_progress_bar_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.gtk_progress_bar_set_ellipsize(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

func Fn_gtk_text_buffer_backspace(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	ret := C.gtk_text_buffer_backspace(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_view_get_iter_at_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 int, param3 int) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	ret := C.gtk_text_view_get_iter_at_position(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

func Fn_gtk_tool_item_rebuild_menu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_rebuild_menu(cValueInstance)
}

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_tree_view_get_fixed_height_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_fixed_height_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_hover_expand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_hover_expand(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_hover_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_hover_selection(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_fixed_height_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_fixed_height_mode(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_hover_expand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_hover_expand(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_hover_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_hover_selection(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_drag_dest_add_image_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_add_image_targets(cValueInstance)
}

func Fn_gtk_drag_dest_add_text_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_add_text_targets(cValueInstance)
}

func Fn_gtk_drag_dest_add_uri_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_add_uri_targets(cValueInstance)
}

func Fn_gtk_drag_source_add_image_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_add_image_targets(cValueInstance)
}

func Fn_gtk_drag_source_add_text_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_add_text_targets(cValueInstance)
}

func Fn_gtk_drag_source_add_uri_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_add_uri_targets(cValueInstance)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_get_focus_on_map(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_focus_on_map(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_window_set_focus_on_map(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_focus_on_map(cValueInstance, cValue0)
}

func Fn_gtk_window_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_window_set_default_icon_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_default_icon_name(cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_file_chooser_get_show_hidden(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_show_hidden(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_show_hidden(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_show_hidden(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

// Code generated - DO NOT EDIT.
// +build gtk_2.4

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static GtkWidget* c_gtk_file_chooser_dialog_new(const gchar* title, GtkWindow* parent, GtkFileChooserAction action, const gchar* first_button_text) {
    return gtk_file_chooser_dialog_new(title, parent, action, first_button_text, NULL);
}
*/
/*

static GtkWidget* c_gtk_message_dialog_new_with_markup(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, const gchar* message_format) {
    return gtk_message_dialog_new_with_markup(parent, flags, type, buttons, message_format, NULL);
}
*/
/*

static void c_gtk_cell_layout_set_attributes(GtkCellLayout* cell_layout, GtkCellRenderer* cell) {
    return gtk_cell_layout_set_attributes(cell_layout, cell, NULL);
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

func Fn_gtk_text_iter_backward_visible_cursor_position(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_visible_cursor_position(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_cursor_positions(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_visible_cursor_positions(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_word_start(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_visible_word_start(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_word_starts(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_visible_word_starts(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

func Fn_gtk_text_iter_forward_visible_cursor_position(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_visible_cursor_position(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_cursor_positions(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_visible_cursor_positions(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_word_end(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_visible_word_end(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_word_ends(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_visible_word_ends(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func Fn_gtk_bindings_activate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEventKey)(unsafe.Pointer(param1))

	ret := C.gtk_bindings_activate_event(cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

func Fn_gtk_rc_reset_styles(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	C.gtk_rc_reset_styles(cValue0)
}

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

func Fn_gtk_accel_map_get() unsafe.Pointer {
	ret := C.gtk_accel_map_get()

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_map_lock_path(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_map_lock_path(cValue0)
}

func Fn_gtk_accel_map_unlock_path(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_map_unlock_path(cValue0)
}

func Fn_gtk_action_new(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.gtk_action_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_activate(cValueInstance)
}

func Fn_gtk_action_connect_accelerator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_connect_accelerator(cValueInstance)
}

func Fn_gtk_action_create_icon(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkIconSize)(param0)

	ret := C.gtk_action_create_icon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_create_menu_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_create_menu_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_create_tool_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_create_tool_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_disconnect_accelerator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_disconnect_accelerator(cValueInstance)
}

func Fn_gtk_action_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_proxies(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_proxies(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_is_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_is_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_is_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_is_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_action_set_accel_group(cValueInstance, cValue0)
}

func Fn_gtk_action_set_accel_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_accel_path(cValueInstance, cValue0)
}

func Fn_gtk_action_group_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_action_group_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_group_add_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_action_group_add_action(cValueInstance, cValue0)
}

func Fn_gtk_action_group_add_action_with_accel(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_action_group_add_action_with_accel(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_action_group_add_actions(paramInstance unsafe.Pointer, param0 []ActionEntry, param1 uint, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkActionEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.gpointer)(param2)

	C.gtk_action_group_add_actions(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

func Fn_gtk_action_group_add_toggle_actions(paramInstance unsafe.Pointer, param0 []ToggleActionEntry, param1 uint, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToggleActionEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.gpointer)(param2)

	C.gtk_action_group_add_toggle_actions(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

func Fn_gtk_action_group_get_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_action_group_get_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_group_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_group_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_get_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_group_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_group_list_actions(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_list_actions(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_group_remove_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_action_group_remove_action(cValueInstance, cValue0)
}

func Fn_gtk_action_group_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_group_set_sensitive(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

func Fn_gtk_action_group_set_translation_domain(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_group_set_translation_domain(cValueInstance, cValue0)
}

func Fn_gtk_action_group_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_group_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_alignment_get_padding(paramInstance unsafe.Pointer, param0 *uint, param1 *uint, param2 *uint, param3 *uint) {
	cValueInstance := (*C.GtkAlignment)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	cValue3 := (*C.guint)(unsafe.Pointer(param3))

	C.gtk_alignment_get_padding(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_alignment_set_padding(paramInstance unsafe.Pointer, param0 uint, param1 uint, param2 uint, param3 uint) {
	cValueInstance := (*C.GtkAlignment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.guint)(param3)

	C.gtk_alignment_set_padding(cValueInstance, cValue0, cValue1, cValue2, cValue3)
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

func Fn_gtk_button_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gfloat)(unsafe.Pointer(param0))

	cValue1 := (*C.gfloat)(unsafe.Pointer(param1))

	C.gtk_button_get_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_button_get_focus_on_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_focus_on_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_button_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	C.gtk_button_set_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_button_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_button_set_focus_on_click(cValueInstance, cValue0)
}

func Fn_gtk_button_box_get_child_secondary(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_button_box_get_child_secondary(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_calendar_get_display_options(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_calendar_get_display_options(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

func Fn_gtk_calendar_set_display_options(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkCalendarDisplayOptions)(param0)

	C.gtk_calendar_set_display_options(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_area_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : parameter 'callback' is callback

func Fn_gtk_check_menu_item_get_draw_as_radio(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_check_menu_item_get_draw_as_radio(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_check_menu_item_set_draw_as_radio(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_check_menu_item_set_draw_as_radio(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

func Fn_gtk_clipboard_wait_for_targets(paramInstance unsafe.Pointer, param0 *[]unsafe.Pointer, param1 *int) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer (*C.GdkAtom)
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gtk_clipboard_wait_for_targets(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]unsafe.Pointer, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0Out = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
	}
	*param0 = param0Out

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

func Fn_gtk_color_button_new() unsafe.Pointer {
	ret := C.gtk_color_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_button_new_with_color(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	ret := C.gtk_color_button_new_with_color(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_button_get_alpha(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_button_get_alpha(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_color_button_get_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_button_get_color(cValueInstance, cValue0)
}

func Fn_gtk_color_button_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_button_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_color_button_get_use_alpha(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_button_get_use_alpha(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_color_button_set_alpha(paramInstance unsafe.Pointer, param0 uint16) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint16)(param0)

	C.gtk_color_button_set_alpha(cValueInstance, cValue0)
}

func Fn_gtk_color_button_set_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_button_set_color(cValueInstance, cValue0)
}

func Fn_gtk_color_button_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_color_button_set_title(cValueInstance, cValue0)
}

func Fn_gtk_color_button_set_use_alpha(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_color_button_set_use_alpha(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

func Fn_gtk_combo_box_new() unsafe.Pointer {
	ret := C.gtk_combo_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_model(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_combo_box_new_with_model(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_get_active(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_active(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_active_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_combo_box_get_active_iter(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_combo_box_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

func Fn_gtk_combo_box_popdown(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_popdown(cValueInstance)
}

func Fn_gtk_combo_box_popup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_popup(cValueInstance)
}

func Fn_gtk_combo_box_set_active(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_active(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_active_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_combo_box_set_active_iter(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_column_span_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_column_span_column(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_combo_box_set_model(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

func Fn_gtk_combo_box_set_row_span_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_row_span_column(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_wrap_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_wrap_width(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_entry_get_alignment(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_alignment(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_entry_get_completion(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_completion(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_set_alignment(paramInstance unsafe.Pointer, param0 float32) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	C.gtk_entry_set_alignment(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_completion(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkEntryCompletion)(unsafe.Pointer(param0))

	C.gtk_entry_set_completion(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_new() unsafe.Pointer {
	ret := C.gtk_entry_completion_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_completion_complete(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_complete(cValueInstance)
}

func Fn_gtk_entry_completion_delete_action(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_completion_delete_action(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_get_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_completion_get_minimum_key_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_minimum_key_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_completion_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_completion_insert_action_markup(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_completion_insert_action_markup(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_completion_insert_action_text(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_completion_insert_action_text(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

func Fn_gtk_entry_completion_set_minimum_key_length(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_completion_set_minimum_key_length(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_entry_completion_set_model(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_text_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_completion_set_text_column(cValueInstance, cValue0)
}

func Fn_gtk_event_box_get_above_child(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_event_box_get_above_child(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_event_box_get_visible_window(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_event_box_get_visible_window(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_event_box_set_above_child(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_event_box_set_above_child(cValueInstance, cValue0)
}

func Fn_gtk_event_box_set_visible_window(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_event_box_set_visible_window(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

func Fn_gtk_expander_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_expander_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_expander_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_expander_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_expander_get_expanded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_expanded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_expander_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_expander_get_label_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_label_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_expander_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_expander_get_use_markup(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_use_markup(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_expander_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_use_underline(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_expander_set_expanded(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_expanded(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_expander_set_label(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_expander_set_label_widget(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_expander_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_use_markup(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_use_markup(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_dialog_new(param0 string, param1 unsafe.Pointer, param2 int, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	cValue2 := (C.GtkFileChooserAction)(param2)

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.c_gtk_file_chooser_dialog_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_widget_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkFileChooserAction)(param0)

	ret := C.gtk_file_chooser_widget_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_filter_new() unsafe.Pointer {
	ret := C.gtk_file_filter_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

func Fn_gtk_file_filter_add_mime_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_filter_add_mime_type(cValueInstance, cValue0)
}

func Fn_gtk_file_filter_add_pattern(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_filter_add_pattern(cValueInstance, cValue0)
}

func Fn_gtk_file_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilterInfo)(unsafe.Pointer(param0))

	ret := C.gtk_file_filter_filter(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_filter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_filter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_filter_get_needed(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_filter_get_needed(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_file_filter_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_filter_set_name(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_font_button_new() unsafe.Pointer {
	ret := C.gtk_font_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_button_new_with_font(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_font_button_new_with_font(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_button_get_font_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_font_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_button_get_show_size(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_show_size(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_get_show_style(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_show_style(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_button_get_use_font(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_use_font(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_get_use_size(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_use_size(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_set_font_name(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_font_button_set_font_name(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_font_button_set_show_size(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_show_size(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_show_style(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_show_style(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_button_set_title(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_use_font(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_use_font(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_use_size(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_use_size(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

func Fn_gtk_icon_info_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_info_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_free(cValueInstance)
}

func Fn_gtk_icon_info_get_attach_points(paramInstance unsafe.Pointer, param0 *[]unsafe.Pointer, param1 *int) bool {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer (*C.GdkPoint)
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gtk_icon_info_get_attach_points(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]unsafe.Pointer, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0Out = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
	}
	*param0 = param0Out

	return toGoBool(ret)
}

func Fn_gtk_icon_info_get_base_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_base_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_info_get_builtin_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_builtin_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_info_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_icon_info_get_embedded_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gtk_icon_info_get_embedded_rect(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_icon_info_get_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_icon_info_load_icon(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_icon(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

func Fn_gtk_icon_info_set_raw_coordinates(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_info_set_raw_coordinates(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_new() unsafe.Pointer {
	ret := C.gtk_icon_theme_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_append_search_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_append_search_path(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

func Fn_gtk_icon_theme_get_example_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_theme_get_example_icon_name(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

func Fn_gtk_icon_theme_get_search_path(paramInstance unsafe.Pointer, param0 *[]string, param1 *int) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer **C.gchar
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_icon_theme_get_search_path(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]string, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
		for param0Outi, param0OutCString := range param0OutCSlice {
			param0Out[param0Outi] = C.GoString(param0OutCString)
		}
	}
	*param0 = param0Out
}

func Fn_gtk_icon_theme_has_icon(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_icon_theme_has_icon(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_icon_theme_list_icons(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_icon_theme_list_icons(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_load_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	cError := (**C.GError)(error)

	ret := C.gtk_icon_theme_load_icon(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_lookup_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	ret := C.gtk_icon_theme_lookup_icon(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_prepend_search_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_prepend_search_path(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_rescan_if_needed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_theme_rescan_if_needed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_icon_theme_set_custom_theme(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_set_custom_theme(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_icon_theme_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_set_search_path(paramInstance unsafe.Pointer, param0 []string, param1 int) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	cValue1 := (C.gint)(param1)

	C.gtk_icon_theme_set_search_path(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_theme_add_builtin_icon(param0 string, param1 int, param2 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GdkPixbuf)(unsafe.Pointer(param2))

	C.gtk_icon_theme_add_builtin_icon(cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_theme_get_default() unsafe.Pointer {
	ret := C.gtk_icon_theme_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_get_for_screen(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_icon_theme_get_for_screen(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_menu_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.guint)(param3)

	cValue4 := (C.guint)(param4)

	C.gtk_menu_attach(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_set_monitor(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_menu_set_monitor(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	C.gtk_menu_shell_cancel(cValueInstance)
}

func Fn_gtk_message_dialog_new_with_markup(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 string) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkDialogFlags)(param1)

	cValue2 := (C.GtkMessageType)(param2)

	cValue3 := (C.GtkButtonsType)(param3)

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	ret := C.c_gtk_message_dialog_new_with_markup(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_message_dialog_set_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_message_dialog_set_markup(cValueInstance, cValue0)
}

func Fn_gtk_paned_get_child1(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_child1(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paned_get_child2(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_child2(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_radio_action_new(param0 string, param1 string, param2 string, param3 string, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.gint)(param4)

	ret := C.gtk_radio_action_new(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_action_get_current_value(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_action_get_current_value(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_radio_action_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_action_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_action_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_action_set_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_menu_item_new_from_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	ret := C.gtk_radio_menu_item_new_from_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_menu_item_new_with_label_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	ret := C.gtk_radio_tool_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_new_from_stock(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_tool_button_new_from_stock(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_new_from_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRadioToolButton)(unsafe.Pointer(param0))

	ret := C.gtk_radio_tool_button_new_from_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_new_with_stock_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioToolButton)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_tool_button_new_with_stock_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRadioToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_tool_button_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_tool_button_set_group(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

func Fn_gtk_scale_get_layout(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_get_layout(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_scale_get_layout_offsets(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_separator_tool_item_new() unsafe.Pointer {
	ret := C.gtk_separator_tool_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_separator_tool_item_get_draw(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSeparatorToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_separator_tool_item_get_draw(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_separator_tool_item_set_draw(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSeparatorToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_separator_tool_item_set_draw(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

func Fn_gtk_text_buffer_select_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_select_range(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_view_get_accepts_tab(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_accepts_tab(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_overwrite(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_overwrite(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_set_accepts_tab(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_view_set_accepts_tab(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_overwrite(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_view_set_overwrite(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

func Fn_gtk_toggle_action_new(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.gtk_toggle_action_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_action_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_action_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toggle_action_get_draw_as_radio(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_action_get_draw_as_radio(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toggle_action_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_action_set_active(cValueInstance, cValue0)
}

func Fn_gtk_toggle_action_set_draw_as_radio(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_action_set_draw_as_radio(cValueInstance, cValue0)
}

func Fn_gtk_toggle_action_toggled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_action_toggled(cValueInstance)
}

func Fn_gtk_toggle_tool_button_new() unsafe.Pointer {
	ret := C.gtk_toggle_tool_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_tool_button_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_toggle_tool_button_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_tool_button_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_tool_button_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toggle_tool_button_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_tool_button_set_active(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_new(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_tool_button_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_tool_button_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_get_icon_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_icon_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_button_get_label_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_label_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_get_stock_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_stock_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_button_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_use_underline(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_button_set_icon_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tool_button_set_icon_widget(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_button_set_label(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tool_button_set_label_widget(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_stock_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_button_set_stock_id(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_button_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_new() unsafe.Pointer {
	ret := C.gtk_tool_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_get_expand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_expand(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_icon_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_icon_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_is_important(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_is_important(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_proxy_menu_item(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_tool_item_get_proxy_menu_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_get_relief_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_relief_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_toolbar_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_toolbar_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_use_drag_window(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_use_drag_window(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_visible_horizontal(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_visible_horizontal(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_visible_vertical(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_visible_vertical(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_retrieve_proxy_menu_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_retrieve_proxy_menu_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_set_expand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_expand(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_is_important(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_is_important(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_proxy_menu_item(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_tool_item_set_proxy_menu_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_item_set_use_drag_window(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_use_drag_window(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_visible_horizontal(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_visible_horizontal(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_visible_vertical(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_visible_vertical(cValueInstance, cValue0)
}

func Fn_gtk_toolbar_get_drop_index(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_toolbar_get_drop_index(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_item_index(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	ret := C.gtk_toolbar_get_item_index(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_n_items(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_n_items(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_nth_item(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_toolbar_get_nth_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toolbar_get_relief_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_relief_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_show_arrow(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_show_arrow(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toolbar_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_toolbar_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_toolbar_set_drop_highlight_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_toolbar_set_drop_highlight_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_toolbar_set_show_arrow(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toolbar_set_show_arrow(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_filter_clear_cache(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_filter_clear_cache(cValueInstance)
}

func Fn_gtk_tree_model_filter_convert_child_iter_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	ret := C.gtk_tree_model_filter_convert_child_iter_to_iter(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_filter_convert_child_path_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_filter_convert_child_path_to_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_filter_convert_iter_to_child_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_filter_convert_iter_to_child_iter(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_filter_convert_path_to_child_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_filter_convert_path_to_child_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_filter_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_model_filter_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_filter_refilter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_filter_refilter(cValueInstance)
}

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

func Fn_gtk_tree_model_filter_set_visible_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_model_filter_set_visible_column(cValueInstance, cValue0)
}

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

func Fn_gtk_tree_view_column_get_expand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_expand(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

func Fn_gtk_tree_view_column_set_expand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_column_set_expand(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_new() unsafe.Pointer {
	ret := C.gtk_ui_manager_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_add_ui(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 string, param3 string, param4 int, param5 bool) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.GtkUIManagerItemType)(param4)

	cValue5 := toCBool(param5)

	C.gtk_ui_manager_add_ui(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_ui_manager_add_ui_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_ui_manager_add_ui_from_file(cValueInstance, cValue0, cError)

	return (uint)(ret)
}

func Fn_gtk_ui_manager_add_ui_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gssize)(param1)

	cError := (**C.GError)(error)

	ret := C.gtk_ui_manager_add_ui_from_string(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

func Fn_gtk_ui_manager_ensure_update(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_ensure_update(cValueInstance)
}

func Fn_gtk_ui_manager_get_accel_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_accel_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_ui_manager_get_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_action_groups(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_action_groups(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_add_tearoffs(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_add_tearoffs(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_ui_manager_get_toplevels(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUIManagerItemType)(param0)

	ret := C.gtk_ui_manager_get_toplevels(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_ui(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_ui(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_ui_manager_get_widget(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_ui_manager_get_widget(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_insert_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkActionGroup)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_ui_manager_insert_action_group(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_ui_manager_new_merge_id(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_new_merge_id(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_ui_manager_remove_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkActionGroup)(unsafe.Pointer(param0))

	C.gtk_ui_manager_remove_action_group(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_remove_ui(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_ui_manager_remove_ui(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_set_add_tearoffs(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_ui_manager_set_add_tearoffs(cValueInstance, cValue0)
}

func Fn_gtk_widget_add_mnemonic_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_add_mnemonic_label(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_can_activate_accel(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_widget_can_activate_accel(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_drag_source_get_target_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_drag_source_get_target_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_drag_source_set_target_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	C.gtk_drag_source_set_target_list(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_no_show_all(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_no_show_all(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_list_mnemonic_labels(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_list_mnemonic_labels(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_queue_resize_no_redraw(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_queue_resize_no_redraw(cValueInstance)
}

func Fn_gtk_widget_remove_mnemonic_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_remove_mnemonic_label(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_no_show_all(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_no_show_all(cValueInstance, cValue0)
}

func Fn_gtk_window_activate_key(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	ret := C.gtk_window_activate_key(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_window_get_accept_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_accept_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_has_toplevel_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_has_toplevel_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_propagate_key_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	ret := C.gtk_window_propagate_key_event(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_window_set_accept_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_accept_focus(cValueInstance, cValue0)
}

func Fn_gtk_window_set_keep_above(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_keep_above(cValueInstance, cValue0)
}

func Fn_gtk_window_set_keep_below(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_keep_below(cValueInstance, cValue0)
}

func Fn_gtk_window_set_default_icon(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_window_set_default_icon(cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

func Fn_gtk_cell_layout_add_attribute(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_cell_layout_add_attribute(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_layout_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	C.gtk_cell_layout_clear(cValueInstance)
}

func Fn_gtk_cell_layout_clear_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_layout_clear_attributes(cValueInstance, cValue0)
}

func Fn_gtk_cell_layout_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_cell_layout_pack_end(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_layout_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_cell_layout_pack_start(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_layout_reorder(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_cell_layout_reorder(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_layout_set_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.c_gtk_cell_layout_set_attributes(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_file_chooser_add_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilter)(unsafe.Pointer(param0))

	C.gtk_file_chooser_add_filter(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_add_shortcut_folder(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_add_shortcut_folder(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_add_shortcut_folder_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_add_shortcut_folder_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_action(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_action(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_file_chooser_get_current_folder(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_current_folder(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_current_folder_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_current_folder_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_extra_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_extra_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_filenames(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_filenames(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_filter(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_filter(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_local_only(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_local_only(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_preview_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_preview_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_preview_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_preview_widget_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_widget_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_select_multiple(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_select_multiple(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_uris(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_uris(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_list_filters(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_list_filters(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_list_shortcut_folder_uris(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_list_shortcut_folder_uris(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_list_shortcut_folders(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_list_shortcut_folders(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_remove_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilter)(unsafe.Pointer(param0))

	C.gtk_file_chooser_remove_filter(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_remove_shortcut_folder(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_remove_shortcut_folder(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_remove_shortcut_folder_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_remove_shortcut_folder_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	C.gtk_file_chooser_select_all(cValueInstance)
}

func Fn_gtk_file_chooser_select_filename(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_select_filename(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_select_uri(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_select_uri(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_action(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkFileChooserAction)(param0)

	C.gtk_file_chooser_set_action(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_current_folder(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_current_folder(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_current_folder_uri(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_current_folder_uri(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_current_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_set_current_name(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_extra_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_file_chooser_set_extra_widget(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_filename(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_filename(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilter)(unsafe.Pointer(param0))

	C.gtk_file_chooser_set_filter(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_local_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_local_only(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_preview_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_file_chooser_set_preview_widget(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_preview_widget_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_preview_widget_active(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_select_multiple(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_select_multiple(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_uri(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_uri(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_use_preview_label(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_use_preview_label(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	C.gtk_file_chooser_unselect_all(cValueInstance)
}

func Fn_gtk_file_chooser_unselect_filename(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_unselect_filename(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_unselect_uri(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_unselect_uri(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_tree_model_filter_new(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_filter_new(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

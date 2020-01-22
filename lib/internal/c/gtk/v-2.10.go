// Code generated - DO NOT EDIT.
// +build gtk_2.10

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static GtkWidget* c_gtk_recent_chooser_dialog_new(const gchar* title, GtkWindow* parent, const gchar* first_button_text) {
    return gtk_recent_chooser_dialog_new(title, parent, first_button_text, NULL);
}
*/
/*

static GtkWidget* c_gtk_recent_chooser_dialog_new_for_manager(const gchar* title, GtkWindow* parent, GtkRecentManager* manager, const gchar* first_button_text) {
    return gtk_recent_chooser_dialog_new_for_manager(title, parent, manager, first_button_text, NULL);
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
type RecentInfo C.GtkRecentInfo
type RecentManagerClass C.GtkRecentManagerClass

// UNSUPPORTED : StackAccessibleClass : blacklisted
func Fn_gtk_paper_size_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_paper_size_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_new_custom(param0 string, param1 string, param2 float64, param3 float64, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.GtkUnit)(param4)

	ret := C.gtk_paper_size_new_custom(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_new_from_ppd(param0 string, param1 string, param2 float64, param3 float64) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	ret := C.gtk_paper_size_new_from_ppd(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	C.gtk_paper_size_free(cValueInstance)
}

func Fn_gtk_paper_size_get_default_bottom_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_bottom_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_default_left_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_left_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_default_right_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_right_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_default_top_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_top_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_paper_size_get_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_paper_size_get_ppd_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_get_ppd_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_paper_size_get_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_is_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	ret := C.gtk_paper_size_is_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_paper_size_set_size(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 int) {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.GtkUnit)(param2)

	C.gtk_paper_size_set_size(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_paper_size_get_default() string {
	ret := C.gtk_paper_size_get_default()

	return C.GoString(ret)
}

func Fn_gtk_recent_info_exists(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_exists(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_get_added(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_added(cValueInstance)

	return (int64)(ret)
}

func Fn_gtk_recent_info_get_age(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_age(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_recent_info_get_application_info(paramInstance unsafe.Pointer, param0 string, param1 *string, param2 *uint, param3 *int64) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	cValue3 := (*C.time_t)(unsafe.Pointer(param3))

	ret := C.gtk_recent_info_get_application_info(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return toGoBool(ret)
}

func Fn_gtk_recent_info_get_applications(paramInstance unsafe.Pointer, param0 *uint64) []string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	ret := C.gtk_recent_info_get_applications(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]string, retLen, retLen)
	retSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(ret))[:retLen:retLen]
	for retGoi, retGoString := range retSlice {
		retGo[retGoi] = C.GoString(retGoString)
	}
	return retGo
}

func Fn_gtk_recent_info_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_groups(paramInstance unsafe.Pointer, param0 *uint64) []string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	ret := C.gtk_recent_info_get_groups(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]string, retLen, retLen)
	retSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(ret))[:retLen:retLen]
	for retGoi, retGoString := range retSlice {
		retGo[retGoi] = C.GoString(retGoString)
	}
	return retGo
}

func Fn_gtk_recent_info_get_icon(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_recent_info_get_icon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_info_get_mime_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_mime_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_modified(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_modified(cValueInstance)

	return (int64)(ret)
}

func Fn_gtk_recent_info_get_private_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_private_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_get_short_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_short_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_uri_display(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_uri_display(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_visited(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_visited(cValueInstance)

	return (int64)(ret)
}

func Fn_gtk_recent_info_has_application(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_info_has_application(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_has_group(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_info_has_group(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_is_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_is_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_last_application(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_last_application(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_match(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentInfo)(unsafe.Pointer(param0))

	ret := C.gtk_recent_info_match(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	C.gtk_recent_info_unref(cValueInstance)
}

// UNSUPPORTED : gtk_selection_data_get_data : no array length

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

func Fn_gtk_selection_data_targets_include_rich_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	ret := C.gtk_selection_data_targets_include_rich_text(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_targets_include_uri(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_targets_include_uri(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_target_list_add_rich_text_targets(paramInstance unsafe.Pointer, param0 uint, param1 bool, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := toCBool(param1)

	cValue2 := (*C.GtkTextBuffer)(unsafe.Pointer(param2))

	C.gtk_target_list_add_rich_text_targets(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

func Fn_gtk_print_run_page_setup_dialog(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkPageSetup)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkPrintSettings)(unsafe.Pointer(param2))

	ret := C.gtk_print_run_page_setup_dialog(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

func Fn_gtk_target_table_free(param0 []TargetEntry, param1 int) {
	cValue0 := (*C.GtkTargetEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	C.gtk_target_table_free(cValue0, cValue1)
}

func Fn_gtk_target_table_new_from_list(param0 unsafe.Pointer, param1 *int) []TargetEntry {
	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gtk_target_table_new_from_list(cValue0, cValue1)

	retLen := int(*cValue1)
	retGo := make([]TargetEntry, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](TargetEntry))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gtk_targets_include_image(param0 []gdk.Atom, param1 int, param2 bool) bool {
	cValue0 := (*C.GdkAtom)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	cValue2 := toCBool(param2)

	ret := C.gtk_targets_include_image(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_targets_include_rich_text(param0 []gdk.Atom, param1 int, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GdkAtom)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GtkTextBuffer)(unsafe.Pointer(param2))

	ret := C.gtk_targets_include_rich_text(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_targets_include_text(param0 []gdk.Atom, param1 int) bool {
	cValue0 := (*C.GdkAtom)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_targets_include_text(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_targets_include_uri(param0 []gdk.Atom, param1 int) bool {
	cValue0 := (*C.GdkAtom)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_targets_include_uri(cValue0, cValue1)

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

func Fn_gtk_assistant_new() unsafe.Pointer {
	ret := C.gtk_assistant_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_add_action_widget(cValueInstance, cValue0)
}

func Fn_gtk_assistant_append_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_append_page(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_assistant_get_current_page(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	ret := C.gtk_assistant_get_current_page(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_assistant_get_n_pages(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	ret := C.gtk_assistant_get_n_pages(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_assistant_get_nth_page(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_assistant_get_nth_page(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_get_page_complete(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_complete(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_assistant_get_page_header_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_header_image(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_get_page_side_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_side_image(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_get_page_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_title(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_assistant_get_page_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_assistant_insert_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_assistant_insert_page(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_assistant_prepend_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_prepend_page(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_assistant_remove_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_remove_action_widget(cValueInstance, cValue0)
}

func Fn_gtk_assistant_set_current_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_assistant_set_current_page(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

func Fn_gtk_assistant_set_page_complete(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_assistant_set_page_complete(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_header_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_assistant_set_page_header_image(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_side_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_assistant_set_page_side_image(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_assistant_set_page_title(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkAssistantPageType)(param1)

	C.gtk_assistant_set_page_type(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_update_buttons_state(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_update_buttons_state(cValueInstance)
}

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_button_get_image_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_image_position(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_button_set_image_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPositionType)(param0)

	C.gtk_button_set_image_position(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

// UNSUPPORTED : gtk_cell_area_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : parameter 'callback' is callback

func Fn_gtk_cell_renderer_accel_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_accel_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_spin_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_spin_new()

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

func Fn_gtk_clipboard_wait_for_rich_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *uint64) []uint8 {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	ret := C.gtk_clipboard_wait_for_rich_text(cValueInstance, cValue0, cValue1, cValue2)

	retLen := int(*cValue2)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

func Fn_gtk_clipboard_wait_is_rich_text_available(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	ret := C.gtk_clipboard_wait_is_rich_text_available(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

func Fn_gtk_combo_box_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_title(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

func Fn_gtk_combo_box_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_combo_box_set_title(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_entry_get_inner_border(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_inner_border(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_set_inner_border(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBorder)(unsafe.Pointer(param0))

	C.gtk_entry_set_inner_border(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

func Fn_gtk_file_chooser_button_get_focus_on_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_button_get_focus_on_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_button_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_button_set_focus_on_click(cValueInstance, cValue0)
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

func Fn_gtk_label_get_line_wrap_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_line_wrap_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_set_line_wrap_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoWrapMode)(param0)

	C.gtk_label_set_line_wrap_mode(cValueInstance, cValue0)
}

func Fn_gtk_link_button_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_link_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_link_button_new_with_label(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_link_button_new_with_label(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_link_button_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_link_button_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_link_button_set_uri(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_link_button_set_uri(cValueInstance, cValue0)
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

func Fn_gtk_message_dialog_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_message_dialog_set_image(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_tab_detachable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_tab_detachable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_tab_reorderable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_tab_reorderable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_notebook_set_tab_detachable(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_notebook_set_tab_detachable(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_tab_reorderable(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_notebook_set_tab_reorderable(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_new() unsafe.Pointer {
	ret := C.gtk_page_setup_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_page_setup_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_get_bottom_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_bottom_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_left_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_left_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_page_setup_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_page_setup_get_page_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_page_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_page_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_page_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_paper_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_paper_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_paper_size(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_page_setup_get_paper_size(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_get_paper_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_paper_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_right_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_right_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_top_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_top_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_set_bottom_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_bottom_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_set_left_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_left_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_set_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPageOrientation)(param0)

	C.gtk_page_setup_set_orientation(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_set_paper_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	C.gtk_page_setup_set_paper_size(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_set_paper_size_and_default_margins(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	C.gtk_page_setup_set_paper_size_and_default_margins(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_set_right_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_right_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_set_top_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_top_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_context_create_pango_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_create_pango_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_create_pango_layout(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_create_pango_layout(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_cairo_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_cairo_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_dpi_x(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_dpi_x(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_get_dpi_y(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_dpi_y(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_get_height(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_height(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_get_page_setup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_page_setup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_pango_fontmap(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_pango_fontmap(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_width(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_width(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_set_cairo_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 float64) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (C.double)(param1)

	cValue2 := (C.double)(param2)

	C.gtk_print_context_set_cairo_context(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_print_operation_new() unsafe.Pointer {
	ret := C.gtk_print_operation_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_operation_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_cancel(cValueInstance)
}

func Fn_gtk_print_operation_get_default_page_setup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_default_page_setup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_operation_get_error(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	C.gtk_print_operation_get_error(cValueInstance, cError)
}

func Fn_gtk_print_operation_get_print_settings(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_print_settings(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_operation_get_status(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_status(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_operation_get_status_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_status_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_operation_is_finished(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_is_finished(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_run(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintOperationAction)(param0)

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.gtk_print_operation_run(cValueInstance, cValue0, cValue1, cError)

	return (int)(ret)
}

func Fn_gtk_print_operation_set_allow_async(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_allow_async(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_current_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_operation_set_current_page(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_custom_tab_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_operation_set_custom_tab_label(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_default_page_setup(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPageSetup)(unsafe.Pointer(param0))

	C.gtk_print_operation_set_default_page_setup(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_export_filename(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_operation_set_export_filename(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_job_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_operation_set_job_name(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_n_pages(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_operation_set_n_pages(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_print_settings(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPrintSettings)(unsafe.Pointer(param0))

	C.gtk_print_operation_set_print_settings(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_show_progress(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_show_progress(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_track_print_status(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_track_print_status(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_unit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_print_operation_set_unit(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_use_full_page(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_use_full_page(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_new() unsafe.Pointer {
	ret := C.gtk_print_settings_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_settings_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_print_settings_get(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_bool(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get_bool(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_get_collate(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_collate(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_get_default_source(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_default_source(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_dither(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_dither(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_double(paramInstance unsafe.Pointer, param0 string) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get_double(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_double_with_default(paramInstance unsafe.Pointer, param0 string, param1 float64) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	ret := C.gtk_print_settings_get_double_with_default(cValueInstance, cValue0, cValue1)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_duplex(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_duplex(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_finishings(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_finishings(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_int(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get_int(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_int_with_default(paramInstance unsafe.Pointer, param0 string, param1 int) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_print_settings_get_int_with_default(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_length(paramInstance unsafe.Pointer, param0 string, param1 int) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkUnit)(param1)

	ret := C.gtk_print_settings_get_length(cValueInstance, cValue0, cValue1)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_media_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_media_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_n_copies(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_n_copies(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_number_up(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_number_up(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_output_bin(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_output_bin(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_page_ranges(paramInstance unsafe.Pointer, param0 *int) []PageRange {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	ret := C.gtk_print_settings_get_page_ranges(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]PageRange, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](PageRange))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gtk_print_settings_get_page_set(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_page_set(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_paper_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_print_settings_get_paper_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_paper_size(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_paper_size(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_settings_get_paper_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_print_settings_get_paper_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_print_pages(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_print_pages(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_printer(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_printer(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_quality(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_quality(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_resolution(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_resolution(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_reverse(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_reverse(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_get_scale(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_scale(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_use_color(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_use_color(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_has_key(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_has_key(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_set(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_print_settings_set(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_bool(paramInstance unsafe.Pointer, param0 string, param1 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	C.gtk_print_settings_set_bool(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_collate(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_settings_set_collate(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_default_source(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_default_source(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_dither(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_dither(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_double(paramInstance unsafe.Pointer, param0 string, param1 float64) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	C.gtk_print_settings_set_double(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_duplex(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintDuplex)(param0)

	C.gtk_print_settings_set_duplex(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_finishings(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_finishings(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_int(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_print_settings_set_int(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_length(paramInstance unsafe.Pointer, param0 string, param1 float64, param2 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.GtkUnit)(param2)

	C.gtk_print_settings_set_length(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_print_settings_set_media_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_media_type(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_n_copies(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_settings_set_n_copies(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_number_up(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_settings_set_number_up(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPageOrientation)(param0)

	C.gtk_print_settings_set_orientation(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_output_bin(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_output_bin(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_page_ranges(paramInstance unsafe.Pointer, param0 []PageRange, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPageRange)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	C.gtk_print_settings_set_page_ranges(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_page_set(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPageSet)(param0)

	C.gtk_print_settings_set_page_set(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_paper_height(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_print_settings_set_paper_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_paper_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	C.gtk_print_settings_set_paper_size(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_paper_width(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_print_settings_set_paper_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_print_pages(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintPages)(param0)

	C.gtk_print_settings_set_print_pages(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_printer(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_printer(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_quality(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintQuality)(param0)

	C.gtk_print_settings_set_quality(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_resolution(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_settings_set_resolution(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_reverse(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_settings_set_reverse(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_scale(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_print_settings_set_scale(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_use_color(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_settings_set_use_color(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_unset(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_unset(cValueInstance, cValue0)
}

func Fn_gtk_radio_action_set_current_value(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_radio_action_set_current_value(cValueInstance, cValue0)
}

func Fn_gtk_range_get_lower_stepper_sensitivity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_lower_stepper_sensitivity(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_range_get_upper_stepper_sensitivity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_upper_stepper_sensitivity(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_range_set_lower_stepper_sensitivity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSensitivityType)(param0)

	C.gtk_range_set_lower_stepper_sensitivity(cValueInstance, cValue0)
}

func Fn_gtk_range_set_upper_stepper_sensitivity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSensitivityType)(param0)

	C.gtk_range_set_upper_stepper_sensitivity(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_dialog_new(param0 string, param1 unsafe.Pointer, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.c_gtk_recent_chooser_dialog_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_dialog_new_for_manager(param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkRecentManager)(unsafe.Pointer(param2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.c_gtk_recent_chooser_dialog_new_for_manager(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_menu_new() unsafe.Pointer {
	ret := C.gtk_recent_chooser_menu_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_menu_new_for_manager(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRecentManager)(unsafe.Pointer(param0))

	ret := C.gtk_recent_chooser_menu_new_for_manager(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_menu_get_show_numbers(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooserMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_menu_get_show_numbers(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_menu_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooserMenu)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_menu_set_show_numbers(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_widget_new() unsafe.Pointer {
	ret := C.gtk_recent_chooser_widget_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_widget_new_for_manager(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRecentManager)(unsafe.Pointer(param0))

	ret := C.gtk_recent_chooser_widget_new_for_manager(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_filter_new() unsafe.Pointer {
	ret := C.gtk_recent_filter_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_filter_add_age(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_recent_filter_add_age(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_application(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_application(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

func Fn_gtk_recent_filter_add_group(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_group(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_mime_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_mime_type(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_pattern(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_pattern(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_pixbuf_formats(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	C.gtk_recent_filter_add_pixbuf_formats(cValueInstance)
}

func Fn_gtk_recent_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilterInfo)(unsafe.Pointer(param0))

	ret := C.gtk_recent_filter_filter(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_filter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_filter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_filter_get_needed(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_filter_get_needed(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_recent_filter_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_set_name(cValueInstance, cValue0)
}

func Fn_gtk_recent_manager_new() unsafe.Pointer {
	ret := C.gtk_recent_manager_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_manager_add_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkRecentData)(unsafe.Pointer(param1))

	ret := C.gtk_recent_manager_add_full(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_add_item(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_manager_add_item(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_get_items(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_manager_get_items(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_manager_has_item(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_manager_has_item(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_lookup_item(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_lookup_item(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_manager_move_item(paramInstance unsafe.Pointer, param0 string, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_move_item(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_purge_items(paramInstance unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_purge_items(cValueInstance, cError)

	return (int)(ret)
}

func Fn_gtk_recent_manager_remove_item(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_remove_item(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_get_default() unsafe.Pointer {
	ret := C.gtk_recent_manager_get_default()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scrolled_window_unset_placement(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_unset_placement(cValueInstance)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_size_group_get_widgets(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_size_group_get_widgets(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new() unsafe.Pointer {
	ret := C.gtk_status_icon_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_file(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_status_icon_new_from_file(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_icon_name(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_status_icon_new_from_icon_name(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_pixbuf(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_status_icon_new_from_pixbuf(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_status_icon_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_geometry(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, param2 *int) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkOrientation)(unsafe.Pointer(param2))

	ret := C.gtk_status_icon_get_geometry(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_status_icon_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_status_icon_get_stock(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_stock(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_storage_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_storage_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_status_icon_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_status_icon_is_embedded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_is_embedded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_status_icon_set_from_file(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_from_file(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_from_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_from_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_from_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_status_icon_set_from_pixbuf(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_from_stock(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_from_stock(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_status_icon_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_position_menu(param0 unsafe.Pointer, param1 *int, param2 *int, param3 *bool, param4 unsafe.Pointer) {
	cValue0 := (*C.GtkMenu)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gboolean)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	C.gtk_status_icon_position_menu(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_style_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	ret := C.gtk_style_lookup_color(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

func Fn_gtk_text_buffer_deserialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 []uint8, param4 uint64, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(param1)

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	cValue3 := (*C.guint8)(unsafe.Pointer(&param3[0]))

	cValue4 := (C.gsize)(param4)

	cError := (**C.GError)(error)

	ret := C.gtk_text_buffer_deserialize(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_deserialize_get_can_create_tags(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(param0)

	ret := C.gtk_text_buffer_deserialize_get_can_create_tags(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_deserialize_set_can_create_tags(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(param0)

	cValue1 := toCBool(param1)

	C.gtk_text_buffer_deserialize_set_can_create_tags(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_get_copy_target_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_copy_target_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_deserialize_formats(paramInstance unsafe.Pointer, param0 *int) []gdk.Atom {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	ret := C.gtk_text_buffer_get_deserialize_formats(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]gdk.Atom, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](gdk.Atom))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gtk_text_buffer_get_has_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_has_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_get_paste_target_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_paste_target_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_serialize_formats(paramInstance unsafe.Pointer, param0 *int) []gdk.Atom {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	ret := C.gtk_text_buffer_get_serialize_formats(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]gdk.Atom, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](gdk.Atom))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

func Fn_gtk_text_buffer_register_deserialize_tagset(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_text_buffer_register_deserialize_tagset(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

func Fn_gtk_text_buffer_register_serialize_tagset(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_text_buffer_register_serialize_tagset(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_serialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *uint64) []uint8 {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(param1)

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkTextIter)(unsafe.Pointer(param3))

	cValue4 := (*C.gsize)(unsafe.Pointer(param4))

	ret := C.gtk_text_buffer_serialize(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	retLen := int(*cValue4)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gtk_text_buffer_unregister_deserialize_format(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(param0)

	C.gtk_text_buffer_unregister_deserialize_format(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_unregister_serialize_format(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(param0)

	C.gtk_text_buffer_unregister_serialize_format(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

func Fn_gtk_tree_store_insert_with_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 []int, param4 []gobject.Value, param5 int) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(&param3[0]))

	cValue4 := (*C.GValue)(unsafe.Pointer(&param4[0]))

	cValue5 := (C.gint)(param5)

	C.gtk_tree_store_insert_with_valuesv(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

func Fn_gtk_tree_view_get_enable_tree_lines(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_enable_tree_lines(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_grid_lines(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_grid_lines(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_get_headers_clickable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_headers_clickable(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

func Fn_gtk_tree_view_get_rubber_banding(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_rubber_banding(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_search_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_search_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_enable_tree_lines(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_enable_tree_lines(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_grid_lines(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTreeViewGridLines)(param0)

	C.gtk_tree_view_set_grid_lines(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_rubber_banding(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_rubber_banding(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_search_entry(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkEntry)(unsafe.Pointer(param0))

	C.gtk_tree_view_set_search_entry(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_drag_dest_get_track_motion(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_drag_dest_get_track_motion(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_drag_dest_set_track_motion(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_drag_dest_set_track_motion(cValueInstance, cValue0)
}

func Fn_gtk_widget_is_composited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_composited(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_window_get_deletable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_deletable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_set_deletable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_deletable(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

func Fn_gtk_print_operation_preview_end_preview(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperationPreview)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_preview_end_preview(cValueInstance)
}

func Fn_gtk_print_operation_preview_is_selected(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkPrintOperationPreview)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_print_operation_preview_is_selected(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_preview_render_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperationPreview)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_operation_preview_render_page(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_add_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilter)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_add_filter(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_get_current_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_current_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_get_current_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_current_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_chooser_get_filter(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_filter(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_get_items(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_items(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_get_limit(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_limit(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_recent_chooser_get_local_only(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_local_only(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_select_multiple(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_select_multiple(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_icons(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_icons(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_not_found(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_not_found(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_private(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_private(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_tips(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_tips(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_sort_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_sort_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_recent_chooser_get_uris(paramInstance unsafe.Pointer, param0 *uint64) []string {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	ret := C.gtk_recent_chooser_get_uris(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]string, retLen, retLen)
	retSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(ret))[:retLen:retLen]
	for retGoi, retGoString := range retSlice {
		retGo[retGoi] = C.GoString(retGoString)
	}
	return retGo
}

func Fn_gtk_recent_chooser_list_filters(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_list_filters(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_remove_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilter)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_remove_filter(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	C.gtk_recent_chooser_select_all(cValueInstance)
}

func Fn_gtk_recent_chooser_select_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_chooser_select_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_set_current_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_chooser_set_current_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_set_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilter)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_set_filter(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_limit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_recent_chooser_set_limit(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_local_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_local_only(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_select_multiple(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_select_multiple(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_icons(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_icons(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_not_found(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_not_found(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_private(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_private(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_tips(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_tips(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_recent_chooser_set_sort_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkRecentSortType)(param0)

	C.gtk_recent_chooser_set_sort_type(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	C.gtk_recent_chooser_unselect_all(cValueInstance)
}

func Fn_gtk_recent_chooser_unselect_uri(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_chooser_unselect_uri(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

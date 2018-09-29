// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_about_dialog_add_credit_section : unsupported parameter people : no param type

// Unsupported : gtk_about_dialog_get_artists : no return type

// Unsupported : gtk_about_dialog_get_authors : no return type

// Unsupported : gtk_about_dialog_get_documenters : no return type

// Unsupported : gtk_about_dialog_set_artists : unsupported parameter artists : no param type

// Unsupported : gtk_about_dialog_set_authors : unsupported parameter authors : no param type

// Unsupported : gtk_about_dialog_set_documenters : unsupported parameter documenters : no param type

// Unsupported : gtk_accel_group_find : unsupported parameter find_func : no type generator for AccelGroupFindFunc, GtkAccelGroupFindFunc

// Unsupported : gtk_accel_group_query : unsupported parameter n_entries : no type generator for guint, guint*

// Unsupported : gtk_accel_label_get_accel : unsupported parameter accelerator_key : no type generator for guint, guint*

// Unsupported : gtk_action_create_icon : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// Unsupported : gtk_action_get_gicon : no return generator

// Unsupported : gtk_action_set_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GtkTranslateFunc

// Unsupported : gtk_alignment_get_padding : unsupported parameter padding_top : no type generator for guint, guint*

// Unsupported : gtk_app_chooser_button_append_custom_item : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_application_add_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_get_actions_for_accel : no return type

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_remove_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels : no param type

// Unsupported : gtk_assistant_set_forward_page_func : unsupported parameter page_func : no type generator for AssistantPageFunc, GtkAssistantPageFunc

// GetBaselinePosition is a wrapper around the C function gtk_box_get_baseline_position.
func (recv *Box) GetBaselinePosition() BaselinePosition {
	retC := C.gtk_box_get_baseline_position((*C.GtkBox)(recv.native))
	retGo := (BaselinePosition)(retC)

	return retGo
}

// Unsupported : gtk_box_query_child_packing : unsupported parameter padding : no type generator for guint, guint*

// SetBaselinePosition is a wrapper around the C function gtk_box_set_baseline_position.
func (recv *Box) SetBaselinePosition(position BaselinePosition) {
	c_position := (C.GtkBaselinePosition)(position)

	C.gtk_box_set_baseline_position((*C.GtkBox)(recv.native), c_position)

	return
}

// BuilderNewFromFile is a wrapper around the C function gtk_builder_new_from_file.
func BuilderNewFromFile(filename string) *Builder {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_builder_new_from_file(c_filename)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BuilderNewFromResource is a wrapper around the C function gtk_builder_new_from_resource.
func BuilderNewFromResource(resourcePath string) *Builder {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_builder_new_from_resource(c_resource_path)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BuilderNewFromString is a wrapper around the C function gtk_builder_new_from_string.
func BuilderNewFromString(string string, length int64) *Builder {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gssize)(length)

	retC := C.gtk_builder_new_from_string(c_string, c_length)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_objects_from_file : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_resource : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_string : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// Unsupported : gtk_builder_extend_with_template : unsupported parameter template_type : no type generator for GType, GType

// GetApplication is a wrapper around the C function gtk_builder_get_application.
func (recv *Builder) GetApplication() *Application {
	retC := C.gtk_builder_get_application((*C.GtkBuilder)(recv.native))
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_builder_get_type_from_name : no return generator

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// SetApplication is a wrapper around the C function gtk_builder_set_application.
func (recv *Builder) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(application.ToC())

	C.gtk_builder_set_application((*C.GtkBuilder)(recv.native), c_application)

	return
}

// Unsupported : gtk_builder_value_from_string_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_button_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_calendar_get_date : unsupported parameter year : no type generator for guint, guint*

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc, GtkCalendarDetailFunc

// Unsupported : gtk_cell_area_activate : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_activate_cell : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_area_add_with_properties : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_apply_attributes : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_cell_area_cell_get : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_cell_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_cell_area_cell_set : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_cell_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_cell_area_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_area_foreach : unsupported parameter callback : no type generator for CellCallback, GtkCellCallback

// Unsupported : gtk_cell_area_foreach_alloc : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_allocation : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_at_position : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_edit_widget : no return generator

// Unsupported : gtk_cell_area_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_inner_cell_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_request_renderer : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_allocation : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_activate : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_renderer_get_aligned_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_cell_renderer_get_fixed_size : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_width : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_start_editing : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_view_get_model : no return generator

// Unsupported : gtk_cell_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_clipboard_get_selection : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_request_contents : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_request_image : unsupported parameter callback : no type generator for ClipboardImageReceivedFunc, GtkClipboardImageReceivedFunc

// Unsupported : gtk_clipboard_request_rich_text : unsupported parameter callback : no type generator for ClipboardRichTextReceivedFunc, GtkClipboardRichTextReceivedFunc

// Unsupported : gtk_clipboard_request_targets : unsupported parameter callback : no type generator for ClipboardTargetsReceivedFunc, GtkClipboardTargetsReceivedFunc

// Unsupported : gtk_clipboard_request_text : unsupported parameter callback : no type generator for ClipboardTextReceivedFunc, GtkClipboardTextReceivedFunc

// Unsupported : gtk_clipboard_request_uris : unsupported parameter callback : no type generator for ClipboardURIReceivedFunc, GtkClipboardURIReceivedFunc

// Unsupported : gtk_clipboard_set_can_store : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_set_with_data : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_set_with_owner : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_wait_for_contents : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_wait_for_rich_text : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_wait_for_targets : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_wait_for_uris : no return type

// Unsupported : gtk_clipboard_wait_is_target_available : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_get_model : no return generator

// Unsupported : gtk_combo_box_get_row_separator_func : no return generator

// Unsupported : gtk_combo_box_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc, GtkTreeViewRowSeparatorFunc

// Unsupported : gtk_container_add_with_properties : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_get : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_set : unsupported parameter ... : varargs

// Unsupported : gtk_container_child_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_type : no return generator

// Unsupported : gtk_container_forall : unsupported parameter callback : no type generator for Callback, GtkCallback

// Unsupported : gtk_container_foreach : unsupported parameter callback : no type generator for Callback, GtkCallback

// Unsupported : gtk_container_get_focus_chain : unsupported parameter focusable_widgets : record with indirection level of 2

// Unsupported : gtk_css_provider_load_from_data : unsupported parameter data : no param type

// Unsupported : gtk_css_provider_load_from_file : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_add_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_set_alternative_button_order : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_set_alternative_button_order_from_array : unsupported parameter new_order : no param type

// Unsupported : gtk_entry_get_icon_area : unsupported parameter icon_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_get_icon_gicon : no return generator

// Unsupported : gtk_entry_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// GetTabs is a wrapper around the C function gtk_entry_get_tabs.
func (recv *Entry) GetTabs() *pango.TabArray {
	retC := C.gtk_entry_get_tabs((*C.GtkEntry)(recv.native))
	retGo := pango.TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_set_icon_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SetTabs is a wrapper around the C function gtk_entry_set_tabs.
func (recv *Entry) SetTabs(tabs *pango.TabArray) {
	c_tabs := (*C.PangoTabArray)(tabs.ToC())

	C.gtk_entry_set_tabs((*C.GtkEntry)(recv.native), c_tabs)

	return
}

// Unsupported : gtk_entry_completion_get_model : no return generator

// Unsupported : gtk_entry_completion_set_match_func : unsupported parameter func : no type generator for EntryCompletionMatchFunc, GtkEntryCompletionMatchFunc

// Unsupported : gtk_entry_completion_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_event_controller_handle_event : unsupported parameter event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_file_filter_add_custom : unsupported parameter func : no type generator for FileFilterFunc, GtkFileFilterFunc

// Unsupported : gtk_file_filter_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_flow_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// Unsupported : gtk_flow_box_selected_foreach : unsupported parameter func : no type generator for FlowBoxForeachFunc, GtkFlowBoxForeachFunc

// Unsupported : gtk_flow_box_set_filter_func : unsupported parameter filter_func : no type generator for FlowBoxFilterFunc, GtkFlowBoxFilterFunc

// Unsupported : gtk_flow_box_set_sort_func : unsupported parameter sort_func : no type generator for FlowBoxSortFunc, GtkFlowBoxSortFunc

// Unsupported : gtk_frame_get_label_align : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_gl_area_get_required_version : unsupported parameter major : no type generator for gint, gint*

// Unsupported : gtk_gesture_get_bounding_box : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_get_bounding_box_center : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_get_last_event : no return generator

// Unsupported : gtk_gesture_get_point : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_drag_get_offset : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_drag_get_start_point : unsupported parameter x : no type generator for gdouble, gdouble*

// Unsupported : gtk_gesture_multi_press_get_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_multi_press_set_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_gesture_swipe_get_velocity : unsupported parameter velocity_x : no type generator for gdouble, gdouble*

// GetBaselineRow is a wrapper around the C function gtk_grid_get_baseline_row.
func (recv *Grid) GetBaselineRow() int32 {
	retC := C.gtk_grid_get_baseline_row((*C.GtkGrid)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRowBaselinePosition is a wrapper around the C function gtk_grid_get_row_baseline_position.
func (recv *Grid) GetRowBaselinePosition(row int32) BaselinePosition {
	c_row := (C.gint)(row)

	retC := C.gtk_grid_get_row_baseline_position((*C.GtkGrid)(recv.native), c_row)
	retGo := (BaselinePosition)(retC)

	return retGo
}

// RemoveColumn is a wrapper around the C function gtk_grid_remove_column.
func (recv *Grid) RemoveColumn(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_column((*C.GtkGrid)(recv.native), c_position)

	return
}

// RemoveRow is a wrapper around the C function gtk_grid_remove_row.
func (recv *Grid) RemoveRow(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_row((*C.GtkGrid)(recv.native), c_position)

	return
}

// SetBaselineRow is a wrapper around the C function gtk_grid_set_baseline_row.
func (recv *Grid) SetBaselineRow(row int32) {
	c_row := (C.gint)(row)

	C.gtk_grid_set_baseline_row((*C.GtkGrid)(recv.native), c_row)

	return
}

// SetRowBaselinePosition is a wrapper around the C function gtk_grid_set_row_baseline_position.
func (recv *Grid) SetRowBaselinePosition(row int32, pos BaselinePosition) {
	c_row := (C.gint)(row)

	c_pos := (C.GtkBaselinePosition)(pos)

	C.gtk_grid_set_row_baseline_position((*C.GtkGrid)(recv.native), c_row, c_pos)

	return
}

// Unsupported : gtk_hsv_get_color : unsupported parameter h : no type generator for gdouble, gdouble*

// Unsupported : gtk_hsv_get_metrics : unsupported parameter size : no type generator for gint, gint*

// HeaderBarNew is a wrapper around the C function gtk_header_bar_new.
func HeaderBarNew() *Widget {
	retC := C.gtk_header_bar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCustomTitle is a wrapper around the C function gtk_header_bar_get_custom_title.
func (recv *HeaderBar) GetCustomTitle() *Widget {
	retC := C.gtk_header_bar_get_custom_title((*C.GtkHeaderBar)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShowCloseButton is a wrapper around the C function gtk_header_bar_get_show_close_button.
func (recv *HeaderBar) GetShowCloseButton() bool {
	retC := C.gtk_header_bar_get_show_close_button((*C.GtkHeaderBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSubtitle is a wrapper around the C function gtk_header_bar_get_subtitle.
func (recv *HeaderBar) GetSubtitle() string {
	retC := C.gtk_header_bar_get_subtitle((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTitle is a wrapper around the C function gtk_header_bar_get_title.
func (recv *HeaderBar) GetTitle() string {
	retC := C.gtk_header_bar_get_title((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// PackEnd is a wrapper around the C function gtk_header_bar_pack_end.
func (recv *HeaderBar) PackEnd(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_header_bar_pack_end((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// PackStart is a wrapper around the C function gtk_header_bar_pack_start.
func (recv *HeaderBar) PackStart(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_header_bar_pack_start((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// SetCustomTitle is a wrapper around the C function gtk_header_bar_set_custom_title.
func (recv *HeaderBar) SetCustomTitle(titleWidget *Widget) {
	c_title_widget := (*C.GtkWidget)(titleWidget.ToC())

	C.gtk_header_bar_set_custom_title((*C.GtkHeaderBar)(recv.native), c_title_widget)

	return
}

// SetShowCloseButton is a wrapper around the C function gtk_header_bar_set_show_close_button.
func (recv *HeaderBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_header_bar_set_show_close_button((*C.GtkHeaderBar)(recv.native), c_setting)

	return
}

// SetSubtitle is a wrapper around the C function gtk_header_bar_set_subtitle.
func (recv *HeaderBar) SetSubtitle(subtitle string) {
	c_subtitle := C.CString(subtitle)
	defer C.free(unsafe.Pointer(c_subtitle))

	C.gtk_header_bar_set_subtitle((*C.GtkHeaderBar)(recv.native), c_subtitle)

	return
}

// SetTitle is a wrapper around the C function gtk_header_bar_set_title.
func (recv *HeaderBar) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_header_bar_set_title((*C.GtkHeaderBar)(recv.native), c_title)

	return
}

// Unsupported : gtk_im_context_get_preedit_string : unsupported parameter attrs : record with indirection level of 2

// Unsupported : gtk_im_context_get_surrounding : unsupported parameter cursor_index : no type generator for gint, gint*

// Unsupported : gtk_im_context_set_cursor_location : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_im_context_simple_add_table : unsupported parameter data : no param type

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : no param type

// GetBaseScale is a wrapper around the C function gtk_icon_info_get_base_scale.
func (recv *IconInfo) GetBaseScale() int32 {
	retC := C.gtk_icon_info_get_base_scale((*C.GtkIconInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_icon_info_get_embedded_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gtk_icon_info_load_icon_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_icon_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// LoadSurface is a wrapper around the C function gtk_icon_info_load_surface.
func (recv *IconInfo) LoadSurface(forWindow *gdk.Window) (*cairo.Surface, error) {
	c_for_window := (*C.GdkWindow)(forWindow.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_surface((*C.GtkIconInfo)(recv.native), c_for_window, &cThrowableError)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_for_context_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_theme_choose_icon : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_get_icon_sizes : no return type

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : no param type

// LoadIconForScale is a wrapper around the C function gtk_icon_theme_load_icon_for_scale.
func (recv *IconTheme) LoadIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LoadSurface is a wrapper around the C function gtk_icon_theme_load_surface.
func (recv *IconTheme) LoadSurface(iconName string, size int32, scale int32, forWindow *gdk.Window, flags IconLookupFlags) (*cairo.Surface, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_for_window := (*C.GdkWindow)(forWindow.ToC())

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_surface((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_for_window, c_flags, &cThrowableError)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_icon_theme_lookup_by_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_lookup_by_gicon_for_scale : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// LookupIconForScale is a wrapper around the C function gtk_icon_theme_lookup_icon_for_scale.
func (recv *IconTheme) LookupIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags)
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_theme_set_search_path : unsupported parameter path : no param type

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_icon_view_convert_widget_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_icon_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_get_cell_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_icon_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_dest_item_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_drag_dest_item : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_item_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_model : no return generator

// Unsupported : gtk_icon_view_get_tooltip_context : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_icon_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// Unsupported : gtk_icon_view_selected_foreach : unsupported parameter func : no type generator for IconViewForeachFunc, GtkIconViewForeachFunc

// Unsupported : gtk_icon_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// ImageNewFromSurface is a wrapper around the C function gtk_image_new_from_surface.
func ImageNewFromSurface(surface *cairo.Surface) *Widget {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	retC := C.gtk_image_new_from_surface(c_surface)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon**

// Unsupported : gtk_image_get_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize*

// Unsupported : gtk_image_get_icon_set : unsupported parameter icon_set : record with indirection level of 2

// Unsupported : gtk_image_get_stock : unsupported parameter size : no type generator for gint, GtkIconSize*

// Unsupported : gtk_image_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_set_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// SetFromSurface is a wrapper around the C function gtk_image_set_from_surface.
func (recv *Image) SetFromSurface(surface *cairo.Surface) {
	c_surface := (*C.cairo_surface_t)(surface.ToC())

	C.gtk_image_set_from_surface((*C.GtkImage)(recv.native), c_surface)

	return
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// GetShowCloseButton is a wrapper around the C function gtk_info_bar_get_show_close_button.
func (recv *InfoBar) GetShowCloseButton() bool {
	retC := C.gtk_info_bar_get_show_close_button((*C.GtkInfoBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowCloseButton is a wrapper around the C function gtk_info_bar_set_show_close_button.
func (recv *InfoBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_info_bar_set_show_close_button((*C.GtkInfoBar)(recv.native), c_setting)

	return
}

// Unsupported : gtk_label_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// GetLines is a wrapper around the C function gtk_label_get_lines.
func (recv *Label) GetLines() int32 {
	retC := C.gtk_label_get_lines((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_label_get_selection_bounds : unsupported parameter start : no type generator for gint, gint*

// SetLines is a wrapper around the C function gtk_label_set_lines.
func (recv *Label) SetLines(lines int32) {
	c_lines := (C.gint)(lines)

	C.gtk_label_set_lines((*C.GtkLabel)(recv.native), c_lines)

	return
}

// Unsupported : gtk_layout_get_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gtk_level_bar_get_offset_value : unsupported parameter value : no type generator for gdouble, gdouble*

// ListBoxNew is a wrapper around the C function gtk_list_box_new.
func ListBoxNew() *Widget {
	retC := C.gtk_list_box_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_list_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// DragHighlightRow is a wrapper around the C function gtk_list_box_drag_highlight_row.
func (recv *ListBox) DragHighlightRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(row.ToC())

	C.gtk_list_box_drag_highlight_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// DragUnhighlightRow is a wrapper around the C function gtk_list_box_drag_unhighlight_row.
func (recv *ListBox) DragUnhighlightRow() {
	C.gtk_list_box_drag_unhighlight_row((*C.GtkListBox)(recv.native))

	return
}

// GetActivateOnSingleClick is a wrapper around the C function gtk_list_box_get_activate_on_single_click.
func (recv *ListBox) GetActivateOnSingleClick() bool {
	retC := C.gtk_list_box_get_activate_on_single_click((*C.GtkListBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetAdjustment is a wrapper around the C function gtk_list_box_get_adjustment.
func (recv *ListBox) GetAdjustment() *Adjustment {
	retC := C.gtk_list_box_get_adjustment((*C.GtkListBox)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRowAtIndex is a wrapper around the C function gtk_list_box_get_row_at_index.
func (recv *ListBox) GetRowAtIndex(index int32) *ListBoxRow {
	c_index_ := (C.gint)(index)

	retC := C.gtk_list_box_get_row_at_index((*C.GtkListBox)(recv.native), c_index_)
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRowAtY is a wrapper around the C function gtk_list_box_get_row_at_y.
func (recv *ListBox) GetRowAtY(y int32) *ListBoxRow {
	c_y := (C.gint)(y)

	retC := C.gtk_list_box_get_row_at_y((*C.GtkListBox)(recv.native), c_y)
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectedRow is a wrapper around the C function gtk_list_box_get_selected_row.
func (recv *ListBox) GetSelectedRow() *ListBoxRow {
	retC := C.gtk_list_box_get_selected_row((*C.GtkListBox)(recv.native))
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionMode is a wrapper around the C function gtk_list_box_get_selection_mode.
func (recv *ListBox) GetSelectionMode() SelectionMode {
	retC := C.gtk_list_box_get_selection_mode((*C.GtkListBox)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Insert is a wrapper around the C function gtk_list_box_insert.
func (recv *ListBox) Insert(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_position := (C.gint)(position)

	C.gtk_list_box_insert((*C.GtkListBox)(recv.native), c_child, c_position)

	return
}

// InvalidateFilter is a wrapper around the C function gtk_list_box_invalidate_filter.
func (recv *ListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter((*C.GtkListBox)(recv.native))

	return
}

// InvalidateHeaders is a wrapper around the C function gtk_list_box_invalidate_headers.
func (recv *ListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers((*C.GtkListBox)(recv.native))

	return
}

// InvalidateSort is a wrapper around the C function gtk_list_box_invalidate_sort.
func (recv *ListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort((*C.GtkListBox)(recv.native))

	return
}

// Prepend is a wrapper around the C function gtk_list_box_prepend.
func (recv *ListBox) Prepend(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_list_box_prepend((*C.GtkListBox)(recv.native), c_child)

	return
}

// SelectRow is a wrapper around the C function gtk_list_box_select_row.
func (recv *ListBox) SelectRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(row.ToC())

	C.gtk_list_box_select_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// Unsupported : gtk_list_box_selected_foreach : unsupported parameter func : no type generator for ListBoxForeachFunc, GtkListBoxForeachFunc

// SetActivateOnSingleClick is a wrapper around the C function gtk_list_box_set_activate_on_single_click.
func (recv *ListBox) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_list_box_set_activate_on_single_click((*C.GtkListBox)(recv.native), c_single)

	return
}

// SetAdjustment is a wrapper around the C function gtk_list_box_set_adjustment.
func (recv *ListBox) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_list_box_set_adjustment((*C.GtkListBox)(recv.native), c_adjustment)

	return
}

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc, GtkListBoxFilterFunc

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc, GtkListBoxUpdateHeaderFunc

// SetPlaceholder is a wrapper around the C function gtk_list_box_set_placeholder.
func (recv *ListBox) SetPlaceholder(placeholder *Widget) {
	c_placeholder := (*C.GtkWidget)(placeholder.ToC())

	C.gtk_list_box_set_placeholder((*C.GtkListBox)(recv.native), c_placeholder)

	return
}

// SetSelectionMode is a wrapper around the C function gtk_list_box_set_selection_mode.
func (recv *ListBox) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_list_box_set_selection_mode((*C.GtkListBox)(recv.native), c_mode)

	return
}

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc, GtkListBoxSortFunc

// ListBoxRowNew is a wrapper around the C function gtk_list_box_row_new.
func ListBoxRowNew() *Widget {
	retC := C.gtk_list_box_row_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Changed is a wrapper around the C function gtk_list_box_row_changed.
func (recv *ListBoxRow) Changed() {
	C.gtk_list_box_row_changed((*C.GtkListBoxRow)(recv.native))

	return
}

// GetHeader is a wrapper around the C function gtk_list_box_row_get_header.
func (recv *ListBoxRow) GetHeader() *Widget {
	retC := C.gtk_list_box_row_get_header((*C.GtkListBoxRow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIndex is a wrapper around the C function gtk_list_box_row_get_index.
func (recv *ListBoxRow) GetIndex() int32 {
	retC := C.gtk_list_box_row_get_index((*C.GtkListBoxRow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetHeader is a wrapper around the C function gtk_list_box_row_set_header.
func (recv *ListBoxRow) SetHeader(header *Widget) {
	c_header := (*C.GtkWidget)(header.ToC())

	C.gtk_list_box_row_set_header((*C.GtkListBoxRow)(recv.native), c_header)

	return
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_list_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_insert_with_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_list_store_reorder : unsupported parameter new_order : no param type

// Unsupported : gtk_list_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_list_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_menu_attach_to_widget : unsupported parameter detacher : no type generator for MenuDetachFunc, GtkMenuDetachFunc

// Unsupported : gtk_menu_popup : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// Unsupported : gtk_menu_popup_at_pointer : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_at_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_menu_popup_at_widget : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_for_device : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// Unsupported : gtk_menu_item_toggle_size_request : unsupported parameter requisition : no type generator for gint, gint*

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_text : unsupported parameter ... : varargs

// Unsupported : gtk_misc_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_misc_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_numerable_icon_get_background_gicon : no return generator

// Unsupported : gtk_numerable_icon_set_background_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_pad_controller_set_action_entries : unsupported parameter entries : no param type

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_page_setup_to_gvariant : return type : Blacklisted record : GVariant

// PlacesSidebarNew is a wrapper around the C function gtk_places_sidebar_new.
func PlacesSidebarNew() *Widget {
	retC := C.gtk_places_sidebar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_places_sidebar_add_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_get_location : no return generator

// Unsupported : gtk_places_sidebar_get_nth_bookmark : no return generator

// GetOpenFlags is a wrapper around the C function gtk_places_sidebar_get_open_flags.
func (recv *PlacesSidebar) GetOpenFlags() PlacesOpenFlags {
	retC := C.gtk_places_sidebar_get_open_flags((*C.GtkPlacesSidebar)(recv.native))
	retGo := (PlacesOpenFlags)(retC)

	return retGo
}

// GetShowDesktop is a wrapper around the C function gtk_places_sidebar_get_show_desktop.
func (recv *PlacesSidebar) GetShowDesktop() bool {
	retC := C.gtk_places_sidebar_get_show_desktop((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListShortcuts is a wrapper around the C function gtk_places_sidebar_list_shortcuts.
func (recv *PlacesSidebar) ListShortcuts() *glib.SList {
	retC := C.gtk_places_sidebar_list_shortcuts((*C.GtkPlacesSidebar)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_places_sidebar_remove_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_set_location : unsupported parameter location : no type generator for Gio.File, GFile*

// SetOpenFlags is a wrapper around the C function gtk_places_sidebar_set_open_flags.
func (recv *PlacesSidebar) SetOpenFlags(flags PlacesOpenFlags) {
	c_flags := (C.GtkPlacesOpenFlags)(flags)

	C.gtk_places_sidebar_set_open_flags((*C.GtkPlacesSidebar)(recv.native), c_flags)

	return
}

// SetShowConnectToServer is a wrapper around the C function gtk_places_sidebar_set_show_connect_to_server.
func (recv *PlacesSidebar) SetShowConnectToServer(showConnectToServer bool) {
	c_show_connect_to_server :=
		boolToGboolean(showConnectToServer)

	C.gtk_places_sidebar_set_show_connect_to_server((*C.GtkPlacesSidebar)(recv.native), c_show_connect_to_server)

	return
}

// SetShowDesktop is a wrapper around the C function gtk_places_sidebar_set_show_desktop.
func (recv *PlacesSidebar) SetShowDesktop(showDesktop bool) {
	c_show_desktop :=
		boolToGboolean(showDesktop)

	C.gtk_places_sidebar_set_show_desktop((*C.GtkPlacesSidebar)(recv.native), c_show_desktop)

	return
}

// Blacklisted : GtkPlug

// Unsupported : gtk_popover_get_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_popover_set_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_print_context_get_hard_margins : unsupported parameter top : no type generator for gdouble, gdouble*

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_foreach : unsupported parameter func : no type generator for PrintSettingsFunc, GtkPrintSettingsFunc

// Unsupported : gtk_print_settings_get_page_ranges : unsupported parameter num_ranges : no type generator for gint, gint*

// Unsupported : gtk_print_settings_set_page_ranges : unsupported parameter page_ranges : no param type

// Unsupported : gtk_print_settings_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_range_get_range_rect : unsupported parameter range_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_range_get_slider_range : unsupported parameter slider_start : no type generator for gint, gint*

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_recent_filter_add_custom : unsupported parameter func : no type generator for RecentFilterFunc, GtkRecentFilterFunc

// RevealerNew is a wrapper around the C function gtk_revealer_new.
func RevealerNew() *Widget {
	retC := C.gtk_revealer_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChildRevealed is a wrapper around the C function gtk_revealer_get_child_revealed.
func (recv *Revealer) GetChildRevealed() bool {
	retC := C.gtk_revealer_get_child_revealed((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRevealChild is a wrapper around the C function gtk_revealer_get_reveal_child.
func (recv *Revealer) GetRevealChild() bool {
	retC := C.gtk_revealer_get_reveal_child((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTransitionDuration is a wrapper around the C function gtk_revealer_get_transition_duration.
func (recv *Revealer) GetTransitionDuration() uint32 {
	retC := C.gtk_revealer_get_transition_duration((*C.GtkRevealer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTransitionType is a wrapper around the C function gtk_revealer_get_transition_type.
func (recv *Revealer) GetTransitionType() RevealerTransitionType {
	retC := C.gtk_revealer_get_transition_type((*C.GtkRevealer)(recv.native))
	retGo := (RevealerTransitionType)(retC)

	return retGo
}

// SetRevealChild is a wrapper around the C function gtk_revealer_set_reveal_child.
func (recv *Revealer) SetRevealChild(revealChild bool) {
	c_reveal_child :=
		boolToGboolean(revealChild)

	C.gtk_revealer_set_reveal_child((*C.GtkRevealer)(recv.native), c_reveal_child)

	return
}

// SetTransitionDuration is a wrapper around the C function gtk_revealer_set_transition_duration.
func (recv *Revealer) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_revealer_set_transition_duration((*C.GtkRevealer)(recv.native), c_duration)

	return
}

// SetTransitionType is a wrapper around the C function gtk_revealer_set_transition_type.
func (recv *Revealer) SetTransitionType(transition RevealerTransitionType) {
	c_transition := (C.GtkRevealerTransitionType)(transition)

	C.gtk_revealer_set_transition_type((*C.GtkRevealer)(recv.native), c_transition)

	return
}

// Unsupported : gtk_scale_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons : no param type

// Unsupported : gtk_scrolled_window_get_policy : unsupported parameter hscrollbar_policy : GtkPolicyType* with indirection level of 1

// SearchBarNew is a wrapper around the C function gtk_search_bar_new.
func SearchBarNew() *Widget {
	retC := C.gtk_search_bar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ConnectEntry is a wrapper around the C function gtk_search_bar_connect_entry.
func (recv *SearchBar) ConnectEntry(entry *Entry) {
	c_entry := (*C.GtkEntry)(entry.ToC())

	C.gtk_search_bar_connect_entry((*C.GtkSearchBar)(recv.native), c_entry)

	return
}

// GetSearchMode is a wrapper around the C function gtk_search_bar_get_search_mode.
func (recv *SearchBar) GetSearchMode() bool {
	retC := C.gtk_search_bar_get_search_mode((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowCloseButton is a wrapper around the C function gtk_search_bar_get_show_close_button.
func (recv *SearchBar) GetShowCloseButton() bool {
	retC := C.gtk_search_bar_get_show_close_button((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// SetSearchMode is a wrapper around the C function gtk_search_bar_set_search_mode.
func (recv *SearchBar) SetSearchMode(searchMode bool) {
	c_search_mode :=
		boolToGboolean(searchMode)

	C.gtk_search_bar_set_search_mode((*C.GtkSearchBar)(recv.native), c_search_mode)

	return
}

// SetShowCloseButton is a wrapper around the C function gtk_search_bar_set_show_close_button.
func (recv *SearchBar) SetShowCloseButton(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_search_bar_set_show_close_button((*C.GtkSearchBar)(recv.native), c_visible)

	return
}

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Blacklisted : GtkSocket

// Unsupported : gtk_spin_button_get_increments : unsupported parameter step : no type generator for gdouble, gdouble*

// Unsupported : gtk_spin_button_get_range : unsupported parameter min : no type generator for gdouble, gdouble*

// StackNew is a wrapper around the C function gtk_stack_new.
func StackNew() *Widget {
	retC := C.gtk_stack_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddNamed is a wrapper around the C function gtk_stack_add_named.
func (recv *Stack) AddNamed(child *Widget, name string) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_add_named((*C.GtkStack)(recv.native), c_child, c_name)

	return
}

// AddTitled is a wrapper around the C function gtk_stack_add_titled.
func (recv *Stack) AddTitled(child *Widget, name string, title string) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_stack_add_titled((*C.GtkStack)(recv.native), c_child, c_name, c_title)

	return
}

// GetHomogeneous is a wrapper around the C function gtk_stack_get_homogeneous.
func (recv *Stack) GetHomogeneous() bool {
	retC := C.gtk_stack_get_homogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTransitionDuration is a wrapper around the C function gtk_stack_get_transition_duration.
func (recv *Stack) GetTransitionDuration() uint32 {
	retC := C.gtk_stack_get_transition_duration((*C.GtkStack)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTransitionType is a wrapper around the C function gtk_stack_get_transition_type.
func (recv *Stack) GetTransitionType() StackTransitionType {
	retC := C.gtk_stack_get_transition_type((*C.GtkStack)(recv.native))
	retGo := (StackTransitionType)(retC)

	return retGo
}

// GetVisibleChild is a wrapper around the C function gtk_stack_get_visible_child.
func (recv *Stack) GetVisibleChild() *Widget {
	retC := C.gtk_stack_get_visible_child((*C.GtkStack)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisibleChildName is a wrapper around the C function gtk_stack_get_visible_child_name.
func (recv *Stack) GetVisibleChildName() string {
	retC := C.gtk_stack_get_visible_child_name((*C.GtkStack)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetHomogeneous is a wrapper around the C function gtk_stack_set_homogeneous.
func (recv *Stack) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_stack_set_homogeneous((*C.GtkStack)(recv.native), c_homogeneous)

	return
}

// SetTransitionDuration is a wrapper around the C function gtk_stack_set_transition_duration.
func (recv *Stack) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_stack_set_transition_duration((*C.GtkStack)(recv.native), c_duration)

	return
}

// SetTransitionType is a wrapper around the C function gtk_stack_set_transition_type.
func (recv *Stack) SetTransitionType(transition StackTransitionType) {
	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_transition_type((*C.GtkStack)(recv.native), c_transition)

	return
}

// SetVisibleChild is a wrapper around the C function gtk_stack_set_visible_child.
func (recv *Stack) SetVisibleChild(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_stack_set_visible_child((*C.GtkStack)(recv.native), c_child)

	return
}

// SetVisibleChildFull is a wrapper around the C function gtk_stack_set_visible_child_full.
func (recv *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_visible_child_full((*C.GtkStack)(recv.native), c_name, c_transition)

	return
}

// SetVisibleChildName is a wrapper around the C function gtk_stack_set_visible_child_name.
func (recv *Stack) SetVisibleChildName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_set_visible_child_name((*C.GtkStack)(recv.native), c_name)

	return
}

// Blacklisted : GtkStackAccessible

// StackSwitcherNew is a wrapper around the C function gtk_stack_switcher_new.
func StackSwitcherNew() *Widget {
	retC := C.gtk_stack_switcher_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStack is a wrapper around the C function gtk_stack_switcher_get_stack.
func (recv *StackSwitcher) GetStack() *Stack {
	retC := C.gtk_stack_switcher_get_stack((*C.GtkStackSwitcher)(recv.native))
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetStack is a wrapper around the C function gtk_stack_switcher_set_stack.
func (recv *StackSwitcher) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(stack.ToC())

	C.gtk_stack_switcher_set_stack((*C.GtkStackSwitcher)(recv.native), c_stack)

	return
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_status_icon_get_geometry : unsupported parameter screen : record with indirection level of 2

// Unsupported : gtk_status_icon_get_gicon : no return generator

// Unsupported : gtk_status_icon_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_style_get : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_style_property : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_valist : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_style_context_add_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// Unsupported : gtk_style_context_get : unsupported parameter ... : varargs

// GetScale is a wrapper around the C function gtk_style_context_get_scale.
func (recv *StyleContext) GetScale() int32 {
	retC := C.gtk_style_context_get_scale((*C.GtkStyleContext)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_style_context_remove_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// SetScale is a wrapper around the C function gtk_style_context_set_scale.
func (recv *StyleContext) SetScale(scale int32) {
	c_scale := (C.gint)(scale)

	C.gtk_style_context_set_scale((*C.GtkStyleContext)(recv.native), c_scale)

	return
}

// Unsupported : gtk_style_context_state_is_running : unsupported parameter progress : no type generator for gdouble, gdouble*

// Unsupported : gtk_style_properties_get : unsupported parameter ... : varargs

// Unsupported : gtk_style_properties_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_properties_set : unsupported parameter ... : varargs

// Unsupported : gtk_style_properties_set_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_table_get_size : unsupported parameter rows : no type generator for guint, guint*

// Unsupported : gtk_text_buffer_create_tag : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_deserialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_get_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_deserialize_set_can_create_tags : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_get_deserialize_formats : unsupported parameter n_formats : no type generator for gint, gint*

// Unsupported : gtk_text_buffer_get_serialize_formats : unsupported parameter n_formats : no type generator for gint, gint*

// Unsupported : gtk_text_buffer_insert_with_tags : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_insert_with_tags_by_name : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_register_deserialize_format : unsupported parameter function : no type generator for TextBufferDeserializeFunc, GtkTextBufferDeserializeFunc

// Unsupported : gtk_text_buffer_register_deserialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_register_serialize_format : unsupported parameter function : no type generator for TextBufferSerializeFunc, GtkTextBufferSerializeFunc

// Unsupported : gtk_text_buffer_register_serialize_tagset : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_serialize : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_unregister_deserialize_format : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_buffer_unregister_serialize_format : unsupported parameter format : Blacklisted record : GdkAtom

// Unsupported : gtk_text_tag_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_text_tag_table_foreach : unsupported parameter func : no type generator for TextTagTableForeach, GtkTextTagTableForeach

// Unsupported : gtk_text_view_buffer_to_window_coords : unsupported parameter window_x : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_cursor_locations : unsupported parameter strong : Blacklisted record : GdkRectangle

// Unsupported : gtk_text_view_get_iter_at_position : unsupported parameter trailing : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_iter_location : unsupported parameter location : Blacklisted record : GdkRectangle

// Unsupported : gtk_text_view_get_line_at_y : unsupported parameter line_top : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_line_yrange : unsupported parameter y : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_text_view_window_to_buffer_coords : unsupported parameter buffer_x : no type generator for gint, gint*

// Unsupported : gtk_theming_engine_get : unsupported parameter ... : varargs

// Unsupported : gtk_theming_engine_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_theming_engine_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_theming_engine_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_theming_engine_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_theming_engine_state_is_running : unsupported parameter progress : no type generator for gdouble, gdouble*

// Unsupported : gtk_tool_item_get_icon_size : no return generator

// Unsupported : gtk_tool_palette_get_icon_size : no return generator

// Unsupported : gtk_tool_palette_set_icon_size : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_icon_from_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_tooltip_set_icon_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_icon_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_tooltip_set_tip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_model_filter_get_model : no return generator

// Unsupported : gtk_tree_model_filter_set_modify_func : unsupported parameter types : no param type

// Unsupported : gtk_tree_model_filter_set_visible_func : unsupported parameter func : no type generator for TreeModelFilterVisibleFunc, GtkTreeModelFilterVisibleFunc

// Unsupported : gtk_tree_model_sort_get_model : no return generator

// Unsupported : gtk_tree_selection_get_select_function : no return generator

// Unsupported : gtk_tree_selection_get_selected : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_selection_get_selected_rows : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_selection_selected_foreach : unsupported parameter func : no type generator for TreeSelectionForeachFunc, GtkTreeSelectionForeachFunc

// Unsupported : gtk_tree_selection_set_select_function : unsupported parameter func : no type generator for TreeSelectionFunc, GtkTreeSelectionFunc

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_insert_with_values : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_insert_with_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_tree_store_reorder : unsupported parameter new_order : no param type

// Unsupported : gtk_tree_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_tree_store_set_valuesv : unsupported parameter columns : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_convert_bin_window_to_tree_coords : unsupported parameter tx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_bin_window_to_widget_coords : unsupported parameter wx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_tree_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_tree_to_widget_coords : unsupported parameter wx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_widget_to_bin_window_coords : unsupported parameter bx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_convert_widget_to_tree_coords : unsupported parameter tx : no type generator for gint, gint*

// Unsupported : gtk_tree_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_tree_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_tree_view_get_background_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_get_cell_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_dest_row_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_drag_dest_row : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_model : no return generator

// Unsupported : gtk_tree_view_get_path_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_row_separator_func : no return generator

// Unsupported : gtk_tree_view_get_search_equal_func : no return generator

// Unsupported : gtk_tree_view_get_search_position_func : no return generator

// Unsupported : gtk_tree_view_get_tooltip_context : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_tree_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_insert_column_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_insert_column_with_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_tree_view_is_blank_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_map_expanded_rows : unsupported parameter func : no type generator for TreeViewMappingFunc, GtkTreeViewMappingFunc

// Unsupported : gtk_tree_view_set_column_drag_function : unsupported parameter func : no type generator for TreeViewColumnDropFunc, GtkTreeViewColumnDropFunc

// Unsupported : gtk_tree_view_set_destroy_count_func : unsupported parameter func : no type generator for TreeDestroyCountFunc, GtkTreeDestroyCountFunc

// Unsupported : gtk_tree_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc, GtkTreeViewRowSeparatorFunc

// Unsupported : gtk_tree_view_set_search_equal_func : unsupported parameter search_equal_func : no type generator for TreeViewSearchEqualFunc, GtkTreeViewSearchEqualFunc

// Unsupported : gtk_tree_view_set_search_position_func : unsupported parameter func : no type generator for TreeViewSearchPositionFunc, GtkTreeViewSearchPositionFunc

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_cell_get_position : unsupported parameter x_offset : no type generator for gint, gint*

// Unsupported : gtk_tree_view_column_cell_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_column_cell_set_cell_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_set_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_set_cell_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_add_tick_callback : unsupported parameter callback : no type generator for TickCallback, GtkTickCallback

// Unsupported : gtk_widget_class_path : unsupported parameter path_length : no type generator for guint, guint*

// Unsupported : gtk_widget_destroyed : unsupported parameter widget_pointer : record with indirection level of 2

// Unsupported : gtk_drag_begin : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_drag_dest_find_target : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_drag_dest_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_get_data : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_drag_source_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_source_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_widget_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_get_action_group : no return generator

// GetAllocatedBaseline is a wrapper around the C function gtk_widget_get_allocated_baseline.
func (recv *Widget) GetAllocatedBaseline() int32 {
	retC := C.gtk_widget_get_allocated_baseline((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_ancestor : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_widget_get_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_clipboard : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_widget_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_and_baseline_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// GetScaleFactor is a wrapper around the C function gtk_widget_get_scale_factor.
func (recv *Widget) GetScaleFactor() int32 {
	retC := C.gtk_widget_get_scale_factor((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_size_request : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_template_child : unsupported parameter widget_type : no type generator for GType, GType

// GetValignWithBaseline is a wrapper around the C function gtk_widget_get_valign_with_baseline.
func (recv *Widget) GetValignWithBaseline() Align {
	retC := C.gtk_widget_get_valign_with_baseline((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// InitTemplate is a wrapper around the C function gtk_widget_init_template.
func (recv *Widget) InitTemplate() {
	C.gtk_widget_init_template((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_insert_action_group : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_widget_intersect : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_list_action_prefixes : no return type

// Unsupported : gtk_widget_path : unsupported parameter path_length : no type generator for guint, guint*

// Unsupported : gtk_widget_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_widget_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_widget_send_expose : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_send_focus_change : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_set_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_set_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_size_allocate : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_size_allocate_with_baseline : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_style_get : unsupported parameter ... : varargs

// Unsupported : gtk_widget_style_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_widget_translate_coordinates : unsupported parameter dest_x : no type generator for gint, gint*

// Close is a wrapper around the C function gtk_window_close.
func (recv *Window) Close() {
	C.gtk_window_close((*C.GtkWindow)(recv.native))

	return
}

// Unsupported : gtk_window_get_default_size : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_window_get_position : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_window_get_size : unsupported parameter width : no type generator for gint, gint*

// SetTitlebar is a wrapper around the C function gtk_window_set_titlebar.
func (recv *Window) SetTitlebar(titlebar *Widget) {
	c_titlebar := (*C.GtkWidget)(titlebar.ToC())

	C.gtk_window_set_titlebar((*C.GtkWindow)(recv.native), c_titlebar)

	return
}

// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
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

// SetAccel is a wrapper around the C function gtk_accel_label_set_accel.
func (recv *AccelLabel) SetAccel(acceleratorKey uint32, acceleratorMods gdk.ModifierType) {
	c_accelerator_key := (C.guint)(acceleratorKey)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	C.gtk_accel_label_set_accel((*C.GtkAccelLabel)(recv.native), c_accelerator_key, c_accelerator_mods)

	return
}

// Unsupported : gtk_action_create_icon : unsupported parameter icon_size : no type generator for gint, GtkIconSize

// Unsupported : gtk_action_get_gicon : no return generator

// Unsupported : gtk_action_set_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries : no param type

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries : no param type

// GetAccelGroup is a wrapper around the C function gtk_action_group_get_accel_group.
func (recv *ActionGroup) GetAccelGroup() *AccelGroup {
	retC := C.gtk_action_group_get_accel_group((*C.GtkActionGroup)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAccelGroup is a wrapper around the C function gtk_action_group_set_accel_group.
func (recv *ActionGroup) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	C.gtk_action_group_set_accel_group((*C.GtkActionGroup)(recv.native), c_accel_group)

	return
}

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GtkTranslateFunc

// Unsupported : gtk_alignment_get_padding : unsupported parameter padding_top : no type generator for guint, guint*

// Unsupported : gtk_app_chooser_button_append_custom_item : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_application_add_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_get_actions_for_accel : no return type

// GetActiveWindow is a wrapper around the C function gtk_application_get_active_window.
func (recv *Application) GetActiveWindow() *Window {
	retC := C.gtk_application_get_active_window((*C.GtkApplication)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindowById is a wrapper around the C function gtk_application_get_window_by_id.
func (recv *Application) GetWindowById(id uint32) *Window {
	c_id := (C.guint)(id)

	retC := C.gtk_application_get_window_by_id((*C.GtkApplication)(recv.native), c_id)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_remove_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels : no param type

// GetId is a wrapper around the C function gtk_application_window_get_id.
func (recv *ApplicationWindow) GetId() uint32 {
	retC := C.gtk_application_window_get_id((*C.GtkApplicationWindow)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_assistant_set_forward_page_func : unsupported parameter page_func : no type generator for AssistantPageFunc, GtkAssistantPageFunc

// Unsupported : gtk_box_query_child_packing : unsupported parameter padding : no type generator for guint, guint*

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback, GCallback

// Unsupported : gtk_builder_add_objects_from_file : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_resource : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_add_objects_from_string : unsupported parameter object_ids : no param type

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc, GtkBuilderConnectFunc

// Unsupported : gtk_builder_extend_with_template : unsupported parameter template_type : no type generator for GType, GType

// Unsupported : gtk_builder_get_type_from_name : no return generator

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// Unsupported : gtk_builder_value_from_string_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_button_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// GetAlwaysShowImage is a wrapper around the C function gtk_button_get_always_show_image.
func (recv *Button) GetAlwaysShowImage() bool {
	retC := C.gtk_button_get_always_show_image((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlwaysShowImage is a wrapper around the C function gtk_button_set_always_show_image.
func (recv *Button) SetAlwaysShowImage(alwaysShow bool) {
	c_always_show :=
		boolToGboolean(alwaysShow)

	C.gtk_button_set_always_show_image((*C.GtkButton)(recv.native), c_always_show)

	return
}

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

// GetAttributes is a wrapper around the C function gtk_entry_get_attributes.
func (recv *Entry) GetAttributes() *pango.AttrList {
	retC := C.gtk_entry_get_attributes((*C.GtkEntry)(recv.native))
	retGo := pango.AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_entry_get_icon_area : unsupported parameter icon_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_get_icon_gicon : no return generator

// GetInputHints is a wrapper around the C function gtk_entry_get_input_hints.
func (recv *Entry) GetInputHints() InputHints {
	retC := C.gtk_entry_get_input_hints((*C.GtkEntry)(recv.native))
	retGo := (InputHints)(retC)

	return retGo
}

// GetInputPurpose is a wrapper around the C function gtk_entry_get_input_purpose.
func (recv *Entry) GetInputPurpose() InputPurpose {
	retC := C.gtk_entry_get_input_purpose((*C.GtkEntry)(recv.native))
	retGo := (InputPurpose)(retC)

	return retGo
}

// Unsupported : gtk_entry_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// SetAttributes is a wrapper around the C function gtk_entry_set_attributes.
func (recv *Entry) SetAttributes(attrs *pango.AttrList) {
	c_attrs := (*C.PangoAttrList)(attrs.ToC())

	C.gtk_entry_set_attributes((*C.GtkEntry)(recv.native), c_attrs)

	return
}

// Unsupported : gtk_entry_set_icon_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SetInputHints is a wrapper around the C function gtk_entry_set_input_hints.
func (recv *Entry) SetInputHints(hints InputHints) {
	c_hints := (C.GtkInputHints)(hints)

	C.gtk_entry_set_input_hints((*C.GtkEntry)(recv.native), c_hints)

	return
}

// SetInputPurpose is a wrapper around the C function gtk_entry_set_input_purpose.
func (recv *Entry) SetInputPurpose(purpose InputPurpose) {
	c_purpose := (C.GtkInputPurpose)(purpose)

	C.gtk_entry_set_input_purpose((*C.GtkEntry)(recv.native), c_purpose)

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

// Unsupported : gtk_hsv_get_color : unsupported parameter h : no type generator for gdouble, gdouble*

// Unsupported : gtk_hsv_get_metrics : unsupported parameter size : no type generator for gint, gint*

// Unsupported : gtk_im_context_get_preedit_string : unsupported parameter attrs : record with indirection level of 2

// Unsupported : gtk_im_context_get_surrounding : unsupported parameter cursor_index : no type generator for gint, gint*

// Unsupported : gtk_im_context_set_cursor_location : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_im_context_simple_add_table : unsupported parameter data : no param type

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : no param type

// Unsupported : gtk_icon_info_get_embedded_rect : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gtk_icon_info_load_icon_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_icon_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_for_context_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gtk_icon_theme_choose_icon : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_get_icon_sizes : no return type

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : no param type

// Unsupported : gtk_icon_theme_lookup_by_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_lookup_by_gicon_for_scale : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

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

// Unsupported : gtk_image_get_gicon : unsupported parameter gicon : no type generator for Gio.Icon, GIcon**

// Unsupported : gtk_image_get_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize*

// Unsupported : gtk_image_get_icon_set : unsupported parameter icon_set : record with indirection level of 2

// Unsupported : gtk_image_get_stock : unsupported parameter size : no type generator for gint, GtkIconSize*

// Unsupported : gtk_image_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_set_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_set_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_info_bar_add_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_label_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_label_get_selection_bounds : unsupported parameter start : no type generator for gint, gint*

// Unsupported : gtk_layout_get_size : unsupported parameter width : no type generator for guint, guint*

// LevelBarNew is a wrapper around the C function gtk_level_bar_new.
func LevelBarNew() *Widget {
	retC := C.gtk_level_bar_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LevelBarNewForInterval is a wrapper around the C function gtk_level_bar_new_for_interval.
func LevelBarNewForInterval(minValue float64, maxValue float64) *Widget {
	c_min_value := (C.gdouble)(minValue)

	c_max_value := (C.gdouble)(maxValue)

	retC := C.gtk_level_bar_new_for_interval(c_min_value, c_max_value)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddOffsetValue is a wrapper around the C function gtk_level_bar_add_offset_value.
func (recv *LevelBar) AddOffsetValue(name string, value float64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (C.gdouble)(value)

	C.gtk_level_bar_add_offset_value((*C.GtkLevelBar)(recv.native), c_name, c_value)

	return
}

// GetMaxValue is a wrapper around the C function gtk_level_bar_get_max_value.
func (recv *LevelBar) GetMaxValue() float64 {
	retC := C.gtk_level_bar_get_max_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetMinValue is a wrapper around the C function gtk_level_bar_get_min_value.
func (recv *LevelBar) GetMinValue() float64 {
	retC := C.gtk_level_bar_get_min_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetMode is a wrapper around the C function gtk_level_bar_get_mode.
func (recv *LevelBar) GetMode() LevelBarMode {
	retC := C.gtk_level_bar_get_mode((*C.GtkLevelBar)(recv.native))
	retGo := (LevelBarMode)(retC)

	return retGo
}

// Unsupported : gtk_level_bar_get_offset_value : unsupported parameter value : no type generator for gdouble, gdouble*

// GetValue is a wrapper around the C function gtk_level_bar_get_value.
func (recv *LevelBar) GetValue() float64 {
	retC := C.gtk_level_bar_get_value((*C.GtkLevelBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// RemoveOffsetValue is a wrapper around the C function gtk_level_bar_remove_offset_value.
func (recv *LevelBar) RemoveOffsetValue(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_level_bar_remove_offset_value((*C.GtkLevelBar)(recv.native), c_name)

	return
}

// SetMaxValue is a wrapper around the C function gtk_level_bar_set_max_value.
func (recv *LevelBar) SetMaxValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_max_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// SetMinValue is a wrapper around the C function gtk_level_bar_set_min_value.
func (recv *LevelBar) SetMinValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_min_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// SetMode is a wrapper around the C function gtk_level_bar_set_mode.
func (recv *LevelBar) SetMode(mode LevelBarMode) {
	c_mode := (C.GtkLevelBarMode)(mode)

	C.gtk_level_bar_set_mode((*C.GtkLevelBar)(recv.native), c_mode)

	return
}

// SetValue is a wrapper around the C function gtk_level_bar_set_value.
func (recv *LevelBar) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_level_bar_set_value((*C.GtkLevelBar)(recv.native), c_value)

	return
}

// Unsupported : gtk_list_box_bind_model : unsupported parameter model : no type generator for Gio.ListModel, GListModel*

// Unsupported : gtk_list_box_selected_foreach : unsupported parameter func : no type generator for ListBoxForeachFunc, GtkListBoxForeachFunc

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc, GtkListBoxFilterFunc

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc, GtkListBoxUpdateHeaderFunc

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc, GtkListBoxSortFunc

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

// MenuButtonNew is a wrapper around the C function gtk_menu_button_new.
func MenuButtonNew() *Widget {
	retC := C.gtk_menu_button_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAlignWidget is a wrapper around the C function gtk_menu_button_get_align_widget.
func (recv *MenuButton) GetAlignWidget() *Widget {
	retC := C.gtk_menu_button_get_align_widget((*C.GtkMenuButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDirection is a wrapper around the C function gtk_menu_button_get_direction.
func (recv *MenuButton) GetDirection() ArrowType {
	retC := C.gtk_menu_button_get_direction((*C.GtkMenuButton)(recv.native))
	retGo := (ArrowType)(retC)

	return retGo
}

// GetMenuModel is a wrapper around the C function gtk_menu_button_get_menu_model.
func (recv *MenuButton) GetMenuModel() *gio.MenuModel {
	retC := C.gtk_menu_button_get_menu_model((*C.GtkMenuButton)(recv.native))
	retGo := gio.MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPopup is a wrapper around the C function gtk_menu_button_get_popup.
func (recv *MenuButton) GetPopup() *Menu {
	retC := C.gtk_menu_button_get_popup((*C.GtkMenuButton)(recv.native))
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAlignWidget is a wrapper around the C function gtk_menu_button_set_align_widget.
func (recv *MenuButton) SetAlignWidget(alignWidget *Widget) {
	c_align_widget := (*C.GtkWidget)(alignWidget.ToC())

	C.gtk_menu_button_set_align_widget((*C.GtkMenuButton)(recv.native), c_align_widget)

	return
}

// SetDirection is a wrapper around the C function gtk_menu_button_set_direction.
func (recv *MenuButton) SetDirection(direction ArrowType) {
	c_direction := (C.GtkArrowType)(direction)

	C.gtk_menu_button_set_direction((*C.GtkMenuButton)(recv.native), c_direction)

	return
}

// SetMenuModel is a wrapper around the C function gtk_menu_button_set_menu_model.
func (recv *MenuButton) SetMenuModel(menuModel *gio.MenuModel) {
	c_menu_model := (*C.GMenuModel)(menuModel.ToC())

	C.gtk_menu_button_set_menu_model((*C.GtkMenuButton)(recv.native), c_menu_model)

	return
}

// SetPopup is a wrapper around the C function gtk_menu_button_set_popup.
func (recv *MenuButton) SetPopup(menu *Widget) {
	c_menu := (*C.GtkWidget)(menu.ToC())

	C.gtk_menu_button_set_popup((*C.GtkMenuButton)(recv.native), c_menu)

	return
}

// Unsupported : gtk_menu_item_toggle_size_request : unsupported parameter requisition : no type generator for gint, gint*

// BindModel is a wrapper around the C function gtk_menu_shell_bind_model.
func (recv *MenuShell) BindModel(model *gio.MenuModel, actionNamespace string, withSeparators bool) {
	c_model := (*C.GMenuModel)(model.ToC())

	c_action_namespace := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(c_action_namespace))

	c_with_separators :=
		boolToGboolean(withSeparators)

	C.gtk_menu_shell_bind_model((*C.GtkMenuShell)(recv.native), c_model, c_action_namespace, c_with_separators)

	return
}

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

// Unsupported : gtk_places_sidebar_add_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_get_location : no return generator

// Unsupported : gtk_places_sidebar_get_nth_bookmark : no return generator

// Unsupported : gtk_places_sidebar_remove_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_set_location : unsupported parameter location : no type generator for Gio.File, GFile*

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

// Unsupported : gtk_scale_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons : no param type

// Unsupported : gtk_scrolled_window_get_policy : unsupported parameter hscrollbar_policy : GtkPolicyType* with indirection level of 1

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// SearchEntryNew is a wrapper around the C function gtk_search_entry_new.
func SearchEntryNew() *Widget {
	retC := C.gtk_search_entry_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Blacklisted : GtkSocket

// Unsupported : gtk_spin_button_get_increments : unsupported parameter step : no type generator for gdouble, gdouble*

// Unsupported : gtk_spin_button_get_range : unsupported parameter min : no type generator for gdouble, gdouble*

// Blacklisted : GtkStackAccessible

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

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_style_context_remove_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

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

// GetInputHints is a wrapper around the C function gtk_text_view_get_input_hints.
func (recv *TextView) GetInputHints() InputHints {
	retC := C.gtk_text_view_get_input_hints((*C.GtkTextView)(recv.native))
	retGo := (InputHints)(retC)

	return retGo
}

// GetInputPurpose is a wrapper around the C function gtk_text_view_get_input_purpose.
func (recv *TextView) GetInputPurpose() InputPurpose {
	retC := C.gtk_text_view_get_input_purpose((*C.GtkTextView)(recv.native))
	retGo := (InputPurpose)(retC)

	return retGo
}

// Unsupported : gtk_text_view_get_iter_at_position : unsupported parameter trailing : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_iter_location : unsupported parameter location : Blacklisted record : GdkRectangle

// Unsupported : gtk_text_view_get_line_at_y : unsupported parameter line_top : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_line_yrange : unsupported parameter y : no type generator for gint, gint*

// Unsupported : gtk_text_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// SetInputHints is a wrapper around the C function gtk_text_view_set_input_hints.
func (recv *TextView) SetInputHints(hints InputHints) {
	c_hints := (C.GtkInputHints)(hints)

	C.gtk_text_view_set_input_hints((*C.GtkTextView)(recv.native), c_hints)

	return
}

// SetInputPurpose is a wrapper around the C function gtk_text_view_set_input_purpose.
func (recv *TextView) SetInputPurpose(purpose InputPurpose) {
	c_purpose := (C.GtkInputPurpose)(purpose)

	C.gtk_text_view_set_input_purpose((*C.GtkTextView)(recv.native), c_purpose)

	return
}

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

// Unsupported : gtk_widget_get_size_request : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_template_child : unsupported parameter widget_type : no type generator for GType, GType

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

// Unsupported : gtk_window_get_default_size : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_window_get_position : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_window_get_size : unsupported parameter width : no type generator for gint, gint*

// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
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

// GetIsLocked is a wrapper around the C function gtk_accel_group_get_is_locked.
func (recv *AccelGroup) GetIsLocked() bool {
	retC := C.gtk_accel_group_get_is_locked((*C.GtkAccelGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetModifierMask is a wrapper around the C function gtk_accel_group_get_modifier_mask.
func (recv *AccelGroup) GetModifierMask() gdk.ModifierType {
	retC := C.gtk_accel_group_get_modifier_mask((*C.GtkAccelGroup)(recv.native))
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

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

// Configure is a wrapper around the C function gtk_adjustment_configure.
func (recv *Adjustment) Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) {
	c_value := (C.gdouble)(value)

	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	c_step_increment := (C.gdouble)(stepIncrement)

	c_page_increment := (C.gdouble)(pageIncrement)

	c_page_size := (C.gdouble)(pageSize)

	C.gtk_adjustment_configure((*C.GtkAdjustment)(recv.native), c_value, c_lower, c_upper, c_step_increment, c_page_increment, c_page_size)

	return
}

// GetLower is a wrapper around the C function gtk_adjustment_get_lower.
func (recv *Adjustment) GetLower() float64 {
	retC := C.gtk_adjustment_get_lower((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetPageIncrement is a wrapper around the C function gtk_adjustment_get_page_increment.
func (recv *Adjustment) GetPageIncrement() float64 {
	retC := C.gtk_adjustment_get_page_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetPageSize is a wrapper around the C function gtk_adjustment_get_page_size.
func (recv *Adjustment) GetPageSize() float64 {
	retC := C.gtk_adjustment_get_page_size((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetStepIncrement is a wrapper around the C function gtk_adjustment_get_step_increment.
func (recv *Adjustment) GetStepIncrement() float64 {
	retC := C.gtk_adjustment_get_step_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetUpper is a wrapper around the C function gtk_adjustment_get_upper.
func (recv *Adjustment) GetUpper() float64 {
	retC := C.gtk_adjustment_get_upper((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetLower is a wrapper around the C function gtk_adjustment_set_lower.
func (recv *Adjustment) SetLower(lower float64) {
	c_lower := (C.gdouble)(lower)

	C.gtk_adjustment_set_lower((*C.GtkAdjustment)(recv.native), c_lower)

	return
}

// SetPageIncrement is a wrapper around the C function gtk_adjustment_set_page_increment.
func (recv *Adjustment) SetPageIncrement(pageIncrement float64) {
	c_page_increment := (C.gdouble)(pageIncrement)

	C.gtk_adjustment_set_page_increment((*C.GtkAdjustment)(recv.native), c_page_increment)

	return
}

// SetPageSize is a wrapper around the C function gtk_adjustment_set_page_size.
func (recv *Adjustment) SetPageSize(pageSize float64) {
	c_page_size := (C.gdouble)(pageSize)

	C.gtk_adjustment_set_page_size((*C.GtkAdjustment)(recv.native), c_page_size)

	return
}

// SetStepIncrement is a wrapper around the C function gtk_adjustment_set_step_increment.
func (recv *Adjustment) SetStepIncrement(stepIncrement float64) {
	c_step_increment := (C.gdouble)(stepIncrement)

	C.gtk_adjustment_set_step_increment((*C.GtkAdjustment)(recv.native), c_step_increment)

	return
}

// SetUpper is a wrapper around the C function gtk_adjustment_set_upper.
func (recv *Adjustment) SetUpper(upper float64) {
	c_upper := (C.gdouble)(upper)

	C.gtk_adjustment_set_upper((*C.GtkAdjustment)(recv.native), c_upper)

	return
}

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

// Unsupported : gtk_calendar_get_date : unsupported parameter year : no type generator for guint, guint*

// GetDetailHeightRows is a wrapper around the C function gtk_calendar_get_detail_height_rows.
func (recv *Calendar) GetDetailHeightRows() int32 {
	retC := C.gtk_calendar_get_detail_height_rows((*C.GtkCalendar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDetailWidthChars is a wrapper around the C function gtk_calendar_get_detail_width_chars.
func (recv *Calendar) GetDetailWidthChars() int32 {
	retC := C.gtk_calendar_get_detail_width_chars((*C.GtkCalendar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc, GtkCalendarDetailFunc

// SetDetailHeightRows is a wrapper around the C function gtk_calendar_set_detail_height_rows.
func (recv *Calendar) SetDetailHeightRows(rows int32) {
	c_rows := (C.gint)(rows)

	C.gtk_calendar_set_detail_height_rows((*C.GtkCalendar)(recv.native), c_rows)

	return
}

// SetDetailWidthChars is a wrapper around the C function gtk_calendar_set_detail_width_chars.
func (recv *Calendar) SetDetailWidthChars(chars int32) {
	c_chars := (C.gint)(chars)

	C.gtk_calendar_set_detail_width_chars((*C.GtkCalendar)(recv.native), c_chars)

	return
}

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

// WaitIsUrisAvailable is a wrapper around the C function gtk_clipboard_wait_is_uris_available.
func (recv *Clipboard) WaitIsUrisAvailable() bool {
	retC := C.gtk_clipboard_wait_is_uris_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetColorSelection is a wrapper around the C function gtk_color_selection_dialog_get_color_selection.
func (recv *ColorSelectionDialog) GetColorSelection() *Widget {
	retC := C.gtk_color_selection_dialog_get_color_selection((*C.GtkColorSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetButtonSensitivity is a wrapper around the C function gtk_combo_box_get_button_sensitivity.
func (recv *ComboBox) GetButtonSensitivity() SensitivityType {
	retC := C.gtk_combo_box_get_button_sensitivity((*C.GtkComboBox)(recv.native))
	retGo := (SensitivityType)(retC)

	return retGo
}

// Unsupported : gtk_combo_box_get_model : no return generator

// Unsupported : gtk_combo_box_get_row_separator_func : no return generator

// SetButtonSensitivity is a wrapper around the C function gtk_combo_box_set_button_sensitivity.
func (recv *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	c_sensitivity := (C.GtkSensitivityType)(sensitivity)

	C.gtk_combo_box_set_button_sensitivity((*C.GtkComboBox)(recv.native), c_sensitivity)

	return
}

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

// GetFocusChild is a wrapper around the C function gtk_container_get_focus_child.
func (recv *Container) GetFocusChild() *Widget {
	retC := C.gtk_container_get_focus_child((*C.GtkContainer)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_css_provider_load_from_data : unsupported parameter data : no param type

// Unsupported : gtk_css_provider_load_from_file : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_add_buttons : unsupported parameter ... : varargs

// GetActionArea is a wrapper around the C function gtk_dialog_get_action_area.
func (recv *Dialog) GetActionArea() *Widget {
	retC := C.gtk_dialog_get_action_area((*C.GtkDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContentArea is a wrapper around the C function gtk_dialog_get_content_area.
func (recv *Dialog) GetContentArea() *Box {
	retC := C.gtk_dialog_get_content_area((*C.GtkDialog)(recv.native))
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_dialog_set_alternative_button_order : unsupported parameter ... : varargs

// Unsupported : gtk_dialog_set_alternative_button_order_from_array : unsupported parameter new_order : no param type

// Unsupported : gtk_entry_get_icon_area : unsupported parameter icon_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_get_icon_gicon : no return generator

// Unsupported : gtk_entry_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// GetOverwriteMode is a wrapper around the C function gtk_entry_get_overwrite_mode.
func (recv *Entry) GetOverwriteMode() bool {
	retC := C.gtk_entry_get_overwrite_mode((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// GetTextLength is a wrapper around the C function gtk_entry_get_text_length.
func (recv *Entry) GetTextLength() uint16 {
	retC := C.gtk_entry_get_text_length((*C.GtkEntry)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Unsupported : gtk_entry_set_icon_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SetOverwriteMode is a wrapper around the C function gtk_entry_set_overwrite_mode.
func (recv *Entry) SetOverwriteMode(overwrite bool) {
	c_overwrite :=
		boolToGboolean(overwrite)

	C.gtk_entry_set_overwrite_mode((*C.GtkEntry)(recv.native), c_overwrite)

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

// GetFace is a wrapper around the C function gtk_font_selection_get_face.
func (recv *FontSelection) GetFace() *pango.FontFace {
	retC := C.gtk_font_selection_get_face((*C.GtkFontSelection)(recv.native))
	retGo := pango.FontFaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFaceList is a wrapper around the C function gtk_font_selection_get_face_list.
func (recv *FontSelection) GetFaceList() *Widget {
	retC := C.gtk_font_selection_get_face_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamily is a wrapper around the C function gtk_font_selection_get_family.
func (recv *FontSelection) GetFamily() *pango.FontFamily {
	retC := C.gtk_font_selection_get_family((*C.GtkFontSelection)(recv.native))
	retGo := pango.FontFamilyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamilyList is a wrapper around the C function gtk_font_selection_get_family_list.
func (recv *FontSelection) GetFamilyList() *Widget {
	retC := C.gtk_font_selection_get_family_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewEntry is a wrapper around the C function gtk_font_selection_get_preview_entry.
func (recv *FontSelection) GetPreviewEntry() *Widget {
	retC := C.gtk_font_selection_get_preview_entry((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSize is a wrapper around the C function gtk_font_selection_get_size.
func (recv *FontSelection) GetSize() int32 {
	retC := C.gtk_font_selection_get_size((*C.GtkFontSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSizeEntry is a wrapper around the C function gtk_font_selection_get_size_entry.
func (recv *FontSelection) GetSizeEntry() *Widget {
	retC := C.gtk_font_selection_get_size_entry((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSizeList is a wrapper around the C function gtk_font_selection_get_size_list.
func (recv *FontSelection) GetSizeList() *Widget {
	retC := C.gtk_font_selection_get_size_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCancelButton is a wrapper around the C function gtk_font_selection_dialog_get_cancel_button.
func (recv *FontSelectionDialog) GetCancelButton() *Widget {
	retC := C.gtk_font_selection_dialog_get_cancel_button((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOkButton is a wrapper around the C function gtk_font_selection_dialog_get_ok_button.
func (recv *FontSelectionDialog) GetOkButton() *Widget {
	retC := C.gtk_font_selection_dialog_get_ok_button((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// HSVNew is a wrapper around the C function gtk_hsv_new.
func HSVNew() *Widget {
	retC := C.gtk_hsv_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_hsv_get_color : unsupported parameter h : no type generator for gdouble, gdouble*

// Unsupported : gtk_hsv_get_metrics : unsupported parameter size : no type generator for gint, gint*

// IsAdjusting is a wrapper around the C function gtk_hsv_is_adjusting.
func (recv *HSV) IsAdjusting() bool {
	retC := C.gtk_hsv_is_adjusting((*C.GtkHSV)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetColor is a wrapper around the C function gtk_hsv_set_color.
func (recv *HSV) SetColor(h float64, s float64, v float64) {
	c_h := (C.double)(h)

	c_s := (C.double)(s)

	c_v := (C.double)(v)

	C.gtk_hsv_set_color((*C.GtkHSV)(recv.native), c_h, c_s, c_v)

	return
}

// SetMetrics is a wrapper around the C function gtk_hsv_set_metrics.
func (recv *HSV) SetMetrics(size int32, ringWidth int32) {
	c_size := (C.gint)(size)

	c_ring_width := (C.gint)(ringWidth)

	C.gtk_hsv_set_metrics((*C.GtkHSV)(recv.native), c_size, c_ring_width)

	return
}

// GetChildDetached is a wrapper around the C function gtk_handle_box_get_child_detached.
func (recv *HandleBox) GetChildDetached() bool {
	retC := C.gtk_handle_box_get_child_detached((*C.GtkHandleBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_im_context_get_preedit_string : unsupported parameter attrs : record with indirection level of 2

// Unsupported : gtk_im_context_get_surrounding : unsupported parameter cursor_index : no type generator for gint, gint*

// Unsupported : gtk_im_context_set_cursor_location : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_im_context_simple_add_table : unsupported parameter data : no param type

// IconInfoNewForPixbuf is a wrapper around the C function gtk_icon_info_new_for_pixbuf.
func IconInfoNewForPixbuf(iconTheme *IconTheme, pixbuf *gdkpixbuf.Pixbuf) *IconInfo {
	c_icon_theme := (*C.GtkIconTheme)(iconTheme.ToC())

	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	retC := C.gtk_icon_info_new_for_pixbuf(c_icon_theme, c_pixbuf)
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetBinWindow is a wrapper around the C function gtk_layout_get_bin_window.
func (recv *Layout) GetBinWindow() *gdk.Window {
	retC := C.gtk_layout_get_bin_window((*C.GtkLayout)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_layout_get_size : unsupported parameter width : no type generator for guint, guint*

// Unsupported : gtk_level_bar_get_offset_value : unsupported parameter value : no type generator for gdouble, gdouble*

// GetVisited is a wrapper around the C function gtk_link_button_get_visited.
func (recv *LinkButton) GetVisited() bool {
	retC := C.gtk_link_button_get_visited((*C.GtkLinkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetVisited is a wrapper around the C function gtk_link_button_set_visited.
func (recv *LinkButton) SetVisited(visited bool) {
	c_visited :=
		boolToGboolean(visited)

	C.gtk_link_button_set_visited((*C.GtkLinkButton)(recv.native), c_visited)

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

// GetAccelPath is a wrapper around the C function gtk_menu_get_accel_path.
func (recv *Menu) GetAccelPath() string {
	retC := C.gtk_menu_get_accel_path((*C.GtkMenu)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMonitor is a wrapper around the C function gtk_menu_get_monitor.
func (recv *Menu) GetMonitor() int32 {
	retC := C.gtk_menu_get_monitor((*C.GtkMenu)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_menu_popup : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// Unsupported : gtk_menu_popup_at_pointer : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_at_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_menu_popup_at_widget : unsupported parameter trigger_event : no type generator for Gdk.Event, const GdkEvent*

// Unsupported : gtk_menu_popup_for_device : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// GetAccelPath is a wrapper around the C function gtk_menu_item_get_accel_path.
func (recv *MenuItem) GetAccelPath() string {
	retC := C.gtk_menu_item_get_accel_path((*C.GtkMenuItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_menu_item_toggle_size_request : unsupported parameter requisition : no type generator for gint, gint*

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_text : unsupported parameter ... : varargs

// GetImage is a wrapper around the C function gtk_message_dialog_get_image.
func (recv *MessageDialog) GetImage() *Widget {
	retC := C.gtk_message_dialog_get_image((*C.GtkMessageDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_misc_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_misc_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// MountOperationNew is a wrapper around the C function gtk_mount_operation_new.
func MountOperationNew(parent *Window) *gio.MountOperation {
	c_parent := (*C.GtkWindow)(parent.ToC())

	retC := C.gtk_mount_operation_new(c_parent)
	retGo := gio.MountOperationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetParent is a wrapper around the C function gtk_mount_operation_get_parent.
func (recv *MountOperation) GetParent() *Window {
	retC := C.gtk_mount_operation_get_parent((*C.GtkMountOperation)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_mount_operation_get_screen.
func (recv *MountOperation) GetScreen() *gdk.Screen {
	retC := C.gtk_mount_operation_get_screen((*C.GtkMountOperation)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsShowing is a wrapper around the C function gtk_mount_operation_is_showing.
func (recv *MountOperation) IsShowing() bool {
	retC := C.gtk_mount_operation_is_showing((*C.GtkMountOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetParent is a wrapper around the C function gtk_mount_operation_set_parent.
func (recv *MountOperation) SetParent(parent *Window) {
	c_parent := (*C.GtkWindow)(parent.ToC())

	C.gtk_mount_operation_set_parent((*C.GtkMountOperation)(recv.native), c_parent)

	return
}

// SetScreen is a wrapper around the C function gtk_mount_operation_set_screen.
func (recv *MountOperation) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_mount_operation_set_screen((*C.GtkMountOperation)(recv.native), c_screen)

	return
}

// Unsupported : gtk_numerable_icon_get_background_gicon : no return generator

// Unsupported : gtk_numerable_icon_set_background_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_pad_controller_set_action_entries : unsupported parameter entries : no param type

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// LoadFile is a wrapper around the C function gtk_page_setup_load_file.
func (recv *PageSetup) LoadFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_load_file((*C.GtkPageSetup)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LoadKeyFile is a wrapper around the C function gtk_page_setup_load_key_file.
func (recv *PageSetup) LoadKeyFile(keyFile *glib.KeyFile, groupName string) (bool, error) {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_load_key_file((*C.GtkPageSetup)(recv.native), c_key_file, c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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

// GetNumberUpLayout is a wrapper around the C function gtk_print_settings_get_number_up_layout.
func (recv *PrintSettings) GetNumberUpLayout() NumberUpLayout {
	retC := C.gtk_print_settings_get_number_up_layout((*C.GtkPrintSettings)(recv.native))
	retGo := (NumberUpLayout)(retC)

	return retGo
}

// Unsupported : gtk_print_settings_get_page_ranges : unsupported parameter num_ranges : no type generator for gint, gint*

// LoadFile is a wrapper around the C function gtk_print_settings_load_file.
func (recv *PrintSettings) LoadFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_load_file((*C.GtkPrintSettings)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LoadKeyFile is a wrapper around the C function gtk_print_settings_load_key_file.
func (recv *PrintSettings) LoadKeyFile(keyFile *glib.KeyFile, groupName string) (bool, error) {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_load_key_file((*C.GtkPrintSettings)(recv.native), c_key_file, c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetNumberUpLayout is a wrapper around the C function gtk_print_settings_set_number_up_layout.
func (recv *PrintSettings) SetNumberUpLayout(numberUpLayout NumberUpLayout) {
	c_number_up_layout := (C.GtkNumberUpLayout)(numberUpLayout)

	C.gtk_print_settings_set_number_up_layout((*C.GtkPrintSettings)(recv.native), c_number_up_layout)

	return
}

// Unsupported : gtk_print_settings_set_page_ranges : unsupported parameter page_ranges : no param type

// Unsupported : gtk_print_settings_to_gvariant : return type : Blacklisted record : GVariant

// Unsupported : gtk_range_get_range_rect : unsupported parameter range_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_range_get_slider_range : unsupported parameter slider_start : no type generator for gint, gint*

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_recent_filter_add_custom : unsupported parameter func : no type generator for RecentFilterFunc, GtkRecentFilterFunc

// Unsupported : gtk_scale_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// GetMinusButton is a wrapper around the C function gtk_scale_button_get_minus_button.
func (recv *ScaleButton) GetMinusButton() *Button {
	retC := C.gtk_scale_button_get_minus_button((*C.GtkScaleButton)(recv.native))
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPlusButton is a wrapper around the C function gtk_scale_button_get_plus_button.
func (recv *ScaleButton) GetPlusButton() *Button {
	retC := C.gtk_scale_button_get_plus_button((*C.GtkScaleButton)(recv.native))
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPopup is a wrapper around the C function gtk_scale_button_get_popup.
func (recv *ScaleButton) GetPopup() *Widget {
	retC := C.gtk_scale_button_get_popup((*C.GtkScaleButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons : no param type

// Unsupported : gtk_scrolled_window_get_policy : unsupported parameter hscrollbar_policy : GtkPolicyType* with indirection level of 1

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Blacklisted : GtkSocket

// Unsupported : gtk_spin_button_get_increments : unsupported parameter step : no type generator for gdouble, gdouble*

// Unsupported : gtk_spin_button_get_range : unsupported parameter min : no type generator for gdouble, gdouble*

// Blacklisted : GtkStackAccessible

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_status_icon_get_geometry : unsupported parameter screen : record with indirection level of 2

// Unsupported : gtk_status_icon_get_gicon : no return generator

// GetX11WindowId is a wrapper around the C function gtk_status_icon_get_x11_window_id.
func (recv *StatusIcon) GetX11WindowId() uint32 {
	retC := C.gtk_status_icon_get_x11_window_id((*C.GtkStatusIcon)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

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

// ToolbarReconfigured is a wrapper around the C function gtk_tool_item_toolbar_reconfigured.
func (recv *ToolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured((*C.GtkToolItem)(recv.native))

	return
}

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

// GetWindow is a wrapper around the C function gtk_widget_get_window.
func (recv *Widget) GetWindow() *gdk.Window {
	retC := C.gtk_widget_get_window((*C.GtkWidget)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
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

// Unsupported : gtk_window_get_default_size : unsupported parameter width : no type generator for gint, gint*

// GetDefaultWidget is a wrapper around the C function gtk_window_get_default_widget.
func (recv *Window) GetDefaultWidget() *Widget {
	retC := C.gtk_window_get_default_widget((*C.GtkWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_window_get_position : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_window_get_size : unsupported parameter width : no type generator for gint, gint*

// ListWindows is a wrapper around the C function gtk_window_group_list_windows.
func (recv *WindowGroup) ListWindows() *glib.List {
	retC := C.gtk_window_group_list_windows((*C.GtkWindowGroup)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

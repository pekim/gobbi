// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// GetLicenseType is a wrapper around the C function gtk_about_dialog_get_license_type.
func (recv *AboutDialog) GetLicenseType() License {
	retC := C.gtk_about_dialog_get_license_type((*C.GtkAboutDialog)(recv.native))
	retGo := (License)(retC)

	return retGo
}

// Unsupported : gtk_about_dialog_set_artists : unsupported parameter artists : no param type

// Unsupported : gtk_about_dialog_set_authors : unsupported parameter authors : no param type

// Unsupported : gtk_about_dialog_set_documenters : unsupported parameter documenters : no param type

// SetLicenseType is a wrapper around the C function gtk_about_dialog_set_license_type.
func (recv *AboutDialog) SetLicenseType(licenseType License) {
	c_license_type := (C.GtkLicense)(licenseType)

	C.gtk_about_dialog_set_license_type((*C.GtkAboutDialog)(recv.native), c_license_type)

	return
}

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

// AppChooserButtonNew is a wrapper around the C function gtk_app_chooser_button_new.
func AppChooserButtonNew(contentType string) *Widget {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_button_new(c_content_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_app_chooser_button_append_custom_item : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// AppendSeparator is a wrapper around the C function gtk_app_chooser_button_append_separator.
func (recv *AppChooserButton) AppendSeparator() {
	C.gtk_app_chooser_button_append_separator((*C.GtkAppChooserButton)(recv.native))

	return
}

// GetShowDialogItem is a wrapper around the C function gtk_app_chooser_button_get_show_dialog_item.
func (recv *AppChooserButton) GetShowDialogItem() bool {
	retC := C.gtk_app_chooser_button_get_show_dialog_item((*C.GtkAppChooserButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActiveCustomItem is a wrapper around the C function gtk_app_chooser_button_set_active_custom_item.
func (recv *AppChooserButton) SetActiveCustomItem(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_app_chooser_button_set_active_custom_item((*C.GtkAppChooserButton)(recv.native), c_name)

	return
}

// SetShowDialogItem is a wrapper around the C function gtk_app_chooser_button_set_show_dialog_item.
func (recv *AppChooserButton) SetShowDialogItem(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_button_set_show_dialog_item((*C.GtkAppChooserButton)(recv.native), c_setting)

	return
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// AppChooserDialogNewForContentType is a wrapper around the C function gtk_app_chooser_dialog_new_for_content_type.
func AppChooserDialogNewForContentType(parent *Window, flags DialogFlags, contentType string) *Widget {
	c_parent := (*C.GtkWindow)(parent.ToC())

	c_flags := (C.GtkDialogFlags)(flags)

	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_dialog_new_for_content_type(c_parent, c_flags, c_content_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidget is a wrapper around the C function gtk_app_chooser_dialog_get_widget.
func (recv *AppChooserDialog) GetWidget() *Widget {
	retC := C.gtk_app_chooser_dialog_get_widget((*C.GtkAppChooserDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppChooserWidgetNew is a wrapper around the C function gtk_app_chooser_widget_new.
func AppChooserWidgetNew(contentType string) *Widget {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_widget_new(c_content_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultText is a wrapper around the C function gtk_app_chooser_widget_get_default_text.
func (recv *AppChooserWidget) GetDefaultText() string {
	retC := C.gtk_app_chooser_widget_get_default_text((*C.GtkAppChooserWidget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetShowAll is a wrapper around the C function gtk_app_chooser_widget_get_show_all.
func (recv *AppChooserWidget) GetShowAll() bool {
	retC := C.gtk_app_chooser_widget_get_show_all((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowDefault is a wrapper around the C function gtk_app_chooser_widget_get_show_default.
func (recv *AppChooserWidget) GetShowDefault() bool {
	retC := C.gtk_app_chooser_widget_get_show_default((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowFallback is a wrapper around the C function gtk_app_chooser_widget_get_show_fallback.
func (recv *AppChooserWidget) GetShowFallback() bool {
	retC := C.gtk_app_chooser_widget_get_show_fallback((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowOther is a wrapper around the C function gtk_app_chooser_widget_get_show_other.
func (recv *AppChooserWidget) GetShowOther() bool {
	retC := C.gtk_app_chooser_widget_get_show_other((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowRecommended is a wrapper around the C function gtk_app_chooser_widget_get_show_recommended.
func (recv *AppChooserWidget) GetShowRecommended() bool {
	retC := C.gtk_app_chooser_widget_get_show_recommended((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowAll is a wrapper around the C function gtk_app_chooser_widget_set_show_all.
func (recv *AppChooserWidget) SetShowAll(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_all((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowDefault is a wrapper around the C function gtk_app_chooser_widget_set_show_default.
func (recv *AppChooserWidget) SetShowDefault(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_default((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowFallback is a wrapper around the C function gtk_app_chooser_widget_set_show_fallback.
func (recv *AppChooserWidget) SetShowFallback(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_fallback((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowOther is a wrapper around the C function gtk_app_chooser_widget_set_show_other.
func (recv *AppChooserWidget) SetShowOther(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_other((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowRecommended is a wrapper around the C function gtk_app_chooser_widget_set_show_recommended.
func (recv *AppChooserWidget) SetShowRecommended(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_recommended((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// ApplicationNew is a wrapper around the C function gtk_application_new.
func ApplicationNew(applicationId string, flags gio.ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.gtk_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_application_add_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// AddWindow is a wrapper around the C function gtk_application_add_window.
func (recv *Application) AddWindow(window *Window) {
	c_window := (*C.GtkWindow)(window.ToC())

	C.gtk_application_add_window((*C.GtkApplication)(recv.native), c_window)

	return
}

// Unsupported : gtk_application_get_accels_for_action : no return type

// Unsupported : gtk_application_get_actions_for_accel : no return type

// GetWindows is a wrapper around the C function gtk_application_get_windows.
func (recv *Application) GetWindows() *glib.List {
	retC := C.gtk_application_get_windows((*C.GtkApplication)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_application_list_action_descriptions : no return type

// Unsupported : gtk_application_remove_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// RemoveWindow is a wrapper around the C function gtk_application_remove_window.
func (recv *Application) RemoveWindow(window *Window) {
	c_window := (*C.GtkWindow)(window.ToC())

	C.gtk_application_remove_window((*C.GtkApplication)(recv.native), c_window)

	return
}

// Unsupported : gtk_application_set_accels_for_action : unsupported parameter accels : no param type

// NextPage is a wrapper around the C function gtk_assistant_next_page.
func (recv *Assistant) NextPage() {
	C.gtk_assistant_next_page((*C.GtkAssistant)(recv.native))

	return
}

// PreviousPage is a wrapper around the C function gtk_assistant_previous_page.
func (recv *Assistant) PreviousPage() {
	C.gtk_assistant_previous_page((*C.GtkAssistant)(recv.native))

	return
}

// Unsupported : gtk_assistant_set_forward_page_func : unsupported parameter page_func : no type generator for AssistantPageFunc, GtkAssistantPageFunc

// BoxNew is a wrapper around the C function gtk_box_new.
func BoxNew(orientation Orientation, spacing int32) *Widget {
	c_orientation := (C.GtkOrientation)(orientation)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_box_new(c_orientation, c_spacing)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// ButtonBoxNew is a wrapper around the C function gtk_button_box_new.
func ButtonBoxNew(orientation Orientation) *Widget {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_button_box_new(c_orientation)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_calendar_get_date : unsupported parameter year : no type generator for guint, guint*

// GetDayIsMarked is a wrapper around the C function gtk_calendar_get_day_is_marked.
func (recv *Calendar) GetDayIsMarked(day uint32) bool {
	c_day := (C.guint)(day)

	retC := C.gtk_calendar_get_day_is_marked((*C.GtkCalendar)(recv.native), c_day)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc, GtkCalendarDetailFunc

// Unsupported : gtk_cell_area_activate : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_activate_cell : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Add is a wrapper around the C function gtk_cell_area_add.
func (recv *CellArea) Add(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	C.gtk_cell_area_add((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// AddFocusSibling is a wrapper around the C function gtk_cell_area_add_focus_sibling.
func (recv *CellArea) AddFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_sibling := (*C.GtkCellRenderer)(sibling.ToC())

	C.gtk_cell_area_add_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)

	return
}

// Unsupported : gtk_cell_area_add_with_properties : unsupported parameter ... : varargs

// Unsupported : gtk_cell_area_apply_attributes : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// AttributeConnect is a wrapper around the C function gtk_cell_area_attribute_connect.
func (recv *CellArea) AttributeConnect(renderer *CellRenderer, attribute string, column int32) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_column := (C.gint)(column)

	C.gtk_cell_area_attribute_connect((*C.GtkCellArea)(recv.native), c_renderer, c_attribute, c_column)

	return
}

// AttributeDisconnect is a wrapper around the C function gtk_cell_area_attribute_disconnect.
func (recv *CellArea) AttributeDisconnect(renderer *CellRenderer, attribute string) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	C.gtk_cell_area_attribute_disconnect((*C.GtkCellArea)(recv.native), c_renderer, c_attribute)

	return
}

// Unsupported : gtk_cell_area_cell_get : unsupported parameter ... : varargs

// CellGetProperty is a wrapper around the C function gtk_cell_area_cell_get_property.
func (recv *CellArea) CellGetProperty(renderer *CellRenderer, propertyName string, value *gobject.Value) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(value.ToC())

	C.gtk_cell_area_cell_get_property((*C.GtkCellArea)(recv.native), c_renderer, c_property_name, c_value)

	return
}

// Unsupported : gtk_cell_area_cell_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_cell_area_cell_set : unsupported parameter ... : varargs

// CellSetProperty is a wrapper around the C function gtk_cell_area_cell_set_property.
func (recv *CellArea) CellSetProperty(renderer *CellRenderer, propertyName string, value *gobject.Value) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(value.ToC())

	C.gtk_cell_area_cell_set_property((*C.GtkCellArea)(recv.native), c_renderer, c_property_name, c_value)

	return
}

// Unsupported : gtk_cell_area_cell_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// CopyContext is a wrapper around the C function gtk_cell_area_copy_context.
func (recv *CellArea) CopyContext(context *CellAreaContext) *CellAreaContext {
	c_context := (*C.GtkCellAreaContext)(context.ToC())

	retC := C.gtk_cell_area_copy_context((*C.GtkCellArea)(recv.native), c_context)
	retGo := CellAreaContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateContext is a wrapper around the C function gtk_cell_area_create_context.
func (recv *CellArea) CreateContext() *CellAreaContext {
	retC := C.gtk_cell_area_create_context((*C.GtkCellArea)(recv.native))
	retGo := CellAreaContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_cell_area_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Focus is a wrapper around the C function gtk_cell_area_focus.
func (recv *CellArea) Focus(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_cell_area_focus((*C.GtkCellArea)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_area_foreach : unsupported parameter callback : no type generator for CellCallback, GtkCellCallback

// Unsupported : gtk_cell_area_foreach_alloc : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_allocation : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_get_cell_at_position : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// GetCurrentPathString is a wrapper around the C function gtk_cell_area_get_current_path_string.
func (recv *CellArea) GetCurrentPathString() string {
	retC := C.gtk_cell_area_get_current_path_string((*C.GtkCellArea)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_cell_area_get_edit_widget : no return generator

// GetEditedCell is a wrapper around the C function gtk_cell_area_get_edited_cell.
func (recv *CellArea) GetEditedCell() *CellRenderer {
	retC := C.gtk_cell_area_get_edited_cell((*C.GtkCellArea)(recv.native))
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusCell is a wrapper around the C function gtk_cell_area_get_focus_cell.
func (recv *CellArea) GetFocusCell() *CellRenderer {
	retC := C.gtk_cell_area_get_focus_cell((*C.GtkCellArea)(recv.native))
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusFromSibling is a wrapper around the C function gtk_cell_area_get_focus_from_sibling.
func (recv *CellArea) GetFocusFromSibling(renderer *CellRenderer) *CellRenderer {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	retC := C.gtk_cell_area_get_focus_from_sibling((*C.GtkCellArea)(recv.native), c_renderer)
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusSiblings is a wrapper around the C function gtk_cell_area_get_focus_siblings.
func (recv *CellArea) GetFocusSiblings(renderer *CellRenderer) *glib.List {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	retC := C.gtk_cell_area_get_focus_siblings((*C.GtkCellArea)(recv.native), c_renderer)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_cell_area_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// GetRequestMode is a wrapper around the C function gtk_cell_area_get_request_mode.
func (recv *CellArea) GetRequestMode() SizeRequestMode {
	retC := C.gtk_cell_area_get_request_mode((*C.GtkCellArea)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// HasRenderer is a wrapper around the C function gtk_cell_area_has_renderer.
func (recv *CellArea) HasRenderer(renderer *CellRenderer) bool {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	retC := C.gtk_cell_area_has_renderer((*C.GtkCellArea)(recv.native), c_renderer)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_area_inner_cell_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// IsActivatable is a wrapper around the C function gtk_cell_area_is_activatable.
func (recv *CellArea) IsActivatable() bool {
	retC := C.gtk_cell_area_is_activatable((*C.GtkCellArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsFocusSibling is a wrapper around the C function gtk_cell_area_is_focus_sibling.
func (recv *CellArea) IsFocusSibling(renderer *CellRenderer, sibling *CellRenderer) bool {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_sibling := (*C.GtkCellRenderer)(sibling.ToC())

	retC := C.gtk_cell_area_is_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)
	retGo := retC == C.TRUE

	return retGo
}

// Remove is a wrapper around the C function gtk_cell_area_remove.
func (recv *CellArea) Remove(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	C.gtk_cell_area_remove((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// RemoveFocusSibling is a wrapper around the C function gtk_cell_area_remove_focus_sibling.
func (recv *CellArea) RemoveFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_sibling := (*C.GtkCellRenderer)(sibling.ToC())

	C.gtk_cell_area_remove_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)

	return
}

// Unsupported : gtk_cell_area_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_area_request_renderer : unsupported parameter minimum_size : no type generator for gint, gint*

// SetFocusCell is a wrapper around the C function gtk_cell_area_set_focus_cell.
func (recv *CellArea) SetFocusCell(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	C.gtk_cell_area_set_focus_cell((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// StopEditing is a wrapper around the C function gtk_cell_area_stop_editing.
func (recv *CellArea) StopEditing(canceled bool) {
	c_canceled :=
		boolToGboolean(canceled)

	C.gtk_cell_area_stop_editing((*C.GtkCellArea)(recv.native), c_canceled)

	return
}

// CellAreaBoxNew is a wrapper around the C function gtk_cell_area_box_new.
func CellAreaBoxNew() *CellArea {
	retC := C.gtk_cell_area_box_new()
	retGo := CellAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_cell_area_box_get_spacing.
func (recv *CellAreaBox) GetSpacing() int32 {
	retC := C.gtk_cell_area_box_get_spacing((*C.GtkCellAreaBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// PackEnd is a wrapper around the C function gtk_cell_area_box_pack_end.
func (recv *CellAreaBox) PackEnd(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_expand :=
		boolToGboolean(expand)

	c_align :=
		boolToGboolean(align)

	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_cell_area_box_pack_end((*C.GtkCellAreaBox)(recv.native), c_renderer, c_expand, c_align, c_fixed)

	return
}

// PackStart is a wrapper around the C function gtk_cell_area_box_pack_start.
func (recv *CellAreaBox) PackStart(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_expand :=
		boolToGboolean(expand)

	c_align :=
		boolToGboolean(align)

	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_cell_area_box_pack_start((*C.GtkCellAreaBox)(recv.native), c_renderer, c_expand, c_align, c_fixed)

	return
}

// SetSpacing is a wrapper around the C function gtk_cell_area_box_set_spacing.
func (recv *CellAreaBox) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_cell_area_box_set_spacing((*C.GtkCellAreaBox)(recv.native), c_spacing)

	return
}

// Unsupported : gtk_cell_area_context_get_allocation : unsupported parameter width : no type generator for gint, gint*

// GetArea is a wrapper around the C function gtk_cell_area_context_get_area.
func (recv *CellAreaContext) GetArea() *CellArea {
	retC := C.gtk_cell_area_context_get_area((*C.GtkCellAreaContext)(recv.native))
	retGo := CellAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_cell_area_context_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_cell_area_context_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// PushPreferredHeight is a wrapper around the C function gtk_cell_area_context_push_preferred_height.
func (recv *CellAreaContext) PushPreferredHeight(minimumHeight int32, naturalHeight int32) {
	c_minimum_height := (C.gint)(minimumHeight)

	c_natural_height := (C.gint)(naturalHeight)

	C.gtk_cell_area_context_push_preferred_height((*C.GtkCellAreaContext)(recv.native), c_minimum_height, c_natural_height)

	return
}

// PushPreferredWidth is a wrapper around the C function gtk_cell_area_context_push_preferred_width.
func (recv *CellAreaContext) PushPreferredWidth(minimumWidth int32, naturalWidth int32) {
	c_minimum_width := (C.gint)(minimumWidth)

	c_natural_width := (C.gint)(naturalWidth)

	C.gtk_cell_area_context_push_preferred_width((*C.GtkCellAreaContext)(recv.native), c_minimum_width, c_natural_width)

	return
}

// Unsupported : gtk_cell_renderer_activate : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_cell_renderer_get_aligned_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_cell_renderer_get_fixed_size : unsupported parameter width : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// GetPreferredSize is a wrapper around the C function gtk_cell_renderer_get_preferred_size.
func (recv *CellRenderer) GetPreferredSize(widget *Widget) (*Requisition, *Requisition) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	var c_minimum_size C.GtkRequisition

	var c_natural_size C.GtkRequisition

	C.gtk_cell_renderer_get_preferred_size((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := RequisitionNewFromC(unsafe.Pointer(&c_minimum_size))

	naturalSize := RequisitionNewFromC(unsafe.Pointer(&c_natural_size))

	return minimumSize, naturalSize
}

// Unsupported : gtk_cell_renderer_get_preferred_width : unsupported parameter minimum_size : no type generator for gint, gint*

// Unsupported : gtk_cell_renderer_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// GetRequestMode is a wrapper around the C function gtk_cell_renderer_get_request_mode.
func (recv *CellRenderer) GetRequestMode() SizeRequestMode {
	retC := C.gtk_cell_renderer_get_request_mode((*C.GtkCellRenderer)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// Unsupported : gtk_cell_renderer_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// GetState is a wrapper around the C function gtk_cell_renderer_get_state.
func (recv *CellRenderer) GetState(widget *Widget, cellState CellRendererState) StateFlags {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_cell_state := (C.GtkCellRendererState)(cellState)

	retC := C.gtk_cell_renderer_get_state((*C.GtkCellRenderer)(recv.native), c_widget, c_cell_state)
	retGo := (StateFlags)(retC)

	return retGo
}

// IsActivatable is a wrapper around the C function gtk_cell_renderer_is_activatable.
func (recv *CellRenderer) IsActivatable() bool {
	retC := C.gtk_cell_renderer_is_activatable((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_renderer_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_start_editing : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetDrawSensitive is a wrapper around the C function gtk_cell_view_get_draw_sensitive.
func (recv *CellView) GetDrawSensitive() bool {
	retC := C.gtk_cell_view_get_draw_sensitive((*C.GtkCellView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFitModel is a wrapper around the C function gtk_cell_view_get_fit_model.
func (recv *CellView) GetFitModel() bool {
	retC := C.gtk_cell_view_get_fit_model((*C.GtkCellView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_view_get_model : no return generator

// SetBackgroundRgba is a wrapper around the C function gtk_cell_view_set_background_rgba.
func (recv *CellView) SetBackgroundRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(rgba.ToC())

	C.gtk_cell_view_set_background_rgba((*C.GtkCellView)(recv.native), c_rgba)

	return
}

// SetDrawSensitive is a wrapper around the C function gtk_cell_view_set_draw_sensitive.
func (recv *CellView) SetDrawSensitive(drawSensitive bool) {
	c_draw_sensitive :=
		boolToGboolean(drawSensitive)

	C.gtk_cell_view_set_draw_sensitive((*C.GtkCellView)(recv.native), c_draw_sensitive)

	return
}

// SetFitModel is a wrapper around the C function gtk_cell_view_set_fit_model.
func (recv *CellView) SetFitModel(fitModel bool) {
	c_fit_model :=
		boolToGboolean(fitModel)

	C.gtk_cell_view_set_fit_model((*C.GtkCellView)(recv.native), c_fit_model)

	return
}

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

// ColorButtonNewWithRgba is a wrapper around the C function gtk_color_button_new_with_rgba.
func ColorButtonNewWithRgba(rgba *gdk.RGBA) *Widget {
	c_rgba := (*C.GdkRGBA)(rgba.ToC())

	retC := C.gtk_color_button_new_with_rgba(c_rgba)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRgba is a wrapper around the C function gtk_color_button_get_rgba.
func (recv *ColorButton) GetRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_button_get_rgba((*C.GtkColorButton)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// SetRgba is a wrapper around the C function gtk_color_button_set_rgba.
func (recv *ColorButton) SetRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(rgba.ToC())

	C.gtk_color_button_set_rgba((*C.GtkColorButton)(recv.native), c_rgba)

	return
}

// GetCurrentRgba is a wrapper around the C function gtk_color_selection_get_current_rgba.
func (recv *ColorSelection) GetCurrentRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_selection_get_current_rgba((*C.GtkColorSelection)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// GetPreviousRgba is a wrapper around the C function gtk_color_selection_get_previous_rgba.
func (recv *ColorSelection) GetPreviousRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_selection_get_previous_rgba((*C.GtkColorSelection)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// SetCurrentRgba is a wrapper around the C function gtk_color_selection_set_current_rgba.
func (recv *ColorSelection) SetCurrentRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(rgba.ToC())

	C.gtk_color_selection_set_current_rgba((*C.GtkColorSelection)(recv.native), c_rgba)

	return
}

// SetPreviousRgba is a wrapper around the C function gtk_color_selection_set_previous_rgba.
func (recv *ColorSelection) SetPreviousRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(rgba.ToC())

	C.gtk_color_selection_set_previous_rgba((*C.GtkColorSelection)(recv.native), c_rgba)

	return
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetActiveId is a wrapper around the C function gtk_combo_box_get_active_id.
func (recv *ComboBox) GetActiveId() string {
	retC := C.gtk_combo_box_get_active_id((*C.GtkComboBox)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIdColumn is a wrapper around the C function gtk_combo_box_get_id_column.
func (recv *ComboBox) GetIdColumn() int32 {
	retC := C.gtk_combo_box_get_id_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_combo_box_get_model : no return generator

// GetPopupFixedWidth is a wrapper around the C function gtk_combo_box_get_popup_fixed_width.
func (recv *ComboBox) GetPopupFixedWidth() bool {
	retC := C.gtk_combo_box_get_popup_fixed_width((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_combo_box_get_row_separator_func : no return generator

// PopupForDevice is a wrapper around the C function gtk_combo_box_popup_for_device.
func (recv *ComboBox) PopupForDevice(device *gdk.Device) {
	c_device := (*C.GdkDevice)(device.ToC())

	C.gtk_combo_box_popup_for_device((*C.GtkComboBox)(recv.native), c_device)

	return
}

// SetActiveId is a wrapper around the C function gtk_combo_box_set_active_id.
func (recv *ComboBox) SetActiveId(activeId string) bool {
	c_active_id := C.CString(activeId)
	defer C.free(unsafe.Pointer(c_active_id))

	retC := C.gtk_combo_box_set_active_id((*C.GtkComboBox)(recv.native), c_active_id)
	retGo := retC == C.TRUE

	return retGo
}

// SetIdColumn is a wrapper around the C function gtk_combo_box_set_id_column.
func (recv *ComboBox) SetIdColumn(idColumn int32) {
	c_id_column := (C.gint)(idColumn)

	C.gtk_combo_box_set_id_column((*C.GtkComboBox)(recv.native), c_id_column)

	return
}

// Unsupported : gtk_combo_box_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// SetPopupFixedWidth is a wrapper around the C function gtk_combo_box_set_popup_fixed_width.
func (recv *ComboBox) SetPopupFixedWidth(fixed bool) {
	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_combo_box_set_popup_fixed_width((*C.GtkComboBox)(recv.native), c_fixed)

	return
}

// Unsupported : gtk_combo_box_set_row_separator_func : unsupported parameter func : no type generator for TreeViewRowSeparatorFunc, GtkTreeViewRowSeparatorFunc

// Insert is a wrapper around the C function gtk_combo_box_text_insert.
func (recv *ComboBoxText) Insert(position int32, id string, text string) {
	c_position := (C.gint)(position)

	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_insert((*C.GtkComboBoxText)(recv.native), c_position, c_id, c_text)

	return
}

// RemoveAll is a wrapper around the C function gtk_combo_box_text_remove_all.
func (recv *ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all((*C.GtkComboBoxText)(recv.native))

	return
}

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

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_set_icon_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// EntryCompletionNewWithArea is a wrapper around the C function gtk_entry_completion_new_with_area.
func EntryCompletionNewWithArea(area *CellArea) *EntryCompletion {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_entry_completion_new_with_area(c_area)
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
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

// LoadSymbolic is a wrapper around the C function gtk_icon_info_load_symbolic.
func (recv *IconInfo) LoadSymbolic(fg *gdk.RGBA, successColor *gdk.RGBA, warningColor *gdk.RGBA, errorColor *gdk.RGBA) (*gdkpixbuf.Pixbuf, bool, error) {
	c_fg := (*C.GdkRGBA)(fg.ToC())

	c_success_color := (*C.GdkRGBA)(successColor.ToC())

	c_warning_color := (*C.GdkRGBA)(warningColor.ToC())

	c_error_color := (*C.GdkRGBA)(errorColor.ToC())

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic((*C.GtkIconInfo)(recv.native), c_fg, c_success_color, c_warning_color, c_error_color, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Unsupported : gtk_icon_info_load_symbolic_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// LoadSymbolicForContext is a wrapper around the C function gtk_icon_info_load_symbolic_for_context.
func (recv *IconInfo) LoadSymbolicForContext(context *StyleContext) (*gdkpixbuf.Pixbuf, bool, error) {
	c_context := (*C.GtkStyleContext)(context.ToC())

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_context((*C.GtkIconInfo)(recv.native), c_context, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Unsupported : gtk_icon_info_load_symbolic_for_context_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gtk_icon_info_load_symbolic_for_context_finish : unsupported parameter res : no type generator for Gio.AsyncResult, GAsyncResult*

// LoadSymbolicForStyle is a wrapper around the C function gtk_icon_info_load_symbolic_for_style.
func (recv *IconInfo) LoadSymbolicForStyle(style *Style, state StateType) (*gdkpixbuf.Pixbuf, bool, error) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_state := (C.GtkStateType)(state)

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_style((*C.GtkIconInfo)(recv.native), c_style, c_state, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goThrowableError
}

// Unsupported : gtk_icon_theme_choose_icon : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_choose_icon_for_scale : unsupported parameter icon_names : no param type

// Unsupported : gtk_icon_theme_get_icon_sizes : no return type

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : no param type

// Unsupported : gtk_icon_theme_lookup_by_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_lookup_by_gicon_for_scale : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_icon_theme_set_search_path : unsupported parameter path : no param type

// IconViewNewWithArea is a wrapper around the C function gtk_icon_view_new_with_area.
func IconViewNewWithArea(area *CellArea) *Widget {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_icon_view_new_with_area(c_area)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported : gtk_level_bar_get_offset_value : unsupported parameter value : no type generator for gdouble, gdouble*

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

// GetReserveIndicator is a wrapper around the C function gtk_menu_item_get_reserve_indicator.
func (recv *MenuItem) GetReserveIndicator() bool {
	retC := C.gtk_menu_item_get_reserve_indicator((*C.GtkMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetReserveIndicator is a wrapper around the C function gtk_menu_item_set_reserve_indicator.
func (recv *MenuItem) SetReserveIndicator(reserve bool) {
	c_reserve :=
		boolToGboolean(reserve)

	C.gtk_menu_item_set_reserve_indicator((*C.GtkMenuItem)(recv.native), c_reserve)

	return
}

// Unsupported : gtk_menu_item_toggle_size_request : unsupported parameter requisition : no type generator for gint, gint*

// GetParentShell is a wrapper around the C function gtk_menu_shell_get_parent_shell.
func (recv *MenuShell) GetParentShell() *Widget {
	retC := C.gtk_menu_shell_get_parent_shell((*C.GtkMenuShell)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectedItem is a wrapper around the C function gtk_menu_shell_get_selected_item.
func (recv *MenuShell) GetSelectedItem() *Widget {
	retC := C.gtk_menu_shell_get_selected_item((*C.GtkMenuShell)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_markup : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_format_secondary_text : unsupported parameter ... : varargs

// Unsupported : gtk_misc_get_alignment : unsupported parameter xalign : no type generator for gfloat, gfloat*

// Unsupported : gtk_misc_get_padding : unsupported parameter xpad : no type generator for gint, gint*

// Unsupported : gtk_numerable_icon_get_background_gicon : no return generator

// GetBackgroundIconName is a wrapper around the C function gtk_numerable_icon_get_background_icon_name.
func (recv *NumerableIcon) GetBackgroundIconName() string {
	retC := C.gtk_numerable_icon_get_background_icon_name((*C.GtkNumerableIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetCount is a wrapper around the C function gtk_numerable_icon_get_count.
func (recv *NumerableIcon) GetCount() int32 {
	retC := C.gtk_numerable_icon_get_count((*C.GtkNumerableIcon)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLabel is a wrapper around the C function gtk_numerable_icon_get_label.
func (recv *NumerableIcon) GetLabel() string {
	retC := C.gtk_numerable_icon_get_label((*C.GtkNumerableIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStyleContext is a wrapper around the C function gtk_numerable_icon_get_style_context.
func (recv *NumerableIcon) GetStyleContext() *StyleContext {
	retC := C.gtk_numerable_icon_get_style_context((*C.GtkNumerableIcon)(recv.native))
	retGo := StyleContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_numerable_icon_set_background_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// SetBackgroundIconName is a wrapper around the C function gtk_numerable_icon_set_background_icon_name.
func (recv *NumerableIcon) SetBackgroundIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_numerable_icon_set_background_icon_name((*C.GtkNumerableIcon)(recv.native), c_icon_name)

	return
}

// SetCount is a wrapper around the C function gtk_numerable_icon_set_count.
func (recv *NumerableIcon) SetCount(count int32) {
	c_count := (C.gint)(count)

	C.gtk_numerable_icon_set_count((*C.GtkNumerableIcon)(recv.native), c_count)

	return
}

// SetLabel is a wrapper around the C function gtk_numerable_icon_set_label.
func (recv *NumerableIcon) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_numerable_icon_set_label((*C.GtkNumerableIcon)(recv.native), c_label)

	return
}

// SetStyleContext is a wrapper around the C function gtk_numerable_icon_set_style_context.
func (recv *NumerableIcon) SetStyleContext(style *StyleContext) {
	c_style := (*C.GtkStyleContext)(style.ToC())

	C.gtk_numerable_icon_set_style_context((*C.GtkNumerableIcon)(recv.native), c_style)

	return
}

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_pad_controller_set_action_entries : unsupported parameter entries : no param type

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_page_setup_to_gvariant : return type : Blacklisted record : GVariant

// PanedNew is a wrapper around the C function gtk_paned_new.
func PanedNew(orientation Orientation) *Widget {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_paned_new(c_orientation)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_places_sidebar_add_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_get_location : no return generator

// Unsupported : gtk_places_sidebar_get_nth_bookmark : no return generator

// Unsupported : gtk_places_sidebar_remove_shortcut : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_places_sidebar_set_location : unsupported parameter location : no type generator for Gio.File, GFile*

// Unsupported : gtk_popover_get_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_popover_set_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_print_context_get_hard_margins : unsupported parameter top : no type generator for gdouble, gdouble*

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_foreach : unsupported parameter func : no type generator for PrintSettingsFunc, GtkPrintSettingsFunc

// Unsupported : gtk_print_settings_get_page_ranges : unsupported parameter num_ranges : no type generator for gint, gint*

// Unsupported : gtk_print_settings_set_page_ranges : unsupported parameter page_ranges : no param type

// Unsupported : gtk_print_settings_to_gvariant : return type : Blacklisted record : GVariant

// GetShowText is a wrapper around the C function gtk_progress_bar_get_show_text.
func (recv *ProgressBar) GetShowText() bool {
	retC := C.gtk_progress_bar_get_show_text((*C.GtkProgressBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowText is a wrapper around the C function gtk_progress_bar_set_show_text.
func (recv *ProgressBar) SetShowText(showText bool) {
	c_show_text :=
		boolToGboolean(showText)

	C.gtk_progress_bar_set_show_text((*C.GtkProgressBar)(recv.native), c_show_text)

	return
}

// JoinGroup is a wrapper around the C function gtk_radio_action_join_group.
func (recv *RadioAction) JoinGroup(groupSource *RadioAction) {
	c_group_source := (*C.GtkRadioAction)(groupSource.ToC())

	C.gtk_radio_action_join_group((*C.GtkRadioAction)(recv.native), c_group_source)

	return
}

// JoinGroup is a wrapper around the C function gtk_radio_button_join_group.
func (recv *RadioButton) JoinGroup(groupSource *RadioButton) {
	c_group_source := (*C.GtkRadioButton)(groupSource.ToC())

	C.gtk_radio_button_join_group((*C.GtkRadioButton)(recv.native), c_group_source)

	return
}

// Unsupported : gtk_range_get_range_rect : unsupported parameter range_rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_range_get_slider_range : unsupported parameter slider_start : no type generator for gint, gint*

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_recent_filter_add_custom : unsupported parameter func : no type generator for RecentFilterFunc, GtkRecentFilterFunc

// ScaleNew is a wrapper around the C function gtk_scale_new.
func ScaleNew(orientation Orientation, adjustment *Adjustment) *Widget {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_scale_new(c_orientation, c_adjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScaleNewWithRange is a wrapper around the C function gtk_scale_new_with_range.
func ScaleNewWithRange(orientation Orientation, min float64, max float64, step float64) *Widget {
	c_orientation := (C.GtkOrientation)(orientation)

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_scale_new_with_range(c_orientation, c_min, c_max, c_step)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_scale_get_layout_offsets : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons : no param type

// ScrollbarNew is a wrapper around the C function gtk_scrollbar_new.
func ScrollbarNew(orientation Orientation, adjustment *Adjustment) *Widget {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_scrollbar_new(c_orientation, c_adjustment)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMinContentHeight is a wrapper around the C function gtk_scrolled_window_get_min_content_height.
func (recv *ScrolledWindow) GetMinContentHeight() int32 {
	retC := C.gtk_scrolled_window_get_min_content_height((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinContentWidth is a wrapper around the C function gtk_scrolled_window_get_min_content_width.
func (recv *ScrolledWindow) GetMinContentWidth() int32 {
	retC := C.gtk_scrolled_window_get_min_content_width((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_scrolled_window_get_policy : unsupported parameter hscrollbar_policy : GtkPolicyType* with indirection level of 1

// SetMinContentHeight is a wrapper around the C function gtk_scrolled_window_set_min_content_height.
func (recv *ScrolledWindow) SetMinContentHeight(height int32) {
	c_height := (C.gint)(height)

	C.gtk_scrolled_window_set_min_content_height((*C.GtkScrolledWindow)(recv.native), c_height)

	return
}

// SetMinContentWidth is a wrapper around the C function gtk_scrolled_window_set_min_content_width.
func (recv *ScrolledWindow) SetMinContentWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_scrolled_window_set_min_content_width((*C.GtkScrolledWindow)(recv.native), c_width)

	return
}

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_search_entry_handle_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// SeparatorNew is a wrapper around the C function gtk_separator_new.
func SeparatorNew(orientation Orientation) *Widget {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_separator_new(c_orientation)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_spin_button_get_increments : unsupported parameter step : no type generator for gdouble, gdouble*

// Unsupported : gtk_spin_button_get_range : unsupported parameter min : no type generator for gdouble, gdouble*

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_status_icon_get_geometry : unsupported parameter screen : record with indirection level of 2

// Unsupported : gtk_status_icon_get_gicon : no return generator

// Unsupported : gtk_status_icon_set_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_style_get : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_style_property : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_style_get_valist : unsupported parameter widget_type : no type generator for GType, GType

// HasContext is a wrapper around the C function gtk_style_has_context.
func (recv *Style) HasContext() bool {
	retC := C.gtk_style_has_context((*C.GtkStyle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_style_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// AddClass is a wrapper around the C function gtk_style_context_add_class.
func (recv *StyleContext) AddClass(className string) {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	C.gtk_style_context_add_class((*C.GtkStyleContext)(recv.native), c_class_name)

	return
}

// Unsupported : gtk_style_context_add_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// AddRegion is a wrapper around the C function gtk_style_context_add_region.
func (recv *StyleContext) AddRegion(regionName string, flags RegionFlags) {
	c_region_name := C.CString(regionName)
	defer C.free(unsafe.Pointer(c_region_name))

	c_flags := (C.GtkRegionFlags)(flags)

	C.gtk_style_context_add_region((*C.GtkStyleContext)(recv.native), c_region_name, c_flags)

	return
}

// CancelAnimations is a wrapper around the C function gtk_style_context_cancel_animations.
func (recv *StyleContext) CancelAnimations(regionId uintptr) {
	c_region_id := (C.gpointer)(regionId)

	C.gtk_style_context_cancel_animations((*C.GtkStyleContext)(recv.native), c_region_id)

	return
}

// Unsupported : gtk_style_context_get : unsupported parameter ... : varargs

// GetBackgroundColor is a wrapper around the C function gtk_style_context_get_background_color.
func (recv *StyleContext) GetBackgroundColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_background_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetBorder is a wrapper around the C function gtk_style_context_get_border.
func (recv *StyleContext) GetBorder(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_border C.GtkBorder

	C.gtk_style_context_get_border((*C.GtkStyleContext)(recv.native), c_state, &c_border)

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return border
}

// GetBorderColor is a wrapper around the C function gtk_style_context_get_border_color.
func (recv *StyleContext) GetBorderColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_border_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetColor is a wrapper around the C function gtk_style_context_get_color.
func (recv *StyleContext) GetColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetDirection is a wrapper around the C function gtk_style_context_get_direction.
func (recv *StyleContext) GetDirection() TextDirection {
	retC := C.gtk_style_context_get_direction((*C.GtkStyleContext)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// GetFont is a wrapper around the C function gtk_style_context_get_font.
func (recv *StyleContext) GetFont(state StateFlags) *pango.FontDescription {
	c_state := (C.GtkStateFlags)(state)

	retC := C.gtk_style_context_get_font((*C.GtkStyleContext)(recv.native), c_state)
	retGo := pango.FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetJunctionSides is a wrapper around the C function gtk_style_context_get_junction_sides.
func (recv *StyleContext) GetJunctionSides() JunctionSides {
	retC := C.gtk_style_context_get_junction_sides((*C.GtkStyleContext)(recv.native))
	retGo := (JunctionSides)(retC)

	return retGo
}

// GetMargin is a wrapper around the C function gtk_style_context_get_margin.
func (recv *StyleContext) GetMargin(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_margin C.GtkBorder

	C.gtk_style_context_get_margin((*C.GtkStyleContext)(recv.native), c_state, &c_margin)

	margin := BorderNewFromC(unsafe.Pointer(&c_margin))

	return margin
}

// GetPadding is a wrapper around the C function gtk_style_context_get_padding.
func (recv *StyleContext) GetPadding(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_padding C.GtkBorder

	C.gtk_style_context_get_padding((*C.GtkStyleContext)(recv.native), c_state, &c_padding)

	padding := BorderNewFromC(unsafe.Pointer(&c_padding))

	return padding
}

// GetPath is a wrapper around the C function gtk_style_context_get_path.
func (recv *StyleContext) GetPath() *WidgetPath {
	retC := C.gtk_style_context_get_path((*C.GtkStyleContext)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProperty is a wrapper around the C function gtk_style_context_get_property.
func (recv *StyleContext) GetProperty(property string, state StateFlags) *gobject.Value {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	C.gtk_style_context_get_property((*C.GtkStyleContext)(recv.native), c_property, c_state, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// GetState is a wrapper around the C function gtk_style_context_get_state.
func (recv *StyleContext) GetState() StateFlags {
	retC := C.gtk_style_context_get_state((*C.GtkStyleContext)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list, va_list

// HasClass is a wrapper around the C function gtk_style_context_has_class.
func (recv *StyleContext) HasClass(className string) bool {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	retC := C.gtk_style_context_has_class((*C.GtkStyleContext)(recv.native), c_class_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Invalidate is a wrapper around the C function gtk_style_context_invalidate.
func (recv *StyleContext) Invalidate() {
	C.gtk_style_context_invalidate((*C.GtkStyleContext)(recv.native))

	return
}

// ListClasses is a wrapper around the C function gtk_style_context_list_classes.
func (recv *StyleContext) ListClasses() *glib.List {
	retC := C.gtk_style_context_list_classes((*C.GtkStyleContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListRegions is a wrapper around the C function gtk_style_context_list_regions.
func (recv *StyleContext) ListRegions() *glib.List {
	retC := C.gtk_style_context_list_regions((*C.GtkStyleContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NotifyStateChange is a wrapper around the C function gtk_style_context_notify_state_change.
func (recv *StyleContext) NotifyStateChange(window *gdk.Window, regionId uintptr, state StateType, stateValue bool) {
	c_window := (*C.GdkWindow)(window.ToC())

	c_region_id := (C.gpointer)(regionId)

	c_state := (C.GtkStateType)(state)

	c_state_value :=
		boolToGboolean(stateValue)

	C.gtk_style_context_notify_state_change((*C.GtkStyleContext)(recv.native), c_window, c_region_id, c_state, c_state_value)

	return
}

// PopAnimatableRegion is a wrapper around the C function gtk_style_context_pop_animatable_region.
func (recv *StyleContext) PopAnimatableRegion() {
	C.gtk_style_context_pop_animatable_region((*C.GtkStyleContext)(recv.native))

	return
}

// PushAnimatableRegion is a wrapper around the C function gtk_style_context_push_animatable_region.
func (recv *StyleContext) PushAnimatableRegion(regionId uintptr) {
	c_region_id := (C.gpointer)(regionId)

	C.gtk_style_context_push_animatable_region((*C.GtkStyleContext)(recv.native), c_region_id)

	return
}

// RemoveClass is a wrapper around the C function gtk_style_context_remove_class.
func (recv *StyleContext) RemoveClass(className string) {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	C.gtk_style_context_remove_class((*C.GtkStyleContext)(recv.native), c_class_name)

	return
}

// Unsupported : gtk_style_context_remove_provider : unsupported parameter provider : no type generator for StyleProvider, GtkStyleProvider*

// RemoveRegion is a wrapper around the C function gtk_style_context_remove_region.
func (recv *StyleContext) RemoveRegion(regionName string) {
	c_region_name := C.CString(regionName)
	defer C.free(unsafe.Pointer(c_region_name))

	C.gtk_style_context_remove_region((*C.GtkStyleContext)(recv.native), c_region_name)

	return
}

// Restore is a wrapper around the C function gtk_style_context_restore.
func (recv *StyleContext) Restore() {
	C.gtk_style_context_restore((*C.GtkStyleContext)(recv.native))

	return
}

// Save is a wrapper around the C function gtk_style_context_save.
func (recv *StyleContext) Save() {
	C.gtk_style_context_save((*C.GtkStyleContext)(recv.native))

	return
}

// ScrollAnimations is a wrapper around the C function gtk_style_context_scroll_animations.
func (recv *StyleContext) ScrollAnimations(window *gdk.Window, dx int32, dy int32) {
	c_window := (*C.GdkWindow)(window.ToC())

	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gtk_style_context_scroll_animations((*C.GtkStyleContext)(recv.native), c_window, c_dx, c_dy)

	return
}

// SetBackground is a wrapper around the C function gtk_style_context_set_background.
func (recv *StyleContext) SetBackground(window *gdk.Window) {
	c_window := (*C.GdkWindow)(window.ToC())

	C.gtk_style_context_set_background((*C.GtkStyleContext)(recv.native), c_window)

	return
}

// SetDirection is a wrapper around the C function gtk_style_context_set_direction.
func (recv *StyleContext) SetDirection(direction TextDirection) {
	c_direction := (C.GtkTextDirection)(direction)

	C.gtk_style_context_set_direction((*C.GtkStyleContext)(recv.native), c_direction)

	return
}

// SetJunctionSides is a wrapper around the C function gtk_style_context_set_junction_sides.
func (recv *StyleContext) SetJunctionSides(sides JunctionSides) {
	c_sides := (C.GtkJunctionSides)(sides)

	C.gtk_style_context_set_junction_sides((*C.GtkStyleContext)(recv.native), c_sides)

	return
}

// SetPath is a wrapper around the C function gtk_style_context_set_path.
func (recv *StyleContext) SetPath(path *WidgetPath) {
	c_path := (*C.GtkWidgetPath)(path.ToC())

	C.gtk_style_context_set_path((*C.GtkStyleContext)(recv.native), c_path)

	return
}

// SetScreen is a wrapper around the C function gtk_style_context_set_screen.
func (recv *StyleContext) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_style_context_set_screen((*C.GtkStyleContext)(recv.native), c_screen)

	return
}

// SetState is a wrapper around the C function gtk_style_context_set_state.
func (recv *StyleContext) SetState(flags StateFlags) {
	c_flags := (C.GtkStateFlags)(flags)

	C.gtk_style_context_set_state((*C.GtkStyleContext)(recv.native), c_flags)

	return
}

// Unsupported : gtk_style_context_state_is_running : unsupported parameter progress : no type generator for gdouble, gdouble*

// Unsupported : gtk_style_properties_get : unsupported parameter ... : varargs

// GetProperty is a wrapper around the C function gtk_style_properties_get_property.
func (recv *StyleProperties) GetProperty(property string, state StateFlags) (bool, *gobject.Value) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	retC := C.gtk_style_properties_get_property((*C.GtkStyleProperties)(recv.native), c_property, c_state, &c_value)
	retGo := retC == C.TRUE

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value
}

// Unsupported : gtk_style_properties_get_valist : unsupported parameter args : no type generator for va_list, va_list

// LookupColor is a wrapper around the C function gtk_style_properties_lookup_color.
func (recv *StyleProperties) LookupColor(name string) *SymbolicColor {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_style_properties_lookup_color((*C.GtkStyleProperties)(recv.native), c_name)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MapColor is a wrapper around the C function gtk_style_properties_map_color.
func (recv *StyleProperties) MapColor(name string, color *SymbolicColor) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_color := (*C.GtkSymbolicColor)(color.ToC())

	C.gtk_style_properties_map_color((*C.GtkStyleProperties)(recv.native), c_name, c_color)

	return
}

// Merge is a wrapper around the C function gtk_style_properties_merge.
func (recv *StyleProperties) Merge(propsToMerge *StyleProperties, replace bool) {
	c_props_to_merge := (*C.GtkStyleProperties)(propsToMerge.ToC())

	c_replace :=
		boolToGboolean(replace)

	C.gtk_style_properties_merge((*C.GtkStyleProperties)(recv.native), c_props_to_merge, c_replace)

	return
}

// Unsupported : gtk_style_properties_set : unsupported parameter ... : varargs

// SetProperty is a wrapper around the C function gtk_style_properties_set_property.
func (recv *StyleProperties) SetProperty(property string, state StateFlags, value *gobject.Value) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	c_value := (*C.GValue)(value.ToC())

	C.gtk_style_properties_set_property((*C.GtkStyleProperties)(recv.native), c_property, c_state, c_value)

	return
}

// Unsupported : gtk_style_properties_set_valist : unsupported parameter args : no type generator for va_list, va_list

// UnsetProperty is a wrapper around the C function gtk_style_properties_unset_property.
func (recv *StyleProperties) UnsetProperty(property string, state StateFlags) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	C.gtk_style_properties_unset_property((*C.GtkStyleProperties)(recv.native), c_property, c_state)

	return
}

// SwitchNew is a wrapper around the C function gtk_switch_new.
func SwitchNew() *Widget {
	retC := C.gtk_switch_new()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_switch_get_active.
func (recv *Switch) GetActive() bool {
	retC := C.gtk_switch_get_active((*C.GtkSwitch)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_switch_set_active.
func (recv *Switch) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_switch_set_active((*C.GtkSwitch)(recv.native), c_is_active)

	return
}

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

// GetBackgroundColor is a wrapper around the C function gtk_theming_engine_get_background_color.
func (recv *ThemingEngine) GetBackgroundColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_background_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetBorder is a wrapper around the C function gtk_theming_engine_get_border.
func (recv *ThemingEngine) GetBorder(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_border C.GtkBorder

	C.gtk_theming_engine_get_border((*C.GtkThemingEngine)(recv.native), c_state, &c_border)

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return border
}

// GetBorderColor is a wrapper around the C function gtk_theming_engine_get_border_color.
func (recv *ThemingEngine) GetBorderColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_border_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetColor is a wrapper around the C function gtk_theming_engine_get_color.
func (recv *ThemingEngine) GetColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetDirection is a wrapper around the C function gtk_theming_engine_get_direction.
func (recv *ThemingEngine) GetDirection() TextDirection {
	retC := C.gtk_theming_engine_get_direction((*C.GtkThemingEngine)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// GetFont is a wrapper around the C function gtk_theming_engine_get_font.
func (recv *ThemingEngine) GetFont(state StateFlags) *pango.FontDescription {
	c_state := (C.GtkStateFlags)(state)

	retC := C.gtk_theming_engine_get_font((*C.GtkThemingEngine)(recv.native), c_state)
	retGo := pango.FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetJunctionSides is a wrapper around the C function gtk_theming_engine_get_junction_sides.
func (recv *ThemingEngine) GetJunctionSides() JunctionSides {
	retC := C.gtk_theming_engine_get_junction_sides((*C.GtkThemingEngine)(recv.native))
	retGo := (JunctionSides)(retC)

	return retGo
}

// GetMargin is a wrapper around the C function gtk_theming_engine_get_margin.
func (recv *ThemingEngine) GetMargin(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_margin C.GtkBorder

	C.gtk_theming_engine_get_margin((*C.GtkThemingEngine)(recv.native), c_state, &c_margin)

	margin := BorderNewFromC(unsafe.Pointer(&c_margin))

	return margin
}

// GetPadding is a wrapper around the C function gtk_theming_engine_get_padding.
func (recv *ThemingEngine) GetPadding(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_padding C.GtkBorder

	C.gtk_theming_engine_get_padding((*C.GtkThemingEngine)(recv.native), c_state, &c_padding)

	padding := BorderNewFromC(unsafe.Pointer(&c_padding))

	return padding
}

// GetPath is a wrapper around the C function gtk_theming_engine_get_path.
func (recv *ThemingEngine) GetPath() *WidgetPath {
	retC := C.gtk_theming_engine_get_path((*C.GtkThemingEngine)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProperty is a wrapper around the C function gtk_theming_engine_get_property.
func (recv *ThemingEngine) GetProperty(property string, state StateFlags) *gobject.Value {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	C.gtk_theming_engine_get_property((*C.GtkThemingEngine)(recv.native), c_property, c_state, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// GetState is a wrapper around the C function gtk_theming_engine_get_state.
func (recv *ThemingEngine) GetState() StateFlags {
	retC := C.gtk_theming_engine_get_state((*C.GtkThemingEngine)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Unsupported : gtk_theming_engine_get_style : unsupported parameter ... : varargs

// GetStyleProperty is a wrapper around the C function gtk_theming_engine_get_style_property.
func (recv *ThemingEngine) GetStyleProperty(propertyName string) *gobject.Value {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	var c_value C.GValue

	C.gtk_theming_engine_get_style_property((*C.GtkThemingEngine)(recv.native), c_property_name, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Unsupported : gtk_theming_engine_get_style_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : gtk_theming_engine_get_valist : unsupported parameter args : no type generator for va_list, va_list

// HasClass is a wrapper around the C function gtk_theming_engine_has_class.
func (recv *ThemingEngine) HasClass(styleClass string) bool {
	c_style_class := C.CString(styleClass)
	defer C.free(unsafe.Pointer(c_style_class))

	retC := C.gtk_theming_engine_has_class((*C.GtkThemingEngine)(recv.native), c_style_class)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_theming_engine_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// LookupColor is a wrapper around the C function gtk_theming_engine_lookup_color.
func (recv *ThemingEngine) LookupColor(colorName string) (bool, *gdk.RGBA) {
	c_color_name := C.CString(colorName)
	defer C.free(unsafe.Pointer(c_color_name))

	var c_color C.GdkRGBA

	retC := C.gtk_theming_engine_lookup_color((*C.GtkThemingEngine)(recv.native), c_color_name, &c_color)
	retGo := retC == C.TRUE

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

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

// TreeViewColumnNewWithArea is a wrapper around the C function gtk_tree_view_column_new_with_area.
func TreeViewColumnNewWithArea(area *CellArea) *TreeViewColumn {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_tree_view_column_new_with_area(c_area)
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_cell_get_position : unsupported parameter x_offset : no type generator for gint, gint*

// Unsupported : gtk_tree_view_column_cell_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_view_column_cell_set_cell_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// GetButton is a wrapper around the C function gtk_tree_view_column_get_button.
func (recv *TreeViewColumn) GetButton() *Widget {
	retC := C.gtk_tree_view_column_get_button((*C.GtkTreeViewColumn)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_column_set_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_set_cell_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// AddDeviceEvents is a wrapper around the C function gtk_widget_add_device_events.
func (recv *Widget) AddDeviceEvents(device *gdk.Device, events gdk.EventMask) {
	c_device := (*C.GdkDevice)(device.ToC())

	c_events := (C.GdkEventMask)(events)

	C.gtk_widget_add_device_events((*C.GtkWidget)(recv.native), c_device, c_events)

	return
}

// Unsupported : gtk_widget_add_tick_callback : unsupported parameter callback : no type generator for TickCallback, GtkTickCallback

// Unsupported : gtk_widget_class_path : unsupported parameter path_length : no type generator for guint, guint*

// Unsupported : gtk_widget_destroyed : unsupported parameter widget_pointer : record with indirection level of 2

// DeviceIsShadowed is a wrapper around the C function gtk_widget_device_is_shadowed.
func (recv *Widget) DeviceIsShadowed(device *gdk.Device) bool {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gtk_widget_device_is_shadowed((*C.GtkWidget)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_drag_begin : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_drag_dest_find_target : return type : Blacklisted record : GdkAtom

// Unsupported : gtk_drag_dest_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_get_data : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_drag_source_set : unsupported parameter targets : no param type

// Unsupported : gtk_drag_source_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Draw is a wrapper around the C function gtk_widget_draw.
func (recv *Widget) Draw(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(cr.ToC())

	C.gtk_widget_draw((*C.GtkWidget)(recv.native), c_cr)

	return
}

// Unsupported : gtk_widget_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_get_action_group : no return generator

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_ancestor : unsupported parameter widget_type : no type generator for GType, GType

// Unsupported : gtk_widget_get_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_get_clipboard : unsupported parameter selection : Blacklisted record : GdkAtom

// GetDeviceEnabled is a wrapper around the C function gtk_widget_get_device_enabled.
func (recv *Widget) GetDeviceEnabled(device *gdk.Device) bool {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gtk_widget_get_device_enabled((*C.GtkWidget)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// GetDeviceEvents is a wrapper around the C function gtk_widget_get_device_events.
func (recv *Widget) GetDeviceEvents(device *gdk.Device) gdk.EventMask {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gtk_widget_get_device_events((*C.GtkWidget)(recv.native), c_device)
	retGo := (gdk.EventMask)(retC)

	return retGo
}

// GetMarginBottom is a wrapper around the C function gtk_widget_get_margin_bottom.
func (recv *Widget) GetMarginBottom() int32 {
	retC := C.gtk_widget_get_margin_bottom((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginLeft is a wrapper around the C function gtk_widget_get_margin_left.
func (recv *Widget) GetMarginLeft() int32 {
	retC := C.gtk_widget_get_margin_left((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginRight is a wrapper around the C function gtk_widget_get_margin_right.
func (recv *Widget) GetMarginRight() int32 {
	retC := C.gtk_widget_get_margin_right((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginTop is a wrapper around the C function gtk_widget_get_margin_top.
func (recv *Widget) GetMarginTop() int32 {
	retC := C.gtk_widget_get_margin_top((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_pointer : unsupported parameter x : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_and_baseline_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_height_for_width : unsupported parameter minimum_height : no type generator for gint, gint*

// GetPreferredSize is a wrapper around the C function gtk_widget_get_preferred_size.
func (recv *Widget) GetPreferredSize() (*Requisition, *Requisition) {
	var c_minimum_size C.GtkRequisition

	var c_natural_size C.GtkRequisition

	C.gtk_widget_get_preferred_size((*C.GtkWidget)(recv.native), &c_minimum_size, &c_natural_size)

	minimumSize := RequisitionNewFromC(unsafe.Pointer(&c_minimum_size))

	naturalSize := RequisitionNewFromC(unsafe.Pointer(&c_natural_size))

	return minimumSize, naturalSize
}

// Unsupported : gtk_widget_get_preferred_width : unsupported parameter minimum_width : no type generator for gint, gint*

// Unsupported : gtk_widget_get_preferred_width_for_height : unsupported parameter minimum_width : no type generator for gint, gint*

// GetRequestMode is a wrapper around the C function gtk_widget_get_request_mode.
func (recv *Widget) GetRequestMode() SizeRequestMode {
	retC := C.gtk_widget_get_request_mode((*C.GtkWidget)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_size_request : unsupported parameter width : no type generator for gint, gint*

// GetStateFlags is a wrapper around the C function gtk_widget_get_state_flags.
func (recv *Widget) GetStateFlags() StateFlags {
	retC := C.gtk_widget_get_state_flags((*C.GtkWidget)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_template_child : unsupported parameter widget_type : no type generator for GType, GType

// InputShapeCombineRegion is a wrapper around the C function gtk_widget_input_shape_combine_region.
func (recv *Widget) InputShapeCombineRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(region.ToC())

	C.gtk_widget_input_shape_combine_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// Unsupported : gtk_widget_insert_action_group : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_widget_intersect : unsupported parameter area : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_list_action_prefixes : no return type

// OverrideBackgroundColor is a wrapper around the C function gtk_widget_override_background_color.
func (recv *Widget) OverrideBackgroundColor(state StateFlags, color *gdk.RGBA) {
	c_state := (C.GtkStateFlags)(state)

	c_color := (*C.GdkRGBA)(color.ToC())

	C.gtk_widget_override_background_color((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// OverrideColor is a wrapper around the C function gtk_widget_override_color.
func (recv *Widget) OverrideColor(state StateFlags, color *gdk.RGBA) {
	c_state := (C.GtkStateFlags)(state)

	c_color := (*C.GdkRGBA)(color.ToC())

	C.gtk_widget_override_color((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// OverrideCursor is a wrapper around the C function gtk_widget_override_cursor.
func (recv *Widget) OverrideCursor(cursor *gdk.RGBA, secondaryCursor *gdk.RGBA) {
	c_cursor := (*C.GdkRGBA)(cursor.ToC())

	c_secondary_cursor := (*C.GdkRGBA)(secondaryCursor.ToC())

	C.gtk_widget_override_cursor((*C.GtkWidget)(recv.native), c_cursor, c_secondary_cursor)

	return
}

// OverrideFont is a wrapper around the C function gtk_widget_override_font.
func (recv *Widget) OverrideFont(fontDesc *pango.FontDescription) {
	c_font_desc := (*C.PangoFontDescription)(fontDesc.ToC())

	C.gtk_widget_override_font((*C.GtkWidget)(recv.native), c_font_desc)

	return
}

// OverrideSymbolicColor is a wrapper around the C function gtk_widget_override_symbolic_color.
func (recv *Widget) OverrideSymbolicColor(name string, color *gdk.RGBA) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_color := (*C.GdkRGBA)(color.ToC())

	C.gtk_widget_override_symbolic_color((*C.GtkWidget)(recv.native), c_name, c_color)

	return
}

// Unsupported : gtk_widget_path : unsupported parameter path_length : no type generator for guint, guint*

// QueueDrawRegion is a wrapper around the C function gtk_widget_queue_draw_region.
func (recv *Widget) QueueDrawRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(region.ToC())

	C.gtk_widget_queue_draw_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// Unsupported : gtk_widget_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_widget_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// ResetStyle is a wrapper around the C function gtk_widget_reset_style.
func (recv *Widget) ResetStyle() {
	C.gtk_widget_reset_style((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_send_expose : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_send_focus_change : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_widget_set_allocation : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_set_clip : unsupported parameter clip : Blacklisted record : GdkRectangle

// SetDeviceEnabled is a wrapper around the C function gtk_widget_set_device_enabled.
func (recv *Widget) SetDeviceEnabled(device *gdk.Device, enabled bool) {
	c_device := (*C.GdkDevice)(device.ToC())

	c_enabled :=
		boolToGboolean(enabled)

	C.gtk_widget_set_device_enabled((*C.GtkWidget)(recv.native), c_device, c_enabled)

	return
}

// SetDeviceEvents is a wrapper around the C function gtk_widget_set_device_events.
func (recv *Widget) SetDeviceEvents(device *gdk.Device, events gdk.EventMask) {
	c_device := (*C.GdkDevice)(device.ToC())

	c_events := (C.GdkEventMask)(events)

	C.gtk_widget_set_device_events((*C.GtkWidget)(recv.native), c_device, c_events)

	return
}

// SetMarginBottom is a wrapper around the C function gtk_widget_set_margin_bottom.
func (recv *Widget) SetMarginBottom(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_bottom((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginLeft is a wrapper around the C function gtk_widget_set_margin_left.
func (recv *Widget) SetMarginLeft(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_left((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginRight is a wrapper around the C function gtk_widget_set_margin_right.
func (recv *Widget) SetMarginRight(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_right((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginTop is a wrapper around the C function gtk_widget_set_margin_top.
func (recv *Widget) SetMarginTop(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_top((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetStateFlags is a wrapper around the C function gtk_widget_set_state_flags.
func (recv *Widget) SetStateFlags(flags StateFlags, clear bool) {
	c_flags := (C.GtkStateFlags)(flags)

	c_clear :=
		boolToGboolean(clear)

	C.gtk_widget_set_state_flags((*C.GtkWidget)(recv.native), c_flags, c_clear)

	return
}

// SetSupportMultidevice is a wrapper around the C function gtk_widget_set_support_multidevice.
func (recv *Widget) SetSupportMultidevice(supportMultidevice bool) {
	c_support_multidevice :=
		boolToGboolean(supportMultidevice)

	C.gtk_widget_set_support_multidevice((*C.GtkWidget)(recv.native), c_support_multidevice)

	return
}

// ShapeCombineRegion is a wrapper around the C function gtk_widget_shape_combine_region.
func (recv *Widget) ShapeCombineRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(region.ToC())

	C.gtk_widget_shape_combine_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// Unsupported : gtk_widget_size_allocate : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_size_allocate_with_baseline : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_widget_style_get : unsupported parameter ... : varargs

// Unsupported : gtk_widget_style_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_widget_translate_coordinates : unsupported parameter dest_x : no type generator for gint, gint*

// UnsetStateFlags is a wrapper around the C function gtk_widget_unset_state_flags.
func (recv *Widget) UnsetStateFlags(flags StateFlags) {
	c_flags := (C.GtkStateFlags)(flags)

	C.gtk_widget_unset_state_flags((*C.GtkWidget)(recv.native), c_flags)

	return
}

// GetApplication is a wrapper around the C function gtk_window_get_application.
func (recv *Window) GetApplication() *Application {
	retC := C.gtk_window_get_application((*C.GtkWindow)(recv.native))
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_window_get_default_size : unsupported parameter width : no type generator for gint, gint*

// GetHasResizeGrip is a wrapper around the C function gtk_window_get_has_resize_grip.
func (recv *Window) GetHasResizeGrip() bool {
	retC := C.gtk_window_get_has_resize_grip((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_get_position : unsupported parameter root_x : no type generator for gint, gint*

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_window_get_size : unsupported parameter width : no type generator for gint, gint*

// ResizeGripIsVisible is a wrapper around the C function gtk_window_resize_grip_is_visible.
func (recv *Window) ResizeGripIsVisible() bool {
	retC := C.gtk_window_resize_grip_is_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ResizeToGeometry is a wrapper around the C function gtk_window_resize_to_geometry.
func (recv *Window) ResizeToGeometry(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_resize_to_geometry((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// SetApplication is a wrapper around the C function gtk_window_set_application.
func (recv *Window) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(application.ToC())

	C.gtk_window_set_application((*C.GtkWindow)(recv.native), c_application)

	return
}

// SetDefaultGeometry is a wrapper around the C function gtk_window_set_default_geometry.
func (recv *Window) SetDefaultGeometry(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_set_default_geometry((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// SetHasResizeGrip is a wrapper around the C function gtk_window_set_has_resize_grip.
func (recv *Window) SetHasResizeGrip(value bool) {
	c_value :=
		boolToGboolean(value)

	C.gtk_window_set_has_resize_grip((*C.GtkWindow)(recv.native), c_value)

	return
}

// SetHasUserRefCount is a wrapper around the C function gtk_window_set_has_user_ref_count.
func (recv *Window) SetHasUserRefCount(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_has_user_ref_count((*C.GtkWindow)(recv.native), c_setting)

	return
}

// GetCurrentDeviceGrab is a wrapper around the C function gtk_window_group_get_current_device_grab.
func (recv *WindowGroup) GetCurrentDeviceGrab(device *gdk.Device) *Widget {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gtk_window_group_get_current_device_grab((*C.GtkWindowGroup)(recv.native), c_device)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

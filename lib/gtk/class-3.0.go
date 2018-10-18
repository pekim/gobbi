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
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void CellArea_focusChangedHandler();

	static gulong CellArea_signal_connect_focus_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-changed", CellArea_focusChangedHandler, data);
	}

*/
/*

	void StyleContext_changedHandler();

	static gulong StyleContext_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", StyleContext_changedHandler, data);
	}

*/
/*

	void Widget_drawHandler();

	static gulong Widget_signal_connect_draw(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "draw", Widget_drawHandler, data);
	}

*/
/*

	void Widget_stateFlagsChangedHandler();

	static gulong Widget_signal_connect_state_flags_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-flags-changed", Widget_stateFlagsChangedHandler, data);
	}

*/
/*

	void Widget_styleUpdatedHandler();

	static gulong Widget_signal_connect_style_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-updated", Widget_styleUpdatedHandler, data);
	}

*/
import "C"

// GetLicenseType is a wrapper around the C function gtk_about_dialog_get_license_type.
func (recv *AboutDialog) GetLicenseType() License {
	retC := C.gtk_about_dialog_get_license_type((*C.GtkAboutDialog)(recv.native))
	retGo := (License)(retC)

	return retGo
}

// SetLicenseType is a wrapper around the C function gtk_about_dialog_set_license_type.
func (recv *AboutDialog) SetLicenseType(licenseType License) {
	c_license_type := (C.GtkLicense)(licenseType)

	C.gtk_about_dialog_set_license_type((*C.GtkAboutDialog)(recv.native), c_license_type)

	return
}

// AppChooserButtonNew is a wrapper around the C function gtk_app_chooser_button_new.
func AppChooserButtonNew(contentType string) *AppChooserButton {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_button_new(c_content_type)
	retGo := AppChooserButtonNewFromC(unsafe.Pointer(retC))

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
func AppChooserDialogNewForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	c_parent := (*C.GtkWindow)(parent.ToC())

	c_flags := (C.GtkDialogFlags)(flags)

	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_dialog_new_for_content_type(c_parent, c_flags, c_content_type)
	retGo := AppChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidget is a wrapper around the C function gtk_app_chooser_dialog_get_widget.
func (recv *AppChooserDialog) GetWidget() *Widget {
	retC := C.gtk_app_chooser_dialog_get_widget((*C.GtkAppChooserDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// AppChooserWidgetNew is a wrapper around the C function gtk_app_chooser_widget_new.
func AppChooserWidgetNew(contentType string) *AppChooserWidget {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_widget_new(c_content_type)
	retGo := AppChooserWidgetNewFromC(unsafe.Pointer(retC))

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

// AddWindow is a wrapper around the C function gtk_application_add_window.
func (recv *Application) AddWindow(window *Window) {
	c_window := (*C.GtkWindow)(window.ToC())

	C.gtk_application_add_window((*C.GtkApplication)(recv.native), c_window)

	return
}

// GetWindows is a wrapper around the C function gtk_application_get_windows.
func (recv *Application) GetWindows() *glib.List {
	retC := C.gtk_application_get_windows((*C.GtkApplication)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveWindow is a wrapper around the C function gtk_application_remove_window.
func (recv *Application) RemoveWindow(window *Window) {
	c_window := (*C.GtkWindow)(window.ToC())

	C.gtk_application_remove_window((*C.GtkApplication)(recv.native), c_window)

	return
}

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

// BoxNew is a wrapper around the C function gtk_box_new.
func BoxNew(orientation Orientation, spacing int32) *Box {
	c_orientation := (C.GtkOrientation)(orientation)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_box_new(c_orientation, c_spacing)
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// ButtonBoxNew is a wrapper around the C function gtk_button_box_new.
func ButtonBoxNew(orientation Orientation) *ButtonBox {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_button_box_new(c_orientation)
	retGo := ButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDayIsMarked is a wrapper around the C function gtk_calendar_get_day_is_marked.
func (recv *Calendar) GetDayIsMarked(day uint32) bool {
	c_day := (C.guint)(day)

	retC := C.gtk_calendar_get_day_is_marked((*C.GtkCalendar)(recv.native), c_day)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

var signalCellAreaFocusChangedId int
var signalCellAreaFocusChangedMap = make(map[int]CellAreaSignalFocusChangedCallback)
var signalCellAreaFocusChangedLock sync.Mutex

// CellAreaSignalFocusChangedCallback is a callback function for a 'focus-changed' signal emitted from a CellArea.
type CellAreaSignalFocusChangedCallback func(renderer *CellRenderer, path string)

func (recv *CellArea) ConnectFocusChanged(callback CellAreaSignalFocusChangedCallback) {
	signalCellAreaFocusChangedLock.Lock()
	defer signalCellAreaFocusChangedLock.Unlock()

	signalCellAreaFocusChangedId++
	signalCellAreaFocusChangedMap[signalCellAreaFocusChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.CellArea_signal_connect_focus_changed(instance, C.gpointer(uintptr(signalCellAreaFocusChangedId)))
}

//export CellArea_focusChangedHandler
func CellArea_focusChangedHandler() {}

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

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

// GetPreferredHeight is a wrapper around the C function gtk_cell_area_get_preferred_height.
func (recv *CellArea) GetPreferredHeight(context *CellAreaContext, widget *Widget) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(context.ToC())

	c_widget := (*C.GtkWidget)(widget.ToC())

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_get_preferred_height((*C.GtkCellArea)(recv.native), c_context, c_widget, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_cell_area_get_preferred_height_for_width.
func (recv *CellArea) GetPreferredHeightForWidth(context *CellAreaContext, widget *Widget, width int32) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(context.ToC())

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_get_preferred_height_for_width((*C.GtkCellArea)(recv.native), c_context, c_widget, c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredWidth is a wrapper around the C function gtk_cell_area_get_preferred_width.
func (recv *CellArea) GetPreferredWidth(context *CellAreaContext, widget *Widget) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(context.ToC())

	c_widget := (*C.GtkWidget)(widget.ToC())

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_get_preferred_width((*C.GtkCellArea)(recv.native), c_context, c_widget, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_cell_area_get_preferred_width_for_height.
func (recv *CellArea) GetPreferredWidthForHeight(context *CellAreaContext, widget *Widget, height int32) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(context.ToC())

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_get_preferred_width_for_height((*C.GtkCellArea)(recv.native), c_context, c_widget, c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

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

// RequestRenderer is a wrapper around the C function gtk_cell_area_request_renderer.
func (recv *CellArea) RequestRenderer(renderer *CellRenderer, orientation Orientation, widget *Widget, forSize int32) (int32, int32) {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	c_orientation := (C.GtkOrientation)(orientation)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_for_size := (C.gint)(forSize)

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_area_request_renderer((*C.GtkCellArea)(recv.native), c_renderer, c_orientation, c_widget, c_for_size, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

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
func CellAreaBoxNew() *CellAreaBox {
	retC := C.gtk_cell_area_box_new()
	retGo := CellAreaBoxNewFromC(unsafe.Pointer(retC))

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

// GetAllocation is a wrapper around the C function gtk_cell_area_context_get_allocation.
func (recv *CellAreaContext) GetAllocation() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.gtk_cell_area_context_get_allocation((*C.GtkCellAreaContext)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GetArea is a wrapper around the C function gtk_cell_area_context_get_area.
func (recv *CellAreaContext) GetArea() *CellArea {
	retC := C.gtk_cell_area_context_get_area((*C.GtkCellAreaContext)(recv.native))
	retGo := CellAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPreferredHeight is a wrapper around the C function gtk_cell_area_context_get_preferred_height.
func (recv *CellAreaContext) GetPreferredHeight() (int32, int32) {
	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_context_get_preferred_height((*C.GtkCellAreaContext)(recv.native), &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_cell_area_context_get_preferred_height_for_width.
func (recv *CellAreaContext) GetPreferredHeightForWidth(width int32) (int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_context_get_preferred_height_for_width((*C.GtkCellAreaContext)(recv.native), c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredWidth is a wrapper around the C function gtk_cell_area_context_get_preferred_width.
func (recv *CellAreaContext) GetPreferredWidth() (int32, int32) {
	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_context_get_preferred_width((*C.GtkCellAreaContext)(recv.native), &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_cell_area_context_get_preferred_width_for_height.
func (recv *CellAreaContext) GetPreferredWidthForHeight(height int32) (int32, int32) {
	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_context_get_preferred_width_for_height((*C.GtkCellAreaContext)(recv.native), c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

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

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported : gtk_cell_renderer_get_aligned_area : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// GetPreferredHeight is a wrapper around the C function gtk_cell_renderer_get_preferred_height.
func (recv *CellRenderer) GetPreferredHeight(widget *Widget) (int32, int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_renderer_get_preferred_height((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_cell_renderer_get_preferred_height_for_width.
func (recv *CellRenderer) GetPreferredHeightForWidth(widget *Widget, width int32) (int32, int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_renderer_get_preferred_height_for_width((*C.GtkCellRenderer)(recv.native), c_widget, c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

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

// GetPreferredWidth is a wrapper around the C function gtk_cell_renderer_get_preferred_width.
func (recv *CellRenderer) GetPreferredWidth(widget *Widget) (int32, int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_renderer_get_preferred_width((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_cell_renderer_get_preferred_width_for_height.
func (recv *CellRenderer) GetPreferredWidthForHeight(widget *Widget, height int32) (int32, int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_renderer_get_preferred_width_for_height((*C.GtkCellRenderer)(recv.native), c_widget, c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetRequestMode is a wrapper around the C function gtk_cell_renderer_get_request_mode.
func (recv *CellRenderer) GetRequestMode() SizeRequestMode {
	retC := C.gtk_cell_renderer_get_request_mode((*C.GtkCellRenderer)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

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

// ColorButtonNewWithRgba is a wrapper around the C function gtk_color_button_new_with_rgba.
func ColorButtonNewWithRgba(rgba *gdk.RGBA) *ColorButton {
	c_rgba := (*C.GdkRGBA)(rgba.ToC())

	retC := C.gtk_color_button_new_with_rgba(c_rgba)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

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

// GetPopupFixedWidth is a wrapper around the C function gtk_combo_box_get_popup_fixed_width.
func (recv *ComboBox) GetPopupFixedWidth() bool {
	retC := C.gtk_combo_box_get_popup_fixed_width((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

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

// SetPopupFixedWidth is a wrapper around the C function gtk_combo_box_set_popup_fixed_width.
func (recv *ComboBox) SetPopupFixedWidth(fixed bool) {
	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_combo_box_set_popup_fixed_width((*C.GtkComboBox)(recv.native), c_fixed)

	return
}

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

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_entry_get_icon_area : unsupported parameter icon_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_entry_get_text_area : unsupported parameter text_area : Blacklisted record : GdkRectangle

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// EntryCompletionNewWithArea is a wrapper around the C function gtk_entry_completion_new_with_area.
func EntryCompletionNewWithArea(area *CellArea) *EntryCompletion {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_entry_completion_new_with_area(c_area)
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

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

// IconViewNewWithArea is a wrapper around the C function gtk_icon_view_new_with_area.
func IconViewNewWithArea(area *CellArea) *IconView {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_icon_view_new_with_area(c_area)
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

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

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// PanedNew is a wrapper around the C function gtk_paned_new.
func PanedNew(orientation Orientation) *Paned {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_paned_new(c_orientation)
	retGo := PanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal : unsupported parameter selected_item : no type generator for Gio.File,

// Unsupported signal : unsupported parameter preview : no type generator for PrintOperationPreview,

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

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

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// ScaleNew is a wrapper around the C function gtk_scale_new.
func ScaleNew(orientation Orientation, adjustment *Adjustment) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_scale_new(c_orientation, c_adjustment)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScaleNewWithRange is a wrapper around the C function gtk_scale_new_with_range.
func ScaleNewWithRange(orientation Orientation, min float64, max float64, step float64) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_scale_new_with_range(c_orientation, c_min, c_max, c_step)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// ScrollbarNew is a wrapper around the C function gtk_scrollbar_new.
func ScrollbarNew(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_scrollbar_new(c_orientation, c_adjustment)
	retGo := ScrollbarNewFromC(unsafe.Pointer(retC))

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

// SeparatorNew is a wrapper around the C function gtk_separator_new.
func SeparatorNew(orientation Orientation) *Separator {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_separator_new(c_orientation)
	retGo := SeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// HasContext is a wrapper around the C function gtk_style_has_context.
func (recv *Style) HasContext() bool {
	retC := C.gtk_style_has_context((*C.GtkStyle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

var signalStyleContextChangedId int
var signalStyleContextChangedMap = make(map[int]StyleContextSignalChangedCallback)
var signalStyleContextChangedLock sync.Mutex

// StyleContextSignalChangedCallback is a callback function for a 'changed' signal emitted from a StyleContext.
type StyleContextSignalChangedCallback func()

func (recv *StyleContext) ConnectChanged(callback StyleContextSignalChangedCallback) {
	signalStyleContextChangedLock.Lock()
	defer signalStyleContextChangedLock.Unlock()

	signalStyleContextChangedId++
	signalStyleContextChangedMap[signalStyleContextChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.StyleContext_signal_connect_changed(instance, C.gpointer(uintptr(signalStyleContextChangedId)))
}

//export StyleContext_changedHandler
func StyleContext_changedHandler() {}

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

// StateIsRunning is a wrapper around the C function gtk_style_context_state_is_running.
func (recv *StyleContext) StateIsRunning(state StateType) (bool, float64) {
	c_state := (C.GtkStateType)(state)

	var c_progress C.gdouble

	retC := C.gtk_style_context_state_is_running((*C.GtkStyleContext)(recv.native), c_state, &c_progress)
	retGo := retC == C.TRUE

	progress := (float64)(c_progress)

	return retGo, progress
}

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
func SwitchNew() *Switch {
	retC := C.gtk_switch_new()
	retGo := SwitchNewFromC(unsafe.Pointer(retC))

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

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported : gtk_text_view_get_cursor_locations : unsupported parameter strong : Blacklisted record : GdkRectangle

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

// StateIsRunning is a wrapper around the C function gtk_theming_engine_state_is_running.
func (recv *ThemingEngine) StateIsRunning(state StateType) (bool, float64) {
	c_state := (C.GtkStateType)(state)

	var c_progress C.gdouble

	retC := C.gtk_theming_engine_state_is_running((*C.GtkThemingEngine)(recv.native), c_state, &c_progress)
	retGo := retC == C.TRUE

	progress := (float64)(c_progress)

	return retGo, progress
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_is_blank_at_pos : unsupported parameter path : record with indirection level of 2

// TreeViewColumnNewWithArea is a wrapper around the C function gtk_tree_view_column_new_with_area.
func TreeViewColumnNewWithArea(area *CellArea) *TreeViewColumn {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_tree_view_column_new_with_area(c_area)
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// GetButton is a wrapper around the C function gtk_tree_view_column_get_button.
func (recv *TreeViewColumn) GetButton() *Widget {
	retC := C.gtk_tree_view_column_get_button((*C.GtkTreeViewColumn)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

var signalWidgetDrawId int
var signalWidgetDrawMap = make(map[int]WidgetSignalDrawCallback)
var signalWidgetDrawLock sync.Mutex

// WidgetSignalDrawCallback is a callback function for a 'draw' signal emitted from a Widget.
type WidgetSignalDrawCallback func(cr *cairo.Context) bool

func (recv *Widget) ConnectDraw(callback WidgetSignalDrawCallback) {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	signalWidgetDrawId++
	signalWidgetDrawMap[signalWidgetDrawId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Widget_signal_connect_draw(instance, C.gpointer(uintptr(signalWidgetDrawId)))
}

//export Widget_drawHandler
func Widget_drawHandler() {}

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

var signalWidgetStateFlagsChangedId int
var signalWidgetStateFlagsChangedMap = make(map[int]WidgetSignalStateFlagsChangedCallback)
var signalWidgetStateFlagsChangedLock sync.Mutex

// WidgetSignalStateFlagsChangedCallback is a callback function for a 'state-flags-changed' signal emitted from a Widget.
type WidgetSignalStateFlagsChangedCallback func(flags StateFlags)

func (recv *Widget) ConnectStateFlagsChanged(callback WidgetSignalStateFlagsChangedCallback) {
	signalWidgetStateFlagsChangedLock.Lock()
	defer signalWidgetStateFlagsChangedLock.Unlock()

	signalWidgetStateFlagsChangedId++
	signalWidgetStateFlagsChangedMap[signalWidgetStateFlagsChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Widget_signal_connect_state_flags_changed(instance, C.gpointer(uintptr(signalWidgetStateFlagsChangedId)))
}

//export Widget_stateFlagsChangedHandler
func Widget_stateFlagsChangedHandler() {}

var signalWidgetStyleUpdatedId int
var signalWidgetStyleUpdatedMap = make(map[int]WidgetSignalStyleUpdatedCallback)
var signalWidgetStyleUpdatedLock sync.Mutex

// WidgetSignalStyleUpdatedCallback is a callback function for a 'style-updated' signal emitted from a Widget.
type WidgetSignalStyleUpdatedCallback func()

func (recv *Widget) ConnectStyleUpdated(callback WidgetSignalStyleUpdatedCallback) {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	signalWidgetStyleUpdatedId++
	signalWidgetStyleUpdatedMap[signalWidgetStyleUpdatedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Widget_signal_connect_style_updated(instance, C.gpointer(uintptr(signalWidgetStyleUpdatedId)))
}

//export Widget_styleUpdatedHandler
func Widget_styleUpdatedHandler() {}

// Unsupported signal : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// AddDeviceEvents is a wrapper around the C function gtk_widget_add_device_events.
func (recv *Widget) AddDeviceEvents(device *gdk.Device, events gdk.EventMask) {
	c_device := (*C.GdkDevice)(device.ToC())

	c_events := (C.GdkEventMask)(events)

	C.gtk_widget_add_device_events((*C.GtkWidget)(recv.native), c_device, c_events)

	return
}

// DeviceIsShadowed is a wrapper around the C function gtk_widget_device_is_shadowed.
func (recv *Widget) DeviceIsShadowed(device *gdk.Device) bool {
	c_device := (*C.GdkDevice)(device.ToC())

	retC := C.gtk_widget_device_is_shadowed((*C.GtkWidget)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Draw is a wrapper around the C function gtk_widget_draw.
func (recv *Widget) Draw(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(cr.ToC())

	C.gtk_widget_draw((*C.GtkWidget)(recv.native), c_cr)

	return
}

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

// GetPreferredHeight is a wrapper around the C function gtk_widget_get_preferred_height.
func (recv *Widget) GetPreferredHeight() (int32, int32) {
	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_widget_get_preferred_height((*C.GtkWidget)(recv.native), &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_widget_get_preferred_height_for_width.
func (recv *Widget) GetPreferredHeightForWidth(width int32) (int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_widget_get_preferred_height_for_width((*C.GtkWidget)(recv.native), c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredSize is a wrapper around the C function gtk_widget_get_preferred_size.
func (recv *Widget) GetPreferredSize() (*Requisition, *Requisition) {
	var c_minimum_size C.GtkRequisition

	var c_natural_size C.GtkRequisition

	C.gtk_widget_get_preferred_size((*C.GtkWidget)(recv.native), &c_minimum_size, &c_natural_size)

	minimumSize := RequisitionNewFromC(unsafe.Pointer(&c_minimum_size))

	naturalSize := RequisitionNewFromC(unsafe.Pointer(&c_natural_size))

	return minimumSize, naturalSize
}

// GetPreferredWidth is a wrapper around the C function gtk_widget_get_preferred_width.
func (recv *Widget) GetPreferredWidth() (int32, int32) {
	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_widget_get_preferred_width((*C.GtkWidget)(recv.native), &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_widget_get_preferred_width_for_height.
func (recv *Widget) GetPreferredWidthForHeight(height int32) (int32, int32) {
	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_widget_get_preferred_width_for_height((*C.GtkWidget)(recv.native), c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetRequestMode is a wrapper around the C function gtk_widget_get_request_mode.
func (recv *Widget) GetRequestMode() SizeRequestMode {
	retC := C.gtk_widget_get_request_mode((*C.GtkWidget)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// GetStateFlags is a wrapper around the C function gtk_widget_get_state_flags.
func (recv *Widget) GetStateFlags() StateFlags {
	retC := C.gtk_widget_get_state_flags((*C.GtkWidget)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// InputShapeCombineRegion is a wrapper around the C function gtk_widget_input_shape_combine_region.
func (recv *Widget) InputShapeCombineRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(region.ToC())

	C.gtk_widget_input_shape_combine_region((*C.GtkWidget)(recv.native), c_region)

	return
}

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

// QueueDrawRegion is a wrapper around the C function gtk_widget_queue_draw_region.
func (recv *Widget) QueueDrawRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(region.ToC())

	C.gtk_widget_queue_draw_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// Unsupported : gtk_widget_render_icon_pixbuf : unsupported parameter size : no type generator for gint, GtkIconSize

// ResetStyle is a wrapper around the C function gtk_widget_reset_style.
func (recv *Widget) ResetStyle() {
	C.gtk_widget_reset_style((*C.GtkWidget)(recv.native))

	return
}

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

// GetHasResizeGrip is a wrapper around the C function gtk_window_get_has_resize_grip.
func (recv *Window) GetHasResizeGrip() bool {
	retC := C.gtk_window_get_has_resize_grip((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_window_get_resize_grip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

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

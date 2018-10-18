// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
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

	void LevelBar_offsetChangedHandler();

	static gulong LevelBar_signal_connect_offset_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "offset-changed", LevelBar_offsetChangedHandler, data);
	}

*/
import "C"

// SetAccel is a wrapper around the C function gtk_accel_label_set_accel.
func (recv *AccelLabel) SetAccel(acceleratorKey uint32, acceleratorMods gdk.ModifierType) {
	c_accelerator_key := (C.guint)(acceleratorKey)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	C.gtk_accel_label_set_accel((*C.GtkAccelLabel)(recv.native), c_accelerator_key, c_accelerator_mods)

	return
}

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

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal 'application-activated' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'application-selected' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'populate-popup' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

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

// GetId is a wrapper around the C function gtk_application_window_get_id.
func (recv *ApplicationWindow) GetId() uint32 {
	retC := C.gtk_application_window_get_id((*C.GtkApplicationWindow)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

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

// Unsupported signal 'add-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'apply-attributes' for CellArea : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'remove-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'editing-started' for CellRenderer : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'format-entry-text' for ComboBox : return value utf8 :

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetAttributes is a wrapper around the C function gtk_entry_get_attributes.
func (recv *Entry) GetAttributes() *pango.AttrList {
	retC := C.gtk_entry_get_attributes((*C.GtkEntry)(recv.native))
	retGo := pango.AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// SetAttributes is a wrapper around the C function gtk_entry_set_attributes.
func (recv *Entry) SetAttributes(attrs *pango.AttrList) {
	c_attrs := (*C.PangoAttrList)(attrs.ToC())

	C.gtk_entry_set_attributes((*C.GtkEntry)(recv.native), c_attrs)

	return
}

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

// Unsupported signal 'cursor-on-match' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'match-selected' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_icon_view_get_cell_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

var signalLevelBarOffsetChangedId int
var signalLevelBarOffsetChangedMap = make(map[int]LevelBarSignalOffsetChangedCallback)
var signalLevelBarOffsetChangedLock sync.Mutex

// LevelBarSignalOffsetChangedCallback is a callback function for a 'offset-changed' signal emitted from a LevelBar.
type LevelBarSignalOffsetChangedCallback func(name string)

/*
ConnectOffsetChanged connects the callback to the 'offset-changed' signal for the LevelBar.

The returned value represents the connection, and may be passed to DisconnectOffsetChanged to remove it.
*/
func (recv *LevelBar) ConnectOffsetChanged(callback LevelBarSignalOffsetChangedCallback) int {
	signalLevelBarOffsetChangedLock.Lock()
	defer signalLevelBarOffsetChangedLock.Unlock()

	signalLevelBarOffsetChangedId++
	signalLevelBarOffsetChangedMap[signalLevelBarOffsetChangedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.LevelBar_signal_connect_offset_changed(instance, C.gpointer(uintptr(signalLevelBarOffsetChangedId)))
	return int(retC)
}

/*
DisconnectOffsetChanged disconnects a callback from the 'offset-changed' signal for the LevelBar.

The connectionID should be a value returned from a call to ConnectOffsetChanged.
*/
func (recv *LevelBar) DisconnectOffsetChanged(connectionID int) {
	_, exists := signalLevelBarOffsetChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalLevelBarOffsetChangedMap, connectionID)
}

//export LevelBar_offsetChangedHandler
func LevelBar_offsetChangedHandler() {
	fmt.Println("cb")
}

// LevelBarNew is a wrapper around the C function gtk_level_bar_new.
func LevelBarNew() *LevelBar {
	retC := C.gtk_level_bar_new()
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LevelBarNewForInterval is a wrapper around the C function gtk_level_bar_new_for_interval.
func LevelBarNewForInterval(minValue float64, maxValue float64) *LevelBar {
	c_min_value := (C.gdouble)(minValue)

	c_max_value := (C.gdouble)(maxValue)

	retC := C.gtk_level_bar_new_for_interval(c_min_value, c_max_value)
	retGo := LevelBarNewFromC(unsafe.Pointer(retC))

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

// GetOffsetValue is a wrapper around the C function gtk_level_bar_get_offset_value.
func (recv *LevelBar) GetOffsetValue(name string) (bool, float64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_value C.gdouble

	retC := C.gtk_level_bar_get_offset_value((*C.GtkLevelBar)(recv.native), c_name, &c_value)
	retGo := retC == C.TRUE

	value := (float64)(c_value)

	return retGo, value
}

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

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// MenuButtonNew is a wrapper around the C function gtk_menu_button_new.
func MenuButtonNew() *MenuButton {
	retC := C.gtk_menu_button_new()
	retGo := MenuButtonNewFromC(unsafe.Pointer(retC))

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

// Unsupported signal 'create-window' for Notebook : return value Notebook :

// Unsupported signal 'get-child-position' for Overlay : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported signal 'drag-action-ask' for PlacesSidebar : return value gint :

// Unsupported signal 'drag-action-requested' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal 'open-location' for PlacesSidebar : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal 'populate-popup' for PlacesSidebar : unsupported parameter selected_item : no type generator for Gio.File,

// Unsupported signal 'create-custom-widget' for PrintOperation : return value GObject.Object :

// Unsupported signal 'preview' for PrintOperation : unsupported parameter preview : no type generator for PrintOperationPreview,

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported signal 'format-value' for Scale : return value utf8 :

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// SearchEntryNew is a wrapper around the C function gtk_search_entry_new.
func SearchEntryNew() *SearchEntry {
	retC := C.gtk_search_entry_new()
	retGo := SearchEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'input' for SpinButton : return value gint :

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported signal 'event' for TextTag : unsupported parameter event : no type generator for Gdk.Event,

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

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Unsupported signal 'child-notify' for Widget : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal 'delete-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'destroy-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event-after' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'size-allocate' for Widget : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal 'touch-event' for Widget : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_widget_insert_action_group : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

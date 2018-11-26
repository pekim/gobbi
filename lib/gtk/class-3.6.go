// This is a generated file - DO NOT EDIT
// +build gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
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

	void levelbar_offsetChangedHandler(GObject *, gchar*, gpointer);

	static gulong LevelBar_signal_connect_offset_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "offset-changed", G_CALLBACK(levelbar_offsetChangedHandler), data);
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
	c_accel_group := (*C.GtkAccelGroup)(C.NULL)
	if accelGroup != nil {
		c_accel_group = (*C.GtkAccelGroup)(accelGroup.ToC())
	}

	C.gtk_action_group_set_accel_group((*C.GtkActionGroup)(recv.native), c_accel_group)

	return
}

// GetActiveWindow is a wrapper around the C function gtk_application_get_active_window.
func (recv *Application) GetActiveWindow() *Window {
	retC := C.gtk_application_get_active_window((*C.GtkApplication)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetWindowById is a wrapper around the C function gtk_application_get_window_by_id.
func (recv *Application) GetWindowById(id uint32) *Window {
	c_id := (C.guint)(id)

	retC := C.gtk_application_get_window_by_id((*C.GtkApplication)(recv.native), c_id)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetId is a wrapper around the C function gtk_application_window_get_id.
func (recv *ApplicationWindow) GetId() uint32 {
	retC := C.gtk_application_window_get_id((*C.GtkApplicationWindow)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

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

// GetAttributes is a wrapper around the C function gtk_entry_get_attributes.
func (recv *Entry) GetAttributes() *pango.AttrList {
	retC := C.gtk_entry_get_attributes((*C.GtkEntry)(recv.native))
	var retGo (*pango.AttrList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.AttrListNewFromC(unsafe.Pointer(retC))
	}

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
	c_attrs := (*C.PangoAttrList)(C.NULL)
	if attrs != nil {
		c_attrs = (*C.PangoAttrList)(attrs.ToC())
	}

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

// Unsupported : gtk_icon_view_get_cell_rect : unsupported parameter rect : Blacklisted record : GdkRectangle

type signalLevelBarOffsetChangedDetail struct {
	callback  LevelBarSignalOffsetChangedCallback
	handlerID C.gulong
}

var signalLevelBarOffsetChangedId int
var signalLevelBarOffsetChangedMap = make(map[int]signalLevelBarOffsetChangedDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.LevelBar_signal_connect_offset_changed(instance, C.gpointer(uintptr(signalLevelBarOffsetChangedId)))

	detail := signalLevelBarOffsetChangedDetail{callback, handlerID}
	signalLevelBarOffsetChangedMap[signalLevelBarOffsetChangedId] = detail

	return signalLevelBarOffsetChangedId
}

/*
DisconnectOffsetChanged disconnects a callback from the 'offset-changed' signal for the LevelBar.

The connectionID should be a value returned from a call to ConnectOffsetChanged.
*/
func (recv *LevelBar) DisconnectOffsetChanged(connectionID int) {
	signalLevelBarOffsetChangedLock.Lock()
	defer signalLevelBarOffsetChangedLock.Unlock()

	detail, exists := signalLevelBarOffsetChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalLevelBarOffsetChangedMap, connectionID)
}

//export levelbar_offsetChangedHandler
func levelbar_offsetChangedHandler(_ *C.GObject, c_name *C.gchar, data C.gpointer) {
	name := C.GoString(c_name)

	index := int(uintptr(data))
	callback := signalLevelBarOffsetChangedMap[index].callback
	callback(name)
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

// MenuButtonNew is a wrapper around the C function gtk_menu_button_new.
func MenuButtonNew() *MenuButton {
	retC := C.gtk_menu_button_new()
	retGo := MenuButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAlignWidget is a wrapper around the C function gtk_menu_button_get_align_widget.
func (recv *MenuButton) GetAlignWidget() *Widget {
	retC := C.gtk_menu_button_get_align_widget((*C.GtkMenuButton)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

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
	var retGo (*gio.MenuModel)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.MenuModelNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetPopup is a wrapper around the C function gtk_menu_button_get_popup.
func (recv *MenuButton) GetPopup() *Menu {
	retC := C.gtk_menu_button_get_popup((*C.GtkMenuButton)(recv.native))
	var retGo (*Menu)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MenuNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetAlignWidget is a wrapper around the C function gtk_menu_button_set_align_widget.
func (recv *MenuButton) SetAlignWidget(alignWidget *Widget) {
	c_align_widget := (*C.GtkWidget)(C.NULL)
	if alignWidget != nil {
		c_align_widget = (*C.GtkWidget)(alignWidget.ToC())
	}

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
	c_menu_model := (*C.GMenuModel)(C.NULL)
	if menuModel != nil {
		c_menu_model = (*C.GMenuModel)(menuModel.ToC())
	}

	C.gtk_menu_button_set_menu_model((*C.GtkMenuButton)(recv.native), c_menu_model)

	return
}

// SetPopup is a wrapper around the C function gtk_menu_button_set_popup.
func (recv *MenuButton) SetPopup(menu *Widget) {
	c_menu := (*C.GtkWidget)(C.NULL)
	if menu != nil {
		c_menu = (*C.GtkWidget)(menu.ToC())
	}

	C.gtk_menu_button_set_popup((*C.GtkMenuButton)(recv.native), c_menu)

	return
}

// BindModel is a wrapper around the C function gtk_menu_shell_bind_model.
func (recv *MenuShell) BindModel(model *gio.MenuModel, actionNamespace string, withSeparators bool) {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	c_action_namespace := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(c_action_namespace))

	c_with_separators :=
		boolToGboolean(withSeparators)

	C.gtk_menu_shell_bind_model((*C.GtkMenuShell)(recv.native), c_model, c_action_namespace, c_with_separators)

	return
}

// SearchEntryNew is a wrapper around the C function gtk_search_entry_new.
func SearchEntryNew() *SearchEntry {
	retC := C.gtk_search_entry_new()
	retGo := SearchEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// InsertActionGroup is a wrapper around the C function gtk_widget_insert_action_group.
func (recv *Widget) InsertActionGroup(name string, group *gio.ActionGroup) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_group := (*C.GActionGroup)(group.ToC())

	C.gtk_widget_insert_action_group((*C.GtkWidget)(recv.native), c_name, c_group)

	return
}

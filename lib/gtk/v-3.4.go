// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	static void _gtk_actionable_set_action_target(GtkActionable* actionable, const gchar* format_string) {
		return gtk_actionable_set_action_target(actionable, format_string);
    }
*/
/*

	void colorchooser_colorActivatedHandler(GObject *, GdkRGBA *, gpointer);

	static gulong ColorChooser_signal_connect_color_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "color-activated", G_CALLBACK(colorchooser_colorActivatedHandler), data);
	}

*/
import "C"

type ApplicationInhibitFlags C.GtkApplicationInhibitFlags

const (
	GTK_APPLICATION_INHIBIT_LOGOUT  ApplicationInhibitFlags = 1
	GTK_APPLICATION_INHIBIT_SWITCH  ApplicationInhibitFlags = 2
	GTK_APPLICATION_INHIBIT_SUSPEND ApplicationInhibitFlags = 4
	GTK_APPLICATION_INHIBIT_IDLE    ApplicationInhibitFlags = 8
)

// AddCreditSection is a wrapper around the C function gtk_about_dialog_add_credit_section.
func (recv *AboutDialog) AddCreditSection(sectionName string, people []string) {
	c_section_name := C.CString(sectionName)
	defer C.free(unsafe.Pointer(c_section_name))

	c_people_array := make([]*C.gchar, len(people)+1, len(people)+1)
	for i, item := range people {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_people_array[i] = c
	}
	c_people_array[len(people)] = nil
	c_people_arrayPtr := &c_people_array[0]
	c_people := (**C.gchar)(unsafe.Pointer(c_people_arrayPtr))

	C.gtk_about_dialog_add_credit_section((*C.GtkAboutDialog)(recv.native), c_section_name, c_people)

	return
}

// AddAccelerator is a wrapper around the C function gtk_application_add_accelerator.
func (recv *Application) AddAccelerator(accelerator string, actionName string, parameter *glib.Variant) {
	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_parameter := (*C.GVariant)(C.NULL)
	if parameter != nil {
		c_parameter = (*C.GVariant)(parameter.ToC())
	}

	C.gtk_application_add_accelerator((*C.GtkApplication)(recv.native), c_accelerator, c_action_name, c_parameter)

	return
}

// GetAppMenu is a wrapper around the C function gtk_application_get_app_menu.
func (recv *Application) GetAppMenu() *gio.MenuModel {
	retC := C.gtk_application_get_app_menu((*C.GtkApplication)(recv.native))
	var retGo (*gio.MenuModel)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.MenuModelNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetMenubar is a wrapper around the C function gtk_application_get_menubar.
func (recv *Application) GetMenubar() *gio.MenuModel {
	retC := C.gtk_application_get_menubar((*C.GtkApplication)(recv.native))
	retGo := gio.MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inhibit is a wrapper around the C function gtk_application_inhibit.
func (recv *Application) Inhibit(window *Window, flags ApplicationInhibitFlags, reason string) uint32 {
	c_window := (*C.GtkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GtkWindow)(window.ToC())
	}

	c_flags := (C.GtkApplicationInhibitFlags)(flags)

	c_reason := C.CString(reason)
	defer C.free(unsafe.Pointer(c_reason))

	retC := C.gtk_application_inhibit((*C.GtkApplication)(recv.native), c_window, c_flags, c_reason)
	retGo := (uint32)(retC)

	return retGo
}

// IsInhibited is a wrapper around the C function gtk_application_is_inhibited.
func (recv *Application) IsInhibited(flags ApplicationInhibitFlags) bool {
	c_flags := (C.GtkApplicationInhibitFlags)(flags)

	retC := C.gtk_application_is_inhibited((*C.GtkApplication)(recv.native), c_flags)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveAccelerator is a wrapper around the C function gtk_application_remove_accelerator.
func (recv *Application) RemoveAccelerator(actionName string, parameter *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_parameter := (*C.GVariant)(C.NULL)
	if parameter != nil {
		c_parameter = (*C.GVariant)(parameter.ToC())
	}

	C.gtk_application_remove_accelerator((*C.GtkApplication)(recv.native), c_action_name, c_parameter)

	return
}

// SetAppMenu is a wrapper around the C function gtk_application_set_app_menu.
func (recv *Application) SetAppMenu(appMenu *gio.MenuModel) {
	c_app_menu := (*C.GMenuModel)(C.NULL)
	if appMenu != nil {
		c_app_menu = (*C.GMenuModel)(appMenu.ToC())
	}

	C.gtk_application_set_app_menu((*C.GtkApplication)(recv.native), c_app_menu)

	return
}

// SetMenubar is a wrapper around the C function gtk_application_set_menubar.
func (recv *Application) SetMenubar(menubar *gio.MenuModel) {
	c_menubar := (*C.GMenuModel)(C.NULL)
	if menubar != nil {
		c_menubar = (*C.GMenuModel)(menubar.ToC())
	}

	C.gtk_application_set_menubar((*C.GtkApplication)(recv.native), c_menubar)

	return
}

// Uninhibit is a wrapper around the C function gtk_application_uninhibit.
func (recv *Application) Uninhibit(cookie uint32) {
	c_cookie := (C.guint)(cookie)

	C.gtk_application_uninhibit((*C.GtkApplication)(recv.native), c_cookie)

	return
}

// ApplicationWindowNew is a wrapper around the C function gtk_application_window_new.
func ApplicationWindowNew(application *Application) *ApplicationWindow {
	c_application := (*C.GtkApplication)(C.NULL)
	if application != nil {
		c_application = (*C.GtkApplication)(application.ToC())
	}

	retC := C.gtk_application_window_new(c_application)
	retGo := ApplicationWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShowMenubar is a wrapper around the C function gtk_application_window_get_show_menubar.
func (recv *ApplicationWindow) GetShowMenubar() bool {
	retC := C.gtk_application_window_get_show_menubar((*C.GtkApplicationWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowMenubar is a wrapper around the C function gtk_application_window_set_show_menubar.
func (recv *ApplicationWindow) SetShowMenubar(showMenubar bool) {
	c_show_menubar :=
		boolToGboolean(showMenubar)

	C.gtk_application_window_set_show_menubar((*C.GtkApplicationWindow)(recv.native), c_show_menubar)

	return
}

// AddFromResource is a wrapper around the C function gtk_builder_add_from_resource.
func (recv *Builder) AddFromResource(resourcePath string) (uint32, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_from_resource((*C.GtkBuilder)(recv.native), c_resource_path, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddObjectsFromResource is a wrapper around the C function gtk_builder_add_objects_from_resource.
func (recv *Builder) AddObjectsFromResource(resourcePath string, objectIds []string) (uint32, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	c_object_ids_array := make([]*C.gchar, len(objectIds)+1, len(objectIds)+1)
	for i, item := range objectIds {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_object_ids_array[i] = c
	}
	c_object_ids_array[len(objectIds)] = nil
	c_object_ids_arrayPtr := &c_object_ids_array[0]
	c_object_ids := (**C.gchar)(unsafe.Pointer(c_object_ids_arrayPtr))

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_objects_from_resource((*C.GtkBuilder)(recv.native), c_resource_path, c_object_ids, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ColorChooserDialogNew is a wrapper around the C function gtk_color_chooser_dialog_new.
func ColorChooserDialogNew(title string, parent *Window) *ColorChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_color_chooser_dialog_new(c_title, c_parent)
	retGo := ColorChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorChooserWidgetNew is a wrapper around the C function gtk_color_chooser_widget_new.
func ColorChooserWidgetNew() *ColorChooserWidget {
	retC := C.gtk_color_chooser_widget_new()
	retGo := ColorChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'format-entry-text' for ComboBox : return value utf8 :

// ComputePrefix is a wrapper around the C function gtk_entry_completion_compute_prefix.
func (recv *EntryCompletion) ComputePrefix(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_entry_completion_compute_prefix((*C.GtkEntryCompletion)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromResource is a wrapper around the C function gtk_image_new_from_resource.
func ImageNewFromResource(resourcePath string) *Image {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_image_new_from_resource(c_resource_path)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuNewFromModel is a wrapper around the C function gtk_menu_new_from_model.
func MenuNewFromModel(model *gio.MenuModel) *Menu {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	retC := C.gtk_menu_new_from_model(c_model)
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuBarNewFromModel is a wrapper around the C function gtk_menu_bar_new_from_model.
func MenuBarNewFromModel(model *gio.MenuModel) *MenuBar {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	retC := C.gtk_menu_bar_new_from_model(c_model)
	retGo := MenuBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHasOrigin is a wrapper around the C function gtk_scale_get_has_origin.
func (recv *Scale) GetHasOrigin() bool {
	retC := C.gtk_scale_get_has_origin((*C.GtkScale)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetHasOrigin is a wrapper around the C function gtk_scale_set_has_origin.
func (recv *Scale) SetHasOrigin(hasOrigin bool) {
	c_has_origin :=
		boolToGboolean(hasOrigin)

	C.gtk_scale_set_has_origin((*C.GtkScale)(recv.native), c_has_origin)

	return
}

// GetCaptureButtonPress is a wrapper around the C function gtk_scrolled_window_get_capture_button_press.
func (recv *ScrolledWindow) GetCaptureButtonPress() bool {
	retC := C.gtk_scrolled_window_get_capture_button_press((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetKineticScrolling is a wrapper around the C function gtk_scrolled_window_get_kinetic_scrolling.
func (recv *ScrolledWindow) GetKineticScrolling() bool {
	retC := C.gtk_scrolled_window_get_kinetic_scrolling((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetCaptureButtonPress is a wrapper around the C function gtk_scrolled_window_set_capture_button_press.
func (recv *ScrolledWindow) SetCaptureButtonPress(captureButtonPress bool) {
	c_capture_button_press :=
		boolToGboolean(captureButtonPress)

	C.gtk_scrolled_window_set_capture_button_press((*C.GtkScrolledWindow)(recv.native), c_capture_button_press)

	return
}

// SetKineticScrolling is a wrapper around the C function gtk_scrolled_window_set_kinetic_scrolling.
func (recv *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	c_kinetic_scrolling :=
		boolToGboolean(kineticScrolling)

	C.gtk_scrolled_window_set_kinetic_scrolling((*C.GtkScrolledWindow)(recv.native), c_kinetic_scrolling)

	return
}

// GetParent is a wrapper around the C function gtk_style_context_get_parent.
func (recv *StyleContext) GetParent() *StyleContext {
	retC := C.gtk_style_context_get_parent((*C.GtkStyleContext)(recv.native))
	var retGo (*StyleContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StyleContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetParent is a wrapper around the C function gtk_style_context_set_parent.
func (recv *StyleContext) SetParent(parent *StyleContext) {
	c_parent := (*C.GtkStyleContext)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkStyleContext)(parent.ToC())
	}

	C.gtk_style_context_set_parent((*C.GtkStyleContext)(recv.native), c_parent)

	return
}

// GetNColumns is a wrapper around the C function gtk_tree_view_get_n_columns.
func (recv *TreeView) GetNColumns() uint32 {
	retC := C.gtk_tree_view_get_n_columns((*C.GtkTreeView)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// AddUiFromResource is a wrapper around the C function gtk_ui_manager_add_ui_from_resource.
func (recv *UIManager) AddUiFromResource(resourcePath string) (uint32, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	var cThrowableError *C.GError

	retC := C.gtk_ui_manager_add_ui_from_resource((*C.GtkUIManager)(recv.native), c_resource_path, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetModifierMask is a wrapper around the C function gtk_widget_get_modifier_mask.
func (recv *Widget) GetModifierMask(intent gdk.ModifierIntent) gdk.ModifierType {
	c_intent := (C.GdkModifierIntent)(intent)

	retC := C.gtk_widget_get_modifier_mask((*C.GtkWidget)(recv.native), c_intent)
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// GetAttachedTo is a wrapper around the C function gtk_window_get_attached_to.
func (recv *Window) GetAttachedTo() *Widget {
	retC := C.gtk_window_get_attached_to((*C.GtkWindow)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetHideTitlebarWhenMaximized is a wrapper around the C function gtk_window_get_hide_titlebar_when_maximized.
func (recv *Window) GetHideTitlebarWhenMaximized() bool {
	retC := C.gtk_window_get_hide_titlebar_when_maximized((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAttachedTo is a wrapper around the C function gtk_window_set_attached_to.
func (recv *Window) SetAttachedTo(attachWidget *Widget) {
	c_attach_widget := (*C.GtkWidget)(C.NULL)
	if attachWidget != nil {
		c_attach_widget = (*C.GtkWidget)(attachWidget.ToC())
	}

	C.gtk_window_set_attached_to((*C.GtkWindow)(recv.native), c_attach_widget)

	return
}

// SetHideTitlebarWhenMaximized is a wrapper around the C function gtk_window_set_hide_titlebar_when_maximized.
func (recv *Window) SetHideTitlebarWhenMaximized(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_hide_titlebar_when_maximized((*C.GtkWindow)(recv.native), c_setting)

	return
}

// AcceleratorGetLabelWithKeycode is a wrapper around the C function gtk_accelerator_get_label_with_keycode.
func AcceleratorGetLabelWithKeycode(display *gdk.Display, acceleratorKey uint32, keycode uint32, acceleratorMods gdk.ModifierType) string {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_accelerator_key := (C.guint)(acceleratorKey)

	c_keycode := (C.guint)(keycode)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_get_label_with_keycode(c_display, c_accelerator_key, c_keycode, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AcceleratorNameWithKeycode is a wrapper around the C function gtk_accelerator_name_with_keycode.
func AcceleratorNameWithKeycode(display *gdk.Display, acceleratorKey uint32, keycode uint32, acceleratorMods gdk.ModifierType) string {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_accelerator_key := (C.guint)(acceleratorKey)

	c_keycode := (C.guint)(keycode)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_name_with_keycode(c_display, c_accelerator_key, c_keycode, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_codes : output array param accelerator_codes

// RenderInsertionCursor is a wrapper around the C function gtk_render_insertion_cursor.
func RenderInsertionCursor(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout, index int32, direction pango.Direction) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	c_index := (C.int)(index)

	c_direction := (C.PangoDirection)(direction)

	C.gtk_render_insertion_cursor(c_context, c_cr, c_x, c_y, c_layout, c_index, c_direction)

	return
}

// GetActionName is a wrapper around the C function gtk_actionable_get_action_name.
func (recv *Actionable) GetActionName() string {
	retC := C.gtk_actionable_get_action_name((*C.GtkActionable)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetActionTargetValue is a wrapper around the C function gtk_actionable_get_action_target_value.
func (recv *Actionable) GetActionTargetValue() *glib.Variant {
	retC := C.gtk_actionable_get_action_target_value((*C.GtkActionable)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetActionName is a wrapper around the C function gtk_actionable_set_action_name.
func (recv *Actionable) SetActionName(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.gtk_actionable_set_action_name((*C.GtkActionable)(recv.native), c_action_name)

	return
}

// SetActionTarget is a wrapper around the C function gtk_actionable_set_action_target.
func (recv *Actionable) SetActionTarget(formatString string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(formatString, args...)
	c_format_string := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format_string))

	C._gtk_actionable_set_action_target((*C.GtkActionable)(recv.native), c_format_string)

	return
}

// SetActionTargetValue is a wrapper around the C function gtk_actionable_set_action_target_value.
func (recv *Actionable) SetActionTargetValue(targetValue *glib.Variant) {
	c_target_value := (*C.GVariant)(C.NULL)
	if targetValue != nil {
		c_target_value = (*C.GVariant)(targetValue.ToC())
	}

	C.gtk_actionable_set_action_target_value((*C.GtkActionable)(recv.native), c_target_value)

	return
}

// SetDetailedActionName is a wrapper around the C function gtk_actionable_set_detailed_action_name.
func (recv *Actionable) SetDetailedActionName(detailedActionName string) {
	c_detailed_action_name := C.CString(detailedActionName)
	defer C.free(unsafe.Pointer(c_detailed_action_name))

	C.gtk_actionable_set_detailed_action_name((*C.GtkActionable)(recv.native), c_detailed_action_name)

	return
}

type signalColorChooserColorActivatedDetail struct {
	callback  ColorChooserSignalColorActivatedCallback
	handlerID C.gulong
}

var signalColorChooserColorActivatedId int
var signalColorChooserColorActivatedMap = make(map[int]signalColorChooserColorActivatedDetail)
var signalColorChooserColorActivatedLock sync.RWMutex

// ColorChooserSignalColorActivatedCallback is a callback function for a 'color-activated' signal emitted from a ColorChooser.
type ColorChooserSignalColorActivatedCallback func(color *gdk.RGBA)

/*
ConnectColorActivated connects the callback to the 'color-activated' signal for the ColorChooser.

The returned value represents the connection, and may be passed to DisconnectColorActivated to remove it.
*/
func (recv *ColorChooser) ConnectColorActivated(callback ColorChooserSignalColorActivatedCallback) int {
	signalColorChooserColorActivatedLock.Lock()
	defer signalColorChooserColorActivatedLock.Unlock()

	signalColorChooserColorActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ColorChooser_signal_connect_color_activated(instance, C.gpointer(uintptr(signalColorChooserColorActivatedId)))

	detail := signalColorChooserColorActivatedDetail{callback, handlerID}
	signalColorChooserColorActivatedMap[signalColorChooserColorActivatedId] = detail

	return signalColorChooserColorActivatedId
}

/*
DisconnectColorActivated disconnects a callback from the 'color-activated' signal for the ColorChooser.

The connectionID should be a value returned from a call to ConnectColorActivated.
*/
func (recv *ColorChooser) DisconnectColorActivated(connectionID int) {
	signalColorChooserColorActivatedLock.Lock()
	defer signalColorChooserColorActivatedLock.Unlock()

	detail, exists := signalColorChooserColorActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalColorChooserColorActivatedMap, connectionID)
}

//export colorchooser_colorActivatedHandler
func colorchooser_colorActivatedHandler(_ *C.GObject, c_color *C.GdkRGBA, data C.gpointer) {
	signalColorChooserColorActivatedLock.RLock()
	defer signalColorChooserColorActivatedLock.RUnlock()

	color := gdk.RGBANewFromC(unsafe.Pointer(c_color))

	index := int(uintptr(data))
	callback := signalColorChooserColorActivatedMap[index].callback
	callback(color)
}

// Unsupported : gtk_color_chooser_add_palette : unsupported parameter colors :

// GetRgba is a wrapper around the C function gtk_color_chooser_get_rgba.
func (recv *ColorChooser) GetRgba() *gdk.RGBA {
	var c_color C.GdkRGBA

	C.gtk_color_chooser_get_rgba((*C.GtkColorChooser)(recv.native), &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetUseAlpha is a wrapper around the C function gtk_color_chooser_get_use_alpha.
func (recv *ColorChooser) GetUseAlpha() bool {
	retC := C.gtk_color_chooser_get_use_alpha((*C.GtkColorChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetRgba is a wrapper around the C function gtk_color_chooser_set_rgba.
func (recv *ColorChooser) SetRgba(color *gdk.RGBA) {
	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_color_chooser_set_rgba((*C.GtkColorChooser)(recv.native), c_color)

	return
}

// SetUseAlpha is a wrapper around the C function gtk_color_chooser_set_use_alpha.
func (recv *ColorChooser) SetUseAlpha(useAlpha bool) {
	c_use_alpha :=
		boolToGboolean(useAlpha)

	C.gtk_color_chooser_set_use_alpha((*C.GtkColorChooser)(recv.native), c_use_alpha)

	return
}

// SymbolicColorNewWin32 is a wrapper around the C function gtk_symbolic_color_new_win32.
func SymbolicColorNewWin32(themeClass string, id int32) *SymbolicColor {
	c_theme_class := C.CString(themeClass)
	defer C.free(unsafe.Pointer(c_theme_class))

	c_id := (C.gint)(id)

	retC := C.gtk_symbolic_color_new_win32(c_theme_class, c_id)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

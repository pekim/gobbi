// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_about_dialog_add_credit_section : unsupported parameter people : no param type

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_application_add_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// GetAppMenu is a wrapper around the C function gtk_application_get_app_menu.
func (recv *Application) GetAppMenu() *gio.MenuModel {
	retC := C.gtk_application_get_app_menu((*C.GtkApplication)(recv.native))
	retGo := gio.MenuModelNewFromC(unsafe.Pointer(retC))

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
	c_window := (*C.GtkWindow)(window.ToC())

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

// Unsupported : gtk_application_remove_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// SetAppMenu is a wrapper around the C function gtk_application_set_app_menu.
func (recv *Application) SetAppMenu(appMenu *gio.MenuModel) {
	c_app_menu := (*C.GMenuModel)(appMenu.ToC())

	C.gtk_application_set_app_menu((*C.GtkApplication)(recv.native), c_app_menu)

	return
}

// SetMenubar is a wrapper around the C function gtk_application_set_menubar.
func (recv *Application) SetMenubar(menubar *gio.MenuModel) {
	c_menubar := (*C.GMenuModel)(menubar.ToC())

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
	c_application := (*C.GtkApplication)(application.ToC())

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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_builder_add_objects_from_resource : unsupported parameter object_ids : no param type

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// ColorChooserDialogNew is a wrapper around the C function gtk_color_chooser_dialog_new.
func ColorChooserDialogNew(title string, parent *Window) *ColorChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(parent.ToC())

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

// Unsupported signal 'format-entry-text' for ComboBox : unsupported parameter path : type utf8 :

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// ComputePrefix is a wrapper around the C function gtk_entry_completion_compute_prefix.
func (recv *EntryCompletion) ComputePrefix(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_entry_completion_compute_prefix((*C.GtkEntryCompletion)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// ImageNewFromResource is a wrapper around the C function gtk_image_new_from_resource.
func ImageNewFromResource(resourcePath string) *Image {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_image_new_from_resource(c_resource_path)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// MenuNewFromModel is a wrapper around the C function gtk_menu_new_from_model.
func MenuNewFromModel(model *gio.MenuModel) *Menu {
	c_model := (*C.GMenuModel)(model.ToC())

	retC := C.gtk_menu_new_from_model(c_model)
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuBarNewFromModel is a wrapper around the C function gtk_menu_bar_new_from_model.
func MenuBarNewFromModel(model *gio.MenuModel) *MenuBar {
	c_model := (*C.GMenuModel)(model.ToC())

	retC := C.gtk_menu_bar_new_from_model(c_model)
	retGo := MenuBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

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

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

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

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// GetParent is a wrapper around the C function gtk_style_context_get_parent.
func (recv *StyleContext) GetParent() *StyleContext {
	retC := C.gtk_style_context_get_parent((*C.GtkStyleContext)(recv.native))
	retGo := StyleContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetParent is a wrapper around the C function gtk_style_context_set_parent.
func (recv *StyleContext) SetParent(parent *StyleContext) {
	c_parent := (*C.GtkStyleContext)(parent.ToC())

	C.gtk_style_context_set_parent((*C.GtkStyleContext)(recv.native), c_parent)

	return
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetNColumns is a wrapper around the C function gtk_tree_view_get_n_columns.
func (recv *TreeView) GetNColumns() uint32 {
	retC := C.gtk_tree_view_get_n_columns((*C.GtkTreeView)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// AddUiFromResource is a wrapper around the C function gtk_ui_manager_add_ui_from_resource.
func (recv *UIManager) AddUiFromResource(resourcePath string) (uint32, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	var cThrowableError *C.GError

	retC := C.gtk_ui_manager_add_ui_from_resource((*C.GtkUIManager)(recv.native), c_resource_path, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

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
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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
	c_attach_widget := (*C.GtkWidget)(attachWidget.ToC())

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

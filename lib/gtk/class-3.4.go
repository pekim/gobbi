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

// Unsupported : gtk_about_dialog_add_credit_section : unsupported parameter people :

// Unsupported : gtk_application_add_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Returns the menu model that has been set with
// gtk_application_set_app_menu().
/*

C function : gtk_application_get_app_menu
*/
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

// Returns the menu model that has been set with
// gtk_application_set_menubar().
/*

C function : gtk_application_get_menubar
*/
func (recv *Application) GetMenubar() *gio.MenuModel {
	retC := C.gtk_application_get_menubar((*C.GtkApplication)(recv.native))
	retGo := gio.MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inform the session manager that certain types of actions should be
// inhibited. This is not guaranteed to work on all platforms and for
// all types of actions.
//
// Applications should invoke this method when they begin an operation
// that should not be interrupted, such as creating a CD or DVD. The
// types of actions that may be blocked are specified by the @flags
// parameter. When the application completes the operation it should
// call gtk_application_uninhibit() to remove the inhibitor. Note that
// an application can have multiple inhibitors, and all of them must
// be individually removed. Inhibitors are also cleared when the
// application exits.
//
// Applications should not expect that they will always be able to block
// the action. In most cases, users will be given the option to force
// the action to take place.
//
// Reasons should be short and to the point.
//
// If @window is given, the session manager may point the user to
// this window to find out more about why the action is inhibited.
/*

C function : gtk_application_inhibit
*/
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

// Determines if any of the actions specified in @flags are
// currently inhibited (possibly by another application).
//
// Note that this information may not be available (for example
// when the application is running in a sandbox).
/*

C function : gtk_application_is_inhibited
*/
func (recv *Application) IsInhibited(flags ApplicationInhibitFlags) bool {
	c_flags := (C.GtkApplicationInhibitFlags)(flags)

	retC := C.gtk_application_is_inhibited((*C.GtkApplication)(recv.native), c_flags)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_application_remove_accelerator : unsupported parameter parameter : Blacklisted record : GVariant

// Sets or unsets the application menu for @application.
//
// This can only be done in the primary instance of the application,
// after it has been registered.  #GApplication::startup is a good place
// to call this.
//
// The application menu is a single menu containing items that typically
// impact the application as a whole, rather than acting on a specific
// window or document.  For example, you would expect to see
// “Preferences” or “Quit” in an application menu, but not “Save” or
// “Print”.
//
// If supported, the application menu will be rendered by the desktop
// environment.
//
// Use the base #GActionMap interface to add actions, to respond to the user
// selecting these menu items.
/*

C function : gtk_application_set_app_menu
*/
func (recv *Application) SetAppMenu(appMenu *gio.MenuModel) {
	c_app_menu := (*C.GMenuModel)(C.NULL)
	if appMenu != nil {
		c_app_menu = (*C.GMenuModel)(appMenu.ToC())
	}

	C.gtk_application_set_app_menu((*C.GtkApplication)(recv.native), c_app_menu)

	return
}

// Sets or unsets the menubar for windows of @application.
//
// This is a menubar in the traditional sense.
//
// This can only be done in the primary instance of the application,
// after it has been registered.  #GApplication::startup is a good place
// to call this.
//
// Depending on the desktop environment, this may appear at the top of
// each window, or at the top of the screen.  In some environments, if
// both the application menu and the menubar are set, the application
// menu will be presented as if it were the first item of the menubar.
// Other environments treat the two as completely separate — for example,
// the application menu may be rendered by the desktop shell while the
// menubar (if set) remains in each individual window.
//
// Use the base #GActionMap interface to add actions, to respond to the
// user selecting these menu items.
/*

C function : gtk_application_set_menubar
*/
func (recv *Application) SetMenubar(menubar *gio.MenuModel) {
	c_menubar := (*C.GMenuModel)(C.NULL)
	if menubar != nil {
		c_menubar = (*C.GMenuModel)(menubar.ToC())
	}

	C.gtk_application_set_menubar((*C.GtkApplication)(recv.native), c_menubar)

	return
}

// Removes an inhibitor that has been established with gtk_application_inhibit().
// Inhibitors are also cleared when the application exits.
/*

C function : gtk_application_uninhibit
*/
func (recv *Application) Uninhibit(cookie uint32) {
	c_cookie := (C.guint)(cookie)

	C.gtk_application_uninhibit((*C.GtkApplication)(recv.native), c_cookie)

	return
}

// Creates a new #GtkApplicationWindow.
/*

C function : gtk_application_window_new
*/
func ApplicationWindowNew(application *Application) *ApplicationWindow {
	c_application := (*C.GtkApplication)(C.NULL)
	if application != nil {
		c_application = (*C.GtkApplication)(application.ToC())
	}

	retC := C.gtk_application_window_new(c_application)
	retGo := ApplicationWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the window will display a menubar for the app menu
// and menubar as needed.
/*

C function : gtk_application_window_get_show_menubar
*/
func (recv *ApplicationWindow) GetShowMenubar() bool {
	retC := C.gtk_application_window_get_show_menubar((*C.GtkApplicationWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the window will display a menubar for the app menu
// and menubar as needed.
/*

C function : gtk_application_window_set_show_menubar
*/
func (recv *ApplicationWindow) SetShowMenubar(showMenubar bool) {
	c_show_menubar :=
		boolToGboolean(showMenubar)

	C.gtk_application_window_set_show_menubar((*C.GtkApplicationWindow)(recv.native), c_show_menubar)

	return
}

// Parses a resource file containing a [GtkBuilder UI definition][BUILDER-UI]
// and merges it with the current contents of @builder.
//
// Most users will probably want to use gtk_builder_new_from_resource().
//
// If an error occurs, 0 will be returned and @error will be assigned a
// #GError from the #GTK_BUILDER_ERROR, #G_MARKUP_ERROR or #G_RESOURCE_ERROR
// domain.
//
// It’s not really reasonable to attempt to handle failures of this
// call.  The only reasonable thing to do when an error is detected is
// to call g_error().
/*

C function : gtk_builder_add_from_resource
*/
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

// Unsupported : gtk_builder_add_objects_from_resource : unsupported parameter object_ids :

// Creates a new #GtkColorChooserDialog.
/*

C function : gtk_color_chooser_dialog_new
*/
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

// Creates a new #GtkColorChooserWidget.
/*

C function : gtk_color_chooser_widget_new
*/
func ColorChooserWidgetNew() *ColorChooserWidget {
	retC := C.gtk_color_chooser_widget_new()
	retGo := ColorChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'format-entry-text' for ComboBox : unsupported parameter path : type utf8 :

// Computes the common prefix that is shared by all rows in @completion
// that start with @key. If no row matches @key, %NULL will be returned.
// Note that a text column must have been set for this function to work,
// see gtk_entry_completion_set_text_column() for details.
/*

C function : gtk_entry_completion_compute_prefix
*/
func (recv *EntryCompletion) ComputePrefix(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gtk_entry_completion_compute_prefix((*C.GtkEntryCompletion)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkImage displaying the resource file @resource_path. If the file
// isn’t found or can’t be loaded, the resulting #GtkImage will
// display a “broken image” icon. This function never returns %NULL,
// it always returns a valid #GtkImage widget.
//
// If the file contains an animation, the image will contain an
// animation.
//
// If you need to detect failures to load the file, use
// gdk_pixbuf_new_from_file() to load the file yourself, then create
// the #GtkImage from the pixbuf. (Or for animations, use
// gdk_pixbuf_animation_new_from_file()).
//
// The storage type (gtk_image_get_storage_type()) of the returned
// image is not defined, it will be whatever is appropriate for
// displaying the file.
/*

C function : gtk_image_new_from_resource
*/
func ImageNewFromResource(resourcePath string) *Image {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_image_new_from_resource(c_resource_path)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GtkMenu and populates it with menu items and
// submenus according to @model.
//
// The created menu items are connected to actions found in the
// #GtkApplicationWindow to which the menu belongs - typically
// by means of being attached to a widget (see gtk_menu_attach_to_widget())
// that is contained within the #GtkApplicationWindows widget hierarchy.
//
// Actions can also be added using gtk_widget_insert_action_group() on the menu's
// attach widget or on any of its parent widgets.
/*

C function : gtk_menu_new_from_model
*/
func MenuNewFromModel(model *gio.MenuModel) *Menu {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	retC := C.gtk_menu_new_from_model(c_model)
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkMenuBar and populates it with menu items
// and submenus according to @model.
//
// The created menu items are connected to actions found in the
// #GtkApplicationWindow to which the menu bar belongs - typically
// by means of being contained within the #GtkApplicationWindows
// widget hierarchy.
/*

C function : gtk_menu_bar_new_from_model
*/
func MenuBarNewFromModel(model *gio.MenuModel) *MenuBar {
	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	retC := C.gtk_menu_bar_new_from_model(c_model)
	retGo := MenuBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the scale has an origin.
/*

C function : gtk_scale_get_has_origin
*/
func (recv *Scale) GetHasOrigin() bool {
	retC := C.gtk_scale_get_has_origin((*C.GtkScale)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// If #GtkScale:has-origin is set to %TRUE (the default), the scale will
// highlight the part of the trough between the origin (bottom or left side)
// and the current value.
/*

C function : gtk_scale_set_has_origin
*/
func (recv *Scale) SetHasOrigin(hasOrigin bool) {
	c_has_origin :=
		boolToGboolean(hasOrigin)

	C.gtk_scale_set_has_origin((*C.GtkScale)(recv.native), c_has_origin)

	return
}

// Return whether button presses are captured during kinetic
// scrolling. See gtk_scrolled_window_set_capture_button_press().
/*

C function : gtk_scrolled_window_get_capture_button_press
*/
func (recv *ScrolledWindow) GetCaptureButtonPress() bool {
	retC := C.gtk_scrolled_window_get_capture_button_press((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the specified kinetic scrolling behavior.
/*

C function : gtk_scrolled_window_get_kinetic_scrolling
*/
func (recv *ScrolledWindow) GetKineticScrolling() bool {
	retC := C.gtk_scrolled_window_get_kinetic_scrolling((*C.GtkScrolledWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Changes the behaviour of @scrolled_window with regard to the initial
// event that possibly starts kinetic scrolling. When @capture_button_press
// is set to %TRUE, the event is captured by the scrolled window, and
// then later replayed if it is meant to go to the child widget.
//
// This should be enabled if any child widgets perform non-reversible
// actions on #GtkWidget::button-press-event. If they don't, and handle
// additionally handle #GtkWidget::grab-broken-event, it might be better
// to set @capture_button_press to %FALSE.
//
// This setting only has an effect if kinetic scrolling is enabled.
/*

C function : gtk_scrolled_window_set_capture_button_press
*/
func (recv *ScrolledWindow) SetCaptureButtonPress(captureButtonPress bool) {
	c_capture_button_press :=
		boolToGboolean(captureButtonPress)

	C.gtk_scrolled_window_set_capture_button_press((*C.GtkScrolledWindow)(recv.native), c_capture_button_press)

	return
}

// Turns kinetic scrolling on or off.
// Kinetic scrolling only applies to devices with source
// %GDK_SOURCE_TOUCHSCREEN.
/*

C function : gtk_scrolled_window_set_kinetic_scrolling
*/
func (recv *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	c_kinetic_scrolling :=
		boolToGboolean(kineticScrolling)

	C.gtk_scrolled_window_set_kinetic_scrolling((*C.GtkScrolledWindow)(recv.native), c_kinetic_scrolling)

	return
}

// Gets the parent context set via gtk_style_context_set_parent().
// See that function for details.
/*

C function : gtk_style_context_get_parent
*/
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

// Sets the parent style context for @context. The parent style
// context is used to implement
// [inheritance](http://www.w3.org/TR/css3-cascade/#inheritance)
// of properties.
//
// If you are using a #GtkStyleContext returned from
// gtk_widget_get_style_context(), the parent will be set for you.
/*

C function : gtk_style_context_set_parent
*/
func (recv *StyleContext) SetParent(parent *StyleContext) {
	c_parent := (*C.GtkStyleContext)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkStyleContext)(parent.ToC())
	}

	C.gtk_style_context_set_parent((*C.GtkStyleContext)(recv.native), c_parent)

	return
}

// Queries the number of columns in the given @tree_view.
/*

C function : gtk_tree_view_get_n_columns
*/
func (recv *TreeView) GetNColumns() uint32 {
	retC := C.gtk_tree_view_get_n_columns((*C.GtkTreeView)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Parses a resource file containing a [UI definition][XML-UI] and
// merges it with the current contents of @manager.
/*

C function : gtk_ui_manager_add_ui_from_resource
*/
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

// Returns the modifier mask the @widget’s windowing system backend
// uses for a particular purpose.
//
// See gdk_keymap_get_modifier_mask().
/*

C function : gtk_widget_get_modifier_mask
*/
func (recv *Widget) GetModifierMask(intent gdk.ModifierIntent) gdk.ModifierType {
	c_intent := (C.GdkModifierIntent)(intent)

	retC := C.gtk_widget_get_modifier_mask((*C.GtkWidget)(recv.native), c_intent)
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// Fetches the attach widget for this window. See
// gtk_window_set_attached_to().
/*

C function : gtk_window_get_attached_to
*/
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

// Returns whether the window has requested to have its titlebar hidden
// when maximized. See gtk_window_set_hide_titlebar_when_maximized ().
/*

C function : gtk_window_get_hide_titlebar_when_maximized
*/
func (recv *Window) GetHideTitlebarWhenMaximized() bool {
	retC := C.gtk_window_get_hide_titlebar_when_maximized((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Marks @window as attached to @attach_widget. This creates a logical binding
// between the window and the widget it belongs to, which is used by GTK+ to
// propagate information such as styling or accessibility to @window as if it
// was a children of @attach_widget.
//
// Examples of places where specifying this relation is useful are for instance
// a #GtkMenu created by a #GtkComboBox, a completion popup window
// created by #GtkEntry or a typeahead search entry created by #GtkTreeView.
//
// Note that this function should not be confused with
// gtk_window_set_transient_for(), which specifies a window manager relation
// between two toplevels instead.
//
// Passing %NULL for @attach_widget detaches the window.
/*

C function : gtk_window_set_attached_to
*/
func (recv *Window) SetAttachedTo(attachWidget *Widget) {
	c_attach_widget := (*C.GtkWidget)(C.NULL)
	if attachWidget != nil {
		c_attach_widget = (*C.GtkWidget)(attachWidget.ToC())
	}

	C.gtk_window_set_attached_to((*C.GtkWindow)(recv.native), c_attach_widget)

	return
}

// If @setting is %TRUE, then @window will request that it’s titlebar
// should be hidden when maximized.
// This is useful for windows that don’t convey any information other
// than the application name in the titlebar, to put the available
// screen space to better use. If the underlying window system does not
// support the request, the setting will not have any effect.
//
// Note that custom titlebars set with gtk_window_set_titlebar() are
// not affected by this. The application is in full control of their
// content and visibility anyway.
/*

C function : gtk_window_set_hide_titlebar_when_maximized
*/
func (recv *Window) SetHideTitlebarWhenMaximized(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_hide_titlebar_when_maximized((*C.GtkWindow)(recv.native), c_setting)

	return
}

// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void application_windowAddedHandler(GObject *, GtkWindow *, gpointer);

	static gulong Application_signal_connect_window_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-added", G_CALLBACK(application_windowAddedHandler), data);
	}

*/
/*

	void application_windowRemovedHandler(GObject *, GtkWindow *, gpointer);

	static gulong Application_signal_connect_window_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-removed", G_CALLBACK(application_windowRemovedHandler), data);
	}

*/
import "C"

// Gets the smaller of step increment and page increment.
/*

C function : gtk_adjustment_get_minimum_increment
*/
func (recv *Adjustment) GetMinimumIncrement() float64 {
	retC := C.gtk_adjustment_get_minimum_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Returns the current value of the #GtkAppChooserButton:show-default-item
// property.
/*

C function : gtk_app_chooser_button_get_show_default_item
*/
func (recv *AppChooserButton) GetShowDefaultItem() bool {
	retC := C.gtk_app_chooser_button_get_show_default_item((*C.GtkAppChooserButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the dropdown menu of this button should show the
// default application for the given content type at top.
/*

C function : gtk_app_chooser_button_set_show_default_item
*/
func (recv *AppChooserButton) SetShowDefaultItem(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_button_set_show_default_item((*C.GtkAppChooserButton)(recv.native), c_setting)

	return
}

type signalApplicationWindowAddedDetail struct {
	callback  ApplicationSignalWindowAddedCallback
	handlerID C.gulong
}

var signalApplicationWindowAddedId int
var signalApplicationWindowAddedMap = make(map[int]signalApplicationWindowAddedDetail)
var signalApplicationWindowAddedLock sync.Mutex

// ApplicationSignalWindowAddedCallback is a callback function for a 'window-added' signal emitted from a Application.
type ApplicationSignalWindowAddedCallback func(window *Window)

/*
ConnectWindowAdded connects the callback to the 'window-added' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectWindowAdded to remove it.
*/
func (recv *Application) ConnectWindowAdded(callback ApplicationSignalWindowAddedCallback) int {
	signalApplicationWindowAddedLock.Lock()
	defer signalApplicationWindowAddedLock.Unlock()

	signalApplicationWindowAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_window_added(instance, C.gpointer(uintptr(signalApplicationWindowAddedId)))

	detail := signalApplicationWindowAddedDetail{callback, handlerID}
	signalApplicationWindowAddedMap[signalApplicationWindowAddedId] = detail

	return signalApplicationWindowAddedId
}

/*
DisconnectWindowAdded disconnects a callback from the 'window-added' signal for the Application.

The connectionID should be a value returned from a call to ConnectWindowAdded.
*/
func (recv *Application) DisconnectWindowAdded(connectionID int) {
	signalApplicationWindowAddedLock.Lock()
	defer signalApplicationWindowAddedLock.Unlock()

	detail, exists := signalApplicationWindowAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationWindowAddedMap, connectionID)
}

//export application_windowAddedHandler
func application_windowAddedHandler(_ *C.GObject, c_window *C.GtkWindow, data C.gpointer) {
	window := WindowNewFromC(unsafe.Pointer(c_window))

	index := int(uintptr(data))
	callback := signalApplicationWindowAddedMap[index].callback
	callback(window)
}

type signalApplicationWindowRemovedDetail struct {
	callback  ApplicationSignalWindowRemovedCallback
	handlerID C.gulong
}

var signalApplicationWindowRemovedId int
var signalApplicationWindowRemovedMap = make(map[int]signalApplicationWindowRemovedDetail)
var signalApplicationWindowRemovedLock sync.Mutex

// ApplicationSignalWindowRemovedCallback is a callback function for a 'window-removed' signal emitted from a Application.
type ApplicationSignalWindowRemovedCallback func(window *Window)

/*
ConnectWindowRemoved connects the callback to the 'window-removed' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectWindowRemoved to remove it.
*/
func (recv *Application) ConnectWindowRemoved(callback ApplicationSignalWindowRemovedCallback) int {
	signalApplicationWindowRemovedLock.Lock()
	defer signalApplicationWindowRemovedLock.Unlock()

	signalApplicationWindowRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_window_removed(instance, C.gpointer(uintptr(signalApplicationWindowRemovedId)))

	detail := signalApplicationWindowRemovedDetail{callback, handlerID}
	signalApplicationWindowRemovedMap[signalApplicationWindowRemovedId] = detail

	return signalApplicationWindowRemovedId
}

/*
DisconnectWindowRemoved disconnects a callback from the 'window-removed' signal for the Application.

The connectionID should be a value returned from a call to ConnectWindowRemoved.
*/
func (recv *Application) DisconnectWindowRemoved(connectionID int) {
	signalApplicationWindowRemovedLock.Lock()
	defer signalApplicationWindowRemovedLock.Unlock()

	detail, exists := signalApplicationWindowRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationWindowRemovedMap, connectionID)
}

//export application_windowRemovedHandler
func application_windowRemovedHandler(_ *C.GObject, c_window *C.GtkWindow, data C.gpointer) {
	window := WindowNewFromC(unsafe.Pointer(c_window))

	index := int(uintptr(data))
	callback := signalApplicationWindowRemovedMap[index].callback
	callback(window)
}

// Removes the @page_num’s page from @assistant.
/*

C function : gtk_assistant_remove_page
*/
func (recv *Assistant) RemovePage(pageNum int32) {
	c_page_num := (C.gint)(pageNum)

	C.gtk_assistant_remove_page((*C.GtkAssistant)(recv.native), c_page_num)

	return
}

// Returns whether the child is exempted from homogenous
// sizing.
/*

C function : gtk_button_box_get_child_non_homogeneous
*/
func (recv *ButtonBox) GetChildNonHomogeneous(child *Widget) bool {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	retC := C.gtk_button_box_get_child_non_homogeneous((*C.GtkButtonBox)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the child is exempted from homogeous sizing.
/*

C function : gtk_button_box_set_child_non_homogeneous
*/
func (recv *ButtonBox) SetChildNonHomogeneous(child *Widget, nonHomogeneous bool) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_non_homogeneous :=
		boolToGboolean(nonHomogeneous)

	C.gtk_button_box_set_child_non_homogeneous((*C.GtkButtonBox)(recv.native), c_child, c_non_homogeneous)

	return
}

// Emits a #GtkWidget::child-notify signal for the
// [child property][child-properties]
// @child_property on the child.
//
// This is an analogue of g_object_notify() for child properties.
//
// Also see gtk_widget_child_notify().
/*

C function : gtk_container_child_notify
*/
func (recv *Container) ChildNotify(child *Widget, childProperty string) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_child_property := C.CString(childProperty)
	defer C.free(unsafe.Pointer(c_child_property))

	C.gtk_container_child_notify((*C.GtkContainer)(recv.native), c_child, c_child_property)

	return
}

// Converts the @provider into a string representation in CSS
// format.
//
// Using gtk_css_provider_load_from_data() with the return value
// from this function on a new provider created with
// gtk_css_provider_new() will basically create a duplicate of
// this @provider.
/*

C function : gtk_css_provider_to_string
*/
func (recv *CssProvider) ToString() string {
	retC := C.gtk_css_provider_to_string((*C.GtkCssProvider)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Retrieves the text that will be displayed when @entry is empty and unfocused
/*

C function : gtk_entry_get_placeholder_text
*/
func (recv *Entry) GetPlaceholderText() string {
	retC := C.gtk_entry_get_placeholder_text((*C.GtkEntry)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets text to be displayed in @entry when it is empty and unfocused.
// This can be used to give a visual hint of the expected contents of
// the #GtkEntry.
//
// Note that since the placeholder text gets removed when the entry
// received focus, using this feature is a bit problematic if the entry
// is given the initial focus in a window. Sometimes this can be
// worked around by delaying the initial focus setting until the
// first key event arrives.
/*

C function : gtk_entry_set_placeholder_text
*/
func (recv *Entry) SetPlaceholderText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_set_placeholder_text((*C.GtkEntry)(recv.native), c_text)

	return
}

// Returns whether the expander will resize the toplevel widget
// containing the expander upon resizing and collpasing.
/*

C function : gtk_expander_get_resize_toplevel
*/
func (recv *Expander) GetResizeToplevel() bool {
	retC := C.gtk_expander_get_resize_toplevel((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the expander will resize the toplevel widget
// containing the expander upon resizing and collpasing.
/*

C function : gtk_expander_set_resize_toplevel
*/
func (recv *Expander) SetResizeToplevel(resizeToplevel bool) {
	c_resize_toplevel :=
		boolToGboolean(resizeToplevel)

	C.gtk_expander_set_resize_toplevel((*C.GtkExpander)(recv.native), c_resize_toplevel)

	return
}

// Creates a new #GtkFontChooserDialog.
/*

C function : gtk_font_chooser_dialog_new
*/
func FontChooserDialogNew(title string, parent *Window) *FontChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_font_chooser_dialog_new(c_title, c_parent)
	retGo := FontChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkFontChooserWidget.
/*

C function : gtk_font_chooser_widget_new
*/
func FontChooserWidgetNew() *FontChooserWidget {
	retC := C.gtk_font_chooser_widget_new()
	retGo := FontChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the child of @grid whose area covers the grid
// cell whose upper left corner is at @left, @top.
/*

C function : gtk_grid_get_child_at
*/
func (recv *Grid) GetChildAt(left int32, top int32) *Widget {
	c_left := (C.gint)(left)

	c_top := (C.gint)(top)

	retC := C.gtk_grid_get_child_at((*C.GtkGrid)(recv.native), c_left, c_top)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Inserts a column at the specified position.
//
// Children which are attached at or to the right of this position
// are moved one column to the right. Children which span across this
// position are grown to span the new column.
/*

C function : gtk_grid_insert_column
*/
func (recv *Grid) InsertColumn(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_insert_column((*C.GtkGrid)(recv.native), c_position)

	return
}

// Inserts a row or column at the specified position.
//
// The new row or column is placed next to @sibling, on the side
// determined by @side. If @side is %GTK_POS_TOP or %GTK_POS_BOTTOM,
// a row is inserted. If @side is %GTK_POS_LEFT of %GTK_POS_RIGHT,
// a column is inserted.
/*

C function : gtk_grid_insert_next_to
*/
func (recv *Grid) InsertNextTo(sibling *Widget, side PositionType) {
	c_sibling := (*C.GtkWidget)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkWidget)(sibling.ToC())
	}

	c_side := (C.GtkPositionType)(side)

	C.gtk_grid_insert_next_to((*C.GtkGrid)(recv.native), c_sibling, c_side)

	return
}

// Inserts a row at the specified position.
//
// Children which are attached at or below this position
// are moved one row down. Children which span across this
// position are grown to span the new row.
/*

C function : gtk_grid_insert_row
*/
func (recv *Grid) InsertRow(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_insert_row((*C.GtkGrid)(recv.native), c_position)

	return
}

// Creates a new lock button which reflects the @permission.
/*

C function : gtk_lock_button_new
*/
func LockButtonNew(permission *gio.Permission) *LockButton {
	c_permission := (*C.GPermission)(C.NULL)
	if permission != nil {
		c_permission = (*C.GPermission)(permission.ToC())
	}

	retC := C.gtk_lock_button_new(c_permission)
	retGo := LockButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains the #GPermission object that controls @button.
/*

C function : gtk_lock_button_get_permission
*/
func (recv *LockButton) GetPermission() *gio.Permission {
	retC := C.gtk_lock_button_get_permission((*C.GtkLockButton)(recv.native))
	retGo := gio.PermissionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the #GPermission object that controls @button.
/*

C function : gtk_lock_button_set_permission
*/
func (recv *LockButton) SetPermission(permission *gio.Permission) {
	c_permission := (*C.GPermission)(C.NULL)
	if permission != nil {
		c_permission = (*C.GPermission)(permission.ToC())
	}

	C.gtk_lock_button_set_permission((*C.GtkLockButton)(recv.native), c_permission)

	return
}

// Unsupported signal 'insert' for MenuShell : unsupported parameter position : type gint :

// Creates a new #GtkOverlay.
/*

C function : gtk_overlay_new
*/
func OverlayNew() *Overlay {
	retC := C.gtk_overlay_new()
	retGo := OverlayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds @widget to @overlay.
//
// The widget will be stacked on top of the main widget
// added with gtk_container_add().
//
// The position at which @widget is placed is determined
// from its #GtkWidget:halign and #GtkWidget:valign properties.
/*

C function : gtk_overlay_add_overlay
*/
func (recv *Overlay) AddOverlay(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_overlay_add_overlay((*C.GtkOverlay)(recv.native), c_widget)

	return
}

// Returns the current X offset of @tree_column in pixels.
/*

C function : gtk_tree_view_column_get_x_offset
*/
func (recv *TreeViewColumn) GetXOffset() int32 {
	retC := C.gtk_tree_view_column_get_x_offset((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets the icon that will be used for drags from a particular source
// to @icon. See the docs for #GtkIconTheme for more details.
/*

C function : gtk_drag_source_set_icon_gicon
*/
func (recv *Widget) DragSourceSetIconGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_drag_source_set_icon_gicon((*C.GtkWidget)(recv.native), c_icon)

	return
}

// Determines if the widget should show a visible indication that
// it has the global input focus. This is a convenience function for
// use in ::draw handlers that takes into account whether focus
// indication should currently be shown in the toplevel window of
// @widget. See gtk_window_get_focus_visible() for more information
// about focus indication.
//
// To find out if the widget has the global input focus, use
// gtk_widget_has_focus().
/*

C function : gtk_widget_has_visible_focus
*/
func (recv *Widget) HasVisibleFocus() bool {
	retC := C.gtk_widget_has_visible_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value of the #GtkWindow:focus-visible property.
/*

C function : gtk_window_get_focus_visible
*/
func (recv *Window) GetFocusVisible() bool {
	retC := C.gtk_window_get_focus_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the #GtkWindow:focus-visible property.
/*

C function : gtk_window_set_focus_visible
*/
func (recv *Window) SetFocusVisible(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_focus_visible((*C.GtkWindow)(recv.native), c_setting)

	return
}

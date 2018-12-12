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
/*

	void menushell_insertHandler(GObject *, GtkWidget *, gint, gpointer);

	static gulong MenuShell_signal_connect_insert(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert", G_CALLBACK(menushell_insertHandler), data);
	}

*/
import "C"

// GetMinimumIncrement is a wrapper around the C function gtk_adjustment_get_minimum_increment.
func (recv *Adjustment) GetMinimumIncrement() float64 {
	retC := C.gtk_adjustment_get_minimum_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetShowDefaultItem is a wrapper around the C function gtk_app_chooser_button_get_show_default_item.
func (recv *AppChooserButton) GetShowDefaultItem() bool {
	retC := C.gtk_app_chooser_button_get_show_default_item((*C.GtkAppChooserButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowDefaultItem is a wrapper around the C function gtk_app_chooser_button_set_show_default_item.
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
var signalApplicationWindowAddedLock sync.RWMutex

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
	signalApplicationWindowAddedLock.RLock()
	defer signalApplicationWindowAddedLock.RUnlock()

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
var signalApplicationWindowRemovedLock sync.RWMutex

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
	signalApplicationWindowRemovedLock.RLock()
	defer signalApplicationWindowRemovedLock.RUnlock()

	window := WindowNewFromC(unsafe.Pointer(c_window))

	index := int(uintptr(data))
	callback := signalApplicationWindowRemovedMap[index].callback
	callback(window)
}

// RemovePage is a wrapper around the C function gtk_assistant_remove_page.
func (recv *Assistant) RemovePage(pageNum int32) {
	c_page_num := (C.gint)(pageNum)

	C.gtk_assistant_remove_page((*C.GtkAssistant)(recv.native), c_page_num)

	return
}

// GetChildNonHomogeneous is a wrapper around the C function gtk_button_box_get_child_non_homogeneous.
func (recv *ButtonBox) GetChildNonHomogeneous(child *Widget) bool {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	retC := C.gtk_button_box_get_child_non_homogeneous((*C.GtkButtonBox)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// SetChildNonHomogeneous is a wrapper around the C function gtk_button_box_set_child_non_homogeneous.
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

// ChildNotify is a wrapper around the C function gtk_container_child_notify.
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

// ToString is a wrapper around the C function gtk_css_provider_to_string.
func (recv *CssProvider) ToString() string {
	retC := C.gtk_css_provider_to_string((*C.GtkCssProvider)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPlaceholderText is a wrapper around the C function gtk_entry_get_placeholder_text.
func (recv *Entry) GetPlaceholderText() string {
	retC := C.gtk_entry_get_placeholder_text((*C.GtkEntry)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetPlaceholderText is a wrapper around the C function gtk_entry_set_placeholder_text.
func (recv *Entry) SetPlaceholderText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_set_placeholder_text((*C.GtkEntry)(recv.native), c_text)

	return
}

// GetResizeToplevel is a wrapper around the C function gtk_expander_get_resize_toplevel.
func (recv *Expander) GetResizeToplevel() bool {
	retC := C.gtk_expander_get_resize_toplevel((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetResizeToplevel is a wrapper around the C function gtk_expander_set_resize_toplevel.
func (recv *Expander) SetResizeToplevel(resizeToplevel bool) {
	c_resize_toplevel :=
		boolToGboolean(resizeToplevel)

	C.gtk_expander_set_resize_toplevel((*C.GtkExpander)(recv.native), c_resize_toplevel)

	return
}

// FontChooserDialogNew is a wrapper around the C function gtk_font_chooser_dialog_new.
func FontChooserDialogNew(title string, parent *Window) *FontChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_font_chooser_dialog_new(c_title, c_parent)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := FontChooserDialogNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	return retGo
}

// FontChooserWidgetNew is a wrapper around the C function gtk_font_chooser_widget_new.
func FontChooserWidgetNew() *FontChooserWidget {
	retC := C.gtk_font_chooser_widget_new()
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := FontChooserWidgetNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	return retGo
}

// GetChildAt is a wrapper around the C function gtk_grid_get_child_at.
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

// InsertColumn is a wrapper around the C function gtk_grid_insert_column.
func (recv *Grid) InsertColumn(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_insert_column((*C.GtkGrid)(recv.native), c_position)

	return
}

// InsertNextTo is a wrapper around the C function gtk_grid_insert_next_to.
func (recv *Grid) InsertNextTo(sibling *Widget, side PositionType) {
	c_sibling := (*C.GtkWidget)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkWidget)(sibling.ToC())
	}

	c_side := (C.GtkPositionType)(side)

	C.gtk_grid_insert_next_to((*C.GtkGrid)(recv.native), c_sibling, c_side)

	return
}

// InsertRow is a wrapper around the C function gtk_grid_insert_row.
func (recv *Grid) InsertRow(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_insert_row((*C.GtkGrid)(recv.native), c_position)

	return
}

// LockButtonNew is a wrapper around the C function gtk_lock_button_new.
func LockButtonNew(permission *gio.Permission) *LockButton {
	c_permission := (*C.GPermission)(C.NULL)
	if permission != nil {
		c_permission = (*C.GPermission)(permission.ToC())
	}

	retC := C.gtk_lock_button_new(c_permission)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := LockButtonNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	return retGo
}

// GetPermission is a wrapper around the C function gtk_lock_button_get_permission.
func (recv *LockButton) GetPermission() *gio.Permission {
	retC := C.gtk_lock_button_get_permission((*C.GtkLockButton)(recv.native))
	retGo := gio.PermissionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetPermission is a wrapper around the C function gtk_lock_button_set_permission.
func (recv *LockButton) SetPermission(permission *gio.Permission) {
	c_permission := (*C.GPermission)(C.NULL)
	if permission != nil {
		c_permission = (*C.GPermission)(permission.ToC())
	}

	C.gtk_lock_button_set_permission((*C.GtkLockButton)(recv.native), c_permission)

	return
}

type signalMenuShellInsertDetail struct {
	callback  MenuShellSignalInsertCallback
	handlerID C.gulong
}

var signalMenuShellInsertId int
var signalMenuShellInsertMap = make(map[int]signalMenuShellInsertDetail)
var signalMenuShellInsertLock sync.RWMutex

// MenuShellSignalInsertCallback is a callback function for a 'insert' signal emitted from a MenuShell.
type MenuShellSignalInsertCallback func(child *Widget, position int32)

/*
ConnectInsert connects the callback to the 'insert' signal for the MenuShell.

The returned value represents the connection, and may be passed to DisconnectInsert to remove it.
*/
func (recv *MenuShell) ConnectInsert(callback MenuShellSignalInsertCallback) int {
	signalMenuShellInsertLock.Lock()
	defer signalMenuShellInsertLock.Unlock()

	signalMenuShellInsertId++
	instance := C.gpointer(recv.native)
	handlerID := C.MenuShell_signal_connect_insert(instance, C.gpointer(uintptr(signalMenuShellInsertId)))

	detail := signalMenuShellInsertDetail{callback, handlerID}
	signalMenuShellInsertMap[signalMenuShellInsertId] = detail

	return signalMenuShellInsertId
}

/*
DisconnectInsert disconnects a callback from the 'insert' signal for the MenuShell.

The connectionID should be a value returned from a call to ConnectInsert.
*/
func (recv *MenuShell) DisconnectInsert(connectionID int) {
	signalMenuShellInsertLock.Lock()
	defer signalMenuShellInsertLock.Unlock()

	detail, exists := signalMenuShellInsertMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMenuShellInsertMap, connectionID)
}

//export menushell_insertHandler
func menushell_insertHandler(_ *C.GObject, c_child *C.GtkWidget, c_position C.gint, data C.gpointer) {
	signalMenuShellInsertLock.RLock()
	defer signalMenuShellInsertLock.RUnlock()

	child := WidgetNewFromC(unsafe.Pointer(c_child))

	position := int32(c_position)

	index := int(uintptr(data))
	callback := signalMenuShellInsertMap[index].callback
	callback(child, position)
}

// OverlayNew is a wrapper around the C function gtk_overlay_new.
func OverlayNew() *Overlay {
	retC := C.gtk_overlay_new()
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := OverlayNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	return retGo
}

// AddOverlay is a wrapper around the C function gtk_overlay_add_overlay.
func (recv *Overlay) AddOverlay(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_overlay_add_overlay((*C.GtkOverlay)(recv.native), c_widget)

	return
}

// GetXOffset is a wrapper around the C function gtk_tree_view_column_get_x_offset.
func (recv *TreeViewColumn) GetXOffset() int32 {
	retC := C.gtk_tree_view_column_get_x_offset((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// DragSourceSetIconGicon is a wrapper around the C function gtk_drag_source_set_icon_gicon.
func (recv *Widget) DragSourceSetIconGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_drag_source_set_icon_gicon((*C.GtkWidget)(recv.native), c_icon)

	return
}

// HasVisibleFocus is a wrapper around the C function gtk_widget_has_visible_focus.
func (recv *Widget) HasVisibleFocus() bool {
	retC := C.gtk_widget_has_visible_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFocusVisible is a wrapper around the C function gtk_window_get_focus_visible.
func (recv *Window) GetFocusVisible() bool {
	retC := C.gtk_window_get_focus_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFocusVisible is a wrapper around the C function gtk_window_set_focus_visible.
func (recv *Window) SetFocusVisible(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_focus_visible((*C.GtkWindow)(recv.native), c_setting)

	return
}

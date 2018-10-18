// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
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

	void Application_windowAddedHandler();

	static gulong Application_signal_connect_window_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-added", Application_windowAddedHandler, data);
	}

*/
/*

	void Application_windowRemovedHandler();

	static gulong Application_signal_connect_window_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-removed", Application_windowRemovedHandler, data);
	}

*/
/*

	void MenuShell_insertHandler();

	static gulong MenuShell_signal_connect_insert(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert", MenuShell_insertHandler, data);
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

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal 'application-activated' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'application-selected' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'populate-popup' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

var signalApplicationWindowAddedId int
var signalApplicationWindowAddedMap = make(map[int]ApplicationSignalWindowAddedCallback)
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
	signalApplicationWindowAddedMap[signalApplicationWindowAddedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Application_signal_connect_window_added(instance, C.gpointer(uintptr(signalApplicationWindowAddedId)))
	return int(retC)
}

/*
DisconnectWindowAdded disconnects a callback from the 'window-added' signal for the Application.

The connectionID should be a value returned from a call to ConnectWindowAdded.
*/
func (recv *Application) DisconnectWindowAdded(connectionID int) {
	_, exists := signalApplicationWindowAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalApplicationWindowAddedMap, connectionID)
}

//export Application_windowAddedHandler
func Application_windowAddedHandler() {
	fmt.Println("cb")
}

var signalApplicationWindowRemovedId int
var signalApplicationWindowRemovedMap = make(map[int]ApplicationSignalWindowRemovedCallback)
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
	signalApplicationWindowRemovedMap[signalApplicationWindowRemovedId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Application_signal_connect_window_removed(instance, C.gpointer(uintptr(signalApplicationWindowRemovedId)))
	return int(retC)
}

/*
DisconnectWindowRemoved disconnects a callback from the 'window-removed' signal for the Application.

The connectionID should be a value returned from a call to ConnectWindowRemoved.
*/
func (recv *Application) DisconnectWindowRemoved(connectionID int) {
	_, exists := signalApplicationWindowRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalApplicationWindowRemovedMap, connectionID)
}

//export Application_windowRemovedHandler
func Application_windowRemovedHandler() {
	fmt.Println("cb")
}

// RemovePage is a wrapper around the C function gtk_assistant_remove_page.
func (recv *Assistant) RemovePage(pageNum int32) {
	c_page_num := (C.gint)(pageNum)

	C.gtk_assistant_remove_page((*C.GtkAssistant)(recv.native), c_page_num)

	return
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// GetChildNonHomogeneous is a wrapper around the C function gtk_button_box_get_child_non_homogeneous.
func (recv *ButtonBox) GetChildNonHomogeneous(child *Widget) bool {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_button_box_get_child_non_homogeneous((*C.GtkButtonBox)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// SetChildNonHomogeneous is a wrapper around the C function gtk_button_box_set_child_non_homogeneous.
func (recv *ButtonBox) SetChildNonHomogeneous(child *Widget, nonHomogeneous bool) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_non_homogeneous :=
		boolToGboolean(nonHomogeneous)

	C.gtk_button_box_set_child_non_homogeneous((*C.GtkButtonBox)(recv.native), c_child, c_non_homogeneous)

	return
}

// Unsupported signal 'add-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'apply-attributes' for CellArea : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'remove-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'editing-started' for CellRenderer : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'format-entry-text' for ComboBox : return value utf8 :

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// ChildNotify is a wrapper around the C function gtk_container_child_notify.
func (recv *Container) ChildNotify(child *Widget, childProperty string) {
	c_child := (*C.GtkWidget)(child.ToC())

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

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

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

// Unsupported signal 'cursor-on-match' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'match-selected' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported : EntryIconAccessible : no CType

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

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// FontChooserDialogNew is a wrapper around the C function gtk_font_chooser_dialog_new.
func FontChooserDialogNew(title string, parent *Window) *FontChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(parent.ToC())

	retC := C.gtk_font_chooser_dialog_new(c_title, c_parent)
	retGo := FontChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontChooserWidgetNew is a wrapper around the C function gtk_font_chooser_widget_new.
func FontChooserWidgetNew() *FontChooserWidget {
	retC := C.gtk_font_chooser_widget_new()
	retGo := FontChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChildAt is a wrapper around the C function gtk_grid_get_child_at.
func (recv *Grid) GetChildAt(left int32, top int32) *Widget {
	c_left := (C.gint)(left)

	c_top := (C.gint)(top)

	retC := C.gtk_grid_get_child_at((*C.GtkGrid)(recv.native), c_left, c_top)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

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
	c_sibling := (*C.GtkWidget)(sibling.ToC())

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

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// LockButtonNew is a wrapper around the C function gtk_lock_button_new.
func LockButtonNew(permission *gio.Permission) *LockButton {
	c_permission := (*C.GPermission)(permission.ToC())

	retC := C.gtk_lock_button_new(c_permission)
	retGo := LockButtonNewFromC(unsafe.Pointer(retC))

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
	c_permission := (*C.GPermission)(permission.ToC())

	C.gtk_lock_button_set_permission((*C.GtkLockButton)(recv.native), c_permission)

	return
}

var signalMenuShellInsertId int
var signalMenuShellInsertMap = make(map[int]MenuShellSignalInsertCallback)
var signalMenuShellInsertLock sync.Mutex

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
	signalMenuShellInsertMap[signalMenuShellInsertId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.MenuShell_signal_connect_insert(instance, C.gpointer(uintptr(signalMenuShellInsertId)))
	return int(retC)
}

/*
DisconnectInsert disconnects a callback from the 'insert' signal for the MenuShell.

The connectionID should be a value returned from a call to ConnectInsert.
*/
func (recv *MenuShell) DisconnectInsert(connectionID int) {
	_, exists := signalMenuShellInsertMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalMenuShellInsertMap, connectionID)
}

//export MenuShell_insertHandler
func MenuShell_insertHandler() {
	fmt.Println("cb")
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Unsupported signal 'create-window' for Notebook : return value Notebook :

// Unsupported signal 'get-child-position' for Overlay : unsupported parameter allocation : Blacklisted record : GdkRectangle

// OverlayNew is a wrapper around the C function gtk_overlay_new.
func OverlayNew() *Overlay {
	retC := C.gtk_overlay_new()
	retGo := OverlayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddOverlay is a wrapper around the C function gtk_overlay_add_overlay.
func (recv *Overlay) AddOverlay(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_overlay_add_overlay((*C.GtkOverlay)(recv.native), c_widget)

	return
}

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

// Unsupported signal 'input' for SpinButton : return value gint :

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported signal 'event' for TextTag : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// GetXOffset is a wrapper around the C function gtk_tree_view_column_get_x_offset.
func (recv *TreeViewColumn) GetXOffset() int32 {
	retC := C.gtk_tree_view_column_get_x_offset((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported signal 'child-notify' for Widget : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal 'delete-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'destroy-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event-after' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'size-allocate' for Widget : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal 'touch-event' for Widget : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_drag_source_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

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

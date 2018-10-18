// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void Widget_grabBrokenEventHandler();

	static gulong Widget_signal_connect_grab_broken_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "grab-broken-event", Widget_grabBrokenEventHandler, data);
	}

*/
import "C"

// GetWrapLicense is a wrapper around the C function gtk_about_dialog_get_wrap_license.
func (recv *AboutDialog) GetWrapLicense() bool {
	retC := C.gtk_about_dialog_get_wrap_license((*C.GtkAboutDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetWrapLicense is a wrapper around the C function gtk_about_dialog_set_wrap_license.
func (recv *AboutDialog) SetWrapLicense(wrapLicense bool) {
	c_wrap_license :=
		boolToGboolean(wrapLicense)

	C.gtk_about_dialog_set_wrap_license((*C.GtkAboutDialog)(recv.native), c_wrap_license)

	return
}

// GetAccelClosure is a wrapper around the C function gtk_action_get_accel_closure.
func (recv *Action) GetAccelClosure() *gobject.Closure {
	retC := C.gtk_action_get_accel_closure((*C.GtkAction)(recv.native))
	retGo := gobject.ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported signal 'application-activated' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'application-selected' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal 'populate-popup' for AppChooserWidget : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported signal 'add-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'apply-attributes' for CellArea : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'remove-editable' for CellArea : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'editing-started' for CellRenderer : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal 'format-entry-text' for ComboBox : return value utf8 :

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetResponseForWidget is a wrapper around the C function gtk_dialog_get_response_for_widget.
func (recv *Dialog) GetResponseForWidget(widget *Widget) int32 {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_dialog_get_response_for_widget((*C.GtkDialog)(recv.native), c_widget)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported signal 'cursor-on-match' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal 'match-selected' for EntryCompletion : unsupported parameter model : no type generator for TreeModel,

// GetPopupSetWidth is a wrapper around the C function gtk_entry_completion_get_popup_set_width.
func (recv *EntryCompletion) GetPopupSetWidth() bool {
	retC := C.gtk_entry_completion_get_popup_set_width((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPopupSingleMatch is a wrapper around the C function gtk_entry_completion_get_popup_single_match.
func (recv *EntryCompletion) GetPopupSingleMatch() bool {
	retC := C.gtk_entry_completion_get_popup_single_match((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPopupSetWidth is a wrapper around the C function gtk_entry_completion_set_popup_set_width.
func (recv *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	c_popup_set_width :=
		boolToGboolean(popupSetWidth)

	C.gtk_entry_completion_set_popup_set_width((*C.GtkEntryCompletion)(recv.native), c_popup_set_width)

	return
}

// SetPopupSingleMatch is a wrapper around the C function gtk_entry_completion_set_popup_single_match.
func (recv *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	c_popup_single_match :=
		boolToGboolean(popupSingleMatch)

	C.gtk_entry_completion_set_popup_single_match((*C.GtkEntryCompletion)(recv.native), c_popup_single_match)

	return
}

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// CreateDragIcon is a wrapper around the C function gtk_icon_view_create_drag_icon.
func (recv *IconView) CreateDragIcon(path *TreePath) *cairo.Surface {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_icon_view_create_drag_icon((*C.GtkIconView)(recv.native), c_path)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_enable_model_drag_source : unsupported parameter targets : no param type

// Unsupported : gtk_icon_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_dest_item_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_drag_dest_item : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_icon_view_get_item_at_pos : unsupported parameter path : record with indirection level of 2

// GetReorderable is a wrapper around the C function gtk_icon_view_get_reorderable.
func (recv *IconView) GetReorderable() bool {
	retC := C.gtk_icon_view_get_reorderable((*C.GtkIconView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_icon_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// ScrollToPath is a wrapper around the C function gtk_icon_view_scroll_to_path.
func (recv *IconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign float32, colAlign float32) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_use_align :=
		boolToGboolean(useAlign)

	c_row_align := (C.gfloat)(rowAlign)

	c_col_align := (C.gfloat)(colAlign)

	C.gtk_icon_view_scroll_to_path((*C.GtkIconView)(recv.native), c_path, c_use_align, c_row_align, c_col_align)

	return
}

// SetCursor is a wrapper around the C function gtk_icon_view_set_cursor.
func (recv *IconView) SetCursor(path *TreePath, cell *CellRenderer, startEditing bool) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_icon_view_set_cursor((*C.GtkIconView)(recv.native), c_path, c_cell, c_start_editing)

	return
}

// SetDragDestItem is a wrapper around the C function gtk_icon_view_set_drag_dest_item.
func (recv *IconView) SetDragDestItem(path *TreePath, pos IconViewDropPosition) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_pos := (C.GtkIconViewDropPosition)(pos)

	C.gtk_icon_view_set_drag_dest_item((*C.GtkIconView)(recv.native), c_path, c_pos)

	return
}

// SetReorderable is a wrapper around the C function gtk_icon_view_set_reorderable.
func (recv *IconView) SetReorderable(reorderable bool) {
	c_reorderable :=
		boolToGboolean(reorderable)

	C.gtk_icon_view_set_reorderable((*C.GtkIconView)(recv.native), c_reorderable)

	return
}

// UnsetModelDragDest is a wrapper around the C function gtk_icon_view_unset_model_drag_dest.
func (recv *IconView) UnsetModelDragDest() {
	C.gtk_icon_view_unset_model_drag_dest((*C.GtkIconView)(recv.native))

	return
}

// UnsetModelDragSource is a wrapper around the C function gtk_icon_view_unset_model_drag_source.
func (recv *IconView) UnsetModelDragSource() {
	C.gtk_icon_view_unset_model_drag_source((*C.GtkIconView)(recv.native))

	return
}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Clear is a wrapper around the C function gtk_image_clear.
func (recv *Image) Clear() {
	C.gtk_image_clear((*C.GtkImage)(recv.native))

	return
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// GetChildPackDirection is a wrapper around the C function gtk_menu_bar_get_child_pack_direction.
func (recv *MenuBar) GetChildPackDirection() PackDirection {
	retC := C.gtk_menu_bar_get_child_pack_direction((*C.GtkMenuBar)(recv.native))
	retGo := (PackDirection)(retC)

	return retGo
}

// GetPackDirection is a wrapper around the C function gtk_menu_bar_get_pack_direction.
func (recv *MenuBar) GetPackDirection() PackDirection {
	retC := C.gtk_menu_bar_get_pack_direction((*C.GtkMenuBar)(recv.native))
	retGo := (PackDirection)(retC)

	return retGo
}

// SetChildPackDirection is a wrapper around the C function gtk_menu_bar_set_child_pack_direction.
func (recv *MenuBar) SetChildPackDirection(childPackDir PackDirection) {
	c_child_pack_dir := (C.GtkPackDirection)(childPackDir)

	C.gtk_menu_bar_set_child_pack_direction((*C.GtkMenuBar)(recv.native), c_child_pack_dir)

	return
}

// SetPackDirection is a wrapper around the C function gtk_menu_bar_set_pack_direction.
func (recv *MenuBar) SetPackDirection(packDir PackDirection) {
	c_pack_dir := (C.GtkPackDirection)(packDir)

	C.gtk_menu_bar_set_pack_direction((*C.GtkMenuBar)(recv.native), c_pack_dir)

	return
}

// GetTakeFocus is a wrapper around the C function gtk_menu_shell_get_take_focus.
func (recv *MenuShell) GetTakeFocus() bool {
	retC := C.gtk_menu_shell_get_take_focus((*C.GtkMenuShell)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetTakeFocus is a wrapper around the C function gtk_menu_shell_set_take_focus.
func (recv *MenuShell) SetTakeFocus(takeFocus bool) {
	c_take_focus :=
		boolToGboolean(takeFocus)

	C.gtk_menu_shell_set_take_focus((*C.GtkMenuShell)(recv.native), c_take_focus)

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

// GetHscrollbar is a wrapper around the C function gtk_scrolled_window_get_hscrollbar.
func (recv *ScrolledWindow) GetHscrollbar() *Widget {
	retC := C.gtk_scrolled_window_get_hscrollbar((*C.GtkScrolledWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVscrollbar is a wrapper around the C function gtk_scrolled_window_get_vscrollbar.
func (recv *ScrolledWindow) GetVscrollbar() *Widget {
	retC := C.gtk_scrolled_window_get_vscrollbar((*C.GtkScrolledWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIgnoreHidden is a wrapper around the C function gtk_size_group_get_ignore_hidden.
func (recv *SizeGroup) GetIgnoreHidden() bool {
	retC := C.gtk_size_group_get_ignore_hidden((*C.GtkSizeGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetIgnoreHidden is a wrapper around the C function gtk_size_group_set_ignore_hidden.
func (recv *SizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	c_ignore_hidden :=
		boolToGboolean(ignoreHidden)

	C.gtk_size_group_set_ignore_hidden((*C.GtkSizeGroup)(recv.native), c_ignore_hidden)

	return
}

// Unsupported signal 'input' for SpinButton : return value gint :

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported signal 'event' for TextTag : unsupported parameter event : no type generator for Gdk.Event,

// GetIconName is a wrapper around the C function gtk_tool_button_get_icon_name.
func (recv *ToolButton) GetIconName() string {
	retC := C.gtk_tool_button_get_icon_name((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetIconName is a wrapper around the C function gtk_tool_button_set_icon_name.
func (recv *ToolButton) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_tool_button_set_icon_name((*C.GtkToolButton)(recv.native), c_icon_name)

	return
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_get_visible_range : unsupported parameter start_path : record with indirection level of 2

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// QueueResize is a wrapper around the C function gtk_tree_view_column_queue_resize.
func (recv *TreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize((*C.GtkTreeViewColumn)(recv.native))

	return
}

// Unsupported signal 'child-notify' for Widget : unsupported parameter child_property : Blacklisted record : GParamSpec

// Unsupported signal 'delete-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'destroy-event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event' for Widget : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal 'event-after' for Widget : unsupported parameter event : no type generator for Gdk.Event,

var signalWidgetGrabBrokenEventId int
var signalWidgetGrabBrokenEventMap = make(map[int]WidgetSignalGrabBrokenEventCallback)
var signalWidgetGrabBrokenEventLock sync.Mutex

// WidgetSignalGrabBrokenEventCallback is a callback function for a 'grab-broken-event' signal emitted from a Widget.
type WidgetSignalGrabBrokenEventCallback func(event *gdk.EventGrabBroken) bool

/*
ConnectGrabBrokenEvent connects the callback to the 'grab-broken-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectGrabBrokenEvent to remove it.
*/
func (recv *Widget) ConnectGrabBrokenEvent(callback WidgetSignalGrabBrokenEventCallback) int {
	signalWidgetGrabBrokenEventLock.Lock()
	defer signalWidgetGrabBrokenEventLock.Unlock()

	signalWidgetGrabBrokenEventId++
	signalWidgetGrabBrokenEventMap[signalWidgetGrabBrokenEventId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Widget_signal_connect_grab_broken_event(instance, C.gpointer(uintptr(signalWidgetGrabBrokenEventId)))
	return int(retC)
}

/*
DisconnectGrabBrokenEvent disconnects a callback from the 'grab-broken-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectGrabBrokenEvent.
*/
func (recv *Widget) DisconnectGrabBrokenEvent(connectionID int) {
	_, exists := signalWidgetGrabBrokenEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalWidgetGrabBrokenEventMap, connectionID)
}

//export Widget_grabBrokenEventHandler
func Widget_grabBrokenEventHandler() C.boolean {
	fmt.Println("cb")
}

// Unsupported signal 'size-allocate' for Widget : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported signal 'touch-event' for Widget : unsupported parameter object : no type generator for Gdk.Event,

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// DragSourceSetIconName is a wrapper around the C function gtk_drag_source_set_icon_name.
func (recv *Widget) DragSourceSetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_drag_source_set_icon_name((*C.GtkWidget)(recv.native), c_icon_name)

	return
}

// GetUrgencyHint is a wrapper around the C function gtk_window_get_urgency_hint.
func (recv *Window) GetUrgencyHint() bool {
	retC := C.gtk_window_get_urgency_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PresentWithTime is a wrapper around the C function gtk_window_present_with_time.
func (recv *Window) PresentWithTime(timestamp uint32) {
	c_timestamp := (C.guint32)(timestamp)

	C.gtk_window_present_with_time((*C.GtkWindow)(recv.native), c_timestamp)

	return
}

// SetUrgencyHint is a wrapper around the C function gtk_window_set_urgency_hint.
func (recv *Window) SetUrgencyHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_urgency_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

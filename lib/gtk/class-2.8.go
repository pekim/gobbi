// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
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

	gboolean widget_grabBrokenEventHandler(GObject *, GdkEventGrabBroken *, gpointer);

	static gulong Widget_signal_connect_grab_broken_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "grab-broken-event", G_CALLBACK(widget_grabBrokenEventHandler), data);
	}

*/
import "C"

// Returns whether the license text in @about is
// automatically wrapped.
/*

C function : gtk_about_dialog_get_wrap_license
*/
func (recv *AboutDialog) GetWrapLicense() bool {
	retC := C.gtk_about_dialog_get_wrap_license((*C.GtkAboutDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the license text in @about is
// automatically wrapped.
/*

C function : gtk_about_dialog_set_wrap_license
*/
func (recv *AboutDialog) SetWrapLicense(wrapLicense bool) {
	c_wrap_license :=
		boolToGboolean(wrapLicense)

	C.gtk_about_dialog_set_wrap_license((*C.GtkAboutDialog)(recv.native), c_wrap_license)

	return
}

// Returns the accel closure for this action.
/*

C function : gtk_action_get_accel_closure
*/
func (recv *Action) GetAccelClosure() *gobject.Closure {
	retC := C.gtk_action_get_accel_closure((*C.GtkAction)(recv.native))
	retGo := gobject.ClosureNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the response id of a widget in the action area
// of a dialog.
/*

C function : gtk_dialog_get_response_for_widget
*/
func (recv *Dialog) GetResponseForWidget(widget *Widget) int32 {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_dialog_get_response_for_widget((*C.GtkDialog)(recv.native), c_widget)
	retGo := (int32)(retC)

	return retGo
}

// Returns whether the  completion popup window will be resized to the
// width of the entry.
/*

C function : gtk_entry_completion_get_popup_set_width
*/
func (recv *EntryCompletion) GetPopupSetWidth() bool {
	retC := C.gtk_entry_completion_get_popup_set_width((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether the completion popup window will appear even if there is
// only a single match.
/*

C function : gtk_entry_completion_get_popup_single_match
*/
func (recv *EntryCompletion) GetPopupSingleMatch() bool {
	retC := C.gtk_entry_completion_get_popup_single_match((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the completion popup window will be resized to be the same
// width as the entry.
/*

C function : gtk_entry_completion_set_popup_set_width
*/
func (recv *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	c_popup_set_width :=
		boolToGboolean(popupSetWidth)

	C.gtk_entry_completion_set_popup_set_width((*C.GtkEntryCompletion)(recv.native), c_popup_set_width)

	return
}

// Sets whether the completion popup window will appear even if there is
// only a single match. You may want to set this to %FALSE if you
// are using [inline completion][GtkEntryCompletion--inline-completion].
/*

C function : gtk_entry_completion_set_popup_single_match
*/
func (recv *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	c_popup_single_match :=
		boolToGboolean(popupSingleMatch)

	C.gtk_entry_completion_set_popup_single_match((*C.GtkEntryCompletion)(recv.native), c_popup_single_match)

	return
}

// Creates a #cairo_surface_t representation of the item at @path.
// This image is used for a drag icon.
/*

C function : gtk_icon_view_create_drag_icon
*/
func (recv *IconView) CreateDragIcon(path *TreePath) *cairo.Surface {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	retC := C.gtk_icon_view_create_drag_icon((*C.GtkIconView)(recv.native), c_path)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_view_enable_model_drag_dest : unsupported parameter targets :

// Unsupported : gtk_icon_view_enable_model_drag_source : unsupported parameter targets :

// Fills in @path and @cell with the current cursor path and cell.
// If the cursor isn’t currently set, then *@path will be %NULL.
// If no cell currently has focus, then *@cell will be %NULL.
//
// The returned #GtkTreePath must be freed with gtk_tree_path_free().
/*

C function : gtk_icon_view_get_cursor
*/
func (recv *IconView) GetCursor() (bool, *TreePath, *CellRenderer) {
	var c_path *C.GtkTreePath

	var c_cell *C.GtkCellRenderer

	retC := C.gtk_icon_view_get_cursor((*C.GtkIconView)(recv.native), &c_path, &c_cell)
	retGo := retC == C.TRUE

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	cell := CellRendererNewFromC(unsafe.Pointer(c_cell))

	return retGo, path, cell
}

// Unsupported : gtk_icon_view_get_dest_item_at_pos : unsupported parameter pos : GtkIconViewDropPosition* with indirection level of 1

// Unsupported : gtk_icon_view_get_drag_dest_item : unsupported parameter pos : GtkIconViewDropPosition* with indirection level of 1

// Finds the path at the point (@x, @y), relative to bin_window coordinates.
// In contrast to gtk_icon_view_get_path_at_pos(), this function also
// obtains the cell at the specified position. The returned path should
// be freed with gtk_tree_path_free().
// See gtk_icon_view_convert_widget_to_bin_window_coords() for converting
// widget coordinates to bin_window coordinates.
/*

C function : gtk_icon_view_get_item_at_pos
*/
func (recv *IconView) GetItemAtPos(x int32, y int32) (bool, *TreePath, *CellRenderer) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	var c_path *C.GtkTreePath

	var c_cell *C.GtkCellRenderer

	retC := C.gtk_icon_view_get_item_at_pos((*C.GtkIconView)(recv.native), c_x, c_y, &c_path, &c_cell)
	retGo := retC == C.TRUE

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	cell := CellRendererNewFromC(unsafe.Pointer(c_cell))

	return retGo, path, cell
}

// Retrieves whether the user can reorder the list via drag-and-drop.
// See gtk_icon_view_set_reorderable().
/*

C function : gtk_icon_view_get_reorderable
*/
func (recv *IconView) GetReorderable() bool {
	retC := C.gtk_icon_view_get_reorderable((*C.GtkIconView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets @start_path and @end_path to be the first and last visible path.
// Note that there may be invisible paths in between.
//
// Both paths should be freed with gtk_tree_path_free() after use.
/*

C function : gtk_icon_view_get_visible_range
*/
func (recv *IconView) GetVisibleRange() (bool, *TreePath, *TreePath) {
	var c_start_path *C.GtkTreePath

	var c_end_path *C.GtkTreePath

	retC := C.gtk_icon_view_get_visible_range((*C.GtkIconView)(recv.native), &c_start_path, &c_end_path)
	retGo := retC == C.TRUE

	startPath := TreePathNewFromC(unsafe.Pointer(c_start_path))

	endPath := TreePathNewFromC(unsafe.Pointer(c_end_path))

	return retGo, startPath, endPath
}

// Moves the alignments of @icon_view to the position specified by @path.
// @row_align determines where the row is placed, and @col_align determines
// where @column is placed.  Both are expected to be between 0.0 and 1.0.
// 0.0 means left/top alignment, 1.0 means right/bottom alignment, 0.5 means
// center.
//
// If @use_align is %FALSE, then the alignment arguments are ignored, and the
// tree does the minimum amount of work to scroll the item onto the screen.
// This means that the item will be scrolled to the edge closest to its current
// position.  If the item is currently visible on the screen, nothing is done.
//
// This function only works if the model is set, and @path is a valid row on
// the model. If the model changes before the @icon_view is realized, the
// centered path will be modified to reflect this change.
/*

C function : gtk_icon_view_scroll_to_path
*/
func (recv *IconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign float32, colAlign float32) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_use_align :=
		boolToGboolean(useAlign)

	c_row_align := (C.gfloat)(rowAlign)

	c_col_align := (C.gfloat)(colAlign)

	C.gtk_icon_view_scroll_to_path((*C.GtkIconView)(recv.native), c_path, c_use_align, c_row_align, c_col_align)

	return
}

// Sets the current keyboard focus to be at @path, and selects it.  This is
// useful when you want to focus the user’s attention on a particular item.
// If @cell is not %NULL, then focus is given to the cell specified by
// it. Additionally, if @start_editing is %TRUE, then editing should be
// started in the specified cell.
//
// This function is often followed by `gtk_widget_grab_focus
// (icon_view)` in order to give keyboard focus to the widget.
// Please note that editing can only happen when the widget is realized.
/*

C function : gtk_icon_view_set_cursor
*/
func (recv *IconView) SetCursor(path *TreePath, cell *CellRenderer, startEditing bool) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_icon_view_set_cursor((*C.GtkIconView)(recv.native), c_path, c_cell, c_start_editing)

	return
}

// Sets the item that is highlighted for feedback.
/*

C function : gtk_icon_view_set_drag_dest_item
*/
func (recv *IconView) SetDragDestItem(path *TreePath, pos IconViewDropPosition) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_pos := (C.GtkIconViewDropPosition)(pos)

	C.gtk_icon_view_set_drag_dest_item((*C.GtkIconView)(recv.native), c_path, c_pos)

	return
}

// This function is a convenience function to allow you to reorder models that
// support the #GtkTreeDragSourceIface and the #GtkTreeDragDestIface.  Both
// #GtkTreeStore and #GtkListStore support these.  If @reorderable is %TRUE, then
// the user can reorder the model by dragging and dropping rows.  The
// developer can listen to these changes by connecting to the model's
// row_inserted and row_deleted signals. The reordering is implemented by setting up
// the icon view as a drag source and destination. Therefore, drag and
// drop can not be used in a reorderable view for any other purpose.
//
// This function does not give you any degree of control over the order -- any
// reordering is allowed.  If more control is needed, you should probably
// handle drag and drop manually.
/*

C function : gtk_icon_view_set_reorderable
*/
func (recv *IconView) SetReorderable(reorderable bool) {
	c_reorderable :=
		boolToGboolean(reorderable)

	C.gtk_icon_view_set_reorderable((*C.GtkIconView)(recv.native), c_reorderable)

	return
}

// Undoes the effect of gtk_icon_view_enable_model_drag_dest(). Calling this
// method sets #GtkIconView:reorderable to %FALSE.
/*

C function : gtk_icon_view_unset_model_drag_dest
*/
func (recv *IconView) UnsetModelDragDest() {
	C.gtk_icon_view_unset_model_drag_dest((*C.GtkIconView)(recv.native))

	return
}

// Undoes the effect of gtk_icon_view_enable_model_drag_source(). Calling this
// method sets #GtkIconView:reorderable to %FALSE.
/*

C function : gtk_icon_view_unset_model_drag_source
*/
func (recv *IconView) UnsetModelDragSource() {
	C.gtk_icon_view_unset_model_drag_source((*C.GtkIconView)(recv.native))

	return
}

// Resets the image to be empty.
/*

C function : gtk_image_clear
*/
func (recv *Image) Clear() {
	C.gtk_image_clear((*C.GtkImage)(recv.native))

	return
}

// Retrieves the current child pack direction of the menubar.
// See gtk_menu_bar_set_child_pack_direction().
/*

C function : gtk_menu_bar_get_child_pack_direction
*/
func (recv *MenuBar) GetChildPackDirection() PackDirection {
	retC := C.gtk_menu_bar_get_child_pack_direction((*C.GtkMenuBar)(recv.native))
	retGo := (PackDirection)(retC)

	return retGo
}

// Retrieves the current pack direction of the menubar.
// See gtk_menu_bar_set_pack_direction().
/*

C function : gtk_menu_bar_get_pack_direction
*/
func (recv *MenuBar) GetPackDirection() PackDirection {
	retC := C.gtk_menu_bar_get_pack_direction((*C.GtkMenuBar)(recv.native))
	retGo := (PackDirection)(retC)

	return retGo
}

// Sets how widgets should be packed inside the children of a menubar.
/*

C function : gtk_menu_bar_set_child_pack_direction
*/
func (recv *MenuBar) SetChildPackDirection(childPackDir PackDirection) {
	c_child_pack_dir := (C.GtkPackDirection)(childPackDir)

	C.gtk_menu_bar_set_child_pack_direction((*C.GtkMenuBar)(recv.native), c_child_pack_dir)

	return
}

// Sets how items should be packed inside a menubar.
/*

C function : gtk_menu_bar_set_pack_direction
*/
func (recv *MenuBar) SetPackDirection(packDir PackDirection) {
	c_pack_dir := (C.GtkPackDirection)(packDir)

	C.gtk_menu_bar_set_pack_direction((*C.GtkMenuBar)(recv.native), c_pack_dir)

	return
}

// Returns %TRUE if the menu shell will take the keyboard focus on popup.
/*

C function : gtk_menu_shell_get_take_focus
*/
func (recv *MenuShell) GetTakeFocus() bool {
	retC := C.gtk_menu_shell_get_take_focus((*C.GtkMenuShell)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// If @take_focus is %TRUE (the default) the menu shell will take
// the keyboard focus so that it will receive all keyboard events
// which is needed to enable keyboard navigation in menus.
//
// Setting @take_focus to %FALSE is useful only for special applications
// like virtual keyboard implementations which should not take keyboard
// focus.
//
// The @take_focus state of a menu or menu bar is automatically
// propagated to submenus whenever a submenu is popped up, so you
// don’t have to worry about recursively setting it for your entire
// menu hierarchy. Only when programmatically picking a submenu and
// popping it up manually, the @take_focus property of the submenu
// needs to be set explicitly.
//
// Note that setting it to %FALSE has side-effects:
//
// If the focus is in some other app, it keeps the focus and keynav in
// the menu doesn’t work. Consequently, keynav on the menu will only
// work if the focus is on some toplevel owned by the onscreen keyboard.
//
// To avoid confusing the user, menus with @take_focus set to %FALSE
// should not display mnemonics or accelerators, since it cannot be
// guaranteed that they will work.
//
// See also gdk_keyboard_grab()
/*

C function : gtk_menu_shell_set_take_focus
*/
func (recv *MenuShell) SetTakeFocus(takeFocus bool) {
	c_take_focus :=
		boolToGboolean(takeFocus)

	C.gtk_menu_shell_set_take_focus((*C.GtkMenuShell)(recv.native), c_take_focus)

	return
}

// Returns the horizontal scrollbar of @scrolled_window.
/*

C function : gtk_scrolled_window_get_hscrollbar
*/
func (recv *ScrolledWindow) GetHscrollbar() *Widget {
	retC := C.gtk_scrolled_window_get_hscrollbar((*C.GtkScrolledWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the vertical scrollbar of @scrolled_window.
/*

C function : gtk_scrolled_window_get_vscrollbar
*/
func (recv *ScrolledWindow) GetVscrollbar() *Widget {
	retC := C.gtk_scrolled_window_get_vscrollbar((*C.GtkScrolledWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns if invisible widgets are ignored when calculating the size.
/*

C function : gtk_size_group_get_ignore_hidden
*/
func (recv *SizeGroup) GetIgnoreHidden() bool {
	retC := C.gtk_size_group_get_ignore_hidden((*C.GtkSizeGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether unmapped widgets should be ignored when
// calculating the size.
/*

C function : gtk_size_group_set_ignore_hidden
*/
func (recv *SizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	c_ignore_hidden :=
		boolToGboolean(ignoreHidden)

	C.gtk_size_group_set_ignore_hidden((*C.GtkSizeGroup)(recv.native), c_ignore_hidden)

	return
}

// Returns the name of the themed icon for the tool button,
// see gtk_tool_button_set_icon_name().
/*

C function : gtk_tool_button_get_icon_name
*/
func (recv *ToolButton) GetIconName() string {
	retC := C.gtk_tool_button_get_icon_name((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the icon for the tool button from a named themed icon.
// See the docs for #GtkIconTheme for more details.
// The #GtkToolButton:icon-name property only has an effect if not
// overridden by non-%NULL #GtkToolButton:label-widget,
// #GtkToolButton:icon-widget and #GtkToolButton:stock-id properties.
/*

C function : gtk_tool_button_set_icon_name
*/
func (recv *ToolButton) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_tool_button_set_icon_name((*C.GtkToolButton)(recv.native), c_icon_name)

	return
}

// Sets @start_path and @end_path to be the first and last visible path.
// Note that there may be invisible paths in between.
//
// The paths should be freed with gtk_tree_path_free() after use.
/*

C function : gtk_tree_view_get_visible_range
*/
func (recv *TreeView) GetVisibleRange() (bool, *TreePath, *TreePath) {
	var c_start_path *C.GtkTreePath

	var c_end_path *C.GtkTreePath

	retC := C.gtk_tree_view_get_visible_range((*C.GtkTreeView)(recv.native), &c_start_path, &c_end_path)
	retGo := retC == C.TRUE

	startPath := TreePathNewFromC(unsafe.Pointer(c_start_path))

	endPath := TreePathNewFromC(unsafe.Pointer(c_end_path))

	return retGo, startPath, endPath
}

// Flags the column, and the cell renderers added to this column, to have
// their sizes renegotiated.
/*

C function : gtk_tree_view_column_queue_resize
*/
func (recv *TreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize((*C.GtkTreeViewColumn)(recv.native))

	return
}

type signalWidgetGrabBrokenEventDetail struct {
	callback  WidgetSignalGrabBrokenEventCallback
	handlerID C.gulong
}

var signalWidgetGrabBrokenEventId int
var signalWidgetGrabBrokenEventMap = make(map[int]signalWidgetGrabBrokenEventDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_grab_broken_event(instance, C.gpointer(uintptr(signalWidgetGrabBrokenEventId)))

	detail := signalWidgetGrabBrokenEventDetail{callback, handlerID}
	signalWidgetGrabBrokenEventMap[signalWidgetGrabBrokenEventId] = detail

	return signalWidgetGrabBrokenEventId
}

/*
DisconnectGrabBrokenEvent disconnects a callback from the 'grab-broken-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectGrabBrokenEvent.
*/
func (recv *Widget) DisconnectGrabBrokenEvent(connectionID int) {
	signalWidgetGrabBrokenEventLock.Lock()
	defer signalWidgetGrabBrokenEventLock.Unlock()

	detail, exists := signalWidgetGrabBrokenEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetGrabBrokenEventMap, connectionID)
}

//export widget_grabBrokenEventHandler
func widget_grabBrokenEventHandler(_ *C.GObject, c_event *C.GdkEventGrabBroken, data C.gpointer) C.gboolean {
	event := gdk.EventGrabBrokenNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetGrabBrokenEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Sets the icon that will be used for drags from a particular source
// to a themed icon. See the docs for #GtkIconTheme for more details.
/*

C function : gtk_drag_source_set_icon_name
*/
func (recv *Widget) DragSourceSetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_drag_source_set_icon_name((*C.GtkWidget)(recv.native), c_icon_name)

	return
}

// Gets the value set by gtk_window_set_urgency_hint()
/*

C function : gtk_window_get_urgency_hint
*/
func (recv *Window) GetUrgencyHint() bool {
	retC := C.gtk_window_get_urgency_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Presents a window to the user in response to a user interaction.
// If you need to present a window without a timestamp, use
// gtk_window_present(). See gtk_window_present() for details.
/*

C function : gtk_window_present_with_time
*/
func (recv *Window) PresentWithTime(timestamp uint32) {
	c_timestamp := (C.guint32)(timestamp)

	C.gtk_window_present_with_time((*C.GtkWindow)(recv.native), c_timestamp)

	return
}

// Windows may set a hint asking the desktop environment to draw
// the users attention to the window. This function sets this hint.
/*

C function : gtk_window_set_urgency_hint
*/
func (recv *Window) SetUrgencyHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_urgency_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

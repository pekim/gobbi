// This is a generated file - DO NOT EDIT
// +build gtk_2.2 gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Gets the #GdkDisplay associated with @clipboard
/*

C function

gtk_clipboard_get_display
*/
func (recv *Clipboard) GetDisplay() *gdk.Display {
	retC := C.gtk_clipboard_get_display((*C.GtkClipboard)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkInvisible object for a specified screen
/*

C function

gtk_invisible_new_for_screen
*/
func InvisibleNewForScreen(screen *gdk.Screen) *Invisible {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gtk_invisible_new_for_screen(c_screen)
	retGo := InvisibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GdkScreen object associated with @invisible
/*

C function

gtk_invisible_get_screen
*/
func (recv *Invisible) GetScreen() *gdk.Screen {
	retC := C.gtk_invisible_get_screen((*C.GtkInvisible)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the #GdkScreen where the #GtkInvisible object will be displayed.
/*

C function

gtk_invisible_set_screen
*/
func (recv *Invisible) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_invisible_set_screen((*C.GtkInvisible)(recv.native), c_screen)

	return
}

// > This function is slow. Only use it for debugging and/or testing
// > purposes.
//
// Checks if the given iter is a valid iter for this #GtkListStore.
/*

C function

gtk_list_store_iter_is_valid
*/
func (recv *ListStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_list_store_iter_is_valid((*C.GtkListStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Moves @iter in @store to the position after @position. Note that this
// function only works with unsorted stores. If @position is %NULL, @iter
// will be moved to the start of the list.
/*

C function

gtk_list_store_move_after
*/
func (recv *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_list_store_move_after((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// Moves @iter in @store to the position before @position. Note that this
// function only works with unsorted stores. If @position is %NULL, @iter
// will be moved to the end of the list.
/*

C function

gtk_list_store_move_before
*/
func (recv *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_list_store_move_before((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// Reorders @store to follow the order indicated by @new_order. Note that
// this function only works with unsorted stores.
/*

C function

gtk_list_store_reorder
*/
func (recv *ListStore) Reorder(newOrder []int32) {
	c_new_order := &newOrder[0]

	C.gtk_list_store_reorder((*C.GtkListStore)(recv.native), (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Swaps @a and @b in @store. Note that this function only works with
// unsorted stores.
/*

C function

gtk_list_store_swap
*/
func (recv *ListStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(C.NULL)
	if a != nil {
		c_a = (*C.GtkTreeIter)(a.ToC())
	}

	c_b := (*C.GtkTreeIter)(C.NULL)
	if b != nil {
		c_b = (*C.GtkTreeIter)(b.ToC())
	}

	C.gtk_list_store_swap((*C.GtkListStore)(recv.native), c_a, c_b)

	return
}

// Sets the #GdkScreen on which the menu will be displayed.
/*

C function

gtk_menu_set_screen
*/
func (recv *Menu) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_menu_set_screen((*C.GtkMenu)(recv.native), c_screen)

	return
}

// Select the first visible or selectable child of the menu shell;
// don’t select tearoff items unless the only item is a tearoff
// item.
/*

C function

gtk_menu_shell_select_first
*/
func (recv *MenuShell) SelectFirst(searchSensitive bool) {
	c_search_sensitive :=
		boolToGboolean(searchSensitive)

	C.gtk_menu_shell_select_first((*C.GtkMenuShell)(recv.native), c_search_sensitive)

	return
}

// Gets the number of pages in a notebook.
/*

C function

gtk_notebook_get_n_pages
*/
func (recv *Notebook) GetNPages() int32 {
	retC := C.gtk_notebook_get_n_pages((*C.GtkNotebook)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// > This function is slow. Only use it for debugging and/or testing
// > purposes.
//
// Checks if the given iter is a valid iter for this #GtkTreeModelSort.
/*

C function

gtk_tree_model_sort_iter_is_valid
*/
func (recv *TreeModelSort) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_sort_iter_is_valid((*C.GtkTreeModelSort)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Returns the number of rows that have been selected in @tree.
/*

C function

gtk_tree_selection_count_selected_rows
*/
func (recv *TreeSelection) CountSelectedRows() int32 {
	retC := C.gtk_tree_selection_count_selected_rows((*C.GtkTreeSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Creates a list of path of all selected rows. Additionally, if you are
// planning on modifying the model after calling this function, you may
// want to convert the returned list into a list of #GtkTreeRowReferences.
// To do this, you can use gtk_tree_row_reference_new().
//
// To free the return value, use:
// |[<!-- language="C" -->
// g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);
// ]|
/*

C function

gtk_tree_selection_get_selected_rows
*/
func (recv *TreeSelection) GetSelectedRows() (*glib.List, *TreeModel) {
	var c_model *C.GtkTreeModel

	retC := C.gtk_tree_selection_get_selected_rows((*C.GtkTreeSelection)(recv.native), &c_model)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	return retGo, model
}

// Unselects a range of nodes, determined by @start_path and @end_path
// inclusive.
/*

C function

gtk_tree_selection_unselect_range
*/
func (recv *TreeSelection) UnselectRange(startPath *TreePath, endPath *TreePath) {
	c_start_path := (*C.GtkTreePath)(C.NULL)
	if startPath != nil {
		c_start_path = (*C.GtkTreePath)(startPath.ToC())
	}

	c_end_path := (*C.GtkTreePath)(C.NULL)
	if endPath != nil {
		c_end_path = (*C.GtkTreePath)(endPath.ToC())
	}

	C.gtk_tree_selection_unselect_range((*C.GtkTreeSelection)(recv.native), c_start_path, c_end_path)

	return
}

// WARNING: This function is slow. Only use it for debugging and/or testing
// purposes.
//
// Checks if the given iter is a valid iter for this #GtkTreeStore.
/*

C function

gtk_tree_store_iter_is_valid
*/
func (recv *TreeStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_store_iter_is_valid((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Moves @iter in @tree_store to the position after @position. @iter and
// @position should be in the same level. Note that this function only
// works with unsorted stores. If @position is %NULL, @iter will be moved
// to the start of the level.
/*

C function

gtk_tree_store_move_after
*/
func (recv *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_tree_store_move_after((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// Moves @iter in @tree_store to the position before @position. @iter and
// @position should be in the same level. Note that this function only
// works with unsorted stores. If @position is %NULL, @iter will be
// moved to the end of the level.
/*

C function

gtk_tree_store_move_before
*/
func (recv *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_position := (*C.GtkTreeIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTreeIter)(position.ToC())
	}

	C.gtk_tree_store_move_before((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// Reorders the children of @parent in @tree_store to follow the order
// indicated by @new_order. Note that this function only works with
// unsorted stores.
/*

C function

gtk_tree_store_reorder
*/
func (recv *TreeStore) Reorder(parent *TreeIter, newOrder []int32) {
	c_parent := (*C.GtkTreeIter)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkTreeIter)(parent.ToC())
	}

	c_new_order := &newOrder[0]

	C.gtk_tree_store_reorder((*C.GtkTreeStore)(recv.native), c_parent, (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Swaps @a and @b in the same level of @tree_store. Note that this function
// only works with unsorted stores.
/*

C function

gtk_tree_store_swap
*/
func (recv *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(C.NULL)
	if a != nil {
		c_a = (*C.GtkTreeIter)(a.ToC())
	}

	c_b := (*C.GtkTreeIter)(C.NULL)
	if b != nil {
		c_b = (*C.GtkTreeIter)(b.ToC())
	}

	C.gtk_tree_store_swap((*C.GtkTreeStore)(recv.native), c_a, c_b)

	return
}

// Expands the row at @path. This will also expand all parent rows of
// @path as necessary.
/*

C function

gtk_tree_view_expand_to_path
*/
func (recv *TreeView) ExpandToPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_view_expand_to_path((*C.GtkTreeView)(recv.native), c_path)

	return
}

// Sets the current keyboard focus to be at @path, and selects it.  This is
// useful when you want to focus the user’s attention on a particular row.  If
// @focus_column is not %NULL, then focus is given to the column specified by
// it. If @focus_column and @focus_cell are not %NULL, and @focus_column
// contains 2 or more editable or activatable cells, then focus is given to
// the cell specified by @focus_cell. Additionally, if @focus_column is
// specified, and @start_editing is %TRUE, then editing should be started in
// the specified cell.  This function is often followed by
// @gtk_widget_grab_focus (@tree_view) in order to give keyboard focus to the
// widget.  Please note that editing can only happen when the widget is
// realized.
//
// If @path is invalid for @model, the current cursor (if any) will be unset
// and the function will return without failing.
/*

C function

gtk_tree_view_set_cursor_on_cell
*/
func (recv *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing bool) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_focus_column := (*C.GtkTreeViewColumn)(C.NULL)
	if focusColumn != nil {
		c_focus_column = (*C.GtkTreeViewColumn)(focusColumn.ToC())
	}

	c_focus_cell := (*C.GtkCellRenderer)(C.NULL)
	if focusCell != nil {
		c_focus_cell = (*C.GtkCellRenderer)(focusCell.ToC())
	}

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_tree_view_set_cursor_on_cell((*C.GtkTreeView)(recv.native), c_path, c_focus_column, c_focus_cell, c_start_editing)

	return
}

// Sets the current keyboard focus to be at @cell, if the column contains
// 2 or more editable and activatable cells.
/*

C function

gtk_tree_view_column_focus_cell
*/
func (recv *TreeViewColumn) FocusCell(cell *CellRenderer) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_tree_view_column_focus_cell((*C.GtkTreeViewColumn)(recv.native), c_cell)

	return
}

// Unsupported : gtk_widget_get_clipboard : unsupported parameter selection : Blacklisted record : GdkAtom

// Get the #GdkDisplay for the toplevel window associated with
// this widget. This function can only be called after the widget
// has been added to a widget hierarchy with a #GtkWindow at the top.
//
// In general, you should only create display specific
// resources when a widget has been realized, and you should
// free those resources when the widget is unrealized.
/*

C function

gtk_widget_get_display
*/
func (recv *Widget) GetDisplay() *gdk.Display {
	retC := C.gtk_widget_get_display((*C.GtkWidget)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get the root window where this widget is located. This function can
// only be called after the widget has been added to a widget
// hierarchy with #GtkWindow at the top.
//
// The root window is useful for such purposes as creating a popup
// #GdkWindow associated with the window. In general, you should only
// create display specific resources when a widget has been realized,
// and you should free those resources when the widget is unrealized.
/*

C function

gtk_widget_get_root_window
*/
func (recv *Widget) GetRootWindow() *gdk.Window {
	retC := C.gtk_widget_get_root_window((*C.GtkWidget)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get the #GdkScreen from the toplevel window associated with
// this widget. This function can only be called after the widget
// has been added to a widget hierarchy with a #GtkWindow
// at the top.
//
// In general, you should only create screen specific
// resources when a widget has been realized, and you should
// free those resources when the widget is unrealized.
/*

C function

gtk_widget_get_screen
*/
func (recv *Widget) GetScreen() *gdk.Screen {
	retC := C.gtk_widget_get_screen((*C.GtkWidget)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks whether there is a #GdkScreen is associated with
// this widget. All toplevel widgets have an associated
// screen, and all widgets added into a hierarchy with a toplevel
// window at the top.
/*

C function

gtk_widget_has_screen
*/
func (recv *Widget) HasScreen() bool {
	retC := C.gtk_widget_has_screen((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Asks to place @window in the fullscreen state. Note that you
// shouldn’t assume the window is definitely full screen afterward,
// because other entities (e.g. the user or
// [window manager][gtk-X11-arch]) could unfullscreen it
// again, and not all window managers honor requests to fullscreen
// windows. But normally the window will end up fullscreen. Just
// don’t write code that crashes if not.
//
// You can track the fullscreen state via the “window-state-event” signal
// on #GtkWidget.
/*

C function

gtk_window_fullscreen
*/
func (recv *Window) Fullscreen() {
	C.gtk_window_fullscreen((*C.GtkWindow)(recv.native))

	return
}

// Returns the #GdkScreen associated with @window.
/*

C function

gtk_window_get_screen
*/
func (recv *Window) GetScreen() *gdk.Screen {
	retC := C.gtk_window_get_screen((*C.GtkWindow)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the value set by gtk_window_set_skip_pager_hint().
/*

C function

gtk_window_get_skip_pager_hint
*/
func (recv *Window) GetSkipPagerHint() bool {
	retC := C.gtk_window_get_skip_pager_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value set by gtk_window_set_skip_taskbar_hint()
/*

C function

gtk_window_get_skip_taskbar_hint
*/
func (recv *Window) GetSkipTaskbarHint() bool {
	retC := C.gtk_window_get_skip_taskbar_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the icon for @window.
// Warns on failure if @err is %NULL.
//
// This function is equivalent to calling gtk_window_set_icon()
// with a pixbuf created by loading the image from @filename.
/*

C function

gtk_window_set_icon_from_file
*/
func (recv *Window) SetIconFromFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_window_set_icon_from_file((*C.GtkWindow)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the #GdkScreen where the @window is displayed; if
// the window is already mapped, it will be unmapped, and
// then remapped on the new screen.
/*

C function

gtk_window_set_screen
*/
func (recv *Window) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_window_set_screen((*C.GtkWindow)(recv.native), c_screen)

	return
}

// Windows may set a hint asking the desktop environment not to display
// the window in the pager. This function sets this hint.
// (A "pager" is any desktop navigation tool such as a workspace
// switcher that displays a thumbnail representation of the windows
// on the screen.)
/*

C function

gtk_window_set_skip_pager_hint
*/
func (recv *Window) SetSkipPagerHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_pager_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Windows may set a hint asking the desktop environment not to display
// the window in the task bar. This function sets this hint.
/*

C function

gtk_window_set_skip_taskbar_hint
*/
func (recv *Window) SetSkipTaskbarHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_taskbar_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Asks to toggle off the fullscreen state for @window. Note that you
// shouldn’t assume the window is definitely not full screen
// afterward, because other entities (e.g. the user or
// [window manager][gtk-X11-arch]) could fullscreen it
// again, and not all window managers honor requests to unfullscreen
// windows. But normally the window will end up restored to its normal
// state. Just don’t write code that crashes if not.
//
// You can track the fullscreen state via the “window-state-event” signal
// on #GtkWidget.
/*

C function

gtk_window_unfullscreen
*/
func (recv *Window) Unfullscreen() {
	C.gtk_window_unfullscreen((*C.GtkWindow)(recv.native))

	return
}

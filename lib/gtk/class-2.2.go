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

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File (GFile*) for param file

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// GetDisplay is a wrapper around the C function gtk_clipboard_get_display.
func (recv *Clipboard) GetDisplay() *gdk.Display {
	retC := C.gtk_clipboard_get_display((*C.GtkClipboard)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// InvisibleNewForScreen is a wrapper around the C function gtk_invisible_new_for_screen.
func InvisibleNewForScreen(screen *gdk.Screen) *Invisible {
	c_screen := (*C.GdkScreen)(screen.ToC())

	retC := C.gtk_invisible_new_for_screen(c_screen)
	retGo := InvisibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_invisible_get_screen.
func (recv *Invisible) GetScreen() *gdk.Screen {
	retC := C.gtk_invisible_get_screen((*C.GtkInvisible)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetScreen is a wrapper around the C function gtk_invisible_set_screen.
func (recv *Invisible) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_invisible_set_screen((*C.GtkInvisible)(recv.native), c_screen)

	return
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// IterIsValid is a wrapper around the C function gtk_list_store_iter_is_valid.
func (recv *ListStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_list_store_iter_is_valid((*C.GtkListStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// MoveAfter is a wrapper around the C function gtk_list_store_move_after.
func (recv *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_list_store_move_after((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// MoveBefore is a wrapper around the C function gtk_list_store_move_before.
func (recv *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_list_store_move_before((*C.GtkListStore)(recv.native), c_iter, c_position)

	return
}

// Reorder is a wrapper around the C function gtk_list_store_reorder.
func (recv *ListStore) Reorder(newOrder []int32) {
	c_new_order := &newOrder[0]

	C.gtk_list_store_reorder((*C.GtkListStore)(recv.native), (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Swap is a wrapper around the C function gtk_list_store_swap.
func (recv *ListStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(a.ToC())

	c_b := (*C.GtkTreeIter)(b.ToC())

	C.gtk_list_store_swap((*C.GtkListStore)(recv.native), c_a, c_b)

	return
}

// SetScreen is a wrapper around the C function gtk_menu_set_screen.
func (recv *Menu) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_menu_set_screen((*C.GtkMenu)(recv.native), c_screen)

	return
}

// SelectFirst is a wrapper around the C function gtk_menu_shell_select_first.
func (recv *MenuShell) SelectFirst(searchSensitive bool) {
	c_search_sensitive :=
		boolToGboolean(searchSensitive)

	C.gtk_menu_shell_select_first((*C.GtkMenuShell)(recv.native), c_search_sensitive)

	return
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// GetNPages is a wrapper around the C function gtk_notebook_get_n_pages.
func (recv *Notebook) GetNPages() int32 {
	retC := C.gtk_notebook_get_n_pages((*C.GtkNotebook)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// IterIsValid is a wrapper around the C function gtk_tree_model_sort_iter_is_valid.
func (recv *TreeModelSort) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_model_sort_iter_is_valid((*C.GtkTreeModelSort)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// CountSelectedRows is a wrapper around the C function gtk_tree_selection_count_selected_rows.
func (recv *TreeSelection) CountSelectedRows() int32 {
	retC := C.gtk_tree_selection_count_selected_rows((*C.GtkTreeSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_selection_get_selected_rows : unsupported parameter model : no type generator for TreeModel (GtkTreeModel**) for param model

// UnselectRange is a wrapper around the C function gtk_tree_selection_unselect_range.
func (recv *TreeSelection) UnselectRange(startPath *TreePath, endPath *TreePath) {
	c_start_path := (*C.GtkTreePath)(startPath.ToC())

	c_end_path := (*C.GtkTreePath)(endPath.ToC())

	C.gtk_tree_selection_unselect_range((*C.GtkTreeSelection)(recv.native), c_start_path, c_end_path)

	return
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// IterIsValid is a wrapper around the C function gtk_tree_store_iter_is_valid.
func (recv *TreeStore) IterIsValid(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_store_iter_is_valid((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// MoveAfter is a wrapper around the C function gtk_tree_store_move_after.
func (recv *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_tree_store_move_after((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// MoveBefore is a wrapper around the C function gtk_tree_store_move_before.
func (recv *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_position := (*C.GtkTreeIter)(position.ToC())

	C.gtk_tree_store_move_before((*C.GtkTreeStore)(recv.native), c_iter, c_position)

	return
}

// Reorder is a wrapper around the C function gtk_tree_store_reorder.
func (recv *TreeStore) Reorder(parent *TreeIter, newOrder []int32) {
	c_parent := (*C.GtkTreeIter)(parent.ToC())

	c_new_order := &newOrder[0]

	C.gtk_tree_store_reorder((*C.GtkTreeStore)(recv.native), c_parent, (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Swap is a wrapper around the C function gtk_tree_store_swap.
func (recv *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	c_a := (*C.GtkTreeIter)(a.ToC())

	c_b := (*C.GtkTreeIter)(b.ToC())

	C.gtk_tree_store_swap((*C.GtkTreeStore)(recv.native), c_a, c_b)

	return
}

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel (GtkTreeModel*) for param model

// ExpandToPath is a wrapper around the C function gtk_tree_view_expand_to_path.
func (recv *TreeView) ExpandToPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_tree_view_expand_to_path((*C.GtkTreeView)(recv.native), c_path)

	return
}

// SetCursorOnCell is a wrapper around the C function gtk_tree_view_set_cursor_on_cell.
func (recv *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell *CellRenderer, startEditing bool) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_focus_column := (*C.GtkTreeViewColumn)(focusColumn.ToC())

	c_focus_cell := (*C.GtkCellRenderer)(focusCell.ToC())

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_tree_view_set_cursor_on_cell((*C.GtkTreeView)(recv.native), c_path, c_focus_column, c_focus_cell, c_start_editing)

	return
}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// FocusCell is a wrapper around the C function gtk_tree_view_column_focus_cell.
func (recv *TreeViewColumn) FocusCell(cell *CellRenderer) {
	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	C.gtk_tree_view_column_focus_cell((*C.GtkTreeViewColumn)(recv.native), c_cell)

	return
}

// Unsupported : gtk_widget_new : unsupported parameter ... : varargs

// Unsupported : gtk_widget_get_clipboard : unsupported parameter selection : Blacklisted record : GdkAtom

// GetDisplay is a wrapper around the C function gtk_widget_get_display.
func (recv *Widget) GetDisplay() *gdk.Display {
	retC := C.gtk_widget_get_display((*C.GtkWidget)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRootWindow is a wrapper around the C function gtk_widget_get_root_window.
func (recv *Widget) GetRootWindow() *gdk.Window {
	retC := C.gtk_widget_get_root_window((*C.GtkWidget)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_widget_get_screen.
func (recv *Widget) GetScreen() *gdk.Screen {
	retC := C.gtk_widget_get_screen((*C.GtkWidget)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasScreen is a wrapper around the C function gtk_widget_has_screen.
func (recv *Widget) HasScreen() bool {
	retC := C.gtk_widget_has_screen((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Fullscreen is a wrapper around the C function gtk_window_fullscreen.
func (recv *Window) Fullscreen() {
	C.gtk_window_fullscreen((*C.GtkWindow)(recv.native))

	return
}

// GetScreen is a wrapper around the C function gtk_window_get_screen.
func (recv *Window) GetScreen() *gdk.Screen {
	retC := C.gtk_window_get_screen((*C.GtkWindow)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSkipPagerHint is a wrapper around the C function gtk_window_get_skip_pager_hint.
func (recv *Window) GetSkipPagerHint() bool {
	retC := C.gtk_window_get_skip_pager_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSkipTaskbarHint is a wrapper around the C function gtk_window_get_skip_taskbar_hint.
func (recv *Window) GetSkipTaskbarHint() bool {
	retC := C.gtk_window_get_skip_taskbar_hint((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetIconFromFile is a wrapper around the C function gtk_window_set_icon_from_file.
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

// SetScreen is a wrapper around the C function gtk_window_set_screen.
func (recv *Window) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(screen.ToC())

	C.gtk_window_set_screen((*C.GtkWindow)(recv.native), c_screen)

	return
}

// SetSkipPagerHint is a wrapper around the C function gtk_window_set_skip_pager_hint.
func (recv *Window) SetSkipPagerHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_pager_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetSkipTaskbarHint is a wrapper around the C function gtk_window_set_skip_taskbar_hint.
func (recv *Window) SetSkipTaskbarHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_skip_taskbar_hint((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Unfullscreen is a wrapper around the C function gtk_window_unfullscreen.
func (recv *Window) Unfullscreen() {
	C.gtk_window_unfullscreen((*C.GtkWindow)(recv.native))

	return
}

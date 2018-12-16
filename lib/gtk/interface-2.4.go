// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AddAttribute is a wrapper around the C function gtk_cell_layout_add_attribute.
func (recv *CellLayout) AddAttribute(cell *CellRenderer, attribute string, column int32) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_column := (C.gint)(column)

	C.gtk_cell_layout_add_attribute((*C.GtkCellLayout)(recv.native), c_cell, c_attribute, c_column)

	return
}

// Clear is a wrapper around the C function gtk_cell_layout_clear.
func (recv *CellLayout) Clear() {
	C.gtk_cell_layout_clear((*C.GtkCellLayout)(recv.native))

	return
}

// ClearAttributes is a wrapper around the C function gtk_cell_layout_clear_attributes.
func (recv *CellLayout) ClearAttributes(cell *CellRenderer) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_cell_layout_clear_attributes((*C.GtkCellLayout)(recv.native), c_cell)

	return
}

// PackEnd is a wrapper around the C function gtk_cell_layout_pack_end.
func (recv *CellLayout) PackEnd(cell *CellRenderer, expand bool) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	C.gtk_cell_layout_pack_end((*C.GtkCellLayout)(recv.native), c_cell, c_expand)

	return
}

// PackStart is a wrapper around the C function gtk_cell_layout_pack_start.
func (recv *CellLayout) PackStart(cell *CellRenderer, expand bool) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	C.gtk_cell_layout_pack_start((*C.GtkCellLayout)(recv.native), c_cell, c_expand)

	return
}

// Reorder is a wrapper around the C function gtk_cell_layout_reorder.
func (recv *CellLayout) Reorder(cell *CellRenderer, position int32) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_cell_layout_reorder((*C.GtkCellLayout)(recv.native), c_cell, c_position)

	return
}

// Unsupported : gtk_cell_layout_set_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_cell_layout_set_cell_data_func : unsupported parameter func : no type generator for CellLayoutDataFunc (GtkCellLayoutDataFunc) for param func

// AddFilter is a wrapper around the C function gtk_file_chooser_add_filter.
func (recv *FileChooser) AddFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_add_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// AddShortcutFolder is a wrapper around the C function gtk_file_chooser_add_shortcut_folder.
func (recv *FileChooser) AddShortcutFolder(folder string) (bool, error) {
	c_folder := C.CString(folder)
	defer C.free(unsafe.Pointer(c_folder))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_add_shortcut_folder((*C.GtkFileChooser)(recv.native), c_folder, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddShortcutFolderUri is a wrapper around the C function gtk_file_chooser_add_shortcut_folder_uri.
func (recv *FileChooser) AddShortcutFolderUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_add_shortcut_folder_uri((*C.GtkFileChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAction is a wrapper around the C function gtk_file_chooser_get_action.
func (recv *FileChooser) GetAction() FileChooserAction {
	retC := C.gtk_file_chooser_get_action((*C.GtkFileChooser)(recv.native))
	retGo := (FileChooserAction)(retC)

	return retGo
}

// GetCurrentFolder is a wrapper around the C function gtk_file_chooser_get_current_folder.
func (recv *FileChooser) GetCurrentFolder() string {
	retC := C.gtk_file_chooser_get_current_folder((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentFolderUri is a wrapper around the C function gtk_file_chooser_get_current_folder_uri.
func (recv *FileChooser) GetCurrentFolderUri() string {
	retC := C.gtk_file_chooser_get_current_folder_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetExtraWidget is a wrapper around the C function gtk_file_chooser_get_extra_widget.
func (recv *FileChooser) GetExtraWidget() *Widget {
	retC := C.gtk_file_chooser_get_extra_widget((*C.GtkFileChooser)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFilename is a wrapper around the C function gtk_file_chooser_get_filename.
func (recv *FileChooser) GetFilename() string {
	retC := C.gtk_file_chooser_get_filename((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetFilenames is a wrapper around the C function gtk_file_chooser_get_filenames.
func (recv *FileChooser) GetFilenames() *glib.SList {
	retC := C.gtk_file_chooser_get_filenames((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFilter is a wrapper around the C function gtk_file_chooser_get_filter.
func (recv *FileChooser) GetFilter() *FileFilter {
	retC := C.gtk_file_chooser_get_filter((*C.GtkFileChooser)(recv.native))
	var retGo (*FileFilter)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileFilterNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLocalOnly is a wrapper around the C function gtk_file_chooser_get_local_only.
func (recv *FileChooser) GetLocalOnly() bool {
	retC := C.gtk_file_chooser_get_local_only((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPreviewFilename is a wrapper around the C function gtk_file_chooser_get_preview_filename.
func (recv *FileChooser) GetPreviewFilename() string {
	retC := C.gtk_file_chooser_get_preview_filename((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewUri is a wrapper around the C function gtk_file_chooser_get_preview_uri.
func (recv *FileChooser) GetPreviewUri() string {
	retC := C.gtk_file_chooser_get_preview_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewWidget is a wrapper around the C function gtk_file_chooser_get_preview_widget.
func (recv *FileChooser) GetPreviewWidget() *Widget {
	retC := C.gtk_file_chooser_get_preview_widget((*C.GtkFileChooser)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetPreviewWidgetActive is a wrapper around the C function gtk_file_chooser_get_preview_widget_active.
func (recv *FileChooser) GetPreviewWidgetActive() bool {
	retC := C.gtk_file_chooser_get_preview_widget_active((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectMultiple is a wrapper around the C function gtk_file_chooser_get_select_multiple.
func (recv *FileChooser) GetSelectMultiple() bool {
	retC := C.gtk_file_chooser_get_select_multiple((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUri is a wrapper around the C function gtk_file_chooser_get_uri.
func (recv *FileChooser) GetUri() string {
	retC := C.gtk_file_chooser_get_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUris is a wrapper around the C function gtk_file_chooser_get_uris.
func (recv *FileChooser) GetUris() *glib.SList {
	retC := C.gtk_file_chooser_get_uris((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListFilters is a wrapper around the C function gtk_file_chooser_list_filters.
func (recv *FileChooser) ListFilters() *glib.SList {
	retC := C.gtk_file_chooser_list_filters((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListShortcutFolderUris is a wrapper around the C function gtk_file_chooser_list_shortcut_folder_uris.
func (recv *FileChooser) ListShortcutFolderUris() *glib.SList {
	retC := C.gtk_file_chooser_list_shortcut_folder_uris((*C.GtkFileChooser)(recv.native))
	var retGo (*glib.SList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.SListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ListShortcutFolders is a wrapper around the C function gtk_file_chooser_list_shortcut_folders.
func (recv *FileChooser) ListShortcutFolders() *glib.SList {
	retC := C.gtk_file_chooser_list_shortcut_folders((*C.GtkFileChooser)(recv.native))
	var retGo (*glib.SList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.SListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// RemoveFilter is a wrapper around the C function gtk_file_chooser_remove_filter.
func (recv *FileChooser) RemoveFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_remove_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// RemoveShortcutFolder is a wrapper around the C function gtk_file_chooser_remove_shortcut_folder.
func (recv *FileChooser) RemoveShortcutFolder(folder string) (bool, error) {
	c_folder := C.CString(folder)
	defer C.free(unsafe.Pointer(c_folder))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_remove_shortcut_folder((*C.GtkFileChooser)(recv.native), c_folder, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveShortcutFolderUri is a wrapper around the C function gtk_file_chooser_remove_shortcut_folder_uri.
func (recv *FileChooser) RemoveShortcutFolderUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_remove_shortcut_folder_uri((*C.GtkFileChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SelectAll is a wrapper around the C function gtk_file_chooser_select_all.
func (recv *FileChooser) SelectAll() {
	C.gtk_file_chooser_select_all((*C.GtkFileChooser)(recv.native))

	return
}

// SelectFilename is a wrapper around the C function gtk_file_chooser_select_filename.
func (recv *FileChooser) SelectFilename(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_select_filename((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// SelectUri is a wrapper around the C function gtk_file_chooser_select_uri.
func (recv *FileChooser) SelectUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_select_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// SetAction is a wrapper around the C function gtk_file_chooser_set_action.
func (recv *FileChooser) SetAction(action FileChooserAction) {
	c_action := (C.GtkFileChooserAction)(action)

	C.gtk_file_chooser_set_action((*C.GtkFileChooser)(recv.native), c_action)

	return
}

// SetCurrentFolder is a wrapper around the C function gtk_file_chooser_set_current_folder.
func (recv *FileChooser) SetCurrentFolder(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_set_current_folder((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// SetCurrentFolderUri is a wrapper around the C function gtk_file_chooser_set_current_folder_uri.
func (recv *FileChooser) SetCurrentFolderUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_set_current_folder_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// SetCurrentName is a wrapper around the C function gtk_file_chooser_set_current_name.
func (recv *FileChooser) SetCurrentName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_file_chooser_set_current_name((*C.GtkFileChooser)(recv.native), c_name)

	return
}

// SetExtraWidget is a wrapper around the C function gtk_file_chooser_set_extra_widget.
func (recv *FileChooser) SetExtraWidget(extraWidget *Widget) {
	c_extra_widget := (*C.GtkWidget)(C.NULL)
	if extraWidget != nil {
		c_extra_widget = (*C.GtkWidget)(extraWidget.ToC())
	}

	C.gtk_file_chooser_set_extra_widget((*C.GtkFileChooser)(recv.native), c_extra_widget)

	return
}

// SetFilename is a wrapper around the C function gtk_file_chooser_set_filename.
func (recv *FileChooser) SetFilename(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_set_filename((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// SetFilter is a wrapper around the C function gtk_file_chooser_set_filter.
func (recv *FileChooser) SetFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_set_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// SetLocalOnly is a wrapper around the C function gtk_file_chooser_set_local_only.
func (recv *FileChooser) SetLocalOnly(localOnly bool) {
	c_local_only :=
		boolToGboolean(localOnly)

	C.gtk_file_chooser_set_local_only((*C.GtkFileChooser)(recv.native), c_local_only)

	return
}

// SetPreviewWidget is a wrapper around the C function gtk_file_chooser_set_preview_widget.
func (recv *FileChooser) SetPreviewWidget(previewWidget *Widget) {
	c_preview_widget := (*C.GtkWidget)(C.NULL)
	if previewWidget != nil {
		c_preview_widget = (*C.GtkWidget)(previewWidget.ToC())
	}

	C.gtk_file_chooser_set_preview_widget((*C.GtkFileChooser)(recv.native), c_preview_widget)

	return
}

// SetPreviewWidgetActive is a wrapper around the C function gtk_file_chooser_set_preview_widget_active.
func (recv *FileChooser) SetPreviewWidgetActive(active bool) {
	c_active :=
		boolToGboolean(active)

	C.gtk_file_chooser_set_preview_widget_active((*C.GtkFileChooser)(recv.native), c_active)

	return
}

// SetSelectMultiple is a wrapper around the C function gtk_file_chooser_set_select_multiple.
func (recv *FileChooser) SetSelectMultiple(selectMultiple bool) {
	c_select_multiple :=
		boolToGboolean(selectMultiple)

	C.gtk_file_chooser_set_select_multiple((*C.GtkFileChooser)(recv.native), c_select_multiple)

	return
}

// SetUri is a wrapper around the C function gtk_file_chooser_set_uri.
func (recv *FileChooser) SetUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_set_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// SetUsePreviewLabel is a wrapper around the C function gtk_file_chooser_set_use_preview_label.
func (recv *FileChooser) SetUsePreviewLabel(useLabel bool) {
	c_use_label :=
		boolToGboolean(useLabel)

	C.gtk_file_chooser_set_use_preview_label((*C.GtkFileChooser)(recv.native), c_use_label)

	return
}

// UnselectAll is a wrapper around the C function gtk_file_chooser_unselect_all.
func (recv *FileChooser) UnselectAll() {
	C.gtk_file_chooser_unselect_all((*C.GtkFileChooser)(recv.native))

	return
}

// UnselectFilename is a wrapper around the C function gtk_file_chooser_unselect_filename.
func (recv *FileChooser) UnselectFilename(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_file_chooser_unselect_filename((*C.GtkFileChooser)(recv.native), c_filename)

	return
}

// UnselectUri is a wrapper around the C function gtk_file_chooser_unselect_uri.
func (recv *FileChooser) UnselectUri(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.gtk_file_chooser_unselect_uri((*C.GtkFileChooser)(recv.native), c_uri)

	return
}

// FilterNew is a wrapper around the C function gtk_tree_model_filter_new.
func (recv *TreeModel) FilterNew(root *TreePath) *TreeModel {
	c_root := (*C.GtkTreePath)(C.NULL)
	if root != nil {
		c_root = (*C.GtkTreePath)(root.ToC())
	}

	retC := C.gtk_tree_model_filter_new((*C.GtkTreeModel)(recv.native), c_root)
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

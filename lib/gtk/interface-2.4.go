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

// Adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the parameter on @cell to be set from the value. So for
// example if column 2 of the model contains strings, you could have the
// “text” attribute of a #GtkCellRendererText get its values from column 2.
/*

C function

gtk_cell_layout_add_attribute
*/
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

// Unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
/*

C function

gtk_cell_layout_clear
*/
func (recv *CellLayout) Clear() {
	C.gtk_cell_layout_clear((*C.GtkCellLayout)(recv.native))

	return
}

// Clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
/*

C function

gtk_cell_layout_clear_attributes
*/
func (recv *CellLayout) ClearAttributes(cell *CellRenderer) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_cell_layout_clear_attributes((*C.GtkCellLayout)(recv.native), c_cell)

	return
}

// Adds the @cell to the end of @cell_layout. If @expand is %FALSE, then the
// @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
/*

C function

gtk_cell_layout_pack_end
*/
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

// Packs the @cell into the beginning of @cell_layout. If @expand is %FALSE,
// then the @cell is allocated no more space than it needs. Any unused space
// is divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
/*

C function

gtk_cell_layout_pack_start
*/
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

// Re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout
// for this to function properly.
/*

C function

gtk_cell_layout_reorder
*/
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

// Adds @filter to the list of filters that the user can select between.
// When a filter is selected, only files that are passed by that
// filter are displayed.
//
// Note that the @chooser takes ownership of the filter, so you have to
// ref and sink it if you want to keep a reference.
/*

C function

gtk_file_chooser_add_filter
*/
func (recv *FileChooser) AddFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_add_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// Adds a folder to be displayed with the shortcut folders in a file chooser.
// Note that shortcut folders do not get saved, as they are provided by the
// application.  For example, you can use this to add a
// “/usr/share/mydrawprogram/Clipart” folder to the volume list.
/*

C function

gtk_file_chooser_add_shortcut_folder
*/
func (recv *FileChooser) AddShortcutFolder(folder string) (bool, error) {
	c_folder := C.CString(folder)
	defer C.free(unsafe.Pointer(c_folder))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_add_shortcut_folder((*C.GtkFileChooser)(recv.native), c_folder, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Adds a folder URI to be displayed with the shortcut folders in a file
// chooser.  Note that shortcut folders do not get saved, as they are provided
// by the application.  For example, you can use this to add a
// “file:///usr/share/mydrawprogram/Clipart” folder to the volume list.
/*

C function

gtk_file_chooser_add_shortcut_folder_uri
*/
func (recv *FileChooser) AddShortcutFolderUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_add_shortcut_folder_uri((*C.GtkFileChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the type of operation that the file chooser is performing; see
// gtk_file_chooser_set_action().
/*

C function

gtk_file_chooser_get_action
*/
func (recv *FileChooser) GetAction() FileChooserAction {
	retC := C.gtk_file_chooser_get_action((*C.GtkFileChooser)(recv.native))
	retGo := (FileChooserAction)(retC)

	return retGo
}

// Gets the current folder of @chooser as a local filename.
// See gtk_file_chooser_set_current_folder().
//
// Note that this is the folder that the file chooser is currently displaying
// (e.g. "/home/username/Documents"), which is not the same
// as the currently-selected folder if the chooser is in
// %GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER mode
// (e.g. "/home/username/Documents/selected-folder/".  To get the
// currently-selected folder in that mode, use gtk_file_chooser_get_uri() as the
// usual way to get the selection.
/*

C function

gtk_file_chooser_get_current_folder
*/
func (recv *FileChooser) GetCurrentFolder() string {
	retC := C.gtk_file_chooser_get_current_folder((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the current folder of @chooser as an URI.
// See gtk_file_chooser_set_current_folder_uri().
//
// Note that this is the folder that the file chooser is currently displaying
// (e.g. "file:///home/username/Documents"), which is not the same
// as the currently-selected folder if the chooser is in
// %GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER mode
// (e.g. "file:///home/username/Documents/selected-folder/".  To get the
// currently-selected folder in that mode, use gtk_file_chooser_get_uri() as the
// usual way to get the selection.
/*

C function

gtk_file_chooser_get_current_folder_uri
*/
func (recv *FileChooser) GetCurrentFolderUri() string {
	retC := C.gtk_file_chooser_get_current_folder_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the current extra widget; see
// gtk_file_chooser_set_extra_widget().
/*

C function

gtk_file_chooser_get_extra_widget
*/
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

// Gets the filename for the currently selected file in
// the file selector. The filename is returned as an absolute path. If
// multiple files are selected, one of the filenames will be returned at
// random.
//
// If the file chooser is in folder mode, this function returns the selected
// folder.
/*

C function

gtk_file_chooser_get_filename
*/
func (recv *FileChooser) GetFilename() string {
	retC := C.gtk_file_chooser_get_filename((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Lists all the selected files and subfolders in the current folder of
// @chooser. The returned names are full absolute paths. If files in the current
// folder cannot be represented as local filenames they will be ignored. (See
// gtk_file_chooser_get_uris())
/*

C function

gtk_file_chooser_get_filenames
*/
func (recv *FileChooser) GetFilenames() *glib.SList {
	retC := C.gtk_file_chooser_get_filenames((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the current filter; see gtk_file_chooser_set_filter().
/*

C function

gtk_file_chooser_get_filter
*/
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

// Gets whether only local files can be selected in the
// file selector. See gtk_file_chooser_set_local_only()
/*

C function

gtk_file_chooser_get_local_only
*/
func (recv *FileChooser) GetLocalOnly() bool {
	retC := C.gtk_file_chooser_get_local_only((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the filename that should be previewed in a custom preview
// widget. See gtk_file_chooser_set_preview_widget().
/*

C function

gtk_file_chooser_get_preview_filename
*/
func (recv *FileChooser) GetPreviewFilename() string {
	retC := C.gtk_file_chooser_get_preview_filename((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the URI that should be previewed in a custom preview
// widget. See gtk_file_chooser_set_preview_widget().
/*

C function

gtk_file_chooser_get_preview_uri
*/
func (recv *FileChooser) GetPreviewUri() string {
	retC := C.gtk_file_chooser_get_preview_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the current preview widget; see
// gtk_file_chooser_set_preview_widget().
/*

C function

gtk_file_chooser_get_preview_widget
*/
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

// Gets whether the preview widget set by gtk_file_chooser_set_preview_widget()
// should be shown for the current filename. See
// gtk_file_chooser_set_preview_widget_active().
/*

C function

gtk_file_chooser_get_preview_widget_active
*/
func (recv *FileChooser) GetPreviewWidgetActive() bool {
	retC := C.gtk_file_chooser_get_preview_widget_active((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether multiple files can be selected in the file
// selector. See gtk_file_chooser_set_select_multiple().
/*

C function

gtk_file_chooser_get_select_multiple
*/
func (recv *FileChooser) GetSelectMultiple() bool {
	retC := C.gtk_file_chooser_get_select_multiple((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the URI for the currently selected file in
// the file selector. If multiple files are selected,
// one of the filenames will be returned at random.
//
// If the file chooser is in folder mode, this function returns the selected
// folder.
/*

C function

gtk_file_chooser_get_uri
*/
func (recv *FileChooser) GetUri() string {
	retC := C.gtk_file_chooser_get_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Lists all the selected files and subfolders in the current folder of
// @chooser. The returned names are full absolute URIs.
/*

C function

gtk_file_chooser_get_uris
*/
func (recv *FileChooser) GetUris() *glib.SList {
	retC := C.gtk_file_chooser_get_uris((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lists the current set of user-selectable filters; see
// gtk_file_chooser_add_filter(), gtk_file_chooser_remove_filter().
/*

C function

gtk_file_chooser_list_filters
*/
func (recv *FileChooser) ListFilters() *glib.SList {
	retC := C.gtk_file_chooser_list_filters((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Queries the list of shortcut folders in the file chooser, as set by
// gtk_file_chooser_add_shortcut_folder_uri().
/*

C function

gtk_file_chooser_list_shortcut_folder_uris
*/
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

// Queries the list of shortcut folders in the file chooser, as set by
// gtk_file_chooser_add_shortcut_folder().
/*

C function

gtk_file_chooser_list_shortcut_folders
*/
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

// Removes @filter from the list of filters that the user can select between.
/*

C function

gtk_file_chooser_remove_filter
*/
func (recv *FileChooser) RemoveFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_remove_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// Removes a folder from a file chooser’s list of shortcut folders.
/*

C function

gtk_file_chooser_remove_shortcut_folder
*/
func (recv *FileChooser) RemoveShortcutFolder(folder string) (bool, error) {
	c_folder := C.CString(folder)
	defer C.free(unsafe.Pointer(c_folder))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_remove_shortcut_folder((*C.GtkFileChooser)(recv.native), c_folder, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Removes a folder URI from a file chooser’s list of shortcut folders.
/*

C function

gtk_file_chooser_remove_shortcut_folder_uri
*/
func (recv *FileChooser) RemoveShortcutFolderUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_remove_shortcut_folder_uri((*C.GtkFileChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Selects all the files in the current folder of a file chooser.
/*

C function

gtk_file_chooser_select_all
*/
func (recv *FileChooser) SelectAll() {
	C.gtk_file_chooser_select_all((*C.GtkFileChooser)(recv.native))

	return
}

// Selects a filename. If the file name isn’t in the current
// folder of @chooser, then the current folder of @chooser will
// be changed to the folder containing @filename.
/*

C function

gtk_file_chooser_select_filename
*/
func (recv *FileChooser) SelectFilename(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_select_filename((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// Selects the file to by @uri. If the URI doesn’t refer to a
// file in the current folder of @chooser, then the current folder of
// @chooser will be changed to the folder containing @filename.
/*

C function

gtk_file_chooser_select_uri
*/
func (recv *FileChooser) SelectUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_select_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the type of operation that the chooser is performing; the
// user interface is adapted to suit the selected action. For example,
// an option to create a new folder might be shown if the action is
// %GTK_FILE_CHOOSER_ACTION_SAVE but not if the action is
// %GTK_FILE_CHOOSER_ACTION_OPEN.
/*

C function

gtk_file_chooser_set_action
*/
func (recv *FileChooser) SetAction(action FileChooserAction) {
	c_action := (C.GtkFileChooserAction)(action)

	C.gtk_file_chooser_set_action((*C.GtkFileChooser)(recv.native), c_action)

	return
}

// Sets the current folder for @chooser from a local filename.
// The user will be shown the full contents of the current folder,
// plus user interface elements for navigating to other folders.
//
// In general, you should not use this function.  See the
// [section on setting up a file chooser dialog][gtkfilechooserdialog-setting-up]
// for the rationale behind this.
/*

C function

gtk_file_chooser_set_current_folder
*/
func (recv *FileChooser) SetCurrentFolder(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_set_current_folder((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the current folder for @chooser from an URI.
// The user will be shown the full contents of the current folder,
// plus user interface elements for navigating to other folders.
//
// In general, you should not use this function.  See the
// [section on setting up a file chooser dialog][gtkfilechooserdialog-setting-up]
// for the rationale behind this.
/*

C function

gtk_file_chooser_set_current_folder_uri
*/
func (recv *FileChooser) SetCurrentFolderUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_set_current_folder_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the current name in the file selector, as if entered
// by the user. Note that the name passed in here is a UTF-8
// string rather than a filename. This function is meant for
// such uses as a suggested name in a “Save As...” dialog.  You can
// pass “Untitled.doc” or a similarly suitable suggestion for the @name.
//
// If you want to preselect a particular existing file, you should use
// gtk_file_chooser_set_filename() or gtk_file_chooser_set_uri() instead.
// Please see the documentation for those functions for an example of using
// gtk_file_chooser_set_current_name() as well.
/*

C function

gtk_file_chooser_set_current_name
*/
func (recv *FileChooser) SetCurrentName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_file_chooser_set_current_name((*C.GtkFileChooser)(recv.native), c_name)

	return
}

// Sets an application-supplied widget to provide extra options to the user.
/*

C function

gtk_file_chooser_set_extra_widget
*/
func (recv *FileChooser) SetExtraWidget(extraWidget *Widget) {
	c_extra_widget := (*C.GtkWidget)(C.NULL)
	if extraWidget != nil {
		c_extra_widget = (*C.GtkWidget)(extraWidget.ToC())
	}

	C.gtk_file_chooser_set_extra_widget((*C.GtkFileChooser)(recv.native), c_extra_widget)

	return
}

// Sets @filename as the current filename for the file chooser, by changing to
// the file’s parent folder and actually selecting the file in list; all other
// files will be unselected.  If the @chooser is in
// %GTK_FILE_CHOOSER_ACTION_SAVE mode, the file’s base name will also appear in
// the dialog’s file name entry.
//
// Note that the file must exist, or nothing will be done except
// for the directory change.
//
// You should use this function only when implementing a save
// dialog for which you already have a file name to which
// the user may save.  For example, when the user opens an existing file and
// then does Save As... to save a copy or
// a modified version.  If you don’t have a file name already — for
// example, if the user just created a new file and is saving it for the first
// time, do not call this function.  Instead, use something similar to this:
// |[<!-- language="C" -->
// if (document_is_new)
// {
// the user just created a new document
// gtk_file_chooser_set_current_name (chooser, "Untitled document");
// }
// else
// {
// the user edited an existing document
// gtk_file_chooser_set_filename (chooser, existing_filename);
// }
// ]|
//
// In the first case, the file chooser will present the user with useful suggestions
// as to where to save his new file.  In the second case, the file’s existing location
// is already known, so the file chooser will use it.
/*

C function

gtk_file_chooser_set_filename
*/
func (recv *FileChooser) SetFilename(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_set_filename((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the current filter; only the files that pass the
// filter will be displayed. If the user-selectable list of filters
// is non-empty, then the filter should be one of the filters
// in that list. Setting the current filter when the list of
// filters is empty is useful if you want to restrict the displayed
// set of files without letting the user change it.
/*

C function

gtk_file_chooser_set_filter
*/
func (recv *FileChooser) SetFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_set_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// Sets whether only local files can be selected in the
// file selector. If @local_only is %TRUE (the default),
// then the selected file or files are guaranteed to be
// accessible through the operating systems native file
// system and therefore the application only
// needs to worry about the filename functions in
// #GtkFileChooser, like gtk_file_chooser_get_filename(),
// rather than the URI functions like
// gtk_file_chooser_get_uri(),
//
// On some systems non-native files may still be
// available using the native filesystem via a userspace
// filesystem (FUSE).
/*

C function

gtk_file_chooser_set_local_only
*/
func (recv *FileChooser) SetLocalOnly(localOnly bool) {
	c_local_only :=
		boolToGboolean(localOnly)

	C.gtk_file_chooser_set_local_only((*C.GtkFileChooser)(recv.native), c_local_only)

	return
}

// Sets an application-supplied widget to use to display a custom preview
// of the currently selected file. To implement a preview, after setting the
// preview widget, you connect to the #GtkFileChooser::update-preview
// signal, and call gtk_file_chooser_get_preview_filename() or
// gtk_file_chooser_get_preview_uri() on each change. If you can
// display a preview of the new file, update your widget and
// set the preview active using gtk_file_chooser_set_preview_widget_active().
// Otherwise, set the preview inactive.
//
// When there is no application-supplied preview widget, or the
// application-supplied preview widget is not active, the file chooser
// will display no preview at all.
/*

C function

gtk_file_chooser_set_preview_widget
*/
func (recv *FileChooser) SetPreviewWidget(previewWidget *Widget) {
	c_preview_widget := (*C.GtkWidget)(C.NULL)
	if previewWidget != nil {
		c_preview_widget = (*C.GtkWidget)(previewWidget.ToC())
	}

	C.gtk_file_chooser_set_preview_widget((*C.GtkFileChooser)(recv.native), c_preview_widget)

	return
}

// Sets whether the preview widget set by
// gtk_file_chooser_set_preview_widget() should be shown for the
// current filename. When @active is set to false, the file chooser
// may display an internally generated preview of the current file
// or it may display no preview at all. See
// gtk_file_chooser_set_preview_widget() for more details.
/*

C function

gtk_file_chooser_set_preview_widget_active
*/
func (recv *FileChooser) SetPreviewWidgetActive(active bool) {
	c_active :=
		boolToGboolean(active)

	C.gtk_file_chooser_set_preview_widget_active((*C.GtkFileChooser)(recv.native), c_active)

	return
}

// Sets whether multiple files can be selected in the file selector.  This is
// only relevant if the action is set to be %GTK_FILE_CHOOSER_ACTION_OPEN or
// %GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
/*

C function

gtk_file_chooser_set_select_multiple
*/
func (recv *FileChooser) SetSelectMultiple(selectMultiple bool) {
	c_select_multiple :=
		boolToGboolean(selectMultiple)

	C.gtk_file_chooser_set_select_multiple((*C.GtkFileChooser)(recv.native), c_select_multiple)

	return
}

// Sets the file referred to by @uri as the current file for the file chooser,
// by changing to the URI’s parent folder and actually selecting the URI in the
// list.  If the @chooser is %GTK_FILE_CHOOSER_ACTION_SAVE mode, the URI’s base
// name will also appear in the dialog’s file name entry.
//
// Note that the URI must exist, or nothing will be done except for the
// directory change.
//
// You should use this function only when implementing a save
// dialog for which you already have a file name to which
// the user may save.  For example, when the user opens an existing file and then
// does Save As... to save a copy or a
// modified version.  If you don’t have a file name already — for example,
// if the user just created a new file and is saving it for the first time, do
// not call this function.  Instead, use something similar to this:
// |[<!-- language="C" -->
// if (document_is_new)
// {
// the user just created a new document
// gtk_file_chooser_set_current_name (chooser, "Untitled document");
// }
// else
// {
// the user edited an existing document
// gtk_file_chooser_set_uri (chooser, existing_uri);
// }
// ]|
//
//
// In the first case, the file chooser will present the user with useful suggestions
// as to where to save his new file.  In the second case, the file’s existing location
// is already known, so the file chooser will use it.
/*

C function

gtk_file_chooser_set_uri
*/
func (recv *FileChooser) SetUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_set_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the file chooser should display a stock label with the name of
// the file that is being previewed; the default is %TRUE.  Applications that
// want to draw the whole preview area themselves should set this to %FALSE and
// display the name themselves in their preview widget.
//
// See also: gtk_file_chooser_set_preview_widget()
/*

C function

gtk_file_chooser_set_use_preview_label
*/
func (recv *FileChooser) SetUsePreviewLabel(useLabel bool) {
	c_use_label :=
		boolToGboolean(useLabel)

	C.gtk_file_chooser_set_use_preview_label((*C.GtkFileChooser)(recv.native), c_use_label)

	return
}

// Unselects all the files in the current folder of a file chooser.
/*

C function

gtk_file_chooser_unselect_all
*/
func (recv *FileChooser) UnselectAll() {
	C.gtk_file_chooser_unselect_all((*C.GtkFileChooser)(recv.native))

	return
}

// Unselects a currently selected filename. If the filename
// is not in the current directory, does not exist, or
// is otherwise not currently selected, does nothing.
/*

C function

gtk_file_chooser_unselect_filename
*/
func (recv *FileChooser) UnselectFilename(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_file_chooser_unselect_filename((*C.GtkFileChooser)(recv.native), c_filename)

	return
}

// Unselects the file referred to by @uri. If the file
// is not in the current directory, does not exist, or
// is otherwise not currently selected, does nothing.
/*

C function

gtk_file_chooser_unselect_uri
*/
func (recv *FileChooser) UnselectUri(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.gtk_file_chooser_unselect_uri((*C.GtkFileChooser)(recv.native), c_uri)

	return
}

// Creates a new #GtkTreeModel, with @child_model as the child_model
// and @root as the virtual root.
/*

C function

gtk_tree_model_filter_new
*/
func (recv *TreeModel) FilterNew(root *TreePath) *TreeModel {
	c_root := (*C.GtkTreePath)(C.NULL)
	if root != nil {
		c_root = (*C.GtkTreePath)(root.ToC())
	}

	retC := C.gtk_tree_model_filter_new((*C.GtkTreeModel)(recv.native), c_root)
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

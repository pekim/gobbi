// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
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

// Gets the current folder of @chooser as #GFile.
// See gtk_file_chooser_get_current_folder_uri().
/*

C function : gtk_file_chooser_get_current_folder_file
*/
func (recv *FileChooser) GetCurrentFolderFile() *gio.File {
	retC := C.gtk_file_chooser_get_current_folder_file((*C.GtkFileChooser)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GFile for the currently selected file in
// the file selector. If multiple files are selected,
// one of the files will be returned at random.
//
// If the file chooser is in folder mode, this function returns the selected
// folder.
/*

C function : gtk_file_chooser_get_file
*/
func (recv *FileChooser) GetFile() *gio.File {
	retC := C.gtk_file_chooser_get_file((*C.GtkFileChooser)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lists all the selected files and subfolders in the current folder of @chooser
// as #GFile. An internal function, see gtk_file_chooser_get_uris().
/*

C function : gtk_file_chooser_get_files
*/
func (recv *FileChooser) GetFiles() *glib.SList {
	retC := C.gtk_file_chooser_get_files((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GFile that should be previewed in a custom preview
// Internal function, see gtk_file_chooser_get_preview_uri().
/*

C function : gtk_file_chooser_get_preview_file
*/
func (recv *FileChooser) GetPreviewFile() *gio.File {
	retC := C.gtk_file_chooser_get_preview_file((*C.GtkFileChooser)(recv.native))
	var retGo (*gio.File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Selects the file referred to by @file. An internal function. See
// _gtk_file_chooser_select_uri().
/*

C function : gtk_file_chooser_select_file
*/
func (recv *FileChooser) SelectFile(file *gio.File) (bool, error) {
	c_file := (*C.GFile)(file.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_select_file((*C.GtkFileChooser)(recv.native), c_file, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the current folder for @chooser from a #GFile.
// Internal function, see gtk_file_chooser_set_current_folder_uri().
/*

C function : gtk_file_chooser_set_current_folder_file
*/
func (recv *FileChooser) SetCurrentFolderFile(file *gio.File) (bool, error) {
	c_file := (*C.GFile)(file.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_set_current_folder_file((*C.GtkFileChooser)(recv.native), c_file, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @file as the current filename for the file chooser, by changing
// to the file’s parent folder and actually selecting the file in list.  If
// the @chooser is in %GTK_FILE_CHOOSER_ACTION_SAVE mode, the file’s base name
// will also appear in the dialog’s file name entry.
//
// If the file name isn’t in the current folder of @chooser, then the current
// folder of @chooser will be changed to the folder containing @filename. This
// is equivalent to a sequence of gtk_file_chooser_unselect_all() followed by
// gtk_file_chooser_select_filename().
//
// Note that the file must exist, or nothing will be done except
// for the directory change.
//
// If you are implementing a save dialog,
// you should use this function if you already have a file name to which the
// user may save; for example, when the user opens an existing file and then
// does Save As...  If you don’t have
// a file name already — for example, if the user just created a new
// file and is saving it for the first time, do not call this function.
// Instead, use something similar to this:
// |[<!-- language="C" -->
// if (document_is_new)
// {
// the user just created a new document
// gtk_file_chooser_set_current_folder_file (chooser, default_file_for_saving);
// gtk_file_chooser_set_current_name (chooser, "Untitled document");
// }
// else
// {
// the user edited an existing document
// gtk_file_chooser_set_file (chooser, existing_file);
// }
// ]|
/*

C function : gtk_file_chooser_set_file
*/
func (recv *FileChooser) SetFile(file *gio.File) (bool, error) {
	c_file := (*C.GFile)(file.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_set_file((*C.GtkFileChooser)(recv.native), c_file, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unselects the file referred to by @file. If the file is not in the current
// directory, does not exist, or is otherwise not currently selected, does nothing.
/*

C function : gtk_file_chooser_unselect_file
*/
func (recv *FileChooser) UnselectFile(file *gio.File) {
	c_file := (*C.GFile)(file.ToC())

	C.gtk_file_chooser_unselect_file((*C.GtkFileChooser)(recv.native), c_file)

	return
}

// Retrieves the icon size for the tool shell. Tool items must not call this
// function directly, but rely on gtk_tool_item_get_icon_size() instead.
/*

C function : gtk_tool_shell_get_icon_size
*/
func (recv *ToolShell) GetIconSize() IconSize {
	retC := C.gtk_tool_shell_get_icon_size((*C.GtkToolShell)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// Retrieves the current orientation for the tool shell. Tool items must not
// call this function directly, but rely on gtk_tool_item_get_orientation()
// instead.
/*

C function : gtk_tool_shell_get_orientation
*/
func (recv *ToolShell) GetOrientation() Orientation {
	retC := C.gtk_tool_shell_get_orientation((*C.GtkToolShell)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// Returns the relief style of buttons on @shell. Tool items must not call this
// function directly, but rely on gtk_tool_item_get_relief_style() instead.
/*

C function : gtk_tool_shell_get_relief_style
*/
func (recv *ToolShell) GetReliefStyle() ReliefStyle {
	retC := C.gtk_tool_shell_get_relief_style((*C.GtkToolShell)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// Retrieves whether the tool shell has text, icons, or both. Tool items must
// not call this function directly, but rely on gtk_tool_item_get_toolbar_style()
// instead.
/*

C function : gtk_tool_shell_get_style
*/
func (recv *ToolShell) GetStyle() ToolbarStyle {
	retC := C.gtk_tool_shell_get_style((*C.GtkToolShell)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// Calling this function signals the tool shell that the overflow menu item for
// tool items have changed. If there is an overflow menu and if it is visible
// when this function it called, the menu will be rebuilt.
//
// Tool items must not call this function directly, but rely on
// gtk_tool_item_rebuild_menu() instead.
/*

C function : gtk_tool_shell_rebuild_menu
*/
func (recv *ToolShell) RebuildMenu() {
	C.gtk_tool_shell_rebuild_menu((*C.GtkToolShell)(recv.native))

	return
}

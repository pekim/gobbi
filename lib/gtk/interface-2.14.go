// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// Unsupported : gtk_file_chooser_get_current_folder_file : no return generator

// Unsupported : gtk_file_chooser_get_file : no return generator

// GetFiles is a wrapper around the C function gtk_file_chooser_get_files.
func (recv *FileChooser) GetFiles() *glib.SList {
	retC := C.gtk_file_chooser_get_files((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_file_chooser_get_preview_file : no return generator

// Unsupported : gtk_file_chooser_select_file : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_file_chooser_set_current_folder_file : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_file_chooser_set_file : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_file_chooser_unselect_file : unsupported parameter file : no type generator for Gio.File, GFile*

// Unsupported : gtk_tool_shell_get_icon_size : no return generator

// GetOrientation is a wrapper around the C function gtk_tool_shell_get_orientation.
func (recv *ToolShell) GetOrientation() Orientation {
	retC := C.gtk_tool_shell_get_orientation((*C.GtkToolShell)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// GetReliefStyle is a wrapper around the C function gtk_tool_shell_get_relief_style.
func (recv *ToolShell) GetReliefStyle() ReliefStyle {
	retC := C.gtk_tool_shell_get_relief_style((*C.GtkToolShell)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetStyle is a wrapper around the C function gtk_tool_shell_get_style.
func (recv *ToolShell) GetStyle() ToolbarStyle {
	retC := C.gtk_tool_shell_get_style((*C.GtkToolShell)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// RebuildMenu is a wrapper around the C function gtk_tool_shell_rebuild_menu.
func (recv *ToolShell) RebuildMenu() {
	C.gtk_tool_shell_rebuild_menu((*C.GtkToolShell)(recv.native))

	return
}

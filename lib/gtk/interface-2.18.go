// This is a generated file - DO NOT EDIT
// +build gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "C"

// GetCreateFolders is a wrapper around the C function gtk_file_chooser_get_create_folders.
func (recv *FileChooser) GetCreateFolders() bool {
	retC := C.gtk_file_chooser_get_create_folders((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetCreateFolders is a wrapper around the C function gtk_file_chooser_set_create_folders.
func (recv *FileChooser) SetCreateFolders(createFolders bool) {
	c_create_folders :=
		boolToGboolean(createFolders)

	C.gtk_file_chooser_set_create_folders((*C.GtkFileChooser)(recv.native), c_create_folders)

	return
}

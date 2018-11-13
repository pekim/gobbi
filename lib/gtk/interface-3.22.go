// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_file_chooser_add_choice : unsupported parameter options :

// Gets the currently selected option in the 'choice' with the given ID.
/*

C function : gtk_file_chooser_get_choice
*/
func (recv *FileChooser) GetChoice(id string) string {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	retC := C.gtk_file_chooser_get_choice((*C.GtkFileChooser)(recv.native), c_id)
	retGo := C.GoString(retC)

	return retGo
}

// Removes a 'choice' that has been added with gtk_file_chooser_add_choice().
/*

C function : gtk_file_chooser_remove_choice
*/
func (recv *FileChooser) RemoveChoice(id string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	C.gtk_file_chooser_remove_choice((*C.GtkFileChooser)(recv.native), c_id)

	return
}

// Selects an option in a 'choice' that has been added with
// gtk_file_chooser_add_choice(). For a boolean choice, the
// possible options are "true" and "false".
/*

C function : gtk_file_chooser_set_choice
*/
func (recv *FileChooser) SetChoice(id string, option string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.gtk_file_chooser_set_choice((*C.GtkFileChooser)(recv.native), c_id, c_option)

	return
}

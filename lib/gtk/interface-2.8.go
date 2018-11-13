// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported signal 'confirm-overwrite' for FileChooser : return value FileChooserConfirmation :

// Queries whether a file chooser is set to confirm for overwriting when the user
// types a file name that already exists.
/*

C function

gtk_file_chooser_get_do_overwrite_confirmation
*/
func (recv *FileChooser) GetDoOverwriteConfirmation() bool {
	retC := C.gtk_file_chooser_get_do_overwrite_confirmation((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether a file chooser in %GTK_FILE_CHOOSER_ACTION_SAVE mode will present
// a confirmation dialog if the user types a file name that already exists.  This
// is %FALSE by default.
//
// If set to %TRUE, the @chooser will emit the
// #GtkFileChooser::confirm-overwrite signal when appropriate.
//
// If all you need is the stock confirmation dialog, set this property to %TRUE.
// You can override the way confirmation is done by actually handling the
// #GtkFileChooser::confirm-overwrite signal; please refer to its documentation
// for the details.
/*

C function

gtk_file_chooser_set_do_overwrite_confirmation
*/
func (recv *FileChooser) SetDoOverwriteConfirmation(doOverwriteConfirmation bool) {
	c_do_overwrite_confirmation :=
		boolToGboolean(doOverwriteConfirmation)

	C.gtk_file_chooser_set_do_overwrite_confirmation((*C.GtkFileChooser)(recv.native), c_do_overwrite_confirmation)

	return
}

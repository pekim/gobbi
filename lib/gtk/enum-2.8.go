// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Used as a return value of handlers for the
// #GtkFileChooser::confirm-overwrite signal of a #GtkFileChooser. This
// value determines whether the file chooser will present the stock
// confirmation dialog, accept the user’s choice of a filename, or
// let the user choose another filename.
type FileChooserConfirmation C.GtkFileChooserConfirmation

const (
	// The file chooser will present
	// its stock dialog to confirm about overwriting an existing file.
	GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM FileChooserConfirmation = 0

	// The file chooser will
	// terminate and accept the user’s choice of a file name.
	GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME FileChooserConfirmation = 1

	// The file chooser will
	// continue running, so as to let the user select another file name.
	GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN FileChooserConfirmation = 2
)

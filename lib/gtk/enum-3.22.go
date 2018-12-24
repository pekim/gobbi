// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type PadActionType C.GtkPadActionType

const (
	GTK_PAD_ACTION_BUTTON PadActionType = 0
	GTK_PAD_ACTION_RING   PadActionType = 1
	GTK_PAD_ACTION_STRIP  PadActionType = 2
)

// This is a generated file - DO NOT EDIT
// +build gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
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

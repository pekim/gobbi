// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// The type of a pad action.
type PadActionType C.GtkPadActionType

const (
	// Action is triggered by a pad button
	GTK_PAD_ACTION_BUTTON PadActionType = 0
	// Action is triggered by a pad ring
	GTK_PAD_ACTION_RING PadActionType = 1
	// Action is triggered by a pad strip
	GTK_PAD_ACTION_STRIP PadActionType = 2
)

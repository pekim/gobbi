// This is a generated file - DO NOT EDIT
// +build gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type EntryIconPosition C.GtkEntryIconPosition

const (
	GTK_ENTRY_ICON_PRIMARY   EntryIconPosition = 0
	GTK_ENTRY_ICON_SECONDARY EntryIconPosition = 1
)
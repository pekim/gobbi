// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type BaselinePosition C.GtkBaselinePosition

const (
	GTK_BASELINE_POSITION_TOP    BaselinePosition = 0
	GTK_BASELINE_POSITION_CENTER BaselinePosition = 1
	GTK_BASELINE_POSITION_BOTTOM BaselinePosition = 2
)

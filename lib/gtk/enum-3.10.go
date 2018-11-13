// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Whenever a container has some form of natural row it may align
// children in that row along a common typographical baseline. If
// the amount of verical space in the row is taller than the total
// requested height of the baseline-aligned children then it can use a
// #GtkBaselinePosition to select where to put the baseline inside the
// extra availible space.
type BaselinePosition C.GtkBaselinePosition

const (
	// Align the baseline at the top
	GTK_BASELINE_POSITION_TOP BaselinePosition = 0
	// Center the baseline
	GTK_BASELINE_POSITION_CENTER BaselinePosition = 1
	// Align the baseline at the bottom
	GTK_BASELINE_POSITION_BOTTOM BaselinePosition = 2
)

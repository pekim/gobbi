// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Granularity types that extend the text selection. Use the
// #GtkTextView::extend-selection signal to customize the selection.
type TextExtendSelection C.GtkTextExtendSelection

const (
	// Selects the current word. It is triggered by
	// a double-click for example.
	GTK_TEXT_EXTEND_SELECTION_WORD TextExtendSelection = 0
	// Selects the current line. It is triggered by
	// a triple-click for example.
	GTK_TEXT_EXTEND_SELECTION_LINE TextExtendSelection = 1
)

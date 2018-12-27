// This is a generated file - DO NOT EDIT
// +build gtk_2.2 gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import gdk "github.com/pekim/gobbi/lib/gdk"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// SelectionOwnerSetForDisplay is a wrapper around the C function gtk_selection_owner_set_for_display.
func SelectionOwnerSetForDisplay(display *gdk.Display, widget *Widget, selection *gdk.Atom, time uint32) bool {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_time_ := (C.guint32)(time)

	retC := C.gtk_selection_owner_set_for_display(c_display, c_widget, c_selection, c_time_)
	retGo := retC == C.TRUE

	return retGo
}

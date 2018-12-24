// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gobject "github.com/pekim/gobbi/lib/gobject"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// BindingsActivateEvent is a wrapper around the C function gtk_bindings_activate_event.
func BindingsActivateEvent(object *gobject.Object, event *gdk.EventKey) bool {
	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_bindings_activate_event(c_object, c_event)
	retGo := retC == C.TRUE

	return retGo
}

// RcResetStyles is a wrapper around the C function gtk_rc_reset_styles.
func RcResetStyles(settings *Settings) {
	c_settings := (*C.GtkSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSettings)(settings.ToC())
	}

	C.gtk_rc_reset_styles(c_settings)

	return
}

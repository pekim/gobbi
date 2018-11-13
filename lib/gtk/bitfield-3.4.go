// This is a generated file - DO NOT EDIT
// +build gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Types of user actions that may be blocked by gtk_application_inhibit().
type ApplicationInhibitFlags C.GtkApplicationInhibitFlags

const (
	// Inhibit ending the user session
	// by logging out or by shutting down the computer
	GTK_APPLICATION_INHIBIT_LOGOUT ApplicationInhibitFlags = 1
	// Inhibit user switching
	GTK_APPLICATION_INHIBIT_SWITCH ApplicationInhibitFlags = 2
	// Inhibit suspending the
	// session or computer
	GTK_APPLICATION_INHIBIT_SUSPEND ApplicationInhibitFlags = 4
	// Inhibit the session being
	// marked as idle (and possibly locked)
	GTK_APPLICATION_INHIBIT_IDLE ApplicationInhibitFlags = 8
)

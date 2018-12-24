// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type License C.GtkLicense

const (
	GTK_LICENSE_UNKNOWN       License = 0
	GTK_LICENSE_CUSTOM        License = 1
	GTK_LICENSE_GPL_2_0       License = 2
	GTK_LICENSE_GPL_3_0       License = 3
	GTK_LICENSE_LGPL_2_1      License = 4
	GTK_LICENSE_LGPL_3_0      License = 5
	GTK_LICENSE_BSD           License = 6
	GTK_LICENSE_MIT_X11       License = 7
	GTK_LICENSE_ARTISTIC      License = 8
	GTK_LICENSE_GPL_2_0_ONLY  License = 9
	GTK_LICENSE_GPL_3_0_ONLY  License = 10
	GTK_LICENSE_LGPL_2_1_ONLY License = 11
	GTK_LICENSE_LGPL_3_0_ONLY License = 12
	GTK_LICENSE_AGPL_3_0      License = 13
	GTK_LICENSE_AGPL_3_0_ONLY License = 14
)

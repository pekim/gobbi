// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// The type of license for an application.
//
// This enumeration can be expanded at later date.
type License C.GtkLicense

const (
	// No license specified
	GTK_LICENSE_UNKNOWN License = 0
	// A license text is going to be specified by the
	// developer
	GTK_LICENSE_CUSTOM License = 1
	// The GNU General Public License, version 2.0 or later
	GTK_LICENSE_GPL_2_0 License = 2
	// The GNU General Public License, version 3.0 or later
	GTK_LICENSE_GPL_3_0 License = 3
	// The GNU Lesser General Public License, version 2.1 or later
	GTK_LICENSE_LGPL_2_1 License = 4
	// The GNU Lesser General Public License, version 3.0 or later
	GTK_LICENSE_LGPL_3_0 License = 5
	// The BSD standard license
	GTK_LICENSE_BSD License = 6
	// The MIT/X11 standard license
	GTK_LICENSE_MIT_X11 License = 7
	// The Artistic License, version 2.0
	GTK_LICENSE_ARTISTIC License = 8
	// The GNU General Public License, version 2.0 only. Since 3.12.
	GTK_LICENSE_GPL_2_0_ONLY License = 9
	// The GNU General Public License, version 3.0 only. Since 3.12.
	GTK_LICENSE_GPL_3_0_ONLY License = 10
	// The GNU Lesser General Public License, version 2.1 only. Since 3.12.
	GTK_LICENSE_LGPL_2_1_ONLY License = 11
	// The GNU Lesser General Public License, version 3.0 only. Since 3.12.
	GTK_LICENSE_LGPL_3_0_ONLY License = 12
	// The GNU Affero General Public License, version 3.0 or later. Since: 3.22.
	GTK_LICENSE_AGPL_3_0 License = 13
	// The GNU Affero General Public License, version 3.0 only. Since: 3.22.27.
	GTK_LICENSE_AGPL_3_0_ONLY License = 14
)

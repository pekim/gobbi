// This is a generated file - DO NOT EDIT
// +build gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// ActivatableIface is a wrapper around the C record GtkActivatableIface.
type ActivatableIface struct {
	native *C.GtkActivatableIface
	// Private : g_iface
	// no type for update
	// no type for sync_action_properties
}

func ActivatableIfaceNewFromC(u unsafe.Pointer) *ActivatableIface {
	c := (*C.GtkActivatableIface)(u)
	if c == nil {
		return nil
	}

	g := &ActivatableIface{native: c}

	return g
}

func (recv *ActivatableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

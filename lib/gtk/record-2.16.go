// This is a generated file - DO NOT EDIT
// +build gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gdk "github.com/pekim/gobbi/lib/gdk"
	"unsafe"
)

// Equals compares this ActivatableIface with another ActivatableIface, and returns true if they represent the same GObject.
func (recv *ActivatableIface) Equals(other *ActivatableIface) bool {
	return other.ToC() == recv.ToC()
}

// GetSelection is a wrapper around the C function gtk_selection_data_get_selection.
func (recv *SelectionData) GetSelection() gdk.Atom {
	retC := C.gtk_selection_data_get_selection((*C.GtkSelectionData)(recv.native))
	retGo := *gdk.AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

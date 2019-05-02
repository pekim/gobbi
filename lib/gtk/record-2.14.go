// This is a generated file - DO NOT EDIT
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gdk "github.com/pekim/gobbi/lib/gdk"
	"unsafe"
)

// Unsupported : gtk_selection_data_get_data : array return type :

// GetDataType is a wrapper around the C function gtk_selection_data_get_data_type.
func (recv *SelectionData) GetDataType() gdk.Atom {
	retC := C.gtk_selection_data_get_data_type((*C.GtkSelectionData)(recv.native))
	retGo := *gdk.AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gtk_selection_data_get_display.
func (recv *SelectionData) GetDisplay() *gdk.Display {
	retC := C.gtk_selection_data_get_display((*C.GtkSelectionData)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFormat is a wrapper around the C function gtk_selection_data_get_format.
func (recv *SelectionData) GetFormat() int32 {
	retC := C.gtk_selection_data_get_format((*C.GtkSelectionData)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLength is a wrapper around the C function gtk_selection_data_get_length.
func (recv *SelectionData) GetLength() int32 {
	retC := C.gtk_selection_data_get_length((*C.GtkSelectionData)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTarget is a wrapper around the C function gtk_selection_data_get_target.
func (recv *SelectionData) GetTarget() gdk.Atom {
	retC := C.gtk_selection_data_get_target((*C.GtkSelectionData)(recv.native))
	retGo := *gdk.AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

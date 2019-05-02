// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// ToGvariant is a wrapper around the C function gtk_paper_size_to_gvariant.
func (recv *PaperSize) ToGvariant() *glib.Variant {
	retC := C.gtk_paper_size_to_gvariant((*C.GtkPaperSize)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

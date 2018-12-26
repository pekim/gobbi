// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// PaperSizeNewFromGvariant is a wrapper around the C function gtk_paper_size_new_from_gvariant.
func PaperSizeNewFromGvariant(variant *glib.Variant) *PaperSize {
	c_variant := (*C.GVariant)(C.NULL)
	if variant != nil {
		c_variant = (*C.GVariant)(variant.ToC())
	}

	retC := C.gtk_paper_size_new_from_gvariant(c_variant)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToGvariant is a wrapper around the C function gtk_paper_size_to_gvariant.
func (recv *PaperSize) ToGvariant() *glib.Variant {
	retC := C.gtk_paper_size_to_gvariant((*C.GtkPaperSize)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

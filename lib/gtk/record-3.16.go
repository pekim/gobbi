// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GLAreaClass is a wrapper around the C record GtkGLAreaClass.
type GLAreaClass struct {
	native *C.GtkGLAreaClass
	// Private : parent_class
	// no type for render
	// no type for resize
	// no type for create_context
	// Private : _padding
}

func GLAreaClassNewFromC(u unsafe.Pointer) *GLAreaClass {
	c := (*C.GtkGLAreaClass)(u)
	if c == nil {
		return nil
	}

	g := &GLAreaClass{native: c}

	return g
}

func (recv *GLAreaClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GLAreaClass with another GLAreaClass, and returns true if they represent the same GObject.
func (recv *GLAreaClass) Equals(other *GLAreaClass) bool {
	return other.ToC() == recv.ToC()
}

// PaperSizeNewFromIpp is a wrapper around the C function gtk_paper_size_new_from_ipp.
func PaperSizeNewFromIpp(ippName string, width float64, height float64) *PaperSize {
	c_ipp_name := C.CString(ippName)
	defer C.free(unsafe.Pointer(c_ipp_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	retC := C.gtk_paper_size_new_from_ipp(c_ipp_name, c_width, c_height)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

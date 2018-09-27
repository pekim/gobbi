// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.20 gtk_3.22

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

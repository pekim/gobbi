// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// GLAreaClass is a wrapper around the C record GtkGLAreaClass.
type GLAreaClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for render
	// no type for resize
	// no type for create_context
	// Private : _padding
}

func GLAreaClassNewFromC(u unsafe.Pointer) *GLAreaClass {
	if u == nil {
		return nil
	}

	g := &GLAreaClass{native: u}

	return g
}

func (recv *GLAreaClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_paper_size_new_from_ipp : return type :

// This is a generated file - DO NOT EDIT
// +build gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.18 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.16 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// AtomInternStaticString is a wrapper around the C function gdk_atom_intern_static_string.
func AtomInternStaticString(atomName string) Atom {
	c_atom_name := C.CString(atomName)
	defer C.free(unsafe.Pointer(c_atom_name))

	retC := C.gdk_atom_intern_static_string(c_atom_name)
	retGo := AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

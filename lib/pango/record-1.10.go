// This is a generated file - DO NOT EDIT
// +build pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Ref is a wrapper around the C function pango_attr_list_ref.
func (recv *AttrList) Ref() *AttrList {
	retC := C.pango_attr_list_ref((*C.PangoAttrList)(recv.native))
	retGo := AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function pango_layout_line_ref.
func (recv *LayoutLine) Ref() *LayoutLine {
	retC := C.pango_layout_line_ref((*C.PangoLayoutLine)(recv.native))
	retGo := LayoutLineNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Create a new background alpha attribute.
/*

C function

pango_attr_background_alpha_new
*/
func AttrBackgroundAlphaNew(alpha uint16) *Attribute {
	c_alpha := (C.guint16)(alpha)

	retC := C.pango_attr_background_alpha_new(c_alpha)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new foreground alpha attribute.
/*

C function

pango_attr_foreground_alpha_new
*/
func AttrForegroundAlphaNew(alpha uint16) *Attribute {
	c_alpha := (C.guint16)(alpha)

	retC := C.pango_attr_foreground_alpha_new(c_alpha)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

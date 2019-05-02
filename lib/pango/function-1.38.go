// This is a generated file - DO NOT EDIT
// +build pango_1.38

package pango

import (
	"C"
	"unsafe"
)

// AttrBackgroundAlphaNew is a wrapper around the C function pango_attr_background_alpha_new.
func AttrBackgroundAlphaNew(alpha uint16) *Attribute {
	c_alpha := (C.guint16)(alpha)

	retC := C.pango_attr_background_alpha_new(c_alpha)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AttrForegroundAlphaNew is a wrapper around the C function pango_attr_foreground_alpha_new.
func AttrForegroundAlphaNew(alpha uint16) *Attribute {
	c_alpha := (C.guint16)(alpha)

	retC := C.pango_attr_foreground_alpha_new(c_alpha)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

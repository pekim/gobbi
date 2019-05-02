// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// pango_attr_shape_new_with_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// AttrSizeNewAbsolute is a wrapper around the C function pango_attr_size_new_absolute.
func AttrSizeNewAbsolute(size int32) *Attribute {
	c_size := (C.int)(size)

	retC := C.pango_attr_size_new_absolute(c_size)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSizeIsAbsolute is a wrapper around the C function pango_font_description_get_size_is_absolute.
func (recv *FontDescription) GetSizeIsAbsolute() bool {
	retC := C.pango_font_description_get_size_is_absolute((*C.PangoFontDescription)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAbsoluteSize is a wrapper around the C function pango_font_description_set_absolute_size.
func (recv *FontDescription) SetAbsoluteSize(size float64) {
	c_size := (C.double)(size)

	C.pango_font_description_set_absolute_size((*C.PangoFontDescription)(recv.native), c_size)

	return
}

// Equals compares this RendererClass with another RendererClass, and returns true if they represent the same GObject.
func (recv *RendererClass) Equals(other *RendererClass) bool {
	return other.ToC() == recv.ToC()
}

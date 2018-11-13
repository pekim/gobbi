// This is a generated file - DO NOT EDIT
// +build pango_1.8 pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Create a new strikethrough color attribute. This attribute
// modifies the color of strikethrough lines. If not set, strikethrough
// lines will use the foreground color.
/*

C function : pango_attr_strikethrough_color_new
*/
func AttrStrikethroughColorNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_strikethrough_color_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Create a new underline color attribute. This attribute
// modifies the color of underlines. If not set, underlines
// will use the foreground color.
/*

C function : pango_attr_underline_color_new
*/
func AttrUnderlineColorNew(red uint16, green uint16, blue uint16) *Attribute {
	c_red := (C.guint16)(red)

	c_green := (C.guint16)(green)

	c_blue := (C.guint16)(blue)

	retC := C.pango_attr_underline_color_new(c_red, c_green, c_blue)
	retGo := AttributeNewFromC(unsafe.Pointer(retC))

	return retGo
}

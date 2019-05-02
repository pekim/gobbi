// Code generated - DO NOT EDIT.
// +build pango_1.10 pango_1.12 pango_1.14 pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// GetFontMap is a wrapper around the C function pango_font_get_font_map.
func (recv *Font) GetFontMap() *FontMap {
	retC := C.pango_font_get_font_map((*C.PangoFont)(recv.native))
	var retGo (*FontMap)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FontMapNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// IsZeroWidth is a wrapper around the C function pango_is_zero_width.
func IsZeroWidth(ch rune) bool {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_is_zero_width(c_ch)
	retGo := retC == C.TRUE

	return retGo
}

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

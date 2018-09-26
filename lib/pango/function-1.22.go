// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.31.0 pango_1.32 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// AttrTypeGetName is a wrapper around the C function pango_attr_type_get_name.
func AttrTypeGetName(type_ AttrType) string {
	c_type := (C.PangoAttrType)(type_)

	retC := C.pango_attr_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// BidiTypeForUnichar is a wrapper around the C function pango_bidi_type_for_unichar.
func BidiTypeForUnichar(ch rune) BidiType {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_bidi_type_for_unichar(c_ch)
	retGo := (BidiType)(retC)

	return retGo
}

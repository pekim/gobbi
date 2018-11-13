// This is a generated file - DO NOT EDIT
// +build pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Fetches the attribute type name passed in when registering the type using
// pango_attr_type_register().
//
// The returned value is an interned string (see g_intern_string() for what
// that means) that should not be modified or freed.
/*

C function

pango_attr_type_get_name
*/
func AttrTypeGetName(type_ AttrType) string {
	c_type := (C.PangoAttrType)(type_)

	retC := C.pango_attr_type_get_name(c_type)
	retGo := C.GoString(retC)

	return retGo
}

// Determines the normative bidirectional character type of a
// character, as specified in the Unicode Character Database.
//
// A simplified version of this function is available as
// pango_unichar_direction().
/*

C function

pango_bidi_type_for_unichar
*/
func BidiTypeForUnichar(ch rune) BidiType {
	c_ch := (C.gunichar)(ch)

	retC := C.pango_bidi_type_for_unichar(c_ch)
	retGo := (BidiType)(retC)

	return retGo
}

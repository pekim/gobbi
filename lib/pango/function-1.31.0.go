// This is a generated file - DO NOT EDIT
// +build pango_1.31.0 pango_1.32 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// Unsupported : pango_markup_parser_finish : unsupported parameter attr_list : record with indirection level of 2

// MarkupParserNew is a wrapper around the C function pango_markup_parser_new.
func MarkupParserNew(accelMarker rune) *glib.MarkupParseContext {
	c_accel_marker := (C.gunichar)(accelMarker)

	retC := C.pango_markup_parser_new(c_accel_marker)
	retGo := glib.MarkupParseContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

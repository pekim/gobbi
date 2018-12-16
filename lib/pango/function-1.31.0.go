// This is a generated file - DO NOT EDIT
// +build pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
// #include <stdlib.h>
import "C"

// MarkupParserFinish is a wrapper around the C function pango_markup_parser_finish.
func MarkupParserFinish(context *glib.MarkupParseContext) (bool, *AttrList, string, rune, error) {
	c_context := (*C.GMarkupParseContext)(C.NULL)
	if context != nil {
		c_context = (*C.GMarkupParseContext)(context.ToC())
	}

	var c_attr_list *C.PangoAttrList

	var c_text *C.char

	var c_accel_char C.gunichar

	var cThrowableError *C.GError

	retC := C.pango_markup_parser_finish(c_context, &c_attr_list, &c_text, &c_accel_char, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	attrList := AttrListNewFromC(unsafe.Pointer(c_attr_list))

	text := C.GoString(c_text)
	defer C.free(unsafe.Pointer(c_text))

	accelChar := (rune)(c_accel_char)

	return retGo, attrList, text, accelChar, goError
}

// MarkupParserNew is a wrapper around the C function pango_markup_parser_new.
func MarkupParserNew(accelMarker rune) *glib.MarkupParseContext {
	c_accel_marker := (C.gunichar)(accelMarker)

	retC := C.pango_markup_parser_new(c_accel_marker)
	retGo := glib.MarkupParseContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

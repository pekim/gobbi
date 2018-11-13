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

// After feeding a pango markup parser some data with g_markup_parse_context_parse(),
// use this function to get the list of pango attributes and text out of the
// markup. This function will not free @context, use g_markup_parse_context_free()
// to do so.
/*

C function

pango_markup_parser_finish
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	attrList := AttrListNewFromC(unsafe.Pointer(c_attr_list))

	text := C.GoString(c_text)
	defer C.free(unsafe.Pointer(c_text))

	accelChar := (rune)(c_accel_char)

	return retGo, attrList, text, accelChar, goThrowableError
}

// Parses marked-up text (see
// <link linkend="PangoMarkupFormat">markup format</link>) to create
// a plain-text string and an attribute list.
//
// If @accel_marker is nonzero, the given character will mark the
// character following it as an accelerator. For example, @accel_marker
// might be an ampersand or underscore. All characters marked
// as an accelerator will receive a %PANGO_UNDERLINE_LOW attribute,
// and the first character so marked will be returned in @accel_char,
// when calling finish(). Two @accel_marker characters following each
// other produce a single literal @accel_marker character.
//
// To feed markup to the parser, use g_markup_parse_context_parse()
// on the returned #GMarkupParseContext. When done with feeding markup
// to the parser, use pango_markup_parser_finish() to get the data out
// of it, and then use g_markup_parse_context_free() to free it.
//
// This function is designed for applications that read pango markup
// from streams. To simply parse a string containing pango markup,
// the simpler pango_parse_markup() API is recommended instead.
/*

C function

pango_markup_parser_new
*/
func MarkupParserNew(accelMarker rune) *glib.MarkupParseContext {
	c_accel_marker := (C.gunichar)(accelMarker)

	retC := C.pango_markup_parser_new(c_accel_marker)
	retGo := glib.MarkupParseContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

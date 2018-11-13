// This is a generated file - DO NOT EDIT
// +build glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Error codes returned by functions converting a string to a number.
type NumberParserError C.GNumberParserError

const (
	// String was not a valid number.
	NUMBER_PARSER_ERROR_INVALID NumberParserError = 0
	// String was a number, but out of bounds.
	NUMBER_PARSER_ERROR_OUT_OF_BOUNDS NumberParserError = 1
)

// +build glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

type NumberParserError C.GNumberParserError

const (
	NUMBER_PARSER_ERROR_INVALID       NumberParserError = 0
	NUMBER_PARSER_ERROR_OUT_OF_BOUNDS NumberParserError = 1
)

// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Logfield is a wrapper around the C record GLogField.
type Logfield struct {
	native *C.GLogField
	Key    int
	Value  int
	Length int
}

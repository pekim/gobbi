// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

/*
A format specifier that can be used in printf()-style format strings
when printing a #GPid.
*/
const PID_FORMAT string = C.G_PID_FORMAT

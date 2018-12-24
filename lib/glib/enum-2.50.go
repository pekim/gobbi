// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

type LogWriterOutput C.GLogWriterOutput

const (
	LOG_WRITER_HANDLED   LogWriterOutput = 1
	LOG_WRITER_UNHANDLED LogWriterOutput = 0
)

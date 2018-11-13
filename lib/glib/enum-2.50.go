// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Return values from #GLogWriterFuncs to indicate whether the given log entry
// was successfully handled by the writer, or whether there was an error in
// handling it (and hence a fallback writer should be used).
//
// If a #GLogWriterFunc ignores a log entry, it should return
// %G_LOG_WRITER_HANDLED.
type LogWriterOutput C.GLogWriterOutput

const (
	// Log writer has handled the log entry.
	LOG_WRITER_HANDLED LogWriterOutput = 1

	// Log writer could not handle the log entry.
	LOG_WRITER_UNHANDLED LogWriterOutput = 0
)

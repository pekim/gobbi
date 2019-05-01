// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_compute_hmac_for_bytes : return type :

// Unsupported : g_log_set_writer_func : unsupported parameter func : no type generator for LogWriterFunc (GLogWriterFunc) for param func

// Unsupported : g_log_structured : unsupported parameter ... : varargs

// Unsupported : g_log_structured_array : unsupported parameter fields :

// Unsupported : g_log_writer_default : unsupported parameter fields :

// Unsupported : g_log_writer_format_fields : unsupported parameter fields :

// Unsupported : g_log_writer_is_journald : return type :

// Unsupported : g_log_writer_journald : unsupported parameter fields :

// Unsupported : g_log_writer_standard_streams : unsupported parameter fields :

// Unsupported : g_log_writer_supports_color : return type :

// LogField is a wrapper around the C record GLogField.
type LogField struct {
	native *C.GLogField
	Key    string
	// value : no type generator for gpointer, gconstpointer
	Length int64
}

func LogFieldNewFromC(u unsafe.Pointer) *LogField {
	c := (*C.GLogField)(u)
	if c == nil {
		return nil
	}

	g := &LogField{
		Key:    C.GoString(c.key),
		Length: (int64)(c.length),
		native: c,
	}

	return g
}

func (recv *LogField) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

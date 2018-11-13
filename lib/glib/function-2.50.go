// This is a generated file - DO NOT EDIT
// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Computes the HMAC for a binary @data. This is a
// convenience wrapper for g_hmac_new(), g_hmac_get_string()
// and g_hmac_unref().
//
// The hexadecimal string returned will be in lower case.
/*

C function

g_compute_hmac_for_bytes
*/
func ComputeHmacForBytes(digestType ChecksumType, key *Bytes, data *Bytes) string {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key := (*C.GBytes)(C.NULL)
	if key != nil {
		c_key = (*C.GBytes)(key.ToC())
	}

	c_data := (*C.GBytes)(C.NULL)
	if data != nil {
		c_data = (*C.GBytes)(data.ToC())
	}

	retC := C.g_compute_hmac_for_bytes(c_digest_type, c_key, c_data)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_log_set_writer_func : unsupported parameter func : no type generator for LogWriterFunc (GLogWriterFunc) for param func

// Unsupported : g_log_structured : unsupported parameter ... : varargs

// Unsupported : g_log_structured_array : unsupported parameter fields :

// Unsupported : g_log_variant : unsupported parameter fields : Blacklisted record : GVariant

// Unsupported : g_log_writer_default : unsupported parameter fields :

// Unsupported : g_log_writer_format_fields : unsupported parameter fields :

// Check whether the given @output_fd file descriptor is a connection to the
// systemd journal, or something else (like a log file or `stdout` or
// `stderr`).
//
// Invalid file descriptors are accepted and return %FALSE, which allows for
// the following construct without needing any additional error handling:
// |[<!-- language="C" -->
// is_journald = g_log_writer_is_journald (fileno (stderr));
// ]|
/*

C function

g_log_writer_is_journald
*/
func LogWriterIsJournald(outputFd int32) bool {
	c_output_fd := (C.gint)(outputFd)

	retC := C.g_log_writer_is_journald(c_output_fd)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_log_writer_journald : unsupported parameter fields :

// Unsupported : g_log_writer_standard_streams : unsupported parameter fields :

// Check whether the given @output_fd file descriptor supports ANSI color
// escape sequences. If so, they can safely be used when formatting log
// messages.
/*

C function

g_log_writer_supports_color
*/
func LogWriterSupportsColor(outputFd int32) bool {
	c_output_fd := (C.gint)(outputFd)

	retC := C.g_log_writer_supports_color(c_output_fd)
	retGo := retC == C.TRUE

	return retGo
}

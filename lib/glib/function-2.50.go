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
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ComputeHmacForBytes is a wrapper around the C function g_compute_hmac_for_bytes.
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

// LogVariant is a wrapper around the C function g_log_variant.
func LogVariant(logDomain string, logLevel LogLevelFlags, fields *Variant) {
	c_log_domain := C.CString(logDomain)
	defer C.free(unsafe.Pointer(c_log_domain))

	c_log_level := (C.GLogLevelFlags)(logLevel)

	c_fields := (*C.GVariant)(C.NULL)
	if fields != nil {
		c_fields = (*C.GVariant)(fields.ToC())
	}

	C.g_log_variant(c_log_domain, c_log_level, c_fields)

	return
}

// Unsupported : g_log_writer_default : unsupported parameter fields :

// Unsupported : g_log_writer_format_fields : unsupported parameter fields :

// LogWriterIsJournald is a wrapper around the C function g_log_writer_is_journald.
func LogWriterIsJournald(outputFd int32) bool {
	c_output_fd := (C.gint)(outputFd)

	retC := C.g_log_writer_is_journald(c_output_fd)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_log_writer_journald : unsupported parameter fields :

// Unsupported : g_log_writer_standard_streams : unsupported parameter fields :

// LogWriterSupportsColor is a wrapper around the C function g_log_writer_supports_color.
func LogWriterSupportsColor(outputFd int32) bool {
	c_output_fd := (C.gint)(outputFd)

	retC := C.g_log_writer_supports_color(c_output_fd)
	retGo := retC == C.TRUE

	return retGo
}

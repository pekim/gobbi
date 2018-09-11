// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_compute_hmac_for_bytes : unsupported parameter digest_type : no param type

// Unsupported : g_log_set_writer_func : unsupported parameter func : no param type

// Unsupported : g_log_structured : unsupported parameter log_level : no param type

// Unsupported : g_log_structured_array : unsupported parameter log_level : no param type

// Unsupported : g_log_variant : unsupported parameter log_level : no param type

// Unsupported : g_log_writer_default : unsupported parameter log_level : no param type

// Unsupported : g_log_writer_format_fields : unsupported parameter log_level : no param type

// LogWriterIsJournald is a wrapper around the C function g_log_writer_is_journald.
func LogWriterIsJournald(outputFd int32) {
	c_output_fd := (C.gint)(outputFd)

	C.g_log_writer_is_journald(c_output_fd)
}

// Unsupported : g_log_writer_journald : unsupported parameter log_level : no param type

// Unsupported : g_log_writer_standard_streams : unsupported parameter log_level : no param type

// LogWriterSupportsColor is a wrapper around the C function g_log_writer_supports_color.
func LogWriterSupportsColor(outputFd int32) {
	c_output_fd := (C.gint)(outputFd)

	C.g_log_writer_supports_color(c_output_fd)
}

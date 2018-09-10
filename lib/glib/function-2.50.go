// +build glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_compute_hmac_for_bytes : unsupported parameter digest_type : type ChecksumType, GChecksumType

// Unsupported function: g_log_set_writer_func : unsupported parameter func : type LogWriterFunc, GLogWriterFunc

// Unsupported function: g_log_structured : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_structured_array : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_variant : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_writer_default : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_writer_format_fields : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// LogWriterIsJournald is a wrapper around the C function g_log_writer_is_journald.
func LogWriterIsJournald(outputFd int32) {}

// Unsupported function: g_log_writer_journald : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// Unsupported function: g_log_writer_standard_streams : unsupported parameter log_level : type LogLevelFlags, GLogLevelFlags

// LogWriterSupportsColor is a wrapper around the C function g_log_writer_supports_color.
func LogWriterSupportsColor(outputFd int32) {}

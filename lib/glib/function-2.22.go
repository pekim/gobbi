// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_byte_array_unref : unsupported parameter array : no type

// Unsupported : g_double_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported : g_double_hash : unsupported parameter v : type gpointer, gconstpointer

// HostnameIsAsciiEncoded is a wrapper around the C function g_hostname_is_ascii_encoded.
func HostnameIsAsciiEncoded(hostname string) {}

// HostnameIsIpAddress is a wrapper around the C function g_hostname_is_ip_address.
func HostnameIsIpAddress(hostname string) {}

// HostnameIsNonAscii is a wrapper around the C function g_hostname_is_non_ascii.
func HostnameIsNonAscii(hostname string) {}

// HostnameToAscii is a wrapper around the C function g_hostname_to_ascii.
func HostnameToAscii(hostname string) {}

// HostnameToUnicode is a wrapper around the C function g_hostname_to_unicode.
func HostnameToUnicode(hostname string) {}

// Unsupported : g_int64_equal : unsupported parameter v1 : type gpointer, gconstpointer

// Unsupported : g_int64_hash : unsupported parameter v : type gpointer, gconstpointer

// MainContextGetThreadDefault is a wrapper around the C function g_main_context_get_thread_default.
func MainContextGetThreadDefault() {}

// Unsupported : g_mkstemp_full : unsupported parameter tmpl : type filename, gchar*

// ReloadUserSpecialDirsCache is a wrapper around the C function g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {}

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : type TestLogFatalFunc, GTestLogFatalFunc

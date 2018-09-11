// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_byte_array_unref : unsupported parameter array : no param type

// Unsupported : g_double_equal : unsupported parameter v1 : no param type

// Unsupported : g_double_hash : unsupported parameter v : no param type

// HostnameIsAsciiEncoded is a wrapper around the C function g_hostname_is_ascii_encoded.
func HostnameIsAsciiEncoded(hostname string) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	C.g_hostname_is_ascii_encoded(c_hostname)
}

// HostnameIsIpAddress is a wrapper around the C function g_hostname_is_ip_address.
func HostnameIsIpAddress(hostname string) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	C.g_hostname_is_ip_address(c_hostname)
}

// HostnameIsNonAscii is a wrapper around the C function g_hostname_is_non_ascii.
func HostnameIsNonAscii(hostname string) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	C.g_hostname_is_non_ascii(c_hostname)
}

// HostnameToAscii is a wrapper around the C function g_hostname_to_ascii.
func HostnameToAscii(hostname string) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	C.g_hostname_to_ascii(c_hostname)
}

// HostnameToUnicode is a wrapper around the C function g_hostname_to_unicode.
func HostnameToUnicode(hostname string) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	C.g_hostname_to_unicode(c_hostname)
}

// Unsupported : g_int64_equal : unsupported parameter v1 : no param type

// Unsupported : g_int64_hash : unsupported parameter v : no param type

// MainContextGetThreadDefault is a wrapper around the C function g_main_context_get_thread_default.
func MainContextGetThreadDefault() {
	C.g_main_context_get_thread_default()
}

// Unsupported : g_mkstemp_full : unsupported parameter tmpl : no param type

// ReloadUserSpecialDirsCache is a wrapper around the C function g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()
}

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no param type

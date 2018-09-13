// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_byte_array_unref : unsupported parameter array : no param type

// Unsupported : g_double_equal : no return type

// DoubleHash is a wrapper around the C function g_double_hash.
func DoubleHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_double_hash(c_v)
	retGo :=
		(uint32)(retC)

	return retGo
}

// Unsupported : g_hostname_is_ascii_encoded : no return type

// Unsupported : g_hostname_is_ip_address : no return type

// Unsupported : g_hostname_is_non_ascii : no return type

// HostnameToAscii is a wrapper around the C function g_hostname_to_ascii.
func HostnameToAscii(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_ascii(c_hostname)
	retGo :=
		C.GoString(retC)

	return retGo
}

// HostnameToUnicode is a wrapper around the C function g_hostname_to_unicode.
func HostnameToUnicode(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_unicode(c_hostname)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_int64_equal : no return type

// Int64Hash is a wrapper around the C function g_int64_hash.
func Int64Hash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_int64_hash(c_v)
	retGo :=
		(uint32)(retC)

	return retGo
}

// Unsupported : g_main_context_get_thread_default : no return type

// MkstempFull is a wrapper around the C function g_mkstemp_full.
func MkstempFull(tmpl string, flags int32, mode int32) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_flags := (C.gint)(flags)

	c_mode := (C.gint)(mode)

	retC := C.g_mkstemp_full(c_tmpl, c_flags, c_mode)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_reload_user_special_dirs_cache : no return type

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no param type for TestLogFatalFunc, GTestLogFatalFunc

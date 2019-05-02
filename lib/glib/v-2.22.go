// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

const GINTPTR_FORMAT string = C.G_GINTPTR_FORMAT
const GINTPTR_MODIFIER string = C.G_GINTPTR_MODIFIER
const GUINTPTR_FORMAT string = C.G_GUINTPTR_FORMAT

// Unsupported : g_double_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_double_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// HostnameIsAsciiEncoded is a wrapper around the C function g_hostname_is_ascii_encoded.
func HostnameIsAsciiEncoded(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ascii_encoded(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameIsIpAddress is a wrapper around the C function g_hostname_is_ip_address.
func HostnameIsIpAddress(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ip_address(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameIsNonAscii is a wrapper around the C function g_hostname_is_non_ascii.
func HostnameIsNonAscii(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_non_ascii(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// HostnameToAscii is a wrapper around the C function g_hostname_to_ascii.
func HostnameToAscii(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_ascii(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// HostnameToUnicode is a wrapper around the C function g_hostname_to_unicode.
func HostnameToUnicode(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_unicode(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_int64_equal : unsupported parameter v1 : no type generator for gpointer (gconstpointer) for param v1

// Unsupported : g_int64_hash : unsupported parameter v : no type generator for gpointer (gconstpointer) for param v

// MkstempFull is a wrapper around the C function g_mkstemp_full.
func MkstempFull(tmpl string, flags int32, mode int32) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_flags := (C.gint)(flags)

	c_mode := (C.gint)(mode)

	retC := C.g_mkstemp_full(c_tmpl, c_flags, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// ReloadUserSpecialDirsCache is a wrapper around the C function g_reload_user_special_dirs_cache.
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()

	return
}

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc (GTestLogFatalFunc) for param log_func

// g_array_get_element_size : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_ref : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_array_unref : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// MainContextGetThreadDefault is a wrapper around the C function g_main_context_get_thread_default.
func MainContextGetThreadDefault() *MainContext {
	retC := C.g_main_context_get_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopThreadDefault is a wrapper around the C function g_main_context_pop_thread_default.
func (recv *MainContext) PopThreadDefault() {
	C.g_main_context_pop_thread_default((*C.GMainContext)(recv.native))

	return
}

// PushThreadDefault is a wrapper around the C function g_main_context_push_thread_default.
func (recv *MainContext) PushThreadDefault() {
	C.g_main_context_push_thread_default((*C.GMainContext)(recv.native))

	return
}

// Ref is a wrapper around the C function g_mapped_file_ref.
func (recv *MappedFile) Ref() *MappedFile {
	retC := C.g_mapped_file_ref((*C.GMappedFile)(recv.native))
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_tree_ref.
func (recv *Tree) Ref() *Tree {
	retC := C.g_tree_ref((*C.GTree)(recv.native))
	retGo := TreeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_tree_unref.
func (recv *Tree) Unref() {
	C.g_tree_unref((*C.GTree)(recv.native))

	return
}

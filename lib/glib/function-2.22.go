// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Atomically decrements the reference count of @array by one. If the
// reference count drops to 0, all memory allocated by the array is
// released. This function is thread-safe and may be called from any
// thread.
/*

C function : g_byte_array_unref
*/
func ByteArrayUnref(array []uint8) {
	c_array := &array[0]

	C.g_byte_array_unref((*C.GByteArray)(unsafe.Pointer(c_array)))

	return
}

// Compares the two #gdouble values being pointed to and returns
// %TRUE if they are equal.
// It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-%NULL pointers to doubles as keys in a
// #GHashTable.
/*

C function : g_double_equal
*/
func DoubleEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_double_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Converts a pointer to a #gdouble to a hash value.
// It can be passed to g_hash_table_new() as the @hash_func parameter,
// It can be passed to g_hash_table_new() as the @hash_func parameter,
// when using non-%NULL pointers to doubles as keys in a #GHashTable.
/*

C function : g_double_hash
*/
func DoubleHash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_double_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Tests if @hostname contains segments with an ASCII-compatible
// encoding of an Internationalized Domain Name. If this returns
// %TRUE, you should decode the hostname with g_hostname_to_unicode()
// before displaying it to the user.
//
// Note that a hostname might contain a mix of encoded and unencoded
// segments, and so it is possible for g_hostname_is_non_ascii() and
// g_hostname_is_ascii_encoded() to both return %TRUE for a name.
/*

C function : g_hostname_is_ascii_encoded
*/
func HostnameIsAsciiEncoded(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ascii_encoded(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// Tests if @hostname is the string form of an IPv4 or IPv6 address.
// (Eg, "192.168.0.1".)
/*

C function : g_hostname_is_ip_address
*/
func HostnameIsIpAddress(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_ip_address(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// Tests if @hostname contains Unicode characters. If this returns
// %TRUE, you need to encode the hostname with g_hostname_to_ascii()
// before using it in non-IDN-aware contexts.
//
// Note that a hostname might contain a mix of encoded and unencoded
// segments, and so it is possible for g_hostname_is_non_ascii() and
// g_hostname_is_ascii_encoded() to both return %TRUE for a name.
/*

C function : g_hostname_is_non_ascii
*/
func HostnameIsNonAscii(hostname string) bool {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_is_non_ascii(c_hostname)
	retGo := retC == C.TRUE

	return retGo
}

// Converts @hostname to its canonical ASCII form; an ASCII-only
// string containing no uppercase letters and not ending with a
// trailing dot.
/*

C function : g_hostname_to_ascii
*/
func HostnameToAscii(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_ascii(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Converts @hostname to its canonical presentation form; a UTF-8
// string in Unicode normalization form C, containing no uppercase
// letters, no forbidden characters, and no ASCII-encoded segments,
// and not ending with a trailing dot.
//
// Of course if @hostname is not an internationalized hostname, then
// the canonical presentation form will be entirely ASCII.
/*

C function : g_hostname_to_unicode
*/
func HostnameToUnicode(hostname string) string {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	retC := C.g_hostname_to_unicode(c_hostname)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Compares the two #gint64 values being pointed to and returns
// %TRUE if they are equal.
// It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-%NULL pointers to 64-bit integers as keys in a
// #GHashTable.
/*

C function : g_int64_equal
*/
func Int64Equal(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.g_int64_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Converts a pointer to a #gint64 to a hash value.
//
// It can be passed to g_hash_table_new() as the @hash_func parameter,
// when using non-%NULL pointers to 64-bit integer values as keys in a
// #GHashTable.
/*

C function : g_int64_hash
*/
func Int64Hash(v uintptr) uint32 {
	c_v := (C.gconstpointer)(v)

	retC := C.g_int64_hash(c_v)
	retGo := (uint32)(retC)

	return retGo
}

// Gets the thread-default #GMainContext for this thread. Asynchronous
// operations that want to be able to be run in contexts other than
// the default one should call this method or
// g_main_context_ref_thread_default() to get a #GMainContext to add
// their #GSources to. (Note that even in single-threaded
// programs applications may sometimes want to temporarily push a
// non-default context, so it is not safe to assume that this will
// always return %NULL if you are running in the default thread.)
//
// If you need to hold a reference on the context, use
// g_main_context_ref_thread_default() instead.
/*

C function : g_main_context_get_thread_default
*/
func MainContextGetThreadDefault() *MainContext {
	retC := C.g_main_context_get_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Opens a temporary file. See the mkstemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkstemp() templates, i.e. contain the string "XXXXXX".
// g_mkstemp_full() is slightly more flexible than mkstemp()
// in that the sequence does not have to occur at the very end of the
// template and you can pass a @mode and additional @flags. The X
// string will be modified to form the name of a file that didn't exist.
// The string should be in the GLib file name encoding. Most importantly,
// on Windows it should be in UTF-8.
/*

C function : g_mkstemp_full
*/
func MkstempFull(tmpl string, flags int32, mode int32) int32 {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_flags := (C.gint)(flags)

	c_mode := (C.gint)(mode)

	retC := C.g_mkstemp_full(c_tmpl, c_flags, c_mode)
	retGo := (int32)(retC)

	return retGo
}

// Resets the cache used for g_get_user_special_dir(), so
// that the latest on-disk version is used. Call this only
// if you just changed the data on disk yourself.
//
// Due to threadsafety issues this may cause leaking of strings
// that were previously returned from g_get_user_special_dir()
// that can't be freed. We ensure to only leak the data for
// the directories that actually changed value though.
/*

C function : g_reload_user_special_dirs_cache
*/
func ReloadUserSpecialDirsCache() {
	C.g_reload_user_special_dirs_cache()

	return
}

// Unsupported : g_test_log_set_fatal_handler : unsupported parameter log_func : no type generator for TestLogFatalFunc (GTestLogFatalFunc) for param log_func

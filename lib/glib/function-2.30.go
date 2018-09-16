// This is a generated file - DO NOT EDIT
// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"fmt"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_atomic_int_and : unsupported parameter atomic : no type generator for guint, volatile guint*

// Unsupported : g_atomic_int_or : unsupported parameter atomic : no type generator for guint, volatile guint*

// Unsupported : g_atomic_int_xor : unsupported parameter atomic : no type generator for guint, volatile guint*

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer, void*

// Unsupported : g_compute_hmac_for_data : unsupported parameter digest_type : no type generator for ChecksumType, GChecksumType

// Unsupported : g_compute_hmac_for_string : unsupported parameter digest_type : no type generator for ChecksumType, GChecksumType

// DirMakeTmp is a wrapper around the C function g_dir_make_tmp.
func DirMakeTmp(tmpl string) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var cThrowableError *C.GError

	retC := C.g_dir_make_tmp(c_tmpl, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := errorNewFromC(cThrowableError)
	fmt.Println(goThrowableError)

	return retGo
}

// FormatSize is a wrapper around the C function g_format_size.
func FormatSize(size uint64) string {
	c_size := (C.guint64)(size)

	retC := C.g_format_size(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_format_size_full : unsupported parameter flags : no type generator for FormatSizeFlags, GFormatSizeFlags

// Mkdtemp is a wrapper around the C function g_mkdtemp.
func Mkdtemp(tmpl string) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkdtemp(c_tmpl)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// MkdtempFull is a wrapper around the C function g_mkdtemp_full.
func MkdtempFull(tmpl string, mode int32) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdtemp_full(c_tmpl, c_mode)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer, void*

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer, void*

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer, void*

// RegexEscapeNul is a wrapper around the C function g_regex_escape_nul.
func RegexEscapeNul(string string, length int32) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gint)(length)

	retC := C.g_regex_escape_nul(c_string, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_test_fail : no return generator

// Unsupported : g_unichar_compose : unsupported parameter ch : no type generator for gunichar, gunichar*

// Unsupported : g_unichar_decompose : unsupported parameter a : no type generator for gunichar, gunichar*

// Unsupported : g_unichar_fully_decompose : unsupported parameter compat : no type generator for gboolean, gboolean

// Unsupported : g_unicode_script_from_iso15924 : no return generator

// Unsupported : g_unicode_script_to_iso15924 : unsupported parameter script : no type generator for UnicodeScript, GUnicodeScript

// Unsupported : g_unix_open_pipe : unsupported parameter fds : no type generator for gint, gint*

// Unsupported : g_unix_set_fd_nonblocking : unsupported parameter nonblock : no type generator for gboolean, gboolean

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_unix_signal_source_new : no return generator

// Utf8Substring is a wrapper around the C function g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_start_pos := (C.glong)(startPos)

	c_end_pos := (C.glong)(endPos)

	retC := C.g_utf8_substring(c_str, c_start_pos, c_end_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

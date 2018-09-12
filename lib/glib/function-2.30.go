// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_atomic_int_and : unsupported parameter atomic : no param type for guint, volatile guint*

// Unsupported : g_atomic_int_or : unsupported parameter atomic : no param type for guint, volatile guint*

// Unsupported : g_atomic_int_xor : unsupported parameter atomic : no param type for guint, volatile guint*

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no param type for gpointer, void*

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no param type for gpointer, void*

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no param type for gpointer, void*

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no param type for gpointer, void*

// Unsupported : g_compute_hmac_for_data : unsupported parameter digest_type : no param type for ChecksumType, GChecksumType

// Unsupported : g_compute_hmac_for_string : unsupported parameter digest_type : no param type for ChecksumType, GChecksumType

// Unsupported : g_dir_make_tmp : unsupported parameter tmpl : no param type for filename, const gchar*

// FormatSize is a wrapper around the C function g_format_size.
func FormatSize(size uint64) string {
	c_size := (C.guint64)(size)

	retC := C.g_format_size(c_size)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_format_size_full : unsupported parameter flags : no param type for FormatSizeFlags, GFormatSizeFlags

// Unsupported : g_mkdtemp : unsupported parameter tmpl : no param type for filename, gchar*

// Unsupported : g_mkdtemp_full : unsupported parameter tmpl : no param type for filename, gchar*

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no param type for gpointer, void*

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no param type for gpointer, void*

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no param type for gpointer, void*

// RegexEscapeNul is a wrapper around the C function g_regex_escape_nul.
func RegexEscapeNul(string string, length int32) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gint)(length)

	retC := C.g_regex_escape_nul(c_string, c_length)
	retGo :=
		C.GoString(retC)

	return retGo
}

// Unsupported : g_test_fail : no return type

// Unsupported : g_unichar_compose : unsupported parameter ch : no param type for gunichar, gunichar*

// Unsupported : g_unichar_decompose : unsupported parameter a : no param type for gunichar, gunichar*

// Unsupported : g_unichar_fully_decompose : unsupported parameter compat : no param type for gboolean, gboolean

// Unsupported : g_unicode_script_from_iso15924 : no return type

// Unsupported : g_unicode_script_to_iso15924 : unsupported parameter script : no param type for UnicodeScript, GUnicodeScript

// Unsupported : g_unix_open_pipe : unsupported parameter fds : no param type for gint, gint*

// Unsupported : g_unix_set_fd_nonblocking : unsupported parameter nonblock : no param type for gboolean, gboolean

// Unsupported : g_unix_signal_add : unsupported parameter handler : no param type for SourceFunc, GSourceFunc

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no param type for SourceFunc, GSourceFunc

// Unsupported : g_unix_signal_source_new : no return type

// Utf8Substring is a wrapper around the C function g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_start_pos := (C.glong)(startPos)

	c_end_pos := (C.glong)(endPos)

	retC := C.g_utf8_substring(c_str, c_start_pos, c_end_pos)
	retGo :=
		C.GoString(retC)

	return retGo
}

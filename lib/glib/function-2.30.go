// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_atomic_int_and : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_or : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_xor : unsupported parameter atomic : no param type

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no param type

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no param type

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no param type

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no param type

// Unsupported : g_compute_hmac_for_data : unsupported parameter digest_type : no param type

// Unsupported : g_compute_hmac_for_string : unsupported parameter digest_type : no param type

// Unsupported : g_dir_make_tmp : unsupported parameter tmpl : no param type

// FormatSize is a wrapper around the C function g_format_size.
func FormatSize(size uint64) {
	c_size := (C.guint64)(size)

	C.g_format_size(c_size)
}

// Unsupported : g_format_size_full : unsupported parameter flags : no param type

// Unsupported : g_mkdtemp : unsupported parameter tmpl : no param type

// Unsupported : g_mkdtemp_full : unsupported parameter tmpl : no param type

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no param type

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no param type

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no param type

// RegexEscapeNul is a wrapper around the C function g_regex_escape_nul.
func RegexEscapeNul(string string, length int32) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gint)(length)

	C.g_regex_escape_nul(c_string, c_length)
}

// TestFail is a wrapper around the C function g_test_fail.
func TestFail() {
	C.g_test_fail()
}

// Unsupported : g_unichar_compose : unsupported parameter ch : no param type

// Unsupported : g_unichar_decompose : unsupported parameter a : no param type

// Unsupported : g_unichar_fully_decompose : unsupported parameter compat : no param type

// UnicodeScriptFromIso15924 is a wrapper around the C function g_unicode_script_from_iso15924.
func UnicodeScriptFromIso15924(iso15924 uint32) {
	c_iso15924 := (C.guint32)(iso15924)

	C.g_unicode_script_from_iso15924(c_iso15924)
}

// Unsupported : g_unicode_script_to_iso15924 : unsupported parameter script : no param type

// Unsupported : g_unix_open_pipe : unsupported parameter fds : no param type

// Unsupported : g_unix_set_fd_nonblocking : unsupported parameter nonblock : no param type

// Unsupported : g_unix_signal_add : unsupported parameter handler : no param type

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no param type

// UnixSignalSourceNew is a wrapper around the C function g_unix_signal_source_new.
func UnixSignalSourceNew(signum int32) {
	c_signum := (C.gint)(signum)

	C.g_unix_signal_source_new(c_signum)
}

// Utf8Substring is a wrapper around the C function g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_start_pos := (C.glong)(startPos)

	c_end_pos := (C.glong)(endPos)

	C.g_utf8_substring(c_str, c_start_pos, c_end_pos)
}

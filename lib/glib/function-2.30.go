// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_atomic_int_and : unsupported parameter atomic : type guint, volatile guint*

// Unsupported function: g_atomic_int_or : unsupported parameter atomic : type guint, volatile guint*

// Unsupported function: g_atomic_int_xor : unsupported parameter atomic : type guint, volatile guint*

// Unsupported function: g_atomic_pointer_add : unsupported parameter atomic : type gpointer, void*

// Unsupported function: g_atomic_pointer_and : unsupported parameter atomic : type gpointer, void*

// Unsupported function: g_atomic_pointer_or : unsupported parameter atomic : type gpointer, void*

// Unsupported function: g_atomic_pointer_xor : unsupported parameter atomic : type gpointer, void*

// Unsupported function: g_compute_hmac_for_data : unsupported parameter digest_type : type ChecksumType, GChecksumType

// Unsupported function: g_compute_hmac_for_string : unsupported parameter digest_type : type ChecksumType, GChecksumType

// Unsupported function: g_dir_make_tmp : unsupported parameter tmpl : type filename, const gchar*

// FormatSize is a wrapper around the C function g_format_size.
func FormatSize(size uint64) {}

// Unsupported function: g_format_size_full : unsupported parameter flags : type FormatSizeFlags, GFormatSizeFlags

// Unsupported function: g_mkdtemp : unsupported parameter tmpl : type filename, gchar*

// Unsupported function: g_mkdtemp_full : unsupported parameter tmpl : type filename, gchar*

// Unsupported function: g_pointer_bit_lock : unsupported parameter address : type gpointer, void*

// Unsupported function: g_pointer_bit_trylock : unsupported parameter address : type gpointer, void*

// Unsupported function: g_pointer_bit_unlock : unsupported parameter address : type gpointer, void*

// RegexEscapeNul is a wrapper around the C function g_regex_escape_nul.
func RegexEscapeNul(string string, length int32) {}

// TestFail is a wrapper around the C function g_test_fail.
func TestFail() {}

// Unsupported function: g_unichar_compose : unsupported parameter ch : type gunichar, gunichar*

// Unsupported function: g_unichar_decompose : unsupported parameter a : type gunichar, gunichar*

// Unsupported function: g_unichar_fully_decompose : unsupported parameter compat : type gboolean, gboolean

// UnicodeScriptFromIso15924 is a wrapper around the C function g_unicode_script_from_iso15924.
func UnicodeScriptFromIso15924(iso15924 uint32) {}

// Unsupported function: g_unicode_script_to_iso15924 : unsupported parameter script : type UnicodeScript, GUnicodeScript

// Unsupported function: g_unix_open_pipe : unsupported parameter fds : type gint, gint*

// Unsupported function: g_unix_set_fd_nonblocking : unsupported parameter nonblock : type gboolean, gboolean

// Unsupported function: g_unix_signal_add : unsupported parameter handler : type SourceFunc, GSourceFunc

// Unsupported function: g_unix_signal_add_full : unsupported parameter handler : type SourceFunc, GSourceFunc

// UnixSignalSourceNew is a wrapper around the C function g_unix_signal_source_new.
func UnixSignalSourceNew(signum int32) {}

// Utf8Substring is a wrapper around the C function g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) {}

// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_get_user_special_dir : unsupported parameter directory : type UserDirectory, GUserDirectory

// Unsupported : g_once_init_enter : unsupported parameter location : type gpointer, void*

// Unsupported : g_once_init_leave : unsupported parameter location : type gpointer, void*

// Unsupported : g_regex_check_replacement : unsupported parameter has_references : type gboolean, gboolean*

// Unsupported : g_regex_escape_string : unsupported parameter string : no type

// Unsupported : g_regex_match_simple : unsupported parameter compile_options : type RegexCompileFlags, GRegexCompileFlags

// Unsupported : g_regex_split_simple : unsupported parameter compile_options : type RegexCompileFlags, GRegexCompileFlags

// Unsupported : g_sequence_get : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_insert_before : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_move : unsupported parameter src : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_move_range : unsupported parameter dest : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_range_get_midpoint : unsupported parameter begin : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_remove : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_remove_range : unsupported parameter begin : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_set : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported : g_sequence_swap : unsupported parameter a : type SequenceIter, GSequenceIter*

// Unsupported : g_slice_copy : unsupported parameter mem_block : type gpointer, gconstpointer

// Unsupported : g_timeout_add_seconds : unsupported parameter function : type SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : type SourceFunc, GSourceFunc

// TimeoutSourceNewSeconds is a wrapper around the C function g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) {
	c_interval := (C.guint)(interval)

	C.g_timeout_source_new_seconds(c_interval)
}

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) {
	c_uc := (C.gunichar)(uc)

	C.g_unichar_combining_class(c_uc)
}

// UnicharGetScript is a wrapper around the C function g_unichar_get_script.
func UnicharGetScript(ch rune) {
	c_ch := (C.gunichar)(ch)

	C.g_unichar_get_script(c_ch)
}

// UnicharIsmark is a wrapper around the C function g_unichar_ismark.
func UnicharIsmark(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_ismark(c_c)
}

// UnicharIszerowidth is a wrapper around the C function g_unichar_iszerowidth.
func UnicharIszerowidth(c rune) {
	c_c := (C.gunichar)(c)

	C.g_unichar_iszerowidth(c_c)
}

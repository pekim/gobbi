// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported function: g_get_user_special_dir : unsupported parameter directory : type UserDirectory, GUserDirectory

// Blacklisted function: g_once_init_enter

// Blacklisted function: g_once_init_leave

// Unsupported function: g_regex_check_replacement : unsupported parameter has_references : type gboolean, gboolean*

// Unsupported function: g_regex_escape_string : unsupported parameter string : no type

// Unsupported function: g_regex_match_simple : unsupported parameter compile_options : type RegexCompileFlags, GRegexCompileFlags

// Unsupported function: g_regex_split_simple : unsupported parameter compile_options : type RegexCompileFlags, GRegexCompileFlags

// Unsupported function: g_sequence_get : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_insert_before : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_move : unsupported parameter src : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_move_range : unsupported parameter dest : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_range_get_midpoint : unsupported parameter begin : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_remove : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_remove_range : unsupported parameter begin : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_set : unsupported parameter iter : type SequenceIter, GSequenceIter*

// Unsupported function: g_sequence_swap : unsupported parameter a : type SequenceIter, GSequenceIter*

// Unsupported function: g_slice_copy : unsupported parameter mem_block : type gpointer, gconstpointer

// Unsupported function: g_timeout_add_seconds : unsupported parameter function : type SourceFunc, GSourceFunc

// Unsupported function: g_timeout_add_seconds_full : unsupported parameter function : type SourceFunc, GSourceFunc

// TimeoutSourceNewSeconds is a wrapper around the C function g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) {}

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) {}

// UnicharGetScript is a wrapper around the C function g_unichar_get_script.
func UnicharGetScript(ch rune) {}

// UnicharIsmark is a wrapper around the C function g_unichar_ismark.
func UnicharIsmark(c rune) {}

// UnicharIszerowidth is a wrapper around the C function g_unichar_iszerowidth.
func UnicharIszerowidth(c rune) {}

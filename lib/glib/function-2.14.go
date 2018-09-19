// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// GetUserSpecialDir is a wrapper around the C function g_get_user_special_dir.
func GetUserSpecialDir(directory UserDirectory) string {
	c_directory := (C.GUserDirectory)(directory)

	retC := C.g_get_user_special_dir(c_directory)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer, void*

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer, void*

// Unsupported : g_regex_check_replacement : unsupported parameter has_references : no type generator for gboolean, gboolean*

// Unsupported : g_regex_escape_string : unsupported parameter string : no param type

// Unsupported : g_regex_match_simple : unsupported parameter compile_options : no type generator for RegexCompileFlags, GRegexCompileFlags

// Unsupported : g_regex_split_simple : unsupported parameter compile_options : no type generator for RegexCompileFlags, GRegexCompileFlags

// Unsupported : g_sequence_get : unsupported parameter iter : record param - coming soon

// Unsupported : g_sequence_insert_before : unsupported parameter iter : record param - coming soon

// Unsupported : g_sequence_move : unsupported parameter src : record param - coming soon

// Unsupported : g_sequence_move_range : unsupported parameter dest : record param - coming soon

// Unsupported : g_sequence_range_get_midpoint : unsupported parameter begin : record param - coming soon

// Unsupported : g_sequence_remove : unsupported parameter iter : record param - coming soon

// Unsupported : g_sequence_remove_range : unsupported parameter begin : record param - coming soon

// Unsupported : g_sequence_set : unsupported parameter iter : record param - coming soon

// Unsupported : g_sequence_swap : unsupported parameter a : record param - coming soon

// SliceCopy is a wrapper around the C function g_slice_copy.
func SliceCopy(blockSize uint64, memBlock uintptr) uintptr {
	c_block_size := (C.gsize)(blockSize)

	c_mem_block := (C.gconstpointer)(memBlock)

	retC := C.g_slice_copy(c_block_size, c_mem_block)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// TimeoutSourceNewSeconds is a wrapper around the C function g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new_seconds(c_interval)
	retGo := sourceNewFromC(retC)

	return retGo
}

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) int32 {
	c_uc := (C.gunichar)(uc)

	retC := C.g_unichar_combining_class(c_uc)
	retGo := (int32)(retC)

	return retGo
}

// UnicharGetScript is a wrapper around the C function g_unichar_get_script.
func UnicharGetScript(ch rune) UnicodeScript {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_get_script(c_ch)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// Unsupported : g_unichar_ismark : no return generator

// Unsupported : g_unichar_iszerowidth : no return generator

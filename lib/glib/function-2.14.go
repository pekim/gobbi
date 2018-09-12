// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_get_user_special_dir : unsupported parameter directory : no param type for UserDirectory, GUserDirectory

// Unsupported : g_once_init_enter : no return type

// Unsupported : g_once_init_leave : no return type

// Unsupported : g_regex_check_replacement : unsupported parameter has_references : no param type for gboolean, gboolean*

// Unsupported : g_regex_escape_string : unsupported parameter string : no param type

// Unsupported : g_regex_match_simple : unsupported parameter compile_options : no param type for RegexCompileFlags, GRegexCompileFlags

// Unsupported : g_regex_split_simple : unsupported parameter compile_options : no param type for RegexCompileFlags, GRegexCompileFlags

// Unsupported : g_sequence_get : unsupported parameter iter : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_insert_before : unsupported parameter iter : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_move : unsupported parameter src : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_move_range : unsupported parameter dest : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_range_get_midpoint : unsupported parameter begin : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_remove : unsupported parameter iter : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_remove_range : unsupported parameter begin : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_set : unsupported parameter iter : no param type for SequenceIter, GSequenceIter*

// Unsupported : g_sequence_swap : unsupported parameter a : no param type for SequenceIter, GSequenceIter*

// SliceCopy is a wrapper around the C function g_slice_copy.
func SliceCopy(blockSize uint64, memBlock uintptr) uintptr {
	c_block_size := (C.gsize)(blockSize)

	c_mem_block := (C.gpointer)(memBlock)

	retC := C.g_slice_copy(c_block_size, c_mem_block)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no param type for SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no param type for SourceFunc, GSourceFunc

// Unsupported : g_timeout_source_new_seconds : no return type

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) int32 {
	c_uc := (C.gunichar)(uc)

	retC := C.g_unichar_combining_class(c_uc)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_unichar_get_script : no return type

// Unsupported : g_unichar_ismark : no return type

// Unsupported : g_unichar_iszerowidth : no return type

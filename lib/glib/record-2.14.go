// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Blacklisted : g_hash_table_get_keys

// Blacklisted : g_hash_table_get_values

// Blacklisted : g_key_file_load_from_dirs

// Blacklisted : g_match_info_expand_references

// Blacklisted : g_match_info_fetch

// Blacklisted : g_match_info_fetch_all

// Blacklisted : g_match_info_fetch_named

// Blacklisted : g_match_info_fetch_named_pos

// Blacklisted : g_match_info_fetch_pos

// Blacklisted : g_match_info_free

// Blacklisted : g_match_info_get_match_count

// Unsupported : g_match_info_get_regex : return type : Blacklisted record : GRegex

// Blacklisted : g_match_info_get_string

// Blacklisted : g_match_info_is_partial_match

// Blacklisted : g_match_info_matches

// Blacklisted : g_match_info_next

// g_once_init_enter : unsupported parameter location : no type generator for gpointer (void*) for param location
// g_once_init_leave : unsupported parameter location : no type generator for gpointer (void*) for param location
// Blacklisted : g_option_context_get_help

// Blacklisted : g_queue_clear

// Init is a wrapper around the C function g_queue_init.
func (recv *Queue) Init() {
	C.g_queue_init((*C.GQueue)(recv.native))

	return
}

// Blacklisted : GRegex

// g_sequence_foreach_range : unsupported parameter func : no type generator for Func (GFunc) for param func
// g_sequence_get : no return generator
// g_sequence_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// Blacklisted : g_sequence_move

// Blacklisted : g_sequence_move_range

// g_sequence_new : unsupported parameter data_destroy : no type generator for DestroyNotify (GDestroyNotify) for param data_destroy
// Blacklisted : g_sequence_range_get_midpoint

// Blacklisted : g_sequence_remove

// Blacklisted : g_sequence_remove_range

// g_sequence_set : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// g_sequence_sort_changed : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func
// g_sequence_sort_changed_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param iter_cmp
// Blacklisted : g_sequence_swap

// Unsupported : g_sequence_append : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func

// Blacklisted : g_sequence_free

// Blacklisted : g_sequence_get_begin_iter

// Blacklisted : g_sequence_get_end_iter

// Blacklisted : g_sequence_get_iter_at_pos

// Blacklisted : g_sequence_get_length

// Unsupported : g_sequence_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_prepend : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_search : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_search_iter : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param cmp_func

// Blacklisted : g_sequence_iter_compare

// Blacklisted : g_sequence_iter_get_position

// Blacklisted : g_sequence_iter_get_sequence

// Blacklisted : g_sequence_iter_is_begin

// Blacklisted : g_sequence_iter_is_end

// Blacklisted : g_sequence_iter_move

// Blacklisted : g_sequence_iter_next

// Blacklisted : g_sequence_iter_prev

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : g_string_overwrite

// Blacklisted : g_string_overwrite_len

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : g_string_chunk_clear

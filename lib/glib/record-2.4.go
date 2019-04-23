// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// g_array_remove_range : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_hash_table_find : unsupported parameter predicate : no type generator for HRFunc (GHRFunc) for param predicate
// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc (GCopyFunc) for param copy_func

// Once is a wrapper around the C record GOnce.
type Once struct {
	native *C.GOnce
	Status OnceStatus
	// retval : no type generator for gpointer, volatile gpointer
}

func OnceNewFromC(u unsafe.Pointer) *Once {
	c := (*C.GOnce)(u)
	if c == nil {
		return nil
	}

	g := &Once{
		Status: (OnceStatus)(c.status),
		native: c,
	}

	return g
}

func (recv *Once) ToC() unsafe.Pointer {
	recv.native.status =
		(C.GOnceStatus)(recv.Status)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Once with another Once, and returns true if they represent the same GObject.
func (recv *Once) Equals(other *Once) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_once_impl : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Blacklisted : g_queue_copy

// Blacklisted : g_queue_delete_link

// Unsupported : g_queue_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func

// Blacklisted : g_queue_get_length

// Unsupported : g_queue_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_insert_after : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_queue_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_queue_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_queue_link_index

// Blacklisted : g_queue_peek_head_link

// Unsupported : g_queue_peek_nth : no return generator

// Blacklisted : g_queue_peek_nth_link

// Blacklisted : g_queue_peek_tail_link

// Unsupported : g_queue_pop_nth : no return generator

// Blacklisted : g_queue_pop_nth_link

// Unsupported : g_queue_push_nth : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Blacklisted : g_queue_push_nth_link

// Unsupported : g_queue_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Blacklisted : g_queue_reverse

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func

// Blacklisted : g_queue_unlink

// Blacklisted : g_rand_new_with_seed_array

// Blacklisted : g_rand_copy

// Blacklisted : g_rand_set_seed_array

// Blacklisted : g_string_chunk_insert_len

// Blacklisted : g_timer_continue

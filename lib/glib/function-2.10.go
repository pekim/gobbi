// +build glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_hash_table_unref : unsupported parameter hash_table : type GLib.HashTable, GHashTable*

// InternStaticString is a wrapper around the C function g_intern_static_string.
func InternStaticString(string string) {}

// InternString is a wrapper around the C function g_intern_string.
func InternString(string string) {}

// SliceAlloc is a wrapper around the C function g_slice_alloc.
func SliceAlloc(blockSize uint64) {}

// SliceAlloc0 is a wrapper around the C function g_slice_alloc0.
func SliceAlloc0(blockSize uint64) {}

// SliceFree1 is a wrapper around the C function g_slice_free1.
func SliceFree1(blockSize uint64, memBlock uintptr) {}

// SliceFreeChainWithOffset is a wrapper around the C function g_slice_free_chain_with_offset.
func SliceFreeChainWithOffset(blockSize uint64, memChain uintptr, nextOffset uint64) {}

// ThreadPoolGetMaxIdleTime is a wrapper around the C function g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() {}

// ThreadPoolSetMaxIdleTime is a wrapper around the C function g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {}

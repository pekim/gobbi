// This is a generated file - DO NOT EDIT
// +build glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Atomically decrements the reference count of @hash_table by one.
// If the reference count drops to 0, all keys and values will be
// destroyed, and all memory allocated by the hash table is released.
// This function is MT-safe and may be called from any thread.
/*

C function : g_hash_table_unref
*/
func HashTableUnref(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_unref(c_hash_table)

	return
}

// Returns a canonical representation for @string. Interned strings
// can be compared for equality by comparing the pointers, instead of
// using strcmp(). g_intern_static_string() does not copy the string,
// therefore @string must not be freed or modified.
/*

C function : g_intern_static_string
*/
func InternStaticString(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_static_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Returns a canonical representation for @string. Interned strings
// can be compared for equality by comparing the pointers, instead of
// using strcmp().
/*

C function : g_intern_string
*/
func InternString(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// Allocates a block of memory from the slice allocator.
// The block address handed out can be expected to be aligned
// to at least 1 * sizeof (void*),
// though in general slices are 2 * sizeof (void*) bytes aligned,
// if a malloc() fallback implementation is used instead,
// the alignment may be reduced in a libc dependent fashion.
// Note that the underlying slice allocation mechanism can
// be changed with the [`G_SLICE=always-malloc`][G_SLICE]
// environment variable.
/*

C function : g_slice_alloc
*/
func SliceAlloc(blockSize uint64) uintptr {
	c_block_size := (C.gsize)(blockSize)

	retC := C.g_slice_alloc(c_block_size)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Allocates a block of memory via g_slice_alloc() and initializes
// the returned memory to 0. Note that the underlying slice allocation
// mechanism can be changed with the [`G_SLICE=always-malloc`][G_SLICE]
// environment variable.
/*

C function : g_slice_alloc0
*/
func SliceAlloc0(blockSize uint64) uintptr {
	c_block_size := (C.gsize)(blockSize)

	retC := C.g_slice_alloc0(c_block_size)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Frees a block of memory.
//
// The memory must have been allocated via g_slice_alloc() or
// g_slice_alloc0() and the @block_size has to match the size
// specified upon allocation. Note that the exact release behaviour
// can be changed with the [`G_DEBUG=gc-friendly`][G_DEBUG] environment
// variable, also see [`G_SLICE`][G_SLICE] for related debugging options.
//
// If @mem_block is %NULL, this function does nothing.
/*

C function : g_slice_free1
*/
func SliceFree1(blockSize uint64, memBlock uintptr) {
	c_block_size := (C.gsize)(blockSize)

	c_mem_block := (C.gpointer)(memBlock)

	C.g_slice_free1(c_block_size, c_mem_block)

	return
}

// Frees a linked list of memory blocks of structure type @type.
//
// The memory blocks must be equal-sized, allocated via
// g_slice_alloc() or g_slice_alloc0() and linked together by a
// @next pointer (similar to #GSList). The offset of the @next
// field in each block is passed as third argument.
// Note that the exact release behaviour can be changed with the
// [`G_DEBUG=gc-friendly`][G_DEBUG] environment variable, also see
// [`G_SLICE`][G_SLICE] for related debugging options.
//
// If @mem_chain is %NULL, this function does nothing.
/*

C function : g_slice_free_chain_with_offset
*/
func SliceFreeChainWithOffset(blockSize uint64, memChain uintptr, nextOffset uint64) {
	c_block_size := (C.gsize)(blockSize)

	c_mem_chain := (C.gpointer)(memChain)

	c_next_offset := (C.gsize)(nextOffset)

	C.g_slice_free_chain_with_offset(c_block_size, c_mem_chain, c_next_offset)

	return
}

// This function will return the maximum @interval that a
// thread will wait in the thread pool for new tasks before
// being stopped.
//
// If this function returns 0, threads waiting in the thread
// pool for new work are not stopped.
/*

C function : g_thread_pool_get_max_idle_time
*/
func ThreadPoolGetMaxIdleTime() uint32 {
	retC := C.g_thread_pool_get_max_idle_time()
	retGo := (uint32)(retC)

	return retGo
}

// This function will set the maximum @interval that a thread
// waiting in the pool for new tasks can be idle for before
// being stopped. This function is similar to calling
// g_thread_pool_stop_unused_threads() on a regular timeout,
// except this is done on a per thread basis.
//
// By setting @interval to 0, idle threads will not be stopped.
//
// The default value is 15000 (15 seconds).
/*

C function : g_thread_pool_set_max_idle_time
*/
func ThreadPoolSetMaxIdleTime(interval uint32) {
	c_interval := (C.guint)(interval)

	C.g_thread_pool_set_max_idle_time(c_interval)

	return
}

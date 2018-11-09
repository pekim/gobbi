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

// HashTableUnref is a wrapper around the C function g_hash_table_unref.
func HashTableUnref(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	C.g_hash_table_unref(c_hash_table)

	return
}

// InternStaticString is a wrapper around the C function g_intern_static_string.
func InternStaticString(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_static_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// InternString is a wrapper around the C function g_intern_string.
func InternString(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_intern_string(c_string)
	retGo := C.GoString(retC)

	return retGo
}

// SliceAlloc is a wrapper around the C function g_slice_alloc.
func SliceAlloc(blockSize uint64) uintptr {
	c_block_size := (C.gsize)(blockSize)

	retC := C.g_slice_alloc(c_block_size)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// SliceAlloc0 is a wrapper around the C function g_slice_alloc0.
func SliceAlloc0(blockSize uint64) uintptr {
	c_block_size := (C.gsize)(blockSize)

	retC := C.g_slice_alloc0(c_block_size)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// SliceFree1 is a wrapper around the C function g_slice_free1.
func SliceFree1(blockSize uint64, memBlock uintptr) {
	c_block_size := (C.gsize)(blockSize)

	c_mem_block := (C.gpointer)(memBlock)

	C.g_slice_free1(c_block_size, c_mem_block)

	return
}

// SliceFreeChainWithOffset is a wrapper around the C function g_slice_free_chain_with_offset.
func SliceFreeChainWithOffset(blockSize uint64, memChain uintptr, nextOffset uint64) {
	c_block_size := (C.gsize)(blockSize)

	c_mem_chain := (C.gpointer)(memChain)

	c_next_offset := (C.gsize)(nextOffset)

	C.g_slice_free_chain_with_offset(c_block_size, c_mem_chain, c_next_offset)

	return
}

// ThreadPoolGetMaxIdleTime is a wrapper around the C function g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() uint32 {
	retC := C.g_thread_pool_get_max_idle_time()
	retGo := (uint32)(retC)

	return retGo
}

// ThreadPoolSetMaxIdleTime is a wrapper around the C function g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {
	c_interval := (C.guint)(interval)

	C.g_thread_pool_set_max_idle_time(c_interval)

	return
}

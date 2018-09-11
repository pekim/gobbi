// +build glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_hash_table_unref : unsupported parameter hash_table : no param type

// InternStaticString is a wrapper around the C function g_intern_static_string.
func InternStaticString(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_intern_static_string(c_string)
}

// InternString is a wrapper around the C function g_intern_string.
func InternString(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_intern_string(c_string)
}

// SliceAlloc is a wrapper around the C function g_slice_alloc.
func SliceAlloc(blockSize uint64) {
	c_block_size := (C.gsize)(blockSize)

	C.g_slice_alloc(c_block_size)
}

// SliceAlloc0 is a wrapper around the C function g_slice_alloc0.
func SliceAlloc0(blockSize uint64) {
	c_block_size := (C.gsize)(blockSize)

	C.g_slice_alloc0(c_block_size)
}

// SliceFree1 is a wrapper around the C function g_slice_free1.
func SliceFree1(blockSize uint64, memBlock uintptr) {
	c_block_size := (C.gsize)(blockSize)

	c_mem_block := (C.gpointer)(memBlock)

	C.g_slice_free1(c_block_size, c_mem_block)
}

// SliceFreeChainWithOffset is a wrapper around the C function g_slice_free_chain_with_offset.
func SliceFreeChainWithOffset(blockSize uint64, memChain uintptr, nextOffset uint64) {
	c_block_size := (C.gsize)(blockSize)

	c_mem_chain := (C.gpointer)(memChain)

	c_next_offset := (C.gsize)(nextOffset)

	C.g_slice_free_chain_with_offset(c_block_size, c_mem_chain, c_next_offset)
}

// ThreadPoolGetMaxIdleTime is a wrapper around the C function g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() {
	C.g_thread_pool_get_max_idle_time()
}

// ThreadPoolSetMaxIdleTime is a wrapper around the C function g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {
	c_interval := (C.guint)(interval)

	C.g_thread_pool_set_max_idle_time(c_interval)
}

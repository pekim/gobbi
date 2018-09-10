// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bit_lock : unsupported parameter address : type gint, volatile gint*

// Unsupported : g_bit_trylock : unsupported parameter address : type gint, volatile gint*

// Unsupported : g_bit_unlock : unsupported parameter address : type gint, volatile gint*

// Malloc0N is a wrapper around the C function g_malloc0_n.
func Malloc0N(nBlocks uint64, nBlockBytes uint64) {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	C.g_malloc0_n(c_n_blocks, c_n_block_bytes)
}

// MallocN is a wrapper around the C function g_malloc_n.
func MallocN(nBlocks uint64, nBlockBytes uint64) {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	C.g_malloc_n(c_n_blocks, c_n_block_bytes)
}

// ReallocN is a wrapper around the C function g_realloc_n.
func ReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	C.g_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
}

// TryMalloc0N is a wrapper around the C function g_try_malloc0_n.
func TryMalloc0N(nBlocks uint64, nBlockBytes uint64) {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	C.g_try_malloc0_n(c_n_blocks, c_n_block_bytes)
}

// TryMallocN is a wrapper around the C function g_try_malloc_n.
func TryMallocN(nBlocks uint64, nBlockBytes uint64) {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	C.g_try_malloc_n(c_n_blocks, c_n_block_bytes)
}

// TryReallocN is a wrapper around the C function g_try_realloc_n.
func TryReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	C.g_try_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
}

// VariantIsObjectPath is a wrapper around the C function g_variant_is_object_path.
func VariantIsObjectPath(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_variant_is_object_path(c_string)
}

// VariantIsSignature is a wrapper around the C function g_variant_is_signature.
func VariantIsSignature(string string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_variant_is_signature(c_string)
}

// VariantTypeStringScan is a wrapper around the C function g_variant_type_string_scan.
func VariantTypeStringScan(string string, limit string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_limit := C.CString(limit)
	defer C.free(unsafe.Pointer(c_limit))

	C.g_variant_type_string_scan(c_string, c_limit, c_endptr)
}

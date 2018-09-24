// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bit_lock : unsupported parameter address : no type generator for gint, volatile gint*

// Unsupported : g_bit_trylock : unsupported parameter address : no type generator for gint, volatile gint*

// Unsupported : g_bit_unlock : unsupported parameter address : no type generator for gint, volatile gint*

// Malloc0N is a wrapper around the C function g_malloc0_n.
func Malloc0N(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_malloc0_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(retC)

	return retGo
}

// MallocN is a wrapper around the C function g_malloc_n.
func MallocN(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_malloc_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(retC)

	return retGo
}

// ReallocN is a wrapper around the C function g_realloc_n.
func ReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(retC)

	return retGo
}

// TryMalloc0N is a wrapper around the C function g_try_malloc0_n.
func TryMalloc0N(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_malloc0_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(retC)

	return retGo
}

// TryMallocN is a wrapper around the C function g_try_malloc_n.
func TryMallocN(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_malloc_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(retC)

	return retGo
}

// TryReallocN is a wrapper around the C function g_try_realloc_n.
func TryReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(retC)

	return retGo
}

// VariantIsObjectPath is a wrapper around the C function g_variant_is_object_path.
func VariantIsObjectPath(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_object_path(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// VariantIsSignature is a wrapper around the C function g_variant_is_signature.
func VariantIsSignature(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_signature(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// VariantTypeStringScan is a wrapper around the C function g_variant_type_string_scan.
func VariantTypeStringScan(string string, limit string) (bool, string) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_limit := C.CString(limit)
	defer C.free(unsafe.Pointer(c_limit))

	var c_endptr *C.gchar

	retC := C.g_variant_type_string_scan(c_string, c_limit, &c_endptr)
	retGo := retC == C.TRUE

	endptr := C.GoString(c_endptr)
	defer C.free(unsafe.Pointer(c_endptr))

	return retGo, endptr
}

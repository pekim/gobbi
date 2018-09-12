// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bit_lock : unsupported parameter address : no param type for gint, volatile gint*

// Unsupported : g_bit_trylock : unsupported parameter address : no param type for gint, volatile gint*

// Unsupported : g_bit_unlock : unsupported parameter address : no param type for gint, volatile gint*

// Malloc0N is a wrapper around the C function g_malloc0_n.
func Malloc0N(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_malloc0_n(c_n_blocks, c_n_block_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// MallocN is a wrapper around the C function g_malloc_n.
func MallocN(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_malloc_n(c_n_blocks, c_n_block_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// ReallocN is a wrapper around the C function g_realloc_n.
func ReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// TryMalloc0N is a wrapper around the C function g_try_malloc0_n.
func TryMalloc0N(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_malloc0_n(c_n_blocks, c_n_block_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// TryMallocN is a wrapper around the C function g_try_malloc_n.
func TryMallocN(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_malloc_n(c_n_blocks, c_n_block_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// TryReallocN is a wrapper around the C function g_try_realloc_n.
func TryReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
	retGo :=
		(uintptr)(retC)

	return retGo
}

// Unsupported : g_variant_is_object_path : no return type

// Unsupported : g_variant_is_signature : no return type

// Unsupported : g_variant_type_string_scan : no return type

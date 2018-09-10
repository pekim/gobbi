// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Blacklisted function: g_bit_lock

// Blacklisted function: g_bit_trylock

// Blacklisted function: g_bit_unlock

// Blacklisted function: g_malloc0_n

// Blacklisted function: g_malloc_n

// ReallocN is a wrapper around the C function g_realloc_n.
func ReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) {}

// TryMalloc0N is a wrapper around the C function g_try_malloc0_n.
func TryMalloc0N(nBlocks uint64, nBlockBytes uint64) {}

// TryMallocN is a wrapper around the C function g_try_malloc_n.
func TryMallocN(nBlocks uint64, nBlockBytes uint64) {}

// TryReallocN is a wrapper around the C function g_try_realloc_n.
func TryReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) {}

// VariantIsObjectPath is a wrapper around the C function g_variant_is_object_path.
func VariantIsObjectPath(string string) {}

// VariantIsSignature is a wrapper around the C function g_variant_is_signature.
func VariantIsSignature(string string) {}

// VariantTypeStringScan is a wrapper around the C function g_variant_type_string_scan.
func VariantTypeStringScan(string string, limit string, endptr string) {}

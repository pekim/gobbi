// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Sets the indicated @lock_bit in @address.  If the bit is already
// set, this call will block until g_bit_unlock() unsets the
// corresponding bit.
//
// Attempting to lock on two different bits within the same integer is
// not supported and will very probably cause deadlocks.
//
// The value of the bit that is set is (1u << @bit).  If @bit is not
// between 0 and 31 then the result is undefined.
//
// This function accesses @address atomically.  All other accesses to
// @address must be atomic in order for this function to work
// reliably.
/*

C function : g_bit_lock
*/
func BitLock(address int32, lockBit int32) {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	C.g_bit_lock(&c_address, c_lock_bit)

	return
}

// Sets the indicated @lock_bit in @address, returning %TRUE if
// successful.  If the bit is already set, returns %FALSE immediately.
//
// Attempting to lock on two different bits within the same integer is
// not supported.
//
// The value of the bit that is set is (1u << @bit).  If @bit is not
// between 0 and 31 then the result is undefined.
//
// This function accesses @address atomically.  All other accesses to
// @address must be atomic in order for this function to work
// reliably.
/*

C function : g_bit_trylock
*/
func BitTrylock(address int32, lockBit int32) bool {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	retC := C.g_bit_trylock(&c_address, c_lock_bit)
	retGo := retC == C.TRUE

	return retGo
}

// Clears the indicated @lock_bit in @address.  If another thread is
// currently blocked in g_bit_lock() on this same bit then it will be
// woken up.
//
// This function accesses @address atomically.  All other accesses to
// @address must be atomic in order for this function to work
// reliably.
/*

C function : g_bit_unlock
*/
func BitUnlock(address int32, lockBit int32) {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	C.g_bit_unlock(&c_address, c_lock_bit)

	return
}

// This function is similar to g_malloc0(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
/*

C function : g_malloc0_n
*/
func Malloc0N(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_malloc0_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This function is similar to g_malloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
/*

C function : g_malloc_n
*/
func MallocN(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_malloc_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This function is similar to g_realloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
/*

C function : g_realloc_n
*/
func ReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This function is similar to g_try_malloc0(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
/*

C function : g_try_malloc0_n
*/
func TryMalloc0N(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_malloc0_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This function is similar to g_try_malloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
/*

C function : g_try_malloc_n
*/
func TryMallocN(nBlocks uint64, nBlockBytes uint64) uintptr {
	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_malloc_n(c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// This function is similar to g_try_realloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
/*

C function : g_try_realloc_n
*/
func TryReallocN(mem uintptr, nBlocks uint64, nBlockBytes uint64) uintptr {
	c_mem := (C.gpointer)(mem)

	c_n_blocks := (C.gsize)(nBlocks)

	c_n_block_bytes := (C.gsize)(nBlockBytes)

	retC := C.g_try_realloc_n(c_mem, c_n_blocks, c_n_block_bytes)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Determines if a given string is a valid D-Bus object path.  You
// should ensure that a string is a valid D-Bus object path before
// passing it to g_variant_new_object_path().
//
// A valid object path starts with '/' followed by zero or more
// sequences of characters separated by '/' characters.  Each sequence
// must contain only the characters "[A-Z][a-z][0-9]_".  No sequence
// (including the one following the final '/' character) may be empty.
/*

C function : g_variant_is_object_path
*/
func VariantIsObjectPath(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_object_path(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if a given string is a valid D-Bus type signature.  You
// should ensure that a string is a valid D-Bus type signature before
// passing it to g_variant_new_signature().
//
// D-Bus type signatures consist of zero or more definite #GVariantType
// strings in sequence.
/*

C function : g_variant_is_signature
*/
func VariantIsSignature(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_variant_is_signature(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Scan for a single complete and valid GVariant type string in @string.
// The memory pointed to by @limit (or bytes beyond it) is never
// accessed.
//
// If a valid type string is found, @endptr is updated to point to the
// first character past the end of the string that was found and %TRUE
// is returned.
//
// If there is no valid type string starting at @string, or if the type
// string does not end before @limit then %FALSE is returned.
//
// For the simple case of checking if a string is a valid type string,
// see g_variant_type_string_is_valid().
/*

C function : g_variant_type_string_scan
*/
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

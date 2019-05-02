// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "C"

// BitLock is a wrapper around the C function g_bit_lock.
func BitLock(address int32, lockBit int32) {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	C.g_bit_lock(&c_address, c_lock_bit)

	return
}

// BitTrylock is a wrapper around the C function g_bit_trylock.
func BitTrylock(address int32, lockBit int32) bool {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	retC := C.g_bit_trylock(&c_address, c_lock_bit)
	retGo := retC == C.TRUE

	return retGo
}

// BitUnlock is a wrapper around the C function g_bit_unlock.
func BitUnlock(address int32, lockBit int32) {
	c_address := (C.gint)(address)

	c_lock_bit := (C.gint)(lockBit)

	C.g_bit_unlock(&c_address, c_lock_bit)

	return
}

// Unsupported : g_malloc0_n : no return generator

// Unsupported : g_malloc_n : no return generator

// Unsupported : g_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_try_malloc0_n : no return generator

// Unsupported : g_try_malloc_n : no return generator

// Unsupported : g_try_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

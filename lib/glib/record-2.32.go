// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Bytes is a wrapper around the C record GBytes.
type Bytes struct {
	native *C.GBytes
}

func bytesNewFromC(c *C.GBytes) *Bytes {
	if c == nil {
		return nil
	}

	g := &Bytes{native: c}

	return g
}

// Unsupported : g_bytes_new : unsupported parameter data : no param type

// Unsupported : g_bytes_new_static : unsupported parameter data : no param type

// Unsupported : g_bytes_new_take : unsupported parameter data : no param type

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no param type

// Compare is a wrapper around the C function g_bytes_compare.
func Compare(bytes2 uintptr) int32 {
	c_bytes2 := (C.gconstpointer)(bytes2)

	retC := C.g_bytes_compare(c_bytes2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_bytes_equal : no return generator

// Unsupported : g_bytes_get_data : unsupported parameter size : no type generator for gsize, gsize*

// GetSize is a wrapper around the C function g_bytes_get_size.
func GetSize() uint64 {
	retC := C.g_bytes_get_size()
	retGo := (uint64)(retC)

	return retGo
}

// Hash is a wrapper around the C function g_bytes_hash.
func Hash() uint32 {
	retC := C.g_bytes_hash()
	retGo := (uint32)(retC)

	return retGo
}

// NewFromBytes is a wrapper around the C function g_bytes_new_from_bytes.
func NewFromBytes(offset uint64, length uint64) *Bytes {
	c_offset := (C.gsize)(offset)

	c_length := (C.gsize)(length)

	retC := C.g_bytes_new_from_bytes(c_offset, c_length)
	retGo := bytesNewFromC(retC)

	return retGo
}

// Ref is a wrapper around the C function g_bytes_ref.
func Ref() *Bytes {
	retC := C.g_bytes_ref()
	retGo := bytesNewFromC(retC)

	return retGo
}

// Unsupported : g_bytes_unref : no return generator

// Unsupported : g_bytes_unref_to_array : no return type

// Unsupported : g_bytes_unref_to_data : unsupported parameter size : no type generator for gsize, gsize*

// RWLock is a wrapper around the C record GRWLock.
type RWLock struct {
	native *C.GRWLock
	P      uintptr
	// no type for i
}

func rWLockNewFromC(c *C.GRWLock) *RWLock {
	if c == nil {
		return nil
	}

	g := &RWLock{
		P:      (uintptr)(c.p),
		native: c,
	}

	return g
}

// Unsupported : g_rw_lock_clear : no return generator

// Unsupported : g_rw_lock_init : no return generator

// Unsupported : g_rw_lock_reader_lock : no return generator

// Unsupported : g_rw_lock_reader_trylock : no return generator

// Unsupported : g_rw_lock_reader_unlock : no return generator

// Unsupported : g_rw_lock_writer_lock : no return generator

// Unsupported : g_rw_lock_writer_trylock : no return generator

// Unsupported : g_rw_lock_writer_unlock : no return generator

// RecMutex is a wrapper around the C record GRecMutex.
type RecMutex struct {
	native *C.GRecMutex
	P      uintptr
	// no type for i
}

func recMutexNewFromC(c *C.GRecMutex) *RecMutex {
	if c == nil {
		return nil
	}

	g := &RecMutex{
		P:      (uintptr)(c.p),
		native: c,
	}

	return g
}

// Unsupported : g_rec_mutex_clear : no return generator

// Unsupported : g_rec_mutex_init : no return generator

// Unsupported : g_rec_mutex_lock : no return generator

// Unsupported : g_rec_mutex_trylock : no return generator

// Unsupported : g_rec_mutex_unlock : no return generator

// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// g_array_set_clear_func : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// Bytes is a wrapper around the C record GBytes.
type Bytes struct {
	native *C.GBytes
}

func BytesNewFromC(u unsafe.Pointer) *Bytes {
	c := (*C.GBytes)(u)
	if c == nil {
		return nil
	}

	g := &Bytes{native: c}

	return g
}

func (recv *Bytes) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Bytes with another Bytes, and returns true if they represent the same GObject.
func (recv *Bytes) Equals(other *Bytes) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_bytes_new

// Blacklisted : g_bytes_new_static

// Blacklisted : g_bytes_new_take

// Unsupported : g_bytes_new_with_free_func : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// Blacklisted : g_bytes_compare

// Blacklisted : g_bytes_equal

// Unsupported : g_bytes_get_data : array return type :

// Blacklisted : g_bytes_get_size

// Blacklisted : g_bytes_hash

// Blacklisted : g_bytes_new_from_bytes

// Blacklisted : g_bytes_ref

// Blacklisted : g_bytes_unref

// Unsupported : g_bytes_unref_to_array : array return type :

// Unsupported : g_bytes_unref_to_data : array return type :

// Blacklisted : g_cond_clear

// Init is a wrapper around the C function g_cond_init.
func (recv *Cond) Init() {
	C.g_cond_init((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// g_hash_table_add : unsupported parameter key : no type generator for gpointer (gpointer) for param key
// g_hash_table_contains : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// Blacklisted : g_key_file_ref

// Blacklisted : g_key_file_unref

// Blacklisted : g_main_context_ref_thread_default

// Blacklisted : g_mapped_file_new_from_fd

// Unsupported : g_private_replace : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Unsupported : g_queue_free_full : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// RWLock is a wrapper around the C record GRWLock.
type RWLock struct {
	native *C.GRWLock
	// Private : p
	// Private : i
}

func RWLockNewFromC(u unsafe.Pointer) *RWLock {
	c := (*C.GRWLock)(u)
	if c == nil {
		return nil
	}

	g := &RWLock{native: c}

	return g
}

func (recv *RWLock) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RWLock with another RWLock, and returns true if they represent the same GObject.
func (recv *RWLock) Equals(other *RWLock) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_rw_lock_clear

// Init is a wrapper around the C function g_rw_lock_init.
func (recv *RWLock) Init() {
	C.g_rw_lock_init((*C.GRWLock)(recv.native))

	return
}

// Blacklisted : g_rw_lock_reader_lock

// Blacklisted : g_rw_lock_reader_trylock

// Blacklisted : g_rw_lock_reader_unlock

// Blacklisted : g_rw_lock_writer_lock

// Blacklisted : g_rw_lock_writer_trylock

// Blacklisted : g_rw_lock_writer_unlock

// RecMutex is a wrapper around the C record GRecMutex.
type RecMutex struct {
	native *C.GRecMutex
	// Private : p
	// Private : i
}

func RecMutexNewFromC(u unsafe.Pointer) *RecMutex {
	c := (*C.GRecMutex)(u)
	if c == nil {
		return nil
	}

	g := &RecMutex{native: c}

	return g
}

func (recv *RecMutex) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RecMutex with another RecMutex, and returns true if they represent the same GObject.
func (recv *RecMutex) Equals(other *RecMutex) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_rec_mutex_clear

// Init is a wrapper around the C function g_rec_mutex_init.
func (recv *RecMutex) Init() {
	C.g_rec_mutex_init((*C.GRecMutex)(recv.native))

	return
}

// Blacklisted : g_rec_mutex_lock

// Blacklisted : g_rec_mutex_trylock

// Blacklisted : g_rec_mutex_unlock

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Blacklisted : g_thread_ref

// Blacklisted : g_thread_unref

// Unsupported : g_variant_new_fixed_array : unsupported parameter elements : no type generator for gpointer (gconstpointer) for param elements

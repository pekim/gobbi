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

// BytesNew is a wrapper around the C function g_bytes_new.
func BytesNew(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data), len(data))
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BytesNewStatic is a wrapper around the C function g_bytes_new_static.
func BytesNewStatic(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data), len(data))
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_static(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BytesNewTake is a wrapper around the C function g_bytes_new_take.
func BytesNewTake(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data), len(data))
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_take(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_bytes_new_with_free_func : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// Compare is a wrapper around the C function g_bytes_compare.
func (recv *Bytes) Compare(bytes2 *Bytes) int32 {
	c_bytes2 := (C.gconstpointer)(C.NULL)
	if bytes2 != nil {
		c_bytes2 = (C.gconstpointer)(bytes2.ToC())
	}

	retC := C.g_bytes_compare((C.gconstpointer)(recv.native), c_bytes2)
	retGo := (int32)(retC)

	return retGo
}

// Equal is a wrapper around the C function g_bytes_equal.
func (recv *Bytes) Equal(bytes2 *Bytes) bool {
	c_bytes2 := (C.gconstpointer)(C.NULL)
	if bytes2 != nil {
		c_bytes2 = (C.gconstpointer)(bytes2.ToC())
	}

	retC := C.g_bytes_equal((C.gconstpointer)(recv.native), c_bytes2)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_bytes_get_data : array return type :

// GetSize is a wrapper around the C function g_bytes_get_size.
func (recv *Bytes) GetSize() uint64 {
	retC := C.g_bytes_get_size((*C.GBytes)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Hash is a wrapper around the C function g_bytes_hash.
func (recv *Bytes) Hash() uint32 {
	retC := C.g_bytes_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NewFromBytes is a wrapper around the C function g_bytes_new_from_bytes.
func (recv *Bytes) NewFromBytes(offset uint64, length uint64) *Bytes {
	c_offset := (C.gsize)(offset)

	c_length := (C.gsize)(length)

	retC := C.g_bytes_new_from_bytes((*C.GBytes)(recv.native), c_offset, c_length)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_bytes_ref.
func (recv *Bytes) Ref() *Bytes {
	retC := C.g_bytes_ref((*C.GBytes)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_bytes_unref.
func (recv *Bytes) Unref() {
	C.g_bytes_unref((*C.GBytes)(recv.native))

	return
}

// Unsupported : g_bytes_unref_to_array : array return type :

// Unsupported : g_bytes_unref_to_data : array return type :

// Clear is a wrapper around the C function g_cond_clear.
func (recv *Cond) Clear() {
	C.g_cond_clear((*C.GCond)(recv.native))

	return
}

// Init is a wrapper around the C function g_cond_init.
func (recv *Cond) Init() {
	C.g_cond_init((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// g_hash_table_add : unsupported parameter key : no type generator for gpointer (gpointer) for param key
// g_hash_table_contains : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key
// Ref is a wrapper around the C function g_key_file_ref.
func (recv *KeyFile) Ref() *KeyFile {
	retC := C.g_key_file_ref((*C.GKeyFile)(recv.native))
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_key_file_unref.
func (recv *KeyFile) Unref() {
	C.g_key_file_unref((*C.GKeyFile)(recv.native))

	return
}

// MainContextRefThreadDefault is a wrapper around the C function g_main_context_ref_thread_default.
func MainContextRefThreadDefault() *MainContext {
	retC := C.g_main_context_ref_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MappedFileNewFromFd is a wrapper around the C function g_mapped_file_new_from_fd.
func MappedFileNewFromFd(fd int32, writable bool) (*MappedFile, error) {
	c_fd := (C.gint)(fd)

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new_from_fd(c_fd, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

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

// Clear is a wrapper around the C function g_rw_lock_clear.
func (recv *RWLock) Clear() {
	C.g_rw_lock_clear((*C.GRWLock)(recv.native))

	return
}

// Init is a wrapper around the C function g_rw_lock_init.
func (recv *RWLock) Init() {
	C.g_rw_lock_init((*C.GRWLock)(recv.native))

	return
}

// ReaderLock is a wrapper around the C function g_rw_lock_reader_lock.
func (recv *RWLock) ReaderLock() {
	C.g_rw_lock_reader_lock((*C.GRWLock)(recv.native))

	return
}

// ReaderTrylock is a wrapper around the C function g_rw_lock_reader_trylock.
func (recv *RWLock) ReaderTrylock() bool {
	retC := C.g_rw_lock_reader_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ReaderUnlock is a wrapper around the C function g_rw_lock_reader_unlock.
func (recv *RWLock) ReaderUnlock() {
	C.g_rw_lock_reader_unlock((*C.GRWLock)(recv.native))

	return
}

// WriterLock is a wrapper around the C function g_rw_lock_writer_lock.
func (recv *RWLock) WriterLock() {
	C.g_rw_lock_writer_lock((*C.GRWLock)(recv.native))

	return
}

// WriterTrylock is a wrapper around the C function g_rw_lock_writer_trylock.
func (recv *RWLock) WriterTrylock() bool {
	retC := C.g_rw_lock_writer_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// WriterUnlock is a wrapper around the C function g_rw_lock_writer_unlock.
func (recv *RWLock) WriterUnlock() {
	C.g_rw_lock_writer_unlock((*C.GRWLock)(recv.native))

	return
}

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

// Clear is a wrapper around the C function g_rec_mutex_clear.
func (recv *RecMutex) Clear() {
	C.g_rec_mutex_clear((*C.GRecMutex)(recv.native))

	return
}

// Init is a wrapper around the C function g_rec_mutex_init.
func (recv *RecMutex) Init() {
	C.g_rec_mutex_init((*C.GRecMutex)(recv.native))

	return
}

// Lock is a wrapper around the C function g_rec_mutex_lock.
func (recv *RecMutex) Lock() {
	C.g_rec_mutex_lock((*C.GRecMutex)(recv.native))

	return
}

// Trylock is a wrapper around the C function g_rec_mutex_trylock.
func (recv *RecMutex) Trylock() bool {
	retC := C.g_rec_mutex_trylock((*C.GRecMutex)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unlock is a wrapper around the C function g_rec_mutex_unlock.
func (recv *RecMutex) Unlock() {
	C.g_rec_mutex_unlock((*C.GRecMutex)(recv.native))

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Ref is a wrapper around the C function g_thread_ref.
func (recv *Thread) Ref() *Thread {
	retC := C.g_thread_ref((*C.GThread)(recv.native))
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_thread_unref.
func (recv *Thread) Unref() {
	C.g_thread_unref((*C.GThread)(recv.native))

	return
}

// Unsupported : g_variant_new_fixed_array : unsupported parameter elements : no type generator for gpointer (gconstpointer) for param elements

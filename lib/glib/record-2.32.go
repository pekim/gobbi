// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

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

// Creates a new #GBytes from @data.
//
// @data is copied. If @size is 0, @data may be %NULL.
/*

C function : g_bytes_new
*/
func BytesNew(data []uint8) *Bytes {
	c_data := &data[0]

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new((C.gconstpointer)(unsafe.Pointer(c_data)), c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GBytes from static data.
//
// @data must be static (ie: never modified or freed). It may be %NULL if @size
// is 0.
/*

C function : g_bytes_new_static
*/
func BytesNewStatic(data []uint8) *Bytes {
	c_data := &data[0]

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_static((C.gconstpointer)(unsafe.Pointer(c_data)), c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GBytes from @data.
//
// After this call, @data belongs to the bytes and may no longer be
// modified by the caller.  g_free() will be called on @data when the
// bytes is no longer in use. Because of this @data must have been created by
// a call to g_malloc(), g_malloc0() or g_realloc() or by one of the many
// functions that wrap these calls (such as g_new(), g_strdup(), etc).
//
// For creating #GBytes with memory from other allocators, see
// g_bytes_new_with_free_func().
//
// @data may be %NULL if @size is 0.
/*

C function : g_bytes_new_take
*/
func BytesNewTake(data []uint8) *Bytes {
	c_data := &data[0]

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_take((C.gpointer)(unsafe.Pointer(c_data)), c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_bytes_new_with_free_func : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// Compares the two #GBytes values.
//
// This function can be used to sort GBytes instances in lexographical order.
/*

C function : g_bytes_compare
*/
func (recv *Bytes) Compare(bytes2 uintptr) int32 {
	c_bytes2 := (C.gconstpointer)(bytes2)

	retC := C.g_bytes_compare((C.gconstpointer)(recv.native), c_bytes2)
	retGo := (int32)(retC)

	return retGo
}

// Compares the two #GBytes values being pointed to and returns
// %TRUE if they are equal.
//
// This function can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-%NULL #GBytes pointers as keys in a #GHashTable.
/*

C function : g_bytes_equal
*/
func (recv *Bytes) Equal(bytes2 uintptr) bool {
	c_bytes2 := (C.gconstpointer)(bytes2)

	retC := C.g_bytes_equal((C.gconstpointer)(recv.native), c_bytes2)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_bytes_get_data : no return type

// Get the size of the byte data in the #GBytes.
//
// This function will always return the same value for a given #GBytes.
/*

C function : g_bytes_get_size
*/
func (recv *Bytes) GetSize() uint64 {
	retC := C.g_bytes_get_size((*C.GBytes)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Creates an integer hash code for the byte data in the #GBytes.
//
// This function can be passed to g_hash_table_new() as the @key_hash_func
// parameter, when using non-%NULL #GBytes pointers as keys in a #GHashTable.
/*

C function : g_bytes_hash
*/
func (recv *Bytes) Hash() uint32 {
	retC := C.g_bytes_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Creates a #GBytes which is a subsection of another #GBytes. The @offset +
// @length may not be longer than the size of @bytes.
//
// A reference to @bytes will be held by the newly created #GBytes until
// the byte data is no longer needed.
//
// Since 2.56, if @offset is 0 and @length matches the size of @bytes, then
// @bytes will be returned with the reference count incremented by 1. If @bytes
// is a slice of another #GBytes, then the resulting #GBytes will reference
// the same #GBytes instead of @bytes. This allows consumers to simplify the
// usage of #GBytes when asynchronously writing to streams.
/*

C function : g_bytes_new_from_bytes
*/
func (recv *Bytes) NewFromBytes(offset uint64, length uint64) *Bytes {
	c_offset := (C.gsize)(offset)

	c_length := (C.gsize)(length)

	retC := C.g_bytes_new_from_bytes((*C.GBytes)(recv.native), c_offset, c_length)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Increase the reference count on @bytes.
/*

C function : g_bytes_ref
*/
func (recv *Bytes) Ref() *Bytes {
	retC := C.g_bytes_ref((*C.GBytes)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Releases a reference on @bytes.  This may result in the bytes being
// freed. If @bytes is %NULL, it will return immediately.
/*

C function : g_bytes_unref
*/
func (recv *Bytes) Unref() {
	C.g_bytes_unref((*C.GBytes)(recv.native))

	return
}

// Unsupported : g_bytes_unref_to_array : no return type

// Unsupported : g_bytes_unref_to_data : no return type

// Frees the resources allocated to a #GCond with g_cond_init().
//
// This function should not be used with a #GCond that has been
// statically allocated.
//
// Calling g_cond_clear() for a #GCond on which threads are
// blocking leads to undefined behaviour.
/*

C function : g_cond_clear
*/
func (recv *Cond) Clear() {
	C.g_cond_clear((*C.GCond)(recv.native))

	return
}

// Initialises a #GCond so that it can be used.
//
// This function is useful to initialise a #GCond that has been
// allocated as part of a larger structure.  It is not necessary to
// initialise a #GCond that has been statically allocated.
//
// To undo the effect of g_cond_init() when a #GCond is no longer
// needed, use g_cond_clear().
//
// Calling g_cond_init() on an already-initialised #GCond leads
// to undefined behaviour.
/*

C function : g_cond_init
*/
func (recv *Cond) Init() {
	C.g_cond_init((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// Increases the reference count of @key_file.
/*

C function : g_key_file_ref
*/
func (recv *KeyFile) Ref() *KeyFile {
	retC := C.g_key_file_ref((*C.GKeyFile)(recv.native))
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decreases the reference count of @key_file by 1. If the reference count
// reaches zero, frees the key file and all its allocated memory.
/*

C function : g_key_file_unref
*/
func (recv *KeyFile) Unref() {
	C.g_key_file_unref((*C.GKeyFile)(recv.native))

	return
}

// Maps a file into memory. On UNIX, this is using the mmap() function.
//
// If @writable is %TRUE, the mapped buffer may be modified, otherwise
// it is an error to modify the mapped buffer. Modifications to the buffer
// are not visible to other processes mapping the same file, and are not
// written back to the file.
//
// Note that modifications of the underlying file might affect the contents
// of the #GMappedFile. Therefore, mapping should only be used if the file
// will not be modified, or if all modifications of the file are done
// atomically (e.g. using g_file_set_contents()).
/*

C function : g_mapped_file_new_from_fd
*/
func MappedFileNewFromFd(fd int32, writable bool) (*MappedFile, error) {
	c_fd := (C.gint)(fd)

	c_writable :=
		boolToGboolean(writable)

	var cThrowableError *C.GError

	retC := C.g_mapped_file_new_from_fd(c_fd, c_writable, &cThrowableError)
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the thread local variable @key to have the value @value in the
// current thread.
//
// This function differs from g_private_set() in the following way: if
// the previous value was non-%NULL then the #GDestroyNotify handler for
// @key is run on it.
/*

C function : g_private_replace
*/
func (recv *Private) Replace(value uintptr) {
	c_value := (C.gpointer)(value)

	C.g_private_replace((*C.GPrivate)(recv.native), c_value)

	return
}

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

// Frees the resources allocated to a lock with g_rw_lock_init().
//
// This function should not be used with a #GRWLock that has been
// statically allocated.
//
// Calling g_rw_lock_clear() when any thread holds the lock
// leads to undefined behaviour.
//
// Sine: 2.32
/*

C function : g_rw_lock_clear
*/
func (recv *RWLock) Clear() {
	C.g_rw_lock_clear((*C.GRWLock)(recv.native))

	return
}

// Initializes a #GRWLock so that it can be used.
//
// This function is useful to initialize a lock that has been
// allocated on the stack, or as part of a larger structure.  It is not
// necessary to initialise a reader-writer lock that has been statically
// allocated.
//
// |[<!-- language="C" -->
// typedef struct {
// GRWLock l;
// ...
// } Blob;
//
// Blob *b;
//
// b = g_new (Blob, 1);
// g_rw_lock_init (&b->l);
// ]|
//
// To undo the effect of g_rw_lock_init() when a lock is no longer
// needed, use g_rw_lock_clear().
//
// Calling g_rw_lock_init() on an already initialized #GRWLock leads
// to undefined behaviour.
/*

C function : g_rw_lock_init
*/
func (recv *RWLock) Init() {
	C.g_rw_lock_init((*C.GRWLock)(recv.native))

	return
}

// Obtain a read lock on @rw_lock. If another thread currently holds
// the write lock on @rw_lock or blocks waiting for it, the current
// thread will block. Read locks can be taken recursively.
//
// It is implementation-defined how many threads are allowed to
// hold read locks on the same lock simultaneously. If the limit is hit,
// or if a deadlock is detected, a critical warning will be emitted.
/*

C function : g_rw_lock_reader_lock
*/
func (recv *RWLock) ReaderLock() {
	C.g_rw_lock_reader_lock((*C.GRWLock)(recv.native))

	return
}

// Tries to obtain a read lock on @rw_lock and returns %TRUE if
// the read lock was successfully obtained. Otherwise it
// returns %FALSE.
/*

C function : g_rw_lock_reader_trylock
*/
func (recv *RWLock) ReaderTrylock() bool {
	retC := C.g_rw_lock_reader_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Release a read lock on @rw_lock.
//
// Calling g_rw_lock_reader_unlock() on a lock that is not held
// by the current thread leads to undefined behaviour.
/*

C function : g_rw_lock_reader_unlock
*/
func (recv *RWLock) ReaderUnlock() {
	C.g_rw_lock_reader_unlock((*C.GRWLock)(recv.native))

	return
}

// Obtain a write lock on @rw_lock. If any thread already holds
// a read or write lock on @rw_lock, the current thread will block
// until all other threads have dropped their locks on @rw_lock.
/*

C function : g_rw_lock_writer_lock
*/
func (recv *RWLock) WriterLock() {
	C.g_rw_lock_writer_lock((*C.GRWLock)(recv.native))

	return
}

// Tries to obtain a write lock on @rw_lock. If any other thread holds
// a read or write lock on @rw_lock, it immediately returns %FALSE.
// Otherwise it locks @rw_lock and returns %TRUE.
/*

C function : g_rw_lock_writer_trylock
*/
func (recv *RWLock) WriterTrylock() bool {
	retC := C.g_rw_lock_writer_trylock((*C.GRWLock)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Release a write lock on @rw_lock.
//
// Calling g_rw_lock_writer_unlock() on a lock that is not held
// by the current thread leads to undefined behaviour.
/*

C function : g_rw_lock_writer_unlock
*/
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

// Frees the resources allocated to a recursive mutex with
// g_rec_mutex_init().
//
// This function should not be used with a #GRecMutex that has been
// statically allocated.
//
// Calling g_rec_mutex_clear() on a locked recursive mutex leads
// to undefined behaviour.
//
// Sine: 2.32
/*

C function : g_rec_mutex_clear
*/
func (recv *RecMutex) Clear() {
	C.g_rec_mutex_clear((*C.GRecMutex)(recv.native))

	return
}

// Initializes a #GRecMutex so that it can be used.
//
// This function is useful to initialize a recursive mutex
// that has been allocated on the stack, or as part of a larger
// structure.
//
// It is not necessary to initialise a recursive mutex that has been
// statically allocated.
//
// |[<!-- language="C" -->
// typedef struct {
// GRecMutex m;
// ...
// } Blob;
//
// Blob *b;
//
// b = g_new (Blob, 1);
// g_rec_mutex_init (&b->m);
// ]|
//
// Calling g_rec_mutex_init() on an already initialized #GRecMutex
// leads to undefined behaviour.
//
// To undo the effect of g_rec_mutex_init() when a recursive mutex
// is no longer needed, use g_rec_mutex_clear().
/*

C function : g_rec_mutex_init
*/
func (recv *RecMutex) Init() {
	C.g_rec_mutex_init((*C.GRecMutex)(recv.native))

	return
}

// Locks @rec_mutex. If @rec_mutex is already locked by another
// thread, the current thread will block until @rec_mutex is
// unlocked by the other thread. If @rec_mutex is already locked
// by the current thread, the 'lock count' of @rec_mutex is increased.
// The mutex will only become available again when it is unlocked
// as many times as it has been locked.
/*

C function : g_rec_mutex_lock
*/
func (recv *RecMutex) Lock() {
	C.g_rec_mutex_lock((*C.GRecMutex)(recv.native))

	return
}

// Tries to lock @rec_mutex. If @rec_mutex is already locked
// by another thread, it immediately returns %FALSE. Otherwise
// it locks @rec_mutex and returns %TRUE.
/*

C function : g_rec_mutex_trylock
*/
func (recv *RecMutex) Trylock() bool {
	retC := C.g_rec_mutex_trylock((*C.GRecMutex)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unlocks @rec_mutex. If another thread is blocked in a
// g_rec_mutex_lock() call for @rec_mutex, it will become unblocked
// and can lock @rec_mutex itself.
//
// Calling g_rec_mutex_unlock() on a recursive mutex that is not
// locked by the current thread leads to undefined behaviour.
/*

C function : g_rec_mutex_unlock
*/
func (recv *RecMutex) Unlock() {
	C.g_rec_mutex_unlock((*C.GRecMutex)(recv.native))

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Increase the reference count on @thread.
/*

C function : g_thread_ref
*/
func (recv *Thread) Ref() *Thread {
	retC := C.g_thread_ref((*C.GThread)(recv.native))
	retGo := ThreadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrease the reference count on @thread, possibly freeing all
// resources associated with it.
//
// Note that each thread holds a reference to its #GThread while
// it is running, so it is safe to drop your own reference to it
// if you don't need it anymore.
/*

C function : g_thread_unref
*/
func (recv *Thread) Unref() {
	C.g_thread_unref((*C.GThread)(recv.native))

	return
}

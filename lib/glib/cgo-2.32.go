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

// Unsupported : g_byte_array_new_take : array return type :

// Unsupported : g_hash_table_add : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : g_hash_table_contains : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

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

// BytesNew is a wrapper around the C function g_bytes_new.
func BytesNew(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BytesNewStatic is a wrapper around the C function g_bytes_new_static.
func BytesNewStatic(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_static(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BytesNewTake is a wrapper around the C function g_bytes_new_take.
func BytesNewTake(data []uint8) *Bytes {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_size := (C.gsize)(len(data))

	retC := C.g_bytes_new_take(c_data, c_size)
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_bytes_new_with_free_func : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

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

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_variant_new_fixed_array : unsupported parameter elements : no type generator for gpointer (gconstpointer) for param elements

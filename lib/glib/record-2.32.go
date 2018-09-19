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

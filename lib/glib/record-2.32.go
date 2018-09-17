// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "runtime"

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

func bytesNewFromC(c *C.GBytes, finalizeFree bool) *Bytes {
	if c == nil {
		return nil
	}

	g := &Bytes{native: c}

	if finalizeFree {
		runtime.SetFinalizer(g, func(obj interface{}) {
			C.g_free((C.gpointer)(c))
		})
	}

	return g
}

// Rwlock is a wrapper around the C record GRWLock.
type Rwlock struct {
	native *C.GRWLock
	P      uintptr
	// no type for i
}

func rwlockNewFromC(c *C.GRWLock, finalizeFree bool) *Rwlock {
	if c == nil {
		return nil
	}

	g := &Rwlock{
		P:      (uintptr)(c.p),
		native: c,
	}

	if finalizeFree {
		runtime.SetFinalizer(g, func(obj interface{}) {
			C.g_free((C.gpointer)(c))
		})
	}

	return g
}

// Recmutex is a wrapper around the C record GRecMutex.
type Recmutex struct {
	native *C.GRecMutex
	P      uintptr
	// no type for i
}

func recmutexNewFromC(c *C.GRecMutex, finalizeFree bool) *Recmutex {
	if c == nil {
		return nil
	}

	g := &Recmutex{
		P:      (uintptr)(c.p),
		native: c,
	}

	if finalizeFree {
		runtime.SetFinalizer(g, func(obj interface{}) {
			C.g_free((C.gpointer)(c))
		})
	}

	return g
}

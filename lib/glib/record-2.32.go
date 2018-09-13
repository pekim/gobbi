// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Bytes is a wrapper around the C function GBytes.
type Bytes struct{}

// Rwlock is a wrapper around the C function GRWLock.
type Rwlock struct{}

// Recmutex is a wrapper around the C function GRecMutex.
type Recmutex struct{}

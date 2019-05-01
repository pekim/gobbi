// This is a generated file - DO NOT EDIT
// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Hmac is a wrapper around the C record GHmac.
type Hmac struct {
	native *C.GHmac
}

func HmacNewFromC(u unsafe.Pointer) *Hmac {
	c := (*C.GHmac)(u)
	if c == nil {
		return nil
	}

	g := &Hmac{native: c}

	return g
}

func (recv *Hmac) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Blacklisted : g_atomic_int_add

// Unsupported : g_atomic_int_compare_and_exchange : return type :

// Unsupported : g_atomic_int_dec_and_test : return type :

// Blacklisted : g_atomic_int_exchange_and_add

// Blacklisted : g_atomic_int_get

// Blacklisted : g_atomic_int_inc

// Blacklisted : g_atomic_int_set

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_source_new : return type :

// Unsupported : g_file_read_link : return type :

// Unsupported : g_markup_printf_escaped : return type :

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_setenv : return type :

// Unsupported : g_strip_context : return type :

// Unsupported : g_strsplit_set : array return type :

// Unsupported : g_unichar_get_mirror_char : return type :

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2

// Once is a wrapper around the C record GOnce.
type Once struct {
	native *C.GOnce
	Status OnceStatus
	// retval : no type generator for gpointer, volatile gpointer
}

func OnceNewFromC(u unsafe.Pointer) *Once {
	c := (*C.GOnce)(u)
	if c == nil {
		return nil
	}

	g := &Once{
		Status: (OnceStatus)(c.status),
		native: c,
	}

	return g
}

func (recv *Once) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

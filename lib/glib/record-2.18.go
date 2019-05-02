// This is a generated file - DO NOT EDIT
// +build glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Reset is a wrapper around the C function g_checksum_reset.
func (recv *Checksum) Reset() {
	C.g_checksum_reset((*C.GChecksum)(recv.native))

	return
}

// Unsupported : g_markup_parse_context_get_user_data : no return generator

// Unsupported : g_markup_parse_context_pop : no return generator

// Unsupported : g_markup_parse_context_push : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

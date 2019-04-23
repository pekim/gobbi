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

// Blacklisted : g_dir_make_tmp

// Unsupported : g_hash_table_iter_replace : unsupported parameter value : no type generator for gpointer (gpointer) for param value

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

// Equals compares this Hmac with another Hmac, and returns true if they represent the same GObject.
func (recv *Hmac) Equals(other *Hmac) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_hmac_new

// Blacklisted : g_hmac_copy

// Blacklisted : g_hmac_get_digest

// Blacklisted : g_hmac_get_string

// Blacklisted : g_hmac_ref

// Blacklisted : g_hmac_unref

// Blacklisted : g_hmac_update

// Blacklisted : g_match_info_ref

// Blacklisted : g_match_info_unref

// Blacklisted : g_variant_new_objv

// Blacklisted : g_variant_dup_objv

// Blacklisted : g_variant_get_objv

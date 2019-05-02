// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// VariantNewFromBytes is a wrapper around the C function g_variant_new_from_bytes.
func VariantNewFromBytes(type_ *VariantType, bytes *Bytes, trusted bool) *Variant {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	c_trusted :=
		boolToGboolean(trusted)

	retC := C.g_variant_new_from_bytes(c_type, c_bytes, c_trusted)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

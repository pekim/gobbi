// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Checksum is a wrapper around the C record GChecksum.
type Checksum struct {
	native *C.GChecksum
}

func checksumNewFromC(c *C.GChecksum) *Checksum {
	if c == nil {
		return nil
	}

	g := &Checksum{native: c}

	return g
}

func (recv *Checksum) toC() *C.GChecksum {

	return recv.native
}

// ChecksumNew is a wrapper around the C function g_checksum_new.
func ChecksumNew(checksumType ChecksumType) *Checksum {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_new(c_checksum_type)
	retGo := checksumNewFromC(retC)

	return retGo
}

// Copy is a wrapper around the C function g_checksum_copy.
func (recv *Checksum) Copy() *Checksum {
	retC := C.g_checksum_copy((*C.GChecksum)(recv.native))
	retGo := checksumNewFromC(retC)

	return retGo
}

// Unsupported : g_checksum_free : no return generator

// Unsupported : g_checksum_get_digest : unsupported parameter buffer : no type generator for guint8, guint8*

// GetString is a wrapper around the C function g_checksum_get_string.
func (recv *Checksum) GetString() string {
	retC := C.g_checksum_get_string((*C.GChecksum)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_checksum_reset : no return generator

// Unsupported : g_checksum_update : unsupported parameter data : no param type

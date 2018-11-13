// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// The hashing algorithm to be used by #GChecksum when performing the
// digest of some data.
//
// Note that the #GChecksumType enumeration may be extended at a later
// date to include new hashing algorithm types.
type ChecksumType C.GChecksumType

const (
	// Use the MD5 hashing algorithm
	CHECKSUM_MD5 ChecksumType = 0
	// Use the SHA-1 hashing algorithm
	CHECKSUM_SHA1 ChecksumType = 1
	// Use the SHA-256 hashing algorithm
	CHECKSUM_SHA256 ChecksumType = 2
	// Use the SHA-512 hashing algorithm (Since: 2.36)
	CHECKSUM_SHA512 ChecksumType = 3
	// Use the SHA-384 hashing algorithm (Since: 2.51)
	CHECKSUM_SHA384 ChecksumType = 4
)

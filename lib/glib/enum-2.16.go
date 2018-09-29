// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

type ChecksumType C.GChecksumType

const (
	CHECKSUM_MD5    ChecksumType = 0
	CHECKSUM_SHA1   ChecksumType = 1
	CHECKSUM_SHA256 ChecksumType = 2
	CHECKSUM_SHA512 ChecksumType = 3
	CHECKSUM_SHA384 ChecksumType = 4
)

// Blacklisted : GTestResult

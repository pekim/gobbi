// This is a generated file - DO NOT EDIT
// +build glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// IsEmpty is a wrapper around the C function g_sequence_is_empty.
func (recv *Sequence) IsEmpty() bool {
	retC := C.g_sequence_is_empty((*C.GSequence)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// This is a generated file - DO NOT EDIT
// +build glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Returns %TRUE if the sequence contains zero items.
//
// This function is functionally identical to checking the result of
// g_sequence_get_length() being equal to zero. However this function is
// implemented in O(1) running time.
/*

C function : g_sequence_is_empty
*/
func (recv *Sequence) IsEmpty() bool {
	retC := C.g_sequence_is_empty((*C.GSequence)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

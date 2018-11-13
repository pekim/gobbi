// This is a generated file - DO NOT EDIT
// +build glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Creates a new #GBytes which references the data mapped from @file.
// The mapped contents of the file must not be modified after creating this
// bytes object, because a #GBytes should be immutable.
/*

C function

g_mapped_file_get_bytes
*/
func (recv *MappedFile) GetBytes() *Bytes {
	retC := C.g_mapped_file_get_bytes((*C.GMappedFile)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks whether the pattern contains explicit CR or LF references.
/*

C function

g_regex_get_has_cr_or_lf
*/
func (recv *Regex) GetHasCrOrLf() bool {
	retC := C.g_regex_get_has_cr_or_lf((*C.GRegex)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Transfers ownership of the contents of @string to a newly allocated
// #GBytes.  The #GString structure itself is deallocated, and it is
// therefore invalid to use @string after invoking this function.
//
// Note that while #GString ensures that its buffer always has a
// trailing nul character (not reflected in its "len"), the returned
// #GBytes does not include this extra nul; i.e. it has length exactly
// equal to the "len" member.
/*

C function

g_string_free_to_bytes
*/
func (recv *String) FreeToBytes() *Bytes {
	retC := C.g_string_free_to_bytes((*C.GString)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

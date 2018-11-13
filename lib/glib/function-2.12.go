// This is a generated file - DO NOT EDIT
// +build glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Converts a string to a #gint64 value.
// This function behaves like the standard strtoll() function
// does in the C locale. It does this without actually
// changing the current locale, since that would not be
// thread-safe.
//
// This function is typically used when reading configuration
// files or other non-user input that should be locale independent.
// To handle input from the user you should normally use the
// locale-sensitive system strtoll() function.
//
// If the correct value would cause overflow, %G_MAXINT64 or %G_MININT64
// is returned, and `ERANGE` is stored in `errno`.
// If the base is outside the valid range, zero is returned, and
// `EINVAL` is stored in `errno`. If the
// string conversion fails, zero is returned, and @endptr returns @nptr
// (if @endptr is non-%NULL).
/*

C function

g_ascii_strtoll
*/
func AsciiStrtoll(nptr string, base uint32) (int64, string) {
	c_nptr := C.CString(nptr)
	defer C.free(unsafe.Pointer(c_nptr))

	var c_endptr *C.gchar

	c_base := (C.guint)(base)

	retC := C.g_ascii_strtoll(c_nptr, &c_endptr, c_base)
	retGo := (int64)(retC)

	endptr := C.GoString(c_endptr)

	return retGo, endptr
}

// Unsupported : g_base64_decode : no return type

// Unsupported : g_base64_decode_step : unsupported parameter out : output array param out

// Encode a sequence of binary data into its Base-64 stringified
// representation.
/*

C function

g_base64_encode
*/
func Base64Encode(data []uint8) string {
	c_data := &data[0]

	c_len := (C.gsize)(len(data))

	retC := C.g_base64_encode((*C.guchar)(unsafe.Pointer(c_data)), c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_base64_encode_close : unsupported parameter out : output array param out

// Unsupported : g_base64_encode_step : unsupported parameter out : output array param out

// Removes all keys and their associated values from a #GHashTable.
//
// If the #GHashTable was created using g_hash_table_new_full(),
// the keys and values are freed using the supplied destroy functions,
// otherwise you have to make sure that any dynamically allocated
// values are freed yourself.
/*

C function

g_hash_table_remove_all
*/
func HashTableRemoveAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_remove_all(c_hash_table)

	return
}

// Removes all keys and their associated values from a #GHashTable
// without calling the key and value destroy functions.
/*

C function

g_hash_table_steal_all
*/
func HashTableStealAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_steal_all(c_hash_table)

	return
}

// Returns the currently firing source for this thread.
/*

C function

g_main_current_source
*/
func MainCurrentSource() *Source {
	retC := C.g_main_current_source()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts a string containing an ISO 8601 encoded date and time
// to a #GTimeVal and puts it into @time_.
//
// @iso_date must include year, month, day, hours, minutes, and
// seconds. It can optionally include fractions of a second and a time
// zone indicator. (In the absence of any time zone indication, the
// timestamp is assumed to be in local time.)
/*

C function

g_time_val_from_iso8601
*/
func TimeValFromIso8601(isoDate string) (bool, *TimeVal) {
	c_iso_date := C.CString(isoDate)
	defer C.free(unsafe.Pointer(c_iso_date))

	var c_time_ C.GTimeVal

	retC := C.g_time_val_from_iso8601(c_iso_date, &c_time_)
	retGo := retC == C.TRUE

	time := TimeValNewFromC(unsafe.Pointer(&c_time_))

	return retGo, time
}

// Determines if a character is typically rendered in a double-width
// cell under legacy East Asian locales.  If a character is wide according to
// g_unichar_iswide(), then it is also reported wide with this function, but
// the converse is not necessarily true. See the
// [Unicode Standard Annex #11](http://www.unicode.org/reports/tr11/)
// for details.
//
// If a character passes the g_unichar_iswide() test then it will also pass
// this test, but not the other way around.  Note that some characters may
// pass both this test and g_unichar_iszerowidth().
/*

C function

g_unichar_iswide_cjk
*/
func UnicharIswideCjk(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iswide_cjk(c_c)
	retGo := retC == C.TRUE

	return retGo
}

// This is a generated file - DO NOT EDIT
// +build glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// AsciiStrtoll is a wrapper around the C function g_ascii_strtoll.
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

// Unsupported : g_base64_decode : unsupported parameter out_len : no type generator for gsize, gsize*

// Unsupported : g_base64_decode_step : unsupported parameter in : no param type

// Unsupported : g_base64_encode : unsupported parameter data : no param type

// Unsupported : g_base64_encode_close : unsupported parameter out : no param type

// Unsupported : g_base64_encode_step : unsupported parameter in : no param type

// HashTableRemoveAll is a wrapper around the C function g_hash_table_remove_all.
func HashTableRemoveAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	C.g_hash_table_remove_all(c_hash_table)

	return
}

// HashTableStealAll is a wrapper around the C function g_hash_table_steal_all.
func HashTableStealAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(hashTable.ToC())

	C.g_hash_table_steal_all(c_hash_table)

	return
}

// MainCurrentSource is a wrapper around the C function g_main_current_source.
func MainCurrentSource() *Source {
	retC := C.g_main_current_source()
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TimeValFromIso8601 is a wrapper around the C function g_time_val_from_iso8601.
func TimeValFromIso8601(isoDate string) (bool, *TimeVal) {
	c_iso_date := C.CString(isoDate)
	defer C.free(unsafe.Pointer(c_iso_date))

	var c_time_ C.GTimeVal

	retC := C.g_time_val_from_iso8601(c_iso_date, &c_time_)
	retGo := retC == C.TRUE

	time := TimeValNewFromC(unsafe.Pointer(&c_time_))

	return retGo, time
}

// UnicharIswideCjk is a wrapper around the C function g_unichar_iswide_cjk.
func UnicharIswideCjk(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iswide_cjk(c_c)
	retGo := retC == C.TRUE

	return retGo
}

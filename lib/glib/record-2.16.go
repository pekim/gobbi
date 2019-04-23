// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// g_async_queue_new_full : unsupported parameter item_free_func : no type generator for DestroyNotify (GDestroyNotify) for param item_free_func
// Checksum is a wrapper around the C record GChecksum.
type Checksum struct {
	native *C.GChecksum
}

func ChecksumNewFromC(u unsafe.Pointer) *Checksum {
	c := (*C.GChecksum)(u)
	if c == nil {
		return nil
	}

	g := &Checksum{native: c}

	return g
}

func (recv *Checksum) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Checksum with another Checksum, and returns true if they represent the same GObject.
func (recv *Checksum) Equals(other *Checksum) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_checksum_new

// Blacklisted : g_checksum_type_get_length

// Blacklisted : g_checksum_copy

// Blacklisted : g_checksum_free

// Blacklisted : g_checksum_get_digest

// Blacklisted : g_checksum_get_string

// Blacklisted : g_checksum_update

// Blacklisted : g_hash_table_iter_get_hash_table

// Init is a wrapper around the C function g_hash_table_iter_init.
func (recv *HashTableIter) Init(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_iter_init((*C.GHashTableIter)(recv.native), c_hash_table)

	return
}

// Unsupported : g_hash_table_iter_next : unsupported parameter key : no type generator for gpointer (gpointer*) for param key

// Blacklisted : g_hash_table_iter_remove

// Blacklisted : g_hash_table_iter_steal

// Blacklisted : g_markup_parse_context_get_element_stack

// Blacklisted : g_string_append_uri_escaped

// Blacklisted : g_test_suite_add

// Blacklisted : g_test_suite_add_suite

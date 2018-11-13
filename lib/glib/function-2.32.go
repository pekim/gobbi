// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Transfers the data from the #GByteArray into a new immutable #GBytes.
//
// The #GByteArray is freed unless the reference count of @array is greater
// than one, the #GByteArray wrapper is preserved but the size of @array
// will be set to zero.
//
// This is identical to using g_bytes_new_take() and g_byte_array_free()
// together.
/*

C function

g_byte_array_free_to_bytes
*/
func ByteArrayFreeToBytes(array []uint8) *Bytes {
	c_array := &array[0]

	retC := C.g_byte_array_free_to_bytes((*C.GByteArray)(unsafe.Pointer(c_array)))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_byte_array_new_take : no return type

// Unsupported : g_environ_getenv : unsupported parameter envp :

// Unsupported : g_environ_setenv : unsupported parameter envp :

// Unsupported : g_environ_unsetenv : unsupported parameter envp :

// This is a convenience function for using a #GHashTable as a set.  It
// is equivalent to calling g_hash_table_replace() with @key as both the
// key and the value.
//
// When a hash table only ever contains keys that have themselves as the
// corresponding value it is able to be stored more efficiently.  See
// the discussion in the section description.
//
// Starting from GLib 2.40, this function returns a boolean value to
// indicate whether the newly added value was already in the hash table
// or not.
/*

C function

g_hash_table_add
*/
func HashTableAdd(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gpointer)(key)

	retC := C.g_hash_table_add(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @key is in @hash_table.
/*

C function

g_hash_table_contains
*/
func HashTableContains(hashTable *HashTable, key uintptr) bool {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	c_key := (C.gconstpointer)(key)

	retC := C.g_hash_table_contains(c_hash_table, c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the thread-default #GMainContext for this thread, as with
// g_main_context_get_thread_default(), but also adds a reference to
// it with g_main_context_ref(). In addition, unlike
// g_main_context_get_thread_default(), if the thread-default context
// is the global default context, this will return that #GMainContext
// (with a ref added to it) rather than returning %NULL.
/*

C function

g_main_context_ref_thread_default
*/
func MainContextRefThreadDefault() *MainContext {
	retC := C.g_main_context_ref_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

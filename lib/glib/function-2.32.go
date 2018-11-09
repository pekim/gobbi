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

// ByteArrayFreeToBytes is a wrapper around the C function g_byte_array_free_to_bytes.
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

// HashTableAdd is a wrapper around the C function g_hash_table_add.
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

// HashTableContains is a wrapper around the C function g_hash_table_contains.
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

// MainContextRefThreadDefault is a wrapper around the C function g_main_context_ref_thread_default.
func MainContextRefThreadDefault() *MainContext {
	retC := C.g_main_context_ref_thread_default()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

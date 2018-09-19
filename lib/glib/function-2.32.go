// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_byte_array_free_to_bytes : unsupported parameter array : no param type

// Unsupported : g_byte_array_new_take : unsupported parameter data : no param type

// Unsupported : g_environ_getenv : unsupported parameter envp : no param type

// Unsupported : g_environ_setenv : unsupported parameter envp : no param type

// Unsupported : g_environ_unsetenv : unsupported parameter envp : no param type

// Unsupported : g_hash_table_add : unsupported parameter hash_table : no type generator for GLib.HashTable, GHashTable*

// Unsupported : g_hash_table_contains : unsupported parameter hash_table : no type generator for GLib.HashTable, GHashTable*

// MainContextRefThreadDefault is a wrapper around the C function g_main_context_ref_thread_default.
func MainContextRefThreadDefault() *MainContext {
	retC := C.g_main_context_ref_thread_default()
	retGo := mainContextNewFromC(retC)

	return retGo
}

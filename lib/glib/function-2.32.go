// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_byte_array_free_to_bytes : unsupported parameter array : no param type

// Unsupported : g_byte_array_new_take : unsupported parameter data : no param type

// Unsupported : g_environ_getenv : unsupported parameter envp : no param type

// Unsupported : g_environ_setenv : unsupported parameter envp : no param type

// Unsupported : g_environ_unsetenv : unsupported parameter envp : no param type

// Unsupported : g_hash_table_add : unsupported parameter hash_table : no param type

// Unsupported : g_hash_table_contains : unsupported parameter hash_table : no param type

// MainContextRefThreadDefault is a wrapper around the C function g_main_context_ref_thread_default.
func MainContextRefThreadDefault() {
	C.g_main_context_ref_thread_default()
}

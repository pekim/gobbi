// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_atomic_int_add : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_compare_and_exchange : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_dec_and_test : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_exchange_and_add : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_get : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_inc : unsupported parameter atomic : no param type

// Unsupported : g_atomic_int_set : unsupported parameter atomic : no param type

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no param type

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no param type

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no param type

// Unsupported : g_child_watch_add : unsupported parameter pid : no param type

// Unsupported : g_child_watch_add_full : unsupported parameter pid : no param type

// Unsupported : g_child_watch_source_new : unsupported parameter pid : no param type

// Unsupported : g_file_read_link : unsupported parameter filename : no param type

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no param type

// Unsupported : g_setenv : unsupported parameter variable : no param type

// StripContext is a wrapper around the C function g_strip_context.
func StripContext(msgid string, msgval string) {
	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgval := C.CString(msgval)
	defer C.free(unsafe.Pointer(c_msgval))

	C.g_strip_context(c_msgid, c_msgval)
}

// StrsplitSet is a wrapper around the C function g_strsplit_set.
func StrsplitSet(string string, delimiters string, maxTokens int32) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiters := C.CString(delimiters)
	defer C.free(unsafe.Pointer(c_delimiters))

	c_max_tokens := (C.gint)(maxTokens)

	C.g_strsplit_set(c_string, c_delimiters, c_max_tokens)
}

// Unsupported : g_unichar_get_mirror_char : unsupported parameter mirrored_ch : no param type

// Unsupported : g_unsetenv : unsupported parameter variable : no param type

// Unsupported : g_vasprintf : unsupported parameter string : in for string with indirection level of 2
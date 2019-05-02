// This is a generated file - DO NOT EDIT
// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Unsupported : g_strv_length : unsupported parameter str_array : in string with indirection level of 2

// KeyFileNew is a wrapper around the C function g_key_file_new.
func KeyFileNew() *KeyFile {
	retC := C.g_key_file_new()
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_option_group_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

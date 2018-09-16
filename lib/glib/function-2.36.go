// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_close : no return generator

// GetNumProcessors is a wrapper around the C function g_get_num_processors.
func GetNumProcessors() uint32 {
	retC := C.g_get_num_processors()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_unix_fd_add : unsupported parameter condition : no type generator for IOCondition, GIOCondition

// Unsupported : g_unix_fd_add_full : unsupported parameter condition : no type generator for IOCondition, GIOCondition

// Unsupported : g_unix_fd_source_new : unsupported parameter condition : no type generator for IOCondition, GIOCondition

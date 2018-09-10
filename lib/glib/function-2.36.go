// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Close is a wrapper around the C function g_close.
func Close(fd int32) {}

// GetNumProcessors is a wrapper around the C function g_get_num_processors.
func GetNumProcessors() {}

// Unsupported : g_unix_fd_add : unsupported parameter condition : type IOCondition, GIOCondition

// Unsupported : g_unix_fd_add_full : unsupported parameter condition : type IOCondition, GIOCondition

// Unsupported : g_unix_fd_source_new : unsupported parameter condition : type IOCondition, GIOCondition

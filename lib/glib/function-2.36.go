// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Close is a wrapper around the C function g_close.
func Close(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_close(c_fd)
}

// GetNumProcessors is a wrapper around the C function g_get_num_processors.
func GetNumProcessors() {
	C.g_get_num_processors()
}

// Unsupported : g_unix_fd_add : unsupported parameter condition : no param type

// Unsupported : g_unix_fd_add_full : unsupported parameter condition : no param type

// Unsupported : g_unix_fd_source_new : unsupported parameter condition : no param type

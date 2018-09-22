// This is a generated file - DO NOT EDIT
// +build glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_base64_decode_inplace : unsupported parameter text : no param type

// Poll is a wrapper around the C function g_poll.
func Poll(fds *PollFD, nfds uint32, timeout int32) int32 {
	c_fds := fds.toC()

	c_nfds := (C.guint)(nfds)

	c_timeout := (C.gint)(timeout)

	retC := C.g_poll(c_fds, c_nfds, c_timeout)
	retGo := (int32)(retC)

	return retGo
}

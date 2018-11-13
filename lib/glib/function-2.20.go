// This is a generated file - DO NOT EDIT
// +build glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_base64_decode_inplace : unsupported parameter out_len : array length param out_len is pointer (gsize*)

// Polls @fds, as with the poll() system call, but portably. (On
// systems that don't have poll(), it is emulated using select().)
// This is used internally by #GMainContext, but it can be called
// directly if you need to block until a file descriptor is ready, but
// don't want to run the full main loop.
//
// Each element of @fds is a #GPollFD describing a single file
// descriptor to poll. The @fd field indicates the file descriptor,
// and the @events field indicates the events to poll for. On return,
// the @revents fields will be filled with the events that actually
// occurred.
//
// On POSIX systems, the file descriptors in @fds can be any sort of
// file descriptor, but the situation is much more complicated on
// Windows. If you need to use g_poll() in code that has to run on
// Windows, the easiest solution is to construct all of your
// #GPollFDs with g_io_channel_win32_make_pollfd().
/*

C function

g_poll
*/
func Poll(fds *PollFD, nfds uint32, timeout int32) int32 {
	c_fds := (*C.GPollFD)(C.NULL)
	if fds != nil {
		c_fds = (*C.GPollFD)(fds.ToC())
	}

	c_nfds := (C.guint)(nfds)

	c_timeout := (C.gint)(timeout)

	retC := C.g_poll(c_fds, c_nfds, c_timeout)
	retGo := (int32)(retC)

	return retGo
}
